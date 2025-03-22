TARGET = pico-w
SOURCE = blinky.go
BINARY = main.uf2
LDFLAGS = -size short -monitor

build:
	tinygo build -o $(BINARY) $(LDFLAGS) -target $(TARGET) $(SOURCE)

run:
	tinygo flash $(LDFLAGS) -target $(TARGET) $(SOURCE)
	

fmt:
	go fmt *.go

clean:
	-rm -f $(BINARY)