package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64) uint16 {
    if true {
        v6 := GenS1{F0: 2147483648, F1: 42}
        _ = v6
    } else {
        fmt.Println(p0)
    }
    var v7 int64 = p0
    _ = p0
    _ = v7
    return 0
}

func main() {
    v1 := []int{4294967295, 867, 2147483647, 9223372036854775807}
    var v2 uint64 = 42
    var v3 float64 = float64((float32(float64(1e10)) + (0.1 + 1e10)))
    var v4 uint16 = 65535
    var v5 int16 = 1000
    var v8 int64 = 9223372036854775807
    v2 += (v2 ^ 0)
    v9 := GenS1{F0: 2147483647, F1: 42}
    v10 := GenS1{F0: 492, F1: 42}
    var v11 uint8 = 156
    switch "a" {
    case "\n\t":
        var v12 int16 = 542
        _ = v12
    case "a":
        var v13 uint64 = v2
        _ = v13
    }
    switch 42 {
    case 32767:
        switch "" {
        case "\n\t":
            switch -1 {
            case 127:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                v4 += v4
            }
        case "hello":
            if (v11 >= (^(v11))) {
                v14 := GenS2{F0: (v5 - v5)}
                for fi15 := 0; fi15 < 1; fi15++ {
                    var v16 uint64 = 4294967295
                    for fi17 := 0; fi17 < 6; fi17++ {
                        switch 127 {
                        case -1:
                            v2 += v2
                            var v18 uint16 = v4
                            var v19 int = (fi15 / ((fi17) | 1))
                            _ = v18
                            _ = v19
                        default:
                            switch (-(fi15)) {
                            case -707:
                                var v20 int = fi15
                                _ = v20
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = (fi15 ^ 128)
                        }
                        var v21 int32 = 484
                        var v22 float32 = (0.0 + (0.0 - 1e10))
                        fmt.Println(v5, v11)
                        _ = fi17
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS2{F0: int16(fi15)}
                    _ = fi15
                    _ = v16
                    _ = v23
                }
                if true {
                    v24 := GenS2{F0: (v5 / (((v5 * (-(v5)))) | 1))}
                    if false {
                        var v25 int = -2147483648
                        _ = v25
                    } else {
                        v4 += (v4 * 163)
                    }
                    _ = v24
                }
                _ = v14
            } else {
                v8 = (v8 << ((18446744073709551615) & 63))
            }
        case "a":
            switch "" {
            case "hello":
                _ = len(v1)
            }
        default:
            v26 := GenS1{F0: 42, F1: -2147483648}
            _ = v26
        }
    }
    fmt.Println(v11)
    if len(v1) > 0 {
        v1[0] = 0
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
