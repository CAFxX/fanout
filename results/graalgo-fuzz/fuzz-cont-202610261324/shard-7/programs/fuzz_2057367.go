package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint32; F2 int64 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 uint64; F1 int32; F2 int16 }
type GenS7 struct { F0 uint8; F1 uint64; F2 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int16; F1 int16; F2 string }
type GenS11 struct { F0 uint; F1 int; F2 uint64 }
type GenS12 struct { F0 uint8; F1 int32; F2 int16 }
type GenS13 struct { F0 int; F1 uint16 }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: 85, F1: 42}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS2{F0: 18446744073709551615, F1: 0}
        switch (fi7 | fi7) {
        case 88:
            switch ("foo bar" + ("\n\t" + (("hello" + "a") + "\n\t"))) {
            case "a":
                var v9 int32 = 42
                _ = v9
            case "\n\t":
                v10 := GenS3{F0: 43, F1: 255, F2: 255}
                _ = v10
            case "foo bar":
                var v11 uint16 = 42
                _ = v11
            }
        default:
            var v12 bool = (((((0.1 / 1.0) / 0.1) <= ((((1.5 * 3.14159) * 1.5) / 0.1) - 1.5)) && (2147483647 == 1)) || (((!false) || (false && false)) && true))
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    v13 := []int{-1, 32768}
    if ((!true) || (2 > 65535)) {
        var v14 int16 = -856
        _ = v14
    } else {
        switch ("" + "hello") {
        case "a":
            v15 := GenS1{F0: 42, F1: 508}
            _ = v15
        case "hello":
            var v16 int8 = -1
            _ = v16
        case "":
            var v17 int = 554
            _ = v17
        default:
            switch "foo bar" {
            case "a":
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            case "hello":
                for fi18 := 0; fi18 < 15; fi18++ {
                    switch ("a" + ("\n\t" + (("a" + ((("hello" + ("hello" + (("a" + "a") + "hello"))) + ("" + "\n\t")) + "")) + ("foo bar" + (("a" + "foo bar") + ("" + ("\n\t" + "\n\t"))))))) {
                    case "\n\t":
                        var v19 string = "foo bar"
                        _ = v19
                    case "":
                        for fi20 := 0; fi20 < 16; fi20++ {
                            var v21 uint8 = 128
                            v21 = 102
                            var v22 float32 = ((-2.25 - 3.14159) + (float32(3.14159) - 3.14159))
                            switch fi20 {
                            case 777:
                                var v23 bool = false
                                _ = v23
                            case 960:
                                var v24 uint16 = 591
                                _ = v24
                            }
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                    var v25 int = (fi18 << ((2) & 63))
                    v25 += (v25 % (((v25 - fi18)) | 1))
                    _ = fi18
                    _ = v25
                }
            }
        }
        var v26 float64 = 0.0
        var v27 int8 = 1
        _ = v26
        _ = v27
    }
    if (float32((0.0 - (float64(float32(1.5)) + (-2.25 - ((-2.25 / ((3.14159 + (1.5 - (0.0 / 3.14159))) - ((1.5 / 3.14159) + 1.5))) * -2.25))))) >= 0.1) {
        if ((-2.25 / ((1.5 + (0.0 + 0.1)) * 1.5)) < (((float64(((3.14159 / 1.0) - (3.14159 - 3.14159))) - 3.14159) - 0.1) + 0.1)) {
            var v28 float64 = 3.14159
            var v29 uint64 = 2
            _ = v28
            _ = v29
        } else {
            var v30 uint32 = 2147483647
            _ = v30
        }
        var v31 string = ("foo bar" + "hello")
        var v32 uint64 = 0
        _ = v31
        _ = v32
    } else {
        for fi33 := 0; fi33 < 8; fi33++ {
            var v34 string = ("foo bar" + "")
            if (-128 != -128) {
                v34 = "hello"
            }
            _ = fi33
            _ = v34
        }
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 uint16 = 255
            if (3.14159 == (1.5 + (((1.5 / ((0.1 - 1.5) - -2.25)) / 3.14159) + 1.5))) {
                if (false && ((false || (!false)) || false)) {
                    v37 := GenS4{F0: 2, F1: 932, F2: -1}
                    for fi38 := 0; fi38 < 6; fi38++ {
                        switch fi38 {
                        case -1:
                            if false {
                                var v39 uint64 = uint64(v36)
                                v36 += (v36 + (^(v36)))
                                var v40 int32 = -2147483648
                                _ = v39
                                _ = v40
                            } else {
                                v41 := GenS2{F0: 739, F1: 42}
                                var v42 uint16 = (^(v36))
                                _ = v41
                                _ = v42
                            }
                            var v43 float32 = 3.14159
                            _ = v43
                        case -9223372036854775808:
                            if (((-385 == 42) || true) || true) {
                                v36 = 1
                            } else {
                                if false {
                                    if (32767 < int16(v36)) {
                                        if false {
                                            _ = len(v13)
                                            switch fi38 {
                                            case -159:
                                                v44 := GenS5{F0: 1, F1: -39}
                                                var v45 uint64 = 942
                                                if true {
                                                    fmt.Println(fi35, v45)
                                                    v46 := GenS6{F0: 523, F1: 553, F2: 32767}
                                                    v45 -= 0
                                                    _ = v46
                                                }
                                                if true {
                                                    var v47 int16 = -473
                                                    var v48 int64 = 0
                                                    var v49 float32 = 0.0
                                                    v36 -= v36
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    fmt.Println(fi35, v45)
                                                }
                                                _ = v44
                                                _ = v45
                                            case 42:
                                                var v50 string = ("" + "foo bar")
                                                _ = v50
                                            case -195:
                                                var v51 uint32 = 42
                                                _ = v51
                                            default:
                                                var v52 uint32 = 943
                                                _ = v52
                                            }
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                            v53 := GenS5{F0: -2147483648, F1: 1}
                                            if false {
                                                var v54 int16 = -1
                                                switch ("hello" + "hello") {
                                                case "a":
                                                    var v55 int32 = 0
                                                    _ = v55
                                                case "foo bar":
                                                    if false {
                                                        var v56 bool = false
                                                        _ = len(v13)
                                                        _ = v56
                                                    }
                                                case "\n\t":
                                                    var v57 float64 = ((0.0 - (0.0 * 0.0)) + 0.0)
                                                    _ = v57
                                                }
                                                v36 = v36
                                                if ((0 == -126) || (v36 == v36)) {
                                                    if false {
                                                        switch "hello" {
                                                        case "hello":
                                                            var v58 uint8 = 2
                                                            _ = v58
                                                        default:
                                                            v54 = (-(v54))
                                                        }
                                                    }
                                                    v59 := GenS6{F0: 386, F1: -1, F2: -32768}
                                                    v36 -= (^(v36))
                                                    _ = v59
                                                }
                                                v36 = v36
                                                _ = v54
                                            } else {
                                                v36 += 32767
                                                var v60 uint32 = 453
                                                v36 += (v36 - v36)
                                                _ = v60
                                            }
                                            _ = v53
                                        }
                                    }
                                    if false {
                                        var v61 int32 = 42
                                        _ = v61
                                    }
                                    fmt.Println(v36, fi38, fi38)
                                    fmt.Println(fi35, fi38)
                                    v36 = v36
                                } else {
                                    fmt.Println(fi38, v36, fi35)
                                }
                                var v62 uint64 = 990
                                fmt.Println(fi35, fi38, v36)
                                var v63 uint8 = 0
                                switch fi35 {
                                case 1:
                                    if len(v13) > 0 {
                                        v13[0] = fi38
                                    }
                                case 127:
                                    fmt.Println(fi38, v36, v36)
                                default:
                                    var v64 int8 = 42
                                    _ = v64
                                }
                                var v65 int = (fi35 % (((fi38 << ((2) & 63))) | 1))
                                _ = v62
                                _ = v63
                                _ = v65
                            }
                        case -196:
                            var v66 int8 = 37
                            _ = v66
                        default:
                            v67 := GenS3{F0: 9, F1: 42, F2: (^(v36))}
                            _ = v67
                        }
                        if (4294967295 >= 162) {
                            var v68 uint64 = 4294967295
                            v36 = v36
                            if len(v13) > 0 {
                                v13[0] = (fi35 & 0)
                            }
                            _ = v68
                        } else {
                            var v69 bool = (255 >= 65535)
                            switch ("foo bar" + "hello") {
                            case "":
                                v70 := GenS3{F0: 0, F1: 127, F2: (v36 ^ v36)}
                                _ = v70
                            }
                            fmt.Println(fi38, v69, v36)
                            _ = v69
                        }
                        var v71 int64 = int64(fi35)
                        _ = fi38
                        _ = v71
                    }
                    switch fi35 {
                    case -9223372036854775808:
                        var v72 int32 = 1
                        _ = v72
                    case 32768:
                        var v73 float64 = 0.1
                        _ = v73
                    default:
                        v36 = v36
                    }
                    _ = v37
                } else {
                    v74 := GenS7{F0: 255, F1: 4294967295, F2: true}
                    if len(v13) > 0 {
                        v13[0] = 0
                    }
                    _ = v74
                }
            } else {
                v36 = 712
                for fi75 := 0; fi75 < 16; fi75++ {
                    v36 = v36
                    _ = fi75
                }
            }
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = fi35
            _ = v36
        }
        var v76 float64 = -2.25
        _ = v76
    }
    _ = v13
    return 228
}
func genFn1(p0 bool, p1 int64) int16 {
    v77 := GenS8{F0: "a"}
    switch ("foo bar" + "") {
    case "hello":
        var v78 int32 = 0
        _ = v78
    case "\n\t":
        var v79 int16 = 699
        _ = v79
    case "":
        var v80 float32 = (((1e10 - 1.5) / 1.5) / (1e10 - float32(1.5)))
        _ = v80
    }
    p1 -= (p1 << ((188) & 63))
    _ = p0
    _ = p1
    _ = v77
    return -1
}

func main() {
    var v1 float32 = ((1.5 - 0.0) / 3.14159)
    var v2 int32 = -992
    var v3 uint = 65535
    v4 := make(map[int]int)
    v81 := GenS9{F0: -1}
    if (255 != 1) {
        for fi82 := 0; fi82 < 16; fi82++ {
            v83 := v4[fi82]
            _ = fi82
            _ = v83
        }
        v84 := []int{0, 4294967295, 1, 1}
        _ = v84
    } else {
        var v85 int32 = v2
        v1 += (0.0 + v1)
        var v86 uint16 = 455
        _ = v85
        _ = v86
    }
    v87 := []int{32767, -217}
    v3 += 36
    if (-9223372036854775808 <= 9223372036854775807) {
        var v88 float32 = ((1e10 + float32(1.5)) / (0.1 / 0.1))
        if true {
            for fi89 := 0; fi89 < 6; fi89++ {
                switch (-(fi89)) {
                case -272:
                    v4[fi89] = fi89
                case 221:
                    v1 += v88
                default:
                    fmt.Println(v3)
                }
                var v90 int64 = -94
                _ = fi89
                _ = v90
            }
            v2 = v2
            if true {
                var v91 float32 = v1
                if true {
                    switch -683 {
                    case -2147483648:
                        v1 = ((((v91 - (0.0 / 0.1)) + (v1 * 0.0)) * v91) * float32(-2.25))
                    }
                } else {
                    v88 += (v88 * (v88 / 3.14159))
                    var v92 uint16 = 65535
                    v2 += v2
                    _ = v92
                }
                _ = v91
            }
        } else {
            fmt.Println(v88)
        }
        _ = v88
    } else {
        for fi93 := 0; fi93 < 14; fi93++ {
            switch fi93 {
            case -435:
                var v94 float64 = 1.5
                _ = v94
            }
            v95 := GenS10{F0: -1, F1: genFn1(true, 0), F2: "a"}
            var v96 string = "\n\t"
            _ = fi93
            _ = v95
            _ = v96
        }
    }
    v3 = v3
    fmt.Println(v3)
    v1 += 1.5
    for fi97 := 0; fi97 < 11; fi97++ {
        var v98 uint8 = 2
        v99 := GenS11{F0: (^(v3)), F1: -217, F2: 2}
        var v100 uint32 = 244
        _ = fi97
        _ = v100
        _ = v98
        _ = v99
    }
    switch 209 {
    case 1:
        var v101 int8 = -72
        _ = v101
    case -475:
        if true {
            v3 = (v3 % (((v3 | (v3 ^ v3))) | 1))
            v1 -= 0.1
            var v102 uint32 = 0
            _ = v102
        }
    case -91:
        var v103 uint32 = 2
        _ = v103
    default:
        v104 := GenS1{F0: (^(v2)), F1: 1}
        _ = v104
    }
    v2 = 0
    if ((false && true) && (!(702 != 1))) {
        if (true && (428 >= 32768)) {
            v105 := GenS2{F0: v3, F1: 2147483647}
            _ = v105
        }
    } else {
        switch 504 {
        case 208:
            for fi106 := 0; fi106 < 10; fi106++ {
                if false {
                    for fi107 := 0; fi107 < 3; fi107++ {
                        if (false && (!true)) {
                            var v108 int8 = 100
                            _ = v108
                        } else {
                            var v109 uint = (v3 >> ((1) & 63))
                            v3 -= 338
                            _ = v109
                        }
                        _ = len(v87)
                        var v110 int8 = 42
                        _ = fi107
                        _ = v110
                    }
                    switch ("\n\t" + ("\n\t" + "foo bar")) {
                    case "\n\t":
                        if true {
                            switch "a" {
                            case "foo bar":
                                for fi111 := 0; fi111 < 7; fi111++ {
                                    var v112 int16 = 32767
                                    _ = fi111
                                    _ = v112
                                }
                            }
                            for fi113 := 0; fi113 < 20; fi113++ {
                                v114 := GenS12{F0: 128, F1: v2, F2: 1}
                                _ = len(v87)
                                _ = fi113
                                _ = v114
                            }
                            if true {
                                v115 := v4[42]
                                _ = v115
                            } else {
                                switch fi106 {
                                case 127:
                                    var v116 uint16 = 1
                                    _ = v116
                                default:
                                    var v117 int16 = -1
                                    _ = v117
                                }
                                if (v3 == v3) {
                                    v1 = v1
                                    v2 += v2
                                    for fi118 := 0; fi118 < 16; fi118++ {
                                        var v119 uint8 = 5
                                        fmt.Println(v3)
                                        if (!false) {
                                            var v120 int32 = v2
                                            _ = v120
                                        }
                                        _ = fi118
                                        _ = v119
                                    }
                                }
                                v1 = v1
                            }
                        }
                    }
                    fmt.Println(v3, v1, fi106)
                }
                _ = fi106
            }
        case -1:
            v87 = append(v87, -2147483648)
        }
        v3 = (v3 ^ (^(v3)))
        for fi121 := 0; fi121 < 15; fi121++ {
            v122 := GenS13{F0: (fi121 / ((int(v2)) | 1)), F1: 0}
            v1 += (v1 - v1)
            if true {
                v123 := GenS10{F0: -20, F1: 444, F2: "hello"}
                _ = v123
            }
            _ = fi121
            _ = v122
        }
    }
    var v124 int16 = 0
    for fi125 := 0; fi125 < 11; fi125++ {
        var v126 uint16 = 252
        _ = fi125
        _ = v126
    }
    v3 += 255
    _ = len(v87)
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v87:", len(v87), v87)
    _ = v1
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v81
    _ = v87
}
