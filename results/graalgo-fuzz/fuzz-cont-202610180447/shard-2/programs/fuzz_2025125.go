package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("" + "hello"), "", ((("\n\t" + "") + ("\n\t" + "")) + "foo bar"), ("" + "hello")}
    v3 := make(map[string]string)
    var v4 uint = 247
    switch "a" {
    case "foo bar":
        v4 = v4
    case "hello":
        v5 := GenS1{F0: 42, F1: -1}
        _ = v5
    default:
        v4 = v4
    }
    switch "\n\t" {
    case "\n\t":
        v2 = append(v2, ("a" + "hello"))
    default:
        fmt.Println(v4)
    }
    var v6 int64 = 1
    var v7 uint64 = 2
    v6 += (^(v6))
    v7 = 735
    fmt.Println(v4, v7, v4)
    var v8 uint16 = 1
    var v9 int64 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
