# vertex-engine

A super-fast open-source game engine built with Go and OpenGL.

## Setup for Developers

### Needed packages
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

### Setup
To set up the development environment, follow these steps:
Open a terminal.
Navigate to the root directory of the vertex-engine project.

Then run:
```bash
make run
```
## Running Examples
After setting up the development environment, you can run the examples.

Open a terminal.
Navigate to the root directory of the vertex-engine project.
To run the default example, use the following command:
```bash
make run
```
This will run the default example.

To specify a different example, use the EXAMPLE variable, like this(currently not availible):

```bash
make run EXAMPLE=2
```
