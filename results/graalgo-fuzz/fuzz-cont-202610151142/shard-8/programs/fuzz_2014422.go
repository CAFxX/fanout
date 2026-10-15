package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{("a" + "hello")}
    var v3 int8 = 100
    v4 := make(map[int]string)
    var v5 uint = 502
    v6 := []uint32{0}
    var v7 bool = false
    v3 = (-(v3))
    v3 -= v3
    v7 = v7
    v1 -= v1
    v1 -= -386
    for fi8 := 0; fi8 < 16; fi8++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v9 uint64 = 909
        var v10 uint = v5
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: v5, F1: v7, F2: (v3 & v3)}
    if (!v7) {
        _ = len(v4)
    } else {
        fmt.Println(v7, v7)
        v12 := v4[-38]
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v5 = 42
        _ = fi13
    }
    v14 := GenS1{F0: (v5 * v5), F1: false, F2: v3}
    if v7 {
        v15 := GenS1{F0: 4294967295, F1: (!(!true)), F2: -1}
        v16 := GenS2{F0: 42}
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v4[fi17] = ""
        v18 := GenS3{F0: (fi17 | (fi17 << ((255) & 63)))}
        v1 = v1
        _ = fi17
        _ = v18
    }
    _ = len(v4)
    var v19 uint32 = 65535
    fmt.Println(v5, v1)
    var v20 float64 = (3.14159 / (0.1 - ((((1e10 * 3.14159) - 0.0) * 0.1) - 3.14159)))
    for fi21 := 0; fi21 < 19; fi21++ {
        v1 = v1
        v22 := GenS3{F0: 127}
        var v23 uint8 = 121
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
