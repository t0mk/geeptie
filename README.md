# Geeptie

Geeptie is a command-line interface (CLI) tool for ChatGPT, a large language model trained by OpenAI. With Geeptie, you can interact with ChatGPT directly from your terminal.

## Installation

To install Geeptie, you will need to have Go installed on your system. If you don't have Go installed, you can download it from the [official website](https://golang.org/dl/). Once you have Go installed, you can install Geeptie using the following command:

```
go get github.com/t0mk/geeptie
```

This will download the Geeptie source code and install the `geeptie` binary in your `$GOPATH/bin` directory. Make sure that your `$GOPATH/bin` directory is in your `$PATH` environment variable so that you can use the `geeptie` command from anywhere.

## Usage

To use Geeptie, simply run the `geeptie` command in your terminal followed by your query. For example:

```
geeptie What is the capital of France\?
```

or 

```
geeptie "What is the capital of France?"
```


This will send your query to ChatGPT and display the response in your terminal. You can also use Geeptie to have a conversation with ChatGPT by running the `geeptie` command without any arguments:

```
geeptie
```

This will start a conversation with ChatGPT. You can exit the conversation by typing `exit` or pressing `Ctrl+C`.

## Contributing

If you find a bug or have a feature request, please create an issue on the [GitHub repository](https://github.com/t0mk/geeptie/issues). If you would like to contribute to Geeptie, please submit a pull request. 

## License

Geeptie is licensed under the [MIT License](https://github.com/t0mk/geeptie/blob/main/LICENSE).
