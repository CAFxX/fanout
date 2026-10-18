package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int16; F2 int32 }

func main() {
    var v1 uint16 = 32767
    v2 := []int{-828}
    var v3 float32 = -2.25
    var v4 float32 = (1e10 * 3.14159)
    v5 := make(map[int]uint32)
    var v6 int16 = 42
    _ = len(v2)
    v7 := GenS1{F0: 451}
    v6 = (v6 & v6)
    var v8 int = 32767
    var v9 uint8 = 2
    switch "" {
    case "foo bar":
        switch (((("a" + "\n\t") + ("foo bar" + (("\n\t" + "foo bar") + "foo bar"))) + "hello") + "\n\t") {
        case "a":
            switch v8 {
            case -9223372036854775808:
                v8 = v8
            case 1:
                v8 = v8
            }
        default:
            switch v8 {
            case 1:
                var v10 uint64 = 812
                _ = v10
            case 0:
                switch v8 {
                case -831:
                    var v11 uint32 = 2147483648
                    _ = v11
                case 9223372036854775807:
                    switch 128 {
                    case 128:
                        v1 = 65535
                    case 42:
                        fmt.Println(v8, v9)
                    case 806:
                        for fi12 := 0; fi12 < 6; fi12++ {
                            v6 -= v6
                            _ = fi12
                        }
                    default:
                        v13 := GenS2{F0: 18446744073709551615, F1: v6, F2: -2147483648}
                        _ = v13
                    }
                case -17:
                    v4 = v4
                default:
                    v14 := GenS1{F0: 809}
                    _ = v14
                }
            case 32768:
                var v15 int8 = 42
                _ = v15
            default:
                v8 = (v8 - v8)
            }
        }
    default:
        var v16 uint8 = (v9 + 128)
        _ = v16
    }
    var v17 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
