package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint16; F1 string; F2 int8 }
type GenS3 struct { F0 int; F1 uint; F2 uint32 }
type GenS4 struct { F0 int; F1 string; F2 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16 }

func main() {
    var v1 int8 = -1
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 uint = 678
    var v5 uint64 = 0
    v6 := GenS1{F0: 293, F1: 1, F2: 333}
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(v1, v1, v4)
        _ = fi7
    }
    fmt.Println(v5, v5, v5)
    for fi8 := 0; fi8 < 1; fi8++ {
        if (1.5 >= float32(float64(float32(float64(-2.25))))) {
            if true {
                if (1 <= 1000) {
                    if true {
                        fmt.Println(v1)
                    } else {
                        for fi9 := 0; fi9 < 11; fi9++ {
                            var v10 int8 = 127
                            delete(v3, (("\n\t" + "foo bar") + "hello"))
                            if (!false) {
                                fmt.Println(fi9, fi9)
                            } else {
                                var v11 int16 = 42
                                var v12 uint8 = 128
                                _ = v11
                                _ = v12
                            }
                            _ = fi9
                            _ = v10
                        }
                    }
                    var v13 uint64 = 186
                    if false {
                        var v14 uint8 = 85
                        _ = v14
                    }
                    _ = v13
                } else {
                    if (!true) {
                        var v15 string = ""
                        _ = len(v2)
                        v3[v15] = fi8
                        _ = v15
                    }
                    if true {
                        v16 := []int{(fi8 * (^(fi8))), (-(fi8))}
                        _ = v16
                    } else {
                        v5 += v5
                        var v17 float32 = ((3.14159 * 3.14159) * 0.0)
                        _ = v17
                    }
                    v18 := []int{(fi8 >> (((v4 - v4)) & 63))}
                    _ = v18
                }
                v19 := GenS2{F0: 420, F1: (("" + "\n\t") + "foo bar"), F2: v1}
                var v20 int16 = int16(v1)
                _ = v19
                _ = v20
            } else {
                v5 += 1
                v21 := v2["\n\t"]
                var v22 uint8 = 128
                _ = v21
                _ = v22
            }
            var v23 uint32 = 2
            v4 = v4
            _ = v23
        } else {
            v5 = (v5 / (((v5 * (v5 | v5))) | 1))
            for fi24 := 0; fi24 < 15; fi24++ {
                var v25 float32 = -2.25
                _ = fi24
                _ = v25
            }
            if ((false && (345 != 661)) && (!((1 > -32768) || (!false)))) {
                for fi26 := 0; fi26 < 11; fi26++ {
                    switch ("" + ("foo bar" + ("foo bar" + "foo bar"))) {
                    case "a":
                        if (true || false) {
                            v27 := []int{}
                            var v28 int8 = v1
                            if true {
                                _ = len(v3)
                            } else {
                                var v29 bool = ((42 < 9223372036854775807) && (!("" >= ("hello" + ("\n\t" + ("\n\t" + "\n\t"))))))
                                var v30 int16 = 1
                                switch (fi26 / ((-9223372036854775808) | 1)) {
                                case 4294967295:
                                    var v31 float32 = (0.0 * float32(float64(((0.1 - 3.14159) * 3.14159))))
                                    _ = v31
                                default:
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v30)
                                        if v29 {
                                            v4 = 576
                                            var v32 uint8 = 13
                                            fmt.Println(v5, v30)
                                            switch "\n\t" {
                                            case "\n\t":
                                                switch -9223372036854775808 {
                                                case 841:
                                                    if len(v27) > 0 {
                                                        v27[0] = (-(fi8))
                                                    }
                                                default:
                                                    v2["a"] = 9223372036854775807
                                                }
                                            case "hello":
                                                if (float32((-2.25 / (3.14159 * float64(0.1)))) > 1e10) {
                                                    v1 = v28
                                                    v33 := GenS3{F0: 42, F1: (^(v4)), F2: 0}
                                                    v34 := GenS2{F0: 1, F1: ("foo bar" + ("hello" + "")), F2: v28}
                                                    _ = v33
                                                    _ = v34
                                                } else {
                                                    if true {
                                                        var v35 uint32 = 1
                                                        _ = v35
                                                    } else {
                                                        if ((float64(-2.25) + (float64(((-2.25 + 0.1) + (3.14159 * 0.0))) / 0.1)) != ((-2.25 * 1.5) * -2.25)) {
                                                            var v36 uint64 = 42
                                                            var v37 int = 4294967295
                                                            if (!v29) {
                                                                fmt.Println(v36)
                                                            } else {
                                                                v28 -= (v28 - int8(v4))
                                                                delete(v3, (("" + "foo bar") + "\n\t"))
                                                                v38 := v2[""]
                                                                _ = v38
                                                            }
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        fmt.Println(v4, v5, v32)
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v39 bool = ((v30 / ((v30) | 1)) > v30)
                                                            _ = v39
                                                        case "":
                                                            v28 -= (v1 + (v28 << ((v4) & 63)))
                                                        }
                                                    }
                                                    v40 := GenS4{F0: fi26, F1: "", F2: 1000}
                                                    _ = v40
                                                }
                                            }
                                            _ = v32
                                        }
                                    case "":
                                        _ = len(v27)
                                    default:
                                        v41 := GenS3{F0: fi8, F1: (v4 - (v4 | uint(v5))), F2: 2147483648}
                                        _ = v41
                                    }
                                }
                                _ = v29
                                _ = v30
                            }
                            _ = v27
                            _ = v28
                        } else {
                            v5 = 18446744073709551615
                        }
                    case "":
                        var v42 int8 = (^(v1))
                        _ = v42
                    default:
                        v2[(("\n\t" + (("foo bar" + ("\n\t" + "a")) + "")) + "foo bar")] = fi8
                    }
                    v1 += -1
                    fmt.Println(v4, v5)
                    _ = fi26
                }
            }
        }
        if (1 <= 598) {
            var v43 int32 = 140
            v44 := GenS2{F0: 501, F1: "foo bar", F2: 127}
            v45 := GenS2{F0: 2, F1: "", F2: (v1 % ((42) | 1))}
            _ = v43
            _ = v44
            _ = v45
        }
        if false {
            var v46 float64 = (3.14159 - 0.0)
            v4 -= (v4 << ((801) & 63))
            _ = v46
        } else {
            for fi47 := 0; fi47 < 14; fi47++ {
                var v48 uint16 = 0
                if false {
                    var v49 uint8 = 2
                    switch "\n\t" {
                    case "":
                        v3["\n\t"] = fi8
                        var v50 bool = (0.1 <= ((((-2.25 / float64((1e10 * 0.1))) / (0.1 - (-2.25 / (0.0 + 1e10)))) + 3.14159) + (1e10 - 1e10)))
                        if v50 {
                            switch "a" {
                            case "foo bar":
                                _ = len(v3)
                            default:
                                _ = len(v2)
                            }
                            v4 -= v4
                            if (int64(fi8) < 9223372036854775807) {
                                var v51 bool = v50
                                switch fi8 {
                                case 4294967295:
                                    var v52 uint64 = v5
                                    _ = v52
                                case 128:
                                    v53 := []int{(fi47 << ((v4) & 63)), (fi8 ^ (fi8 % ((int(v48)) | 1))), 127, fi8}
                                    _ = v53
                                default:
                                    fmt.Println(v49, v4, v50)
                                }
                                _ = v51
                            }
                        } else {
                            switch fi47 {
                            case -1:
                                fmt.Println(fi8)
                            }
                            v54 := GenS3{F0: fi8, F1: (v4 ^ (v4 + v4)), F2: 65535}
                            var v55 bool = v50
                            v56 := GenS4{F0: fi8, F1: "foo bar", F2: 42}
                            fmt.Println(v50, v55)
                            if v50 {
                                v49 -= v49
                                v57 := []int{fi47, fi8}
                                switch ((("\n\t" + ("" + "a")) + "") + ("hello" + ("foo bar" + "hello"))) {
                                case "hello":
                                    v55 = (!(1.5 != -2.25))
                                case "a":
                                    v2["foo bar"] = fi47
                                    var v58 int32 = 2147483647
                                    _ = v58
                                case "\n\t":
                                    v55 = v55
                                default:
                                    if v50 {
                                        var v59 int = fi8
                                        switch "\n\t" {
                                        case "":
                                            var v60 string = ""
                                            _ = v60
                                        case "hello":
                                            v57 = append(v57, v59)
                                        case "a":
                                            v55 = v55
                                        default:
                                            fmt.Println(fi8, v4)
                                        }
                                        _ = v59
                                    } else {
                                        var v61 uint32 = uint32(v49)
                                        _ = v61
                                    }
                                }
                                _ = v57
                            }
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                        _ = v50
                    }
                    _ = v49
                }
                _ = fi47
                _ = v48
            }
        }
        _ = fi8
    }
    if true {
        if (true && (1e10 > 1.5)) {
            if false {
                var v62 bool = (42 != 1)
                fmt.Println(v62, v62, v4)
                _ = v62
            } else {
                var v63 uint = 255
                _ = v63
            }
            var v64 uint8 = 128
            _ = v64
        }
        for fi65 := 0; fi65 < 11; fi65++ {
            switch fi65 {
            case 42:
                var v66 bool = (-454 == -566)
                _ = v66
            case -1:
                for fi67 := 0; fi67 < 7; fi67++ {
                    switch fi65 {
                    case -277:
                        switch 0 {
                        case 9223372036854775807:
                            var v68 string = "\n\t"
                            _ = v68
                        case 42:
                            if true {
                                v2["hello"] = fi65
                                v1 = -54
                            }
                        }
                    }
                    v69 := []int{fi67, 10, fi65, fi65}
                    if (!false) {
                        v5 -= v5
                    }
                    _ = fi67
                    _ = v69
                }
            case 0:
                if false {
                    for fi70 := 0; fi70 < 9; fi70++ {
                        switch "foo bar" {
                        case "\n\t":
                            var v71 uint32 = 1
                            _ = v71
                        case "hello":
                            v4 = (v4 + (v4 >> ((992) & 63)))
                        default:
                            switch 42 {
                            case 32768:
                                var v72 bool = (127 == 127)
                                _ = v72
                            default:
                                if false {
                                    switch (fi65 << (((v4 >> ((42) & 63))) & 63)) {
                                    case -541:
                                        v4 = v4
                                    }
                                    v1 += 127
                                    switch fi70 {
                                    case -2147483648:
                                        switch (fi70 >> (((v4 + (^(v4)))) & 63)) {
                                        case 42:
                                            var v73 string = ("\n\t" + "\n\t")
                                            _ = v73
                                        case 9223372036854775807:
                                            if false {
                                                var v74 uint64 = (v5 - v5)
                                                _ = v74
                                            }
                                        }
                                    case 1:
                                        var v75 int = (^(fi70))
                                        fmt.Println(v75, v5, v5)
                                        _ = v75
                                    case 42:
                                        var v76 int32 = 934
                                        _ = v76
                                    default:
                                        v1 = v1
                                    }
                                }
                            }
                        }
                        _ = fi70
                    }
                } else {
                    for fi77 := 0; fi77 < 2; fi77++ {
                        if (true && (false || (65535 != 0))) {
                            if (!(1 >= 1)) {
                                _ = len(v3)
                                var v78 float64 = float64(0.1)
                                _ = len(v2)
                                if ((!true) && false) {
                                    v5 += 868
                                    var v79 uint8 = 237
                                    fmt.Println(v1)
                                    _ = v79
                                }
                                var v80 float32 = 3.14159
                                var v81 uint = (v4 & (^(v4)))
                                if ((v78 != v78) && (1.5 > v78)) {
                                    fmt.Println(fi65, v1, v1)
                                    var v82 uint8 = 127
                                    _ = v82
                                }
                                _ = v78
                                _ = v80
                                _ = v81
                            } else {
                                if (1 == (^(v4))) {
                                    if ((!(425 <= 255)) || false) {
                                        v1 -= v1
                                        _ = len(v2)
                                    }
                                    var v83 float32 = 0.0
                                    _ = v83
                                } else {
                                    v4 = 10
                                }
                                v1 = 1
                                v1 += 100
                                if true {
                                    if false {
                                        var v84 uint64 = v5
                                        fmt.Println(v4)
                                        _ = v84
                                    }
                                    v4 = (v4 ^ (v4 + 4294967295))
                                    fmt.Println(v4, v1)
                                } else {
                                    if ((4294967295 == 940) || false) {
                                        v85 := GenS4{F0: fi77, F1: ("hello" + "\n\t"), F2: 336}
                                        var v86 int32 = -1
                                        v87 := []int{(fi77 ^ 32768)}
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                    } else {
                                        var v88 bool = true
                                        var v89 string = "foo bar"
                                        var v90 float64 = 1.5
                                        v91 := GenS1{F0: fi65, F1: 1, F2: 724}
                                        _ = v88
                                        _ = v89
                                        _ = v90
                                        _ = v91
                                    }
                                    fmt.Println(fi77, v1)
                                    fmt.Println(fi77, v4, v5)
                                    switch fi65 {
                                    case 0:
                                        v4 = (v4 << (((v4 ^ (^(v4)))) & 63))
                                    case 32767:
                                        v92 := GenS3{F0: 9223372036854775807, F1: v4, F2: 2147483647}
                                        _ = v92
                                    }
                                    var v93 bool = (!(1 < 128))
                                    _ = v93
                                }
                            }
                        }
                        _ = fi77
                    }
                    for fi94 := 0; fi94 < 2; fi94++ {
                        v95 := GenS3{F0: (fi94 - (fi65 + -22)), F1: (v4 % ((v4) | 1)), F2: 7}
                        v96 := v3["a"]
                        _ = fi94
                        _ = v95
                        _ = v96
                    }
                    v5 -= (v5 * v5)
                }
            }
            var v97 uint64 = (^(v5))
            fmt.Println(v5)
            _ = fi65
            _ = v97
        }
        var v98 float32 = 0.0
        _ = v98
    } else {
        if (false || true) {
            _ = len(v2)
        }
    }
    v99 := GenS5{F0: 128}
    if true {
        v4 = 255
        var v100 uint = v4
        for fi101 := 0; fi101 < 10; fi101++ {
            switch "hello" {
            case "\n\t":
                v102 := GenS3{F0: 4294967295, F1: 65535, F2: 2147483648}
                _ = v102
            case "":
                v103 := GenS3{F0: fi101, F1: v100, F2: 322}
                _ = v103
            default:
                if ("foo bar" <= "\n\t") {
                    if (!((!((151 <= 127) || (0 >= 42))) && false)) {
                        v100 -= (v4 << (((^(v100))) & 63))
                    }
                    for fi104 := 0; fi104 < 15; fi104++ {
                        if true {
                            switch 916 {
                            case 1:
                                var v105 uint16 = 1
                                _ = v105
                            case -234:
                                fmt.Println(fi101, v1)
                            case -174:
                                fmt.Println(fi104, v1)
                            default:
                                v106 := GenS3{F0: fi101, F1: v100, F2: 842}
                                _ = v106
                            }
                            switch (fi104 & (fi104 % ((fi101) | 1))) {
                            case -335:
                                var v107 uint8 = 42
                                _ = v107
                            case -2147483648:
                                var v108 uint8 = 88
                                _ = v108
                            }
                        } else {
                            switch (("hello" + ("a" + "\n\t")) + ("foo bar" + (("" + "") + ("\n\t" + "foo bar")))) {
                            case "":
                                v4 = v4
                                var v109 uint = 18446744073709551615
                                _ = v109
                            case "\n\t":
                                if false {
                                    var v110 uint = 42
                                    _ = v110
                                }
                            }
                            fmt.Println(fi104)
                            v111 := []int{32767}
                            fmt.Println(v4, fi101)
                            v100 = (v100 + v100)
                            _ = v111
                        }
                        switch fi101 {
                        case 32767:
                            if false {
                                v112 := []int{fi101, fi104}
                                switch "foo bar" {
                                case "a":
                                    var v113 int64 = 9223372036854775807
                                    _ = v113
                                case "\n\t":
                                    if len(v112) >= 2 {
                                        v112 = v112[1:len(v112)]
                                    }
                                    v5 = 0
                                    var v114 int64 = -1
                                    v1 += -128
                                    fmt.Println(v114)
                                    _ = v114
                                case "":
                                    if (true || false) {
                                        v4 = v100
                                        if len(v112) >= 2 {
                                            v112 = v112[1:len(v112)]
                                        }
                                        switch fi104 {
                                        case 406:
                                            fmt.Println(v4, v5, fi101)
                                        case 127:
                                            if ((((0.1 * (1e10 - ((1e10 * -2.25) + (3.14159 - -2.25)))) - -2.25) - (3.14159 * -2.25)) <= 1.5) {
                                                if ("hello" < ("hello" + "a")) {
                                                    if ((1e10 - -2.25) != (-2.25 / float64(((-2.25 + (float32(float64(1.5)) - 1e10)) / -2.25)))) {
                                                        fmt.Println(v4, fi104, fi101)
                                                        if true {
                                                            _ = len(v3)
                                                            v3[("a" + ("a" + ""))] = fi101
                                                            _ = len(v3)
                                                        }
                                                        v115 := v2["a"]
                                                        _ = v115
                                                    }
                                                } else {
                                                    var v116 string = ("a" + ("" + "hello"))
                                                    if (v116 <= (v116 + v116)) {
                                                        switch (fi104 % ((2147483647) | 1)) {
                                                        case -2147483648:
                                                            var v117 string = "a"
                                                            _ = v117
                                                        case 32767:
                                                            v4 += 255
                                                        default:
                                                            var v118 int = (fi101 & fi104)
                                                            _ = v118
                                                        }
                                                        if (v5 < (v5 << (((^(v100))) & 63))) {
                                                            var v119 uint64 = 42
                                                            var v120 float64 = 3.14159
                                                            fmt.Println(v120)
                                                            _ = v119
                                                            _ = v120
                                                        }
                                                        var v121 int = (fi104 * fi104)
                                                        _ = v121
                                                    } else {
                                                        if len(v112) >= 2 {
                                                            v112 = v112[1:len(v112)]
                                                        }
                                                    }
                                                    var v122 uint8 = 255
                                                    _ = v116
                                                    _ = v122
                                                }
                                                v123 := v3[(("" + "hello") + ("foo bar" + ("\n\t" + "a")))]
                                                delete(v2, "\n\t")
                                                _ = v123
                                            } else {
                                                v4 = 0
                                                fmt.Println(fi101, v1)
                                                var v124 float64 = (0.0 / 1.0)
                                                _ = v124
                                            }
                                        case 128:
                                            fmt.Println(fi104, v5, v100)
                                        }
                                    }
                                    if (!(1 >= 42)) {
                                        v100 = 18446744073709551615
                                        var v125 uint = v100
                                        fmt.Println(fi104)
                                        v126 := GenS3{F0: (^(fi101)), F1: (v100 >> ((1) & 63)), F2: 1}
                                        fmt.Println(fi104, v1)
                                        v5 = (v5 & 589)
                                        fmt.Println(v4)
                                        v5 = (v5 & (^(v5)))
                                        fmt.Println(fi101, v1, fi101)
                                        _ = v125
                                        _ = v126
                                    }
                                default:
                                    var v127 int16 = 0
                                    _ = v127
                                }
                                v4 += v100
                                _ = v112
                            }
                        case 127:
                            v4 -= (v4 - (v100 + (^(v4))))
                        case -1:
                            switch (("foo bar" + "a") + (("" + ("\n\t" + "")) + "hello")) {
                            case "":
                                switch fi101 {
                                case -9223372036854775808:
                                    v100 = v100
                                }
                            default:
                                v128 := GenS1{F0: fi104, F1: 929, F2: -709}
                                _ = v128
                            }
                        default:
                            v100 = v100
                        }
                        if true {
                            var v129 int32 = -1
                            var v130 uint16 = uint16(v5)
                            _ = v129
                            _ = v130
                        }
                        _ = fi104
                    }
                } else {
                    var v131 int32 = 0
                    var v132 uint32 = 995
                    fmt.Println(v4, v5, fi101)
                    _ = v131
                    _ = v132
                }
            }
            _ = fi101
        }
        _ = v100
    } else {
        v133 := []int{}
        _ = v133
    }
    v5 += v5
    var v134 uint16 = 255
    var v135 float32 = float32(0.1)
    v136 := v2["hello"]
    switch "foo bar" {
    case "a":
        v137 := GenS2{F0: (^(v134)), F1: "hello", F2: (^(v1))}
        _ = v137
    case "foo bar":
        for fi138 := 0; fi138 < 12; fi138++ {
            v134 = (^(v134))
            switch "hello" {
            case "":
                var v139 uint64 = 747
                _ = v139
            case "hello":
                var v140 string = ("" + "foo bar")
                _ = v140
            }
            v141 := GenS6{F0: 1000}
            _ = fi138
            _ = v141
        }
    case "\n\t":
        _ = len(v3)
    default:
        v142 := []int{-264, v136, v136, v136}
        _ = v142
    }
    v143 := GenS5{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v134:", v134)
    fmt.Println("v135:", v135)
    fmt.Println("v136:", v136)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v134
    _ = v135
    _ = v136
    _ = v143
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
