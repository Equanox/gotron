UNAME := $(shell uname)

EXEC = ./example


#https://stackoverflow.com/questions/4058840/makefile-that-distincts-between-windows-and-unix-like-systems
# ifdef OS 
# 	EXEC = example.exe
# endif

run:
	@cd example && go build && $(EXEC)

install-ui:
	@cd ui && npm run install

build-ui:
	@cd ui && npm run build

builder:
	@cd cmd/gotron-builder && go build -o ../../gotron-builder && cd ../../ && ./gotron-builder

install-builder:
	@cd cmd/gotron-builder && go install

clean:
	@-rm -r .gotron
	@-rm -r .gotron-builder