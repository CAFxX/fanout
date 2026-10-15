package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 string }
type GenS2 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 uint32; F1 string; F2 uint8 }
type GenS9 struct { F0 uint32; F1 uint64; F2 uint }
type GenS10 struct { F0 uint8; F1 uint16 }
type GenS11 struct { F0 int8; F1 string }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint64; F1 bool; F2 uint }
func genFn0(p0 int64, p1 int32, p2 uint) int {
    fmt.Println(p1, p0)
    if true {
        v7 := []int{int(p2), 32767, -9223372036854775808, 670}
        if false {
            v8 := GenS1{F0: 656, F1: ((3.14159 * float64(0.0)) >= 3.14159), F2: ("foo bar" + "")}
            var v9 uint16 = 42
            switch 9223372036854775807 {
            case 1:
                var v10 string = "hello"
                _ = v10
            case 374:
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 uint16 = (v9 % ((32767) | 1))
                    _ = fi11
                    _ = v12
                }
            }
            _ = v8
            _ = v9
        } else {
            if len(v7) > 0 {
                v7[0] = -1
            }
        }
        for fi13 := 0; fi13 < 8; fi13++ {
            if len(v7) > 0 {
                v7[0] = fi13
            }
            var v14 bool = false
            switch fi13 {
            case 127:
                var v15 int = fi13
                _ = v15
            case 0:
                var v16 string = ("foo bar" + "\n\t")
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        _ = v7
    } else {
        v17 := []int{}
        _ = v17
    }
    if (!(559 != 2)) {
        for fi18 := 0; fi18 < 9; fi18++ {
            p2 = (p2 / ((206) | 1))
            var v19 uint64 = 189
            switch fi18 {
            case -170:
                if (p2 <= (p2 | p2)) {
                    p2 -= p2
                }
            default:
                v20 := GenS2{F0: 9223372036854775807, F1: 40, F2: 38}
                _ = v20
            }
            _ = fi18
            _ = v19
        }
    } else {
        var v21 int = 4294967295
        _ = v21
    }
    var v22 uint64 = 895
    if (-219 == 1000) {
        switch 2147483647 {
        case -939:
            var v23 float32 = 1e10
            _ = v23
        default:
            switch ("hello" + ("hello" + (("hello" + (("hello" + "\n\t") + "")) + "foo bar"))) {
            case "foo bar":
                p2 = 0
            case "":
                fmt.Println(p0)
            case "\n\t":
                if true {
                    v24 := []int{-2147483648, -9223372036854775808, 557, -1}
                    switch 533 {
                    case 966:
                        v25 := GenS2{F0: p0, F1: 1, F2: 0}
                        _ = v25
                    case 673:
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                    case 42:
                        for fi26 := 0; fi26 < 3; fi26++ {
                            for fi27 := 0; fi27 < 17; fi27++ {
                                switch (^(fi27)) {
                                case 188:
                                    p2 += p2
                                default:
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                }
                                if false {
                                    var v28 uint16 = 42
                                    var v29 string = (((("hello" + (("foo bar" + "") + "")) + "a") + "foo bar") + "\n\t")
                                    p2 = 689
                                    _ = v28
                                    _ = v29
                                }
                                v22 -= (^(v22))
                                _ = fi27
                            }
                            p2 = (p2 / ((25) | 1))
                            var v30 int64 = p0
                            _ = fi26
                            _ = v30
                        }
                    default:
                        fmt.Println(v22, p2, p1)
                    }
                    for fi31 := 0; fi31 < 9; fi31++ {
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                        if ((p0 % ((p0) | 1)) < 392) {
                            v24 = append(v24, (fi31 % ((fi31) | 1)))
                            v24 = append(v24, (fi31 - fi31))
                        }
                        var v32 bool = ((731 <= 42) || ((3.14159 - 1.5) == 1.5))
                        _ = fi31
                        _ = v32
                    }
                    _ = v24
                }
            }
        }
        switch 0 {
        case 127:
            var v33 float32 = 3.14159
            _ = v33
        case 9223372036854775807:
            v34 := GenS3{F0: 723, F1: -1}
            _ = v34
        case -1:
            p1 -= (p1 % ((p1) | 1))
        }
    } else {
        p0 = int64(v22)
        for fi35 := 0; fi35 < 7; fi35++ {
            fmt.Println(fi35, p1)
            _ = fi35
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{}
    v3 := make(map[uint32]bool)
    var v4 uint32 = 1
    var v5 float64 = float64(((1e10 * (((1e10 / float32((-2.25 * (3.14159 * float64(1.5))))) * 0.0) + (0.1 - -2.25))) + 1e10))
    var v6 int32 = 0
    var v36 float32 = 0.0
    fmt.Println(v36, v36, v4)
    var v37 bool = ((v4 >> ((574) & 63)) <= (v4 << ((18446744073709551615) & 63)))
    var v38 int8 = 100
    var v39 float64 = float64(float32(((1e10 - (float64(v36) / v5)) - v5)))
    for fi40 := 0; fi40 < 18; fi40++ {
        v41 := GenS3{F0: v4, F1: fi40}
        if (41 < 565) {
            if ((1 > 1) && (((v37 && ((v38 * v38) < v38)) || v37) || (730 >= 772))) {
                var v42 uint32 = (v4 ^ v4)
                for fi43 := 0; fi43 < 20; fi43++ {
                    v37 = (1 < 1)
                    var v44 float64 = 3.14159
                    var v45 int = (fi43 + genFn0(180, int32(v38), 978))
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                _ = v42
            } else {
                for fi46 := 0; fi46 < 6; fi46++ {
                    v6 = (-(v6))
                    var v47 int8 = v38
                    _ = fi46
                    _ = v47
                }
                fmt.Println(v5, v4, v36)
            }
            v2 = append(v2, ("hello" + "a"))
            var v48 int8 = -127
            _ = v48
        }
        v4 -= (^(v4))
        _ = fi40
        _ = v41
    }
    var v49 uint32 = v4
    v37 = v37
    v39 = (-2.25 + (v5 * (1e10 + 1.5)))
    switch ("a" + "foo bar") {
    case "foo bar":
        for fi50 := 0; fi50 < 4; fi50++ {
            if (!(42 < -651)) {
                _ = len(v3)
            } else {
                v5 = v5
                for fi51 := 0; fi51 < 5; fi51++ {
                    v39 = float64((1.5 / 1e10))
                    _ = fi51
                }
            }
            var v52 uint8 = 127
            fmt.Println(v6, v37, v6)
            _ = fi50
            _ = v52
        }
    case "\n\t":
        fmt.Println(v39)
    }
    v37 = (18446744073709551615 >= 200)
    switch ("a" + (("hello" + ("foo bar" + "foo bar")) + "a")) {
    case "":
        for fi53 := 0; fi53 < 16; fi53++ {
            if (42 >= 0) {
                if (((fi53 % ((fi53) | 1)) <= 128) && (-101 <= v38)) {
                    var v54 float32 = float32((3.14159 + v5))
                    v55 := GenS2{F0: int64(v38), F1: 128, F2: 128}
                    _ = v54
                    _ = v55
                }
                var v56 uint8 = 3
                _ = v56
            }
            var v57 string = "foo bar"
            if v37 {
                if v37 {
                    if false {
                        if (32767 < 0) {
                            fmt.Println(v4, v4)
                            if ((v37 && (203 == 1)) && true) {
                                var v58 float32 = (-2.25 * v36)
                                if v37 {
                                    var v59 uint8 = 2
                                    v60 := GenS4{F0: uint32(v59), F1: (v4 | 358)}
                                    _ = v59
                                    _ = v60
                                }
                                for fi61 := 0; fi61 < 18; fi61++ {
                                    if true {
                                        if (v36 > v36) {
                                            if (!false) {
                                                var v62 uint64 = 18446744073709551615
                                                v63 := GenS2{F0: 42, F1: uint8(v4), F2: 255}
                                                _ = v62
                                                _ = v63
                                            }
                                            if v37 {
                                                var v64 int = (^(fi53))
                                                v65 := GenS4{F0: (v49 * (^(v49))), F1: (v49 % ((2) | 1))}
                                                var v66 float32 = (v36 + (v36 - ((v36 / 1.0) / float32(-2.25))))
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            }
                                        }
                                        var v67 uint8 = 1
                                        _ = v67
                                    } else {
                                        v39 += v5
                                    }
                                    _ = fi61
                                }
                                _ = v58
                            }
                        } else {
                            if (v5 >= 3.14159) {
                                fmt.Println(v37)
                                v2 = append(v2, v57)
                                if v37 {
                                    var v68 float64 = 1.5
                                    _ = v68
                                } else {
                                    var v69 float32 = v36
                                    v70 := GenS1{F0: 65535, F1: true, F2: "a"}
                                    v37 = false
                                    _ = v69
                                    _ = v70
                                }
                            } else {
                                var v71 int64 = 512
                                _ = v71
                            }
                        }
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v36, v57, v5)
                        }
                        _ = len(v1)
                    }
                    v38 = v38
                } else {
                    v39 += (v39 * v5)
                }
                v4 = 42
                var v72 int8 = v38
                _ = v72
            } else {
                switch (fi53 ^ genFn0(42, v6, 65535)) {
                case 4294967295:
                    var v73 float64 = (float64(v36) - (0.1 * v39))
                    _ = v73
                default:
                    for fi74 := 0; fi74 < 5; fi74++ {
                        var v75 uint = 4294967295
                        v6 -= (v6 >> ((v75) & 63))
                        _ = fi74
                        _ = v75
                    }
                }
            }
            _ = fi53
            _ = v57
        }
    case "hello":
        v76 := GenS3{F0: v49, F1: 128}
        _ = v76
    case "a":
        v77 := GenS4{F0: (v49 << ((231) & 63)), F1: (v49 - v49)}
        _ = v77
    default:
        if (4294967295 == 766) {
            for fi78 := 0; fi78 < 15; fi78++ {
                switch ("foo bar" + (("" + ("a" + "")) + "hello")) {
                case "\n\t":
                    for fi79 := 0; fi79 < 17; fi79++ {
                        var v80 uint32 = (v4 | (^(v4)))
                        v37 = v37
                        v81 := GenS3{F0: (v4 - 735), F1: (fi79 ^ genFn0(42, v6, 65535))}
                        _ = fi79
                        _ = v80
                        _ = v81
                    }
                default:
                    var v82 uint64 = 42
                    _ = v82
                }
                _ = fi78
            }
        } else {
            var v83 int8 = v38
            _ = v83
        }
    }
    var v84 int32 = v6
    if v37 {
        v3[v49] = (42 != 1)
        for fi85 := 0; fi85 < 11; fi85++ {
            var v86 int8 = int8(fi85)
            if (32767 != 32767) {
                for fi87 := 0; fi87 < 8; fi87++ {
                    switch (fi87 ^ 32768) {
                    case 9223372036854775807:
                        var v88 int16 = -32768
                        switch (-(fi87)) {
                        case -760:
                            if v37 {
                                if v37 {
                                    v38 += 127
                                    v6 = 0
                                    switch 32767 {
                                    case -1:
                                        fmt.Println(v88, fi87, v86)
                                    case 1:
                                        v89 := GenS4{F0: v49, F1: v4}
                                        _ = v89
                                    default:
                                        v90 := GenS5{F0: 2}
                                        _ = v90
                                    }
                                } else {
                                    var v91 bool = ("\n\t" < ("a" + ("\n\t" + ("hello" + ("a" + "")))))
                                    v4 -= (^(v49))
                                    _ = v91
                                }
                                v39 = v5
                                var v92 bool = v37
                                _ = v92
                            }
                        default:
                            v93 := GenS1{F0: 2, F1: (("a" + "hello") < ("a" + "a")), F2: "foo bar"}
                            _ = v93
                        }
                        switch "foo bar" {
                        case "\n\t":
                            v37 = (v37 && v37)
                        case "":
                            if false {
                                var v94 uint32 = v4
                                var v95 uint64 = 1
                                _ = v94
                                _ = v95
                            }
                        }
                        _ = v88
                    case -2147483648:
                        if true {
                            fmt.Println(v38)
                        }
                    case -738:
                        switch fi87 {
                        case 128:
                            var v96 uint = 556
                            _ = v96
                        case 9223372036854775807:
                            if false {
                                fmt.Println(v49)
                                var v97 int = genFn0(-431, (v6 << ((813) & 63)), 2)
                                var v98 uint64 = 700
                                _ = v97
                                _ = v98
                            } else {
                                v49 -= v49
                                var v99 int8 = v38
                                _ = v99
                            }
                        case -278:
                            if (((v39 * (v39 - (((v39 / float64(v36)) * v5) * v39))) * 3.14159) >= v39) {
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                                v100 := GenS6{F0: 127, F1: (^(fi85))}
                                v38 -= int8(v49)
                                _ = v100
                            } else {
                                fmt.Println(v84, v86, v36)
                                if v37 {
                                    v2 = append(v2, "a")
                                    var v101 int = genFn0(0, (^(v6)), 0)
                                    v102 := GenS7{F0: (v4 - (^(v4))), F1: (127 == 128)}
                                    _ = v101
                                    _ = v102
                                } else {
                                    v4 += v49
                                }
                                v39 -= v5
                                if v37 {
                                    fmt.Println(fi87, fi87)
                                    if v37 {
                                        if true {
                                            v103 := GenS7{F0: (v49 + 4294967295), F1: ((v86 / ((127) | 1)) < (v38 / ((v86) | 1)))}
                                            var v104 int64 = -378
                                            _ = v103
                                            _ = v104
                                        }
                                        fmt.Println(v86)
                                    }
                                    var v105 string = "\n\t"
                                    var v106 int16 = -881
                                    _ = v105
                                    _ = v106
                                } else {
                                    v4 = (v49 << ((267) & 63))
                                    var v107 uint16 = 2
                                    _ = v107
                                }
                                v6 = v6
                            }
                        }
                    default:
                        if (v37 && (-436 < 0)) {
                            var v108 uint32 = v4
                            switch "a" {
                            case "\n\t":
                                v109 := GenS6{F0: 227, F1: fi85}
                                _ = v109
                            case "":
                                v38 = v86
                            default:
                                if false {
                                    v110 := GenS8{F0: (^(v108)), F1: "", F2: 127}
                                    v49 += v49
                                    _ = v110
                                }
                                var v111 uint = 255
                                switch "\n\t" {
                                case "":
                                    var v112 uint32 = (^(v108))
                                    var v113 uint64 = 4294967295
                                    v114 := GenS6{F0: 176, F1: (fi87 - fi85)}
                                    _ = v112
                                    _ = v113
                                    _ = v114
                                case "hello":
                                    var v115 uint32 = v108
                                    var v116 uint = 65535
                                    fmt.Println(v116, v116, fi85)
                                    v5 -= (((v5 + 0.0) / v5) * 0.0)
                                    _ = v115
                                    _ = v116
                                default:
                                    var v117 float64 = (v39 + (((float64(v36) / ((v5 * v5) / v39)) / ((float64(v36) + float64(float32((1.5 + v5)))) * 1e10)) / ((v5 - v5) + 1e10)))
                                    _ = v117
                                }
                                _ = v111
                            }
                            if true {
                                var v118 int64 = 42
                                fmt.Println(v39, fi87, v84)
                                _ = v118
                            }
                            switch "" {
                            case "foo bar":
                                if false {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = "\n\t"
                                    }
                                    switch "\n\t" {
                                    case "a":
                                        var v119 uint32 = (^(v108))
                                        _ = v119
                                    case "\n\t":
                                        v4 = (v4 & v108)
                                    }
                                    switch fi85 {
                                    case 32767:
                                        switch (fi87 * (-(fi85))) {
                                        case 128:
                                            switch ("\n\t" + "foo bar") {
                                            case "hello":
                                                v120 := GenS9{F0: (v49 & 42), F1: 4294967295, F2: 782}
                                                _ = v120
                                            case "":
                                                if v37 {
                                                    var v121 int = (fi87 - (fi85 % ((int(v6)) | 1)))
                                                    var v122 string = (((("a" + ("a" + "a")) + "foo bar") + (("" + "") + "\n\t")) + "a")
                                                    var v123 int8 = (v38 / (((v38 + (v86 << ((828) & 63)))) | 1))
                                                    _ = v121
                                                    _ = v122
                                                    _ = v123
                                                }
                                            }
                                        }
                                    case -2147483648:
                                        if false {
                                            var v124 float64 = ((v39 - v5) * float64(((float32(v39) / float32((((0.1 / (v5 * v39)) + v5) + (v5 - ((v5 * v5) - v39))))) / v36)))
                                            var v125 int64 = 846
                                            _ = v124
                                            _ = v125
                                        }
                                    case 1:
                                        var v126 int8 = (v38 + v38)
                                        _ = v126
                                    default:
                                        v127 := GenS10{F0: 255, F1: 72}
                                        _ = v127
                                    }
                                    switch "" {
                                    case "foo bar":
                                        v86 += (v86 % ((42) | 1))
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                } else {
                                    var v128 int16 = 35
                                    fmt.Println(v108)
                                    var v129 uint32 = v108
                                    _ = v128
                                    _ = v129
                                }
                            }
                            _ = v108
                        }
                    }
                    _ = fi87
                }
            }
            _ = fi85
            _ = v86
        }
        v39 = 1e10
    } else {
        fmt.Println(v84, v84)
    }
    v39 -= v39
    for fi130 := 0; fi130 < 17; fi130++ {
        fmt.Println(fi130)
        if (!(!v37)) {
            fmt.Println(v5)
            switch "hello" {
            case "hello":
                v5 = 1.5
            case "":
                if len(v2) > 0 {
                    v2[0] = "a"
                }
            default:
                if (42 >= 799) {
                    v37 = ((v6 / ((v6) | 1)) >= (v6 & v84))
                    fmt.Println(v4)
                } else {
                    if ((v6 - v84) > (v6 >> ((449) & 63))) {
                        for fi131 := 0; fi131 < 5; fi131++ {
                            fmt.Println(v36, v84)
                            if (!v37) {
                                var v132 int32 = 1
                                if (!(1 >= 0)) {
                                    var v133 int16 = 1
                                    if v37 {
                                        var v134 string = "a"
                                        switch -2147483648 {
                                        case -269:
                                            switch (v134 + v134) {
                                            case "a":
                                                fmt.Println(v133)
                                            case "":
                                                if v37 {
                                                    v37 = true
                                                    fmt.Println(v39, v36)
                                                    v5 += ((3.14159 + v5) * (float64(v36) * v5))
                                                }
                                            case "foo bar":
                                                if ((fi130 != 128) && false) {
                                                    var v135 int8 = v38
                                                    if v37 {
                                                        v136 := GenS11{F0: v38, F1: "a"}
                                                        _ = v136
                                                    } else {
                                                        v4 += (v4 << ((2) & 63))
                                                        fmt.Println(v132)
                                                    }
                                                    if ((v132 ^ (v132 - v6)) == v84) {
                                                        v134 = v134
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                    _ = v135
                                                } else {
                                                    if v37 {
                                                        v1[2147483647] = (!((v36 <= float32(v5)) || false))
                                                        var v137 uint = 0
                                                        v4 += 1
                                                        _ = v137
                                                    }
                                                }
                                            }
                                        case -664:
                                            _ = len(v2)
                                        case 32767:
                                            if v37 {
                                                var v138 int64 = 1
                                                v139 := GenS6{F0: 93, F1: (fi131 + (fi131 ^ int(v138)))}
                                                fmt.Println(fi130)
                                                _ = v138
                                                _ = v139
                                            } else {
                                                var v140 string = v134
                                                _ = v140
                                            }
                                        default:
                                            switch (("a" + v134) + "foo bar") {
                                            case "hello":
                                                if v37 {
                                                    v141 := GenS12{F0: (^(v4))}
                                                    _ = v141
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                            default:
                                                fmt.Println(v49)
                                                v5 += 0.0
                                            }
                                        }
                                        _ = v134
                                    } else {
                                        fmt.Println(v5, v36, v132)
                                    }
                                    _ = v133
                                }
                                switch fi131 {
                                case 782:
                                    fmt.Println(v132, v37, v4)
                                case -9223372036854775808:
                                    fmt.Println(v5)
                                }
                                _ = v132
                            } else {
                                switch (("foo bar" + ((("\n\t" + ("\n\t" + "hello")) + "\n\t") + ("a" + ("foo bar" + ("hello" + (("hello" + "") + "")))))) + "foo bar") {
                                case "":
                                    switch fi130 {
                                    case 913:
                                        var v142 int64 = -983
                                        _ = v142
                                    case 655:
                                        var v143 int32 = v84
                                        _ = v143
                                    case 152:
                                        v144 := GenS13{F0: 4294967295, F1: ("foo bar" >= "foo bar"), F2: 859}
                                        _ = v144
                                    default:
                                        v39 = ((float64((float32(v5) + float32(-2.25))) * (v5 - 3.14159)) + (v39 / ((((v5 + (v5 / 1.0)) * v5) + 1.5) * ((v5 * 1.5) + 0.0))))
                                    }
                                }
                                var v145 uint32 = (^(v4))
                                var v146 int = 32767
                                _ = v145
                                _ = v146
                            }
                            _ = len(v2)
                            _ = fi131
                        }
                        _ = len(v1)
                    } else {
                        v147 := GenS13{F0: 717, F1: (!v37), F2: 17}
                        var v148 uint8 = 145
                        if (float32(v39) != (float32(v39) / v36)) {
                            var v149 float32 = v36
                            _ = v149
                        }
                        _ = v147
                        _ = v148
                    }
                }
            }
        } else {
            v5 -= (0.1 + (v39 + v39))
            v3[(v4 & (v49 ^ (v49 / (((^(v49))) | 1))))] = v37
        }
        for fi150 := 0; fi150 < 1; fi150++ {
            fmt.Println(v4, v6)
            v2 = append(v2, ("\n\t" + ("foo bar" + "hello")))
            switch (fi130 % ((fi130) | 1)) {
            case -1:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                if len(v2) > 0 {
                    v2[0] = ""
                }
                switch fi150 {
                case 127:
                    var v151 uint8 = 1
                    _ = v151
                default:
                    fmt.Println(v49, v6, v84)
                }
            }
            _ = len(v2)
            _ = fi150
        }
        _ = fi130
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v84
}
