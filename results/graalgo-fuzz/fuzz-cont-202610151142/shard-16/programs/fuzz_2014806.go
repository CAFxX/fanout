package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int64; F1 int16; F2 uint8 }
type GenS4 struct { F0 bool }
func genFn0(p0 int8) uint {
    if (9223372036854775807 >= -9223372036854775808) {
        p0 = (^(p0))
    } else {
        p0 -= p0
    }
    p0 = 42
    var v4 int32 = 42
    var v5 uint16 = 65535
    _ = p0
    _ = v4
    _ = v5
    return 2
}
func genFn1() uint8 {
    var v6 string = (("a" + "foo bar") + ("\n\t" + "a"))
    for fi7 := 0; fi7 < 4; fi7++ {
        if false {
            for fi8 := 0; fi8 < 4; fi8++ {
                fmt.Println(fi8, fi7)
                _ = fi8
            }
            for fi9 := 0; fi9 < 6; fi9++ {
                if (false || (65535 > 67)) {
                    var v10 string = ("" + v6)
                    var v11 int32 = 2147483647
                    _ = v10
                    _ = v11
                } else {
                    switch 0 {
                    case 128:
                        fmt.Println(fi9, fi7, v6)
                    case 18:
                        if false {
                            v12 := GenS1{F0: 9223372036854775807}
                            v6 = v6
                            var v13 uint = 437
                            var v14 uint64 = 1
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    case 127:
                        v6 = ""
                    default:
                        var v15 int = fi9
                        _ = v15
                    }
                }
                v6 = "a"
                _ = fi9
            }
            v16 := []int{(-(fi7)), (fi7 >> ((4294967295) & 63))}
            _ = v16
        } else {
            var v17 int64 = 42
            _ = v17
        }
        v6 = ""
        v18 := GenS2{F0: 42, F1: 17}
        _ = fi7
        _ = v18
    }
    var v19 uint32 = 735
    v20 := []int{128, -691}
    _ = v19
    _ = v20
    _ = v6
    return 2
}

func main() {
    var v1 float64 = ((1e10 - 0.0) / 1.0)
    var v2 int64 = 130
    var v3 int64 = v2
    fmt.Println(v3, v1)
    v21 := GenS1{F0: 490}
    v22 := GenS3{F0: v2, F1: 42, F2: 42}
    v1 -= float64(1.5)
    v3 = 1
    v2 += (^(v3))
    for fi23 := 0; fi23 < 9; fi23++ {
        if (4294967295 >= 2) {
            v1 -= v1
            v2 = (v3 * (v3 ^ (v3 << ((203) & 63))))
            for fi24 := 0; fi24 < 20; fi24++ {
                switch "hello" {
                case "\n\t":
                    v25 := GenS4{F0: ("a" < ("hello" + ""))}
                    _ = v25
                }
                _ = fi24
            }
        }
        fmt.Println(v1, v2)
        _ = fi23
    }
    v3 = 534
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
