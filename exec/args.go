package exec

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rehiy/tencent-cloud-cmd/api"
)

func ParseFlag() *api.Params {

	var args = api.Params{}

	flag.StringVar(&args.Service, "service", "", "服务名")
	flag.StringVar(&args.Version, "version", "", "服务版本")
	flag.StringVar(&args.Action, "action", "", "执行动作")
	flag.StringVar(&args.Region, "region", "", "地域")
	flag.StringVar(&args.Payload, "payload", "", "结构化数据")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, readme)
		flag.PrintDefaults()
	}

	flag.Parse()

	return &args

}

func CheckSecret(args *api.Params) {

	args.SecretId, _ = os.LookupEnv("TENCENTCLOUD_SECRET_ID")
	args.SecretKey, _ = os.LookupEnv("TENCENTCLOUD_SECRET_KEY")

	if args.SecretId == "" || args.SecretKey == "" {
		log.Fatal("请设置环境变量 TENCENTCLOUD_SECRET_ID 和 TENCENTCLOUD_SECRET_KEY")
	}

}

const readme = `
使用方法:

tcmd --service cvm --version 2017-03-12 --action DescribeRegions --region ap-guangzhou --payload "{}"

选项说明:

`
