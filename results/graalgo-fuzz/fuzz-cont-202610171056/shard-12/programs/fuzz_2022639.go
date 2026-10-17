package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint }
type GenS2 struct { F0 uint16; F1 uint32; F2 int32 }
func genFn0() bool {
    v7 := []int{-156}
    var v8 int16 = 42
    v8 = v8
    v9 := GenS1{F0: 97, F1: int64(v8), F2: 1}
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 bool = true
    v2 := []int{0}
    var v3 int16 = 255
    var v4 uint = 1
    var v5 uint8 = 154
    var v6 int64 = 607
    switch 32767 {
    case 505:
        var v10 float64 = 1.5
        _ = v10
    case 128:
        v11 := GenS2{F0: 255, F1: 584, F2: 42}
        _ = v11
    default:
        var v12 int64 = (^(v6))
        _ = v12
    }
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    _ = len(v2)
    if true {
        v6 -= v6
    }
    v4 -= (v4 ^ (v4 >> (((v4 ^ (v4 / ((v4) | 1)))) & 63)))
    var v13 bool = false
    for fi14 := 0; fi14 < 3; fi14++ {
        v13 = v13
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 int8 = 0
            if (("" >= "foo bar") || (v4 != (^(v4)))) {
                var v17 bool = v13
                fmt.Println(fi15, v13, v1)
                if true {
                    if v1 {
                        var v18 int8 = (v16 + -96)
                        _ = v18
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v19 int = 127
                    _ = v19
                } else {
                    if v1 {
                        var v20 int8 = 127
                        v21 := GenS1{F0: 127, F1: v6, F2: v4}
                        switch fi14 {
                        case -1:
                            var v22 string = ("\n\t" + "hello")
                            _ = v22
                        }
                        if (2147483647 > -1) {
                            v5 += v5
                            var v23 int16 = v3
                            var v24 int8 = v16
                            _ = v23
                            _ = v24
                        } else {
                            var v25 float32 = 1e10
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v25
                        }
                        v20 += (v16 / ((v16) | 1))
                        fmt.Println(v17)
                        _ = v20
                        _ = v21
                    } else {
                        var v26 float32 = float32(((-2.25 / (3.14159 / -2.25)) / 1e10))
                        v27 := GenS1{F0: (v5 >> ((v4) & 63)), F1: v6, F2: (v4 % ((v4) | 1))}
                        _ = v26
                        _ = v27
                    }
                }
                _ = v17
            }
            _ = fi15
            _ = v16
        }
        var v28 uint = v4
        _ = fi14
        _ = v28
    }
    v2 = append(v2, -2147483648)
    v29 := GenS1{F0: v5, F1: v6, F2: v4}
    switch "a" {
    case "foo bar":
        fmt.Println(v1, v5, v3)
    case "\n\t":
        var v30 string = ""
        _ = v30
    default:
        v5 -= 127
    }
    v1 = (0 < 127)
    v13 = (-1 != -852)
    for fi31 := 0; fi31 < 3; fi31++ {
        v2 = append(v2, 127)
        var v32 uint8 = v5
        _ = fi31
        _ = v32
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
