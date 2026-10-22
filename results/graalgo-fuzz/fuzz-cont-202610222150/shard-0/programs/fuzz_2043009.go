package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 bool }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 bool; F2 int16 }
type GenS7 struct { F0 int64; F1 uint16; F2 bool }
type GenS8 struct { F0 uint8; F1 int32; F2 uint }
type GenS9 struct { F0 uint64; F1 uint64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 bool; F1 uint32; F2 string }
type GenS12 struct { F0 uint64; F1 bool; F2 int8 }

func main() {
    v1 := []uint32{410, 2147483647, 469, 0}
    v2 := make(map[int]bool)
    var v3 uint16 = 32767
    var v4 int32 = 42
    var v5 int8 = 127
    if false {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 int16 = 1
            var v8 float32 = (0.0 * 1.5)
            _ = fi6
            _ = v7
            _ = v8
        }
        v4 = 2147483647
        var v9 bool = (v3 <= v3)
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 int32 = (-(v4))
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int = fi12
        _ = fi12
        _ = v13
    }
    _ = len(v1)
    switch ("" + "\n\t") {
    case "a":
        fmt.Println(v5, v4, v4)
    case "hello":
        v5 += (v5 - 38)
    case "":
        var v14 int8 = (-(v5))
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 uint64 = 18446744073709551615
        v17 := GenS1{F0: v5, F1: fi15, F2: false}
        var v18 int8 = (v5 / ((-1) | 1))
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    fmt.Println(v5, v5)
    var v19 uint16 = (v3 << ((309) & 63))
    var v20 bool = false
    _ = len(v1)
    delete(v2, 42)
    if false {
        var v21 uint32 = 2147483647
        v1 = append(v1, v21)
        v20 = (!v20)
        _ = v21
    } else {
        switch ("a" + "a") {
        case "foo bar":
            var v22 uint64 = 1
            _ = v22
        }
        if (v20 || v20) {
            if v20 {
                v19 += v3
                fmt.Println(v20, v3)
            }
            switch "\n\t" {
            case "\n\t":
                for fi23 := 0; fi23 < 16; fi23++ {
                    if v20 {
                        var v24 string = "\n\t"
                        switch v24 {
                        case "hello":
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                        }
                        var v25 int32 = v4
                        _ = v24
                        _ = v25
                    } else {
                        var v26 int16 = 0
                        for fi27 := 0; fi27 < 2; fi27++ {
                            fmt.Println(v5, v20, v4)
                            if v20 {
                                if v20 {
                                    var v28 float64 = float64((3.14159 * (0.0 + float32(-2.25))))
                                    _ = v28
                                } else {
                                    _ = len(v1)
                                    var v29 int16 = 1000
                                    _ = v29
                                }
                                var v30 int32 = (v4 << ((2) & 63))
                                _ = v30
                            } else {
                                v19 -= (^(v3))
                                var v31 int32 = v4
                                _ = v31
                            }
                            v3 += (v19 / ((v3) | 1))
                            switch "" {
                            case "foo bar":
                                v32 := GenS2{F0: 166, F1: 2}
                                _ = v32
                            case "hello":
                                _ = len(v1)
                            case "\n\t":
                                v33 := GenS2{F0: 2, F1: 42}
                                _ = v33
                            }
                            _ = fi27
                        }
                        _ = v26
                    }
                    fmt.Println(v20, v19)
                    var v34 uint64 = 18446744073709551615
                    _ = fi23
                    _ = v34
                }
            case "":
                if (false || v20) {
                    if (1 >= 42) {
                        fmt.Println(v20, v19, v5)
                        if true {
                            var v35 uint32 = 97
                            for fi36 := 0; fi36 < 6; fi36++ {
                                if ((v35 / ((v35) | 1)) != v35) {
                                    switch "a" {
                                    case "foo bar":
                                        if (v20 && (v5 <= 1)) {
                                            if ((^(v19)) >= (v3 + v3)) {
                                                var v37 uint16 = (v3 * v3)
                                                if (722 < 1) {
                                                    var v38 bool = (("" == "hello") || (42 >= 1))
                                                    _ = v38
                                                } else {
                                                    var v39 int8 = (^(v5))
                                                    for fi40 := 0; fi40 < 7; fi40++ {
                                                        fmt.Println(fi36, v5, v3)
                                                        if (((^(v35)) > 897) || (-2.25 < (1e10 - 1.5))) {
                                                            if v20 {
                                                                var v41 uint16 = v37
                                                                v37 -= v19
                                                                v3 = (v3 / ((v19) | 1))
                                                                _ = v41
                                                            } else {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                v39 = -128
                                                                fmt.Println(v3, fi40, v35)
                                                            }
                                                        } else {
                                                            switch ("\n\t" + "foo bar") {
                                                            case "":
                                                                v20 = v20
                                                            case "\n\t":
                                                                var v42 int = -793
                                                                _ = v42
                                                            case "foo bar":
                                                                _ = len(v2)
                                                            }
                                                            v3 = v37
                                                        }
                                                        switch (fi40 | 145) {
                                                        case 77:
                                                            v3 = 65535
                                                        case -230:
                                                            v5 = 60
                                                        default:
                                                            v43 := GenS3{F0: (fi40 + fi40)}
                                                            _ = v43
                                                        }
                                                        v37 = uint16(v35)
                                                        switch (("hello" + "") + ("" + "a")) {
                                                        case "a":
                                                            v1 = append(v1, 2)
                                                        case "":
                                                            if (3.14159 <= 1e10) {
                                                                if len(v1) > 0 {
                                                                    v1[0] = (v35 >> ((959) & 63))
                                                                }
                                                                if (((0.0 * 3.14159) / (1.5 * 3.14159)) == 3.14159) {
                                                                    var v44 uint16 = (v3 | (v19 >> ((856) & 63)))
                                                                    v1 = append(v1, v35)
                                                                    if false {
                                                                        switch fi40 {
                                                                        case 128:
                                                                            fmt.Println(v5, v19, v39)
                                                                        case 700:
                                                                            var v45 uint = 1
                                                                            _ = v45
                                                                        default:
                                                                            if true {
                                                                                var v46 string = "\n\t"
                                                                                if (-2147483648 >= fi36) {
                                                                                    var v47 int = (fi40 * fi40)
                                                                                    var v48 float64 = (0.0 - ((3.14159 - 3.14159) * 0.1))
                                                                                    v39 = 1
                                                                                    _ = v47
                                                                                    _ = v48
                                                                                } else {
                                                                                    v49 := GenS2{F0: 1, F1: 42}
                                                                                    v19 = v37
                                                                                    if v20 {
                                                                                        v20 = (("hello" > (v46 + v46)) || ((v46 + ("foo bar" + v46)) <= (v46 + ("\n\t" + v46))))
                                                                                        v50 := GenS2{F0: 643, F1: 770}
                                                                                        _ = v50
                                                                                    }
                                                                                    _ = v49
                                                                                }
                                                                                switch (fi40 * (fi36 + fi40)) {
                                                                                case -1:
                                                                                    switch (v46 + (v46 + ((((((v46 + v46) + (v46 + v46)) + "hello") + v46) + (v46 + ("" + "foo bar"))) + "hello"))) {
                                                                                    case "a":
                                                                                        v20 = (42 >= (fi36 << ((0) & 63)))
                                                                                    case "\n\t":
                                                                                        var v51 string = (v46 + "foo bar")
                                                                                        _ = v51
                                                                                    case "":
                                                                                        switch 9223372036854775807 {
                                                                                        case 4294967295:
                                                                                            var v52 float64 = 1e10
                                                                                            _ = v52
                                                                                        case 1:
                                                                                            fmt.Println(v44, fi36, v3)
                                                                                        case 127:
                                                                                            v3 = v19
                                                                                        default:
                                                                                            var v53 uint16 = (v3 ^ (v19 * v37))
                                                                                            _ = v53
                                                                                        }
                                                                                    }
                                                                                    var v54 float32 = -2.25
                                                                                    fmt.Println(v54)
                                                                                    _ = v54
                                                                                case 2147483647:
                                                                                    v55 := GenS2{F0: 1, F1: 209}
                                                                                    _ = v55
                                                                                }
                                                                                _ = v46
                                                                            } else {
                                                                                _ = len(v2)
                                                                            }
                                                                        }
                                                                    }
                                                                    if (!(v5 != (^(v39)))) {
                                                                        if (!(129 < 0)) {
                                                                            var v56 uint32 = 545
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                if v20 {
                                                                                    v35 = 2147483648
                                                                                    v57 := GenS2{F0: 42, F1: 65535}
                                                                                    _ = v57
                                                                                }
                                                                            case "a":
                                                                                v39 = 42
                                                                            case "":
                                                                                if true {
                                                                                    v39 = -128
                                                                                    v3 = 523
                                                                                } else {
                                                                                    fmt.Println(v4, v20, v44)
                                                                                    v39 = -128
                                                                                }
                                                                            default:
                                                                                var v58 int64 = 9223372036854775807
                                                                                _ = v58
                                                                            }
                                                                            switch fi40 {
                                                                            case 9223372036854775807:
                                                                                var v59 float64 = (3.14159 * -2.25)
                                                                                _ = v59
                                                                            case 32768:
                                                                                v60 := GenS4{F0: fi36, F1: fi40}
                                                                                _ = v60
                                                                            default:
                                                                                var v61 uint = 0
                                                                                _ = v61
                                                                            }
                                                                            _ = v56
                                                                        }
                                                                        _ = len(v1)
                                                                        v35 += v35
                                                                    }
                                                                    _ = v44
                                                                }
                                                            }
                                                        default:
                                                            if v20 {
                                                                if v20 {
                                                                    if v20 {
                                                                        v62 := GenS4{F0: 4294967295, F1: (fi36 >> ((65535) & 63))}
                                                                        delete(v2, 328)
                                                                        v19 += 42
                                                                        _ = v62
                                                                    }
                                                                } else {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = v35
                                                                    }
                                                                    v4 += (v4 * (v4 ^ int32(v39)))
                                                                }
                                                                if (!(((v39 / ((-1) | 1)) > v5) || (v4 >= (v4 | 0)))) {
                                                                    fmt.Println(v4, v35)
                                                                    v1 = append(v1, 1)
                                                                } else {
                                                                    if (!(!v20)) {
                                                                        switch 1 {
                                                                        case 483:
                                                                            switch "hello" {
                                                                            case "":
                                                                                fmt.Println(v35, fi36, v39)
                                                                            case "a":
                                                                                v39 += v5
                                                                                v35 = 857
                                                                                switch "" {
                                                                                case "hello":
                                                                                    if false {
                                                                                        v5 = v5
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        switch (^(fi36)) {
                                                                                        case 400:
                                                                                            var v63 int64 = -1
                                                                                            _ = v63
                                                                                        case -712:
                                                                                            _ = len(v1)
                                                                                        case -279:
                                                                                            v64 := GenS4{F0: (fi36 << ((2) & 63)), F1: (fi40 / (((fi36 + -329)) | 1))}
                                                                                            _ = v64
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v65 uint = 4294967295
                                                                                    _ = v65
                                                                                }
                                                                            case "foo bar":
                                                                                var v66 uint = 65535
                                                                                _ = v66
                                                                            }
                                                                        case 9223372036854775807:
                                                                            if (!true) {
                                                                                if v20 {
                                                                                    v67 := GenS1{F0: (v5 + -91), F1: (fi40 + -934), F2: (v20 && (v37 >= v37))}
                                                                                    if v20 {
                                                                                        var v68 float32 = 0.0
                                                                                        _ = v68
                                                                                    }
                                                                                    var v69 uint16 = 255
                                                                                    fmt.Println(v20, v19, v19)
                                                                                    _ = v67
                                                                                    _ = v69
                                                                                } else {
                                                                                    fmt.Println(v37)
                                                                                    v35 = (^(v35))
                                                                                }
                                                                                var v70 uint64 = 489
                                                                                _ = v70
                                                                            } else {
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    var v71 uint32 = (v35 ^ v35)
                                                                                    _ = v71
                                                                                case "a":
                                                                                    v72 := GenS1{F0: 127, F1: -2147483648, F2: v20}
                                                                                    _ = v72
                                                                                case "foo bar":
                                                                                    v73 := GenS3{F0: (^(fi40))}
                                                                                    _ = v73
                                                                                }
                                                                                v5 = 44
                                                                            }
                                                                        case 2147483647:
                                                                            v74 := GenS5{F0: -32768}
                                                                            _ = v74
                                                                        }
                                                                        fmt.Println(v37, v20)
                                                                        v19 -= (v3 * (v3 << ((0) & 63)))
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = fi40
                                                    }
                                                    for fi75 := 0; fi75 < 9; fi75++ {
                                                        delete(v2, fi36)
                                                        v76 := GenS4{F0: fi75, F1: 42}
                                                        v77 := GenS2{F0: 1, F1: 385}
                                                        var v78 float64 = -2.25
                                                        _ = fi75
                                                        _ = v76
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                    _ = v39
                                                }
                                                for fi79 := 0; fi79 < 8; fi79++ {
                                                    fmt.Println(v37)
                                                    _ = fi79
                                                }
                                                _ = v37
                                            }
                                            v80 := GenS3{F0: (-(fi36))}
                                            v81 := GenS3{F0: (-(fi36))}
                                            _ = v80
                                            _ = v81
                                        } else {
                                            v82 := GenS6{F0: -9223372036854775808, F1: (42 < 2147483647), F2: int16(v19)}
                                            fmt.Println(v20, v5, v3)
                                            _ = v82
                                        }
                                    }
                                    if v20 {
                                        for fi83 := 0; fi83 < 19; fi83++ {
                                            if (!false) {
                                                v19 = (v19 | (v3 / ((v19) | 1)))
                                                var v84 uint = 42
                                                switch ("hello" + "a") {
                                                case "":
                                                    var v85 int = (fi36 << (((v84 + (v84 << (((v84 << ((v84) & 63))) & 63)))) & 63))
                                                    _ = v85
                                                case "hello":
                                                    v86 := GenS7{F0: -1, F1: (v3 % (((^(v3))) | 1)), F2: false}
                                                    _ = v86
                                                case "a":
                                                    fmt.Println(v84)
                                                default:
                                                    fmt.Println(v3, v4)
                                                }
                                                var v87 uint16 = (v3 - (^(v3)))
                                                v5 += (-(v5))
                                                var v88 string = "\n\t"
                                                _ = v84
                                                _ = v87
                                                _ = v88
                                            }
                                            v5 += (-(v5))
                                            _ = fi83
                                        }
                                    }
                                } else {
                                    v89 := GenS8{F0: 127, F1: (^(v4)), F2: 1}
                                    for fi90 := 0; fi90 < 16; fi90++ {
                                        fmt.Println(v35, v4)
                                        _ = fi90
                                    }
                                    v20 = (v20 || (18446744073709551615 == 1))
                                    _ = v89
                                }
                                var v91 int16 = 0
                                _ = fi36
                                _ = v91
                            }
                            _ = v35
                        } else {
                            switch "foo bar" {
                            case "a":
                                var v92 int64 = 527
                                _ = v92
                            case "":
                                for fi93 := 0; fi93 < 6; fi93++ {
                                    if v20 {
                                        if (2 > 4294967295) {
                                            for fi94 := 0; fi94 < 16; fi94++ {
                                                switch (^(fi94)) {
                                                case 569:
                                                    fmt.Println(v4)
                                                case 0:
                                                    if true {
                                                        var v95 int8 = v5
                                                        if true {
                                                            var v96 uint8 = 42
                                                            if true {
                                                                v96 = (^(v96))
                                                            } else {
                                                                v5 = 24
                                                                if (("hello" >= "hello") || v20) {
                                                                    if true {
                                                                        var v97 int8 = (v95 + v95)
                                                                        _ = v97
                                                                    }
                                                                } else {
                                                                    var v98 float32 = float32(1e10)
                                                                    _ = v98
                                                                }
                                                                switch "" {
                                                                case "":
                                                                    if (("" + ("\n\t" + "")) > "hello") {
                                                                        v99 := GenS6{F0: 9223372036854775807, F1: (1 < 2), F2: 42}
                                                                        _ = v99
                                                                    } else {
                                                                        if ((519 <= 42) && ((3.14159 < (-2.25 / (0.0 + 3.14159))) && (-9223372036854775808 <= 42))) {
                                                                            v100 := GenS2{F0: 1, F1: 1}
                                                                            _ = v100
                                                                        } else {
                                                                            var v101 uint8 = (^(v96))
                                                                            if ("hello" <= "foo bar") {
                                                                                var v102 float64 = -2.25
                                                                                var v103 uint64 = 891
                                                                                switch fi93 {
                                                                                case -1:
                                                                                    v3 += (v19 & (^(v19)))
                                                                                }
                                                                                if v20 {
                                                                                    fmt.Println(v20)
                                                                                    switch (fi93 % (((^(fi94))) | 1)) {
                                                                                    case 39:
                                                                                        var v104 uint8 = 127
                                                                                        _ = v104
                                                                                    case 2147483647:
                                                                                        v105 := v2[fi94]
                                                                                        _ = v105
                                                                                    }
                                                                                    switch fi93 {
                                                                                    case -23:
                                                                                        v2[(fi94 % ((fi93) | 1))] = (3.14159 > ((-2.25 / (3.14159 - 1.5)) + 0.1))
                                                                                    case -2147483648:
                                                                                        var v106 int64 = -731
                                                                                        _ = v106
                                                                                    case 112:
                                                                                        var v107 uint16 = v19
                                                                                        _ = v107
                                                                                    default:
                                                                                        v108 := GenS9{F0: (^(v103)), F1: v103}
                                                                                        var v109 bool = (int64(fi93) < -9223372036854775808)
                                                                                        _ = v108
                                                                                        _ = v109
                                                                                    }
                                                                                    v101 += (v101 + v101)
                                                                                }
                                                                                v4 = v4
                                                                                v110 := v2[fi93]
                                                                                v4 -= v4
                                                                                _ = v102
                                                                                _ = v103
                                                                                _ = v110
                                                                            }
                                                                            var v111 int64 = 1
                                                                            _ = v101
                                                                            _ = v111
                                                                        }
                                                                        _ = len(v2)
                                                                    }
                                                                }
                                                            }
                                                            var v112 uint16 = (v3 - 2)
                                                            _ = v112
                                                            _ = v96
                                                        }
                                                        _ = v95
                                                    } else {
                                                        var v113 bool = (v20 || ((fi94 % ((fi93) | 1)) < 155))
                                                        _ = v113
                                                    }
                                                default:
                                                    v3 = (v19 & 511)
                                                }
                                                v114 := GenS4{F0: 986, F1: fi93}
                                                v2[2147483647] = false
                                                _ = fi94
                                                _ = v114
                                            }
                                            switch "hello" {
                                            case "":
                                                v115 := GenS4{F0: 9223372036854775807, F1: (fi93 >> ((0) & 63))}
                                                _ = v115
                                            case "a":
                                                v3 = v3
                                            case "foo bar":
                                                var v116 bool = v20
                                                _ = v116
                                            }
                                            if false {
                                                var v117 uint16 = 369
                                                if true {
                                                    v19 = (v117 % ((32767) | 1))
                                                } else {
                                                    for fi118 := 0; fi118 < 8; fi118++ {
                                                        delete(v2, fi118)
                                                        _ = fi118
                                                    }
                                                    v119 := GenS10{F0: 18446744073709551615}
                                                    _ = v119
                                                }
                                                _ = v117
                                            }
                                        } else {
                                            for fi120 := 0; fi120 < 2; fi120++ {
                                                var v121 int16 = -1
                                                _ = len(v1)
                                                fmt.Println(v4, v19, v20)
                                                v20 = false
                                                v3 = 42
                                                v122 := v2[(fi120 | 824)]
                                                fmt.Println(v121)
                                                _ = fi120
                                                _ = v121
                                                _ = v122
                                            }
                                            v20 = (560 == -1)
                                            if len(v1) > 0 {
                                                v1[0] = 0
                                            }
                                        }
                                        v4 += v4
                                        fmt.Println(v3, v4, v4)
                                    }
                                    for fi123 := 0; fi123 < 4; fi123++ {
                                        v4 -= v4
                                        var v124 int64 = -9223372036854775808
                                        var v125 int8 = 42
                                        _ = fi123
                                        _ = v124
                                        _ = v125
                                    }
                                    var v126 int64 = 1
                                    _ = fi93
                                    _ = v126
                                }
                            default:
                                for fi127 := 0; fi127 < 7; fi127++ {
                                    fmt.Println(v3)
                                    switch "\n\t" {
                                    case "\n\t":
                                        v128 := GenS11{F0: ((v5 * v5) >= (-(v5))), F1: 65535, F2: ("foo bar" + (("hello" + "\n\t") + "hello"))}
                                        _ = v128
                                    case "hello":
                                        var v129 int8 = (v5 | 127)
                                        _ = v129
                                    }
                                    switch (^(fi127)) {
                                    case -422:
                                        v5 += 52
                                    default:
                                        if false {
                                            for fi130 := 0; fi130 < 12; fi130++ {
                                                fmt.Println(fi130, fi130, v5)
                                                fmt.Println(v3, v20)
                                                v5 -= v5
                                                _ = fi130
                                            }
                                            var v131 bool = ((v3 + (v3 + v19)) <= v19)
                                            var v132 uint64 = 2
                                            _ = v131
                                            _ = v132
                                        }
                                    }
                                    _ = fi127
                                }
                            }
                            v133 := v2[0]
                            _ = v133
                        }
                        for fi134 := 0; fi134 < 2; fi134++ {
                            fmt.Println(v4)
                            if (994 <= 2) {
                                v19 += (v19 / (((v19 * 1)) | 1))
                                fmt.Println(v19, fi134)
                            }
                            if false {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v1 = append(v1, 85)
                            }
                            _ = fi134
                        }
                    } else {
                        v135 := GenS12{F0: 1, F1: v20, F2: (^(v5))}
                        _ = len(v2)
                        delete(v2, 42)
                        _ = v135
                    }
                    var v136 int32 = v4
                    v5 += (-(v5))
                    _ = v136
                }
            }
            var v137 int16 = 1000
            _ = v137
        } else {
            v138 := v2[9223372036854775807]
            _ = v138
        }
    }
    var v139 uint64 = 42
    switch "a" {
    case "hello":
        v2[-9223372036854775808] = v20
    case "a":
        v139 = v139
    case "\n\t":
        var v140 float64 = (-2.25 - (((0.1 * float64(3.14159)) / 1.0) * -2.25))
        _ = v140
    default:
        var v141 uint8 = 127
        _ = v141
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v139:", v139)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v139
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
