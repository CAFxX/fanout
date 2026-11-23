package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 int8; F2 int8 }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 int64; F1 int16; F2 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint; F1 int64 }
func genFn0(p0 int8, p1 bool, p2 int32) uint64 {
    switch 42 {
    case 1:
        v4 := GenS1{F0: 2147483647}
        _ = v4
    case 0:
        p2 = p2
    case 9223372036854775807:
        p0 = (p0 % ((-115) | 1))
    }
    p0 += p0
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(fi5)
        v6 := []int{(-(fi5)), -2147483648, fi5, fi5}
        _ = fi5
        _ = v6
    }
    p2 = p2
    if (1 != 9223372036854775807) {
        v7 := []int{-76}
        var v8 string = "a"
        _ = v7
        _ = v8
    } else {
        switch 127 {
        case -9223372036854775808:
            var v9 int32 = p2
            _ = v9
        case 128:
            if (9223372036854775807 >= 128) {
                v10 := GenS1{F0: 0}
                v11 := []int{106, 32768}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = v10
                _ = v11
            }
        case -2147483648:
            var v12 int8 = (p0 - 2)
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 818
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 285
    v3 := []uint32{2147483648, 0, 4294967295, 1}
    v2 -= v2
    switch "hello" {
    case "":
        switch "a" {
        case "foo bar":
            fmt.Println(v2)
        }
    default:
        v13 := GenS2{F0: -57, F1: 42}
        _ = v13
    }
    if (2147483647 > 0) {
        v2 += (v2 << ((174) & 63))
        var v14 uint32 = 781
        _ = v14
    }
    v3 = append(v3, 1)
    for fi15 := 0; fi15 < 2; fi15++ {
        _ = len(v1)
        var v16 int = (fi15 * (fi15 / (((fi15 / ((32768) | 1))) | 1)))
        _ = fi15
        _ = v16
    }
    switch ((((("a" + "") + "") + "foo bar") + "hello") + "foo bar") {
    case "":
        v17 := GenS3{F0: 2, F1: 0}
        _ = v17
    case "foo bar":
        for fi18 := 0; fi18 < 1; fi18++ {
            v2 = (v2 >> ((0) & 63))
            v2 += v2
            if false {
                v2 += 748
                if (true || false) {
                    var v19 int = 1
                    fmt.Println(fi18, v19)
                    for fi20 := 0; fi20 < 9; fi20++ {
                        v1[(fi18 | fi18)] = (!((244 <= 2) || false))
                        delete(v1, fi20)
                        switch ("hello" + (("a" + ("a" + "foo bar")) + "")) {
                        case "":
                            if false {
                                var v21 float32 = 1e10
                                _ = v21
                            } else {
                                switch 0 {
                                case 9223372036854775807:
                                    var v22 uint32 = 787
                                    _ = v22
                                case 2147483647:
                                    v19 -= fi18
                                    var v23 uint = 2
                                    delete(v1, (v19 << ((65535) & 63)))
                                    _ = v23
                                default:
                                    _ = len(v1)
                                }
                                var v24 float64 = (1.5 / 1.5)
                                _ = v24
                            }
                        case "foo bar":
                            v25 := GenS1{F0: 0}
                            _ = v25
                        case "hello":
                            fmt.Println(fi20, fi18, fi20)
                        }
                        v26 := GenS2{F0: 42, F1: -467}
                        v19 = 0
                        _ = fi20
                        _ = v26
                    }
                    _ = v19
                }
            }
            _ = fi18
        }
    case "a":
        var v27 int64 = 42
        _ = v27
    }
    if ((2 > 1) || false) {
        switch -928 {
        case 1:
            for fi28 := 0; fi28 < 10; fi28++ {
                var v29 uint32 = 955
                _ = fi28
                _ = v29
            }
        default:
            for fi30 := 0; fi30 < 3; fi30++ {
                var v31 int64 = -1
                var v32 bool = (-1 == 385)
                fmt.Println(fi30, v2, v32)
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        v2 -= -1
    } else {
        var v33 float64 = float64(0.0)
        var v34 string = ("" + "hello")
        _ = v33
        _ = v34
    }
    fmt.Println(v2)
    v35 := GenS4{F0: ""}
    var v36 float32 = (3.14159 - (((1e10 / -2.25) + 0.1) * 0.0))
    for fi37 := 0; fi37 < 13; fi37++ {
        var v38 uint64 = genFn0(-128, true, 967)
        switch (fi37 + -904) {
        case 128:
            v36 -= (((float32(float64(((v36 + v36) / v36))) / v36) - v36) + v36)
        case 0:
            switch "\n\t" {
            case "":
                v38 -= (v38 >> ((349) & 63))
            default:
                var v39 uint8 = 128
                _ = v39
            }
        }
        v40 := v1[-654]
        _ = fi37
        _ = v38
        _ = v40
    }
    switch -717 {
    case 297:
        if (2 >= 4294967295) {
            if false {
                v2 -= v2
                switch 4294967295 {
                case 32768:
                    v36 += v36
                case -9223372036854775808:
                    var v41 uint8 = 0
                    _ = v41
                case 1:
                    for fi42 := 0; fi42 < 16; fi42++ {
                        v43 := GenS2{F0: 1, F1: 2147483647}
                        _ = len(v1)
                        _ = fi42
                        _ = v43
                    }
                default:
                    var v44 uint32 = 42
                    _ = v44
                }
                switch ("" + "hello") {
                case "\n\t":
                    var v45 float32 = float32(1.5)
                    _ = v45
                case "hello":
                    for fi46 := 0; fi46 < 12; fi46++ {
                        v2 = v2
                        fmt.Println(fi46, v2, fi46)
                        _ = fi46
                    }
                case "foo bar":
                    _ = len(v3)
                }
            }
        }
    }
    for fi47 := 0; fi47 < 5; fi47++ {
        if true {
            fmt.Println(fi47, v2, v2)
        }
        for fi48 := 0; fi48 < 5; fi48++ {
            var v49 float32 = (v36 + v36)
            var v50 uint8 = 191
            fmt.Println(v36)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi48
            _ = v49
            _ = v50
        }
        _ = fi47
    }
    v36 = 1.5
    if ((v36 - 1.5) <= v36) {
        if false {
            var v51 uint16 = 696
            fmt.Println(v51, v51, v51)
            if (true || true) {
                v51 -= (v51 | v51)
                if (85 == 110) {
                    for fi52 := 0; fi52 < 15; fi52++ {
                        if (false || true) {
                            if (!false) {
                                v53 := GenS5{F0: 18446744073709551615, F1: 127, F2: 55}
                                for fi54 := 0; fi54 < 6; fi54++ {
                                    fmt.Println(fi52, fi52)
                                    _ = fi54
                                }
                                _ = v53
                            } else {
                                switch fi52 {
                                case 42:
                                    var v55 uint = 688
                                    _ = v55
                                case -1:
                                    var v56 uint = 418
                                    _ = v56
                                case 128:
                                    var v57 int = -2147483648
                                    _ = v57
                                default:
                                    if false {
                                        fmt.Println(v51, v51)
                                    }
                                }
                                for fi58 := 0; fi58 < 2; fi58++ {
                                    v36 += ((v36 / 0.1) - v36)
                                    v36 = v36
                                    _ = fi58
                                }
                                switch ("hello" + "a") {
                                case "\n\t":
                                    fmt.Println(v51, v2, v51)
                                case "hello":
                                    switch "\n\t" {
                                    case "\n\t":
                                        if ((^(v51)) < v51) {
                                            for fi59 := 0; fi59 < 10; fi59++ {
                                                switch (fi59 - fi59) {
                                                case 530:
                                                    var v60 uint64 = 4294967295
                                                    _ = v60
                                                case 32768:
                                                    v36 += 0.0
                                                default:
                                                    v61 := GenS1{F0: 42}
                                                    _ = v61
                                                }
                                                switch "foo bar" {
                                                case "foo bar":
                                                    switch ("foo bar" + ((("hello" + (("a" + "a") + "a")) + "hello") + "\n\t")) {
                                                    case "a":
                                                        v51 = v51
                                                    case "":
                                                        fmt.Println(v51)
                                                    default:
                                                        switch ("foo bar" + "hello") {
                                                        case "hello":
                                                            v51 -= v51
                                                        case "a":
                                                            v2 += v2
                                                        default:
                                                            switch "a" {
                                                            case "a":
                                                                var v62 bool = (!true)
                                                                _ = v62
                                                            case "":
                                                                if false {
                                                                    if false {
                                                                        v2 = 1000
                                                                    } else {
                                                                        var v63 float32 = 0.0
                                                                        _ = v63
                                                                    }
                                                                }
                                                            default:
                                                                fmt.Println(v51, fi59)
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = fi59
                                            }
                                            switch fi52 {
                                            case 32767:
                                                var v64 uint16 = v51
                                                _ = v64
                                            default:
                                                for fi65 := 0; fi65 < 6; fi65++ {
                                                    v2 += v2
                                                    if (("\n\t" + "\n\t") >= "") {
                                                        var v66 int64 = 9223372036854775807
                                                        var v67 uint64 = 658
                                                        if false {
                                                            v36 -= v36
                                                            v67 -= genFn0(0, (false && (v51 >= v51)), 42)
                                                            _ = len(v1)
                                                        } else {
                                                            if (true && ("\n\t" >= ("foo bar" + "hello"))) {
                                                                fmt.Println(fi52, fi65)
                                                                if false {
                                                                    var v68 int16 = (v2 % (((v2 << ((4294967295) & 63))) | 1))
                                                                    _ = v68
                                                                } else {
                                                                    var v69 float32 = v36
                                                                    switch (fi65 | 127) {
                                                                    case 128:
                                                                        fmt.Println(v69, v51, fi52)
                                                                    }
                                                                    _ = v69
                                                                }
                                                                if (!((42 < 128) && (v2 < -495))) {
                                                                    v36 = (3.14159 - v36)
                                                                }
                                                                if true {
                                                                    v70 := GenS4{F0: ""}
                                                                    var v71 int8 = 1
                                                                    switch ("hello" + ((("a" + ("" + ("\n\t" + ((("foo bar" + (("foo bar" + "a") + "foo bar")) + "hello") + (((("\n\t" + "foo bar") + ("" + "\n\t")) + "\n\t") + ((("" + "a") + "foo bar") + "a")))))) + "a") + ("hello" + "foo bar"))) {
                                                                    case "a":
                                                                        delete(v1, fi52)
                                                                    case "foo bar":
                                                                        v72 := GenS4{F0: "a"}
                                                                        _ = v72
                                                                    }
                                                                    _ = v70
                                                                    _ = v71
                                                                } else {
                                                                    var v73 string = "\n\t"
                                                                    v74 := GenS4{F0: v73}
                                                                    _ = v73
                                                                    _ = v74
                                                                }
                                                                var v75 uint32 = 0
                                                                _ = v75
                                                            }
                                                        }
                                                        if false {
                                                            v66 = (^(v66))
                                                        }
                                                        var v76 int = (fi52 | (fi52 << ((0) & 63)))
                                                        _ = v66
                                                        _ = v67
                                                        _ = v76
                                                    } else {
                                                        fmt.Println(fi65)
                                                        v77 := GenS3{F0: 4294967295, F1: -598}
                                                        _ = v77
                                                    }
                                                    switch 127 {
                                                    case 370:
                                                        if (858 != 42) {
                                                            var v78 uint32 = 614
                                                            var v79 int8 = 42
                                                            _ = v78
                                                            _ = v79
                                                        } else {
                                                            var v80 uint = 291
                                                            fmt.Println(v51, v51, fi52)
                                                            _ = v80
                                                        }
                                                    }
                                                    _ = fi65
                                                }
                                            }
                                            var v81 uint = 0
                                            _ = v81
                                        } else {
                                            var v82 uint16 = (^(v51))
                                            if true {
                                                v83 := GenS6{F0: 42, F1: 1000}
                                                var v84 float64 = -2.25
                                                var v85 int = (^(fi52))
                                                _ = v83
                                                _ = v84
                                                _ = v85
                                            }
                                            _ = v82
                                        }
                                    default:
                                        for fi86 := 0; fi86 < 16; fi86++ {
                                            var v87 int = fi52
                                            if true {
                                                if (!(42 <= 100)) {
                                                    v88 := GenS3{F0: 443, F1: 2147483647}
                                                    if (65535 < 293) {
                                                        v89 := GenS7{F0: -595, F1: 1000, F2: (255 > 1)}
                                                        _ = v89
                                                    } else {
                                                        if (true && true) {
                                                            if (true || false) {
                                                                if (((1e10 / (3.14159 / (v36 * float32(3.14159)))) + 1.5) <= (float32(-2.25) * (3.14159 + 1.5))) {
                                                                    var v90 uint8 = 255
                                                                    _ = v90
                                                                }
                                                                fmt.Println(v36, v51, v36)
                                                            } else {
                                                                var v91 int8 = -89
                                                                _ = len(v1)
                                                                _ = v91
                                                            }
                                                            v87 -= -9223372036854775808
                                                            var v92 string = "a"
                                                            fmt.Println(fi52, v51, v2)
                                                            _ = v92
                                                        } else {
                                                            switch ("hello" + ("foo bar" + ("foo bar" + "hello"))) {
                                                            case "a":
                                                                var v93 uint16 = v51
                                                                _ = v93
                                                            default:
                                                                var v94 int8 = 42
                                                                _ = v94
                                                            }
                                                            if false {
                                                                var v95 int64 = -503
                                                                _ = len(v1)
                                                                v95 += (v95 & 109)
                                                                var v96 uint64 = genFn0(100, false, 1)
                                                                if ((v51 * (v51 / ((v51) | 1))) > v51) {
                                                                    v36 += v36
                                                                }
                                                                _ = v95
                                                                _ = v96
                                                            }
                                                        }
                                                        v36 = v36
                                                        v87 = (v87 + (v87 ^ (^(fi86))))
                                                        v97 := GenS8{F0: "hello"}
                                                        var v98 uint32 = 2147483648
                                                        _ = v97
                                                        _ = v98
                                                    }
                                                    _ = v88
                                                } else {
                                                    v51 = 255
                                                    if true {
                                                        if false {
                                                            v99 := GenS7{F0: 0, F1: (v2 / ((1) | 1)), F2: (false || true)}
                                                            fmt.Println(v2, v36)
                                                            switch "hello" {
                                                            case "":
                                                                if len(v3) > 0 {
                                                                    v3[0] = 42
                                                                }
                                                            case "foo bar":
                                                                if true {
                                                                    var v100 uint64 = 18446744073709551615
                                                                    var v101 int16 = v2
                                                                    _ = v100
                                                                    _ = v101
                                                                } else {
                                                                    if true {
                                                                        fmt.Println(v36, fi86)
                                                                        v102 := GenS8{F0: ("\n\t" + "hello")}
                                                                        var v103 float64 = 0.1
                                                                        v36 -= 3.14159
                                                                        _ = v102
                                                                        _ = v103
                                                                    }
                                                                    var v104 uint64 = 443
                                                                    _ = v104
                                                                }
                                                            }
                                                            _ = v99
                                                        }
                                                        fmt.Println(v2)
                                                        v105 := v1[fi86]
                                                        _ = v105
                                                    }
                                                    fmt.Println(v36, fi86, v2)
                                                    var v106 float64 = 1.5
                                                    _ = v106
                                                }
                                                fmt.Println(v51, v2, v36)
                                            } else {
                                                if ((!false) || false) {
                                                    switch ("\n\t" + "\n\t") {
                                                    case "a":
                                                        v107 := GenS9{F0: -983}
                                                        _ = v107
                                                    case "foo bar":
                                                        v51 = 65535
                                                        fmt.Println(v51)
                                                    case "\n\t":
                                                        v108 := GenS9{F0: -922}
                                                        _ = v108
                                                    default:
                                                        var v109 uint64 = genFn0(0, false, 309)
                                                        v2 = v2
                                                        v51 = (v51 - (^(v51)))
                                                        _ = v109
                                                    }
                                                    var v110 string = ""
                                                    if (2147483647 == 167) {
                                                        fmt.Println(fi86)
                                                        fmt.Println(v2, fi52)
                                                        v51 -= v51
                                                        fmt.Println(v2, fi86)
                                                        switch fi86 {
                                                        case -775:
                                                            var v111 float64 = 3.14159
                                                            _ = v111
                                                        default:
                                                            v112 := v1[(v87 % (((^(v87))) | 1))]
                                                            _ = v112
                                                        }
                                                        v51 = 355
                                                        var v113 int16 = v2
                                                        v36 = (v36 - (v36 - ((float32(0.0) + ((v36 * ((1e10 / (float32(3.14159) - float32(float64((v36 * v36))))) - 1e10)) * v36)) - v36)))
                                                        _ = v113
                                                    }
                                                    _ = v110
                                                } else {
                                                    v2 = v2
                                                    var v114 uint8 = 42
                                                    var v115 int16 = 0
                                                    _ = v114
                                                    _ = v115
                                                }
                                                v87 += v87
                                            }
                                            _ = fi86
                                            _ = v87
                                        }
                                    }
                                case "":
                                    fmt.Println(fi52, fi52)
                                default:
                                    var v116 uint = 464
                                    _ = v116
                                }
                            }
                            v117 := GenS10{F0: 65535, F1: -9223372036854775808}
                            _ = v117
                        }
                        if true {
                            v36 = v36
                            if ((!false) || true) {
                                fmt.Println(v51, v2, fi52)
                                for fi118 := 0; fi118 < 12; fi118++ {
                                    v2 -= 42
                                    var v119 float64 = 3.14159
                                    _ = fi118
                                    _ = v119
                                }
                            }
                            delete(v1, 2147483647)
                        } else {
                            for fi120 := 0; fi120 < 12; fi120++ {
                                fmt.Println(v2)
                                _ = fi120
                            }
                            v2 -= 0
                            var v121 int64 = 9223372036854775807
                            _ = v121
                        }
                        v36 += ((v36 - (v36 * ((1e10 * (float32((-2.25 / ((float64((1.5 - v36)) * 3.14159) + 1.5))) + v36)) * v36))) - v36)
                        _ = fi52
                    }
                }
                delete(v1, 2147483647)
            }
            _ = v51
        }
    }
    v36 += ((float32(0.1) - ((v36 + v36) - 0.0)) * -2.25)
    for fi122 := 0; fi122 < 11; fi122++ {
        var v123 int64 = 9223372036854775807
        v124 := GenS8{F0: (("a" + "") + ((("foo bar" + "\n\t") + "a") + (("hello" + "hello") + ("foo bar" + ""))))}
        _ = fi122
        _ = v123
        _ = v124
    }
    v125 := GenS7{F0: -9223372036854775808, F1: (v2 * (v2 - (-(v2)))), F2: ((100 == 42) && (4294967295 <= 65535))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v125
    _ = v2
    _ = v3
    _ = v35
    _ = v36
}
