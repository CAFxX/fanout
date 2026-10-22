package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int16 = -1
    var v2 float64 = (-2.25 - 1e10)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]bool)
    var v5 int64 = 42
    for fi6 := 0; fi6 < 3; fi6++ {
        v2 = v2
        _ = fi6
    }
    if (1 == 18446744073709551615) {
        if (false && false) {
            v7 := GenS1{F0: (-(v5)), F1: -128}
            fmt.Println(v2, v5, v2)
            _ = v7
        }
        v1 = v1
        _ = len(v4)
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        v5 = (v5 % (((v5 >> ((18446744073709551615) & 63))) | 1))
        for fi9 := 0; fi9 < 11; fi9++ {
            v2 -= v2
            var v10 uint32 = 769
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := GenS1{F0: (v5 ^ (v5 >> ((18446744073709551615) & 63))), F1: -31}
    v1 += v1
    var v12 uint = 236
    v13 := GenS2{F0: (^(v5))}
    var v14 uint32 = 2147483648
    var v15 string = "hello"
    var v16 uint32 = 65535
    switch -654 {
    case 498:
        var v17 int64 = 731
        _ = v17
    case 9223372036854775807:
        if true {
            var v18 float32 = 1.5
            if (887 <= 42) {
                if false {
                    switch v15 {
                    case "\n\t":
                        var v19 uint8 = 2
                        _ = v19
                    case "a":
                        v5 = (-(v5))
                    default:
                        var v20 int32 = -1
                        _ = v20
                    }
                }
            }
            fmt.Println(v14)
            _ = v18
        }
    default:
        fmt.Println(v2, v15)
    }
    var v21 float32 = 1e10
    var v22 uint = v12
    v23 := []int{}
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 string = (v15 + "a")
        v26 := GenS2{F0: v5}
        v15 = ("a" + "foo bar")
        _ = fi24
        _ = v25
        _ = v26
    }
    switch v15 {
    case "a":
        for fi27 := 0; fi27 < 19; fi27++ {
            if false {
                var v28 int8 = -1
                _ = v28
            }
            v14 = (v14 * v14)
            var v29 float64 = ((1.5 + v2) - v2)
            _ = fi27
            _ = v29
        }
    default:
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
