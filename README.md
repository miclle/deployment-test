# deployment-test

用于验证 LAS 一键部署流程的最小 Go HTTP 服务。

## 本地运行

```bash
go run .
```

服务默认监听 `4173` 端口，也可以通过 `PORT` 环境变量覆盖。

## 部署配置

- 配置文件：`.qiniu/deploy.yaml`
- 工作目录：仓库根目录
- 安装命令：下载并校验 Go 1.24.6 工具链，然后执行 `go build -o deployment-test .`
- 启动命令：`./deployment-test`
- 端口：`4173`
