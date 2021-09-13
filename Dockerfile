FROM golang:1.17

WORKDIR /projetos/go/src/digital-bank

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /digital-bank

EXPOSE 4500

CMD [ "/digital-bank" ]