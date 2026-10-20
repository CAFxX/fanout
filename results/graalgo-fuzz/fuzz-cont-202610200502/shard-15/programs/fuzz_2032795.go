package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 int32 }

func main() {
    var v1 int64 = -1
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 uint8 = 210
    if false {
        if true {
            _ = len(v3)
        }
        var v5 bool = (!(17 != 2))
        _ = v5
    }
    v6 := GenS1{F0: 552, F1: 321, F2: v1}
    for fi7 := 0; fi7 < 16; fi7++ {
        if (v1 <= v1) {
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := GenS1{F0: 1, F1: 255, F2: 42}
                if (!(0.0 < 1.5)) {
                    var v10 int = (fi8 % ((fi8) | 1))
                    _ = v10
                }
                switch "" {
                case "a":
                    v1 -= v1
                default:
                    if false {
                        v11 := []int{int(v1)}
                        v12 := GenS2{F0: fi7}
                        _ = v11
                        _ = v12
                    } else {
                        var v13 bool = (1 == 61)
                        fmt.Println(fi8, fi7)
                        v14 := GenS3{F0: 0, F1: 42}
                        _ = len(v3)
                        _ = v13
                        _ = v14
                    }
                }
                _ = fi8
                _ = v9
            }
        } else {
            v1 = (v1 * (v1 / (((v1 >> ((466) & 63))) | 1)))
        }
        var v15 int32 = 164
        var v16 float64 = 3.14159
        _ = fi7
        _ = v15
        _ = v16
    }
    v4 -= v4
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS4{F0: 0, F1: 127}
        var v19 float64 = ((float64((3.14159 / 0.1)) + float64(0.1)) / (-2.25 / 1.0))
        _ = fi17
        _ = v18
        _ = v19
    }
    if (((v1 > (-(v1))) || false) && false) {
        fmt.Println(v1)
        v4 -= v4
    }
    v20 := GenS1{F0: 18446744073709551615, F1: 42, F2: v1}
    var v21 float64 = (3.14159 * 0.0)
    if false {
        var v22 bool = true
        if (v1 != (^(v1))) {
            switch 2147483647 {
            case 42:
                fmt.Println(v1, v21, v1)
            case -9223372036854775808:
                v22 = v22
            case -1:
                v23 := GenS4{F0: 2, F1: v4}
                _ = v23
            }
            delete(v3, 1)
            v21 += v21
        } else {
            fmt.Println(v1, v21, v1)
            v24 := GenS5{F0: "", F1: v4}
            delete(v3, -9223372036854775808)
            _ = v24
        }
        if (!true) {
            var v25 int8 = -128
            var v26 int64 = -9223372036854775808
            switch "hello" {
            case "\n\t":
                if (65535 > 2147483647) {
                    fmt.Println(v21, v26)
                } else {
                    v27 := GenS6{F0: 42}
                    _ = v27
                }
            }
            _ = v25
            _ = v26
        } else {
            var v28 int32 = 134
            _ = v28
        }
        _ = v22
    } else {
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 uint64 = 18446744073709551615
            var v31 uint16 = 65535
            _ = fi29
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 float64 = v21
            _ = fi32
            _ = v33
        }
    }
    var v34 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v6
}
