package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int8; F2 string }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 2
    var v3 uint32 = 0
    v4 := []uint32{v3, (v3 ^ (v3 ^ (^(v3)))), v3}
    var v5 uint16 = 42
    v6 := []int{-299, 127}
    fmt.Println(v2, v3, v5)
    v7 := GenS1{F0: (v2 | uint(v5))}
    delete(v1, int(v2))
    var v8 bool = true
    if (-9223372036854775808 != -590) {
        switch "" {
        case "foo bar":
            fmt.Println(v3, v8)
        case "hello":
            v9 := GenS2{F0: 100, F1: -128, F2: ("hello" + "a")}
            _ = v9
        default:
            var v10 int32 = 373
            _ = v10
        }
        _ = len(v1)
    }
    _ = len(v1)
    var v11 int32 = -2147483648
    var v12 uint16 = 42
    if false {
        var v13 int = -9223372036854775808
        if (!v8) {
            if (!true) {
                v3 -= (v3 / (((v3 * 0)) | 1))
                switch "hello" {
                case "\n\t":
                    v14 := GenS2{F0: 0, F1: 80, F2: ("" + "hello")}
                    _ = v14
                case "foo bar":
                    var v15 int16 = -32768
                    _ = v15
                case "hello":
                    for fi16 := 0; fi16 < 16; fi16++ {
                        v8 = ((301 == int16(v11)) || true)
                        if (!(v11 < (v11 % (((^(v11))) | 1)))) {
                            v17 := GenS2{F0: 0, F1: -1, F2: "foo bar"}
                            _ = v17
                        }
                        _ = fi16
                    }
                }
            }
        } else {
            switch "hello" {
            case "\n\t":
                delete(v1, (v13 | v13))
            default:
                v2 = (v2 % (((v2 - 38)) | 1))
            }
            v12 += v5
            var v18 uint64 = 1
            _ = v18
        }
        _ = v13
    } else {
        v19 := GenS2{F0: 42, F1: -12, F2: "foo bar"}
        _ = v19
    }
    var v20 uint64 = 18446744073709551615
    var v21 int16 = int16(v2)
    if (!((((v11 > v11) && v8) || true) || (-2.25 == (0.0 / ((1.5 - 0.1) / 1.0))))) {
        if len(v6) > 0 {
            v6[0] = -211
        }
        var v22 uint = (v2 ^ (v2 + v2))
        _ = v22
    }
    v20 += 584
    v23 := v1[2147483647]
    v24 := GenS1{F0: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
