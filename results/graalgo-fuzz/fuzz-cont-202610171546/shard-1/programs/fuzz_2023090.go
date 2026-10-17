package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 int8; F2 uint16 }
type GenS5 struct { F0 int32; F1 uint16; F2 int32 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS7 struct { F0 string; F1 int8 }
type GenS8 struct { F0 int32; F1 bool; F2 int8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint32; F1 int8; F2 int16 }
type GenS11 struct { F0 uint; F1 string; F2 uint32 }
type GenS12 struct { F0 int16; F1 uint32 }
type GenS13 struct { F0 uint32; F1 uint; F2 int32 }
type GenS14 struct { F0 int64; F1 int32; F2 uint8 }
type GenS15 struct { F0 int32; F1 uint32; F2 int }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 bool; F1 uint; F2 string }
func genFn0(p0 uint8, p1 int16) int8 {
    v4 := GenS1{F0: (p1 - p1), F1: ("a" + "\n\t"), F2: -1}
    var v5 int8 = -128
    v6 := []int{-337}
    p1 = (-(p1))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int32, p1 int64) int32 {
    var v7 int = 32767
    v8 := []int{v7, 42, (v7 - v7)}
    fmt.Println(v7, p1, p0)
    if true {
        _ = len(v8)
    } else {
        var v9 int32 = (p0 * p0)
        switch 0 {
        case 4294967295:
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := GenS1{F0: -83, F1: "a", F2: 42}
                _ = fi10
                _ = v11
            }
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 1
    v3 := []uint32{0, 4294967295, 393, 557}
    var v12 uint16 = 32767
    fmt.Println(v2)
    v12 = v12
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            v2 -= (v2 - v2)
            v15 := GenS2{F0: 4294967295, F1: 2147483647}
            v2 = v2
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v12 = (v12 >> (((v2 % (((^(v2))) | 1))) & 63))
        for fi17 := 0; fi17 < 15; fi17++ {
            v2 -= (v2 % ((v2) | 1))
            switch "foo bar" {
            case "":
                switch (fi16 % ((fi16) | 1)) {
                case -1:
                    v18 := GenS3{F0: (0.1 > 0.0)}
                    _ = v18
                default:
                    if false {
                        fmt.Println(v12, v2, fi17)
                        delete(v1, fi16)
                        switch (fi16 * fi17) {
                        case -2147483648:
                            switch (fi17 | fi17) {
                            case 32768:
                                v2 += (v2 * (v2 * (^(v2))))
                                var v19 uint16 = (v12 * (v12 - v12))
                                switch (fi16 << ((v2) & 63)) {
                                case -868:
                                    switch "foo bar" {
                                    case "foo bar":
                                        v2 = (v2 * 726)
                                    case "\n\t":
                                        v20 := GenS3{F0: true}
                                        _ = v20
                                    case "a":
                                        switch fi16 {
                                        case -456:
                                            v12 -= v12
                                        case 4294967295:
                                            if ((933 != 2147483647) || (!false)) {
                                                fmt.Println(fi17)
                                                v2 = (v2 + 0)
                                                v21 := GenS1{F0: 1, F1: "hello", F2: -32768}
                                                v19 = (v12 % (((v12 % (((v19 / ((v19) | 1))) | 1))) | 1))
                                                _ = v21
                                            } else {
                                                v22 := GenS4{F0: -2147483648, F1: 42, F2: v19}
                                                _ = v22
                                            }
                                        case 32767:
                                            v23 := GenS4{F0: int32(v19), F1: 127, F2: v12}
                                            _ = v23
                                        }
                                    default:
                                        if false {
                                            var v24 int64 = 9223372036854775807
                                            v19 -= 1
                                            if true {
                                                var v25 float64 = 1.5
                                                if (v25 == -2.25) {
                                                    v2 = (^(v2))
                                                    switch (-(fi16)) {
                                                    case 0:
                                                        fmt.Println(v24)
                                                    case 2147483647:
                                                        fmt.Println(v19, fi16)
                                                    case -1:
                                                        var v26 uint32 = 2147483647
                                                        _ = v26
                                                    default:
                                                        v2 = v2
                                                        switch ("foo bar" + "hello") {
                                                        case "hello":
                                                            if false {
                                                                var v27 uint32 = 660
                                                                _ = v27
                                                            } else {
                                                                fmt.Println(v19)
                                                            }
                                                        default:
                                                            var v28 int = fi17
                                                            _ = v28
                                                        }
                                                        _ = len(v1)
                                                    }
                                                    var v29 uint32 = 528
                                                    _ = v29
                                                }
                                                _ = v25
                                            }
                                            _ = v24
                                        }
                                    }
                                case -950:
                                    var v30 int32 = -2147483648
                                    _ = v30
                                default:
                                    switch (fi17 - (fi17 * -1)) {
                                    case -1:
                                        var v31 bool = false
                                        _ = v31
                                    default:
                                        fmt.Println(v12)
                                    }
                                }
                                if true {
                                    switch ("\n\t" + ("\n\t" + ((("hello" + "a") + "a") + "a"))) {
                                    case "a":
                                        v32 := v1[fi16]
                                        _ = v32
                                    case "\n\t":
                                        v19 = v12
                                    default:
                                        v19 += v12
                                    }
                                    v19 = (v12 | (v12 ^ (v12 % ((v19) | 1))))
                                }
                                if (-488 <= 863) {
                                    v2 += v2
                                    fmt.Println(fi16)
                                } else {
                                    v2 += v2
                                    if (1e10 == -2.25) {
                                        fmt.Println(fi16, fi17, fi16)
                                        if false {
                                            v33 := GenS5{F0: -2147483648, F1: v12, F2: 642}
                                            v2 += (^(v2))
                                            _ = len(v1)
                                            _ = v33
                                        } else {
                                            v34 := GenS2{F0: 802, F1: fi16}
                                            _ = v34
                                        }
                                    } else {
                                        _ = len(v3)
                                        fmt.Println(fi17, fi17, fi17)
                                    }
                                }
                                _ = v19
                            case -408:
                                var v35 int64 = 285
                                _ = v35
                            case 597:
                                v36 := GenS2{F0: 4294967295, F1: fi17}
                                _ = v36
                            }
                        case 97:
                            var v37 int32 = genFn1(-1, 9223372036854775807)
                            _ = v37
                        default:
                            switch fi17 {
                            case -1:
                                var v38 string = "\n\t"
                                _ = v38
                            case 32767:
                                v39 := GenS2{F0: 42, F1: 660}
                                _ = v39
                            default:
                                switch "" {
                                case "a":
                                    var v40 uint8 = 41
                                    _ = v40
                                }
                            }
                        }
                        var v41 int32 = 42
                        fmt.Println(v41)
                        _ = v41
                    } else {
                        fmt.Println(fi16)
                        if false {
                            v42 := v1[-2147483648]
                            v3 = append(v3, 65535)
                            var v43 uint16 = (v12 % (((^(v12))) | 1))
                            v42 = v42
                            _ = v42
                            _ = v43
                        }
                    }
                }
                var v44 float32 = (3.14159 + 1.5)
                var v45 int8 = int8(v12)
                _ = v44
                _ = v45
            default:
                if (fi16 > fi17) {
                    if true {
                        fmt.Println(fi16, fi17)
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v46 uint8 = 1
                            _ = v46
                        }
                    } else {
                        switch (fi17 << (((^(v2))) & 63)) {
                        case -2147483648:
                            fmt.Println(v12)
                        case 2147483647:
                            v47 := GenS6{F0: false, F1: 274, F2: 128}
                            _ = v47
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    if true {
                        if true {
                            if true {
                                if false {
                                    fmt.Println(v2, fi17)
                                    v48 := GenS1{F0: -32768, F1: "hello", F2: -1}
                                    _ = v48
                                } else {
                                    if false {
                                        var v49 uint16 = (v12 | (^(v12)))
                                        _ = v49
                                    }
                                    v2 = v2
                                    v2 = v2
                                }
                                switch "hello" {
                                case "foo bar":
                                    var v50 int16 = 676
                                    _ = v50
                                case "\n\t":
                                    if true {
                                        var v51 int64 = 9223372036854775807
                                        v51 = 9223372036854775807
                                        _ = v51
                                    } else {
                                        v2 += (v2 % ((v2) | 1))
                                    }
                                case "":
                                    if false {
                                        v52 := GenS6{F0: (2 != 42), F1: 4294967295, F2: 2}
                                        _ = v52
                                    } else {
                                        if (("" + "foo bar") == "hello") {
                                            v2 = (v2 & v2)
                                            v12 -= (v12 + 0)
                                        }
                                    }
                                }
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = 840
                                    }
                                    v53 := GenS7{F0: "", F1: -36}
                                    _ = v53
                                } else {
                                    var v54 int8 = genFn0(4, 1)
                                    v12 = (v12 - v12)
                                    if (!false) {
                                        v54 -= v54
                                        v12 -= uint16(v2)
                                        switch fi16 {
                                        case 1:
                                            var v55 uint32 = 65535
                                            _ = v55
                                        case 857:
                                            v56 := GenS8{F0: 1, F1: true, F2: genFn0(128, 1000)}
                                            _ = v56
                                        }
                                    } else {
                                        _ = len(v3)
                                        switch 680 {
                                        case -2147483648:
                                            v57 := GenS9{F0: 485}
                                            _ = v57
                                        default:
                                            switch "\n\t" {
                                            case "foo bar":
                                                if true {
                                                    v58 := GenS6{F0: ("" != (("foo bar" + "\n\t") + ("" + "foo bar"))), F1: 2, F2: 1}
                                                    _ = v58
                                                }
                                                fmt.Println(fi16, fi17, v54)
                                                switch "\n\t" {
                                                case "\n\t":
                                                    _ = len(v1)
                                                }
                                                if (true && true) {
                                                    switch "" {
                                                    case "foo bar":
                                                        v54 = genFn0(1, 1000)
                                                        if (((!((((-424 <= 42) || false) || (!true)) && false)) || ((^(v12)) != v12)) && true) {
                                                            if (!(2147483647 > 0)) {
                                                                switch fi17 {
                                                                case 128:
                                                                    v2 = (v2 % ((v2) | 1))
                                                                case 694:
                                                                    v2 = 18446744073709551615
                                                                case 1:
                                                                    var v59 int = fi16
                                                                    _ = v59
                                                                }
                                                                fmt.Println(fi16, v54)
                                                            } else {
                                                                if ((42 < 769) || (true && (42 != 42))) {
                                                                    fmt.Println(fi16, v54, fi16)
                                                                    var v60 int64 = -645
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 351
                                                                    }
                                                                    v61 := GenS6{F0: (1.5 < 0.1), F1: 1, F2: 127}
                                                                    _ = v60
                                                                    _ = v61
                                                                } else {
                                                                    var v62 uint32 = 4294967295
                                                                    var v63 int16 = 42
                                                                    v64 := GenS7{F0: "foo bar", F1: 0}
                                                                    _ = v62
                                                                    _ = v63
                                                                    _ = v64
                                                                }
                                                                if (false || (v54 < v54)) {
                                                                    v2 += (v2 + (v2 & v2))
                                                                    fmt.Println(v12, v12, v12)
                                                                } else {
                                                                    var v65 float32 = 1e10
                                                                    _ = v65
                                                                }
                                                                switch "a" {
                                                                case "":
                                                                    v2 = v2
                                                                    var v66 int = (^(fi17))
                                                                    v67 := GenS3{F0: false}
                                                                    _ = v66
                                                                    _ = v67
                                                                case "hello":
                                                                    var v68 string = (((("a" + "a") + ("\n\t" + ("hello" + "foo bar"))) + "foo bar") + "a")
                                                                    _ = v68
                                                                default:
                                                                    v12 = v12
                                                                }
                                                            }
                                                        } else {
                                                            var v69 int32 = 0
                                                            var v70 int8 = genFn0(2, 756)
                                                            _ = v69
                                                            _ = v70
                                                        }
                                                        var v71 int64 = 9223372036854775807
                                                        var v72 int8 = v54
                                                        var v73 bool = (466 >= 0)
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                    default:
                                                        var v74 uint16 = v12
                                                        _ = v74
                                                    }
                                                } else {
                                                    switch (fi16 / (((fi17 >> ((v2) & 63))) | 1)) {
                                                    case -2147483648:
                                                        v3 = append(v3, 453)
                                                    }
                                                    if false {
                                                        _ = len(v3)
                                                        v75 := v1[0]
                                                        if false {
                                                            v76 := GenS4{F0: 857, F1: (v54 ^ genFn0(42, 0)), F2: (v12 - 0)}
                                                            _ = v76
                                                        }
                                                        v75 = (v75 - fi17)
                                                        _ = v75
                                                    }
                                                    v77 := GenS1{F0: 1000, F1: "foo bar", F2: 0}
                                                    var v78 int8 = 71
                                                    v2 += 185
                                                    v54 += v54
                                                    _ = v77
                                                    _ = v78
                                                }
                                            case "\n\t":
                                                fmt.Println(v54, v12, fi16)
                                            case "hello":
                                                v79 := GenS1{F0: 1000, F1: "hello", F2: -32768}
                                                _ = v79
                                            }
                                        }
                                        fmt.Println(v2)
                                    }
                                    _ = v54
                                }
                            }
                            fmt.Println(fi17, v2)
                            v12 = 255
                            var v80 float64 = 1.5
                            _ = v80
                        }
                    } else {
                        v81 := GenS10{F0: 4294967295, F1: 1, F2: 32767}
                        var v82 uint16 = v12
                        _ = v81
                        _ = v82
                    }
                } else {
                    switch "a" {
                    case "a":
                        v83 := GenS7{F0: ("hello" + "a"), F1: 25}
                        _ = v83
                    case "foo bar":
                        v84 := GenS8{F0: -575, F1: false, F2: 100}
                        v2 -= 4294967295
                        _ = v84
                    case "":
                        fmt.Println(fi17, v12)
                    default:
                        var v85 int = fi17
                        _ = v85
                    }
                    v86 := GenS10{F0: 65535, F1: 127, F2: 1000}
                    _ = v86
                }
            }
            if false {
                var v87 uint64 = 18446744073709551615
                _ = v87
            }
            _ = fi17
        }
        v12 -= (v12 << ((v2) & 63))
        _ = fi16
    }
    if true {
        switch 32768 {
        case 2147483647:
            for fi88 := 0; fi88 < 12; fi88++ {
                v2 = (^(v2))
                var v89 int64 = -672
                _ = fi88
                _ = v89
            }
        default:
            if false {
                fmt.Println(v12, v2)
            }
        }
    }
    var v90 int = 637
    var v91 uint = v2
    for fi92 := 0; fi92 < 7; fi92++ {
        fmt.Println(v2, v91, fi92)
        _ = len(v1)
        _ = fi92
    }
    _ = len(v1)
    v2 -= (v91 ^ (v91 / (((v2 / ((v91) | 1))) | 1)))
    for fi93 := 0; fi93 < 5; fi93++ {
        if ((!true) || (2147483648 >= 0)) {
            if ("" < "a") {
                switch "foo bar" {
                case "hello":
                    v12 -= v12
                }
                switch 32767 {
                case -725:
                    _ = len(v3)
                case -9223372036854775808:
                    if (3.14159 != 1e10) {
                        v3 = append(v3, 4294967295)
                        if true {
                            switch "a" {
                            case "\n\t":
                                v90 = (v90 & 32768)
                            }
                            for fi94 := 0; fi94 < 14; fi94++ {
                                var v95 int8 = 42
                                _ = fi94
                                _ = v95
                            }
                        } else {
                            if (18446744073709551615 <= 0) {
                                for fi96 := 0; fi96 < 2; fi96++ {
                                    var v97 uint64 = 994
                                    _ = fi96
                                    _ = v97
                                }
                                _ = len(v3)
                            } else {
                                switch (((("hello" + "hello") + "a") + "foo bar") + "a") {
                                case "":
                                    switch 9223372036854775807 {
                                    case -633:
                                        var v98 uint = 0
                                        _ = v98
                                    case 4294967295:
                                        for fi99 := 0; fi99 < 18; fi99++ {
                                            v1[fi99] = (fi99 >> (((^(v91))) & 63))
                                            var v100 float32 = float32(1e10)
                                            _ = fi99
                                            _ = v100
                                        }
                                    case -906:
                                        switch fi93 {
                                        case -2147483648:
                                            v101 := GenS1{F0: 475, F1: ("" + ((("\n\t" + "foo bar") + "a") + (("hello" + "") + ((("foo bar" + "a") + ("a" + (((("a" + "foo bar") + "foo bar") + "hello") + ""))) + "")))), F2: -1}
                                            _ = v101
                                        case 32768:
                                            var v102 int64 = -924
                                            _ = v102
                                        case -1:
                                            for fi103 := 0; fi103 < 3; fi103++ {
                                                if true {
                                                    if false {
                                                        if ((v91 / (((v2 & v91)) | 1)) >= v2) {
                                                            v12 += v12
                                                            v104 := GenS11{F0: 65535, F1: ((("" + ("hello" + ("hello" + "foo bar"))) + ("" + "a")) + ("" + "")), F2: 42}
                                                            var v105 uint32 = 2147483647
                                                            _ = v104
                                                            _ = v105
                                                        }
                                                        if false {
                                                            if (!false) {
                                                                v12 += (v12 - v12)
                                                                v12 += v12
                                                            } else {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                v2 = (v2 | (v91 - (v2 % ((42) | 1))))
                                                            }
                                                        }
                                                        switch (v90 | (fi103 / ((-955) | 1))) {
                                                        case -854:
                                                            if true {
                                                                var v106 int16 = 1000
                                                                switch (-(v90)) {
                                                                case 32767:
                                                                    v2 += 255
                                                                }
                                                                _ = v106
                                                            } else {
                                                                v107 := GenS7{F0: ("" + "a"), F1: genFn0(1, -182)}
                                                                v108 := GenS12{F0: 42, F1: 682}
                                                                _ = v107
                                                                _ = v108
                                                            }
                                                        case 42:
                                                            if false {
                                                                v12 = (v12 >> ((507) & 63))
                                                            } else {
                                                                v2 = (^(v91))
                                                                v109 := GenS13{F0: uint32(v2), F1: uint(v12), F2: -1}
                                                                if len(v3) > 0 {
                                                                    v3[0] = 4294967295
                                                                }
                                                                _ = v109
                                                            }
                                                        }
                                                    } else {
                                                        fmt.Println(v91, fi93, v2)
                                                    }
                                                    if (9223372036854775807 >= int64(v91)) {
                                                        v91 -= 65535
                                                    }
                                                    if (295 <= 4294967295) {
                                                        v110 := GenS14{F0: 371, F1: 894, F2: 0}
                                                        v111 := GenS13{F0: 42, F1: v2, F2: -1}
                                                        _ = len(v1)
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                } else {
                                                    v90 += (v90 >> ((135) & 63))
                                                    var v112 uint8 = 59
                                                    _ = v112
                                                }
                                                if true {
                                                    var v113 bool = (0 != 32767)
                                                    switch "" {
                                                    case "":
                                                        fmt.Println(v91)
                                                    case "a":
                                                        switch (fi103 >> ((v2) & 63)) {
                                                        case -2147483648:
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v114 float64 = 3.14159
                                                                var v115 string = ""
                                                                var v116 uint32 = 195
                                                                _ = v114
                                                                _ = v115
                                                                _ = v116
                                                            case "\n\t":
                                                                if (65535 >= 0) {
                                                                    var v117 int16 = -248
                                                                    _ = v117
                                                                } else {
                                                                    v91 += v91
                                                                }
                                                            case "":
                                                                if (2 >= 367) {
                                                                    if v113 {
                                                                        v91 = v91
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 1
                                                                        }
                                                                        delete(v1, -737)
                                                                    }
                                                                    var v118 uint32 = 0
                                                                    v119 := v1[(fi93 & v90)]
                                                                    if v113 {
                                                                        v3 = append(v3, (v118 >> ((v2) & 63)))
                                                                    }
                                                                    v91 = 2
                                                                    v118 -= (v118 - (v118 | 130))
                                                                    _ = v118
                                                                    _ = v119
                                                                } else {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    v120 := GenS15{F0: -2147483648, F1: 1, F2: (-(fi93))}
                                                                    _ = v120
                                                                }
                                                            default:
                                                                if (true || false) {
                                                                    v2 += v2
                                                                }
                                                            }
                                                        case -842:
                                                            v113 = v113
                                                        }
                                                    case "hello":
                                                        v1[0] = 76
                                                    }
                                                    var v121 int16 = 0
                                                    _ = v113
                                                    _ = v121
                                                }
                                                v2 -= v2
                                                _ = fi103
                                            }
                                        }
                                    }
                                case "hello":
                                    for fi122 := 0; fi122 < 17; fi122++ {
                                        var v123 uint = v91
                                        var v124 uint16 = (v12 / ((v12) | 1))
                                        _ = fi122
                                        _ = v123
                                        _ = v124
                                    }
                                case "foo bar":
                                    v91 = v2
                                default:
                                    v125 := GenS14{F0: 42, F1: 1, F2: 1}
                                    _ = v125
                                }
                                var v126 uint32 = 484
                                _ = v126
                            }
                        }
                    }
                case -788:
                    var v127 uint = (^(v91))
                    _ = v127
                default:
                    switch "\n\t" {
                    case "foo bar":
                        v128 := GenS4{F0: 1, F1: 1, F2: v12}
                        _ = v128
                    case "hello":
                        v129 := GenS8{F0: 2147483647, F1: false, F2: 127}
                        _ = v129
                    }
                }
                if false {
                    switch "" {
                    case "\n\t":
                        v12 = v12
                    case "hello":
                        if (true || true) {
                            v130 := GenS16{F0: 18446744073709551615}
                            for fi131 := 0; fi131 < 1; fi131++ {
                                v90 = 2147483647
                                _ = fi131
                            }
                            _ = v130
                        } else {
                            for fi132 := 0; fi132 < 11; fi132++ {
                                if false {
                                    fmt.Println(v91, v2, fi93)
                                    switch "\n\t" {
                                    case "hello":
                                        v90 -= v90
                                    case "a":
                                        if (!true) {
                                            switch "" {
                                            case "\n\t":
                                                var v133 int64 = 0
                                                _ = v133
                                            }
                                            v134 := GenS11{F0: (^(v91)), F1: "\n\t", F2: 2147483648}
                                            switch 4294967295 {
                                            case 9223372036854775807:
                                                var v135 string = ("a" + (("foo bar" + "a") + ""))
                                                if true {
                                                    if len(v3) > 0 {
                                                        v3[0] = 802
                                                    }
                                                } else {
                                                    v3 = append(v3, 63)
                                                }
                                                _ = v135
                                            case 0:
                                                if (true && (42 >= 0)) {
                                                    v12 -= v12
                                                    v91 = (v2 >> (((^(v2))) & 63))
                                                    var v136 float32 = ((((float32(0.1) - (0.0 / (3.14159 / float32(-2.25)))) - (0.0 / (3.14159 / (1.5 + 0.1)))) + float32(((1e10 / float64(1.5)) + -2.25))) + ((float32(1e10) - 0.0) - 0.1))
                                                    v137 := GenS13{F0: 65535, F1: v91, F2: 0}
                                                    if len(v3) > 0 {
                                                        v3[0] = 0
                                                    }
                                                    var v138 int16 = 32767
                                                    v90 = (^(fi132))
                                                    var v139 uint64 = 853
                                                    v140 := GenS17{F0: 1000}
                                                    v138 = -1
                                                    _ = v136
                                                    _ = v137
                                                    _ = v138
                                                    _ = v139
                                                    _ = v140
                                                } else {
                                                    v141 := GenS13{F0: 780, F1: (v2 * (v2 ^ 0)), F2: 718}
                                                    if true {
                                                        v142 := GenS18{F0: 9223372036854775807}
                                                        _ = v142
                                                    }
                                                    _ = v141
                                                }
                                            case 32768:
                                                v91 -= v2
                                            default:
                                                v90 += 834
                                            }
                                            _ = v134
                                        }
                                    }
                                }
                                var v143 float64 = 0.1
                                v91 = (v91 * v91)
                                _ = fi132
                                _ = v143
                            }
                            for fi144 := 0; fi144 < 10; fi144++ {
                                fmt.Println(fi144, fi93, v90)
                                _ = fi144
                            }
                            fmt.Println(v12)
                        }
                    case "a":
                        if true {
                            var v145 string = (("a" + "hello") + "foo bar")
                            fmt.Println(v90, v145)
                            v90 -= fi93
                            _ = v145
                        }
                    }
                    fmt.Println(v90, fi93)
                    var v146 uint = (v91 % (((v2 | v91)) | 1))
                    _ = v146
                } else {
                    fmt.Println(fi93, v2)
                }
            }
            var v147 int16 = -295
            _ = v147
        } else {
            switch ("foo bar" + "a") {
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "":
                v148 := GenS19{F0: false, F1: v91, F2: "a"}
                _ = v148
            }
        }
        for fi149 := 0; fi149 < 15; fi149++ {
            var v150 int16 = 66
            _ = fi149
            _ = v150
        }
        for fi151 := 0; fi151 < 13; fi151++ {
            var v152 uint16 = v12
            if true {
                v153 := v1[fi93]
                _ = v153
            }
            switch (("\n\t" + "hello") + "hello") {
            case "foo bar":
                v91 = 0
            case "hello":
                v152 = (v152 << (((v91 | (v91 + (v91 >> ((947) & 63))))) & 63))
            default:
                fmt.Println(fi93, v90, v90)
            }
            v154 := v1[fi151]
            if (!true) {
                var v155 uint16 = 32767
                v90 -= fi151
                _ = v155
            }
            v156 := v1[int(v2)]
            _ = fi151
            _ = v152
            _ = v154
            _ = v156
        }
        _ = fi93
    }
    if len(v3) > 0 {
        v3[0] = 7
    }
    for fi157 := 0; fi157 < 8; fi157++ {
        fmt.Println(v12)
        if len(v3) > 0 {
            v3[0] = 420
        }
        if len(v3) > 0 {
            v3[0] = 2
        }
        _ = fi157
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v90
    _ = v91
}
