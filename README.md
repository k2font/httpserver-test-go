# httpserver-test-go
## Features
- 基本的なハンドラのテスト
- HTTPリクエストの内容を解析し、いい感じのレスポンスを返す

## Requirements
- gorilla/mux@1.8.0
  - https://github.com/gorilla/mux

## Usage
- `$ go run path/to/main.go`
- http://localhost/workspace/testworkspace/user/k2font などにアクセスすると、以下のようなメッセージを出力します。
  - `あなたの所属ワークスペース名は testworkspace で、あなたは k2font さんです。`