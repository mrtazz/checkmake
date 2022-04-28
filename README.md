# checkmake

[![Build Status](https://travis-ci.org/mrtazz/checkmake.svg?branch=master)](https://travis-ci.org/mrtazz/checkmake)
[![Coverage Status](https://coveralls.io/repos/github/mrtazz/checkmake/badge.svg?branch=master)](https://coveralls.io/github/mrtazz/checkmake?branch=master)
[![Code Climate](https://codeclimate.com/github/mrtazz/checkmake/badges/gpa.svg)](https://codeclimate.com/github/mrtazz/checkmake)
[![Packagecloud](https://img.shields.io/badge/packagecloud-available-brightgreen.svg)](https://packagecloud.io/mrtazz/checkmake)
[![MIT license](https://img.shields.io/badge/license-MIT-blue.svg)](http://opensource.org/licenses/MIT)

## Overview
checkmake is an experimental tool for linting and checking Makefiles. It may
not do what you want it to.

## Usage

```
% checkmake Makefile

% checkmake Makefile foo.mk bar.mk baz.mk

% checkmake --help
checkmake.

Usage:
checkmake [--debug|--config=<configPath>] <makefile>...
checkmake -h | --help
checkmake --version

Options:
-h --help               Show this screen.
--version               Show version.
--debug                 Enable debug mode
--config=<configPath>   Configuration file to read
--list-rules            List registered rules

% checkmake fixtures/missing_phony.make

      RULE                 DESCRIPTION             LINE NUMBER

  minphony        Missing required phony target    0
                  "all"
  minphony        Missing required phony target    0
                  "test"
  phonydeclared   Target "all" should be           18
                  declared PHONY.

```

## Docker usage
Build the image, or pull it:
```sh
docker build --build-arg BUILDER_NAME='Your Name' --build-arg BUILDER_EMAIL=your.name@example.com . -t checker
```

Then run it with your Makefile attached, below is an example of it assuming the Makefile is in your current working directory:
```sh
docker run -v "$PWD"/Makefile:/Makefile checker
```

## Installation

### Requirements
The [pandoc](https://pandoc.org/) document converter utility is required to run checkmate. You can find out if you have it via `which pandoc`. [Install pandoc](https://pandoc.org/installing.html) if the command was not found.

### Packages
There are packages for linux up [on packagecloud.io](https://packagecloud.io/mrtazz/checkmake) or build it yourself with the steps below.

### Build
To build checkmake you will need to have [golang](https://golang.org/) installed. Once you have Go installed, you can simply clone the repo and build the binary and man page yourself with the following commands.

```sh
git clone https://github.com/mrtazz/checkmake
cd checkmake
make
```

## Inspiration
This is totally inspired by an idea by [Dan
Buch](https://twitter.com/meatballhat/status/768112351924985856).
