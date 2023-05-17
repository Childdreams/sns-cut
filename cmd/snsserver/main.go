package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kevin-zx/wordcut/pkg/corpus/clear"
	"github.com/kevin-zx/wordcut/pkg/extractor"
	"git.fxt.cn/ec/sns-cut/pkg/sns_rpc"
	"git.fxt.cn/ec/sns-cut/pkg/xhsclear"
	"google.golang.org/grpc"
)

type server struct {
	sns_rpc.UnimplementedCuterServer
}

func (s *server) Cut(stream sns_rpc.Cuter_CutServer) error {
	contents := []string{}
	var err error
	for {
		r, err := stream.Recv()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		for _, c := range r.Content {
			contents = append(contents, c)
		}
	}
	raw := xhsclear.CleanAll(strings.Join(contents, "\n"))
	datar := clear.PureCorpusWithLettersAndDigitals(raw)
	raw = ""
	core := runtime.NumCPU()
	core = core / 2
	if core < 2 {
		core = 2
	}
	e := extractor.NewParallelBuilder(datar, 10, core)
	datar = nil
	ws := e.Extract()
	// 每次传输 10000条 否则会触发最大长度限制
	batchId := 0
	batchSize := 10000
	for {
		if batchId*batchSize+batchSize > len(ws) {
			err = stream.Send(convert(ws[batchId*batchSize:]))
			if err != nil {
				ws = nil
				return err
			}
			break
		}
		err = stream.Send(convert(ws[batchId*batchSize : batchId*batchSize+batchSize]))
		if err != nil {
			ws = nil
			return err
		}
		batchId++
	}
	e = nil
	ws = nil
	runtime.GC()
	return nil
}

// convert 转 extractor.Word 成为RPC对象
func convert(ws []*extractor.Word) *sns_rpc.SnsReply {
	sr := &sns_rpc.SnsReply{}
	for _, w := range ws {
		c := &sns_rpc.Cut{
			Word:  w.Word,
			Count: int32(w.Count),
			Freq:  float32(w.Freq),
			Poly:  float32(w.Poly),
			Flex:  float32(w.Flex),
		}
		sr.Cuts = append(sr.Cuts, c)
	}
	return sr
}

func main() {
	_ = godotenv.Load(".env")
	host := os.Getenv("RPC_HOST")
	port := os.Getenv("RPC_PORT")
	fmt.Printf("%s:%s\n", host, port)
	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	sns_rpc.RegisterCuterServer(s, &server{})
	log.Printf("server sns cuter service at %s:%s\n", host, port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
