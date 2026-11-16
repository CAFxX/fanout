package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int8; F1 int32; F2 int }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint8; F1 string }
type GenS11 struct { F0 int16; F1 uint16; F2 int }
type GenS12 struct { F0 uint; F1 uint8 }
type GenS13 struct { F0 int16; F1 int64 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int16; F1 int16 }
type GenS16 struct { F0 int; F1 uint32; F2 int8 }
type GenS17 struct { F0 uint32; F1 string; F2 uint32 }
type GenS18 struct { F0 uint8; F1 bool }
type GenS19 struct { F0 uint64; F1 int32; F2 int }
func genFn0(p0 int8) int {
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(fi7, p0)
        v8 := []int{-459, (fi7 / ((-578) | 1))}
        for fi9 := 0; fi9 < 3; fi9++ {
            switch 0 {
            case 32768:
                if (!true) {
                    if (2147483647 == 1) {
                        p0 = p0
                    }
                    switch fi9 {
                    case 174:
                        var v10 uint32 = 4294967295
                        _ = v10
                    case 32767:
                        if true {
                            var v11 bool = true
                            var v12 uint64 = 2
                            _ = v11
                            _ = v12
                        }
                    }
                } else {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v13 int = (fi7 ^ 32767)
                    _ = len(v8)
                    _ = v13
                }
            case 32767:
                var v14 float32 = ((1.5 / 3.14159) + 0.0)
                _ = v14
            case -2147483648:
                if false {
                    _ = len(v8)
                    fmt.Println(fi9, p0)
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                } else {
                    _ = len(v8)
                }
                fmt.Println(p0, fi9)
            }
            var v15 string = "\n\t"
            v8 = append(v8, (-(fi7)))
            p0 -= (-(p0))
            fmt.Println(p0)
            _ = fi9
            _ = v15
        }
        _ = fi7
        _ = v8
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        if ((561 != -83) && ((42 > 0) || (fi16 >= 127))) {
            if false {
                fmt.Println(p0)
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := []int{127, (fi17 & 32768), fi16, (fi16 & (fi16 % ((fi16) | 1)))}
                    _ = fi17
                    _ = v18
                }
            } else {
                p0 += (p0 & (p0 ^ p0))
                switch 2147483647 {
                case -2147483648:
                    v19 := GenS1{F0: 2, F1: 0, F2: 0}
                    _ = v19
                case -9223372036854775808:
                    v20 := []int{fi16, 989, fi16}
                    _ = v20
                }
            }
            for fi21 := 0; fi21 < 20; fi21++ {
                p0 += (p0 | p0)
                _ = fi21
            }
            switch "hello" {
            case "hello":
                p0 = p0
            }
        } else {
            fmt.Println(p0, p0)
        }
        p0 -= p0
        _ = fi16
    }
    _ = p0
    return 467
}
func genFn1(p0 int8, p1 string) int32 {
    v22 := []int{}
    v23 := GenS1{F0: 1, F1: 1, F2: 32767}
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 uint8 = 127
        _ = len(v22)
        for fi26 := 0; fi26 < 1; fi26++ {
            switch 4294967295 {
            case 32767:
                v27 := GenS1{F0: 0, F1: 4294967295, F2: 32767}
                _ = v27
            case 0:
                if false {
                    switch (fi26 + (fi26 >> ((0) & 63))) {
                    case 2147483647:
                        p1 = "foo bar"
                    case 544:
                        p0 += (-(p0))
                    }
                    var v28 int8 = (p0 + (p0 >> ((2) & 63)))
                    v28 = 42
                    _ = v28
                } else {
                    switch (p1 + "foo bar") {
                    case "hello":
                        var v29 uint8 = 212
                        _ = v29
                    case "\n\t":
                        v30 := GenS1{F0: 65535, F1: 2, F2: 32767}
                        _ = v30
                    default:
                        switch (fi24 % ((42) | 1)) {
                        case 757:
                            v31 := GenS1{F0: 65535, F1: 492, F2: 1000}
                            _ = v31
                        case 42:
                            switch fi26 {
                            case 234:
                                if false {
                                    v25 -= (^(v25))
                                    var v32 float32 = -2.25
                                    _ = v32
                                } else {
                                    switch (p1 + p1) {
                                    case "foo bar":
                                        fmt.Println(p1, p0)
                                    default:
                                        fmt.Println(v25)
                                    }
                                }
                            case 9223372036854775807:
                                v33 := GenS1{F0: 689, F1: 14, F2: -32768}
                                _ = v33
                            case 1:
                                fmt.Println(p0, fi26, v25)
                            default:
                                if (fi26 > fi26) {
                                    v34 := GenS1{F0: 862, F1: 2147483648, F2: 0}
                                    var v35 string = ("\n\t" + p1)
                                    p1 = ("hello" + (v35 + (v35 + p1)))
                                    fmt.Println(v25, p0)
                                    _ = v34
                                    _ = v35
                                }
                                v25 += v25
                            }
                        case 0:
                            v36 := GenS1{F0: 555, F1: 2147483647, F2: -898}
                            _ = v36
                        default:
                            p0 = (p0 | 100)
                        }
                    }
                    var v37 bool = true
                    var v38 uint32 = 1
                    _ = v37
                    _ = v38
                }
            }
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    if false {
        v22 = append(v22, 127)
        if false {
            if ((true || true) || (((((float32((3.14159 / 3.14159)) - ((0.0 - (0.0 - 0.0)) - 1e10)) + 3.14159) == float32((0.1 + 0.1))) || false) && true)) {
                v39 := GenS1{F0: 65535, F1: 379, F2: -1}
                _ = v39
            }
            for fi40 := 0; fi40 < 8; fi40++ {
                if len(v22) > 0 {
                    v22[0] = genFn0(p0)
                }
                for fi41 := 0; fi41 < 16; fi41++ {
                    if false {
                        if len(v22) > 0 {
                            v22[0] = (^(fi41))
                        }
                        var v42 int32 = 13
                        if true {
                            if (!(!true)) {
                                switch "\n\t" {
                                case "hello":
                                    var v43 float64 = 1.5
                                    _ = v43
                                }
                                var v44 float64 = -2.25
                                switch ("hello" + "") {
                                case "a":
                                    v22 = append(v22, (fi40 + fi40))
                                case "":
                                    p1 = p1
                                case "\n\t":
                                    v44 += -2.25
                                }
                                _ = v44
                            }
                            if (42 > 0) {
                                if false {
                                    p0 = 1
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                }
                            }
                            if false {
                                fmt.Println(p1, p0, p1)
                            } else {
                                if (float64(((3.14159 - 0.0) / 1e10)) >= (((1e10 * -2.25) + 3.14159) - float64(1.5))) {
                                    switch fi40 {
                                    case 9223372036854775807:
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                    case -57:
                                        var v45 float32 = 1.5
                                        _ = v45
                                    case 32768:
                                        var v46 string = p1
                                        _ = v46
                                    }
                                } else {
                                    v47 := GenS2{F0: 255, F1: 554}
                                    var v48 int8 = (^(p0))
                                    _ = v47
                                    _ = v48
                                }
                            }
                            switch "a" {
                            case "a":
                                switch p1 {
                                case "hello":
                                    switch p1 {
                                    case "foo bar":
                                        p0 -= 1
                                    }
                                case "foo bar":
                                    var v49 uint16 = 2
                                    _ = v49
                                case "a":
                                    p1 = (("hello" + "a") + (p1 + ("" + p1)))
                                default:
                                    v22 = append(v22, (fi40 | -226))
                                    v22 = append(v22, genFn0(p0))
                                }
                            }
                        }
                        _ = v42
                    }
                    switch "foo bar" {
                    case "\n\t":
                        _ = len(v22)
                    case "foo bar":
                        v50 := GenS1{F0: 65535, F1: 2147483648, F2: 42}
                        _ = v50
                    case "a":
                        p0 = -128
                    }
                    switch genFn0((p0 ^ (-(p0)))) {
                    case -217:
                        var v51 uint8 = 42
                        _ = v51
                    case 2147483647:
                        p1 = (("\n\t" + p1) + "hello")
                    default:
                        if false {
                            _ = len(v22)
                            if false {
                                var v52 int64 = -723
                                _ = len(v22)
                                if (((!false) && ((382 >= 65535) && true)) || (true || true)) {
                                    var v53 uint = 147
                                    v54 := GenS2{F0: 0, F1: 259}
                                    _ = v53
                                    _ = v54
                                } else {
                                    var v55 float32 = 3.14159
                                    p1 = "hello"
                                    _ = v55
                                }
                                _ = v52
                            } else {
                                switch p1 {
                                case "hello":
                                    var v56 uint8 = 83
                                    v57 := GenS3{F0: (^(p0)), F1: -288, F2: genFn0(p0)}
                                    v58 := GenS1{F0: 2147483648, F1: 0, F2: -1}
                                    var v59 uint = 18446744073709551615
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                default:
                                    var v60 string = ("" + p1)
                                    switch ("hello" + p1) {
                                    case "foo bar":
                                        var v61 int64 = 513
                                        _ = v61
                                    case "":
                                        var v62 uint16 = 65535
                                        _ = v62
                                    case "\n\t":
                                        v63 := GenS4{F0: 1, F1: -931}
                                        p1 = (((("\n\t" + (p1 + ((p1 + p1) + "hello"))) + v60) + p1) + "a")
                                        _ = v63
                                    }
                                    _ = v60
                                }
                                if true {
                                    p0 -= p0
                                    p0 -= p0
                                    if false {
                                        if false {
                                            if (genFn0(p0) <= genFn0(p0)) {
                                                v64 := GenS3{F0: (p0 - (p0 - (p0 + 100))), F1: -732, F2: int(p0)}
                                                _ = v64
                                            }
                                            if true {
                                                p0 = 0
                                            }
                                            var v65 uint32 = 990
                                            _ = v65
                                        }
                                    }
                                } else {
                                    v66 := GenS5{F0: 1, F1: 111}
                                    if (4294967295 >= 65535) {
                                        v67 := GenS4{F0: 0, F1: -180}
                                        _ = v67
                                    }
                                    _ = v66
                                }
                                var v68 string = (("foo bar" + ("hello" + p1)) + p1)
                                _ = v68
                            }
                            p1 = ""
                        } else {
                            if true {
                                if (true && (!(42 <= 65535))) {
                                    fmt.Println(fi40)
                                } else {
                                    var v69 float32 = ((float32((3.14159 / ((((-2.25 - 1.5) + 3.14159) * 1.5) + (-2.25 + 0.0)))) * ((1.5 / (3.14159 * (1e10 + (1.5 + ((-2.25 + (float32(((((0.1 / -2.25) * (0.0 / 1.5)) + 3.14159) - 3.14159)) - 1.5)) / float32(1e10)))))) * float32((1.5 + 0.1)))) * -2.25)
                                    v70 := GenS6{F0: -1, F1: fi41}
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                    _ = v69
                                    _ = v70
                                }
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            } else {
                                p1 = (("\n\t" + "") + p1)
                                var v71 int64 = 1
                                var v72 float32 = 1.5
                                if ((v71 > -361) || (false && (true && true))) {
                                    var v73 uint16 = 0
                                    _ = v73
                                }
                                _ = v71
                                _ = v72
                            }
                            switch fi41 {
                            case 0:
                                p1 = ((p1 + p1) + "")
                            case -866:
                                v74 := GenS7{F0: 540}
                                _ = v74
                            default:
                                fmt.Println(fi40, fi41, p0)
                            }
                            _ = len(v22)
                        }
                    }
                    _ = fi41
                }
                p0 = (p0 - p0)
                _ = fi40
            }
        } else {
            for fi75 := 0; fi75 < 1; fi75++ {
                var v76 int16 = -771
                if true {
                    p0 = (p0 ^ (p0 * p0))
                } else {
                    switch 1 {
                    case 1:
                        if false {
                            v77 := GenS4{F0: 2147483647, F1: 0}
                            var v78 int8 = (p0 & p0)
                            v22 = append(v22, 128)
                            _ = v77
                            _ = v78
                        } else {
                            p0 += (p0 >> ((481) & 63))
                            p0 -= p0
                            v76 = (v76 | v76)
                        }
                    default:
                        if (42 < 467) {
                            v79 := GenS5{F0: -298, F1: 255}
                            for fi80 := 0; fi80 < 3; fi80++ {
                                switch (fi75 ^ 4294967295) {
                                case -986:
                                    v81 := GenS8{F0: (v76 << ((836) & 63))}
                                    _ = v81
                                case 0:
                                    v82 := GenS9{F0: -261}
                                    _ = v82
                                case 9223372036854775807:
                                    var v83 uint = 65535
                                    _ = v83
                                }
                                _ = fi80
                            }
                            for fi84 := 0; fi84 < 20; fi84++ {
                                switch "hello" {
                                case "":
                                    if (0.1 == (float64(1.5) - 0.1)) {
                                        v76 -= 353
                                        v76 += 42
                                        p1 = p1
                                    } else {
                                        var v85 int8 = p0
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                        _ = len(v22)
                                        _ = v85
                                    }
                                    switch p1 {
                                    case "a":
                                        switch "a" {
                                        case "\n\t":
                                            v76 += v76
                                        case "":
                                            var v86 uint32 = 4294967295
                                            _ = v86
                                        default:
                                            v87 := GenS9{F0: 1}
                                            _ = v87
                                        }
                                    case "foo bar":
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                    case "hello":
                                        if true {
                                            var v88 int32 = 1
                                            var v89 int8 = (-(p0))
                                            v76 = v76
                                            _ = v88
                                            _ = v89
                                        } else {
                                            switch (fi84 + (fi84 & fi84)) {
                                            case 42:
                                                v90 := GenS10{F0: 128, F1: (p1 + (p1 + p1))}
                                                _ = v90
                                            case -23:
                                                if true {
                                                    switch ((p1 + (((p1 + p1) + (p1 + p1)) + ("hello" + "\n\t"))) + "foo bar") {
                                                    case "foo bar":
                                                        v91 := GenS7{F0: 42}
                                                        _ = v91
                                                    }
                                                } else {
                                                    v92 := GenS9{F0: -258}
                                                    p0 += p0
                                                    var v93 int = genFn0(p0)
                                                    if false {
                                                        p0 += 1
                                                    } else {
                                                        switch v93 {
                                                        case 42:
                                                            var v94 int8 = (^(p0))
                                                            _ = v94
                                                        case 2147483647:
                                                            var v95 uint8 = 31
                                                            _ = v95
                                                        case 684:
                                                            v96 := GenS11{F0: -529, F1: 458, F2: fi75}
                                                            _ = v96
                                                        }
                                                        v76 = 0
                                                        if (!true) {
                                                            p0 -= (p0 >> ((0) & 63))
                                                            var v97 int64 = 0
                                                            fmt.Println(v97, v97)
                                                            _ = v97
                                                        }
                                                        v98 := GenS1{F0: 42, F1: 2147483647, F2: (v76 * v76)}
                                                        _ = v98
                                                    }
                                                    _ = v92
                                                    _ = v93
                                                }
                                            }
                                            var v99 int32 = 0
                                            _ = v99
                                        }
                                    }
                                case "hello":
                                    if true {
                                        v100 := GenS2{F0: 0, F1: 263}
                                        var v101 uint32 = 173
                                        if len(v22) > 0 {
                                            v22[0] = fi75
                                        }
                                        _ = v100
                                        _ = v101
                                    } else {
                                        p0 = (p0 % ((-128) | 1))
                                        var v102 uint16 = 2
                                        switch ((p1 + "a") + ("a" + (p1 + p1))) {
                                        case "":
                                            fmt.Println(v76, fi84, fi75)
                                        case "a":
                                            switch fi75 {
                                            case -1:
                                                if (!((v76 >> ((42) & 63)) == -32768)) {
                                                    v76 = v76
                                                } else {
                                                    v103 := GenS3{F0: -83, F1: 1, F2: -2147483648}
                                                    _ = len(v22)
                                                    v104 := GenS10{F0: 42, F1: "a"}
                                                    _ = v103
                                                    _ = v104
                                                }
                                            default:
                                                if len(v22) >= 2 {
                                                    v22 = v22[1:len(v22)]
                                                }
                                            }
                                            fmt.Println(p0, fi84)
                                            var v105 uint8 = 0
                                            _ = v105
                                        }
                                        switch (fi84 >> ((42) & 63)) {
                                        case -2147483648:
                                            v106 := GenS12{F0: 65535, F1: 128}
                                            _ = v106
                                        }
                                        if true {
                                            v107 := GenS13{F0: (v76 << ((18446744073709551615) & 63)), F1: 9223372036854775807}
                                            fmt.Println(fi75, v76)
                                            _ = v107
                                        } else {
                                            if (1e10 <= float32((((0.0 / 1.5) + (3.14159 * (1e10 - 0.1))) / -2.25))) {
                                                var v108 uint16 = (v102 * v102)
                                                v109 := GenS14{F0: 126}
                                                _ = v108
                                                _ = v109
                                            } else {
                                                if (957 == 0) {
                                                    var v110 uint32 = 647
                                                    if ((0.1 * 1e10) < 0.1) {
                                                        if true {
                                                            var v111 int32 = -589
                                                            if true {
                                                                if len(v22) >= 2 {
                                                                    v22 = v22[1:len(v22)]
                                                                }
                                                                v112 := GenS7{F0: 0}
                                                                _ = v112
                                                            }
                                                            _ = v111
                                                        } else {
                                                            switch genFn0(p0) {
                                                            case 128:
                                                                if (!false) {
                                                                    v102 = 255
                                                                    p1 = (("foo bar" + "") + ("\n\t" + p1))
                                                                    var v113 int = fi84
                                                                    _ = v113
                                                                } else {
                                                                    if (v76 != v76) {
                                                                        var v114 uint = 0
                                                                        v110 -= uint32(v114)
                                                                        v114 -= (^(v114))
                                                                        _ = v114
                                                                    }
                                                                    v102 += (v102 / (((v102 & (v102 | (^(v102))))) | 1))
                                                                    v22 = append(v22, fi84)
                                                                }
                                                            case 459:
                                                                if ((^(v110)) != 2147483648) {
                                                                    v115 := GenS10{F0: 128, F1: p1}
                                                                    _ = v115
                                                                }
                                                            }
                                                            var v116 float32 = 1.5
                                                            if (false && ((v102 < v102) && (!false))) {
                                                                if len(v22) >= 2 {
                                                                    v22 = v22[1:len(v22)]
                                                                }
                                                                switch (fi84 % ((127) | 1)) {
                                                                case 2147483647:
                                                                    switch -9223372036854775808 {
                                                                    case -259:
                                                                        v117 := GenS15{F0: v76, F1: -32768}
                                                                        _ = v117
                                                                    case -9223372036854775808:
                                                                        var v118 uint64 = uint64(v102)
                                                                        _ = v118
                                                                    case 1:
                                                                        fmt.Println(fi84, p1)
                                                                    default:
                                                                        v119 := GenS10{F0: 38, F1: p1}
                                                                        _ = v119
                                                                    }
                                                                case 32768:
                                                                    v120 := GenS8{F0: -858}
                                                                    _ = v120
                                                                }
                                                            } else {
                                                                v121 := GenS4{F0: 39, F1: 2147483647}
                                                                v122 := GenS9{F0: 2147483647}
                                                                v123 := GenS3{F0: (p0 - (p0 + (p0 | -128))), F1: 42, F2: (fi75 << ((467) & 63))}
                                                                _ = v121
                                                                _ = v122
                                                                _ = v123
                                                            }
                                                            if (!((p0 - (^(p0))) >= (p0 * (p0 >> ((255) & 63))))) {
                                                                if false {
                                                                    v116 += ((0.0 * ((v116 / v116) - float32(0.0))) + v116)
                                                                    _ = len(v22)
                                                                    var v124 int32 = 867
                                                                    _ = v124
                                                                }
                                                                v125 := GenS11{F0: (v76 * -1), F1: (v102 / ((42) | 1)), F2: (fi75 ^ fi84)}
                                                                switch (fi84 & fi75) {
                                                                case -2147483648:
                                                                    if (false || (-1 == (p0 & 0))) {
                                                                        var v126 int64 = 9223372036854775807
                                                                        fmt.Println(p1, v76)
                                                                        _ = v126
                                                                    }
                                                                default:
                                                                    fmt.Println(v102, fi84, fi75)
                                                                    var v127 float32 = ((v116 - v116) / ((v116 - (v116 - -2.25)) + (1e10 * (v116 * v116))))
                                                                    v128 := GenS16{F0: genFn0((^(p0))), F1: 0, F2: 0}
                                                                    _ = v127
                                                                    _ = v128
                                                                }
                                                                _ = v125
                                                            } else {
                                                                p1 = "foo bar"
                                                            }
                                                            _ = v116
                                                        }
                                                    } else {
                                                        var v129 uint64 = 1
                                                        v130 := GenS17{F0: 387, F1: p1, F2: (v110 - v110)}
                                                        fmt.Println(p1, fi75, p1)
                                                        _ = v129
                                                        _ = v130
                                                    }
                                                    _ = v110
                                                } else {
                                                    if false {
                                                        if ((0 != 2) && (true && false)) {
                                                            var v131 int32 = 974
                                                            if len(v22) >= 2 {
                                                                v22 = v22[1:len(v22)]
                                                            }
                                                            var v132 int = 0
                                                            v133 := GenS17{F0: 157, F1: "", F2: 2147483647}
                                                            _ = v131
                                                            _ = v132
                                                            _ = v133
                                                        }
                                                        if len(v22) > 0 {
                                                            v22[0] = (fi84 % ((fi84) | 1))
                                                        }
                                                    }
                                                    if (2147483647 < 2) {
                                                        var v134 uint8 = 128
                                                        v76 = 42
                                                        if len(v22) > 0 {
                                                            v22[0] = fi75
                                                        }
                                                        _ = v134
                                                    } else {
                                                        p0 += (p0 & (^(p0)))
                                                        v135 := GenS6{F0: (p0 << ((0) & 63)), F1: fi84}
                                                        fmt.Println(p0, p0)
                                                        if len(v22) >= 2 {
                                                            v22 = v22[1:len(v22)]
                                                        }
                                                        p0 = 100
                                                        _ = v135
                                                    }
                                                }
                                                switch p1 {
                                                case "\n\t":
                                                    var v136 int32 = -1
                                                    _ = v136
                                                case "hello":
                                                    var v137 int16 = (v76 ^ v76)
                                                    _ = v137
                                                }
                                            }
                                            var v138 bool = true
                                            _ = v138
                                        }
                                        var v139 float32 = 0.0
                                        _ = v102
                                        _ = v139
                                    }
                                default:
                                    switch (^(fi75)) {
                                    case -101:
                                        var v140 uint = 737
                                        _ = v140
                                    case 1:
                                        if true {
                                            var v141 uint = 0
                                            v142 := GenS18{F0: 1, F1: true}
                                            _ = v141
                                            _ = v142
                                        } else {
                                            var v143 float64 = 1e10
                                            var v144 int16 = v76
                                            _ = v143
                                            _ = v144
                                        }
                                    case 4294967295:
                                        fmt.Println(v76, fi84, fi84)
                                    }
                                }
                                _ = fi84
                            }
                            _ = v79
                        }
                    }
                    for fi145 := 0; fi145 < 10; fi145++ {
                        _ = len(v22)
                        p0 = p0
                        fmt.Println(p1, fi75, p1)
                        _ = fi145
                    }
                }
                _ = fi75
                _ = v76
            }
            p0 = p0
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
        }
        for fi146 := 0; fi146 < 1; fi146++ {
            p1 = p1
            v147 := GenS16{F0: (fi146 ^ 9223372036854775807), F1: 0, F2: 40}
            switch genFn0((-(p0))) {
            case 1:
                if (false && (false && true)) {
                    p1 = p1
                    for fi148 := 0; fi148 < 19; fi148++ {
                        var v149 uint64 = 2
                        p0 -= (p0 << ((4294967295) & 63))
                        _ = fi148
                        _ = v149
                    }
                    v150 := GenS19{F0: 83, F1: 2147483647, F2: fi146}
                    _ = v150
                } else {
                    p0 = (^(p0))
                }
            }
            _ = fi146
            _ = v147
        }
    }
    for fi151 := 0; fi151 < 12; fi151++ {
        p0 = (p0 | (p0 ^ -128))
        _ = fi151
    }
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 0
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v5 := []string{("" + ("" + "a"))}
    var v6 string = "foo bar"
    switch -9223372036854775808 {
    case -2147483648:
        v1 -= float64((1e10 + (-2.25 * 1e10)))
    }
    switch -9223372036854775808 {
    case 322:
        var v152 int = 1
        _ = v152
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v153 := GenS19{F0: 0, F1: 42, F2: 309}
    if (-340 != -1) {
        switch "a" {
        case "hello":
            for fi154 := 0; fi154 < 11; fi154++ {
                for fi155 := 0; fi155 < 3; fi155++ {
                    fmt.Println(v6)
                    fmt.Println(v1)
                    v6 = "a"
                    _ = fi155
                }
                _ = fi154
            }
        case "\n\t":
            var v156 int8 = 127
            _ = v156
        }
        switch ("foo bar" + (v6 + "\n\t")) {
        case "":
            v1 += 3.14159
        default:
            v1 += v1
        }
        if false {
            var v157 uint8 = 1
            _ = v157
        }
    }
    var v158 uint16 = 1
    var v159 uint8 = 127
    for fi160 := 0; fi160 < 10; fi160++ {
        fmt.Println(v159)
        v6 = (((("\n\t" + ((("foo bar" + (v6 + "hello")) + v6) + ((((v6 + v6) + v6) + "a") + (v6 + (("hello" + "foo bar") + (v6 + v6)))))) + ((("hello" + v6) + "foo bar") + (v6 + ((v6 + v6) + v6)))) + ("" + ("hello" + ((v6 + v6) + v6)))) + v6)
        for fi161 := 0; fi161 < 18; fi161++ {
            fmt.Println(v158)
            var v162 uint16 = 0
            v6 = "\n\t"
            _ = fi161
            _ = v162
        }
        _ = fi160
    }
    var v163 int16 = -1
    for fi164 := 0; fi164 < 7; fi164++ {
        for fi165 := 0; fi165 < 19; fi165++ {
            var v166 int = fi165
            if ((^(v158)) < 794) {
                fmt.Println(v6)
            }
            _ = fi165
            _ = v166
        }
        _ = fi164
    }
    fmt.Println(v158, v6, v1)
    var v167 float64 = float64(0.1)
    var v168 int32 = -1
    if (!true) {
        fmt.Println(v163)
    } else {
        v6 = (v6 + "")
        var v169 uint8 = v159
        _ = v169
    }
    if true {
        v167 = -2.25
        var v170 uint32 = 65535
        _ = v170
    }
    if true {
        for fi171 := 0; fi171 < 13; fi171++ {
            fmt.Println(v168, v158, v1)
            v172 := GenS17{F0: 0, F1: (((v6 + (((v6 + "hello") + "a") + v6)) + "hello") + v6), F2: 73}
            v163 = (v163 << ((953) & 63))
            _ = fi171
            _ = v172
        }
        v159 += 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v158:", v158)
    fmt.Println("v159:", v159)
    fmt.Println("v163:", v163)
    fmt.Println("v167:", v167)
    fmt.Println("v168:", v168)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v153
    _ = v158
    _ = v159
    _ = v163
    _ = v167
    _ = v168
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
