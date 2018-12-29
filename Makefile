EXEC = ./example
VERSION = 0.2.2

run:
	@cd example && go build && $(EXEC)

install-ui:
	@cd example/ui && npm run install

build-ui:
	@cd example/ui && npm run build

builder:
	@make install-builder
	@gotron-builder -g=example --win
	
install-builder:
	@cd cmd/gotron-builder && go install

clean:
	@-rm -r .gotron
	@-rm -r .gotron-builder

################ test ################
test-ci:
	@-rm -r .gotron .gotron-builder example/.gotron example/.gotron-builder

	@make release
	@make release-clean

	@make install-builder
	@gotron-builder -g example -l --ia32
	@cd example && gotron-builder
	@gotron-builder -g example -a example/ui/build -w
	@gotron-builder -g example --out example -w --ia32

test-clean-build:
	@docker build -f test/Dockerfile .

################ release #############
# Create releasable executable of gotron-builder
release:
	@-mkdir release 
	@-rm release/*
	
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -o ./release/gotron-builder-amd64-linux \
	-ldflags="-X main.gotronBuilderVersion=$(VERSION)" \
	-a cmd/gotron-builder/main.go
	
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 \
	go build -o ./release/gotron-builder-amd64-darwin \
	-ldflags="-X main.gotronBuilderVersion=$(VERSION)" \
	-a cmd/gotron-builder/main.go
	
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
	go build -o ./release/gotron-builder-amd64-win \
	-ldflags="-X main.gotronBuilderVersion=$(VERSION)" \
	-a cmd/gotron-builder/main.go

release-clean:
	@-rm -r release