# Dependencies

```
go get fyne.io/fyne
```

# Windows

```
brew install mingw-w64
```

32 bit:
```
CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc go build
```

64 bit:
```
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build
```

# macOS

```
go build
```
