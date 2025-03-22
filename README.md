# npminfo

`npminfo` is a command-line tool that allows users to fetch and display detailed information about any NPM (Node Package Manager) package. Built by Mycin, this tool makes it easier to retrieve package details and display them in a user-friendly table format.

## Features
- Fetches detailed information about any NPM package.
- Displays the information in a clear, tabular format.
- Allows users to specify a package name with a flag.

## Installation

You can install `npminfo` using `go install`:

```bash
go install github.com/mmycin/npmdc/cmd/npminfo@latest
```

Make sure your `$GOPATH/bin` directory is included in your system's `PATH` so that you can run the command from anywhere in your terminal.

## Usage

After installation, you can run `npminfo` from your terminal to get detailed information about an NPM package.

### Command Syntax

```bash
npminfo -p <package_name>
```

### Options

- `-p`, `--package` (required): The name of the NPM package you want to fetch information about.

### Example

```bash
npminfo -p express
```

This command will fetch and display detailed information about the `express` package.

### Output

The information will be displayed in a clear table format, providing details like:
- Name
- Version
- Description
- Author
- License
- Dependencies

## Contributing

We welcome contributions to `npminfo`. If you'd like to help improve the project, please feel free to fork the repository, create a branch, and submit a pull request.

### Steps for contributing:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and write tests.
4. Submit a pull request with a clear description of the changes.

## License

`npminfo` is licensed under the MIT License. See [LICENSE](./LICENSE) for more information.

## Acknowledgements

`npminfo` uses the following libraries:
- [cobra](https://github.com/spf13/cobra) for building the CLI
- [npmdc](https://github.com/mmycin/npmdc) for fetching and displaying NPM package info.

## Contact

For any questions or issues, feel free to reach out:
- GitHub: [mmycin](https://github.com/mmycin)

---
