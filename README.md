# wordcount

Small CLI to count lines, words, bytes, and Unicode characters in a text file.

## Requirements

- Go 1.22+

## Build

```bash
go build -o wordcount ./cmd/wordcount
```

## Usage

Exactly one mode flag is required:

```bash
./wordcount --lines sample.txt
./wordcount --words sample.txt
./wordcount --bytes sample.txt
./wordcount --chars sample.txt
```

Example:

```bash
$ echo -e "hello\nworld" > sample.txt
$ ./wordcount --lines sample.txt
2
$ ./wordcount --words sample.txt
2
```

## Tests

```bash
go test ./...
```

## License

MIT — see [LICENSE](LICENSE).
