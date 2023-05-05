#!/usr/bin/env bash

cd ./src || exit
go build -o ../build/pc
cd ../build || exit
zip -9 "pc.zip" pc
