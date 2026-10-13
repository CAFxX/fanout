package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int32; F2 uint }
type GenS3 struct { F0 int8; F1 uint32; F2 int32 }
func genFn0(p0 uint64) int64 {
    if (true || (1e10 < 0.1)) {
        var v4 float64 = float64(((((1e10 / (3.14159 / 1e10)) + 0.1) + 0.0) - ((0.1 - float32((1e10 / -2.25))) / 1.0)))
        _ = v4
    } else {
        var v5 uint64 = (^(p0))
        var v6 uint = 1
        var v7 int16 = -32768
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := []int{-2147483648}
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 0
}
func genFn1() int {
    v9 := GenS1{F0: false}
    v10 := GenS1{F0: true}
    v11 := GenS2{F0: 127, F1: 42, F2: 255}
    switch ("a" + ("\n\t" + (("" + "") + "foo bar"))) {
    case "foo bar":
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(fi12)
            var v13 uint16 = 189
            _ = fi12
            _ = v13
        }
    case "":
        var v14 float32 = -2.25
        _ = v14
    }
    _ = v10
    _ = v11
    _ = v9
    return 32767
}

func main() {
    var v1 uint8 = 235
    v2 := []string{}
    v3 := []uint32{}
    v1 -= 127
    if (1.5 != 1.5) {
        _ = len(v2)
        switch -711 {
        case 9223372036854775807:
            for fi15 := 0; fi15 < 8; fi15++ {
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 bool = (((!(0.1 >= 1e10)) || false) || false)
                    switch "foo bar" {
                    case "\n\t":
                        if v17 {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case "hello":
                        var v18 bool = (1.5 <= ((((float32(1.5) - -2.25) * 0.1) * 0.1) - (0.0 - 0.0)))
                        _ = v18
                    case "":
                        v1 -= (v1 << ((1) & 63))
                    default:
                        fmt.Println(fi15)
                    }
                    _ = fi16
                    _ = v17
                }
                v1 -= (v1 * (v1 / ((v1) | 1)))
                _ = fi15
            }
        case 32768:
            var v19 int64 = 0
            _ = v19
        case 42:
            if len(v2) > 0 {
                v2[0] = (("a" + "\n\t") + "hello")
            }
        default:
            if ((((113 > (v1 + 0)) || (genFn0(572) >= genFn0(42))) || true) || false) {
                v1 = (v1 - 0)
            } else {
                v1 -= v1
                var v20 int64 = 0
                _ = v20
            }
        }
    } else {
        var v21 int32 = 2147483647
        _ = v21
    }
    var v22 int = 585
    v22 = v22
    v1 = (v1 + (^(v1)))
    for fi23 := 0; fi23 < 9; fi23++ {
        v22 += fi23
        for fi24 := 0; fi24 < 10; fi24++ {
            v1 = (v1 ^ v1)
            _ = fi24
        }
        _ = fi23
    }
    v25 := GenS1{F0: true}
    for fi26 := 0; fi26 < 11; fi26++ {
        if (false && false) {
            if false {
                v27 := GenS2{F0: 1, F1: 898, F2: 255}
                v22 = (^(v22))
                var v28 bool = (1.5 > -2.25)
                _ = v27
                _ = v28
            } else {
                var v29 string = ("hello" + (("\n\t" + ("" + ("hello" + "\n\t"))) + "foo bar"))
                _ = v29
            }
            v30 := GenS3{F0: 42, F1: 65535, F2: int32(v22)}
            var v31 string = ("hello" + "\n\t")
            _ = v30
            _ = v31
        }
        _ = fi26
    }
    fmt.Println(v22)
    var v32 int = v22
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v32
}
