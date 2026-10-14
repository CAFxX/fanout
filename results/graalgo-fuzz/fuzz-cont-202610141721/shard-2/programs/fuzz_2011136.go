package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 string; F2 int8 }
type GenS7 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS8 struct { F0 uint8; F1 int64 }
type GenS9 struct { F0 bool; F1 int; F2 string }
type GenS10 struct { F0 uint; F1 uint32 }
type GenS11 struct { F0 int8; F1 string; F2 int8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int32; F1 bool }
type GenS14 struct { F0 uint32; F1 uint8 }
type GenS15 struct { F0 uint; F1 int32 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 int8; F1 uint64; F2 bool }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS20 struct { F0 uint16; F1 int; F2 uint8 }
type GenS21 struct { F0 string; F1 int32 }
func genFn0(p0 uint16, p1 uint8) int {
    if false {
        p1 = (p1 | p1)
    } else {
        p1 += 255
        v5 := GenS1{F0: -1}
        _ = v5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        if false {
            for fi7 := 0; fi7 < 6; fi7++ {
                switch (fi6 | (-(fi7))) {
                case 42:
                    fmt.Println(p1, fi7)
                case 2147483647:
                    fmt.Println(fi7, fi6)
                case 127:
                    p1 -= 42
                default:
                    if true {
                        var v8 float64 = 3.14159
                        _ = v8
                    } else {
                        var v9 uint16 = 32767
                        v9 = p0
                        _ = v9
                    }
                }
                p1 -= p1
                _ = fi7
            }
            v10 := GenS2{F0: ("\n\t" + ("foo bar" + "foo bar"))}
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 uint = 2
                p0 += p0
                var v13 uint = v12
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        } else {
            p0 = p0
            if false {
                for fi14 := 0; fi14 < 12; fi14++ {
                    v15 := []int{}
                    _ = fi14
                    _ = v15
                }
                if ((!((3.14159 <= (-2.25 - 0.1)) && (127 < -95))) || (!true)) {
                    fmt.Println(p1, p1)
                    p0 = (p0 & p0)
                }
            } else {
                p0 = p0
                var v16 float64 = 3.14159
                p1 = 42
                _ = v16
            }
            if (!((!((p0 ^ (p0 << ((2) & 63))) > p0)) && false)) {
                if (false || false) {
                    if (628 == 377) {
                        fmt.Println(p0, p1)
                        if true {
                            for fi17 := 0; fi17 < 8; fi17++ {
                                if (true || true) {
                                    p0 = (p0 ^ p0)
                                    var v18 int32 = 100
                                    _ = v18
                                } else {
                                    fmt.Println(p1, p0)
                                    v19 := GenS3{F0: 2, F1: 50}
                                    _ = v19
                                }
                                fmt.Println(fi6, p1, fi6)
                                if true {
                                    p1 = p1
                                    p1 = uint8(p0)
                                }
                                _ = fi17
                            }
                        }
                        switch (fi6 / (((fi6 >> ((812) & 63))) | 1)) {
                        case -690:
                            v20 := []int{9223372036854775807, fi6, 2147483647, 1}
                            _ = v20
                        case 9223372036854775807:
                            var v21 int8 = 127
                            _ = v21
                        }
                    }
                    p0 += 2
                    p1 = (^(p1))
                } else {
                    var v22 uint16 = 42
                    for fi23 := 0; fi23 < 6; fi23++ {
                        var v24 int = (fi23 & -725)
                        v25 := []int{(fi23 * (fi6 & (fi23 & (fi6 * (v24 * 32767)))))}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    var v26 int8 = 42
                    _ = v22
                    _ = v26
                }
            }
        }
        var v27 string = "a"
        _ = fi6
        _ = v27
    }
    p0 += (p0 >> ((255) & 63))
    for fi28 := 0; fi28 < 17; fi28++ {
        if false {
            v29 := GenS4{F0: (p0 >> ((42) & 63)), F1: true}
            switch fi28 {
            case 32767:
                var v30 float32 = float32(-2.25)
                _ = v30
            default:
                if false {
                    v31 := []int{fi28}
                    p0 = p0
                    _ = v31
                } else {
                    var v32 float32 = ((0.0 - 1.5) + 1e10)
                    _ = v32
                }
            }
            _ = v29
        }
        _ = fi28
    }
    _ = p0
    _ = p1
    return -430
}
func genFn1(p0 int64, p1 uint) uint16 {
    var v33 int = -2147483648
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = v33
    return 65535
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]bool)
    var v3 uint32 = 1
    var v4 int16 = 32767
    v1 += v1
    v4 += 1
    if false {
        switch "hello" {
        case "":
            switch 127 {
            case -2147483648:
                v34 := GenS5{F0: 227}
                _ = v34
            default:
                if false {
                    switch 9223372036854775807 {
                    case 1:
                        v4 = -101
                    case -562:
                        switch 9223372036854775807 {
                        case 127:
                            switch 4294967295 {
                            case 127:
                                switch "\n\t" {
                                case "":
                                    v4 = (-(v4))
                                case "hello":
                                    for fi35 := 0; fi35 < 13; fi35++ {
                                        for fi36 := 0; fi36 < 15; fi36++ {
                                            var v37 int32 = (v1 - (v1 << ((42) & 63)))
                                            _ = fi36
                                            _ = v37
                                        }
                                        var v38 uint = 68
                                        v39 := GenS6{F0: (v3 ^ 2147483648), F1: ("" + "\n\t"), F2: 0}
                                        _ = fi35
                                        _ = v38
                                        _ = v39
                                    }
                                case "\n\t":
                                    for fi40 := 0; fi40 < 18; fi40++ {
                                        if (42 < 402) {
                                            switch fi40 {
                                            case 952:
                                                if ((v1 & v1) >= int32(v3)) {
                                                    var v41 uint64 = 4294967295
                                                    if (1.5 != -2.25) {
                                                        fmt.Println(v4, v41)
                                                        v41 += v41
                                                        fmt.Println(v1, v4)
                                                    }
                                                    for fi42 := 0; fi42 < 10; fi42++ {
                                                        v43 := GenS6{F0: (v3 & v3), F1: "\n\t", F2: 42}
                                                        var v44 uint = 18446744073709551615
                                                        v41 += v41
                                                        _ = fi42
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                    _ = v41
                                                } else {
                                                    if true {
                                                        delete(v2, (v3 * 42))
                                                        v45 := GenS7{F0: 1, F1: -128, F2: 744}
                                                        v46 := []int{127, (-(fi40))}
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                }
                                            case 0:
                                                v4 -= v4
                                            case 127:
                                                fmt.Println(v4)
                                            default:
                                                var v47 int = (fi40 << ((18446744073709551615) & 63))
                                                _ = v47
                                            }
                                            var v48 uint = 1
                                            var v49 int16 = v4
                                            _ = v48
                                            _ = v49
                                        } else {
                                            for fi50 := 0; fi50 < 11; fi50++ {
                                                v51 := []int{fi40, (fi50 & 42), (fi50 & (fi40 | 2147483647))}
                                                _ = fi50
                                                _ = v51
                                            }
                                            v1 = v1
                                            if true {
                                                if true {
                                                    switch (fi40 & genFn0(1, 161)) {
                                                    case 1:
                                                        var v52 int32 = v1
                                                        _ = v52
                                                    case 32768:
                                                        if (true || (genFn1(1, 2) != 65535)) {
                                                            for fi53 := 0; fi53 < 14; fi53++ {
                                                                if (0.1 < 1.5) {
                                                                    if false {
                                                                        fmt.Println(v3)
                                                                    } else {
                                                                        if true {
                                                                            if (!true) {
                                                                                if true {
                                                                                    v54 := []int{127, genFn0(255, 127), (fi40 >> ((2) & 63)), (fi40 + 42)}
                                                                                    switch fi40 {
                                                                                    case -528:
                                                                                        v1 += v1
                                                                                    default:
                                                                                        v4 += (v4 * (-(v4)))
                                                                                    }
                                                                                    _ = v54
                                                                                }
                                                                                fmt.Println(fi53, v3)
                                                                                if true {
                                                                                    v1 -= (v1 >> ((670) & 63))
                                                                                    v55 := []int{(fi53 - fi53)}
                                                                                    fmt.Println(fi53, v3, v1)
                                                                                    _ = v55
                                                                                }
                                                                                var v56 uint = 65535
                                                                                v1 += 1
                                                                                _ = v56
                                                                            }
                                                                            v4 -= v4
                                                                        }
                                                                        if (!((v4 >> ((42) & 63)) == v4)) {
                                                                            var v57 uint16 = 499
                                                                            v58 := []int{genFn0(v57, 8), int(v57), -910}
                                                                            v3 += v3
                                                                            _ = v57
                                                                            _ = v58
                                                                        }
                                                                        if true {
                                                                            if false {
                                                                                if true {
                                                                                    var v59 int64 = 9223372036854775807
                                                                                    var v60 uint = 2
                                                                                    v60 = 18446744073709551615
                                                                                    v61 := GenS5{F0: 86}
                                                                                    _ = v59
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                }
                                                                                v4 = (-(v4))
                                                                                v1 -= (v1 ^ 0)
                                                                            }
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                switch (fi53 & 2147483647) {
                                                                                case 0:
                                                                                    v62 := GenS8{F0: 84, F1: 1}
                                                                                    _ = v62
                                                                                case 564:
                                                                                    var v63 int16 = (v4 | -620)
                                                                                    _ = v63
                                                                                }
                                                                            }
                                                                            v64 := []int{}
                                                                            _ = v64
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v65 int32 = (v1 % (((-(v1))) | 1))
                                                                    delete(v2, v3)
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        var v66 string = "foo bar"
                                                                        _ = v66
                                                                    case "a":
                                                                        if (false && true) {
                                                                            v67 := GenS9{F0: true, F1: (fi53 % ((-903) | 1)), F2: "hello"}
                                                                            var v68 uint32 = (v3 >> ((4294967295) & 63))
                                                                            _ = v67
                                                                            _ = v68
                                                                        } else {
                                                                            v69 := GenS5{F0: 255}
                                                                            _ = v69
                                                                        }
                                                                    case "hello":
                                                                        v1 -= -218
                                                                    }
                                                                    if false {
                                                                        var v70 int = (fi53 >> ((255) & 63))
                                                                        _ = v70
                                                                    } else {
                                                                        var v71 int32 = (^(v1))
                                                                        v1 -= 3
                                                                        var v72 float64 = (1e10 / 1.0)
                                                                        v4 = v4
                                                                        v73 := GenS7{F0: 2, F1: 111, F2: 383}
                                                                        v74 := []int{604, genFn0(982, 255), (fi40 ^ -532), (fi40 & fi40)}
                                                                        _ = v71
                                                                        _ = v72
                                                                        _ = v73
                                                                        _ = v74
                                                                    }
                                                                    _ = v65
                                                                }
                                                                _ = fi53
                                                            }
                                                            v4 += -32768
                                                        }
                                                    case -2147483648:
                                                        if (3.14159 == -2.25) {
                                                            if (!(0 <= -9223372036854775808)) {
                                                                v3 = (v3 * v3)
                                                            }
                                                        }
                                                    }
                                                }
                                                if (2 >= 2) {
                                                    v75 := GenS7{F0: 18446744073709551615, F1: 42, F2: 0}
                                                    var v76 string = (("hello" + ("" + "")) + "hello")
                                                    _ = v75
                                                    _ = v76
                                                } else {
                                                    for fi77 := 0; fi77 < 16; fi77++ {
                                                        v78 := GenS4{F0: 937, F1: (614 < 4294967295)}
                                                        var v79 bool = false
                                                        v79 = ((v79 && (v79 && v79)) || ((18446744073709551615 != 886) || (v3 <= 2147483647)))
                                                        _ = fi77
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                }
                                            }
                                        }
                                        _ = fi40
                                    }
                                default:
                                    for fi80 := 0; fi80 < 19; fi80++ {
                                        var v81 int8 = -56
                                        v82 := []int{genFn0(32767, 42), (fi80 | (fi80 | int(v4)))}
                                        _ = fi80
                                        _ = v81
                                        _ = v82
                                    }
                                }
                            default:
                                v1 -= (v1 >> ((833) & 63))
                            }
                        case -126:
                            switch -267 {
                            case 42:
                                switch "hello" {
                                case "\n\t":
                                    var v83 int32 = (v1 * v1)
                                    _ = v83
                                case "foo bar":
                                    for fi84 := 0; fi84 < 1; fi84++ {
                                        for fi85 := 0; fi85 < 6; fi85++ {
                                            if false {
                                                v3 = (v3 + v3)
                                            } else {
                                                v3 = (v3 + v3)
                                                v3 -= v3
                                                v3 = 2
                                            }
                                            v86 := GenS9{F0: false, F1: (fi84 << ((779) & 63)), F2: "hello"}
                                            v87 := []int{9223372036854775807, fi84, 2147483647}
                                            switch -9223372036854775808 {
                                            case 127:
                                                v1 = 0
                                            }
                                            _ = fi85
                                            _ = v86
                                            _ = v87
                                        }
                                        for fi88 := 0; fi88 < 10; fi88++ {
                                            if ((!true) && true) {
                                                v89 := GenS10{F0: 219, F1: v3}
                                                v90 := GenS5{F0: 2}
                                                fmt.Println(v4, fi88)
                                                v4 = int16(fi88)
                                                var v91 uint64 = 1
                                                _ = v89
                                                _ = v90
                                                _ = v91
                                            }
                                            v92 := []int{(fi88 ^ (fi88 ^ genFn0(509, 255))), fi84, (fi88 | (-(fi84))), genFn0(uint16(v3), 153)}
                                            v2[(^(v3))] = ((v3 | 42) != v3)
                                            _ = fi88
                                            _ = v92
                                        }
                                        v1 = 42
                                        _ = fi84
                                    }
                                }
                            default:
                                v93 := GenS6{F0: v3, F1: ("" + ""), F2: 105}
                                _ = v93
                            }
                        case 32768:
                            fmt.Println(v3)
                        }
                    }
                    switch -1 {
                    case 9223372036854775807:
                        if (-363 >= -2147483648) {
                            fmt.Println(v4, v4)
                            var v94 int32 = (v1 - 1)
                            for fi95 := 0; fi95 < 7; fi95++ {
                                fmt.Println(fi95, v94)
                                v96 := GenS11{F0: 73, F1: "a", F2: 1}
                                var v97 int64 = -856
                                _ = fi95
                                _ = v96
                                _ = v97
                            }
                            _ = v94
                        }
                    case -1:
                        if true {
                            v1 = v1
                        } else {
                            switch ("hello" + ("a" + "\n\t")) {
                            case "foo bar":
                                v98 := []int{-2}
                                _ = v98
                            case "a":
                                var v99 int16 = 726
                                _ = v99
                            default:
                                v100 := GenS5{F0: 42}
                                _ = v100
                            }
                            var v101 uint = 18446744073709551615
                            _ = v101
                        }
                    case 117:
                        var v102 int8 = 1
                        _ = v102
                    }
                    v4 += (v4 | v4)
                } else {
                    switch (("a" + "foo bar") + "foo bar") {
                    case "\n\t":
                        switch 9223372036854775807 {
                        case 1:
                            for fi103 := 0; fi103 < 3; fi103++ {
                                v104 := v2[v3]
                                var v105 uint8 = 2
                                _ = fi103
                                _ = v104
                                _ = v105
                            }
                        case 687:
                            for fi106 := 0; fi106 < 16; fi106++ {
                                fmt.Println(v4, v1, v3)
                                v4 = 32767
                                _ = fi106
                            }
                        case 127:
                            switch int(v4) {
                            case -1:
                                v4 = v4
                            }
                        }
                    case "":
                        if (!false) {
                            v107 := []int{}
                            v2[(v3 & 4294967295)] = false
                            for fi108 := 0; fi108 < 3; fi108++ {
                                v109 := GenS12{F0: ("" + "hello")}
                                v4 = v4
                                _ = fi108
                                _ = v109
                            }
                            _ = v107
                        } else {
                            var v110 float32 = 1e10
                            _ = v110
                        }
                    default:
                        v1 += v1
                    }
                }
            }
        case "\n\t":
            v111 := GenS3{F0: 1, F1: v1}
            _ = v111
        case "hello":
            var v112 uint64 = 42
            _ = v112
        default:
            v113 := GenS13{F0: v1, F1: (!false)}
            _ = v113
        }
        switch (("\n\t" + "") + ((((("hello" + "a") + "foo bar") + "\n\t") + "a") + "")) {
        case "":
            var v114 string = "foo bar"
            _ = v114
        case "\n\t":
            fmt.Println(v1)
        case "foo bar":
            v115 := []int{343, -9223372036854775808, 9223372036854775807, -9223372036854775808}
            _ = v115
        }
        v3 = 0
    }
    if true {
        if (!true) {
            for fi116 := 0; fi116 < 12; fi116++ {
                fmt.Println(fi116)
                switch ("hello" + "foo bar") {
                case "foo bar":
                    for fi117 := 0; fi117 < 17; fi117++ {
                        v3 -= 846
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v3, v1)
                        default:
                            delete(v2, (v3 | v3))
                        }
                        if false {
                            v4 = v4
                        } else {
                            var v118 int8 = 1
                            var v119 bool = ((true && ("hello" < "hello")) && (-32768 <= (-(v4))))
                            if (!false) {
                                v120 := GenS14{F0: v3, F1: 0}
                                v3 = (^(v3))
                                _ = v120
                            } else {
                                v121 := []int{fi117, fi117, (fi116 + fi116)}
                                _ = v121
                            }
                            _ = v118
                            _ = v119
                        }
                        _ = fi117
                    }
                case "\n\t":
                    var v122 string = "a"
                    _ = v122
                default:
                    if (18446744073709551615 < 2) {
                        if false {
                            v123 := GenS13{F0: v1, F1: (2 > 255)}
                            _ = v123
                        } else {
                            var v124 int8 = 42
                            _ = v124
                        }
                    }
                }
                _ = fi116
            }
            if false {
                if false {
                    switch "a" {
                    case "a":
                        for fi125 := 0; fi125 < 18; fi125++ {
                            var v126 uint = 65535
                            var v127 int32 = v1
                            if false {
                                switch "foo bar" {
                                case "hello":
                                    v4 -= (v4 ^ 42)
                                default:
                                    v128 := GenS15{F0: (v126 >> (((v126 >> ((v126) & 63))) & 63)), F1: (v1 & (v1 % (((-(v127))) | 1)))}
                                    _ = v128
                                }
                                v3 -= v3
                                switch genFn0(genFn1(1, v126), 255) {
                                case -273:
                                    for fi129 := 0; fi129 < 20; fi129++ {
                                        var v130 string = "hello"
                                        if (1 > 9223372036854775807) {
                                            v131 := []int{fi129, (fi125 & (fi129 << ((v126) & 63))), (fi129 - genFn0(32767, 128))}
                                            switch (fi129 / ((439) | 1)) {
                                            case 1:
                                                if (true || false) {
                                                    var v132 int32 = (v1 | (v1 >> ((v126) & 63)))
                                                    var v133 float32 = -2.25
                                                    _ = v132
                                                    _ = v133
                                                }
                                            }
                                            var v134 uint64 = 0
                                            _ = v131
                                            _ = v134
                                        }
                                        v135 := []int{}
                                        if false {
                                            fmt.Println(fi129)
                                        }
                                        v1 -= (v1 | v1)
                                        _ = fi129
                                        _ = v130
                                        _ = v135
                                    }
                                case 32767:
                                    _ = len(v2)
                                case 42:
                                    v136 := []int{int(v3)}
                                    _ = v136
                                }
                            } else {
                                if (false && false) {
                                    v137 := GenS13{F0: 1, F1: (!true)}
                                    delete(v2, (v3 ^ (v3 - v3)))
                                    _ = v137
                                }
                                var v138 uint8 = 127
                                delete(v2, (v3 << (((v126 % (((v126 ^ (v126 - 887))) | 1))) & 63)))
                                _ = v138
                            }
                            _ = fi125
                            _ = v126
                            _ = v127
                        }
                    case "hello":
                        if false {
                            v1 += (v1 << ((479) & 63))
                            for fi139 := 0; fi139 < 4; fi139++ {
                                v1 += (v1 / (((v1 >> ((610) & 63))) | 1))
                                _ = fi139
                            }
                            var v140 bool = (2147483647 <= -2147483648)
                            _ = v140
                        } else {
                            var v141 int64 = -105
                            for fi142 := 0; fi142 < 6; fi142++ {
                                v143 := v2[(v3 % ((2) | 1))]
                                if v143 {
                                    v144 := GenS16{F0: (^(v1))}
                                    var v145 uint16 = 2
                                    for fi146 := 0; fi146 < 7; fi146++ {
                                        fmt.Println(v4, v145)
                                        var v147 bool = false
                                        v2[(v3 ^ v3)] = (0 == v3)
                                        _ = fi146
                                        _ = v147
                                    }
                                    _ = v144
                                    _ = v145
                                }
                                for fi148 := 0; fi148 < 9; fi148++ {
                                    if v143 {
                                        v149 := GenS5{F0: 255}
                                        var v150 int32 = (v1 * (^(v1)))
                                        _ = v149
                                        _ = v150
                                    } else {
                                        switch (^(fi148)) {
                                        case 127:
                                            v1 = (-(v1))
                                        case 2147483647:
                                            var v151 int8 = int8(v1)
                                            _ = v151
                                        }
                                    }
                                    _ = fi148
                                }
                                _ = fi142
                                _ = v143
                            }
                            if ((-(v141)) < v141) {
                                v152 := GenS17{F0: -128, F1: 1, F2: (int16(v141) >= v4)}
                                if ((3.14159 * (float64(0.0) / 1e10)) < -2.25) {
                                    v153 := GenS18{F0: genFn1((v141 ^ v141), 1)}
                                    var v154 int8 = 1
                                    for fi155 := 0; fi155 < 8; fi155++ {
                                        fmt.Println(v1, v141)
                                        fmt.Println(v154, v154)
                                        _ = fi155
                                    }
                                    _ = v153
                                    _ = v154
                                }
                                fmt.Println(v3)
                                _ = v152
                            }
                            _ = v141
                        }
                    case "\n\t":
                        if false {
                            var v156 bool = false
                            _ = v156
                        } else {
                            for fi157 := 0; fi157 < 2; fi157++ {
                                if (13 > 160) {
                                    v2[2147483648] = true
                                } else {
                                    switch (("" + "a") + "a") {
                                    case "\n\t":
                                        if false {
                                            v4 = 60
                                            _ = len(v2)
                                            for fi158 := 0; fi158 < 2; fi158++ {
                                                var v159 int16 = 42
                                                var v160 uint = 1
                                                var v161 uint32 = v3
                                                _ = fi158
                                                _ = v159
                                                _ = v160
                                                _ = v161
                                            }
                                        }
                                    default:
                                        v162 := []int{(fi157 * (fi157 % ((42) | 1))), (fi157 >> ((2) & 63)), fi157}
                                        _ = v162
                                    }
                                }
                                v4 = (v4 | v4)
                                _ = fi157
                            }
                            for fi163 := 0; fi163 < 9; fi163++ {
                                v1 = (-(v1))
                                v164 := v2[v3]
                                var v165 int16 = (v4 % ((1000) | 1))
                                _ = fi163
                                _ = v164
                                _ = v165
                            }
                        }
                    }
                    v4 += (^(v4))
                }
                if ((3.14159 - 0.1) != (0.1 - (float64(1e10) + 0.1))) {
                    v3 += 1
                    v166 := []int{127, 773, genFn0(genFn1(1, 670), 74)}
                    v167 := GenS19{F0: genFn1(9223372036854775807, 4294967295), F1: (v4 ^ (v4 & (v4 * v4))), F2: genFn1(-9223372036854775808, 18446744073709551615)}
                    _ = v166
                    _ = v167
                }
            }
        }
    } else {
        var v168 uint32 = 0
        var v169 bool = true
        _ = v168
        _ = v169
    }
    v1 = v1
    switch 32767 {
    case 1:
        if false {
            v4 += (v4 >> ((547) & 63))
        }
    case 866:
        v170 := GenS10{F0: 1, F1: (v3 & 106)}
        _ = v170
    case 128:
        var v171 int16 = v4
        _ = v171
    }
    v172 := GenS8{F0: 42, F1: 1}
    v173 := GenS20{F0: 1, F1: 2147483647, F2: 200}
    if (!(!(42 < -128))) {
        var v174 uint8 = 224
        switch ("hello" + "foo bar") {
        case "":
            var v175 uint32 = (^(v3))
            _ = v175
        case "\n\t":
            for fi176 := 0; fi176 < 4; fi176++ {
                var v177 uint64 = 2
                _ = fi176
                _ = v177
            }
        case "hello":
            for fi178 := 0; fi178 < 1; fi178++ {
                var v179 bool = false
                var v180 uint16 = uint16(v174)
                _ = fi178
                _ = v179
                _ = v180
            }
        default:
            fmt.Println(v3, v174, v4)
        }
        _ = v174
    }
    v1 = (v1 / ((v1) | 1))
    v181 := []int{32768, 32768, -567, 9223372036854775807}
    switch -478 {
    case 466:
        if false {
            var v182 uint32 = v3
            switch "foo bar" {
            case "\n\t":
                v183 := GenS21{F0: ("foo bar" + "hello"), F1: v1}
                _ = v183
            default:
                v182 += (v3 / (((v182 & (^(v3)))) | 1))
            }
            _ = v182
        } else {
            v184 := GenS19{F0: 42, F1: v4, F2: 0}
            v1 += -1
            _ = v184
        }
    case 65:
        v3 += (v3 << ((1) & 63))
    case 1:
        var v185 int8 = 21
        _ = v185
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v181:", len(v181), v181)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v172
    _ = v173
    _ = v181
    _ = v2
    _ = v3
    _ = v4
}
