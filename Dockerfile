# 	#ベースとなるイメージを取得
# FROM golang:latest
# #作成者情報
# MAINTAINER career<allofcareerapp@gmail.com>

# #カレントディレクトリは "/go"
# WORKDIR ./src/career

# #コンテナにローカルからコピー
# COPY ./ /go/src/career

# #ホットリロード用に環境変数をセットしておく
# ENV GO111MODULE=on

# #go言語におけるホットリロード用パッケージを取得
# RUN go get github.com/pilu/fresh

# #ポートの指定
# EXPOSE 8081


# #ホットリロードの開始
# CMD ["fresh"]

#===for k8s===

# # Use the official Golang image to create a build artifact.
#     # This is based on Debian and sets the GOPATH to /go.
#     # https://hub.docker.com/_/golang
#     FROM golang:latest as builder

# 	MAINTAINER career<allofcareerapp@gmail.com>

#     # Copy local code to the container image.
#     COPY . /usr/local/go/src/career-web-backend 
# 	WORKDIR /usr/local/go/src/career-web-backend
# 	RUN go get -d
#     # Build the binary. remove -mod=readonly
#     RUN CGO_ENABLED=0 go build -o career-web-backend . 

#     # Use the official Alpine image for a lean production container.
#     # https://hub.docker.com/_/alpine
#     # https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
# 	# alpineのシェルはbashではなくash
#     FROM alpine:latest
#     RUN apk add --no-cache ca-certificates
#     RUN apk add --no-cache bash

#     # Copy the binary to the production image from the builder stage.
#     COPY --from=builder /usr/local/go/src/career-web-backend/career-web-backend .
#     COPY --from=builder /usr/local/go/src/career-web-backend/config.ini .
#     COPY --from=builder /usr/local/go/src/career-web-backend/wait-for-it.sh .
#     RUN chmod 777 wait-for-it.sh

# 	# ENV PORT 8081
#     # Run the web service on container startup.
#     ENTRYPOINT [ "/bin/ash", "-c" ]
# 	# CMD [ ]

FROM golang:1.13 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server
RUN chmod 777 /server

# Run the web service on container startup.
CMD ["/server/career-web-backend","-runOption=PRO"]