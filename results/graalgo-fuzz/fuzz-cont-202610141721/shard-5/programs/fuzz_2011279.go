package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 int; F2 uint32 }
type GenS7 struct { F0 int64; F1 uint; F2 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16; F1 int; F2 bool }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS15 struct { F0 int32; F1 uint8 }
type GenS16 struct { F0 int16; F1 int16; F2 string }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint; F1 int64 }
type GenS19 struct { F0 uint8; F1 int }
type GenS20 struct { F0 int16; F1 uint32 }
type GenS21 struct { F0 int8; F1 uint }
func genFn0(p0 int32, p1 bool, p2 int16) uint64 {
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 float64 = float64(((3.14159 * 3.14159) / (3.14159 / (0.1 + 3.14159))))
        var v7 float32 = 1.5
        var v8 uint8 = 255
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    fmt.Println(p2)
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 string = "foo bar"
        var v11 uint16 = 255
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "foo bar" {
    case "\n\t":
        v12 := GenS1{F0: ("\n\t" + "a")}
        _ = v12
    default:
        if p1 {
            if false {
                var v13 float32 = (0.0 * 0.0)
                var v14 int16 = p2
                var v15 int32 = (-(p0))
                _ = v13
                _ = v14
                _ = v15
            }
        }
    }
    var v16 string = ("" + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 4294967295
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -252
    var v3 int8 = 42
    var v4 uint64 = 0
    var v17 int = 2147483647
    switch ("\n\t" + "\n\t") {
    case "a":
        var v18 int64 = 0
        _ = v18
    default:
        v3 = -72
    }
    switch v17 {
    case 0:
        if true {
            for fi19 := 0; fi19 < 3; fi19++ {
                v4 = (^(v4))
                _ = fi19
            }
        } else {
            var v20 int16 = 1
            v21 := GenS1{F0: ("foo bar" + "\n\t")}
            _ = v20
            _ = v21
        }
    case -970:
        fmt.Println(v4, v4)
    case 32767:
        fmt.Println(v17, v2, v3)
    default:
        var v22 float32 = ((0.0 / (((((0.0 + (((1e10 / 1e10) / 1.0) + 1.5)) + (1e10 + ((float32(-2.25) - 0.0) / float32(1.5)))) + (1.5 * 1.5)) + 1.5) * ((0.1 * ((1e10 - 0.1) - float32(-2.25))) - 1.5))) * 1.5)
        _ = v22
    }
    v23 := GenS1{F0: ("hello" + "a")}
    if (!true) {
        v24 := GenS1{F0: "foo bar"}
        switch "a" {
        case "":
            v3 += -71
        case "a":
            if false {
                v17 -= v17
                v3 += v3
            } else {
                switch (v17 + (v17 / (((v17 % ((v17) | 1))) | 1))) {
                case 0:
                    if false {
                        if (!((32767 < 405) || (v4 == v4))) {
                            for fi25 := 0; fi25 < 7; fi25++ {
                                for fi26 := 0; fi26 < 13; fi26++ {
                                    v4 += 0
                                    var v27 uint8 = 127
                                    _ = fi26
                                    _ = v27
                                }
                                _ = fi25
                            }
                            v17 = -375
                        } else {
                            v28 := []int{(v17 & (v17 + 127)), v17, v17, (v17 & 189)}
                            _ = v28
                        }
                        if true {
                            if (!true) {
                                var v29 string = "foo bar"
                                fmt.Println(v2, v17)
                                _ = v29
                            }
                            v30 := GenS1{F0: "foo bar"}
                            v4 = v4
                            _ = v30
                        } else {
                            var v31 uint = 0
                            var v32 int = (^(v17))
                            v33 := GenS1{F0: ("a" + "\n\t")}
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                    }
                default:
                    var v34 uint64 = v4
                    _ = v34
                }
            }
        }
        fmt.Println(v2, v4, v17)
        _ = v24
    } else {
        var v35 bool = true
        var v36 uint16 = 32767
        _ = v35
        _ = v36
    }
    _ = len(v1)
    if true {
        for fi37 := 0; fi37 < 9; fi37++ {
            v38 := GenS1{F0: (("" + ("foo bar" + "foo bar")) + ("a" + "\n\t"))}
            var v39 uint64 = genFn0(v2, true, 1)
            _ = fi37
            _ = v38
            _ = v39
        }
    } else {
        var v40 int = v17
        _ = v40
    }
    v4 = genFn0(219, false, 32767)
    var v41 int8 = v3
    fmt.Println(v17, v41, v3)
    switch ("hello" + "hello") {
    case "foo bar":
        v42 := GenS2{F0: true}
        _ = v42
    case "a":
        if ((0 == 1) || (18446744073709551615 == 0)) {
            if (10 > 0) {
                switch "hello" {
                case "foo bar":
                    fmt.Println(v2)
                }
                switch (v17 / ((v17) | 1)) {
                case 9223372036854775807:
                    for fi43 := 0; fi43 < 18; fi43++ {
                        var v44 float64 = 3.14159
                        v45 := GenS3{F0: 1, F1: 42}
                        for fi46 := 0; fi46 < 18; fi46++ {
                            v47 := GenS2{F0: true}
                            switch "hello" {
                            case "":
                                v2 -= (v2 % ((v2) | 1))
                            case "hello":
                                var v48 uint64 = 631
                                _ = v48
                            case "a":
                                v17 = (v17 & 2147483647)
                            }
                            v49 := []int{}
                            _ = fi46
                            _ = v47
                            _ = v49
                        }
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                default:
                    var v50 uint = 42
                    _ = v50
                }
                switch "\n\t" {
                case "hello":
                    var v51 uint16 = 65535
                    _ = v51
                case "a":
                    var v52 float32 = (3.14159 * 3.14159)
                    _ = v52
                }
            }
        } else {
            for fi53 := 0; fi53 < 18; fi53++ {
                if false {
                    for fi54 := 0; fi54 < 16; fi54++ {
                        v55 := GenS2{F0: true}
                        v4 += v4
                        v56 := GenS4{F0: (v2 + (-(v2))), F1: (!true)}
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                    v57 := GenS5{F0: 42}
                    _ = v57
                } else {
                    for fi58 := 0; fi58 < 1; fi58++ {
                        v59 := []int{v17}
                        var v60 uint16 = 65535
                        if false {
                            if ((fi58 < (fi53 << ((2) & 63))) && true) {
                                if true {
                                    fmt.Println(v60, v17, v17)
                                    switch "hello" {
                                    case "a":
                                        var v61 uint = 255
                                        _ = v61
                                    }
                                    fmt.Println(v41, v3, v17)
                                }
                            }
                            if false {
                                var v62 uint = 42
                                v59 = append(v59, v17)
                                if false {
                                    v3 = (^(v41))
                                    if (0 <= 1) {
                                        var v63 string = ""
                                        switch v17 {
                                        case 4294967295:
                                            if false {
                                                if (1000 != -1) {
                                                    var v64 bool = false
                                                    v3 -= v3
                                                    v4 -= uint64(v62)
                                                    _ = v64
                                                } else {
                                                    v65 := GenS6{F0: 0, F1: fi53, F2: 2147483647}
                                                    if false {
                                                        v66 := GenS1{F0: (((("a" + "a") + v63) + (v63 + (v63 + "a"))) + "hello")}
                                                        _ = v66
                                                    } else {
                                                        v41 += v41
                                                        var v67 uint = v62
                                                        _ = v67
                                                    }
                                                    if (604 != 1) {
                                                        var v68 uint8 = 17
                                                        _ = v68
                                                    } else {
                                                        v69 := GenS3{F0: uint(v60), F1: -1}
                                                        _ = v69
                                                    }
                                                    _ = v65
                                                }
                                                if ((!(v4 <= genFn0((v2 + v2), true, 1000))) && ((v2 >> ((280) & 63)) == (v2 * (v2 * v2)))) {
                                                    if false {
                                                        var v70 uint64 = v4
                                                        _ = v70
                                                    }
                                                    if true {
                                                        var v71 uint32 = 42
                                                        var v72 uint64 = (v4 / ((v4) | 1))
                                                        var v73 uint32 = v71
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                    } else {
                                                        if (789 == 2147483647) {
                                                            v74 := GenS7{F0: 9223372036854775807, F1: v62, F2: v17}
                                                            fmt.Println(v63, v4, v60)
                                                            v75 := GenS3{F0: v62, F1: 42}
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        var v76 bool = false
                                                        v77 := GenS8{F0: v60}
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                } else {
                                                    fmt.Println(v2, v4)
                                                }
                                                v78 := GenS4{F0: v2, F1: (!((!false) && (0.0 <= ((1e10 * (1e10 + (float64(1.5) / (1e10 - float64(float32(0.1)))))) + ((((0.0 / 0.1) / -2.25) * 0.1) * (3.14159 - -2.25))))))}
                                                var v79 int8 = v41
                                                _ = v78
                                                _ = v79
                                            }
                                        default:
                                            fmt.Println(v41, v2)
                                        }
                                        var v80 uint = (v62 | 0)
                                        _ = v63
                                        _ = v80
                                    } else {
                                        if false {
                                            if false {
                                                if true {
                                                    v81 := GenS9{F0: false}
                                                    _ = v81
                                                }
                                                var v82 int16 = -1
                                                if (false && (!true)) {
                                                    v83 := GenS9{F0: (!true)}
                                                    if ((v2 % ((v2) | 1)) >= (v2 + 42)) {
                                                        fmt.Println(fi58, v3)
                                                    } else {
                                                        if false {
                                                            switch int(v82) {
                                                            case 127:
                                                                _ = len(v59)
                                                            case 1:
                                                                switch (v17 % ((fi58) | 1)) {
                                                                case 1:
                                                                    fmt.Println(v2, v82, v17)
                                                                case -1:
                                                                    v3 -= v41
                                                                case 9223372036854775807:
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        fmt.Println(fi53, fi53)
                                                                        v17 += v17
                                                                    }
                                                                }
                                                            default:
                                                                v84 := GenS10{F0: v82, F1: (v17 % ((fi53) | 1)), F2: ("foo bar" == ("\n\t" + "foo bar"))}
                                                                _ = v84
                                                            }
                                                            var v85 int = v17
                                                            v1[739] = false
                                                            _ = v85
                                                        }
                                                        v3 = 35
                                                        v86 := GenS7{F0: 42, F1: v62, F2: (fi53 << ((1) & 63))}
                                                        _ = v86
                                                    }
                                                    _ = v83
                                                } else {
                                                    v4 += 94
                                                    v82 -= (v82 & 782)
                                                    var v87 uint8 = 1
                                                    if (true || ("a" < ("" + ("a" + ("" + ("\n\t" + "hello")))))) {
                                                        var v88 uint = 42
                                                        _ = v88
                                                    }
                                                    _ = v87
                                                }
                                                _ = v82
                                            }
                                            var v89 string = "\n\t"
                                            var v90 uint = v62
                                            _ = v89
                                            _ = v90
                                        } else {
                                            var v91 uint = v62
                                            _ = v91
                                        }
                                        if (v2 == (v2 + (v2 ^ v2))) {
                                            if true {
                                                v92 := GenS3{F0: v62, F1: -171}
                                                var v93 uint16 = (^(v60))
                                                if (2147483647 != 4294967295) {
                                                    v93 -= v60
                                                } else {
                                                    v94 := v1[fi53]
                                                    _ = v94
                                                }
                                                _ = v92
                                                _ = v93
                                            }
                                            v95 := GenS9{F0: true}
                                            _ = v95
                                        } else {
                                            var v96 float64 = 1.5
                                            v97 := v1[(v17 % (((fi53 / (((fi53 & (fi58 * (fi58 & int(v62))))) | 1))) | 1))]
                                            _ = v96
                                            _ = v97
                                        }
                                    }
                                    var v98 float64 = 1e10
                                    var v99 uint32 = uint32(v60)
                                    v100 := GenS6{F0: (v99 >> ((v62) & 63)), F1: fi53, F2: (v99 + (v99 % ((v99) | 1)))}
                                    _ = v100
                                    _ = v98
                                    _ = v99
                                } else {
                                    if (248 != 27) {
                                        var v101 uint8 = 108
                                        v102 := GenS11{F0: v2}
                                        v103 := v1[fi53]
                                        _ = v101
                                        _ = v102
                                        _ = v103
                                    }
                                }
                                _ = v62
                            }
                        } else {
                            v104 := GenS1{F0: "\n\t"}
                            _ = v104
                        }
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                    v41 += v41
                }
                v17 -= (v17 | -402)
                _ = fi53
            }
            if true {
                if false {
                    var v105 uint64 = v4
                    _ = v105
                } else {
                    var v106 int8 = -128
                    v4 += 0
                    _ = v106
                }
                v107 := GenS12{F0: 625}
                for fi108 := 0; fi108 < 17; fi108++ {
                    v109 := GenS13{F0: 1}
                    if false {
                        switch fi108 {
                        case 128:
                            v2 = (v2 * (v2 / ((-1) | 1)))
                        default:
                            switch ("foo bar" + "hello") {
                            case "":
                                for fi110 := 0; fi110 < 20; fi110++ {
                                    v111 := []int{v17, 42, (fi108 % ((fi108) | 1))}
                                    if len(v111) > 0 {
                                        v111[0] = (fi108 | (fi110 >> ((255) & 63)))
                                    }
                                    _ = len(v1)
                                    _ = fi110
                                    _ = v111
                                }
                            case "hello":
                                var v112 uint64 = v4
                                _ = v112
                            case "\n\t":
                                if ((v2 * v2) <= 270) {
                                    v3 += (v3 << ((2) & 63))
                                }
                            }
                        }
                    } else {
                        if false {
                            if true {
                                v113 := []int{}
                                var v114 int16 = 32767
                                _ = v113
                                _ = v114
                            } else {
                                if true {
                                    switch 42 {
                                    case 796:
                                        v3 -= (v41 / ((v41) | 1))
                                    case 32767:
                                        fmt.Println(v2, v3, v3)
                                    case 4294967295:
                                        for fi115 := 0; fi115 < 9; fi115++ {
                                            var v116 string = "hello"
                                            if true {
                                                fmt.Println(fi108, fi108)
                                            } else {
                                                v117 := GenS14{F0: 0, F1: 1000, F2: 526}
                                                _ = v117
                                            }
                                            _ = fi115
                                            _ = v116
                                        }
                                    default:
                                        if (!false) {
                                            for fi118 := 0; fi118 < 14; fi118++ {
                                                var v119 string = "hello"
                                                fmt.Println(v41, v17)
                                                _ = fi118
                                                _ = v119
                                            }
                                        }
                                    }
                                    switch fi108 {
                                    case 9223372036854775807:
                                        switch -371 {
                                        case -611:
                                            for fi120 := 0; fi120 < 1; fi120++ {
                                                v121 := GenS11{F0: v2}
                                                v2 -= 312
                                                v122 := v1[449]
                                                _ = fi120
                                                _ = v121
                                                _ = v122
                                            }
                                        case 32768:
                                            var v123 int = v17
                                            _ = v123
                                        }
                                    case 128:
                                        v17 += fi108
                                    case -60:
                                        v41 -= v3
                                    default:
                                        for fi124 := 0; fi124 < 4; fi124++ {
                                            v125 := GenS15{F0: v2, F1: 196}
                                            var v126 uint16 = 32767
                                            v1[32767] = (677 == 2147483648)
                                            fmt.Println(v126, v3)
                                            var v127 uint8 = 2
                                            fmt.Println(v127, v17)
                                            v41 = v41
                                            v126 += (v126 + (^(v126)))
                                            _ = fi124
                                            _ = v125
                                            _ = v126
                                            _ = v127
                                        }
                                    }
                                } else {
                                    v128 := GenS16{F0: -35, F1: 32767, F2: "a"}
                                    var v129 uint64 = (v4 >> ((18446744073709551615) & 63))
                                    if true {
                                        fmt.Println(fi108, v17)
                                    }
                                    _ = v128
                                    _ = v129
                                }
                                if ((float64((1e10 * 0.1)) - ((1.5 - 0.0) + (((1.5 + 1.5) + (1e10 + float64(3.14159))) * 0.0))) <= (float64(0.0) - 1.5)) {
                                    v2 = v2
                                    var v130 int16 = 526
                                    _ = v130
                                } else {
                                    fmt.Println(v17)
                                    v41 -= 127
                                    if true {
                                        v131 := GenS17{F0: -652}
                                        _ = v131
                                    }
                                }
                                switch v17 {
                                case -360:
                                    var v132 float32 = 3.14159
                                    _ = v132
                                case 32768:
                                    if false {
                                        fmt.Println(v2, v3, v2)
                                        v133 := GenS10{F0: 741, F1: (fi108 + (v17 << ((65535) & 63))), F2: true}
                                        var v134 uint = 259
                                        _ = v133
                                        _ = v134
                                    }
                                }
                            }
                            v2 = v2
                        } else {
                            fmt.Println(fi108, v3)
                            var v135 int32 = 2147483647
                            _ = v135
                        }
                        v136 := []int{128, (fi108 | fi108), fi108}
                        var v137 uint16 = 42
                        _ = v136
                        _ = v137
                    }
                    _ = fi108
                    _ = v109
                }
                _ = v107
            } else {
                v41 = v3
                for fi138 := 0; fi138 < 14; fi138++ {
                    switch 4294967295 {
                    case -2147483648:
                        fmt.Println(v17)
                    }
                    _ = fi138
                }
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v4)
                case "a":
                    var v139 int16 = -1
                    _ = v139
                case "hello":
                    var v140 uint16 = 2
                    _ = v140
                default:
                    var v141 uint32 = 2147483647
                    _ = v141
                }
            }
        }
    }
    if ((v41 * 100) > 100) {
        for fi142 := 0; fi142 < 5; fi142++ {
            fmt.Println(fi142)
            v143 := GenS14{F0: 155, F1: -1, F2: 214}
            _ = fi142
            _ = v143
        }
        for fi144 := 0; fi144 < 5; fi144++ {
            for fi145 := 0; fi145 < 19; fi145++ {
                if false {
                    if ((1 > int32(v41)) || true) {
                        fmt.Println(v4, v4, v3)
                        if (true && (true && (-1 < 1))) {
                            v2 = v2
                            var v146 uint = 255
                            _ = v146
                        }
                    } else {
                        v147 := GenS6{F0: 2147483648, F1: (fi144 - 42), F2: 36}
                        _ = v147
                    }
                }
                _ = fi145
            }
            var v148 int16 = 32767
            _ = fi144
            _ = v148
        }
    }
    v1[v17] = (false || false)
    for fi149 := 0; fi149 < 17; fi149++ {
        v150 := []int{128}
        v17 = v17
        _ = fi149
        _ = v150
    }
    if false {
        v151 := GenS17{F0: v17}
        _ = v151
    }
    switch (v17 + (v17 / (((v17 >> ((470) & 63))) | 1))) {
    case -774:
        for fi152 := 0; fi152 < 9; fi152++ {
            if (!true) {
                if (((233 <= 469) && false) || (true || (65535 <= 0))) {
                    for fi153 := 0; fi153 < 17; fi153++ {
                        v154 := GenS15{F0: (v2 | v2), F1: 255}
                        _ = fi153
                        _ = v154
                    }
                    v155 := GenS18{F0: 0, F1: -458}
                    _ = v155
                }
                for fi156 := 0; fi156 < 2; fi156++ {
                    var v157 uint64 = v4
                    if ((-32768 < 32767) || (!(42 == 0))) {
                        v158 := GenS6{F0: 298, F1: 1, F2: 1}
                        _ = v158
                    }
                    v159 := GenS18{F0: 583, F1: 42}
                    v160 := GenS19{F0: 127, F1: fi152}
                    _ = fi156
                    _ = v157
                    _ = v159
                    _ = v160
                }
                var v161 float32 = (1.5 * 3.14159)
                _ = v161
            } else {
                if true {
                    if false {
                        v162 := GenS20{F0: 32767, F1: 2}
                        _ = v162
                    }
                    var v163 string = "a"
                    _ = v163
                } else {
                    var v164 int16 = 773
                    v165 := []int{32767, fi152, v17}
                    _ = v164
                    _ = v165
                }
            }
            v41 = (v41 | (v3 << ((42) & 63)))
            var v166 int8 = (v3 << ((1) & 63))
            _ = fi152
            _ = v166
        }
    }
    v167 := GenS21{F0: (v41 ^ (v3 - (v41 % ((1) | 1)))), F1: 1}
    v168 := v1[v17]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v168:", v168)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v167
    _ = v168
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v41
}
