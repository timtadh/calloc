# C Allocator (calloc)

by Tim Henderson

Licensed under the GPL version 3

## What is this?

This package allows you to do manual memory management via the standard
C allocator for your system. It is a thin wrapper on top of `malloc`,
`calloc` and `free` from `<stdlib.h>`. See `man malloc` for details on
these functions for your system. This library uses cgo.

### Why would you want this?

When a program is causing memory pressure or the system is running out
of memory it can be helpful to manually control memory allocations and
deallocations. Go can help you control allocations but it is not
possible to explicitly deallocate unneeded data. Using this library may
not solve the memory issues the program is experiencing but it is a tool
which may help.

### Danger!

You cannot mix data which was allocated with the regular Go allocator
and data which is allocated with this allocator. If you for instance
allocate a slice of pointers using `calloc.Make` and then insert
pointers allocated with the normal Go allocator Go will attempt to
collect those pointers! This will cause weird memory glitches when the
program tries to access that data later.

So be careful! Use this library as a last resort!

## Usage

See [go doc](https://godoc.org/github.com/timtadh/calloc) for API docs.

    package main

    import (
        "fmt"
        "reflect"
    )

    import(
        "github.com/timtadh/calloc"
    )


    type myStruct struct {
        a int
        b int
        c float32
        d []byte
    }

    func main() {
        var s *myStruct
        s = calloc.New(reflect.TypeOf(s)).(*myStruct)
        defer calloc.Free(s)
        s.a = 1
        s.b = 2
        s.c = 3.2
        str := []byte("hello world")
        s.d = calloc.Make(reflect.TypeOf(s.d), len(str), len(str)).([]byte)
        defer calloc.Free(s.d)
        copy(s.d, str)
        fmt.Println(s)
    }

