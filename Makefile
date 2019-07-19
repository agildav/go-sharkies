# Load env variables, installation required
# go get github.com/joho/godotenv/cmd/godotenv
GODOTENV= godotenv
# Go parameters
GO_BUILD= go build
GO_CLEAN= go clean
GO_TEST= go test
BINARY_NAME= sharkies

all: test build
build: clean
	$(GO_BUILD) -o $(BINARY_NAME) -v
clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)
run: build
	$(GODOTENV) ./$(BINARY_NAME)
test:
	$(GODOTENV) $(GO_TEST) ./... -count=1
test-v:
	$(GODOTENV) $(GO_TEST) ./... -count=1 -v