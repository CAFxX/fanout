package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint; F1 int8; F2 int8 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 string; F1 string; F2 bool }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint; F1 int8; F2 int64 }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 int; F1 int32; F2 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int16; F1 int32; F2 int64 }
type GenS12 struct { F0 int32; F1 bool }
type GenS13 struct { F0 uint; F1 int16; F2 uint8 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int64; F1 bool }
type GenS16 struct { F0 uint8; F1 uint8 }
type GenS17 struct { F0 uint; F1 int32; F2 uint }
type GenS18 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 uint; F1 int64 }
type GenS21 struct { F0 int; F1 uint8 }
type GenS22 struct { F0 uint8; F1 int; F2 int }
type GenS23 struct { F0 int; F1 int8; F2 int32 }
type GenS24 struct { F0 string }
type GenS25 struct { F0 uint16; F1 int64 }
type GenS26 struct { F0 uint; F1 string; F2 int8 }
type GenS27 struct { F0 uint32; F1 int8 }
type GenS28 struct { F0 bool }

func main() {
    var v1 uint32 = 65535
    v2 := []string{"foo bar", "a"}
    v3 := []uint32{(v1 - 544), v1}
    for fi4 := 0; fi4 < 5; fi4++ {
        for fi5 := 0; fi5 < 8; fi5++ {
            v2 = append(v2, "")
            fmt.Println(v1, fi4)
            switch "foo bar" {
            case "foo bar":
                var v6 uint64 = 1
                _ = v6
            case "":
                fmt.Println(v1, fi5)
            case "hello":
                fmt.Println(v1, fi4, fi4)
            }
            _ = fi5
        }
        _ = fi4
    }
    fmt.Println(v1)
    if ((((1e10 * ((1.5 * (0.1 + ((1.5 + 1e10) - 0.0))) * 3.14159)) * 0.1) / float32(0.1)) == (float32((-2.25 / 3.14159)) * float32((((1e10 + float64(3.14159)) / -2.25) + 3.14159)))) {
        if true {
            for fi7 := 0; fi7 < 4; fi7++ {
                for fi8 := 0; fi8 < 2; fi8++ {
                    v9 := GenS1{F0: 0, F1: 801}
                    var v10 int64 = 9223372036854775807
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi7
            }
        } else {
            if true {
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                v3 = append(v3, v1)
                for fi11 := 0; fi11 < 9; fi11++ {
                    _ = len(v3)
                    switch "foo bar" {
                    case "foo bar":
                        v1 = v1
                    }
                    switch (("\n\t" + ("" + "\n\t")) + (("a" + ((("a" + "foo bar") + "hello") + "hello")) + "\n\t")) {
                    case "":
                        if (!true) {
                            if len(v3) > 0 {
                                v3[0] = (v1 - 287)
                            }
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                            var v12 int16 = 42
                            _ = v12
                        }
                    case "\n\t":
                        if false {
                            _ = len(v3)
                        }
                    case "foo bar":
                        switch (((("\n\t" + "\n\t") + ("" + "a")) + (("\n\t" + "") + (("a" + "foo bar") + "\n\t"))) + "hello") {
                        case "a":
                            v1 += 2147483647
                        case "hello":
                            var v13 int16 = 1
                            _ = v13
                        case "\n\t":
                            switch fi11 {
                            case -396:
                                if len(v3) > 0 {
                                    v3[0] = 2147483648
                                }
                            case 2147483647:
                                if (100 != 109) {
                                    for fi14 := 0; fi14 < 6; fi14++ {
                                        v1 -= (^(v1))
                                        _ = fi14
                                    }
                                    for fi15 := 0; fi15 < 17; fi15++ {
                                        v1 -= (v1 ^ (^(v1)))
                                        v16 := GenS2{F0: 18446744073709551615, F1: 1, F2: 100}
                                        if (!false) {
                                            v17 := GenS3{F0: "\n\t", F1: 18446744073709551615}
                                            var v18 int64 = -9223372036854775808
                                            _ = v17
                                            _ = v18
                                        }
                                        if true {
                                            if false {
                                                fmt.Println(fi11, v1)
                                                v1 += (v1 + v1)
                                                _ = len(v2)
                                            } else {
                                                v19 := GenS2{F0: 255, F1: -126, F2: -1}
                                                v1 = 783
                                                _ = v19
                                            }
                                            v1 = (v1 / ((2147483647) | 1))
                                        }
                                        v1 += 127
                                        _ = len(v2)
                                        if true {
                                            var v20 int64 = -1
                                            v21 := GenS2{F0: 4294967295, F1: 127, F2: -128}
                                            v22 := GenS2{F0: 0, F1: 42, F2: 59}
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        } else {
                                            _ = len(v3)
                                            fmt.Println(fi15)
                                        }
                                        _ = fi15
                                        _ = v16
                                    }
                                    if (false && (231 == 211)) {
                                        if (!false) {
                                            if (0 >= -32768) {
                                                if ((0 == 940) || false) {
                                                    var v23 int16 = -147
                                                    v24 := GenS4{F0: ("" + ("hello" + "")), F1: "", F2: (v1 < (v1 ^ 132))}
                                                    _ = v23
                                                    _ = v24
                                                }
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = (v1 ^ 2)
                                            }
                                            v25 := GenS5{F0: 128, F1: 32767}
                                            _ = v25
                                        }
                                        switch ("" + ("\n\t" + "a")) {
                                        case "hello":
                                            for fi26 := 0; fi26 < 14; fi26++ {
                                                switch (-(fi11)) {
                                                case -448:
                                                    if len(v2) > 0 {
                                                        v2[0] = ""
                                                    }
                                                case 73:
                                                    if false {
                                                        v1 = (v1 + 0)
                                                        v2 = append(v2, ("a" + ("" + "hello")))
                                                        fmt.Println(fi11, fi11, fi11)
                                                    } else {
                                                        v1 += v1
                                                        if true {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            fmt.Println(fi11)
                                                            v27 := GenS5{F0: (fi26 | fi26), F1: 128}
                                                            _ = v27
                                                        }
                                                        v1 = (v1 ^ v1)
                                                    }
                                                case 4294967295:
                                                    v28 := GenS1{F0: 1, F1: 18446744073709551615}
                                                    _ = v28
                                                }
                                                _ = fi26
                                            }
                                        }
                                        switch fi11 {
                                        case 9223372036854775807:
                                            v29 := GenS6{F0: -9223372036854775808}
                                            _ = v29
                                        case -799:
                                            var v30 uint16 = 42
                                            _ = v30
                                        default:
                                            switch "" {
                                            case "a":
                                                if true {
                                                    switch (^(fi11)) {
                                                    case 2147483647:
                                                        fmt.Println(fi11, fi11)
                                                    case 4294967295:
                                                        v31 := GenS2{F0: 2, F1: -91, F2: 42}
                                                        _ = v31
                                                    case 42:
                                                        var v32 int64 = 1
                                                        _ = v32
                                                    default:
                                                        var v33 uint = 42
                                                        _ = v33
                                                    }
                                                    v34 := GenS2{F0: 1, F1: 45, F2: 0}
                                                    _ = v34
                                                } else {
                                                    var v35 uint64 = 1
                                                    _ = v35
                                                }
                                            case "\n\t":
                                                switch fi11 {
                                                case 9223372036854775807:
                                                    for fi36 := 0; fi36 < 7; fi36++ {
                                                        fmt.Println(fi36, v1)
                                                        v37 := GenS5{F0: fi11, F1: 42}
                                                        _ = fi36
                                                        _ = v37
                                                    }
                                                default:
                                                    var v38 uint8 = 247
                                                    _ = v38
                                                }
                                            }
                                        }
                                    } else {
                                        switch fi11 {
                                        case -2147483648:
                                            var v39 int64 = 0
                                            _ = v39
                                        case 42:
                                            if true {
                                                switch 626 {
                                                case 200:
                                                    if false {
                                                        v40 := GenS7{F0: 4294967295, F1: 10, F2: -1}
                                                        for fi41 := 0; fi41 < 16; fi41++ {
                                                            v42 := GenS8{F0: false, F1: -1}
                                                            switch "foo bar" {
                                                            case "a":
                                                                v1 = (v1 / ((1) | 1))
                                                            case "hello":
                                                                _ = len(v2)
                                                            case "foo bar":
                                                                v1 = v1
                                                            default:
                                                                fmt.Println(fi41)
                                                            }
                                                            v1 = (v1 / ((v1) | 1))
                                                            _ = fi41
                                                            _ = v42
                                                        }
                                                        _ = v40
                                                    }
                                                case 273:
                                                    v1 -= v1
                                                default:
                                                    if (51 <= 967) {
                                                        switch 32767 {
                                                        case 9223372036854775807:
                                                            v43 := GenS9{F0: fi11, F1: -2147483648, F2: (fi11 << ((uint(v1)) & 63))}
                                                            _ = v43
                                                        default:
                                                            fmt.Println(fi11)
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = (^(v1))
                                                        }
                                                    }
                                                }
                                                v1 = v1
                                                var v44 uint8 = 1
                                                _ = v44
                                            } else {
                                                var v45 int8 = 0
                                                switch "foo bar" {
                                                case "":
                                                    if false {
                                                        if false {
                                                            switch fi11 {
                                                            case 9223372036854775807:
                                                                v1 = v1
                                                            case -9223372036854775808:
                                                                v1 = 768
                                                            case -2147483648:
                                                                v1 -= (^(v1))
                                                            default:
                                                                var v46 int = 127
                                                                _ = v46
                                                            }
                                                            var v47 float64 = -2.25
                                                            _ = v47
                                                        }
                                                        for fi48 := 0; fi48 < 10; fi48++ {
                                                            v49 := GenS10{F0: ((("hello" + "a") < ("" + "hello")) || true)}
                                                            fmt.Println(v45, fi11)
                                                            _ = fi48
                                                            _ = v49
                                                        }
                                                        switch (^(fi11)) {
                                                        case 32768:
                                                            v45 -= (v45 << ((uint(v1)) & 63))
                                                        case 128:
                                                            if false {
                                                                v1 = (^(v1))
                                                                fmt.Println(v1)
                                                            } else {
                                                                v50 := GenS11{F0: -32768, F1: 316, F2: 0}
                                                                var v51 uint8 = 29
                                                                if false {
                                                                    var v52 uint64 = 4294967295
                                                                    for fi53 := 0; fi53 < 5; fi53++ {
                                                                        v52 += v52
                                                                        v51 = (v51 - (v51 | (v51 ^ 83)))
                                                                        if (4294967295 > v52) {
                                                                            v52 = (v52 >> ((uint(v52)) & 63))
                                                                            var v54 int32 = -611
                                                                            v54 = (v54 % ((220) | 1))
                                                                            _ = v54
                                                                        } else {
                                                                            var v55 float32 = 3.14159
                                                                            if ((("" + ("a" + "hello")) >= "") && (!true)) {
                                                                                var v56 uint64 = (v52 + (v52 - v52))
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = v1
                                                                                }
                                                                                if (v51 <= (v51 ^ 42)) {
                                                                                    v57 := GenS12{F0: int32(fi11), F1: true}
                                                                                    if (42 <= 18446744073709551615) {
                                                                                        switch 642 {
                                                                                        case 241:
                                                                                            v58 := GenS13{F0: 613, F1: -1, F2: 128}
                                                                                            _ = v58
                                                                                        case 128:
                                                                                            var v59 uint = 1
                                                                                            v51 = (v51 << ((65535) & 63))
                                                                                            _ = v59
                                                                                        }
                                                                                        v51 = 255
                                                                                    }
                                                                                    var v60 int8 = 127
                                                                                    _ = v57
                                                                                    _ = v60
                                                                                }
                                                                                _ = v56
                                                                            } else {
                                                                                var v61 uint16 = 32767
                                                                                var v62 uint = 65535
                                                                                if false {
                                                                                    v62 -= (v62 + uint(v52))
                                                                                } else {
                                                                                    switch (fi53 / (((fi53 * 128)) | 1)) {
                                                                                    case 0:
                                                                                        v63 := GenS14{F0: (v51 / (((v51 ^ v51)) | 1))}
                                                                                        var v64 uint = v62
                                                                                        var v65 uint32 = (v1 | v1)
                                                                                        _ = v63
                                                                                        _ = v64
                                                                                        _ = v65
                                                                                    case -9223372036854775808:
                                                                                        v52 = v52
                                                                                    case -1:
                                                                                        var v66 uint8 = (^(v51))
                                                                                        _ = v66
                                                                                    }
                                                                                    v67 := GenS9{F0: (fi11 | 127), F1: -922, F2: (fi53 | int(v1))}
                                                                                    fmt.Println(v51, v61, v51)
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    _ = v67
                                                                                }
                                                                                _ = v61
                                                                                _ = v62
                                                                            }
                                                                            _ = v55
                                                                        }
                                                                        switch 4294967295 {
                                                                        case 1:
                                                                            v45 = v45
                                                                        default:
                                                                            v51 -= (v51 / (((v51 - v51)) | 1))
                                                                        }
                                                                        switch fi53 {
                                                                        case -103:
                                                                            v2 = append(v2, "a")
                                                                        default:
                                                                            switch ("hello" + ("hello" + "")) {
                                                                            case "a":
                                                                                var v68 uint64 = v52
                                                                                _ = v68
                                                                            default:
                                                                                switch fi53 {
                                                                                case 4294967295:
                                                                                    v45 = (v45 * v45)
                                                                                case -2147483648:
                                                                                    v51 += (v51 | v51)
                                                                                }
                                                                            }
                                                                            if (fi11 >= (fi11 / ((fi53) | 1))) {
                                                                                fmt.Println(v52, v1, v45)
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    var v69 bool = (!((v51 % ((v51) | 1)) >= uint8(v1)))
                                                                                    v69 = (1 < (v1 & v1))
                                                                                    _ = v69
                                                                                case "a":
                                                                                    if (100 <= (v45 * v45)) {
                                                                                        v70 := GenS15{F0: 1, F1: false}
                                                                                        switch 32767 {
                                                                                        case -9223372036854775808:
                                                                                            var v71 int32 = 2147483647
                                                                                            _ = v71
                                                                                        default:
                                                                                            var v72 uint32 = v1
                                                                                            _ = v72
                                                                                        }
                                                                                        _ = v70
                                                                                    } else {
                                                                                        v73 := GenS9{F0: (^(fi11)), F1: 0, F2: -2147483648}
                                                                                        switch "hello" {
                                                                                        case "hello":
                                                                                            fmt.Println(fi11, fi11, v51)
                                                                                            switch "" {
                                                                                            case "":
                                                                                                if true {
                                                                                                    var v74 uint32 = (v1 << ((400) & 63))
                                                                                                    _ = v74
                                                                                                }
                                                                                            case "a":
                                                                                                var v75 string = "a"
                                                                                                _ = v75
                                                                                            default:
                                                                                                v76 := GenS16{F0: v51, F1: (v51 << ((422) & 63))}
                                                                                                _ = v76
                                                                                            }
                                                                                        case "a":
                                                                                            v52 += (v52 >> ((65535) & 63))
                                                                                        case "foo bar":
                                                                                            v77 := GenS10{F0: (-2.25 > 0.1)}
                                                                                            _ = v77
                                                                                        default:
                                                                                            switch 255 {
                                                                                            case -9223372036854775808:
                                                                                                switch (fi11 ^ fi53) {
                                                                                                case 128:
                                                                                                    switch fi11 {
                                                                                                    case 228:
                                                                                                        var v78 bool = (!false)
                                                                                                        _ = v78
                                                                                                    default:
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                    }
                                                                                                case 42:
                                                                                                    if false {
                                                                                                        if false {
                                                                                                            if false {
                                                                                                                var v79 uint32 = (v1 % (((v1 & v1)) | 1))
                                                                                                                _ = v79
                                                                                                            } else {
                                                                                                                var v80 uint32 = v1
                                                                                                                fmt.Println(fi11)
                                                                                                                if true {
                                                                                                                    if true {
                                                                                                                        var v81 int = (fi11 & int(v1))
                                                                                                                        var v82 int32 = 2147483647
                                                                                                                        v80 += (v1 | v1)
                                                                                                                        if (0 > (fi11 | (-(v81)))) {
                                                                                                                            v82 = v82
                                                                                                                        }
                                                                                                                        _ = v81
                                                                                                                        _ = v82
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v83 := GenS10{F0: true}
                                                                                                                    _ = v83
                                                                                                                }
                                                                                                                _ = v80
                                                                                                            }
                                                                                                            switch ((("hello" + "") + (("foo bar" + "hello") + "a")) + "hello") {
                                                                                                            case "\n\t":
                                                                                                                _ = len(v3)
                                                                                                            default:
                                                                                                                if true {
                                                                                                                    var v84 float64 = ((3.14159 + 1.5) - 1.5)
                                                                                                                    _ = v84
                                                                                                                } else {
                                                                                                                    v85 := GenS13{F0: 255, F1: 0, F2: v51}
                                                                                                                    _ = v85
                                                                                                                }
                                                                                                            }
                                                                                                            v86 := GenS13{F0: 612, F1: 32767, F2: (v51 - (v51 / (((^(v51))) | 1)))}
                                                                                                            _ = v86
                                                                                                        }
                                                                                                        v87 := GenS2{F0: 18446744073709551615, F1: (v45 & (v45 << ((1) & 63))), F2: 68}
                                                                                                        _ = len(v2)
                                                                                                        _ = v87
                                                                                                    }
                                                                                                default:
                                                                                                    v3 = append(v3, 2147483648)
                                                                                                }
                                                                                            case 128:
                                                                                                v45 -= -1
                                                                                            case 32768:
                                                                                                if true {
                                                                                                    if false {
                                                                                                        switch "" {
                                                                                                        case "":
                                                                                                            var v88 bool = (-1 == 9223372036854775807)
                                                                                                            _ = v88
                                                                                                        case "\n\t":
                                                                                                            var v89 int16 = 42
                                                                                                            _ = v89
                                                                                                        case "foo bar":
                                                                                                            v45 -= (v45 + v45)
                                                                                                        default:
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                        }
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                        v51 -= (v51 << ((4294967295) & 63))
                                                                                                    } else {
                                                                                                        if (124 < 1) {
                                                                                                            v90 := GenS5{F0: (fi53 & (fi53 | int(v1))), F1: fi11}
                                                                                                            if true {
                                                                                                                switch "a" {
                                                                                                                case "\n\t":
                                                                                                                    v91 := GenS17{F0: 331, F1: 856, F2: 1}
                                                                                                                    _ = v91
                                                                                                                }
                                                                                                                var v92 float32 = -2.25
                                                                                                                v52 -= 2
                                                                                                                _ = v92
                                                                                                            } else {
                                                                                                                var v93 int = fi53
                                                                                                                if false {
                                                                                                                    if len(v3) >= 2 {
                                                                                                                        v3 = v3[1:len(v3)]
                                                                                                                    }
                                                                                                                    v45 -= (-(v45))
                                                                                                                }
                                                                                                                _ = v93
                                                                                                            }
                                                                                                            _ = v90
                                                                                                        } else {
                                                                                                            if len(v2) >= 2 {
                                                                                                                v2 = v2[1:len(v2)]
                                                                                                            }
                                                                                                            var v94 float64 = float64(float32(0.1))
                                                                                                            _ = v94
                                                                                                        }
                                                                                                        var v95 int32 = -852
                                                                                                        _ = v95
                                                                                                    }
                                                                                                    if (0.0 != 0.1) {
                                                                                                        var v96 uint32 = v1
                                                                                                        var v97 int64 = int64(v45)
                                                                                                        if (0.0 > float32(0.0)) {
                                                                                                            if false {
                                                                                                                var v98 uint16 = 1
                                                                                                                _ = v98
                                                                                                            }
                                                                                                        }
                                                                                                        var v99 string = "foo bar"
                                                                                                        _ = v96
                                                                                                        _ = v97
                                                                                                        _ = v99
                                                                                                    } else {
                                                                                                        fmt.Println(fi11, v1, fi53)
                                                                                                        var v100 int8 = (v45 + v45)
                                                                                                        v1 = v1
                                                                                                        _ = v100
                                                                                                    }
                                                                                                    v101 := GenS18{F0: 42, F1: 568, F2: v52}
                                                                                                    _ = v101
                                                                                                } else {
                                                                                                    v45 -= v45
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v73
                                                                                    }
                                                                                default:
                                                                                    if false {
                                                                                        fmt.Println(v1, v45, v52)
                                                                                        v3 = append(v3, 2147483648)
                                                                                    }
                                                                                }
                                                                                var v102 bool = (true || ((1.5 / 0.1) < 1e10))
                                                                                switch (fi53 % ((fi53) | 1)) {
                                                                                case 42:
                                                                                    var v103 uint8 = uint8(v1)
                                                                                    if (!(((v102 && v102) || (true || (!v102))) && (v52 != (v52 / (((v52 + 1)) | 1))))) {
                                                                                        v102 = v102
                                                                                        v104 := GenS19{F0: -1}
                                                                                        if (!v102) {
                                                                                            v1 -= v1
                                                                                        }
                                                                                        _ = v104
                                                                                    } else {
                                                                                        if v102 {
                                                                                            if v102 {
                                                                                                v3 = append(v3, (v1 + (v1 & (v1 % ((v1) | 1)))))
                                                                                                if v102 {
                                                                                                    v102 = v102
                                                                                                    fmt.Println(fi53, v103, v1)
                                                                                                    v1 += (^(v1))
                                                                                                }
                                                                                            } else {
                                                                                                v102 = (3.14159 < 0.1)
                                                                                                var v105 string = (("foo bar" + "") + ("foo bar" + ""))
                                                                                                v103 += (v51 + v51)
                                                                                                switch -842 {
                                                                                                case 1:
                                                                                                    v45 += v45
                                                                                                }
                                                                                                v106 := GenS18{F0: uint8(v52), F1: v1, F2: 2}
                                                                                                _ = v105
                                                                                                _ = v106
                                                                                            }
                                                                                            v45 = v45
                                                                                        } else {
                                                                                            if (2 <= 0) {
                                                                                                fmt.Println(v52, v102)
                                                                                            }
                                                                                            if false {
                                                                                                v107 := GenS10{F0: ((!(!(v102 || false))) && false)}
                                                                                                _ = v107
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v103
                                                                                case 0:
                                                                                    if v102 {
                                                                                        switch 32768 {
                                                                                        case 2147483647:
                                                                                            var v108 float32 = 3.14159
                                                                                            _ = v108
                                                                                        case 42:
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case 1:
                                                                                    var v109 int16 = 0
                                                                                    _ = v109
                                                                                default:
                                                                                    fmt.Println(v52, v52)
                                                                                }
                                                                                var v110 int8 = v45
                                                                                v111 := GenS6{F0: 11}
                                                                                _ = v102
                                                                                _ = v110
                                                                                _ = v111
                                                                            }
                                                                            v112 := GenS3{F0: "\n\t", F1: (v52 & (v52 & (v52 / (((v52 - 1)) | 1))))}
                                                                            _ = v112
                                                                        }
                                                                        if true {
                                                                            switch fi53 {
                                                                            case 128:
                                                                                var v113 int64 = 0
                                                                                _ = v113
                                                                            default:
                                                                                _ = len(v2)
                                                                            }
                                                                        }
                                                                        fmt.Println(v51, v45)
                                                                        v114 := GenS18{F0: v51, F1: v1, F2: v52}
                                                                        _ = fi53
                                                                        _ = v114
                                                                    }
                                                                    _ = v52
                                                                }
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        }
                                                    }
                                                case "hello":
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v115 uint32 = v1
                                                        _ = v115
                                                    case "":
                                                        for fi116 := 0; fi116 < 16; fi116++ {
                                                            var v117 int8 = 100
                                                            _ = fi116
                                                            _ = v117
                                                        }
                                                    }
                                                case "foo bar":
                                                    if true {
                                                        switch 32768 {
                                                        case -1:
                                                            var v118 uint = 2
                                                            _ = v118
                                                        }
                                                    }
                                                }
                                                _ = v45
                                            }
                                        }
                                    }
                                }
                            case 760:
                                var v119 uint32 = (v1 - v1)
                                _ = v119
                            default:
                                fmt.Println(v1, fi11)
                            }
                        default:
                            for fi120 := 0; fi120 < 17; fi120++ {
                                v121 := GenS20{F0: 1, F1: 42}
                                switch (fi11 * (fi11 / ((fi120) | 1))) {
                                case 42:
                                    var v122 uint64 = 4294967295
                                    _ = v122
                                default:
                                    if (-2.25 > (((-2.25 - 0.1) / ((3.14159 + (0.0 / (float32(-2.25) + 3.14159))) * (1e10 * (1e10 / 3.14159)))) * 1e10)) {
                                        var v123 uint64 = 4294967295
                                        _ = v123
                                    } else {
                                        v1 = v1
                                        switch (fi11 * -2147483648) {
                                        case 4294967295:
                                            v124 := GenS21{F0: (fi11 & (^(fi120))), F1: 42}
                                            _ = v124
                                        case 42:
                                            switch "hello" {
                                            case "foo bar":
                                                fmt.Println(fi11, fi120)
                                            case "hello":
                                                v125 := GenS2{F0: 386, F1: 0, F2: 100}
                                                _ = v125
                                            case "":
                                                fmt.Println(v1, v1)
                                            }
                                            if ((0 >= 2) || false) {
                                                fmt.Println(fi11)
                                                var v126 uint = uint(v1)
                                                v126 = v126
                                                v127 := GenS22{F0: 159, F1: (-(fi120)), F2: 331}
                                                v1 += (v1 + (^(v1)))
                                                _ = v126
                                                _ = v127
                                            } else {
                                                v128 := GenS2{F0: 18446744073709551615, F1: 42, F2: -128}
                                                if false {
                                                    switch 779 {
                                                    case 128:
                                                        switch 398 {
                                                        case 42:
                                                            fmt.Println(fi11)
                                                        case 149:
                                                            var v129 uint8 = 194
                                                            _ = v129
                                                        }
                                                    case -2147483648:
                                                        var v130 int8 = -27
                                                        _ = v130
                                                    case 4294967295:
                                                        v131 := GenS9{F0: (fi11 - (fi120 & (fi11 >> ((255) & 63)))), F1: 0, F2: fi11}
                                                        v1 += 0
                                                        if false {
                                                            if true {
                                                                if len(v2) > 0 {
                                                                    v2[0] = "a"
                                                                }
                                                                v132 := GenS23{F0: (fi11 * (fi120 << ((255) & 63))), F1: 127, F2: -266}
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = v132
                                                            } else {
                                                                switch ("a" + ("" + (("hello" + "") + "a"))) {
                                                                case "hello":
                                                                    fmt.Println(fi11, v1)
                                                                case "a":
                                                                    fmt.Println(fi120)
                                                                }
                                                                var v133 uint16 = 572
                                                                var v134 string = "a"
                                                                _ = v133
                                                                _ = v134
                                                            }
                                                            v1 += v1
                                                        }
                                                        _ = v131
                                                    default:
                                                        v1 = 844
                                                    }
                                                    fmt.Println(fi120)
                                                    fmt.Println(v1)
                                                }
                                                v3 = append(v3, (v1 & v1))
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v135 uint16 = 692
                                                    _ = v135
                                                }
                                                _ = v128
                                            }
                                        case 0:
                                            v1 = v1
                                        }
                                    }
                                }
                                var v136 int = fi11
                                _ = fi120
                                _ = v121
                                _ = v136
                            }
                        }
                    default:
                        if (4294967295 == 4294967295) {
                            v137 := GenS19{F0: -680}
                            _ = len(v3)
                            _ = v137
                        }
                    }
                    _ = fi11
                }
            }
        }
        var v138 float32 = (1.5 + (float32(-2.25) - ((0.1 + (3.14159 + 1.5)) - 3.14159)))
        _ = v138
    } else {
        var v139 uint64 = 1
        _ = v139
    }
    var v140 int = 161
    var v141 int = (v140 % ((v140) | 1))
    var v142 float64 = ((((((((1e10 + 0.1) - float64(-2.25)) / 1.0) - 1.5) + 1e10) + (0.1 + 0.1)) / 3.14159) / ((((1e10 + 0.1) / 1.0) + 1e10) * 1e10))
    switch v141 {
    case 525:
        switch "hello" {
        case "":
            var v143 int32 = 0
            _ = v143
        case "hello":
            var v144 int = int(v1)
            _ = v144
        case "a":
            switch v140 {
            case 467:
                v141 += (v141 + (^(v141)))
            case -136:
                var v145 int64 = 59
                _ = v145
            case 2147483647:
                var v146 uint64 = 4294967295
                _ = v146
            }
        default:
            switch ((((("hello" + "\n\t") + "a") + ("a" + "")) + "hello") + ("hello" + (((("hello" + (("\n\t" + "a") + "hello")) + "a") + (("\n\t" + "a") + "a")) + ""))) {
            case "a":
                fmt.Println(v141, v1, v142)
            }
        }
    case -9223372036854775808:
        for fi147 := 0; fi147 < 17; fi147++ {
            v1 += (v1 % ((v1) | 1))
            _ = fi147
        }
    case 209:
        for fi148 := 0; fi148 < 3; fi148++ {
            v149 := GenS11{F0: -1, F1: 865, F2: -286}
            if false {
                fmt.Println(v1, v141)
            } else {
                for fi150 := 0; fi150 < 1; fi150++ {
                    switch (fi148 << ((2) & 63)) {
                    case 601:
                        if false {
                            v151 := GenS20{F0: 1, F1: -754}
                            v1 += v1
                            if true {
                                fmt.Println(fi150, v140)
                                var v152 string = "hello"
                                v153 := GenS24{F0: ""}
                                _ = v152
                                _ = v153
                            } else {
                                var v154 bool = false
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if (1.5 < 0.0) {
                                    if (488 <= 1) {
                                        v1 -= (v1 | 891)
                                        v155 := GenS13{F0: 42, F1: 32767, F2: 2}
                                        fmt.Println(v154, v154)
                                        _ = v155
                                    } else {
                                        v156 := GenS25{F0: 959, F1: 251}
                                        _ = v156
                                    }
                                    var v157 int8 = 15
                                    v1 = (v1 >> ((1) & 63))
                                    _ = v157
                                }
                                _ = v154
                            }
                            _ = v151
                        } else {
                            v142 -= v142
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                        }
                    }
                    var v158 int32 = 0
                    switch "hello" {
                    case "":
                        var v159 uint64 = 503
                        _ = v159
                    }
                    _ = fi150
                    _ = v158
                }
            }
            _ = fi148
            _ = v149
        }
    default:
        fmt.Println(v142)
    }
    if true {
        if false {
            var v160 int8 = 1
            v161 := GenS26{F0: uint(v1), F1: (("\n\t" + "\n\t") + "\n\t"), F2: -1}
            _ = v160
            _ = v161
        } else {
            if ("" >= "foo bar") {
                v162 := GenS6{F0: 9223372036854775807}
                if false {
                    if false {
                        fmt.Println(v141, v141)
                        for fi163 := 0; fi163 < 11; fi163++ {
                            if false {
                                _ = len(v2)
                                if true {
                                    var v164 int8 = 109
                                    for fi165 := 0; fi165 < 5; fi165++ {
                                        _ = len(v3)
                                        _ = fi165
                                    }
                                    _ = v164
                                }
                            }
                            _ = fi163
                        }
                        switch (("hello" + "") + "a") {
                        case "\n\t":
                            for fi166 := 0; fi166 < 19; fi166++ {
                                switch -2147483648 {
                                case 1:
                                    for fi167 := 0; fi167 < 8; fi167++ {
                                        v168 := GenS20{F0: 2, F1: 42}
                                        switch v141 {
                                        case 32767:
                                            fmt.Println(fi166, fi166)
                                        case -504:
                                            v140 = 32768
                                        case 981:
                                            v169 := GenS27{F0: v1, F1: -1}
                                            _ = v169
                                        default:
                                            v141 -= v140
                                        }
                                        _ = fi167
                                        _ = v168
                                    }
                                case 32767:
                                    fmt.Println(fi166, v1, v140)
                                default:
                                    var v170 int8 = 100
                                    _ = v170
                                }
                                _ = fi166
                            }
                        case "foo bar":
                            for fi171 := 0; fi171 < 11; fi171++ {
                                var v172 uint32 = v1
                                var v173 int32 = 146
                                _ = fi171
                                _ = v172
                                _ = v173
                            }
                        }
                    }
                    var v174 int64 = -1
                    _ = v174
                } else {
                    switch v141 {
                    case -110:
                        fmt.Println(v142, v141, v140)
                    case 587:
                        v1 += v1
                    case 928:
                        v141 += v140
                    default:
                        v175 := GenS24{F0: "hello"}
                        _ = v175
                    }
                    v176 := GenS2{F0: 42, F1: 42, F2: 0}
                    _ = v176
                }
                _ = v162
            }
        }
        _ = len(v2)
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    } else {
        var v177 int64 = 708
        _ = len(v3)
        v178 := GenS28{F0: false}
        _ = v177
        _ = v178
    }
    if false {
        if false {
            fmt.Println(v141)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("v142:", v142)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v140
    _ = v141
    _ = v142
    _ = v2
    _ = v3
}
