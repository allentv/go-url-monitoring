FROM golang:1.18 as base

WORKDIR /app

# Install UPX compressor
# Ref: https://github.com/upx/upx

RUN apt-get update && apt-get install xz-utils
RUN wget https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz
RUN tar -xvf upx-* && cp upx-*/upx ./upx

# Copy Go files
COPY go.* ./
RUN go mod download
COPY . .

# Build the application and strip debug info
RUN go build -ldflags="-s -w" -o main cmd/app/main.go

# Compress the binary
RUN ./upx --best --lzma ./main


FROM gcr.io/distroless/static

WORKDIR /app

COPY --from=base /app/main /app/main

EXPOSE 9000

CMD ["./main"]
