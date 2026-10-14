package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int8; F1 int; F2 int64 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool; F1 uint; F2 uint64 }
type GenS10 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint16; F1 uint8 }
type GenS13 struct { F0 string; F1 uint16 }

func main() {
    v1 := []uint32{65535, 686, 0, 2}
    var v2 int16 = -1
    v3 := []uint32{246, 2147483647, 42, 65535}
    var v4 bool = (128 <= 188)
    v5 := []int{}
    v6 := []string{""}
    v7 := GenS1{F0: 4294967295, F1: 1, F2: 959}
    v4 = v4
    var v8 int8 = 0
    switch 652 {
    case 32768:
        var v9 int8 = (v8 << ((987) & 63))
        _ = v9
    case -293:
        fmt.Println(v4, v8)
    case -2147483648:
        var v10 uint64 = 0
        _ = v10
    default:
        if (((1e10 + float32(0.0)) - float32(float64(3.14159))) == (0.0 * ((0.1 - (float32(1.5) + ((1.5 - 1.5) - float32(0.0)))) - (0.0 / float32(1.5))))) {
            v4 = (v4 && false)
            var v11 bool = (!v4)
            _ = v11
        }
    }
    v12 := GenS1{F0: 1, F1: -609, F2: -1}
    for fi13 := 0; fi13 < 6; fi13++ {
        if (!true) {
            fmt.Println(v8)
        } else {
            if (v2 >= v2) {
                if v4 {
                    var v14 int32 = -92
                    _ = v14
                } else {
                    var v15 uint8 = 2
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v17 := GenS1{F0: uint64(v15), F1: -423, F2: 1}
                        if v4 {
                            v18 := GenS1{F0: 347, F1: -201, F2: -995}
                            var v19 int = -1
                            v15 = 127
                            _ = v18
                            _ = v19
                        } else {
                            var v20 int = (fi13 & (fi13 / ((0) | 1)))
                            switch 32767 {
                            case -9223372036854775808:
                                v21 := GenS1{F0: 1, F1: 42, F2: 42}
                                _ = v21
                            default:
                                v8 = (v8 | 0)
                            }
                            var v22 uint8 = 152
                            _ = v20
                            _ = v22
                        }
                        v8 = v8
                        var v23 float32 = 0.1
                        if true {
                            v23 = (v23 + ((((v23 * float32(0.1)) * (v23 / (v23 - (v23 + (0.1 - -2.25))))) * (3.14159 / (((v23 * (v23 / v23)) - v23) - 1.5))) - float32(0.0)))
                        }
                        switch "a" {
                        case "hello":
                            v5 = append(v5, -351)
                        case "foo bar":
                            if len(v5) > 0 {
                                v5[0] = 128
                            }
                            if true {
                                var v24 int = (fi13 << ((0) & 63))
                                fmt.Println(v8)
                                v8 = (v8 & v8)
                                v25 := GenS2{F0: 950, F1: ((2 <= 2) || false)}
                                var v26 int64 = -249
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                v15 += (^(v15))
                                v23 += ((0.1 / (v23 + v23)) - -2.25)
                            }
                            fmt.Println(v8, v23)
                            switch ("foo bar" + (("a" + (((("" + "hello") + ((("\n\t" + "foo bar") + ("a" + "hello")) + ("a" + ("foo bar" + ("" + "hello"))))) + "") + "hello")) + (((("a" + "hello") + ("hello" + ("hello" + "hello"))) + "hello") + ("" + "hello")))) {
                            case "\n\t":
                                var v27 int16 = v2
                                _ = v27
                            case "a":
                                v23 -= 0.0
                            case "":
                                var v28 int = (fi13 + fi16)
                                _ = v28
                            }
                            if v4 {
                                if ((32767 <= (fi16 - (fi13 * 42))) || (false && v4)) {
                                    var v29 float64 = 0.1
                                    _ = v29
                                } else {
                                    _ = len(v5)
                                }
                                v15 += v15
                                if (!false) {
                                    switch (fi13 % (((fi13 - (fi13 >> ((0) & 63)))) | 1)) {
                                    case 128:
                                        var v30 float32 = v23
                                        _ = v30
                                    case 127:
                                        v31 := GenS2{F0: 276, F1: (1 != 2147483647)}
                                        _ = v31
                                    default:
                                        var v32 uint16 = 0
                                        v8 = v8
                                        _ = v32
                                    }
                                    if (((v23 - float32((((float64(v23) - 0.1) / (float64(v23) / (0.0 + (3.14159 / 0.1)))) + 0.1))) > v23) || (fi16 == (fi16 & fi16))) {
                                        var v33 int64 = 1
                                        v5 = append(v5, 127)
                                        v33 = v33
                                        fmt.Println(fi13, fi13)
                                        _ = v33
                                    } else {
                                        var v34 uint = 65535
                                        var v35 int64 = int64(v2)
                                        switch fi16 {
                                        case 9223372036854775807:
                                            var v36 int = fi16
                                            _ = v36
                                        case 292:
                                            var v37 float64 = 3.14159
                                            _ = v37
                                        case 844:
                                            if v4 {
                                                fmt.Println(fi16, v2)
                                                if ((v23 >= ((float32((3.14159 * 1e10)) + float32(((-2.25 - 3.14159) * (3.14159 + 0.1)))) / v23)) || false) {
                                                    var v38 int64 = (^(v35))
                                                    v38 = (v38 - (v38 + v38))
                                                    _ = v38
                                                }
                                            } else {
                                                fmt.Println(v34, v4, v15)
                                            }
                                        }
                                        _ = v34
                                        _ = v35
                                    }
                                }
                            } else {
                                var v39 uint64 = 42
                                v40 := GenS1{F0: (v39 >> ((uint(v39)) & 63)), F1: -1, F2: 2147483647}
                                var v41 uint64 = (v39 << ((65535) & 63))
                                v8 += 100
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                            if (9223372036854775807 == 267) {
                                fmt.Println(v15, v23)
                                switch "foo bar" {
                                case "foo bar":
                                    v23 = ((v23 * -2.25) + (((v23 - 1e10) + float32(1e10)) + (v23 + float32(0.1))))
                                case "hello":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            }
                            v2 += (v2 + (v2 / (((^(v2))) | 1)))
                        case "\n\t":
                            v15 = (^(v15))
                        }
                        v2 = -32768
                        v8 = 75
                        _ = fi16
                        _ = v17
                        _ = v23
                    }
                    var v42 int32 = 1
                    _ = v15
                    _ = v42
                }
                var v43 uint16 = 2
                for fi44 := 0; fi44 < 8; fi44++ {
                    switch (fi44 / (((fi13 ^ -1)) | 1)) {
                    case -563:
                        if v4 {
                            var v45 string = (("" + "hello") + "foo bar")
                            v2 = v2
                            _ = v45
                        }
                    case 1:
                        if len(v1) > 0 {
                            v1[0] = 2147483647
                        }
                        switch "\n\t" {
                        case "hello":
                            if ((1e10 / 3.14159) > -2.25) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v8 += 127
                            }
                        case "foo bar":
                            if v4 {
                                v43 = (v43 & 42)
                                fmt.Println(fi44, fi44)
                            } else {
                                var v46 uint32 = 65535
                                _ = v46
                            }
                        case "\n\t":
                            switch ("foo bar" + "foo bar") {
                            case "\n\t":
                                if true {
                                    switch "hello" {
                                    case "foo bar":
                                        if v4 {
                                            v47 := GenS3{F0: (v8 ^ (v8 | 42)), F1: fi44, F2: 42}
                                            v4 = v4
                                            switch fi44 {
                                            case 19:
                                                v48 := GenS1{F0: 4294967295, F1: 481, F2: -530}
                                                _ = v48
                                            case 127:
                                                v4 = true
                                            default:
                                                v8 = v8
                                            }
                                            if false {
                                                switch (fi44 >> ((272) & 63)) {
                                                case 2147483647:
                                                    var v49 int32 = int32(v2)
                                                    _ = v49
                                                case -2147483648:
                                                    var v50 uint8 = 128
                                                    _ = v50
                                                default:
                                                    var v51 uint64 = 0
                                                    _ = v51
                                                }
                                            } else {
                                                v52 := GenS3{F0: (v8 << ((2) & 63)), F1: (fi44 >> ((0) & 63)), F2: 941}
                                                _ = v52
                                            }
                                            switch "\n\t" {
                                            case "":
                                                var v53 float32 = ((-2.25 * 1e10) * float32(1.5))
                                                _ = v53
                                            case "\n\t":
                                                var v54 uint = 4294967295
                                                _ = v54
                                            case "hello":
                                                v43 -= 32767
                                            default:
                                                if v4 {
                                                    var v55 uint32 = 2
                                                    v8 -= v8
                                                    if v4 {
                                                        var v56 uint16 = 2
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v57 := GenS4{F0: v2, F1: v4}
                                                            _ = v57
                                                        case "\n\t":
                                                            if ("foo bar" <= "") {
                                                                v58 := GenS3{F0: (v8 >> ((0) & 63)), F1: (fi13 ^ (fi13 << ((0) & 63))), F2: 9223372036854775807}
                                                                if v4 {
                                                                    if (false && false) {
                                                                        if (v4 && v4) {
                                                                            if (!(v2 != v2)) {
                                                                                v59 := GenS5{F0: 1}
                                                                                _ = v59
                                                                            } else {
                                                                                if false {
                                                                                    var v60 bool = v4
                                                                                    v61 := GenS5{F0: 492}
                                                                                    v2 += (v2 - v2)
                                                                                    v62 := GenS4{F0: 944, F1: (!(42 < 42))}
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                    _ = v62
                                                                                } else {
                                                                                    switch 32768 {
                                                                                    case 32767:
                                                                                        switch 128 {
                                                                                        case -9223372036854775808:
                                                                                            _ = len(v3)
                                                                                        case 416:
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = 0
                                                                                            }
                                                                                        default:
                                                                                            if v4 {
                                                                                                v63 := GenS6{F0: v55, F1: (v8 ^ -1)}
                                                                                                var v64 int64 = 491
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = v55
                                                                                                }
                                                                                                _ = v63
                                                                                                _ = v64
                                                                                            } else {
                                                                                                fmt.Println(fi44, v2, v56)
                                                                                                v65 := GenS1{F0: 1, F1: 183, F2: 42}
                                                                                                if v4 {
                                                                                                    switch (("foo bar" + (("" + "") + "\n\t")) + "foo bar") {
                                                                                                    case "hello":
                                                                                                        v66 := GenS2{F0: 562, F1: false}
                                                                                                        _ = v66
                                                                                                    default:
                                                                                                        fmt.Println(v2)
                                                                                                    }
                                                                                                    if ((0 >= 2147483647) && v4) {
                                                                                                        var v67 int32 = 1
                                                                                                        var v68 uint = 255
                                                                                                        var v69 bool = v4
                                                                                                        _ = v67
                                                                                                        _ = v68
                                                                                                        _ = v69
                                                                                                    } else {
                                                                                                        if v4 {
                                                                                                            var v70 uint8 = 1
                                                                                                            _ = v70
                                                                                                        } else {
                                                                                                            var v71 uint32 = v55
                                                                                                            _ = v71
                                                                                                        }
                                                                                                        var v72 uint32 = v55
                                                                                                        _ = v72
                                                                                                    }
                                                                                                    if v4 {
                                                                                                        if false {
                                                                                                            if len(v6) >= 2 {
                                                                                                                v6 = v6[1:len(v6)]
                                                                                                            }
                                                                                                            if (0 < 42) {
                                                                                                                v73 := GenS3{F0: 100, F1: (^(fi13)), F2: 0}
                                                                                                                v56 += (v56 >> ((18446744073709551615) & 63))
                                                                                                                var v74 int16 = (v2 + 1)
                                                                                                                _ = v73
                                                                                                                _ = v74
                                                                                                            }
                                                                                                            v75 := GenS7{F0: (v55 << ((0) & 63)), F1: 42}
                                                                                                            _ = v75
                                                                                                        } else {
                                                                                                            switch "" {
                                                                                                            case "":
                                                                                                                fmt.Println(v56)
                                                                                                            case "hello":
                                                                                                                _ = len(v3)
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    v5 = append(v5, 9223372036854775807)
                                                                                                    if false {
                                                                                                        if true {
                                                                                                            fmt.Println(v43, v8, v56)
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = v65
                                                                                            }
                                                                                        }
                                                                                    case -451:
                                                                                        v8 -= -1
                                                                                    default:
                                                                                        v8 -= v8
                                                                                    }
                                                                                }
                                                                                _ = len(v6)
                                                                                var v76 int16 = v2
                                                                                _ = v76
                                                                            }
                                                                            if v4 {
                                                                                fmt.Println(v43, v4)
                                                                            } else {
                                                                                fmt.Println(v56)
                                                                                if (0 > 232) {
                                                                                    var v77 int64 = -542
                                                                                    _ = v77
                                                                                }
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = (^(v55))
                                                                                }
                                                                            }
                                                                            v43 += uint16(v55)
                                                                            var v78 int = fi44
                                                                            var v79 float32 = ((1e10 * float32(float64((0.0 / 1e10)))) + -2.25)
                                                                            if ((^(v78)) < (fi44 - fi44)) {
                                                                                v43 = v43
                                                                                _ = len(v3)
                                                                            }
                                                                            v80 := GenS8{F0: -605}
                                                                            _ = v78
                                                                            _ = v79
                                                                            _ = v80
                                                                        } else {
                                                                            v81 := GenS9{F0: (1e10 >= float32(1e10)), F1: 264, F2: 2}
                                                                            _ = v81
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            var v82 int32 = 2147483647
                                                                            _ = v82
                                                                        }
                                                                    } else {
                                                                        if true {
                                                                            var v83 uint16 = 441
                                                                            _ = v83
                                                                        }
                                                                        switch (fi44 << ((1) & 63)) {
                                                                        case 2147483647:
                                                                            if v4 {
                                                                                v84 := GenS8{F0: (v2 % (((-(v2))) | 1))}
                                                                                _ = v84
                                                                            } else {
                                                                                v2 = (v2 << ((uint(v43)) & 63))
                                                                            }
                                                                        case 9223372036854775807:
                                                                            fmt.Println(v8)
                                                                        }
                                                                    }
                                                                    var v85 uint8 = 0
                                                                    var v86 float32 = float32(3.14159)
                                                                    _ = v85
                                                                    _ = v86
                                                                }
                                                                _ = v58
                                                            }
                                                        case "foo bar":
                                                            switch (fi44 << ((uint(v56)) & 63)) {
                                                            case 0:
                                                                if len(v3) > 0 {
                                                                    v3[0] = 42
                                                                }
                                                            case -111:
                                                                var v87 uint16 = (v56 >> ((973) & 63))
                                                                _ = v87
                                                            default:
                                                                v43 = 2
                                                            }
                                                        default:
                                                            if v4 {
                                                                v88 := GenS1{F0: 1, F1: 42, F2: -1}
                                                                v89 := GenS4{F0: v2, F1: (-2.25 < 0.0)}
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                        }
                                                        _ = v56
                                                    } else {
                                                        v90 := GenS10{F0: 1, F1: (v8 - (v8 << ((42) & 63))), F2: (v43 ^ v43)}
                                                        var v91 int32 = 2147483647
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                    _ = v55
                                                } else {
                                                    v2 += 0
                                                }
                                                v43 -= v43
                                                v92 := GenS5{F0: 433}
                                                _ = v92
                                            }
                                            fmt.Println(v4, fi44)
                                            _ = v47
                                        } else {
                                            v8 -= v8
                                            v93 := GenS11{F0: 2147483648}
                                            if v4 {
                                                v3 = append(v3, 2147483648)
                                                var v94 uint32 = 910
                                                _ = v94
                                            } else {
                                                fmt.Println(fi13, v4, v43)
                                                var v95 int16 = -44
                                                if (true && v4) {
                                                    var v96 uint8 = 166
                                                    var v97 uint32 = 4294967295
                                                    v4 = v4
                                                    _ = v96
                                                    _ = v97
                                                }
                                                var v98 uint = 1
                                                _ = v95
                                                _ = v98
                                            }
                                            _ = v93
                                        }
                                    default:
                                        v8 = int8(fi13)
                                    }
                                    if (v2 == -1) {
                                        if v4 {
                                            v2 -= (v2 << ((680) & 63))
                                            var v99 uint = 1
                                            _ = v99
                                        }
                                        _ = len(v6)
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                switch -322 {
                                case 32767:
                                    v4 = (-1 >= -2147483648)
                                case 42:
                                    if (v4 && v4) {
                                        v4 = v4
                                        var v100 uint = 256
                                        _ = v100
                                    }
                                case 1:
                                    v101 := GenS4{F0: v2, F1: v4}
                                    _ = v101
                                default:
                                    v4 = true
                                }
                                var v102 string = ""
                                _ = v102
                            case "foo bar":
                                if (1.5 > 1e10) {
                                    v43 -= (^(v43))
                                    var v103 uint16 = (^(v43))
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                    _ = v103
                                }
                            default:
                                v4 = (2 != (^(v43)))
                            }
                        }
                        v104 := GenS12{F0: v43, F1: 218}
                        _ = v104
                    case 2147483647:
                        var v105 int32 = -2147483648
                        _ = v105
                    }
                    v8 += (v8 % ((v8) | 1))
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                    _ = fi44
                }
                _ = v43
            } else {
                v106 := GenS1{F0: 18446744073709551615, F1: 9223372036854775807, F2: 26}
                var v107 uint8 = 0
                var v108 uint16 = 255
                _ = v106
                _ = v107
                _ = v108
            }
            v3 = append(v3, 646)
        }
        _ = fi13
    }
    for fi109 := 0; fi109 < 14; fi109++ {
        switch "" {
        case "":
            v2 = v2
        case "\n\t":
            v4 = false
        }
        v2 += (v2 | v2)
        _ = fi109
    }
    if (-1 >= 0) {
        if len(v5) > 0 {
            v5[0] = 32768
        }
    }
    var v110 uint64 = 1
    switch 1 {
    case 768:
        v8 -= (^(v8))
    }
    v111 := GenS11{F0: 1}
    if true {
        fmt.Println(v110, v8)
        var v112 float32 = (1e10 / (1.5 - 1e10))
        _ = v112
    } else {
        var v113 int16 = 1
        v114 := GenS13{F0: ("a" + "hello"), F1: 65535}
        _ = v113
        _ = v114
    }
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v110
    _ = v111
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
