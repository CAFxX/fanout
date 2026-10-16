package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS3 struct { F0 int16; F1 int64; F2 uint16 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 int32; F1 uint16; F2 int32 }
type GenS7 struct { F0 int16; F1 int; F2 uint16 }
type GenS8 struct { F0 bool; F1 uint16; F2 int64 }
type GenS9 struct { F0 int8; F1 uint16 }
type GenS10 struct { F0 int; F1 bool }
type GenS11 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int8 }
func genFn0(p0 uint16, p1 uint, p2 string) string {
    var v5 int8 = 42
    var v6 int8 = (^(v5))
    var v7 int = 128
    var v8 uint8 = 0
    v9 := []int{(-(v7)), (v7 % (((v7 * 884)) | 1)), 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    v1 := []uint32{}
    var v2 uint64 = 4294967295
    v3 := []string{"foo bar", ""}
    v4 := make(map[string]bool)
    v10 := v4["\n\t"]
    if false {
        if (!(("\n\t" <= "hello") || (2147483647 >= 42))) {
            for fi11 := 0; fi11 < 4; fi11++ {
                delete(v4, "hello")
                v2 -= 2
                _ = fi11
            }
            v12 := GenS1{F0: 2, F1: -122}
            _ = v12
        }
        if false {
            v13 := GenS2{F0: 127, F1: 749, F2: 42}
            _ = v13
        }
    }
    if len(v3) > 0 {
        v3[0] = (("a" + "hello") + "hello")
    }
    fmt.Println(v2, v2)
    if (-718 >= int(v2)) {
        fmt.Println(v2, v10)
        fmt.Println(v2)
        if (958 == 817) {
            for fi14 := 0; fi14 < 20; fi14++ {
                switch "" {
                case "":
                    for fi15 := 0; fi15 < 18; fi15++ {
                        v2 = 2
                        fmt.Println(fi14, v2, fi15)
                        fmt.Println(v2, v2)
                        _ = fi15
                    }
                case "\n\t":
                    fmt.Println(v10)
                default:
                    switch (("a" + "hello") + "foo bar") {
                    case "\n\t":
                        var v16 uint8 = 128
                        _ = v16
                    default:
                        switch fi14 {
                        case -9223372036854775808:
                            if true {
                                switch (fi14 + fi14) {
                                case 127:
                                    for fi17 := 0; fi17 < 2; fi17++ {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        switch fi17 {
                                        case -235:
                                            if v10 {
                                                var v18 bool = (1.5 != ((0.0 - 0.0) + (3.14159 / 1.5)))
                                                _ = v18
                                            } else {
                                                v19 := GenS3{F0: 32767, F1: 42, F2: 0}
                                                v20 := GenS3{F0: 32767, F1: 828, F2: 0}
                                                if v10 {
                                                    if ((!false) || (!v10)) {
                                                        v1 = append(v1, 65535)
                                                        switch (fi17 << ((42) & 63)) {
                                                        case 32768:
                                                            v21 := GenS2{F0: 2, F1: 573, F2: -1}
                                                            var v22 uint8 = 128
                                                            _ = v21
                                                            _ = v22
                                                        default:
                                                            v10 = true
                                                        }
                                                        switch fi14 {
                                                        case -9223372036854775808:
                                                            if ("foo bar" != (("\n\t" + "") + ("" + ("" + (("a" + "") + "hello"))))) {
                                                                var v23 uint32 = 65535
                                                                _ = v23
                                                            } else {
                                                                var v24 string = ""
                                                                _ = v24
                                                            }
                                                        case -2147483648:
                                                            v2 -= v2
                                                        case -847:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        default:
                                                            v25 := GenS4{F0: v10, F1: 116}
                                                            v2 = 325
                                                            if false {
                                                                v2 -= v2
                                                                if (255 < 46) {
                                                                    switch -9223372036854775808 {
                                                                    case -89:
                                                                        _ = len(v3)
                                                                    case 1:
                                                                        fmt.Println(fi17, v10)
                                                                    case -290:
                                                                        v2 += v2
                                                                    default:
                                                                        var v26 int64 = -767
                                                                        _ = v26
                                                                    }
                                                                    _ = len(v1)
                                                                    v27 := GenS5{F0: (fi14 / ((fi14) | 1)), F1: -9223372036854775808}
                                                                    _ = v27
                                                                } else {
                                                                    v10 = (!true)
                                                                }
                                                                var v28 float64 = (3.14159 * 3.14159)
                                                                _ = v28
                                                            }
                                                            _ = v25
                                                        }
                                                    } else {
                                                        var v29 uint16 = 284
                                                        var v30 uint = 947
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                    v31 := GenS6{F0: -847, F1: 65535, F2: -2147483648}
                                                    _ = v31
                                                } else {
                                                    var v32 uint16 = 4
                                                    switch "hello" {
                                                    case "a":
                                                        if v10 {
                                                            v2 = v2
                                                            v33 := GenS4{F0: (int8(v32) > 125), F1: 496}
                                                            v10 = v10
                                                            v10 = v10
                                                            _ = v33
                                                        }
                                                    }
                                                    _ = v32
                                                }
                                                v34 := GenS6{F0: 2147483647, F1: 2, F2: 1}
                                                var v35 int64 = -903
                                                if len(v1) > 0 {
                                                    v1[0] = 2147483647
                                                }
                                                _ = v19
                                                _ = v20
                                                _ = v34
                                                _ = v35
                                            }
                                        }
                                        _ = fi17
                                    }
                                default:
                                    if (-881 != (-(fi14))) {
                                        var v36 int = fi14
                                        _ = v36
                                    } else {
                                        if v10 {
                                            if (float32(1.5) == ((0.1 - -2.25) * float32(-2.25))) {
                                                fmt.Println(v2, v10, v10)
                                                v37 := GenS1{F0: 127, F1: 0}
                                                v38 := GenS1{F0: 0, F1: -128}
                                                _ = v37
                                                _ = v38
                                            } else {
                                                if (((18446744073709551615 <= 4294967295) && false) || true) {
                                                    v10 = v10
                                                }
                                            }
                                            var v39 int8 = 110
                                            v40 := GenS5{F0: (fi14 << ((0) & 63)), F1: -9223372036854775808}
                                            _ = v39
                                            _ = v40
                                        }
                                    }
                                }
                                switch "hello" {
                                case "":
                                    for fi41 := 0; fi41 < 16; fi41++ {
                                        var v42 bool = (999 < 255)
                                        v43 := GenS1{F0: 34, F1: 0}
                                        _ = fi41
                                        _ = v42
                                        _ = v43
                                    }
                                case "a":
                                    if v10 {
                                        switch (("\n\t" + ("" + (("\n\t" + ("" + "foo bar")) + "\n\t"))) + "hello") {
                                        case "a":
                                            var v44 uint32 = 708
                                            _ = v44
                                        case "foo bar":
                                            if (1 != -28) {
                                                for fi45 := 0; fi45 < 16; fi45++ {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        if (!(42 < 255)) {
                                                            var v46 uint8 = 127
                                                            switch -9223372036854775808 {
                                                            case -1:
                                                                v47 := GenS7{F0: 1, F1: fi14, F2: 65535}
                                                                _ = v47
                                                            case 32767:
                                                                if v10 {
                                                                    switch fi14 {
                                                                    case 178:
                                                                        v46 += v46
                                                                    case 359:
                                                                        if v10 {
                                                                            v46 += (v46 / (((v46 << ((15) & 63))) | 1))
                                                                        }
                                                                    }
                                                                }
                                                            case 42:
                                                                if v10 {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = ("" + "hello")
                                                                    }
                                                                    fmt.Println(v46, fi45, v46)
                                                                } else {
                                                                    var v48 string = "a"
                                                                    var v49 float64 = 0.1
                                                                    var v50 float64 = (float64(((float32(1e10) * ((0.0 / 1.5) - 3.14159)) - float32(v49))) + float64(1e10))
                                                                    _ = v48
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                            default:
                                                                var v51 int32 = 1
                                                                _ = v51
                                                            }
                                                            _ = v46
                                                        }
                                                    case "\n\t":
                                                        v2 += (v2 - v2)
                                                    default:
                                                        var v52 int = (fi14 * fi45)
                                                        _ = v52
                                                    }
                                                    if (42 >= 9223372036854775807) {
                                                        var v53 int = (fi45 * fi14)
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v2 = (^(v2))
                                                            v54 := GenS1{F0: 86, F1: 1}
                                                            _ = v54
                                                        case "foo bar":
                                                            switch (v53 % ((703) | 1)) {
                                                            case 4294967295:
                                                                if v10 {
                                                                    v4[("foo bar" + "")] = (-542 <= -32768)
                                                                    fmt.Println(fi45)
                                                                    v53 -= fi45
                                                                } else {
                                                                    switch (v53 << ((4294967295) & 63)) {
                                                                    case 222:
                                                                        var v55 uint32 = 2147483647
                                                                        _ = v55
                                                                    case -57:
                                                                        if v10 {
                                                                            var v56 float32 = 1e10
                                                                            _ = len(v1)
                                                                            var v57 uint8 = 62
                                                                            fmt.Println(v2, v56, v56)
                                                                            _ = v56
                                                                            _ = v57
                                                                        } else {
                                                                            _ = len(v3)
                                                                            var v58 int8 = 56
                                                                            _ = len(v1)
                                                                            _ = v58
                                                                        }
                                                                    case 32768:
                                                                        var v59 uint = 980
                                                                        _ = v59
                                                                    default:
                                                                        var v60 float64 = float64(1e10)
                                                                        _ = v60
                                                                    }
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v10 = v10
                                                                    }
                                                                    var v61 bool = true
                                                                    var v62 int = fi45
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                            case 32768:
                                                                v2 = (v2 + (v2 * (v2 & v2)))
                                                            case 840:
                                                                v63 := GenS8{F0: v10, F1: 0, F2: 9223372036854775807}
                                                                _ = v63
                                                            }
                                                        case "":
                                                            var v64 uint32 = 2147483648
                                                            _ = v64
                                                        default:
                                                            _ = len(v1)
                                                        }
                                                        v4["a"] = v10
                                                        switch (-(v53)) {
                                                        case -2147483648:
                                                            fmt.Println(fi45, fi14, v2)
                                                        case 1:
                                                            switch "" {
                                                            case "a":
                                                                if false {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        delete(v4, (("a" + "a") + ("\n\t" + "")))
                                                                    case "foo bar":
                                                                        v65 := GenS9{F0: -1, F1: 452}
                                                                        _ = v65
                                                                    case "":
                                                                        if ((float32(3.14159) + -2.25) == ((0.0 - (1e10 + -2.25)) - (float32(1e10) + ((0.0 + 3.14159) * ((1e10 + 0.1) - (1e10 - (-2.25 / float32(((-2.25 + 1e10) * 0.1))))))))) {
                                                                            fmt.Println(fi14, v10, fi14)
                                                                            fmt.Println(v53, v2)
                                                                            v1 = append(v1, 0)
                                                                        }
                                                                    default:
                                                                        var v66 uint16 = 32767
                                                                        _ = v66
                                                                    }
                                                                }
                                                            case "hello":
                                                                var v67 uint32 = 385
                                                                var v68 float64 = float64((1e10 + float32((-2.25 * 1e10))))
                                                                switch (v53 - 0) {
                                                                case 1:
                                                                    var v69 uint64 = (v2 >> ((65535) & 63))
                                                                    _ = v69
                                                                default:
                                                                    v2 += 42
                                                                }
                                                                _ = v67
                                                                _ = v68
                                                            case "":
                                                                if len(v1) > 0 {
                                                                    v1[0] = 2
                                                                }
                                                            default:
                                                                if (-156 >= 1) {
                                                                    switch fi45 {
                                                                    case -774:
                                                                        if (-124 > -128) {
                                                                            v10 = (255 > 393)
                                                                            v53 = (v53 % ((fi14) | 1))
                                                                        } else {
                                                                            switch (-(fi45)) {
                                                                            case 128:
                                                                                v10 = v10
                                                                            case -478:
                                                                                v70 := GenS6{F0: 42, F1: 922, F2: -2147483648}
                                                                                _ = v70
                                                                            case 2147483647:
                                                                                var v71 int8 = 42
                                                                                _ = v71
                                                                            }
                                                                            switch fi14 {
                                                                            case -367:
                                                                                _ = len(v4)
                                                                            case 4294967295:
                                                                                var v72 uint64 = (v2 % ((1) | 1))
                                                                                _ = v72
                                                                            }
                                                                            v73 := GenS5{F0: (^(v53)), F1: 0}
                                                                            _ = v73
                                                                        }
                                                                    case -2147483648:
                                                                        var v74 int16 = 1000
                                                                        _ = v74
                                                                    }
                                                                    var v75 uint = 2
                                                                    v10 = v10
                                                                    v76 := GenS10{F0: v53, F1: ((1e10 - 0.1) >= (float64(((0.1 * 0.1) * 0.0)) - ((1e10 - (3.14159 - -2.25)) * 0.0)))}
                                                                    _ = v75
                                                                    _ = v76
                                                                } else {
                                                                    var v77 uint8 = 128
                                                                    switch int(v2) {
                                                                    case 42:
                                                                        _ = len(v1)
                                                                    case 9223372036854775807:
                                                                        v78 := GenS11{F0: 32767, F1: (v77 ^ v77), F2: v77}
                                                                        var v79 int16 = 1000
                                                                        var v80 uint32 = 0
                                                                        _ = v78
                                                                        _ = v79
                                                                        _ = v80
                                                                    case 128:
                                                                        v3 = append(v3, "foo bar")
                                                                    }
                                                                    _ = v77
                                                                }
                                                                var v81 uint32 = 2
                                                                var v82 uint = 963
                                                                _ = v81
                                                                _ = v82
                                                            }
                                                        default:
                                                            v1 = append(v1, 2147483648)
                                                        }
                                                        v83 := v4["\n\t"]
                                                        _ = v53
                                                        _ = v83
                                                    }
                                                    v84 := GenS11{F0: -32768, F1: 255, F2: 32}
                                                    if v10 {
                                                        var v85 int64 = -9223372036854775808
                                                        _ = v85
                                                    }
                                                    _ = fi45
                                                    _ = v84
                                                }
                                                for fi86 := 0; fi86 < 12; fi86++ {
                                                    v87 := GenS7{F0: 42, F1: (^(fi14)), F2: 0}
                                                    _ = fi86
                                                    _ = v87
                                                }
                                            }
                                        case "":
                                            v4[(("foo bar" + "foo bar") + "a")] = (651 < uint16(v2))
                                        }
                                        for fi88 := 0; fi88 < 7; fi88++ {
                                            v4["\n\t"] = v10
                                            v89 := GenS4{F0: v10, F1: -1}
                                            _ = fi88
                                            _ = v89
                                        }
                                    } else {
                                        v10 = (9223372036854775807 <= 9223372036854775807)
                                        v2 -= (v2 / ((v2) | 1))
                                    }
                                case "\n\t":
                                    v90 := GenS8{F0: ((v2 + (^(v2))) != 874), F1: 255, F2: 1}
                                    _ = v90
                                }
                                var v91 uint = 42
                                _ = v91
                            }
                        default:
                            for fi92 := 0; fi92 < 10; fi92++ {
                                var v93 bool = (-2.25 == 0.1)
                                if (!v10) {
                                    switch (fi14 | 42) {
                                    case 42:
                                        v93 = false
                                    case -9223372036854775808:
                                        var v94 uint64 = (v2 - 42)
                                        _ = v94
                                    case -2147483648:
                                        var v95 float64 = 1e10
                                        _ = v95
                                    default:
                                        if true {
                                            switch fi14 {
                                            case 320:
                                                var v96 int32 = -2147483648
                                                _ = v96
                                            case 128:
                                                if (false && v10) {
                                                    var v97 int8 = 127
                                                    v10 = (("\n\t" + ("foo bar" + ("a" + ""))) >= (("hello" + "hello") + "hello"))
                                                    var v98 uint8 = 113
                                                    if v10 {
                                                        var v99 uint64 = 943
                                                        _ = v99
                                                    } else {
                                                        var v100 int16 = 32767
                                                        switch "hello" {
                                                        case "hello":
                                                            if false {
                                                                var v101 float32 = ((0.0 - (1e10 / (1e10 - (float32(float64((0.1 - 1e10))) + (1e10 / float32((((3.14159 - 1.5) - (0.0 + float64((1.5 * ((3.14159 + 0.1) - 1.5))))) / (float64(float32(1e10)) + -2.25)))))))) - 3.14159)
                                                                v102 := v4[(("hello" + "foo bar") + ("foo bar" + (("hello" + "") + ("hello" + ("foo bar" + "a")))))]
                                                                _ = v101
                                                                _ = v102
                                                            } else {
                                                                _ = len(v3)
                                                                if (893 <= 325) {
                                                                    if false {
                                                                        if (-317 != 0) {
                                                                            if (586 != 795) {
                                                                                switch (fi14 & -9223372036854775808) {
                                                                                case 1:
                                                                                    if ((!v93) && ((^(v2)) < v2)) {
                                                                                        v103 := GenS10{F0: 330, F1: false}
                                                                                        var v104 int8 = v97
                                                                                        _ = v103
                                                                                        _ = v104
                                                                                    }
                                                                                case -491:
                                                                                    if v93 {
                                                                                        switch "a" {
                                                                                        case "":
                                                                                            var v105 float64 = float64((((3.14159 * (1e10 / float32(-2.25))) - 0.1) - ((float32(1e10) - 1e10) - -2.25)))
                                                                                            _ = v105
                                                                                        case "hello":
                                                                                            v2 += v2
                                                                                        }
                                                                                        v93 = v10
                                                                                    } else {
                                                                                        if v93 {
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            var v106 int64 = 194
                                                                                            _ = v106
                                                                                        } else {
                                                                                            if true {
                                                                                                v107 := GenS11{F0: (v100 + v100), F1: 128, F2: (v98 + 175)}
                                                                                                _ = v107
                                                                                            } else {
                                                                                                switch fi92 {
                                                                                                case 224:
                                                                                                    var v108 string = "hello"
                                                                                                    _ = v108
                                                                                                }
                                                                                                v2 -= 2
                                                                                                fmt.Println(fi92, v100)
                                                                                                v100 += (v100 - int16(v97))
                                                                                            }
                                                                                            if (fi92 < (fi14 << ((65535) & 63))) {
                                                                                                fmt.Println(v97, v10)
                                                                                                v97 += v97
                                                                                                var v109 int32 = -1
                                                                                                _ = v109
                                                                                            }
                                                                                            v2 -= (v2 | v2)
                                                                                        }
                                                                                        var v110 int32 = 42
                                                                                        _ = v110
                                                                                    }
                                                                                    v2 = (v2 - v2)
                                                                                }
                                                                                delete(v4, (("" + ((("hello" + "hello") + (((("a" + "hello") + "hello") + "foo bar") + "foo bar")) + ("" + "hello"))) + "\n\t"))
                                                                                v100 += (v100 - v100)
                                                                            } else {
                                                                                delete(v4, "a")
                                                                            }
                                                                        } else {
                                                                            v93 = (!(fi14 == fi92))
                                                                            var v111 int32 = 2147483647
                                                                            _ = v111
                                                                        }
                                                                        var v112 bool = (uint32(v98) < 0)
                                                                        v97 = (v97 + int8(v98))
                                                                        _ = v112
                                                                    }
                                                                    switch fi14 {
                                                                    case 0:
                                                                        fmt.Println(fi14)
                                                                    case 32767:
                                                                        v97 -= v97
                                                                    case 1:
                                                                        switch (-(fi14)) {
                                                                        case 540:
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                if (0.1 != 1.5) {
                                                                                    if (("" + ("" + "\n\t")) < ("" + "hello")) {
                                                                                        fmt.Println(fi14)
                                                                                    }
                                                                                    v113 := GenS7{F0: (v100 ^ (v100 % ((int16(v98)) | 1))), F1: fi14, F2: 65535}
                                                                                    var v114 uint32 = 33
                                                                                    fmt.Println(v2)
                                                                                    var v115 string = ""
                                                                                    _ = v113
                                                                                    _ = v114
                                                                                    _ = v115
                                                                                }
                                                                            case "":
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    var v116 float64 = (1.5 / (3.14159 + 1e10))
                                                                                    _ = v116
                                                                                case "":
                                                                                    v117 := GenS6{F0: -1, F1: 301, F2: 42}
                                                                                    _ = v117
                                                                                case "foo bar":
                                                                                    v98 += (^(v98))
                                                                                default:
                                                                                    var v118 uint32 = 2147483647
                                                                                    _ = v118
                                                                                }
                                                                            case "a":
                                                                                var v119 uint8 = (v98 ^ (v98 ^ v98))
                                                                                _ = v119
                                                                            }
                                                                        case -477:
                                                                            if (((3.14159 / (1.5 / 1.0)) / 3.14159) >= 0.1) {
                                                                                var v120 uint64 = 42
                                                                                _ = len(v1)
                                                                                v10 = (42 == 65535)
                                                                                _ = v120
                                                                            } else {
                                                                                if (v2 < (v2 - v2)) {
                                                                                    fmt.Println(v2, v93)
                                                                                    _ = len(v4)
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        var v121 float32 = 0.1
                                                                        _ = v121
                                                                    }
                                                                    v4[(("hello" + "") + "a")] = ((-2.25 >= float32((3.14159 + (1.5 + (0.1 / (1e10 - float64((1e10 + 1.5)))))))) && v93)
                                                                }
                                                            }
                                                        default:
                                                            var v122 uint32 = uint32(v2)
                                                            _ = v122
                                                        }
                                                        _ = v100
                                                    }
                                                    v123 := v4["hello"]
                                                    fmt.Println(v97)
                                                    _ = v123
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = "foo bar"
                                                    }
                                                    var v124 string = ""
                                                    _ = v124
                                                }
                                            default:
                                                switch "a" {
                                                case "foo bar":
                                                    v93 = (!v10)
                                                case "":
                                                    var v125 int32 = -307
                                                    _ = v125
                                                }
                                            }
                                            v2 += v2
                                            var v126 string = ("\n\t" + "hello")
                                            _ = v126
                                        }
                                        var v127 float32 = 0.0
                                        var v128 float64 = float64(((((1e10 + (v127 * v127)) + (v127 / v127)) / (v127 + ((float32(3.14159) * 0.0) * float32(-2.25)))) / (0.0 - (float32((-2.25 / (0.1 + -2.25))) / v127))))
                                        _ = v127
                                        _ = v128
                                    }
                                }
                                var v129 int32 = 2147483647
                                v130 := GenS12{F0: 1}
                                v10 = (127 < -1)
                                _ = fi92
                                _ = v129
                                _ = v130
                                _ = v93
                            }
                        }
                    }
                }
                for fi131 := 0; fi131 < 2; fi131++ {
                    var v132 uint16 = 879
                    fmt.Println(fi14)
                    _ = fi131
                    _ = v132
                }
                v4["foo bar"] = (true || (1.5 <= 0.1))
                _ = fi14
            }
            var v133 int8 = -128
            var v134 uint8 = 1
            _ = v133
            _ = v134
        }
    }
    fmt.Println(v2, v10)
    v2 = (v2 | (v2 % ((2) | 1)))
    v135 := GenS13{F0: -1}
    var v136 int16 = -195
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v135
    _ = v136
    _ = v2
    _ = v3
    _ = v4
}
