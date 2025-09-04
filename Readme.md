TinyGo is a Go Compiler For Small Places. It allows Go to start getting involved with microcontrollers, embedded systems.

## Install TinyGo

https://tinygo.org/docs/

Install: https://tinygo.org/getting-started/install/

### Windows

Quick Install via [Scoop](https://mrotaru.co.uk/blog/windows-package-manager-scoop/)

You can use Scoop to install TinyGo and dependencies.

If you haven’t installed Go already, you can do so with the following command:

> scoop install go

Followed by TinyGo itself:

> scoop install tinygo

Your $PATH environment variable will be updated via the scoop package. By default a shim is created in ~/scoop/shims/tinygo.

You can test that the installation was successful by running the version command which should display the version number:

> tinygo version
> tinygo version 0.36.0 windows/amd64 (using go version go1.24 and LLVM version 19.1.2)

### Linux

https://tinygo.org/getting-started/install/linux/

## Pico2w documentation

https://www.raspberrypi.com/documentation/microcontrollers/pico-series.html#pico-2-family

![417682171-89be49a1-b381-4cd1-b109-21f744a02b64](https://github.com/user-attachments/assets/37f50285-c34e-4ee5-be57-5646e404991e)

The Pinout Diagram [PDF](https://datasheets.raspberrypi.com/pico/Pico-2-Pinout.pdf)

Tinygo [Interface and Pin diagram](https://tinygo.org/docs/reference/microcontrollers/pico2-w/)

## Install Visual Studio Code support for TinyGo

https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo

See: https://tinygo.org/docs/guides/ide-integration/vscode/

https://pragmatik.tech/set-up-your-pico-with-tinygo-and-vscode

Install a Reset Button:

Connect a push button between the RUN pin (pin 30) on your Pico and ground.

## Use it

Create the directory

Open the command palette (Ctrl+Maj+P) and search for TinyGo target.
This set the right environment variables in the .vscode/settings.json of your workspace. If target is pico2-w:

``` bash
go mod init c/git/Golang/tinygo/tests/helloworld
go mod tidy
tinygo build -target=pico2-w -o helloworld.uf2 helloworld.go
```

Copy uf2 to disk.

> tinygo flash -target=pico2-w ./hello.go

To see the output of the program, you can use a serial monitor.
Test available ports with ports option:

``` bash
tinygo ports
Port                 ID        Boards
COM3                     :     
COM9                 2E8A:000A pico2-w
```

Display program output:

``` bash
tinygo monitor -target=pico2-w
Connected to COM9. Press Ctrl-C to exit.
hello world!
hello world!
...
```

Linux:

``` bash
tinygo flash -target=pico2-w ./main.go -port=/dev/ttyACM0
or
tinygo build -o main.uf2 -target=pico2-w ./main.go
```

## Programming

### Drivers

https://tinygo.org/docs/concepts/drivers/

https://github.com/tinygo-org/drivers

### Examples

https://tinygo.org/docs/tutorials/

https://github.com/tinygo-org/awesome-tinygo


## Resetting Flash memory

For Pico-series devices, BOOTSEL mode lives in read-only memory inside the RP2040 or RP2350 chip, and can’t be overwritten accidentally. No matter what, if you hold down the BOOTSEL button when you plug in your Pico, it will appear as a drive onto which you can drag a new UF2 file. There is no way to brick the board through software. However, there are some circumstances where you might want to make sure your flash memory is empty. You can do this by dragging and dropping a special UF2 binary onto your Pico when it is in mass storage mode.

    Download the UF2 file:
    
https://datasheets.raspberrypi.com/soft/flash_nuke.uf2





