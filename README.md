# golang-server-template

## ディレクトリ構成

```bash
.
├── Makefile  // 各種コマンドを登録
├── README.md  // ディレクトリ構築の説明や環境構築手順
├── app  // アプリケーション開発としてエンジニアが頑張るコード(サービスのコア部分)をこのディレクトリ配下で管理！！
│   ├── domain  // ドメインモデルの管理 サードパーティライブラリの依存はなるべく回避
│   │   ├── error
│   │   │   └── error.go
│   │   └── template  // ドメインモデル/値オブジェクト/リポジトリ(Interface)/ユースケース(Interface)
│   │       ├── id.go
│   │       ├── sample.go
│   │       ├── template.go
│   │       ├── template_list.go
│   │       ├── template_repository.go
│   │       ├── template_test.go
│   │       ├── template_usecase.go
│   │       └── time.go
│   ├── error
│   │   └── controller
│   │       └── error_translator.go
│   ├── template  // ドメインモデルに紐づく各要素の管理
│   │   ├── controller  // 入手におけるデータの管理
│   │   │   ├── template_controller_test.go
│   │   │   ├── template_controller.go
│   │   │   └── template_translator.go
│   │   ├── gateway  // リポジトリの実装
│   │   └── usecase  // ユースケースの実装
│   └── main.go // アプリ起動コード
├── docker-compose.yml // ローカル環境の構築やビルドイメージの作成
├── ent  // .entによるテーブル管理
│   ├── schema  // 自前管理
│   │   └── template.go  // 各テーブル設計
│   ├── ... // 自動生成
│   └── database.go // **必須** Databaseとの接続部実装を追加
├── go.mod
├── go.sum
├── migration  // .entを利用したマイグレーション -> 要ビルドイメージ
│   └── main.go
└── proto // protoファイル管理 -> 外部リポジトリへ切出
│   └── template
│       ├── template.pb.go  // 自動生成
│       ├── template_grpc.pb.go  // 自動生成
│       └── v1/template.proto
└── script // 各種スクリプトファイル
    ├── migration.sh  // ローカル環境でのmigration実行
    ├── run.sh // ローカル環境でのサーバー起動
    └── test.sh
```

## テスト
以下2点のテストコードを各ドメインモデルで用意する(必須)
- app/domain/${domain_name}_test.go  
    ビジネスロジックの正当性検証
- app/\${domain\_name}/controller/${domain_name}_controller_test.go  
    リクエスト単位の確認 ValueObjectのロジック正当性(バリデーション)もここで検証

## サーボパーティライブラリ
domain/usecase配下ではなるべく依存しないようにする(標準ライブラリで頑張る)

### サーバー
- [gRPC](https://grpc.io/)

### データベース
- [ent](https://entgo.io/)
- [pg](https://pkg.go.dev/github.com/lib/pq)

### ログ  
- [zap](https://pkg.go.dev/go.uber.org/zap)

### テスト
- [testify](https://pkg.go.dev/github.com/stretchr/testify)

## 参照
- [go-clean-arch](https://github.com/bxcodec/go-clean-arch)
