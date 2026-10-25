package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int64 }
func genFn0() int64 {
    var v5 uint64 = 18446744073709551615
    v6 := []int{32768, 42, 4294967295}
    v5 -= 4294967295
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 string = "hello"
    v2 := make(map[int]string)
    var v3 int64 = -9223372036854775808
    var v4 int64 = -555
    v4 += (-(v3))
    v1 = (v1 + ((v1 + "hello") + (v1 + v1)))
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 bool = (-2.25 > 1.5)
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 1
    switch "a" {
    case "a":
        v10 := GenS1{F0: 4294967295, F1: 1}
        _ = v10
    case "hello":
        if ((0 < 1) || true) {
            switch v1 {
            case "":
                var v11 uint32 = 89
                _ = v11
            case "\n\t":
                var v12 int16 = -32768
                _ = v12
            case "a":
                if false {
                    v13 := []int{-1, 2147483647}
                    v9 -= (v9 - v9)
                    v9 = (^(v9))
                    _ = v13
                }
            default:
                switch (((((v1 + v1) + ("\n\t" + v1)) + v1) + v1) + (v1 + v1)) {
                case "\n\t":
                    for fi14 := 0; fi14 < 4; fi14++ {
                        if false {
                            v15 := []int{854, 32767, (fi14 ^ 901), (fi14 << ((498) & 63))}
                            _ = v15
                        }
                        v3 += (v4 / (((v4 & (v3 ^ (v3 * (v4 | v4))))) | 1))
                        var v16 int16 = 0
                        _ = fi14
                        _ = v16
                    }
                case "foo bar":
                    if (4294967295 >= 4294967295) {
                        var v17 int16 = 42
                        var v18 float64 = 0.0
                        _ = v17
                        _ = v18
                    } else {
                        v19 := []int{42, 587, -521, 4294967295}
                        _ = v19
                    }
                }
            }
            for fi20 := 0; fi20 < 16; fi20++ {
                for fi21 := 0; fi21 < 8; fi21++ {
                    v4 += (v4 + 259)
                    var v22 uint64 = 2
                    v23 := GenS1{F0: 65535, F1: 2}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                var v24 bool = false
                _ = fi20
                _ = v24
            }
            v4 += (v4 | v3)
        }
    default:
        fmt.Println(v3, v3, v4)
    }
    v4 = genFn0()
    v25 := []int{2147483647, 32767, -137}
    for fi26 := 0; fi26 < 4; fi26++ {
        v27 := v2[fi26]
        v28 := v2[int(v9)]
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        fmt.Println(v3, v4)
        v30 := GenS2{F0: v4}
        _ = fi29
        _ = v30
    }
    if len(v25) > 0 {
        v25[0] = 128
    }
    fmt.Println(v1)
    _ = len(v2)
    v31 := GenS1{F0: 1, F1: 0}
    v1 = (v1 + "foo bar")
    fmt.Println(v1, v4)
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 int32 = -2147483648
        v33 -= (v33 - v33)
        fmt.Println(v3, fi32, v9)
        _ = fi32
        _ = v33
    }
    v2[-1] = ""
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v9
}
