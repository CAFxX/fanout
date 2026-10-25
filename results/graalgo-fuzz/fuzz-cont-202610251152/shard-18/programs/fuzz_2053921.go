package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint8; F2 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 int; F1 int; F2 bool }
type GenS10 struct { F0 uint8; F1 int; F2 int32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int64; F1 uint32; F2 uint64 }
func genFn0(p0 bool, p1 string, p2 uint64) int64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 string = ("hello" + p1)
        v7 = (p1 + "\n\t")
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        p1 = p1
        var v9 bool = true
        v10 := []int{fi8, 9223372036854775807, -2147483648, (fi8 - fi8)}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 float32 = 0.1
    v2 := []int{-327}
    var v3 float64 = 0.0
    var v4 uint32 = 936
    v5 := make(map[string]int)
    var v11 float32 = 1e10
    var v12 int16 = 32767
    var v13 string = (("hello" + "hello") + ("\n\t" + "\n\t"))
    if true {
        var v14 string = "\n\t"
        _ = v14
    }
    var v15 uint8 = 127
    switch v13 {
    case "hello":
        switch -2147483648 {
        case -9223372036854775808:
            switch v13 {
            case "hello":
                var v16 int32 = 400
                _ = v16
            }
        default:
            var v17 string = (("hello" + "") + v13)
            _ = v17
        }
    case "":
        var v18 int16 = v12
        _ = v18
    default:
        for fi19 := 0; fi19 < 16; fi19++ {
            switch v13 {
            case "\n\t":
                switch "\n\t" {
                case "\n\t":
                    var v20 int = fi19
                    _ = v20
                case "foo bar":
                    v21 := GenS1{F0: uint(v4), F1: (v12 << ((2) & 63))}
                    _ = v21
                case "hello":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    _ = len(v5)
                }
            case "a":
                for fi22 := 0; fi22 < 3; fi22++ {
                    var v23 uint = 255
                    v4 = (^(v4))
                    _ = fi22
                    _ = v23
                }
            }
            v11 -= v11
            for fi24 := 0; fi24 < 9; fi24++ {
                if len(v2) > 0 {
                    v2[0] = (fi19 * (fi24 >> ((4294967295) & 63)))
                }
                v15 -= v15
                switch (v13 + "\n\t") {
                case "\n\t":
                    switch (-(fi19)) {
                    case -2147483648:
                        if ((v15 * v15) > v15) {
                            var v25 string = v13
                            _ = v25
                        } else {
                            switch (fi24 * 32767) {
                            case -9223372036854775808:
                                var v26 uint16 = 255
                                _ = v26
                            case 674:
                                var v27 uint64 = 18446744073709551615
                                _ = v27
                            case 0:
                                fmt.Println(v4, fi24, fi19)
                            default:
                                var v28 int16 = 42
                                _ = v28
                            }
                            var v29 int8 = 100
                            _ = v29
                        }
                        switch fi19 {
                        case -549:
                            v30 := GenS2{F0: v4}
                            _ = v30
                        case 127:
                            if (747 < 928) {
                                var v31 uint8 = v15
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v32 uint8 = (v31 ^ (v31 | 221))
                                _ = v31
                                _ = v32
                            } else {
                                fmt.Println(v4, v4, v4)
                            }
                        case 956:
                            v33 := GenS2{F0: 2147483648}
                            var v34 string = ""
                            switch "foo bar" {
                            case "a":
                                var v35 int32 = 0
                                _ = v35
                            }
                            _ = v33
                            _ = v34
                        default:
                            v3 += (float64((float32(0.1) + 0.0)) + v3)
                        }
                    case -792:
                        v36 := GenS2{F0: (v4 | v4)}
                        _ = v36
                    case 4294967295:
                        var v37 bool = (!false)
                        v38 := GenS3{F0: 18446744073709551615}
                        _ = v37
                        _ = v38
                    default:
                        var v39 string = ("" + v13)
                        _ = v39
                    }
                }
                fmt.Println(v11)
                if ((true || false) || false) {
                    switch (fi24 % ((-83) | 1)) {
                    case 32767:
                        v11 = v11
                    case 280:
                        if ((v4 | (^(v4))) >= (v4 / (((v4 << ((42) & 63))) | 1))) {
                            if true {
                                if true {
                                    switch fi24 {
                                    case -179:
                                        fmt.Println(v15)
                                        v40 := GenS3{F0: 18446744073709551615}
                                        v11 -= v11
                                        _ = v40
                                    case 0:
                                        var v41 int64 = 0
                                        _ = v41
                                    }
                                    v3 = -2.25
                                } else {
                                    v12 = v12
                                    var v42 bool = true
                                    _ = v42
                                }
                            }
                        }
                    case 0:
                        v43 := GenS3{F0: 42}
                        _ = v43
                    default:
                        v44 := GenS4{F0: (v13 + ""), F1: false, F2: 634}
                        _ = v44
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v45 := GenS3{F0: 42}
                    _ = v45
                }
                _ = fi24
            }
            _ = fi19
        }
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        var v47 float64 = (v3 + 1e10)
        v4 = v4
        var v48 int16 = v12
        _ = fi46
        _ = v47
        _ = v48
    }
    _ = len(v5)
    if false {
        var v49 uint16 = 883
        _ = v49
    } else {
        switch "\n\t" {
        case "hello":
            var v50 uint32 = v4
            _ = v50
        case "":
            v51 := GenS4{F0: v13, F1: false, F2: 2}
            _ = v51
        default:
            for fi52 := 0; fi52 < 7; fi52++ {
                var v53 float32 = v11
                var v54 uint32 = (^(v4))
                _ = fi52
                _ = v53
                _ = v54
            }
        }
        v55 := GenS5{F0: (v15 & (^(v15))), F1: (v15 | v15), F2: 300}
        _ = v55
    }
    if (1.5 < 0.0) {
        for fi56 := 0; fi56 < 15; fi56++ {
            fmt.Println(v3, v12, v1)
            var v57 uint32 = (v4 - v4)
            _ = fi56
            _ = v57
        }
        if (v13 != v13) {
            var v58 int32 = 42
            v59 := GenS4{F0: (v13 + v13), F1: false, F2: 42}
            _ = v58
            _ = v59
        } else {
            var v60 uint = 679
            for fi61 := 0; fi61 < 13; fi61++ {
                if (3 != -1) {
                    if true {
                        for fi62 := 0; fi62 < 16; fi62++ {
                            v63 := GenS3{F0: 313}
                            _ = fi62
                            _ = v63
                        }
                    } else {
                        for fi64 := 0; fi64 < 3; fi64++ {
                            _ = len(v5)
                            switch (fi61 / ((fi61) | 1)) {
                            case 128:
                                v4 = (v4 / (((v4 + v4)) | 1))
                            case 32767:
                                if false {
                                    v13 = v13
                                    switch fi61 {
                                    case -1:
                                        switch (-(fi61)) {
                                        case 32767:
                                            switch (fi64 * fi61) {
                                            case -9223372036854775808:
                                                if ((true && false) && (!false)) {
                                                    if (0 >= 541) {
                                                        v65 := GenS3{F0: 327}
                                                        v4 += v4
                                                        v60 += (v60 | (^(v60)))
                                                        _ = v65
                                                    }
                                                    v3 -= (3.14159 - (0.1 * (v3 - float64(v11))))
                                                    var v66 uint32 = (v4 << (((^(v60))) & 63))
                                                    v12 = (v12 | v12)
                                                    _ = v66
                                                } else {
                                                    if (((42 <= 42) || (true && false)) && false) {
                                                        var v67 uint8 = 53
                                                        _ = v67
                                                    } else {
                                                        var v68 uint = 18446744073709551615
                                                        if false {
                                                            v1 -= v1
                                                            if false {
                                                                if (true && false) {
                                                                    if true {
                                                                        v4 = (^(v4))
                                                                    } else {
                                                                        v60 = (v60 | v68)
                                                                        var v69 int8 = 127
                                                                        _ = v69
                                                                    }
                                                                    if false {
                                                                        fmt.Println(v11, fi64, v15)
                                                                    } else {
                                                                        if false {
                                                                            v12 -= 0
                                                                            if (false && true) {
                                                                                v1 = 1e10
                                                                            }
                                                                            var v70 string = v13
                                                                            var v71 int = (fi64 / (((fi64 * fi61)) | 1))
                                                                            switch (v71 + (^(v71))) {
                                                                            case -516:
                                                                                if true {
                                                                                    var v72 int32 = 42
                                                                                    v12 -= (v12 << ((v68) & 63))
                                                                                    fmt.Println(v68, v68, fi64)
                                                                                    _ = v72
                                                                                }
                                                                            }
                                                                            _ = v70
                                                                            _ = v71
                                                                        } else {
                                                                            var v73 string = "a"
                                                                            _ = v73
                                                                        }
                                                                    }
                                                                    _ = len(v5)
                                                                    v68 -= 609
                                                                    v68 = uint(v15)
                                                                }
                                                                var v74 int64 = -869
                                                                v4 = v4
                                                                _ = v74
                                                            } else {
                                                                switch -1 {
                                                                case -240:
                                                                    v5[("\n\t" + v13)] = fi61
                                                                }
                                                                v75 := GenS6{F0: v13}
                                                                var v76 int16 = v12
                                                                v77 := GenS5{F0: 96, F1: 91, F2: 128}
                                                                v78 := GenS4{F0: ("a" + "\n\t"), F1: true, F2: 4294967295}
                                                                _ = v75
                                                                _ = v76
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                            switch (((v13 + v13) + v13) + ("foo bar" + v13)) {
                                                            case "":
                                                                v79 := GenS4{F0: (v13 + ((v13 + ("foo bar" + ((v13 + v13) + (v13 + v13)))) + ((v13 + (v13 + ((v13 + v13) + (v13 + "\n\t")))) + v13))), F1: (42 > 0), F2: uint64(v15)}
                                                                _ = v79
                                                            case "foo bar":
                                                                var v80 uint64 = 42
                                                                _ = v80
                                                            case "a":
                                                                var v81 int8 = 42
                                                                _ = v81
                                                            }
                                                        }
                                                        _ = v68
                                                    }
                                                    fmt.Println(fi64)
                                                }
                                            case 403:
                                                v82 := GenS7{F0: 170}
                                                _ = v82
                                            case -781:
                                                switch int(v12) {
                                                case 4294967295:
                                                    var v83 int = (fi61 >> (((v60 | v60)) & 63))
                                                    _ = v83
                                                case 328:
                                                    var v84 bool = true
                                                    _ = v84
                                                }
                                            default:
                                                var v85 int = fi64
                                                v13 = v13
                                                v86 := v5[((v13 + v13) + v13)]
                                                _ = v85
                                                _ = v86
                                            }
                                        }
                                    case 2147483647:
                                        switch ((v13 + ("hello" + "a")) + (v13 + v13)) {
                                        case "a":
                                            v12 -= v12
                                        case "hello":
                                            v87 := GenS8{F0: 2, F1: 99}
                                            _ = v87
                                        case "":
                                            if false {
                                                var v88 int = (fi64 & fi61)
                                                if len(v2) > 0 {
                                                    v2[0] = 2147483647
                                                }
                                                _ = v88
                                            } else {
                                                if len(v2) > 0 {
                                                    v2[0] = fi64
                                                }
                                                var v89 float64 = v3
                                                _ = v89
                                            }
                                        }
                                    case 4294967295:
                                        delete(v5, v13)
                                    default:
                                        v3 -= v3
                                    }
                                    v90 := GenS6{F0: (v13 + v13)}
                                    _ = v90
                                }
                            case 184:
                                switch fi61 {
                                case 970:
                                    v5[("foo bar" + "a")] = 2147483647
                                case 4294967295:
                                    v5["a"] = 128
                                case -9223372036854775808:
                                    if true {
                                        v11 = 1e10
                                        var v91 bool = (true && ((fi64 / ((4294967295) | 1)) < (fi61 >> ((245) & 63))))
                                        switch ("\n\t" + "a") {
                                        case "foo bar":
                                            fmt.Println(fi61, v3)
                                        case "":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        _ = v91
                                    }
                                default:
                                    v15 += v15
                                }
                            }
                            _ = len(v2)
                            v15 = (v15 << (((v60 << ((306) & 63))) & 63))
                            _ = fi64
                        }
                        var v92 int64 = 1
                        v93 := v5[""]
                        _ = v92
                        _ = v93
                    }
                    if (1 != 9223372036854775807) {
                        if false {
                            fmt.Println(v1, v15)
                            var v94 uint64 = 0
                            _ = v94
                        }
                    } else {
                        delete(v5, (v13 + (v13 + v13)))
                        if true {
                            if (-2.25 != v3) {
                                if (42 == 1) {
                                    if (v15 >= 1) {
                                        var v95 float32 = (v11 - 0.1)
                                        v96 := GenS4{F0: v13, F1: false, F2: 0}
                                        v95 = 3.14159
                                        _ = v95
                                        _ = v96
                                    }
                                    if false {
                                        switch v13 {
                                        case "foo bar":
                                            for fi97 := 0; fi97 < 5; fi97++ {
                                                if len(v2) > 0 {
                                                    v2[0] = fi61
                                                }
                                                _ = fi97
                                            }
                                        case "":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        default:
                                            v98 := GenS8{F0: 255, F1: 869}
                                            _ = v98
                                        }
                                    } else {
                                        var v99 string = "\n\t"
                                        switch "\n\t" {
                                        case "hello":
                                            var v100 int8 = 0
                                            _ = v100
                                        default:
                                            v60 = 439
                                        }
                                        delete(v5, "\n\t")
                                        _ = v99
                                    }
                                } else {
                                    v13 = ((v13 + ((v13 + v13) + v13)) + ((v13 + v13) + v13))
                                    var v101 bool = (!((((!false) || false) && true) && false))
                                    fmt.Println(v101)
                                    _ = v101
                                }
                            } else {
                                switch ("a" + v13) {
                                case "a":
                                    switch fi61 {
                                    case -863:
                                        if true {
                                            v102 := GenS9{F0: (fi61 >> (((^(v60))) & 63)), F1: fi61, F2: false}
                                            switch v13 {
                                            case "\n\t":
                                                var v103 int32 = 934
                                                _ = v103
                                            }
                                            _ = v102
                                        } else {
                                            v1 = v11
                                            var v104 uint16 = 42
                                            _ = v104
                                        }
                                    }
                                case "":
                                    var v105 float64 = v3
                                    _ = v105
                                case "hello":
                                    var v106 uint64 = 2
                                    _ = v106
                                }
                                if true {
                                    v107 := GenS5{F0: (v15 & (v15 * 128)), F1: v15, F2: (fi61 + (^(fi61)))}
                                    var v108 int = -987
                                    if ((v4 % (((^(v4))) | 1)) >= uint32(v15)) {
                                        switch (v13 + "hello") {
                                        case "":
                                            v13 = v13
                                        default:
                                            for fi109 := 0; fi109 < 8; fi109++ {
                                                v60 = (v60 << ((v60) & 63))
                                                fmt.Println(v60)
                                                fmt.Println(v15, v15, v11)
                                                fmt.Println(v1, v108)
                                                if (!false) {
                                                    v13 = v13
                                                }
                                                _ = fi109
                                            }
                                        }
                                        for fi110 := 0; fi110 < 8; fi110++ {
                                            if false {
                                                switch "hello" {
                                                case "":
                                                    v60 += (v60 / ((v60) | 1))
                                                case "\n\t":
                                                    fmt.Println(v1, v4, v13)
                                                default:
                                                    v5["\n\t"] = (^(fi110))
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            } else {
                                                var v111 uint16 = 713
                                                _ = v111
                                            }
                                            _ = fi110
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    _ = v107
                                    _ = v108
                                }
                                var v112 uint32 = (v4 / ((253) | 1))
                                _ = v112
                            }
                            var v113 string = "a"
                            v2 = append(v2, int(v4))
                            _ = v113
                        } else {
                            v114 := v5[((v13 + (v13 + "\n\t")) + v13)]
                            var v115 uint = (v60 & v60)
                            var v116 uint8 = 1
                            _ = v114
                            _ = v115
                            _ = v116
                        }
                        v2 = append(v2, fi61)
                    }
                    var v117 int8 = 100
                    _ = v117
                } else {
                    switch v13 {
                    case "foo bar":
                        v12 = -1
                    case "\n\t":
                        if (true && ((v60 < (v60 / (((v60 - 0)) | 1))) && ((false && true) || ((1 != 541) && ((^(v60)) < v60))))) {
                            switch ("\n\t" + v13) {
                            case "hello":
                                if true {
                                    for fi118 := 0; fi118 < 14; fi118++ {
                                        if false {
                                            var v119 int64 = 1
                                            var v120 string = v13
                                            _ = len(v5)
                                            _ = v119
                                            _ = v120
                                        }
                                        _ = fi118
                                    }
                                    var v121 uint = v60
                                    fmt.Println(v3, fi61)
                                    _ = v121
                                } else {
                                    if false {
                                        v5["a"] = -142
                                        for fi122 := 0; fi122 < 7; fi122++ {
                                            var v123 uint8 = v15
                                            v5["foo bar"] = int(v12)
                                            if (float64((float32((v3 - v3)) + v1)) >= (v3 * v3)) {
                                                v124 := GenS3{F0: 231}
                                                switch (fi122 >> ((v60) & 63)) {
                                                case -697:
                                                    v3 = 0.0
                                                case 9223372036854775807:
                                                    switch fi122 {
                                                    case 42:
                                                        var v125 uint8 = (v123 + v15)
                                                        _ = v125
                                                    }
                                                case 128:
                                                    var v126 uint = v60
                                                    _ = v126
                                                }
                                                var v127 string = ""
                                                _ = v124
                                                _ = v127
                                            }
                                            _ = fi122
                                            _ = v123
                                        }
                                    } else {
                                        var v128 int16 = (v12 ^ v12)
                                        _ = len(v5)
                                        _ = len(v5)
                                        _ = v128
                                    }
                                }
                            case "\n\t":
                                switch 2147483647 {
                                case -756:
                                    v11 = 0.0
                                case -483:
                                    v5["a"] = (fi61 & fi61)
                                }
                            case "":
                                v1 += v11
                            }
                            var v129 int64 = -9223372036854775808
                            _ = v129
                        } else {
                            v60 -= v60
                        }
                    default:
                        var v130 string = v13
                        _ = v130
                    }
                    var v131 uint = 0
                    v132 := GenS10{F0: 241, F1: fi61, F2: 625}
                    _ = v131
                    _ = v132
                }
                _ = fi61
            }
            delete(v5, ("foo bar" + "foo bar"))
            _ = v60
        }
    }
    for fi133 := 0; fi133 < 5; fi133++ {
        var v134 uint = 255
        if len(v2) > 0 {
            v2[0] = fi133
        }
        for fi135 := 0; fi135 < 5; fi135++ {
            var v136 uint64 = 4294967295
            v137 := v5[v13]
            _ = fi135
            _ = v136
            _ = v137
        }
        _ = fi133
        _ = v134
    }
    switch ("foo bar" + "hello") {
    case "foo bar":
        switch ("\n\t" + v13) {
        case "":
            switch "foo bar" {
            case "foo bar":
                if true {
                    var v138 int8 = 0
                    delete(v5, v13)
                    for fi139 := 0; fi139 < 16; fi139++ {
                        var v140 float64 = 0.0
                        var v141 int32 = 222
                        fmt.Println(v3, v3)
                        _ = fi139
                        _ = v140
                        _ = v141
                    }
                    _ = v138
                } else {
                    var v142 uint32 = v4
                    if false {
                        var v143 uint32 = (v142 | 1)
                        v11 = v11
                        var v144 bool = (0 > 9223372036854775807)
                        _ = v143
                        _ = v144
                    }
                    for fi145 := 0; fi145 < 8; fi145++ {
                        v146 := GenS11{F0: 2}
                        for fi147 := 0; fi147 < 20; fi147++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v148 float32 = 1e10
                            v149 := v5[v13]
                            _ = fi147
                            _ = v148
                            _ = v149
                        }
                        _ = fi145
                        _ = v146
                    }
                    _ = v142
                }
            }
        case "foo bar":
            v15 = (^(v15))
        case "hello":
            for fi150 := 0; fi150 < 13; fi150++ {
                v12 += (v12 & 467)
                for fi151 := 0; fi151 < 4; fi151++ {
                    switch (v13 + "") {
                    case "foo bar":
                        if true {
                            v1 -= float32(v3)
                            v1 = v11
                        } else {
                            var v152 float64 = 1.5
                            v11 = (((v1 - (v11 * (v11 + float32((float64((v1 - float32(1.5))) + float64(v1)))))) * 0.1) + 3.14159)
                            _ = v152
                        }
                        fmt.Println(v4, fi150)
                    case "\n\t":
                        if true {
                            if false {
                                if (65535 > 2) {
                                    fmt.Println(v13, v3, fi150)
                                    v153 := GenS12{F0: 65535}
                                    v154 := GenS13{F0: 37}
                                    _ = v153
                                    _ = v154
                                } else {
                                    var v155 int64 = 42
                                    _ = v155
                                }
                                v1 -= -2.25
                            } else {
                                v156 := GenS13{F0: 746}
                                _ = v156
                            }
                            var v157 int = fi150
                            _ = v157
                        }
                    default:
                        var v158 uint16 = 758
                        _ = v158
                    }
                    _ = fi151
                }
                v5[v13] = fi150
                _ = fi150
            }
        }
    default:
        fmt.Println(v15, v4)
    }
    for fi159 := 0; fi159 < 11; fi159++ {
        for fi160 := 0; fi160 < 6; fi160++ {
            v161 := GenS6{F0: "\n\t"}
            _ = fi160
            _ = v161
        }
        _ = fi159
    }
    for fi162 := 0; fi162 < 19; fi162++ {
        if (358 == -2147483648) {
            v4 += (^(v4))
            fmt.Println(v13, v1, v1)
            fmt.Println(v11, v15, v15)
        }
        _ = fi162
    }
    v163 := GenS14{F0: -9223372036854775808, F1: v4, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v163
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
