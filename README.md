# 腾讯云跨平台命令行工具

这是一个无第三方库依赖的腾讯云API接入实现，你可以作为 Library 或 CLI 程序使用。

## Library 用法

```go
import "github.com/rehiy/tencent-cloud-cmd/api"

func main() {
    res, err := api.Request(&api.Params{
      // 参数表
    })
}
```

## Command 用法

```shell
export TENCENTCLOUD_SECRET_ID=xxxx
export TENCENTCLOUD_SECRET_KEY=yyyy

tcmd --service cvm --version 2017-03-12 --action DescribeRegions --region ap-guangzhou --payload "{}"
```

### Command 参数

```shell
  -action string
        执行动作
  -payload string
        JSON数据
  -region string
        地域
  -service string
        服务名
  -version string
        服务版本
```
