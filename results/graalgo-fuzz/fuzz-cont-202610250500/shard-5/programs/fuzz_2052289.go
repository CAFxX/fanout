package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int16; F2 int }
type GenS4 struct { F0 int; F1 uint64; F2 uint }
func genFn0() bool {
    var v4 int = 32767
    v5 := []int{}
    v6 := GenS1{F0: 194}
    v4 = v4
    v4 = v4
    _ = v4
    _ = v5
    _ = v6
    return (1 != 0)
}
func genFn1(p0 uint, p1 int) uint {
    switch "a" {
    case "\n\t":
        v7 := GenS2{F0: 127}
        _ = v7
    }
    var v8 string = "foo bar"
    if false {
        v9 := GenS2{F0: (p1 / ((p1) | 1))}
        _ = v9
    }
    p0 += p0
    if true {
        v10 := GenS1{F0: 101}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return (p0 - 255)
}

func main() {
    var v1 float64 = -2.25
    var v2 int64 = 42
    var v3 bool = false
    v11 := GenS2{F0: 32768}
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: 218}
        _ = fi12
        _ = v13
    }
    v2 -= (v2 / ((v2) | 1))
    var v14 uint8 = 128
    if (!(true && (v3 || v3))) {
        var v15 float32 = (((float32(v1) + -2.25) + (1e10 / (0.1 + 3.14159))) + 0.1)
        v16 := GenS3{F0: v14, F1: 1000, F2: 128}
        _ = v15
        _ = v16
    }
    v1 = float64(1.5)
    v17 := GenS3{F0: v14, F1: -1, F2: 640}
    switch 127 {
    case -1:
        v18 := []int{-9223372036854775808, 42}
        _ = v18
    case 0:
        v1 += float64(1.5)
    default:
        for fi19 := 0; fi19 < 4; fi19++ {
            for fi20 := 0; fi20 < 15; fi20++ {
                v21 := GenS2{F0: fi19}
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    }
    fmt.Println(v3, v2, v2)
    var v22 int64 = v2
    v14 -= (v14 >> ((595) & 63))
    v23 := []int{-2147483648, -1, -562}
    v14 = v14
    if (-32768 <= 28) {
        v24 := GenS4{F0: 0, F1: 2, F2: 0}
        switch 42 {
        case 4294967295:
            if (uint(v14) == 42) {
                var v25 int8 = 0
                var v26 uint = 65535
                v26 += (v26 * 0)
                _ = v25
                _ = v26
            } else {
                var v27 float64 = float64(1.5)
                _ = v27
            }
        case -2147483648:
            for fi28 := 0; fi28 < 18; fi28++ {
                for fi29 := 0; fi29 < 17; fi29++ {
                    if v3 {
                        var v30 uint16 = 865
                        _ = v30
                    }
                    _ = fi29
                }
                var v31 int8 = -117
                var v32 uint16 = 660
                _ = fi28
                _ = v31
                _ = v32
            }
        default:
            v3 = v3
        }
        _ = v24
    } else {
        v33 := GenS4{F0: 127, F1: uint64(v14), F2: 255}
        v2 = (v2 ^ -1)
        _ = v33
    }
    var v34 int16 = -182
    for fi35 := 0; fi35 < 13; fi35++ {
        for fi36 := 0; fi36 < 5; fi36++ {
            v34 -= (v34 << ((524) & 63))
            _ = fi36
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v34
}
