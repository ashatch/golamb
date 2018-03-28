# golamb

Tiny utility to wrap the lambda build flags for OSX.

## Usage

    $golamb --help    
    golamb v0.1

    Usage of golamb:
      -handler string
          Handler name (default "main")

## What it does

 It essentially executes:

    GOOS=linux GOARCH=amd64 go build -o main main.go
    zip main.zip main

