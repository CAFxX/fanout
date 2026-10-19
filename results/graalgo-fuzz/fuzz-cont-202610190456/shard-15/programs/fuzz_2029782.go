package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint32 }

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]int)
    var v3 string = "a"
    v4 := []int{32768}
    var v5 string = v3
    var v6 float64 = float64((float32(0.0) / (((3.14159 / 0.1) * (((-2.25 - float32(0.1)) - 3.14159) * (1.5 * float32(((3.14159 - 1.5) - (float64((3.14159 - 1.5)) - (0.1 - -2.25))))))) - 1e10)))
    switch 862 {
    case -753:
        if false {
            var v7 uint16 = 748
            for fi8 := 0; fi8 < 3; fi8++ {
                if ((fi8 >> ((2) & 63)) > fi8) {
                    if len(v4) > 0 {
                        v4[0] = -1
                    }
                } else {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        v5 = v5
                        _ = fi9
                    }
                    var v10 uint16 = (v7 % ((v7) | 1))
                    _ = v10
                }
                _ = fi8
            }
            var v11 string = ("a" + v5)
            _ = v11
            _ = v7
        } else {
            if true {
                if ((!false) || false) {
                    fmt.Println(v3, v5)
                } else {
                    _ = len(v4)
                    v12 := GenS1{F0: v5}
                    _ = v12
                }
                delete(v2, 42)
                for fi13 := 0; fi13 < 11; fi13++ {
                    fmt.Println(v1)
                    fmt.Println(v5, v3)
                    _ = fi13
                }
            }
        }
    default:
        v14 := GenS2{F0: 1, F1: 63}
        _ = v14
    }
    fmt.Println(v6, v1, v1)
    v5 = v3
    if (65535 != 1) {
        v15 := GenS2{F0: -9223372036854775808, F1: 2}
        _ = v15
    }
    v5 = v5
    v6 = -2.25
    v5 = v5
    var v16 int16 = -1
    if (true || (-2.25 > ((float64(0.1) * (v6 / v6)) * v6))) {
        v6 += (1e10 * ((((3.14159 / v6) / (v6 + v6)) - 0.1) * v6))
    }
    if (127 >= 42) {
        var v17 uint = 0
        _ = v17
    } else {
        v1 = (v1 & v1)
        for fi18 := 0; fi18 < 12; fi18++ {
            v3 = v5
            _ = fi18
        }
        var v19 int8 = 42
        _ = v19
    }
    switch 42 {
    case 42:
        var v20 uint32 = 169
        _ = v20
    case -384:
        if ((v3 + "") == v5) {
            v21 := GenS2{F0: 9223372036854775807, F1: 0}
            _ = len(v2)
            v6 -= (1e10 - float64(0.1))
            _ = v21
        }
    }
    v22 := v2[65535]
    if (0 == 0) {
        for fi23 := 0; fi23 < 2; fi23++ {
            v4 = append(v4, (fi23 | v22))
            var v24 uint64 = 4294967295
            _ = fi23
            _ = v24
        }
    } else {
        v22 = v22
        if true {
            _ = len(v4)
            for fi25 := 0; fi25 < 20; fi25++ {
                for fi26 := 0; fi26 < 2; fi26++ {
                    var v27 int8 = -1
                    _ = len(v2)
                    _ = fi26
                    _ = v27
                }
                var v28 int = (fi25 / (((v22 + (^(fi25)))) | 1))
                _ = fi25
                _ = v28
            }
        }
        var v29 int64 = -1
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
