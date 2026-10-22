package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 string }
type GenS2 struct { F0 string; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 uint32; F1 uint16 }
type GenS10 struct { F0 uint; F1 uint16 }
type GenS11 struct { F0 uint; F1 uint8 }
type GenS12 struct { F0 int64; F1 string }
type GenS13 struct { F0 uint64; F1 uint32; F2 string }
type GenS14 struct { F0 string }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 uint16; F1 uint }
type GenS17 struct { F0 uint16; F1 int32 }
type GenS18 struct { F0 string; F1 uint16 }
type GenS19 struct { F0 int32; F1 int8; F2 uint }
type GenS20 struct { F0 int64 }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 int64; F1 int16; F2 bool }
type GenS23 struct { F0 int64; F1 bool; F2 int }
type GenS24 struct { F0 int64; F1 int }
type GenS25 struct { F0 int32 }
func genFn0(p0 int16) uint {
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 8; fi5++ {
            fmt.Println(fi4, fi5)
            switch (fi4 ^ fi5) {
            case 42:
                p0 -= (p0 + (p0 | (p0 - (p0 ^ (p0 | (p0 % ((1) | 1)))))))
            case 4294967295:
                p0 += -1
            }
            _ = fi5
        }
        p0 -= (^(p0))
        var v6 uint = 18446744073709551615
        _ = fi4
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        if false {
            var v8 uint64 = 18446744073709551615
            if ("" == ("foo bar" + ("" + "hello"))) {
                p0 -= p0
            }
            _ = v8
        } else {
            v9 := []int{}
            if (float64(-2.25) >= 3.14159) {
                for fi10 := 0; fi10 < 17; fi10++ {
                    fmt.Println(fi10, p0)
                    _ = fi10
                }
            }
            for fi11 := 0; fi11 < 6; fi11++ {
                switch -379 {
                case 4294967295:
                    p0 -= p0
                case 32767:
                    if (!true) {
                        if (860 == -1) {
                            var v12 uint16 = 1
                            switch "" {
                            case "foo bar":
                                v13 := GenS1{F0: 42, F1: -1, F2: ("hello" + "a")}
                                _ = v13
                            case "hello":
                                if false {
                                    v14 := GenS2{F0: (("foo bar" + "foo bar") + ""), F1: -128, F2: 2}
                                    _ = v14
                                } else {
                                    var v15 uint8 = 209
                                    v16 := GenS3{F0: 0, F1: v12}
                                    var v17 int32 = -2147483648
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                            }
                            _ = v12
                        }
                        p0 = -989
                        var v18 string = "foo bar"
                        _ = v18
                    } else {
                        v19 := GenS1{F0: 0, F1: 2147483647, F2: "foo bar"}
                        p0 += p0
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        _ = v19
                    }
                case 128:
                    if true {
                        v9 = append(v9, 128)
                        switch "foo bar" {
                        case "\n\t":
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(fi7, fi11)
                            case "a":
                                switch ("hello" + "\n\t") {
                                case "\n\t":
                                    v20 := GenS1{F0: 0, F1: 2147483647, F2: "hello"}
                                    _ = v20
                                default:
                                    var v21 float32 = 3.14159
                                    _ = v21
                                }
                            case "\n\t":
                                var v22 uint32 = 65535
                                _ = v22
                            default:
                                v23 := GenS3{F0: 65535, F1: 773}
                                _ = v23
                            }
                        case "foo bar":
                            v24 := GenS2{F0: "a", F1: 0, F2: 128}
                            _ = v24
                        }
                        v25 := GenS3{F0: 42, F1: 368}
                        _ = v25
                    } else {
                        fmt.Println(fi7)
                        v26 := GenS3{F0: 27, F1: 42}
                        _ = v26
                    }
                    p0 = p0
                default:
                    var v27 uint8 = 2
                    _ = v27
                }
                if ((752 <= p0) || false) {
                    var v28 float32 = (0.0 + (1.5 / (1e10 - 0.1)))
                    var v29 int32 = -991
                    _ = v28
                    _ = v29
                } else {
                    if false {
                        fmt.Println(fi11, fi7, p0)
                        var v30 float32 = -2.25
                        _ = v30
                    } else {
                        v9 = append(v9, fi11)
                        if true {
                            if true {
                                if ((949 < 963) || ((1 != 1) || ((65535 <= 953) && (18 != 101)))) {
                                    var v31 uint64 = 2
                                    var v32 uint = 65535
                                    v33 := GenS2{F0: "", F1: -128, F2: 127}
                                    fmt.Println(v32, fi7, fi7)
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v34 int = fi11
                                    if true {
                                        if false {
                                            switch (fi11 >> ((65535) & 63)) {
                                            case 649:
                                                if true {
                                                    _ = len(v9)
                                                    if (!(true && false)) {
                                                        if false {
                                                            p0 -= 32767
                                                        }
                                                        var v35 bool = false
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                        v36 := GenS4{F0: 1, F1: 2147483647}
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                } else {
                                                    fmt.Println(fi11, v34, fi11)
                                                    p0 += 32767
                                                }
                                            case 127:
                                                p0 = 42
                                            case 32767:
                                                if false {
                                                    if (1.5 <= 1e10) {
                                                        v37 := GenS5{F0: 255, F1: fi11}
                                                        var v38 uint64 = 2
                                                        v9 = append(v9, fi7)
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                } else {
                                                    if (false && false) {
                                                        p0 -= p0
                                                        fmt.Println(p0, v34, v34)
                                                        var v39 int16 = p0
                                                        _ = v39
                                                    } else {
                                                        if len(v9) > 0 {
                                                            v9[0] = 612
                                                        }
                                                        p0 += -1
                                                    }
                                                    switch 9223372036854775807 {
                                                    case 128:
                                                        _ = len(v9)
                                                    case -490:
                                                        var v40 bool = true
                                                        _ = v40
                                                    }
                                                    if len(v9) > 0 {
                                                        v9[0] = fi7
                                                    }
                                                }
                                            default:
                                                var v41 string = ""
                                                _ = v41
                                            }
                                            var v42 string = ("a" + "foo bar")
                                            v43 := GenS4{F0: 4294967295, F1: -797}
                                            var v44 uint8 = 127
                                            if ((false || true) || (255 != 42)) {
                                                v45 := GenS6{F0: 4294967295}
                                                _ = v45
                                            } else {
                                                if true {
                                                    p0 -= p0
                                                } else {
                                                    v44 -= v44
                                                    p0 = 42
                                                    var v46 uint8 = 2
                                                    _ = v46
                                                }
                                                if false {
                                                    switch fi7 {
                                                    case 1:
                                                        switch "a" {
                                                        case "a":
                                                            v47 := GenS7{F0: (v44 << ((255) & 63))}
                                                            _ = v47
                                                        default:
                                                            var v48 uint = 18446744073709551615
                                                            _ = v48
                                                        }
                                                    default:
                                                        p0 += -32768
                                                    }
                                                }
                                                _ = len(v9)
                                                switch 42 {
                                                case 628:
                                                    var v49 float32 = (3.14159 / 1.0)
                                                    _ = v49
                                                default:
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                }
                                                v50 := GenS8{F0: v42, F1: false}
                                                p0 = (^(p0))
                                                if true {
                                                    fmt.Println(fi11)
                                                } else {
                                                    var v51 uint32 = 357
                                                    _ = v51
                                                }
                                                var v52 uint32 = 858
                                                var v53 int16 = (p0 - int16(v52))
                                                _ = v50
                                                _ = v52
                                                _ = v53
                                            }
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        } else {
                                            switch fi7 {
                                            case 32768:
                                                fmt.Println(v34, fi7)
                                            case 4294967295:
                                                v34 = fi7
                                            case 424:
                                                switch ((("a" + "hello") + "\n\t") + "") {
                                                case "foo bar":
                                                    p0 -= 0
                                                case "a":
                                                    switch (fi7 % ((fi11) | 1)) {
                                                    case 2147483647:
                                                        v9 = append(v9, (fi7 ^ fi11))
                                                    default:
                                                        var v54 float32 = float32(3.14159)
                                                        _ = v54
                                                    }
                                                }
                                            }
                                            if false {
                                                fmt.Println(p0, fi7)
                                                var v55 uint = 494
                                                v56 := GenS1{F0: 42, F1: 2147483647, F2: ""}
                                                _ = v55
                                                _ = v56
                                            } else {
                                                if (!false) {
                                                    v34 = (^(v34))
                                                } else {
                                                    if false {
                                                        var v57 string = ("" + ("a" + "a"))
                                                        var v58 int32 = -1
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        switch (("hello" + "\n\t") + "") {
                                                        case "\n\t":
                                                            if (true || (18446744073709551615 == 2)) {
                                                                v9 = append(v9, fi11)
                                                            }
                                                            fmt.Println(p0, fi7, fi7)
                                                            var v59 int16 = (^(p0))
                                                            _ = v59
                                                        case "":
                                                            if true {
                                                                v60 := GenS9{F0: 42, F1: 65535}
                                                                _ = v60
                                                            } else {
                                                                switch "foo bar" {
                                                                case "a":
                                                                    if true {
                                                                        var v61 uint64 = 1
                                                                        if (-351 <= -837) {
                                                                            if len(v9) > 0 {
                                                                                v9[0] = (-(fi11))
                                                                            }
                                                                            v62 := GenS10{F0: 255, F1: 32767}
                                                                            var v63 uint = 4294967295
                                                                            _ = v62
                                                                            _ = v63
                                                                        }
                                                                        _ = v61
                                                                    } else {
                                                                        var v64 int32 = -2147483648
                                                                        var v65 uint8 = 128
                                                                        _ = v64
                                                                        _ = v65
                                                                    }
                                                                case "":
                                                                    switch ("a" + ("a" + "\n\t")) {
                                                                    case "hello":
                                                                        fmt.Println(fi7, fi11)
                                                                    case "":
                                                                        switch "" {
                                                                        case "hello":
                                                                            var v66 int64 = -9223372036854775808
                                                                            switch -1 {
                                                                            case 2147483647:
                                                                                v67 := GenS6{F0: 298}
                                                                                fmt.Println(p0, fi11)
                                                                                v68 := GenS5{F0: 2, F1: (v34 / ((fi11) | 1))}
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                            _ = v66
                                                                        case "a":
                                                                            v69 := GenS11{F0: 799, F1: 42}
                                                                            _ = v69
                                                                        }
                                                                    case "foo bar":
                                                                        switch 2147483647 {
                                                                        case 127:
                                                                            var v70 int = v34
                                                                            _ = v70
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    var v71 int8 = 100
                                                                    _ = v71
                                                                }
                                                                v34 = 0
                                                            }
                                                        }
                                                        fmt.Println(fi11, fi11, fi11)
                                                    }
                                                    p0 -= p0
                                                    fmt.Println(p0, v34, fi7)
                                                }
                                            }
                                            var v72 uint = 593
                                            _ = v72
                                        }
                                    }
                                    switch (fi7 ^ fi11) {
                                    case 4294967295:
                                        p0 += (-(p0))
                                    case 32768:
                                        v73 := GenS11{F0: 42, F1: 42}
                                        _ = v73
                                    }
                                    _ = v34
                                }
                                if ("hello" < "") {
                                    if false {
                                        v74 := GenS7{F0: 255}
                                        switch (-(fi7)) {
                                        case 4294967295:
                                            switch (^(fi11)) {
                                            case 128:
                                                if (!(!false)) {
                                                    if (!true) {
                                                        v9 = append(v9, fi7)
                                                        var v75 uint32 = 667
                                                        var v76 int = fi11
                                                        _ = v75
                                                        _ = v76
                                                    } else {
                                                        fmt.Println(p0, p0)
                                                        v77 := GenS6{F0: 215}
                                                        var v78 float64 = 0.0
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                }
                                                p0 += 1000
                                            case 2147483647:
                                                v79 := GenS12{F0: 231, F1: "\n\t"}
                                                _ = v79
                                            case 1:
                                                p0 = p0
                                            }
                                        case 32768:
                                            if false {
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                            } else {
                                                if false {
                                                    if (0 < (fi11 | fi7)) {
                                                        if true {
                                                            if true {
                                                                p0 = (^(p0))
                                                                v80 := GenS2{F0: "", F1: 100, F2: 1}
                                                                v81 := GenS6{F0: 42}
                                                                _ = v80
                                                                _ = v81
                                                            }
                                                            v9 = append(v9, -9223372036854775808)
                                                            switch "\n\t" {
                                                            case "":
                                                                if (363 != 42) {
                                                                    var v82 uint8 = 128
                                                                    var v83 int8 = -114
                                                                    var v84 float32 = 1.5
                                                                    _ = v82
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                            case "\n\t":
                                                                if len(v9) >= 2 {
                                                                    v9 = v9[1:len(v9)]
                                                                }
                                                            case "a":
                                                                p0 += (p0 / ((0) | 1))
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    _ = len(v9)
                                                    p0 -= 32767
                                                    v85 := GenS4{F0: 947, F1: 2147483647}
                                                    _ = v85
                                                }
                                            }
                                        default:
                                            var v86 float32 = 0.1
                                            _ = v86
                                        }
                                        _ = v74
                                    } else {
                                        switch "\n\t" {
                                        case "foo bar":
                                            p0 = (p0 >> ((971) & 63))
                                        case "a":
                                            if (-1 > p0) {
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                p0 += (p0 % ((1000) | 1))
                                            }
                                            var v87 int32 = -610
                                            _ = v87
                                        case "\n\t":
                                            p0 = (p0 >> ((18446744073709551615) & 63))
                                        }
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v88 := GenS13{F0: 4294967295, F1: 147, F2: "\n\t"}
                                        _ = v88
                                    }
                                }
                            } else {
                                fmt.Println(p0)
                                var v89 float64 = (0.1 - 0.1)
                                _ = v89
                            }
                        } else {
                            v90 := GenS7{F0: 128}
                            _ = v90
                        }
                    }
                    var v91 uint64 = 42
                    _ = v91
                }
                p0 += (^(p0))
                _ = fi11
            }
            _ = v9
        }
        _ = fi7
    }
    v92 := GenS5{F0: 42, F1: 127}
    var v93 uint16 = 831
    _ = p0
    _ = v92
    _ = v93
    return 65535
}

func main() {
    v1 := []int{128, 222, 9223372036854775807}
    var v2 uint64 = 2
    var v3 int32 = 42
    for fi94 := 0; fi94 < 20; fi94++ {
        switch "a" {
        case "a":
            v95 := GenS14{F0: ""}
            _ = v95
        }
        _ = fi94
    }
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    if (4294967295 != 127) {
        for fi96 := 0; fi96 < 3; fi96++ {
            var v97 uint32 = 42
            if (!false) {
                fmt.Println(v2, v3, v2)
                if false {
                    v2 = (v2 >> ((uint(v2)) & 63))
                    v3 += (v3 >> ((1) & 63))
                }
                for fi98 := 0; fi98 < 6; fi98++ {
                    var v99 bool = false
                    v100 := GenS15{F0: (!v99)}
                    v2 += v2
                    _ = fi98
                    _ = v100
                    _ = v99
                }
            }
            _ = fi96
            _ = v97
        }
        if false {
            var v101 int64 = int64(v3)
            for fi102 := 0; fi102 < 13; fi102++ {
                fmt.Println(v101, v101, v101)
                fmt.Println(v3, v101)
                v1 = append(v1, 32767)
                _ = fi102
            }
            _ = v101
        } else {
            fmt.Println(v3, v3)
            fmt.Println(v3, v2)
        }
    } else {
        v3 -= -1
    }
    if true {
        v3 -= v3
    }
    if false {
        v103 := GenS7{F0: 248}
        var v104 uint8 = 2
        for fi105 := 0; fi105 < 15; fi105++ {
            v104 = (v104 + 238)
            v106 := GenS15{F0: true}
            _ = fi105
            _ = v106
        }
        _ = v103
        _ = v104
    } else {
        var v107 uint32 = 4294967295
        var v108 uint32 = (v107 % (((v107 << ((484) & 63))) | 1))
        if (!(("foo bar" < "") && (!(true || false)))) {
            var v109 int8 = -1
            _ = v109
        } else {
            v107 -= (v107 | 2147483648)
            for fi110 := 0; fi110 < 6; fi110++ {
                var v111 uint8 = 42
                switch fi110 {
                case 0:
                    if true {
                        v3 = 73
                        v2 -= 2
                    }
                case 128:
                    fmt.Println(v107, v3)
                }
                _ = fi110
                _ = v111
            }
            for fi112 := 0; fi112 < 16; fi112++ {
                if (false && (!(4294967295 >= 774))) {
                    for fi113 := 0; fi113 < 19; fi113++ {
                        switch (fi113 + 32767) {
                        case 9223372036854775807:
                            v114 := GenS2{F0: ("foo bar" + ""), F1: 1, F2: 128}
                            fmt.Println(v107, fi113, v108)
                            if (-77 < 127) {
                                switch "" {
                                case "":
                                    switch "a" {
                                    case "":
                                        v107 = (v107 - (^(v107)))
                                    }
                                case "\n\t":
                                    switch "foo bar" {
                                    case "a":
                                        var v115 int64 = 1
                                        _ = v115
                                    case "\n\t":
                                        fmt.Println(v3, fi113)
                                    }
                                }
                            }
                            _ = v114
                        case 42:
                            switch "hello" {
                            case "\n\t":
                                v1 = append(v1, fi112)
                            case "a":
                                var v116 bool = (-128 <= 0)
                                _ = v116
                            }
                        case 613:
                            var v117 int8 = int8(fi113)
                            _ = v117
                        default:
                            var v118 uint32 = 6
                            _ = v118
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) > 0 {
                            v1[0] = (fi112 % (((fi112 % ((fi113) | 1))) | 1))
                        }
                        _ = fi113
                    }
                } else {
                    v3 = 42
                }
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        if (42 != 0) {
                            v119 := GenS16{F0: 255, F1: 1}
                            v107 = (v108 & v107)
                            var v120 string = "a"
                            _ = v119
                            _ = v120
                        }
                    case "":
                        v121 := GenS16{F0: 42, F1: 65535}
                        _ = v121
                    case "\n\t":
                        if ((^(v107)) >= (^(v108))) {
                            if false {
                                var v122 float64 = 0.1
                                _ = v122
                            } else {
                                if true {
                                    v2 -= 0
                                } else {
                                    v123 := GenS17{F0: 0, F1: v3}
                                    _ = v123
                                }
                                if ((fi112 / (((fi112 + fi112)) | 1)) != 4294967295) {
                                    if false {
                                        v124 := GenS7{F0: 127}
                                        var v125 uint16 = 2
                                        var v126 uint64 = v2
                                        _ = v124
                                        _ = v125
                                        _ = v126
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        v2 = (v2 ^ (v2 / (((^(v2))) | 1)))
                                    case "hello":
                                        var v127 int16 = int16(v3)
                                        _ = v127
                                    case "foo bar":
                                        v2 = (v2 >> ((65535) & 63))
                                    }
                                    v107 -= v107
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            v128 := GenS18{F0: (("hello" + "a") + ("\n\t" + "foo bar")), F1: 32767}
                            var v129 int = fi112
                            _ = v128
                            _ = v129
                        } else {
                            v108 = v107
                            for fi130 := 0; fi130 < 17; fi130++ {
                                var v131 float64 = 1.5
                                _ = fi130
                                _ = v131
                            }
                            v132 := GenS2{F0: "\n\t", F1: 0, F2: 101}
                            _ = v132
                        }
                    default:
                        var v133 int = fi112
                        _ = v133
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v108 -= v107
                    if (0.1 >= 0.1) {
                        var v134 uint64 = 18446744073709551615
                        _ = v134
                    } else {
                        v2 -= v2
                        switch (-(fi112)) {
                        case 32768:
                            var v135 int16 = 32767
                            _ = v135
                        case 30:
                            v136 := GenS19{F0: (-(v3)), F1: 100, F2: 1}
                            _ = v136
                        case 4294967295:
                            if (32767 >= 1000) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                v137 := GenS13{F0: v2, F1: 192, F2: ""}
                                _ = v137
                            }
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        fmt.Println(v2, v2, v108)
                    }
                }
                var v138 uint64 = (^(v2))
                _ = fi112
                _ = v138
            }
        }
        _ = v107
        _ = v108
    }
    v1 = append(v1, 32767)
    v3 -= (v3 * (-(v3)))
    v139 := GenS20{F0: 0}
    for fi140 := 0; fi140 < 15; fi140++ {
        var v141 float32 = float32(0.0)
        _ = fi140
        _ = v141
    }
    for fi142 := 0; fi142 < 7; fi142++ {
        _ = len(v1)
        _ = fi142
    }
    if true {
        switch "hello" {
        case "\n\t":
            var v143 uint16 = 32767
            _ = v143
        case "":
            v3 -= v3
        case "foo bar":
            for fi144 := 0; fi144 < 1; fi144++ {
                var v145 int64 = 9223372036854775807
                for fi146 := 0; fi146 < 2; fi146++ {
                    v147 := GenS10{F0: 496, F1: 65535}
                    if false {
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                if true {
                                    v2 = v2
                                }
                            }
                            if true {
                                fmt.Println(v2)
                                v1 = append(v1, (fi146 * -9223372036854775808))
                                v2 = 4294967295
                            } else {
                                v148 := GenS21{F0: 1}
                                var v149 uint32 = 2147483647
                                _ = v148
                                _ = v149
                            }
                            switch -499 {
                            case 2147483647:
                                var v150 uint64 = v2
                                _ = v150
                            }
                        } else {
                            if (!false) {
                                fmt.Println(fi146, v2)
                                var v151 uint64 = 42
                                v152 := GenS22{F0: (v145 & 659), F1: 0, F2: false}
                                _ = v151
                                _ = v152
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = (fi144 ^ fi146)
                                }
                                if true {
                                    if true {
                                        var v153 string = ("\n\t" + ("a" + "a"))
                                        _ = v153
                                    }
                                    v2 = (v2 + v2)
                                    v2 = (^(v2))
                                }
                                v3 -= v3
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v2 = v2
                        var v154 int64 = 42
                        if true {
                            switch "" {
                            case "hello":
                                v3 -= v3
                            }
                        }
                        fmt.Println(v145, v154, fi144)
                        _ = v154
                    }
                    _ = fi146
                    _ = v147
                }
                _ = fi144
                _ = v145
            }
        }
        if (31 > -128) {
            v155 := GenS23{F0: -9223372036854775808, F1: (true || (128 <= 32767)), F2: 128}
            if (!false) {
                if (65535 < 32767) {
                    switch -1 {
                    case 127:
                        v156 := GenS4{F0: v2, F1: v3}
                        _ = v156
                    default:
                        fmt.Println(v3)
                    }
                    if ((!true) && false) {
                        switch 42 {
                        case 9223372036854775807:
                            var v157 string = "\n\t"
                            _ = v157
                        default:
                            v3 += v3
                        }
                    }
                    if false {
                        fmt.Println(v3, v2)
                    }
                } else {
                    var v158 uint16 = 42
                    if false {
                        switch "\n\t" {
                        case "a":
                            var v159 int16 = -1
                            _ = v159
                        case "":
                            var v160 uint64 = 1
                            _ = v160
                        default:
                            v158 = (v158 % (((v158 | (v158 - v158))) | 1))
                        }
                    }
                    _ = v158
                }
                fmt.Println(v2)
            }
            _ = v155
        }
    }
    if (!true) {
        if (65535 <= 291) {
            var v161 string = ("" + "\n\t")
            switch v161 {
            case "\n\t":
                if (!((1 != -1) || true)) {
                    v162 := GenS7{F0: 128}
                    if false {
                        fmt.Println(v2, v2)
                    }
                    _ = v162
                } else {
                    v161 = "a"
                    fmt.Println(v2, v2)
                    fmt.Println(v2, v2, v3)
                }
            default:
                var v163 int64 = 42
                _ = v163
            }
            _ = v161
        } else {
            var v164 int16 = 0
            _ = v164
        }
        if len(v1) > 0 {
            v1[0] = 128
        }
        v165 := GenS24{F0: 1, F1: 767}
        _ = v165
    } else {
        v166 := GenS15{F0: true}
        v2 += v2
        fmt.Println(v3, v3)
        _ = v166
    }
    v1 = append(v1, 4294967295)
    switch -9223372036854775808 {
    case 2147483647:
        if len(v1) > 0 {
            v1[0] = 42
        }
    case 374:
        v2 += v2
    case 1:
        for fi167 := 0; fi167 < 12; fi167++ {
            v3 -= (v3 << ((65535) & 63))
            switch (fi167 | (fi167 << ((0) & 63))) {
            case -1:
                if len(v1) > 0 {
                    v1[0] = -263
                }
            }
            _ = fi167
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v168 := GenS25{F0: v3}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v139
    _ = v168
    _ = v2
    _ = v3
}
