mumax3 
======
[![Build Status](https://travis-ci.org/jomafu/Mumax3_DH.svg?branch=master)](https://travis-ci.org/jomafu/Mumax3_DH)

GPU accelerated micromagnetic simulator.


Downloads and documentation
---------------------------

http://mumax.github.io


Paper
-----

The Design and Verification of mumax3:

http://scitation.aip.org/content/aip/journal/adva/4/10/10.1063/1.4899186


Tools
-----

https://godoc.org/github.com/jomafu/Mumax3_DH/cmd


Building from source (for linux)
--------------------

Consider downloading a pre-compiled binary. If you want to compile nevertheless:

  * install the nvidia proprietary driver, if not yet present.
   - if unsure, it's probably already there
   - version 440.44 recommended
  * install Go 
    - https://golang.org/dl/
    - set $GOPATH
  * install CUDA 
    - https://developer.nvidia.com/cuda-downloads (pick default installation path)
    - or `sudo apt-get install nvidia-cuda-toolkit`
  * install a C compiler
    - on Ubuntu: `sudo apt-get install gcc`
  * if you have git installed: 
    - `go get github.com/jomafu/Mumax3_DH/cmd/mumax3`
  * if you don't have git:
    - seriously, no git?
    - get the source from https://github.com/jomafu/Mumax3_DH/releases
    - unzip the source into $GOPATH/src/github.com/jomafu/Mumax3_DH
    - `cd $GOPATH/src/github.com/jomafu/Mumax3_DH/cmd/mumax3`
    - `go install`
  * optional: install gnuplot if you want pretty graphs
    - on ubuntu: `sudo apt-get install gnuplot`
  * use the Makefile if there is a need to recompile the cuda kernels
    - `make realclean && make`

Your binary is now at `$GOPATH/bin/mumax3`

Contributing
------------

Contributions are gratefully accepted. To contribute code, fork our repo on github and send a pull request.
