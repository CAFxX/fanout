package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 int64; F1 int64; F2 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8; F1 int8; F2 uint16 }
type GenS9 struct { F0 string; F1 uint; F2 int64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 string; F1 uint; F2 string }
type GenS14 struct { F0 string; F1 int8; F2 bool }
type GenS15 struct { F0 uint32; F1 bool }
type GenS16 struct { F0 uint32; F1 int8 }
type GenS17 struct { F0 uint16; F1 uint64 }
type GenS18 struct { F0 uint; F1 uint64 }
type GenS19 struct { F0 int8; F1 int64 }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 uint8; F1 uint }
type GenS22 struct { F0 uint }
type GenS23 struct { F0 string; F1 uint64 }
type GenS24 struct { F0 uint16; F1 uint16 }
func genFn0(p0 string, p1 uint32) int8 {
    fmt.Println(p0)
    v4 := []int{4294967295, -9223372036854775808, -990}
    _ = p0
    _ = p1
    _ = v4
    return 42
}
func genFn1(p0 int16, p1 uint64) uint64 {
    if (false && true) {
        switch 9223372036854775807 {
        case 0:
            if false {
                fmt.Println(p1, p1)
            } else {
                var v5 uint64 = 4294967295
                if true {
                    var v6 uint64 = v5
                    v6 = (^(v5))
                    for fi7 := 0; fi7 < 9; fi7++ {
                        v8 := []int{fi7}
                        for fi9 := 0; fi9 < 15; fi9++ {
                            var v10 float64 = 0.0
                            _ = fi9
                            _ = v10
                        }
                        v11 := GenS1{F0: 32767, F1: 42}
                        _ = fi7
                        _ = v11
                        _ = v8
                    }
                    _ = v6
                }
                _ = v5
            }
        case -1:
            for fi12 := 0; fi12 < 1; fi12++ {
                if (2147483647 > 2147483647) {
                    v13 := GenS2{F0: 1}
                    for fi14 := 0; fi14 < 9; fi14++ {
                        if ((((1e10 * 0.0) / 1.5) * (3.14159 * 0.0)) < (1.5 - (1.5 - ((3.14159 - 1e10) - 3.14159)))) {
                            switch (fi12 ^ (fi14 / ((fi12) | 1))) {
                            case 523:
                                switch "a" {
                                case "\n\t":
                                    fmt.Println(p0)
                                case "a":
                                    switch fi12 {
                                    case 32768:
                                        v15 := []int{(fi12 * (fi14 + (fi12 >> ((4294967295) & 63))))}
                                        _ = v15
                                    case -501:
                                        var v16 float32 = 0.0
                                        _ = v16
                                    case 128:
                                        v17 := GenS3{F0: fi12, F1: 0, F2: 255}
                                        _ = v17
                                    }
                                default:
                                    p0 = p0
                                }
                            case 1:
                                var v18 string = "\n\t"
                                _ = v18
                            }
                        }
                        _ = fi14
                    }
                    _ = v13
                }
                var v19 int64 = 336
                _ = fi12
                _ = v19
            }
        default:
            v20 := []int{32767}
            _ = v20
        }
        var v21 int32 = 1
        _ = v21
    } else {
        p0 += (p0 & -1)
        switch 9223372036854775807 {
        case 1:
            for fi22 := 0; fi22 < 1; fi22++ {
                v23 := []int{(^(fi22)), (-(fi22))}
                for fi24 := 0; fi24 < 10; fi24++ {
                    switch ("hello" + ("" + "hello")) {
                    case "\n\t":
                        var v25 uint32 = 2147483648
                        _ = v25
                    case "hello":
                        if (((float64((0.1 - 3.14159)) + (1e10 - (-2.25 * ((1e10 / 1.0) / 0.1)))) <= (0.1 * 0.1)) || ((439 < 524) || (uint8(p1) < 174))) {
                            switch fi24 {
                            case -129:
                                var v26 bool = false
                                _ = v26
                            case 32768:
                                var v27 float32 = 1.5
                                _ = v27
                            case -295:
                                var v28 int = fi22
                                _ = v28
                            }
                            v29 := GenS4{F0: 42, F1: 1, F2: "\n\t"}
                            v30 := GenS1{F0: 32767, F1: 65535}
                            var v31 int64 = 9223372036854775807
                            var v32 int16 = (p0 << ((823) & 63))
                            _ = v29
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    default:
                        p1 = p1
                    }
                    var v33 int = 42
                    _ = fi24
                    _ = v33
                }
                _ = fi22
                _ = v23
            }
        }
        if true {
            v34 := GenS5{F0: p1}
            var v35 string = ("foo bar" + (("" + "") + "foo bar"))
            _ = v34
            _ = v35
        } else {
            v36 := GenS6{F0: (-(p0))}
            fmt.Println(p1, p1)
            _ = v36
        }
    }
    var v37 bool = true
    _ = p0
    _ = p1
    _ = v37
    return 1
}

func main() {
    var v1 float64 = (1e10 + float64(1e10))
    var v2 uint = 2
    v3 := []uint32{}
    switch "hello" {
    case "":
        var v38 int16 = 0
        _ = v38
    case "\n\t":
        var v39 string = "\n\t"
        _ = v39
    default:
        v1 = v1
    }
    if ((1e10 <= -2.25) && false) {
        var v40 int64 = -9223372036854775808
        _ = v40
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v41 uint64 = 2
        var v42 float64 = v1
        _ = v41
        _ = v42
    }
    if true {
        var v43 uint32 = 988
        _ = v43
    } else {
        for fi44 := 0; fi44 < 20; fi44++ {
            if (!true) {
                v3 = append(v3, 534)
                if true {
                    v2 = v2
                }
            }
            var v45 uint = v2
            v46 := GenS7{F0: 4294967295}
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    if false {
        v1 += v1
        v1 += float64(0.1)
        v47 := GenS6{F0: 398}
        _ = v47
    } else {
        switch -926 {
        case 1:
            for fi48 := 0; fi48 < 13; fi48++ {
                v2 -= (v2 << ((v2) & 63))
                var v49 int16 = 635
                var v50 int64 = -9223372036854775808
                _ = fi48
                _ = v49
                _ = v50
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v51 := GenS7{F0: 513}
        _ = v51
    }
    var v52 uint = v2
    v53 := GenS7{F0: (v52 + 0)}
    v2 = (v52 + (v2 % ((v2) | 1)))
    switch ("foo bar" + ("" + "a")) {
    case "foo bar":
        switch -337 {
        case -897:
            v2 = 65535
        case 2147483647:
            if false {
                var v54 uint8 = 34
                if (!true) {
                    _ = len(v3)
                    for fi55 := 0; fi55 < 13; fi55++ {
                        v54 -= v54
                        _ = fi55
                    }
                } else {
                    var v56 uint16 = 32767
                    _ = len(v3)
                    _ = v56
                }
                v57 := GenS8{F0: 100, F1: -123, F2: 1}
                _ = v54
                _ = v57
            } else {
                _ = len(v3)
                for fi58 := 0; fi58 < 18; fi58++ {
                    if (!false) {
                        switch "foo bar" {
                        case "foo bar":
                            v3 = append(v3, 4294967295)
                        case "hello":
                            if ((980 > 42) && false) {
                                for fi59 := 0; fi59 < 7; fi59++ {
                                    var v60 int32 = 910
                                    switch (fi58 >> (((^(v2))) & 63)) {
                                    case 0:
                                        v1 = (v1 / v1)
                                        v61 := GenS9{F0: "hello", F1: v2, F2: 9223372036854775807}
                                        _ = v61
                                    case 2147483647:
                                        fmt.Println(v1, v52, v52)
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        _ = len(v3)
                                    case "hello":
                                        var v62 int16 = 1
                                        _ = v62
                                    case "\n\t":
                                        var v63 bool = true
                                        if true {
                                            switch ("foo bar" + ("hello" + ("hello" + ("foo bar" + "hello")))) {
                                            case "hello":
                                                var v64 uint8 = 1
                                                _ = v64
                                            case "a":
                                                v1 -= v1
                                            }
                                            v52 = (v52 * 1)
                                            switch fi58 {
                                            case -1:
                                                v1 += ((v1 + v1) / (v1 * v1))
                                            default:
                                                switch "" {
                                                case "a":
                                                    v52 += (v52 / ((v2) | 1))
                                                    if (0 > 98) {
                                                        switch fi59 {
                                                        case 9223372036854775807:
                                                            if (v1 > (v1 + v1)) {
                                                                v1 = v1
                                                                var v65 int16 = 0
                                                                _ = v65
                                                            } else {
                                                                fmt.Println(v60)
                                                                if (!(309 < 42)) {
                                                                    fmt.Println(fi59, fi58, fi59)
                                                                    var v66 int64 = 9223372036854775807
                                                                    _ = v66
                                                                }
                                                                var v67 int16 = 1
                                                                _ = v67
                                                            }
                                                        case 396:
                                                            v68 := GenS2{F0: -128}
                                                            _ = v68
                                                        case 32768:
                                                            var v69 uint8 = 128
                                                            _ = v69
                                                        default:
                                                            v3 = append(v3, 2147483648)
                                                        }
                                                    }
                                                case "hello":
                                                    if len(v3) > 0 {
                                                        v3[0] = 65535
                                                    }
                                                default:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v70 int32 = -1
                                                        _ = v70
                                                    default:
                                                        if v63 {
                                                            v1 = v1
                                                        } else {
                                                            v2 -= (v2 ^ (v52 + v2))
                                                            var v71 bool = true
                                                            v72 := GenS1{F0: 873, F1: v2}
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                    }
                                                }
                                            }
                                            if ((fi59 <= fi58) || v63) {
                                                v60 += v60
                                                v73 := GenS10{F0: 65535}
                                                v74 := GenS7{F0: (^(v2))}
                                                _ = v73
                                                _ = v74
                                            } else {
                                                _ = len(v3)
                                                v75 := GenS11{F0: 69}
                                                var v76 int64 = 42
                                                _ = v75
                                                _ = v76
                                            }
                                            if v63 {
                                                fmt.Println(v52, v52)
                                                var v77 float64 = 1.5
                                                _ = v77
                                            } else {
                                                if (v63 && ((uint16(v52) == 697) || true)) {
                                                    switch ("\n\t" + "hello") {
                                                    case "":
                                                        v78 := GenS2{F0: -128}
                                                        _ = v78
                                                    case "foo bar":
                                                        if (false && v63) {
                                                            switch (fi58 % ((2147483647) | 1)) {
                                                            case 32768:
                                                                v52 = 4294967295
                                                            case -2147483648:
                                                                switch ((("a" + (("a" + ("a" + ((("a" + ("a" + "a")) + "\n\t") + ("a" + "hello")))) + "\n\t")) + "\n\t") + ("foo bar" + "hello")) {
                                                                case "foo bar":
                                                                    v52 = v2
                                                                }
                                                            }
                                                        } else {
                                                            switch (fi59 << ((0) & 63)) {
                                                            case 526:
                                                                fmt.Println(v52, v52)
                                                            default:
                                                                v52 -= (^(v52))
                                                            }
                                                        }
                                                    case "a":
                                                        if false {
                                                            if (false || (-9223372036854775808 >= -9223372036854775808)) {
                                                                var v79 int64 = 1
                                                                var v80 int64 = (v79 | -1)
                                                                _ = v79
                                                                _ = v80
                                                            } else {
                                                                v2 += 4
                                                                v1 = v1
                                                            }
                                                            switch ("" + "foo bar") {
                                                            case "":
                                                                v2 += (v2 ^ 4294967295)
                                                            default:
                                                                v2 = 170
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    if true {
                                                        if ((fi59 % ((2147483647) | 1)) != -9223372036854775808) {
                                                            switch 32768 {
                                                            case 276:
                                                                if (((^(fi59)) >= fi59) || (v2 == v52)) {
                                                                    switch fi58 {
                                                                    case 1:
                                                                        if (((3.14159 - (float32(((((-2.25 / 3.14159) - ((((v1 * -2.25) / v1) * v1) - (1e10 + v1))) / 3.14159) - (v1 + (v1 + v1)))) + (1.5 * 0.1))) - 1.5) <= (1e10 + ((0.0 / 3.14159) / 1.0))) {
                                                                            if true {
                                                                                v63 = (v63 && v63)
                                                                                var v81 uint = v2
                                                                                var v82 uint64 = genFn1(-1, 804)
                                                                                _ = v81
                                                                                _ = v82
                                                                            }
                                                                            var v83 float32 = 3.14159
                                                                            v60 = v60
                                                                            _ = v83
                                                                        } else {
                                                                            fmt.Println(fi59, fi59, fi58)
                                                                            var v84 uint16 = 855
                                                                            var v85 bool = ((1.5 > (1.5 / (v1 / float64((((3.14159 / 1e10) + 1.5) + 3.14159))))) && (!(9223372036854775807 > -9223372036854775808)))
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                v3 = append(v3, 2)
                                                                            default:
                                                                                if (3.14159 == ((float32(0.0) / 1.5) - 0.0)) {
                                                                                    fmt.Println(fi59, fi58)
                                                                                    v86 := GenS6{F0: 42}
                                                                                    var v87 uint8 = 42
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                }
                                                                            }
                                                                            _ = v84
                                                                            _ = v85
                                                                        }
                                                                    case 2147483647:
                                                                        var v88 int8 = 0
                                                                        var v89 float64 = 1.5
                                                                        v90 := GenS9{F0: ((("hello" + ("foo bar" + "a")) + "") + "\n\t"), F1: v2, F2: -9223372036854775808}
                                                                        _ = v88
                                                                        _ = v89
                                                                        _ = v90
                                                                    case 904:
                                                                        var v91 uint32 = 2147483647
                                                                        _ = v91
                                                                    }
                                                                    if v63 {
                                                                        v63 = false
                                                                    }
                                                                    if v63 {
                                                                        v3 = append(v3, 1)
                                                                        v92 := GenS8{F0: 124, F1: genFn0("hello", 923), F2: 2}
                                                                        _ = v92
                                                                    }
                                                                }
                                                            default:
                                                                v2 += 1
                                                            }
                                                            v63 = true
                                                            switch (fi58 / (((fi58 >> (((^(v52))) & 63))) | 1)) {
                                                            case -212:
                                                                fmt.Println(v60, v52, fi58)
                                                            case 128:
                                                                var v93 uint32 = 0
                                                                _ = v93
                                                            case -331:
                                                                var v94 int = 127
                                                                _ = v94
                                                            }
                                                        } else {
                                                            var v95 int64 = 429
                                                            v2 += v52
                                                            switch ((("\n\t" + ("a" + "foo bar")) + "a") + "a") {
                                                            case "\n\t":
                                                                switch 1 {
                                                                case -957:
                                                                    fmt.Println(v52, v63, fi59)
                                                                case -845:
                                                                    fmt.Println(v63, v1)
                                                                }
                                                                switch (fi59 << ((v52) & 63)) {
                                                                case 4294967295:
                                                                    v63 = (v60 != v60)
                                                                case 127:
                                                                    if (0 <= v95) {
                                                                        if (!v63) {
                                                                            v96 := GenS4{F0: v95, F1: v95, F2: (((("foo bar" + "a") + "a") + "") + "hello")}
                                                                            v52 = 255
                                                                            _ = v96
                                                                        }
                                                                    } else {
                                                                        v63 = v63
                                                                        v97 := GenS2{F0: -15}
                                                                        _ = v97
                                                                    }
                                                                default:
                                                                    v1 = (1.5 + 0.1)
                                                                }
                                                                switch fi59 {
                                                                case 42:
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        fmt.Println(v1, fi58)
                                                                        if (3.14159 <= (1e10 / float32(float64(float32(v1))))) {
                                                                            v98 := GenS9{F0: "a", F1: 65535, F2: (v95 % ((-9223372036854775808) | 1))}
                                                                            var v99 uint32 = 936
                                                                            _ = v98
                                                                            _ = v99
                                                                        }
                                                                        v100 := GenS3{F0: fi58, F1: 0, F2: v2}
                                                                        _ = v100
                                                                    case "hello":
                                                                        _ = len(v3)
                                                                    case "\n\t":
                                                                        var v101 float32 = 1.5
                                                                        _ = v101
                                                                    default:
                                                                        if true {
                                                                            var v102 float32 = 0.1
                                                                            _ = v102
                                                                        } else {
                                                                            v103 := GenS10{F0: 2147483648}
                                                                            v1 = 1.5
                                                                            if (-2.25 < 3.14159) {
                                                                                if v63 {
                                                                                    fmt.Println(v63, v52)
                                                                                }
                                                                                if v63 {
                                                                                    switch -999 {
                                                                                    case 948:
                                                                                        if v63 {
                                                                                            var v104 int = 42
                                                                                            _ = v104
                                                                                        } else {
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v105 bool = v63
                                                                                    v1 = (v1 * float64(1.5))
                                                                                    _ = v105
                                                                                }
                                                                            }
                                                                            _ = v103
                                                                        }
                                                                    }
                                                                case -1:
                                                                    v3 = append(v3, uint32(v2))
                                                                case -981:
                                                                    if ((((v2 > v52) && v63) && ((1.5 / 0.1) > (1.5 * ((3.14159 + 3.14159) * float32((float64((0.0 - -2.25)) - (1e10 + v1))))))) || ((^(fi58)) >= fi58)) {
                                                                        if false {
                                                                            var v106 int64 = -9223372036854775808
                                                                            fmt.Println(fi58, v52)
                                                                            if true {
                                                                                switch ("a" + "foo bar") {
                                                                                case "foo bar":
                                                                                    v63 = (4294967295 >= 2147483648)
                                                                                case "hello":
                                                                                    v52 -= (v2 ^ (v52 - (^(v2))))
                                                                                case "\n\t":
                                                                                    fmt.Println(v52)
                                                                                default:
                                                                                    v2 -= v2
                                                                                }
                                                                                fmt.Println(v63, v60)
                                                                                if ((v2 | (v52 ^ v2)) == (v52 / (((v2 | (v2 & v2))) | 1))) {
                                                                                    v60 = 247
                                                                                    v107 := GenS11{F0: 174}
                                                                                    _ = v107
                                                                                } else {
                                                                                    var v108 uint = 403
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        v52 += 221
                                                                                    }
                                                                                    var v109 float64 = v1
                                                                                    var v110 int = int(v95)
                                                                                    var v111 uint16 = 1
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = 1
                                                                                    }
                                                                                    _ = v108
                                                                                    _ = v109
                                                                                    _ = v110
                                                                                    _ = v111
                                                                                }
                                                                                v63 = (v106 == (v95 | v95))
                                                                                var v112 int32 = (v60 * (v60 % ((v60) | 1)))
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 0
                                                                                }
                                                                                var v113 float32 = 1e10
                                                                                _ = v112
                                                                                _ = v113
                                                                            }
                                                                            _ = v106
                                                                        } else {
                                                                            var v114 string = ("foo bar" + (("" + "foo bar") + "a"))
                                                                            _ = v114
                                                                        }
                                                                    }
                                                                }
                                                            case "":
                                                                v115 := GenS12{F0: (v60 * -2147483648)}
                                                                _ = v115
                                                            default:
                                                                v3 = append(v3, 42)
                                                            }
                                                            _ = v95
                                                        }
                                                    } else {
                                                        v116 := GenS13{F0: ("a" + ("a" + "a")), F1: 255, F2: "\n\t"}
                                                        fmt.Println(fi58, v1)
                                                        _ = v116
                                                    }
                                                    v3 = append(v3, 2147483648)
                                                    var v117 string = ("\n\t" + ("foo bar" + "\n\t"))
                                                    var v118 int32 = -916
                                                    v3 = append(v3, 4294967295)
                                                    _ = v117
                                                    _ = v118
                                                }
                                                v119 := GenS6{F0: 1}
                                                var v120 uint = v52
                                                if v63 {
                                                    v1 -= float64(3.14159)
                                                    var v121 uint8 = 127
                                                    var v122 bool = (v63 || true)
                                                    var v123 uint32 = 189
                                                    _ = v121
                                                    _ = v122
                                                    _ = v123
                                                } else {
                                                    var v124 int16 = 42
                                                    _ = v124
                                                }
                                                if (!v63) {
                                                    var v125 int = (^(fi58))
                                                    _ = v125
                                                }
                                                v1 = (v1 * float64(-2.25))
                                                fmt.Println(v60, fi59, v63)
                                                if (0.1 >= v1) {
                                                    v60 -= (v60 % ((int32(v2)) | 1))
                                                    var v126 int16 = 1
                                                    switch fi58 {
                                                    case -1:
                                                        v127 := GenS12{F0: 0}
                                                        _ = v127
                                                    case 571:
                                                        if v63 {
                                                            var v128 int64 = 42
                                                            v129 := GenS9{F0: ("foo bar" + "foo bar"), F1: 4294967295, F2: -1}
                                                            _ = len(v3)
                                                            var v130 int8 = 100
                                                            var v131 uint16 = 232
                                                            var v132 int8 = genFn0((("foo bar" + "foo bar") + "foo bar"), 0)
                                                            _ = v128
                                                            _ = v129
                                                            _ = v130
                                                            _ = v131
                                                            _ = v132
                                                        }
                                                    default:
                                                        var v133 uint32 = 51
                                                        _ = v133
                                                    }
                                                    _ = v126
                                                }
                                                v134 := GenS10{F0: 2}
                                                var v135 int8 = 48
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                switch 615 {
                                                case 1:
                                                    var v136 uint64 = 0
                                                    _ = v136
                                                case 128:
                                                    fmt.Println(fi58, v120)
                                                default:
                                                    v60 = (v60 / ((v60) | 1))
                                                }
                                                _ = v119
                                                _ = v120
                                                _ = v134
                                                _ = v135
                                            }
                                        }
                                        switch (^(fi58)) {
                                        case -9223372036854775808:
                                            var v137 uint16 = 255
                                            _ = v137
                                        case 1:
                                            var v138 bool = v63
                                            _ = v138
                                        case -235:
                                            switch "a" {
                                            case "foo bar":
                                                if v63 {
                                                    v52 = (v52 + v52)
                                                    switch (fi58 >> ((4294967295) & 63)) {
                                                    case 32767:
                                                        v139 := GenS14{F0: ("hello" + "\n\t"), F1: -1, F2: false}
                                                        _ = v139
                                                    case 733:
                                                        v140 := GenS15{F0: 955, F1: v63}
                                                        var v141 uint16 = 335
                                                        _ = v140
                                                        _ = v141
                                                    default:
                                                        if ((!(!(!v63))) && true) {
                                                            v52 -= (v2 + 1)
                                                            v142 := GenS13{F0: "hello", F1: v2, F2: "foo bar"}
                                                            _ = v142
                                                        } else {
                                                            if v63 {
                                                                if (!(-1 >= 1)) {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    var v143 uint8 = 128
                                                                    _ = v143
                                                                }
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    if v63 {
                                                                        fmt.Println(v60, v63)
                                                                    }
                                                                case "hello":
                                                                    fmt.Println(v52, fi59)
                                                                case "foo bar":
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                                if false {
                                                                    if v63 {
                                                                        if v63 {
                                                                            switch (fi59 & 226) {
                                                                            case 32767:
                                                                                v2 -= v52
                                                                            case -9223372036854775808:
                                                                                if (v2 == 65535) {
                                                                                    v60 += 2147483647
                                                                                } else {
                                                                                    v52 += 18446744073709551615
                                                                                    if v63 {
                                                                                        v60 -= (v60 % (((v60 | (v60 | 454))) | 1))
                                                                                    }
                                                                                    if (!(1 >= 4294967295)) {
                                                                                        if (!(v60 != v60)) {
                                                                                            v3 = append(v3, 1)
                                                                                            v63 = (v52 > (v2 + v2))
                                                                                            v144 := GenS16{F0: 2, F1: genFn0(("" + "\n\t"), 42)}
                                                                                            var v145 uint16 = 371
                                                                                            _ = v144
                                                                                            _ = v145
                                                                                        } else {
                                                                                            var v146 int64 = 42
                                                                                            var v147 int16 = 1000
                                                                                            _ = v146
                                                                                            _ = v147
                                                                                        }
                                                                                        if (-842 >= -488) {
                                                                                            v52 = v52
                                                                                        } else {
                                                                                            fmt.Println(v60, v2, fi59)
                                                                                            v63 = ((1e10 / (((3.14159 * ((0.0 / (1.5 - -2.25)) / (-2.25 / 1e10))) / (-2.25 * -2.25)) / float32(v1))) != 3.14159)
                                                                                        }
                                                                                        fmt.Println(v52, v1)
                                                                                    }
                                                                                }
                                                                            }
                                                                            fmt.Println(v2, fi59, fi59)
                                                                        } else {
                                                                            v63 = v63
                                                                            switch (fi58 - fi58) {
                                                                            case 42:
                                                                                v52 -= 4294967295
                                                                            case 4294967295:
                                                                                v148 := GenS17{F0: 32767, F1: 1}
                                                                                _ = v148
                                                                            case -116:
                                                                                v149 := GenS13{F0: ("hello" + "hello"), F1: v52, F2: "\n\t"}
                                                                                _ = v149
                                                                            }
                                                                        }
                                                                        v1 -= ((v1 * 1.5) - v1)
                                                                    }
                                                                    v150 := GenS7{F0: 0}
                                                                    var v151 uint = (v52 | v52)
                                                                    _ = v150
                                                                    _ = v151
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = 65535
                                                    }
                                                    v152 := GenS4{F0: -9223372036854775808, F1: 0, F2: ("a" + ("a" + ("foo bar" + "")))}
                                                    _ = v152
                                                } else {
                                                    v60 = 1
                                                    v153 := GenS13{F0: "", F1: (v52 - v2), F2: (("hello" + (("a" + "hello") + "hello")) + ("\n\t" + ("a" + "\n\t")))}
                                                    if len(v3) > 0 {
                                                        v3[0] = 0
                                                    }
                                                    _ = v153
                                                }
                                                fmt.Println(fi59, fi58, v1)
                                            case "a":
                                                v60 += v60
                                            case "\n\t":
                                                v154 := GenS9{F0: "a", F1: (v52 & 255), F2: 1}
                                                _ = v154
                                            }
                                        }
                                        _ = v63
                                    default:
                                        switch ("" + "") {
                                        case "a":
                                            var v155 uint8 = 243
                                            _ = v155
                                        case "":
                                            if (1 >= 1) {
                                                switch "hello" {
                                                case "":
                                                    if false {
                                                        var v156 float32 = ((((-2.25 * ((1e10 / 1.0) - (1e10 / float32(float64(float32(v1)))))) / 1e10) / 3.14159) - 1e10)
                                                        v157 := GenS6{F0: 1}
                                                        v156 = (0.1 + (v156 / v156))
                                                        _ = v156
                                                        _ = v157
                                                    }
                                                }
                                                var v158 uint64 = 0
                                                _ = v158
                                            }
                                        case "\n\t":
                                            if len(v3) > 0 {
                                                v3[0] = 1
                                            }
                                        }
                                    }
                                    _ = fi59
                                    _ = v60
                                }
                                var v159 bool = (!false)
                                _ = v159
                            } else {
                                var v160 bool = false
                                if len(v3) > 0 {
                                    v3[0] = 569
                                }
                                fmt.Println(v1, v52)
                                _ = v160
                            }
                        case "a":
                            var v161 int64 = -1
                            _ = v161
                        }
                        switch fi58 {
                        case 0:
                            v52 = (v52 | v52)
                        case 9223372036854775807:
                            switch ((("a" + "foo bar") + "\n\t") + "a") {
                            case "":
                                for fi162 := 0; fi162 < 6; fi162++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v163 := GenS18{F0: 47, F1: 1}
                                    _ = fi162
                                    _ = v163
                                }
                            case "\n\t":
                                var v164 float32 = (0.1 - -2.25)
                                _ = v164
                            default:
                                v165 := GenS19{F0: 42, F1: -117}
                                _ = v165
                            }
                        case 127:
                            v166 := GenS11{F0: 82}
                            _ = v166
                        default:
                            for fi167 := 0; fi167 < 12; fi167++ {
                                if (!(65535 >= 32767)) {
                                    fmt.Println(fi58, v52)
                                } else {
                                    v52 = 661
                                }
                                v2 = 255
                                fmt.Println(fi58, v52)
                                _ = fi167
                            }
                        }
                        fmt.Println(v1, v52)
                    } else {
                        switch "\n\t" {
                        case "":
                            fmt.Println(v52)
                        case "hello":
                            _ = len(v3)
                        case "foo bar":
                            v52 -= (^(v2))
                        }
                        var v168 float32 = ((3.14159 * float32(v1)) * 1e10)
                        var v169 uint64 = 42
                        _ = v168
                        _ = v169
                    }
                    _ = fi58
                }
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
            }
        case 1:
            v52 = 521
        }
    case "a":
        for fi170 := 0; fi170 < 11; fi170++ {
            if false {
                if (2147483647 != 2) {
                    v1 -= v1
                    if (!(-418 == -1)) {
                        for fi171 := 0; fi171 < 16; fi171++ {
                            if ((v2 % ((v2) | 1)) != v52) {
                                v52 = (v2 + (^(v52)))
                                var v172 float32 = ((((1.5 + 1.5) - 1e10) / 1.0) + 0.1)
                                _ = v172
                            } else {
                                v173 := GenS6{F0: 1000}
                                v174 := GenS11{F0: 0}
                                _ = v173
                                _ = v174
                            }
                            var v175 int32 = 659
                            v175 -= v175
                            v3 = append(v3, 42)
                            var v176 uint64 = 0
                            v176 -= genFn1(927, (v176 << (((v52 << ((0) & 63))) & 63)))
                            switch ((("\n\t" + "hello") + (("foo bar" + "a") + ("hello" + (("foo bar" + ("foo bar" + ("a" + ""))) + "")))) + "") {
                            case "hello":
                                if (739 < 937) {
                                    if (genFn1(1, v176) >= v176) {
                                        v177 := GenS20{F0: -2147483648}
                                        v3 = append(v3, 65535)
                                        fmt.Println(v2, fi171)
                                        _ = v177
                                    } else {
                                        var v178 uint8 = 42
                                        if true {
                                            v175 += (v175 << (((^(v2))) & 63))
                                            switch (fi170 * int(v176)) {
                                            case -9223372036854775808:
                                                v179 := GenS21{F0: v178, F1: v2}
                                                _ = v179
                                            case 42:
                                                switch (^(fi171)) {
                                                case 32767:
                                                    v180 := GenS22{F0: (v2 | v2)}
                                                    _ = v180
                                                case -1:
                                                    fmt.Println(v2, v1, v1)
                                                case 2147483647:
                                                    if len(v3) > 0 {
                                                        v3[0] = 725
                                                    }
                                                default:
                                                    _ = len(v3)
                                                }
                                            default:
                                                var v181 int = 481
                                                _ = v181
                                            }
                                            var v182 uint16 = 42
                                            _ = v182
                                        } else {
                                            v1 = (v1 - float64(1e10))
                                            v52 += (^(v52))
                                            var v183 string = "\n\t"
                                            _ = v183
                                        }
                                        v176 = v176
                                        _ = v178
                                    }
                                } else {
                                    var v184 int8 = -62
                                    if false {
                                        v2 = v52
                                    }
                                    var v185 bool = ((v52 & 728) >= v2)
                                    _ = v184
                                    _ = v185
                                }
                            case "\n\t":
                                v175 = (v175 + v175)
                            case "":
                                v176 += genFn1(1, (^(v176)))
                                v186 := GenS23{F0: "foo bar", F1: genFn1(32767, v176)}
                                _ = v186
                            default:
                                v187 := GenS24{F0: 255, F1: 32767}
                                _ = v187
                            }
                            _ = fi171
                            _ = v175
                            _ = v176
                        }
                    }
                }
                v188 := GenS22{F0: v52}
                if true {
                    switch ("hello" + "\n\t") {
                    case "":
                        if (!((-2.25 - float32(v1)) != (-2.25 * 0.1))) {
                            v2 = v2
                        }
                    default:
                        fmt.Println(v1)
                    }
                }
                _ = v188
            }
            _ = fi170
        }
    default:
        v52 = v52
    }
    fmt.Println(v2, v1, v52)
    switch -940 {
    case 1:
        fmt.Println(v52, v1)
    default:
        for fi189 := 0; fi189 < 20; fi189++ {
            for fi190 := 0; fi190 < 9; fi190++ {
                if (42 <= 127) {
                    if true {
                        fmt.Println(v1, v52, v2)
                        v52 += 42
                        _ = len(v3)
                        var v191 string = ""
                        _ = v191
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 2
                        }
                    }
                } else {
                    v52 = v52
                    var v192 uint8 = 42
                    v193 := GenS20{F0: -970}
                    _ = v192
                    _ = v193
                }
                v194 := GenS20{F0: 0}
                var v195 float32 = 0.0
                _ = fi190
                _ = v194
                _ = v195
            }
            v1 -= v1
            if true {
                fmt.Println(v52, v1)
            } else {
                var v196 uint16 = 162
                var v197 bool = (65535 != v2)
                var v198 int32 = -228
                _ = v196
                _ = v197
                _ = v198
            }
            _ = fi189
        }
    }
    if (!((!true) && true)) {
        v199 := GenS10{F0: 450}
        v200 := GenS5{F0: 18446744073709551615}
        v2 -= (^(v52))
        _ = v199
        _ = v200
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v53
}
