# Echo AppRunner Project
Go言語のフレームワークEchoを用いて開発するTODOアプリケーションをAppRunnerにデプロイするプロジェクト

## アプリケーションのコンテナ化
Dockerfileではマルチステージビルドを用いた
goアプリケーションのコンテナ化と起動
```bash
make go-up
```

停止
```bash
make go-down
```

## 備忘録
- [開発手順書](./docs/dev-process.md)