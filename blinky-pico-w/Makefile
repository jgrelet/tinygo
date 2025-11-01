TARGET = pico-w
SOURCE = blinky.go
BINARY = flash.uf2
LDFLAGS = -size short -monitor

build:
	tinygo build -o $(BINARY) $(LDFLAGS) -target $(TARGET) $(SOURCE)

flash:
	tinygo flash $(LDFLAGS) -target $(TARGET) $(SOURCE)

monitor: 
	tinygo monitor -target=$(TARGET)		

fmt:
	go fmt *.go

clean:
	-rm -f $(BINARY)