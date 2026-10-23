package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 bool }
type GenS2 struct { F0 uint; F1 string; F2 int16 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int64; F1 uint16 }
type GenS11 struct { F0 uint16; F1 uint8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 string; F1 int64; F2 string }
func genFn0(p0 uint8, p1 uint64, p2 int32) uint32 {
    p1 = (p1 / ((83) | 1))
    switch 968 {
    case 128:
        v5 := GenS1{F0: 0, F1: (((false && false) || (0.0 > (-2.25 * 1e10))) || false), F2: false}
        _ = v5
    }
    fmt.Println(p2, p0, p1)
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 bool, p1 uint8, p2 uint32) int16 {
    p0 = p0
    if p0 {
        fmt.Println(p1, p1, p2)
        for fi6 := 0; fi6 < 5; fi6++ {
            p2 = p2
            _ = fi6
        }
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint16 = 719
            var v9 uint16 = 294
            if (((0.1 + 0.0) - 1e10) > 0.1) {
                if p0 {
                    if true {
                        var v10 int64 = -697
                        v11 := GenS1{F0: 1, F1: false, F2: (48 <= 42)}
                        for fi12 := 0; fi12 < 4; fi12++ {
                            var v13 uint8 = p1
                            _ = fi12
                            _ = v13
                        }
                        _ = v10
                        _ = v11
                    }
                    var v14 uint8 = p1
                    v9 -= v8
                    _ = v14
                }
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS2{F0: uint(p1), F1: "hello", F2: 32767}
                    _ = fi15
                    _ = v16
                }
            } else {
                if (18446744073709551615 < 1) {
                    p0 = ((0 > -110) || false)
                } else {
                    fmt.Println(p2, v9, v8)
                }
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 float32 = ((1.5 / ((((-2.25 - -2.25) - (3.14159 - -2.25)) - 1e10) - (float32(0.1) / (1.5 / 1e10)))) / 1e10)
                    var v19 uint64 = uint64(p2)
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        fmt.Println(p2)
        switch "" {
        case "hello":
            for fi20 := 0; fi20 < 14; fi20++ {
                v21 := GenS1{F0: uint64(p2), F1: false, F2: p0}
                v22 := GenS3{F0: 210, F1: (float32(0.1) >= 1.5)}
                for fi23 := 0; fi23 < 12; fi23++ {
                    if p0 {
                        if false {
                            var v24 bool = (!((0 == 1000) && (!("hello" < "hello"))))
                            if true {
                                v24 = true
                                switch (fi20 % ((42) | 1)) {
                                case 2147483647:
                                    switch "\n\t" {
                                    case "a":
                                        var v25 int16 = 32767
                                        _ = v25
                                    default:
                                        v24 = false
                                    }
                                }
                                v26 := []int{128}
                                _ = v26
                            } else {
                                fmt.Println(v24, p0)
                                v27 := GenS2{F0: 162, F1: "hello", F2: -31}
                                _ = v27
                            }
                            p2 -= p2
                            _ = v24
                        }
                        var v28 bool = (65535 >= 32767)
                        _ = v28
                    }
                    var v29 string = (("" + "a") + "hello")
                    switch "a" {
                    case "\n\t":
                        var v30 int32 = 0
                        _ = v30
                    case "foo bar":
                        v31 := []int{fi23, fi23}
                        _ = v31
                    case "hello":
                        p0 = p0
                    }
                    _ = fi23
                    _ = v29
                }
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        var v32 int16 = -285
        _ = v32
    }
    p0 = false
    fmt.Println(p1, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    return -32768
}

func main() {
    var v1 float32 = 0.0
    var v2 float64 = 1.5
    var v3 int64 = 0
    v4 := []int{32767, 0, 364}
    var v33 float64 = v2
    fmt.Println(v3, v1)
    var v34 bool = true
    var v35 int64 = (v3 + (v3 & v3))
    if v34 {
        var v36 uint16 = 42
        var v37 int64 = v3
        var v38 int64 = (v35 & (v35 << ((2) & 63)))
        _ = v36
        _ = v37
        _ = v38
    }
    v39 := GenS4{F0: 2, F1: v34}
    switch "\n\t" {
    case "a":
        var v40 int = 796
        _ = v40
    case "foo bar":
        v41 := GenS5{F0: 42}
        _ = v41
    }
    var v42 uint32 = 71
    v3 = v35
    v35 = (v35 ^ 0)
    fmt.Println(v2, v2)
    var v43 int32 = 0
    for fi44 := 0; fi44 < 14; fi44++ {
        v45 := GenS2{F0: 65535, F1: (("\n\t" + "hello") + "\n\t"), F2: -32768}
        for fi46 := 0; fi46 < 10; fi46++ {
            v47 := GenS2{F0: 18446744073709551615, F1: "\n\t", F2: -1}
            _ = fi46
            _ = v47
        }
        v48 := GenS2{F0: 462, F1: (("foo bar" + "") + "foo bar"), F2: -1}
        _ = fi44
        _ = v45
        _ = v48
    }
    if (-1 < 0) {
        var v49 uint64 = 155
        if (false && (!v34)) {
            fmt.Println(v42, v49, v2)
            switch 1 {
            case 1:
                var v50 int = 648
                _ = v50
            }
            for fi51 := 0; fi51 < 7; fi51++ {
                var v52 uint64 = (v49 | (v49 * (v49 % (((^(v49))) | 1))))
                v2 += v33
                for fi53 := 0; fi53 < 13; fi53++ {
                    var v54 int32 = int32(v42)
                    var v55 uint16 = 1
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                _ = fi51
                _ = v52
            }
        } else {
            v56 := GenS4{F0: v49, F1: (((v49 <= v49) && v34) && (v2 != (float64(float32(v2)) - float64((1.5 + (v1 + ((v1 - v1) - (3.14159 + v1))))))))}
            _ = v56
        }
        _ = v49
    } else {
        var v57 int32 = (v43 + (v43 * 1))
        for fi58 := 0; fi58 < 19; fi58++ {
            var v59 float64 = (1e10 - -2.25)
            if v34 {
                if false {
                    var v60 string = (("" + "") + "foo bar")
                    for fi61 := 0; fi61 < 4; fi61++ {
                        var v62 int8 = 0
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if v34 {
                            v42 -= (v42 + v42)
                            v4 = append(v4, -106)
                        } else {
                            fmt.Println(v35, v42, v34)
                            _ = len(v4)
                            var v63 int64 = v35
                            v57 = (v43 / (((v57 - (-(v43)))) | 1))
                            if (!(true || ((v34 || (!true)) && v34))) {
                                var v64 uint64 = 42
                                var v65 uint32 = (^(v42))
                                _ = v64
                                _ = v65
                            } else {
                                var v66 uint8 = 2
                                v43 = v57
                                switch -464 {
                                case 9223372036854775807:
                                    v42 = (v42 * (v42 >> ((1) & 63)))
                                case 42:
                                    v4 = append(v4, 127)
                                    fmt.Println(v57)
                                }
                                _ = v66
                            }
                            _ = v63
                        }
                        _ = fi61
                        _ = v62
                    }
                    _ = v60
                }
                for fi67 := 0; fi67 < 20; fi67++ {
                    if true {
                        v2 = v33
                        fmt.Println(v1, v57, v35)
                        if v34 {
                            v33 -= v33
                            if true {
                                if ((-74 > -128) || true) {
                                    v68 := GenS4{F0: 509, F1: ((!(v34 || true)) && false)}
                                    var v69 int32 = (v43 + 774)
                                    v70 := GenS1{F0: 592, F1: true, F2: v34}
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                } else {
                                    if (v57 < (v43 * v57)) {
                                        var v71 int8 = 0
                                        v34 = v34
                                        switch "\n\t" {
                                        case "foo bar":
                                            v1 = 1.5
                                        default:
                                            v4 = append(v4, (fi58 - (fi67 - fi67)))
                                        }
                                        _ = v71
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = (fi58 / ((2147483647) | 1))
                                    }
                                    if v34 {
                                        switch fi67 {
                                        case -1:
                                            var v72 int64 = v35
                                            _ = v72
                                        case 32767:
                                            if (861 == 2) {
                                                if len(v4) > 0 {
                                                    v4[0] = -199
                                                }
                                                var v73 uint32 = 471
                                                fmt.Println(v42, v34, fi58)
                                                _ = v73
                                            } else {
                                                var v74 string = "hello"
                                                var v75 uint16 = 42
                                                v74 = "hello"
                                                _ = v74
                                                _ = v75
                                            }
                                        case -766:
                                            if ((1.5 * v59) != ((1.5 - (((float64((((v1 * v1) / (v1 - ((float32((((v2 / 1e10) + 3.14159) / v2)) * 0.0) - v1))) - v1)) / v33) - v33) - (v33 / v59))) - 1e10)) {
                                                v42 = (v42 - genFn0(255, uint64(v42), v43))
                                                switch (fi58 << ((1) & 63)) {
                                                case 1:
                                                    if (!(235 < 18446744073709551615)) {
                                                        if v34 {
                                                            var v76 uint32 = (v42 >> ((18446744073709551615) & 63))
                                                            _ = v76
                                                        }
                                                        var v77 int64 = (v35 % ((42) | 1))
                                                        _ = v77
                                                    } else {
                                                        v78 := GenS1{F0: 746, F1: v34, F2: false}
                                                        _ = v78
                                                    }
                                                case 128:
                                                    var v79 uint32 = (v42 << ((255) & 63))
                                                    _ = v79
                                                case 106:
                                                    v80 := GenS6{F0: genFn0(200, 18446744073709551615, 2147483647)}
                                                    _ = v80
                                                default:
                                                    if (!true) {
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v81 uint16 = 42
                                                            _ = v81
                                                        case "a":
                                                            v82 := GenS7{F0: "hello", F1: ("foo bar" + ("a" + ("a" + ("\n\t" + ("\n\t" + "\n\t")))))}
                                                            var v83 float32 = ((((-2.25 - 3.14159) / v1) / v1) / ((v1 / v1) / v1))
                                                            _ = v82
                                                            _ = v83
                                                        case "hello":
                                                            if v34 {
                                                                if (v34 && (!true)) {
                                                                    var v84 int8 = 127
                                                                    if (((("\n\t" + (((("hello" + "a") + "\n\t") + ("\n\t" + "\n\t")) + ("" + ((("\n\t" + "\n\t") + "foo bar") + "a")))) < "a") && false) && v34) {
                                                                        v4 = append(v4, (fi67 + (fi67 - fi58)))
                                                                    } else {
                                                                        fmt.Println(v35, v35, v42)
                                                                        v85 := GenS6{F0: (v42 * v42)}
                                                                        var v86 string = "a"
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                    _ = v84
                                                                } else {
                                                                    v87 := GenS7{F0: ("" + ""), F1: ("\n\t" + "a")}
                                                                    if false {
                                                                        v34 = v34
                                                                    }
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    _ = v87
                                                                }
                                                                switch 4294967295 {
                                                                case 830:
                                                                    switch fi67 {
                                                                    case -361:
                                                                        v88 := GenS7{F0: "", F1: ""}
                                                                        _ = v88
                                                                    case 583:
                                                                        if v34 {
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            fmt.Println(v1)
                                                                            switch "a" {
                                                                            case "a":
                                                                                if true {
                                                                                    v4 = append(v4, (-(fi58)))
                                                                                    v34 = (int8(v3) < 42)
                                                                                } else {
                                                                                    if true {
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = 9223372036854775807
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case "hello":
                                                                                v34 = false
                                                                            case "foo bar":
                                                                                var v89 uint64 = 4294967295
                                                                                _ = v89
                                                                            }
                                                                        }
                                                                        switch (("hello" + "") + "hello") {
                                                                        case "hello":
                                                                            v90 := GenS7{F0: ("a" + (("hello" + "a") + "\n\t")), F1: "foo bar"}
                                                                            _ = v90
                                                                        case "foo bar":
                                                                            fmt.Println(v33, v1)
                                                                        }
                                                                        v91 := GenS1{F0: 480, F1: true, F2: (!v34)}
                                                                        _ = v91
                                                                    case -1:
                                                                        fmt.Println(v2)
                                                                    default:
                                                                        v92 := GenS5{F0: 42}
                                                                        _ = v92
                                                                    }
                                                                    if v34 {
                                                                        var v93 float32 = v1
                                                                        _ = len(v4)
                                                                        switch 32768 {
                                                                        case 42:
                                                                            var v94 int = fi67
                                                                            _ = v94
                                                                        case 81:
                                                                            fmt.Println(fi58, v57)
                                                                        case 772:
                                                                            v34 = v34
                                                                            if true {
                                                                                v95 := GenS1{F0: 18446744073709551615, F1: ((v35 ^ v35) == int64(v42)), F2: ((v3 != (v35 ^ (v35 ^ (v35 + v3)))) && ((255 <= 42) || true))}
                                                                                _ = v95
                                                                            }
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                        }
                                                                        _ = v93
                                                                    } else {
                                                                        var v96 int64 = (v35 + (v35 % ((v35) | 1)))
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 127
                                                                        }
                                                                        _ = v96
                                                                    }
                                                                case -393:
                                                                    switch (fi58 * -129) {
                                                                    case 42:
                                                                        if len(v4) > 0 {
                                                                            v4[0] = (fi67 / ((fi67) | 1))
                                                                        }
                                                                        if (("\n\t" == ((("hello" + "\n\t") + "") + "foo bar")) || (((^(v42)) >= v42) || (false && v34))) {
                                                                            v42 = genFn0(127, 831, (v57 + v43))
                                                                            switch "" {
                                                                            case "a":
                                                                                if false {
                                                                                    v97 := GenS3{F0: 896, F1: (!(1 < 65535))}
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = fi58
                                                                                    }
                                                                                    _ = v97
                                                                                } else {
                                                                                    if ((!(("foo bar" == (("foo bar" + "foo bar") + (((("" + "\n\t") + "\n\t") + "\n\t") + ("" + "a")))) || v34)) || (!(929 <= 811))) {
                                                                                        v1 -= (0.1 / 1.0)
                                                                                        v57 += -1
                                                                                    } else {
                                                                                        var v98 uint64 = 2
                                                                                        fmt.Println(v43)
                                                                                        var v99 int8 = -1
                                                                                        _ = v98
                                                                                        _ = v99
                                                                                    }
                                                                                }
                                                                            case "foo bar":
                                                                                switch "\n\t" {
                                                                                case "\n\t":
                                                                                    fmt.Println(v43, v35, v57)
                                                                                case "hello":
                                                                                    var v100 uint64 = 18446744073709551615
                                                                                    _ = v100
                                                                                case "a":
                                                                                    fmt.Println(v57)
                                                                                }
                                                                                var v101 int16 = 1
                                                                                _ = v101
                                                                            }
                                                                            if true {
                                                                                if v34 {
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    if (1 == 228) {
                                                                                        fmt.Println(v34, v59)
                                                                                        var v102 uint32 = (v42 >> ((382) & 63))
                                                                                        _ = v102
                                                                                    } else {
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = (fi58 ^ 9223372036854775807)
                                                                                        }
                                                                                        var v103 uint8 = 168
                                                                                        fmt.Println(v34, v59)
                                                                                        v104 := GenS5{F0: 351}
                                                                                        _ = v103
                                                                                        _ = v104
                                                                                    }
                                                                                } else {
                                                                                    v105 := GenS7{F0: "foo bar", F1: "a"}
                                                                                    var v106 float64 = v2
                                                                                    _ = v105
                                                                                    _ = v106
                                                                                }
                                                                                v107 := GenS7{F0: "\n\t", F1: "foo bar"}
                                                                                if (!(!v34)) {
                                                                                    v57 = v57
                                                                                    v108 := GenS8{F0: (v43 ^ (v43 * v43))}
                                                                                    _ = v108
                                                                                }
                                                                                _ = v107
                                                                            } else {
                                                                                var v109 int16 = -76
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = fi67
                                                                                }
                                                                                fmt.Println(v42)
                                                                                _ = v109
                                                                            }
                                                                        }
                                                                        var v110 uint8 = 1
                                                                        _ = len(v4)
                                                                        v42 = (v42 + (v42 ^ (v42 - v42)))
                                                                        v4 = append(v4, 127)
                                                                        var v111 int8 = 42
                                                                        fmt.Println(v42)
                                                                        _ = v110
                                                                        _ = v111
                                                                    case -91:
                                                                        var v112 uint8 = 2
                                                                        _ = v112
                                                                    default:
                                                                        if ((v34 || ((false || true) || v34)) && (!true)) {
                                                                            if (fi58 < (fi67 ^ fi67)) {
                                                                                if (v34 && (189 > v35)) {
                                                                                    if (!(351 != 0)) {
                                                                                        v4 = append(v4, fi58)
                                                                                        if (220 == 1) {
                                                                                            fmt.Println(v3, v59)
                                                                                            v35 += 489
                                                                                        } else {
                                                                                            var v113 string = ""
                                                                                            _ = v113
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi67, v2, fi58)
                                                                                    var v114 uint = 65535
                                                                                    v57 += v57
                                                                                    _ = v114
                                                                                }
                                                                            } else {
                                                                                var v115 uint64 = 0
                                                                                switch ("\n\t" + ("\n\t" + "a")) {
                                                                                case "foo bar":
                                                                                    v59 -= float64(((v1 - (float32(float64(3.14159)) / -2.25)) * v1))
                                                                                case "\n\t":
                                                                                    v35 += (v3 / ((42) | 1))
                                                                                }
                                                                                _ = v115
                                                                            }
                                                                            if v34 {
                                                                                v4 = append(v4, (fi58 % (((-(fi67))) | 1)))
                                                                                v116 := GenS9{F0: -128}
                                                                                var v117 int16 = 0
                                                                                _ = v116
                                                                                _ = v117
                                                                            }
                                                                            fmt.Println(v34, fi67, v33)
                                                                            switch fi58 {
                                                                            case 4294967295:
                                                                                v57 = v57
                                                                            case 9223372036854775807:
                                                                                var v118 uint = 1
                                                                                _ = v118
                                                                            case 1:
                                                                                var v119 uint32 = (v42 << ((18446744073709551615) & 63))
                                                                                _ = v119
                                                                            default:
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    v120 := GenS10{F0: v3, F1: 349}
                                                                                    _ = v120
                                                                                case "":
                                                                                    var v121 string = "\n\t"
                                                                                    _ = v121
                                                                                case "hello":
                                                                                    switch ("foo bar" + (("foo bar" + "hello") + "foo bar")) {
                                                                                    case "hello":
                                                                                        v3 -= -717
                                                                                        var v122 uint = 1
                                                                                        switch "foo bar" {
                                                                                        case "foo bar":
                                                                                            v1 += 1e10
                                                                                        case "\n\t":
                                                                                            if (-617 != genFn1(v34, 62, genFn0(255, 18446744073709551615, (^(v57))))) {
                                                                                                v33 -= -2.25
                                                                                            } else {
                                                                                                v123 := GenS11{F0: 473, F1: 201}
                                                                                                v57 += (v57 << (((^(v122))) & 63))
                                                                                                var v124 uint8 = 204
                                                                                                _ = v123
                                                                                                _ = v124
                                                                                            }
                                                                                        case "hello":
                                                                                            if false {
                                                                                                var v125 string = ("a" + "\n\t")
                                                                                                _ = v125
                                                                                            }
                                                                                        }
                                                                                        _ = v122
                                                                                    }
                                                                                default:
                                                                                    switch (fi58 >> ((183) & 63)) {
                                                                                    case 1:
                                                                                        v33 = (v33 - float64((1e10 * (v1 * v1))))
                                                                                    case 673:
                                                                                        var v126 bool = ((((v1 - v1) / v1) != ((((-2.25 - float32(v33)) + (3.14159 + v1)) + 0.1) - 0.1)) && (!false))
                                                                                        _ = v126
                                                                                    case -126:
                                                                                        v4 = append(v4, (fi67 >> ((1) & 63)))
                                                                                    default:
                                                                                        fmt.Println(v57, v2, v34)
                                                                                    }
                                                                                }
                                                                            }
                                                                        } else {
                                                                            _ = len(v4)
                                                                            v127 := GenS4{F0: 18446744073709551615, F1: ((!((4294967295 == 1) || (genFn0(2, 2, -363) == 4294967295))) && false)}
                                                                            v33 = v33
                                                                            _ = v127
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v2 += v2
                                                    } else {
                                                        v59 -= v2
                                                        fmt.Println(v33, v57, v57)
                                                        if false {
                                                            v128 := GenS8{F0: (v57 * int32(fi58))}
                                                            if v34 {
                                                                v35 = -198
                                                            } else {
                                                                v35 -= (v35 | (v35 + v3))
                                                                fmt.Println(v57, v59, fi58)
                                                                v3 = v35
                                                            }
                                                            var v129 uint64 = 42
                                                            _ = v128
                                                            _ = v129
                                                        } else {
                                                            var v130 float32 = ((3.14159 - v1) - v1)
                                                            v131 := GenS6{F0: genFn0(127, 0, (v57 ^ int32(v35)))}
                                                            fmt.Println(fi67)
                                                            var v132 int = -946
                                                            var v133 int64 = (^(v3))
                                                            v133 = (v133 / ((9223372036854775807) | 1))
                                                            var v134 float32 = ((v1 * 0.1) - v130)
                                                            _ = v130
                                                            _ = v131
                                                            _ = v132
                                                            _ = v133
                                                            _ = v134
                                                        }
                                                    }
                                                }
                                                var v135 float64 = 1e10
                                                _ = v135
                                            } else {
                                                _ = len(v4)
                                                v4 = append(v4, 9223372036854775807)
                                            }
                                        default:
                                            v35 = (v35 ^ (v35 << ((65535) & 63)))
                                            fmt.Println(v42, fi58, v42)
                                        }
                                        v4 = append(v4, (fi67 ^ fi58))
                                        v136 := GenS3{F0: 42, F1: (v34 && (v34 && v34))}
                                        v137 := GenS12{F0: 1}
                                        if v34 {
                                            v138 := GenS12{F0: 0}
                                            v139 := GenS1{F0: 246, F1: v34, F2: true}
                                            _ = v138
                                            _ = v139
                                        } else {
                                            v34 = (!v34)
                                            v2 += v59
                                        }
                                        _ = v136
                                        _ = v137
                                    } else {
                                        if true {
                                            fmt.Println(v35, fi67)
                                            v42 -= (v42 % ((genFn0(0, 18446744073709551615, (v43 >> ((614) & 63)))) | 1))
                                            var v140 float64 = ((v33 - v59) / (v2 * float64(v1)))
                                            _ = v140
                                        }
                                        v141 := GenS1{F0: 4294967295, F1: v34, F2: (v33 <= v33)}
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v142 float64 = (v59 * v2)
                                            _ = v142
                                        case "foo bar":
                                            if v34 {
                                                var v143 float64 = v59
                                                v59 += (0.0 + float64(1.5))
                                                _ = v143
                                            } else {
                                                var v144 float64 = 0.1
                                                v145 := GenS13{F0: 243}
                                                _ = v144
                                                _ = v145
                                            }
                                        }
                                        fmt.Println(v2, v33, v1)
                                        v57 -= 2147483647
                                        _ = v141
                                    }
                                }
                            } else {
                                v57 += (v57 * v57)
                                var v146 uint = 784
                                v57 = -2147483648
                                _ = v146
                            }
                        }
                    } else {
                        if (32767 > -404) {
                            fmt.Println(fi58, fi67, v57)
                            var v147 int64 = 42
                            v59 += -2.25
                            _ = v147
                        }
                        v148 := GenS11{F0: 2, F1: 128}
                        v149 := GenS12{F0: 94}
                        _ = v148
                        _ = v149
                    }
                    var v150 float64 = v59
                    if true {
                        v151 := GenS14{F0: (("a" + "foo bar") + "a"), F1: (v35 | -672), F2: "foo bar"}
                        switch "" {
                        case "":
                            var v152 uint64 = 614
                            _ = v152
                        case "hello":
                            if v34 {
                                switch "\n\t" {
                                case "a":
                                    v153 := GenS10{F0: (v35 & v35), F1: 255}
                                    _ = v153
                                default:
                                    v4 = append(v4, -504)
                                }
                                var v154 int16 = 544
                                _ = v154
                            }
                        case "a":
                            var v155 uint32 = (^(v42))
                            _ = v155
                        default:
                            v4 = append(v4, fi58)
                        }
                        v156 := GenS8{F0: (^(v57))}
                        _ = v151
                        _ = v156
                    }
                    _ = fi67
                    _ = v150
                }
                var v157 uint8 = 128
                _ = v157
            } else {
                var v158 uint64 = 1
                _ = v158
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi58
            _ = v59
        }
        for fi159 := 0; fi159 < 14; fi159++ {
            if ((("a" + "hello") < ("hello" + "foo bar")) && v34) {
                var v160 uint8 = 128
                _ = v160
            }
            _ = fi159
        }
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v42
    _ = v43
}
