package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 uint8; F1 int32; F2 int64 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string; F1 int32; F2 int64 }
type GenS10 struct { F0 bool; F1 int8; F2 int64 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int8; F1 int8; F2 uint8 }
type GenS13 struct { F0 uint8; F1 uint64; F2 int16 }
func genFn0(p0 uint16, p1 int32, p2 string) int32 {
    var v5 int = 895
    v6 := []int{v5, (v5 | int(p0)), (^(v5)), v5}
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int64 = 591
        var v9 uint8 = 2
        var v10 int32 = (p1 / ((p1) | 1))
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (^(p1))
}
func genFn1() int16 {
    var v11 int64 = 1
    v11 = -9223372036854775808
    var v12 uint8 = 0
    v13 := []int{}
    _ = v11
    _ = v12
    _ = v13
    return -40
}

func main() {
    v1 := []uint32{42}
    var v2 float64 = 1.5
    v3 := make(map[uint32]string)
    var v4 bool = (255 >= 247)
    v14 := GenS1{F0: -115}
    v2 = 3.14159
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 bool = (false && ((0.0 - ((1e10 * (1.5 / float32(-2.25))) * 1.5)) == float32(v2)))
        v17 := GenS1{F0: (fi15 >> ((255) & 63))}
        _ = fi15
        _ = v16
        _ = v17
    }
    v2 -= (v2 - 0.0)
    if (!(52 > 65535)) {
        fmt.Println(v2, v4)
        v1 = append(v1, 474)
        v2 = float64(((float32((v2 + v2)) * float32((v2 + (v2 - v2)))) - 3.14159))
    } else {
        switch -2147483648 {
        case -2147483648:
            if len(v1) > 0 {
                v1[0] = 99
            }
        case 950:
            switch 2147483647 {
            case 42:
                if v4 {
                    if (699 >= 249) {
                        for fi18 := 0; fi18 < 6; fi18++ {
                            v3[393] = "\n\t"
                            v4 = v4
                            _ = fi18
                        }
                    }
                } else {
                    if v4 {
                        fmt.Println(v2, v4)
                        if (-1 <= 163) {
                            v2 = v2
                            v19 := GenS2{F0: -172}
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 16; fi20++ {
                                v4 = (!(!v4))
                                if (!(100 >= 1)) {
                                    switch fi20 {
                                    case 4294967295:
                                        for fi21 := 0; fi21 < 14; fi21++ {
                                            v4 = v4
                                            var v22 uint8 = 0
                                            _ = fi21
                                            _ = v22
                                        }
                                    case -9223372036854775808:
                                        var v23 string = (("a" + "") + ("\n\t" + "foo bar"))
                                        _ = v23
                                    default:
                                        var v24 int32 = 0
                                        _ = v24
                                    }
                                    v2 = v2
                                }
                                _ = fi20
                            }
                            v2 = v2
                        }
                    }
                    v25 := GenS2{F0: 851}
                    var v26 int8 = 42
                    _ = v25
                    _ = v26
                }
            default:
                switch ((("" + "") + ("\n\t" + (("" + ("hello" + ("a" + ("" + ("foo bar" + ("foo bar" + "\n\t")))))) + "a"))) + "hello") {
                case "\n\t":
                    if (!(127 >= 583)) {
                        if (!(2147483647 != -129)) {
                            var v27 uint32 = 42
                            var v28 int8 = 42
                            _ = v27
                            _ = v28
                        }
                    } else {
                        _ = len(v1)
                        switch -2147483648 {
                        case 127:
                            _ = len(v3)
                        }
                        for fi29 := 0; fi29 < 4; fi29++ {
                            for fi30 := 0; fi30 < 9; fi30++ {
                                fmt.Println(v2)
                                _ = fi30
                            }
                            v4 = false
                            var v31 int16 = 42
                            _ = fi29
                            _ = v31
                        }
                    }
                }
            }
        default:
            v2 -= (v2 * v2)
        }
    }
    v2 = (v2 - float64(1e10))
    fmt.Println(v4, v4)
    var v32 bool = true
    for fi33 := 0; fi33 < 6; fi33++ {
        v32 = (1 < -32768)
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 uint32 = 211
            var v36 uint8 = 1
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi33
    }
    var v37 float32 = (3.14159 * 0.0)
    var v38 uint64 = 42
    switch (("\n\t" + ("foo bar" + "")) + "hello") {
    case "a":
        switch "a" {
        case "\n\t":
            for fi39 := 0; fi39 < 2; fi39++ {
                fmt.Println(v38, v38)
                _ = fi39
            }
        case "foo bar":
            _ = len(v3)
        case "a":
            _ = len(v1)
        default:
            var v40 uint = 18446744073709551615
            _ = v40
        }
    case "\n\t":
        v41 := GenS3{F0: 57, F1: 42}
        _ = v41
    case "hello":
        v32 = (v2 != float64(v37))
    default:
        var v42 bool = false
        _ = v42
    }
    if v4 {
        var v43 string = "a"
        if (983 != 32768) {
            if true {
                for fi44 := 0; fi44 < 19; fi44++ {
                    switch fi44 {
                    case -435:
                        var v45 int64 = -955
                        _ = v45
                    case -207:
                        switch 1 {
                        case 128:
                            if (!(false || (0 == -1))) {
                                v46 := GenS2{F0: (fi44 / (((fi44 + (fi44 % ((fi44) | 1)))) | 1))}
                                v47 := GenS3{F0: -91, F1: 461}
                                _ = v46
                                _ = v47
                            } else {
                                fmt.Println(v43, v43, v2)
                                if v4 {
                                    if v32 {
                                        v2 -= ((0.0 * v2) * 1e10)
                                        v37 += float32((1.5 / 1.0))
                                        switch fi44 {
                                        case 32767:
                                            _ = len(v3)
                                        default:
                                            if (false && v32) {
                                                for fi48 := 0; fi48 < 14; fi48++ {
                                                    v3[32] = "\n\t"
                                                    v1 = append(v1, 2147483648)
                                                    _ = fi48
                                                }
                                                switch v43 {
                                                case "a":
                                                    switch 42 {
                                                    case -880:
                                                        if ("foo bar" >= v43) {
                                                            for fi49 := 0; fi49 < 8; fi49++ {
                                                                v1 = append(v1, 2147483647)
                                                                if (!false) {
                                                                    var v50 uint16 = 300
                                                                    _ = v50
                                                                } else {
                                                                    v51 := GenS1{F0: (fi44 / ((475) | 1))}
                                                                    _ = v51
                                                                }
                                                                delete(v3, 731)
                                                                switch ((((v43 + v43) + ((("hello" + v43) + (v43 + v43)) + v43)) + "a") + v43) {
                                                                case "a":
                                                                    v32 = false
                                                                case "hello":
                                                                    if v32 {
                                                                        v52 := GenS4{F0: 0, F1: 1, F2: -843}
                                                                        _ = len(v1)
                                                                        v38 -= (v38 * (v38 - v38))
                                                                        if ((v38 % ((v38) | 1)) < (v38 & 548)) {
                                                                            _ = len(v3)
                                                                        } else {
                                                                            var v53 uint8 = 49
                                                                            switch v43 {
                                                                            case "foo bar":
                                                                                v54 := GenS5{F0: 2147483647, F1: 2147483648}
                                                                                _ = v54
                                                                            case "":
                                                                                var v55 int = fi44
                                                                                _ = v55
                                                                            case "hello":
                                                                                delete(v3, 0)
                                                                            default:
                                                                                var v56 int8 = 100
                                                                                _ = v56
                                                                            }
                                                                            switch (("\n\t" + ((v43 + (("" + ("a" + "")) + ("foo bar" + (("" + v43) + "\n\t")))) + "hello")) + v43) {
                                                                            case "":
                                                                                _ = len(v3)
                                                                            case "hello":
                                                                                if v4 {
                                                                                    if v32 {
                                                                                        var v57 uint32 = 2147483648
                                                                                        _ = v57
                                                                                    }
                                                                                    if v4 {
                                                                                        var v58 int64 = -419
                                                                                        if (!v32) {
                                                                                            v59 := GenS6{F0: v58, F1: v38}
                                                                                            if false {
                                                                                                v2 = v2
                                                                                                v60 := GenS4{F0: (v53 * v53), F1: 1, F2: (v58 + v58)}
                                                                                                _ = v60
                                                                                            }
                                                                                            var v61 string = (v43 + "foo bar")
                                                                                            _ = v59
                                                                                            _ = v61
                                                                                        }
                                                                                        v38 -= 1
                                                                                        _ = v58
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v53
                                                                        }
                                                                        _ = v52
                                                                    }
                                                                case "":
                                                                    v4 = (0 == 42)
                                                                default:
                                                                    if v32 {
                                                                        if (56 > 127) {
                                                                            if v4 {
                                                                                v2 = 0.1
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v62 := GenS7{F0: 42, F1: 1}
                                                                        v32 = false
                                                                        if v4 {
                                                                            if v32 {
                                                                                var v63 uint16 = 65535
                                                                                if ((v43 >= ("" + (((v43 + "") + v43) + v43))) && true) {
                                                                                    delete(v3, 42)
                                                                                } else {
                                                                                    _ = len(v3)
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        var v64 float32 = float32(v2)
                                                                                        _ = v64
                                                                                    default:
                                                                                        v32 = (-1 <= 0)
                                                                                    }
                                                                                    var v65 int16 = 42
                                                                                    v43 = v43
                                                                                    v66 := GenS3{F0: 116, F1: 0}
                                                                                    _ = v65
                                                                                    _ = v66
                                                                                }
                                                                                _ = v63
                                                                            }
                                                                            v67 := GenS8{F0: 2}
                                                                            _ = v67
                                                                        } else {
                                                                            v37 = float32(float64(v37))
                                                                            var v68 uint = 255
                                                                            var v69 int8 = 100
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                        _ = v62
                                                                    }
                                                                }
                                                                var v70 float32 = float32(v2)
                                                                _ = fi49
                                                                _ = v70
                                                            }
                                                            v43 = v43
                                                        }
                                                    default:
                                                        if (2147483648 >= 278) {
                                                            v32 = v32
                                                        } else {
                                                            var v71 int32 = 299
                                                            v38 += (v38 - (^(v38)))
                                                            _ = v71
                                                        }
                                                    }
                                                case "foo bar":
                                                    var v72 int8 = 100
                                                    _ = v72
                                                }
                                                v43 = v43
                                            } else {
                                                if (2 > 7) {
                                                    v4 = (v32 && true)
                                                    var v73 float32 = v37
                                                    v43 = "foo bar"
                                                    _ = v73
                                                }
                                                if false {
                                                    for fi74 := 0; fi74 < 19; fi74++ {
                                                        var v75 int = fi44
                                                        if (127 != 42) {
                                                            var v76 uint8 = 128
                                                            _ = v76
                                                        }
                                                        var v77 int8 = 127
                                                        _ = fi74
                                                        _ = v75
                                                        _ = v77
                                                    }
                                                    var v78 uint = 65535
                                                    _ = v78
                                                }
                                                var v79 int16 = 1
                                                _ = v79
                                            }
                                        }
                                    } else {
                                        var v80 uint = 2
                                        _ = v80
                                    }
                                    var v81 int16 = 32767
                                    _ = v81
                                } else {
                                    v38 = v38
                                    v82 := GenS9{F0: v43, F1: 2147483647, F2: 677}
                                    switch v43 {
                                    case "a":
                                        var v83 float64 = 1e10
                                        _ = v83
                                    case "hello":
                                        delete(v3, 42)
                                    case "\n\t":
                                        var v84 int64 = -1
                                        _ = v84
                                    }
                                    _ = v82
                                }
                            }
                        }
                    default:
                        v85 := GenS7{F0: 42, F1: uint32(v38)}
                        _ = v85
                    }
                    for fi86 := 0; fi86 < 16; fi86++ {
                        if v4 {
                            v32 = v4
                            v87 := GenS1{F0: 0}
                            _ = v87
                        } else {
                            switch v43 {
                            case "foo bar":
                                v43 = ""
                            default:
                                if true {
                                    v1 = append(v1, 780)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if v4 {
                                        if (!((65535 <= 888) && false)) {
                                            switch v43 {
                                            case "foo bar":
                                                v37 -= float32((v2 * v2))
                                            case "":
                                                var v88 uint64 = (v38 - 1)
                                                _ = v88
                                            }
                                            _ = len(v1)
                                            v38 = (^(v38))
                                        }
                                        var v89 string = v43
                                        _ = v89
                                    } else {
                                        switch "hello" {
                                        case "hello":
                                            v2 -= -2.25
                                        case "\n\t":
                                            if (1 < 0) {
                                                if (385 <= 232) {
                                                    v90 := GenS10{F0: v4, F1: -128, F2: 9223372036854775807}
                                                    if false {
                                                        if true {
                                                            v2 = 0.1
                                                            if (!((true || v32) && v32)) {
                                                                if true {
                                                                    if v32 {
                                                                        v91 := GenS5{F0: -1, F1: 42}
                                                                        fmt.Println(v2, v38, v38)
                                                                        if v4 {
                                                                            v37 += float32(0.0)
                                                                        }
                                                                        _ = v91
                                                                    }
                                                                    v32 = (("\n\t" + (("hello" + v43) + v43)) < (v43 + "a"))
                                                                    if true {
                                                                        if v32 {
                                                                            var v92 uint = 42
                                                                            v93 := GenS6{F0: 800, F1: (v38 % (((v38 - 2)) | 1))}
                                                                            v43 = v43
                                                                            _ = v92
                                                                            _ = v93
                                                                        }
                                                                        v94 := GenS11{F0: 255}
                                                                        var v95 uint16 = 42
                                                                        _ = v94
                                                                        _ = v95
                                                                    }
                                                                } else {
                                                                    if (!(true || ((!false) && v32))) {
                                                                        delete(v3, 4294967295)
                                                                        switch (fi86 & fi86) {
                                                                        case 4294967295:
                                                                            v32 = ((fi86 << ((4294967295) & 63)) >= (fi44 * fi44))
                                                                        case -316:
                                                                            var v96 float64 = 0.0
                                                                            _ = v96
                                                                        }
                                                                        var v97 uint32 = 1
                                                                        _ = v97
                                                                    } else {
                                                                        var v98 int = 128
                                                                        v4 = v32
                                                                        v2 = 0.1
                                                                        switch 32768 {
                                                                        case 2147483647:
                                                                            var v99 uint64 = (v38 - v38)
                                                                            _ = v99
                                                                        case 42:
                                                                            delete(v3, 843)
                                                                        case 9223372036854775807:
                                                                            if (-1 < 42) {
                                                                                v32 = false
                                                                                v32 = (4294967295 > 465)
                                                                                var v100 uint8 = 2
                                                                                var v101 int16 = 0
                                                                                switch -1 {
                                                                                case 2147483647:
                                                                                    switch 464 {
                                                                                    case 4294967295:
                                                                                        v102 := GenS1{F0: v98}
                                                                                        _ = v102
                                                                                    case 127:
                                                                                        if (1 < 0) {
                                                                                            v2 -= v2
                                                                                            fmt.Println(v2)
                                                                                            switch v43 {
                                                                                            case "a":
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            case "foo bar":
                                                                                                var v103 string = (("\n\t" + v43) + "")
                                                                                                _ = v103
                                                                                            default:
                                                                                                switch fi44 {
                                                                                                case 2147483647:
                                                                                                    if (32767 < 0) {
                                                                                                        switch 32768 {
                                                                                                        case 32768:
                                                                                                            var v104 float32 = v37
                                                                                                            _ = v104
                                                                                                        }
                                                                                                    }
                                                                                                case -2147483648:
                                                                                                    v101 += v101
                                                                                                default:
                                                                                                    var v105 float32 = v37
                                                                                                    _ = v105
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case 127:
                                                                                    fmt.Println(v38)
                                                                                case 1:
                                                                                    fmt.Println(v2)
                                                                                }
                                                                                _ = v100
                                                                                _ = v101
                                                                            }
                                                                        }
                                                                        v106 := GenS11{F0: 531}
                                                                        _ = v106
                                                                        _ = v98
                                                                    }
                                                                    var v107 bool = ((107 <= 42) || (1 <= 255))
                                                                    var v108 int = (-(fi86))
                                                                    _ = v107
                                                                    _ = v108
                                                                }
                                                            }
                                                            if ((1 == 42) && v32) {
                                                                if false {
                                                                    v2 = v2
                                                                    var v109 uint32 = 2
                                                                    _ = v109
                                                                } else {
                                                                    v4 = v32
                                                                    var v110 int64 = -844
                                                                    if true {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 2
                                                                        }
                                                                        var v111 string = v43
                                                                        _ = v111
                                                                    }
                                                                    _ = v110
                                                                }
                                                                if v32 {
                                                                    _ = len(v3)
                                                                } else {
                                                                    var v112 float32 = v37
                                                                    _ = v112
                                                                }
                                                                var v113 int64 = -1
                                                                fmt.Println(v37, fi44, v43)
                                                                v114 := v3[2]
                                                                var v115 uint16 = 152
                                                                v115 = (v115 >> ((0) & 63))
                                                                var v116 float32 = 3.14159
                                                                _ = v113
                                                                _ = v114
                                                                _ = v115
                                                                _ = v116
                                                            }
                                                        } else {
                                                            if v4 {
                                                                fmt.Println(v4, v38, v4)
                                                                var v117 float64 = v2
                                                                _ = v117
                                                            }
                                                            switch v43 {
                                                            case "foo bar":
                                                                var v118 int8 = -128
                                                                _ = v118
                                                            default:
                                                                v119 := GenS5{F0: -2147483648, F1: 2147483648}
                                                                _ = v119
                                                            }
                                                            _ = len(v3)
                                                            v120 := GenS6{F0: 0, F1: v38}
                                                            _ = v120
                                                        }
                                                        var v121 int16 = 651
                                                        _ = v121
                                                    } else {
                                                        var v122 uint16 = 2
                                                        _ = v122
                                                    }
                                                    var v123 int8 = 1
                                                    _ = v123
                                                    _ = v90
                                                }
                                            } else {
                                                v38 += 0
                                            }
                                        }
                                    }
                                }
                            }
                            v1 = append(v1, 0)
                        }
                        fmt.Println(v32, fi86, v32)
                        var v124 float32 = 0.1
                        _ = fi86
                        _ = v124
                    }
                    for fi125 := 0; fi125 < 3; fi125++ {
                        if v32 {
                            v37 -= v37
                            _ = len(v3)
                            switch "a" {
                            case "foo bar":
                                var v126 int64 = 42
                                _ = v126
                            default:
                                v4 = v32
                            }
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch v43 {
                            case "\n\t":
                                var v127 int16 = 42
                                _ = v127
                            case "hello":
                                v128 := GenS12{F0: -82, F1: 1, F2: 255}
                                _ = v128
                            default:
                                var v129 float64 = v2
                                _ = v129
                            }
                            fmt.Println(v37)
                        }
                        _ = fi125
                    }
                    _ = fi44
                }
            }
        }
        _ = v43
    } else {
        var v130 bool = v32
        switch 664 {
        case -9223372036854775808:
            if v32 {
                var v131 uint16 = 472
                v2 -= (v2 - v2)
                if v32 {
                    v4 = (false && (128 < 254))
                }
                _ = v131
            } else {
                v132 := v3[42]
                _ = v132
            }
        }
        _ = v130
    }
    for fi133 := 0; fi133 < 17; fi133++ {
        v134 := GenS7{F0: 307, F1: 0}
        v135 := GenS13{F0: 42, F1: v38, F2: 0}
        _ = fi133
        _ = v134
        _ = v135
    }
    v38 += 1
    v32 = v32
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v38
    _ = v4
}
