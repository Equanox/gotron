UNAME := $(shell uname)

EXEC = ./example


#https://stackoverflow.com/questions/4058840/makefile-that-distincts-between-windows-and-unix-like-systems
# ifdef OS 
# 	EXEC = example.exe
# endif

run:
	@cd example && go build && $(EXEC)

install-ui:
	@cd example/ui && npm run install

build-ui:
	@cd example/ui && npm run build

builder:
	@make install-builder
	@gotron-builder -g=example
	
install-builder:
	@cd cmd/gotron-builder && go install

clean:
	@-rm -r .gotron
	@-rm -r .gotron-builder