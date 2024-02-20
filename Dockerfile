FROM docker.arvancloud.ir/golang

WORKDIR /app

COPY . .

RUN go build -o vending-machine .

EXPOSE 8080

ENTRYPOINT  ["./vending-machine"]