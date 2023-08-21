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