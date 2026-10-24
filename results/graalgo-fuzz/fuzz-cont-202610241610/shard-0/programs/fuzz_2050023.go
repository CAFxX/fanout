package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16; F1 uint32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint64; F1 string; F2 uint32 }

func main() {
    var v1 uint8 = 255
    var v2 uint = 4294967295
    v3 := make(map[string]bool)
    var v4 int = 32767
    var v5 uint8 = (v1 - 1)
    var v6 float32 = (1e10 * 1e10)
    for fi7 := 0; fi7 < 8; fi7++ {
        switch (fi7 >> ((v2) & 63)) {
        case 32768:
            v5 += 127
        case 869:
            if ((v5 % ((0) | 1)) < 0) {
                _ = len(v3)
                var v8 uint32 = 660
                fmt.Println(v8, v8)
                _ = v8
            } else {
                if false {
                    for fi9 := 0; fi9 < 18; fi9++ {
                        switch (((("foo bar" + ("" + ("hello" + "hello"))) + ("foo bar" + "hello")) + "hello") + "hello") {
                        case "foo bar":
                            v1 = 42
                        case "hello":
                            var v10 bool = (-128 <= -1)
                            _ = v10
                        case "\n\t":
                            var v11 int16 = 1000
                            _ = v11
                        default:
                            if (!((!(true || (715 <= 32767))) || (((^(v2)) >= (v2 % ((1) | 1))) || true))) {
                                delete(v3, ("hello" + ("" + "foo bar")))
                                v5 -= v5
                            }
                        }
                        _ = fi9
                    }
                    for fi12 := 0; fi12 < 6; fi12++ {
                        if (0 != -937) {
                            var v13 bool = true
                            var v14 int16 = -1
                            if ((1 >= 681) && v13) {
                                if (-563 > 42) {
                                    v15 := GenS1{F0: 42, F1: 1}
                                    var v16 float64 = (3.14159 - 0.0)
                                    if v13 {
                                        v1 -= v1
                                        delete(v3, "")
                                        var v17 bool = (59 >= -685)
                                        _ = v17
                                    }
                                    _ = v15
                                    _ = v16
                                }
                                if (false || (1 == 1)) {
                                    if false {
                                        if ("foo bar" > "foo bar") {
                                            switch (fi7 * (^(v4))) {
                                            case 925:
                                                var v18 int32 = int32(fi12)
                                                _ = v18
                                            case 497:
                                                fmt.Println(fi12)
                                            }
                                            var v19 uint32 = 2147483647
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v20 uint32 = v19
                                                _ = v20
                                            }
                                            _ = v19
                                        }
                                        fmt.Println(v6, v1)
                                    }
                                } else {
                                    v4 = -462
                                }
                                var v21 uint64 = 1
                                _ = v21
                            } else {
                                var v22 int16 = v14
                                if (0 > 71) {
                                    if ("" != "hello") {
                                        v4 = (fi12 * v4)
                                        v22 += (v14 ^ int16(v4))
                                        var v23 float32 = v6
                                        _ = v23
                                    } else {
                                        fmt.Println(v2, v1)
                                    }
                                    var v24 bool = (!false)
                                    var v25 uint8 = (v1 - 2)
                                    _ = v24
                                    _ = v25
                                }
                                v3["hello"] = (-320 != 42)
                                var v26 float64 = 0.0
                                v4 = 128
                                var v27 int32 = 2147483647
                                var v28 uint = uint(v5)
                                fmt.Println(v13, v22, fi7)
                                v29 := GenS1{F0: 1, F1: 589}
                                _ = v22
                                _ = v26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            if (v13 || false) {
                                fmt.Println(fi7, v13)
                                v30 := GenS1{F0: -662, F1: 255}
                                v1 = (v5 >> ((v2) & 63))
                                if v13 {
                                    v14 += v14
                                    switch "a" {
                                    case "hello":
                                        switch 25 {
                                        case 32768:
                                            if v13 {
                                                v31 := GenS1{F0: 0, F1: 0}
                                                delete(v3, ((("hello" + "") + ("a" + ((("" + "a") + (("foo bar" + ("hello" + "a")) + (("" + "hello") + "foo bar"))) + "\n\t"))) + "a"))
                                                _ = v31
                                            } else {
                                                v4 -= (fi7 & int(v1))
                                                v32 := GenS1{F0: -1, F1: 255}
                                                _ = v32
                                            }
                                        case 32767:
                                            var v33 uint = (v2 >> ((v2) & 63))
                                            _ = v33
                                        case -784:
                                            v34 := GenS1{F0: 463, F1: 32767}
                                            _ = v34
                                        }
                                    }
                                    v35 := []int{(^(v4)), 2147483647, fi7, fi12}
                                    _ = v35
                                } else {
                                    v36 := GenS1{F0: -868, F1: 2}
                                    v37 := GenS2{F0: 4294967295, F1: 0}
                                    if ((v4 * fi12) == fi7) {
                                        v4 = (fi12 | (fi7 + (v4 | -1)))
                                        var v38 bool = false
                                        v39 := GenS3{F0: (^(v1)), F1: v5, F2: (v14 | -1)}
                                        _ = v38
                                        _ = v39
                                    }
                                    v14 -= 454
                                    var v40 uint16 = 2
                                    if v13 {
                                        switch (^(fi7)) {
                                        case 2147483647:
                                            v13 = (("a" + "") <= ("foo bar" + "hello"))
                                        case -9223372036854775808:
                                            switch fi12 {
                                            case -2147483648:
                                                if (-1 < -1) {
                                                    var v41 string = "foo bar"
                                                    v42 := v3[v41]
                                                    var v43 int32 = -1
                                                    switch int(v2) {
                                                    case 4294967295:
                                                        switch (fi7 & (v4 - (v4 ^ v4))) {
                                                        case 1:
                                                            switch 1 {
                                                            case -569:
                                                                v44 := GenS4{F0: 4294967295, F1: v40}
                                                                _ = v44
                                                            case 4294967295:
                                                                _ = len(v3)
                                                            }
                                                        case 42:
                                                            var v45 uint = v2
                                                            _ = v45
                                                        case 785:
                                                            var v46 uint64 = 301
                                                            _ = v46
                                                        default:
                                                            if v42 {
                                                                if v42 {
                                                                    v42 = (-2.25 >= 3.14159)
                                                                    delete(v3, "")
                                                                }
                                                                v13 = (65535 <= (^(v2)))
                                                            } else {
                                                                if ((("\n\t" < "a") || (v5 != v5)) && v42) {
                                                                    if (42 != 414) {
                                                                        v4 = (v4 | (v4 ^ fi7))
                                                                        if false {
                                                                            v2 = (v2 - (v2 * v2))
                                                                            v47 := []int{(fi7 * v4), (-(v4)), fi12}
                                                                            _ = v47
                                                                        } else {
                                                                            if (18446744073709551615 < 160) {
                                                                                v48 := GenS5{F0: v42, F1: -1}
                                                                                _ = v48
                                                                            } else {
                                                                                var v49 int64 = 1
                                                                                _ = v49
                                                                            }
                                                                            v2 += (v2 - v2)
                                                                        }
                                                                        v5 = v5
                                                                    }
                                                                    v42 = v42
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        v50 := v3[v41]
                                                        _ = v50
                                                    }
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                            case -188:
                                                switch "" {
                                                case "":
                                                    v51 := GenS6{F0: 151, F1: uint32(v2)}
                                                    _ = v51
                                                case "foo bar":
                                                    v52 := GenS6{F0: v2, F1: 0}
                                                    _ = v52
                                                case "hello":
                                                    switch "\n\t" {
                                                    case "":
                                                        switch (("a" + ("hello" + "hello")) + ("foo bar" + "")) {
                                                        case "\n\t":
                                                            v53 := GenS7{F0: 0, F1: (v40 << (((v2 * (v2 / ((v2) | 1)))) & 63)), F2: -9223372036854775808}
                                                            _ = v53
                                                        default:
                                                            var v54 float64 = 3.14159
                                                            _ = v54
                                                        }
                                                    case "hello":
                                                        var v55 uint32 = 517
                                                        _ = v55
                                                    default:
                                                        var v56 uint16 = v40
                                                        _ = v56
                                                    }
                                                    var v57 string = (("a" + ("\n\t" + ("hello" + "\n\t"))) + "\n\t")
                                                    _ = v57
                                                }
                                            }
                                        }
                                        var v58 int32 = -2147483648
                                        if v13 {
                                            if (!(4294967295 == 4294967295)) {
                                                if ((!false) || true) {
                                                    var v59 string = "a"
                                                    _ = v59
                                                } else {
                                                    v60 := GenS5{F0: v13, F1: 42}
                                                    if true {
                                                        var v61 float64 = float64(v6)
                                                        switch "hello" {
                                                        case "hello":
                                                            v13 = ((true && v13) && v13)
                                                        }
                                                        v6 -= (float32((v61 + 0.1)) - v6)
                                                        if v13 {
                                                            v13 = (!v13)
                                                            if (272 == 467) {
                                                                if (!(v13 && (0.1 <= (v6 + v6)))) {
                                                                    v62 := GenS7{F0: (^(v5)), F1: v40, F2: -1}
                                                                    _ = v62
                                                                }
                                                                v63 := GenS4{F0: 0, F1: v40}
                                                                _ = v63
                                                            }
                                                            if (!(!(639 > 2))) {
                                                                delete(v3, "foo bar")
                                                                var v64 uint16 = v40
                                                                v65 := GenS3{F0: v5, F1: v5, F2: (^(v14))}
                                                                v66 := []int{}
                                                                fmt.Println(v2, v40, fi12)
                                                                v1 += uint8(v2)
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                            var v67 uint16 = v40
                                                            var v68 int16 = -886
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                v6 -= -2.25
                                                            case "":
                                                                v69 := []int{(^(fi7))}
                                                                _ = v69
                                                            }
                                                            _ = v67
                                                            _ = v68
                                                        } else {
                                                            switch ("foo bar" + ("a" + "foo bar")) {
                                                            case "a":
                                                                v70 := GenS8{F0: 4294967295}
                                                                _ = v70
                                                            case "":
                                                                if v13 {
                                                                    var v71 uint8 = 0
                                                                    _ = v71
                                                                }
                                                            }
                                                            if (v13 && (float64(v6) > (1e10 - v61))) {
                                                                v58 -= v58
                                                                v61 -= -2.25
                                                                fmt.Println(v4, v5, v1)
                                                            }
                                                            switch ("\n\t" + ("a" + "hello")) {
                                                            case "":
                                                                v5 = uint8(v2)
                                                            case "\n\t":
                                                                v2 = (v2 / ((v2) | 1))
                                                            case "foo bar":
                                                                if v13 {
                                                                    v4 += 24
                                                                    switch fi7 {
                                                                    case 42:
                                                                        v14 = v14
                                                                    default:
                                                                        var v72 int64 = 213
                                                                        _ = v72
                                                                    }
                                                                } else {
                                                                    v73 := v3[""]
                                                                    var v74 uint = 591
                                                                    _ = v73
                                                                    _ = v74
                                                                }
                                                            }
                                                        }
                                                        var v75 string = ("" + "\n\t")
                                                        _ = v61
                                                        _ = v75
                                                    } else {
                                                        v76 := GenS9{F0: 0, F1: 1}
                                                        var v77 float64 = float64(1.5)
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    v2 = (^(v2))
                                                    _ = v60
                                                }
                                                v40 = (v40 % ((v40) | 1))
                                                fmt.Println(v2, v14, v1)
                                                v13 = (v2 >= 680)
                                            } else {
                                                v78 := GenS6{F0: v2, F1: 1}
                                                _ = v78
                                            }
                                            v5 = v5
                                            var v79 uint = 255
                                            _ = v79
                                        }
                                        _ = v58
                                    } else {
                                        v80 := GenS10{F0: -327}
                                        _ = v80
                                    }
                                    fmt.Println(v4, v2, v5)
                                    _ = v36
                                    _ = v37
                                    _ = v40
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    v6 += 1e10
                                case "foo bar":
                                    var v81 string = ("\n\t" + ((("foo bar" + "a") + ("hello" + ("" + "foo bar"))) + "hello"))
                                    _ = v81
                                default:
                                    if ((2147483647 > 316) || v13) {
                                        v82 := v3[""]
                                        _ = v82
                                    }
                                }
                                _ = v30
                            }
                            fmt.Println(fi12, v4)
                            _ = v13
                            _ = v14
                        }
                        v83 := GenS4{F0: 42, F1: 2}
                        v2 = (v2 % ((v2) | 1))
                        _ = fi12
                        _ = v83
                    }
                } else {
                    fmt.Println(v6, v4)
                    v5 = v1
                }
            }
        case 944:
            v84 := []int{(fi7 | (v4 >> (((v2 - v2)) & 63))), fi7, (v4 & fi7), (^(v4))}
            _ = v84
        }
        if (!false) {
            var v85 int = v4
            v1 -= (v5 & v1)
            fmt.Println(v6, v4, v85)
            _ = v85
        } else {
            if true {
                v1 -= 1
            } else {
                var v86 string = ""
                v87 := GenS11{F0: (v2 ^ v2)}
                switch (v86 + (v86 + v86)) {
                case "":
                    v88 := []int{fi7, (fi7 / ((v4) | 1)), v4, v4}
                    _ = v88
                case "hello":
                    var v89 float64 = ((-2.25 * 1e10) - 0.0)
                    _ = v89
                }
                _ = v86
                _ = v87
            }
            switch "hello" {
            case "hello":
                if false {
                    var v90 int8 = 100
                    switch ("foo bar" + "foo bar") {
                    case "\n\t":
                        v91 := []int{v4, (fi7 & (v4 - 128)), (fi7 + v4)}
                        _ = v91
                    case "a":
                        v92 := []int{2147483647, v4, v4, (fi7 >> (((v2 | 18446744073709551615)) & 63))}
                        _ = v92
                    }
                    _ = v90
                }
            default:
                _ = len(v3)
            }
        }
        _ = fi7
    }
    if (1 < 100) {
        fmt.Println(v6, v6)
    }
    v4 = int(v2)
    var v93 uint64 = 0
    if (!false) {
        var v94 int32 = -882
        _ = v94
    } else {
        fmt.Println(v93, v93)
    }
    fmt.Println(v5, v93)
    v95 := GenS1{F0: 9223372036854775807, F1: 65535}
    for fi96 := 0; fi96 < 8; fi96++ {
        _ = len(v3)
        _ = fi96
    }
    switch ("foo bar" + "") {
    case "hello":
        var v97 uint16 = uint16(v93)
        _ = v97
    case "\n\t":
        v98 := GenS2{F0: 31, F1: -23}
        _ = v98
    default:
        v99 := GenS3{F0: v1, F1: v1, F2: -796}
        _ = v99
    }
    v93 = v93
    if (!false) {
        var v100 int32 = -375
        if (32767 <= uint16(v93)) {
            if false {
                for fi101 := 0; fi101 < 16; fi101++ {
                    if (uint64(v1) != v93) {
                        var v102 string = "\n\t"
                        _ = v102
                    } else {
                        if false {
                            v103 := GenS12{F0: v5}
                            _ = v103
                        } else {
                            for fi104 := 0; fi104 < 7; fi104++ {
                                var v105 uint = (v2 >> (((^(v2))) & 63))
                                v106 := []int{(fi101 % ((fi101) | 1)), (fi101 ^ (fi101 + (-(fi101)))), fi101, v4}
                                v107 := GenS10{F0: 792}
                                _ = fi104
                                _ = v105
                                _ = v106
                                _ = v107
                            }
                            v108 := GenS13{F0: "\n\t"}
                            _ = v108
                        }
                        fmt.Println(v4)
                    }
                    for fi109 := 0; fi109 < 10; fi109++ {
                        v100 = 1
                        v4 = (fi101 | fi109)
                        v93 = (v93 | v93)
                        _ = fi109
                    }
                    _ = fi101
                }
            }
            v93 = (v93 | 710)
            for fi110 := 0; fi110 < 1; fi110++ {
                for fi111 := 0; fi111 < 20; fi111++ {
                    var v112 uint16 = 2
                    var v113 int16 = 1
                    v112 = (v112 % (((v112 + 255)) | 1))
                    _ = fi111
                    _ = v112
                    _ = v113
                }
                v114 := GenS1{F0: 941, F1: 1}
                _ = fi110
                _ = v114
            }
        }
        v115 := GenS4{F0: 4294967295, F1: 510}
        _ = v100
        _ = v115
    } else {
        delete(v3, "hello")
        fmt.Println(v6, v6, v93)
    }
    v116 := v3[("foo bar" + "hello")]
    v117 := GenS14{F0: v93, F1: "hello", F2: 65535}
    v3["a"] = v116
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v93
    _ = v95
}
