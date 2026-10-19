package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int64; F1 int16; F2 bool }
func genFn0(p0 int32, p1 int) int32 {
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{(p1 << ((1) & 63)), -1, p1, 32768}
        var v8 uint = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 float32 = 1.5
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: "\n\t", F1: 597}
    var v12 string = "a"
    if false {
        v13 := GenS2{F0: 1, F1: 18446744073709551615}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return p0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 127
    v3 := make(map[uint32]string)
    v4 := []int{-1, -303}
    var v5 int32 = -636
    v2 = (^(v2))
    for fi14 := 0; fi14 < 16; fi14++ {
        v4 = append(v4, -1)
        fmt.Println(fi14)
        _ = fi14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint32 = 0
        fmt.Println(v5, v16, v5)
        var v17 uint8 = (v2 / ((0) | 1))
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint8 = 1
    fmt.Println(v18, v18, v5)
    var v19 int32 = 1
    v20 := GenS1{F0: "hello", F1: 0}
    var v21 bool = ((v18 << ((255) & 63)) > v2)
    v18 = v18
    for fi22 := 0; fi22 < 4; fi22++ {
        v4 = append(v4, (fi22 / (((-(fi22))) | 1)))
        _ = fi22
    }
    if v21 {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 int = (^(fi23))
            _ = len(v4)
            var v25 bool = v21
            _ = fi23
            _ = v24
            _ = v25
        }
        v4 = append(v4, 127)
    } else {
        for fi26 := 0; fi26 < 1; fi26++ {
            if (496 == int16(v18)) {
                v27 := GenS3{F0: 9223372036854775807, F1: 42, F2: v21}
                var v28 uint16 = 32767
                _ = v27
                _ = v28
            } else {
                var v29 uint8 = (v18 ^ 42)
                _ = v29
            }
            _ = fi26
        }
        v19 -= genFn0(1, 9223372036854775807)
    }
    var v30 int32 = (v19 * v19)
    v21 = (-2.25 <= (3.14159 / ((1.5 + 1e10) + (3.14159 + 1.5))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
