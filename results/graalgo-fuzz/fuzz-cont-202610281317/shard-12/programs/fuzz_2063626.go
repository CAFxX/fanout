package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int }

func main() {
    v1 := []string{"a", "hello", "\n\t", "a"}
    v2 := []string{"hello", "foo bar", (("" + "a") + "")}
    v3 := make(map[uint32]uint32)
    var v4 float32 = 3.14159
    var v5 int8 = 100
    var v6 uint32 = 491
    var v7 string = "\n\t"
    var v8 uint8 = 127
    if true {
        v8 = (^(v8))
        v4 += v4
    }
    delete(v3, (v6 | 684))
    fmt.Println(v4, v7, v6)
    v9 := GenS1{F0: (v8 * (v8 >> ((2) & 63))), F1: (v5 % ((v5) | 1))}
    switch -1 {
    case -723:
        var v10 uint64 = 138
        _ = v10
    }
    if (true && (0.0 != v4)) {
        if ((^(v6)) == v6) {
            var v11 int64 = 0
            _ = v11
        } else {
            v8 += 42
            v8 -= (^(v8))
        }
        var v12 int32 = 42
        switch -70 {
        case -710:
            for fi13 := 0; fi13 < 4; fi13++ {
                if (75 < v8) {
                    if ((float32(-2.25) / ((v4 * v4) / (v4 + 1e10))) >= (v4 - (float32(1e10) * (v4 / v4)))) {
                        v5 -= int8(fi13)
                    }
                } else {
                    if ((395 >= 9223372036854775807) && false) {
                        v5 -= (^(v5))
                        var v14 int32 = v12
                        for fi15 := 0; fi15 < 12; fi15++ {
                            v6 += 678
                            _ = fi15
                        }
                        _ = v14
                    }
                }
                _ = fi13
            }
        case 32767:
            switch 128 {
            case 1:
                _ = len(v1)
            case -43:
                v6 = v6
            }
        case 127:
            v6 = v6
        default:
            v8 += v8
        }
        _ = v12
    } else {
        if true {
            var v16 float64 = 0.0
            var v17 bool = false
            v17 = false
            _ = v16
            _ = v17
        }
        v18 := GenS2{F0: (v5 << ((4294967295) & 63)), F1: false}
        _ = v18
    }
    v19 := GenS3{F0: 328}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
