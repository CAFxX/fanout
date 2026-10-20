package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 bool }
type GenS2 struct { F0 uint16; F1 bool; F2 bool }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]bool)
    var v4 int64 = -172
    v5 := make(map[uint32]uint32)
    v4 += v4
    for fi6 := 0; fi6 < 6; fi6++ {
        if (1e10 < float64(0.0)) {
            var v7 uint32 = 65535
            _ = v7
        } else {
            if (!(32 > 42)) {
                for fi8 := 0; fi8 < 4; fi8++ {
                    v9 := GenS1{F0: 964, F1: v4, F2: false}
                    v2 += v2
                    _ = fi8
                    _ = v9
                }
                v10 := GenS1{F0: 18446744073709551615, F1: (v4 >> ((18446744073709551615) & 63)), F2: true}
                v2 -= (^(v2))
                _ = v10
            }
        }
        v2 = (v2 * v2)
        _ = fi6
    }
    v11 := []int{}
    var v12 uint = 1
    var v13 int64 = v4
    var v14 int64 = v13
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 bool = true
        _ = fi15
        _ = v16
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v17 := GenS2{F0: 411, F1: false, F2: false}
    v4 = v14
    if (((1.5 + ((((1.5 + (((1e10 / 0.1) + 0.1) + ((3.14159 / 1.0) / 1.5))) / -2.25) - (1e10 * -2.25)) * 3.14159)) / float64(0.1)) == (-2.25 / float64((1e10 + 1e10)))) {
        v12 = (^(v12))
    }
    if false {
        if (!false) {
            fmt.Println(v2, v13, v2)
            if false {
                switch 0 {
                case 32768:
                    fmt.Println(v14)
                case 630:
                    for fi18 := 0; fi18 < 5; fi18++ {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            var v20 uint64 = 1
                            v13 = (v4 % ((-1) | 1))
                            _ = fi19
                            _ = v20
                        }
                        var v21 bool = true
                        _ = fi18
                        _ = v21
                    }
                case 943:
                    v12 = v12
                default:
                    v22 := GenS2{F0: 65535, F1: (1.5 >= (0.1 - 3.14159)), F2: true}
                    _ = v22
                }
            } else {
                var v23 uint8 = 255
                var v24 int32 = 2147483647
                _ = v23
                _ = v24
            }
            v25 := GenS3{F0: uint(v2)}
            _ = v25
        }
        v26 := GenS3{F0: (v12 - (v12 & (^(v12))))}
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
