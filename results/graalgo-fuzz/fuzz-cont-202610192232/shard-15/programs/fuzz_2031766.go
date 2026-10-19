package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 uint; F2 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool; F1 uint32 }
func genFn0() int8 {
    var v5 uint8 = 42
    v6 := []int{int(v5), 992, 42}
    _ = len(v6)
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{("foo bar" + "hello"), ((((("foo bar" + ("foo bar" + "foo bar")) + "hello") + "a") + "hello") + "a"), ("" + ("hello" + "a"))}
    var v3 int64 = 860
    var v4 int16 = 1
    var v7 int8 = genFn0()
    delete(v1, -2147483648)
    fmt.Println(v4, v3, v3)
    var v8 int16 = v4
    for fi9 := 0; fi9 < 18; fi9++ {
        v8 += v4
        _ = fi9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        if (913 != 32767) {
            switch fi10 {
            case 42:
                if ((fi10 >> ((1) & 63)) < (fi10 ^ (fi10 ^ fi10))) {
                    if false {
                        switch ("a" + "a") {
                        case "a":
                            v4 += -1
                        }
                        var v11 int = 0
                        for fi12 := 0; fi12 < 8; fi12++ {
                            if true {
                                v13 := GenS1{F0: "a", F1: -232, F2: v4}
                                var v14 float64 = 3.14159
                                switch (("\n\t" + "") + "foo bar") {
                                case "":
                                    fmt.Println(fi10, v11)
                                case "foo bar":
                                    v1[(fi10 / (((v11 * (fi10 << ((293) & 63)))) | 1))] = (genFn0() <= -113)
                                }
                                _ = v13
                                _ = v14
                            }
                            var v15 uint16 = 0
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi12
                            _ = v15
                        }
                        _ = v11
                    }
                    switch (fi10 * (fi10 << ((255) & 63))) {
                    case -2147483648:
                        var v16 float64 = float64((float32(-2.25) - -2.25))
                        _ = v16
                    case 9223372036854775807:
                        if (!false) {
                            v3 = (v3 ^ v3)
                            if (!false) {
                                for fi17 := 0; fi17 < 4; fi17++ {
                                    switch (fi17 ^ fi10) {
                                    case -690:
                                        if ((("a" + ("foo bar" + "a")) == "") && ((false && (fi17 == int(v7))) && false)) {
                                            switch fi17 {
                                            case -707:
                                                switch (-(fi17)) {
                                                case -77:
                                                    v18 := GenS1{F0: "foo bar", F1: 1, F2: v4}
                                                    _ = v18
                                                case 127:
                                                    if (fi17 <= (fi17 + (fi17 << ((4294967295) & 63)))) {
                                                        v19 := GenS2{F0: 560}
                                                        fmt.Println(v7, v7, v4)
                                                        var v20 int64 = (v3 << ((24) & 63))
                                                        _ = v19
                                                        _ = v20
                                                    }
                                                }
                                            case -193:
                                                fmt.Println(fi17, v7, v3)
                                            case -9223372036854775808:
                                                v21 := GenS2{F0: -2147483648}
                                                fmt.Println(v7, v8)
                                                switch "a" {
                                                case "":
                                                    if true {
                                                        v2 = append(v2, "\n\t")
                                                    } else {
                                                        v8 = v8
                                                    }
                                                case "\n\t":
                                                    v8 = (v8 & -1)
                                                }
                                                if false {
                                                    v7 = v7
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = "\n\t"
                                                    }
                                                    if (true && false) {
                                                        v22 := GenS1{F0: (("a" + ("foo bar" + ("" + "foo bar"))) + "\n\t"), F1: -424, F2: (v4 | (v8 % ((961) | 1)))}
                                                        v3 -= (v3 << ((65535) & 63))
                                                        v8 += (^(v8))
                                                        _ = v22
                                                    }
                                                    delete(v1, fi10)
                                                }
                                                _ = v21
                                            }
                                            if false {
                                                v8 = v8
                                            } else {
                                                v4 -= v4
                                            }
                                            var v23 int = (fi10 - fi10)
                                            _ = v23
                                        }
                                    case -533:
                                        var v24 string = ("\n\t" + ("hello" + (("\n\t" + "") + "")))
                                        _ = v24
                                    case 80:
                                        var v25 uint32 = 65535
                                        _ = v25
                                    }
                                    var v26 string = "\n\t"
                                    _ = fi17
                                    _ = v26
                                }
                            }
                        }
                    case 4294967295:
                        v7 -= (v7 & v7)
                    }
                    var v27 uint64 = 4294967295
                    _ = v27
                }
            case 9223372036854775807:
                var v28 int32 = 272
                _ = v28
            }
            fmt.Println(v8, v4)
            v7 = (v7 * v7)
        }
        _ = fi10
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        for fi30 := 0; fi30 < 9; fi30++ {
            v8 = (-(v4))
            if (32767 < 32767) {
                v7 += v7
                var v31 uint32 = 42
                _ = v31
            }
            switch ("a" + "a") {
            case "":
                var v32 uint16 = 2
                _ = v32
            }
            _ = fi30
        }
        v3 = v3
        v3 += (v3 + (v3 + v3))
        _ = fi29
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        switch fi33 {
        case -550:
            v34 := GenS2{F0: 42}
            _ = v34
        }
        var v35 uint8 = 1
        var v36 int64 = (v3 - v3)
        _ = fi33
        _ = v35
        _ = v36
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v7 += 0
    if (!(!true)) {
        for fi37 := 0; fi37 < 9; fi37++ {
            var v38 int32 = 552
            _ = fi37
            _ = v38
        }
    } else {
        switch 128 {
        case 4294967295:
            if false {
                if false {
                    switch 161 {
                    case 677:
                        for fi39 := 0; fi39 < 20; fi39++ {
                            fmt.Println(v4, v8)
                            _ = fi39
                        }
                    }
                    for fi40 := 0; fi40 < 17; fi40++ {
                        var v41 float32 = (1.5 - (1e10 * (float32(3.14159) + (1e10 / 1e10))))
                        for fi42 := 0; fi42 < 16; fi42++ {
                            v43 := GenS1{F0: ("hello" + "hello"), F1: -2147483648, F2: 1}
                            if false {
                                if true {
                                    if (!(true || (135 != 845))) {
                                        var v44 uint = 4294967295
                                        var v45 uint = v44
                                        _ = v44
                                        _ = v45
                                    }
                                    if ("hello" <= ((("\n\t" + "hello") + "foo bar") + "a")) {
                                        var v46 int32 = int32(v7)
                                        v3 += -473
                                        _ = v46
                                    } else {
                                        if true {
                                            var v47 int32 = 1
                                            if (42 < 2) {
                                                fmt.Println(v47, v3)
                                                v8 -= v8
                                                var v48 uint64 = 729
                                                fmt.Println(v7, v7)
                                                if (!false) {
                                                    v4 = v8
                                                } else {
                                                    var v49 int16 = 494
                                                    _ = v49
                                                }
                                                v8 = v8
                                                _ = v48
                                            }
                                            switch ("" + ("hello" + ("foo bar" + "\n\t"))) {
                                            case "\n\t":
                                                switch (fi42 + -9223372036854775808) {
                                                case 2147483647:
                                                    if false {
                                                        v50 := GenS3{F0: ("a" + "")}
                                                        v51 := GenS2{F0: -1}
                                                        v4 += (v4 ^ v4)
                                                        v7 = v7
                                                        switch "a" {
                                                        case "a":
                                                            var v52 uint32 = 42
                                                            _ = v52
                                                        }
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        var v53 uint64 = 0
                                                        _ = v53
                                                    }
                                                case -9223372036854775808:
                                                    v54 := GenS2{F0: 42}
                                                    _ = v54
                                                case -1:
                                                    switch (fi40 % ((fi40) | 1)) {
                                                    case 9223372036854775807:
                                                        v55 := GenS4{F0: v3}
                                                        _ = v55
                                                    case 0:
                                                        fmt.Println(v4, v4)
                                                    default:
                                                        v56 := GenS5{F0: -853, F1: 4294967295, F2: ""}
                                                        _ = v56
                                                    }
                                                default:
                                                    v47 = 0
                                                    var v57 int64 = (v3 * 1)
                                                    _ = v57
                                                }
                                            }
                                            _ = v47
                                        }
                                        switch (fi42 & 32767) {
                                        case 127:
                                            v58 := GenS6{F0: 65535}
                                            _ = v58
                                        }
                                        v7 = (v7 % (((v7 | (v7 >> ((531) & 63)))) | 1))
                                        fmt.Println(v3, fi40, v3)
                                    }
                                    var v59 int = 128
                                    _ = v59
                                } else {
                                    v60 := GenS4{F0: (v3 << ((0) & 63))}
                                    switch fi40 {
                                    case 4294967295:
                                        v3 -= (-(v3))
                                    default:
                                        var v61 uint = 4294967295
                                        _ = v61
                                    }
                                    v62 := GenS2{F0: 42}
                                    fmt.Println(v3, v3)
                                    v8 -= v8
                                    var v63 uint16 = 65535
                                    _ = v60
                                    _ = v62
                                    _ = v63
                                }
                                if false {
                                    fmt.Println(v8)
                                    var v64 float64 = 1e10
                                    switch -9223372036854775808 {
                                    case -512:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case -288:
                                        switch "\n\t" {
                                        case "a":
                                            fmt.Println(fi42, v4, v7)
                                        case "hello":
                                            v2 = append(v2, (("foo bar" + "foo bar") + "\n\t"))
                                        case "\n\t":
                                            var v65 int8 = v7
                                            _ = v65
                                        }
                                    default:
                                        var v66 uint32 = 65535
                                        _ = v66
                                    }
                                    v3 = v3
                                    v67 := GenS3{F0: ""}
                                    _ = v64
                                    _ = v67
                                } else {
                                    var v68 int16 = v8
                                    var v69 uint16 = 54
                                    var v70 int32 = 691
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                }
                            }
                            v4 = -1
                            _ = fi42
                            _ = v43
                        }
                        _ = fi40
                        _ = v41
                    }
                    if true {
                        v71 := GenS1{F0: "foo bar", F1: 42, F2: (-(v4))}
                        var v72 bool = false
                        _ = v71
                        _ = v72
                    } else {
                        v73 := GenS7{F0: true}
                        for fi74 := 0; fi74 < 16; fi74++ {
                            for fi75 := 0; fi75 < 14; fi75++ {
                                v7 = v7
                                switch (^(fi75)) {
                                case 0:
                                    var v76 int64 = 9223372036854775807
                                    _ = v76
                                }
                                _ = fi75
                            }
                            if false {
                                var v77 int = fi74
                                for fi78 := 0; fi78 < 10; fi78++ {
                                    switch ("hello" + ("hello" + "a")) {
                                    case "hello":
                                        if len(v2) > 0 {
                                            v2[0] = "hello"
                                        }
                                    case "\n\t":
                                        if false {
                                            var v79 int64 = v3
                                            var v80 uint = 2
                                            if true {
                                                var v81 int = 1
                                                v81 -= (fi74 * int(v80))
                                                var v82 int8 = (v7 >> (((v80 >> ((v80) & 63))) & 63))
                                                if (42 <= (^(v80))) {
                                                    v83 := GenS2{F0: -2147483648}
                                                    var v84 int8 = v82
                                                    _ = v83
                                                    _ = v84
                                                }
                                                switch fi78 {
                                                case 128:
                                                    fmt.Println(v8, fi78)
                                                default:
                                                    if (!true) {
                                                        switch -599 {
                                                        case 32767:
                                                            fmt.Println(v79, v3)
                                                        case 128:
                                                            v85 := GenS3{F0: "hello"}
                                                            _ = v85
                                                        case -9223372036854775808:
                                                            if (false || true) {
                                                                v2 = append(v2, "\n\t")
                                                                _ = len(v1)
                                                            }
                                                        default:
                                                            fmt.Println(v77, v8)
                                                        }
                                                    } else {
                                                        var v86 float32 = 1.5
                                                        _ = v86
                                                    }
                                                }
                                                _ = v81
                                                _ = v82
                                            } else {
                                                var v87 uint = (^(v80))
                                                var v88 int32 = -23
                                                v3 = (^(v79))
                                                _ = v87
                                                _ = v88
                                            }
                                            _ = v79
                                            _ = v80
                                        } else {
                                            var v89 bool = (898 < v3)
                                            fmt.Println(v4)
                                            fmt.Println(v4, v4)
                                            _ = v89
                                        }
                                    }
                                    var v90 int16 = v8
                                    v2 = append(v2, (("hello" + "hello") + "a"))
                                    v8 += (^(v8))
                                    v91 := GenS8{F0: 255}
                                    _ = fi78
                                    _ = v90
                                    _ = v91
                                }
                                if true {
                                    switch (v77 + (-(fi74))) {
                                    case 390:
                                        var v92 float64 = 0.1
                                        _ = v92
                                    default:
                                        if (!(2147483647 < 935)) {
                                            v93 := GenS5{F0: -522, F1: 1, F2: "a"}
                                            _ = v93
                                        } else {
                                            var v94 int = -212
                                            v95 := GenS8{F0: 127}
                                            if false {
                                                var v96 string = ("" + ((("foo bar" + "\n\t") + "a") + ""))
                                                _ = v96
                                            }
                                            _ = v94
                                            _ = v95
                                        }
                                    }
                                }
                                _ = v77
                            }
                            var v97 string = "foo bar"
                            _ = fi74
                            _ = v97
                        }
                        _ = v73
                    }
                }
                var v98 int32 = 42
                _ = v98
            } else {
                v7 = 1
                v99 := GenS8{F0: 1}
                _ = v99
            }
        default:
            if false {
                fmt.Println(v4)
                if (4294967295 == 0) {
                    v7 -= (v7 ^ v7)
                    switch 128 {
                    case 128:
                        v8 += (-(v8))
                    }
                } else {
                    switch -9223372036854775808 {
                    case 908:
                        for fi100 := 0; fi100 < 1; fi100++ {
                            v8 = v8
                            fmt.Println(fi100)
                            _ = fi100
                        }
                    default:
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                    }
                    v101 := GenS9{F0: true, F1: 4294967295}
                    fmt.Println(v8, v7)
                    _ = v101
                }
                var v102 int32 = 0
                _ = v102
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
