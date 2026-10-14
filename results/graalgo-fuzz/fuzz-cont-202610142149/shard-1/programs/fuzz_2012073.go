package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }
type GenS2 struct { F0 bool; F1 uint; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 string; F1 uint16 }
type GenS11 struct { F0 uint32; F1 int32 }
type GenS12 struct { F0 string; F1 int16 }
type GenS13 struct { F0 int8; F1 int8 }
type GenS14 struct { F0 uint32; F1 int64 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int; F1 int64; F2 string }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 int8; F1 int8 }
type GenS19 struct { F0 uint64; F1 int32; F2 int }
type GenS20 struct { F0 int }
type GenS21 struct { F0 int; F1 int16 }
type GenS22 struct { F0 uint; F1 uint32 }
type GenS23 struct { F0 int8 }
type GenS24 struct { F0 uint64; F1 bool }
type GenS25 struct { F0 uint32; F1 uint16 }
type GenS26 struct { F0 int16 }
type GenS27 struct { F0 bool; F1 uint64; F2 uint64 }

func main() {
    v1 := []string{}
    var v2 uint8 = 69
    var v3 int = -1
    v4 := []string{"a", "", ""}
    var v5 uint8 = (v2 - v2)
    var v6 int32 = 627
    var v7 int64 = -1
    v8 := GenS1{F0: 18446744073709551615, F1: 42, F2: (true && (!(!(false && true))))}
    for fi9 := 0; fi9 < 15; fi9++ {
        if false {
            var v10 int8 = -68
            _ = len(v4)
            _ = v10
        }
        v11 := GenS1{F0: 51, F1: 42, F2: (true && false)}
        _ = fi9
        _ = v11
    }
    var v12 uint = 42
    v3 = v3
    if (1.5 < 0.1) {
        switch (v3 - 128) {
        case 1:
            switch (v3 * (v3 % (((v3 + v3)) | 1))) {
            case 576:
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    switch v3 {
                    case 4294967295:
                        if (v12 <= 255) {
                            var v13 int32 = v6
                            _ = v13
                        }
                    case -2147483648:
                        fmt.Println(v3, v6)
                    case 42:
                        v1 = append(v1, "foo bar")
                    }
                default:
                    var v14 uint = 534
                    _ = v14
                }
            case -374:
                for fi15 := 0; fi15 < 16; fi15++ {
                    if (843 < 1) {
                        fmt.Println(fi15)
                        for fi16 := 0; fi16 < 2; fi16++ {
                            v2 = v5
                            _ = fi16
                        }
                        for fi17 := 0; fi17 < 12; fi17++ {
                            var v18 int = (^(v3))
                            v19 := GenS2{F0: false, F1: v12, F2: (-(v3))}
                            v18 = int(v12)
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    } else {
                        switch ((("hello" + "foo bar") + "foo bar") + "a") {
                        case "a":
                            if false {
                                if (("hello" + "") <= "\n\t") {
                                    var v20 bool = ((1 > 891) && false)
                                    v21 := GenS2{F0: ((v20 || (v7 > (v7 * (v7 << ((uint(v2)) & 63))))) || (-114 != -128)), F1: (v12 & v12), F2: (v3 << ((v12) & 63))}
                                    _ = v20
                                    _ = v21
                                }
                                for fi22 := 0; fi22 < 10; fi22++ {
                                    var v23 int8 = -1
                                    v3 = fi22
                                    _ = fi22
                                    _ = v23
                                }
                                for fi24 := 0; fi24 < 4; fi24++ {
                                    if (((((3.14159 + 3.14159) / ((1.5 / (3.14159 + 0.0)) + (1.5 - 3.14159))) - 1.5) + float32((0.1 * 1.5))) <= -2.25) {
                                        var v25 uint32 = 2147483648
                                        var v26 uint32 = uint32(v12)
                                        v27 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: true}
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                    var v28 uint8 = (v2 / (((v2 & 127)) | 1))
                                    _ = fi24
                                    _ = v28
                                }
                            }
                        case "":
                            if (v3 < (fi15 << (((v12 / ((v12) | 1))) & 63))) {
                                var v29 uint64 = 1
                                for fi30 := 0; fi30 < 1; fi30++ {
                                    if len(v4) > 0 {
                                        v4[0] = ""
                                    }
                                    v31 := GenS3{F0: 1}
                                    v32 := GenS4{F0: (v12 | v12)}
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                                _ = v29
                            } else {
                                var v33 uint16 = 65535
                                v34 := GenS5{F0: 4294967295}
                                var v35 float64 = 0.0
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                        }
                        fmt.Println(v5, fi15, v5)
                    }
                    _ = fi15
                }
            }
        case -2147483648:
            if (0.1 <= ((3.14159 / 1e10) - -2.25)) {
                fmt.Println(v2)
                v3 -= v3
                for fi36 := 0; fi36 < 15; fi36++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v12 += 42
                    _ = fi36
                }
            } else {
                var v37 int16 = -1
                switch "hello" {
                case "a":
                    v38 := GenS5{F0: 513}
                    _ = v38
                case "":
                    switch v3 {
                    case 32768:
                        v2 = v2
                    case 4294967295:
                        fmt.Println(v12, v6, v7)
                    default:
                        v12 = (v12 % ((v12) | 1))
                    }
                case "\n\t":
                    for fi39 := 0; fi39 < 7; fi39++ {
                        v12 = v12
                        _ = fi39
                    }
                }
                for fi40 := 0; fi40 < 17; fi40++ {
                    v5 = v5
                    _ = fi40
                }
                _ = v37
            }
        case 347:
            v4 = append(v4, "\n\t")
        }
        var v41 int = v3
        var v42 float32 = 1.5
        _ = v41
        _ = v42
    }
    var v43 string = "foo bar"
    v44 := GenS1{F0: 562, F1: 59, F2: false}
    var v45 uint8 = (v2 >> ((v12) & 63))
    for fi46 := 0; fi46 < 10; fi46++ {
        v47 := GenS2{F0: true, F1: (v12 >> ((v12) & 63)), F2: (v3 | v3)}
        for fi48 := 0; fi48 < 9; fi48++ {
            switch (v43 + (v43 + (v43 + v43))) {
            case "a":
                var v49 int32 = (v6 * v6)
                _ = v49
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "hello":
                if (!(fi48 <= fi46)) {
                    v45 -= v5
                    v2 -= (v45 & (v45 - v45))
                    if false {
                        v5 -= uint8(v12)
                        var v50 int = (fi48 ^ int(v45))
                        _ = v50
                    }
                }
            }
            if (false || true) {
                switch (^(v3)) {
                case 160:
                    v51 := GenS1{F0: uint64(v12), F1: 0, F2: true}
                    _ = v51
                default:
                    if true {
                        v52 := GenS6{F0: -58}
                        v2 += (^(v2))
                        _ = v52
                    } else {
                        v1 = append(v1, v43)
                    }
                }
                switch (v43 + "hello") {
                case "hello":
                    v43 = v43
                    var v53 string = "foo bar"
                    if (!false) {
                        switch v3 {
                        case 32768:
                            v2 -= uint8(v12)
                        case 2147483647:
                            var v54 float32 = 0.1
                            _ = v54
                        case -1:
                            if len(v1) > 0 {
                                v1[0] = (("hello" + (("foo bar" + (v43 + v43)) + ("hello" + ((v43 + v43) + (v43 + ""))))) + (v43 + v43))
                            }
                        }
                        if ((0.0 * 0.0) < 0.0) {
                            v7 = 428
                        }
                    }
                    _ = v53
                case "":
                    var v55 bool = (0.1 > -2.25)
                    _ = v55
                default:
                    switch ((v43 + ((v43 + v43) + v43)) + v43) {
                    case "foo bar":
                        v5 = v45
                        var v56 string = v43
                        if (v56 >= v43) {
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                            switch ("foo bar" + v43) {
                            case "\n\t":
                                v57 := GenS7{F0: v43, F1: ((("\n\t" + v56) + v43) + (v43 + v56))}
                                _ = v57
                            case "foo bar":
                                var v58 int8 = 42
                                v4 = append(v4, "\n\t")
                                var v59 int32 = (^(v6))
                                if (-2.25 <= (float64(((0.0 - 0.0) + 1e10)) + 0.1)) {
                                    v45 -= (v45 % ((0) | 1))
                                    v59 = 1
                                    var v60 int = (fi48 >> (((v12 >> ((v12) & 63))) & 63))
                                    _ = v60
                                } else {
                                    v7 = (v7 | (-(v7)))
                                    v1 = append(v1, v43)
                                    v61 := GenS7{F0: (v56 + v56), F1: v56}
                                    _ = v61
                                }
                                if false {
                                    var v62 uint16 = 169
                                    if (true || false) {
                                        var v63 int = 0
                                        fmt.Println(v59)
                                        var v64 uint = v12
                                        _ = v63
                                        _ = v64
                                    } else {
                                        switch v43 {
                                        case "hello":
                                            switch ((v43 + (v56 + (v56 + "a"))) + v43) {
                                            case "\n\t":
                                                if true {
                                                    var v65 uint64 = 4294967295
                                                    _ = v65
                                                } else {
                                                    if true {
                                                        if (!(306 <= 2147483647)) {
                                                            var v66 uint = v12
                                                            var v67 int16 = -32768
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        if (!(!(!true))) {
                                                            v43 = (((v56 + v43) + v56) + ("" + (v43 + (v56 + (v43 + v56)))))
                                                        } else {
                                                            _ = len(v4)
                                                        }
                                                        v68 := GenS8{F0: 484, F1: uint32(v45)}
                                                        _ = v68
                                                    } else {
                                                        v43 = v56
                                                        v56 = v56
                                                        v2 -= (^(v2))
                                                    }
                                                }
                                            case "a":
                                                v58 += v58
                                            default:
                                                if false {
                                                    var v69 uint32 = 2147483648
                                                    var v70 uint16 = 255
                                                    if true {
                                                        v12 = (v12 * (v12 * (v12 >> ((v12) & 63))))
                                                        v12 = 1
                                                        if false {
                                                            var v71 uint = (v12 | (^(v12)))
                                                            switch ("foo bar" + v56) {
                                                            case "foo bar":
                                                                switch (fi48 - fi48) {
                                                                case -2147483648:
                                                                    v72 := GenS8{F0: 42, F1: v69}
                                                                    _ = v72
                                                                case -696:
                                                                    v62 = v62
                                                                case -9223372036854775808:
                                                                    var v73 uint64 = 1
                                                                    _ = v73
                                                                }
                                                            case "a":
                                                                if (2147483647 >= v59) {
                                                                    v43 = "foo bar"
                                                                }
                                                            case "\n\t":
                                                                switch 0 {
                                                                case -694:
                                                                    v59 -= v6
                                                                }
                                                            }
                                                            var v74 int16 = 1000
                                                            _ = v71
                                                            _ = v74
                                                        }
                                                        fmt.Println(v70)
                                                        v75 := GenS9{F0: uint32(v70)}
                                                        _ = v75
                                                    } else {
                                                        v76 := GenS10{F0: (("" + "") + "hello"), F1: (^(v62))}
                                                        _ = v76
                                                    }
                                                    _ = v69
                                                    _ = v70
                                                } else {
                                                    var v77 uint8 = v5
                                                    if false {
                                                        fmt.Println(fi48, v6, v2)
                                                        var v78 int16 = 0
                                                        _ = v78
                                                    }
                                                    _ = v77
                                                }
                                            }
                                        default:
                                            var v79 uint32 = 2
                                            _ = v79
                                        }
                                        var v80 uint8 = v2
                                        _ = v80
                                    }
                                    var v81 float32 = (0.1 - 3.14159)
                                    var v82 float32 = v81
                                    _ = v62
                                    _ = v81
                                    _ = v82
                                }
                                v83 := GenS10{F0: v43, F1: 908}
                                _ = v58
                                _ = v59
                                _ = v83
                            }
                        }
                        _ = v56
                    case "\n\t":
                        var v84 uint = (v12 + v12)
                        _ = v84
                    default:
                        var v85 int64 = v7
                        var v86 uint = (v12 % ((v12) | 1))
                        _ = v85
                        _ = v86
                    }
                }
            } else {
                var v87 uint8 = v5
                v7 -= v7
                fmt.Println(v87, v3)
                _ = v87
            }
            switch "" {
            case "foo bar":
                if ((fi46 + int(v6)) <= fi48) {
                    v88 := GenS11{F0: 673, F1: v6}
                    v6 = (v6 % (((-(v6))) | 1))
                    fmt.Println(v43, fi46, v12)
                    _ = v88
                } else {
                    _ = len(v1)
                    if (!(false || (false || true))) {
                        if (!(42 < 2)) {
                            var v89 uint64 = 121
                            switch (v43 + ("a" + ("" + ("hello" + (v43 + v43))))) {
                            case "a":
                                var v90 uint16 = 1
                                _ = v90
                            case "\n\t":
                                var v91 float32 = 1e10
                                _ = v91
                            case "":
                                var v92 uint8 = (v2 - v2)
                                _ = v92
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            _ = v89
                        } else {
                            fmt.Println(v7, v3)
                        }
                        v45 += (^(v5))
                        fmt.Println(v7, v43, v43)
                    } else {
                        if ((!true) && true) {
                            var v93 int32 = v6
                            if true {
                                switch ("foo bar" + (((v43 + v43) + v43) + "\n\t")) {
                                case "hello":
                                    v94 := GenS8{F0: 42, F1: 25}
                                    _ = v94
                                default:
                                    v93 = v93
                                }
                                var v95 string = v43
                                _ = v95
                            } else {
                                var v96 int32 = v93
                                _ = v96
                            }
                            v97 := GenS7{F0: "foo bar", F1: v43}
                            _ = v93
                            _ = v97
                        }
                        v98 := GenS2{F0: (0 >= 65535), F1: 42, F2: (fi46 - 4294967295)}
                        switch (fi48 | (fi48 % (((v3 * (fi46 >> (((v12 | v12)) & 63)))) | 1))) {
                        case 423:
                            var v99 bool = true
                            _ = v99
                        default:
                            var v100 int64 = v7
                            _ = v100
                        }
                        _ = v98
                    }
                    v101 := GenS11{F0: uint32(v2), F1: (-(v6))}
                    _ = v101
                }
            case "hello":
                switch v43 {
                case "hello":
                    var v102 uint64 = 433
                    if false {
                        v45 -= (v45 & 128)
                        v3 = fi48
                    } else {
                        if ((42 < 588) || false) {
                            fmt.Println(v3, v2, v102)
                            var v103 float64 = 1e10
                            _ = v103
                        }
                    }
                    if false {
                        if len(v4) > 0 {
                            v4[0] = v43
                        }
                    } else {
                        v102 = v102
                        if (1e10 != (3.14159 + 1.5)) {
                            var v104 int8 = 0
                            if (!(true && true)) {
                                v43 = ((v43 + (v43 + "a")) + v43)
                                v104 -= int8(v7)
                                if (false || (((v102 | 763) != v102) || true)) {
                                    if true {
                                        switch fi46 {
                                        case -886:
                                            switch "a" {
                                            case "a":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case "\n\t":
                                                if (!(v104 > (v104 | v104))) {
                                                    v12 = (v12 * (^(v12)))
                                                    v105 := GenS12{F0: v43, F1: int16(v7)}
                                                    var v106 uint32 = 2147483647
                                                    _ = v105
                                                    _ = v106
                                                }
                                            }
                                        case 4294967295:
                                            if true {
                                                var v107 float64 = (0.0 / float64(((float32((-2.25 + 0.0)) - (0.0 * 1e10)) + ((0.1 + (1e10 * 1e10)) / -2.25))))
                                                if ((v107 * v107) < v107) {
                                                    var v108 bool = false
                                                    v109 := GenS13{F0: 1, F1: (v104 / (((v104 & v104)) | 1))}
                                                    _ = v108
                                                    _ = v109
                                                } else {
                                                    v110 := GenS14{F0: uint32(v2), F1: v7}
                                                    _ = v110
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v107
                                            } else {
                                                v111 := GenS2{F0: (v6 <= (v6 << ((v12) & 63))), F1: v12, F2: 742}
                                                if (42 > uint8(v102)) {
                                                    v3 = (fi48 % ((-699) | 1))
                                                    if (!true) {
                                                        switch ("a" + (v43 + ("\n\t" + ("a" + v43)))) {
                                                        case "a":
                                                            fmt.Println(v43)
                                                        case "\n\t":
                                                            var v112 int16 = 683
                                                            _ = v112
                                                        case "hello":
                                                            if ((v12 >> ((v12) & 63)) >= (v12 - v12)) {
                                                                var v113 uint64 = v102
                                                                var v114 int8 = 100
                                                                v3 = 4294967295
                                                                v114 = 110
                                                                _ = v113
                                                                _ = v114
                                                            } else {
                                                                if (v6 <= (v6 >> (((v12 * v12)) & 63))) {
                                                                    v115 := GenS11{F0: 2147483648, F1: (v6 | (^(v6)))}
                                                                    if (v102 > 4294967295) {
                                                                        switch v43 {
                                                                        case "hello":
                                                                            v116 := GenS4{F0: (v12 / ((v12) | 1))}
                                                                            _ = v116
                                                                        case "a":
                                                                            fmt.Println(v104, v6, v12)
                                                                        }
                                                                        if true {
                                                                            if true {
                                                                                if false {
                                                                                    v104 += int8(v12)
                                                                                } else {
                                                                                    var v117 float32 = 3.14159
                                                                                    _ = v117
                                                                                }
                                                                                v3 = fi46
                                                                                switch v43 {
                                                                                case "\n\t":
                                                                                    fmt.Println(v3)
                                                                                case "a":
                                                                                    v102 = v102
                                                                                case "foo bar":
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = ("hello" + v43)
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v118 int64 = v7
                                                                            _ = v118
                                                                        }
                                                                    } else {
                                                                        v6 = (^(v6))
                                                                        switch v43 {
                                                                        case "a":
                                                                            if (v104 == 53) {
                                                                                switch fi48 {
                                                                                case -2147483648:
                                                                                    v5 -= v5
                                                                                default:
                                                                                    v12 += v12
                                                                                }
                                                                                var v119 uint16 = uint16(v102)
                                                                                fmt.Println(fi48, v104)
                                                                                _ = v119
                                                                            } else {
                                                                                if (!true) {
                                                                                    var v120 int32 = v6
                                                                                    var v121 uint = v12
                                                                                    var v122 uint16 = 1
                                                                                    _ = v120
                                                                                    _ = v121
                                                                                    _ = v122
                                                                                }
                                                                            }
                                                                            if true {
                                                                                switch (-(fi46)) {
                                                                                case 4294967295:
                                                                                    v123 := GenS14{F0: 4294967295, F1: 899}
                                                                                    _ = v123
                                                                                case 0:
                                                                                    fmt.Println(v2)
                                                                                case -432:
                                                                                    var v124 uint64 = 4294967295
                                                                                    _ = v124
                                                                                default:
                                                                                    v125 := GenS13{F0: -54, F1: 127}
                                                                                    _ = v125
                                                                                }
                                                                                var v126 float32 = float32(((3.14159 - ((1e10 / 1.0) - (float64(-2.25) + ((float64((-2.25 + 0.0)) * (float64((-2.25 * 0.0)) / (((1.5 / (0.1 * (float64(3.14159) - (0.1 * ((1e10 + 1e10) / 1.0))))) - -2.25) * ((1.5 / 1e10) - 3.14159)))) - (0.1 / 0.1))))) * 0.1))
                                                                                _ = v126
                                                                            }
                                                                            v127 := GenS10{F0: v43, F1: 255}
                                                                            _ = v127
                                                                        case "hello":
                                                                            if true {
                                                                                var v128 uint16 = uint16(v45)
                                                                                v129 := GenS15{F0: v2}
                                                                                v12 = (v12 >> ((318) & 63))
                                                                                _ = v128
                                                                                _ = v129
                                                                            } else {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = ""
                                                                                }
                                                                                v130 := GenS6{F0: v104}
                                                                                _ = v130
                                                                            }
                                                                        case "foo bar":
                                                                            var v131 string = v43
                                                                            _ = v131
                                                                        }
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            if false {
                                                                                var v132 int = fi46
                                                                                v102 -= 1
                                                                                _ = v132
                                                                            } else {
                                                                                if ((v5 | (v45 >> (((v12 + (v12 + 18446744073709551615))) & 63))) == (v45 >> ((v12) & 63))) {
                                                                                    v133 := GenS14{F0: 2147483647, F1: int64(v45)}
                                                                                    var v134 int32 = (v6 + (-(v6)))
                                                                                    if (v102 >= 310) {
                                                                                        v2 += v45
                                                                                        switch fi48 {
                                                                                        case 9223372036854775807:
                                                                                            v135 := GenS14{F0: 123, F1: -1}
                                                                                            _ = v135
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v43, v134, v12)
                                                                                        v136 := GenS15{F0: 175}
                                                                                        if false {
                                                                                            v7 = -1
                                                                                        }
                                                                                        _ = v136
                                                                                    }
                                                                                    _ = v133
                                                                                    _ = v134
                                                                                }
                                                                            }
                                                                        case "\n\t":
                                                                            if len(v1) > 0 {
                                                                                v1[0] = ((v43 + v43) + v43)
                                                                            }
                                                                        case "a":
                                                                            v137 := GenS16{F0: v3, F1: 9223372036854775807, F2: (v43 + "foo bar")}
                                                                            _ = v137
                                                                        default:
                                                                            var v138 int64 = (v7 - (v7 >> (((v12 / ((v12) | 1))) & 63)))
                                                                            _ = v138
                                                                        }
                                                                    }
                                                                    if (-1 != -432) {
                                                                        if (1 > -58) {
                                                                            var v139 int16 = 42
                                                                            v140 := GenS1{F0: v102, F1: v102, F2: true}
                                                                            _ = v139
                                                                            _ = v140
                                                                        }
                                                                    } else {
                                                                        switch ("foo bar" + (v43 + v43)) {
                                                                        case "a":
                                                                            if (v102 < 317) {
                                                                                v43 = v43
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                var v141 uint32 = uint32(v12)
                                                                                _ = v141
                                                                            } else {
                                                                                v43 = v43
                                                                                if (!(v7 < v7)) {
                                                                                    switch v43 {
                                                                                    case "":
                                                                                        v104 = v104
                                                                                    case "a":
                                                                                        v5 -= v45
                                                                                    case "\n\t":
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                    default:
                                                                                        v5 = v5
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "":
                                                                            fmt.Println(fi46, fi48, v6)
                                                                        case "foo bar":
                                                                            v7 = (v7 << ((v12) & 63))
                                                                        default:
                                                                            fmt.Println(v3, v45, v102)
                                                                        }
                                                                        switch v3 {
                                                                        case 1:
                                                                            var v142 uint64 = v102
                                                                            _ = v142
                                                                        case 2147483647:
                                                                            v143 := GenS17{F0: 2}
                                                                            _ = v143
                                                                        }
                                                                    }
                                                                    _ = v115
                                                                }
                                                                switch v43 {
                                                                case "foo bar":
                                                                    v144 := GenS18{F0: 127, F1: v104}
                                                                    _ = v144
                                                                case "\n\t":
                                                                    var v145 int16 = 1000
                                                                    _ = v145
                                                                case "a":
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                                v146 := GenS4{F0: v12}
                                                                _ = v146
                                                            }
                                                        default:
                                                            var v147 bool = (float64((0.0 - 1.5)) == 3.14159)
                                                            _ = v147
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    } else {
                                                        var v148 uint16 = 997
                                                        if true {
                                                            var v149 int8 = (v104 & v104)
                                                            _ = v149
                                                        } else {
                                                            if ((v7 - (^(v7))) == v7) {
                                                                v6 -= v6
                                                            } else {
                                                                v150 := GenS19{F0: v102, F1: (v6 >> ((v12) & 63)), F2: fi46}
                                                                _ = v150
                                                            }
                                                        }
                                                        if (1 >= 32767) {
                                                            var v151 uint8 = 10
                                                            v152 := GenS20{F0: int(v12)}
                                                            var v153 int8 = v104
                                                            _ = v151
                                                            _ = v152
                                                            _ = v153
                                                        } else {
                                                            var v154 uint16 = v148
                                                            _ = v154
                                                        }
                                                        if true {
                                                            v3 = (^(v3))
                                                            v104 -= -56
                                                        }
                                                        var v155 int = (fi46 / (((fi48 * 32768)) | 1))
                                                        _ = v148
                                                        _ = v155
                                                    }
                                                }
                                                _ = v111
                                            }
                                        case 0:
                                            v43 = v43
                                        }
                                        v45 += (v2 >> ((uint(v2)) & 63))
                                        v4 = append(v4, ((("foo bar" + (v43 + ("a" + ""))) + ((v43 + v43) + v43)) + (v43 + (v43 + v43))))
                                        v5 = (v5 | (v45 << (((v12 - v12)) & 63)))
                                    } else {
                                        v3 += fi46
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    v6 -= -2147483648
                                }
                            }
                            v156 := GenS21{F0: (-(v3)), F1: int16(v102)}
                            var v157 string = v43
                            if true {
                                var v158 float64 = (1e10 - 1e10)
                                fmt.Println(v43, v6)
                                _ = v158
                            }
                            switch -760 {
                            case 42:
                                var v159 uint = v12
                                _ = v159
                            case -9223372036854775808:
                                if ((v5 == v5) || (float64(float32(-2.25)) != 3.14159)) {
                                    var v160 uint = (^(v12))
                                    _ = v160
                                }
                            case -727:
                                if true {
                                    var v161 uint16 = 0
                                    _ = v161
                                } else {
                                    v6 += v6
                                }
                            default:
                                if false {
                                    if (1 != 731) {
                                        var v162 uint16 = 0
                                        switch (v3 ^ (v3 / (((v3 % ((v3) | 1))) | 1))) {
                                        case -1:
                                            v104 = 42
                                        }
                                        _ = v162
                                    } else {
                                        v4 = append(v4, (v43 + ""))
                                    }
                                    if (128 <= v5) {
                                        var v163 uint32 = 1
                                        v163 += v163
                                        _ = v163
                                    }
                                }
                            }
                            v164 := GenS22{F0: v12, F1: 2147483648}
                            var v165 uint = (v12 + 65535)
                            _ = v104
                            _ = v156
                            _ = v157
                            _ = v164
                            _ = v165
                        } else {
                            if false {
                                v7 += v7
                                v4 = append(v4, (((v43 + "hello") + (v43 + v43)) + "a"))
                                v166 := GenS23{F0: 42}
                                _ = v166
                            }
                            if (!(((v5 / ((v45) | 1)) >= (v5 / ((v45) | 1))) || true)) {
                                if ("" < "hello") {
                                    var v167 int64 = (v7 / ((42) | 1))
                                    fmt.Println(v5)
                                    _ = v167
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                fmt.Println(v6, v43)
                                fmt.Println(v3)
                            }
                            var v168 float32 = (1.5 * 1.5)
                            _ = v168
                        }
                    }
                    _ = v102
                case "a":
                    var v169 uint32 = 4294967295
                    _ = v169
                case "foo bar":
                    v1 = append(v1, "\n\t")
                }
            default:
                switch "hello" {
                case "":
                    var v170 bool = true
                    _ = v170
                case "a":
                    fmt.Println(v12)
                default:
                    fmt.Println(v5)
                }
            }
            var v171 int32 = (v6 / (((v6 - (v6 * (v6 & v6)))) | 1))
            _ = fi48
            _ = v171
        }
        v5 = (v45 + v5)
        _ = fi46
        _ = v47
    }
    if false {
        var v172 int64 = v7
        _ = v172
    } else {
        var v173 int = 32767
        for fi174 := 0; fi174 < 15; fi174++ {
            if true {
                fmt.Println(v6)
                if (v3 == fi174) {
                    v175 := GenS24{F0: 1, F1: (32767 >= 42)}
                    v43 = "foo bar"
                    for fi176 := 0; fi176 < 13; fi176++ {
                        if (32767 >= 32767) {
                            var v177 int8 = 100
                            _ = v177
                        }
                        if false {
                            if true {
                                switch (v3 - (fi176 % ((v3) | 1))) {
                                case -9223372036854775808:
                                    v45 = v5
                                case 127:
                                    v178 := GenS25{F0: 73, F1: 32767}
                                    _ = v178
                                }
                                var v179 bool = false
                                if (-1 != 32767) {
                                    v45 -= (^(v45))
                                    v5 = v2
                                    var v180 uint32 = 100
                                    _ = v180
                                } else {
                                    v45 = (^(v5))
                                    if (1e10 >= -2.25) {
                                        var v181 int8 = 127
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v181
                                    } else {
                                        var v182 string = v43
                                        _ = len(v1)
                                        fmt.Println(v3)
                                        _ = v182
                                    }
                                    fmt.Println(v6, v5)
                                }
                                v173 += (fi174 % (((^(fi176))) | 1))
                                if true {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v183 bool = true
                                    if true {
                                        v184 := GenS26{F0: 727}
                                        _ = v184
                                    } else {
                                        if false {
                                            v45 = (v2 & v2)
                                        }
                                    }
                                    _ = v183
                                }
                                switch (fi176 / ((v173) | 1)) {
                                case -885:
                                    v43 = ((v43 + v43) + v43)
                                case -804:
                                    v43 = "foo bar"
                                case 127:
                                    v6 = 42
                                }
                                _ = v179
                            } else {
                                v3 += (fi176 | fi176)
                                if true {
                                    var v185 float64 = -2.25
                                    switch v43 {
                                    case "":
                                        if true {
                                            var v186 uint32 = 0
                                            v187 := GenS5{F0: 18446744073709551615}
                                            v1 = append(v1, "foo bar")
                                            _ = v186
                                            _ = v187
                                        } else {
                                            var v188 float32 = ((0.0 - 0.1) / 0.1)
                                            v185 = float64(v188)
                                            v12 += v12
                                            _ = v188
                                        }
                                    }
                                    var v189 float32 = 0.0
                                    _ = v185
                                    _ = v189
                                }
                                if (!(!true)) {
                                    v5 += v2
                                } else {
                                    var v190 float64 = 1e10
                                    _ = v190
                                }
                                _ = len(v4)
                            }
                        }
                        _ = fi176
                    }
                    _ = v175
                }
                for fi191 := 0; fi191 < 3; fi191++ {
                    v192 := GenS27{F0: false, F1: 2, F2: 1}
                    var v193 int16 = 42
                    fmt.Println(v6)
                    _ = fi191
                    _ = v192
                    _ = v193
                }
            }
            var v194 float32 = -2.25
            _ = fi174
            _ = v194
        }
        var v195 float64 = -2.25
        _ = v173
        _ = v195
    }
    var v196 uint8 = 255
    for fi197 := 0; fi197 < 8; fi197++ {
        var v198 uint64 = 992
        _ = fi197
        _ = v198
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v196:", v196)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v196
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
