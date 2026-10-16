package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 string; F1 int64; F2 int8 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 uint; F1 int8; F2 uint8 }
type GenS10 struct { F0 int; F1 int8; F2 bool }
func genFn0(p0 uint, p1 int8, p2 uint64) int32 {
    v6 := []int{128, -198, 1}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 bool = true
        if true {
            v9 := GenS1{F0: p1, F1: 170}
            p1 = (-(p1))
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 604
}
func genFn1() uint8 {
    v10 := GenS2{F0: (42 <= 1)}
    switch "hello" {
    case "\n\t":
        var v11 uint32 = 939
        _ = v11
    case "foo bar":
        switch -1 {
        case -9223372036854775808:
            for fi12 := 0; fi12 < 6; fi12++ {
                if ((!(!true)) || false) {
                    var v13 uint = 2
                    _ = v13
                }
                var v14 uint16 = 255
                _ = fi12
                _ = v14
            }
        case 0:
            switch -1 {
            case 128:
                for fi15 := 0; fi15 < 18; fi15++ {
                    if ("" != "\n\t") {
                        v16 := []int{32768, (fi15 ^ fi15), 0}
                        for fi17 := 0; fi17 < 19; fi17++ {
                            v18 := GenS2{F0: true}
                            switch (("" + "foo bar") + (("" + "hello") + "a")) {
                            case "\n\t":
                                var v19 bool = false
                                _ = v19
                            }
                            fmt.Println(fi15, fi15)
                            _ = fi17
                            _ = v18
                        }
                        var v20 uint32 = 0
                        _ = v16
                        _ = v20
                    }
                    v21 := []int{fi15, 1, fi15, -2147483648}
                    var v22 float64 = 0.1
                    _ = fi15
                    _ = v21
                    _ = v22
                }
            default:
                var v23 string = ""
                _ = v23
            }
        case -2147483648:
            if true {
                v24 := GenS3{F0: 42, F1: 746}
                for fi25 := 0; fi25 < 12; fi25++ {
                    var v26 uint = 65535
                    if false {
                        v26 += (v26 - (v26 + (v26 >> ((67) & 63))))
                        if (!(578 <= 65535)) {
                            var v27 int32 = -2147483648
                            _ = v27
                        } else {
                            v28 := []int{}
                            switch "\n\t" {
                            case "a":
                                if (true && false) {
                                    var v29 int8 = -1
                                    _ = v29
                                }
                            case "\n\t":
                                var v30 uint32 = 739
                                _ = v30
                            }
                            _ = v28
                        }
                    } else {
                        switch "hello" {
                        case "a":
                            if false {
                                var v31 uint64 = 2
                                v32 := GenS2{F0: (false && ((!(!(!false))) && (-1 < 1)))}
                                if true {
                                    fmt.Println(v31, v31, v31)
                                    switch ("hello" + "a") {
                                    case "":
                                        var v33 bool = false
                                        _ = v33
                                    case "hello":
                                        if (("" + (("a" + "\n\t") + "\n\t")) != "") {
                                            for fi34 := 0; fi34 < 7; fi34++ {
                                                v35 := []int{}
                                                _ = fi34
                                                _ = v35
                                            }
                                            for fi36 := 0; fi36 < 9; fi36++ {
                                                fmt.Println(v31)
                                                if (100 > 100) {
                                                    var v37 int64 = -9223372036854775808
                                                    fmt.Println(v31, v26, v37)
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        fmt.Println(fi36, v26, v31)
                                                    case "foo bar":
                                                        if true {
                                                            v38 := GenS4{F0: v26, F1: 100}
                                                            if true {
                                                                v26 = v26
                                                                v39 := GenS4{F0: (v26 % ((65535) | 1)), F1: 42}
                                                                v40 := []int{(fi25 << (((v26 & (v26 + 255))) & 63)), (fi25 >> ((v26) & 63)), fi25}
                                                                v37 -= 0
                                                                v37 -= v37
                                                                _ = v39
                                                                _ = v40
                                                            }
                                                            _ = v38
                                                        } else {
                                                            v41 := GenS3{F0: 32767, F1: 42}
                                                            _ = v41
                                                        }
                                                    }
                                                    _ = v37
                                                }
                                                var v42 int64 = 9223372036854775807
                                                _ = fi36
                                                _ = v42
                                            }
                                            switch 945 {
                                            case -353:
                                                for fi43 := 0; fi43 < 18; fi43++ {
                                                    var v44 string = ""
                                                    var v45 uint32 = 1
                                                    fmt.Println(fi25, v44, v44)
                                                    if false {
                                                        var v46 float32 = 1.5
                                                        fmt.Println(v44, v31, v26)
                                                        v46 = 0.0
                                                        _ = v46
                                                    }
                                                    var v47 uint32 = 543
                                                    _ = fi43
                                                    _ = v44
                                                    _ = v45
                                                    _ = v47
                                                }
                                            default:
                                                v48 := GenS5{F0: (("" + ("a" + "\n\t")) + (("" + "") + "foo bar")), F1: 9223372036854775807, F2: 1}
                                                _ = v48
                                            }
                                        } else {
                                            var v49 int32 = -2147483648
                                            _ = v49
                                        }
                                    }
                                } else {
                                    switch "a" {
                                    case "\n\t":
                                        v31 = v31
                                    }
                                }
                                _ = v31
                                _ = v32
                            }
                        case "foo bar":
                            v50 := GenS1{F0: 0, F1: 2}
                            _ = v50
                        }
                    }
                    fmt.Println(v26, fi25)
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
        default:
            for fi51 := 0; fi51 < 17; fi51++ {
                v52 := GenS6{F0: -32768, F1: 18446744073709551615}
                switch "a" {
                case "hello":
                    v53 := []int{}
                    _ = v53
                case "foo bar":
                    v54 := GenS4{F0: 631, F1: -1}
                    _ = v54
                default:
                    var v55 int16 = 1
                    _ = v55
                }
                var v56 int16 = -640
                _ = fi51
                _ = v52
                _ = v56
            }
        }
    default:
        v57 := GenS7{F0: 0}
        _ = v57
    }
    _ = v10
    return 127
}

func main() {
    v1 := []uint32{967, 42, 65535}
    var v2 int = 709
    v3 := []string{}
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]uint32)
    if ((!(true && (true || true))) && (543 > 132)) {
        var v58 int32 = 956
        var v59 string = (("" + "\n\t") + "a")
        switch 920 {
        case 317:
            var v60 uint = 255
            _ = v60
        case 32768:
            fmt.Println(v2, v58)
        case 4294967295:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = v58
        _ = v59
    } else {
        v2 = v2
    }
    for fi61 := 0; fi61 < 15; fi61++ {
        switch -1 {
        case 483:
            v2 = 593
        case 32767:
            if (42 != 42) {
                var v62 float64 = (0.0 * 0.0)
                for fi63 := 0; fi63 < 16; fi63++ {
                    if (581 != fi63) {
                        if (906 <= 297) {
                            switch "hello" {
                            case "foo bar":
                                switch v2 {
                                case 1:
                                    switch (v2 * -732) {
                                    case 42:
                                        v62 += (float64(1.5) + (((v62 * v62) + ((float64(float32(1e10)) + (3.14159 - 3.14159)) + v62)) * float64(float32(((v62 / v62) + v62)))))
                                    case 0:
                                        v2 += 401
                                        fmt.Println(v2, fi63, v2)
                                    default:
                                        v64 := GenS8{F0: "foo bar", F1: 0}
                                        _ = v64
                                    }
                                case -9223372036854775808:
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v65 float64 = v62
                                        _ = v65
                                    default:
                                        v62 = v62
                                    }
                                case 127:
                                    if false {
                                        fmt.Println(v2)
                                        switch "" {
                                        case "hello":
                                            fmt.Println(v2, v2, fi63)
                                        }
                                        var v66 int32 = -180
                                        _ = v66
                                    } else {
                                        if false {
                                            v62 += float64((0.1 * ((float32(3.14159) / 1.0) / 1.0)))
                                        }
                                        fmt.Println(v62, fi63)
                                    }
                                default:
                                    if (!(0.1 >= 1e10)) {
                                        v62 += v62
                                        v67 := GenS7{F0: 0}
                                        _ = v67
                                    } else {
                                        v68 := GenS9{F0: 65535, F1: -72, F2: 0}
                                        _ = v68
                                    }
                                }
                            }
                            if true {
                                fmt.Println(fi61, v62, v2)
                            }
                            var v69 int = 242
                            _ = v69
                        }
                    } else {
                        v2 -= fi61
                    }
                    _ = fi63
                }
                fmt.Println(fi61, fi61)
                _ = v62
            } else {
                if (100 <= 100) {
                    for fi70 := 0; fi70 < 6; fi70++ {
                        fmt.Println(fi61, fi61)
                        var v71 int8 = 100
                        if ((float64(float32((0.1 * (1.5 + (0.1 * 0.1))))) - (((3.14159 - 1.5) * 0.1) - (((1.5 * 1.5) / 1.0) * (0.1 * float64(3.14159))))) == (-2.25 - (1e10 + (((3.14159 - 1.5) - 1.5) - 0.1)))) {
                            if len(v3) > 0 {
                                v3[0] = ""
                            }
                        }
                        _ = fi70
                        _ = v71
                    }
                    var v72 uint16 = 895
                    switch 1 {
                    case 989:
                        v2 = v2
                    case 127:
                        var v73 uint64 = 147
                        _ = v73
                    default:
                        for fi74 := 0; fi74 < 20; fi74++ {
                            var v75 uint8 = 127
                            _ = fi74
                            _ = v75
                        }
                    }
                    _ = v72
                }
                var v76 int8 = -84
                _ = v76
            }
        }
        _ = fi61
    }
    var v77 uint64 = 18446744073709551615
    var v78 int8 = 51
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v79 float32 = (0.1 - 1e10)
    switch v2 {
    case 0:
        var v80 float32 = float32(0.1)
        _ = v80
    case 118:
        v78 -= 0
    case 100:
        if (((-(v2)) > -1) || true) {
            v81 := GenS8{F0: "", F1: 858}
            switch (v2 & 1) {
            case 9223372036854775807:
                switch int(v77) {
                case 4294967295:
                    var v82 uint32 = 426
                    _ = v82
                case -301:
                    switch ("a" + (("a" + "foo bar") + "a")) {
                    case "foo bar":
                        if false {
                            switch v2 {
                            case 32767:
                                switch "foo bar" {
                                case "a":
                                    if (324 > 0) {
                                        fmt.Println(v79)
                                        for fi83 := 0; fi83 < 12; fi83++ {
                                            fmt.Println(v78)
                                            _ = fi83
                                        }
                                    }
                                case "\n\t":
                                    switch (^(v2)) {
                                    case 128:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case 999:
                                        v77 += v77
                                    }
                                default:
                                    v77 = (v77 - 451)
                                }
                            }
                            var v84 uint16 = 779
                            _ = v84
                        } else {
                            if true {
                                var v85 uint16 = 670
                                switch "foo bar" {
                                case "hello":
                                    switch v2 {
                                    case 98:
                                        v77 -= (^(v77))
                                    case -880:
                                        var v86 uint64 = uint64(v85)
                                        _ = v86
                                    }
                                case "":
                                    var v87 float32 = 1.5
                                    _ = v87
                                }
                                _ = v85
                            }
                            v88 := GenS10{F0: v2, F1: v78, F2: (v77 < 42)}
                            _ = v88
                        }
                    default:
                        var v89 uint = 1
                        _ = v89
                    }
                case -1:
                    for fi90 := 0; fi90 < 10; fi90++ {
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                        v2 -= fi90
                        _ = fi90
                    }
                default:
                    fmt.Println(v77, v78, v79)
                }
            case 32768:
                _ = len(v5)
            }
            var v91 uint32 = 0
            _ = v81
            _ = v91
        } else {
            switch v2 {
            case -9223372036854775808:
                for fi92 := 0; fi92 < 10; fi92++ {
                    var v93 float64 = 1.5
                    v3 = append(v3, "")
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                    _ = fi92
                    _ = v93
                }
            case 405:
                if true {
                    var v94 int = v2
                    for fi95 := 0; fi95 < 6; fi95++ {
                        for fi96 := 0; fi96 < 14; fi96++ {
                            if false {
                                var v97 int32 = 0
                                _ = v97
                            } else {
                                var v98 uint16 = 997
                                var v99 int64 = 589
                                v100 := GenS3{F0: 42, F1: 1}
                                _ = v100
                                _ = v98
                                _ = v99
                            }
                            _ = fi96
                        }
                        _ = fi95
                    }
                    _ = v94
                }
            case -206:
                v78 += (v78 % ((v78) | 1))
            default:
                var v101 uint16 = 32767
                _ = v101
            }
        }
    default:
        for fi102 := 0; fi102 < 20; fi102++ {
            var v103 int16 = 42
            _ = fi102
            _ = v103
        }
    }
    for fi104 := 0; fi104 < 11; fi104++ {
        if true {
            v79 = v79
            if (2 >= 1) {
                for fi105 := 0; fi105 < 10; fi105++ {
                    v106 := GenS3{F0: 1000, F1: 4294967295}
                    fmt.Println(v78)
                    _ = fi105
                    _ = v106
                }
            }
        } else {
            switch "a" {
            case "":
                fmt.Println(fi104)
            case "\n\t":
                for fi107 := 0; fi107 < 1; fi107++ {
                    if (false || false) {
                        var v108 bool = (1000 < 1000)
                        v109 := GenS3{F0: 0, F1: 990}
                        _ = v108
                        _ = v109
                    } else {
                        var v110 int64 = -9223372036854775808
                        _ = v110
                    }
                    switch "a" {
                    case "hello":
                        v77 += (^(v77))
                    case "a":
                        v2 += fi107
                        v79 = float32(float64(0.0))
                    default:
                        _ = len(v5)
                    }
                    _ = fi107
                }
            }
        }
        _ = fi104
    }
    var v111 uint = 185
    if (1 > 1) {
        var v112 uint32 = 1
        _ = v112
    }
    v113 := v5[728]
    if true {
        var v114 bool = true
        _ = v114
    } else {
        var v115 float32 = ((0.1 / 1e10) * v79)
        v115 += v79
        _ = v115
    }
    v116 := GenS6{F0: -1, F1: (v77 - (^(v77)))}
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v111:", v111)
    fmt.Println("v113:", v113)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v111
    _ = v113
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v78
    _ = v79
}
