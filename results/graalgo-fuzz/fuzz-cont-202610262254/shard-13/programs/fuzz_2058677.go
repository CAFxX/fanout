package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 int32; F2 int64 }
type GenS8 struct { F0 uint16; F1 bool; F2 int16 }
type GenS9 struct { F0 uint; F1 int16; F2 int32 }
type GenS10 struct { F0 int16; F1 uint8 }
type GenS11 struct { F0 uint32; F1 int32 }
type GenS12 struct { F0 uint32; F1 uint32 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int16; F1 uint16 }
type GenS15 struct { F0 uint32; F1 int32; F2 int64 }
func genFn0() int {
    v6 := []int{2147483647, -214, 32768, 127}
    if (0.1 == -2.25) {
        if true {
            switch "hello" {
            case "":
                switch -464 {
                case 42:
                    v7 := GenS1{F0: 18446744073709551615}
                    _ = v7
                }
            default:
                var v8 int = 127
                _ = v8
            }
            v9 := GenS1{F0: 18446744073709551615}
            var v10 bool = true
            _ = v10
            _ = v9
        }
        var v11 int32 = 0
        _ = v11
    }
    var v12 uint64 = 244
    _ = v12
    _ = v6
    return -930
}
func genFn1(p0 uint64) int32 {
    var v13 uint16 = 255
    switch "foo bar" {
    case "foo bar":
        for fi14 := 0; fi14 < 10; fi14++ {
            switch "a" {
            case "\n\t":
                v15 := []int{(fi14 & (fi14 ^ fi14)), (fi14 << ((336) & 63)), 2147483647, (fi14 & (fi14 | (^(fi14))))}
                _ = v15
            case "hello":
                v16 := GenS1{F0: 42}
                _ = v16
            default:
                var v17 uint16 = 835
                _ = v17
            }
            v18 := []int{(fi14 + genFn0()), -1, genFn0(), (fi14 / ((fi14) | 1))}
            v19 := GenS1{F0: 4294967295}
            _ = fi14
            _ = v18
            _ = v19
        }
    default:
        v13 += (v13 ^ v13)
    }
    v20 := []int{}
    var v21 int8 = 1
    var v22 int16 = 42
    _ = p0
    _ = v13
    _ = v20
    _ = v21
    _ = v22
    return -852
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{4294967295}
    var v3 int32 = 865
    var v4 uint32 = 2147483647
    var v5 int = -357
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 int64 = 1
        if ((v4 & v4) >= v4) {
            v25 := GenS1{F0: 4294967295}
            _ = v25
        }
        var v26 uint16 = 32767
        _ = fi23
        _ = v24
        _ = v26
    }
    v27 := GenS2{F0: -128, F1: 1}
    var v28 uint64 = 2
    switch ("hello" + ("\n\t" + ("" + ""))) {
    case "\n\t":
        v5 -= (^(v5))
    case "a":
        fmt.Println(v5, v5, v5)
    case "hello":
        v4 = (v4 + (v4 << ((630) & 63)))
    default:
        var v29 uint8 = 127
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        if false {
            for fi31 := 0; fi31 < 2; fi31++ {
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
                _ = fi31
            }
        }
        for fi32 := 0; fi32 < 19; fi32++ {
            v4 -= 952
            v28 = 42
            if (!true) {
                v5 -= genFn0()
            } else {
                v4 = v4
                if (!(true || ((-(v5)) != (v5 & (fi30 ^ (v5 << ((0) & 63))))))) {
                    switch "" {
                    case "a":
                        var v33 int16 = 32767
                        _ = v33
                    }
                    v34 := GenS3{F0: v4, F1: false}
                    v35 := v1["foo bar"]
                    _ = v34
                    _ = v35
                } else {
                    if false {
                        v3 = (v3 + v3)
                        var v36 string = "\n\t"
                        _ = v36
                    }
                }
            }
            v2 = append(v2, (v4 * (v4 % (((v4 << ((2) & 63))) | 1))))
            _ = fi32
        }
        if true {
            var v37 int32 = (v3 - v3)
            var v38 int16 = 462
            if (!false) {
                v39 := GenS4{F0: (false || (127 >= 249)), F1: 9223372036854775807}
                v40 := GenS3{F0: 1, F1: ((^(v28)) >= 2)}
                _ = v39
                _ = v40
            }
            _ = v37
            _ = v38
        }
        _ = fi30
    }
    switch "" {
    case "":
        var v41 float64 = 0.0
        _ = v41
    }
    fmt.Println(v3, v5, v28)
    if true {
        var v42 bool = true
        v28 = v28
        _ = v42
    } else {
        for fi43 := 0; fi43 < 8; fi43++ {
            switch ("\n\t" + "hello") {
            case "\n\t":
                if ((-9223372036854775808 == -9223372036854775808) || true) {
                    switch ("\n\t" + "a") {
                    case "a":
                        switch -686 {
                        case 127:
                            delete(v1, "\n\t")
                        case 42:
                            fmt.Println(v5)
                        case -357:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    default:
                        for fi44 := 0; fi44 < 11; fi44++ {
                            v1["\n\t"] = (^(v4))
                            _ = fi44
                        }
                    }
                    if (!(1 < 2)) {
                        if true {
                            if (32767 != -359) {
                                switch (v5 - (fi43 + genFn0())) {
                                case 0:
                                    v5 -= (v5 ^ v5)
                                case 127:
                                    for fi45 := 0; fi45 < 13; fi45++ {
                                        if true {
                                            v4 = 24
                                        } else {
                                            var v46 int64 = 1
                                            _ = v46
                                        }
                                        v4 -= (v4 - (v4 / (((v4 * 776)) | 1)))
                                        _ = fi45
                                    }
                                case 32768:
                                    if true {
                                        v28 -= v28
                                    }
                                }
                                v47 := GenS5{F0: int16(v4), F1: int64(fi43)}
                                v48 := GenS5{F0: 1000, F1: 9223372036854775807}
                                _ = v47
                                _ = v48
                            }
                            var v49 int64 = -1
                            v5 = (v5 / ((2147483647) | 1))
                            _ = v49
                        }
                        if false {
                            switch (v5 + fi43) {
                            case -330:
                                for fi50 := 0; fi50 < 14; fi50++ {
                                    var v51 int16 = 32767
                                    _ = fi50
                                    _ = v51
                                }
                            case -464:
                                var v52 uint8 = 128
                                _ = v52
                            }
                        } else {
                            var v53 uint64 = 42
                            switch ("hello" + "hello") {
                            case "hello":
                                _ = len(v2)
                            case "\n\t":
                                fmt.Println(fi43, v5, v5)
                            case "":
                                for fi54 := 0; fi54 < 12; fi54++ {
                                    var v55 float64 = 3.14159
                                    var v56 float32 = (0.1 / -2.25)
                                    var v57 uint16 = 1
                                    switch "a" {
                                    case "\n\t":
                                        if ((v4 >= (v4 % (((v4 & (^(v4)))) | 1))) || (v55 > 3.14159)) {
                                            v53 = (v28 | 635)
                                            if false {
                                                var v58 string = "foo bar"
                                                v2 = append(v2, v4)
                                                if ((-(v3)) != (v3 + genFn1(v28))) {
                                                    if true {
                                                        v53 = v28
                                                        v59 := GenS6{F0: 0}
                                                        _ = v59
                                                    }
                                                } else {
                                                    fmt.Println(v28, v58)
                                                    v60 := GenS7{F0: 2, F1: v3, F2: 42}
                                                    _ = v60
                                                }
                                                _ = v58
                                            } else {
                                                v55 = 1e10
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v61 := GenS8{F0: 1, F1: true, F2: 1000}
                                                _ = v61
                                            }
                                            switch ("\n\t" + "hello") {
                                            case "a":
                                                fmt.Println(v55, v3, fi43)
                                            case "hello":
                                                var v62 int32 = genFn1((v53 % ((v28) | 1)))
                                                _ = v62
                                            case "\n\t":
                                                if false {
                                                    if (-2147483648 == (fi43 << ((1) & 63))) {
                                                        if (v56 <= (v56 * 1e10)) {
                                                            v4 = (v4 ^ 2)
                                                            var v63 int8 = 127
                                                            if len(v2) > 0 {
                                                                v2[0] = v4
                                                            }
                                                            switch (fi54 / (((fi43 | (fi54 >> ((548) & 63)))) | 1)) {
                                                            case 2147483647:
                                                                switch "a" {
                                                                case "a":
                                                                    var v64 uint32 = v4
                                                                    _ = v64
                                                                case "hello":
                                                                    switch (fi43 >> ((18446744073709551615) & 63)) {
                                                                    case -676:
                                                                        if len(v2) > 0 {
                                                                            v2[0] = v4
                                                                        }
                                                                    }
                                                                default:
                                                                    v1["foo bar"] = 0
                                                                }
                                                            case 42:
                                                                var v65 uint8 = 128
                                                                fmt.Println(v5)
                                                                _ = v65
                                                            case 609:
                                                                v28 = v28
                                                            default:
                                                                v66 := GenS9{F0: 85, F1: 540, F2: 0}
                                                                _ = v66
                                                            }
                                                            var v67 float32 = v56
                                                            _ = v63
                                                            _ = v67
                                                        } else {
                                                            v68 := GenS4{F0: (20 == 31), F1: -1}
                                                            v5 -= (v5 | fi54)
                                                            _ = v68
                                                        }
                                                        switch fi54 {
                                                        case -2147483648:
                                                            var v69 uint = 642
                                                            _ = v69
                                                        default:
                                                            v53 = v53
                                                            delete(v1, "hello")
                                                            v28 = (v53 & (^(v53)))
                                                        }
                                                        delete(v1, "foo bar")
                                                    } else {
                                                        _ = len(v2)
                                                        delete(v1, ("" + ("foo bar" + "hello")))
                                                        switch "" {
                                                        case "a":
                                                            v5 = (fi54 + fi54)
                                                        default:
                                                            switch (v5 ^ (fi54 % ((fi54) | 1))) {
                                                            case -424:
                                                                var v70 int16 = 1
                                                                _ = v70
                                                            case -637:
                                                                var v71 bool = (v56 != (-2.25 - v56))
                                                                _ = v71
                                                            case 341:
                                                                fmt.Println(fi54, v53, v4)
                                                            }
                                                        }
                                                        fmt.Println(v57, v28, v3)
                                                        var v72 int8 = 127
                                                        _ = v72
                                                    }
                                                    if (42 != 1) {
                                                        v73 := GenS10{F0: 32767, F1: 42}
                                                        switch fi43 {
                                                        case -636:
                                                            fmt.Println(v28)
                                                        case 128:
                                                            var v74 bool = true
                                                            _ = v74
                                                        default:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        _ = v73
                                                    }
                                                    v57 = (v57 | v57)
                                                } else {
                                                    fmt.Println(v3, v56, v53)
                                                    var v75 uint64 = (v28 & v28)
                                                    _ = v75
                                                }
                                            default:
                                                v57 = (v57 << ((42) & 63))
                                                var v76 int8 = -1
                                                _ = v76
                                            }
                                        }
                                    case "a":
                                        switch fi54 {
                                        case 1:
                                            v77 := v1[(((("\n\t" + "hello") + "\n\t") + "hello") + "a")]
                                            _ = v77
                                        case 42:
                                            var v78 uint16 = 42
                                            _ = v78
                                        default:
                                            v28 += (v53 >> ((767) & 63))
                                        }
                                        v3 = genFn1(v28)
                                        var v79 int64 = -9223372036854775808
                                        if true {
                                            fmt.Println(v55, v57)
                                            v53 = v53
                                        } else {
                                            var v80 int8 = -96
                                            _ = v80
                                        }
                                        _ = v79
                                    case "foo bar":
                                        v4 = 4294967295
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    fmt.Println(v56)
                                    _ = fi54
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                }
                            }
                            var v81 bool = false
                            _ = v53
                            _ = v81
                        }
                        v4 = v4
                    } else {
                        v28 += (v28 & (^(v28)))
                        if ((-2.25 / (((0.1 / float64(1e10)) * -2.25) - 3.14159)) == 1.5) {
                            fmt.Println(fi43)
                            if true {
                                switch (v5 & 4294967295) {
                                case -1:
                                    var v82 uint16 = 98
                                    _ = v82
                                case 32767:
                                    var v83 float64 = -2.25
                                    _ = v83
                                case -560:
                                    var v84 uint32 = v4
                                    _ = v84
                                }
                            } else {
                                v85 := GenS11{F0: 4294967295, F1: (v3 ^ genFn1((v28 * (v28 % ((18446744073709551615) | 1)))))}
                                _ = v85
                            }
                        } else {
                            for fi86 := 0; fi86 < 1; fi86++ {
                                delete(v1, ("\n\t" + "a"))
                                var v87 uint64 = (v28 / ((18446744073709551615) | 1))
                                _ = fi86
                                _ = v87
                            }
                        }
                        switch "hello" {
                        case "foo bar":
                            switch fi43 {
                            case -9223372036854775808:
                                v88 := GenS3{F0: 2147483648, F1: true}
                                _ = v88
                            case 497:
                                switch "" {
                                case "":
                                    v89 := GenS4{F0: (100 != -38), F1: 1}
                                    _ = v89
                                case "hello":
                                    var v90 bool = true
                                    _ = v90
                                case "\n\t":
                                    if (!true) {
                                        if (!(-1 <= -128)) {
                                            v91 := GenS4{F0: false, F1: 42}
                                            _ = v91
                                        }
                                        v3 -= 464
                                        v92 := GenS12{F0: v4, F1: v4}
                                        _ = v92
                                    }
                                default:
                                    var v93 int8 = -128
                                    _ = v93
                                }
                            case -762:
                                v28 -= 631
                            }
                        case "a":
                            for fi94 := 0; fi94 < 13; fi94++ {
                                if false {
                                    var v95 int64 = -653
                                    _ = v95
                                }
                                if len(v2) > 0 {
                                    v2[0] = v4
                                }
                                _ = fi94
                            }
                        default:
                            var v96 uint8 = 255
                            _ = v96
                        }
                    }
                    v97 := v1["a"]
                    _ = v97
                }
            case "foo bar":
                for fi98 := 0; fi98 < 2; fi98++ {
                    v4 -= uint32(v28)
                    v3 += (v3 * 1)
                    if false {
                        var v99 bool = false
                        _ = v99
                    } else {
                        var v100 int32 = (v3 >> ((65535) & 63))
                        if ((true && (true || true)) || true) {
                            switch "" {
                            case "\n\t":
                                delete(v1, "foo bar")
                            case "":
                                var v101 uint16 = 42
                                fmt.Println(v3)
                                _ = len(v2)
                                _ = v101
                            case "hello":
                                switch "a" {
                                case "\n\t":
                                    if true {
                                        delete(v1, "\n\t")
                                    } else {
                                        v100 -= (v3 | -2147483648)
                                        if true {
                                            v3 -= v3
                                            if (-1 >= 0) {
                                                v102 := GenS5{F0: 32767, F1: 1}
                                                _ = len(v1)
                                                fmt.Println(fi98)
                                                _ = v102
                                            } else {
                                                var v103 int16 = 42
                                                var v104 uint32 = (v4 / (((^(v4))) | 1))
                                                v100 -= genFn1((v28 >> ((255) & 63)))
                                                _ = v103
                                                _ = v104
                                            }
                                        }
                                    }
                                    v28 += (^(v28))
                                case "a":
                                    v3 -= genFn1(v28)
                                default:
                                    fmt.Println(v100)
                                }
                            default:
                                switch "foo bar" {
                                case "":
                                    var v105 int8 = 48
                                    _ = v105
                                case "hello":
                                    fmt.Println(v4, v3, v100)
                                case "\n\t":
                                    var v106 uint8 = 42
                                    _ = v106
                                }
                            }
                            v5 += fi98
                        }
                        _ = v100
                    }
                    _ = fi98
                }
            default:
                switch "a" {
                case "a":
                    var v107 int = fi43
                    _ = v107
                case "hello":
                    if true {
                        if (("hello" + "\n\t") <= ("foo bar" + "hello")) {
                            fmt.Println(v3, v5)
                            if (!(55 < 1)) {
                                var v108 int = (v5 << ((1) & 63))
                                _ = v108
                            } else {
                                v109 := GenS6{F0: 1000}
                                for fi110 := 0; fi110 < 15; fi110++ {
                                    if (false && (!(!true))) {
                                        switch fi43 {
                                        case -9223372036854775808:
                                            var v111 uint16 = 32767
                                            _ = v111
                                        case 32768:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        default:
                                            var v112 uint64 = v28
                                            _ = v112
                                        }
                                        switch (fi110 >> ((18446744073709551615) & 63)) {
                                        case -2147483648:
                                            fmt.Println(v28, v3, v3)
                                        default:
                                            v113 := GenS2{F0: 0, F1: 9223372036854775807}
                                            _ = v113
                                        }
                                        fmt.Println(fi110, v5, fi43)
                                    }
                                    v3 = (v3 / ((genFn1(v28)) | 1))
                                    if true {
                                        v5 = (fi43 & v5)
                                    } else {
                                        v114 := GenS12{F0: (v4 % ((v4) | 1)), F1: (v4 << ((533) & 63))}
                                        _ = v114
                                    }
                                    _ = fi110
                                }
                                for fi115 := 0; fi115 < 9; fi115++ {
                                    v116 := GenS4{F0: true, F1: 42}
                                    _ = fi115
                                    _ = v116
                                }
                                _ = v109
                            }
                            var v117 uint8 = 0
                            _ = v117
                        } else {
                            v118 := GenS12{F0: v4, F1: 42}
                            if false {
                                var v119 uint16 = 42
                                for fi120 := 0; fi120 < 17; fi120++ {
                                    var v121 uint16 = 185
                                    var v122 uint32 = v4
                                    _ = fi120
                                    _ = v121
                                    _ = v122
                                }
                                if ((v4 % ((v4) | 1)) < (^(v4))) {
                                    switch 1 {
                                    case 4294967295:
                                        v2 = append(v2, 1)
                                    case 0:
                                        if true {
                                            for fi123 := 0; fi123 < 15; fi123++ {
                                                v124 := GenS13{F0: 42}
                                                _ = fi123
                                                _ = v124
                                            }
                                        } else {
                                            for fi125 := 0; fi125 < 7; fi125++ {
                                                if true {
                                                    if len(v2) > 0 {
                                                        v2[0] = (v4 / (((v4 & v4)) | 1))
                                                    }
                                                    delete(v1, "")
                                                }
                                                if (true || (true && (true && (v119 == 0)))) {
                                                    var v126 int = (-(fi125))
                                                    if false {
                                                        v5 -= (v126 % ((genFn0()) | 1))
                                                        if false {
                                                            var v127 int32 = 1
                                                            v128 := v1[""]
                                                            var v129 uint16 = (v119 & v119)
                                                            _ = v127
                                                            _ = v128
                                                            _ = v129
                                                        } else {
                                                            if (true || ((float32((0.1 * 0.1)) / 1.0) >= float32(-2.25))) {
                                                                var v130 string = ("foo bar" + ("hello" + ""))
                                                                if ((9223372036854775807 < -9223372036854775808) && false) {
                                                                    v119 = v119
                                                                    if ((^(v4)) == (^(v4))) {
                                                                        v131 := GenS1{F0: 65535}
                                                                        fmt.Println(v3, fi125)
                                                                        v132 := GenS7{F0: (v28 ^ (v28 | (v28 & (v28 & (v28 << ((255) & 63)))))), F1: 2147483647, F2: 9223372036854775807}
                                                                        _ = v131
                                                                        _ = v132
                                                                    }
                                                                    var v133 int8 = 77
                                                                    _ = v133
                                                                }
                                                                _ = v130
                                                            }
                                                            var v134 int8 = -95
                                                            _ = v134
                                                        }
                                                        fmt.Println(v119, v5)
                                                    } else {
                                                        switch (-(fi125)) {
                                                        case 32768:
                                                            var v135 int = int(v3)
                                                            _ = v135
                                                        case -9223372036854775808:
                                                            v119 = (v119 % ((v119) | 1))
                                                        default:
                                                            var v136 uint16 = 255
                                                            _ = v136
                                                        }
                                                        if (1 >= 1) {
                                                            var v137 uint32 = v4
                                                            v138 := GenS14{F0: 328, F1: (v119 % (((v119 | v119)) | 1))}
                                                            v137 += v4
                                                            _ = v137
                                                            _ = v138
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    _ = v126
                                                } else {
                                                    switch genFn0() {
                                                    case 128:
                                                        fmt.Println(v119)
                                                        v3 -= (v3 & (v3 << ((93) & 63)))
                                                        fmt.Println(v4)
                                                        _ = len(v2)
                                                        v139 := GenS8{F0: 65535, F1: true, F2: -32768}
                                                        switch ("a" + "") {
                                                        case "":
                                                            if (507 != 9223372036854775807) {
                                                                v140 := GenS4{F0: (v3 < (v3 ^ -12)), F1: 42}
                                                                var v141 int8 = 127
                                                                v3 = 98
                                                                _ = v140
                                                                _ = v141
                                                            } else {
                                                                fmt.Println(fi43)
                                                                if (0 != -1) {
                                                                    var v142 uint32 = 137
                                                                    _ = v142
                                                                } else {
                                                                    if false {
                                                                        if (("" != "\n\t") && true) {
                                                                            v119 += v119
                                                                        } else {
                                                                            v2 = append(v2, v4)
                                                                            v143 := GenS4{F0: ((0.1 - 3.14159) >= ((0.0 * ((1e10 - (float64(0.0) * 3.14159)) + 0.1)) + 0.1)), F1: 1}
                                                                            var v144 int16 = -1
                                                                            _ = v143
                                                                            _ = v144
                                                                        }
                                                                        var v145 string = "hello"
                                                                        var v146 float64 = (1.5 - float64((-2.25 - (0.0 * -2.25))))
                                                                        _ = v145
                                                                        _ = v146
                                                                    }
                                                                }
                                                            }
                                                        case "\n\t":
                                                            _ = len(v2)
                                                        case "foo bar":
                                                            var v147 uint32 = 474
                                                            _ = v147
                                                        }
                                                        var v148 int32 = v3
                                                        v28 += 18446744073709551615
                                                        v148 = (v148 >> ((18446744073709551615) & 63))
                                                        _ = v139
                                                        _ = v148
                                                    default:
                                                        fmt.Println(v119)
                                                    }
                                                }
                                                var v149 uint16 = 42
                                                switch "\n\t" {
                                                case "a":
                                                    v150 := GenS8{F0: v119, F1: (v4 >= 65535), F2: 32767}
                                                    _ = v150
                                                case "foo bar":
                                                    v4 = (v4 * (v4 << ((uint(v4)) & 63)))
                                                default:
                                                    var v151 int = fi43
                                                    _ = v151
                                                }
                                                switch (fi43 * (fi125 * 9223372036854775807)) {
                                                case 128:
                                                    var v152 int16 = 152
                                                    _ = v152
                                                default:
                                                    switch v5 {
                                                    case 1:
                                                        if (!false) {
                                                            var v153 bool = (70 < 2)
                                                            v154 := GenS15{F0: (v4 % ((65535) | 1)), F1: v3, F2: 9223372036854775807}
                                                            _ = v153
                                                            _ = v154
                                                        }
                                                    case 127:
                                                        var v155 uint8 = 0
                                                        _ = v155
                                                    case -2147483648:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    default:
                                                        var v156 string = "a"
                                                        _ = v156
                                                    }
                                                }
                                                v157 := GenS14{F0: -1, F1: (v119 % ((v119) | 1))}
                                                v119 = v119
                                                _ = fi125
                                                _ = v149
                                                _ = v157
                                            }
                                        }
                                    default:
                                        v158 := GenS10{F0: -1, F1: 1}
                                        _ = v158
                                    }
                                }
                                _ = v119
                            } else {
                                v28 += v28
                            }
                            _ = v118
                        }
                        if false {
                            var v159 uint = 1
                            _ = v159
                        }
                    } else {
                        if ((!(18 < 184)) && false) {
                            v2 = append(v2, v4)
                        }
                        for fi160 := 0; fi160 < 17; fi160++ {
                            if true {
                                v5 = v5
                                var v161 int64 = 0
                                fmt.Println(v28, fi43)
                                _ = v161
                            }
                            v162 := GenS8{F0: 42, F1: true, F2: -100}
                            fmt.Println(v4)
                            _ = fi160
                            _ = v162
                        }
                        var v163 int8 = -1
                        _ = v163
                    }
                case "\n\t":
                    var v164 bool = false
                    _ = v164
                }
            }
            switch genFn0() {
            case 127:
                var v165 string = ""
                _ = v165
            case 9223372036854775807:
                fmt.Println(v4, v3, fi43)
            default:
                v166 := GenS2{F0: 83, F1: 42}
                _ = v166
            }
            _ = fi43
        }
        if (false || true) {
            _ = len(v2)
            v28 -= 1
            v28 = v28
        }
    }
    fmt.Println(v4)
    if len(v2) > 0 {
        v2[0] = uint32(v28)
    }
    if (false && true) {
        v3 += (v3 ^ 42)
        v3 = (v3 | genFn1(v28))
        v28 -= v28
    }
    v167 := GenS6{F0: int16(v4)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v167
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
