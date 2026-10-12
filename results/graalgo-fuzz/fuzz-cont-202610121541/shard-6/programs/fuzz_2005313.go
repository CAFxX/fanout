package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }

func main() {
    var v1 int16 = -513
    v2 := []int{9223372036854775807, -9223372036854775808, -280, -408}
    var v3 uint = 1
    var v4 int = -1
    var v5 int64 = 0
    v6 := make(map[int]int)
    if (v3 >= v3) {
        v3 = (^(v3))
        _ = len(v6)
        v4 -= (v4 >> ((18446744073709551615) & 63))
    } else {
        switch -9223372036854775808 {
        case -127:
            v3 = 227
        case 735:
            if len(v2) > 0 {
                v2[0] = (-(v4))
            }
        case 42:
            fmt.Println(v1, v1)
        }
    }
    v4 -= v4
    v7 := GenS1{F0: v4, F1: 18446744073709551615}
    var v8 bool = true
    var v9 int16 = -83
    var v10 float64 = 1e10
    v11 := GenS1{F0: 32768, F1: 18446744073709551615}
    for fi12 := 0; fi12 < 14; fi12++ {
        if v8 {
            v10 += ((float64(1.5) + (v10 + 1e10)) - 1.5)
            var v13 uint8 = 126
            v14 := v6[v4]
            _ = v13
            _ = v14
        }
        v9 += (v9 >> (((^(v3))) & 63))
        switch -553 {
        case -9223372036854775808:
            delete(v6, (fi12 ^ 127))
        case -2147483648:
            for fi15 := 0; fi15 < 15; fi15++ {
                switch "a" {
                case "foo bar":
                    var v16 int16 = 303
                    _ = v16
                default:
                    v9 += v9
                }
                switch 1 {
                case -738:
                    _ = len(v6)
                case 0:
                    switch (^(fi15)) {
                    case 128:
                        v5 = 162
                        switch "foo bar" {
                        case "a":
                            v1 = (-(v9))
                        }
                    default:
                        fmt.Println(v3, v9)
                    }
                case 4294967295:
                    v17 := GenS1{F0: v4, F1: 4}
                    _ = v17
                default:
                    var v18 uint32 = 2
                    _ = v18
                }
                _ = fi15
            }
        }
        _ = fi12
    }
    var v19 float64 = (float64(((-2.25 + float32(v10)) - -2.25)) / v10)
    var v20 int64 = (v5 | v5)
    var v21 uint32 = 4294967295
    v4 += (v4 >> ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
