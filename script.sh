#!/bin/bash

# load the execution date
date >> date.txt

## Replace text
sed 's/This is a Golang example/This is an example with a pre-build step/g'  ./main.go > ./main_aux.go

mv main_aux.go main.go

