#!/usr/bin/env bash
LD_LIBRARY_PATH=$(pwd)/lib/linux64:$(pwd)/lib/linux64/HCNetSDKCom go run main.go
# LD_LIBRARY_PATH=$(pwd)/lib/linux64:$(pwd)/lib/linux64/HCNetSDKCom ./hkweb
