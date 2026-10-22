package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint16; F2 uint }
type GenS4 struct { F0 uint32; F1 int8 }

func main() {
    var v1 bool = true
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    var v4 uint64 = 4294967295
    switch -142 {
    case -2147483648:
        var v5 uint16 = 32767
        _ = v5
    case 0:
        var v6 int8 = 42
        _ = v6
    default:
        v4 -= v4
    }
    v1 = v1
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 bool = (-635 < int16(fi7))
            var v10 int8 = 124
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v4 += v4
    v4 -= v4
    v11 := GenS1{F0: 1}
    if false {
        v12 := GenS2{F0: 42}
        v1 = v1
        v4 = (v4 + v4)
        _ = v12
    } else {
        var v13 uint = 18446744073709551615
        _ = v13
    }
    var v14 int32 = 165
    var v15 uint64 = v4
    var v16 int = 42
    var v17 uint8 = 255
    var v18 float32 = (float32(0.0) + (((float32((3.14159 + 3.14159)) * -2.25) - (3.14159 - (1.5 / -2.25))) / 1.5))
    if false {
        v17 = v17
    } else {
        v19 := GenS3{F0: "", F1: 65535, F2: 588}
        v20 := GenS4{F0: 2147483648, F1: 1}
        v17 += 50
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
