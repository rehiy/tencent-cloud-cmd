package exec

import (
	"fmt"
	"log"

	"github.com/rehiy/tencent-cloud-cmd/api"
)

func Caller() {

	flags := ParseFlag()
	CheckSecret(flags)

	if flags.Service == "" {
		log.Fatal("请设置 -service 参数，-h 查看帮助")
		return
	}

	if flags.Version == "" {
		log.Fatal("请设置 -version 参数，-h 查看帮助")
		return
	}

	if flags.Action == "" {
		log.Fatal("请设置 -action 参数，-h 查看帮助")
		return
	}

	res, err := api.Request(flags)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

}
