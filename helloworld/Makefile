TARGET = pico2-w
SOURCE = helloworld.go
BINARY = main.uf2
LDFLAGS = -size short

all: build flash run

build:
	tinygo build -o $(BINARY) $(LDFLAGS) -target $(TARGET) $(SOURCE)

flash:
	tinygo flash $(LDFLAGS) -target $(TARGET) $(SOURCE)
	
run: 
	tinygo monitor -target=$(TARGET)

fmt:
	go fmt *.go

clean:
	-rm -f $(BINARY)