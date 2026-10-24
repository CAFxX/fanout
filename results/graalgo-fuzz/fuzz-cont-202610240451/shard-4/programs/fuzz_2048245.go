package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint8 = 246
    var v2 float64 = 1.5
    v3 := []uint32{191, 363, 827, 204}
    var v4 float64 = (v2 + (float64(0.0) + v2))
    var v5 float32 = 3.14159
    switch -2147483648 {
    case -9223372036854775808:
        switch "foo bar" {
        case "a":
            fmt.Println(v1, v2, v5)
        case "hello":
            v5 = (v5 + 0.0)
        case "foo bar":
            for fi6 := 0; fi6 < 3; fi6++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if (!(2147483647 == 2147483647)) {
                    v2 += (float64(((v5 * (v5 * (3.14159 - 0.1))) / 1e10)) * float64((v5 / v5)))
                } else {
                    for fi7 := 0; fi7 < 13; fi7++ {
                        fmt.Println(fi6, v2, v1)
                        _ = fi7
                    }
                    if (-9223372036854775808 > 538) {
                        var v8 bool = false
                        var v9 uint16 = 715
                        _ = v8
                        _ = v9
                    }
                }
                if false {
                    var v10 uint32 = 1
                    v10 += (v10 & 65535)
                    _ = v10
                } else {
                    fmt.Println(v4)
                    v11 := GenS1{F0: -894, F1: 18446744073709551615}
                    fmt.Println(v2, v2, v5)
                    _ = v11
                }
                _ = fi6
            }
        default:
            for fi12 := 0; fi12 < 5; fi12++ {
                v5 = (v5 - (v5 * 1.5))
                v5 -= v5
                var v13 float32 = (((v5 + float32((v4 + v2))) + v5) - (0.1 + float32(0.1)))
                _ = fi12
                _ = v13
            }
        }
    default:
        v14 := GenS2{F0: 2147483647, F1: ((1 >= 1000) && (!("hello" < "foo bar")))}
        _ = v14
    }
    v3 = append(v3, 0)
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 uint64 = 68
        var v17 uint8 = (^(v1))
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 float32 = v5
        var v20 int = -427
        var v21 uint32 = 2147483648
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    var v22 bool = (42 >= 502)
    v1 = v1
    if (-9223372036854775808 <= 42) {
        var v23 int = -17
        var v24 uint32 = 42
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
