
.PHONY: clean build

build: clean
	go build -o flsvc -C ./build/ ../main.go
	cp -fr ./build/flsvc ./flsvc

clean:
	rm -f ./build/flsvc
	rm -f ./flsvc
	
