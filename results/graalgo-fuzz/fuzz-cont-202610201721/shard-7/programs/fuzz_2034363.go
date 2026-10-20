package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS9 struct { F0 int64; F1 uint64 }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int16; F1 string }

func main() {
    v1 := []string{""}
    v2 := make(map[int]bool)
    var v3 string = "a"
    v4 := make(map[int]int)
    var v5 int32 = 0
    fmt.Println(v3, v3)
    switch 32767 {
    case 128:
        v5 = 1
    case -1:
        v5 += v5
    default:
        switch -2147483648 {
        case 4294967295:
            fmt.Println(v5, v5)
        case 1:
            var v6 uint64 = 184
            _ = v6
        }
    }
    v5 -= (v5 - (v5 - (v5 << ((42) & 63))))
    if false {
        v5 = (-(v5))
        for fi7 := 0; fi7 < 17; fi7++ {
            v5 -= (-(v5))
            if false {
                if true {
                    fmt.Println(v5, v5)
                    v8 := GenS1{F0: (fi7 | fi7), F1: true}
                    _ = v8
                } else {
                    var v9 uint16 = 1
                    if (true || (v3 < v3)) {
                        var v10 float64 = ((0.1 + -2.25) * (0.0 - 1.5))
                        _ = v10
                    } else {
                        var v11 bool = false
                        for fi12 := 0; fi12 < 6; fi12++ {
                            if v11 {
                                if (!(fi7 <= fi12)) {
                                    delete(v2, fi7)
                                }
                            } else {
                                v13 := GenS2{F0: fi7, F1: 1}
                                if (0 <= 42) {
                                    switch 0 {
                                    case -934:
                                        v3 = (v3 + v3)
                                    case 32767:
                                        v14 := GenS3{F0: 212}
                                        _ = v14
                                    default:
                                        if (!true) {
                                            var v15 uint16 = (v9 / (((^(v9))) | 1))
                                            fmt.Println(v11)
                                            v16 := GenS1{F0: (fi7 >> ((4294967295) & 63)), F1: (!(v11 || v11))}
                                            var v17 float64 = ((-2.25 - 0.1) - 1.5)
                                            if true {
                                                v15 += v15
                                            }
                                            var v18 int16 = 1000
                                            if len(v1) > 0 {
                                                v1[0] = ("\n\t" + ((((v3 + "foo bar") + "a") + ((("" + (v3 + "a")) + v3) + "foo bar")) + v3))
                                            }
                                            _ = v15
                                            _ = v16
                                            _ = v17
                                            _ = v18
                                        }
                                    }
                                    v19 := GenS1{F0: 32767, F1: (float64(1e10) < float64(1.5))}
                                    v20 := v4[1]
                                    _ = v19
                                    _ = v20
                                }
                                var v21 uint32 = 1
                                _ = v13
                                _ = v21
                            }
                            _ = fi12
                        }
                        var v22 uint8 = 128
                        _ = v11
                        _ = v22
                    }
                    _ = v9
                }
                v23 := GenS1{F0: 32767, F1: ((fi7 - fi7) < 4294967295)}
                _ = v23
            } else {
                if (false && (18446744073709551615 <= 42)) {
                    v5 = v5
                    var v24 string = v3
                    fmt.Println(v5)
                    _ = v24
                } else {
                    v25 := GenS4{F0: 18446744073709551615, F1: 0}
                    switch ((v3 + "a") + "a") {
                    case "foo bar":
                        _ = len(v1)
                    default:
                        if false {
                            v26 := GenS2{F0: 127, F1: 1000}
                            switch fi7 {
                            case 9223372036854775807:
                                for fi27 := 0; fi27 < 5; fi27++ {
                                    if (!true) {
                                        switch "a" {
                                        case "":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case "a":
                                            switch v3 {
                                            case "hello":
                                                if (32767 >= -1) {
                                                    switch (fi7 % (((^(fi7))) | 1)) {
                                                    case 4294967295:
                                                        v5 -= v5
                                                    case 2147483647:
                                                        delete(v2, fi27)
                                                    default:
                                                        if false {
                                                            var v28 uint = 375
                                                            v4[(-(fi7))] = -488
                                                            if (v5 >= v5) {
                                                                v29 := GenS3{F0: 19}
                                                                _ = v29
                                                            }
                                                            var v30 uint8 = 92
                                                            _ = v28
                                                            _ = v30
                                                        }
                                                    }
                                                    v5 = v5
                                                } else {
                                                    fmt.Println(fi27, v5, v3)
                                                    switch "\n\t" {
                                                    case "":
                                                        fmt.Println(fi7)
                                                    default:
                                                        fmt.Println(v3, v5)
                                                    }
                                                    switch v3 {
                                                    case "\n\t":
                                                        v5 = 298
                                                    }
                                                }
                                            }
                                            var v31 int32 = 0
                                            fmt.Println(v3, v5, v3)
                                            _ = v31
                                        case "\n\t":
                                            v32 := GenS1{F0: (-(fi7)), F1: ((42 != 0) || (1 != 127))}
                                            _ = v32
                                        default:
                                            var v33 float32 = 1e10
                                            fmt.Println(v33)
                                            _ = v33
                                        }
                                        if true {
                                            if false {
                                                if ((false || false) || (!true)) {
                                                    if (!false) {
                                                        v3 = v3
                                                        if (v3 >= v3) {
                                                            var v34 int8 = 23
                                                            if true {
                                                                var v35 uint16 = 0
                                                                if false {
                                                                    switch (fi27 + fi7) {
                                                                    case 32767:
                                                                        if false {
                                                                            var v36 int = fi27
                                                                            _ = v36
                                                                        } else {
                                                                            v5 = v5
                                                                        }
                                                                    case 32768:
                                                                        var v37 int = fi7
                                                                        _ = v37
                                                                    }
                                                                    v38 := GenS5{F0: -19, F1: (v3 + "a")}
                                                                    var v39 string = v3
                                                                    v3 = v3
                                                                    _ = len(v2)
                                                                    v40 := GenS6{F0: v39, F1: v5}
                                                                    _ = v38
                                                                    _ = v39
                                                                    _ = v40
                                                                } else {
                                                                    v41 := GenS3{F0: 127}
                                                                    _ = v41
                                                                }
                                                                _ = v35
                                                            }
                                                            fmt.Println(v3, v3)
                                                            _ = v34
                                                        } else {
                                                            v42 := GenS1{F0: (fi27 * (^(fi7))), F1: (fi7 < 1)}
                                                            var v43 uint32 = 65535
                                                            v43 -= v43
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        v44 := GenS6{F0: v3, F1: 724}
                                                        _ = v44
                                                    }
                                                    v3 = (v3 + (("\n\t" + v3) + v3))
                                                    v3 = v3
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v3 = (v3 + v3)
                                                } else {
                                                    switch (^(fi27)) {
                                                    case 2147483647:
                                                        v4[-107] = (fi27 + (fi7 + -9223372036854775808))
                                                    case 127:
                                                        v45 := GenS2{F0: 42, F1: -310}
                                                        _ = v45
                                                    case -1:
                                                        fmt.Println(v3)
                                                    default:
                                                        fmt.Println(fi7, v5, v5)
                                                    }
                                                    var v46 int16 = -1
                                                    _ = v46
                                                }
                                            }
                                            if false {
                                                _ = len(v1)
                                            }
                                            delete(v2, 955)
                                        }
                                    } else {
                                        v47 := GenS7{F0: 255}
                                        if (1 <= -128) {
                                            if true {
                                                var v48 int16 = -1
                                                var v49 int16 = (^(v48))
                                                if (15 < 49) {
                                                    var v50 int64 = 1
                                                    fmt.Println(fi27)
                                                    fmt.Println(v50, v5)
                                                    v5 = (v5 & v5)
                                                    _ = v50
                                                } else {
                                                    v48 = v48
                                                    switch (fi7 & (^(fi27))) {
                                                    case 32768:
                                                        v3 = v3
                                                    case 178:
                                                        v51 := v4[fi27]
                                                        _ = v51
                                                    case -2147483648:
                                                        var v52 int8 = 0
                                                        _ = v52
                                                    }
                                                    if true {
                                                        v48 = v49
                                                    }
                                                }
                                                _ = v48
                                                _ = v49
                                            } else {
                                                v3 = (v3 + "foo bar")
                                                var v53 float32 = ((0.1 + 0.1) + 1.5)
                                                _ = v53
                                            }
                                        }
                                        _ = v47
                                    }
                                    var v54 uint32 = 594
                                    if (!true) {
                                        v54 = v54
                                        v5 -= (v5 % ((42) | 1))
                                        if (731 != 4294967295) {
                                            fmt.Println(fi27, fi27)
                                            var v55 bool = false
                                            _ = v55
                                        }
                                    } else {
                                        switch fi7 {
                                        case -2147483648:
                                            if false {
                                                switch (v3 + v3) {
                                                case "foo bar":
                                                    var v56 float64 = 1e10
                                                    _ = v56
                                                }
                                                var v57 int8 = 100
                                                v3 = (v3 + (("a" + v3) + v3))
                                                _ = v57
                                            }
                                        }
                                    }
                                    _ = fi27
                                    _ = v54
                                }
                            default:
                                v58 := GenS8{F0: 0, F1: true, F2: 0}
                                _ = v58
                            }
                            v5 = -346
                            _ = v26
                        } else {
                            if (2 >= 42) {
                                v3 = v3
                                switch 128 {
                                case -9223372036854775808:
                                    for fi59 := 0; fi59 < 20; fi59++ {
                                        v3 = (v3 + (v3 + v3))
                                        v4[0] = fi7
                                        _ = fi59
                                    }
                                case 2147483647:
                                    v60 := GenS4{F0: 1, F1: 33}
                                    _ = v60
                                default:
                                    var v61 int16 = -32768
                                    _ = v61
                                }
                            }
                        }
                    }
                    fmt.Println(fi7, v5, fi7)
                    _ = v25
                }
                for fi62 := 0; fi62 < 12; fi62++ {
                    fmt.Println(fi62)
                    switch (fi7 % (((fi7 / (((-(fi62))) | 1))) | 1)) {
                    case -2147483648:
                        if (!(!(-2147483648 == (v5 / ((282) | 1))))) {
                            fmt.Println(v5)
                        } else {
                            switch (v3 + "a") {
                            case "hello":
                                var v63 uint32 = 42
                                _ = v63
                            default:
                                var v64 float64 = 1e10
                                _ = v64
                            }
                            var v65 uint16 = 886
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v65
                        }
                    default:
                        fmt.Println(fi62, fi7, v3)
                    }
                    _ = len(v4)
                    switch "hello" {
                    case "":
                        switch fi7 {
                        case 4294967295:
                            v3 = (v3 + (v3 + v3))
                        }
                    case "hello":
                        if (206 > 65535) {
                            v1 = append(v1, (v3 + ("foo bar" + v3)))
                            v66 := GenS5{F0: 42, F1: "a"}
                            v3 = v3
                            _ = v66
                        } else {
                            delete(v4, 481)
                            var v67 string = ("a" + v3)
                            v3 = ("" + ("foo bar" + ""))
                            _ = v67
                        }
                    case "a":
                        var v68 float32 = ((((-2.25 + ((1.5 - (3.14159 / 0.1)) - -2.25)) * 1.5) - 1.5) - 1.5)
                        _ = v68
                    default:
                        v3 = v3
                    }
                    _ = fi62
                }
                if false {
                    v69 := GenS9{F0: -9223372036854775808, F1: 930}
                    _ = v69
                } else {
                    var v70 bool = false
                    _ = v70
                }
            }
            var v71 uint16 = 65535
            _ = fi7
            _ = v71
        }
        if len(v1) > 0 {
            v1[0] = v3
        }
    }
    v72 := GenS1{F0: 219, F1: true}
    for fi73 := 0; fi73 < 11; fi73++ {
        v3 = ""
        var v74 string = ("\n\t" + "\n\t")
        _ = fi73
        _ = v74
    }
    switch 2147483647 {
    case 128:
        for fi75 := 0; fi75 < 20; fi75++ {
            var v76 int = -2147483648
            switch (v3 + v3) {
            case "a":
                fmt.Println(v3)
            case "hello":
                v77 := GenS10{F0: 32767, F1: 0}
                _ = v77
            }
            if (!false) {
                var v78 int32 = v5
                _ = v78
            }
            _ = fi75
            _ = v76
        }
    case -130:
        _ = len(v4)
    case 0:
        switch (v3 + (v3 + v3)) {
        case "\n\t":
            switch (("" + "foo bar") + (v3 + "\n\t")) {
            case "a":
                v5 = (v5 / ((-1) | 1))
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "":
                v79 := GenS9{F0: -362, F1: 18446744073709551615}
                _ = v79
            default:
                v3 = (v3 + "hello")
            }
        case "a":
            if false {
                if true {
                    var v80 bool = (v5 == v5)
                    var v81 uint = 373
                    fmt.Println(v3, v80)
                    _ = v80
                    _ = v81
                }
                fmt.Println(v3, v5)
                v3 = (("foo bar" + v3) + ((("" + "foo bar") + ("foo bar" + "foo bar")) + "foo bar"))
            }
        case "hello":
            v5 += v5
        default:
            if true {
                var v82 int16 = 233
                _ = v82
            } else {
                if (!(((-223 > 127) && (59 != 2)) && true)) {
                    switch ("a" + ("a" + v3)) {
                    case "a":
                        var v83 bool = (!true)
                        _ = v83
                    default:
                        var v84 int8 = 92
                        _ = v84
                    }
                }
            }
        }
    default:
        if len(v1) > 0 {
            v1[0] = ((("a" + ((v3 + (v3 + ((v3 + v3) + v3))) + v3)) + ((v3 + "hello") + v3)) + ("a" + ((v3 + (v3 + v3)) + v3)))
        }
    }
    if len(v1) > 0 {
        v1[0] = v3
    }
    fmt.Println(v3)
    for fi85 := 0; fi85 < 8; fi85++ {
        if false {
            switch "hello" {
            case "a":
                fmt.Println(v3)
            case "\n\t":
                v5 = 1
            case "":
                for fi86 := 0; fi86 < 14; fi86++ {
                    v3 = ""
                    var v87 int = 766
                    if false {
                        var v88 float32 = -2.25
                        _ = v88
                    } else {
                        var v89 int = fi85
                        _ = v89
                    }
                    if (42 > -32768) {
                        switch fi86 {
                        case -1:
                            v87 = fi85
                        case 0:
                            var v90 int64 = -9223372036854775808
                            _ = v90
                        default:
                            v87 += v87
                        }
                        delete(v4, v87)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if true {
                            v91 := GenS11{F0: (fi85 * (v87 / ((fi86) | 1)))}
                            _ = v91
                        } else {
                            fmt.Println(fi86)
                            var v92 int64 = -9223372036854775808
                            v87 = -502
                            _ = v92
                        }
                        var v93 int8 = -1
                        var v94 int = 128
                        _ = v93
                        _ = v94
                    } else {
                        v5 -= v5
                        if (false || false) {
                            v1 = append(v1, v3)
                            var v95 int = fi85
                            _ = len(v1)
                            v96 := GenS2{F0: fi86, F1: 186}
                            v97 := GenS3{F0: 0}
                            _ = v95
                            _ = v96
                            _ = v97
                        } else {
                            fmt.Println(fi85, fi86)
                            v4[fi86] = fi85
                        }
                        fmt.Println(fi86, v87, v5)
                    }
                    var v98 bool = (!(v5 == v5))
                    v99 := GenS7{F0: 629}
                    _ = fi86
                    _ = v87
                    _ = v98
                    _ = v99
                }
            default:
                var v100 uint8 = 2
                _ = v100
            }
            fmt.Println(fi85, fi85, v5)
            var v101 uint8 = 1
            _ = v101
        }
        _ = fi85
    }
    for fi102 := 0; fi102 < 6; fi102++ {
        var v103 int16 = 1000
        var v104 uint16 = 136
        if false {
            v105 := GenS11{F0: 128}
            _ = v105
        }
        _ = fi102
        _ = v103
        _ = v104
    }
    var v106 bool = (true && (!true))
    switch 127 {
    case 32767:
        var v107 uint8 = 2
        _ = v107
    case -371:
        v108 := GenS3{F0: 2}
        _ = v108
    default:
        _ = len(v1)
    }
    if false {
        if (v106 || (0 == 110)) {
            if (true && v106) {
                fmt.Println(v3, v5, v106)
                fmt.Println(v5, v5, v3)
            }
        } else {
            switch v3 {
            case "foo bar":
                var v109 uint = 157
                _ = v109
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                v5 = (v5 + (v5 | v5))
            }
            var v110 int64 = -1
            v106 = v106
            _ = v110
        }
        _ = len(v2)
    }
    v111 := GenS12{F0: -42, F1: (v3 + v3)}
    switch 127 {
    case 32767:
        fmt.Println(v3, v3, v5)
    default:
        var v112 uint = 2
        _ = v112
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v106:", v106)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v106
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
}
