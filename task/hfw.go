package exec

import (
	"log"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/rehiy/tencent-cloud-cmd/api"
	"github.com/rehiy/tencent-cloud-cmd/exec"
)

func SetLighthouseFirewalls() {

	flags := exec.ParseFlag()
	exec.CheckSecret(flags)

	if flags.Region == "" {
		log.Fatal("请设置 -region 参数，-h 查看帮助")
		return
	}

	ready := 1
	limit := 100

	for i := 0; ; i++ {

		log.Println("正在获取第", i, "页实例信息")

		res, err := api.Request(&api.Params{
			Service:   "lighthouse",
			Version:   "2020-03-24",
			Action:    "DescribeInstances",
			Region:    flags.Region,
			Payload:   `{"Offset": ` + strconv.Itoa(limit*i) + `, "Limit": ` + strconv.Itoa(limit) + `}`,
			SecretId:  flags.SecretId,
			SecretKey: flags.SecretKey,
		})

		if err != nil {
			log.Println(err)
			continue
		}

		obj, err := simplejson.NewJson([]byte(res))

		if err != nil {
			log.Println(err)
			continue
		}

		total := obj.GetPath("Response", "TotalCount").MustInt()
		instanceSet := obj.GetPath("Response", "InstanceSet").MustArray()

		if len(instanceSet) == 0 {
			log.Println("未找到实例")
			break
		}

		for _, item := range instanceSet {
			id := item.(map[string]any)["InstanceId"].(string)
			log.Println("正在设置实例 ", id, " 防火墙规则（", ready, "/", total, "）")
			SetLighthouseFirewall(flags, id)
			ready++
		}
	}

}

func SetLighthouseFirewall(flags *api.Params, id string) {

	res, err := api.Request(&api.Params{
		Service:   "lighthouse",
		Version:   "2020-03-24",
		Action:    "CreateFirewallRules",
		Region:    flags.Region,
		Payload:   `{"InstanceId":"` + id + `","FirewallRules":` + flags.Payload + `}`,
		SecretId:  flags.SecretId,
		SecretKey: flags.SecretKey,
	})

	if err != nil {
		log.Println(err)
	} else {
		log.Println(res)
	}

}
