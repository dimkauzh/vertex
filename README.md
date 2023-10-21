# vertex-engine

A lightweight super-fast open-source game engine built with Go and OpenGL.

## Building

### Prerequisites
#### Ubuntu/Debian-like Linux Distributions
On Ubuntu/Debian-like Linux distributions, you need to install the following packages:

```bash
sudo apt-get install libgl1-mesa-dev xorg-dev
```
#### macOS
On macOS, you need to have either Xcode or Command Line Tools for Xcode installed. You can install the Command Line Tools by running:

```bash
xcode-select --install
```

#### Windows
For Windows, you will need a C compiler. We recommend using [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/).

### Getting the Go package
Vertex is an ordinary go package, so you can get it using this command:
```bash
go get github.com/dimkauzh/vertex@latest
```

### Running Examples
1. Open a terminal.
2. Clone the repo
3. Navigate to the root directory of the vertex-engine project.
4. To run the example, use the following command:
```bash
make example
```
## License
This project is licensed under the GPLv3 License - see the LICENSE file for details.
