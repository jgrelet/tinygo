TARGET = pico-w
SOURCE = main.go
BINARY = main.uf2
LDFLAGS = -size short -monitor
CURRENT_TIME = "-X main.buildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'`"

build:
	tinygo build -o $(BINARY) $(LDFLAGS) -target $(TARGET) $(SOURCE)

flash:
	tinygo flash $(LDFLAGS) -target $(TARGET) -ldflags=$(CURRENT_TIME) $(SOURCE)

monitor: 
	tinygo monitor -target=$(TARGET)	

fmt:
	go fmt *.go

clean:
	-rm -f $(BINARY)