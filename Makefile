BUILD_DIR=$(shell pwd)/bin
VANILLA_DIR=$(shell pwd)/cli/vanilla
clean:
	rm -rf $(BUILD_DIR)/*

build:
	cp -R $(VANILLA_DIR)/assets $(BUILD_DIR)/
	cp -R $(VANILLA_DIR)/config $(BUILD_DIR)/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	go build -v -o $(BUILD_DIR)/vanilla $(VANILLA_DIR)
