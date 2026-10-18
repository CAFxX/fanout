package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint64; F1 uint64; F2 string }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int32, p2 int32) int64 {
    for fi5 := 0; fi5 < 9; fi5++ {
        fmt.Println(p0)
        var v6 int64 = 363
        _ = fi5
        _ = v6
    }
    fmt.Println(p2, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 float32 = (0.0 * 0.0)
    var v2 bool = (676 != 411)
    var v3 uint32 = 1
    var v4 uint32 = (v3 ^ (v3 + v3))
    var v7 string = ("" + "a")
    v1 = 1.5
    if v2 {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := []int{0}
            v7 = "a"
            v4 = v4
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 584}
        _ = v10
    } else {
        if v2 {
            var v11 uint32 = (^(v3))
            _ = v11
        }
    }
    fmt.Println(v1, v7, v7)
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint = 1
        v4 = 2147483647
        switch 2147483647 {
        case -9223372036854775808:
            for fi14 := 0; fi14 < 16; fi14++ {
                v15 := []int{127}
                v16 := GenS2{F0: 127, F1: 138, F2: 179}
                if (!true) {
                    v17 := GenS3{F0: 285, F1: 1000}
                    if (!v2) {
                        v18 := GenS4{F0: (v13 >> (((v13 + 125)) & 63)), F1: 1}
                        v4 = v3
                        _ = v18
                    } else {
                        var v19 uint64 = 0
                        if ((uint16(v19) != 658) || v2) {
                            var v20 bool = true
                            _ = v20
                        }
                        var v21 int8 = 127
                        v22 := GenS1{F0: 204}
                        v23 := GenS5{F0: (v19 >> ((18446744073709551615) & 63)), F1: v19, F2: v7}
                        _ = v19
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = v17
                } else {
                    switch v7 {
                    case "hello":
                        v3 -= v4
                    case "foo bar":
                        v2 = false
                    default:
                        var v24 int8 = -128
                        _ = v24
                    }
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        case -654:
            v25 := GenS3{F0: int16(v13), F1: 1000}
            _ = v25
        case 2147483647:
            var v26 uint = v13
            _ = v26
        }
        _ = fi12
        _ = v13
    }
    v27 := GenS3{F0: -195, F1: 42}
    var v28 uint32 = v3
    if (!(genFn0(1, -2147483648, 0) >= genFn0(42, 469, 1))) {
        v1 = (v1 * v1)
        v29 := []int{2147483647, 2147483647, -9223372036854775808, -1}
        _ = v29
    } else {
        v28 = 42
        for fi30 := 0; fi30 < 6; fi30++ {
            fmt.Println(v4, v1)
            _ = fi30
        }
    }
    v1 += v1
    if v2 {
        var v31 float32 = (float32((0.0 - float64(3.14159))) * v1)
        var v32 float64 = -2.25
        _ = v31
        _ = v32
    }
    v33 := GenS6{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v7
}
