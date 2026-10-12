package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 string; F1 int; F2 int8 }
func genFn0(p0 uint64) uint64 {
    var v6 string = "a"
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint16 = 1
        p0 = (p0 | 42)
        p0 += (^(p0))
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 646, F1: 0}
    if (!(!true)) {
        v10 := []int{-698, 0, 1}
        var v11 float64 = 3.14159
        if false {
            for fi12 := 0; fi12 < 12; fi12++ {
                v10 = append(v10, fi12)
                _ = fi12
            }
            if ((!((0.0 * 1.5) > 0.1)) && true) {
                if (((97 < -128) || (!(1 >= 100))) && (0.1 < ((0.0 * ((((0.1 / (-2.25 * (((1e10 + (-2.25 * -2.25)) - 3.14159) + (float32(1.5) + 3.14159)))) + 0.0) - float32((3.14159 * 3.14159))) / 3.14159)) + float32(0.1)))) {
                    if true {
                        if true {
                            var v13 uint32 = 4294967295
                            for fi14 := 0; fi14 < 1; fi14++ {
                                var v15 int16 = 1000
                                var v16 int64 = 1
                                fmt.Println(v11)
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            _ = v13
                        }
                    } else {
                        if true {
                            var v17 float32 = ((float32((v11 - ((1e10 + float64(-2.25)) + 3.14159))) - 1e10) - 0.0)
                            _ = len(v10)
                            _ = v17
                        }
                        switch 4294967295 {
                        case 1:
                            v6 = "hello"
                        case 4294967295:
                            _ = len(v10)
                        default:
                            fmt.Println(p0)
                        }
                    }
                    var v18 float64 = 0.0
                    if true {
                        v18 -= v18
                        var v19 int64 = 9223372036854775807
                        _ = v19
                    }
                    _ = v18
                }
            }
        } else {
            fmt.Println(v6)
            _ = len(v10)
        }
        _ = v10
        _ = v11
    } else {
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 string = v6
            switch "foo bar" {
            case "foo bar":
                v22 := GenS2{F0: (((("foo bar" + "foo bar") + "foo bar") + v6) + ((v6 + v21) + ("a" + v6))), F1: -893, F2: 127}
                _ = v22
            case "\n\t":
                var v23 uint64 = p0
                _ = v23
            default:
                fmt.Println(fi20)
            }
            p0 -= (p0 % ((p0) | 1))
            _ = fi20
            _ = v21
        }
    }
    _ = p0
    _ = v6
    _ = v9
    return 801
}

func main() {
    var v1 int32 = 0
    var v2 int8 = 1
    var v3 int64 = 9223372036854775807
    var v4 int16 = 0
    var v5 float64 = 3.14159
    var v24 uint8 = 0
    var v25 float32 = (0.0 * 0.1)
    v26 := GenS2{F0: (("" + "hello") + (("\n\t" + "hello") + ("" + "hello"))), F1: -180, F2: (v2 / (((-(v2))) | 1))}
    for fi27 := 0; fi27 < 3; fi27++ {
        var v28 string = "a"
        _ = fi27
        _ = v28
    }
    v29 := []int{628, 32768, 2147483647, 127}
    v4 = (v4 ^ v4)
    if len(v29) > 0 {
        v29[0] = -359
    }
    v3 = v3
    var v30 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
