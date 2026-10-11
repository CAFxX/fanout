package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 bool; F1 int; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int32; F1 int16; F2 uint16 }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int; F1 int }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 bool; F1 uint8 }
type GenS12 struct { F0 int16; F1 uint8; F2 int64 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint16; F1 int32; F2 uint32 }
func genFn0(p0 int, p1 uint32, p2 uint16) uint {
    switch p0 {
    case -9223372036854775808:
        v5 := []int{p0}
        _ = v5
    case 57:
        v6 := GenS1{F0: 255, F1: 1}
        _ = v6
    case 1:
        fmt.Println(p2, p2, p0)
    default:
        fmt.Println(p1, p0, p0)
    }
    switch "" {
    case "":
        if ((!(true && true)) || false) {
            switch p0 {
            case -953:
                var v7 uint16 = (^(p2))
                _ = v7
            case 1:
                for fi8 := 0; fi8 < 3; fi8++ {
                    switch "\n\t" {
                    case "hello":
                        v9 := GenS1{F0: 42, F1: 4294967295}
                        _ = v9
                    case "foo bar":
                        switch "hello" {
                        case "":
                            v10 := GenS1{F0: 563, F1: 4294967295}
                            _ = v10
                        case "foo bar":
                            for fi11 := 0; fi11 < 4; fi11++ {
                                p2 = (p2 / ((p2) | 1))
                                p1 -= (p1 - (p1 / ((65535) | 1)))
                                v12 := GenS1{F0: 40, F1: 255}
                                _ = fi11
                                _ = v12
                            }
                        case "hello":
                            p0 += (fi8 | (fi8 | -735))
                        }
                    }
                    _ = fi8
                }
            default:
                for fi13 := 0; fi13 < 14; fi13++ {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        v15 := []int{(fi14 + (fi14 - fi13))}
                        v16 := GenS1{F0: uint(p2), F1: 1}
                        fmt.Println(p2, p0)
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    p1 += 2
                    _ = fi13
                }
            }
            if (0 > int8(p0)) {
                var v17 uint64 = 832
                var v18 float32 = 3.14159
                _ = v17
                _ = v18
            } else {
                p2 -= (p2 << ((2) & 63))
                var v19 int32 = 42
                p2 = (p2 >> ((2) & 63))
                _ = v19
            }
        } else {
            v20 := GenS1{F0: 18446744073709551615, F1: 65535}
            switch "a" {
            case "hello":
                if (!true) {
                    var v21 bool = true
                    p1 -= (^(p1))
                    _ = v21
                }
            default:
                v22 := GenS1{F0: 636, F1: 6}
                _ = v22
            }
            v23 := []int{p0, (p0 / ((p0) | 1)), p0}
            _ = v20
            _ = v23
        }
    case "\n\t":
        v24 := GenS1{F0: 2, F1: 65535}
        _ = v24
    case "a":
        if true {
            if (false && false) {
                switch ("a" + ("" + "foo bar")) {
                case "\n\t":
                    var v25 int8 = 127
                    _ = v25
                case "":
                    v26 := []int{}
                    _ = v26
                }
                var v27 bool = ((127 > -2147483648) && (true && false))
                _ = v27
            } else {
                v28 := GenS1{F0: 466, F1: 2}
                _ = v28
            }
        } else {
            var v29 uint32 = p1
            p2 -= p2
            _ = v29
        }
    }
    v30 := []int{p0, 127, (p0 | p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 65535
}
func genFn1(p0 int8, p1 int16) uint8 {
    for fi31 := 0; fi31 < 16; fi31++ {
        p1 = 32767
        p1 = p1
        p0 -= (p0 ^ 42)
        _ = fi31
    }
    v32 := []int{4294967295}
    v33 := GenS1{F0: 303, F1: genFn0(32767, 0, 1)}
    for fi34 := 0; fi34 < 1; fi34++ {
        var v35 uint16 = 42
        v35 -= (^(v35))
        fmt.Println(p1, p1, v35)
        _ = fi34
        _ = v35
    }
    var v36 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v32
    _ = v33
    _ = v36
    return 42
}

func main() {
    var v1 float32 = 1e10
    var v2 uint64 = 42
    var v3 bool = (106 < 160)
    var v4 float32 = 0.0
    v37 := GenS1{F0: 4294967295, F1: 4294967295}
    if ((-1 <= 32767) && (!(-607 != -32768))) {
        for fi38 := 0; fi38 < 3; fi38++ {
            fmt.Println(v3, v4)
            for fi39 := 0; fi39 < 18; fi39++ {
                if v3 {
                    fmt.Println(v1, v3)
                    var v40 uint8 = 42
                    _ = v40
                } else {
                    v41 := []int{}
                    if ("foo bar" != ("hello" + ("a" + "\n\t"))) {
                        v4 += v4
                        if false {
                            var v42 uint32 = uint32(v2)
                            v43 := GenS1{F0: 42, F1: 1}
                            var v44 uint16 = 42
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                    }
                    fmt.Println(v2)
                    _ = v41
                }
                v45 := GenS1{F0: 18446744073709551615, F1: 4294967295}
                _ = fi39
                _ = v45
            }
            switch "a" {
            case "a":
                v46 := []int{}
                _ = v46
            case "foo bar":
                if (-813 < 1) {
                    if (18446744073709551615 >= (v2 >> ((18446744073709551615) & 63))) {
                        fmt.Println(fi38, v1)
                        v4 += -2.25
                        for fi47 := 0; fi47 < 2; fi47++ {
                            var v48 string = ("a" + "a")
                            v49 := GenS1{F0: genFn0(-9223372036854775808, 735, 772), F1: 255}
                            var v50 string = ""
                            _ = fi47
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                    } else {
                        switch (fi38 % (((fi38 % (((fi38 / ((fi38) | 1))) | 1))) | 1)) {
                        case 2147483647:
                            v51 := []int{2147483647}
                            _ = v51
                        case 806:
                            v52 := GenS2{F0: (v4 != v1), F1: 127}
                            _ = v52
                        default:
                            v53 := []int{}
                            _ = v53
                        }
                    }
                } else {
                    for fi54 := 0; fi54 < 17; fi54++ {
                        if (v3 && v3) {
                            var v55 int32 = 2147483647
                            if ((2 <= 2) || v3) {
                                var v56 int = fi54
                                switch 4294967295 {
                                case 1:
                                    var v57 int16 = -279
                                    _ = v57
                                case 42:
                                    if true {
                                        v2 = v2
                                        v1 = v1
                                        v58 := []int{fi38, -9223372036854775808}
                                        _ = v58
                                    }
                                case -1:
                                    switch 9223372036854775807 {
                                    case 1:
                                        if (v3 && ((v2 / (((^(v2))) | 1)) <= (v2 ^ v2))) {
                                            var v59 uint64 = (v2 << ((65535) & 63))
                                            _ = v59
                                        } else {
                                            var v60 int16 = 0
                                            var v61 uint8 = 127
                                            _ = v60
                                            _ = v61
                                        }
                                    default:
                                        v62 := []int{fi38, (v56 - 853)}
                                        _ = v62
                                    }
                                }
                                _ = v56
                            }
                            _ = v55
                        }
                        _ = fi54
                    }
                    v63 := []int{42, (fi38 | -9223372036854775808)}
                    _ = v63
                }
            case "hello":
                for fi64 := 0; fi64 < 1; fi64++ {
                    var v65 float64 = 0.0
                    switch fi64 {
                    case 4294967295:
                        var v66 string = "a"
                        _ = v66
                    default:
                        switch "hello" {
                        case "hello":
                            v67 := []int{(fi64 + 2147483647), fi38, fi38}
                            _ = v67
                        case "a":
                            v68 := []int{}
                            _ = v68
                        default:
                            v2 += (v2 ^ v2)
                        }
                    }
                    if v3 {
                        v69 := GenS2{F0: (!v3), F1: 128}
                        _ = v69
                    } else {
                        if false {
                            v70 := GenS1{F0: 859, F1: 18446744073709551615}
                            v65 = float64(0.1)
                            _ = v70
                        } else {
                            fmt.Println(v2, v65, v4)
                        }
                        var v71 int64 = 9223372036854775807
                        var v72 uint32 = 65535
                        if v3 {
                            var v73 bool = (false && (65535 != 1))
                            v74 := []int{(fi38 | (fi64 & fi38)), fi38, fi64, fi38}
                            var v75 uint = 1
                            _ = v73
                            _ = v74
                            _ = v75
                        } else {
                            v3 = (-133 <= 1000)
                            v76 := GenS3{F0: ((v72 | (^(v72))) == (v72 >> ((685) & 63))), F1: 42, F2: (v2 & (v2 ^ 42))}
                            v77 := GenS4{F0: int64(v72), F1: 201}
                            _ = v76
                            _ = v77
                        }
                        _ = v71
                        _ = v72
                    }
                    _ = fi64
                    _ = v65
                }
            }
            _ = fi38
        }
        if v3 {
            v4 += -2.25
            for fi78 := 0; fi78 < 5; fi78++ {
                switch "hello" {
                case "\n\t":
                    if v3 {
                        var v79 uint64 = v2
                        v80 := GenS5{F0: 42, F1: -32768, F2: 65535}
                        _ = v79
                        _ = v80
                    } else {
                        fmt.Println(v1)
                        v81 := GenS6{F0: -827, F1: 634}
                        _ = v81
                    }
                case "foo bar":
                    var v82 uint16 = 255
                    _ = v82
                }
                for fi83 := 0; fi83 < 1; fi83++ {
                    if ((genFn1(-128, 42) < 31) || true) {
                        fmt.Println(v1, v3)
                    } else {
                        if ((0 > 984) || false) {
                            switch int(v2) {
                            case -178:
                                if v3 {
                                    switch "" {
                                    case "\n\t":
                                        var v84 int64 = 237
                                        _ = v84
                                    default:
                                        if v3 {
                                            v4 -= float32(-2.25)
                                            v2 = (v2 ^ (v2 & (^(v2))))
                                            v85 := []int{(fi78 & 127)}
                                            _ = v85
                                        }
                                    }
                                    var v86 int16 = -32768
                                    _ = v86
                                } else {
                                    if (4294967295 < 2) {
                                        if (!(!v3)) {
                                            v87 := []int{}
                                            switch fi83 {
                                            case 605:
                                                var v88 string = "foo bar"
                                                _ = v88
                                            default:
                                                fmt.Println(v1, fi83, v3)
                                            }
                                            _ = v87
                                        } else {
                                            v89 := []int{(^(fi83)), (-(fi78)), 32767, fi83}
                                            _ = v89
                                        }
                                    } else {
                                        switch ("foo bar" + ("foo bar" + "hello")) {
                                        case "":
                                            fmt.Println(fi83, fi78)
                                        case "\n\t":
                                            v90 := GenS7{F0: -95}
                                            _ = v90
                                        case "foo bar":
                                            switch 9223372036854775807 {
                                            case 784:
                                                v4 = ((((v1 - v1) - (((v4 - 1.5) * v4) - 0.1)) * 1.5) * v1)
                                            case 128:
                                                var v91 int32 = 42
                                                _ = v91
                                            }
                                        default:
                                            fmt.Println(fi83)
                                        }
                                    }
                                    var v92 uint = 0
                                    _ = v92
                                }
                            case -825:
                                v3 = true
                            case 4294967295:
                                v93 := GenS7{F0: 0}
                                _ = v93
                            }
                        }
                    }
                    if (!true) {
                        v4 -= (v4 - 0.1)
                        v94 := GenS2{F0: v3, F1: 128}
                        switch ("hello" + ("hello" + ("" + "foo bar"))) {
                        case "\n\t":
                            if v3 {
                                v4 += 1e10
                                var v95 uint64 = (v2 / (((v2 | v2)) | 1))
                                _ = v95
                            } else {
                                v1 = v4
                                if true {
                                    var v96 int = (fi83 ^ 548)
                                    _ = v96
                                } else {
                                    if v3 {
                                        if ((901 <= -1) || (!(true && ((2147483647 >= -817) || true)))) {
                                            switch "foo bar" {
                                            case "foo bar":
                                                v2 = 1
                                                fmt.Println(v2, v4, v1)
                                                v97 := []int{}
                                                _ = v97
                                            }
                                        }
                                        v98 := []int{fi78, fi78}
                                        v4 += (1e10 / v4)
                                        _ = v98
                                    } else {
                                        v99 := []int{(fi78 - (fi78 | fi83)), (fi78 % ((fi78) | 1)), fi83}
                                        var v100 uint64 = v2
                                        if (!v3) {
                                            v1 = (((v1 - -2.25) - v1) / 0.1)
                                            v101 := GenS8{F0: 42}
                                            _ = v101
                                        }
                                        switch "foo bar" {
                                        case "foo bar":
                                            switch (fi78 * fi78) {
                                            case -134:
                                                if v3 {
                                                    v102 := GenS1{F0: 1, F1: 0}
                                                    switch (fi83 * (fi83 % ((-5) | 1))) {
                                                    case 4294967295:
                                                        if true {
                                                            if len(v99) > 0 {
                                                                v99[0] = 32768
                                                            }
                                                            v4 = ((v1 / (((v4 / 0.1) - v4) - (v4 / v1))) / v1)
                                                        }
                                                    case 548:
                                                        v1 += -2.25
                                                    case 9223372036854775807:
                                                        v103 := GenS3{F0: (!v3), F1: fi83, F2: (v100 % ((v100) | 1))}
                                                        fmt.Println(v3, v3)
                                                        if false {
                                                            var v104 uint = 0
                                                            switch ("a" + "a") {
                                                            case "hello":
                                                                switch (("hello" + "\n\t") + ((("\n\t" + "a") + "hello") + "\n\t")) {
                                                                case "\n\t":
                                                                    v100 = v2
                                                                case "":
                                                                    var v105 int32 = 164
                                                                    _ = v105
                                                                }
                                                            case "a":
                                                                var v106 float32 = v1
                                                                _ = v106
                                                            default:
                                                                if len(v99) > 0 {
                                                                    v99[0] = -1
                                                                }
                                                            }
                                                            _ = len(v99)
                                                            _ = v104
                                                        }
                                                        _ = v103
                                                    default:
                                                        var v107 string = ("a" + "a")
                                                        _ = v107
                                                    }
                                                    _ = v102
                                                } else {
                                                    switch ("hello" + "a") {
                                                    case "hello":
                                                        v1 = 0.0
                                                    case "":
                                                        switch "a" {
                                                        case "hello":
                                                            v108 := GenS9{F0: (fi78 - -714), F1: (fi83 - (fi83 >> ((genFn0(fi78, 0, 189)) & 63)))}
                                                            _ = v108
                                                        case "a":
                                                            var v109 bool = true
                                                            _ = v109
                                                        case "":
                                                            v110 := GenS10{F0: 128, F1: 127}
                                                            _ = v110
                                                        default:
                                                            var v111 int64 = 238
                                                            _ = v111
                                                        }
                                                    }
                                                    v112 := GenS11{F0: (!(916 >= 255)), F1: 136}
                                                    v113 := GenS7{F0: 797}
                                                    _ = v112
                                                    _ = v113
                                                }
                                                v100 += (^(v2))
                                            case -9223372036854775808:
                                                v2 += 4294967295
                                            }
                                        case "a":
                                            switch fi83 {
                                            case 9223372036854775807:
                                                v114 := GenS5{F0: 2147483647, F1: 32767, F2: 21}
                                                _ = v114
                                            case 42:
                                                switch "" {
                                                case "\n\t":
                                                    if len(v99) > 0 {
                                                        v99[0] = (fi83 & fi83)
                                                    }
                                                case "":
                                                    fmt.Println(v4, v1, v2)
                                                default:
                                                    if len(v99) > 0 {
                                                        v99[0] = (fi83 * fi83)
                                                    }
                                                }
                                            default:
                                                v1 = (1.5 + float32(-2.25))
                                            }
                                        }
                                        _ = v100
                                        _ = v99
                                    }
                                }
                                switch (("a" + "") + "hello") {
                                case "\n\t":
                                    v115 := []int{fi78, -2147483648, (fi83 % (((fi83 | fi78)) | 1)), fi78}
                                    _ = v115
                                case "hello":
                                    v3 = (-437 <= -900)
                                case "":
                                    switch (fi83 >> ((0) & 63)) {
                                    case -2147483648:
                                        v116 := GenS7{F0: -793}
                                        _ = v116
                                    default:
                                        var v117 string = ("" + (("foo bar" + "") + "a"))
                                        _ = v117
                                    }
                                default:
                                    var v118 string = "hello"
                                    switch -9223372036854775808 {
                                    case 32767:
                                        var v119 int8 = 0
                                        _ = v119
                                    }
                                    _ = v118
                                }
                            }
                        default:
                            var v120 int16 = -32768
                            _ = v120
                        }
                        v121 := GenS12{F0: 0, F1: 42, F2: 292}
                        _ = v121
                        _ = v94
                    }
                    _ = fi83
                }
                _ = fi78
            }
        }
    } else {
        var v122 uint16 = 42
        v123 := GenS13{F0: -644}
        _ = v122
        _ = v123
    }
    var v124 uint32 = 2147483648
    for fi125 := 0; fi125 < 3; fi125++ {
        v126 := []int{}
        if ((27 <= 1) || v3) {
            switch fi125 {
            case 2147483647:
                var v127 uint16 = 255
                _ = v127
            default:
                var v128 float64 = (3.14159 - -2.25)
                _ = v128
            }
        } else {
            if len(v126) > 0 {
                v126[0] = (-(fi125))
            }
            v1 = ((1e10 - 0.0) * v1)
        }
        _ = fi125
        _ = v126
    }
    v4 = v4
    v129 := GenS1{F0: 0, F1: 751}
    var v130 int = 32767
    switch v130 {
    case -2147483648:
        if ((0.1 >= 1e10) && v3) {
            fmt.Println(v130)
            v124 -= (v124 / ((v124) | 1))
        }
    case 229:
        for fi131 := 0; fi131 < 8; fi131++ {
            if v3 {
                v4 = (((v1 / ((float32(1e10) + (v1 / (0.1 / v4))) * v4)) + (3.14159 - (v4 - float32(((0.1 - (3.14159 + 1.5)) + float64(float32(0.1))))))) + 0.0)
                v2 += v2
            }
            _ = fi131
        }
    case 661:
        v130 -= (v130 | v130)
    default:
        for fi132 := 0; fi132 < 7; fi132++ {
            v124 = v124
            _ = fi132
        }
    }
    for fi133 := 0; fi133 < 16; fi133++ {
        var v134 float32 = 3.14159
        v3 = (!true)
        v135 := GenS14{F0: 124, F1: 266, F2: 1}
        _ = fi133
        _ = v134
        _ = v135
    }
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v124
    _ = v129
    _ = v130
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
