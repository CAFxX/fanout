package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 bool }

func main() {
    v1 := []uint32{42}
    var v2 uint = 65535
    var v3 int32 = 0
    for fi4 := 0; fi4 < 13; fi4++ {
        if true {
            var v5 bool = false
            v6 := GenS1{F0: 1000}
            v2 += v2
            _ = v5
            _ = v6
        } else {
            fmt.Println(v2, fi4, fi4)
            _ = len(v1)
        }
        var v7 bool = false
        v8 := GenS1{F0: 1}
        _ = fi4
        _ = v7
        _ = v8
    }
    v2 -= v2
    switch 128 {
    case 915:
        var v9 int8 = -1
        _ = v9
    case 239:
        v2 -= 744
    default:
        var v10 uint16 = 1
        _ = v10
    }
    if (-2.25 <= 0.1) {
        fmt.Println(v3, v2)
    }
    v11 := GenS1{F0: -1}
    v2 += (v2 % ((42) | 1))
    v3 = (v3 % ((v3) | 1))
    if false {
        _ = len(v1)
        v3 = v3
    }
    v2 -= 65535
    v12 := GenS1{F0: 32767}
    v2 += 18446744073709551615
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS2{F0: -32768}
        v15 := GenS3{F0: false, F1: true}
        fmt.Println(fi13, fi13, v3)
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int16 = -32768
    var v17 uint64 = 2
    fmt.Println(v16, v17)
    fmt.Println(v2)
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
