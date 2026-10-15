package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int; F1 bool; F2 uint }
type GenS7 struct { F0 bool; F1 string; F2 uint64 }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int8; F1 bool; F2 int }
type GenS11 struct { F0 string; F1 int32; F2 int32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint16; F1 bool }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 string; F1 string }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int; F1 bool; F2 uint }
type GenS19 struct { F0 uint8; F1 int32 }
type GenS20 struct { F0 string; F1 int16; F2 int16 }
func genFn0(p0 string, p1 uint32) int {
    if false {
        switch ((p0 + (p0 + "hello")) + (p0 + p0)) {
        case "":
            var v7 uint = 443
            _ = v7
        case "hello":
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: 2, F1: -570}
                _ = fi8
                _ = v9
            }
        case "a":
            var v10 int16 = 0
            _ = v10
        default:
            v11 := []int{4294967295}
            _ = v11
        }
    } else {
        var v12 string = p0
        _ = v12
    }
    if (float32(3.14159) == float32(float64(-2.25))) {
        var v13 int = -9223372036854775808
        _ = v13
    }
    v14 := []int{4294967295, 897, 9223372036854775807, 32768}
    if (-246 >= -980) {
        for fi15 := 0; fi15 < 15; fi15++ {
            fmt.Println(p0)
            v16 := GenS2{F0: (^(fi15)), F1: 2147483647, F2: (p1 >> ((2) & 63))}
            _ = fi15
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := GenS2{F0: (fi17 - fi17), F1: -1, F2: p1}
            _ = fi17
            _ = v18
        }
        p1 = (p1 - p1)
    }
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = v14
    return -391
}
func genFn1() uint64 {
    if true {
        var v19 int32 = 0
        v20 := GenS3{F0: 1, F1: 1}
        _ = v19
        _ = v20
    }
    if (true || ((!(("foo bar" >= ("\n\t" + "a")) && true)) || true)) {
        v21 := GenS1{F0: 65535, F1: 9223372036854775807}
        var v22 int16 = -1
        var v23 float64 = (-2.25 * (1e10 * 0.0))
        _ = v21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        v25 := []int{fi24}
        v25 = append(v25, fi24)
        var v26 string = "\n\t"
        _ = fi24
        _ = v25
        _ = v26
    }
    if true {
        if ((0.0 - (-2.25 - 3.14159)) != 0.1) {
            var v27 int8 = 1
            switch "" {
            case "hello":
                switch "\n\t" {
                case "":
                    v27 = (v27 & v27)
                case "foo bar":
                    var v28 uint64 = 2
                    _ = v28
                case "hello":
                    if (203 == 42) {
                        switch 0 {
                        case 775:
                            var v29 bool = (!(704 != 2147483648))
                            _ = v29
                        case 127:
                            v27 = -1
                        default:
                            for fi30 := 0; fi30 < 13; fi30++ {
                                v31 := []int{(fi30 / (((fi30 & fi30)) | 1)), (^(fi30))}
                                v27 = (v27 + (v27 % ((v27) | 1)))
                                for fi32 := 0; fi32 < 5; fi32++ {
                                    switch (("\n\t" + "hello") + (("\n\t" + (("hello" + "hello") + "\n\t")) + ("a" + ""))) {
                                    case "foo bar":
                                        var v33 int16 = 0
                                        _ = v33
                                    case "\n\t":
                                        switch fi30 {
                                        case 32767:
                                            v27 = (v27 + 127)
                                        case 32768:
                                            if false {
                                                _ = len(v31)
                                                switch "\n\t" {
                                                case "hello":
                                                    v27 -= v27
                                                case "a":
                                                    if true {
                                                        if true {
                                                            switch (fi32 / ((fi30) | 1)) {
                                                            case 0:
                                                                fmt.Println(v27, fi30, fi32)
                                                            case 32767:
                                                                _ = len(v31)
                                                            case 2147483647:
                                                                v31 = append(v31, genFn0("foo bar", 764))
                                                            }
                                                            if true {
                                                                v27 += (v27 >> ((661) & 63))
                                                            } else {
                                                                v27 -= 91
                                                            }
                                                            fmt.Println(fi32)
                                                        }
                                                        v27 += (-(v27))
                                                    } else {
                                                        var v34 int = fi30
                                                        _ = v34
                                                    }
                                                case "\n\t":
                                                    switch fi32 {
                                                    case 2147483647:
                                                        var v35 uint64 = 876
                                                        v36 := GenS1{F0: 65535, F1: 1000}
                                                        _ = v35
                                                        _ = v36
                                                    case -9223372036854775808:
                                                        if false {
                                                            v37 := GenS2{F0: genFn0((("foo bar" + "foo bar") + ("hello" + ("a" + "\n\t"))), 0), F1: 2147483647, F2: 65535}
                                                            _ = v37
                                                        }
                                                    default:
                                                        if (471 != 568) {
                                                            if false {
                                                                fmt.Println(fi30, fi32, fi30)
                                                            } else {
                                                                v27 += 100
                                                                v27 += (v27 / (((v27 << ((255) & 63))) | 1))
                                                                v38 := GenS4{F0: 0, F1: int16(v27)}
                                                                _ = v38
                                                            }
                                                            if len(v31) >= 2 {
                                                                v31 = v31[1:len(v31)]
                                                            }
                                                            var v39 uint8 = 128
                                                            v40 := GenS3{F0: 228, F1: 232}
                                                            switch ("a" + ("" + "foo bar")) {
                                                            case "":
                                                                var v41 int = genFn0("", 521)
                                                                v41 = -2147483648
                                                                fmt.Println(v41, v41, v27)
                                                                _ = v41
                                                            case "\n\t":
                                                                fmt.Println(fi32)
                                                            default:
                                                                var v42 int = (fi30 % (((fi32 ^ (fi30 ^ fi32))) | 1))
                                                                _ = v42
                                                            }
                                                            _ = v39
                                                            _ = v40
                                                        } else {
                                                            var v43 float32 = (((0.1 / 1e10) * (-2.25 - float32((((1e10 / 3.14159) + (0.0 * -2.25)) / (3.14159 - (((1.5 * (0.1 + -2.25)) + 1e10) - 3.14159)))))) + -2.25)
                                                            v43 -= (float32((1e10 * -2.25)) * (v43 + (1e10 / (v43 - v43))))
                                                            _ = v43
                                                        }
                                                    }
                                                default:
                                                    switch (fi32 / ((fi32) | 1)) {
                                                    case 9223372036854775807:
                                                        if true {
                                                            v44 := GenS5{F0: genFn0(("foo bar" + "foo bar"), 65535), F1: 4294967295}
                                                            _ = v44
                                                        }
                                                    case -1:
                                                        var v45 string = "foo bar"
                                                        _ = v45
                                                    default:
                                                        switch (fi30 & genFn0("", 974)) {
                                                        case 0:
                                                            v27 = (v27 - (v27 / ((100) | 1)))
                                                        default:
                                                            v27 = (v27 / ((v27) | 1))
                                                        }
                                                    }
                                                }
                                                var v46 float64 = 1e10
                                                _ = v46
                                            } else {
                                                var v47 uint8 = 127
                                                _ = v47
                                            }
                                        }
                                    }
                                    _ = fi32
                                }
                                _ = fi30
                                _ = v31
                            }
                        }
                        var v48 bool = false
                        _ = v48
                    } else {
                        var v49 int32 = 2147483647
                        _ = v49
                    }
                default:
                    v50 := GenS4{F0: 42, F1: 42}
                    _ = v50
                }
            case "":
                fmt.Println(v27)
            case "a":
                for fi51 := 0; fi51 < 12; fi51++ {
                    fmt.Println(v27, fi51)
                    _ = fi51
                }
            }
            var v52 int = genFn0("foo bar", 202)
            _ = v27
            _ = v52
        } else {
            var v53 int16 = 984
            _ = v53
        }
    } else {
        v54 := []int{-9223372036854775808, 978, 32768}
        _ = v54
    }
    v55 := []int{22}
    _ = v55
    return 1
}

func main() {
    var v1 int8 = 70
    v2 := make(map[string]bool)
    var v3 int64 = -9223372036854775808
    var v4 uint = 255
    v5 := make(map[uint32]string)
    v6 := make(map[string]int)
    for fi56 := 0; fi56 < 5; fi56++ {
        if (true || (0 == 1)) {
            var v57 uint64 = 597
            _ = v57
        }
        v58 := []int{128, genFn0("hello", 2), 32768, (fi56 ^ (^(fi56)))}
        _ = fi56
        _ = v58
    }
    var v59 uint = (^(v4))
    var v60 uint8 = 210
    if (v3 != (v3 ^ v3)) {
        v61 := GenS4{F0: v3, F1: 32767}
        var v62 int32 = -1
        if (691 < 568) {
            v63 := v6["\n\t"]
            v64 := []int{}
            _ = v63
            _ = v64
        }
        _ = v61
        _ = v62
    } else {
        if false {
            for fi65 := 0; fi65 < 7; fi65++ {
                v66 := GenS6{F0: 127, F1: (((v4 < (v4 + 0)) || true) || (!(65535 > v4))), F2: (v59 & 2)}
                if (9223372036854775807 == (v3 >> (((v4 & 2)) & 63))) {
                    fmt.Println(v60, v1, fi65)
                    if false {
                        if false {
                            v67 := GenS6{F0: fi65, F1: (!(-135 != genFn0("a", 818))), F2: 65535}
                            var v68 int64 = -818
                            var v69 int8 = 1
                            _ = v67
                            _ = v68
                            _ = v69
                        }
                        var v70 uint32 = 898
                        v71 := GenS7{F0: false, F1: "\n\t", F2: 1}
                        _ = v70
                        _ = v71
                    } else {
                        switch "\n\t" {
                        case "a":
                            _ = len(v2)
                        case "foo bar":
                            var v72 int16 = 194
                            _ = v72
                        case "\n\t":
                            switch "a" {
                            case "\n\t":
                                v1 -= v1
                            case "a":
                                v1 += 0
                            case "":
                                switch "hello" {
                                case "\n\t":
                                    switch (fi65 % ((1) | 1)) {
                                    case -748:
                                        delete(v6, "foo bar")
                                    case 32767:
                                        if true {
                                            if true {
                                                if true {
                                                    for fi73 := 0; fi73 < 5; fi73++ {
                                                        var v74 int32 = 385
                                                        switch 313 {
                                                        case 326:
                                                            delete(v5, 262)
                                                        case -9223372036854775808:
                                                            v60 = v60
                                                        case 513:
                                                            fmt.Println(v3, fi73, v4)
                                                        default:
                                                            _ = len(v6)
                                                        }
                                                        _ = fi73
                                                        _ = v74
                                                    }
                                                    for fi75 := 0; fi75 < 15; fi75++ {
                                                        v1 = (v1 + int8(v59))
                                                        delete(v5, 794)
                                                        _ = fi75
                                                    }
                                                } else {
                                                    v4 += (v4 >> ((v59) & 63))
                                                }
                                                if false {
                                                    v76 := []int{4294967295, fi65, fi65, fi65}
                                                    _ = v76
                                                } else {
                                                    if true {
                                                        var v77 uint32 = 2147483647
                                                        _ = v77
                                                    } else {
                                                        if (42 == (v4 | (^(v59)))) {
                                                            if false {
                                                                for fi78 := 0; fi78 < 8; fi78++ {
                                                                    var v79 int64 = v3
                                                                    v80 := []int{(fi78 * -1)}
                                                                    _ = fi78
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                                if (4294967295 < 1) {
                                                                    for fi81 := 0; fi81 < 17; fi81++ {
                                                                        if (false || true) {
                                                                            v1 -= int8(v59)
                                                                            v6["a"] = fi65
                                                                        }
                                                                        _ = fi81
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(fi65, fi65)
                                                        }
                                                        if false {
                                                            v82 := []int{127, genFn0("", 537)}
                                                            v2[""] = true
                                                            if true {
                                                                if (false || false) {
                                                                    if len(v82) > 0 {
                                                                        v82[0] = (fi65 * -774)
                                                                    }
                                                                    _ = len(v82)
                                                                }
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    fmt.Println(v4, v1)
                                                                case "a":
                                                                    var v83 int32 = -1
                                                                    _ = v83
                                                                case "foo bar":
                                                                    for fi84 := 0; fi84 < 14; fi84++ {
                                                                        if len(v82) >= 2 {
                                                                            v82 = v82[1:len(v82)]
                                                                        }
                                                                        fmt.Println(v59, v4, v3)
                                                                        _ = fi84
                                                                    }
                                                                default:
                                                                    v3 += (v3 | (v3 * 42))
                                                                }
                                                                if ("hello" >= "a") {
                                                                    _ = len(v82)
                                                                }
                                                            } else {
                                                                switch 1 {
                                                                case 894:
                                                                    var v85 float32 = float32(3.14159)
                                                                    _ = v85
                                                                }
                                                                for fi86 := 0; fi86 < 13; fi86++ {
                                                                    _ = len(v82)
                                                                    v82 = append(v82, fi86)
                                                                    v4 = (^(v4))
                                                                    delete(v2, ("" + "\n\t"))
                                                                    _ = fi86
                                                                }
                                                                delete(v2, "\n\t")
                                                            }
                                                            _ = v82
                                                        } else {
                                                            v87 := GenS3{F0: 237, F1: 2}
                                                            _ = v87
                                                        }
                                                        for fi88 := 0; fi88 < 11; fi88++ {
                                                            if false {
                                                                _ = len(v2)
                                                                v1 = 72
                                                            }
                                                            fmt.Println(v4, fi88, v3)
                                                            v89 := []int{}
                                                            _ = fi88
                                                            _ = v89
                                                        }
                                                    }
                                                    if false {
                                                        v90 := GenS8{F0: (fi65 * (fi65 / ((fi65) | 1))), F1: "hello"}
                                                        var v91 int32 = 1
                                                        switch (fi65 | (^(fi65))) {
                                                        case 0:
                                                            v3 += -144
                                                        default:
                                                            var v92 int32 = 1
                                                            _ = v92
                                                        }
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                    fmt.Println(v3, v59)
                                                }
                                            }
                                            for fi93 := 0; fi93 < 14; fi93++ {
                                                v94 := GenS5{F0: fi93, F1: 184}
                                                _ = fi93
                                                _ = v94
                                            }
                                            if true {
                                                if (1000 < 1) {
                                                    v1 = v1
                                                    for fi95 := 0; fi95 < 10; fi95++ {
                                                        switch 32768 {
                                                        case 9223372036854775807:
                                                            switch (("a" + (("foo bar" + "a") + "")) + "\n\t") {
                                                            case "hello":
                                                                v3 = v3
                                                            case "foo bar":
                                                                if ((^(v59)) <= v59) {
                                                                    delete(v6, (("" + "\n\t") + ("\n\t" + ("" + "a"))))
                                                                    var v96 int16 = -32768
                                                                    _ = v96
                                                                } else {
                                                                    if (v3 != (v3 / (((v3 & -1)) | 1))) {
                                                                        if false {
                                                                            var v97 uint8 = v60
                                                                            _ = v97
                                                                        } else {
                                                                            if (v60 > (v60 * (v60 | (v60 + v60)))) {
                                                                                fmt.Println(fi65, v59, v59)
                                                                                switch (fi65 | (fi95 << ((v59) & 63))) {
                                                                                case 128:
                                                                                    v98 := GenS9{F0: 0}
                                                                                    _ = v98
                                                                                case 0:
                                                                                    fmt.Println(v1, v4, v3)
                                                                                case 4294967295:
                                                                                    v99 := GenS10{F0: (v1 / ((v1) | 1)), F1: (int32(fi65) != -2147483648), F2: fi95}
                                                                                    _ = v99
                                                                                default:
                                                                                    var v100 uint16 = 265
                                                                                    _ = v100
                                                                                }
                                                                            } else {
                                                                                if (v60 != v60) {
                                                                                    _ = len(v6)
                                                                                    switch ((("foo bar" + "\n\t") + "") + "foo bar") {
                                                                                    case "foo bar":
                                                                                        var v101 float64 = (0.0 / float64((-2.25 * (3.14159 / 1.0))))
                                                                                        _ = v101
                                                                                    case "a":
                                                                                        if true {
                                                                                            if (!false) {
                                                                                                if (true && false) {
                                                                                                    v102 := GenS11{F0: ((("\n\t" + (((("\n\t" + ("\n\t" + ((("\n\t" + ("" + "foo bar")) + ("" + "a")) + "foo bar"))) + "") + "foo bar") + "")) + "foo bar") + "a"), F1: -808, F2: -2147483648}
                                                                                                    v103 := GenS7{F0: (42 != 1), F1: ("a" + ""), F2: 463}
                                                                                                    _ = v102
                                                                                                    _ = v103
                                                                                                } else {
                                                                                                    switch "foo bar" {
                                                                                                    case "\n\t":
                                                                                                        v104 := []int{fi65, (-(fi95))}
                                                                                                        _ = v104
                                                                                                    default:
                                                                                                        v105 := v6[("hello" + (("" + "hello") + ("hello" + ("foo bar" + "foo bar"))))]
                                                                                                        _ = v105
                                                                                                    }
                                                                                                    v106 := GenS6{F0: 42, F1: (!(!false)), F2: 42}
                                                                                                    var v107 int = (fi65 & int(v3))
                                                                                                    _ = v106
                                                                                                    _ = v107
                                                                                                }
                                                                                                if true {
                                                                                                    switch "hello" {
                                                                                                    case "a":
                                                                                                        v3 -= (v3 % ((v3) | 1))
                                                                                                    default:
                                                                                                        if false {
                                                                                                            var v108 float32 = -2.25
                                                                                                            _ = v108
                                                                                                        } else {
                                                                                                            v109 := []int{fi95, (fi65 << ((v4) & 63)), (fi65 % ((fi65) | 1)), (-(fi65))}
                                                                                                            _ = v109
                                                                                                        }
                                                                                                    }
                                                                                                    v110 := GenS12{F0: 943}
                                                                                                    switch 2147483647 {
                                                                                                    case 461:
                                                                                                        switch ("\n\t" + "\n\t") {
                                                                                                        case "":
                                                                                                            if true {
                                                                                                                fmt.Println(v4, v1)
                                                                                                                fmt.Println(v1, fi95)
                                                                                                                var v111 uint16 = 320
                                                                                                                _ = v111
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            v112 := GenS8{F0: -813, F1: ("hello" + "\n\t")}
                                                                                                            _ = v112
                                                                                                        }
                                                                                                    case -707:
                                                                                                        fmt.Println(v1, fi95, v60)
                                                                                                        v113 := GenS13{F0: v60}
                                                                                                        _ = v113
                                                                                                    case 32768:
                                                                                                        if true {
                                                                                                            v114 := []int{}
                                                                                                            _ = v114
                                                                                                        }
                                                                                                    default:
                                                                                                        var v115 int16 = 304
                                                                                                        _ = v115
                                                                                                    }
                                                                                                    _ = v110
                                                                                                } else {
                                                                                                    v116 := GenS14{F0: 2, F1: (42 >= 1000)}
                                                                                                    fmt.Println(v1)
                                                                                                    _ = v116
                                                                                                }
                                                                                                var v117 float32 = -2.25
                                                                                                _ = v117
                                                                                            } else {
                                                                                                _ = len(v6)
                                                                                                v3 += int64(v59)
                                                                                                var v118 int32 = -1
                                                                                                _ = v118
                                                                                            }
                                                                                            fmt.Println(v60, v4)
                                                                                            switch (((("hello" + "hello") + "hello") + "") + ("hello" + "a")) {
                                                                                            case "foo bar":
                                                                                                if false {
                                                                                                    if true {
                                                                                                        var v119 float32 = -2.25
                                                                                                        fmt.Println(v3)
                                                                                                        _ = v119
                                                                                                    }
                                                                                                } else {
                                                                                                    if false {
                                                                                                        var v120 int16 = 32767
                                                                                                        _ = len(v5)
                                                                                                        _ = v120
                                                                                                    } else {
                                                                                                        var v121 int = (fi95 >> ((v59) & 63))
                                                                                                        _ = v121
                                                                                                    }
                                                                                                    fmt.Println(v3)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v122 := []int{genFn0(("hello" + ("foo bar" + ("" + "hello"))), 4294967295)}
                                                                                        _ = v122
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi65)
                                                                                    v4 = 280
                                                                                    _ = len(v6)
                                                                                }
                                                                            }
                                                                            if true {
                                                                                if (int8(v3) > v1) {
                                                                                    v1 = 42
                                                                                    v1 -= -128
                                                                                    if (v59 < (v59 + v4)) {
                                                                                        switch genFn0("\n\t", 1) {
                                                                                        case 32767:
                                                                                            v123 := GenS15{F0: (v1 >> ((v4) & 63))}
                                                                                            _ = v123
                                                                                        default:
                                                                                            var v124 int64 = 373
                                                                                            _ = v124
                                                                                        }
                                                                                        v125 := v6[""]
                                                                                        _ = v125
                                                                                    } else {
                                                                                        var v126 uint32 = 2
                                                                                        _ = v126
                                                                                    }
                                                                                }
                                                                                var v127 uint32 = uint32(v4)
                                                                                fmt.Println(v3, v1, v1)
                                                                                _ = v127
                                                                            } else {
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    var v128 int16 = -32768
                                                                                    _ = v128
                                                                                case "":
                                                                                    v4 -= 1
                                                                                case "foo bar":
                                                                                    fmt.Println(v60, v59)
                                                                                }
                                                                                v59 = v59
                                                                                switch "" {
                                                                                case "hello":
                                                                                    v4 += 263
                                                                                case "foo bar":
                                                                                    fmt.Println(v1, fi95, fi65)
                                                                                default:
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        v129 := []int{}
                                                                                        _ = len(v129)
                                                                                        _ = v129
                                                                                    case "a":
                                                                                        if (1e10 > 1.5) {
                                                                                            delete(v6, "\n\t")
                                                                                            var v130 int = genFn0(("\n\t" + ("foo bar" + "a")), 350)
                                                                                            v60 = v60
                                                                                            _ = v130
                                                                                        } else {
                                                                                            switch fi95 {
                                                                                            case 864:
                                                                                                v59 += (^(v59))
                                                                                            }
                                                                                            v131 := []int{(fi95 / (((fi95 << (((v4 << ((v4) & 63))) & 63))) | 1)), fi65}
                                                                                            _ = v131
                                                                                        }
                                                                                    case "":
                                                                                        v3 += (v3 * v3)
                                                                                    }
                                                                                }
                                                                            }
                                                                            delete(v6, "a")
                                                                        }
                                                                        if (42 >= 65535) {
                                                                            v132 := GenS9{F0: 364}
                                                                            _ = v132
                                                                        }
                                                                        var v133 int16 = -574
                                                                        _ = v133
                                                                    } else {
                                                                        var v134 int32 = 0
                                                                        v60 -= 1
                                                                        _ = v134
                                                                    }
                                                                    if false {
                                                                        switch 127 {
                                                                        case 788:
                                                                            var v135 uint64 = 251
                                                                            _ = v135
                                                                        default:
                                                                            switch genFn0("\n\t", 4294967295) {
                                                                            case 128:
                                                                                v60 -= (v60 + 255)
                                                                                v59 -= (^(v4))
                                                                                if false {
                                                                                    v136 := GenS16{F0: "a", F1: "hello"}
                                                                                    _ = v136
                                                                                } else {
                                                                                    fmt.Println(v4, v59, v60)
                                                                                }
                                                                            case 127:
                                                                                fmt.Println(v60)
                                                                            case -387:
                                                                                var v137 bool = (255 <= 820)
                                                                                _ = v137
                                                                            }
                                                                        }
                                                                        var v138 uint32 = 1
                                                                        var v139 uint64 = genFn1()
                                                                        _ = v138
                                                                        _ = v139
                                                                    } else {
                                                                        delete(v5, 978)
                                                                        switch "a" {
                                                                        case "":
                                                                            var v140 uint16 = 32767
                                                                            _ = v140
                                                                        default:
                                                                            fmt.Println(v3, v3, fi65)
                                                                        }
                                                                        v4 = v59
                                                                    }
                                                                    var v141 uint = v4
                                                                    _ = v141
                                                                }
                                                            }
                                                        }
                                                        var v142 int = int(v3)
                                                        _ = fi95
                                                        _ = v142
                                                    }
                                                }
                                            }
                                        }
                                    case -599:
                                        var v143 uint = v4
                                        _ = v143
                                    default:
                                        v4 += (v4 ^ (v4 & (v59 % (((v59 % ((339) | 1))) | 1))))
                                    }
                                case "":
                                    for fi144 := 0; fi144 < 17; fi144++ {
                                        v145 := GenS1{F0: (v4 / (((v4 * v4)) | 1)), F1: (v3 & v3)}
                                        _ = fi144
                                        _ = v145
                                    }
                                default:
                                    v1 += v1
                                }
                            default:
                                fmt.Println(v59, fi65, v59)
                            }
                        }
                        v146 := []int{fi65, genFn0(((("hello" + "a") + "\n\t") + "a"), 19)}
                        delete(v5, 65535)
                        _ = v146
                    }
                } else {
                    fmt.Println(v3)
                    var v147 int16 = 1
                    v59 -= (v59 ^ (^(v4)))
                    _ = v147
                }
                switch fi65 {
                case 0:
                    for fi148 := 0; fi148 < 1; fi148++ {
                        if true {
                            v60 += v60
                            delete(v5, 0)
                            v149 := GenS9{F0: 1}
                            fmt.Println(fi148, v60, v4)
                            delete(v6, ("\n\t" + ""))
                            v4 = 1
                            var v150 uint16 = 453
                            if false {
                                switch "" {
                                case "a":
                                    v151 := GenS5{F0: (fi148 >> ((1) & 63)), F1: 42}
                                    _ = v151
                                default:
                                    if false {
                                        fmt.Println(v59, v1, v1)
                                    }
                                }
                            }
                            var v152 int16 = 1
                            _ = v149
                            _ = v150
                            _ = v152
                        } else {
                            var v153 string = ""
                            var v154 float64 = (1e10 - 1.5)
                            _ = v153
                            _ = v154
                        }
                        var v155 int16 = -277
                        var v156 uint16 = 371
                        _ = fi148
                        _ = v155
                        _ = v156
                    }
                case 777:
                    switch fi65 {
                    case -412:
                        var v157 uint32 = 2147483648
                        _ = v157
                    case 2147483647:
                        if false {
                            var v158 int8 = v1
                            if ((false || true) || true) {
                                fmt.Println(v3, fi65, v3)
                                if false {
                                    v158 -= int8(v60)
                                    for fi159 := 0; fi159 < 10; fi159++ {
                                        v59 += v4
                                        fmt.Println(v1)
                                        _ = fi159
                                    }
                                    fmt.Println(v60, v158, v4)
                                } else {
                                    v160 := GenS7{F0: false, F1: "", F2: 0}
                                    _ = v160
                                }
                            } else {
                                v59 -= v59
                                v158 += -90
                                if false {
                                    if (!((^(v4)) >= (v59 >> (((v59 & 86)) & 63)))) {
                                        var v161 uint16 = 2
                                        _ = v161
                                    } else {
                                        v162 := v6["a"]
                                        fmt.Println(v59, v158, v1)
                                        _ = v162
                                    }
                                    v163 := GenS1{F0: (v59 / ((319) | 1)), F1: (v3 >> ((1) & 63))}
                                    _ = v163
                                }
                            }
                            _ = v158
                        }
                    case 128:
                        v164 := v2[("a" + "\n\t")]
                        _ = v164
                    }
                default:
                    delete(v6, "a")
                }
                _ = fi65
                _ = v66
            }
            switch ((("" + "\n\t") + "") + "foo bar") {
            case "a":
                if true {
                    if (false || false) {
                        for fi165 := 0; fi165 < 6; fi165++ {
                            switch "foo bar" {
                            case "a":
                                fmt.Println(v60, v3)
                            case "foo bar":
                                if (v3 <= v3) {
                                    v1 = (^(v1))
                                    var v166 int = genFn0("foo bar", 830)
                                    delete(v2, "foo bar")
                                    _ = v166
                                } else {
                                    for fi167 := 0; fi167 < 19; fi167++ {
                                        if ((v1 ^ v1) < 42) {
                                            var v168 int = 32767
                                            fmt.Println(v59)
                                            _ = v168
                                        }
                                        v169 := []int{fi167, fi165, int(v59)}
                                        _ = fi167
                                        _ = v169
                                    }
                                    v60 += 65
                                }
                            case "hello":
                                if ((true || (true || true)) && (128 <= (^(v60)))) {
                                    var v170 int64 = v3
                                    switch fi165 {
                                    case 9223372036854775807:
                                        if ((((1e10 / (float32(-2.25) / (1e10 / float32(((0.1 * (-2.25 / ((float64(1e10) - 0.1) - float64(((-2.25 - 3.14159) + -2.25))))) * ((((1.5 / (3.14159 / 0.1)) + ((1.5 * (float64(-2.25) * 1e10)) + ((0.1 + (-2.25 + 1.5)) / 1e10))) * 1.5) + 0.1)))))) * 1e10) + 0.0) > float32(0.0)) {
                                            var v171 uint8 = 127
                                            var v172 uint64 = 633
                                            if ((!false) && true) {
                                                v172 += (^(v172))
                                                var v173 uint64 = v172
                                                v174 := GenS14{F0: 926, F1: (v172 > genFn1())}
                                                _ = v173
                                                _ = v174
                                            } else {
                                                for fi175 := 0; fi175 < 18; fi175++ {
                                                    var v176 string = ""
                                                    if (false || true) {
                                                        var v177 float64 = 0.1
                                                        _ = v177
                                                    }
                                                    v59 = 4294967295
                                                    _ = fi175
                                                    _ = v176
                                                }
                                                var v178 float32 = 3.14159
                                                _ = v178
                                            }
                                            _ = v171
                                            _ = v172
                                        }
                                    }
                                    _ = v170
                                } else {
                                    if ((1.5 >= (1.5 - 0.1)) || true) {
                                        v6[""] = 437
                                        fmt.Println(v59, v1, fi165)
                                        var v179 float64 = 1.5
                                        _ = v179
                                    } else {
                                        v180 := GenS12{F0: 65535}
                                        _ = v180
                                    }
                                    v3 = 42
                                }
                            default:
                                v181 := []int{genFn0(("a" + (("foo bar" + (("a" + "\n\t") + "\n\t")) + "foo bar")), 0), 472, genFn0(("foo bar" + "foo bar"), 0), -871}
                                _ = v181
                            }
                            fmt.Println(v60)
                            _ = fi165
                        }
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            if false {
                                var v182 int64 = (v3 - (v3 * (^(v3))))
                                var v183 int32 = 0
                                for fi184 := 0; fi184 < 11; fi184++ {
                                    for fi185 := 0; fi185 < 13; fi185++ {
                                        switch (fi184 & int(v60)) {
                                        case 127:
                                            v186 := GenS3{F0: 42, F1: 1}
                                            _ = v186
                                        case -2147483648:
                                            switch (-(fi184)) {
                                            case 127:
                                                delete(v6, "a")
                                            case -9223372036854775808:
                                                v187 := GenS17{F0: -1}
                                                _ = v187
                                            default:
                                                v59 += 105
                                            }
                                        default:
                                            if false {
                                                fmt.Println(v59, v182)
                                                fmt.Println(fi184, v59, v60)
                                                v188 := GenS13{F0: 0}
                                                _ = v188
                                            }
                                        }
                                        delete(v5, 2147483647)
                                        v189 := GenS18{F0: (fi184 + fi185), F1: (float64(float32((float64(1e10) - ((float64(-2.25) + 0.1) + 3.14159)))) >= -2.25), F2: v4}
                                        var v190 string = "\n\t"
                                        fmt.Println(v190, fi184)
                                        var v191 uint8 = 2
                                        _ = fi185
                                        _ = v189
                                        _ = v190
                                        _ = v191
                                    }
                                    var v192 uint64 = 42
                                    _ = fi184
                                    _ = v192
                                }
                                _ = v182
                                _ = v183
                            } else {
                                v3 = v3
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v59)
                                case "hello":
                                    v59 -= v4
                                case "":
                                    fmt.Println(v1)
                                }
                                if true {
                                    for fi193 := 0; fi193 < 13; fi193++ {
                                        var v194 float32 = -2.25
                                        switch -1 {
                                        case 2147483647:
                                            v1 -= -128
                                        default:
                                            var v195 int8 = (v1 + (v1 - v1))
                                            _ = v195
                                        }
                                        _ = fi193
                                        _ = v194
                                    }
                                }
                            }
                        case "\n\t":
                            var v196 string = "foo bar"
                            _ = v196
                        }
                        for fi197 := 0; fi197 < 20; fi197++ {
                            if (!(true && false)) {
                                var v198 int64 = v3
                                fmt.Println(fi197, v1, v3)
                                _ = v198
                            }
                            if true {
                                if true {
                                    switch ("\n\t" + "a") {
                                    case "hello":
                                        fmt.Println(v3)
                                    case "":
                                        v199 := []int{fi197}
                                        _ = v199
                                    }
                                    v3 = -1
                                } else {
                                    v200 := GenS8{F0: -621, F1: "a"}
                                    _ = v200
                                }
                                var v201 float64 = float64((float32(-2.25) + ((0.1 + 1.5) * 0.0)))
                                var v202 float64 = ((v201 + 3.14159) - v201)
                                _ = v201
                                _ = v202
                            }
                            var v203 float64 = 0.0
                            _ = fi197
                            _ = v203
                        }
                    }
                } else {
                    v204 := []int{}
                    var v205 int32 = -1
                    v60 += v60
                    _ = v204
                    _ = v205
                }
            case "hello":
                if true {
                    fmt.Println(v60, v3, v3)
                    var v206 uint32 = 2
                    delete(v6, "hello")
                    _ = v206
                } else {
                    v1 -= 42
                    var v207 string = (("" + ("hello" + "\n\t")) + (((("\n\t" + "a") + "hello") + "hello") + ((("a" + "hello") + (((("\n\t" + "\n\t") + "\n\t") + "\n\t") + "")) + ("hello" + "a"))))
                    _ = v207
                }
            case "":
                var v208 int64 = v3
                _ = v208
            default:
                fmt.Println(v1, v1, v3)
            }
        }
        v3 -= int64(v1)
        v209 := GenS10{F0: (v1 >> (((v59 & v59)) & 63)), F1: (0 == 225), F2: 32767}
        _ = v209
    }
    if ("\n\t" <= "\n\t") {
        v210 := GenS19{F0: v60, F1: 267}
        v1 = v1
        var v211 int8 = v1
        _ = v210
        _ = v211
    }
    switch 0 {
    case 128:
        v212 := v5[65535]
        _ = v212
    case 1:
        v4 += 46
    case -56:
        v4 -= v59
    }
    v213 := GenS6{F0: 42, F1: (genFn1() > 1), F2: 2}
    if true {
        var v214 uint = (v59 ^ (v4 * (v4 & 0)))
        var v215 bool = false
        _ = v214
        _ = v215
    }
    v216 := GenS20{F0: "", F1: 42, F2: -32768}
    if true {
        var v217 bool = (!false)
        if v217 {
            v3 = 1
        }
        _ = v217
    } else {
        v218 := GenS14{F0: 18, F1: false}
        _ = v218
    }
    var v219 uint64 = 39
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v219:", v219)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v213
    _ = v216
    _ = v219
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
}
