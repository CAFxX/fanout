package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 bool }
type GenS2 struct { F0 int32; F1 bool; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 int8; F2 int16 }
func genFn0(p0 uint64, p1 uint16, p2 string) int {
    var v7 uint64 = (p0 ^ (^(p0)))
    v8 := []int{2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return int(p1)
}
func genFn1() bool {
    if ("\n\t" != (("" + "a") + "a")) {
        var v9 uint64 = 2
        _ = v9
    }
    if (-1 == -1) {
        v10 := []int{262}
        _ = v10
    }
    return true
}

func main() {
    var v1 int32 = -571
    var v2 string = ("foo bar" + "hello")
    var v3 float64 = 0.0
    v4 := make(map[uint32]uint32)
    var v5 string = v2
    var v6 uint64 = 0
    v1 = (v1 & (-(v1)))
    v11 := GenS1{F0: 4294967295, F1: 42, F2: (!true)}
    fmt.Println(v2, v1)
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 string = "hello"
        var v14 float64 = 0.1
        _ = fi12
        _ = v13
        _ = v14
    }
    switch -129 {
    case -1:
        switch (("a" + "hello") + v5) {
        case "":
            var v15 float32 = float32(((v3 / (3.14159 * ((v3 / 1.5) + (((float64(0.0) * float64(0.1)) + (v3 + v3)) - ((0.0 / 0.1) + 1e10))))) - v3))
            _ = v15
        }
    case -473:
        v1 += (v1 ^ v1)
    case -883:
        delete(v4, 2147483647)
    default:
        fmt.Println(v3, v6)
    }
    switch -401 {
    case 2147483647:
        switch v2 {
        case "":
            v16 := GenS2{F0: v1, F1: (false || false), F2: 42}
            _ = v16
        default:
            v17 := GenS2{F0: -1, F1: (100 >= -47), F2: 32767}
            _ = v17
        }
    case -38:
        var v18 string = "a"
        _ = v18
    case 9223372036854775807:
        v19 := GenS3{F0: v1}
        _ = v19
    }
    fmt.Println(v1, v2)
    switch "\n\t" {
    case "\n\t":
        for fi20 := 0; fi20 < 8; fi20++ {
            v4[65535] = 0
            _ = fi20
        }
    case "a":
        v3 += (float64(3.14159) + float64(float32((float64((float32(0.1) / (((-2.25 * 0.0) - (1.5 - float32((((1e10 - 3.14159) - (((0.0 - (v3 + 1.5)) / 1.0) + ((-2.25 / float64((1e10 / 1e10))) / 1.0))) / 1.0)))) - (0.0 / 1.5)))) * ((0.0 / (v3 / (v3 - (0.0 / 0.1)))) - v3)))))
    case "":
        if (v2 == "\n\t") {
            fmt.Println(v2, v6, v5)
            if true {
                var v21 int = -2147483648
                var v22 uint64 = 4294967295
                v23 := GenS4{F0: 145, F1: 0, F2: 1000}
                _ = v21
                _ = v22
                _ = v23
            } else {
                var v24 bool = false
                _ = v24
            }
        }
    }
    var v25 float32 = 1e10
    v6 = (v6 << ((18446744073709551615) & 63))
    if (18446744073709551615 == 1) {
        v26 := GenS3{F0: (v1 * (v1 ^ (v1 / ((-2147483648) | 1))))}
        var v27 string = ""
        v28 := GenS4{F0: 32767, F1: -10, F2: 875}
        _ = v26
        _ = v27
        _ = v28
    } else {
        v29 := []int{-968}
        v6 = (v6 & 2)
        var v30 int = 4294967295
        _ = v29
        _ = v30
    }
    v31 := []int{-9223372036854775808}
    v31 = append(v31, 32768)
    var v32 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
