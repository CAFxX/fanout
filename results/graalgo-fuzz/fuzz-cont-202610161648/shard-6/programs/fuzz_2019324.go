package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 string; F1 int16; F2 int8 }
type GenS3 struct { F0 bool; F1 bool }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 bool = true
    var v3 int32 = 42
    v4 := []uint32{2147483647, 65535, 940}
    v5 := make(map[int]bool)
    v6 := make(map[int]uint32)
    fmt.Println(v2, v2)
    if v2 {
        v2 = (42 > 1)
        if v2 {
            v4 = append(v4, 975)
            v2 = v2
        }
        var v7 int64 = v1
        _ = v7
    } else {
        var v8 float32 = 0.0
        v6[128] = 65535
        _ = v8
    }
    fmt.Println(v1, v2, v3)
    fmt.Println(v1, v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v9 float64 = float64(0.0)
    v5[-9223372036854775808] = false
    v1 -= v1
    var v10 uint8 = 127
    fmt.Println(v10)
    if (!(9223372036854775807 >= 0)) {
        if (1 <= 2) {
            if v2 {
                v3 += (v3 & v3)
            }
            if len(v4) > 0 {
                v4[0] = 42
            }
        } else {
            v11 := GenS1{F0: v10, F1: 224}
            var v12 int8 = 100
            _ = v11
            _ = v12
        }
        var v13 uint8 = v10
        _ = v13
    } else {
        v14 := GenS1{F0: 127, F1: 2}
        switch -107 {
        case -9223372036854775808:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 1:
            var v15 uint32 = 42
            _ = v15
        case -1:
            v16 := GenS2{F0: "\n\t", F1: -1, F2: -1}
            _ = v16
        }
        _ = v14
    }
    var v17 uint = 1
    switch 4294967295 {
    case 32768:
        _ = len(v6)
    case 233:
        var v18 uint8 = v10
        _ = v18
    case 127:
        var v19 float64 = (v9 - (v9 - v9))
        _ = v19
    default:
        for fi20 := 0; fi20 < 10; fi20++ {
            delete(v6, (^(fi20)))
            v21 := GenS3{F0: false, F1: false}
            v17 = v17
            _ = fi20
            _ = v21
        }
    }
    v22 := GenS1{F0: 42, F1: 0}
    for fi23 := 0; fi23 < 7; fi23++ {
        fmt.Println(v10, fi23)
        if v2 {
            v24 := GenS1{F0: v10, F1: 2}
            var v25 uint = (v17 & (v17 + (^(v17))))
            v17 += 18446744073709551615
            _ = v24
            _ = v25
        }
        _ = fi23
    }
    v26 := GenS2{F0: "\n\t", F1: 42, F2: -40}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
