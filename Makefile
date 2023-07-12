.PHONY: setup build clean run

setup:
	python3 -m pip install pybindgen && go install golang.org/x/tools/cmd/goimports@latest && go install github.com/go-python/gopy@latest

build: clean
	gopy build -output=out -vm=python3 ./hi

clean:
	rm -rf ./out

run:
	python main.py
