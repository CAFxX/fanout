package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int64; F1 bool; F2 uint }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 uint8; F1 int; F2 int16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 uint }
type GenS9 struct { F0 int32; F1 int8; F2 int }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS12 struct { F0 int16; F1 int }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int16; F1 uint32 }
type GenS15 struct { F0 uint32; F1 int16 }
type GenS16 struct { F0 uint16; F1 bool }
type GenS17 struct { F0 int64; F1 int8; F2 int16 }
type GenS18 struct { F0 int64; F1 int8; F2 int64 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint16; F1 int8; F2 int }
type GenS21 struct { F0 uint32; F1 uint32; F2 uint8 }
func genFn0(p0 string, p1 int64) uint64 {
    v5 := GenS1{F0: 42, F1: -1}
    var v6 bool = true
    v7 := GenS1{F0: 791, F1: 1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1() bool {
    if false {
        for fi8 := 0; fi8 < 15; fi8++ {
            if (1 <= 32767) {
                if (true && ((1 != 2) || ("" > ""))) {
                    if (724 > 42) {
                        for fi9 := 0; fi9 < 14; fi9++ {
                            var v10 int16 = -32768
                            v11 := []int{}
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                    } else {
                        v12 := GenS2{F0: 402, F1: false, F2: 18446744073709551615}
                        var v13 int16 = -527
                        for fi14 := 0; fi14 < 17; fi14++ {
                            var v15 int = fi8
                            v16 := GenS1{F0: 0, F1: -32768}
                            var v17 uint16 = 915
                            _ = fi14
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        _ = v12
                        _ = v13
                    }
                    if true {
                        v18 := []int{}
                        v19 := GenS1{F0: 4294967295, F1: 1000}
                        _ = v18
                        _ = v19
                    } else {
                        var v20 int = fi8
                        v21 := []int{}
                        var v22 string = (("\n\t" + "a") + (("hello" + "foo bar") + "\n\t"))
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                }
            } else {
                if (128 < 42) {
                    var v23 float64 = 1e10
                    _ = v23
                } else {
                    v24 := GenS3{F0: 0, F1: 104}
                    fmt.Println(fi8)
                    if false {
                        var v25 int32 = -1
                        var v26 uint8 = 135
                        _ = v25
                        _ = v26
                    } else {
                        fmt.Println(fi8)
                        var v27 uint8 = 1
                        switch 9223372036854775807 {
                        case 32768:
                            var v28 bool = (-9223372036854775808 > 938)
                            _ = v28
                        case 9223372036854775807:
                            v29 := GenS1{F0: 18446744073709551615, F1: 1}
                            _ = v29
                        default:
                            var v30 uint8 = (v27 - v27)
                            _ = v30
                        }
                        _ = v27
                    }
                    _ = v24
                }
            }
            switch ("" + ("foo bar" + "a")) {
            case "hello":
                var v31 int8 = -2
                _ = v31
            }
            _ = fi8
        }
        var v32 uint32 = 382
        _ = v32
    } else {
        v33 := []int{-9223372036854775808, 4294967295, -889, 127}
        v34 := GenS2{F0: 0, F1: true, F2: 42}
        for fi35 := 0; fi35 < 20; fi35++ {
            var v36 bool = (false || (100 >= -20))
            if v36 {
                v37 := GenS4{F0: (fi35 ^ fi35)}
                switch 32767 {
                case 177:
                    var v38 float32 = (-2.25 / 1e10)
                    _ = v38
                }
                _ = v37
            } else {
                if (0 < 42) {
                    var v39 int64 = 0
                    var v40 float64 = 0.0
                    var v41 int32 = -2147483648
                    _ = v39
                    _ = v40
                    _ = v41
                }
                if ((!v36) || v36) {
                    var v42 string = ("a" + (("a" + ("" + "")) + (("" + ("\n\t" + "foo bar")) + "\n\t")))
                    _ = v42
                } else {
                    switch -357 {
                    case 9223372036854775807:
                        switch (("a" + "") + "\n\t") {
                        case "":
                            var v43 int8 = -11
                            _ = v43
                        case "hello":
                            for fi44 := 0; fi44 < 19; fi44++ {
                                var v45 int = fi44
                                var v46 int8 = 127
                                v47 := GenS5{F0: -32768, F1: 1}
                                fmt.Println(v45, v46, fi35)
                                fmt.Println(v46)
                                _ = fi44
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                        case "a":
                            v36 = v36
                        default:
                            v36 = false
                        }
                    case -405:
                        v36 = (!(718 > fi35))
                    }
                    switch "\n\t" {
                    case "":
                        if (((1e10 - -2.25) <= 0.0) && false) {
                            var v48 int = (-(fi35))
                            v36 = v36
                            var v49 int = fi35
                            _ = v48
                            _ = v49
                        }
                    }
                }
                switch -9223372036854775808 {
                case 128:
                    switch "" {
                    case "foo bar":
                        for fi50 := 0; fi50 < 6; fi50++ {
                            switch (fi50 - (-(fi50))) {
                            case 127:
                                v51 := GenS4{F0: (fi50 << ((4294967295) & 63))}
                                _ = v51
                            default:
                                v36 = (1 > -9223372036854775808)
                                var v52 uint8 = 2
                                _ = v52
                            }
                            if v36 {
                                if (-128 == 100) {
                                    v33 = append(v33, fi50)
                                    if len(v33) > 0 {
                                        v33[0] = (fi50 ^ fi50)
                                    }
                                }
                                var v53 int64 = -9223372036854775808
                                v53 -= v53
                                _ = v53
                            } else {
                                v54 := GenS3{F0: 153, F1: 128}
                                if false {
                                    v36 = (!(0 != 128))
                                    v36 = ("foo bar" != "\n\t")
                                    if len(v33) > 0 {
                                        v33[0] = 9223372036854775807
                                    }
                                    _ = len(v33)
                                } else {
                                    if len(v33) > 0 {
                                        v33[0] = (fi35 - (fi35 * (fi35 ^ fi50)))
                                    }
                                }
                                v36 = true
                                if len(v33) > 0 {
                                    v33[0] = 127
                                }
                                _ = v54
                            }
                            _ = fi50
                        }
                    case "":
                        for fi55 := 0; fi55 < 20; fi55++ {
                            var v56 uint = 18446744073709551615
                            var v57 uint32 = 0
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                    case "\n\t":
                        if len(v33) > 0 {
                            v33[0] = (^(fi35))
                        }
                    default:
                        if true {
                            v58 := GenS4{F0: (fi35 / ((631) | 1))}
                            switch "\n\t" {
                            case "hello":
                                switch "foo bar" {
                                case "foo bar":
                                    var v59 int16 = 0
                                    _ = v59
                                case "":
                                    var v60 uint8 = 1
                                    _ = v60
                                default:
                                    v36 = v36
                                }
                            default:
                                v61 := GenS6{F0: 2, F1: fi35, F2: 301}
                                _ = v61
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v36 = (-1 == -128)
                            }
                            _ = v58
                        }
                    }
                case -640:
                    v36 = v36
                case -168:
                    if v36 {
                        _ = len(v33)
                    }
                }
            }
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            _ = fi35
            _ = v36
        }
        _ = v33
        _ = v34
    }
    var v62 string = "hello"
    _ = v62
    return (!true)
}

func main() {
    v1 := []uint32{}
    v2 := []int{}
    v3 := []uint32{}
    var v4 bool = true
    switch "\n\t" {
    case "\n\t":
        v1 = append(v1, 414)
    case "a":
        fmt.Println(v4)
    case "foo bar":
        for fi63 := 0; fi63 < 19; fi63++ {
            var v64 uint = 65535
            if false {
                for fi65 := 0; fi65 < 20; fi65++ {
                    v66 := GenS6{F0: 42, F1: fi65, F2: 1}
                    var v67 uint16 = 2
                    _ = fi65
                    _ = v66
                    _ = v67
                }
                v68 := GenS3{F0: 974, F1: 255}
                _ = v68
            } else {
                fmt.Println(fi63)
                var v69 int8 = 127
                for fi70 := 0; fi70 < 19; fi70++ {
                    v2 = append(v2, fi63)
                    var v71 uint64 = 144
                    v4 = (!false)
                    _ = fi70
                    _ = v71
                }
                _ = v69
            }
            _ = fi63
            _ = v64
        }
    default:
        v4 = (!(3.14159 < ((0.1 + (0.0 * 0.0)) * 1e10)))
    }
    if (29 < 992) {
        if (!(4294967295 > 4294967295)) {
            if (!true) {
                switch 0 {
                case 188:
                    v4 = (924 > 2)
                case 2147483647:
                    for fi72 := 0; fi72 < 20; fi72++ {
                        var v73 uint64 = genFn0("\n\t", 1)
                        _ = fi72
                        _ = v73
                    }
                }
            } else {
                v4 = (v4 || (251 <= 128))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v3 = append(v3, 2147483647)
            v4 = (127 <= 255)
        }
        if true {
            v4 = v4
            if (-116 > 0) {
                v4 = v4
                switch ("foo bar" + (("" + ("" + "foo bar")) + ("foo bar" + (("\n\t" + "hello") + "hello")))) {
                case "":
                    for fi74 := 0; fi74 < 7; fi74++ {
                        var v75 int64 = 1
                        v4 = (-1 > (v75 >> ((611) & 63)))
                        _ = fi74
                        _ = v75
                    }
                default:
                    for fi76 := 0; fi76 < 13; fi76++ {
                        for fi77 := 0; fi77 < 20; fi77++ {
                            v4 = ((0.1 * 3.14159) <= 3.14159)
                            if v4 {
                                var v78 float32 = 3.14159
                                v78 -= v78
                                _ = v78
                            } else {
                                v4 = true
                                if len(v3) > 0 {
                                    v3[0] = 777
                                }
                            }
                            var v79 bool = false
                            v79 = v4
                            if v79 {
                                var v80 uint64 = 568
                                var v81 int16 = -32768
                                if (65535 <= 65535) {
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        v4 = (255 > 836)
                                        v82 := GenS2{F0: 285, F1: false, F2: 18446744073709551615}
                                        if false {
                                            v83 := GenS7{F0: 141}
                                            v84 := GenS8{F0: "a", F1: 0}
                                            var v85 uint64 = genFn0(("a" + (("foo bar" + "") + ("a" + "foo bar"))), -9223372036854775808)
                                            v4 = (!v4)
                                            _ = v83
                                            _ = v84
                                            _ = v85
                                        } else {
                                            switch "a" {
                                            case "":
                                                if false {
                                                    if len(v3) > 0 {
                                                        v3[0] = 4294967295
                                                    }
                                                    var v86 uint = 18446744073709551615
                                                    var v87 bool = (4294967295 > (v80 >> ((42) & 63)))
                                                    if (1 >= v80) {
                                                        v79 = (!(0 > 241))
                                                    } else {
                                                        v88 := GenS9{F0: 2147483647, F1: 127, F2: fi76}
                                                        v86 = 18446744073709551615
                                                        v4 = v87
                                                        _ = len(v1)
                                                        _ = v88
                                                    }
                                                    _ = v86
                                                    _ = v87
                                                }
                                            case "foo bar":
                                                v89 := GenS9{F0: -1, F1: 1, F2: (fi76 & -114)}
                                                _ = v89
                                            case "\n\t":
                                                fmt.Println(v81, fi77, fi77)
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                v90 := GenS10{F0: (v81 + 1000)}
                                                _ = v90
                                            case "":
                                                switch "foo bar" {
                                                case "a":
                                                    v81 -= 42
                                                }
                                            default:
                                                v79 = (0 >= 255)
                                            }
                                        }
                                        _ = v82
                                    }
                                }
                                if false {
                                    _ = len(v2)
                                } else {
                                    var v91 int = fi76
                                    switch ("\n\t" + "hello") {
                                    case "foo bar":
                                        var v92 int = (-(fi76))
                                        v93 := GenS6{F0: 215, F1: v92, F2: v81}
                                        var v94 int8 = -1
                                        _ = v92
                                        _ = v93
                                        _ = v94
                                    case "\n\t":
                                        switch -9223372036854775808 {
                                        case 0:
                                            fmt.Println(v80, fi76)
                                        case 4294967295:
                                            v95 := GenS8{F0: ("foo bar" + "foo bar"), F1: 627}
                                            _ = v95
                                        }
                                        v2 = append(v2, 491)
                                        switch "a" {
                                        case "a":
                                            v81 = 1000
                                        case "foo bar":
                                            var v96 float64 = float64((0.1 * 3.14159))
                                            _ = v96
                                        case "":
                                            var v97 int32 = -1
                                            _ = v97
                                        }
                                        var v98 int16 = 1000
                                        var v99 int8 = 42
                                        _ = v98
                                        _ = v99
                                    default:
                                        v91 = fi76
                                    }
                                    _ = v91
                                }
                                _ = v80
                                _ = v81
                            }
                            v79 = (510 == 658)
                            fmt.Println(fi77)
                            if v79 {
                                fmt.Println(v4, v79, v79)
                                var v100 uint16 = 1
                                v101 := GenS2{F0: 1, F1: (-56 >= -38), F2: 286}
                                _ = v100
                                _ = v101
                            } else {
                                v1 = append(v1, 42)
                            }
                            var v102 uint8 = 128
                            _ = fi77
                            _ = v102
                            _ = v79
                        }
                        var v103 uint16 = 255
                        switch fi76 {
                        case 4294967295:
                            v4 = v4
                        case 32768:
                            _ = len(v3)
                        }
                        _ = fi76
                        _ = v103
                    }
                }
                var v104 int32 = -233
                _ = v104
            } else {
                for fi105 := 0; fi105 < 16; fi105++ {
                    var v106 uint8 = 180
                    _ = fi105
                    _ = v106
                }
                for fi107 := 0; fi107 < 7; fi107++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch -2147483648 {
                    case -898:
                        var v108 int32 = 1
                        _ = v108
                    case 1:
                        v109 := GenS10{F0: -146}
                        _ = v109
                    case 32767:
                        v4 = v4
                    default:
                        v4 = (752 < 549)
                    }
                    _ = fi107
                }
                if false {
                    v110 := GenS3{F0: 9223372036854775807, F1: 255}
                    _ = v110
                } else {
                    for fi111 := 0; fi111 < 3; fi111++ {
                        fmt.Println(fi111, fi111)
                        _ = fi111
                    }
                    v4 = (v4 || v4)
                    _ = len(v1)
                }
            }
        } else {
            v4 = (32767 >= -1)
            if ((v4 || (float64(-2.25) != (3.14159 - 0.0))) && (!(2147483647 == -734))) {
                if len(v1) > 0 {
                    v1[0] = 65535
                }
                if ((-2.25 * float32((-2.25 * 3.14159))) <= -2.25) {
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        var v112 int16 = 42
                        _ = v112
                    }
                    var v113 uint = 255
                    _ = v113
                } else {
                    if (32767 < 589) {
                        v4 = (255 < 476)
                        for fi114 := 0; fi114 < 7; fi114++ {
                            switch -285 {
                            case 32768:
                                v3 = append(v3, 65535)
                            case 128:
                                for fi115 := 0; fi115 < 1; fi115++ {
                                    v4 = (32767 >= 0)
                                    var v116 uint64 = 188
                                    if ((202 < 1) && true) {
                                        var v117 float64 = -2.25
                                        v116 -= 2
                                        _ = v117
                                    } else {
                                        v118 := GenS1{F0: 520, F1: int16(fi114)}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v4 = (1 >= 272)
                                        v4 = v4
                                        var v119 uint32 = 421
                                        _ = v118
                                        _ = v119
                                    }
                                    v120 := GenS11{F0: v4, F1: 1, F2: (^(v116))}
                                    _ = fi115
                                    _ = v116
                                    _ = v120
                                }
                            }
                            v4 = (3.14159 >= 1e10)
                            fmt.Println(v4, fi114)
                            _ = fi114
                        }
                        if v4 {
                            for fi121 := 0; fi121 < 2; fi121++ {
                                fmt.Println(v4, v4)
                                v4 = (1.5 <= 1e10)
                                for fi122 := 0; fi122 < 8; fi122++ {
                                    var v123 int32 = -1
                                    v4 = ((v123 >> ((18446744073709551615) & 63)) <= v123)
                                    v4 = v4
                                    _ = fi122
                                    _ = v123
                                }
                                _ = fi121
                            }
                            for fi124 := 0; fi124 < 4; fi124++ {
                                v125 := GenS8{F0: "\n\t", F1: 201}
                                v126 := GenS12{F0: 1000, F1: fi124}
                                var v127 uint16 = 1
                                _ = fi124
                                _ = v125
                                _ = v126
                                _ = v127
                            }
                            v4 = v4
                        } else {
                            v4 = ((("\n\t" + "") + "a") != "\n\t")
                        }
                    }
                }
                for fi128 := 0; fi128 < 6; fi128++ {
                    fmt.Println(v4, fi128)
                    _ = fi128
                }
            }
            v4 = (811 <= 725)
        }
        for fi129 := 0; fi129 < 8; fi129++ {
            if (!((false || (84 < 255)) && (!(2 >= 193)))) {
                if v4 {
                    for fi130 := 0; fi130 < 11; fi130++ {
                        var v131 float64 = 1.5
                        if (!v4) {
                            var v132 uint16 = 421
                            v133 := GenS13{F0: 32768}
                            _ = len(v1)
                            _ = v132
                            _ = v133
                        } else {
                            v131 += v131
                        }
                        if true {
                            _ = len(v3)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi130
                        _ = v131
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v134 := GenS14{F0: -32768, F1: 137}
                    if true {
                        if v4 {
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                        }
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v135 float32 = ((0.1 - (3.14159 + (1.5 / 1e10))) * 3.14159)
                        var v136 float64 = 0.0
                        _ = v135
                        _ = v136
                    }
                    _ = v134
                }
            }
            v4 = v4
            for fi137 := 0; fi137 < 9; fi137++ {
                var v138 uint16 = 710
                if (105 == 968) {
                    fmt.Println(fi129, fi129, v4)
                    var v139 uint16 = (^(v138))
                    _ = v139
                }
                if v4 {
                    v140 := GenS15{F0: 485, F1: -1}
                    v141 := GenS1{F0: 539, F1: 1}
                    _ = v140
                    _ = v141
                } else {
                    fmt.Println(v4, fi137, v138)
                    var v142 uint = 4294967295
                    _ = v142
                }
                _ = fi137
                _ = v138
            }
            _ = fi129
        }
    }
    for fi143 := 0; fi143 < 5; fi143++ {
        if len(v1) > 0 {
            v1[0] = 0
        }
        _ = fi143
    }
    for fi144 := 0; fi144 < 6; fi144++ {
        switch (fi144 << ((4294967295) & 63)) {
        case 530:
            if ((-(fi144)) == 128) {
                fmt.Println(v4, v4)
                v145 := GenS16{F0: 0, F1: (42 != -2147483648)}
                _ = v145
            }
        }
        _ = fi144
    }
    for fi146 := 0; fi146 < 13; fi146++ {
        fmt.Println(fi146)
        if v4 {
            v147 := GenS17{F0: 42, F1: -41, F2: -806}
            _ = v147
        }
        _ = fi146
    }
    switch 1 {
    case 42:
        switch 32768 {
        case 4294967295:
            var v148 int16 = 0
            _ = v148
        case 127:
            if true {
                fmt.Println(v4)
                v4 = (!v4)
            }
        case -1:
            v4 = (4294967295 < 256)
        default:
            v4 = (88 != 1)
        }
    case 32768:
        for fi149 := 0; fi149 < 4; fi149++ {
            if len(v3) > 0 {
                v3[0] = 65535
            }
            _ = fi149
        }
    }
    var v150 int8 = -48
    var v151 float64 = ((0.0 / 1.5) / 1e10)
    if (0.0 < v151) {
        v150 += 127
        var v152 uint = 97
        _ = v152
    }
    if (v4 && (0.1 < 3.14159)) {
        v150 += 100
        fmt.Println(v151, v150)
    } else {
        fmt.Println(v150, v4, v151)
        v150 = (-(v150))
    }
    for fi153 := 0; fi153 < 4; fi153++ {
        v4 = true
        _ = fi153
    }
    if (!(true || v4)) {
        var v154 float32 = (((float32(v151) / -2.25) + 0.0) * ((3.14159 + 1.5) / 3.14159))
        _ = v154
    } else {
        fmt.Println(v4, v151)
    }
    v150 -= (v150 * v150)
    switch (("" + ("a" + ("\n\t" + "a"))) + "") {
    case "\n\t":
        for fi155 := 0; fi155 < 8; fi155++ {
            var v156 int64 = -1
            fmt.Println(v151, v4)
            fmt.Println(v156, v151, v150)
            _ = fi155
            _ = v156
        }
    case "":
        v150 -= (v150 | 42)
    default:
        fmt.Println(v150)
    }
    for fi157 := 0; fi157 < 10; fi157++ {
        var v158 string = ("a" + ("foo bar" + "\n\t"))
        var v159 float64 = v151
        for fi160 := 0; fi160 < 8; fi160++ {
            fmt.Println(v158, v151)
            if (((float64(3.14159) + (3.14159 - 3.14159)) * 1.5) != float64(-2.25)) {
                if v4 {
                    switch (fi157 ^ (fi160 | fi157)) {
                    case 977:
                        var v161 int = -113
                        fmt.Println(fi157)
                        switch v158 {
                        case "hello":
                            if v4 {
                                v151 = 0.0
                                var v162 int = fi160
                                _ = v162
                            }
                        case "":
                            v3 = append(v3, 0)
                        case "foo bar":
                            v151 += 0.1
                        default:
                            v163 := GenS18{F0: -1, F1: (v150 << ((870) & 63)), F2: 42}
                            _ = v163
                        }
                        var v164 float64 = v159
                        if (2 > 728) {
                            v151 += v164
                            var v165 float64 = (v164 + 0.1)
                            _ = v165
                        }
                        var v166 int = fi157
                        _ = v161
                        _ = v164
                        _ = v166
                    case 2147483647:
                        if (4294967295 >= 42) {
                            var v167 uint = 285
                            fmt.Println(v167, fi160, v4)
                            var v168 uint64 = 752
                            _ = v167
                            _ = v168
                        } else {
                            if (2147483647 != -2147483648) {
                                v150 = (v150 & -128)
                                var v169 float32 = (3.14159 - 3.14159)
                                var v170 int32 = -1
                                _ = v169
                                _ = v170
                            }
                        }
                    case 9223372036854775807:
                        fmt.Println(fi157, v150)
                    }
                    if (!false) {
                        v171 := GenS5{F0: -570, F1: 42}
                        var v172 string = v158
                        _ = len(v3)
                        v173 := GenS19{F0: 1}
                        if len(v1) > 0 {
                            v1[0] = 933
                        }
                        if v4 {
                            v174 := GenS14{F0: -712, F1: 485}
                            if true {
                                var v175 uint32 = 4294967295
                                _ = v175
                            }
                            var v176 int8 = v150
                            var v177 uint32 = 2147483647
                            _ = v174
                            _ = v176
                            _ = v177
                        }
                        if (fi160 > fi157) {
                            fmt.Println(v159, v172)
                            v151 += v159
                        }
                        v4 = (v4 && (v4 || (v150 != v150)))
                        var v178 int32 = -367
                        _ = v171
                        _ = v172
                        _ = v173
                        _ = v178
                    } else {
                        v179 := GenS13{F0: fi157}
                        v3 = append(v3, 2)
                        v180 := GenS2{F0: -356, F1: ((((float32((0.1 * v159)) / float32(float64(float32(((0.0 + ((v159 - v159) - (float64(0.0) + -2.25))) / 1.5))))) / 1e10) / 1e10) != -2.25), F2: 0}
                        switch v158 {
                        case "":
                            fmt.Println(v150, v4)
                        case "hello":
                            var v181 int16 = -1
                            _ = v181
                        }
                        _ = v179
                        _ = v180
                    }
                    switch ((v158 + "a") + v158) {
                    case "\n\t":
                        var v182 float32 = 1.5
                        _ = v182
                    case "":
                        v150 += (v150 ^ (v150 ^ (v150 >> ((551) & 63))))
                    }
                } else {
                    v4 = v4
                }
                v183 := GenS10{F0: -32768}
                if (-1 != (fi157 << ((255) & 63))) {
                    fmt.Println(v158, fi160, v151)
                } else {
                    if (4294967295 != 861) {
                        if v4 {
                            var v184 float64 = ((0.0 * v151) / (((1.5 + -2.25) - 0.1) + float64(-2.25)))
                            var v185 uint8 = 42
                            _ = v184
                            _ = v185
                        } else {
                            var v186 int32 = -852
                            var v187 float64 = v151
                            var v188 int32 = (v186 * -479)
                            _ = v186
                            _ = v187
                            _ = v188
                        }
                        if (((!true) || v4) && v4) {
                            fmt.Println(fi160, fi160)
                        } else {
                            fmt.Println(v158, fi157)
                        }
                    } else {
                        v189 := GenS20{F0: 65535, F1: (v150 | 0), F2: fi160}
                        fmt.Println(fi157, v159)
                        v159 -= v151
                        var v190 bool = (!v4)
                        _ = v189
                        _ = v190
                    }
                }
                _ = v183
            }
            v191 := GenS21{F0: 2, F1: 470, F2: 42}
            _ = fi160
            _ = v191
        }
        _ = fi157
        _ = v158
        _ = v159
    }
    fmt.Println(v151, v4, v151)
    var v192 uint8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v150:", v150)
    fmt.Println("v151:", v151)
    fmt.Println("v192:", v192)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v150
    _ = v151
    _ = v192
    _ = v2
    _ = v3
    _ = v4
}
