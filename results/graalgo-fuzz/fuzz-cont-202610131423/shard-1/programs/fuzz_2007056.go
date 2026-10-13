package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 string }
type GenS2 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS9 struct { F0 uint8; F1 int; F2 uint64 }
type GenS10 struct { F0 uint16; F1 uint16; F2 uint16 }
func genFn0(p0 uint8, p1 uint64, p2 uint) uint16 {
    p2 = (^(p2))
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := []int{}
        if false {
            v8 := GenS1{F0: p2, F1: (fi6 - 2147483647), F2: "hello"}
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := GenS2{F0: (p0 >> ((0) & 63)), F1: 29, F2: 78}
                p2 = (^(p2))
                _ = fi9
                _ = v10
            }
            _ = v8
        } else {
            switch "" {
            case "\n\t":
                if len(v7) > 0 {
                    v7[0] = 42
                }
            case "foo bar":
                var v11 int16 = 564
                _ = v11
            case "a":
                if len(v7) > 0 {
                    v7[0] = (^(fi6))
                }
            default:
                p2 = 1
            }
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            switch "" {
            case "a":
                if (false && false) {
                    p2 = p2
                    if true {
                        var v13 int8 = 100
                        fmt.Println(p1, p1)
                        p0 = p0
                        _ = v13
                    }
                    p2 = (^(p2))
                }
            }
            if true {
                if len(v7) > 0 {
                    v7[0] = (fi12 * (fi6 | fi6))
                }
                var v14 int32 = 2147483647
                _ = v14
            }
            fmt.Println(fi12, p0, fi12)
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            if (p1 < (p1 << ((42) & 63))) {
                if (!false) {
                    p1 = (^(p1))
                    if (65535 < 32767) {
                        fmt.Println(fi6, fi6)
                        if true {
                            if false {
                                p0 += p0
                                var v15 bool = false
                                if (!(0 != 317)) {
                                    var v16 uint16 = 255
                                    var v17 uint8 = (^(p0))
                                    v17 -= (v17 ^ (v17 & (p0 & (v17 ^ v17))))
                                    _ = v16
                                    _ = v17
                                } else {
                                    if false {
                                        var v18 int8 = 1
                                        _ = v18
                                    } else {
                                        p2 -= (p2 & 52)
                                        switch 128 {
                                        case -2147483648:
                                            var v19 int32 = 2147483647
                                            _ = v19
                                        case 32767:
                                            p0 -= 25
                                        default:
                                            _ = len(v7)
                                        }
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                    }
                                    if false {
                                        if len(v7) > 0 {
                                            v7[0] = fi12
                                        }
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        if v15 {
                                            switch ("" + "hello") {
                                            case "hello":
                                                if false {
                                                    switch "hello" {
                                                    case "hello":
                                                        p1 -= (^(p1))
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                    case "\n\t":
                                                        var v20 int32 = -1
                                                        _ = v20
                                                    case "":
                                                        v15 = (787 <= -9223372036854775808)
                                                    }
                                                    if v15 {
                                                        if len(v7) > 0 {
                                                            v7[0] = fi6
                                                        }
                                                        v21 := GenS2{F0: 116, F1: -1, F2: 416}
                                                        switch 32768 {
                                                        case 767:
                                                            p2 -= p2
                                                        default:
                                                            p1 -= p1
                                                        }
                                                        _ = v21
                                                    } else {
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                    }
                                                    v22 := GenS3{F0: 1, F1: 1, F2: 127}
                                                    _ = v22
                                                } else {
                                                    var v23 uint64 = (^(p1))
                                                    fmt.Println(p2)
                                                    _ = v23
                                                }
                                            }
                                        } else {
                                            v24 := GenS4{F0: -1}
                                            var v25 uint16 = 0
                                            var v26 uint8 = (p0 - 176)
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                }
                                _ = v15
                            } else {
                                v27 := GenS5{F0: 100}
                                p1 += p1
                                _ = v27
                            }
                        } else {
                            _ = len(v7)
                            var v28 bool = false
                            _ = v28
                        }
                        if (((3.14159 + float64(1e10)) * 1e10) == 0.1) {
                            var v29 float64 = (1e10 + ((((float64(-2.25) - 3.14159) + float64(-2.25)) + (((float64((0.0 * 1e10)) - (float64(0.0) - (0.0 - -2.25))) * 1e10) / 1.0)) + ((3.14159 * (1.5 / float64((1e10 / 1.5)))) / 1e10)))
                            _ = v29
                        } else {
                            _ = len(v7)
                            v30 := GenS6{F0: 207, F1: int(p0)}
                            _ = v30
                        }
                        var v31 uint8 = (p0 ^ p0)
                        _ = v31
                    }
                    p1 = p1
                }
                var v32 uint = (p2 * 2)
                if ((!false) || (42 <= -9223372036854775808)) {
                    fmt.Println(p0, fi6, v32)
                    p1 -= (p1 * p1)
                } else {
                    v7 = append(v7, 128)
                    if false {
                        var v33 uint32 = 0
                        if false {
                            v34 := GenS7{F0: int(p0)}
                            p1 += 18446744073709551615
                            _ = v34
                        }
                        if (true || false) {
                            switch fi6 {
                            case 636:
                                var v35 bool = (false && true)
                                _ = v35
                            }
                            v32 += v32
                            if true {
                                if true {
                                    p1 = p1
                                } else {
                                    if (9223372036854775807 > 9223372036854775807) {
                                        switch "a" {
                                        case "a":
                                            p2 = p2
                                        case "hello":
                                            v33 += v33
                                        }
                                        v36 := GenS2{F0: (p0 | (p0 ^ p0)), F1: -128, F2: 853}
                                        _ = v36
                                    } else {
                                        if (((!((!(42 != 0)) || true)) && (!(-2147483648 > 1))) || true) {
                                            v37 := GenS2{F0: (^(p0)), F1: -1, F2: 65535}
                                            _ = v37
                                        }
                                    }
                                }
                                p0 += p0
                            }
                        } else {
                            _ = len(v7)
                            var v38 uint16 = 368
                            _ = v38
                        }
                        _ = v33
                    }
                }
                _ = v32
            } else {
                var v39 string = "\n\t"
                p1 -= (^(p1))
                var v40 uint64 = (p1 & p1)
                _ = v39
                _ = v40
            }
            fmt.Println(fi12, p2, p2)
            _ = fi12
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 928
    v3 := make(map[uint32]bool)
    var v4 float32 = (float32(0.0) * -2.25)
    v5 := []string{"\n\t", "a"}
    _ = len(v3)
    var v41 float32 = 0.0
    if false {
        v2 -= v2
        switch "" {
        case "\n\t":
            for fi42 := 0; fi42 < 11; fi42++ {
                if false {
                    fmt.Println(fi42)
                } else {
                    v2 -= genFn0(223, 4294967295, 18446744073709551615)
                    var v43 uint = 187
                    if (2 == 4294967295) {
                        for fi44 := 0; fi44 < 13; fi44++ {
                            if true {
                                v4 -= (0.1 - 0.0)
                                v45 := GenS6{F0: 112, F1: -1}
                                v46 := GenS2{F0: 2, F1: -128, F2: 42}
                                _ = len(v3)
                                _ = v45
                                _ = v46
                            } else {
                                v47 := GenS8{F0: -9223372036854775808, F1: genFn0(42, 530, (v43 + v43)), F2: (v2 - (v2 / ((v2) | 1)))}
                                _ = v47
                            }
                            if true {
                                var v48 uint64 = 4294967295
                                var v49 uint8 = 194
                                _ = v48
                                _ = v49
                            } else {
                                if false {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            }
                            _ = fi44
                        }
                        var v50 int32 = -2147483648
                        v51 := GenS9{F0: 2, F1: fi42, F2: 18446744073709551615}
                        _ = v50
                        _ = v51
                    } else {
                        switch (fi42 - (fi42 * int(v43))) {
                        case -2147483648:
                            if (false || true) {
                                var v52 uint32 = 42
                                _ = v52
                            } else {
                                switch 11 {
                                case 32768:
                                    for fi53 := 0; fi53 < 7; fi53++ {
                                        if (!false) {
                                            v54 := GenS3{F0: 0, F1: 1, F2: 255}
                                            fmt.Println(v4)
                                            fmt.Println(fi42, fi42)
                                            _ = v54
                                        } else {
                                            delete(v1, 32768)
                                        }
                                        _ = fi53
                                    }
                                case 0:
                                    for fi55 := 0; fi55 < 2; fi55++ {
                                        var v56 bool = false
                                        v56 = (-901 != 2147483647)
                                        _ = fi55
                                        _ = v56
                                    }
                                }
                                if (v4 != float32((0.0 * (0.0 - ((1e10 - 3.14159) * float64(float32((0.1 * ((3.14159 / ((0.0 * -2.25) + (-2.25 - 0.0))) / 1.0))))))))) {
                                    fmt.Println(fi42)
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = "a"
                                    }
                                }
                            }
                        case 499:
                            v57 := GenS8{F0: 1, F1: (^(v2)), F2: genFn0(187, 18446744073709551615, (v43 >> (((v43 & (v43 & (v43 >> (((v43 * (v43 % ((279) | 1)))) & 63))))) & 63)))}
                            _ = v57
                        default:
                            for fi58 := 0; fi58 < 19; fi58++ {
                                v59 := GenS7{F0: (fi42 & (fi58 << ((141) & 63)))}
                                fmt.Println(v41)
                                _ = fi58
                                _ = v59
                            }
                        }
                        v4 += (-2.25 * v4)
                    }
                    _ = v43
                }
                v41 -= float32(1e10)
                v41 += 1e10
                _ = fi42
            }
        case "hello":
            for fi60 := 0; fi60 < 6; fi60++ {
                fmt.Println(v41, v2)
                for fi61 := 0; fi61 < 4; fi61++ {
                    v4 = v4
                    v62 := GenS3{F0: 9223372036854775807, F1: 2, F2: 1}
                    v4 = 1.5
                    switch 32767 {
                    case -740:
                        v1[1] = ("\n\t" + ((("" + "") + "") + ("hello" + (("foo bar" + ("hello" + ("a" + "hello"))) + "hello"))))
                    default:
                        v63 := GenS10{F0: (v2 | (v2 ^ 42)), F1: 824, F2: (v2 ^ v2)}
                        if false {
                            if true {
                                if false {
                                    v64 := GenS4{F0: -2147483648}
                                    var v65 bool = (!true)
                                    switch fi61 {
                                    case 0:
                                        if (42 < 4294967295) {
                                            if len(v5) > 0 {
                                                v5[0] = ""
                                            }
                                        }
                                    }
                                    _ = v64
                                    _ = v65
                                } else {
                                    v2 = 808
                                    v66 := GenS6{F0: 127, F1: -9223372036854775808}
                                    var v67 uint = 1
                                    _ = v66
                                    _ = v67
                                }
                            }
                            v2 = 65535
                            if false {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                delete(v1, (fi61 + (^(fi60))))
                                var v68 int8 = -128
                                if len(v5) > 0 {
                                    v5[0] = ""
                                }
                                _ = v68
                            }
                        }
                        _ = v63
                    }
                    _ = fi61
                    _ = v62
                }
                if true {
                    var v69 float32 = v4
                    switch 9223372036854775807 {
                    case -332:
                        _ = len(v3)
                    default:
                        for fi70 := 0; fi70 < 17; fi70++ {
                            v4 = v4
                            v41 = v4
                            v69 -= (v41 * v41)
                            if true {
                                switch "" {
                                case "hello":
                                    fmt.Println(v41)
                                default:
                                    v41 += 0.0
                                }
                                if true {
                                    var v71 uint16 = (v2 ^ genFn0(139, 0, 960))
                                    _ = v71
                                } else {
                                    _ = len(v1)
                                    v2 = (v2 << ((1) & 63))
                                    if false {
                                        v4 = v69
                                        var v72 int = fi70
                                        _ = v72
                                    } else {
                                        v69 -= v4
                                    }
                                    if (1 == 2147483647) {
                                        v73 := GenS3{F0: 42, F1: 47, F2: 128}
                                        var v74 float64 = ((3.14159 * (-2.25 + 3.14159)) + 3.14159)
                                        _ = len(v1)
                                        _ = len(v5)
                                        _ = v73
                                        _ = v74
                                    } else {
                                        var v75 int32 = 2147483647
                                        fmt.Println(v41, v2, v41)
                                        v76 := GenS7{F0: (fi70 >> ((445) & 63))}
                                        _ = v75
                                        _ = v76
                                    }
                                }
                            }
                            switch -304 {
                            case 178:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            default:
                                fmt.Println(v4, v69, v4)
                            }
                            _ = fi70
                        }
                    }
                    if len(v5) > 0 {
                        v5[0] = ("hello" + ("foo bar" + "\n\t"))
                    }
                    _ = v69
                }
                _ = fi60
            }
        }
    }
    var v77 int = -9223372036854775808
    delete(v3, 2)
    v78 := GenS3{F0: 9223372036854775807, F1: 42, F2: 255}
    for fi79 := 0; fi79 < 2; fi79++ {
        var v80 int8 = 109
        if false {
            v81 := GenS2{F0: 250, F1: 0, F2: 1}
            _ = v81
        } else {
            v77 += 674
            var v82 float32 = float32(0.1)
            _ = v82
        }
        _ = fi79
        _ = v80
    }
    v41 = ((v41 * (v41 + 1e10)) / v41)
    fmt.Println(v2)
    var v83 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v77:", v77)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v77
    _ = v78
    _ = v83
}
