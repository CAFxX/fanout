package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ("a" + "\n\t")
    var v3 int = -9223372036854775808
    v4 := []string{(("" + "foo bar") + v2)}
    var v5 float32 = 0.0
    v6 := []int{}
    var v7 int64 = 42
    v7 = (^(v7))
    var v8 float32 = v5
    v9 := GenS1{F0: 0}
    if true {
        var v10 int16 = 479
        if true {
            v5 = ((v8 - v8) / v8)
            var v11 float32 = float32(((0.0 + (0.0 * 1.5)) / (0.1 + float64(1.5))))
            _ = v11
        } else {
            _ = len(v1)
            var v12 string = v2
            v13 := v1[913]
            _ = v12
            _ = v13
        }
        _ = v10
    } else {
        if ((-(v3)) <= v3) {
            v1[2147483648] = 629
            if (0 == 0) {
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 uint32 = 42
                    switch v2 {
                    case "":
                        _ = len(v1)
                    case "\n\t":
                        if (true && false) {
                            v6 = append(v6, fi14)
                            _ = len(v6)
                            v16 := v1[v15]
                            _ = v16
                        } else {
                            v8 -= (v8 + float32(1e10))
                        }
                    case "hello":
                        v17 := GenS2{F0: 2, F1: -786, F2: 255}
                        _ = v17
                    default:
                        var v18 string = "\n\t"
                        _ = v18
                    }
                    _ = fi14
                    _ = v15
                }
            } else {
                var v19 string = v2
                _ = v19
            }
            for fi20 := 0; fi20 < 19; fi20++ {
                delete(v1, 1)
                var v21 int64 = 0
                fmt.Println(v2, v5, v2)
                _ = fi20
                _ = v21
            }
        }
        var v22 uint16 = 42
        if (((148 >= 2) && (303 < 18446744073709551615)) && (510 >= 18446744073709551615)) {
            var v23 uint64 = 1
            v22 = (v22 ^ v22)
            v3 += v3
            _ = v23
        }
        _ = v22
    }
    if false {
        switch v3 {
        case 129:
            for fi24 := 0; fi24 < 4; fi24++ {
                v8 -= v5
                for fi25 := 0; fi25 < 6; fi25++ {
                    fmt.Println(v3, v7)
                    var v26 uint = 2
                    _ = len(v1)
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
        case 127:
            v7 -= v7
        case 128:
            v2 = v2
        default:
            fmt.Println(v2, v7)
        }
        v1[409] = 2147483647
    }
    var v27 float64 = (0.1 - 1.5)
    v28 := GenS1{F0: -49}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
