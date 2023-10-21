<h1 align="center">Vertex</h1>
<h3 align="center">A lightweight super-fast open-source game engine built with Go and OpenGL!</h3>

<p align="center">
  <a href="https://github.com/dimkauzh/vertex"><img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/dimkauzh/vertex"></a>
  <a href="https://github.com/dimkauzh/vertex"><img alt="GitHub license" src="https://img.shields.io/github/license/dimkauzh/vertex"></a>
  <a href="https://github.com/dimkauzh/vertex"><img alt="Lines of code" src="https://tokei.rs/b1/github/dimkauzh/vertex?category=lines"></a>
</p>

> **Please note that vertex is currently under high maintenance and is not production-ready.** The project is actively being developed and improved, which is why there is only a `dev` branch available. Its currently in Pre-Alpha status


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
