package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint; F1 int; F2 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    var v3 int = 127
    v4 := []int{(v3 & (v3 | 32768)), v3, v3}
    if len(v4) > 0 {
        v4[0] = -272
    }
    v3 -= v3
    v3 += v3
    switch "" {
    case "hello":
        if (1 < 1) {
            v5 := GenS1{F0: 105, F1: 18446744073709551615, F2: 42}
            _ = v5
        }
    default:
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 uint8 = 128
            switch "\n\t" {
            case "\n\t":
                switch "a" {
                case "":
                    var v8 uint64 = 705
                    _ = v8
                case "a":
                    var v9 bool = (!(411 <= 2147483648))
                    _ = v9
                case "foo bar":
                    switch "" {
                    case "":
                        fmt.Println(v3)
                    case "a":
                        fmt.Println(v3, v3, fi6)
                    default:
                        fmt.Println(v3, fi6, v7)
                    }
                default:
                    var v10 int32 = 0
                    _ = v10
                }
            case "foo bar":
                _ = len(v1)
            case "hello":
                var v11 int = v3
                _ = v11
            default:
                v12 := GenS1{F0: -64, F1: 396, F2: 627}
                _ = v12
            }
            v3 = (fi6 << ((127) & 63))
            _ = fi6
            _ = v7
        }
    }
    v3 = 477
    if false {
        var v13 float32 = -2.25
        v3 -= v3
        for fi14 := 0; fi14 < 20; fi14++ {
            v3 += v3
            if len(v4) > 0 {
                v4[0] = -9223372036854775808
            }
            _ = fi14
        }
        _ = v13
    } else {
        v15 := GenS2{F0: 1, F1: v3, F2: v3}
        var v16 uint32 = 2
        v17 := GenS2{F0: 4294967295, F1: (-(v3)), F2: -1}
        _ = v15
        _ = v16
        _ = v17
    }
    v3 += (v3 * (v3 << ((833) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
