package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint16; F1 uint; F2 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 uint8; F2 int32 }
type GenS8 struct { F0 int8; F1 bool; F2 int64 }
type GenS9 struct { F0 int; F1 int64 }
type GenS10 struct { F0 uint32; F1 int; F2 uint32 }
type GenS11 struct { F0 int32; F1 int }
type GenS12 struct { F0 int64; F1 string }
func genFn0(p0 int8, p1 int) int8 {
    v7 := GenS1{F0: 4294967295, F1: (p1 >> ((506) & 63))}
    p1 -= (p1 - p1)
    v8 := GenS1{F0: 2147483648, F1: (p1 / ((p1) | 1))}
    p0 += (p0 + p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (-(p0))
}
func genFn1() uint32 {
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := []int{(fi9 ^ fi9), fi9, -1, 32767}
        if (0 > 42) {
            var v11 uint64 = 67
            if (1 != -2147483648) {
                v11 = 0
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 int = (fi12 & (fi9 * fi9))
                    v14 := GenS2{F0: 255, F1: ("\n\t" + "\n\t")}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            } else {
                if ((true && (v11 != (v11 | (v11 & (^(v11)))))) && ((fi9 << ((628) & 63)) != -2147483648)) {
                    if (42 >= 411) {
                        var v15 int64 = 677
                        switch (^(fi9)) {
                        case 127:
                            v16 := GenS2{F0: 0, F1: "hello"}
                            _ = v16
                        default:
                            switch 128 {
                            case 1:
                                fmt.Println(v15, v11)
                            default:
                                v15 -= 0
                            }
                        }
                        _ = v15
                    } else {
                        var v17 uint = 836
                        switch fi9 {
                        case 0:
                            switch (("\n\t" + ("foo bar" + ("" + "\n\t"))) + "hello") {
                            case "a":
                                if false {
                                    v11 += (v11 / ((v11) | 1))
                                } else {
                                    var v18 uint32 = 2
                                    _ = v18
                                }
                            case "hello":
                                if (739 > 67) {
                                    if true {
                                        for fi19 := 0; fi19 < 6; fi19++ {
                                            v17 = (v17 * v17)
                                            switch ("" + "foo bar") {
                                            case "\n\t":
                                                v17 = v17
                                            case "a":
                                                v17 -= (v17 | v17)
                                            default:
                                                v20 := GenS3{F0: 2, F1: v17, F2: ""}
                                                _ = v20
                                            }
                                            _ = fi19
                                        }
                                    }
                                    if len(v10) > 0 {
                                        v10[0] = (fi9 % ((269) | 1))
                                    }
                                }
                            default:
                                switch "foo bar" {
                                case "foo bar":
                                    v21 := GenS3{F0: 242, F1: v17, F2: ("hello" + ("a" + "hello"))}
                                    _ = v21
                                case "":
                                    v17 += (v17 >> ((v17) & 63))
                                case "\n\t":
                                    var v22 uint = v17
                                    _ = v22
                                default:
                                    var v23 int8 = 42
                                    _ = v23
                                }
                            }
                        case 42:
                            v24 := GenS1{F0: 565, F1: -2147483648}
                            _ = v24
                        }
                        _ = v17
                    }
                    v11 += (v11 / ((v11) | 1))
                    var v25 uint8 = 1
                    _ = v25
                }
            }
            var v26 int = fi9
            _ = v11
            _ = v26
        }
        _ = len(v10)
        _ = fi9
        _ = v10
    }
    if false {
        switch 961 {
        case 1:
            var v27 int32 = 42
            _ = v27
        case -61:
            v28 := GenS1{F0: 962, F1: 42}
            _ = v28
        case 796:
            var v29 uint16 = 2
            _ = v29
        }
        if (!true) {
            var v30 float64 = 0.0
            fmt.Println(v30)
            var v31 bool = (0 <= 65535)
            _ = v30
            _ = v31
        }
    } else {
        v32 := []int{127, 9223372036854775807, 127, 127}
        switch "hello" {
        case "hello":
            var v33 float32 = -2.25
            _ = v33
        case "\n\t":
            var v34 bool = (float32(1e10) == float32(float64(((float32((3.14159 + 1e10)) - -2.25) / (1e10 / 3.14159)))))
            _ = v34
        case "foo bar":
            var v35 float32 = (1.5 + 3.14159)
            _ = v35
        default:
            switch "\n\t" {
            case "a":
                if false {
                    v36 := GenS4{F0: 2}
                    var v37 int = -333
                    if (((!false) && (true && (!(v37 == (v37 / (((v37 ^ (v37 % ((v37) | 1)))) | 1)))))) || true) {
                        v37 -= v37
                        var v38 string = "a"
                        v39 := GenS5{F0: 2, F1: -1}
                        _ = v38
                        _ = v39
                    } else {
                        switch (v37 - 32768) {
                        case 127:
                            for fi40 := 0; fi40 < 12; fi40++ {
                                var v41 int8 = -128
                                if false {
                                    for fi42 := 0; fi42 < 19; fi42++ {
                                        if true {
                                            fmt.Println(v41, fi42, v41)
                                            _ = len(v32)
                                        } else {
                                            var v43 int64 = -1
                                            fmt.Println(v41, v43)
                                            _ = v43
                                        }
                                        _ = fi42
                                    }
                                    var v44 string = ("hello" + "foo bar")
                                    _ = v44
                                } else {
                                    v32 = append(v32, fi40)
                                }
                                if (!false) {
                                    switch int(v41) {
                                    case 0:
                                        var v45 int32 = 42
                                        _ = v45
                                    default:
                                        _ = len(v32)
                                    }
                                } else {
                                    v41 -= v41
                                    if len(v32) > 0 {
                                        v32[0] = (^(v37))
                                    }
                                    if len(v32) > 0 {
                                        v32[0] = (v37 + (fi40 + fi40))
                                    }
                                }
                                _ = fi40
                                _ = v41
                            }
                        case 32768:
                            var v46 uint8 = 1
                            _ = v46
                        default:
                            var v47 float32 = 0.0
                            _ = v47
                        }
                        if ((("hello" + "foo bar") + "a") < ("a" + "foo bar")) {
                            if len(v32) >= 2 {
                                v32 = v32[1:len(v32)]
                            }
                            if ((0.0 > 3.14159) && false) {
                                v32 = append(v32, 128)
                                v37 = v37
                            }
                        } else {
                            switch (v37 & (^(v37))) {
                            case 9223372036854775807:
                                if true {
                                    var v48 float64 = 0.0
                                    _ = v48
                                } else {
                                    for fi49 := 0; fi49 < 18; fi49++ {
                                        _ = len(v32)
                                        v37 = v37
                                        switch 2147483647 {
                                        case 32768:
                                            if len(v32) >= 2 {
                                                v32 = v32[1:len(v32)]
                                            }
                                        case -9223372036854775808:
                                            var v50 uint16 = 32767
                                            _ = v50
                                        case -707:
                                            fmt.Println(fi49, fi49)
                                        default:
                                            var v51 float32 = 1.5
                                            _ = v51
                                        }
                                        _ = fi49
                                    }
                                    for fi52 := 0; fi52 < 5; fi52++ {
                                        var v53 int32 = 1
                                        v37 -= 194
                                        for fi54 := 0; fi54 < 17; fi54++ {
                                            var v55 uint16 = 32767
                                            if len(v32) > 0 {
                                                v32[0] = fi52
                                            }
                                            var v56 int = fi54
                                            _ = fi54
                                            _ = v55
                                            _ = v56
                                        }
                                        _ = fi52
                                        _ = v53
                                    }
                                    v37 = (v37 - 32767)
                                }
                            case 0:
                                var v57 int64 = 1
                                _ = v57
                            }
                            v37 -= (v37 | (v37 << ((255) & 63)))
                        }
                        for fi58 := 0; fi58 < 11; fi58++ {
                            v37 = fi58
                            _ = fi58
                        }
                    }
                    _ = v36
                    _ = v37
                }
            }
        }
        for fi59 := 0; fi59 < 7; fi59++ {
            var v60 int8 = 1
            _ = fi59
            _ = v60
        }
        _ = v32
    }
    for fi61 := 0; fi61 < 5; fi61++ {
        if ((genFn0(-128, fi61) >= genFn0(genFn0(78, -828), fi61)) && (!(true && false))) {
            for fi62 := 0; fi62 < 1; fi62++ {
                switch "a" {
                case "":
                    v63 := GenS2{F0: 1, F1: "\n\t"}
                    _ = v63
                case "foo bar":
                    if false {
                        var v64 bool = false
                        _ = v64
                    } else {
                        v65 := GenS1{F0: 378, F1: 32767}
                        v66 := []int{(fi62 << ((65535) & 63)), (fi61 ^ fi62), fi61, fi62}
                        _ = v65
                        _ = v66
                    }
                case "\n\t":
                    switch fi61 {
                    case 134:
                        v67 := GenS5{F0: 1, F1: 1}
                        _ = v67
                    case -1:
                        if true {
                            var v68 int16 = 0
                            _ = v68
                        }
                    }
                default:
                    v69 := []int{}
                    _ = v69
                }
                var v70 int32 = 2147483647
                v71 := GenS6{F0: v70}
                _ = fi62
                _ = v70
                _ = v71
            }
        }
        if false {
            switch "foo bar" {
            case "":
                v72 := []int{964, 128}
                _ = v72
            default:
                switch fi61 {
                case 4294967295:
                    fmt.Println(fi61)
                case 1:
                    var v73 float32 = (1e10 + 1.5)
                    _ = v73
                case 128:
                    var v74 uint = 255
                    _ = v74
                }
            }
            v75 := GenS1{F0: 0, F1: (fi61 % (((-(fi61))) | 1))}
            _ = v75
        } else {
            var v76 int16 = 1
            _ = v76
        }
        switch "foo bar" {
        case "":
            var v77 uint32 = 716
            _ = v77
        }
        _ = fi61
    }
    return 699
}

func main() {
    var v1 int = 2147483647
    var v2 uint = 2
    v3 := []string{}
    var v4 string = ""
    v5 := make(map[string]uint32)
    v6 := []string{v4, ("\n\t" + ("" + v4)), v4}
    v78 := GenS1{F0: 2, F1: (v1 << ((v2) & 63))}
    var v79 uint8 = 255
    var v80 uint32 = genFn1()
    v2 = (v2 + 392)
    var v81 uint64 = 42
    for fi82 := 0; fi82 < 10; fi82++ {
        v2 = (v2 + 18446744073709551615)
        v4 = v4
        if true {
            var v83 uint = v2
            _ = v83
        }
        _ = fi82
    }
    var v84 int = v1
    v4 = "hello"
    v80 = (^(v80))
    if ((!(!(255 >= 0))) || (uint(v80) > uint(v81))) {
        v2 = 42
    } else {
        var v85 int8 = 100
        var v86 float32 = 0.0
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v85
        _ = v86
    }
    v87 := GenS2{F0: 2, F1: "\n\t"}
    switch (v1 & 540) {
    case -221:
        if false {
            v1 -= (v84 | 1)
            delete(v5, "hello")
        } else {
            var v88 uint = (v2 % (((v2 % ((42) | 1))) | 1))
            var v89 int64 = -9223372036854775808
            _ = v88
            _ = v89
        }
    default:
        for fi90 := 0; fi90 < 1; fi90++ {
            v2 -= v2
            _ = fi90
        }
    }
    for fi91 := 0; fi91 < 11; fi91++ {
        if (-1 >= -1) {
            var v92 uint16 = 42
            for fi93 := 0; fi93 < 19; fi93++ {
                if false {
                    v81 = (v81 / (((^(v81))) | 1))
                    var v94 int16 = 1000
                    _ = v94
                } else {
                    delete(v5, ("\n\t" + ""))
                    v95 := v5[(("foo bar" + v4) + ("hello" + v4))]
                    _ = v95
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v92 = v92
                if (1 > 1000) {
                    v96 := GenS7{F0: 0, F1: uint8(v2), F2: 42}
                    v79 = (v79 ^ 87)
                    var v97 int8 = 0
                    if (1 != v1) {
                        switch (fi91 * -1) {
                        case -9223372036854775808:
                            var v98 string = v4
                            _ = v98
                        default:
                            fmt.Println(v81, v80)
                        }
                        if (!false) {
                            _ = len(v3)
                        }
                    } else {
                        switch (v4 + "") {
                        case "\n\t":
                            fmt.Println(v84)
                        case "hello":
                            if (-677 <= 423) {
                                if true {
                                    switch 9223372036854775807 {
                                    case 32767:
                                        var v99 int8 = (^(v97))
                                        _ = v99
                                    default:
                                        if false {
                                            var v100 float64 = -2.25
                                            if true {
                                                fmt.Println(v100, v97)
                                            }
                                            _ = v100
                                        }
                                    }
                                    fmt.Println(v97)
                                    switch v1 {
                                    case 32767:
                                        switch v4 {
                                        case "foo bar":
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v101 int32 = -2147483648
                                                _ = v101
                                            default:
                                                v102 := GenS6{F0: -1}
                                                _ = v102
                                            }
                                        case "":
                                            var v103 uint = (v2 & (v2 % ((v2) | 1)))
                                            _ = v103
                                        case "a":
                                            if false {
                                                var v104 bool = false
                                                v105 := GenS8{F0: (-(v97)), F1: v104, F2: -572}
                                                _ = v104
                                                _ = v105
                                            }
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = (v4 + "\n\t")
                                            }
                                        }
                                        var v106 bool = false
                                        var v107 uint = (^(v2))
                                        _ = v106
                                        _ = v107
                                    case 128:
                                        var v108 float32 = 0.1
                                        _ = v108
                                    default:
                                        if ((true || false) && (false || true)) {
                                            if true {
                                                if false {
                                                    if false {
                                                        v84 -= fi91
                                                        switch (fi93 - fi91) {
                                                        case -608:
                                                            if true {
                                                                v109 := GenS2{F0: v92, F1: "\n\t"}
                                                                _ = v109
                                                            } else {
                                                                if (v97 != (v97 % ((0) | 1))) {
                                                                    fmt.Println(fi93, v92)
                                                                } else {
                                                                    fmt.Println(fi91)
                                                                    switch v4 {
                                                                    case "hello":
                                                                        var v110 uint16 = uint16(v80)
                                                                        _ = v110
                                                                    default:
                                                                        if (!false) {
                                                                            v5[v4] = uint32(v2)
                                                                            v79 += (^(v79))
                                                                        } else {
                                                                            if true {
                                                                                v111 := GenS9{F0: (v1 + (v84 / ((2147483647) | 1))), F1: 0}
                                                                                v112 := GenS10{F0: genFn1(), F1: (v84 << (((v2 % (((v2 & uint(v79))) | 1))) & 63)), F2: v80}
                                                                                fmt.Println(fi91, v92, v79)
                                                                                var v113 int = fi93
                                                                                _ = v111
                                                                                _ = v112
                                                                                _ = v113
                                                                            }
                                                                            var v114 uint16 = v92
                                                                            _ = v114
                                                                        }
                                                                    }
                                                                    v115 := GenS8{F0: 1, F1: ((0 >= -32768) || false), F2: -1}
                                                                    v1 -= -133
                                                                    _ = v115
                                                                }
                                                            }
                                                            var v116 int = fi93
                                                            var v117 uint32 = 942
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                        _ = len(v5)
                                                    }
                                                    switch v4 {
                                                    case "foo bar":
                                                        var v118 int64 = -9223372036854775808
                                                        _ = v118
                                                    }
                                                    switch (v4 + "\n\t") {
                                                    case "\n\t":
                                                        v84 = -496
                                                    case "foo bar":
                                                        var v119 int = fi91
                                                        _ = v119
                                                    case "":
                                                        delete(v5, (((v4 + v4) + ("a" + (("a" + "a") + ((((("" + v4) + (v4 + (v4 + v4))) + v4) + v4) + (v4 + "hello"))))) + v4))
                                                    default:
                                                        v120 := GenS11{F0: -2147483648, F1: 127}
                                                        _ = v120
                                                    }
                                                }
                                                if false {
                                                    var v121 uint16 = uint16(v80)
                                                    _ = len(v6)
                                                    v92 += (v92 >> ((uint(v79)) & 63))
                                                    _ = v121
                                                } else {
                                                    fmt.Println(v79, fi93)
                                                    var v122 int64 = 42
                                                    _ = v122
                                                }
                                                v123 := GenS1{F0: uint32(v81), F1: (v1 * int(v79))}
                                                _ = v123
                                            }
                                            if true {
                                                v2 = 102
                                                if (1000 >= 0) {
                                                    switch "" {
                                                    case "\n\t":
                                                        _ = len(v5)
                                                        if false {
                                                            var v124 float32 = 1.5
                                                            var v125 uint32 = 2147483647
                                                            var v126 int64 = -1
                                                            _ = v124
                                                            _ = v125
                                                            _ = v126
                                                        }
                                                    case "":
                                                        switch "\n\t" {
                                                        case "":
                                                            var v127 uint8 = (v79 & 0)
                                                            _ = v127
                                                        case "\n\t":
                                                            if false {
                                                                switch fi91 {
                                                                case 789:
                                                                    v97 = (v97 - (v97 ^ 100))
                                                                case 32768:
                                                                    var v128 string = v4
                                                                    _ = v128
                                                                case 9223372036854775807:
                                                                    v129 := GenS11{F0: 0, F1: (-(v84))}
                                                                    _ = v129
                                                                }
                                                                fmt.Println(v81, v92)
                                                            } else {
                                                                switch v84 {
                                                                case -351:
                                                                    v4 = ("\n\t" + (v4 + ((v4 + "") + (((v4 + "hello") + (v4 + (((v4 + (("hello" + v4) + (v4 + ""))) + ("hello" + v4)) + ((v4 + v4) + v4)))) + "foo bar"))))
                                                                }
                                                                v5["hello"] = (v80 | (v80 << ((uint(v80)) & 63)))
                                                                var v130 uint64 = v81
                                                                _ = v130
                                                            }
                                                        }
                                                    case "foo bar":
                                                        v131 := GenS8{F0: v97, F1: false, F2: 738}
                                                        _ = v131
                                                    default:
                                                        v1 -= 128
                                                    }
                                                    var v132 int32 = -1
                                                    switch "hello" {
                                                    case "hello":
                                                        if (!(!(-1 != -1))) {
                                                            fmt.Println(v97, v79)
                                                            var v133 uint64 = (v81 >> ((uint(v81)) & 63))
                                                            _ = v133
                                                        }
                                                    case "a":
                                                        fmt.Println(fi91)
                                                    case "foo bar":
                                                        if false {
                                                            v134 := GenS1{F0: v80, F1: (v1 ^ (v1 * 0))}
                                                            var v135 uint64 = 1
                                                            _ = v134
                                                            _ = v135
                                                        }
                                                    default:
                                                        v136 := GenS12{F0: 9223372036854775807, F1: ("a" + v4)}
                                                        _ = v136
                                                    }
                                                    _ = len(v6)
                                                    switch v4 {
                                                    case "foo bar":
                                                        var v137 uint8 = v79
                                                        v81 += v81
                                                        _ = v137
                                                    default:
                                                        switch "" {
                                                        case "hello":
                                                            _ = len(v3)
                                                        case "\n\t":
                                                            if false {
                                                                v80 += v80
                                                                var v138 int8 = v97
                                                                var v139 string = ((("" + v4) + (v4 + "hello")) + v4)
                                                                _ = v138
                                                                _ = v139
                                                            }
                                                        case "a":
                                                            var v140 int32 = 42
                                                            _ = v140
                                                        default:
                                                            switch v4 {
                                                            case "hello":
                                                                var v141 int8 = genFn0((v97 >> (((v2 ^ (v2 >> ((v2) & 63)))) & 63)), (-(v84)))
                                                                _ = v141
                                                            }
                                                        }
                                                    }
                                                    var v142 uint64 = 18446744073709551615
                                                    _ = v132
                                                    _ = v142
                                                } else {
                                                    if (0 > uint(v80)) {
                                                        v2 -= 1
                                                        v2 = 237
                                                        var v143 string = v4
                                                        var v144 uint = (v2 * 524)
                                                        fmt.Println(v144, v1, v2)
                                                        _ = v143
                                                        _ = v144
                                                    } else {
                                                        fmt.Println(v81, v92)
                                                    }
                                                    v92 -= (v92 % ((42) | 1))
                                                    v145 := GenS2{F0: v92, F1: v4}
                                                    _ = v145
                                                }
                                                var v146 uint8 = 128
                                                _ = v146
                                            }
                                        } else {
                                            v4 = v4
                                            fmt.Println(v4)
                                            var v147 uint = 4294967295
                                            _ = v147
                                        }
                                    }
                                } else {
                                    var v148 uint32 = (v80 & v80)
                                    fmt.Println(v1, v2)
                                    v80 = 65535
                                    _ = v148
                                }
                                if (255 < 0) {
                                    v80 -= 0
                                    var v149 bool = (-1 < 1)
                                    _ = v149
                                }
                                v1 = 1
                                fmt.Println(v2)
                                fmt.Println(v80)
                            }
                        case "a":
                            v4 = (v4 + v4)
                            v81 = (v81 ^ (v81 - v81))
                            fmt.Println(v80)
                        default:
                            var v150 uint = v2
                            _ = v150
                        }
                        if (100 > 89) {
                            v151 := GenS3{F0: v92, F1: 1, F2: "hello"}
                            fmt.Println(v4, v81)
                            _ = v151
                        }
                    }
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    _ = v96
                    _ = v97
                }
                _ = fi93
            }
            _ = v92
        }
        _ = fi91
    }
    v4 = (v4 + v4)
    _ = len(v5)
    var v152 int64 = -244
    v152 = (v152 / ((v152) | 1))
    for fi153 := 0; fi153 < 13; fi153++ {
        if ((v81 <= (v81 & (v81 >> ((2) & 63)))) || false) {
            var v154 float64 = (1.5 * 0.0)
            for fi155 := 0; fi155 < 18; fi155++ {
                v6 = append(v6, "hello")
                var v156 uint8 = (v79 >> ((v2) & 63))
                _ = fi155
                _ = v156
            }
            switch -2147483648 {
            case 1:
                v80 += v80
            }
            _ = v154
        }
        fmt.Println(v1)
        for fi157 := 0; fi157 < 17; fi157++ {
            if ((true && true) && (!false)) {
                if (-1 <= -2147483648) {
                    v158 := GenS10{F0: (v80 % ((v80) | 1)), F1: (v1 + 0), F2: v80}
                    fmt.Println(fi153, v79, v79)
                    v159 := GenS12{F0: v152, F1: v4}
                    _ = v158
                    _ = v159
                }
                if false {
                    var v160 int64 = -9223372036854775808
                    var v161 int = 1
                    _ = v160
                    _ = v161
                }
            } else {
                v79 = (v79 - v79)
                v162 := GenS2{F0: 0, F1: ((v4 + ("foo bar" + ("" + (v4 + ((v4 + v4) + (("hello" + v4) + "hello")))))) + ((v4 + v4) + "foo bar"))}
                _ = v162
            }
            _ = fi157
        }
        _ = fi153
    }
    fmt.Println("v1:", v1)
    fmt.Println("v152:", v152)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v152
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v79
    _ = v80
    _ = v81
    _ = v84
    _ = v87
}
