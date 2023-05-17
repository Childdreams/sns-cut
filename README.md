# sns rpc cutword

```shell
// grpc build
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.
/--go-grpc_opt=paths=source_relative pkg/sns/sns.proto
```

## 建议

每次发送到服务端的文章数量不宜超过1000篇

## client demo

```go

package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"git.fxt.cn/ec/sns-cut/pkg/sns_rpc"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	host := os.Getenv("RPC_HOST")
	port := os.Getenv("RPC_PORT")
	data, err := ioutil.ReadFile("data/content.data")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s:%s\n", host, port)
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := sns_rpc.NewCuterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()
	cl, err := c.Cut(ctx)
	if err != nil {
		panic(err)
	}
	datas := strings.Split(string(data), "\n")
	batch := 0
	batchSize := 1000
	for {
		fmt.Printf("batch:%d\n", batch)
		if batch*batchSize+batchSize > len(datas) {
			err = cl.Send(&sns_rpc.SnsRawContentReq{
				Content: datas[batch*batchSize:],
			})
			if err != nil {
				panic(err)
			}
			break

		}
		err = cl.Send(&sns_rpc.SnsRawContentReq{
			Content: datas[batch*batchSize : batch*batchSize+batchSize],
		})
		if err != nil {
			panic(err)
		}
		batch++

	}
	cl.CloseSend()
	var cuts []*sns_rpc.Cut
	for {
		r, err := cl.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		cuts = append(cuts, r.Cuts...)
	}

	for _, c := range cuts {
		fmt.Printf("%+v\n", c)
	}
}

```
