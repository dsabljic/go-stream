#!/bin/sh

grep -vFf "$1" coverage.out > coverage.tmp && mv coverage.tmp coverage.out