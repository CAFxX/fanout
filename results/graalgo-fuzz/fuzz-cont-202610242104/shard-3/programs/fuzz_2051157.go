package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint; F2 uint16 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint8 = 255
    v2 := []int{-2147483648, 9223372036854775807}
    var v3 uint64 = 596
    var v4 int64 = 9223372036854775807
    var v5 string = "foo bar"
    _ = len(v2)
    if true {
        for fi6 := 0; fi6 < 8; fi6++ {
            v3 = (v3 * v3)
            var v7 int = (fi6 / (((fi6 ^ (fi6 - fi6))) | 1))
            fmt.Println(v1)
            _ = fi6
            _ = v7
        }
        var v8 float32 = (1e10 + 0.0)
        v9 := GenS1{F0: -1}
        _ = v8
        _ = v9
    }
    var v10 int = 128
    v4 = (v4 * v4)
    var v11 bool = (false || (!true))
    var v12 int8 = 100
    v13 := GenS2{F0: (v4 / ((v4) | 1)), F1: 0, F2: 697}
    if (!true) {
        v2 = append(v2, 127)
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS2{F0: (v4 << ((0) & 63)), F1: 0, F2: 1}
        for fi16 := 0; fi16 < 14; fi16++ {
            if (65535 == 65535) {
                fmt.Println(v5, v5)
                v2 = append(v2, (v10 / ((0) | 1)))
                if (v11 || v11) {
                    var v17 uint16 = uint16(v3)
                    _ = v17
                }
            } else {
                fmt.Println(v5, v5)
                var v18 int32 = 1
                switch fi14 {
                case 128:
                    var v19 int32 = (v18 << ((2) & 63))
                    _ = v19
                case 979:
                    var v20 int32 = v18
                    _ = v20
                case 32767:
                    v21 := GenS3{F0: (v4 ^ (v4 & (-(v4))))}
                    _ = v21
                default:
                    var v22 uint = 225
                    _ = v22
                }
                _ = v18
            }
            if len(v2) > 0 {
                v2[0] = fi16
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    v10 += 2147483647
    var v23 int64 = v4
    if len(v2) > 0 {
        v2[0] = v10
    }
    v2 = append(v2, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
