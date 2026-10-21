package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 uint8 = 127
    var v2 int64 = 1
    v3 := []uint32{0, 352, 1}
    var v4 uint = 42
    var v5 uint64 = 1
    v5 = uint64(v1)
    if false {
        if (v2 < v2) {
            var v6 uint8 = v1
            _ = v6
        }
        var v7 int32 = 42
        _ = v7
    }
    v8 := GenS1{F0: 960, F1: (v1 & 31)}
    fmt.Println(v2, v2)
    fmt.Println(v5)
    if false {
        v9 := GenS1{F0: 65535, F1: (v1 % ((1) | 1))}
        v1 = uint8(v5)
        _ = v9
    }
    v10 := GenS1{F0: 255, F1: (v1 << (((v4 % ((v4) | 1))) & 63))}
    if len(v3) > 0 {
        v3[0] = 745
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
