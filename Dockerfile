# ---------- Build stage ----------
FROM golang:1.24-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# ---------- Run stage ----------
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=build /app/server /server

USER nonroot:nonroot
EXPOSE 8080
ENV PORT=8080

ENTRYPOINT ["/server"]
