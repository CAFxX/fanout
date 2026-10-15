package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int8; F1 int8; F2 int }
type GenS13 struct { F0 uint8; F1 uint }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int8; F1 string }
type GenS16 struct { F0 uint32; F1 int32; F2 uint }
type GenS17 struct { F0 uint; F1 int; F2 bool }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 int32; F1 uint64; F2 int64 }
type GenS21 struct { F0 uint64 }
type GenS22 struct { F0 int; F1 int64 }
type GenS23 struct { F0 uint64; F1 int16 }
func genFn0(p0 int64, p1 uint, p2 int) uint32 {
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := GenS1{F0: (fi4 / ((fi4) | 1)), F1: ("foo bar" + ""), F2: 65535}
        var v6 int64 = p0
        switch "hello" {
        case "\n\t":
            for fi7 := 0; fi7 < 16; fi7++ {
                if (42 >= p1) {
                    switch fi4 {
                    case -709:
                        p2 -= 2147483647
                    case 127:
                        if (1e10 >= 1e10) {
                            v8 := []int{-9223372036854775808}
                            if true {
                                var v9 int8 = 127
                                var v10 float32 = 0.1
                                _ = v10
                                _ = v9
                            }
                            _ = v8
                        } else {
                            fmt.Println(fi7, fi4)
                            var v11 uint8 = 2
                            v12 := GenS1{F0: 732, F1: "\n\t", F2: 255}
                            _ = v11
                            _ = v12
                        }
                    case -477:
                        var v13 uint64 = 1
                        _ = v13
                    }
                    v14 := []int{int(p0), 166, fi7, (fi7 << ((p1) & 63))}
                    _ = v14
                }
                switch "\n\t" {
                case "a":
                    switch -783 {
                    case 32768:
                        fmt.Println(p2, v6, fi7)
                    }
                    switch fi4 {
                    case 253:
                        v15 := GenS2{F0: "a"}
                        _ = v15
                    case 2147483647:
                        v16 := []int{(^(fi7)), -187, (fi4 % ((p2) | 1))}
                        _ = v16
                    case -362:
                        var v17 string = ("" + ((("foo bar" + "\n\t") + "a") + "a"))
                        _ = v17
                    default:
                        var v18 uint16 = 547
                        _ = v18
                    }
                    if (p1 >= (p1 | 285)) {
                        v19 := []int{(^(fi7))}
                        var v20 string = "hello"
                        if (-128 > -1) {
                            v21 := GenS2{F0: v20}
                            var v22 uint64 = 56
                            fmt.Println(v6)
                            _ = v21
                            _ = v22
                        } else {
                            v23 := GenS1{F0: (p2 >> (((p1 + (^(p1)))) & 63)), F1: v20, F2: 32767}
                            _ = v23
                        }
                        _ = v19
                        _ = v20
                    } else {
                        var v24 int32 = 341
                        var v25 int8 = -91
                        v24 = (v24 & 28)
                        _ = v24
                        _ = v25
                    }
                }
                _ = fi7
            }
        case "foo bar":
            if (false || ((99 >= -27) || ((3.14159 + (1.5 - (0.1 / (float64(3.14159) + 1e10)))) <= 0.1))) {
                p1 = (p1 % ((42) | 1))
            }
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    switch "" {
    case "a":
        v26 := []int{}
        _ = v26
    case "\n\t":
        v27 := GenS3{F0: (^(p0)), F1: 32767}
        _ = v27
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 float32 = -2.25
        p0 = 9223372036854775807
        if false {
            v29 = ((v29 + 3.14159) + v29)
            switch 42 {
            case -2147483648:
                if true {
                    for fi30 := 0; fi30 < 12; fi30++ {
                        switch "a" {
                        case "hello":
                            var v31 uint8 = 255
                            _ = v31
                        case "":
                            p0 = p0
                        default:
                            switch "\n\t" {
                            case "\n\t":
                                if (32767 < 32767) {
                                    if (1e10 < 0.1) {
                                        var v32 uint64 = 0
                                        v33 := GenS4{F0: "", F1: 614}
                                        switch 1 {
                                        case 319:
                                            var v34 uint8 = 128
                                            _ = v34
                                        }
                                        _ = v32
                                        _ = v33
                                    } else {
                                        switch (^(fi30)) {
                                        case 127:
                                            if true {
                                                p1 -= p1
                                            }
                                        }
                                        var v35 uint64 = 555
                                        v29 = 3.14159
                                        var v36 float64 = 1e10
                                        _ = v35
                                        _ = v36
                                    }
                                    v37 := []int{fi30, 32768}
                                    var v38 int16 = -32768
                                    _ = v37
                                    _ = v38
                                }
                            case "hello":
                                v29 -= v29
                            case "":
                                if false {
                                    v39 := []int{}
                                    p0 = (p0 & 9223372036854775807)
                                    _ = v39
                                }
                            default:
                                v40 := []int{}
                                _ = v40
                            }
                        }
                        var v41 int = fi30
                        var v42 int16 = 1
                        _ = fi30
                        _ = v41
                        _ = v42
                    }
                    v29 += -2.25
                    var v43 int = fi28
                    _ = v43
                } else {
                    switch fi28 {
                    case 1:
                        fmt.Println(v29, p0, fi28)
                    case 0:
                        v29 -= v29
                    }
                    var v44 uint8 = 128
                    v45 := GenS5{F0: (1 <= 684), F1: 32767}
                    _ = v44
                    _ = v45
                }
            case 32767:
                var v46 int32 = 2147483647
                _ = v46
            case -1:
                for fi47 := 0; fi47 < 14; fi47++ {
                    v48 := []int{}
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        v48 = append(v48, fi28)
                        var v49 int = (fi47 + fi47)
                        _ = v49
                    case "a":
                        switch fi28 {
                        case 128:
                            var v50 string = "foo bar"
                            _ = v50
                        case -48:
                            p1 = 2
                        case 1:
                            fmt.Println(p0, v29)
                        }
                    case "\n\t":
                        if len(v48) >= 2 {
                            v48 = v48[1:len(v48)]
                        }
                    }
                    _ = fi47
                    _ = v48
                }
            default:
                if (42 == 0) {
                    for fi51 := 0; fi51 < 1; fi51++ {
                        fmt.Println(p2, p1, p1)
                        p1 = 18446744073709551615
                        if (!(!(127 >= 42))) {
                            v52 := []int{(p2 | (fi51 ^ fi51))}
                            fmt.Println(fi51, p2, p0)
                            _ = v52
                        }
                        if (float32((3.14159 / (-2.25 + (0.0 / ((-2.25 * (0.1 + float64(((v29 / v29) + (3.14159 / 1e10))))) * 0.0))))) < (3.14159 * (float32(3.14159) - (v29 / v29)))) {
                            if (true && true) {
                                v53 := GenS6{F0: 122, F1: 42}
                                if true {
                                    v54 := []int{}
                                    p2 = 42
                                    _ = v54
                                }
                                v55 := []int{(fi28 >> (((p1 / ((p1) | 1))) & 63)), (-(fi51))}
                                _ = v53
                                _ = v55
                            }
                        }
                        p0 = p0
                        if (true || (65535 < 1)) {
                            if (p2 != p2) {
                                if ((42 != 251) && true) {
                                    if true {
                                        v56 := []int{p2, (^(fi28)), fi28}
                                        _ = v56
                                    } else {
                                        var v57 uint16 = 32767
                                        _ = v57
                                    }
                                    p0 = (p0 * (-(p0)))
                                    switch "" {
                                    case "foo bar":
                                        p2 += (p2 / ((fi51) | 1))
                                    case "hello":
                                        if (false && false) {
                                            v58 := GenS7{F0: (p0 * int64(fi51)), F1: 338}
                                            var v59 int = fi51
                                            _ = v58
                                            _ = v59
                                        }
                                    }
                                }
                                fmt.Println(v29, p1, p2)
                            }
                            v60 := []int{32768, p2}
                            v60 = append(v60, (p2 / ((int(p1)) | 1)))
                            _ = len(v60)
                            if (32767 >= fi28) {
                                var v61 string = ""
                                v62 := GenS4{F0: ((v61 + ((v61 + (("hello" + v61) + (v61 + "foo bar"))) + "\n\t")) + "\n\t"), F1: 42}
                                var v63 uint64 = 595
                                switch fi51 {
                                case 197:
                                    p0 -= (p0 | p0)
                                default:
                                    v60 = append(v60, fi28)
                                }
                                switch (v61 + (v61 + (v61 + v61))) {
                                case "":
                                    v64 := GenS8{F0: -58, F1: v63}
                                    if ((p1 > 42) && true) {
                                        var v65 string = "\n\t"
                                        var v66 int8 = 5
                                        _ = v65
                                        _ = v66
                                    } else {
                                        var v67 uint32 = 973
                                        _ = v67
                                    }
                                    _ = v64
                                }
                                _ = v61
                                _ = v62
                                _ = v63
                            } else {
                                var v68 int = fi51
                                fmt.Println(fi51, fi28)
                                fmt.Println(fi28, p2, p2)
                                _ = v68
                            }
                            var v69 int8 = 55
                            _ = v60
                            _ = v69
                        } else {
                            switch "\n\t" {
                            case "hello":
                                p2 = (fi28 >> ((p1) & 63))
                            case "a":
                                if (!false) {
                                    fmt.Println(p2)
                                    fmt.Println(v29, p0, fi51)
                                    if false {
                                        var v70 string = ""
                                        _ = v70
                                    } else {
                                        fmt.Println(fi28, p1)
                                    }
                                } else {
                                    v71 := GenS9{F0: 65535, F1: 1, F2: 51}
                                    if (false && false) {
                                        v72 := []int{fi51, fi28, -557, 2147483647}
                                        v73 := GenS10{F0: 0}
                                        var v74 uint64 = 42
                                        _ = v72
                                        _ = v73
                                        _ = v74
                                    } else {
                                        switch (fi28 ^ fi51) {
                                        case 758:
                                            fmt.Println(p1)
                                        case 4294967295:
                                            switch "foo bar" {
                                            case "foo bar":
                                                if true {
                                                    fmt.Println(v29, fi28, fi51)
                                                    v75 := []int{(fi51 + (fi51 & 2147483647)), 128}
                                                    _ = v75
                                                } else {
                                                    v76 := []int{128, fi51}
                                                    var v77 int16 = -1
                                                    _ = v76
                                                    _ = v77
                                                }
                                            case "hello":
                                                var v78 int8 = -1
                                                _ = v78
                                            case "\n\t":
                                                if ((true && false) && false) {
                                                    if ("a" <= "a") {
                                                        var v79 float32 = v29
                                                        v80 := []int{}
                                                        var v81 float64 = float64(v79)
                                                        _ = v79
                                                        _ = v80
                                                        _ = v81
                                                    } else {
                                                        v82 := []int{(p2 / ((fi28) | 1))}
                                                        _ = v82
                                                    }
                                                    var v83 string = ("\n\t" + ("foo bar" + ""))
                                                    var v84 uint64 = 927
                                                    v29 += (v29 + ((v29 / (0.1 / 3.14159)) - -2.25))
                                                    _ = v83
                                                    _ = v84
                                                } else {
                                                    fmt.Println(v29, p2)
                                                    p2 += (p2 << ((p1) & 63))
                                                }
                                            }
                                            fmt.Println(v29)
                                        case 32768:
                                            p2 += (fi28 << (((p1 * p1)) & 63))
                                        }
                                        v85 := GenS11{F0: 127}
                                        _ = v85
                                    }
                                    var v86 float64 = 0.0
                                    _ = v71
                                    _ = v86
                                }
                            default:
                                fmt.Println(fi28, fi28, p1)
                            }
                            var v87 int32 = -1
                            p2 += fi51
                            _ = v87
                        }
                        switch "a" {
                        case "a":
                            var v88 uint8 = 1
                            _ = v88
                        case "hello":
                            v89 := GenS12{F0: 127, F1: -14, F2: fi51}
                            _ = v89
                        default:
                            v90 := GenS1{F0: fi51, F1: (("" + "hello") + ""), F2: 255}
                            _ = v90
                        }
                        if false {
                            var v91 uint64 = 18446744073709551615
                            _ = v91
                        } else {
                            var v92 uint8 = 133
                            var v93 float32 = ((v29 / 1.5) * v29)
                            var v94 uint8 = v92
                            _ = v92
                            _ = v93
                            _ = v94
                        }
                        if true {
                            v95 := []int{fi51}
                            _ = v95
                        }
                        _ = fi51
                    }
                    for fi96 := 0; fi96 < 18; fi96++ {
                        fmt.Println(p1)
                        var v97 uint = p1
                        switch fi28 {
                        case 223:
                            switch "a" {
                            case "\n\t":
                                var v98 int16 = 32767
                                _ = v98
                            }
                        case 32768:
                            v99 := []int{fi96, fi28, (p2 / ((int(p0)) | 1))}
                            _ = v99
                        case 128:
                            v100 := []int{fi96, p2, (fi28 / ((int(p0)) | 1))}
                            fmt.Println(fi96, p1)
                            var v101 float64 = float64((v29 - v29))
                            _ = v100
                            _ = v101
                        }
                        _ = fi96
                        _ = v97
                    }
                } else {
                    var v102 bool = (718 == 18446744073709551615)
                    _ = v102
                }
            }
        } else {
            if true {
                var v103 uint16 = 255
                _ = v103
            }
        }
        _ = fi28
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    return 693
}
func genFn1(p0 uint, p1 uint64) int16 {
    p1 = 2
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 343
}

func main() {
    var v1 int64 = 0
    v2 := make(map[string]bool)
    var v3 int64 = v1
    if true {
        v1 += v3
        _ = len(v2)
        var v104 int = 4294967295
        _ = v104
    } else {
        var v105 int32 = -2147483648
        v106 := []int{127, 128}
        for fi107 := 0; fi107 < 16; fi107++ {
            var v108 int64 = (v3 % (((v3 * 42)) | 1))
            switch (^(fi107)) {
            case 42:
                v109 := GenS13{F0: 0, F1: 65535}
                _ = v109
            case 127:
                for fi110 := 0; fi110 < 12; fi110++ {
                    if false {
                        v105 -= -875
                        v111 := GenS14{F0: 591}
                        _ = v111
                    } else {
                        if true {
                            var v112 uint64 = 1
                            switch "" {
                            case "foo bar":
                                delete(v2, "foo bar")
                            case "\n\t":
                                if len(v106) > 0 {
                                    v106[0] = -55
                                }
                            case "hello":
                                delete(v2, "hello")
                            default:
                                if len(v106) >= 2 {
                                    v106 = v106[1:len(v106)]
                                }
                            }
                            _ = v112
                        }
                    }
                    v113 := v2["foo bar"]
                    v3 = (-(v1))
                    _ = fi110
                    _ = v113
                }
            case 32767:
                var v114 string = (("foo bar" + ("a" + "a")) + "")
                _ = v114
            }
            for fi115 := 0; fi115 < 12; fi115++ {
                v1 += v3
                if (!(!true)) {
                    if false {
                        switch "" {
                        case "\n\t":
                            var v116 bool = (fi107 > (fi115 | 128))
                            _ = v116
                        case "":
                            v117 := GenS10{F0: 32767}
                            _ = v117
                        case "a":
                            v105 = -1
                        }
                        v105 += -2147483648
                    } else {
                        var v118 uint = 1
                        v1 = 272
                        var v119 uint16 = 511
                        _ = v118
                        _ = v119
                    }
                    switch ("hello" + "hello") {
                    case "a":
                        v1 = v3
                    case "foo bar":
                        v1 -= -9223372036854775808
                    default:
                        var v120 uint = 730
                        _ = v120
                    }
                    var v121 int64 = (v108 << ((560) & 63))
                    _ = v121
                }
                _ = fi115
            }
            _ = fi107
            _ = v108
        }
        _ = v105
        _ = v106
    }
    for fi122 := 0; fi122 < 6; fi122++ {
        if (!(2147483648 >= 65535)) {
            if (1e10 < -2.25) {
                if true {
                    var v123 uint32 = 0
                    v124 := GenS1{F0: fi122, F1: "a", F2: 458}
                    _ = v123
                    _ = v124
                } else {
                    v3 = (v1 >> ((2) & 63))
                    if (127 < 181) {
                        var v125 int16 = -32768
                        v126 := GenS10{F0: 623}
                        _ = v125
                        _ = v126
                    }
                    switch 9223372036854775807 {
                    case 585:
                        v127 := GenS15{F0: 127, F1: "\n\t"}
                        _ = v127
                    case -268:
                        for fi128 := 0; fi128 < 16; fi128++ {
                            v3 = v3
                            if ((606 <= 255) || true) {
                                var v129 string = ("\n\t" + "foo bar")
                                var v130 int8 = -1
                                if (false || true) {
                                    v3 = (v3 + v1)
                                    switch ((v129 + ((("hello" + v129) + "foo bar") + (((v129 + "a") + (v129 + ("" + v129))) + "foo bar"))) + "a") {
                                    case "":
                                        switch fi128 {
                                        case 9223372036854775807:
                                            v131 := GenS16{F0: 892, F1: 2147483647, F2: 2}
                                            _ = v131
                                        case 128:
                                            switch fi128 {
                                            case -794:
                                                fmt.Println(v1, fi128, fi122)
                                            case -9223372036854775808:
                                                v1 = v3
                                            case -1:
                                                v129 = ("" + v129)
                                                v3 += (v1 + (v3 << ((18446744073709551615) & 63)))
                                                var v132 int16 = -32768
                                                _ = v132
                                            default:
                                                v1 = -68
                                                v3 += v1
                                            }
                                        }
                                    case "\n\t":
                                        if true {
                                            v129 = (v129 + "hello")
                                            switch v129 {
                                            case "foo bar":
                                                var v133 int16 = 1
                                                _ = v133
                                            case "":
                                                v129 = ""
                                            case "a":
                                                v1 = v3
                                            default:
                                                var v134 string = v129
                                                _ = v134
                                            }
                                            v129 = (v129 + v129)
                                        }
                                    case "hello":
                                        v129 = v129
                                    }
                                } else {
                                    var v135 uint64 = 42
                                    _ = v135
                                }
                                var v136 int8 = (v130 << ((42) & 63))
                                _ = v129
                                _ = v130
                                _ = v136
                            } else {
                                var v137 uint16 = 30
                                if false {
                                    if (v137 <= v137) {
                                        v138 := GenS6{F0: -128, F1: 0}
                                        if false {
                                            fmt.Println(v1, fi122, fi128)
                                            v137 -= (v137 - v137)
                                            delete(v2, "\n\t")
                                        }
                                        _ = v138
                                    } else {
                                        delete(v2, "foo bar")
                                        v2[(("a" + ("\n\t" + (("" + ("foo bar" + (("a" + ("a" + "\n\t")) + "a"))) + "hello"))) + "hello")] = (true || false)
                                        v139 := GenS3{F0: v3, F1: 952}
                                        if ((!false) && (255 > 255)) {
                                            v137 = v137
                                            v137 = 255
                                        } else {
                                            v2["a"] = (-2.25 > ((0.0 * (1.5 * 1e10)) * 1e10))
                                        }
                                        _ = v139
                                    }
                                }
                                var v140 uint = 18446744073709551615
                                _ = v137
                                _ = v140
                            }
                            var v141 uint = 503
                            _ = fi128
                            _ = v141
                        }
                    case 1:
                        var v142 float64 = float64(((1e10 - 0.1) - 3.14159))
                        _ = v142
                    default:
                        for fi143 := 0; fi143 < 7; fi143++ {
                            if (2 < 0) {
                                v144 := []int{(-(fi143)), (fi143 % ((32768) | 1)), fi143, (fi143 / (((fi122 ^ (fi143 + fi122))) | 1))}
                                _ = v144
                            } else {
                                fmt.Println(v3)
                            }
                            v1 = (v1 & v3)
                            if (true || true) {
                                if (2 >= 42) {
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v145 int64 = v1
                                        _ = v145
                                    case "":
                                        v146 := []int{(fi122 ^ (^(fi143))), (fi122 << ((436) & 63)), (fi143 + fi122), (fi122 / (((fi143 - fi122)) | 1))}
                                        _ = v146
                                    case "\n\t":
                                        v3 = (v1 % (((v3 << ((833) & 63))) | 1))
                                    }
                                    var v147 uint32 = 877
                                    var v148 uint64 = uint64(v147)
                                    _ = v147
                                    _ = v148
                                } else {
                                    var v149 int64 = v1
                                    v150 := []int{1, (-(fi122)), (fi122 % (((^(fi143))) | 1)), fi122}
                                    _ = v149
                                    _ = v150
                                }
                            } else {
                                var v151 string = "\n\t"
                                var v152 bool = (18446744073709551615 > 0)
                                var v153 uint32 = 2147483647
                                _ = v151
                                _ = v152
                                _ = v153
                            }
                            _ = fi143
                        }
                    }
                }
            }
            var v154 float32 = (0.1 - 0.0)
            v155 := []int{fi122, 521, (-(fi122))}
            _ = v154
            _ = v155
        } else {
            var v156 int16 = -32768
            v3 = (v1 / (((v3 ^ -9223372036854775808)) | 1))
            _ = v156
        }
        v157 := GenS8{F0: 1, F1: 42}
        _ = fi122
        _ = v157
    }
    v3 += v1
    fmt.Println(v1)
    fmt.Println(v1, v1)
    if ((2 < 982) || ((0.1 + 1.5) == -2.25)) {
        if (2 <= 255) {
            if (true || true) {
                v1 -= (v1 ^ (v1 >> ((4294967295) & 63)))
                for fi158 := 0; fi158 < 8; fi158++ {
                    if (4294967295 <= 731) {
                        fmt.Println(fi158)
                        v159 := GenS9{F0: 42, F1: 980, F2: genFn0(v1, 0, 32768)}
                        v160 := GenS17{F0: 2, F1: (fi158 % (((fi158 * 0)) | 1)), F2: (true || ((fi158 * (fi158 + -767)) != fi158))}
                        _ = v159
                        _ = v160
                    }
                    switch "" {
                    case "a":
                        _ = len(v2)
                    case "foo bar":
                        for fi161 := 0; fi161 < 4; fi161++ {
                            if (false || true) {
                                v162 := []int{(fi158 | (fi158 & (-(fi158)))), (fi161 >> ((824) & 63)), fi158, 42}
                                v3 -= 1
                                v3 -= -594
                                _ = v162
                            }
                            fmt.Println(v1, fi161)
                            if (!true) {
                                var v163 int8 = -1
                                if true {
                                    switch fi158 {
                                    case 9223372036854775807:
                                        if (!false) {
                                            fmt.Println(v1, v163)
                                            switch (("foo bar" + "hello") + "foo bar") {
                                            case "foo bar":
                                                v164 := GenS18{F0: genFn1(219, 42)}
                                                _ = v164
                                            }
                                        } else {
                                            v2[""] = true
                                            _ = len(v2)
                                        }
                                    }
                                }
                                _ = v163
                            } else {
                                if true {
                                    v2["foo bar"] = (-2.25 < (0.0 * ((0.0 / ((float32(0.0) - 3.14159) - float32(0.1))) + (0.1 - ((((((float32(0.0) * (1e10 * (1e10 / ((1e10 + 0.0) * 3.14159)))) * float32(((0.1 / 1.5) / (1e10 / 1.0)))) * (0.0 / float32((3.14159 - (-2.25 * 1.5))))) / 0.1) + 0.1) / 0.1)))))
                                } else {
                                    var v165 uint8 = 2
                                    if (true && (v1 >= v3)) {
                                        if true {
                                            delete(v2, "\n\t")
                                            switch -9223372036854775808 {
                                            case 32768:
                                                v166 := GenS19{F0: -1}
                                                _ = v166
                                            case 24:
                                                if true {
                                                    v167 := []int{42, (fi158 % ((fi161) | 1))}
                                                    if (32768 < fi158) {
                                                        if false {
                                                            if len(v167) > 0 {
                                                                v167[0] = (-(fi158))
                                                            }
                                                            switch "foo bar" {
                                                            case "hello":
                                                                if false {
                                                                    v168 := GenS16{F0: 920, F1: 0, F2: 882}
                                                                    v3 = -1
                                                                    if false {
                                                                        v165 = (^(v165))
                                                                        v165 = v165
                                                                        if (!true) {
                                                                            if false {
                                                                                var v169 int32 = 42
                                                                                v170 := v2["a"]
                                                                                var v171 bool = (!v170)
                                                                                _ = v169
                                                                                _ = v170
                                                                                _ = v171
                                                                            } else {
                                                                                var v172 float32 = 1.5
                                                                                _ = v172
                                                                            }
                                                                            switch ("a" + "") {
                                                                            case "":
                                                                                v1 = (v1 % ((v3) | 1))
                                                                            case "foo bar":
                                                                                var v173 string = "\n\t"
                                                                                fmt.Println(v1, v173)
                                                                                _ = v173
                                                                            case "a":
                                                                                var v174 string = ("hello" + ("foo bar" + ("" + "hello")))
                                                                                _ = v174
                                                                            }
                                                                        } else {
                                                                            v3 = v1
                                                                            var v175 uint16 = 220
                                                                            var v176 float32 = ((0.0 / -2.25) / 1e10)
                                                                            _ = v175
                                                                            _ = v176
                                                                        }
                                                                    }
                                                                    if (!false) {
                                                                        v177 := GenS20{F0: -594, F1: 18446744073709551615, F2: (v3 - int64(fi161))}
                                                                        _ = v177
                                                                    }
                                                                    var v178 int64 = (v1 | (v1 * v1))
                                                                    v179 := GenS8{F0: 1, F1: 946}
                                                                    var v180 int32 = 0
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v181 uint64 = 1
                                                                        _ = v181
                                                                    }
                                                                    if true {
                                                                        if (((false || (!false)) || true) || (v1 <= 1)) {
                                                                            v165 = v165
                                                                            v178 += (v3 * (^(v178)))
                                                                            var v182 int16 = 201
                                                                            _ = v182
                                                                        }
                                                                        if len(v167) > 0 {
                                                                            v167[0] = fi158
                                                                        }
                                                                    } else {
                                                                        v183 := GenS21{F0: uint64(v165)}
                                                                        _ = v183
                                                                    }
                                                                    var v184 int16 = -1
                                                                    _ = v168
                                                                    _ = v178
                                                                    _ = v179
                                                                    _ = v180
                                                                    _ = v184
                                                                }
                                                            }
                                                        }
                                                        if (1.5 != ((0.1 - -2.25) - 1.5)) {
                                                            if ((("hello" <= "\n\t") || (!(629 != 1))) || false) {
                                                                v185 := GenS9{F0: 255, F1: -2147483648, F2: 42}
                                                                v165 = (v165 + (v165 & v165))
                                                                if (float64((((-2.25 + 1e10) - float32(1.5)) * -2.25)) < 1.5) {
                                                                    var v186 int64 = v1
                                                                    var v187 int = (fi161 ^ (^(fi161)))
                                                                    _ = v186
                                                                    _ = v187
                                                                } else {
                                                                    var v188 int8 = 1
                                                                    _ = v188
                                                                }
                                                                _ = v185
                                                            }
                                                        }
                                                        var v189 uint = 255
                                                        _ = v189
                                                    } else {
                                                        if false {
                                                            v3 = (^(v3))
                                                            if len(v167) > 0 {
                                                                v167[0] = fi158
                                                            }
                                                            v190 := GenS16{F0: 4294967295, F1: -2147483648, F2: 4294967295}
                                                            _ = v190
                                                        }
                                                    }
                                                    _ = v167
                                                } else {
                                                    delete(v2, "foo bar")
                                                    v3 = -1
                                                    v191 := []int{(fi158 << ((18446744073709551615) & 63)), 2147483647}
                                                    _ = v191
                                                }
                                            }
                                            v192 := GenS10{F0: 255}
                                            _ = v192
                                        }
                                        var v193 int = (^(fi158))
                                        _ = v193
                                    } else {
                                        switch fi161 {
                                        case 1:
                                            var v194 uint8 = (v165 >> ((1) & 63))
                                            _ = v194
                                        case -1:
                                            switch "a" {
                                            case "\n\t":
                                                switch (fi158 ^ 32767) {
                                                case 4294967295:
                                                    v1 += v1
                                                default:
                                                    var v195 int32 = 154
                                                    _ = v195
                                                }
                                            case "a":
                                                v196 := v2["a"]
                                                _ = v196
                                            case "hello":
                                                v3 += (v1 * -374)
                                            default:
                                                var v197 float32 = 0.0
                                                _ = v197
                                            }
                                        default:
                                            if (-273 > v1) {
                                                if true {
                                                    v165 = 163
                                                    if ((fi161 | int(v3)) <= fi161) {
                                                        var v198 uint8 = (^(v165))
                                                        _ = v198
                                                    } else {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            var v199 uint8 = 127
                                                            _ = v199
                                                        case "":
                                                            v200 := GenS11{F0: 0}
                                                            _ = v200
                                                        }
                                                        v201 := GenS20{F0: 1, F1: 1, F2: (v3 << ((696) & 63))}
                                                        _ = v201
                                                    }
                                                    v1 += (v3 - v3)
                                                    _ = len(v2)
                                                    v165 += v165
                                                } else {
                                                    switch fi158 {
                                                    case -2147483648:
                                                        var v202 int64 = v3
                                                        _ = v202
                                                    case 2147483647:
                                                        if false {
                                                            var v203 int16 = -1
                                                            switch "hello" {
                                                            case "a":
                                                                var v204 uint64 = 251
                                                                _ = v204
                                                            case "hello":
                                                                var v205 int16 = -93
                                                                _ = v205
                                                            default:
                                                                v206 := GenS13{F0: (v165 / (((v165 / ((21) | 1))) | 1)), F1: 97}
                                                                _ = v206
                                                            }
                                                            _ = v203
                                                        }
                                                    default:
                                                        v207 := []int{2147483647, 4294967295, (fi161 | (fi158 + fi161)), (fi158 & (^(fi161)))}
                                                        _ = v207
                                                    }
                                                    delete(v2, ("a" + "\n\t"))
                                                    v3 = v3
                                                }
                                                switch "hello" {
                                                case "":
                                                    v165 = v165
                                                case "a":
                                                    v165 = (v165 & (v165 & (v165 >> ((18446744073709551615) & 63))))
                                                case "\n\t":
                                                    v208 := []int{32, fi161, (fi161 | (fi161 / (((fi161 - 128)) | 1))), fi161}
                                                    _ = v208
                                                }
                                                v165 -= (^(v165))
                                                fmt.Println(fi158, fi158, v1)
                                            } else {
                                                fmt.Println(fi161, fi161)
                                                v1 += v1
                                            }
                                        }
                                        switch 646 {
                                        case 32768:
                                            var v209 int16 = 1
                                            _ = v209
                                        default:
                                            v165 -= 2
                                        }
                                        v210 := []int{(-(fi161)), fi158, fi158, fi161}
                                        var v211 int32 = 0
                                        if len(v210) > 0 {
                                            v210[0] = (fi161 | (fi158 - (-(fi158))))
                                        }
                                        _ = v210
                                        _ = v211
                                    }
                                    v1 -= v3
                                    _ = v165
                                }
                            }
                            _ = fi161
                        }
                    case "":
                        if (2147483648 > 257) {
                            var v212 int32 = 1
                            fmt.Println(v1, v1, fi158)
                            _ = v212
                        } else {
                            var v213 int32 = 641
                            v214 := []int{fi158}
                            _ = v213
                            _ = v214
                        }
                    }
                    _ = fi158
                }
            } else {
                fmt.Println(v3)
            }
            if false {
                v3 += v1
            }
            v3 = (v1 ^ (v3 * -9223372036854775808))
        }
        v215 := GenS22{F0: 2147483647, F1: v3}
        for fi216 := 0; fi216 < 12; fi216++ {
            if true {
                v217 := GenS17{F0: 18446744073709551615, F1: (-(fi216)), F2: false}
                v1 -= (^(v1))
                v218 := GenS23{F0: 4294967295, F1: -32768}
                _ = v217
                _ = v218
            }
            v219 := v2[("\n\t" + "foo bar")]
            _ = fi216
            _ = v219
        }
        _ = v215
    } else {
        var v220 int = 128
        var v221 int32 = -44
        _ = v220
        _ = v221
    }
    var v222 string = ("" + ("\n\t" + "hello"))
    v223 := GenS9{F0: 1, F1: 42, F2: 904}
    v2[v222] = (0 == 42)
    v224 := GenS15{F0: 37, F1: ("foo bar" + v222)}
    v3 = (v3 + 42)
    for fi225 := 0; fi225 < 20; fi225++ {
        switch fi225 {
        case 42:
            v3 = (v1 + (v1 + (v3 * -9223372036854775808)))
        case 1:
            v1 += (^(v1))
        case 9223372036854775807:
            if false {
                fmt.Println(v3, v3, v3)
                v226 := GenS4{F0: v222, F1: 175}
                v227 := v2[""]
                _ = v226
                _ = v227
            }
        default:
            v1 = 664
        }
        fmt.Println(v222)
        v228 := []int{(fi225 - (-(fi225))), fi225, -2147483648}
        _ = fi225
        _ = v228
    }
    v229 := GenS11{F0: 0}
    switch ("\n\t" + v222) {
    case "foo bar":
        fmt.Println(v1)
    case "\n\t":
        v1 = (v1 ^ 333)
    default:
        var v230 int8 = -59
        _ = v230
    }
    var v231 uint16 = 109
    for fi232 := 0; fi232 < 14; fi232++ {
        fmt.Println(v231, v222, v3)
        _ = fi232
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v222:", v222)
    fmt.Println("v231:", v231)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v222
    _ = v223
    _ = v224
    _ = v229
    _ = v231
    _ = v3
}
