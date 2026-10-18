package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 42
    v3 := []uint32{301, 4294967295}
    v4 := make(map[string]int)
    v5 := GenS1{F0: 4294967295, F1: false, F2: 781}
    var v6 float32 = 0.0
    v6 = v6
    var v7 int16 = 42
    var v8 int = 32768
    if true {
        v9 := GenS1{F0: 65535, F1: false, F2: 42}
        _ = v9
    } else {
        if true {
            v2 = -1
        }
        var v10 uint8 = 127
        v10 -= (v10 & (v10 & 20))
        _ = v10
    }
    v11 := GenS2{F0: 127, F1: 2}
    if (true && false) {
        var v12 int8 = 1
        for fi13 := 0; fi13 < 11; fi13++ {
            v7 = -32768
            fmt.Println(v12)
            _ = fi13
        }
        var v14 int16 = (v7 - v7)
        _ = v12
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
