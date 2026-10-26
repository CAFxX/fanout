package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 uint16; F1 string; F2 int16 }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 string; F1 uint }
type GenS16 struct { F0 uint; F1 string; F2 int32 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 int64; F1 uint; F2 int16 }
type GenS19 struct { F0 bool; F1 int32; F2 int32 }
type GenS20 struct { F0 uint; F1 uint }
type GenS21 struct { F0 uint32; F1 uint64 }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 uint; F1 uint64 }
type GenS24 struct { F0 string }
type GenS25 struct { F0 string; F1 uint32 }
type GenS26 struct { F0 int64 }
type GenS27 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS28 struct { F0 int8 }
type GenS29 struct { F0 uint16 }
type GenS30 struct { F0 int32; F1 uint64 }
func genFn0(p0 int8, p1 int64) int64 {
    switch 1 {
    case -2147483648:
        if false {
            v6 := []int{0}
            p0 += p0
            _ = v6
        }
    case 9223372036854775807:
        p0 -= -1
    case -307:
        p1 += p1
    }
    var v7 uint = 0
    if (42 == 42) {
        switch 952 {
        case 346:
            for fi8 := 0; fi8 < 7; fi8++ {
                fmt.Println(p1, p1, v7)
                switch fi8 {
                case 788:
                    v9 := GenS1{F0: 1}
                    _ = v9
                case 621:
                    p0 = p0
                default:
                    v7 -= (v7 | 42)
                }
                _ = fi8
            }
        default:
            var v10 float32 = -2.25
            _ = v10
        }
    } else {
        if true {
            var v11 bool = true
            for fi12 := 0; fi12 < 16; fi12++ {
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := GenS2{F0: 42}
                    _ = fi13
                    _ = v14
                }
                p1 += p1
                _ = fi12
            }
            var v15 float64 = 1.5
            _ = v11
            _ = v15
        }
        if true {
            switch ("foo bar" + ("" + "hello")) {
            case "\n\t":
                v16 := []int{128, 42, 32768, 128}
                _ = v16
            default:
                v17 := GenS3{F0: p0, F1: 974, F2: 1}
                _ = v17
            }
        } else {
            var v18 uint8 = 27
            for fi19 := 0; fi19 < 7; fi19++ {
                if (true || ((p0 & p0) >= 0)) {
                    v20 := []int{fi19}
                    _ = v20
                } else {
                    if (32767 >= 2) {
                        p0 += (p0 & (p0 << ((v7) & 63)))
                    } else {
                        var v21 int8 = (p0 - (p0 | p0))
                        for fi22 := 0; fi22 < 19; fi22++ {
                            v23 := []int{fi19, -627}
                            if false {
                                p0 = 100
                                p0 += 0
                            }
                            fmt.Println(v7)
                            _ = fi22
                            _ = v23
                        }
                        if true {
                            var v24 int16 = 32767
                            _ = v24
                        }
                        _ = v21
                    }
                    var v25 float64 = 3.14159
                    _ = v25
                }
                if ((v18 - 1) == (v18 % ((v18) | 1))) {
                    var v26 uint16 = 112
                    switch "foo bar" {
                    case "a":
                        v27 := []int{fi19, fi19, (^(fi19))}
                        _ = v27
                    case "":
                        var v28 int16 = int16(v18)
                        _ = v28
                    default:
                        var v29 uint32 = 2
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 14; fi30++ {
                        v31 := GenS4{F0: (^(p0))}
                        if false {
                            switch (fi30 / ((fi30) | 1)) {
                            case -869:
                                switch "\n\t" {
                                case "hello":
                                    var v32 int16 = 0
                                    _ = v32
                                default:
                                    var v33 uint = v7
                                    _ = v33
                                }
                            case 673:
                                var v34 float32 = (-2.25 / 1.5)
                                v35 := GenS2{F0: 894}
                                _ = v34
                                _ = v35
                            case 32767:
                                if (!(127 >= p0)) {
                                    switch "a" {
                                    case "foo bar":
                                        v26 += v26
                                    }
                                    v36 := GenS5{F0: v18, F1: (v18 << (((v7 | v7)) & 63))}
                                    _ = v36
                                }
                            }
                        }
                        _ = fi30
                        _ = v31
                    }
                    _ = v26
                } else {
                    for fi37 := 0; fi37 < 12; fi37++ {
                        var v38 uint32 = 0
                        _ = fi37
                        _ = v38
                    }
                    var v39 bool = true
                    _ = v39
                }
                if false {
                    fmt.Println(v18)
                    for fi40 := 0; fi40 < 15; fi40++ {
                        v41 := GenS6{F0: 127, F1: p0}
                        _ = fi40
                        _ = v41
                    }
                }
                _ = fi19
            }
            _ = v18
        }
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        if (3.14159 < 0.1) {
            switch fi42 {
            case -2147483648:
                v43 := []int{(fi42 * 127), 4294967295, fi42, fi42}
                _ = v43
            case 127:
                if false {
                    var v44 float32 = 0.1
                    if ((fi42 << ((v7) & 63)) == fi42) {
                        fmt.Println(fi42, v44, p1)
                        if false {
                            p0 += p0
                            v45 := []int{fi42, 127}
                            _ = v45
                        } else {
                            v46 := GenS3{F0: (p0 + p0), F1: 4294967295, F2: 42}
                            v47 := []int{fi42, fi42, 0}
                            _ = v46
                            _ = v47
                        }
                    } else {
                        if (-2147483648 > 2147483647) {
                            var v48 uint8 = 1
                            v49 := []int{fi42}
                            for fi50 := 0; fi50 < 13; fi50++ {
                                if len(v49) >= 2 {
                                    v49 = v49[1:len(v49)]
                                }
                                v7 = 1
                                var v51 int64 = (-(p1))
                                _ = fi50
                                _ = v51
                            }
                            _ = v48
                            _ = v49
                        }
                        if false {
                            v52 := GenS7{F0: 42, F1: ("hello" + "a"), F2: 42}
                            fmt.Println(v44)
                            _ = v52
                        }
                    }
                    var v53 int32 = 751
                    _ = v44
                    _ = v53
                }
            default:
                p1 = p1
            }
            if false {
                v54 := GenS8{F0: p1, F1: 0}
                for fi55 := 0; fi55 < 12; fi55++ {
                    var v56 float32 = 1e10
                    var v57 int32 = 2147483647
                    switch "" {
                    case "foo bar":
                        if (true || ((255 < 105) || false)) {
                            switch fi55 {
                            case -9223372036854775808:
                                var v58 int64 = -1
                                _ = v58
                            }
                            if (1.5 >= -2.25) {
                                v59 := []int{(-(fi55)), 162}
                                switch 133 {
                                case 4294967295:
                                    _ = len(v59)
                                case 9223372036854775807:
                                    if len(v59) >= 2 {
                                        v59 = v59[1:len(v59)]
                                    }
                                default:
                                    p1 = (p1 - (p1 / ((1) | 1)))
                                }
                                v7 += v7
                                _ = v59
                            } else {
                                switch 4294967295 {
                                case -664:
                                    switch "a" {
                                    case "":
                                        var v60 bool = true
                                        _ = v60
                                    default:
                                        v61 := GenS2{F0: 533}
                                        _ = v61
                                    }
                                default:
                                    v57 -= v57
                                }
                                var v62 int8 = p0
                                fmt.Println(v56)
                                _ = v62
                            }
                        } else {
                            if true {
                                p0 = p0
                            } else {
                                v56 = (v56 * 0.1)
                                var v63 int = 344
                                p1 = p1
                                if false {
                                    v64 := GenS7{F0: 42, F1: ((("" + "\n\t") + "a") + ((((("" + (("\n\t" + "hello") + ((("a" + (("a" + "\n\t") + "a")) + (("foo bar" + "a") + "a")) + ("hello" + (("a" + ("a" + "")) + "hello"))))) + "foo bar") + "") + "foo bar") + "a")), F2: 32767}
                                    _ = v64
                                } else {
                                    v65 := []int{}
                                    if false {
                                        switch "foo bar" {
                                        case "foo bar":
                                            if (true || ("" != ("a" + ((((((("foo bar" + "hello") + "\n\t") + "a") + (("hello" + "\n\t") + "foo bar")) + ("a" + "")) + "\n\t") + "\n\t")))) {
                                                v57 = (v57 % (((^(v57))) | 1))
                                                if true {
                                                    v56 = (v56 + v56)
                                                } else {
                                                    v7 -= v7
                                                    fmt.Println(v63, p1)
                                                    v7 = 509
                                                }
                                                if true {
                                                    var v66 bool = false
                                                    p1 = (p1 % ((p1) | 1))
                                                    _ = len(v65)
                                                    _ = v66
                                                } else {
                                                    if true {
                                                        if (-32768 < 42) {
                                                            var v67 int64 = p1
                                                            _ = v67
                                                        } else {
                                                            switch v63 {
                                                            case 2147483647:
                                                                var v68 uint8 = 226
                                                                _ = v68
                                                            case -671:
                                                                if true {
                                                                    var v69 uint64 = 4294967295
                                                                    if len(v65) >= 2 {
                                                                        v65 = v65[1:len(v65)]
                                                                    }
                                                                    var v70 int64 = (p1 * p1)
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                            default:
                                                                fmt.Println(v56, p0, fi42)
                                                            }
                                                            var v71 int16 = -32768
                                                            _ = v71
                                                        }
                                                        p0 -= -6
                                                        var v72 uint32 = 2147483647
                                                        _ = v72
                                                    } else {
                                                        if false {
                                                            if (1.5 > (0.1 / 1.0)) {
                                                                v56 = (v56 - 1e10)
                                                                if (!(!false)) {
                                                                    _ = len(v65)
                                                                    v73 := GenS1{F0: 0}
                                                                    var v74 string = "foo bar"
                                                                    v65 = append(v65, (fi55 - (fi55 | (fi42 % ((v63) | 1)))))
                                                                    _ = v73
                                                                    _ = v74
                                                                }
                                                                switch fi42 {
                                                                case 9223372036854775807:
                                                                    fmt.Println(v56, v7, v57)
                                                                    v75 := GenS1{F0: uint64(v7)}
                                                                    v56 += ((v56 / -2.25) - 0.1)
                                                                    v76 := GenS6{F0: 2, F1: p0}
                                                                    _ = v75
                                                                    _ = v76
                                                                case -469:
                                                                    v77 := GenS9{F0: 65535, F1: 42}
                                                                    _ = v77
                                                                default:
                                                                    if (!(true || (!(((-1 != fi42) || (482 > 64)) || false)))) {
                                                                        if len(v65) > 0 {
                                                                            v65[0] = v63
                                                                        }
                                                                        var v78 int8 = (p0 & (p0 % ((42) | 1)))
                                                                        switch (fi42 ^ 127) {
                                                                        case -450:
                                                                            var v79 float64 = 0.1
                                                                            _ = v79
                                                                        case 42:
                                                                            var v80 int64 = (p1 - 42)
                                                                            _ = v80
                                                                        default:
                                                                            p1 += p1
                                                                        }
                                                                        _ = v78
                                                                    } else {
                                                                        switch (^(v63)) {
                                                                        case -161:
                                                                            v81 := GenS10{F0: (fi42 / ((127) | 1))}
                                                                            _ = v81
                                                                        }
                                                                        if (uint8(v7) < 42) {
                                                                            v57 += -1
                                                                            switch ("" + "a") {
                                                                            case "":
                                                                                if false {
                                                                                    v82 := GenS10{F0: v63}
                                                                                    _ = v82
                                                                                }
                                                                            case "\n\t":
                                                                                if len(v65) >= 2 {
                                                                                    v65 = v65[1:len(v65)]
                                                                                }
                                                                            default:
                                                                                v56 = (v56 - ((v56 - float32(1.5)) * ((v56 / ((0.1 / (v56 + ((float32(1.5) + v56) * 0.1))) / v56)) * 0.1)))
                                                                            }
                                                                            if (548 >= 202) {
                                                                                var v83 string = "\n\t"
                                                                                var v84 float32 = 1e10
                                                                                var v85 uint = v7
                                                                                _ = v83
                                                                                _ = v84
                                                                                _ = v85
                                                                            } else {
                                                                                fmt.Println(p1)
                                                                                v86 := GenS4{F0: (p0 ^ 126)}
                                                                                _ = v86
                                                                            }
                                                                        } else {
                                                                            var v87 string = "a"
                                                                            _ = v87
                                                                        }
                                                                    }
                                                                }
                                                                if ((p0 / (((p0 / (((p0 << ((135) & 63))) | 1))) | 1)) >= (p0 >> ((v7) & 63))) {
                                                                    if (1 < 667) {
                                                                        p0 += (p0 % (((p0 - p0)) | 1))
                                                                    } else {
                                                                        v56 = v56
                                                                        if ((v7 & v7) > (v7 >> ((2) & 63))) {
                                                                            v57 -= -281
                                                                            switch ("a" + "") {
                                                                            case "foo bar":
                                                                                p0 -= p0
                                                                            case "\n\t":
                                                                                if (false || false) {
                                                                                    if false {
                                                                                        v56 -= (float32(1.5) - 3.14159)
                                                                                        var v88 uint = (v7 % ((0) | 1))
                                                                                        if len(v65) >= 2 {
                                                                                            v65 = v65[1:len(v65)]
                                                                                        }
                                                                                        _ = v88
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v89 int8 = int8(v7)
                                                                                _ = v89
                                                                            }
                                                                            v90 := GenS5{F0: 2, F1: 76}
                                                                            _ = v90
                                                                        }
                                                                        switch ((("hello" + "") + ("foo bar" + ("foo bar" + ("" + ("a" + "hello"))))) + "") {
                                                                        case "\n\t":
                                                                            fmt.Println(v57, fi55, p0)
                                                                        case "foo bar":
                                                                            var v91 uint64 = 42
                                                                            _ = v91
                                                                        }
                                                                    }
                                                                    switch ("a" + "foo bar") {
                                                                    case "hello":
                                                                        switch fi55 {
                                                                        case -9223372036854775808:
                                                                            v92 := GenS10{F0: 4294967295}
                                                                            _ = v92
                                                                        case 9223372036854775807:
                                                                            v65 = append(v65, (v63 & fi55))
                                                                        }
                                                                    case "a":
                                                                        v93 := GenS4{F0: (p0 | (p0 / ((p0) | 1)))}
                                                                        _ = v93
                                                                    case "":
                                                                        _ = len(v65)
                                                                    }
                                                                    v94 := GenS10{F0: (fi55 << (((v7 ^ (v7 - (v7 - (^(v7)))))) & 63))}
                                                                    _ = v94
                                                                }
                                                                if true {
                                                                    if true {
                                                                        _ = len(v65)
                                                                    } else {
                                                                        v95 := GenS5{F0: 42, F1: 92}
                                                                        switch ("hello" + ("hello" + "")) {
                                                                        case "a":
                                                                            fmt.Println(v63, fi55, fi42)
                                                                        default:
                                                                            if ((!((("" + "foo bar") + ((("hello" + ("foo bar" + "hello")) + "foo bar") + ("\n\t" + "a"))) >= ("hello" + "hello"))) || false) {
                                                                                v96 := GenS11{F0: ""}
                                                                                if len(v65) >= 2 {
                                                                                    v65 = v65[1:len(v65)]
                                                                                }
                                                                                v7 += (^(v7))
                                                                                _ = v96
                                                                            } else {
                                                                                v97 := GenS2{F0: 0}
                                                                                if ((float32(1.5) + -2.25) > 3.14159) {
                                                                                    v65 = append(v65, v63)
                                                                                    if (!((p0 - p0) <= -1)) {
                                                                                        v65 = append(v65, (^(fi55)))
                                                                                        if (true || ((v7 + 177) == v7)) {
                                                                                            if (255 <= 1) {
                                                                                                var v98 int64 = p1
                                                                                                _ = v98
                                                                                            } else {
                                                                                                v57 -= (v57 >> ((v7) & 63))
                                                                                                if true {
                                                                                                    v56 = v56
                                                                                                    switch "hello" {
                                                                                                    case "hello":
                                                                                                        v57 += v57
                                                                                                    case "foo bar":
                                                                                                        _ = len(v65)
                                                                                                    }
                                                                                                    var v99 uint = v7
                                                                                                    _ = v99
                                                                                                }
                                                                                                switch ("" + (("a" + "") + "a")) {
                                                                                                case "\n\t":
                                                                                                    if true {
                                                                                                        var v100 uint8 = 2
                                                                                                        var v101 int64 = p1
                                                                                                        _ = v100
                                                                                                        _ = v101
                                                                                                    } else {
                                                                                                        fmt.Println(v56, v63, v7)
                                                                                                        v102 := GenS12{F0: v57}
                                                                                                        _ = len(v65)
                                                                                                        _ = v102
                                                                                                    }
                                                                                                case "hello":
                                                                                                    _ = len(v65)
                                                                                                }
                                                                                            }
                                                                                            if len(v65) >= 2 {
                                                                                                v65 = v65[1:len(v65)]
                                                                                            }
                                                                                        }
                                                                                        v103 := GenS13{F0: 128}
                                                                                        _ = v103
                                                                                    } else {
                                                                                        var v104 int64 = (^(p1))
                                                                                        p1 += v104
                                                                                        v105 := GenS14{F0: false}
                                                                                        _ = v104
                                                                                        _ = v105
                                                                                    }
                                                                                    if ((v56 * ((0.1 * v56) + 1e10)) == (v56 - (v56 * v56))) {
                                                                                        fmt.Println(v7, fi42, v63)
                                                                                        switch "a" {
                                                                                        case "foo bar":
                                                                                            var v106 uint8 = 0
                                                                                            _ = v106
                                                                                        default:
                                                                                            if false {
                                                                                                v107 := GenS10{F0: v63}
                                                                                                fmt.Println(fi42, p1)
                                                                                                switch "a" {
                                                                                                case "":
                                                                                                    var v108 bool = false
                                                                                                    _ = v108
                                                                                                case "foo bar":
                                                                                                    switch v63 {
                                                                                                    case 166:
                                                                                                        v109 := GenS13{F0: 128}
                                                                                                        _ = v109
                                                                                                    case 9223372036854775807:
                                                                                                        v110 := GenS15{F0: "\n\t", F1: 861}
                                                                                                        _ = v110
                                                                                                    case 4294967295:
                                                                                                        switch v63 {
                                                                                                        case 165:
                                                                                                            p0 -= (p0 % ((-44) | 1))
                                                                                                        case -139:
                                                                                                            p0 -= (p0 * (p0 + p0))
                                                                                                        }
                                                                                                        fmt.Println(v7)
                                                                                                    default:
                                                                                                        switch "a" {
                                                                                                        case "a":
                                                                                                            var v111 uint = 65535
                                                                                                            _ = v111
                                                                                                        case "\n\t":
                                                                                                            if true {
                                                                                                                v112 := GenS15{F0: "hello", F1: 65535}
                                                                                                                _ = v112
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            switch (fi55 - fi42) {
                                                                                                            case -896:
                                                                                                                switch "" {
                                                                                                                case "\n\t":
                                                                                                                    var v113 int = (fi55 | fi42)
                                                                                                                    _ = v113
                                                                                                                case "foo bar":
                                                                                                                    fmt.Println(p0)
                                                                                                                case "":
                                                                                                                    switch fi42 {
                                                                                                                    case 32768:
                                                                                                                        v114 := GenS1{F0: 1}
                                                                                                                        _ = v114
                                                                                                                    default:
                                                                                                                        if true {
                                                                                                                            var v115 uint8 = 2
                                                                                                                            v7 = v7
                                                                                                                            switch (fi42 << ((v7) & 63)) {
                                                                                                                            case 32767:
                                                                                                                                fmt.Println(fi42)
                                                                                                                            case -9223372036854775808:
                                                                                                                                v116 := GenS5{F0: v115, F1: 92}
                                                                                                                                switch (fi55 + (v63 & (fi42 - (v63 | v63)))) {
                                                                                                                                case 907:
                                                                                                                                    v117 := GenS6{F0: (v115 & (v115 * (v115 - (v115 >> ((v7) & 63))))), F1: (^(p0))}
                                                                                                                                    _ = v117
                                                                                                                                }
                                                                                                                                _ = v116
                                                                                                                            case 128:
                                                                                                                                switch "\n\t" {
                                                                                                                                case "hello":
                                                                                                                                    var v118 uint64 = uint64(v7)
                                                                                                                                    _ = v118
                                                                                                                                case "a":
                                                                                                                                    var v119 float64 = 1e10
                                                                                                                                    _ = v119
                                                                                                                                case "\n\t":
                                                                                                                                    var v120 uint = (v7 & v7)
                                                                                                                                    _ = v120
                                                                                                                                default:
                                                                                                                                    v63 = v63
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                if len(v65) > 0 {
                                                                                                                                    v65[0] = v63
                                                                                                                                }
                                                                                                                                if false {
                                                                                                                                    switch (fi55 / ((fi42) | 1)) {
                                                                                                                                    case 1:
                                                                                                                                        if len(v65) > 0 {
                                                                                                                                            v65[0] = 9223372036854775807
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        switch (v63 << (((^(v7))) & 63)) {
                                                                                                                                        case -9223372036854775808:
                                                                                                                                            if len(v65) >= 2 {
                                                                                                                                                v65 = v65[1:len(v65)]
                                                                                                                                            }
                                                                                                                                        case -344:
                                                                                                                                            if false {
                                                                                                                                                p0 -= (^(p0))
                                                                                                                                                var v121 int16 = -1
                                                                                                                                                _ = v121
                                                                                                                                            } else {
                                                                                                                                                v65 = append(v65, (v63 - 32767))
                                                                                                                                                if len(v65) >= 2 {
                                                                                                                                                    v65 = v65[1:len(v65)]
                                                                                                                                                }
                                                                                                                                                v122 := GenS14{F0: (!(!true))}
                                                                                                                                                _ = v122
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            if false {
                                                                                                                                                if len(v65) >= 2 {
                                                                                                                                                    v65 = v65[1:len(v65)]
                                                                                                                                                }
                                                                                                                                            } else {
                                                                                                                                                v123 := GenS7{F0: 32767, F1: "a", F2: 834}
                                                                                                                                                p0 += -84
                                                                                                                                                _ = v123
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    fmt.Println(fi42, p0)
                                                                                                                                }
                                                                                                                                _ = len(v65)
                                                                                                                            }
                                                                                                                            _ = v115
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v63 = fi42
                                                                                                                    var v124 uint64 = 188
                                                                                                                    _ = v124
                                                                                                                }
                                                                                                                fmt.Println(p0, fi42, p1)
                                                                                                                var v125 int32 = 498
                                                                                                                _ = v125
                                                                                                            case 0:
                                                                                                                switch 32767 {
                                                                                                                case 128:
                                                                                                                    if ((^(v7)) >= v7) {
                                                                                                                        v126 := GenS16{F0: (v7 << ((v7) & 63)), F1: "a", F2: (v57 | v57)}
                                                                                                                        if false {
                                                                                                                            var v127 uint32 = 2147483648
                                                                                                                            if false {
                                                                                                                                var v128 float32 = 0.0
                                                                                                                                _ = v128
                                                                                                                            } else {
                                                                                                                                v129 := GenS17{F0: p1}
                                                                                                                                _ = v129
                                                                                                                            }
                                                                                                                            _ = v127
                                                                                                                        } else {
                                                                                                                            p1 = p1
                                                                                                                            switch ((("" + "\n\t") + "\n\t") + (("hello" + (("a" + "\n\t") + "a")) + "a")) {
                                                                                                                            case "\n\t":
                                                                                                                                p0 = -1
                                                                                                                            case "":
                                                                                                                                var v130 uint64 = uint64(v7)
                                                                                                                                _ = v130
                                                                                                                            case "a":
                                                                                                                                v7 = (^(v7))
                                                                                                                            default:
                                                                                                                                v65 = append(v65, (fi42 * (v63 / (((v63 - 127)) | 1))))
                                                                                                                            }
                                                                                                                            if len(v65) > 0 {
                                                                                                                                v65[0] = v63
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = v126
                                                                                                                    } else {
                                                                                                                        var v131 int = (fi55 ^ (v63 >> ((v7) & 63)))
                                                                                                                        fmt.Println(fi55, p0, fi42)
                                                                                                                        _ = v131
                                                                                                                    }
                                                                                                                case 4294967295:
                                                                                                                    v7 += (v7 * v7)
                                                                                                                case -964:
                                                                                                                    if (!true) {
                                                                                                                        v57 += -1
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                if true {
                                                                                                                    var v132 int = (fi55 % ((fi42) | 1))
                                                                                                                    if true {
                                                                                                                        if true {
                                                                                                                            var v133 uint8 = 255
                                                                                                                            if (v57 <= 971) {
                                                                                                                                v56 = v56
                                                                                                                                fmt.Println(p1, p1)
                                                                                                                            }
                                                                                                                            v134 := GenS7{F0: 42, F1: "a", F2: 0}
                                                                                                                            _ = v133
                                                                                                                            _ = v134
                                                                                                                        }
                                                                                                                        var v135 int32 = v57
                                                                                                                        p0 = (p0 ^ (p0 * 100))
                                                                                                                        _ = v135
                                                                                                                    }
                                                                                                                    switch ((((("foo bar" + "foo bar") + (("\n\t" + ("\n\t" + "")) + "foo bar")) + ("" + ("\n\t" + ("" + ("\n\t" + "foo bar"))))) + "\n\t") + (("hello" + "a") + "foo bar")) {
                                                                                                                    case "foo bar":
                                                                                                                        if false {
                                                                                                                            var v136 string = (("a" + ("hello" + "foo bar")) + "")
                                                                                                                            _ = v136
                                                                                                                        } else {
                                                                                                                            v137 := GenS18{F0: p1, F1: v7, F2: -1}
                                                                                                                            _ = v137
                                                                                                                        }
                                                                                                                    case "hello":
                                                                                                                        if len(v65) > 0 {
                                                                                                                            v65[0] = 32767
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v138 string = "hello"
                                                                                                                    _ = v132
                                                                                                                    _ = v138
                                                                                                                }
                                                                                                                var v139 bool = true
                                                                                                                v140 := GenS1{F0: 846}
                                                                                                                switch ("a" + "\n\t") {
                                                                                                                case "hello":
                                                                                                                    v141 := GenS9{F0: 2, F1: p0}
                                                                                                                    _ = v141
                                                                                                                case "foo bar":
                                                                                                                    var v142 uint = v7
                                                                                                                    _ = v142
                                                                                                                case "a":
                                                                                                                    v56 -= v56
                                                                                                                }
                                                                                                                var v143 uint8 = 234
                                                                                                                v144 := GenS17{F0: p1}
                                                                                                                fmt.Println(v143)
                                                                                                                _ = v139
                                                                                                                _ = v140
                                                                                                                _ = v143
                                                                                                                _ = v144
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(fi42, fi55)
                                                                                                if (!true) {
                                                                                                    var v145 uint32 = 692
                                                                                                    fmt.Println(v7)
                                                                                                    _ = v145
                                                                                                } else {
                                                                                                    switch ("foo bar" + ((("" + "\n\t") + "foo bar") + "hello")) {
                                                                                                    case "a":
                                                                                                        fmt.Println(v56)
                                                                                                    case "\n\t":
                                                                                                        fmt.Println(v7, v56, p0)
                                                                                                    }
                                                                                                }
                                                                                                var v146 int64 = p1
                                                                                                _ = v107
                                                                                                _ = v146
                                                                                            }
                                                                                        }
                                                                                        switch 32768 {
                                                                                        case 128:
                                                                                            fmt.Println(fi55)
                                                                                        default:
                                                                                            if (v57 <= (-(v57))) {
                                                                                                v7 -= v7
                                                                                                if true {
                                                                                                    if len(v65) >= 2 {
                                                                                                        v65 = v65[1:len(v65)]
                                                                                                    }
                                                                                                    _ = len(v65)
                                                                                                } else {
                                                                                                    switch 32767 {
                                                                                                    case -955:
                                                                                                        p1 += 0
                                                                                                    case -2147483648:
                                                                                                        var v147 int8 = p0
                                                                                                        _ = v147
                                                                                                    default:
                                                                                                        v65 = append(v65, 0)
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                if (false || (p0 != 127)) {
                                                                                                    if true {
                                                                                                        p0 += 42
                                                                                                    } else {
                                                                                                        var v148 bool = false
                                                                                                        var v149 string = ("" + ("foo bar" + "\n\t"))
                                                                                                        v65 = append(v65, 7)
                                                                                                        _ = v148
                                                                                                        _ = v149
                                                                                                    }
                                                                                                }
                                                                                                v63 += fi55
                                                                                                var v150 uint8 = 0
                                                                                                _ = v150
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        if (397 >= 4294967295) {
                                                                                            switch (^(fi42)) {
                                                                                            case -740:
                                                                                                var v151 uint = (^(v7))
                                                                                                _ = v151
                                                                                            case 1:
                                                                                                if false {
                                                                                                    if true {
                                                                                                        v152 := GenS3{F0: p0, F1: 497, F2: 2}
                                                                                                        var v153 uint8 = 1
                                                                                                        _ = v152
                                                                                                        _ = v153
                                                                                                    }
                                                                                                }
                                                                                            case 32768:
                                                                                                _ = len(v65)
                                                                                            }
                                                                                            fmt.Println(v56, fi42, v56)
                                                                                        }
                                                                                        switch ("" + ("a" + "a")) {
                                                                                        case "a":
                                                                                            v65 = append(v65, (v63 >> (((v7 / ((v7) | 1))) & 63)))
                                                                                        case "hello":
                                                                                            var v154 int16 = 601
                                                                                            _ = v154
                                                                                        default:
                                                                                            switch (fi55 | fi42) {
                                                                                            case 2147483647:
                                                                                                v155 := GenS2{F0: 18446744073709551615}
                                                                                                _ = v155
                                                                                            case -819:
                                                                                                switch (fi55 | -112) {
                                                                                                case 32768:
                                                                                                    switch fi42 {
                                                                                                    case 42:
                                                                                                        var v156 uint = (v7 >> (((^(v7))) & 63))
                                                                                                        _ = v156
                                                                                                    case 127:
                                                                                                        fmt.Println(p1, fi55)
                                                                                                    default:
                                                                                                        v157 := GenS8{F0: 1, F1: 1}
                                                                                                        _ = v157
                                                                                                    }
                                                                                                    var v158 uint32 = 42
                                                                                                    _ = v158
                                                                                                }
                                                                                            case 4294967295:
                                                                                                v63 = (fi42 - v63)
                                                                                            default:
                                                                                                if len(v65) >= 2 {
                                                                                                    v65 = v65[1:len(v65)]
                                                                                                }
                                                                                            }
                                                                                            var v159 uint32 = 970
                                                                                            if true {
                                                                                                fmt.Println(p1, p0)
                                                                                            }
                                                                                            fmt.Println(fi55, p1, v63)
                                                                                            _ = v159
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v97
                                                                            }
                                                                        }
                                                                        _ = v95
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        if (4294967295 >= 681) {
                                                                            fmt.Println(v57)
                                                                            v160 := GenS1{F0: 42}
                                                                            v161 := GenS19{F0: false, F1: v57, F2: (v57 - -1)}
                                                                            _ = v160
                                                                            _ = v161
                                                                        } else {
                                                                            fmt.Println(p0, p1, fi42)
                                                                            if len(v65) >= 2 {
                                                                                v65 = v65[1:len(v65)]
                                                                            }
                                                                        }
                                                                    default:
                                                                        v162 := GenS1{F0: 2}
                                                                        _ = v162
                                                                    }
                                                                    v163 := GenS15{F0: ("hello" + "hello"), F1: (^(v7))}
                                                                    switch fi55 {
                                                                    case 127:
                                                                        p1 = (-(p1))
                                                                    case -51:
                                                                        var v164 uint64 = 4294967295
                                                                        _ = v164
                                                                    default:
                                                                        fmt.Println(fi55, p1, v57)
                                                                    }
                                                                    _ = v163
                                                                } else {
                                                                    switch (fi55 >> ((v7) & 63)) {
                                                                    case 1:
                                                                        if (false && ((^(v7)) < 42)) {
                                                                            _ = len(v65)
                                                                            v56 = v56
                                                                            fmt.Println(p0)
                                                                            v165 := GenS17{F0: p1}
                                                                            v7 += v7
                                                                            var v166 uint = v7
                                                                            if (100 >= -1) {
                                                                                var v167 int16 = -104
                                                                                var v168 uint8 = 209
                                                                                _ = v167
                                                                                _ = v168
                                                                            }
                                                                            switch (v63 / ((42) | 1)) {
                                                                            case 32767:
                                                                                switch ("\n\t" + (("hello" + "hello") + "\n\t")) {
                                                                                case "hello":
                                                                                    var v169 int16 = 1
                                                                                    _ = v169
                                                                                case "foo bar":
                                                                                    if len(v65) >= 2 {
                                                                                        v65 = v65[1:len(v65)]
                                                                                    }
                                                                                default:
                                                                                    v65 = append(v65, (fi42 + fi42))
                                                                                }
                                                                            case 0:
                                                                                if (float32((1e10 / (3.14159 * float64((3.14159 - ((1e10 - (v56 + v56)) - (float32(-2.25) / (v56 - 0.1)))))))) < (v56 + v56)) {
                                                                                    v170 := GenS6{F0: 128, F1: p0}
                                                                                    v7 += 107
                                                                                    _ = v170
                                                                                } else {
                                                                                    _ = len(v65)
                                                                                }
                                                                            case -1:
                                                                                var v171 uint16 = 65535
                                                                                _ = v171
                                                                            default:
                                                                                v172 := GenS20{F0: v7, F1: v166}
                                                                                _ = v172
                                                                            }
                                                                            fmt.Println(p0)
                                                                            if (p0 > (p0 / (((^(p0))) | 1))) {
                                                                                if false {
                                                                                    fmt.Println(v63, fi55)
                                                                                    var v173 int64 = (p1 + 9223372036854775807)
                                                                                    v174 := GenS21{F0: 574, F1: 18446744073709551615}
                                                                                    var v175 float32 = (v56 / (0.1 - 3.14159))
                                                                                    var v176 int8 = p0
                                                                                    _ = v173
                                                                                    _ = v174
                                                                                    _ = v175
                                                                                    _ = v176
                                                                                } else {
                                                                                    var v177 bool = (false && true)
                                                                                    switch 32767 {
                                                                                    case -1:
                                                                                        if len(v65) >= 2 {
                                                                                            v65 = v65[1:len(v65)]
                                                                                        }
                                                                                    default:
                                                                                        switch ("" + "") {
                                                                                        case "hello":
                                                                                            v56 -= float32(0.1)
                                                                                        case "":
                                                                                            p1 = p1
                                                                                        case "foo bar":
                                                                                            var v178 int16 = 608
                                                                                            _ = v178
                                                                                        }
                                                                                    }
                                                                                    var v179 float32 = v56
                                                                                    if (true || v177) {
                                                                                        if (42 >= v7) {
                                                                                            switch "hello" {
                                                                                            case "\n\t":
                                                                                                var v180 int16 = 817
                                                                                                _ = v180
                                                                                            }
                                                                                            if (("hello" + "\n\t") >= "a") {
                                                                                                var v181 uint64 = 2
                                                                                                _ = len(v65)
                                                                                                _ = v181
                                                                                            } else {
                                                                                                var v182 string = ("" + ("hello" + "a"))
                                                                                                v166 = 2
                                                                                                var v183 int = v63
                                                                                                var v184 uint32 = 428
                                                                                                _ = v182
                                                                                                _ = v183
                                                                                                _ = v184
                                                                                            }
                                                                                        }
                                                                                        v185 := GenS22{F0: v177}
                                                                                        _ = v185
                                                                                    }
                                                                                    v65 = append(v65, 4294967295)
                                                                                    _ = v177
                                                                                    _ = v179
                                                                                }
                                                                                var v186 uint8 = 255
                                                                                v63 = 127
                                                                                var v187 bool = true
                                                                                _ = v186
                                                                                _ = v187
                                                                            } else {
                                                                                var v188 uint = (v166 ^ 18446744073709551615)
                                                                                var v189 int64 = (p1 ^ p1)
                                                                                var v190 uint = v188
                                                                                _ = v188
                                                                                _ = v189
                                                                                _ = v190
                                                                            }
                                                                            var v191 uint16 = 407
                                                                            _ = v165
                                                                            _ = v166
                                                                            _ = v191
                                                                        } else {
                                                                            var v192 float64 = (1.5 * 1e10)
                                                                            _ = len(v65)
                                                                            _ = v192
                                                                        }
                                                                    default:
                                                                        v193 := GenS23{F0: (v7 - (v7 >> ((v7) & 63))), F1: 4294967295}
                                                                        _ = v193
                                                                    }
                                                                    v56 += ((v56 / 1e10) - (0.0 / v56))
                                                                }
                                                            } else {
                                                                v63 = fi55
                                                                v63 += (v63 % (((fi55 / (((-(fi42))) | 1))) | 1))
                                                            }
                                                            var v194 string = ((("hello" + "\n\t") + ((("" + (("hello" + "hello") + "foo bar")) + "a") + "hello")) + "foo bar")
                                                            p0 -= 42
                                                            fmt.Println(v56, v63)
                                                            var v195 int16 = -1
                                                            var v196 string = v194
                                                            if len(v65) >= 2 {
                                                                v65 = v65[1:len(v65)]
                                                            }
                                                            _ = v194
                                                            _ = v195
                                                            _ = v196
                                                        }
                                                        var v197 int8 = p0
                                                        _ = v197
                                                    }
                                                }
                                                v198 := GenS6{F0: 164, F1: (p0 * (p0 | p0))}
                                                _ = v198
                                            }
                                        default:
                                            p0 = -128
                                        }
                                        var v199 float32 = 3.14159
                                        _ = v199
                                    }
                                    v7 -= v7
                                    _ = v65
                                }
                                _ = v63
                            }
                            if false {
                                fmt.Println(p0, p1)
                                switch "a" {
                                case "hello":
                                    p1 -= -913
                                case "":
                                    var v200 int64 = -409
                                    _ = v200
                                }
                            }
                        }
                    case "a":
                        v201 := GenS15{F0: "foo bar", F1: (v7 * (^(v7)))}
                        _ = v201
                    }
                    _ = fi55
                    _ = v56
                    _ = v57
                }
                p1 += p1
                _ = v54
            }
            if true {
                var v202 uint64 = 18446744073709551615
                _ = v202
            }
        }
        _ = fi42
    }
    v203 := GenS23{F0: (v7 - v7), F1: 0}
    _ = p0
    _ = p1
    _ = v203
    _ = v7
    return -1
}
func genFn1(p0 uint32, p1 int64, p2 uint64) int {
    v204 := []int{-2147483648, 9223372036854775807}
    var v205 float32 = 3.14159
    fmt.Println(v205)
    if (!true) {
        v206 := GenS24{F0: ("\n\t" + "a")}
        for fi207 := 0; fi207 < 7; fi207++ {
            switch fi207 {
            case 806:
                for fi208 := 0; fi208 < 8; fi208++ {
                    fmt.Println(fi208, p2)
                    p1 = p1
                    _ = fi208
                }
            case -922:
                switch (-(fi207)) {
                case 1:
                    v205 += v205
                }
            }
            var v209 uint64 = p2
            if len(v204) > 0 {
                v204[0] = int(p1)
            }
            _ = fi207
            _ = v209
        }
        if ("\n\t" == (("" + "\n\t") + "")) {
            p1 += p1
        } else {
            switch 4294967295 {
            case 861:
                if true {
                    v210 := GenS12{F0: -2147483648}
                    _ = v210
                }
            case 9223372036854775807:
                v211 := GenS20{F0: 549, F1: 0}
                _ = v211
            default:
                p0 += (p0 - p0)
            }
            var v212 float32 = v205
            v213 := GenS19{F0: false, F1: -767, F2: -505}
            _ = v212
            _ = v213
        }
        _ = v206
    } else {
        for fi214 := 0; fi214 < 18; fi214++ {
            p0 += 74
            var v215 string = ("\n\t" + "foo bar")
            _ = fi214
            _ = v215
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v204
    _ = v205
    return 1
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    v4 := []string{("foo bar" + "a"), ("\n\t" + "")}
    v5 := make(map[string]string)
    for fi216 := 0; fi216 < 1; fi216++ {
        v217 := GenS13{F0: 0}
        for fi218 := 0; fi218 < 20; fi218++ {
            var v219 int16 = 21
            v220 := GenS25{F0: "", F1: 0}
            v219 = v219
            _ = fi218
            _ = v219
            _ = v220
        }
        _ = fi216
        _ = v217
    }
    if (true && false) {
        v3[2] = "hello"
        v221 := GenS13{F0: 181}
        _ = v221
    }
    if false {
        v1 = v1
        if true {
            switch -9223372036854775808 {
            case 128:
                _ = len(v2)
            default:
                for fi222 := 0; fi222 < 9; fi222++ {
                    for fi223 := 0; fi223 < 9; fi223++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi223
                    }
                    var v224 uint8 = 127
                    for fi225 := 0; fi225 < 11; fi225++ {
                        _ = len(v2)
                        _ = fi225
                    }
                    _ = fi222
                    _ = v224
                }
            }
            for fi226 := 0; fi226 < 9; fi226++ {
                if (!(!false)) {
                    v1 -= v1
                } else {
                    for fi227 := 0; fi227 < 19; fi227++ {
                        v1 += ((v1 - -2.25) + v1)
                        _ = fi227
                    }
                }
                v228 := v2[("a" + ("foo bar" + (("" + "a") + "a")))]
                v1 = 0.0
                _ = fi226
                _ = v228
            }
        }
        var v229 int = -167
        _ = v229
    } else {
        switch "foo bar" {
        case "hello":
            var v230 uint = 4294967295
            _ = v230
        case "\n\t":
            for fi231 := 0; fi231 < 19; fi231++ {
                fmt.Println(fi231)
                _ = fi231
            }
        case "a":
            switch -269 {
            case 193:
                switch 2147483647 {
                case 2147483647:
                    delete(v2, "")
                case 619:
                    if (1 >= 1) {
                        var v232 string = ("" + "")
                        v3[4294967295] = v232
                        switch -1 {
                        case 883:
                            for fi233 := 0; fi233 < 8; fi233++ {
                                fmt.Println(v1, v1, v1)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = fi233
                            }
                        default:
                            var v234 float32 = 0.0
                            _ = v234
                        }
                        _ = v232
                    }
                }
            case 127:
                if false {
                    var v235 bool = (!(v1 <= 0.1))
                    var v236 string = ("" + "foo bar")
                    fmt.Println(v1)
                    _ = v235
                    _ = v236
                }
            default:
                var v237 string = "\n\t"
                _ = v237
            }
        default:
            v238 := GenS15{F0: "hello", F1: 510}
            _ = v238
        }
        switch "a" {
        case "":
            fmt.Println(v1)
        case "hello":
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        case "\n\t":
            var v239 int16 = -271
            _ = v239
        default:
            var v240 bool = (("foo bar" <= "a") && false)
            _ = v240
        }
        var v241 int8 = 127
        _ = v241
    }
    fmt.Println(v1)
    var v242 int = -1
    if (true || true) {
        for fi243 := 0; fi243 < 2; fi243++ {
            v1 -= -2.25
            _ = fi243
        }
        v5["foo bar"] = "\n\t"
        v1 = (0.0 + (-2.25 + -2.25))
    } else {
        v242 -= (v242 / (((v242 / (((v242 / ((genFn1(65535, 0, 2)) | 1))) | 1))) | 1))
        if false {
            v244 := GenS1{F0: 763}
            v1 -= 0.1
            _ = v244
        }
        v1 = v1
    }
    if false {
        var v245 bool = (-632 == -1)
        v1 = v1
        switch "foo bar" {
        case "\n\t":
            _ = len(v4)
        case "foo bar":
            v246 := GenS2{F0: 4294967295}
            _ = v246
        case "a":
            if ((true || ((!v245) || (1 > 775))) && (genFn1(847, 9223372036854775807, 4294967295) != (v242 * (^(v242))))) {
                var v247 int8 = 1
                _ = v247
            } else {
                v3[2147483647] = "a"
            }
        default:
            for fi248 := 0; fi248 < 2; fi248++ {
                var v249 int64 = 42
                v249 -= v249
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi248
                _ = v249
            }
        }
        _ = v245
    }
    if (true || true) {
        if (2 > 65535) {
            v1 -= v1
            for fi250 := 0; fi250 < 7; fi250++ {
                fmt.Println(v242, v1, fi250)
                for fi251 := 0; fi251 < 7; fi251++ {
                    v252 := GenS26{F0: 0}
                    var v253 uint64 = 394
                    fmt.Println(v253)
                    _ = fi251
                    _ = v252
                    _ = v253
                }
                _ = fi250
            }
        }
    }
    fmt.Println(v242, v242)
    v254 := GenS1{F0: 4294967295}
    var v255 string = "\n\t"
    for fi256 := 0; fi256 < 16; fi256++ {
        if true {
            var v257 string = ((v255 + v255) + (v255 + v255))
            _ = v257
        } else {
            for fi258 := 0; fi258 < 6; fi258++ {
                var v259 int16 = -32768
                _ = fi258
                _ = v259
            }
        }
        for fi260 := 0; fi260 < 5; fi260++ {
            var v261 int8 = 65
            v261 = -1
            var v262 uint32 = 4294967295
            v242 = genFn1(v262, genFn0((v261 >> ((0) & 63)), 1), 0)
            _ = fi260
            _ = v261
            _ = v262
        }
        _ = fi256
    }
    for fi263 := 0; fi263 < 19; fi263++ {
        v1 += 3.14159
        v264 := GenS26{F0: -9223372036854775808}
        v255 = v255
        _ = fi263
        _ = v264
    }
    for fi265 := 0; fi265 < 2; fi265++ {
        v266 := GenS1{F0: 1}
        _ = fi265
        _ = v266
    }
    var v267 uint16 = 42
    if (42 >= -128) {
        if true {
            for fi268 := 0; fi268 < 5; fi268++ {
                fmt.Println(fi268, v255, v267)
                _ = fi268
            }
            if (!false) {
                var v269 int32 = 2147483647
                _ = v269
            }
            if (false || true) {
                v255 = v255
                v270 := GenS27{F0: v267, F1: 606, F2: 2}
                v4 = append(v4, ("\n\t" + (v255 + "a")))
                _ = v270
            }
        }
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        if len(v4) > 0 {
            v4[0] = v255
        }
    } else {
        v1 += 0.0
        switch 1 {
        case -667:
            var v271 uint32 = 856
            _ = v271
        case 0:
            var v272 string = v255
            _ = v272
        }
    }
    if (-1 <= 627) {
        v273 := GenS7{F0: (v267 | (^(v267))), F1: v255, F2: 1000}
        if false {
            if true {
                v274 := GenS12{F0: 0}
                _ = v274
            }
            var v275 int8 = 0
            _ = v275
        } else {
            v276 := GenS28{F0: 0}
            if ((1000 > -462) && false) {
                fmt.Println(v1, v267)
                switch 1 {
                case 5:
                    v267 = 255
                case 32767:
                    for fi277 := 0; fi277 < 1; fi277++ {
                        var v278 float32 = 1.5
                        switch v255 {
                        case "a":
                            var v279 int16 = 367
                            _ = v279
                        case "foo bar":
                            v1 -= ((float32(((3.14159 * 1.5) * 0.0)) + ((((v1 - v278) - (1.5 * v1)) / (v1 - (float32(0.1) + v1))) * ((v1 * 1.5) / v1))) / ((((-2.25 + v278) + v278) / v1) + v1))
                        }
                        v280 := GenS5{F0: 0, F1: 0}
                        _ = fi277
                        _ = v278
                        _ = v280
                    }
                default:
                    v1 += v1
                }
                for fi281 := 0; fi281 < 10; fi281++ {
                    delete(v3, 2147483647)
                    switch v242 {
                    case -357:
                        for fi282 := 0; fi282 < 20; fi282++ {
                            if (!((18446744073709551615 != 4294967295) || false)) {
                                v283 := GenS25{F0: (v255 + v255), F1: 181}
                                var v284 uint64 = 742
                                if true {
                                    v284 += v284
                                    delete(v2, "")
                                    v285 := v3[1]
                                    var v286 uint32 = 0
                                    if true {
                                        if false {
                                            var v287 int8 = 1
                                            var v288 bool = false
                                            var v289 uint64 = (v284 | (v284 + (v284 - 167)))
                                            _ = v287
                                            _ = v288
                                            _ = v289
                                        } else {
                                            if false {
                                                if true {
                                                    switch v285 {
                                                    case "hello":
                                                        var v290 uint64 = (v284 ^ v284)
                                                        _ = v290
                                                    default:
                                                        if (false || true) {
                                                            _ = len(v4)
                                                            var v291 uint = 0
                                                            v284 -= 0
                                                            _ = v291
                                                        }
                                                    }
                                                } else {
                                                    _ = len(v5)
                                                    var v292 float32 = v1
                                                    _ = v292
                                                }
                                                v293 := GenS19{F0: (((v267 * 255) < (v267 & (v267 & (v267 + 326)))) && (((v284 << ((565) & 63)) <= v284) || false)), F1: 42, F2: -1}
                                                if (0.1 > (-2.25 - 1e10)) {
                                                    fmt.Println(v285, v286)
                                                }
                                                var v294 uint8 = 152
                                                _ = v293
                                                _ = v294
                                            }
                                            var v295 int32 = 1
                                            fmt.Println(v286, v1)
                                            v1 = float32(0.1)
                                            _ = v295
                                        }
                                        var v296 uint64 = (v284 | uint64(v267))
                                        _ = v296
                                    } else {
                                        switch 655 {
                                        case 117:
                                            if (v284 > 364) {
                                                fmt.Println(v267, v242, v255)
                                                switch v255 {
                                                case "hello":
                                                    switch -30 {
                                                    case -1:
                                                        if (-1 >= 41) {
                                                            if (!(((9223372036854775807 == 9223372036854775807) || false) || ((!false) || (249 <= v267)))) {
                                                                var v297 int8 = int8(v286)
                                                                fmt.Println(v242, fi282, v297)
                                                                v298 := GenS3{F0: 1, F1: v286, F2: (v267 ^ v267)}
                                                                _ = v297
                                                                _ = v298
                                                            } else {
                                                                v267 += v267
                                                                v299 := GenS17{F0: 9223372036854775807}
                                                                var v300 uint = 1
                                                                _ = v299
                                                                _ = v300
                                                            }
                                                            _ = len(v5)
                                                        }
                                                    case -573:
                                                        var v301 float32 = 1e10
                                                        _ = v301
                                                    case 277:
                                                        delete(v2, ("a" + "hello"))
                                                    }
                                                default:
                                                    v286 -= v286
                                                }
                                            } else {
                                                if true {
                                                    v302 := GenS29{F0: 709}
                                                    v285 = (("\n\t" + "hello") + v285)
                                                    _ = v302
                                                }
                                                v1 = v1
                                            }
                                        case 4294967295:
                                            var v303 uint32 = 4294967295
                                            _ = v303
                                        default:
                                            fmt.Println(fi282, v284, fi281)
                                            var v304 string = ""
                                            _ = v304
                                        }
                                        var v305 uint8 = 98
                                        _ = v305
                                    }
                                    _ = v285
                                    _ = v286
                                } else {
                                    switch (-(fi282)) {
                                    case 32767:
                                        _ = len(v4)
                                    }
                                    v242 += fi281
                                }
                                _ = v283
                                _ = v284
                            } else {
                                var v306 bool = false
                                v306 = false
                                var v307 int = 174
                                _ = v306
                                _ = v307
                            }
                            _ = len(v3)
                            v308 := GenS30{F0: -334, F1: 478}
                            _ = fi282
                            _ = v308
                        }
                    }
                    _ = fi281
                }
            }
            fmt.Println(v1, v242, v242)
            _ = v276
        }
        _ = v273
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v242:", v242)
    fmt.Println("v255:", v255)
    fmt.Println("v267:", v267)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v242
    _ = v254
    _ = v255
    _ = v267
    _ = v3
    _ = v4
    _ = v5
}
