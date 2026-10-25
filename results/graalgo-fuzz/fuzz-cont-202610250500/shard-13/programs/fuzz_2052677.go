package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 int64; F1 uint16; F2 int32 }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int32; F1 int64 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint64; F1 uint32 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int32; F1 uint; F2 string }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 int }
type GenS21 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS22 struct { F0 string; F1 int16 }
type GenS23 struct { F0 string }
func genFn0(p0 int64, p1 uint8) uint32 {
    p1 = p1
    v5 := []int{-428, -1, -9223372036854775808, 32768}
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            p1 -= p1
            _ = len(v5)
            var v8 uint = 65535
            _ = fi7
            _ = v8
        }
        if true {
            if (2 < 916) {
                p1 -= 128
                _ = len(v5)
            }
        }
        if true {
            p0 = (p0 / ((1) | 1))
        } else {
            if false {
                v9 := GenS1{F0: 42, F1: 1, F2: 2147483648}
                if len(v5) > 0 {
                    v5[0] = -94
                }
                v10 := GenS1{F0: 2147483647, F1: 747, F2: 2147483647}
                _ = v10
                _ = v9
            } else {
                v11 := GenS2{F0: 42, F1: true}
                _ = v11
            }
            p1 = 0
        }
        _ = fi6
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1() int32 {
    if (1 > 883) {
        switch ("foo bar" + "") {
        case "foo bar":
            switch 32767 {
            case -36:
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 uint64 = 511
                    _ = fi12
                    _ = v13
                }
            case 32768:
                for fi14 := 0; fi14 < 9; fi14++ {
                    if (fi14 >= fi14) {
                        var v15 uint16 = 1
                        fmt.Println(v15)
                        _ = v15
                    } else {
                        v16 := []int{fi14}
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        _ = v16
                    }
                    var v17 string = "hello"
                    _ = fi14
                    _ = v17
                }
            case 32767:
                var v18 bool = ((3.14159 > 0.0) || false)
                _ = v18
            }
        case "\n\t":
            v19 := []int{4294967295, -857}
            _ = v19
        default:
            v20 := GenS3{F0: false, F1: 42}
            _ = v20
        }
        switch "" {
        case "":
            var v21 uint = 485
            _ = v21
        case "\n\t":
            var v22 uint = 4294967295
            _ = v22
        case "hello":
            v23 := GenS2{F0: 666, F1: false}
            _ = v23
        }
    } else {
        v24 := GenS4{F0: ("foo bar" != "foo bar"), F1: 2, F2: 1}
        var v25 string = ("" + "hello")
        v26 := []int{-9223372036854775808, 964, 128}
        _ = v24
        _ = v25
        _ = v26
    }
    if ((1 <= -377) || true) {
        var v27 float32 = (0.0 * -2.25)
        if ((!(!(0 >= 1))) || ("\n\t" > (("\n\t" + "foo bar") + ((("\n\t" + "foo bar") + "foo bar") + "hello")))) {
            v27 -= v27
            v28 := []int{1, 1, -2147483648}
            v27 += (1e10 - v27)
            _ = v28
        }
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 int8 = -102
            _ = fi29
            _ = v30
        }
        _ = v27
    }
    return 0
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{}
    v3 := []uint32{65535, 0, 0}
    v4 := []int{1}
    if (1e10 < 1e10) {
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 uint = 42
            _ = fi31
            _ = v32
        }
    }
    v33 := GenS5{F0: 0, F1: 1}
    var v34 int8 = -128
    fmt.Println(v34)
    var v35 string = "foo bar"
    var v36 int8 = (v34 + (v34 | v34))
    v37 := GenS4{F0: (-181 != -1), F1: 2, F2: 255}
    var v38 float64 = 0.0
    for fi39 := 0; fi39 < 7; fi39++ {
        for fi40 := 0; fi40 < 12; fi40++ {
            v4 = append(v4, fi40)
            _ = fi40
        }
        if false {
            fmt.Println(v38, v35, v38)
        }
        _ = fi39
    }
    if false {
        v36 += v36
    } else {
        v35 = "a"
        for fi41 := 0; fi41 < 14; fi41++ {
            if len(v2) > 0 {
                v2[0] = v35
            }
            if false {
                switch "foo bar" {
                case "hello":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "a":
                    v42 := GenS5{F0: 1, F1: v36}
                    _ = v42
                case "\n\t":
                    if (!(!(true && ((-(v36)) != v36)))) {
                        _ = len(v2)
                        v43 := GenS6{F0: (v34 / ((v36) | 1)), F1: 0}
                        _ = v43
                    } else {
                        _ = len(v3)
                    }
                }
                for fi44 := 0; fi44 < 7; fi44++ {
                    v45 := GenS5{F0: 0, F1: (v36 & (v36 / ((-91) | 1)))}
                    if false {
                        if (!(!false)) {
                            var v46 float64 = ((v38 - (((v38 + -2.25) - v38) - (float64(1.5) * v38))) / v38)
                            v47 := GenS7{F0: -9223372036854775808, F1: 65535, F2: -671}
                            switch (fi44 | fi41) {
                            case 4294967295:
                                v1[(fi44 % (((fi41 % (((fi41 & fi41)) | 1))) | 1))] = (v38 > v46)
                            case -1:
                                v34 -= v36
                            }
                            v48 := GenS5{F0: 42, F1: v36}
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                        v38 -= 1.5
                        v38 = (v38 / (0.1 / v38))
                    }
                    v35 = ("" + "a")
                    fmt.Println(v34, fi41, fi44)
                    if (42 != 4294967295) {
                        v34 = (v34 * v34)
                        v49 := GenS8{F0: 807, F1: true}
                        if (1000 <= -32768) {
                            var v50 float64 = (v38 * (float64(1.5) + -2.25))
                            _ = v50
                        }
                        _ = v49
                    } else {
                        switch v35 {
                        case "\n\t":
                            var v51 uint8 = 255
                            _ = v51
                        case "foo bar":
                            if ((((1.5 * float32(v38)) - (1.5 / (0.1 - (-2.25 - float32(((float64((1e10 * 0.0)) + (((-2.25 + 1.5) * v38) - v38)) + (v38 * 1e10))))))) * 1.5) > (1.5 * -2.25)) {
                                if ((1e10 - 0.1) <= ((1.5 + 1e10) / 1.0)) {
                                    if false {
                                        fmt.Println(v34, v35)
                                    } else {
                                        v34 = (v34 ^ v36)
                                        v52 := GenS9{F0: 0}
                                        fmt.Println(v36, fi41)
                                        _ = v52
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                v1[fi44] = true
                            } else {
                                if (true || false) {
                                    var v53 string = ((v35 + (((((((((("foo bar" + ((v35 + v35) + "foo bar")) + ((v35 + v35) + ("foo bar" + v35))) + "foo bar") + "") + "\n\t") + v35) + (("foo bar" + v35) + "\n\t")) + v35) + "") + v35)) + "\n\t")
                                    if ((!true) && false) {
                                        v35 = v53
                                    } else {
                                        switch (fi44 % (((^(fi44))) | 1)) {
                                        case 32768:
                                            if (true && (42 <= -952)) {
                                                var v54 uint32 = 4294967295
                                                v55 := GenS8{F0: 0, F1: true}
                                                fmt.Println(v34)
                                                _ = v54
                                                _ = v55
                                            } else {
                                                if ((false && (-2147483648 < -926)) && true) {
                                                    var v56 int16 = -460
                                                    _ = v56
                                                } else {
                                                    if (255 <= 443) {
                                                        v35 = ((v35 + ("\n\t" + v35)) + v53)
                                                    }
                                                    v34 += 42
                                                    var v57 uint32 = 65535
                                                    _ = v57
                                                }
                                            }
                                        case 2147483647:
                                            if (!true) {
                                                fmt.Println(v34, v38, v36)
                                                switch fi44 {
                                                case 32767:
                                                    v58 := GenS10{F0: -2147483648, F1: 815}
                                                    _ = v58
                                                case 849:
                                                    fmt.Println(v34)
                                                }
                                                var v59 int8 = v36
                                                _ = v59
                                            }
                                        default:
                                            v60 := GenS7{F0: -9223372036854775808, F1: 2, F2: 2147483647}
                                            _ = v60
                                        }
                                        var v61 bool = ((((v35 + v35) + (((v53 + v35) + "") + (v35 + v53))) + v53) < v53)
                                        var v62 int = 4294967295
                                        _ = v61
                                        _ = v62
                                    }
                                    _ = v53
                                }
                                if len(v2) > 0 {
                                    v2[0] = v35
                                }
                                fmt.Println(fi41, v34, v35)
                            }
                        }
                        fmt.Println(fi41, v35, v35)
                    }
                    if true {
                        switch fi41 {
                        case 993:
                            var v63 uint = 4294967295
                            _ = v63
                        case 127:
                            v38 = float64(3.14159)
                        case 42:
                            v34 = v36
                        }
                        switch (((v35 + "a") + ((("foo bar" + v35) + "\n\t") + "")) + (v35 + "")) {
                        case "a":
                            if true {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if (!false) {
                                    var v64 uint16 = 692
                                    if true {
                                        fmt.Println(fi41, v34)
                                        if (!(v36 < 106)) {
                                            if true {
                                                var v65 int64 = -9223372036854775808
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v65
                                            } else {
                                                if false {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        switch v35 {
                                                        case "foo bar":
                                                            v38 = 1e10
                                                        case "hello":
                                                            var v66 int64 = -1
                                                            _ = v66
                                                        default:
                                                            if ((true && (!true)) && (3.14159 == v38)) {
                                                                var v67 float32 = (float32(1.5) - (((float32(1e10) / (-2.25 * ((1.5 + float32((v38 / ((v38 + v38) * v38)))) - float32(v38)))) / (-2.25 - (0.0 * 0.0))) / 1.0))
                                                                if (0 != -11) {
                                                                    v68 := GenS4{F0: true, F1: 0, F2: 197}
                                                                    v64 += 42
                                                                    fmt.Println(v36, v35, fi44)
                                                                    _ = v68
                                                                } else {
                                                                    v69 := GenS4{F0: true, F1: 128, F2: 127}
                                                                    _ = v69
                                                                }
                                                                if false {
                                                                    fmt.Println(v34, v64)
                                                                    delete(v1, (-(fi41)))
                                                                }
                                                                _ = v67
                                                            } else {
                                                                var v70 int8 = v34
                                                                var v71 float32 = 3.14159
                                                                v2 = append(v2, (v35 + ("\n\t" + v35)))
                                                                _ = v70
                                                                _ = v71
                                                            }
                                                        }
                                                    case "":
                                                        fmt.Println(fi44, fi41)
                                                    }
                                                    fmt.Println(v35, fi41, v35)
                                                    v34 = (v36 ^ -1)
                                                } else {
                                                    if (40 <= 42) {
                                                        switch v35 {
                                                        case "":
                                                            v34 = v36
                                                        case "a":
                                                            fmt.Println(fi44, v36)
                                                        case "\n\t":
                                                            fmt.Println(fi44, v36, v34)
                                                            v64 += v64
                                                            var v72 uint16 = (^(v64))
                                                            _ = v72
                                                        }
                                                    }
                                                    v73 := GenS11{F0: false}
                                                    _ = v73
                                                }
                                                v36 -= 42
                                            }
                                            _ = len(v1)
                                        } else {
                                            fmt.Println(v35, v64)
                                            v36 += (v34 / ((v34) | 1))
                                            var v74 float32 = -2.25
                                            _ = v74
                                        }
                                        v64 = v64
                                    } else {
                                        if (false || true) {
                                            var v75 uint = 18446744073709551615
                                            delete(v1, (fi44 | fi44))
                                            if false {
                                                if ((1 == 329) || (true && (((v38 / v38) / (v38 + v38)) == 0.0))) {
                                                    fmt.Println(fi41, v64, v38)
                                                    v76 := GenS9{F0: v64}
                                                    var v77 float32 = 3.14159
                                                    fmt.Println(v34, v64)
                                                    var v78 uint32 = 4294967295
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                } else {
                                                    if (179 != 1) {
                                                        var v79 float32 = 3.14159
                                                        fmt.Println(v79)
                                                        _ = v79
                                                    } else {
                                                        var v80 string = v35
                                                        if (false || false) {
                                                            fmt.Println(v36, v36)
                                                            v1[(fi41 | fi44)] = true
                                                            v81 := GenS12{F0: v80}
                                                            if len(v3) > 0 {
                                                                v3[0] = 2147483648
                                                            }
                                                            _ = v81
                                                        } else {
                                                            v3 = append(v3, 2)
                                                            if (1 <= 255) {
                                                                v82 := GenS4{F0: true, F1: 255, F2: 255}
                                                                fmt.Println(fi41)
                                                                v83 := v1[0]
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                            switch (^(fi44)) {
                                                            case -350:
                                                                var v84 bool = (-133 >= genFn1())
                                                                _ = v84
                                                            case 2147483647:
                                                                v75 += (v75 & v75)
                                                            }
                                                        }
                                                        v34 -= (v34 | v36)
                                                        _ = v80
                                                    }
                                                    var v85 uint64 = 42
                                                    _ = v85
                                                }
                                                var v86 float32 = float32(float64(-2.25))
                                                if len(v4) > 0 {
                                                    v4[0] = 9223372036854775807
                                                }
                                                v34 += int8(v75)
                                                _ = v86
                                            } else {
                                                if true {
                                                    switch (fi44 << ((v75) & 63)) {
                                                    case 1:
                                                        var v87 int64 = 9223372036854775807
                                                        _ = v87
                                                    }
                                                    v75 -= (v75 << ((v75) & 63))
                                                } else {
                                                    if true {
                                                        v88 := GenS13{F0: 42, F1: 4294967295}
                                                        switch "a" {
                                                        case "foo bar":
                                                            v35 = ("a" + v35)
                                                        case "\n\t":
                                                            v89 := GenS14{F0: 2}
                                                            _ = v89
                                                        }
                                                        _ = v88
                                                    } else {
                                                        var v90 uint16 = 65535
                                                        _ = v90
                                                    }
                                                    switch 658 {
                                                    case 9223372036854775807:
                                                        var v91 float64 = v38
                                                        _ = v91
                                                    }
                                                    v64 -= v64
                                                    var v92 float32 = 0.0
                                                    var v93 string = v35
                                                    _ = v92
                                                    _ = v93
                                                }
                                                var v94 uint8 = uint8(v64)
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v94
                                            }
                                            var v95 int32 = 809
                                            _ = len(v1)
                                            v1[-239] = true
                                            _ = v75
                                            _ = v95
                                        } else {
                                            v96 := GenS14{F0: 350}
                                            if (!true) {
                                                v97 := GenS15{F0: -554, F1: 65535, F2: ("foo bar" + v35)}
                                                _ = v97
                                            }
                                            _ = v96
                                        }
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = (fi44 >> ((879) & 63))
                                    }
                                    _ = v64
                                } else {
                                    v35 = "\n\t"
                                    v2 = append(v2, v35)
                                    v1[988] = true
                                }
                                v35 = "foo bar"
                            }
                        case "\n\t":
                            v36 = 87
                        case "hello":
                            var v98 bool = (65535 != 65535)
                            _ = v98
                        default:
                            var v99 int64 = -902
                            _ = v99
                        }
                    } else {
                        if false {
                            var v100 int16 = -503
                            v101 := GenS10{F0: -784, F1: -9223372036854775808}
                            if (!((!false) && (fi41 <= (fi41 + (fi41 % ((fi41) | 1)))))) {
                                var v102 int64 = -772
                                var v103 int = 32768
                                var v104 bool = (v38 <= v38)
                                v104 = true
                                _ = v102
                                _ = v103
                                _ = v104
                            }
                            _ = v100
                            _ = v101
                        }
                        if ((genFn0(-993, 222) != 2) && (!false)) {
                            v35 = "hello"
                            v105 := v1[fi44]
                            _ = v105
                        } else {
                            var v106 int16 = -32768
                            _ = v106
                        }
                        v107 := GenS16{F0: 2}
                        _ = v107
                    }
                    if (v35 != v35) {
                        v34 += (v34 | v36)
                    } else {
                        v108 := GenS17{F0: false}
                        v109 := GenS10{F0: -1, F1: 1}
                        v34 = (v34 << ((2) & 63))
                        _ = v108
                        _ = v109
                    }
                    var v110 uint = 65535
                    v110 = v110
                    var v111 string = (v35 + (v35 + "hello"))
                    v111 = ("foo bar" + ("hello" + v111))
                    _ = fi44
                    _ = v110
                    _ = v111
                    _ = v45
                }
                v35 = v35
            } else {
                var v112 int64 = -1
                _ = v112
            }
            _ = fi41
        }
    }
    for fi113 := 0; fi113 < 17; fi113++ {
        if true {
            v38 = 1e10
            if true {
                fmt.Println(v35)
                v34 -= (v34 ^ v36)
            } else {
                switch 42 {
                case 32768:
                    v114 := GenS8{F0: 4294967295, F1: true}
                    _ = v114
                default:
                    var v115 uint16 = 65535
                    _ = v115
                }
                for fi116 := 0; fi116 < 10; fi116++ {
                    v36 = 39
                    _ = fi116
                }
            }
            v34 += 1
        }
        if false {
            v117 := GenS18{F0: 80, F1: 913, F2: 181}
            if (!((!true) || (18446744073709551615 == 18446744073709551615))) {
                v1[128] = (42 < 42)
                v36 -= v34
            }
            switch fi113 {
            case -415:
                switch (^(fi113)) {
                case 9223372036854775807:
                    fmt.Println(v38)
                case -1:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            case 193:
                if false {
                    var v118 uint = 2
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v118
                } else {
                    var v119 bool = (((3.14159 / 0.1) / (1e10 + 1.5)) <= 1e10)
                    _ = v119
                }
            default:
                if false {
                    delete(v1, (fi113 & fi113))
                    var v120 uint32 = 784
                    _ = v120
                }
            }
            _ = v117
        } else {
            for fi121 := 0; fi121 < 20; fi121++ {
                fmt.Println(v34)
                _ = fi121
            }
            v35 = "\n\t"
        }
        _ = fi113
    }
    for fi122 := 0; fi122 < 15; fi122++ {
        v3 = append(v3, 850)
        if (0 == 32767) {
            v34 = v36
        }
        _ = fi122
    }
    var v123 int64 = 214
    for fi124 := 0; fi124 < 7; fi124++ {
        if false {
            fmt.Println(v38)
            var v125 float64 = 3.14159
            v38 = v125
            _ = v125
        } else {
            v4 = append(v4, -2147483648)
            var v126 bool = false
            if v126 {
                if false {
                    for fi127 := 0; fi127 < 2; fi127++ {
                        if v126 {
                            var v128 uint64 = 133
                            switch "a" {
                            case "\n\t":
                                if false {
                                    if (42 > 572) {
                                        if v126 {
                                            if (1 > -1) {
                                                if v126 {
                                                    v3 = append(v3, 0)
                                                    v129 := GenS19{F0: (v128 - v128)}
                                                    var v130 int16 = 0
                                                    _ = v129
                                                    _ = v130
                                                }
                                            }
                                            v131 := GenS4{F0: v126, F1: 128, F2: 255}
                                            var v132 int32 = genFn1()
                                            _ = len(v1)
                                            _ = v131
                                            _ = v132
                                        }
                                        if (0.1 >= (0.0 - ((1.5 / 1.0) + float32((1e10 * (v38 - v38)))))) {
                                            _ = len(v4)
                                        }
                                        v133 := v1[(fi127 >> ((18446744073709551615) & 63))]
                                        _ = v133
                                    }
                                } else {
                                    _ = len(v1)
                                    var v134 uint32 = 1
                                    var v135 int32 = -1
                                    _ = v134
                                    _ = v135
                                }
                            case "hello":
                                var v136 int = fi127
                                _ = v136
                            default:
                                var v137 int32 = -2147483648
                                _ = v137
                            }
                            _ = v128
                        } else {
                            switch "\n\t" {
                            case "":
                                switch 1 {
                                case 971:
                                    switch (^(fi124)) {
                                    case 127:
                                        var v138 int64 = -9223372036854775808
                                        _ = v138
                                    case 443:
                                        fmt.Println(v36, fi124)
                                    case 2147483647:
                                        delete(v1, fi124)
                                    default:
                                        var v139 float64 = (v38 * float64(0.0))
                                        _ = v139
                                    }
                                case 51:
                                    v140 := GenS20{F0: (fi124 * (^(fi124)))}
                                    _ = v140
                                case 0:
                                    if v126 {
                                        v38 = (v38 * v38)
                                        var v141 float64 = v38
                                        switch ("" + v35) {
                                        case "":
                                            v38 = float64(-2.25)
                                        case "a":
                                            fmt.Println(v126, v123)
                                        case "hello":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        v38 += v141
                                        _ = v141
                                    }
                                default:
                                    v142 := GenS13{F0: 238, F1: 1}
                                    _ = v142
                                }
                            case "a":
                                var v143 int64 = v123
                                _ = v143
                            default:
                                fmt.Println(v123)
                            }
                            switch -708 {
                            case -705:
                                v2 = append(v2, "a")
                            case 403:
                                fmt.Println(v126, v36, v36)
                            }
                            if ((v123 * -1) != v123) {
                                delete(v1, (fi127 << ((0) & 63)))
                                switch fi127 {
                                case 352:
                                    if (175 <= 76) {
                                        v144 := GenS3{F0: v126, F1: 0}
                                        v145 := GenS21{F0: 2, F1: (^(v123)), F2: 2147483648}
                                        var v146 string = ""
                                        _ = v144
                                        _ = v145
                                        _ = v146
                                    }
                                }
                            }
                        }
                        _ = fi127
                    }
                    for fi147 := 0; fi147 < 4; fi147++ {
                        v36 -= v36
                        v148 := GenS22{F0: (v35 + (v35 + v35)), F1: 367}
                        var v149 uint = 643
                        _ = fi147
                        _ = v148
                        _ = v149
                    }
                    v38 += v38
                } else {
                    if (!v126) {
                        fmt.Println(v35, v38, v35)
                        var v150 int64 = (v123 | 1)
                        _ = v150
                    }
                }
                v151 := GenS23{F0: v35}
                _ = v151
            }
            _ = v126
        }
        _ = fi124
    }
    v152 := v1[83]
    _ = len(v1)
    var v153 int = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v123:", v123)
    fmt.Println("v152:", v152)
    fmt.Println("v153:", v153)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v123
    _ = v152
    _ = v153
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
