package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint }
type GenS2 struct { F0 uint16; F1 int32; F2 int16 }
type GenS3 struct { F0 uint8; F1 int; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 uint16; F1 uint64; F2 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int; F1 int64; F2 string }
type GenS11 struct { F0 uint16; F1 uint64 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int64; F1 int8 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint; F1 uint; F2 string }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 bool = true
    var v3 uint8 = 128
    v4 := make(map[string]uint32)
    var v5 bool = v2
    v6 := make(map[string]uint32)
    var v7 int8 = -1
    var v8 int8 = v7
    switch -9223372036854775808 {
    case 42:
        v9 := GenS1{F0: v7, F1: 42, F2: 75}
        _ = v9
    case -294:
        v2 = (2147483647 <= 873)
    case 128:
        v10 := GenS1{F0: (v7 ^ (v8 % ((v7) | 1))), F1: -1, F2: 65535}
        _ = v10
    }
    fmt.Println(v3, v8, v8)
    v3 = 2
    var v11 uint32 = 864
    v8 -= v8
    v7 = v8
    switch "foo bar" {
    case "a":
        fmt.Println(v7, v2, v5)
    case "\n\t":
        if (!v2) {
            if v5 {
                var v12 float64 = ((float64(1e10) / 3.14159) + 0.1)
                for fi13 := 0; fi13 < 2; fi13++ {
                    fmt.Println(v2, v3, v8)
                    for fi14 := 0; fi14 < 1; fi14++ {
                        switch "foo bar" {
                        case "foo bar":
                            var v15 uint = 42
                            _ = v15
                        default:
                            v16 := GenS2{F0: 32767, F1: -441, F2: -720}
                            _ = v16
                        }
                        if v2 {
                            var v17 uint = 255
                            if (!false) {
                                var v18 uint64 = 618
                                v3 -= v3
                                v7 -= (v8 | v8)
                                var v19 float64 = ((1e10 + v12) * 0.1)
                                delete(v1, ("foo bar" + "a"))
                                _ = v18
                                _ = v19
                            } else {
                                if v5 {
                                    v20 := GenS2{F0: 560, F1: 1, F2: 32767}
                                    v21 := GenS3{F0: 255, F1: 2147483647, F2: v3}
                                    _ = v20
                                    _ = v21
                                }
                                var v22 float64 = 1.5
                                var v23 int = -2147483648
                                if false {
                                    if (270 == 0) {
                                        var v24 uint = (^(v17))
                                        _ = v24
                                    }
                                    v25 := GenS4{F0: fi14}
                                    v22 -= 0.1
                                    _ = v25
                                } else {
                                    fmt.Println(v23)
                                    if v5 {
                                        if (238 == (v3 & v3)) {
                                            var v26 uint64 = 42
                                            v27 := GenS3{F0: 2, F1: -1, F2: (v3 ^ 89)}
                                            _ = v26
                                            _ = v27
                                        }
                                        var v28 uint32 = v11
                                        var v29 uint = 42
                                        v22 += (v22 / 1.0)
                                        switch "a" {
                                        case "a":
                                            v29 = (v17 << (((v29 * 2)) & 63))
                                        case "":
                                            v30 := GenS5{F0: v7, F1: 934}
                                            _ = v30
                                        }
                                        _ = v28
                                        _ = v29
                                    }
                                    var v31 uint32 = (^(v11))
                                    var v32 uint8 = (v3 & (^(v3)))
                                    _ = v31
                                    _ = v32
                                }
                                fmt.Println(v23, v17, v2)
                                _ = v22
                                _ = v23
                            }
                            _ = v17
                        } else {
                            v11 = (v11 - (v11 + v11))
                            if v5 {
                                fmt.Println(v5, v11, v7)
                                v2 = v2
                                v33 := GenS6{F0: "foo bar", F1: 0}
                                v3 += v3
                                fmt.Println(fi13, v8, v5)
                                if ((("\n\t" + (((((("hello" + "hello") + ("a" + ("hello" + "\n\t"))) + (((("foo bar" + ("hello" + ("hello" + "hello"))) + ("" + "a")) + "") + ((("foo bar" + "foo bar") + "hello") + "foo bar"))) + "a") + ("foo bar" + "")) + "foo bar")) + "a") < "") {
                                    delete(v1, "hello")
                                    if (v5 || (!(!v5))) {
                                        fmt.Println(v8, v12)
                                        v34 := GenS5{F0: (v8 ^ (v8 & v7)), F1: 827}
                                        _ = v34
                                    } else {
                                        if v2 {
                                            v5 = (((!(true && true)) || v5) && ((v8 ^ -88) >= v8))
                                            var v35 string = ("foo bar" + "foo bar")
                                            _ = v35
                                        } else {
                                            fmt.Println(v2, v7, v2)
                                        }
                                    }
                                    v36 := v4["\n\t"]
                                    _ = v36
                                }
                                v37 := []int{fi13}
                                _ = v33
                                _ = v37
                            } else {
                                v2 = false
                            }
                            v38 := v1[""]
                            _ = v38
                        }
                        _ = fi14
                    }
                    _ = fi13
                }
                for fi39 := 0; fi39 < 7; fi39++ {
                    var v40 uint8 = 255
                    v41 := GenS4{F0: -1}
                    if v5 {
                        if (false && v2) {
                            switch (fi39 ^ (fi39 % (((fi39 - int(v40))) | 1))) {
                            case -2147483648:
                                if (-2147483648 > 1) {
                                    var v42 int8 = (v7 / ((v8) | 1))
                                    for fi43 := 0; fi43 < 1; fi43++ {
                                        fmt.Println(v3, v2, v11)
                                        _ = fi43
                                    }
                                    v44 := GenS7{F0: 2, F1: 0, F2: 295}
                                    _ = v42
                                    _ = v44
                                } else {
                                    for fi45 := 0; fi45 < 7; fi45++ {
                                        var v46 uint8 = (v3 >> ((721) & 63))
                                        _ = fi45
                                        _ = v46
                                    }
                                    var v47 int16 = 141
                                    _ = v47
                                }
                            default:
                                var v48 uint64 = 42
                                _ = v48
                            }
                        } else {
                            var v49 int16 = -218
                            _ = v49
                        }
                    }
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                _ = v12
            }
            delete(v6, ("" + ""))
            if (("\n\t" + "\n\t") < "hello") {
                for fi50 := 0; fi50 < 5; fi50++ {
                    v3 = (^(v3))
                    switch fi50 {
                    case 1:
                        var v51 bool = (true || v5)
                        _ = v51
                    default:
                        var v52 int = (fi50 - fi50)
                        _ = v52
                    }
                    fmt.Println(v7, v3)
                    _ = fi50
                }
                v53 := []int{}
                for fi54 := 0; fi54 < 14; fi54++ {
                    var v55 uint32 = v11
                    switch (fi54 >> ((4294967295) & 63)) {
                    case 42:
                        _ = len(v53)
                    case 9223372036854775807:
                        var v56 int32 = -2147483648
                        _ = v56
                    case -27:
                        switch fi54 {
                        case -940:
                            switch "hello" {
                            case "foo bar":
                                var v57 int32 = -178
                                _ = v57
                            case "":
                                var v58 uint16 = 394
                                _ = v58
                            case "\n\t":
                                var v59 int32 = int32(v55)
                                _ = v59
                            default:
                                fmt.Println(v2, v11, v3)
                            }
                        default:
                            for fi60 := 0; fi60 < 2; fi60++ {
                                if v2 {
                                    v61 := GenS8{F0: -1}
                                    v62 := GenS9{F0: ((1 != 42) || (2147483648 != 65535))}
                                    _ = v61
                                    _ = v62
                                }
                                var v63 int16 = 0
                                _ = fi60
                                _ = v63
                            }
                        }
                    }
                    _ = fi54
                    _ = v55
                }
                _ = v53
            } else {
                var v64 int64 = 131
                _ = v64
            }
        } else {
            for fi65 := 0; fi65 < 4; fi65++ {
                v66 := GenS3{F0: (v3 & (^(v3))), F1: (^(fi65)), F2: v3}
                if v2 {
                    v5 = (("foo bar" + "foo bar") > "a")
                    switch fi65 {
                    case 4294967295:
                        switch fi65 {
                        case 42:
                            if ((4294967295 <= 42) && (v8 <= v8)) {
                                for fi67 := 0; fi67 < 1; fi67++ {
                                    v8 = -65
                                    var v68 uint32 = v11
                                    _ = fi67
                                    _ = v68
                                }
                            }
                        case 32768:
                            v3 -= v3
                        default:
                            for fi69 := 0; fi69 < 8; fi69++ {
                                var v70 uint16 = 1
                                if v5 {
                                    v5 = v2
                                    delete(v6, "hello")
                                    switch fi65 {
                                    case 0:
                                        v71 := []int{fi65, fi65, fi69, 9223372036854775807}
                                        _ = v71
                                    case 9223372036854775807:
                                        if (true || true) {
                                            fmt.Println(v8)
                                            v11 = (v11 & (v11 << ((0) & 63)))
                                            var v72 string = (("\n\t" + "hello") + "a")
                                            _ = v72
                                        } else {
                                            switch ("hello" + "") {
                                            case "hello":
                                                fmt.Println(fi65, v7)
                                            case "\n\t":
                                                switch "" {
                                                case "\n\t":
                                                    fmt.Println(v7, v3, fi65)
                                                default:
                                                    var v73 uint32 = v11
                                                    _ = v73
                                                }
                                            case "foo bar":
                                                v70 = v70
                                            default:
                                                var v74 bool = (1e10 >= (0.0 * -2.25))
                                                _ = v74
                                            }
                                            var v75 uint8 = 255
                                            var v76 uint64 = 4294967295
                                            if v2 {
                                                var v77 int32 = 394
                                                fmt.Println(v70, v3, fi69)
                                                v78 := GenS10{F0: int(v70), F1: 1, F2: ("\n\t" + "")}
                                                v3 = uint8(v76)
                                                _ = v77
                                                _ = v78
                                            }
                                            switch fi69 {
                                            case 157:
                                                v79 := GenS2{F0: 32767, F1: -273, F2: 0}
                                                _ = v79
                                            case -9223372036854775808:
                                                v5 = (1e10 == (-2.25 - 1.5))
                                            }
                                            _ = v75
                                            _ = v76
                                        }
                                    case -2147483648:
                                        v80 := GenS9{F0: (!((v5 || v2) || v5))}
                                        fmt.Println(v70, v5, fi69)
                                        if ("\n\t" >= "\n\t") {
                                            var v81 uint = 18446744073709551615
                                            fmt.Println(v2)
                                            if ((("a" + "foo bar") + "hello") <= "hello") {
                                                fmt.Println(v8, v3)
                                                if v5 {
                                                    v3 -= (v3 % ((0) | 1))
                                                    if v2 {
                                                        var v82 float32 = (-2.25 - 1.5)
                                                        _ = v82
                                                    } else {
                                                        if false {
                                                            v83 := GenS6{F0: ("hello" + ("foo bar" + (("a" + "a") + "a"))), F1: 0}
                                                            var v84 string = ("foo bar" + "foo bar")
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                        switch fi69 {
                                                        case 42:
                                                            var v85 uint32 = (v11 % ((v11) | 1))
                                                            _ = v85
                                                        }
                                                    }
                                                    var v86 uint16 = 0
                                                    _ = v86
                                                }
                                            }
                                            _ = v81
                                        } else {
                                            var v87 uint32 = 2147483648
                                            switch (-(fi69)) {
                                            case 258:
                                                fmt.Println(v87)
                                            case 42:
                                                if v5 {
                                                    v88 := v1[("foo bar" + "hello")]
                                                    _ = v88
                                                } else {
                                                    v89 := GenS11{F0: v70, F1: 4294967295}
                                                    var v90 int32 = 42
                                                    _ = v89
                                                    _ = v90
                                                }
                                            case 438:
                                                v11 = v11
                                            default:
                                                v91 := []int{(-(fi65))}
                                                _ = v91
                                            }
                                            _ = v87
                                        }
                                        if ("" > "\n\t") {
                                            v92 := GenS12{F0: 42}
                                            _ = v92
                                        } else {
                                            var v93 float64 = 1e10
                                            if v2 {
                                                switch (fi69 >> ((255) & 63)) {
                                                case 127:
                                                    if (v2 && ((-9223372036854775808 < -946) && false)) {
                                                        v94 := []int{}
                                                        if len(v94) > 0 {
                                                            v94[0] = fi69
                                                        }
                                                        _ = v94
                                                    }
                                                case -492:
                                                    if (680 >= (fi65 << ((uint(v11)) & 63))) {
                                                        var v95 uint8 = 42
                                                        v96 := GenS9{F0: ((!v2) || (!(!v2)))}
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                case 32768:
                                                    v11 = v11
                                                default:
                                                    if v5 {
                                                        var v97 uint = 255
                                                        v98 := []int{}
                                                        if true {
                                                            v93 -= 0.0
                                                            if ((32767 == 1000) || v5) {
                                                                v8 -= v8
                                                                v3 = (v3 / ((v3) | 1))
                                                            } else {
                                                                if (v3 < (v3 - v3)) {
                                                                    if v5 {
                                                                        var v99 bool = v5
                                                                        var v100 bool = (-1 == -304)
                                                                        _ = v100
                                                                        _ = v99
                                                                    } else {
                                                                        var v101 uint = v97
                                                                        if len(v98) >= 2 {
                                                                            v98 = v98[1:len(v98)]
                                                                        }
                                                                        switch (-(fi65)) {
                                                                        case 42:
                                                                            var v102 float32 = 0.0
                                                                            _ = v102
                                                                        case -9223372036854775808:
                                                                            switch (fi65 % ((fi69) | 1)) {
                                                                            case 42:
                                                                                v2 = (!(!true))
                                                                            case 9223372036854775807:
                                                                                fmt.Println(v7, v97)
                                                                            default:
                                                                                if len(v98) >= 2 {
                                                                                    v98 = v98[1:len(v98)]
                                                                                }
                                                                            }
                                                                        case -954:
                                                                            var v103 uint16 = v70
                                                                            _ = v103
                                                                        default:
                                                                            v2 = ((1 >= 42) || true)
                                                                        }
                                                                        _ = v101
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(v70)
                                                        } else {
                                                            if true {
                                                                v104 := GenS13{F0: 1, F1: -128}
                                                                _ = v104
                                                            }
                                                        }
                                                        _ = v97
                                                        _ = v98
                                                    } else {
                                                        var v105 uint8 = v3
                                                        if (true && (!v5)) {
                                                            if v2 {
                                                                var v106 int16 = 42
                                                                var v107 int8 = int8(v11)
                                                                _ = v106
                                                                _ = v107
                                                            } else {
                                                                var v108 int32 = -1
                                                                _ = v108
                                                            }
                                                            switch (("\n\t" + "hello") + "foo bar") {
                                                            case "":
                                                                var v109 string = "a"
                                                                _ = v109
                                                            case "a":
                                                                var v110 int16 = 1
                                                                _ = v110
                                                            case "hello":
                                                                var v111 uint = 214
                                                                var v112 uint32 = (v11 * (^(v11)))
                                                                var v113 int64 = -9223372036854775808
                                                                _ = v111
                                                                _ = v112
                                                                _ = v113
                                                            }
                                                        } else {
                                                            v114 := v1["foo bar"]
                                                            var v115 float64 = 0.0
                                                            switch v114 {
                                                            case "\n\t":
                                                                var v116 bool = v5
                                                                _ = v116
                                                            case "a":
                                                                v11 = v11
                                                            }
                                                            _ = v114
                                                            _ = v115
                                                        }
                                                        if false {
                                                            var v117 int8 = (-(v7))
                                                            _ = v117
                                                        }
                                                        fmt.Println(v93, v93)
                                                        _ = v105
                                                    }
                                                }
                                                fmt.Println(fi65, v3)
                                                var v118 uint8 = (v3 - v3)
                                                _ = v118
                                            }
                                            _ = v93
                                        }
                                        v119 := GenS14{F0: (v8 | (v7 + (v7 / ((v8) | 1))))}
                                        v120 := GenS15{F0: 296, F1: 0, F2: "a"}
                                        v121 := GenS14{F0: -128}
                                        v11 -= (v11 + (v11 & v11))
                                        _ = v119
                                        _ = v120
                                        _ = v121
                                        _ = v80
                                    }
                                    v3 += v3
                                } else {
                                    v122 := []int{(fi65 ^ fi69), fi69, (fi65 - fi65), (^(fi69))}
                                    fmt.Println(fi65, v8, v70)
                                    var v123 uint32 = (v11 | v11)
                                    _ = v122
                                    _ = v123
                                }
                                if (815 < -337) {
                                    switch fi65 {
                                    case 128:
                                        var v124 float32 = 1.5
                                        _ = v124
                                    case -9223372036854775808:
                                        v3 -= v3
                                    default:
                                        v125 := GenS4{F0: (fi69 << ((663) & 63))}
                                        _ = v125
                                    }
                                    fmt.Println(v2, v70, v8)
                                    v70 -= v70
                                } else {
                                    v1["a"] = "a"
                                    fmt.Println(fi69)
                                    if v5 {
                                        v11 -= v11
                                        switch (fi69 * (fi69 ^ int(v70))) {
                                        case 42:
                                            var v126 int32 = 0
                                            _ = v126
                                        case 397:
                                            var v127 uint16 = v70
                                            _ = v127
                                        default:
                                            if (((!v5) && v5) || (v3 > (v3 - 0))) {
                                                v7 -= v7
                                            } else {
                                                v3 += 255
                                            }
                                        }
                                        v7 += -1
                                    }
                                }
                                _ = fi69
                                _ = v70
                            }
                        }
                    case -576:
                        v128 := GenS1{F0: v7, F1: 673, F2: 65535}
                        _ = v128
                    case 778:
                        fmt.Println(fi65)
                    default:
                        v5 = true
                    }
                    fmt.Println(v8)
                }
                if (1 >= 255) {
                    var v129 int32 = 381
                    v130 := GenS9{F0: (!true)}
                    var v131 string = "a"
                    _ = v129
                    _ = v130
                    _ = v131
                }
                _ = fi65
                _ = v66
            }
        }
    default:
        v132 := GenS4{F0: 32768}
        _ = v132
    }
    if (v2 || false) {
        if ((!(691 > 647)) && ("a" < "a")) {
            if (!v5) {
                v2 = v5
                v133 := GenS16{F0: v5}
                fmt.Println(v7)
                _ = v133
            }
            fmt.Println(v5, v7)
        }
        for fi134 := 0; fi134 < 18; fi134++ {
            var v135 uint = 4294967295
            v136 := v4[(("" + "a") + "")]
            _ = fi134
            _ = v135
            _ = v136
        }
        var v137 int32 = 1
        _ = v137
    } else {
        var v138 uint64 = 0
        fmt.Println(v138, v3, v8)
        fmt.Println(v8, v2)
        _ = v138
    }
    v11 -= 2147483647
    if (v11 < 65535) {
        v139 := v6[""]
        if (!false) {
            var v140 float32 = (0.1 + (3.14159 - -2.25))
            var v141 uint32 = v11
            fmt.Println(v2)
            _ = v140
            _ = v141
        } else {
            v4[""] = uint32(v3)
        }
        v2 = (317 == 840)
        _ = v139
    } else {
        if v2 {
            if false {
                v8 = v8
                switch "a" {
                case "hello":
                    for fi142 := 0; fi142 < 11; fi142++ {
                        switch "\n\t" {
                        case "hello":
                            v5 = (v5 && v5)
                        case "":
                            v5 = (-407 >= 912)
                        }
                        var v143 int32 = -868
                        fmt.Println(v5, v7, fi142)
                        _ = fi142
                        _ = v143
                    }
                case "\n\t":
                    for fi144 := 0; fi144 < 12; fi144++ {
                        for fi145 := 0; fi145 < 11; fi145++ {
                            delete(v6, "hello")
                            _ = fi145
                        }
                        v146 := GenS17{F0: 9223372036854775807}
                        _ = fi144
                        _ = v146
                    }
                }
                switch 128 {
                case 42:
                    v8 = (v8 * (^(v7)))
                default:
                    var v147 bool = false
                    _ = v147
                }
            }
            v11 -= (^(v11))
            var v148 bool = ((0.0 + (1e10 * 0.1)) <= 3.14159)
            _ = v148
        } else {
            v149 := []int{32768}
            _ = v149
        }
        if true {
            fmt.Println(v7, v8, v2)
        } else {
            v11 = 4294967295
            v5 = true
        }
    }
    v2 = (255 >= 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
