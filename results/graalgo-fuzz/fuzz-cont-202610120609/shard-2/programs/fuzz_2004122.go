package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int32; F1 uint8; F2 int8 }
type GenS5 struct { F0 int16; F1 int16; F2 int32 }
type GenS6 struct { F0 uint8; F1 int; F2 int16 }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int64) uint32 {
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: 65535, F1: 0, F2: 42}
        _ = fi6
        _ = v7
    }
    p0 = (-(p0))
    var v8 float64 = 0.0
    _ = p0
    _ = v8
    return 603
}
func genFn1(p0 uint8) uint64 {
    var v9 int16 = 42
    if (true || (!false)) {
        var v10 int32 = 316
        p0 = (p0 + p0)
        _ = v10
    } else {
        var v11 string = "\n\t"
        if false {
            switch v11 {
            case "\n\t":
                if (356 >= 42) {
                    v12 := []int{128, -51, 1}
                    if (!false) {
                        if (false || true) {
                            for fi13 := 0; fi13 < 3; fi13++ {
                                for fi14 := 0; fi14 < 3; fi14++ {
                                    v11 = (v11 + ("\n\t" + ""))
                                    _ = fi14
                                }
                                v9 += v9
                                p0 = 0
                                _ = fi13
                            }
                            if (false && (42 != 946)) {
                                switch v11 {
                                case "":
                                    v15 := GenS1{F0: 759, F1: 255, F2: 42}
                                    _ = v15
                                default:
                                    var v16 int8 = 127
                                    _ = v16
                                }
                                switch v11 {
                                case "hello":
                                    fmt.Println(v9)
                                case "a":
                                    for fi17 := 0; fi17 < 5; fi17++ {
                                        if true {
                                            fmt.Println(v9, fi17, p0)
                                        } else {
                                            v18 := GenS1{F0: 32767, F1: p0, F2: 42}
                                            _ = v18
                                        }
                                        v9 -= v9
                                        if ((fi17 >> ((0) & 63)) > 4294967295) {
                                            v19 := GenS2{F0: 42, F1: false}
                                            v20 := GenS2{F0: -73, F1: true}
                                            _ = v19
                                            _ = v20
                                        } else {
                                            var v21 int8 = -7
                                            var v22 string = "a"
                                            var v23 int32 = -1
                                            _ = v21
                                            _ = v22
                                            _ = v23
                                        }
                                        _ = fi17
                                    }
                                case "\n\t":
                                    for fi24 := 0; fi24 < 9; fi24++ {
                                        fmt.Println(fi24, v11)
                                        var v25 uint8 = (p0 ^ (p0 % ((p0) | 1)))
                                        var v26 int64 = 9223372036854775807
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                default:
                                    for fi27 := 0; fi27 < 4; fi27++ {
                                        switch -20 {
                                        case -518:
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                        case -1:
                                            var v28 float64 = 1e10
                                            _ = v28
                                        default:
                                            v29 := GenS1{F0: 799, F1: p0, F2: 379}
                                            _ = v29
                                        }
                                        v9 = (v9 / ((-32768) | 1))
                                        v30 := GenS2{F0: -128, F1: true}
                                        _ = fi27
                                        _ = v30
                                    }
                                }
                            } else {
                                for fi31 := 0; fi31 < 10; fi31++ {
                                    var v32 uint8 = 255
                                    v11 = "foo bar"
                                    v12 = append(v12, 128)
                                    _ = fi31
                                    _ = v32
                                }
                            }
                            var v33 uint = 255
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 10; fi34++ {
                            for fi35 := 0; fi35 < 2; fi35++ {
                                p0 += 0
                                v36 := GenS2{F0: 50, F1: (42 == 0)}
                                _ = fi35
                                _ = v36
                            }
                            _ = fi34
                        }
                    }
                    fmt.Println(v9, v9)
                    _ = v12
                } else {
                    if (0.0 >= (((-2.25 - 0.1) - 0.0) + ((1.5 + ((0.1 + float32(3.14159)) / 1.5)) - 3.14159))) {
                        v37 := GenS3{F0: 100, F1: 18446744073709551615}
                        _ = v37
                    }
                    switch 0 {
                    case -357:
                        for fi38 := 0; fi38 < 1; fi38++ {
                            if false {
                                var v39 int32 = 703
                                if (!(v9 == (v9 - (v9 << ((2) & 63))))) {
                                    v11 = v11
                                }
                                switch (v11 + "foo bar") {
                                case "\n\t":
                                    var v40 uint16 = 255
                                    _ = v40
                                case "":
                                    v9 -= (v9 % (((^(v9))) | 1))
                                case "hello":
                                    if ((v39 > -2147483648) || (true && (((v39 >> ((2) & 63)) <= v39) && (!true)))) {
                                        fmt.Println(p0, v39)
                                        var v41 float32 = 0.0
                                        switch v11 {
                                        case "hello":
                                            var v42 string = v11
                                            _ = v42
                                        case "foo bar":
                                            v43 := []int{678, (fi38 + (fi38 - fi38)), -2147483648}
                                            _ = v43
                                        case "":
                                            for fi44 := 0; fi44 < 5; fi44++ {
                                                var v45 uint64 = 640
                                                switch (("foo bar" + "foo bar") + (v11 + ("\n\t" + v11))) {
                                                case "hello":
                                                    if true {
                                                        fmt.Println(fi44, v11)
                                                        if ((65535 > 2) || ("hello" >= ("hello" + (((v11 + (v11 + "hello")) + v11) + ("\n\t" + v11))))) {
                                                            var v46 bool = true
                                                            if ((((!false) && v46) || v46) || v46) {
                                                                var v47 string = "a"
                                                                var v48 int32 = (v39 | (v39 >> ((1) & 63)))
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                            if v46 {
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v11 = (v11 + v11)
                                                                case "hello":
                                                                    v49 := []int{0, 128, (-(fi44)), int(v39)}
                                                                    _ = v49
                                                                }
                                                            } else {
                                                                v9 = v9
                                                                var v50 uint16 = uint16(p0)
                                                                _ = v50
                                                            }
                                                            _ = v46
                                                        }
                                                    }
                                                    fmt.Println(v39, v41, v45)
                                                default:
                                                    v51 := []int{(fi44 / ((fi38) | 1)), (fi44 / ((-9223372036854775808) | 1)), (-(fi44))}
                                                    _ = v51
                                                }
                                                v11 = "a"
                                                _ = fi44
                                                _ = v45
                                            }
                                        default:
                                            p0 = (p0 * p0)
                                        }
                                        _ = v41
                                    }
                                default:
                                    p0 = p0
                                }
                                _ = v39
                            } else {
                                if false {
                                    v52 := GenS4{F0: 42, F1: (p0 * 1), F2: 36}
                                    _ = v52
                                } else {
                                    for fi53 := 0; fi53 < 12; fi53++ {
                                        var v54 bool = false
                                        _ = fi53
                                        _ = v54
                                    }
                                    v55 := GenS5{F0: (v9 % ((-32768) | 1)), F1: v9, F2: -565}
                                    _ = v55
                                }
                            }
                            _ = fi38
                        }
                    }
                }
            case "a":
                if (true || (false || (327 > 277))) {
                    fmt.Println(p0, p0, p0)
                    var v56 float64 = (-2.25 + ((0.0 + 0.0) * 3.14159))
                    _ = v56
                } else {
                    if ((!true) && false) {
                        v57 := GenS5{F0: (v9 ^ v9), F1: v9, F2: 891}
                        _ = v57
                    } else {
                        for fi58 := 0; fi58 < 9; fi58++ {
                            for fi59 := 0; fi59 < 8; fi59++ {
                                var v60 uint32 = 2
                                switch fi59 {
                                case 128:
                                    v61 := GenS4{F0: 0, F1: 106, F2: -128}
                                    _ = v61
                                case -435:
                                    v60 = v60
                                }
                                v62 := GenS6{F0: p0, F1: 0, F2: 538}
                                _ = fi59
                                _ = v60
                                _ = v62
                            }
                            _ = fi58
                        }
                        if true {
                            p0 -= (^(p0))
                            fmt.Println(v9)
                        }
                        for fi63 := 0; fi63 < 15; fi63++ {
                            switch fi63 {
                            case 236:
                                p0 = (^(p0))
                            }
                            _ = fi63
                        }
                    }
                    var v64 uint16 = 206
                    _ = v64
                }
            case "hello":
                v11 = (((v11 + "a") + (("a" + (("" + v11) + ((("a" + "a") + v11) + ("hello" + v11)))) + "")) + "")
            }
            var v65 float32 = (3.14159 / ((1.5 * -2.25) / -2.25))
            _ = v65
        }
        var v66 uint8 = 255
        _ = v11
        _ = v66
    }
    _ = p0
    _ = v9
    return 745
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ("\n\t" + "a")
    v3 := make(map[string]uint32)
    var v4 int8 = 100
    var v5 string = v2
    delete(v3, "\n\t")
    _ = len(v1)
    switch ("foo bar" + v2) {
    case "\n\t":
        v3[v2] = 1
    case "":
        v4 = v4
    }
    delete(v1, 9223372036854775807)
    switch v2 {
    case "hello":
        for fi67 := 0; fi67 < 10; fi67++ {
            fmt.Println(v2, fi67, v2)
            v68 := GenS7{F0: 9223372036854775807, F1: (-(fi67))}
            v3[v5] = 2
            _ = fi67
            _ = v68
        }
    case "":
        fmt.Println(v2, v4, v5)
    case "\n\t":
        switch (v5 + v2) {
        case "":
            var v69 uint = 813
            _ = v69
        }
    default:
        if true {
            for fi70 := 0; fi70 < 7; fi70++ {
                v4 = (v4 ^ (v4 - 68))
                _ = fi70
            }
        } else {
            if (32767 != -32768) {
                fmt.Println(v5)
                v4 = v4
                fmt.Println(v4)
            }
            switch 0 {
            case -2147483648:
                delete(v3, "hello")
            case 239:
                fmt.Println(v4, v5)
            }
            v5 = (v2 + (v2 + (v5 + v5)))
        }
    }
    for fi71 := 0; fi71 < 10; fi71++ {
        var v72 uint16 = 0
        _ = len(v3)
        _ = fi71
        _ = v72
    }
    for fi73 := 0; fi73 < 14; fi73++ {
        v74 := v1[(fi73 << ((288) & 63))]
        for fi75 := 0; fi75 < 1; fi75++ {
            if ((((v2 + (v5 + "a")) + (v2 + v2)) + v2) != "") {
                switch (fi75 % ((fi75) | 1)) {
                case 128:
                    v4 -= v4
                }
                v76 := v3["\n\t"]
                _ = v76
            }
            v77 := []int{}
            fmt.Println(v4)
            _ = fi75
            _ = v77
        }
        _ = fi73
        _ = v74
    }
    v5 = "hello"
    v78 := GenS3{F0: v4, F1: 595}
    if false {
        switch 2147483647 {
        case 32767:
            v79 := GenS8{F0: genFn1(162), F1: 432}
            _ = v79
        }
        if (!(2 <= 1)) {
            var v80 string = ((v5 + v5) + "")
            var v81 int32 = -1
            v2 = v2
            _ = v80
            _ = v81
        }
        for fi82 := 0; fi82 < 14; fi82++ {
            if (0 == 255) {
                var v83 uint = 4294967295
                v84 := []int{fi82, (fi82 | fi82), -1}
                var v85 float32 = ((float32(0.0) + (((1e10 * (((float32(0.1) + 1e10) + (float32((0.0 / (((0.0 - (1e10 / 0.1)) - (1e10 - (float64(((0.0 * 1e10) - (1e10 * 0.0))) - -2.25))) - 1e10))) - float32(1e10))) + 1e10)) - ((1.5 + (0.1 * 0.1)) - (3.14159 / 1e10))) * (1e10 / (((3.14159 - 1.5) + (((0.1 - (float32(3.14159) - -2.25)) - (0.1 - float32((((float64(0.1) + ((1e10 + 1.5) - 3.14159)) - -2.25) - float64(((((((3.14159 - ((3.14159 / 1.0) * (0.0 - 0.1))) / (1e10 / (-2.25 / 1.5))) * -2.25) - (((1e10 - ((1e10 + 0.1) - (0.1 * 1e10))) + 0.1) - 0.0)) - -2.25) / 1.5)))))) * 0.0)) * (1e10 * 3.14159))))) + 3.14159)
                _ = v83
                _ = v84
                _ = v85
            } else {
                v86 := []int{fi82, (fi82 * fi82), -629, fi82}
                var v87 string = "hello"
                delete(v1, -470)
                _ = v86
                _ = v87
            }
            _ = fi82
        }
    } else {
        for fi88 := 0; fi88 < 1; fi88++ {
            var v89 uint = 18446744073709551615
            _ = fi88
            _ = v89
        }
        v4 = v4
        for fi90 := 0; fi90 < 9; fi90++ {
            var v91 int16 = 32767
            v2 = ""
            var v92 uint8 = 51
            _ = fi90
            _ = v91
            _ = v92
        }
    }
    v2 = v5
    v2 = ""
    if (!false) {
        v5 = v5
        v3[v2] = 0
    } else {
        v93 := []int{0, -2147483648}
        fmt.Println(v2, v2)
        switch "hello" {
        case "hello":
            if ((!false) && false) {
                for fi94 := 0; fi94 < 20; fi94++ {
                    v2 = v5
                    var v95 int32 = 0
                    _ = fi94
                    _ = v95
                }
            }
        case "\n\t":
            v2 = (v5 + "hello")
        }
        _ = v93
    }
    var v96 string = (v5 + "")
    fmt.Println(v96)
    if false {
        switch -9223372036854775808 {
        case 9223372036854775807:
            for fi97 := 0; fi97 < 8; fi97++ {
                v98 := GenS6{F0: 2, F1: fi97, F2: 653}
                _ = fi97
                _ = v98
            }
        case -728:
            v99 := []int{}
            _ = v99
        default:
            v2 = (v96 + ("foo bar" + v5))
        }
        var v100 uint8 = 2
        _ = v100
    }
    switch "a" {
    case "foo bar":
        v101 := v1[4294967295]
        _ = v101
    default:
        var v102 int16 = -714
        _ = v102
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v78
    _ = v96
}
