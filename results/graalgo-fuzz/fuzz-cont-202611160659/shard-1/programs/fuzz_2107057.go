package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 int }
type GenS3 struct { F0 int8; F1 uint32; F2 int64 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint; F1 bool; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint8; F1 string }
type GenS12 struct { F0 uint16; F1 int32 }
type GenS13 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS14 struct { F0 uint; F1 int16; F2 int16 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int; F1 int8 }
type GenS17 struct { F0 bool; F1 int32 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 string; F1 uint }
type GenS20 struct { F0 uint8; F1 uint8; F2 uint }
type GenS21 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 int = 732
    var v2 float32 = 1.5
    v3 := make(map[string]uint32)
    var v4 uint = 42
    var v5 int16 = -32768
    v6 := make(map[int]string)
    v1 -= v1
    v5 -= v5
    if false {
        v2 = v2
        var v7 int64 = -9223372036854775808
        v2 = (v2 * 1e10)
        _ = v7
    }
    if (false || ((!((!(!false)) && (255 == 563))) && true)) {
        v2 += float32(-2.25)
    } else {
        fmt.Println(v4)
        fmt.Println(v5)
        if (v4 <= (v4 - v4)) {
            switch -64 {
            case 9223372036854775807:
                v2 = (float32(1.5) - v2)
            case -2147483648:
                switch "a" {
                case "foo bar":
                    v4 += (v4 & v4)
                case "hello":
                    if (false || false) {
                        for fi8 := 0; fi8 < 2; fi8++ {
                            if (-2147483648 > 1) {
                                var v9 string = "foo bar"
                                if (!false) {
                                    switch (v1 ^ (v1 >> (((v4 & v4)) & 63))) {
                                    case 2147483647:
                                        var v10 int16 = (v5 ^ v5)
                                        _ = v10
                                    }
                                    var v11 float64 = 1e10
                                    if true {
                                        if (!((552 == 2147483647) && (((true && true) && false) && true))) {
                                            if true {
                                                v1 = (v1 << (((v4 << ((0) & 63))) & 63))
                                                v11 = v11
                                            } else {
                                                if false {
                                                    var v12 uint64 = 18446744073709551615
                                                    _ = v12
                                                } else {
                                                    v13 := GenS1{F0: 255, F1: 255}
                                                    v4 = 20
                                                    v14 := GenS1{F0: 0, F1: 2}
                                                    _ = v13
                                                    _ = v14
                                                }
                                                if (!((float32(float64(v2)) + -2.25) < v2)) {
                                                    for fi15 := 0; fi15 < 18; fi15++ {
                                                        v4 = (^(v4))
                                                        var v16 int16 = v5
                                                        if ("hello" == v9) {
                                                            v4 -= v4
                                                            var v17 int32 = 42
                                                            var v18 uint32 = 603
                                                            _ = v17
                                                            _ = v18
                                                        }
                                                        _ = fi15
                                                        _ = v16
                                                    }
                                                    switch (v1 ^ fi8) {
                                                    case -304:
                                                        _ = len(v3)
                                                    case 9223372036854775807:
                                                        if (!true) {
                                                            var v19 uint64 = 926
                                                            for fi20 := 0; fi20 < 13; fi20++ {
                                                                var v21 string = v9
                                                                var v22 uint16 = 65535
                                                                _ = fi20
                                                                _ = v21
                                                                _ = v22
                                                            }
                                                            for fi23 := 0; fi23 < 4; fi23++ {
                                                                var v24 uint32 = 2147483648
                                                                switch "\n\t" {
                                                                case "":
                                                                    var v25 uint8 = 0
                                                                    _ = v25
                                                                case "\n\t":
                                                                    var v26 float32 = (((v2 / (0.1 * v2)) * (v2 - v2)) * 3.14159)
                                                                    _ = v26
                                                                default:
                                                                    v27 := GenS1{F0: uint8(v24), F1: uint16(v24)}
                                                                    v2 += v2
                                                                    v28 := GenS2{F0: v5, F1: (v5 & v5), F2: 9223372036854775807}
                                                                    if true {
                                                                        if true {
                                                                            fmt.Println(v1, v11)
                                                                            var v29 int8 = -94
                                                                            fmt.Println(v29, v4)
                                                                            _ = v29
                                                                        }
                                                                        var v30 int64 = 9223372036854775807
                                                                        var v31 float64 = (float64((0.0 * float32(3.14159))) / 1e10)
                                                                        _ = v30
                                                                        _ = v31
                                                                    } else {
                                                                        var v32 int = (fi8 / ((fi23) | 1))
                                                                        v1 -= fi8
                                                                        v33 := []int{-31, (v32 / ((v32) | 1))}
                                                                        _ = v32
                                                                        _ = v33
                                                                    }
                                                                    v34 := GenS3{F0: 0, F1: v24, F2: -4}
                                                                    _ = v27
                                                                    _ = v28
                                                                    _ = v34
                                                                }
                                                                _ = fi23
                                                                _ = v24
                                                            }
                                                            _ = v19
                                                        }
                                                    default:
                                                        v2 -= (0.1 - v2)
                                                    }
                                                    var v35 float64 = -2.25
                                                    _ = v35
                                                } else {
                                                    fmt.Println(v2, v4)
                                                    v5 -= (v5 << (((^(v4))) & 63))
                                                }
                                                if true {
                                                    v36 := []int{fi8, 9223372036854775807}
                                                    v5 = 42
                                                    _ = v36
                                                } else {
                                                    fmt.Println(v1)
                                                }
                                            }
                                            fmt.Println(v5, v9)
                                            var v37 float32 = (((v2 * (v2 - 0.1)) - float32(v11)) + v2)
                                            _ = v37
                                        }
                                    }
                                    _ = v11
                                }
                                _ = v9
                            } else {
                                var v38 string = "\n\t"
                                if false {
                                    var v39 uint = 42
                                    _ = v39
                                } else {
                                    v1 = 32768
                                    if (1 <= 901) {
                                        if (true && ((256 == 2) || false)) {
                                            if ("\n\t" == v38) {
                                                v40 := GenS2{F0: 814, F1: (^(v5)), F2: fi8}
                                                for fi41 := 0; fi41 < 8; fi41++ {
                                                    v2 -= (1e10 + (v2 - v2))
                                                    v42 := []int{}
                                                    _ = fi41
                                                    _ = v42
                                                }
                                                v2 = v2
                                                _ = v40
                                            } else {
                                                if (false && true) {
                                                    v43 := []int{(v1 * 32768), fi8, v1}
                                                    var v44 float64 = 3.14159
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    switch (fi8 | fi8) {
                                                    case 9223372036854775807:
                                                        var v45 float32 = v2
                                                        _ = v45
                                                    case -957:
                                                        var v46 string = (v38 + ("hello" + v38))
                                                        _ = v46
                                                    case 18:
                                                        for fi47 := 0; fi47 < 13; fi47++ {
                                                            v38 = ("" + v38)
                                                            var v48 float32 = ((((v2 - v2) * 0.1) / (v2 - float32(1.5))) + v2)
                                                            _ = fi47
                                                            _ = v48
                                                        }
                                                    }
                                                    switch (v38 + (v38 + v38)) {
                                                    case "":
                                                        if true {
                                                            fmt.Println(v2)
                                                            v49 := GenS3{F0: 0, F1: 592, F2: -1}
                                                            for fi50 := 0; fi50 < 5; fi50++ {
                                                                v38 = "foo bar"
                                                                v4 += (v4 ^ 0)
                                                                if (((0.1 * (float64(float32(1e10)) - -2.25)) - float64(v2)) < (((1e10 - (0.0 + 0.0)) + -2.25) / 1e10)) {
                                                                    if true {
                                                                        var v51 uint = (v4 + (v4 & 250))
                                                                        _ = v51
                                                                    }
                                                                    var v52 int8 = -47
                                                                    var v53 uint8 = 234
                                                                    fmt.Println(fi8)
                                                                    _ = v52
                                                                    _ = v53
                                                                }
                                                                _ = fi50
                                                            }
                                                            _ = v49
                                                        } else {
                                                            var v54 string = ""
                                                            _ = v54
                                                        }
                                                    case "\n\t":
                                                        if false {
                                                            var v55 float64 = ((-2.25 - 1e10) / (float64(float32(3.14159)) * 1.5))
                                                            _ = v55
                                                        } else {
                                                            fmt.Println(v38)
                                                            if true {
                                                                for fi56 := 0; fi56 < 8; fi56++ {
                                                                    switch (fi8 + fi56) {
                                                                    case 881:
                                                                        v1 = 32768
                                                                    case 42:
                                                                        v38 = "foo bar"
                                                                    }
                                                                    v57 := []int{v1, (-(fi8)), fi56, fi56}
                                                                    _ = fi56
                                                                    _ = v57
                                                                }
                                                            } else {
                                                                if false {
                                                                    v58 := GenS4{F0: int(v4), F1: int32(v5)}
                                                                    var v59 int16 = v5
                                                                    _ = v58
                                                                    _ = v59
                                                                }
                                                            }
                                                        }
                                                    case "hello":
                                                        v2 -= (((v2 * v2) / v2) * v2)
                                                    default:
                                                        var v60 int32 = 754
                                                        _ = v60
                                                    }
                                                }
                                                fmt.Println(v5, v4, v38)
                                                var v61 int64 = -9223372036854775808
                                                _ = v61
                                            }
                                            v62 := GenS5{F0: 255, F1: false, F2: v5}
                                            v5 -= v5
                                            _ = v62
                                        } else {
                                            v5 += (v5 + v5)
                                            for fi63 := 0; fi63 < 19; fi63++ {
                                                var v64 uint = 0
                                                _ = len(v3)
                                                _ = fi63
                                                _ = v64
                                            }
                                            v38 = v38
                                        }
                                        v2 = 1e10
                                        v38 = v38
                                    }
                                }
                                v6[(fi8 * (fi8 ^ (v1 % ((v1) | 1))))] = v38
                                _ = v38
                            }
                            for fi65 := 0; fi65 < 16; fi65++ {
                                v1 = (fi8 & -767)
                                v1 = fi8
                                _ = fi65
                            }
                            _ = fi8
                        }
                    } else {
                        if (false && false) {
                            switch ("foo bar" + "\n\t") {
                            case "foo bar":
                                v1 = 32767
                            case "":
                                v66 := []int{-9223372036854775808, 2147483647}
                                _ = v66
                            case "\n\t":
                                for fi67 := 0; fi67 < 6; fi67++ {
                                    switch ("" + "\n\t") {
                                    case "foo bar":
                                        v68 := GenS4{F0: v1, F1: 2147483647}
                                        _ = v68
                                    case "hello":
                                        v69 := GenS4{F0: fi67, F1: 1}
                                        _ = v69
                                    case "a":
                                        var v70 int32 = 123
                                        _ = v70
                                    }
                                    _ = fi67
                                }
                            default:
                                v4 -= (v4 * 18446744073709551615)
                            }
                            v1 = -587
                            if (("\n\t" + ("\n\t" + ("a" + ("" + "\n\t")))) > "\n\t") {
                                if ((v2 / float32(0.1)) >= 1e10) {
                                    v2 = v2
                                    v5 -= v5
                                }
                                v71 := []int{1}
                                if len(v71) > 0 {
                                    v71[0] = (-(v1))
                                }
                                _ = v71
                            }
                        } else {
                            switch "a" {
                            case "foo bar":
                                if true {
                                    var v72 int8 = -128
                                    switch 1 {
                                    case -1:
                                        fmt.Println(v2, v72, v1)
                                    case 199:
                                        fmt.Println(v4)
                                    case 130:
                                        var v73 string = ""
                                        _ = v73
                                    default:
                                        v1 = 13
                                    }
                                    _ = v72
                                }
                            default:
                                v5 += v5
                            }
                        }
                    }
                case "":
                    var v74 float32 = (v2 * 1e10)
                    _ = v74
                }
            default:
                v5 += v5
            }
            var v75 float64 = (1.5 - float64((1e10 * v2)))
            v76 := GenS1{F0: 127, F1: 596}
            _ = v75
            _ = v76
        } else {
            for fi77 := 0; fi77 < 4; fi77++ {
                fmt.Println(v2)
                _ = fi77
            }
            for fi78 := 0; fi78 < 17; fi78++ {
                if true {
                    switch -969 {
                    case 42:
                        switch "hello" {
                        case "hello":
                            v5 += 1
                        case "\n\t":
                            var v79 string = ""
                            _ = v79
                        default:
                            if (2147483647 >= 2147483647) {
                                if true {
                                    switch (-(fi78)) {
                                    case 4294967295:
                                        v80 := GenS4{F0: (fi78 >> ((v4) & 63)), F1: 42}
                                        _ = v80
                                    }
                                    v5 = v5
                                }
                            }
                        }
                    case -9223372036854775808:
                        v5 -= 0
                    case -2147483648:
                        v81 := GenS6{F0: (v4 ^ v4)}
                        _ = v81
                    default:
                        var v82 float64 = (-2.25 * -2.25)
                        _ = v82
                    }
                    fmt.Println(fi78)
                    for fi83 := 0; fi83 < 4; fi83++ {
                        fmt.Println(v4, v2)
                        _ = fi83
                    }
                }
                _ = fi78
            }
            _ = len(v6)
        }
    }
    v4 = v4
    switch v1 {
    case 4294967295:
        fmt.Println(v4, v2, v1)
    case 543:
        switch (("hello" + "foo bar") + ("a" + "hello")) {
        case "hello":
            v2 += v2
        case "\n\t":
            var v84 int32 = 42
            _ = v84
        case "a":
            if (1 != 248) {
                for fi85 := 0; fi85 < 15; fi85++ {
                    var v86 uint = 18446744073709551615
                    v1 += fi85
                    _ = fi85
                    _ = v86
                }
                for fi87 := 0; fi87 < 17; fi87++ {
                    if false {
                        v5 = v5
                        if true {
                            if ((^(v4)) == 18446744073709551615) {
                                v88 := GenS1{F0: 1, F1: 0}
                                _ = v88
                            }
                            for fi89 := 0; fi89 < 4; fi89++ {
                                fmt.Println(v1, fi87)
                                if false {
                                    if true {
                                        var v90 int8 = -44
                                        _ = v90
                                    } else {
                                        var v91 float64 = 1e10
                                        _ = v91
                                    }
                                    switch (fi87 | fi89) {
                                    case -735:
                                        var v92 int64 = 0
                                        _ = v92
                                    case 1:
                                        v93 := []int{2147483647, fi87, (fi87 >> ((v4) & 63)), v1}
                                        _ = v93
                                    }
                                    v4 = 1
                                } else {
                                    v94 := []int{-205}
                                    v95 := v3["\n\t"]
                                    v2 = 1e10
                                    _ = v94
                                    _ = v95
                                }
                                _ = fi89
                            }
                            if true {
                                switch 2147483647 {
                                case -279:
                                    var v96 uint16 = 65535
                                    _ = v96
                                case 4294967295:
                                    fmt.Println(v2)
                                case 9223372036854775807:
                                    delete(v3, "a")
                                }
                                v97 := []int{v1, 42}
                                _ = v97
                            } else {
                                fmt.Println(v5, v1)
                                var v98 int32 = 42
                                var v99 uint16 = 255
                                _ = v98
                                _ = v99
                            }
                        } else {
                            v4 = (v4 >> ((4294967295) & 63))
                            v4 = 558
                            var v100 uint64 = 0
                            _ = v100
                        }
                        var v101 uint = 2
                        _ = v101
                    }
                    var v102 float32 = -2.25
                    for fi103 := 0; fi103 < 15; fi103++ {
                        if (1 > (v4 * (v4 ^ v4))) {
                            delete(v6, -2147483648)
                            v4 -= (^(v4))
                            fmt.Println(v5, v5, v4)
                        }
                        v2 = float32(-2.25)
                        var v104 uint64 = 827
                        fmt.Println(fi87, v104)
                        var v105 bool = true
                        _ = fi103
                        _ = v104
                        _ = v105
                    }
                    _ = fi87
                    _ = v102
                }
                v106 := GenS7{F0: (^(v4))}
                _ = v106
            }
        }
    case 2147483647:
        v2 -= float32((3.14159 + 3.14159))
    default:
        if (!false) {
            for fi107 := 0; fi107 < 2; fi107++ {
                if (670 <= 483) {
                    v108 := GenS2{F0: (v5 + (v5 + (-(v5)))), F1: (v5 & (v5 ^ v5)), F2: v1}
                    var v109 float64 = 1e10
                    v110 := GenS8{F0: false}
                    _ = v108
                    _ = v109
                    _ = v110
                } else {
                    if true {
                        v5 += (v5 * v5)
                        v1 = (v1 & (fi107 >> (((v4 ^ (v4 | v4))) & 63)))
                        var v111 bool = true
                        _ = v111
                    }
                    v112 := GenS4{F0: (v1 - int(v4)), F1: -2147483648}
                    v113 := GenS7{F0: v4}
                    _ = v112
                    _ = v113
                }
                fmt.Println(v4, v5)
                for fi114 := 0; fi114 < 19; fi114++ {
                    if false {
                        fmt.Println(v1, fi114, fi114)
                        v3[(("\n\t" + ("foo bar" + "\n\t")) + ("\n\t" + "hello"))] = 632
                        switch ("hello" + "a") {
                        case "hello":
                            switch -308 {
                            case 32767:
                                v5 += v5
                            case 371:
                                fmt.Println(fi107, fi107)
                            }
                        case "a":
                            switch "foo bar" {
                            case "foo bar":
                                v115 := []int{fi114, fi107, int(v5), fi107}
                                _ = v115
                            default:
                                v5 = v5
                                var v116 float32 = 0.0
                                _ = v116
                            }
                        case "foo bar":
                            switch ("" + "foo bar") {
                            case "":
                                switch (v1 + fi114) {
                                case 128:
                                    var v117 bool = ((false || (0 == 102)) || (!(false && (false && true))))
                                    _ = v117
                                case 1:
                                    switch fi114 {
                                    case -865:
                                        var v118 bool = false
                                        _ = v118
                                    case 263:
                                        var v119 uint16 = 0
                                        _ = v119
                                    case 32768:
                                        v3[("" + ((("foo bar" + "a") + "hello") + ""))] = 266
                                    default:
                                        v120 := GenS4{F0: (v1 & fi107), F1: 1}
                                        _ = v120
                                    }
                                }
                                fmt.Println(fi107, fi114)
                            case "hello":
                                v121 := GenS9{F0: -128, F1: -1}
                                _ = v121
                            case "a":
                                if true {
                                    fmt.Println(fi114)
                                    v122 := GenS10{F0: 0}
                                    var v123 uint = (v4 - 599)
                                    _ = v122
                                    _ = v123
                                }
                            }
                        default:
                            switch fi114 {
                            case 4294967295:
                                fmt.Println(v2, v5, fi114)
                            case 32767:
                                switch 0 {
                                case 127:
                                    v124 := []int{(fi114 >> ((65535) & 63))}
                                    v125 := GenS8{F0: false}
                                    _ = v124
                                    _ = v125
                                case 42:
                                    v126 := GenS11{F0: 253, F1: ("\n\t" + ("hello" + "\n\t"))}
                                    _ = v126
                                case 128:
                                    if (488 != 9223372036854775807) {
                                        v1 += -785
                                        if (592 != 138) {
                                            var v127 int8 = 0
                                            _ = v127
                                        } else {
                                            v128 := GenS10{F0: 487}
                                            switch ("hello" + "") {
                                            case "a":
                                                v4 = v4
                                            case "\n\t":
                                                if false {
                                                    var v129 int = fi107
                                                    switch (fi114 + fi107) {
                                                    case -481:
                                                        v2 = 1e10
                                                    case -400:
                                                        v2 -= v2
                                                    default:
                                                        var v130 uint = v4
                                                        _ = v130
                                                    }
                                                    v2 -= (0.1 * 0.0)
                                                    _ = v129
                                                } else {
                                                    var v131 int8 = 100
                                                    if true {
                                                        v132 := GenS5{F0: (v4 ^ (v4 % (((v4 & v4)) | 1))), F1: false, F2: (-(v5))}
                                                        switch fi114 {
                                                        case 32768:
                                                            if false {
                                                                var v133 int16 = v5
                                                                fmt.Println(v2)
                                                                var v134 int32 = -2147483648
                                                                v1 -= fi114
                                                                _ = v133
                                                                _ = v134
                                                            } else {
                                                                v135 := []int{(fi114 << ((v4) & 63)), (-(v1)), (fi114 * fi114)}
                                                                if (4294967295 != 4294967295) {
                                                                    var v136 bool = ((fi107 & fi114) == 9223372036854775807)
                                                                    v4 = 1
                                                                    if ((9223372036854775807 != -9223372036854775808) || v136) {
                                                                        if false {
                                                                            var v137 uint16 = 251
                                                                            fmt.Println(v136, v2)
                                                                            switch ("foo bar" + (("hello" + "foo bar") + "hello")) {
                                                                            case "\n\t":
                                                                                if (v136 || (v136 && true)) {
                                                                                    if v136 {
                                                                                        v138 := GenS12{F0: 65535, F1: 1}
                                                                                        _ = v138
                                                                                    }
                                                                                }
                                                                            case "a":
                                                                                var v139 bool = true
                                                                                _ = v139
                                                                            case "":
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    fmt.Println(v131)
                                                                                case "hello":
                                                                                    fmt.Println(v1)
                                                                                    if (0.0 > 1e10) {
                                                                                        var v140 string = "foo bar"
                                                                                        v141 := GenS1{F0: 178, F1: 303}
                                                                                        v140 = (v140 + v140)
                                                                                        _ = v140
                                                                                        _ = v141
                                                                                    } else {
                                                                                        switch ((("hello" + ("" + "foo bar")) + ((("a" + "foo bar") + "hello") + ("hello" + "foo bar"))) + "") {
                                                                                        case "\n\t":
                                                                                            if (!v136) {
                                                                                                var v142 uint16 = v137
                                                                                                switch -2147483648 {
                                                                                                case 32768:
                                                                                                    var v143 bool = v136
                                                                                                    _ = v143
                                                                                                default:
                                                                                                    var v144 float64 = (1.5 - 0.1)
                                                                                                    _ = v144
                                                                                                }
                                                                                                v145 := GenS13{F0: int32(v5), F1: 694, F2: (^(v142))}
                                                                                                _ = v142
                                                                                                _ = v145
                                                                                            } else {
                                                                                                v146 := GenS8{F0: true}
                                                                                                _ = v146
                                                                                            }
                                                                                        case "foo bar":
                                                                                            fmt.Println(v137, v137)
                                                                                        default:
                                                                                            _ = len(v135)
                                                                                        }
                                                                                    }
                                                                                case "a":
                                                                                    var v147 uint64 = 154
                                                                                    _ = v147
                                                                                }
                                                                            default:
                                                                                var v148 uint32 = 0
                                                                                _ = v148
                                                                            }
                                                                            var v149 uint8 = 255
                                                                            _ = v137
                                                                            _ = v149
                                                                        }
                                                                        fmt.Println(fi114)
                                                                        v4 += (v4 / ((v4) | 1))
                                                                    } else {
                                                                        if (v5 >= (v5 / ((v5) | 1))) {
                                                                            if (32767 <= 2) {
                                                                                if len(v135) >= 2 {
                                                                                    v135 = v135[1:len(v135)]
                                                                                }
                                                                            } else {
                                                                                v150 := GenS14{F0: (v4 % ((828) | 1)), F1: -32768, F2: int16(v1)}
                                                                                delete(v3, "\n\t")
                                                                                if false {
                                                                                    var v151 bool = ("foo bar" >= "foo bar")
                                                                                    _ = v151
                                                                                }
                                                                                _ = v150
                                                                            }
                                                                            fmt.Println(v4, v136, v2)
                                                                            switch (v1 / ((289) | 1)) {
                                                                            case 32768:
                                                                                if ((v5 < 1) || (-1 <= 0)) {
                                                                                    v3[""] = uint32(v4)
                                                                                    fmt.Println(v4, v136, fi107)
                                                                                    v152 := GenS4{F0: fi107, F1: 1}
                                                                                    _ = v152
                                                                                }
                                                                            case 32767:
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    v1 = (-(v1))
                                                                                default:
                                                                                    switch (("foo bar" + "") + (("\n\t" + ("foo bar" + "\n\t")) + (("hello" + "") + ("hello" + "")))) {
                                                                                    case "\n\t":
                                                                                        var v153 float32 = float32(1e10)
                                                                                        _ = v153
                                                                                    case "foo bar":
                                                                                        v154 := GenS3{F0: -1, F1: 0, F2: -9223372036854775808}
                                                                                        _ = v154
                                                                                    default:
                                                                                        var v155 uint64 = 2
                                                                                        _ = v155
                                                                                    }
                                                                                }
                                                                            case -1:
                                                                                if ((fi114 * (fi107 << (((^(v4))) & 63))) == (v1 << ((v4) & 63))) {
                                                                                    if (-295 > v5) {
                                                                                        switch (fi107 ^ (fi114 & fi114)) {
                                                                                        case -462:
                                                                                            fmt.Println(fi114, v2)
                                                                                        case -2147483648:
                                                                                            var v156 int8 = (v131 >> (((v4 | v4)) & 63))
                                                                                            _ = v156
                                                                                        }
                                                                                        var v157 int8 = -122
                                                                                        if (60 > 4294967295) {
                                                                                            v158 := GenS14{F0: (v4 >> (((v4 / (((v4 << ((v4) & 63))) | 1))) & 63)), F1: 523, F2: (v5 ^ (v5 - 999))}
                                                                                            _ = v158
                                                                                        } else {
                                                                                            fmt.Println(fi114)
                                                                                            v1 = -2147483648
                                                                                            v2 -= (v2 / ((v2 * 1e10) + (v2 / (float32((0.1 + 1.5)) / 1e10))))
                                                                                        }
                                                                                        v1 = 32768
                                                                                        v157 += v131
                                                                                        _ = v157
                                                                                    } else {
                                                                                        var v159 int32 = -2147483648
                                                                                        _ = v159
                                                                                    }
                                                                                    if v136 {
                                                                                        v160 := GenS6{F0: v4}
                                                                                        switch fi107 {
                                                                                        case 1:
                                                                                            if len(v135) > 0 {
                                                                                                v135[0] = v1
                                                                                            }
                                                                                        case -874:
                                                                                            _ = len(v135)
                                                                                        }
                                                                                        _ = v160
                                                                                    } else {
                                                                                        v161 := GenS9{F0: v131, F1: 42}
                                                                                        v162 := GenS15{F0: 914}
                                                                                        var v163 float64 = 1e10
                                                                                        _ = v161
                                                                                        _ = v162
                                                                                        _ = v163
                                                                                    }
                                                                                    v164 := GenS9{F0: (^(v131)), F1: 0}
                                                                                    v2 = ((float32((3.14159 - 1e10)) / float32((float64((0.1 / v2)) - 1.5))) / v2)
                                                                                    var v165 float32 = (v2 + v2)
                                                                                    _ = v164
                                                                                    _ = v165
                                                                                } else {
                                                                                    var v166 string = ""
                                                                                    v167 := GenS8{F0: false}
                                                                                    _ = v166
                                                                                    _ = v167
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v1 = (v1 ^ 128)
                                                                        }
                                                                        fmt.Println(fi107, v136)
                                                                    }
                                                                    _ = v136
                                                                }
                                                                delete(v6, v1)
                                                                _ = v135
                                                            }
                                                        default:
                                                            switch 9223372036854775807 {
                                                            case 576:
                                                                fmt.Println(fi107)
                                                            case -1:
                                                                if false {
                                                                    _ = len(v3)
                                                                    var v168 int16 = (v5 % ((v5) | 1))
                                                                    _ = v168
                                                                }
                                                            case 128:
                                                                fmt.Println(v5, v2)
                                                            default:
                                                                var v169 float32 = (v2 * v2)
                                                                _ = v169
                                                            }
                                                        }
                                                        v5 = -693
                                                        _ = v132
                                                    }
                                                    _ = v131
                                                }
                                            default:
                                                v3[((("hello" + "\n\t") + "") + ("" + ((("hello" + "\n\t") + "hello") + "foo bar")))] = 1
                                            }
                                            if false {
                                                if ((false && (2 <= 743)) && true) {
                                                    var v170 bool = true
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v171 int64 = 42
                                                        _ = v171
                                                    case "":
                                                        v170 = v170
                                                    default:
                                                        var v172 bool = v170
                                                        _ = v172
                                                    }
                                                    v3[""] = 449
                                                    _ = v170
                                                } else {
                                                    var v173 uint8 = 40
                                                    v174 := []int{}
                                                    _ = v173
                                                    _ = v174
                                                }
                                                v175 := v6[(v1 ^ (v1 ^ (fi107 >> (((v4 >> ((42) & 63))) & 63))))]
                                                _ = v175
                                            }
                                            _ = v128
                                        }
                                    }
                                default:
                                    var v176 int64 = 0
                                    _ = v176
                                }
                            case 32768:
                                if true {
                                    switch (("foo bar" + ("foo bar" + "foo bar")) + "hello") {
                                    case "\n\t":
                                        v4 += (v4 % (((^(v4))) | 1))
                                    case "a":
                                        if true {
                                            switch 32768 {
                                            case -929:
                                                v177 := []int{128, v1, (fi107 >> ((255) & 63)), fi107}
                                                var v178 uint = v4
                                                var v179 float64 = 0.1
                                                _ = v177
                                                _ = v178
                                                _ = v179
                                            case 32767:
                                                v5 -= 42
                                                var v180 uint16 = 1
                                                switch "a" {
                                                case "":
                                                    if true {
                                                        fmt.Println(fi107, fi114)
                                                        v181 := GenS1{F0: 107, F1: (v180 | v180)}
                                                        fmt.Println(v4)
                                                        v5 = (v5 % ((1000) | 1))
                                                        if (!true) {
                                                            if (v4 <= (^(v4))) {
                                                                if true {
                                                                    v182 := GenS16{F0: (fi114 - 42), F1: 44}
                                                                    _ = v182
                                                                }
                                                                var v183 int16 = v5
                                                                _ = v183
                                                            }
                                                        }
                                                        v1 = fi107
                                                        v184 := GenS14{F0: v4, F1: 1, F2: v5}
                                                        v2 = (v2 * (v2 + float32(3.14159)))
                                                        var v185 uint = (v4 & v4)
                                                        delete(v6, (^(fi107)))
                                                        _ = v181
                                                        _ = v184
                                                        _ = v185
                                                    } else {
                                                        var v186 uint64 = 2
                                                        _ = v186
                                                    }
                                                case "hello":
                                                    if ("foo bar" < ("" + "")) {
                                                        switch (^(fi114)) {
                                                        case -2147483648:
                                                            var v187 uint = v4
                                                            _ = v187
                                                        case 128:
                                                            switch "" {
                                                            case "a":
                                                                v188 := GenS3{F0: -112, F1: 42, F2: 0}
                                                                _ = v188
                                                            }
                                                        }
                                                        v4 -= (^(v4))
                                                        if (!true) {
                                                            fmt.Println(v2, v4)
                                                            switch "foo bar" {
                                                            case "hello":
                                                                v5 = 1000
                                                            default:
                                                                if false {
                                                                    var v189 uint16 = v180
                                                                    var v190 string = "\n\t"
                                                                    _ = v189
                                                                    _ = v190
                                                                } else {
                                                                    v191 := GenS17{F0: false, F1: -2147483648}
                                                                    _ = v191
                                                                }
                                                            }
                                                        } else {
                                                            v5 = v5
                                                            if (uint64(v4) == 18446744073709551615) {
                                                                if false {
                                                                    v180 -= 236
                                                                    var v192 float32 = ((((1.5 * 3.14159) + (((v2 + (v2 - -2.25)) - (3.14159 - (1.5 - v2))) / v2)) * (-2.25 + v2)) + 1.5)
                                                                    _ = v192
                                                                }
                                                            }
                                                            var v193 string = "a"
                                                            _ = v193
                                                        }
                                                    }
                                                }
                                                v194 := GenS18{F0: ""}
                                                _ = v180
                                                _ = v194
                                            }
                                            var v195 float64 = 0.0
                                            _ = v195
                                        }
                                    case "foo bar":
                                        var v196 int8 = 1
                                        if (255 > 127) {
                                            switch fi114 {
                                            case -913:
                                                fmt.Println(v2, v5, v1)
                                            case 429:
                                                v1 += fi107
                                            default:
                                                v5 -= 42
                                            }
                                            switch fi107 {
                                            case 32768:
                                                fmt.Println(v5, v4, v4)
                                                if true {
                                                    switch ("hello" + "a") {
                                                    case "foo bar":
                                                        fmt.Println(v4, v5)
                                                    case "\n\t":
                                                        var v197 int16 = -32768
                                                        _ = v197
                                                    default:
                                                        v198 := GenS14{F0: v4, F1: -32768, F2: 32767}
                                                        _ = v198
                                                    }
                                                    var v199 uint = (v4 / ((528) | 1))
                                                    _ = len(v6)
                                                    _ = v199
                                                } else {
                                                    fmt.Println(fi114, v196)
                                                    fmt.Println(v1)
                                                    delete(v6, (fi114 >> (((^(v4))) & 63)))
                                                    if (0 < (v5 ^ (v5 - (v5 / (((v5 + (v5 & (v5 >> ((42) & 63))))) | 1))))) {
                                                        v200 := GenS18{F0: "\n\t"}
                                                        _ = v200
                                                    } else {
                                                        var v201 int32 = 207
                                                        v202 := []int{(fi114 >> ((18446744073709551615) & 63)), -1, (fi114 >> (((v4 ^ (v4 >> (((v4 % ((4294967295) | 1))) & 63)))) & 63)), fi107}
                                                        v2 = 1.5
                                                        _ = v201
                                                        _ = v202
                                                    }
                                                }
                                                v196 += (v196 / ((v196) | 1))
                                            case -154:
                                                v1 += (fi114 / ((fi114) | 1))
                                                v2 = (1e10 / 1.5)
                                                v2 += v2
                                            default:
                                                switch (-(fi114)) {
                                                case 9223372036854775807:
                                                    if (0 == 42) {
                                                        v203 := GenS13{F0: 2147483647, F1: 894, F2: 32}
                                                        if true {
                                                            var v204 int64 = -1
                                                            delete(v6, (-(v1)))
                                                            var v205 uint = (v4 & (v4 << ((v4) & 63)))
                                                            if (v4 != 42) {
                                                                v196 = (v196 >> ((v4) & 63))
                                                                v205 = (v4 + 1)
                                                                var v206 int16 = v5
                                                                if ((!((fi107 - 32768) == fi114)) && (!true)) {
                                                                    v4 = v205
                                                                    v2 = float32(0.1)
                                                                } else {
                                                                    v206 = -751
                                                                    if (v196 < (v196 ^ v196)) {
                                                                        if false {
                                                                            v4 = (^(v4))
                                                                            v207 := GenS4{F0: (fi114 % (((-(v1))) | 1)), F1: 42}
                                                                            _ = v207
                                                                        }
                                                                        var v208 uint64 = 854
                                                                        if (((("foo bar" + "") + "") + ("hello" + ("foo bar" + ""))) <= "hello") {
                                                                            var v209 uint8 = 2
                                                                            v210 := []int{fi114, (v1 & -341), fi107}
                                                                            v211 := GenS9{F0: (v196 >> ((v205) & 63)), F1: 42}
                                                                            _ = v209
                                                                            _ = v210
                                                                            _ = v211
                                                                        } else {
                                                                            fmt.Println(v206, v206, v204)
                                                                        }
                                                                        _ = v208
                                                                    } else {
                                                                        _ = len(v3)
                                                                        delete(v3, "foo bar")
                                                                    }
                                                                    var v212 int8 = (v196 << (((^(v4))) & 63))
                                                                    _ = v212
                                                                }
                                                                if false {
                                                                    var v213 uint8 = 42
                                                                    _ = v213
                                                                } else {
                                                                    switch ("foo bar" + "hello") {
                                                                    case "hello":
                                                                        var v214 uint64 = 4294967295
                                                                        _ = v214
                                                                    case "a":
                                                                        switch 1 {
                                                                        case 1:
                                                                            v215 := v3[(("a" + (("a" + "a") + "a")) + "a")]
                                                                            _ = v215
                                                                        case -1:
                                                                            if ((v196 < v196) || false) {
                                                                                fmt.Println(v2, fi114, fi114)
                                                                                var v216 uint64 = 0
                                                                                _ = v216
                                                                            } else {
                                                                                v4 = 1
                                                                                var v217 int64 = (v204 / ((v204) | 1))
                                                                                _ = v217
                                                                            }
                                                                        case 127:
                                                                            v1 = 696
                                                                        }
                                                                    case "":
                                                                        if true {
                                                                            var v218 string = ("\n\t" + ("hello" + "hello"))
                                                                            _ = v218
                                                                        }
                                                                    default:
                                                                        v219 := v6[(-(fi107))]
                                                                        _ = v219
                                                                    }
                                                                    v205 -= (v4 % (((v4 ^ v205)) | 1))
                                                                }
                                                                fmt.Println(v196)
                                                                _ = v206
                                                            } else {
                                                                var v220 int8 = v196
                                                                v204 = -9223372036854775808
                                                                fmt.Println(v1)
                                                                _ = v220
                                                            }
                                                            v221 := GenS13{F0: 42, F1: 4294967295, F2: 42}
                                                            _ = v204
                                                            _ = v205
                                                            _ = v221
                                                        } else {
                                                            var v222 float32 = (3.14159 + v2)
                                                            _ = v222
                                                        }
                                                        var v223 float32 = v2
                                                        _ = v203
                                                        _ = v223
                                                    }
                                                case 42:
                                                    var v224 uint32 = 4294967295
                                                    _ = v224
                                                }
                                            }
                                            v225 := GenS18{F0: "foo bar"}
                                            _ = v225
                                        }
                                        _ = v196
                                    default:
                                        v226 := []int{}
                                        _ = v226
                                    }
                                } else {
                                    v227 := []int{fi107, fi107, (fi114 + (fi114 << ((v4) & 63))), v1}
                                    v2 = v2
                                    v6[2147483647] = "hello"
                                    _ = v227
                                }
                            default:
                                v228 := GenS19{F0: "\n\t", F1: (v4 >> (((^(v4))) & 63))}
                                _ = v228
                            }
                        }
                    } else {
                        switch (fi114 ^ fi114) {
                        case 0:
                            v229 := GenS3{F0: 21, F1: 122, F2: 360}
                            _ = v229
                        case -83:
                            v230 := GenS8{F0: (v2 > (v2 - (v2 + ((float32(float64(v2)) - v2) / v2))))}
                            _ = v230
                        case 9223372036854775807:
                            v2 += 0.0
                        default:
                            if false {
                                if (false || true) {
                                    fmt.Println(v2, v1)
                                    v4 = (v4 << (((v4 / (((v4 & 65535)) | 1))) & 63))
                                } else {
                                    v2 = (0.1 - v2)
                                    v5 -= -32768
                                    var v231 uint32 = 1
                                    _ = v231
                                }
                                var v232 uint16 = 516
                                _ = v232
                            }
                        }
                        v233 := GenS20{F0: 128, F1: 0, F2: (^(v4))}
                        fmt.Println(fi107)
                        _ = v233
                    }
                    v2 -= v2
                    _ = fi114
                }
                _ = fi107
            }
            fmt.Println(v5, v1, v1)
        } else {
            v234 := GenS21{F0: 2, F1: 2147483647}
            _ = v234
        }
    }
    fmt.Println(v4, v1, v2)
    var v235 int8 = 1
    v2 = (1.5 * (v2 * (v2 - ((v2 / 1.5) * (float32(1e10) / v2)))))
    v236 := GenS4{F0: (v1 | v1), F1: -651}
    v235 += (v235 | (v235 | v235))
    for fi237 := 0; fi237 < 19; fi237++ {
        v4 -= v4
        _ = fi237
    }
    v238 := GenS20{F0: 255, F1: 66, F2: (v4 | (v4 & v4))}
    if ((((((("foo bar" + "") + ("\n\t" + ("a" + "a"))) + "\n\t") + "\n\t") + "\n\t") < "hello") && true) {
        v1 -= (v1 + v1)
        var v239 bool = (!true)
        for fi240 := 0; fi240 < 14; fi240++ {
            for fi241 := 0; fi241 < 14; fi241++ {
                var v242 uint = v4
                if ("" > "a") {
                    var v243 bool = true
                    v244 := []int{985, fi241}
                    _ = v243
                    _ = v244
                } else {
                    var v245 uint8 = 1
                    var v246 int8 = (v235 >> ((v242) & 63))
                    _ = v245
                    _ = v246
                }
                v247 := []int{2147483647, (fi240 * v1), (fi241 | 127), (fi241 * (^(v1)))}
                var v248 string = "a"
                fmt.Println(v4, v239, fi241)
                _ = fi241
                _ = v242
                _ = v247
                _ = v248
            }
            v2 += ((1e10 * (1.5 / (v2 * v2))) + v2)
            _ = fi240
        }
        _ = v239
    }
    fmt.Println(v2)
    v1 = 745
    var v249 int = (v1 % ((988) | 1))
    v4 = (v4 << (((^(v4))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v235:", v235)
    fmt.Println("v249:", v249)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v235
    _ = v236
    _ = v238
    _ = v249
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
