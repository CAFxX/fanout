package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64; F1 uint; F2 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint16; F1 int32; F2 bool }
type GenS11 struct { F0 string; F1 bool }
type GenS12 struct { F0 int8; F1 bool }
type GenS13 struct { F0 uint16; F1 string; F2 int64 }
type GenS14 struct { F0 uint64; F1 int64 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS17 struct { F0 uint16; F1 uint64 }
type GenS18 struct { F0 uint8; F1 int64 }
type GenS19 struct { F0 uint64; F1 int64; F2 int }
type GenS20 struct { F0 string; F1 uint; F2 int8 }
type GenS21 struct { F0 int64; F1 int64; F2 bool }
type GenS22 struct { F0 uint; F1 bool; F2 bool }
type GenS23 struct { F0 uint16; F1 int64 }
type GenS24 struct { F0 string }
type GenS25 struct { F0 uint }
type GenS26 struct { F0 uint64 }
func genFn0(p0 int, p1 int) int8 {
    fmt.Println(p0, p0)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 string) int16 {
    var v4 float64 = ((float64((float32(0.1) - 1.5)) + 1e10) + 0.0)
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            v4 = (float64(1e10) * 3.14159)
            _ = fi6
        }
        for fi7 := 0; fi7 < 4; fi7++ {
            fmt.Println(fi7, p0, fi7)
            var v8 int16 = 1
            p0 = p0
            _ = fi7
            _ = v8
        }
        _ = fi5
    }
    fmt.Println(v4, v4)
    fmt.Println(p0, p0)
    _ = p0
    _ = v4
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 1e10
    var v3 uint32 = 1
    v9 := GenS1{F0: 32767}
    var v10 uint64 = 537
    fmt.Println(v3, v2, v10)
    if true {
        v2 -= v2
        fmt.Println(v10, v10)
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS2{F0: v10, F1: 1, F2: v3}
        _ = fi11
        _ = v12
    }
    v13 := []int{-791}
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v2, v2)
    case "":
        if true {
            switch "a" {
            case "a":
                switch -1 {
                case -420:
                    v2 = (1.5 + float32(((0.1 - 3.14159) / 1.0)))
                case 127:
                    for fi14 := 0; fi14 < 9; fi14++ {
                        v15 := GenS3{F0: 42, F1: (fi14 ^ fi14)}
                        for fi16 := 0; fi16 < 12; fi16++ {
                            var v17 int = (fi16 | 128)
                            if true {
                                var v18 uint32 = 2
                                v10 += 2
                                switch (("a" + "hello") + "a") {
                                case "hello":
                                    switch (v17 ^ (v17 % (((fi14 * fi16)) | 1))) {
                                    case -116:
                                        switch (^(fi14)) {
                                        case -552:
                                            switch fi14 {
                                            case 128:
                                                var v19 uint32 = (v3 % ((v18) | 1))
                                                _ = v19
                                            default:
                                                switch (^(v17)) {
                                                case 128:
                                                    if false {
                                                        fmt.Println(v3, fi16, v18)
                                                    } else {
                                                        if true {
                                                            var v20 uint = 42
                                                            if false {
                                                                var v21 uint16 = 1
                                                                fmt.Println(fi16, v3, v2)
                                                                var v22 bool = (128 == 255)
                                                                _ = v21
                                                                _ = v22
                                                            } else {
                                                                v10 += uint64(v20)
                                                                if ((!true) || true) {
                                                                    _ = len(v1)
                                                                    if ((1.5 / 3.14159) <= 0.1) {
                                                                        if (!(("hello" + "\n\t") >= "a")) {
                                                                            var v23 uint64 = uint64(v20)
                                                                            v24 := GenS2{F0: 1, F1: 0, F2: (v3 & v18)}
                                                                            _ = v23
                                                                            _ = v24
                                                                        }
                                                                        switch (v17 << ((241) & 63)) {
                                                                        case 9223372036854775807:
                                                                            v25 := GenS4{F0: 4294967295}
                                                                            _ = v25
                                                                        default:
                                                                            var v26 int16 = 0
                                                                            v27 := v1[(v18 << ((v20) & 63))]
                                                                            fmt.Println(v18, v26, v2)
                                                                            _ = v26
                                                                            _ = v27
                                                                        }
                                                                        fmt.Println(fi14, v10, v10)
                                                                    } else {
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            var v28 int16 = -32768
                                                                            _ = v28
                                                                        case "foo bar":
                                                                            switch ("foo bar" + "\n\t") {
                                                                            case "\n\t":
                                                                                if false {
                                                                                    var v29 int64 = 1
                                                                                    v20 = 4294967295
                                                                                    v30 := GenS5{F0: (-(v29))}
                                                                                    _ = v29
                                                                                    _ = v30
                                                                                } else {
                                                                                    v31 := v1[(^(v18))]
                                                                                    _ = v31
                                                                                }
                                                                            case "hello":
                                                                                var v32 uint16 = 612
                                                                                _ = v32
                                                                            case "foo bar":
                                                                                var v33 uint = (v20 - v20)
                                                                                _ = v33
                                                                            }
                                                                        default:
                                                                            v10 += v10
                                                                        }
                                                                        if ((v18 < 2147483647) && false) {
                                                                            v17 -= (v17 / (((fi14 & fi16)) | 1))
                                                                            v3 = 967
                                                                            if false {
                                                                                v18 += (v18 / (((v18 % ((v3) | 1))) | 1))
                                                                            } else {
                                                                                if len(v13) > 0 {
                                                                                    v13[0] = v17
                                                                                }
                                                                            }
                                                                        }
                                                                        if false {
                                                                            var v34 int64 = -93
                                                                            _ = v34
                                                                        } else {
                                                                            v35 := v1[(v3 ^ 2147483647)]
                                                                            v36 := GenS6{F0: 208}
                                                                            switch v17 {
                                                                            case 812:
                                                                                fmt.Println(v10, v18)
                                                                            case 127:
                                                                                var v37 int8 = -128
                                                                                _ = v37
                                                                            }
                                                                            var v38 float32 = 0.1
                                                                            _ = v35
                                                                            _ = v36
                                                                            _ = v38
                                                                        }
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        if (v18 >= (v3 ^ v3)) {
                                                                            if false {
                                                                                var v39 int32 = 276
                                                                                v40 := GenS7{F0: "foo bar"}
                                                                                v10 = v10
                                                                                _ = v39
                                                                                _ = v40
                                                                            }
                                                                            if true {
                                                                                v2 = (v2 * float32(1.5))
                                                                                if (40 > 2) {
                                                                                    switch 32767 {
                                                                                    case 42:
                                                                                        var v41 float32 = v2
                                                                                        _ = v41
                                                                                    case 3:
                                                                                        v42 := v1[v3]
                                                                                        _ = v42
                                                                                    case 2147483647:
                                                                                        if true {
                                                                                            switch "\n\t" {
                                                                                            case "":
                                                                                                v17 -= -1
                                                                                            case "a":
                                                                                                var v43 uint16 = uint16(v18)
                                                                                                _ = v43
                                                                                            default:
                                                                                                var v44 int = -735
                                                                                                _ = v44
                                                                                            }
                                                                                            v17 += (v17 >> ((555) & 63))
                                                                                        } else {
                                                                                            v45 := GenS7{F0: ("a" + "hello")}
                                                                                            fmt.Println(v17, v18, v3)
                                                                                            _ = v45
                                                                                        }
                                                                                    }
                                                                                    delete(v1, (v18 + (v18 & v3)))
                                                                                } else {
                                                                                    switch fi16 {
                                                                                    case -113:
                                                                                        v46 := GenS6{F0: 2}
                                                                                        _ = v46
                                                                                    case 2147483647:
                                                                                        switch fi16 {
                                                                                        case 0:
                                                                                            v20 = 1
                                                                                        default:
                                                                                            v10 = v10
                                                                                        }
                                                                                    case 4294967295:
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            if false {
                                                                                                if (1000 != -32768) {
                                                                                                    switch ((("foo bar" + (("hello" + "") + ("foo bar" + "hello"))) + ((("\n\t" + "") + "\n\t") + (((("\n\t" + ("hello" + ("hello" + ""))) + "a") + ("\n\t" + "foo bar")) + "\n\t"))) + ("" + "hello")) {
                                                                                                    case "hello":
                                                                                                        fmt.Println(v2, v20, v3)
                                                                                                        v10 += (v10 & v10)
                                                                                                        if (!true) {
                                                                                                            var v47 string = "a"
                                                                                                            v48 := GenS2{F0: v10, F1: uint8(v3), F2: v18}
                                                                                                            if true {
                                                                                                                fmt.Println(fi14)
                                                                                                                var v49 bool = (!((0.1 + (1e10 - 0.1)) >= 3.14159))
                                                                                                                _ = v49
                                                                                                            }
                                                                                                            _ = v47
                                                                                                            _ = v48
                                                                                                        } else {
                                                                                                            switch (fi14 + 127) {
                                                                                                            case 127:
                                                                                                                v10 += v10
                                                                                                            }
                                                                                                            switch "foo bar" {
                                                                                                            case "hello":
                                                                                                                v20 += v20
                                                                                                            case "":
                                                                                                                v3 += (v3 % (((v18 >> (((v20 / ((65535) | 1))) & 63))) | 1))
                                                                                                            case "a":
                                                                                                                v20 -= (v20 % ((2) | 1))
                                                                                                            default:
                                                                                                                if false {
                                                                                                                    var v50 uint32 = v3
                                                                                                                    var v51 int8 = -104
                                                                                                                    switch -9223372036854775808 {
                                                                                                                    case 4294967295:
                                                                                                                        var v52 float64 = (0.0 - -2.25)
                                                                                                                        _ = v52
                                                                                                                    case 32768:
                                                                                                                        var v53 uint8 = 127
                                                                                                                        _ = v53
                                                                                                                    case 0:
                                                                                                                        var v54 uint32 = (v3 ^ (v50 * v50))
                                                                                                                        _ = v54
                                                                                                                    default:
                                                                                                                        v20 = (v20 % (((v20 % (((v20 | 2)) | 1))) | 1))
                                                                                                                        v55 := GenS6{F0: 127}
                                                                                                                        v13 = append(v13, (-(v17)))
                                                                                                                        _ = v55
                                                                                                                    }
                                                                                                                    _ = v50
                                                                                                                    _ = v51
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        v56 := GenS8{F0: 9223372036854775807, F1: v20, F2: 255}
                                                                                                                        _ = v56
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            switch "foo bar" {
                                                                                                            case "\n\t":
                                                                                                                var v57 int = (v17 * (fi14 & fi16))
                                                                                                                _ = v57
                                                                                                            case "foo bar":
                                                                                                                switch "\n\t" {
                                                                                                                case "a":
                                                                                                                    var v58 uint8 = 2
                                                                                                                    _ = v58
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        var v59 int = v17
                                                                                                        if len(v13) > 0 {
                                                                                                            v13[0] = (v59 * 1)
                                                                                                        }
                                                                                                        if (1 > 1000) {
                                                                                                            if (v2 == (float32(0.0) / (v2 - float32((((3.14159 - 1e10) - 0.0) + 0.1))))) {
                                                                                                                fmt.Println(v59, v59, v17)
                                                                                                            }
                                                                                                            v60 := v1[uint32(v20)]
                                                                                                            switch ("\n\t" + "a") {
                                                                                                            case "foo bar":
                                                                                                                if false {
                                                                                                                    if ((^(v17)) < (fi14 << ((42) & 63))) {
                                                                                                                        v61 := v1[v60]
                                                                                                                        _ = v61
                                                                                                                    }
                                                                                                                    if false {
                                                                                                                        v62 := GenS5{F0: -1}
                                                                                                                        _ = v62
                                                                                                                    }
                                                                                                                    v10 = uint64(v3)
                                                                                                                } else {
                                                                                                                    var v63 int16 = 42
                                                                                                                    fmt.Println(v60, v63)
                                                                                                                    if len(v13) >= 2 {
                                                                                                                        v13 = v13[1:len(v13)]
                                                                                                                    }
                                                                                                                    _ = v63
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                switch (-(fi16)) {
                                                                                                                case 0:
                                                                                                                    v60 = 4294967295
                                                                                                                case 128:
                                                                                                                    v10 += 42
                                                                                                                    switch "\n\t" {
                                                                                                                    case "a":
                                                                                                                        switch "a" {
                                                                                                                        case "hello":
                                                                                                                            var v64 int32 = int32(v18)
                                                                                                                            _ = v64
                                                                                                                        case "":
                                                                                                                            var v65 uint32 = v18
                                                                                                                            _ = v65
                                                                                                                        }
                                                                                                                    }
                                                                                                                case 748:
                                                                                                                    if ((fi16 > fi14) && true) {
                                                                                                                        v66 := GenS4{F0: v3}
                                                                                                                        var v67 uint64 = v10
                                                                                                                        _ = v66
                                                                                                                        _ = v67
                                                                                                                    }
                                                                                                                    v68 := GenS3{F0: uint8(v10), F1: fi14}
                                                                                                                    switch v59 {
                                                                                                                    case -9223372036854775808:
                                                                                                                        var v69 uint32 = (v18 ^ v60)
                                                                                                                        _ = v69
                                                                                                                    case 32767:
                                                                                                                        var v70 int64 = int64(v60)
                                                                                                                        _ = v70
                                                                                                                    case -2147483648:
                                                                                                                        var v71 bool = (32767 <= 11)
                                                                                                                        _ = v71
                                                                                                                    default:
                                                                                                                        switch fi16 {
                                                                                                                        case 42:
                                                                                                                            v72 := GenS3{F0: 198, F1: -1}
                                                                                                                            _ = v72
                                                                                                                        case 4294967295:
                                                                                                                            if true {
                                                                                                                                var v73 uint16 = 42
                                                                                                                                _ = v73
                                                                                                                            } else {
                                                                                                                                v20 += v20
                                                                                                                                _ = len(v1)
                                                                                                                                v60 = v60
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            switch ("" + ("\n\t" + ((("hello" + (("\n\t" + "") + (("" + "foo bar") + "foo bar"))) + "\n\t") + ""))) {
                                                                                                                            case "":
                                                                                                                                fmt.Println(fi16, v20, v2)
                                                                                                                            case "a":
                                                                                                                                v1[2147483647] = v60
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v17 = fi16
                                                                                                                    _ = v68
                                                                                                                default:
                                                                                                                    switch "hello" {
                                                                                                                    case "":
                                                                                                                        if (true || false) {
                                                                                                                            var v74 uint8 = 32
                                                                                                                            fmt.Println(v17, v74)
                                                                                                                            _ = v74
                                                                                                                        } else {
                                                                                                                            v75 := GenS7{F0: "hello"}
                                                                                                                            fmt.Println(v18)
                                                                                                                            v20 += 2
                                                                                                                            _ = v75
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            case "":
                                                                                                                var v76 int8 = 15
                                                                                                                _ = v76
                                                                                                            }
                                                                                                            _ = v60
                                                                                                        }
                                                                                                        switch v59 {
                                                                                                        case 1:
                                                                                                            var v77 string = "hello"
                                                                                                            _ = v77
                                                                                                        case -28:
                                                                                                            var v78 int32 = 42
                                                                                                            _ = v78
                                                                                                        default:
                                                                                                            switch ("a" + "foo bar") {
                                                                                                            case "foo bar":
                                                                                                                var v79 uint = (v20 << (((v20 | (v20 << (((v20 - v20)) & 63)))) & 63))
                                                                                                                _ = v79
                                                                                                            case "a":
                                                                                                                v59 -= 42
                                                                                                            default:
                                                                                                                var v80 int64 = 1
                                                                                                                _ = v80
                                                                                                            }
                                                                                                        }
                                                                                                        var v81 uint64 = (v10 & (v10 * (v10 * (v10 >> ((2) & 63)))))
                                                                                                        fmt.Println(fi14)
                                                                                                        _ = v59
                                                                                                        _ = v81
                                                                                                    case "":
                                                                                                        v17 = fi14
                                                                                                    case "foo bar":
                                                                                                        switch 128 {
                                                                                                        case 2147483647:
                                                                                                            if false {
                                                                                                                switch -9223372036854775808 {
                                                                                                                case 2147483647:
                                                                                                                    var v82 uint64 = v10
                                                                                                                    _ = v82
                                                                                                                default:
                                                                                                                    var v83 int32 = 2147483647
                                                                                                                    _ = v83
                                                                                                                }
                                                                                                                if true {
                                                                                                                    v13 = append(v13, (v17 % (((fi14 - 4294967295)) | 1)))
                                                                                                                    var v84 int16 = -32768
                                                                                                                    _ = v84
                                                                                                                } else {
                                                                                                                    if (false || (((^(v3)) < v18) || ((("\n\t" + "a") + "") >= ("" + ("" + (("" + "hello") + ("\n\t" + "a"))))))) {
                                                                                                                        var v85 int32 = 0
                                                                                                                        if false {
                                                                                                                            if (true || (42 > 65535)) {
                                                                                                                                var v86 string = "\n\t"
                                                                                                                                if true {
                                                                                                                                    var v87 uint8 = 1
                                                                                                                                    _ = v87
                                                                                                                                }
                                                                                                                                var v88 int64 = -9223372036854775808
                                                                                                                                _ = len(v1)
                                                                                                                                switch fi14 {
                                                                                                                                case 4294967295:
                                                                                                                                    v17 -= (fi16 % ((fi16) | 1))
                                                                                                                                }
                                                                                                                                _ = v86
                                                                                                                                _ = v88
                                                                                                                            }
                                                                                                                            if (2 >= 0) {
                                                                                                                                var v89 float64 = 3.14159
                                                                                                                                _ = v89
                                                                                                                            }
                                                                                                                            v2 -= float32(-2.25)
                                                                                                                            fmt.Println(v3, v85)
                                                                                                                            if (int16(v18) >= -32768) {
                                                                                                                                switch "" {
                                                                                                                                case "":
                                                                                                                                    v3 += v18
                                                                                                                                case "hello":
                                                                                                                                    v20 = (^(v20))
                                                                                                                                case "\n\t":
                                                                                                                                    var v90 uint64 = v10
                                                                                                                                    _ = v90
                                                                                                                                default:
                                                                                                                                    var v91 int8 = 127
                                                                                                                                    _ = v91
                                                                                                                                }
                                                                                                                                v92 := GenS9{F0: 3}
                                                                                                                                if true {
                                                                                                                                    var v93 float64 = float64(1e10)
                                                                                                                                    v94 := GenS2{F0: 42, F1: 4, F2: (v3 / ((v3) | 1))}
                                                                                                                                    v95 := GenS7{F0: "\n\t"}
                                                                                                                                    _ = v93
                                                                                                                                    _ = v94
                                                                                                                                    _ = v95
                                                                                                                                } else {
                                                                                                                                    v20 = 42
                                                                                                                                    var v96 string = ("" + "")
                                                                                                                                    if len(v13) > 0 {
                                                                                                                                        v13[0] = -2147483648
                                                                                                                                    }
                                                                                                                                    v18 += 2
                                                                                                                                    if false {
                                                                                                                                        v97 := GenS7{F0: v96}
                                                                                                                                        fmt.Println(v85, fi14, v10)
                                                                                                                                        _ = v97
                                                                                                                                    }
                                                                                                                                    switch v96 {
                                                                                                                                    case "hello":
                                                                                                                                        fmt.Println(v85, v17, v3)
                                                                                                                                    default:
                                                                                                                                        v10 = 750
                                                                                                                                        v1[v3] = 4294967295
                                                                                                                                        v2 += (v2 / v2)
                                                                                                                                    }
                                                                                                                                    var v98 float64 = ((((0.0 + (3.14159 * ((((-2.25 - 0.1) + -2.25) / float64(1.5)) + 0.1))) * 3.14159) + 0.0) + (0.0 * (((0.1 - 0.0) - (0.0 * ((1.5 + 1.5) * (1e10 / 1.5)))) * -2.25)))
                                                                                                                                    _ = v96
                                                                                                                                    _ = v98
                                                                                                                                }
                                                                                                                                switch 42 {
                                                                                                                                case 32767:
                                                                                                                                    v18 = (v18 & 1)
                                                                                                                                case 127:
                                                                                                                                    if (!true) {
                                                                                                                                        var v99 uint16 = 2
                                                                                                                                        v100 := GenS10{F0: (v99 * 481), F1: int32(v10), F2: true}
                                                                                                                                        switch ("foo bar" + "") {
                                                                                                                                        case "foo bar":
                                                                                                                                            v101 := GenS7{F0: "foo bar"}
                                                                                                                                            _ = v101
                                                                                                                                        case "\n\t":
                                                                                                                                            if ((!(true && ((^(v18)) >= v18))) && false) {
                                                                                                                                                var v102 bool = false
                                                                                                                                                fmt.Println(v17)
                                                                                                                                                v99 += (^(v99))
                                                                                                                                                _ = v102
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            v3 = (v18 >> (((v20 - (v20 % ((v20) | 1)))) & 63))
                                                                                                                                        }
                                                                                                                                        _ = v100
                                                                                                                                        _ = v99
                                                                                                                                    }
                                                                                                                                    fmt.Println(fi14)
                                                                                                                                    v13 = append(v13, (fi16 & fi16))
                                                                                                                                default:
                                                                                                                                    v18 += v18
                                                                                                                                }
                                                                                                                                var v103 int64 = 42
                                                                                                                                _ = v103
                                                                                                                                _ = v92
                                                                                                                            } else {
                                                                                                                                if false {
                                                                                                                                    v18 += v3
                                                                                                                                    if true {
                                                                                                                                        v85 += -2147483648
                                                                                                                                        if false {
                                                                                                                                            var v104 string = "a"
                                                                                                                                            var v105 string = ((v104 + (v104 + (v104 + v104))) + ("a" + (((((v104 + v104) + "hello") + "\n\t") + ((v104 + v104) + (((v104 + v104) + (v104 + (v104 + v104))) + "foo bar"))) + "foo bar")))
                                                                                                                                            v1[981] = v18
                                                                                                                                            _ = v104
                                                                                                                                            _ = v105
                                                                                                                                        } else {
                                                                                                                                            if true {
                                                                                                                                                var v106 int16 = -107
                                                                                                                                                var v107 int16 = int16(v85)
                                                                                                                                                if (!((-1 != -9223372036854775808) && true)) {
                                                                                                                                                    v108 := GenS8{F0: int64(v106), F1: 42, F2: 42}
                                                                                                                                                    if true {
                                                                                                                                                        v1[v3] = v3
                                                                                                                                                        var v109 bool = (116 > 314)
                                                                                                                                                        _ = v109
                                                                                                                                                    } else {
                                                                                                                                                        v110 := GenS8{F0: int64(v10), F1: 65535, F2: 42}
                                                                                                                                                        _ = v110
                                                                                                                                                    }
                                                                                                                                                    var v111 uint = (v20 | v20)
                                                                                                                                                    _ = v108
                                                                                                                                                    _ = v111
                                                                                                                                                }
                                                                                                                                                _ = v106
                                                                                                                                                _ = v107
                                                                                                                                            } else {
                                                                                                                                                if (false || (255 <= 48)) {
                                                                                                                                                    _ = len(v1)
                                                                                                                                                    var v112 uint32 = (^(v3))
                                                                                                                                                    _ = len(v13)
                                                                                                                                                    switch ("hello" + ("" + "a")) {
                                                                                                                                                    case "":
                                                                                                                                                        var v113 int32 = v85
                                                                                                                                                        _ = v113
                                                                                                                                                    }
                                                                                                                                                    v18 += (v112 >> (((v20 ^ 1)) & 63))
                                                                                                                                                    fmt.Println(fi16, v85)
                                                                                                                                                    v18 = v112
                                                                                                                                                    v18 += (v18 | (v3 % (((v3 % ((2147483647) | 1))) | 1)))
                                                                                                                                                    _ = v112
                                                                                                                                                }
                                                                                                                                                _ = len(v1)
                                                                                                                                                _ = len(v13)
                                                                                                                                            }
                                                                                                                                            var v114 uint32 = v18
                                                                                                                                            _ = v114
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        var v115 int64 = 0
                                                                                                                                        _ = v115
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                var v116 uint64 = 4294967295
                                                                                                                                v116 = (^(v10))
                                                                                                                                _ = v116
                                                                                                                            }
                                                                                                                            v117 := GenS11{F0: ("foo bar" + "hello"), F1: false}
                                                                                                                            switch (fi16 << ((4294967295) & 63)) {
                                                                                                                            case 0:
                                                                                                                                var v118 float32 = 3.14159
                                                                                                                                _ = v118
                                                                                                                            case 4294967295:
                                                                                                                                if (1 > 42) {
                                                                                                                                    var v119 uint = 0
                                                                                                                                    var v120 int64 = 919
                                                                                                                                    _ = v119
                                                                                                                                    _ = v120
                                                                                                                                } else {
                                                                                                                                    var v121 uint64 = 1
                                                                                                                                    var v122 bool = (!true)
                                                                                                                                    var v123 uint = 696
                                                                                                                                    _ = v121
                                                                                                                                    _ = v122
                                                                                                                                    _ = v123
                                                                                                                                }
                                                                                                                            case -2147483648:
                                                                                                                                v124 := GenS1{F0: 243}
                                                                                                                                _ = v124
                                                                                                                            }
                                                                                                                            _ = v117
                                                                                                                        } else {
                                                                                                                            switch fi16 {
                                                                                                                            case 42:
                                                                                                                                var v125 float64 = float64((1e10 - v2))
                                                                                                                                _ = v125
                                                                                                                            }
                                                                                                                            var v126 float64 = 1.5
                                                                                                                            v127 := GenS12{F0: 51, F1: false}
                                                                                                                            var v128 string = ("a" + "hello")
                                                                                                                            v10 += v10
                                                                                                                            var v129 int8 = -1
                                                                                                                            _ = v126
                                                                                                                            _ = v127
                                                                                                                            _ = v128
                                                                                                                            _ = v129
                                                                                                                        }
                                                                                                                        var v130 float64 = (1.5 - -2.25)
                                                                                                                        _ = v130
                                                                                                                        _ = v85
                                                                                                                    }
                                                                                                                }
                                                                                                                if true {
                                                                                                                    var v131 int8 = -117
                                                                                                                    if (55 != uint8(v3)) {
                                                                                                                        switch ("foo bar" + "") {
                                                                                                                        case "\n\t":
                                                                                                                            var v132 bool = false
                                                                                                                            _ = v132
                                                                                                                        case "foo bar":
                                                                                                                            if false {
                                                                                                                                fmt.Println(v10, v10)
                                                                                                                                if (uint8(v20) != 225) {
                                                                                                                                    v133 := GenS6{F0: 128}
                                                                                                                                    v2 -= 3.14159
                                                                                                                                    v18 += 301
                                                                                                                                    _ = v133
                                                                                                                                }
                                                                                                                                if false {
                                                                                                                                    if (!false) {
                                                                                                                                        v18 = (v3 << (((v20 | (^(v20)))) & 63))
                                                                                                                                        switch "" {
                                                                                                                                        case "foo bar":
                                                                                                                                            if (!(2 >= 2)) {
                                                                                                                                                _ = len(v1)
                                                                                                                                            } else {
                                                                                                                                                v134 := GenS2{F0: (^(v10)), F1: 180, F2: (v18 / (((v18 % (((v18 >> ((v20) & 63))) | 1))) | 1))}
                                                                                                                                                v18 = v18
                                                                                                                                                _ = v134
                                                                                                                                            }
                                                                                                                                        case "hello":
                                                                                                                                            switch (^(v17)) {
                                                                                                                                            case -560:
                                                                                                                                                fmt.Println(fi16, v18)
                                                                                                                                            case 42:
                                                                                                                                                if true {
                                                                                                                                                    fmt.Println(v20, v3)
                                                                                                                                                    var v135 int8 = (v131 - 42)
                                                                                                                                                    _ = v135
                                                                                                                                                } else {
                                                                                                                                                    v2 = v2
                                                                                                                                                    var v136 bool = (2 < 1)
                                                                                                                                                    _ = v136
                                                                                                                                                }
                                                                                                                                            case -9223372036854775808:
                                                                                                                                                switch fi16 {
                                                                                                                                                case -864:
                                                                                                                                                    switch (("\n\t" + "\n\t") + "a") {
                                                                                                                                                    case "hello":
                                                                                                                                                        if true {
                                                                                                                                                            fmt.Println(v3, v17)
                                                                                                                                                            var v137 int = fi14
                                                                                                                                                            _ = v137
                                                                                                                                                        }
                                                                                                                                                    case "\n\t":
                                                                                                                                                        v131 = (v131 % ((v131) | 1))
                                                                                                                                                    }
                                                                                                                                                    var v138 int32 = 42
                                                                                                                                                    var v139 int8 = (v131 << ((v20) & 63))
                                                                                                                                                    _ = v138
                                                                                                                                                    _ = v139
                                                                                                                                                case 4294967295:
                                                                                                                                                    var v140 int = (^(v17))
                                                                                                                                                    _ = v140
                                                                                                                                                default:
                                                                                                                                                    delete(v1, 42)
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            v131 -= (v131 << ((255) & 63))
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        if (((v17 << (((^(v20))) & 63)) <= (fi14 * (fi16 << ((255) & 63)))) && ((v20 & v20) <= (v20 * (^(v20))))) {
                                                                                                                                            v141 := GenS6{F0: 236}
                                                                                                                                            _ = v141
                                                                                                                                        } else {
                                                                                                                                            if true {
                                                                                                                                                var v142 uint16 = 42
                                                                                                                                                v3 -= v18
                                                                                                                                                if true {
                                                                                                                                                    switch 4294967295 {
                                                                                                                                                    case -1:
                                                                                                                                                        v20 = v20
                                                                                                                                                    case 2147483647:
                                                                                                                                                        v2 = 1e10
                                                                                                                                                    }
                                                                                                                                                    v1[(v3 * 621)] = v3
                                                                                                                                                } else {
                                                                                                                                                    var v143 int64 = 748
                                                                                                                                                    if true {
                                                                                                                                                        fmt.Println(v10, fi14, fi16)
                                                                                                                                                    } else {
                                                                                                                                                        v144 := GenS13{F0: v142, F1: "a", F2: v143}
                                                                                                                                                        fmt.Println(v18, fi16)
                                                                                                                                                        _ = v144
                                                                                                                                                    }
                                                                                                                                                    var v145 uint = uint(v10)
                                                                                                                                                    _ = v143
                                                                                                                                                    _ = v145
                                                                                                                                                }
                                                                                                                                                var v146 uint64 = 2
                                                                                                                                                v146 = (^(v10))
                                                                                                                                                _ = v142
                                                                                                                                                _ = v146
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        if ((v3 % (((v3 + v3)) | 1)) >= v18) {
                                                                                                                                            delete(v1, (v18 >> (((v20 >> (((^(v20))) & 63))) & 63)))
                                                                                                                                        } else {
                                                                                                                                            v20 = (^(v20))
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    if false {
                                                                                                                                        v13 = append(v13, (v17 ^ (v17 | fi16)))
                                                                                                                                        v10 -= 18446744073709551615
                                                                                                                                    }
                                                                                                                                    _ = len(v1)
                                                                                                                                } else {
                                                                                                                                    var v147 int32 = 0
                                                                                                                                    if true {
                                                                                                                                        v148 := GenS6{F0: 1}
                                                                                                                                        _ = v148
                                                                                                                                    }
                                                                                                                                    v17 -= -473
                                                                                                                                    _ = v147
                                                                                                                                }
                                                                                                                                fmt.Println(v17, v18)
                                                                                                                            }
                                                                                                                        case "hello":
                                                                                                                            var v149 float32 = v2
                                                                                                                            _ = v149
                                                                                                                        default:
                                                                                                                            if true {
                                                                                                                                var v150 uint8 = 110
                                                                                                                                v151 := GenS11{F0: "a", F1: true}
                                                                                                                                fmt.Println(v18, v2)
                                                                                                                                _ = v150
                                                                                                                                _ = v151
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = len(v1)
                                                                                                                        v2 += (((v2 + ((float32(3.14159) + 0.0) * 3.14159)) - -2.25) + (v2 + 0.0))
                                                                                                                    } else {
                                                                                                                        v10 = (v10 << ((v20) & 63))
                                                                                                                        v152 := GenS8{F0: -9223372036854775808, F1: (v20 ^ (^(v20))), F2: 0}
                                                                                                                        v20 += v20
                                                                                                                        if true {
                                                                                                                            v153 := GenS14{F0: (v10 ^ 2), F1: -1}
                                                                                                                            v1[(v18 * (v3 + (v3 / (((v18 - v3)) | 1))))] = v18
                                                                                                                            var v154 uint32 = 65535
                                                                                                                            _ = v153
                                                                                                                            _ = v154
                                                                                                                        }
                                                                                                                        v1[v3] = (v3 ^ 2147483648)
                                                                                                                        var v155 int32 = 42
                                                                                                                        if false {
                                                                                                                            switch (v17 >> (((v20 - (v20 | (v20 >> (((v20 | 255)) & 63))))) & 63)) {
                                                                                                                            case 128:
                                                                                                                                var v156 int32 = -2147483648
                                                                                                                                _ = v156
                                                                                                                            }
                                                                                                                            v13 = append(v13, fi16)
                                                                                                                            switch "hello" {
                                                                                                                            case "\n\t":
                                                                                                                                _ = len(v13)
                                                                                                                                v157 := GenS8{F0: 1, F1: (v20 >> (((v20 << (((^(v20))) & 63))) & 63)), F2: 128}
                                                                                                                                v13 = append(v13, (v17 & (v17 - -2147483648)))
                                                                                                                                _ = v157
                                                                                                                            case "foo bar":
                                                                                                                                if ((v155 / ((v155) | 1)) == v155) {
                                                                                                                                    switch (^(fi14)) {
                                                                                                                                    case -200:
                                                                                                                                        v2 -= v2
                                                                                                                                    case 127:
                                                                                                                                        var v158 uint32 = (v18 + v18)
                                                                                                                                        _ = v158
                                                                                                                                    default:
                                                                                                                                        if true {
                                                                                                                                            switch (("\n\t" + "") + "") {
                                                                                                                                            case "a":
                                                                                                                                                v159 := GenS11{F0: "foo bar", F1: (65535 >= 255)}
                                                                                                                                                _ = v159
                                                                                                                                            case "":
                                                                                                                                                if (((v131 - genFn0(v17, v17)) <= (v131 % ((genFn0((fi16 << ((v20) & 63)), (fi16 - (-(fi14))))) | 1))) || true) {
                                                                                                                                                    if true {
                                                                                                                                                        v131 -= v131
                                                                                                                                                        var v160 float32 = v2
                                                                                                                                                        switch "foo bar" {
                                                                                                                                                        case "":
                                                                                                                                                            if (v160 < 0.1) {
                                                                                                                                                                fmt.Println(v3, v18)
                                                                                                                                                                v131 += (v131 + (v131 + int8(v20)))
                                                                                                                                                                v161 := GenS9{F0: genFn0(fi14, -450)}
                                                                                                                                                                _ = v161
                                                                                                                                                            } else {
                                                                                                                                                                if ((v155 * (^(v155))) <= (v155 - int32(v20))) {
                                                                                                                                                                    if true {
                                                                                                                                                                        if (false && true) {
                                                                                                                                                                            if true {
                                                                                                                                                                                v1[(^(v18))] = (v18 << (((^(v20))) & 63))
                                                                                                                                                                                var v162 int32 = v155
                                                                                                                                                                                _ = v162
                                                                                                                                                                            }
                                                                                                                                                                            v10 = v10
                                                                                                                                                                            v163 := GenS15{F0: ("a" + "")}
                                                                                                                                                                            if len(v13) >= 2 {
                                                                                                                                                                                v13 = v13[1:len(v13)]
                                                                                                                                                                            }
                                                                                                                                                                            if ((fi14 / ((v17) | 1)) == int(v155)) {
                                                                                                                                                                                fmt.Println(v160, v160)
                                                                                                                                                                                v17 = (fi14 / ((int(v3)) | 1))
                                                                                                                                                                                v20 -= v20
                                                                                                                                                                            }
                                                                                                                                                                            _ = v163
                                                                                                                                                                        } else {
                                                                                                                                                                            var v164 int32 = -2147483648
                                                                                                                                                                            _ = v164
                                                                                                                                                                        }
                                                                                                                                                                        v18 = v18
                                                                                                                                                                    } else {
                                                                                                                                                                        if false {
                                                                                                                                                                            v131 += (v131 / ((-78) | 1))
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    if ((-2.25 * ((1.5 * 3.14159) / (1e10 * 1.5))) <= 0.1) {
                                                                                                                                                                        v20 -= 2
                                                                                                                                                                        v2 += ((v2 - (0.1 + (3.14159 + v2))) - (v160 * ((float32(1.5) - v160) * v2)))
                                                                                                                                                                        v1[v3] = uint32(v10)
                                                                                                                                                                    } else {
                                                                                                                                                                        v165 := GenS16{F0: 841, F1: 196, F2: 224}
                                                                                                                                                                        _ = v165
                                                                                                                                                                    }
                                                                                                                                                                    var v166 uint64 = v10
                                                                                                                                                                    _ = v166
                                                                                                                                                                } else {
                                                                                                                                                                    if (!false) {
                                                                                                                                                                        switch "hello" {
                                                                                                                                                                        case "":
                                                                                                                                                                            switch 749 {
                                                                                                                                                                            case 659:
                                                                                                                                                                                var v167 uint32 = 2147483648
                                                                                                                                                                                _ = v167
                                                                                                                                                                            case 390:
                                                                                                                                                                                fmt.Println(fi16)
                                                                                                                                                                            }
                                                                                                                                                                        case "foo bar":
                                                                                                                                                                            if (true && false) {
                                                                                                                                                                                var v168 uint32 = (v18 - v18)
                                                                                                                                                                                _ = v168
                                                                                                                                                                            } else {
                                                                                                                                                                                v160 = float32(-2.25)
                                                                                                                                                                            }
                                                                                                                                                                        case "hello":
                                                                                                                                                                            var v169 uint = (v20 % ((v20) | 1))
                                                                                                                                                                            _ = v169
                                                                                                                                                                        default:
                                                                                                                                                                            v10 -= 796
                                                                                                                                                                        }
                                                                                                                                                                        var v170 uint = 0
                                                                                                                                                                        _ = v170
                                                                                                                                                                    } else {
                                                                                                                                                                        fmt.Println(v17)
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                                v20 -= v20
                                                                                                                                                                v171 := GenS6{F0: 128}
                                                                                                                                                                _ = v171
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        _ = v160
                                                                                                                                                    }
                                                                                                                                                    v17 += 2147483647
                                                                                                                                                    if len(v13) >= 2 {
                                                                                                                                                        v13 = v13[1:len(v13)]
                                                                                                                                                    }
                                                                                                                                                    var v172 float64 = float64(((float32((float64(0.0) / (((1e10 / 1.0) / (-2.25 * 1.5)) * 1e10))) * v2) / float32(3.14159)))
                                                                                                                                                    var v173 bool = true
                                                                                                                                                    _ = v172
                                                                                                                                                    _ = v173
                                                                                                                                                } else {
                                                                                                                                                    fmt.Println(v20, v10)
                                                                                                                                                    if ((v20 - (v20 >> ((v20) & 63))) < 530) {
                                                                                                                                                        fmt.Println(v20)
                                                                                                                                                        v1[v3] = 2
                                                                                                                                                        fmt.Println(v131)
                                                                                                                                                    }
                                                                                                                                                    v174 := GenS17{F0: 42, F1: v10}
                                                                                                                                                    _ = v174
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                v18 = v3
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    var v175 float64 = 0.0
                                                                                                                                    switch "" {
                                                                                                                                    case "a":
                                                                                                                                        switch (fi16 >> (((v20 << (((v20 * v20)) & 63))) & 63)) {
                                                                                                                                        case -336:
                                                                                                                                            var v176 int64 = 1
                                                                                                                                            _ = v176
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        var v177 int32 = 0
                                                                                                                                        _ = v177
                                                                                                                                    }
                                                                                                                                    _ = v175
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = v152
                                                                                                                        _ = v155
                                                                                                                    }
                                                                                                                    v3 -= 4294967295
                                                                                                                    _ = v131
                                                                                                                } else {
                                                                                                                    var v178 uint16 = uint16(v3)
                                                                                                                    var v179 int8 = 42
                                                                                                                    _ = v178
                                                                                                                    _ = v179
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v180 uint32 = (^(v3))
                                                                                                                if ((!true) && (!true)) {
                                                                                                                    if (131 == -9223372036854775808) {
                                                                                                                        v10 -= v10
                                                                                                                        v2 = -2.25
                                                                                                                        var v181 uint16 = 0
                                                                                                                        _ = v181
                                                                                                                    }
                                                                                                                    var v182 int64 = 1
                                                                                                                    _ = v182
                                                                                                                } else {
                                                                                                                    var v183 uint16 = 0
                                                                                                                    var v184 float64 = -2.25
                                                                                                                    v183 -= (v183 | (v183 >> (((v20 - v20)) & 63)))
                                                                                                                    _ = v183
                                                                                                                    _ = v184
                                                                                                                }
                                                                                                                _ = v180
                                                                                                            }
                                                                                                            v185 := GenS5{F0: -564}
                                                                                                            _ = v185
                                                                                                        case 1:
                                                                                                            v18 = (v18 >> ((2) & 63))
                                                                                                        case 693:
                                                                                                            switch (v17 - (^(v17))) {
                                                                                                            case 128:
                                                                                                                var v186 int16 = genFn1((("" + "a") + "a"))
                                                                                                                _ = v186
                                                                                                            case -982:
                                                                                                                v17 -= fi16
                                                                                                            case -9223372036854775808:
                                                                                                                var v187 int16 = 0
                                                                                                                _ = v187
                                                                                                            default:
                                                                                                                v188 := GenS17{F0: 0, F1: v10}
                                                                                                                _ = v188
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = len(v13)
                                                                                                    var v189 int8 = 100
                                                                                                    v10 += (v10 >> ((v20) & 63))
                                                                                                    v190 := GenS18{F0: 128, F1: 42}
                                                                                                    _ = v189
                                                                                                    _ = v190
                                                                                                } else {
                                                                                                    v18 = (v3 | v3)
                                                                                                    v18 -= (v3 & v3)
                                                                                                    var v191 uint64 = v10
                                                                                                    if len(v13) > 0 {
                                                                                                        v13[0] = fi14
                                                                                                    }
                                                                                                    _ = v191
                                                                                                }
                                                                                            } else {
                                                                                                if false {
                                                                                                    v192 := GenS19{F0: v10, F1: 0, F2: 9223372036854775807}
                                                                                                    if (int(v20) == -1) {
                                                                                                        if (!(true && true)) {
                                                                                                            switch ("hello" + "hello") {
                                                                                                            case "":
                                                                                                                v17 = v17
                                                                                                            case "\n\t":
                                                                                                                var v193 uint = v20
                                                                                                                _ = v193
                                                                                                            case "a":
                                                                                                                v18 = v18
                                                                                                                v10 = 0
                                                                                                                fmt.Println(v3, v17, v17)
                                                                                                            }
                                                                                                            v3 = (^(v3))
                                                                                                            v194 := GenS13{F0: 42, F1: "foo bar", F2: 997}
                                                                                                            _ = v194
                                                                                                        } else {
                                                                                                            v195 := v1[v18]
                                                                                                            v195 += v195
                                                                                                            v20 += v20
                                                                                                            _ = v195
                                                                                                        }
                                                                                                        var v196 bool = (772 == 404)
                                                                                                        _ = v196
                                                                                                    }
                                                                                                    fmt.Println(v3, fi14)
                                                                                                    _ = v192
                                                                                                } else {
                                                                                                    var v197 float64 = (1e10 / 1e10)
                                                                                                    switch v17 {
                                                                                                    case 313:
                                                                                                        if false {
                                                                                                            if true {
                                                                                                                var v198 uint = 42
                                                                                                                if (true || ((v20 >> (((^(v198))) & 63)) >= (^(v20)))) {
                                                                                                                    v2 -= (((1e10 / float32(v197)) - 1e10) / (1e10 / 3.14159))
                                                                                                                    v3 -= v18
                                                                                                                } else {
                                                                                                                    if (-2147483648 == -196) {
                                                                                                                        if len(v13) >= 2 {
                                                                                                                            v13 = v13[1:len(v13)]
                                                                                                                        }
                                                                                                                        if true {
                                                                                                                            v199 := GenS18{F0: 255, F1: 9223372036854775807}
                                                                                                                            fmt.Println(v197)
                                                                                                                            fmt.Println(v3, v10)
                                                                                                                            _ = v199
                                                                                                                        } else {
                                                                                                                            if (-484 >= 42) {
                                                                                                                                var v200 int16 = -668
                                                                                                                                v201 := GenS16{F0: uint16(v18), F1: 255, F2: 201}
                                                                                                                                _ = v200
                                                                                                                                _ = v201
                                                                                                                            }
                                                                                                                            fmt.Println(fi16)
                                                                                                                            switch "a" {
                                                                                                                            case "foo bar":
                                                                                                                                delete(v1, v3)
                                                                                                                            case "\n\t":
                                                                                                                                v13 = append(v13, 4294967295)
                                                                                                                            default:
                                                                                                                                var v202 uint = v20
                                                                                                                                _ = v202
                                                                                                                            }
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        switch fi16 {
                                                                                                                        case -743:
                                                                                                                            switch ("foo bar" + (("foo bar" + "") + ("a" + "\n\t"))) {
                                                                                                                            case "foo bar":
                                                                                                                                var v203 float64 = float64(1e10)
                                                                                                                                _ = v203
                                                                                                                            case "hello":
                                                                                                                                var v204 int64 = 920
                                                                                                                                _ = v204
                                                                                                                            }
                                                                                                                        case 128:
                                                                                                                            if true {
                                                                                                                                v10 -= (v10 + v10)
                                                                                                                                var v205 int8 = 127
                                                                                                                                v2 = (v2 / -2.25)
                                                                                                                                _ = v205
                                                                                                                            } else {
                                                                                                                                fmt.Println(v18)
                                                                                                                                var v206 uint64 = (^(v10))
                                                                                                                                _ = v206
                                                                                                                            }
                                                                                                                        case 50:
                                                                                                                            var v207 int = (fi16 | 659)
                                                                                                                            _ = v207
                                                                                                                        default:
                                                                                                                            var v208 uint64 = (v10 + v10)
                                                                                                                            _ = v208
                                                                                                                        }
                                                                                                                        _ = len(v13)
                                                                                                                        v198 -= v198
                                                                                                                    }
                                                                                                                    fmt.Println(v18, v198)
                                                                                                                    switch "foo bar" {
                                                                                                                    case "":
                                                                                                                        fmt.Println(v197)
                                                                                                                    case "hello":
                                                                                                                        if false {
                                                                                                                            fmt.Println(v10, v18)
                                                                                                                            v1[(v18 & (v18 >> ((42) & 63)))] = 2147483648
                                                                                                                            fmt.Println(v2)
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        if len(v13) > 0 {
                                                                                                                            v13[0] = fi14
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                v3 = 4294967295
                                                                                                                _ = v198
                                                                                                            } else {
                                                                                                                v3 = (v3 / ((42) | 1))
                                                                                                            }
                                                                                                            fmt.Println(v3)
                                                                                                            fmt.Println(fi16, v3)
                                                                                                        } else {
                                                                                                            v197 = v197
                                                                                                            var v209 float32 = 3.14159
                                                                                                            _ = v209
                                                                                                        }
                                                                                                    case -2147483648:
                                                                                                        v210 := GenS7{F0: ("hello" + "hello")}
                                                                                                        _ = v210
                                                                                                    }
                                                                                                    _ = v197
                                                                                                }
                                                                                                switch -9223372036854775808 {
                                                                                                case 32767:
                                                                                                    fmt.Println(v2)
                                                                                                case -1:
                                                                                                    _ = len(v1)
                                                                                                }
                                                                                                fmt.Println(fi16)
                                                                                            }
                                                                                        case "hello":
                                                                                            if (-1 < -1) {
                                                                                                v211 := GenS9{F0: 0}
                                                                                                _ = v211
                                                                                            }
                                                                                            var v212 int16 = 0
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                v20 = (v20 - 65535)
                                                                                            default:
                                                                                                if (32767 < 32767) {
                                                                                                    switch "\n\t" {
                                                                                                    case "foo bar":
                                                                                                        var v213 uint = 0
                                                                                                        _ = v213
                                                                                                    case "a":
                                                                                                        v214 := GenS8{F0: 9223372036854775807, F1: v20, F2: 12}
                                                                                                        _ = v214
                                                                                                    case "\n\t":
                                                                                                        var v215 bool = true
                                                                                                        _ = v215
                                                                                                    default:
                                                                                                        v2 -= v2
                                                                                                    }
                                                                                                    v216 := GenS18{F0: 112, F1: -1}
                                                                                                    _ = v216
                                                                                                }
                                                                                            }
                                                                                            _ = v212
                                                                                        case "\n\t":
                                                                                            v217 := GenS17{F0: 445, F1: v10}
                                                                                            _ = v217
                                                                                        default:
                                                                                            var v218 uint = (v20 / ((v20) | 1))
                                                                                            _ = v218
                                                                                        }
                                                                                    default:
                                                                                        var v219 uint32 = v18
                                                                                        _ = v219
                                                                                    }
                                                                                    fmt.Println(v2, v2)
                                                                                    v220 := GenS11{F0: (("foo bar" + "a") + (("\n\t" + "a") + "")), F1: (1 == 240)}
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        v3 = v18
                                                                                    }
                                                                                    _ = v220
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v17, v2)
                                                                            }
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v13 = append(v13, (-(fi16)))
                                                                            default:
                                                                                var v221 int16 = -272
                                                                                _ = v221
                                                                            }
                                                                            v222 := GenS20{F0: "\n\t", F1: (v20 - (v20 << ((18446744073709551615) & 63))), F2: 48}
                                                                            _ = v222
                                                                        }
                                                                        if (v2 >= 3.14159) {
                                                                            fmt.Println(v17)
                                                                        } else {
                                                                            if false {
                                                                                var v223 uint64 = v10
                                                                                if (true && (!true)) {
                                                                                    fmt.Println(v17, v18, v2)
                                                                                    switch ("hello" + "") {
                                                                                    case "":
                                                                                        var v224 int = fi16
                                                                                        _ = v224
                                                                                    case "a":
                                                                                        var v225 uint = 4294967295
                                                                                        _ = v225
                                                                                    default:
                                                                                        v10 = v223
                                                                                    }
                                                                                    if true {
                                                                                        var v226 uint16 = 0
                                                                                        if len(v13) >= 2 {
                                                                                            v13 = v13[1:len(v13)]
                                                                                        }
                                                                                        v18 = v3
                                                                                        _ = v226
                                                                                    }
                                                                                } else {
                                                                                    v1[65535] = (^(v3))
                                                                                    fmt.Println(v17)
                                                                                    v20 = 708
                                                                                }
                                                                                v18 -= v18
                                                                                _ = v223
                                                                            } else {
                                                                                if (645 != -74) {
                                                                                    if ((v3 + v3) != (^(v18))) {
                                                                                        var v227 float32 = v2
                                                                                        v228 := GenS12{F0: 127, F1: true}
                                                                                        _ = v227
                                                                                        _ = v228
                                                                                    } else {
                                                                                        if (v18 == 2147483648) {
                                                                                            var v229 int32 = 1
                                                                                            v230 := GenS6{F0: 128}
                                                                                            _ = v229
                                                                                            _ = v230
                                                                                        } else {
                                                                                            fmt.Println(fi16, v18, v10)
                                                                                        }
                                                                                        v2 -= v2
                                                                                    }
                                                                                    switch 42 {
                                                                                    case 0:
                                                                                        var v231 uint8 = 2
                                                                                        _ = v231
                                                                                    case 2147483647:
                                                                                        switch (-(v17)) {
                                                                                        case 2147483647:
                                                                                            switch "hello" {
                                                                                            case "a":
                                                                                                if (genFn1("a") < 1000) {
                                                                                                    if false {
                                                                                                        if false {
                                                                                                            v20 += (v20 % (((v20 + v20)) | 1))
                                                                                                            switch fi14 {
                                                                                                            case 32767:
                                                                                                                var v232 float64 = 1e10
                                                                                                                _ = v232
                                                                                                            case -723:
                                                                                                                _ = len(v1)
                                                                                                                v13 = append(v13, fi14)
                                                                                                                v3 = v3
                                                                                                            case 32768:
                                                                                                                var v233 uint32 = 903
                                                                                                                _ = v233
                                                                                                            default:
                                                                                                                var v234 uint = (v20 / ((910) | 1))
                                                                                                                _ = v234
                                                                                                            }
                                                                                                            switch fi16 {
                                                                                                            case 4294967295:
                                                                                                                v20 = (v20 / ((v20) | 1))
                                                                                                            case 1:
                                                                                                                v13 = append(v13, fi14)
                                                                                                                if true {
                                                                                                                    switch "\n\t" {
                                                                                                                    case "hello":
                                                                                                                        if (0.0 <= 1.5) {
                                                                                                                            switch ("hello" + ("foo bar" + ("a" + "\n\t"))) {
                                                                                                                            case "a":
                                                                                                                                fmt.Println(fi14)
                                                                                                                                v20 = v20
                                                                                                                                fmt.Println(fi14, fi16, fi14)
                                                                                                                                fmt.Println(v17, fi16, v10)
                                                                                                                            case "foo bar":
                                                                                                                                if true {
                                                                                                                                    v3 = 2147483647
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v235 uint8 = 128
                                                                                                                            _ = v235
                                                                                                                        }
                                                                                                                    case "foo bar":
                                                                                                                        v2 = (-2.25 * (0.1 - 1e10))
                                                                                                                    case "":
                                                                                                                        if (true || ((v20 << ((v20) & 63)) >= (^(v20)))) {
                                                                                                                            switch (fi16 - -796) {
                                                                                                                            case 826:
                                                                                                                                fmt.Println(fi14, v18)
                                                                                                                            case 127:
                                                                                                                                var v236 int32 = -453
                                                                                                                                _ = v236
                                                                                                                            default:
                                                                                                                                var v237 int = -2147483648
                                                                                                                                var v238 int32 = 42
                                                                                                                                _ = v237
                                                                                                                                _ = v238
                                                                                                                            }
                                                                                                                            fmt.Println(v10)
                                                                                                                            v2 += v2
                                                                                                                        } else {
                                                                                                                            var v239 int8 = 1
                                                                                                                            _ = v239
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v240 int = (v17 * 627)
                                                                                                                        _ = v240
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            fmt.Println(v10)
                                                                                                            v10 = v10
                                                                                                        } else {
                                                                                                            v241 := GenS2{F0: v10, F1: 100, F2: 2147483648}
                                                                                                            _ = v241
                                                                                                        }
                                                                                                        var v242 float64 = float64(((-2.25 / (((v2 - (v2 * float32(-2.25))) / v2) * v2)) / v2))
                                                                                                        _ = len(v1)
                                                                                                        _ = v242
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            if len(v13) >= 2 {
                                                                                                                v13 = v13[1:len(v13)]
                                                                                                            }
                                                                                                            fmt.Println(v2, v10, v18)
                                                                                                            var v243 float64 = float64(0.0)
                                                                                                            _ = v243
                                                                                                        } else {
                                                                                                            if (v2 >= (v2 + 3.14159)) {
                                                                                                                switch 32767 {
                                                                                                                case 32767:
                                                                                                                    fmt.Println(v3)
                                                                                                                case 70:
                                                                                                                    fmt.Println(v20, v20)
                                                                                                                case 2147483647:
                                                                                                                    var v244 uint16 = 456
                                                                                                                    _ = v244
                                                                                                                default:
                                                                                                                    if (!false) {
                                                                                                                        fmt.Println(fi14, v20)
                                                                                                                        var v245 int = fi14
                                                                                                                        v246 := GenS6{F0: 128}
                                                                                                                        _ = v245
                                                                                                                        _ = v246
                                                                                                                    }
                                                                                                                }
                                                                                                                var v247 int8 = -118
                                                                                                                switch "" {
                                                                                                                case "foo bar":
                                                                                                                    var v248 uint8 = 128
                                                                                                                    _ = v248
                                                                                                                }
                                                                                                                _ = v247
                                                                                                            }
                                                                                                            if len(v13) > 0 {
                                                                                                                v13[0] = -935
                                                                                                            }
                                                                                                            var v249 string = ((("foo bar" + "\n\t") + "a") + "foo bar")
                                                                                                            _ = v249
                                                                                                        }
                                                                                                        var v250 float32 = v2
                                                                                                        _ = v250
                                                                                                    }
                                                                                                    v2 = (((v2 + (0.1 / v2)) + (v2 / 3.14159)) * 1.5)
                                                                                                } else {
                                                                                                    if len(v13) >= 2 {
                                                                                                        v13 = v13[1:len(v13)]
                                                                                                    }
                                                                                                }
                                                                                            case "":
                                                                                                var v251 int8 = -22
                                                                                                _ = v251
                                                                                            case "\n\t":
                                                                                                var v252 int64 = -662
                                                                                                if len(v13) >= 2 {
                                                                                                    v13 = v13[1:len(v13)]
                                                                                                }
                                                                                                fmt.Println(fi16, v3, v17)
                                                                                                _ = v252
                                                                                            }
                                                                                        case 127:
                                                                                            v17 = (fi16 - fi16)
                                                                                        case -2147483648:
                                                                                            var v253 int8 = -1
                                                                                            _ = v253
                                                                                        default:
                                                                                            v3 -= 2147483647
                                                                                        }
                                                                                    case -2147483648:
                                                                                        v254 := GenS21{F0: 9223372036854775807, F1: -126, F2: true}
                                                                                        _ = v254
                                                                                    }
                                                                                    switch 9223372036854775807 {
                                                                                    case 127:
                                                                                        fmt.Println(v17)
                                                                                    case 0:
                                                                                        switch (fi16 ^ (fi16 & 1)) {
                                                                                        case 128:
                                                                                            if ((v20 >> ((v20) & 63)) > v20) {
                                                                                                switch fi14 {
                                                                                                case -335:
                                                                                                    v13 = append(v13, (fi16 ^ fi14))
                                                                                                case 32767:
                                                                                                    v2 += (v2 + (v2 - (((v2 - float32(((1e10 * (0.1 + (3.14159 / (3.14159 / (1.5 - (-2.25 * (float64((-2.25 + 1e10)) * -2.25))))))) + 1.5))) - v2) / float32(-2.25))))
                                                                                                    if ((3.14159 - 3.14159) < 1.5) {
                                                                                                        switch "foo bar" {
                                                                                                        case "hello":
                                                                                                            fmt.Println(fi14, v3)
                                                                                                            var v255 int8 = 127
                                                                                                            var v256 int16 = 42
                                                                                                            switch fi16 {
                                                                                                            case 0:
                                                                                                                switch "" {
                                                                                                                case "foo bar":
                                                                                                                    var v257 int64 = 0
                                                                                                                    _ = v257
                                                                                                                default:
                                                                                                                    var v258 int32 = -583
                                                                                                                    _ = v258
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v255
                                                                                                            _ = v256
                                                                                                        case "":
                                                                                                            v259 := GenS22{F0: (v20 + 4294967295), F1: false, F2: (2 < 2147483648)}
                                                                                                            _ = v259
                                                                                                        }
                                                                                                        if false {
                                                                                                            switch "a" {
                                                                                                            case "":
                                                                                                                var v260 float64 = -2.25
                                                                                                                _ = v260
                                                                                                            case "foo bar":
                                                                                                                v20 = (v20 << (((^(v20))) & 63))
                                                                                                            case "a":
                                                                                                                var v261 uint16 = 830
                                                                                                                _ = v261
                                                                                                            default:
                                                                                                                var v262 bool = false
                                                                                                                _ = v262
                                                                                                            }
                                                                                                            if true {
                                                                                                                if len(v13) >= 2 {
                                                                                                                    v13 = v13[1:len(v13)]
                                                                                                                }
                                                                                                                var v263 uint32 = v18
                                                                                                                if len(v13) >= 2 {
                                                                                                                    v13 = v13[1:len(v13)]
                                                                                                                }
                                                                                                                v18 = (v18 - (v18 << (((v20 + v20)) & 63)))
                                                                                                                _ = v263
                                                                                                            } else {
                                                                                                                var v264 uint = (v20 / ((195) | 1))
                                                                                                                v265 := GenS23{F0: 2, F1: 1}
                                                                                                                v1[uint32(v20)] = (v3 / ((2147483647) | 1))
                                                                                                                v17 = v17
                                                                                                                v17 -= (fi14 / (((fi14 / ((fi16) | 1))) | 1))
                                                                                                                _ = v264
                                                                                                                _ = v265
                                                                                                            }
                                                                                                            v20 += v20
                                                                                                            var v266 int32 = 1
                                                                                                            _ = v266
                                                                                                        }
                                                                                                    } else {
                                                                                                        switch "hello" {
                                                                                                        case "hello":
                                                                                                            fmt.Println(fi16, fi16)
                                                                                                        default:
                                                                                                            var v267 uint32 = v3
                                                                                                            _ = v267
                                                                                                        }
                                                                                                        fmt.Println(v18, fi16)
                                                                                                    }
                                                                                                default:
                                                                                                    v18 -= v18
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v2)
                                                                                            }
                                                                                        case 4294967295:
                                                                                            var v268 int8 = -128
                                                                                            _ = v268
                                                                                        }
                                                                                    default:
                                                                                        v269 := v1[(v3 & (v18 & (^(v18))))]
                                                                                        _ = v269
                                                                                    }
                                                                                }
                                                                                v17 = int(v20)
                                                                            }
                                                                            v20 -= v20
                                                                            var v270 int8 = -9
                                                                            if len(v13) >= 2 {
                                                                                v13 = v13[1:len(v13)]
                                                                            }
                                                                            if len(v13) > 0 {
                                                                                v13[0] = v17
                                                                            }
                                                                            var v271 bool = ((-1 >= 9223372036854775807) || false)
                                                                            v1[(v18 + 442)] = 0
                                                                            if len(v13) > 0 {
                                                                                v13[0] = (fi16 & (fi14 >> ((v20) & 63)))
                                                                            }
                                                                            v272 := GenS2{F0: uint64(v18), F1: 127, F2: v3}
                                                                            if v271 {
                                                                                switch (fi16 % ((v17) | 1)) {
                                                                                case 2147483647:
                                                                                    v17 += (fi14 & (fi16 & 32767))
                                                                                case 127:
                                                                                    switch ("" + "a") {
                                                                                    case "\n\t":
                                                                                        switch (("foo bar" + ("" + "")) + "foo bar") {
                                                                                        case "\n\t":
                                                                                            fmt.Println(v271)
                                                                                        }
                                                                                    }
                                                                                    var v273 float64 = 0.1
                                                                                    var v274 uint32 = (v3 ^ (v3 - (v18 >> ((479) & 63))))
                                                                                    _ = v273
                                                                                    _ = v274
                                                                                default:
                                                                                    switch 0 {
                                                                                    case -960:
                                                                                        delete(v1, v18)
                                                                                    case -24:
                                                                                        v17 += v17
                                                                                    default:
                                                                                        if (255 <= 153) {
                                                                                            v17 = -1
                                                                                        }
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v275 int16 = -32768
                                                                                _ = v275
                                                                            }
                                                                            _ = v270
                                                                            _ = v271
                                                                            _ = v272
                                                                        }
                                                                    }
                                                                    v2 += (v2 / v2)
                                                                }
                                                                switch fi16 {
                                                                case 547:
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v2 = (v2 + float32(-2.25))
                                                                    case "a":
                                                                        v276 := GenS24{F0: ((("" + "a") + "foo bar") + "a")}
                                                                        _ = v276
                                                                    default:
                                                                        var v277 int16 = 32767
                                                                        _ = v277
                                                                    }
                                                                case 127:
                                                                    switch "a" {
                                                                    case "hello":
                                                                        v278 := GenS25{F0: v20}
                                                                        _ = v278
                                                                    }
                                                                case 0:
                                                                    var v279 int16 = genFn1("hello")
                                                                    _ = v279
                                                                default:
                                                                    if len(v13) > 0 {
                                                                        v13[0] = -9223372036854775808
                                                                    }
                                                                }
                                                            }
                                                            switch (fi14 / ((fi14) | 1)) {
                                                            case 42:
                                                                v280 := v1[(v18 / ((2) | 1))]
                                                                if (1 == genFn0(0, fi16)) {
                                                                    v2 = 0.1
                                                                } else {
                                                                    if true {
                                                                        if (-1 < (fi14 | 2147483647)) {
                                                                            if (915 != 856) {
                                                                                if true {
                                                                                    var v281 int32 = -1
                                                                                    var v282 float32 = (float32(0.1) - (float32(float64(v2)) * 3.14159))
                                                                                    _ = v281
                                                                                    _ = v282
                                                                                } else {
                                                                                    if (42 > -2147483648) {
                                                                                        v283 := GenS2{F0: (^(v10)), F1: 255, F2: v3}
                                                                                        _ = len(v1)
                                                                                        _ = v283
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi16, v18)
                                                                            } else {
                                                                                var v284 uint32 = v280
                                                                                _ = v284
                                                                            }
                                                                            fmt.Println(v3, v2)
                                                                            switch ("foo bar" + "") {
                                                                            case "hello":
                                                                                switch "a" {
                                                                                case "":
                                                                                    _ = len(v13)
                                                                                case "a":
                                                                                    switch (fi16 & (fi14 % ((int(v20)) | 1))) {
                                                                                    case 35:
                                                                                        v285 := GenS2{F0: 744, F1: 2, F2: (v3 ^ (v3 | 1))}
                                                                                        _ = v285
                                                                                    case 1:
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            fmt.Println(v3)
                                                                                        case "a":
                                                                                            v280 = 7
                                                                                        case "hello":
                                                                                            if (!(1.5 >= v2)) {
                                                                                                var v286 uint32 = v18
                                                                                                _ = v286
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v280, v3, v18)
                                                                                        }
                                                                                    case 2147483647:
                                                                                        if true {
                                                                                            var v287 float32 = v2
                                                                                            _ = v287
                                                                                        } else {
                                                                                            _ = len(v13)
                                                                                            var v288 int32 = 1
                                                                                            _ = v288
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v280, v3, v280)
                                                                                    }
                                                                                case "foo bar":
                                                                                    delete(v1, (v280 & (v280 | 326)))
                                                                                    v2 = float32(1e10)
                                                                                default:
                                                                                    v3 += v18
                                                                                }
                                                                            }
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                var v289 int16 = -1
                                                                                _ = v289
                                                                            default:
                                                                                fmt.Println(v17)
                                                                            }
                                                                            var v290 uint64 = v10
                                                                            _ = v290
                                                                        }
                                                                        if true {
                                                                            var v291 uint16 = 255
                                                                            v292 := GenS11{F0: "\n\t", F1: false}
                                                                            _ = v291
                                                                            _ = v292
                                                                        }
                                                                    }
                                                                    _ = len(v1)
                                                                }
                                                                v280 = (v280 & uint32(v20))
                                                                _ = v280
                                                            case -2147483648:
                                                                v20 -= (v20 + 0)
                                                            }
                                                            _ = v20
                                                        } else {
                                                            if len(v13) > 0 {
                                                                v13[0] = 0
                                                            }
                                                            switch ("\n\t" + "") {
                                                            case "a":
                                                                v3 = v3
                                                            }
                                                        }
                                                    }
                                                    v18 = 2147483647
                                                }
                                            }
                                            var v293 string = ""
                                            v294 := GenS2{F0: 18446744073709551615, F1: 42, F2: (v18 & (v18 & (v3 ^ (^(v18)))))}
                                            var v295 float32 = (3.14159 - v2)
                                            var v296 bool = (100 >= 42)
                                            _ = v293
                                            _ = v294
                                            _ = v295
                                            _ = v296
                                        case -735:
                                            if (255 < 4294967295) {
                                                v2 -= v2
                                            }
                                        case 127:
                                            if (false && true) {
                                                var v297 uint32 = v18
                                                var v298 int16 = genFn1((("hello" + "") + ""))
                                                fmt.Println(v298, v18, v298)
                                                _ = v297
                                                _ = v298
                                            }
                                            v299 := GenS19{F0: 970, F1: 42, F2: fi14}
                                            var v300 uint16 = uint16(v10)
                                            _ = v299
                                            _ = v300
                                        default:
                                            v17 = (fi14 ^ (fi16 ^ 728))
                                        }
                                    case -9223372036854775808:
                                        v301 := GenS23{F0: 0, F1: 236}
                                        _ = v301
                                    case 941:
                                        var v302 uint64 = (v10 % (((^(v10))) | 1))
                                        _ = v302
                                    }
                                }
                                _ = v18
                            }
                            var v303 int64 = -51
                            v3 = 1
                            _ = fi16
                            _ = v17
                            _ = v303
                        }
                        _ = fi14
                        _ = v15
                    }
                case 128:
                    switch "foo bar" {
                    case "hello":
                        v304 := GenS26{F0: v10}
                        _ = v304
                    case "foo bar":
                        var v305 uint64 = v10
                        _ = v305
                    case "\n\t":
                        fmt.Println(v10, v10, v10)
                    }
                default:
                    var v306 bool = (true || (35 >= -31))
                    _ = v306
                }
            case "hello":
                var v307 uint32 = (v3 - (^(v3)))
                _ = v307
            case "foo bar":
                if (2 < uint(v3)) {
                    v10 = v10
                    v308 := GenS20{F0: ("hello" + ("a" + ("\n\t" + ((("foo bar" + "") + "a") + "a")))), F1: 65535, F2: 1}
                    _ = v308
                } else {
                    fmt.Println(v3, v3, v3)
                }
            }
            var v309 int16 = 32767
            v2 += (v2 - -2.25)
            _ = v309
        }
    case "\n\t":
        var v310 float32 = (((v2 * v2) / v2) / float32((((0.1 / 1.5) + (3.14159 - (1e10 * float64(v2)))) + (0.0 - 3.14159))))
        _ = v310
    default:
        v2 = v2
    }
    v311 := GenS18{F0: 1, F1: -1}
    v10 -= (v10 * v10)
    v2 -= float32(1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v311
    _ = v9
}
