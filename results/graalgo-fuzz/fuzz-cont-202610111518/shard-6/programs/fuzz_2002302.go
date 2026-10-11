package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int; F2 uint64 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int; F2 uint16 }
type GenS6 struct { F0 int; F1 bool; F2 string }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS9 struct { F0 uint; F1 string; F2 uint32 }
type GenS10 struct { F0 int8; F1 uint8; F2 int }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint; F1 int8 }
type GenS14 struct { F0 int64; F1 int; F2 int32 }
func genFn0() uint16 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint64 = 4294967295
        _ = fi6
        _ = v7
    }
    var v8 float64 = (-2.25 / float64(0.1))
    if (false || (!(679 == 4294967295))) {
        v9 := []int{}
        _ = len(v9)
        _ = v9
    } else {
        v8 = -2.25
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: fi10}
        _ = fi10
        _ = v11
    }
    if (1 == 42) {
        v8 = v8
        for fi12 := 0; fi12 < 6; fi12++ {
            fmt.Println(v8)
            if false {
                v8 = v8
                fmt.Println(v8, v8)
            }
            _ = fi12
        }
    } else {
        v8 += v8
        v8 = v8
        fmt.Println(v8)
    }
    _ = v8
    return 176
}
func genFn1() int8 {
    if (151 > 125) {
        v13 := []int{32768}
        if false {
            var v14 int16 = -1
            var v15 uint8 = 127
            _ = v14
            _ = v15
        }
        var v16 uint = 255
        _ = v13
        _ = v16
    } else {
        switch -912 {
        case 900:
            var v17 uint = 18446744073709551615
            _ = v17
        case -9223372036854775808:
            var v18 float64 = 0.1
            _ = v18
        default:
            var v19 string = ("a" + "\n\t")
            _ = v19
        }
        switch "foo bar" {
        case "hello":
            if (-32768 == 42) {
                var v20 int8 = 86
                _ = v20
            }
        }
    }
    if true {
        var v21 uint16 = genFn0()
        var v22 int64 = 1
        v22 -= v22
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 11; fi23++ {
            v24 := []int{87, 2147483647, (fi23 ^ fi23)}
            if false {
                switch "foo bar" {
                case "\n\t":
                    var v25 uint = 0
                    _ = v25
                case "hello":
                    var v26 string = "foo bar"
                    _ = v26
                default:
                    for fi27 := 0; fi27 < 6; fi27++ {
                        var v28 string = "\n\t"
                        fmt.Println(v28)
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                        _ = fi27
                        _ = v28
                    }
                }
                var v29 float32 = float32((0.0 / (0.0 + (0.1 * -2.25))))
                v30 := GenS1{F0: fi23}
                _ = v29
                _ = v30
            } else {
                var v31 uint64 = 218
                if (((1.5 * float32(float64((1e10 + (float32((3.14159 * 0.0)) - (1.5 / 1e10)))))) - 1.5) > (0.1 * 1e10)) {
                    var v32 uint64 = 4294967295
                    var v33 float32 = float32(0.1)
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v35 := GenS2{F0: 2, F1: (fi34 ^ (fi34 << ((438) & 63))), F2: 351}
                        if len(v24) > 0 {
                            v24[0] = (fi34 % ((fi23) | 1))
                        }
                        v31 = 57
                        _ = fi34
                        _ = v35
                    }
                    _ = v32
                    _ = v33
                }
                v31 -= v31
                _ = v31
            }
            _ = fi23
            _ = v24
        }
    }
    return 127
}

func main() {
    var v1 bool = true
    var v2 int = 773
    v3 := []uint32{2147483648, 2147483648, 65535}
    var v4 uint = 255
    v5 := make(map[int]bool)
    v4 = 18446744073709551615
    var v36 int16 = 42
    var v37 uint8 = 2
    var v38 uint8 = v37
    for fi39 := 0; fi39 < 20; fi39++ {
        fmt.Println(fi39)
        fmt.Println(fi39)
        switch (v2 & (v2 << (((^(v4))) & 63))) {
        case 4294967295:
            var v40 float32 = 0.1
            _ = v40
        case 773:
            var v41 bool = v1
            _ = v41
        default:
            if v1 {
                v37 = 128
                v4 = (v4 + v4)
                v1 = (!v1)
            } else {
                v42 := GenS3{F0: 401, F1: (v4 != 889)}
                v37 -= v38
                _ = v42
            }
        }
        _ = fi39
    }
    var v43 float64 = 1.5
    fmt.Println(v43, v36)
    var v44 uint16 = 1
    fmt.Println(v38, v36, v36)
    switch "\n\t" {
    case "\n\t":
        for fi45 := 0; fi45 < 18; fi45++ {
            if (2 == 769) {
                v43 = 0.0
            } else {
                for fi46 := 0; fi46 < 13; fi46++ {
                    switch v2 {
                    case -9223372036854775808:
                        if true {
                            v47 := GenS3{F0: genFn0(), F1: v1}
                            if ("hello" == ("" + "foo bar")) {
                                if v1 {
                                    v1 = ((fi45 | fi46) < int(v44))
                                    if v1 {
                                        v44 -= v44
                                        v48 := GenS4{F0: 2147483648}
                                        _ = v48
                                    }
                                    v36 -= 321
                                }
                            }
                            _ = v47
                        } else {
                            var v49 string = ""
                            switch fi45 {
                            case -439:
                                var v50 string = ""
                                _ = v50
                            case 128:
                                var v51 uint64 = 831
                                switch "" {
                                case "a":
                                    fmt.Println(v51)
                                case "\n\t":
                                    switch v2 {
                                    case 4294967295:
                                        if (!("" != "hello")) {
                                            v52 := v5[(fi46 * int(v38))]
                                            _ = v52
                                        } else {
                                            v53 := GenS4{F0: 0}
                                            v49 = v49
                                            switch (fi45 | (fi45 | v2)) {
                                            case 1:
                                                _ = len(v5)
                                            }
                                            fmt.Println(v36)
                                            v54 := GenS5{F0: v51, F1: (v2 + -9223372036854775808), F2: (v44 >> ((2) & 63))}
                                            v55 := GenS2{F0: 2, F1: int(v38), F2: v51}
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        }
                                    case 42:
                                        v56 := GenS6{F0: fi46, F1: ((v49 == (v49 + v49)) || false), F2: v49}
                                        _ = v56
                                    case 270:
                                        var v57 int16 = v36
                                        _ = v57
                                    default:
                                        fmt.Println(v43, v44, fi46)
                                    }
                                case "":
                                    var v58 float32 = 1.5
                                    _ = v58
                                }
                                switch (fi45 % ((-312) | 1)) {
                                case 0:
                                    var v59 int = (fi46 | -9223372036854775808)
                                    _ = v59
                                case 32768:
                                    v60 := GenS4{F0: 459}
                                    _ = v60
                                case 32767:
                                    v36 += v36
                                default:
                                    v4 = 42
                                }
                                fmt.Println(v4)
                                switch v2 {
                                case 127:
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                    v44 = v44
                                    v61 := GenS2{F0: v44, F1: -2147483648, F2: v51}
                                    if v1 {
                                        var v62 uint = 2
                                        v36 -= (v36 / (((v36 % (((^(v36))) | 1))) | 1))
                                        fmt.Println(v1)
                                        _ = v62
                                    }
                                    var v63 uint32 = uint32(v51)
                                    _ = v61
                                    _ = v63
                                }
                                v38 = (v37 << (((v4 >> ((v4) & 63))) & 63))
                                v64 := GenS7{F0: 43, F1: -1}
                                _ = v51
                                _ = v64
                            case -523:
                                if v1 {
                                    if (426 == 4294967295) {
                                        _ = len(v5)
                                    }
                                } else {
                                    _ = len(v3)
                                    v65 := GenS2{F0: uint16(v4), F1: (fi45 * fi45), F2: 42}
                                    if (true || (v43 == v43)) {
                                        var v66 uint32 = 136
                                        v3 = append(v3, (v66 * v66))
                                        _ = v66
                                    } else {
                                        v36 = 778
                                        delete(v5, (v2 & (v2 & 126)))
                                    }
                                    v36 -= v36
                                    _ = v65
                                }
                            }
                            _ = v49
                        }
                    }
                    v37 -= v38
                    _ = fi46
                }
            }
            v1 = (!v1)
            for fi67 := 0; fi67 < 5; fi67++ {
                if v1 {
                    switch "foo bar" {
                    case "a":
                        var v68 uint64 = 4294967295
                        _ = v68
                    case "":
                        switch ("a" + "hello") {
                        case "hello":
                            if v1 {
                                var v69 int32 = 775
                                if true {
                                    v3 = append(v3, 0)
                                    v38 = v37
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        fmt.Println(v43, fi45, fi45)
                                    case "":
                                        if (!(!v1)) {
                                            switch ("hello" + ("hello" + "a")) {
                                            case "foo bar":
                                                var v70 uint64 = uint64(v44)
                                                _ = v70
                                            case "":
                                                var v71 int8 = 0
                                                _ = v71
                                            case "a":
                                                v36 = (v36 * int16(v38))
                                            default:
                                                var v72 uint32 = 2147483648
                                                _ = v72
                                            }
                                            fmt.Println(v43, v2)
                                        }
                                    case "hello":
                                        v2 = v2
                                    default:
                                        var v73 int16 = v36
                                        _ = v73
                                    }
                                }
                                fmt.Println(fi45)
                                v4 -= 255
                                _ = v69
                            }
                        default:
                            if v1 {
                                if (-1 <= 74) {
                                    var v74 uint16 = (v44 % (((^(v44))) | 1))
                                    if (v43 != -2.25) {
                                        v1 = v1
                                    } else {
                                        if v1 {
                                            var v75 int = fi67
                                            _ = v75
                                        } else {
                                            if ((^(v44)) < v44) {
                                                if v1 {
                                                    var v76 int = (fi45 + fi45)
                                                    _ = v76
                                                } else {
                                                    if (v43 == v43) {
                                                        var v77 int32 = -776
                                                        v1 = ((v36 & v36) != (-(v36)))
                                                        v78 := GenS8{F0: 259, F1: 781, F2: int64(v4)}
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                    var v79 uint16 = v44
                                                    var v80 int = fi67
                                                    var v81 int8 = genFn1()
                                                    _ = v79
                                                    _ = v80
                                                    _ = v81
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            v82 := GenS9{F0: (v4 | (v4 | (v4 | v4))), F1: ("" + "hello"), F2: 450}
                                            _ = v82
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            v44 = v44
                                        case "hello":
                                            v38 = (v38 | (^(v38)))
                                        case "a":
                                            v1 = ((fi67 < (fi45 % (((fi67 * fi45)) | 1))) || ((v4 < (v4 * (v4 + (v4 >> ((v4) & 63))))) || (v37 != uint8(v4))))
                                        default:
                                            v4 = v4
                                        }
                                    }
                                    _ = v74
                                }
                                v83 := GenS1{F0: fi45}
                                _ = v83
                            }
                            fmt.Println(v37, v38, v1)
                            if v1 {
                                v2 = (v2 >> ((2) & 63))
                                fmt.Println(v2)
                                if v1 {
                                    fmt.Println(v2, v44, fi67)
                                    if true {
                                        var v84 float32 = 1.5
                                        _ = v84
                                    }
                                    v1 = (90 < 42)
                                } else {
                                    if (!v1) {
                                        v1 = (2147483647 <= 12)
                                    } else {
                                        _ = len(v5)
                                        var v85 uint32 = uint32(v4)
                                        v2 = fi45
                                        _ = v85
                                    }
                                    v86 := GenS3{F0: v44, F1: true}
                                    _ = v86
                                }
                            }
                        }
                    default:
                        if (!v1) {
                            if v1 {
                                if (2147483648 >= 2147483647) {
                                    var v87 uint16 = (v44 ^ (v44 * (v44 + (v44 & genFn0()))))
                                    fmt.Println(v87, fi45)
                                    _ = v87
                                }
                            } else {
                                switch (fi67 % ((127) | 1)) {
                                case -620:
                                    var v88 bool = (v43 == v43)
                                    _ = v88
                                case -2147483648:
                                    if (!(false && v1)) {
                                        var v89 uint32 = 2147483647
                                        _ = v89
                                    } else {
                                        var v90 float32 = 1.5
                                        v91 := GenS10{F0: 12, F1: (v38 % (((v37 ^ v37)) | 1)), F2: 4294967295}
                                        v1 = (v43 > v43)
                                        _ = v90
                                        _ = v91
                                    }
                                case 1:
                                    delete(v5, fi45)
                                }
                                if v1 {
                                    delete(v5, fi45)
                                    fmt.Println(fi45, v2, v1)
                                }
                            }
                        } else {
                            var v92 uint = (v4 - 1)
                            v36 -= (v36 % (((^(v36))) | 1))
                            _ = v92
                        }
                    }
                }
                v93 := GenS11{F0: 538}
                _ = fi67
                _ = v93
            }
            _ = fi45
        }
    case "":
        v94 := GenS12{F0: (v44 - genFn0())}
        _ = v94
    default:
        var v95 int16 = 42
        _ = v95
    }
    switch ((("hello" + "a") + "") + (("foo bar" + ("foo bar" + ((("" + "a") + "") + "\n\t"))) + "\n\t")) {
    case "":
        v36 -= -1
    case "a":
        switch v2 {
        case 2147483647:
            for fi96 := 0; fi96 < 11; fi96++ {
                for fi97 := 0; fi97 < 5; fi97++ {
                    switch fi96 {
                    case 4294967295:
                        v44 -= (^(v44))
                    }
                    _ = fi97
                }
                if v1 {
                    v2 += -99
                    v98 := v5[-267]
                    v36 += v36
                    _ = v98
                } else {
                    v43 += ((float64(1e10) + (0.1 + v43)) * 0.1)
                    for fi99 := 0; fi99 < 12; fi99++ {
                        var v100 int16 = (v36 - (v36 << (((v4 + (v4 * (^(v4))))) & 63)))
                        _ = fi99
                        _ = v100
                    }
                    v44 += v44
                }
                var v101 int32 = -409
                _ = fi96
                _ = v101
            }
        case 32768:
            switch ("hello" + "a") {
            case "a":
                if ((float32((float64(1e10) - v43)) >= (0.0 * 3.14159)) || ((-794 <= 42) && (!true))) {
                    var v102 uint16 = (^(v44))
                    _ = v102
                }
            case "hello":
                if (1e10 == float32(1e10)) {
                    if (v43 >= v43) {
                        v103 := GenS11{F0: 42}
                        var v104 int64 = -9223372036854775808
                        v105 := GenS13{F0: v4, F1: 127}
                        _ = v103
                        _ = v104
                        _ = v105
                    } else {
                        v106 := GenS8{F0: 4294967295, F1: 1, F2: 42}
                        for fi107 := 0; fi107 < 18; fi107++ {
                            var v108 uint64 = 4294967295
                            _ = fi107
                            _ = v108
                        }
                        fmt.Println(v1)
                        _ = v106
                    }
                    for fi109 := 0; fi109 < 14; fi109++ {
                        var v110 float64 = v43
                        v111 := GenS14{F0: -246, F1: (v2 - (v2 / ((v2) | 1))), F2: 2147483647}
                        fmt.Println(v37)
                        _ = fi109
                        _ = v110
                        _ = v111
                    }
                    fmt.Println(v44, v38)
                }
            case "\n\t":
                var v112 float64 = (v43 - (0.0 - v43))
                _ = v112
            default:
                v37 -= v38
            }
        case 4294967295:
            v5[-662] = v1
        default:
            switch (("\n\t" + "hello") + "hello") {
            case "\n\t":
                var v113 float64 = 0.0
                _ = v113
            default:
                for fi114 := 0; fi114 < 13; fi114++ {
                    var v115 int32 = 42
                    var v116 uint8 = v37
                    var v117 int64 = 87
                    _ = fi114
                    _ = v115
                    _ = v116
                    _ = v117
                }
            }
        }
    case "hello":
        if (v1 && (v36 != v36)) {
            v36 = v36
            if (true && (!v1)) {
                if ((65535 == v4) || (false && v1)) {
                    var v118 int32 = -2147483648
                    _ = v118
                } else {
                    var v119 uint32 = 534
                    _ = v119
                }
                switch (("hello" + "\n\t") + (("a" + "") + "hello")) {
                case "a":
                    v120 := GenS11{F0: 2}
                    _ = v120
                case "foo bar":
                    for fi121 := 0; fi121 < 5; fi121++ {
                        v122 := GenS7{F0: (v37 + (v38 >> ((v4) & 63))), F1: genFn1()}
                        _ = fi121
                        _ = v122
                    }
                case "":
                    var v123 string = (((("\n\t" + "\n\t") + "\n\t") + "foo bar") + "a")
                    _ = v123
                }
                fmt.Println(v38)
            }
            v37 += v37
        } else {
            if (0 > genFn0()) {
                var v124 uint16 = (v44 ^ (^(v44)))
                _ = v124
            } else {
                if (!(!(!((v36 <= v36) || (!((^(v36)) <= 0)))))) {
                    fmt.Println(v43, v37)
                    var v125 int = v2
                    _ = v125
                }
            }
        }
    default:
        var v126 uint = (v4 - (v4 - (v4 + v4)))
        _ = v126
    }
    var v127 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v127
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
