package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int16; F1 int8; F2 uint16 }
type GenS4 struct { F0 int8; F1 bool; F2 uint8 }
type GenS5 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 int8; F1 int16; F2 bool }
type GenS11 struct { F0 int8; F1 bool; F2 int8 }
type GenS12 struct { F0 uint32; F1 string; F2 uint16 }
type GenS13 struct { F0 int; F1 uint64; F2 int8 }
type GenS14 struct { F0 int64; F1 int16 }
type GenS15 struct { F0 uint16; F1 string; F2 int16 }
type GenS16 struct { F0 int16; F1 uint32; F2 int }
func genFn0(p0 bool, p1 uint64) int {
    if true {
        v4 := GenS1{F0: 4294967295, F1: "foo bar", F2: int8(p1)}
        _ = v4
    } else {
        var v5 int = 1
        switch "" {
        case "foo bar":
            v6 := GenS2{F0: (("" + "foo bar") + "\n\t"), F1: 9223372036854775807}
            _ = v6
        case "":
            switch (-(v5)) {
            case 498:
                switch "foo bar" {
                case "foo bar":
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        for fi7 := 0; fi7 < 13; fi7++ {
                            fmt.Println(p0, p0)
                            v8 := GenS3{F0: -1, F1: -1, F2: 65535}
                            _ = fi7
                            _ = v8
                        }
                    case "a":
                        var v9 int64 = -309
                        _ = v9
                    default:
                        if p0 {
                            if ((!p0) || p0) {
                                v10 := []int{(v5 >> ((2) & 63)), (v5 - v5)}
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                var v11 float64 = 0.1
                                _ = v10
                                _ = v11
                            } else {
                                for fi12 := 0; fi12 < 2; fi12++ {
                                    v13 := []int{}
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v14 float64 = 3.14159
                                        _ = v14
                                    case "\n\t":
                                        switch ("a" + (("foo bar" + "a") + "")) {
                                        case "\n\t":
                                            var v15 uint16 = 874
                                            _ = v15
                                        case "":
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v16 float32 = ((3.14159 + 0.0) - (3.14159 + 3.14159))
                                                _ = v16
                                            case "foo bar":
                                                var v17 uint16 = 42
                                                _ = v17
                                            }
                                        case "foo bar":
                                            if false {
                                                var v18 int = (v5 - 42)
                                                _ = len(v13)
                                                var v19 int32 = 0
                                                _ = v18
                                                _ = v19
                                            }
                                        }
                                    }
                                    _ = fi12
                                    _ = v13
                                }
                                v5 = v5
                                p0 = true
                            }
                        } else {
                            v20 := GenS4{F0: 93, F1: (((!(-1 == -128)) && (p0 && p0)) || false), F2: 0}
                            var v21 float32 = 1e10
                            v22 := GenS1{F0: uint32(p1), F1: ("\n\t" + "a"), F2: 0}
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    }
                case "\n\t":
                    for fi23 := 0; fi23 < 18; fi23++ {
                        if p0 {
                            var v24 uint32 = 349
                            _ = v24
                        }
                        p0 = p0
                        _ = fi23
                    }
                case "a":
                    if p0 {
                        var v25 uint32 = 4294967295
                        _ = v25
                    } else {
                        v26 := []int{(v5 >> ((255) & 63)), (^(v5)), (v5 % (((v5 | (^(v5)))) | 1))}
                        v27 := GenS5{F0: 102, F1: 255, F2: 32767}
                        var v28 bool = p0
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                default:
                    switch v5 {
                    case -1:
                        v29 := []int{}
                        _ = v29
                    case -9223372036854775808:
                        v30 := []int{(v5 & (v5 / ((v5) | 1))), (v5 + v5), (v5 >> ((0) & 63))}
                        _ = v30
                    case 42:
                        p0 = (1 != 1000)
                    }
                }
            }
        case "hello":
            v31 := GenS2{F0: ("a" + ("foo bar" + (("\n\t" + "") + "hello"))), F1: 1}
            _ = v31
        default:
            var v32 string = "a"
            _ = v32
        }
        _ = v5
    }
    v33 := []int{42}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v33
    return -626
}
func genFn1(p0 int16) int64 {
    switch ((((("hello" + "hello") + ("" + "\n\t")) + ("\n\t" + "a")) + "foo bar") + "hello") {
    case "foo bar":
        var v34 uint = 682
        _ = v34
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            if (false || (-2147483648 != 2147483647)) {
                var v35 int32 = -30
                _ = v35
            } else {
                var v36 bool = false
                if (-508 != 0) {
                    p0 = -32768
                    v37 := []int{genFn0((v36 || v36), 0), genFn0(true, 18446744073709551615)}
                    _ = v37
                }
                _ = v36
            }
        }
    default:
        var v38 float32 = ((3.14159 * float32((0.0 / ((-2.25 - 3.14159) + -2.25)))) * 1e10)
        _ = v38
    }
    v39 := GenS6{F0: 1, F1: 2147483648}
    for fi40 := 0; fi40 < 18; fi40++ {
        v41 := GenS4{F0: 78, F1: ("hello" <= (("a" + "\n\t") + ("\n\t" + ("hello" + "\n\t")))), F2: 6}
        _ = fi40
        _ = v41
    }
    v42 := []int{1}
    _ = p0
    _ = v39
    _ = v42
    return -1
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{919}
    v3 := []uint32{1, 10, 1}
    v43 := v1["a"]
    if len(v3) > 0 {
        v3[0] = 65535
    }
    v44 := v1["hello"]
    v43 = (v44 + "hello")
    if true {
        for fi45 := 0; fi45 < 13; fi45++ {
            v44 = "hello"
            v44 = ((v44 + v43) + (v43 + "\n\t"))
            _ = fi45
        }
        if len(v2) > 0 {
            v2[0] = -2147483648
        }
    }
    switch 32767 {
    case 127:
        v46 := GenS5{F0: 872, F1: 32767, F2: -481}
        _ = v46
    case 9223372036854775807:
        for fi47 := 0; fi47 < 5; fi47++ {
            var v48 float64 = (((3.14159 * (1e10 * (float64(0.1) * (1.5 + -2.25)))) * 0.0) / 3.14159)
            if ((!(!(!(true && false)))) && (false || true)) {
                for fi49 := 0; fi49 < 9; fi49++ {
                    if true {
                        switch 128 {
                        case -9223372036854775808:
                            var v50 bool = true
                            _ = v50
                        case -892:
                            if true {
                                if ((0 < 1) || false) {
                                    _ = len(v3)
                                    if (!(!((v43 + "hello") > ("\n\t" + ("" + "foo bar"))))) {
                                        v44 = v44
                                        v51 := GenS1{F0: 4294967295, F1: "", F2: 79}
                                        _ = v51
                                    } else {
                                        var v52 uint8 = 255
                                        _ = v52
                                    }
                                } else {
                                    v43 = "\n\t"
                                }
                                if true {
                                    v44 = "foo bar"
                                }
                                var v53 int16 = 1
                                _ = v53
                            }
                        case 159:
                            v43 = ((v43 + v44) + (("a" + v44) + (("foo bar" + v44) + "foo bar")))
                        }
                        v48 = (float64(3.14159) + v48)
                        v43 = v44
                    } else {
                        v54 := GenS7{F0: false}
                        _ = v54
                    }
                    var v55 float32 = -2.25
                    switch fi49 {
                    case 4294967295:
                        var v56 float64 = v48
                        _ = v56
                    case -136:
                        var v57 uint32 = 409
                        _ = v57
                    }
                    _ = fi49
                    _ = v55
                }
                if ((!true) || false) {
                    v48 -= (0.0 * -2.25)
                } else {
                    var v58 bool = (float32((((0.1 / 1e10) - (v48 * (v48 * ((((v48 * (v48 - v48)) + float64(-2.25)) - v48) + v48)))) / -2.25)) < (0.0 * -2.25))
                    _ = v58
                }
                if (!false) {
                    _ = len(v1)
                    for fi59 := 0; fi59 < 12; fi59++ {
                        var v60 int32 = 42
                        v43 = v44
                        _ = fi59
                        _ = v60
                    }
                }
            }
            _ = fi47
            _ = v48
        }
    case -749:
        for fi61 := 0; fi61 < 5; fi61++ {
            for fi62 := 0; fi62 < 4; fi62++ {
                if ((-2.25 + 1.5) <= (1e10 - 0.1)) {
                    switch "foo bar" {
                    case "a":
                        var v63 float64 = 0.1
                        _ = v63
                    case "\n\t":
                        fmt.Println(fi62, v44, fi61)
                    case "hello":
                        if (true || true) {
                            fmt.Println(fi61, v44)
                            var v64 string = v43
                            _ = v64
                        }
                    default:
                        v43 = v44
                    }
                    if false {
                        v65 := GenS7{F0: true}
                        _ = v65
                    } else {
                        var v66 uint16 = 846
                        switch "a" {
                        case "foo bar":
                            if false {
                                var v67 bool = (100 <= -128)
                                if v67 {
                                    v3 = append(v3, 4294967295)
                                    switch fi62 {
                                    case 128:
                                        v68 := v1[(v43 + v43)]
                                        _ = v68
                                    }
                                    var v69 int = fi61
                                    _ = v69
                                } else {
                                    if (1.5 < float32(0.1)) {
                                        switch (fi62 << ((0) & 63)) {
                                        case -544:
                                            fmt.Println(v67, fi62, v43)
                                        case 42:
                                            if (!(42 == 127)) {
                                                delete(v1, v43)
                                                v70 := GenS4{F0: -128, F1: v67, F2: 127}
                                                v43 = v43
                                                if v67 {
                                                    var v71 int = 4294967295
                                                    var v72 string = v44
                                                    _ = v71
                                                    _ = v72
                                                } else {
                                                    var v73 int16 = 0
                                                    switch ((("\n\t" + v43) + v43) + "") {
                                                    case "\n\t":
                                                        var v74 int8 = 51
                                                        _ = v74
                                                    case "foo bar":
                                                        v43 = v44
                                                    case "":
                                                        switch ("\n\t" + (v44 + ((v44 + (v44 + v43)) + (v43 + v43)))) {
                                                        case "\n\t":
                                                            var v75 int16 = v73
                                                            _ = v75
                                                        }
                                                    default:
                                                        fmt.Println(v43, v73, v44)
                                                    }
                                                    switch v44 {
                                                    case "":
                                                        v73 -= v73
                                                    case "\n\t":
                                                        _ = len(v3)
                                                    default:
                                                        if v67 {
                                                            v67 = (1e10 >= (0.1 / float64((-2.25 * -2.25))))
                                                        }
                                                    }
                                                    _ = v73
                                                }
                                                _ = v70
                                            } else {
                                                _ = len(v1)
                                            }
                                        case 127:
                                            v67 = v67
                                        }
                                        switch (v44 + v44) {
                                        case "a":
                                            switch (v43 + ("foo bar" + ("a" + v43))) {
                                            case "foo bar":
                                                fmt.Println(v67, v44)
                                            case "hello":
                                                if false {
                                                    v76 := GenS1{F0: 2, F1: (v43 + v44), F2: 106}
                                                    switch (fi62 & (fi61 ^ (fi62 + -328))) {
                                                    case 42:
                                                        var v77 int64 = -1
                                                        _ = v77
                                                    case 128:
                                                        v78 := v1["\n\t"]
                                                        _ = v78
                                                    case 32767:
                                                        if v67 {
                                                            var v79 float64 = 3.14159
                                                            _ = v79
                                                        }
                                                    default:
                                                        if len(v2) > 0 {
                                                            v2[0] = -933
                                                        }
                                                    }
                                                    _ = v76
                                                }
                                            }
                                        case "foo bar":
                                            var v80 int64 = 0
                                            _ = v80
                                        }
                                    }
                                }
                                _ = v67
                            }
                        case "\n\t":
                            switch 128 {
                            case 127:
                                v1[v43] = v43
                            }
                            v43 = v43
                        case "":
                            var v81 int = fi61
                            _ = v81
                        default:
                            fmt.Println(v44)
                        }
                        _ = v66
                    }
                }
                _ = len(v3)
                _ = fi62
            }
            fmt.Println(fi61, v43, v44)
            _ = fi61
        }
    }
    switch "\n\t" {
    case "foo bar":
        if (-9223372036854775808 > -1) {
            var v82 uint = 255
            for fi83 := 0; fi83 < 17; fi83++ {
                if false {
                    v84 := GenS2{F0: "foo bar", F1: 1}
                    var v85 uint8 = 2
                    v86 := GenS4{F0: 42, F1: (v85 == (v85 + v85)), F2: v85}
                    _ = v84
                    _ = v85
                    _ = v86
                }
                _ = fi83
            }
            _ = len(v1)
            _ = v82
        } else {
            _ = len(v3)
        }
    default:
        if true {
            v44 = v44
            if true {
                v87 := GenS6{F0: -1, F1: 2}
                v43 = (((v43 + v43) + (("foo bar" + (((("" + v43) + "hello") + "hello") + ((("" + "foo bar") + ("foo bar" + "a")) + "a"))) + v43)) + (v44 + ("a" + "foo bar")))
                v88 := GenS4{F0: -128, F1: (-32768 == -32768), F2: 42}
                _ = v87
                _ = v88
            }
            switch (v43 + v44) {
            case "foo bar":
                v89 := v1["hello"]
                _ = v89
            default:
                var v90 bool = false
                _ = v90
            }
        } else {
            for fi91 := 0; fi91 < 7; fi91++ {
                switch (fi91 & fi91) {
                case 856:
                    var v92 int32 = 600
                    _ = v92
                case -9223372036854775808:
                    var v93 string = v43
                    _ = v93
                default:
                    if (969 < -32768) {
                        switch "foo bar" {
                        case "foo bar":
                            if false {
                                var v94 int64 = -696
                                _ = v94
                            }
                        default:
                            if len(v2) > 0 {
                                v2[0] = fi91
                            }
                        }
                        fmt.Println(v43, v43, fi91)
                    }
                }
                v44 = v43
                _ = fi91
            }
            fmt.Println(v44)
        }
    }
    v1[v44] = "\n\t"
    for fi95 := 0; fi95 < 5; fi95++ {
        if true {
            switch fi95 {
            case -617:
                var v96 uint32 = 647
                _ = v96
            case 6:
                if (!true) {
                    delete(v1, ("" + v44))
                    for fi97 := 0; fi97 < 14; fi97++ {
                        _ = len(v1)
                        v44 = ""
                        v98 := GenS7{F0: (1.5 == 3.14159)}
                        v99 := GenS3{F0: 1000, F1: -1, F2: 32767}
                        v1[v44] = v44
                        _ = fi97
                        _ = v98
                        _ = v99
                    }
                }
            case 32767:
                switch (v44 + v44) {
                case "a":
                    v100 := GenS8{F0: 42}
                    _ = v100
                case "hello":
                    fmt.Println(v44)
                case "\n\t":
                    var v101 int32 = 595
                    _ = v101
                default:
                    v2 = append(v2, fi95)
                }
            default:
                v102 := GenS5{F0: 140, F1: 886, F2: -1}
                _ = v102
            }
            fmt.Println(fi95, fi95)
        } else {
            v44 = v43
            switch (^(fi95)) {
            case 4294967295:
                v103 := GenS8{F0: -1}
                _ = v103
            }
            for fi104 := 0; fi104 < 13; fi104++ {
                v44 = ""
                switch v43 {
                case "":
                    if false {
                        var v105 int16 = 0
                        _ = v105
                    }
                    var v106 bool = true
                    v44 = v44
                    _ = v106
                case "\n\t":
                    var v107 int8 = 26
                    _ = v107
                case "a":
                    v43 = ((((v43 + v44) + v43) + "") + (v43 + v43))
                }
                _ = fi104
            }
        }
        _ = fi95
    }
    fmt.Println(v44)
    v108 := GenS9{F0: -1, F1: 416}
    var v109 float32 = 1.5
    var v110 int64 = -724
    switch "foo bar" {
    case "\n\t":
        v111 := GenS10{F0: -128, F1: 32767, F2: false}
        _ = v111
    default:
        v112 := GenS11{F0: 100, F1: (v44 > ("" + (v43 + "foo bar"))), F2: 100}
        _ = v112
    }
    var v113 string = "hello"
    for fi114 := 0; fi114 < 9; fi114++ {
        switch (v43 + "foo bar") {
        case "hello":
            for fi115 := 0; fi115 < 10; fi115++ {
                if false {
                    if true {
                        v44 = (v113 + v43)
                        var v116 int = (fi115 | 9223372036854775807)
                        fmt.Println(v110, v116, v110)
                        _ = v116
                    }
                    v113 = v113
                }
                _ = fi115
            }
        case "\n\t":
            var v117 float64 = float64((v109 - 1e10))
            _ = v117
        }
        _ = len(v2)
        for fi118 := 0; fi118 < 6; fi118++ {
            v44 = ("foo bar" + "hello")
            var v119 bool = (4294967295 == 18446744073709551615)
            _ = fi118
            _ = v119
        }
        _ = fi114
    }
    if false {
        if (v110 >= genFn1(-32768)) {
            v110 += v110
        }
        if false {
            v120 := GenS12{F0: 2, F1: v43, F2: 759}
            if (1 < 42) {
                switch -9223372036854775808 {
                case 639:
                    var v121 float32 = float32(0.1)
                    _ = v121
                case 679:
                    for fi122 := 0; fi122 < 2; fi122++ {
                        v123 := GenS13{F0: 127, F1: 0, F2: 42}
                        _ = fi122
                        _ = v123
                    }
                case 0:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    var v124 uint8 = 191
                    _ = v124
                }
            } else {
                var v125 int32 = 564
                switch 32768 {
                case -9223372036854775808:
                    if ((0 < 4294967295) || ((v110 * -1) != (v110 % (((-(v110))) | 1)))) {
                        var v126 int8 = 59
                        fmt.Println(v125, v125, v43)
                        _ = v126
                    } else {
                        if (65535 != 18446744073709551615) {
                            switch v43 {
                            case "foo bar":
                                v127 := GenS11{F0: -128, F1: (5 <= 127), F2: 42}
                                _ = v127
                            }
                        } else {
                            switch -978 {
                            case 32768:
                                for fi128 := 0; fi128 < 15; fi128++ {
                                    var v129 int16 = 90
                                    var v130 int16 = (v129 * 1000)
                                    var v131 bool = (v125 == (v125 ^ 1))
                                    _ = fi128
                                    _ = v129
                                    _ = v130
                                    _ = v131
                                }
                            case -2147483648:
                                switch v43 {
                                case "":
                                    var v132 uint32 = 2147483648
                                    _ = v132
                                case "hello":
                                    for fi133 := 0; fi133 < 13; fi133++ {
                                        switch fi133 {
                                        case 269:
                                            v134 := GenS4{F0: 0, F1: (32767 == 1000), F2: 0}
                                            _ = v134
                                        }
                                        _ = fi133
                                    }
                                default:
                                    if false {
                                        if (!(32767 < 65535)) {
                                            var v135 bool = (!false)
                                            var v136 uint64 = 811
                                            if (v135 || ((((3.14159 * 1.5) / (1.5 * -2.25)) * ((0.1 * float64((v109 / (v109 - v109)))) / 0.1)) == 1e10)) {
                                                if (!false) {
                                                    v3 = append(v3, 2)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v137 uint32 = 0
                                                    _ = v137
                                                }
                                            }
                                            _ = v135
                                            _ = v136
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            switch v43 {
                                            case "foo bar":
                                                for fi138 := 0; fi138 < 19; fi138++ {
                                                    v139 := GenS14{F0: (-(v110)), F1: 969}
                                                    if len(v3) > 0 {
                                                        v3[0] = 866
                                                    }
                                                    _ = fi138
                                                    _ = v139
                                                }
                                            case "":
                                                for fi140 := 0; fi140 < 10; fi140++ {
                                                    if false {
                                                        v141 := GenS10{F0: -128, F1: 120, F2: false}
                                                        _ = v141
                                                    }
                                                    var v142 uint = 655
                                                    for fi143 := 0; fi143 < 17; fi143++ {
                                                        _ = len(v2)
                                                        delete(v1, "foo bar")
                                                        _ = fi143
                                                    }
                                                    _ = fi140
                                                    _ = v142
                                                }
                                            case "\n\t":
                                                var v144 uint8 = 0
                                                _ = v144
                                            default:
                                                v145 := v1["foo bar"]
                                                _ = v145
                                            }
                                            var v146 int8 = 1
                                            _ = v146
                                        }
                                        v43 = "\n\t"
                                        for fi147 := 0; fi147 < 1; fi147++ {
                                            v113 = v113
                                            if (!true) {
                                                delete(v1, (v113 + (v44 + v44)))
                                                if (100 <= 1) {
                                                    switch fi147 {
                                                    case 32767:
                                                        var v148 int32 = (v125 - (v125 * (v125 * -2147483648)))
                                                        _ = v148
                                                    case 233:
                                                        if true {
                                                            fmt.Println(fi147, v44)
                                                        } else {
                                                            var v149 float64 = 0.0
                                                            _ = v149
                                                        }
                                                    }
                                                } else {
                                                    switch v43 {
                                                    case "foo bar":
                                                        _ = len(v1)
                                                    }
                                                }
                                            }
                                            var v150 bool = true
                                            _ = fi147
                                            _ = v150
                                        }
                                    }
                                }
                            }
                            for fi151 := 0; fi151 < 20; fi151++ {
                                if (!(!false)) {
                                    v152 := GenS15{F0: 255, F1: "a", F2: -1}
                                    if false {
                                        fmt.Println(v110, v110)
                                        v44 = v44
                                        v153 := GenS9{F0: v110, F1: 345}
                                        _ = v153
                                    } else {
                                        _ = len(v1)
                                        if true {
                                            if (!(525 <= 32767)) {
                                                v113 = ""
                                            }
                                            v109 = (v109 / (v109 - v109))
                                        }
                                        v125 -= (v125 * (v125 - -1))
                                    }
                                    if (4294967295 == 4294967295) {
                                        for fi154 := 0; fi154 < 1; fi154++ {
                                            var v155 uint8 = 0
                                            _ = fi154
                                            _ = v155
                                        }
                                        var v156 int64 = (-(v110))
                                        _ = v156
                                    }
                                    _ = v152
                                }
                                var v157 int32 = -1
                                _ = fi151
                                _ = v157
                            }
                        }
                    }
                case -495:
                    v110 = 1
                case 127:
                    v109 -= float32(3.14159)
                default:
                    v158 := GenS3{F0: 941, F1: -1, F2: 0}
                    _ = v158
                }
                v43 = "\n\t"
                _ = v125
            }
            v159 := GenS8{F0: (v110 >> ((164) & 63))}
            _ = v120
            _ = v159
        } else {
            v160 := GenS16{F0: -462, F1: 102, F2: 220}
            _ = v160
        }
        v1[v44] = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v113:", v113)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v108
    _ = v109
    _ = v110
    _ = v113
    _ = v2
    _ = v3
    _ = v43
    _ = v44
}
