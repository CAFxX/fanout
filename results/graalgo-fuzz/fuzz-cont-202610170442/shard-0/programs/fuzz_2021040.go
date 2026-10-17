package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int8; F1 int64; F2 uint32 }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]bool)
    var v3 int16 = 363
    v4 := []uint32{}
    v5 := GenS1{F0: 4294967295, F1: (v3 | v3), F2: 2}
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint16 = 255
        var v8 uint16 = (v7 % ((v7) | 1))
        var v9 int = (fi6 ^ fi6)
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 int = 128
    for fi11 := 0; fi11 < 2; fi11++ {
        v10 -= (fi11 * (^(v10)))
        _ = fi11
    }
    var v12 int = (v10 ^ (^(v10)))
    v10 = 970
    var v13 uint = 870
    if (((-(v10)) != v10) || true) {
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint16 = 2
            _ = fi14
            _ = v15
        }
        v10 += (v10 / ((2147483647) | 1))
    }
    v2[76] = ((v1 == v1) && (42 == 37))
    var v16 float64 = 0.0
    if false {
        v17 := v2[(v12 + v12)]
        _ = v17
    }
    if ((v13 / (((v13 % ((v13) | 1))) | 1)) == v13) {
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 uint = v13
            v20 := GenS1{F0: 65535, F1: 32767, F2: uint32(v19)}
            _ = fi18
            _ = v19
            _ = v20
        }
        fmt.Println(v3, v13)
    } else {
        var v21 int8 = 1
        v22 := GenS2{F0: true, F1: ((uint64(v13) > 0) || false)}
        _ = v21
        _ = v22
    }
    v12 += -1
    if false {
        for fi23 := 0; fi23 < 2; fi23++ {
            v24 := GenS3{F0: -1, F1: -9223372036854775808, F2: 2147483647}
            if (-2.25 > 1.5) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v25 := GenS1{F0: 42, F1: v3, F2: 0}
                _ = v25
            }
            _ = fi23
            _ = v24
        }
        v10 -= (v12 & (^(v10)))
        var v26 int32 = -2147483648
        _ = v26
    }
    if true {
        v10 += v10
    } else {
        for fi27 := 0; fi27 < 1; fi27++ {
            v28 := GenS1{F0: 2147483647, F1: v3, F2: 2}
            if (1e10 < 0.0) {
                for fi29 := 0; fi29 < 4; fi29++ {
                    v3 = v3
                    var v30 int = 128
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                    _ = fi29
                    _ = v30
                }
            }
            _ = fi27
            _ = v28
        }
        v3 = v3
        if true {
            if true {
                v31 := GenS1{F0: 2147483648, F1: 694, F2: 65535}
                switch "foo bar" {
                case "foo bar":
                    v16 = 0.1
                default:
                    v32 := GenS2{F0: ((0 >= -2147483648) && false), F1: true}
                    _ = v32
                }
                var v33 int16 = (^(v3))
                _ = v31
                _ = v33
            }
            fmt.Println(v3)
            var v34 int8 = -128
            _ = v34
        }
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
