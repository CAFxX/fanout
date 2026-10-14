package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 string; F2 uint32 }
type GenS4 struct { F0 string; F1 int8; F2 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint32; F1 int64; F2 uint16 }
func genFn0() string {
    v6 := GenS1{F0: 2147483648, F1: 1, F2: 2147483647}
    switch "\n\t" {
    case "foo bar":
        var v7 uint8 = 242
        _ = v7
    case "hello":
        v8 := []int{-2147483648, 127, 42, -443}
        _ = v8
    default:
        var v9 uint16 = 1
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS1{F0: 430, F1: 284, F2: -322}
        v12 := []int{fi10, (fi10 / (((fi10 | (fi10 % (((fi10 + (-(fi10)))) | 1)))) | 1))}
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: -1}
    _ = v13
    _ = v6
    return ("\n\t" + ("" + "hello"))
}
func genFn1(p0 uint64, p1 int16, p2 uint32) uint {
    p1 = p1
    var v14 uint32 = 42
    p0 = (p0 / ((p0) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "a"
    var v3 float64 = 0.1
    var v4 int8 = 0
    var v5 string = v2
    var v15 uint = 1
    var v16 uint8 = 42
    switch 42 {
    case 9223372036854775807:
        _ = len(v1)
    case 306:
        fmt.Println(v2, v4)
    default:
        var v17 float32 = 0.0
        _ = v17
    }
    var v18 uint32 = 2147483648
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS2{F0: (fi19 << (((v15 ^ v15)) & 63))}
        _ = fi19
        _ = v20
    }
    switch -2147483648 {
    case -226:
        switch v5 {
        case "foo bar":
            v4 -= (v4 / ((100) | 1))
        case "a":
            if (-2.25 >= ((0.1 + float32(v3)) / ((1.5 / float32(v3)) * float32((v3 + v3))))) {
                for fi21 := 0; fi21 < 13; fi21++ {
                    v3 = (v3 + 1.5)
                    _ = fi21
                }
                for fi22 := 0; fi22 < 1; fi22++ {
                    v23 := GenS3{F0: (v15 & 613), F1: ((((v5 + v5) + ("\n\t" + (v2 + ""))) + v2) + v2), F2: 1}
                    v24 := []int{}
                    var v25 uint32 = 2147483648
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                v18 = v18
            }
        }
    case 32768:
        var v26 string = v2
        _ = v26
    }
    switch 42 {
    case 32768:
        if (0.1 < 0.0) {
            fmt.Println(v5, v15)
            var v27 uint32 = (v18 & 65535)
            _ = v27
        } else {
            if true {
                v4 += -128
            }
            if ("\n\t" != (v2 + v2)) {
                var v28 bool = true
                v29 := []int{32767, 42, 4294967295, int(v18)}
                if false {
                    switch -927 {
                    case -769:
                        var v30 string = v2
                        _ = v30
                    }
                    fmt.Println(v3)
                } else {
                    for fi31 := 0; fi31 < 14; fi31++ {
                        if v28 {
                            fmt.Println(v28, v5)
                            if true {
                                _ = len(v29)
                                for fi32 := 0; fi32 < 14; fi32++ {
                                    fmt.Println(v5, v18)
                                    var v33 uint = v15
                                    _ = fi32
                                    _ = v33
                                }
                                if (1 <= 0) {
                                    if (0 <= 0) {
                                        v18 -= 65535
                                        v28 = (!((42 >= 1) && v28))
                                        v4 += (v4 % (((v4 << (((v15 * (v15 >> (((v15 + (^(v15)))) & 63)))) & 63))) | 1))
                                    }
                                    v34 := GenS1{F0: 0, F1: (v18 * (v18 - v18)), F2: -2147483648}
                                    switch (-(fi31)) {
                                    case -9223372036854775808:
                                        for fi35 := 0; fi35 < 18; fi35++ {
                                            _ = len(v29)
                                            v15 -= genFn1(444, -292, (v18 ^ v18))
                                            _ = fi35
                                        }
                                    default:
                                        var v36 uint32 = v18
                                        _ = v36
                                    }
                                    _ = v34
                                } else {
                                    _ = len(v29)
                                }
                            }
                        } else {
                            var v37 uint = (v15 & 730)
                            v38 := GenS2{F0: fi31}
                            _ = v37
                            _ = v38
                        }
                        if (!(true || false)) {
                            if len(v29) > 0 {
                                v29[0] = 32768
                            }
                        } else {
                            v16 = (v16 % (((^(v16))) | 1))
                            v39 := GenS4{F0: ("" + "foo bar"), F1: 42, F2: 42}
                            v16 -= v16
                            _ = v39
                        }
                        _ = fi31
                    }
                    v40 := GenS5{F0: -1}
                    v41 := GenS6{F0: 791}
                    _ = v40
                    _ = v41
                }
                _ = v28
                _ = v29
            } else {
                v42 := v1[-178]
                _ = v42
            }
        }
    case -663:
        if true {
            var v43 float64 = float64(1e10)
            for fi44 := 0; fi44 < 19; fi44++ {
                for fi45 := 0; fi45 < 18; fi45++ {
                    fmt.Println(v2, v2)
                    v46 := GenS6{F0: -1}
                    _ = fi45
                    _ = v46
                }
                v47 := GenS7{F0: (genFn1(802, 0, uint32(v16)) == v15), F1: 0, F2: (v18 % (((v18 << (((v15 / ((255) | 1))) & 63))) | 1))}
                var v48 int8 = (v4 | 42)
                _ = fi44
                _ = v47
                _ = v48
            }
            switch (v5 + ((v2 + (v5 + v5)) + ("a" + v2))) {
            case "hello":
                v49 := GenS5{F0: 779}
                _ = v49
            case "foo bar":
                var v50 float32 = (0.1 / (((1.5 * 0.1) / -2.25) / 0.1))
                _ = v50
            case "\n\t":
                var v51 uint32 = 415
                _ = v51
            }
            _ = v43
        }
    default:
        v52 := GenS2{F0: 32768}
        _ = v52
    }
    v53 := []int{}
    v53 = append(v53, 646)
    v54 := GenS7{F0: (false && true), F1: 248, F2: v18}
    v2 = ""
    _ = len(v53)
    v55 := v1[1]
    for fi56 := 0; fi56 < 20; fi56++ {
        var v57 uint64 = 18446744073709551615
        _ = fi56
        _ = v57
    }
    v58 := GenS2{F0: v55}
    if true {
        v59 := GenS2{F0: v55}
        _ = v59
    } else {
        var v60 float64 = float64((0.0 - 3.14159))
        switch "hello" {
        case "\n\t":
            for fi61 := 0; fi61 < 3; fi61++ {
                var v62 uint32 = v18
                _ = fi61
                _ = v62
            }
        default:
            switch ("hello" + "hello") {
            case "":
                switch v55 {
                case 2147483647:
                    if (-314 == -9223372036854775808) {
                        for fi63 := 0; fi63 < 5; fi63++ {
                            v64 := GenS8{F0: (v4 + (v4 << ((genFn1(4294967295, -190, v18)) & 63)))}
                            if false {
                                fmt.Println(v15, v4, v18)
                                if true {
                                    v60 = v60
                                    v55 = v55
                                } else {
                                    v15 = 1
                                    v65 := GenS3{F0: (v15 >> ((v15) & 63)), F1: ("" + "\n\t"), F2: (^(v18))}
                                    _ = v65
                                }
                                for fi66 := 0; fi66 < 18; fi66++ {
                                    v67 := v1[42]
                                    if (!((v16 % ((63) | 1)) >= v16)) {
                                        v1[796] = (fi63 | v55)
                                        if ((^(fi66)) <= (^(fi63))) {
                                            v2 = "hello"
                                            var v68 uint64 = uint64(v15)
                                            _ = v68
                                        } else {
                                            v5 = v5
                                        }
                                    }
                                    _ = fi66
                                    _ = v67
                                }
                            } else {
                                v2 = v2
                                v69 := GenS8{F0: -128}
                                _ = v69
                            }
                            v4 = (v4 % ((v4) | 1))
                            _ = fi63
                            _ = v64
                        }
                        var v70 int16 = 42
                        _ = v70
                    } else {
                        for fi71 := 0; fi71 < 9; fi71++ {
                            v72 := GenS9{F0: 42}
                            v4 = 1
                            v73 := GenS10{F0: 2}
                            _ = fi71
                            _ = v72
                            _ = v73
                        }
                        _ = len(v1)
                        switch v55 {
                        case -9223372036854775808:
                            v5 = v2
                        case 128:
                            delete(v1, int(v4))
                        }
                    }
                case -2147483648:
                    var v74 float32 = 1.5
                    _ = v74
                case 931:
                    for fi75 := 0; fi75 < 14; fi75++ {
                        v53 = append(v53, v55)
                        _ = fi75
                    }
                }
            case "a":
                var v76 int64 = -9223372036854775808
                _ = v76
            default:
                for fi77 := 0; fi77 < 11; fi77++ {
                    v55 += fi77
                    _ = fi77
                }
            }
        }
        switch v55 {
        case -926:
            if false {
                v53 = append(v53, v55)
                fmt.Println(v4)
                v2 = "a"
            } else {
                v55 = (v55 % (((v55 >> ((v15) & 63))) | 1))
                if false {
                    if false {
                        v60 = (((v60 + v60) * (v3 / v60)) * 1e10)
                        for fi78 := 0; fi78 < 2; fi78++ {
                            var v79 string = (((((("a" + v5) + v2) + (v2 + v5)) + v5) + v5) + v5)
                            v3 += (v3 - (float64(0.1) - (3.14159 * ((1e10 + -2.25) / ((1.5 - (1e10 - (0.1 / v3))) + ((float64(0.1) * (v3 / float64(1e10))) * v60))))))
                            fmt.Println(v60, v4, v5)
                            _ = fi78
                            _ = v79
                        }
                    } else {
                        fmt.Println(v2)
                        var v80 int32 = 2147483647
                        _ = v80
                    }
                    var v81 int32 = 0
                    _ = v81
                }
            }
        case 128:
            v82 := GenS1{F0: v18, F1: (v18 >> (((v15 >> ((genFn1(42, -200, (^(v18)))) & 63))) & 63)), F2: int32(v18)}
            _ = v82
        case 160:
            fmt.Println(v55, v16)
        }
        _ = v60
    }
    _ = len(v1)
    if true {
        for fi83 := 0; fi83 < 8; fi83++ {
            for fi84 := 0; fi84 < 18; fi84++ {
                if (-2.25 > v3) {
                    switch v5 {
                    case "a":
                        fmt.Println(fi83)
                        fmt.Println(fi83, fi84)
                        v18 = v18
                        var v85 uint16 = 42
                        _ = v85
                    case "\n\t":
                        switch ("hello" + v2) {
                        case "foo bar":
                            var v86 int8 = (v4 - v4)
                            fmt.Println(fi84, fi83)
                            _ = v86
                        default:
                            switch (("a" + v2) + v5) {
                            case "\n\t":
                                v87 := GenS11{F0: (v18 / (((v18 | v18)) | 1)), F1: 9223372036854775807, F2: 695}
                                _ = v87
                            default:
                                _ = len(v53)
                            }
                        }
                    }
                    v2 = v5
                }
                var v88 uint = (v15 ^ genFn1(360, -1, 807))
                _ = fi84
                _ = v88
            }
            switch v2 {
            case "a":
                if ((((("" > "\n\t") && true) && (!(true || true))) || false) || (!false)) {
                    fmt.Println(v3)
                    var v89 uint = genFn1(42, 0, (v18 & (v18 << ((v15) & 63))))
                    _ = v89
                }
            case "foo bar":
                v90 := GenS6{F0: 2147483647}
                _ = v90
            case "":
                for fi91 := 0; fi91 < 8; fi91++ {
                    _ = len(v53)
                    fmt.Println(v2, v55, v2)
                    v55 = v55
                    v3 += (1.5 - (3.14159 * (v3 / v3)))
                    v16 += v16
                    fmt.Println(v2, v2)
                    _ = fi91
                }
            default:
                if true {
                    v4 = int8(v16)
                }
            }
            _ = fi83
        }
        fmt.Println(v15, v2)
        for fi92 := 0; fi92 < 8; fi92++ {
            v93 := GenS1{F0: (v18 ^ v18), F1: (v18 % ((v18) | 1)), F2: -1}
            if true {
                switch "hello" {
                case "a":
                    var v94 bool = true
                    _ = v94
                case "\n\t":
                    v15 += v15
                case "hello":
                    v1[int(v15)] = fi92
                default:
                    fmt.Println(v15)
                }
                fmt.Println(v18, fi92, v5)
                fmt.Println(v16, v18, v4)
            }
            _ = fi92
            _ = v93
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v58
}
