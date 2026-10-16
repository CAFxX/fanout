package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string; F1 bool; F2 uint64 }
type GenS10 struct { F0 int8; F1 int8; F2 uint32 }
type GenS11 struct { F0 uint8; F1 int64 }
type GenS12 struct { F0 uint; F1 bool; F2 int32 }

func main() {
    var v1 uint8 = 2
    var v2 int16 = 1000
    var v3 uint8 = (^(v1))
    var v4 uint64 = 448
    var v5 uint16 = 32767
    v6 := []int{-195, 32768, -1}
    v5 += v5
    var v7 bool = (1e10 < ((float64(1.5) * 0.1) * 0.1))
    v4 -= v4
    switch "a" {
    case "":
        for fi8 := 0; fi8 < 15; fi8++ {
            v6 = append(v6, -553)
            _ = fi8
        }
    case "a":
        v9 := GenS1{F0: (v2 << ((1) & 63)), F1: (^(v3))}
        _ = v9
    }
    v1 -= v1
    switch 580 {
    case 127:
        v2 += -323
    case 128:
        for fi10 := 0; fi10 < 4; fi10++ {
            switch fi10 {
            case -78:
                var v11 int8 = 42
                _ = v11
            case 128:
                for fi12 := 0; fi12 < 13; fi12++ {
                    v7 = (65535 >= 623)
                    v13 := GenS1{F0: v2, F1: 2}
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v14 := GenS2{F0: 2147483647}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            _ = fi10
        }
    }
    if false {
        v15 := GenS2{F0: -986}
        _ = v15
    } else {
        if v7 {
            if (v7 && v7) {
                fmt.Println(v1, v2, v7)
                var v16 float64 = 3.14159
                switch ("foo bar" + "foo bar") {
                case "":
                    v2 += (v2 ^ v2)
                case "hello":
                    for fi17 := 0; fi17 < 7; fi17++ {
                        var v18 int64 = -9223372036854775808
                        switch 9223372036854775807 {
                        case 4294967295:
                            if v7 {
                                var v19 int8 = -56
                                var v20 int8 = v19
                                _ = v19
                                _ = v20
                            }
                        case 32768:
                            _ = len(v6)
                        default:
                            var v21 float64 = (0.0 - (v16 + v16))
                            _ = v21
                        }
                        _ = fi17
                        _ = v18
                    }
                case "\n\t":
                    if (v7 || false) {
                        for fi22 := 0; fi22 < 8; fi22++ {
                            fmt.Println(v1)
                            for fi23 := 0; fi23 < 3; fi23++ {
                                fmt.Println(fi23)
                                var v24 int8 = 42
                                var v25 uint = 97
                                if true {
                                    switch int(v1) {
                                    case -912:
                                        var v26 int8 = (v24 << (((^(v25))) & 63))
                                        var v27 bool = v7
                                        fmt.Println(v26, fi22, v26)
                                        _ = v26
                                        _ = v27
                                    case 32767:
                                        if v7 {
                                            v2 = -32768
                                            var v28 uint8 = 5
                                            fmt.Println(v28, fi22, v5)
                                            _ = v28
                                        }
                                    case -669:
                                        switch "a" {
                                        case "foo bar":
                                            v7 = false
                                        case "a":
                                            if v7 {
                                                var v29 int64 = 267
                                                var v30 uint64 = v4
                                                v31 := GenS2{F0: -2147483648}
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            }
                                        case "hello":
                                            v32 := GenS3{F0: int(v24), F1: -9223372036854775808, F2: 4294967295}
                                            _ = v32
                                        default:
                                            var v33 uint16 = (v5 / ((32767) | 1))
                                            v2 = int16(v24)
                                            _ = v33
                                        }
                                    }
                                    fmt.Println(v5)
                                } else {
                                    if len(v6) > 0 {
                                        v6[0] = int(v2)
                                    }
                                    fmt.Println(v2, v24, v5)
                                    if ((v5 * v5) != 0) {
                                        if ((v25 <= 4294967295) || (((v25 % ((18446744073709551615) | 1)) >= v25) && v7)) {
                                            v1 = (v3 | v3)
                                            if len(v6) > 0 {
                                                v6[0] = 2147483647
                                            }
                                            v2 += v2
                                            fmt.Println(v2, v16, v25)
                                        }
                                    } else {
                                        v25 = (v25 / (((v25 << ((v25) & 63))) | 1))
                                        v4 += 2
                                        v34 := GenS3{F0: (fi22 % (((-(fi23))) | 1)), F1: -9223372036854775808, F2: 393}
                                        _ = v34
                                    }
                                }
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            if (782 == 1) {
                                if len(v6) > 0 {
                                    v6[0] = fi22
                                }
                                if len(v6) > 0 {
                                    v6[0] = 32768
                                }
                            }
                            _ = fi22
                        }
                        var v35 int64 = 9223372036854775807
                        _ = v35
                    } else {
                        switch 32767 {
                        case 42:
                            v4 -= uint64(v5)
                        case 0:
                            v2 = int16(v5)
                        case 1:
                            var v36 uint32 = 978
                            _ = v36
                        default:
                            fmt.Println(v4)
                        }
                        if v7 {
                            v37 := GenS4{F0: (v5 * v5), F1: 2}
                            if true {
                                if (("\n\t" + (("\n\t" + "") + "\n\t")) > "foo bar") {
                                    if (!("a" >= "a")) {
                                        v38 := GenS4{F0: v5, F1: 9}
                                        v39 := GenS4{F0: v5, F1: 18446744073709551615}
                                        _ = v38
                                        _ = v39
                                    } else {
                                        if true {
                                            v40 := GenS3{F0: 128, F1: -1, F2: 42}
                                            _ = v40
                                        }
                                        v4 -= v4
                                        _ = len(v6)
                                    }
                                    v4 -= 824
                                }
                                for fi41 := 0; fi41 < 14; fi41++ {
                                    if v7 {
                                        var v42 float32 = float32(((v16 + v16) / v16))
                                        switch "foo bar" {
                                        case "foo bar":
                                            v43 := GenS2{F0: 0}
                                            _ = v43
                                        case "a":
                                            fmt.Println(fi41)
                                        }
                                        _ = v42
                                    } else {
                                        var v44 bool = false
                                        switch (-(fi41)) {
                                        case 127:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        case 32768:
                                            if (!(!true)) {
                                                fmt.Println(v5, v44)
                                                var v45 float64 = (0.1 - v16)
                                                var v46 int = (fi41 >> ((255) & 63))
                                                _ = v45
                                                _ = v46
                                            }
                                        case -1:
                                            fmt.Println(v7)
                                        default:
                                            var v47 int64 = 0
                                            _ = v47
                                        }
                                        fmt.Println(v16, v5, v4)
                                        _ = v44
                                    }
                                    switch (fi41 - (^(fi41))) {
                                    case -780:
                                        if len(v6) > 0 {
                                            v6[0] = (fi41 << ((4294967295) & 63))
                                        }
                                    default:
                                        var v48 int8 = 7
                                        _ = v48
                                    }
                                    _ = fi41
                                }
                            } else {
                                var v49 uint32 = 1
                                switch 271 {
                                case 557:
                                    v16 = v16
                                case 32767:
                                    var v50 float32 = 3.14159
                                    _ = v50
                                }
                                if v7 {
                                    if true {
                                        for fi51 := 0; fi51 < 17; fi51++ {
                                            if (492 < fi51) {
                                                v52 := GenS1{F0: -32768, F1: 42}
                                                if (v16 < -2.25) {
                                                    var v53 uint = uint(v49)
                                                    _ = v53
                                                } else {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    if (((!(-2.25 == float32(0.1))) || false) && v7) {
                                                        if (((3.14159 != ((1.5 + (1e10 - (0.0 / 1.5))) * ((1e10 + (0.1 * 1.5)) - 1.5))) || v7) || v7) {
                                                            for fi54 := 0; fi54 < 5; fi54++ {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                v55 := GenS5{F0: "\n\t", F1: uint(v49)}
                                                                fmt.Println(v5, v3, v49)
                                                                _ = fi54
                                                                _ = v55
                                                            }
                                                            if v7 {
                                                                if ((v1 >> ((4294967295) & 63)) == (v1 * v1)) {
                                                                    v16 = -2.25
                                                                    var v56 int = fi51
                                                                    _ = v56
                                                                } else {
                                                                    switch (fi51 ^ 0) {
                                                                    case 2147483647:
                                                                        v16 = (v16 - float64(-2.25))
                                                                    case 9223372036854775807:
                                                                        v49 = v49
                                                                    }
                                                                    v2 = v2
                                                                    var v57 uint8 = v1
                                                                    _ = v57
                                                                }
                                                                if true {
                                                                    for fi58 := 0; fi58 < 13; fi58++ {
                                                                        if (100 != 76) {
                                                                            switch "" {
                                                                            case "a":
                                                                                fmt.Println(v2, v49)
                                                                            case "foo bar":
                                                                                var v59 int8 = -1
                                                                                _ = v59
                                                                            case "":
                                                                                var v60 int32 = -1
                                                                                _ = v60
                                                                            }
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            v1 += (v3 % (((v3 + (^(v1)))) | 1))
                                                                        } else {
                                                                            v49 = (v49 | v49)
                                                                        }
                                                                        v6 = append(v6, 42)
                                                                        var v61 float64 = ((1.5 * (-2.25 + 0.1)) - v16)
                                                                        if v7 {
                                                                            v6 = append(v6, (fi51 ^ (fi58 & -681)))
                                                                            switch -2147483648 {
                                                                            case 32767:
                                                                                var v62 uint64 = v4
                                                                                _ = v62
                                                                            case 4294967295:
                                                                                if ((v2 % ((1000) | 1)) < 98) {
                                                                                    v63 := GenS5{F0: "\n\t", F1: 18446744073709551615}
                                                                                    fmt.Println(v5, v49)
                                                                                    _ = v63
                                                                                } else {
                                                                                    fmt.Println(v1)
                                                                                    v6 = append(v6, (fi58 | 9223372036854775807))
                                                                                }
                                                                            case 127:
                                                                                fmt.Println(fi51, v1)
                                                                            default:
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    v64 := GenS5{F0: "a", F1: 18446744073709551615}
                                                                                    _ = v64
                                                                                }
                                                                            }
                                                                            _ = len(v6)
                                                                        } else {
                                                                            var v65 int8 = 100
                                                                            _ = v65
                                                                        }
                                                                        fmt.Println(v16, v4, v49)
                                                                        _ = fi58
                                                                        _ = v61
                                                                    }
                                                                    for fi66 := 0; fi66 < 3; fi66++ {
                                                                        if true {
                                                                            v67 := GenS6{F0: 901}
                                                                            v68 := GenS1{F0: (^(v2)), F1: 85}
                                                                            _ = v67
                                                                            _ = v68
                                                                        }
                                                                        v2 = (^(v2))
                                                                        var v69 float64 = (v16 * 1e10)
                                                                        _ = fi66
                                                                        _ = v69
                                                                    }
                                                                }
                                                            } else {
                                                                v3 += (v3 ^ 65)
                                                            }
                                                            var v70 int = fi51
                                                            _ = v70
                                                        }
                                                        fmt.Println(v3)
                                                        var v71 int32 = 42
                                                        _ = v71
                                                    }
                                                    if (-837 >= -9223372036854775808) {
                                                        v6 = append(v6, fi51)
                                                    } else {
                                                        if false {
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                if v7 {
                                                                    if ((v3 << ((2) & 63)) == v3) {
                                                                        v72 := GenS4{F0: v5, F1: 42}
                                                                        fmt.Println(v2)
                                                                        _ = v72
                                                                    }
                                                                } else {
                                                                    var v73 int64 = 1
                                                                    _ = v73
                                                                }
                                                            case "":
                                                                if (!false) {
                                                                    var v74 int16 = v2
                                                                    var v75 uint = 255
                                                                    for fi76 := 0; fi76 < 11; fi76++ {
                                                                        var v77 int = (-(fi76))
                                                                        if len(v6) > 0 {
                                                                            v6[0] = fi76
                                                                        }
                                                                        if false {
                                                                            v78 := GenS2{F0: -698}
                                                                            if len(v6) > 0 {
                                                                                v6[0] = 0
                                                                            }
                                                                            v74 = 1000
                                                                            var v79 int8 = 1
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                v7 = false
                                                                            case "hello":
                                                                                _ = len(v6)
                                                                            }
                                                                            var v80 int = -9223372036854775808
                                                                            var v81 uint32 = v49
                                                                            _ = v78
                                                                            _ = v79
                                                                            _ = v80
                                                                            _ = v81
                                                                        }
                                                                        _ = fi76
                                                                        _ = v77
                                                                    }
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                            }
                                                            if v7 {
                                                                if ((^(v49)) <= 4294967295) {
                                                                    var v82 bool = v7
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        for fi83 := 0; fi83 < 20; fi83++ {
                                                                            var v84 int64 = 42
                                                                            _ = fi83
                                                                            _ = v84
                                                                        }
                                                                    case "":
                                                                        for fi85 := 0; fi85 < 4; fi85++ {
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            _ = fi85
                                                                        }
                                                                    default:
                                                                        v6 = append(v6, (^(fi51)))
                                                                    }
                                                                    _ = v82
                                                                } else {
                                                                    v86 := GenS7{F0: v5}
                                                                    _ = v86
                                                                }
                                                                for fi87 := 0; fi87 < 8; fi87++ {
                                                                    if v7 {
                                                                        var v88 int16 = v2
                                                                        _ = v88
                                                                    }
                                                                    fmt.Println(v7, fi87)
                                                                    var v89 string = "hello"
                                                                    fmt.Println(v2, v16, v3)
                                                                    switch v89 {
                                                                    case "hello":
                                                                        if len(v6) > 0 {
                                                                            v6[0] = fi51
                                                                        }
                                                                        v49 = (v49 / ((497) | 1))
                                                                    case "foo bar":
                                                                        if (!(!false)) {
                                                                            if v7 {
                                                                                var v90 int32 = 828
                                                                                _ = v90
                                                                            }
                                                                            if (!true) {
                                                                                v89 = v89
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        if v7 {
                                                                            var v91 int8 = -128
                                                                            switch fi51 {
                                                                            case 173:
                                                                                v92 := GenS5{F0: "a", F1: 42}
                                                                                _ = v92
                                                                            default:
                                                                                var v93 uint32 = uint32(v4)
                                                                                _ = v93
                                                                            }
                                                                            if v7 {
                                                                                switch (fi51 & fi87) {
                                                                                case 1:
                                                                                    _ = len(v6)
                                                                                }
                                                                            } else {
                                                                                v49 = v49
                                                                                var v94 int8 = v91
                                                                                var v95 int16 = v2
                                                                                _ = v94
                                                                                _ = v95
                                                                            }
                                                                            _ = v91
                                                                        }
                                                                    default:
                                                                        if v7 {
                                                                            fmt.Println(v7, v2)
                                                                            if v7 {
                                                                                v5 = (v5 & v5)
                                                                            }
                                                                            v3 += (^(v3))
                                                                        }
                                                                    }
                                                                    var v96 float64 = (float64(1e10) + v16)
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    v16 += v16
                                                                    v89 = "hello"
                                                                    var v97 float32 = 0.0
                                                                    _ = fi87
                                                                    _ = v89
                                                                    _ = v96
                                                                    _ = v97
                                                                }
                                                            }
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        }
                                                    }
                                                }
                                                var v98 int64 = 1
                                                _ = v52
                                                _ = v98
                                            } else {
                                                if (-16 >= 39) {
                                                    var v99 int8 = -1
                                                    _ = v99
                                                } else {
                                                    var v100 string = "foo bar"
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    var v101 int = (fi51 % ((fi51) | 1))
                                                    _ = v100
                                                    _ = v101
                                                }
                                                fmt.Println(v49, v3, v3)
                                                _ = len(v6)
                                            }
                                            if v7 {
                                                switch ("foo bar" + "hello") {
                                                case "a":
                                                    var v102 float32 = 1.5
                                                    _ = v102
                                                case "hello":
                                                    var v103 uint16 = v5
                                                    _ = v103
                                                default:
                                                    switch 127 {
                                                    case 128:
                                                        var v104 float64 = ((v16 - (1e10 + 1.5)) * v16)
                                                        _ = v104
                                                    case 127:
                                                        for fi105 := 0; fi105 < 17; fi105++ {
                                                            v6 = append(v6, (fi51 >> ((4294967295) & 63)))
                                                            v6 = append(v6, (fi105 >> ((915) & 63)))
                                                            v7 = v7
                                                            v2 = (v2 << ((42) & 63))
                                                            var v106 int8 = 1
                                                            _ = fi105
                                                            _ = v106
                                                        }
                                                    }
                                                }
                                                v107 := GenS8{F0: v5}
                                                v108 := GenS3{F0: fi51, F1: 9223372036854775807, F2: v49}
                                                _ = v107
                                                _ = v108
                                            } else {
                                                if (68 > 0) {
                                                    for fi109 := 0; fi109 < 1; fi109++ {
                                                        switch (fi51 / ((fi51) | 1)) {
                                                        case -1:
                                                            switch "" {
                                                            case "foo bar":
                                                                v110 := GenS1{F0: v2, F1: v1}
                                                                _ = v110
                                                            case "hello":
                                                                fmt.Println(v2, fi51, v5)
                                                            default:
                                                                var v111 bool = (!(!(float32((v16 + float64(0.0))) < (1e10 / 3.14159))))
                                                                _ = v111
                                                            }
                                                        case 0:
                                                            var v112 int = 128
                                                            _ = v112
                                                        }
                                                        v113 := GenS7{F0: 443}
                                                        v16 += v16
                                                        _ = fi109
                                                        _ = v113
                                                    }
                                                    v114 := GenS1{F0: v2, F1: (v3 / ((128) | 1))}
                                                    _ = v114
                                                }
                                                v6 = append(v6, 828)
                                                switch "foo bar" {
                                                case "a":
                                                    switch 0 {
                                                    case -1:
                                                        if ((v49 - (v49 * (^(v49)))) < (^(v49))) {
                                                            switch (fi51 & (fi51 ^ (fi51 ^ fi51))) {
                                                            case 4294967295:
                                                                v4 -= (v4 + (v4 >> ((2) & 63)))
                                                            case -2147483648:
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    fmt.Println(v2)
                                                                default:
                                                                    if (v16 <= ((v16 - v16) / v16)) {
                                                                        v4 -= (^(v4))
                                                                    } else {
                                                                        fmt.Println(v2)
                                                                        if ((v49 ^ (^(v49))) < (^(v49))) {
                                                                            v2 += 42
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                if v7 {
                                                                                    fmt.Println(fi51)
                                                                                    var v115 int8 = 0
                                                                                    if true {
                                                                                        v116 := GenS9{F0: "foo bar", F1: true, F2: (^(v4))}
                                                                                        _ = v116
                                                                                    }
                                                                                    _ = v115
                                                                                }
                                                                            case "":
                                                                                v117 := GenS2{F0: 1}
                                                                                _ = v117
                                                                            }
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                        }
                                                                        var v118 bool = true
                                                                        _ = v118
                                                                    }
                                                                }
                                                            default:
                                                                switch ("foo bar" + "hello") {
                                                                case "":
                                                                    var v119 float32 = float32(float64(float32(v16)))
                                                                    _ = v119
                                                                case "a":
                                                                    if v7 {
                                                                        switch ("a" + "a") {
                                                                        case "foo bar":
                                                                            v6 = append(v6, fi51)
                                                                        }
                                                                        v6 = append(v6, (fi51 * fi51))
                                                                        v3 -= (v1 % (((^(v1))) | 1))
                                                                    } else {
                                                                        v120 := GenS10{F0: 100, F1: 127, F2: v49}
                                                                        fmt.Println(fi51, v1, v16)
                                                                        for fi121 := 0; fi121 < 3; fi121++ {
                                                                            if v7 {
                                                                                var v122 float64 = 0.0
                                                                                var v123 int = 509
                                                                                _ = v122
                                                                                _ = v123
                                                                            }
                                                                            var v124 float32 = 1e10
                                                                            _ = fi121
                                                                            _ = v124
                                                                        }
                                                                        _ = v120
                                                                    }
                                                                }
                                                            }
                                                            if len(v6) > 0 {
                                                                v6[0] = -2147483648
                                                            }
                                                        }
                                                    case 0:
                                                        fmt.Println(fi51, v49, v3)
                                                    default:
                                                        fmt.Println(fi51)
                                                    }
                                                case "":
                                                    for fi125 := 0; fi125 < 9; fi125++ {
                                                        v6 = append(v6, -791)
                                                        var v126 uint16 = 122
                                                        _ = fi125
                                                        _ = v126
                                                    }
                                                case "\n\t":
                                                    var v127 int8 = -128
                                                    _ = v127
                                                default:
                                                    v7 = ("\n\t" <= ("a" + ""))
                                                }
                                            }
                                            _ = fi51
                                        }
                                    }
                                } else {
                                    if true {
                                        v128 := GenS7{F0: uint16(v4)}
                                        v1 = (^(v3))
                                        _ = v128
                                    }
                                    var v129 int8 = 1
                                    v130 := GenS11{F0: (^(v1)), F1: 72}
                                    _ = v129
                                    _ = v130
                                }
                                _ = v49
                            }
                            var v131 int32 = 0
                            _ = v131
                            _ = v37
                        } else {
                            v2 -= 736
                        }
                    }
                default:
                    v6 = append(v6, 127)
                }
                _ = v16
            } else {
                var v132 float64 = 3.14159
                _ = v132
            }
            for fi133 := 0; fi133 < 4; fi133++ {
                for fi134 := 0; fi134 < 5; fi134++ {
                    if true {
                        var v135 int32 = 2147483647
                        _ = v135
                    } else {
                        var v136 float64 = 3.14159
                        if ("" <= ("" + "foo bar")) {
                            var v137 float32 = 3.14159
                            _ = v137
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            fmt.Println(v3)
                            var v138 uint = 42
                            if false {
                                if false {
                                    v7 = v7
                                    v2 = (v2 >> ((v138) & 63))
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                            } else {
                                v2 += (v2 << ((18446744073709551615) & 63))
                                switch "foo bar" {
                                case "hello":
                                    var v139 int8 = 42
                                    if v7 {
                                        fmt.Println(v138, v139)
                                    }
                                    _ = v139
                                case "\n\t":
                                    fmt.Println(v1)
                                case "":
                                    v1 += (^(v3))
                                }
                                var v140 int64 = 1
                                _ = v140
                            }
                            v136 = (float64((1e10 / 1.0)) / 1.0)
                            _ = len(v6)
                            switch "" {
                            case "":
                                if false {
                                    fmt.Println(fi133)
                                    fmt.Println(v5, v1, v3)
                                }
                            }
                            v3 -= uint8(v138)
                            v141 := GenS7{F0: v5}
                            v3 = v3
                            switch 350 {
                            case -1:
                                v7 = false
                            default:
                                v4 -= (v4 << ((v138) & 63))
                            }
                            fmt.Println(v138)
                            var v142 uint64 = (^(v4))
                            v3 = 255
                            var v143 float32 = (float32(1.5) + (float32(v136) + 0.1))
                            _ = v138
                            _ = v141
                            _ = v142
                            _ = v143
                        }
                        v4 -= 42
                        v2 = (v2 * (^(v2)))
                        _ = v136
                    }
                    if v7 {
                        if v7 {
                            var v144 uint64 = (v4 / ((42) | 1))
                            _ = v144
                        }
                    } else {
                        v2 = (v2 % ((v2) | 1))
                        v145 := GenS10{F0: 0, F1: 127, F2: 280}
                        v3 = uint8(v5)
                        var v146 int = int(v3)
                        v5 = (v5 ^ v5)
                        _ = v145
                        _ = v146
                    }
                    if false {
                        var v147 bool = (v7 || ((-(v2)) >= int16(fi133)))
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = v147
                    } else {
                        var v148 int16 = 0
                        var v149 uint16 = v5
                        var v150 int = fi133
                        v151 := GenS3{F0: (fi134 % (((-(fi133))) | 1)), F1: 334, F2: 1}
                        _ = v148
                        _ = v149
                        _ = v150
                        _ = v151
                    }
                    _ = fi134
                }
                v7 = v7
                _ = fi133
            }
            var v152 string = "\n\t"
            _ = v152
        }
        if v7 {
            var v153 int16 = (v2 + (^(v2)))
            v153 -= int16(v1)
            _ = len(v6)
            _ = v153
        } else {
            var v154 float64 = 0.1
            if v7 {
                var v155 int8 = 86
                _ = v155
            } else {
                var v156 float64 = 1e10
                switch "a" {
                case "\n\t":
                    var v157 string = ((("hello" + "foo bar") + ("" + "a")) + "\n\t")
                    _ = v157
                case "hello":
                    v7 = ((!(((-2.25 > 1.5) && (-81 != -109)) && true)) || v7)
                }
                for fi158 := 0; fi158 < 4; fi158++ {
                    if false {
                        if false {
                            var v159 int16 = v2
                            for fi160 := 0; fi160 < 4; fi160++ {
                                v159 = v2
                                _ = fi160
                            }
                            _ = v159
                        }
                        switch ("\n\t" + (("" + "") + ("\n\t" + "hello"))) {
                        case "hello":
                            var v161 uint8 = v3
                            _ = v161
                        case "":
                            for fi162 := 0; fi162 < 3; fi162++ {
                                var v163 int32 = 0
                                _ = fi162
                                _ = v163
                            }
                        case "a":
                            var v164 uint64 = v4
                            _ = v164
                        }
                        var v165 int32 = 0
                        _ = v165
                    }
                    for fi166 := 0; fi166 < 4; fi166++ {
                        v154 = (0.0 + v156)
                        v5 -= (v5 - (v5 - v5))
                        _ = fi166
                    }
                    for fi167 := 0; fi167 < 16; fi167++ {
                        v4 = 530
                        fmt.Println(v5)
                        v3 = v1
                        if false {
                            if true {
                                var v168 uint32 = 42
                                switch (fi167 << ((42) & 63)) {
                                case 946:
                                    var v169 float32 = (-2.25 - 0.1)
                                    _ = v169
                                case 0:
                                    fmt.Println(fi158)
                                case 127:
                                    v170 := GenS9{F0: "\n\t", F1: (!true), F2: v4}
                                    _ = v170
                                default:
                                    var v171 uint = 1
                                    _ = v171
                                }
                                _ = v168
                            } else {
                                v156 = float64((3.14159 * ((1e10 - 1.5) - 3.14159)))
                                v154 = 3.14159
                                v3 += (^(v1))
                            }
                            var v172 uint64 = v4
                            switch "" {
                            case "\n\t":
                                var v173 uint16 = v5
                                _ = v173
                            }
                            _ = v172
                        }
                        _ = fi167
                    }
                    _ = fi158
                }
                _ = v156
            }
            switch "a" {
            case "\n\t":
                if v7 {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    for fi174 := 0; fi174 < 3; fi174++ {
                        v175 := GenS12{F0: 18446744073709551615, F1: (!(!v7)), F2: 2147483647}
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi174
                        _ = v175
                    }
                    v1 = (v3 ^ v1)
                }
            case "foo bar":
                var v176 int16 = v2
                _ = v176
            case "a":
                var v177 int32 = -2147483648
                _ = v177
            }
            _ = v154
        }
        v178 := GenS12{F0: 886, F1: false, F2: 644}
        _ = v178
    }
    if v7 {
        fmt.Println(v7)
        for fi179 := 0; fi179 < 11; fi179++ {
            switch fi179 {
            case 42:
                if len(v6) > 0 {
                    v6[0] = fi179
                }
            }
            if (1.5 != ((float64(-2.25) * 1.5) - (float64((3.14159 + 0.1)) / 1.0))) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi179
        }
    }
    fmt.Println(v3, v2, v1)
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
