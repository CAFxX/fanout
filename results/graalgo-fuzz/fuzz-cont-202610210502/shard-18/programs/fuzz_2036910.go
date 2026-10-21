package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 bool }
type GenS2 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint8; F1 string; F2 uint16 }
type GenS9 struct { F0 int32; F1 bool; F2 int16 }
type GenS10 struct { F0 string; F1 int }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int8; F1 int64 }
type GenS13 struct { F0 int32; F1 int32; F2 int16 }
type GenS14 struct { F0 int8; F1 uint16; F2 int }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    var v4 uint = 455
    var v5 uint16 = 42
    if false {
        for fi6 := 0; fi6 < 4; fi6++ {
            v4 += 4294967295
            var v7 int8 = 42
            _ = fi6
            _ = v7
        }
        v4 -= v4
    } else {
        fmt.Println(v5, v5)
        switch "a" {
        case "a":
            if false {
                for fi8 := 0; fi8 < 4; fi8++ {
                    fmt.Println(v4, v4)
                    if ((42 < 127) && true) {
                        var v9 string = "a"
                        _ = v9
                    } else {
                        v4 = (v4 ^ 1)
                        for fi10 := 0; fi10 < 15; fi10++ {
                            switch "hello" {
                            case "a":
                                var v11 bool = (-2.25 < 1.5)
                                v12 := GenS1{F0: -56, F1: 9223372036854775807, F2: v11}
                                switch (("\n\t" + "hello") + ("hello" + "foo bar")) {
                                case "foo bar":
                                    var v13 int8 = -1
                                    _ = v13
                                }
                                _ = v11
                                _ = v12
                            case "":
                                if true {
                                    v14 := []int{}
                                    var v15 bool = false
                                    var v16 float32 = ((0.0 - -2.25) - ((-2.25 * 0.1) - ((0.1 * (0.1 / (float32(-2.25) - (0.0 - float32(float64((3.14159 + float32(1.5)))))))) + 0.0)))
                                    var v17 uint8 = 128
                                    fmt.Println(v15, v4, fi8)
                                    v18 := GenS1{F0: -128, F1: -9223372036854775808, F2: ((((((!false) || ((!(18446744073709551615 > 1)) && ((!("" != ("" + (("hello" + "\n\t") + ("hello" + ""))))) && (42 < 4294967295)))) || false) || (uint8(v4) <= (v17 * (v17 / ((v17) | 1))))) || (42 < 603)) || (false || v15))}
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                } else {
                                    fmt.Println(fi8, v5, fi8)
                                    switch (-(fi10)) {
                                    case 251:
                                        v19 := GenS1{F0: -128, F1: 271, F2: true}
                                        _ = v19
                                    }
                                    switch "" {
                                    case "foo bar":
                                        var v20 uint64 = 425
                                        _ = v20
                                    case "a":
                                        v21 := []int{(^(fi8)), (fi10 | (fi8 | 9223372036854775807)), (fi8 + -2147483648), fi10}
                                        _ = v21
                                    default:
                                        var v22 int = -2147483648
                                        _ = v22
                                    }
                                }
                            case "foo bar":
                                var v23 string = ("hello" + ("" + ("a" + ("foo bar" + ("a" + ((("a" + "a") + "a") + ""))))))
                                _ = v23
                            default:
                                v1["a"] = 2147483648
                            }
                            switch "hello" {
                            case "\n\t":
                                v24 := GenS2{F0: v5, F1: 812, F2: 42}
                                _ = v24
                            case "a":
                                var v25 string = ""
                                v5 += v5
                                _ = v25
                            default:
                                if true {
                                    v26 := GenS1{F0: 0, F1: 9223372036854775807, F2: true}
                                    v27 := GenS3{F0: 42, F1: 1}
                                    v5 += (v5 * (v5 ^ v5))
                                    _ = v26
                                    _ = v27
                                } else {
                                    if (511 != v5) {
                                        v5 += 32767
                                        if false {
                                            var v28 uint = v4
                                            var v29 uint8 = 42
                                            _ = v28
                                            _ = v29
                                        } else {
                                            if true {
                                                v30 := GenS2{F0: v5, F1: 4294967295, F2: -32768}
                                                var v31 float32 = 0.1
                                                delete(v2, "\n\t")
                                                _ = v30
                                                _ = v31
                                            }
                                            v1["a"] = 2147483647
                                        }
                                        fmt.Println(fi8, v4)
                                    } else {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v32 := GenS4{F0: (v5 - v5), F1: -1}
                                            _ = v32
                                        case "":
                                            var v33 uint64 = 4294967295
                                            _ = v33
                                        case "\n\t":
                                            var v34 int64 = int64(fi10)
                                            var v35 uint64 = 418
                                            var v36 int16 = 0
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                        default:
                                            switch "foo bar" {
                                            case "foo bar":
                                                v37 := []int{(-(fi8))}
                                                _ = v37
                                            case "\n\t":
                                                if (100 >= -1) {
                                                    if (1 < 2) {
                                                        switch fi8 {
                                                        case 1:
                                                            fmt.Println(fi8)
                                                            _ = len(v2)
                                                        default:
                                                            var v38 int8 = 42
                                                            _ = v38
                                                        }
                                                        var v39 float64 = -2.25
                                                        v40 := []int{(fi10 % ((fi8) | 1))}
                                                        _ = v39
                                                        _ = v40
                                                    } else {
                                                        _ = len(v3)
                                                        v5 = v5
                                                        v5 += (v5 >> ((24) & 63))
                                                    }
                                                    v41 := GenS5{F0: uint64(v4), F1: v4, F2: v5}
                                                    _ = v41
                                                } else {
                                                    v5 = 437
                                                }
                                            default:
                                                switch (^(fi10)) {
                                                case 32768:
                                                    v4 = (v4 ^ v4)
                                                    v4 = 18446744073709551615
                                                case 850:
                                                    var v42 uint16 = uint16(v4)
                                                    _ = v42
                                                case 42:
                                                    v43 := GenS6{F0: (!(("\n\t" + ("a" + "a")) >= ("" + ("\n\t" + "foo bar"))))}
                                                    _ = v43
                                                default:
                                                    var v44 uint = v4
                                                    if false {
                                                        v5 = v5
                                                        v44 -= uint(v5)
                                                        v4 += 971
                                                    } else {
                                                        delete(v1, ("foo bar" + ("foo bar" + "a")))
                                                        v44 = 65535
                                                    }
                                                    v45 := GenS7{F0: 42}
                                                    _ = v44
                                                    _ = v45
                                                }
                                            }
                                            fmt.Println(fi10)
                                            var v46 uint = v4
                                            _ = v46
                                        }
                                    }
                                    if false {
                                        delete(v1, "")
                                        if false {
                                            if true {
                                                var v47 uint32 = 2147483647
                                                v4 = (^(v4))
                                                _ = v47
                                            }
                                        }
                                        v48 := GenS1{F0: -118, F1: 11, F2: false}
                                        _ = v48
                                    } else {
                                        switch "hello" {
                                        case "\n\t":
                                            if (441 < -32768) {
                                                var v49 int8 = 57
                                                if true {
                                                    switch ("\n\t" + "") {
                                                    case "a":
                                                        var v50 int64 = 1
                                                        _ = v50
                                                    case "foo bar":
                                                        if (!(127 >= 2)) {
                                                            v51 := GenS5{F0: 479, F1: 65535, F2: (v5 & v5)}
                                                            var v52 uint32 = 574
                                                            _ = v51
                                                            _ = v52
                                                        } else {
                                                            switch "" {
                                                            case "\n\t":
                                                                if true {
                                                                    var v53 float64 = -2.25
                                                                    v5 -= (v5 ^ v5)
                                                                    if false {
                                                                        if (false && (!(0 >= 42))) {
                                                                            v54 := []int{(fi8 ^ 9223372036854775807)}
                                                                            _ = v54
                                                                        } else {
                                                                            v55 := GenS4{F0: v5, F1: 0}
                                                                            switch "" {
                                                                            case "":
                                                                                var v56 int32 = -1
                                                                                _ = v56
                                                                            }
                                                                            v57 := []int{fi8, -9223372036854775808, fi10, fi8}
                                                                            _ = v55
                                                                            _ = v57
                                                                        }
                                                                        v58 := []int{(fi10 - -685), -1, fi10, (fi10 & fi8)}
                                                                        _ = len(v58)
                                                                        v49 = v49
                                                                        v59 := GenS6{F0: (128 <= 194)}
                                                                        var v60 uint16 = v5
                                                                        if (!(false && true)) {
                                                                            v4 = v4
                                                                        } else {
                                                                            var v61 int8 = 1
                                                                            v61 = 42
                                                                            v60 = 42
                                                                            _ = v61
                                                                        }
                                                                        _ = len(v58)
                                                                        _ = v58
                                                                        _ = v59
                                                                        _ = v60
                                                                    }
                                                                    _ = v53
                                                                }
                                                                v4 = v4
                                                                v49 -= 127
                                                            case "":
                                                                if (374 < -1) {
                                                                    v62 := GenS8{F0: 0, F1: "hello", F2: v5}
                                                                    var v63 int = (fi10 << ((v4) & 63))
                                                                    v64 := GenS5{F0: 0, F1: 255, F2: 266}
                                                                    _ = v62
                                                                    _ = v63
                                                                    _ = v64
                                                                }
                                                            case "a":
                                                                if ((0.1 - 1e10) < 0.0) {
                                                                    var v65 int = fi10
                                                                    _ = v65
                                                                }
                                                            }
                                                        }
                                                    case "":
                                                        v66 := GenS9{F0: 0, F1: false, F2: 1}
                                                        if true {
                                                            delete(v1, "foo bar")
                                                            v67 := GenS5{F0: 190, F1: (v4 | v4), F2: (^(v5))}
                                                            _ = v67
                                                        }
                                                        var v68 int16 = -1
                                                        _ = v66
                                                        _ = v68
                                                    default:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            var v69 uint64 = 367
                                                            _ = v69
                                                        case "hello":
                                                            var v70 int32 = -1
                                                            _ = v70
                                                        case "a":
                                                            v71 := []int{410}
                                                            _ = v71
                                                        }
                                                        fmt.Println(v5)
                                                        if (true || true) {
                                                            v72 := GenS6{F0: (-2.25 >= (0.0 - ((1.5 - 3.14159) * 0.1)))}
                                                            _ = v72
                                                        } else {
                                                            if false {
                                                                fmt.Println(fi8, fi10, v4)
                                                                var v73 uint32 = 65535
                                                                v74 := []int{fi10, 0, (fi8 << ((428) & 63))}
                                                                _ = len(v1)
                                                                _ = v73
                                                                _ = v74
                                                            } else {
                                                                switch ("\n\t" + "") {
                                                                case "hello":
                                                                    v5 -= 255
                                                                case "foo bar":
                                                                    v5 = v5
                                                                default:
                                                                    v75 := GenS10{F0: "", F1: (fi10 | (-(fi10)))}
                                                                    _ = v75
                                                                }
                                                                v49 = 127
                                                                v76 := []int{(fi8 ^ (fi8 / ((-9223372036854775808) | 1)))}
                                                                var v77 uint32 = 662
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                        }
                                                        switch fi10 {
                                                        case -556:
                                                            if true {
                                                                v78 := []int{fi10}
                                                                v1[""] = 782
                                                                fmt.Println(v49, fi8, fi8)
                                                                var v79 string = ("foo bar" + ("hello" + "a"))
                                                                v80 := GenS2{F0: 42, F1: 1, F2: 42}
                                                                _ = v78
                                                                _ = v79
                                                                _ = v80
                                                            }
                                                        case -776:
                                                            switch (fi10 | fi10) {
                                                            case 32767:
                                                                var v81 string = "a"
                                                                _ = v81
                                                            case -272:
                                                                v49 = int8(v4)
                                                            }
                                                        default:
                                                            var v82 int64 = -9223372036854775808
                                                            _ = v82
                                                        }
                                                    }
                                                }
                                                _ = v49
                                            } else {
                                                var v83 uint8 = 128
                                                _ = v83
                                            }
                                            fmt.Println(v4, v5, fi8)
                                            fmt.Println(fi8)
                                        case "":
                                            var v84 uint32 = 2147483648
                                            _ = v84
                                        }
                                        v4 = (^(v4))
                                    }
                                }
                            }
                            v5 -= (^(v5))
                            _ = fi10
                        }
                    }
                    v85 := GenS11{F0: 2}
                    _ = fi8
                    _ = v85
                }
                v4 += uint(v5)
            } else {
                v4 = 132
            }
        default:
            var v86 int = 127
            _ = v86
        }
    }
    for fi87 := 0; fi87 < 12; fi87++ {
        v4 = v4
        if true {
            if true {
                v5 = (v5 % ((v5) | 1))
                var v88 int16 = 1000
                _ = v88
            } else {
                if false {
                    v89 := GenS12{F0: 127, F1: -9223372036854775808}
                    _ = v89
                } else {
                    _ = len(v1)
                    switch ("\n\t" + "") {
                    case "hello":
                        for fi90 := 0; fi90 < 9; fi90++ {
                            v91 := GenS13{F0: -1, F1: 0, F2: -236}
                            v92 := GenS10{F0: ("" + ("a" + ("" + ""))), F1: (fi90 >> (((v4 * (v4 & (v4 * 974)))) & 63))}
                            v93 := []int{1}
                            var v94 int64 = -187
                            var v95 uint32 = 256
                            fmt.Println(v5, fi87, v4)
                            if (false || ((-244 <= 1) || (("foo bar" + "") == "\n\t"))) {
                                v95 = v95
                            } else {
                                v5 -= 95
                                var v96 uint = 255
                                if true {
                                    switch fi90 {
                                    case 2147483647:
                                        if false {
                                            var v97 int32 = 42
                                            _ = v97
                                        }
                                    default:
                                        v98 := v3[-9223372036854775808]
                                        _ = v98
                                    }
                                }
                                _ = v96
                            }
                            _ = fi90
                            _ = v91
                            _ = v92
                            _ = v93
                            _ = v94
                            _ = v95
                        }
                    case "a":
                        var v99 string = "\n\t"
                        _ = v99
                    default:
                        v5 -= 375
                    }
                    fmt.Println(v4)
                }
            }
            fmt.Println(v5)
        }
        _ = fi87
    }
    v100 := GenS10{F0: ("foo bar" + "hello"), F1: -359}
    v101 := []int{719, 128, 833}
    switch 0 {
    case 972:
        if true {
            for fi102 := 0; fi102 < 3; fi102++ {
                v103 := GenS11{F0: 2147483648}
                fmt.Println(v5)
                _ = fi102
                _ = v103
            }
            var v104 int32 = 1
            _ = v104
        }
    case 32768:
        delete(v3, 2147483647)
    case 919:
        switch 224 {
        case 4294967295:
            _ = len(v101)
        default:
            var v105 int = -9223372036854775808
            _ = v105
        }
    default:
        var v106 uint16 = 0
        _ = v106
    }
    v107 := GenS10{F0: ("a" + ("foo bar" + ("\n\t" + ("foo bar" + (("\n\t" + ("a" + "")) + "foo bar"))))), F1: 42}
    fmt.Println(v4, v5)
    for fi108 := 0; fi108 < 3; fi108++ {
        v4 += 18446744073709551615
        _ = fi108
    }
    v5 -= v5
    for fi109 := 0; fi109 < 15; fi109++ {
        var v110 int16 = 0
        for fi111 := 0; fi111 < 12; fi111++ {
            v5 = (v5 / (((v5 | 42)) | 1))
            _ = fi111
        }
        _ = fi109
        _ = v110
    }
    if (!true) {
        if true {
            _ = len(v1)
            for fi112 := 0; fi112 < 15; fi112++ {
                _ = len(v101)
                v3[(fi112 << (((v4 / ((v4) | 1))) & 63))] = 42
                switch "\n\t" {
                case "a":
                    v101 = append(v101, (fi112 - fi112))
                }
                _ = fi112
            }
        }
        v101 = append(v101, 32767)
    } else {
        v113 := GenS14{F0: 117, F1: v5, F2: -9223372036854775808}
        var v114 int64 = -283
        _ = v113
        _ = v114
    }
    v5 += 42
    var v115 int8 = 100
    v115 -= (v115 | -128)
    var v116 float64 = -2.25
    if (1000 == -1) {
        for fi117 := 0; fi117 < 4; fi117++ {
            if false {
                var v118 uint32 = 898
                v115 += v115
                for fi119 := 0; fi119 < 1; fi119++ {
                    switch fi117 {
                    case -1:
                        v116 -= ((v116 + v116) * v116)
                    case 9223372036854775807:
                        if false {
                            var v120 int = fi119
                            _ = v120
                        }
                    }
                    v5 += v5
                    var v121 int64 = 9223372036854775807
                    _ = fi119
                    _ = v121
                }
                _ = v118
            }
            v115 -= (v115 % ((v115) | 1))
            switch (fi117 / ((fi117) | 1)) {
            case 128:
                var v122 string = "foo bar"
                _ = v122
            case 9223372036854775807:
                var v123 uint8 = 255
                _ = v123
            case 32768:
                fmt.Println(v4, v115, v4)
            }
            _ = fi117
        }
    }
    if len(v101) > 0 {
        v101[0] = -949
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v101:", len(v101), v101)
    fmt.Println("v115:", v115)
    fmt.Println("v116:", v116)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v100
    _ = v101
    _ = v107
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
