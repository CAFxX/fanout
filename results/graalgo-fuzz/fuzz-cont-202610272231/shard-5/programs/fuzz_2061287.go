package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 bool }

func main() {
    v1 := []uint32{42, 0, 2147483648, 935}
    v2 := make(map[int]bool)
    var v3 uint32 = 449
    if true {
        delete(v2, -9223372036854775808)
    } else {
        if (2 < 255) {
            var v4 uint8 = 0
            fmt.Println(v4, v3)
            _ = len(v1)
            _ = v4
        }
        v5 := GenS1{F0: 0, F1: -128, F2: false}
        switch ("" + "") {
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "hello":
            var v6 uint32 = (^(v3))
            _ = v6
        default:
            v3 = v3
        }
        _ = v5
    }
    fmt.Println(v3)
    var v7 uint8 = 0
    fmt.Println(v7, v7)
    v7 -= v7
    var v8 uint8 = 0
    v7 = 246
    if (true && (269 > 1)) {
        v1 = append(v1, 992)
    } else {
        var v9 uint = 2
        _ = v9
    }
    v7 -= (v8 ^ v7)
    switch "hello" {
    case "\n\t":
        v7 = (v7 + v7)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
