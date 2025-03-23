TARGET = pico-w
SOURCE = main.go
BINARY = main.uf2
#LDFLAGS = -size short -monitor -scheduler tasks -gc=conservative -size=full -stack-size=20kb
LDFLAGS = -size short -monitor

build:
	tinygo build -o $(BINARY) $(LDFLAGS) -target $(TARGET) $(SOURCE)

run:
	tinygo flash $(LDFLAGS) -target $(TARGET) $(SOURCE)
	

fmt:
	go fmt *.go

clean:
	-rm -f $(BINARY)