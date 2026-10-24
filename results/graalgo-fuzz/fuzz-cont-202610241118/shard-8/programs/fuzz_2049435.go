package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint16; F2 int64 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 int8; F1 uint64; F2 string }
type GenS8 struct { F0 uint64; F1 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int; F1 string; F2 bool }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint8; F1 bool }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 string; F1 uint32; F2 int }
type GenS15 struct { F0 int16; F1 int32 }
type GenS16 struct { F0 int32; F1 uint8 }
type GenS17 struct { F0 int16; F1 int; F2 bool }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 string; F1 int8 }
type GenS20 struct { F0 int32; F1 bool; F2 uint16 }
type GenS21 struct { F0 uint; F1 string }
type GenS22 struct { F0 uint8; F1 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -2147483648
    v3 := make(map[string]string)
    v4 := []string{"hello", "\n\t", ""}
    var v5 float32 = 0.1
    v6 := make(map[int]uint32)
    v6[127] = 65535
    v7 := GenS1{F0: -86, F1: 2147483648, F2: 33}
    for fi8 := 0; fi8 < 18; fi8++ {
        if (-9223372036854775808 > 0) {
            if len(v4) > 0 {
                v4[0] = (("\n\t" + ("" + "\n\t")) + ("\n\t" + ""))
            }
            for fi9 := 0; fi9 < 13; fi9++ {
                v5 = v5
                v10 := GenS2{F0: -32768, F1: 2}
                _ = fi9
                _ = v10
            }
            v5 = 3.14159
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                fmt.Println(v2, v5, v2)
                var v12 int = fi11
                fmt.Println(v2, v2)
                _ = fi11
                _ = v12
            }
            v5 = (float32(0.0) - ((3.14159 - (3.14159 + (v5 + float32(0.0)))) + 1e10))
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 uint16 = 42
            v15 := GenS2{F0: 1000, F1: 1}
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                v16 := GenS3{F0: -61, F1: (v14 | (v14 - v14)), F2: 1}
                _ = v16
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi8
    }
    if (0.0 != 1e10) {
        var v17 uint32 = 2147483648
        _ = v17
    }
    switch "hello" {
    case "hello":
        var v18 int8 = -23
        _ = v18
    case "\n\t":
        var v19 uint64 = 99
        _ = v19
    case "":
        var v20 uint64 = 813
        _ = v20
    }
    v21 := GenS4{F0: 0, F1: -13}
    switch -2147483648 {
    case 2147483647:
        var v22 uint64 = 2
        _ = v22
    case -832:
        for fi23 := 0; fi23 < 5; fi23++ {
            v2 += v2
            if (-484 > -32768) {
                v24 := GenS3{F0: -996, F1: 65535, F2: 42}
                _ = v24
            }
            _ = fi23
        }
    case 476:
        for fi25 := 0; fi25 < 10; fi25++ {
            fmt.Println(v2, v5, fi25)
            v2 -= (v2 * v2)
            for fi26 := 0; fi26 < 16; fi26++ {
                if true {
                    v2 = (v2 - (v2 << ((1) & 63)))
                    var v27 float64 = 1e10
                    _ = v27
                } else {
                    var v28 int8 = 38
                    _ = v28
                }
                _ = fi26
            }
            _ = fi25
        }
    }
    fmt.Println(v2)
    if (32767 <= 1) {
        if (0 < 42) {
            var v29 bool = false
            _ = v29
        }
        if (127 <= -75) {
            if true {
                fmt.Println(v5, v5)
                switch 0 {
                case -120:
                    var v30 uint = 796
                    _ = v30
                }
                fmt.Println(v5)
            } else {
                fmt.Println(v2)
                switch 42 {
                case -1:
                    if true {
                        var v31 uint32 = 0
                        var v32 int32 = -672
                        _ = v31
                        _ = v32
                    } else {
                        if false {
                            switch "hello" {
                            case "\n\t":
                                switch "foo bar" {
                                case "a":
                                    var v33 uint = 0
                                    _ = v33
                                case "\n\t":
                                    fmt.Println(v5, v5)
                                default:
                                    fmt.Println(v2, v2)
                                }
                            case "":
                                for fi34 := 0; fi34 < 13; fi34++ {
                                    var v35 int32 = v2
                                    v36 := GenS5{F0: 324}
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                }
                            case "a":
                                var v37 int64 = 42
                                _ = v37
                            }
                            var v38 float64 = (-2.25 / 0.1)
                            _ = v38
                        } else {
                            if false {
                                var v39 int32 = (v2 >> ((42) & 63))
                                for fi40 := 0; fi40 < 3; fi40++ {
                                    for fi41 := 0; fi41 < 17; fi41++ {
                                        fmt.Println(fi41, v2, fi40)
                                        if true {
                                            switch fi40 {
                                            case 9223372036854775807:
                                                v42 := v3[("\n\t" + "hello")]
                                                v43 := GenS3{F0: 42, F1: 33, F2: 42}
                                                _ = v42
                                                _ = v43
                                            default:
                                                v44 := GenS6{F0: 42, F1: fi41}
                                                _ = v44
                                            }
                                            fmt.Println(v2)
                                            _ = len(v4)
                                            if (42 != 0) {
                                                switch "a" {
                                                case "a":
                                                    v4 = append(v4, "")
                                                case "\n\t":
                                                    switch fi41 {
                                                    case 2147483647:
                                                        _ = len(v1)
                                                    case -159:
                                                        switch (fi41 << ((656) & 63)) {
                                                        case 4294967295:
                                                            if len(v4) > 0 {
                                                                v4[0] = "a"
                                                            }
                                                        case -9223372036854775808:
                                                            v5 += float32((1e10 - ((-2.25 + 0.0) / (0.0 * (float64((float32(1e10) - v5)) + 1.5)))))
                                                        case 127:
                                                            switch (fi40 / (((^(fi40))) | 1)) {
                                                            case 924:
                                                                var v45 int16 = 42
                                                                _ = v45
                                                            default:
                                                                var v46 float32 = (0.0 * (-2.25 * ((v5 * 3.14159) / v5)))
                                                                _ = v46
                                                            }
                                                        default:
                                                            var v47 float64 = float64(((float32((1e10 * 0.0)) / (v5 * (v5 + 1.5))) * v5))
                                                            _ = v47
                                                        }
                                                    }
                                                default:
                                                    var v48 int16 = -573
                                                    _ = v48
                                                }
                                                v49 := GenS4{F0: (fi40 / ((fi40) | 1)), F1: 42}
                                                _ = v49
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            switch "foo bar" {
                                            case "":
                                                v50 := GenS7{F0: 42, F1: 361, F2: "foo bar"}
                                                fmt.Println(v39)
                                                v51 := GenS8{F0: 1, F1: 36}
                                                fmt.Println(v39, fi41, v39)
                                                _ = v50
                                                _ = v51
                                            case "foo bar":
                                                if false {
                                                    v39 = -1
                                                    v39 -= v2
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v52 int64 = -1
                                                        _ = v52
                                                    case "a":
                                                        if (false || true) {
                                                            v39 += -1
                                                            var v53 int16 = -1
                                                            v2 = v2
                                                            _ = v53
                                                        } else {
                                                            fmt.Println(v5, fi41)
                                                            var v54 string = ""
                                                            _ = v54
                                                        }
                                                    case "\n\t":
                                                        v5 += v5
                                                        _ = len(v3)
                                                        if ((v39 - v39) != 196) {
                                                            v55 := v3["hello"]
                                                            var v56 int16 = 1000
                                                            if false {
                                                                _ = len(v6)
                                                                if false {
                                                                    v56 += (v56 >> ((4294967295) & 63))
                                                                    if true {
                                                                        v57 := GenS1{F0: 42, F1: 995, F2: -9}
                                                                        v58 := GenS4{F0: fi41, F1: 47}
                                                                        _ = v57
                                                                        _ = v58
                                                                    }
                                                                }
                                                                v5 -= v5
                                                                var v59 int = 2147483647
                                                                _ = v59
                                                            }
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                        if len(v4) > 0 {
                                                            v4[0] = "a"
                                                        }
                                                    }
                                                } else {
                                                    switch 221 {
                                                    case 0:
                                                        v2 -= (v2 + (v2 << ((18446744073709551615) & 63)))
                                                    case 128:
                                                        if (false || (127 < 1)) {
                                                            var v60 int = (-(fi41))
                                                            var v61 int32 = 1
                                                            _ = v60
                                                            _ = v61
                                                        } else {
                                                            if (63 != 3) {
                                                                fmt.Println(v5)
                                                                if (127 == 1) {
                                                                    v62 := GenS9{F0: "\n\t"}
                                                                    if false {
                                                                        fmt.Println(v39)
                                                                    }
                                                                    _ = v62
                                                                } else {
                                                                    fmt.Println(v2, v5)
                                                                    v39 = v2
                                                                    switch (-(fi41)) {
                                                                    case -847:
                                                                        var v63 int8 = 27
                                                                        _ = v63
                                                                    case 9223372036854775807:
                                                                        var v64 float64 = -2.25
                                                                        _ = v64
                                                                    }
                                                                    v39 += (v39 - (-(v39)))
                                                                }
                                                                fmt.Println(v39, v39)
                                                            } else {
                                                                v65 := GenS3{F0: 42, F1: 42, F2: 42}
                                                                v5 -= (-2.25 - v5)
                                                                var v66 int32 = (v2 / (((v2 | v2)) | 1))
                                                                var v67 float64 = 0.0
                                                                if true {
                                                                    v68 := GenS10{F0: fi40, F1: "", F2: true}
                                                                    _ = v68
                                                                }
                                                                v39 -= v2
                                                                var v69 string = "hello"
                                                                _ = v65
                                                                _ = v66
                                                                _ = v67
                                                                _ = v69
                                                            }
                                                            v2 += (v2 ^ v2)
                                                        }
                                                    case 2147483647:
                                                        if false {
                                                            var v70 bool = false
                                                            _ = v70
                                                        } else {
                                                            fmt.Println(v2, fi40, v5)
                                                        }
                                                    }
                                                    v71 := GenS11{F0: 18446744073709551615}
                                                    _ = v71
                                                }
                                            default:
                                                if (!((fi41 * (fi40 - (fi41 * (^(fi40))))) == fi41)) {
                                                    var v72 float64 = float64(v5)
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v73 int = (^(fi41))
                                                    v73 += fi41
                                                    if true {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        if false {
                                                            v73 = fi40
                                                        }
                                                        fmt.Println(v39)
                                                    }
                                                    if false {
                                                        v72 = ((float64(v5) / (v72 + v72)) - (v72 / v72))
                                                        v5 = 1e10
                                                    }
                                                    _ = v72
                                                    _ = v73
                                                } else {
                                                    switch (-(fi41)) {
                                                    case 32767:
                                                        v39 = v39
                                                    case 4294967295:
                                                        if true {
                                                            if false {
                                                                var v74 int8 = -58
                                                                fmt.Println(fi40, v39, v39)
                                                                _ = v74
                                                            }
                                                            switch ("a" + "\n\t") {
                                                            case "foo bar":
                                                                if (!(!(!(798 < 2147483648)))) {
                                                                    fmt.Println(fi40, v5, v2)
                                                                    v39 = (v2 >> ((2) & 63))
                                                                    v75 := GenS2{F0: -929, F1: 42}
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        if (128 > 127) {
                                                                            delete(v6, (fi40 / (((fi40 >> ((566) & 63))) | 1)))
                                                                            if ((0.1 - ((float64(v5) + (0.0 + (float64(1e10) - float64((3.14159 * v5))))) + 0.0)) <= ((1.5 - 0.0) + 0.0)) {
                                                                                _ = len(v6)
                                                                            }
                                                                            if ("" != "a") {
                                                                                switch "a" {
                                                                                case "\n\t":
                                                                                    v2 -= v39
                                                                                default:
                                                                                    if (!true) {
                                                                                        fmt.Println(v2, v2)
                                                                                        v4 = append(v4, "foo bar")
                                                                                    } else {
                                                                                        var v76 string = "a"
                                                                                        switch v76 {
                                                                                        case "foo bar":
                                                                                            switch 9223372036854775807 {
                                                                                            case 730:
                                                                                                v2 = v39
                                                                                            case 2147483647:
                                                                                                switch v76 {
                                                                                                case "hello":
                                                                                                    v2 -= 0
                                                                                                    var v77 string = "hello"
                                                                                                    var v78 int32 = -2147483648
                                                                                                    _ = v77
                                                                                                    _ = v78
                                                                                                default:
                                                                                                    v5 -= v5
                                                                                                }
                                                                                            case 32768:
                                                                                                fmt.Println(v39, fi41)
                                                                                            }
                                                                                        case "a":
                                                                                            v1[42] = -927
                                                                                        }
                                                                                        if false {
                                                                                            v2 += (v2 << ((1) & 63))
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = ""
                                                                                            }
                                                                                            v79 := GenS11{F0: 4294967295}
                                                                                            _ = v79
                                                                                        } else {
                                                                                            _ = len(v4)
                                                                                            var v80 uint32 = 2
                                                                                            _ = v80
                                                                                        }
                                                                                        _ = v76
                                                                                    }
                                                                                }
                                                                                v5 += float32(0.1)
                                                                            } else {
                                                                                switch (fi40 & (fi41 * fi41)) {
                                                                                case 32767:
                                                                                    var v81 uint8 = 42
                                                                                    var v82 float32 = v5
                                                                                    _ = v81
                                                                                    _ = v82
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v5 -= v5
                                                                            _ = len(v4)
                                                                            switch (fi40 & -2147483648) {
                                                                            case 32767:
                                                                                v3["foo bar"] = "foo bar"
                                                                                fmt.Println(fi41, fi41, v5)
                                                                            }
                                                                            var v83 float64 = (float64(v5) + 0.1)
                                                                            _ = v83
                                                                        }
                                                                    case "hello":
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            if true {
                                                                                var v84 uint = 0
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                switch "" {
                                                                                case "hello":
                                                                                    v5 = ((v5 + 0.0) + v5)
                                                                                case "\n\t":
                                                                                    var v85 int16 = 32767
                                                                                    _ = v85
                                                                                }
                                                                                _ = v84
                                                                            } else {
                                                                                fmt.Println(v2, fi40, v39)
                                                                                v86 := GenS11{F0: 4294967295}
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = "a"
                                                                                }
                                                                                _ = v86
                                                                            }
                                                                            _ = len(v3)
                                                                        default:
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                switch "a" {
                                                                                case "":
                                                                                    v87 := GenS11{F0: 926}
                                                                                    _ = v87
                                                                                case "a":
                                                                                    v88 := GenS12{F0: 0, F1: (454 <= 220)}
                                                                                    _ = v88
                                                                                case "foo bar":
                                                                                    v89 := GenS13{F0: 1}
                                                                                    _ = v89
                                                                                default:
                                                                                    v90 := GenS11{F0: 18446744073709551615}
                                                                                    _ = v90
                                                                                }
                                                                            default:
                                                                                v91 := GenS14{F0: "foo bar", F1: 65535, F2: fi40}
                                                                                _ = v91
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        var v92 float64 = float64(v5)
                                                                        _ = v92
                                                                    default:
                                                                        var v93 float32 = 0.1
                                                                        _ = v93
                                                                    }
                                                                    switch "a" {
                                                                    case "":
                                                                        if len(v4) > 0 {
                                                                            v4[0] = ((("foo bar" + "") + "hello") + ("hello" + "a"))
                                                                        }
                                                                    case "hello":
                                                                        var v94 uint64 = 1
                                                                        _ = v94
                                                                    default:
                                                                        if (((("hello" + "") + "hello") + "") >= "\n\t") {
                                                                            if (1.5 == -2.25) {
                                                                                fmt.Println(fi40, v2, fi40)
                                                                                if (!(!(!((!((65535 == 381) || true)) && ((341 >= 164) || ((!true) || true)))))) {
                                                                                    delete(v3, ("\n\t" + "hello"))
                                                                                    var v95 int64 = 9223372036854775807
                                                                                    v2 -= (v39 - 42)
                                                                                    var v96 int = 32768
                                                                                    _ = v95
                                                                                    _ = v96
                                                                                }
                                                                            }
                                                                            v97 := v1[0]
                                                                            switch 32768 {
                                                                            case 2147483647:
                                                                                v39 = (v39 + v39)
                                                                            }
                                                                            v6[907] = 1
                                                                            var v98 bool = ((0 == 1) && ((!((float32(-2.25) + v5) > (v5 + 0.0))) || ((-2.25 - v5) > (float32((1.5 * 1.5)) * v5))))
                                                                            _ = v97
                                                                            _ = v98
                                                                        }
                                                                    }
                                                                    _ = v75
                                                                }
                                                            case "a":
                                                                v2 = (-(v39))
                                                            case "hello":
                                                                v1[2] = (fi41 - fi40)
                                                            }
                                                            var v99 uint = 671
                                                            _ = v99
                                                        }
                                                    }
                                                }
                                            }
                                        } else {
                                            var v100 uint = 42
                                            v100 = v100
                                            _ = v100
                                        }
                                        v101 := GenS15{F0: 0, F1: (v39 >> ((255) & 63))}
                                        _ = fi41
                                        _ = v101
                                    }
                                    if false {
                                        var v102 float32 = 1.5
                                        _ = v102
                                    } else {
                                        v103 := GenS1{F0: 100, F1: 2, F2: 73}
                                        _ = v103
                                    }
                                    _ = fi40
                                }
                                var v104 uint32 = 4294967295
                                _ = v104
                                _ = v39
                            }
                            for fi105 := 0; fi105 < 8; fi105++ {
                                if false {
                                    switch (fi105 << ((0) & 63)) {
                                    case -1:
                                        if (false || (3 > 128)) {
                                            if (!true) {
                                                if ((fi105 * fi105) < (fi105 / ((fi105) | 1))) {
                                                    v106 := GenS16{F0: v2, F1: 199}
                                                    v107 := GenS9{F0: "foo bar"}
                                                    _ = v106
                                                    _ = v107
                                                } else {
                                                    if false {
                                                        fmt.Println(v5, v5)
                                                        switch (fi105 ^ fi105) {
                                                        case 875:
                                                            var v108 int64 = -1
                                                            _ = v108
                                                        case -162:
                                                            v109 := GenS13{F0: 42}
                                                            _ = v109
                                                        case -2147483648:
                                                            v5 = ((v5 / v5) - v5)
                                                        default:
                                                            v110 := GenS4{F0: fi105, F1: 100}
                                                            _ = v110
                                                        }
                                                        for fi111 := 0; fi111 < 17; fi111++ {
                                                            if true {
                                                                v2 += 0
                                                                v2 += v2
                                                                if false {
                                                                    var v112 float64 = float64(((((((0.0 * (1e10 - v5)) + v5) + (v5 - v5)) - 3.14159) / v5) * (((3.14159 * 0.1) / v5) / 1.5)))
                                                                    var v113 uint8 = 1
                                                                    fmt.Println(v113, v2, fi105)
                                                                    _ = v112
                                                                    _ = v113
                                                                }
                                                            } else {
                                                                if ((255 != 2) || ("\n\t" == ("\n\t" + ""))) {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    switch (fi111 * fi111) {
                                                                    case 864:
                                                                        if (!(false || (1 == 903))) {
                                                                            v114 := GenS14{F0: ("\n\t" + "hello"), F1: 4294967295, F2: (fi111 >> ((238) & 63))}
                                                                            v5 = v5
                                                                            _ = v114
                                                                        }
                                                                    case 0:
                                                                        var v115 uint32 = 0
                                                                        _ = v115
                                                                    case 1:
                                                                        v2 += (v2 / (((^(v2))) | 1))
                                                                    default:
                                                                        var v116 string = "foo bar"
                                                                        _ = v116
                                                                    }
                                                                } else {
                                                                    v2 -= v2
                                                                    fmt.Println(v5, v5, v2)
                                                                }
                                                                if (false || (!true)) {
                                                                    if (42 == 1) {
                                                                        fmt.Println(fi111, v2, fi105)
                                                                        if (100 >= 127) {
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                v2 = 2147483647
                                                                            default:
                                                                                if true {
                                                                                    var v117 string = (("hello" + ((("a" + "a") + "") + "")) + "\n\t")
                                                                                    _ = v117
                                                                                }
                                                                            }
                                                                            var v118 float64 = ((0.1 - (0.0 + 0.1)) - (0.1 * 0.1))
                                                                            _ = v118
                                                                        }
                                                                        fmt.Println(fi111, fi111)
                                                                        var v119 int32 = (v2 % ((v2) | 1))
                                                                        switch "hello" {
                                                                        case "":
                                                                            v120 := GenS8{F0: 1, F1: 780}
                                                                            _ = v120
                                                                        case "hello":
                                                                            fmt.Println(v2, v5, fi111)
                                                                        }
                                                                        _ = v119
                                                                    } else {
                                                                        if true {
                                                                            var v121 uint32 = 404
                                                                            v121 += (v121 << ((65535) & 63))
                                                                            _ = v121
                                                                        }
                                                                        var v122 uint32 = 42
                                                                        var v123 uint16 = 0
                                                                        _ = v122
                                                                        _ = v123
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        if true {
                                                                            if (0 != 0) {
                                                                                if false {
                                                                                    fmt.Println(fi111)
                                                                                    var v124 int16 = -560
                                                                                    _ = v124
                                                                                } else {
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    var v125 int32 = v2
                                                                                    var v126 int = (fi111 / (((fi105 % (((^(fi105))) | 1))) | 1))
                                                                                    _ = v125
                                                                                    _ = v126
                                                                                }
                                                                                var v127 uint32 = 42
                                                                                if true {
                                                                                    var v128 int = -1
                                                                                    var v129 uint32 = v127
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        var v130 bool = (-32768 <= 1)
                                                                                        _ = v130
                                                                                    default:
                                                                                        if true {
                                                                                            var v131 bool = (!false)
                                                                                            v132 := GenS9{F0: "foo bar"}
                                                                                            if false {
                                                                                                v129 += v127
                                                                                                v133 := v1[(v127 / (((v127 * (v129 % (((^(v127))) | 1)))) | 1))]
                                                                                                switch "\n\t" {
                                                                                                case "":
                                                                                                    switch "" {
                                                                                                    case "a":
                                                                                                        v134 := GenS4{F0: fi105, F1: 102}
                                                                                                        _ = v134
                                                                                                    }
                                                                                                case "a":
                                                                                                    v135 := GenS7{F0: -78, F1: 4294967295, F2: ""}
                                                                                                    if len(v4) >= 2 {
                                                                                                        v4 = v4[1:len(v4)]
                                                                                                    }
                                                                                                    _ = v135
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v129, v131, v2)
                                                                                                default:
                                                                                                    var v136 uint64 = 328
                                                                                                    _ = v136
                                                                                                }
                                                                                                _ = v133
                                                                                            }
                                                                                            if v131 {
                                                                                                fmt.Println(v131, v127)
                                                                                                v128 = fi105
                                                                                            } else {
                                                                                                v4 = append(v4, ("foo bar" + "a"))
                                                                                            }
                                                                                            switch "a" {
                                                                                            case "hello":
                                                                                                v2 = -133
                                                                                            case "foo bar":
                                                                                                switch ("a" + "") {
                                                                                                case "\n\t":
                                                                                                    var v137 int16 = 32767
                                                                                                    _ = v137
                                                                                                case "a":
                                                                                                    var v138 int16 = -859
                                                                                                    _ = v138
                                                                                                }
                                                                                                _ = len(v1)
                                                                                            case "a":
                                                                                                fmt.Println(v127, v5)
                                                                                            }
                                                                                            _ = v131
                                                                                            _ = v132
                                                                                        } else {
                                                                                            v2 = v2
                                                                                            _ = len(v4)
                                                                                        }
                                                                                        var v139 uint = 92
                                                                                        _ = v139
                                                                                    }
                                                                                    var v140 int8 = 1
                                                                                    var v141 float64 = (0.0 + ((float64(v5) + 1e10) * -2.25))
                                                                                    _ = v128
                                                                                    _ = v129
                                                                                    _ = v140
                                                                                    _ = v141
                                                                                } else {
                                                                                    var v142 int16 = 1
                                                                                    if false {
                                                                                        v2 += (v2 << ((962) & 63))
                                                                                        _ = len(v3)
                                                                                    }
                                                                                    var v143 uint8 = 59
                                                                                    _ = v142
                                                                                    _ = v143
                                                                                }
                                                                                v144 := GenS14{F0: "foo bar", F1: 2147483647, F2: (fi105 << ((4294967295) & 63))}
                                                                                _ = v127
                                                                                _ = v144
                                                                            } else {
                                                                                var v145 float64 = 1.5
                                                                                v145 -= (v145 * -2.25)
                                                                                _ = v145
                                                                            }
                                                                        } else {
                                                                            var v146 uint32 = 2147483647
                                                                            var v147 uint64 = 42
                                                                            if false {
                                                                                if (((1.5 / 1.5) >= (3.14159 * -2.25)) && ((((0.0 / 1.0) - 1e10) + ((3.14159 / 1.5) - ((0.0 - 3.14159) + (1e10 * 0.0)))) == 1e10)) {
                                                                                    fmt.Println(v2, v5)
                                                                                    v148 := GenS7{F0: -1, F1: 445, F2: ""}
                                                                                    _ = v148
                                                                                }
                                                                                if ((-691 > 42) && false) {
                                                                                    var v149 int32 = -1
                                                                                    if false {
                                                                                        var v150 int = fi111
                                                                                        var v151 uint32 = v146
                                                                                        _ = v150
                                                                                        _ = v151
                                                                                    }
                                                                                    _ = v149
                                                                                }
                                                                            } else {
                                                                                v152 := GenS17{F0: -1, F1: (fi111 ^ 4294967295), F2: ((0 <= (v147 ^ 2)) || true)}
                                                                                _ = v152
                                                                            }
                                                                            if true {
                                                                                if false {
                                                                                    v4 = append(v4, "hello")
                                                                                }
                                                                                v153 := GenS11{F0: 382}
                                                                                v4 = append(v4, "a")
                                                                                fmt.Println(v5, fi111)
                                                                                var v154 int32 = v2
                                                                                _ = v153
                                                                                _ = v154
                                                                            } else {
                                                                                switch (fi105 >> ((65535) & 63)) {
                                                                                case 962:
                                                                                    if true {
                                                                                        v155 := GenS18{F0: 9223372036854775807}
                                                                                        _ = v155
                                                                                    } else {
                                                                                        if true {
                                                                                            if (v5 > ((1e10 / (0.1 / v5)) - v5)) {
                                                                                                v5 -= (1e10 + 0.0)
                                                                                                var v156 uint64 = 1
                                                                                                _ = v156
                                                                                            } else {
                                                                                                v2 += v2
                                                                                            }
                                                                                            if (float32(0.0) != 3.14159) {
                                                                                                var v157 float64 = float64(float32((1e10 * ((-2.25 + 0.0) - ((-2.25 / 1e10) / 1.5)))))
                                                                                                fmt.Println(v147, fi111, v147)
                                                                                                _ = v157
                                                                                            }
                                                                                            var v158 int = fi105
                                                                                            var v159 float32 = 0.1
                                                                                            _ = v158
                                                                                            _ = v159
                                                                                        }
                                                                                        switch (fi111 - (fi111 + (fi105 | (fi111 % ((42) | 1))))) {
                                                                                        case 42:
                                                                                            v147 += v147
                                                                                        }
                                                                                        v160 := GenS6{F0: (v146 - 2147483647), F1: fi105}
                                                                                        _ = v160
                                                                                    }
                                                                                case 418:
                                                                                    fmt.Println(v146)
                                                                                }
                                                                                fmt.Println(v146, v147)
                                                                            }
                                                                            _ = v146
                                                                            _ = v147
                                                                        }
                                                                    case "":
                                                                        v5 -= 0.1
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        if true {
                                                                            fmt.Println(v2, v2)
                                                                            v161 := GenS5{F0: 42}
                                                                            fmt.Println(fi105, fi111)
                                                                            v5 += 1.5
                                                                            var v162 int8 = 1
                                                                            _ = v161
                                                                            _ = v162
                                                                        }
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    }
                                                                    v5 = (v5 - ((3.14159 / 3.14159) / v5))
                                                                    v2 -= (v2 | 42)
                                                                }
                                                                switch ("hello" + "hello") {
                                                                case "foo bar":
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        var v163 int16 = -32768
                                                                        _ = v163
                                                                    }
                                                                case "a":
                                                                    if false {
                                                                        v5 += float32(3.14159)
                                                                        v5 = 1.5
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            var v164 uint = 42
                                                                            _ = v164
                                                                        case "foo bar":
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                v165 := GenS9{F0: "foo bar"}
                                                                                _ = v165
                                                                            default:
                                                                                _ = len(v4)
                                                                            }
                                                                        default:
                                                                            v5 = (v5 * v5)
                                                                        }
                                                                    } else {
                                                                        switch 42 {
                                                                        case 2147483647:
                                                                            if (false && (true || (fi105 == (fi105 | fi111)))) {
                                                                                v2 = v2
                                                                            } else {
                                                                                fmt.Println(v5)
                                                                            }
                                                                        case 0:
                                                                            fmt.Println(fi105)
                                                                        case 32767:
                                                                            v5 -= -2.25
                                                                        default:
                                                                            var v166 uint16 = 42
                                                                            _ = v166
                                                                        }
                                                                        v5 -= float32(1e10)
                                                                        v5 = (float32(0.0) * 1.5)
                                                                    }
                                                                case "hello":
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                default:
                                                                    delete(v1, 904)
                                                                }
                                                            }
                                                            _ = fi111
                                                        }
                                                    } else {
                                                        v167 := v3[(("\n\t" + ("foo bar" + "hello")) + "\n\t")]
                                                        v167 = ((((v167 + (v167 + (("" + v167) + (("" + (("foo bar" + v167) + v167)) + v167)))) + "") + "foo bar") + "a")
                                                        _ = v167
                                                    }
                                                    v168 := GenS19{F0: "\n\t", F1: 100}
                                                    for fi169 := 0; fi169 < 17; fi169++ {
                                                        if len(v4) > 0 {
                                                            v4[0] = "foo bar"
                                                        }
                                                        _ = fi169
                                                    }
                                                    _ = v168
                                                }
                                            }
                                        } else {
                                            var v170 string = "a"
                                            _ = v170
                                        }
                                    default:
                                        fmt.Println(v5, fi105, v5)
                                    }
                                }
                                _ = fi105
                            }
                            _ = len(v4)
                        }
                    }
                case 127:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 32767:
                    if true {
                        v6[127] = 932
                        v171 := GenS4{F0: 0, F1: 100}
                        _ = v171
                    } else {
                        if false {
                            for fi172 := 0; fi172 < 6; fi172++ {
                                var v173 uint64 = 4294967295
                                for fi174 := 0; fi174 < 20; fi174++ {
                                    fmt.Println(v2, fi174)
                                    v173 = (v173 >> ((2) & 63))
                                    _ = fi174
                                }
                                var v175 uint8 = 42
                                _ = fi172
                                _ = v173
                                _ = v175
                            }
                            fmt.Println(v5)
                        } else {
                            switch 4294967295 {
                            case -2147483648:
                                if (42 <= 0) {
                                    for fi176 := 0; fi176 < 16; fi176++ {
                                        v177 := GenS6{F0: 65535, F1: (fi176 & fi176)}
                                        v2 = (v2 & v2)
                                        if true {
                                            v178 := GenS20{F0: (v2 >> ((42) & 63)), F1: (true || (88 <= 2)), F2: 42}
                                            _ = v178
                                        } else {
                                            fmt.Println(fi176, v2)
                                            v5 -= v5
                                            if false {
                                                if (1 <= 42) {
                                                    if (2147483647 != 157) {
                                                        var v179 string = (("foo bar" + ("\n\t" + "a")) + ("a" + ("\n\t" + "hello")))
                                                        _ = v179
                                                    }
                                                    v2 -= v2
                                                } else {
                                                    if false {
                                                        var v180 uint = 18446744073709551615
                                                        var v181 uint64 = 4294967295
                                                        v181 = (v181 + (^(v181)))
                                                        _ = v180
                                                        _ = v181
                                                    }
                                                }
                                                v2 = (v2 >> ((255) & 63))
                                            }
                                        }
                                        _ = fi176
                                        _ = v177
                                    }
                                } else {
                                    v2 += v2
                                    for fi182 := 0; fi182 < 17; fi182++ {
                                        v2 = v2
                                        for fi183 := 0; fi183 < 7; fi183++ {
                                            if (42 == 0) {
                                                v184 := GenS21{F0: 1, F1: "foo bar"}
                                                v2 += -2147483648
                                                _ = v184
                                            }
                                            _ = fi183
                                        }
                                        var v185 uint64 = 2
                                        _ = fi182
                                        _ = v185
                                    }
                                    v5 = (v5 / ((v5 + 0.1) / v5))
                                }
                            case 128:
                                switch "foo bar" {
                                case "foo bar":
                                    var v186 int16 = -1
                                    _ = v186
                                case "a":
                                    v4 = append(v4, "a")
                                case "\n\t":
                                    v187 := GenS17{F0: -736, F1: -1, F2: true}
                                    _ = v187
                                }
                            case 0:
                                if (true || (156 <= 45)) {
                                    var v188 uint64 = 1
                                    v2 += 0
                                    switch 4294967295 {
                                    case 419:
                                        switch -133 {
                                        case -9223372036854775808:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        case 42:
                                            v6[2147483647] = 2147483647
                                        case 0:
                                            var v189 float32 = ((v5 - v5) / v5)
                                            _ = v189
                                        default:
                                            fmt.Println(v5, v5, v188)
                                        }
                                    default:
                                        v2 = (v2 % ((v2) | 1))
                                    }
                                    _ = v188
                                } else {
                                    v5 += ((float32((-2.25 + 0.0)) + v5) - v5)
                                }
                            }
                            if false {
                                v2 += (v2 & (v2 - 339))
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v190 := GenS11{F0: 2}
                                for fi191 := 0; fi191 < 3; fi191++ {
                                    for fi192 := 0; fi192 < 15; fi192++ {
                                        switch (fi192 ^ (fi192 ^ fi192)) {
                                        case -2147483648:
                                            if true {
                                                if len(v4) > 0 {
                                                    v4[0] = "foo bar"
                                                }
                                                var v193 int8 = -39
                                                var v194 int32 = (v2 >> ((942) & 63))
                                                _ = v193
                                                _ = v194
                                            }
                                        default:
                                            v4 = append(v4, "")
                                        }
                                        switch "\n\t" {
                                        case "\n\t":
                                            v195 := GenS22{F0: 48, F1: ("\n\t" + "a")}
                                            _ = v195
                                        case "":
                                            v196 := GenS3{F0: -953, F1: 42, F2: -378}
                                            _ = v196
                                        default:
                                            if true {
                                                if true {
                                                    if (0 > 403) {
                                                        v2 += 2147483647
                                                        v197 := v6[0]
                                                        _ = v197
                                                    } else {
                                                        var v198 bool = (-1 != 1)
                                                        switch fi192 {
                                                        case -2147483648:
                                                            fmt.Println(v198)
                                                        case 128:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        case 9223372036854775807:
                                                            v199 := GenS13{F0: 18446744073709551615}
                                                            _ = v199
                                                        default:
                                                            if (v198 && (!(-145 >= 0))) {
                                                                var v200 uint = 255
                                                                fmt.Println(fi191, v5, fi192)
                                                                var v201 float32 = 0.1
                                                                _ = v200
                                                                _ = v201
                                                            } else {
                                                                v3[("a" + ("" + "hello"))] = ""
                                                                switch (^(fi191)) {
                                                                case -321:
                                                                    var v202 int32 = -661
                                                                    _ = v202
                                                                case -24:
                                                                    var v203 uint16 = 42
                                                                    _ = v203
                                                                case -261:
                                                                    var v204 int16 = 1
                                                                    _ = v204
                                                                }
                                                            }
                                                        }
                                                        v205 := GenS6{F0: 988, F1: fi191}
                                                        _ = v198
                                                        _ = v205
                                                    }
                                                }
                                            } else {
                                                var v206 string = ("hello" + "\n\t")
                                                _ = v206
                                            }
                                        }
                                        switch fi192 {
                                        case 2147483647:
                                            var v207 int = (fi192 >> ((2) & 63))
                                            _ = v207
                                        case 1:
                                            v2 = -2147483648
                                            var v208 bool = false
                                            _ = v208
                                        case 32767:
                                            fmt.Println(fi191, fi191, fi191)
                                            v209 := GenS16{F0: (v2 & (v2 + v2)), F1: 42}
                                            v210 := GenS18{F0: 456}
                                            _ = v209
                                            _ = v210
                                        }
                                        var v211 uint64 = 1
                                        var v212 uint64 = (v211 - 42)
                                        _ = fi192
                                        _ = v211
                                        _ = v212
                                    }
                                    v2 = v2
                                    _ = fi191
                                }
                                _ = v190
                            }
                            v5 = float32(1.5)
                        }
                    }
                default:
                    var v213 float32 = -2.25
                    _ = v213
                }
                v5 = (v5 - (v5 + (((v5 * (v5 / ((v5 - (float32(((((0.1 / float64(float32(((0.0 + (float64((v5 + (-2.25 * 0.0))) + -2.25)) + 1.5)))) + 0.1) + 1e10) * (0.1 - float64(0.1)))) + 1e10)) * v5))) + (float32(3.14159) - v5)) / v5)))
            }
            var v214 int32 = (v2 << ((152) & 63))
            _ = v214
        } else {
            if false {
                if true {
                    v2 += 491
                    delete(v6, 127)
                } else {
                    var v215 int8 = -48
                    _ = v215
                }
                var v216 uint = 18446744073709551615
                v217 := GenS14{F0: ("hello" + "a"), F1: 2, F2: 9223372036854775807}
                _ = v216
                _ = v217
            } else {
                if false {
                    if false {
                        switch 4294967295 {
                        case 127:
                            _ = len(v4)
                        default:
                            if (float32(1.5) != 3.14159) {
                                var v218 string = "\n\t"
                                for fi219 := 0; fi219 < 15; fi219++ {
                                    v2 -= (^(v2))
                                    _ = fi219
                                }
                                _ = v218
                            } else {
                                var v220 int32 = (v2 - v2)
                                v220 = 1
                                _ = v220
                            }
                        }
                        for fi221 := 0; fi221 < 11; fi221++ {
                            var v222 uint64 = 42
                            _ = fi221
                            _ = v222
                        }
                    }
                }
            }
        }
    }
    v223 := GenS13{F0: 973}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v223
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
