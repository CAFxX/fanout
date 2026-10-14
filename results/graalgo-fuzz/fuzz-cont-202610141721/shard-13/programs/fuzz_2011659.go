package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint }
func genFn0(p0 int) uint8 {
    for fi6 := 0; fi6 < 11; fi6++ {
        switch (p0 ^ (-(fi6))) {
        case 128:
            v7 := GenS1{F0: 0, F1: 1}
            _ = v7
        case -9223372036854775808:
            var v8 uint64 = 0
            _ = v8
        default:
            v9 := []int{32767, 582, fi6}
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := []int{}
            _ = fi10
            _ = v11
        }
        _ = fi6
    }
    var v12 uint32 = 2
    _ = p0
    _ = v12
    return 125
}
func genFn1(p0 uint8, p1 uint64, p2 uint64) uint {
    var v13 uint8 = (p0 / ((p0) | 1))
    var v14 uint = 65535
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 string = "foo bar"
        for fi17 := 0; fi17 < 10; fi17++ {
            p1 = 18446744073709551615
            switch "" {
            case "":
                v18 := GenS1{F0: (v14 << ((584) & 63)), F1: uint32(p0)}
                _ = v18
            case "foo bar":
                var v19 uint32 = 1
                _ = v19
            case "hello":
                v14 -= 2
            default:
                var v20 uint8 = (p0 << ((0) & 63))
                _ = v20
            }
            _ = fi17
        }
        v21 := GenS1{F0: v14, F1: 8}
        _ = fi15
        _ = v16
        _ = v21
    }
    fmt.Println(p1)
    if true {
        v22 := GenS1{F0: v14, F1: 448}
        v23 := GenS1{F0: 514, F1: 724}
        if false {
            v24 := GenS2{F0: -2147483648, F1: 1000}
            p0 = p0
            _ = v24
        }
        _ = v22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return (^(v14))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = -2.25
    v3 := []string{}
    var v4 uint32 = 84
    v5 := []uint32{}
    v4 -= 65535
    var v25 float64 = (0.0 - 1e10)
    var v26 int32 = -2147483648
    v2 = 3.14159
    v27 := GenS2{F0: -9223372036854775808, F1: 1000}
    v28 := GenS3{F0: 230, F1: 0, F2: 1}
    var v29 int32 = (-(v26))
    v25 = (v25 - v25)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1[v4] = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
