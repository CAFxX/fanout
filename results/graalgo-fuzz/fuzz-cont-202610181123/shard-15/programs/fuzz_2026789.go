package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint16; F1 bool; F2 bool }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int; F1 int32; F2 uint }
type GenS5 struct { F0 string; F1 int32; F2 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int8; F1 int16; F2 uint64 }
type GenS10 struct { F0 uint32; F1 int64; F2 int64 }
type GenS11 struct { F0 int32; F1 uint; F2 int }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 bool; F1 uint8 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int64; F1 string; F2 uint32 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 bool; F1 uint16; F2 int8 }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 int16; F1 int32 }
type GenS20 struct { F0 uint64; F1 int32 }
type GenS21 struct { F0 uint16; F1 int64; F2 int32 }
type GenS22 struct { F0 uint; F1 int; F2 int16 }
type GenS23 struct { F0 int32 }

func main() {
    v1 := []uint32{2, 2147483648, 1, 1}
    v2 := make(map[int]bool)
    var v3 uint32 = 869
    v4 := make(map[uint32]uint32)
    var v5 int8 = -128
    if ((!(1e10 <= 0.0)) && false) {
        for fi6 := 0; fi6 < 7; fi6++ {
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 int = fi7
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        if false {
            switch 128 {
            case -2147483648:
                if false {
                    for fi9 := 0; fi9 < 3; fi9++ {
                        v1 = append(v1, (v3 % (((v3 >> ((255) & 63))) | 1)))
                        _ = fi9
                    }
                    fmt.Println(v5, v5)
                } else {
                    if false {
                        var v10 uint16 = 0
                        _ = v10
                    } else {
                        if len(v1) > 0 {
                            v1[0] = (v3 & v3)
                        }
                        v1 = append(v1, (^(v3)))
                        var v11 int = 2147483647
                        _ = v11
                    }
                }
            case 170:
                for fi12 := 0; fi12 < 2; fi12++ {
                    fmt.Println(v3)
                    var v13 uint8 = 42
                    v13 -= 128
                    _ = fi12
                    _ = v13
                }
            case 0:
                v3 += v3
            }
            var v14 int32 = 1
            _ = v14
        }
        v15 := GenS1{F0: 0, F1: (v3 >> ((4294967295) & 63))}
        _ = v15
    } else {
        var v16 uint64 = 1
        _ = v16
    }
    if ((!false) || false) {
        switch -678 {
        case 1:
            if true {
                var v17 uint8 = 79
                switch "a" {
                case "hello":
                    for fi18 := 0; fi18 < 11; fi18++ {
                        v19 := GenS1{F0: (v17 * v17), F1: v3}
                        var v20 int32 = 0
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                case "":
                    v3 = 2
                case "foo bar":
                    var v21 int64 = int64(v3)
                    _ = v21
                default:
                    switch 32767 {
                    case -107:
                        var v22 uint16 = 255
                        _ = v22
                    }
                }
                _ = v17
            } else {
                v5 = v5
                var v23 string = (("a" + "") + "hello")
                _ = len(v1)
                _ = v23
            }
        case 2147483647:
            v24 := GenS1{F0: 0, F1: v3}
            _ = v24
        default:
            if false {
                var v25 int = 9223372036854775807
                var v26 uint = 18446744073709551615
                var v27 uint16 = 2
                _ = v25
                _ = v26
                _ = v27
            } else {
                fmt.Println(v5, v3)
            }
        }
        v28 := GenS2{F0: 42, F1: false, F2: true}
        _ = v28
    } else {
        v29 := GenS3{F0: v5, F1: int16(v3)}
        _ = v29
    }
    v3 += 692
    switch "" {
    case "hello":
        var v30 float32 = 1e10
        _ = v30
    case "foo bar":
        if false {
            var v31 uint32 = v3
            _ = v31
        }
    }
    if (!(!((float64(1e10) > (((-2.25 + (1.5 + (float64(float32(-2.25)) - 3.14159))) / -2.25) + (1e10 - -2.25))) || (false && ((566 < 0) || (!false)))))) {
        for fi32 := 0; fi32 < 2; fi32++ {
            v33 := GenS3{F0: v5, F1: 0}
            _ = fi32
            _ = v33
        }
        v34 := GenS4{F0: -2147483648, F1: -816, F2: 1}
        if (2 > 202) {
            v5 += -1
        } else {
            for fi35 := 0; fi35 < 16; fi35++ {
                for fi36 := 0; fi36 < 1; fi36++ {
                    if false {
                        var v37 string = (("a" + "foo bar") + ("foo bar" + ("a" + (("" + "\n\t") + ""))))
                        v5 -= v5
                        v3 -= v3
                        _ = v37
                    } else {
                        if true {
                            var v38 bool = (true && (42 <= -685))
                            if v38 {
                                v38 = v38
                                v38 = v38
                            }
                            delete(v2, (fi36 * fi36))
                            _ = v38
                        } else {
                            switch 127 {
                            case 20:
                                v5 -= (v5 * (-(v5)))
                            default:
                                var v39 int8 = -128
                                _ = v39
                            }
                        }
                        v5 += v5
                    }
                    var v40 int16 = 32767
                    _ = fi36
                    _ = v40
                }
                _ = fi35
            }
            v3 = (v3 / (((^(v3))) | 1))
            if (845 > 4294967295) {
                var v41 uint8 = 141
                _ = v41
            } else {
                switch "a" {
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "\n\t":
                    var v42 int32 = -1
                    _ = v42
                case "":
                    for fi43 := 0; fi43 < 7; fi43++ {
                        var v44 float64 = 1e10
                        fmt.Println(v3, v44)
                        for fi45 := 0; fi45 < 20; fi45++ {
                            if false {
                                v3 += (v3 ^ 1)
                                v44 += float64(0.1)
                            }
                            v46 := GenS2{F0: 32767, F1: false, F2: ("foo bar" > "a")}
                            if false {
                                fmt.Println(v5, fi43, fi45)
                                if (true && false) {
                                    var v47 bool = (((true && false) || (false || false)) || false)
                                    switch ((((((("" + "foo bar") + (("foo bar" + "") + "")) + ("\n\t" + ("a" + (("" + "") + "hello")))) + ("" + "foo bar")) + (((((("a" + "") + "") + ("hello" + "")) + "foo bar") + "a") + "hello")) + "") + "") {
                                    case "\n\t":
                                        if (((^(fi43)) <= fi45) && true) {
                                            v47 = v47
                                            fmt.Println(fi45, v47, fi43)
                                            if (1 == (v5 / ((16) | 1))) {
                                                var v48 uint16 = 42
                                                v49 := v2[4294967295]
                                                _ = v48
                                                _ = v49
                                            } else {
                                                switch ("\n\t" + ("" + (("a" + ("hello" + ("a" + "hello"))) + ("" + "a")))) {
                                                case "foo bar":
                                                    var v50 uint8 = uint8(v3)
                                                    if (v3 < (v3 - v3)) {
                                                        fmt.Println(fi45, v3, fi45)
                                                        v51 := GenS3{F0: (-(v5)), F1: 1}
                                                        _ = v51
                                                    } else {
                                                        if true {
                                                            v52 := GenS5{F0: ("hello" + "hello"), F1: 758, F2: 596}
                                                            var v53 float64 = (((0.0 * (float64(-2.25) + 0.1)) / 0.1) + v44)
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                    }
                                                    var v54 int16 = 32767
                                                    _ = v50
                                                    _ = v54
                                                default:
                                                    v47 = false
                                                }
                                                if v47 {
                                                    v55 := GenS3{F0: (v5 >> ((825) & 63)), F1: 214}
                                                    v56 := GenS6{F0: fi43}
                                                    if len(v1) > 0 {
                                                        v1[0] = v3
                                                    }
                                                    _ = v55
                                                    _ = v56
                                                }
                                                var v57 string = "\n\t"
                                                if v47 {
                                                    v5 = (v5 / (((v5 - v5)) | 1))
                                                }
                                                _ = v57
                                            }
                                            var v58 uint16 = 713
                                            _ = v58
                                        }
                                    default:
                                        v3 -= v3
                                    }
                                    switch fi43 {
                                    case -9223372036854775808:
                                        var v59 float32 = (float32(float64(((1e10 - 3.14159) / (-2.25 * 3.14159)))) / 3.14159)
                                        _ = v59
                                    }
                                    v47 = (v3 == v3)
                                    switch "a" {
                                    case "a":
                                        var v60 int64 = -1
                                        _ = v60
                                    case "\n\t":
                                        v1 = append(v1, (^(v3)))
                                    case "foo bar":
                                        switch 128 {
                                        case -210:
                                            var v61 uint16 = 2
                                            _ = v61
                                        case 4294967295:
                                            v44 = v44
                                        case 128:
                                            v62 := GenS4{F0: fi43, F1: 42, F2: 4294967295}
                                            _ = v62
                                        }
                                    default:
                                        switch "foo bar" {
                                        case "":
                                            if (v47 || ((!(v47 && true)) || ((v5 == (v5 * v5)) && v47))) {
                                                var v63 uint16 = 0
                                                switch "" {
                                                case "foo bar":
                                                    delete(v4, (v3 << ((65535) & 63)))
                                                case "\n\t":
                                                    if len(v1) > 0 {
                                                        v1[0] = (v3 | v3)
                                                    }
                                                default:
                                                    v3 = v3
                                                }
                                                v64 := GenS7{F0: ("hello" + ("a" + "hello")), F1: 0}
                                                _ = v63
                                                _ = v64
                                            } else {
                                                v47 = v47
                                                v5 -= -117
                                                v65 := GenS8{F0: -54}
                                                _ = v65
                                            }
                                        case "\n\t":
                                            v66 := GenS7{F0: (((("foo bar" + ("a" + "\n\t")) + "hello") + ((("foo bar" + "hello") + "hello") + ("foo bar" + "a"))) + "\n\t"), F1: 18446744073709551615}
                                            _ = v66
                                        case "foo bar":
                                            switch -276 {
                                            case 2147483647:
                                                switch (fi43 * fi45) {
                                                case 2147483647:
                                                    v44 = v44
                                                    if (1.5 <= -2.25) {
                                                        var v67 float64 = 3.14159
                                                        if v47 {
                                                            var v68 uint16 = 255
                                                            if ((0.1 / ((3.14159 - (1.5 + -2.25)) * float32(v44))) < (1.5 / 1.0)) {
                                                                switch (fi43 + fi43) {
                                                                case 9223372036854775807:
                                                                    v44 -= v67
                                                                default:
                                                                    v69 := v2[(fi43 * fi43)]
                                                                    _ = v69
                                                                }
                                                                if (144 <= uint8(v3)) {
                                                                    if (!(!(1.5 >= v67))) {
                                                                        if (42 == 18446744073709551615) {
                                                                            v70 := GenS6{F0: (fi45 * 127)}
                                                                            var v71 uint = 255
                                                                            switch 32767 {
                                                                            case 0:
                                                                                v72 := GenS9{F0: 42, F1: 1, F2: 1}
                                                                                var v73 int64 = -394
                                                                                v74 := GenS1{F0: 0, F1: v3}
                                                                                _ = v72
                                                                                _ = v73
                                                                                _ = v74
                                                                            default:
                                                                                if (!(v47 && v47)) {
                                                                                    var v75 int64 = 9223372036854775807
                                                                                    var v76 uint = v71
                                                                                    v77 := GenS10{F0: (v3 / (((v3 % ((1) | 1))) | 1)), F1: (^(v75)), F2: (v75 ^ (v75 << (((v71 / ((259) | 1))) & 63)))}
                                                                                    v68 += v68
                                                                                    v71 = (^(v76))
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                } else {
                                                                                    if (2 != 128) {
                                                                                        fmt.Println(v68)
                                                                                        v47 = ((!(-32768 > 1)) || (((0.0 - -2.25) * ((0.0 - ((3.14159 - 1e10) - (0.0 + 0.1))) * (((0.0 * 0.0) + 1.5) + 0.1))) >= (((0.1 + (((-2.25 / (-2.25 - 3.14159)) - 0.0) + (1e10 * 0.0))) / 0.1) + float32((float64(0.1) - v44)))))
                                                                                        if true {
                                                                                            v78 := GenS8{F0: 40}
                                                                                            v79 := GenS10{F0: 864, F1: 987, F2: -1}
                                                                                            _ = v78
                                                                                            _ = v79
                                                                                        } else {
                                                                                            var v80 uint16 = v68
                                                                                            _ = len(v2)
                                                                                            _ = v80
                                                                                        }
                                                                                    } else {
                                                                                        var v81 int = 4294967295
                                                                                        _ = v81
                                                                                    }
                                                                                    if ((v44 >= (-2.25 - 1.5)) && (-1 <= 0)) {
                                                                                        v44 += (v44 * (v67 - v67))
                                                                                        v82 := GenS5{F0: ("a" + "foo bar"), F1: 306, F2: (fi43 << ((v71) & 63))}
                                                                                        _ = v82
                                                                                    }
                                                                                    var v83 string = ""
                                                                                    _ = v83
                                                                                }
                                                                            }
                                                                            _ = v70
                                                                            _ = v71
                                                                        }
                                                                    }
                                                                    var v84 uint16 = (^(v68))
                                                                    var v85 uint8 = uint8(v3)
                                                                    _ = v84
                                                                    _ = v85
                                                                } else {
                                                                    if v47 {
                                                                        v67 -= 0.0
                                                                    } else {
                                                                        v44 += v67
                                                                        var v86 uint = 255
                                                                        var v87 float32 = (((-2.25 * (0.1 - (0.0 * 1e10))) / -2.25) + 3.14159)
                                                                        _ = v86
                                                                        _ = v87
                                                                    }
                                                                }
                                                                v5 -= (v5 ^ (-(v5)))
                                                            } else {
                                                                if (false || false) {
                                                                    v3 += (v3 << ((65535) & 63))
                                                                    v88 := v2[-118]
                                                                    _ = v88
                                                                } else {
                                                                    v47 = (!((1.5 + 3.14159) >= 0.1))
                                                                }
                                                            }
                                                            _ = v68
                                                        }
                                                        v44 = v44
                                                        _ = v67
                                                    }
                                                    var v89 uint8 = 72
                                                    _ = v89
                                                case -178:
                                                    v90 := GenS11{F0: 0, F1: 0, F2: (^(fi43))}
                                                    _ = v90
                                                case -1:
                                                    v91 := GenS5{F0: "a", F1: -696, F2: fi43}
                                                    _ = v91
                                                default:
                                                    delete(v4, 2)
                                                }
                                            case 4294967295:
                                                v47 = ((true && false) && (v47 || (((!((v3 << ((4294967295) & 63)) > (^(v3)))) && v47) || false)))
                                            default:
                                                if v47 {
                                                    var v92 int = (fi43 + int(v5))
                                                    v93 := GenS1{F0: 2, F1: v3}
                                                    v94 := GenS7{F0: ((("hello" + (("a" + ("" + "foo bar")) + (("hello" + "\n\t") + ("" + "")))) + ("\n\t" + ("hello" + (("\n\t" + ((("foo bar" + "") + "\n\t") + "foo bar")) + "hello")))) + ("foo bar" + "a")), F1: 2}
                                                    _ = v92
                                                    _ = v93
                                                    _ = v94
                                                } else {
                                                    var v95 string = "hello"
                                                    _ = v95
                                                }
                                                v96 := GenS7{F0: "\n\t", F1: 2}
                                                if v47 {
                                                    var v97 int32 = -1
                                                    var v98 string = (("foo bar" + ("foo bar" + (("\n\t" + ("foo bar" + "foo bar")) + "foo bar"))) + (("" + "foo bar") + ""))
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    fmt.Println(v47)
                                                }
                                                _ = v96
                                            }
                                        }
                                    }
                                    var v99 int64 = 993
                                    v99 = (^(v99))
                                    _ = v47
                                    _ = v99
                                } else {
                                    if (true && false) {
                                        if false {
                                            v100 := GenS12{F0: 42}
                                            _ = v100
                                        }
                                        v44 = v44
                                        var v101 float64 = v44
                                        _ = v101
                                    }
                                    var v102 uint64 = 42
                                    switch 1 {
                                    case -2147483648:
                                        switch "" {
                                        case "\n\t":
                                            v103 := GenS2{F0: 255, F1: false, F2: (true || (121 <= 1))}
                                            _ = v103
                                        case "a":
                                            if (1000 >= 32767) {
                                                fmt.Println(fi45, fi45, v3)
                                                switch "\n\t" {
                                                case "a":
                                                    v4[354] = 339
                                                case "":
                                                    var v104 bool = (47 >= 446)
                                                    if ((((float64(3.14159) - 0.1) + v44) < v44) && false) {
                                                        var v105 bool = v104
                                                        v106 := GenS11{F0: 42, F1: 0, F2: int(v102)}
                                                        _ = v105
                                                        _ = v106
                                                    } else {
                                                        fmt.Println(v5, v3)
                                                        var v107 int16 = -87
                                                        _ = v107
                                                    }
                                                    v108 := GenS13{F0: ((3.14159 >= 1.5) || false), F1: 29}
                                                    _ = v104
                                                    _ = v108
                                                case "hello":
                                                    v4[2147483648] = uint32(v102)
                                                default:
                                                    var v109 uint8 = 161
                                                    _ = v109
                                                }
                                                v110 := GenS14{F0: fi45}
                                                _ = v110
                                            }
                                        case "hello":
                                            v3 -= v3
                                        }
                                    }
                                    _ = v102
                                }
                            } else {
                                if false {
                                    v111 := GenS12{F0: 1}
                                    _ = v111
                                } else {
                                    switch "a" {
                                    case "a":
                                        switch (fi43 << ((1) & 63)) {
                                        case -432:
                                            v44 = 0.1
                                        case 0:
                                            v5 = v5
                                            var v112 int8 = -128
                                            var v113 uint16 = 2
                                            _ = v112
                                            _ = v113
                                        case 1:
                                            var v114 int8 = 7
                                            _ = v114
                                        default:
                                            var v115 uint8 = 2
                                            _ = v115
                                        }
                                        switch ((("foo bar" + "foo bar") + ("a" + "")) + "\n\t") {
                                        case "foo bar":
                                            var v116 string = (("" + ("foo bar" + "foo bar")) + ("" + "a"))
                                            if false {
                                                var v117 uint8 = 1
                                                v3 += v3
                                                _ = v117
                                            }
                                            _ = v116
                                        case "\n\t":
                                            var v118 float64 = v44
                                            _ = v118
                                        case "":
                                            switch (fi43 & (fi45 | (-(fi45)))) {
                                            case -9223372036854775808:
                                                var v119 int16 = -1
                                                _ = v119
                                            }
                                        default:
                                            switch (fi45 + (-(fi45))) {
                                            case 0:
                                                switch "hello" {
                                                case "":
                                                    v5 = 0
                                                default:
                                                    fmt.Println(fi45, fi45, v5)
                                                }
                                            case 414:
                                                v44 = v44
                                            }
                                        }
                                    case "hello":
                                        var v120 uint64 = 18446744073709551615
                                        _ = v120
                                    case "\n\t":
                                        _ = len(v4)
                                        var v121 uint64 = 2
                                        v121 -= v121
                                        switch fi45 {
                                        case 32767:
                                            v5 -= (-(v5))
                                        default:
                                            var v122 uint32 = 1
                                            _ = v122
                                        }
                                        v3 = v3
                                        _ = v121
                                    }
                                    v123 := GenS15{F0: int64(v5), F1: "a", F2: 2}
                                    v5 = v5
                                    if true {
                                        v124 := GenS2{F0: 879, F1: true, F2: (!(-9223372036854775808 < 159))}
                                        v5 = (v5 | (^(v5)))
                                        fmt.Println(fi45)
                                        switch "" {
                                        case "":
                                            var v125 int8 = (^(v5))
                                            switch "" {
                                            case "foo bar":
                                                if (127 != 42) {
                                                    if (0 < 124) {
                                                        if false {
                                                            v126 := GenS16{F0: 42}
                                                            v127 := GenS3{F0: v125, F1: -1}
                                                            _ = v126
                                                            _ = v127
                                                        } else {
                                                            if ((false && (255 > 170)) && (fi43 >= (fi43 % ((-25) | 1)))) {
                                                                switch fi45 {
                                                                case 32767:
                                                                    fmt.Println(v44, v44)
                                                                }
                                                            } else {
                                                                v125 += v5
                                                                v5 += (v125 - v5)
                                                            }
                                                            v3 -= 2
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                switch (fi45 / ((fi45) | 1)) {
                                                                case -9223372036854775808:
                                                                    fmt.Println(v3, fi43, v125)
                                                                case 301:
                                                                    if (fi45 <= int(v125)) {
                                                                        var v128 int8 = v125
                                                                        if len(v1) > 0 {
                                                                            v1[0] = v3
                                                                        }
                                                                        _ = v128
                                                                    }
                                                                    var v129 int16 = -410
                                                                    _ = v129
                                                                case 9223372036854775807:
                                                                    var v130 uint = 903
                                                                    _ = v130
                                                                }
                                                            case "a":
                                                                if false {
                                                                    if false {
                                                                        fmt.Println(fi45, fi45, v5)
                                                                        var v131 int8 = 22
                                                                        v131 -= v131
                                                                        _ = v131
                                                                    } else {
                                                                        if true {
                                                                            var v132 uint8 = 128
                                                                            v5 -= (-(v125))
                                                                            _ = v132
                                                                        } else {
                                                                            v5 = 0
                                                                            var v133 int = 42
                                                                            v133 = 1
                                                                            _ = v133
                                                                        }
                                                                        v3 -= 42
                                                                    }
                                                                    delete(v4, v3)
                                                                }
                                                            }
                                                            v3 -= 47
                                                        }
                                                        v5 = (v125 | -49)
                                                        var v134 uint64 = 42
                                                        fmt.Println(v5, v134, v3)
                                                        delete(v2, (fi43 >> ((604) & 63)))
                                                        _ = v134
                                                    } else {
                                                        var v135 uint64 = 1
                                                        v135 = (v135 << ((65535) & 63))
                                                        v125 += 127
                                                        _ = v135
                                                    }
                                                    if ("foo bar" <= "") {
                                                        var v136 uint8 = 42
                                                        if (((-(v125)) > -48) && true) {
                                                            _ = len(v4)
                                                        }
                                                        v125 = (v5 ^ (v125 << ((520) & 63)))
                                                        v44 = v44
                                                        _ = v136
                                                    } else {
                                                        v125 += (-(v125))
                                                        v125 += v5
                                                        if false {
                                                            var v137 uint16 = 0
                                                            _ = v137
                                                        }
                                                        var v138 uint = 121
                                                        _ = len(v2)
                                                        _ = v138
                                                    }
                                                    v5 -= v125
                                                }
                                            case "a":
                                                v3 = (v3 >> ((681) & 63))
                                                if true {
                                                    var v139 int = (fi43 & fi45)
                                                    if (!(!false)) {
                                                        var v140 int = v139
                                                        _ = v140
                                                    }
                                                    _ = v139
                                                } else {
                                                    var v141 bool = false
                                                    _ = v141
                                                }
                                                fmt.Println(v5, v44, v44)
                                            case "":
                                                if false {
                                                    var v142 int = (fi43 - fi43)
                                                    v143 := GenS17{F0: true, F1: 0, F2: v125}
                                                    fmt.Println(v125)
                                                    v144 := GenS4{F0: (-(fi45)), F1: int32(v5), F2: 18446744073709551615}
                                                    if (128 <= 0) {
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            if true {
                                                                switch (fi43 ^ fi45) {
                                                                case -1:
                                                                    fmt.Println(v5)
                                                                case 32768:
                                                                    var v145 uint = 1
                                                                    if false {
                                                                        if (true || (1 >= v3)) {
                                                                            if true {
                                                                                var v146 bool = false
                                                                                v125 = v125
                                                                                v147 := GenS18{F0: -2147483648}
                                                                                var v148 bool = v146
                                                                                switch "" {
                                                                                case "a":
                                                                                    fmt.Println(v146, fi45, v142)
                                                                                }
                                                                                var v149 uint64 = 1
                                                                                _ = v146
                                                                                _ = v147
                                                                                _ = v148
                                                                                _ = v149
                                                                            }
                                                                            v125 -= 50
                                                                        } else {
                                                                            var v150 int = v142
                                                                            _ = v150
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            v151 := GenS19{F0: 32767, F1: 2147483647}
                                                                            _ = v151
                                                                        case "hello":
                                                                            var v152 uint32 = v3
                                                                            _ = v152
                                                                        case "foo bar":
                                                                            var v153 int8 = (v5 % ((v125) | 1))
                                                                            _ = v153
                                                                        default:
                                                                            if true {
                                                                                var v154 uint64 = 2
                                                                                var v155 float32 = -2.25
                                                                                var v156 uint16 = 2
                                                                                if (v154 >= v154) {
                                                                                    v44 = v44
                                                                                    v5 += 40
                                                                                    var v157 uint = 0
                                                                                    _ = v157
                                                                                } else {
                                                                                    v155 = 3.14159
                                                                                    v158 := GenS11{F0: 843, F1: v145, F2: v142}
                                                                                    _ = v158
                                                                                }
                                                                                _ = v154
                                                                                _ = v155
                                                                                _ = v156
                                                                            } else {
                                                                                v145 = (v145 >> ((42) & 63))
                                                                                var v159 int8 = v5
                                                                                var v160 int = int(v145)
                                                                                _ = v159
                                                                                _ = v160
                                                                            }
                                                                        }
                                                                        var v161 uint16 = 0
                                                                        var v162 uint = v145
                                                                        switch ("foo bar" + "foo bar") {
                                                                        case "a":
                                                                            fmt.Println(v5, v161, v44)
                                                                        case "\n\t":
                                                                            switch fi45 {
                                                                            case 0:
                                                                                v3 = 0
                                                                            case 42:
                                                                                switch v142 {
                                                                                case 658:
                                                                                    v44 = ((v44 * v44) - float64((1.5 + (0.1 * (((0.0 - ((float32(0.0) - (float32(((v44 + v44) * (v44 / 1e10))) + float32((1e10 * v44)))) - -2.25)) + 1e10) + 0.1)))))
                                                                                case 127:
                                                                                    v163 := GenS11{F0: 0, F1: v162, F2: v142}
                                                                                    fmt.Println(v3, v5, v125)
                                                                                    _ = v163
                                                                                case -9223372036854775808:
                                                                                    var v164 float64 = float64(1.5)
                                                                                    _ = v164
                                                                                default:
                                                                                    v161 = v161
                                                                                }
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = v3
                                                                                }
                                                                                v165 := GenS19{F0: 971, F1: 858}
                                                                                _ = v165
                                                                            case -1:
                                                                                if true {
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        v166 := GenS4{F0: (fi43 * (v142 ^ v142)), F1: 42, F2: 113}
                                                                                        _ = v166
                                                                                    case "a":
                                                                                        var v167 int32 = -2147483648
                                                                                        _ = v167
                                                                                    case "foo bar":
                                                                                        var v168 int32 = 42
                                                                                        _ = v168
                                                                                    default:
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            v162 = v162
                                                                                        case "hello":
                                                                                            var v169 int64 = -1
                                                                                            _ = v169
                                                                                        case "":
                                                                                            var v170 float32 = 3.14159
                                                                                            _ = v170
                                                                                        default:
                                                                                            v171 := GenS18{F0: int32(fi45)}
                                                                                            _ = v171
                                                                                        }
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if false {
                                                                                    var v172 float64 = v44
                                                                                    var v173 string = "foo bar"
                                                                                    _ = v172
                                                                                    _ = v173
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v161
                                                                        _ = v162
                                                                    }
                                                                    if ((3.14159 + ((0.1 - 1.5) * float32((v44 / v44)))) <= (0.0 - 0.1)) {
                                                                        if (true && true) {
                                                                            if (-32768 != 42) {
                                                                                switch ("a" + "a") {
                                                                                case "a":
                                                                                    var v174 uint = 730
                                                                                    var v175 int = (fi43 | v142)
                                                                                    var v176 uint8 = 128
                                                                                    _ = v174
                                                                                    _ = v175
                                                                                    _ = v176
                                                                                case "\n\t":
                                                                                    v142 = (v142 / ((fi43) | 1))
                                                                                }
                                                                                var v177 uint64 = 519
                                                                                _ = len(v1)
                                                                                _ = v177
                                                                            }
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            v145 = v145
                                                                        }
                                                                    } else {
                                                                        v145 = 255
                                                                    }
                                                                    _ = v145
                                                                case -300:
                                                                    v125 -= (v125 * -103)
                                                                }
                                                                if false {
                                                                    fmt.Println(v5, v44, v44)
                                                                    fmt.Println(fi45, v142, v5)
                                                                    _ = len(v2)
                                                                    fmt.Println(v44, v5, v142)
                                                                    v178 := GenS20{F0: uint64(v3), F1: -2147483648}
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 1
                                                                    }
                                                                    v3 -= 2147483647
                                                                    _ = v178
                                                                }
                                                                delete(v4, 229)
                                                            } else {
                                                                if false {
                                                                    _ = len(v1)
                                                                }
                                                            }
                                                        default:
                                                            fmt.Println(v44)
                                                            v179 := GenS21{F0: 255, F1: 1, F2: 0}
                                                            if (!(!(((v44 + float64(float32((v44 + v44)))) == (0.1 / 3.14159)) || false))) {
                                                                var v180 uint16 = 65535
                                                                v142 = (fi45 >> ((929) & 63))
                                                                _ = v180
                                                            }
                                                            _ = v179
                                                        }
                                                        fmt.Println(v3, v125, v44)
                                                    } else {
                                                        v44 = 3.14159
                                                        v5 -= (v125 ^ -1)
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v142
                                                    _ = v143
                                                    _ = v144
                                                }
                                            }
                                            v44 -= 1.5
                                            _ = v125
                                        default:
                                            fmt.Println(v3, v3)
                                        }
                                        v3 -= v3
                                        switch "a" {
                                        case "foo bar":
                                            fmt.Println(fi43, fi43, fi43)
                                        case "hello":
                                            var v181 uint = 4294967295
                                            _ = v181
                                        case "\n\t":
                                            var v182 float64 = v44
                                            v5 = -75
                                            var v183 uint16 = 702
                                            switch ("" + (("hello" + "\n\t") + ("\n\t" + "a"))) {
                                            case "foo bar":
                                                if true {
                                                    switch (("\n\t" + "a") + "foo bar") {
                                                    case "":
                                                        v184 := GenS22{F0: 0, F1: 127, F2: 32767}
                                                        _ = v184
                                                    case "a":
                                                        var v185 int64 = -9223372036854775808
                                                        _ = v185
                                                    default:
                                                        var v186 uint64 = 2
                                                        _ = v186
                                                    }
                                                } else {
                                                    v183 = (^(v183))
                                                }
                                            case "a":
                                                v44 = 0.0
                                            default:
                                                v187 := GenS13{F0: ((v5 << ((581) & 63)) > v5), F1: 2}
                                                _ = v187
                                            }
                                            _ = v182
                                            _ = v183
                                        }
                                        _ = v124
                                    } else {
                                        v188 := GenS13{F0: false, F1: 1}
                                        v2[fi43] = ((!(32767 < 1000)) && false)
                                        _ = v188
                                    }
                                    var v189 uint64 = 18446744073709551615
                                    var v190 uint = 0
                                    _ = v123
                                    _ = v189
                                    _ = v190
                                }
                                var v191 uint32 = (^(v3))
                                _ = v191
                            }
                            _ = fi45
                            _ = v46
                        }
                        _ = fi43
                        _ = v44
                    }
                default:
                    v3 = 2
                }
                var v192 uint = uint(v3)
                v4[(v3 * v3)] = v3
                _ = v192
            }
        }
        _ = v34
    }
    delete(v2, 128)
    switch (((("\n\t" + "foo bar") + "foo bar") + "foo bar") + "hello") {
    case "\n\t":
        if (false || false) {
            fmt.Println(v5)
            v193 := GenS23{F0: 0}
            _ = v193
        }
    case "":
        if (!true) {
            fmt.Println(v5, v5)
            v3 -= v3
            v3 = v3
        }
    }
    v3 -= 2
    v3 = (v3 | (^(v3)))
    for fi194 := 0; fi194 < 17; fi194++ {
        for fi195 := 0; fi195 < 6; fi195++ {
            v2[(fi194 + fi195)] = (v3 <= (^(v3)))
            v196 := v2[1]
            var v197 float64 = ((1.5 + 1e10) - (-2.25 - (1.5 - 1e10)))
            var v198 uint64 = 2
            _ = fi195
            _ = v196
            _ = v197
            _ = v198
        }
        _ = fi194
    }
    v3 = (^(v3))
    var v199 int32 = 42
    var v200 float32 = 0.1
    for fi201 := 0; fi201 < 16; fi201++ {
        v199 = 217
        fmt.Println(v5, v199)
        var v202 uint64 = 4294967295
        _ = fi201
        _ = v202
    }
    if false {
        fmt.Println(v3, v199, v200)
        var v203 uint16 = 2
        _ = v203
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v199:", v199)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v200:", v200)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v199
    _ = v2
    _ = v200
    _ = v3
    _ = v4
    _ = v5
}
