


GOARCH=wasm GOOS=js go build -o  ./web-app/app.wasm # Build app.wasm:
go run main.go # Build and generate static website
git add . && git commit -m "Update" && git push # Push to github
