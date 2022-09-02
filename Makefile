TARGET = gobackendversion

build:
	go build -o $(TARGET) cmd/$(TARGET).go

clean:
	go clean
	rm -f $(TARGET)
	find . -name "*~" -exec rm -f {} \;
