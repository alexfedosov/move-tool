# Move tool

A simple CLI for slicing long samples into Ableton Note / Ableton Move presets

## TLDR; Quick start
1. Download app for your platform from [the latest release](https://github.com/alexfedosov/move-tool/releases/latest)
2. Make an audio sample contaning N parts (up to 16). Let's say you got a sample containing 16 slices of the equal length
3. Run in your terminal app ```move-tool slice -i my-long-sample.wav -n 16 -o /Users/alex/Desktop ```
4. Move tool will slice the original sample into 16 pieces and create a Move/Note preset with random name at `/Users/alex/Desktop/`
5. Star this repo
6. Keep reading if you want to contribute



# Read this if you want to contribute

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Ensure that you have [Go 1.22](https://golang.org/dl/) installed.
2. Clone the repo:

    ```sh
    git clone https://github.com/alexfedosov/move-tool.git
    ```

3. Navigate to the project directory:

    ```sh
    cd move-tool
    ```

4. Install dependencies:

    ```sh
    go mod tidy
    ```

## Usage

```sh
go run . slice -i <file path> -n <number of samples> -o <output directory>
```

### Example
Lets say you have prepared a long wav sample with up to 16 sounds of equal length. 
To slice it up into .ablpresetbundle you need to run the tool as

```sh
go run . slice -i my-sample.wav -n 16 -o /Users/alex/Desktop
```

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
