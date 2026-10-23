package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 string; F1 int; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint }
type GenS6 struct { F0 bool; F1 uint8; F2 int16 }
type GenS7 struct { F0 uint8; F1 uint32; F2 bool }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 uint64; F1 int; F2 int64 }
type GenS10 struct { F0 int8; F1 uint16 }
type GenS11 struct { F0 int16; F1 int16; F2 int32 }
type GenS12 struct { F0 int32; F1 string }
type GenS13 struct { F0 int8; F1 string; F2 uint8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int; F1 int32; F2 uint16 }
type GenS16 struct { F0 uint32; F1 bool }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int32; F1 string }
type GenS19 struct { F0 uint8; F1 int64; F2 string }
func genFn0(p0 uint32, p1 string) string {
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 8; fi5++ {
            p1 = p1
            v6 := GenS1{F0: 1}
            v7 := GenS2{F0: 100, F1: 152}
            if true {
                v8 := GenS1{F0: 0}
                if false {
                    fmt.Println(p0, fi4)
                    v9 := GenS2{F0: 127, F1: 32767}
                    var v10 float64 = 0.0
                    if true {
                        if (!false) {
                            fmt.Println(fi5, p1)
                            v11 := []int{(-(fi4)), (fi4 + fi4), (fi5 % (((-(fi4))) | 1)), fi4}
                            fmt.Println(p1)
                            _ = v11
                        }
                        switch p1 {
                        case "":
                            p1 = ("" + p1)
                        default:
                            switch (fi4 ^ fi5) {
                            case -808:
                                if ((!true) && false) {
                                    v12 := GenS3{F0: p1, F1: fi4, F2: true}
                                    v10 += v10
                                    fmt.Println(fi4, fi5, p0)
                                    _ = v12
                                }
                            }
                        }
                    } else {
                        var v13 uint16 = 255
                        v14 := GenS3{F0: p1, F1: fi4, F2: true}
                        _ = v13
                        _ = v14
                    }
                    v15 := []int{(fi4 >> ((4294967295) & 63)), -1, (fi4 ^ fi4)}
                    _ = v10
                    _ = v15
                    _ = v9
                } else {
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        if true {
                            if (((0.1 * float64((0.1 / (1e10 * 3.14159)))) * 1e10) >= 0.0) {
                                p0 -= (p0 * p0)
                            } else {
                                var v16 string = "hello"
                                _ = v16
                            }
                            switch "\n\t" {
                            case "a":
                                var v17 float32 = (-2.25 * ((-2.25 / 0.1) * 0.1))
                                _ = v17
                            case "foo bar":
                                if (!true) {
                                    switch fi4 {
                                    case 42:
                                        var v18 int8 = -1
                                        p1 = ("hello" + (p1 + ""))
                                        _ = v18
                                    case 9223372036854775807:
                                        v19 := GenS3{F0: "foo bar", F1: (^(fi5)), F2: (false && false)}
                                        _ = v19
                                    default:
                                        v20 := GenS2{F0: 107, F1: 1}
                                        switch "\n\t" {
                                        case "foo bar":
                                            v21 := []int{(fi4 / (((^(fi5))) | 1)), fi5, (fi4 + fi5), fi4}
                                            _ = v21
                                        }
                                        fmt.Println(fi5, fi5)
                                        p0 += 963
                                        v22 := GenS3{F0: ((("hello" + p1) + ((p1 + p1) + p1)) + p1), F1: (^(fi4)), F2: true}
                                        _ = v20
                                        _ = v22
                                    }
                                    p0 += p0
                                    var v23 bool = ((0 < 2147483647) || false)
                                    _ = v23
                                }
                            default:
                                fmt.Println(fi5, fi4)
                            }
                            var v24 int16 = 787
                            _ = v24
                        }
                    case 2147483647:
                        if ((!(false && false)) && true) {
                            p1 = ((p1 + "") + (p1 + ("hello" + p1)))
                            fmt.Println(p1, fi5)
                            if false {
                                p1 = p1
                            }
                        }
                    case 32767:
                        switch "hello" {
                        case "":
                            v25 := GenS1{F0: 1000}
                            _ = v25
                        case "a":
                            fmt.Println(fi4)
                        case "\n\t":
                            var v26 int16 = 1000
                            _ = v26
                        default:
                            if (0 != 0) {
                                if (true || (!false)) {
                                    p0 -= p0
                                } else {
                                    var v27 int32 = -1
                                    var v28 int16 = -1
                                    _ = v27
                                    _ = v28
                                }
                                if true {
                                    fmt.Println(fi5)
                                    if false {
                                        v29 := GenS4{F0: -18}
                                        _ = v29
                                    }
                                    fmt.Println(fi4)
                                }
                                var v30 float64 = 3.14159
                                var v31 uint64 = 2
                                var v32 uint16 = 747
                                p0 = 189
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                        }
                    }
                    v33 := []int{fi5, fi4, fi4}
                    var v34 uint32 = (p0 / (((p0 & 65535)) | 1))
                    _ = v33
                    _ = v34
                }
                v35 := GenS5{F0: 1, F1: 4294967295, F2: 18446744073709551615}
                p0 = (p0 & p0)
                _ = v35
                _ = v8
            }
            switch 32768 {
            case 0:
                if true {
                    p1 = "\n\t"
                }
            }
            v36 := []int{(fi4 << ((42) & 63))}
            _ = fi5
            _ = v36
            _ = v6
            _ = v7
        }
        for fi37 := 0; fi37 < 15; fi37++ {
            v38 := []int{fi4, fi4}
            var v39 int16 = 801
            v38 = append(v38, (fi4 & (fi4 + fi4)))
            _ = fi37
            _ = v38
            _ = v39
        }
        if ((3.14159 - -2.25) >= (((3.14159 * 0.1) + 1e10) * 0.0)) {
            p0 += (p0 + 957)
            if false {
                v40 := GenS5{F0: 1, F1: 18446744073709551615, F2: 42}
                fmt.Println(fi4)
                _ = v40
            } else {
                var v41 int = (fi4 ^ fi4)
                fmt.Println(p0, p0)
                p1 = (("a" + "") + "")
                _ = v41
            }
            if (uint(p0) != 2) {
                p1 = p1
                v42 := GenS2{F0: -128, F1: 1000}
                _ = v42
            } else {
                v43 := []int{(fi4 & (-(fi4))), (fi4 >> ((683) & 63)), 127, (^(fi4))}
                _ = v43
            }
        } else {
            var v44 int32 = -2147483648
            _ = v44
        }
        _ = fi4
    }
    p0 += 824
    for fi45 := 0; fi45 < 8; fi45++ {
        v46 := GenS6{F0: (0 < 237), F1: 0, F2: -9}
        _ = fi45
        _ = v46
    }
    switch 127 {
    case 32767:
        if (true && true) {
            for fi47 := 0; fi47 < 2; fi47++ {
                switch "" {
                case "foo bar":
                    for fi48 := 0; fi48 < 6; fi48++ {
                        if false {
                            if false {
                                if false {
                                    p1 = p1
                                    if true {
                                        var v49 uint = 42
                                        p1 = p1
                                        _ = v49
                                    }
                                }
                                if true {
                                    p0 = p0
                                    if true {
                                        v50 := []int{fi48}
                                        p0 += 2147483647
                                        _ = v50
                                    }
                                }
                            } else {
                                v51 := []int{fi48, (fi47 + fi47), fi48}
                                p1 = ((p1 + "foo bar") + p1)
                                var v52 int32 = -950
                                v53 := GenS2{F0: int8(v52), F1: 6}
                                p0 = (p0 - (p0 | (^(p0))))
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                        } else {
                            p0 = 65535
                            v54 := []int{}
                            p1 = p1
                            _ = v54
                        }
                        v55 := []int{-9223372036854775808, fi48, fi48, fi48}
                        _ = len(v55)
                        _ = fi48
                        _ = v55
                    }
                }
                _ = fi47
            }
            if (888 <= 42) {
                fmt.Println(p0)
            }
            v56 := GenS7{F0: 127, F1: (^(p0)), F2: (-1 != 127)}
            _ = v56
        } else {
            var v57 uint = 4294967295
            _ = v57
        }
    case 275:
        if (!(0 == -9223372036854775808)) {
            v58 := []int{}
            _ = v58
        } else {
            var v59 uint8 = 252
            v60 := GenS7{F0: (v59 & (v59 & (v59 % ((v59) | 1)))), F1: (p0 ^ p0), F2: true}
            _ = v59
            _ = v60
        }
    case -2147483648:
        for fi61 := 0; fi61 < 19; fi61++ {
            v62 := GenS8{F0: -1, F1: 604}
            _ = fi61
            _ = v62
        }
    }
    _ = p0
    _ = p1
    return p1
}
func genFn1() string {
    for fi63 := 0; fi63 < 12; fi63++ {
        switch "\n\t" {
        case "foo bar":
            v64 := []int{}
            _ = v64
        case "\n\t":
            for fi65 := 0; fi65 < 10; fi65++ {
                var v66 int = (fi63 % (((fi65 >> ((2) & 63))) | 1))
                _ = fi65
                _ = v66
            }
        case "":
            var v67 string = ("\n\t" + "")
            _ = v67
        }
        _ = fi63
    }
    var v68 uint8 = 235
    v69 := []int{-337, 1, 608, 32767}
    switch ("\n\t" + (("hello" + (((("" + (("a" + "hello") + ("hello" + (("foo bar" + "foo bar") + "foo bar")))) + "a") + "\n\t") + "\n\t")) + "")) {
    case "\n\t":
        v68 += 0
    case "a":
        for fi70 := 0; fi70 < 7; fi70++ {
            v71 := GenS5{F0: 42, F1: 0, F2: 1}
            _ = fi70
            _ = v71
        }
    case "hello":
        if (true || (false && true)) {
            switch (("a" + "foo bar") + "hello") {
            case "":
                for fi72 := 0; fi72 < 11; fi72++ {
                    switch ((("foo bar" + "") + ("\n\t" + ("a" + ""))) + "") {
                    case "":
                        v68 -= v68
                    case "foo bar":
                        for fi73 := 0; fi73 < 20; fi73++ {
                            if true {
                                switch (fi73 & 0) {
                                case 519:
                                    _ = len(v69)
                                }
                                var v74 int8 = 1
                                _ = v74
                            } else {
                                switch fi73 {
                                case 42:
                                    var v75 float32 = float32((3.14159 - (3.14159 + 0.1)))
                                    _ = v75
                                case 4294967295:
                                    v69 = append(v69, (fi72 / (((fi72 << ((18446744073709551615) & 63))) | 1)))
                                case 9223372036854775807:
                                    switch "foo bar" {
                                    case "\n\t":
                                        fmt.Println(v68, fi73)
                                    case "a":
                                        if false {
                                            _ = len(v69)
                                            if false {
                                                _ = len(v69)
                                                v76 := GenS4{F0: -1}
                                                _ = v76
                                            } else {
                                                var v77 uint64 = 4294967295
                                                v68 = (^(v68))
                                                fmt.Println(v77, v68, fi72)
                                                _ = v77
                                            }
                                            v68 += v68
                                        } else {
                                            var v78 uint32 = 0
                                            switch -1 {
                                            case 1:
                                                switch 2147483647 {
                                                case 2147483647:
                                                    if len(v69) > 0 {
                                                        v69[0] = (fi72 ^ 9223372036854775807)
                                                    }
                                                case 32767:
                                                    var v79 uint = 65535
                                                    _ = v79
                                                default:
                                                    var v80 uint8 = (v68 ^ 161)
                                                    _ = v80
                                                }
                                            case 9223372036854775807:
                                                var v81 int64 = 42
                                                _ = v81
                                            case -2147483648:
                                                fmt.Println(fi73)
                                            }
                                            v82 := GenS2{F0: 0, F1: 0}
                                            if len(v69) >= 2 {
                                                v69 = v69[1:len(v69)]
                                            }
                                            v68 -= (v68 << ((4294967295) & 63))
                                            v68 += 1
                                            if len(v69) > 0 {
                                                v69[0] = (fi73 ^ fi72)
                                            }
                                            _ = v78
                                            _ = v82
                                        }
                                    case "":
                                        var v83 int16 = -32768
                                        if ("foo bar" > "hello") {
                                            var v84 bool = ((("a" + "a") + "") == "foo bar")
                                            switch 32767 {
                                            case 9223372036854775807:
                                                switch -205 {
                                                case 2147483647:
                                                    fmt.Println(fi72)
                                                case 4294967295:
                                                    var v85 uint32 = 248
                                                    _ = v85
                                                }
                                            case -528:
                                                v84 = false
                                            case -949:
                                                if len(v69) >= 2 {
                                                    v69 = v69[1:len(v69)]
                                                }
                                            default:
                                                switch ("" + ("hello" + (("\n\t" + "hello") + "\n\t"))) {
                                                case "":
                                                    v86 := GenS8{F0: -1, F1: 591}
                                                    _ = v86
                                                case "a":
                                                    if len(v69) > 0 {
                                                        v69[0] = (fi73 & fi72)
                                                    }
                                                case "hello":
                                                    var v87 int = -772
                                                    _ = v87
                                                default:
                                                    _ = len(v69)
                                                }
                                            }
                                            _ = v84
                                        } else {
                                            v68 -= 79
                                        }
                                        _ = v83
                                    }
                                }
                                fmt.Println(v68)
                                v68 = v68
                            }
                            _ = fi73
                        }
                    case "a":
                        if false {
                            var v88 float64 = (0.0 + (0.1 - 1e10))
                            var v89 uint16 = 22
                            _ = v88
                            _ = v89
                        }
                    }
                    _ = len(v69)
                    fmt.Println(v68)
                    _ = fi72
                }
            case "hello":
                for fi90 := 0; fi90 < 8; fi90++ {
                    for fi91 := 0; fi91 < 5; fi91++ {
                        v68 -= (^(v68))
                        var v92 int8 = 127
                        switch (fi90 | fi91) {
                        case -2147483648:
                            var v93 float64 = 1.5
                            _ = v93
                        case 128:
                            if false {
                                var v94 float32 = 0.0
                                _ = len(v69)
                                if (false && ((true || false) && (1 == 1))) {
                                    var v95 uint8 = (^(v68))
                                    fmt.Println(v92, fi90, v68)
                                    v96 := GenS2{F0: (^(v92)), F1: 48}
                                    v94 -= v94
                                    _ = v95
                                    _ = v96
                                }
                                _ = v94
                            }
                        case -1:
                            var v97 int8 = -1
                            if (!false) {
                                var v98 uint64 = 0
                                if ((!true) && (v98 >= v98)) {
                                    v97 -= v97
                                    fmt.Println(fi91, fi91)
                                    v98 = v98
                                } else {
                                    if (1 != 32767) {
                                        v99 := GenS8{F0: v92, F1: 4294967295}
                                        if true {
                                            if len(v69) > 0 {
                                                v69[0] = fi90
                                            }
                                        } else {
                                            v68 += v68
                                            fmt.Println(v98)
                                            switch fi90 {
                                            case 0:
                                                if (fi90 == (fi90 + 128)) {
                                                    v68 = (v68 & (v68 >> ((18446744073709551615) & 63)))
                                                } else {
                                                    if (1 == -32768) {
                                                        var v100 int16 = 977
                                                        switch (fi90 - (fi90 / ((9223372036854775807) | 1))) {
                                                        case -774:
                                                            v101 := GenS3{F0: "foo bar", F1: 741, F2: (-9223372036854775808 == -210)}
                                                            _ = v101
                                                        case 9223372036854775807:
                                                            if len(v69) > 0 {
                                                                v69[0] = fi91
                                                            }
                                                        }
                                                        if len(v69) > 0 {
                                                            v69[0] = (fi91 | -9223372036854775808)
                                                        }
                                                        _ = v100
                                                    } else {
                                                        var v102 int64 = -9223372036854775808
                                                        if len(v69) >= 2 {
                                                            v69 = v69[1:len(v69)]
                                                        }
                                                        _ = v102
                                                    }
                                                }
                                            default:
                                                if ((true && true) || ("foo bar" > "")) {
                                                    v103 := GenS5{F0: 995, F1: uint(v98), F2: 89}
                                                    var v104 uint64 = v98
                                                    var v105 int32 = 562
                                                    fmt.Println(fi91)
                                                    _ = v103
                                                    _ = v104
                                                    _ = v105
                                                } else {
                                                    v106 := GenS9{F0: v98, F1: fi90, F2: 0}
                                                    v98 += v98
                                                    v107 := GenS10{F0: 1, F1: uint16(v98)}
                                                    switch (fi91 & -770) {
                                                    case 32767:
                                                        if true {
                                                            var v108 int32 = 2147483647
                                                            _ = v108
                                                        } else {
                                                            fmt.Println(fi90, fi90)
                                                            fmt.Println(v68)
                                                        }
                                                    case 42:
                                                        if ((^(v68)) < (v68 * 0)) {
                                                            var v109 int16 = -702
                                                            v97 = int8(v109)
                                                            if len(v69) > 0 {
                                                                v69[0] = -58
                                                            }
                                                            _ = v109
                                                        } else {
                                                            switch (fi90 / (((fi91 << ((42) & 63))) | 1)) {
                                                            case 940:
                                                                var v110 int32 = -2147483648
                                                                _ = v110
                                                            default:
                                                                fmt.Println(v97, v68)
                                                            }
                                                            v68 = 3
                                                        }
                                                    default:
                                                        if (!(65535 == 581)) {
                                                            var v111 uint32 = 710
                                                            _ = v111
                                                        }
                                                    }
                                                    v97 = v97
                                                    var v112 int16 = 42
                                                    if (!(1.5 == -2.25)) {
                                                        v68 = v68
                                                        v113 := GenS1{F0: int16(v97)}
                                                        switch "hello" {
                                                        case "foo bar":
                                                            if (!false) {
                                                                v114 := GenS7{F0: v68, F1: 65535, F2: false}
                                                                v112 = -92
                                                                _ = v114
                                                            }
                                                        default:
                                                            v115 := GenS5{F0: 481, F1: 0, F2: 4294967295}
                                                            _ = v115
                                                        }
                                                        _ = v113
                                                    } else {
                                                        v98 = v98
                                                        v112 = v112
                                                        var v116 uint = 42
                                                        _ = v116
                                                    }
                                                    fmt.Println(fi91, v68, v112)
                                                    _ = v106
                                                    _ = v107
                                                    _ = v112
                                                }
                                            }
                                        }
                                        v97 += -1
                                        _ = v99
                                    } else {
                                        if (2147483647 != 65535) {
                                            var v117 int = (^(fi91))
                                            if false {
                                                if ((v97 | v97) == v97) {
                                                    var v118 uint = 633
                                                    switch (^(v117)) {
                                                    case 9223372036854775807:
                                                        v69 = append(v69, fi90)
                                                        if len(v69) > 0 {
                                                            v69[0] = v117
                                                        }
                                                        if len(v69) >= 2 {
                                                            v69 = v69[1:len(v69)]
                                                        }
                                                    }
                                                    _ = v118
                                                }
                                                v119 := GenS11{F0: 0, F1: -32768, F2: -1}
                                                _ = v119
                                            }
                                            _ = v117
                                        }
                                    }
                                }
                                _ = len(v69)
                                switch "\n\t" {
                                case "a":
                                    if (65535 > 332) {
                                        if len(v69) >= 2 {
                                            v69 = v69[1:len(v69)]
                                        }
                                        v68 = v68
                                        if (1e10 <= (0.0 + 0.0)) {
                                            switch int(v92) {
                                            case 0:
                                                if false {
                                                    v92 -= 42
                                                }
                                            case 1:
                                                if true {
                                                    var v120 uint32 = 663
                                                    if (1e10 != (0.0 / 1.0)) {
                                                        v121 := GenS1{F0: 32767}
                                                        if len(v69) >= 2 {
                                                            v69 = v69[1:len(v69)]
                                                        }
                                                        if false {
                                                            v69 = append(v69, (-(fi90)))
                                                        }
                                                        _ = v121
                                                    }
                                                    switch fi90 {
                                                    case 128:
                                                        v122 := GenS10{F0: -124, F1: 359}
                                                        _ = v122
                                                    case 32767:
                                                        _ = len(v69)
                                                        v123 := GenS6{F0: true, F1: 42, F2: -83}
                                                        v124 := GenS11{F0: 0, F1: 37, F2: int32(v68)}
                                                        _ = v123
                                                        _ = v124
                                                    }
                                                    _ = v120
                                                }
                                            default:
                                                v68 += 1
                                            }
                                        }
                                    }
                                    v98 = v98
                                    fmt.Println(fi90)
                                    if (-1 <= (-(fi91))) {
                                        v125 := GenS6{F0: ((3.14159 <= 3.14159) && (0 > 0)), F1: (v68 ^ 127), F2: 0}
                                        v68 = v68
                                        v97 -= (v92 >> ((365) & 63))
                                        _ = v125
                                    }
                                }
                                _ = v98
                            } else {
                                var v126 uint32 = 2147483647
                                _ = v126
                            }
                            _ = v97
                        default:
                            v127 := GenS6{F0: false, F1: v68, F2: 42}
                            _ = v127
                        }
                        _ = fi91
                        _ = v92
                    }
                    switch ("foo bar" + ("" + "\n\t")) {
                    case "\n\t":
                        var v128 uint32 = 1
                        _ = v128
                    case "":
                        v129 := GenS2{F0: 42, F1: -32768}
                        _ = v129
                    default:
                        _ = len(v69)
                    }
                    for fi130 := 0; fi130 < 20; fi130++ {
                        if len(v69) >= 2 {
                            v69 = v69[1:len(v69)]
                        }
                        var v131 bool = false
                        _ = fi130
                        _ = v131
                    }
                    _ = fi90
                }
            case "\n\t":
                var v132 uint64 = 2
                _ = v132
            default:
                var v133 float32 = (1e10 + 3.14159)
                _ = v133
            }
            var v134 int64 = -9223372036854775808
            _ = len(v69)
            _ = v134
        } else {
            v68 = v68
        }
    default:
        switch 32767 {
        case -2147483648:
            v68 += v68
        case -474:
            v135 := GenS8{F0: -128, F1: uint32(v68)}
            _ = v135
        }
    }
    if len(v69) > 0 {
        v69[0] = 32768
    }
    _ = v68
    _ = v69
    return "a"
}

func main() {
    var v1 uint64 = 42
    v2 := []int{695, 1, 128}
    var v3 string = "\n\t"
    fmt.Println(v3)
    switch ("\n\t" + "a") {
    case "foo bar":
        for fi136 := 0; fi136 < 7; fi136++ {
            fmt.Println(fi136)
            var v137 uint32 = 388
            var v138 int32 = -634
            _ = fi136
            _ = v137
            _ = v138
        }
    case "\n\t":
        v3 = "hello"
    case "hello":
        switch 127 {
        case 127:
            for fi139 := 0; fi139 < 12; fi139++ {
                if (false && false) {
                    var v140 uint8 = 42
                    var v141 uint8 = (v140 & v140)
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        if true {
                            v141 = (v141 | 249)
                            if (true || false) {
                                if false {
                                    var v142 float64 = -2.25
                                    var v143 int16 = 1
                                    var v144 uint32 = 42
                                    _ = v142
                                    _ = v143
                                    _ = v144
                                }
                            } else {
                                if (("\n\t" + v3) <= ("\n\t" + v3)) {
                                    var v145 int8 = -4
                                    switch v3 {
                                    case "\n\t":
                                        switch fi139 {
                                        case 1:
                                            v146 := GenS12{F0: 380, F1: v3}
                                            _ = v146
                                        case -1:
                                            if true {
                                                fmt.Println(fi139, fi139, v141)
                                                switch -2147483648 {
                                                case -1:
                                                    fmt.Println(v141, v3, v140)
                                                default:
                                                    fmt.Println(v141)
                                                }
                                                v140 += (v140 ^ v140)
                                            }
                                        case -9223372036854775808:
                                            if false {
                                                var v147 bool = false
                                                var v148 int16 = 1
                                                fmt.Println(v140)
                                                _ = v147
                                                _ = v148
                                            } else {
                                                var v149 int64 = -1000
                                                if true {
                                                    fmt.Println(v141, v145, fi139)
                                                } else {
                                                    var v150 float32 = 1e10
                                                    _ = v150
                                                }
                                                _ = v149
                                            }
                                        default:
                                            v1 += v1
                                        }
                                    }
                                    _ = v145
                                } else {
                                    if true {
                                        for fi151 := 0; fi151 < 16; fi151++ {
                                            if true {
                                                if true {
                                                    fmt.Println(v3, v141, fi151)
                                                    var v152 uint = 4294967295
                                                    if false {
                                                        v3 = v3
                                                    } else {
                                                        var v153 int = (fi151 + fi151)
                                                        v154 := GenS13{F0: 127, F1: ((v3 + "") + "\n\t"), F2: (^(v140))}
                                                        switch 4294967295 {
                                                        case 791:
                                                            v140 = (v141 + (v140 % ((v141) | 1)))
                                                        }
                                                        _ = v153
                                                        _ = v154
                                                    }
                                                    _ = v152
                                                }
                                                var v155 float32 = (3.14159 + -2.25)
                                                _ = v155
                                            }
                                            if true {
                                                if len(v2) > 0 {
                                                    v2[0] = (fi139 % (((-(fi139))) | 1))
                                                }
                                                fmt.Println(fi151)
                                            } else {
                                                switch v3 {
                                                case "hello":
                                                    v141 -= v140
                                                }
                                                var v156 int32 = 497
                                                switch fi151 {
                                                case 0:
                                                    var v157 int32 = 2147483647
                                                    _ = v157
                                                case 32767:
                                                    var v158 float64 = ((0.1 - ((1e10 * 1.5) * (1e10 - (-2.25 - 0.1)))) * 0.0)
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        if (v3 < v3) {
                                                            var v159 uint64 = v1
                                                            var v160 bool = false
                                                            var v161 uint32 = 2147483648
                                                            _ = v159
                                                            _ = v160
                                                            _ = v161
                                                        } else {
                                                            if (1 >= 4294967295) {
                                                                switch (^(fi139)) {
                                                                case 127:
                                                                    v140 -= 210
                                                                case 128:
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                case -2147483648:
                                                                    var v162 int16 = 0
                                                                    _ = v162
                                                                }
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                            v156 += 0
                                                            v163 := GenS14{F0: 2}
                                                            switch (^(fi151)) {
                                                            case 1:
                                                                v164 := GenS9{F0: 1, F1: (-(fi139)), F2: 1}
                                                                _ = v164
                                                            case 32767:
                                                                var v165 string = ((v3 + v3) + "hello")
                                                                _ = v165
                                                            }
                                                            if false {
                                                                fmt.Println(fi139)
                                                                if (2 < 2147483648) {
                                                                    switch -753 {
                                                                    case -702:
                                                                        fmt.Println(v1)
                                                                    case -9223372036854775808:
                                                                        switch ((v3 + "a") + (v3 + v3)) {
                                                                        case "\n\t":
                                                                            var v166 uint = 4294967295
                                                                            _ = v166
                                                                        case "a":
                                                                            v167 := GenS15{F0: fi151, F1: (v156 * v156), F2: 702}
                                                                            _ = v167
                                                                        case "":
                                                                            switch -561 {
                                                                            case -2147483648:
                                                                                var v168 uint64 = (v1 << ((240) & 63))
                                                                                _ = v168
                                                                            }
                                                                        default:
                                                                            v1 -= v1
                                                                        }
                                                                    default:
                                                                        var v169 string = (v3 + v3)
                                                                        _ = v169
                                                                    }
                                                                    switch (("\n\t" + "\n\t") + v3) {
                                                                    case "foo bar":
                                                                        if (true && (1 < -9223372036854775808)) {
                                                                            v1 += (v1 | 42)
                                                                        }
                                                                    case "hello":
                                                                        var v170 bool = (int32(fi139) >= (v156 + -1))
                                                                        _ = v170
                                                                    case "\n\t":
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 2147483647
                                                                        }
                                                                    }
                                                                } else {
                                                                    _ = len(v2)
                                                                    switch v3 {
                                                                    case "hello":
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    case "foo bar":
                                                                        var v171 int64 = -1
                                                                        v158 = v158
                                                                        _ = v171
                                                                    case "":
                                                                        v156 += v156
                                                                    }
                                                                    var v172 uint8 = v140
                                                                    _ = v172
                                                                }
                                                                var v173 int16 = 1000
                                                                _ = v173
                                                            }
                                                            v174 := GenS6{F0: true, F1: (v141 - 186), F2: -1}
                                                            v175 := GenS7{F0: (v141 * v140), F1: 65535, F2: false}
                                                            _ = v163
                                                            _ = v174
                                                            _ = v175
                                                        }
                                                    case "foo bar":
                                                        switch 9223372036854775807 {
                                                        case -460:
                                                            switch fi151 {
                                                            case 42:
                                                                v141 += 2
                                                            default:
                                                                var v176 float32 = 1e10
                                                                _ = v176
                                                            }
                                                        case -304:
                                                            var v177 uint32 = 4294967295
                                                            _ = v177
                                                        default:
                                                            if (42 < 887) {
                                                                if false {
                                                                    if false {
                                                                        v141 += (^(v140))
                                                                        if true {
                                                                            if true {
                                                                                v158 -= (-2.25 + ((0.0 - 3.14159) * ((v158 * (v158 * v158)) / 1.0)))
                                                                                var v178 int16 = -734
                                                                                var v179 uint64 = 354
                                                                                _ = v178
                                                                                _ = v179
                                                                            }
                                                                            v156 = 42
                                                                            if len(v2) > 0 {
                                                                                v2[0] = fi151
                                                                            }
                                                                            if len(v2) > 0 {
                                                                                v2[0] = fi151
                                                                            }
                                                                            var v180 uint64 = v1
                                                                            _ = v180
                                                                        }
                                                                        if true {
                                                                            v1 = v1
                                                                        } else {
                                                                            switch fi139 {
                                                                            case -2147483648:
                                                                                var v181 int8 = -1
                                                                                _ = v181
                                                                            }
                                                                            var v182 int32 = (v156 | 42)
                                                                            _ = v182
                                                                        }
                                                                    }
                                                                    var v183 int32 = (v156 ^ (v156 ^ (v156 ^ (v156 * v156))))
                                                                    v183 -= 42
                                                                    _ = v183
                                                                }
                                                            }
                                                        }
                                                        switch (v3 + v3) {
                                                        case "foo bar":
                                                            var v184 int64 = -9223372036854775808
                                                            _ = v184
                                                        case "hello":
                                                            v141 = (v141 >> ((244) & 63))
                                                            if len(v2) > 0 {
                                                                v2[0] = fi151
                                                            }
                                                        }
                                                        _ = len(v2)
                                                    case "":
                                                        v158 = (-2.25 * (float64(-2.25) / 3.14159))
                                                    default:
                                                        v2 = append(v2, fi139)
                                                    }
                                                    v3 = ""
                                                    _ = v158
                                                case -209:
                                                    if false {
                                                        if (true || true) {
                                                            v156 -= v156
                                                            var v185 string = "hello"
                                                            _ = v185
                                                        }
                                                        var v186 int8 = -38
                                                        v187 := GenS16{F0: 42, F1: true}
                                                        _ = v186
                                                        _ = v187
                                                    }
                                                default:
                                                    v3 = ("" + v3)
                                                }
                                                _ = v156
                                            }
                                            _ = fi151
                                        }
                                    }
                                    var v188 int64 = 0
                                    switch (v3 + ("foo bar" + v3)) {
                                    case "":
                                        v189 := GenS9{F0: (v1 % (((^(v1))) | 1)), F1: fi139, F2: int64(v1)}
                                        _ = v189
                                    }
                                    _ = v188
                                }
                            }
                            var v190 int32 = 195
                            _ = v190
                        } else {
                            var v191 int64 = -1
                            for fi192 := 0; fi192 < 10; fi192++ {
                                v2 = append(v2, (fi192 | (fi192 ^ 0)))
                                _ = fi192
                            }
                            _ = v191
                        }
                    default:
                        v2 = append(v2, (fi139 ^ (-(fi139))))
                    }
                    _ = v140
                    _ = v141
                } else {
                    if false {
                        if (false || true) {
                            v1 = (v1 % (((^(v1))) | 1))
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            for fi193 := 0; fi193 < 4; fi193++ {
                                v194 := GenS17{F0: (((v3 + "a") + ((v3 + v3) + "\n\t")) + ((v3 + v3) + "foo bar"))}
                                _ = fi193
                                _ = v194
                            }
                        }
                        v1 = v1
                        var v195 uint16 = 36
                        _ = v195
                    }
                }
                _ = fi139
            }
        case -83:
            if (42 <= 65535) {
                if true {
                    v2 = append(v2, 42)
                    v2 = append(v2, 4294967295)
                } else {
                    fmt.Println(v1)
                    fmt.Println(v3, v1)
                    v196 := GenS4{F0: 1}
                    _ = v196
                }
                if false {
                    v197 := GenS18{F0: -219, F1: (v3 + v3)}
                    _ = v197
                } else {
                    v198 := GenS4{F0: -124}
                    for fi199 := 0; fi199 < 14; fi199++ {
                        v200 := GenS19{F0: 255, F1: -1, F2: "hello"}
                        _ = fi199
                        _ = v200
                    }
                    var v201 int32 = -2147483648
                    _ = v198
                    _ = v201
                }
            }
        default:
            if true {
                var v202 float64 = (3.14159 / -2.25)
                if len(v2) > 0 {
                    v2[0] = -201
                }
                if (((!((true && (!(-29 <= -2147483648))) && (2147483647 > 2147483647))) || false) && (!(-32768 > 550))) {
                    v202 = v202
                } else {
                    v202 = ((0.1 * v202) * ((1.5 - (3.14159 * v202)) + v202))
                }
                _ = v202
            }
        }
    }
    var v203 uint8 = 42
    var v204 int8 = 100
    switch -9223372036854775808 {
    case 884:
        var v205 uint8 = (^(v203))
        _ = v205
    case 345:
        v204 = v204
    }
    for fi206 := 0; fi206 < 13; fi206++ {
        var v207 int16 = 1
        _ = fi206
        _ = v207
    }
    if true {
        v3 = v3
        v1 = (v1 << ((0) & 63))
    }
    if true {
        for fi208 := 0; fi208 < 8; fi208++ {
            fmt.Println(v1, v203)
            _ = fi208
        }
        if (-933 >= -32768) {
            if (0 == 42) {
                for fi209 := 0; fi209 < 17; fi209++ {
                    for fi210 := 0; fi210 < 7; fi210++ {
                        var v211 float64 = (3.14159 + 1.5)
                        var v212 string = "hello"
                        if ((float32((((v211 - (v211 - v211)) * 0.0) + -2.25)) * 1e10) >= 0.1) {
                            var v213 uint16 = 42
                            _ = v213
                        }
                        _ = fi210
                        _ = v211
                        _ = v212
                    }
                    _ = fi209
                }
            }
            var v214 float32 = ((-2.25 + 3.14159) - ((float32(1e10) - 0.0) / (0.1 - (-2.25 + -2.25))))
            _ = v214
        }
    }
    for fi215 := 0; fi215 < 10; fi215++ {
        var v216 uint8 = v203
        _ = fi215
        _ = v216
    }
    v217 := GenS7{F0: (v203 * v203), F1: 2147483647, F2: false}
    var v218 float64 = float64(1e10)
    v204 += v204
    var v219 float64 = 1e10
    v220 := GenS4{F0: v204}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v203:", v203)
    fmt.Println("v204:", v204)
    fmt.Println("v218:", v218)
    fmt.Println("v219:", v219)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v203
    _ = v204
    _ = v217
    _ = v218
    _ = v219
    _ = v220
    _ = v3
}
