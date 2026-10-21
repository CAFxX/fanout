package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 string; F1 string; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 int; F2 string }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 bool; F1 int16; F2 uint8 }
type GenS9 struct { F0 int8; F1 uint8 }
type GenS10 struct { F0 int16; F1 int; F2 int }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint8; F1 int; F2 uint64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int32; F1 int64 }
type GenS15 struct { F0 uint8; F1 string; F2 int16 }
type GenS16 struct { F0 int8; F1 int16 }
type GenS17 struct { F0 int64; F1 uint }
type GenS18 struct { F0 int32; F1 uint8; F2 bool }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 uint64; F1 int; F2 int32 }
func genFn0(p0 uint) bool {
    var v5 int32 = 1
    for fi6 := 0; fi6 < 5; fi6++ {
        p0 -= (p0 >> (((p0 >> ((p0) & 63))) & 63))
        v5 = v5
        _ = fi6
    }
    p0 = p0
    if false {
        if false {
            v7 := []int{23, 127}
            v8 := GenS1{F0: 2147483648, F1: 94}
            if (true && (false && ((^(p0)) > 610))) {
                v5 += v5
            }
            _ = v7
            _ = v8
        }
        v9 := []int{127, 32767}
        var v10 int8 = -69
        _ = v10
        _ = v9
    } else {
        v11 := GenS2{F0: "foo bar", F1: "\n\t", F2: 18446744073709551615}
        switch "a" {
        case "hello":
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := GenS1{F0: 65535, F1: 2}
                for fi14 := 0; fi14 < 4; fi14++ {
                    if (413 > 741) {
                        var v15 bool = (((("" + "\n\t") < (("foo bar" + "a") + "a")) && (true || (((0.0 / 1.0) <= (1.5 + 1e10)) || (!true)))) || (1 == 9223372036854775807))
                        _ = v15
                    } else {
                        p0 = p0
                        v5 = (v5 - v5)
                    }
                    v5 = (v5 * (v5 * (^(v5))))
                    var v16 int = fi12
                    _ = fi14
                    _ = v16
                }
                for fi17 := 0; fi17 < 5; fi17++ {
                    v18 := []int{(-(fi17)), 42, fi12}
                    if (!(1 >= 2)) {
                        var v19 uint8 = 42
                        v19 -= 0
                        v18 = append(v18, (fi17 + -1))
                        _ = v19
                    } else {
                        p0 = p0
                        if true {
                            if true {
                                switch 1 {
                                case 128:
                                    var v20 float32 = 1.5
                                    switch "hello" {
                                    case "hello":
                                        v21 := GenS3{F0: 91}
                                        _ = v21
                                    case "":
                                        p0 += (p0 ^ (^(p0)))
                                    }
                                    _ = v20
                                }
                            } else {
                                _ = len(v18)
                                switch fi12 {
                                case -9223372036854775808:
                                    v5 = v5
                                case 4294967295:
                                    if len(v18) > 0 {
                                        v18[0] = fi12
                                    }
                                case 275:
                                    if (true || (false && false)) {
                                        v22 := GenS2{F0: ("\n\t" + ("\n\t" + ("" + "\n\t"))), F1: (("\n\t" + ((("\n\t" + "a") + "\n\t") + "\n\t")) + ("" + ("foo bar" + "foo bar"))), F2: (p0 << ((p0) & 63))}
                                        v23 := GenS1{F0: 4294967295, F1: 467}
                                        if ("hello" > ("a" + "")) {
                                            switch "hello" {
                                            case "foo bar":
                                                fmt.Println(p0)
                                            case "\n\t":
                                                p0 += (p0 - (p0 - p0))
                                            case "a":
                                                v24 := GenS4{F0: 4294967295, F1: 360}
                                                _ = v24
                                            }
                                            v5 = (-(v5))
                                            var v25 uint64 = 42
                                            _ = v25
                                        } else {
                                            if true {
                                                var v26 uint16 = 42
                                                var v27 float64 = 0.1
                                                v5 = v5
                                                _ = v26
                                                _ = v27
                                            }
                                            v28 := GenS5{F0: (false && false)}
                                            v5 = (v5 << ((42) & 63))
                                            _ = v28
                                        }
                                        _ = v22
                                        _ = v23
                                    }
                                default:
                                    if (-1 < 127) {
                                        p0 -= p0
                                        switch "foo bar" {
                                        case "a":
                                            v5 += (v5 - (v5 | (v5 % (((v5 % ((-1) | 1))) | 1))))
                                        case "":
                                            v5 = (v5 % ((-267) | 1))
                                        case "hello":
                                            var v29 int8 = -11
                                            _ = v29
                                        default:
                                            v30 := GenS6{F0: fi17, F1: 127, F2: "hello"}
                                            _ = v30
                                        }
                                        fmt.Println(fi17, fi17)
                                    } else {
                                        switch fi12 {
                                        case 128:
                                            fmt.Println(v5, fi17)
                                        case -9223372036854775808:
                                            if (-2147483648 != 32768) {
                                                v31 := GenS7{F0: 2, F1: false}
                                                var v32 float32 = 1.5
                                                if len(v18) > 0 {
                                                    v18[0] = fi12
                                                }
                                                _ = v31
                                                _ = v32
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v33 int32 = (v5 % (((v5 | v5)) | 1))
                                                    _ = v33
                                                case "foo bar":
                                                    v34 := GenS4{F0: 107, F1: 32767}
                                                    _ = v34
                                                case "hello":
                                                    var v35 uint8 = 127
                                                    _ = v35
                                                default:
                                                    switch "" {
                                                    case "hello":
                                                        var v36 int = -1
                                                        _ = v36
                                                    case "":
                                                        var v37 int32 = -926
                                                        _ = v37
                                                    case "a":
                                                        v5 += v5
                                                        var v38 int64 = -9223372036854775808
                                                        if len(v18) > 0 {
                                                            v18[0] = -9223372036854775808
                                                        }
                                                        _ = v38
                                                    }
                                                }
                                            }
                                        }
                                        var v39 uint = (p0 - (p0 - (p0 << (((p0 & 42)) & 63))))
                                        v39 -= (^(p0))
                                        _ = v39
                                    }
                                }
                            }
                            var v40 int64 = 796
                            _ = v40
                        }
                        fmt.Println(fi17, fi17)
                    }
                    _ = fi17
                    _ = v18
                }
                _ = fi12
                _ = v13
            }
        default:
            v41 := GenS7{F0: 651, F1: false}
            _ = v41
        }
        if true {
            v42 := GenS2{F0: ("\n\t" + ("foo bar" + "\n\t")), F1: "hello", F2: 18446744073709551615}
            fmt.Println(v5)
            switch 127 {
            case 128:
                for fi43 := 0; fi43 < 6; fi43++ {
                    if (1 <= 2) {
                        p0 += (p0 * (p0 / (((p0 / ((65535) | 1))) | 1)))
                        fmt.Println(fi43)
                    } else {
                        fmt.Println(v5, fi43)
                        if (-53 > 0) {
                            var v44 int = (fi43 << ((2) & 63))
                            _ = v44
                        }
                        var v45 uint64 = 42
                        _ = v45
                    }
                    _ = fi43
                }
            default:
                for fi46 := 0; fi46 < 4; fi46++ {
                    var v47 string = "hello"
                    _ = fi46
                    _ = v47
                }
            }
            _ = v42
        } else {
            p0 = (p0 ^ 255)
            if (("a" + "hello") == ((("hello" + "hello") + "foo bar") + ("foo bar" + "hello"))) {
                if (811 > 18446744073709551615) {
                    fmt.Println(v5, p0)
                    for fi48 := 0; fi48 < 11; fi48++ {
                        var v49 uint64 = 42
                        for fi50 := 0; fi50 < 9; fi50++ {
                            p0 = (p0 / (((p0 * (p0 & (p0 << ((4294967295) & 63))))) | 1))
                            var v51 uint16 = 1
                            _ = fi50
                            _ = v51
                        }
                        _ = fi48
                        _ = v49
                    }
                } else {
                    if true {
                        if (false && false) {
                            v52 := GenS8{F0: (255 != 128), F1: 0, F2: 42}
                            if (848 == 600) {
                                if false {
                                    if false {
                                        for fi53 := 0; fi53 < 2; fi53++ {
                                            v54 := GenS9{F0: -47, F1: 1}
                                            v55 := GenS10{F0: 1000, F1: (fi53 * 844), F2: (fi53 - fi53)}
                                            _ = fi53
                                            _ = v54
                                            _ = v55
                                        }
                                        v56 := GenS2{F0: "a", F1: "\n\t", F2: p0}
                                        _ = v56
                                    }
                                    var v57 float64 = (0.1 - 0.1)
                                    var v58 string = ("" + "\n\t")
                                    _ = v57
                                    _ = v58
                                } else {
                                    if true {
                                        for fi59 := 0; fi59 < 8; fi59++ {
                                            var v60 uint = p0
                                            v61 := []int{(fi59 | fi59)}
                                            _ = fi59
                                            _ = v60
                                            _ = v61
                                        }
                                    }
                                    fmt.Println(v5)
                                    var v62 uint = p0
                                    _ = v62
                                }
                                v63 := []int{}
                                for fi64 := 0; fi64 < 2; fi64++ {
                                    for fi65 := 0; fi65 < 15; fi65++ {
                                        p0 = p0
                                        p0 = (p0 + (p0 / (((p0 << ((p0) & 63))) | 1)))
                                        v66 := GenS11{F0: (^(v5))}
                                        _ = fi65
                                        _ = v66
                                    }
                                    _ = fi64
                                }
                                _ = v63
                            } else {
                                switch 0 {
                                case -158:
                                    v67 := []int{}
                                    _ = v67
                                default:
                                    fmt.Println(v5, p0)
                                }
                            }
                            _ = v52
                        } else {
                            p0 += (p0 * (p0 ^ p0))
                            v68 := GenS1{F0: 977, F1: 22}
                            _ = v68
                        }
                        for fi69 := 0; fi69 < 9; fi69++ {
                            for fi70 := 0; fi70 < 9; fi70++ {
                                p0 = 176
                                v71 := []int{9223372036854775807, (fi70 * fi69), fi70, fi70}
                                _ = fi70
                                _ = v71
                            }
                            v72 := GenS6{F0: fi69, F1: fi69, F2: "a"}
                            _ = fi69
                            _ = v72
                        }
                        v5 += v5
                    } else {
                        v73 := []int{-2147483648, -874, 42}
                        _ = v73
                    }
                    if (328 < 0) {
                        p0 -= 18446744073709551615
                        switch (("" + ("a" + ("a" + ""))) + "\n\t") {
                        case "a":
                            for fi74 := 0; fi74 < 7; fi74++ {
                                if (v5 == (v5 & v5)) {
                                    var v75 int = fi74
                                    v5 += v5
                                    _ = v75
                                } else {
                                    var v76 uint = (^(p0))
                                    switch 9223372036854775807 {
                                    case -9223372036854775808:
                                        switch ("a" + "") {
                                        case "":
                                            for fi77 := 0; fi77 < 19; fi77++ {
                                                switch "hello" {
                                                case "a":
                                                    v76 += (^(p0))
                                                default:
                                                    var v78 uint8 = 2
                                                    _ = v78
                                                }
                                                var v79 uint = (^(p0))
                                                _ = fi77
                                                _ = v79
                                            }
                                        case "hello":
                                            switch "hello" {
                                            case "a":
                                                if (true && true) {
                                                    switch ("" + ("a" + "\n\t")) {
                                                    case "\n\t":
                                                        v5 = 228
                                                    case "":
                                                        fmt.Println(p0)
                                                    }
                                                    var v80 int = 9223372036854775807
                                                    _ = v80
                                                } else {
                                                    if (!(280 < 850)) {
                                                        for fi81 := 0; fi81 < 14; fi81++ {
                                                            v5 = int32(p0)
                                                            _ = fi81
                                                        }
                                                        var v82 uint16 = 984
                                                        _ = v82
                                                    }
                                                }
                                            case "hello":
                                                if false {
                                                    if (1 != 780) {
                                                        p0 = p0
                                                        fmt.Println(v5)
                                                        if (2147483647 > 529) {
                                                            v83 := GenS12{F0: 151, F1: (^(fi74)), F2: 42}
                                                            v84 := GenS2{F0: "\n\t", F1: "", F2: (^(p0))}
                                                            if false {
                                                                switch 0 {
                                                                case 9223372036854775807:
                                                                    var v85 uint32 = 2
                                                                    _ = v85
                                                                case 2147483647:
                                                                    for fi86 := 0; fi86 < 6; fi86++ {
                                                                        p0 += v76
                                                                        v76 = 42
                                                                        v87 := GenS13{F0: 18446744073709551615}
                                                                        _ = fi86
                                                                        _ = v87
                                                                    }
                                                                case 4294967295:
                                                                    fmt.Println(v76, v76, v76)
                                                                default:
                                                                    for fi88 := 0; fi88 < 10; fi88++ {
                                                                        v89 := []int{(^(fi88))}
                                                                        v5 -= v5
                                                                        _ = len(v89)
                                                                        _ = fi88
                                                                        _ = v89
                                                                    }
                                                                }
                                                            }
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                    } else {
                                                        v76 += v76
                                                    }
                                                    if false {
                                                        v76 -= 0
                                                    } else {
                                                        var v90 int16 = 32767
                                                        v91 := []int{}
                                                        p0 = (p0 % (((p0 ^ (p0 / ((v76) | 1)))) | 1))
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                }
                                            case "\n\t":
                                                var v92 bool = true
                                                _ = v92
                                            default:
                                                if false {
                                                    if false {
                                                        v5 += (v5 | 1)
                                                    } else {
                                                        if (-9223372036854775808 < 42) {
                                                            fmt.Println(fi74)
                                                        } else {
                                                            if (!(!(true && (32768 < fi74)))) {
                                                                var v93 uint32 = 731
                                                                var v94 uint32 = (v93 + v93)
                                                                switch 32768 {
                                                                case 42:
                                                                    v95 := GenS14{F0: -2147483648, F1: 42}
                                                                    _ = v95
                                                                case 127:
                                                                    for fi96 := 0; fi96 < 4; fi96++ {
                                                                        v76 = 18446744073709551615
                                                                        fmt.Println(v5, p0)
                                                                        v94 = uint32(p0)
                                                                        var v97 float32 = 0.1
                                                                        _ = fi96
                                                                        _ = v97
                                                                    }
                                                                }
                                                                _ = v93
                                                                _ = v94
                                                            }
                                                            v5 = -1
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v98 int16 = 1
                                                        _ = v98
                                                    case "a":
                                                        v76 -= p0
                                                    case "foo bar":
                                                        var v99 bool = (!true)
                                                        _ = v99
                                                    }
                                                    var v100 int = fi74
                                                    _ = v100
                                                }
                                            }
                                        case "a":
                                            v76 = p0
                                        default:
                                            if (!(v5 <= v5)) {
                                                v101 := []int{0, int(v76), fi74, (^(fi74))}
                                                if (0.1 != (-2.25 - 1.5)) {
                                                    var v102 string = ("" + "a")
                                                    switch fi74 {
                                                    case 2147483647:
                                                        v76 -= v76
                                                    case 817:
                                                        v76 = (p0 / ((v76) | 1))
                                                    default:
                                                        v103 := GenS15{F0: 127, F1: "hello", F2: -918}
                                                        _ = v103
                                                    }
                                                    v76 += v76
                                                    _ = v102
                                                } else {
                                                    p0 += v76
                                                }
                                                _ = v101
                                            } else {
                                                fmt.Println(v76)
                                                for fi104 := 0; fi104 < 16; fi104++ {
                                                    switch (fi74 >> (((^(p0))) & 63)) {
                                                    case 32767:
                                                        v105 := GenS16{F0: 127, F1: -727}
                                                        _ = v105
                                                    case -258:
                                                        var v106 uint32 = 2147483648
                                                        _ = v106
                                                    case 0:
                                                        fmt.Println(fi74, fi104, fi104)
                                                    default:
                                                        fmt.Println(v5, v76)
                                                    }
                                                    _ = fi104
                                                }
                                            }
                                        }
                                    case -75:
                                        v5 += (v5 + (v5 >> ((p0) & 63)))
                                    case 42:
                                        for fi107 := 0; fi107 < 3; fi107++ {
                                            if false {
                                                if false {
                                                    switch "" {
                                                    case "hello":
                                                        p0 = 42
                                                    case "":
                                                        if ((("a" + "a") + "") < "\n\t") {
                                                            v5 -= v5
                                                        }
                                                        var v108 int = -412
                                                        fmt.Println(v76, fi107, v108)
                                                        _ = v108
                                                    case "a":
                                                        var v109 float64 = (float64(0.1) - (0.1 / float64(1e10)))
                                                        _ = v109
                                                    }
                                                    v110 := []int{(fi107 | 128)}
                                                    v76 = v76
                                                    switch ("hello" + "") {
                                                    case "foo bar":
                                                        if false {
                                                            switch fi74 {
                                                            case -936:
                                                                v5 = v5
                                                                var v111 string = "foo bar"
                                                                v112 := GenS14{F0: (v5 | -761), F1: -9223372036854775808}
                                                                _ = v111
                                                                _ = v112
                                                            }
                                                            if (!true) {
                                                                switch ("" + ("foo bar" + "a")) {
                                                                case "":
                                                                    fmt.Println(p0, v76, fi74)
                                                                default:
                                                                    if len(v110) > 0 {
                                                                        v110[0] = fi74
                                                                    }
                                                                    var v113 uint8 = 195
                                                                    var v114 int16 = 42
                                                                    _ = v113
                                                                    _ = v114
                                                                }
                                                                fmt.Println(p0, v5)
                                                                v115 := GenS10{F0: 1000, F1: (fi74 >> ((147) & 63)), F2: fi107}
                                                                var v116 uint32 = uint32(v76)
                                                                _ = v115
                                                                _ = v116
                                                            }
                                                            v110 = append(v110, fi74)
                                                            var v117 uint8 = 215
                                                            var v118 uint = (p0 + p0)
                                                            _ = v117
                                                            _ = v118
                                                        }
                                                    case "\n\t":
                                                        var v119 string = "foo bar"
                                                        switch -273 {
                                                        case 4294967295:
                                                            if len(v110) > 0 {
                                                                v110[0] = (fi74 * fi74)
                                                            }
                                                        case 128:
                                                            if (false || (false && true)) {
                                                                if false {
                                                                    if len(v110) > 0 {
                                                                        v110[0] = (-(fi74))
                                                                    }
                                                                    var v120 float64 = 3.14159
                                                                    var v121 int32 = v5
                                                                    _ = v120
                                                                    _ = v121
                                                                } else {
                                                                    var v122 string = "foo bar"
                                                                    if true {
                                                                        fmt.Println(p0, p0, fi74)
                                                                        var v123 uint16 = 529
                                                                        v124 := GenS13{F0: 547}
                                                                        var v125 uint16 = v123
                                                                        _ = v123
                                                                        _ = v124
                                                                        _ = v125
                                                                    }
                                                                    fmt.Println(p0)
                                                                    var v126 uint64 = 615
                                                                    fmt.Println(v122)
                                                                    _ = v122
                                                                    _ = v126
                                                                }
                                                                var v127 int64 = 1
                                                                _ = v127
                                                            } else {
                                                                var v128 int16 = -1
                                                                _ = v128
                                                            }
                                                        case 9223372036854775807:
                                                            var v129 int32 = -976
                                                            _ = v129
                                                        default:
                                                            var v130 uint64 = 963
                                                            _ = v130
                                                        }
                                                        var v131 uint8 = 93
                                                        _ = v119
                                                        _ = v131
                                                    }
                                                    v76 -= 4294967295
                                                    _ = v110
                                                }
                                                p0 = 65535
                                            }
                                            v132 := GenS9{F0: 42, F1: 2}
                                            v133 := []int{fi107, 1}
                                            _ = fi107
                                            _ = v132
                                            _ = v133
                                        }
                                    default:
                                        v134 := GenS17{F0: 0, F1: p0}
                                        _ = v134
                                    }
                                    _ = v76
                                }
                                v135 := []int{}
                                var v136 int64 = 788
                                _ = fi74
                                _ = v135
                                _ = v136
                            }
                        default:
                            v137 := GenS8{F0: (3.14159 != (3.14159 + float32((0.0 + (1e10 / 1.0))))), F1: 1000, F2: 4}
                            _ = v137
                        }
                        var v138 int32 = (v5 ^ (v5 << ((p0) & 63)))
                        _ = v138
                    }
                }
                v139 := GenS2{F0: "", F1: "\n\t", F2: 1}
                _ = v139
            } else {
                fmt.Println(v5)
                var v140 int32 = v5
                p0 = (p0 - (p0 / ((p0) | 1)))
                _ = v140
            }
        }
        _ = v11
    }
    _ = p0
    _ = v5
    return (479 == 1000)
}
func genFn1(p0 uint64) uint64 {
    switch 714 {
    case -9223372036854775808:
        fmt.Println(p0)
    case -2147483648:
        v141 := []int{32768, 32767, 472, -9223372036854775808}
        _ = v141
    }
    var v142 uint32 = 930
    _ = p0
    _ = v142
    return p0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 255
    var v3 uint8 = 42
    var v4 bool = (1 >= 589)
    fmt.Println(v4)
    for fi143 := 0; fi143 < 19; fi143++ {
        var v144 uint32 = 32
        _ = fi143
        _ = v144
    }
    v145 := GenS11{F0: 0}
    v146 := []int{-9223372036854775808}
    v147 := GenS18{F0: 0, F1: (v3 * (v3 << ((804) & 63))), F2: v4}
    var v148 int8 = -1
    if ((v2 >> ((uint(v3)) & 63)) > (v2 >> ((704) & 63))) {
        if false {
            fmt.Println(v2)
            v149 := v1[(("foo bar" + ("a" + "\n\t")) + (("foo bar" + "a") + "foo bar"))]
            fmt.Println(v2, v4)
            _ = v149
        } else {
            v150 := GenS16{F0: (v148 & 0), F1: -1}
            _ = v150
        }
        if (2147483647 >= -483) {
            if v4 {
                switch "a" {
                case "foo bar":
                    switch (("" + "hello") + "") {
                    case "":
                        v4 = (v4 && v4)
                    case "foo bar":
                        v148 = 107
                    case "\n\t":
                        var v151 uint64 = 0
                        _ = v151
                    default:
                        var v152 int8 = v148
                        _ = v152
                    }
                case "":
                    v153 := v1["a"]
                    _ = v153
                }
                v154 := v1["a"]
                v3 = v3
                _ = v154
            }
            switch "a" {
            case "foo bar":
                v155 := GenS19{F0: -1}
                _ = v155
            case "\n\t":
                if len(v146) >= 2 {
                    v146 = v146[1:len(v146)]
                }
            case "hello":
                delete(v1, ("foo bar" + (("" + (((("hello" + "hello") + ("" + "foo bar")) + "") + ("hello" + "foo bar"))) + "hello")))
            }
        } else {
            v1["foo bar"] = 373
        }
        v4 = ((42 <= 846) || true)
    } else {
        for fi156 := 0; fi156 < 3; fi156++ {
            var v157 uint16 = 962
            v2 = v157
            if (1 >= 660) {
                for fi158 := 0; fi158 < 9; fi158++ {
                    if len(v146) >= 2 {
                        v146 = v146[1:len(v146)]
                    }
                    v157 -= v157
                    var v159 string = ("hello" + "")
                    _ = fi158
                    _ = v159
                }
                var v160 uint32 = 0
                _ = v160
            } else {
                v1[""] = 673
            }
            _ = fi156
            _ = v157
        }
    }
    if v4 {
        for fi161 := 0; fi161 < 5; fi161++ {
            v162 := v1["\n\t"]
            _ = fi161
            _ = v162
        }
    }
    _ = len(v146)
    var v163 int = 2147483647
    if ((v2 >> ((2) & 63)) >= 1) {
        for fi164 := 0; fi164 < 3; fi164++ {
            var v165 float32 = -2.25
            _ = fi164
            _ = v165
        }
        if len(v146) >= 2 {
            v146 = v146[1:len(v146)]
        }
        _ = len(v146)
    } else {
        switch (v163 << ((0) & 63)) {
        case -385:
            v146 = append(v146, 20)
        case 128:
            delete(v1, "\n\t")
        }
    }
    v166 := GenS20{F0: (v148 | v148)}
    var v167 int32 = 496
    v168 := v1["foo bar"]
    fmt.Println(v148)
    v169 := GenS21{F0: 0, F1: (-(v163)), F2: v167}
    v3 = v3
    switch v163 {
    case 42:
        switch "\n\t" {
        case "a":
            switch v163 {
            case 4294967295:
                if v4 {
                    var v170 bool = v4
                    v163 = (v163 + (v163 & v163))
                    _ = v170
                }
            default:
                var v171 uint = 4294967295
                _ = v171
            }
        case "\n\t":
            switch (v163 + (v163 % ((v163) | 1))) {
            case 2147483647:
                for fi172 := 0; fi172 < 16; fi172++ {
                    var v173 int = v163
                    _ = fi172
                    _ = v173
                }
            case 0:
                var v174 int8 = v148
                _ = v174
            default:
                v175 := v1[(("" + "") + "a")]
                _ = v175
            }
        case "":
            v146 = append(v146, 127)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v146:", len(v146), v146)
    fmt.Println("v148:", v148)
    fmt.Println("v163:", v163)
    fmt.Println("v167:", v167)
    fmt.Println("v168:", v168)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v145
    _ = v146
    _ = v147
    _ = v148
    _ = v163
    _ = v166
    _ = v167
    _ = v168
    _ = v169
    _ = v2
    _ = v3
    _ = v4
}
