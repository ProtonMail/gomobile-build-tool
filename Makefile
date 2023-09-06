LINTERS_VERSION=v1.33.0
LINTERS_REPO=github.com/golangci/golangci-lint/cmd/golangci-lint

setup:
	git submodule init && git submodule update

build: setup
	go run . $(cfg)

clean:
	rm -rf build out

install-linters:
	go install golang.org/x/lint/golint@latest
	go install $(LINTERS_REPO)@$(LINTERS_VERSION)

lint:
	golangci-lint run ./...

zip-apple:
	cd out && zip -r Crypto.xcframework.zip Crypto.xcframework/
	
zip-android:
	cd out && zip -r android.zip android

