package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint64; F2 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 0
    v3 := make(map[string]bool)
    v4 := []int{-2147483648, 815}
    if false {
        v1[""] = 2147483648
    } else {
        var v5 uint64 = 4294967295
        _ = v5
    }
    var v6 string = "foo bar"
    v6 = ((v6 + (("foo bar" + "a") + v6)) + v6)
    if ((65535 < v2) && false) {
        v2 -= (v2 % ((65535) | 1))
        var v7 float64 = ((3.14159 + (-2.25 + (0.0 - (3.14159 / 1.5)))) * -2.25)
        switch -154 {
        case 127:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case -688:
            fmt.Println(v2, v6, v6)
        default:
            if (-877 < -2147483648) {
                v6 = "\n\t"
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2 += 642
            }
        }
        _ = v7
    } else {
        delete(v1, v6)
        var v8 float64 = ((1e10 / 3.14159) - 0.1)
        fmt.Println(v2, v6, v8)
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(v6)
            var v11 bool = (!false)
            if false {
                var v12 int = fi9
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        v13 := GenS1{F0: -128}
        _ = fi9
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v6 = "foo bar"
        v15 := GenS2{F0: 0, F1: 255}
        if true {
            if len(v4) > 0 {
                v4[0] = (^(fi14))
            }
            if (!false) {
                v16 := GenS1{F0: -1}
                v17 := GenS3{F0: 32767, F1: 1, F2: -1}
                _ = len(v4)
                _ = v16
                _ = v17
            } else {
                switch 4294967295 {
                case -805:
                    var v18 float32 = (1.5 + 0.1)
                    _ = v18
                case 32767:
                    if true {
                        var v19 uint8 = 67
                        fmt.Println(v19, v6)
                        _ = v19
                    }
                case -319:
                    v6 = v6
                default:
                    for fi20 := 0; fi20 < 14; fi20++ {
                        switch (fi14 ^ (fi20 >> ((255) & 63))) {
                        case 0:
                            v6 = "a"
                        default:
                            v2 += (v2 & v2)
                        }
                        switch 2147483647 {
                        case -2147483648:
                            fmt.Println(v6, v6)
                        default:
                            _ = len(v4)
                        }
                        _ = fi20
                    }
                }
                fmt.Println(fi14, v6, fi14)
                var v21 uint = 4294967295
                _ = v21
            }
            v22 := GenS4{F0: 19}
            _ = v22
        }
        _ = fi14
        _ = v15
    }
    var v23 uint8 = 42
    v2 = v2
    v4 = append(v4, 32767)
    v24 := GenS5{F0: 0, F1: 433}
    v6 = v6
    v23 = (v23 | v23)
    var v25 int32 = -993
    v26 := GenS2{F0: 32, F1: v23}
    v23 -= v23
    v23 += v23
    v27 := GenS6{F0: v6}
    for fi28 := 0; fi28 < 19; fi28++ {
        v29 := v1[v6]
        v30 := GenS2{F0: -617, F1: (v23 % ((255) | 1))}
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v6
}
