package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool; F1 uint64; F2 int }
func genFn0(p0 uint32, p1 int64, p2 uint32) uint8 {
    var v5 int8 = 127
    v6 := GenS1{F0: "", F1: p2, F2: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int) uint8 {
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int = p0
            _ = fi8
            _ = v9
        }
        var v10 float32 = (((((-2.25 + 1.5) + 1.5) - (float32((3.14159 * -2.25)) * (1.5 / (1.5 - ((0.1 * 1.5) * -2.25))))) - 0.0) - 1.5)
        if true {
            if (0 == 1) {
                for fi11 := 0; fi11 < 4; fi11++ {
                    if ((!true) || ((1.5 / (((((v10 / v10) * ((3.14159 / v10) / ((float32(1.5) + v10) * v10))) - (3.14159 - v10)) / 0.1) / v10)) != float32(-2.25))) {
                        var v12 string = "\n\t"
                        _ = v12
                    } else {
                        v13 := []int{fi11, fi11, (-(p0))}
                        if len(v13) > 0 {
                            v13[0] = 4294967295
                        }
                        _ = v13
                    }
                    _ = fi11
                }
            } else {
                p0 -= 285
                for fi14 := 0; fi14 < 18; fi14++ {
                    if false {
                        fmt.Println(p0)
                        if true {
                            v15 := GenS1{F0: "hello", F1: 2147483648, F2: 4294967295}
                            if (1 <= 0) {
                                switch p0 {
                                case 128:
                                    var v16 uint16 = 0
                                    _ = v16
                                case 4294967295:
                                    if (!(719 == 394)) {
                                        var v17 float64 = 1e10
                                        v18 := GenS2{F0: true, F1: (("hello" + ("\n\t" + ("\n\t" + "hello"))) + ("hello" + ""))}
                                        p0 = (fi7 / (((^(fi7))) | 1))
                                        _ = v17
                                        _ = v18
                                    } else {
                                        var v19 int64 = 42
                                        _ = v19
                                    }
                                case 9223372036854775807:
                                    p0 -= 32767
                                default:
                                    p0 += fi14
                                }
                            } else {
                                v20 := GenS3{F0: (v10 > 3.14159), F1: 1}
                                p0 += fi14
                                _ = v20
                            }
                            fmt.Println(fi14)
                            p0 = (fi14 - 32768)
                            v10 += 1.5
                            _ = v15
                        } else {
                            fmt.Println(v10, fi7, fi14)
                        }
                    }
                    if false {
                        switch fi14 {
                        case -2147483648:
                            if false {
                                var v21 int16 = -525
                                var v22 float64 = (3.14159 + (-2.25 + 1e10))
                                _ = v21
                                _ = v22
                            }
                        case -9223372036854775808:
                            fmt.Println(p0, fi7)
                        }
                    }
                    v23 := []int{-571, p0, (fi7 | 127)}
                    _ = fi14
                    _ = v23
                }
                if (-32768 <= -431) {
                    var v24 uint16 = 65535
                    if true {
                        v25 := GenS3{F0: ((!true) && true), F1: (v24 & v24)}
                        p0 -= p0
                        for fi26 := 0; fi26 < 17; fi26++ {
                            var v27 bool = (!true)
                            v28 := []int{(fi7 ^ (fi7 * fi7)), (fi7 + (fi26 - fi7)), fi7}
                            fmt.Println(v24, p0)
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        _ = v25
                    } else {
                        var v29 uint16 = (v24 >> ((65535) & 63))
                        _ = v29
                    }
                    _ = v24
                } else {
                    if false {
                        v30 := []int{}
                        _ = v30
                    } else {
                        if (!(-39 <= 370)) {
                            v10 = v10
                        }
                        v10 += 1.5
                        if (1 > -494) {
                            v31 := GenS4{F0: (42 < 18446744073709551615), F1: 255, F2: 682}
                            v32 := GenS5{F0: ""}
                            _ = v31
                            _ = v32
                        }
                    }
                }
            }
        }
        _ = fi7
        _ = v10
    }
    v33 := GenS1{F0: ("" + "hello"), F1: 2147483647, F2: 2147483648}
    v34 := []int{(p0 << ((42) & 63)), (p0 / ((-1) | 1)), (p0 % ((128) | 1)), (p0 >> ((4294967295) & 63))}
    _ = p0
    _ = v33
    _ = v34
    return 2
}

func main() {
    v1 := []string{"", ("foo bar" + ("\n\t" + "hello")), "a", (("" + (((("foo bar" + "") + ("hello" + ("\n\t" + "\n\t"))) + (("a" + "foo bar") + (((("hello" + "a") + ("foo bar" + "foo bar")) + "a") + ""))) + ("" + "foo bar"))) + "")}
    v2 := make(map[string]string)
    var v3 uint8 = 127
    v4 := make(map[string]int)
    fmt.Println(v3)
    var v35 uint64 = 42
    fmt.Println(v3, v35)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = (^(v3))
    var v36 int64 = -773
    switch "hello" {
    case "\n\t":
        v36 = v36
    default:
        for fi37 := 0; fi37 < 7; fi37++ {
            v3 = v3
            _ = len(v2)
            if (v3 <= (v3 ^ 0)) {
                fmt.Println(fi37, v36)
                for fi38 := 0; fi38 < 5; fi38++ {
                    v39 := GenS6{F0: (v35 + 1)}
                    switch fi37 {
                    case 9223372036854775807:
                        var v40 uint32 = 2147483647
                        _ = v40
                    case 128:
                        v35 = v35
                    case 2147483647:
                        var v41 int = fi38
                        _ = v41
                    }
                    _ = fi38
                    _ = v39
                }
                fmt.Println(v35)
            }
            _ = fi37
        }
    }
    var v42 int64 = v36
    var v43 bool = (false || ((v42 + 9223372036854775807) == v36))
    fmt.Println(v36, v3, v43)
    if (2147483647 != 42) {
        switch 127 {
        case 127:
            var v44 int32 = -1
            _ = v44
        case 236:
            for fi45 := 0; fi45 < 19; fi45++ {
                var v46 int64 = (v36 - (v36 ^ (v36 / (((v42 / ((v36) | 1))) | 1))))
                v47 := GenS1{F0: "", F1: 235, F2: 2}
                _ = fi45
                _ = v46
                _ = v47
            }
        }
    } else {
        var v48 uint8 = 128
        _ = v48
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        var v50 string = "a"
        var v51 int64 = 9223372036854775807
        _ = fi49
        _ = v50
        _ = v51
    }
    v36 = v42
    fmt.Println(v42)
    if len(v1) > 0 {
        v1[0] = (("" + ("\n\t" + "\n\t")) + ("foo bar" + "\n\t"))
    }
    for fi52 := 0; fi52 < 19; fi52++ {
        switch fi52 {
        case -2147483648:
            v53 := GenS7{F0: v43, F1: v35, F2: (fi52 * 2147483647)}
            _ = v53
        case -1:
            v54 := v2["foo bar"]
            _ = v54
        default:
            var v55 uint64 = v35
            _ = v55
        }
        _ = fi52
    }
    var v56 bool = v43
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v56
}
