# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/work
# 軽量なウェブツールキットを導入
RUN go get github.com/gorilla/mux
RUN go run github.com/gorilla/handlers
RUN go get firebase.google.com/go
RUN go get google.golang.org/api/option
RUN go get firebase.google.com/go/auth
RUN go get golang.org/x/net/context

