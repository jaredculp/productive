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
CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc go build -o productive_0.1.0_macos
```

64 bit:
```
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -ldflags -H=windowsgui -o productive_0.1.0_windows-amd64.exe
```

# macOS

```
go build
```
