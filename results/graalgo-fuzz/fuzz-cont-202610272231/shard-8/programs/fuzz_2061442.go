package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int64 }
type GenS2 struct { F0 int64; F1 int64; F2 uint32 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int16; F1 int16; F2 int64 }
type GenS6 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS7 struct { F0 uint64; F1 int64; F2 int64 }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 string; F1 string }
type GenS13 struct { F0 string; F1 int64; F2 uint64 }
type GenS14 struct { F0 uint32; F1 uint64 }
type GenS15 struct { F0 uint64; F1 uint64; F2 int }
func genFn0(p0 int, p1 int64, p2 string) string {
    for fi6 := 0; fi6 < 14; fi6++ {
        if true {
            if false {
                p0 = p0
                for fi7 := 0; fi7 < 18; fi7++ {
                    var v8 int32 = -2147483648
                    switch ("hello" + "a") {
                    case "a":
                        var v9 int32 = int32(fi7)
                        _ = v9
                    case "foo bar":
                        p2 = p2
                    case "":
                        switch p0 {
                        case -571:
                            var v10 float32 = -2.25
                            _ = v10
                        case 337:
                            fmt.Println(p1, p0)
                        }
                    }
                    v11 := []int{42}
                    _ = fi7
                    _ = v11
                    _ = v8
                }
                p1 -= 0
            } else {
                if true {
                    fmt.Println(p0)
                    for fi12 := 0; fi12 < 3; fi12++ {
                        v13 := GenS1{F0: 310, F1: fi6, F2: p1}
                        v14 := []int{(-(p0)), 0, (fi12 / ((p0) | 1))}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    v15 := []int{(-(fi6)), p0, 42}
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        p1 += (p1 / (((p1 << ((2) & 63))) | 1))
                        v17 := []int{(fi6 & 127)}
                        if (-366 <= 0) {
                            p0 = (fi16 | p0)
                            switch (fi16 << ((75) & 63)) {
                            case -2147483648:
                                var v18 int16 = 895
                                _ = v18
                            case 889:
                                var v19 int = (p0 & (^(fi16)))
                                _ = v19
                            default:
                                v20 := GenS1{F0: 18446744073709551615, F1: p0, F2: 1}
                                _ = v20
                            }
                            switch p2 {
                            case "foo bar":
                                fmt.Println(p0, p0, p2)
                            default:
                                v21 := GenS1{F0: 1, F1: fi6, F2: p1}
                                _ = v21
                            }
                        }
                        _ = fi16
                        _ = v17
                    }
                }
            }
            if true {
                v22 := GenS2{F0: (p1 ^ p1), F1: (^(p1)), F2: 235}
                _ = v22
            }
            if true {
                switch "" {
                case "":
                    p2 = p2
                case "\n\t":
                    var v23 uint16 = 733
                    _ = v23
                case "a":
                    switch p2 {
                    case "a":
                        var v24 int32 = 1
                        _ = v24
                    case "\n\t":
                        fmt.Println(fi6, fi6, p1)
                    }
                }
                var v25 int = fi6
                fmt.Println(p1, p0, v25)
                _ = v25
            } else {
                v26 := []int{fi6, 117, (fi6 / (((p0 << ((18446744073709551615) & 63))) | 1)), (fi6 * (p0 | fi6))}
                p2 = p2
                _ = v26
            }
        } else {
            if (1 != 924) {
                var v27 int = 129
                v28 := []int{p0, 4294967295}
                fmt.Println(p1, p2, fi6)
                _ = v27
                _ = v28
            }
        }
        _ = fi6
    }
    if false {
        if true {
            if (1 >= 163) {
                fmt.Println(p0, p1, p1)
                var v29 float32 = float32(((((0.0 * 0.1) / 3.14159) + 1.5) - 3.14159))
                _ = v29
            } else {
                v30 := GenS2{F0: (p1 - 951), F1: (p1 >> ((2) & 63)), F2: 2147483647}
                _ = v30
            }
        }
        p2 = ("foo bar" + "a")
    }
    var v31 int = (p0 << ((0) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    return p2
}

func main() {
    var v1 bool = (926 == 1)
    var v2 uint = 18446744073709551615
    v3 := []uint32{}
    var v4 bool = (((4294967295 > 405) && (-839 < 0)) || (-365 <= 2147483647))
    v5 := []int{4294967295, -1, 42, -2147483648}
    v4 = v1
    v32 := GenS3{F0: (("" + "\n\t") + ("" + "\n\t")), F1: -2147483648}
    switch "hello" {
    case "\n\t":
        v33 := GenS4{F0: 42, F1: 803}
        _ = v33
    case "a":
        var v34 float32 = 1e10
        _ = v34
    case "foo bar":
        if (true && ((uint8(v2) <= 180) || v4)) {
            v1 = (1e10 != 1.5)
        } else {
            _ = len(v5)
            var v35 uint64 = 42
            var v36 int16 = 1
            _ = v35
            _ = v36
        }
    }
    v37 := GenS5{F0: -979, F1: 32767, F2: -288}
    switch 32767 {
    case 315:
        v4 = v4
    case 32767:
        fmt.Println(v4, v1)
    case 32768:
        if ("a" <= (("hello" + "a") + "\n\t")) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (1e10 > 0.1) {
                for fi38 := 0; fi38 < 1; fi38++ {
                    var v39 bool = (377 != 42)
                    _ = fi38
                    _ = v39
                }
            } else {
                var v40 uint = (v2 << ((44) & 63))
                _ = v40
            }
        } else {
            v41 := GenS2{F0: -558, F1: 42, F2: uint32(v2)}
            var v42 uint = 850
            _ = v41
            _ = v42
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if v1 {
        v43 := GenS5{F0: 448, F1: -1, F2: -9223372036854775808}
        if false {
            var v44 string = "foo bar"
            var v45 float32 = -2.25
            for fi46 := 0; fi46 < 14; fi46++ {
                switch v44 {
                case "a":
                    v47 := GenS4{F0: 635, F1: (^(v2))}
                    _ = v47
                case "\n\t":
                    var v48 uint = (v2 + 366)
                    _ = v48
                default:
                    var v49 uint16 = 397
                    _ = v49
                }
                _ = fi46
            }
            _ = v44
            _ = v45
        }
        v50 := GenS2{F0: -1, F1: -1, F2: 0}
        _ = v43
        _ = v50
    }
    switch "hello" {
    case "\n\t":
        v51 := GenS6{F0: 32767, F1: 0, F2: 0}
        _ = v51
    case "hello":
        if len(v3) > 0 {
            v3[0] = 988
        }
    default:
        if true {
            var v52 bool = v1
            v53 := GenS7{F0: 42, F1: 0, F2: 796}
            switch ("hello" + "a") {
            case "foo bar":
                for fi54 := 0; fi54 < 4; fi54++ {
                    v55 := GenS7{F0: 2, F1: -1, F2: 42}
                    _ = fi54
                    _ = v55
                }
            }
            _ = v52
            _ = v53
        } else {
            switch ("\n\t" + ((("foo bar" + "hello") + ("" + "")) + "hello")) {
            case "foo bar":
                fmt.Println(v1, v1, v4)
            case "hello":
                v4 = (v4 || false)
            case "\n\t":
                var v56 int64 = -52
                _ = v56
            default:
                v4 = v1
            }
            v4 = (!(false && v4))
        }
    }
    var v57 uint8 = 65
    var v58 uint8 = 127
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    if v4 {
        switch -2147483648 {
        case 42:
            fmt.Println(v58, v58, v1)
        case 32768:
            if ("hello" >= "") {
                var v59 int64 = 1
                switch 620 {
                case 127:
                    v4 = v1
                default:
                    var v60 int16 = -922
                    _ = v60
                }
                if (!false) {
                    if v1 {
                        if len(v5) > 0 {
                            v5[0] = 128
                        }
                    }
                    fmt.Println(v59, v57)
                    v3 = append(v3, 1)
                }
                _ = v59
            }
        }
    }
    if v1 {
        switch 4294967295 {
        case -506:
            var v61 string = "hello"
            _ = v61
        case 32768:
            for fi62 := 0; fi62 < 5; fi62++ {
                for fi63 := 0; fi63 < 18; fi63++ {
                    if (810 < -32768) {
                        fmt.Println(fi62)
                        if true {
                            if (int8(fi63) < 100) {
                                var v64 float32 = ((1e10 + 0.0) * float32(float64(-2.25)))
                                var v65 string = ("hello" + "\n\t")
                                _ = v64
                                _ = v65
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = 2
                                }
                                fmt.Println(fi63)
                                var v66 float32 = ((((float32((1.5 * 1.5)) / 1.5) / ((0.1 - -2.25) * float32(3.14159))) / ((0.0 / float32(((((1.5 + 1e10) / (-2.25 / (-2.25 + 1.5))) + (0.0 + float64(1e10))) + 3.14159))) + 1e10)) + (1.5 / 1e10))
                                v4 = v4
                                var v67 uint64 = 0
                                _ = v66
                                _ = v67
                            }
                            var v68 bool = (("foo bar" + "hello") < ("a" + "foo bar"))
                            switch -9223372036854775808 {
                            case 2147483647:
                                var v69 int16 = 42
                                _ = v69
                            case 127:
                                v70 := GenS8{F0: 277, F1: (^(fi62))}
                                _ = v70
                            case -9223372036854775808:
                                var v71 string = "foo bar"
                                var v72 bool = false
                                fmt.Println(v72, v72)
                                _ = v71
                                _ = v72
                            default:
                                if (-57 >= -53) {
                                    if v1 {
                                        switch "" {
                                        case "":
                                            v5 = append(v5, fi62)
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        fmt.Println(v57)
                                    }
                                } else {
                                    v73 := GenS1{F0: 606, F1: fi63, F2: -1}
                                    if v1 {
                                        fmt.Println(v57)
                                    } else {
                                        var v74 uint32 = 0
                                        v3 = append(v3, v74)
                                        _ = v74
                                    }
                                    switch ("foo bar" + "") {
                                    case "foo bar":
                                        if (!false) {
                                            switch (^(fi63)) {
                                            case 128:
                                                fmt.Println(v58)
                                            case 32767:
                                                v75 := GenS9{F0: ((v57 << ((uint(v57)) & 63)) <= v58)}
                                                _ = v75
                                            default:
                                                var v76 int = fi62
                                                _ = v76
                                            }
                                            v77 := GenS10{F0: -307}
                                            _ = v77
                                        } else {
                                            v68 = ("\n\t" >= "")
                                            switch "" {
                                            case "\n\t":
                                                fmt.Println(fi62, fi62, v2)
                                            case "hello":
                                                v2 -= v2
                                            case "foo bar":
                                                if false {
                                                    v68 = (!true)
                                                    v2 += 994
                                                    if true {
                                                        v78 := GenS11{F0: (^(v57))}
                                                        if len(v3) > 0 {
                                                            v3[0] = 1
                                                        }
                                                        _ = v78
                                                    }
                                                } else {
                                                    v1 = (127 <= 60)
                                                    _ = len(v5)
                                                }
                                            }
                                            if v1 {
                                                v2 = 4294967295
                                                v79 := GenS12{F0: "\n\t", F1: ""}
                                                _ = v79
                                            } else {
                                                _ = len(v3)
                                                var v80 uint32 = 518
                                                if (0 < 244) {
                                                    fmt.Println(fi63)
                                                    switch (((("" + "\n\t") + "\n\t") + (((("foo bar" + ("hello" + "foo bar")) + "\n\t") + "a") + "hello")) + ("a" + "hello")) {
                                                    case "foo bar":
                                                        if false {
                                                            var v81 float64 = 3.14159
                                                            _ = len(v5)
                                                            _ = v81
                                                        } else {
                                                            fmt.Println(v58, fi62, v2)
                                                        }
                                                    case "\n\t":
                                                        var v82 int8 = 127
                                                        _ = v82
                                                    case "a":
                                                        switch fi63 {
                                                        case 638:
                                                            v80 = (v80 << ((v2) & 63))
                                                        default:
                                                            v83 := GenS4{F0: 1, F1: v2}
                                                            _ = v83
                                                        }
                                                    }
                                                    v1 = false
                                                } else {
                                                    var v84 uint8 = v58
                                                    switch "" {
                                                    case "":
                                                        if v68 {
                                                            if v1 {
                                                                var v85 float32 = ((float32(((3.14159 + (-2.25 + ((-2.25 - float64((float32(1.5) + (1e10 * -2.25)))) * (float64((1e10 / float32(1e10))) - 1.5)))) - (float64((1.5 * 1e10)) * 0.0))) - 0.0) * -2.25)
                                                                v84 += v57
                                                                _ = v85
                                                            }
                                                            fmt.Println(v2, v80)
                                                            fmt.Println(v58, fi63, v80)
                                                            if ((v80 << ((v2) & 63)) <= 2147483648) {
                                                                var v86 uint = (v2 + (v2 % ((v2) | 1)))
                                                                if len(v3) > 0 {
                                                                    v3[0] = (^(v80))
                                                                }
                                                                if false {
                                                                    v68 = false
                                                                } else {
                                                                    v87 := GenS6{F0: -1, F1: 631, F2: 255}
                                                                    _ = v87
                                                                }
                                                                _ = v86
                                                            }
                                                            _ = len(v5)
                                                        }
                                                        v2 += v2
                                                        switch (fi63 * (fi62 ^ (fi62 % (((-(fi62))) | 1)))) {
                                                        case 127:
                                                            if v1 {
                                                                if v4 {
                                                                    var v88 int16 = 32767
                                                                    switch 2147483647 {
                                                                    case 2147483647:
                                                                        v57 -= 225
                                                                    }
                                                                    var v89 int8 = 42
                                                                    _ = v88
                                                                    _ = v89
                                                                } else {
                                                                    fmt.Println(v80, v80)
                                                                    var v90 float32 = 1.5
                                                                    v91 := GenS10{F0: -142}
                                                                    _ = v90
                                                                    _ = v91
                                                                }
                                                                v84 += 128
                                                            }
                                                        default:
                                                            switch "hello" {
                                                            case "foo bar":
                                                                switch "" {
                                                                case "hello":
                                                                    if (("\n\t" + (("hello" + ("\n\t" + "hello")) + "foo bar")) >= "") {
                                                                        var v92 uint32 = (v80 + v80)
                                                                        v92 += v80
                                                                        _ = v92
                                                                    }
                                                                    var v93 uint = (v2 - v2)
                                                                    v94 := GenS13{F0: ("" + "a"), F1: -177, F2: 170}
                                                                    _ = v93
                                                                    _ = v94
                                                                case "":
                                                                    v95 := GenS8{F0: 699, F1: (fi62 + fi63)}
                                                                    _ = v95
                                                                default:
                                                                    v80 += (v80 << ((v2) & 63))
                                                                }
                                                            default:
                                                                var v96 uint32 = (v80 + 693)
                                                                _ = v96
                                                            }
                                                        }
                                                    case "foo bar":
                                                        if (true && true) {
                                                            var v97 string = "foo bar"
                                                            fmt.Println(v2, v58)
                                                            _ = v97
                                                        }
                                                    default:
                                                        v3 = append(v3, v80)
                                                    }
                                                    _ = v84
                                                }
                                                var v98 uint16 = 1
                                                _ = v80
                                                _ = v98
                                            }
                                        }
                                    default:
                                        v5 = append(v5, fi63)
                                    }
                                    _ = v73
                                }
                            }
                            _ = v68
                        }
                        if (v1 || true) {
                            var v99 uint64 = 578
                            v1 = false
                            _ = v99
                        } else {
                            var v100 int16 = -1
                            _ = v100
                        }
                    } else {
                        switch 9223372036854775807 {
                        case 42:
                            switch (^(fi62)) {
                            case 1:
                                switch 369 {
                                case 42:
                                    v57 = (v58 * (v58 >> ((4294967295) & 63)))
                                case -1:
                                    if len(v3) > 0 {
                                        v3[0] = 65535
                                    }
                                case 1:
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                default:
                                    if (v4 && v4) {
                                        var v101 uint32 = 0
                                        var v102 int16 = -271
                                        _ = v101
                                        _ = v102
                                    }
                                }
                            default:
                                _ = len(v3)
                            }
                        case 32768:
                            switch (fi62 + fi62) {
                            case 127:
                                v58 += v57
                            case 454:
                                switch fi63 {
                                case 4294967295:
                                    if false {
                                        v5 = append(v5, (fi62 & 32767))
                                    } else {
                                        var v103 uint64 = 1
                                        if (1.5 >= float64(float32(1.5))) {
                                            if (0.1 < 0.0) {
                                                v104 := GenS2{F0: 42, F1: -775, F2: 4294967295}
                                                v105 := GenS10{F0: 751}
                                                var v106 uint32 = 2147483647
                                                _ = v104
                                                _ = v105
                                                _ = v106
                                            }
                                            switch (fi63 - (-(fi63))) {
                                            case -9223372036854775808:
                                                var v107 float64 = 0.1
                                                _ = v107
                                            case 32767:
                                                var v108 int16 = 42
                                                _ = v108
                                            }
                                            var v109 int32 = 1
                                            _ = v109
                                        }
                                        if true {
                                            if true {
                                                v2 = (v2 << ((v2) & 63))
                                                if false {
                                                    v103 = (v103 ^ v103)
                                                    fmt.Println(v57)
                                                    switch (fi63 ^ fi63) {
                                                    case -2147483648:
                                                        v5 = append(v5, (fi63 * fi63))
                                                    case -9223372036854775808:
                                                        var v110 uint16 = 0
                                                        _ = v110
                                                    }
                                                    v111 := GenS14{F0: 995, F1: v103}
                                                    _ = v111
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                        }
                                        _ = v103
                                    }
                                    v4 = false
                                    var v112 string = "foo bar"
                                    _ = v112
                                case 847:
                                    if v1 {
                                        switch (("" + "foo bar") + "hello") {
                                        case "a":
                                            if len(v3) > 0 {
                                                v3[0] = 1
                                            }
                                        case "foo bar":
                                            _ = len(v3)
                                        case "\n\t":
                                            v4 = false
                                        }
                                    } else {
                                        v2 += v2
                                        if ((4294967295 <= 913) && (18446744073709551615 < 2)) {
                                            fmt.Println(v58)
                                            v2 += v2
                                            var v113 int8 = -1
                                            _ = v113
                                        }
                                    }
                                case -9223372036854775808:
                                    v4 = v4
                                }
                            }
                        }
                        v114 := GenS5{F0: 1000, F1: 1, F2: 9223372036854775807}
                        _ = v114
                    }
                    v115 := GenS15{F0: 2, F1: 703, F2: (fi62 & fi62)}
                    _ = fi63
                    _ = v115
                }
                fmt.Println(v57)
                for fi116 := 0; fi116 < 13; fi116++ {
                    var v117 uint8 = v58
                    if (!true) {
                        fmt.Println(v117, v117)
                        switch fi116 {
                        case 1:
                            var v118 uint64 = 4294967295
                            _ = v118
                        case 0:
                            var v119 int = (fi116 + fi116)
                            _ = v119
                        default:
                            fmt.Println(v117, fi62, fi116)
                        }
                    }
                    _ = fi116
                    _ = v117
                }
                _ = fi62
            }
        case -927:
            for fi120 := 0; fi120 < 9; fi120++ {
                for fi121 := 0; fi121 < 10; fi121++ {
                    v122 := GenS8{F0: 2, F1: -568}
                    v1 = (((0.1 - float32((-2.25 / 1.5))) < 1e10) && v4)
                    var v123 int16 = 1000
                    _ = fi121
                    _ = v122
                    _ = v123
                }
                var v124 int64 = 42
                _ = fi120
                _ = v124
            }
        }
        v2 += (v2 / ((0) | 1))
        fmt.Println(v2, v1)
    } else {
        switch 399 {
        case 2147483647:
            if (100 <= 84) {
                var v125 uint8 = 15
                if v1 {
                    if (((3.14159 - float64(3.14159)) <= 1e10) && v4) {
                        v4 = (0.0 < (-2.25 + ((1.5 / (-2.25 + float32(3.14159))) + (float32(3.14159) + (0.0 + float32(1e10))))))
                    }
                } else {
                    _ = len(v5)
                    v3 = append(v3, 2147483648)
                    var v126 int8 = 127
                    _ = v126
                }
                _ = v125
            }
        case 42:
            var v127 float64 = 0.0
            _ = v127
        }
        v58 += (v57 >> (((^(v2))) & 63))
        var v128 int16 = 592
        _ = v128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v57
    _ = v58
}
