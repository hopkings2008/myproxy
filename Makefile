.PHONY: binary clean

ROOT=$(shell pwd)/../..

CURDIR=$(shell pwd)

TARGET=goproxy

binary:
	go build -gcflags "-N -l" -o $(TARGET)

clean:
	-@rm -f $(TARGET)
