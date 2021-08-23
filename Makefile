setup:
	git submodule init && git submodule update

build: setup
	go run . $(cfg)

clean:
	rm -rf build out
