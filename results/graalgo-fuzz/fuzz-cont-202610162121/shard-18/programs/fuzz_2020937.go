package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }
func genFn0() uint8 {
    var v5 string = "foo bar"
    var v6 int8 = 100
    _ = v5
    _ = v6
    return 128
}

func main() {
    var v1 int = 128
    var v2 uint32 = 405
    v3 := make(map[uint32]string)
    v4 := []string{"a", "foo bar"}
    v7 := GenS1{F0: 0}
    if ((v1 ^ 299) != (^(v1))) {
        fmt.Println(v2, v1)
    } else {
        if (1 > 18446744073709551615) {
            v2 = 0
        }
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 float32 = 1.5
        for fi10 := 0; fi10 < 17; fi10++ {
            if true {
                if false {
                    if (!false) {
                        v11 := GenS1{F0: -920}
                        if (1 > 989) {
                            if false {
                                fmt.Println(v1, v9)
                            }
                        } else {
                            _ = len(v3)
                        }
                        var v12 int8 = -1
                        _ = v11
                        _ = v12
                    } else {
                        fmt.Println(v1, v2, v2)
                        var v13 uint32 = 2147483648
                        var v14 int64 = -1
                        _ = v13
                        _ = v14
                    }
                    var v15 int16 = 1000
                    v1 = 32767
                    _ = v15
                } else {
                    if (1.5 > (((1e10 - (0.1 - 0.0)) - (((0.0 + 1.5) - 1e10) / (3.14159 * (1e10 / 1e10)))) - (1e10 * 1.5))) {
                        var v16 uint32 = (v2 | v2)
                        _ = v16
                    }
                    var v17 int32 = 967
                    if true {
                        v18 := GenS2{F0: 2, F1: -1, F2: uint8(v2)}
                        v19 := v3[(v2 & v2)]
                        v2 = v2
                        _ = v18
                        _ = v19
                    }
                    _ = v17
                }
                var v20 float64 = 3.14159
                v1 = 892
                _ = v20
            }
            fmt.Println(fi10, v9, v2)
            v3[v2] = "foo bar"
            _ = fi10
        }
        v1 -= -993
        _ = fi8
        _ = v9
    }
    var v21 int = 32767
    delete(v3, (v2 ^ 889))
    v1 -= (v21 << ((1) & 63))
    fmt.Println(v21, v1)
    var v22 uint = 0
    for fi23 := 0; fi23 < 9; fi23++ {
        if (((v1 >= 15) || false) && false) {
            var v24 int32 = -375
            v25 := v3[v2]
            var v26 int = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        v1 += -9223372036854775808
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v7
}
