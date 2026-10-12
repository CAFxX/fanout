package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint32; F1 int; F2 bool }
type GenS3 struct { F0 uint8; F1 uint64; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 uint32; F2 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int32; F1 string; F2 uint32 }
type GenS9 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS10 struct { F0 uint16; F1 uint16; F2 int }
func genFn0() int8 {
    var v4 uint = 905
    v4 = v4
    fmt.Println(v4)
    var v5 float32 = 3.14159
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 int64 = 1
    v2 := []uint32{4294967295, 2}
    v3 := make(map[uint32]bool)
    var v6 float32 = ((0.0 * 0.0) / 1e10)
    var v7 int16 = 1000
    fmt.Println(v1, v7, v1)
    switch -1 {
    case 354:
        var v8 uint16 = 255
        _ = v8
    }
    var v9 int8 = genFn0()
    v9 -= v9
    _ = len(v3)
    if true {
        var v10 int32 = -225
        _ = v10
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v11 := GenS1{F0: 563, F1: 1}
        _ = v11
    }
    switch ("" + "hello") {
    case "":
        if len(v2) > 0 {
            v2[0] = 65535
        }
    case "a":
        v7 = (v7 / (((-(v7))) | 1))
    default:
        if false {
            v6 += v6
        } else {
            for fi12 := 0; fi12 < 1; fi12++ {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v6 += float32(0.0)
                    v14 := GenS2{F0: 444, F1: (fi13 * -9223372036854775808), F2: (((v7 >> ((0) & 63)) < v7) || (163 < 784))}
                    v6 = 0.0
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
        }
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 float32 = v6
            var v18 string = "hello"
            fmt.Println(fi15)
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    fmt.Println(v1, v9)
    v19 := GenS3{F0: 127, F1: 2, F2: -658}
    if false {
        switch 4294967295 {
        case 4294967295:
            var v20 uint64 = 487
            _ = v20
        case -2147483648:
            for fi21 := 0; fi21 < 17; fi21++ {
                var v22 int8 = -1
                v23 := GenS1{F0: -629, F1: 788}
                _ = fi21
                _ = v22
                _ = v23
            }
        case -1:
            var v24 uint16 = 189
            _ = v24
        default:
            fmt.Println(v7, v1)
        }
    } else {
        var v25 uint = 97
        _ = v25
    }
    fmt.Println(v9, v6, v7)
    if (127 < 1) {
        if (1 != (v1 + (v1 << ((255) & 63)))) {
            var v26 int8 = (v9 << ((0) & 63))
            var v27 float64 = ((1e10 + (-2.25 - 0.0)) / -2.25)
            _ = v26
            _ = v27
        } else {
            v1 = (v1 << ((42) & 63))
            if (true || false) {
                v6 -= float32(1.5)
            }
            var v28 uint32 = 4294967295
            _ = v28
        }
        v7 += -1
        v9 += v9
    } else {
        var v29 uint = 67
        _ = v29
    }
    v2 = append(v2, 65535)
    fmt.Println(v7, v1)
    if false {
        v2 = append(v2, 2147483647)
        fmt.Println(v9, v6)
    } else {
        if (!false) {
            if true {
                var v30 uint = 65535
                _ = v30
            } else {
                if true {
                    for fi31 := 0; fi31 < 16; fi31++ {
                        var v32 uint64 = 18446744073709551615
                        for fi33 := 0; fi33 < 4; fi33++ {
                            v34 := GenS3{F0: 234, F1: (v32 % (((v32 ^ v32)) | 1)), F2: (-(fi33))}
                            _ = fi33
                            _ = v34
                        }
                        _ = fi31
                        _ = v32
                    }
                } else {
                    var v35 int = int(v7)
                    v6 = v6
                    _ = v35
                }
                for fi36 := 0; fi36 < 2; fi36++ {
                    var v37 uint32 = 2
                    for fi38 := 0; fi38 < 18; fi38++ {
                        switch (fi38 + (fi38 * fi36)) {
                        case -9223372036854775808:
                            if true {
                                fmt.Println(v9, fi38)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v39 := GenS4{F0: 0}
                                _ = v39
                            }
                            fmt.Println(v1, fi36)
                        case 9223372036854775807:
                            switch ("foo bar" + "hello") {
                            case "foo bar":
                                switch "" {
                                case "a":
                                    v40 := GenS1{F0: -1, F1: -7}
                                    _ = v40
                                case "foo bar":
                                    if false {
                                        fmt.Println(fi36, fi36, v6)
                                    }
                                case "\n\t":
                                    if true {
                                        v41 := v3[(^(v37))]
                                        var v42 string = "\n\t"
                                        var v43 uint = 4294967295
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    } else {
                                        v3[(^(v37))] = false
                                        var v44 int = fi38
                                        _ = v44
                                    }
                                default:
                                    var v45 uint32 = (v37 * v37)
                                    _ = v45
                                }
                                v1 -= v1
                                var v46 string = "a"
                                _ = v46
                            case "\n\t":
                                if len(v2) > 0 {
                                    v2[0] = v37
                                }
                                switch -2147483648 {
                                case 42:
                                    var v47 int8 = (-(v9))
                                    _ = v47
                                case -239:
                                    if ((-2147483648 < -139) || (1000 >= 263)) {
                                        v9 += (v9 / ((v9) | 1))
                                        v9 -= (v9 % (((v9 << ((346) & 63))) | 1))
                                    } else {
                                        var v48 uint16 = 42
                                        v1 -= (^(v1))
                                        v49 := GenS5{F0: (^(v7))}
                                        fmt.Println(v7, fi36, fi36)
                                        if ((v7 - v7) > v7) {
                                            if (!(!(v48 == v48))) {
                                                switch fi38 {
                                                case 32767:
                                                    v3[(^(v37))] = true
                                                    var v50 int8 = (v9 % ((-1) | 1))
                                                    v51 := v3[65535]
                                                    _ = v50
                                                    _ = v51
                                                case -9223372036854775808:
                                                    if (-862 < 1) {
                                                        v37 += (^(v37))
                                                        var v52 bool = true
                                                        var v53 uint64 = 943
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                case -791:
                                                    switch ("foo bar" + "hello") {
                                                    case "\n\t":
                                                        v54 := GenS3{F0: 42, F1: 0, F2: -2147483648}
                                                        _ = v54
                                                    case "foo bar":
                                                        var v55 uint16 = (v48 + v48)
                                                        _ = v55
                                                    }
                                                }
                                                if true {
                                                    v56 := GenS3{F0: uint8(v48), F1: 4294967295, F2: fi38}
                                                    _ = v56
                                                } else {
                                                    _ = len(v2)
                                                }
                                            }
                                            if true {
                                                v1 += 9223372036854775807
                                                if len(v2) > 0 {
                                                    v2[0] = (v37 - (^(v37)))
                                                }
                                                if true {
                                                    fmt.Println(v48)
                                                    fmt.Println(fi36, fi36, v7)
                                                    var v57 float32 = (v6 / v6)
                                                    _ = v57
                                                } else {
                                                    v2 = append(v2, uint32(v48))
                                                    switch "a" {
                                                    case "hello":
                                                        v58 := GenS6{F0: 128, F1: (v37 / ((v37) | 1)), F2: (42 >= 2)}
                                                        _ = v58
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = v37
                                                    }
                                                }
                                                if false {
                                                    var v59 int16 = v7
                                                    if true {
                                                        v48 = v48
                                                    }
                                                    if (!(!true)) {
                                                        var v60 float32 = v6
                                                        v60 = ((0.1 - ((v6 / 1.0) / v60)) - (v60 / 1e10))
                                                        var v61 int16 = (v59 % (((v59 << ((636) & 63))) | 1))
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        switch 32767 {
                                                        case -9223372036854775808:
                                                            v37 = v37
                                                        case -325:
                                                            if false {
                                                                v62 := GenS5{F0: 1}
                                                                if (1 <= 438) {
                                                                    var v63 int16 = (v59 | (v59 & (v59 & v59)))
                                                                    switch "" {
                                                                    case "hello":
                                                                        if ((false && (228 == 255)) && ((v48 + v48) != v48)) {
                                                                            if (!((^(v9)) <= (v9 & (v9 % (((v9 + (v9 & v9))) | 1))))) {
                                                                                fmt.Println(fi38)
                                                                                fmt.Println(v9)
                                                                                if (!false) {
                                                                                    if (v6 <= v6) {
                                                                                        var v64 int64 = (v1 * v1)
                                                                                        _ = v64
                                                                                    } else {
                                                                                        if true {
                                                                                            switch "\n\t" {
                                                                                            case "":
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                            case "a":
                                                                                                v3[(v37 + (v37 & 2147483647))] = (555 < 18446744073709551615)
                                                                                            case "\n\t":
                                                                                                v59 = v59
                                                                                            default:
                                                                                                _ = len(v2)
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v59 -= v59
                                                                                    if (39 < 90) {
                                                                                        if ((v48 & (v48 | 1)) <= (v48 + v48)) {
                                                                                            fmt.Println(v37, v48)
                                                                                            var v65 int = int(v7)
                                                                                            switch "a" {
                                                                                            case "hello":
                                                                                                var v66 int64 = (v1 * v1)
                                                                                                _ = v66
                                                                                            case "":
                                                                                                v9 = (v9 << ((735) & 63))
                                                                                            }
                                                                                            _ = v65
                                                                                        } else {
                                                                                            var v67 int32 = 1
                                                                                            v68 := GenS5{F0: (v7 * (^(v7)))}
                                                                                            var v69 int = (fi38 % ((fi38) | 1))
                                                                                            _ = v67
                                                                                            _ = v68
                                                                                            _ = v69
                                                                                        }
                                                                                        delete(v3, (^(v37)))
                                                                                    } else {
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = (v37 * v37)
                                                                                        }
                                                                                        fmt.Println(v1, v9, fi38)
                                                                                        _ = len(v2)
                                                                                        var v70 float32 = v6
                                                                                        v59 -= 32767
                                                                                        _ = v70
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        v71 := GenS5{F0: (-(v59))}
                                                                                        _ = v71
                                                                                    }
                                                                                    v6 = 1.5
                                                                                    var v72 uint16 = (v48 / (((v48 % (((v48 | (v48 / ((831) | 1)))) | 1))) | 1))
                                                                                    _ = v72
                                                                                } else {
                                                                                    var v73 string = ((("" + "") + "hello") + "\n\t")
                                                                                    _ = v73
                                                                                }
                                                                            } else {
                                                                                v74 := GenS4{F0: v48}
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = 42
                                                                                }
                                                                                switch (fi38 & 128) {
                                                                                case 2147483647:
                                                                                    fmt.Println(v48)
                                                                                default:
                                                                                    v48 = (^(v48))
                                                                                }
                                                                                _ = v74
                                                                            }
                                                                            var v75 uint32 = (v37 ^ (^(v37)))
                                                                            v9 += -1
                                                                            _ = v75
                                                                        }
                                                                    default:
                                                                        delete(v3, v37)
                                                                    }
                                                                    _ = v63
                                                                }
                                                                _ = v62
                                                            } else {
                                                                var v76 uint = 0
                                                                if false {
                                                                    var v77 int8 = genFn0()
                                                                    v78 := GenS2{F0: (v37 % ((v37) | 1)), F1: 529, F2: ((v1 & (-(v1))) > (-(v1)))}
                                                                    _ = v77
                                                                    _ = v78
                                                                } else {
                                                                    v79 := GenS5{F0: v7}
                                                                    v80 := GenS2{F0: (v37 << (((v76 | 0)) & 63)), F1: (fi38 % ((fi36) | 1)), F2: false}
                                                                    v6 += ((float32((-2.25 - 1e10)) / (float32(0.0) * ((float32(3.14159) * ((v6 - ((v6 - (float32(0.0) - v6)) + -2.25)) - v6)) / v6))) + v6)
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                                v9 = (v9 / (((v9 >> ((v76) & 63))) | 1))
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v37 = v37
                                                                delete(v3, 54)
                                                                _ = v76
                                                            }
                                                        }
                                                        var v81 int = 628
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v81
                                                    }
                                                    switch ("foo bar" + "\n\t") {
                                                    case "":
                                                        switch ("foo bar" + ("hello" + "hello")) {
                                                        case "\n\t":
                                                            fmt.Println(v59)
                                                        default:
                                                            switch (fi38 & fi38) {
                                                            case -71:
                                                                v48 += v48
                                                            case 9223372036854775807:
                                                                var v82 int = (fi38 / ((751) | 1))
                                                                _ = v82
                                                            case 42:
                                                                v83 := GenS7{F0: 0}
                                                                _ = v83
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v84 := GenS6{F0: 255, F1: 2, F2: true}
                                                        _ = v84
                                                    }
                                                    switch 127 {
                                                    case -2147483648:
                                                        if false {
                                                            var v85 int32 = 0
                                                            v6 = (3.14159 + -2.25)
                                                            v86 := GenS1{F0: int32(v59), F1: v59}
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                    case -529:
                                                        v1 = v1
                                                    }
                                                    _ = v59
                                                } else {
                                                    v87 := GenS8{F0: -1, F1: (("a" + "") + "hello"), F2: 387}
                                                    v9 -= (v9 | (-(v9)))
                                                    _ = v87
                                                }
                                            } else {
                                                v6 += v6
                                            }
                                            v9 = 127
                                            v88 := GenS3{F0: 128, F1: 18446744073709551615, F2: fi36}
                                            if false {
                                                if false {
                                                    if false {
                                                        v9 = (v9 & 0)
                                                    }
                                                    if (2147483648 != (v37 >> ((65535) & 63))) {
                                                        fmt.Println(v7, v48, v1)
                                                        v89 := GenS8{F0: -2147483648, F1: "foo bar", F2: (v37 >> ((18446744073709551615) & 63))}
                                                        _ = v89
                                                    } else {
                                                        var v90 float32 = (0.0 / v6)
                                                        v37 += v37
                                                        _ = v90
                                                    }
                                                    var v91 string = "a"
                                                    var v92 string = ""
                                                    var v93 uint16 = 42
                                                    _ = v91
                                                    _ = v92
                                                    _ = v93
                                                }
                                                var v94 float32 = (1e10 + (3.14159 * -2.25))
                                                v95 := GenS7{F0: 2147483647}
                                                _ = v94
                                                _ = v95
                                            } else {
                                                var v96 uint8 = 210
                                                v97 := GenS1{F0: 903, F1: (v7 >> ((42) & 63))}
                                                _ = v96
                                                _ = v97
                                            }
                                            var v98 float64 = ((float64(v6) + (float64((v6 * v6)) + 1.5)) / (0.1 - 0.0))
                                            v2 = append(v2, (^(v37)))
                                            _ = v88
                                            _ = v98
                                        }
                                        _ = v48
                                        _ = v49
                                    }
                                case -9223372036854775808:
                                    if (255 > 0) {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = (v37 | (v37 / (((^(v37))) | 1)))
                                        }
                                    } else {
                                        var v99 int8 = (v9 >> ((971) & 63))
                                        v99 = v9
                                        var v100 float32 = v6
                                        _ = v100
                                        _ = v99
                                    }
                                }
                                fmt.Println(fi38)
                            }
                        case 2147483647:
                            if true {
                                fmt.Println(v1)
                                var v101 int32 = 1
                                _ = v101
                            } else {
                                _ = len(v2)
                                var v102 int32 = 0
                                v103 := v3[507]
                                var v104 int64 = -9223372036854775808
                                if false {
                                    _ = len(v3)
                                    _ = len(v3)
                                    if v103 {
                                        switch 128 {
                                        case -888:
                                            delete(v3, v37)
                                        case -2147483648:
                                            var v105 uint64 = 1
                                            _ = v105
                                        case 388:
                                            switch "hello" {
                                            case "foo bar":
                                                var v106 uint8 = 0
                                                _ = v106
                                            case "\n\t":
                                                if (true || (0 >= 0)) {
                                                    v37 -= (v37 % ((v37) | 1))
                                                    v107 := GenS1{F0: 1, F1: (-(v7))}
                                                    v7 = (v7 / ((188) | 1))
                                                    var v108 float64 = 1.5
                                                    _ = v107
                                                    _ = v108
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = v37
                                    }
                                    var v109 int32 = v102
                                    v110 := GenS5{F0: (v7 / (((v7 >> ((656) & 63))) | 1))}
                                    if v103 {
                                        if (2 != 31) {
                                            var v111 int64 = (v1 + (v104 % ((-1) | 1)))
                                            _ = v111
                                        }
                                        switch "\n\t" {
                                        case "\n\t":
                                            v6 -= v6
                                        }
                                        v37 -= (v37 << ((301) & 63))
                                        v112 := GenS9{F0: 42, F1: (v109 >> ((65535) & 63)), F2: 231}
                                        fmt.Println(fi38)
                                        _ = v112
                                    } else {
                                        v37 = v37
                                        if v103 {
                                            fmt.Println(fi36, fi38, v7)
                                        } else {
                                            var v113 int8 = (^(v9))
                                            if v103 {
                                                fmt.Println(v104, v104)
                                            }
                                            _ = v113
                                        }
                                        var v114 uint16 = 255
                                        _ = v114
                                    }
                                    _ = v109
                                    _ = v110
                                }
                                delete(v3, (^(v37)))
                                var v115 float64 = 3.14159
                                _ = v102
                                _ = v103
                                _ = v104
                                _ = v115
                            }
                        }
                        _ = fi38
                    }
                    _ = fi36
                    _ = v37
                }
                for fi116 := 0; fi116 < 4; fi116++ {
                    v117 := GenS3{F0: 255, F1: 4294967295, F2: (fi116 ^ 561)}
                    switch "foo bar" {
                    case "a":
                        if (42 <= 127) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v1)
                            for fi118 := 0; fi118 < 17; fi118++ {
                                v119 := GenS10{F0: 773, F1: 1, F2: 42}
                                _ = len(v2)
                                _ = fi118
                                _ = v119
                            }
                        }
                    case "hello":
                        for fi120 := 0; fi120 < 5; fi120++ {
                            var v121 int16 = v7
                            var v122 float32 = (v6 * 1e10)
                            v6 -= (v6 - float32(float64((v122 * v122))))
                            if len(v2) > 0 {
                                v2[0] = 313
                            }
                            if true {
                                fmt.Println(v121, v6)
                                if true {
                                    var v123 int = (^(fi120))
                                    v124 := GenS3{F0: 255, F1: 42, F2: v123}
                                    v125 := GenS5{F0: (v121 - v7)}
                                    _ = v123
                                    _ = v124
                                    _ = v125
                                } else {
                                    fmt.Println(fi120, v121)
                                    v126 := v3[0]
                                    var v127 uint16 = 2
                                    _ = v126
                                    _ = v127
                                }
                                v128 := GenS7{F0: 89}
                                _ = v128
                            }
                            v121 -= -32768
                            _ = fi120
                            _ = v121
                            _ = v122
                        }
                    default:
                        v7 -= v7
                    }
                    switch "a" {
                    case "\n\t":
                        var v129 int64 = v1
                        _ = v129
                    case "":
                        var v130 float32 = v6
                        _ = v130
                    }
                    _ = fi116
                    _ = v117
                }
            }
            v3[0] = false
            v131 := v3[1]
            _ = v131
        }
        if (v7 <= (v7 * v7)) {
            v6 -= ((float32(0.1) * v6) * v6)
        } else {
            var v132 int16 = v7
            _ = v132
        }
        var v133 float32 = 1e10
        _ = v133
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
