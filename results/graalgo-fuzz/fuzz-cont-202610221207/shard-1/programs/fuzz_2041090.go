package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int64 }
type GenS2 struct { F0 uint8; F1 int; F2 uint16 }
func genFn0(p0 int32, p1 uint8, p2 int) uint {
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := []int{(^(p2)), p2}
        p2 = fi4
        _ = fi4
        _ = v5
    }
    var v6 int32 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 702
}

func main() {
    var v1 bool = false
    var v2 uint = 4294967295
    var v3 int8 = 0
    var v7 uint32 = 1
    v2 -= (v2 & (^(v2)))
    v8 := []int{-737, 32768, -2147483648, 1}
    fmt.Println(v2, v2, v7)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        switch "hello" {
        case "":
            for fi10 := 0; fi10 < 6; fi10++ {
                switch ("\n\t" + "") {
                case "foo bar":
                    var v11 int64 = -9223372036854775808
                    _ = v11
                case "\n\t":
                    fmt.Println(v2, fi9)
                case "":
                    var v12 uint64 = 1
                    _ = v12
                default:
                    if (v7 != v7) {
                        var v13 uint = genFn0(1, 19, fi10)
                        v14 := GenS1{F0: 2, F1: genFn0(1, 42, fi9), F2: 1}
                        _ = v13
                        _ = v14
                    } else {
                        v15 := GenS1{F0: 128, F1: genFn0(0, 160, 9223372036854775807), F2: -9223372036854775808}
                        fmt.Println(v1, v3)
                        fmt.Println(v1, v1)
                        _ = v15
                    }
                }
                v3 += (v3 % ((v3) | 1))
                switch 650 {
                case -9223372036854775808:
                    v3 += v3
                case 32768:
                    v8 = append(v8, (-(fi10)))
                }
                _ = fi10
            }
        case "\n\t":
            if (-1 <= 1) {
                if (1000 == 1000) {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        var v17 uint8 = 2
                        var v18 string = ("hello" + "foo bar")
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    if len(v8) > 0 {
                        v8[0] = 42
                    }
                } else {
                    var v19 int8 = -128
                    var v20 uint8 = 150
                    _ = v19
                    _ = v20
                }
                v7 -= (v7 | (^(v7)))
            }
        case "a":
            if (("hello" + "foo bar") != "") {
                _ = len(v8)
            } else {
                v21 := GenS2{F0: 42, F1: (fi9 | fi9), F2: 800}
                _ = v21
            }
        }
        _ = fi9
    }
    v2 = (^(v2))
    v3 += v3
    if (v3 < 42) {
        var v22 uint = (v2 * v2)
        _ = v22
    } else {
        for fi23 := 0; fi23 < 5; fi23++ {
            for fi24 := 0; fi24 < 19; fi24++ {
                v3 = v3
                fmt.Println(v7, v1, v7)
                _ = fi24
            }
            _ = fi23
        }
        v25 := GenS2{F0: 194, F1: 128, F2: 32767}
        var v26 uint8 = 128
        _ = v25
        _ = v26
    }
    switch -806 {
    case -793:
        _ = len(v8)
    default:
        var v27 int8 = -36
        _ = v27
    }
    v7 = (v7 * v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
