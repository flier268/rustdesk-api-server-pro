# Build the backend
FROM golang:alpine AS golang
WORKDIR /backend
COPY ./backend .
RUN go build

# Build the frontend
FROM node:20-alpine AS node
WORKDIR /frontend
COPY ./soybean-admin .
RUN rm -rf node_modules
RUN npm install -g pnpm
RUN pnpm i && pnpm build

# Build the init program
FROM golang:alpine AS init-builder
WORKDIR /app
COPY ./docker/init.go .
RUN go mod init init
RUN go build -o init

# Final stage
FROM gcr.io/distroless/base-debian12
ENV ADMIN_USER=
ENV ADMIN_PASS=
WORKDIR /app
COPY --from=golang /backend/rustdesk-api-server-pro .
COPY --from=golang /backend/server.yaml .
COPY --from=node /frontend/dist ./dist
COPY --from=init-builder /app/init .
EXPOSE 8080
CMD ["/app/init"]
