# go-textwriter

Writer for DOS line code

## Usage

```
f, _ := os.Create("newfile.txt")
defer f.Close()
tw := textwriter.NewWriter(f)
defer tw.Close()
tw.Write([]byte("hello\n"))
```

## Installation

```
$ go get github.com/mattn/go-textwriter
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
