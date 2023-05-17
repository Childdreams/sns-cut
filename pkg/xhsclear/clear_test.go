package xhsclear

import (
	"reflect"
	"testing"
)

func TestEmojiFilter(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "1",
			args: args{
				content: []byte(`才会心甘情愿为女生的DR梦买单。[萌萌哒R][萌萌哒R]一直以来他是一个很温,笑着说还热快喝。[害羞R]我到`),
			},
			want: []byte(`才会心甘情愿为女生的DR梦买单。一直以来他是一个很温,笑着说还热快喝。我到`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EmojiFilter(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmojiFilter() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestTagFilter(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "1",
			args: args{
				content: []byte(`脏粉鱼骨纹法兰绒#兰州西服##兰州西装##兰州定制##兰州定制西服##兰州定制西装##兰州西服定制##兰州西装定制##兰州西装高级定制##兰州商务西装定制##兰州商务西服定制##兰州婚礼##兰州婚礼新郎西服##兰州婚礼新郎西装##兰州量体定制##兰州主持人形象##兰州主持人服饰##兰州艺考生定制##兰州旗袍定制##兰州男士服装##兰州婚礼##兰州婚庆##兰州婚纱##兰州结婚##兰州专业西服定制#全网最可爱发型，小姐姐做完发型年轻10岁`),
			},
			want: []byte(`脏粉鱼骨纹法兰绒，小姐姐做完发型年轻10岁`),
		}, {
			name: "2",
			args: args{
				content: []byte(`悉尼水果蛋糕 悉尼蛋糕店 悉尼蛋糕配送
Crave Cake 水果蛋糕 ❤️
#悉尼复古蛋糕#悉尼甜品#悉尼蛋糕#悉尼生日蛋糕#悉尼蛋糕配送#悉尼甜品配送#悉尼婚礼蛋糕#悉尼蛋糕#甜品店#悉尼婚礼甜品#悉尼翻糖蛋糕#悉尼蛋糕订制#悉尼生日礼物#悉尼慕斯蛋糕#慕斯蛋糕#爱心蛋糕#悉尼婚礼蛋糕#悉尼蛋糕店#悉尼网红蛋糕#悉尼女士蛋糕#网红蛋糕#悉尼蛋糕店#韩式裱花#悉尼订制饼干#悉尼饼干订制#悉尼慕斯蛋糕#悉尼小仙女蛋糕#奥利奥蛋糕#悉尼奥利奥蛋糕#悉尼男士蛋糕#悉尼女生蛋糕#悉尼女神蛋糕#悉尼裱花蛋糕#悉尼魔方慕斯蛋糕#悉尼外卖#悉尼爱心蛋糕#悉尼手绘蛋糕#悉尼男孩蛋糕#悉尼下午茶#悉尼甜品外卖#悉尼外卖#悉尼甜品店##悉尼柠檬蛋糕#悉尼草莓花束#悉尼草莓蛋糕#悉尼奥利奥饼#悉尼甜品外卖#悉尼甜品外卖#悉尼皇冠蛋糕#悉尼女神蛋糕#悉尼蛋糕外送
#悉尼复古蛋糕
#悉尼复古蛋糕
#悉尼复古蛋糕
配饰|就想戴着你❤️`),
			},
			want: []byte(`悉尼水果蛋糕 悉尼蛋糕店 悉尼蛋糕配送
Crave Cake 水果蛋糕 ❤️




配饰|就想戴着你❤️`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TagFilter(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagFilter() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestAtFilter(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "1",
			args: args{
				content: []byte(`文艺@日常薯  @小红叔  @穿搭薯  @视频薯 超显白@郑州哪里剪发好@郑州哪里染发好@郑州哪里剪短发好@郑州哪里烫发好@郑州短发，郑州长发，冷棕色，蓝黑色，闷青色，茶棕色，百搭色，郑州浩天郑州MAX`),
			},
			want: []byte(`文艺       超显白，郑州长发，冷棕色，蓝黑色，闷青色，茶棕色，百搭色，郑州浩天郑州MAX`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AtFilter(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagFilter() = \n%v, want \n%v", string(got), string(tt.want))
			}
		})
	}
}
