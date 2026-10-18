package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8; F1 uint64 }
type GenS11 struct { F0 bool; F1 uint64; F2 int32 }
type GenS12 struct { F0 int; F1 int8; F2 int }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 string; F1 int32; F2 int8 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS17 struct { F0 uint64; F1 bool }
type GenS18 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS19 struct { F0 int8; F1 int64 }
type GenS20 struct { F0 uint64; F1 int16; F2 int }
func genFn0(p0 bool) bool {
    var v6 float64 = 1.5
    p0 = p0
    var v7 uint = 255
    _ = p0
    _ = v6
    _ = v7
    return (!(-806 != 922))
}

func main() {
    var v1 uint = 1
    var v2 uint64 = 18446744073709551615
    var v3 uint32 = 4294967295
    var v4 int8 = 1
    var v5 int32 = -2147483648
    v1 -= (v1 + (^(v1)))
    var v8 uint16 = 32767
    if true {
        var v9 int = 32767
        switch 111 {
        case -1:
            fmt.Println(v2)
        case 0:
            switch ("foo bar" + "\n\t") {
            case "a":
                v9 = 9223372036854775807
            case "\n\t":
                var v10 float32 = (0.1 - (0.1 * 1.5))
                _ = v10
            }
        }
        v11 := GenS1{F0: -2147483648}
        _ = v11
        _ = v9
    }
    if (((v3 <= (v3 & 2)) && ((v5 | v5) <= (v5 % ((v5) | 1)))) && (false || ((0.0 >= 1e10) && (true && ((v2 / (((v2 + v2)) | 1)) <= v2))))) {
        v12 := GenS1{F0: 31}
        var v13 uint32 = 64
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := GenS1{F0: (v5 << ((v1) & 63))}
            v16 := GenS1{F0: 2147483647}
            fmt.Println(v8)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v12
        _ = v13
    } else {
        if true {
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := GenS1{F0: v5}
                v19 := []int{32768, fi17, fi17}
                switch "foo bar" {
                case "hello":
                    var v20 bool = (false || false)
                    _ = v20
                case "\n\t":
                    if (!true) {
                        v4 = v4
                        var v21 bool = true
                        _ = v21
                    } else {
                        v8 = (v8 % ((v8) | 1))
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        for fi22 := 0; fi22 < 6; fi22++ {
                            var v23 int32 = int32(fi22)
                            _ = fi22
                            _ = v23
                        }
                    }
                }
                _ = fi17
                _ = v18
                _ = v19
            }
        } else {
            fmt.Println(v4, v5)
        }
        var v24 uint64 = 2
        _ = v24
    }
    v5 = (v5 % ((v5) | 1))
    v1 = (v1 * (v1 >> ((v1) & 63)))
    if true {
        if true {
            var v25 int32 = (v5 | 1)
            v26 := GenS1{F0: (-(v5))}
            v1 = (^(v1))
            _ = v25
            _ = v26
        } else {
            v27 := []int{650, 127}
            v28 := GenS1{F0: -1}
            v2 += (v2 / ((v2) | 1))
            _ = v27
            _ = v28
        }
        switch ("\n\t" + "a") {
        case "foo bar":
            switch ((("foo bar" + "a") + ("hello" + "")) + "a") {
            case "":
                v29 := []int{int(v4)}
                _ = v29
            default:
                if (1 < -519) {
                    var v30 uint64 = (^(v2))
                    if (!true) {
                        if true {
                            for fi31 := 0; fi31 < 18; fi31++ {
                                switch "a" {
                                case "a":
                                    var v32 uint8 = uint8(v8)
                                    _ = v32
                                case "foo bar":
                                    var v33 float32 = 1e10
                                    _ = v33
                                }
                                var v34 uint32 = v3
                                _ = fi31
                                _ = v34
                            }
                            var v35 int16 = 913
                            var v36 int16 = v35
                            _ = v35
                            _ = v36
                        } else {
                            fmt.Println(v1)
                            v37 := GenS1{F0: (v5 << ((v1) & 63))}
                            v8 = v8
                            _ = v37
                        }
                    }
                    switch "a" {
                    case "":
                        for fi38 := 0; fi38 < 6; fi38++ {
                            for fi39 := 0; fi39 < 19; fi39++ {
                                fmt.Println(fi39, v3)
                                if true {
                                    v40 := GenS1{F0: v5}
                                    switch (fi38 | 554) {
                                    case -866:
                                        v4 = v4
                                    case -359:
                                        v8 = v8
                                    case -847:
                                        fmt.Println(fi39, v2, v2)
                                    default:
                                        var v41 string = "foo bar"
                                        _ = v41
                                    }
                                    _ = v40
                                }
                                _ = fi39
                            }
                            v42 := GenS2{F0: (v4 / ((v4) | 1))}
                            switch "hello" {
                            case "hello":
                                switch 0 {
                                case 127:
                                    v5 -= v5
                                }
                            case "a":
                                var v43 uint64 = (v30 << ((uint(v2)) & 63))
                                _ = v43
                            case "foo bar":
                                if (fi38 > fi38) {
                                    if (!false) {
                                        switch (fi38 >> ((v1) & 63)) {
                                        case 2147483647:
                                            if ((42 < 1) || ((-2.25 + 3.14159) > (-2.25 / 1.5))) {
                                                var v44 uint64 = v2
                                                v45 := []int{(fi38 % ((fi38) | 1)), (^(fi38))}
                                                _ = v44
                                                _ = v45
                                            } else {
                                                v46 := GenS3{F0: 0, F1: v3}
                                                _ = v46
                                            }
                                        case 0:
                                            v1 -= v1
                                        case 9223372036854775807:
                                            switch ("hello" + "") {
                                            case "foo bar":
                                                switch (fi38 - fi38) {
                                                case -1:
                                                    v47 := GenS4{F0: v5, F1: v4}
                                                    _ = v47
                                                }
                                            default:
                                                if false {
                                                    v48 := GenS5{F0: (^(v2))}
                                                    v49 := []int{fi38}
                                                    switch ("" + ((("\n\t" + ("\n\t" + "foo bar")) + "") + "")) {
                                                    case "\n\t":
                                                        _ = len(v49)
                                                    case "foo bar":
                                                        var v50 uint16 = (v8 % ((65535) | 1))
                                                        _ = v50
                                                    default:
                                                        var v51 int = (fi38 ^ (fi38 ^ -936))
                                                        _ = v51
                                                    }
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    fmt.Println(v5)
                                                    var v52 int16 = 1
                                                    _ = v52
                                                }
                                            }
                                        }
                                    } else {
                                        switch ("" + ("" + "a")) {
                                        case "\n\t":
                                            var v53 uint16 = v8
                                            _ = v53
                                        case "":
                                            for fi54 := 0; fi54 < 11; fi54++ {
                                                var v55 int8 = v4
                                                _ = fi54
                                                _ = v55
                                            }
                                        }
                                    }
                                    v56 := []int{fi38, fi38}
                                    v30 -= (v2 + (^(v30)))
                                    _ = v56
                                }
                            default:
                                var v57 string = (("a" + "a") + ("\n\t" + "a"))
                                _ = v57
                            }
                            _ = fi38
                            _ = v42
                        }
                    default:
                        v58 := []int{32767, 110, 32768}
                        _ = v58
                    }
                    _ = v30
                }
            }
        case "":
            var v59 int8 = v4
            _ = v59
        }
        switch 144 {
        case 127:
            v5 = (v5 << (((v1 + (v1 | 594))) & 63))
        case 42:
            v60 := []int{557}
            _ = v60
        case 4294967295:
            switch 4294967295 {
            case 32768:
                for fi61 := 0; fi61 < 16; fi61++ {
                    if false {
                        v5 = v5
                    } else {
                        v2 += 42
                        v62 := []int{(fi61 + (fi61 / ((fi61) | 1))), 1, -910, 76}
                        _ = v62
                    }
                    _ = fi61
                }
            case 9223372036854775807:
                v1 = 404
            case 128:
                if ((32767 == -287) || (v2 >= v2)) {
                    v4 += v4
                    v4 -= 127
                    for fi63 := 0; fi63 < 4; fi63++ {
                        var v64 int16 = -360
                        _ = fi63
                        _ = v64
                    }
                } else {
                    switch "a" {
                    case "foo bar":
                        v65 := GenS2{F0: v4}
                        _ = v65
                    }
                    v66 := []int{0, 2147483647}
                    v3 -= v3
                    _ = v66
                }
            }
        }
    } else {
        for fi67 := 0; fi67 < 6; fi67++ {
            var v68 int8 = 127
            _ = fi67
            _ = v68
        }
        if (((9223372036854775807 == -9223372036854775808) && false) || false) {
            var v69 uint = v1
            var v70 uint64 = (^(v2))
            _ = v69
            _ = v70
        } else {
            for fi71 := 0; fi71 < 12; fi71++ {
                for fi72 := 0; fi72 < 15; fi72++ {
                    v73 := []int{(fi71 & (fi71 >> (((^(v1))) & 63))), fi72, fi72}
                    _ = fi72
                    _ = v73
                }
                var v74 float32 = float32(1.5)
                if (((-486 < 0) || false) || true) {
                    v3 -= 1
                } else {
                    if ((false && false) || (!(-9223372036854775808 <= (fi71 & 4294967295)))) {
                        if ((fi71 << (((v1 | (v1 / ((v1) | 1)))) & 63)) > (fi71 * fi71)) {
                            if true {
                                var v75 bool = ((v4 % ((v4) | 1)) != (-(v4)))
                                var v76 bool = (128 < 127)
                                switch "a" {
                                case "a":
                                    var v77 int64 = 42
                                    _ = v77
                                case "\n\t":
                                    switch -1 {
                                    case 9223372036854775807:
                                        switch "\n\t" {
                                        case "a":
                                            if false {
                                                var v78 bool = false
                                                if v75 {
                                                    switch (fi71 + fi71) {
                                                    case 32767:
                                                        switch fi71 {
                                                        case 0:
                                                            var v79 float64 = (1e10 / (float64((v74 / (v74 + (v74 * (1e10 * float32(((3.14159 - 0.1) - 0.1))))))) * (1e10 - 1e10)))
                                                            _ = v79
                                                        case -544:
                                                            v80 := []int{fi71, (-(fi71))}
                                                            _ = v80
                                                        case 9223372036854775807:
                                                            switch (fi71 + fi71) {
                                                            case 0:
                                                                v5 += -2147483648
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(v5)
                                                    }
                                                    v2 = v2
                                                    v4 = (v4 * v4)
                                                }
                                                v81 := []int{(fi71 << (((v1 >> (((v1 ^ uint(v2))) & 63))) & 63)), (fi71 | fi71)}
                                                _ = v78
                                                _ = v81
                                            } else {
                                                var v82 string = "\n\t"
                                                var v83 int64 = 42
                                                v84 := []int{}
                                                _ = v82
                                                _ = v83
                                                _ = v84
                                            }
                                        case "foo bar":
                                            v76 = v75
                                        case "\n\t":
                                            if (-51 != 2147483647) {
                                                for fi85 := 0; fi85 < 12; fi85++ {
                                                    var v86 int64 = -1
                                                    switch (("\n\t" + "a") + "\n\t") {
                                                    case "\n\t":
                                                        var v87 uint64 = 4294967295
                                                        _ = v87
                                                    }
                                                    _ = fi85
                                                    _ = v86
                                                }
                                            }
                                        }
                                    case -413:
                                        var v88 uint = (v1 ^ (^(v1)))
                                        _ = v88
                                    default:
                                        fmt.Println(v5, fi71)
                                    }
                                case "foo bar":
                                    v89 := []int{(fi71 + 2147483647), fi71, 2147483647}
                                    _ = v89
                                default:
                                    var v90 uint32 = (v3 ^ (v3 / (((^(v3))) | 1)))
                                    _ = v90
                                }
                                _ = v75
                                _ = v76
                            }
                            for fi91 := 0; fi91 < 7; fi91++ {
                                v4 = (v4 * -4)
                                _ = fi91
                            }
                            v74 -= ((float32(-2.25) + v74) - v74)
                        } else {
                            for fi92 := 0; fi92 < 2; fi92++ {
                                v93 := GenS3{F0: v8, F1: v3}
                                _ = fi92
                                _ = v93
                            }
                            for fi94 := 0; fi94 < 14; fi94++ {
                                var v95 uint64 = 1
                                if false {
                                    v1 -= v1
                                }
                                var v96 uint64 = 42
                                _ = fi94
                                _ = v95
                                _ = v96
                            }
                        }
                        var v97 uint64 = v2
                        _ = v97
                    } else {
                        if true {
                            for fi98 := 0; fi98 < 13; fi98++ {
                                v99 := GenS6{F0: 4294967295, F1: v8}
                                if true {
                                    var v100 uint64 = (v2 | v2)
                                    _ = v100
                                }
                                v5 += (v5 - (-(v5)))
                                _ = fi98
                                _ = v99
                            }
                            switch (fi71 | (fi71 - (fi71 + 16))) {
                            case 229:
                                v101 := []int{2147483647, fi71, (fi71 ^ -9223372036854775808)}
                                _ = v101
                            case 1:
                                v8 += v8
                            case 0:
                                v102 := []int{32768, fi71, fi71, (fi71 << ((v1) & 63))}
                                _ = v102
                            default:
                                v103 := []int{(-(fi71)), fi71}
                                _ = v103
                            }
                        } else {
                            var v104 int = fi71
                            switch 4294967295 {
                            case 128:
                                switch "" {
                                case "foo bar":
                                    if (v4 == (v4 >> (((v1 | v1)) & 63))) {
                                        if (true || true) {
                                            v74 += 3.14159
                                            v4 = 1
                                            v4 = 42
                                        } else {
                                            for fi105 := 0; fi105 < 12; fi105++ {
                                                var v106 int32 = v5
                                                _ = fi105
                                                _ = v106
                                            }
                                            for fi107 := 0; fi107 < 9; fi107++ {
                                                fmt.Println(v8)
                                                v3 = 2
                                                v108 := []int{}
                                                _ = fi107
                                                _ = v108
                                            }
                                        }
                                        for fi109 := 0; fi109 < 8; fi109++ {
                                            v5 = (v5 * 297)
                                            switch (v104 % (((fi71 * int(v5))) | 1)) {
                                            case 490:
                                                v110 := GenS6{F0: -1, F1: v8}
                                                var v111 int = v104
                                                v2 = v2
                                                if false {
                                                    var v112 float32 = v74
                                                    _ = v112
                                                }
                                                fmt.Println(v111, v8)
                                                fmt.Println(v2, v3)
                                                _ = v110
                                                _ = v111
                                            case 32768:
                                                var v113 int32 = (v5 % (((v5 << ((v1) & 63))) | 1))
                                                _ = v113
                                            default:
                                                v114 := GenS6{F0: 42, F1: v8}
                                                _ = v114
                                            }
                                            fmt.Println(v104, v2)
                                            _ = fi109
                                        }
                                        for fi115 := 0; fi115 < 16; fi115++ {
                                            switch "" {
                                            case "\n\t":
                                                switch ("" + (("" + "foo bar") + ("" + (("\n\t" + "hello") + "hello")))) {
                                                case "foo bar":
                                                    switch fi71 {
                                                    case 42:
                                                        var v116 uint = 783
                                                        _ = v116
                                                    case 128:
                                                        v104 = (v104 ^ fi71)
                                                    case 127:
                                                        v117 := GenS3{F0: v8, F1: (v3 + (v3 << (((v1 + uint(v3))) & 63)))}
                                                        _ = v117
                                                    default:
                                                        v8 += (v8 << ((255) & 63))
                                                    }
                                                }
                                            case "":
                                                if ((true && (1 < v4)) || (-9223372036854775808 > -1)) {
                                                    fmt.Println(v1, v74, v8)
                                                    switch fi115 {
                                                    case 9223372036854775807:
                                                        if false {
                                                            v118 := GenS7{F0: ("hello" + ("hello" + "a"))}
                                                            if true {
                                                                if true {
                                                                    v119 := []int{(v104 / ((fi71) | 1)), fi71, -113}
                                                                    _ = v119
                                                                }
                                                                v2 = v2
                                                            }
                                                            v3 = (v3 ^ 2147483648)
                                                            _ = v118
                                                        } else {
                                                            switch (fi71 + v104) {
                                                            case 32768:
                                                                v120 := GenS4{F0: 2147483647, F1: v4}
                                                                _ = v120
                                                            case 0:
                                                                fmt.Println(v1, v1, v4)
                                                            case 1:
                                                                if (-1 < -9223372036854775808) {
                                                                    v121 := GenS8{F0: false, F1: (-(v5))}
                                                                    fmt.Println(fi115)
                                                                    fmt.Println(fi115, fi115)
                                                                    _ = v121
                                                                }
                                                            }
                                                            if (true && false) {
                                                                if true {
                                                                    v2 -= v2
                                                                    var v122 int32 = 1
                                                                    fmt.Println(v4, fi115, v4)
                                                                    _ = v122
                                                                } else {
                                                                    fmt.Println(v74)
                                                                    v123 := GenS8{F0: ((-9223372036854775808 != (-(v104))) || (!false)), F1: int32(v2)}
                                                                    _ = v123
                                                                }
                                                            }
                                                            switch v104 {
                                                            case 32767:
                                                                v124 := []int{}
                                                                _ = v124
                                                            }
                                                            var v125 int64 = 0
                                                            fmt.Println(v125, v8)
                                                            if false {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v3 -= (^(v3))
                                                                    v8 += 42
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        if false {
                                                                            var v126 bool = ("\n\t" > ((("" + ("\n\t" + "a")) + "hello") + "hello"))
                                                                            _ = v126
                                                                        } else {
                                                                            v127 := []int{int(v125)}
                                                                            if len(v127) > 0 {
                                                                                v127[0] = fi71
                                                                            }
                                                                            var v128 string = "foo bar"
                                                                            _ = v127
                                                                            _ = v128
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi71, v104, v74)
                                                                    }
                                                                    v129 := GenS7{F0: ("a" + "hello")}
                                                                    _ = v129
                                                                default:
                                                                    var v130 uint = 0
                                                                    _ = v130
                                                                }
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    fmt.Println(v74)
                                                                case "a":
                                                                    if (((!false) && true) || ((v1 + (v1 % ((42) | 1))) <= (v1 >> ((v1) & 63)))) {
                                                                        switch (^(fi115)) {
                                                                        case -345:
                                                                            if (1 == int32(v2)) {
                                                                                fmt.Println(fi115, v4, v3)
                                                                            } else {
                                                                                var v131 int = (v104 + (v104 % ((fi71) | 1)))
                                                                                var v132 uint64 = v2
                                                                                v1 = uint(v132)
                                                                                _ = v131
                                                                                _ = v132
                                                                            }
                                                                        case 1:
                                                                            var v133 int64 = (-(v125))
                                                                            _ = v133
                                                                        default:
                                                                            if (true || (-992 >= -32768)) {
                                                                                fmt.Println(v2, fi71, v104)
                                                                                if (v3 < (^(v3))) {
                                                                                    if ((v74 > 0.1) || (!((v3 << ((v1) & 63)) == v3))) {
                                                                                        v104 += int(v5)
                                                                                        if ((!(v8 == uint16(v3))) && (false && false)) {
                                                                                            if false {
                                                                                                v134 := GenS9{F0: (^(v2))}
                                                                                                var v135 uint8 = uint8(v3)
                                                                                                _ = v134
                                                                                                _ = v135
                                                                                            }
                                                                                            fmt.Println(v8, v8)
                                                                                        }
                                                                                        if true {
                                                                                            fmt.Println(v104)
                                                                                            v104 = int(v5)
                                                                                        }
                                                                                        if true {
                                                                                            var v136 string = (("foo bar" + "\n\t") + "foo bar")
                                                                                            var v137 string = v136
                                                                                            _ = v136
                                                                                            _ = v137
                                                                                        } else {
                                                                                            switch fi115 {
                                                                                            case 1:
                                                                                                v138 := GenS1{F0: v5}
                                                                                                _ = v138
                                                                                            case 42:
                                                                                                v139 := []int{(fi115 & int(v5))}
                                                                                                _ = v139
                                                                                            default:
                                                                                                v3 = 139
                                                                                            }
                                                                                            fmt.Println(v2)
                                                                                            v140 := []int{fi115, (v104 + (^(fi115)))}
                                                                                            v141 := GenS9{F0: v2}
                                                                                            _ = v140
                                                                                            _ = v141
                                                                                        }
                                                                                    }
                                                                                    var v142 bool = true
                                                                                    if (uint16(v1) >= uint16(v1)) {
                                                                                        var v143 int16 = -32768
                                                                                        v2 += (v2 & 42)
                                                                                        v144 := []int{(fi115 & (fi71 / ((int(v143)) | 1))), (v104 * (v104 << (((v1 % (((^(v1))) | 1))) & 63))), 1, int(v3)}
                                                                                        switch (v104 - (fi115 + (v104 / ((128) | 1)))) {
                                                                                        case 32767:
                                                                                            fmt.Println(v4, fi71, v4)
                                                                                            fmt.Println(v143)
                                                                                            var v145 uint = (v1 & v1)
                                                                                            _ = v145
                                                                                        case 32768:
                                                                                            var v146 float64 = 1e10
                                                                                            _ = v146
                                                                                        }
                                                                                        _ = v143
                                                                                        _ = v144
                                                                                    }
                                                                                    v147 := []int{v104}
                                                                                    _ = v142
                                                                                    _ = v147
                                                                                } else {
                                                                                    fmt.Println(v8, v1, fi115)
                                                                                    var v148 int32 = 2147483647
                                                                                    switch (v104 | (fi71 | v104)) {
                                                                                    case 1:
                                                                                        v149 := []int{}
                                                                                        _ = v149
                                                                                    case 441:
                                                                                        switch (fi115 & v104) {
                                                                                        case 1000:
                                                                                            var v150 int8 = (v4 & -60)
                                                                                            _ = v150
                                                                                        case 32768:
                                                                                            if (true && true) {
                                                                                                var v151 uint32 = v3
                                                                                                var v152 uint16 = (v8 & (v8 / ((v8) | 1)))
                                                                                                v153 := GenS10{F0: int8(fi71), F1: v2}
                                                                                                _ = v151
                                                                                                _ = v152
                                                                                                _ = v153
                                                                                            } else {
                                                                                                fmt.Println(v74, v8)
                                                                                                switch ("" + "foo bar") {
                                                                                                case "hello":
                                                                                                    fmt.Println(v3, fi115, v148)
                                                                                                default:
                                                                                                    fmt.Println(v125, v8)
                                                                                                    if true {
                                                                                                        if false {
                                                                                                            v2 = (^(v2))
                                                                                                            var v154 float32 = 3.14159
                                                                                                            v5 += v148
                                                                                                            var v155 int16 = 42
                                                                                                            var v156 int32 = (-(v148))
                                                                                                            v157 := []int{(v104 >> ((2) & 63)), v104, (v104 + (fi115 >> ((v1) & 63))), 32767}
                                                                                                            v154 -= (v154 * v154)
                                                                                                            _ = len(v157)
                                                                                                            if false {
                                                                                                                v8 -= (v8 * (^(v8)))
                                                                                                                v125 = (^(v125))
                                                                                                                v1 = 1
                                                                                                            }
                                                                                                            fmt.Println(v2, v148, v154)
                                                                                                            fmt.Println(v148)
                                                                                                            _ = v154
                                                                                                            _ = v155
                                                                                                            _ = v156
                                                                                                            _ = v157
                                                                                                        } else {
                                                                                                            switch (fi71 | (v104 >> (((v1 ^ 4294967295)) & 63))) {
                                                                                                            case 32767:
                                                                                                                v158 := []int{fi71, -2147483648, v104, (fi71 | fi71)}
                                                                                                                _ = v158
                                                                                                            case 947:
                                                                                                                v159 := []int{(^(fi71)), (fi71 | (v104 % ((v104) | 1)))}
                                                                                                                _ = v159
                                                                                                            default:
                                                                                                                v1 -= v1
                                                                                                            }
                                                                                                            var v160 int16 = int16(v2)
                                                                                                            _ = v160
                                                                                                        }
                                                                                                        var v161 uint8 = 42
                                                                                                        v162 := []int{(-(fi115)), 42, 128}
                                                                                                        _ = v161
                                                                                                        _ = v162
                                                                                                    } else {
                                                                                                        if (v74 < v74) {
                                                                                                            v163 := GenS11{F0: true, F1: v2, F2: 536}
                                                                                                            v74 = -2.25
                                                                                                            if false {
                                                                                                                v2 -= v2
                                                                                                                var v164 int = (fi71 << ((v1) & 63))
                                                                                                                var v165 uint16 = v8
                                                                                                                _ = v164
                                                                                                                _ = v165
                                                                                                            } else {
                                                                                                                v166 := GenS12{F0: -841, F1: (v4 / ((42) | 1)), F2: 32767}
                                                                                                                _ = v166
                                                                                                            }
                                                                                                            _ = v163
                                                                                                        }
                                                                                                    }
                                                                                                    v148 += v5
                                                                                                    v4 -= v4
                                                                                                    v104 = 0
                                                                                                    if ((v1 ^ v1) != 820) {
                                                                                                        fmt.Println(fi71)
                                                                                                        switch ("foo bar" + "") {
                                                                                                        case "a":
                                                                                                            switch "" {
                                                                                                            case "hello":
                                                                                                                v8 += 65535
                                                                                                            default:
                                                                                                                var v167 float32 = -2.25
                                                                                                                _ = v167
                                                                                                            }
                                                                                                            switch "a" {
                                                                                                            case "a":
                                                                                                                var v168 int64 = (v125 * v125)
                                                                                                                _ = v168
                                                                                                            case "\n\t":
                                                                                                                if (!true) {
                                                                                                                    if (("hello" + ("a" + "a")) == "foo bar") {
                                                                                                                        v169 := []int{(fi115 & (v104 >> (((v1 | v1)) & 63)))}
                                                                                                                        var v170 uint16 = (v8 ^ (v8 + v8))
                                                                                                                        _ = v169
                                                                                                                        _ = v170
                                                                                                                    } else {
                                                                                                                        var v171 uint16 = (v8 | (^(v8)))
                                                                                                                        v74 = float32(3.14159)
                                                                                                                        v4 = v4
                                                                                                                        _ = v171
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v172 uint16 = v8
                                                                                                                    _ = v172
                                                                                                                }
                                                                                                            }
                                                                                                            v173 := GenS12{F0: (fi115 >> ((4294967295) & 63)), F1: -17, F2: (fi71 & fi71)}
                                                                                                            _ = v173
                                                                                                        }
                                                                                                        switch "foo bar" {
                                                                                                        case "\n\t":
                                                                                                            v174 := []int{4294967295, 128}
                                                                                                            _ = v174
                                                                                                        }
                                                                                                        switch 2147483647 {
                                                                                                        case 9223372036854775807:
                                                                                                            fmt.Println(v5, v2, v4)
                                                                                                        case 32767:
                                                                                                            v175 := []int{-1}
                                                                                                            _ = v175
                                                                                                        case 128:
                                                                                                            switch "\n\t" {
                                                                                                            case "":
                                                                                                                if ((v3 / ((uint32(v8)) | 1)) == 462) {
                                                                                                                    v176 := GenS9{F0: v2}
                                                                                                                    _ = v176
                                                                                                                } else {
                                                                                                                    var v177 uint16 = v8
                                                                                                                    var v178 uint64 = (v2 - (v2 | 1))
                                                                                                                    v179 := []int{}
                                                                                                                    _ = v177
                                                                                                                    _ = v178
                                                                                                                    _ = v179
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        if false {
                                                                                                            v74 = 0.0
                                                                                                        } else {
                                                                                                            v3 = (v3 << ((v1) & 63))
                                                                                                            switch "foo bar" {
                                                                                                            case "":
                                                                                                                var v180 uint32 = v3
                                                                                                                _ = v180
                                                                                                            }
                                                                                                            v181 := GenS11{F0: (false || true), F1: v2, F2: int32(v3)}
                                                                                                            _ = v181
                                                                                                        }
                                                                                                        v182 := []int{(-(fi71)), 764, (fi115 % (((^(fi71))) | 1)), 32767}
                                                                                                        v2 = v2
                                                                                                        if len(v182) > 0 {
                                                                                                            v182[0] = 298
                                                                                                        }
                                                                                                        _ = v182
                                                                                                    }
                                                                                                    v183 := GenS13{F0: (v8 | 32767)}
                                                                                                    _ = v183
                                                                                                }
                                                                                                var v184 int = -1
                                                                                                _ = v184
                                                                                            }
                                                                                        }
                                                                                    case -406:
                                                                                        v185 := GenS14{F0: ("\n\t" + (((("\n\t" + ("foo bar" + ("\n\t" + "\n\t"))) + "a") + "hello") + "hello")), F1: (v5 ^ v148), F2: v4}
                                                                                        _ = v185
                                                                                    default:
                                                                                        switch "hello" {
                                                                                        case "hello":
                                                                                            v1 = v1
                                                                                        case "a":
                                                                                            if false {
                                                                                                var v186 int = 32768
                                                                                                _ = v186
                                                                                            }
                                                                                        case "":
                                                                                            v187 := GenS6{F0: fi71, F1: (v8 % ((1) | 1))}
                                                                                            _ = v187
                                                                                        default:
                                                                                            if (((1e10 - (((1.5 / 1e10) * (0.1 + float64(-2.25))) + ((1e10 - (0.1 + ((1e10 - 1.5) + ((3.14159 * float64(0.1)) / (3.14159 / -2.25))))) - 0.1))) < (1.5 + 3.14159)) && ((v8 | uint16(v3)) <= 65535)) {
                                                                                                switch "a" {
                                                                                                case "":
                                                                                                    var v188 int16 = 32767
                                                                                                    _ = v188
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "\n\t":
                                                                                        var v189 bool = false
                                                                                        _ = v189
                                                                                    case "":
                                                                                        fmt.Println(v148, v4, v4)
                                                                                    }
                                                                                    v190 := GenS10{F0: (v4 - v4), F1: (v2 << (((^(v1))) & 63))}
                                                                                    v191 := []int{}
                                                                                    v192 := GenS3{F0: (v8 / ((v8) | 1)), F1: (v3 | v3)}
                                                                                    _ = v148
                                                                                    _ = v190
                                                                                    _ = v191
                                                                                    _ = v192
                                                                                }
                                                                                var v193 int = -214
                                                                                _ = v193
                                                                            } else {
                                                                                v194 := []int{9223372036854775807, 9223372036854775807, fi71, v104}
                                                                                if (666 >= 42) {
                                                                                    v4 = (v4 * (v4 / ((v4) | 1)))
                                                                                    var v195 int16 = -449
                                                                                    _ = v195
                                                                                }
                                                                                _ = v194
                                                                            }
                                                                        }
                                                                        var v196 int8 = 0
                                                                        var v197 string = ((("hello" + "foo bar") + ("foo bar" + "foo bar")) + ("\n\t" + "hello"))
                                                                        _ = v196
                                                                        _ = v197
                                                                    } else {
                                                                        if true {
                                                                            var v198 int16 = 1
                                                                            v199 := []int{}
                                                                            var v200 float32 = v74
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                var v201 int64 = -1
                                                                                _ = v201
                                                                            }
                                                                            var v202 int = (fi115 % ((32767) | 1))
                                                                            _ = v198
                                                                            _ = v199
                                                                            _ = v200
                                                                            _ = v202
                                                                        } else {
                                                                            v1 -= (v1 | 0)
                                                                            switch (("a" + "a") + "a") {
                                                                            case "a":
                                                                                fmt.Println(fi115)
                                                                            case "":
                                                                                if (v8 != (v8 % ((699) | 1))) {
                                                                                    var v203 int64 = -166
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        switch fi71 {
                                                                                        case 80:
                                                                                            v125 = -992
                                                                                        default:
                                                                                            v125 -= (^(v203))
                                                                                        }
                                                                                    default:
                                                                                        v1 = (v1 >> ((uint(v8)) & 63))
                                                                                    }
                                                                                    if (float64(((v74 - 1e10) / float32(float64(-2.25)))) >= 1.5) {
                                                                                        var v204 float64 = 3.14159
                                                                                        _ = v204
                                                                                    }
                                                                                    _ = v203
                                                                                }
                                                                            case "hello":
                                                                                v205 := []int{(fi71 << ((2) & 63)), v104, (v104 << ((0) & 63))}
                                                                                _ = v205
                                                                            default:
                                                                                v5 += 42
                                                                            }
                                                                        }
                                                                        switch (("hello" + "") + "") {
                                                                        case "foo bar":
                                                                            v2 += 778
                                                                        case "a":
                                                                            if (uint(v8) < 255) {
                                                                                if ((v2 >= v2) || (65535 != (v8 ^ (^(v8))))) {
                                                                                    fmt.Println(v125)
                                                                                    v206 := []int{fi115}
                                                                                    var v207 uint64 = v2
                                                                                    fmt.Println(v1)
                                                                                    _ = v206
                                                                                    _ = v207
                                                                                } else {
                                                                                    v208 := []int{(v104 * (-(fi115)))}
                                                                                    var v209 int16 = -250
                                                                                    v5 = (-(v5))
                                                                                    v208 = append(v208, v104)
                                                                                    _ = v208
                                                                                    _ = v209
                                                                                }
                                                                                v210 := GenS10{F0: (v4 + (v4 ^ v4)), F1: (v2 / (((v2 + (^(v2)))) | 1))}
                                                                                v2 = v2
                                                                                _ = v210
                                                                            }
                                                                        case "hello":
                                                                            fmt.Println(v1, v5, v4)
                                                                        default:
                                                                            v1 = 4294967295
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    if true {
                                                                        var v211 string = "a"
                                                                        fmt.Println(v8, v3)
                                                                        _ = v211
                                                                    } else {
                                                                        v212 := []int{(^(v104)), -2147483648, fi71}
                                                                        var v213 float64 = (1e10 + float64((v74 - float32(float64(v74)))))
                                                                        fmt.Println(v74, v104, v5)
                                                                        var v214 uint = uint(v8)
                                                                        var v215 uint64 = uint64(v214)
                                                                        fmt.Println(v214, v2)
                                                                        v104 -= (v104 + (-(fi71)))
                                                                        _ = v212
                                                                        _ = v213
                                                                        _ = v214
                                                                        _ = v215
                                                                    }
                                                                }
                                                                fmt.Println(v5)
                                                                v216 := GenS13{F0: v8}
                                                                var v217 uint16 = 1
                                                                _ = v216
                                                                _ = v217
                                                            } else {
                                                                var v218 int8 = (v4 ^ 100)
                                                                _ = v218
                                                            }
                                                            v2 = uint64(v8)
                                                            _ = v125
                                                        }
                                                    case 127:
                                                        switch ("" + "hello") {
                                                        case "\n\t":
                                                            v1 -= v1
                                                        case "foo bar":
                                                            v8 += (^(v8))
                                                        case "":
                                                            v219 := []int{}
                                                            _ = v219
                                                        default:
                                                            if true {
                                                                v3 += 2147483648
                                                                v8 -= (v8 / (((v8 % ((93) | 1))) | 1))
                                                            }
                                                        }
                                                    case 128:
                                                        var v220 float64 = 0.1
                                                        _ = v220
                                                    default:
                                                        v1 += 2
                                                    }
                                                    v221 := []int{fi115, fi71}
                                                    _ = v221
                                                } else {
                                                    if true {
                                                        v222 := []int{(v104 & int(v1)), (fi115 ^ -9223372036854775808), (v104 | 127)}
                                                        v1 += uint(v3)
                                                        switch (fi71 / ((-9223372036854775808) | 1)) {
                                                        case 127:
                                                            v223 := GenS15{F0: (float64((v74 - (v74 + float32(-2.25)))) != 1e10)}
                                                            _ = v223
                                                        case -414:
                                                            v74 -= (((3.14159 / v74) + v74) - -2.25)
                                                        default:
                                                            if len(v222) > 0 {
                                                                v222[0] = (fi71 % ((v104) | 1))
                                                            }
                                                        }
                                                        _ = v222
                                                    }
                                                    v224 := GenS13{F0: 42}
                                                    v225 := []int{42}
                                                    _ = v224
                                                    _ = v225
                                                }
                                            case "foo bar":
                                                v74 += -2.25
                                            }
                                            _ = fi115
                                        }
                                    }
                                case "":
                                    v1 -= (v1 & 889)
                                }
                            case -9223372036854775808:
                                var v226 uint16 = 42
                                _ = v226
                            default:
                                v1 = (v1 & v1)
                            }
                            v2 -= (v2 << (((v1 & 255)) & 63))
                            _ = v104
                        }
                        v227 := GenS9{F0: v2}
                        _ = v227
                    }
                    switch fi71 {
                    case 1:
                        var v228 uint64 = v2
                        _ = v228
                    case -582:
                        for fi229 := 0; fi229 < 9; fi229++ {
                            v230 := []int{(^(fi229))}
                            v1 += 65535
                            var v231 int16 = 739
                            if (!(false || (!(fi229 > -264)))) {
                                if (!(790 >= v231)) {
                                    var v232 uint64 = 792
                                    if (!true) {
                                        v8 -= (^(v8))
                                        v232 = v232
                                        v3 += v3
                                        switch fi229 {
                                        case 1:
                                            var v233 uint8 = 34
                                            if true {
                                                v3 += v3
                                            }
                                            switch ("foo bar" + "hello") {
                                            case "foo bar":
                                                fmt.Println(fi71, v8, v233)
                                            case "":
                                                var v234 int32 = v5
                                                _ = v234
                                            case "\n\t":
                                                v230 = append(v230, (fi71 * 4294967295))
                                            }
                                            _ = v233
                                        case 991:
                                            var v235 uint = 303
                                            _ = v235
                                        case 816:
                                            v236 := GenS16{F0: v1, F1: (v3 ^ v3), F2: v8}
                                            _ = v236
                                        }
                                        var v237 uint8 = 0
                                        _ = v237
                                    }
                                    if (v1 != (v1 - (v1 >> ((v1) & 63)))) {
                                        switch (fi229 % (((fi71 ^ 32767)) | 1)) {
                                        case 42:
                                            if true {
                                                v5 += v5
                                                var v238 int8 = int8(fi71)
                                                var v239 float32 = 0.1
                                                _ = v238
                                                _ = v239
                                            } else {
                                                switch (^(fi229)) {
                                                case -310:
                                                    v4 -= v4
                                                case 761:
                                                    var v240 uint64 = (^(v232))
                                                    _ = v240
                                                case 42:
                                                    if len(v230) > 0 {
                                                        v230[0] = (fi229 >> ((255) & 63))
                                                    }
                                                default:
                                                    v241 := GenS14{F0: "hello", F1: int32(v3), F2: v4}
                                                    _ = v241
                                                }
                                            }
                                        case 127:
                                            _ = len(v230)
                                        case -2147483648:
                                            var v242 int16 = (v231 / ((v231) | 1))
                                            _ = v242
                                        default:
                                            switch -1 {
                                            case 32767:
                                                v1 = v1
                                            default:
                                                v8 -= 423
                                            }
                                        }
                                    }
                                    _ = v232
                                } else {
                                    var v243 bool = (42 >= (v231 % ((v231) | 1)))
                                    switch ((((("a" + ("foo bar" + "foo bar")) + "") + "a") + "\n\t") + ((("a" + ("hello" + "hello")) + "hello") + "\n\t")) {
                                    case "\n\t":
                                        switch (fi229 ^ fi71) {
                                        case 1:
                                            v244 := GenS17{F0: 1, F1: ((v243 && (v74 >= v74)) && false)}
                                            _ = v244
                                        default:
                                            if (true || ("foo bar" < "foo bar")) {
                                                fmt.Println(v5, v231, v8)
                                                if (879 < (v8 - 761)) {
                                                    v245 := GenS18{F0: 4294967295, F1: 61, F2: (v3 ^ (^(v3)))}
                                                    switch int(v4) {
                                                    case 42:
                                                        if true {
                                                            if len(v230) > 0 {
                                                                v230[0] = 42
                                                            }
                                                        }
                                                        v246 := GenS19{F0: v4, F1: int64(v4)}
                                                        v8 = (v8 << (((v1 & v1)) & 63))
                                                        _ = v246
                                                    }
                                                    _ = v245
                                                }
                                                fmt.Println(v1, v1, v2)
                                            } else {
                                                if len(v230) > 0 {
                                                    v230[0] = fi71
                                                }
                                                if (v4 < (v4 << (((v1 * (v1 % (((v1 + 718)) | 1)))) & 63))) {
                                                    v243 = (true && v243)
                                                } else {
                                                    if false {
                                                        fmt.Println(v231)
                                                    }
                                                }
                                            }
                                            var v247 int8 = -4
                                            var v248 uint64 = 4294967295
                                            fmt.Println(v247, v231)
                                            v2 -= (^(v248))
                                            _ = v247
                                            _ = v248
                                        }
                                    }
                                    v231 += v231
                                    _ = v243
                                }
                                v231 -= (v231 / (((^(v231))) | 1))
                            } else {
                                var v249 int32 = int32(v4)
                                v8 -= 602
                                _ = v249
                            }
                            var v250 uint64 = (^(v2))
                            _ = fi229
                            _ = v230
                            _ = v231
                            _ = v250
                        }
                    case 32767:
                        if ((v4 / (((v4 / ((v4) | 1))) | 1)) < v4) {
                            var v251 int8 = v4
                            if true {
                                for fi252 := 0; fi252 < 18; fi252++ {
                                    if true {
                                        switch (fi252 / ((fi71) | 1)) {
                                        case 32768:
                                            switch int(v4) {
                                            case -1:
                                                fmt.Println(v251, v1)
                                            case 128:
                                                var v253 float32 = (0.1 - 3.14159)
                                                _ = v253
                                            case -57:
                                                var v254 uint64 = 2
                                                _ = v254
                                            default:
                                                fmt.Println(v74)
                                            }
                                        default:
                                            v255 := GenS17{F0: v2, F1: true}
                                            _ = v255
                                        }
                                    } else {
                                        switch (^(fi252)) {
                                        case 127:
                                            var v256 uint = v1
                                            _ = v256
                                        case 4294967295:
                                            var v257 uint16 = v8
                                            _ = v257
                                        case -493:
                                            v258 := []int{}
                                            _ = v258
                                        default:
                                            if false {
                                                if false {
                                                    var v259 int = (fi252 % ((fi252) | 1))
                                                    _ = v259
                                                }
                                                fmt.Println(v2, fi71)
                                            } else {
                                                v251 = int8(v3)
                                            }
                                            var v260 int = (fi71 / ((fi252) | 1))
                                            _ = v260
                                        }
                                        fmt.Println(fi252, v4, v5)
                                    }
                                    v261 := []int{}
                                    _ = fi252
                                    _ = v261
                                }
                            }
                            _ = v251
                        } else {
                            var v262 uint = (^(v1))
                            v263 := GenS18{F0: v3, F1: 255, F2: (v3 & (v3 & 0))}
                            if false {
                                var v264 int16 = 32767
                                _ = v264
                            }
                            _ = v262
                            _ = v263
                        }
                    }
                }
                _ = fi71
                _ = v74
            }
            var v265 float64 = 0.1
            v266 := []int{691, 4294967295, int(v3)}
            _ = v265
            _ = v266
        }
        if ((v4 % ((127) | 1)) < (-(v4))) {
            v267 := []int{-322, 2147483647}
            _ = v267
        }
    }
    switch 42 {
    case 708:
        v268 := GenS3{F0: (^(v8)), F1: (v3 / ((v3) | 1))}
        _ = v268
    default:
        v269 := []int{128, -2147483648}
        _ = v269
    }
    switch "\n\t" {
    case "\n\t":
        var v270 int = 32767
        _ = v270
    }
    fmt.Println(v3, v2, v5)
    switch -102 {
    case -1:
        if true {
            fmt.Println(v3)
        }
    }
    var v271 int64 = 1
    v2 += v2
    if (!true) {
        v5 = v5
    }
    v272 := []int{32767, -323}
    v3 -= 921
    switch "\n\t" {
    case "hello":
        for fi273 := 0; fi273 < 10; fi273++ {
            var v274 int64 = (v271 - v271)
            var v275 bool = false
            v276 := GenS20{F0: uint64(v3), F1: 280, F2: int(v4)}
            _ = fi273
            _ = v274
            _ = v275
            _ = v276
        }
    }
    for fi277 := 0; fi277 < 4; fi277++ {
        var v278 uint8 = 1
        for fi279 := 0; fi279 < 4; fi279++ {
            if false {
                v8 += (^(v8))
            }
            v271 -= v271
            if false {
                var v280 uint64 = v2
                if len(v272) > 0 {
                    v272[0] = fi277
                }
                fmt.Println(v271, v2)
                if ((float64(((1.5 + (0.0 / 1e10)) * float32((3.14159 * 0.0)))) + -2.25) > 3.14159) {
                    var v281 uint64 = (v280 + v2)
                    var v282 float64 = (0.0 - 0.0)
                    v272 = append(v272, (fi279 % (((^(fi279))) | 1)))
                    v282 += 3.14159
                    var v283 bool = false
                    _ = v281
                    _ = v282
                    _ = v283
                } else {
                    var v284 float32 = 0.1
                    if len(v272) > 0 {
                        v272[0] = fi277
                    }
                    var v285 float64 = (1.5 + 3.14159)
                    _ = v284
                    _ = v285
                }
                var v286 uint8 = v278
                _ = v280
                _ = v286
            }
            _ = fi279
        }
        fmt.Println(v8, v5)
        _ = fi277
        _ = v278
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v271:", v271)
    fmt.Println("slicelen v272:", len(v272), v272)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v271
    _ = v272
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
