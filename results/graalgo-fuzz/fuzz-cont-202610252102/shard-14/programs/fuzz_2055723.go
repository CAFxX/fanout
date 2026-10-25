package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 string }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 uint8; F1 int8; F2 bool }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 int; F1 int32 }
type GenS11 struct { F0 int32; F1 uint8; F2 int16 }
type GenS12 struct { F0 int8; F1 bool; F2 uint16 }
type GenS13 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS14 struct { F0 bool; F1 int32; F2 uint64 }
type GenS15 struct { F0 bool; F1 int16 }
type GenS16 struct { F0 uint; F1 int8 }
type GenS17 struct { F0 int64; F1 int8 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 int8; F1 uint16; F2 uint32 }
func genFn0() int64 {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float32 = ((1.5 / -2.25) * 1e10)
        var v8 uint16 = 42
        for fi9 := 0; fi9 < 18; fi9++ {
            switch fi6 {
            case 127:
                v10 := GenS1{F0: -211, F1: 1, F2: "hello"}
                _ = v10
            case 42:
                var v11 int64 = -9223372036854775808
                _ = v11
            case 163:
                v8 = v8
            }
            switch ("hello" + ("foo bar" + "a")) {
            case "a":
                v7 = v7
            case "hello":
                if true {
                    v12 := GenS1{F0: 2147483647, F1: 42, F2: "a"}
                    _ = v12
                } else {
                    v7 = v7
                    var v13 uint = 18446744073709551615
                    v14 := GenS2{F0: 1, F1: 18446744073709551615}
                    fmt.Println(fi9, fi6)
                    _ = v13
                    _ = v14
                }
            default:
                v15 := []int{fi9, 9223372036854775807, (fi6 * (fi9 - (fi9 << ((831) & 63))))}
                _ = v15
            }
            if (false || (0.1 > -2.25)) {
                var v16 bool = (42 != -1)
                var v17 int64 = 14
                switch fi9 {
                case 650:
                    var v18 float64 = -2.25
                    _ = v18
                }
                switch 2147483647 {
                case -9223372036854775808:
                    var v19 int = (fi9 * 128)
                    _ = v19
                case 4294967295:
                    switch 9223372036854775807 {
                    case 128:
                        if (((v7 + v7) - v7) < v7) {
                            v16 = (float64((-2.25 - ((float32(((0.1 - -2.25) * 0.1)) + ((3.14159 * v7) + v7)) - ((v7 / 1e10) / ((0.1 - (float32((0.1 + 1.5)) - 3.14159)) / v7))))) == (-2.25 + 1.5))
                            var v20 string = "hello"
                            switch v20 {
                            case "hello":
                                var v21 uint8 = 0
                                _ = v21
                            case "":
                                if (48 <= 0) {
                                    v16 = v16
                                    v17 = int64(fi9)
                                } else {
                                    v16 = v16
                                    v20 = ""
                                    var v22 float32 = ((v7 - v7) * 0.1)
                                    _ = v22
                                }
                            case "a":
                                if (true && false) {
                                    if v16 {
                                        if (v16 && v16) {
                                            v7 += v7
                                        } else {
                                            v23 := []int{fi9, fi9}
                                            _ = v23
                                        }
                                    }
                                    var v24 uint64 = 514
                                    v25 := []int{fi6}
                                    _ = v24
                                    _ = v25
                                } else {
                                    v26 := []int{}
                                    v20 = ("" + (v20 + ""))
                                    var v27 uint8 = 146
                                    _ = v26
                                    _ = v27
                                }
                            default:
                                if v16 {
                                    fmt.Println(fi9, v16, v7)
                                }
                            }
                            _ = v20
                        }
                    }
                case 2147483647:
                    fmt.Println(v17)
                default:
                    v28 := GenS2{F0: 0, F1: 18446744073709551615}
                    var v29 string = "a"
                    _ = v28
                    _ = v29
                }
                _ = v16
                _ = v17
            }
            _ = fi9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v30 uint64 = 1
    _ = v30
    return 680
}
func genFn1() int {
    if (true || true) {
        var v31 int = 298
        if false {
            var v32 float64 = 3.14159
            _ = v32
        } else {
            v33 := GenS2{F0: 18446744073709551615, F1: 2}
            for fi34 := 0; fi34 < 12; fi34++ {
                v35 := GenS2{F0: 965, F1: 1}
                _ = fi34
                _ = v35
            }
            _ = v33
        }
        _ = v31
    }
    v36 := []int{-29, -2147483648, 4294967295, -1}
    var v37 int64 = 1
    switch int(v37) {
    case 0:
        v37 -= (-(v37))
    case 42:
        v37 -= 0
    default:
        v37 = (v37 % ((genFn0()) | 1))
    }
    switch -318 {
    case 0:
        var v38 uint16 = 255
        _ = v38
    case -1:
        v37 += v37
    }
    _ = v36
    _ = v37
    return 351
}

func main() {
    var v1 uint16 = 2
    var v2 int = 128
    var v3 int = v2
    var v4 string = ((("a" + "") + ("hello" + "hello")) + "foo bar")
    var v5 uint = 2
    v1 += v1
    var v39 int16 = -235
    var v40 float32 = 0.0
    for fi41 := 0; fi41 < 13; fi41++ {
        var v42 int64 = -504
        _ = fi41
        _ = v42
    }
    var v43 uint8 = 127
    fmt.Println(v3, v43, v43)
    v44 := []int{}
    switch (v3 << ((v5) & 63)) {
    case 2147483647:
        if true {
            for fi45 := 0; fi45 < 10; fi45++ {
                fmt.Println(v39, v2)
                for fi46 := 0; fi46 < 20; fi46++ {
                    var v47 uint64 = 1
                    _ = fi46
                    _ = v47
                }
                _ = fi45
            }
            switch 4294967295 {
            case -864:
                var v48 string = v4
                _ = v48
            case -286:
                var v49 uint64 = 540
                _ = v49
            }
            v50 := GenS3{F0: 42, F1: 1}
            _ = v50
        } else {
            switch (v2 - v2) {
            case -224:
                if true {
                    var v51 int32 = -2147483648
                    for fi52 := 0; fi52 < 18; fi52++ {
                        var v53 string = (("foo bar" + v4) + "\n\t")
                        v54 := GenS3{F0: (v51 * v51), F1: 748}
                        if (false && false) {
                            v40 = v40
                            fmt.Println(v53, v4, fi52)
                        } else {
                            for fi55 := 0; fi55 < 13; fi55++ {
                                var v56 uint16 = v1
                                v3 += 9223372036854775807
                                if false {
                                    var v57 string = "foo bar"
                                    if (1 != 100) {
                                        _ = len(v44)
                                    }
                                    _ = v57
                                }
                                _ = fi55
                                _ = v56
                            }
                            v40 += float32(0.1)
                        }
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                    v1 = (v1 + v1)
                    _ = v51
                }
            case 694:
                fmt.Println(v4, v1, v43)
            }
            var v58 string = v4
            var v59 uint16 = v1
            _ = v58
            _ = v59
        }
    case 128:
        var v60 bool = true
        _ = v60
    }
    v40 = 0.1
    switch (v3 + v2) {
    case 0:
        if true {
            switch ("hello" + v4) {
            case "foo bar":
                for fi61 := 0; fi61 < 19; fi61++ {
                    v3 = genFn1()
                    var v62 uint = (v5 | (v5 / ((v5) | 1)))
                    var v63 int8 = 1
                    _ = fi61
                    _ = v62
                    _ = v63
                }
            case "hello":
                switch (^(v2)) {
                case -2147483648:
                    v40 -= (3.14159 + (0.0 / ((1e10 / 3.14159) * v40)))
                default:
                    if (-575 > v39) {
                        if true {
                            switch v3 {
                            case 4294967295:
                                var v64 uint8 = (v43 + v43)
                                _ = v64
                            case 32767:
                                var v65 float32 = 3.14159
                                _ = v65
                            default:
                                if false {
                                    switch "\n\t" {
                                    case "a":
                                        var v66 uint32 = 65535
                                        _ = v66
                                    default:
                                        for fi67 := 0; fi67 < 6; fi67++ {
                                            for fi68 := 0; fi68 < 1; fi68++ {
                                                if (true && (989 > 42)) {
                                                    var v69 float32 = v40
                                                    _ = v69
                                                } else {
                                                    switch 32767 {
                                                    case 894:
                                                        if len(v44) > 0 {
                                                            v44[0] = fi67
                                                        }
                                                    case 0:
                                                        v4 = (("" + ((v4 + v4) + "hello")) + "")
                                                    case 127:
                                                        if len(v44) >= 2 {
                                                            v44 = v44[1:len(v44)]
                                                        }
                                                    default:
                                                        if false {
                                                            v2 += (fi68 - 703)
                                                        } else {
                                                            if false {
                                                                var v70 uint = 18446744073709551615
                                                                _ = v70
                                                            } else {
                                                                var v71 float64 = -2.25
                                                                _ = v71
                                                            }
                                                            var v72 int16 = v39
                                                            switch "hello" {
                                                            case "\n\t":
                                                                if (false && (false && (-1 <= -2147483648))) {
                                                                    var v73 int8 = 0
                                                                    _ = v73
                                                                } else {
                                                                    switch v4 {
                                                                    case "foo bar":
                                                                        _ = len(v44)
                                                                    case "\n\t":
                                                                        if true {
                                                                            fmt.Println(fi68)
                                                                        } else {
                                                                            var v74 int32 = 0
                                                                            _ = v74
                                                                        }
                                                                    case "":
                                                                        var v75 int = 1
                                                                        _ = v75
                                                                    }
                                                                    var v76 string = v4
                                                                    if (333 <= genFn0()) {
                                                                        v40 += float32((float64((-2.25 + 3.14159)) + float64(float32(1.5))))
                                                                        v77 := GenS4{F0: (v72 | 0)}
                                                                        _ = len(v44)
                                                                        _ = v77
                                                                    } else {
                                                                        if (true && true) {
                                                                            v40 = v40
                                                                            switch v3 {
                                                                            case 265:
                                                                                var v78 uint32 = uint32(v1)
                                                                                _ = v78
                                                                            }
                                                                            v3 = (v3 - (-(fi67)))
                                                                        } else {
                                                                            fmt.Println(v2, v2)
                                                                        }
                                                                        switch ((v4 + "\n\t") + (v76 + ("\n\t" + ((((v76 + (v4 + v4)) + "a") + "\n\t") + v76)))) {
                                                                        case "a":
                                                                            v79 := GenS5{F0: v5, F1: v2}
                                                                            _ = v79
                                                                        default:
                                                                            if true {
                                                                                fmt.Println(v4)
                                                                            } else {
                                                                                var v80 int8 = -26
                                                                                _ = v80
                                                                            }
                                                                        }
                                                                        var v81 uint16 = 1
                                                                        _ = len(v44)
                                                                        switch fi67 {
                                                                        case -482:
                                                                            v40 = v40
                                                                        case -771:
                                                                            v82 := GenS5{F0: v5, F1: 9223372036854775807}
                                                                            _ = v82
                                                                        case 0:
                                                                            if false {
                                                                                switch (fi68 * -876) {
                                                                                case -2147483648:
                                                                                    v40 = (v40 - 0.0)
                                                                                case 0:
                                                                                    var v83 int8 = 1
                                                                                    _ = v83
                                                                                }
                                                                                v84 := GenS6{F0: (v43 & (v43 - v43)), F1: 1}
                                                                                _ = v84
                                                                            } else {
                                                                                switch int(v43) {
                                                                                case 128:
                                                                                    v44 = append(v44, (v2 | fi68))
                                                                                case -644:
                                                                                    var v85 bool = false
                                                                                    _ = v85
                                                                                case 121:
                                                                                    v86 := GenS6{F0: v43, F1: 9223372036854775807}
                                                                                    switch (v3 & fi67) {
                                                                                    case -355:
                                                                                        switch (fi68 ^ fi67) {
                                                                                        case 128:
                                                                                            var v87 int64 = -215
                                                                                            _ = v87
                                                                                        default:
                                                                                            v88 := GenS6{F0: 60, F1: 9223372036854775807}
                                                                                            _ = v88
                                                                                        }
                                                                                    default:
                                                                                        switch v2 {
                                                                                        case 994:
                                                                                            switch (((v4 + (v76 + ("" + ((v76 + (v4 + v4)) + (("foo bar" + (v4 + ("foo bar" + "a"))) + v76))))) + "hello") + v4) {
                                                                                            case "a":
                                                                                                v2 = v2
                                                                                            case "":
                                                                                                v4 = v76
                                                                                            }
                                                                                        case -9223372036854775808:
                                                                                            if ((v43 + 127) >= (v43 << ((759) & 63))) {
                                                                                                switch 128 {
                                                                                                case 128:
                                                                                                    var v89 uint32 = 65535
                                                                                                    _ = v89
                                                                                                case -834:
                                                                                                    v90 := GenS7{F0: 127, F1: 100, F2: true}
                                                                                                    _ = v90
                                                                                                default:
                                                                                                    v4 = v76
                                                                                                }
                                                                                            } else {
                                                                                                if (((v4 + "a") + v76) == (v76 + v76)) {
                                                                                                    v91 := GenS3{F0: 293, F1: 2}
                                                                                                    _ = v91
                                                                                                } else {
                                                                                                    var v92 int8 = -1
                                                                                                    if len(v44) > 0 {
                                                                                                        v44[0] = (v2 >> ((v5) & 63))
                                                                                                    }
                                                                                                    _ = v92
                                                                                                }
                                                                                                if (1000 < v72) {
                                                                                                    var v93 int = v2
                                                                                                    _ = len(v44)
                                                                                                    if len(v44) > 0 {
                                                                                                        v44[0] = -2147483648
                                                                                                    }
                                                                                                    switch (fi67 * (fi68 >> (((^(v5))) & 63))) {
                                                                                                    case 9223372036854775807:
                                                                                                        switch (v4 + "a") {
                                                                                                        case "":
                                                                                                            var v94 int8 = int8(v39)
                                                                                                            _ = v94
                                                                                                        case "hello":
                                                                                                            fmt.Println(v72, v81, v43)
                                                                                                        case "a":
                                                                                                            v95 := GenS6{F0: (v43 % ((v43) | 1)), F1: 144}
                                                                                                            _ = v95
                                                                                                        default:
                                                                                                            var v96 float64 = 1e10
                                                                                                            _ = v96
                                                                                                        }
                                                                                                        v44 = append(v44, (fi67 - v3))
                                                                                                        if false {
                                                                                                            v1 = (^(v1))
                                                                                                            if (true && ((2147483648 != uint32(v43)) && ((false && ((v2 != v93) && (2 <= 4294967295))) && false))) {
                                                                                                                v1 -= 32767
                                                                                                                if false {
                                                                                                                    var v97 uint8 = 0
                                                                                                                    _ = v97
                                                                                                                } else {
                                                                                                                    var v98 bool = false
                                                                                                                    _ = v98
                                                                                                                }
                                                                                                            }
                                                                                                            v99 := GenS8{F0: false, F1: (v1 + v81)}
                                                                                                            var v100 string = v4
                                                                                                            _ = v100
                                                                                                            _ = v99
                                                                                                        } else {
                                                                                                            v40 = (3.14159 - v40)
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v93
                                                                                                } else {
                                                                                                    if false {
                                                                                                        v44 = append(v44, 2147483647)
                                                                                                        if len(v44) >= 2 {
                                                                                                            v44 = v44[1:len(v44)]
                                                                                                        }
                                                                                                    } else {
                                                                                                        v5 = v5
                                                                                                    }
                                                                                                    v101 := GenS9{F0: -1, F1: 212}
                                                                                                    v102 := GenS10{F0: fi68, F1: -1}
                                                                                                    _ = v101
                                                                                                    _ = v102
                                                                                                }
                                                                                                var v103 uint8 = v43
                                                                                                _ = v103
                                                                                            }
                                                                                        case 32768:
                                                                                            if ((-(v2)) <= fi68) {
                                                                                                if len(v44) > 0 {
                                                                                                    v44[0] = 4294967295
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v104 int = v2
                                                                                    _ = v104
                                                                                    _ = v86
                                                                                }
                                                                            }
                                                                        default:
                                                                            switch int(v72) {
                                                                            case 128:
                                                                                switch (fi68 ^ (v2 % (((v2 - genFn1())) | 1))) {
                                                                                case 32767:
                                                                                    fmt.Println(v43, v39, v5)
                                                                                }
                                                                            case 393:
                                                                                switch (v76 + v4) {
                                                                                case "":
                                                                                    v1 = 255
                                                                                    var v105 uint64 = 2
                                                                                    _ = v105
                                                                                case "foo bar":
                                                                                    v2 -= 4294967295
                                                                                default:
                                                                                    switch fi67 {
                                                                                    case 9223372036854775807:
                                                                                        switch v2 {
                                                                                        case -1:
                                                                                            v106 := GenS1{F0: -2147483648, F1: v39, F2: v76}
                                                                                            _ = v106
                                                                                        case 621:
                                                                                            if true {
                                                                                                fmt.Println(v4, v1, v3)
                                                                                                if (!true) {
                                                                                                    v107 := GenS4{F0: v39}
                                                                                                    _ = v107
                                                                                                } else {
                                                                                                    v44 = append(v44, v3)
                                                                                                    var v108 uint8 = (^(v43))
                                                                                                    _ = v108
                                                                                                }
                                                                                                var v109 uint64 = 2
                                                                                                _ = v109
                                                                                            } else {
                                                                                                var v110 int = 9223372036854775807
                                                                                                _ = v110
                                                                                            }
                                                                                        case 339:
                                                                                            if false {
                                                                                                if false {
                                                                                                    v111 := GenS11{F0: 689, F1: v43, F2: 32767}
                                                                                                    if (!((v72 & v39) <= (v39 | v72))) {
                                                                                                        fmt.Println(v76)
                                                                                                        if len(v44) >= 2 {
                                                                                                            v44 = v44[1:len(v44)]
                                                                                                        }
                                                                                                        switch 261 {
                                                                                                        case -9223372036854775808:
                                                                                                            v112 := GenS11{F0: 1, F1: (v43 - (v43 | (v43 + (v43 << (((v5 << ((v5) & 63))) & 63))))), F2: v39}
                                                                                                            _ = v112
                                                                                                        case 32768:
                                                                                                            v43 = v43
                                                                                                        case 32767:
                                                                                                            v1 -= 2
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v76)
                                                                                                    v72 -= (^(v72))
                                                                                                    fmt.Println(v4, v39)
                                                                                                    _ = v111
                                                                                                } else {
                                                                                                    switch int(v43) {
                                                                                                    case -733:
                                                                                                        if len(v44) > 0 {
                                                                                                            v44[0] = 32768
                                                                                                        }
                                                                                                    }
                                                                                                    v113 := GenS3{F0: 431, F1: 4294967295}
                                                                                                    _ = v113
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            _ = len(v44)
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 341:
                                                                                v114 := GenS12{F0: 1, F1: true, F2: (v1 << ((v5) & 63))}
                                                                                _ = v114
                                                                            default:
                                                                                var v115 int64 = 0
                                                                                _ = v115
                                                                            }
                                                                        }
                                                                        if false {
                                                                            v116 := GenS13{F0: 1, F1: 1, F2: 576}
                                                                            var v117 int32 = 0
                                                                            _ = v116
                                                                            _ = v117
                                                                        } else {
                                                                            if false {
                                                                                var v118 int64 = genFn0()
                                                                                _ = v118
                                                                            }
                                                                        }
                                                                        v119 := GenS3{F0: 2147483647, F1: 0}
                                                                        v76 = v4
                                                                        _ = v119
                                                                        _ = v81
                                                                    }
                                                                    var v120 string = v76
                                                                    _ = v120
                                                                    _ = v76
                                                                }
                                                            case "foo bar":
                                                                if (!true) {
                                                                    if (!false) {
                                                                        v72 += (v39 & 1000)
                                                                        if false {
                                                                            _ = len(v44)
                                                                        } else {
                                                                            var v121 bool = (v40 <= (v40 / v40))
                                                                            _ = v121
                                                                        }
                                                                    } else {
                                                                        var v122 int8 = 122
                                                                        switch (v2 & (v2 >> ((v5) & 63))) {
                                                                        case 127:
                                                                            v123 := GenS14{F0: false, F1: 2147483647, F2: 16}
                                                                            _ = v123
                                                                        }
                                                                        _ = v122
                                                                    }
                                                                    var v124 string = (v4 + v4)
                                                                    _ = v124
                                                                }
                                                            case "hello":
                                                                var v125 bool = (v4 == v4)
                                                                _ = v125
                                                            default:
                                                                if (32768 <= v3) {
                                                                    var v126 float32 = v40
                                                                    v40 -= ((1e10 + ((v40 - 0.0) + (v40 * (v126 - -2.25)))) - v40)
                                                                    v1 -= (v1 % ((uint16(v5)) | 1))
                                                                    _ = v126
                                                                }
                                                            }
                                                            _ = v72
                                                        }
                                                    }
                                                    var v127 uint32 = 0
                                                    var v128 int = fi67
                                                    _ = v127
                                                    _ = v128
                                                }
                                                _ = len(v44)
                                                _ = fi68
                                            }
                                            if len(v44) > 0 {
                                                v44[0] = (fi67 % ((fi67) | 1))
                                            }
                                            _ = fi67
                                        }
                                    }
                                    fmt.Println(v3, v3, v3)
                                }
                            }
                            for fi129 := 0; fi129 < 6; fi129++ {
                                v2 -= genFn1()
                                _ = fi129
                            }
                            _ = len(v44)
                        }
                        if false {
                            var v130 int8 = 42
                            for fi131 := 0; fi131 < 6; fi131++ {
                                var v132 uint64 = 42
                                var v133 int32 = 832
                                _ = fi131
                                _ = v132
                                _ = v133
                            }
                            v5 += v5
                            _ = v130
                        }
                    } else {
                        switch v2 {
                        case -755:
                            _ = len(v44)
                        case 32767:
                            for fi134 := 0; fi134 < 12; fi134++ {
                                if false {
                                    var v135 uint64 = 2
                                    _ = v135
                                } else {
                                    v136 := GenS2{F0: 2, F1: uint64(v5)}
                                    fmt.Println(v2, v4)
                                    v137 := GenS13{F0: 296, F1: 0, F2: uint64(v5)}
                                    _ = v136
                                    _ = v137
                                }
                                switch 9223372036854775807 {
                                case 2147483647:
                                    v2 = (v3 - v2)
                                case 1:
                                    fmt.Println(v1, v4, v43)
                                case 42:
                                    v4 = ("\n\t" + v4)
                                default:
                                    for fi138 := 0; fi138 < 15; fi138++ {
                                        var v139 uint64 = 1
                                        var v140 uint32 = 1
                                        _ = fi138
                                        _ = v139
                                        _ = v140
                                    }
                                }
                                _ = fi134
                            }
                        case 1:
                            switch v4 {
                            case "\n\t":
                                var v141 int32 = -1
                                _ = v141
                            case "":
                                for fi142 := 0; fi142 < 13; fi142++ {
                                    if true {
                                        switch ((v4 + (((v4 + "\n\t") + v4) + ("\n\t" + "a"))) + (v4 + "")) {
                                        case "":
                                            var v143 uint32 = 260
                                            _ = v143
                                        default:
                                            if (!true) {
                                                if false {
                                                    _ = len(v44)
                                                } else {
                                                    var v144 uint64 = 4294967295
                                                    v3 = (fi142 * (v3 / ((genFn1()) | 1)))
                                                    switch v3 {
                                                    case 128:
                                                        if false {
                                                            if true {
                                                                v4 = v4
                                                            } else {
                                                                if (410 > (^(v5))) {
                                                                    v145 := GenS1{F0: 2147483647, F1: (-(v39)), F2: ""}
                                                                    _ = v145
                                                                } else {
                                                                    var v146 uint64 = (v144 - (v144 >> ((v5) & 63)))
                                                                    v147 := GenS10{F0: v3, F1: -2147483648}
                                                                    for fi148 := 0; fi148 < 9; fi148++ {
                                                                        v149 := GenS15{F0: (1 < 561), F1: (^(v39))}
                                                                        _ = fi148
                                                                        _ = v149
                                                                    }
                                                                    _ = v146
                                                                    _ = v147
                                                                }
                                                                if (612 <= (v1 | (^(v1)))) {
                                                                    v4 = v4
                                                                    var v150 int64 = 139
                                                                    _ = v150
                                                                } else {
                                                                    if len(v44) > 0 {
                                                                        v44[0] = genFn1()
                                                                    }
                                                                    _ = len(v44)
                                                                }
                                                            }
                                                            var v151 int = (v3 & fi142)
                                                            switch v4 {
                                                            case "foo bar":
                                                                v5 = (v5 + (v5 >> ((1) & 63)))
                                                            case "\n\t":
                                                                var v152 uint64 = 18446744073709551615
                                                                _ = v152
                                                            case "hello":
                                                                switch 32767 {
                                                                case -823:
                                                                    fmt.Println(v4, v40, v4)
                                                                case 32768:
                                                                    v40 -= float32((float64(0.0) * (0.1 - float64(float32(0.1)))))
                                                                case -2147483648:
                                                                    var v153 float32 = v40
                                                                    _ = v153
                                                                default:
                                                                    v154 := GenS16{F0: 255, F1: 127}
                                                                    _ = v154
                                                                }
                                                            }
                                                            _ = v151
                                                        }
                                                    default:
                                                        for fi155 := 0; fi155 < 5; fi155++ {
                                                            var v156 int32 = -24
                                                            v40 = (v40 + (v40 - ((v40 - -2.25) * v40)))
                                                            v157 := GenS1{F0: v156, F1: v39, F2: v4}
                                                            _ = fi155
                                                            _ = v156
                                                            _ = v157
                                                        }
                                                    }
                                                    _ = v144
                                                }
                                                var v158 string = ((v4 + "\n\t") + v4)
                                                _ = v158
                                            } else {
                                                for fi159 := 0; fi159 < 19; fi159++ {
                                                    if false {
                                                        var v160 uint = v5
                                                        _ = v160
                                                    } else {
                                                        v43 += uint8(v5)
                                                        var v161 int32 = -978
                                                        _ = v161
                                                    }
                                                    _ = len(v44)
                                                    _ = fi159
                                                }
                                                switch (-(fi142)) {
                                                case 981:
                                                    v39 = (v39 >> (((v5 % ((v5) | 1))) & 63))
                                                case 127:
                                                    _ = len(v44)
                                                case 2147483647:
                                                    var v162 int32 = 0
                                                    _ = v162
                                                default:
                                                    if true {
                                                        if (947 < -1) {
                                                            v5 = 959
                                                            fmt.Println(v43, v2)
                                                        } else {
                                                            fmt.Println(v2, v1)
                                                        }
                                                        if len(v44) >= 2 {
                                                            v44 = v44[1:len(v44)]
                                                        }
                                                    }
                                                }
                                                v3 += 0
                                            }
                                        }
                                        switch (v4 + "\n\t") {
                                        case "hello":
                                            if len(v44) >= 2 {
                                                v44 = v44[1:len(v44)]
                                            }
                                        case "\n\t":
                                            if true {
                                                if false {
                                                    _ = len(v44)
                                                }
                                                v163 := GenS10{F0: (v3 % (((v2 << (((v5 * v5)) & 63))) | 1)), F1: 1}
                                                _ = v163
                                            } else {
                                                var v164 int = v3
                                                if len(v44) >= 2 {
                                                    v44 = v44[1:len(v44)]
                                                }
                                                _ = len(v44)
                                                _ = v164
                                            }
                                        default:
                                            for fi165 := 0; fi165 < 7; fi165++ {
                                                v166 := GenS7{F0: v43, F1: -76, F2: false}
                                                var v167 float64 = -2.25
                                                v168 := GenS8{F0: (1000 != (-(v39))), F1: v1}
                                                _ = fi165
                                                _ = v166
                                                _ = v167
                                                _ = v168
                                            }
                                        }
                                    } else {
                                        if false {
                                            v169 := GenS15{F0: true, F1: 1}
                                            var v170 bool = true
                                            _ = v169
                                            _ = v170
                                        }
                                    }
                                    v1 = v1
                                    _ = fi142
                                }
                            case "a":
                                for fi171 := 0; fi171 < 1; fi171++ {
                                    if (true || false) {
                                        fmt.Println(v2, v5, v4)
                                        if len(v44) >= 2 {
                                            v44 = v44[1:len(v44)]
                                        }
                                    } else {
                                        v172 := GenS17{F0: genFn0(), F1: 0}
                                        _ = v172
                                    }
                                    var v173 uint16 = v1
                                    if true {
                                        fmt.Println(fi171)
                                    } else {
                                        v4 = ("foo bar" + v4)
                                        var v174 int = (v3 ^ (^(v3)))
                                        if true {
                                            v4 = v4
                                        }
                                        _ = v174
                                    }
                                    _ = fi171
                                    _ = v173
                                }
                            default:
                                v4 = "foo bar"
                            }
                        default:
                            var v175 uint = 1
                            _ = v175
                        }
                        if false {
                            var v176 uint8 = (v43 >> ((65535) & 63))
                            _ = v176
                        }
                        v44 = append(v44, (v2 ^ genFn1()))
                    }
                }
            default:
                v177 := GenS18{F0: v4}
                _ = v177
            }
            fmt.Println(v40, v4)
        } else {
            var v178 int64 = genFn0()
            for fi179 := 0; fi179 < 4; fi179++ {
                switch "a" {
                case "hello":
                    v178 = (v178 >> ((uint(v1)) & 63))
                case "a":
                    fmt.Println(v3, v5, v4)
                case "":
                    v180 := GenS19{F0: -128, F1: v1, F2: 4294967295}
                    _ = v180
                default:
                    var v181 float32 = v40
                    _ = v181
                }
                var v182 uint = (^(v5))
                if (!false) {
                    v183 := GenS9{F0: (v178 / (((v178 % ((-9223372036854775808) | 1))) | 1)), F1: (^(v182))}
                    _ = v183
                }
                _ = fi179
                _ = v182
            }
            _ = v178
        }
    case -1:
        var v184 float64 = 0.0
        _ = v184
    case -907:
        fmt.Println(v39, v40)
    default:
        for fi185 := 0; fi185 < 7; fi185++ {
            v3 -= (fi185 ^ v3)
            v40 -= (((v40 + v40) - ((v40 * (v40 - 0.1)) + 1.5)) + (v40 + (((((v40 - v40) - float32(3.14159)) - (((v40 + v40) + 0.0) - float32(0.0))) + v40) / ((v40 / 1e10) * -2.25))))
            _ = fi185
        }
    }
    var v186 int16 = (-(v39))
    if (666 == 519) {
        for fi187 := 0; fi187 < 4; fi187++ {
            var v188 uint16 = (v1 ^ (v1 - v1))
            _ = fi187
            _ = v188
        }
    } else {
        for fi189 := 0; fi189 < 11; fi189++ {
            var v190 bool = false
            switch -9223372036854775808 {
            case -139:
                if false {
                    var v191 uint16 = v1
                    _ = v191
                }
            case 46:
                _ = len(v44)
            }
            var v192 uint16 = (^(v1))
            _ = fi189
            _ = v190
            _ = v192
        }
    }
    fmt.Println(v3, v39)
    fmt.Println(v186, v43, v3)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v186:", v186)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v186
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v5
}
