TARGET = gobackend

build:
	go build -o $(TARGET)

clean:
	go clean
	rm -f $(TARGET)
	rm -f *~ 
