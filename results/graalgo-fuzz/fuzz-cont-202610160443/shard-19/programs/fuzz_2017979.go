package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
func genFn0(p0 int32, p1 uint8) uint8 {
    var v5 int64 = 42
    var v6 int = -769
    v7 := []int{(v6 - -830), v6}
    if false {
        var v8 float32 = (float32(1.5) * ((3.14159 * (float32((0.1 / 3.14159)) - 1.5)) * (3.14159 * 0.1)))
        _ = v8
    } else {
        v9 := GenS1{F0: true, F1: 2}
        fmt.Println(p0, p1, v5)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return p1
}
func genFn1(p0 uint64, p1 uint32) int64 {
    var v10 int64 = 0
    var v11 bool = false
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return (v10 + v10)
}

func main() {
    var v1 uint64 = 0
    v2 := []uint32{963, 2, 2}
    var v3 uint16 = 2
    var v4 uint32 = 4294967295
    var v12 uint8 = 42
    switch 9223372036854775807 {
    case 127:
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 int = (fi13 ^ 128)
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 float64 = 1.5
                var v17 uint = 255
                v4 = v4
                _ = fi15
                _ = v16
                _ = v17
            }
            switch fi13 {
            case 128:
                v14 = 4294967295
            default:
                fmt.Println(v12, fi13)
            }
            _ = fi13
            _ = v14
        }
    }
    var v18 uint32 = (^(v4))
    v18 = v18
    fmt.Println(v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 = (v18 + v18)
    var v19 uint16 = 900
    v3 += v19
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v18 += (^(v18))
        _ = fi20
    }
    if true {
        v4 = uint32(v1)
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 float32 = -2.25
            _ = fi22
            _ = v23
        }
        var v24 uint = 2
        _ = fi21
        _ = v24
    }
    var v25 uint8 = (v12 << ((18446744073709551615) & 63))
    for fi26 := 0; fi26 < 3; fi26++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v19, fi26, v4)
        _ = fi26
    }
    if false {
        var v27 int32 = 42
        var v28 float32 = ((0.1 - -2.25) - float32((1e10 - (1.5 + 3.14159))))
        var v29 int32 = v27
        _ = v27
        _ = v28
        _ = v29
    } else {
        fmt.Println(v25, v3, v1)
        for fi30 := 0; fi30 < 15; fi30++ {
            v12 -= uint8(v4)
            _ = fi30
        }
    }
    fmt.Println(v3, v25)
    v31 := GenS1{F0: true, F1: (v1 - v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v4
}
