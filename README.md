# 実行方法

1. `git clone` したら、 `docker compose build` コマンドでイメージを作成してください。
2. `docker-compose up -d` コマンドでを実行し、ginのサーバーを起動してください。

## ディレクトリ構成
```mermaid
graph LR
A[gin-docker-sample] --> B[Dockerfile]
A --> C[README.md]
A --> D[docker-compose.yml]
A --> E[src]
E --> F[go.mod]
E --> G[go.sum]
E --> H[hello-world.go]
E --> I[main.go]
```
上記のディレクトリ構成を持つことを前提としています。`src` ディレクトリには、ginで使用するファイルが含まれています。