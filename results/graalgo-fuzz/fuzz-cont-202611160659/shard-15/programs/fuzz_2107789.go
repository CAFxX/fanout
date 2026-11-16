package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint32; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 380
    v3 := []uint32{2147483648}
    var v4 uint8 = 255
    v5 := GenS1{F0: "hello"}
    v6 := GenS1{F0: ("\n\t" + (("\n\t" + "\n\t") + ""))}
    if true {
        v4 = v4
    }
    v4 += (v4 & (^(v4)))
    if true {
        v7 := GenS2{F0: -1, F1: 1, F2: 9223372036854775807}
        _ = v7
    } else {
        v2 = (^(v2))
        v4 = (v4 - (v4 / (((v4 | v4)) | 1)))
    }
    var v8 uint64 = 4294967295
    fmt.Println(v8, v4, v4)
    switch "hello" {
    case "hello":
        v8 = v8
    case "":
        var v9 int64 = 0
        _ = v9
    default:
        var v10 uint64 = v8
        _ = v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
