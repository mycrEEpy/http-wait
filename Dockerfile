FROM golang:1.8

LABEL maintainer="Tobias Germer"

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080

CMD ["go-wrapper", "run"] 
