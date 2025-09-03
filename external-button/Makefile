TARGET = pico-w
SOURCE = main.go
BINARY = main.uf2
LDFLAGS = -size short

build:
	tinygo build -o $(BINARY) $(LDFLAGS) -target $(TARGET) $(SOURCE)

run:
	tinygo flash $(LDFLAGS) -target $(TARGET) $(SOURCE)
	
monitor: 
	tinygo monitor -target=$(TARGET)	
	
fmt:
	go fmt *.go

clean:
	-rm -f $(BINARY)