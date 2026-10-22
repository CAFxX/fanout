package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }
type GenS2 struct { F0 int32; F1 bool; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 uint16; F2 int16 }
func genFn0() int8 {
    var v7 float64 = 3.14159
    fmt.Println(v7)
    if (42 > 100) {
        v8 := []int{127, 32767, 32767, -2147483648}
        if (0 >= 1) {
            fmt.Println(v7)
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint32 = 1
                switch fi9 {
                case 1:
                    for fi11 := 0; fi11 < 17; fi11++ {
                        fmt.Println(fi11)
                        _ = fi11
                    }
                case 9223372036854775807:
                    if (!true) {
                        v10 = 2147483648
                    } else {
                        for fi12 := 0; fi12 < 14; fi12++ {
                            v10 -= 611
                            switch fi12 {
                            case -1:
                                var v13 int = -339
                                _ = v13
                            case -9223372036854775808:
                                if len(v8) > 0 {
                                    v8[0] = fi9
                                }
                            case 1:
                                switch "\n\t" {
                                case "a":
                                    v10 += 130
                                case "":
                                    v10 -= (^(v10))
                                }
                            }
                            _ = fi12
                        }
                        for fi14 := 0; fi14 < 1; fi14++ {
                            v8 = append(v8, fi14)
                            if (2147483647 < 1) {
                                if (true || (((0 > -1) || false) && true)) {
                                    if false {
                                        v10 += (v10 - (^(v10)))
                                    } else {
                                        if (v7 == (1.5 + (v7 / v7))) {
                                            var v15 uint8 = 2
                                            _ = v15
                                        }
                                    }
                                    if (-1 < 9223372036854775807) {
                                        if true {
                                            v16 := GenS1{F0: 0, F1: -9223372036854775808, F2: ("a" + ("\n\t" + ""))}
                                            _ = v16
                                        } else {
                                            v8 = append(v8, fi14)
                                            var v17 int32 = -656
                                            var v18 uint = 2
                                            _ = v17
                                            _ = v18
                                        }
                                    } else {
                                        var v19 int64 = 642
                                        switch fi9 {
                                        case -739:
                                            v19 = 42
                                        }
                                        _ = v19
                                    }
                                    if false {
                                        v7 = v7
                                        var v20 int8 = 1
                                        if ((-(fi9)) <= fi9) {
                                            v20 += (^(v20))
                                        }
                                        _ = v20
                                    } else {
                                        var v21 uint8 = 42
                                        _ = v21
                                    }
                                }
                                fmt.Println(v10, fi9)
                            }
                            _ = fi14
                        }
                    }
                case 128:
                    for fi22 := 0; fi22 < 20; fi22++ {
                        v10 = 576
                        if (!true) {
                            _ = len(v8)
                            v10 = v10
                            fmt.Println(fi9)
                        }
                        if (!false) {
                            var v23 uint = uint(v10)
                            v10 -= 0
                            _ = len(v8)
                            _ = v23
                        }
                        _ = fi22
                    }
                default:
                    v24 := GenS2{F0: 0, F1: true, F2: 403}
                    _ = v24
                }
                _ = fi9
                _ = v10
            }
            v7 = v7
        } else {
            if (!true) {
                if true {
                    switch "" {
                    case "\n\t":
                        v8 = append(v8, 617)
                    }
                    var v25 int8 = -102
                    v7 = v7
                    _ = v25
                }
            }
        }
        _ = v8
    } else {
        var v26 uint16 = 41
        v27 := GenS3{F0: 1000}
        v28 := GenS4{F0: 18446744073709551615}
        _ = v26
        _ = v27
        _ = v28
    }
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 106
    v3 := make(map[string]bool)
    var v4 int64 = 9223372036854775807
    v5 := []int{-9223372036854775808, -2147483648}
    var v6 int16 = 286
    v2 -= v2
    v2 += (v2 & 181)
    v6 += 32767
    switch 32767 {
    case 562:
        if false {
            v29 := v1[("a" + "foo bar")]
            _ = v29
        } else {
            fmt.Println(v2, v2)
            v30 := v1["a"]
            v31 := GenS2{F0: 160, F1: true, F2: 42}
            _ = v30
            _ = v31
        }
    case -546:
        delete(v3, "")
    default:
        v6 = (^(v6))
    }
    var v32 uint64 = 2
    v4 = v4
    v33 := GenS3{F0: v6}
    if false {
        if true {
            v34 := GenS1{F0: 65535, F1: -98, F2: "foo bar"}
            var v35 int64 = v4
            for fi36 := 0; fi36 < 9; fi36++ {
                v32 = v32
                v37 := GenS3{F0: (v6 ^ v6)}
                if len(v5) > 0 {
                    v5[0] = (fi36 ^ (fi36 | 32767))
                }
                _ = fi36
                _ = v37
            }
            _ = v34
            _ = v35
        }
        var v38 int16 = v6
        v39 := GenS5{F0: "foo bar", F1: uint16(v32), F2: (v38 * v38)}
        _ = v38
        _ = v39
    } else {
        v40 := GenS4{F0: (v2 / ((uint(v32)) | 1))}
        _ = v40
    }
    var v41 uint8 = 22
    var v42 uint64 = (^(v32))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
