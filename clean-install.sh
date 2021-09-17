#!/bin/sh

rm -f $GOPATH/bin/meow
go install .
echo "Successfully installed meow"
