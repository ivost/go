SR4_PREFIX=aarch64-linux-gnu-
SR4_IP=192.168.6.100
CC4 = $(SR4_PREFIX)gcc
GO_FLAGS = -ldflags '-extldflags "-static"'
NAME = server
TARGET = /tmp/$(NAME)

.PHONY: dist clean client server

# local build
dist: clean
	go run $(NAME).go

help:
	@echo Usage:
	@echo make clean
	@echo make sr
	@echo make bench
	
# SR4 build	
sr:
	CGO=0 GOARCH=arm64 GOOS=linux go build -o $(TARGET) $(GO_FLAGS) $(NAME).go
	ls -ltrh $(TARGET)

scp:
	scp $(TARGET) root@$(SR4_IP):/data/bench/$(NAME)

clean:
	rm -f $(NAME) $(TARGET)

bench:
	#baton -u http://$(SR4_IP):8080/foo -c 4 -r 10000 
	baton -u http://localhost:8080/foo -c 4 -r 10000 
 
mem:
	go tool pprof --http localhost:6060 server mem.prof &

cpu:
	go tool pprof --http localhost:6061 server cpu.prof &

quit:
	curl -i localhost:8080/quit

client:
	go test ../client/client_test.go