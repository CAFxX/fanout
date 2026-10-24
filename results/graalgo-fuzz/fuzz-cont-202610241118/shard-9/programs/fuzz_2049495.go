package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint) int64 {
    v6 := []int{}
    if (true && false) {
        v6 = append(v6, 42)
        v7 := GenS1{F0: 127, F1: 200}
        p0 = (p0 / (((p0 + (p0 << ((p0) & 63)))) | 1))
        _ = v7
    } else {
        if ((((("" + "\n\t") <= "foo bar") && ((1e10 - 3.14159) > -2.25)) && false) && true) {
            p0 -= (p0 / ((696) | 1))
            switch "hello" {
            case "hello":
                v8 := GenS1{F0: 42, F1: 2}
                _ = v8
            case "foo bar":
                if len(v6) > 0 {
                    v6[0] = 880
                }
            default:
                var v9 uint = p0
                _ = v9
            }
        } else {
            switch "a" {
            case "foo bar":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case "\n\t":
                if ((^(p0)) == p0) {
                    if true {
                        v10 := GenS1{F0: 2, F1: 127}
                        _ = v10
                    }
                } else {
                    var v11 bool = (44 <= 347)
                    _ = v11
                }
            case "hello":
                v12 := GenS1{F0: 128, F1: 144}
                _ = v12
            }
            var v13 uint = p0
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS1{F0: 243, F1: 42}
            var v16 int32 = 42
            switch "\n\t" {
            case "foo bar":
                switch "\n\t" {
                case "":
                    var v17 int64 = 775
                    _ = v17
                case "hello":
                    for fi18 := 0; fi18 < 8; fi18++ {
                        v19 := GenS1{F0: 255, F1: 127}
                        switch (fi14 ^ fi18) {
                        case 929:
                            var v20 int8 = -128
                            _ = v20
                        }
                        v6 = append(v6, (fi18 & fi14))
                        _ = fi18
                        _ = v19
                    }
                case "\n\t":
                    fmt.Println(v16)
                default:
                    for fi21 := 0; fi21 < 18; fi21++ {
                        v6 = append(v6, (fi14 / ((0) | 1)))
                        _ = fi21
                    }
                }
            case "":
                if (0 <= 163) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            default:
                if false {
                    var v22 int = fi14
                    _ = v22
                }
            }
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    _ = p0
    _ = v6
    return 510
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    v3 := []uint32{810, 2147483647}
    v4 := make(map[string]string)
    v5 := make(map[uint32]uint32)
    var v23 float64 = ((float64(((0.1 - 0.1) * (3.14159 * (0.1 + ((1.5 + float32(float64(1.5))) * 0.0))))) + -2.25) + 3.14159)
    v23 += (3.14159 / v23)
    v23 -= ((v23 * float64(((-2.25 - 3.14159) + ((0.1 * 1e10) + (-2.25 / 0.1))))) * v23)
    if len(v3) > 0 {
        v3[0] = 2
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v23 = v23
    if (43 != 1) {
        switch 1 {
        case 4294967295:
            if false {
                v23 = float64(1.5)
                v24 := GenS1{F0: 255, F1: 255}
                var v25 uint32 = 2147483648
                _ = v24
                _ = v25
            } else {
                v23 += v23
            }
        default:
            if false {
                if (!((719 > 42) && true)) {
                    fmt.Println(v23)
                    var v26 int8 = 100
                    _ = v26
                }
            } else {
                var v27 int8 = -97
                v28 := GenS2{F0: 0}
                _ = v27
                _ = v28
            }
        }
        _ = len(v3)
        var v29 int32 = -2147483648
        _ = v29
    } else {
        var v30 uint8 = 0
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v30
    }
    fmt.Println(v23)
    var v31 float32 = float32((3.14159 * v23))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
