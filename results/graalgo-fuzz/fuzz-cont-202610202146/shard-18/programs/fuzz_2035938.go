package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint; F1 bool; F2 string }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64; F1 int; F2 int32 }
type GenS8 struct { F0 bool; F1 bool; F2 uint16 }
type GenS9 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int; F1 uint32; F2 string }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint }
func genFn0(p0 int8, p1 uint16, p2 int) int32 {
    if true {
        var v5 uint = 955
        p1 -= p1
        switch p2 {
        case -649:
            fmt.Println(p2)
        case 981:
            var v6 uint = 1
            _ = v6
        case 273:
            v7 := []int{}
            _ = v7
        }
        _ = v5
    } else {
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(p2, fi8, p2)
            v9 := GenS1{F0: 32767, F1: 1}
            _ = fi8
            _ = v9
        }
        v10 := GenS2{F0: 771, F1: (2147483647 == 944), F2: ("hello" + (("" + "hello") + "hello"))}
        v11 := []int{(p2 << ((65535) & 63)), p2}
        _ = v10
        _ = v11
    }
    var v12 float64 = ((0.1 / (0.0 + 1.5)) / -2.25)
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 float64 = ((v12 / 3.14159) * (v12 * v12))
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -2147483648
}
func genFn1(p0 int, p1 uint) uint16 {
    var v15 uint32 = 0
    switch (p0 / ((42) | 1)) {
    case 2147483647:
        var v16 int64 = 0
        _ = v16
    default:
        if (((-2.25 * 1e10) + 0.1) >= 3.14159) {
            fmt.Println(v15)
            if (("" <= "\n\t") || true) {
                switch (("a" + ((("foo bar" + "hello") + "a") + ("a" + "a"))) + "\n\t") {
                case "":
                    for fi17 := 0; fi17 < 19; fi17++ {
                        if false {
                            p1 += (p1 + 450)
                            p1 -= uint(v15)
                            for fi18 := 0; fi18 < 6; fi18++ {
                                p0 = 180
                                _ = fi18
                            }
                        }
                        fmt.Println(v15, p1)
                        _ = fi17
                    }
                case "foo bar":
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v15, p0)
                    case "":
                        switch "\n\t" {
                        case "foo bar":
                            var v19 int = p0
                            _ = v19
                        case "hello":
                            var v20 uint8 = 127
                            _ = v20
                        case "":
                            p0 = (-(p0))
                        }
                    case "\n\t":
                        v21 := []int{p0, (p0 >> ((p1) & 63)), p0}
                        _ = v21
                    default:
                        var v22 int16 = 1
                        _ = v22
                    }
                case "\n\t":
                    switch (("foo bar" + "") + ("hello" + "hello")) {
                    case "hello":
                        for fi23 := 0; fi23 < 12; fi23++ {
                            fmt.Println(fi23)
                            _ = fi23
                        }
                    case "":
                        v24 := GenS3{F0: 255, F1: 255}
                        _ = v24
                    case "\n\t":
                        var v25 int64 = 9223372036854775807
                        _ = v25
                    default:
                        for fi26 := 0; fi26 < 14; fi26++ {
                            if (false && false) {
                                var v27 int32 = 2147483647
                                _ = v27
                            }
                            fmt.Println(v15)
                            v28 := []int{}
                            _ = fi26
                            _ = v28
                        }
                    }
                default:
                    v29 := GenS1{F0: 42, F1: 1}
                    _ = v29
                }
                if (!(!(((p0 | p0) != p0) || true))) {
                    for fi30 := 0; fi30 < 12; fi30++ {
                        if (974 != 0) {
                            switch "" {
                            case "\n\t":
                                fmt.Println(p1)
                            }
                            var v31 uint64 = 42
                            _ = v31
                        } else {
                            v32 := []int{42}
                            v32 = append(v32, (p0 | 127))
                            _ = v32
                        }
                        var v33 float64 = -2.25
                        fmt.Println(p0, v15, v15)
                        _ = fi30
                        _ = v33
                    }
                } else {
                    var v34 float64 = (0.0 / ((1e10 + (1.5 - -2.25)) / 1e10))
                    fmt.Println(v34, v34, v15)
                    var v35 bool = false
                    _ = v34
                    _ = v35
                }
                var v36 bool = true
                _ = v36
            }
            var v37 float32 = (float32(0.1) * 3.14159)
            _ = v37
        }
    }
    p0 -= (p0 - p0)
    _ = p0
    _ = p1
    _ = v15
    return 42
}

func main() {
    var v1 int8 = -1
    var v2 uint32 = 0
    v3 := []uint32{}
    var v4 int32 = -2147483648
    for fi38 := 0; fi38 < 6; fi38++ {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        _ = fi38
    }
    for fi39 := 0; fi39 < 19; fi39++ {
        var v40 int = 139
        var v41 float32 = 0.1
        _ = fi39
        _ = v40
        _ = v41
    }
    var v42 int16 = 124
    var v43 int8 = (v1 / ((v1) | 1))
    fmt.Println(v43)
    if (2147483647 >= 0) {
        v43 = v1
        var v44 int = -301
        v45 := GenS1{F0: int16(v1), F1: 1}
        _ = v44
        _ = v45
    }
    v43 -= (v1 << ((961) & 63))
    for fi46 := 0; fi46 < 12; fi46++ {
        v47 := GenS2{F0: 255, F1: true, F2: "hello"}
        switch "\n\t" {
        case "\n\t":
            for fi48 := 0; fi48 < 19; fi48++ {
                if (398 >= v42) {
                    switch "a" {
                    case "foo bar":
                        var v49 int64 = -875
                        var v50 uint = 0
                        if (!(!(!true))) {
                            var v51 float32 = -2.25
                            if (v42 <= (v42 & 42)) {
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v52 float32 = (((1.5 / v51) - ((1.5 + float32(-2.25)) * v51)) + ((3.14159 / (float32((((3.14159 * 0.1) + ((((((1.5 * (3.14159 + 1.5)) - 3.14159) - 1e10) * -2.25) / (-2.25 + 0.0)) - 3.14159)) / 3.14159)) - v51)) * v51))
                                        _ = v52
                                    case "foo bar":
                                        switch "a" {
                                        case "":
                                            var v53 int64 = (v49 >> ((v50) & 63))
                                            _ = v53
                                        case "\n\t":
                                            fmt.Println(v51, v42)
                                        case "a":
                                            fmt.Println(v4, v2)
                                        default:
                                            v51 += (float32(-2.25) - (v51 - v51))
                                            _ = len(v3)
                                        }
                                    case "hello":
                                        v54 := GenS3{F0: 138, F1: 2}
                                        _ = v54
                                    }
                                }
                            }
                            v55 := GenS4{F0: 1, F1: v4, F2: 341}
                            _ = v51
                            _ = v55
                        } else {
                            v42 = 1000
                        }
                        v56 := GenS5{F0: int64(v4), F1: 419}
                        var v57 int64 = 0
                        _ = v49
                        _ = v50
                        _ = v56
                        _ = v57
                    default:
                        if (!(2 == 2)) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch ("hello" + ("foo bar" + "")) {
                            case "\n\t":
                                var v58 uint64 = 34
                                _ = v58
                            case "hello":
                                v59 := GenS1{F0: int16(v1), F1: v42}
                                _ = v59
                            case "":
                                v4 = (-(v4))
                            }
                        } else {
                            var v60 uint32 = v2
                            if ((!false) && false) {
                                switch 2147483647 {
                                case 32768:
                                    if true {
                                        v2 = (^(v60))
                                        v1 -= (v1 + (v1 >> ((1) & 63)))
                                        v61 := GenS6{F0: 638}
                                        switch "foo bar" {
                                        case "hello":
                                            switch "a" {
                                            case "\n\t":
                                                v60 -= v2
                                            case "foo bar":
                                                var v62 int8 = v43
                                                _ = v62
                                            }
                                        case "":
                                            var v63 uint64 = 2
                                            _ = v63
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = v2
                                            }
                                        }
                                        switch "" {
                                        case "":
                                            switch "foo bar" {
                                            case "":
                                                if len(v3) > 0 {
                                                    v3[0] = v60
                                                }
                                            case "hello":
                                                if false {
                                                    v64 := GenS6{F0: 2}
                                                    v65 := GenS7{F0: 42, F1: fi48, F2: genFn0((v1 * (-(v1))), 42, fi46)}
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case "":
                                                        v66 := GenS1{F0: v42, F1: 0}
                                                        _ = v66
                                                    case "\n\t":
                                                        switch (fi48 & 32768) {
                                                        case 42:
                                                            var v67 int = (-(fi46))
                                                            _ = v67
                                                        case -103:
                                                            if true {
                                                                v3 = append(v3, 399)
                                                                var v68 uint64 = 18446744073709551615
                                                                _ = v68
                                                            }
                                                        case 730:
                                                            var v69 int64 = 478
                                                            _ = v69
                                                        default:
                                                            switch ((("a" + ("\n\t" + (("hello" + ("a" + "foo bar")) + ((("\n\t" + ((("a" + "") + "foo bar") + ("foo bar" + ""))) + (("hello" + ((("foo bar" + ("" + ((("a" + "a") + "\n\t") + ("foo bar" + "hello")))) + "a") + "hello")) + ("\n\t" + "a"))) + "hello")))) + "foo bar") + "hello") {
                                                            case "a":
                                                                v60 -= v60
                                                            case "\n\t":
                                                                v42 = (v42 / ((v42) | 1))
                                                            case "hello":
                                                                fmt.Println(fi48, v4, fi48)
                                                            default:
                                                                var v70 bool = false
                                                                _ = v70
                                                            }
                                                        }
                                                    }
                                                    if false {
                                                        var v71 float32 = 0.0
                                                        var v72 int32 = (v4 - -1)
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        if true {
                                                            v2 += v2
                                                            v73 := GenS8{F0: true, F1: false, F2: 1}
                                                            var v74 int32 = v4
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                    }
                                                    _ = v64
                                                    _ = v65
                                                }
                                            case "\n\t":
                                                if ("foo bar" != "foo bar") {
                                                    v60 -= (v2 & (v2 >> ((0) & 63)))
                                                }
                                            default:
                                                v1 = (v43 - int8(v60))
                                            }
                                        case "foo bar":
                                            v75 := GenS8{F0: (0.1 < 3.14159), F1: (!false), F2: 0}
                                            _ = v75
                                        default:
                                            fmt.Println(v2, v1)
                                        }
                                        v4 += genFn0(v43, genFn1(fi48, 526), (fi46 >> ((42) & 63)))
                                        _ = v61
                                    }
                                case -9223372036854775808:
                                    fmt.Println(fi48)
                                }
                            }
                            _ = v60
                        }
                    }
                    fmt.Println(fi46, v1)
                } else {
                    v2 -= (^(v2))
                    switch 1 {
                    case 42:
                        if (826 > 235) {
                            if false {
                                if ((true || (((1e10 / 1.5) * 0.0) > -2.25)) || (3.14159 < ((0.1 + 1.5) / 1.0))) {
                                    switch ("" + "\n\t") {
                                    case "":
                                        var v76 uint64 = uint64(v2)
                                        _ = v76
                                    }
                                }
                            } else {
                                if ("hello" > "hello") {
                                    fmt.Println(fi46, fi46)
                                    var v77 float64 = 0.1
                                    switch (fi48 << ((18446744073709551615) & 63)) {
                                    case 4294967295:
                                        v78 := GenS9{F0: 42, F1: -780, F2: 2}
                                        _ = v78
                                    case 417:
                                        var v79 uint64 = 1
                                        _ = v79
                                    default:
                                        _ = len(v3)
                                    }
                                    if (true || (1e10 == 3.14159)) {
                                        var v80 int16 = v42
                                        v77 += (-2.25 / float64((float32((v77 + (v77 - (v77 * v77)))) - 1e10)))
                                        v81 := GenS1{F0: 1, F1: 1}
                                        _ = v80
                                        _ = v81
                                    } else {
                                        v82 := GenS10{F0: 18446744073709551615}
                                        switch (^(fi46)) {
                                        case 32768:
                                            fmt.Println(v77, v2, v43)
                                        case 23:
                                            if false {
                                                _ = len(v3)
                                            }
                                            v1 = (v43 % ((v1) | 1))
                                        case -2147483648:
                                            if false {
                                                fmt.Println(v2, v77)
                                                var v83 int8 = v1
                                                _ = len(v3)
                                                var v84 bool = false
                                                if len(v3) > 0 {
                                                    v3[0] = v2
                                                }
                                                _ = v83
                                                _ = v84
                                            } else {
                                                v3 = append(v3, 820)
                                                var v85 int64 = 9223372036854775807
                                                if true {
                                                    var v86 uint64 = 148
                                                    var v87 uint16 = 42
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v86
                                                    _ = v87
                                                }
                                                _ = v85
                                            }
                                        }
                                        v88 := GenS8{F0: false, F1: false, F2: 1}
                                        _ = v82
                                        _ = v88
                                    }
                                    var v89 int64 = -848
                                    var v90 string = "\n\t"
                                    switch ("hello" + v90) {
                                    case "foo bar":
                                        if false {
                                            if true {
                                                v2 += v2
                                            } else {
                                                var v91 bool = true
                                                fmt.Println(fi48, v90)
                                                v4 += v4
                                                if v91 {
                                                    var v92 uint64 = 54
                                                    v93 := GenS11{F0: fi46, F1: (^(v2)), F2: ""}
                                                    _ = v92
                                                    _ = v93
                                                } else {
                                                    v43 = (v43 + (v1 << ((82) & 63)))
                                                    var v94 uint = 931
                                                    _ = v94
                                                }
                                                if (v77 > -2.25) {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                _ = v91
                                            }
                                            v3 = append(v3, v2)
                                        }
                                    case "a":
                                        v42 = v42
                                        if false {
                                            switch ("hello" + v90) {
                                            case "hello":
                                                if true {
                                                    fmt.Println(fi46, v4, v2)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v95 := GenS12{F0: fi46}
                                                    _ = v95
                                                }
                                            default:
                                                switch fi48 {
                                                case 842:
                                                    if len(v3) > 0 {
                                                        v3[0] = v2
                                                    }
                                                case 2147483647:
                                                    _ = len(v3)
                                                }
                                            }
                                            var v96 float64 = ((float64(0.0) + 1.5) - 0.0)
                                            _ = v96
                                        } else {
                                            fmt.Println(fi46, v90)
                                        }
                                    case "":
                                        v89 = (v89 + (v89 % (((v89 & -1)) | 1)))
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = 986
                                        }
                                    }
                                    _ = v77
                                    _ = v89
                                    _ = v90
                                } else {
                                    var v97 int16 = -1
                                    switch (("" + "foo bar") + "\n\t") {
                                    case "\n\t":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case "hello":
                                        fmt.Println(v43, v2, v2)
                                    }
                                    _ = v97
                                }
                            }
                            v2 += (^(v2))
                        } else {
                            if true {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                v98 := GenS13{F0: 18446744073709551615, F1: 32, F2: genFn0((v1 | 42), 1, (fi48 % ((fi48) | 1)))}
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = 42
                                    }
                                    v2 -= (v2 | v2)
                                    v3 = append(v3, (v2 + 65535))
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                var v99 string = ("a" + "hello")
                                _ = v98
                                _ = v99
                            }
                            var v100 float64 = (float64(0.1) * (1.5 - 1e10))
                            _ = v100
                        }
                    }
                }
                _ = fi48
            }
        case "a":
            _ = len(v3)
        case "foo bar":
            if false {
                fmt.Println(v43, v43)
                v101 := GenS12{F0: (fi46 / (((fi46 >> ((42) & 63))) | 1))}
                _ = v101
            } else {
                var v102 float64 = 3.14159
                var v103 uint64 = 743
                _ = v102
                _ = v103
            }
        }
        switch "foo bar" {
        case "":
            for fi104 := 0; fi104 < 19; fi104++ {
                v1 = v1
                switch fi104 {
                case 128:
                    var v105 uint64 = 0
                    _ = v105
                case 4294967295:
                    if false {
                        v106 := GenS12{F0: 1}
                        switch ("hello" + "foo bar") {
                        case "\n\t":
                            switch (("\n\t" + "a") + ("foo bar" + "\n\t")) {
                            case "a":
                                fmt.Println(v43, fi104, v42)
                            case "foo bar":
                                var v107 int = (^(fi104))
                                _ = v107
                            case "\n\t":
                                var v108 string = (("foo bar" + "hello") + "")
                                _ = v108
                            }
                        case "":
                            fmt.Println(v4)
                        case "foo bar":
                            var v109 float64 = 0.1
                            _ = v109
                        default:
                            var v110 int16 = v42
                            _ = v110
                        }
                        _ = v106
                    }
                    var v111 float32 = 0.0
                    _ = v111
                case 0:
                    v4 = v4
                }
                v4 -= (v4 - genFn0(int8(v4), genFn1(fi46, 4294967295), (fi46 ^ fi104)))
                _ = fi104
            }
        case "a":
            for fi112 := 0; fi112 < 6; fi112++ {
                v3 = append(v3, (^(v2)))
                var v113 string = "\n\t"
                _ = fi112
                _ = v113
            }
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = v2
            }
        default:
            if ((v42 % ((v42) | 1)) <= (v42 | v42)) {
                v42 += v42
            }
        }
        _ = fi46
        _ = v47
    }
    var v114 int8 = -1
    fmt.Println(v42)
    var v115 int16 = -1
    var v116 float32 = float32((0.1 / (1.5 - float64(float32((float64(-2.25) + -2.25))))))
    for fi117 := 0; fi117 < 10; fi117++ {
        v1 = v1
        _ = fi117
    }
    var v118 float64 = 1e10
    if (v114 > int8(v4)) {
        v42 -= v42
        v119 := GenS7{F0: 0, F1: 9223372036854775807, F2: v4}
        v4 -= (^(v4))
        _ = v119
    } else {
        var v120 bool = (!(!true))
        v121 := GenS14{F0: -965}
        _ = v120
        _ = v121
    }
    if len(v3) > 0 {
        v3[0] = v2
    }
    v122 := GenS15{F0: 42}
    fmt.Println(v115, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v116:", v116)
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v114
    _ = v115
    _ = v116
    _ = v118
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}
