package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int8; F1 string }
type GenS11 struct { F0 uint16; F1 uint8; F2 uint }
type GenS12 struct { F0 uint64; F1 int64; F2 string }
type GenS13 struct { F0 uint; F1 uint32 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint8; F1 int32 }
type GenS16 struct { F0 uint32; F1 string }
type GenS17 struct { F0 int16; F1 uint32 }
type GenS18 struct { F0 uint64; F1 uint64 }
type GenS19 struct { F0 string; F1 int64; F2 int32 }
func genFn0() int32 {
    v5 := []int{473, 236, 32768, 128}
    var v6 bool = false
    switch "hello" {
    case "a":
        var v7 uint64 = 442
        _ = v7
    case "":
        v6 = false
    }
    _ = len(v5)
    v6 = (-798 == -682)
    _ = v5
    _ = v6
    return -2147483648
}
func genFn1(p0 int64, p1 uint8, p2 uint64) bool {
    if ((!false) && ((p1 == (p1 ^ p1)) && (32767 <= 1))) {
        p1 += (p1 + uint8(p2))
        p1 = (p1 << ((1) & 63))
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                p0 -= (^(p0))
                var v10 uint8 = 0
                _ = fi9
                _ = v10
            }
            p2 -= p2
            _ = fi8
        }
    } else {
        p1 = p1
        v11 := []int{-805, 42}
        var v12 float32 = 1.5
        _ = v11
        _ = v12
    }
    var v13 uint64 = p2
    p2 = 42
    v14 := GenS1{F0: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return false
}

func main() {
    var v1 uint32 = 971
    v2 := []string{}
    v3 := make(map[string]bool)
    var v4 uint = 42
    v4 += 4294967295
    if (32767 >= -9223372036854775808) {
        v15 := GenS1{F0: true}
        v16 := GenS1{F0: (!(42 > 1))}
        _ = v15
        _ = v16
    }
    if true {
        var v17 float32 = (float32(3.14159) - (0.0 * 3.14159))
        var v18 uint64 = 0
        _ = v17
        _ = v18
    } else {
        var v19 int16 = 32767
        if (!(v19 >= 0)) {
            v20 := GenS2{F0: v19, F1: (^(v4)), F2: 1}
            v19 -= (v19 ^ (^(v19)))
            _ = v20
        }
        _ = v19
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS3{F0: 0}
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 uint32 = 42
            v4 = (v4 / (((v4 + (v4 % ((v4) | 1)))) | 1))
            v25 := GenS1{F0: (!false)}
            v26 := GenS2{F0: 0, F1: v4, F2: 103}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        if ((v1 / (((v1 / (((v1 / ((v1) | 1))) | 1))) | 1)) == (v1 / ((2147483648) | 1))) {
            for fi27 := 0; fi27 < 7; fi27++ {
                if (false || (0.1 < 0.1)) {
                    v2 = append(v2, (("a" + "hello") + "a"))
                    v28 := v3[(((("hello" + "") + "hello") + "a") + "")]
                    _ = v28
                } else {
                    v29 := GenS3{F0: 0}
                    _ = v29
                }
                v30 := GenS4{F0: "a", F1: v1}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch "\n\t" {
                case "foo bar":
                    if true {
                        var v31 uint32 = 2147483648
                        _ = v31
                    } else {
                        var v32 uint32 = 5
                        v4 += v4
                        _ = v32
                    }
                    switch fi27 {
                    case -599:
                        v3["a"] = true
                        fmt.Println(fi21, v1, v4)
                    case 128:
                        v33 := GenS5{F0: "\n\t"}
                        _ = v33
                    case 127:
                        v34 := GenS6{F0: 346}
                        _ = v34
                    }
                case "":
                    var v35 uint32 = (v1 | v1)
                    _ = v35
                case "a":
                    if false {
                        var v36 int = (^(fi27))
                        switch fi27 {
                        case 127:
                            _ = len(v3)
                        case 42:
                            _ = len(v2)
                        case 1:
                            var v37 int8 = -128
                            _ = v37
                        }
                        _ = v36
                    }
                }
                if true {
                    if (v4 >= v4) {
                        if true {
                            var v38 int8 = -113
                            switch (fi21 & fi21) {
                            case -737:
                                var v39 uint = 0
                                _ = v39
                            case 42:
                                if true {
                                    var v40 int16 = -1
                                    _ = v40
                                }
                            case 128:
                                switch "hello" {
                                case "hello":
                                    if false {
                                        var v41 uint8 = 127
                                        v1 += v1
                                        _ = v41
                                    } else {
                                        v42 := v3[("\n\t" + "")]
                                        if v42 {
                                            var v43 float32 = float32((3.14159 / 1.0))
                                            v3[("foo bar" + "foo bar")] = (v1 == (^(v1)))
                                            _ = v43
                                        } else {
                                            v44 := v3["\n\t"]
                                            var v45 float64 = 3.14159
                                            _ = v44
                                            _ = v45
                                        }
                                        if (v42 || false) {
                                            if len(v2) > 0 {
                                                v2[0] = ""
                                            }
                                            switch (("\n\t" + "foo bar") + "foo bar") {
                                            case "foo bar":
                                                if false {
                                                    var v46 uint16 = 523
                                                    var v47 uint16 = 783
                                                    v48 := GenS6{F0: (^(v47))}
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                }
                                            }
                                        }
                                        _ = v42
                                    }
                                case "a":
                                    v49 := GenS7{F0: ("\n\t" + "\n\t"), F1: 140}
                                    _ = v49
                                case "foo bar":
                                    if false {
                                        v50 := GenS2{F0: 767, F1: 42, F2: 1}
                                        var v51 float64 = ((3.14159 + -2.25) * 3.14159)
                                        v52 := GenS8{F0: (v38 << (((v4 ^ (v4 / (((v4 >> ((v4) & 63))) | 1)))) & 63))}
                                        switch "" {
                                        case "":
                                            _ = len(v2)
                                        case "\n\t":
                                            fmt.Println(v51, fi21, v51)
                                        default:
                                            v1 += (v1 << (((v4 - (^(v4)))) & 63))
                                        }
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    }
                                }
                            }
                            v53 := GenS9{F0: 0}
                            _ = v38
                            _ = v53
                        }
                        v1 -= (v1 & (v1 & (^(v1))))
                        switch ("" + "") {
                        case "a":
                            v3[""] = true
                            fmt.Println(v1, fi21)
                            _ = len(v3)
                        case "":
                            v1 -= v1
                            v4 -= (v4 * (v4 ^ 745))
                            switch (fi27 / ((-493) | 1)) {
                            case 1:
                                var v54 uint = 0
                                _ = v54
                            }
                            v1 -= v1
                            v55 := GenS2{F0: 0, F1: (v4 / ((v4) | 1)), F2: 42}
                            _ = v55
                        default:
                            var v56 float64 = 0.1
                            _ = v56
                        }
                    } else {
                        var v57 uint64 = 18446744073709551615
                        if (1e10 <= 3.14159) {
                            switch "a" {
                            case "hello":
                                if true {
                                    var v58 uint = (^(v4))
                                    v57 = uint64(v4)
                                    _ = v58
                                } else {
                                    var v59 string = "\n\t"
                                    var v60 uint16 = 1
                                    v4 -= (v4 >> (((v4 << ((v4) & 63))) & 63))
                                    if ((-106 < 100) && (!(v59 > ""))) {
                                        v61 := GenS8{F0: -1}
                                        var v62 uint16 = (v60 - (v60 * (v60 ^ 345)))
                                        _ = v61
                                        _ = v62
                                    } else {
                                        v59 = ((v59 + "hello") + ((v59 + v59) + v59))
                                        if false {
                                            fmt.Println(v57, v4)
                                            if (false && false) {
                                                if true {
                                                    v63 := GenS1{F0: false}
                                                    var v64 uint8 = 1
                                                    v1 = 244
                                                    _ = v63
                                                    _ = v64
                                                }
                                                v59 = v59
                                            } else {
                                                delete(v3, "a")
                                            }
                                            v65 := GenS10{F0: 0, F1: v59}
                                            _ = v65
                                        }
                                    }
                                    if false {
                                        v66 := GenS7{F0: (((((("a" + ((v59 + v59) + v59)) + (v59 + "a")) + v59) + "") + "\n\t") + ""), F1: 20}
                                        switch (-(fi27)) {
                                        case 127:
                                            delete(v3, v59)
                                        case 56:
                                            var v67 uint = 0
                                            _ = v67
                                        case -346:
                                            fmt.Println(v1, fi27)
                                        }
                                        _ = v66
                                    } else {
                                        v1 = v1
                                        fmt.Println(fi27, v60, fi21)
                                        v57 = 232
                                    }
                                    v57 = v57
                                    v57 -= v57
                                    _ = v59
                                    _ = v60
                                }
                            case "a":
                                if (0.0 != 1e10) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if (((float64(1e10) - ((1e10 + 0.0) * float64(0.0))) - 0.1) < (1e10 - 3.14159)) {
                                        var v68 uint = uint(v57)
                                        v57 -= (^(v57))
                                        switch "foo bar" {
                                        case "":
                                            if false {
                                                if true {
                                                    var v69 int64 = 9223372036854775807
                                                    switch (fi27 + 127) {
                                                    case -149:
                                                        var v70 uint32 = 868
                                                        _ = v70
                                                    case 241:
                                                        if ((fi27 & (fi27 | 32768)) <= fi21) {
                                                            if len(v2) > 0 {
                                                                v2[0] = "\n\t"
                                                            }
                                                        } else {
                                                            v57 += 0
                                                            var v71 int32 = -2147483648
                                                            if false {
                                                                if true {
                                                                    fmt.Println(v68, v4, fi27)
                                                                }
                                                            }
                                                            _ = v71
                                                        }
                                                    }
                                                    if false {
                                                        v1 -= (v1 * 4294967295)
                                                        fmt.Println(v4, v68, fi27)
                                                    }
                                                    _ = v69
                                                }
                                                if false {
                                                    v72 := GenS11{F0: 0, F1: 182, F2: v68}
                                                    v68 = (v4 - v4)
                                                    _ = v72
                                                } else {
                                                    var v73 uint8 = 0
                                                    if (531 < 1) {
                                                        if len(v2) > 0 {
                                                            v2[0] = ("\n\t" + ("hello" + ("hello" + "foo bar")))
                                                        }
                                                    } else {
                                                        fmt.Println(v68, v73)
                                                    }
                                                    fmt.Println(v73, v4, v57)
                                                    _ = v73
                                                }
                                            }
                                        case "foo bar":
                                            var v74 uint = v4
                                            _ = v74
                                        case "a":
                                            if len(v2) > 0 {
                                                v2[0] = "a"
                                            }
                                        default:
                                            v1 += 4294967295
                                        }
                                        _ = v68
                                    }
                                }
                            case "\n\t":
                                v57 += 4294967295
                                var v75 float64 = 1.5
                                var v76 int64 = 447
                                var v77 bool = true
                                if (true && v77) {
                                    v1 = (v1 & (v1 + v1))
                                    fmt.Println(v76)
                                    var v78 uint16 = 42
                                    if (-38 >= 100) {
                                        if v77 {
                                            v4 = v4
                                        }
                                    } else {
                                        v75 = v75
                                        if false {
                                            v57 -= (v57 >> ((v4) & 63))
                                        } else {
                                            switch ("hello" + "hello") {
                                            case "a":
                                                v78 = 255
                                            case "\n\t":
                                                v79 := v3[("hello" + "hello")]
                                                _ = v79
                                            case "hello":
                                                fmt.Println(v78, fi21, v75)
                                            default:
                                                v75 = ((((0.1 + float64(0.0)) + (float64(0.1) - v75)) * 0.0) + float64(3.14159))
                                            }
                                            var v80 string = "hello"
                                            if true {
                                                fmt.Println(v80)
                                                var v81 int32 = genFn0()
                                                _ = v81
                                            } else {
                                                var v82 int32 = 252
                                                switch (fi21 >> (((v4 * (v4 - (^(v4))))) & 63)) {
                                                case 231:
                                                    v76 = v76
                                                case 32767:
                                                    fmt.Println(v77, v77, v75)
                                                case -1:
                                                    v83 := GenS3{F0: v57}
                                                    _ = v83
                                                }
                                                _ = v82
                                            }
                                            _ = v80
                                        }
                                    }
                                    _ = v78
                                } else {
                                    v77 = (((float32(v75) + (3.14159 / 3.14159)) > float32(v75)) && v77)
                                    switch "a" {
                                    case "foo bar":
                                        v84 := GenS10{F0: -1, F1: ""}
                                        _ = v84
                                    case "\n\t":
                                        if v77 {
                                            if true {
                                                v76 += v76
                                                delete(v3, "\n\t")
                                                fmt.Println(v76, fi21)
                                            } else {
                                                if true {
                                                    v85 := GenS12{F0: v57, F1: v76, F2: ("a" + ("foo bar" + ""))}
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = v85
                                                } else {
                                                    fmt.Println(v4, fi21, v1)
                                                    if (v77 && (!true)) {
                                                        v77 = v77
                                                        var v86 bool = v77
                                                        _ = v86
                                                    } else {
                                                        fmt.Println(fi27, v1)
                                                        var v87 int16 = 719
                                                        _ = v87
                                                    }
                                                    if ((-2.25 / float32(1.5)) == (-2.25 + 0.1)) {
                                                        fmt.Println(v77, fi21)
                                                        var v88 int16 = -66
                                                        v89 := GenS8{F0: 100}
                                                        _ = v88
                                                        _ = v89
                                                    } else {
                                                        v3[("" + "hello")] = ((1 != v1) && true)
                                                        if true {
                                                            v90 := GenS13{F0: 517, F1: v1}
                                                            _ = v90
                                                        } else {
                                                            v76 += v76
                                                            var v91 float64 = 3.14159
                                                            if v77 {
                                                                v92 := GenS14{F0: 1}
                                                                _ = v92
                                                            }
                                                            v3[""] = (!v77)
                                                            _ = v91
                                                        }
                                                    }
                                                }
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v75 = float64(0.1)
                                                case "":
                                                    v75 += v75
                                                case "hello":
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v76 = -494
                                                    default:
                                                        var v93 int16 = 32767
                                                        _ = v93
                                                    }
                                                default:
                                                    _ = len(v2)
                                                }
                                                delete(v3, (("\n\t" + ("\n\t" + "\n\t")) + ("foo bar" + (("" + ("" + "\n\t")) + "a"))))
                                            }
                                            v1 -= (v1 / (((v1 - 166)) | 1))
                                            switch fi27 {
                                            case -1:
                                                var v94 float32 = (0.1 - 3.14159)
                                                _ = v94
                                            case 185:
                                                if false {
                                                    var v95 float64 = 0.0
                                                    _ = v95
                                                } else {
                                                    if ((!(false || (!true))) || v77) {
                                                        _ = len(v3)
                                                        if v77 {
                                                            switch "\n\t" {
                                                            case "a":
                                                                fmt.Println(v75, v57)
                                                            }
                                                            if v77 {
                                                                var v96 int = (^(fi21))
                                                                _ = v96
                                                            } else {
                                                                if (((uint64(v4) >= (v57 & v57)) || (v76 <= (-(v76)))) || (true && false)) {
                                                                    var v97 uint64 = (v57 ^ v57)
                                                                    v98 := GenS11{F0: 1, F1: 2, F2: v4}
                                                                    _ = v97
                                                                    _ = v98
                                                                }
                                                            }
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v4 += v4
                                                            case "hello":
                                                                v2 = append(v2, "hello")
                                                            case "a":
                                                                v57 -= (v57 << ((4294967295) & 63))
                                                            }
                                                            var v99 uint8 = 0
                                                            _ = v99
                                                        }
                                                        _ = len(v3)
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        switch "a" {
                                                        case "hello":
                                                            var v100 uint8 = 255
                                                            _ = v100
                                                        case "foo bar":
                                                            v3["hello"] = v77
                                                        case "":
                                                            _ = len(v3)
                                                        }
                                                        switch fi27 {
                                                        case 128:
                                                            v101 := v3["\n\t"]
                                                            _ = v101
                                                        case 32768:
                                                            switch fi27 {
                                                            case 128:
                                                                v77 = false
                                                            case 4294967295:
                                                                if v77 {
                                                                    v102 := GenS15{F0: 0, F1: 0}
                                                                    _ = v102
                                                                }
                                                            }
                                                        default:
                                                            var v103 uint64 = v57
                                                            _ = v103
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    case "hello":
                                        v77 = ((^(v57)) < v57)
                                    }
                                    v104 := GenS16{F0: 415, F1: "hello"}
                                    v4 = (v4 >> (((v4 & 815)) & 63))
                                    _ = v104
                                }
                                _ = v75
                                _ = v76
                                _ = v77
                            }
                        } else {
                            fmt.Println(v1)
                        }
                        v4 += 0
                        var v105 uint16 = 2
                        var v106 float64 = ((float64((-2.25 + -2.25)) + (1e10 * 1.5)) / (0.0 - 1e10))
                        var v107 uint32 = v1
                        _ = v105
                        _ = v106
                        _ = v107
                        _ = v57
                    }
                    if (fi21 >= 127) {
                        var v108 string = ("a" + ("\n\t" + "foo bar"))
                        var v109 float64 = (3.14159 / (3.14159 + (float64(0.1) * 3.14159)))
                        _ = v108
                        _ = v109
                    }
                    switch (fi27 / ((fi21) | 1)) {
                    case -860:
                        fmt.Println(v4)
                    case -844:
                        v110 := GenS8{F0: -128}
                        _ = v110
                    }
                    v4 += uint(v1)
                    fmt.Println(v1, fi27)
                } else {
                    if (fi21 != fi21) {
                        fmt.Println(v1, v4)
                    }
                    fmt.Println(fi21)
                    var v111 string = "hello"
                    _ = v111
                }
                delete(v3, "a")
                _ = fi27
                _ = v30
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if true {
                v4 = (v4 & v4)
                for fi112 := 0; fi112 < 13; fi112++ {
                    v3[""] = (1 < 42)
                    var v113 bool = (!true)
                    _ = fi112
                    _ = v113
                }
            } else {
                fmt.Println(v1, v4, v1)
                var v114 int = fi21
                for fi115 := 0; fi115 < 18; fi115++ {
                    if true {
                        delete(v3, (("foo bar" + "") + ("hello" + ("foo bar" + "hello"))))
                        v116 := GenS17{F0: -1, F1: (v1 & (v1 % (((v1 << ((4294967295) & 63))) | 1)))}
                        var v117 bool = (42 < -1)
                        _ = v116
                        _ = v117
                    }
                    v4 = (v4 | v4)
                    _ = fi115
                }
                _ = v114
            }
        }
        _ = fi21
        _ = v22
    }
    v118 := v3[("a" + "foo bar")]
    v119 := GenS4{F0: "a", F1: (v1 - (v1 & (^(v1))))}
    if false {
        switch 963 {
        case -881:
            switch "" {
            case "hello":
                for fi120 := 0; fi120 < 17; fi120++ {
                    var v121 bool = ((v4 % ((v4) | 1)) < v4)
                    _ = fi120
                    _ = v121
                }
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
            case "\n\t":
                v122 := GenS18{F0: 903, F1: 42}
                _ = v122
            }
        default:
            v123 := v3["a"]
            _ = v123
        }
    } else {
        var v124 int64 = -124
        v125 := GenS11{F0: 346, F1: uint8(v1), F2: (v4 ^ (v4 << ((v4) & 63)))}
        _ = v124
        _ = v125
    }
    v1 += (v1 << (((^(v4))) & 63))
    v126 := GenS19{F0: (("hello" + "hello") + ("\n\t" + ("a" + ("a" + ("a" + ""))))), F1: -9223372036854775808, F2: 1}
    fmt.Println(v4, v118)
    fmt.Println("v1:", v1)
    fmt.Println("v118:", v118)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v118
    _ = v119
    _ = v126
    _ = v2
    _ = v3
    _ = v4
}
