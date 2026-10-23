package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int8; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS6 struct { F0 int32; F1 uint32; F2 int }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 string; F1 uint8; F2 int }
type GenS11 struct { F0 string; F1 uint8; F2 int32 }
type GenS12 struct { F0 bool; F1 uint; F2 uint64 }
type GenS13 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS14 struct { F0 int8; F1 bool }
type GenS15 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS16 struct { F0 int32; F1 uint16; F2 int32 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 int32; F1 uint32; F2 bool }
type GenS19 struct { F0 uint32; F1 uint }
type GenS20 struct { F0 int8; F1 int64; F2 uint64 }
type GenS21 struct { F0 int }
type GenS22 struct { F0 string; F1 string }
type GenS23 struct { F0 uint32; F1 int32; F2 int64 }
type GenS24 struct { F0 int64 }
type GenS25 struct { F0 uint }
type GenS26 struct { F0 int; F1 int; F2 uint64 }

func main() {
    var v1 uint64 = 41
    var v2 uint = 65535
    var v3 uint8 = 42
    var v4 int = 4294967295
    v5 := make(map[string]uint32)
    var v6 uint = uint(v3)
    var v7 int16 = 32767
    var v8 uint64 = 0
    v9 := GenS1{F0: v4, F1: v3, F2: v3}
    var v10 int16 = -1
    switch (("a" + "") + ("a" + "foo bar")) {
    case "a":
        fmt.Println(v3, v4)
    case "hello":
        v11 := GenS1{F0: v4, F1: (v3 * v3), F2: (^(v3))}
        _ = v11
    case "foo bar":
        v3 += (^(v3))
    default:
        if (0.0 <= (((1e10 + (3.14159 / -2.25)) - (1.5 * 1e10)) / 1.5)) {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 int64 = 42
                _ = fi12
                _ = v13
            }
            v14 := GenS2{F0: v7}
            var v15 int8 = 127
            _ = v14
            _ = v15
        } else {
            _ = len(v5)
        }
    }
    fmt.Println(v10)
    v16 := []int{(v4 - 4294967295), v4}
    switch ((("hello" + "foo bar") + "a") + "hello") {
    case "\n\t":
        delete(v5, "")
    case "":
        var v17 int16 = int16(v8)
        _ = v17
    case "hello":
        if true {
            if ((3.14159 - 0.1) == -2.25) {
                v18 := GenS3{F0: 32767, F1: ("a" + "foo bar")}
                _ = v18
            } else {
                if false {
                    v16 = append(v16, 1)
                    v2 += 4294967295
                    if true {
                        delete(v5, (("a" + ((("hello" + "foo bar") + ("a" + ("\n\t" + ("a" + "\n\t")))) + "")) + "foo bar"))
                        var v19 float32 = -2.25
                        v3 = 127
                        _ = v19
                    }
                } else {
                    fmt.Println(v1, v10)
                }
                var v20 bool = true
                v1 = (v1 >> ((0) & 63))
                _ = v20
            }
            if true {
                v21 := GenS1{F0: -2147483648, F1: 0, F2: v3}
                v22 := GenS4{F0: -128, F1: (v2 | v6), F2: 42}
                for fi23 := 0; fi23 < 9; fi23++ {
                    if ((-187 != -9223372036854775808) || ("a" <= (("hello" + ("\n\t" + "\n\t")) + ("a" + ("a" + "hello"))))) {
                        v8 = v1
                        switch (v4 + v4) {
                        case 9223372036854775807:
                            v24 := GenS5{F0: uint64(v3), F1: 2147483648, F2: (v3 + 25)}
                            _ = v24
                        case -2147483648:
                            if false {
                                v6 += v2
                            } else {
                                for fi25 := 0; fi25 < 5; fi25++ {
                                    switch ("a" + "hello") {
                                    case "hello":
                                        v4 = fi25
                                    default:
                                        delete(v5, ("hello" + ("a" + "hello")))
                                    }
                                    v26 := GenS4{F0: 1, F1: 1, F2: uint16(v6)}
                                    delete(v5, ("hello" + "\n\t"))
                                    _ = fi25
                                    _ = v26
                                }
                                var v27 bool = true
                                _ = v27
                            }
                        }
                    }
                    _ = fi23
                }
                _ = v21
                _ = v22
            } else {
                v28 := GenS6{F0: -1, F1: 857, F2: v4}
                if ((!true) && (!false)) {
                    _ = len(v5)
                    for fi29 := 0; fi29 < 1; fi29++ {
                        v30 := GenS7{F0: 42, F1: (v8 * (v8 | (^(v1))))}
                        var v31 uint32 = 1
                        v2 -= (v6 % ((v6) | 1))
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    fmt.Println(v1)
                } else {
                    switch 815 {
                    case -9223372036854775808:
                        switch "a" {
                        case "\n\t":
                            v32 := GenS5{F0: v8, F1: uint32(v3), F2: (^(v3))}
                            _ = v32
                        }
                    case 128:
                        fmt.Println(v2, v2)
                    case 32768:
                        var v33 int8 = 127
                        _ = v33
                    default:
                        var v34 int = v4
                        _ = v34
                    }
                    v35 := GenS6{F0: 2147483647, F1: 261, F2: -2147483648}
                    if (!true) {
                        if (1e10 < 0.1) {
                            v36 := v5[((("foo bar" + "hello") + "\n\t") + "a")]
                            switch ((("\n\t" + "hello") + "foo bar") + "foo bar") {
                            case "a":
                                if true {
                                    var v37 uint64 = (^(v8))
                                    switch (v4 << (((v2 | 255)) & 63)) {
                                    case -888:
                                        v4 = (v4 + (v4 | (v4 ^ (v4 / ((828) | 1)))))
                                    case 32768:
                                        v10 += (v10 - -200)
                                    case 9223372036854775807:
                                        v1 = 425
                                    default:
                                        if (!true) {
                                            v8 = uint64(v2)
                                            v5["a"] = 496
                                            v38 := GenS2{F0: v7}
                                            _ = v38
                                        } else {
                                            var v39 bool = ((!(!false)) || (!(3.14159 != 0.1)))
                                            _ = v39
                                        }
                                    }
                                    v7 -= v10
                                    _ = v37
                                } else {
                                    v3 = (v3 & (v3 ^ (^(v3))))
                                    var v40 int8 = 0
                                    switch "foo bar" {
                                    case "a":
                                        v41 := GenS6{F0: -1, F1: v36, F2: v4}
                                        _ = v41
                                    case "hello":
                                        if (1 == v40) {
                                            if ((-(v40)) == v40) {
                                                fmt.Println(v40)
                                                v6 = (v2 + v6)
                                                v40 = v40
                                            } else {
                                                for fi42 := 0; fi42 < 8; fi42++ {
                                                    var v43 uint = 18446744073709551615
                                                    v16 = append(v16, fi42)
                                                    v7 = v7
                                                    _ = fi42
                                                    _ = v43
                                                }
                                            }
                                        }
                                    case "":
                                        var v44 uint8 = (v3 << ((2) & 63))
                                        _ = v44
                                    }
                                    _ = v40
                                }
                            case "foo bar":
                                for fi45 := 0; fi45 < 12; fi45++ {
                                    _ = len(v5)
                                    _ = fi45
                                }
                            case "hello":
                                for fi46 := 0; fi46 < 8; fi46++ {
                                    switch "hello" {
                                    case "":
                                        if (false || false) {
                                            v7 = 139
                                            if true {
                                                for fi47 := 0; fi47 < 4; fi47++ {
                                                    if ((!(!true)) && true) {
                                                        var v48 int16 = -1
                                                        var v49 float32 = (0.0 + ((((-2.25 - 1.5) / (1e10 + (float32(1e10) * 1e10))) / 0.1) * ((-2.25 * 0.1) - (float32((float64(3.14159) / 0.1)) / (float32(3.14159) + -2.25)))))
                                                        switch 42 {
                                                        case 653:
                                                            if (2147483648 <= v36) {
                                                                v50 := GenS8{F0: -38, F1: -244}
                                                                var v51 int = (-(v4))
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                            v6 = v6
                                                            var v52 uint32 = (v36 ^ (^(v36)))
                                                            _ = v52
                                                        case -2147483648:
                                                            v4 -= v4
                                                            v3 -= (v3 / (((v3 & v3)) | 1))
                                                            fmt.Println(v10, v48)
                                                            var v53 bool = ((100 >= int8(v36)) || (1000 > v48))
                                                            _ = v53
                                                        case -9223372036854775808:
                                                            v48 += int16(v36)
                                                        }
                                                        v48 = v7
                                                        var v54 int64 = 47
                                                        _ = v48
                                                        _ = v49
                                                        _ = v54
                                                    } else {
                                                        v4 = (fi47 % (((v4 / ((-865) | 1))) | 1))
                                                        var v55 string = "a"
                                                        switch "\n\t" {
                                                        case "a":
                                                            v8 = v8
                                                        }
                                                        _ = v55
                                                    }
                                                    v56 := GenS9{F0: -25}
                                                    _ = fi47
                                                    _ = v56
                                                }
                                            } else {
                                                var v57 uint8 = v3
                                                _ = v57
                                            }
                                        } else {
                                            switch (^(v4)) {
                                            case 42:
                                                v6 += uint(v1)
                                            case 9223372036854775807:
                                                var v58 int8 = -128
                                                _ = v58
                                            default:
                                                if (!false) {
                                                    v16 = append(v16, fi46)
                                                    v2 = (v6 + (v2 ^ (v6 + (v2 | (^(v6))))))
                                                    if (true || false) {
                                                        _ = len(v5)
                                                        for fi59 := 0; fi59 < 6; fi59++ {
                                                            v60 := GenS7{F0: -460, F1: 1}
                                                            switch "hello" {
                                                            case "":
                                                                var v61 uint32 = (v36 | (v36 >> ((1) & 63)))
                                                                var v62 int64 = 9223372036854775807
                                                                v1 = uint64(v2)
                                                                _ = v61
                                                                _ = v62
                                                            case "a":
                                                                v8 += v8
                                                            case "\n\t":
                                                                switch (fi59 << ((v6) & 63)) {
                                                                case 128:
                                                                    switch fi59 {
                                                                    case -2147483648:
                                                                        var v63 string = ("" + "\n\t")
                                                                        _ = v63
                                                                    case -1:
                                                                        _ = len(v16)
                                                                    case 127:
                                                                        var v64 uint8 = v3
                                                                        _ = v64
                                                                    }
                                                                case 127:
                                                                    if false {
                                                                        if true {
                                                                            v65 := v5["foo bar"]
                                                                            if (1.5 <= (1.5 + (0.0 - ((((((0.0 - 1e10) * 0.0) * 0.1) * 0.0) - 1e10) / 3.14159)))) {
                                                                                var v66 int8 = int8(v3)
                                                                                switch 32767 {
                                                                                case 4294967295:
                                                                                    v36 += v65
                                                                                    var v67 float64 = (-2.25 * ((1.5 - (0.1 / 1e10)) - 0.0))
                                                                                    switch ("hello" + ("foo bar" + ("\n\t" + "hello"))) {
                                                                                    case "\n\t":
                                                                                        var v68 uint32 = (v36 | uint32(v6))
                                                                                        _ = v68
                                                                                    case "":
                                                                                        v6 = uint(v3)
                                                                                        var v69 int32 = 1
                                                                                        _ = v69
                                                                                    }
                                                                                    _ = v67
                                                                                }
                                                                                v36 = v65
                                                                                fmt.Println(v2, v7, v8)
                                                                                _ = v66
                                                                            }
                                                                            _ = len(v16)
                                                                            _ = v65
                                                                        }
                                                                    } else {
                                                                        v70 := GenS7{F0: -9223372036854775808, F1: (v8 + (v8 & v1))}
                                                                        var v71 int32 = 626
                                                                        fmt.Println(fi46)
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                case 42:
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        v10 += v10
                                                                    case "hello":
                                                                        if len(v16) >= 2 {
                                                                            v16 = v16[1:len(v16)]
                                                                        }
                                                                    default:
                                                                        v7 = (v7 & (v7 & -32768))
                                                                    }
                                                                }
                                                            default:
                                                                v72 := GenS10{F0: "", F1: (v3 - v3), F2: (v4 >> (((v6 >> (((v2 * v2)) & 63))) & 63))}
                                                                _ = v72
                                                            }
                                                            var v73 int64 = 0
                                                            _ = fi59
                                                            _ = v60
                                                            _ = v73
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    case "foo bar":
                                        switch "hello" {
                                        case "a":
                                            for fi74 := 0; fi74 < 6; fi74++ {
                                                fmt.Println(v8)
                                                _ = fi74
                                            }
                                        case "foo bar":
                                            v6 -= (v2 / (((v6 % (((v2 | (^(v2)))) | 1))) | 1))
                                        case "\n\t":
                                            for fi75 := 0; fi75 < 12; fi75++ {
                                                v1 += 1
                                                _ = fi75
                                            }
                                        default:
                                            for fi76 := 0; fi76 < 2; fi76++ {
                                                var v77 uint = v2
                                                fmt.Println(v3, v6, v4)
                                                _ = fi76
                                                _ = v77
                                            }
                                        }
                                    }
                                    fmt.Println(v7, v36)
                                    _ = fi46
                                }
                            }
                            _ = v36
                        }
                        switch v4 {
                        case -305:
                            var v78 uint64 = v8
                            _ = v78
                        case -665:
                            v10 = (v7 << ((v2) & 63))
                        default:
                            v2 = v2
                        }
                    } else {
                        for fi79 := 0; fi79 < 1; fi79++ {
                            var v80 string = ((("\n\t" + "\n\t") + "hello") + "")
                            _ = fi79
                            _ = v80
                        }
                        fmt.Println(v2, v4)
                    }
                    _ = v35
                }
                if true {
                    v81 := GenS9{F0: 51}
                    _ = v81
                } else {
                    var v82 uint8 = (^(v3))
                    _ = v82
                }
                _ = v28
            }
            var v83 int = v4
            _ = v83
        }
    }
    for fi84 := 0; fi84 < 4; fi84++ {
        switch (v4 - 1) {
        case 9223372036854775807:
            v5["foo bar"] = 65535
        case -2147483648:
            switch fi84 {
            case 32768:
                v8 = 18446744073709551615
            case 9223372036854775807:
                for fi85 := 0; fi85 < 3; fi85++ {
                    var v86 int16 = v7
                    if (32767 == (v86 >> (((^(v2))) & 63))) {
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        if (false || (821 >= 48)) {
                            delete(v5, "\n\t")
                            var v87 float32 = 0.1
                            _ = v87
                        } else {
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            v86 -= 840
                            v10 -= (v86 / (((v86 >> ((v2) & 63))) | 1))
                            _ = len(v16)
                        }
                        if true {
                            var v88 int64 = -346
                            var v89 float64 = ((3.14159 / (1.5 * (((0.0 + 3.14159) * 0.0) - 0.1))) * (1e10 - (float64((-2.25 + 0.1)) / 1.0)))
                            _ = v88
                            _ = v89
                        }
                        var v90 bool = true
                        _ = len(v5)
                        _ = v90
                    } else {
                        var v91 uint = v2
                        if true {
                            v5[""] = 820
                            v10 = -618
                            _ = len(v5)
                        }
                        if (42 != -645) {
                            fmt.Println(v7, v91, fi85)
                        }
                        _ = v91
                    }
                    if true {
                        var v92 int = -535
                        _ = v92
                    } else {
                        v6 -= (v2 << ((v2) & 63))
                        fmt.Println(fi84, v10)
                        var v93 string = (("foo bar" + ("\n\t" + ("hello" + "foo bar"))) + ("\n\t" + ("hello" + ("" + "foo bar"))))
                        v94 := GenS11{F0: v93, F1: (v3 - v3), F2: -2147483648}
                        _ = v93
                        _ = v94
                    }
                    _ = fi85
                    _ = v86
                }
            case 42:
                v95 := GenS8{F0: 0, F1: (v7 - (v10 * v10))}
                _ = v95
            }
        case 128:
            for fi96 := 0; fi96 < 5; fi96++ {
                switch (("hello" + "hello") + "hello") {
                case "foo bar":
                    _ = len(v5)
                default:
                    v8 = 4294967295
                }
                v6 += (^(v2))
                _ = fi96
            }
        default:
            var v97 uint32 = 993
            _ = v97
        }
        var v98 int8 = 96
        _ = fi84
        _ = v98
    }
    v1 = v8
    if false {
        switch 128 {
        case -9223372036854775808:
            v2 = (^(v6))
        case 2147483647:
            var v99 uint = v6
            _ = v99
        case 39:
            fmt.Println(v10, v4)
        default:
            var v100 uint64 = (v1 - (v1 << ((v6) & 63)))
            _ = v100
        }
        var v101 uint8 = v3
        for fi102 := 0; fi102 < 16; fi102++ {
            var v103 uint8 = v101
            switch -9223372036854775808 {
            case 32767:
                var v104 uint32 = 0
                _ = v104
            }
            switch (fi102 % ((32767) | 1)) {
            case 1:
                switch fi102 {
                case 42:
                    fmt.Println(v3, v1, v8)
                case 32767:
                    v105 := GenS7{F0: -650, F1: (v8 & uint64(v2))}
                    _ = v105
                case -441:
                    if true {
                        var v106 int16 = 32767
                        for fi107 := 0; fi107 < 5; fi107++ {
                            var v108 uint16 = uint16(v8)
                            switch (v4 - (v4 & (v4 - fi107))) {
                            case 9223372036854775807:
                                v109 := GenS2{F0: -32768}
                                v110 := GenS12{F0: true, F1: 65535, F2: (v8 + v1)}
                                _ = v109
                                _ = v110
                            }
                            if (v1 > (v8 + v1)) {
                                if (0.1 > (0.1 / -2.25)) {
                                    v103 = v103
                                    var v111 uint16 = v108
                                    _ = len(v16)
                                    v2 = (v2 ^ (v6 >> ((uint(v3)) & 63)))
                                    var v112 int64 = -9223372036854775808
                                    v103 -= v3
                                    _ = v111
                                    _ = v112
                                } else {
                                    v6 = 914
                                    if false {
                                        fmt.Println(v10, v7)
                                        v4 += (fi107 & v4)
                                        switch v4 {
                                        case 541:
                                            v7 -= (^(v7))
                                            var v113 uint = 299
                                            v114 := GenS13{F0: 0, F1: (^(v103)), F2: 4}
                                            _ = v113
                                            _ = v114
                                        case 128:
                                            v103 += 128
                                        case 127:
                                            v106 += (v106 / (((v106 / ((v10) | 1))) | 1))
                                        default:
                                            if ((("a" + "\n\t") + "hello") > ("foo bar" + "hello")) {
                                                delete(v5, "\n\t")
                                                v8 += (^(v1))
                                                v115 := GenS14{F0: -128, F1: true}
                                                if (-2.25 < 3.14159) {
                                                    v116 := GenS10{F0: ("" + ((("" + "\n\t") + "hello") + "a")), F1: v101, F2: (fi107 | fi102)}
                                                    _ = v116
                                                }
                                                switch "a" {
                                                case "a":
                                                    if len(v16) > 0 {
                                                        v16[0] = (v4 - v4)
                                                    }
                                                case "foo bar":
                                                    switch (("foo bar" + ("foo bar" + (("foo bar" + "hello") + ""))) + "foo bar") {
                                                    case "foo bar":
                                                        if len(v16) >= 2 {
                                                            v16 = v16[1:len(v16)]
                                                        }
                                                    case "":
                                                        if true {
                                                            if ((v3 | (v101 ^ 2)) != (v101 - v101)) {
                                                                v6 = 1
                                                                v117 := v5["foo bar"]
                                                                _ = v117
                                                            } else {
                                                                fmt.Println(fi107)
                                                            }
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(v101)
                                                        switch "foo bar" {
                                                        case "":
                                                            if true {
                                                                var v118 uint = (v2 / (((^(v6))) | 1))
                                                                _ = v118
                                                            } else {
                                                                if ((100 != -1) || (1 > 42)) {
                                                                    if true {
                                                                        var v119 uint32 = 970
                                                                        v6 -= v2
                                                                        v108 -= (v108 ^ v108)
                                                                        _ = v119
                                                                    } else {
                                                                        switch (v4 - (fi102 << (((v6 * (^(v2)))) & 63))) {
                                                                        case 658:
                                                                            if (!false) {
                                                                                v4 = -1
                                                                                if (0 < uint32(v2)) {
                                                                                    v120 := GenS2{F0: int16(fi102)}
                                                                                    _ = v120
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v6, v106, v103)
                                                                                v121 := GenS9{F0: -1}
                                                                                _ = v121
                                                                            }
                                                                        default:
                                                                            if (!true) {
                                                                                v1 -= (v1 | v1)
                                                                            }
                                                                        }
                                                                        var v122 uint64 = v8
                                                                        _ = v122
                                                                    }
                                                                    var v123 uint8 = (^(v101))
                                                                    if true {
                                                                        var v124 int8 = 100
                                                                        v101 = (v101 * v123)
                                                                        _ = v124
                                                                    }
                                                                    _ = v123
                                                                }
                                                            }
                                                        }
                                                        v125 := GenS8{F0: 127, F1: (^(v10))}
                                                        v126 := GenS13{F0: -1, F1: v101, F2: v103}
                                                        _ = v125
                                                        _ = v126
                                                    default:
                                                        var v127 uint16 = 2
                                                        _ = v127
                                                    }
                                                case "hello":
                                                    v128 := GenS15{F0: -11, F1: v108, F2: 0}
                                                    _ = v128
                                                default:
                                                    if (-579 == v7) {
                                                        v1 = v8
                                                        v101 = (v103 / ((v103) | 1))
                                                    }
                                                }
                                                _ = v115
                                            }
                                        }
                                        if (v108 < v108) {
                                            var v129 float32 = (1.5 / 1.0)
                                            v5[(("" + "foo bar") + (("" + "\n\t") + ("" + "")))] = 2
                                            switch (fi102 ^ (fi102 ^ (^(v4)))) {
                                            case 127:
                                                fmt.Println(v6, v106, fi107)
                                            case -273:
                                                var v130 string = ("foo bar" + "\n\t")
                                                _ = v130
                                            case 9223372036854775807:
                                                v2 = v2
                                            default:
                                                v10 = v10
                                            }
                                            var v131 string = ""
                                            _ = v129
                                            _ = v131
                                        }
                                    }
                                }
                                v101 = v3
                            } else {
                                v132 := v5[("foo bar" + (("foo bar" + ("\n\t" + "a")) + "foo bar"))]
                                var v133 int32 = -2147483648
                                if true {
                                    fmt.Println(v8, v10)
                                    v6 = v6
                                } else {
                                    if false {
                                        var v134 bool = false
                                        _ = len(v16)
                                        _ = v134
                                    } else {
                                        if (v2 <= v6) {
                                            delete(v5, "")
                                            v135 := GenS1{F0: int(v108), F1: v3, F2: v103}
                                            _ = v135
                                        } else {
                                            v106 = (v7 ^ (v10 % ((-32768) | 1)))
                                        }
                                        var v136 uint32 = 71
                                        v10 = -1
                                        _ = v136
                                    }
                                    var v137 uint8 = (^(v3))
                                    v138 := GenS16{F0: (v133 >> (((v6 / ((524) | 1))) & 63)), F1: v108, F2: v133}
                                    if true {
                                        fmt.Println(v8, v106, v8)
                                        v139 := GenS17{F0: 128}
                                        _ = v139
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        if ((false || true) || (-1 <= 1)) {
                                            if (1.5 == ((-2.25 / 0.1) * (1.5 - 0.1))) {
                                                if (1 < (fi102 + fi102)) {
                                                    if true {
                                                        fmt.Println(v3)
                                                        v140 := GenS18{F0: (v133 << (((v6 | uint(v8))) & 63)), F1: (v132 - v132), F2: true}
                                                        if true {
                                                            v2 += v2
                                                        }
                                                        _ = v140
                                                    }
                                                    var v141 int16 = (v7 / ((1) | 1))
                                                    v108 -= uint16(v6)
                                                    _ = v141
                                                } else {
                                                    var v142 uint = (v6 / ((2) | 1))
                                                    _ = v142
                                                }
                                            }
                                            var v143 string = "\n\t"
                                            _ = v143
                                        }
                                    case "":
                                        v103 = v101
                                    case "foo bar":
                                        delete(v5, (((("a" + "hello") + "a") + ("hello" + "\n\t")) + "\n\t"))
                                    }
                                    _ = v137
                                    _ = v138
                                }
                                _ = v132
                                _ = v133
                            }
                            _ = fi107
                            _ = v108
                        }
                        var v144 uint64 = 1
                        _ = v106
                        _ = v144
                    }
                default:
                    v4 += 528
                }
            default:
                var v145 uint8 = v101
                _ = v145
            }
            _ = fi102
            _ = v103
        }
        _ = v101
    }
    if (-9223372036854775808 < -9223372036854775808) {
        if (!(!false)) {
            fmt.Println(v10, v10, v2)
            v6 -= (v2 << ((v6) & 63))
            var v146 int64 = 310
            _ = v146
        } else {
            v2 += v6
            if true {
                var v147 uint16 = 0
                if true {
                    v148 := GenS19{F0: 143, F1: v2}
                    v7 = 32767
                    _ = v148
                }
                var v149 float32 = ((0.1 * float32(-2.25)) * 0.0)
                _ = v147
                _ = v149
            } else {
                var v150 int32 = -2147483648
                var v151 uint64 = v8
                if len(v16) > 0 {
                    v16[0] = v4
                }
                _ = v150
                _ = v151
            }
        }
        v3 -= (v3 * v3)
    }
    if ((1.5 + (0.1 + -2.25)) >= (3.14159 * (float64(3.14159) - float64(float32((0.0 + float64(-2.25))))))) {
        v5[("\n\t" + "hello")] = 2
        if (v3 < v3) {
            switch ("a" + "\n\t") {
            case "a":
                if len(v16) > 0 {
                    v16[0] = (v4 / ((int(v2)) | 1))
                }
            case "hello":
                var v152 uint64 = v1
                _ = v152
            default:
                v1 += v1
            }
            fmt.Println(v2, v7, v4)
            if (false || false) {
                for fi153 := 0; fi153 < 15; fi153++ {
                    v3 = (v3 << (((v2 << ((v2) & 63))) & 63))
                    for fi154 := 0; fi154 < 9; fi154++ {
                        if (4294967295 <= 520) {
                            if false {
                                fmt.Println(v3, fi153, v1)
                            } else {
                                var v155 string = "hello"
                                switch fi154 {
                                case 1:
                                    switch int(v2) {
                                    case -313:
                                        var v156 float32 = 1e10
                                        _ = v156
                                    case 32767:
                                        v3 += v3
                                    }
                                case 42:
                                    if (26 > -1) {
                                        var v157 int64 = -9223372036854775808
                                        var v158 int16 = (v10 | -13)
                                        delete(v5, v155)
                                        _ = v157
                                        _ = v158
                                    } else {
                                        _ = len(v5)
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                    }
                                case 128:
                                    fmt.Println(v7, v3)
                                default:
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                }
                                _ = v155
                            }
                            if false {
                                if true {
                                    switch (fi153 + fi154) {
                                    case -2147483648:
                                        _ = len(v16)
                                    case 111:
                                        delete(v5, "a")
                                    case 9223372036854775807:
                                        var v159 int16 = (-(v10))
                                        fmt.Println(v7)
                                        v8 = v1
                                        v2 += (v2 >> ((757) & 63))
                                        if false {
                                            if (v2 > 822) {
                                                switch (((("" + "a") + "\n\t") + ("" + ("\n\t" + ("hello" + "foo bar")))) + (("" + ("" + "foo bar")) + "hello")) {
                                                case "hello":
                                                    var v160 uint64 = (v1 | v1)
                                                    _ = v160
                                                case "a":
                                                    switch (("a" + "") + "foo bar") {
                                                    case "a":
                                                        switch (fi153 >> ((v2) & 63)) {
                                                        case 1:
                                                            _ = len(v16)
                                                        default:
                                                            delete(v5, "hello")
                                                        }
                                                    }
                                                case "":
                                                    var v161 uint = 472
                                                    _ = v161
                                                }
                                                fmt.Println(fi153, v2)
                                            } else {
                                                fmt.Println(fi154, v2)
                                                var v162 uint64 = v1
                                                var v163 uint = v2
                                                var v164 int16 = v159
                                                v5["hello"] = 2147483647
                                                _ = v162
                                                _ = v163
                                                _ = v164
                                            }
                                            var v165 int8 = 0
                                            if (false || false) {
                                                if len(v16) > 0 {
                                                    v16[0] = -305
                                                }
                                                if false {
                                                    var v166 float32 = 1.5
                                                    v167 := GenS20{F0: v165, F1: int64(fi154), F2: v1}
                                                    v168 := GenS21{F0: 4294967295}
                                                    _ = v166
                                                    _ = v167
                                                    _ = v168
                                                } else {
                                                    v16 = append(v16, v4)
                                                }
                                            }
                                            _ = v165
                                        }
                                        _ = v159
                                    }
                                }
                                var v169 bool = ((true && (!(!(!(-1 != 347))))) || (309 != 42))
                                if false {
                                    if false {
                                        v5["\n\t"] = 0
                                        var v170 int = fi154
                                        switch "" {
                                        case "":
                                            if (!(v169 || v169)) {
                                                v171 := GenS22{F0: "foo bar", F1: "foo bar"}
                                                var v172 uint8 = (v3 + v3)
                                                v3 = v3
                                                _ = v171
                                                _ = v172
                                            } else {
                                                v173 := GenS18{F0: -2147483648, F1: 4294967295, F2: (-2.25 == (float32(0.0) - 1e10))}
                                                v2 += v2
                                                _ = v173
                                            }
                                        default:
                                            v174 := GenS8{F0: 42, F1: (v10 % ((v7) | 1))}
                                            _ = v174
                                        }
                                        _ = v170
                                    }
                                    var v175 int64 = -1
                                    _ = v175
                                }
                                _ = v169
                            } else {
                                switch "" {
                                case "hello":
                                    var v176 string = "foo bar"
                                    _ = v176
                                case "foo bar":
                                    v1 -= uint64(v2)
                                default:
                                    var v177 float32 = 3.14159
                                    v178 := GenS4{F0: 42, F1: v6, F2: 42}
                                    switch (("" + "hello") + "\n\t") {
                                    case "hello":
                                        var v179 uint = 0
                                        v8 = (v8 & 834)
                                        v180 := GenS19{F0: uint32(v3), F1: v2}
                                        v3 += (v3 | (v3 - (v3 >> (((v6 * v6)) & 63))))
                                        v8 = (^(v1))
                                        _ = v179
                                        _ = v180
                                    case "foo bar":
                                        v181 := GenS13{F0: -1, F1: v3, F2: 0}
                                        _ = v181
                                    case "a":
                                        if true {
                                            if true {
                                                var v182 uint32 = 2147483647
                                                v182 -= (v182 / (((v182 >> ((1) & 63))) | 1))
                                                _ = v182
                                            }
                                            v183 := GenS23{F0: 0, F1: 0, F2: -1}
                                            v184 := GenS24{F0: 644}
                                            _ = v183
                                            _ = v184
                                        }
                                    }
                                    _ = v177
                                    _ = v178
                                }
                                switch 542 {
                                case -1:
                                    v6 += 1
                                }
                                v8 += (v1 * (^(v8)))
                            }
                        } else {
                            v185 := GenS15{F0: 0, F1: uint16(v2), F2: 2147483647}
                            _ = v185
                        }
                        v8 = (v1 ^ 0)
                        _ = fi154
                    }
                    _ = fi153
                }
                if len(v16) > 0 {
                    v16[0] = -2147483648
                }
            } else {
                v10 = v7
                if (0.1 == float32(1e10)) {
                    var v186 uint32 = 2147483648
                    v187 := GenS20{F0: 42, F1: -1, F2: (^(v8))}
                    for fi188 := 0; fi188 < 2; fi188++ {
                        var v189 uint = v2
                        var v190 uint64 = 737
                        _ = fi188
                        _ = v189
                        _ = v190
                    }
                    _ = v186
                    _ = v187
                }
            }
        }
        fmt.Println(v3)
    } else {
        if true {
            var v191 int64 = -9223372036854775808
            for fi192 := 0; fi192 < 12; fi192++ {
                v8 = v1
                v193 := GenS16{F0: -590, F1: 1, F2: 2147483647}
                fmt.Println(fi192)
                _ = fi192
                _ = v193
            }
            _ = v191
        }
        v4 -= v4
    }
    v194 := GenS25{F0: (v6 & (v6 % (((v6 - 2)) | 1)))}
    v195 := GenS26{F0: v4, F1: (v4 << ((26) & 63)), F2: (^(v8))}
    switch ("foo bar" + "") {
    case "":
        v10 -= v7
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v16
    _ = v194
    _ = v195
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
