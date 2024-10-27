#!/bin/sh

BINARY='/usr/local/bin'

echo "Building ds"
go build ds.go

echo "Installing ds to $BINARY"
install -v ds $BINARY

echo "Removing the build"
rm ds
