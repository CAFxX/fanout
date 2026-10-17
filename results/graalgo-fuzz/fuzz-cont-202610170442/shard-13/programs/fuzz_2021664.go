package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint32; F2 string }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint8; F1 bool }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 bool; F1 bool }
type GenS11 struct { F0 bool }
func genFn0() uint {
    var v6 int32 = 2147483647
    if false {
        var v7 uint16 = 32767
        v7 -= 2
        _ = v7
    }
    var v8 int8 = -128
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint16 = 42
        _ = fi9
        _ = v10
    }
    _ = v6
    _ = v8
    return 2
}

func main() {
    var v1 uint32 = 106
    v2 := make(map[string]string)
    var v3 uint = 18446744073709551615
    var v4 float32 = 1.5
    v5 := make(map[string]bool)
    if true {
        var v11 int = -2147483648
        v11 = (v11 | (v11 % ((973) | 1)))
        fmt.Println(v3, v3, v1)
        _ = v11
    } else {
        if ((255 < 0) && true) {
            fmt.Println(v4, v3, v4)
        }
        v12 := GenS1{F0: -9223372036854775808, F1: 28}
        _ = v12
    }
    delete(v2, "a")
    switch 4294967295 {
    case 1:
        v13 := []int{880}
        _ = v13
    case 127:
        var v14 int8 = 66
        _ = v14
    case 128:
        if (32767 <= 255) {
            switch 1 {
            case 32767:
                var v15 float32 = v4
                _ = v15
            case 127:
                fmt.Println(v1, v1, v4)
            case 32768:
                if false {
                    if true {
                        var v16 int32 = -1
                        if (!((^(v1)) <= v1)) {
                            if (!(4294967295 > 1)) {
                                switch 2147483647 {
                                case 128:
                                    v17 := GenS1{F0: 0, F1: 127}
                                    _ = v17
                                case 32768:
                                    var v18 int = -9223372036854775808
                                    _ = v18
                                case 127:
                                    v16 = v16
                                }
                                switch "a" {
                                case "foo bar":
                                    switch ("hello" + "foo bar") {
                                    case "a":
                                        v19 := GenS2{F0: 18446744073709551615}
                                        _ = v19
                                    case "foo bar":
                                        v3 += genFn0()
                                    case "\n\t":
                                        for fi20 := 0; fi20 < 10; fi20++ {
                                            var v21 uint = 18446744073709551615
                                            _ = fi20
                                            _ = v21
                                        }
                                    }
                                default:
                                    switch 128 {
                                    case 4294967295:
                                        if (900 != -1) {
                                            var v22 uint = v3
                                            _ = v22
                                        } else {
                                            v4 = 1.5
                                            for fi23 := 0; fi23 < 3; fi23++ {
                                                var v24 float64 = 3.14159
                                                _ = fi23
                                                _ = v24
                                            }
                                            v4 = v4
                                        }
                                    case 128:
                                        if false {
                                            var v25 int8 = -128
                                            var v26 int32 = 0
                                            var v27 int32 = v16
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        }
                                    default:
                                        for fi28 := 0; fi28 < 2; fi28++ {
                                            for fi29 := 0; fi29 < 15; fi29++ {
                                                if (!(-290 != (v16 / ((0) | 1)))) {
                                                    fmt.Println(v1)
                                                } else {
                                                    v16 = (v16 * (v16 - -219))
                                                    delete(v5, ("hello" + "foo bar"))
                                                    if false {
                                                        var v30 int64 = 0
                                                        _ = v30
                                                    }
                                                    v3 = genFn0()
                                                    if true {
                                                        if (1.5 < (v4 + ((v4 * 0.1) * (1.5 + v4)))) {
                                                            v1 = v1
                                                        } else {
                                                            var v31 uint64 = 4294967295
                                                            _ = v31
                                                        }
                                                        var v32 bool = false
                                                        _ = v32
                                                    }
                                                }
                                                fmt.Println(v1, fi28, v1)
                                                _ = len(v5)
                                                _ = fi29
                                            }
                                            fmt.Println(fi28, v16, v3)
                                            _ = fi28
                                        }
                                    }
                                }
                            }
                        }
                        if true {
                            v3 += 18446744073709551615
                        } else {
                            var v33 int = 42
                            _ = v33
                        }
                        _ = v16
                    }
                    v1 = (v1 << (((v3 ^ 4294967295)) & 63))
                } else {
                    for fi34 := 0; fi34 < 18; fi34++ {
                        v1 = 741
                        _ = fi34
                    }
                }
            default:
                v3 = (v3 / (((v3 ^ (v3 >> (((v3 & (v3 - v3))) & 63)))) | 1))
            }
            switch "hello" {
            case "a":
                v35 := []int{32767, -2147483648, 42}
                _ = v35
            case "":
                fmt.Println(v4, v3)
            default:
                if (!true) {
                    v1 -= v1
                    var v36 uint16 = 406
                    fmt.Println(v3, v3, v3)
                    _ = v36
                }
            }
            var v37 uint16 = 0
            _ = v37
        }
    }
    v38 := GenS2{F0: 42}
    v39 := GenS3{F0: 298, F1: v1, F2: ""}
    var v40 uint32 = v1
    var v41 int16 = 1000
    v1 -= (^(v40))
    v40 = 2147483648
    if (9223372036854775807 <= 1) {
        switch "" {
        case "foo bar":
            for fi42 := 0; fi42 < 18; fi42++ {
                var v43 float32 = ((v4 - 0.1) - v4)
                _ = fi42
                _ = v43
            }
        case "\n\t":
            v3 += (v3 + v3)
        case "hello":
            if (!(!true)) {
                var v44 int32 = -1
                var v45 bool = (1 != 18446744073709551615)
                _ = v44
                _ = v45
            } else {
                var v46 int16 = (v41 << ((v3) & 63))
                _ = v46
            }
        default:
            v47 := GenS4{F0: 2, F1: "foo bar"}
            _ = v47
        }
    }
    switch ((("\n\t" + "\n\t") + "") + ("" + "foo bar")) {
    case "":
        v3 = v3
    case "a":
        v1 -= 1
    case "foo bar":
        for fi48 := 0; fi48 < 16; fi48++ {
            var v49 bool = false
            switch "foo bar" {
            case "\n\t":
                var v50 int64 = -676
                _ = v50
            case "a":
                v51 := GenS3{F0: 4294967295, F1: 376, F2: "hello"}
                _ = v51
            case "":
                if (!v49) {
                    if (v49 && ((v41 - v41) != -32768)) {
                        v52 := GenS2{F0: 18446744073709551615}
                        if (v49 && v49) {
                            switch "foo bar" {
                            case "hello":
                                v53 := GenS3{F0: 0, F1: (^(v1)), F2: "hello"}
                                _ = v53
                            }
                        }
                        _ = v52
                    }
                } else {
                    switch (^(fi48)) {
                    case -1:
                        var v54 float64 = 0.1
                        _ = v54
                    case 32767:
                        fmt.Println(v41)
                    case 399:
                        if v49 {
                            if v49 {
                                if v49 {
                                    switch "\n\t" {
                                    case "":
                                        v55 := []int{-1, fi48, 1}
                                        _ = v55
                                    case "foo bar":
                                        v1 = (^(v1))
                                    case "a":
                                        v56 := GenS3{F0: 2, F1: 1, F2: ""}
                                        _ = v56
                                    }
                                    v4 = 0.0
                                    v40 -= uint32(v3)
                                }
                                var v57 bool = v49
                                switch -1 {
                                case 32768:
                                    var v58 int = fi48
                                    _ = v58
                                }
                                _ = v57
                            } else {
                                _ = len(v5)
                                for fi59 := 0; fi59 < 16; fi59++ {
                                    switch "\n\t" {
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v60 int = (^(fi48))
                                            _ = v60
                                        }
                                    case "":
                                        if (uint8(v3) <= 1) {
                                            var v61 uint8 = 128
                                            v62 := GenS1{F0: 1, F1: 10}
                                            switch (fi48 / ((fi48) | 1)) {
                                            case 0:
                                                switch fi48 {
                                                case 32767:
                                                    var v63 int16 = 42
                                                    _ = v63
                                                default:
                                                    var v64 int8 = -128
                                                    _ = v64
                                                }
                                            case -167:
                                                v61 = 128
                                            case -2147483648:
                                                v65 := GenS2{F0: 168}
                                                _ = v65
                                            default:
                                                if (!((v41 * (v41 * (-(v41)))) != v41)) {
                                                    v66 := GenS3{F0: 931, F1: v40, F2: ("hello" + "\n\t")}
                                                    _ = v66
                                                } else {
                                                    var v67 string = "hello"
                                                    var v68 int64 = -9223372036854775808
                                                    _ = v67
                                                    _ = v68
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    var v69 uint32 = 762
                                                    _ = v69
                                                default:
                                                    switch ("a" + ("foo bar" + "foo bar")) {
                                                    case "\n\t":
                                                        switch "foo bar" {
                                                        case "":
                                                            var v70 uint16 = 255
                                                            _ = v70
                                                        }
                                                    case "":
                                                        switch (fi48 - (fi48 * fi59)) {
                                                        case -151:
                                                            var v71 uint8 = v61
                                                            _ = v71
                                                        default:
                                                            var v72 bool = false
                                                            switch fi48 {
                                                            case 2147483647:
                                                                if v49 {
                                                                    var v73 uint8 = (v61 / (((v61 | uint8(v3))) | 1))
                                                                    switch (fi48 + 127) {
                                                                    case -1:
                                                                        v4 -= v4
                                                                    }
                                                                    _ = v73
                                                                }
                                                            case 651:
                                                                if v72 {
                                                                    v74 := GenS5{F0: uint16(v3), F1: 42, F2: 18446744073709551615}
                                                                    _ = v74
                                                                }
                                                            case -2147483648:
                                                                fmt.Println(v40, fi48, v41)
                                                            }
                                                            var v75 string = "foo bar"
                                                            _ = v72
                                                            _ = v75
                                                        }
                                                    default:
                                                        if (955 <= 1) {
                                                            if (2 == 326) {
                                                                v76 := []int{fi48, fi59}
                                                                var v77 uint8 = (v61 * v61)
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                            fmt.Println(v3)
                                                            if false {
                                                                v78 := GenS6{F0: 42}
                                                                _ = v78
                                                            }
                                                            switch ((("a" + "") + "a") + "foo bar") {
                                                            case "foo bar":
                                                                fmt.Println(fi48, v40)
                                                            }
                                                            if ((v4 - ((0.0 / v4) / v4)) == float32(1e10)) {
                                                                var v79 uint64 = 322
                                                                v5["\n\t"] = v49
                                                                _ = v79
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            v80 := v2["hello"]
                                            var v81 uint8 = v61
                                            _ = v61
                                            _ = v62
                                            _ = v80
                                            _ = v81
                                        }
                                    case "a":
                                        fmt.Println(v3, v49, fi48)
                                    }
                                    v49 = v49
                                    if false {
                                        var v82 float32 = (v4 * (v4 - ((v4 - 0.0) * (v4 / (((3.14159 * v4) / 1e10) * v4)))))
                                        _ = v82
                                    }
                                    _ = fi59
                                }
                                var v83 uint = v3
                                _ = v83
                            }
                            v40 = 2147483648
                            var v84 int8 = 115
                            _ = v84
                        } else {
                            for fi85 := 0; fi85 < 20; fi85++ {
                                delete(v2, "")
                                switch (^(fi48)) {
                                case 2147483647:
                                    v41 = 42
                                    v86 := []int{fi85, (^(fi85)), (fi85 ^ fi48), fi48}
                                    switch 650 {
                                    case -1:
                                        switch ("hello" + "") {
                                        case "\n\t":
                                            if false {
                                                v87 := GenS2{F0: 0}
                                                if (170 <= -1) {
                                                    v4 -= (-2.25 + ((float32(((0.1 / -2.25) * ((3.14159 + 3.14159) * 0.1))) - v4) - v4))
                                                    v49 = ((true && ((((284 > 42) && v49) || (v3 < (v3 | v3))) && (9223372036854775807 != 9223372036854775807))) && false)
                                                    var v88 float32 = v4
                                                    if len(v86) > 0 {
                                                        v86[0] = (fi85 & (fi48 | (fi48 * (fi85 * 2147483647))))
                                                    }
                                                    var v89 int32 = 1
                                                    _ = v88
                                                    _ = v89
                                                } else {
                                                    var v90 uint8 = 255
                                                    _ = v90
                                                }
                                                _ = v87
                                            } else {
                                                v5[("foo bar" + "\n\t")] = false
                                                if true {
                                                    v49 = (42 == 2)
                                                    if false {
                                                        if len(v86) > 0 {
                                                            v86[0] = -1
                                                        }
                                                    }
                                                    v4 -= ((v4 / 0.1) / 1.5)
                                                }
                                            }
                                        case "foo bar":
                                            v91 := GenS1{F0: 42, F1: 127}
                                            _ = v91
                                        default:
                                            if ((fi48 * (fi85 % ((fi48) | 1))) == fi48) {
                                                v92 := GenS4{F0: 128, F1: "a"}
                                                v93 := GenS2{F0: 4294967295}
                                                var v94 uint = 4294967295
                                                v94 = 0
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                            }
                                        }
                                    case 32767:
                                        v3 -= (^(v3))
                                    default:
                                        var v95 int32 = 516
                                        _ = v95
                                    }
                                    _ = v86
                                case 32767:
                                    fmt.Println(v3, fi85)
                                default:
                                    v96 := GenS7{F0: 42, F1: -2147483648}
                                    _ = v96
                                }
                                v40 = 65535
                                _ = fi85
                            }
                            v97 := GenS6{F0: 1}
                            if v49 {
                                v98 := v5["hello"]
                                var v99 int32 = int32(fi48)
                                for fi100 := 0; fi100 < 11; fi100++ {
                                    var v101 float32 = v4
                                    var v102 int = fi100
                                    switch ("a" + "hello") {
                                    case "\n\t":
                                        v101 = v4
                                    case "hello":
                                        switch fi100 {
                                        case -1:
                                            var v103 uint64 = 2
                                            _ = v103
                                        case 128:
                                            fmt.Println(v1, fi100, v98)
                                        case 77:
                                            var v104 uint = (v3 ^ (v3 >> (((^(v3))) & 63)))
                                            _ = v104
                                        }
                                    }
                                    _ = fi100
                                    _ = v101
                                    _ = v102
                                }
                                _ = v98
                                _ = v99
                            } else {
                                if v49 {
                                    var v105 string = ("\n\t" + ("foo bar" + "hello"))
                                    v41 = 42
                                    for fi106 := 0; fi106 < 12; fi106++ {
                                        if v49 {
                                            _ = len(v2)
                                        } else {
                                            if v49 {
                                                var v107 bool = true
                                                var v108 uint32 = v1
                                                _ = v107
                                                _ = v108
                                            } else {
                                                fmt.Println(v40)
                                            }
                                            delete(v5, "a")
                                        }
                                        v109 := GenS8{F0: 41, F1: (!v49)}
                                        if v49 {
                                            fmt.Println(fi48, v40)
                                            var v110 int64 = 9223372036854775807
                                            _ = v110
                                        }
                                        _ = fi106
                                        _ = v109
                                    }
                                    _ = v105
                                }
                                fmt.Println(v49, fi48, fi48)
                            }
                            _ = v97
                        }
                    default:
                        v111 := []int{(fi48 - (fi48 / ((127) | 1))), fi48, fi48}
                        _ = v111
                    }
                }
            default:
                v112 := GenS2{F0: 42}
                _ = v112
            }
            for fi113 := 0; fi113 < 10; fi113++ {
                v41 = -657
                var v114 uint16 = 65535
                switch ("" + ("" + "a")) {
                case "foo bar":
                    if true {
                        if v49 {
                            if v49 {
                                v115 := GenS9{F0: (^(v114)), F1: 0}
                                var v116 int8 = -86
                                switch (("a" + "hello") + "a") {
                                case "":
                                    v1 -= 536
                                case "foo bar":
                                    v117 := []int{}
                                    _ = v117
                                default:
                                    v118 := GenS3{F0: 18446744073709551615, F1: v40, F2: "\n\t"}
                                    _ = v118
                                }
                                _ = v115
                                _ = v116
                            }
                            var v119 int64 = 42
                            _ = v119
                        }
                        v41 += (v41 - 562)
                        v49 = v49
                    }
                case "hello":
                    var v120 float64 = 0.0
                    _ = v120
                default:
                    var v121 uint = (v3 * v3)
                    _ = v121
                }
                _ = fi113
                _ = v114
            }
            _ = fi48
            _ = v49
        }
    }
    var v122 bool = true
    if (9223372036854775807 <= -1) {
        if (false || (!v122)) {
            for fi123 := 0; fi123 < 10; fi123++ {
                var v124 int16 = (v41 & 42)
                _ = fi123
                _ = v124
            }
            var v125 int32 = 2147483647
            fmt.Println(v122)
            _ = v125
        } else {
            var v126 int16 = (v41 << ((v3) & 63))
            _ = v126
        }
        v1 -= (v1 / (((v40 + v1)) | 1))
        v3 += (^(v3))
    } else {
        if ((float32((1e10 + float64(v4))) * float32(0.1)) > 1e10) {
            if v122 {
                v40 += v1
            } else {
                v122 = ((3.14159 == (0.0 * 3.14159)) && false)
                var v127 int8 = 1
                _ = len(v2)
                _ = v127
            }
            v128 := GenS10{F0: v122, F1: false}
            for fi129 := 0; fi129 < 20; fi129++ {
                if v122 {
                    var v130 uint16 = 2
                    var v131 int32 = 0
                    _ = v130
                    _ = v131
                } else {
                    fmt.Println(v41)
                    fmt.Println(v122, v1, v122)
                    v1 -= (^(v40))
                }
                v132 := GenS9{F0: 2, F1: 0}
                _ = fi129
                _ = v132
            }
            _ = v128
        } else {
            for fi133 := 0; fi133 < 6; fi133++ {
                fmt.Println(v4, v1, v41)
                for fi134 := 0; fi134 < 7; fi134++ {
                    v2["foo bar"] = ("a" + ("\n\t" + ("foo bar" + (("a" + "foo bar") + "a"))))
                    fmt.Println(fi134)
                    v135 := []int{(-(fi133)), fi133}
                    _ = fi134
                    _ = v135
                }
                fmt.Println(v40, v1)
                _ = fi133
            }
            switch 42 {
            case -2147483648:
                for fi136 := 0; fi136 < 15; fi136++ {
                    v137 := GenS7{F0: 2, F1: -956}
                    v138 := []int{(fi136 >> ((v3) & 63)), (fi136 + (fi136 % (((fi136 ^ (fi136 & (-(fi136))))) | 1)))}
                    _ = fi136
                    _ = v137
                    _ = v138
                }
            case 32767:
                v3 -= v3
            case -1:
                var v139 int16 = v41
                _ = v139
            default:
                if (!v122) {
                    var v140 int8 = 127
                    v141 := []int{-9223372036854775808}
                    _ = v140
                    _ = v141
                }
            }
            v142 := GenS11{F0: v122}
            _ = v142
        }
        var v143 float32 = 0.1
        _ = v143
    }
    var v144 int16 = (v41 << ((v3) & 63))
    fmt.Println(v40, v40)
    fmt.Println("v1:", v1)
    fmt.Println("v122:", v122)
    fmt.Println("v144:", v144)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v122
    _ = v144
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
