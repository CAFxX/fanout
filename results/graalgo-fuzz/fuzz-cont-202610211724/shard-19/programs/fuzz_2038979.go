package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint; F1 int; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 string; F2 uint32 }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 uint64; F1 int16; F2 int16 }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int; F1 int32; F2 uint64 }
type GenS13 struct { F0 int32; F1 int32 }
type GenS14 struct { F0 int64; F1 uint16 }
func genFn0(p0 bool) uint8 {
    v4 := GenS1{F0: 0}
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 float32 = float32(0.0)
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 bool = (!(!p0))
            v8 = true
            _ = fi7
            _ = v8
        }
        var v9 int64 = 1
        _ = fi5
        _ = v6
        _ = v9
    }
    v10 := GenS1{F0: -1}
    fmt.Println(p0)
    var v11 int64 = 9223372036854775807
    _ = p0
    _ = v10
    _ = v11
    _ = v4
    return 42
}
func genFn1(p0 uint32) uint8 {
    v12 := GenS1{F0: 42}
    p0 -= 65535
    var v13 int64 = -869
    _ = p0
    _ = v12
    _ = v13
    return 120
}

func main() {
    var v1 uint8 = 183
    var v2 int64 = 0
    v3 := make(map[int]bool)
    switch 0 {
    case -1:
        if false {
            v14 := []int{32767}
            v2 -= -345
            _ = v14
        }
    case 1:
        switch (("foo bar" + "hello") + "\n\t") {
        case "":
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := []int{}
                _ = fi15
                _ = v16
            }
        case "\n\t":
            v1 -= genFn1(409)
        default:
            for fi17 := 0; fi17 < 17; fi17++ {
                for fi18 := 0; fi18 < 7; fi18++ {
                    var v19 uint16 = 2
                    v1 = (v1 << ((2) & 63))
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
        }
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 uint32 = 42
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v23 uint16 = 65535
    switch "foo bar" {
    case "hello":
        if (1e10 != (1.5 / float64(float32((1.5 * (float64(1e10) * (3.14159 + (float64(3.14159) - (1.5 - (((((0.1 / 1.0) + 0.1) - 1e10) - float64(-2.25)) + 1e10)))))))))) {
            for fi24 := 0; fi24 < 16; fi24++ {
                fmt.Println(v2, v2)
                var v25 int = (fi24 + 4294967295)
                _ = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 4; fi26++ {
                v27 := GenS1{F0: v2}
                _ = fi26
                _ = v27
            }
            v1 = v1
        }
    case "\n\t":
        v23 = 65535
    case "foo bar":
        v1 = (^(v1))
    default:
        if ((!false) || true) {
            fmt.Println(v2, v2, v23)
            v28 := []int{0, 2147483647}
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        v30 := GenS2{F0: -123, F1: v23}
        var v31 int64 = 42
        _ = fi29
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        switch (fi32 * fi32) {
        case -854:
            switch 32767 {
            case -1:
                for fi33 := 0; fi33 < 7; fi33++ {
                    v34 := GenS3{F0: 0, F1: -780, F2: 586}
                    var v35 string = (("\n\t" + (("a" + "a") + "a")) + "")
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            case -2147483648:
                var v36 string = ("\n\t" + ("" + (("foo bar" + "a") + "\n\t")))
                _ = v36
            }
        }
        _ = fi32
    }
    v23 += v23
    var v37 int64 = -319
    var v38 int16 = 173
    if false {
        v39 := []int{}
        _ = v39
    }
    v40 := GenS4{F0: "\n\t"}
    v3[-1] = false
    switch -369 {
    case 9223372036854775807:
        var v41 uint = 18446744073709551615
        _ = v41
    case 1:
        if (!false) {
            switch -948 {
            case 127:
                var v42 string = "foo bar"
                _ = v42
            case 128:
                if (127 > -92) {
                    for fi43 := 0; fi43 < 17; fi43++ {
                        switch fi43 {
                        case -1:
                            if false {
                                var v44 uint16 = (^(v23))
                                v38 += -1
                                v44 += (v23 + (v44 & (v44 / ((0) | 1))))
                                _ = v44
                            }
                        default:
                            if true {
                                v2 += v2
                            } else {
                                var v45 bool = (v1 <= (^(v1)))
                                var v46 bool = (-88 > 100)
                                v47 := []int{(fi43 << ((0) & 63)), (fi43 / ((fi43) | 1)), 9223372036854775807}
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                        }
                        v1 += v1
                        _ = fi43
                    }
                    for fi48 := 0; fi48 < 10; fi48++ {
                        switch "\n\t" {
                        case "\n\t":
                            switch "" {
                            case "hello":
                                fmt.Println(fi48)
                            case "":
                                v38 = (v38 >> ((62) & 63))
                            case "foo bar":
                                v49 := GenS5{F0: (v23 % (((v23 * (^(v23)))) | 1)), F1: "", F2: 916}
                                _ = v49
                            default:
                                v50 := GenS5{F0: v23, F1: "hello", F2: 2147483647}
                                _ = v50
                            }
                        case "a":
                            fmt.Println(v37, v37)
                        case "hello":
                            if (uint64(v1) > 403) {
                                switch fi48 {
                                case 40:
                                    v51 := []int{fi48}
                                    _ = v51
                                case 460:
                                    if false {
                                        v52 := GenS6{F0: 508, F1: v2}
                                        v38 = 1000
                                        var v53 uint32 = 0
                                        _ = v52
                                        _ = v53
                                    }
                                default:
                                    switch "\n\t" {
                                    case "\n\t":
                                        for fi54 := 0; fi54 < 7; fi54++ {
                                            var v55 int = fi54
                                            if (1.5 >= 1.5) {
                                                v56 := GenS7{F0: 350, F1: (v38 << ((479) & 63)), F2: (v38 & v38)}
                                                var v57 uint = 42
                                                _ = v56
                                                _ = v57
                                            }
                                            if false {
                                                v58 := GenS4{F0: ("a" + ("foo bar" + "\n\t"))}
                                                v38 = v38
                                                _ = v58
                                            }
                                            _ = fi54
                                            _ = v55
                                        }
                                    }
                                }
                            } else {
                                for fi59 := 0; fi59 < 6; fi59++ {
                                    v60 := GenS2{F0: 1, F1: 475}
                                    var v61 int = (fi48 ^ (fi48 + fi48))
                                    _ = fi59
                                    _ = v60
                                    _ = v61
                                }
                                _ = len(v3)
                            }
                        default:
                            for fi62 := 0; fi62 < 12; fi62++ {
                                v63 := []int{(-(fi48)), fi48}
                                fmt.Println(v38, fi48)
                                switch "\n\t" {
                                case "":
                                    v2 -= v2
                                case "\n\t":
                                    v23 += 986
                                case "foo bar":
                                    switch "hello" {
                                    case "hello":
                                        switch fi48 {
                                        case 127:
                                            var v64 uint64 = 223
                                            _ = v64
                                        case 128:
                                            var v65 uint8 = genFn0(((v38 >> ((42) & 63)) != (v38 / ((v38) | 1))))
                                            _ = v65
                                        case -739:
                                            if ("foo bar" < ("" + "\n\t")) {
                                                var v66 float32 = 3.14159
                                                if false {
                                                    if (((!(65535 != uint(v1))) || ("" != "\n\t")) && true) {
                                                        v2 = (v2 & v37)
                                                        if (1 <= 245) {
                                                            v67 := GenS8{F0: 2, F1: (v37 % (((v2 - -142)) | 1))}
                                                            _ = v67
                                                        }
                                                    } else {
                                                        if len(v63) > 0 {
                                                            v63[0] = (fi48 + (fi62 >> ((18446744073709551615) & 63)))
                                                        }
                                                        var v68 bool = false
                                                        _ = v68
                                                    }
                                                    if false {
                                                        var v69 int = fi62
                                                        v3[v69] = true
                                                        var v70 uint64 = 2
                                                        if ((1.5 / (0.1 / float64((v66 - (float32(1e10) * 1e10))))) >= (0.1 - 3.14159)) {
                                                            if false {
                                                                if len(v63) > 0 {
                                                                    v63[0] = (fi48 ^ 0)
                                                                }
                                                                var v71 uint64 = v70
                                                                _ = v71
                                                            }
                                                            _ = len(v3)
                                                        } else {
                                                            var v72 uint32 = 409
                                                            fmt.Println(v72)
                                                            _ = v72
                                                        }
                                                        fmt.Println(v2, fi62, v69)
                                                        v73 := v3[int(v70)]
                                                        _ = v69
                                                        _ = v70
                                                        _ = v73
                                                    } else {
                                                        if len(v63) > 0 {
                                                            v63[0] = fi62
                                                        }
                                                        var v74 uint8 = 128
                                                        var v75 float64 = 1.5
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    if (!(!(2 >= v23))) {
                                                        fmt.Println(v2, fi48)
                                                        var v76 uint8 = genFn1(2)
                                                        if true {
                                                            v76 += (v1 - v76)
                                                            v37 = v2
                                                            if (!false) {
                                                                if (465 != (v23 & (v23 - v23))) {
                                                                    v77 := GenS7{F0: 2, F1: v38, F2: (v38 % ((v38) | 1))}
                                                                    switch 32767 {
                                                                    case 0:
                                                                        v37 = (v37 - v37)
                                                                    case -1:
                                                                        v76 = (v1 - (v1 - v76))
                                                                    default:
                                                                        if len(v63) >= 2 {
                                                                            v63 = v63[1:len(v63)]
                                                                        }
                                                                    }
                                                                    switch fi48 {
                                                                    case 261:
                                                                        v1 -= genFn1(801)
                                                                    }
                                                                    if ((4294967295 > 324) && (!((v23 ^ (v23 << ((uint(v23)) & 63))) > (v23 % (((v23 + 2)) | 1))))) {
                                                                        v1 = v1
                                                                    }
                                                                    _ = v77
                                                                }
                                                                var v78 int8 = 1
                                                                switch -944 {
                                                                case -9223372036854775808:
                                                                    if (!false) {
                                                                        delete(v3, 9223372036854775807)
                                                                    }
                                                                }
                                                                _ = v78
                                                            }
                                                        }
                                                        _ = v76
                                                    } else {
                                                        v79 := GenS6{F0: (v23 - (v23 / (((^(v23))) | 1))), F1: 1}
                                                        _ = v79
                                                    }
                                                } else {
                                                    v80 := GenS9{F0: 1}
                                                    v38 -= (v38 >> ((4294967295) & 63))
                                                    var v81 float32 = (v66 - v66)
                                                    var v82 int16 = v38
                                                    var v83 uint32 = 914
                                                    if false {
                                                        if true {
                                                            var v84 int = fi48
                                                            fmt.Println(v84, v83, v38)
                                                            _ = v84
                                                        } else {
                                                            var v85 int32 = int32(v37)
                                                            _ = v85
                                                        }
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                    _ = v80
                                                    _ = v81
                                                    _ = v82
                                                    _ = v83
                                                }
                                                v66 = (v66 + 0.1)
                                                _ = v66
                                            } else {
                                                fmt.Println(fi62)
                                            }
                                        }
                                    case "":
                                        fmt.Println(fi62, fi62, v2)
                                    case "\n\t":
                                        if true {
                                            var v86 uint64 = 42
                                            var v87 int = fi62
                                            _ = v86
                                            _ = v87
                                        }
                                        fmt.Println(v23)
                                        var v88 float32 = (((-2.25 / 3.14159) + 1e10) * 0.1)
                                        _ = v88
                                    }
                                default:
                                    if (1.5 >= -2.25) {
                                        fmt.Println(fi48, v2, fi62)
                                        v89 := GenS9{F0: 1}
                                        if false {
                                            switch fi48 {
                                            case -2147483648:
                                                _ = len(v3)
                                            case -849:
                                                if true {
                                                    v90 := GenS2{F0: 0, F1: 42}
                                                    _ = len(v3)
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v91 uint16 = (v23 + v23)
                                                        _ = v91
                                                    case "foo bar":
                                                        fmt.Println(v37, v38, v1)
                                                    case "":
                                                        v38 = v38
                                                    default:
                                                        var v92 string = ("a" + "\n\t")
                                                        _ = v92
                                                    }
                                                    v1 -= 2
                                                    fmt.Println(v2, v2, v37)
                                                    if false {
                                                        var v93 uint32 = 850
                                                        _ = v93
                                                    } else {
                                                        v94 := GenS1{F0: (^(v2))}
                                                        v3[(fi62 - -9223372036854775808)] = ((v23 >> ((42) & 63)) >= (v23 % ((v23) | 1)))
                                                        _ = v94
                                                    }
                                                    fmt.Println(fi48, v37)
                                                    _ = v90
                                                } else {
                                                    v37 = v2
                                                }
                                            default:
                                                v63 = append(v63, fi48)
                                            }
                                            v95 := GenS10{F0: 921}
                                            var v96 string = (((("hello" + "") + "foo bar") + "a") + "a")
                                            var v97 int8 = 100
                                            v98 := GenS7{F0: 2, F1: v38, F2: v38}
                                            fmt.Println(v37, v37)
                                            switch (fi62 | fi48) {
                                            case 0:
                                                v96 = v96
                                            case 4294967295:
                                                if true {
                                                    var v99 int16 = (v38 >> ((255) & 63))
                                                    _ = v99
                                                } else {
                                                    var v100 string = v96
                                                    if true {
                                                        fmt.Println(v96, v100, v23)
                                                        if (255 >= (v23 + v23)) {
                                                            delete(v3, 0)
                                                            _ = len(v3)
                                                            v101 := GenS8{F0: 770, F1: v37}
                                                            _ = v101
                                                        }
                                                        switch v100 {
                                                        case "a":
                                                            fmt.Println(v1, v96, v2)
                                                        case "hello":
                                                            v102 := GenS11{F0: 42}
                                                            _ = v102
                                                        case "\n\t":
                                                            var v103 int32 = -884
                                                            _ = v103
                                                        default:
                                                            v2 = -9223372036854775808
                                                        }
                                                        fmt.Println(v38, v37)
                                                    } else {
                                                        switch (v100 + v96) {
                                                        case "":
                                                            v96 = "\n\t"
                                                        default:
                                                            v104 := GenS12{F0: int(v2), F1: -1, F2: 1}
                                                            _ = v104
                                                        }
                                                        var v105 uint64 = 55
                                                        _ = v105
                                                    }
                                                    v63 = append(v63, int(v2))
                                                    _ = v100
                                                }
                                            default:
                                                v106 := GenS10{F0: 564}
                                                _ = v106
                                            }
                                            var v107 float64 = ((1.5 + 0.1) / 1.0)
                                            v23 = 108
                                            _ = v107
                                            _ = v95
                                            _ = v96
                                            _ = v97
                                            _ = v98
                                        } else {
                                            var v108 string = ""
                                            fmt.Println(fi48, v1, v23)
                                            _ = v108
                                        }
                                        _ = v89
                                    } else {
                                        var v109 uint64 = 847
                                        v37 += v37
                                        v109 += (v109 - v109)
                                        _ = v109
                                    }
                                }
                                _ = fi62
                                _ = v63
                            }
                        }
                        _ = fi48
                    }
                    v38 -= -32768
                }
            case 42:
                v110 := GenS5{F0: v23, F1: ("a" + ("a" + ("" + ""))), F2: 65535}
                _ = v110
            }
            if false {
                if ((1.5 <= 1.5) && (42 > 0)) {
                    if true {
                        if (!(-1 >= 127)) {
                            switch ("foo bar" + ("hello" + "")) {
                            case "":
                                var v111 uint = uint(v23)
                                _ = v111
                            case "hello":
                                var v112 uint16 = v23
                                _ = v112
                            default:
                                v37 -= v37
                            }
                        } else {
                            var v113 string = (("a" + "") + "a")
                            switch (v113 + (v113 + "foo bar")) {
                            case "\n\t":
                                if (v37 == (-(v37))) {
                                    v1 = (v1 - 255)
                                    v38 += v38
                                    var v114 bool = (false && true)
                                    _ = v114
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        v115 := GenS12{F0: 4294967295, F1: -2147483648, F2: 4294967295}
                                        _ = v115
                                    default:
                                        if false {
                                            v1 = v1
                                            if (!(false && true)) {
                                                v113 = (v113 + "foo bar")
                                            } else {
                                                for fi116 := 0; fi116 < 11; fi116++ {
                                                    v117 := []int{fi116, (fi116 * -1), (fi116 - fi116)}
                                                    fmt.Println(v37, fi116)
                                                    if ((-(v38)) >= (v38 | 0)) {
                                                        var v118 int8 = 0
                                                        switch 2147483647 {
                                                        case 133:
                                                            switch ((v113 + "\n\t") + "foo bar") {
                                                            case "hello":
                                                                switch "a" {
                                                                case "\n\t":
                                                                    for fi119 := 0; fi119 < 3; fi119++ {
                                                                        v37 += v2
                                                                        _ = fi119
                                                                    }
                                                                default:
                                                                    v37 = 0
                                                                }
                                                            }
                                                        case -673:
                                                            _ = len(v3)
                                                        case -922:
                                                            if true {
                                                                v113 = v113
                                                            }
                                                        }
                                                        _ = v118
                                                    }
                                                    _ = fi116
                                                    _ = v117
                                                }
                                            }
                                            v120 := GenS4{F0: v113}
                                            _ = v120
                                        }
                                    }
                                }
                            case "foo bar":
                                v113 = v113
                            case "hello":
                                if true {
                                    v121 := []int{4294967295, -2147483648, 32768, -1}
                                    _ = v121
                                } else {
                                    var v122 uint = uint(v23)
                                    var v123 string = (v113 + v113)
                                    _ = v122
                                    _ = v123
                                }
                            }
                            v124 := GenS13{F0: -2147483648, F1: int32(v2)}
                            _ = v113
                            _ = v124
                        }
                        var v125 bool = false
                        v23 = v23
                        _ = v125
                    }
                } else {
                    var v126 float32 = 1e10
                    _ = v126
                }
                v38 += v38
                if true {
                    v127 := v3[648]
                    _ = v127
                } else {
                    fmt.Println(v2)
                    var v128 bool = (4294967295 < 1)
                    _ = v128
                }
            } else {
                v129 := GenS14{F0: (v2 >> ((240) & 63)), F1: (v23 << ((18446744073709551615) & 63))}
                _ = v129
            }
        }
    case 29:
        for fi130 := 0; fi130 < 6; fi130++ {
            for fi131 := 0; fi131 < 4; fi131++ {
                v1 = (v1 >> ((4294967295) & 63))
                _ = fi131
            }
            _ = fi130
        }
    default:
        v37 -= int64(v23)
    }
    var v132 bool = true
    v133 := GenS5{F0: v23, F1: "hello", F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v132:", v132)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v132
    _ = v133
    _ = v2
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v40
}
