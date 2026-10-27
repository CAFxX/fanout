package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint64; F1 string; F2 int16 }
type GenS5 struct { F0 bool; F1 bool; F2 uint32 }
func genFn0() uint8 {
    v4 := []int{4294967295}
    var v5 int64 = -1
    v5 += v5
    _ = v4
    _ = v5
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 2147483647
    v3 := make(map[uint32]bool)
    v2 = v2
    v6 := GenS1{F0: -283, F1: (0 == 628), F2: 417}
    v7 := GenS2{F0: 1000}
    var v8 int = -2147483648
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int64 = -9223372036854775808
        delete(v1, 2)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        if true {
            switch "\n\t" {
            case "foo bar":
                switch (fi11 ^ v8) {
                case 4294967295:
                    v2 -= (v8 - (v2 % ((v2) | 1)))
                default:
                    v12 := GenS2{F0: 1}
                    _ = v12
                }
            default:
                for fi13 := 0; fi13 < 16; fi13++ {
                    switch (^(fi13)) {
                    case -877:
                        v14 := GenS1{F0: 313, F1: (255 != 1), F2: 1}
                        _ = v14
                    case 0:
                        if true {
                            v15 := GenS1{F0: -1, F1: true, F2: 785}
                            if (55 >= -451) {
                                v2 = (v8 << ((65535) & 63))
                            } else {
                                var v16 float32 = ((1e10 + float32((float64(((3.14159 * (float32(-2.25) * 1e10)) - 0.1)) + 3.14159))) * (((0.1 + float32(float64(float32(3.14159)))) / (1e10 / -2.25)) * (1e10 - 1.5)))
                                if (180 <= 18446744073709551615) {
                                    switch "hello" {
                                    case "\n\t":
                                        v8 -= (fi13 & (v8 ^ (v2 >> ((498) & 63))))
                                    default:
                                        _ = len(v3)
                                    }
                                } else {
                                    v17 := GenS3{F0: 2147483647, F1: 255, F2: -1}
                                    var v18 int16 = -17
                                    _ = v17
                                    _ = v18
                                }
                                v8 += 2147483647
                                v8 -= (fi13 & (v2 % (((^(fi13))) | 1)))
                                _ = v16
                            }
                            var v19 uint = 753
                            _ = v15
                            _ = v19
                        } else {
                            var v20 float64 = 1e10
                            _ = v20
                        }
                    default:
                        v2 = (fi11 >> ((255) & 63))
                    }
                    v8 = (v8 | v8)
                    if ((false && true) || (1 > 2)) {
                        fmt.Println(v2, v2, v2)
                        var v21 float32 = 0.1
                        _ = v21
                    } else {
                        v22 := GenS3{F0: 2, F1: 95, F2: 0}
                        v2 = -2147483648
                        _ = v22
                    }
                    _ = fi13
                }
            }
            for fi23 := 0; fi23 < 10; fi23++ {
                if true {
                    v8 += (fi11 ^ v8)
                    if true {
                        v2 = fi23
                        var v24 int8 = -128
                        var v25 uint32 = 819
                        var v26 uint8 = 127
                        _ = v24
                        _ = v25
                        _ = v26
                    } else {
                        var v27 float64 = 0.1
                        _ = v27
                    }
                    switch (v8 * v2) {
                    case 32767:
                        var v28 int64 = 0
                        _ = v28
                    case 32768:
                        var v29 int8 = -1
                        _ = v29
                    case -127:
                        if (-1 > 958) {
                            v30 := GenS4{F0: 0, F1: "foo bar", F2: 0}
                            _ = v30
                        }
                    default:
                        var v31 bool = false
                        _ = v31
                    }
                    var v32 bool = true
                    _ = v32
                }
                var v33 uint8 = genFn0()
                _ = fi23
                _ = v33
            }
        }
        var v34 int8 = -1
        _ = fi11
        _ = v34
    }
    v35 := GenS5{F0: (!false), F1: true, F2: 2147483647}
    for fi36 := 0; fi36 < 19; fi36++ {
        v8 = (fi36 * (fi36 | (^(fi36))))
        _ = fi36
    }
    var v37 uint16 = 333
    v37 = v37
    for fi38 := 0; fi38 < 19; fi38++ {
        v3[0] = true
        _ = fi38
    }
    v1[577] = v2
    var v39 float64 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v39
    _ = v6
    _ = v7
    _ = v8
}
