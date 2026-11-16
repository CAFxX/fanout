package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    v3 := []uint32{1, 541}
    var v4 uint32 = 304
    v5 := []uint32{v4}
    var v6 int16 = 32767
    var v7 uint = 42
    var v8 int16 = -1
    v7 -= 478
    for fi9 := 0; fi9 < 2; fi9++ {
        delete(v2, (v4 ^ v4))
        delete(v2, (v4 - (^(v4))))
        _ = fi9
    }
    if ((false && (-93 <= -125)) || (!(201 < 0))) {
        if (32767 != 32767) {
            fmt.Println(v8)
            v10 := GenS1{F0: 1, F1: 53}
            _ = v10
        }
    } else {
        v2[2] = ""
        v4 = 1
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(v6, v8)
        _ = fi11
    }
    v4 = v4
    if false {
        v12 := GenS1{F0: -430, F1: uint64(v4)}
        _ = v12
    }
    var v13 string = "a"
    _ = len(v2)
    v6 += v6
    var v14 uint16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
