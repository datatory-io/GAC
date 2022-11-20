#!/bin/bash

GAC=bin/gac

# remove after dev phase
rm $GAC
########################


if [ ! -f "$GAC" ]; then
    (cd ./generator && go build -o ../bin/gac .)
fi

$GAC -d=0 -o $2 -include-boilerplate=0 -include-runtime=0 $3
