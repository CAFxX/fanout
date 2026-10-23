package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16; F1 int64; F2 uint32 }
func genFn0() uint16 {
    var v6 int64 = 0
    v6 = 42
    v6 += v6
    if (0.1 >= (0.0 - (0.0 / 1.0))) {
        fmt.Println(v6)
        fmt.Println(v6)
        var v7 uint8 = 0
        _ = v7
    } else {
        var v8 int8 = 100
        v9 := []int{-1, -512, 9223372036854775807, 228}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        if false {
            var v11 bool = false
            for fi12 := 0; fi12 < 17; fi12++ {
                switch fi12 {
                case 2147483647:
                    var v13 float64 = -2.25
                    _ = v13
                }
                switch fi12 {
                case -429:
                    var v14 uint16 = 2
                    _ = v14
                }
                v11 = (fi12 <= fi12)
                _ = fi12
            }
            _ = v11
        }
        _ = fi10
    }
    _ = v6
    return 65535
}
func genFn1(p0 uint8, p1 bool, p2 uint16) int {
    v15 := GenS1{F0: 0}
    var v16 float64 = (1.5 + -2.25)
    for fi17 := 0; fi17 < 10; fi17++ {
        if p1 {
            var v18 int16 = 32767
            v19 := GenS2{F0: -760, F1: 556}
            v20 := GenS2{F0: (v18 | v18), F1: 1}
            _ = v18
            _ = v19
            _ = v20
        }
        var v21 int32 = 2147483647
        var v22 float32 = 1.5
        _ = fi17
        _ = v21
        _ = v22
    }
    p2 += (p2 >> ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-2147483648, 127}
    var v3 float32 = 3.14159
    var v4 float32 = (v3 / float32(0.1))
    v5 := make(map[uint32]string)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "" {
    case "\n\t":
        fmt.Println(v3, v4)
    case "hello":
        var v23 uint16 = 806
        _ = v23
    case "foo bar":
        v3 -= -2.25
    default:
        v24 := GenS2{F0: 1, F1: 42}
        _ = v24
    }
    v25 := GenS3{F0: 0}
    v26 := GenS4{F0: 42, F1: 9223372036854775807}
    if true {
        v4 += 1e10
        v3 = (v4 + float32(1.5))
    }
    v3 = (v4 * (v3 - v3))
    v27 := GenS5{F0: 65535}
    if false {
        switch ("" + "a") {
        case "a":
            v28 := GenS6{F0: 1000, F1: 248, F2: 148}
            _ = v28
        case "hello":
            for fi29 := 0; fi29 < 20; fi29++ {
                _ = len(v1)
                _ = fi29
            }
        }
        fmt.Println(v3, v3)
    } else {
        var v30 uint = 1
        _ = v30
    }
    v3 -= v4
    v3 = 0.1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
