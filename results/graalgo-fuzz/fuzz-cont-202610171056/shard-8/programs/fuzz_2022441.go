package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }

func main() {
    v1 := []string{"foo bar"}
    var v2 int16 = 966
    var v3 bool = true
    fmt.Println(v2)
    if (!true) {
        var v4 uint16 = 1
        _ = v4
    }
    fmt.Println(v2, v3)
    fmt.Println(v2)
    v2 = v2
    fmt.Println(v3, v3)
    v5 := GenS1{F0: -992, F1: 1}
    if (149 < 4294967295) {
        if v3 {
            var v6 uint16 = 2
            var v7 int = 0
            _ = v6
            _ = v7
        }
    }
    switch "a" {
    case "\n\t":
        _ = len(v1)
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "a":
        var v8 int = 0
        _ = v8
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
}
