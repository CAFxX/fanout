package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint32 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 string; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int, p1 string) uint64 {
    v6 := []int{}
    fmt.Println(p1)
    p1 = p1
    var v7 uint8 = 42
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 464
}
func genFn1() uint16 {
    v8 := []int{32768, -787, 32768}
    switch -9223372036854775808 {
    case 2147483647:
        var v9 uint16 = 2
        _ = v9
    default:
        for fi10 := 0; fi10 < 11; fi10++ {
            if ("" <= "foo bar") {
                v11 := GenS1{F0: 127, F1: fi10, F2: 747}
                fmt.Println(fi10)
                switch (fi10 << ((65535) & 63)) {
                case 128:
                    fmt.Println(fi10)
                case 801:
                    var v12 int16 = -180
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 14; fi13++ {
                        var v14 uint = 18446744073709551615
                        var v15 float32 = -2.25
                        if (!("foo bar" < "")) {
                            v14 += v14
                            v16 := GenS1{F0: 127, F1: (^(fi13)), F2: 42}
                            _ = v16
                        }
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
                _ = v11
            } else {
                fmt.Println(fi10)
                var v17 uint16 = 1
                _ = v17
            }
            fmt.Println(fi10)
            var v18 int = -956
            _ = fi10
            _ = v18
        }
    }
    switch "a" {
    case "foo bar":
        var v19 uint = 341
        _ = v19
    default:
        if len(v8) > 0 {
            v8[0] = -767
        }
    }
    _ = v8
    return 909
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 127
    var v3 int = -9223372036854775808
    var v4 uint32 = 65535
    var v5 uint8 = 127
    fmt.Println(v5, v2)
    if ((-1 != 100) && (!false)) {
        var v20 int64 = -9223372036854775808
        _ = v20
    } else {
        v21 := GenS1{F0: (-(v2)), F1: (v3 * (^(v3))), F2: (v4 | v4)}
        _ = v21
    }
    if true {
        fmt.Println(v2)
    } else {
        var v22 int = v3
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 uint16 = 0
            v2 -= v2
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    v25 := []int{}
    fmt.Println(v5, v2, v3)
    if (!true) {
        switch 834 {
        case -975:
            for fi26 := 0; fi26 < 7; fi26++ {
                v2 = -128
                _ = fi26
            }
        }
        v27 := GenS1{F0: (v2 + (v2 + v2)), F1: v3, F2: v4}
        _ = v27
    } else {
        fmt.Println(v5, v3)
        var v28 float64 = float64(0.0)
        _ = v28
    }
    if (!((255 < 65535) || false)) {
        v29 := GenS2{F0: v3, F1: -286}
        for fi30 := 0; fi30 < 9; fi30++ {
            v2 = -18
            v31 := GenS3{F0: "\n\t", F1: 0, F2: v5}
            _ = fi30
            _ = v31
        }
        v3 = v3
        _ = v29
    } else {
        fmt.Println(v3, v3)
        fmt.Println(v2, v2)
        if (!false) {
            v32 := GenS4{F0: 188}
            _ = v32
        } else {
            fmt.Println(v3, v5, v4)
        }
    }
    var v33 float32 = 0.1
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    var v34 int32 = 84
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
