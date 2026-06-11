# wordcount

Small CLI to count lines, words, bytes, Unicode characters, and the longest line in a text file.

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
./wordcount --longest sample.txt
```

The `--longest` flag prints the length of the longest line in Unicode code points (not bytes).

Example:

```bash
$ cat sample.txt
hello
world wide
$ ./wordcount --longest sample.txt
10
```

## Tests

```bash
go test ./...
```

## License

MIT — see [LICENSE](LICENSE).
