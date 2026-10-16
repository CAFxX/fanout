package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 string }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint }
func genFn0() uint {
    var v4 uint64 = 2
    var v5 string = "a"
    fmt.Println(v4, v5)
    switch 9223372036854775807 {
    case -597:
        var v6 string = ((v5 + v5) + v5)
        _ = v6
    case 150:
        var v7 string = v5
        _ = v7
    case 540:
        var v8 uint64 = (^(v4))
        _ = v8
    }
    _ = v4
    _ = v5
    return 739
}
func genFn1(p0 uint64) int64 {
    switch "" {
    case "a":
        var v9 uint8 = 128
        _ = v9
    case "foo bar":
        p0 = p0
    case "hello":
        if true {
            if true {
                if (!((-32768 >= 206) && true)) {
                    p0 -= (p0 << ((4294967295) & 63))
                    var v10 int64 = 9223372036854775807
                    v11 := GenS1{F0: (v10 - 0), F1: 1000, F2: "hello"}
                    _ = v10
                    _ = v11
                }
                switch 0 {
                case -1:
                    if (true && (-1 < 2147483647)) {
                        p0 = p0
                        v12 := GenS1{F0: 9223372036854775807, F1: 16, F2: "foo bar"}
                        _ = v12
                    }
                case 32767:
                    if (0.1 == (float32(3.14159) * (0.0 + 0.1))) {
                        for fi13 := 0; fi13 < 2; fi13++ {
                            if true {
                                switch "" {
                                case "":
                                    v14 := []int{}
                                    _ = v14
                                default:
                                    v15 := []int{(fi13 & fi13), fi13}
                                    _ = v15
                                }
                                p0 = 4294967295
                                for fi16 := 0; fi16 < 17; fi16++ {
                                    fmt.Println(fi13, p0)
                                    p0 -= 32
                                    _ = fi16
                                }
                            } else {
                                v17 := []int{fi13, 42}
                                var v18 float64 = float64(((0.0 / (1e10 - -2.25)) - float32(1.5)))
                                v18 = (1e10 / (3.14159 + v18))
                                _ = v17
                                _ = v18
                            }
                            v19 := []int{}
                            _ = fi13
                            _ = v19
                        }
                    } else {
                        var v20 int16 = 1
                        if true {
                            v21 := []int{127, 127}
                            p0 = (p0 % ((1) | 1))
                            _ = v21
                        } else {
                            v22 := GenS2{F0: "", F1: 1}
                            fmt.Println(v20)
                            _ = v22
                        }
                        _ = v20
                    }
                default:
                    v23 := GenS2{F0: ("foo bar" + "a"), F1: 127}
                    _ = v23
                }
            } else {
                switch ("hello" + ("" + "foo bar")) {
                case "a":
                    if false {
                        for fi24 := 0; fi24 < 2; fi24++ {
                            var v25 uint8 = 150
                            _ = fi24
                            _ = v25
                        }
                    }
                case "\n\t":
                    p0 += (p0 - p0)
                case "hello":
                    v26 := []int{32767}
                    _ = v26
                }
            }
            if (-828 >= -1) {
                if false {
                    v27 := GenS3{F0: (p0 >> ((genFn0()) & 63)), F1: false}
                    p0 -= p0
                    var v28 int8 = 1
                    _ = v27
                    _ = v28
                }
                p0 = 42
            }
        }
    }
    p0 = (p0 / ((2) | 1))
    _ = p0
    return 9223372036854775807
}

func main() {
    var v1 int16 = -32768
    var v2 uint8 = 128
    var v3 int8 = 100
    var v29 int = 42
    for fi30 := 0; fi30 < 11; fi30++ {
        v31 := []int{9223372036854775807, (fi30 + fi30)}
        fmt.Println(v29)
        _ = fi30
        _ = v31
    }
    v32 := []int{}
    v29 = v29
    var v33 uint = 2
    if (760 < uint64(v2)) {
        v34 := GenS1{F0: 0, F1: (-(v1)), F2: "hello"}
        for fi35 := 0; fi35 < 7; fi35++ {
            _ = len(v32)
            v2 = (v2 * v2)
            v2 -= v2
            _ = fi35
        }
        var v36 uint32 = 2147483647
        _ = v34
        _ = v36
    } else {
        if true {
            for fi37 := 0; fi37 < 11; fi37++ {
                v1 = -1
                var v38 int16 = (v1 - (v1 - v1))
                fmt.Println(v2, v33)
                _ = fi37
                _ = v38
            }
            var v39 int64 = genFn1(18446744073709551615)
            _ = v39
        }
        var v40 uint = 0
        switch (((("hello" + "") + "a") + "a") + "a") {
        case "\n\t":
            v29 += (v29 | (-(v29)))
        }
        _ = v40
    }
    v1 = 32767
    var v41 int32 = -1
    switch ("a" + "\n\t") {
    case "foo bar":
        if len(v32) > 0 {
            v32[0] = v29
        }
    case "\n\t":
        fmt.Println(v33)
    case "hello":
        v41 -= -1
    }
    v2 = (^(v2))
    var v42 float32 = 0.0
    v43 := GenS4{F0: (v1 << ((v33) & 63))}
    if true {
        v32 = append(v32, v29)
        fmt.Println(v3, v3)
    } else {
        v44 := GenS4{F0: (v1 - v1)}
        for fi45 := 0; fi45 < 11; fi45++ {
            v46 := GenS5{F0: "hello"}
            var v47 float64 = 3.14159
            fmt.Println(v2, v2, v47)
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = v44
    }
    if (0 < (v3 | -96)) {
        v41 += (-(v41))
        v48 := GenS2{F0: "\n\t", F1: v3}
        for fi49 := 0; fi49 < 17; fi49++ {
            v2 -= 22
            var v50 int = -473
            v51 := GenS5{F0: ("" + "foo bar")}
            _ = fi49
            _ = v50
            _ = v51
        }
        _ = v48
    } else {
        v52 := GenS6{F0: 220}
        for fi53 := 0; fi53 < 14; fi53++ {
            v29 = (^(v29))
            _ = fi53
        }
        v2 = (^(v2))
        _ = v52
    }
    for fi54 := 0; fi54 < 6; fi54++ {
        var v55 int32 = int32(v33)
        _ = fi54
        _ = v55
    }
    if false {
        v56 := GenS7{F0: (v33 ^ 18446744073709551615)}
        _ = v56
    } else {
        switch -9223372036854775808 {
        case -2147483648:
            var v57 int64 = 0
            _ = v57
        case -781:
            if len(v32) > 0 {
                v32[0] = 4294967295
            }
        default:
            fmt.Println(v41)
        }
        v29 += v29
        for fi58 := 0; fi58 < 20; fi58++ {
            var v59 uint = (v33 * (v33 << (((v33 ^ v33)) & 63)))
            v41 = (v41 | 1)
            _ = fi58
            _ = v59
        }
    }
    var v60 int64 = -1
    var v61 uint8 = (v2 - v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v41
    _ = v42
    _ = v43
    _ = v60
    _ = v61
}
