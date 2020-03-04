# ベースとなるDockerイメージ指定
FROM golang:latest

# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/work
# 軽量なウェブツールキットを導入
RUN go get golang.org/x/net/context
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get google.golang.org/api/option
RUN go get github.com/derekparker/delve/cmd/dlv
RUN go get firebase.google.com/go
RUN go get firebase.google.com/go/auth

