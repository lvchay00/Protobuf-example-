package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	// 为 AllPerson 填充数据
	p1 := Person{
		Id:   *proto.Int32(1),
		Name: *proto.String("xieyanke"),
	}

	p2 := Person{
		Id:   2,
		Name: "gopher",
	}

	p3 := Person{
		Id:   3,
		Name: "gopher",
	}

	all_p := AllPerson{
		Per: []*Person{&p1, &p2, &p3},
	}

	// 对数据进行序列化
	data, err := proto.Marshal(&all_p)
	if err != nil {
		log.Fatalln("Mashal data error:", err)
	}
	fmt.Printf("%v\n", data)

	// 对已经序列化的数据进行反序列化
	var target AllPerson
	err = proto.Unmarshal(data, &target)
	if err != nil {
		log.Fatalln("UnMashal data error:", err)
	}

	fmt.Printf("%v", target.Per) // 打印person Name 的值进行反序列化验证
}
