FROM golang:alpine AS Builder
ARG ENV
LABEL maintainer="Abdulsamet İleri <abdulsamet.ileri@ceng.deu.edu.tr>"
WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go test --tags $ENV ./...
RUN CGO_ENABLED=0 GOOS=linux go build --tags $ENV -o main main.go

FROM scratch
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"]