FROM golang:1.14-stretch

# コンテナ内に作業ディレクトリを作成
RUN mkdir /quizen
# コンテナログイン時のディレクトリ指定
WORKDIR /quizen
# コンテナにソースをコピー
COPY . /quizen
# ポートを公開
EXPOSE 8080
# コンテナにコピーしたソースをビルド
RUN make build_linux
# Serverを起動
CMD ./build/quizen