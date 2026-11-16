package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int16; F2 int32 }
type GenS4 struct { F0 uint8; F1 int32; F2 string }
type GenS5 struct { F0 int16; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 string; F1 int16; F2 int32 }
type GenS14 struct { F0 int64; F1 uint32 }
type GenS15 struct { F0 uint64; F1 uint64 }
type GenS16 struct { F0 bool; F1 uint8; F2 int8 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 int; F1 uint16; F2 uint16 }
type GenS21 struct { F0 int64; F1 bool; F2 uint64 }
func genFn0() int64 {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int32 = 42
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int64 = -408
        v8 += int64(fi7)
        v8 = (v8 & (v8 | v8))
        _ = fi7
        _ = v8
    }
    var v9 float64 = 0.0
    switch -367 {
    case 2147483647:
        var v10 uint64 = 397
        _ = v10
    }
    _ = v9
    return 224
}
func genFn1() int32 {
    if (-128 <= -26) {
        var v11 uint16 = 2
        _ = v11
    } else {
        v12 := GenS1{F0: 49}
        var v13 int = -633
        fmt.Println(v13)
        _ = v12
        _ = v13
    }
    var v14 int32 = 0
    switch 868 {
    case 714:
        fmt.Println(v14)
    case -946:
        if (false || (4294967295 != 4294967295)) {
            switch ("a" + "a") {
            case "hello":
                v15 := GenS2{F0: "hello"}
                _ = v15
            case "":
                var v16 uint8 = 127
                _ = v16
            default:
                v17 := []int{127, 9223372036854775807}
                _ = v17
            }
            for fi18 := 0; fi18 < 10; fi18++ {
                v14 += -991
                v14 -= 1
                if ((!true) || (((!("hello" != "foo bar")) && (42 >= 0)) || ((2147483647 <= 2147483648) && (v14 < (v14 * (v14 ^ v14)))))) {
                    if false {
                        if true {
                            switch "a" {
                            case "a":
                                if false {
                                    var v19 uint64 = 4294967295
                                    _ = v19
                                } else {
                                    var v20 bool = (true && true)
                                    _ = v20
                                }
                            }
                            for fi21 := 0; fi21 < 18; fi21++ {
                                v22 := []int{(fi18 + fi18), 1, fi18}
                                _ = fi21
                                _ = v22
                            }
                        }
                        if (!(!false)) {
                            v23 := []int{fi18}
                            _ = v23
                        } else {
                            v24 := GenS2{F0: "foo bar"}
                            if false {
                                var v25 uint16 = 255
                                _ = v25
                            }
                            _ = v24
                        }
                    } else {
                        v14 -= (v14 / (((v14 - v14)) | 1))
                        for fi26 := 0; fi26 < 10; fi26++ {
                            if (942 >= 1) {
                                var v27 uint8 = 250
                                if true {
                                    var v28 int32 = v14
                                    _ = v28
                                } else {
                                    var v29 int = fi18
                                    v27 += (v27 - (v27 / ((v27) | 1)))
                                    v30 := GenS1{F0: 0}
                                    _ = v29
                                    _ = v30
                                }
                                v31 := []int{(fi18 * (fi18 << ((18446744073709551615) & 63)))}
                                _ = v27
                                _ = v31
                            } else {
                                if false {
                                    v14 = (v14 & v14)
                                }
                                v14 = -908
                                if ((3.14159 / float64(-2.25)) <= -2.25) {
                                    switch 127 {
                                    case -998:
                                        if true {
                                            v32 := []int{127}
                                            _ = len(v32)
                                            _ = v32
                                        } else {
                                            v14 += (v14 / ((0) | 1))
                                        }
                                    }
                                    var v33 uint32 = 582
                                    switch "hello" {
                                    case "":
                                        var v34 bool = (!(65535 > 255))
                                        _ = v34
                                    case "a":
                                        var v35 uint = 255
                                        _ = v35
                                    case "foo bar":
                                        v33 -= 436
                                    default:
                                        switch fi18 {
                                        case -1:
                                            var v36 uint = 1
                                            _ = v36
                                        case 403:
                                            if (!(false || (1 > 255))) {
                                                v37 := []int{fi18, 32768, int(v14), (fi26 & (^(fi18)))}
                                                switch ("\n\t" + ("a" + "\n\t")) {
                                                case "hello":
                                                    v38 := GenS3{F0: -1, F1: 32767, F2: v14}
                                                    _ = v38
                                                case "\n\t":
                                                    _ = len(v37)
                                                case "a":
                                                    fmt.Println(fi18, v33)
                                                }
                                                _ = v37
                                            }
                                            fmt.Println(fi26)
                                            var v39 float64 = (1.5 + 1e10)
                                            _ = v39
                                        case 0:
                                            v33 = (v33 | v33)
                                        }
                                    }
                                    _ = v33
                                }
                            }
                            _ = fi26
                        }
                        v40 := []int{fi18, fi18}
                        _ = v40
                    }
                    var v41 int32 = (v14 | (v14 * 0))
                    _ = v41
                }
                _ = fi18
            }
            v42 := []int{-365, 42}
            _ = v42
        }
    case -2147483648:
        v14 = 42
    default:
        var v43 uint32 = 2147483648
        _ = v43
    }
    switch "\n\t" {
    case "a":
        var v44 float64 = 1.5
        _ = v44
    case "\n\t":
        for fi45 := 0; fi45 < 9; fi45++ {
            v46 := GenS4{F0: 141, F1: 42, F2: ("\n\t" + "hello")}
            _ = fi45
            _ = v46
        }
    default:
        var v47 bool = false
        _ = v47
    }
    _ = v14
    return (v14 - -353)
}

func main() {
    var v1 int = 157
    var v2 float32 = ((1e10 / 1e10) / (0.1 / (float32(float64(3.14159)) - ((1e10 / (1.5 - 0.0)) * ((3.14159 - float32(1e10)) / 1.0)))))
    v3 := make(map[uint32]int)
    var v4 int64 = 42
    v48 := []int{}
    fmt.Println(v1)
    switch (v1 / ((v1) | 1)) {
    case -129:
        if false {
            v4 = v4
            if len(v48) >= 2 {
                v48 = v48[1:len(v48)]
            }
        }
    case 4294967295:
        switch v1 {
        case 9223372036854775807:
            fmt.Println(v4, v1)
        case -2147483648:
            if (18446744073709551615 <= 65535) {
                switch v1 {
                case -1:
                    for fi49 := 0; fi49 < 2; fi49++ {
                        var v50 uint32 = 2
                        var v51 bool = (42 < 413)
                        for fi52 := 0; fi52 < 2; fi52++ {
                            if ((v51 && v51) || true) {
                                switch "" {
                                case "a":
                                    v53 := GenS4{F0: 174, F1: 42, F2: ((("" + "foo bar") + ("" + "hello")) + "\n\t")}
                                    _ = v53
                                default:
                                    v54 := GenS3{F0: v4, F1: -675, F2: -2147483648}
                                    _ = v54
                                }
                            }
                            v55 := GenS1{F0: 4294967295}
                            v56 := GenS5{F0: 805, F1: (v50 | v50), F2: 0}
                            _ = fi52
                            _ = v55
                            _ = v56
                        }
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                default:
                    v57 := GenS3{F0: (^(v4)), F1: 1000, F2: -143}
                    _ = v57
                }
                if true {
                    if false {
                        if true {
                            if true {
                                fmt.Println(v1, v2)
                            }
                        } else {
                            var v58 uint8 = 0
                            switch 1 {
                            case 2147483647:
                                var v59 uint = 65535
                                _ = v59
                            case 128:
                                var v60 int64 = -528
                                _ = v60
                            case 546:
                                fmt.Println(v2, v58, v2)
                            }
                            if (1.5 < v2) {
                                fmt.Println(v1, v58, v4)
                            } else {
                                if len(v48) >= 2 {
                                    v48 = v48[1:len(v48)]
                                }
                                switch "" {
                                case "foo bar":
                                    switch "hello" {
                                    case "a":
                                        v58 -= v58
                                    }
                                case "":
                                    for fi61 := 0; fi61 < 16; fi61++ {
                                        if true {
                                            v62 := GenS3{F0: -141, F1: 1, F2: -2147483648}
                                            v63 := GenS6{F0: 65535}
                                            _ = v62
                                            _ = v63
                                        } else {
                                            v3[828] = 127
                                            v2 = v2
                                            var v64 string = ("a" + "a")
                                            _ = v64
                                        }
                                        _ = fi61
                                    }
                                case "\n\t":
                                    fmt.Println(v1, v2)
                                default:
                                    v2 -= v2
                                }
                                if (-9223372036854775808 != 42) {
                                    var v65 uint8 = 0
                                    fmt.Println(v65, v4, v2)
                                    _ = v65
                                }
                            }
                            _ = v58
                        }
                        if (261 >= 4294967295) {
                            switch v1 {
                            case -1:
                                v66 := GenS6{F0: 236}
                                _ = v66
                            case 514:
                                v67 := GenS7{F0: -79}
                                _ = v67
                            default:
                                var v68 bool = ((!(-1 < -1)) && (false && true))
                                _ = v68
                            }
                            var v69 int16 = 32767
                            _ = len(v3)
                            _ = v69
                        }
                        if (0 < 4294967295) {
                            var v70 int = (v1 / (((v1 ^ (^(v1)))) | 1))
                            _ = v70
                        }
                    }
                    for fi71 := 0; fi71 < 3; fi71++ {
                        v72 := GenS8{F0: "\n\t"}
                        var v73 uint = 18446744073709551615
                        fmt.Println(v4, v1, fi71)
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                } else {
                    if true {
                        var v74 int64 = (v4 * (v4 + -1))
                        if false {
                            fmt.Println(v74, v74, v2)
                            fmt.Println(v1)
                            switch (("a" + "\n\t") + "foo bar") {
                            case "hello":
                                if (false || true) {
                                    var v75 int32 = -2147483648
                                    _ = v75
                                }
                            default:
                                fmt.Println(v2, v1, v74)
                            }
                        } else {
                            if (true && false) {
                                v76 := GenS5{F0: -32768, F1: 2147483647, F2: 1}
                                _ = v76
                            } else {
                                var v77 int8 = 1
                                _ = v77
                            }
                        }
                        switch v1 {
                        case 0:
                            fmt.Println(v4, v2)
                        default:
                            v78 := v3[65535]
                            _ = v78
                        }
                        _ = v74
                    } else {
                        var v79 uint64 = 672
                        _ = v79
                    }
                    for fi80 := 0; fi80 < 1; fi80++ {
                        v81 := GenS2{F0: "\n\t"}
                        v2 -= (-2.25 / 1.5)
                        _ = fi80
                        _ = v81
                    }
                    switch 2147483647 {
                    case 127:
                        var v82 bool = (2 >= 2147483647)
                        _ = v82
                    }
                }
                if (!(86 != 128)) {
                    if false {
                        for fi83 := 0; fi83 < 14; fi83++ {
                            if len(v48) >= 2 {
                                v48 = v48[1:len(v48)]
                            }
                            for fi84 := 0; fi84 < 17; fi84++ {
                                v85 := GenS2{F0: ("" + ("a" + ("hello" + "\n\t")))}
                                _ = fi84
                                _ = v85
                            }
                            v86 := v3[240]
                            _ = fi83
                            _ = v86
                        }
                        var v87 int8 = 106
                        v1 += 97
                        _ = v87
                    }
                    switch 127 {
                    case 372:
                        switch ("a" + ((("" + ("hello" + "a")) + "hello") + "\n\t")) {
                        case "a":
                            if (42 == 127) {
                                var v88 uint64 = 18446744073709551615
                                _ = v88
                            } else {
                                for fi89 := 0; fi89 < 16; fi89++ {
                                    if true {
                                        if ((2147483647 < 669) || true) {
                                            var v90 uint16 = 32767
                                            switch (fi89 - v1) {
                                            case -381:
                                                for fi91 := 0; fi91 < 6; fi91++ {
                                                    v90 = v90
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v3[4294967295] = 1
                                                    }
                                                    v92 := GenS1{F0: 0}
                                                    _ = fi91
                                                    _ = v92
                                                }
                                            }
                                            _ = v90
                                        }
                                        delete(v3, 1)
                                    } else {
                                        v48 = append(v48, v1)
                                    }
                                    var v93 int16 = 0
                                    _ = fi89
                                    _ = v93
                                }
                                v94 := GenS8{F0: ("" + "a")}
                                _ = v94
                            }
                        case "":
                            var v95 bool = true
                            _ = v95
                        case "foo bar":
                            switch "\n\t" {
                            case "foo bar":
                                for fi96 := 0; fi96 < 10; fi96++ {
                                    var v97 string = "a"
                                    for fi98 := 0; fi98 < 17; fi98++ {
                                        v99 := GenS8{F0: "foo bar"}
                                        v100 := GenS1{F0: 2}
                                        var v101 uint64 = 42
                                        _ = fi98
                                        _ = v100
                                        _ = v101
                                        _ = v99
                                    }
                                    for fi102 := 0; fi102 < 15; fi102++ {
                                        if (!true) {
                                            switch v1 {
                                            case 32768:
                                                if (!false) {
                                                    switch -9223372036854775808 {
                                                    case 33:
                                                        var v103 string = ("foo bar" + (v97 + ((v97 + ((v97 + v97) + ("foo bar" + ("foo bar" + v97)))) + v97)))
                                                        v104 := GenS7{F0: 45}
                                                        delete(v3, 99)
                                                        _ = v103
                                                        _ = v104
                                                    case -590:
                                                        if (4294967295 >= 2) {
                                                            var v105 int32 = 1
                                                            _ = v105
                                                        } else {
                                                            var v106 bool = ((65535 != 1) || false)
                                                            v3[1] = fi102
                                                            _ = v106
                                                        }
                                                    case 516:
                                                        var v107 uint16 = 42
                                                        _ = v107
                                                    }
                                                    v108 := GenS3{F0: (v4 * 1), F1: -1, F2: -1}
                                                    var v109 uint = 23
                                                    var v110 int = 32767
                                                    _ = v108
                                                    _ = v109
                                                    _ = v110
                                                }
                                            case -806:
                                                switch (((v97 + ((v97 + "a") + v97)) + v97) + (v97 + (v97 + v97))) {
                                                case "a":
                                                    fmt.Println(v4, v2)
                                                case "hello":
                                                    if true {
                                                        if true {
                                                            var v111 float32 = (float32(((0.0 + 0.1) + 0.0)) + -2.25)
                                                            v1 = fi96
                                                            if (!(417 < 4294967295)) {
                                                                v112 := GenS7{F0: -1}
                                                                v4 += v4
                                                                v1 = 127
                                                                _ = v112
                                                            }
                                                            _ = v111
                                                        }
                                                        var v113 int64 = 42
                                                        _ = v113
                                                    } else {
                                                        v1 += 4294967295
                                                        v114 := GenS5{F0: 0, F1: 42, F2: 1}
                                                        _ = v114
                                                    }
                                                }
                                            case 398:
                                                var v115 float32 = float32(3.14159)
                                                _ = v115
                                            default:
                                                switch (fi102 | 1) {
                                                case 9223372036854775807:
                                                    if ((false || false) || false) {
                                                        var v116 uint8 = 255
                                                        if (v4 <= -665) {
                                                            v117 := GenS5{F0: -32768, F1: 841, F2: 496}
                                                            _ = len(v48)
                                                            _ = v117
                                                        } else {
                                                            if (false && (42 > (v116 >> ((973) & 63)))) {
                                                                v2 = v2
                                                                switch v97 {
                                                                case "foo bar":
                                                                    switch 127 {
                                                                    case -345:
                                                                        var v118 uint = 18446744073709551615
                                                                        _ = v118
                                                                    }
                                                                case "hello":
                                                                    v97 = "a"
                                                                case "\n\t":
                                                                    fmt.Println(fi96)
                                                                default:
                                                                    v48 = append(v48, -725)
                                                                }
                                                            } else {
                                                                switch (^(fi96)) {
                                                                case -9223372036854775808:
                                                                    if (!true) {
                                                                        var v119 uint8 = v116
                                                                        _ = v119
                                                                    } else {
                                                                        switch -485 {
                                                                        case 1:
                                                                            var v120 uint = 29
                                                                            _ = v120
                                                                        case 42:
                                                                            v116 = v116
                                                                        case -169:
                                                                            var v121 uint32 = 4294967295
                                                                            _ = v121
                                                                        default:
                                                                            v97 = v97
                                                                            var v122 uint32 = 1
                                                                            _ = v122
                                                                        }
                                                                    }
                                                                case -2147483648:
                                                                    if false {
                                                                        v97 = "\n\t"
                                                                    }
                                                                case 1:
                                                                    var v123 string = v97
                                                                    _ = v123
                                                                }
                                                            }
                                                            v4 = (v4 % ((v4) | 1))
                                                        }
                                                        v4 -= v4
                                                        _ = v116
                                                    } else {
                                                        v97 = v97
                                                        if (v97 < v97) {
                                                            v97 = v97
                                                            v1 -= fi96
                                                            var v124 int8 = 100
                                                            switch v97 {
                                                            case "a":
                                                                switch v97 {
                                                                case "":
                                                                    fmt.Println(fi102)
                                                                case "hello":
                                                                    v125 := GenS2{F0: ("foo bar" + v97)}
                                                                    _ = v125
                                                                case "foo bar":
                                                                    var v126 float64 = float64(v2)
                                                                    _ = v126
                                                                default:
                                                                    switch fi102 {
                                                                    case 32768:
                                                                        var v127 int16 = 556
                                                                        _ = v127
                                                                    case -1:
                                                                        v128 := GenS9{F0: 42}
                                                                        _ = v128
                                                                    default:
                                                                        var v129 uint16 = 65535
                                                                        _ = v129
                                                                    }
                                                                }
                                                            case "hello":
                                                                switch fi96 {
                                                                case 128:
                                                                    v2 += ((float32((3.14159 * 1.5)) + (-2.25 * 1.5)) * 0.1)
                                                                case 42:
                                                                    _ = len(v3)
                                                                case -633:
                                                                    var v130 bool = (1 > 42)
                                                                    _ = v130
                                                                }
                                                            }
                                                            _ = v124
                                                        } else {
                                                            v2 = v2
                                                        }
                                                    }
                                                default:
                                                    if len(v48) >= 2 {
                                                        v48 = v48[1:len(v48)]
                                                    }
                                                }
                                            }
                                        } else {
                                            var v131 int8 = -1
                                            fmt.Println(v2)
                                            _ = v131
                                        }
                                        v132 := GenS10{F0: "hello"}
                                        _ = fi102
                                        _ = v132
                                    }
                                    _ = fi96
                                    _ = v97
                                }
                            case "a":
                                if true {
                                    for fi133 := 0; fi133 < 4; fi133++ {
                                        v3[0] = (-(v1))
                                        if (2147483648 >= 2147483647) {
                                            for fi134 := 0; fi134 < 12; fi134++ {
                                                var v135 int64 = v4
                                                _ = fi134
                                                _ = v135
                                            }
                                        } else {
                                            var v136 int = v1
                                            _ = v136
                                        }
                                        var v137 int16 = 0
                                        _ = fi133
                                        _ = v137
                                    }
                                    if false {
                                        v1 = (v1 ^ (v1 | 4294967295))
                                    }
                                    var v138 float64 = ((0.1 / -2.25) / 0.1)
                                    _ = v138
                                } else {
                                    if len(v48) >= 2 {
                                        v48 = v48[1:len(v48)]
                                    }
                                    fmt.Println(v1, v4, v1)
                                }
                            case "\n\t":
                                v2 = 0.0
                            default:
                                fmt.Println(v4)
                            }
                        }
                    }
                    _ = len(v48)
                }
            } else {
                var v139 string = ("a" + "hello")
                _ = v139
            }
        }
    case 9223372036854775807:
        fmt.Println(v1, v2, v2)
    default:
        var v140 float32 = v2
        _ = v140
    }
    if true {
        if true {
            for fi141 := 0; fi141 < 15; fi141++ {
                v2 -= -2.25
                if (-2.25 >= 1.5) {
                    switch (fi141 ^ v1) {
                    case 2147483647:
                        var v142 bool = false
                        _ = v142
                    case 9223372036854775807:
                        for fi143 := 0; fi143 < 11; fi143++ {
                            v4 += 42
                            v2 -= v2
                            fmt.Println(v2, v1)
                            v144 := GenS2{F0: ("foo bar" + ("" + "foo bar"))}
                            _ = fi143
                            _ = v144
                        }
                    default:
                        switch fi141 {
                        case 686:
                            v145 := GenS3{F0: -754, F1: 32767, F2: 261}
                            _ = v145
                        default:
                            v146 := GenS1{F0: 2}
                            _ = v146
                        }
                    }
                    var v147 uint = 18446744073709551615
                    _ = v147
                } else {
                    fmt.Println(v1, v2)
                }
                _ = fi141
            }
            v4 -= v4
            for fi148 := 0; fi148 < 2; fi148++ {
                switch (fi148 * (fi148 | (v1 << ((248) & 63)))) {
                case 1:
                    v149 := GenS8{F0: "hello"}
                    _ = v149
                case 371:
                    switch "a" {
                    case "\n\t":
                        for fi150 := 0; fi150 < 8; fi150++ {
                            if true {
                                v151 := GenS11{F0: (v4 | v4)}
                                switch (v1 ^ fi150) {
                                case 4294967295:
                                    v1 = v1
                                case 42:
                                    if true {
                                        if ((!false) || ((v4 & (v4 << ((0) & 63))) == (v4 % (((v4 / (((v4 + v4)) | 1))) | 1)))) {
                                            var v152 string = ("hello" + (("\n\t" + "\n\t") + "a"))
                                            v2 = v2
                                            var v153 int16 = 1
                                            _ = v152
                                            _ = v153
                                        } else {
                                            if (42 > 2) {
                                                var v154 uint = 255
                                                var v155 int8 = -12
                                                _ = v154
                                                _ = v155
                                            } else {
                                                if (!false) {
                                                    v4 = genFn0()
                                                    if (!true) {
                                                        var v156 int64 = (^(v4))
                                                        var v157 uint64 = 938
                                                        _ = v156
                                                        _ = v157
                                                    } else {
                                                        var v158 uint8 = 42
                                                        v4 = v4
                                                        _ = v158
                                                    }
                                                } else {
                                                    if ((fi148 & 42) >= 1) {
                                                        v159 := GenS9{F0: -970}
                                                        fmt.Println(fi148)
                                                        if ((v4 ^ v4) > v4) {
                                                            var v160 uint32 = 2
                                                            var v161 bool = (false && true)
                                                            _ = v160
                                                            _ = v161
                                                        } else {
                                                            var v162 uint32 = 2147483648
                                                            _ = v162
                                                        }
                                                        _ = v159
                                                    }
                                                }
                                            }
                                            if true {
                                                v1 += (^(fi150))
                                            } else {
                                                fmt.Println(v1, v4)
                                                if false {
                                                    v1 = 32767
                                                    var v163 uint8 = 162
                                                    v164 := GenS1{F0: 0}
                                                    _ = v163
                                                    _ = v164
                                                }
                                            }
                                            switch fi148 {
                                            case -2147483648:
                                                if ((v1 & fi148) < v1) {
                                                    v2 -= v2
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                        }
                                        var v165 uint32 = 2
                                        _ = v165
                                    }
                                }
                                _ = v151
                            } else {
                                v166 := GenS2{F0: (("hello" + (("foo bar" + "") + "hello")) + "")}
                                _ = v166
                            }
                            _ = fi150
                        }
                    case "a":
                        if true {
                            var v167 int8 = -1
                            _ = v167
                        }
                    default:
                        v168 := GenS12{F0: "\n\t"}
                        _ = v168
                    }
                default:
                    for fi169 := 0; fi169 < 13; fi169++ {
                        v170 := GenS13{F0: "hello", F1: 24, F2: -2147483648}
                        _ = fi169
                        _ = v170
                    }
                }
                _ = fi148
            }
        }
        v171 := GenS12{F0: ("a" + (("a" + ("foo bar" + (("\n\t" + "\n\t") + "foo bar"))) + ("\n\t" + "a")))}
        var v172 float64 = -2.25
        _ = v171
        _ = v172
    }
    for fi173 := 0; fi173 < 20; fi173++ {
        v2 = v2
        v4 = (v4 * (v4 ^ (v4 >> ((736) & 63))))
        v1 += (fi173 << ((566) & 63))
        _ = fi173
    }
    var v174 float64 = (float64((((v2 / 1.5) / (v2 / v2)) + v2)) * 0.0)
    switch (^(v1)) {
    case 769:
        if len(v48) > 0 {
            v48[0] = v1
        }
    case 32767:
        v48 = append(v48, (v1 + 42))
    case 0:
        for fi175 := 0; fi175 < 12; fi175++ {
            v48 = append(v48, (fi175 ^ fi175))
            _ = fi175
        }
    }
    for fi176 := 0; fi176 < 20; fi176++ {
        v177 := GenS11{F0: v4}
        fmt.Println(v2, v1, v1)
        _ = fi176
        _ = v177
    }
    v178 := GenS12{F0: "hello"}
    switch "hello" {
    case "a":
        if true {
            if (840 >= 2) {
                v179 := GenS14{F0: 1, F1: 4294967295}
                if ((false && false) || true) {
                    var v180 int64 = -284
                    _ = v180
                } else {
                    fmt.Println(v1, v2)
                    if (1 <= 814) {
                        if (652 < 0) {
                            v2 += 0.1
                            if len(v48) >= 2 {
                                v48 = v48[1:len(v48)]
                            }
                            for fi181 := 0; fi181 < 3; fi181++ {
                                fmt.Println(v2)
                                _ = fi181
                            }
                        } else {
                            v2 = v2
                            switch "" {
                            case "a":
                                fmt.Println(v174, v174, v1)
                            }
                            for fi182 := 0; fi182 < 5; fi182++ {
                                v183 := GenS15{F0: 2, F1: 248}
                                fmt.Println(v1, v2, v174)
                                v184 := GenS5{F0: 42, F1: 916, F2: 0}
                                _ = fi182
                                _ = v183
                                _ = v184
                            }
                        }
                        if false {
                            if false {
                                var v185 uint = 65535
                                for fi186 := 0; fi186 < 6; fi186++ {
                                    v174 = 0.0
                                    if false {
                                        switch "" {
                                        case "\n\t":
                                            v185 = v185
                                        }
                                        var v187 int64 = -726
                                        _ = v187
                                    }
                                    _ = fi186
                                }
                                v188 := GenS3{F0: (v4 << (((v185 | v185)) & 63)), F1: 32767, F2: 0}
                                _ = v185
                                _ = v188
                            } else {
                                v2 = v2
                                switch "foo bar" {
                                case "\n\t":
                                    if (("foo bar" + "foo bar") != "foo bar") {
                                        v174 = (v174 + v174)
                                        if (102 < 42) {
                                            if true {
                                                switch (((("" + "hello") + "a") + ("\n\t" + (("foo bar" + "") + "a"))) + "\n\t") {
                                                case "":
                                                    if len(v48) > 0 {
                                                        v48[0] = v1
                                                    }
                                                default:
                                                    switch "" {
                                                    case "":
                                                        switch 4294967295 {
                                                        case 2147483647:
                                                            switch -653 {
                                                            case 9223372036854775807:
                                                                var v189 int32 = 266
                                                                _ = v189
                                                            case 128:
                                                                v4 += (v4 | genFn0())
                                                            }
                                                        case 127:
                                                            if true {
                                                                switch "a" {
                                                                case "\n\t":
                                                                    v4 = (v4 | genFn0())
                                                                case "foo bar":
                                                                    var v190 uint16 = 990
                                                                    _ = v190
                                                                }
                                                                var v191 int64 = genFn0()
                                                                _ = v191
                                                            }
                                                        default:
                                                            v2 = 1e10
                                                        }
                                                    case "a":
                                                        v1 += v1
                                                    case "hello":
                                                        v174 += (v174 * 0.1)
                                                    default:
                                                        v192 := GenS13{F0: "", F1: 1000, F2: 26}
                                                        _ = v192
                                                    }
                                                }
                                            } else {
                                                var v193 uint8 = 202
                                                v1 = (v1 << ((65535) & 63))
                                                _ = len(v3)
                                                _ = v193
                                            }
                                        } else {
                                            v4 += 9223372036854775807
                                            if false {
                                                fmt.Println(v2, v174)
                                                for fi194 := 0; fi194 < 7; fi194++ {
                                                    fmt.Println(v4, fi194)
                                                    _ = fi194
                                                }
                                                for fi195 := 0; fi195 < 8; fi195++ {
                                                    for fi196 := 0; fi196 < 11; fi196++ {
                                                        v1 = (fi196 >> ((327) & 63))
                                                        v197 := GenS13{F0: (("\n\t" + "a") + ((("" + ((("\n\t" + "hello") + "a") + "")) + "foo bar") + "")), F1: 0, F2: -1}
                                                        v4 += 999
                                                        _ = fi196
                                                        _ = v197
                                                    }
                                                    var v198 int8 = 58
                                                    switch 9223372036854775807 {
                                                    case 32767:
                                                        switch (-(fi195)) {
                                                        case 32768:
                                                            if (4294967295 > 42) {
                                                                fmt.Println(v1)
                                                            } else {
                                                                var v199 int32 = -803
                                                                _ = v199
                                                            }
                                                        case 128:
                                                            for fi200 := 0; fi200 < 16; fi200++ {
                                                                _ = len(v48)
                                                                var v201 uint8 = 128
                                                                _ = fi200
                                                                _ = v201
                                                            }
                                                        default:
                                                            fmt.Println(v174)
                                                        }
                                                    case -1000:
                                                        var v202 string = "hello"
                                                        _ = v202
                                                    }
                                                    _ = fi195
                                                    _ = v198
                                                }
                                            }
                                        }
                                    }
                                case "":
                                    delete(v3, 502)
                                case "a":
                                    switch (v1 << ((0) & 63)) {
                                    case -9223372036854775808:
                                        if false {
                                            v203 := GenS16{F0: (18446744073709551615 <= 148), F1: 0, F2: 100}
                                            _ = v203
                                        } else {
                                            var v204 uint = 2
                                            var v205 int64 = (v4 << ((v204) & 63))
                                            v206 := GenS17{F0: (v204 + 2)}
                                            _ = v204
                                            _ = v205
                                            _ = v206
                                        }
                                    case 2147483647:
                                        v4 -= v4
                                    case 32768:
                                        for fi207 := 0; fi207 < 18; fi207++ {
                                            v3[0] = v1
                                            _ = fi207
                                        }
                                    default:
                                        for fi208 := 0; fi208 < 20; fi208++ {
                                            v1 = v1
                                            _ = fi208
                                        }
                                    }
                                }
                            }
                            if (float32(float64(((v2 / (-2.25 / (v2 * 3.14159))) + v2))) > 0.1) {
                                if ((v1 & v1) <= (v1 % ((61) | 1))) {
                                    v2 = ((v2 - v2) - v2)
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(v1)
                                    case "foo bar":
                                        v209 := GenS18{F0: 32767}
                                        _ = v209
                                    }
                                } else {
                                    var v210 uint64 = 2
                                    var v211 uint8 = 1
                                    if len(v48) > 0 {
                                        v48[0] = -2147483648
                                    }
                                    _ = v210
                                    _ = v211
                                }
                                v212 := GenS1{F0: 2}
                                _ = v212
                            } else {
                                var v213 int = (-(v1))
                                fmt.Println(v2)
                                _ = v213
                            }
                            _ = len(v48)
                        }
                    } else {
                        v174 = v174
                        for fi214 := 0; fi214 < 1; fi214++ {
                            var v215 float32 = 3.14159
                            v216 := GenS12{F0: "a"}
                            v2 -= 3.14159
                            _ = fi214
                            _ = v215
                            _ = v216
                        }
                    }
                    if true {
                        fmt.Println(v1, v2)
                        _ = len(v3)
                    } else {
                        var v217 int32 = -1
                        switch ("foo bar" + "a") {
                        case "foo bar":
                            var v218 uint16 = 1
                            _ = v218
                        case "":
                            v1 = (^(v1))
                        }
                        _ = v217
                    }
                }
                switch v1 {
                case -9223372036854775808:
                    switch "hello" {
                    case "":
                        if true {
                            var v219 int64 = v4
                            for fi220 := 0; fi220 < 2; fi220++ {
                                v174 = v174
                                _ = fi220
                            }
                            _ = v219
                        }
                    default:
                        v221 := GenS19{F0: (204 > 127)}
                        _ = v221
                    }
                case 128:
                    switch 128 {
                    case 32768:
                        if ((v1 / ((v1) | 1)) > v1) {
                            var v222 uint8 = 35
                            var v223 uint8 = (v222 - (v222 - v222))
                            switch "" {
                            case "foo bar":
                                v224 := GenS19{F0: (v222 >= v223)}
                                _ = v224
                            case "hello":
                                v225 := GenS20{F0: 32767, F1: 83, F2: 65535}
                                _ = v225
                            case "\n\t":
                                for fi226 := 0; fi226 < 15; fi226++ {
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v2, v2)
                                    default:
                                        v174 += (v174 + v174)
                                    }
                                    _ = fi226
                                }
                            default:
                                v227 := GenS20{F0: 127, F1: 65535, F2: 2}
                                _ = v227
                            }
                            _ = v222
                            _ = v223
                        } else {
                            v48 = append(v48, (v1 / (((v1 ^ v1)) | 1)))
                            v48 = append(v48, -1)
                            fmt.Println(v2)
                        }
                    case 9223372036854775807:
                        v228 := GenS21{F0: (v4 % ((v4) | 1)), F1: true, F2: 171}
                        _ = v228
                    default:
                        var v229 int8 = 100
                        _ = v229
                    }
                }
                _ = v179
            }
            if len(v48) > 0 {
                v48[0] = -9223372036854775808
            }
        } else {
            for fi230 := 0; fi230 < 17; fi230++ {
                v231 := GenS6{F0: 0}
                _ = fi230
                _ = v231
            }
        }
    case "\n\t":
        if (v1 < v1) {
            v1 -= v1
            var v232 int8 = -99
            v233 := GenS9{F0: 2147483647}
            _ = v232
            _ = v233
        }
    default:
        v48 = append(v48, v1)
    }
    switch (v1 | v1) {
    case -319:
        v3[197] = 128
    case 1:
        if ((v2 - 0.1) > 1.5) {
            v1 = -1
            var v234 int = v1
            _ = v234
        }
    default:
        var v235 string = "\n\t"
        _ = v235
    }
    fmt.Println(v174, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v174:", v174)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v48:", len(v48), v48)
    _ = v1
    _ = v174
    _ = v178
    _ = v2
    _ = v3
    _ = v4
    _ = v48
}
