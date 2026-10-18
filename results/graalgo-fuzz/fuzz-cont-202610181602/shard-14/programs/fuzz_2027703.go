package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{32768, -879, -907, -1}
    var v3 int8 = -128
    var v4 float64 = float64(float32(0.0))
    v5 := []uint32{2, 65535}
    var v6 int64 = -9223372036854775808
    switch -333 {
    case 536:
        v2 = append(v2, 42)
    case -824:
        if false {
            v5 = append(v5, 0)
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v7 uint32 = 2
            _ = v7
        }
    default:
        var v8 int = -1
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: (1000 == -32768)}
    if true {
        v10 := GenS2{F0: (v6 ^ v6)}
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            _ = len(v5)
            _ = fi12
        }
        switch fi11 {
        case -9223372036854775808:
            fmt.Println(v3, v3)
        case 832:
            var v13 int64 = 42
            _ = v13
        }
        _ = fi11
    }
    if true {
        v6 = v6
    }
    _ = len(v1)
    if (255 < 32767) {
        if (!false) {
            for fi14 := 0; fi14 < 4; fi14++ {
                switch 2147483647 {
                case 9223372036854775807:
                    if true {
                        v15 := GenS1{F0: 2, F1: false}
                        v4 -= float64(1e10)
                        _ = v15
                    } else {
                        v5 = append(v5, 0)
                        v16 := GenS2{F0: (v6 ^ v6)}
                        v17 := v1[0]
                        _ = v16
                        _ = v17
                    }
                }
                _ = fi14
            }
        }
        if true {
            for fi18 := 0; fi18 < 5; fi18++ {
                if false {
                    _ = len(v5)
                    fmt.Println(v6, v4, v3)
                } else {
                    _ = len(v2)
                    v4 = (v4 / (float64(1e10) / ((1.5 - (-2.25 - v4)) + (v4 * v4))))
                    v19 := v1[639]
                    _ = v19
                }
                _ = fi18
            }
        } else {
            if (9223372036854775807 > 128) {
                var v20 uint32 = 2147483648
                _ = v20
            } else {
                fmt.Println(v3, v4)
                v21 := GenS2{F0: v6}
                _ = v21
            }
            if false {
                switch -327 {
                case -2147483648:
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v4 += float64(float32(float64(((float32(((float64(((0.0 * ((float32((v4 + 0.0)) / 3.14159) / 1.5)) * ((0.1 + 1e10) - float32((v4 * (v4 + v4)))))) * 0.1) - v4)) - 0.0) * 0.0))))
                        _ = fi22
                    }
                default:
                    switch 2147483647 {
                    case -1:
                        fmt.Println(v3, v6, v4)
                    case 953:
                        v23 := GenS2{F0: int64(v3)}
                        _ = v23
                    case 0:
                        for fi24 := 0; fi24 < 15; fi24++ {
                            var v25 uint64 = 131
                            _ = fi24
                            _ = v25
                        }
                    default:
                        var v26 int64 = v6
                        _ = v26
                    }
                }
                fmt.Println(v4, v6, v6)
                for fi27 := 0; fi27 < 20; fi27++ {
                    if ("foo bar" < ("\n\t" + ("hello" + (("a" + "foo bar") + ("foo bar" + ("a" + "hello")))))) {
                        var v28 uint64 = 18446744073709551615
                        fmt.Println(v28, v6)
                        _ = v28
                    } else {
                        var v29 int = (fi27 + (fi27 | -468))
                        _ = v29
                    }
                    _ = fi27
                }
            }
            for fi30 := 0; fi30 < 13; fi30++ {
                v31 := GenS1{F0: 42, F1: (v4 >= v4)}
                for fi32 := 0; fi32 < 9; fi32++ {
                    v6 += (^(v6))
                    v6 -= v6
                    _ = fi32
                }
                _ = fi30
                _ = v31
            }
        }
    } else {
        fmt.Println(v4)
        v4 = (0.1 * v4)
        v6 -= v6
    }
    var v33 string = "\n\t"
    fmt.Println(v4, v6)
    v34 := GenS2{F0: 903}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
