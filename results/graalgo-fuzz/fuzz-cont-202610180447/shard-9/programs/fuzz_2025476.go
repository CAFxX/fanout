package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }

func main() {
    var v1 bool = true
    var v2 int16 = 1
    v3 := []uint32{}
    for fi4 := 0; fi4 < 19; fi4++ {
        fmt.Println(v1, fi4, fi4)
        var v5 uint64 = 518
        _ = fi4
        _ = v5
    }
    v2 -= (v2 + v2)
    var v6 int = 4294967295
    var v7 int8 = -91
    v7 = (v7 >> ((849) & 63))
    if (("foo bar" + (((("a" + ("foo bar" + "foo bar")) + "foo bar") + ("a" + "foo bar")) + ("\n\t" + "a"))) >= ("a" + ("a" + "hello"))) {
        v8 := GenS1{F0: 0, F1: v2}
        _ = v8
    }
    v2 = 32767
    v7 = 42
    var v9 bool = (0 <= -1)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
