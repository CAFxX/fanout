package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"", ""}
    v3 := []int{-9223372036854775808}
    switch "foo bar" {
    case "hello":
        var v4 int64 = -1
        _ = v4
    }
    var v5 int = 4294967295
    delete(v1, v5)
    fmt.Println(v5)
    switch "\n\t" {
    case "a":
        delete(v1, v5)
    }
    if (!true) {
        if (!false) {
            v5 = 980
            var v6 int32 = 2147483647
            _ = v6
        }
        _ = len(v1)
    }
    v7 := GenS1{F0: (v5 - (v5 & (v5 >> ((929) & 63)))), F1: 42}
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v7
}
