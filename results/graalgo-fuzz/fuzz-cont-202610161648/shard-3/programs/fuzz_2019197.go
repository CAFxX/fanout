package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint }

func main() {
    var v1 uint32 = 65535
    var v2 float32 = ((0.0 - -2.25) / 0.1)
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := []int{2147483647, 1, (fi5 >> ((727) & 63)), (fi5 * fi5)}
        var v7 string = "\n\t"
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint = 0
    if (true && (1 <= 127)) {
        for fi9 := 0; fi9 < 20; fi9++ {
            v2 = float32(0.1)
            _ = fi9
        }
    }
    v10 := GenS1{F0: 497, F1: 65535, F2: 0}
    v8 = 0
    for fi11 := 0; fi11 < 18; fi11++ {
        v2 -= (v2 - 0.0)
        v8 += 2
        _ = fi11
    }
    v12 := GenS1{F0: 105, F1: 293, F2: 128}
    if (1 > 1) {
        for fi13 := 0; fi13 < 12; fi13++ {
            for fi14 := 0; fi14 < 10; fi14++ {
                fmt.Println(fi13)
                switch ("" + "a") {
                case "foo bar":
                    if true {
                        var v15 string = "foo bar"
                        v2 += v2
                        switch 285 {
                        case 1:
                            v1 += (v1 & v1)
                        case 0:
                            switch v15 {
                            case "\n\t":
                                var v16 int32 = 42
                                _ = v16
                            case "":
                                if (!(v15 <= v15)) {
                                    var v17 int8 = 127
                                    _ = v17
                                } else {
                                    v1 += (v1 >> ((v8) & 63))
                                    v15 = "foo bar"
                                }
                            case "a":
                                if false {
                                    var v18 uint64 = 2
                                    fmt.Println(fi13, v15)
                                    _ = v18
                                }
                            default:
                                if (-1 == 712) {
                                    v2 -= v2
                                    var v19 int = fi13
                                    _ = v19
                                } else {
                                    var v20 int64 = 0
                                    v21 := GenS1{F0: 0, F1: 646, F2: uint8(v1)}
                                    v8 += (v8 << (((v8 - v8)) & 63))
                                    _ = v20
                                    _ = v21
                                }
                            }
                        case 32767:
                            v22 := GenS1{F0: 603, F1: 2, F2: 1}
                            _ = v22
                        }
                        var v23 float64 = ((((-2.25 + -2.25) / float64(v2)) - 1.5) + (3.14159 + 1e10))
                        _ = v15
                        _ = v23
                    } else {
                        var v24 float32 = 0.0
                        if (18446744073709551615 <= 2) {
                            v25 := GenS1{F0: 18446744073709551615, F1: 65535, F2: 218}
                            v26 := []int{fi13}
                            _ = v25
                            _ = v26
                        }
                        _ = v24
                    }
                default:
                    fmt.Println(v2, v1, v1)
                }
                v27 := GenS2{F0: 42, F1: v8}
                _ = fi14
                _ = v27
            }
            _ = fi13
        }
        fmt.Println(v1, v2, v2)
    } else {
        for fi28 := 0; fi28 < 6; fi28++ {
            v1 = 629
            v1 = v1
            _ = fi28
        }
        v1 += v1
        v3[682] = 2
    }
    v29 := []int{32767, 436}
    var v30 string = "\n\t"
    v31 := GenS2{F0: -128, F1: uint(v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v8
}
