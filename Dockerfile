##
## Build
##

FROM golang:1.16-buster AS build

# timezone 설정
# ARG DEBIAN_FRONTEND=noninteractive
# ENV TZ=Asia/Seoul
# RUN apt-get install -y tzdata

WORKDIR /go/src/app
COPY . .

RUN go mod download

RUN go build -o go-docker-boilerplate

##
## Deploy
##

FROM gcr.io/distroless/base-debian10:debug
WORKDIR /app/
COPY --from=build /go/src/app /app

#EXPOSE 3201

USER nonroot:nonroot
ENV APP_ENV=production

ENTRYPOINT ["/app/go-docker-boilerplate"]
