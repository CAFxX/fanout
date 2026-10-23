package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 bool; F1 bool; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 int32; F1 bool; F2 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int64; F1 string }
type GenS11 struct { F0 int8; F1 uint32 }
type GenS12 struct { F0 uint8; F1 int32 }
type GenS13 struct { F0 uint16; F1 int32 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int64; F1 bool; F2 int16 }
func genFn0(p0 int) uint64 {
    v6 := []int{(p0 & p0), (p0 + (-(p0))), (p0 % ((-1) | 1)), p0}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        if (!(!((!false) || (128 < 0)))) {
            v8 := GenS1{F0: (fi7 - 791), F1: -73}
            for fi9 := 0; fi9 < 11; fi9++ {
                fmt.Println(p0, fi9, fi7)
                var v10 int = -9223372036854775808
                if len(v6) > 0 {
                    v6[0] = fi9
                }
                _ = fi9
                _ = v10
            }
            switch ("a" + "foo bar") {
            case "":
                var v11 int16 = 32767
                _ = v11
            default:
                var v12 uint64 = 42
                _ = v12
            }
            _ = v8
        }
        p0 = (p0 & fi7)
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 bool = false
            p0 += (fi13 + (fi7 % ((fi7) | 1)))
            v14 = ((!(!("hello" < (("" + "\n\t") + ("a" + "hello"))))) && v14)
            _ = fi13
            _ = v14
        }
        _ = fi7
    }
    if len(v6) > 0 {
        v6[0] = 1
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        if true {
            if true {
                for fi16 := 0; fi16 < 19; fi16++ {
                    var v17 uint32 = 42
                    v17 += (^(v17))
                    _ = fi16
                    _ = v17
                }
                var v18 uint16 = 65535
                if ((-995 <= -119) && false) {
                    var v19 string = "\n\t"
                    var v20 float64 = 0.1
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(p0)
                    if true {
                        v21 := GenS2{F0: ((1 < 65535) || ("a" != "")), F1: true, F2: p0}
                        var v22 int64 = 9223372036854775807
                        var v23 float32 = 0.1
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        switch (("" + "hello") + ("hello" + "")) {
                        case "":
                            v24 := GenS3{F0: -1}
                            _ = v24
                        }
                    }
                    v18 -= (^(v18))
                }
                _ = v18
            } else {
                fmt.Println(p0, p0)
            }
            if false {
                fmt.Println(fi15)
                switch "foo bar" {
                case "foo bar":
                    p0 += -468
                }
                var v25 bool = false
                _ = v25
            } else {
                v26 := GenS2{F0: true, F1: true, F2: (fi15 & fi15)}
                _ = v26
            }
        } else {
            switch ("" + "\n\t") {
            case "hello":
                switch 32768 {
                case 2147483647:
                    var v27 uint8 = 0
                    _ = v27
                case 42:
                    var v28 uint = 94
                    _ = v28
                }
            }
            for fi29 := 0; fi29 < 1; fi29++ {
                var v30 uint = 18446744073709551615
                v30 = v30
                _ = fi29
                _ = v30
            }
            if false {
                v31 := GenS2{F0: true, F1: (42 <= -4), F2: fi15}
                fmt.Println(fi15, p0)
                _ = v31
            }
        }
        _ = fi15
    }
    _ = p0
    _ = v6
    return 772
}
func genFn1(p0 int64) int8 {
    v32 := []int{}
    var v33 int16 = 1000
    v34 := GenS3{F0: (v33 - v33)}
    fmt.Println(v33)
    _ = p0
    _ = v32
    _ = v33
    _ = v34
    return 127
}

func main() {
    var v1 float32 = (((-2.25 - (0.1 / ((1e10 + float32(0.1)) / (0.1 + (0.1 * 0.1))))) - (1.5 - 1e10)) - -2.25)
    var v2 uint = 1
    var v3 uint64 = 263
    v4 := make(map[int]uint32)
    v5 := make(map[int]bool)
    if (!true) {
        v35 := GenS1{F0: -399, F1: 127}
        for fi36 := 0; fi36 < 17; fi36++ {
            switch "\n\t" {
            case "hello":
                var v37 int32 = 2147483647
                _ = v37
            case "a":
                switch fi36 {
                case 128:
                    switch (-(fi36)) {
                    case -2147483648:
                        v3 += 4294967295
                    case 2147483647:
                        v2 += 622
                    default:
                        v1 -= (v1 + ((3.14159 * (v1 + float32(-2.25))) + 0.0))
                    }
                case 32768:
                    fmt.Println(fi36, fi36)
                case 266:
                    for fi38 := 0; fi38 < 10; fi38++ {
                        if true {
                            var v39 float32 = v1
                            _ = v39
                        } else {
                            _ = len(v5)
                            if (v2 <= (v2 << ((553) & 63))) {
                                v40 := []int{(fi38 & (fi38 / (((-(fi36))) | 1))), (fi36 << (((^(v2))) & 63)), (fi36 << ((42) & 63)), fi36}
                                switch 2147483647 {
                                case 1:
                                    if ((!false) && false) {
                                        v5[127] = true
                                        v41 := GenS4{F0: 2147483647, F1: 927}
                                        if (2147483647 < 2147483647) {
                                            v42 := v4[fi38]
                                            _ = v42
                                        }
                                        v1 -= v1
                                        _ = v41
                                    }
                                }
                                v43 := GenS5{F0: -2147483648}
                                _ = v40
                                _ = v43
                            }
                            if ((float32(-2.25) >= (float32(1.5) * -2.25)) || true) {
                                v3 = v3
                                v3 += v3
                            } else {
                                var v44 uint = (v2 * v2)
                                _ = v44
                            }
                        }
                        var v45 int = fi36
                        if false {
                            if (-1 != -827) {
                                v46 := GenS4{F0: 2147483647, F1: 959}
                                switch (fi36 % ((v45) | 1)) {
                                case 127:
                                    switch ("foo bar" + "hello") {
                                    case "hello":
                                        if true {
                                            v45 = -1
                                        }
                                    case "foo bar":
                                        if true {
                                            v47 := []int{-412, (v45 - 32767), (fi38 | fi38), fi38}
                                            v3 = genFn0((fi38 & fi36))
                                            if len(v47) >= 2 {
                                                v47 = v47[1:len(v47)]
                                            }
                                            var v48 int16 = -653
                                            switch -207 {
                                            case -1:
                                                if false {
                                                    if (0 == 708) {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v49 := v5[-9223372036854775808]
                                                            _ = v49
                                                        case "foo bar":
                                                            v50 := GenS6{F0: (^(v3)), F1: 2}
                                                            _ = v50
                                                        case "a":
                                                            var v51 int16 = (v48 | (v48 * 32767))
                                                            _ = v51
                                                        }
                                                        var v52 uint64 = (v3 * v3)
                                                        switch 42 {
                                                        case -1:
                                                            v53 := GenS7{F0: 2147483647, F1: true, F2: ""}
                                                            _ = v53
                                                        default:
                                                            switch ("" + (("hello" + (("hello" + "\n\t") + ("\n\t" + "foo bar"))) + "hello")) {
                                                            case "\n\t":
                                                                if (!true) {
                                                                    v54 := GenS1{F0: 42, F1: 42}
                                                                    var v55 int32 = 1
                                                                    fmt.Println(v45)
                                                                    _ = v54
                                                                    _ = v55
                                                                } else {
                                                                    switch v45 {
                                                                    case 9223372036854775807:
                                                                        var v56 int8 = genFn1(943)
                                                                        _ = v56
                                                                    case 127:
                                                                        switch 32768 {
                                                                        case 32767:
                                                                            v52 = genFn0(fi36)
                                                                        case 1:
                                                                            switch (("\n\t" + ("\n\t" + "hello")) + "foo bar") {
                                                                            case "":
                                                                                fmt.Println(v48)
                                                                            case "foo bar":
                                                                                var v57 int64 = 0
                                                                                _ = v57
                                                                            case "hello":
                                                                                v58 := GenS6{F0: v3, F1: 255}
                                                                                _ = v58
                                                                            default:
                                                                                if false {
                                                                                    fmt.Println(v3, v45)
                                                                                } else {
                                                                                    v3 += genFn0((fi36 / ((-160) | 1)))
                                                                                    if true {
                                                                                        var v59 float32 = v1
                                                                                        v60 := GenS7{F0: -1, F1: true, F2: "hello"}
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                    }
                                                                                    fmt.Println(v52, v2)
                                                                                    v3 += (v3 - (^(v3)))
                                                                                }
                                                                            }
                                                                        default:
                                                                            fmt.Println(fi36)
                                                                        }
                                                                    case -194:
                                                                        switch (fi36 >> (((v2 * 151)) & 63)) {
                                                                        case 2147483647:
                                                                            if (!false) {
                                                                                switch fi36 {
                                                                                case 127:
                                                                                    if false {
                                                                                        var v61 string = ("\n\t" + "hello")
                                                                                        var v62 uint64 = genFn0(v45)
                                                                                        _ = v61
                                                                                        _ = v62
                                                                                    } else {
                                                                                        switch "a" {
                                                                                        case "":
                                                                                            _ = len(v4)
                                                                                        case "foo bar":
                                                                                            var v63 float64 = (0.1 + 3.14159)
                                                                                            _ = v63
                                                                                        case "\n\t":
                                                                                            var v64 string = ("\n\t" + "a")
                                                                                            _ = v64
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v65 float32 = (((1e10 * -2.25) - (float32(0.0) / v1)) - -2.25)
                                                                                var v66 bool = true
                                                                                _ = v65
                                                                                _ = v66
                                                                            } else {
                                                                                switch -353 {
                                                                                case 32768:
                                                                                    fmt.Println(v48, fi38, v48)
                                                                                    var v67 int16 = 1
                                                                                    var v68 uint32 = 65535
                                                                                    v69 := GenS4{F0: 546, F1: v68}
                                                                                    _ = v67
                                                                                    _ = v68
                                                                                    _ = v69
                                                                                }
                                                                                var v70 uint32 = 2
                                                                                var v71 uint64 = (v3 ^ uint64(v2))
                                                                                _ = v70
                                                                                _ = v71
                                                                            }
                                                                        case -1:
                                                                            v45 += 279
                                                                        case -563:
                                                                            v2 = v2
                                                                        default:
                                                                            v1 -= v1
                                                                        }
                                                                    default:
                                                                        v2 -= 18446744073709551615
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        _ = len(v5)
                                                                    case "a":
                                                                        fmt.Println(v52)
                                                                    case "\n\t":
                                                                        if (2 == 255) {
                                                                            v48 = (v48 - (v48 - v48))
                                                                        }
                                                                    default:
                                                                        v52 = v52
                                                                    }
                                                                    v2 = 42
                                                                }
                                                            case "":
                                                                var v72 uint = (v2 >> ((uint(v52)) & 63))
                                                                var v73 int64 = 9223372036854775807
                                                                switch "a" {
                                                                case "a":
                                                                    var v74 uint8 = 127
                                                                    _ = v74
                                                                case "foo bar":
                                                                    switch 127 {
                                                                    case -9223372036854775808:
                                                                        v2 = (^(v72))
                                                                    case 128:
                                                                        var v75 uint8 = 42
                                                                        _ = v75
                                                                    default:
                                                                        if (3.14159 > float64(((0.1 / (float32((0.0 - float64(((v1 + v1) + (v1 / (v1 * v1)))))) * (((v1 + (v1 - (v1 + (v1 / 3.14159)))) / 0.1) * float32(1e10)))) * v1))) {
                                                                            if false {
                                                                                switch (fi36 % ((fi36) | 1)) {
                                                                                case 32767:
                                                                                    switch -9223372036854775808 {
                                                                                    case 127:
                                                                                        if len(v47) >= 2 {
                                                                                            v47 = v47[1:len(v47)]
                                                                                        }
                                                                                    case 42:
                                                                                        if true {
                                                                                            if false {
                                                                                                switch "a" {
                                                                                                case "a":
                                                                                                    v76 := GenS8{F0: -128}
                                                                                                    _ = v76
                                                                                                }
                                                                                                switch ("\n\t" + "") {
                                                                                                case "":
                                                                                                    var v77 uint64 = (^(v3))
                                                                                                    _ = v77
                                                                                                case "foo bar":
                                                                                                    if false {
                                                                                                        switch (("a" + "") + "\n\t") {
                                                                                                        case "a":
                                                                                                            if false {
                                                                                                                v45 = (v45 - v45)
                                                                                                                var v78 uint32 = 65535
                                                                                                                var v79 string = "\n\t"
                                                                                                                _ = v78
                                                                                                                _ = v79
                                                                                                            } else {
                                                                                                                var v80 uint16 = 65535
                                                                                                                _ = len(v4)
                                                                                                                var v81 uint32 = 890
                                                                                                                _ = v80
                                                                                                                _ = v81
                                                                                                            }
                                                                                                            v47 = append(v47, v45)
                                                                                                            fmt.Println(v52, fi38, v48)
                                                                                                        case "\n\t":
                                                                                                            fmt.Println(v2, v48, v48)
                                                                                                        default:
                                                                                                            delete(v5, 127)
                                                                                                        }
                                                                                                        var v82 uint = 2
                                                                                                        v52 = (v3 + (^(v52)))
                                                                                                        _ = v82
                                                                                                    } else {
                                                                                                        fmt.Println(fi38, v72)
                                                                                                        if (!false) {
                                                                                                            if (false || true) {
                                                                                                                fmt.Println(v3)
                                                                                                            }
                                                                                                            v83 := GenS3{F0: v48}
                                                                                                            if len(v47) > 0 {
                                                                                                                v47[0] = -2147483648
                                                                                                            }
                                                                                                            _ = v83
                                                                                                        } else {
                                                                                                            v3 -= (v52 / ((1) | 1))
                                                                                                        }
                                                                                                        switch "" {
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(v2, v52, v45)
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    switch "hello" {
                                                                                                    case "":
                                                                                                        var v84 uint16 = 2
                                                                                                        _ = v84
                                                                                                    case "hello":
                                                                                                        _ = len(v47)
                                                                                                    case "foo bar":
                                                                                                        v85 := GenS9{F0: v2}
                                                                                                        if (float32(float64(float32(1.5))) >= v1) {
                                                                                                            if ((v72 < (v2 >> ((v2) & 63))) && true) {
                                                                                                                if true {
                                                                                                                    v86 := GenS4{F0: 2147483647, F1: 1}
                                                                                                                    v87 := GenS10{F0: v73, F1: "hello"}
                                                                                                                    var v88 uint = v72
                                                                                                                    _ = v86
                                                                                                                    _ = v87
                                                                                                                    _ = v88
                                                                                                                }
                                                                                                                if (105 == 100) {
                                                                                                                    var v89 int = -614
                                                                                                                    _ = v89
                                                                                                                } else {
                                                                                                                    fmt.Println(fi38, v72)
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v90 uint = (v72 * 42)
                                                                                                                v52 = v52
                                                                                                                if (!true) {
                                                                                                                    if ((^(v90)) <= v2) {
                                                                                                                        v48 = -32768
                                                                                                                        v2 = v72
                                                                                                                        v72 += v72
                                                                                                                    }
                                                                                                                    v1 -= 0.1
                                                                                                                } else {
                                                                                                                    v1 -= v1
                                                                                                                }
                                                                                                                _ = v90
                                                                                                            }
                                                                                                            var v91 string = ""
                                                                                                            _ = v91
                                                                                                        } else {
                                                                                                            var v92 float32 = ((0.1 / (3.14159 - 1.5)) / v1)
                                                                                                            _ = v92
                                                                                                        }
                                                                                                        _ = len(v4)
                                                                                                        _ = v85
                                                                                                    default:
                                                                                                        if false {
                                                                                                            if false {
                                                                                                                var v93 float64 = 0.1
                                                                                                                var v94 int16 = (v48 >> ((4294967295) & 63))
                                                                                                                if true {
                                                                                                                    if false {
                                                                                                                        if (v2 == (v2 & uint(v52))) {
                                                                                                                            v3 = v3
                                                                                                                            switch fi36 {
                                                                                                                            case 9223372036854775807:
                                                                                                                                if (0 < 42) {
                                                                                                                                    v48 = int16(v3)
                                                                                                                                    var v95 int64 = v73
                                                                                                                                    _ = v95
                                                                                                                                }
                                                                                                                            }
                                                                                                                            fmt.Println(v72, v1, v94)
                                                                                                                        }
                                                                                                                        var v96 uint = (v2 / ((v2) | 1))
                                                                                                                        switch (^(v45)) {
                                                                                                                        case -492:
                                                                                                                            fmt.Println(v93, v52)
                                                                                                                        default:
                                                                                                                            var v97 uint8 = 255
                                                                                                                            _ = v97
                                                                                                                        }
                                                                                                                        var v98 uint32 = 326
                                                                                                                        delete(v4, (v45 | fi36))
                                                                                                                        _ = v96
                                                                                                                        _ = v98
                                                                                                                    }
                                                                                                                    v1 = float32(((v93 / v93) / v93))
                                                                                                                } else {
                                                                                                                    var v99 uint16 = 2
                                                                                                                    switch ("foo bar" + "\n\t") {
                                                                                                                    case "a":
                                                                                                                        v100 := GenS9{F0: v72}
                                                                                                                        _ = v100
                                                                                                                    case "hello":
                                                                                                                        if false {
                                                                                                                            v101 := GenS1{F0: (fi36 >> ((864) & 63)), F1: -8}
                                                                                                                            var v102 uint16 = v99
                                                                                                                            delete(v5, 128)
                                                                                                                            _ = v101
                                                                                                                            _ = v102
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v99
                                                                                                                }
                                                                                                                var v103 int64 = (-(v73))
                                                                                                                _ = v103
                                                                                                                _ = v93
                                                                                                                _ = v94
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v104 := GenS7{F0: -778, F1: ("foo bar" >= ""), F2: "\n\t"}
                                                                                                v3 -= (v52 / (((^(v3))) | 1))
                                                                                                fmt.Println(v72)
                                                                                                v105 := GenS11{F0: 93, F1: 2}
                                                                                                var v106 int64 = -1
                                                                                                v52 += genFn0(fi36)
                                                                                                _ = v104
                                                                                                _ = v105
                                                                                                _ = v106
                                                                                            }
                                                                                            v48 += int16(v2)
                                                                                        } else {
                                                                                            fmt.Println(v72, fi38, v1)
                                                                                        }
                                                                                    default:
                                                                                        var v107 uint = 328
                                                                                        _ = v107
                                                                                    }
                                                                                }
                                                                                if ((v52 - (v3 >> ((v72) & 63))) == v52) {
                                                                                    var v108 int8 = 100
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v1)
                                                                                    case "\n\t":
                                                                                        var v109 string = "foo bar"
                                                                                        _ = v109
                                                                                    case "hello":
                                                                                        v47 = append(v47, v45)
                                                                                    }
                                                                                    _ = v108
                                                                                } else {
                                                                                    if (true && (523 >= 894)) {
                                                                                        var v110 bool = true
                                                                                        v3 = (v3 << (((v72 >> (((v2 % (((^(v2))) | 1))) & 63))) & 63))
                                                                                        var v111 uint64 = v3
                                                                                        _ = v110
                                                                                        _ = v111
                                                                                    } else {
                                                                                        fmt.Println(v73, fi38)
                                                                                        var v112 int8 = 1
                                                                                        _ = v112
                                                                                    }
                                                                                    if (!false) {
                                                                                        var v113 string = ("a" + "a")
                                                                                        _ = v113
                                                                                    } else {
                                                                                        v2 += v2
                                                                                    }
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v114 float32 = (v1 * v1)
                                                                            v72 = (v72 | 65535)
                                                                            _ = v114
                                                                        }
                                                                        _ = len(v4)
                                                                        v115 := GenS12{F0: 128, F1: int32(v52)}
                                                                        if false {
                                                                            if false {
                                                                                v116 := v5[-9223372036854775808]
                                                                                delete(v5, -1)
                                                                                _ = v116
                                                                            } else {
                                                                                var v117 uint = (v72 ^ v2)
                                                                                _ = v117
                                                                            }
                                                                            var v118 bool = true
                                                                            _ = v118
                                                                        }
                                                                        _ = v115
                                                                    }
                                                                case "\n\t":
                                                                    delete(v5, 758)
                                                                default:
                                                                    v48 -= (v48 << ((v72) & 63))
                                                                }
                                                                _ = v72
                                                                _ = v73
                                                            default:
                                                                v47 = append(v47, (-(fi38)))
                                                            }
                                                        }
                                                        _ = v52
                                                    }
                                                    if (false && true) {
                                                        switch (("a" + "foo bar") + "") {
                                                        case "foo bar":
                                                            v119 := GenS11{F0: -128, F1: 811}
                                                            _ = v119
                                                        case "a":
                                                            v3 -= (v3 / ((v3) | 1))
                                                        case "":
                                                            var v120 uint8 = 127
                                                            _ = v120
                                                        }
                                                    }
                                                    switch (fi36 & -9223372036854775808) {
                                                    case -128:
                                                        if (!true) {
                                                            v3 = (v3 % (((v3 * v3)) | 1))
                                                            var v121 int = (fi38 ^ v45)
                                                            fmt.Println(v2, fi36, v45)
                                                            _ = v121
                                                        }
                                                    case -9223372036854775808:
                                                        if (!(1.5 > 0.0)) {
                                                            switch fi36 {
                                                            case 42:
                                                                _ = len(v5)
                                                            case 128:
                                                                v122 := GenS13{F0: 0, F1: -1}
                                                                _ = v122
                                                            case -424:
                                                                if len(v47) >= 2 {
                                                                    v47 = v47[1:len(v47)]
                                                                }
                                                            }
                                                            if (1 < 2) {
                                                                var v123 uint = 65535
                                                                _ = v123
                                                            }
                                                            switch (fi36 * v45) {
                                                            case 32768:
                                                                if (42 >= 405) {
                                                                    fmt.Println(v48, v45, fi38)
                                                                }
                                                                var v124 string = ((("a" + "foo bar") + "hello") + ((("\n\t" + "\n\t") + "foo bar") + ("hello" + "\n\t")))
                                                                fmt.Println(v2, v3, fi38)
                                                                _ = v124
                                                            case 42:
                                                                if (true || false) {
                                                                    v125 := GenS14{F0: 913}
                                                                    _ = v125
                                                                }
                                                            case -937:
                                                                v45 = fi36
                                                            }
                                                            if (true && false) {
                                                                v3 = v3
                                                            }
                                                            switch (fi38 & 32768) {
                                                            case -1:
                                                                v126 := GenS4{F0: -563, F1: uint32(v3)}
                                                                _ = v126
                                                            case 4294967295:
                                                                switch "a" {
                                                                case "":
                                                                    v47 = append(v47, v45)
                                                                default:
                                                                    if (v3 <= v3) {
                                                                        if true {
                                                                            var v127 uint32 = 2
                                                                            var v128 int = (-(v45))
                                                                            _ = v127
                                                                            _ = v128
                                                                        } else {
                                                                            var v129 float64 = -2.25
                                                                            var v130 int = v45
                                                                            v131 := GenS10{F0: 42, F1: ""}
                                                                            _ = v129
                                                                            _ = v130
                                                                            _ = v131
                                                                        }
                                                                    } else {
                                                                        v45 += -2147483648
                                                                        v45 += v45
                                                                        v45 += 792
                                                                    }
                                                                }
                                                            case 9223372036854775807:
                                                                if len(v47) >= 2 {
                                                                    v47 = v47[1:len(v47)]
                                                                }
                                                                v3 = v3
                                                            default:
                                                                var v132 string = ("" + (("foo bar" + "a") + "\n\t"))
                                                                _ = v132
                                                            }
                                                        }
                                                    case 127:
                                                        fmt.Println(v2)
                                                    }
                                                }
                                            }
                                            switch 460 {
                                            case 32767:
                                                switch "\n\t" {
                                                case "a":
                                                    switch (fi36 % (((v45 % ((v45) | 1))) | 1)) {
                                                    case 606:
                                                        var v133 int = fi38
                                                        _ = v133
                                                    default:
                                                        var v134 uint = 18446744073709551615
                                                        v135 := GenS11{F0: genFn1(42), F1: 821}
                                                        _ = v134
                                                        _ = v135
                                                    }
                                                default:
                                                    v45 = (fi36 % ((fi36) | 1))
                                                }
                                            case 128:
                                                if true {
                                                    v2 = v2
                                                }
                                            }
                                            fmt.Println(fi38, fi38, v3)
                                            _ = v47
                                            _ = v48
                                        }
                                    }
                                case 9223372036854775807:
                                    v136 := v5[(-(fi38))]
                                    _ = v136
                                }
                                var v137 uint8 = 128
                                _ = v137
                                _ = v46
                            } else {
                                var v138 int32 = 42
                                _ = v138
                            }
                        }
                        _ = fi38
                        _ = v45
                    }
                default:
                    var v139 int8 = -43
                    _ = v139
                }
            case "":
                var v140 bool = true
                _ = v140
            }
            _ = fi36
        }
        _ = v35
    } else {
        var v141 bool = ((-2.25 * (float64(v1) * (-2.25 + 1.5))) == (0.1 * 0.1))
        _ = v141
    }
    v142 := GenS15{F0: 9223372036854775807, F1: true, F2: 42}
    fmt.Println(v3, v3, v3)
    for fi143 := 0; fi143 < 6; fi143++ {
        var v144 uint32 = 229
        _ = fi143
        _ = v144
    }
    for fi145 := 0; fi145 < 17; fi145++ {
        v146 := GenS8{F0: -50}
        v147 := []int{fi145, fi145, fi145, (^(fi145))}
        _ = fi145
        _ = v146
        _ = v147
    }
    v2 = (v2 * (v2 << ((v2) & 63)))
    if true {
        fmt.Println(v1, v2)
        fmt.Println(v2)
        delete(v4, -547)
    }
    v3 = (v3 - v3)
    fmt.Println(v2)
    var v148 int16 = 32767
    v148 -= v148
    if true {
        v149 := []int{126}
        var v150 uint64 = (v3 ^ genFn0(int(v148)))
        switch "\n\t" {
        case "foo bar":
            switch "\n\t" {
            case "foo bar":
                switch 279 {
                case 1:
                    var v151 int64 = 42
                    _ = v151
                case 9223372036854775807:
                    _ = len(v149)
                case 0:
                    switch 42 {
                    case -9223372036854775808:
                        var v152 int64 = -9223372036854775808
                        _ = v152
                    }
                default:
                    v149 = append(v149, 32767)
                }
            case "\n\t":
                fmt.Println(v150, v1, v2)
            }
        default:
            _ = len(v149)
        }
        _ = v149
        _ = v150
    }
    v153 := GenS4{F0: -2147483648, F1: 506}
    fmt.Println("v1:", v1)
    fmt.Println("v148:", v148)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v142
    _ = v148
    _ = v153
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
