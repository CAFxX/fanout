package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 bool; F1 uint64; F2 bool }
type GenS7 struct { F0 int32; F1 int; F2 int32 }
type GenS8 struct { F0 bool; F1 int64; F2 uint }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint16; F1 uint64 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int64; F1 uint16 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int32; F1 int32 }
type GenS15 struct { F0 int64; F1 int8; F2 string }
type GenS16 struct { F0 uint64; F1 int16 }
type GenS17 struct { F0 uint64; F1 int8; F2 bool }
func genFn0(p0 uint16) int32 {
    fmt.Println(p0)
    v4 := []int{}
    if false {
        fmt.Println(p0)
        var v5 int64 = 42
        _ = v5
    }
    _ = p0
    _ = v4
    return 40
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{65535}
    v3 := []uint32{65535, 733, 543, 4294967295}
    if (4294967295 < 2147483647) {
        if true {
            var v6 uint = 4294967295
            var v7 int8 = -22
            v7 = v7
            _ = v6
            _ = v7
        } else {
            var v8 int8 = -128
            _ = v8
        }
        var v9 bool = (-2.25 == 3.14159)
        _ = v9
    }
    if (true && false) {
        v10 := GenS1{F0: 100, F1: 42}
        _ = v10
    }
    var v11 int8 = 1
    switch "\n\t" {
    case "":
        fmt.Println(v11)
    default:
        if false {
            var v12 uint = 65535
            v13 := GenS2{F0: false}
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 14; fi14++ {
                for fi15 := 0; fi15 < 4; fi15++ {
                    v11 += v11
                    if (!false) {
                        var v16 int64 = -9223372036854775808
                        v17 := GenS2{F0: (828 < 42)}
                        if (!true) {
                            v18 := GenS3{F0: 1, F1: v16}
                            v1[""] = 733
                            v19 := GenS2{F0: (2147483648 > 65535)}
                            _ = v18
                            _ = v19
                        }
                        v16 = v16
                        if len(v3) > 0 {
                            v3[0] = 4294967295
                        }
                        _ = v16
                        _ = v17
                    } else {
                        v20 := GenS1{F0: (v11 / (((^(v11))) | 1)), F1: -121}
                        _ = v20
                    }
                    _ = fi15
                }
                v11 += (v11 & -128)
                _ = fi14
            }
            delete(v1, "hello")
            switch 127 {
            case 338:
                var v21 float64 = 0.0
                _ = v21
            case -1:
                v11 -= v11
            }
        }
    }
    if (602 == 1) {
        if (0.0 != (3.14159 - (float32(0.0) + 0.1))) {
            delete(v1, ("a" + ""))
        } else {
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 int = -285
                _ = fi22
                _ = v23
            }
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := v1[""]
            if false {
                var v26 float64 = -2.25
                _ = v26
            }
            _ = fi24
            _ = v25
        }
    } else {
        var v27 uint16 = 578
        fmt.Println(v27, v27)
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        fmt.Println(v11, v11)
        var v29 uint32 = 2147483648
        _ = fi28
        _ = v29
    }
    v30 := GenS3{F0: 0, F1: 9223372036854775807}
    if ((1 <= 2147483647) && ("hello" == ("a" + "hello"))) {
        var v31 int16 = -32768
        for fi32 := 0; fi32 < 16; fi32++ {
            v33 := v1[""]
            switch -864 {
            case -9223372036854775808:
                fmt.Println(v31, fi32)
            case 32768:
                var v34 int64 = -810
                _ = v34
            case 653:
                v35 := GenS3{F0: 42, F1: -46}
                _ = v35
            default:
                v33 = fi32
            }
            v33 += (fi32 ^ (^(v33)))
            _ = fi32
            _ = v33
        }
        v36 := GenS2{F0: (!(!(!false)))}
        _ = v31
        _ = v36
    } else {
        if true {
            v11 += (v11 + (v11 + v11))
            switch "a" {
            case "\n\t":
                var v37 uint64 = 18446744073709551615
                _ = v37
            case "a":
                var v38 int8 = 0
                _ = v38
            case "hello":
                fmt.Println(v11)
            default:
                for fi39 := 0; fi39 < 12; fi39++ {
                    for fi40 := 0; fi40 < 2; fi40++ {
                        v41 := GenS4{F0: -197, F1: 130, F2: 38}
                        _ = fi40
                        _ = v41
                    }
                    _ = fi39
                }
            }
            var v42 int32 = -2147483648
            _ = v42
        } else {
            v43 := GenS5{F0: "hello", F1: 4294967295}
            v11 += (-(v11))
            _ = v43
        }
    }
    switch ("" + "\n\t") {
    case "\n\t":
        var v44 int16 = 42
        _ = v44
    case "hello":
        v11 -= (v11 + v11)
    default:
        var v45 uint = 577
        _ = v45
    }
    var v46 int8 = v11
    if (-18 != (v11 + (v11 ^ v11))) {
        for fi47 := 0; fi47 < 13; fi47++ {
            if (531 >= -594) {
                for fi48 := 0; fi48 < 13; fi48++ {
                    fmt.Println(v46, v11, fi47)
                    _ = fi48
                }
                var v49 int64 = 1
                _ = v49
            } else {
                if false {
                    var v50 uint8 = 255
                    var v51 float32 = 1e10
                    if ("foo bar" == ("hello" + ("foo bar" + "hello"))) {
                        var v52 uint = 65535
                        fmt.Println(fi47, v46)
                        for fi53 := 0; fi53 < 9; fi53++ {
                            v54 := GenS3{F0: v52, F1: 772}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v46)
                            _ = fi53
                            _ = v54
                        }
                        _ = v52
                    }
                    _ = v50
                    _ = v51
                } else {
                    var v55 float64 = (0.0 / 3.14159)
                    v46 = (v11 * 1)
                    _ = v55
                }
            }
            v46 += v11
            v56 := v1["\n\t"]
            _ = fi47
            _ = v56
        }
        delete(v1, "a")
        var v57 float32 = (float32(0.0) + (1.5 + 1.5))
        _ = v57
    }
    if true {
        for fi58 := 0; fi58 < 1; fi58++ {
            v59 := v1[(("hello" + ("a" + "\n\t")) + ("foo bar" + "\n\t"))]
            if false {
                if false {
                    v60 := GenS4{F0: 0, F1: 128, F2: 128}
                    v46 += v11
                    if (false && (-419 == -377)) {
                        v61 := GenS4{F0: 672, F1: 127, F2: 1}
                        v1[""] = fi58
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                        _ = v61
                    } else {
                        if true {
                            switch ("a" + "a") {
                            case "a":
                                switch ("foo bar" + ("a" + "foo bar")) {
                                case "foo bar":
                                    v62 := GenS4{F0: 114, F1: 128, F2: 1}
                                    _ = v62
                                case "hello":
                                    v1["hello"] = (v59 % ((-831) | 1))
                                case "":
                                    _ = len(v2)
                                default:
                                    v63 := GenS6{F0: true, F1: 1, F2: (((3.14159 / 0.1) <= float64(float32(-2.25))) && (false && (2147483647 == -1)))}
                                    _ = v63
                                }
                            case "":
                                for fi64 := 0; fi64 < 13; fi64++ {
                                    fmt.Println(v11)
                                    var v65 uint = 51
                                    v66 := GenS6{F0: false, F1: uint64(v65), F2: ((!true) || ((v59 << ((42) & 63)) > 945))}
                                    _ = fi64
                                    _ = v65
                                    _ = v66
                                }
                            case "hello":
                                if false {
                                    for fi67 := 0; fi67 < 2; fi67++ {
                                        v46 += v46
                                        switch (fi67 & (fi67 ^ (v59 << ((65535) & 63)))) {
                                        case -1:
                                            delete(v1, "foo bar")
                                        case 710:
                                            switch ("a" + "\n\t") {
                                            case "\n\t":
                                                if true {
                                                    v68 := GenS6{F0: (42 >= 924), F1: 694, F2: (811 > 4294967295)}
                                                    fmt.Println(fi58, fi58)
                                                    var v69 uint32 = 472
                                                    _ = v68
                                                    _ = v69
                                                }
                                            default:
                                                v70 := GenS4{F0: -978, F1: 255, F2: 242}
                                                _ = v70
                                            }
                                        case 0:
                                            v2 = append(v2, 4294967295)
                                        default:
                                            fmt.Println(v59)
                                        }
                                        switch "a" {
                                        case "":
                                            if true {
                                                var v71 int64 = 42
                                                _ = v71
                                            } else {
                                                var v72 int = v59
                                                if (true && true) {
                                                    var v73 int32 = 266
                                                    _ = v73
                                                } else {
                                                    var v74 int32 = -456
                                                    _ = v74
                                                }
                                                _ = v72
                                            }
                                        case "\n\t":
                                            if false {
                                                v46 += v46
                                            } else {
                                                var v75 bool = false
                                                _ = v75
                                            }
                                        case "hello":
                                            var v76 float32 = -2.25
                                            _ = v76
                                        }
                                        v59 -= fi67
                                        _ = fi67
                                    }
                                    for fi77 := 0; fi77 < 16; fi77++ {
                                        v1["\n\t"] = -600
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v59 -= (fi77 ^ -179)
                                        fmt.Println(fi77, v46)
                                        v78 := GenS7{F0: 0, F1: fi77, F2: -2147483648}
                                        _ = fi77
                                        _ = v78
                                    }
                                    v79 := GenS5{F0: "a", F1: 2}
                                    _ = v79
                                }
                            default:
                                _ = len(v2)
                            }
                            if (-2.25 >= 1e10) {
                                if true {
                                    if true {
                                        for fi80 := 0; fi80 < 16; fi80++ {
                                            v1["hello"] = fi80
                                            switch 2147483647 {
                                            case 32767:
                                                fmt.Println(fi80, v11)
                                            case 9223372036854775807:
                                                var v81 uint64 = 2
                                                _ = v81
                                            case 874:
                                                if (42 == 912) {
                                                    v11 += (v11 * v11)
                                                    v1["a"] = fi58
                                                    switch 0 {
                                                    case 4294967295:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            fmt.Println(v59, fi80, v59)
                                                        }
                                                    }
                                                } else {
                                                    switch "hello" {
                                                    case "a":
                                                        v11 += (v11 / (((v11 * v11)) | 1))
                                                    case "foo bar":
                                                        var v82 int8 = (v46 >> ((65535) & 63))
                                                        _ = v82
                                                    }
                                                }
                                            default:
                                                switch v59 {
                                                case 912:
                                                    v83 := GenS8{F0: true, F1: 905, F2: 107}
                                                    _ = v83
                                                case 42:
                                                    v46 = v46
                                                case 32767:
                                                    v46 += 127
                                                }
                                            }
                                            _ = fi80
                                        }
                                        for fi84 := 0; fi84 < 12; fi84++ {
                                            v85 := GenS3{F0: 255, F1: 42}
                                            v86 := v1["hello"]
                                            var v87 uint16 = 42
                                            _ = fi84
                                            _ = v85
                                            _ = v86
                                            _ = v87
                                        }
                                    }
                                    v11 += v11
                                    v46 = v46
                                } else {
                                    if true {
                                        var v88 uint32 = 2
                                        _ = v88
                                    } else {
                                        var v89 uint = 4294967295
                                        fmt.Println(v89)
                                        _ = v89
                                    }
                                    switch ("hello" + "a") {
                                    case "hello":
                                        var v90 float64 = -2.25
                                        _ = v90
                                    }
                                    if true {
                                        v1["\n\t"] = (v59 / ((128) | 1))
                                        v2 = append(v2, 872)
                                        switch "a" {
                                        case "a":
                                            v11 -= (-(v46))
                                        case "foo bar":
                                            v1[""] = (fi58 + v59)
                                        case "hello":
                                            if ((false || (!true)) && (683 > 1)) {
                                                v91 := GenS4{F0: genFn0(1), F1: 2, F2: 127}
                                                v46 -= 100
                                                var v92 bool = true
                                                _ = v91
                                                _ = v92
                                            }
                                        }
                                    }
                                }
                                _ = len(v2)
                            } else {
                                v11 = (v46 + (v11 | (^(v46))))
                                var v93 uint = 65535
                                for fi94 := 0; fi94 < 1; fi94++ {
                                    var v95 float64 = 1.5
                                    _ = fi94
                                    _ = v95
                                }
                                _ = v93
                            }
                            switch 929 {
                            case 42:
                                var v96 int = (v59 / (((fi58 << ((0) & 63))) | 1))
                                _ = v96
                            case 9223372036854775807:
                                fmt.Println(v46)
                            case 206:
                                var v97 int64 = 539
                                _ = v97
                            default:
                                fmt.Println(fi58, v46)
                            }
                        } else {
                            if true {
                                var v98 uint64 = 1
                                v1[""] = v59
                                _ = v98
                            }
                        }
                        var v99 uint32 = 1
                        v100 := GenS9{F0: 4294967295}
                        _ = v100
                        _ = v99
                    }
                    _ = v60
                } else {
                    v1["foo bar"] = (fi58 << ((460) & 63))
                }
                v101 := v1["hello"]
                fmt.Println(v46)
                _ = v101
            } else {
                v59 = (v59 / ((v59) | 1))
            }
            v11 -= v11
            _ = fi58
            _ = v59
        }
    } else {
        switch "" {
        case "\n\t":
            fmt.Println(v11, v46)
        case "a":
            v11 = v11
        case "":
            v46 -= (v46 * 0)
        default:
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        var v102 int16 = -317
        v103 := GenS1{F0: (v11 / ((0) | 1)), F1: (v11 & v11)}
        _ = v102
        _ = v103
    }
    switch 42 {
    case 135:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 655:
        for fi104 := 0; fi104 < 1; fi104++ {
            for fi105 := 0; fi105 < 14; fi105++ {
                var v106 int = 2147483647
                fmt.Println(v106, v11)
                v106 += 2147483647
                _ = fi105
                _ = v106
            }
            _ = fi104
        }
    }
    for fi107 := 0; fi107 < 13; fi107++ {
        switch (fi107 >> ((18446744073709551615) & 63)) {
        case 32767:
            var v108 bool = true
            _ = v108
        case 32768:
            fmt.Println(v46, v11, fi107)
        case 2147483647:
            for fi109 := 0; fi109 < 3; fi109++ {
                if (((("a" + ("foo bar" + "\n\t")) < "foo bar") && false) && false) {
                    v11 = 77
                    switch "" {
                    case "\n\t":
                        v110 := v1[""]
                        _ = v110
                    }
                    if false {
                        v11 = v11
                    } else {
                        v111 := GenS8{F0: (!true), F1: 1, F2: 65535}
                        switch fi109 {
                        case 127:
                            switch ("" + "foo bar") {
                            case "foo bar":
                                if true {
                                    v112 := GenS8{F0: (!(true || true)), F1: 196, F2: 758}
                                    if (("" + "a") < ("foo bar" + ("a" + "a"))) {
                                        if false {
                                            fmt.Println(fi109, v11)
                                            _ = len(v1)
                                        }
                                        v11 += (v11 / ((v46) | 1))
                                        fmt.Println(v11, fi109, fi109)
                                    }
                                    _ = v112
                                } else {
                                    if false {
                                        if true {
                                            var v113 bool = false
                                            v114 := GenS10{F0: 2, F1: 2}
                                            switch "a" {
                                            case "a":
                                                v115 := GenS11{F0: 872}
                                                _ = v115
                                            default:
                                                v116 := GenS12{F0: 42, F1: 42}
                                                _ = v116
                                            }
                                            _ = v113
                                            _ = v114
                                        } else {
                                            switch (fi107 + 2147483647) {
                                            case 128:
                                                v46 = 0
                                            }
                                            v46 -= v11
                                        }
                                        v1["hello"] = fi109
                                        var v117 int16 = -1
                                        if true {
                                            switch fi107 {
                                            case 9223372036854775807:
                                                v118 := GenS3{F0: 2, F1: 674}
                                                if (fi109 > (fi109 * fi109)) {
                                                    v119 := GenS2{F0: (false && true)}
                                                    _ = v119
                                                } else {
                                                    var v120 int64 = 9223372036854775807
                                                    fmt.Println(fi109, v11, fi109)
                                                    switch "\n\t" {
                                                    case "hello":
                                                        if false {
                                                            switch ((("" + "a") + "foo bar") + ("\n\t" + "foo bar")) {
                                                            case "hello":
                                                                v1["foo bar"] = (fi109 | (fi107 << ((1) & 63)))
                                                            case "":
                                                                v121 := GenS1{F0: (v11 & int8(fi109)), F1: v11}
                                                                _ = v121
                                                            case "foo bar":
                                                                if (392 >= 2) {
                                                                    switch (("\n\t" + ("a" + "a")) + (("" + "") + (("hello" + ("hello" + ("" + "hello"))) + "foo bar"))) {
                                                                    case "hello":
                                                                        var v122 float64 = float64((1.5 + 1.5))
                                                                        _ = v122
                                                                    case "":
                                                                        v123 := GenS13{F0: 637}
                                                                        _ = v123
                                                                    case "\n\t":
                                                                        var v124 int16 = -1
                                                                        _ = v124
                                                                    }
                                                                    fmt.Println(fi109, v11)
                                                                    var v125 uint64 = 495
                                                                    _ = v125
                                                                } else {
                                                                    var v126 string = ""
                                                                    v2 = append(v2, 4294967295)
                                                                    fmt.Println(v46)
                                                                    _ = v126
                                                                }
                                                            default:
                                                                if false {
                                                                    v120 += (v120 >> ((255) & 63))
                                                                    fmt.Println(v120, fi107)
                                                                    if (42 != 923) {
                                                                        v127 := GenS9{F0: 2}
                                                                        _ = v127
                                                                    }
                                                                    if (!true) {
                                                                        v11 = 42
                                                                        fmt.Println(v120, v11, fi109)
                                                                        v46 = v46
                                                                    }
                                                                } else {
                                                                    fmt.Println(v120, v117)
                                                                    var v128 uint16 = 32767
                                                                    _ = v128
                                                                }
                                                            }
                                                        }
                                                    case "a":
                                                        fmt.Println(fi107)
                                                    case "foo bar":
                                                        v46 = v11
                                                    }
                                                    var v129 uint = 0
                                                    v46 = (-(v11))
                                                    _ = v120
                                                    _ = v129
                                                }
                                                if (true && (((3.14159 - 0.0) - 3.14159) > 0.0)) {
                                                    switch ("foo bar" + "foo bar") {
                                                    case "foo bar":
                                                        var v130 int64 = 662
                                                        _ = v130
                                                    }
                                                } else {
                                                    var v131 int16 = (v117 >> ((0) & 63))
                                                    _ = v131
                                                }
                                                _ = v118
                                            case -9223372036854775808:
                                                if ((float64(0.0) - (1e10 / 0.1)) != 0.0) {
                                                    var v132 uint = 42
                                                    _ = v132
                                                }
                                            case -2147483648:
                                                v133 := GenS7{F0: genFn0(432), F1: (fi107 | (fi107 ^ (fi109 / ((int(v46)) | 1)))), F2: -910}
                                                _ = v133
                                            default:
                                                v134 := GenS6{F0: true, F1: 106, F2: (42 == 42)}
                                                _ = v134
                                            }
                                            v135 := GenS14{F0: 2147483647, F1: 2147483647}
                                            v136 := GenS13{F0: 961}
                                            v11 = v11
                                            if false {
                                                _ = len(v1)
                                                delete(v1, ("" + ("" + "\n\t")))
                                            }
                                            v137 := GenS15{F0: 9223372036854775807, F1: (v46 | 42), F2: "a"}
                                            fmt.Println(v11)
                                            v138 := GenS16{F0: 0, F1: v117}
                                            _ = v135
                                            _ = v136
                                            _ = v137
                                            _ = v138
                                        }
                                        if false {
                                            v3 = append(v3, 662)
                                        } else {
                                            var v139 string = ((((("hello" + "\n\t") + "a") + "\n\t") + "\n\t") + "a")
                                            var v140 uint32 = 2
                                            var v141 int64 = 703
                                            _ = v139
                                            _ = v140
                                            _ = v141
                                        }
                                        _ = v117
                                    } else {
                                        fmt.Println(fi109, v46)
                                    }
                                    switch fi109 {
                                    case -9223372036854775808:
                                        delete(v1, "hello")
                                    case 42:
                                        v11 = v46
                                    }
                                    switch (("a" + "a") + "foo bar") {
                                    case "hello":
                                        v142 := GenS13{F0: 2}
                                        _ = v142
                                    case "":
                                        var v143 bool = (65535 > 447)
                                        _ = v143
                                    case "a":
                                        v144 := GenS11{F0: -2147483648}
                                        _ = v144
                                    }
                                }
                            case "hello":
                                var v145 int8 = (v46 | (^(v11)))
                                _ = v145
                            default:
                                var v146 uint16 = 0
                                _ = v146
                            }
                            switch "a" {
                            case "\n\t":
                                var v147 bool = (false && (671 < 4294967295))
                                _ = v147
                            case "hello":
                                var v148 int32 = -2147483648
                                _ = v148
                            default:
                                var v149 uint = 4294967295
                                _ = v149
                            }
                        case -429:
                            v11 = (v46 / (((^(v11))) | 1))
                        case -886:
                            var v150 int = fi107
                            _ = v150
                        default:
                            if (-600 < 432) {
                                switch "\n\t" {
                                case "\n\t":
                                    _ = len(v3)
                                }
                                v46 = v46
                            } else {
                                v11 += (v46 >> ((4294967295) & 63))
                                v1[(("foo bar" + "\n\t") + ("hello" + "foo bar"))] = (-(fi107))
                                if ((255 <= 1) && (0.0 >= 0.0)) {
                                    v151 := GenS7{F0: genFn0(296), F1: (fi109 & fi109), F2: 1}
                                    var v152 int64 = 42
                                    switch (fi109 ^ fi109) {
                                    case -2147483648:
                                        v46 = v11
                                    case 128:
                                        var v153 uint = 0
                                        _ = v153
                                    }
                                    var v154 uint16 = 2
                                    v155 := GenS17{F0: 42, F1: (v11 % ((v46) | 1)), F2: true}
                                    _ = len(v1)
                                    _ = v151
                                    _ = v152
                                    _ = v154
                                    _ = v155
                                } else {
                                    fmt.Println(v11, fi109)
                                    v156 := GenS13{F0: 1}
                                    _ = v156
                                }
                            }
                            if (0 != 127) {
                                v46 = (v46 % (((v46 + (v11 << ((4294967295) & 63)))) | 1))
                                v157 := GenS5{F0: "\n\t", F1: 101}
                                if ((false || (0 >= -1)) && (0.1 != float32((1.5 + 3.14159)))) {
                                    var v158 uint16 = 920
                                    if (!false) {
                                        v158 = v158
                                        v1[("" + "")] = (fi109 ^ 1)
                                        var v159 int = 308
                                        _ = v159
                                    }
                                    _ = v158
                                }
                                _ = v157
                            }
                            var v160 uint64 = 0
                            v160 = (v160 / ((v160) | 1))
                            v161 := v1["\n\t"]
                            _ = v160
                            _ = v161
                        }
                        v11 = 0
                        _ = v111
                    }
                }
                _ = fi109
            }
        }
        fmt.Println(v46, v11)
        _ = fi107
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi162 := 0; fi162 < 2; fi162++ {
        var v163 int64 = -1
        _ = fi162
        _ = v163
    }
    v2 = append(v2, 1)
    fmt.Println(v46, v46)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v46
}
