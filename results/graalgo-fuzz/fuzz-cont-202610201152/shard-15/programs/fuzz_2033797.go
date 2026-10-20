package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint; F1 int64; F2 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 int }
func genFn0(p0 bool) uint {
    var v6 float32 = (-2.25 + 0.1)
    switch "" {
    case "foo bar":
        v7 := GenS1{F0: 338, F1: 650}
        _ = v7
    case "\n\t":
        if false {
            v8 := []int{32768, 32767, 128}
            p0 = p0
            switch 32767 {
            case 127:
                var v9 float64 = 1e10
                _ = v9
            case 469:
                fmt.Println(v6, p0)
            case 32768:
                v10 := GenS1{F0: 2147483648, F1: 1}
                _ = v10
            default:
                if (p0 || true) {
                    for fi11 := 0; fi11 < 16; fi11++ {
                        for fi12 := 0; fi12 < 15; fi12++ {
                            var v13 uint16 = 413
                            _ = fi12
                            _ = v13
                        }
                        _ = fi11
                    }
                    v14 := GenS1{F0: 43, F1: 42}
                    var v15 uint64 = 4294967295
                    _ = v14
                    _ = v15
                }
            }
            _ = v8
        } else {
            if p0 {
                var v16 float64 = 0.1
                v17 := []int{-2147483648, -2147483648, -9223372036854775808}
                v16 = (v16 + 3.14159)
                _ = v16
                _ = v17
            }
        }
    default:
        p0 = p0
    }
    var v18 int32 = -2147483648
    v19 := GenS2{F0: 0, F1: 1}
    fmt.Println(v18, v6)
    _ = p0
    _ = v18
    _ = v19
    _ = v6
    return 255
}

func main() {
    v1 := []uint32{65535}
    v2 := make(map[uint32]int)
    var v3 uint = 115
    var v4 float32 = (1e10 - 0.0)
    var v5 string = ("\n\t" + (("foo bar" + "") + ""))
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 10; fi21++ {
            _ = len(v2)
            _ = fi21
        }
        _ = fi20
    }
    v5 = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch (v5 + v5) {
    case "\n\t":
        delete(v2, 4294967295)
    default:
        var v22 uint32 = 0
        _ = v22
    }
    delete(v2, 1)
    for fi23 := 0; fi23 < 9; fi23++ {
        switch v5 {
        case "hello":
            if (((!false) && ((1e10 - 0.1) < 1e10)) || (fi23 != 127)) {
                var v24 float64 = 0.0
                var v25 int16 = 42
                _ = v24
                _ = v25
            } else {
                v3 = (v3 >> (((v3 ^ (v3 >> (((v3 / (((v3 << (((v3 % ((genFn0(false)) | 1))) & 63))) | 1))) & 63)))) & 63))
            }
        case "\n\t":
            v2[4294967295] = 4294967295
        default:
            v5 = "foo bar"
        }
        _ = fi23
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        v5 = v5
        for fi27 := 0; fi27 < 6; fi27++ {
            if true {
                var v28 int32 = -2147483648
                _ = v28
            }
            switch ((v5 + v5) + "hello") {
            case "foo bar":
                var v29 string = "\n\t"
                fmt.Println(v3, fi27, v4)
                _ = v29
            }
            switch (fi27 + fi26) {
            case 1:
                if false {
                    var v30 uint64 = 42
                    _ = v30
                }
            case -9223372036854775808:
                fmt.Println(fi27, v4, fi27)
            }
            _ = fi27
        }
        _ = fi26
    }
    v31 := GenS1{F0: 759, F1: 641}
    fmt.Println(v5)
    if (true || true) {
        if (false && ((1 <= 255) || (false && (!(-128 >= 42))))) {
            for fi32 := 0; fi32 < 12; fi32++ {
                for fi33 := 0; fi33 < 7; fi33++ {
                    if (4294967295 != 2) {
                        if (-1 < 42) {
                            v34 := GenS3{F0: 0}
                            switch v5 {
                            case "a":
                                fmt.Println(fi32, v4, v3)
                            case "":
                                if true {
                                    v5 = "foo bar"
                                    var v35 uint = (v3 + genFn0((!false)))
                                    _ = v35
                                }
                            }
                            if true {
                                var v36 float32 = (float32(1e10) * v4)
                                v37 := GenS1{F0: 2, F1: 42}
                                _ = v36
                                _ = v37
                            }
                            _ = v34
                        } else {
                            _ = len(v2)
                            var v38 float64 = 3.14159
                            _ = v38
                        }
                        if (!false) {
                            var v39 int = (fi33 << (((v3 >> ((v3) & 63))) & 63))
                            _ = v39
                        } else {
                            if ((^(fi32)) == -9223372036854775808) {
                                var v40 float64 = (0.0 * 3.14159)
                                v40 = v40
                                _ = v40
                            } else {
                                fmt.Println(fi33)
                                _ = len(v2)
                                var v41 int8 = 127
                                _ = v41
                            }
                            var v42 float64 = 3.14159
                            _ = len(v2)
                            v43 := GenS2{F0: 2, F1: 127}
                            var v44 float64 = v42
                            v4 = float32(-2.25)
                            v5 = (v5 + (v5 + ("hello" + "hello")))
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        v45 := GenS3{F0: 457}
                        var v46 int32 = 332
                        var v47 uint64 = 312
                        var v48 int32 = (v46 + -1)
                        _ = v45
                        _ = v46
                        _ = v47
                        _ = v48
                    }
                    switch (fi32 & (fi32 | fi33)) {
                    case 4294967295:
                        v49 := GenS4{F0: ""}
                        _ = v49
                    }
                    var v50 float32 = 3.14159
                    _ = fi33
                    _ = v50
                }
                for fi51 := 0; fi51 < 12; fi51++ {
                    var v52 uint32 = 852
                    if true {
                        var v53 uint8 = 243
                        _ = v53
                    }
                    v4 -= float32(float64(1.5))
                    v4 += 3.14159
                    _ = fi51
                    _ = v52
                }
                _ = fi32
            }
            if (2147483648 == 2147483648) {
                delete(v2, 0)
                v54 := GenS5{F0: 1, F1: 129, F2: 42}
                var v55 uint64 = 580
                _ = v54
                _ = v55
            } else {
                v4 = 1e10
                for fi56 := 0; fi56 < 14; fi56++ {
                    var v57 int16 = -326
                    _ = fi56
                    _ = v57
                }
                switch 291 {
                case -120:
                    if true {
                        var v58 int16 = -26
                        _ = v58
                    } else {
                        var v59 uint8 = 42
                        for fi60 := 0; fi60 < 5; fi60++ {
                            var v61 uint16 = 45
                            _ = len(v2)
                            for fi62 := 0; fi62 < 2; fi62++ {
                                fmt.Println(v4, v4)
                                var v63 float64 = (1.5 - (((0.0 - 0.0) * 1e10) + 1e10))
                                _ = fi62
                                _ = v63
                            }
                            _ = fi60
                            _ = v61
                        }
                        _ = v59
                    }
                case -1:
                    var v64 string = v5
                    _ = v64
                default:
                    var v65 int = 1
                    _ = v65
                }
            }
        }
    }
    fmt.Println(v3, v4, v3)
    v66 := GenS6{F0: 63, F1: 881, F2: 4294967295}
    for fi67 := 0; fi67 < 7; fi67++ {
        v3 += 2
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi67
    }
    for fi68 := 0; fi68 < 7; fi68++ {
        for fi69 := 0; fi69 < 19; fi69++ {
            switch v5 {
            case "hello":
                if (952 == -32768) {
                    switch (fi69 | fi68) {
                    case 3:
                        _ = len(v1)
                    case 242:
                        var v70 int = (fi69 << ((v3) & 63))
                        _ = v70
                    }
                }
            case "":
                fmt.Println(v4)
                var v71 int16 = 32767
                _ = v71
            case "a":
                if true {
                    v72 := GenS6{F0: 65535, F1: 0, F2: (fi68 * (fi69 - fi69))}
                    v73 := GenS7{F0: 42}
                    _ = v72
                    _ = v73
                } else {
                    if false {
                        var v74 int8 = 1
                        switch (fi68 / ((int(v3)) | 1)) {
                        case -538:
                            fmt.Println(v4)
                        default:
                            fmt.Println(v4)
                        }
                        _ = v74
                    }
                }
                v75 := GenS7{F0: -1}
                var v76 int8 = 100
                var v77 uint64 = 262
                if (!false) {
                    var v78 uint32 = 4294967295
                    var v79 uint8 = 0
                    switch "\n\t" {
                    case "a":
                        _ = len(v1)
                    case "foo bar":
                        switch "\n\t" {
                        case "foo bar":
                            v77 = v77
                        case "\n\t":
                            fmt.Println(v3, fi68)
                        case "hello":
                            fmt.Println(fi68)
                        }
                    default:
                        var v80 uint8 = (v79 ^ uint8(v3))
                        v78 -= v78
                        _ = v80
                    }
                    _ = v78
                    _ = v79
                } else {
                    if false {
                        if (2147483647 > 1) {
                            fmt.Println(fi69)
                            v81 := GenS7{F0: 42}
                            _ = v81
                        }
                    } else {
                        v77 = v77
                        v2[0] = fi69
                        var v82 uint64 = 18446744073709551615
                        if (!((true && true) && false)) {
                            var v83 bool = (!true)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch (fi69 | 313) {
                            case 2147483647:
                                if false {
                                    var v84 int64 = 42
                                    _ = v84
                                }
                            case 127:
                                v85 := GenS4{F0: ""}
                                _ = v85
                            default:
                                fmt.Println(fi68, fi68, v77)
                            }
                            _ = v83
                        }
                        v86 := GenS8{F0: int8(fi68)}
                        switch (fi68 & (fi68 << (((^(v3))) & 63))) {
                        case 32767:
                            var v87 uint32 = 2147483647
                            _ = v87
                        case 32768:
                            v88 := GenS1{F0: 65535, F1: 9223372036854775807}
                            _ = v88
                        default:
                            if false {
                                if false {
                                    v82 = uint64(v3)
                                    v76 -= v76
                                    var v89 int8 = v76
                                    v5 = ((("foo bar" + v5) + ("\n\t" + (v5 + v5))) + "hello")
                                    switch (("foo bar" + v5) + "a") {
                                    case "a":
                                        v90 := GenS9{F0: 0, F1: fi69}
                                        _ = v90
                                    case "":
                                        switch fi69 {
                                        case 4294967295:
                                            var v91 uint16 = 723
                                            _ = v91
                                        default:
                                            fmt.Println(v89, v82)
                                        }
                                    }
                                    _ = v89
                                } else {
                                    var v92 bool = (127 != 0)
                                    var v93 int16 = -1
                                    if (34 != 204) {
                                        var v94 uint32 = 65535
                                        _ = v94
                                    } else {
                                        v82 += v77
                                        v95 := GenS7{F0: v93}
                                        _ = v95
                                    }
                                    _ = v92
                                    _ = v93
                                }
                                v3 = (v3 + 746)
                            } else {
                                v77 += (v77 * (v82 | (^(v77))))
                                var v96 float64 = 0.0
                                _ = v96
                            }
                        }
                        _ = v82
                        _ = v86
                    }
                    switch 0 {
                    case 2147483647:
                        if (228 > 0) {
                            var v97 int16 = 42
                            _ = v97
                        } else {
                            if true {
                                var v98 float32 = (0.1 / 1.0)
                                switch v5 {
                                case "":
                                    if (true && ((fi68 ^ (fi68 - -9223372036854775808)) > (fi69 | fi69))) {
                                        var v99 uint = (v3 - (v3 - 1))
                                        if false {
                                            if (!true) {
                                                if len(v1) > 0 {
                                                    v1[0] = 682
                                                }
                                            } else {
                                                var v100 int32 = 2147483647
                                                var v101 bool = (255 != 53)
                                                var v102 uint8 = 212
                                                _ = v100
                                                _ = v101
                                                _ = v102
                                            }
                                            _ = len(v2)
                                        }
                                        if false {
                                            v103 := v2[uint32(v3)]
                                            if len(v1) > 0 {
                                                v1[0] = 787
                                            }
                                            fmt.Println(v77)
                                            fmt.Println(v103, v3)
                                            _ = len(v2)
                                            v98 = float32(3.14159)
                                            var v104 string = v5
                                            _ = v103
                                            _ = v104
                                        }
                                        var v105 int32 = 0
                                        _ = v105
                                        _ = v99
                                    } else {
                                        v106 := v2[318]
                                        _ = v106
                                    }
                                    fmt.Println(fi69, v4)
                                    if (true || false) {
                                        delete(v2, 2147483648)
                                        switch 32767 {
                                        case 499:
                                            if ((false && (!((93 > -122) || ((-237 >= 9223372036854775807) && (fi68 == 127))))) || ((v3 + genFn0((v3 < v3))) > 799)) {
                                                var v107 int64 = -501
                                                v108 := GenS8{F0: 0}
                                                v76 = v76
                                                _ = v107
                                                _ = v108
                                            }
                                        case -2147483648:
                                            v109 := GenS2{F0: v77, F1: 252}
                                            var v110 uint8 = 16
                                            var v111 uint32 = 4294967295
                                            _ = v109
                                            _ = v110
                                            _ = v111
                                        }
                                    }
                                case "foo bar":
                                    switch fi69 {
                                    case 32767:
                                        var v112 int = fi69
                                        _ = v112
                                    case -9223372036854775808:
                                        v1 = append(v1, 65535)
                                    }
                                case "\n\t":
                                    var v113 float64 = -2.25
                                    _ = v113
                                }
                                _ = v98
                            } else {
                                var v114 int = (fi69 & 127)
                                _ = v114
                            }
                            fmt.Println(fi69, v5, v4)
                            var v115 uint8 = 124
                            _ = v115
                        }
                    }
                    switch (v5 + "hello") {
                    case "":
                        _ = len(v2)
                    case "foo bar":
                        delete(v2, uint32(v77))
                    case "hello":
                        var v116 float32 = 0.1
                        _ = v116
                    }
                }
                _ = v75
                _ = v76
                _ = v77
            default:
                var v117 uint8 = 0
                _ = v117
            }
            v4 = v4
            v118 := GenS4{F0: v5}
            v5 = v5
            fmt.Println(v5, fi69)
            _ = fi69
            _ = v118
        }
        _ = fi68
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v66
}
