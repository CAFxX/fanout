package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int16; F2 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 int8; F2 uint }

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]bool)
    var v3 uint = 65535
    if false {
        v4 := GenS1{F0: 32767}
        v5 := GenS2{F0: 219, F1: 17, F2: -772}
        _ = v4
        _ = v5
    } else {
        v6 := GenS3{F0: ((1 < 602) || true)}
        if false {
            v3 = v3
            var v7 float64 = 0.1
            v8 := GenS4{F0: ("a" + "\n\t")}
            _ = v7
            _ = v8
        } else {
            var v9 uint64 = 425
            v10 := v2[4294967295]
            _ = v10
            _ = v9
        }
        fmt.Println(v1, v1)
        _ = v6
    }
    v1 = 0
    v1 = -1
    v11 := []int{1, 32768, -166, 2147483647}
    switch -2147483648 {
    case 223:
        var v12 uint = 1
        _ = v12
    case -2147483648:
        var v13 uint64 = 1
        _ = v13
    case 128:
        fmt.Println(v1)
    }
    v3 += (v3 / ((v3) | 1))
    if len(v11) > 0 {
        v11[0] = 0
    }
    v14 := GenS5{F0: 128, F1: 69, F2: (v3 / (((^(v3))) | 1))}
    var v15 float64 = 0.1
    var v16 int32 = v1
    var v17 float32 = (float32(float64((0.0 / 1e10))) + 0.0)
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            v1 = (-(v1))
            delete(v2, 4294967295)
            _ = fi19
        }
        var v20 bool = (v3 < 255)
        _ = fi18
        _ = v20
    }
    v17 = v17
    v2[2147483647] = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
