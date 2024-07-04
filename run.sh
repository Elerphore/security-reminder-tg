#!/bin/sh
export PATH=$PATH:/usr/local/go/bin
nohup go run ./cmd/flowerJournal/ &
