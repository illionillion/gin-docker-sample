FROM golang:latest

WORKDIR /app

# ファイルをコンテナのルートにコピー
COPY src/go.mod . 
COPY src/go.sum .
# コマンドでmodファイルにあるパッケージをダウンロード
RUN go mod download
# COPY src .

# CMD [ "go", "run", "main.go" ]
# docker build -t mygo .