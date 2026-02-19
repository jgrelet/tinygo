TinyGo is a Go Compiler For Small Places. It allows Go to start getting involved with microcontrollers, embedded systems.

## Install TinyGo

https://tinygo.org/docs/

Install: https://tinygo.org/getting-started/install/

### Windows

Quick Install via [Scoop](https://mrotaru.co.uk/blog/windows-package-manager-scoop/)

You can use Scoop to install TinyGo and dependencies.

If you haven’t installed Go already, you can do so with the following command:

``` bash
scoop install go
```

Followed by TinyGo itself:

``` bash
scoop install tinygo
```

Your $PATH environment variable will be updated via the scoop package. By default a shim is created in ~/scoop/shims/tinygo.

You can test that the installation was successful by running the version command which should display the version number:

``` bash
tinygo version
tinygo version 0.36.0 windows/amd64 (using go version go1.24 and LLVM version 19.1.2)
```

### Linux

https://tinygo.org/getting-started/install/linux/

## Pico2w documentation

https://www.raspberrypi.com/documentation/microcontrollers/pico-series.html#pico-2-family

![417682171-89be49a1-b381-4cd1-b109-21f744a02b64](https://github.com/user-attachments/assets/37f50285-c34e-4ee5-be57-5646e404991e)

The Pinout Diagram [PDF](https://datasheets.raspberrypi.com/pico/Pico-2-Pinout.pdf) and Raspberry Pi Pico and Pico W Pinout Guide: [GPIOs Explained](https://randomnerdtutorials.com/raspberry-pi-pico-w-pinout-gpios)

Tinygo [Interface and Pin diagram](https://tinygo.org/docs/reference/microcontrollers/pico2-w/)

## Install Visual Studio Code support for TinyGo

https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo

https://tinygo.org/docs/guides/ide-integration/vscode/

https://pragmatik.tech/set-up-your-pico-with-tinygo-and-vscode

To use it, click on the TinyGo item in the status bar at the bottom of the screen and select a target. You can also open the command palette, search for the TinyGo target, and select your target. In our example, a pico-w or pico2-w board.

For example, it may set the following configuration to work with the Raspberry Pico2-w:

``` bash
{
    "go.toolsEnvVars": {
        "GOOS": "linux",
        "GOARCH": "arm",
        "GOROOT": "C:\\Users\\yourName\\AppData\\Local\\tinygo\\goroot-fe15d27687627bb9dacd906a845880b8406b8c133b7c087813b7f5767641f01a",
        "GOFLAGS": "-tags=cortexm,baremetal,linux,arm,rp2350,rp,pico2,pico2-w,cyw43439,tinygo,purego,osusergo,math_big_pure_go,gc.conservative,scheduler.cores,serial.usb"
    }
}
```

## Install a Reset Button

Connect a push button between the RUN pin (pin 30) on your Pico and ground.

## Use it

Create the directory

Open the command palette (Ctrl+Maj+P) and search for TinyGo target.
This set the right environment variables in the .vscode/settings.json of your workspace. If target is pico2-w:

``` bash
go mod init <your path>/tinygo/tests/helloworld
go mod tidy
tinygo build -target=pico2-w -o helloworld.uf2 helloworld.go
```

Copy uf2 to disk.

``` bash
tinygo flash -target=pico2-w ./hello.go
```

To view the program output, you can use a serial monitor.
Test the available ports with the ports option:

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
It is also possible to flash the program and launch monitor mode to display the result with the following option:
``` bash
tinygo flash -monitor -target=pico2-w ./hello.go
```

Use the -size option to reduce the size of the binary
``` bash
tinygo flash -monitor -size short -target=pico2-w ./hello.go
```

Linux:

``` bash
tinygo flash -target=pico2-w ./main.go -port=/dev/ttyACM0
or
tinygo build -o main.uf2 -target=pico2-w ./main.go
```

## Tips

When using monitor mode, it is essential to set a delay of 1 to 2 seconds at the start of the main function, otherwise the first outputs will not be visible on the serial device (USB).

``` go
import "time"

    time.Sleep(time.Second)
``` 

The GOROOT variable defines a cache directory under <user>/AppData/Local/tinygo/goroot-2aa6dda4...
When inconsistent compilation issues occur, you can clear this cache using the clean option.
Then reselect the target under VSC and restart it.
Check environment variables GOROOT with go env or tinygo env command. [Understanding Go Environment Variables](https://medium.com/@dilandashintha/understanding-go-env-1109bcba9a9c) :

``` bash
go env GOROOT
```

When you deploy your modules after testing them on GitHub and changing the path names, the following error may occur:

``` bash
go: github.com/<path>/project: parsing go.mod:
    module declares its path as: github.com/<path>/test
    but was required as: github.com/<path>/project
```

Even when I run <code>go get -u github.com/\<path\>/project</code> or
<code>go mod tidy</code>

In this case, execute:

``` bash
go env -w GOPROXY="direct"
go mod tidy
```

## Programming

### Drivers

https://tinygo.org/docs/concepts/drivers/

https://github.com/tinygo-org/drivers

### Examples

https://tinygo.org/docs/tutorials/

https://github.com/tinygo-org/awesome-tinygo


## Resetting Flash memory

For Pico-series devices, BOOTSEL mode lives in read-only memory inside the RP2040 or RP2350 chip, and can’t be overwritten accidentally. No matter what, if you hold down the BOOTSEL button when you plug in your Pico, it will appear as a drive onto which you can drag a new UF2 file. There is no way to brick the board through software. However, there are some circumstances where you might want to make sure your flash memory is empty. You can do this by dragging and dropping a special UF2 flash_nuke.uf2 binary onto your Pico when it is in mass storage mode.

## Download the UF2 file:
    
https://datasheets.raspberrypi.com/soft/flash_nuke.uf2





