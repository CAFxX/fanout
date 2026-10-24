package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int }
type GenS2 struct { F0 uint64; F1 int8; F2 int16 }
func genFn0() string {
    var v4 bool = (-1 != -229)
    v5 := GenS1{F0: 916, F1: 4294967295, F2: 32768}
    v4 = v4
    fmt.Println(v4)
    _ = v4
    _ = v5
    return "foo bar"
}

func main() {
    var v1 float32 = 0.1
    var v2 float32 = 0.0
    v3 := make(map[string]uint32)
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            v2 += 0.0
            _ = fi7
        }
        _ = fi6
    }
    if (2 >= 274) {
        v2 -= (v2 - 0.0)
    } else {
        if (((true || false) && true) && (-1 >= 9223372036854775807)) {
            var v8 int8 = 42
            v9 := GenS2{F0: 385, F1: v8, F2: 787}
            _ = v8
            _ = v9
        } else {
            var v10 float64 = ((0.0 / 1.5) - 1e10)
            _ = v10
        }
    }
    var v11 uint64 = 600
    var v12 bool = false
    for fi13 := 0; fi13 < 16; fi13++ {
        v2 += v2
        if v12 {
            if (!v12) {
                switch (fi13 / (((fi13 << ((18446744073709551615) & 63))) | 1)) {
                case 1:
                    var v14 int64 = -6
                    _ = v14
                case -2147483648:
                    for fi15 := 0; fi15 < 12; fi15++ {
                        var v16 uint = 0
                        v12 = true
                        var v17 int16 = -266
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case -847:
                    v12 = false
                default:
                    var v18 uint64 = (v11 % ((338) | 1))
                    _ = v18
                }
                if ((!true) && v12) {
                    v19 := GenS2{F0: v11, F1: -128, F2: -448}
                    v1 = v1
                    v11 += (v11 * (^(v11)))
                    _ = v19
                } else {
                    v20 := []int{}
                    switch ("" + "\n\t") {
                    case "hello":
                        switch "hello" {
                        case "\n\t":
                            delete(v3, ((("a" + ("foo bar" + ((("hello" + "hello") + "foo bar") + "a"))) + "foo bar") + "hello"))
                        case "hello":
                            v2 -= (((float32(1.5) / v2) - float32(0.0)) * v2)
                        case "":
                            switch "a" {
                            case "":
                                var v21 uint = 18446744073709551615
                                _ = v21
                            case "a":
                                v2 = (((1.5 + 1e10) / 3.14159) - float32(-2.25))
                            default:
                                fmt.Println(v2, v12, v1)
                            }
                        }
                    }
                    _ = v20
                }
            }
            v22 := GenS1{F0: 42, F1: 42, F2: (fi13 ^ fi13)}
            switch fi13 {
            case 204:
                var v23 int64 = 9223372036854775807
                _ = v23
            case 2147483647:
                var v24 int8 = -52
                _ = v24
            case -1:
                v12 = false
            default:
                v2 = -2.25
            }
            _ = v22
        } else {
            for fi25 := 0; fi25 < 11; fi25++ {
                _ = len(v3)
                var v26 float64 = 0.1
                _ = fi25
                _ = v26
            }
            var v27 float32 = -2.25
            var v28 int64 = -1
            _ = v27
            _ = v28
        }
        _ = fi13
    }
    fmt.Println(v1, v1)
    v29 := []int{-9223372036854775808, 376}
    var v30 uint16 = 359
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
