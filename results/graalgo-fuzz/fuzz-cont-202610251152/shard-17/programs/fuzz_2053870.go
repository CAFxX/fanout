package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int16 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0() uint16 {
    var v6 float32 = -2.25
    v7 := []int{32768, 0}
    var v8 int16 = 0
    fmt.Println(v8, v6)
    var v9 uint8 = 127
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 float32 = (((-2.25 - float32(-2.25)) - float32((0.0 * (3.14159 / 1.0)))) * 0.1)
    var v2 float32 = (((-2.25 * (-2.25 / (v1 * float32((1e10 + (3.14159 - 3.14159)))))) - 0.1) * (((0.1 / (0.0 + v1)) * v1) - ((3.14159 / 1e10) + (v1 - v1))))
    v3 := make(map[uint32]bool)
    var v4 float64 = 0.0
    var v5 int8 = -1
    v10 := []int{-306}
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(v2)
        for fi12 := 0; fi12 < 20; fi12++ {
            if ((v5 < v5) && true) {
                v1 -= v2
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            v13 := GenS1{F0: 297, F1: (!true), F2: 923}
            _ = fi12
            _ = v13
        }
        var v14 uint = 724
        _ = fi11
        _ = v14
    }
    v15 := GenS1{F0: 2147483647, F1: (1 > 119), F2: -1}
    if false {
        v1 = (v2 * ((v1 - v2) - v1))
        var v16 float64 = -2.25
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v2 -= v2
        v5 -= v5
        _ = fi17
    }
    var v18 uint64 = 546
    switch 161 {
    case 0:
        fmt.Println(v5, v4, v18)
    case -699:
        fmt.Println(v4, v2)
    }
    v5 = 22
    var v19 int = 128
    if (true && false) {
        switch ("\n\t" + "hello") {
        case "hello":
            _ = len(v10)
        case "a":
            v20 := GenS2{F0: (v19 + v19), F1: v18}
            _ = v20
        default:
            if (true && (!true)) {
                v21 := GenS2{F0: 32768, F1: (v18 & (^(v18)))}
                _ = v21
            } else {
                v2 = float32(-2.25)
                switch "a" {
                case "a":
                    var v22 float64 = float64(0.0)
                    _ = v22
                case "\n\t":
                    fmt.Println(v1, v4)
                case "hello":
                    switch v19 {
                    case -754:
                        v5 = 100
                    case -9223372036854775808:
                        v3[2147483648] = false
                    case 260:
                        var v23 uint64 = (v18 << ((255) & 63))
                        _ = v23
                    default:
                        var v24 int64 = 42
                        _ = v24
                    }
                }
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
        }
        v5 -= (v5 & v5)
    } else {
        v5 = (v5 % (((v5 << ((662) & 63))) | 1))
        for fi25 := 0; fi25 < 1; fi25++ {
            if false {
                var v26 float32 = v2
                v3[2] = true
                for fi27 := 0; fi27 < 12; fi27++ {
                    v3[785] = false
                    if false {
                        var v28 int8 = (v5 << ((4294967295) & 63))
                        _ = v28
                    } else {
                        switch "\n\t" {
                        case "hello":
                            v29 := GenS2{F0: v19, F1: 729}
                            _ = v29
                        case "foo bar":
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        case "a":
                            var v30 string = "foo bar"
                            _ = v30
                        }
                    }
                    _ = fi27
                }
                _ = v26
            } else {
                var v31 uint32 = 28
                v2 -= 1.5
                _ = v31
            }
            _ = fi25
        }
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        v33 := GenS1{F0: 1, F1: true, F2: 1000}
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
