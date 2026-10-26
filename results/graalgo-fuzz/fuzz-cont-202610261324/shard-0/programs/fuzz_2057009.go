package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int) uint {
    if true {
        p0 = (p0 >> ((2) & 63))
        fmt.Println(p0)
    }
    p0 = (p0 << ((255) & 63))
    _ = p0
    return 42
}
func genFn1(p0 int32) uint32 {
    switch 762 {
    case -265:
        v7 := GenS1{F0: 587, F1: 128, F2: 0}
        _ = v7
    case 2147483647:
        if true {
            v8 := []int{0}
            var v9 int = -9223372036854775808
            _ = v8
            _ = v9
        }
    case -59:
        v10 := []int{32768, 2147483647, 127}
        _ = v10
    }
    switch 42 {
    case -1:
        v11 := GenS1{F0: 0, F1: 1, F2: 100}
        _ = v11
    default:
        v12 := []int{-2147483648, 1, 1, -786}
        _ = v12
    }
    v13 := GenS1{F0: 228, F1: 255, F2: -128}
    var v14 uint = 65535
    p0 = (p0 % ((p0) | 1))
    _ = p0
    _ = v13
    _ = v14
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 1.5
    var v3 float32 = -2.25
    var v4 float64 = (v2 + (((3.14159 - 0.0) / v2) + ((((v2 + (((v2 + 1.5) - v2) / (v2 / -2.25))) * v2) / 1e10) - (1e10 + v2))))
    v5 := make(map[int]uint32)
    var v6 uint64 = 42
    if (v4 != 1.5) {
        var v15 int64 = 0
        v16 := GenS1{F0: (v6 >> ((18446744073709551615) & 63)), F1: 42, F2: 127}
        var v17 int16 = -248
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 uint32 = 166
    v4 = v2
    fmt.Println(v6, v6)
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 uint64 = 18446744073709551615
        var v21 uint32 = (v18 | (v18 ^ v18))
        var v22 uint32 = (v21 + (v18 - 2))
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: 42, F1: 127, F2: 49}
    if (!true) {
        var v24 int64 = 1
        v25 := []int{}
        _ = v24
        _ = v25
    } else {
        if false {
            var v26 int16 = 42
            var v27 int8 = -36
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 uint8 = 42
                var v30 uint32 = (v18 - genFn1(950))
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = v26
            _ = v27
        }
    }
    switch "hello" {
    case "":
        v3 = 0.0
    case "foo bar":
        switch 127 {
        case -37:
            v31 := GenS2{F0: 1000}
            _ = v31
        case 127:
            v5[32767] = (v18 & (v18 ^ 697))
        case 519:
            for fi32 := 0; fi32 < 14; fi32++ {
                if true {
                    v6 += v6
                    for fi33 := 0; fi33 < 12; fi33++ {
                        v4 -= 0.1
                        var v34 int = 4294967295
                        v18 -= v18
                        _ = fi33
                        _ = v34
                    }
                } else {
                    var v35 int64 = 42
                    if (0.0 < v3) {
                        var v36 int64 = -372
                        _ = v36
                    }
                    _ = v35
                }
                v37 := v5[88]
                v38 := GenS1{F0: 2, F1: 0, F2: -23}
                _ = fi32
                _ = v37
                _ = v38
            }
        default:
            var v39 uint64 = 811
            _ = v39
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
