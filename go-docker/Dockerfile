# ベースとなるDockerイメージ指定
FROM golang:1.14

ENV GO111MODULE=on
ENV GOPATH /go

# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/app

# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/app

# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

# 軽量なウェブツールキットを導入
#RUN go get github.com/labstack/echo/v4

RUN apt-get update \
	&& apt-get install -y git \
	&& go get -v \
	github.com/urfave/cli/v2 \
	github.com/go-delve/delve/cmd/dlv@latest \
	github.com/rogpeppe/godef@latest \
	golang.org/x/tools/gopls@latest \
	golang.org/x/tools/cmd/goimports@latest \
	github.com/ramya-rao-a/go-outline@latest  \
	golang.org/x/net/context@latest \
	github.com/gorilla/mux@latest \
	github.com/gorilla/handlers@latest \
	google.golang.org/api/option@latest \
	firebase.google.com/go@latest \
	firebase.google.com/go/auth@latest \
	&& apt-get autoremove -y \
	&& apt-get clean -y \
	&& rm -rf /var/lib/apt/lists/*


RUN GO111MODULE=off go get github.com/oxequa/realize
RUN GO111MODULE=on