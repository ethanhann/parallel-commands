#!/usr/bin/env bash

go build -o ./build/pc
cd ./build || exit
zip "pc.zip" pc