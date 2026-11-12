package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 bool; F1 int; F2 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 int8 }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 7; fi4++ {
        if true {
            var v5 int16 = 42
            v6 := GenS1{F0: 0}
            fmt.Println(v5)
            _ = v5
            _ = v6
        } else {
            v7 := GenS1{F0: -799}
            v8 := GenS2{F0: (fi4 | fi4), F1: 0}
            _ = v7
            _ = v8
        }
        var v9 string = "hello"
        var v10 uint32 = 0
        _ = fi4
        _ = v10
        _ = v9
    }
    v11 := GenS3{F0: (false && true), F1: 2147483647, F2: 50}
    var v12 uint32 = 2147483647
    _ = v11
    _ = v12
    return 18446744073709551615
}
func genFn1(p0 int32, p1 int32) uint32 {
    var v13 uint8 = 255
    v13 = (v13 >> ((658) & 63))
    _ = p0
    _ = p1
    _ = v13
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 2
    v3 := []uint32{}
    var v14 uint16 = 683
    v15 := GenS4{F0: 128}
    var v16 float64 = -2.25
    v17 := GenS1{F0: 362}
    fmt.Println(v2)
    v18 := GenS4{F0: (v2 * (v2 / (((v2 ^ v2)) | 1)))}
    v14 = (^(v14))
    v2 = v2
    fmt.Println(v16, v16)
    var v19 uint = 42
    v14 = 599
    for fi20 := 0; fi20 < 7; fi20++ {
        v21 := GenS5{F0: genFn1(-1, 0), F1: -47}
        _ = fi20
        _ = v21
    }
    var v22 uint8 = (v2 | v2)
    var v23 int16 = 569
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
