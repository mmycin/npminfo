# Nim Package Directory Web Crawler

This project is a web crawler that scrapes the Nim package directory and outputs the data into a CSV file. It uses `BeautifulSoup` for web scraping, `Polars` for data manipulation, and `UV` (a Rust-based package manager) to run the project. The resulting CSV file is saved in the `data` folder.

## Features

- Scrapes the Nim package directory.
- Outputs the results as a CSV file.
- Written in Python using `BeautifulSoup`, `Polars`, and `UV` for asynchronous execution.
- Easily run the project with `make`.

## Requirements

- Python 3.8+
- `UV` package manager (alternative to `pip`)
- Required Python packages:
  - `beautifulsoup4`
  - `polars`
  - `requests`
  - `python-dotenv`

## Installation

### 1. Clone the repository:

```bash
git clone https://github.com/mmycin/npmdc
cd npmdc
```

### 2. Install dependencies with `UV`:

If you have `UV` installed, you can use it to install dependencies by running:

```bash
uv install
```

Alternatively, you can manually install the required Python packages:

```bash
pip install beautifulsoup4 polars requests python-dotenv
```

### 3. Set up the `.env` file:

In the root directory, create a `.env` file to store the necessary environment variables. The `.env` file should include:

```
URL=<URL_of_the_Nim_package_directory>
```

Make sure to replace `<URL_of_the_Nim_package_directory>` with the actual URL of the Nim package directory you want to crawl.

## Usage

### 1. Run with `make`:

If you have `make` installed, you can use the following command to run the crawler:

```bash
make
```

The `make` command will execute the crawler, scrape the Nim package directory, and output the data into a CSV file in the `data` folder.

### 2. Running manually (without `make`):

If you prefer to run the crawler manually, simply run the Python script:

```bash
python crawler.py
```

This will start the web crawler and save the CSV output in the `data` folder.

## Output

The CSV file will be saved in the `data` folder with the name `nim_package_data.csv`. The columns in the CSV file will represent the data extracted from the Nim package directory.

### Example structure:

- `Name`: The name of the package.
- `Description`: A brief description of the package.
- `Link`: A link to the package’s page.
  
## Troubleshooting

- If you encounter any issues related to dependencies, ensure that you have installed the correct versions of the required packages.
- If the `.env` file is not set up correctly, the crawler will fail to retrieve the Nim package directory URL.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
