package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int64; F1 bool; F2 uint8 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 uint; F2 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint64; F1 uint64; F2 bool }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int64; F1 int32 }
type GenS11 struct { F0 int64; F1 int8; F2 uint16 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int32; F1 int64 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int; F1 string }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 1
    var v4 int32 = 182
    if (-1 == (v4 + -425)) {
        fmt.Println(v4, v3, v4)
    } else {
        if (true || true) {
            var v5 float64 = float64((((0.1 + 0.1) * ((1e10 / ((-2.25 + -2.25) - 3.14159)) * 3.14159)) * float32(3.14159)))
            var v6 int16 = -837
            fmt.Println(v2, v4, v5)
            _ = v5
            _ = v6
        } else {
            switch 375 {
            case -101:
                var v7 int8 = -1
                _ = v7
            case 568:
                var v8 int32 = v4
                _ = v8
            }
        }
        v9 := GenS1{F0: 4294967295, F1: (v4 >> ((361) & 63))}
        _ = v9
    }
    v3 = v3
    for fi10 := 0; fi10 < 5; fi10++ {
        v3 -= (v3 * v3)
        var v11 uint8 = 255
        _ = fi10
        _ = v11
    }
    if false {
        v12 := GenS2{F0: v2, F1: (false || (v4 > v4)), F2: 0}
        if (2147483647 > 538) {
            switch "\n\t" {
            case "":
                for fi13 := 0; fi13 < 14; fi13++ {
                    if true {
                        v14 := []int{fi13, -2147483648, 32767, (fi13 / ((-9223372036854775808) | 1))}
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 5; fi15++ {
                            v3 = (v3 & (v3 % ((42) | 1)))
                            var v16 uint = 99
                            _ = fi15
                            _ = v16
                        }
                        var v17 uint32 = 290
                        _ = v17
                    }
                    v3 = (v3 << ((684) & 63))
                    _ = fi13
                }
            case "\n\t":
                for fi18 := 0; fi18 < 10; fi18++ {
                    for fi19 := 0; fi19 < 2; fi19++ {
                        if (true || (3.14159 == 3.14159)) {
                            var v20 bool = false
                            if (4294967295 >= 18446744073709551615) {
                                var v21 float32 = -2.25
                                var v22 uint8 = 101
                                fmt.Println(v3)
                                _ = v21
                                _ = v22
                            } else {
                                if (!(v20 || (65535 > 2))) {
                                    if (!((fi18 - (fi19 * (fi18 ^ fi19))) != 42)) {
                                        var v23 uint32 = 2
                                        _ = v23
                                    }
                                    if ((42 > -34) || v20) {
                                        v24 := GenS3{F0: (fi19 / ((fi18) | 1)), F1: 897}
                                        switch (fi19 % ((fi18) | 1)) {
                                        case 32768:
                                            v25 := GenS3{F0: (fi19 >> ((255) & 63)), F1: 0}
                                            _ = v25
                                        case 42:
                                            v26 := []int{fi18, -750, (fi19 * fi18)}
                                            _ = v26
                                        case 128:
                                            var v27 int8 = 100
                                            _ = v27
                                        }
                                        if false {
                                            fmt.Println(v20, v4, v4)
                                            delete(v1, 2)
                                            var v28 uint32 = 49
                                            _ = v28
                                        } else {
                                            var v29 bool = v20
                                            v30 := []int{fi19, 2147483647}
                                            _ = v29
                                            _ = v30
                                        }
                                        if v20 {
                                            if v20 {
                                                v31 := []int{(^(fi18)), 2147483647}
                                                if ((985 > 679) || true) {
                                                    v4 = (^(v4))
                                                    v32 := GenS1{F0: 881, F1: 42}
                                                    var v33 uint32 = 2147483648
                                                    _ = v32
                                                    _ = v33
                                                }
                                                _ = v31
                                            } else {
                                                if (v20 || v20) {
                                                    var v34 float64 = 3.14159
                                                    var v35 uint = 65535
                                                    v36 := GenS4{F0: (^(v35))}
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                }
                                            }
                                        } else {
                                            var v37 string = "foo bar"
                                            fmt.Println(v20, v3, v37)
                                            var v38 uint32 = 2147483648
                                            v37 = (v37 + ((v37 + (v37 + v37)) + v37))
                                            v39 := []int{fi19, fi18, (fi19 + 9223372036854775807), int(v38)}
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        }
                                        _ = v24
                                    }
                                    fmt.Println(fi18, v2, v20)
                                    var v40 int16 = 116
                                    v41 := []int{}
                                    _ = v40
                                    _ = v41
                                } else {
                                    v42 := GenS4{F0: 0}
                                    v4 -= (v4 % ((v4) | 1))
                                    v3 += (v3 | 997)
                                    _ = v42
                                }
                                switch (("foo bar" + "") + ((("hello" + "foo bar") + ("hello" + "a")) + "foo bar")) {
                                case "hello":
                                    if (!((!v20) || (0 != int16(v3)))) {
                                        var v43 int16 = -1
                                        if ((v3 % ((v3) | 1)) >= v3) {
                                            if (!v20) {
                                                var v44 uint16 = 255
                                                var v45 uint16 = uint16(v3)
                                                if (1 == 65535) {
                                                    var v46 bool = (!(127 >= 127))
                                                    var v47 uint8 = 42
                                                    _ = v46
                                                    _ = v47
                                                } else {
                                                    v48 := GenS4{F0: 18446744073709551615}
                                                    _ = v48
                                                }
                                                var v49 string = "hello"
                                                switch 128 {
                                                case -114:
                                                    _ = len(v1)
                                                case 127:
                                                    var v50 uint16 = 2
                                                    _ = v50
                                                }
                                                if false {
                                                    if (2147483648 <= 4294967295) {
                                                        if (v20 || false) {
                                                            if false {
                                                                fmt.Println(v49, v4, v45)
                                                                if ((v3 * (v3 * (v3 ^ v3))) != (v3 - v3)) {
                                                                    fmt.Println(v49, v44, v49)
                                                                } else {
                                                                    fmt.Println(v4, v2, fi19)
                                                                    switch (fi19 | fi19) {
                                                                    case 127:
                                                                        v51 := GenS5{F0: 42}
                                                                        _ = v51
                                                                    case 809:
                                                                        var v52 uint16 = 2
                                                                        _ = v52
                                                                    case 0:
                                                                        var v53 int = (fi19 >> ((18446744073709551615) & 63))
                                                                        _ = v53
                                                                    }
                                                                    v54 := GenS3{F0: int(v43), F1: 1}
                                                                    v45 -= v45
                                                                    _ = v54
                                                                }
                                                                fmt.Println(v3)
                                                                var v55 int8 = 1
                                                                var v56 float64 = 1e10
                                                                v43 -= (v43 * int16(v55))
                                                                _ = v55
                                                                _ = v56
                                                            } else {
                                                                var v57 uint64 = 995
                                                                _ = v57
                                                            }
                                                            fmt.Println(v43, v20)
                                                        }
                                                    } else {
                                                        v4 -= (v4 ^ 2147483647)
                                                        var v58 int16 = (v43 >> ((42) & 63))
                                                        _ = v58
                                                    }
                                                } else {
                                                    v59 := []int{(-(fi18)), (fi19 & -492)}
                                                    if true {
                                                        v60 := GenS6{F0: 907, F1: uint(v3), F2: 127}
                                                        if v20 {
                                                            v61 := GenS7{F0: 2147483648}
                                                            if v20 {
                                                                v62 := GenS8{F0: (v3 / ((v3) | 1)), F1: v3, F2: (((v4 % (((-(v4))) | 1)) <= (v4 | 252)) && ((v49 != (("" + (((v49 + v49) + v49) + "\n\t")) + "")) || ((v43 << ((588) & 63)) >= (v43 & 1))))}
                                                                v20 = ((-2.25 / 0.1) <= 1e10)
                                                                fmt.Println(v45)
                                                                _ = v62
                                                            }
                                                            v45 -= 32767
                                                            _ = v61
                                                        }
                                                        _ = v60
                                                    }
                                                    v2 -= (v2 / ((v2) | 1))
                                                    _ = v59
                                                }
                                                v4 = -421
                                                _ = v44
                                                _ = v45
                                                _ = v49
                                            } else {
                                                v63 := GenS9{F0: (fi19 % ((-9223372036854775808) | 1))}
                                                fmt.Println(v43, v2, v4)
                                                var v64 bool = (217 > 124)
                                                _ = v63
                                                _ = v64
                                            }
                                            if false {
                                                var v65 int32 = (v4 & v4)
                                                var v66 int16 = -1
                                                _ = v65
                                                _ = v66
                                            } else {
                                                var v67 int16 = v43
                                                _ = v67
                                            }
                                        } else {
                                            v20 = (!v20)
                                            v68 := GenS10{F0: v2, F1: v4}
                                            fmt.Println(v4, v3)
                                            _ = v68
                                        }
                                        _ = v43
                                    } else {
                                        var v69 string = "foo bar"
                                        var v70 uint = uint(v3)
                                        _ = v69
                                        _ = v70
                                    }
                                    v3 = 2
                                case "foo bar":
                                    v71 := v1[9]
                                    if (v4 >= v4) {
                                        switch "\n\t" {
                                        case "foo bar":
                                            if (!v20) {
                                                if v20 {
                                                    var v72 int32 = v4
                                                    switch (("foo bar" + "a") + ((("\n\t" + "") + ("foo bar" + ("" + ("a" + ("a" + "\n\t"))))) + "\n\t")) {
                                                    case "foo bar":
                                                        v71 -= v71
                                                    }
                                                    if false {
                                                        fmt.Println(v3, fi19)
                                                        if (v2 > v2) {
                                                            fmt.Println(fi19, fi18)
                                                        }
                                                        fmt.Println(v3)
                                                    } else {
                                                        var v73 uint = 0
                                                        if v20 {
                                                            v4 = 868
                                                            if (!v20) {
                                                                switch "a" {
                                                                case "foo bar":
                                                                    v2 = -38
                                                                case "":
                                                                    var v74 uint = v73
                                                                    _ = v74
                                                                case "\n\t":
                                                                    v71 = (v71 ^ (v71 | v71))
                                                                }
                                                                v75 := GenS11{F0: 416, F1: 22, F2: 418}
                                                                v76 := GenS12{F0: 0}
                                                                v73 = 4294967295
                                                                _ = v75
                                                                _ = v76
                                                            } else {
                                                                fmt.Println(v2, v2, v4)
                                                                v20 = true
                                                                switch "" {
                                                                case "foo bar":
                                                                    if ((-128 <= 75) && (!((255 >= 255) && v20))) {
                                                                        var v77 int64 = v2
                                                                        _ = v77
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        switch (fi19 >> (((v73 >> (((v73 * v73)) & 63))) & 63)) {
                                                        case 434:
                                                            v78 := GenS13{F0: 2147483647, F1: (v2 + (v2 * (-(v2))))}
                                                            _ = v78
                                                        case 32768:
                                                            var v79 int16 = -32768
                                                            _ = v79
                                                        default:
                                                            var v80 int32 = (v72 * (v72 % (((v72 | int32(v2))) | 1)))
                                                            v1[v71] = v71
                                                            _ = v80
                                                        }
                                                        _ = v73
                                                    }
                                                    _ = v72
                                                }
                                            }
                                        default:
                                            v71 = v71
                                            if (v4 != (v4 | v4)) {
                                                var v81 int = fi18
                                                var v82 int32 = (v4 + v4)
                                                var v83 int = (v81 ^ fi18)
                                                _ = v81
                                                _ = v82
                                                _ = v83
                                            } else {
                                                v84 := []int{(fi19 | fi18), fi18, 42, (fi19 | fi19)}
                                                v1[v71] = v71
                                                _ = v84
                                            }
                                        }
                                        var v85 uint16 = 47
                                        var v86 float32 = 0.0
                                        _ = v85
                                        _ = v86
                                    } else {
                                        var v87 uint = 255
                                        if false {
                                            v88 := GenS14{F0: ("" + "a")}
                                            switch fi19 {
                                            case 926:
                                                var v89 bool = (uint8(v3) > 255)
                                                _ = v89
                                            case -583:
                                                if (v4 < (v4 ^ v4)) {
                                                    if false {
                                                        v4 -= (v4 & (v4 + v4))
                                                        var v90 uint16 = 42
                                                        _ = v90
                                                    } else {
                                                        fmt.Println(fi18, v87, fi18)
                                                    }
                                                    var v91 uint64 = (^(v3))
                                                    _ = v91
                                                } else {
                                                    var v92 float32 = 1e10
                                                    v4 = 0
                                                    _ = v92
                                                }
                                            }
                                            var v93 float32 = 3.14159
                                            _ = v88
                                            _ = v93
                                        } else {
                                            v94 := GenS13{F0: v4, F1: v2}
                                            _ = v94
                                        }
                                        _ = v87
                                    }
                                    var v95 uint64 = v3
                                    _ = v71
                                    _ = v95
                                }
                            }
                            v2 = (v2 * -152)
                            v3 = 42
                            fmt.Println(v2, v2)
                            _ = v20
                        }
                        fmt.Println(fi18, v2, fi18)
                        if (((42 == 248) && (-2.25 >= (0.1 + 3.14159))) || false) {
                            fmt.Println(v3)
                        } else {
                            v96 := GenS13{F0: v4, F1: v2}
                            v97 := GenS5{F0: 182}
                            _ = v96
                            _ = v97
                        }
                        v4 = v4
                        v98 := GenS6{F0: 352, F1: 65535, F2: 127}
                        v99 := []int{}
                        fmt.Println(fi18, v2)
                        v100 := GenS10{F0: (v2 * -842), F1: 2147483647}
                        var v101 string = "a"
                        switch fi19 {
                        case -632:
                            v2 -= (v2 * v2)
                        default:
                            fmt.Println(v3, fi19, fi18)
                        }
                        _ = fi19
                        _ = v100
                        _ = v101
                        _ = v98
                        _ = v99
                    }
                    _ = fi18
                }
            default:
                switch ("foo bar" + ("hello" + "foo bar")) {
                case "a":
                    v1[1] = 4294967295
                case "hello":
                    if (188 > 2) {
                        var v102 uint8 = 2
                        fmt.Println(v102, v4, v3)
                        _ = v102
                    }
                default:
                    for fi103 := 0; fi103 < 10; fi103++ {
                        if (!true) {
                            v4 += v4
                            var v104 int64 = (v2 | (v2 - (-(v2))))
                            _ = v104
                        }
                        _ = fi103
                    }
                }
            }
            v105 := GenS8{F0: (v3 & (v3 << ((622) & 63))), F1: (v3 * (v3 / ((1) | 1))), F2: (2 >= 2)}
            v106 := GenS15{F0: 0}
            _ = v105
            _ = v106
        } else {
            switch "\n\t" {
            case "a":
                _ = len(v1)
            case "\n\t":
                var v107 int32 = (^(v4))
                _ = v107
            case "":
                v4 -= -501
            default:
                fmt.Println(v2)
            }
            switch "hello" {
            case "hello":
                switch ("hello" + "\n\t") {
                case "\n\t":
                    v3 = v3
                case "":
                    var v108 float32 = (1.5 / (1e10 - (1.5 / 1.5)))
                    _ = v108
                default:
                    if true {
                        var v109 string = "\n\t"
                        var v110 uint64 = 667
                        fmt.Println(v2, v109)
                        _ = v109
                        _ = v110
                    } else {
                        v2 = v2
                    }
                }
            }
        }
        if (("foo bar" < "") && (false || (!true))) {
            switch 351 {
            case 127:
                v111 := []int{}
                _ = v111
            }
            v4 -= v4
        }
        _ = v12
    } else {
        switch "a" {
        case "a":
            v3 += v3
        }
        fmt.Println(v4, v2)
        v3 += v3
    }
    fmt.Println(v2)
    fmt.Println(v2, v4, v3)
    fmt.Println(v4)
    v4 = v4
    for fi112 := 0; fi112 < 14; fi112++ {
        if false {
            switch ("\n\t" + "a") {
            case "":
                var v113 uint8 = 42
                _ = v113
            case "foo bar":
                v2 -= -1
            case "hello":
                if ((false && (1000 == 156)) || false) {
                    v114 := []int{(fi112 << ((255) & 63)), (fi112 + (fi112 / ((32768) | 1)))}
                    _ = v114
                }
            }
        } else {
            v3 += v3
            for fi115 := 0; fi115 < 11; fi115++ {
                v4 = 20
                var v116 uint16 = 186
                _ = fi115
                _ = v116
            }
        }
        v117 := []int{-2147483648, fi112, 128}
        _ = fi112
        _ = v117
    }
    if false {
        var v118 int64 = -645
        if false {
            var v119 uint8 = 104
            var v120 int64 = (v2 / ((v118) | 1))
            v119 += (^(v119))
            _ = v119
            _ = v120
        } else {
            v121 := []int{32768}
            for fi122 := 0; fi122 < 16; fi122++ {
                var v123 int32 = v4
                switch "a" {
                case "a":
                    var v124 bool = false
                    _ = v124
                }
                _ = fi122
                _ = v123
            }
            _ = v121
        }
        if ((^(v3)) < 42) {
            v125 := GenS16{F0: (v118 << ((447) & 63))}
            _ = v125
        } else {
            if ((1 < 42) && true) {
                v4 = (v4 - 0)
                _ = len(v1)
                if (!true) {
                    if true {
                        var v126 float64 = (3.14159 + 1.5)
                        for fi127 := 0; fi127 < 2; fi127++ {
                            for fi128 := 0; fi128 < 4; fi128++ {
                                v118 -= (v2 % ((-9223372036854775808) | 1))
                                _ = fi128
                            }
                            if false {
                                if ("a" < ("hello" + "")) {
                                    for fi129 := 0; fi129 < 5; fi129++ {
                                        var v130 int = (fi127 / (((fi129 - fi127)) | 1))
                                        if (!(!true)) {
                                            switch v130 {
                                            case 127:
                                                v131 := v1[2]
                                                _ = v131
                                            case 153:
                                                var v132 uint32 = 723
                                                _ = v132
                                            default:
                                                if false {
                                                    if false {
                                                        var v133 int32 = v4
                                                        if (58 < 100) {
                                                            v1[58] = 938
                                                            var v134 int = -465
                                                            switch ("foo bar" + ("hello" + ("foo bar" + (("foo bar" + "a") + ((("foo bar" + "hello") + "\n\t") + ("foo bar" + ("a" + "a"))))))) {
                                                            case "a":
                                                                v2 += int64(v3)
                                                            case "foo bar":
                                                                var v135 int = fi129
                                                                _ = v135
                                                            default:
                                                                v136 := GenS17{F0: (v130 << ((uint(v3)) & 63))}
                                                                _ = v136
                                                            }
                                                            _ = v134
                                                        } else {
                                                            if false {
                                                                v137 := GenS8{F0: 1, F1: v3, F2: false}
                                                                var v138 float64 = (((v126 / (1.5 + v126)) + 1e10) - ((1.5 + 0.1) - v126))
                                                                switch (fi127 ^ 258) {
                                                                case 2147483647:
                                                                    v133 -= v133
                                                                case 652:
                                                                    fmt.Println(v118)
                                                                default:
                                                                    v1[261] = 963
                                                                }
                                                                _ = v137
                                                                _ = v138
                                                            } else {
                                                                var v139 int64 = -9223372036854775808
                                                                switch fi129 {
                                                                case 127:
                                                                    v140 := GenS6{F0: 42, F1: 665, F2: -128}
                                                                    _ = v140
                                                                case 691:
                                                                    v141 := GenS16{F0: v118}
                                                                    _ = v141
                                                                case 32767:
                                                                    var v142 float64 = 1e10
                                                                    _ = v142
                                                                }
                                                                _ = v139
                                                            }
                                                        }
                                                        v143 := GenS16{F0: (v2 - v118)}
                                                        v126 = (v126 * (v126 - float64(1.5)))
                                                        _ = v133
                                                        _ = v143
                                                    }
                                                    v118 = v118
                                                    switch (fi127 % (((fi129 + (v130 * (fi129 - fi129)))) | 1)) {
                                                    case 42:
                                                        v3 += 1
                                                    default:
                                                        switch "" {
                                                        case "foo bar":
                                                            v144 := GenS17{F0: 0}
                                                            _ = v144
                                                        case "":
                                                            v4 = -2147483648
                                                            var v145 uint = 714
                                                            var v146 uint16 = 1
                                                            v3 += (v3 / ((v3) | 1))
                                                            fmt.Println(v130)
                                                            _ = v145
                                                            _ = v146
                                                        default:
                                                            var v147 float64 = v126
                                                            _ = v147
                                                        }
                                                    }
                                                    var v148 uint16 = 244
                                                    var v149 uint32 = 396
                                                    switch ("a" + (("hello" + "") + "\n\t")) {
                                                    case "":
                                                        var v150 float64 = v126
                                                        _ = v150
                                                    case "foo bar":
                                                        v149 -= (v149 - 724)
                                                    default:
                                                        v151 := []int{42, (fi129 * 32767), (fi127 >> ((650) & 63))}
                                                        _ = v151
                                                    }
                                                    _ = v148
                                                    _ = v149
                                                }
                                            }
                                        } else {
                                            v2 = (v2 | v2)
                                            v152 := GenS18{F0: -77}
                                            _ = v152
                                        }
                                        v118 = (v118 | v118)
                                        if (0 <= (v130 << ((113) & 63))) {
                                            v130 -= -9223372036854775808
                                            var v153 uint32 = 1
                                            _ = v153
                                        } else {
                                            var v154 float32 = 1e10
                                            var v155 uint16 = 1
                                            _ = v154
                                            _ = v155
                                        }
                                        v156 := []int{}
                                        var v157 string = ""
                                        v118 = 42
                                        _ = fi129
                                        _ = v130
                                        _ = v156
                                        _ = v157
                                    }
                                    v158 := []int{fi127, (fi127 + fi127)}
                                    fmt.Println(v118, fi127)
                                    _ = v158
                                } else {
                                    var v159 int16 = 1000
                                    for fi160 := 0; fi160 < 1; fi160++ {
                                        switch "" {
                                        case "a":
                                            var v161 int = fi160
                                            v162 := GenS7{F0: 4294967295}
                                            _ = v161
                                            _ = v162
                                        case "\n\t":
                                            if true {
                                                v163 := GenS19{F0: fi127, F1: "a"}
                                                var v164 uint64 = (v3 ^ (v3 + (^(v3))))
                                                _ = v163
                                                _ = v164
                                            }
                                        }
                                        _ = fi160
                                    }
                                    _ = v159
                                }
                            }
                            v3 = 42
                            _ = fi127
                        }
                        fmt.Println(v118)
                        _ = v126
                    }
                    fmt.Println(v118, v118)
                    var v165 uint16 = 255
                    _ = v165
                } else {
                    fmt.Println(v2)
                }
            } else {
                switch 360 {
                case 4294967295:
                    fmt.Println(v4, v4)
                case -9223372036854775808:
                    var v166 int64 = v2
                    _ = v166
                }
            }
        }
        _ = v118
    }
    var v167 int64 = (v2 | 9223372036854775807)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v167:", v167)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v167
    _ = v2
    _ = v3
    _ = v4
}
