package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int8; F1 int32; F2 int16 }
type GenS5 struct { F0 int64; F1 uint; F2 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64; F1 uint; F2 bool }
type GenS10 struct { F0 uint; F1 int8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 bool; F1 bool }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 uint32; F1 int16 }
type GenS17 struct { F0 uint8; F1 uint32 }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 uint8; F1 int8 }
type GenS20 struct { F0 int16; F1 bool; F2 bool }
func genFn0() int8 {
    var v4 int64 = 0
    var v5 int16 = 1000
    v4 = -9223372036854775808
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint16, p1 uint16) int32 {
    var v6 int8 = -128
    v7 := []int{-939, -9223372036854775808, -2147483648}
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int8 = (v6 >> ((0) & 63))
            v7 = append(v7, (^(fi8)))
            var v11 float64 = -2.25
            _ = fi9
            _ = v10
            _ = v11
        }
        if false {
            if false {
                v7 = append(v7, (fi8 >> ((619) & 63)))
                v6 += -44
            } else {
                v12 := GenS1{F0: 2147483647, F1: fi8}
                _ = v12
            }
            if len(v7) > 0 {
                v7[0] = 429
            }
            fmt.Println(v6, p1)
        } else {
            if true {
                p0 += 0
                for fi13 := 0; fi13 < 10; fi13++ {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    _ = fi13
                }
                var v14 int8 = v6
                _ = v14
            }
            if (((float32(0.1) * 0.0) - 0.1) == float32(1e10)) {
                p1 = 2
                v6 = genFn0()
                for fi15 := 0; fi15 < 1; fi15++ {
                    if true {
                        p0 = 848
                    }
                    v6 -= -38
                    _ = fi15
                }
            }
            if false {
                v7 = append(v7, fi8)
                for fi16 := 0; fi16 < 11; fi16++ {
                    fmt.Println(p0, fi8)
                    v17 := GenS1{F0: 2147483647, F1: 1}
                    fmt.Println(fi8)
                    _ = fi16
                    _ = v17
                }
                var v18 float64 = 3.14159
                _ = v18
            } else {
                var v19 string = ""
                fmt.Println(p1, v19, fi8)
                var v20 int16 = -1
                _ = v19
                _ = v20
            }
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{2, 1, 344, 2147483648}
    var v3 uint16 = 2
    fmt.Println(v3)
    switch "foo bar" {
    case "":
        for fi21 := 0; fi21 < 7; fi21++ {
            for fi22 := 0; fi22 < 9; fi22++ {
                v3 += v3
                var v23 float64 = 0.1
                _ = fi22
                _ = v23
            }
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 float32 = 1.5
                if false {
                    if true {
                        if len(v2) > 0 {
                            v2[0] = 2147483647
                        }
                        v3 -= 42
                        fmt.Println(v3)
                    } else {
                        fmt.Println(fi21, v25)
                        v26 := GenS1{F0: 42, F1: -177}
                        fmt.Println(v3, fi21)
                        _ = v26
                    }
                }
                if true {
                    var v27 uint8 = 230
                    v28 := GenS1{F0: 2, F1: fi21}
                    if true {
                        v27 -= 202
                        var v29 bool = true
                        v30 := GenS1{F0: 2147483648, F1: (-(fi21))}
                        if false {
                            fmt.Println(v29)
                        }
                        _ = v29
                        _ = v30
                    }
                    if (883 < 136) {
                        v27 = 164
                        v3 += (v3 << ((2) & 63))
                        switch (fi24 & (fi24 % ((14) | 1))) {
                        case 0:
                            if (v27 <= (v27 ^ v27)) {
                                if true {
                                    v27 -= 1
                                    switch 9223372036854775807 {
                                    case -160:
                                        v27 += 91
                                        switch 9223372036854775807 {
                                        case -9223372036854775808:
                                            fmt.Println(v25, v27, v25)
                                        }
                                        fmt.Println(v27, v3)
                                    case 0:
                                        var v31 uint16 = 48
                                        _ = v31
                                    }
                                } else {
                                    var v32 float32 = v25
                                    v33 := GenS2{F0: 42}
                                    _ = v32
                                    _ = v33
                                }
                            }
                        case 230:
                            v27 += (v27 % ((128) | 1))
                        default:
                            var v34 int64 = 1
                            _ = v34
                        }
                        delete(v1, "foo bar")
                    } else {
                        fmt.Println(v27)
                        var v35 float64 = 3.14159
                        _ = v35
                    }
                    _ = v27
                    _ = v28
                } else {
                    var v36 uint = 255
                    if true {
                        var v37 float64 = (0.1 + 0.1)
                        v37 = float64((v25 / (v25 - v25)))
                        _ = v37
                    } else {
                        if (true || false) {
                            if true {
                                if (2 == 0) {
                                    v25 = v25
                                }
                                var v38 int16 = -420
                                v39 := GenS1{F0: 811, F1: (fi21 >> ((593) & 63))}
                                _ = v38
                                _ = v39
                            } else {
                                if (!true) {
                                    var v40 uint8 = 58
                                    var v41 int32 = 1
                                    if false {
                                        v25 = 0.1
                                        _ = len(v1)
                                    }
                                    var v42 float64 = ((0.1 - 1.5) * 1e10)
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if (true || (false || false)) {
                                        v25 = ((-2.25 + v25) / (1e10 + -2.25))
                                        v36 = (v36 / (((^(v36))) | 1))
                                        v43 := GenS3{F0: v3, F1: (false || ((!(-600 == 0)) || false))}
                                        _ = v43
                                    }
                                }
                                switch "\n\t" {
                                case "hello":
                                    _ = len(v2)
                                }
                                if (v3 >= 65535) {
                                    fmt.Println(v25, fi24, v3)
                                }
                            }
                            v1["\n\t"] = 128
                            var v44 int = 42
                            _ = v44
                        } else {
                            var v45 int64 = int64(v36)
                            if false {
                                _ = len(v2)
                                switch 1 {
                                case -2147483648:
                                    var v46 float64 = (float64((v25 / -2.25)) * 3.14159)
                                    _ = v46
                                case 32767:
                                    v25 = (float32(1.5) / v25)
                                case 2147483647:
                                    v47 := GenS4{F0: 0, F1: 1, F2: -172}
                                    v1["\n\t"] = (fi24 * fi21)
                                    _ = v47
                                default:
                                    v48 := GenS2{F0: 0}
                                    _ = v48
                                }
                                var v49 int32 = -257
                                _ = v49
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = 2
                                }
                                if (true && (!(850 >= (v36 + (v36 ^ 1))))) {
                                    switch (("hello" + ("foo bar" + "")) + ("\n\t" + ((("\n\t" + "hello") + "foo bar") + "\n\t"))) {
                                    case "\n\t":
                                        v50 := GenS5{F0: -379, F1: v36, F2: 42}
                                        _ = v50
                                    case "hello":
                                        var v51 string = "\n\t"
                                        _ = v51
                                    }
                                    v25 = 3.14159
                                } else {
                                    var v52 int8 = 16
                                    switch "" {
                                    case "hello":
                                        var v53 float64 = ((((0.0 + (1.5 * 0.1)) - (1e10 * (-2.25 / ((1.5 + (0.0 + (1e10 / 1.0))) / (0.1 * 1.5))))) + (3.14159 - 3.14159)) - (1e10 - 1.5))
                                        _ = v53
                                    case "\n\t":
                                        switch (fi21 << ((v36) & 63)) {
                                        case 2147483647:
                                            var v54 bool = false
                                            _ = v54
                                        }
                                    case "a":
                                        v52 -= (v52 >> (((v36 >> (((v36 ^ v36)) & 63))) & 63))
                                    }
                                    var v55 int64 = 42
                                    var v56 uint16 = (v3 << ((255) & 63))
                                    v3 -= 255
                                    var v57 string = ("\n\t" + "")
                                    fmt.Println(v55)
                                    _ = v52
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                }
                            }
                            v58 := GenS2{F0: 2}
                            v45 = (v45 & (v45 % (((v45 + v45)) | 1)))
                            if (-128 >= -57) {
                                var v59 string = ("" + "foo bar")
                                _ = v59
                            } else {
                                var v60 int8 = 42
                                if len(v2) > 0 {
                                    v2[0] = 65535
                                }
                                _ = v60
                            }
                            _ = v45
                            _ = v58
                        }
                        v61 := GenS5{F0: -16, F1: 1, F2: 1000}
                        var v62 uint16 = v3
                        var v63 float64 = (-2.25 + 0.0)
                        switch 627 {
                        case 127:
                            var v64 uint64 = 2
                            _ = v64
                        case -753:
                            fmt.Println(v25)
                        case 32767:
                            var v65 float32 = (float32((float64(v25) + (float64(-2.25) - 1.5))) + 3.14159)
                            _ = v65
                        }
                        _ = v61
                        _ = v62
                        _ = v63
                    }
                    v66 := GenS6{F0: true}
                    _ = v36
                    _ = v66
                }
                _ = fi24
                _ = v25
            }
            if (1.5 >= 0.0) {
                for fi67 := 0; fi67 < 13; fi67++ {
                    v3 = (v3 >> ((4294967295) & 63))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(fi67, fi21, v3)
                    _ = fi67
                }
                var v68 int64 = 9223372036854775807
                _ = v68
            }
            _ = fi21
        }
    case "a":
        fmt.Println(v3)
    }
    if (0.0 != 1.5) {
        switch "foo bar" {
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "hello":
            switch "" {
            case "a":
                if (70 > 65535) {
                    var v69 int = 478
                    _ = v69
                }
            case "hello":
                v3 += (v3 | v3)
            case "":
                for fi70 := 0; fi70 < 5; fi70++ {
                    v3 = v3
                    fmt.Println(v3, v3)
                    for fi71 := 0; fi71 < 7; fi71++ {
                        if ((42 < -2147483648) && false) {
                            v72 := GenS7{F0: 50}
                            if false {
                                if (((false || (!(true && ((v3 & 0) < v3)))) && true) || (!true)) {
                                    if (false || false) {
                                        v3 = 1
                                        switch 1 {
                                        case 9223372036854775807:
                                            if (float32(1e10) <= (0.0 - 1e10)) {
                                                v3 += (v3 ^ 0)
                                            } else {
                                                v3 = v3
                                            }
                                        case 340:
                                            if true {
                                                switch ("hello" + "foo bar") {
                                                case "foo bar":
                                                    v2 = append(v2, 1)
                                                }
                                                if (65535 >= 65535) {
                                                    switch ("" + ((("hello" + "a") + (("foo bar" + "foo bar") + "")) + "foo bar")) {
                                                    case "hello":
                                                        var v73 string = (("" + ("a" + "a")) + "\n\t")
                                                        _ = v73
                                                    default:
                                                        v3 = (v3 + 2)
                                                    }
                                                } else {
                                                    switch fi70 {
                                                    case 1:
                                                        var v74 float64 = 1.5
                                                        _ = v74
                                                    default:
                                                        var v75 uint64 = uint64(v3)
                                                        _ = v75
                                                    }
                                                    v76 := v1["a"]
                                                    v76 -= (-(fi70))
                                                    if true {
                                                        switch ("foo bar" + "a") {
                                                        case "foo bar":
                                                            v3 += (^(v3))
                                                        case "hello":
                                                            fmt.Println(v76)
                                                        }
                                                    } else {
                                                        v77 := GenS8{F0: -9223372036854775808}
                                                        v78 := GenS2{F0: 777}
                                                        if true {
                                                            var v79 uint8 = 251
                                                            _ = v79
                                                        } else {
                                                            var v80 int8 = -8
                                                            _ = v80
                                                        }
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                    v81 := GenS4{F0: -128, F1: -440, F2: -1}
                                                    _ = len(v1)
                                                    if (2147483647 <= 2147483648) {
                                                        v1[(("" + "foo bar") + "a")] = (^(fi70))
                                                        var v82 uint64 = 18446744073709551615
                                                        var v83 uint8 = 2
                                                        _ = v82
                                                        _ = v83
                                                    } else {
                                                        v84 := GenS5{F0: -1, F1: 2, F2: 42}
                                                        if true {
                                                            v85 := GenS9{F0: 280, F1: 18446744073709551615, F2: true}
                                                            _ = v85
                                                        }
                                                        _ = v84
                                                    }
                                                    _ = v76
                                                    _ = v81
                                                }
                                                if false {
                                                    var v86 int16 = -32768
                                                    if true {
                                                        _ = len(v1)
                                                        var v87 uint = 855
                                                        _ = v87
                                                    } else {
                                                        var v88 int16 = (v86 % (((v86 >> ((473) & 63))) | 1))
                                                        _ = v88
                                                    }
                                                    var v89 uint64 = 4294967295
                                                    _ = v86
                                                    _ = v89
                                                }
                                            } else {
                                                var v90 bool = false
                                                delete(v1, "")
                                                if true {
                                                    _ = len(v1)
                                                    if true {
                                                        fmt.Println(fi70, fi70, fi71)
                                                        switch "foo bar" {
                                                        case "":
                                                            var v91 float64 = (0.0 / 3.14159)
                                                            _ = v91
                                                        case "hello":
                                                            switch "\n\t" {
                                                            case "a":
                                                                v92 := GenS9{F0: 4294967295, F1: 18446744073709551615, F2: v90}
                                                                _ = v92
                                                            case "":
                                                                v93 := GenS10{F0: 42, F1: genFn0()}
                                                                _ = v93
                                                            case "foo bar":
                                                                v94 := GenS11{F0: 1}
                                                                _ = v94
                                                            }
                                                        }
                                                        v3 -= (v3 * (v3 + (v3 >> ((0) & 63))))
                                                    } else {
                                                        fmt.Println(fi70)
                                                        v2 = append(v2, 1)
                                                    }
                                                    v95 := GenS12{F0: v90, F1: v90}
                                                    _ = v95
                                                }
                                                v3 = v3
                                                v90 = false
                                                if (2 == 42) {
                                                    if (4294967295 <= 4294967295) {
                                                        var v96 float64 = ((((0.1 + 1.5) * (-2.25 / (3.14159 - float64(((1.5 - (0.0 + 3.14159)) / (float32(3.14159) / 1e10)))))) + float64(0.1)) / (-2.25 * 1e10))
                                                        v1[("a" + "hello")] = 32768
                                                        v97 := GenS4{F0: 47, F1: 42, F2: 914}
                                                        var v98 uint = 929
                                                        _ = v96
                                                        _ = v97
                                                        _ = v98
                                                    } else {
                                                        v99 := GenS9{F0: 42, F1: 1, F2: (!(-54 < 793))}
                                                        _ = v99
                                                    }
                                                    v90 = (4294967295 == 616)
                                                }
                                                v100 := GenS13{F0: 42}
                                                v2 = append(v2, 0)
                                                v3 -= (v3 & (^(v3)))
                                                fmt.Println(v90, fi70, fi71)
                                                if (true && false) {
                                                    var v101 uint32 = 42
                                                    _ = v101
                                                }
                                                _ = v100
                                                _ = v90
                                            }
                                        }
                                        fmt.Println(fi70, fi71)
                                    }
                                    fmt.Println(v3, fi71, fi71)
                                    switch ("hello" + "a") {
                                    case "foo bar":
                                        if true {
                                            var v102 int16 = -32768
                                            switch ("hello" + ("" + ("\n\t" + "a"))) {
                                            case "a":
                                                v103 := GenS14{F0: 9223372036854775807}
                                                _ = v103
                                            default:
                                                fmt.Println(fi71)
                                            }
                                            _ = v102
                                        }
                                    case "a":
                                        switch "hello" {
                                        case "\n\t":
                                            if ("foo bar" >= "a") {
                                                v3 = (v3 - v3)
                                                fmt.Println(v3, fi71)
                                                var v104 uint = 255
                                                _ = v104
                                            }
                                        }
                                    }
                                } else {
                                    v105 := GenS15{F0: (0 == -858)}
                                    if true {
                                        v3 -= v3
                                        _ = len(v1)
                                    }
                                    _ = v105
                                }
                                switch fi70 {
                                case 2147483647:
                                    if (("\n\t" + ("hello" + "hello")) != ("foo bar" + (("" + "foo bar") + "foo bar"))) {
                                        var v106 uint32 = 42
                                        v106 += 756
                                        if (0 >= 982) {
                                            _ = len(v2)
                                            fmt.Println(v106, fi70)
                                            switch 42 {
                                            case -2147483648:
                                                switch fi71 {
                                                case 127:
                                                    if true {
                                                        var v107 bool = false
                                                        _ = len(v1)
                                                        _ = v107
                                                    } else {
                                                        fmt.Println(v106)
                                                    }
                                                case 107:
                                                    switch (("\n\t" + "hello") + "hello") {
                                                    case "":
                                                        v106 -= (^(v106))
                                                        var v108 float64 = 1e10
                                                        var v109 uint64 = 0
                                                        _ = v108
                                                        _ = v109
                                                    default:
                                                        delete(v1, ("a" + "foo bar"))
                                                    }
                                                default:
                                                    fmt.Println(fi70)
                                                }
                                            }
                                        } else {
                                            delete(v1, (("hello" + "foo bar") + ("foo bar" + "")))
                                            v3 = (v3 % (((v3 * v3)) | 1))
                                            if true {
                                                v110 := GenS15{F0: ((v106 << ((566) & 63)) < v106)}
                                                switch fi71 {
                                                case 5:
                                                    v106 = (v106 << ((2) & 63))
                                                }
                                                var v111 int64 = -9223372036854775808
                                                if false {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if (127 != 127) {
                                                        var v112 uint8 = 175
                                                        _ = v112
                                                    } else {
                                                        if ((true || (true || (42 <= 42))) || (!(!(v111 <= v111)))) {
                                                            fmt.Println(v106, v111)
                                                            v113 := GenS16{F0: v106, F1: -1}
                                                            _ = v113
                                                        }
                                                    }
                                                }
                                                if (!(!true)) {
                                                    v114 := GenS1{F0: (^(v106)), F1: fi70}
                                                    _ = v114
                                                }
                                                _ = v110
                                                _ = v111
                                            }
                                            v3 += v3
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v115 uint8 = 128
                                        if true {
                                            v116 := GenS9{F0: 18446744073709551615, F1: 449, F2: false}
                                            _ = v116
                                        }
                                        var v117 float64 = 0.0
                                        _ = v106
                                        _ = v115
                                        _ = v117
                                    }
                                case -231:
                                    var v118 string = ("hello" + "a")
                                    _ = v118
                                }
                                v119 := GenS15{F0: true}
                                _ = v119
                            }
                            v120 := GenS1{F0: 2147483647, F1: (fi70 >> ((4294967295) & 63))}
                            _ = v120
                            _ = v72
                        } else {
                            switch (("a" + "foo bar") + "\n\t") {
                            case "\n\t":
                                if true {
                                    fmt.Println(v3)
                                    v3 = (v3 / (((v3 / ((v3) | 1))) | 1))
                                    var v121 uint8 = 0
                                    _ = v121
                                }
                            case "a":
                                var v122 int = fi71
                                _ = v122
                            case "hello":
                                var v123 uint32 = 4294967295
                                _ = v123
                            }
                            var v124 uint64 = 42
                            _ = v124
                        }
                        _ = fi71
                    }
                    _ = fi70
                }
            }
        case "foo bar":
            v3 = (v3 + (v3 >> ((255) & 63)))
        default:
            v3 += v3
        }
    } else {
        v125 := GenS17{F0: 208, F1: 2147483648}
        for fi126 := 0; fi126 < 6; fi126++ {
            for fi127 := 0; fi127 < 12; fi127++ {
                var v128 int16 = -516
                fmt.Println(fi126)
                _ = fi127
                _ = v128
            }
            if (true || (((1.5 / 1e10) + (3.14159 / (3.14159 / 1.0))) >= 0.0)) {
                v129 := GenS9{F0: 2, F1: 440, F2: true}
                v3 = 42
                fmt.Println(v3)
                _ = v129
            } else {
                var v130 bool = (fi126 < -9223372036854775808)
                v131 := GenS3{F0: v3, F1: v130}
                v132 := GenS9{F0: 1, F1: 1, F2: (42 == 34)}
                _ = v130
                _ = v131
                _ = v132
            }
            v133 := GenS18{F0: 2}
            _ = fi126
            _ = v133
        }
        v3 = v3
        _ = v125
    }
    var v134 uint8 = 255
    if (!(!(float32(-2.25) >= 3.14159))) {
        for fi135 := 0; fi135 < 17; fi135++ {
            v136 := GenS19{F0: v134, F1: genFn0()}
            for fi137 := 0; fi137 < 2; fi137++ {
                v134 -= 2
                v3 = (^(v3))
                fmt.Println(fi135, v134)
                var v138 int64 = -101
                _ = fi137
                _ = v138
            }
            var v139 uint = 239
            _ = fi135
            _ = v136
            _ = v139
        }
    } else {
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    }
    v140 := GenS10{F0: 282, F1: -128}
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        var v141 uint64 = 42
        _ = v141
    case "":
        _ = len(v2)
    default:
        var v142 uint16 = v3
        _ = v142
    }
    var v143 uint = 789
    for fi144 := 0; fi144 < 13; fi144++ {
        switch "foo bar" {
        case "":
            if (((float32(0.1) * 0.0) * -2.25) > 1e10) {
                switch (^(fi144)) {
                case -371:
                    v143 = v143
                case 0:
                    switch "\n\t" {
                    case "a":
                        _ = len(v2)
                    }
                case 709:
                    v143 = v143
                }
                var v145 int32 = -2147483648
                _ = v145
            }
        case "foo bar":
            if true {
                for fi146 := 0; fi146 < 13; fi146++ {
                    v134 = v134
                    switch (fi144 * 599) {
                    case -720:
                        if (((0.0 - 1e10) - (1e10 * 0.0)) <= 1.5) {
                            var v147 int8 = -128
                            var v148 string = "\n\t"
                            if len(v2) > 0 {
                                v2[0] = 191
                            }
                            _ = v147
                            _ = v148
                        }
                    }
                    if false {
                        var v149 string = "hello"
                        _ = v149
                    } else {
                        v150 := GenS20{F0: 42, F1: true, F2: true}
                        switch ("" + "") {
                        case "foo bar":
                            var v151 int16 = 1
                            _ = v151
                        case "":
                            v2 = append(v2, 924)
                            fmt.Println(v3, v134)
                            var v152 uint8 = (^(v134))
                            fmt.Println(v143)
                            _ = v152
                        default:
                            v3 -= v3
                        }
                        var v153 int = -2147483648
                        _ = v150
                        _ = v153
                    }
                    _ = fi146
                }
            }
        case "\n\t":
            if (3.14159 > 1.5) {
                v1["a"] = -9223372036854775808
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v154 int16 = int16(v134)
                _ = v154
            }
        default:
            var v155 bool = (-128 <= -128)
            _ = v155
        }
        _ = fi144
    }
    fmt.Println(v134, v3, v3)
    var v156 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v134:", v134)
    fmt.Println("v143:", v143)
    fmt.Println("v156:", v156)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v134
    _ = v140
    _ = v143
    _ = v156
    _ = v2
    _ = v3
}
