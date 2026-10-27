package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 uint8; F2 int64 }
type GenS6 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 int16; F1 uint64 }
type GenS10 struct { F0 bool; F1 int16; F2 int64 }
type GenS11 struct { F0 uint16; F1 string; F2 int16 }
type GenS12 struct { F0 uint64; F1 bool; F2 int16 }
type GenS13 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS16 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 int32 = 42
    var v4 int32 = v3
    v5 := []int{}
    switch -2147483648 {
    case 42:
        if false {
            var v6 int16 = 1000
            if (true || true) {
                var v7 int8 = 42
                v4 = 42
                var v8 bool = true
                _ = v7
                _ = v8
            } else {
                if (-1 > 1) {
                    switch 32767 {
                    case 0:
                        v4 = 0
                    case 4294967295:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        var v9 int64 = 919
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 13; fi10++ {
                        fmt.Println(v6, fi10, v4)
                        _ = fi10
                    }
                    if (42 < 2) {
                        switch ("foo bar" + "") {
                        case "\n\t":
                            var v11 uint = 65535
                            _ = v11
                        case "hello":
                            _ = len(v1)
                        case "foo bar":
                            var v12 float32 = (0.0 * ((0.0 / ((1e10 * (((0.0 - (1e10 * float32(3.14159))) + ((float32(3.14159) - (3.14159 * 1e10)) + 3.14159)) * 1.5)) + 3.14159)) + (0.1 * float32(((-2.25 * 0.1) * 1e10)))))
                            _ = v12
                        }
                        v3 = v3
                    } else {
                        v13 := GenS1{F0: 0}
                        if ((true && (-1 >= 1)) && ("foo bar" > "a")) {
                            var v14 string = ""
                            switch v14 {
                            case "hello":
                                if false {
                                    var v15 float32 = ((0.0 + ((1.5 / (float32(1.5) + 0.1)) + (0.0 + 0.0))) - 3.14159)
                                    _ = v15
                                }
                            default:
                                switch (v14 + "a") {
                                case "\n\t":
                                    v16 := v2[(v14 + "hello")]
                                    _ = v16
                                case "a":
                                    v17 := v1[-314]
                                    _ = v17
                                }
                            }
                            _ = v14
                        }
                        _ = v13
                    }
                }
                var v18 int32 = 1
                _ = v18
            }
            var v19 uint32 = 240
            _ = v19
            _ = v6
        }
    case 4294967295:
        v3 = (v4 >> ((0) & 63))
    case 127:
        var v20 float64 = 1.5
        _ = v20
    }
    var v21 bool = ((!(65535 < 72)) && (-199 != 0))
    for fi22 := 0; fi22 < 8; fi22++ {
        switch ("\n\t" + (("a" + ("a" + "a")) + "hello")) {
        case "foo bar":
            v23 := GenS1{F0: 1}
            _ = v23
        case "\n\t":
            v21 = v21
        case "":
            for fi24 := 0; fi24 < 1; fi24++ {
                _ = len(v5)
                _ = fi24
            }
        default:
            fmt.Println(v21, fi22, v4)
        }
        _ = fi22
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := GenS1{F0: -9223372036854775808}
            var v28 uint = 42
            v29 := GenS2{F0: 42, F1: 1}
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi25
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        if v21 {
            var v31 int = fi30
            for fi32 := 0; fi32 < 10; fi32++ {
                var v33 float32 = 0.0
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = len(v2)
                _ = fi32
                _ = v33
            }
            _ = len(v5)
            _ = v31
        } else {
            var v34 uint32 = 4294967295
            var v35 float64 = 1.5
            _ = v34
            _ = v35
        }
        for fi36 := 0; fi36 < 14; fi36++ {
            var v37 float32 = 0.0
            fmt.Println(v37)
            v38 := GenS2{F0: 840, F1: 42}
            if len(v5) > 0 {
                v5[0] = fi30
            }
            v4 -= 42
            if true {
                var v39 int = (fi30 * -369)
                switch (fi36 << ((603) & 63)) {
                case -408:
                    fmt.Println(v4, v3)
                case 1:
                    fmt.Println(v4, v39, v3)
                case 4294967295:
                    fmt.Println(v37)
                default:
                    switch v39 {
                    case 32768:
                        if len(v5) > 0 {
                            v5[0] = (v39 / ((v39) | 1))
                        }
                    case 128:
                        fmt.Println(v3, fi36)
                    case -9223372036854775808:
                        v4 = v4
                    default:
                        switch ("foo bar" + "\n\t") {
                        case "hello":
                            if (42 != 2147483647) {
                                var v40 uint8 = 124
                                v39 = (fi30 / ((-456) | 1))
                                _ = v40
                            } else {
                                var v41 int64 = -9223372036854775808
                                if len(v5) > 0 {
                                    v5[0] = (v39 & fi30)
                                }
                                _ = v41
                            }
                        default:
                            v4 = v3
                        }
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if (32767 >= 0) {
                    v42 := GenS2{F0: 0, F1: 332}
                    var v43 int64 = -213
                    v44 := GenS2{F0: 42, F1: -32768}
                    _ = v42
                    _ = v43
                    _ = v44
                } else {
                    var v45 float32 = float32((0.0 * -2.25))
                    _ = v45
                }
                v46 := GenS1{F0: -1}
                _ = v39
                _ = v46
            }
            if (false || v21) {
                v21 = (2147483647 != 2147483648)
                if (false || (1 <= 436)) {
                    v4 -= (v3 + v4)
                    v37 -= ((1.5 - v37) - float32(3.14159))
                }
            }
            v47 := GenS3{F0: v21, F1: 849}
            _ = fi36
            _ = v37
            _ = v38
            _ = v47
        }
        _ = fi30
    }
    if (float64(1.5) >= 3.14159) {
        var v48 int64 = 590
        _ = v48
    }
    if v21 {
        if true {
            v49 := GenS3{F0: v21, F1: 758}
            _ = v49
        } else {
            for fi50 := 0; fi50 < 2; fi50++ {
                delete(v2, ("\n\t" + "hello"))
                _ = fi50
            }
            v4 = (v4 & v3)
        }
        var v51 float32 = 0.1
        v52 := GenS1{F0: 243}
        _ = v51
        _ = v52
    }
    var v53 uint16 = 42
    _ = len(v1)
    _ = len(v5)
    for fi54 := 0; fi54 < 20; fi54++ {
        var v55 int8 = -1
        v2[(("hello" + "hello") + ("a" + ""))] = true
        _ = len(v5)
        _ = fi54
        _ = v55
    }
    if v21 {
        if (false || v21) {
            var v56 string = "hello"
            _ = v56
        }
    } else {
        v57 := GenS4{F0: (!(!v21))}
        for fi58 := 0; fi58 < 9; fi58++ {
            for fi59 := 0; fi59 < 19; fi59++ {
                switch fi58 {
                case 127:
                    var v60 int8 = 100
                    _ = v60
                case -174:
                    switch (-(fi59)) {
                    case -2147483648:
                        v21 = ((903 != -32768) || v21)
                    }
                }
                if v21 {
                    _ = len(v5)
                    v61 := GenS4{F0: (!((!(v21 || v21)) || v21))}
                    switch "" {
                    case "":
                        switch "a" {
                        case "a":
                            if v21 {
                                var v62 bool = false
                                var v63 int32 = (v4 + (v4 + -1))
                                _ = v62
                                _ = v63
                            }
                        case "":
                            if (false && (v21 || ("a" > ("a" + "")))) {
                                v64 := GenS3{F0: v21, F1: 2147483647}
                                switch (fi59 * (fi59 >> ((447) & 63))) {
                                case 42:
                                    switch ("foo bar" + "hello") {
                                    case "\n\t":
                                        v53 = v53
                                    case "a":
                                        var v65 uint64 = 42
                                        _ = v65
                                    case "":
                                        v21 = ((3.14159 > ((((float32((-2.25 * (3.14159 / (0.1 * 3.14159)))) + (0.0 / float32((0.1 + float64((0.1 * 3.14159)))))) - 1e10) + 1.5) * 3.14159)) && false)
                                    default:
                                        var v66 uint8 = 255
                                        _ = v66
                                    }
                                    var v67 int8 = 42
                                    _ = v67
                                case -9223372036854775808:
                                    if len(v5) > 0 {
                                        v5[0] = fi59
                                    }
                                case 32767:
                                    var v68 int16 = -32768
                                    _ = v68
                                default:
                                    if true {
                                        var v69 uint8 = 0
                                        _ = v69
                                    } else {
                                        v70 := v2["foo bar"]
                                        switch "\n\t" {
                                        case "a":
                                            if (v21 && ((0.0 != -2.25) && (false && (!((-(v4)) > 1))))) {
                                                var v71 float64 = 1e10
                                                _ = v71
                                            }
                                            v70 = (!false)
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        v21 = (127 >= 0)
                                        _ = v70
                                    }
                                }
                                if ("foo bar" > "foo bar") {
                                    v72 := v1[(^(fi58))]
                                    if (-2.25 >= 1.5) {
                                        fmt.Println(v3, fi59)
                                        v4 = v4
                                    } else {
                                        if false {
                                            if (false && v21) {
                                                fmt.Println(v3)
                                                var v73 uint64 = 4294967295
                                                _ = v73
                                            }
                                            if (16 < -62) {
                                                delete(v1, 32767)
                                                v74 := GenS4{F0: v21}
                                                var v75 int16 = -423
                                                _ = v74
                                                _ = v75
                                            }
                                            switch fi58 {
                                            case -1:
                                                switch fi58 {
                                                case -701:
                                                    if (4294967295 <= 589) {
                                                        v72 = ("" + (v72 + v72))
                                                        v76 := GenS1{F0: -272}
                                                        _ = v76
                                                    } else {
                                                        if false {
                                                            fmt.Println(v4, v4, v72)
                                                            v4 = (v4 | (v4 * -1))
                                                        }
                                                        var v77 int64 = -460
                                                        _ = v77
                                                    }
                                                case 2147483647:
                                                    _ = len(v5)
                                                }
                                            }
                                            switch (("" + "hello") + "hello") {
                                            case "\n\t":
                                                var v78 int32 = v4
                                                _ = v78
                                            case "hello":
                                                if ((false || v21) || v21) {
                                                    if ((-2.25 - 0.1) <= 3.14159) {
                                                        v3 -= (^(v4))
                                                        if (v21 && (!false)) {
                                                            v72 = (("\n\t" + "\n\t") + v72)
                                                            var v79 uint = 577
                                                            _ = v79
                                                        } else {
                                                            var v80 int64 = 9223372036854775807
                                                            _ = v80
                                                        }
                                                    } else {
                                                        var v81 uint64 = 18446744073709551615
                                                        switch fi59 {
                                                        case -2147483648:
                                                            switch (fi58 >> ((65535) & 63)) {
                                                            case 128:
                                                                v82 := GenS2{F0: (v81 << ((4294967295) & 63)), F1: 1}
                                                                _ = v82
                                                            default:
                                                                v83 := GenS5{F0: 356, F1: 1, F2: 1}
                                                                _ = v83
                                                            }
                                                        case 127:
                                                            if false {
                                                                if len(v5) > 0 {
                                                                    v5[0] = fi58
                                                                }
                                                            }
                                                        case 709:
                                                            switch v72 {
                                                            case "\n\t":
                                                                var v84 float64 = -2.25
                                                                _ = v84
                                                            default:
                                                                v85 := GenS6{F0: 82, F1: int8(v53), F2: 65535}
                                                                v72 = v72
                                                                var v86 uint64 = (v81 % (((v81 ^ (v81 << ((18446744073709551615) & 63)))) | 1))
                                                                var v87 float32 = (0.1 + 1.5)
                                                                _ = len(v5)
                                                                switch v72 {
                                                                case "\n\t":
                                                                    v88 := GenS3{F0: ((v21 || (1 == 65535)) || true), F1: 65535}
                                                                    _ = v88
                                                                default:
                                                                    if (2 <= 2) {
                                                                        var v89 uint32 = 1
                                                                        var v90 float64 = ((1e10 + (1e10 * (3.14159 * -2.25))) - (1.5 / 1.0))
                                                                        switch v72 {
                                                                        case "a":
                                                                            switch -9223372036854775808 {
                                                                            case 128:
                                                                                delete(v1, (fi59 * 891))
                                                                                v1[(fi59 ^ 982)] = v72
                                                                                switch "hello" {
                                                                                case "hello":
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v90 = v90
                                                                                    case "hello":
                                                                                        var v91 int32 = v3
                                                                                        _ = v91
                                                                                    }
                                                                                case "a":
                                                                                    v92 := GenS5{F0: -296, F1: 1, F2: 1}
                                                                                    _ = v92
                                                                                default:
                                                                                    var v93 uint8 = 255
                                                                                    _ = v93
                                                                                }
                                                                            case -148:
                                                                                if false {
                                                                                    var v94 string = (v72 + (v72 + v72))
                                                                                    v3 += (-(v3))
                                                                                    _ = v94
                                                                                }
                                                                            }
                                                                        case "foo bar":
                                                                            v95 := GenS3{F0: v21, F1: 42}
                                                                            _ = v95
                                                                        case "\n\t":
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                fmt.Println(v3, v4)
                                                                            case "foo bar":
                                                                                v96 := GenS7{F0: v21}
                                                                                if v21 {
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                    fmt.Println(fi59, v81, v89)
                                                                                    v72 = v72
                                                                                } else {
                                                                                    v5 = append(v5, (fi59 - 732))
                                                                                }
                                                                                _ = v96
                                                                            case "a":
                                                                                v72 = v72
                                                                            default:
                                                                                v21 = (v72 <= "")
                                                                            }
                                                                        }
                                                                        _ = v89
                                                                        _ = v90
                                                                    } else {
                                                                        v21 = ((417 == 32767) && v21)
                                                                        v86 += (v86 - v81)
                                                                    }
                                                                }
                                                                if v21 {
                                                                    var v97 uint64 = 42
                                                                    switch fi59 {
                                                                    case 32767:
                                                                        var v98 int16 = -1
                                                                        _ = v98
                                                                    case -396:
                                                                        var v99 int = (fi58 << ((18446744073709551615) & 63))
                                                                        _ = v99
                                                                    }
                                                                    if (!(!v21)) {
                                                                        var v100 bool = false
                                                                        _ = v100
                                                                    }
                                                                    _ = v97
                                                                }
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    var v101 int64 = int64(v81)
                                                                    _ = v101
                                                                default:
                                                                    if (128 > 128) {
                                                                        switch (v72 + ((v72 + "a") + v72)) {
                                                                        case "a":
                                                                            v81 -= v86
                                                                        case "\n\t":
                                                                            var v102 bool = false
                                                                            v3 = (v3 | (v4 * (v4 ^ (^(v3)))))
                                                                            _ = v102
                                                                        case "hello":
                                                                            if (-72 <= 0) {
                                                                                fmt.Println(v87, v4, v86)
                                                                            } else {
                                                                                var v103 bool = v21
                                                                                delete(v1, 32767)
                                                                                if (18446744073709551615 >= 290) {
                                                                                    v104 := GenS8{F0: 29, F1: 42}
                                                                                    v105 := v2[(v72 + v72)]
                                                                                    _ = v104
                                                                                    _ = v105
                                                                                } else {
                                                                                    v81 = v86
                                                                                    v106 := GenS9{F0: 42, F1: 1}
                                                                                    _ = v106
                                                                                }
                                                                                _ = v103
                                                                            }
                                                                        default:
                                                                            if ((((v53 - v53) <= uint16(v81)) || (1e10 == (0.1 * -2.25))) && (2147483648 <= 65535)) {
                                                                                v107 := GenS10{F0: v21, F1: 0, F2: 9223372036854775807}
                                                                                var v108 int = fi58
                                                                                _ = v107
                                                                                _ = v108
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v81 = v81
                                                                        fmt.Println(v4, fi59, v3)
                                                                        fmt.Println(v21)
                                                                        var v109 uint8 = 198
                                                                        var v110 bool = true
                                                                        _ = v109
                                                                        _ = v110
                                                                    }
                                                                }
                                                                _ = v85
                                                                _ = v86
                                                                _ = v87
                                                            }
                                                        default:
                                                            v111 := GenS11{F0: (v53 % ((v53) | 1)), F1: (v72 + (v72 + ("" + (v72 + (v72 + (v72 + ((v72 + v72) + (v72 + v72)))))))), F2: 32767}
                                                            _ = v111
                                                        }
                                                        fmt.Println(fi58, v72)
                                                        v72 = "a"
                                                        _ = v81
                                                    }
                                                    var v112 uint8 = 255
                                                    _ = v112
                                                } else {
                                                    fmt.Println(v21, v53, v21)
                                                    v5 = append(v5, (fi58 ^ 1))
                                                    var v113 uint = 255
                                                    if v21 {
                                                        var v114 int32 = int32(v113)
                                                        _ = v114
                                                    }
                                                    fmt.Println(fi58, v21, v72)
                                                    _ = v113
                                                }
                                            }
                                            if v21 {
                                                var v115 int32 = (v4 >> ((2) & 63))
                                                var v116 int64 = -9223372036854775808
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                _ = v115
                                                _ = v116
                                            }
                                            switch ("a" + "a") {
                                            case "":
                                                v21 = v21
                                            case "a":
                                                v21 = false
                                            }
                                            var v117 bool = ((1.5 * 3.14159) < 1e10)
                                            _ = v117
                                        } else {
                                            var v118 int = (fi59 & fi58)
                                            _ = v118
                                        }
                                        if (!((v21 || (1.5 >= (1.5 / 3.14159))) || (!(("foo bar" + v72) <= ("a" + "foo bar"))))) {
                                            if (!v21) {
                                                _ = len(v2)
                                                fmt.Println(v53, fi59)
                                            } else {
                                                var v119 uint64 = 926
                                                _ = v119
                                            }
                                            if (v21 && (816 <= 982)) {
                                                fmt.Println(fi59, v21)
                                            }
                                        } else {
                                            var v120 int8 = 100
                                            switch fi59 {
                                            case 2147483647:
                                                var v121 int = -9223372036854775808
                                                _ = v121
                                            case 972:
                                                _ = len(v5)
                                            default:
                                                var v122 uint8 = 255
                                                _ = v122
                                            }
                                            _ = v120
                                        }
                                        _ = len(v2)
                                    }
                                    _ = v72
                                } else {
                                    if v21 {
                                        var v123 uint64 = 42
                                        _ = v123
                                    }
                                    v124 := GenS5{F0: -32768, F1: 1, F2: -9223372036854775808}
                                    var v125 uint16 = (v53 + v53)
                                    _ = v124
                                    _ = v125
                                }
                                _ = v64
                            } else {
                                _ = len(v5)
                                fmt.Println(v53)
                                if v21 {
                                    if (2 <= 0) {
                                        var v126 uint64 = 882
                                        switch (("a" + "") + ("foo bar" + ("a" + "\n\t"))) {
                                        case "hello":
                                            v126 -= v126
                                        }
                                        v127 := GenS8{F0: 2, F1: -654}
                                        if ((fi59 << ((1) & 63)) == (fi59 + (fi59 * (fi59 | (fi58 | fi58))))) {
                                            v128 := GenS3{F0: ((42 <= 127) && ((^(v126)) <= 0)), F1: 2147483647}
                                            _ = v128
                                        }
                                        _ = v126
                                        _ = v127
                                    }
                                    v129 := GenS2{F0: 18446744073709551615, F1: 42}
                                    if false {
                                        v130 := GenS12{F0: 1, F1: v21, F2: 1}
                                        v21 = false
                                        _ = v130
                                    } else {
                                        var v131 uint8 = 0
                                        _ = v131
                                    }
                                    v132 := GenS10{F0: v21, F1: 215, F2: 9223372036854775807}
                                    _ = v129
                                    _ = v132
                                }
                            }
                        }
                    case "\n\t":
                        v5 = append(v5, (fi58 - (fi59 >> ((0) & 63))))
                    default:
                        v3 += (v4 << ((0) & 63))
                    }
                    var v133 int32 = (v3 % (((-(v3))) | 1))
                    _ = v133
                    _ = v61
                } else {
                    if (!(v21 && true)) {
                        fmt.Println(v53, v4, fi58)
                    } else {
                        if (189 <= 42) {
                            switch -923 {
                            case 127:
                                if v21 {
                                    if (v21 || true) {
                                        fmt.Println(v3)
                                    }
                                }
                            case 0:
                                v53 += v53
                            case 4294967295:
                                v21 = (((3.14159 + 3.14159) / 3.14159) > -2.25)
                            default:
                                var v134 int32 = v3
                                _ = v134
                            }
                        }
                        var v135 uint16 = v53
                        v136 := GenS8{F0: 18446744073709551615, F1: 1000}
                        v137 := GenS13{F0: uint(v53), F1: 81, F2: 1}
                        _ = v135
                        _ = v136
                        _ = v137
                    }
                    if (-9223372036854775808 < 321) {
                        v138 := v1[47]
                        _ = v138
                    } else {
                        switch 4294967295 {
                        case 0:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        default:
                            v4 -= (^(v4))
                        }
                    }
                }
                v139 := GenS14{F0: v4}
                _ = fi59
                _ = v139
            }
            if v21 {
                for fi140 := 0; fi140 < 16; fi140++ {
                    v53 += v53
                    var v141 int64 = 0
                    _ = fi140
                    _ = v141
                }
            } else {
                v53 = (v53 - 343)
                var v142 int32 = v3
                _ = v142
            }
            fmt.Println(fi58, fi58)
            _ = fi58
        }
        v21 = (-102 == 1)
        _ = v57
    }
    if v21 {
        for fi143 := 0; fi143 < 2; fi143++ {
            v144 := GenS15{F0: 2, F1: 42, F2: v53}
            fmt.Println(v4, v21)
            _ = fi143
            _ = v144
        }
        switch -798 {
        case -368:
            v5 = append(v5, 2147483647)
        case -9223372036854775808:
            v145 := GenS12{F0: 42, F1: false, F2: 1000}
            _ = v145
        }
        var v146 int32 = (v4 % ((int32(v53)) | 1))
        _ = v146
    } else {
        switch ("\n\t" + ((("" + "hello") + "a") + ("" + "\n\t"))) {
        case "hello":
            v147 := GenS4{F0: (v21 && false)}
            _ = v147
        case "foo bar":
            var v148 float32 = (3.14159 * -2.25)
            _ = v148
        case "a":
            v4 = (v4 + (^(v3)))
        default:
            v149 := GenS16{F0: 42, F1: 2147483647}
            _ = v149
        }
    }
    if v21 {
        v53 -= (v53 + (v53 >> ((492) & 63)))
    }
    delete(v2, ("a" + "hello"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}
