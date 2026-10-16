package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int8; F1 uint64; F2 uint }
type GenS3 struct { F0 int; F1 uint64; F2 uint8 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint32; F1 int16 }
type GenS10 struct { F0 int32; F1 int8 }
type GenS11 struct { F0 uint16; F1 string }
type GenS12 struct { F0 uint64; F1 uint16; F2 string }
type GenS13 struct { F0 int64; F1 int }
type GenS14 struct { F0 int32; F1 string; F2 bool }
type GenS15 struct { F0 string }

func main() {
    v1 := []uint32{705, 65535, 2}
    var v2 float32 = -2.25
    var v3 uint64 = 243
    var v4 int32 = -2147483648
    _ = len(v1)
    v4 = 0
    switch 127 {
    case 127:
        if (((!true) || (2147483648 >= 2)) && true) {
            for fi5 := 0; fi5 < 18; fi5++ {
                switch "foo bar" {
                case "foo bar":
                    switch (fi5 * fi5) {
                    case -519:
                        var v6 int = fi5
                        _ = v6
                    default:
                        var v7 uint8 = 42
                        _ = v7
                    }
                case "\n\t":
                    v8 := GenS1{F0: 255, F1: (312 == 0)}
                    _ = v8
                }
                _ = fi5
            }
        }
    case 2147483647:
        for fi9 := 0; fi9 < 19; fi9++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi9
        }
    case -9223372036854775808:
        switch 0 {
        case 4294967295:
            for fi10 := 0; fi10 < 11; fi10++ {
                if (0.1 > ((v2 * float32(0.0)) + (v2 * float32(1e10)))) {
                    for fi11 := 0; fi11 < 12; fi11++ {
                        if false {
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                            if ((int8(v4) != 127) || (((("foo bar" + "hello") + "\n\t") >= "a") && (255 < 128))) {
                                var v12 uint16 = 562
                                if true {
                                    v12 += v12
                                    if false {
                                        if (false && (1 > -718)) {
                                            v13 := GenS1{F0: 162, F1: true}
                                            switch ("\n\t" + ("a" + "a")) {
                                            case "a":
                                                switch "\n\t" {
                                                case "a":
                                                    if false {
                                                        v12 = 30
                                                        var v14 uint64 = (^(v3))
                                                        var v15 uint8 = 2
                                                        _ = v14
                                                        _ = v15
                                                    } else {
                                                        v12 += v12
                                                        fmt.Println(v3)
                                                        v2 -= (0.1 + (3.14159 / v2))
                                                    }
                                                    fmt.Println(v4, v4)
                                                    var v16 uint64 = (v3 - 4294967295)
                                                    _ = v16
                                                case "\n\t":
                                                    var v17 uint32 = 4294967295
                                                    _ = v17
                                                default:
                                                    if len(v1) > 0 {
                                                        v1[0] = 932
                                                    }
                                                }
                                            }
                                            if false {
                                                v3 = (v3 >> ((65535) & 63))
                                                v18 := GenS1{F0: 214, F1: ((^(fi10)) == (fi11 - -9223372036854775808))}
                                                _ = v18
                                            }
                                            _ = v13
                                        }
                                    } else {
                                        v4 += v4
                                        v19 := GenS2{F0: -9, F1: v3, F2: 18446744073709551615}
                                        _ = v19
                                    }
                                }
                                fmt.Println(v4)
                                _ = v12
                            }
                            _ = len(v1)
                        }
                        _ = fi11
                    }
                    for fi20 := 0; fi20 < 5; fi20++ {
                        switch "" {
                        case "hello":
                            v21 := GenS2{F0: -33, F1: v3, F2: 255}
                            _ = v21
                        case "\n\t":
                            v22 := GenS1{F0: 128, F1: (!false)}
                            _ = v22
                        default:
                            var v23 uint32 = 1
                            _ = v23
                        }
                        var v24 int8 = -1
                        var v25 float32 = v2
                        _ = fi20
                        _ = v24
                        _ = v25
                    }
                    v26 := GenS2{F0: 114, F1: v3, F2: 2}
                    _ = v26
                } else {
                    var v27 bool = (!true)
                    var v28 bool = (128 > 84)
                    _ = v27
                    _ = v28
                }
                v1 = append(v1, 2)
                _ = fi10
            }
        case 0:
            var v29 float32 = v2
            _ = v29
        default:
            var v30 int8 = 1
            _ = v30
        }
    default:
        _ = len(v1)
    }
    v4 += v4
    for fi31 := 0; fi31 < 11; fi31++ {
        for fi32 := 0; fi32 < 3; fi32++ {
            v33 := GenS2{F0: 29, F1: v3, F2: 255}
            if (2 > 127) {
                if false {
                    v34 := GenS1{F0: 255, F1: (fi32 != fi32)}
                    switch -954 {
                    case 1:
                        fmt.Println(v3, v4)
                    case -2147483648:
                        _ = len(v1)
                        var v35 uint64 = (v3 ^ (v3 / ((v3) | 1)))
                        _ = v35
                    default:
                        v2 += 1e10
                    }
                    switch (fi32 - (^(fi31))) {
                    case 0:
                        v36 := GenS2{F0: 118, F1: v3, F2: 65535}
                        _ = v36
                    case 127:
                        if (!false) {
                            v37 := GenS3{F0: (fi32 | fi32), F1: v3, F2: 128}
                            _ = v37
                        } else {
                            v38 := GenS4{F0: "foo bar", F1: (!(9223372036854775807 < 42))}
                            _ = v38
                        }
                    default:
                        var v39 uint16 = 333
                        _ = v39
                    }
                    v2 = 3.14159
                    v40 := GenS3{F0: (fi31 & fi31), F1: v3, F2: 2}
                    _ = v34
                    _ = v40
                } else {
                    v41 := GenS5{F0: -824}
                    _ = v41
                }
                var v42 float32 = 0.1
                _ = v42
            }
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    v1 = append(v1, 42)
    switch 4294967295 {
    case 2147483647:
        if (!(2147483647 != (v4 % ((1) | 1)))) {
            for fi43 := 0; fi43 < 16; fi43++ {
                v4 = (v4 * v4)
                _ = fi43
            }
            if true {
                for fi44 := 0; fi44 < 13; fi44++ {
                    _ = len(v1)
                    v45 := GenS5{F0: -1}
                    for fi46 := 0; fi46 < 11; fi46++ {
                        _ = len(v1)
                        var v47 int8 = 79
                        _ = fi46
                        _ = v47
                    }
                    _ = fi44
                    _ = v45
                }
            }
        } else {
            var v48 int = 32767
            for fi49 := 0; fi49 < 16; fi49++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi49
            }
            if false {
                for fi50 := 0; fi50 < 15; fi50++ {
                    switch ("hello" + "foo bar") {
                    case "\n\t":
                        v3 -= (v3 & (v3 * (v3 ^ v3)))
                    default:
                        var v51 float64 = (0.0 + (0.0 + (1.5 * 0.1)))
                        _ = v51
                    }
                    if false {
                        fmt.Println(fi50)
                    }
                    _ = fi50
                }
            } else {
                for fi52 := 0; fi52 < 16; fi52++ {
                    for fi53 := 0; fi53 < 11; fi53++ {
                        v1 = append(v1, 2147483648)
                        v54 := GenS6{F0: 192, F1: (^(fi52))}
                        _ = fi53
                        _ = v54
                    }
                    if (0 < 126) {
                        var v55 uint = 0
                        if (!true) {
                            if ((!(false || (-81 > 42))) && true) {
                                v3 = (v3 + 618)
                                var v56 bool = false
                                _ = v56
                            } else {
                                if true {
                                    for fi57 := 0; fi57 < 20; fi57++ {
                                        if (42 == 100) {
                                            if true {
                                                switch (fi57 + (fi52 >> ((v55) & 63))) {
                                                case -311:
                                                    v55 += 1
                                                case -337:
                                                    if false {
                                                        v4 += int32(v55)
                                                        if (!false) {
                                                            v3 = (v3 * 589)
                                                            v4 -= v4
                                                            v4 = -438
                                                        } else {
                                                            var v58 int64 = 9223372036854775807
                                                            v59 := GenS7{F0: -868, F1: 726}
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                    } else {
                                                        v2 = 1.5
                                                        v60 := GenS4{F0: "hello", F1: true}
                                                        v2 = v2
                                                        v55 -= 65535
                                                        _ = v60
                                                    }
                                                case -246:
                                                    v61 := GenS8{F0: fi57}
                                                    _ = v61
                                                default:
                                                    switch (v48 >> (((v55 | (v55 ^ 141))) & 63)) {
                                                    case 128:
                                                        _ = len(v1)
                                                    case 1:
                                                        if false {
                                                            var v62 uint32 = 4294967295
                                                            fmt.Println(v3)
                                                            _ = v62
                                                        } else {
                                                            var v63 int64 = 1
                                                            _ = v63
                                                        }
                                                    }
                                                }
                                                fmt.Println(v3)
                                                var v64 bool = (-9223372036854775808 > -1)
                                                if false {
                                                    var v65 int32 = (v4 | 42)
                                                    _ = v65
                                                } else {
                                                    v66 := GenS7{F0: -883, F1: 218}
                                                    _ = v66
                                                }
                                                _ = v64
                                            } else {
                                                if true {
                                                    if false {
                                                        v55 = v55
                                                        var v67 uint64 = (v3 << ((v55) & 63))
                                                        if (-32768 > -108) {
                                                            var v68 string = "a"
                                                            _ = v68
                                                        } else {
                                                            var v69 int = 32768
                                                            var v70 int16 = 42
                                                            switch v48 {
                                                            case 0:
                                                                v71 := GenS6{F0: v3, F1: -2147483648}
                                                                _ = v71
                                                            case 9223372036854775807:
                                                                if false {
                                                                    fmt.Println(v69, v4, v67)
                                                                } else {
                                                                    v72 := GenS9{F0: 2147483648, F1: (^(v70))}
                                                                    fmt.Println(v4, v2, fi52)
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v3 = v67
                                                                    case "a":
                                                                        v70 = (^(v70))
                                                                    }
                                                                    _ = v72
                                                                }
                                                            default:
                                                                switch "\n\t" {
                                                                case "":
                                                                    v67 = 1
                                                                case "a":
                                                                    var v73 int32 = (v4 ^ (v4 * v4))
                                                                    _ = v73
                                                                default:
                                                                    if true {
                                                                        v67 = v67
                                                                        var v74 int8 = 1
                                                                        _ = len(v1)
                                                                        switch "hello" {
                                                                        case "a":
                                                                            var v75 bool = (true && true)
                                                                            _ = v75
                                                                        default:
                                                                            v55 = (v55 + 120)
                                                                        }
                                                                        _ = v74
                                                                    }
                                                                }
                                                            }
                                                            _ = v69
                                                            _ = v70
                                                        }
                                                        if (!true) {
                                                            v76 := GenS10{F0: v4, F1: 1}
                                                            _ = v76
                                                        }
                                                        v77 := GenS9{F0: 42, F1: -914}
                                                        _ = v67
                                                        _ = v77
                                                    }
                                                } else {
                                                    if (2 <= 42) {
                                                        v2 += 1e10
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v78 := GenS6{F0: 42, F1: 928}
                                                            _ = v78
                                                        }
                                                    }
                                                    var v79 int64 = -144
                                                    v3 -= 76
                                                    _ = v79
                                                }
                                                fmt.Println(fi57)
                                            }
                                            fmt.Println(fi57, v48, v48)
                                            v1 = append(v1, 65535)
                                        }
                                        fmt.Println(v3, v2)
                                        _ = fi57
                                    }
                                    v48 = v48
                                    switch 42 {
                                    case 0:
                                        for fi80 := 0; fi80 < 1; fi80++ {
                                            var v81 int32 = (v4 | (v4 ^ -687))
                                            _ = fi80
                                            _ = v81
                                        }
                                    case 127:
                                        v82 := GenS7{F0: 32767, F1: 2147483647}
                                        _ = v82
                                    default:
                                        if true {
                                            for fi83 := 0; fi83 < 19; fi83++ {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v48 = 2147483647
                                                v2 += 1.5
                                                _ = fi83
                                            }
                                            if true {
                                                v84 := GenS11{F0: 146, F1: "hello"}
                                                if len(v1) > 0 {
                                                    v1[0] = 2
                                                }
                                                _ = v84
                                            } else {
                                                if false {
                                                    v1 = append(v1, 0)
                                                } else {
                                                    var v85 uint = 4294967295
                                                    for fi86 := 0; fi86 < 4; fi86++ {
                                                        switch (fi52 & int(v85)) {
                                                        case 197:
                                                            var v87 uint8 = 127
                                                            _ = v87
                                                        case 0:
                                                            fmt.Println(v3)
                                                        }
                                                        switch "hello" {
                                                        case "hello":
                                                            v4 -= v4
                                                            _ = len(v1)
                                                            v2 = 1.5
                                                        case "\n\t":
                                                            _ = len(v1)
                                                        default:
                                                            fmt.Println(v55)
                                                        }
                                                        v88 := GenS10{F0: 1, F1: 53}
                                                        _ = fi86
                                                        _ = v88
                                                    }
                                                    _ = v85
                                                }
                                                var v89 float32 = v2
                                                _ = v89
                                            }
                                        } else {
                                            for fi90 := 0; fi90 < 20; fi90++ {
                                                _ = len(v1)
                                                switch (fi52 + v48) {
                                                case 1:
                                                    var v91 uint8 = 0
                                                    _ = v91
                                                case 2147483647:
                                                    var v92 uint16 = 2
                                                    _ = v92
                                                }
                                                _ = fi90
                                            }
                                            v48 -= fi52
                                            v93 := GenS9{F0: 0, F1: 32767}
                                            _ = v93
                                        }
                                    }
                                }
                                var v94 int16 = 435
                                _ = len(v1)
                                _ = v94
                            }
                            v55 = 0
                        } else {
                            for fi95 := 0; fi95 < 15; fi95++ {
                                fmt.Println(v48, v48)
                                switch "hello" {
                                case "foo bar":
                                    if false {
                                        if (false && (643 <= (v3 % ((4294967295) | 1)))) {
                                            if (true && (false || ((v55 ^ 2) <= (v55 ^ (^(v55)))))) {
                                                var v96 uint8 = 255
                                                _ = v96
                                            } else {
                                                var v97 int = v48
                                                if ("foo bar" < "\n\t") {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = len(v1)
                                                    v3 = 380
                                                    if len(v1) > 0 {
                                                        v1[0] = 2147483647
                                                    }
                                                    v97 = (fi95 ^ int(v3))
                                                } else {
                                                    v2 = v2
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v98 uint = v55
                                                    _ = v98
                                                }
                                                _ = v97
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = 2147483647
                                            }
                                            v99 := GenS8{F0: (^(fi95))}
                                            var v100 uint = (v55 % ((v55) | 1))
                                            _ = v100
                                            _ = v99
                                        } else {
                                            v55 -= (v55 * v55)
                                            if true {
                                                if (-1 == -32768) {
                                                    if len(v1) > 0 {
                                                        v1[0] = 42
                                                    }
                                                    var v101 uint16 = 42
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v101
                                                } else {
                                                    fmt.Println(fi95, v55, fi95)
                                                }
                                            }
                                            v4 = 0
                                        }
                                        switch "" {
                                        case "hello":
                                            v3 -= v3
                                        case "foo bar":
                                            v2 = 1e10
                                        }
                                        if false {
                                            if len(v1) > 0 {
                                                v1[0] = 1
                                            }
                                        }
                                        var v102 string = "hello"
                                        var v103 uint32 = 2
                                        _ = v102
                                        _ = v103
                                    } else {
                                        v104 := GenS1{F0: 1, F1: (v2 <= (1.5 / ((v2 * (v2 - v2)) + v2)))}
                                        switch (v48 * (v48 - (fi95 / ((127) | 1)))) {
                                        case 662:
                                            fmt.Println(v2, fi52)
                                        case -2147483648:
                                            _ = len(v1)
                                        case 2147483647:
                                            fmt.Println(fi52, fi95)
                                        }
                                        v1 = append(v1, 65535)
                                        _ = v104
                                    }
                                case "a":
                                    fmt.Println(v55, v48)
                                case "hello":
                                    v2 += float32(1.5)
                                }
                                v48 = v48
                                switch ("a" + "hello") {
                                case "a":
                                    var v105 int32 = (v4 * v4)
                                    _ = v105
                                case "":
                                    switch (v48 << ((v55) & 63)) {
                                    case 4294967295:
                                        fmt.Println(fi52, fi52)
                                    case -786:
                                        v1 = append(v1, 563)
                                    case -849:
                                        fmt.Println(v2, fi95)
                                    default:
                                        if true {
                                            _ = len(v1)
                                        } else {
                                            fmt.Println(v3)
                                            v2 -= 1.5
                                        }
                                    }
                                }
                                _ = fi95
                            }
                            var v106 uint8 = 0
                            switch 0 {
                            case 0:
                                fmt.Println(v48)
                            case -2147483648:
                                fmt.Println(v2)
                            default:
                                v4 -= v4
                            }
                            _ = v106
                        }
                        _ = len(v1)
                        _ = v55
                    }
                    _ = fi52
                }
                if false {
                    v48 += v48
                    fmt.Println(v4)
                    fmt.Println(v4, v3, v4)
                } else {
                    var v107 uint16 = 298
                    v2 = v2
                    _ = v107
                }
                v108 := GenS12{F0: 786, F1: 2, F2: "a"}
                _ = v108
            }
            _ = v48
        }
    case -9223372036854775808:
        fmt.Println(v3, v2, v2)
    }
    var v109 uint8 = 128
    v110 := GenS13{F0: 42, F1: -2147483648}
    var v111 uint = 954
    v2 += (1.5 - v2)
    var v112 uint32 = 2147483648
    if true {
        for fi113 := 0; fi113 < 1; fi113++ {
            var v114 int16 = 0
            if true {
                var v115 bool = ((!false) || true)
                for fi116 := 0; fi116 < 5; fi116++ {
                    _ = len(v1)
                    _ = fi116
                }
                var v117 uint8 = v109
                _ = v115
                _ = v117
            }
            _ = fi113
            _ = v114
        }
        if (!true) {
            if false {
                v2 = v2
                v111 = 0
            } else {
                v118 := GenS14{F0: v4, F1: "", F2: false}
                _ = v118
            }
            var v119 int8 = 76
            fmt.Println(v4)
            _ = v119
        } else {
            for fi120 := 0; fi120 < 4; fi120++ {
                v4 = v4
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi120
            }
        }
        for fi121 := 0; fi121 < 7; fi121++ {
            v1 = append(v1, 1)
            _ = fi121
        }
    } else {
        v122 := GenS15{F0: "hello"}
        var v123 string = ""
        if false {
            var v124 int8 = int8(v111)
            if (319 <= -1) {
                if ((v4 % (((v4 << (((^(v111))) & 63))) | 1)) == int32(v3)) {
                    for fi125 := 0; fi125 < 19; fi125++ {
                        v111 -= (v111 + uint(v3))
                        _ = fi125
                    }
                }
                v126 := GenS9{F0: v112, F1: -32768}
                _ = v126
            } else {
                v127 := GenS6{F0: (v3 ^ 508), F1: -2147483648}
                v4 = (v4 >> (((v111 / ((v111) | 1))) & 63))
                var v128 int16 = 623
                _ = v127
                _ = v128
            }
            v111 -= (v111 / ((4294967295) | 1))
            _ = v124
        } else {
            v3 -= v3
            var v129 int16 = 32767
            _ = v129
        }
        _ = v122
        _ = v123
    }
    var v130 int64 = -132
    v111 = v111
    var v131 uint32 = (v112 >> (((v111 >> (((v111 >> (((v111 - v111)) & 63))) & 63))) & 63))
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v109:", v109)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v109
    _ = v110
    _ = v111
    _ = v112
    _ = v130
    _ = v131
    _ = v2
    _ = v3
    _ = v4
}
