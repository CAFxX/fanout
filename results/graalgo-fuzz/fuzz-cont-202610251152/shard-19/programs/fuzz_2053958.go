package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint16; F2 bool }
type GenS5 struct { F0 int64; F1 int; F2 uint8 }
type GenS6 struct { F0 string; F1 uint32 }
func genFn0() uint64 {
    var v4 bool = true
    if v4 {
        v5 := GenS1{F0: 277, F1: 1}
        v6 := GenS1{F0: 42, F1: -88}
        _ = v5
        _ = v6
    } else {
        v4 = (v4 && v4)
        v7 := GenS2{F0: 128}
        v4 = v4
        _ = v7
    }
    _ = v4
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = (float64((0.1 * ((-2.25 - 3.14159) * (3.14159 / -2.25)))) - (((0.1 + 0.1) * 0.0) / -2.25))
    v3 := []uint32{}
    var v8 int = 0
    var v9 bool = (((("\n\t" + "") + ("\n\t" + "\n\t")) + ("a" + ("\n\t" + "a"))) == (("a" + "foo bar") + (("hello" + "foo bar") + ("\n\t" + "foo bar"))))
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 uint64 = genFn0()
        switch v8 {
        case 99:
            v8 = (fi10 / ((fi10) | 1))
        case 32768:
            var v12 int64 = 9223372036854775807
            _ = v12
        case -346:
            v13 := GenS3{F0: 255}
            _ = v13
        default:
            if (9223372036854775807 < -1) {
                for fi14 := 0; fi14 < 4; fi14++ {
                    v15 := GenS1{F0: 1, F1: 1}
                    switch 127 {
                    case -124:
                        var v16 uint = 0
                        _ = v16
                    case 32768:
                        fmt.Println(fi14, fi10, v9)
                    case -791:
                        v17 := GenS2{F0: 42}
                        _ = v17
                    default:
                        if v9 {
                            switch ("\n\t" + ("\n\t" + ((((("foo bar" + "\n\t") + "\n\t") + "hello") + "a") + ("foo bar" + "foo bar")))) {
                            case "hello":
                                var v18 int16 = 42
                                _ = v18
                            }
                            switch 4294967295 {
                            case 876:
                                switch 361 {
                                case 4294967295:
                                    if false {
                                        switch ("" + (("foo bar" + "foo bar") + ("a" + (("a" + "a") + "foo bar")))) {
                                        case "\n\t":
                                            v11 += (v11 - (v11 + 2))
                                        default:
                                            v3 = append(v3, 91)
                                        }
                                    }
                                default:
                                    var v19 int64 = 1
                                    _ = v19
                                }
                            default:
                                if (-614 >= -908) {
                                    v8 = (fi14 % ((32767) | 1))
                                } else {
                                    v20 := GenS2{F0: 0}
                                    _ = v20
                                }
                            }
                        } else {
                            var v21 int = (^(fi14))
                            var v22 float64 = ((((0.0 * (((v2 + ((1e10 - v2) * (v2 - -2.25))) - 1e10) * (1.5 - float64(0.0)))) / 1e10) * -2.25) / (1e10 * v2))
                            _ = v21
                            _ = v22
                        }
                    }
                    if (-2.25 < 0.0) {
                        v23 := GenS4{F0: 4294967295, F1: 32767, F2: v9}
                        if (4294967295 > 1) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if v9 {
                                if v9 {
                                    var v24 uint = 42
                                    v2 += ((0.0 + v2) - 0.1)
                                    v2 = float64((-2.25 + 0.1))
                                    if ((float32((v2 * (v2 / 0.1))) != 1.5) || (42 >= 0)) {
                                        v25 := GenS5{F0: -9223372036854775808, F1: -1, F2: 255}
                                        _ = v25
                                    }
                                    v3 = append(v3, uint32(v11))
                                    var v26 int16 = -76
                                    var v27 float32 = 3.14159
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v28 uint64 = genFn0()
                                    _ = v24
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 int64 = 42
                                    var v30 uint64 = (v11 | 42)
                                    v3 = append(v3, 1)
                                    _ = v29
                                    _ = v30
                                }
                                _ = len(v3)
                            }
                            var v31 int32 = -2147483648
                            _ = v31
                        }
                        _ = v23
                    }
                    _ = fi14
                    _ = v15
                }
                v9 = v9
                v32 := GenS1{F0: 129, F1: 32767}
                _ = v32
            }
        }
        v11 = (v11 ^ v11)
        _ = fi10
        _ = v11
    }
    var v33 uint = 42
    fmt.Println(v8)
    v34 := GenS6{F0: ("" + ""), F1: 424}
    for fi35 := 0; fi35 < 7; fi35++ {
        for fi36 := 0; fi36 < 8; fi36++ {
            var v37 bool = v9
            _ = fi36
            _ = v37
        }
        v33 = (v33 / (((^(v33))) | 1))
        _ = fi35
    }
    for fi38 := 0; fi38 < 3; fi38++ {
        v8 = (^(fi38))
        _ = fi38
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v39 uint8 = 128
    fmt.Println(v33, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v39
    _ = v8
    _ = v9
}
