package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS7 struct { F0 uint; F1 uint8; F2 int }
type GenS8 struct { F0 int16; F1 int }
type GenS9 struct { F0 int16; F1 string; F2 bool }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 string; F1 uint8 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int64; F1 int16 }
type GenS15 struct { F0 int32; F1 uint16 }
type GenS16 struct { F0 int64; F1 string; F2 uint32 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 uint8; F1 int }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 int8; F1 uint }
func genFn0(p0 uint, p1 int16) string {
    var v7 float32 = 0.1
    var v8 int64 = 1
    v8 = v8
    fmt.Println(p1, v8)
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := []int{-1, (fi9 << ((p0) & 63)), (fi9 % (((-(fi9))) | 1)), fi9}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return "a"
}
func genFn1() uint64 {
    var v11 bool = (false || (293 > -9223372036854775808))
    v12 := GenS1{F0: false, F1: 1}
    for fi13 := 0; fi13 < 5; fi13++ {
        v14 := GenS1{F0: v11, F1: 42}
        switch "hello" {
        case "":
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := GenS2{F0: 0}
                fmt.Println(fi13)
                _ = fi15
                _ = v16
            }
        case "a":
            switch ("" + "\n\t") {
            case "\n\t":
                v17 := GenS3{F0: 0, F1: 361}
                _ = v17
            }
        }
        _ = fi13
        _ = v14
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        switch 2147483647 {
        case 128:
            var v19 uint16 = 757
            _ = v19
        case 2147483647:
            switch ("foo bar" + "\n\t") {
            case "a":
                var v20 uint8 = 75
                _ = v20
            case "\n\t":
                switch fi18 {
                case 1:
                    v21 := []int{fi18, -170, (fi18 << ((457) & 63))}
                    _ = v21
                case 4294967295:
                    fmt.Println(fi18, fi18)
                }
            case "foo bar":
                for fi22 := 0; fi22 < 20; fi22++ {
                    v11 = v11
                    var v23 int16 = -817
                    var v24 int8 = 0
                    v25 := GenS1{F0: ((678 <= 255) && v11), F1: 81}
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            }
        case 9223372036854775807:
            v26 := GenS4{F0: (0 <= 42)}
            _ = v26
        }
        _ = fi18
    }
    var v27 int = 282
    _ = v11
    _ = v12
    _ = v27
    return 4294967295
}

func main() {
    var v1 int32 = -1
    var v2 int8 = 1
    var v3 uint8 = 116
    var v4 uint = 652
    var v5 int32 = 663
    v6 := make(map[int]uint32)
    if false {
        if ((v2 * (^(v2))) > 17) {
            v4 += v4
            v1 += (v5 % ((v1) | 1))
        } else {
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 uint64 = 555
                for fi30 := 0; fi30 < 13; fi30++ {
                    switch "a" {
                    case "foo bar":
                        v1 += v5
                    case "\n\t":
                        if true {
                            var v31 int16 = -295
                            _ = v31
                        } else {
                            v32 := GenS5{F0: 65535, F1: v1}
                            var v33 float32 = 1e10
                            v34 := GenS5{F0: 651, F1: (v1 & v5)}
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    case "hello":
                        v35 := []int{int(v2), fi30, 633}
                        _ = v35
                    }
                    v36 := v6[(fi30 ^ -1)]
                    _ = fi30
                    _ = v36
                }
                _ = fi28
                _ = v29
            }
            var v37 float64 = 1.5
            if (true || (-93 <= 1000)) {
                v6[42] = 668
            } else {
                v4 += v4
                v37 += 1.5
                for fi38 := 0; fi38 < 2; fi38++ {
                    switch ("hello" + "") {
                    case "foo bar":
                        var v39 string = "foo bar"
                        _ = v39
                    case "":
                        if true {
                            if (false && (v4 != v4)) {
                                v4 = (v4 - v4)
                                switch fi38 {
                                case 128:
                                    v40 := []int{32767, fi38}
                                    _ = v40
                                case 42:
                                    v4 = 541
                                case 0:
                                    if (v5 == -1) {
                                        var v41 string = "foo bar"
                                        v41 = ("foo bar" + (("hello" + (v41 + (v41 + v41))) + v41))
                                        fmt.Println(v2, v5)
                                        _ = v41
                                    }
                                }
                            } else {
                                v37 = (v37 - (v37 + ((((v37 + v37) + (float64(3.14159) / v37)) / v37) - (1e10 * 1.5))))
                                if true {
                                    var v42 float32 = (1e10 * (((3.14159 * 0.0) * -2.25) + -2.25))
                                    _ = v42
                                } else {
                                    v43 := GenS3{F0: v2, F1: v4}
                                    switch 17 {
                                    case 4294967295:
                                        switch int(v3) {
                                        case -9223372036854775808:
                                            var v44 uint = (v4 / (((v4 - (v4 / (((v4 + (v4 - (v4 + 2)))) | 1)))) | 1))
                                            _ = v44
                                        case 127:
                                            fmt.Println(v1, v37, v3)
                                        default:
                                            v45 := GenS1{F0: false, F1: (v3 | v3)}
                                            _ = v45
                                        }
                                    }
                                    _ = v43
                                }
                                v46 := []int{(^(fi38)), fi38, fi38}
                                _ = v46
                            }
                            var v47 int8 = (-(v2))
                            _ = v47
                        }
                    case "hello":
                        switch "\n\t" {
                        case "hello":
                            var v48 uint64 = 42
                            _ = v48
                        default:
                            v6[(-(fi38))] = 65535
                        }
                    }
                    v49 := GenS6{F0: 681, F1: 42, F2: (v5 << ((v4) & 63))}
                    var v50 bool = true
                    _ = fi38
                    _ = v49
                    _ = v50
                }
            }
            _ = v37
        }
        var v51 float64 = 1e10
        for fi52 := 0; fi52 < 4; fi52++ {
            var v53 uint16 = 2
            var v54 int32 = (v5 << (((v4 & (v4 * (v4 >> ((v4) & 63))))) & 63))
            v4 += v4
            _ = fi52
            _ = v53
            _ = v54
        }
        _ = v51
    }
    v1 = 42
    if (255 <= 313) {
        for fi55 := 0; fi55 < 1; fi55++ {
            var v56 int32 = v1
            v3 = (^(v3))
            _ = fi55
            _ = v56
        }
        if true {
            if false {
                for fi57 := 0; fi57 < 6; fi57++ {
                    if true {
                        var v58 bool = true
                        _ = v58
                    }
                    var v59 uint8 = 1
                    _ = fi57
                    _ = v59
                }
                v3 = v3
                v60 := GenS4{F0: true}
                _ = v60
            } else {
                var v61 int32 = v1
                v61 += (^(v5))
                var v62 bool = ((v61 + v1) == v1)
                _ = v61
                _ = v62
            }
            for fi63 := 0; fi63 < 18; fi63++ {
                v2 += (v2 + (v2 * 1))
                for fi64 := 0; fi64 < 15; fi64++ {
                    v65 := GenS6{F0: 4294967295, F1: 96, F2: (v5 >> ((255) & 63))}
                    _ = fi64
                    _ = v65
                }
                _ = fi63
            }
            v2 = (^(v2))
        }
        var v66 string = ("hello" + "")
        _ = v66
    } else {
        v4 += (^(v4))
        v2 = (v2 - (v2 & 100))
    }
    v67 := []int{-1, 32767, 32767, 999}
    v3 = 59
    var v68 string = "a"
    var v69 uint64 = 83
    if true {
        for fi70 := 0; fi70 < 4; fi70++ {
            for fi71 := 0; fi71 < 12; fi71++ {
                v72 := GenS4{F0: false}
                _ = len(v6)
                var v73 int = fi70
                var v74 string = ("hello" + (v68 + "a"))
                v5 += (v5 - 0)
                _ = fi71
                _ = v72
                _ = v73
                _ = v74
            }
            for fi75 := 0; fi75 < 19; fi75++ {
                if true {
                    v76 := GenS4{F0: (-849 >= -615)}
                    _ = v76
                }
                var v77 int64 = 975
                var v78 string = v68
                v79 := GenS1{F0: ((false || false) && true), F1: 159}
                switch (fi70 + (fi75 | fi70)) {
                case 814:
                    var v80 uint = v4
                    _ = v80
                case 739:
                    switch ((("\n\t" + ("hello" + v68)) + v78) + v68) {
                    case "a":
                        var v81 int32 = v5
                        _ = v81
                    case "foo bar":
                        v1 += -2147483648
                    default:
                        switch v78 {
                        case "foo bar":
                            var v82 float64 = 3.14159
                            v6[32767] = 2
                            if ((v69 & 2) == (v69 % (((v69 ^ genFn1())) | 1))) {
                                var v83 int64 = v77
                                fmt.Println(v3, v68)
                                if (!true) {
                                    v68 = v78
                                    switch "hello" {
                                    case "":
                                        fmt.Println(v5)
                                    }
                                }
                                _ = v83
                            } else {
                                v1 -= v5
                                switch v78 {
                                case "foo bar":
                                    var v84 uint = v4
                                    _ = v84
                                case "a":
                                    v85 := GenS7{F0: v4, F1: v3, F2: (fi75 & (fi75 & (fi70 << (((v4 ^ v4)) & 63))))}
                                    _ = v85
                                case "":
                                    v86 := GenS6{F0: 42, F1: 1, F2: 0}
                                    if true {
                                        var v87 float32 = (1e10 * 1e10)
                                        fmt.Println(v78, fi70, fi75)
                                        if false {
                                            var v88 float32 = v87
                                            _ = v88
                                        }
                                        _ = len(v6)
                                        _ = v87
                                    }
                                    var v89 int32 = 0
                                    var v90 uint16 = 0
                                    _ = v86
                                    _ = v89
                                    _ = v90
                                }
                                v77 -= (v77 | v77)
                            }
                            switch fi75 {
                            case 127:
                                switch fi70 {
                                case -534:
                                    if (((v69 << ((v4) & 63)) < 18446744073709551615) && (v77 <= v77)) {
                                        var v91 uint8 = (v3 % (((v3 + v3)) | 1))
                                        v92 := GenS8{F0: 42, F1: fi70}
                                        v77 = (v77 ^ v77)
                                        _ = v91
                                        _ = v92
                                    }
                                case 961:
                                    var v93 int32 = v1
                                    _ = v93
                                }
                            default:
                                if (3.14159 > ((1.5 + float32((1e10 / ((0.1 - v82) - v82)))) * (0.1 * (1.5 * 1e10)))) {
                                    if false {
                                        v94 := GenS5{F0: 65535, F1: (^(v5))}
                                        var v95 uint64 = (v69 | 2)
                                        _ = v94
                                        _ = v95
                                    } else {
                                        switch fi70 {
                                        case -115:
                                            v96 := v6[(fi75 / ((fi75) | 1))]
                                            switch v68 {
                                            case "foo bar":
                                                var v97 int16 = 32767
                                                _ = v97
                                            default:
                                                if len(v67) >= 2 {
                                                    v67 = v67[1:len(v67)]
                                                }
                                            }
                                            if true {
                                                switch (fi75 / (((fi70 / (((fi70 % ((fi70) | 1))) | 1))) | 1)) {
                                                case -9223372036854775808:
                                                    if false {
                                                        var v98 int32 = (v1 + -128)
                                                        _ = v98
                                                    } else {
                                                        fmt.Println(v4)
                                                        fmt.Println(v96, v68)
                                                    }
                                                case -179:
                                                    v77 = (v77 ^ 9223372036854775807)
                                                default:
                                                    switch (fi70 / ((289) | 1)) {
                                                    case 4294967295:
                                                        v99 := v6[(fi75 * (fi70 % (((fi70 ^ (fi75 - (fi75 + fi70)))) | 1)))]
                                                        _ = v99
                                                    case 127:
                                                        fmt.Println(v1, v2, v78)
                                                    case 612:
                                                        v100 := GenS5{F0: 42, F1: (v1 - 187)}
                                                        _ = v100
                                                    default:
                                                        v78 = (v78 + v78)
                                                    }
                                                }
                                                v101 := GenS9{F0: -32768, F1: ("a" + v78), F2: true}
                                                var v102 int = 287
                                                _ = v101
                                                _ = v102
                                            }
                                            _ = v96
                                        case 127:
                                            switch 42 {
                                            case 127:
                                                if false {
                                                    fmt.Println(v2, v3, v4)
                                                    delete(v6, (fi75 ^ (fi75 ^ (^(fi75)))))
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        fmt.Println(v82)
                                                    case "a":
                                                        if (!true) {
                                                            v1 -= 2147483647
                                                            fmt.Println(v82)
                                                            var v103 int16 = 1
                                                            _ = v103
                                                        }
                                                        var v104 float32 = 0.1
                                                        _ = v104
                                                    default:
                                                        fmt.Println(v2, fi75, v5)
                                                    }
                                                    switch 0 {
                                                    case 9223372036854775807:
                                                        v1 -= (v5 & 0)
                                                    case -2147483648:
                                                        if false {
                                                            v105 := GenS10{F0: uint16(v69)}
                                                            v2 = v2
                                                            _ = v105
                                                        } else {
                                                            v106 := GenS4{F0: (-1 >= fi70)}
                                                            _ = v106
                                                        }
                                                    }
                                                    delete(v6, int(v2))
                                                }
                                            case 9223372036854775807:
                                                v6[int(v1)] = 226
                                            case -86:
                                                v4 = uint(v3)
                                            default:
                                                switch int(v77) {
                                                case -262:
                                                    v107 := GenS8{F0: 1, F1: fi75}
                                                    _ = v107
                                                }
                                            }
                                        case 9223372036854775807:
                                            if (-2.25 != 1e10) {
                                                switch fi75 {
                                                case 32767:
                                                    v6[135] = 0
                                                case 1:
                                                    fmt.Println(v78, v3)
                                                case -1:
                                                    v108 := GenS11{F0: v1}
                                                    _ = v108
                                                }
                                                var v109 uint = uint(v3)
                                                _ = v109
                                            }
                                        }
                                        if len(v67) > 0 {
                                            v67[0] = int(v3)
                                        }
                                        switch (v78 + ("" + v78)) {
                                        case "\n\t":
                                            v110 := GenS12{F0: v68, F1: v3}
                                            _ = v110
                                        default:
                                            delete(v6, fi70)
                                        }
                                    }
                                    v6[(^(fi75))] = 1
                                } else {
                                    if true {
                                        v111 := GenS13{F0: 4294967295}
                                        var v112 int64 = v77
                                        _ = v111
                                        _ = v112
                                    }
                                    if true {
                                        v5 += v1
                                        fmt.Println(v5, v82, v78)
                                        v113 := GenS2{F0: fi75}
                                        _ = v113
                                    } else {
                                        v114 := v6[fi70]
                                        var v115 float32 = 0.1
                                        if true {
                                            switch (((v78 + (("\n\t" + "a") + v68)) + "foo bar") + (((v78 + "a") + (v68 + v68)) + ((v78 + ("\n\t" + (((v78 + v78) + v68) + v78))) + (((("" + (v78 + v68)) + (v78 + "a")) + v68) + v68)))) {
                                            case "\n\t":
                                                if true {
                                                    var v116 uint32 = v114
                                                    _ = v116
                                                }
                                            case "foo bar":
                                                if (true || (745 <= uint16(v4))) {
                                                    v115 += 1e10
                                                }
                                            case "a":
                                                v117 := GenS4{F0: true}
                                                _ = v117
                                            }
                                        } else {
                                            switch v68 {
                                            case "foo bar":
                                                v118 := GenS14{F0: v77, F1: -215}
                                                v119 := GenS14{F0: 1, F1: 0}
                                                _ = len(v67)
                                                _ = v118
                                                _ = v119
                                            case "\n\t":
                                                var v120 int16 = -581
                                                _ = v120
                                            default:
                                                v121 := GenS15{F0: v1, F1: 141}
                                                _ = v121
                                            }
                                        }
                                        _ = v114
                                        _ = v115
                                    }
                                    switch (fi75 * fi75) {
                                    case 128:
                                        fmt.Println(v3, fi70, v3)
                                    case 32768:
                                        switch "hello" {
                                        case "":
                                            var v122 int16 = 0
                                            if false {
                                                var v123 string = v68
                                                _ = v123
                                            }
                                            v82 -= (1e10 - v82)
                                            _ = v122
                                        case "hello":
                                            v6[(fi75 + int(v3))] = 0
                                        case "\n\t":
                                            switch v78 {
                                            case "hello":
                                                fmt.Println(v77, v82, v68)
                                            case "":
                                                v82 = 3.14159
                                            case "a":
                                                fmt.Println(v1)
                                            default:
                                                var v124 uint32 = 1
                                                _ = v124
                                            }
                                        default:
                                            v3 = v3
                                        }
                                    case 0:
                                        var v125 uint8 = v3
                                        _ = v125
                                    default:
                                        switch (v68 + "") {
                                        case "hello":
                                            switch "" {
                                            case "a":
                                                if (!true) {
                                                    switch "foo bar" {
                                                    case "":
                                                        if (0 != int16(v69)) {
                                                            fmt.Println(fi70, v68)
                                                            fmt.Println(fi70, v69, v69)
                                                            switch "\n\t" {
                                                            case "":
                                                                if (true && (false || true)) {
                                                                    if len(v67) >= 2 {
                                                                        v67 = v67[1:len(v67)]
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                var v126 int = (-(fi70))
                                                                _ = v126
                                                            case "hello":
                                                                fmt.Println(v82)
                                                            }
                                                        } else {
                                                            v4 -= v4
                                                            v127 := GenS16{F0: v77, F1: v68, F2: 2147483648}
                                                            _ = v127
                                                        }
                                                    case "a":
                                                        switch v78 {
                                                        case "hello":
                                                            if false {
                                                                var v128 int8 = -15
                                                                fmt.Println(v82, fi70, fi75)
                                                                fmt.Println(v77, fi70, v77)
                                                                _ = v128
                                                            }
                                                        case "\n\t":
                                                            if (false || (!false)) {
                                                                v77 = (v77 + (v77 >> (((v4 | (v4 / (((v4 | v4)) | 1)))) & 63)))
                                                            }
                                                        default:
                                                            if len(v67) > 0 {
                                                                v67[0] = 0
                                                            }
                                                            if true {
                                                                var v129 float32 = ((1.5 / float32(v82)) * (((0.0 * float32(0.1)) - 1e10) - 0.1))
                                                                _ = v129
                                                            } else {
                                                                var v130 uint32 = 286
                                                                v131 := GenS12{F0: v78, F1: v3}
                                                                _ = v130
                                                                _ = v131
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v68 = "a"
                                                    default:
                                                        _ = len(v6)
                                                    }
                                                } else {
                                                    switch 9223372036854775807 {
                                                    case 32767:
                                                        v82 = (float64(1e10) + v82)
                                                    case -1:
                                                        v132 := GenS16{F0: (v77 * 367), F1: v68, F2: 1}
                                                        _ = v132
                                                    case 42:
                                                        if (((v5 >> (((v4 % (((v4 - v4)) | 1))) & 63)) < (v1 << (((v4 ^ 55)) & 63))) && (42 < (v5 & v1))) {
                                                            fmt.Println(v1, v69)
                                                            fmt.Println(v2)
                                                            v133 := v6[fi70]
                                                            _ = v133
                                                        } else {
                                                            v78 = "\n\t"
                                                            if false {
                                                                switch 438 {
                                                                case 128:
                                                                    var v134 uint = v4
                                                                    _ = v134
                                                                case -9223372036854775808:
                                                                    var v135 bool = true
                                                                    _ = v135
                                                                default:
                                                                    var v136 string = v68
                                                                    _ = v136
                                                                }
                                                                if false {
                                                                    var v137 int8 = (v2 + v2)
                                                                    _ = v137
                                                                } else {
                                                                    var v138 uint = 618
                                                                    var v139 int16 = -3
                                                                    var v140 uint = (v138 >> ((v4) & 63))
                                                                    v141 := v6[fi75]
                                                                    switch (-(fi75)) {
                                                                    case 518:
                                                                        v67 = append(v67, 42)
                                                                    default:
                                                                        v142 := GenS17{F0: v141}
                                                                        _ = v142
                                                                    }
                                                                    _ = v138
                                                                    _ = v139
                                                                    _ = v140
                                                                    _ = v141
                                                                }
                                                            } else {
                                                                var v143 int = fi75
                                                                v144 := GenS18{F0: 189, F1: -2147483648}
                                                                _ = v143
                                                                _ = v144
                                                            }
                                                            v6[fi75] = 2147483647
                                                        }
                                                        if true {
                                                            v145 := GenS5{F0: 497, F1: v1}
                                                            var v146 int = fi75
                                                            _ = v145
                                                            _ = v146
                                                        }
                                                    }
                                                    v147 := GenS16{F0: (^(v77)), F1: v68, F2: 42}
                                                    if true {
                                                        var v148 int32 = v5
                                                        _ = len(v67)
                                                        v4 = (v4 ^ v4)
                                                        _ = v148
                                                    }
                                                    if (true && true) {
                                                        v149 := GenS1{F0: (v4 > uint(v3)), F1: v3}
                                                        _ = v149
                                                    }
                                                    _ = v147
                                                }
                                            case "":
                                                v1 = v1
                                            default:
                                                if true {
                                                    v150 := v6[fi75]
                                                    _ = v150
                                                }
                                            }
                                        case "foo bar":
                                            if (true || true) {
                                                var v151 uint16 = 362
                                                var v152 int16 = -827
                                                var v153 float64 = 0.1
                                                if len(v67) >= 2 {
                                                    v67 = v67[1:len(v67)]
                                                }
                                                _ = v151
                                                _ = v152
                                                _ = v153
                                            }
                                            v2 = v2
                                            var v154 int = fi70
                                            _ = v154
                                        }
                                    }
                                }
                                if (!(447 < 649)) {
                                    v155 := GenS19{F0: (v4 << (((v4 % ((42) | 1))) & 63))}
                                    v156 := v6[fi70]
                                    _ = v155
                                    _ = v156
                                } else {
                                    delete(v6, fi75)
                                    v69 += v69
                                    var v157 int64 = v77
                                    _ = v157
                                }
                                v2 = v2
                                var v158 float32 = 0.1
                                if true {
                                    var v159 int32 = -998
                                    var v160 uint = (v4 / (((^(v4))) | 1))
                                    _ = v159
                                    _ = v160
                                }
                                _ = v158
                            }
                            _ = v82
                        case "hello":
                            switch (fi70 % ((fi70) | 1)) {
                            case -929:
                                v3 += v3
                            }
                        default:
                            v68 = v68
                            if true {
                                if (!((!true) || ((fi75 < fi70) && true))) {
                                    if (949 == 893) {
                                        var v161 int32 = v1
                                        var v162 uint = v4
                                        if true {
                                            v161 = (-(v161))
                                            fmt.Println(v1, fi75, fi70)
                                            if true {
                                                if true {
                                                    v163 := v6[fi75]
                                                    if true {
                                                        v3 += uint8(v162)
                                                        v164 := GenS13{F0: v163}
                                                        _ = v164
                                                    }
                                                    _ = v163
                                                } else {
                                                    v165 := GenS17{F0: 2}
                                                    fmt.Println(fi75, fi70)
                                                    var v166 float32 = 0.1
                                                    _ = v165
                                                    _ = v166
                                                }
                                            } else {
                                                v6[(fi70 >> ((v162) & 63))] = 42
                                                delete(v6, int(v3))
                                                if (-7 == -32768) {
                                                    v6[fi70] = 4294967295
                                                    var v167 uint8 = (v3 << ((255) & 63))
                                                    _ = v167
                                                } else {
                                                    if false {
                                                        fmt.Println(v5, v77, v3)
                                                        if (198 > 2) {
                                                            v168 := GenS15{F0: v1, F1: uint16(v3)}
                                                            _ = v168
                                                        } else {
                                                            v1 = v1
                                                            switch 128 {
                                                            case -120:
                                                                v169 := GenS9{F0: 42, F1: "a", F2: false}
                                                                v162 -= 255
                                                                var v170 uint8 = 158
                                                                _ = v169
                                                                _ = v170
                                                            case 32768:
                                                                switch (fi70 & int(v2)) {
                                                                case 0:
                                                                    v2 -= (v2 ^ v2)
                                                                case 32768:
                                                                    fmt.Println(v162, v4)
                                                                default:
                                                                    fmt.Println(v1, v4, v2)
                                                                }
                                                            }
                                                            if (v69 <= (^(v69))) {
                                                                var v171 uint = v162
                                                                var v172 float32 = float32((3.14159 - (3.14159 + 1e10)))
                                                                _ = v171
                                                                _ = v172
                                                            } else {
                                                                if false {
                                                                    v4 = 4294967295
                                                                }
                                                                v67 = append(v67, fi70)
                                                            }
                                                        }
                                                    }
                                                    var v173 int64 = (-(v77))
                                                    var v174 uint32 = 711
                                                    if ((-909 != (fi75 ^ (fi75 | (fi70 + (fi75 & (fi70 | (-(fi70)))))))) && true) {
                                                        var v175 int32 = v1
                                                        _ = v175
                                                    } else {
                                                        v173 = v173
                                                    }
                                                    if false {
                                                        var v176 uint64 = v69
                                                        _ = v176
                                                    } else {
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v177 bool = (32768 >= (fi75 | fi70))
                                                            _ = v177
                                                        default:
                                                            _ = len(v6)
                                                        }
                                                        v178 := GenS20{F0: 0, F1: (^(v162))}
                                                        _ = v178
                                                    }
                                                    delete(v6, fi75)
                                                    v5 += v1
                                                    _ = v173
                                                    _ = v174
                                                }
                                            }
                                            v3 = v3
                                            v78 = (("foo bar" + "") + "\n\t")
                                        } else {
                                            if (!true) {
                                                v179 := GenS8{F0: 1000, F1: (fi70 & (fi70 + 9223372036854775807))}
                                                v6[(fi75 << ((42) & 63))] = 4294967295
                                                _ = len(v67)
                                                _ = v179
                                            } else {
                                                v78 = "hello"
                                                v1 = v5
                                            }
                                            fmt.Println(v68, fi75)
                                        }
                                        _ = v161
                                        _ = v162
                                    } else {
                                        switch "foo bar" {
                                        case "a":
                                            fmt.Println(v3)
                                        case "":
                                            v180 := v6[fi75]
                                            _ = v180
                                        default:
                                            if len(v67) > 0 {
                                                v67[0] = (fi70 / ((fi70) | 1))
                                            }
                                        }
                                        v1 = -1
                                    }
                                    v3 = (v3 ^ (^(v3)))
                                    v5 -= (-(v5))
                                    v5 = (v5 ^ -1)
                                }
                                var v181 int64 = int64(v69)
                                delete(v6, int(v181))
                                _ = v181
                            } else {
                                fmt.Println(v5)
                                var v182 int64 = (v77 >> (((v4 << ((v4) & 63))) & 63))
                                var v183 int16 = 32767
                                _ = v182
                                _ = v183
                            }
                            var v184 int16 = 32767
                            fmt.Println(v68, v68)
                            _ = v184
                        }
                    }
                }
                _ = fi75
                _ = v77
                _ = v78
                _ = v79
            }
            _ = fi70
        }
    } else {
        if true {
            v185 := GenS1{F0: true, F1: (v3 << (((v4 << ((679) & 63))) & 63))}
            _ = v185
        }
        var v186 uint8 = (v3 >> (((v4 % (((v4 << ((v4) & 63))) | 1))) & 63))
        if false {
            v68 = v68
        }
        _ = v186
    }
    if false {
        for fi187 := 0; fi187 < 6; fi187++ {
            v69 += (v69 ^ (v69 & v69))
            v188 := GenS19{F0: 895}
            _ = len(v67)
            _ = fi187
            _ = v188
        }
        var v189 float64 = 1e10
        for fi190 := 0; fi190 < 12; fi190++ {
            fmt.Println(v2)
            switch "\n\t" {
            case "":
                var v191 float64 = (1.5 - v189)
                _ = v191
            case "hello":
                var v192 string = (v68 + ("a" + ""))
                _ = v192
            default:
                v6[fi190] = 423
            }
            _ = fi190
        }
        _ = v189
    }
    if len(v67) > 0 {
        v67[0] = 2147483647
    }
    var v193 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v193:", v193)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v193
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v69
}
