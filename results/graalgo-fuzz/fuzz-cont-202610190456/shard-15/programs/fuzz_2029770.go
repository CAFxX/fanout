package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 bool }
func genFn0() bool {
    var v7 int8 = 1
    for fi8 := 0; fi8 < 11; fi8++ {
        v7 = (v7 & (-(v7)))
        _ = fi8
    }
    var v9 int32 = -2147483648
    switch 4294967295 {
    case 0:
        v10 := []int{}
        _ = v10
    default:
        v9 = (v9 ^ (^(v9)))
    }
    v11 := []int{}
    _ = v11
    _ = v7
    _ = v9
    return true
}
func genFn1(p0 int32, p1 string) bool {
    p1 = ""
    fmt.Println(p1, p0)
    fmt.Println(p1, p0)
    if (32767 != 1) {
        v12 := GenS1{F0: (p0 << ((768) & 63)), F1: 249}
        _ = v12
    }
    _ = p0
    _ = p1
    return (0.0 != (((-2.25 + 0.0) / (float64(1.5) + (1.5 - 0.1))) + 1e10))
}

func main() {
    v1 := []int{}
    v2 := []uint32{2}
    var v3 bool = false
    var v4 float64 = (float64((-2.25 - 1.5)) + -2.25)
    var v5 int8 = 31
    v6 := make(map[uint32]string)
    if false {
        fmt.Println(v5, v3, v3)
    } else {
        switch 127 {
        case 0:
            if ("foo bar" > "a") {
                _ = len(v6)
                var v13 int8 = (v5 << ((65535) & 63))
                var v14 bool = (false && v3)
                _ = v13
                _ = v14
            }
        case 4294967295:
            var v15 uint16 = 255
            _ = v15
        case 1:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        switch ("a" + "hello") {
        case "":
            v17 := GenS2{F0: 976}
            _ = v17
        case "a":
            fmt.Println(v4)
        }
        v4 = float64(float32(((v4 * -2.25) * v4)))
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := GenS1{F0: 864, F1: 127}
            v3 = false
            _ = fi18
            _ = v19
        }
        _ = fi16
    }
    var v20 uint64 = 1
    v6[903] = ("hello" + "")
    v21 := GenS3{F0: 1}
    if false {
        for fi22 := 0; fi22 < 12; fi22++ {
            if false {
                var v23 float64 = (-2.25 / ((((float64(float32(0.0)) + v4) * v4) / v4) - v4))
                var v24 uint = 1
                _ = v23
                _ = v24
            }
            for fi25 := 0; fi25 < 9; fi25++ {
                if (1.5 <= 3.14159) {
                    v26 := GenS4{F0: fi22}
                    v5 -= (-(v5))
                    var v27 bool = (4294967295 != 0)
                    _ = v26
                    _ = v27
                }
                v4 -= (float64((1e10 / (0.0 * float32(v4)))) - 0.1)
                var v28 int64 = -9223372036854775808
                _ = fi25
                _ = v28
            }
            v29 := GenS5{F0: 362, F1: v3}
            _ = fi22
            _ = v29
        }
        v30 := GenS2{F0: 42}
        _ = v30
    } else {
        switch -9223372036854775808 {
        case -1:
            _ = len(v2)
        default:
            fmt.Println(v4)
        }
    }
    v20 -= (v20 | v20)
    var v31 int = 9223372036854775807
    var v32 uint = 335
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
