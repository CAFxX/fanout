package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int64; F1 bool; F2 int32 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint16 }

func main() {
    var v1 uint16 = 2
    var v2 float64 = 0.1
    v3 := []uint32{548, 864, 155, 2}
    var v4 uint16 = (v1 * (v1 % ((286) | 1)))
    v5 := []uint32{}
    var v6 uint16 = v1
    v1 = v1
    var v7 bool = false
    fmt.Println(v6)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v8 := GenS1{F0: 147, F1: 950, F2: 255}
    switch 9223372036854775807 {
    case 1:
        if len(v5) > 0 {
            v5[0] = 1
        }
    default:
        if (float64(-2.25) != ((v2 + v2) * ((((3.14159 * (v2 + float64(1.5))) - v2) + 1.5) * v2))) {
            _ = len(v5)
        }
    }
    v9 := GenS1{F0: 255, F1: 42, F2: 128}
    var v10 float32 = -2.25
    v11 := GenS2{F0: 42, F1: (42 >= 2147483647), F2: 532}
    v12 := GenS3{F0: 0, F1: 42, F2: (v1 ^ (v4 % (((v6 & v6)) | 1)))}
    fmt.Println(v6)
    fmt.Println(v4, v6, v4)
    v4 += 255
    v13 := GenS1{F0: 255, F1: 2147483647, F2: 255}
    var v14 uint32 = 4294967295
    if (!(!true)) {
        var v15 float64 = (float64(1e10) / v2)
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
