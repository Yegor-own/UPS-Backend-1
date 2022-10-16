FROM golang:alpine

WORKDIR /go/src/backend

COPY . .

RUN go get -d -v ./
# RUN go install -v ./

RUN go build .

EXPOSE 3000

CMD [ "./ups" ]