#! /bin/bash

# builds with -race and runs tests with browser open.

set -e

go install -race github.com/jomafu/Mumax3_DH/cmd/mumax3

google-chrome http://localhost:35367 &

for f in *.mx3; do
	mumax3 $f 
done

go install github.com/jomafu/Mumax3_DH/cmd/mumax3 # re-build without race detector

