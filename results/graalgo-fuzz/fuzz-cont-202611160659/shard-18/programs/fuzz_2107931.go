package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int; F2 uint }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int16, p1 bool, p2 uint8) uint8 {
    fmt.Println(p0, p0, p0)
    var v5 uint64 = 42
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{}
        p1 = false
        var v8 uint64 = (^(v5))
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{42, 245, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return (p2 - (^(p2)))
}
func genFn1(p0 int16) bool {
    v10 := GenS1{F0: 34, F1: 42, F2: 127}
    var v11 uint = 0
    var v12 uint8 = genFn0(p0, true, 101)
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return false
}

func main() {
    var v1 uint32 = 0
    var v2 uint = 81
    var v3 bool = (542 < 16)
    v4 := make(map[string]bool)
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            if false {
                switch "" {
                case "foo bar":
                    if (!(true && v3)) {
                        var v15 int = 437
                        _ = v15
                    }
                case "a":
                    var v16 float64 = 1.5
                    _ = v16
                default:
                    var v17 uint64 = 499
                    v18 := GenS1{F0: 100, F1: 56, F2: 128}
                    v17 -= (v17 % ((v17) | 1))
                    _ = v17
                    _ = v18
                }
            } else {
                v19 := GenS1{F0: -128, F1: 128, F2: 165}
                var v20 int16 = -32768
                _ = v19
                _ = v20
            }
            _ = fi14
        }
        if v3 {
            if v3 {
                switch fi13 {
                case 4294967295:
                    v21 := GenS2{F0: 289, F1: fi13, F2: v2}
                    _ = v21
                case -572:
                    if (float64(1e10) > (float64((((((0.1 + 3.14159) + 1.5) - 3.14159) + 1e10) + 1e10)) + -2.25)) {
                        delete(v4, "foo bar")
                        if (224 >= 2) {
                            v1 = v1
                        } else {
                            var v22 bool = (-2.25 == (float32(0.1) / 1.0))
                            _ = v22
                        }
                        if v3 {
                            v1 -= (v1 - 386)
                        }
                    }
                }
                var v23 uint16 = 957
                v1 -= v1
                _ = v23
            }
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 string = "foo bar"
            v26 := GenS2{F0: 336, F1: fi13, F2: (^(v2))}
            v27 := []int{}
            v4[((v25 + (v25 + v25)) + "\n\t")] = (-1 != 0)
            fmt.Println(fi13, fi24)
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi13
    }
    fmt.Println(v3)
    v28 := GenS2{F0: 877, F1: 32768, F2: v2}
    var v29 uint64 = 0
    fmt.Println(v29, v29, v1)
    switch "\n\t" {
    case "a":
        v2 += (v2 * 658)
    case "hello":
        switch 32767 {
        case -773:
            v30 := []int{207, 2147483647, 532}
            _ = v30
        case 32767:
            switch "" {
            case "\n\t":
                v3 = false
            }
        case 128:
            fmt.Println(v29)
        default:
            v31 := GenS3{F0: 1}
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        v4["\n\t"] = ((v1 | 65535) < (v1 >> ((v2) & 63)))
        fmt.Println(fi32)
        _ = fi32
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
