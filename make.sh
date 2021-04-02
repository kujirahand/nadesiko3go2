#!/bin/bash

SCR_DIR=$(cd $(dirname $0); pwd)

cd $SCR_DIR
python3 gen_vmcodename.py
cd $SCR_DIR/cmd/cnako
go build .
