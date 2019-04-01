# Dependencies

```
go get fyne.io/fyne
go get fyne.io/fyne/cmd/fyne
```

# Windows

```
brew install mingw-w64
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build
fyne package -os windows -icon icon.png
```

# macOS

```
go build
fyne package -os darwin -icon icon.png
```
