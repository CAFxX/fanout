package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint8; F1 int }
func genFn0() int16 {
    var v6 float64 = -2.25
    var v7 uint64 = 18446744073709551615
    _ = v6
    _ = v7
    return 1000
}
func genFn1(p0 string, p1 int16, p2 uint16) int32 {
    fmt.Println(p0, p1)
    var v8 int64 = -729
    if false {
        v9 := []int{127, 127}
        var v10 uint = 1
        var v11 uint8 = 128
        _ = v10
        _ = v11
        _ = v9
    } else {
        if false {
            if (-1 != -128) {
                p2 += p2
                v12 := GenS1{F0: 18446744073709551615, F1: true, F2: 591}
                v13 := GenS2{F0: 42}
                _ = v12
                _ = v13
            } else {
                v14 := GenS2{F0: v8}
                if (false || ("" == "\n\t")) {
                    v15 := GenS2{F0: v8}
                    p2 = 842
                    _ = v15
                } else {
                    fmt.Println(v8)
                }
                var v16 int64 = 9223372036854775807
                _ = v14
                _ = v16
            }
            v17 := []int{-379, 32768}
            _ = v17
        }
    }
    var v18 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v8
    return int32(v8)
}

func main() {
    var v1 uint16 = 215
    v2 := make(map[int]uint32)
    var v3 bool = false
    var v4 uint8 = 194
    var v5 int8 = 80
    v4 += (v4 & 164)
    v19 := GenS3{F0: 0, F1: 740}
    v20 := []int{-479}
    v4 -= 73
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 float64 = 1e10
        v23 := GenS1{F0: 1, F1: v3, F2: 42}
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS4{F0: v4, F1: 0}
    var v25 float32 = float32(0.1)
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 uint32 = 65535
        _ = fi26
        _ = v27
    }
    fmt.Println(v4, v1)
    v3 = v3
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
