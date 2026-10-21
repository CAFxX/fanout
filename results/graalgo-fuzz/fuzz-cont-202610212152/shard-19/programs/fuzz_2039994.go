package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint }
type GenS4 struct { F0 int64; F1 int64; F2 string }
type GenS5 struct { F0 int64; F1 string; F2 uint64 }
type GenS6 struct { F0 int8; F1 int32; F2 uint32 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint; F1 int64 }
type GenS11 struct { F0 int16; F1 string }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int16; F1 string; F2 int8 }
type GenS15 struct { F0 string; F1 uint; F2 int32 }
type GenS16 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS17 struct { F0 uint16; F1 uint16 }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 int8; F1 uint8 }
type GenS20 struct { F0 int; F1 int }
type GenS21 struct { F0 bool; F1 uint16 }
type GenS22 struct { F0 uint8 }
type GenS23 struct { F0 int64; F1 uint32; F2 int }
type GenS24 struct { F0 uint8 }
type GenS25 struct { F0 uint32 }
type GenS26 struct { F0 int64 }
type GenS27 struct { F0 uint8; F1 bool }
type GenS28 struct { F0 int8; F1 int32 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]uint32)
    var v3 uint = 866
    v4 := make(map[string]uint32)
    v5 := []string{"foo bar", "\n\t", ("foo bar" + "foo bar")}
    v6 := make(map[uint32]string)
    if false {
        if (32767 < 461) {
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 int = (fi7 - -2147483648)
                var v9 int8 = 127
                var v10 string = ((("" + "\n\t") + "foo bar") + ((("foo bar" + "a") + "\n\t") + ("hello" + ("" + ("hello" + "a")))))
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            for fi11 := 0; fi11 < 16; fi11++ {
                switch fi11 {
                case 32768:
                    switch "hello" {
                    case "\n\t":
                        v1 = v1
                    case "a":
                        var v12 int32 = 1
                        _ = v12
                    }
                case 128:
                    switch (fi11 * fi11) {
                    case 0:
                        var v13 uint64 = 0
                        _ = v13
                    default:
                        switch fi11 {
                        case 127:
                            v3 += uint(v1)
                        case 2147483647:
                            if ((v3 + (v3 >> ((v3) & 63))) > (v3 << ((uint(v1)) & 63))) {
                                for fi14 := 0; fi14 < 19; fi14++ {
                                    v15 := GenS1{F0: 1, F1: 1}
                                    _ = fi14
                                    _ = v15
                                }
                                switch (-(fi11)) {
                                case 42:
                                    var v16 uint64 = 2
                                    _ = v16
                                case -295:
                                    if ((3.14159 < 1.5) && (-32768 <= -32768)) {
                                        switch ("foo bar" + "") {
                                        case "hello":
                                            if false {
                                                var v17 uint16 = 255
                                                var v18 int = (^(fi11))
                                                _ = v17
                                                _ = v18
                                            } else {
                                                switch (fi11 << ((707) & 63)) {
                                                case -234:
                                                    v3 = v3
                                                case 540:
                                                    var v19 float32 = float32(1e10)
                                                    _ = v19
                                                case -332:
                                                    var v20 uint64 = 2
                                                    _ = v20
                                                default:
                                                    var v21 int32 = 798
                                                    _ = v21
                                                }
                                            }
                                        case "foo bar":
                                            v1 = 0
                                        case "a":
                                            if (42 > 100) {
                                                var v22 float32 = float32(3.14159)
                                                _ = v22
                                            }
                                        default:
                                            for fi23 := 0; fi23 < 3; fi23++ {
                                                var v24 float64 = 1e10
                                                v5 = append(v5, ("hello" + "hello"))
                                                _ = fi23
                                                _ = v24
                                            }
                                        }
                                        var v25 int = (fi11 & (fi11 - fi11))
                                        _ = v25
                                    } else {
                                        _ = len(v6)
                                        v2[uint32(v3)] = 65535
                                        var v26 int = 0
                                        _ = v26
                                    }
                                }
                                if (42 <= 365) {
                                    v3 += 65535
                                }
                            } else {
                                if (fi11 >= (fi11 + (fi11 - (fi11 * (fi11 * (^(fi11))))))) {
                                    v27 := GenS1{F0: 26, F1: -1}
                                    v1 += 65535
                                    _ = v27
                                }
                                v3 = v3
                            }
                        case 1:
                            fmt.Println(v3)
                        }
                    }
                default:
                    var v28 float64 = 1.5
                    _ = v28
                }
                var v29 int64 = 1
                switch fi11 {
                case -177:
                    v29 = 1
                case 220:
                    v4["\n\t"] = 65535
                case -1:
                    switch "" {
                    case "\n\t":
                        v4["\n\t"] = 4294967295
                    }
                default:
                    fmt.Println(fi11, v1)
                }
                _ = fi11
                _ = v29
            }
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 uint8 = 127
                v1 -= (^(v1))
                _ = fi30
                _ = v31
            }
        }
        v1 = v1
    } else {
        v5 = append(v5, ("" + "\n\t"))
        if (true && true) {
            v1 -= v1
            delete(v4, "\n\t")
            if false {
                for fi32 := 0; fi32 < 2; fi32++ {
                    switch ("foo bar" + (((("hello" + "hello") + ((((("foo bar" + "a") + "a") + (("foo bar" + "foo bar") + ("" + ""))) + "a") + "a")) + "\n\t") + ("a" + ("\n\t" + "")))) {
                    case "":
                        if (!(false || (160 == 44))) {
                            switch (^(fi32)) {
                            case 310:
                                v33 := v4["\n\t"]
                                _ = v33
                            case -282:
                                switch ("a" + ("" + "hello")) {
                                case "\n\t":
                                    v5 = append(v5, ("" + "hello"))
                                case "a":
                                    switch ("" + ((("" + ("\n\t" + (("foo bar" + "") + ("hello" + "hello")))) + "\n\t") + "")) {
                                    case "":
                                        var v34 int16 = -819
                                        _ = v34
                                    case "a":
                                        for fi35 := 0; fi35 < 18; fi35++ {
                                            switch "foo bar" {
                                            case "":
                                                switch (-(fi35)) {
                                                case 128:
                                                    v36 := GenS1{F0: 0, F1: 42}
                                                    _ = v36
                                                case 4294967295:
                                                    var v37 int64 = 9223372036854775807
                                                    _ = v37
                                                default:
                                                    switch (fi32 & 710) {
                                                    case -502:
                                                        v38 := GenS1{F0: 21, F1: 32767}
                                                        _ = v38
                                                    case 127:
                                                        if (-9223372036854775808 >= 9223372036854775807) {
                                                            fmt.Println(fi32, fi35)
                                                        }
                                                    case -532:
                                                        v39 := GenS2{F0: 1}
                                                        _ = v39
                                                    default:
                                                        v40 := GenS1{F0: uint8(v3), F1: 1000}
                                                        _ = v40
                                                    }
                                                }
                                            }
                                            fmt.Println(v1, v1)
                                            _ = fi35
                                        }
                                    }
                                case "hello":
                                    v1 -= v1
                                }
                            }
                        }
                    case "hello":
                        v5 = append(v5, ("foo bar" + ""))
                    default:
                        fmt.Println(fi32)
                    }
                    v5 = append(v5, ("\n\t" + "a"))
                    switch 843 {
                    case -9223372036854775808:
                        for fi41 := 0; fi41 < 7; fi41++ {
                            if len(v5) > 0 {
                                v5[0] = ("a" + "\n\t")
                            }
                            if false {
                                if (0 <= -61) {
                                    v1 = v1
                                    var v42 int64 = -1
                                    if ("foo bar" > "\n\t") {
                                        v42 -= (v42 | v42)
                                        v43 := GenS1{F0: 255, F1: 1000}
                                        var v44 uint64 = 891
                                        _ = v43
                                        _ = v44
                                    }
                                    fmt.Println(fi41, v3, v1)
                                    v42 += 0
                                    v4[("foo bar" + "")] = (v1 << (((v3 | 2)) & 63))
                                    _ = len(v5)
                                    v1 = v1
                                    v1 = v1
                                    v1 = v1
                                    _ = v42
                                } else {
                                    v3 = v3
                                }
                                if true {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                } else {
                                    if true {
                                        if false {
                                            v3 -= 4294967295
                                        }
                                        switch ((("a" + "") + "foo bar") + (("foo bar" + ("" + "\n\t")) + "hello")) {
                                        case "hello":
                                            switch ("\n\t" + ("a" + "")) {
                                            case "":
                                                var v45 uint64 = 199
                                                _ = v45
                                            case "a":
                                                if len(v5) > 0 {
                                                    v5[0] = "hello"
                                                }
                                            }
                                        case "\n\t":
                                            if false {
                                                var v46 uint8 = 128
                                                v46 = (v46 | (v46 ^ (v46 % (((v46 >> ((4294967295) & 63))) | 1))))
                                                v3 = 18446744073709551615
                                                _ = v46
                                            } else {
                                                fmt.Println(fi41, v1)
                                            }
                                        }
                                        var v47 uint16 = 65535
                                        _ = v47
                                    }
                                    v48 := GenS2{F0: 0}
                                    v3 = 848
                                    _ = v48
                                }
                            }
                            if (!(int8(fi32) >= 0)) {
                                var v49 int8 = -1
                                _ = v49
                            } else {
                                switch "a" {
                                case "\n\t":
                                    v50 := GenS3{F0: -9223372036854775808, F1: 802, F2: v3}
                                    _ = v50
                                case "a":
                                    var v51 float32 = -2.25
                                    _ = v51
                                default:
                                    if false {
                                        switch "hello" {
                                        case "foo bar":
                                            v5 = append(v5, "")
                                            if true {
                                                if false {
                                                    var v52 float32 = ((0.0 / 0.1) - 1.5)
                                                    _ = v52
                                                } else {
                                                    if (true || (!false)) {
                                                        v53 := GenS3{F0: 42, F1: 1, F2: v3}
                                                        _ = v53
                                                    }
                                                }
                                            } else {
                                                var v54 int = (^(fi41))
                                                _ = v54
                                            }
                                            var v55 bool = true
                                            _ = v55
                                        case "\n\t":
                                            if false {
                                                switch 9223372036854775807 {
                                                case 4294967295:
                                                    switch fi32 {
                                                    case 127:
                                                        v1 = v1
                                                        var v56 int64 = -9223372036854775808
                                                        if ((v3 ^ v3) <= (^(v3))) {
                                                            fmt.Println(fi41)
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                fmt.Println(v3)
                                                            }
                                                        }
                                                        _ = v56
                                                    case 515:
                                                        v3 -= v3
                                                    }
                                                case 665:
                                                    fmt.Println(fi41)
                                                case 32767:
                                                    v3 = v3
                                                }
                                            }
                                        case "":
                                            var v57 uint8 = 2
                                            _ = v57
                                        default:
                                            v1 = 2147483648
                                        }
                                        v3 -= (v3 + v3)
                                        v58 := GenS1{F0: 2, F1: 0}
                                        _ = v58
                                    } else {
                                        var v59 uint8 = 1
                                        v59 -= v59
                                        delete(v4, ("foo bar" + ("a" + ("a" + "\n\t"))))
                                        _ = v59
                                    }
                                }
                                var v60 int64 = 0
                                _ = v60
                            }
                            _ = fi41
                        }
                    case -2147483648:
                        for fi61 := 0; fi61 < 18; fi61++ {
                            v3 = v3
                            _ = fi61
                        }
                    }
                    _ = fi32
                }
            } else {
                v1 -= 755
                v62 := GenS2{F0: 42}
                v63 := GenS4{F0: 42, F1: 1, F2: "foo bar"}
                _ = v62
                _ = v63
            }
        }
        var v64 float32 = (((((0.0 - -2.25) - (1.5 + 0.0)) * ((3.14159 + (3.14159 - -2.25)) / -2.25)) + 3.14159) - (1e10 * (((0.0 + float32(3.14159)) + -2.25) + ((0.0 / ((float32(0.1) * 0.0) + 1.5)) / (0.1 * float32(1e10))))))
        _ = v64
    }
    if true {
        var v65 uint16 = 42
        v66 := GenS5{F0: 9223372036854775807, F1: "\n\t", F2: 878}
        var v67 float64 = 3.14159
        _ = v65
        _ = v66
        _ = v67
    } else {
        v68 := GenS1{F0: 2, F1: 17}
        if false {
            for fi69 := 0; fi69 < 19; fi69++ {
                var v70 int8 = 75
                if (128 != 42) {
                    fmt.Println(v1, v70, v3)
                } else {
                    var v71 bool = true
                    var v72 float64 = 0.1
                    _ = v71
                    _ = v72
                }
                _ = fi69
                _ = v70
            }
            var v73 string = "foo bar"
            _ = v73
        }
        v1 = (^(v1))
        _ = v68
    }
    var v74 uint16 = 2
    if false {
        switch 42 {
        case 127:
            v74 = v74
        case 32767:
            for fi75 := 0; fi75 < 15; fi75++ {
                switch (-(fi75)) {
                case -577:
                    fmt.Println(v74)
                case -484:
                    var v76 uint = 625
                    _ = v76
                case -2147483648:
                    switch ("" + ("\n\t" + "foo bar")) {
                    case "hello":
                        if (!false) {
                            var v77 int8 = 1
                            v1 += v1
                            v77 += 42
                            _ = v77
                        } else {
                            var v78 uint = 0
                            fmt.Println(v1, v78)
                            var v79 int8 = 100
                            _ = v78
                            _ = v79
                        }
                    case "a":
                        for fi80 := 0; fi80 < 18; fi80++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi80
                        }
                    default:
                        for fi81 := 0; fi81 < 14; fi81++ {
                            fmt.Println(v3, fi75, fi81)
                            var v82 uint = v3
                            _ = fi81
                            _ = v82
                        }
                    }
                default:
                    if true {
                        if (!(180 > 83)) {
                            v83 := GenS5{F0: 555, F1: "hello", F2: 4294967295}
                            if false {
                                for fi84 := 0; fi84 < 2; fi84++ {
                                    switch (fi75 ^ (fi75 >> ((v3) & 63))) {
                                    case 1:
                                        v85 := GenS3{F0: -9223372036854775808, F1: 2, F2: v3}
                                        _ = v85
                                    case 32768:
                                        if (v1 >= v1) {
                                            v6[v1] = ("" + "hello")
                                        } else {
                                            v3 += v3
                                            var v86 int16 = -32768
                                            v87 := GenS6{F0: 100, F1: 0, F2: (v1 ^ (v1 * v1))}
                                            _ = v86
                                            _ = v87
                                        }
                                    case -978:
                                        if true {
                                            if false {
                                                v74 -= (v74 * v74)
                                                fmt.Println(v1, fi84)
                                            }
                                            v4["\n\t"] = (v1 * 406)
                                        }
                                    }
                                    fmt.Println(fi84, v1)
                                    fmt.Println(fi75)
                                    _ = fi84
                                }
                                fmt.Println(v1)
                            } else {
                                var v88 bool = true
                                if (0 < -671) {
                                    var v89 uint = 761
                                    _ = len(v2)
                                    var v90 uint = (v3 << ((255) & 63))
                                    _ = v89
                                    _ = v90
                                } else {
                                    v3 = (^(v3))
                                    var v91 int32 = 423
                                    if len(v5) > 0 {
                                        v5[0] = "foo bar"
                                    }
                                    _ = v91
                                }
                                _ = v88
                            }
                            if ((^(v74)) <= v74) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            _ = v83
                        } else {
                            if false {
                                var v92 int32 = -545
                                for fi93 := 0; fi93 < 6; fi93++ {
                                    _ = len(v5)
                                    v92 = (v92 ^ v92)
                                    _ = fi93
                                }
                                _ = v92
                            }
                        }
                        switch "foo bar" {
                        case "a":
                            var v94 int64 = 0
                            _ = v94
                        case "\n\t":
                            if (42 > 644) {
                                var v95 int8 = 15
                                v96 := GenS7{F0: 1, F1: 4294967295}
                                _ = v95
                                _ = v96
                            }
                        case "foo bar":
                            var v97 int32 = 1
                            _ = v97
                        default:
                            v1 -= uint32(v74)
                        }
                    }
                }
                if (1.5 < float64(-2.25)) {
                    for fi98 := 0; fi98 < 19; fi98++ {
                        if (false && (1.5 <= (1e10 - 3.14159))) {
                            var v99 float64 = -2.25
                            v99 -= v99
                            _ = v99
                        } else {
                            var v100 uint16 = (v74 + (v74 >> (((^(v3))) & 63)))
                            if false {
                                switch fi75 {
                                case 2147483647:
                                    fmt.Println(v74, v74, v1)
                                }
                                if true {
                                    v101 := GenS8{F0: 1, F1: (v100 & (v74 % ((65535) | 1)))}
                                    _ = v101
                                } else {
                                    v100 -= (v74 + (v74 - (v100 - v100)))
                                    fmt.Println(v3)
                                    switch -9223372036854775808 {
                                    case -2147483648:
                                        if true {
                                            v102 := GenS6{F0: 127, F1: -2147483648, F2: (^(v1))}
                                            v103 := GenS7{F0: 2, F1: 18446744073709551615}
                                            _ = v102
                                            _ = v103
                                        } else {
                                            fmt.Println(fi98, fi98, v100)
                                            fmt.Println(v100, v100, fi98)
                                            v3 -= v3
                                            switch fi98 {
                                            case -2147483648:
                                                v3 -= v3
                                            }
                                        }
                                    case 1:
                                        v104 := GenS2{F0: uint8(v1)}
                                        _ = v104
                                    default:
                                        v74 = (^(v100))
                                    }
                                    var v105 uint16 = (v100 << ((v3) & 63))
                                    var v106 uint64 = 83
                                    _ = v105
                                    _ = v106
                                }
                                delete(v6, v1)
                            } else {
                                switch -1 {
                                case 4294967295:
                                    if false {
                                        delete(v4, "")
                                        v107 := v2[188]
                                        _ = v107
                                    } else {
                                        v74 += 498
                                        switch "" {
                                        case "hello":
                                            v100 = 546
                                        }
                                    }
                                case 2147483647:
                                    v108 := GenS9{F0: (-2.25 < (((3.14159 + 1e10) - 0.0) / 1.0))}
                                    _ = v108
                                default:
                                    v3 += (v3 / (((v3 & (v3 / ((v3) | 1)))) | 1))
                                }
                            }
                            _ = v100
                        }
                        _ = fi98
                    }
                    var v109 string = ((("foo bar" + ((("foo bar" + "a") + "a") + ("" + "a"))) + ("" + (("hello" + "") + "\n\t"))) + ("foo bar" + ("" + "a")))
                    var v110 float64 = 0.0
                    _ = v109
                    _ = v110
                }
                _ = fi75
            }
        default:
            v111 := GenS10{F0: uint(v1), F1: 542}
            _ = v111
        }
        var v112 float64 = 0.0
        v113 := GenS10{F0: uint(v74), F1: -9223372036854775808}
        _ = v112
        _ = v113
    } else {
        v3 += v3
        fmt.Println(v74, v3, v1)
    }
    if true {
        var v114 int16 = 42
        _ = v114
    } else {
        v115 := GenS2{F0: 179}
        var v116 uint32 = uint32(v74)
        _ = v115
        _ = v116
    }
    if (!true) {
        fmt.Println(v1, v74, v3)
        for fi117 := 0; fi117 < 11; fi117++ {
            for fi118 := 0; fi118 < 2; fi118++ {
                _ = len(v5)
                _ = fi118
            }
            _ = fi117
        }
        for fi119 := 0; fi119 < 17; fi119++ {
            for fi120 := 0; fi120 < 9; fi120++ {
                v121 := GenS7{F0: 0, F1: 1}
                var v122 int8 = -1
                v3 -= v3
                v1 = (v1 ^ 0)
                v123 := v4["foo bar"]
                _ = fi120
                _ = v121
                _ = v122
                _ = v123
            }
            for fi124 := 0; fi124 < 18; fi124++ {
                fmt.Println(fi119, fi124)
                _ = fi124
            }
            _ = fi119
        }
    } else {
        v1 = (^(v1))
        for fi125 := 0; fi125 < 20; fi125++ {
            for fi126 := 0; fi126 < 7; fi126++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v127 := GenS1{F0: 255, F1: 1000}
                _ = fi126
                _ = v127
            }
            if true {
                v4["foo bar"] = (^(v1))
            }
            _ = fi125
        }
        for fi128 := 0; fi128 < 19; fi128++ {
            for fi129 := 0; fi129 < 10; fi129++ {
                v130 := GenS9{F0: false}
                _ = fi129
                _ = v130
            }
            v74 -= (v74 | 19)
            _ = fi128
        }
    }
    for fi131 := 0; fi131 < 1; fi131++ {
        switch "\n\t" {
        case "\n\t":
            v3 = v3
        case "":
            v132 := GenS1{F0: uint8(v74), F1: 0}
            _ = v132
        case "foo bar":
            v5 = append(v5, "")
        }
        var v133 uint64 = 42
        v1 = 4294967295
        _ = fi131
        _ = v133
    }
    v134 := GenS4{F0: 537, F1: 1, F2: ("foo bar" + "")}
    for fi135 := 0; fi135 < 20; fi135++ {
        v136 := GenS2{F0: 128}
        if false {
            switch "\n\t" {
            case "hello":
                switch "hello" {
                case "":
                    v74 = uint16(v3)
                case "\n\t":
                    switch fi135 {
                    case 19:
                        v137 := GenS9{F0: (0 > -1)}
                        _ = v137
                    default:
                        if ((!true) && false) {
                            v74 -= (v74 ^ v74)
                            v1 -= (^(v1))
                        }
                    }
                }
            }
            for fi138 := 0; fi138 < 18; fi138++ {
                if true {
                    var v139 string = ("foo bar" + "")
                    var v140 int64 = -456
                    fmt.Println(v140, fi135)
                    _ = v139
                    _ = v140
                } else {
                    if (true || true) {
                        if false {
                            var v141 bool = false
                            var v142 string = ""
                            v142 = v142
                            _ = v141
                            _ = v142
                        } else {
                            if false {
                                switch fi135 {
                                case 42:
                                    v1 -= 361
                                case -6:
                                    fmt.Println(v1, v1, v74)
                                default:
                                    v143 := GenS2{F0: 127}
                                    _ = v143
                                }
                                if false {
                                    v3 = v3
                                    if (!false) {
                                        if ((!(v74 < v74)) || true) {
                                            v1 = 348
                                            v74 += 578
                                        } else {
                                            if (!(false && (!true))) {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v1 = v1
                                                switch fi138 {
                                                case -746:
                                                    if (true && ((-96 < -1) && false)) {
                                                        v3 -= (^(v3))
                                                        v2[473] = v1
                                                        var v144 uint16 = (v74 - v74)
                                                        if ((v3 >> ((272) & 63)) > v3) {
                                                            switch "" {
                                                            case "\n\t":
                                                                if (1e10 == float32(0.0)) {
                                                                    if (199 == 0) {
                                                                        v145 := GenS11{F0: 1, F1: "\n\t"}
                                                                        var v146 string = "a"
                                                                        if true {
                                                                            v1 -= 2147483648
                                                                            v1 = v1
                                                                        } else {
                                                                            var v147 float32 = 1e10
                                                                            if ((!true) && true) {
                                                                                v3 = v3
                                                                                var v148 int64 = 1
                                                                                _ = v148
                                                                            }
                                                                            _ = v147
                                                                        }
                                                                        _ = v145
                                                                        _ = v146
                                                                    } else {
                                                                        switch fi135 {
                                                                        case 42:
                                                                            v3 += (v3 ^ (v3 / ((567) | 1)))
                                                                        }
                                                                    }
                                                                    var v149 string = "foo bar"
                                                                    _ = v149
                                                                } else {
                                                                    switch fi135 {
                                                                    case 985:
                                                                        v150 := GenS2{F0: 1}
                                                                        _ = len(v4)
                                                                        _ = v150
                                                                    case 0:
                                                                        fmt.Println(fi135)
                                                                    case -663:
                                                                        if (((false || true) && false) && false) {
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        }
                                                                    default:
                                                                        v151 := GenS10{F0: v3, F1: 740}
                                                                        _ = v151
                                                                    }
                                                                    if true {
                                                                        if (536 < 65535) {
                                                                            v74 -= v74
                                                                            switch (fi135 % ((fi135) | 1)) {
                                                                            case -9223372036854775808:
                                                                                v1 = 2147483647
                                                                            case 42:
                                                                                v6[(^(v1))] = "hello"
                                                                            default:
                                                                                v1 -= (v1 + v1)
                                                                            }
                                                                        }
                                                                        var v152 int16 = 837
                                                                        _ = v152
                                                                    }
                                                                }
                                                            default:
                                                                if len(v5) > 0 {
                                                                    v5[0] = ("\n\t" + "")
                                                                }
                                                            }
                                                            var v153 float32 = 1.5
                                                            if (float32(1e10) < -2.25) {
                                                                v3 = (v3 << (((v3 % ((0) | 1))) & 63))
                                                                if false {
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    default:
                                                                        if len(v5) > 0 {
                                                                            v5[0] = ""
                                                                        }
                                                                        fmt.Println(fi135, v1)
                                                                    }
                                                                }
                                                            } else {
                                                                v154 := GenS4{F0: 793, F1: -276, F2: "foo bar"}
                                                                _ = v154
                                                            }
                                                            _ = v153
                                                        } else {
                                                            v1 = (v1 - 405)
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            switch "a" {
                                                            case "\n\t":
                                                                var v155 uint = 18446744073709551615
                                                                _ = v155
                                                            case "hello":
                                                                if (77 <= 127) {
                                                                    fmt.Println(v144)
                                                                    v5 = append(v5, "foo bar")
                                                                    v1 += (v1 << ((v3) & 63))
                                                                }
                                                            case "":
                                                                v3 += (v3 << ((2) & 63))
                                                            }
                                                            var v156 float32 = (-2.25 * 3.14159)
                                                            _ = v156
                                                        }
                                                        _ = v144
                                                    } else {
                                                        switch ((((("a" + "") + "foo bar") + "") + "foo bar") + (("hello" + "") + "a")) {
                                                        case "a":
                                                            var v157 string = "hello"
                                                            _ = v157
                                                        case "foo bar":
                                                            v158 := v6[v1]
                                                            _ = v158
                                                        default:
                                                            v159 := GenS12{F0: 42}
                                                            _ = v159
                                                        }
                                                        fmt.Println(fi138, v3)
                                                    }
                                                }
                                                v160 := GenS6{F0: 100, F1: -2147483648, F2: (v1 & (v1 / (((^(v1))) | 1)))}
                                                if (0 > 62) {
                                                    var v161 int8 = 125
                                                    fmt.Println(fi138, v1, fi138)
                                                    _ = v161
                                                } else {
                                                    fmt.Println(v74, fi135, v1)
                                                }
                                                _ = v160
                                            } else {
                                                v2[(^(v1))] = 2147483648
                                                var v162 string = ""
                                                fmt.Println(fi138, fi135, v74)
                                                var v163 uint8 = 209
                                                _ = v162
                                                _ = v163
                                            }
                                            delete(v4, "a")
                                            v3 = (v3 * 2)
                                        }
                                    } else {
                                        v164 := GenS13{F0: (v1 | (v1 % (((v1 / (((v1 | (^(v1)))) | 1))) | 1)))}
                                        v6[uint32(v74)] = (((("hello" + (("\n\t" + ("hello" + "\n\t")) + ("hello" + (("a" + "foo bar") + ("a" + "a"))))) + "") + "\n\t") + "a")
                                        _ = v164
                                    }
                                    switch "a" {
                                    case "":
                                        if false {
                                            fmt.Println(v3)
                                        } else {
                                            v165 := GenS14{F0: 354, F1: "\n\t", F2: 127}
                                            v1 = (v1 % ((v1) | 1))
                                            _ = len(v2)
                                            _ = v165
                                        }
                                    case "\n\t":
                                        v166 := GenS1{F0: 255, F1: 32767}
                                        _ = v166
                                    }
                                } else {
                                    switch (fi135 >> ((v3) & 63)) {
                                    case 4294967295:
                                        switch fi138 {
                                        case 578:
                                            var v167 uint = (v3 | v3)
                                            _ = v167
                                        case -1:
                                            if (!false) {
                                                if (!false) {
                                                    if (1 >= 0) {
                                                        v74 = (v74 << (((v3 % ((v3) | 1))) & 63))
                                                    } else {
                                                        v74 += (v74 << ((65535) & 63))
                                                        var v168 float64 = 0.1
                                                        fmt.Println(v3, fi135)
                                                        _ = v168
                                                    }
                                                } else {
                                                    v74 = v74
                                                    if (v74 >= (^(v74))) {
                                                        v169 := GenS9{F0: (("\n\t" + (("" + "") + "hello")) < ("foo bar" + ("a" + "\n\t")))}
                                                        fmt.Println(fi138, v74)
                                                        var v170 uint8 = uint8(v1)
                                                        _ = v169
                                                        _ = v170
                                                    }
                                                    _ = len(v2)
                                                }
                                                v171 := GenS2{F0: 41}
                                                _ = v171
                                            } else {
                                                v172 := GenS5{F0: 1, F1: (("a" + ("\n\t" + "\n\t")) + ("\n\t" + "\n\t")), F2: 362}
                                                if false {
                                                    v74 -= v74
                                                    if (!(1 < 42)) {
                                                        var v173 int16 = -32768
                                                        fmt.Println(v74)
                                                        _ = v173
                                                    } else {
                                                        v3 += 707
                                                        var v174 int64 = 9223372036854775807
                                                        v1 -= (v1 << ((v3) & 63))
                                                        _ = v174
                                                    }
                                                    v1 = v1
                                                }
                                                _ = v172
                                            }
                                        case 0:
                                            if false {
                                                var v175 string = "a"
                                                var v176 uint16 = (^(v74))
                                                if true {
                                                    v176 -= v176
                                                } else {
                                                    v74 -= (v74 << ((v3) & 63))
                                                    v175 = v175
                                                }
                                                _ = v175
                                                _ = v176
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v177 uint32 = 4294967295
                                                    _ = v177
                                                case "foo bar":
                                                    if true {
                                                        if (!false) {
                                                            v1 -= (v1 - (v1 - (v1 + 2147483647)))
                                                            v178 := v4[(("" + "a") + "")]
                                                            _ = v178
                                                        }
                                                        if (!(true || (false || false))) {
                                                            v74 = (v74 << (((^(v3))) & 63))
                                                            switch int(v1) {
                                                            case 542:
                                                                var v179 uint32 = v1
                                                                _ = v179
                                                            }
                                                        } else {
                                                            v180 := GenS15{F0: ("" + "foo bar"), F1: (v3 % (((v3 - 255)) | 1)), F2: 1}
                                                            _ = v180
                                                        }
                                                        v181 := GenS16{F0: v74, F1: -1, F2: (v1 - v1)}
                                                        _ = v181
                                                    } else {
                                                        fmt.Println(fi135)
                                                    }
                                                case "":
                                                    v3 = v3
                                                default:
                                                    v182 := GenS17{F0: 1, F1: (v74 | (v74 / (((v74 << ((v3) & 63))) | 1)))}
                                                    _ = v182
                                                }
                                            }
                                        }
                                    default:
                                        switch int(v74) {
                                        case -463:
                                            v1 -= (^(v1))
                                        case -2147483648:
                                            switch "\n\t" {
                                            case "a":
                                                if true {
                                                    if len(v5) > 0 {
                                                        v5[0] = "hello"
                                                    }
                                                }
                                            case "hello":
                                                v183 := GenS16{F0: (v74 * (v74 / ((902) | 1))), F1: 0, F2: (^(v1))}
                                                _ = v183
                                            }
                                        default:
                                            var v184 uint64 = 1
                                            _ = v184
                                        }
                                    }
                                }
                                if true {
                                    var v185 uint8 = 82
                                    if (278 > 1) {
                                        var v186 int64 = -710
                                        if (true || false) {
                                            fmt.Println(v1, v3, v3)
                                            if false {
                                                if len(v5) > 0 {
                                                    v5[0] = "foo bar"
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = ""
                                                }
                                                if (v1 <= (^(v1))) {
                                                    if true {
                                                        v74 = (v74 - v74)
                                                        var v187 int = fi135
                                                        if false {
                                                            v1 += (v1 / ((629) | 1))
                                                            var v188 float64 = 0.0
                                                            _ = v188
                                                        } else {
                                                            var v189 int64 = v186
                                                            var v190 int32 = 1
                                                            var v191 int8 = int8(v189)
                                                            v1 = 798
                                                            fmt.Println(v191, v191)
                                                            v192 := GenS10{F0: (v3 / ((v3) | 1)), F1: 319}
                                                            _ = v189
                                                            _ = v190
                                                            _ = v191
                                                            _ = v192
                                                        }
                                                        _ = v187
                                                    }
                                                    var v193 bool = false
                                                    _ = v193
                                                }
                                            }
                                            var v194 uint16 = v74
                                            _ = v194
                                        } else {
                                            _ = len(v2)
                                            v185 += 76
                                            v3 -= (v3 & v3)
                                        }
                                        if ((!(!(!(0 >= 42)))) && (float32(1e10) < ((float32(0.1) + 1e10) / 1.5))) {
                                            var v195 uint = uint(v74)
                                            var v196 int64 = (v186 >> ((v195) & 63))
                                            _ = v195
                                            _ = v196
                                        } else {
                                            var v197 string = "hello"
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = v197
                                        }
                                        _ = v186
                                    }
                                    _ = v185
                                } else {
                                    fmt.Println(v1)
                                    v198 := GenS18{F0: 0}
                                    _ = len(v2)
                                    _ = v198
                                }
                            }
                            var v199 uint8 = 111
                            _ = v199
                        }
                        fmt.Println(v1, v3, fi138)
                    }
                    v200 := GenS19{F0: 1, F1: 88}
                    _ = v200
                }
                _ = fi138
            }
        }
        if (!false) {
            v1 += (^(v1))
            v3 = (v3 & v3)
        }
        _ = fi135
        _ = v136
    }
    var v201 uint64 = 42
    v201 = v201
    switch "\n\t" {
    case "":
        switch 42 {
        case -2147483648:
            for fi202 := 0; fi202 < 5; fi202++ {
                for fi203 := 0; fi203 < 13; fi203++ {
                    v1 = (v1 * 2147483647)
                    fmt.Println(fi202)
                    _ = fi203
                }
                if (false && false) {
                    v3 += 42
                    var v204 uint = 65535
                    _ = v204
                }
                v3 -= (v3 >> ((uint(v1)) & 63))
                _ = fi202
            }
        default:
            switch "a" {
            case "hello":
                for fi205 := 0; fi205 < 6; fi205++ {
                    switch "hello" {
                    case "\n\t":
                        v206 := GenS16{F0: v74, F1: 929, F2: v1}
                        _ = v206
                    case "":
                        for fi207 := 0; fi207 < 10; fi207++ {
                            v74 = v74
                            _ = fi207
                        }
                    case "a":
                        v208 := GenS20{F0: fi205, F1: fi205}
                        _ = v208
                    }
                    if false {
                        if (v1 != v1) {
                            v201 = 2
                            v209 := GenS14{F0: 0, F1: "\n\t", F2: -1}
                            _ = v209
                        } else {
                            var v210 uint32 = 42
                            _ = v210
                        }
                        v211 := GenS21{F0: (("\n\t" + ("\n\t" + (("a" + ("a" + "\n\t")) + ("\n\t" + "a")))) != ("hello" + "")), F1: v74}
                        _ = v211
                    } else {
                        v212 := GenS22{F0: 86}
                        switch (fi205 >> (((v3 & 18446744073709551615)) & 63)) {
                        case -9223372036854775808:
                            var v213 uint8 = 233
                            _ = v213
                        case -132:
                            var v214 int64 = -9223372036854775808
                            _ = v214
                        }
                        _ = v212
                    }
                    _ = fi205
                }
            case "":
                v1 = (v1 | 42)
            case "foo bar":
                v215 := GenS23{F0: 401, F1: v1, F2: -754}
                _ = v215
            default:
                var v216 int64 = int64(v3)
                _ = v216
            }
        }
    case "hello":
        for fi217 := 0; fi217 < 17; fi217++ {
            v3 = (^(v3))
            switch fi217 {
            case 4294967295:
                for fi218 := 0; fi218 < 13; fi218++ {
                    fmt.Println(fi218, fi218)
                    fmt.Println(v201, fi217)
                    if len(v5) > 0 {
                        v5[0] = "a"
                    }
                    _ = fi218
                }
            }
            _ = fi217
        }
    case "\n\t":
        var v219 uint32 = 2
        _ = v219
    default:
        v74 += 535
    }
    for fi220 := 0; fi220 < 19; fi220++ {
        var v221 bool = true
        _ = fi220
        _ = v221
    }
    switch -689 {
    case 220:
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    case -1:
        for fi222 := 0; fi222 < 13; fi222++ {
            if true {
                var v223 int16 = -802
                v5 = append(v5, "foo bar")
                _ = v223
            }
            _ = fi222
        }
    default:
        v224 := GenS1{F0: 172, F1: 32767}
        _ = v224
    }
    switch 127 {
    case 1:
        if ("hello" < ("a" + "hello")) {
            if true {
                var v225 int32 = -1
                var v226 uint8 = 241
                switch ("\n\t" + (("a" + "a") + "\n\t")) {
                case "hello":
                    var v227 float32 = 3.14159
                    _ = v227
                }
                _ = v225
                _ = v226
            }
            for fi228 := 0; fi228 < 12; fi228++ {
                if false {
                    for fi229 := 0; fi229 < 8; fi229++ {
                        switch fi228 {
                        case 0:
                            v74 = v74
                        case 9223372036854775807:
                            var v230 int32 = -762
                            _ = v230
                        }
                        fmt.Println(fi228, v3, v74)
                        var v231 uint64 = 2
                        _ = fi229
                        _ = v231
                    }
                    v232 := GenS1{F0: 255, F1: -165}
                    _ = v232
                }
                v233 := GenS7{F0: v201, F1: v201}
                _ = fi228
                _ = v233
            }
        } else {
            for fi234 := 0; fi234 < 9; fi234++ {
                v235 := v2[uint32(v201)]
                _ = fi234
                _ = v235
            }
            for fi236 := 0; fi236 < 5; fi236++ {
                var v237 float32 = (-2.25 * 3.14159)
                var v238 uint8 = 252
                _ = fi236
                _ = v237
                _ = v238
            }
            var v239 int64 = 799
            _ = v239
        }
    case -1:
        v74 -= (v74 + (^(v74)))
    case 589:
        if (79 >= 2) {
            fmt.Println(v1)
            var v240 int64 = -9223372036854775808
            _ = v240
        } else {
            v201 = v201
            if (-1 > -1) {
                v201 = (v201 | v201)
                if (!(true && (42 >= 1))) {
                    if (int32(v74) < 2147483647) {
                        if len(v5) > 0 {
                            v5[0] = ""
                        }
                        _ = len(v6)
                        v74 -= v74
                    }
                } else {
                    v5 = append(v5, "")
                    if (("" + ((("hello" + (((("\n\t" + "") + "a") + ("foo bar" + "\n\t")) + "a")) + ("a" + "foo bar")) + "")) != ("" + (("\n\t" + ((("hello" + "foo bar") + "a") + "hello")) + "a"))) {
                        switch "\n\t" {
                        case "":
                            v241 := GenS9{F0: true}
                            _ = v241
                        case "hello":
                            if (!((^(v1)) != v1)) {
                                fmt.Println(v1, v201, v201)
                                for fi242 := 0; fi242 < 18; fi242++ {
                                    fmt.Println(v201, v201)
                                    fmt.Println(v1)
                                    v1 += v1
                                    _ = fi242
                                }
                                if false {
                                    v243 := GenS24{F0: 2}
                                    fmt.Println(v74)
                                    for fi244 := 0; fi244 < 6; fi244++ {
                                        var v245 int = 1
                                        var v246 bool = true
                                        _ = fi244
                                        _ = v245
                                        _ = v246
                                    }
                                    _ = v243
                                } else {
                                    v74 += v74
                                    v247 := GenS4{F0: 1, F1: -732, F2: ("a" + "a")}
                                    _ = v247
                                }
                            }
                        case "foo bar":
                            for fi248 := 0; fi248 < 15; fi248++ {
                                v249 := GenS25{F0: (v1 | 0)}
                                _ = fi248
                                _ = v249
                            }
                        default:
                            for fi250 := 0; fi250 < 11; fi250++ {
                                for fi251 := 0; fi251 < 6; fi251++ {
                                    var v252 bool = (!(false || false))
                                    v252 = true
                                    v253 := GenS2{F0: 1}
                                    _ = fi251
                                    _ = v252
                                    _ = v253
                                }
                                switch fi250 {
                                case 128:
                                    for fi254 := 0; fi254 < 12; fi254++ {
                                        var v255 int8 = 87
                                        fmt.Println(v74, fi250)
                                        _ = fi254
                                        _ = v255
                                    }
                                case -1:
                                    switch "foo bar" {
                                    case "\n\t":
                                        if false {
                                            v256 := GenS13{F0: (^(v1))}
                                            _ = v256
                                        } else {
                                            var v257 int = (fi250 + (fi250 & 127))
                                            for fi258 := 0; fi258 < 10; fi258++ {
                                                var v259 uint = v3
                                                _ = fi258
                                                _ = v259
                                            }
                                            _ = v257
                                        }
                                    case "foo bar":
                                        v3 += v3
                                    case "":
                                        v260 := GenS24{F0: 1}
                                        _ = v260
                                    }
                                default:
                                    v3 = v3
                                }
                                _ = fi250
                            }
                        }
                        for fi261 := 0; fi261 < 20; fi261++ {
                            v262 := GenS14{F0: 1000, F1: "foo bar", F2: 0}
                            switch "foo bar" {
                            case "\n\t":
                                var v263 string = "hello"
                                _ = v263
                            default:
                                switch (^(fi261)) {
                                case 128:
                                    v5 = append(v5, "")
                                case 282:
                                    for fi264 := 0; fi264 < 20; fi264++ {
                                        fmt.Println(fi261)
                                        _ = fi264
                                    }
                                case 2147483647:
                                    if false {
                                        v201 = (v201 * (v201 | uint64(v1)))
                                        v74 = v74
                                    } else {
                                        fmt.Println(v1)
                                        for fi265 := 0; fi265 < 3; fi265++ {
                                            if false {
                                                var v266 int32 = 136
                                                delete(v2, v1)
                                                _ = v266
                                            } else {
                                                fmt.Println(v74, v74, v3)
                                                v267 := v4[""]
                                                v268 := GenS26{F0: 44}
                                                _ = v267
                                                _ = v268
                                            }
                                            _ = fi265
                                        }
                                        v74 += (v74 | (v74 >> (((v3 << (((v3 ^ uint(v1))) & 63))) & 63)))
                                    }
                                default:
                                    var v269 int16 = 1
                                    _ = v269
                                }
                            }
                            _ = fi261
                            _ = v262
                        }
                    } else {
                        fmt.Println(v201)
                    }
                }
                fmt.Println(v201, v74, v3)
            } else {
                if true {
                    v5 = append(v5, (("\n\t" + ("\n\t" + (("foo bar" + "a") + "hello"))) + ("foo bar" + "hello")))
                    v1 = v1
                    v270 := GenS27{F0: 128, F1: false}
                    _ = v270
                } else {
                    v271 := GenS13{F0: v1}
                    var v272 int32 = 1
                    if ((float32(-2.25) != 1.5) || ((((42 >= 100) && true) || false) && true)) {
                        v5 = append(v5, ("" + "\n\t"))
                        switch -2147483648 {
                        case 133:
                            var v273 bool = true
                            _ = v273
                        case 0:
                            var v274 uint64 = 681
                            _ = v274
                        case -867:
                            var v275 uint8 = 2
                            _ = v275
                        }
                        if false {
                            fmt.Println(v1, v1, v1)
                            var v276 uint8 = 2
                            _ = v276
                        } else {
                            v272 -= v272
                        }
                    }
                    _ = v271
                    _ = v272
                }
                for fi277 := 0; fi277 < 15; fi277++ {
                    _ = len(v5)
                    var v278 float32 = (float32((1e10 + (-2.25 + 1.5))) / (1e10 - float32((-2.25 * ((-2.25 - (float64(0.0) / 1.0)) + 0.0)))))
                    if (42 < -9223372036854775808) {
                        fmt.Println(v74, v3, fi277)
                        var v279 int16 = -32768
                        _ = v279
                    }
                    _ = fi277
                    _ = v278
                }
                for fi280 := 0; fi280 < 12; fi280++ {
                    v74 = (v74 << ((v3) & 63))
                    if (1 > 1000) {
                        v74 = (^(v74))
                        var v281 string = "foo bar"
                        if (0.0 < (float64(1.5) * 3.14159)) {
                            if len(v5) > 0 {
                                v5[0] = "hello"
                            }
                            if (!true) {
                                var v282 string = (v281 + v281)
                                v201 = 42
                                _ = v282
                            } else {
                                v283 := GenS28{F0: int8(v3), F1: 42}
                                fmt.Println(v74, v3, v201)
                                if false {
                                    if true {
                                        for fi284 := 0; fi284 < 7; fi284++ {
                                            if true {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            _ = fi284
                                        }
                                        var v285 uint16 = v74
                                        v201 = v201
                                        _ = v285
                                    }
                                    for fi286 := 0; fi286 < 20; fi286++ {
                                        if false {
                                            fmt.Println(fi286, v281, fi286)
                                        }
                                        v3 += v3
                                        if len(v5) > 0 {
                                            v5[0] = v281
                                        }
                                        var v287 float64 = 3.14159
                                        var v288 int16 = 32767
                                        _ = fi286
                                        _ = v287
                                        _ = v288
                                    }
                                    v289 := GenS11{F0: 42, F1: "\n\t"}
                                    _ = v289
                                }
                                _ = v283
                            }
                        } else {
                            v6[4294967295] = ("\n\t" + ((v281 + ("" + v281)) + v281))
                        }
                        _ = v281
                    }
                    _ = fi280
                }
            }
        }
    }
    if true {
        v290 := GenS10{F0: 65535, F1: -97}
        _ = v290
    }
    v74 += v74
    v1 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v201:", v201)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v74:", v74)
    _ = v1
    _ = v134
    _ = v2
    _ = v201
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
}
