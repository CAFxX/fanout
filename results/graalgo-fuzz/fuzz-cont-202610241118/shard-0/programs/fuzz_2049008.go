package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 int64; F1 string; F2 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int8; F1 uint64 }
type GenS12 struct { F0 bool; F1 int32 }
type GenS13 struct { F0 int64; F1 uint; F2 int }
type GenS14 struct { F0 uint32; F1 uint16; F2 string }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint }
func genFn0(p0 uint64, p1 uint8) string {
    var v4 bool = (2147483647 < 0)
    v4 = (v4 || (-85 >= -87))
    _ = p0
    _ = p1
    _ = v4
    return ((("foo bar" + "hello") + "") + ("a" + ("hello" + "")))
}

func main() {
    v1 := []int{4294967295, 2147483647, 9223372036854775807, -1}
    v2 := []int{127, 2147483647}
    v3 := make(map[string]bool)
    var v5 float32 = 3.14159
    fmt.Println(v5)
    if len(v1) > 0 {
        v1[0] = 32768
    }
    switch 0 {
    case -2147483648:
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 uint8 = 246
            fmt.Println(fi6, v5, v7)
            _ = fi6
            _ = v7
        }
    case 128:
        v5 = v5
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v5 += float32(0.0)
        _ = fi8
    }
    if ((32767 >= 42) || true) {
        fmt.Println(v5)
        for fi9 := 0; fi9 < 20; fi9++ {
            v3["hello"] = false
            v5 = 1e10
            switch "a" {
            case "":
                var v10 string = "a"
                _ = v10
            default:
                for fi11 := 0; fi11 < 1; fi11++ {
                    switch ("" + "") {
                    case "hello":
                        var v12 bool = (!(1 == 100))
                        _ = v12
                    case "\n\t":
                        if false {
                            if (-915 >= 1) {
                                if true {
                                    if (!true) {
                                        _ = len(v2)
                                        if (!(true || false)) {
                                            v5 += (float32(3.14159) + (float32(-2.25) + v5))
                                        } else {
                                            v5 = (float32(0.1) * 1e10)
                                            fmt.Println(fi9)
                                        }
                                        v3[("" + "a")] = (!(991 >= 65535))
                                    } else {
                                        var v13 int64 = 9223372036854775807
                                        switch (-(fi11)) {
                                        case 127:
                                            v5 = float32((-2.25 + 3.14159))
                                        }
                                        v5 = ((v5 * v5) - v5)
                                        _ = v13
                                    }
                                    switch 32768 {
                                    case 32768:
                                        if (!false) {
                                            if (!(741 >= -656)) {
                                                _ = len(v3)
                                                if false {
                                                    var v14 string = ("hello" + "foo bar")
                                                    _ = v14
                                                } else {
                                                    var v15 uint32 = 502
                                                    if true {
                                                        v16 := GenS1{F0: 225}
                                                        fmt.Println(fi11, v15)
                                                        _ = v16
                                                    } else {
                                                        fmt.Println(v5, fi11, v5)
                                                    }
                                                    v17 := GenS2{F0: -1}
                                                    _ = v15
                                                    _ = v17
                                                }
                                            } else {
                                                switch "a" {
                                                case "":
                                                    v18 := GenS3{F0: false, F1: 255, F2: 0}
                                                    _ = v18
                                                case "foo bar":
                                                    _ = len(v3)
                                                case "a":
                                                    if len(v1) > 0 {
                                                        v1[0] = fi11
                                                    }
                                                default:
                                                    v19 := GenS4{F0: 42}
                                                    _ = v19
                                                }
                                                var v20 uint8 = 247
                                                switch "\n\t" {
                                                case "":
                                                    var v21 float32 = -2.25
                                                    _ = v21
                                                case "hello":
                                                    var v22 string = ""
                                                    _ = v22
                                                case "a":
                                                    v3[("" + "foo bar")] = true
                                                }
                                                _ = v20
                                            }
                                            var v23 uint16 = 2
                                            switch ("\n\t" + "hello") {
                                            case "a":
                                                var v24 int = fi9
                                                _ = v24
                                            }
                                            _ = v23
                                        }
                                    case 2147483647:
                                        var v25 int = -284
                                        delete(v3, "hello")
                                        _ = v25
                                    default:
                                        v5 = 0.1
                                    }
                                    v26 := GenS1{F0: 128}
                                    switch "hello" {
                                    case "\n\t":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(fi9)
                                        var v27 int = (fi9 % ((-2147483648) | 1))
                                        _ = v27
                                    }
                                    _ = v26
                                }
                            }
                            var v28 int32 = int32(fi9)
                            if (27 != 69) {
                                fmt.Println(v28, fi9)
                            }
                            _ = v28
                        } else {
                            if len(v1) > 0 {
                                v1[0] = fi11
                            }
                        }
                    }
                    if ((100 <= -1) || false) {
                        v29 := GenS5{F0: "", F1: 212}
                        _ = v29
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            v5 = (1e10 / -2.25)
                            fmt.Println(fi9)
                            v5 = float32((1.5 + 0.0))
                        }
                    }
                    if (false && true) {
                        switch fi11 {
                        case 128:
                            fmt.Println(v5, fi9, v5)
                        default:
                            fmt.Println(fi11, fi11, v5)
                        }
                        var v30 float64 = float64(float32((0.0 + 0.0)))
                        v30 -= (v30 + v30)
                        _ = v30
                    }
                    _ = fi11
                }
            }
            _ = fi9
        }
    }
    _ = len(v3)
    if true {
        v5 = (((v5 * v5) - float32(3.14159)) * -2.25)
        if true {
            var v31 uint = 744
            var v32 string = ""
            _ = v31
            _ = v32
        } else {
            fmt.Println(v5)
        }
        fmt.Println(v5)
    } else {
        for fi33 := 0; fi33 < 20; fi33++ {
            fmt.Println(v5, fi33)
            _ = fi33
        }
        v5 += (((v5 * 3.14159) / v5) * (v5 - v5))
    }
    for fi34 := 0; fi34 < 3; fi34++ {
        if (false && true) {
            fmt.Println(v5, fi34)
        }
        v35 := GenS1{F0: 128}
        v5 += ((v5 * 0.0) + 0.1)
        _ = fi34
        _ = v35
    }
    switch ("" + ((("" + "") + ("hello" + (("hello" + "foo bar") + "hello"))) + ("a" + "a"))) {
    case "\n\t":
        v36 := v3[("hello" + "foo bar")]
        _ = v36
    case "":
        for fi37 := 0; fi37 < 5; fi37++ {
            if (fi37 >= 32768) {
                if ((((563 != -1) && (("a" + "hello") <= ("hello" + ""))) && (!((((3.14159 + 0.0) + -2.25) + 1.5) < 1.5))) || true) {
                    switch fi37 {
                    case 2147483647:
                        switch (fi37 & fi37) {
                        case 1:
                            if (("foo bar" + "foo bar") >= "hello") {
                                for fi38 := 0; fi38 < 19; fi38++ {
                                    var v39 string = "\n\t"
                                    v39 = "a"
                                    v5 += (v5 - (0.0 + v5))
                                    _ = fi38
                                    _ = v39
                                }
                            }
                        case 183:
                            var v40 uint8 = 180
                            _ = v40
                        case -9223372036854775808:
                            fmt.Println(v5, v5)
                        default:
                            delete(v3, "")
                        }
                    case 1:
                        v5 += v5
                    case -1:
                        v2 = append(v2, -358)
                    default:
                        v41 := GenS6{F0: -9223372036854775808, F1: 32767}
                        _ = v41
                    }
                    var v42 uint16 = 65535
                    if true {
                        v42 = 427
                    } else {
                        var v43 uint8 = 127
                        _ = v43
                    }
                    _ = v42
                }
            } else {
                if (!true) {
                    v5 += 0.1
                    v44 := GenS5{F0: ("a" + ("hello" + "\n\t")), F1: 32767}
                    switch "hello" {
                    case "foo bar":
                        if true {
                            for fi45 := 0; fi45 < 5; fi45++ {
                                if false {
                                    v5 -= (float32((1.5 / (3.14159 - 1e10))) + (-2.25 / v5))
                                } else {
                                    fmt.Println(v5, fi45)
                                    var v46 uint = 4294967295
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(fi37, v46, v5)
                                    case "":
                                        v46 = (v46 / ((v46) | 1))
                                    case "\n\t":
                                        v47 := GenS7{F0: -9223372036854775808, F1: "\n\t", F2: (fi45 & (fi45 ^ fi37))}
                                        _ = v47
                                    }
                                    _ = v46
                                }
                                _ = fi45
                            }
                        } else {
                            v48 := GenS5{F0: "\n\t", F1: -669}
                            var v49 string = "foo bar"
                            _ = v48
                            _ = v49
                        }
                    default:
                        var v50 int64 = 1
                        _ = v50
                    }
                    _ = v44
                } else {
                    switch "" {
                    case "hello":
                        if (0 == fi37) {
                            v51 := GenS8{F0: (1 == 2)}
                            for fi52 := 0; fi52 < 9; fi52++ {
                                delete(v3, ("a" + "foo bar"))
                                var v53 int = 4294967295
                                switch "a" {
                                case "":
                                    var v54 uint32 = 42
                                    _ = v54
                                case "\n\t":
                                    var v55 int = 1
                                    _ = v55
                                case "a":
                                    var v56 uint8 = 1
                                    _ = v56
                                }
                                _ = fi52
                                _ = v53
                            }
                            v57 := GenS9{F0: -128, F1: 0}
                            _ = v51
                            _ = v57
                        }
                    case "a":
                        v5 -= 3.14159
                    case "\n\t":
                        if (39 < 128) {
                            if (!(127 >= fi37)) {
                                if true {
                                    switch (fi37 * fi37) {
                                    case 32767:
                                        var v58 bool = true
                                        _ = v58
                                    case 164:
                                        if false {
                                            v59 := GenS10{F0: (true && true)}
                                            _ = v59
                                        } else {
                                            v60 := GenS11{F0: -128, F1: 729}
                                            if true {
                                                _ = len(v3)
                                                for fi61 := 0; fi61 < 18; fi61++ {
                                                    v5 = v5
                                                    if (!(false && true)) {
                                                        if false {
                                                            switch 4294967295 {
                                                            case -2147483648:
                                                                switch fi61 {
                                                                case 4294967295:
                                                                    fmt.Println(fi61, v5, fi37)
                                                                default:
                                                                    if (!false) {
                                                                        v5 -= float32(-2.25)
                                                                        v62 := v3["a"]
                                                                        _ = v62
                                                                    } else {
                                                                        v5 = float32((-2.25 + 3.14159))
                                                                    }
                                                                }
                                                            case 0:
                                                                if false {
                                                                    var v63 uint16 = 42
                                                                    _ = v63
                                                                } else {
                                                                    v64 := GenS12{F0: true, F1: 820}
                                                                    _ = v64
                                                                }
                                                            default:
                                                                switch (fi61 + 32767) {
                                                                case 42:
                                                                    v5 = v5
                                                                case 4294967295:
                                                                    v5 = v5
                                                                }
                                                            }
                                                        } else {
                                                            var v65 float32 = 1e10
                                                            fmt.Println(fi61, fi37, v5)
                                                            v5 -= (-2.25 * 1e10)
                                                            _ = v65
                                                        }
                                                        if (42 == 42) {
                                                            v5 -= 0.1
                                                        }
                                                    } else {
                                                        v66 := GenS13{F0: 9223372036854775807, F1: 311, F2: (^(fi37))}
                                                        fmt.Println(fi61)
                                                        _ = v66
                                                    }
                                                    v5 = v5
                                                    _ = fi61
                                                }
                                                switch -2147483648 {
                                                case 32768:
                                                    for fi67 := 0; fi67 < 11; fi67++ {
                                                        v68 := GenS13{F0: -1, F1: 18446744073709551615, F2: fi67}
                                                        v5 = ((1.5 * ((v5 + 3.14159) + ((v5 - 1e10) - (0.1 * 0.1)))) * 3.14159)
                                                        _ = fi67
                                                        _ = v68
                                                    }
                                                case 706:
                                                    switch (fi37 ^ -9223372036854775808) {
                                                    case 188:
                                                        switch 949 {
                                                        case -2147483648:
                                                            var v69 uint32 = 65535
                                                            _ = v69
                                                        case -375:
                                                            var v70 int64 = 465
                                                            _ = v70
                                                        case 127:
                                                            switch fi37 {
                                                            case 141:
                                                                v71 := GenS9{F0: 0, F1: -727}
                                                                _ = v71
                                                            case -402:
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                        default:
                                                            v72 := GenS14{F0: 852, F1: 451, F2: "hello"}
                                                            _ = v72
                                                        }
                                                    case 128:
                                                        var v73 float32 = 1e10
                                                        _ = v73
                                                    case -1:
                                                        var v74 int = fi37
                                                        _ = v74
                                                    default:
                                                        v5 -= (v5 - v5)
                                                    }
                                                default:
                                                    v5 = (float32(1e10) * (float32(3.14159) - v5))
                                                }
                                            }
                                            _ = v60
                                        }
                                    default:
                                        var v75 uint8 = 103
                                        _ = v75
                                    }
                                }
                            }
                            v5 -= v5
                        } else {
                            if (((-2.25 - float64(1e10)) != (1.5 + (float64(v5) * 0.0))) || true) {
                                v76 := GenS6{F0: -928, F1: 2}
                                if true {
                                    fmt.Println(fi37)
                                    var v77 uint64 = 4294967295
                                    _ = v77
                                } else {
                                    switch (("\n\t" + ("foo bar" + "hello")) + ("" + "hello")) {
                                    case "a":
                                        switch "\n\t" {
                                        case "":
                                            v5 = (1e10 * 1e10)
                                        case "\n\t":
                                            for fi78 := 0; fi78 < 15; fi78++ {
                                                switch ("" + ("" + ((("foo bar" + "a") + "foo bar") + "foo bar"))) {
                                                case "":
                                                    v1 = append(v1, (fi78 * (fi78 * -1)))
                                                }
                                                var v79 int = 95
                                                switch "" {
                                                case "":
                                                    v80 := v3[""]
                                                    _ = v80
                                                case "foo bar":
                                                    fmt.Println(fi37, v5, fi37)
                                                default:
                                                    switch "" {
                                                    case "\n\t":
                                                        var v81 uint16 = 32767
                                                        _ = v81
                                                    case "a":
                                                        fmt.Println(fi37, fi78)
                                                    default:
                                                        v82 := GenS15{F0: "hello"}
                                                        _ = v82
                                                    }
                                                }
                                                var v83 int = v79
                                                switch "hello" {
                                                case "\n\t":
                                                    switch "foo bar" {
                                                    case "hello":
                                                        fmt.Println(v79, fi37, fi37)
                                                    case "\n\t":
                                                        var v84 float64 = -2.25
                                                        _ = v84
                                                    case "foo bar":
                                                        if len(v2) > 0 {
                                                            v2[0] = v83
                                                        }
                                                        var v85 float64 = (-2.25 * (3.14159 / 0.1))
                                                        var v86 uint16 = 42
                                                        fmt.Println(v86, v86)
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                case "foo bar":
                                                    v83 = (v79 ^ fi78)
                                                case "":
                                                    switch (fi37 - (v83 & 42)) {
                                                    case 32767:
                                                        delete(v3, ("hello" + ("foo bar" + "a")))
                                                    case 42:
                                                        var v87 uint = 2
                                                        _ = v87
                                                    default:
                                                        v5 = (v5 - v5)
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v5 = ((float32(-2.25) - (1.5 / float32(-2.25))) / v5)
                                                        case "a":
                                                            v88 := GenS11{F0: 127, F1: 668}
                                                            _ = v88
                                                        case "\n\t":
                                                            if false {
                                                                v89 := GenS16{F0: 255}
                                                                _ = v89
                                                            }
                                                        }
                                                        v90 := GenS10{F0: (127 <= 68)}
                                                        _ = v90
                                                    }
                                                }
                                                fmt.Println(fi37, v83, v83)
                                                switch ("hello" + "\n\t") {
                                                case "foo bar":
                                                    fmt.Println(v79, fi78)
                                                case "a":
                                                    v79 -= v83
                                                case "hello":
                                                    var v91 uint = 2
                                                    _ = v91
                                                default:
                                                    var v92 uint8 = 1
                                                    _ = v92
                                                }
                                                _ = fi78
                                                _ = v79
                                                _ = v83
                                            }
                                        default:
                                            var v93 bool = (537 == 1)
                                            _ = v93
                                        }
                                    }
                                }
                                _ = v76
                            }
                        }
                    default:
                        v94 := GenS15{F0: "foo bar"}
                        _ = v94
                    }
                    switch (fi37 | 9223372036854775807) {
                    case 708:
                        _ = len(v1)
                    case 0:
                        fmt.Println(fi37)
                    default:
                        var v95 int = fi37
                        _ = v95
                    }
                }
                v5 = 1.5
                fmt.Println(fi37)
            }
            var v96 string = "\n\t"
            _ = len(v3)
            _ = fi37
            _ = v96
        }
    default:
        v5 = ((v5 + 0.1) / ((v5 / v5) * (v5 / 1.5)))
    }
    fmt.Println(v5)
    v5 += float32(1.5)
    switch 2147483647 {
    case 873:
        fmt.Println(v5)
    case 32767:
        v97 := GenS14{F0: 1, F1: 42, F2: ""}
        _ = v97
    case 0:
        for fi98 := 0; fi98 < 10; fi98++ {
            switch "" {
            case "":
                if false {
                    v5 = (v5 * float32(0.1))
                    fmt.Println(fi98, fi98)
                }
            default:
                v5 = v5
            }
            v99 := GenS6{F0: -9223372036854775808, F1: 876}
            _ = fi98
            _ = v99
        }
    }
    switch 1 {
    case 359:
        if (!true) {
            v100 := GenS16{F0: 18446744073709551615}
            _ = v100
        }
    case 4294967295:
        v5 += 0.0
    }
    v5 += v5
    var v101 uint16 = 65535
    if (!true) {
        var v102 uint16 = 27
        var v103 int32 = 842
        _ = v102
        _ = v103
    } else {
        var v104 int32 = 2147483647
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            var v105 uint8 = 164
            var v106 float32 = -2.25
            _ = v105
            _ = v106
        }
        _ = v104
    }
    var v107 int64 = 241
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v101:", v101)
    fmt.Println("v107:", v107)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v101
    _ = v107
    _ = v2
    _ = v3
    _ = v5
}
