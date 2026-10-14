package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int; F2 int16 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 string; F1 int64; F2 int32 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 uint32; F1 int; F2 uint }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64; F1 int16; F2 uint }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int32; F1 int32 }
type GenS13 struct { F0 bool; F1 int }
type GenS14 struct { F0 string; F1 bool; F2 bool }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint16 }

func main() {
    var v1 float32 = 0.1
    v2 := []string{"foo bar", "\n\t"}
    v3 := make(map[int]int)
    v4 := []string{"a", (("hello" + "foo bar") + "hello"), "", "hello"}
    v5 := make(map[uint32]string)
    var v6 uint64 = 4294967295
    switch "" {
    case "foo bar":
        var v7 float32 = ((v1 * ((float32(1.5) * 1e10) * float32(3.14159))) - 3.14159)
        _ = v7
    case "\n\t":
        if false {
            v1 += (v1 / float32((0.0 + 3.14159)))
            fmt.Println(v1, v6)
        } else {
            var v8 int = 4294967295
            _ = v8
        }
    case "hello":
        fmt.Println(v6)
    default:
        var v9 int64 = -942
        _ = v9
    }
    v10 := GenS1{F0: 496}
    v11 := GenS1{F0: 361}
    switch "hello" {
    case "foo bar":
        if true {
            v1 = v1
            switch "hello" {
            case "hello":
                switch ("" + "\n\t") {
                case "\n\t":
                    if (true || true) {
                        var v12 int8 = 0
                        _ = v12
                    } else {
                        if false {
                            for fi13 := 0; fi13 < 13; fi13++ {
                                var v14 int8 = 1
                                for fi15 := 0; fi15 < 11; fi15++ {
                                    v1 = 1.5
                                    var v16 uint64 = 197
                                    switch "" {
                                    case "":
                                        var v17 int32 = 699
                                        _ = v17
                                    case "a":
                                        var v18 uint64 = v16
                                        _ = v18
                                    case "\n\t":
                                        v3[fi13] = (fi15 - (fi13 ^ fi13))
                                    default:
                                        v14 -= (v14 + int8(fi13))
                                    }
                                    v19 := GenS1{F0: 994}
                                    v20 := GenS1{F0: uint(v6)}
                                    _ = fi15
                                    _ = v16
                                    _ = v19
                                    _ = v20
                                }
                                _ = fi13
                                _ = v14
                            }
                        }
                        if false {
                            for fi21 := 0; fi21 < 6; fi21++ {
                                v6 = v6
                                v6 -= v6
                                var v22 int8 = -95
                                _ = fi21
                                _ = v22
                            }
                            v6 -= v6
                            v2 = append(v2, (("" + "") + "\n\t"))
                        } else {
                            switch -449 {
                            case 128:
                                v23 := GenS1{F0: 18446744073709551615}
                                _ = v23
                            case 32767:
                                v1 = v1
                            }
                            for fi24 := 0; fi24 < 14; fi24++ {
                                if false {
                                    v6 = v6
                                    var v25 bool = true
                                    var v26 uint8 = 2
                                    _ = v25
                                    _ = v26
                                }
                                fmt.Println(fi24, fi24, fi24)
                                _ = len(v2)
                                _ = fi24
                            }
                            for fi27 := 0; fi27 < 12; fi27++ {
                                fmt.Println(fi27, v1, v1)
                                var v28 int64 = -1
                                v1 = (v1 * (float32(1.5) + v1))
                                _ = fi27
                                _ = v28
                            }
                        }
                        v1 -= v1
                    }
                default:
                    for fi29 := 0; fi29 < 18; fi29++ {
                        v30 := GenS1{F0: 0}
                        var v31 string = "\n\t"
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                }
            case "\n\t":
                var v32 int32 = 0
                _ = v32
            case "":
                if true {
                    if (884 == (^(v6))) {
                        var v33 int32 = 0
                        _ = v33
                    }
                    v34 := GenS1{F0: 18446744073709551615}
                    v1 += (-2.25 - 0.1)
                    _ = v34
                } else {
                    var v35 string = "\n\t"
                    _ = v35
                }
            default:
                for fi36 := 0; fi36 < 20; fi36++ {
                    v6 -= v6
                    _ = fi36
                }
            }
        }
    case "a":
        switch 0 {
        case -40:
            if (42 < -2147483648) {
                v1 = ((v1 / v1) / float32((3.14159 - (1e10 - -2.25))))
            } else {
                var v37 int32 = 2147483647
                v38 := v3[288]
                for fi39 := 0; fi39 < 11; fi39++ {
                    var v40 uint64 = (v6 + v6)
                    if true {
                        if false {
                            v37 += (v37 % (((v37 ^ v37)) | 1))
                            v38 -= (fi39 | v38)
                            for fi41 := 0; fi41 < 12; fi41++ {
                                var v42 uint16 = 255
                                if (!true) {
                                    if false {
                                        v43 := GenS2{F0: v40}
                                        _ = v43
                                    }
                                    v42 -= 0
                                    if (2 <= 1) {
                                        switch (fi41 << ((65535) & 63)) {
                                        case 42:
                                            v37 += 1
                                        case -9223372036854775808:
                                            if (-2147483648 == v37) {
                                                switch fi41 {
                                                case -949:
                                                    if false {
                                                        v1 = float32(1e10)
                                                        v44 := GenS3{F0: (v40 % (((^(v6))) | 1)), F1: 785, F2: 1}
                                                        v5[2] = "foo bar"
                                                        _ = v44
                                                    }
                                                default:
                                                    v45 := GenS1{F0: 1}
                                                    _ = v45
                                                }
                                                if false {
                                                    v42 = v42
                                                    if len(v2) > 0 {
                                                        v2[0] = "hello"
                                                    }
                                                } else {
                                                    switch -2147483648 {
                                                    case 945:
                                                        if true {
                                                            if ("" > "a") {
                                                                fmt.Println(v37, v40, v42)
                                                            }
                                                        }
                                                    }
                                                    var v46 bool = (false && (0.0 > (((0.1 * 0.1) - float64(1.5)) / 1.0)))
                                                    v47 := GenS2{F0: (^(v6))}
                                                    _ = v46
                                                    _ = v47
                                                }
                                            }
                                        }
                                        if (39 > 1000) {
                                            var v48 int32 = (v37 ^ 2147483647)
                                            _ = v48
                                        } else {
                                            if (false && true) {
                                                v38 = (v38 % (((fi41 | (v38 + (fi39 >> ((42) & 63))))) | 1))
                                            } else {
                                                if false {
                                                    var v49 int = (v38 / ((-788) | 1))
                                                    var v50 uint = uint(v42)
                                                    if true {
                                                        switch v49 {
                                                        case 711:
                                                            v51 := GenS3{F0: 4294967295, F1: (fi41 % (((v38 ^ (fi41 | v38))) | 1)), F2: -1}
                                                            _ = v51
                                                        case -982:
                                                            if (-1 >= 564) {
                                                                var v52 bool = (false || false)
                                                                _ = v52
                                                            }
                                                        case 175:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        if true {
                                                            _ = len(v5)
                                                        }
                                                        var v53 uint32 = 2147483648
                                                        _ = v53
                                                    }
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    delete(v3, v38)
                                                    var v54 float64 = ((((((1e10 * 1e10) / (((-2.25 - 0.1) * 1e10) - 1e10)) * 0.1) - 0.1) + 0.0) - -2.25)
                                                    _ = v54
                                                }
                                                fmt.Println(v42)
                                            }
                                            var v55 bool = true
                                            v38 = -9223372036854775808
                                            _ = v55
                                        }
                                        var v56 uint32 = 4294967295
                                        _ = v56
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        switch (v38 / (((fi41 % (((fi41 * fi41)) | 1))) | 1)) {
                                        case -1:
                                            var v57 int64 = 9223372036854775807
                                            _ = v57
                                        case 9223372036854775807:
                                            var v58 uint16 = 2
                                            _ = v58
                                        }
                                    }
                                } else {
                                    v42 -= 255
                                    switch fi41 {
                                    case 215:
                                        if ((("hello" + "\n\t") + "a") == "hello") {
                                            switch fi41 {
                                            case 127:
                                                if true {
                                                    var v59 uint = 18446744073709551615
                                                    v1 = (v1 + (-2.25 * 3.14159))
                                                    if false {
                                                        fmt.Println(v59)
                                                        v6 = (v40 ^ v6)
                                                        v37 = 979
                                                    }
                                                    v38 -= 42
                                                    _ = v59
                                                } else {
                                                    var v60 bool = true
                                                    var v61 uint8 = 42
                                                    var v62 int16 = -1
                                                    var v63 int = 827
                                                    var v64 float32 = v1
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                }
                                            default:
                                                v65 := GenS2{F0: 1}
                                                _ = v65
                                            }
                                        }
                                    case 4294967295:
                                        v6 += v40
                                    default:
                                        if true {
                                            v38 -= (fi39 - -828)
                                            switch v38 {
                                            case 127:
                                                _ = len(v4)
                                            case 9223372036854775807:
                                                var v66 string = ("foo bar" + "foo bar")
                                                _ = v66
                                            }
                                        }
                                    }
                                    v1 += (v1 * 0.0)
                                }
                                _ = fi41
                                _ = v42
                            }
                        }
                        var v67 float32 = v1
                        _ = v67
                    } else {
                        v1 = float32(3.14159)
                        switch "foo bar" {
                        case "foo bar":
                            delete(v5, 656)
                        }
                    }
                    for fi68 := 0; fi68 < 7; fi68++ {
                        var v69 int64 = 42
                        var v70 string = "foo bar"
                        _ = fi68
                        _ = v69
                        _ = v70
                    }
                    _ = fi39
                    _ = v40
                }
                _ = v37
                _ = v38
            }
        case 901:
            fmt.Println(v1, v6)
        case 9223372036854775807:
            var v71 uint64 = 2
            _ = v71
        default:
            for fi72 := 0; fi72 < 8; fi72++ {
                delete(v5, 1)
                v6 = v6
                for fi73 := 0; fi73 < 3; fi73++ {
                    var v74 int = fi72
                    v75 := v5[2]
                    if true {
                        var v76 float32 = (((float32(1e10) + (0.1 / (v1 + v1))) / float32((1.5 * 1.5))) + (v1 + (v1 + (v1 - 0.1))))
                        v6 = v6
                        _ = v76
                    }
                    _ = fi73
                    _ = v74
                    _ = v75
                }
                _ = fi72
            }
        }
    case "\n\t":
        v1 = v1
    }
    var v77 uint8 = 255
    if len(v4) > 0 {
        v4[0] = ("" + "")
    }
    for fi78 := 0; fi78 < 7; fi78++ {
        if false {
            fmt.Println(v6)
            switch ("a" + "") {
            case "":
                switch ("hello" + ("" + "a")) {
                case "\n\t":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "a":
                    for fi79 := 0; fi79 < 5; fi79++ {
                        v80 := GenS3{F0: 42, F1: 274, F2: -573}
                        _ = fi79
                        _ = v80
                    }
                case "":
                    var v81 float32 = (v1 * 1e10)
                    _ = v81
                default:
                    switch "" {
                    case "a":
                        for fi82 := 0; fi82 < 10; fi82++ {
                            v6 = uint64(v77)
                            v6 = (^(v6))
                            _ = fi82
                        }
                    }
                }
            case "a":
                var v83 int64 = 1
                _ = v83
            case "foo bar":
                switch (("foo bar" + (("\n\t" + "") + ("\n\t" + "\n\t"))) + ((("" + "foo bar") + "") + "hello")) {
                case "hello":
                    for fi84 := 0; fi84 < 15; fi84++ {
                        v6 -= (^(v6))
                        v77 = (v77 & (v77 / (((v77 + v77)) | 1)))
                        _ = fi84
                    }
                case "a":
                    v1 = 0.0
                default:
                    if false {
                        var v85 uint16 = 65535
                        for fi86 := 0; fi86 < 14; fi86++ {
                            if (((("foo bar" + "a") + "a") >= "foo bar") || false) {
                                v87 := GenS2{F0: (v6 % ((42) | 1))}
                                if true {
                                    v77 -= 255
                                    v85 -= (v85 & (v85 + uint16(v77)))
                                    if false {
                                        var v88 uint8 = 0
                                        v89 := GenS4{F0: 65535, F1: -257}
                                        _ = v88
                                        _ = v89
                                    }
                                } else {
                                    if (65535 <= 307) {
                                        v90 := GenS5{F0: "a", F1: 141, F2: 1}
                                        var v91 string = "a"
                                        var v92 uint16 = 806
                                        _ = v90
                                        _ = v91
                                        _ = v92
                                    } else {
                                        fmt.Println(fi78, v85)
                                        delete(v3, fi86)
                                        if true {
                                            fmt.Println(v6, v77, v77)
                                            var v93 uint8 = (v77 | (v77 >> ((255) & 63)))
                                            _ = v93
                                        } else {
                                            if false {
                                                v6 = 1
                                                v77 -= 255
                                                switch "hello" {
                                                case "\n\t":
                                                    if false {
                                                        v6 = (^(v6))
                                                        switch 9223372036854775807 {
                                                        case -827:
                                                            var v94 int16 = 32767
                                                            _ = v94
                                                        default:
                                                            v95 := GenS6{F0: (v77 - v77), F1: 1}
                                                            _ = v95
                                                        }
                                                    }
                                                case "foo bar":
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v2 = append(v2, "")
                                                    case "":
                                                        switch "a" {
                                                        case "foo bar":
                                                            switch "hello" {
                                                            case "hello":
                                                                var v96 uint64 = (v6 << ((875) & 63))
                                                                _ = v96
                                                            default:
                                                                if (!true) {
                                                                    v3[fi86] = fi86
                                                                } else {
                                                                    var v97 int64 = 42
                                                                    v85 -= 1
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v98 float64 = ((float64((v1 * ((v1 / (((((v1 * (v1 + 0.1)) - float32(float64(v1))) + (3.14159 - v1)) - 1e10) / (1e10 - 0.1))) / -2.25))) + 1.5) * (1.5 - 3.14159))
                                                                        _ = v98
                                                                    default:
                                                                        v2 = append(v2, "hello")
                                                                    }
                                                                    _ = v97
                                                                }
                                                            }
                                                        case "\n\t":
                                                            switch (("hello" + "") + "\n\t") {
                                                            case "foo bar":
                                                                v99 := GenS7{F0: 4294967295, F1: fi78, F2: 42}
                                                                _ = v99
                                                            case "\n\t":
                                                                v1 = (v1 * float32((((0.1 / 0.1) - float64(((1.5 / (((-2.25 / v1) * v1) / (v1 / v1))) + 0.1))) / float64(1e10))))
                                                            case "hello":
                                                                v100 := GenS6{F0: (^(v77)), F1: 397}
                                                                _ = v100
                                                            default:
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    if false {
                                                                        var v101 float32 = v1
                                                                        _ = v101
                                                                    } else {
                                                                        v77 = (^(v77))
                                                                        fmt.Println(v6, v6, v1)
                                                                        var v102 uint64 = uint64(v85)
                                                                        _ = v102
                                                                    }
                                                                }
                                                            }
                                                            if ("\n\t" >= "hello") {
                                                                switch "\n\t" {
                                                                case "a":
                                                                    switch 32767 {
                                                                    case -113:
                                                                        if (true && true) {
                                                                            fmt.Println(fi78)
                                                                            var v103 int64 = 42
                                                                            if (4294967295 != 679) {
                                                                                var v104 int32 = 1
                                                                                _ = v104
                                                                            } else {
                                                                                var v105 int64 = (v103 / ((-1) | 1))
                                                                                switch "foo bar" {
                                                                                case "a":
                                                                                    switch (fi78 - (fi78 - fi78)) {
                                                                                    case 32768:
                                                                                        v105 += 9223372036854775807
                                                                                    case 319:
                                                                                        fmt.Println(fi86, fi78)
                                                                                        if (!((v103 | -1) <= (v103 >> ((uint(v85)) & 63)))) {
                                                                                            var v106 uint32 = 42
                                                                                            v107 := GenS8{F0: (fi78 ^ fi78)}
                                                                                            _ = v106
                                                                                            _ = v107
                                                                                        }
                                                                                    case -1:
                                                                                        var v108 uint16 = v85
                                                                                        _ = v108
                                                                                    }
                                                                                }
                                                                                switch (("hello" + ("foo bar" + "a")) + ("\n\t" + "")) {
                                                                                case "":
                                                                                    v109 := GenS4{F0: v85, F1: -665}
                                                                                    _ = v109
                                                                                case "\n\t":
                                                                                    if (!((("" + "a") >= "foo bar") || (-2147483648 < -136))) {
                                                                                        var v110 int16 = 923
                                                                                        fmt.Println(v85, v103)
                                                                                        _ = v110
                                                                                    } else {
                                                                                        if true {
                                                                                            var v111 int16 = -32768
                                                                                            var v112 int = fi86
                                                                                            if true {
                                                                                                v113 := GenS3{F0: v6, F1: (fi86 * fi78), F2: 601}
                                                                                                v77 += v77
                                                                                                _ = v113
                                                                                            }
                                                                                            _ = v111
                                                                                            _ = v112
                                                                                        }
                                                                                        if (4294967295 < 2) {
                                                                                            var v114 uint = 0
                                                                                            if (!true) {
                                                                                                v103 += v103
                                                                                                switch fi78 {
                                                                                                case 4294967295:
                                                                                                    v6 = (v6 * (v6 + uint64(v114)))
                                                                                                case -1:
                                                                                                    if true {
                                                                                                        v114 -= (v114 + (v114 ^ 1))
                                                                                                    }
                                                                                                case 1:
                                                                                                    switch (-(fi78)) {
                                                                                                    case 1:
                                                                                                        if len(v4) > 0 {
                                                                                                            v4[0] = "hello"
                                                                                                        }
                                                                                                    case -442:
                                                                                                        fmt.Println(v1, v1, v77)
                                                                                                    default:
                                                                                                        v6 = v6
                                                                                                    }
                                                                                                default:
                                                                                                    if ((false && true) && (true && false)) {
                                                                                                        if false {
                                                                                                            v85 = (v85 >> ((2) & 63))
                                                                                                            v114 += (v114 & uint(v85))
                                                                                                            v115 := v3[32767]
                                                                                                            var v116 int32 = -67
                                                                                                            fmt.Println(v85)
                                                                                                            _ = v115
                                                                                                            _ = v116
                                                                                                        }
                                                                                                        var v117 int8 = -39
                                                                                                        if false {
                                                                                                            v105 += int64(v77)
                                                                                                            if (42 == 0) {
                                                                                                                v85 = (v85 | v85)
                                                                                                            }
                                                                                                            switch ("" + ("\n\t" + "")) {
                                                                                                            case "\n\t":
                                                                                                                fmt.Println(v114)
                                                                                                            case "hello":
                                                                                                                fmt.Println(v117, v85, v1)
                                                                                                                if false {
                                                                                                                    _ = len(v2)
                                                                                                                    delete(v3, 773)
                                                                                                                    v6 += v6
                                                                                                                } else {
                                                                                                                    var v118 bool = true
                                                                                                                    v119 := GenS9{F0: (v6 % ((1) | 1)), F1: 708, F2: v114}
                                                                                                                    v1 = v1
                                                                                                                    _ = v118
                                                                                                                    _ = v119
                                                                                                                }
                                                                                                                if ((fi86 * fi78) < 32768) {
                                                                                                                    if false {
                                                                                                                        var v120 int16 = 1
                                                                                                                        switch fi86 {
                                                                                                                        case 2147483647:
                                                                                                                            var v121 uint64 = 42
                                                                                                                            _ = v121
                                                                                                                        case 488:
                                                                                                                            v5[1] = ("" + "foo bar")
                                                                                                                        }
                                                                                                                        var v122 int8 = (^(v117))
                                                                                                                        _ = v120
                                                                                                                        _ = v122
                                                                                                                    } else {
                                                                                                                        var v123 int = fi86
                                                                                                                        v6 = v6
                                                                                                                        fmt.Println(v123, v117)
                                                                                                                        _ = v123
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                v124 := GenS10{F0: v77}
                                                                                                                _ = v124
                                                                                                            }
                                                                                                        } else {
                                                                                                            v85 += (v85 >> (((v114 + v114)) & 63))
                                                                                                            fmt.Println(v117)
                                                                                                        }
                                                                                                        _ = v117
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v125 := GenS6{F0: 2, F1: uint32(v77)}
                                                                                            v126 := GenS7{F0: 464, F1: (-(fi86)), F2: (^(v114))}
                                                                                            _ = v114
                                                                                            _ = v125
                                                                                            _ = v126
                                                                                        }
                                                                                        if (!true) {
                                                                                            var v127 int32 = 2147483647
                                                                                            v105 = v103
                                                                                            _ = v127
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v105
                                                                            }
                                                                            _ = v103
                                                                        } else {
                                                                            var v128 uint32 = 2147483648
                                                                            v4 = append(v4, "a")
                                                                            if false {
                                                                                var v129 string = ""
                                                                                fmt.Println(fi86, fi78, v129)
                                                                                _ = v129
                                                                            }
                                                                            _ = v128
                                                                        }
                                                                    default:
                                                                        delete(v5, 2147483647)
                                                                    }
                                                                case "\n\t":
                                                                    v77 -= 1
                                                                }
                                                                if false {
                                                                    if true {
                                                                        if true {
                                                                            fmt.Println(v6, v77)
                                                                            var v130 uint64 = v6
                                                                            if false {
                                                                                if true {
                                                                                    if (!true) {
                                                                                        if (fi78 >= fi86) {
                                                                                            var v131 int = fi86
                                                                                            _ = len(v3)
                                                                                            v132 := GenS11{F0: true}
                                                                                            if (1.5 > 0.0) {
                                                                                                v133 := GenS6{F0: (v77 ^ v77), F1: 2147483648}
                                                                                                var v134 string = ("hello" + "foo bar")
                                                                                                _ = v133
                                                                                                _ = v134
                                                                                            }
                                                                                            var v135 bool = false
                                                                                            _ = v131
                                                                                            _ = v132
                                                                                            _ = v135
                                                                                        }
                                                                                        v1 -= (0.1 + v1)
                                                                                    } else {
                                                                                        if true {
                                                                                            if (!true) {
                                                                                                v130 = v130
                                                                                                v5[0] = "\n\t"
                                                                                                if (1.5 >= 3.14159) {
                                                                                                    var v136 uint8 = (v77 >> ((42) & 63))
                                                                                                    _ = v136
                                                                                                }
                                                                                            } else {
                                                                                                switch fi86 {
                                                                                                case -372:
                                                                                                    v137 := GenS12{F0: 1, F1: 964}
                                                                                                    _ = v137
                                                                                                case 2147483647:
                                                                                                    v6 += v130
                                                                                                case 49:
                                                                                                    fmt.Println(fi86, v77, v85)
                                                                                                }
                                                                                                switch "foo bar" {
                                                                                                case "":
                                                                                                    v138 := GenS4{F0: (v85 >> ((484) & 63)), F1: 42}
                                                                                                    _ = v138
                                                                                                case "hello":
                                                                                                    switch ("" + "\n\t") {
                                                                                                    case "":
                                                                                                        if len(v2) > 0 {
                                                                                                            v2[0] = ("hello" + ("a" + "foo bar"))
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        v85 = 65535
                                                                                                    }
                                                                                                default:
                                                                                                    v139 := GenS5{F0: "hello", F1: 0, F2: -639}
                                                                                                    _ = v139
                                                                                                }
                                                                                                switch fi86 {
                                                                                                case 1:
                                                                                                    if (9 >= -128) {
                                                                                                        v3[fi78] = fi86
                                                                                                        var v140 int32 = 179
                                                                                                        v141 := v5[42]
                                                                                                        fmt.Println(v1, v130, v85)
                                                                                                        var v142 uint8 = (^(v77))
                                                                                                        _ = v140
                                                                                                        _ = v141
                                                                                                        _ = v142
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v6 -= (^(v6))
                                                                                        } else {
                                                                                            var v143 int16 = 1
                                                                                            var v144 string = ""
                                                                                            _ = v143
                                                                                            _ = v144
                                                                                        }
                                                                                        var v145 int = (fi86 | (-(fi86)))
                                                                                        v130 = 18446744073709551615
                                                                                        _ = v145
                                                                                    }
                                                                                    v6 = v6
                                                                                }
                                                                                v2 = append(v2, "foo bar")
                                                                            } else {
                                                                                var v146 int8 = -1
                                                                                var v147 float64 = 1.5
                                                                                var v148 float64 = v147
                                                                                _ = v146
                                                                                _ = v147
                                                                                _ = v148
                                                                            }
                                                                            _ = v130
                                                                        } else {
                                                                            v149 := GenS6{F0: (^(v77)), F1: 2147483648}
                                                                            v1 -= (float32(0.0) - v1)
                                                                            fmt.Println(fi78, v1, v6)
                                                                            _ = v149
                                                                        }
                                                                    }
                                                                    v5[2] = ""
                                                                } else {
                                                                    if (42 > 42) {
                                                                        var v150 float64 = (0.1 + (3.14159 / (((3.14159 - 3.14159) - 1e10) - 0.0)))
                                                                        v1 += (v1 + (0.0 / ((v1 - v1) / 1.5)))
                                                                        _ = v150
                                                                    } else {
                                                                        fmt.Println(v6, v85, v85)
                                                                    }
                                                                    v1 = 1.5
                                                                }
                                                                v151 := GenS4{F0: (v85 + v85), F1: -895}
                                                                v77 = v77
                                                                _ = v151
                                                            } else {
                                                                v77 = v77
                                                                var v152 uint = 7
                                                                _ = v152
                                                            }
                                                            if true {
                                                                var v153 uint8 = (v77 & (v77 + v77))
                                                                var v154 float64 = 1e10
                                                                _ = v153
                                                                _ = v154
                                                            }
                                                        default:
                                                            var v155 int64 = -9223372036854775808
                                                            _ = v155
                                                        }
                                                    case "\n\t":
                                                        var v156 uint64 = v6
                                                        _ = v156
                                                    }
                                                    var v157 int8 = -70
                                                    _ = v157
                                                case "hello":
                                                    v1 += 0.1
                                                default:
                                                    v1 -= (0.0 - v1)
                                                }
                                            }
                                        }
                                    }
                                    var v158 string = ("\n\t" + "hello")
                                    v1 = 3.14159
                                    _ = v158
                                }
                                var v159 int = (fi86 ^ fi86)
                                _ = v159
                                _ = v87
                            }
                            v1 = v1
                            var v160 string = ""
                            _ = fi86
                            _ = v160
                        }
                        _ = v85
                    } else {
                        fmt.Println(v1, v1, fi78)
                        if true {
                            fmt.Println(v6, v77)
                            v77 = 37
                            switch "a" {
                            case "a":
                                for fi161 := 0; fi161 < 5; fi161++ {
                                    if (true && ((0.1 / -2.25) == 0.1)) {
                                        var v162 float64 = 1.5
                                        v163 := GenS13{F0: true, F1: fi78}
                                        _ = v162
                                        _ = v163
                                    }
                                    if (2147483647 == 971) {
                                        delete(v3, 873)
                                        v1 += v1
                                    } else {
                                        fmt.Println(v77, fi161)
                                    }
                                    _ = fi161
                                }
                            case "hello":
                                v164 := GenS14{F0: ("hello" + ""), F1: ((v1 + v1) < (v1 - v1)), F2: false}
                                _ = v164
                            }
                        } else {
                            v165 := GenS12{F0: 0, F1: -238}
                            var v166 uint32 = 2147483648
                            _ = len(v2)
                            _ = v165
                            _ = v166
                        }
                        v3[2147483647] = (fi78 - fi78)
                    }
                }
            default:
                delete(v3, (fi78 % (((fi78 - fi78)) | 1)))
            }
            if false {
                var v167 bool = ((127 <= 121) || true)
                if true {
                    if true {
                        switch "foo bar" {
                        case "":
                            switch fi78 {
                            case 127:
                                v168 := GenS15{F0: 1}
                                _ = v168
                            case 128:
                                var v169 int16 = 889
                                _ = v169
                            case 0:
                                var v170 uint16 = 65535
                                _ = v170
                            default:
                                for fi171 := 0; fi171 < 17; fi171++ {
                                    v6 += 705
                                    var v172 int16 = 1
                                    _ = fi171
                                    _ = v172
                                }
                            }
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                    for fi173 := 0; fi173 < 17; fi173++ {
                        switch (fi78 | 32768) {
                        case 0:
                            var v174 uint = 746
                            _ = v174
                        case -2147483648:
                            var v175 int8 = int8(v6)
                            _ = v175
                        }
                        v176 := GenS12{F0: 1, F1: -302}
                        var v177 uint8 = v77
                        _ = fi173
                        _ = v176
                        _ = v177
                    }
                    if v167 {
                        v167 = (!v167)
                        v178 := GenS1{F0: 1}
                        _ = v178
                    } else {
                        for fi179 := 0; fi179 < 17; fi179++ {
                            var v180 int64 = -769
                            if ((("a" + "hello") + ("\n\t" + "foo bar")) != ("hello" + "a")) {
                                _ = len(v4)
                                fmt.Println(v1, v167, v6)
                            } else {
                                v1 -= v1
                                v181 := GenS16{F0: 167}
                                _ = v181
                            }
                            if v167 {
                                var v182 uint32 = 1
                                switch (fi78 * (^(fi179))) {
                                case -2147483648:
                                    fmt.Println(v77)
                                case -1:
                                    v180 = 0
                                case 0:
                                    var v183 int32 = 1
                                    _ = v183
                                }
                                _ = v182
                            }
                            delete(v3, fi78)
                            _ = fi179
                            _ = v180
                        }
                    }
                } else {
                    v6 -= (v6 - 888)
                    v167 = (("hello" + "hello") <= ("\n\t" + "a"))
                }
                _ = v167
            }
        }
        _ = fi78
    }
    if false {
        var v184 bool = true
        if (1 == 4294967295) {
            for fi185 := 0; fi185 < 14; fi185++ {
                v184 = false
                v77 = (v77 / ((0) | 1))
                v186 := GenS17{F0: 65535}
                _ = fi185
                _ = v186
            }
            for fi187 := 0; fi187 < 4; fi187++ {
                var v188 int64 = 864
                _ = fi187
                _ = v188
            }
        } else {
            for fi189 := 0; fi189 < 7; fi189++ {
                delete(v3, -132)
                var v190 uint = 556
                _ = fi189
                _ = v190
            }
            _ = len(v5)
            fmt.Println(v1, v6)
        }
        _ = v184
    } else {
        var v191 int16 = -445
        _ = v191
    }
    var v192 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v192:", v192)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v10
    _ = v11
    _ = v192
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
}
