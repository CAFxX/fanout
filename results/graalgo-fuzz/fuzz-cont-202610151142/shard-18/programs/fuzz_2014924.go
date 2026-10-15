package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 bool }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int8; F1 string; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64; F1 uint16; F2 int }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint16; F1 uint8 }
type GenS13 struct { F0 int8 }
func genFn0(p0 string, p1 int32) uint64 {
    var v4 int8 = 127
    p0 = "\n\t"
    if true {
        var v5 bool = (2 != 2)
        v6 := []int{127}
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint = 1
        var v9 float32 = 1e10
        var v10 int64 = 42
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    return 18446744073709551615
}
func genFn1() int64 {
    var v11 uint8 = 127
    var v12 float32 = 0.0
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 uint8 = 42
    var v2 float32 = 1.5
    var v3 uint16 = 310
    fmt.Println(v2)
    v13 := GenS1{F0: -128, F1: 1, F2: false}
    v14 := GenS2{F0: (v1 - (v1 << ((255) & 63))), F1: -433}
    if (true || false) {
        v15 := GenS1{F0: 127, F1: 4294967295, F2: true}
        v16 := GenS3{F0: -128, F1: "foo bar", F2: -2147483648}
        if true {
            var v17 int = -604
            switch (("foo bar" + "") + "a") {
            case "\n\t":
                for fi18 := 0; fi18 < 19; fi18++ {
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v20 := GenS2{F0: (v1 << ((18446744073709551615) & 63)), F1: 0}
                        if true {
                            switch (fi19 / ((fi19) | 1)) {
                            case 0:
                                var v21 uint16 = (v3 << ((0) & 63))
                                _ = v21
                            case -2147483648:
                                v2 = v2
                            case -1:
                                if (!(!(0.1 == (float64(1.5) - (0.0 + (0.1 * float64(-2.25))))))) {
                                    v3 += (v3 << ((784) & 63))
                                    var v22 float32 = v2
                                    _ = v22
                                } else {
                                    var v23 uint16 = v3
                                    _ = v23
                                }
                            default:
                                var v24 int8 = 42
                                fmt.Println(v2, v17)
                                switch "hello" {
                                case "hello":
                                    v25 := GenS4{F0: 42}
                                    _ = v25
                                }
                                v2 = 1e10
                                var v26 uint32 = 2147483648
                                _ = v24
                                _ = v26
                            }
                            var v27 uint64 = 932
                            _ = v27
                        }
                        v3 = (v3 - (v3 % (((v3 & v3)) | 1)))
                        _ = fi19
                        _ = v20
                    }
                    v28 := GenS4{F0: 486}
                    _ = fi18
                    _ = v28
                }
            case "a":
                for fi29 := 0; fi29 < 4; fi29++ {
                    v1 = v1
                    v2 += (v2 + v2)
                    _ = fi29
                }
            }
            _ = v17
        }
        _ = v15
        _ = v16
    } else {
        fmt.Println(v2)
        v3 -= (v3 + (v3 / ((255) | 1)))
    }
    v30 := GenS5{F0: 42}
    v1 = uint8(v3)
    v2 = v2
    var v31 bool = (1 < 650)
    var v32 int = 42
    var v33 bool = (true && (1 <= -1))
    var v34 uint64 = 0
    if false {
        fmt.Println(v31, v3, v2)
    } else {
        switch "" {
        case "hello":
            switch ("foo bar" + (("\n\t" + "hello") + "")) {
            case "a":
                v35 := []int{v32, (v32 & (v32 >> ((1) & 63)))}
                _ = v35
            case "hello":
                if false {
                    v36 := []int{}
                    switch v32 {
                    case -683:
                        switch ("hello" + ("foo bar" + (("\n\t" + "foo bar") + "hello"))) {
                        case "foo bar":
                            _ = len(v36)
                        case "hello":
                            v31 = true
                        }
                    case 32768:
                        if len(v36) >= 2 {
                            v36 = v36[1:len(v36)]
                        }
                    default:
                        var v37 int = v32
                        _ = v37
                    }
                    _ = v36
                } else {
                    if (!v33) {
                        var v38 float64 = 0.1
                        for fi39 := 0; fi39 < 11; fi39++ {
                            v40 := GenS4{F0: 734}
                            _ = fi39
                            _ = v40
                        }
                        v31 = ((205 == v1) && ((-2.25 > v38) || (40 == 42)))
                        _ = v38
                    }
                }
            case "":
                switch ((("" + (("foo bar" + "\n\t") + (("foo bar" + "hello") + "hello"))) + "hello") + "hello") {
                case "a":
                    fmt.Println(v3, v2)
                default:
                    var v41 int32 = 2147483647
                    _ = v41
                }
            default:
                v42 := GenS2{F0: (v1 * 42), F1: -762}
                _ = v42
            }
        case "foo bar":
            var v43 uint16 = v3
            _ = v43
        case "\n\t":
            if true {
                switch ("" + "\n\t") {
                case "":
                    v31 = (!(1 != 29))
                default:
                    for fi44 := 0; fi44 < 14; fi44++ {
                        var v45 uint32 = 42
                        v46 := GenS6{F0: v32, F1: (fi44 << ((18446744073709551615) & 63))}
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                }
                for fi47 := 0; fi47 < 15; fi47++ {
                    var v48 int16 = 0
                    fmt.Println(v48)
                    v32 = (fi47 / (((fi47 << ((2) & 63))) | 1))
                    _ = fi47
                    _ = v48
                }
            } else {
                switch v32 {
                case -9223372036854775808:
                    v49 := GenS5{F0: 166}
                    _ = v49
                case 0:
                    switch ("foo bar" + ("foo bar" + "a")) {
                    case "foo bar":
                        var v50 uint64 = v34
                        _ = v50
                    case "a":
                        for fi51 := 0; fi51 < 13; fi51++ {
                            for fi52 := 0; fi52 < 20; fi52++ {
                                v33 = v31
                                switch v32 {
                                case -80:
                                    switch (fi52 % ((1) | 1)) {
                                    case 127:
                                        v53 := []int{fi51, (fi51 << ((0) & 63)), fi52}
                                        _ = v53
                                    case 0:
                                        v31 = true
                                    case 42:
                                        var v54 float64 = 0.0
                                        v55 := GenS7{F0: 0}
                                        if v33 {
                                            v56 := []int{fi52, -9223372036854775808, fi51}
                                            _ = v56
                                        } else {
                                            var v57 uint64 = (v34 & (v34 & 18446744073709551615))
                                            v58 := []int{-9223372036854775808, (fi51 | (fi52 * (-(fi51))))}
                                            _ = v57
                                            _ = v58
                                        }
                                        switch ("foo bar" + "") {
                                        case "a":
                                            var v59 uint = 2
                                            _ = v59
                                        default:
                                            var v60 float32 = 1.5
                                            _ = v60
                                        }
                                        _ = v54
                                        _ = v55
                                    default:
                                        v3 = 0
                                    }
                                case -1:
                                    v1 = (v1 << ((18446744073709551615) & 63))
                                default:
                                    v61 := []int{}
                                    _ = v61
                                }
                                _ = fi52
                            }
                            _ = fi51
                        }
                    }
                }
                if v33 {
                    v62 := GenS1{F0: 42, F1: 1, F2: (1 < 42)}
                    if false {
                        v63 := GenS4{F0: 2}
                        _ = v63
                    }
                    fmt.Println(v3, v31, v32)
                    _ = v62
                } else {
                    v31 = (!v31)
                    for fi64 := 0; fi64 < 19; fi64++ {
                        switch "\n\t" {
                        case "hello":
                            if (v33 || v33) {
                                if v31 {
                                    v1 = (v1 / (((v1 | (v1 - 153))) | 1))
                                }
                                for fi65 := 0; fi65 < 5; fi65++ {
                                    switch "foo bar" {
                                    case "a":
                                        var v66 int32 = 42
                                        _ = v66
                                    case "":
                                        if (!v31) {
                                            var v67 int32 = 19
                                            v67 = v67
                                            v3 -= 0
                                            v68 := []int{fi64, v32, -2147483648}
                                            _ = v67
                                            _ = v68
                                        } else {
                                            var v69 float32 = (v2 + v2)
                                            if false {
                                                v70 := GenS8{F0: 2, F1: v3, F2: 0}
                                                var v71 int = (fi65 & fi64)
                                                v3 += v3
                                                _ = v70
                                                _ = v71
                                            } else {
                                                if (469 > (v34 ^ v34)) {
                                                    var v72 uint = 255
                                                    v3 = v3
                                                    _ = v72
                                                } else {
                                                    if v33 {
                                                        var v73 uint64 = (v34 - (v34 / (((v34 % (((v34 ^ (v34 >> ((65535) & 63)))) | 1))) | 1)))
                                                        if v33 {
                                                            var v74 uint64 = (^(v34))
                                                            v75 := GenS9{F0: int(v3)}
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        var v76 uint16 = 758
                                                        _ = v73
                                                        _ = v76
                                                    } else {
                                                        var v77 string = "\n\t"
                                                        _ = v77
                                                    }
                                                    var v78 int64 = 0
                                                    fmt.Println(v1, v31, fi64)
                                                    _ = v78
                                                }
                                            }
                                            var v79 uint64 = (v34 & (^(v34)))
                                            v80 := []int{2147483647, fi65}
                                            _ = v69
                                            _ = v79
                                            _ = v80
                                        }
                                    default:
                                        v1 = (v1 << ((1) & 63))
                                    }
                                    v81 := []int{(v32 << ((1) & 63)), v32, (fi64 / ((fi65) | 1))}
                                    v82 := GenS6{F0: 42, F1: (v32 + int(v1))}
                                    _ = fi65
                                    _ = v81
                                    _ = v82
                                }
                                v83 := []int{(^(fi64))}
                                _ = v83
                            }
                        case "foo bar":
                            v1 = (v1 / ((v1) | 1))
                        case "":
                            for fi84 := 0; fi84 < 4; fi84++ {
                                switch (("hello" + "\n\t") + "") {
                                case "":
                                    switch ("" + "foo bar") {
                                    case "a":
                                        if (v2 < v2) {
                                            var v85 uint16 = (^(v3))
                                            fmt.Println(v1, v33, v2)
                                            _ = v85
                                        } else {
                                            var v86 int8 = -1
                                            _ = v86
                                        }
                                    case "foo bar":
                                        switch (("hello" + "hello") + "hello") {
                                        case "foo bar":
                                            var v87 uint8 = (^(v1))
                                            _ = v87
                                        case "a":
                                            v88 := []int{128, fi84, fi64}
                                            _ = v88
                                        case "\n\t":
                                            v3 += uint16(v34)
                                        }
                                    case "hello":
                                        v89 := []int{fi64}
                                        _ = v89
                                    }
                                case "foo bar":
                                    var v90 int8 = 0
                                    _ = v90
                                }
                                if v33 {
                                    switch "hello" {
                                    case "foo bar":
                                        switch "hello" {
                                        case "foo bar":
                                            fmt.Println(v1)
                                        case "":
                                            v34 = (v34 * 2)
                                            v33 = false
                                        case "\n\t":
                                            v91 := []int{}
                                            _ = v91
                                        }
                                    case "\n\t":
                                        v1 = 1
                                    default:
                                        fmt.Println(fi64, v3, v32)
                                    }
                                    v32 = fi64
                                } else {
                                    v92 := []int{(v32 ^ fi84)}
                                    v32 = fi64
                                    var v93 float32 = 0.0
                                    switch 32768 {
                                    case 127:
                                        if (434 >= (^(v34))) {
                                            v2 += v93
                                        }
                                    case 879:
                                        fmt.Println(v2, v32, v93)
                                    default:
                                        var v94 float32 = 0.0
                                        _ = v94
                                    }
                                    v34 += v34
                                    var v95 int = (fi64 % (((fi64 << ((440) & 63))) | 1))
                                    _ = v92
                                    _ = v93
                                    _ = v95
                                }
                                _ = fi84
                            }
                        default:
                            if (!((755 < 127) && ((v1 % (((v1 | v1)) | 1)) > v1))) {
                                fmt.Println(v1, v33)
                            } else {
                                v32 += fi64
                                if v31 {
                                    switch 2147483647 {
                                    case 9223372036854775807:
                                        fmt.Println(v1, v31, v32)
                                    case 4294967295:
                                        switch fi64 {
                                        case 0:
                                            v96 := []int{v32}
                                            _ = v96
                                        case 2147483647:
                                            var v97 uint32 = 2147483647
                                            _ = v97
                                        }
                                    case 32767:
                                        v98 := GenS10{F0: 42}
                                        _ = v98
                                    default:
                                        v3 = (v3 + v3)
                                    }
                                    switch (fi64 + (fi64 + 1)) {
                                    case -9223372036854775808:
                                        switch -2147483648 {
                                        case -2147483648:
                                            var v99 int64 = -283
                                            _ = v99
                                        }
                                    }
                                    v100 := GenS1{F0: 127, F1: 2, F2: (v33 || v33)}
                                    _ = v100
                                } else {
                                    if v31 {
                                        v101 := GenS11{F0: v33}
                                        for fi102 := 0; fi102 < 1; fi102++ {
                                            v103 := []int{-914}
                                            if true {
                                                var v104 string = "foo bar"
                                                switch v32 {
                                                case 0:
                                                    fmt.Println(v31)
                                                case 127:
                                                    v105 := GenS9{F0: fi64}
                                                    _ = v105
                                                case -742:
                                                    if v33 {
                                                        if len(v103) >= 2 {
                                                            v103 = v103[1:len(v103)]
                                                        }
                                                    }
                                                }
                                                if ((v34 | uint64(v3)) == (v34 * genFn0(v104, -622))) {
                                                    v103 = append(v103, fi102)
                                                } else {
                                                    v106 := GenS12{F0: v3, F1: v1}
                                                    var v107 int16 = -32768
                                                    v1 += (v1 - (v1 + 23))
                                                    var v108 uint = 164
                                                    _ = v106
                                                    _ = v107
                                                    _ = v108
                                                }
                                                if (344 <= uint(v34)) {
                                                    v109 := GenS7{F0: 984}
                                                    var v110 int8 = 1
                                                    _ = v109
                                                    _ = v110
                                                }
                                                var v111 string = v104
                                                _ = v104
                                                _ = v111
                                            }
                                            var v112 int16 = -752
                                            _ = fi102
                                            _ = v103
                                            _ = v112
                                        }
                                        _ = v101
                                    }
                                    var v113 bool = (13 > -128)
                                    for fi114 := 0; fi114 < 20; fi114++ {
                                        switch (^(fi114)) {
                                        case 127:
                                            var v115 int64 = -1
                                            _ = v115
                                        case 0:
                                            var v116 uint32 = 654
                                            _ = v116
                                        case -2147483648:
                                            var v117 int8 = 0
                                            _ = v117
                                        default:
                                            v34 += (v34 * v34)
                                        }
                                        switch fi64 {
                                        case 211:
                                            v32 = (fi64 >> ((1) & 63))
                                        default:
                                            v118 := []int{int(v1), (-(v32)), fi64, fi114}
                                            _ = v118
                                        }
                                        if (!false) {
                                            if (false || (138 <= 1)) {
                                                v119 := []int{959, v32, -266, v32}
                                                fmt.Println(v2, v113, fi64)
                                                if ("hello" > "\n\t") {
                                                    fmt.Println(fi114, v2)
                                                } else {
                                                    var v120 int8 = 94
                                                    var v121 int8 = v120
                                                    _ = v120
                                                    _ = v121
                                                }
                                                _ = v119
                                            } else {
                                                var v122 int8 = -16
                                                v1 -= v1
                                                var v123 int32 = -903
                                                _ = v122
                                                _ = v123
                                            }
                                            fmt.Println(v1)
                                            if (!v113) {
                                                v124 := GenS6{F0: v32, F1: (fi114 & 211)}
                                                _ = v124
                                            } else {
                                                v2 = (float32(float64(v2)) + v2)
                                            }
                                        } else {
                                            var v125 uint = 65535
                                            v3 += 0
                                            _ = v125
                                        }
                                        v126 := GenS4{F0: 18446744073709551615}
                                        _ = fi114
                                        _ = v126
                                    }
                                    _ = v113
                                }
                                v31 = v31
                            }
                        }
                        if ((v32 >> ((772) & 63)) <= v32) {
                            if (((!(!(v31 || ((("hello" + ("a" + ("" + "hello"))) + "") <= ("" + "\n\t"))))) || v31) || (genFn0("foo bar", 42) >= v34)) {
                                if true {
                                    fmt.Println(v3, v1)
                                    fmt.Println(v32, v2, v3)
                                    if (v1 == (v1 / ((127) | 1))) {
                                        var v127 int32 = -867
                                        v34 += (v34 >> ((73) & 63))
                                        var v128 int64 = 42
                                        _ = v127
                                        _ = v128
                                    }
                                } else {
                                    v129 := GenS13{F0: 0}
                                    _ = v129
                                }
                                if v33 {
                                    v130 := GenS4{F0: 0}
                                    for fi131 := 0; fi131 < 19; fi131++ {
                                        v32 = (fi64 * (-(fi64)))
                                        v3 = (v3 % ((v3) | 1))
                                        _ = fi131
                                    }
                                    for fi132 := 0; fi132 < 6; fi132++ {
                                        switch (v32 - (-(v32))) {
                                        case 0:
                                            switch (v32 << ((65535) & 63)) {
                                            case -599:
                                                var v133 uint16 = v3
                                                _ = v133
                                            default:
                                                fmt.Println(v3, fi132)
                                            }
                                        case 128:
                                            v134 := GenS7{F0: 42}
                                            _ = v134
                                        }
                                        switch v32 {
                                        case 32767:
                                            v135 := []int{}
                                            _ = v135
                                        case -2147483648:
                                            switch "" {
                                            case "foo bar":
                                                var v136 int64 = 320
                                                _ = v136
                                            case "a":
                                                if ((-896 < (v32 * v32)) || v31) {
                                                    switch "" {
                                                    case "":
                                                        if v31 {
                                                            v3 = v3
                                                        }
                                                    case "foo bar":
                                                        var v137 uint8 = v1
                                                        _ = v137
                                                    case "a":
                                                        var v138 uint64 = (v34 - 18446744073709551615)
                                                        _ = v138
                                                    default:
                                                        fmt.Println(fi64, v32)
                                                    }
                                                } else {
                                                    v139 := []int{}
                                                    var v140 int64 = genFn1()
                                                    _ = v139
                                                    _ = v140
                                                }
                                            default:
                                                v141 := []int{2147483647, fi64}
                                                _ = v141
                                            }
                                        }
                                        v33 = (-1 > -573)
                                        _ = fi132
                                    }
                                    _ = v130
                                }
                                v32 = v32
                            }
                            var v142 float64 = 1e10
                            _ = v142
                        }
                        _ = fi64
                    }
                }
                var v143 int32 = 42
                _ = v143
            }
        default:
            var v144 string = "hello"
            _ = v144
        }
    }
    v145 := GenS11{F0: v31}
    var v146 int64 = -415
    var v147 int8 = -128
    v32 += v32
    var v148 int = v32
    if ((false && (v33 || (453 <= 2147483647))) || v31) {
        if (uint(v34) <= 42) {
            v146 = 0
        }
    } else {
        switch v148 {
        case 42:
            var v149 uint8 = (v1 + v1)
            _ = v149
        case 127:
            for fi150 := 0; fi150 < 10; fi150++ {
                var v151 int = fi150
                switch "a" {
                case "":
                    if v33 {
                        var v152 uint64 = 0
                        _ = v152
                    } else {
                        v153 := []int{(v148 ^ v151), (fi150 - fi150), v151}
                        var v154 int64 = 42
                        _ = v153
                        _ = v154
                    }
                default:
                    fmt.Println(v3)
                }
                _ = fi150
                _ = v151
            }
        case -600:
            v155 := GenS6{F0: (v148 / ((v148) | 1)), F1: int(v147)}
            _ = v155
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v146:", v146)
    fmt.Println("v147:", v147)
    fmt.Println("v148:", v148)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v13
    _ = v14
    _ = v145
    _ = v146
    _ = v147
    _ = v148
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
