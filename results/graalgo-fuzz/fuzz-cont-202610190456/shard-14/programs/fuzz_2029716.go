package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int32; F1 uint; F2 int }
type GenS6 struct { F0 bool; F1 int64; F2 bool }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 int64; F1 bool; F2 int32 }
type GenS9 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS10 struct { F0 string; F1 int32; F2 int }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 bool; F1 string }
type GenS15 struct { F0 int16; F1 int32 }
type GenS16 struct { F0 int32; F1 uint; F2 uint }
type GenS17 struct { F0 int8; F1 int8 }
type GenS18 struct { F0 int16; F1 uint8; F2 int32 }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 bool; F1 uint64 }
type GenS21 struct { F0 bool; F1 int8 }
type GenS22 struct { F0 uint8 }
type GenS23 struct { F0 bool }
func genFn0(p0 uint16, p1 uint32, p2 bool) uint {
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(p2)
        var v6 int32 = 42
        p2 = p2
        _ = fi5
        _ = v6
    }
    switch ("a" + ("foo bar" + "foo bar")) {
    case "hello":
        if p2 {
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 float64 = 1e10
                switch 128 {
                case 128:
                    v9 := []int{(fi7 & fi7), 1, (fi7 & (fi7 >> ((48) & 63))), (^(fi7))}
                    _ = v9
                case 0:
                    if p2 {
                        if p2 {
                            fmt.Println(fi7, fi7)
                        } else {
                            for fi10 := 0; fi10 < 19; fi10++ {
                                if false {
                                    var v11 uint64 = 1
                                    _ = v11
                                }
                                _ = fi10
                            }
                            fmt.Println(p0)
                        }
                        v12 := []int{-178, (fi7 / ((fi7) | 1)), 0, int(p0)}
                        p1 = (p1 % ((p1) | 1))
                        _ = v12
                    } else {
                        if true {
                            v13 := GenS1{F0: 65535, F1: 543}
                            _ = v13
                        } else {
                            var v14 uint = 483
                            fmt.Println(v8, fi7)
                            _ = v14
                        }
                        var v15 bool = (!(((^(p1)) == (p1 ^ p1)) && p2))
                        fmt.Println(v15, p0, p2)
                        _ = v15
                    }
                default:
                    if (-32768 <= 1) {
                        v16 := []int{fi7}
                        p1 -= (p1 - p1)
                        _ = v16
                    } else {
                        if p2 {
                            if p2 {
                                v17 := []int{}
                                for fi18 := 0; fi18 < 17; fi18++ {
                                    switch fi7 {
                                    case 127:
                                        p0 += 255
                                    default:
                                        v19 := GenS2{F0: 1}
                                        _ = v19
                                    }
                                    var v20 uint8 = 42
                                    v17 = append(v17, (fi18 * (fi7 ^ -2147483648)))
                                    _ = fi18
                                    _ = v20
                                }
                                if p2 {
                                    fmt.Println(fi7)
                                    p1 += 0
                                    v21 := GenS2{F0: 280}
                                    _ = v21
                                }
                                _ = v17
                            } else {
                                if true {
                                    fmt.Println(p1, v8, p0)
                                    for fi22 := 0; fi22 < 10; fi22++ {
                                        var v23 uint16 = 42
                                        v24 := GenS1{F0: 966, F1: 42}
                                        p2 = true
                                        _ = fi22
                                        _ = v23
                                        _ = v24
                                    }
                                    var v25 float64 = v8
                                    _ = v25
                                }
                                fmt.Println(fi7, p0, p1)
                                for fi26 := 0; fi26 < 14; fi26++ {
                                    p1 = p1
                                    fmt.Println(p1)
                                    switch ("\n\t" + "") {
                                    case "a":
                                        switch fi26 {
                                        case 127:
                                            v27 := []int{78}
                                            _ = v27
                                        case 2147483647:
                                            if p2 {
                                                var v28 uint64 = 4294967295
                                                _ = v28
                                            } else {
                                                p2 = p2
                                                p0 -= 37
                                                p0 += 65535
                                            }
                                        default:
                                            fmt.Println(p2, fi26, p2)
                                            var v29 int8 = -1
                                            if ((((v8 / -2.25) + v8) - 3.14159) != (v8 - v8)) {
                                                v30 := GenS3{F0: 1, F1: p1}
                                                v31 := GenS4{F0: 117, F1: 358}
                                                var v32 float64 = (3.14159 * ((0.0 + v8) - (float64(-2.25) + v8)))
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            } else {
                                                if ((p1 + p1) >= (p1 >> ((255) & 63))) {
                                                    fmt.Println(p2)
                                                } else {
                                                    v33 := GenS5{F0: int32(p0), F1: 42, F2: fi26}
                                                    switch ((("\n\t" + "foo bar") + ("\n\t" + "")) + (("\n\t" + "a") + "\n\t")) {
                                                    case "a":
                                                        if (false || (!("\n\t" == ""))) {
                                                            p1 -= (p1 << ((0) & 63))
                                                            v34 := GenS6{F0: ((!(1 > 199)) || false), F1: 42, F2: (295 <= fi26)}
                                                            _ = v34
                                                        }
                                                    default:
                                                        v35 := []int{}
                                                        _ = v35
                                                    }
                                                    v36 := []int{0, (fi26 % ((fi7) | 1)), (^(fi26)), fi26}
                                                    switch "" {
                                                    case "a":
                                                        v29 = 42
                                                    }
                                                    p1 -= p1
                                                    v37 := GenS4{F0: 244, F1: 42}
                                                    if ((v29 << ((18446744073709551615) & 63)) <= -1) {
                                                        switch "hello" {
                                                        case "hello":
                                                            p1 += (^(p1))
                                                        case "\n\t":
                                                            v38 := GenS2{F0: 1}
                                                            _ = v38
                                                        }
                                                        v39 := GenS7{F0: 4294967295, F1: -88}
                                                        if (2 <= 189) {
                                                            if p2 {
                                                                v29 += (v29 | v29)
                                                                if true {
                                                                    if p2 {
                                                                        switch fi26 {
                                                                        case 32767:
                                                                            v29 += v29
                                                                        case 128:
                                                                            if len(v36) > 0 {
                                                                                v36[0] = fi7
                                                                            }
                                                                        case 42:
                                                                            p0 += p0
                                                                        default:
                                                                            v36 = append(v36, fi7)
                                                                        }
                                                                        if len(v36) > 0 {
                                                                            v36[0] = (fi26 - (fi26 + (fi7 + (fi26 - (fi7 ^ fi7)))))
                                                                        }
                                                                        v40 := GenS6{F0: (!(9223372036854775807 < -736)), F1: -1, F2: p2}
                                                                        fmt.Println(v29, fi26, v8)
                                                                        _ = v40
                                                                    }
                                                                    v36 = append(v36, fi26)
                                                                    var v41 float32 = (1e10 - (float32(3.14159) / (float32((1.5 + -2.25)) / (1e10 / 3.14159))))
                                                                    _ = v41
                                                                } else {
                                                                    v29 = (v29 ^ (v29 - -86))
                                                                    fmt.Println(v29, v29, p0)
                                                                }
                                                            }
                                                            var v42 int16 = 893
                                                            v36 = append(v36, fi26)
                                                            var v43 int32 = 42
                                                            v42 -= int16(p1)
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        p0 -= p0
                                                        _ = v39
                                                    } else {
                                                        if len(v36) > 0 {
                                                            v36[0] = fi7
                                                        }
                                                        fmt.Println(p2)
                                                        fmt.Println(p2, p2)
                                                        p1 -= p1
                                                    }
                                                    if false {
                                                        fmt.Println(p0, v29, p2)
                                                        var v44 uint32 = p1
                                                        p1 -= p1
                                                        _ = v44
                                                    } else {
                                                        v45 := GenS8{F0: -465, F1: true, F2: 743}
                                                        if ((p0 << ((622) & 63)) <= p0) {
                                                            if false {
                                                                v46 := GenS1{F0: 943, F1: 65535}
                                                                _ = v46
                                                            } else {
                                                                v47 := GenS4{F0: 127, F1: 2147483647}
                                                                if ((p0 / ((p0) | 1)) == 1) {
                                                                    var v48 int32 = -2147483648
                                                                    _ = v48
                                                                }
                                                                _ = v47
                                                            }
                                                            if p2 {
                                                                var v49 float64 = -2.25
                                                                _ = v49
                                                            }
                                                            switch fi26 {
                                                            case 151:
                                                                var v50 float64 = v8
                                                                _ = v50
                                                            default:
                                                                v51 := GenS9{F0: 255, F1: p0, F2: 181}
                                                                _ = v51
                                                            }
                                                            var v52 uint = 4294967295
                                                            _ = v52
                                                        } else {
                                                            switch fi26 {
                                                            case 2147483647:
                                                                fmt.Println(fi26, v8, v29)
                                                            case -625:
                                                                v53 := GenS1{F0: 1, F1: 1}
                                                                v29 -= (v29 << ((2) & 63))
                                                                var v54 uint64 = uint64(p0)
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                        }
                                                        _ = v45
                                                    }
                                                    if p2 {
                                                        switch (("\n\t" + ("a" + ("a" + ("\n\t" + "a")))) + "") {
                                                        case "a":
                                                            if (25 > 1) {
                                                                v55 := GenS10{F0: "", F1: 42, F2: (fi26 / ((666) | 1))}
                                                                if len(v36) >= 2 {
                                                                    v36 = v36[1:len(v36)]
                                                                }
                                                                var v56 bool = (0.1 > 1.5)
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                        case "\n\t":
                                                            p0 = 1
                                                        case "":
                                                            v57 := GenS11{F0: 2}
                                                            _ = v57
                                                        }
                                                        v58 := GenS9{F0: uint8(p1), F1: p0, F2: 42}
                                                        _ = v58
                                                    }
                                                    _ = v33
                                                    _ = v36
                                                    _ = v37
                                                }
                                                var v59 int64 = -9223372036854775808
                                                p0 -= 255
                                                v59 = v59
                                                v59 += v59
                                                v60 := GenS10{F0: "a", F1: -1, F2: (^(fi26))}
                                                _ = v59
                                                _ = v60
                                            }
                                            _ = v29
                                        }
                                    }
                                    var v61 string = (("hello" + "a") + "foo bar")
                                    var v62 uint = 18446744073709551615
                                    _ = fi26
                                    _ = v61
                                    _ = v62
                                }
                            }
                        }
                        var v63 float32 = float32(v8)
                        var v64 bool = p2
                        _ = v63
                        _ = v64
                    }
                }
                if (!p2) {
                    v65 := []int{fi7}
                    _ = v65
                } else {
                    var v66 int8 = -45
                    _ = v66
                }
                _ = fi7
                _ = v8
            }
        } else {
            var v67 uint64 = uint64(p1)
            _ = v67
        }
    case "foo bar":
        if p2 {
            if p2 {
                var v68 int = -9223372036854775808
                _ = v68
            }
            if false {
                for fi69 := 0; fi69 < 7; fi69++ {
                    v70 := []int{(fi69 - 200), -2147483648, 9223372036854775807}
                    var v71 int64 = 411
                    v71 -= (^(v71))
                    _ = fi69
                    _ = v70
                    _ = v71
                }
                if false {
                    v72 := GenS12{F0: 1}
                    fmt.Println(p1)
                    var v73 float32 = 0.0
                    _ = v72
                    _ = v73
                }
                v74 := GenS7{F0: 2, F1: 100}
                _ = v74
            } else {
                v75 := []int{128, 86, -1, 244}
                _ = v75
            }
        } else {
            fmt.Println(p0)
        }
    default:
        for fi76 := 0; fi76 < 7; fi76++ {
            fmt.Println(p1, p2)
            v77 := []int{}
            _ = fi76
            _ = v77
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 float64 = 3.14159
    var v2 string = ("" + (((("a" + ("foo bar" + "\n\t")) + "hello") + "foo bar") + ""))
    v3 := []int{32768, -305, 173}
    var v4 uint16 = 1
    if (!(!(42 != 2))) {
        if ((3.14159 - 1e10) <= 1.5) {
            switch 4294967295 {
            case -2147483648:
                if (-571 != 817) {
                    switch "a" {
                    case "\n\t":
                        v2 = ((v2 + v2) + ("hello" + (v2 + ("a" + "\n\t"))))
                    }
                    for fi78 := 0; fi78 < 10; fi78++ {
                        switch (fi78 >> ((4294967295) & 63)) {
                        case 42:
                            var v79 uint8 = 99
                            _ = v79
                        case -703:
                            var v80 uint64 = 18446744073709551615
                            _ = v80
                        case 614:
                            v81 := GenS8{F0: -997, F1: ((0 > int16(fi78)) && false), F2: 887}
                            _ = v81
                        }
                        _ = fi78
                    }
                }
            case -9223372036854775808:
                if (42 != 42) {
                    v82 := GenS13{F0: 42}
                    v4 += (v4 ^ (^(v4)))
                    v2 = v2
                    _ = v82
                }
            case 9223372036854775807:
                var v83 float64 = -2.25
                _ = v83
            }
            if len(v3) > 0 {
                v3[0] = -787
            }
            fmt.Println(v1)
        } else {
            if (2147483647 == 65535) {
                switch 42 {
                case 32767:
                    v3 = append(v3, 4294967295)
                case 1:
                    v4 -= 32767
                case 128:
                    for fi84 := 0; fi84 < 7; fi84++ {
                        if (v2 == "hello") {
                            v1 = (v1 - (v1 - float64(-2.25)))
                            for fi85 := 0; fi85 < 20; fi85++ {
                                var v86 string = "a"
                                if false {
                                    var v87 uint64 = 249
                                    _ = v87
                                } else {
                                    v88 := GenS14{F0: true, F1: v2}
                                    v4 = v4
                                    v2 = (v2 + (v86 + (v2 + "a")))
                                    _ = v88
                                }
                                var v89 float64 = v1
                                _ = fi85
                                _ = v86
                                _ = v89
                            }
                        } else {
                            if false {
                                v90 := GenS11{F0: 1}
                                _ = v90
                            } else {
                                fmt.Println(v4, v2)
                                v91 := GenS15{F0: 1000, F1: -321}
                                if (v1 >= v1) {
                                    var v92 int = fi84
                                    for fi93 := 0; fi93 < 8; fi93++ {
                                        fmt.Println(v4, v2, v4)
                                        var v94 int16 = 1000
                                        _ = fi93
                                        _ = v94
                                    }
                                    v3 = append(v3, fi84)
                                    _ = v92
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                _ = v91
                            }
                            v1 = v1
                            _ = len(v3)
                        }
                        v4 = (v4 & 1)
                        if (true || (true || (!false))) {
                            v3 = append(v3, -1)
                            var v95 int8 = 1
                            v96 := GenS16{F0: -903, F1: 2, F2: 18446744073709551615}
                            _ = v95
                            _ = v96
                        } else {
                            fmt.Println(v2, fi84, v2)
                            var v97 bool = true
                            v4 = (v4 / ((v4) | 1))
                            _ = v97
                        }
                        _ = fi84
                    }
                }
            } else {
                if true {
                    for fi98 := 0; fi98 < 2; fi98++ {
                        var v99 uint16 = v4
                        v100 := GenS17{F0: 100, F1: -1}
                        _ = fi98
                        _ = v100
                        _ = v99
                    }
                    var v101 uint8 = 2
                    if true {
                        v102 := GenS13{F0: -344}
                        if (true && false) {
                            fmt.Println(v101, v4, v101)
                        } else {
                            if true {
                                v101 -= v101
                            } else {
                                for fi103 := 0; fi103 < 17; fi103++ {
                                    for fi104 := 0; fi104 < 18; fi104++ {
                                        if (false || true) {
                                            fmt.Println(fi104, v1, v1)
                                            fmt.Println(fi103)
                                        } else {
                                            _ = len(v3)
                                            if false {
                                                if true {
                                                    var v105 string = (v2 + ("hello" + v2))
                                                    switch v2 {
                                                    case "":
                                                        v101 += v101
                                                    }
                                                    _ = v105
                                                }
                                                if (!(v1 >= (v1 * ((v1 / (v1 / (v1 + (v1 / ((v1 / 1e10) * 1e10))))) / v1)))) {
                                                    fmt.Println(v1)
                                                    fmt.Println(fi103, fi103, v1)
                                                    if false {
                                                        var v106 uint8 = (v101 | (v101 & v101))
                                                        if false {
                                                            v107 := GenS11{F0: 440}
                                                            v108 := GenS5{F0: 1, F1: genFn0(255, 275, true), F2: fi104}
                                                            _ = v107
                                                            _ = v108
                                                        }
                                                        v4 -= (v4 % ((v4) | 1))
                                                        _ = v106
                                                    } else {
                                                        v1 = v1
                                                        if len(v3) > 0 {
                                                            v3[0] = (fi103 + fi103)
                                                        }
                                                    }
                                                }
                                                if (true || (false && true)) {
                                                    fmt.Println(fi103)
                                                    fmt.Println(fi103, v4, v101)
                                                    if (float64((0.1 + 1.5)) == (-2.25 - 0.1)) {
                                                        var v109 uint32 = 1
                                                        var v110 string = "\n\t"
                                                        var v111 int16 = -1
                                                        _ = v109
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                }
                                                switch (fi103 / (((fi103 - 375)) | 1)) {
                                                case 2147483647:
                                                    fmt.Println(fi103)
                                                case 260:
                                                    var v112 int32 = -2147483648
                                                    _ = v112
                                                default:
                                                    var v113 int16 = 1
                                                    _ = v113
                                                }
                                                var v114 uint = uint(v101)
                                                v115 := GenS11{F0: 18446744073709551615}
                                                v4 -= 896
                                                var v116 string = ""
                                                _ = v114
                                                _ = v115
                                                _ = v116
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            v117 := GenS12{F0: 4294967295}
                                            _ = v117
                                        }
                                        v118 := GenS18{F0: 0, F1: (v101 / (((v101 - (v101 + 255))) | 1)), F2: 0}
                                        _ = fi104
                                        _ = v118
                                    }
                                    fmt.Println(v1, v4)
                                    _ = fi103
                                }
                                for fi119 := 0; fi119 < 1; fi119++ {
                                    for fi120 := 0; fi120 < 18; fi120++ {
                                        v1 = 0.0
                                        if true {
                                            if false {
                                                fmt.Println(fi120)
                                            } else {
                                                var v121 int8 = -98
                                                if len(v3) > 0 {
                                                    v3[0] = (^(fi120))
                                                }
                                                fmt.Println(fi120, v121, v121)
                                                var v122 string = v2
                                                v101 = (v101 - 255)
                                                _ = v121
                                                _ = v122
                                            }
                                            var v123 int32 = -2147483648
                                            v3 = append(v3, fi120)
                                            v1 += v1
                                            var v124 int8 = 10
                                            _ = v123
                                            _ = v124
                                        } else {
                                            var v125 uint32 = 2
                                            var v126 int64 = -9223372036854775808
                                            v1 = float64((((0.0 / (float32(v1) - ((float32((v1 + float64(1.5))) - 0.0) + (0.1 * -2.25)))) - 0.0) * -2.25))
                                            if true {
                                                var v127 uint = 801
                                                v128 := GenS16{F0: 0, F1: (v127 - v127), F2: 255}
                                                _ = len(v3)
                                                switch "hello" {
                                                case "":
                                                    var v129 int32 = 1
                                                    _ = v129
                                                case "hello":
                                                    v130 := GenS19{F0: 1000}
                                                    _ = v130
                                                }
                                                _ = v127
                                                _ = v128
                                            } else {
                                                v3 = append(v3, (fi120 * 9223372036854775807))
                                            }
                                            _ = v125
                                            _ = v126
                                        }
                                        _ = fi120
                                    }
                                    _ = fi119
                                }
                                v131 := GenS9{F0: v101, F1: 0, F2: (^(v101))}
                                _ = v131
                            }
                            if (2147483647 != 42) {
                                var v132 int = 1
                                switch (v132 >> ((4294967295) & 63)) {
                                case 1:
                                    v133 := GenS17{F0: -1, F1: -38}
                                    _ = v133
                                case 0:
                                    var v134 int32 = int32(v4)
                                    _ = v134
                                default:
                                    var v135 int32 = 581
                                    _ = v135
                                }
                                _ = v132
                            }
                            switch 32767 {
                            case -1:
                                switch (v2 + "\n\t") {
                                case "a":
                                    switch v2 {
                                    case "foo bar":
                                        v136 := GenS20{F0: true, F1: 4294967295}
                                        _ = v136
                                    case "hello":
                                        v101 = v101
                                    case "\n\t":
                                        var v137 string = "hello"
                                        _ = v137
                                    default:
                                        var v138 uint8 = (v101 / ((v101) | 1))
                                        _ = v138
                                    }
                                case "":
                                    v1 = 0.1
                                default:
                                    var v139 string = v2
                                    _ = v139
                                }
                            case 32768:
                                switch -1 {
                                case -2147483648:
                                    for fi140 := 0; fi140 < 15; fi140++ {
                                        fmt.Println(v4, v4, v4)
                                        _ = fi140
                                    }
                                case 128:
                                    v4 -= 65535
                                case -864:
                                    fmt.Println(v2, v101)
                                }
                            default:
                                switch int(v4) {
                                case 574:
                                    v1 = (((v1 / ((v1 * (float64(((1e10 - (0.0 + 1e10)) * -2.25)) * 0.1)) * float64((-2.25 + (float32((1e10 / v1)) - -2.25))))) + 3.14159) + 0.0)
                                case 9223372036854775807:
                                    var v141 uint64 = 865
                                    _ = v141
                                default:
                                    v1 -= v1
                                }
                            }
                        }
                        _ = v102
                    } else {
                        v1 += v1
                        if (255 == 42) {
                            switch 4294967295 {
                            case -1:
                                var v142 float32 = 3.14159
                                _ = v142
                            }
                            var v143 int8 = -31
                            _ = len(v3)
                            _ = v143
                        } else {
                            if true {
                                v101 = 0
                                v144 := GenS19{F0: 42}
                                _ = v144
                            } else {
                                var v145 uint32 = 1
                                if false {
                                    v146 := GenS21{F0: ((true && (3.14159 <= v1)) || (9223372036854775807 >= 0)), F1: 1}
                                    fmt.Println(v4)
                                    _ = v146
                                } else {
                                    var v147 int64 = -9223372036854775808
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v147
                                }
                                _ = v145
                            }
                        }
                        v148 := GenS22{F0: (v101 << ((0) & 63))}
                        _ = v148
                    }
                    _ = v101
                }
                v149 := GenS21{F0: true, F1: 127}
                v2 = v2
                _ = v149
            }
            if true {
                if true {
                    v2 = ((("\n\t" + v2) + ("a" + v2)) + ("a" + v2))
                    v2 = "hello"
                }
                _ = len(v3)
                _ = len(v3)
            } else {
                v2 = v2
            }
        }
        v4 = v4
    }
    v150 := GenS23{F0: (!false)}
    v3 = append(v3, 1)
    v4 = v4
    var v151 uint16 = v4
    var v152 float32 = 0.1
    var v153 uint64 = 2
    for fi154 := 0; fi154 < 5; fi154++ {
        var v155 uint16 = (v151 % ((v151) | 1))
        _ = fi154
        _ = v155
    }
    var v156 bool = (255 == 42)
    switch -965 {
    case -1:
        v157 := GenS20{F0: (!false), F1: uint64(v4)}
        _ = v157
    }
    switch 4294967295 {
    case 437:
        for fi158 := 0; fi158 < 7; fi158++ {
            v159 := GenS10{F0: ("hello" + "a"), F1: -2147483648, F2: (fi158 - (fi158 * -748))}
            _ = fi158
            _ = v159
        }
    case -965:
        for fi160 := 0; fi160 < 16; fi160++ {
            var v161 int32 = 659
            v4 -= v151
            _ = fi160
            _ = v161
        }
    default:
        switch 386 {
        case -449:
            v151 = (v4 & 0)
        }
    }
    switch 32768 {
    case 128:
        v162 := GenS10{F0: v2, F1: 42, F2: 4294967295}
        _ = v162
    }
    var v163 int8 = -1
    v3 = append(v3, -1)
    if len(v3) > 0 {
        v3[0] = 32768
    }
    _ = len(v3)
    var v164 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v151:", v151)
    fmt.Println("v152:", v152)
    fmt.Println("v153:", v153)
    fmt.Println("v156:", v156)
    fmt.Println("v163:", v163)
    fmt.Println("v164:", v164)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v150
    _ = v151
    _ = v152
    _ = v153
    _ = v156
    _ = v163
    _ = v164
    _ = v2
    _ = v3
    _ = v4
}
