package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 string; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32; F1 uint }
type GenS11 struct { F0 uint64 }
func genFn0(p0 int16, p1 int16, p2 uint32) uint16 {
    var v7 int64 = int64(p2)
    v8 := GenS1{F0: -1, F1: -128}
    p2 = (p2 - p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 255
}
func genFn1() int16 {
    for fi9 := 0; fi9 < 15; fi9++ {
        if true {
            if ("a" != "hello") {
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 int = fi9
                    v11 = (v11 - fi9)
                    _ = fi10
                    _ = v11
                }
            } else {
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 int = (fi9 * 9223372036854775807)
                    var v14 int32 = -1
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                v15 := GenS1{F0: fi9, F1: 0}
                _ = v15
            }
            var v16 float32 = ((1.5 + 0.1) + 3.14159)
            v17 := GenS2{F0: "", F1: 2, F2: 0}
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 uint16 = 2
                fmt.Println(fi18, fi9)
                var v20 uint = 4294967295
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        var v21 float32 = -2.25
        v22 := []int{fi9, (fi9 ^ fi9), fi9, fi9}
        _ = fi9
        _ = v21
        _ = v22
    }
    var v23 int64 = 0
    switch "hello" {
    case "":
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 int8 = -15
            _ = fi24
            _ = v25
        }
    case "a":
        for fi26 := 0; fi26 < 16; fi26++ {
            v23 = (v23 * 9223372036854775807)
            for fi27 := 0; fi27 < 8; fi27++ {
                if (fi26 >= (fi27 << ((255) & 63))) {
                    switch "hello" {
                    case "foo bar":
                        v28 := []int{(fi26 | 128)}
                        _ = v28
                    case "\n\t":
                        var v29 int16 = -1
                        _ = v29
                    case "hello":
                        v30 := []int{}
                        _ = v30
                    }
                    var v31 uint16 = 550
                    v32 := []int{}
                    _ = v31
                    _ = v32
                } else {
                    fmt.Println(fi26)
                    v33 := GenS3{F0: 0}
                    if true {
                        v23 -= (v23 + v23)
                        v23 = v23
                        v34 := []int{(fi26 | 0), fi27, (-(fi26)), fi26}
                        v23 = v23
                        if len(v34) > 0 {
                            v34[0] = fi26
                        }
                        switch fi26 {
                        case 879:
                            if false {
                                var v35 bool = ((0.1 == 1.5) && (true || (!(false && ("foo bar" == "a")))))
                                _ = v35
                            }
                        case -1:
                            v23 = (^(v23))
                        case 732:
                            var v36 uint = 1
                            _ = v36
                        }
                        var v37 int = fi26
                        _ = len(v34)
                        v38 := GenS2{F0: "\n\t", F1: 18446744073709551615, F2: 1}
                        var v39 uint64 = 2
                        if (((v39 ^ v39) <= 961) && ((false && (2 <= 2147483647)) || false)) {
                            var v40 string = "a"
                            fmt.Println(v37, v40)
                            if len(v34) > 0 {
                                v34[0] = (fi27 / ((fi27) | 1))
                            }
                            var v41 int8 = 127
                            _ = v40
                            _ = v41
                        } else {
                            v39 -= (v39 % ((18446744073709551615) | 1))
                        }
                        _ = v34
                        _ = v37
                        _ = v38
                        _ = v39
                    } else {
                        fmt.Println(fi26, fi27, fi26)
                    }
                    _ = v33
                }
                if (!((fi27 < fi27) || ((fi26 <= -1) && true))) {
                    var v42 float64 = (float64(3.14159) / (float64((3.14159 + 1.5)) * ((float64((0.1 + 3.14159)) + 0.1) + (0.1 / 1.0))))
                    _ = v42
                }
                var v43 int = fi26
                if (true && (("\n\t" + "hello") != "")) {
                    v44 := GenS1{F0: -9223372036854775808, F1: 1}
                    v43 = fi27
                    _ = v44
                }
                _ = fi27
                _ = v43
            }
            if (fi26 < fi26) {
                v23 = v23
                v45 := GenS2{F0: "hello", F1: 42, F2: 1}
                _ = v45
            }
            _ = fi26
        }
    default:
        fmt.Println(v23)
    }
    _ = v23
    return 741
}

func main() {
    var v1 uint = 4294967295
    var v2 uint = v1
    var v3 float64 = 3.14159
    var v4 float64 = v3
    v5 := []int{9223372036854775807}
    var v6 int64 = 9223372036854775807
    if true {
        switch ("\n\t" + "\n\t") {
        case "hello":
            v5 = append(v5, 0)
        case "foo bar":
            var v46 uint16 = 2
            _ = v46
        case "":
            v6 = v6
        default:
            switch "a" {
            case "foo bar":
                fmt.Println(v1, v4, v1)
            case "\n\t":
                v47 := GenS2{F0: "hello", F1: 402, F2: 255}
                _ = v47
            }
        }
        if (false && (false || (2 < 18446744073709551615))) {
            for fi48 := 0; fi48 < 9; fi48++ {
                fmt.Println(v3)
                var v49 float32 = ((1e10 * 0.0) / 1.5)
                v49 += (0.0 / 1.0)
                _ = fi48
                _ = v49
            }
            var v50 float64 = v3
            _ = v50
        }
        v51 := GenS2{F0: "\n\t", F1: 18446744073709551615, F2: 1}
        _ = v51
    } else {
        if true {
            switch 4294967295 {
            case 9223372036854775807:
                var v52 int8 = 100
                _ = v52
            case 4294967295:
                fmt.Println(v2)
            default:
                var v53 uint64 = 42
                _ = v53
            }
        } else {
            v4 += float64(float32(3.14159))
            _ = len(v5)
        }
    }
    switch 4294967295 {
    case 2147483647:
        var v54 uint8 = 114
        _ = v54
    case -590:
        v5 = append(v5, 128)
    }
    if true {
        switch "" {
        case "a":
            for fi55 := 0; fi55 < 12; fi55++ {
                for fi56 := 0; fi56 < 1; fi56++ {
                    var v57 int64 = v6
                    v58 := GenS4{F0: -918, F1: 128}
                    _ = fi56
                    _ = v57
                    _ = v58
                }
                var v59 int64 = (v6 >> (((v2 * (v1 * 65535))) & 63))
                v60 := GenS1{F0: 32767, F1: 1}
                _ = fi55
                _ = v59
                _ = v60
            }
        }
        v5 = append(v5, int(v2))
        for fi61 := 0; fi61 < 19; fi61++ {
            var v62 float32 = 0.1
            fmt.Println(v4, v6, v6)
            var v63 float64 = (((v4 / ((v3 + v3) - (v4 * (float64(v62) / 1.0)))) - v4) * 0.0)
            _ = fi61
            _ = v62
            _ = v63
        }
    }
    switch "\n\t" {
    case "foo bar":
        for fi64 := 0; fi64 < 17; fi64++ {
            for fi65 := 0; fi65 < 5; fi65++ {
                if (true || true) {
                    var v66 int16 = -598
                    _ = v66
                }
                v67 := GenS1{F0: (^(fi64)), F1: 0}
                if ((((fi65 > (fi65 + fi64)) && true) || false) && ((v6 & -866) <= v6)) {
                    var v68 int32 = 1
                    switch -9223372036854775808 {
                    case -412:
                        if (true && (!((42 <= (v68 ^ (v68 ^ (^(v68))))) && false))) {
                            v68 = -2147483648
                            if (true || (true && (!false))) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            } else {
                                v6 += v6
                                v5 = append(v5, (fi65 & fi65))
                            }
                            v2 -= v2
                        }
                    case -738:
                        var v69 float64 = 1e10
                        _ = v69
                    case 128:
                        if (true || (1 >= 42)) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v70 := GenS3{F0: 190}
                            _ = v70
                        }
                    }
                    _ = v68
                }
                _ = fi65
                _ = v67
            }
            v1 -= 18446744073709551615
            _ = fi64
        }
    default:
        for fi71 := 0; fi71 < 5; fi71++ {
            for fi72 := 0; fi72 < 19; fi72++ {
                fmt.Println(v4, fi71, v3)
                if (true && (!true)) {
                    var v73 int16 = -979
                    v5 = append(v5, -1)
                    _ = v73
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(v4, v4)
                    var v74 int = (fi72 & fi72)
                    _ = v74
                }
                if (!false) {
                    switch 127 {
                    case -111:
                        v75 := GenS5{F0: 65535}
                        _ = v75
                    case 0:
                        v5 = append(v5, (fi71 + (fi71 << (((v2 - 4294967295)) & 63))))
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                if false {
                    v76 := GenS6{F0: 0}
                    switch (fi72 << ((v2) & 63)) {
                    case 661:
                        if false {
                            var v77 uint8 = 127
                            fmt.Println(fi72, v3, v3)
                            _ = v77
                        } else {
                            v4 -= 1e10
                            v78 := GenS6{F0: (v6 / (((v6 | v6)) | 1))}
                            v1 = (v2 + v1)
                            if ((1e10 != 3.14159) || (!(-1 != -1))) {
                                fmt.Println(v1)
                                var v79 string = ""
                                _ = v79
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v80 uint64 = 42
                                var v81 bool = true
                                _ = v80
                                _ = v81
                            }
                            if false {
                                fmt.Println(v1, v6, v4)
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(v1, v6)
                                    switch "\n\t" {
                                    case "hello":
                                        v82 := GenS2{F0: "foo bar", F1: 380, F2: 65535}
                                        _ = v82
                                    case "\n\t":
                                        if true {
                                            v3 = v3
                                            var v83 string = ("hello" + "hello")
                                            var v84 int = (fi72 + fi72)
                                            _ = v83
                                            _ = v84
                                        } else {
                                            if false {
                                                var v85 uint16 = 531
                                                _ = v85
                                            } else {
                                                var v86 uint32 = 657
                                                _ = v86
                                            }
                                        }
                                    default:
                                        var v87 int8 = 100
                                        _ = v87
                                    }
                                    if (1000 != 42) {
                                        if (!((-115 != 121) && (true && true))) {
                                            if true {
                                                if (!false) {
                                                    if true {
                                                        var v88 uint16 = 255
                                                        v1 += (v1 + (v2 / ((v2) | 1)))
                                                        _ = v88
                                                    }
                                                    v1 += (^(v2))
                                                }
                                            } else {
                                                v2 = v1
                                                var v89 int32 = 853
                                                if (!((^(v2)) == v1)) {
                                                    var v90 int64 = int64(v2)
                                                    _ = v90
                                                } else {
                                                    v91 := GenS4{F0: (fi71 ^ (fi71 * -1)), F1: 187}
                                                    _ = v91
                                                }
                                                _ = v89
                                            }
                                            if false {
                                                var v92 int8 = 127
                                                _ = v92
                                            } else {
                                                var v93 uint8 = 0
                                                var v94 uint16 = 255
                                                var v95 uint64 = 1
                                                var v96 int8 = 1
                                                var v97 string = ("hello" + "foo bar")
                                                _ = v93
                                                _ = v94
                                                _ = v95
                                                _ = v96
                                                _ = v97
                                            }
                                            if (42 < 720) {
                                                switch fi71 {
                                                case 647:
                                                    if len(v5) > 0 {
                                                        v5[0] = fi71
                                                    }
                                                case 42:
                                                    if true {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        v5 = append(v5, 415)
                                                    } else {
                                                        v2 += v2
                                                        v4 = ((((1e10 * (v4 * 1e10)) * (v3 + (-2.25 - 3.14159))) * v4) / (((v3 + v4) * 0.1) + v4))
                                                        fmt.Println(v3)
                                                    }
                                                case 358:
                                                    v2 = v1
                                                }
                                            } else {
                                                v3 = 1e10
                                            }
                                            v5 = append(v5, fi71)
                                            fmt.Println(fi71)
                                            switch fi72 {
                                            case -1:
                                                if true {
                                                    fmt.Println(fi72, fi71, v6)
                                                    v5 = append(v5, 128)
                                                } else {
                                                    var v98 uint8 = 1
                                                    switch -9223372036854775808 {
                                                    case 0:
                                                        if false {
                                                            if len(v5) > 0 {
                                                                v5[0] = fi71
                                                            }
                                                            if (v98 > v98) {
                                                                v98 += 255
                                                                v99 := GenS7{F0: 42, F1: 0}
                                                                _ = v99
                                                            } else {
                                                                v100 := GenS5{F0: (v1 % ((2) | 1))}
                                                                v2 = (v2 - (v2 % ((18446744073709551615) | 1)))
                                                                var v101 float32 = -2.25
                                                                _ = v100
                                                                _ = v101
                                                            }
                                                        }
                                                    case 32767:
                                                        if len(v5) > 0 {
                                                            v5[0] = (^(fi71))
                                                        }
                                                    case -9:
                                                        var v102 int16 = 1
                                                        _ = v102
                                                    default:
                                                        v5 = append(v5, (-(fi72)))
                                                    }
                                                    if (genFn0(genFn1(), 1000, 42) > 348) {
                                                        var v103 int = (fi71 ^ fi71)
                                                        v6 = 9223372036854775807
                                                        switch "a" {
                                                        case "hello":
                                                            v3 += (v3 + v4)
                                                        case "a":
                                                            var v104 string = "hello"
                                                            _ = v104
                                                        case "foo bar":
                                                            var v105 int8 = 100
                                                            _ = v105
                                                        }
                                                        _ = v103
                                                    } else {
                                                        fmt.Println(fi71, v6)
                                                        switch "hello" {
                                                        case "\n\t":
                                                            fmt.Println(v1)
                                                        }
                                                        if false {
                                                            v3 = 0.0
                                                            var v106 int16 = -32768
                                                            v4 = 1.5
                                                            _ = v106
                                                        } else {
                                                            if len(v5) > 0 {
                                                                v5[0] = (fi71 + fi71)
                                                            }
                                                            v107 := GenS8{F0: -238}
                                                            v1 += (v1 & (v2 ^ (v2 % ((v2) | 1))))
                                                            _ = v107
                                                        }
                                                    }
                                                    var v108 string = ("" + "foo bar")
                                                    var v109 uint8 = (v98 % ((128) | 1))
                                                    v110 := GenS4{F0: fi72, F1: (v109 + v98)}
                                                    _ = v108
                                                    _ = v109
                                                    _ = v110
                                                    _ = v98
                                                }
                                            case 0:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            default:
                                                var v111 uint8 = 74
                                                _ = v111
                                            }
                                        }
                                    } else {
                                        var v112 int16 = 1
                                        switch fi72 {
                                        case -9223372036854775808:
                                            var v113 float32 = 3.14159
                                            _ = v113
                                        case 598:
                                            v2 = v1
                                        }
                                        if (199 < 255) {
                                            var v114 string = ("foo bar" + ((("a" + ("a" + "hello")) + "a") + ("\n\t" + ("" + ("foo bar" + ("\n\t" + "\n\t"))))))
                                            var v115 int32 = -2147483648
                                            switch v114 {
                                            case "a":
                                                var v116 string = v114
                                                _ = v116
                                            }
                                            _ = v114
                                            _ = v115
                                        }
                                        _ = v112
                                    }
                                case "":
                                    v117 := GenS1{F0: fi71, F1: -1}
                                    _ = v117
                                case "hello":
                                    var v118 uint64 = 42
                                    _ = v118
                                default:
                                    v119 := GenS9{F0: 18446744073709551615}
                                    _ = v119
                                }
                                var v120 int8 = 42
                                _ = v120
                            }
                            _ = v78
                        }
                        if false {
                            fmt.Println(fi72)
                            _ = len(v5)
                        } else {
                            v1 = (^(v1))
                            var v121 uint = v2
                            var v122 uint32 = 513
                            v2 -= 18446744073709551615
                            var v123 uint = 65535
                            fmt.Println(fi72)
                            _ = v121
                            _ = v122
                            _ = v123
                        }
                        var v124 string = (("hello" + "foo bar") + "\n\t")
                        _ = v124
                    case -1:
                        var v125 uint32 = 1
                        _ = v125
                    case -690:
                        var v126 bool = false
                        _ = v126
                    }
                    v1 += 2
                    _ = v76
                } else {
                    if false {
                        v127 := GenS2{F0: "a", F1: 2, F2: 1}
                        if (true || (18446744073709551615 != 1)) {
                            var v128 int = fi72
                            switch ("hello" + "") {
                            case "\n\t":
                                fmt.Println(v2)
                            case "":
                                var v129 int64 = (v6 | v6)
                                _ = v129
                            }
                            _ = v128
                        }
                        _ = v127
                    } else {
                        v130 := GenS10{F0: 0, F1: (v2 ^ v1)}
                        v2 = v2
                        _ = v130
                    }
                }
                v131 := GenS11{F0: 119}
                _ = fi72
                _ = v131
            }
            var v132 int8 = 14
            _ = fi71
            _ = v132
        }
    }
    fmt.Println(v4, v3, v3)
    if true {
        var v133 uint8 = 168
        var v134 int = 0
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v133
        _ = v134
    } else {
        var v135 float32 = 0.1
        if len(v5) > 0 {
            v5[0] = 433
        }
        _ = v135
    }
    for fi136 := 0; fi136 < 19; fi136++ {
        switch fi136 {
        case -2147483648:
            v5 = append(v5, fi136)
        }
        _ = fi136
    }
    v2 += (v2 + 4294967295)
    v1 += v1
    var v137 string = "a"
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v1 = (^(v1))
    if true {
        v138 := GenS4{F0: -344, F1: 2}
        v139 := GenS9{F0: 2}
        for fi140 := 0; fi140 < 2; fi140++ {
            var v141 int16 = -60
            _ = len(v5)
            _ = fi140
            _ = v141
        }
        _ = v138
        _ = v139
    } else {
        var v142 uint8 = 104
        var v143 float32 = ((0.0 - 1e10) / 1.0)
        _ = v142
        _ = v143
    }
    var v144 float64 = (3.14159 + v3)
    switch (("\n\t" + v137) + v137) {
    case "a":
        for fi145 := 0; fi145 < 20; fi145++ {
            var v146 uint64 = 34
            for fi147 := 0; fi147 < 15; fi147++ {
                v148 := GenS8{F0: v6}
                _ = fi147
                _ = v148
            }
            _ = fi145
            _ = v146
        }
    case "\n\t":
        var v149 int8 = 42
        _ = v149
    default:
        var v150 uint = 835
        _ = v150
    }
    var v151 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v137:", v137)
    fmt.Println("v144:", v144)
    fmt.Println("v151:", v151)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v137
    _ = v144
    _ = v151
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
