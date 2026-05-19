setup:
    bun install
    go mod download
    go install github.com/a-h/templ/cmd/templ@latest
    go install github.com/air-verse/air@latest

dev:
    templ generate --watch &
    bun dev &
    air
