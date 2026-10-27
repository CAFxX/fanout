package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint32) string {
    v5 := []int{1, 32767, 32767}
    p0 += 2147483648
    if true {
        var v6 int16 = 42
        _ = v6
    } else {
        var v7 uint8 = uint8(p0)
        _ = v7
    }
    v8 := GenS1{F0: 4294967295, F1: -1}
    if (2 > 42) {
        _ = len(v5)
    } else {
        var v9 string = "hello"
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v8
    return "a"
}
func genFn1(p0 uint32, p1 int) uint64 {
    p0 = p0
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{-205, (fi10 + p1), fi10}
        _ = fi10
        _ = v11
    }
    fmt.Println(p1)
    v12 := GenS2{F0: (-(p1)), F1: (!(!true)), F2: p0}
    if ((false && true) || (true || (128 == 8))) {
        for fi13 := 0; fi13 < 11; fi13++ {
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 int64 = -9223372036854775808
                var v16 float64 = (0.0 * ((1.5 / 1e10) + float64(1.5)))
                _ = fi14
                _ = v15
                _ = v16
            }
            v17 := []int{}
            _ = fi13
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = v12
    return 0
}

func main() {
    v1 := []uint32{687, 2, 0}
    var v2 int8 = 0
    v3 := make(map[uint32]uint32)
    v4 := []uint32{2147483648}
    v3[4294967295] = 65535
    if true {
        if false {
            if (65535 >= 2147483647) {
                v2 += 42
                switch 9223372036854775807 {
                case 128:
                    for fi18 := 0; fi18 < 12; fi18++ {
                        var v19 int16 = -331
                        fmt.Println(v2)
                        for fi20 := 0; fi20 < 4; fi20++ {
                            fmt.Println(v19, v19)
                            v19 = (v19 + (v19 ^ (v19 * (^(v19)))))
                            _ = fi20
                        }
                        _ = fi18
                        _ = v19
                    }
                case -9223372036854775808:
                    var v21 uint32 = 2147483648
                    _ = v21
                default:
                    for fi22 := 0; fi22 < 20; fi22++ {
                        fmt.Println(fi22, fi22)
                        for fi23 := 0; fi23 < 5; fi23++ {
                            switch (fi22 - fi22) {
                            case 200:
                                v2 = v2
                            case -2147483648:
                                fmt.Println(fi23, fi23)
                            default:
                                v2 = int8(fi22)
                            }
                            v2 -= (v2 * (v2 | (v2 / (((v2 << ((384) & 63))) | 1))))
                            _ = fi23
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi22
                    }
                }
                var v24 int = -576
                _ = v24
            }
            if true {
                switch -424 {
                case -1:
                    v2 = (-(v2))
                }
                switch 17 {
                case -357:
                    fmt.Println(v2)
                default:
                    for fi25 := 0; fi25 < 20; fi25++ {
                        var v26 int64 = 883
                        if true {
                            for fi27 := 0; fi27 < 10; fi27++ {
                                var v28 uint8 = 1
                                switch (fi25 - fi27) {
                                case 0:
                                    v2 += v2
                                case -508:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case -255:
                                    var v29 int8 = v2
                                    _ = v29
                                default:
                                    var v30 int8 = v2
                                    _ = v30
                                }
                                _ = fi27
                                _ = v28
                            }
                        } else {
                            delete(v3, 0)
                            var v31 int32 = 1
                            _ = v31
                        }
                        _ = fi25
                        _ = v26
                    }
                }
            }
        }
        v32 := GenS1{F0: 0, F1: -1}
        switch "hello" {
        case "a":
            if (709 <= 1) {
                var v33 int16 = -1
                var v34 bool = (2 < 255)
                v35 := v3[2]
                _ = v33
                _ = v34
                _ = v35
            } else {
                fmt.Println(v2)
                if (!false) {
                    for fi36 := 0; fi36 < 19; fi36++ {
                        if ((fi36 & fi36) < (^(fi36))) {
                            if false {
                                var v37 uint = 0
                                v38 := GenS3{F0: 1}
                                for fi39 := 0; fi39 < 9; fi39++ {
                                    v40 := GenS4{F0: (v2 * v2), F1: 4294967295, F2: v2}
                                    var v41 float32 = (3.14159 * 0.0)
                                    fmt.Println(v2)
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                }
                                _ = v37
                                _ = v38
                            }
                            v2 = v2
                            fmt.Println(v2)
                        } else {
                            switch fi36 {
                            case -379:
                                if (42 != 2147483647) {
                                    if len(v1) > 0 {
                                        v1[0] = 2147483648
                                    }
                                } else {
                                    for fi42 := 0; fi42 < 8; fi42++ {
                                        switch fi36 {
                                        case 0:
                                            v43 := GenS3{F0: 2147483648}
                                            _ = v43
                                        case 279:
                                            switch (fi42 / (((fi36 - fi36)) | 1)) {
                                            case 128:
                                                var v44 float64 = 1e10
                                                _ = v44
                                            case -2147483648:
                                                v2 -= v2
                                            case -232:
                                                var v45 string = ""
                                                _ = v45
                                            }
                                        case -702:
                                            v2 = (v2 << ((4294967295) & 63))
                                            v46 := GenS1{F0: 255, F1: 83}
                                            if ((-(v2)) > -1) {
                                                if (v2 > v2) {
                                                    if false {
                                                        v2 += v2
                                                        _ = len(v3)
                                                        switch (fi42 ^ (fi36 + 127)) {
                                                        case 2147483647:
                                                            var v47 float64 = 1.5
                                                            _ = v47
                                                        case 32768:
                                                            v48 := GenS5{F0: 788, F1: 128}
                                                            _ = v48
                                                        case 857:
                                                            v2 = (^(v2))
                                                        }
                                                    }
                                                    v2 += v2
                                                }
                                                switch fi42 {
                                                case 42:
                                                    v2 += 100
                                                default:
                                                    if (!(("" + "") != ("\n\t" + "a"))) {
                                                        v49 := GenS6{F0: 720}
                                                        _ = v49
                                                    }
                                                }
                                            }
                                            if true {
                                                v3[485] = 2147483648
                                                switch "a" {
                                                case "a":
                                                    fmt.Println(fi42, fi42, fi42)
                                                default:
                                                    v50 := v3[65535]
                                                    _ = v50
                                                }
                                                fmt.Println(fi42, fi42, fi36)
                                            } else {
                                                v51 := GenS5{F0: 1, F1: 1}
                                                _ = v51
                                            }
                                            _ = v46
                                        }
                                        v52 := GenS6{F0: 1}
                                        if (false || ((v2 / (((v2 % (((-(v2))) | 1))) | 1)) >= v2)) {
                                            var v53 int16 = -32768
                                            if true {
                                                v53 -= (v53 ^ (-(v53)))
                                                _ = len(v3)
                                                switch fi36 {
                                                case 42:
                                                    v54 := v3[2147483647]
                                                    _ = v54
                                                case -2147483648:
                                                    if false {
                                                        fmt.Println(v2, v2)
                                                        switch (fi36 % ((32768) | 1)) {
                                                        case -852:
                                                            fmt.Println(v2, v53)
                                                        }
                                                        if false {
                                                            v53 -= v53
                                                            var v55 uint = 1
                                                            v56 := GenS5{F0: 18446744073709551615, F1: 127}
                                                            _ = v55
                                                            _ = v56
                                                        } else {
                                                            fmt.Println(v2)
                                                            if (true || (!(!false))) {
                                                                var v57 uint64 = 18446744073709551615
                                                                var v58 uint16 = 0
                                                                var v59 bool = (!(!false))
                                                                _ = v57
                                                                _ = v58
                                                                _ = v59
                                                            } else {
                                                                var v60 float64 = 0.1
                                                                fmt.Println(fi36, v53, fi42)
                                                                fmt.Println(fi42, v2, v53)
                                                                _ = v60
                                                            }
                                                            v53 += v53
                                                        }
                                                        v53 -= (v53 ^ v53)
                                                        v61 := GenS3{F0: 4294967295}
                                                        _ = v61
                                                    }
                                                case 2147483647:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                default:
                                                    v2 = v2
                                                }
                                            } else {
                                                if ((v53 / ((v53) | 1)) == v53) {
                                                    v62 := GenS7{F0: ("\n\t" + "")}
                                                    if len(v1) > 0 {
                                                        v1[0] = 65535
                                                    }
                                                    _ = v62
                                                }
                                                fmt.Println(fi42, v53)
                                                var v63 int64 = -9223372036854775808
                                                _ = v63
                                            }
                                            _ = v53
                                        } else {
                                            var v64 string = "hello"
                                            v2 = (v2 / ((-128) | 1))
                                            var v65 uint64 = 2
                                            _ = v64
                                            _ = v65
                                        }
                                        var v66 int64 = -1
                                        var v67 uint8 = 207
                                        _ = fi42
                                        _ = v52
                                        _ = v66
                                        _ = v67
                                    }
                                    var v68 int64 = 1
                                    v68 += (v68 << ((1) & 63))
                                    _ = v68
                                }
                            default:
                                var v69 int32 = -249
                                _ = v69
                            }
                        }
                        if true {
                            v70 := GenS7{F0: ("\n\t" + "\n\t")}
                            v2 = v2
                            var v71 string = "foo bar"
                            _ = v70
                            _ = v71
                        }
                        _ = fi36
                    }
                    if (true || true) {
                        switch -892 {
                        case 2147483647:
                            v2 = -1
                        }
                        v2 -= v2
                    }
                }
                for fi72 := 0; fi72 < 4; fi72++ {
                    if false {
                        v73 := GenS5{F0: 42, F1: 127}
                        for fi74 := 0; fi74 < 4; fi74++ {
                            v2 = 127
                            if true {
                                var v75 uint32 = 117
                                var v76 bool = true
                                _ = v75
                                _ = v76
                            } else {
                                var v77 uint8 = 1
                                _ = len(v3)
                                var v78 uint64 = 0
                                _ = v77
                                _ = v78
                            }
                            v2 += v2
                            v2 -= v2
                            v79 := GenS8{F0: "foo bar"}
                            v80 := GenS5{F0: 1, F1: 1}
                            _ = fi74
                            _ = v79
                            _ = v80
                        }
                        _ = v73
                    }
                    switch fi72 {
                    case 4294967295:
                        if (true || (!true)) {
                            if true {
                                v81 := GenS1{F0: 447, F1: 42}
                                _ = v81
                            } else {
                                var v82 int64 = int64(v2)
                                var v83 int = (fi72 ^ (fi72 | (fi72 >> ((18446744073709551615) & 63))))
                                _ = v82
                                _ = v83
                            }
                            v84 := GenS4{F0: v2, F1: 476, F2: (v2 % ((v2) | 1))}
                            _ = v84
                        } else {
                            if true {
                                switch 986 {
                                case 42:
                                    v2 += v2
                                }
                                _ = len(v3)
                            } else {
                                if (255 < 127) {
                                    switch int(v2) {
                                    case 2147483647:
                                        var v85 uint32 = 1
                                        _ = v85
                                    default:
                                        v86 := GenS5{F0: 275, F1: 84}
                                        _ = v86
                                    }
                                    fmt.Println(v2, fi72)
                                    if false {
                                        fmt.Println(fi72, fi72)
                                        if false {
                                            var v87 uint16 = 421
                                            fmt.Println(v2)
                                            fmt.Println(fi72, v2, fi72)
                                            _ = v87
                                        } else {
                                            v2 -= v2
                                            var v88 int8 = v2
                                            _ = v88
                                        }
                                    } else {
                                        switch "" {
                                        case "\n\t":
                                            if true {
                                                v2 = v2
                                                v4 = append(v4, 42)
                                            }
                                        case "hello":
                                            var v89 float32 = float32(0.0)
                                            _ = v89
                                        default:
                                            fmt.Println(v2, v2)
                                        }
                                    }
                                }
                                for fi90 := 0; fi90 < 11; fi90++ {
                                    var v91 float64 = 3.14159
                                    var v92 uint32 = 408
                                    delete(v3, (v92 % ((v92) | 1)))
                                    var v93 uint8 = 199
                                    v91 -= 3.14159
                                    _ = fi90
                                    _ = v91
                                    _ = v92
                                    _ = v93
                                }
                            }
                        }
                    case -717:
                        var v94 bool = true
                        _ = v94
                    case -556:
                        var v95 int = (fi72 << ((2) & 63))
                        _ = v95
                    }
                    _ = fi72
                }
            }
        case "":
            for fi96 := 0; fi96 < 4; fi96++ {
                v2 -= 1
                _ = fi96
            }
        default:
            var v97 uint8 = 42
            _ = v97
        }
        _ = v32
    }
    for fi98 := 0; fi98 < 4; fi98++ {
        v2 = v2
        if false {
            var v99 int32 = -2147483648
            _ = v99
        }
        _ = fi98
    }
    var v100 uint = 965
    fmt.Println(v100, v2)
    fmt.Println(v100, v100)
    if len(v4) > 0 {
        v4[0] = 42
    }
    for fi101 := 0; fi101 < 17; fi101++ {
        v100 = v100
        _ = fi101
    }
    var v102 int16 = 1000
    switch ("a" + "") {
    case "":
        if false {
            if false {
                if true {
                    delete(v3, 2)
                    var v103 int16 = int16(v100)
                    _ = v103
                }
                var v104 float64 = 1e10
                _ = v104
            } else {
                v102 = v102
                if true {
                    fmt.Println(v102, v100)
                    v4 = append(v4, 65535)
                    if (!true) {
                        fmt.Println(v2)
                        v105 := v3[587]
                        _ = v105
                    } else {
                        var v106 int = 4294967295
                        switch "\n\t" {
                        case "foo bar":
                            var v107 int8 = (v2 | -1)
                            _ = v107
                        case "hello":
                            v106 += v106
                        default:
                            var v108 uint32 = 2147483647
                            _ = v108
                        }
                        var v109 float64 = 1.5
                        _ = v106
                        _ = v109
                    }
                } else {
                    fmt.Println(v102, v100, v100)
                    switch (("foo bar" + ((("" + "") + "foo bar") + (("" + "a") + "\n\t"))) + ("" + "foo bar")) {
                    case "":
                        v110 := GenS7{F0: ("a" + "")}
                        _ = v110
                    default:
                        v100 = (^(v100))
                    }
                }
                switch "hello" {
                case "hello":
                    v111 := GenS4{F0: (^(v2)), F1: 402, F2: (v2 / ((-1) | 1))}
                    _ = v111
                default:
                    var v112 string = ((("a" + "foo bar") + ("\n\t" + "\n\t")) + "foo bar")
                    _ = v112
                }
            }
            fmt.Println(v2, v100)
        }
    case "a":
        var v113 float32 = 0.1
        _ = v113
    case "\n\t":
        for fi114 := 0; fi114 < 12; fi114++ {
            v115 := GenS8{F0: ("hello" + "")}
            if false {
                var v116 uint16 = 42
                v102 = v102
                v116 = v116
                _ = v116
            }
            var v117 uint16 = 1
            _ = fi114
            _ = v115
            _ = v117
        }
    default:
        if true {
            v102 = (v102 % ((v102) | 1))
            for fi118 := 0; fi118 < 10; fi118++ {
                var v119 uint16 = 1
                _ = fi118
                _ = v119
            }
            for fi120 := 0; fi120 < 19; fi120++ {
                switch "\n\t" {
                case "\n\t":
                    if (("foo bar" == (("" + "foo bar") + "hello")) || true) {
                        v4 = append(v4, 42)
                        switch ("foo bar" + ("\n\t" + ("\n\t" + ""))) {
                        case "a":
                            var v121 uint8 = 42
                            _ = v121
                        case "foo bar":
                            v3[0] = 0
                        case "hello":
                            v102 = 197
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                    }
                case "a":
                    var v122 float64 = ((1.5 * 1e10) / (float64(1e10) - (3.14159 * 0.1)))
                    _ = v122
                }
                v123 := GenS7{F0: "hello"}
                for fi124 := 0; fi124 < 11; fi124++ {
                    if ((42 != (v2 >> ((v100) & 63))) && (0.0 > (0.1 - (((3.14159 * 1e10) - (0.1 - (-2.25 - 0.0))) * -2.25)))) {
                        var v125 float64 = 0.0
                        var v126 int = fi120
                        if true {
                            delete(v3, 42)
                            var v127 int32 = -2147483648
                            _ = v127
                        }
                        _ = v125
                        _ = v126
                    }
                    switch int(v100) {
                    case 32767:
                        v128 := GenS8{F0: ("" + "")}
                        _ = v128
                    case 42:
                        var v129 bool = ((v102 % (((v102 >> (((v100 % ((1) | 1))) & 63))) | 1)) < (v102 * v102))
                        _ = v129
                    }
                    _ = fi124
                }
                _ = fi120
                _ = v123
            }
        } else {
            v100 = v100
            for fi130 := 0; fi130 < 3; fi130++ {
                var v131 int16 = (v102 << ((v100) & 63))
                v2 -= (v2 ^ v2)
                _ = fi130
                _ = v131
            }
        }
    }
    for fi132 := 0; fi132 < 7; fi132++ {
        v100 -= 0
        _ = fi132
    }
    switch "\n\t" {
    case "hello":
        if len(v4) > 0 {
            v4[0] = 2
        }
    case "":
        switch ("hello" + ("foo bar" + "")) {
        case "foo bar":
            if false {
                v133 := GenS9{F0: 0, F1: 42}
                for fi134 := 0; fi134 < 13; fi134++ {
                    fmt.Println(v102, v100, v100)
                    var v135 uint64 = 643
                    _ = fi134
                    _ = v135
                }
                _ = v133
            } else {
                var v136 uint16 = 976
                for fi137 := 0; fi137 < 12; fi137++ {
                    for fi138 := 0; fi138 < 18; fi138++ {
                        var v139 uint8 = 2
                        if (0 != 42) {
                            v139 = v139
                            var v140 int8 = (v2 << ((1) & 63))
                            v140 -= -128
                            _ = v140
                        }
                        _ = fi138
                        _ = v139
                    }
                    var v141 float32 = -2.25
                    v142 := GenS2{F0: 42, F1: true, F2: 4294967295}
                    _ = fi137
                    _ = v141
                    _ = v142
                }
                delete(v3, 720)
                _ = v136
            }
        case "a":
            var v143 float32 = 1.5
            _ = v143
        default:
            v144 := GenS7{F0: ("\n\t" + "hello")}
            _ = v144
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v100
    _ = v102
    _ = v2
    _ = v3
    _ = v4
}
