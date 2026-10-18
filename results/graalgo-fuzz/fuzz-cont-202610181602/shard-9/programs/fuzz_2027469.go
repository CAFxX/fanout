package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint; F1 int32; F2 string }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 string; F1 uint64; F2 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 uint64; F2 int8 }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 bool; F1 uint64; F2 uint }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int; F1 uint64; F2 bool }
type GenS13 struct { F0 int8; F1 uint; F2 uint8 }
type GenS14 struct { F0 int; F1 int8 }
type GenS15 struct { F0 int; F1 int16; F2 uint16 }
type GenS16 struct { F0 uint16; F1 string; F2 int32 }
type GenS17 struct { F0 int32; F1 int16; F2 int32 }
type GenS18 struct { F0 int64; F1 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 255
    var v3 uint64 = 4294967295
    v4 := make(map[int]uint32)
    var v5 float64 = 1e10
    var v6 uint8 = (v2 / ((v2) | 1))
    v7 := v1[("a" + ("" + "hello"))]
    v8 := GenS1{F0: 372, F1: 2147483647}
    if false {
        var v9 int32 = -1
        v10 := []int{32768, 128}
        _ = v10
        _ = v9
    } else {
        var v11 uint8 = 0
        v12 := []int{0, 0}
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        if false {
            v6 -= v2
            if true {
                v3 = 1
            }
        }
        _ = fi13
    }
    switch -895 {
    case 32768:
        var v14 uint = 42
        _ = v14
    default:
        for fi15 := 0; fi15 < 17; fi15++ {
            if ((v7 / (((v7 - (v7 + v7))) | 1)) == 65535) {
                v16 := GenS1{F0: 65535, F1: -1}
                _ = v16
            }
            _ = fi15
        }
    }
    delete(v1, ("hello" + ""))
    switch 32768 {
    case 204:
        v7 = v7
    }
    switch "hello" {
    case "\n\t":
        var v17 int64 = 42
        _ = v17
    case "foo bar":
        v4[84] = v7
    case "a":
        if true {
            if (false && (((float32(v5) != float32(v5)) || (42 < 42)) && ("hello" != ("\n\t" + ((("" + "") + "\n\t") + ("foo bar" + "a")))))) {
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := GenS1{F0: 65535, F1: 488}
                    for fi20 := 0; fi20 < 9; fi20++ {
                        if (32768 == fi18) {
                            var v21 uint = 461
                            _ = v21
                        }
                        v22 := GenS1{F0: 2, F1: 0}
                        var v23 float64 = float64(0.0)
                        _ = fi20
                        _ = v22
                        _ = v23
                    }
                    _ = fi18
                    _ = v19
                }
                v3 = v3
                v6 = (^(v6))
            } else {
                if true {
                    v6 = (v2 / ((v2) | 1))
                    if false {
                        var v24 int64 = -9223372036854775808
                        if false {
                            var v25 uint8 = 0
                            _ = v25
                        }
                        if (-1 <= -71) {
                            var v26 float32 = (0.1 * (3.14159 - 3.14159))
                            if false {
                                if (-32768 != 42) {
                                    var v27 float64 = float64((v26 - v26))
                                    for fi28 := 0; fi28 < 17; fi28++ {
                                        var v29 int = fi28
                                        fmt.Println(v29, v27)
                                        _ = fi28
                                        _ = v29
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        v30 := GenS1{F0: uint16(v3), F1: 42}
                                        _ = v30
                                    case "":
                                        for fi31 := 0; fi31 < 6; fi31++ {
                                            v32 := GenS1{F0: 80, F1: 1}
                                            switch fi31 {
                                            case -664:
                                                var v33 string = ("foo bar" + "a")
                                                _ = v33
                                            }
                                            _ = fi31
                                            _ = v32
                                        }
                                    case "\n\t":
                                        var v34 int32 = -230
                                        _ = v34
                                    }
                                    _ = v27
                                }
                                var v35 uint = 2
                                var v36 int8 = -128
                                _ = v35
                                _ = v36
                            }
                            v37 := GenS1{F0: 845, F1: 346}
                            _ = v26
                            _ = v37
                        }
                        _ = v24
                    }
                    var v38 bool = (42 <= -1)
                    _ = v38
                } else {
                    if (!(!false)) {
                        v2 -= v2
                        var v39 int = -2147483648
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 1; fi40++ {
                        var v41 int16 = 42
                        v41 = (v41 * v41)
                        for fi42 := 0; fi42 < 6; fi42++ {
                            v41 = v41
                            if true {
                                v7 = (v7 << ((978) & 63))
                                _ = len(v1)
                                var v43 uint = 1
                                var v44 uint32 = (v7 + (v7 >> ((255) & 63)))
                                _ = v43
                                _ = v44
                            }
                            var v45 int = (fi40 + -865)
                            if ((161 == 896) && false) {
                                v46 := v4[fi40]
                                _ = v46
                            }
                            if false {
                                var v47 uint = 2
                                v4[(fi40 + (fi42 << ((1) & 63)))] = 533
                                switch (^(fi40)) {
                                case 767:
                                    switch v45 {
                                    case 32768:
                                        var v48 uint = v47
                                        _ = v48
                                    default:
                                        var v49 bool = true
                                        _ = v49
                                    }
                                case 127:
                                    var v50 uint = v47
                                    _ = v50
                                case 1:
                                    var v51 int64 = 1
                                    _ = v51
                                default:
                                    var v52 int64 = 42
                                    _ = v52
                                }
                                _ = v47
                            }
                            if (!false) {
                                var v53 int8 = -1
                                fmt.Println(v5)
                                _ = v53
                            } else {
                                v54 := []int{610}
                                var v55 float32 = -2.25
                                _ = v54
                                _ = v55
                            }
                            _ = fi42
                            _ = v45
                        }
                        _ = fi40
                        _ = v41
                    }
                }
            }
            var v56 float32 = ((0.1 * 1e10) / 1.0)
            _ = v56
        } else {
            switch -9223372036854775808 {
            case 523:
                v57 := []int{42, 4294967295, 474}
                _ = v57
            case 32768:
                fmt.Println(v2, v2, v2)
            case 4294967295:
                delete(v4, 127)
            default:
                var v58 float64 = float64(-2.25)
                _ = v58
            }
            if (!true) {
                fmt.Println(v6, v2)
            }
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                var v59 string = "\n\t"
                _ = v59
            case "foo bar":
                if true {
                    if ((v5 + v5) != ((1e10 / ((v5 * 3.14159) * v5)) - (((0.1 * (v5 + (v5 + 1.5))) * float64(-2.25)) * (v5 * 1e10)))) {
                        var v60 uint32 = 65535
                        switch ("a" + ("foo bar" + "hello")) {
                        case "":
                            v5 -= v5
                        default:
                            var v61 float32 = float32((float64(0.1) / ((3.14159 * 0.0) + 1.5)))
                            _ = v61
                        }
                        for fi62 := 0; fi62 < 2; fi62++ {
                            fmt.Println(v7, v7)
                            v63 := v1[("foo bar" + "\n\t")]
                            _ = fi62
                            _ = v63
                        }
                        _ = v60
                    } else {
                        var v64 int8 = 1
                        for fi65 := 0; fi65 < 15; fi65++ {
                            for fi66 := 0; fi66 < 3; fi66++ {
                                fmt.Println(v2)
                                _ = fi66
                            }
                            if true {
                                var v67 float32 = float32(v5)
                                for fi68 := 0; fi68 < 12; fi68++ {
                                    var v69 uint64 = v3
                                    var v70 int = fi65
                                    if (0 == -1) {
                                        fmt.Println(v5, v6)
                                        switch "foo bar" {
                                        case "":
                                            if (v6 != v2) {
                                                v71 := GenS2{F0: v6, F1: "hello"}
                                                _ = v71
                                            }
                                        case "\n\t":
                                            if true {
                                                v72 := v1[(("a" + "") + "")]
                                                v73 := []int{(fi68 ^ (v70 / ((fi65) | 1))), fi65, 127, (fi68 % ((-1) | 1))}
                                                delete(v1, "\n\t")
                                                _ = v72
                                                _ = v73
                                            }
                                        }
                                    }
                                    _ = fi68
                                    _ = v69
                                    _ = v70
                                }
                                for fi74 := 0; fi74 < 14; fi74++ {
                                    v75 := []int{(fi65 << ((830) & 63))}
                                    switch (("foo bar" + ("" + (("hello" + ("foo bar" + "\n\t")) + "hello"))) + ("" + (("a" + (("foo bar" + ("foo bar" + "hello")) + "foo bar")) + ("\n\t" + "a")))) {
                                    case "":
                                        var v76 uint16 = 0
                                        _ = v76
                                    }
                                    _ = fi74
                                    _ = v75
                                }
                                _ = v67
                            }
                            fmt.Println(fi65, v64)
                            _ = fi65
                        }
                        var v77 uint16 = 42
                        _ = v64
                        _ = v77
                    }
                }
            case "":
                for fi78 := 0; fi78 < 5; fi78++ {
                    for fi79 := 0; fi79 < 19; fi79++ {
                        if (((0 != 42) || (32767 != 473)) && true) {
                            v6 -= v6
                            var v80 int8 = -113
                            if false {
                                if (3.14159 != 0.1) {
                                    v2 += v6
                                    if (!(9223372036854775807 >= 0)) {
                                        if false {
                                            fmt.Println(v80)
                                            var v81 uint16 = 255
                                            v82 := []int{fi78, fi78, 526}
                                            _ = v81
                                            _ = v82
                                        } else {
                                            v5 += ((v5 * v5) * v5)
                                            v83 := GenS3{F0: 4294967295, F1: int32(fi79), F2: ("hello" + "")}
                                            if false {
                                                v2 += v6
                                                if false {
                                                    v84 := GenS2{F0: v6, F1: "a"}
                                                    _ = v84
                                                } else {
                                                    fmt.Println(v5, v6, v2)
                                                    v85 := GenS4{F0: 726, F1: 255}
                                                    v7 += 42
                                                    _ = v85
                                                }
                                            } else {
                                                if (!true) {
                                                    if ((((1.5 * float32(v5)) - float32(v5)) - 0.1) <= ((0.1 * 1.5) + 1e10)) {
                                                        switch "" {
                                                        case "foo bar":
                                                            v3 -= 2
                                                            fmt.Println(fi78, fi78)
                                                        }
                                                        switch 127 {
                                                        case 1:
                                                            v5 = float64(3.14159)
                                                        default:
                                                            v5 = v5
                                                        }
                                                        var v86 string = "a"
                                                        _ = v86
                                                    }
                                                    fmt.Println(v3, v2, v5)
                                                }
                                                switch (fi78 << ((42) & 63)) {
                                                case -9223372036854775808:
                                                    var v87 float64 = v5
                                                    _ = v87
                                                default:
                                                    var v88 int8 = (^(v80))
                                                    _ = v88
                                                }
                                                switch (("\n\t" + ("\n\t" + ("a" + ""))) + (("" + "hello") + "\n\t")) {
                                                case "a":
                                                    fmt.Println(fi79)
                                                case "foo bar":
                                                    switch ("a" + "") {
                                                    case "a":
                                                        v5 = ((((3.14159 - 1.5) / v5) - (float64(3.14159) - v5)) / float64(float32(v5)))
                                                    case "foo bar":
                                                        fmt.Println(v5)
                                                    default:
                                                        v89 := []int{}
                                                        switch "" {
                                                        case "\n\t":
                                                            v2 = v2
                                                        case "foo bar":
                                                            v90 := GenS3{F0: 255, F1: -2147483648, F2: (("\n\t" + "hello") + "hello")}
                                                            _ = v90
                                                        case "a":
                                                            v6 = (v2 << ((uint(v3)) & 63))
                                                        default:
                                                            var v91 uint16 = 65535
                                                            _ = v91
                                                        }
                                                        v92 := GenS1{F0: 65535, F1: int32(v7)}
                                                        _ = v89
                                                        _ = v92
                                                    }
                                                default:
                                                    if (670 >= -2147483648) {
                                                        if false {
                                                            fmt.Println(v3, v5)
                                                            v7 = (v7 - 407)
                                                            var v93 uint64 = (v3 / ((v3) | 1))
                                                            fmt.Println(v93)
                                                            v94 := GenS5{F0: "a", F1: 740, F2: 0}
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                    }
                                                }
                                            }
                                            var v95 int = 32768
                                            _ = v83
                                            _ = v95
                                        }
                                        v96 := GenS6{F0: ""}
                                        _ = v96
                                    } else {
                                        v97 := GenS3{F0: 564, F1: -1, F2: "\n\t"}
                                        _ = v97
                                    }
                                    v80 -= -51
                                }
                            } else {
                                if false {
                                    fmt.Println(v80)
                                    if (v3 == v3) {
                                        v98 := GenS7{F0: 1, F1: 18446744073709551615, F2: (v80 + v80)}
                                        _ = v98
                                    }
                                } else {
                                    v5 = v5
                                }
                                var v99 uint16 = 0
                                delete(v1, ("\n\t" + ((("a" + "hello") + ("hello" + "a")) + ("foo bar" + ((("hello" + ("" + (("\n\t" + (("a" + "") + "")) + "\n\t"))) + ("a" + "hello")) + ("\n\t" + ("\n\t" + ("foo bar" + ""))))))))
                                _ = v99
                            }
                            _ = v80
                        }
                        v100 := []int{4294967295, (^(fi78)), -9223372036854775808, fi79}
                        switch ("a" + "") {
                        case "hello":
                            var v101 int8 = -1
                            _ = v101
                        }
                        _ = fi79
                        _ = v100
                    }
                    for fi102 := 0; fi102 < 13; fi102++ {
                        v7 -= 4294967295
                        v103 := []int{(^(fi102)), fi78, (fi102 * fi78)}
                        v104 := GenS8{F0: v7, F1: false}
                        var v105 uint32 = uint32(v2)
                        var v106 bool = ((53 == (v7 << ((18446744073709551615) & 63))) || (32767 != -32768))
                        fmt.Println(v6, v3, v105)
                        _ = fi102
                        _ = v103
                        _ = v104
                        _ = v105
                        _ = v106
                    }
                    if (true || true) {
                        switch (fi78 / ((fi78) | 1)) {
                        case 128:
                            if true {
                                fmt.Println(v3)
                            } else {
                                var v107 uint32 = v7
                                v2 -= 255
                                v108 := []int{(-(fi78)), -2147483648, (fi78 << ((255) & 63))}
                                _ = v107
                                _ = v108
                            }
                        }
                        if false {
                            v7 = v7
                        } else {
                            v2 += v2
                            if false {
                                var v109 int32 = 54
                                switch "" {
                                case "a":
                                    if true {
                                        if false {
                                            switch (fi78 + 2147483647) {
                                            case 40:
                                                if (-87 <= 127) {
                                                    for fi110 := 0; fi110 < 9; fi110++ {
                                                        v111 := []int{fi110, (fi110 * fi78), fi110}
                                                        v3 = 260
                                                        var v112 int8 = int8(v7)
                                                        _ = fi110
                                                        _ = v111
                                                        _ = v112
                                                    }
                                                } else {
                                                    for fi113 := 0; fi113 < 18; fi113++ {
                                                        v6 = 255
                                                        switch (fi78 >> ((4294967295) & 63)) {
                                                        case -1:
                                                            fmt.Println(v5)
                                                        case 32768:
                                                            v5 += (v5 - (v5 * float64((1e10 / 1e10))))
                                                        }
                                                        v114 := []int{0, (^(fi113)), (fi78 * -296)}
                                                        _ = fi113
                                                        _ = v114
                                                    }
                                                    var v115 int64 = 0
                                                    _ = v115
                                                }
                                            case 4294967295:
                                                v109 = (v109 - 1)
                                            case 127:
                                                v6 = v6
                                            }
                                            for fi116 := 0; fi116 < 8; fi116++ {
                                                v117 := v1["a"]
                                                if true {
                                                    if ("a" == ("foo bar" + "hello")) {
                                                        v118 := GenS5{F0: ("hello" + "\n\t"), F1: v3, F2: 660}
                                                        var v119 int8 = 127
                                                        if true {
                                                            v109 -= (v109 | (v109 % ((-730) | 1)))
                                                            if true {
                                                                v120 := []int{}
                                                                _ = v120
                                                            } else {
                                                                v1["a"] = 579
                                                            }
                                                        }
                                                        _ = v118
                                                        _ = v119
                                                    } else {
                                                        if false {
                                                            v5 = v5
                                                            var v121 float32 = float32(0.0)
                                                            v122 := GenS5{F0: (("\n\t" + ("\n\t" + "\n\t")) + (("" + "foo bar") + ("" + "\n\t"))), F1: v3, F2: (v109 | v109)}
                                                            _ = v121
                                                            _ = v122
                                                        } else {
                                                            switch (fi78 + fi78) {
                                                            case 573:
                                                                if true {
                                                                    if (v5 == v5) {
                                                                        switch "hello" {
                                                                        case "a":
                                                                            if (("foo bar" <= ("hello" + "foo bar")) || true) {
                                                                                var v123 bool = false
                                                                                var v124 bool = (v123 && (297 < 1))
                                                                                switch (fi116 / ((int(v3)) | 1)) {
                                                                                case -1:
                                                                                    var v125 int8 = int8(v109)
                                                                                    _ = v125
                                                                                case 32768:
                                                                                    switch ("hello" + "foo bar") {
                                                                                    case "a":
                                                                                        switch (fi116 * (fi116 % (((fi78 * (fi116 ^ fi116))) | 1))) {
                                                                                        case -871:
                                                                                            var v126 uint8 = v2
                                                                                            _ = v126
                                                                                        case 127:
                                                                                            if false {
                                                                                                switch (fi116 * fi116) {
                                                                                                case 32767:
                                                                                                    switch "foo bar" {
                                                                                                    case "\n\t":
                                                                                                        var v127 uint32 = v117
                                                                                                        _ = v127
                                                                                                    case "a":
                                                                                                        if (fi116 != fi116) {
                                                                                                            var v128 uint32 = (^(v7))
                                                                                                            _ = v128
                                                                                                        } else {
                                                                                                            fmt.Println(v7, fi78, v7)
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        delete(v4, (fi116 ^ 581))
                                                                                                    }
                                                                                                }
                                                                                                v129 := GenS9{F0: v124, F1: v3, F2: 42}
                                                                                                v130 := []int{(^(fi78)), 318, 100, (^(fi78))}
                                                                                                _ = v129
                                                                                                _ = v130
                                                                                            }
                                                                                        case -2147483648:
                                                                                            if v124 {
                                                                                                var v131 int = (fi78 >> ((423) & 63))
                                                                                                _ = v131
                                                                                            }
                                                                                        default:
                                                                                            var v132 string = "hello"
                                                                                            _ = v132
                                                                                        }
                                                                                    case "foo bar":
                                                                                        switch ("\n\t" + "") {
                                                                                        case "":
                                                                                            _ = len(v4)
                                                                                        case "\n\t":
                                                                                            var v133 bool = (!(!(!((false || v123) && (-1 <= -1)))))
                                                                                            _ = v133
                                                                                        case "foo bar":
                                                                                            v134 := GenS10{F0: 2147483647}
                                                                                            fmt.Println(v3, fi78, v7)
                                                                                            v109 = v109
                                                                                            _ = v134
                                                                                        }
                                                                                    case "\n\t":
                                                                                        if (v2 >= (v6 % ((0) | 1))) {
                                                                                            v109 += 42
                                                                                            v2 += 2
                                                                                            var v135 bool = (!v124)
                                                                                            _ = v135
                                                                                        } else {
                                                                                            var v136 float64 = ((v5 - (v5 / 1.0)) + (v5 / v5))
                                                                                            _ = v136
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v137 float64 = (3.14159 + 0.0)
                                                                                switch (fi116 | (fi78 << ((357) & 63))) {
                                                                                case 980:
                                                                                    var v138 bool = v124
                                                                                    _ = v138
                                                                                }
                                                                                _ = v123
                                                                                _ = v124
                                                                                _ = v137
                                                                            } else {
                                                                                var v139 int8 = 42
                                                                                v3 += 2
                                                                                _ = v139
                                                                            }
                                                                            v140 := []int{}
                                                                            _ = v140
                                                                        case "\n\t":
                                                                            if (true || (0 <= 100)) {
                                                                                v141 := GenS5{F0: "", F1: (v3 >> ((625) & 63)), F2: -1}
                                                                                if false {
                                                                                    v142 := GenS4{F0: 1000, F1: 1}
                                                                                    v5 += float64(3.14159)
                                                                                    _ = v142
                                                                                }
                                                                                v143 := GenS11{F0: ""}
                                                                                _ = v141
                                                                                _ = v143
                                                                            }
                                                                            fmt.Println(v5, v117)
                                                                        default:
                                                                            v144 := []int{}
                                                                            _ = v144
                                                                        }
                                                                        v6 += (v6 & 128)
                                                                        fmt.Println(fi78, v2, v109)
                                                                        if (true && false) {
                                                                            if (!false) {
                                                                                switch 32767 {
                                                                                case 128:
                                                                                    if false {
                                                                                        var v145 float64 = 3.14159
                                                                                        _ = v145
                                                                                    } else {
                                                                                        v7 = 51
                                                                                    }
                                                                                case 127:
                                                                                    var v146 int8 = 1
                                                                                    _ = v146
                                                                                }
                                                                                if false {
                                                                                    if (!false) {
                                                                                        switch (fi78 % (((fi78 >> ((42) & 63))) | 1)) {
                                                                                        case 42:
                                                                                            if (-1 < int16(v109)) {
                                                                                                switch fi116 {
                                                                                                case 792:
                                                                                                    if true {
                                                                                                        v117 += v7
                                                                                                        switch (^(fi78)) {
                                                                                                        case 9223372036854775807:
                                                                                                            switch fi116 {
                                                                                                            case 128:
                                                                                                                v147 := GenS12{F0: (-(fi116)), F1: 677, F2: true}
                                                                                                                _ = v147
                                                                                                            case -733:
                                                                                                                var v148 int32 = v109
                                                                                                                _ = v148
                                                                                                            }
                                                                                                        case -19:
                                                                                                            var v149 int32 = (v109 >> ((1) & 63))
                                                                                                            _ = v149
                                                                                                        case 127:
                                                                                                            v109 += (-(v109))
                                                                                                            v150 := []int{fi116, fi78, fi78, fi116}
                                                                                                            _ = v150
                                                                                                        default:
                                                                                                            v117 -= 2147483648
                                                                                                        }
                                                                                                        delete(v1, ("hello" + "hello"))
                                                                                                    }
                                                                                                case 417:
                                                                                                    switch 620 {
                                                                                                    case 128:
                                                                                                        if false {
                                                                                                            v109 = (v109 >> ((996) & 63))
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                v109 = 1
                                                                                                                switch (fi116 | fi78) {
                                                                                                                case 128:
                                                                                                                    if false {
                                                                                                                        v4[(fi116 + (^(fi116)))] = 2147483647
                                                                                                                        var v151 uint16 = 726
                                                                                                                        v152 := []int{fi78, (fi116 * fi78), (fi116 & fi116)}
                                                                                                                        _ = v151
                                                                                                                        _ = v152
                                                                                                                    }
                                                                                                                case 4294967295:
                                                                                                                    var v153 int64 = 9223372036854775807
                                                                                                                    _ = v153
                                                                                                                default:
                                                                                                                    switch int(v117) {
                                                                                                                    case 2147483647:
                                                                                                                        var v154 int = -2147483648
                                                                                                                        _ = v154
                                                                                                                    default:
                                                                                                                        var v155 uint64 = (^(v3))
                                                                                                                        _ = v155
                                                                                                                    }
                                                                                                                }
                                                                                                                var v156 int = fi78
                                                                                                                _ = v156
                                                                                                            }
                                                                                                            v7 -= v7
                                                                                                            v157 := GenS3{F0: 528, F1: (v109 >> ((493) & 63)), F2: ""}
                                                                                                            switch (fi78 % ((int(v6)) | 1)) {
                                                                                                            case -239:
                                                                                                                var v158 uint = 0
                                                                                                                _ = v158
                                                                                                            case 9223372036854775807:
                                                                                                                if false {
                                                                                                                    fmt.Println(v2, v2, v109)
                                                                                                                } else {
                                                                                                                    v159 := GenS13{F0: -107, F1: 1, F2: (v6 * 2)}
                                                                                                                    v3 = 368
                                                                                                                    fmt.Println(v109)
                                                                                                                    var v160 int = (fi78 << ((255) & 63))
                                                                                                                    _ = v159
                                                                                                                    _ = v160
                                                                                                                }
                                                                                                            }
                                                                                                            if ((uint8(v3) <= 255) && true) {
                                                                                                                fmt.Println(v7, v2, v3)
                                                                                                                switch "" {
                                                                                                                case "":
                                                                                                                    if (false || true) {
                                                                                                                        var v161 int64 = 0
                                                                                                                        if (v6 <= v2) {
                                                                                                                            fmt.Println(fi116, v109)
                                                                                                                            v162 := []int{42}
                                                                                                                            fmt.Println(v109, v3)
                                                                                                                            _ = v162
                                                                                                                        }
                                                                                                                        v7 = 2147483647
                                                                                                                        _ = v161
                                                                                                                    } else {
                                                                                                                        v7 = v117
                                                                                                                        var v163 uint8 = v6
                                                                                                                        v5 = (v5 - 0.1)
                                                                                                                        _ = v163
                                                                                                                    }
                                                                                                                default:
                                                                                                                    fmt.Println(fi78, v3)
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v157
                                                                                                        }
                                                                                                    case 32767:
                                                                                                        v5 += float64((((float32(1e10) / -2.25) + 0.1) * (1e10 / -2.25)))
                                                                                                    case -830:
                                                                                                        var v164 int8 = 0
                                                                                                        _ = v164
                                                                                                    }
                                                                                                }
                                                                                                if ((^(v3)) <= v3) {
                                                                                                    v165 := v4[-9223372036854775808]
                                                                                                    _ = v165
                                                                                                } else {
                                                                                                    var v166 int32 = v109
                                                                                                    v5 -= 0.0
                                                                                                    var v167 int64 = 384
                                                                                                    _ = v166
                                                                                                    _ = v167
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        switch (fi78 ^ (fi78 * (fi78 * fi78))) {
                                                                                        case 0:
                                                                                            var v168 float64 = v5
                                                                                            _ = v168
                                                                                        case 442:
                                                                                            if (127 <= (v6 / (((v6 << ((0) & 63))) | 1))) {
                                                                                                v6 -= 127
                                                                                            }
                                                                                        case 9223372036854775807:
                                                                                            if true {
                                                                                                var v169 string = "a"
                                                                                                _ = v169
                                                                                            } else {
                                                                                                if false {
                                                                                                    var v170 int64 = 1
                                                                                                    if true {
                                                                                                        switch "\n\t" {
                                                                                                        case "\n\t":
                                                                                                            var v171 int = 4294967295
                                                                                                            _ = v171
                                                                                                        case "a":
                                                                                                            fmt.Println(v109)
                                                                                                        default:
                                                                                                            switch 1 {
                                                                                                            case 42:
                                                                                                                var v172 int64 = v170
                                                                                                                _ = v172
                                                                                                            case 9223372036854775807:
                                                                                                                v7 += (^(v117))
                                                                                                            }
                                                                                                        }
                                                                                                        var v173 bool = (-100 <= 353)
                                                                                                        if v173 {
                                                                                                            v3 = 1
                                                                                                            if (!((v109 * v109) < (v109 / ((-2147483648) | 1)))) {
                                                                                                                var v174 uint64 = 18446744073709551615
                                                                                                                var v175 int = (fi78 << ((255) & 63))
                                                                                                                _ = v174
                                                                                                                _ = v175
                                                                                                            }
                                                                                                            _ = len(v4)
                                                                                                            if (fi116 <= (fi116 - (fi78 * fi78))) {
                                                                                                                var v176 string = "a"
                                                                                                                _ = v176
                                                                                                            } else {
                                                                                                                switch "hello" {
                                                                                                                case "\n\t":
                                                                                                                    v177 := []int{(^(fi116))}
                                                                                                                    _ = v177
                                                                                                                }
                                                                                                                v170 = v170
                                                                                                                v3 -= (v3 + v3)
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v173
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            v6 -= (v6 >> ((65535) & 63))
                                                                                                            v178 := []int{fi78, fi78, fi116, -32}
                                                                                                            _ = v178
                                                                                                        }
                                                                                                        v117 = v7
                                                                                                    }
                                                                                                    _ = v170
                                                                                                } else {
                                                                                                    v179 := GenS14{F0: (fi116 % ((127) | 1)), F1: 7}
                                                                                                    _ = v179
                                                                                                }
                                                                                                v6 += (v2 >> ((2) & 63))
                                                                                            }
                                                                                        }
                                                                                        v6 = v2
                                                                                        v180 := []int{}
                                                                                        _ = v180
                                                                                    }
                                                                                }
                                                                                fmt.Println(v2, v7)
                                                                            }
                                                                            v5 = v5
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                var v181 int = -2147483648
                                                                                _ = v181
                                                                            case "hello":
                                                                                var v182 uint8 = (v2 + v6)
                                                                                _ = v182
                                                                            case "a":
                                                                                var v183 uint32 = 1
                                                                                _ = v183
                                                                            default:
                                                                                v184 := GenS15{F0: (-(fi116)), F1: 1000, F2: 255}
                                                                                _ = v184
                                                                            }
                                                                        } else {
                                                                            v185 := []int{}
                                                                            var v186 int16 = 1
                                                                            v5 = v5
                                                                            v2 -= (v2 % ((1) | 1))
                                                                            v186 -= -846
                                                                            _ = v185
                                                                            _ = v186
                                                                        }
                                                                        delete(v4, (fi78 + (fi78 >> ((625) & 63))))
                                                                        var v187 uint64 = v3
                                                                        if true {
                                                                            var v188 bool = ((^(fi116)) != fi116)
                                                                            fmt.Println(v188, v117, v6)
                                                                            _ = v188
                                                                        } else {
                                                                            var v189 uint = 558
                                                                            if (float64(3.14159) > ((1.5 - 0.0) * v5)) {
                                                                                if (false && ((^(v6)) != 42)) {
                                                                                    v3 = 4294967295
                                                                                    v190 := []int{42, (-(fi116)), (fi78 << ((65535) & 63)), (fi78 & int(v6))}
                                                                                    v187 -= uint64(v7)
                                                                                    _ = v190
                                                                                } else {
                                                                                    v191 := v1["foo bar"]
                                                                                    _ = v191
                                                                                }
                                                                                v192 := GenS15{F0: (fi78 | (fi78 + fi116)), F1: -1, F2: 255}
                                                                                _ = v192
                                                                            }
                                                                            _ = v189
                                                                        }
                                                                        v3 += 1
                                                                        var v193 uint = 255
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            var v194 int16 = 383
                                                                            var v195 bool = ((v194 >> (((v193 + (^(v193)))) & 63)) != 1)
                                                                            _ = v194
                                                                            _ = v195
                                                                        case "a":
                                                                            var v196 uint64 = v187
                                                                            var v197 bool = (true && true)
                                                                            _ = v196
                                                                            _ = v197
                                                                        }
                                                                        v5 += v5
                                                                        v187 -= v187
                                                                        _ = v187
                                                                        _ = v193
                                                                    }
                                                                    v198 := GenS12{F0: 723, F1: 1, F2: (0 < 42)}
                                                                    _ = v198
                                                                } else {
                                                                    v109 = v109
                                                                }
                                                            }
                                                            v4[(fi78 * fi78)] = (v7 + v7)
                                                            var v199 int64 = 1
                                                            _ = v199
                                                        }
                                                        switch "\n\t" {
                                                        case "":
                                                            v200 := GenS3{F0: 2, F1: v109, F2: ("a" + "hello")}
                                                            _ = v200
                                                        }
                                                    }
                                                    v2 += (v6 - (v6 - 43))
                                                    v201 := GenS12{F0: (fi116 ^ fi78), F1: (v3 << ((4294967295) & 63)), F2: (42 != 42)}
                                                    _ = v201
                                                } else {
                                                    v202 := []int{(fi78 + (fi116 % (((fi78 & 4294967295)) | 1))), 409}
                                                    v203 := GenS8{F0: (^(v117)), F1: ((((0.0 - -2.25) != -2.25) && false) || (false || (fi78 != -9223372036854775808)))}
                                                    v7 = 2
                                                    _ = v202
                                                    _ = v203
                                                }
                                                var v204 int16 = -32768
                                                _ = fi116
                                                _ = v117
                                                _ = v204
                                            }
                                        }
                                        v5 = v5
                                    }
                                default:
                                    v205 := []int{128, fi78}
                                    _ = v205
                                }
                                _ = v109
                            }
                        }
                    } else {
                        for fi206 := 0; fi206 < 4; fi206++ {
                            v6 += v2
                            v2 = (v2 % (((^(v2))) | 1))
                            v207 := GenS5{F0: "hello", F1: v3, F2: 1}
                            _ = fi206
                            _ = v207
                        }
                        fmt.Println(v2, v7)
                        if true {
                            for fi208 := 0; fi208 < 20; fi208++ {
                                v5 = 3.14159
                                switch ("a" + "\n\t") {
                                case "a":
                                    v6 -= 255
                                case "":
                                    v1[""] = (v7 - 668)
                                case "\n\t":
                                    var v209 float32 = float32(v5)
                                    _ = v209
                                default:
                                    v7 += v7
                                }
                                if false {
                                    var v210 uint64 = (v3 >> ((18446744073709551615) & 63))
                                    _ = v210
                                } else {
                                    if true {
                                        v211 := GenS16{F0: 2, F1: "hello", F2: 2147483647}
                                        v212 := GenS17{F0: -1, F1: 42, F2: 993}
                                        v213 := v4[351]
                                        _ = v211
                                        _ = v212
                                        _ = v213
                                    } else {
                                        if (!(false || true)) {
                                            v3 += v3
                                            var v214 bool = (9223372036854775807 != -9223372036854775808)
                                            var v215 uint8 = v2
                                            _ = v214
                                            _ = v215
                                        } else {
                                            switch (fi208 - (fi78 + -1)) {
                                            case 0:
                                                if (true || false) {
                                                    var v216 uint32 = 713
                                                    switch (("" + "a") + ("\n\t" + ("" + ""))) {
                                                    case "\n\t":
                                                        fmt.Println(v5, fi78, fi78)
                                                    case "foo bar":
                                                        if false {
                                                            v217 := GenS15{F0: fi208, F1: 42, F2: 1}
                                                            _ = v217
                                                        } else {
                                                            v5 = v5
                                                            var v218 uint64 = (v3 / ((4294967295) | 1))
                                                            var v219 int32 = -1
                                                            if true {
                                                                if (((v2 - v6) <= 128) && true) {
                                                                    v219 = v219
                                                                } else {
                                                                    fmt.Println(v218, v216, v219)
                                                                    v220 := GenS14{F0: fi208, F1: 42}
                                                                    _ = v220
                                                                }
                                                                v221 := GenS3{F0: 0, F1: v219, F2: ""}
                                                                _ = v221
                                                            }
                                                            switch 127 {
                                                            case 32768:
                                                                if (!(((0.0 - ((float32(1.5) + -2.25) - float32((0.1 + 0.0)))) > 3.14159) && false)) {
                                                                    v2 = v2
                                                                }
                                                            }
                                                            _ = v218
                                                            _ = v219
                                                        }
                                                    }
                                                    v222 := GenS18{F0: 1, F1: (^(v216))}
                                                    _ = v216
                                                    _ = v222
                                                } else {
                                                    if (fi208 >= (fi208 - (fi208 * (fi78 + -21)))) {
                                                        if true {
                                                            v223 := GenS1{F0: 1, F1: -2147483648}
                                                            var v224 int16 = -1
                                                            v225 := GenS5{F0: "foo bar", F1: 2, F2: -2147483648}
                                                            v226 := []int{fi208, fi78, (fi208 & 127), fi78}
                                                            v2 += 2
                                                            _ = v223
                                                            _ = v224
                                                            _ = v225
                                                            _ = v226
                                                        }
                                                        v3 = v3
                                                        var v227 uint8 = 0
                                                        _ = v227
                                                    }
                                                    v6 -= (v6 & (v6 / ((127) | 1)))
                                                }
                                            }
                                        }
                                        if (v3 < 2) {
                                            var v228 int = (^(fi208))
                                            _ = v228
                                        }
                                        fmt.Println(fi78, v3)
                                    }
                                    v229 := GenS2{F0: v2, F1: ((("" + ((("hello" + "a") + "") + ("" + ("\n\t" + "\n\t")))) + ("" + ((("foo bar" + (("hello" + "") + "\n\t")) + ("" + "a")) + "hello"))) + "hello")}
                                    var v230 uint64 = v3
                                    switch fi78 {
                                    case 0:
                                        delete(v4, fi208)
                                    default:
                                        if ((255 >= v6) || (((v5 / (float64(0.1) / v5)) >= v5) || true)) {
                                            v231 := []int{}
                                            _ = v231
                                        } else {
                                            v5 += v5
                                            v232 := []int{(fi78 / ((fi78) | 1)), (fi78 << ((708) & 63))}
                                            delete(v4, fi208)
                                            _ = v232
                                        }
                                    }
                                    _ = v229
                                    _ = v230
                                }
                                v3 = v3
                                _ = len(v4)
                                v6 += v2
                                _ = fi208
                            }
                            _ = len(v4)
                            fmt.Println(v2, v6, v2)
                        } else {
                            fmt.Println(v7, v5, v3)
                            for fi233 := 0; fi233 < 20; fi233++ {
                                var v234 float32 = (3.14159 / (0.0 + float32(((v5 - 0.0) * v5))))
                                _ = fi233
                                _ = v234
                            }
                        }
                    }
                    _ = fi78
                }
            }
        }
    default:
        fmt.Println(v3)
    }
    for fi235 := 0; fi235 < 7; fi235++ {
        v7 = (v7 - (v7 / (((v7 - v7)) | 1)))
        v236 := []int{fi235, 2147483647, fi235}
        _ = fi235
        _ = v236
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
