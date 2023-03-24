# qwe

qwe is a command-line interface (CLI) tool for ChatGPT, a large language model trained by OpenAI. With qwe, you can interact with ChatGPT directly from your terminal.

## Installation

To install qwe, you will need to have Go installed on your system. If you don't have Go installed, you can download it from the [official website](https://golang.org/dl/). Once you have Go installed, you can install qwe using the following command:

```
go get github.com/t0mk/qwe
```

This will download the qwe source code and install the `qwe` binary in your `$GOPATH/bin` directory. Make sure that your `$GOPATH/bin` directory is in your `$PATH` environment variable so that you can use the `qwe` command from anywhere.

## Usage

To use qwe, simply run the `qwe` command in your terminal followed by your query. For example:

```
qwe What is the capital of France\?
```

or 

```
qwe "What is the capital of France?"
```


This will send your query to ChatGPT and display the response in your terminal. You can also use qwe to have a conversation with ChatGPT by running the `qwe` command without any arguments:

```
qwe
```

This will start a conversation with ChatGPT. You can exit the conversation by typing `exit` or pressing `Ctrl+C`.

## Contributing

If you find a bug or have a feature request, please create an issue on the [GitHub repository](https://github.com/t0mk/qwe/issues). If you would like to contribute to qwe, please submit a pull request. 

## License

qwe is licensed under the [MIT License](https://github.com/t0mk/qwe/blob/main/LICENSE).
