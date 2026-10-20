package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 string; F1 bool }

func main() {
    var v1 string = (("\n\t" + "a") + "hello")
    v2 := make(map[uint32]uint32)
    v3 := []string{"\n\t", (v1 + "\n\t"), "foo bar", v1}
    var v4 float64 = (0.0 * 1.5)
    var v5 uint64 = 42
    v6 := []int{128, 32767}
    if (!false) {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 uint8 = 215
            _ = len(v3)
            v9 := GenS1{F0: 42, F1: 170}
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(v5, v1, v5)
    } else {
        if len(v6) > 0 {
            v6[0] = -691
        }
    }
    if true {
        v5 += (^(v5))
    }
    var v10 float32 = (0.1 - 1e10)
    v11 := GenS1{F0: 1000, F1: 4294967295}
    v5 -= v5
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(fi12, v10)
        if true {
            var v13 uint32 = 65535
            v14 := v2[v13]
            if true {
                if ((-450 <= 2147483647) && false) {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        v16 := GenS1{F0: -459, F1: 0}
                        if false {
                            v17 := GenS1{F0: 491, F1: 2}
                            _ = v17
                        } else {
                            v1 = v1
                            var v18 uint8 = 255
                            _ = v18
                        }
                        switch fi12 {
                        case -9223372036854775808:
                            switch v1 {
                            case "foo bar":
                                if (!(783 > 0)) {
                                    fmt.Println(v14, v5, v5)
                                    v1 = ""
                                    var v19 uint32 = 0
                                    _ = v19
                                }
                            default:
                                switch (fi15 >> ((2) & 63)) {
                                case 737:
                                    var v20 float32 = 3.14159
                                    _ = v20
                                case 4294967295:
                                    var v21 int = (fi15 ^ fi15)
                                    _ = v21
                                default:
                                    v2[(v14 << ((255) & 63))] = (v14 << ((166) & 63))
                                }
                            }
                        default:
                            switch v1 {
                            case "\n\t":
                                if ((v10 - float32((v4 + 1e10))) == ((1e10 + v10) + v10)) {
                                    v2[v13] = (v13 / (((v13 - (v14 - v14))) | 1))
                                    v10 += v10
                                } else {
                                    if len(v6) > 0 {
                                        v6[0] = fi12
                                    }
                                }
                                if (42 > -1) {
                                    var v22 int32 = 468
                                    _ = v22
                                } else {
                                    v2[v13] = 65535
                                }
                                v23 := GenS1{F0: 32767, F1: 255}
                                _ = v23
                            default:
                                if false {
                                    var v24 float32 = 0.0
                                    switch 614 {
                                    case 32768:
                                        fmt.Println(v13)
                                    case 127:
                                        if true {
                                            v5 -= v5
                                            v10 -= v24
                                        }
                                    default:
                                        v5 -= (v5 % ((v5) | 1))
                                    }
                                    _ = v24
                                } else {
                                    var v25 int16 = -32768
                                    switch (fi12 + (fi12 * 127)) {
                                    case 1:
                                        var v26 int = 42
                                        _ = v26
                                    }
                                    v27 := GenS2{F0: ((((v1 + "hello") + "\n\t") + "foo bar") + v1), F1: ((v10 + (((v10 - v10) / (0.0 + (v10 + -2.25))) / (0.1 - ((v10 / v10) - v10)))) == ((((v10 + (v10 * float32(v4))) - 0.1) - (v10 + (v10 / (float32(v4) * v10)))) * v10))}
                                    _ = v25
                                    _ = v27
                                }
                            }
                        }
                        var v28 int32 = 0
                        _ = fi15
                        _ = v16
                        _ = v28
                    }
                }
            }
            _ = v13
            _ = v14
        }
        _ = fi12
    }
    var v29 float32 = (float32((v4 * (1.5 / v4))) + ((v10 - -2.25) - 3.14159))
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 int16 = 42
        _ = fi30
        _ = v31
    }
    switch ("\n\t" + "") {
    case "hello":
        v3 = append(v3, v1)
    default:
        var v32 uint32 = 2147483647
        _ = v32
    }
    var v33 int32 = 167
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
