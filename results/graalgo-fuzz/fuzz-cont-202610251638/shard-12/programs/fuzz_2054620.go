package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint16; F2 bool }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 int8; F1 int32; F2 int8 }
type GenS7 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS8 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS9 struct { F0 uint8; F1 uint8 }
type GenS10 struct { F0 int32; F1 uint }
type GenS11 struct { F0 int8; F1 uint64 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint16; F1 int16 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 uint64; F1 int8 }
type GenS18 struct { F0 int; F1 uint8 }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 int64; F1 int8 }
type GenS22 struct { F0 string; F1 int; F2 int32 }

func main() {
    var v1 int64 = -310
    v2 := []int{-918}
    v3 := make(map[uint32]string)
    var v4 int8 = 1
    var v5 uint = 65535
    var v6 int8 = (-(v4))
    if (!(0.0 == 1e10)) {
        if false {
            switch "foo bar" {
            case "\n\t":
                if (!false) {
                    _ = len(v2)
                } else {
                    for fi7 := 0; fi7 < 13; fi7++ {
                        switch fi7 {
                        case -353:
                            v1 += (v1 << (((v5 | 255)) & 63))
                        case -2147483648:
                            for fi8 := 0; fi8 < 14; fi8++ {
                                v6 = -114
                                var v9 uint16 = 480
                                var v10 uint8 = 127
                                _ = fi8
                                _ = v10
                                _ = v9
                            }
                        case 32768:
                            switch "\n\t" {
                            case "foo bar":
                                var v11 uint8 = 255
                                _ = v11
                            case "a":
                                v2 = append(v2, fi7)
                            }
                        default:
                            v5 = v5
                        }
                        _ = fi7
                    }
                    v4 = (v6 << (((v5 / ((2) | 1))) & 63))
                    if (v5 == v5) {
                        for fi12 := 0; fi12 < 12; fi12++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi12
                        }
                        if (false && ((!true) && (false || (false || (3.14159 != ((((0.0 / 1e10) - (-2.25 - 1.5)) / (0.0 - 3.14159)) - float64(3.14159))))))) {
                            v4 = 42
                            var v13 int8 = (v4 - (v6 + v6))
                            _ = v13
                        } else {
                            fmt.Println(v1, v5)
                            var v14 bool = (!(!true))
                            fmt.Println(v5, v14, v1)
                            _ = v14
                        }
                    }
                }
            default:
                switch 128 {
                case -295:
                    if len(v2) > 0 {
                        v2[0] = 2147483647
                    }
                case 127:
                    v5 = (v5 << (((v5 << ((v5) & 63))) & 63))
                default:
                    v6 -= v4
                }
            }
            switch 2147483647 {
            case 127:
                if false {
                    if true {
                        if (73 == (v6 + 42)) {
                            var v15 uint8 = 255
                            var v16 uint32 = 139
                            switch ("" + ("\n\t" + "foo bar")) {
                            case "":
                                if (v16 >= v16) {
                                    var v17 uint64 = 0
                                    var v18 uint32 = (^(v16))
                                    for fi19 := 0; fi19 < 6; fi19++ {
                                        v20 := GenS1{F0: v5, F1: -1}
                                        if false {
                                            if true {
                                                switch "hello" {
                                                case "a":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                var v21 uint32 = (^(v18))
                                                var v22 int32 = 0
                                                _ = v21
                                                _ = v22
                                            } else {
                                                if ((fi19 < (fi19 | fi19)) && (1 < int32(v5))) {
                                                    var v23 int64 = -1
                                                    v18 = 508
                                                    for fi24 := 0; fi24 < 14; fi24++ {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            if (!(-1 == -1)) {
                                                                v25 := GenS1{F0: (v5 & (v5 + uint(v18))), F1: (v1 - 9223372036854775807)}
                                                                _ = v25
                                                            } else {
                                                                fmt.Println(v17, v1, v23)
                                                                var v26 float64 = (0.0 * (float64((0.0 / (3.14159 - 1.5))) - (0.1 / 1.5)))
                                                                _ = v26
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(fi24, v16, v5)
                                                            switch ("\n\t" + "") {
                                                            case "\n\t":
                                                                if ((v15 << (((v5 >> ((v5) & 63))) & 63)) >= 1) {
                                                                    v23 = v23
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v18 = (v18 << (((^(v5))) & 63))
                                                                }
                                                            case "hello":
                                                                v27 := GenS1{F0: (v5 % ((v5) | 1)), F1: 9223372036854775807}
                                                                _ = v27
                                                            }
                                                            v23 -= (-(v1))
                                                        }
                                                        v6 = (-(v6))
                                                        _ = fi24
                                                    }
                                                    _ = v23
                                                } else {
                                                    v5 += 686
                                                    if true {
                                                        v17 = (^(v17))
                                                        for fi28 := 0; fi28 < 5; fi28++ {
                                                            v29 := GenS1{F0: 345, F1: int64(fi19)}
                                                            _ = len(v3)
                                                            fmt.Println(v16, v18, v17)
                                                            _ = fi28
                                                            _ = v29
                                                        }
                                                        var v30 uint8 = v15
                                                        _ = v30
                                                    }
                                                    v31 := GenS2{F0: uint(v17)}
                                                    _ = v31
                                                }
                                                var v32 string = (("a" + "a") + (("foo bar" + "a") + (("\n\t" + "hello") + (("a" + ("a" + "\n\t")) + "\n\t"))))
                                                switch "a" {
                                                case "":
                                                    var v33 uint = (v5 * v5)
                                                    _ = v33
                                                case "hello":
                                                    for fi34 := 0; fi34 < 9; fi34++ {
                                                        var v35 uint32 = (v18 + (v18 % ((v16) | 1)))
                                                        v35 -= (v35 ^ 2)
                                                        _ = fi34
                                                        _ = v35
                                                    }
                                                }
                                                _ = v32
                                            }
                                        }
                                        _ = fi19
                                        _ = v20
                                    }
                                    _ = v17
                                    _ = v18
                                }
                            case "a":
                                v36 := GenS1{F0: 2, F1: 9223372036854775807}
                                _ = v36
                            }
                            _ = v15
                            _ = v16
                        }
                        for fi37 := 0; fi37 < 11; fi37++ {
                            if true {
                                v5 = (v5 - (v5 / (((v5 + 890)) | 1)))
                                v5 = (v5 * v5)
                            } else {
                                var v38 uint16 = 0
                                _ = v38
                            }
                            if (!true) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v4 += 100
                                var v39 int32 = 2147483647
                                _ = v39
                            } else {
                                v1 += -9223372036854775808
                                _ = len(v3)
                            }
                            _ = fi37
                        }
                    }
                } else {
                    if true {
                        var v40 int64 = v1
                        var v41 int64 = v1
                        fmt.Println(v4)
                        _ = v40
                        _ = v41
                    }
                }
            case -9223372036854775808:
                v3[2147483648] = "foo bar"
            default:
                for fi42 := 0; fi42 < 6; fi42++ {
                    var v43 uint16 = 692
                    _ = fi42
                    _ = v43
                }
            }
            if ((^(v5)) >= (v5 << ((v5) & 63))) {
                v4 = 84
            } else {
                var v44 uint16 = 255
                _ = v44
            }
        } else {
            if true {
                var v45 uint8 = 127
                for fi46 := 0; fi46 < 11; fi46++ {
                    _ = len(v2)
                    for fi47 := 0; fi47 < 14; fi47++ {
                        var v48 uint8 = (v45 % ((v45) | 1))
                        _ = fi47
                        _ = v48
                    }
                    _ = fi46
                }
                _ = v45
            } else {
                for fi49 := 0; fi49 < 1; fi49++ {
                    v4 += 127
                    v6 = (v6 * v4)
                    _ = fi49
                }
                fmt.Println(v5)
            }
            if (!(1 >= -32768)) {
                switch -754 {
                case 4294967295:
                    v5 -= v5
                }
                v50 := GenS1{F0: v5, F1: (v1 % (((v1 >> (((v5 + v5)) & 63))) | 1))}
                if false {
                    for fi51 := 0; fi51 < 13; fi51++ {
                        for fi52 := 0; fi52 < 4; fi52++ {
                            var v53 int8 = (v6 ^ v4)
                            v53 += (v4 * 127)
                            v54 := GenS1{F0: (^(v5)), F1: v1}
                            if ((32767 < 127) || ((0 < 42) && true)) {
                                if false {
                                    v4 = v4
                                    if false {
                                        v6 += 0
                                        switch (fi51 - fi52) {
                                        case 806:
                                            fmt.Println(fi52, fi51)
                                        case -9223372036854775808:
                                            if ("a" > ("foo bar" + ("a" + "foo bar"))) {
                                                fmt.Println(v53)
                                                v2 = append(v2, (-(fi52)))
                                            } else {
                                                var v55 uint64 = 974
                                                _ = v55
                                            }
                                        }
                                        switch (("hello" + "") + ("foo bar" + (("\n\t" + "\n\t") + ("foo bar" + "hello")))) {
                                        case "\n\t":
                                            v1 = (^(v1))
                                        case "a":
                                            v6 = 100
                                        default:
                                            v5 = v5
                                        }
                                    }
                                    v1 = v1
                                } else {
                                    v3[0] = "\n\t"
                                }
                            }
                            var v56 uint8 = 128
                            _ = len(v3)
                            switch -1 {
                            case -2147483648:
                                v6 -= (v4 * v6)
                            }
                            _ = fi52
                            _ = v53
                            _ = v54
                            _ = v56
                        }
                        _ = fi51
                    }
                    if (374 != 1000) {
                        v4 += v4
                    }
                }
                _ = v50
            } else {
                v5 -= v5
            }
        }
    } else {
        v57 := v3[2147483648]
        _ = v57
    }
    switch "hello" {
    case "foo bar":
        v5 = (v5 / ((v5) | 1))
    case "\n\t":
        v5 = 526
    default:
        v58 := GenS1{F0: v5, F1: (v1 >> ((v5) & 63))}
        _ = v58
    }
    switch ("\n\t" + "a") {
    case "hello":
        fmt.Println(v1, v4, v6)
    case "foo bar":
        switch ("a" + (("" + "\n\t") + ("a" + "\n\t"))) {
        case "\n\t":
            if (!true) {
                var v59 float64 = float64((0.1 + 1e10))
                for fi60 := 0; fi60 < 14; fi60++ {
                    v61 := GenS1{F0: (v5 % ((v5) | 1)), F1: (^(v1))}
                    _ = fi60
                    _ = v61
                }
                _ = v59
            } else {
                for fi62 := 0; fi62 < 2; fi62++ {
                    var v63 int16 = 32767
                    if true {
                        if false {
                            v6 = 100
                            switch (("" + "hello") + "a") {
                            case "a":
                                if (42 <= 1) {
                                    v2 = append(v2, (-(fi62)))
                                }
                            case "\n\t":
                                var v64 int8 = v4
                                _ = v64
                            case "hello":
                                switch 725 {
                                case -67:
                                    if true {
                                        v65 := GenS1{F0: (v5 ^ (v5 * 18446744073709551615)), F1: 0}
                                        var v66 float32 = 3.14159
                                        _ = v65
                                        _ = v66
                                    } else {
                                        if true {
                                            fmt.Println(v1, v4, fi62)
                                            switch ("" + ("foo bar" + "hello")) {
                                            case "a":
                                                var v67 int16 = (^(v63))
                                                _ = v67
                                            case "foo bar":
                                                for fi68 := 0; fi68 < 11; fi68++ {
                                                    v1 = 42
                                                    v5 = (v5 % (((v5 ^ 42)) | 1))
                                                    _ = len(v3)
                                                    var v69 bool = (float64(0.1) <= (((-2.25 * -2.25) / 3.14159) - 1e10))
                                                    var v70 uint = (v5 ^ v5)
                                                    v5 = v70
                                                    switch "" {
                                                    case "a":
                                                        if v69 {
                                                            switch -713 {
                                                            case -412:
                                                                if (173 == 1) {
                                                                    v6 = (v4 ^ (v6 | (v6 - (v6 | (-(v4))))))
                                                                    v71 := GenS3{F0: (v1 << (((v70 >> (((v70 & (v5 << ((v70) & 63)))) & 63))) & 63)), F1: 32767, F2: v69}
                                                                    _ = v71
                                                                }
                                                                v4 = int8(fi62)
                                                            case 127:
                                                                switch ("hello" + "hello") {
                                                                case "foo bar":
                                                                    v72 := GenS1{F0: (v5 >> ((v70) & 63)), F1: v1}
                                                                    _ = v72
                                                                }
                                                            default:
                                                                _ = len(v3)
                                                            }
                                                            v73 := GenS4{F0: 1, F1: 140, F2: 0}
                                                            v1 = v1
                                                            if v69 {
                                                                v74 := GenS5{F0: false, F1: true}
                                                                _ = v74
                                                            } else {
                                                                v75 := GenS6{F0: v4, F1: 0, F2: (v6 - (v4 >> ((838) & 63)))}
                                                                if v69 {
                                                                    v63 = -307
                                                                }
                                                                _ = v75
                                                            }
                                                            _ = v73
                                                        } else {
                                                            v6 = v6
                                                            fmt.Println(v63, v6, v6)
                                                        }
                                                    case "":
                                                        v76 := GenS7{F0: 2, F1: 1, F2: (v6 * v4)}
                                                        if (fi68 <= (fi62 >> ((v70) & 63))) {
                                                            v69 = v69
                                                            v69 = v69
                                                            var v77 bool = v69
                                                            _ = v77
                                                        }
                                                        if (false || (!(18446744073709551615 <= 42))) {
                                                            v69 = (!v69)
                                                            var v78 int = (fi68 ^ (fi68 >> (((v70 / ((v70) | 1))) & 63)))
                                                            _ = v78
                                                        } else {
                                                            v6 = v6
                                                            switch ("" + "\n\t") {
                                                            case "hello":
                                                                delete(v3, 42)
                                                            case "":
                                                                var v79 float32 = -2.25
                                                                v80 := GenS8{F0: 0, F1: 4294967295, F2: v1}
                                                                if true {
                                                                    switch fi62 {
                                                                    case 9223372036854775807:
                                                                        var v81 int = (fi62 / ((32768) | 1))
                                                                        _ = v81
                                                                    case 0:
                                                                        fmt.Println(v70, v5, v4)
                                                                        v82 := GenS9{F0: 215, F1: 42}
                                                                        _ = v82
                                                                    default:
                                                                        v69 = v69
                                                                    }
                                                                } else {
                                                                    var v83 uint64 = 18446744073709551615
                                                                    v63 -= int16(v5)
                                                                    _ = v83
                                                                }
                                                                _ = v79
                                                                _ = v80
                                                            default:
                                                                var v84 int = 128
                                                                _ = v84
                                                            }
                                                            v85 := GenS10{F0: -264, F1: 18446744073709551615}
                                                            _ = v85
                                                        }
                                                        _ = v76
                                                    }
                                                    v70 = (v5 + (v70 * (v70 * v5)))
                                                    _ = fi68
                                                    _ = v69
                                                    _ = v70
                                                }
                                            case "":
                                                if len(v2) > 0 {
                                                    v2[0] = fi62
                                                }
                                            }
                                        } else {
                                            for fi86 := 0; fi86 < 7; fi86++ {
                                                v2 = append(v2, (-(fi62)))
                                                var v87 uint32 = 916
                                                if false {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v5 -= (v5 & (v5 / (((v5 >> (((v5 & v5)) & 63))) | 1)))
                                                    default:
                                                        var v88 uint8 = 128
                                                        _ = v88
                                                    }
                                                } else {
                                                    var v89 string = ""
                                                    _ = v89
                                                }
                                                v90 := v3[v87]
                                                _ = fi86
                                                _ = v87
                                                _ = v90
                                            }
                                            var v91 string = ("hello" + "foo bar")
                                            v5 += (^(v5))
                                            _ = v91
                                        }
                                        v3[0] = (("a" + "a") + "a")
                                        if true {
                                            if (!((((v5 * v5) <= (v5 << (((v5 ^ (^(v5)))) & 63))) && false) || false)) {
                                                for fi92 := 0; fi92 < 18; fi92++ {
                                                    v2 = append(v2, fi62)
                                                    var v93 string = ("foo bar" + "foo bar")
                                                    _ = fi92
                                                    _ = v93
                                                }
                                                if false {
                                                    v63 = v63
                                                }
                                            }
                                        } else {
                                            var v94 uint32 = 2147483647
                                            var v95 uint16 = 652
                                            _ = v94
                                            _ = v95
                                        }
                                    }
                                case -1:
                                    for fi96 := 0; fi96 < 9; fi96++ {
                                        switch fi96 {
                                        case 980:
                                            var v97 uint32 = 877
                                            _ = v97
                                        case 9223372036854775807:
                                            var v98 int = (fi96 * (fi96 + 127))
                                            _ = v98
                                        case 592:
                                            switch "" {
                                            case "foo bar":
                                                var v99 float32 = float32(-2.25)
                                                _ = v99
                                            case "hello":
                                                v5 -= (v5 << ((255) & 63))
                                            case "\n\t":
                                                switch int(v63) {
                                                case 64:
                                                    switch "hello" {
                                                    case "a":
                                                        v1 = v1
                                                    case "foo bar":
                                                        fmt.Println(fi96)
                                                    }
                                                default:
                                                    _ = len(v3)
                                                }
                                            }
                                        default:
                                            v100 := GenS11{F0: v4, F1: 563}
                                            _ = v100
                                        }
                                        if (v5 <= 713) {
                                            if len(v2) > 0 {
                                                v2[0] = (-(fi62))
                                            }
                                            var v101 uint = 667
                                            _ = v101
                                        }
                                        var v102 int16 = -275
                                        _ = fi96
                                        _ = v102
                                    }
                                }
                            default:
                                v63 = (-(v63))
                            }
                            fmt.Println(v6)
                        } else {
                            if (0 > 65535) {
                                if (-2.25 < (-2.25 - (0.1 - ((-2.25 - (-2.25 * 1e10)) / 1e10)))) {
                                    var v103 uint64 = 18446744073709551615
                                    v5 = v5
                                    _ = v103
                                }
                                for fi104 := 0; fi104 < 20; fi104++ {
                                    if (42 != (v5 / ((v5) | 1))) {
                                        var v105 float64 = float64(0.0)
                                        var v106 int32 = int32(v5)
                                        if (!(!(42 > 520))) {
                                            v63 -= v63
                                            var v107 int16 = (v63 << (((v5 >> ((0) & 63))) & 63))
                                            var v108 uint64 = 1
                                            if (1.5 != v105) {
                                                fmt.Println(v5, v5, v108)
                                            }
                                            _ = v107
                                            _ = v108
                                        } else {
                                            var v109 uint64 = 18446744073709551615
                                            _ = v109
                                        }
                                        _ = v105
                                        _ = v106
                                    } else {
                                        var v110 float32 = (float32((3.14159 - 1.5)) - 0.1)
                                        var v111 int16 = (v63 - v63)
                                        _ = v110
                                        _ = v111
                                    }
                                    if false {
                                        v4 = 42
                                    }
                                    v112 := v3[65535]
                                    _ = fi104
                                    _ = v112
                                }
                                for fi113 := 0; fi113 < 16; fi113++ {
                                    v114 := GenS5{F0: (1.5 < (((1.5 / 0.1) + 3.14159) / (1.5 + 1e10))), F1: (!true)}
                                    _ = fi113
                                    _ = v114
                                }
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        for fi115 := 0; fi115 < 2; fi115++ {
                            v63 = (v63 - (v63 ^ (v63 * -11)))
                            fmt.Println(fi62, v4)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi115
                        }
                    } else {
                        if (false && true) {
                            if (1 <= 65535) {
                                var v116 uint32 = 553
                                switch fi62 {
                                case -2147483648:
                                    for fi117 := 0; fi117 < 15; fi117++ {
                                        v6 -= v6
                                        _ = fi117
                                    }
                                case 32768:
                                    var v118 int = (fi62 % (((fi62 & (fi62 | (fi62 | (fi62 - (fi62 / ((fi62) | 1))))))) | 1))
                                    _ = v118
                                }
                                _ = v116
                            } else {
                                switch 988 {
                                case 2147483647:
                                    v63 = (v63 << (((v5 ^ 368)) & 63))
                                case 0:
                                    v3[4294967295] = ""
                                case -1:
                                    var v119 uint32 = 0
                                    _ = v119
                                default:
                                    v120 := GenS12{F0: v1}
                                    _ = v120
                                }
                                switch fi62 {
                                case 57:
                                    switch ("a" + ("" + "")) {
                                    case "\n\t":
                                        v1 += -1
                                    default:
                                        if (true && (3.14159 == 0.0)) {
                                            if (((("\n\t" + "foo bar") + "hello") + ("hello" + ("" + "hello"))) > ("hello" + (("hello" + "a") + "a"))) {
                                                v6 -= (v6 | v6)
                                                if (2 > 608) {
                                                    var v121 uint8 = 2
                                                    _ = v121
                                                } else {
                                                    v122 := GenS13{F0: -2147483648}
                                                    var v123 bool = ("a" >= "foo bar")
                                                    _ = v122
                                                    _ = v123
                                                }
                                                if true {
                                                    var v124 int16 = v63
                                                    v124 = (v124 + v124)
                                                    v125 := GenS11{F0: (v6 % ((v6) | 1)), F1: 328}
                                                    _ = v124
                                                    _ = v125
                                                } else {
                                                    for fi126 := 0; fi126 < 10; fi126++ {
                                                        v127 := GenS14{F0: (0 <= 499)}
                                                        _ = fi126
                                                        _ = v127
                                                    }
                                                }
                                            }
                                            switch 2147483647 {
                                            case 1:
                                                for fi128 := 0; fi128 < 20; fi128++ {
                                                    var v129 float32 = (1e10 + 0.1)
                                                    v4 -= v6
                                                    _ = fi128
                                                    _ = v129
                                                }
                                            case -2147483648:
                                                v63 -= v63
                                            case 32767:
                                                for fi130 := 0; fi130 < 4; fi130++ {
                                                    fmt.Println(v4, v4, v63)
                                                    v4 -= 98
                                                    var v131 int = fi62
                                                    _ = fi130
                                                    _ = v131
                                                }
                                            default:
                                                var v132 uint8 = 128
                                                _ = v132
                                            }
                                        } else {
                                            if ((0.1 > ((((float64(-2.25) / 1.0) * -2.25) - 0.1) + (((((3.14159 * 0.1) / (3.14159 / 1.5)) / 0.1) / (3.14159 + ((1.5 + float64(((float32(1.5) + 1.5) - (3.14159 - ((float32(0.1) - 0.0) / 1.5))))) - 3.14159))) - 0.1))) || false) {
                                                v133 := GenS2{F0: (v5 / (((v5 << ((v5) & 63))) | 1))}
                                                v134 := v3[273]
                                                _ = v133
                                                _ = v134
                                            }
                                            if ("" < "hello") {
                                                for fi135 := 0; fi135 < 20; fi135++ {
                                                    v1 += v1
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v136 := GenS11{F0: (v4 ^ v4), F1: 2}
                                                    _ = fi135
                                                    _ = v136
                                                }
                                            }
                                        }
                                    }
                                default:
                                    for fi137 := 0; fi137 < 7; fi137++ {
                                        v138 := GenS15{F0: 0, F1: 1000}
                                        v139 := GenS16{F0: 0}
                                        var v140 int16 = -1
                                        _ = fi137
                                        _ = v138
                                        _ = v139
                                        _ = v140
                                    }
                                }
                            }
                            var v141 uint = (^(v5))
                            v6 += 1
                            _ = v141
                        } else {
                            fmt.Println(v5)
                        }
                    }
                    var v142 int8 = -1
                    _ = fi62
                    _ = v142
                    _ = v63
                }
                var v143 uint16 = 42
                _ = v143
            }
        case "foo bar":
            delete(v3, 98)
        default:
            for fi144 := 0; fi144 < 13; fi144++ {
                fmt.Println(fi144)
                _ = fi144
            }
        }
    default:
        v145 := GenS17{F0: 317, F1: v6}
        _ = v145
    }
    fmt.Println(v4)
    fmt.Println(v1)
    for fi146 := 0; fi146 < 20; fi146++ {
        if (v6 != v6) {
            var v147 uint32 = 2147483647
            for fi148 := 0; fi148 < 12; fi148++ {
                switch "\n\t" {
                case "a":
                    if false {
                        switch (fi148 >> (((v5 / ((42) | 1))) & 63)) {
                        case 32768:
                            v149 := GenS18{F0: (fi146 + (fi146 | fi148)), F1: 1}
                            _ = v149
                        case -9223372036854775808:
                            v1 -= v1
                        case -226:
                            if true {
                                if (!(true || (127 < (v6 % ((v6) | 1))))) {
                                    if (1 > 233) {
                                        var v150 uint = 255
                                        v151 := GenS19{F0: 0}
                                        _ = v150
                                        _ = v151
                                    }
                                    if (1 > 1) {
                                        if true {
                                            var v152 int8 = v6
                                            v153 := GenS17{F0: 2, F1: (v6 % (((-(v6))) | 1))}
                                            _ = v152
                                            _ = v153
                                        }
                                        if true {
                                            v154 := GenS20{F0: v147}
                                            fmt.Println(v6, v6)
                                            if false {
                                                if (("\n\t" == "a") && true) {
                                                    if true {
                                                        _ = len(v3)
                                                        var v155 int32 = 533
                                                        _ = v155
                                                    }
                                                }
                                                v156 := GenS21{F0: (-(v1)), F1: (v4 - (^(v4)))}
                                                _ = v156
                                            }
                                            _ = v154
                                        } else {
                                            var v157 int16 = 601
                                            if (!false) {
                                                switch fi148 {
                                                case 604:
                                                    fmt.Println(v157)
                                                case -608:
                                                    v157 += (^(v157))
                                                case -1:
                                                    if ((32 <= 1) && false) {
                                                        if false {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            fmt.Println(v157)
                                                        } else {
                                                            var v158 int64 = (-(v1))
                                                            v159 := GenS4{F0: 760, F1: 41, F2: v5}
                                                            fmt.Println(v5, v4, v157)
                                                            _ = v158
                                                            _ = v159
                                                        }
                                                        var v160 uint = v5
                                                        fmt.Println(v147, v6, fi146)
                                                        _ = v160
                                                    }
                                                }
                                                v161 := GenS17{F0: 130, F1: (v4 * v6)}
                                                fmt.Println(v4, v6)
                                                _ = v161
                                            } else {
                                                v2 = append(v2, fi146)
                                                var v162 uint64 = 1
                                                v162 = (v162 << ((v5) & 63))
                                                _ = v162
                                            }
                                            if false {
                                                var v163 int16 = (v157 + v157)
                                                _ = v163
                                            } else {
                                                v157 = -1
                                            }
                                            _ = v157
                                        }
                                    } else {
                                        if ("hello" <= "\n\t") {
                                            if false {
                                                if false {
                                                    if (!(float64(float32((1e10 * (float64((float32(3.14159) * float32(0.0))) - (((float64(1e10) + 1.5) * (1.5 + 1.5)) * (0.1 - 0.1)))))) <= (0.1 + 1.5))) {
                                                        if ((true || true) && (true && (!(!false)))) {
                                                            fmt.Println(fi148, v1)
                                                            var v164 string = "foo bar"
                                                            _ = v164
                                                        }
                                                        var v165 uint16 = 0
                                                        _ = v165
                                                    }
                                                    v166 := v3[416]
                                                    var v167 int = fi146
                                                    _ = v166
                                                    _ = v167
                                                }
                                                if ("hello" >= "") {
                                                    switch fi146 {
                                                    case 4294967295:
                                                        v168 := GenS20{F0: 1}
                                                        _ = v168
                                                    }
                                                    var v169 int32 = 516
                                                    v6 = int8(v169)
                                                    _ = v169
                                                } else {
                                                    var v170 float32 = 0.1
                                                    v147 -= 2147483647
                                                    _ = v170
                                                }
                                                var v171 uint64 = 743
                                                var v172 uint16 = 32767
                                                _ = v171
                                                _ = v172
                                            }
                                            var v173 int8 = v6
                                            var v174 uint64 = uint64(v5)
                                            if false {
                                                var v175 int = fi146
                                                fmt.Println(v147)
                                                _ = v175
                                            } else {
                                                v173 -= (v6 / ((v6) | 1))
                                                var v176 int = (fi146 % (((fi146 << ((v5) & 63))) | 1))
                                                v6 = v4
                                                _ = v176
                                            }
                                            _ = v173
                                            _ = v174
                                        }
                                    }
                                    if true {
                                        var v177 int16 = 1000
                                        _ = v177
                                    }
                                } else {
                                    var v178 float32 = (1e10 + ((-2.25 + (0.0 / 0.1)) / -2.25))
                                    _ = v178
                                }
                                v3[(v147 + (^(v147)))] = (("hello" + ("hello" + "a")) + "foo bar")
                            } else {
                                v2 = append(v2, fi146)
                                var v179 float32 = (3.14159 * (0.0 / -2.25))
                                _ = v179
                            }
                        }
                    }
                case "foo bar":
                    v147 = v147
                    v5 = (v5 * 65535)
                    switch "" {
                    case "":
                        v180 := GenS3{F0: (v1 >> ((727) & 63)), F1: 281, F2: (!true)}
                        _ = v180
                    case "\n\t":
                        var v181 bool = false
                        _ = v181
                    case "a":
                        var v182 int32 = -2147483648
                        _ = v182
                    default:
                        v183 := GenS6{F0: (v6 * v6), F1: 145, F2: v6}
                        _ = v183
                    }
                }
                _ = fi148
            }
            _ = v147
        }
        _ = fi146
    }
    if (true || true) {
        if (3.14159 >= 3.14159) {
            fmt.Println(v5, v6)
        }
        for fi184 := 0; fi184 < 13; fi184++ {
            fmt.Println(v6, fi184)
            _ = fi184
        }
        v185 := GenS10{F0: 42, F1: v5}
        _ = v185
    } else {
        var v186 int64 = (v1 - (v1 / ((v1) | 1)))
        var v187 int32 = 0
        _ = v186
        _ = v187
    }
    var v188 uint32 = 798
    v189 := GenS4{F0: 2, F1: 128, F2: (v5 << ((v5) & 63))}
    v4 = v4
    var v190 int8 = (v6 >> ((321) & 63))
    for fi191 := 0; fi191 < 5; fi191++ {
        var v192 uint32 = v188
        _ = fi191
        _ = v192
    }
    if false {
        var v193 uint32 = v188
        if true {
            var v194 int8 = v4
            for fi195 := 0; fi195 < 15; fi195++ {
                v5 -= v5
                v194 -= (^(v190))
                _ = fi195
            }
            v196 := GenS16{F0: 255}
            _ = v194
            _ = v196
        }
        v193 += (v193 & v188)
        _ = v193
    } else {
        var v197 string = "foo bar"
        if (1 >= 0) {
            for fi198 := 0; fi198 < 8; fi198++ {
                switch fi198 {
                case 598:
                    for fi199 := 0; fi199 < 19; fi199++ {
                        v197 = (v197 + v197)
                        v3[v188] = "foo bar"
                        _ = fi199
                    }
                case 127:
                    v200 := GenS22{F0: v197, F1: fi198, F2: -1}
                    _ = v200
                case -245:
                    var v201 float32 = ((1.5 - 0.1) / -2.25)
                    _ = v201
                default:
                    var v202 uint32 = (v188 >> (((v5 * (^(v5)))) & 63))
                    _ = v202
                }
                switch "\n\t" {
                case "a":
                    v4 -= (^(v6))
                default:
                    var v203 float64 = 0.1
                    _ = v203
                }
                _ = fi198
            }
            switch 127 {
            case -2147483648:
                _ = len(v3)
            }
        } else {
            v1 = v1
            v204 := GenS1{F0: 65535, F1: v1}
            _ = v204
        }
        _ = v197
    }
    var v205 int32 = 849
    fmt.Println("v1:", v1)
    fmt.Println("v188:", v188)
    fmt.Println("v190:", v190)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v205:", v205)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v188
    _ = v189
    _ = v190
    _ = v2
    _ = v205
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
