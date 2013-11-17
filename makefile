export CGO_LDFLAGS=-lprocps 

build: 
	$(info ****** Starting Compiling main application ******)
	go build

run: build
	./raspGo

execute:
	./raspGo

clean:
	-rm raspGo