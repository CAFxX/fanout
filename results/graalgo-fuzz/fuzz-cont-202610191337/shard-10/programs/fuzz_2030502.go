package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint8; F1 uint32; F2 uint }
type GenS6 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS7 struct { F0 string; F1 int8; F2 int32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint64; F1 bool; F2 int }
type GenS10 struct { F0 uint8; F1 int64 }
type GenS11 struct { F0 bool; F1 uint64 }
type GenS12 struct { F0 uint8; F1 uint8; F2 bool }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{"", ("" + "a"), ("" + "foo bar"), "foo bar"}
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]int)
    var v6 float64 = float64(1e10)
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v3 = (^(v3))
    if ((-9223372036854775808 == 324) && (true && (!(42 > 255)))) {
        var v7 float32 = 1.5
        v2 = append(v2, "")
        if false {
            if false {
                if (255 < 2) {
                    v2 = append(v2, "foo bar")
                }
            } else {
                var v8 int64 = 9223372036854775807
                v2 = append(v2, "")
                var v9 string = ("a" + ("foo bar" + ("foo bar" + (("hello" + ((("foo bar" + "foo bar") + ("a" + "")) + "")) + "foo bar"))))
                _ = v8
                _ = v9
            }
            v6 = (float64(-2.25) + float64((v7 / 1.0)))
            fmt.Println(v6, v3, v6)
        } else {
            for fi10 := 0; fi10 < 19; fi10++ {
                if (!(0 > (^(v3)))) {
                    _ = len(v5)
                    for fi11 := 0; fi11 < 16; fi11++ {
                        v6 += v6
                        switch "\n\t" {
                        case "":
                            v3 = (v3 + 42)
                        case "a":
                            var v12 float64 = v6
                            _ = v12
                        default:
                            v7 -= 3.14159
                        }
                        if true {
                            if ((true || true) || (!(!(846 < 707)))) {
                                if (1 < -1) {
                                    fmt.Println(v7, v7, v3)
                                    switch "hello" {
                                    case "foo bar":
                                        switch "foo bar" {
                                        case "a":
                                            switch "hello" {
                                            case "hello":
                                                v13 := GenS1{F0: (-1 > 0)}
                                                _ = v13
                                            case "":
                                                if (((true && (-121 <= 25)) && (true && false)) && (float64((3.14159 * v7)) > v6)) {
                                                    var v14 int = (fi10 | fi11)
                                                    var v15 int64 = 9223372036854775807
                                                    var v16 string = "hello"
                                                    _ = v14
                                                    _ = v15
                                                    _ = v16
                                                }
                                            }
                                        case "":
                                            v17 := GenS1{F0: true}
                                            _ = v17
                                        case "foo bar":
                                            if ((^(v3)) >= v3) {
                                                v2 = append(v2, "a")
                                            }
                                        }
                                    case "a":
                                        switch ("" + ("hello" + "foo bar")) {
                                        case "":
                                            var v18 float32 = v7
                                            _ = v18
                                        case "\n\t":
                                            var v19 float64 = 3.14159
                                            _ = v19
                                        }
                                    case "\n\t":
                                        switch (((("\n\t" + ("foo bar" + "")) + ("\n\t" + "a")) + ((("" + "") + ((("a" + "hello") + "hello") + "\n\t")) + "\n\t")) + "foo bar") {
                                        case "":
                                            if false {
                                                fmt.Println(fi11, v3)
                                                var v20 int = fi11
                                                _ = v20
                                            } else {
                                                fmt.Println(v7, v7)
                                                v21 := GenS2{F0: 18446744073709551615, F1: 99, F2: 1}
                                                if true {
                                                    v22 := GenS3{F0: 0, F1: (-(fi11))}
                                                    _ = v22
                                                } else {
                                                    var v23 float64 = (((v6 + (0.0 - (float64(v7) - v6))) / float64(v7)) / v6)
                                                    var v24 float32 = ((v7 / v7) * 1e10)
                                                    _ = v23
                                                    _ = v24
                                                }
                                                _ = v21
                                            }
                                        case "hello":
                                            switch fi11 {
                                            case 128:
                                                var v25 uint64 = 1
                                                _ = v25
                                            case 768:
                                                if (2 != 0) {
                                                    if (true || false) {
                                                        v26 := GenS4{F0: 255, F1: 100}
                                                        v27 := v1[(("hello" + "a") + ("foo bar" + ""))]
                                                        _ = v26
                                                        _ = v27
                                                    }
                                                }
                                            case 2147483647:
                                                var v28 int8 = 29
                                                _ = v28
                                            }
                                        }
                                    default:
                                        var v29 float32 = (v7 - float32(1e10))
                                        _ = v29
                                    }
                                } else {
                                    if false {
                                        if ((0 >= 961) && true) {
                                            if true {
                                                v2 = append(v2, ("hello" + "a"))
                                                _ = len(v1)
                                                var v30 uint16 = 0
                                                _ = v30
                                            }
                                            if (fi11 < (fi11 ^ 42)) {
                                                if (!(!((false && true) || (true || (true || (false && ("foo bar" >= "hello"))))))) {
                                                    switch (fi10 << ((18446744073709551615) & 63)) {
                                                    case 9223372036854775807:
                                                        fmt.Println(fi11, fi10)
                                                    }
                                                    var v31 int64 = 1
                                                    v32 := GenS3{F0: 1, F1: -2147483648}
                                                    _ = v31
                                                    _ = v32
                                                } else {
                                                    if false {
                                                        fmt.Println(v3, v7, fi11)
                                                        v33 := GenS3{F0: 2, F1: (-(fi10))}
                                                        v7 = 1.5
                                                        _ = v33
                                                    }
                                                    var v34 int8 = 0
                                                    v34 += (v34 >> ((2) & 63))
                                                    _ = v34
                                                }
                                                switch "hello" {
                                                case "\n\t":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                fmt.Println(v3)
                                            } else {
                                                v5[42] = fi10
                                                if true {
                                                    fmt.Println(fi10, v3, v3)
                                                    _ = len(v1)
                                                }
                                                var v35 int64 = 1
                                                _ = v35
                                            }
                                            v36 := GenS1{F0: ((1 > 609) || true)}
                                            var v37 uint32 = 2147483647
                                            v38 := GenS4{F0: 127, F1: 127}
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        } else {
                                            var v39 uint16 = 42
                                            if (255 <= 0) {
                                                if true {
                                                    switch "foo bar" {
                                                    case "":
                                                        if false {
                                                            v40 := GenS3{F0: 18446744073709551615, F1: (fi10 << (((v3 ^ v3)) & 63))}
                                                            _ = v40
                                                        } else {
                                                            if false {
                                                                var v41 bool = true
                                                                v39 += v39
                                                                var v42 float64 = v6
                                                                _ = v41
                                                                _ = v42
                                                            } else {
                                                                var v43 int8 = 42
                                                                _ = v43
                                                            }
                                                            v44 := GenS3{F0: 18446744073709551615, F1: (^(fi11))}
                                                            switch fi10 {
                                                            case -635:
                                                                if false {
                                                                    v2 = append(v2, "foo bar")
                                                                    v4[2147483648] = 996
                                                                    v7 += (v7 / (3.14159 / v7))
                                                                }
                                                            default:
                                                                var v45 int16 = -32768
                                                                _ = v45
                                                            }
                                                            _ = v44
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v39, v3)
                                                    case "\n\t":
                                                        if ((fi10 % (((fi10 << (((v3 & v3)) & 63))) | 1)) <= fi10) {
                                                            v6 = v6
                                                            if len(v2) > 0 {
                                                                v2[0] = "\n\t"
                                                            }
                                                        }
                                                    }
                                                }
                                                v5[uint32(v3)] = fi10
                                                var v46 int32 = -1
                                                _ = v46
                                            } else {
                                                if len(v2) > 0 {
                                                    v2[0] = ""
                                                }
                                                v47 := GenS2{F0: 255, F1: 2147483647, F2: 127}
                                                if (true || (1.5 != v7)) {
                                                    v2 = append(v2, "foo bar")
                                                    var v48 uint = 679
                                                    _ = v48
                                                } else {
                                                    if false {
                                                        v49 := GenS1{F0: true}
                                                        var v50 uint = v3
                                                        var v51 int8 = 1
                                                        _ = v49
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        v52 := GenS2{F0: 42, F1: -1, F2: 42}
                                                        v2 = append(v2, "")
                                                        _ = v52
                                                    }
                                                    var v53 float32 = ((v7 / float32(((float64(v7) / 1e10) - ((float64(float32((v6 / (v6 + v6)))) - v6) - (1.5 + 1.5))))) / 1e10)
                                                    v54 := GenS5{F0: 42, F1: 1, F2: v3}
                                                    _ = v53
                                                    _ = v54
                                                }
                                                _ = v47
                                            }
                                            switch "\n\t" {
                                            case "foo bar":
                                                _ = len(v1)
                                            }
                                            _ = v39
                                        }
                                    } else {
                                        if false {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v55 := GenS4{F0: 43, F1: 42}
                                            switch fi10 {
                                            case 9223372036854775807:
                                                if ((!false) || ((v7 / float32(1.5)) != v7)) {
                                                    if (!((true || (-32768 != 244)) || (1 >= 0))) {
                                                        v3 -= v3
                                                        switch 125 {
                                                        case -9223372036854775808:
                                                            if false {
                                                                var v56 uint64 = 42
                                                                fmt.Println(v56, fi11, v3)
                                                                _ = v56
                                                            }
                                                        default:
                                                            if true {
                                                                v3 = 255
                                                                if ((((false && ((554 == 42) && false)) || (!true)) && (9223372036854775807 < -1)) && ((v3 < (v3 + 1)) && false)) {
                                                                    var v57 bool = true
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v58 := GenS4{F0: 127, F1: 1}
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    var v59 int8 = 75
                                                                    var v60 uint8 = 2
                                                                    v61 := GenS6{F0: 1, F1: 42, F2: 2147483647}
                                                                    _ = v59
                                                                    _ = v60
                                                                    _ = v61
                                                                }
                                                                var v62 string = ("hello" + ((("" + "a") + ("" + "hello")) + "\n\t"))
                                                                _ = v62
                                                            }
                                                        }
                                                        var v63 float64 = 0.1
                                                        _ = v63
                                                    } else {
                                                        v3 = 42
                                                        if ((false || false) && true) {
                                                            v7 -= 3.14159
                                                        } else {
                                                            v3 = (v3 + v3)
                                                            var v64 int16 = 32767
                                                            var v65 int16 = 1000
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                    }
                                                    delete(v5, 2)
                                                } else {
                                                    fmt.Println(v3, v3, fi11)
                                                    v66 := GenS1{F0: false}
                                                    _ = v66
                                                }
                                            default:
                                                if len(v2) > 0 {
                                                    v2[0] = "hello"
                                                }
                                            }
                                            _ = v55
                                        } else {
                                            switch "hello" {
                                            case "hello":
                                                if true {
                                                    if false {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v3 = 65535
                                                    } else {
                                                        v67 := GenS2{F0: v3, F1: -810, F2: 128}
                                                        v3 -= v3
                                                        v5[625] = (-(fi11))
                                                        _ = len(v2)
                                                        _ = v67
                                                    }
                                                    if true {
                                                        v68 := GenS3{F0: 4294967295, F1: fi11}
                                                        v7 += 0.1
                                                        v69 := GenS4{F0: 131, F1: 42}
                                                        v3 = (v3 ^ (v3 / ((v3) | 1)))
                                                        if true {
                                                            v6 -= float64(v7)
                                                            switch ("" + ("a" + "foo bar")) {
                                                            case "hello":
                                                                fmt.Println(v7, v3, v3)
                                                                v3 = 842
                                                                var v70 int = fi10
                                                                _ = v70
                                                            case "foo bar":
                                                                switch (fi10 << ((v3) & 63)) {
                                                                case 32768:
                                                                    v6 = v6
                                                                default:
                                                                    fmt.Println(v3, fi11, v3)
                                                                }
                                                            }
                                                            if (!(false && ((!false) && (v3 == v3)))) {
                                                                v71 := GenS4{F0: 139, F1: -128}
                                                                var v72 uint = (^(v3))
                                                                _ = v71
                                                                _ = v72
                                                            }
                                                        }
                                                        _ = v68
                                                        _ = v69
                                                    } else {
                                                        fmt.Println(v6)
                                                        var v73 int8 = -1
                                                        switch ("foo bar" + "") {
                                                        case "hello":
                                                            if len(v2) > 0 {
                                                                v2[0] = "foo bar"
                                                            }
                                                            v74 := GenS7{F0: (("hello" + "foo bar") + ((("\n\t" + "\n\t") + ("a" + "foo bar")) + "a")), F1: v73, F2: 0}
                                                            v7 = float32(0.1)
                                                            _ = v74
                                                        }
                                                        _ = v73
                                                    }
                                                    v6 = 1.5
                                                } else {
                                                    fmt.Println(v3)
                                                    switch -2147483648 {
                                                    case 0:
                                                        v75 := GenS8{F0: false}
                                                        if true {
                                                            var v76 uint8 = 128
                                                            _ = v76
                                                        } else {
                                                            var v77 uint = 25
                                                            _ = v77
                                                        }
                                                        switch "" {
                                                        case "a":
                                                            v78 := v4[39]
                                                            _ = v78
                                                        case "foo bar":
                                                            v3 = v3
                                                        }
                                                        _ = v75
                                                    case 1:
                                                        var v79 uint8 = 6
                                                        _ = v79
                                                    }
                                                }
                                            case "\n\t":
                                                if (float32(0.0) > -2.25) {
                                                    if false {
                                                        v80 := GenS4{F0: 205, F1: -63}
                                                        _ = v80
                                                    } else {
                                                        v3 = (v3 * v3)
                                                        v6 += v6
                                                        if (1 <= 69) {
                                                            fmt.Println(v3)
                                                            if false {
                                                                var v81 int32 = 645
                                                                var v82 float32 = ((v7 - (-2.25 + v7)) / v7)
                                                                _ = v81
                                                                _ = v82
                                                            }
                                                            var v83 bool = (true || (true || true))
                                                            _ = v83
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "":
                                                        v3 = v3
                                                    }
                                                    v84 := GenS9{F0: 439, F1: (!true), F2: 9223372036854775807}
                                                    _ = v84
                                                }
                                            case "foo bar":
                                                if (!false) {
                                                    var v85 int8 = 1
                                                    delete(v4, 2147483648)
                                                    v86 := v5[363]
                                                    _ = v85
                                                    _ = v86
                                                }
                                            }
                                            v7 -= ((float32(1.5) * v7) / v7)
                                            v7 += v7
                                        }
                                        v6 = v6
                                        if (false && true) {
                                            var v87 bool = true
                                            if false {
                                                if (!v87) {
                                                    v87 = v87
                                                    switch 1 {
                                                    case 42:
                                                        v3 = 2
                                                    default:
                                                        switch -1 {
                                                        case -9223372036854775808:
                                                            var v88 uint = (v3 & (v3 + v3))
                                                            _ = v88
                                                        default:
                                                            fmt.Println(v6, v3, fi11)
                                                        }
                                                    }
                                                    switch (^(fi11)) {
                                                    case 9223372036854775807:
                                                        v89 := GenS4{F0: 42, F1: -128}
                                                        _ = v89
                                                    case 127:
                                                        if v87 {
                                                            if ((v3 ^ v3) > (^(v3))) {
                                                                v90 := GenS6{F0: 42, F1: -1, F2: 2147483647}
                                                                if (255 >= 128) {
                                                                    switch 32767 {
                                                                    case -1:
                                                                        fmt.Println(fi11, v3)
                                                                    case 32768:
                                                                        if len(v2) > 0 {
                                                                            v2[0] = "foo bar"
                                                                        }
                                                                    }
                                                                    _ = len(v5)
                                                                    if v87 {
                                                                        switch fi11 {
                                                                        case 32767:
                                                                            fmt.Println(v7, v7, v87)
                                                                        case 32768:
                                                                            var v91 uint16 = 1
                                                                            _ = v91
                                                                        case 127:
                                                                            switch 127 {
                                                                            case 2147483647:
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    if v87 {
                                                                                        if (1e10 <= v6) {
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            if v87 {
                                                                                                var v92 int16 = -416
                                                                                                if (3.14159 != v6) {
                                                                                                    if v87 {
                                                                                                        var v93 uint32 = 4294967295
                                                                                                        var v94 float32 = 0.0
                                                                                                        _ = v93
                                                                                                        _ = v94
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            var v95 string = (("" + "") + "foo bar")
                                                                                                            var v96 uint64 = 1
                                                                                                            if len(v2) >= 2 {
                                                                                                                v2 = v2[1:len(v2)]
                                                                                                            }
                                                                                                            _ = v95
                                                                                                            _ = v96
                                                                                                        }
                                                                                                        fmt.Println(v6, v92)
                                                                                                    }
                                                                                                    fmt.Println(v3)
                                                                                                } else {
                                                                                                    var v97 bool = v87
                                                                                                    var v98 uint64 = 4294967295
                                                                                                    _ = v97
                                                                                                    _ = v98
                                                                                                }
                                                                                                _ = v92
                                                                                            }
                                                                                        } else {
                                                                                            v99 := GenS3{F0: 469, F1: -890}
                                                                                            if ((210 > 735) && v87) {
                                                                                                var v100 int8 = -1
                                                                                                v101 := GenS2{F0: 613, F1: 765, F2: 1}
                                                                                                v1["hello"] = "hello"
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                            } else {
                                                                                                var v102 float64 = v6
                                                                                                switch fi10 {
                                                                                                case -2147483648:
                                                                                                    switch fi10 {
                                                                                                    case 2147483647:
                                                                                                        v7 += ((((3.14159 + (v7 * 1.5)) / (0.1 / (-2.25 * v7))) / -2.25) / 0.1)
                                                                                                    }
                                                                                                default:
                                                                                                    if v87 {
                                                                                                        v7 = 3.14159
                                                                                                    }
                                                                                                }
                                                                                                _ = v102
                                                                                            }
                                                                                            v103 := GenS6{F0: 122, F1: 484, F2: 42}
                                                                                            _ = v103
                                                                                            _ = v99
                                                                                        }
                                                                                        fmt.Println(v6, fi11)
                                                                                        switch (fi11 ^ -1) {
                                                                                        case -1:
                                                                                            var v104 string = "hello"
                                                                                            _ = v104
                                                                                        case 32767:
                                                                                            var v105 uint8 = 1
                                                                                            _ = v105
                                                                                        case -264:
                                                                                            v7 -= v7
                                                                                        default:
                                                                                            v1["hello"] = "\n\t"
                                                                                        }
                                                                                    } else {
                                                                                        if ("hello" >= "\n\t") {
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            v87 = (241 <= 255)
                                                                                            fmt.Println(fi10)
                                                                                            var v106 bool = v87
                                                                                            var v107 int64 = 42
                                                                                            _ = v106
                                                                                            _ = v107
                                                                                        } else {
                                                                                            v108 := v1[(("hello" + ("\n\t" + "")) + "a")]
                                                                                            delete(v4, 42)
                                                                                            var v109 int16 = 42
                                                                                            switch ("" + v108) {
                                                                                            case "hello":
                                                                                                fmt.Println(v7, v109)
                                                                                                var v110 int16 = v109
                                                                                                var v111 float64 = v6
                                                                                                _ = v110
                                                                                                _ = v111
                                                                                            case "":
                                                                                                v87 = ((fi10 - (^(fi10))) == fi11)
                                                                                                var v112 int8 = -128
                                                                                                v113 := GenS7{F0: (v108 + v108), F1: (v112 & v112), F2: 765}
                                                                                                _ = v112
                                                                                                _ = v113
                                                                                            default:
                                                                                                var v114 int64 = 430
                                                                                                _ = v114
                                                                                            }
                                                                                            _ = v108
                                                                                            _ = v109
                                                                                        }
                                                                                        var v115 int16 = 42
                                                                                        if ((((3.14159 / 1.0) / 1.5) * 0.1) <= (v7 * ((0.0 - (v7 + (0.1 - (v7 * (v7 / (v7 / (v7 - 1e10))))))) - 1e10))) {
                                                                                            switch (("" + "a") + ("hello" + ("hello" + ""))) {
                                                                                            case "foo bar":
                                                                                                var v116 uint64 = 0
                                                                                                _ = v116
                                                                                            case "a":
                                                                                                var v117 int32 = 42
                                                                                                _ = v117
                                                                                            case "hello":
                                                                                                switch ("" + ("hello" + "hello")) {
                                                                                                case "":
                                                                                                    switch (fi11 % (((fi11 & fi10)) | 1)) {
                                                                                                    case 2147483647:
                                                                                                        var v118 uint64 = 2
                                                                                                        _ = v118
                                                                                                    case 128:
                                                                                                        if v87 {
                                                                                                            v7 = v7
                                                                                                            v3 -= (^(v3))
                                                                                                        } else {
                                                                                                            v119 := GenS3{F0: 1, F1: -472}
                                                                                                            var v120 bool = v87
                                                                                                            delete(v1, ("" + "\n\t"))
                                                                                                            v3 = 65535
                                                                                                            _ = v119
                                                                                                            _ = v120
                                                                                                        }
                                                                                                    case 555:
                                                                                                        switch "foo bar" {
                                                                                                        case "a":
                                                                                                            v121 := GenS2{F0: (v3 ^ (v3 << ((v3) & 63))), F1: -549, F2: 149}
                                                                                                            _ = v121
                                                                                                        case "\n\t":
                                                                                                            if true {
                                                                                                                var v122 uint64 = 2
                                                                                                                _ = v122
                                                                                                            } else {
                                                                                                                var v123 uint16 = 510
                                                                                                                var v124 string = "foo bar"
                                                                                                                v125 := GenS6{F0: 127, F1: -950, F2: 2147483648}
                                                                                                                v124 = v124
                                                                                                                if false {
                                                                                                                    var v126 uint32 = 42
                                                                                                                    _ = v126
                                                                                                                } else {
                                                                                                                    var v127 uint16 = 65535
                                                                                                                    v7 -= float32((float64((v7 * v7)) * (v6 * (v6 + v6))))
                                                                                                                    v123 -= (v123 | v127)
                                                                                                                    _ = v127
                                                                                                                }
                                                                                                                _ = v123
                                                                                                                _ = v124
                                                                                                                _ = v125
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        if v87 {
                                                                                                            var v128 uint32 = 1
                                                                                                            _ = v128
                                                                                                        }
                                                                                                    }
                                                                                                    v6 = v6
                                                                                                case "a":
                                                                                                    delete(v5, 42)
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v7)
                                                                                                default:
                                                                                                    v129 := GenS10{F0: 0, F1: 42}
                                                                                                    _ = v129
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            var v130 string = "foo bar"
                                                                                            v131 := GenS11{F0: v87, F1: 42}
                                                                                            _ = v130
                                                                                            _ = v131
                                                                                        }
                                                                                        _ = v115
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(fi11, fi11, fi10)
                                                                                }
                                                                            default:
                                                                                v132 := GenS12{F0: 255, F1: 42, F2: (v87 || v87)}
                                                                                _ = v132
                                                                            }
                                                                        }
                                                                        var v133 uint64 = 42
                                                                        v134 := GenS6{F0: 127, F1: 1, F2: 2}
                                                                        _ = v133
                                                                        _ = v134
                                                                    }
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                                var v135 uint16 = 42
                                                                v135 = (^(v135))
                                                                var v136 uint = v3
                                                                v136 = v3
                                                                fmt.Println(v7, fi11)
                                                                var v137 float64 = v6
                                                                v138 := v4[65535]
                                                                _ = v135
                                                                _ = v136
                                                                _ = v137
                                                                _ = v138
                                                                _ = v90
                                                            } else {
                                                                if (("hello" + "foo bar") < "a") {
                                                                    v139 := GenS13{F0: 2}
                                                                    v6 += 0.1
                                                                    fmt.Println(v7)
                                                                    var v140 int32 = 204
                                                                    _ = v139
                                                                    _ = v140
                                                                }
                                                            }
                                                            v7 = float32(v6)
                                                            fmt.Println(fi10, v6, v3)
                                                        }
                                                    default:
                                                        var v141 uint8 = 2
                                                        _ = v141
                                                    }
                                                } else {
                                                    var v142 float32 = (((v7 + (float32(((0.0 - (v6 * v6)) - v6)) - (v7 + -2.25))) - 3.14159) * float32(v6))
                                                    delete(v4, 1)
                                                    var v143 bool = (!(!(687 <= 32767)))
                                                    v7 -= 1.5
                                                    _ = v142
                                                    _ = v143
                                                }
                                                if (fi11 < fi11) {
                                                    v144 := GenS14{F0: "hello", F1: -1}
                                                    var v145 uint = (v3 | (v3 & v3))
                                                    _ = v144
                                                    _ = v145
                                                }
                                                var v146 uint64 = 0
                                                _ = v146
                                            }
                                            _ = v87
                                        }
                                    }
                                }
                            } else {
                                _ = len(v4)
                            }
                        }
                        _ = fi11
                    }
                    fmt.Println(fi10, v6)
                }
                var v147 uint64 = 945
                _ = fi10
                _ = v147
            }
            fmt.Println(v3, v6, v6)
        }
        _ = v7
    } else {
        var v148 uint64 = 910
        switch "\n\t" {
        case "foo bar":
            switch 127 {
            case 128:
                var v149 int32 = -588
                _ = v149
            case 0:
                var v150 int16 = -942
                _ = v150
            case -21:
                v3 = (v3 ^ 368)
            default:
                var v151 int64 = -487
                _ = v151
            }
        default:
            _ = len(v1)
        }
        _ = v148
    }
    switch 2147483647 {
    case 32768:
        fmt.Println(v6, v3)
    case 32767:
        fmt.Println(v3, v6)
    default:
        if true {
            for fi152 := 0; fi152 < 14; fi152++ {
                for fi153 := 0; fi153 < 19; fi153++ {
                    var v154 int8 = -1
                    var v155 float64 = 0.0
                    _ = fi153
                    _ = v154
                    _ = v155
                }
                if (v6 != v6) {
                    v6 += (((v6 * (v6 + (float64(1e10) - 0.1))) * 0.1) / (1.5 / 1.0))
                    if ((2 != 1) || false) {
                        v3 = v3
                        var v156 int16 = 1000
                        v3 = (^(v3))
                        _ = v156
                    }
                }
                _ = fi152
            }
            v6 -= v6
            v1[""] = "foo bar"
        }
    }
    var v157 string = "hello"
    _ = len(v2)
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v158 float32 = (((float32((v6 / ((float64(((3.14159 - 1e10) / ((1e10 * 0.0) - 3.14159))) + v6) / float64((float32(v6) / -2.25))))) / (1e10 * 1.5)) + -2.25) + 0.0)
        _ = v158
    case -124:
        _ = len(v2)
    case 32768:
        v2 = append(v2, v157)
    }
    for fi159 := 0; fi159 < 12; fi159++ {
        v6 += v6
        _ = fi159
    }
    v6 -= -2.25
    var v160 float64 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v157:", v157)
    fmt.Println("v160:", v160)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v157
    _ = v160
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
