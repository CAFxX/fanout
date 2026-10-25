package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32; F1 int; F2 int32 }
type GenS9 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int32; F1 uint32; F2 int64 }
type GenS12 struct { F0 uint16; F1 bool }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint16; F1 int32 }
type GenS15 struct { F0 int; F1 int32 }
type GenS16 struct { F0 uint; F1 uint64 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 bool; F1 int8; F2 uint64 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 uint32; F1 string }
type GenS23 struct { F0 string; F1 int32; F2 uint8 }
type GenS24 struct { F0 int }
func genFn0() uint64 {
    v6 := GenS1{F0: 128, F1: true, F2: 65535}
    var v7 uint32 = 1
    v8 := []int{-2147483648, 1, 408}
    _ = v6
    _ = v7
    _ = v8
    return 391
}

func main() {
    v1 := []int{-1, 127}
    var v2 float32 = 1.5
    var v3 int32 = 42
    var v4 float32 = v2
    v5 := []uint32{2147483647, 736, 1, 0}
    switch 231 {
    case 32768:
        v3 = v3
    case 1:
        switch ("foo bar" + "a") {
        case "":
            v9 := GenS2{F0: 127}
            _ = v9
        }
    }
    var v10 uint = 342
    v1 = append(v1, 32767)
    v3 += v3
    var v11 bool = false
    v12 := GenS2{F0: 42}
    v10 -= v10
    var v13 int16 = 32767
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            fmt.Println(v10, v10)
            var v16 int16 = v13
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v13 += (v13 ^ v13)
    for fi17 := 0; fi17 < 4; fi17++ {
        if (v11 && v11) {
            if (42 < -9223372036854775808) {
                switch fi17 {
                case -1:
                    switch ("foo bar" + ((("a" + "hello") + "a") + "hello")) {
                    case "foo bar":
                        fmt.Println(v2, fi17, fi17)
                    case "hello":
                        fmt.Println(v3, v13, v10)
                    case "\n\t":
                        fmt.Println(v4)
                    }
                case 32768:
                    v13 = -489
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                if (!(42 < 820)) {
                    if (v11 || (343 < 0)) {
                        for fi18 := 0; fi18 < 5; fi18++ {
                            var v19 int16 = v13
                            _ = fi18
                            _ = v19
                        }
                        if v11 {
                            v13 = v13
                        }
                        var v20 uint64 = 975
                        _ = v20
                    } else {
                        if v11 {
                            var v21 float64 = float64(v2)
                            var v22 uint8 = 2
                            var v23 int = (fi17 & -2147483648)
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            if v11 {
                                fmt.Println(v2, v3, v11)
                                if len(v1) > 0 {
                                    v1[0] = (^(fi17))
                                }
                                var v24 string = ""
                                _ = v24
                            }
                            for fi25 := 0; fi25 < 8; fi25++ {
                                v26 := GenS3{F0: 2}
                                if ("\n\t" < "foo bar") {
                                    v27 := GenS4{F0: v3, F1: v10}
                                    fmt.Println(v10, fi17)
                                    fmt.Println(v4)
                                    _ = v27
                                }
                                _ = fi25
                                _ = v26
                            }
                            if v11 {
                                if v11 {
                                    for fi28 := 0; fi28 < 2; fi28++ {
                                        v5 = append(v5, 431)
                                        _ = fi28
                                    }
                                    v29 := GenS4{F0: 958, F1: (v10 & (v10 % ((v10) | 1)))}
                                    _ = v29
                                }
                                v10 = (^(v10))
                                fmt.Println(v11, v2, v3)
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                for fi30 := 0; fi30 < 11; fi30++ {
                                    v31 := GenS1{F0: int(v10), F1: ((42 < 1) && (v11 || (!(!(v10 <= v10))))), F2: 1}
                                    var v32 int16 = 0
                                    if (!((v10 > v10) && (!v11))) {
                                        if (fi17 < fi30) {
                                            v13 += v32
                                            v3 -= v3
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = (fi17 % ((fi30) | 1))
                                            }
                                        }
                                    }
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                                for fi33 := 0; fi33 < 13; fi33++ {
                                    v3 = -410
                                    _ = fi33
                                }
                            }
                        }
                        v34 := GenS5{F0: -1, F1: (("\n\t" + "\n\t") + "\n\t")}
                        _ = v34
                    }
                    if false {
                        v5 = append(v5, 4294967295)
                        fmt.Println(v4, fi17)
                    } else {
                        for fi35 := 0; fi35 < 10; fi35++ {
                            fmt.Println(fi35, v3)
                            v36 := GenS2{F0: 2}
                            _ = fi35
                            _ = v36
                        }
                        if true {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v37 int16 = v13
                            switch "a" {
                            case "hello":
                                v5 = append(v5, 0)
                            case "\n\t":
                                var v38 int8 = 38
                                _ = v38
                            default:
                                v3 -= (v3 % ((v3) | 1))
                            }
                            _ = v37
                        }
                    }
                    for fi39 := 0; fi39 < 3; fi39++ {
                        v13 = v13
                        switch "" {
                        case "a":
                            fmt.Println(v13, fi39)
                        }
                        if false {
                            switch "" {
                            case "a":
                                v3 = -1
                            case "hello":
                                v40 := GenS6{F0: fi17, F1: (fi17 | (-(fi17)))}
                                _ = v40
                            default:
                                switch fi39 {
                                case 128:
                                    v11 = false
                                case 32768:
                                    switch fi17 {
                                    case 542:
                                        var v41 int = -2147483648
                                        _ = v41
                                    }
                                case 229:
                                    v11 = (1 < 255)
                                default:
                                    switch fi39 {
                                    case -583:
                                        var v42 string = "foo bar"
                                        _ = v42
                                    case 0:
                                        var v43 int16 = v13
                                        if (127 == (^(fi39))) {
                                            v44 := GenS1{F0: (-(fi17)), F1: v11, F2: 4294967295}
                                            v10 = (v10 % (((^(v10))) | 1))
                                            _ = v44
                                        }
                                        if (v2 <= v2) {
                                            if ((!((3.14159 * 0.1) != 0.1)) || (42 <= 42)) {
                                                if (((!(-2.25 < 1.5)) && v11) || (v11 && (!false))) {
                                                    var v45 int16 = (v13 % ((v13) | 1))
                                                    if len(v5) > 0 {
                                                        v5[0] = 607
                                                    }
                                                    v46 := GenS2{F0: 42}
                                                    if (v11 || ((!(v4 <= 1.5)) || true)) {
                                                        fmt.Println(fi39, v4, v43)
                                                        if ((v4 * (float32((((((float64(0.1) + (1.5 + 3.14159)) - 1.5) * 1e10) + 0.0) / (1.5 / 1.0))) + v4)) != 1.5) {
                                                            if false {
                                                                if len(v5) > 0 {
                                                                    v5[0] = 2
                                                                }
                                                                var v47 uint64 = 18446744073709551615
                                                                _ = v47
                                                            } else {
                                                                var v48 int32 = 0
                                                                var v49 int8 = 47
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        } else {
                                                            if true {
                                                                fmt.Println(fi17)
                                                            }
                                                            v50 := GenS2{F0: 42}
                                                            if v11 {
                                                                v51 := GenS7{F0: "foo bar"}
                                                                v10 -= v10
                                                                v52 := GenS7{F0: "a"}
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                            _ = v50
                                                        }
                                                        v53 := GenS8{F0: v3, F1: fi39, F2: v3}
                                                        v3 = v3
                                                        _ = v53
                                                    }
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    v4 -= v2
                                                    var v54 uint16 = 610
                                                    _ = v54
                                                }
                                                fmt.Println(v4, fi39)
                                                var v55 int = (fi39 + (fi39 % (((fi17 / (((fi39 | fi39)) | 1))) | 1)))
                                                _ = v55
                                            }
                                            fmt.Println(v43, v13)
                                            var v56 uint16 = 255
                                            _ = v56
                                        } else {
                                            var v57 int32 = (v3 + 42)
                                            if false {
                                                var v58 uint32 = 560
                                                fmt.Println(v10, v4)
                                                fmt.Println(v4, v58, fi39)
                                                _ = v58
                                            } else {
                                                switch "" {
                                                case "a":
                                                    if v11 {
                                                        v13 = v43
                                                        switch fi17 {
                                                        case 830:
                                                            v59 := GenS1{F0: fi39, F1: (!v11), F2: 750}
                                                            _ = v59
                                                        case -671:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v60 := GenS7{F0: ("hello" + "foo bar")}
                                                                _ = v60
                                                            default:
                                                                if (42 > (^(v57))) {
                                                                    if (genFn0() > 285) {
                                                                        v43 = (v43 ^ 1000)
                                                                        if v11 {
                                                                            var v61 int32 = v57
                                                                            fmt.Println(v61, v57, v3)
                                                                            if ("\n\t" < "") {
                                                                                v62 := GenS9{F0: 2, F1: 65535, F2: 0}
                                                                                _ = v62
                                                                            }
                                                                            _ = v61
                                                                        } else {
                                                                            var v63 string = "a"
                                                                            fmt.Println(v10, v4)
                                                                            _ = v63
                                                                        }
                                                                    } else {
                                                                        _ = len(v5)
                                                                    }
                                                                    v64 := GenS10{F0: 92}
                                                                    switch (fi39 / (((fi39 & (fi17 - (fi17 >> (((v10 * 623)) & 63))))) | 1)) {
                                                                    case 1:
                                                                        v11 = false
                                                                    case -9223372036854775808:
                                                                        v10 = 42
                                                                        if false {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                fmt.Println(fi39, v43, fi39)
                                                                            default:
                                                                                if false {
                                                                                    v65 := GenS11{F0: (v57 / ((2147483647) | 1)), F1: 42, F2: 0}
                                                                                    var v66 uint = 2
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                    _ = v65
                                                                                    _ = v66
                                                                                } else {
                                                                                    if true {
                                                                                        if ((v57 < v3) || true) {
                                                                                            switch (fi39 - fi17) {
                                                                                            case -9223372036854775808:
                                                                                                if v11 {
                                                                                                    fmt.Println(fi39)
                                                                                                    v3 = (v3 & v3)
                                                                                                } else {
                                                                                                    switch "\n\t" {
                                                                                                    case "foo bar":
                                                                                                        var v67 int64 = 1
                                                                                                        _ = v67
                                                                                                    case "hello":
                                                                                                        if v11 {
                                                                                                            switch ("\n\t" + "a") {
                                                                                                            case "foo bar":
                                                                                                                if ((v13 % ((v43) | 1)) < v13) {
                                                                                                                    var v68 uint = 42
                                                                                                                    _ = v68
                                                                                                                } else {
                                                                                                                    _ = len(v5)
                                                                                                                    var v69 int8 = 1
                                                                                                                    _ = v69
                                                                                                                }
                                                                                                            case "a":
                                                                                                                v70 := GenS9{F0: 45, F1: 42, F2: 42}
                                                                                                                fmt.Println(v13, fi17)
                                                                                                                v57 = 2147483647
                                                                                                                v43 += 0
                                                                                                                _ = v70
                                                                                                            case "\n\t":
                                                                                                                var v71 uint8 = 1
                                                                                                                _ = v71
                                                                                                            }
                                                                                                            _ = len(v5)
                                                                                                            v3 = v3
                                                                                                            var v72 string = ("foo bar" + ((("hello" + "\n\t") + "foo bar") + ((("a" + "foo bar") + "foo bar") + "hello")))
                                                                                                            _ = v72
                                                                                                        } else {
                                                                                                            v43 = (v43 ^ -186)
                                                                                                        }
                                                                                                    }
                                                                                                    var v73 int16 = (v13 << (((v10 % ((v10) | 1))) & 63))
                                                                                                    _ = v73
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v74 uint = (^(v10))
                                                                                        _ = v74
                                                                                    }
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                    var v75 float64 = (0.0 * 3.14159)
                                                                                    _ = v75
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        v13 = (v43 / ((v43) | 1))
                                                                    }
                                                                    _ = v64
                                                                } else {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 0
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        switch "a" {
                                                        case "foo bar":
                                                            var v76 uint = (v10 >> (((v10 >> ((v10) & 63))) & 63))
                                                            _ = v76
                                                        case "a":
                                                            v5 = append(v5, 2147483648)
                                                        default:
                                                            if false {
                                                                if (0 >= 2) {
                                                                    if (true && ((^(v13)) >= 1)) {
                                                                        if true {
                                                                            var v77 int16 = 0
                                                                            _ = v77
                                                                        }
                                                                        v10 -= 0
                                                                        v4 += v2
                                                                        var v78 int8 = 42
                                                                        var v79 uint8 = 42
                                                                        _ = v78
                                                                        _ = v79
                                                                    }
                                                                }
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                if len(v5) > 0 {
                                                                    v5[0] = 4294967295
                                                                }
                                                            } else {
                                                                _ = len(v1)
                                                                switch fi39 {
                                                                case 127:
                                                                    v43 -= (v43 << (((v10 - 65535)) & 63))
                                                                    v13 = (v43 | (^(v13)))
                                                                    if (v11 || true) {
                                                                        var v80 int64 = -317
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        if (4294967295 == genFn0()) {
                                                                            v43 = v13
                                                                            v11 = v11
                                                                            v81 := GenS12{F0: 317, F1: (2 <= 42)}
                                                                            v82 := GenS9{F0: 1, F1: 1, F2: v80}
                                                                            v83 := GenS13{F0: 56}
                                                                            switch fi17 {
                                                                            case 659:
                                                                                var v84 int = fi39
                                                                                _ = v84
                                                                            default:
                                                                                if false {
                                                                                    switch "a" {
                                                                                    case "a":
                                                                                        v57 = 42
                                                                                    case "hello":
                                                                                        _ = len(v5)
                                                                                    case "foo bar":
                                                                                        v85 := GenS4{F0: (v57 % ((390) | 1)), F1: (v10 >> ((v10) & 63))}
                                                                                        _ = v85
                                                                                    }
                                                                                    var v86 uint64 = 18446744073709551615
                                                                                    v87 := GenS9{F0: 251, F1: 32767, F2: v80}
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                }
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = 0
                                                                                }
                                                                            }
                                                                            _ = v81
                                                                            _ = v82
                                                                            _ = v83
                                                                        }
                                                                        var v88 int8 = -30
                                                                        switch ("foo bar" + ((("foo bar" + ("" + ("a" + "foo bar"))) + (("a" + "\n\t") + ((("a" + "") + "a") + "foo bar"))) + "a")) {
                                                                        case "foo bar":
                                                                            if (255 > 2) {
                                                                                var v89 uint = (v10 | (v10 | v10))
                                                                                _ = v89
                                                                            }
                                                                        default:
                                                                            fmt.Println(v13, v88)
                                                                        }
                                                                        var v90 bool = v11
                                                                        if v11 {
                                                                            var v91 int = fi39
                                                                            if len(v5) > 0 {
                                                                                v5[0] = 65535
                                                                            }
                                                                            v88 = v88
                                                                            _ = v91
                                                                        } else {
                                                                            v92 := GenS7{F0: "a"}
                                                                            _ = v92
                                                                        }
                                                                        if true {
                                                                            if (true && true) {
                                                                                fmt.Println(fi17, v80, v80)
                                                                            }
                                                                            var v93 int8 = -1
                                                                            if (!(!false)) {
                                                                                v4 += (-2.25 + v2)
                                                                                v94 := GenS14{F0: 482, F1: v3}
                                                                                _ = v94
                                                                            }
                                                                            v57 = v57
                                                                            _ = v93
                                                                        }
                                                                        _ = v80
                                                                        _ = v88
                                                                        _ = v90
                                                                    }
                                                                    fmt.Println(v10, fi39, v2)
                                                                    v95 := GenS2{F0: 207}
                                                                    var v96 float32 = (v4 / v2)
                                                                    fmt.Println(v10, v43, v43)
                                                                    if (v11 || v11) {
                                                                        if (true && v11) {
                                                                            v97 := GenS6{F0: fi17, F1: (fi17 / (((fi17 % ((fi39) | 1))) | 1))}
                                                                            v2 -= (0.1 + 1.5)
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            _ = v97
                                                                        }
                                                                        v98 := GenS12{F0: 203, F1: v11}
                                                                        _ = v98
                                                                    } else {
                                                                        if (!false) {
                                                                            var v99 int16 = (v43 - (v43 >> ((42) & 63)))
                                                                            _ = v99
                                                                        } else {
                                                                            if (518 <= 620) {
                                                                                switch 4294967295 {
                                                                                case 1:
                                                                                    v5 = append(v5, 0)
                                                                                case 200:
                                                                                    if false {
                                                                                        v100 := GenS15{F0: 42, F1: (v57 / (((v57 - v3)) | 1))}
                                                                                        fmt.Println(v13, v4, v96)
                                                                                        _ = v100
                                                                                    } else {
                                                                                        v5 = append(v5, 4294967295)
                                                                                        v57 = (v57 / ((633) | 1))
                                                                                    }
                                                                                case -2147483648:
                                                                                    v43 += (-(v43))
                                                                                default:
                                                                                    if (100 <= 45) {
                                                                                        switch (((("\n\t" + "foo bar") + ("" + (("" + ("hello" + "hello")) + "a"))) + "a") + "a") {
                                                                                        case "\n\t":
                                                                                            var v101 int8 = 42
                                                                                            _ = v101
                                                                                        default:
                                                                                            var v102 string = "\n\t"
                                                                                            _ = v102
                                                                                        }
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v57, v4)
                                                                            }
                                                                        }
                                                                        var v103 uint8 = uint8(v10)
                                                                        _ = v103
                                                                    }
                                                                    _ = v95
                                                                    _ = v96
                                                                case -675:
                                                                    if true {
                                                                        switch ("foo bar" + (("\n\t" + "") + "hello")) {
                                                                        case "a":
                                                                            var v104 int64 = -349
                                                                            _ = v104
                                                                        case "":
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                fmt.Println(v57, v11, v4)
                                                                            case "a":
                                                                                if v11 {
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                    var v105 uint64 = 2
                                                                                    _ = v105
                                                                                }
                                                                            case "foo bar":
                                                                                if true {
                                                                                    v106 := GenS10{F0: 950}
                                                                                    v3 -= 0
                                                                                    v43 = -32768
                                                                                    _ = v106
                                                                                }
                                                                            default:
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        var v107 float64 = (0.0 * ((0.0 - (-2.25 - 1.5)) + 1e10))
                                                                                        v2 = (v4 - v4)
                                                                                        _ = v107
                                                                                    } else {
                                                                                        var v108 int8 = -128
                                                                                        fmt.Println(v13, v108)
                                                                                        fmt.Println(v11)
                                                                                        _ = v108
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    v43 += v43
                                                                }
                                                            }
                                                        }
                                                        v11 = ((-2.25 - v2) <= v4)
                                                        switch ("hello" + "foo bar") {
                                                        case "hello":
                                                            if (!v11) {
                                                                v109 := GenS10{F0: 1}
                                                                _ = v109
                                                            } else {
                                                                v110 := GenS8{F0: v57, F1: fi17, F2: v57}
                                                                var v111 string = "\n\t"
                                                                _ = v110
                                                                _ = v111
                                                            }
                                                        case "foo bar":
                                                            switch 4294967295 {
                                                            case -805:
                                                                if v11 {
                                                                    switch (fi17 * 2147483647) {
                                                                    case 1:
                                                                        if ("\n\t" > "\n\t") {
                                                                            _ = len(v1)
                                                                        } else {
                                                                            v3 += v57
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            fmt.Println(fi39)
                                                                            fmt.Println(fi39)
                                                                        }
                                                                    default:
                                                                        v1 = append(v1, int(v43))
                                                                    }
                                                                    v112 := GenS14{F0: 32767, F1: -116}
                                                                    if ((v11 && v11) && v11) {
                                                                        switch fi39 {
                                                                        case 32767:
                                                                            v2 = v4
                                                                        case 4294967295:
                                                                            v10 += v10
                                                                        case 128:
                                                                            v113 := GenS16{F0: 4294967295, F1: 1}
                                                                            _ = v113
                                                                        default:
                                                                            fmt.Println(v2, v43, fi39)
                                                                        }
                                                                        v57 += v57
                                                                    }
                                                                    _ = v112
                                                                }
                                                            case -617:
                                                                switch int(v10) {
                                                                case 127:
                                                                    if v11 {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        fmt.Println(v2, v3, v11)
                                                                        var v114 float64 = (1.5 / float64(v2))
                                                                        _ = v114
                                                                    }
                                                                case -9223372036854775808:
                                                                    if ((v43 + (v13 % ((0) | 1))) != (v13 & (v13 / ((int16(fi17)) | 1)))) {
                                                                        fmt.Println(fi39)
                                                                    }
                                                                case -910:
                                                                    v1 = append(v1, (fi39 ^ (fi39 & fi17)))
                                                                }
                                                            }
                                                        default:
                                                            if true {
                                                                v115 := GenS16{F0: (v10 * v10), F1: 1}
                                                                _ = v115
                                                            } else {
                                                                fmt.Println(v43, v43)
                                                                if (!v11) {
                                                                    fmt.Println(v10, v57, v2)
                                                                    var v116 int16 = v13
                                                                    var v117 int64 = -9223372036854775808
                                                                    _ = v116
                                                                    _ = v117
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                v3 -= v57
                                                if (1 >= v43) {
                                                    if ((v10 | (v10 | (^(v10)))) <= 65535) {
                                                        _ = len(v5)
                                                        if len(v5) > 0 {
                                                            v5[0] = 870
                                                        }
                                                        v118 := GenS17{F0: "\n\t"}
                                                        _ = v118
                                                    }
                                                    var v119 int = (fi39 | fi17)
                                                    if len(v1) > 0 {
                                                        v1[0] = (fi17 + (fi39 - v119))
                                                    }
                                                    _ = v119
                                                } else {
                                                    v4 -= v2
                                                    fmt.Println(v10, fi39, fi17)
                                                    v120 := GenS9{F0: 0, F1: 630, F2: 42}
                                                    _ = v120
                                                }
                                            }
                                            var v121 uint8 = 127
                                            _ = v121
                                            _ = v57
                                        }
                                        _ = v43
                                    default:
                                        v10 = v10
                                    }
                                }
                            }
                            fmt.Println(v2, fi39)
                            v2 -= 1e10
                        } else {
                            switch fi17 {
                            case -884:
                                var v122 int32 = (v3 - 551)
                                _ = v122
                            case 935:
                                switch 331 {
                                case 2147483647:
                                    if len(v1) > 0 {
                                        v1[0] = (fi39 << (((v10 / ((v10) | 1))) & 63))
                                    }
                                case -1:
                                    switch 0 {
                                    case 32768:
                                        var v123 int16 = (v13 << (((v10 | v10)) & 63))
                                        switch (fi39 << ((v10) & 63)) {
                                        case 32768:
                                            v124 := GenS6{F0: (fi17 << ((304) & 63)), F1: fi17}
                                            v13 += v13
                                            var v125 int32 = v3
                                            _ = v124
                                            _ = v125
                                        default:
                                            fmt.Println(v3, v123)
                                        }
                                        fmt.Println(fi39)
                                        _ = v123
                                    case 127:
                                        switch (fi39 ^ 4294967295) {
                                        case -2147483648:
                                            var v126 int = (fi39 - (fi39 << (((^(v10))) & 63)))
                                            _ = v126
                                        default:
                                            var v127 float64 = 0.0
                                            _ = v127
                                        }
                                    case 128:
                                        v10 = v10
                                    }
                                case 128:
                                    fmt.Println(v3, v4, fi17)
                                }
                            default:
                                if true {
                                    if true {
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v128 uint16 = 0
                                            _ = v128
                                        case "a":
                                            if (fi39 != fi17) {
                                                var v129 float64 = ((0.1 / float64(v4)) - 0.0)
                                                switch "hello" {
                                                case "\n\t":
                                                    var v130 string = ""
                                                    _ = v130
                                                default:
                                                    var v131 uint32 = 0
                                                    _ = v131
                                                }
                                                _ = v129
                                            } else {
                                                var v132 int16 = (v13 + v13)
                                                v133 := GenS18{F0: 862, F1: (v2 >= v2), F2: 74}
                                                _ = v132
                                                _ = v133
                                            }
                                        case "\n\t":
                                            v134 := GenS18{F0: 42, F1: v11, F2: 4294967295}
                                            _ = v134
                                        }
                                        v10 -= (v10 | (v10 >> ((4294967295) & 63)))
                                        var v135 float64 = ((1e10 - (float64(v4) + 3.14159)) - 0.0)
                                        _ = v135
                                    } else {
                                        v11 = (164 != 2)
                                    }
                                    v136 := GenS10{F0: 2}
                                    _ = v136
                                } else {
                                    v137 := GenS17{F0: "hello"}
                                    _ = v137
                                }
                            }
                            v138 := GenS8{F0: int32(v10), F1: -70, F2: v3}
                            _ = v138
                        }
                        _ = fi39
                    }
                } else {
                    var v139 float64 = (float64(1e10) + 0.0)
                    if false {
                        v5 = append(v5, 0)
                        v140 := GenS15{F0: (-(fi17)), F1: (v3 + (v3 >> (((v10 | v10)) & 63)))}
                        _ = len(v5)
                        _ = v140
                    }
                    _ = v139
                }
                for fi141 := 0; fi141 < 13; fi141++ {
                    v10 -= v10
                    var v142 bool = false
                    _ = fi141
                    _ = v142
                }
            }
            var v143 int = 2147483647
            _ = v143
        } else {
            for fi144 := 0; fi144 < 11; fi144++ {
                v4 = (-2.25 * v2)
                var v145 uint8 = 127
                v10 = (v10 << (((v10 - (v10 >> ((v10) & 63)))) & 63))
                v145 = (v145 & 54)
                if (225 == 255) {
                    v146 := GenS19{F0: 840}
                    v147 := GenS1{F0: 127, F1: false, F2: 56}
                    switch fi17 {
                    case 1:
                        var v148 int64 = 0
                        _ = v148
                    default:
                        var v149 int64 = 1
                        _ = v149
                    }
                    _ = v146
                    _ = v147
                } else {
                    v150 := GenS20{F0: (9223372036854775807 >= 9223372036854775807), F1: 42, F2: 355}
                    v3 += (v3 * v3)
                    if true {
                        _ = len(v5)
                        var v151 int8 = 127
                        var v152 uint64 = 42
                        v153 := GenS21{F0: v11}
                        if v11 {
                            var v154 uint64 = (v152 ^ (v152 % ((8) | 1)))
                            if len(v1) > 0 {
                                v1[0] = (^(fi17))
                            }
                            v155 := GenS22{F0: uint32(v152), F1: (((("a" + "\n\t") + "\n\t") + "hello") + ((("hello" + "hello") + ("hello" + "hello")) + ("\n\t" + "")))}
                            _ = v154
                            _ = v155
                        }
                        _ = v151
                        _ = v152
                        _ = v153
                    }
                    v156 := GenS17{F0: "hello"}
                    _ = v150
                    _ = v156
                }
                switch 0 {
                case 32767:
                    var v157 float64 = -2.25
                    _ = v157
                case 42:
                    v158 := GenS7{F0: ""}
                    var v159 uint32 = 2147483648
                    fmt.Println(v4)
                    v1 = append(v1, fi144)
                    _ = v158
                    _ = v159
                case 4294967295:
                    v160 := GenS18{F0: 483, F1: (!v11), F2: 2147483648}
                    _ = v160
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi144
                _ = v145
            }
            v161 := GenS23{F0: "", F1: (v3 << (((v10 / ((1) | 1))) & 63)), F2: 128}
            switch ("foo bar" + ("" + ("foo bar" + "hello"))) {
            case "hello":
                if (!true) {
                    var v162 int8 = int8(v13)
                    if (127 <= 42) {
                        var v163 uint64 = 4294967295
                        fmt.Println(v11, v10)
                        _ = v163
                    }
                    _ = v162
                }
            case "\n\t":
                v164 := GenS20{F0: (!v11), F1: 100, F2: 2}
                _ = v164
            case "a":
                v165 := GenS24{F0: (fi17 | fi17)}
                _ = v165
            }
            _ = v161
        }
        if (("a" != ("" + "hello")) && true) {
            if (65535 < 255) {
                v11 = (1 == 42)
                var v166 uint16 = 355
                v10 = v10
                _ = v166
            } else {
                _ = len(v1)
                v2 += ((((v2 - v4) + (v2 + (3.14159 / (1e10 - v2)))) - 0.0) * (v2 - ((((float32(-2.25) * v2) + v2) / (v4 * v4)) * 0.0)))
                var v167 uint64 = 1
                _ = v167
            }
        } else {
            v168 := GenS15{F0: fi17, F1: -51}
            fmt.Println(v11, v3, v11)
            var v169 float32 = (float32(float64(v4)) * (0.1 / v4))
            _ = v168
            _ = v169
        }
        _ = fi17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
