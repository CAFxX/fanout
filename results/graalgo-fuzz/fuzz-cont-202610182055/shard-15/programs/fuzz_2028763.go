package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int32; F1 uint64; F2 int64 }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS11 struct { F0 int32; F1 bool; F2 int16 }
type GenS12 struct { F0 int32; F1 uint64; F2 int64 }
type GenS13 struct { F0 int64; F1 int64; F2 int8 }
type GenS14 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int16; F1 int64; F2 int64 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 int32; F1 bool; F2 int64 }
type GenS19 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 uint32; F1 string }
type GenS22 struct { F0 int16; F1 int64; F2 int16 }
type GenS23 struct { F0 int16; F1 uint32 }
type GenS24 struct { F0 int8; F1 int32; F2 uint32 }
type GenS25 struct { F0 uint8 }
type GenS26 struct { F0 uint32; F1 uint8; F2 int }
type GenS27 struct { F0 int64; F1 uint64 }
type GenS28 struct { F0 uint; F1 int; F2 int16 }
type GenS29 struct { F0 int64; F1 bool }
type GenS30 struct { F0 int16; F1 uint16; F2 int8 }
type GenS31 struct { F0 uint8 }
type GenS32 struct { F0 int32; F1 int8 }
type GenS33 struct { F0 uint32; F1 int8 }
type GenS34 struct { F0 bool; F1 int32 }

func main() {
    v1 := []uint32{2147483648}
    var v2 int32 = -54
    v3 := make(map[int]string)
    var v4 uint8 = 42
    v5 := GenS1{F0: 831}
    if false {
        if (false || (0.1 >= ((1e10 / 1.5) - 1e10))) {
            v2 += (v2 | (v2 >> ((907) & 63)))
            v6 := GenS2{F0: -1}
            _ = v6
        } else {
            if false {
                if (3.14159 >= ((1.5 + 1.5) - 3.14159)) {
                    v4 = v4
                    fmt.Println(v2)
                } else {
                    v7 := GenS2{F0: 1}
                    if ((2 > 42) || (2 <= 2)) {
                        switch "hello" {
                        case "\n\t":
                            v2 = (v2 | (-(v2)))
                        default:
                            switch "\n\t" {
                            case "foo bar":
                                var v8 uint16 = 457
                                _ = v8
                            }
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                        v4 = (v4 << ((88) & 63))
                    }
                    if true {
                        var v9 string = "foo bar"
                        for fi10 := 0; fi10 < 14; fi10++ {
                            switch (fi10 % (((fi10 / (((fi10 % ((fi10) | 1))) | 1))) | 1)) {
                            case 32768:
                                var v11 string = v9
                                _ = v11
                            case -1:
                                v12 := GenS2{F0: 0}
                                _ = v12
                            case 889:
                                v4 += v4
                            default:
                                if true {
                                    for fi13 := 0; fi13 < 5; fi13++ {
                                        var v14 int16 = 250
                                        _ = fi13
                                        _ = v14
                                    }
                                    v3[(fi10 & int(v4))] = v9
                                }
                            }
                            _ = fi10
                        }
                        _ = v9
                    }
                    _ = v7
                }
                if len(v1) > 0 {
                    v1[0] = 67
                }
            } else {
                switch (("\n\t" + "hello") + "\n\t") {
                case "a":
                    switch "\n\t" {
                    case "foo bar":
                        v2 = (v2 / (((v2 >> ((42) & 63))) | 1))
                    case "a":
                        switch -9223372036854775808 {
                        case 122:
                            v15 := GenS2{F0: 0}
                            _ = v15
                        case 0:
                            v1 = append(v1, 1)
                        default:
                            var v16 int32 = v2
                            _ = v16
                        }
                    case "":
                        v4 -= (v4 % ((v4) | 1))
                    default:
                        for fi17 := 0; fi17 < 19; fi17++ {
                            v18 := GenS3{F0: "a"}
                            if ((fi17 >> ((4294967295) & 63)) < fi17) {
                                var v19 float32 = 1.5
                                var v20 uint8 = 44
                                _ = v19
                                _ = v20
                            }
                            v2 += 87
                            _ = fi17
                            _ = v18
                        }
                    }
                case "foo bar":
                    fmt.Println(v4)
                }
            }
            for fi21 := 0; fi21 < 3; fi21++ {
                for fi22 := 0; fi22 < 8; fi22++ {
                    v4 = (v4 / (((^(v4))) | 1))
                    switch ("foo bar" + "\n\t") {
                    case "hello":
                        switch fi22 {
                        case 4294967295:
                            fmt.Println(v4)
                        case -2147483648:
                            v2 += -1
                        case 247:
                            switch "a" {
                            case "":
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(v4, fi21)
                                case "hello":
                                    var v23 int = 0
                                    _ = v23
                                }
                            case "\n\t":
                                var v24 float64 = -2.25
                                _ = v24
                            case "hello":
                                switch "foo bar" {
                                case "foo bar":
                                    if (42 == 42) {
                                        v25 := GenS4{F0: (^(v4)), F1: 2}
                                        v2 -= (v2 + v2)
                                        _ = v25
                                    } else {
                                        switch (fi21 & fi22) {
                                        case 127:
                                            delete(v3, fi21)
                                        case 9223372036854775807:
                                            v2 += v2
                                        case 128:
                                            var v26 uint64 = 407
                                            _ = v26
                                        default:
                                            var v27 int64 = 9223372036854775807
                                            _ = v27
                                        }
                                        delete(v3, (fi21 % ((fi22) | 1)))
                                        v2 += v2
                                        v1 = append(v1, 777)
                                        v4 = 255
                                    }
                                case "a":
                                    v4 += v4
                                case "":
                                    if ((!((541 < 32767) && ((!false) || true))) && (false && (!true))) {
                                        v2 = (v2 >> ((4294967295) & 63))
                                    } else {
                                        v4 = v4
                                        var v28 int = -688
                                        _ = v28
                                    }
                                default:
                                    switch (fi21 & (fi22 & fi21)) {
                                    case 0:
                                        switch 32768 {
                                        case 42:
                                            if true {
                                                switch ("foo bar" + (("\n\t" + "a") + ("hello" + ""))) {
                                                case "":
                                                    if false {
                                                        var v29 uint64 = 18446744073709551615
                                                        v30 := GenS5{F0: -1, F1: 0, F2: -9223372036854775808}
                                                        _ = v29
                                                        _ = v30
                                                    } else {
                                                        if (true || (361 >= 0)) {
                                                            if false {
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v31 bool = (!true)
                                                                    _ = v31
                                                                }
                                                                var v32 string = (("foo bar" + ("" + (("\n\t" + "a") + "\n\t"))) + ((("\n\t" + "hello") + ("" + "hello")) + ((("foo bar" + ("a" + "a")) + ("a" + ("foo bar" + "hello"))) + ("foo bar" + ""))))
                                                                v2 = int32(v4)
                                                                var v33 int64 = 0
                                                                _ = v32
                                                                _ = v33
                                                            } else {
                                                                switch 1 {
                                                                case 599:
                                                                    var v34 int32 = (^(v2))
                                                                    _ = v34
                                                                case -1:
                                                                    var v35 string = "hello"
                                                                    _ = v35
                                                                case 1:
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        if false {
                                                                            if true {
                                                                                v4 = v4
                                                                                v4 += (^(v4))
                                                                                switch ("foo bar" + "") {
                                                                                case "a":
                                                                                    var v36 int32 = (v2 - (v2 - (v2 | (-(v2)))))
                                                                                    _ = v36
                                                                                default:
                                                                                    switch fi22 {
                                                                                    case 127:
                                                                                        v4 += v4
                                                                                    case 1:
                                                                                        if (!(false && false)) {
                                                                                            fmt.Println(fi22, v4)
                                                                                            var v37 uint32 = 1
                                                                                            _ = v37
                                                                                        }
                                                                                    case -278:
                                                                                        var v38 bool = (0.0 < 1e10)
                                                                                        _ = v38
                                                                                    default:
                                                                                        v4 = v4
                                                                                    }
                                                                                }
                                                                            }
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                v2 -= 2147483647
                                                                            }
                                                                            if true {
                                                                                delete(v3, (^(fi21)))
                                                                                fmt.Println(v4, v2, fi22)
                                                                            } else {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 2147483647
                                                                                }
                                                                                v2 -= -867
                                                                                var v39 uint64 = 637
                                                                                _ = v39
                                                                            }
                                                                        }
                                                                    default:
                                                                        fmt.Println(v2, fi21, fi22)
                                                                        var v40 uint8 = v4
                                                                        _ = v40
                                                                    }
                                                                }
                                                                var v41 string = ("hello" + (("a" + ("foo bar" + "")) + (("foo bar" + ("a" + "\n\t")) + "foo bar")))
                                                                _ = v41
                                                            }
                                                            v2 = v2
                                                            switch -877 {
                                                            case 32768:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                            var v42 int16 = 993
                                                            _ = v42
                                                        } else {
                                                            switch ("" + "hello") {
                                                            case "foo bar":
                                                                if (("" + ("" + "\n\t")) == "foo bar") {
                                                                    _ = len(v3)
                                                                    if true {
                                                                        v2 -= (v2 + (v2 ^ (v2 % ((42) | 1))))
                                                                        _ = len(v3)
                                                                        var v43 uint16 = uint16(v4)
                                                                        _ = v43
                                                                    } else {
                                                                        v44 := GenS1{F0: 0}
                                                                        if ((100 < 0) && false) {
                                                                            v4 += v4
                                                                            v4 -= 139
                                                                            var v45 int8 = -128
                                                                            _ = v45
                                                                        }
                                                                        _ = v44
                                                                    }
                                                                    var v46 int64 = 0
                                                                    _ = v46
                                                                } else {
                                                                    v47 := GenS6{F0: -2147483648, F1: 0}
                                                                    _ = v47
                                                                }
                                                            }
                                                        }
                                                        if (0 <= 1000) {
                                                            switch "a" {
                                                            case "hello":
                                                                if false {
                                                                    if ("a" != ("\n\t" + "\n\t")) {
                                                                        var v48 uint8 = 128
                                                                        if true {
                                                                            v48 = 255
                                                                            fmt.Println(fi22)
                                                                            if true {
                                                                                if ("\n\t" <= "") {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 42
                                                                                    }
                                                                                } else {
                                                                                    if true {
                                                                                        v49 := GenS7{F0: (v2 >> ((2) & 63))}
                                                                                        _ = v49
                                                                                    }
                                                                                }
                                                                                v50 := GenS8{F0: 0}
                                                                                v4 = 198
                                                                                _ = v50
                                                                            }
                                                                            v51 := GenS8{F0: 1}
                                                                            v3[-725] = "hello"
                                                                            _ = v51
                                                                        } else {
                                                                            var v52 uint64 = 0
                                                                            _ = v52
                                                                        }
                                                                        v1 = append(v1, 4)
                                                                        _ = v48
                                                                    }
                                                                }
                                                            }
                                                            var v53 int8 = 22
                                                            fmt.Println(fi21, v4)
                                                            _ = v53
                                                        }
                                                        fmt.Println(v2, v4, v2)
                                                    }
                                                case "foo bar":
                                                    var v54 bool = true
                                                    if false {
                                                        v1 = append(v1, 2147483647)
                                                        var v55 int8 = -79
                                                        _ = v55
                                                    }
                                                    if v54 {
                                                        _ = len(v3)
                                                        if ((((("hello" + "\n\t") < "foo bar") && (!v54)) && ("hello" >= "hello")) || v54) {
                                                            var v56 uint8 = v4
                                                            _ = v56
                                                        }
                                                        var v57 int32 = v2
                                                        if v54 {
                                                            v58 := GenS4{F0: v4, F1: 1}
                                                            _ = len(v3)
                                                            var v59 int32 = v57
                                                            switch "hello" {
                                                            case "a":
                                                                v60 := GenS3{F0: ("foo bar" + (("" + "") + ""))}
                                                                _ = v60
                                                            case "\n\t":
                                                                if ((42 > 1) && true) {
                                                                    v57 = (v57 % (((v57 * (^(v59)))) | 1))
                                                                    v61 := GenS6{F0: -1, F1: 18446744073709551615}
                                                                    v62 := GenS6{F0: fi21, F1: 357}
                                                                    _ = v61
                                                                    _ = v62
                                                                } else {
                                                                    if true {
                                                                        fmt.Println(fi22, fi22, fi21)
                                                                        var v63 uint = 18446744073709551615
                                                                        var v64 string = "hello"
                                                                        _ = v63
                                                                        _ = v64
                                                                    } else {
                                                                        var v65 int32 = v59
                                                                        if true {
                                                                            v66 := v3[(fi21 >> ((1) & 63))]
                                                                            v67 := GenS9{F0: 100}
                                                                            _ = v66
                                                                            _ = v67
                                                                        } else {
                                                                            var v68 string = "hello"
                                                                            var v69 uint = 0
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                        _ = v65
                                                                    }
                                                                    v70 := GenS10{F0: 65535, F1: 959, F2: 181}
                                                                    if v54 {
                                                                        if true {
                                                                            v4 -= (v4 % (((v4 - (v4 / (((v4 << ((255) & 63))) | 1)))) | 1))
                                                                        }
                                                                        fmt.Println(fi22)
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            var v71 string = ""
                                                                            _ = v71
                                                                        case "hello":
                                                                            v4 -= v4
                                                                        case "a":
                                                                            v72 := GenS7{F0: v59}
                                                                            _ = v72
                                                                        default:
                                                                            v54 = (255 >= 65535)
                                                                        }
                                                                    } else {
                                                                        v73 := GenS11{F0: (^(v2)), F1: (18446744073709551615 == 2), F2: -498}
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            fmt.Println(fi21, v57)
                                                                        case "a":
                                                                            switch ("\n\t" + ("" + (("a" + ("" + "")) + "hello"))) {
                                                                            case "hello":
                                                                                switch (("\n\t" + "hello") + "a") {
                                                                                case "":
                                                                                    switch "foo bar" {
                                                                                    case "\n\t":
                                                                                        fmt.Println(v2)
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(fi21, v59, v59)
                                                                                }
                                                                            case "":
                                                                                switch ("hello" + "a") {
                                                                                case "foo bar":
                                                                                    if v54 {
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 1
                                                                                        }
                                                                                        v59 -= v57
                                                                                    } else {
                                                                                        if true {
                                                                                            var v74 uint8 = (v4 ^ 2)
                                                                                            fmt.Println(v4, v57)
                                                                                            _ = v74
                                                                                        } else {
                                                                                            v59 = (v57 * (v2 << ((255) & 63)))
                                                                                            v3[fi21] = "foo bar"
                                                                                        }
                                                                                        var v75 bool = ((v4 | v4) >= (v4 / ((1) | 1)))
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 2147483648
                                                                                        }
                                                                                        _ = v75
                                                                                    }
                                                                                case "":
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                case "hello":
                                                                                    var v76 uint32 = 718
                                                                                    _ = v76
                                                                                }
                                                                            case "a":
                                                                                fmt.Println(v2)
                                                                            default:
                                                                                fmt.Println(v57, v57)
                                                                            }
                                                                        case "\n\t":
                                                                            delete(v3, (fi21 ^ fi21))
                                                                        }
                                                                        v57 = v57
                                                                        _ = v73
                                                                    }
                                                                    fmt.Println(v57, fi22, v54)
                                                                    if false {
                                                                        switch -1 {
                                                                        case 0:
                                                                            v1 = append(v1, 272)
                                                                        case -576:
                                                                            var v77 uint8 = v4
                                                                            _ = v77
                                                                        default:
                                                                            if ((true || (true || v54)) && false) {
                                                                                switch 127 {
                                                                                case 0:
                                                                                    fmt.Println(v4, v4)
                                                                                case -2147483648:
                                                                                    var v78 float64 = (float64(((0.1 - 3.14159) / 0.1)) - 1.5)
                                                                                    _ = v78
                                                                                case 128:
                                                                                    v79 := GenS12{F0: (v59 / (((-(v2))) | 1)), F1: 573, F2: 9223372036854775807}
                                                                                    _ = v79
                                                                                }
                                                                                v80 := GenS4{F0: v4, F1: 59}
                                                                                _ = v80
                                                                            } else {
                                                                                if false {
                                                                                    if true {
                                                                                        var v81 uint32 = 2
                                                                                        _ = v81
                                                                                    } else {
                                                                                        v1 = append(v1, 281)
                                                                                    }
                                                                                    var v82 uint32 = 364
                                                                                    switch fi22 {
                                                                                    case -2147483648:
                                                                                        v83 := GenS11{F0: v59, F1: v54, F2: 42}
                                                                                        _ = v83
                                                                                    default:
                                                                                        var v84 int = fi21
                                                                                        _ = v84
                                                                                    }
                                                                                    _ = v82
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v70
                                                                }
                                                            }
                                                            v4 = 42
                                                            var v85 int64 = -150
                                                            if ((v85 ^ (v85 + v85)) <= v85) {
                                                                var v86 int8 = 91
                                                                _ = len(v3)
                                                                delete(v3, fi21)
                                                                _ = v86
                                                            } else {
                                                                v87 := GenS5{F0: 1, F1: 148, F2: int64(fi22)}
                                                                if ((!v54) || v54) {
                                                                    fmt.Println(fi22)
                                                                    var v88 uint = uint(v4)
                                                                    v89 := v3[428]
                                                                    _ = v88
                                                                    _ = v89
                                                                } else {
                                                                    var v90 uint64 = 18446744073709551615
                                                                    v57 = 597
                                                                    _ = v90
                                                                }
                                                                var v91 int8 = -83
                                                                _ = v87
                                                                _ = v91
                                                            }
                                                            _ = v58
                                                            _ = v59
                                                            _ = v85
                                                        } else {
                                                            var v92 uint32 = uint32(v4)
                                                            v92 -= uint32(v4)
                                                            _ = v92
                                                        }
                                                        _ = v57
                                                    } else {
                                                        var v93 uint = 838
                                                        _ = v93
                                                    }
                                                    _ = v54
                                                case "\n\t":
                                                    v94 := GenS4{F0: v4, F1: 0}
                                                    _ = v94
                                                default:
                                                    if true {
                                                        v4 -= v4
                                                        if ((int(v2) >= (fi22 * fi22)) && (((float32(0.1) + -2.25) + 0.1) >= 0.0)) {
                                                            v95 := GenS13{F0: -1, F1: 0, F2: 42}
                                                            if (false || false) {
                                                                var v96 uint64 = 1
                                                                fmt.Println(fi22, fi21, v4)
                                                                _ = v96
                                                            } else {
                                                                v2 = 1
                                                                var v97 int64 = -1
                                                                var v98 int64 = v97
                                                                _ = v97
                                                                _ = v98
                                                            }
                                                            _ = v95
                                                        } else {
                                                            if false {
                                                                v4 = v4
                                                                v99 := GenS6{F0: -9223372036854775808, F1: 18446744073709551615}
                                                                v3[fi21] = ("\n\t" + "")
                                                                var v100 int = 32767
                                                                switch "a" {
                                                                case "foo bar":
                                                                    switch "" {
                                                                    case "":
                                                                        v101 := GenS14{F0: 128, F1: uint16(v4), F2: int32(v4)}
                                                                        v100 += 42
                                                                        _ = v101
                                                                    case "a":
                                                                        if (((0.0 - float64(((1e10 / (1e10 + -2.25)) * (((-2.25 + -2.25) * 0.0) * (0.1 / float32(0.1)))))) * 0.0) >= (float64(1.5) * ((0.0 + 3.14159) * 1e10))) {
                                                                            var v102 uint16 = 860
                                                                            fmt.Println(fi21)
                                                                            v103 := GenS15{F0: 100}
                                                                            var v104 uint64 = 2
                                                                            _ = v102
                                                                            _ = v103
                                                                            _ = v104
                                                                        }
                                                                        v2 = v2
                                                                        if (-1 <= -623) {
                                                                            v105 := GenS16{F0: -1, F1: -1, F2: 1}
                                                                            v100 -= v100
                                                                            v4 = (v4 / (((v4 * 42)) | 1))
                                                                            v106 := GenS17{F0: 127}
                                                                            _ = v105
                                                                            _ = v106
                                                                        } else {
                                                                            if ((v2 - v2) < v2) {
                                                                                if false {
                                                                                    v107 := GenS10{F0: 42, F1: 208, F2: -1}
                                                                                    v100 -= 9223372036854775807
                                                                                    _ = v107
                                                                                }
                                                                                v1 = append(v1, 356)
                                                                                switch 676 {
                                                                                case 32768:
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                case -677:
                                                                                    v100 -= (fi21 & v100)
                                                                                default:
                                                                                    fmt.Println(fi21)
                                                                                }
                                                                                if false {
                                                                                    var v108 int32 = (v2 - v2)
                                                                                    if (!(!false)) {
                                                                                        if (18446744073709551615 < 189) {
                                                                                            var v109 float32 = 1.5
                                                                                            fmt.Println(v100, v109)
                                                                                            switch v100 {
                                                                                            case 9223372036854775807:
                                                                                                var v110 int16 = -32768
                                                                                                _ = v110
                                                                                            case 1:
                                                                                                if false {
                                                                                                    switch (v100 | -767) {
                                                                                                    case 128:
                                                                                                        v100 -= (fi22 % (((fi22 << ((1) & 63))) | 1))
                                                                                                    case -979:
                                                                                                        v111 := GenS1{F0: 42}
                                                                                                        if (0.0 < -2.25) {
                                                                                                            switch "\n\t" {
                                                                                                            case "hello":
                                                                                                                v108 = -1
                                                                                                                fmt.Println(v108, v100)
                                                                                                                _ = len(v3)
                                                                                                            case "foo bar":
                                                                                                                var v112 uint64 = 0
                                                                                                                _ = v112
                                                                                                            case "":
                                                                                                                var v113 uint8 = (v4 - (v4 << ((65535) & 63)))
                                                                                                                _ = v113
                                                                                                            }
                                                                                                            if false {
                                                                                                                var v114 uint = 42
                                                                                                                switch ("foo bar" + "hello") {
                                                                                                                case "a":
                                                                                                                    switch (v100 << ((1) & 63)) {
                                                                                                                    case 9223372036854775807:
                                                                                                                        var v115 uint64 = 4294967295
                                                                                                                        _ = v115
                                                                                                                    case 4294967295:
                                                                                                                        switch "" {
                                                                                                                        case "":
                                                                                                                            v116 := GenS18{F0: -1, F1: (("hello" >= "foo bar") || (true && false)), F2: 9223372036854775807}
                                                                                                                            _ = v116
                                                                                                                        default:
                                                                                                                            if (65535 >= 32767) {
                                                                                                                                switch "foo bar" {
                                                                                                                                case "hello":
                                                                                                                                    var v117 string = "\n\t"
                                                                                                                                    _ = v117
                                                                                                                                }
                                                                                                                                fmt.Println(v2)
                                                                                                                                v118 := GenS1{F0: 42}
                                                                                                                                _ = v118
                                                                                                                            } else {
                                                                                                                                if len(v1) > 0 {
                                                                                                                                    v1[0] = 23
                                                                                                                                }
                                                                                                                                if (v100 != fi22) {
                                                                                                                                    v119 := v3[fi21]
                                                                                                                                    if false {
                                                                                                                                        if true {
                                                                                                                                            if (-1 >= -54) {
                                                                                                                                                if true {
                                                                                                                                                    if (fi21 > (fi21 << (((v114 * (v114 & 255))) & 63))) {
                                                                                                                                                        v120 := GenS5{F0: v2, F1: 42, F2: 1}
                                                                                                                                                        var v121 int64 = -1
                                                                                                                                                        _ = v120
                                                                                                                                                        _ = v121
                                                                                                                                                    }
                                                                                                                                                    var v122 uint = v114
                                                                                                                                                    var v123 float32 = v109
                                                                                                                                                    _ = v122
                                                                                                                                                    _ = v123
                                                                                                                                                }
                                                                                                                                                v114 += v114
                                                                                                                                            }
                                                                                                                                            fmt.Println(v108, v4, v109)
                                                                                                                                        } else {
                                                                                                                                            v124 := GenS19{F0: v4, F1: 2, F2: 402}
                                                                                                                                            switch v100 {
                                                                                                                                            case -9223372036854775808:
                                                                                                                                                fmt.Println(v119, v114)
                                                                                                                                            case -339:
                                                                                                                                                var v125 float32 = v109
                                                                                                                                                _ = v125
                                                                                                                                            }
                                                                                                                                            _ = v124
                                                                                                                                        }
                                                                                                                                        if (295 > -2147483648) {
                                                                                                                                            switch "\n\t" {
                                                                                                                                            case "foo bar":
                                                                                                                                                switch fi21 {
                                                                                                                                                case 4294967295:
                                                                                                                                                    if (2 == 250) {
                                                                                                                                                        v126 := GenS10{F0: 719, F1: 225, F2: 1}
                                                                                                                                                        if len(v1) >= 2 {
                                                                                                                                                            v1 = v1[1:len(v1)]
                                                                                                                                                        }
                                                                                                                                                        _ = v126
                                                                                                                                                    } else {
                                                                                                                                                        v127 := v3[2147483647]
                                                                                                                                                        if false {
                                                                                                                                                            var v128 uint32 = 42
                                                                                                                                                            v127 = v127
                                                                                                                                                            var v129 int = (fi22 + fi22)
                                                                                                                                                            _ = v128
                                                                                                                                                            _ = v129
                                                                                                                                                        }
                                                                                                                                                        _ = v127
                                                                                                                                                    }
                                                                                                                                                case -9223372036854775808:
                                                                                                                                                    delete(v3, v100)
                                                                                                                                                case 2147483647:
                                                                                                                                                    if false {
                                                                                                                                                        var v130 int = fi22
                                                                                                                                                        v131 := GenS9{F0: -1}
                                                                                                                                                        var v132 int64 = -378
                                                                                                                                                        if true {
                                                                                                                                                            var v133 int32 = -2147483648
                                                                                                                                                            var v134 int32 = (v133 - v108)
                                                                                                                                                            v135 := GenS10{F0: 324, F1: 473, F2: (v132 | v132)}
                                                                                                                                                            _ = v133
                                                                                                                                                            _ = v134
                                                                                                                                                            _ = v135
                                                                                                                                                        }
                                                                                                                                                        if len(v1) > 0 {
                                                                                                                                                            v1[0] = 0
                                                                                                                                                        }
                                                                                                                                                        _ = v130
                                                                                                                                                        _ = v131
                                                                                                                                                        _ = v132
                                                                                                                                                    } else {
                                                                                                                                                        v136 := GenS20{F0: 321}
                                                                                                                                                        v114 -= (v114 >> (((v114 / (((v114 * 255)) | 1))) & 63))
                                                                                                                                                        v137 := GenS3{F0: v119}
                                                                                                                                                        _ = v136
                                                                                                                                                        _ = v137
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            case "hello":
                                                                                                                                                var v138 int8 = 1
                                                                                                                                                _ = v138
                                                                                                                                            case "a":
                                                                                                                                                v2 -= 1
                                                                                                                                                if false {
                                                                                                                                                    var v139 bool = (1.5 == v109)
                                                                                                                                                    var v140 int64 = 927
                                                                                                                                                    _ = v139
                                                                                                                                                    _ = v140
                                                                                                                                                } else {
                                                                                                                                                    v119 = v119
                                                                                                                                                    delete(v3, -135)
                                                                                                                                                    fmt.Println(v2, fi21)
                                                                                                                                                }
                                                                                                                                                v141 := GenS1{F0: uint64(v4)}
                                                                                                                                                _ = v141
                                                                                                                                            default:
                                                                                                                                                v108 = (v2 & v108)
                                                                                                                                            }
                                                                                                                                            v1 = append(v1, 283)
                                                                                                                                        }
                                                                                                                                        var v142 uint = v114
                                                                                                                                        _ = v142
                                                                                                                                    }
                                                                                                                                    _ = v119
                                                                                                                                } else {
                                                                                                                                    switch ("a" + ("foo bar" + "a")) {
                                                                                                                                    case "a":
                                                                                                                                        fmt.Println(v108, v100, fi21)
                                                                                                                                        v3[v100] = (("a" + "foo bar") + "")
                                                                                                                                    case "\n\t":
                                                                                                                                        v143 := GenS21{F0: 2147483648, F1: (("" + "hello") + "hello")}
                                                                                                                                        _ = v143
                                                                                                                                    case "":
                                                                                                                                        v109 -= 0.1
                                                                                                                                    default:
                                                                                                                                        v100 = fi21
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                var v144 int16 = 1
                                                                                                                                _ = v144
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case 32767:
                                                                                                                        v145 := GenS22{F0: 1000, F1: 42, F2: 1000}
                                                                                                                        _ = v145
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    fmt.Println(v109, v109)
                                                                                                                }
                                                                                                                _ = v114
                                                                                                            }
                                                                                                            var v146 int32 = v2
                                                                                                            _ = v146
                                                                                                        } else {
                                                                                                            fmt.Println(v108)
                                                                                                            delete(v3, 2147483647)
                                                                                                            v100 += v100
                                                                                                        }
                                                                                                        v100 -= fi21
                                                                                                        _ = v111
                                                                                                    default:
                                                                                                        v2 = (^(v2))
                                                                                                    }
                                                                                                } else {
                                                                                                    fmt.Println(v100)
                                                                                                }
                                                                                                fmt.Println(v100)
                                                                                            default:
                                                                                                var v147 string = ("a" + "")
                                                                                                _ = v147
                                                                                            }
                                                                                            _ = v109
                                                                                        } else {
                                                                                            v1 = append(v1, 154)
                                                                                            v4 -= (^(v4))
                                                                                            var v148 uint32 = 202
                                                                                            var v149 int = fi22
                                                                                            _ = v148
                                                                                            _ = v149
                                                                                        }
                                                                                        var v150 bool = false
                                                                                        v151 := GenS15{F0: 42}
                                                                                        _ = v150
                                                                                        _ = v151
                                                                                    }
                                                                                    _ = v108
                                                                                } else {
                                                                                    fmt.Println(v2, fi22, v4)
                                                                                    var v152 uint16 = 42
                                                                                    v1 = append(v1, 65535)
                                                                                    _ = v152
                                                                                }
                                                                                if false {
                                                                                    if (!(false || (false && (!((-1 <= 127) && (((784 == 42) && (false && (42 < 0))) && true)))))) {
                                                                                        if (!false) {
                                                                                            switch (fi21 - 128) {
                                                                                            case 4294967295:
                                                                                                var v153 uint32 = 161
                                                                                                _ = v153
                                                                                            }
                                                                                            v154 := v3[fi22]
                                                                                            fmt.Println(v154, fi22)
                                                                                            _ = v154
                                                                                        } else {
                                                                                            if false {
                                                                                                fmt.Println(fi22)
                                                                                                v155 := GenS5{F0: (v2 + (v2 % ((v2) | 1))), F1: 120, F2: 9223372036854775807}
                                                                                                _ = v155
                                                                                            } else {
                                                                                                if (((1.5 < (3.14159 / 1.5)) && (!(((1.5 == 1.5) || true) && (((v4 ^ v4) >= (v4 | (v4 - 42))) || false)))) || ((v2 ^ v2) <= (^(v2)))) {
                                                                                                    v2 -= v2
                                                                                                    var v156 uint64 = 42
                                                                                                    _ = v156
                                                                                                } else {
                                                                                                    switch "\n\t" {
                                                                                                    case "":
                                                                                                        fmt.Println(fi22, fi22)
                                                                                                    case "foo bar":
                                                                                                        switch (v100 * 9223372036854775807) {
                                                                                                        case 1:
                                                                                                            if (!(true && false)) {
                                                                                                                fmt.Println(v2)
                                                                                                                var v157 uint8 = (v4 % (((v4 / ((v4) | 1))) | 1))
                                                                                                                _ = v157
                                                                                                            } else {
                                                                                                                v158 := v3[(^(v100))]
                                                                                                                _ = v158
                                                                                                            }
                                                                                                            var v159 int16 = 616
                                                                                                            _ = v159
                                                                                                        case 32767:
                                                                                                            var v160 float64 = 3.14159
                                                                                                            _ = v160
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        var v161 string = (("a" + "foo bar") + "hello")
                                                                                                        _ = v161
                                                                                                    default:
                                                                                                        var v162 uint = 0
                                                                                                        _ = v162
                                                                                                    }
                                                                                                    v163 := GenS23{F0: 32767, F1: 174}
                                                                                                    fmt.Println(fi22, v4, v100)
                                                                                                    _ = v163
                                                                                                }
                                                                                                var v164 int32 = 42
                                                                                                _ = v164
                                                                                            }
                                                                                            fmt.Println(v100, v4, v100)
                                                                                            var v165 float64 = 0.1
                                                                                            _ = v165
                                                                                        }
                                                                                        if (!true) {
                                                                                            switch (("\n\t" + "") + "a") {
                                                                                            case "":
                                                                                                v166 := GenS24{F0: -1, F1: (v2 << ((0) & 63)), F2: 535}
                                                                                                _ = v166
                                                                                            case "hello":
                                                                                                switch "hello" {
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v100, v2, v4)
                                                                                                }
                                                                                            case "a":
                                                                                                switch fi21 {
                                                                                                case 0:
                                                                                                    v167 := GenS25{F0: v4}
                                                                                                    _ = v167
                                                                                                case 4294967295:
                                                                                                    v168 := GenS24{F0: 1, F1: (v2 / ((v2) | 1)), F2: 2147483648}
                                                                                                    _ = v168
                                                                                                default:
                                                                                                    var v169 bool = false
                                                                                                    _ = v169
                                                                                                }
                                                                                            default:
                                                                                                if (0.1 < ((0.1 - 0.1) + 0.1)) {
                                                                                                    var v170 int8 = 1
                                                                                                    var v171 int64 = 0
                                                                                                    v4 += v4
                                                                                                    _ = v170
                                                                                                    _ = v171
                                                                                                }
                                                                                            }
                                                                                            v100 -= (v100 * 127)
                                                                                            v172 := GenS22{F0: 1, F1: 1, F2: 1}
                                                                                            var v173 int32 = (^(v2))
                                                                                            _ = v172
                                                                                            _ = v173
                                                                                        } else {
                                                                                            v174 := GenS26{F0: 4294967295, F1: (v4 ^ (v4 >> ((uint(v4)) & 63))), F2: 589}
                                                                                            v175 := GenS22{F0: 1000, F1: -1, F2: -714}
                                                                                            switch ("" + ("hello" + ("" + ""))) {
                                                                                            case "\n\t":
                                                                                                _ = len(v1)
                                                                                            default:
                                                                                                _ = len(v3)
                                                                                            }
                                                                                            _ = v174
                                                                                            _ = v175
                                                                                        }
                                                                                        if (32768 == fi21) {
                                                                                            v1 = append(v1, 0)
                                                                                        } else {
                                                                                            v4 -= (^(v4))
                                                                                            delete(v3, fi21)
                                                                                            var v176 bool = false
                                                                                            fmt.Println(v4, v100, fi22)
                                                                                            fmt.Println(fi22, fi21)
                                                                                            if (42 == 100) {
                                                                                                v176 = v176
                                                                                            } else {
                                                                                                fmt.Println(v176, fi21, v176)
                                                                                                v177 := GenS21{F0: 281, F1: ("foo bar" + ("\n\t" + "foo bar"))}
                                                                                                if v176 {
                                                                                                    delete(v3, v100)
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                } else {
                                                                                                    switch 43 {
                                                                                                    case -604:
                                                                                                        if (2 >= 2147483648) {
                                                                                                            v176 = ((1e10 + 0.0) == (1.5 + float32((-2.25 * 1.5))))
                                                                                                            _ = len(v1)
                                                                                                            v178 := GenS15{F0: -63}
                                                                                                            _ = v178
                                                                                                        }
                                                                                                    case 0:
                                                                                                        delete(v3, 127)
                                                                                                    case 99:
                                                                                                        var v179 float64 = float64(1.5)
                                                                                                        _ = v179
                                                                                                    default:
                                                                                                        var v180 uint = 136
                                                                                                        _ = v180
                                                                                                    }
                                                                                                    var v181 int32 = 2147483647
                                                                                                    _ = v181
                                                                                                }
                                                                                                _ = v177
                                                                                            }
                                                                                            if true {
                                                                                                switch "hello" {
                                                                                                case "foo bar":
                                                                                                    if true {
                                                                                                        delete(v3, (fi21 + v100))
                                                                                                    }
                                                                                                default:
                                                                                                    fmt.Println(v4, fi22, fi21)
                                                                                                }
                                                                                                v182 := GenS27{F0: -819, F1: 1}
                                                                                                v2 -= -290
                                                                                                _ = v182
                                                                                            } else {
                                                                                                v4 += 4
                                                                                            }
                                                                                            _ = v176
                                                                                        }
                                                                                    } else {
                                                                                        var v183 int = fi22
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        switch (^(v183)) {
                                                                                        case 2147483647:
                                                                                            var v184 uint32 = 4294967295
                                                                                            _ = v184
                                                                                        case 532:
                                                                                            v183 -= v100
                                                                                        default:
                                                                                            switch (-(v183)) {
                                                                                            case 128:
                                                                                                var v185 string = "hello"
                                                                                                _ = v185
                                                                                            case 330:
                                                                                                switch -790 {
                                                                                                case 42:
                                                                                                    if (((1.5 + -2.25) * 1e10) == ((float32(3.14159) - 0.1) - 1e10)) {
                                                                                                        if (false || false) {
                                                                                                            var v186 uint32 = 65535
                                                                                                            switch v183 {
                                                                                                            case -2147483648:
                                                                                                                v187 := GenS25{F0: v4}
                                                                                                                _ = v187
                                                                                                            default:
                                                                                                                var v188 uint64 = 726
                                                                                                                _ = v188
                                                                                                            }
                                                                                                            if ((1.5 + float32(0.0)) > 1.5) {
                                                                                                                switch fi22 {
                                                                                                                case 32767:
                                                                                                                    if (!false) {
                                                                                                                        v189 := GenS28{F0: 815, F1: -9223372036854775808, F2: 271}
                                                                                                                        var v190 uint64 = 2
                                                                                                                        v190 = 2
                                                                                                                        _ = v189
                                                                                                                        _ = v190
                                                                                                                    } else {
                                                                                                                        var v191 int16 = -32768
                                                                                                                        _ = v191
                                                                                                                    }
                                                                                                                case -363:
                                                                                                                    switch "" {
                                                                                                                    case "\n\t":
                                                                                                                        v2 -= v2
                                                                                                                    case "":
                                                                                                                        if (1 > 9223372036854775807) {
                                                                                                                            v183 = 335
                                                                                                                        } else {
                                                                                                                            v186 -= v186
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v192 uint8 = 2
                                                                                                                        _ = v192
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if false {
                                                                                                                        v193 := GenS7{F0: v2}
                                                                                                                        if (42 >= 4294967295) {
                                                                                                                            v194 := GenS9{F0: -1}
                                                                                                                            var v195 uint16 = 2
                                                                                                                            switch v100 {
                                                                                                                            case -494:
                                                                                                                                var v196 uint16 = (v195 & v195)
                                                                                                                                _ = v196
                                                                                                                            }
                                                                                                                            _ = v194
                                                                                                                            _ = v195
                                                                                                                        } else {
                                                                                                                            v4 -= (v4 << ((18446744073709551615) & 63))
                                                                                                                            v197 := v3[int(v2)]
                                                                                                                            switch 92 {
                                                                                                                            case -2147483648:
                                                                                                                                var v198 bool = (!false)
                                                                                                                                _ = v198
                                                                                                                            default:
                                                                                                                                switch v197 {
                                                                                                                                case "foo bar":
                                                                                                                                    if len(v1) > 0 {
                                                                                                                                        v1[0] = (v186 << ((0) & 63))
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _ = v197
                                                                                                                        }
                                                                                                                        _ = v193
                                                                                                                    } else {
                                                                                                                        fmt.Println(fi22, v100, fi22)
                                                                                                                        v2 += (v2 & v2)
                                                                                                                        if (false && (-1 < -9223372036854775808)) {
                                                                                                                            v3[(fi22 << ((494) & 63))] = "hello"
                                                                                                                            v2 -= v2
                                                                                                                            var v199 int8 = 127
                                                                                                                            v200 := GenS14{F0: (^(v4)), F1: 32767, F2: (v2 + 610)}
                                                                                                                            if (0 <= -903) {
                                                                                                                                var v201 int64 = -9223372036854775808
                                                                                                                                if false {
                                                                                                                                    v100 -= fi21
                                                                                                                                } else {
                                                                                                                                    var v202 uint8 = 0
                                                                                                                                    var v203 int16 = 1000
                                                                                                                                    if true {
                                                                                                                                        if true {
                                                                                                                                            if false {
                                                                                                                                                _ = len(v1)
                                                                                                                                            } else {
                                                                                                                                                fmt.Println(v202, fi21, fi22)
                                                                                                                                                if len(v1) >= 2 {
                                                                                                                                                    v1 = v1[1:len(v1)]
                                                                                                                                                }
                                                                                                                                                var v204 float32 = 3.14159
                                                                                                                                                _ = v204
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            switch fi21 {
                                                                                                                                            case -536:
                                                                                                                                                var v205 int32 = v2
                                                                                                                                                _ = v205
                                                                                                                                            default:
                                                                                                                                                fmt.Println(v186)
                                                                                                                                            }
                                                                                                                                            v203 += v203
                                                                                                                                            v202 += 127
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v202
                                                                                                                                    _ = v203
                                                                                                                                }
                                                                                                                                _ = v201
                                                                                                                            }
                                                                                                                            delete(v3, fi21)
                                                                                                                            _ = v199
                                                                                                                            _ = v200
                                                                                                                        } else {
                                                                                                                            fmt.Println(fi22, fi21)
                                                                                                                        }
                                                                                                                        var v206 int64 = 9223372036854775807
                                                                                                                        fmt.Println(fi22)
                                                                                                                        _ = v206
                                                                                                                    }
                                                                                                                }
                                                                                                                switch 519 {
                                                                                                                case 32768:
                                                                                                                    fmt.Println(v100, fi22, v4)
                                                                                                                case 2147483647:
                                                                                                                    v2 = v2
                                                                                                                case 1:
                                                                                                                    var v207 float32 = 3.14159
                                                                                                                    _ = v207
                                                                                                                default:
                                                                                                                    v3[(fi22 ^ int(v4))] = ("" + "a")
                                                                                                                }
                                                                                                                v208 := v3[980]
                                                                                                                _ = v208
                                                                                                            }
                                                                                                            _ = v186
                                                                                                        }
                                                                                                        v2 += v2
                                                                                                        v183 -= v183
                                                                                                    } else {
                                                                                                        v2 -= (^(v2))
                                                                                                    }
                                                                                                case 9223372036854775807:
                                                                                                    var v209 bool = false
                                                                                                    _ = v209
                                                                                                default:
                                                                                                    if true {
                                                                                                        v183 = 42
                                                                                                        var v210 bool = (!false)
                                                                                                        _ = v210
                                                                                                    } else {
                                                                                                        v211 := GenS23{F0: 0, F1: 133}
                                                                                                        v4 -= (v4 - v4)
                                                                                                        _ = v211
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                var v212 uint16 = 2
                                                                                                _ = v212
                                                                                            }
                                                                                            v213 := GenS17{F0: 111}
                                                                                            fmt.Println(fi22)
                                                                                            _ = v213
                                                                                        }
                                                                                        var v214 bool = true
                                                                                        v215 := GenS29{F0: int64(fi21), F1: v214}
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        _ = v183
                                                                                        _ = v214
                                                                                        _ = v215
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v216 := GenS30{F0: -594, F1: 255, F2: 0}
                                                                                _ = v216
                                                                            }
                                                                            if ((v4 > (^(v4))) || (false || false)) {
                                                                                var v217 uint8 = (v4 * 42)
                                                                                var v218 uint = 955
                                                                                v219 := GenS14{F0: 255, F1: 2, F2: v2}
                                                                                _ = v217
                                                                                _ = v218
                                                                                _ = v219
                                                                            } else {
                                                                                if false {
                                                                                    var v220 uint64 = 1
                                                                                    _ = v220
                                                                                }
                                                                                v2 += v2
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        v221 := GenS15{F0: 101}
                                                                        _ = v221
                                                                    default:
                                                                        v222 := GenS31{F0: v4}
                                                                        _ = v222
                                                                    }
                                                                default:
                                                                    fmt.Println(v100, v100)
                                                                }
                                                                v223 := GenS32{F0: (v2 / ((42) | 1)), F1: 1}
                                                                fmt.Println(fi21, v100)
                                                                _ = v100
                                                                _ = v223
                                                                _ = v99
                                                            }
                                                        }
                                                    }
                                                }
                                            } else {
                                                v2 += (v2 >> ((572) & 63))
                                                var v224 int16 = 32767
                                                v4 = v4
                                                fmt.Println(v2, v4, v224)
                                                fmt.Println(v2, v224, v224)
                                                _ = v224
                                            }
                                        case 4294967295:
                                            var v225 uint = 65535
                                            _ = v225
                                        case 460:
                                            v2 -= v2
                                        }
                                    case 32767:
                                        if true {
                                            if len(v1) > 0 {
                                                v1[0] = 4294967295
                                            }
                                            if false {
                                                var v226 uint = 42
                                                _ = v226
                                            } else {
                                                fmt.Println(v4, fi22, fi22)
                                                v227 := GenS31{F0: (^(v4))}
                                                var v228 int64 = -1
                                                _ = v227
                                                _ = v228
                                            }
                                            if true {
                                                var v229 float64 = 1.5
                                                var v230 uint64 = 1
                                                _ = v229
                                                _ = v230
                                            }
                                        }
                                    case 2147483647:
                                        if false {
                                            var v231 int16 = -862
                                            _ = v231
                                        } else {
                                            switch 856 {
                                            case -541:
                                                if (!(42 == -1)) {
                                                    v3[(fi22 % ((fi22) | 1))] = "a"
                                                }
                                            default:
                                                _ = len(v1)
                                            }
                                            var v232 int64 = -250
                                            _ = v232
                                        }
                                    default:
                                        if true {
                                            switch int(v4) {
                                            case 1:
                                                v233 := GenS33{F0: 2147483647, F1: 1}
                                                _ = v233
                                            default:
                                                v2 += (v2 / ((2147483647) | 1))
                                            }
                                            v4 = v4
                                        }
                                    }
                                }
                            }
                        default:
                            if false {
                                _ = len(v1)
                                if (!((!true) && true)) {
                                    var v234 uint32 = 424
                                    _ = len(v3)
                                    var v235 float64 = 0.0
                                    _ = v234
                                    _ = v235
                                }
                            }
                        }
                    }
                    _ = fi22
                }
                _ = fi21
            }
            v4 -= (v4 - (v4 & v4))
        }
        v236 := GenS29{F0: 9223372036854775807, F1: (!(24 <= 2))}
        for fi237 := 0; fi237 < 5; fi237++ {
            var v238 bool = true
            fmt.Println(fi237)
            _ = fi237
            _ = v238
        }
        _ = v236
    }
    var v239 int64 = 42
    switch "foo bar" {
    case "\n\t":
        var v240 bool = (42 < -128)
        _ = v240
    default:
        v241 := GenS31{F0: 2}
        _ = v241
    }
    var v242 float32 = (0.0 * (-2.25 + 0.0))
    fmt.Println(v4, v2, v242)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        for fi243 := 0; fi243 < 8; fi243++ {
            for fi244 := 0; fi244 < 17; fi244++ {
                v242 += (v242 + (0.0 * v242))
                _ = fi244
            }
            fmt.Println(fi243)
            _ = fi243
        }
        for fi245 := 0; fi245 < 9; fi245++ {
            if (139 >= 18446744073709551615) {
                if (((-9223372036854775808 == (-(v239))) || false) && (1000 >= 1000)) {
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                    v246 := GenS18{F0: (-(v2)), F1: (false || false), F2: (v239 >> ((56) & 63))}
                    var v247 float32 = v242
                    _ = v246
                    _ = v247
                }
                for fi248 := 0; fi248 < 1; fi248++ {
                    switch fi245 {
                    case 0:
                        v4 = 253
                        v249 := GenS13{F0: -408, F1: v239, F2: -97}
                        if true {
                            var v250 int32 = (v2 ^ v2)
                            _ = v250
                        }
                        var v251 float64 = -2.25
                        _ = v249
                        _ = v251
                    case -1:
                        fmt.Println(v239, v4, fi248)
                    default:
                        var v252 float64 = 1e10
                        _ = v252
                    }
                    v4 = (v4 ^ (v4 << ((65535) & 63)))
                    if (-2147483648 != v2) {
                        v253 := GenS12{F0: int32(v239), F1: 2, F2: int64(v2)}
                        _ = v253
                    }
                    switch (-(fi245)) {
                    case 32767:
                        switch ("a" + "a") {
                        case "hello":
                            v254 := GenS13{F0: v239, F1: 42, F2: 0}
                            _ = v254
                        case "a":
                            _ = len(v3)
                        default:
                            var v255 int8 = -66
                            _ = v255
                        }
                    case 9223372036854775807:
                        var v256 string = (((("a" + "a") + "") + ((("foo bar" + ("\n\t" + (("foo bar" + "\n\t") + ("foo bar" + "foo bar")))) + "hello") + (("hello" + "hello") + "a"))) + (((("foo bar" + "foo bar") + "") + "hello") + ""))
                        _ = v256
                    default:
                        var v257 int32 = v2
                        _ = v257
                    }
                    _ = fi248
                }
                v4 = (v4 / ((v4) | 1))
            } else {
                var v258 float32 = float32(-2.25)
                _ = v258
            }
            _ = fi245
        }
        switch (("" + "foo bar") + ("" + ("foo bar" + "foo bar"))) {
        case "foo bar":
            v259 := GenS34{F0: (true || ("\n\t" >= (("foo bar" + "\n\t") + ""))), F1: v2}
            _ = v259
        }
    } else {
        fmt.Println(v4, v242, v242)
    }
    if len(v1) > 0 {
        v1[0] = 65535
    }
    var v260 int = 42
    if false {
        _ = len(v3)
        var v261 int8 = 14
        _ = v261
    }
    switch "hello" {
    case "hello":
        v2 = v2
    case "a":
        v242 += (v242 + 1.5)
    }
    var v262 uint8 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v239:", v239)
    fmt.Println("v242:", v242)
    fmt.Println("v260:", v260)
    fmt.Println("v262:", v262)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v239
    _ = v242
    _ = v260
    _ = v262
    _ = v3
    _ = v4
    _ = v5
}
