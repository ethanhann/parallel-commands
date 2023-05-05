#!/usr/bin/env bash

curl -s https://api.github.com/repos/ethanhann/parallel-commands/releases/latest |
  grep "pc.zip" |
  cut -d : -f 2,3 |
  tr -d \" |
  wget -qi -

unzip pc.zip
mv ./pc /usr/local/bin
rm pc.zip
