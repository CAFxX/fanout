package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint; F1 int8; F2 uint64 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 uint64; F1 int64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int16; F1 int8 }
type GenS12 struct { F0 int16; F1 int32 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint64; F1 string }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int8; F1 int32; F2 uint8 }
type GenS17 struct { F0 int64; F1 int; F2 uint8 }
type GenS18 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 int16; F1 uint }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 int16; F1 int32 }
type GenS24 struct { F0 uint32; F1 int }
func genFn0(p0 string) uint8 {
    v4 := GenS1{F0: 42, F1: 42}
    v5 := GenS1{F0: 6, F1: -795}
    _ = p0
    _ = v4
    _ = v5
    return 255
}
func genFn1(p0 int32, p1 uint16) bool {
    var v6 int = 1
    fmt.Println(v6, p1)
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = v6
    return false
}

func main() {
    var v1 uint16 = 32767
    var v2 float64 = 0.1
    v3 := []string{("a" + "\n\t")}
    var v7 int16 = -32768
    switch "" {
    case "hello":
        var v8 int = 2147483647
        _ = v8
    case "foo bar":
        v9 := GenS1{F0: 18446744073709551615, F1: -379}
        _ = v9
    case "a":
        switch ((("a" + "a") + "foo bar") + (("foo bar" + ((("\n\t" + "foo bar") + "\n\t") + "hello")) + "hello")) {
        case "foo bar":
            var v10 uint = 4294967295
            _ = v10
        case "\n\t":
            switch -1 {
            case 268:
                var v11 float32 = float32(-2.25)
                _ = v11
            case -1:
                if (0 <= 0) {
                    if false {
                        for fi12 := 0; fi12 < 14; fi12++ {
                            var v13 float32 = 1e10
                            switch "foo bar" {
                            case "a":
                                var v14 uint32 = 368
                                _ = v14
                            case "\n\t":
                                v15 := GenS1{F0: 0, F1: 42}
                                _ = v15
                            case "hello":
                                fmt.Println(v13)
                            default:
                                switch (fi12 << ((18446744073709551615) & 63)) {
                                case -47:
                                    if true {
                                        for fi16 := 0; fi16 < 19; fi16++ {
                                            _ = len(v3)
                                            fmt.Println(fi12, v7)
                                            v17 := GenS2{F0: -950}
                                            _ = fi16
                                            _ = v17
                                        }
                                        if false {
                                            v18 := GenS2{F0: -441}
                                            fmt.Println(v1, v1)
                                            v19 := GenS3{F0: 0}
                                            _ = v18
                                            _ = v19
                                        }
                                    } else {
                                        v1 += 359
                                        v3 = append(v3, "")
                                    }
                                case 4294967295:
                                    switch (fi12 % ((2147483647) | 1)) {
                                    case 891:
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                    case 127:
                                        var v20 int8 = 100
                                        _ = v20
                                    default:
                                        _ = len(v3)
                                    }
                                case -2147483648:
                                    _ = len(v3)
                                }
                            }
                            var v21 uint8 = 113
                            _ = fi12
                            _ = v13
                            _ = v21
                        }
                        var v22 int = 971
                        _ = v22
                    } else {
                        v23 := GenS4{F0: 283, F1: 459}
                        switch ("hello" + "hello") {
                        case "a":
                            fmt.Println(v1)
                        default:
                            v7 = (^(v7))
                        }
                        v2 += 3.14159
                        _ = v23
                    }
                    _ = len(v3)
                    v24 := GenS4{F0: 567, F1: 18446744073709551615}
                    _ = v24
                }
            default:
                var v25 uint32 = 581
                _ = v25
            }
        default:
            if false {
                v26 := GenS5{F0: 1, F1: 33, F2: 4294967295}
                _ = v26
            } else {
                var v27 uint32 = uint32(v1)
                if false {
                    v28 := GenS3{F0: 52}
                    v27 = 2147483648
                    v1 += v1
                    _ = v28
                }
                _ = v27
            }
        }
    }
    v3 = append(v3, ("hello" + ("\n\t" + "")))
    v2 -= v2
    _ = len(v3)
    v1 = (^(v1))
    v3 = append(v3, "a")
    switch 127 {
    case 795:
        v29 := GenS6{F0: "hello", F1: 4294967295}
        _ = v29
    case 111:
        for fi30 := 0; fi30 < 7; fi30++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 = (v2 - v2)
            _ = fi30
        }
    case 32768:
        for fi31 := 0; fi31 < 16; fi31++ {
            if false {
                _ = len(v3)
            }
            for fi32 := 0; fi32 < 3; fi32++ {
                v33 := GenS6{F0: (("foo bar" + "a") + ""), F1: 255}
                v7 -= (^(v7))
                v1 += 1
                _ = fi32
                _ = v33
            }
            v7 -= v7
            _ = fi31
        }
    default:
        switch 127 {
        case 32767:
            for fi34 := 0; fi34 < 9; fi34++ {
                v2 -= 3.14159
                v2 -= -2.25
                _ = fi34
            }
        case 0:
            v2 -= float64(3.14159)
        case -525:
            var v35 int32 = 410
            _ = v35
        default:
            _ = len(v3)
        }
    }
    for fi36 := 0; fi36 < 15; fi36++ {
        if false {
            v37 := GenS7{F0: 1, F1: 109}
            _ = v37
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch ("" + "foo bar") {
            case "foo bar":
                var v38 int = 4294967295
                _ = v38
            default:
                v39 := GenS8{F0: -94, F1: -107}
                _ = v39
            }
        }
        _ = fi36
    }
    if (42 >= int64(v1)) {
        v40 := GenS2{F0: 2147483647}
        for fi41 := 0; fi41 < 19; fi41++ {
            v7 = v7
            var v42 int64 = 1
            _ = fi41
            _ = v42
        }
        v1 = v1
        _ = v40
    }
    switch 128 {
    case 2147483647:
        fmt.Println(v7, v7, v7)
    case -2147483648:
        for fi43 := 0; fi43 < 1; fi43++ {
            var v44 uint16 = (v1 << ((373) & 63))
            v45 := GenS7{F0: 490, F1: 38}
            var v46 int32 = 42
            _ = fi43
            _ = v44
            _ = v45
            _ = v46
        }
    case -975:
        v3 = append(v3, "a")
    default:
        if (4294967295 > 4294967295) {
            v47 := GenS9{F0: 2, F1: 699}
            for fi48 := 0; fi48 < 11; fi48++ {
                switch "" {
                case "":
                    switch "" {
                    case "a":
                        if (1 > 42) {
                            _ = len(v3)
                        }
                    case "\n\t":
                        v1 += (v1 + v1)
                    case "foo bar":
                        v49 := GenS7{F0: uint64(v1), F1: 255}
                        _ = v49
                    default:
                        v50 := GenS10{F0: 255}
                        _ = v50
                    }
                case "foo bar":
                    for fi51 := 0; fi51 < 8; fi51++ {
                        switch "" {
                        case "\n\t":
                            if len(v3) > 0 {
                                v3[0] = ""
                            }
                            if (!(true && (!true))) {
                                v2 = 0.0
                                v1 -= (^(v1))
                            } else {
                                if (((197 < 571) || (1 < 127)) && (9223372036854775807 >= -294)) {
                                    _ = len(v3)
                                    switch ("\n\t" + "foo bar") {
                                    case "a":
                                        var v52 uint16 = (v1 << ((2) & 63))
                                        _ = v52
                                    case "hello":
                                        fmt.Println(fi48, fi51)
                                    case "\n\t":
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                        v1 -= (^(v1))
                                    default:
                                        if (!(18446744073709551615 > 4294967295)) {
                                            switch (fi48 >> ((431) & 63)) {
                                            case 4294967295:
                                                if true {
                                                    switch (fi48 % ((166) | 1)) {
                                                    case -1:
                                                        if len(v3) > 0 {
                                                            v3[0] = ("foo bar" + "")
                                                        }
                                                    case 1:
                                                        v53 := GenS8{F0: int8(v7), F1: 100}
                                                        _ = v53
                                                    case 571:
                                                        if false {
                                                            if ((^(v1)) > (v1 ^ 123)) {
                                                                if false {
                                                                    v54 := GenS8{F0: -56, F1: 39}
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v55 := GenS11{F0: (v7 - 1), F1: 1}
                                                                        _ = v55
                                                                    default:
                                                                        switch ("hello" + "a") {
                                                                        case "":
                                                                            fmt.Println(v1)
                                                                        case "a":
                                                                            if (-39 >= (-(v7))) {
                                                                                v7 = 0
                                                                            } else {
                                                                                if false {
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    v1 = (^(v1))
                                                                                    if (!true) {
                                                                                        v7 -= v7
                                                                                        var v56 int = fi51
                                                                                        var v57 float64 = (v2 / (v2 - (v2 * v2)))
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = "a"
                                                                                        }
                                                                                        var v58 bool = (!true)
                                                                                        _ = v56
                                                                                        _ = v57
                                                                                        _ = v58
                                                                                    } else {
                                                                                        v59 := GenS12{F0: (v7 | (v7 ^ v7)), F1: 865}
                                                                                        var v60 string = ("" + (("foo bar" + "") + (("hello" + ("" + "")) + "foo bar")))
                                                                                        _ = len(v3)
                                                                                        v2 -= 1e10
                                                                                        v61 := GenS7{F0: 2, F1: 128}
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                        _ = v61
                                                                                    }
                                                                                }
                                                                                v62 := GenS7{F0: 1, F1: 127}
                                                                                _ = v62
                                                                            }
                                                                        case "foo bar":
                                                                            v3 = append(v3, "a")
                                                                        }
                                                                    }
                                                                    var v63 int = (fi48 << ((889) & 63))
                                                                    v2 = float64(0.1)
                                                                    _ = v54
                                                                    _ = v63
                                                                } else {
                                                                    v1 = (v1 / (((v1 << ((1) & 63))) | 1))
                                                                }
                                                                v2 += v2
                                                            } else {
                                                                if false {
                                                                    v7 = (v7 << ((18446744073709551615) & 63))
                                                                    v64 := GenS13{F0: v7}
                                                                    var v65 int8 = 0
                                                                    _ = v64
                                                                    _ = v65
                                                                }
                                                            }
                                                            if true {
                                                                switch "" {
                                                                case "\n\t":
                                                                    v7 = v7
                                                                case "hello":
                                                                    v2 += v2
                                                                default:
                                                                    if true {
                                                                        if false {
                                                                            v7 = (-(v7))
                                                                            v66 := GenS3{F0: -1}
                                                                            _ = v66
                                                                        }
                                                                        var v67 int8 = -38
                                                                        _ = v67
                                                                    } else {
                                                                        if true {
                                                                            v2 = v2
                                                                            v3 = append(v3, "hello")
                                                                        }
                                                                    }
                                                                }
                                                                switch (fi48 % ((42) | 1)) {
                                                                case 127:
                                                                    v68 := GenS14{F0: 0, F1: ""}
                                                                    _ = v68
                                                                }
                                                                v7 += -908
                                                                v69 := GenS13{F0: 0}
                                                                v7 = int16(v1)
                                                                _ = v69
                                                            } else {
                                                                v2 = 1.5
                                                            }
                                                            if false {
                                                                switch (fi51 + -538) {
                                                                case 617:
                                                                    if (128 == 2) {
                                                                        if (true || true) {
                                                                            v1 -= v1
                                                                            switch ("" + "foo bar") {
                                                                            case "a":
                                                                                _ = len(v3)
                                                                                if false {
                                                                                    v1 = (v1 ^ (^(v1)))
                                                                                }
                                                                                v2 = float64(1e10)
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    if (0 >= 292) {
                                                                                        fmt.Println(v1, v7)
                                                                                    }
                                                                                case "a":
                                                                                    fmt.Println(v1, fi48)
                                                                                }
                                                                                switch (("hello" + (("foo bar" + "\n\t") + "")) + ((("\n\t" + "foo bar") + "\n\t") + ("foo bar" + "a"))) {
                                                                                case "a":
                                                                                    v1 = (v1 >> ((347) & 63))
                                                                                }
                                                                            case "hello":
                                                                                fmt.Println(fi51, v1, v1)
                                                                            case "\n\t":
                                                                                if ((65535 == 506) && false) {
                                                                                    v70 := GenS15{F0: 2}
                                                                                    if false {
                                                                                        switch "a" {
                                                                                        case "":
                                                                                            v7 += (v7 - v7)
                                                                                        }
                                                                                        v71 := GenS16{F0: 66, F1: 42, F2: 0}
                                                                                        _ = v71
                                                                                    } else {
                                                                                        v1 += v1
                                                                                        v2 += 1.5
                                                                                        if false {
                                                                                            if (!false) {
                                                                                                v2 = 1e10
                                                                                            } else {
                                                                                                v72 := GenS17{F0: 9223372036854775807, F1: fi51, F2: 115}
                                                                                                _ = v72
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v70
                                                                                } else {
                                                                                    var v73 int8 = -128
                                                                                    if false {
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = "a"
                                                                                        }
                                                                                        fmt.Println(fi48, v7)
                                                                                        var v74 uint64 = 42
                                                                                        v75 := GenS10{F0: 432}
                                                                                        v76 := GenS7{F0: 4294967295, F1: 217}
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = (("\n\t" + (("foo bar" + "a") + ("a" + "a"))) + "")
                                                                                        }
                                                                                        switch (fi48 * fi51) {
                                                                                        case 2147483647:
                                                                                            v77 := GenS15{F0: 18446744073709551615}
                                                                                            _ = v77
                                                                                        }
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                        _ = v74
                                                                                        _ = v75
                                                                                        _ = v76
                                                                                    } else {
                                                                                        var v78 int16 = v7
                                                                                        _ = v78
                                                                                    }
                                                                                    var v79 int16 = v7
                                                                                    _ = v73
                                                                                    _ = v79
                                                                                }
                                                                            }
                                                                            v80 := GenS18{F0: -9223372036854775808, F1: 4294967295, F2: 618}
                                                                            _ = v80
                                                                        }
                                                                        var v81 int32 = -348
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (("foo bar" + "") + "hello")
                                                                        }
                                                                        _ = v81
                                                                    } else {
                                                                        var v82 uint = 718
                                                                        var v83 int32 = 97
                                                                        _ = v82
                                                                        _ = v83
                                                                    }
                                                                case 32768:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    switch fi51 {
                                                                    case 9223372036854775807:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "\n\t"
                                                                        }
                                                                    case 580:
                                                                        v84 := GenS19{F0: 255}
                                                                        _ = v84
                                                                    case 32768:
                                                                        var v85 int64 = 42
                                                                        _ = v85
                                                                    default:
                                                                        fmt.Println(fi48)
                                                                        if (v2 != (v2 / float64(-2.25))) {
                                                                            v1 = (^(v1))
                                                                            v7 += (v7 + (v7 % ((v7) | 1)))
                                                                            fmt.Println(v1)
                                                                        }
                                                                    }
                                                                    if len(v3) > 0 {
                                                                        v3[0] = "a"
                                                                    }
                                                                default:
                                                                    v86 := GenS12{F0: v7, F1: 1}
                                                                    _ = v86
                                                                }
                                                                fmt.Println(v2)
                                                            } else {
                                                                var v87 float32 = (1e10 / 1e10)
                                                                _ = v87
                                                            }
                                                        } else {
                                                            if true {
                                                                var v88 int8 = int8(v1)
                                                                var v89 int16 = (^(v7))
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                            var v90 int = fi48
                                                            _ = v90
                                                        }
                                                    }
                                                }
                                            }
                                            v91 := GenS6{F0: "a", F1: 851}
                                            _ = v91
                                        }
                                    }
                                    if (!(195 <= 36)) {
                                        var v92 uint32 = 585
                                        _ = v92
                                    }
                                } else {
                                    var v93 float64 = 1.5
                                    v93 += float64(float32((v93 * v93)))
                                    if false {
                                        v1 += 0
                                    }
                                    _ = v93
                                }
                                v94 := GenS3{F0: 1}
                                _ = v94
                            }
                            v2 = 1e10
                        case "hello":
                            v7 += 32767
                        case "":
                            v95 := GenS17{F0: 1, F1: int(v7), F2: 2}
                            _ = v95
                        }
                        fmt.Println(fi48, fi51)
                        if len(v3) > 0 {
                            v3[0] = ""
                        }
                        _ = fi51
                    }
                case "hello":
                    for fi96 := 0; fi96 < 4; fi96++ {
                        v97 := GenS12{F0: v7, F1: -2147483648}
                        v7 -= (v7 * (v7 * 42))
                        var v98 int16 = (v7 + 42)
                        _ = fi96
                        _ = v97
                        _ = v98
                    }
                }
                switch (fi48 + (fi48 & fi48)) {
                case -833:
                    var v99 int8 = 42
                    _ = v99
                case 4294967295:
                    if false {
                        var v100 float64 = ((v2 + float64(1.5)) + v2)
                        if ("foo bar" >= ("\n\t" + ("hello" + "\n\t"))) {
                            v100 += -2.25
                            if false {
                                if false {
                                    if ((true && true) || (2 >= 42)) {
                                        v101 := GenS5{F0: 117, F1: 1, F2: 1}
                                        _ = v101
                                    } else {
                                        var v102 int16 = v7
                                        _ = v102
                                    }
                                    var v103 int32 = -2147483648
                                    _ = v103
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    v2 = float64(0.1)
                                default:
                                    for fi104 := 0; fi104 < 4; fi104++ {
                                        if (false && (false || true)) {
                                            v105 := GenS6{F0: "a", F1: 18446744073709551615}
                                            fmt.Println(v100, v100)
                                            _ = v105
                                        } else {
                                            if true {
                                                var v106 uint32 = 1
                                                var v107 float32 = (1e10 - 0.0)
                                                if len(v3) > 0 {
                                                    v3[0] = (("a" + ("hello" + "foo bar")) + "")
                                                }
                                                _ = v106
                                                _ = v107
                                            }
                                        }
                                        v108 := GenS19{F0: 2}
                                        v109 := GenS12{F0: v7, F1: 2147483647}
                                        _ = fi104
                                        _ = v108
                                        _ = v109
                                    }
                                }
                            } else {
                                for fi110 := 0; fi110 < 5; fi110++ {
                                    v1 += (v1 | (v1 << ((4294967295) & 63)))
                                    var v111 int64 = 0
                                    if ((16 >= 5) && false) {
                                        v2 -= v2
                                        v3 = append(v3, "hello")
                                        var v112 uint16 = (^(v1))
                                        _ = v112
                                    } else {
                                        switch "a" {
                                        case "hello":
                                            if (255 != (v1 & v1)) {
                                                v1 = (v1 - v1)
                                            }
                                        case "foo bar":
                                            v111 = v111
                                        default:
                                            var v113 int16 = v7
                                            _ = v113
                                        }
                                        v114 := GenS15{F0: 4294967295}
                                        _ = v114
                                    }
                                    _ = fi110
                                    _ = v111
                                }
                            }
                        } else {
                            _ = len(v3)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v3 = append(v3, ("" + ""))
                        }
                        var v115 uint16 = v1
                        _ = v100
                        _ = v115
                    }
                case 810:
                    var v116 int32 = 809
                    _ = v116
                }
                var v117 uint8 = genFn0("foo bar")
                _ = fi48
                _ = v117
            }
            var v118 uint32 = 927
            _ = v118
            _ = v47
        }
    }
    switch 9223372036854775807 {
    case 374:
        fmt.Println(v7, v7)
    case 502:
        fmt.Println(v2, v7, v7)
    case 127:
        if false {
            var v119 bool = false
            var v120 int32 = -1
            _ = v119
            _ = v120
        } else {
            v121 := GenS20{F0: 39}
            if (!false) {
                v122 := GenS21{F0: (v7 - (v7 % ((v7) | 1))), F1: 65535}
                if true {
                    var v123 float32 = (0.1 + (1.5 - -2.25))
                    if (4294967295 >= 4294967295) {
                        v123 = (0.0 - v123)
                        switch 32768 {
                        case 2147483647:
                            v2 = (((((float64(v123) - (v2 + float64(1e10))) / 1.0) / 1.0) - (float64(0.1) * (3.14159 + -2.25))) * v2)
                        }
                    }
                    _ = v123
                }
                _ = len(v3)
                _ = v122
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "" {
                case "a":
                    switch 128 {
                    case -362:
                        var v124 int32 = 255
                        _ = v124
                    case 500:
                        var v125 int = 127
                        _ = v125
                    default:
                        switch "" {
                        case "\n\t":
                            for fi126 := 0; fi126 < 1; fi126++ {
                                fmt.Println(v2, fi126)
                                fmt.Println(v7, fi126)
                                _ = fi126
                            }
                        case "a":
                            for fi127 := 0; fi127 < 3; fi127++ {
                                v1 += v1
                                v2 -= 0.0
                                v2 += v2
                                _ = fi127
                            }
                        }
                    }
                case "":
                    v128 := GenS10{F0: 255}
                    _ = v128
                case "\n\t":
                    v129 := GenS22{F0: ((!(1 != 42)) || true)}
                    _ = v129
                default:
                    if (!(1 != 0)) {
                        if false {
                            if true {
                                v3 = append(v3, "foo bar")
                                fmt.Println(v1, v1)
                            }
                            v3 = append(v3, "a")
                            if true {
                                if false {
                                    for fi130 := 0; fi130 < 16; fi130++ {
                                        fmt.Println(fi130, v2, v2)
                                        _ = len(v3)
                                        var v131 int32 = 2147483647
                                        _ = fi130
                                        _ = v131
                                    }
                                    v7 -= (v7 / ((-1) | 1))
                                    if (1 != 2) {
                                        switch "" {
                                        case "foo bar":
                                            if len(v3) > 0 {
                                                v3[0] = "foo bar"
                                            }
                                        }
                                        var v132 int32 = 42
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v132
                                    }
                                }
                                v133 := GenS23{F0: v7, F1: -935}
                                v2 -= v2
                                _ = v133
                            }
                        } else {
                            var v134 uint16 = 255
                            _ = v134
                        }
                        if false {
                            for fi135 := 0; fi135 < 19; fi135++ {
                                v1 = (^(v1))
                                var v136 string = ""
                                v136 = ("" + v136)
                                _ = fi135
                                _ = v136
                            }
                        }
                    }
                }
                for fi137 := 0; fi137 < 17; fi137++ {
                    v2 -= float64((3.14159 / -2.25))
                    switch (fi137 | -2147483648) {
                    case 766:
                        var v138 int8 = 100
                        _ = v138
                    case 128:
                        v139 := GenS15{F0: 0}
                        _ = v139
                    case 1:
                        v1 += (v1 * 42)
                    default:
                        for fi140 := 0; fi140 < 5; fi140++ {
                            v1 -= 2
                            fmt.Println(v7, fi137, v7)
                            if (2 == 1) {
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = (("" + ("a" + ("" + ("a" + "")))) + ((((((("a" + ("a" + "a")) + "hello") + (((("foo bar" + "hello") + "foo bar") + ("\n\t" + "hello")) + ("a" + (("hello" + "\n\t") + "hello")))) + ("hello" + "\n\t")) + "foo bar") + (("hello" + (("a" + ("a" + ("\n\t" + ("" + "\n\t")))) + (("hello" + "a") + "\n\t"))) + "hello")) + ""))
                                    }
                                    v141 := GenS22{F0: (12 >= -128)}
                                    _ = v141
                                } else {
                                    v1 += v1
                                }
                                var v142 int16 = v7
                                fmt.Println(fi137)
                                switch ("a" + "a") {
                                case "\n\t":
                                    if len(v3) > 0 {
                                        v3[0] = ("hello" + (("hello" + ("hello" + (("hello" + "\n\t") + "hello"))) + "a"))
                                    }
                                }
                                var v143 uint32 = 1
                                _ = v142
                                _ = v143
                            } else {
                                v144 := GenS4{F0: 345, F1: 190}
                                v7 += (v7 << ((65535) & 63))
                                _ = v144
                            }
                            v145 := GenS8{F0: 64, F1: -128}
                            _ = fi140
                            _ = v145
                        }
                    }
                    v146 := GenS6{F0: ("" + ("foo bar" + ("a" + "a"))), F1: 0}
                    _ = fi137
                    _ = v146
                }
            }
            _ = v121
        }
    }
    fmt.Println(v1)
    for fi147 := 0; fi147 < 8; fi147++ {
        var v148 int64 = -1
        var v149 uint64 = 2
        _ = fi147
        _ = v148
        _ = v149
    }
    var v150 uint32 = 4294967295
    switch "a" {
    case "\n\t":
        if false {
            if true {
                var v151 string = ("" + "hello")
                _ = v151
            } else {
                var v152 uint16 = 340
                for fi153 := 0; fi153 < 1; fi153++ {
                    var v154 int16 = v7
                    switch fi153 {
                    case -2147483648:
                        var v155 string = (("a" + ("foo bar" + "foo bar")) + ("hello" + "\n\t"))
                        _ = v155
                    case 2147483647:
                        var v156 bool = (((("" + "") >= (("\n\t" + "") + (("\n\t" + "a") + ""))) || true) && true)
                        _ = v156
                    case 128:
                        v157 := GenS24{F0: 42, F1: (fi153 << ((770) & 63))}
                        _ = v157
                    default:
                        v154 = 1
                    }
                    for fi158 := 0; fi158 < 11; fi158++ {
                        v159 := GenS23{F0: v7, F1: -284}
                        _ = fi158
                        _ = v159
                    }
                    _ = fi153
                    _ = v154
                }
                v150 = (^(v150))
                _ = v152
            }
            var v160 float64 = v2
            _ = v160
        }
    case "":
        if true {
            _ = len(v3)
        }
    case "hello":
        var v161 uint8 = 255
        _ = v161
    default:
        for fi162 := 0; fi162 < 6; fi162++ {
            var v163 int32 = 778
            v3 = append(v3, ("a" + "foo bar"))
            _ = fi162
            _ = v163
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v150:", v150)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v150
    _ = v2
    _ = v3
    _ = v7
}
