
## API定義書作成
StoplightでTODOアプリケーション用のAPI定義書作成

## コード生成
oapi-codegenをインストールしてコード生成
```bash
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
oapi-codegen -package petstore petstore-expanded.yaml > petstore.gen.go
```

以下のパッケージをインストールする
```text
"github.com/deepmap/oapi-codegen/pkg/runtime"
"github.com/getkin/kin-openapi/openapi3"
"github.com/labstack/echo/v4"
```

## AWS CopilotによるAppRunnerへのデプロイ
Applicationの初期化
```bash
copilot app init
```

Environmentの初期化とデプロイ
```bash
copilot env init
copilot env deploy --name dev
```

Serviceの初期化とデプロイ
```bash
copilot svc init
copilot svc deploy --name {service-name} --env dev
```

デプロイに成功するとエンドポイントが表示される
```bash
Recommended follow-up action:
  - You can access your service at https://yhjp2pr4z3.ap-northeast-1.awsapprunner.com over the internet.
```


# 参考
- [AWS Copilot CLI](https://aws.github.io/copilot-cli/ja/docs/overview/)
