package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int8) int16 {
    for fi6 := 0; fi6 < 1; fi6++ {
        if true {
            v7 := []int{fi6, 9223372036854775807}
            _ = v7
        }
        switch (^(fi6)) {
        case -9223372036854775808:
            p0 = p0
        case -476:
            v8 := GenS1{F0: 845, F1: fi6, F2: 405}
            _ = v8
        default:
            if true {
                var v9 int8 = p0
                for fi10 := 0; fi10 < 2; fi10++ {
                    v11 := GenS1{F0: 2147483647, F1: (^(fi6)), F2: 255}
                    _ = fi10
                    _ = v11
                }
                fmt.Println(v9)
                _ = v9
            } else {
                v12 := []int{fi6, (fi6 + fi6)}
                v13 := GenS1{F0: 1, F1: (-(fi6)), F2: 2}
                _ = v12
                _ = v13
            }
        }
        p0 = p0
        _ = fi6
    }
    fmt.Println(p0)
    _ = p0
    return 42
}
func genFn1(p0 uint64) uint {
    var v14 uint16 = 255
    v14 += v14
    v15 := GenS2{F0: 1000}
    _ = p0
    _ = v14
    _ = v15
    return 879
}

func main() {
    var v1 uint32 = 2
    var v2 int32 = 2147483647
    var v3 float64 = 3.14159
    v4 := make(map[uint32]uint32)
    var v5 string = "hello"
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := GenS3{F0: 255}
        _ = fi16
        _ = v17
    }
    v18 := []int{-9223372036854775808, 32767, 1, -249}
    var v19 uint32 = (v1 ^ (v1 & v1))
    v20 := GenS4{F0: -9223372036854775808, F1: 0}
    v21 := GenS5{F0: 128}
    var v22 float64 = 0.0
    var v23 float32 = 3.14159
    v22 += v22
    if true {
        v2 += (v2 ^ -475)
        v24 := v4[(v1 ^ v1)]
        for fi25 := 0; fi25 < 8; fi25++ {
            fmt.Println(v1)
            var v26 uint16 = 32767
            v27 := GenS6{F0: -1}
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = v24
    } else {
        var v28 uint = 42
        var v29 bool = ((^(v19)) != v1)
        v19 = (^(v1))
        _ = v28
        _ = v29
    }
    v19 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
