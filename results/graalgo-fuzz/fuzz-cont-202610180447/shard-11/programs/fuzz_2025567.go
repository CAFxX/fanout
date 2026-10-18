package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16; F1 int16 }
type GenS9 struct { F0 int32; F1 int64; F2 int }
type GenS10 struct { F0 uint; F1 string }
type GenS11 struct { F0 int64; F1 uint8 }
type GenS12 struct { F0 int16; F1 int64; F2 bool }
type GenS13 struct { F0 uint8; F1 int64 }
type GenS14 struct { F0 string; F1 uint; F2 int }
type GenS15 struct { F0 uint; F1 int32 }
type GenS16 struct { F0 uint64; F1 int32 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 uint32; F1 int32 }
type GenS20 struct { F0 uint32; F1 uint; F2 int64 }
type GenS21 struct { F0 bool; F1 int }
type GenS22 struct { F0 bool; F1 int16 }
type GenS23 struct { F0 uint16; F1 uint8; F2 int }
type GenS24 struct { F0 string }
type GenS25 struct { F0 uint }
type GenS26 struct { F0 int32; F1 uint }

func main() {
    var v1 string = "\n\t"
    v2 := []string{v1, "a", v1}
    v3 := []int{1, 57, -2147483648, -9223372036854775808}
    var v4 int16 = -32768
    v5 := []int{-2147483648, 4294967295}
    v1 = (v1 + "a")
    if (!false) {
        v1 = v1
        fmt.Println(v1, v4)
        if false {
            fmt.Println(v4, v4)
            var v6 float32 = (0.1 - ((float32(1.5) + 1e10) * ((((0.0 * 3.14159) - (1.5 / ((float32(1.5) + (((((1e10 - (((1.5 * (3.14159 - -2.25)) * (1.5 / 1e10)) + 0.1)) / 1.0) * 3.14159) - 0.0) + (-2.25 - (1.5 + 1e10)))) * (((0.1 + -2.25) + 1.5) - -2.25)))) + 0.0) / 1.5)))
            _ = v6
        } else {
            fmt.Println(v1, v4)
        }
    } else {
        if true {
            if (float32(-2.25) > -2.25) {
                var v7 int64 = 1
                _ = v7
            }
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := GenS1{F0: -83}
                fmt.Println(fi8, v4)
                for fi10 := 0; fi10 < 6; fi10++ {
                    if (!false) {
                        var v11 float32 = (float32(0.1) - float32(1.5))
                        _ = v11
                    }
                    v12 := GenS1{F0: 127}
                    _ = fi10
                    _ = v12
                }
                _ = fi8
                _ = v9
            }
        } else {
            if ((float32(0.0) * (0.0 - float32(1e10))) < ((0.1 + -2.25) * 0.1)) {
                switch -1 {
                case 42:
                    var v13 int8 = 100
                    _ = v13
                case -809:
                    var v14 uint8 = 0
                    _ = v14
                case 127:
                    var v15 uint8 = 128
                    _ = v15
                }
                if (42 == 944) {
                    var v16 string = "foo bar"
                    if len(v2) > 0 {
                        v2[0] = v1
                    }
                    for fi17 := 0; fi17 < 19; fi17++ {
                        v16 = (("a" + "\n\t") + v16)
                        _ = fi17
                    }
                    _ = v16
                }
                var v18 string = "foo bar"
                _ = v18
            } else {
                var v19 int = -2147483648
                for fi20 := 0; fi20 < 5; fi20++ {
                    if false {
                        for fi21 := 0; fi21 < 1; fi21++ {
                            v4 = v4
                            _ = fi21
                        }
                    }
                    var v22 int = -312
                    _ = fi20
                    _ = v22
                }
                for fi23 := 0; fi23 < 13; fi23++ {
                    for fi24 := 0; fi24 < 16; fi24++ {
                        v25 := GenS2{F0: 1, F1: 65535}
                        fmt.Println(v1, fi23, fi24)
                        _ = fi24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 20; fi26++ {
                        v27 := GenS2{F0: 42, F1: 255}
                        v1 = ("hello" + v1)
                        _ = fi26
                        _ = v27
                    }
                    _ = fi23
                }
                _ = v19
            }
            switch "\n\t" {
            case "foo bar":
                for fi28 := 0; fi28 < 14; fi28++ {
                    if (4294967295 < 152) {
                        for fi29 := 0; fi29 < 4; fi29++ {
                            if (!true) {
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = 128
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = (((("" + (v1 + v1)) + "\n\t") + ("foo bar" + ((v1 + (("" + (v1 + (("" + ((v1 + (((((((v1 + "a") + "foo bar") + "a") + "") + "a") + "\n\t") + "\n\t")) + (v1 + v1))) + "hello"))) + (v1 + v1))) + (v1 + (v1 + v1))))) + "a")
                                    }
                                } else {
                                    v30 := GenS2{F0: 851, F1: 18446744073709551615}
                                    if true {
                                        switch v1 {
                                        case "\n\t":
                                            v4 = (-(v4))
                                        }
                                    }
                                    fmt.Println(fi28, v4)
                                    _ = v30
                                }
                                v1 = "\n\t"
                            }
                            var v31 bool = false
                            _ = fi29
                            _ = v31
                        }
                        for fi32 := 0; fi32 < 1; fi32++ {
                            v5 = append(v5, fi28)
                            var v33 uint16 = 255
                            _ = fi32
                            _ = v33
                        }
                    } else {
                        v34 := GenS3{F0: 83}
                        _ = v34
                    }
                    switch "hello" {
                    case "foo bar":
                        if false {
                            v1 = ((v1 + v1) + ((v1 + (v1 + ((v1 + v1) + (v1 + v1)))) + ("" + v1)))
                            var v35 int64 = 9223372036854775807
                            fmt.Println(fi28, v1)
                            _ = v35
                        }
                    case "a":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    if true {
                        fmt.Println(fi28, v4)
                    }
                    _ = fi28
                }
            case "hello":
                v1 = v1
            }
        }
        if (6 >= 16) {
            v4 = v4
            if true {
                fmt.Println(v1, v1)
                for fi36 := 0; fi36 < 12; fi36++ {
                    fmt.Println(fi36)
                    v37 := GenS2{F0: 2, F1: 42}
                    _ = fi36
                    _ = v37
                }
            }
        }
        var v38 uint8 = 139
        _ = v38
    }
    if false {
        var v39 float64 = -2.25
        var v40 uint16 = 715
        _ = v39
        _ = v40
    }
    var v41 int32 = 2147483647
    for fi42 := 0; fi42 < 4; fi42++ {
        for fi43 := 0; fi43 < 5; fi43++ {
            switch (v1 + v1) {
            case "\n\t":
                v1 = (((v1 + "hello") + (v1 + (((v1 + (v1 + v1)) + v1) + ""))) + v1)
            case "hello":
                v44 := GenS4{F0: 255, F1: ("a" + v1)}
                _ = v44
            case "a":
                v45 := GenS5{F0: (("hello" + (v1 + (v1 + v1))) + v1)}
                _ = v45
            default:
                switch ("\n\t" + "") {
                case "hello":
                    fmt.Println(v1, v1)
                default:
                    fmt.Println(fi42, v41, v4)
                }
            }
            fmt.Println(fi43, fi42)
            v3 = append(v3, fi42)
            var v46 int64 = 1
            switch (("a" + (v1 + ((v1 + "foo bar") + v1))) + "") {
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "foo bar":
                v3 = append(v3, fi42)
            case "\n\t":
                v4 = 427
            }
            v46 = 9223372036854775807
            var v47 uint64 = 0
            v46 += (v46 & v46)
            v46 -= v46
            _ = fi43
            _ = v46
            _ = v47
        }
        _ = fi42
    }
    var v48 uint = 65535
    fmt.Println(v48, v48)
    if false {
        switch 42 {
        case 1:
            fmt.Println(v4)
        case -888:
            switch 42 {
            case 127:
                var v49 int = -2147483648
                _ = v49
            case -394:
                var v50 float32 = ((-2.25 - 0.1) * 0.0)
                _ = v50
            case 32767:
                v48 += (v48 >> ((42) & 63))
            }
        }
    }
    switch 32767 {
    case 13:
        fmt.Println(v48)
    case 758:
        v1 = (v1 + ((v1 + (v1 + ("foo bar" + "\n\t"))) + v1))
    case 32767:
        var v51 uint16 = 255
        _ = v51
    }
    if (true && (true && true)) {
        fmt.Println(v41, v41)
        if true {
            v52 := GenS3{F0: 0}
            v53 := GenS2{F0: 255, F1: 38}
            v5 = append(v5, int(v4))
            _ = v52
            _ = v53
        }
        v1 = "hello"
    } else {
        v54 := GenS6{F0: 0}
        var v55 int16 = (v4 | -629)
        _ = v54
        _ = v55
    }
    switch (v1 + v1) {
    case "hello":
        v48 += v48
    case "":
        var v56 float32 = 1e10
        _ = v56
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = v1
        }
    default:
        v57 := GenS7{F0: (v41 & v41)}
        _ = v57
    }
    if false {
        switch v1 {
        case "\n\t":
            fmt.Println(v4, v4, v48)
        case "foo bar":
            for fi58 := 0; fi58 < 16; fi58++ {
                if (3.14159 == 0.1) {
                    v59 := GenS8{F0: 892, F1: v4}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch fi58 {
                    case 42:
                        var v60 bool = (-318 >= -1)
                        _ = v60
                    case -1:
                        var v61 int = fi58
                        _ = v61
                    }
                    _ = v59
                } else {
                    var v62 uint32 = 65535
                    switch v1 {
                    case "a":
                        v62 = v62
                    default:
                        if len(v2) > 0 {
                            v2[0] = ((v1 + ("hello" + v1)) + "foo bar")
                        }
                    }
                    var v63 bool = true
                    _ = v62
                    _ = v63
                }
                v4 = (v4 << ((v48) & 63))
                fmt.Println(v41, v1, v48)
                _ = fi58
            }
        case "a":
            if (0 != 50) {
                for fi64 := 0; fi64 < 7; fi64++ {
                    var v65 uint16 = 32767
                    _ = fi64
                    _ = v65
                }
                var v66 int16 = v4
                for fi67 := 0; fi67 < 12; fi67++ {
                    var v68 int8 = -1
                    _ = fi67
                    _ = v68
                }
                _ = v66
            }
        default:
            v4 += v4
        }
    }
    var v69 int8 = -1
    var v70 float64 = -2.25
    fmt.Println(v4, v48, v4)
    for fi71 := 0; fi71 < 1; fi71++ {
        if (!(!(false || (false || (2 != 42))))) {
            switch (fi71 >> ((200) & 63)) {
            case 4294967295:
                if (true || false) {
                    for fi72 := 0; fi72 < 9; fi72++ {
                        v70 += v70
                        _ = fi72
                    }
                } else {
                    if ((0 > 228) && (!true)) {
                        if (3.14159 < ((3.14159 - v70) + 0.0)) {
                            switch "\n\t" {
                            case "a":
                                switch v1 {
                                case "a":
                                    v73 := GenS6{F0: 2147483647}
                                    _ = v73
                                case "\n\t":
                                    var v74 string = (v1 + "\n\t")
                                    _ = v74
                                case "":
                                    v75 := GenS9{F0: (v41 - v41), F1: 9223372036854775807, F2: fi71}
                                    _ = v75
                                }
                            }
                        } else {
                            v76 := GenS10{F0: v48, F1: ((((v1 + ((v1 + (v1 + ((v1 + v1) + v1))) + v1)) + v1) + (v1 + (("a" + (v1 + "hello")) + v1))) + "foo bar")}
                            if true {
                                v4 = (-(v4))
                                var v77 string = ("" + v1)
                                v70 -= ((float64(float32(1e10)) / float64(1.5)) * float64(((-2.25 * 0.1) / (3.14159 - 0.0))))
                                _ = v77
                            }
                            v41 = (v41 ^ v41)
                            _ = v76
                        }
                        v78 := GenS11{F0: 1, F1: 42}
                        if (-893 < 42) {
                            fmt.Println(v1, v69)
                            switch v1 {
                            case "\n\t":
                                switch fi71 {
                                case 0:
                                    for fi79 := 0; fi79 < 9; fi79++ {
                                        v41 = 1
                                        var v80 uint64 = 152
                                        var v81 string = v1
                                        _ = fi79
                                        _ = v80
                                        _ = v81
                                    }
                                case 557:
                                    if false {
                                        if ((true && true) || true) {
                                            fmt.Println(v70, v70, v69)
                                        }
                                        var v82 uint64 = 2
                                        fmt.Println(fi71, v69, v48)
                                        _ = v82
                                    }
                                default:
                                    if (1e10 != ((-2.25 + -2.25) + 3.14159)) {
                                        v48 -= 0
                                    } else {
                                        for fi83 := 0; fi83 < 12; fi83++ {
                                            var v84 string = v1
                                            _ = fi83
                                            _ = v84
                                        }
                                    }
                                }
                            case "a":
                                var v85 float64 = ((3.14159 + 1e10) / float64(float32(0.1)))
                                _ = v85
                            default:
                                switch ((v1 + (v1 + "foo bar")) + (("hello" + v1) + "foo bar")) {
                                case "":
                                    var v86 int64 = 127
                                    _ = v86
                                case "a":
                                    v4 += -859
                                case "foo bar":
                                    switch v1 {
                                    case "a":
                                        for fi87 := 0; fi87 < 2; fi87++ {
                                            var v88 uint8 = 128
                                            _ = fi87
                                            _ = v88
                                        }
                                    case "\n\t":
                                        for fi89 := 0; fi89 < 10; fi89++ {
                                            if true {
                                                if len(v3) > 0 {
                                                    v3[0] = fi71
                                                }
                                                if true {
                                                    fmt.Println(v1, v4)
                                                    v4 += (^(v4))
                                                    if (((v70 != 0.1) && (v48 >= (v48 & v48))) || ((!(-692 <= 0)) || (!true))) {
                                                        if true {
                                                            switch (v1 + "foo bar") {
                                                            case "foo bar":
                                                                var v90 uint32 = 481
                                                                _ = v90
                                                            case "\n\t":
                                                                v69 -= v69
                                                            case "":
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            default:
                                                                fmt.Println(fi71)
                                                            }
                                                            switch 196 {
                                                            case 9223372036854775807:
                                                                if (!false) {
                                                                    v69 += (v69 | (v69 & (v69 - v69)))
                                                                    v4 = (v4 << (((v48 << ((v48) & 63))) & 63))
                                                                } else {
                                                                    var v91 int32 = v41
                                                                    _ = v91
                                                                }
                                                            case -2147483648:
                                                                var v92 int32 = (v41 | (v41 * 42))
                                                                _ = v92
                                                            case 920:
                                                                if true {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                            }
                                                            v93 := GenS12{F0: v4, F1: -1, F2: (!true)}
                                                            _ = v93
                                                        }
                                                        v94 := GenS13{F0: 42, F1: -401}
                                                        _ = v94
                                                    } else {
                                                        v95 := GenS11{F0: 0, F1: 128}
                                                        var v96 string = (("\n\t" + "") + v1)
                                                        if false {
                                                            if false {
                                                                var v97 float64 = v70
                                                                _ = v97
                                                            }
                                                        } else {
                                                            v48 += (v48 & 683)
                                                            switch fi89 {
                                                            case 32767:
                                                                if true {
                                                                    v2 = append(v2, v1)
                                                                    var v98 uint8 = 2
                                                                    _ = v98
                                                                } else {
                                                                    switch v96 {
                                                                    case "\n\t":
                                                                        fmt.Println(v96)
                                                                    }
                                                                    v41 -= (v41 * v41)
                                                                    switch (v1 + (v96 + "foo bar")) {
                                                                    case "hello":
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    case "\n\t":
                                                                        v4 += (v4 << ((65535) & 63))
                                                                    case "":
                                                                        v99 := GenS7{F0: (v41 + v41)}
                                                                        _ = v99
                                                                    }
                                                                }
                                                            case -2147483648:
                                                                v100 := GenS5{F0: v1}
                                                                _ = v100
                                                            default:
                                                                if len(v3) > 0 {
                                                                    v3[0] = 753
                                                                }
                                                            }
                                                        }
                                                        v70 = (1e10 - float64(0.0))
                                                        v69 = 81
                                                        fmt.Println(v48)
                                                        if true {
                                                            v69 += v69
                                                            var v101 uint = (v48 << (((v48 << ((v48) & 63))) & 63))
                                                            var v102 uint8 = 127
                                                            _ = v101
                                                            _ = v102
                                                        }
                                                        var v103 int32 = v41
                                                        if len(v5) > 0 {
                                                            v5[0] = fi89
                                                        }
                                                        _ = v103
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                    switch (fi71 % ((fi71) | 1)) {
                                                    case 32768:
                                                        v104 := GenS14{F0: "\n\t", F1: (v48 ^ v48), F2: fi71}
                                                        _ = v104
                                                    default:
                                                        if false {
                                                            var v105 float32 = (0.1 * 1e10)
                                                            _ = v105
                                                        } else {
                                                            var v106 string = v1
                                                            v41 += v41
                                                            _ = v106
                                                        }
                                                    }
                                                    if (!((1 >= 2) || true)) {
                                                        var v107 uint16 = 0
                                                        switch fi71 {
                                                        case 2147483647:
                                                            v4 = -1
                                                        case 9223372036854775807:
                                                            if (11 == 42) {
                                                                v108 := GenS15{F0: (v48 / ((v48) | 1)), F1: -1}
                                                                _ = v108
                                                            } else {
                                                                v107 = v107
                                                                v109 := GenS2{F0: v107, F1: 42}
                                                                _ = v109
                                                            }
                                                        case 127:
                                                            v110 := GenS16{F0: 654, F1: v41}
                                                            _ = v110
                                                        }
                                                        switch fi71 {
                                                        case 558:
                                                            var v111 int32 = v41
                                                            _ = v111
                                                        case 42:
                                                            if true {
                                                                v112 := GenS17{F0: 953}
                                                                v2 = append(v2, ((v1 + (("hello" + "hello") + (v1 + ("a" + ((v1 + ("foo bar" + v1)) + ("" + (v1 + v1))))))) + ((((v1 + v1) + (v1 + v1)) + "\n\t") + ((v1 + v1) + ""))))
                                                                _ = v112
                                                            }
                                                        case -9223372036854775808:
                                                            v113 := GenS6{F0: 2147483648}
                                                            _ = v113
                                                        default:
                                                            v114 := GenS18{F0: (v4 << (((v48 << (((v48 % (((v48 / ((uint(v107)) | 1))) | 1))) & 63))) & 63))}
                                                            _ = v114
                                                        }
                                                        var v115 int = (fi89 + fi89)
                                                        if (!false) {
                                                            fmt.Println(v107, v69)
                                                            v116 := GenS4{F0: 1, F1: v1}
                                                            v69 = (v69 + (v69 >> (((v48 / (((v48 / ((v48) | 1))) | 1))) & 63)))
                                                            v48 += 472
                                                            _ = v116
                                                        }
                                                        fmt.Println(fi89, v69)
                                                        var v117 uint64 = 863
                                                        _ = v107
                                                        _ = v115
                                                        _ = v117
                                                    } else {
                                                        if (18446744073709551615 != 484) {
                                                            v118 := GenS19{F0: 612, F1: 1}
                                                            _ = v118
                                                        } else {
                                                            if false {
                                                                v70 = (v70 + (v70 * (v70 + (((1e10 - v70) - v70) - ((v70 * v70) + 0.1)))))
                                                                if len(v3) > 0 {
                                                                    v3[0] = -811
                                                                }
                                                                if (901 < 255) {
                                                                    v119 := GenS4{F0: 237, F1: ""}
                                                                    v2 = append(v2, (v1 + v1))
                                                                    _ = v119
                                                                } else {
                                                                    if false {
                                                                        if ((false && (false && (true && ((0.1 * 1.5) >= 3.14159)))) || true) {
                                                                            if false {
                                                                                if (v48 > v48) {
                                                                                    var v120 int = int(v41)
                                                                                    var v121 int16 = v4
                                                                                    fmt.Println(v48)
                                                                                    _ = v120
                                                                                    _ = v121
                                                                                }
                                                                                v122 := GenS14{F0: v1, F1: (v48 - v48), F2: 192}
                                                                                v3 = append(v3, (fi71 / ((9223372036854775807) | 1)))
                                                                                _ = v122
                                                                            } else {
                                                                                var v123 float64 = (v70 * v70)
                                                                                _ = v123
                                                                            }
                                                                            v124 := GenS14{F0: ("" + "\n\t"), F1: (v48 | v48), F2: fi71}
                                                                            _ = v124
                                                                        } else {
                                                                            v5 = append(v5, fi89)
                                                                            var v125 int32 = (v41 >> ((v48) & 63))
                                                                            var v126 uint32 = 2147483647
                                                                            var v127 int = -839
                                                                            v128 := GenS20{F0: (v126 / ((v126) | 1)), F1: 42, F2: 42}
                                                                            if (!true) {
                                                                                switch v1 {
                                                                                case "":
                                                                                    if false {
                                                                                        var v129 string = v1
                                                                                        switch v1 {
                                                                                        case "\n\t":
                                                                                            _ = len(v2)
                                                                                        case "hello":
                                                                                            switch "a" {
                                                                                            case "a":
                                                                                                v130 := GenS16{F0: 0, F1: v125}
                                                                                                var v131 uint8 = 42
                                                                                                fmt.Println(v4)
                                                                                                _ = v130
                                                                                                _ = v131
                                                                                            case "foo bar":
                                                                                                var v132 float32 = -2.25
                                                                                                _ = v132
                                                                                            default:
                                                                                                v133 := GenS7{F0: v125}
                                                                                                _ = v133
                                                                                            }
                                                                                        case "foo bar":
                                                                                            var v134 string = "\n\t"
                                                                                            _ = v134
                                                                                        }
                                                                                        _ = v129
                                                                                    }
                                                                                case "a":
                                                                                    var v135 int64 = 42
                                                                                    _ = v135
                                                                                case "\n\t":
                                                                                    switch fi89 {
                                                                                    case -583:
                                                                                        v136 := GenS21{F0: (true && true), F1: (fi89 * fi89)}
                                                                                        _ = v136
                                                                                    case 86:
                                                                                        v125 -= v125
                                                                                    }
                                                                                }
                                                                                var v137 uint64 = 42
                                                                                var v138 bool = true
                                                                                _ = v137
                                                                                _ = v138
                                                                            }
                                                                            if (!false) {
                                                                                switch (v1 + "foo bar") {
                                                                                case "a":
                                                                                    if false {
                                                                                        fmt.Println(v127, v70)
                                                                                        switch (v1 + v1) {
                                                                                        case "foo bar":
                                                                                            if (!false) {
                                                                                                fmt.Println(v126, v127)
                                                                                                v139 := GenS3{F0: 0}
                                                                                                v140 := GenS22{F0: false, F1: v4}
                                                                                                _ = v139
                                                                                                _ = v140
                                                                                            } else {
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                v1 = v1
                                                                                            }
                                                                                        case "a":
                                                                                            v41 += 42
                                                                                        default:
                                                                                            v126 = 2
                                                                                        }
                                                                                    } else {
                                                                                        _ = len(v3)
                                                                                        var v141 uint16 = 2
                                                                                        v126 -= 65535
                                                                                        if true {
                                                                                            v142 := GenS13{F0: 1, F1: 42}
                                                                                            switch v1 {
                                                                                            case "a":
                                                                                                var v143 uint16 = 1
                                                                                                _ = v143
                                                                                            default:
                                                                                                var v144 uint32 = uint32(v141)
                                                                                                _ = v144
                                                                                            }
                                                                                            _ = v142
                                                                                        }
                                                                                        var v145 int8 = int8(v41)
                                                                                        _ = v141
                                                                                        _ = v145
                                                                                    }
                                                                                case "foo bar":
                                                                                    var v146 uint = v48
                                                                                    _ = v146
                                                                                case "\n\t":
                                                                                    v1 = (v1 + v1)
                                                                                }
                                                                            }
                                                                            _ = v125
                                                                            _ = v126
                                                                            _ = v127
                                                                            _ = v128
                                                                        }
                                                                    }
                                                                    _ = len(v5)
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    v48 = (v48 - v48)
                                                                }
                                                            }
                                                            var v147 uint64 = 2
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v147 += (v147 & (v147 & (v147 | (v147 / ((2) | 1)))))
                                                            case "foo bar":
                                                                fmt.Println(v147, fi71, v48)
                                                            default:
                                                                if false {
                                                                    v69 = 100
                                                                }
                                                            }
                                                            _ = v147
                                                        }
                                                        var v148 int = fi71
                                                        var v149 uint16 = 65535
                                                        _ = v148
                                                        _ = v149
                                                    }
                                                }
                                            } else {
                                                switch fi89 {
                                                case 4294967295:
                                                    fmt.Println(v48, fi71)
                                                }
                                                var v150 bool = (true || (false && false))
                                                var v151 float64 = v70
                                                _ = v150
                                                _ = v151
                                            }
                                            _ = fi89
                                        }
                                    case "hello":
                                        var v152 uint16 = 1
                                        _ = v152
                                    default:
                                        for fi153 := 0; fi153 < 17; fi153++ {
                                            switch (fi153 ^ (fi153 ^ fi71)) {
                                            case 128:
                                                switch "hello" {
                                                case "\n\t":
                                                    var v154 int8 = v69
                                                    _ = v154
                                                case "":
                                                    if len(v5) > 0 {
                                                        v5[0] = -9223372036854775808
                                                    }
                                                default:
                                                    switch "hello" {
                                                    case "a":
                                                        if false {
                                                            switch v1 {
                                                            case "hello":
                                                                v5 = append(v5, int(v4))
                                                            default:
                                                                _ = len(v5)
                                                            }
                                                            v155 := GenS11{F0: 42, F1: 1}
                                                            var v156 float32 = 0.1
                                                            _ = v155
                                                            _ = v156
                                                        } else {
                                                            fmt.Println(v48)
                                                            if len(v3) > 0 {
                                                                v3[0] = (fi153 >> (((v48 | (v48 ^ v48))) & 63))
                                                            }
                                                        }
                                                    case "\n\t":
                                                        if true {
                                                            v41 += (v41 | v41)
                                                            switch v1 {
                                                            case "\n\t":
                                                                v157 := GenS15{F0: 330, F1: (v41 & 0)}
                                                                _ = v157
                                                            }
                                                        } else {
                                                            fmt.Println(v48, v48)
                                                        }
                                                    case "foo bar":
                                                        if (2147483647 == 116) {
                                                            fmt.Println(v41)
                                                            v48 = 0
                                                            if (2 < 18446744073709551615) {
                                                                var v158 uint = 4294967295
                                                                v4 = 1000
                                                                var v159 uint64 = 18446744073709551615
                                                                _ = v158
                                                                _ = v159
                                                            }
                                                        } else {
                                                            fmt.Println(fi71, fi153, v70)
                                                            fmt.Println(v1, v48)
                                                        }
                                                    }
                                                }
                                            default:
                                                var v160 float64 = float64(float32(-2.25))
                                                _ = v160
                                            }
                                            if (2 <= 2) {
                                                var v161 uint8 = 204
                                                var v162 uint8 = (v161 / (((v161 % ((v161) | 1))) | 1))
                                                _ = v161
                                                _ = v162
                                            } else {
                                                v2 = append(v2, v1)
                                                v41 = 2147483647
                                            }
                                            v163 := GenS23{F0: 189, F1: 128, F2: int(v48)}
                                            _ = fi153
                                            _ = v163
                                        }
                                    }
                                }
                            }
                            switch 0 {
                            case 32768:
                                if (v41 > 1) {
                                    v48 = v48
                                }
                            case 32767:
                                if len(v3) > 0 {
                                    v3[0] = fi71
                                }
                            default:
                                for fi164 := 0; fi164 < 9; fi164++ {
                                    v165 := GenS11{F0: 962, F1: 128}
                                    _ = fi164
                                    _ = v165
                                }
                            }
                        } else {
                            if false {
                                v41 -= (v41 % ((v41) | 1))
                            } else {
                                for fi166 := 0; fi166 < 15; fi166++ {
                                    fmt.Println(v1)
                                    if (-2.25 > ((1e10 / (float64(1e10) - float64(3.14159))) + float64((float32(1.5) * 3.14159)))) {
                                        fmt.Println(fi71, fi166, v4)
                                        if false {
                                            var v167 int16 = 1000
                                            _ = v167
                                        }
                                    }
                                    switch "hello" {
                                    case "hello":
                                        switch fi166 {
                                        case -2147483648:
                                            fmt.Println(v69)
                                        case 0:
                                            if true {
                                                var v168 int64 = -824
                                                _ = v168
                                            } else {
                                                var v169 int64 = -1
                                                var v170 uint16 = 65535
                                                _ = v169
                                                _ = v170
                                            }
                                        case -436:
                                            var v171 uint32 = 4294967295
                                            var v172 uint32 = (^(v171))
                                            _ = v171
                                            _ = v172
                                        default:
                                            v173 := GenS8{F0: 0, F1: (v4 + -15)}
                                            _ = v173
                                        }
                                        switch 4294967295 {
                                        case 2147483647:
                                            if false {
                                                fmt.Println(fi71, v69, v4)
                                                fmt.Println(fi166, v4, v4)
                                            }
                                        case -822:
                                            var v174 int64 = int64(v4)
                                            _ = v174
                                        default:
                                            if ((("foo bar" + "a") < "a") && false) {
                                                var v175 string = v1
                                                _ = v175
                                            } else {
                                                var v176 float64 = v70
                                                fmt.Println(v176, v4)
                                                v2 = append(v2, v1)
                                                var v177 uint = (^(v48))
                                                var v178 uint64 = 18446744073709551615
                                                v179 := GenS15{F0: 18446744073709551615, F1: -594}
                                                _ = v176
                                                _ = v177
                                                _ = v178
                                                _ = v179
                                            }
                                        }
                                        v41 -= (v41 ^ (v41 & -2147483648))
                                    case "\n\t":
                                        v180 := GenS2{F0: 65535, F1: (v48 ^ v48)}
                                        _ = v180
                                    case "a":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        var v181 int32 = v41
                                        _ = v181
                                    }
                                    _ = fi166
                                }
                                _ = len(v5)
                            }
                            if (false && false) {
                                switch ((v1 + v1) + "") {
                                case "foo bar":
                                    switch fi71 {
                                    case 9223372036854775807:
                                        switch "a" {
                                        case "a":
                                            for fi182 := 0; fi182 < 13; fi182++ {
                                                v183 := GenS24{F0: v1}
                                                v2 = append(v2, "")
                                                _ = fi182
                                                _ = v183
                                            }
                                        default:
                                            var v184 uint64 = 0
                                            _ = v184
                                        }
                                    case -2147483648:
                                        v1 = (v1 + "a")
                                    default:
                                        for fi185 := 0; fi185 < 1; fi185++ {
                                            v4 = (v4 >> (((^(v48))) & 63))
                                            _ = fi185
                                        }
                                    }
                                case "\n\t":
                                    v41 -= v41
                                default:
                                    var v186 int8 = 100
                                    _ = v186
                                }
                                var v187 uint32 = 65535
                                var v188 bool = ((4294967295 >= 2) && false)
                                _ = v187
                                _ = v188
                            } else {
                                for fi189 := 0; fi189 < 3; fi189++ {
                                    var v190 int64 = -350
                                    _ = fi189
                                    _ = v190
                                }
                                if false {
                                    v4 -= 227
                                    v41 = 42
                                }
                                v70 += 3.14159
                            }
                            switch v1 {
                            case "a":
                                v191 := GenS17{F0: v48}
                                _ = v191
                            case "\n\t":
                                var v192 float32 = (((-2.25 / 1.5) * 0.1) - (0.1 * ((1e10 / 1e10) / (1e10 / (((3.14159 / -2.25) * ((((1.5 - (float32((v70 + (((v70 * ((0.0 - (((float64((3.14159 + 0.0)) - float64(-2.25)) / v70) + 0.1)) / v70)) * v70) / v70))) / 1.0)) - 1e10) * ((1e10 * ((((1e10 * (0.0 + -2.25)) * 0.1) * float32(v70)) - 3.14159)) * (0.1 * (float32((float64(1e10) + ((v70 + v70) * v70))) * (0.1 * 3.14159))))) * (-2.25 / 1.5))) * float32((v70 / 1.5)))))))
                                _ = v192
                            }
                        }
                        _ = v78
                    }
                    var v193 uint = 799
                    if true {
                        v194 := GenS11{F0: -9223372036854775808, F1: 128}
                        v195 := GenS20{F0: 212, F1: (^(v48)), F2: 1}
                        _ = v194
                        _ = v195
                    } else {
                        fmt.Println(v1, v41, v4)
                        var v196 bool = (((-(v69)) <= (v69 - v69)) || false)
                        var v197 uint32 = 1
                        _ = v196
                        _ = v197
                    }
                    _ = v193
                }
            default:
                switch v1 {
                case "\n\t":
                    _ = len(v3)
                default:
                    if (v4 != (-(v4))) {
                        var v198 float64 = ((((v70 / v70) * ((1e10 / (v70 / (v70 - v70))) * float64(1e10))) / (1.5 - 3.14159)) * (0.0 - (v70 / 1.0)))
                        _ = v198
                    }
                }
            }
            if ((-2147483648 >= -2147483648) && ((v69 % (((v69 % ((42) | 1))) | 1)) != (v69 + (^(v69))))) {
                v48 += 287
                fmt.Println(v70, fi71)
                for fi199 := 0; fi199 < 18; fi199++ {
                    var v200 float64 = v70
                    if (false || true) {
                        switch 434 {
                        case -222:
                            fmt.Println(v69, v48)
                        case 127:
                            var v201 int32 = -1
                            _ = v201
                        case 42:
                            _ = len(v5)
                        default:
                            if false {
                                switch v1 {
                                case "a":
                                    var v202 float64 = 0.1
                                    _ = v202
                                case "\n\t":
                                    var v203 uint64 = 698
                                    _ = v203
                                }
                                var v204 string = ""
                                if true {
                                    var v205 float64 = v70
                                    if true {
                                        var v206 uint = 2
                                        var v207 uint = (v48 * (^(v206)))
                                        var v208 uint64 = 42
                                        _ = v206
                                        _ = v207
                                        _ = v208
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        var v209 uint8 = 255
                                        _ = v209
                                    case "foo bar":
                                        var v210 int8 = v69
                                        _ = v210
                                    case "":
                                        fmt.Println(v4, fi199, fi199)
                                    default:
                                        v4 += 0
                                        v2 = append(v2, v1)
                                        if ((true || (272 > 42)) || (((!((v205 > (0.0 + ((((float64(3.14159) / ((v70 + v205) + 0.0)) * (v200 + -2.25)) - 0.0) / v205))) && true)) && true) || true)) {
                                            switch (fi71 & 42) {
                                            case 856:
                                                v200 += (v200 - 0.1)
                                            default:
                                                fmt.Println(v41)
                                            }
                                            if (42 == 2147483647) {
                                                v205 = (v205 / v205)
                                                if len(v5) > 0 {
                                                    v5[0] = (fi71 | 2147483647)
                                                }
                                                switch fi71 {
                                                case 42:
                                                    fmt.Println(v4, v205, v48)
                                                case -2147483648:
                                                    v41 = -554
                                                case 128:
                                                    switch v1 {
                                                    case "":
                                                        fmt.Println(fi71, v200, v200)
                                                    case "hello":
                                                        if (18446744073709551615 != 2) {
                                                            v200 = v70
                                                        }
                                                    default:
                                                        fmt.Println(v41, v48)
                                                    }
                                                default:
                                                    var v211 int64 = 989
                                                    _ = v211
                                                }
                                            } else {
                                                v41 = 0
                                            }
                                            switch v1 {
                                            case "\n\t":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            v212 := GenS21{F0: ((true || false) && (false && ((v204 == v204) || (!false)))), F1: (fi199 / ((fi199) | 1))}
                                            var v213 uint64 = 2
                                            _ = v212
                                            _ = v213
                                        } else {
                                            var v214 uint16 = 628
                                            _ = v214
                                        }
                                    }
                                    _ = v205
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v215 := GenS25{F0: 18446744073709551615}
                                _ = v204
                                _ = v215
                            } else {
                                switch (fi71 << (((^(v48))) & 63)) {
                                case 42:
                                    var v216 float32 = -2.25
                                    _ = v216
                                case -123:
                                    v200 += v200
                                }
                                switch v1 {
                                case "":
                                    switch (v1 + (v1 + v1)) {
                                    case "foo bar":
                                        var v217 uint16 = 65535
                                        _ = v217
                                    }
                                case "\n\t":
                                    var v218 uint32 = 243
                                    _ = v218
                                case "hello":
                                    switch fi199 {
                                    case 673:
                                        v2 = append(v2, v1)
                                        var v219 uint64 = 42
                                        _ = v219
                                    }
                                default:
                                    v220 := GenS19{F0: 65535, F1: v41}
                                    v200 -= 1e10
                                    _ = v220
                                }
                                if (((!true) && (!((float64(1e10) == float64((3.14159 * 1.5))) || (!(false && true))))) || (!true)) {
                                    if false {
                                        fmt.Println(v4)
                                        if (false && (((1e10 + ((((((v70 + v200) - v70) - 0.0) * 1.5) - v70) / (float64((1e10 - 0.1)) / 1.5))) != 0.0) && (0.0 != v70))) {
                                            v41 -= v41
                                        }
                                        _ = len(v5)
                                        v200 = 3.14159
                                    }
                                    fmt.Println(v4)
                                }
                                fmt.Println(v48)
                                if (v69 >= -114) {
                                    if true {
                                        fmt.Println(v48, fi199)
                                        v221 := GenS26{F0: 350, F1: (^(v48))}
                                        _ = v221
                                    } else {
                                        var v222 bool = ((fi199 / ((fi199) | 1)) >= (fi71 * 32767))
                                        _ = v222
                                    }
                                    switch v1 {
                                    case "a":
                                        v200 = (v70 / v70)
                                    case "hello":
                                        var v223 int16 = 697
                                        _ = v223
                                    default:
                                        switch v1 {
                                        case "":
                                            v48 = v48
                                            v48 += v48
                                        }
                                    }
                                } else {
                                    var v224 int16 = v4
                                    var v225 int = 4294967295
                                    if ((v41 >> ((v48) & 63)) == (v41 - -2147483648)) {
                                        var v226 uint32 = 294
                                        _ = v226
                                    }
                                    _ = v224
                                    _ = v225
                                }
                            }
                            var v227 uint8 = 1
                            _ = len(v2)
                            _ = v227
                        }
                        v69 = -1
                    }
                    _ = fi199
                    _ = v200
                }
            }
        }
        v48 -= (v48 >> ((347) & 63))
        v4 = 42
        _ = fi71
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v48
    _ = v5
    _ = v69
    _ = v70
}
