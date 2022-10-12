FROM golang:1.19-alpine3.16

##buat folder APP
RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh file ke app
ADD . /app

##buat executeable
RUN go run .

##jalankan executeable
# CMD ["/app/main"]