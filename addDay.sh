#!/bin/bash

function usage {
    echo "Wrong number of arguments" 
}

if [[ $# != 1 ]]; then
    usage
    exit 1
fi

aoc_root=$HOME/go/src/github.com/iCiaran/AoC19

if [[ ! -d "${aoc_root}/day_$1" ]]; then
    echo "Creating directory"
    mkdir -p ${aoc_root}/day_$1/inputs
    cp ${aoc_root}/template.go ${aoc_root}/day_$1/main.go
    cp ${aoc_root}/template_test.go ${aoc_root}/day_$1/main_test.go
    touch ${aoc_root}/day_$1/README.md
fi
