package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS7 struct { F0 uint8; F1 int64; F2 int64 }
type GenS8 struct { F0 uint; F1 int; F2 string }
type GenS9 struct { F0 int; F1 uint64; F2 int64 }
type GenS10 struct { F0 uint; F1 int64; F2 string }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int32 }

func main() {
    v1 := []int{127, -157, -2147483648}
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    var v4 bool = false
    var v5 bool = false
    var v6 uint16 = 65535
    if (1.5 >= 0.1) {
        if false {
            if ((^(v6)) == (v6 - v6)) {
                if (301 < 1) {
                    var v7 float32 = 1.5
                    for fi8 := 0; fi8 < 2; fi8++ {
                        v4 = v5
                        for fi9 := 0; fi9 < 3; fi9++ {
                            v3["hello"] = fi8
                            _ = fi9
                        }
                        if v5 {
                            fmt.Println(v5, v4, fi8)
                            v7 = v7
                            v5 = v5
                        } else {
                            v10 := GenS1{F0: 135, F1: 42}
                            _ = v10
                        }
                        _ = fi8
                    }
                    switch (("foo bar" + ("foo bar" + ("hello" + ""))) + "hello") {
                    case "":
                        v11 := GenS2{F0: 597, F1: 42}
                        _ = v11
                    case "foo bar":
                        v1 = append(v1, -9223372036854775808)
                    case "a":
                        if true {
                            if false {
                                if v5 {
                                    v6 += (v6 - v6)
                                }
                            } else {
                                switch (("foo bar" + "foo bar") + "") {
                                case "foo bar":
                                    v7 = 1.5
                                case "a":
                                    var v12 uint16 = v6
                                    _ = v12
                                case "\n\t":
                                    v5 = true
                                }
                                v7 -= v7
                                for fi13 := 0; fi13 < 1; fi13++ {
                                    var v14 uint = 663
                                    v15 := GenS2{F0: 2147483648, F1: 1}
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                            }
                            if false {
                                if false {
                                    for fi16 := 0; fi16 < 11; fi16++ {
                                        switch ("foo bar" + "") {
                                        case "a":
                                            v17 := GenS3{F0: 881, F1: "hello"}
                                            _ = v17
                                        }
                                        var v18 int8 = -39
                                        var v19 int16 = 142
                                        _ = fi16
                                        _ = v18
                                        _ = v19
                                    }
                                    v20 := GenS4{F0: 42, F1: "\n\t"}
                                    for fi21 := 0; fi21 < 5; fi21++ {
                                        _ = len(v2)
                                        _ = fi21
                                    }
                                    _ = v20
                                } else {
                                    v22 := GenS5{F0: 9223372036854775807, F1: 6}
                                    var v23 float32 = ((((1.5 + (0.0 * 0.0)) - (v7 * ((-2.25 - 1.5) / v7))) * (v7 + 1e10)) / 1e10)
                                    switch "\n\t" {
                                    case "foo bar":
                                        v4 = true
                                    case "":
                                        switch 4294967295 {
                                        case -8:
                                            fmt.Println(v6, v7, v4)
                                        default:
                                            var v24 uint8 = 128
                                            _ = v24
                                        }
                                    }
                                    _ = v22
                                    _ = v23
                                }
                            } else {
                                switch "foo bar" {
                                case "foo bar":
                                    v7 += v7
                                }
                            }
                        }
                    }
                    _ = v7
                }
                switch -1 {
                case -9223372036854775808:
                    v25 := GenS6{F0: 2, F1: 376, F2: v6}
                    _ = v25
                case -185:
                    v6 = v6
                case 2147483647:
                    fmt.Println(v6, v5, v6)
                default:
                    for fi26 := 0; fi26 < 15; fi26++ {
                        var v27 int32 = -1
                        _ = fi26
                        _ = v27
                    }
                }
            } else {
                v28 := GenS1{F0: 0, F1: -1}
                _ = v28
            }
            for fi29 := 0; fi29 < 5; fi29++ {
                v30 := GenS3{F0: 1, F1: "hello"}
                if v4 {
                    if ((9223372036854775807 <= -658) && (0 != -1)) {
                        if ((!(!(!v4))) && (1 > 32767)) {
                            v31 := GenS4{F0: 255, F1: "foo bar"}
                            v4 = true
                            _ = v31
                        }
                        v4 = false
                    } else {
                        fmt.Println(v6, fi29)
                        var v32 int8 = 42
                        for fi33 := 0; fi33 < 5; fi33++ {
                            if v5 {
                                v3[""] = fi29
                                v32 = v32
                            }
                            fmt.Println(v5)
                            var v34 uint32 = 4294967295
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    }
                    fmt.Println(v6, fi29)
                    v5 = v5
                } else {
                    var v35 float64 = ((1.5 + float64(0.0)) - 0.0)
                    for fi36 := 0; fi36 < 8; fi36++ {
                        delete(v3, "hello")
                        v1 = append(v1, (fi36 | fi36))
                        var v37 uint32 = 42
                        if ("a" < "") {
                            v35 = -2.25
                            v38 := GenS4{F0: v6, F1: ""}
                            v39 := GenS7{F0: 245, F1: -1, F2: 0}
                            _ = v38
                            _ = v39
                        } else {
                            switch ("foo bar" + ("" + ("foo bar" + ("hello" + ("a" + ""))))) {
                            case "\n\t":
                                fmt.Println(v5, v4, fi36)
                                var v40 int64 = 503
                                _ = v40
                            case "foo bar":
                                var v41 int32 = 586
                                _ = v41
                            default:
                                if true {
                                    v6 -= 43
                                    v42 := GenS8{F0: 402, F1: 4294967295, F2: ("foo bar" + "")}
                                    var v43 uint16 = 762
                                    _ = v42
                                    _ = v43
                                } else {
                                    switch ("\n\t" + ("a" + "")) {
                                    case "":
                                        v44 := GenS1{F0: 1, F1: 1}
                                        _ = v44
                                    case "\n\t":
                                        if ((65535 != 2) || (!v4)) {
                                            _ = len(v3)
                                            fmt.Println(v5, fi36)
                                            v45 := GenS9{F0: (-(fi36)), F1: 608, F2: 9223372036854775807}
                                            _ = v45
                                        } else {
                                            var v46 bool = false
                                            var v47 uint16 = v6
                                            _ = v46
                                            _ = v47
                                        }
                                    case "foo bar":
                                        v48 := v2[(v37 >> ((970) & 63))]
                                        _ = v48
                                    }
                                    switch fi29 {
                                    case -2147483648:
                                        v49 := GenS10{F0: 0, F1: -9223372036854775808, F2: "hello"}
                                        if v4 {
                                            var v50 uint = 490
                                            var v51 float64 = ((-2.25 / v35) - (-2.25 + v35))
                                            _ = v50
                                            _ = v51
                                        }
                                        if (-879 > 528) {
                                            var v52 int64 = -791
                                            v52 = v52
                                            _ = v52
                                        } else {
                                            switch ("hello" + (("\n\t" + ("a" + "a")) + (("a" + "\n\t") + "hello"))) {
                                            case "foo bar":
                                                if (-539 < 8) {
                                                    var v53 uint = 2
                                                    var v54 int8 = 42
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    switch "a" {
                                                    case "foo bar":
                                                        if v4 {
                                                            var v55 uint = 1
                                                            if v5 {
                                                                v56 := GenS10{F0: 1, F1: -9223372036854775808, F2: "foo bar"}
                                                                if (!(!(!false))) {
                                                                    fmt.Println(v4, v5, v5)
                                                                } else {
                                                                    var v57 int8 = -128
                                                                    switch (fi36 * fi29) {
                                                                    case 346:
                                                                        var v58 float32 = float32((((v35 + 0.0) + -2.25) / (float64((1.5 + 1e10)) / (v35 / v35))))
                                                                        _ = v58
                                                                    case 2147483647:
                                                                        var v59 float64 = v35
                                                                        _ = v59
                                                                    default:
                                                                        v60 := GenS9{F0: 32767, F1: 802, F2: 328}
                                                                        v61 := GenS11{F0: (fi29 ^ int(v6))}
                                                                        if (977 < 18446744073709551615) {
                                                                            if (fi36 != (fi36 >> ((v55) & 63))) {
                                                                                v3["\n\t"] = (fi29 + (fi36 & (fi36 >> (((v55 >> ((42) & 63))) & 63))))
                                                                                _ = len(v1)
                                                                            }
                                                                        }
                                                                        switch fi36 {
                                                                        case -310:
                                                                            var v62 int8 = 0
                                                                            _ = v62
                                                                        default:
                                                                            var v63 int8 = (-(v57))
                                                                            _ = v63
                                                                        }
                                                                        var v64 uint32 = (^(v37))
                                                                        _ = v60
                                                                        _ = v61
                                                                        _ = v64
                                                                    }
                                                                    _ = v57
                                                                }
                                                                _ = v56
                                                            }
                                                            var v65 int32 = 0
                                                            _ = v55
                                                            _ = v65
                                                        } else {
                                                            v1 = append(v1, fi36)
                                                            if (!true) {
                                                                v66 := GenS7{F0: 0, F1: 42, F2: -1}
                                                                _ = v66
                                                            }
                                                            v2[v37] = (fi36 * -670)
                                                        }
                                                    default:
                                                        var v67 float64 = 3.14159
                                                        _ = v67
                                                    }
                                                }
                                            case "a":
                                                v5 = true
                                            default:
                                                if v4 {
                                                    if (v4 && v4) {
                                                        v4 = v5
                                                    } else {
                                                        switch ("hello" + "foo bar") {
                                                        case "\n\t":
                                                            v35 = (-2.25 / 3.14159)
                                                        case "a":
                                                            var v68 uint32 = v37
                                                            var v69 uint8 = 245
                                                            if v5 {
                                                                if v4 {
                                                                    v35 += -2.25
                                                                    var v70 float32 = 0.0
                                                                    _ = v70
                                                                } else {
                                                                    fmt.Println(fi29)
                                                                    v5 = (v4 && v4)
                                                                }
                                                            } else {
                                                                v37 += (^(v37))
                                                                var v71 float32 = (1.5 + -2.25)
                                                                _ = v71
                                                            }
                                                            _ = v68
                                                            _ = v69
                                                        case "hello":
                                                            _ = len(v2)
                                                        }
                                                    }
                                                } else {
                                                    _ = len(v2)
                                                }
                                            }
                                            switch (fi36 | fi36) {
                                            case 1:
                                                fmt.Println(v37, v37, v5)
                                            case 9223372036854775807:
                                                var v72 float64 = v35
                                                _ = v72
                                            default:
                                                switch ("foo bar" + "a") {
                                                case "":
                                                    if (100 < -64) {
                                                        v73 := GenS12{F0: -694}
                                                        var v74 float32 = 1.5
                                                        v5 = v5
                                                        _ = v73
                                                        _ = v74
                                                    }
                                                case "\n\t":
                                                    fmt.Println(v6)
                                                }
                                            }
                                            var v75 uint64 = 0
                                            _ = v75
                                        }
                                        _ = v49
                                    }
                                    v76 := GenS11{F0: (-(fi36))}
                                    _ = v76
                                }
                            }
                        }
                        delete(v3, "foo bar")
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                }
                for fi77 := 0; fi77 < 5; fi77++ {
                    var v78 uint16 = v6
                    if true {
                        var v79 bool = (!v4)
                        switch -1 {
                        case -2147483648:
                            switch "a" {
                            case "\n\t":
                                switch ("hello" + ((("" + ("foo bar" + (("foo bar" + "") + ""))) + "") + "\n\t")) {
                                case "hello":
                                    var v80 uint64 = 18446744073709551615
                                    _ = v80
                                case "":
                                    delete(v2, 2)
                                case "foo bar":
                                    v79 = true
                                default:
                                    _ = len(v1)
                                }
                                if ("\n\t" >= ("" + "foo bar")) {
                                    switch 32768 {
                                    case 118:
                                        var v81 uint8 = 255
                                        _ = v81
                                    case -2147483648:
                                        switch (fi29 * (fi77 * fi77)) {
                                        case 32768:
                                            switch fi29 {
                                            case 242:
                                                fmt.Println(fi77)
                                            default:
                                                _ = len(v1)
                                            }
                                        }
                                    case 9223372036854775807:
                                        if (18446744073709551615 < 42) {
                                            v82 := GenS4{F0: 835, F1: "hello"}
                                            _ = v82
                                        }
                                    }
                                    var v83 bool = v5
                                    fmt.Println(v6, v79, v79)
                                    _ = v83
                                }
                                if v79 {
                                    var v84 uint = 2
                                    v85 := GenS6{F0: 1, F1: 2, F2: (^(v6))}
                                    _ = v84
                                    _ = v85
                                }
                            case "hello":
                                v1 = append(v1, (fi77 % (((fi29 + fi77)) | 1)))
                            case "a":
                                var v86 int8 = 1
                                _ = v86
                            }
                        case 203:
                            v4 = (0 == 65535)
                        }
                        v87 := GenS2{F0: 620, F1: 0}
                        switch (fi29 / (((fi77 / ((fi29) | 1))) | 1)) {
                        case 1:
                            v6 = v6
                        case 212:
                            fmt.Println(fi29, v6, v78)
                        }
                        _ = v79
                        _ = v87
                    } else {
                        var v88 uint = 63
                        _ = v88
                    }
                    v89 := GenS9{F0: (-(fi77)), F1: 221, F2: -9223372036854775808}
                    v4 = (0 <= -420)
                    _ = fi77
                    _ = v78
                    _ = v89
                }
                _ = fi29
                _ = v30
            }
            for fi90 := 0; fi90 < 18; fi90++ {
                if len(v1) > 0 {
                    v1[0] = fi90
                }
                var v91 uint64 = 0
                v4 = v5
                _ = fi90
                _ = v91
            }
        } else {
            for fi92 := 0; fi92 < 10; fi92++ {
                for fi93 := 0; fi93 < 18; fi93++ {
                    var v94 int = (-(fi93))
                    if (!(!(!true))) {
                        var v95 int32 = 42
                        v94 = (fi93 * int(v95))
                        fmt.Println(fi92, v94)
                        v4 = v4
                        _ = v95
                    }
                    _ = fi93
                    _ = v94
                }
                _ = fi92
            }
        }
        var v96 int8 = 1
        _ = v96
    } else {
        if v5 {
            fmt.Println(v6, v4, v4)
        }
        v4 = v5
        v97 := GenS12{F0: 1}
        _ = v97
    }
    if true {
        v4 = (127 > -128)
    }
    for fi98 := 0; fi98 < 2; fi98++ {
        v6 -= (v6 | 42)
        switch ("a" + "") {
        case "hello":
            delete(v2, 2)
        case "a":
            var v99 uint8 = 0
            _ = v99
        case "foo bar":
            for fi100 := 0; fi100 < 5; fi100++ {
                if false {
                    fmt.Println(v5, v6)
                    if true {
                        var v101 string = "hello"
                        v3["hello"] = fi98
                        var v102 float64 = 0.0
                        _ = v101
                        _ = v102
                    }
                }
                v103 := GenS6{F0: 4294967295, F1: 4294967295, F2: (v6 ^ v6)}
                fmt.Println(fi98, v4)
                _ = fi100
                _ = v103
            }
        default:
            var v104 int32 = 0
            _ = v104
        }
        var v105 int = fi98
        _ = fi98
        _ = v105
    }
    fmt.Println(v4, v5)
    v6 = v6
    var v106 bool = false
    v107 := GenS1{F0: 0, F1: -69}
    for fi108 := 0; fi108 < 8; fi108++ {
        v5 = v4
        v5 = v4
        _ = fi108
    }
    for fi109 := 0; fi109 < 14; fi109++ {
        var v110 int8 = -128
        v4 = v5
        fmt.Println(v106)
        _ = fi109
        _ = v110
    }
    v111 := GenS12{F0: 1}
    _ = len(v2)
    v106 = (float64(1.5) > -2.25)
    for fi112 := 0; fi112 < 9; fi112++ {
        _ = len(v3)
        for fi113 := 0; fi113 < 14; fi113++ {
            _ = len(v1)
            _ = fi113
        }
        v106 = v106
        _ = fi112
    }
    var v114 float64 = -2.25
    if v5 {
        if (false && v5) {
            var v115 float32 = 0.0
            var v116 int = 0
            _ = v115
            _ = v116
        }
        v117 := GenS11{F0: 920}
        v4 = v5
        _ = v117
    } else {
        var v118 string = "hello"
        switch ("" + v118) {
        case "foo bar":
            if v5 {
                v114 -= v114
                var v119 float32 = 3.14159
                _ = v119
            }
        default:
            for fi120 := 0; fi120 < 18; fi120++ {
                var v121 int = (fi120 % (((fi120 >> ((42) & 63))) | 1))
                _ = fi120
                _ = v121
            }
        }
        _ = v118
    }
    v4 = v4
    fmt.Println(v106, v4, v106)
    var v122 uint32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v106:", v106)
    fmt.Println("v114:", v114)
    fmt.Println("v122:", v122)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v106
    _ = v107
    _ = v111
    _ = v114
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
