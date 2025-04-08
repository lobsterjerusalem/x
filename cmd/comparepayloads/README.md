# install

```console
go install github.com/lobsterjerusalem/x/cmd/comparepayloads@latest
```

# usage

```console
$ comparehexpayloads
Usage: comparehexpayloads <file one> <file two>
Note: The files should be mostly the same this was made to compare hex of payloads that are one byte off in order to find length offsets.
The output is derived by running xxd on the payloads and only comparing the hex
```
