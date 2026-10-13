package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int32; F2 uint8 }
type GenS4 struct { F0 int32; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 string; F1 uint64; F2 int }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint; F1 string; F2 int8 }
type GenS12 struct { F0 bool; F1 string }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint; F1 int64 }
type GenS16 struct { F0 uint; F1 int64 }
type GenS17 struct { F0 int; F1 uint8; F2 bool }
type GenS18 struct { F0 uint; F1 string; F2 uint16 }
type GenS19 struct { F0 uint64; F1 uint32; F2 uint }
type GenS20 struct { F0 int8; F1 uint }
type GenS21 struct { F0 int }
type GenS22 struct { F0 int8; F1 int; F2 uint8 }
func genFn0() int8 {
    if true {
        var v7 int64 = 42
        _ = v7
    }
    if ((!true) || ("foo bar" > ("\n\t" + ("foo bar" + "")))) {
        if (4294967295 < 9223372036854775807) {
            switch ("hello" + "") {
            case "hello":
                var v8 uint8 = 128
                _ = v8
            case "foo bar":
                v9 := GenS1{F0: 2147483647}
                _ = v9
            case "\n\t":
                v10 := GenS1{F0: 1}
                _ = v10
            }
            v11 := GenS1{F0: 2147483647}
            v12 := []int{520}
            _ = v11
            _ = v12
        } else {
            switch "a" {
            case "foo bar":
                for fi13 := 0; fi13 < 18; fi13++ {
                    if ((!false) && (42 <= 2)) {
                        for fi14 := 0; fi14 < 8; fi14++ {
                            var v15 int8 = -128
                            _ = fi14
                            _ = v15
                        }
                        var v16 bool = (1 <= 176)
                        _ = v16
                    }
                    _ = fi13
                }
            case "hello":
                var v17 float32 = 0.1
                _ = v17
            case "a":
                var v18 uint = 42
                _ = v18
            default:
                switch 1 {
                case 1:
                    var v19 int32 = 0
                    _ = v19
                case 4294967295:
                    var v20 int16 = 1000
                    _ = v20
                }
            }
            switch "\n\t" {
            case "hello":
                v21 := GenS2{F0: -853}
                _ = v21
            case "foo bar":
                v22 := []int{}
                _ = v22
            case "\n\t":
                var v23 uint8 = 57
                _ = v23
            }
        }
        if (!(970 == 42)) {
            for fi24 := 0; fi24 < 12; fi24++ {
                v25 := GenS3{F0: (fi24 * (fi24 & 42)), F1: 0, F2: 50}
                var v26 int = (fi24 & 127)
                _ = fi24
                _ = v25
                _ = v26
            }
            for fi27 := 0; fi27 < 11; fi27++ {
                var v28 bool = ((0.1 + 1e10) < 0.1)
                var v29 bool = false
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        v30 := []int{4294967295, -876, -966, 1}
        _ = v30
    }
    var v31 uint8 = 43
    switch -823 {
    case 1:
        fmt.Println(v31)
    case 9223372036854775807:
        if (true && true) {
            fmt.Println(v31)
        } else {
            switch 127 {
            case 357:
                var v32 float64 = -2.25
                _ = v32
            default:
                v31 += v31
            }
        }
    case 32767:
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 int8 = -128
            v34 = v34
            v35 := GenS4{F0: -1, F1: (v34 * 45), F2: v31}
            _ = fi33
            _ = v34
            _ = v35
        }
    default:
        fmt.Println(v31)
    }
    _ = v31
    return -1
}
func genFn1(p0 int64, p1 int64, p2 int64) uint16 {
    for fi36 := 0; fi36 < 5; fi36++ {
        for fi37 := 0; fi37 < 18; fi37++ {
            if false {
                var v38 int8 = -1
                p2 = int64(v38)
                _ = v38
            }
            _ = fi37
        }
        v39 := []int{}
        _ = fi36
        _ = v39
    }
    fmt.Println(p0, p2)
    if (1 == 78) {
        v40 := []int{51, 32768, 127}
        _ = v40
    } else {
        fmt.Println(p0, p2, p1)
        v41 := GenS5{F0: 208}
        _ = v41
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := []string{"a"}
    v2 := []int{550}
    v3 := make(map[string]bool)
    var v4 float64 = (0.0 + 1e10)
    var v5 float32 = 0.0
    var v6 int = 4294967295
    v3[("" + "foo bar")] = ((v6 << ((255) & 63)) != 32767)
    if ((1e10 > (((float32(1e10) * v5) * v5) + (v5 * (v5 + float32(float64(v5)))))) || true) {
        v5 = 1e10
        for fi42 := 0; fi42 < 1; fi42++ {
            var v43 int16 = -32768
            if true {
                switch ("foo bar" + ("foo bar" + "")) {
                case "foo bar":
                    switch fi42 {
                    case 487:
                        v44 := GenS2{F0: (^(fi42))}
                        _ = v44
                    }
                case "\n\t":
                    if (!(712 != 65535)) {
                        v45 := GenS1{F0: 1}
                        var v46 uint64 = 913
                        var v47 int64 = 1
                        _ = v45
                        _ = v46
                        _ = v47
                    } else {
                        for fi48 := 0; fi48 < 14; fi48++ {
                            v1 = append(v1, (("a" + ("foo bar" + "hello")) + ""))
                            v43 += v43
                            var v49 bool = true
                            v4 -= v4
                            v50 := GenS1{F0: 0}
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                        var v51 int32 = -908
                        if len(v2) > 0 {
                            v2[0] = (v6 / (((v6 >> ((42) & 63))) | 1))
                        }
                        _ = v51
                    }
                }
                v52 := GenS4{F0: 207, F1: 127, F2: 2}
                v3["hello"] = true
                _ = v52
            }
            var v53 bool = false
            _ = fi42
            _ = v43
            _ = v53
        }
        switch ((("foo bar" + "hello") + "hello") + "\n\t") {
        case "hello":
            var v54 uint64 = 2
            _ = v54
        default:
            var v55 uint8 = 127
            _ = v55
        }
    }
    fmt.Println(v5, v5, v5)
    switch v6 {
    case 507:
        var v56 string = ("a" + "a")
        _ = v56
    case 9223372036854775807:
        var v57 bool = true
        _ = v57
    case 32768:
        for fi58 := 0; fi58 < 6; fi58++ {
            if true {
                var v59 int16 = 32767
                fmt.Println(v59)
                _ = v59
            } else {
                delete(v3, ("foo bar" + "\n\t"))
                v6 = (v6 << ((1) & 63))
            }
            v60 := GenS6{F0: 1, F1: 128}
            for fi61 := 0; fi61 < 6; fi61++ {
                v6 -= v6
                _ = fi61
            }
            _ = fi58
            _ = v60
        }
    default:
        v5 = ((-2.25 * 0.0) - (v5 - ((float32(1e10) / (0.1 * 1.5)) / ((0.0 - 0.0) * v5))))
    }
    for fi62 := 0; fi62 < 5; fi62++ {
        for fi63 := 0; fi63 < 13; fi63++ {
            v64 := GenS6{F0: 0, F1: 42}
            fmt.Println(v6, fi63)
            _ = fi63
            _ = v64
        }
        _ = fi62
    }
    if (((true && (-279 > -77)) || true) || true) {
        var v65 int8 = genFn0()
        v65 -= (v65 + (v65 * (v65 & (v65 | -128))))
        v5 += (v5 - (v5 * float32(3.14159)))
        _ = v65
    }
    v3["\n\t"] = true
    for fi66 := 0; fi66 < 4; fi66++ {
        var v67 string = "\n\t"
        if true {
            if false {
                for fi68 := 0; fi68 < 8; fi68++ {
                    v5 += 1e10
                    if (v5 <= ((v5 * (v5 * v5)) / v5)) {
                        var v69 int8 = -29
                        _ = v69
                    } else {
                        switch v67 {
                        case "hello":
                            switch fi66 {
                            case 34:
                                v5 += 0.1
                            default:
                                var v70 bool = ((!(genFn1(-9223372036854775808, 300, 1) != 1)) || true)
                                _ = v70
                            }
                        case "a":
                            v5 += (float32((-2.25 * ((v4 + 0.0) / v4))) + 0.1)
                            if (((0.0 * -2.25) + (float64(-2.25) * v4)) >= v4) {
                                fmt.Println(fi68)
                            } else {
                                fmt.Println(v4, v67, v4)
                            }
                            v3[("a" + ("a" + (v67 + v67)))] = (960 == 2)
                            _ = len(v3)
                        case "foo bar":
                            v5 += 1e10
                        }
                        fmt.Println(v6, fi66)
                    }
                    if (-2147483648 <= 42) {
                        fmt.Println(v4, v5)
                        var v71 int8 = 1
                        fmt.Println(v6)
                        _ = v71
                    } else {
                        fmt.Println(v5, v67, v4)
                        v72 := GenS4{F0: -1, F1: 100, F2: 140}
                        _ = v72
                    }
                    _ = fi68
                }
                if (-2.25 <= v5) {
                    var v73 int32 = 34
                    _ = v73
                }
            } else {
                for fi74 := 0; fi74 < 13; fi74++ {
                    v75 := GenS1{F0: 2147483647}
                    v6 = (fi74 % (((fi74 / (((fi66 - fi74)) | 1))) | 1))
                    _ = fi74
                    _ = v75
                }
                v67 = v67
            }
        }
        var v76 uint = 65535
        _ = fi66
        _ = v67
        _ = v76
    }
    if (93 >= 2) {
        v6 = 4294967295
    } else {
        v6 = 708
        v5 -= float32(v4)
    }
    v2 = append(v2, (-(v6)))
    switch ("a" + "hello") {
    case "foo bar":
        if (float64(3.14159) <= ((0.0 + 0.1) * float64((0.1 + -2.25)))) {
            for fi77 := 0; fi77 < 7; fi77++ {
                var v78 bool = (2147483647 > -2147483648)
                v6 -= fi77
                if v78 {
                    var v79 uint64 = 537
                    var v80 int64 = -1
                    for fi81 := 0; fi81 < 18; fi81++ {
                        if v78 {
                            var v82 uint64 = v79
                            var v83 float32 = v5
                            _ = v82
                            _ = v83
                        } else {
                            var v84 bool = (!v78)
                            switch v6 {
                            case 933:
                                v1 = append(v1, "a")
                            case 32767:
                                var v85 int = (-(fi77))
                                _ = v85
                            default:
                                var v86 string = ("foo bar" + "")
                                _ = v86
                            }
                            _ = v84
                        }
                        var v87 uint = 891
                        _ = fi81
                        _ = v87
                    }
                    _ = v79
                    _ = v80
                }
                _ = fi77
                _ = v78
            }
            v4 = v4
        } else {
            v88 := v3[("a" + "a")]
            if true {
                v6 += (v6 * 9223372036854775807)
            } else {
                if (42 >= 0) {
                    fmt.Println(v5, v88, v5)
                    if (("" != "foo bar") || (818 >= 42)) {
                        v89 := v3[""]
                        if (v88 || v88) {
                            var v90 uint8 = 42
                            if (277 < 533) {
                                var v91 uint = 2
                                for fi92 := 0; fi92 < 19; fi92++ {
                                    for fi93 := 0; fi93 < 19; fi93++ {
                                        var v94 int32 = -2147483648
                                        if (!v88) {
                                            if len(v2) > 0 {
                                                v2[0] = (fi92 | (fi93 & (fi93 | (-(v6)))))
                                            }
                                            v95 := GenS7{F0: 1}
                                            _ = v95
                                        }
                                        fmt.Println(fi92, v91)
                                        _ = fi93
                                        _ = v94
                                    }
                                    v4 = (float64(-2.25) / -2.25)
                                    if (float32(float64((1e10 / 1.5))) < 1.5) {
                                        switch (^(fi92)) {
                                        case -9223372036854775808:
                                            for fi96 := 0; fi96 < 18; fi96++ {
                                                var v97 int64 = 1
                                                fmt.Println(fi92)
                                                var v98 uint = 255
                                                v5 = v5
                                                v99 := GenS1{F0: 1}
                                                _ = fi96
                                                _ = v97
                                                _ = v98
                                                _ = v99
                                            }
                                        case 32768:
                                            if v89 {
                                                for fi100 := 0; fi100 < 19; fi100++ {
                                                    fmt.Println(v4)
                                                    var v101 int = v6
                                                    v102 := GenS2{F0: (v101 % ((-9223372036854775808) | 1))}
                                                    var v103 bool = v89
                                                    v3[("\n\t" + "foo bar")] = (-9223372036854775808 >= 2)
                                                    switch fi92 {
                                                    case 42:
                                                        v104 := GenS1{F0: -1}
                                                        _ = v104
                                                    case 128:
                                                        v3["foo bar"] = v88
                                                    default:
                                                        var v105 int = fi92
                                                        _ = v105
                                                    }
                                                    if (42 <= -32768) {
                                                        switch ("a" + ("a" + "a")) {
                                                        case "":
                                                            switch ("\n\t" + (("foo bar" + "foo bar") + "foo bar")) {
                                                            case "hello":
                                                                v3[(("a" + ("" + "hello")) + "foo bar")] = false
                                                            case "\n\t":
                                                                var v106 int32 = 399
                                                                _ = v106
                                                            default:
                                                                v107 := v3["foo bar"]
                                                                _ = v107
                                                            }
                                                        case "a":
                                                            var v108 int32 = 1
                                                            _ = v108
                                                        case "hello":
                                                            v6 = (v6 + fi100)
                                                            if true {
                                                                v109 := v3["a"]
                                                                fmt.Println(v91)
                                                                _ = v109
                                                            }
                                                        default:
                                                            fmt.Println(v6, v89, v5)
                                                        }
                                                        if false {
                                                            if (9223372036854775807 < -1) {
                                                                v5 = 3.14159
                                                            } else {
                                                                v90 -= (v90 >> (((v91 + v91)) & 63))
                                                            }
                                                            v110 := GenS4{F0: -2147483648, F1: 1, F2: (v90 | (v90 & v90))}
                                                            v111 := GenS8{F0: -1, F1: 32767}
                                                            _ = v110
                                                            _ = v111
                                                        } else {
                                                            v4 = 0.1
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = "foo bar"
                                                            }
                                                        }
                                                        v112 := GenS9{F0: "foo bar", F1: 18446744073709551615, F2: (v6 * (fi100 >> ((v91) & 63)))}
                                                        _ = v112
                                                    } else {
                                                        var v113 uint64 = 18446744073709551615
                                                        v103 = (!false)
                                                        var v114 uint8 = v90
                                                        if (v103 || v88) {
                                                            switch int(v90) {
                                                            case 4294967295:
                                                                v115 := v3[""]
                                                                _ = v115
                                                            }
                                                        } else {
                                                            var v116 uint = (^(v91))
                                                            v89 = (v89 && (!v103))
                                                            fmt.Println(fi100)
                                                            _ = v116
                                                        }
                                                        var v117 float64 = v4
                                                        _ = len(v3)
                                                        _ = v113
                                                        _ = v114
                                                        _ = v117
                                                    }
                                                    var v118 uint32 = 771
                                                    switch v101 {
                                                    case 592:
                                                        _ = len(v3)
                                                    default:
                                                        v5 -= 3.14159
                                                    }
                                                    v119 := GenS10{F0: v90}
                                                    fmt.Println(fi100)
                                                    var v120 int8 = -1
                                                    _ = fi100
                                                    _ = v101
                                                    _ = v102
                                                    _ = v103
                                                    _ = v118
                                                    _ = v119
                                                    _ = v120
                                                }
                                                v4 -= v4
                                                v88 = v89
                                            } else {
                                                v90 += (v90 >> ((v91) & 63))
                                                var v121 int16 = -1
                                                if (!(0 != -1)) {
                                                    v6 = (v6 & 4294967295)
                                                    switch ("hello" + "hello") {
                                                    case "\n\t":
                                                        for fi122 := 0; fi122 < 19; fi122++ {
                                                            var v123 uint64 = 1
                                                            if ((v121 != v121) && false) {
                                                                fmt.Println(fi92)
                                                                v124 := GenS11{F0: (v91 - v91), F1: ("a" + (("hello" + "\n\t") + ("foo bar" + ("hello" + "a")))), F2: 57}
                                                                _ = v124
                                                            } else {
                                                                if (v89 && true) {
                                                                    if true {
                                                                        var v125 string = ""
                                                                        _ = v125
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            v91 = 702
                                                                        default:
                                                                            if v89 {
                                                                                switch (v6 ^ (fi92 << (((v91 & uint(v90))) & 63))) {
                                                                                case 2147483647:
                                                                                    var v126 float64 = float64(v5)
                                                                                    _ = v126
                                                                                }
                                                                                switch v6 {
                                                                                case 127:
                                                                                    v127 := GenS10{F0: (^(v90))}
                                                                                    _ = v127
                                                                                case 2147483647:
                                                                                    var v128 int = v6
                                                                                    _ = v128
                                                                                default:
                                                                                    if v88 {
                                                                                        fmt.Println(v88, v88, v90)
                                                                                    }
                                                                                }
                                                                                v91 -= uint(v90)
                                                                            }
                                                                        }
                                                                        fmt.Println(v91)
                                                                        v3["\n\t"] = (v88 && ("" != "a"))
                                                                    }
                                                                    if ((v121 ^ (v121 / (((v121 - (v121 | (v121 + v121)))) | 1))) < v121) {
                                                                        if false {
                                                                            switch ("hello" + "foo bar") {
                                                                            case "a":
                                                                                var v129 int64 = 1
                                                                                _ = v129
                                                                            case "\n\t":
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = "foo bar"
                                                                                }
                                                                            default:
                                                                                if ((!v89) && false) {
                                                                                    var v130 bool = false
                                                                                    _ = v130
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v131 int = fi122
                                                                            v131 -= 0
                                                                            if v89 {
                                                                                v5 = ((v5 / 0.1) - 1e10)
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                if (("a" + "foo bar") < "a") {
                                                                                    v132 := v3["hello"]
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        switch -9223372036854775808 {
                                                                                        case -9223372036854775808:
                                                                                            var v133 int16 = (v121 + v121)
                                                                                            _ = v133
                                                                                        case 1:
                                                                                            fmt.Println(fi92)
                                                                                        default:
                                                                                            if ((2 != 42) || v88) {
                                                                                                var v134 uint16 = 32767
                                                                                                v91 = (v91 % ((v91) | 1))
                                                                                                v121 -= v121
                                                                                                _ = v134
                                                                                            } else {
                                                                                                v4 = v4
                                                                                            }
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v4 -= float64(1e10)
                                                                                    }
                                                                                    switch (-(v6)) {
                                                                                    case -9223372036854775808:
                                                                                        var v135 uint = v91
                                                                                        _ = v135
                                                                                    case 341:
                                                                                        v90 = v90
                                                                                    case -1:
                                                                                        v90 = (v90 << (((v91 / ((v91) | 1))) & 63))
                                                                                    }
                                                                                    _ = v132
                                                                                }
                                                                                var v136 int16 = (v121 & (v121 * (v121 & v121)))
                                                                                _ = v136
                                                                            }
                                                                            _ = v131
                                                                        }
                                                                        if (!(v4 <= v4)) {
                                                                            v137 := GenS12{F0: ((((!v89) || ("foo bar" >= "a")) && v89) || ("" <= "foo bar")), F1: ""}
                                                                            fmt.Println(v123, v88)
                                                                            _ = len(v1)
                                                                            var v138 uint32 = 455
                                                                            _ = v137
                                                                            _ = v138
                                                                        }
                                                                        v139 := GenS2{F0: int(v91)}
                                                                        _ = v139
                                                                    }
                                                                    v91 += 4294967295
                                                                }
                                                            }
                                                            v91 = 4294967295
                                                            v123 += (v123 * v123)
                                                            v6 -= 9223372036854775807
                                                            _ = fi122
                                                            _ = v123
                                                        }
                                                    case "a":
                                                        v121 = v121
                                                    default:
                                                        if (!((("foo bar" + "foo bar") + "") <= "foo bar")) {
                                                            switch 266 {
                                                            case 4294967295:
                                                                fmt.Println(v88, v91)
                                                            case -948:
                                                                var v140 string = "a"
                                                                _ = v140
                                                            default:
                                                                v90 += (v90 & 255)
                                                            }
                                                            v4 = v4
                                                        }
                                                    }
                                                }
                                                _ = v121
                                            }
                                        case -701:
                                            v88 = true
                                        default:
                                            switch (((("foo bar" + "") + "hello") + "\n\t") + "") {
                                            case "\n\t":
                                                for fi141 := 0; fi141 < 14; fi141++ {
                                                    if (((float32((float64(v5) + 0.1)) * (v5 - -2.25)) - (v5 - 3.14159)) >= (0.0 / (float32(float64(((v5 * ((v5 * (v5 - (v5 + (1e10 * float32((v4 * 1.5)))))) - 1e10)) - v5))) / v5))) {
                                                        v1 = append(v1, "foo bar")
                                                        if (!v89) {
                                                            if (-32768 <= -951) {
                                                                v142 := GenS6{F0: 65535, F1: (^(v90))}
                                                                v91 = 42
                                                                _ = v142
                                                            } else {
                                                                switch ("a" + "a") {
                                                                case "a":
                                                                    var v143 string = "\n\t"
                                                                    _ = v143
                                                                case "":
                                                                    var v144 uint16 = 32767
                                                                    _ = v144
                                                                default:
                                                                    v5 += (v5 - 1e10)
                                                                }
                                                                v89 = v89
                                                                fmt.Println(v90)
                                                                v145 := GenS1{F0: 442}
                                                                _ = v145
                                                            }
                                                        }
                                                    }
                                                    v4 = (float64(v5) + v4)
                                                    v4 -= 0.1
                                                    _ = fi141
                                                }
                                            case "":
                                                var v146 float64 = 1e10
                                                _ = v146
                                            case "a":
                                                v3["a"] = (1.5 != 3.14159)
                                            }
                                        }
                                    }
                                    _ = fi92
                                }
                                _ = v91
                            }
                            _ = v90
                        }
                        v147 := GenS13{F0: 42}
                        _ = v147
                        _ = v89
                    } else {
                        v148 := GenS14{F0: ("hello" + (("\n\t" + ("" + ("" + ""))) + (("hello" + "\n\t") + ((("\n\t" + "\n\t") + ("foo bar" + "\n\t")) + "foo bar"))))}
                        v4 += ((0.1 - (((0.1 * 1e10) * 0.0) * v4)) - float64(0.1))
                        switch "a" {
                        case "":
                            v88 = true
                        case "\n\t":
                            var v149 uint32 = 0
                            _ = v149
                        default:
                            fmt.Println(v88, v4, v6)
                        }
                        _ = v148
                    }
                }
                var v150 int8 = -128
                _ = v150
            }
            var v151 uint16 = 65535
            _ = v151
            _ = v88
        }
    }
    for fi152 := 0; fi152 < 6; fi152++ {
        if ((42 == 2147483647) && (v5 > v5)) {
            if (1e10 >= ((1e10 / v5) + 1e10)) {
                v153 := GenS15{F0: 2, F1: 9223372036854775807}
                var v154 bool = ((v6 * (fi152 + fi152)) != fi152)
                v1 = append(v1, "\n\t")
                _ = v153
                _ = v154
            } else {
                if false {
                    switch "hello" {
                    case "":
                        if false {
                            if false {
                                var v155 bool = (1 < 0)
                                for fi156 := 0; fi156 < 4; fi156++ {
                                    if ((0 > 0) && ((fi156 - -1) == fi152)) {
                                        var v157 float32 = v5
                                        v158 := GenS9{F0: (("\n\t" + (("foo bar" + ("foo bar" + "\n\t")) + "a")) + "foo bar"), F1: 18446744073709551615, F2: (^(fi156))}
                                        v159 := GenS1{F0: 2147483647}
                                        _ = v157
                                        _ = v158
                                        _ = v159
                                    } else {
                                        delete(v3, ("" + "a"))
                                        fmt.Println(fi156)
                                        switch fi152 {
                                        case 737:
                                            if (((3.14159 - v5) - v5) == v5) {
                                                fmt.Println(v5)
                                                fmt.Println(v155)
                                                v160 := GenS16{F0: 1, F1: 0}
                                                _ = v160
                                            } else {
                                                v5 = v5
                                            }
                                        case 32767:
                                            v5 = float32(((v4 / 1e10) - (v4 / 0.1)))
                                        default:
                                            fmt.Println(v4, v5)
                                        }
                                    }
                                    var v161 uint8 = 1
                                    _ = fi156
                                    _ = v161
                                }
                                _ = v155
                            }
                            v6 += (v6 | 4294967295)
                        } else {
                            v4 += 3.14159
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(fi152)
                        }
                    default:
                        v4 -= (-2.25 * float64((v5 / v5)))
                    }
                    if (!(429 != 765)) {
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                    } else {
                        for fi162 := 0; fi162 < 15; fi162++ {
                            if (fi152 < v6) {
                                v5 += ((v5 + 0.1) - (float32(v4) / (v5 + v5)))
                                if (!(867 == -92)) {
                                    if true {
                                        fmt.Println(v6, v4, v6)
                                        var v163 float32 = 1.5
                                        var v164 bool = (32767 <= 784)
                                        _ = v163
                                        _ = v164
                                    } else {
                                        fmt.Println(v6)
                                        _ = len(v2)
                                    }
                                    var v165 int = 42
                                    _ = v165
                                }
                                v166 := GenS13{F0: 1}
                                v167 := GenS9{F0: "\n\t", F1: 308, F2: fi162}
                                _ = v166
                                _ = v167
                            } else {
                                _ = len(v3)
                            }
                            v4 = (3.14159 / (1.5 / 0.1))
                            v5 -= 1.5
                            v3["hello"] = (42 > 1)
                            var v168 uint32 = 0
                            _ = fi162
                            _ = v168
                        }
                        v169 := GenS17{F0: 127, F1: 42, F2: ((fi152 & (fi152 + 127)) <= (v6 / (((fi152 % ((314) | 1))) | 1)))}
                        _ = v169
                    }
                    for fi170 := 0; fi170 < 19; fi170++ {
                        if (0 < 0) {
                            fmt.Println(fi170)
                            fmt.Println(v4)
                        }
                        if true {
                            switch v6 {
                            case -2147483648:
                                if false {
                                    v5 -= ((float32(v4) + v5) * v5)
                                    if false {
                                        fmt.Println(fi170, v4, v4)
                                    } else {
                                        var v171 float32 = v5
                                        var v172 uint16 = 0
                                        fmt.Println(v172, v4)
                                        _ = v171
                                        _ = v172
                                    }
                                } else {
                                    var v173 bool = (true || (42 == 160))
                                    switch (-(fi152)) {
                                    case -2147483648:
                                        v2 = append(v2, 1)
                                    case -228:
                                        if len(v2) > 0 {
                                            v2[0] = (fi152 & (-(v6)))
                                        }
                                    case -629:
                                        var v174 string = (("\n\t" + "hello") + ("\n\t" + "foo bar"))
                                        _ = v174
                                    }
                                    v175 := v3["a"]
                                    _ = v173
                                    _ = v175
                                }
                            case 977:
                                if (2147483648 == 526) {
                                    v176 := GenS15{F0: 18446744073709551615, F1: int64(fi152)}
                                    v5 = v5
                                    switch -1 {
                                    case 42:
                                        v177 := GenS18{F0: 18446744073709551615, F1: (("hello" + ("foo bar" + "foo bar")) + ("a" + "")), F2: 452}
                                        _ = v177
                                    case 9223372036854775807:
                                        delete(v3, "hello")
                                        var v178 float64 = ((((3.14159 + v4) * (float64(((v5 * v5) * v5)) - v4)) - v4) - (((-2.25 + v4) * (v4 / 0.1)) / v4))
                                        v179 := GenS19{F0: 4294967295, F1: 329, F2: 255}
                                        _ = v178
                                        _ = v179
                                    case -580:
                                        v180 := GenS14{F0: ("\n\t" + "a")}
                                        _ = v180
                                    default:
                                        v5 = 0.0
                                        if true {
                                            v5 = (((v5 + 3.14159) - v5) - ((0.0 / (0.1 + ((0.0 * float32(0.0)) * v5))) * v5))
                                            v6 = (v6 * 42)
                                            v2 = append(v2, (-(fi152)))
                                        }
                                        v181 := GenS15{F0: 4294967295, F1: -302}
                                        _ = v181
                                    }
                                    _ = v176
                                } else {
                                    _ = len(v3)
                                }
                            }
                            var v182 int32 = 1
                            var v183 uint8 = 0
                            _ = v182
                            _ = v183
                        } else {
                            var v184 int8 = genFn0()
                            _ = v184
                        }
                        _ = fi170
                    }
                } else {
                    var v185 int16 = 42
                    for fi186 := 0; fi186 < 9; fi186++ {
                        if (((v185 ^ 42) < (v185 - -1)) || (241 < 1)) {
                            v5 = 0.1
                        } else {
                            v187 := GenS1{F0: 1}
                            v188 := GenS1{F0: -1}
                            _ = v187
                            _ = v188
                        }
                        var v189 float64 = float64((v5 * ((v5 - v5) / v5)))
                        _ = fi186
                        _ = v189
                    }
                    _ = v185
                }
            }
            for fi190 := 0; fi190 < 4; fi190++ {
                v4 = (0.0 * (((v4 - v4) - 0.0) * (v4 - 0.0)))
                if ((255 <= 103) && true) {
                    v191 := GenS3{F0: 436, F1: 42, F2: 93}
                    _ = v191
                }
                _ = fi190
            }
            switch 0 {
            case 128:
                v192 := GenS20{F0: genFn0(), F1: 0}
                _ = v192
            default:
                for fi193 := 0; fi193 < 11; fi193++ {
                    if false {
                        var v194 int16 = 32767
                        var v195 float32 = (((1e10 + (1.5 / v5)) + v5) / (v5 - v5))
                        _ = v194
                        _ = v195
                    }
                    v196 := GenS2{F0: (v6 / ((32767) | 1))}
                    var v197 string = ("\n\t" + ((((("a" + "\n\t") + "") + ("" + "\n\t")) + "\n\t") + ""))
                    _ = fi193
                    _ = v196
                    _ = v197
                }
            }
        }
        v4 += (0.1 * (-2.25 / float64((v5 + float32(3.14159)))))
        if true {
            for fi198 := 0; fi198 < 3; fi198++ {
                if len(v1) > 0 {
                    v1[0] = (((("hello" + "\n\t") + ("" + ((("foo bar" + (("" + (("" + ("foo bar" + "")) + ("hello" + ""))) + "")) + "foo bar") + "a"))) + ("\n\t" + "")) + ("a" + "\n\t"))
                }
                _ = fi198
            }
        } else {
            v5 += float32(v4)
            switch 0 {
            case 42:
                v199 := GenS8{F0: 63, F1: 65535}
                _ = v199
            }
        }
        _ = fi152
    }
    v200 := GenS21{F0: 2147483647}
    var v201 int8 = 1
    for fi202 := 0; fi202 < 19; fi202++ {
        var v203 uint64 = 874
        _ = len(v2)
        _ = fi202
        _ = v203
    }
    v204 := GenS22{F0: v201, F1: v6, F2: 2}
    var v205 string = ""
    if false {
        v206 := v3[v205]
        _ = v206
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v201:", v201)
    fmt.Println("v205:", v205)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v200
    _ = v201
    _ = v204
    _ = v205
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
