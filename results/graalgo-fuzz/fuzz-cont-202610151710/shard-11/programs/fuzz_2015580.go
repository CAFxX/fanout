package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint }
type GenS2 struct { F0 bool; F1 int64 }
func genFn0(p0 uint8, p1 int16, p2 int) bool {
    var v5 int64 = 0
    fmt.Println(v5, v5, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return true
}
func genFn1(p0 int64, p1 bool, p2 string) uint8 {
    for fi6 := 0; fi6 < 19; fi6++ {
        if ((-841 > -1) && p1) {
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 int64 = -9223372036854775808
                _ = fi7
                _ = v8
            }
        }
        var v9 int64 = p0
        p2 = (p2 + ("\n\t" + p2))
        _ = fi6
        _ = v9
    }
    var v10 bool = (!false)
    v11 := GenS1{F0: -72, F1: p1, F2: 877}
    for fi12 := 0; fi12 < 8; fi12++ {
        p1 = p1
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 2
}

func main() {
    var v1 uint = 4294967295
    var v2 bool = false
    var v3 int8 = 0
    v4 := make(map[uint32]string)
    v13 := []int{-9223372036854775808, 32767, 1, 0}
    v1 -= (^(v1))
    var v14 int64 = 9223372036854775807
    switch ("a" + "a") {
    case "hello":
        v2 = false
    case "a":
        switch "foo bar" {
        case "":
            v15 := GenS1{F0: -648, F1: (2147483647 < 1), F2: (v1 % ((v1) | 1))}
            _ = v15
        case "foo bar":
            var v16 uint16 = 519
            _ = v16
        case "a":
            switch 134 {
            case 32767:
                if ((((((float64((0.0 - 1.5)) - 0.0) * (3.14159 / 1.0)) / (1e10 + (0.0 / ((3.14159 - float64(0.0)) * (1e10 * float64((1.5 + -2.25))))))) < -2.25) || v2) || (!v2)) {
                    var v17 int64 = -9223372036854775808
                    switch 919 {
                    case 0:
                        switch ("" + ("\n\t" + ("\n\t" + ("foo bar" + "hello")))) {
                        case "":
                            switch (("a" + "a") + "foo bar") {
                            case "\n\t":
                                v18 := GenS1{F0: -487, F1: (705 > 96), F2: (v1 * (v1 + 42))}
                                _ = v18
                            case "hello":
                                var v19 int = -596
                                _ = v19
                            default:
                                for fi20 := 0; fi20 < 11; fi20++ {
                                    fmt.Println(v14, fi20)
                                    if (int(v14) <= fi20) {
                                        v1 -= v1
                                    } else {
                                        var v21 bool = (true && true)
                                        _ = v21
                                    }
                                    _ = fi20
                                }
                            }
                        case "hello":
                            fmt.Println(v1, v2, v17)
                        }
                    case 32768:
                        for fi22 := 0; fi22 < 1; fi22++ {
                            fmt.Println(v14, v1, v17)
                            _ = fi22
                        }
                    case 848:
                        switch "a" {
                        case "\n\t":
                            var v23 uint64 = 354
                            _ = v23
                        case "hello":
                            v17 += (v14 * v14)
                        }
                    default:
                        var v24 uint8 = 0
                        _ = v24
                    }
                    v14 -= v14
                    _ = v17
                } else {
                    v2 = false
                    fmt.Println(v1, v3)
                }
            case 28:
                v1 += (v1 * v1)
            default:
                switch "hello" {
                case "foo bar":
                    fmt.Println(v2, v2)
                default:
                    v3 -= -128
                }
            }
        }
    }
    v3 += v3
    var v25 bool = (v2 && v2)
    v2 = (("hello" + "hello") < ("" + ("\n\t" + "a")))
    if (!v25) {
        if (!v2) {
            var v26 uint8 = 0
            v27 := GenS1{F0: -9223372036854775808, F1: (v1 != (^(v1))), F2: (^(v1))}
            _ = v26
            _ = v27
        }
        var v28 int = 2147483647
        v29 := GenS2{F0: ((-2147483648 == 2147483647) && (1 < 0)), F1: v14}
        _ = v28
        _ = v29
    }
    var v30 uint64 = 1
    v3 = (v3 * (v3 | v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
}
