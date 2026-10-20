package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }

func main() {
    v1 := []string{"foo bar", ("\n\t" + ""), ""}
    var v2 uint8 = 1
    var v3 uint8 = (v2 << ((18446744073709551615) & 63))
    var v4 int16 = 566
    v5 := GenS1{F0: "\n\t", F1: v2}
    v1 = append(v1, "foo bar")
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = (v3 | (v2 >> ((4294967295) & 63)))
    var v6 string = "a"
    v3 = (v3 % ((v3) | 1))
    _ = len(v1)
    var v7 int8 = 1
    _ = len(v1)
    var v8 uint64 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
