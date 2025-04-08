# to install

```console
go install github.com/lobsterjerusalem/x/cmd/hexdecoder@latest
```

# usage

```console
$ hexdecoder
Usage: hexdecoder <hex or hex filepath> <optional: output filepath prefix>

$ hexdecoder oracle_parsed some
saved as binary to some.bin
saved as hex str to some.hexstr
$ head -n 1 some.hexstr
"\x2e\x4e\x45\x54\x01\x00\x00\x00\x00\x00\x9e\x02\x00\x00\x06\x00\x01"+
```
