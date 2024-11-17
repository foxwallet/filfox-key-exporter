TARGET=filfox-key-exporter

GO_BUILD=go build

OUT_DIR=bin

all: mac_arm windows_x86_64

mac_arm:
	GOOS=darwin GOARCH=arm64 $(GO_BUILD) -o $(OUT_DIR)/$(TARGET)_darwin_arm64

windows_x86_64:
	GOOS=windows GOARCH=amd64 $(GO_BUILD) -o $(OUT_DIR)/$(TARGET)_win64.exe

sha256:
	sha256sum $(OUT_DIR)/*

clean:
	rm -rf $(OUT_DIR)

$(shell mkdir -p $(OUT_DIR))
