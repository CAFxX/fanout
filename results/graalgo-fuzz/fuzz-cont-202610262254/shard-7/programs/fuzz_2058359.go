package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int; F1 bool }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int32; F1 uint32; F2 int16 }
func genFn0(p0 int16, p1 uint8) uint32 {
    var v6 float32 = 0.0
    var v7 uint32 = 65535
    v8 := GenS1{F0: 408, F1: v7}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint = 4294967295
    var v2 string = "hello"
    v3 := []string{((("\n\t" + v2) + (v2 + "\n\t")) + "hello")}
    v4 := make(map[uint32]uint32)
    v5 := []string{(v2 + "a")}
    for fi9 := 0; fi9 < 5; fi9++ {
        v4[1] = 2147483648
        var v10 int8 = 100
        _ = fi9
        _ = v10
    }
    _ = len(v4)
    if false {
        if false {
            v11 := v4[2]
            var v12 int16 = 32767
            _ = v11
            _ = v12
        }
        var v13 float32 = 1.5
        v14 := GenS2{F0: 128}
        _ = v13
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 12; fi15++ {
            fmt.Println(v2, v1)
            v16 := GenS3{F0: int8(fi15)}
            _ = fi15
            _ = v16
        }
        if true {
            v17 := v4[65535]
            if false {
                if false {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        if ((fi18 >> ((1) & 63)) < fi18) {
                            for fi19 := 0; fi19 < 4; fi19++ {
                                if false {
                                    if false {
                                        var v20 bool = true
                                        var v21 uint = v1
                                        _ = v20
                                        _ = v21
                                    }
                                    v2 = "foo bar"
                                } else {
                                    switch (fi19 * (^(fi18))) {
                                    case 9223372036854775807:
                                        v1 -= uint(v17)
                                    case 42:
                                        var v22 uint32 = (^(v17))
                                        _ = v22
                                    case 127:
                                        delete(v4, 65535)
                                        fmt.Println(fi18, fi19, fi18)
                                    }
                                }
                                _ = fi19
                            }
                            var v23 int8 = 3
                            v1 -= (^(v1))
                            _ = v23
                        } else {
                            var v24 int64 = -9223372036854775808
                            if true {
                                v25 := GenS2{F0: 0}
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = v2
                                    }
                                    var v26 uint8 = 255
                                    _ = v26
                                } else {
                                    for fi27 := 0; fi27 < 8; fi27++ {
                                        v1 = (v1 ^ (v1 - v1))
                                        _ = fi27
                                    }
                                    if false {
                                        v28 := GenS1{F0: 705, F1: (^(v17))}
                                        if (false || false) {
                                            v17 = v17
                                            if (1000 < 0) {
                                                if false {
                                                    fmt.Println(fi18, v2)
                                                    v29 := GenS2{F0: 1}
                                                    _ = v29
                                                } else {
                                                    fmt.Println(v2)
                                                }
                                            }
                                        }
                                        _ = v28
                                    } else {
                                        if (!(0.1 >= -2.25)) {
                                            v5 = append(v5, v2)
                                        } else {
                                            fmt.Println(v24, v17)
                                            for fi30 := 0; fi30 < 3; fi30++ {
                                                if (true || false) {
                                                    v4[v17] = (v17 - (v17 | v17))
                                                }
                                                _ = fi30
                                            }
                                        }
                                        switch fi18 {
                                        case 0:
                                            _ = len(v4)
                                        case 32767:
                                            for fi31 := 0; fi31 < 1; fi31++ {
                                                v24 += 9223372036854775807
                                                if ((255 > 2) && false) {
                                                    fmt.Println(v17, v2, v24)
                                                    var v32 int32 = -681
                                                    _ = v32
                                                } else {
                                                    var v33 int = 0
                                                    var v34 bool = true
                                                    _ = v33
                                                    _ = v34
                                                }
                                                _ = fi31
                                            }
                                        case 2147483647:
                                            for fi35 := 0; fi35 < 20; fi35++ {
                                                v2 = v2
                                                v1 = (v1 - (^(v1)))
                                                if true {
                                                    v24 += 858
                                                    switch ("foo bar" + v2) {
                                                    case "":
                                                        var v36 uint = 1
                                                        _ = v36
                                                    case "a":
                                                        fmt.Println(v24, v24)
                                                    }
                                                    var v37 uint32 = v17
                                                    _ = v37
                                                }
                                                _ = fi35
                                            }
                                        default:
                                            _ = len(v5)
                                        }
                                        var v38 int = 42
                                        _ = v38
                                    }
                                }
                                _ = v25
                            } else {
                                if (!true) {
                                    v1 -= 4294967295
                                    switch "hello" {
                                    case "foo bar":
                                        for fi39 := 0; fi39 < 14; fi39++ {
                                            v40 := GenS4{F0: -128, F1: 42}
                                            v41 := GenS5{F0: 1, F1: (v24 & (v24 ^ 570))}
                                            v5 = append(v5, v2)
                                            _ = fi39
                                            _ = v40
                                            _ = v41
                                        }
                                    case "\n\t":
                                        fmt.Println(fi18)
                                    case "":
                                        switch v2 {
                                        case "\n\t":
                                            for fi42 := 0; fi42 < 12; fi42++ {
                                                var v43 bool = false
                                                _ = fi42
                                                _ = v43
                                            }
                                        case "a":
                                            v24 = v24
                                        }
                                    default:
                                        for fi44 := 0; fi44 < 17; fi44++ {
                                            v45 := GenS5{F0: 0, F1: -765}
                                            if false {
                                                if false {
                                                    fmt.Println(v17, v24, fi18)
                                                    var v46 float32 = (float32(-2.25) + float32(1e10))
                                                    _ = v46
                                                } else {
                                                    if len(v5) > 0 {
                                                        v5[0] = v2
                                                    }
                                                    var v47 int64 = v24
                                                    switch v2 {
                                                    case "\n\t":
                                                        if false {
                                                            if (1 >= -1) {
                                                                v48 := GenS6{F0: true, F1: 2, F2: 65535}
                                                                var v49 int = fi18
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                            if true {
                                                                if (1e10 > (float64((((0.0 - 0.1) / float32(-2.25)) * 3.14159)) - float64((-2.25 / 1.0)))) {
                                                                    var v50 uint64 = 42
                                                                    _ = v50
                                                                }
                                                                switch (fi18 | fi44) {
                                                                case 42:
                                                                    var v51 int32 = 42
                                                                    _ = v51
                                                                default:
                                                                    v1 = v1
                                                                }
                                                                v52 := GenS7{F0: (v47 ^ int64(fi44))}
                                                                _ = v52
                                                            }
                                                            var v53 int64 = 0
                                                            _ = v53
                                                        } else {
                                                            v54 := GenS7{F0: int64(v17)}
                                                            v55 := GenS3{F0: -1}
                                                            fmt.Println(v17, v17)
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                    case "":
                                                        switch ("\n\t" + (("foo bar" + v2) + v2)) {
                                                        case "hello":
                                                            if false {
                                                                v24 = 9223372036854775807
                                                                if true {
                                                                    delete(v4, v17)
                                                                    if false {
                                                                        if false {
                                                                            var v56 uint = v1
                                                                            v56 = v56
                                                                            var v57 uint = v1
                                                                            _ = v56
                                                                            _ = v57
                                                                        }
                                                                    }
                                                                } else {
                                                                    _ = len(v4)
                                                                    var v58 uint64 = 437
                                                                    _ = v58
                                                                }
                                                            } else {
                                                                v17 = (^(v17))
                                                                switch (v2 + v2) {
                                                                case "\n\t":
                                                                    var v59 uint32 = v17
                                                                    _ = v59
                                                                case "":
                                                                    var v60 int32 = -958
                                                                    _ = v60
                                                                default:
                                                                    v24 += (v24 * v24)
                                                                }
                                                                v1 += (v1 >> ((42) & 63))
                                                                if (false || (((^(v1)) > (v1 << ((4294967295) & 63))) && true)) {
                                                                    _ = len(v4)
                                                                    v61 := GenS8{F0: 70}
                                                                    _ = v61
                                                                }
                                                            }
                                                            var v62 int16 = 465
                                                            if (0.0 != (float64(0.1) - 3.14159)) {
                                                                var v63 uint8 = 152
                                                                if true {
                                                                    fmt.Println(v24, v2, v62)
                                                                    v47 -= v24
                                                                    if (!true) {
                                                                        v17 = 892
                                                                        v5 = append(v5, ("" + v2))
                                                                        var v64 string = v2
                                                                        _ = v64
                                                                    } else {
                                                                        v2 = "a"
                                                                        var v65 int8 = 100
                                                                        var v66 int = 42
                                                                        _ = v65
                                                                        _ = v66
                                                                    }
                                                                } else {
                                                                    v67 := GenS4{F0: 0, F1: 580}
                                                                    v24 -= (v24 * v24)
                                                                    _ = v67
                                                                }
                                                                _ = v63
                                                            } else {
                                                                v3 = append(v3, v2)
                                                            }
                                                            _ = v62
                                                        case "":
                                                            delete(v4, (^(v17)))
                                                        case "foo bar":
                                                            v17 = v17
                                                        default:
                                                            var v68 int8 = 100
                                                            _ = v68
                                                        }
                                                    case "a":
                                                        fmt.Println(v24, v47, v2)
                                                    }
                                                    _ = v47
                                                }
                                                v69 := GenS5{F0: 42, F1: (v24 ^ (-(v24)))}
                                                _ = v69
                                            }
                                            v2 = ""
                                            _ = len(v3)
                                            _ = fi44
                                            _ = v45
                                        }
                                    }
                                }
                                var v70 string = v2
                                _ = v70
                            }
                            _ = v24
                        }
                        v2 = v2
                        var v71 uint16 = 42
                        _ = fi18
                        _ = v71
                    }
                    var v72 string = (v2 + v2)
                    for fi73 := 0; fi73 < 10; fi73++ {
                        if true {
                            for fi74 := 0; fi74 < 1; fi74++ {
                                v17 += (v17 * (v17 ^ v17))
                                if len(v5) > 0 {
                                    v5[0] = "\n\t"
                                }
                                if (true && (65535 != v1)) {
                                    if (32767 <= 1000) {
                                        if true {
                                            v75 := GenS2{F0: 114}
                                            _ = v75
                                        }
                                        var v76 int32 = 1
                                        fmt.Println(v1, v1)
                                        _ = v76
                                    } else {
                                        var v77 int16 = int16(fi73)
                                        var v78 int32 = int32(v17)
                                        switch fi74 {
                                        case 32767:
                                            var v79 uint16 = 1
                                            var v80 uint = (v1 * (v1 ^ (^(v1))))
                                            _ = v79
                                            _ = v80
                                        case 128:
                                            v77 -= v77
                                        }
                                        _ = v77
                                        _ = v78
                                    }
                                    var v81 uint = v1
                                    switch (fi73 - -545) {
                                    case 0:
                                        if true {
                                            var v82 float32 = -2.25
                                            _ = v82
                                        } else {
                                            v83 := GenS9{F0: fi73, F1: true}
                                            _ = v83
                                        }
                                    case -9223372036854775808:
                                        var v84 int8 = 1
                                        _ = v84
                                    }
                                    _ = v81
                                }
                                _ = fi74
                            }
                            if ((657 <= 1) || ((v1 & 65535) > 643)) {
                                fmt.Println(v72, fi73, v2)
                            }
                        }
                        _ = fi73
                    }
                    _ = v72
                }
                v85 := GenS1{F0: 42, F1: 42}
                _ = v85
            } else {
                for fi86 := 0; fi86 < 3; fi86++ {
                    var v87 uint32 = v17
                    var v88 uint32 = (v17 - (v17 + genFn0(1, 74)))
                    _ = fi86
                    _ = v87
                    _ = v88
                }
                v89 := GenS5{F0: 2147483647, F1: 9223372036854775807}
                for fi90 := 0; fi90 < 5; fi90++ {
                    var v91 string = ("" + v2)
                    v92 := GenS10{F0: true}
                    _ = fi90
                    _ = v91
                    _ = v92
                }
                _ = v89
            }
            _ = v17
        } else {
            if true {
                var v93 uint32 = 2147483647
                fmt.Println(v1, v93, v2)
                _ = v93
            }
        }
    }
    switch -810 {
    case -333:
        fmt.Println(v2)
    default:
        var v94 float32 = ((0.1 / 1e10) - float32(0.1))
        _ = v94
    }
    switch v2 {
    case "a":
        for fi95 := 0; fi95 < 10; fi95++ {
            switch ("\n\t" + v2) {
            case "hello":
                var v96 uint32 = 0
                _ = v96
            default:
                v1 -= v1
            }
            v2 = v2
            _ = fi95
        }
    case "foo bar":
        for fi97 := 0; fi97 < 17; fi97++ {
            fmt.Println(v2)
            for fi98 := 0; fi98 < 12; fi98++ {
                fmt.Println(v2, fi97)
                _ = fi98
            }
            _ = fi97
        }
    default:
        var v99 int = 1
        _ = v99
    }
    fmt.Println(v2, v1)
    v100 := GenS3{F0: 42}
    var v101 int = 0
    v2 = "foo bar"
    switch 127 {
    case -1:
        var v102 int = (v101 * (v101 - v101))
        _ = v102
    case 2147483647:
        switch "foo bar" {
        case "foo bar":
            var v103 uint64 = 61
            _ = v103
        default:
            if (true || true) {
                switch v2 {
                case "\n\t":
                    v104 := GenS11{F0: 1, F1: 4294967295, F2: 0}
                    _ = v104
                case "a":
                    v101 = v101
                case "":
                    v2 = v2
                }
                delete(v4, 2147483648)
                if true {
                    v105 := GenS5{F0: 2147483647, F1: -9223372036854775808}
                    for fi106 := 0; fi106 < 6; fi106++ {
                        v5 = append(v5, (("foo bar" + "") + "hello"))
                        for fi107 := 0; fi107 < 13; fi107++ {
                            switch (v2 + "a") {
                            case "":
                                if (4294967295 > v1) {
                                    v101 -= v101
                                    var v108 int = 128
                                    var v109 int8 = 42
                                    _ = v108
                                    _ = v109
                                }
                            case "a":
                                var v110 bool = (!(0 <= 768))
                                _ = v110
                            }
                            _ = fi107
                        }
                        var v111 uint16 = 117
                        _ = fi106
                        _ = v111
                    }
                    _ = v105
                } else {
                    v1 = v1
                    for fi112 := 0; fi112 < 10; fi112++ {
                        _ = len(v5)
                        _ = fi112
                    }
                    v2 = "\n\t"
                }
            } else {
                if (0 >= 1) {
                    var v113 uint = v1
                    var v114 string = ("\n\t" + "")
                    var v115 string = (((v114 + v114) + ((("a" + (v114 + "foo bar")) + ("a" + v114)) + "\n\t")) + v114)
                    _ = v113
                    _ = v114
                    _ = v115
                } else {
                    for fi116 := 0; fi116 < 15; fi116++ {
                        v1 = (v1 ^ 2)
                        for fi117 := 0; fi117 < 3; fi117++ {
                            var v118 string = v2
                            v2 = ("a" + v118)
                            _ = fi117
                            _ = v118
                        }
                        if len(v3) > 0 {
                            v3[0] = v2
                        }
                        _ = fi116
                    }
                    switch "foo bar" {
                    case "hello":
                        v1 = v1
                    }
                    if (!true) {
                        v3 = append(v3, (v2 + "foo bar"))
                    }
                }
            }
        }
    default:
        for fi119 := 0; fi119 < 2; fi119++ {
            fmt.Println(v101, fi119)
            switch v2 {
            case "hello":
                if (fi119 == (^(v101))) {
                    switch "hello" {
                    case "\n\t":
                        if true {
                            v120 := GenS11{F0: -19, F1: genFn0(int16(fi119), 127), F2: -580}
                            _ = v120
                        }
                    case "":
                        v2 = ("\n\t" + v2)
                    case "a":
                        if true {
                            _ = len(v4)
                        }
                    default:
                        var v121 int16 = 1
                        _ = v121
                    }
                } else {
                    for fi122 := 0; fi122 < 7; fi122++ {
                        var v123 uint16 = 42
                        fmt.Println(v2, fi122, v1)
                        fmt.Println(v1, fi122, v1)
                        _ = fi122
                        _ = v123
                    }
                    var v124 string = ("hello" + v2)
                    v1 += (v1 * (v1 / ((v1) | 1)))
                    _ = v124
                }
            case "\n\t":
                switch (fi119 - fi119) {
                case -104:
                    var v125 float64 = (0.0 / (3.14159 * -2.25))
                    _ = v125
                case 9223372036854775807:
                    var v126 uint8 = 1
                    _ = v126
                case 1:
                    var v127 uint64 = uint64(v1)
                    _ = v127
                default:
                    for fi128 := 0; fi128 < 5; fi128++ {
                        switch ("\n\t" + "\n\t") {
                        case "hello":
                            v2 = v2
                        case "\n\t":
                            var v129 bool = true
                            _ = v129
                        }
                        var v130 float32 = ((-2.25 - ((1e10 - 0.1) + 3.14159)) / 1e10)
                        _ = fi128
                        _ = v130
                    }
                }
            case "foo bar":
                v101 = -862
            default:
                for fi131 := 0; fi131 < 17; fi131++ {
                    switch "hello" {
                    case "hello":
                        v1 = v1
                    }
                    _ = fi131
                }
            }
            fmt.Println(v101, fi119, v1)
            _ = fi119
        }
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
