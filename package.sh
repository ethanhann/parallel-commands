#!/usr/bin/env bash

go build -o ./build/pc
cd ./build || exit
zip -9 "pc.zip" pc