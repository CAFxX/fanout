package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 int64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint; F1 uint }
type GenS12 struct { F0 uint64; F1 string; F2 uint32 }
type GenS13 struct { F0 uint16 }
func genFn0(p0 int64, p1 int16, p2 uint16) uint {
    v6 := []int{-868, 0, 2147483647}
    var v7 float32 = 3.14159
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v8 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 bool, p1 string, p2 int64) uint {
    if p0 {
        p1 = ("\n\t" + ("\n\t" + p1))
    }
    p0 = true
    _ = p0
    _ = p1
    _ = p2
    return 462
}

func main() {
    var v1 string = "hello"
    var v2 int8 = -7
    var v3 int16 = -535
    v4 := make(map[int]string)
    var v5 int32 = 1
    if (((v5 >> ((4294967295) & 63)) >= v5) || (0 >= 0)) {
        switch "" {
        case "hello":
            switch v1 {
            case "a":
                delete(v4, 523)
            case "hello":
                v2 = v2
            case "foo bar":
                for fi9 := 0; fi9 < 1; fi9++ {
                    v10 := GenS1{F0: 1, F1: 41}
                    _ = fi9
                    _ = v10
                }
            default:
                if false {
                    if true {
                        for fi11 := 0; fi11 < 20; fi11++ {
                            v5 -= (v5 / ((v5) | 1))
                            _ = fi11
                        }
                        switch v1 {
                        case "foo bar":
                            if true {
                                v12 := []int{}
                                for fi13 := 0; fi13 < 2; fi13++ {
                                    var v14 uint = genFn1(true, v1, -931)
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    v15 := GenS1{F0: 18446744073709551615, F1: v2}
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                                if true {
                                    v1 = v1
                                    v2 += (v2 << ((2) & 63))
                                }
                                _ = v12
                            } else {
                                switch (v1 + v1) {
                                case "foo bar":
                                    v1 = "\n\t"
                                case "":
                                    var v16 bool = (3.14159 != 1.5)
                                    _ = v16
                                default:
                                    switch v1 {
                                    case "foo bar":
                                        switch 32768 {
                                        case -2147483648:
                                            if false {
                                                var v17 int = 0
                                                _ = v17
                                            } else {
                                                for fi18 := 0; fi18 < 17; fi18++ {
                                                    v3 = (-(v3))
                                                    switch ((v1 + v1) + v1) {
                                                    case "foo bar":
                                                        for fi19 := 0; fi19 < 2; fi19++ {
                                                            fmt.Println(fi19)
                                                            v20 := GenS2{F0: 957, F1: 0}
                                                            if true {
                                                                fmt.Println(fi18, v3)
                                                            } else {
                                                                v21 := GenS3{F0: 1}
                                                                v22 := []int{fi19}
                                                                _ = v21
                                                                _ = v22
                                                            }
                                                            v23 := GenS4{F0: 42}
                                                            _ = fi19
                                                            _ = v20
                                                            _ = v23
                                                        }
                                                    default:
                                                        switch v1 {
                                                        case "":
                                                            v24 := GenS4{F0: 47}
                                                            _ = v24
                                                        case "foo bar":
                                                            var v25 uint32 = 0
                                                            _ = v25
                                                        case "\n\t":
                                                            var v26 int = fi18
                                                            _ = v26
                                                        }
                                                    }
                                                    _ = fi18
                                                }
                                                v27 := GenS1{F0: 2, F1: (v2 - 100)}
                                                _ = v27
                                            }
                                        case 1:
                                            v2 = -84
                                        default:
                                            if (1.5 < 1e10) {
                                                for fi28 := 0; fi28 < 10; fi28++ {
                                                    fmt.Println(v3, v1, v1)
                                                    _ = fi28
                                                }
                                                if ((false || false) || (1 != 42)) {
                                                    switch v1 {
                                                    case "hello":
                                                        v1 = v1
                                                    case "":
                                                        for fi29 := 0; fi29 < 4; fi29++ {
                                                            v30 := GenS3{F0: 42}
                                                            if false {
                                                                for fi31 := 0; fi31 < 11; fi31++ {
                                                                    var v32 uint64 = 932
                                                                    fmt.Println(fi29, v1)
                                                                    if (v5 <= (v5 << ((798) & 63))) {
                                                                        var v33 uint = 42
                                                                        _ = v33
                                                                    }
                                                                    _ = fi31
                                                                    _ = v32
                                                                }
                                                            }
                                                            v3 = -121
                                                            _ = fi29
                                                            _ = v30
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(v2, v3, v2)
                                                    }
                                                    v34 := GenS2{F0: 42, F1: 4294967295}
                                                    if (!(42 == (v3 << ((837) & 63)))) {
                                                        v35 := v4[9223372036854775807]
                                                        v36 := []int{-9223372036854775808, 655, -2147483648}
                                                        var v37 uint16 = 1
                                                        _ = v35
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    _ = v34
                                                }
                                                var v38 float32 = 3.14159
                                                _ = v38
                                            } else {
                                                var v39 bool = (42 >= 18446744073709551615)
                                                switch "\n\t" {
                                                case "\n\t":
                                                    fmt.Println(v3, v39)
                                                default:
                                                    v40 := GenS5{F0: (^(v5))}
                                                    _ = v40
                                                }
                                                v39 = v39
                                                _ = v39
                                            }
                                        }
                                    case "\n\t":
                                        v41 := []int{9223372036854775807, 42}
                                        _ = v41
                                    default:
                                        fmt.Println(v3, v3)
                                    }
                                }
                                if true {
                                    v2 -= (^(v2))
                                    var v42 float32 = 1.5
                                    for fi43 := 0; fi43 < 9; fi43++ {
                                        v44 := GenS6{F0: -301}
                                        _ = fi43
                                        _ = v44
                                    }
                                    _ = v42
                                } else {
                                    v3 = (-(v3))
                                }
                            }
                        case "hello":
                            switch ("foo bar" + "foo bar") {
                            case "a":
                                _ = len(v4)
                            }
                        }
                        var v45 uint64 = 0
                        _ = v45
                    } else {
                        v5 -= (v5 * (v5 | 1))
                        switch "" {
                        case "":
                            var v46 uint16 = 32767
                            _ = v46
                        default:
                            switch -9223372036854775808 {
                            case 127:
                                switch ("\n\t" + "hello") {
                                case "hello":
                                    var v47 int = 406
                                    _ = v47
                                case "\n\t":
                                    if (42 >= 255) {
                                        for fi48 := 0; fi48 < 17; fi48++ {
                                            v49 := GenS7{F0: true, F1: -1}
                                            var v50 uint16 = 255
                                            v51 := GenS4{F0: 1}
                                            _ = fi48
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        }
                                        v3 = 32767
                                        if (!true) {
                                            v1 = "hello"
                                            v2 -= v2
                                        } else {
                                            for fi52 := 0; fi52 < 5; fi52++ {
                                                var v53 uint64 = 626
                                                if false {
                                                    for fi54 := 0; fi54 < 14; fi54++ {
                                                        v2 += v2
                                                        fmt.Println(fi52)
                                                        _ = fi54
                                                    }
                                                    if (1000 >= -32768) {
                                                        for fi55 := 0; fi55 < 7; fi55++ {
                                                            v56 := []int{}
                                                            var v57 uint8 = 1
                                                            var v58 int32 = (v5 >> ((65535) & 63))
                                                            if len(v56) >= 2 {
                                                                v56 = v56[1:len(v56)]
                                                            }
                                                            _ = fi55
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                        fmt.Println(fi52)
                                                        var v59 int32 = (^(v5))
                                                        _ = v59
                                                    }
                                                    v60 := []int{42, (fi52 | (^(fi52)))}
                                                    _ = v60
                                                } else {
                                                    fmt.Println(fi52)
                                                }
                                                _ = fi52
                                                _ = v53
                                            }
                                            v61 := []int{0, 32768, 0}
                                            fmt.Println(v1)
                                            _ = v61
                                        }
                                    }
                                case "a":
                                    fmt.Println(v3)
                                }
                            }
                        }
                    }
                    var v62 int32 = v5
                    _ = v62
                } else {
                    fmt.Println(v5, v1, v5)
                    if false {
                        var v63 uint64 = 42
                        v1 = v1
                        v4[1] = (v1 + v1)
                        _ = v63
                    }
                }
            }
        case "a":
            switch v1 {
            case "\n\t":
                v2 = 0
            default:
                var v64 uint8 = 2
                _ = v64
            }
        case "\n\t":
            if true {
                v3 -= -32768
                if true {
                    v3 += v3
                    v65 := GenS8{F0: -1}
                    _ = v65
                } else {
                    var v66 bool = ((v5 < v5) && (262 == 1))
                    switch -271 {
                    case -2147483648:
                        if (true || (false && v66)) {
                            v4[716] = ("a" + "hello")
                        } else {
                            var v67 int16 = 1
                            v67 = (v3 ^ (v3 % ((v3) | 1)))
                            _ = v67
                        }
                    }
                    _ = v66
                }
            } else {
                if false {
                    v68 := []int{}
                    _ = v68
                } else {
                    if (32767 > 0) {
                        if false {
                            _ = len(v4)
                            switch (v1 + v1) {
                            case "\n\t":
                                var v69 int64 = 9223372036854775807
                                _ = v69
                            case "foo bar":
                                fmt.Println(v3, v2)
                            default:
                                var v70 int32 = v5
                                _ = v70
                            }
                            for fi71 := 0; fi71 < 4; fi71++ {
                                v72 := GenS2{F0: genFn0(-1, -98, 255), F1: genFn0(42, (v3 / ((1000) | 1)), 527)}
                                _ = fi71
                                _ = v72
                            }
                        } else {
                            if true {
                                var v73 bool = false
                                for fi74 := 0; fi74 < 19; fi74++ {
                                    var v75 uint8 = 2
                                    if false {
                                        v1 = (v1 + (v1 + (v1 + "hello")))
                                    } else {
                                        if true {
                                            switch ((v1 + (((v1 + (("a" + "\n\t") + v1)) + "\n\t") + "\n\t")) + v1) {
                                            case "":
                                                v76 := []int{128, fi74}
                                                _ = v76
                                            case "\n\t":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v77 string = v1
                                                    _ = v77
                                                }
                                            default:
                                                v78 := GenS2{F0: genFn1((9223372036854775807 < 0), "hello", 9223372036854775807), F1: 2}
                                                _ = v78
                                            }
                                        } else {
                                            switch fi74 {
                                            case 794:
                                                v2 = (v2 * (v2 << ((18446744073709551615) & 63)))
                                            case 2147483647:
                                                for fi79 := 0; fi79 < 17; fi79++ {
                                                    v73 = v73
                                                    if (691 == 9223372036854775807) {
                                                        var v80 uint64 = 98
                                                        v81 := v4[int(v3)]
                                                        if v73 {
                                                            var v82 float32 = -2.25
                                                            v83 := GenS2{F0: 2, F1: 543}
                                                            _ = v82
                                                            _ = v83
                                                        } else {
                                                            v4[(fi74 << ((1) & 63))] = ("hello" + v81)
                                                            v84 := GenS9{F0: v75, F1: 2, F2: (v5 - v5)}
                                                            _ = v84
                                                        }
                                                        _ = v80
                                                        _ = v81
                                                    } else {
                                                        var v85 int = fi79
                                                        v73 = (!v73)
                                                        v86 := GenS9{F0: (v75 & v75), F1: 276, F2: v5}
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                    if (-2.25 >= 0.0) {
                                                        var v87 float32 = 0.1
                                                        _ = v87
                                                    } else {
                                                        v88 := []int{-9223372036854775808, 128, (fi79 / ((int(v3)) | 1)), (fi74 | (fi79 & (fi74 + (fi74 - int(v3)))))}
                                                        _ = v88
                                                    }
                                                    _ = fi79
                                                }
                                            }
                                            v3 = (v3 % ((v3) | 1))
                                        }
                                        var v89 uint64 = 18446744073709551615
                                        _ = v89
                                    }
                                    _ = fi74
                                    _ = v75
                                }
                                var v90 bool = v73
                                _ = v73
                                _ = v90
                            } else {
                                v91 := []int{-2147483648, 129, -532, -9223372036854775808}
                                _ = v91
                            }
                            var v92 uint32 = 2147483647
                            _ = v92
                        }
                        v2 = 116
                    } else {
                        fmt.Println(v2, v1, v1)
                        for fi93 := 0; fi93 < 8; fi93++ {
                            for fi94 := 0; fi94 < 6; fi94++ {
                                v1 = "foo bar"
                                switch v1 {
                                case "hello":
                                    v1 = ((v1 + ((v1 + (((((v1 + v1) + (v1 + v1)) + (v1 + v1)) + v1) + (v1 + "foo bar"))) + v1)) + "foo bar")
                                }
                                var v95 int8 = (v2 + (v2 * 100))
                                _ = fi94
                                _ = v95
                            }
                            v96 := GenS10{F0: (fi93 >> ((70) & 63))}
                            _ = fi93
                            _ = v96
                        }
                        switch v1 {
                        case "\n\t":
                            switch 128 {
                            case 32768:
                                var v97 uint16 = 42
                                _ = v97
                            case 1:
                                var v98 uint64 = 1
                                _ = v98
                            }
                        default:
                            if ((v2 << ((2) & 63)) <= (^(v2))) {
                                v3 += v3
                                v99 := GenS11{F0: 42, F1: 1}
                                for fi100 := 0; fi100 < 19; fi100++ {
                                    v1 = v1
                                    var v101 uint16 = 103
                                    for fi102 := 0; fi102 < 16; fi102++ {
                                        var v103 uint64 = 18446744073709551615
                                        _ = fi102
                                        _ = v103
                                    }
                                    _ = fi100
                                    _ = v101
                                }
                                _ = v99
                            }
                        }
                    }
                    if (!true) {
                        v104 := v4[791]
                        _ = v104
                    }
                }
                var v105 int8 = 18
                if (!false) {
                    v1 = "foo bar"
                    var v106 string = "hello"
                    v5 = (v5 * (v5 ^ (v5 % ((v5) | 1))))
                    _ = v106
                }
                _ = v105
            }
        default:
            if (false || (!(true && false))) {
                if (-195 == 4294967295) {
                    fmt.Println(v2)
                    v3 = v3
                    v107 := []int{127, 0}
                    _ = v107
                } else {
                    fmt.Println(v3, v2)
                    switch (((v1 + (v1 + v1)) + (v1 + (v1 + "a"))) + v1) {
                    case "foo bar":
                        v108 := GenS8{F0: 9223372036854775807}
                        _ = v108
                    case "":
                        var v109 uint = 36
                        _ = v109
                    default:
                        v110 := []int{}
                        _ = v110
                    }
                    fmt.Println(v3, v2, v3)
                }
                for fi111 := 0; fi111 < 14; fi111++ {
                    var v112 int16 = 42
                    switch v1 {
                    case "":
                        v113 := []int{int(v112), 1, int(v112), fi111}
                        _ = v113
                    default:
                        fmt.Println(v3, v2)
                    }
                    _ = fi111
                    _ = v112
                }
            }
        }
        var v114 int8 = v2
        _ = v114
    }
    v115 := []int{0, -1}
    fmt.Println(v3, v2)
    v116 := GenS5{F0: v5}
    switch (("foo bar" + "foo bar") + "") {
    case "foo bar":
        v117 := GenS8{F0: -502}
        _ = v117
    case "a":
        for fi118 := 0; fi118 < 12; fi118++ {
            v119 := GenS3{F0: (-(fi118))}
            v120 := GenS10{F0: 741}
            _ = fi118
            _ = v119
            _ = v120
        }
    }
    fmt.Println(v2)
    switch 1 {
    case 128:
        if false {
            fmt.Println(v5, v5)
            for fi121 := 0; fi121 < 6; fi121++ {
                switch (-(fi121)) {
                case 42:
                    v2 = -1
                case 4294967295:
                    v3 += 68
                case 1:
                    v2 = (-(v2))
                default:
                    switch fi121 {
                    case 127:
                        for fi122 := 0; fi122 < 18; fi122++ {
                            v2 = v2
                            _ = fi122
                        }
                    case 42:
                        fmt.Println(v2, v1)
                    default:
                        v123 := GenS3{F0: fi121}
                        _ = v123
                    }
                }
                var v124 int8 = (v2 * v2)
                if true {
                    for fi125 := 0; fi125 < 8; fi125++ {
                        var v126 int32 = (v5 << ((0) & 63))
                        _ = fi125
                        _ = v126
                    }
                    if len(v115) >= 2 {
                        v115 = v115[1:len(v115)]
                    }
                    v3 -= (-(v3))
                } else {
                    if (!(!true)) {
                        v124 = v124
                        var v127 uint32 = 65535
                        for fi128 := 0; fi128 < 3; fi128++ {
                            var v129 int16 = v3
                            fmt.Println(v127, v129, v3)
                            var v130 float64 = ((((1e10 - ((0.0 / 1.5) * (0.0 / -2.25))) - 1e10) + 1e10) + (0.1 * (1.5 / (3.14159 + (0.1 + 1e10)))))
                            v131 := GenS3{F0: (^(fi128))}
                            _ = fi128
                            _ = v129
                            _ = v130
                            _ = v131
                        }
                        _ = v127
                    } else {
                        v132 := GenS5{F0: v5}
                        _ = v132
                    }
                    fmt.Println(v1)
                }
                _ = fi121
                _ = v124
            }
            for fi133 := 0; fi133 < 3; fi133++ {
                var v134 uint32 = 2147483648
                if true {
                    fmt.Println(v3, v3)
                    if (!false) {
                        for fi135 := 0; fi135 < 19; fi135++ {
                            v134 = (v134 | 65535)
                            switch v1 {
                            case "hello":
                                v136 := GenS12{F0: 14, F1: v1, F2: v134}
                                _ = v136
                            case "foo bar":
                                fmt.Println(fi135)
                            default:
                                v134 = (v134 + (v134 - v134))
                            }
                            _ = fi135
                        }
                        v1 = ""
                        v3 += (v3 ^ 310)
                    }
                    for fi137 := 0; fi137 < 8; fi137++ {
                        _ = len(v115)
                        _ = fi137
                    }
                }
                for fi138 := 0; fi138 < 13; fi138++ {
                    var v139 int16 = (v3 ^ (-(v3)))
                    _ = fi138
                    _ = v139
                }
                _ = fi133
                _ = v134
            }
        } else {
            v3 += v3
            v1 = ((v1 + "") + v1)
        }
    case -9223372036854775808:
        fmt.Println(v2, v1)
    case 127:
        _ = len(v115)
    default:
        v2 = v2
    }
    var v140 uint = genFn0(-799, (v3 * 0), 2)
    switch v1 {
    case "":
        for fi141 := 0; fi141 < 6; fi141++ {
            if (((fi141 | (fi141 + 280)) == 169) && (!(!(1 == 1)))) {
                if len(v115) > 0 {
                    v115[0] = (fi141 ^ (^(fi141)))
                }
                var v142 bool = true
                v1 = v1
                _ = v142
            }
            _ = fi141
        }
    }
    for fi143 := 0; fi143 < 3; fi143++ {
        v144 := GenS13{F0: 32767}
        var v145 int64 = -669
        _ = fi143
        _ = v144
        _ = v145
    }
    if len(v115) > 0 {
        v115[0] = 128
    }
    var v146 float64 = -2.25
    v4[-9223372036854775808] = (v1 + v1)
    if len(v115) >= 2 {
        v115 = v115[1:len(v115)]
    }
    v140 = (v140 % ((42) | 1))
    if false {
        var v147 bool = (true || (0 != 2))
        _ = v147
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v115:", len(v115), v115)
    fmt.Println("v140:", v140)
    fmt.Println("v146:", v146)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v115
    _ = v116
    _ = v140
    _ = v146
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
