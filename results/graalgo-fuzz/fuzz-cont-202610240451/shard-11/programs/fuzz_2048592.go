package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 int32; F1 uint16; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 bool; F2 uint8 }
type GenS6 struct { F0 bool; F1 string; F2 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int32; F1 int64; F2 uint16 }
type GenS11 struct { F0 uint16; F1 uint64 }
func genFn0() int32 {
    v5 := GenS1{F0: false, F1: 32767, F2: "\n\t"}
    var v6 int = 883
    var v7 uint32 = 4294967295
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    v3 := []int{734}
    var v4 uint8 = 2
    var v8 uint64 = 1
    if v2 {
        v1[""] = "\n\t"
        if (v2 || (v2 || (42 == 100))) {
            var v9 uint32 = 65535
            _ = len(v1)
            _ = v9
        }
        var v10 uint64 = (^(v8))
        _ = v10
    } else {
        v11 := GenS2{F0: 42, F1: 65535, F2: 1}
        _ = v11
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        if len(v3) > 0 {
            v3[0] = -248
        }
    }
    if (1 < 614) {
        var v12 int16 = 32767
        if v2 {
            v13 := GenS3{F0: ("\n\t" + "\n\t")}
            _ = v13
        } else {
            for fi14 := 0; fi14 < 15; fi14++ {
                for fi15 := 0; fi15 < 6; fi15++ {
                    var v16 float32 = float32((0.1 / (0.0 + ((-2.25 * float64(float32((3.14159 + 3.14159)))) + -2.25))))
                    var v17 int8 = -128
                    if true {
                        v17 = v17
                        fmt.Println(v12, v2, v17)
                    } else {
                        v12 -= v12
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                if true {
                    var v18 uint = 2
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v3 = append(v3, (fi19 << ((v18) & 63)))
                        v20 := GenS3{F0: ("hello" + ((("" + "") + "\n\t") + ""))}
                        _ = fi19
                        _ = v20
                    }
                    v4 += 42
                    _ = v18
                }
                fmt.Println(v12, v4, v12)
                _ = fi14
            }
            if v2 {
                var v21 float64 = (((float64((float32(3.14159) * 0.1)) * float64(0.1)) / 1.5) * (1.5 + 0.0))
                if v2 {
                    v21 += (v21 - v21)
                    if ((v21 * v21) > (v21 + (v21 + float64((3.14159 + -2.25))))) {
                        if v2 {
                            if (v21 != v21) {
                                v22 := GenS2{F0: -1, F1: 0, F2: 2147483647}
                                for fi23 := 0; fi23 < 17; fi23++ {
                                    var v24 uint = 0
                                    v25 := GenS4{F0: 18446744073709551615}
                                    for fi26 := 0; fi26 < 18; fi26++ {
                                        v27 := GenS2{F0: -1, F1: 65535, F2: fi26}
                                        _ = fi26
                                        _ = v27
                                    }
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                                _ = v22
                            }
                            for fi28 := 0; fi28 < 16; fi28++ {
                                var v29 int16 = -1
                                for fi30 := 0; fi30 < 3; fi30++ {
                                    v12 = v29
                                    _ = fi30
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi28
                                _ = v29
                            }
                            var v31 uint64 = (v8 ^ v8)
                            _ = v31
                        }
                        switch "a" {
                        case "a":
                            for fi32 := 0; fi32 < 3; fi32++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v33 := GenS4{F0: uint64(v4)}
                                _ = fi32
                                _ = v33
                            }
                        case "hello":
                            switch "\n\t" {
                            case "foo bar":
                                v8 += 975
                            case "":
                                v34 := GenS4{F0: v8}
                                _ = v34
                            case "hello":
                                fmt.Println(v2)
                            default:
                                fmt.Println(v4)
                            }
                        }
                        if len(v3) > 0 {
                            v3[0] = -679
                        }
                    }
                } else {
                    for fi35 := 0; fi35 < 4; fi35++ {
                        var v36 int8 = 1
                        _ = fi35
                        _ = v36
                    }
                    switch ("\n\t" + ("" + "a")) {
                    case "hello":
                        switch ("" + "a") {
                        case "hello":
                            switch -9223372036854775808 {
                            case 32768:
                                if len(v3) > 0 {
                                    v3[0] = -415
                                }
                            case 42:
                                v4 += 255
                            case 2147483647:
                                if (-128 == 1) {
                                    for fi37 := 0; fi37 < 20; fi37++ {
                                        for fi38 := 0; fi38 < 1; fi38++ {
                                            delete(v1, (("a" + "\n\t") + "\n\t"))
                                            var v39 int32 = 42
                                            _ = fi38
                                            _ = v39
                                        }
                                        _ = fi37
                                    }
                                    fmt.Println(v12, v4)
                                } else {
                                    v40 := GenS5{F0: 32767, F1: (671 <= 1), F2: v4}
                                    _ = v40
                                }
                            }
                        case "a":
                            switch "\n\t" {
                            case "\n\t":
                                switch 32767 {
                                case 32768:
                                    for fi41 := 0; fi41 < 16; fi41++ {
                                        v2 = (((0.1 + float32((v21 + v21))) / 3.14159) >= float32(-2.25))
                                        switch fi41 {
                                        case 306:
                                            fmt.Println(v2)
                                        }
                                        _ = fi41
                                    }
                                default:
                                    v2 = v2
                                }
                            case "a":
                                v2 = v2
                            default:
                                for fi42 := 0; fi42 < 7; fi42++ {
                                    v8 -= 1
                                    if ((!v2) || true) {
                                        _ = len(v1)
                                    }
                                    if v2 {
                                        var v43 int64 = 23
                                        _ = v43
                                    } else {
                                        if (!("" < "hello")) {
                                            v44 := GenS1{F0: false, F1: 65535, F2: "hello"}
                                            _ = v44
                                        }
                                    }
                                    _ = fi42
                                }
                            }
                        case "\n\t":
                            v45 := v1["foo bar"]
                            _ = v45
                        default:
                            v21 = (0.1 + float64(((1e10 / (1e10 - ((1.5 * (0.0 * (3.14159 + 1e10))) / ((((((0.1 + 1e10) - 1e10) / 1e10) / 1e10) / 1.0) - (0.0 / 3.14159))))) + 1.5)))
                        }
                    case "":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "foo bar":
                        var v46 uint64 = (v8 % ((42) | 1))
                        _ = v46
                    default:
                        v2 = ((v4 % ((v4) | 1)) <= v4)
                    }
                    for fi47 := 0; fi47 < 19; fi47++ {
                        v21 = (1e10 / ((0.1 + (v21 + v21)) / 0.1))
                        _ = fi47
                    }
                }
                v48 := GenS4{F0: v8}
                _ = v21
                _ = v48
            }
            for fi49 := 0; fi49 < 3; fi49++ {
                if true {
                    v4 = 2
                }
                _ = fi49
            }
        }
        if (!((!(false || ((1e10 < float32(0.0)) && (42 < 0)))) && (v2 && true))) {
            var v50 string = "foo bar"
            _ = v50
        }
        _ = v12
    } else {
        var v51 int16 = 32767
        switch 32767 {
        case -461:
            v8 = 575
        default:
            fmt.Println(v8, v2)
        }
        _ = v51
    }
    for fi52 := 0; fi52 < 8; fi52++ {
        if (!true) {
            v4 = v4
            for fi53 := 0; fi53 < 5; fi53++ {
                if v2 {
                    if true {
                        var v54 bool = v2
                        switch "foo bar" {
                        case "a":
                            var v55 uint64 = (^(v8))
                            _ = v55
                        default:
                            v56 := GenS6{F0: (65535 != 65535), F1: "foo bar", F2: (^(fi52))}
                            _ = v56
                        }
                        var v57 uint = 65535
                        _ = v54
                        _ = v57
                    }
                    switch (fi52 * -649) {
                    case 128:
                        v4 += 2
                    case -9223372036854775808:
                        switch fi53 {
                        case 1:
                            if false {
                                fmt.Println(v2, fi53)
                                var v58 int64 = 9223372036854775807
                                switch ("\n\t" + "hello") {
                                case "foo bar":
                                    v59 := v1[("foo bar" + ("" + "\n\t"))]
                                    if (true || (v58 <= 9223372036854775807)) {
                                        _ = len(v3)
                                        v60 := GenS7{F0: (v8 | (v8 ^ 610))}
                                        _ = v60
                                    } else {
                                        if true {
                                            v2 = (v4 > (v4 << ((78) & 63)))
                                            v8 = v8
                                        } else {
                                            fmt.Println(v4)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v58 = 1
                                            v58 += (v58 | v58)
                                            var v61 uint32 = uint32(v4)
                                            _ = v61
                                        }
                                        v8 = 2
                                        v62 := GenS8{F0: 0}
                                        v3 = append(v3, fi52)
                                        v58 -= v58
                                        _ = v62
                                    }
                                    var v63 int32 = 1
                                    _ = v59
                                    _ = v63
                                case "":
                                    v8 -= v8
                                case "a":
                                    delete(v1, "")
                                }
                                _ = v58
                            }
                        case -462:
                            fmt.Println(fi53, v8)
                        default:
                            v4 = (v4 / ((96) | 1))
                            v8 -= (v8 % (((v8 >> ((265) & 63))) | 1))
                            switch -361 {
                            case 143:
                                if v2 {
                                    v64 := GenS9{F0: 9223372036854775807}
                                    var v65 int16 = 294
                                    _ = v64
                                    _ = v65
                                } else {
                                    var v66 int8 = 0
                                    _ = v66
                                }
                            default:
                                var v67 float64 = (0.1 * ((-2.25 / 1.0) - 1e10))
                                _ = v67
                            }
                        }
                    case 1:
                        if false {
                            fmt.Println(v2)
                            v68 := GenS6{F0: v2, F1: "\n\t", F2: fi52}
                            var v69 uint = 4294967295
                            v69 -= v69
                            v69 += (v69 - (v69 + 607))
                            _ = v68
                            _ = v69
                        } else {
                            _ = len(v1)
                            if (!(!(1 != 42))) {
                                v70 := GenS2{F0: -417, F1: 636, F2: (fi52 % (((fi53 + (fi52 & fi53))) | 1))}
                                var v71 uint32 = 1
                                v71 += 0
                                _ = v70
                                _ = v71
                            }
                        }
                    default:
                        _ = len(v3)
                    }
                    var v72 uint64 = (v8 >> ((uint(v4)) & 63))
                    _ = v72
                } else {
                    if (v2 || (521 > 65535)) {
                        if (v2 || true) {
                            if (!(v4 == 229)) {
                                switch (fi52 + fi53) {
                                case -9223372036854775808:
                                    if (42 != 2147483647) {
                                        fmt.Println(fi53)
                                        v2 = v2
                                        var v73 int = fi53
                                        _ = v73
                                    } else {
                                        if (0 >= 42) {
                                            v74 := v1[("foo bar" + "a")]
                                            fmt.Println(v74, fi52, fi52)
                                            var v75 int8 = 100
                                            _ = v74
                                            _ = v75
                                        }
                                    }
                                default:
                                    v2 = ((fi53 >> ((65535) & 63)) == 9223372036854775807)
                                }
                            } else {
                                var v76 uint16 = 1
                                v8 = v8
                                var v77 uint16 = (v76 + 255)
                                _ = v76
                                _ = v77
                            }
                            var v78 int8 = 39
                            _ = v78
                        }
                        var v79 uint = 18446744073709551615
                        v8 = (v8 - v8)
                        _ = v79
                    } else {
                        if (0 != -13) {
                            v8 = (^(v8))
                        } else {
                            var v80 string = "\n\t"
                            _ = v80
                        }
                    }
                    switch "foo bar" {
                    case "\n\t":
                        var v81 uint32 = 943
                        _ = v81
                    case "hello":
                        if true {
                            v82 := GenS5{F0: 42, F1: v2, F2: (^(v4))}
                            if ((("\n\t" + "hello") + "hello") == "\n\t") {
                                if true {
                                    switch fi52 {
                                    case 32767:
                                        v83 := GenS3{F0: ("a" + ("\n\t" + "hello"))}
                                        _ = v83
                                    case 128:
                                        v4 += (v4 ^ (v4 + 128))
                                        v84 := v1["foo bar"]
                                        fmt.Println(v8, v8)
                                        _ = v84
                                    default:
                                        v2 = v2
                                    }
                                    var v85 uint = 42
                                    v4 += 42
                                    var v86 float32 = (float32(0.1) / ((1e10 + 3.14159) * -2.25))
                                    _ = v85
                                    _ = v86
                                }
                                _ = len(v3)
                                v2 = v2
                            }
                            v3 = append(v3, fi52)
                            _ = v82
                        }
                        fmt.Println(v2, fi53, v8)
                    case "":
                        fmt.Println(v2, fi52, v8)
                    default:
                        switch "" {
                        case "a":
                            if len(v3) > 0 {
                                v3[0] = fi53
                            }
                        case "\n\t":
                            var v87 int64 = 42
                            if false {
                                switch ("\n\t" + ("foo bar" + "a")) {
                                case "hello":
                                    v8 = v8
                                case "foo bar":
                                    var v88 uint8 = (^(v4))
                                    _ = v88
                                default:
                                    if v2 {
                                        var v89 uint8 = v4
                                        var v90 float64 = 1e10
                                        if v2 {
                                            if false {
                                                v8 = (v8 - v8)
                                            }
                                            fmt.Println(v90, fi53, fi52)
                                            _ = len(v1)
                                        }
                                        _ = v89
                                        _ = v90
                                    } else {
                                        if false {
                                            switch "\n\t" {
                                            case "a":
                                                _ = len(v1)
                                            case "":
                                                v4 = (v4 * uint8(v8))
                                            case "hello":
                                                v91 := GenS7{F0: (^(v8))}
                                                _ = v91
                                            }
                                            _ = len(v3)
                                            v2 = v2
                                            fmt.Println(v2)
                                            fmt.Println(v87)
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = fi53
                                        }
                                    }
                                }
                                var v92 uint8 = 255
                                v93 := GenS3{F0: "hello"}
                                var v94 uint16 = 255
                                _ = v92
                                _ = v93
                                _ = v94
                            }
                            delete(v1, "\n\t")
                            _ = v87
                        }
                    }
                    v8 -= (v8 ^ (v8 / (((v8 - (v8 ^ v8))) | 1)))
                }
                _ = fi53
            }
        }
        _ = fi52
    }
    fmt.Println(v2)
    if v2 {
        var v95 uint = 1
        v96 := v1[("foo bar" + "a")]
        _ = v95
        _ = v96
    } else {
        var v97 string = "a"
        _ = v97
    }
    v98 := GenS5{F0: -834, F1: false, F2: v4}
    v2 = false
    for fi99 := 0; fi99 < 11; fi99++ {
        for fi100 := 0; fi100 < 4; fi100++ {
            var v101 int32 = -2147483648
            if v2 {
                v2 = (1000 >= 0)
                v1[("a" + "hello")] = "a"
                v2 = true
            } else {
                var v102 int8 = 12
                if false {
                    delete(v1, "\n\t")
                    switch ("foo bar" + "") {
                    case "a":
                        v102 -= (v102 + -1)
                    }
                    switch (("" + "foo bar") + "hello") {
                    case "":
                        v102 += v102
                    case "foo bar":
                        fmt.Println(v4, v8, v4)
                    case "a":
                        switch ("hello" + ("" + "foo bar")) {
                        case "a":
                            if (-2.25 <= 1e10) {
                                v8 = (v8 ^ (v8 & v8))
                            } else {
                                if (3.14159 <= 0.1) {
                                    fmt.Println(v8)
                                    v101 = (^(v101))
                                }
                            }
                        case "\n\t":
                            v2 = false
                        default:
                            var v103 uint = 255
                            _ = v103
                        }
                    default:
                        if (463 > 4294967295) {
                            if false {
                                v4 = v4
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = -9223372036854775808
                                    }
                                    if (0.1 != (0.1 * 1e10)) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    var v104 uint16 = 2
                                    var v105 int64 = -9223372036854775808
                                    var v106 int32 = genFn0()
                                    _ = v104
                                    _ = v105
                                    _ = v106
                                } else {
                                    if false {
                                        if (fi99 >= -190) {
                                            v102 = (-(v102))
                                        }
                                        v8 += v8
                                        fmt.Println(fi99)
                                        var v107 int16 = -870
                                        v2 = false
                                        _ = v107
                                    }
                                    v108 := GenS8{F0: v101}
                                    fmt.Println(fi100, v8, v8)
                                    _ = len(v1)
                                    if len(v3) > 0 {
                                        v3[0] = (fi100 + (fi100 + fi99))
                                    }
                                    _ = v108
                                }
                                switch fi99 {
                                case 32767:
                                    v109 := v1[((("" + "hello") + "a") + "hello")]
                                    _ = v109
                                case 844:
                                    v110 := v1["a"]
                                    _ = v110
                                default:
                                    v111 := GenS10{F0: v101, F1: 564, F2: 1}
                                    _ = v111
                                }
                            }
                            var v112 float64 = (0.1 + 1.5)
                            _ = v112
                        } else {
                            v102 = v102
                        }
                    }
                }
                v8 = (v8 & (v8 ^ (v8 + 42)))
                _ = v102
            }
            v1["a"] = "foo bar"
            v113 := GenS4{F0: v8}
            v114 := GenS4{F0: uint64(v4)}
            switch "foo bar" {
            case "hello":
                v4 += v4
            case "foo bar":
                v115 := GenS11{F0: 255, F1: (v8 / ((242) | 1))}
                _ = v115
            case "\n\t":
                var v116 float32 = 1e10
                _ = v116
            default:
                v8 -= 1
            }
            var v117 float32 = -2.25
            _ = fi100
            _ = v101
            _ = v113
            _ = v114
            _ = v117
        }
        if (!(((v4 << ((1) & 63)) >= 2) && true)) {
            v118 := GenS9{F0: 98}
            _ = v118
        } else {
            switch "hello" {
            case "hello":
                _ = len(v1)
            }
        }
        _ = fi99
    }
    for fi119 := 0; fi119 < 2; fi119++ {
        var v120 int64 = 1
        for fi121 := 0; fi121 < 16; fi121++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch "" {
            case "hello":
                _ = len(v3)
            case "":
                _ = len(v1)
            }
            _ = fi121
        }
        _ = fi119
        _ = v120
    }
    var v122 uint64 = v8
    v123 := GenS4{F0: (v8 - 97)}
    var v124 uint32 = 42
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v122:", v122)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v122
    _ = v123
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v98
}
