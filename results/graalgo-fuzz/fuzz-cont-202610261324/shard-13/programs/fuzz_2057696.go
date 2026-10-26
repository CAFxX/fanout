package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint; F1 int64; F2 int32 }
type GenS6 struct { F0 int; F1 string; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 bool; F1 int8 }
type GenS12 struct { F0 uint; F1 int16; F2 bool }
type GenS13 struct { F0 int16; F1 int16 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int64; F1 uint; F2 int }
func genFn0() uint {
    v4 := GenS1{F0: 42}
    var v5 int = -673
    v6 := GenS1{F0: 1}
    for fi7 := 0; fi7 < 15; fi7++ {
        v5 = v5
        v5 += (fi7 - (fi7 / (((v5 * v5)) | 1)))
        var v8 float32 = (3.14159 - float32(-2.25))
        _ = fi7
        _ = v8
    }
    _ = v4
    _ = v5
    _ = v6
    return 65535
}

func main() {
    var v1 bool = false
    v2 := make(map[string]bool)
    v3 := []int{9223372036854775807, 1}
    for fi9 := 0; fi9 < 15; fi9++ {
        if (1 == -9223372036854775808) {
            var v10 uint16 = 11
            if (2 != 128) {
                if (v1 || (false && v1)) {
                    switch (fi9 - 0) {
                    case -516:
                        if ((!(v1 && v1)) || v1) {
                            switch (fi9 * -2147483648) {
                            case 128:
                                fmt.Println(v1, fi9)
                            case 4294967295:
                                v11 := GenS1{F0: 768}
                                _ = v11
                            default:
                                switch (fi9 & 32768) {
                                case 9223372036854775807:
                                    v1 = (!(v1 && (284 != 278)))
                                case 248:
                                    v1 = false
                                default:
                                    var v12 uint64 = 504
                                    _ = v12
                                }
                            }
                            switch (fi9 | fi9) {
                            case 710:
                                for fi13 := 0; fi13 < 19; fi13++ {
                                    var v14 string = ""
                                    fmt.Println(v14)
                                    v2[(v14 + v14)] = ((127 >= -1) || (-728 < 1))
                                    v10 = (v10 ^ (v10 + (v10 % ((1) | 1))))
                                    switch v14 {
                                    case "":
                                        if (!((v14 + (v14 + "foo bar")) > "hello")) {
                                            if true {
                                                var v15 float64 = 1e10
                                                v14 = ("foo bar" + v14)
                                                if (!(498 >= 42)) {
                                                    v16 := GenS1{F0: -19}
                                                    v17 := GenS1{F0: 0}
                                                    _ = len(v2)
                                                    _ = v16
                                                    _ = v17
                                                }
                                                _ = v15
                                            } else {
                                                var v18 int16 = 1
                                                v2[v14] = ((!false) && (-147 == 42))
                                                switch fi13 {
                                                case 4294967295:
                                                    v18 += (v18 % (((v18 - -32768)) | 1))
                                                }
                                                _ = v18
                                            }
                                            var v19 uint32 = 865
                                            _ = v19
                                        }
                                    case "hello":
                                        if (fi13 == (fi13 / (((fi13 << ((65535) & 63))) | 1))) {
                                            v20 := GenS1{F0: 99}
                                            if false {
                                                var v21 uint32 = 770
                                                var v22 float32 = 0.0
                                                delete(v2, v14)
                                                _ = v21
                                                _ = v22
                                            }
                                            _ = v20
                                        }
                                    case "\n\t":
                                        _ = len(v2)
                                    }
                                    _ = fi13
                                    _ = v14
                                }
                            case -901:
                                v23 := GenS1{F0: -408}
                                _ = v23
                            case 32767:
                                var v24 float64 = 3.14159
                                _ = v24
                            default:
                                if v1 {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    var v25 uint16 = v10
                                    v10 -= (v25 >> ((4294967295) & 63))
                                    _ = v25
                                }
                            }
                        } else {
                            fmt.Println(v1, v10, v10)
                        }
                    case 2147483647:
                        var v26 uint16 = (v10 / ((562) | 1))
                        _ = v26
                    case 127:
                        fmt.Println(fi9, v10, v1)
                    default:
                        if ((1e10 + (0.1 + (((3.14159 / 1.5) / 0.1) - -2.25))) == 1.5) {
                            if len(v3) > 0 {
                                v3[0] = fi9
                            }
                        } else {
                            v27 := GenS1{F0: 42}
                            for fi28 := 0; fi28 < 3; fi28++ {
                                if (!v1) {
                                    if (v1 || v1) {
                                        fmt.Println(fi9, v1)
                                        if false {
                                            v29 := v2["\n\t"]
                                            var v30 float32 = -2.25
                                            _ = v29
                                            _ = v30
                                        } else {
                                            var v31 string = "a"
                                            var v32 uint64 = 1
                                            var v33 uint64 = v32
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                        }
                                    } else {
                                        var v34 uint64 = 18446744073709551615
                                        _ = v34
                                    }
                                    v1 = false
                                    fmt.Println(v1)
                                }
                                fmt.Println(fi28, v1)
                                v1 = v1
                                _ = fi28
                            }
                            _ = v27
                        }
                    }
                    v1 = true
                    var v35 string = "foo bar"
                    _ = v35
                }
                var v36 uint = 65535
                _ = len(v3)
                _ = v36
            } else {
                for fi37 := 0; fi37 < 13; fi37++ {
                    if false {
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v1)
                        case "":
                            v38 := GenS2{F0: 463, F1: 862, F2: 100}
                            _ = v38
                        default:
                            if false {
                                v39 := v2["a"]
                                if true {
                                    var v40 uint16 = v10
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v10 -= v10
                                    _ = v40
                                } else {
                                    var v41 float32 = (float32(0.1) + -2.25)
                                    switch fi37 {
                                    case 2147483647:
                                        v42 := GenS3{F0: 475, F1: "a"}
                                        _ = v42
                                    case 785:
                                        v10 = (v10 + (v10 << ((2) & 63)))
                                    }
                                    if v1 {
                                        switch fi37 {
                                        case 128:
                                            v43 := GenS4{F0: -9223372036854775808, F1: (^(v10))}
                                            _ = v43
                                        default:
                                            var v44 int8 = 107
                                            _ = v44
                                        }
                                        v45 := GenS2{F0: 2147483648, F1: 18446744073709551615, F2: 127}
                                        _ = v45
                                    } else {
                                        v10 = (^(v10))
                                        v46 := GenS2{F0: 916, F1: 1, F2: 1}
                                        _ = v46
                                    }
                                    _ = v41
                                }
                                v47 := GenS1{F0: -1}
                                var v48 float64 = 1e10
                                _ = v39
                                _ = v47
                                _ = v48
                            }
                        }
                        v10 += (v10 % (((^(v10))) | 1))
                    }
                    if true {
                        v10 = 2
                        switch "hello" {
                        case "a":
                            v49 := GenS3{F0: 1, F1: ""}
                            _ = v49
                        case "foo bar":
                            switch "\n\t" {
                            case "":
                                v50 := GenS2{F0: 271, F1: 442, F2: -128}
                                _ = v50
                            case "foo bar":
                                fmt.Println(v1)
                                switch ("" + "\n\t") {
                                case "":
                                    v1 = v1
                                case "a":
                                    switch "foo bar" {
                                    case "\n\t":
                                        v51 := GenS5{F0: 65535, F1: 1, F2: 0}
                                        var v52 int = (fi37 - (-(fi37)))
                                        _ = v51
                                        _ = v52
                                    case "foo bar":
                                        var v53 string = (("foo bar" + ("" + ("hello" + "a"))) + ("" + ""))
                                        _ = v53
                                    case "hello":
                                        switch -562 {
                                        case 529:
                                            var v54 int = (-(fi37))
                                            _ = v54
                                        default:
                                            switch "hello" {
                                            case "hello":
                                                v1 = false
                                            case "\n\t":
                                                if (-116 > 113) {
                                                    if (false || v1) {
                                                        _ = len(v3)
                                                        var v55 string = "\n\t"
                                                        var v56 float32 = -2.25
                                                        _ = v55
                                                        _ = v56
                                                    } else {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            fmt.Println(fi37, fi37)
                                                        case "a":
                                                            if v1 {
                                                                v10 -= (^(v10))
                                                            }
                                                        case "":
                                                            v10 += (v10 ^ v10)
                                                            if (42 < 543) {
                                                                v57 := GenS1{F0: 42}
                                                                _ = v57
                                                            } else {
                                                                if ((fi9 >> ((167) & 63)) < 32768) {
                                                                    if (1 < 4294967295) {
                                                                        if (127 >= (fi9 * fi9)) {
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            v2[(("\n\t" + (("a" + "") + "")) + "\n\t")] = v1
                                                                        }
                                                                    } else {
                                                                        v58 := GenS4{F0: -506, F1: (v10 >> ((255) & 63))}
                                                                        _ = v58
                                                                    }
                                                                    switch "" {
                                                                    case "hello":
                                                                        var v59 int16 = 471
                                                                        _ = v59
                                                                    default:
                                                                        fmt.Println(v1)
                                                                    }
                                                                }
                                                                v60 := GenS6{F0: 127, F1: "foo bar", F2: 0}
                                                                v10 = 1
                                                                v61 := GenS7{F0: 42, F1: 547}
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                        }
                                                        switch (fi37 % ((32767) | 1)) {
                                                        case -2147483648:
                                                            var v62 uint = 566
                                                            _ = v62
                                                        case 2147483647:
                                                            v3 = append(v3, (fi9 ^ 2147483647))
                                                        case -1:
                                                            v63 := GenS4{F0: -1, F1: 42}
                                                            _ = v63
                                                        }
                                                        if v1 {
                                                            v10 = v10
                                                            var v64 uint16 = v10
                                                            if (!v1) {
                                                                v64 = (v10 | 674)
                                                                v65 := GenS1{F0: -2147483648}
                                                                delete(v2, ("hello" + ("a" + (("foo bar" + "\n\t") + "hello"))))
                                                                v1 = (-124 != (fi9 << ((823) & 63)))
                                                                v66 := GenS5{F0: 255, F1: 9223372036854775807, F2: -2147483648}
                                                                switch "a" {
                                                                case "a":
                                                                    _ = len(v2)
                                                                }
                                                                _ = v65
                                                                _ = v66
                                                            } else {
                                                                v67 := GenS1{F0: -1}
                                                                _ = v67
                                                            }
                                                            _ = v64
                                                        } else {
                                                            var v68 float32 = (-2.25 * (-2.25 + 1e10))
                                                            switch ("" + "a") {
                                                            case "\n\t":
                                                                if false {
                                                                    v69 := v2["\n\t"]
                                                                    var v70 uint = genFn0()
                                                                    if v1 {
                                                                        v70 = genFn0()
                                                                        var v71 int8 = 100
                                                                        var v72 uint32 = 4294967295
                                                                        _ = v71
                                                                        _ = v72
                                                                    }
                                                                    var v73 float32 = v68
                                                                    v1 = (v1 && false)
                                                                    if v69 {
                                                                        v10 = 403
                                                                        v68 += (1e10 - v73)
                                                                    }
                                                                    v70 += v70
                                                                    v1 = true
                                                                    _ = v69
                                                                    _ = v70
                                                                    _ = v73
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        v74 := GenS8{F0: 1, F1: fi9}
                                                                        _ = v74
                                                                    case "foo bar":
                                                                        var v75 uint32 = 4294967295
                                                                        _ = v75
                                                                    case "":
                                                                        v76 := GenS9{F0: true}
                                                                        _ = v76
                                                                    }
                                                                    switch "a" {
                                                                    case "hello":
                                                                        v1 = (false && true)
                                                                    case "foo bar":
                                                                        v3 = append(v3, (fi9 >> ((1) & 63)))
                                                                    case "":
                                                                        if (false && (v1 || v1)) {
                                                                            var v77 int16 = -32768
                                                                            _ = v77
                                                                        } else {
                                                                            var v78 uint64 = 405
                                                                            var v79 int64 = -9223372036854775808
                                                                            v2["hello"] = (!(!true))
                                                                            _ = v78
                                                                            _ = v79
                                                                        }
                                                                    }
                                                                }
                                                            case "a":
                                                                v1 = (!v1)
                                                            default:
                                                                if (v1 && v1) {
                                                                    switch ("hello" + "hello") {
                                                                    case "":
                                                                        var v80 int8 = 127
                                                                        _ = v80
                                                                    case "a":
                                                                        var v81 bool = (-9223372036854775808 <= -1)
                                                                        _ = v81
                                                                    }
                                                                    var v82 uint = 2
                                                                    _ = v82
                                                                }
                                                            }
                                                            v83 := GenS6{F0: fi9, F1: "a", F2: 2}
                                                            v10 += 32767
                                                            _ = v68
                                                            _ = v83
                                                        }
                                                        var v84 float32 = 1e10
                                                        _ = v84
                                                    }
                                                } else {
                                                    var v85 int = fi9
                                                    var v86 uint8 = 255
                                                    v87 := GenS8{F0: 196, F1: (^(fi9))}
                                                    v85 += v85
                                                    _ = v85
                                                    _ = v86
                                                    _ = v87
                                                }
                                            }
                                        }
                                    default:
                                        fmt.Println(fi9)
                                    }
                                case "\n\t":
                                    var v88 int = 1
                                    _ = v88
                                default:
                                    switch (fi9 / (((fi37 / ((fi9) | 1))) | 1)) {
                                    case 42:
                                        v89 := v2[("a" + ((("hello" + (("hello" + "hello") + "\n\t")) + "\n\t") + (("a" + "") + ("foo bar" + "foo bar"))))]
                                        _ = v89
                                    default:
                                        v90 := GenS10{F0: fi37, F1: 42}
                                        _ = v90
                                    }
                                }
                                v1 = v1
                            }
                        default:
                            v91 := GenS11{F0: v1, F1: 42}
                            _ = v91
                        }
                        v3 = append(v3, (-(fi9)))
                        v10 += (v10 | v10)
                        fmt.Println(fi9, fi37)
                    }
                    _ = fi37
                }
            }
            v2[(("" + "hello") + "a")] = ((v10 << ((0) & 63)) >= v10)
            _ = v10
        } else {
            for fi92 := 0; fi92 < 14; fi92++ {
                switch (("" + "foo bar") + (("hello" + "hello") + ((("" + "\n\t") + ("\n\t" + "\n\t")) + "foo bar"))) {
                case "foo bar":
                    var v93 string = "\n\t"
                    _ = v93
                }
                var v94 uint64 = 4294967295
                fmt.Println(fi9, fi92, v1)
                _ = fi92
                _ = v94
            }
            if ("" == "a") {
                if true {
                    delete(v2, ((("hello" + (("hello" + "") + (("hello" + "a") + ("" + "hello")))) + "\n\t") + "hello"))
                    v95 := GenS10{F0: (fi9 ^ (-(fi9))), F1: 1}
                    _ = v95
                } else {
                    v96 := GenS12{F0: 748, F1: 1, F2: v1}
                    _ = v96
                }
            } else {
                v1 = v1
            }
            for fi97 := 0; fi97 < 17; fi97++ {
                v1 = ((((("foo bar" + "a") + "foo bar") + "") + "a") >= "")
                fmt.Println(fi9)
                _ = fi97
            }
        }
        if v1 {
            fmt.Println(v1, v1)
        }
        _ = fi9
    }
    v98 := GenS7{F0: 95, F1: 845}
    delete(v2, "\n\t")
    v3 = append(v3, 4294967295)
    switch "hello" {
    case "\n\t":
        v1 = (!v1)
    default:
        for fi99 := 0; fi99 < 5; fi99++ {
            var v100 int32 = -1
            _ = fi99
            _ = v100
        }
    }
    switch 279 {
    case 4294967295:
        v3 = append(v3, -899)
    default:
        if v1 {
            v1 = (false || (!(v1 && (!(0 != 902)))))
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
        }
    }
    switch ("a" + "hello") {
    case "hello":
        for fi101 := 0; fi101 < 2; fi101++ {
            var v102 bool = v1
            var v103 int16 = 1000
            var v104 uint8 = 127
            _ = fi101
            _ = v102
            _ = v103
            _ = v104
        }
    }
    switch -9223372036854775808 {
    case 4294967295:
        if false {
            var v105 uint8 = 82
            var v106 uint16 = 42
            _ = v105
            _ = v106
        }
    default:
        switch 128 {
        case 4294967295:
            v1 = false
        default:
            v1 = (!v1)
        }
    }
    var v107 float64 = 0.1
    v3 = append(v3, 32767)
    v107 = (v107 / 1.5)
    v107 = ((v107 + 3.14159) + v107)
    switch "" {
    case "a":
        for fi108 := 0; fi108 < 4; fi108++ {
            var v109 uint64 = 1
            v1 = v1
            _ = fi108
            _ = v109
        }
    case "":
        v110 := GenS13{F0: 1000, F1: 42}
        _ = v110
    case "hello":
        fmt.Println(v107)
    default:
        if ((1 == 502) || (!(255 >= 536))) {
            if ("\n\t" < "a") {
                if len(v3) > 0 {
                    v3[0] = 0
                }
            }
            switch 32768 {
            case -82:
                if (!(0 >= 0)) {
                    v2["\n\t"] = true
                    v107 = v107
                    for fi111 := 0; fi111 < 2; fi111++ {
                        var v112 bool = (v1 || (v107 < (v107 - 0.0)))
                        var v113 uint = 65535
                        var v114 int16 = 32767
                        _ = fi111
                        _ = v112
                        _ = v113
                        _ = v114
                    }
                }
            case 42:
                var v115 uint8 = 83
                _ = v115
            case 2147483647:
                var v116 int = 579
                _ = v116
            }
            v117 := v2["a"]
            _ = v117
        }
    }
    for fi118 := 0; fi118 < 11; fi118++ {
        v119 := GenS11{F0: (18446744073709551615 >= 2), F1: -128}
        var v120 float32 = float32(float64(1e10))
        v121 := GenS14{F0: -764}
        _ = fi118
        _ = v119
        _ = v120
        _ = v121
    }
    if (-9223372036854775808 >= 560) {
        var v122 uint16 = 643
        if true {
            var v123 int = 1
            v124 := v2["hello"]
            var v125 float32 = 1.5
            _ = v123
            _ = v124
            _ = v125
        }
        for fi126 := 0; fi126 < 19; fi126++ {
            if v1 {
                v122 -= (v122 >> ((65535) & 63))
            } else {
                var v127 bool = ((100 < 1) || v1)
                for fi128 := 0; fi128 < 5; fi128++ {
                    v107 = float64(float32(-2.25))
                    fmt.Println(v127)
                    _ = fi128
                }
                _ = v127
            }
            v107 = (((v107 * (float64((1.5 / (-2.25 / -2.25))) / v107)) - ((-2.25 + 3.14159) / v107)) - v107)
            fmt.Println(fi126, fi126, v107)
            _ = fi126
        }
        _ = v122
    } else {
        if len(v3) > 0 {
            v3[0] = -9223372036854775808
        }
        var v129 int = 2147483647
        for fi130 := 0; fi130 < 3; fi130++ {
            var v131 uint8 = 44
            v132 := GenS15{F0: 0, F1: 65535, F2: (-(fi130))}
            _ = fi130
            _ = v131
            _ = v132
        }
        _ = v129
    }
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v107
    _ = v2
    _ = v3
    _ = v98
}
