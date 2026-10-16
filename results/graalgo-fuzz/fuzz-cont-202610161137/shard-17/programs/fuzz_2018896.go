package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int8; F1 int }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint8; F1 bool }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int32; F1 string; F2 uint16 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 int32 }
func genFn0(p0 int, p1 int32) uint64 {
    var v6 uint = 18446744073709551615
    v7 := GenS1{F0: 255, F1: p1, F2: (p0 % ((p0) | 1))}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 756
    v2 := make(map[string]bool)
    var v3 uint16 = 0
    var v4 uint64 = 333
    var v5 uint64 = (v1 << ((42) & 63))
    if (2 != 42) {
        switch 88 {
        case -9223372036854775808:
            if (int64(v3) != -1) {
                v5 = v1
                v5 = 695
                var v8 bool = ((!(true || false)) || (!true))
                _ = v8
            } else {
                switch 128 {
                case 1:
                    if (1.5 != 1.5) {
                        for fi9 := 0; fi9 < 9; fi9++ {
                            v5 = (v4 >> ((495) & 63))
                            for fi10 := 0; fi10 < 3; fi10++ {
                                v11 := GenS1{F0: 148, F1: -2147483648, F2: (fi9 - fi9)}
                                switch fi10 {
                                case 32767:
                                    fmt.Println(v1, fi10, fi10)
                                }
                                if (1e10 <= 1.5) {
                                    if (false || (!true)) {
                                        switch (fi10 >> ((0) & 63)) {
                                        case 42:
                                            v3 -= (v3 % ((v3) | 1))
                                            var v12 bool = false
                                            var v13 uint = 0
                                            fmt.Println(v4)
                                            var v14 uint64 = 0
                                            _ = v12
                                            _ = v13
                                            _ = v14
                                        case 128:
                                            v15 := []int{fi9, 4294967295, fi9}
                                            _ = v15
                                        case 32767:
                                            v16 := GenS1{F0: 255, F1: -1, F2: fi10}
                                            _ = len(v2)
                                            _ = v16
                                        default:
                                            v17 := []int{fi9, 127, (fi10 * fi10), fi10}
                                            if (!(0 > 853)) {
                                                switch fi9 {
                                                case 4294967295:
                                                    v1 += v4
                                                case 1:
                                                    v18 := v2["foo bar"]
                                                    v2["\n\t"] = v18
                                                    v17 = append(v17, fi10)
                                                    _ = v18
                                                default:
                                                    var v19 int32 = -931
                                                    _ = v19
                                                }
                                                v20 := GenS1{F0: 1, F1: 42, F2: fi9}
                                                fmt.Println(v3, fi9, v4)
                                                switch (fi10 >> ((65535) & 63)) {
                                                case -727:
                                                    switch ("" + ("\n\t" + ("\n\t" + "foo bar"))) {
                                                    case "hello":
                                                        v21 := GenS1{F0: 65535, F1: -102, F2: fi9}
                                                        _ = v21
                                                    }
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v17 = append(v17, -2147483648)
                                                    case "hello":
                                                        v4 -= genFn0((fi10 >> ((42) & 63)), 450)
                                                    case "a":
                                                        v4 += (^(v4))
                                                    }
                                                    v1 -= v1
                                                case 259:
                                                    fmt.Println(fi10)
                                                }
                                                fmt.Println(fi9)
                                                _ = v20
                                            } else {
                                                if ((v5 > 42) || false) {
                                                    v5 = genFn0((fi9 & fi9), 501)
                                                    v17 = append(v17, fi10)
                                                    switch ("hello" + "foo bar") {
                                                    case "foo bar":
                                                        if true {
                                                            var v22 int = (fi9 ^ fi9)
                                                            var v23 uint16 = v3
                                                            _ = v22
                                                            _ = v23
                                                        } else {
                                                            var v24 float64 = 3.14159
                                                            switch fi10 {
                                                            case -9223372036854775808:
                                                                fmt.Println(v5)
                                                            case -2147483648:
                                                                switch (^(fi9)) {
                                                                case 0:
                                                                    fmt.Println(v4, v3)
                                                                case 42:
                                                                    v25 := v2[(((("" + "hello") + "foo bar") + ("hello" + "a")) + "")]
                                                                    _ = v25
                                                                case 32767:
                                                                    v24 = -2.25
                                                                }
                                                            case 127:
                                                                var v26 uint32 = 956
                                                                _ = v26
                                                            default:
                                                                switch fi9 {
                                                                case 32767:
                                                                    v27 := GenS2{F0: "foo bar", F1: 4294967295}
                                                                    _ = v27
                                                                case 1:
                                                                    v28 := GenS2{F0: "foo bar", F1: 42}
                                                                    _ = v28
                                                                case -9223372036854775808:
                                                                    v2["\n\t"] = (!true)
                                                                }
                                                            }
                                                            _ = v24
                                                        }
                                                    case "hello":
                                                        if (!(3.14159 <= -2.25)) {
                                                            v4 = v1
                                                            if (1.5 <= 0.1) {
                                                                v29 := v2["foo bar"]
                                                                var v30 float32 = (1e10 - -2.25)
                                                                _ = len(v2)
                                                                v31 := v2[(("foo bar" + "foo bar") + "foo bar")]
                                                                if (1000 != 1) {
                                                                    _ = len(v2)
                                                                    if false {
                                                                        if (!v29) {
                                                                            v30 = v30
                                                                            fmt.Println(fi10, v3, v1)
                                                                            var v32 int8 = -1
                                                                            _ = v32
                                                                        }
                                                                        if v29 {
                                                                            if true {
                                                                                v5 = (v5 | (v4 | (^(v1))))
                                                                                v33 := GenS2{F0: "", F1: 805}
                                                                                var v34 bool = false
                                                                                _ = v33
                                                                                _ = v34
                                                                            }
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "":
                                                                                var v35 uint16 = (v3 - v3)
                                                                                _ = v35
                                                                            case "hello":
                                                                                v36 := v2["a"]
                                                                                _ = v36
                                                                            case "foo bar":
                                                                                var v37 float64 = float64(float32(1e10))
                                                                                _ = v37
                                                                            default:
                                                                                var v38 uint16 = v3
                                                                                _ = v38
                                                                            }
                                                                            if false {
                                                                                v30 = v30
                                                                                if ((("foo bar" + ("" + "")) < "\n\t") || ((128 == 2) && v29)) {
                                                                                    fmt.Println(v29)
                                                                                    v39 := GenS2{F0: "hello", F1: 65535}
                                                                                    _ = v39
                                                                                } else {
                                                                                    v40 := v2[""]
                                                                                    if false {
                                                                                        fmt.Println(v40, v29, v3)
                                                                                        switch (fi9 - (fi9 + (fi10 ^ 32767))) {
                                                                                        case 32768:
                                                                                            var v41 int32 = -713
                                                                                            _ = v41
                                                                                        case -1:
                                                                                            var v42 uint8 = 236
                                                                                            _ = v42
                                                                                        }
                                                                                        var v43 bool = true
                                                                                        _ = v43
                                                                                    }
                                                                                    _ = v40
                                                                                }
                                                                                v31 = ((float32(0.0) / (v30 - v30)) <= 1.5)
                                                                            }
                                                                            var v44 int16 = 534
                                                                            _ = v44
                                                                        }
                                                                    }
                                                                } else {
                                                                    _ = len(v2)
                                                                }
                                                                _ = v29
                                                                _ = v30
                                                                _ = v31
                                                            } else {
                                                                var v45 string = "foo bar"
                                                                v1 = (v4 | 0)
                                                                if false {
                                                                    var v46 string = v45
                                                                    var v47 string = (v46 + "")
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                                _ = v45
                                                            }
                                                            v3 = (v3 + v3)
                                                            v48 := GenS2{F0: ("foo bar" + (("\n\t" + "\n\t") + "hello")), F1: 42}
                                                            _ = v48
                                                        } else {
                                                            switch 32767 {
                                                            case 32767:
                                                                fmt.Println(v5, fi9)
                                                            case 127:
                                                                var v49 float64 = 0.0
                                                                _ = v49
                                                            default:
                                                                if (2147483647 == 322) {
                                                                    switch (("hello" + "hello") + "hello") {
                                                                    case "\n\t":
                                                                        if false {
                                                                            _ = len(v17)
                                                                            v17 = append(v17, 127)
                                                                            v5 += v5
                                                                        } else {
                                                                            var v50 uint32 = 2147483648
                                                                            var v51 uint8 = 20
                                                                            _ = v50
                                                                            _ = v51
                                                                        }
                                                                    default:
                                                                        var v52 int64 = 9223372036854775807
                                                                        _ = v52
                                                                    }
                                                                    switch fi10 {
                                                                    case -572:
                                                                        v53 := v2[""]
                                                                        _ = v53
                                                                    case 4294967295:
                                                                        v54 := GenS3{F0: 0}
                                                                        _ = v54
                                                                    case 0:
                                                                        var v55 bool = (1 != v4)
                                                                        _ = v55
                                                                    default:
                                                                        v1 += (v1 >> ((65535) & 63))
                                                                    }
                                                                } else {
                                                                    v56 := GenS4{F0: v1}
                                                                    _ = v56
                                                                }
                                                            }
                                                            fmt.Println(fi10)
                                                        }
                                                    }
                                                }
                                                var v57 string = "\n\t"
                                                v58 := GenS3{F0: 106}
                                                _ = v57
                                                _ = v58
                                            }
                                            if false {
                                                v1 -= v5
                                                v1 += v1
                                            } else {
                                                v1 += v5
                                                if (376 != 943) {
                                                    v3 = v3
                                                    var v59 int16 = 42
                                                    delete(v2, ("a" + "foo bar"))
                                                    if (-1 < -424) {
                                                        switch (fi9 * (fi9 + fi9)) {
                                                        case 923:
                                                            var v60 int64 = 268
                                                            if true {
                                                                switch "foo bar" {
                                                                case "":
                                                                    v1 = (^(v1))
                                                                case "\n\t":
                                                                    var v61 uint = 42
                                                                    v62 := GenS5{F0: "a"}
                                                                    v63 := GenS2{F0: (((("" + "foo bar") + "hello") + ("hello" + "hello")) + ""), F1: uint(v3)}
                                                                    _ = v61
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                                v64 := v2["hello"]
                                                                v3 = v3
                                                                _ = v64
                                                            } else {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    fmt.Println(fi9, fi10)
                                                                case "a":
                                                                    var v65 int8 = 1
                                                                    _ = v65
                                                                case "\n\t":
                                                                    var v66 uint32 = 764
                                                                    _ = v66
                                                                }
                                                                if (-123 != -51) {
                                                                    var v67 uint = 816
                                                                    _ = v67
                                                                }
                                                            }
                                                            var v68 float32 = 3.14159
                                                            _ = v60
                                                            _ = v68
                                                        case 734:
                                                            v69 := v2["hello"]
                                                            _ = v69
                                                        default:
                                                            v5 = v4
                                                        }
                                                    } else {
                                                        v4 += v5
                                                        var v70 int32 = 42
                                                        switch fi9 {
                                                        case 42:
                                                            delete(v2, "hello")
                                                        }
                                                        _ = v70
                                                    }
                                                    var v71 int16 = v59
                                                    var v72 uint = 0
                                                    if (true && (!(42 > 255))) {
                                                        _ = len(v17)
                                                        v1 = v1
                                                    } else {
                                                        if (uint8(v72) > 255) {
                                                            v73 := GenS6{F0: 42, F1: (v4 + (v4 & (v4 << (((v72 - v72)) & 63))))}
                                                            var v74 string = "foo bar"
                                                            _ = v73
                                                            _ = v74
                                                        } else {
                                                            v75 := GenS7{F0: 1, F1: 169}
                                                            switch ("foo bar" + "") {
                                                            case "foo bar":
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v76 int = (fi10 % ((32768) | 1))
                                                                    _ = v76
                                                                case "\n\t":
                                                                    var v77 int32 = 0
                                                                    _ = v77
                                                                case "hello":
                                                                    var v78 uint = (^(v72))
                                                                    _ = v78
                                                                default:
                                                                    var v79 int16 = -18
                                                                    _ = v79
                                                                }
                                                            case "\n\t":
                                                                if false {
                                                                    _ = len(v2)
                                                                    var v80 uint32 = 15
                                                                    _ = v80
                                                                } else {
                                                                    v3 = 1
                                                                    var v81 float32 = (0.1 * float32(0.1))
                                                                    _ = v81
                                                                }
                                                            case "":
                                                                v17 = append(v17, -711)
                                                            default:
                                                                v82 := GenS4{F0: genFn0((fi10 * fi10), 0)}
                                                                _ = v82
                                                            }
                                                            switch "" {
                                                            case "":
                                                                var v83 int64 = int64(v1)
                                                                _ = v83
                                                            case "hello":
                                                                v71 = (^(v59))
                                                            default:
                                                                v84 := GenS8{F0: -128}
                                                                _ = v84
                                                            }
                                                            _ = v75
                                                        }
                                                        var v85 int64 = 9223372036854775807
                                                        v86 := GenS8{F0: int8(v85)}
                                                        fmt.Println(v71, v5)
                                                        var v87 bool = false
                                                        switch -9223372036854775808 {
                                                        case 817:
                                                            var v88 int16 = v71
                                                            _ = v88
                                                        case 128:
                                                            if (v87 || ("foo bar" <= ("" + "a"))) {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    v89 := GenS9{F0: 42, F1: (v71 + (v71 & 638))}
                                                                    _ = v89
                                                                case "foo bar":
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        v4 -= (^(v5))
                                                                        if len(v17) >= 2 {
                                                                            v17 = v17[1:len(v17)]
                                                                        }
                                                                    default:
                                                                        fmt.Println(v87, v5)
                                                                    }
                                                                }
                                                                v87 = false
                                                                v3 = (v3 / ((v3) | 1))
                                                            } else {
                                                                var v90 int32 = -2147483648
                                                                fmt.Println(v4)
                                                                _ = v90
                                                            }
                                                        }
                                                        _ = v85
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                    var v91 float64 = 3.14159
                                                    var v92 uint64 = v5
                                                    _ = v59
                                                    _ = v71
                                                    _ = v72
                                                    _ = v91
                                                    _ = v92
                                                } else {
                                                    v3 = v3
                                                    if true {
                                                        var v93 int = (-(fi10))
                                                        v2[("hello" + ("\n\t" + ("a" + "a")))] = true
                                                        v94 := GenS2{F0: ("hello" + "hello"), F1: 1}
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                }
                                            }
                                            _ = v17
                                        }
                                        if false {
                                            if true {
                                                v1 = 572
                                                v95 := []int{(^(fi9)), (fi9 ^ (fi10 + 42)), (-(fi10))}
                                                switch "" {
                                                case "a":
                                                    v4 -= v4
                                                case "hello":
                                                    var v96 uint16 = (v3 ^ (v3 % (((v3 << ((4294967295) & 63))) | 1)))
                                                    _ = v96
                                                }
                                                _ = v95
                                            }
                                            v3 = 984
                                        }
                                    }
                                } else {
                                    fmt.Println(v1, v4, v5)
                                    if true {
                                        if ((4294967295 != 189) && (42 > 906)) {
                                            var v97 float64 = 1e10
                                            if (4294967295 != (v4 ^ (v4 - (^(v5))))) {
                                                v4 = (v5 | genFn0((fi9 - fi10), int32(fi10)))
                                                var v98 float64 = v97
                                                if false {
                                                    var v99 uint64 = v1
                                                    fmt.Println(v3, v4)
                                                    var v100 uint32 = 65535
                                                    _ = v100
                                                    _ = v99
                                                }
                                                v3 = v3
                                                _ = v98
                                            } else {
                                                v3 = 617
                                            }
                                            var v101 float64 = v97
                                            var v102 uint16 = (v3 << ((65535) & 63))
                                            _ = v101
                                            _ = v102
                                            _ = v97
                                        }
                                        v3 -= (v3 * 529)
                                        var v103 uint64 = (v1 + v4)
                                        _ = v103
                                    }
                                }
                                _ = fi10
                                _ = v11
                            }
                            _ = fi9
                        }
                        if true {
                            fmt.Println(v3, v4, v5)
                        }
                    } else {
                        if false {
                            v2[((("hello" + ((("foo bar" + (("a" + "hello") + ("\n\t" + "\n\t"))) + ("hello" + "foo bar")) + "")) + "a") + ("\n\t" + "\n\t"))] = false
                            for fi104 := 0; fi104 < 16; fi104++ {
                                if ((v5 % ((v5) | 1)) == (v1 % ((genFn0(fi104, -240)) | 1))) {
                                    _ = len(v2)
                                    var v105 int16 = 249
                                    switch "" {
                                    case "foo bar":
                                        fmt.Println(v4, fi104)
                                    case "\n\t":
                                        v106 := []int{fi104, fi104}
                                        _ = v106
                                    }
                                    _ = v105
                                } else {
                                    var v107 int8 = 127
                                    _ = v107
                                }
                                switch fi104 {
                                case -170:
                                    var v108 uint16 = v3
                                    _ = v108
                                case 951:
                                    v109 := GenS9{F0: 100, F1: 158}
                                    _ = v109
                                case -748:
                                    var v110 uint16 = v3
                                    _ = v110
                                default:
                                    v111 := GenS10{F0: 381}
                                    _ = v111
                                }
                                _ = fi104
                            }
                        } else {
                            fmt.Println(v5, v4, v5)
                            var v112 string = (("\n\t" + "\n\t") + ((("\n\t" + ("" + ("hello" + ((("hello" + "hello") + "") + "a")))) + (("\n\t" + "foo bar") + "a")) + "foo bar"))
                            _ = v112
                        }
                    }
                default:
                    v113 := GenS11{F0: -128, F1: 2147483647}
                    _ = v113
                }
            }
        case -2147483648:
            for fi114 := 0; fi114 < 9; fi114++ {
                v5 -= v4
                _ = fi114
            }
        default:
            for fi115 := 0; fi115 < 13; fi115++ {
                var v116 int8 = -128
                switch "foo bar" {
                case "foo bar":
                    for fi117 := 0; fi117 < 8; fi117++ {
                        v1 = (v4 + (^(v1)))
                        if false {
                            fmt.Println(fi115, v1)
                        } else {
                            switch "hello" {
                            case "\n\t":
                                var v118 float64 = (((1e10 * float64(float32(-2.25))) * (3.14159 - 0.1)) + -2.25)
                                _ = v118
                            case "foo bar":
                                v3 = v3
                            case "":
                                if (!true) {
                                    v119 := GenS9{F0: v116, F1: 42}
                                    v120 := GenS12{F0: 2}
                                    var v121 int16 = 0
                                    _ = v119
                                    _ = v120
                                    _ = v121
                                }
                            default:
                                var v122 float32 = ((0.1 + (((1.5 + 0.0) * float32(-2.25)) + 0.0)) / 1.0)
                                _ = v122
                            }
                            var v123 uint64 = genFn0(-9223372036854775808, 735)
                            _ = v123
                        }
                        delete(v2, "foo bar")
                        _ = fi117
                    }
                case "a":
                    if true {
                        if true {
                            v4 = genFn0(fi115, 629)
                            v3 = v3
                            switch -426 {
                            case -835:
                                if false {
                                    var v124 float64 = (((((0.0 + 0.1) + 3.14159) / -2.25) - ((1e10 + (-2.25 + 1.5)) / (-2.25 * (3.14159 * (0.1 / 3.14159))))) * 3.14159)
                                    var v125 float32 = (0.0 - (3.14159 - 0.0))
                                    v126 := GenS12{F0: 0}
                                    _ = v124
                                    _ = v125
                                    _ = v126
                                } else {
                                    if false {
                                        v1 = v1
                                        if (!(!true)) {
                                            switch fi115 {
                                            case 42:
                                                switch (((("\n\t" + "hello") + "\n\t") + "hello") + "foo bar") {
                                                case "hello":
                                                    var v127 int8 = v116
                                                    _ = v127
                                                case "\n\t":
                                                    for fi128 := 0; fi128 < 10; fi128++ {
                                                        if true {
                                                            v129 := []int{fi128, (fi115 + (fi115 - -190)), (fi128 - fi115), fi115}
                                                            _ = v129
                                                        } else {
                                                            if false {
                                                                v130 := GenS13{F0: v116}
                                                                switch -352 {
                                                                case -972:
                                                                    if ((32767 != 1) && false) {
                                                                        v131 := v2["\n\t"]
                                                                        v131 = (((float32((1.5 / (((1.5 - 1.5) * 0.0) - ((-2.25 - ((3.14159 - 0.0) * 3.14159)) / 3.14159)))) + float32(float64(float32(float64((1e10 + (-2.25 - ((0.0 / 3.14159) * 1.5)))))))) - -2.25) == (3.14159 / float32(0.1)))
                                                                        _ = v131
                                                                    }
                                                                }
                                                                _ = v130
                                                            }
                                                        }
                                                        v132 := GenS4{F0: v4}
                                                        _ = fi128
                                                        _ = v132
                                                    }
                                                }
                                            }
                                        } else {
                                            var v133 int = fi115
                                            for fi134 := 0; fi134 < 11; fi134++ {
                                                v135 := GenS7{F0: 4294967295, F1: 0}
                                                v1 -= (v4 * 4294967295)
                                                _ = fi134
                                                _ = v135
                                            }
                                            _ = v133
                                        }
                                    } else {
                                        v3 = v3
                                        fmt.Println(v1)
                                        v136 := []int{4294967295}
                                        _ = v136
                                    }
                                }
                            case 490:
                                var v137 uint = 452
                                _ = v137
                            case 36:
                                fmt.Println(fi115, v4, v1)
                            default:
                                v116 -= (^(v116))
                            }
                        } else {
                            v138 := GenS14{F0: 42, F1: (2 > 999)}
                            var v139 uint = 662
                            switch fi115 {
                            case 814:
                                var v140 uint8 = 255
                                _ = v140
                            case 42:
                                var v141 int16 = 42
                                _ = v141
                            case -21:
                                for fi142 := 0; fi142 < 14; fi142++ {
                                    fmt.Println(fi142)
                                    _ = fi142
                                }
                            default:
                                if false {
                                    v139 = 1
                                    _ = len(v2)
                                }
                            }
                            _ = v138
                            _ = v139
                        }
                        for fi143 := 0; fi143 < 13; fi143++ {
                            v144 := []int{42}
                            _ = fi143
                            _ = v144
                        }
                        v1 -= (v5 + (v1 + (^(v4))))
                    } else {
                        _ = len(v2)
                        for fi145 := 0; fi145 < 11; fi145++ {
                            var v146 uint8 = 2
                            var v147 string = ("foo bar" + "a")
                            v148 := v2["\n\t"]
                            _ = fi145
                            _ = v146
                            _ = v147
                            _ = v148
                        }
                    }
                }
                _ = fi115
                _ = v116
            }
        }
    }
    if false {
        var v149 uint8 = 117
        switch "a" {
        case "":
            v150 := GenS15{F0: v3}
            _ = v150
        case "\n\t":
            var v151 uint16 = uint16(v4)
            _ = v151
        default:
            var v152 uint32 = uint32(v3)
            _ = v152
        }
        if (18446744073709551615 < 18446744073709551615) {
            fmt.Println(v1, v149, v3)
        }
        _ = v149
    }
    v1 = (v1 / ((v4) | 1))
    if true {
        fmt.Println(v4)
        var v153 int = 177
        var v154 bool = true
        _ = v153
        _ = v154
    } else {
        for fi155 := 0; fi155 < 10; fi155++ {
            var v156 int = fi155
            var v157 int8 = -80
            var v158 int64 = -719
            _ = fi155
            _ = v156
            _ = v157
            _ = v158
        }
        v159 := []int{42, 861}
        _ = v159
    }
    _ = len(v2)
    var v160 uint64 = 4294967295
    var v161 uint8 = 128
    fmt.Println(v3)
    fmt.Println(v5, v3, v3)
    switch int(v161) {
    case 2147483647:
        var v162 float64 = (1e10 * 0.1)
        _ = v162
    }
    if ((v161 & v161) < (^(v161))) {
        var v163 string = "a"
        if (!(true && true)) {
            v3 -= v3
            for fi164 := 0; fi164 < 10; fi164++ {
                fmt.Println(v161)
                var v165 float32 = 1e10
                _ = fi164
                _ = v165
            }
            v166 := []int{145}
            _ = v166
        }
        _ = v163
    }
    if true {
        var v167 uint = 1
        switch 42 {
        case 42:
            for fi168 := 0; fi168 < 14; fi168++ {
                delete(v2, "foo bar")
                _ = fi168
            }
        case -9223372036854775808:
            v169 := v2[("a" + "a")]
            _ = v169
        default:
            fmt.Println(v161, v167, v167)
        }
        _ = v167
    } else {
        if false {
            v160 = (v4 - v1)
            var v170 float64 = (((1e10 * (0.0 * 1.5)) + 1e10) + ((-2.25 / -2.25) + float64(0.0)))
            _ = v170
        }
        v171 := []int{-655, -2147483648, 1, 0}
        fmt.Println(v3)
        _ = v171
    }
    var v172 float32 = ((-2.25 - float32(3.14159)) / 1.5)
    if ((-2.25 - 0.0) < v172) {
        v172 += (1.5 * 0.1)
        v1 = (v4 & v5)
        v173 := GenS16{F0: 0, F1: "hello", F2: 65535}
        _ = v173
    } else {
        var v174 int32 = 42
        for fi175 := 0; fi175 < 3; fi175++ {
            v5 -= v1
            var v176 int = (fi175 | -756)
            fmt.Println(fi175, v172)
            _ = fi175
            _ = v176
        }
        if (int(v174) > -566) {
            v2[(("a" + "a") + "\n\t")] = (!false)
            if (!false) {
                var v177 int64 = 70
                if (genFn0(4294967295, (-(v174))) >= (v1 & v4)) {
                    var v178 int = 128
                    if true {
                        fmt.Println(v1, v177)
                        v179 := GenS17{F0: 77}
                        for fi180 := 0; fi180 < 5; fi180++ {
                            var v181 uint8 = (v161 + (v161 << ((750) & 63)))
                            if true {
                                v1 = 0
                                v182 := []int{4294967295, fi180}
                                v178 += 1
                                _ = v182
                            }
                            v183 := GenS6{F0: v161, F1: (v1 / (((v1 << ((42) & 63))) | 1))}
                            _ = fi180
                            _ = v181
                            _ = v183
                        }
                        _ = v179
                    }
                    _ = v178
                } else {
                    v3 = 0
                }
                v160 += 954
                _ = v177
            } else {
                for fi184 := 0; fi184 < 6; fi184++ {
                    switch "foo bar" {
                    case "a":
                        for fi185 := 0; fi185 < 11; fi185++ {
                            if true {
                                v161 = 0
                                var v186 uint8 = 255
                                v161 = (^(v161))
                                var v187 float64 = 1e10
                                var v188 int16 = int16(fi184)
                                switch ("foo bar" + ("a" + ("\n\t" + "a"))) {
                                case "foo bar":
                                    if (false || true) {
                                        v189 := GenS12{F0: 2147483648}
                                        _ = v189
                                    } else {
                                        v2[((((("\n\t" + "") + ("foo bar" + "\n\t")) + "hello") + "\n\t") + "hello")] = (!false)
                                        v3 += v3
                                    }
                                case "":
                                    v190 := []int{fi185, -1}
                                    if (true && true) {
                                        var v191 float32 = (v172 + ((-2.25 + 0.0) + (((v172 * v172) - 1.5) + (v172 * v172))))
                                        switch fi184 {
                                        case 0:
                                            v192 := GenS17{F0: -128}
                                            _ = v192
                                        case 32768:
                                            var v193 string = "hello"
                                            _ = v193
                                        case -777:
                                            var v194 float32 = 0.0
                                            _ = v194
                                        default:
                                            v187 = (v187 + (v187 / 3.14159))
                                        }
                                        _ = v191
                                    }
                                    v190 = append(v190, (fi185 >> ((255) & 63)))
                                    switch "" {
                                    case "a":
                                        _ = len(v2)
                                    case "foo bar":
                                        switch "foo bar" {
                                        case "":
                                            if ((v3 & v3) >= 435) {
                                                var v195 string = "a"
                                                _ = v195
                                            }
                                            var v196 bool = (v186 >= 12)
                                            v190 = append(v190, (-(fi184)))
                                            _ = v196
                                        case "hello":
                                            fmt.Println(v188)
                                        default:
                                            var v197 int8 = 100
                                            _ = v197
                                        }
                                    }
                                    fmt.Println(v5)
                                    _ = v190
                                case "hello":
                                    var v198 int32 = 2147483647
                                    _ = v198
                                default:
                                    switch "\n\t" {
                                    case "":
                                        v199 := []int{32767}
                                        _ = v199
                                    case "a":
                                        v174 -= -1
                                    default:
                                        v200 := []int{(fi184 | 127), 32768, fi184}
                                        _ = v200
                                    }
                                }
                                switch "" {
                                case "\n\t":
                                    fmt.Println(v3, v3)
                                case "foo bar":
                                    if ((fi185 << ((uint(v1)) & 63)) != (fi184 & (fi185 | 42))) {
                                        if ((int32(v4) <= -1) || false) {
                                            fmt.Println(v160)
                                        }
                                    } else {
                                        v174 -= v174
                                    }
                                case "hello":
                                    v201 := []int{(fi185 + fi184)}
                                    _ = v201
                                }
                                _ = v186
                                _ = v187
                                _ = v188
                            } else {
                                fmt.Println(v174, v172, v5)
                                fmt.Println(v4)
                            }
                            v4 += uint64(v3)
                            v161 = 2
                            v202 := GenS12{F0: 65535}
                            if true {
                                if (true || false) {
                                    fmt.Println(fi185, fi184)
                                    if false {
                                        var v203 uint8 = 2
                                        fmt.Println(fi184)
                                        var v204 int8 = -1
                                        var v205 uint64 = genFn0(127, (v174 << ((uint(v3)) & 63)))
                                        var v206 string = "foo bar"
                                        v5 = (v205 << ((4294967295) & 63))
                                        v207 := []int{fi185, (^(fi184)), 127, (-(fi185))}
                                        _ = v203
                                        _ = v204
                                        _ = v205
                                        _ = v206
                                        _ = v207
                                    }
                                } else {
                                    var v208 int32 = -2147483648
                                    var v209 int16 = 41
                                    fmt.Println(v160, fi184, v208)
                                    fmt.Println(v174)
                                    _ = v208
                                    _ = v209
                                }
                                if false {
                                    v2["\n\t"] = false
                                    v210 := GenS5{F0: "hello"}
                                    v211 := GenS18{F0: 42}
                                    _ = v210
                                    _ = v211
                                } else {
                                    v212 := GenS17{F0: 1}
                                    _ = v212
                                }
                            } else {
                                var v213 float64 = float64(-2.25)
                                _ = v213
                            }
                            switch (-(fi185)) {
                            case -1:
                                delete(v2, "foo bar")
                            case 0:
                                var v214 int32 = (^(v174))
                                _ = v214
                            case 9223372036854775807:
                                v174 -= v174
                            }
                            _ = fi185
                            _ = v202
                        }
                    }
                    _ = fi184
                }
                v215 := GenS3{F0: -102}
                _ = v215
            }
        } else {
            v5 += (v4 + v4)
        }
        _ = v174
    }
    switch -338 {
    case 575:
        v172 += 0.1
    }
    v1 -= (^(v5))
    v160 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v160:", v160)
    fmt.Println("v161:", v161)
    fmt.Println("v172:", v172)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v160
    _ = v161
    _ = v172
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
