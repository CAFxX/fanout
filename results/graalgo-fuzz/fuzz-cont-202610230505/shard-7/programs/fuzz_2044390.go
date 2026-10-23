package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 780
    v3 := make(map[string]bool)
    for fi4 := 0; fi4 < 19; fi4++ {
        delete(v3, "foo bar")
        var v5 int16 = 32767
        _ = fi4
        _ = v5
    }
    v1 = v1
    if (v1 > (v1 & 0)) {
        v1 += (v1 - (v1 * v1))
        var v6 int32 = 2147483647
        v1 = 1
        _ = v6
    } else {
        var v7 uint8 = 119
        if (2147483648 < 0) {
            if false {
                v8 := []int{-1}
                if (100 <= -1) {
                    v3[""] = true
                    if len(v8) > 0 {
                        v8[0] = 4294967295
                    }
                }
                var v9 uint32 = uint32(v2)
                _ = v8
                _ = v9
            } else {
                switch "" {
                case "":
                    var v10 uint32 = 2147483648
                    _ = v10
                default:
                    if false {
                        v2 += (v2 ^ (^(v2)))
                        var v11 int16 = 0
                        fmt.Println(v2, v2, v7)
                        _ = v11
                    }
                }
                fmt.Println(v1)
                if (1.5 <= (0.0 - (((1.5 / float64(((float32(1e10) + 0.0) - 0.1))) * 0.1) / 0.1))) {
                    v1 = (^(v1))
                }
            }
            v12 := v3["foo bar"]
            if (1 <= 9223372036854775807) {
                fmt.Println(v1, v12)
                if true {
                    if (32767 < -32768) {
                        var v13 int64 = 9223372036854775807
                        v12 = (42 < -2147483648)
                        var v14 int64 = -1
                        _ = v13
                        _ = v14
                    }
                } else {
                    fmt.Println(v7)
                    var v15 bool = (137 < uint8(v2))
                    _ = v15
                }
            } else {
                fmt.Println(v7)
            }
            _ = v12
        }
        _ = v7
    }
    fmt.Println(v1, v1)
    v1 = v1
    v1 -= (v1 | 0)
    v16 := GenS1{F0: v1, F1: 933}
    v3["hello"] = false
    switch 32768 {
    case 0:
        var v17 int32 = 0
        _ = v17
    case -1:
        var v18 float64 = float64((-2.25 * 1e10))
        _ = v18
    default:
        v19 := []int{-546, 799}
        _ = v19
    }
    var v20 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
}
