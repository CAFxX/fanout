package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint; F2 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 int; F2 uint16 }
type GenS8 struct { F0 uint8; F1 int64; F2 uint }
type GenS9 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS10 struct { F0 int; F1 uint64; F2 string }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS13 struct { F0 bool; F1 uint8 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int8; F1 uint8 }

func main() {
    var v1 string = "foo bar"
    var v2 uint16 = 42
    var v3 float64 = float64(1e10)
    if true {
        v4 := GenS1{F0: 0}
        switch v1 {
        case "a":
            fmt.Println(v2)
        case "\n\t":
            var v5 float32 = 0.0
            _ = v5
        case "foo bar":
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 float32 = 1e10
                fmt.Println(v2, v3, v3)
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 string = "foo bar"
                    if (fi6 != fi8) {
                        fmt.Println(v1, fi6)
                        fmt.Println(fi8)
                    } else {
                        switch (v9 + v9) {
                        case "":
                            if (1 > 1) {
                                var v10 int32 = 0
                                if (!true) {
                                    var v11 uint = 1
                                    switch (fi6 << ((v11) & 63)) {
                                    case -837:
                                        v11 = v11
                                    default:
                                        var v12 int32 = v10
                                        _ = v12
                                    }
                                    v7 -= (((0.1 / float32(3.14159)) * (v7 * 0.0)) / -2.25)
                                    v13 := []int{fi8, (fi8 / ((1) | 1)), (fi6 ^ (fi8 | 2147483647))}
                                    v14 := GenS2{F0: 0}
                                    if false {
                                        var v15 uint64 = 4294967295
                                        v9 = "foo bar"
                                        _ = v15
                                    }
                                    v7 = (v7 / 0.1)
                                    _ = v11
                                    _ = v13
                                    _ = v14
                                }
                                switch 412 {
                                case 1:
                                    var v16 int = fi6
                                    _ = v16
                                case -2147483648:
                                    switch v9 {
                                    case "":
                                        v17 := GenS3{F0: 982, F1: 255, F2: 2147483647}
                                        _ = v17
                                    default:
                                        v18 := GenS3{F0: 1000, F1: 255, F2: 1}
                                        _ = v18
                                    }
                                }
                                _ = v10
                            } else {
                                v19 := []int{}
                                fmt.Println(v2, fi8)
                                _ = v19
                            }
                        default:
                            var v20 uint64 = 42
                            _ = v20
                        }
                        if (!false) {
                            v2 -= v2
                        } else {
                            var v21 uint = 0
                            var v22 float32 = v7
                            v23 := GenS4{F0: 1}
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    }
                    switch fi8 {
                    case 127:
                        v9 = (v9 + v9)
                    case -486:
                        switch ("hello" + "\n\t") {
                        case "a":
                            var v24 bool = (9223372036854775807 >= 797)
                            _ = v24
                        case "hello":
                            var v25 string = ("a" + v1)
                            _ = v25
                        }
                    case 978:
                        if (!(162 > 0)) {
                            switch (fi6 << ((2) & 63)) {
                            case -1:
                                v26 := GenS1{F0: 65535}
                                _ = v26
                            case 564:
                                fmt.Println(fi6)
                            case -2147483648:
                                var v27 int32 = 2147483647
                                _ = v27
                            }
                            if (false || true) {
                                v28 := []int{fi6, -1}
                                if true {
                                    var v29 int8 = int8(fi8)
                                    if ((false && true) && true) {
                                        if ((-32768 == 572) && (!true)) {
                                            if false {
                                                var v30 int64 = 9223372036854775807
                                                v31 := GenS5{F0: 42}
                                                fmt.Println(fi6, v1)
                                                _ = v30
                                                _ = v31
                                            } else {
                                                fmt.Println(v7, v2, fi8)
                                            }
                                        } else {
                                            var v32 string = "a"
                                            v33 := GenS2{F0: 424}
                                            v34 := GenS6{F0: v2}
                                            var v35 int32 = 0
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                        if len(v28) > 0 {
                                            v28[0] = 458
                                        }
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                    } else {
                                        v28 = append(v28, 128)
                                    }
                                    var v36 uint16 = 209
                                    if false {
                                        var v37 float64 = v3
                                        if len(v28) > 0 {
                                            v28[0] = (fi6 | -1)
                                        }
                                        _ = v37
                                    } else {
                                        fmt.Println(v1)
                                        if false {
                                            v28 = append(v28, (fi6 % ((fi8) | 1)))
                                        }
                                        v7 += 0.1
                                    }
                                    _ = v29
                                    _ = v36
                                } else {
                                    var v38 float64 = (v3 - v3)
                                    var v39 uint32 = 2147483647
                                    switch (v1 + "\n\t") {
                                    case "\n\t":
                                        var v40 int16 = 653
                                        _ = v40
                                    case "foo bar":
                                        fmt.Println(v7, v39, fi8)
                                    }
                                    _ = v38
                                    _ = v39
                                }
                                if false {
                                    var v41 float64 = (v3 * -2.25)
                                    var v42 string = v9
                                    _ = v41
                                    _ = v42
                                }
                                _ = v28
                            } else {
                                v43 := GenS5{F0: 2}
                                if false {
                                    v3 = 3.14159
                                    if ((!(9223372036854775807 <= 1)) && true) {
                                        var v44 int8 = -128
                                        var v45 int64 = 163
                                        _ = v44
                                        _ = v45
                                    } else {
                                        v3 = (v3 + v3)
                                        v46 := GenS7{F0: 2, F1: (fi8 << ((255) & 63)), F2: v2}
                                        v1 = v1
                                        _ = v46
                                    }
                                } else {
                                    v1 = (v9 + "a")
                                    v47 := []int{fi8, (fi8 >> ((1) & 63)), (fi8 ^ 32767), fi8}
                                    switch (fi6 << ((385) & 63)) {
                                    case -715:
                                        v3 += -2.25
                                    case 2147483647:
                                        switch v9 {
                                        case "":
                                            v7 = ((((v7 * (0.1 / -2.25)) * -2.25) + (v7 * float32(float64(v7)))) + v7)
                                        }
                                    case -1:
                                        if true {
                                            if true {
                                                var v48 string = v1
                                                v49 := GenS5{F0: 42}
                                                v50 := GenS6{F0: (v2 << ((0) & 63))}
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                            } else {
                                                switch (fi8 / ((42) | 1)) {
                                                case -37:
                                                    switch ("a" + v9) {
                                                    case "foo bar":
                                                        v7 += (0.1 + 0.1)
                                                    case "hello":
                                                        var v51 float64 = v3
                                                        _ = v51
                                                    case "\n\t":
                                                        v52 := GenS5{F0: 18446744073709551615}
                                                        var v53 uint16 = (v2 - (v2 + v2))
                                                        v53 -= v53
                                                        _ = v52
                                                        _ = v53
                                                    default:
                                                        var v54 int64 = 768
                                                        _ = v54
                                                    }
                                                case 2147483647:
                                                    var v55 uint8 = 42
                                                    _ = v55
                                                }
                                                v56 := GenS4{F0: 42}
                                                if (fi6 <= fi6) {
                                                    if true {
                                                        if false {
                                                            var v57 float64 = (v3 * (v3 - float64(v7)))
                                                            var v58 uint64 = 42
                                                            var v59 int = 0
                                                            v1 = v9
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            v1 = v1
                                                            _ = len(v47)
                                                            v60 := GenS8{F0: 116, F1: -928, F2: 2}
                                                            _ = v60
                                                        }
                                                        if false {
                                                            var v61 uint16 = v2
                                                            v7 += (-2.25 / v7)
                                                            _ = v61
                                                        }
                                                        v2 = (v2 << ((42) & 63))
                                                        switch 1 {
                                                        case -770:
                                                            v1 = ("" + "hello")
                                                        }
                                                        var v62 uint64 = 4294967295
                                                        var v63 float32 = v7
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        if len(v47) >= 2 {
                                                            v47 = v47[1:len(v47)]
                                                        }
                                                        v7 += 1.5
                                                        v1 = v9
                                                        var v64 float32 = v7
                                                        _ = v64
                                                    }
                                                    if (18446744073709551615 < 0) {
                                                        if true {
                                                            _ = len(v47)
                                                        }
                                                        fmt.Println(v7, v1)
                                                        var v65 uint = 0
                                                        _ = v65
                                                    } else {
                                                        fmt.Println(v1)
                                                        if (true || false) {
                                                            if true {
                                                                _ = len(v47)
                                                            }
                                                        } else {
                                                            switch v1 {
                                                            case "a":
                                                                var v66 bool = false
                                                                _ = v66
                                                            case "":
                                                                var v67 bool = (v1 == v9)
                                                                _ = v67
                                                            case "\n\t":
                                                                v3 += (v3 * v3)
                                                            }
                                                            v68 := GenS9{F0: 42, F1: -32768, F2: (^(v2))}
                                                            _ = v68
                                                        }
                                                        var v69 uint16 = (v2 - 913)
                                                        _ = v69
                                                    }
                                                    fmt.Println(fi8, fi8, v1)
                                                } else {
                                                    fmt.Println(fi6, v1)
                                                    v2 += 42
                                                    _ = len(v47)
                                                    fmt.Println(v7, v7, v1)
                                                }
                                                _ = v56
                                            }
                                            v47 = append(v47, fi6)
                                            var v70 uint = 255
                                            _ = v70
                                        }
                                    }
                                    _ = v47
                                }
                                _ = v43
                            }
                        }
                    }
                    _ = fi8
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
        default:
            if false {
                for fi71 := 0; fi71 < 12; fi71++ {
                    var v72 int32 = 0
                    _ = fi71
                    _ = v72
                }
                v1 = v1
            } else {
                v73 := GenS8{F0: 112, F1: 948, F2: 229}
                var v74 int64 = -376
                if (1 != -66) {
                    var v75 int32 = -1
                    _ = v75
                } else {
                    for fi76 := 0; fi76 < 11; fi76++ {
                        switch (fi76 * (-(fi76))) {
                        case 493:
                            if true {
                                var v77 int32 = 42
                                for fi78 := 0; fi78 < 17; fi78++ {
                                    v3 -= 1.5
                                    _ = fi78
                                }
                                v1 = v1
                                _ = v77
                            } else {
                                if (!(1 < 398)) {
                                    for fi79 := 0; fi79 < 15; fi79++ {
                                        v80 := []int{(fi76 | fi76)}
                                        _ = fi79
                                        _ = v80
                                    }
                                    var v81 int8 = 0
                                    _ = v81
                                }
                                fmt.Println(v74, v3)
                            }
                        case 0:
                            fmt.Println(fi76)
                        case 1:
                            var v82 float64 = ((v3 / 1.0) - float64(((((0.0 / 1.5) - 0.0) - ((1.5 / 1e10) / ((float32(1.5) + (1e10 / (1e10 / 0.1))) / (1e10 * 0.1)))) - 1e10)))
                            _ = v82
                        default:
                            var v83 uint32 = 1
                            _ = v83
                        }
                        _ = fi76
                    }
                    v84 := GenS10{F0: 42, F1: 115, F2: ((v1 + "foo bar") + "hello")}
                    _ = v84
                }
                _ = v73
                _ = v74
            }
        }
        var v85 uint16 = 42
        _ = v4
        _ = v85
    } else {
        fmt.Println(v3, v2, v1)
    }
    v1 = (("foo bar" + (v1 + v1)) + v1)
    for fi86 := 0; fi86 < 9; fi86++ {
        for fi87 := 0; fi87 < 11; fi87++ {
            switch 1 {
            case -2147483648:
                fmt.Println(fi87, v2)
            case 1:
                v3 -= 0.0
            default:
                if (((3.14159 / -2.25) / (0.1 / 1.0)) <= 1.5) {
                    var v88 int32 = int32(fi86)
                    _ = v88
                }
            }
            var v89 int64 = 42
            _ = fi87
            _ = v89
        }
        var v90 uint8 = 128
        v91 := []int{(fi86 / (((fi86 << ((988) & 63))) | 1)), -833, fi86}
        _ = fi86
        _ = v90
        _ = v91
    }
    v3 += (3.14159 * v3)
    var v92 int8 = -13
    switch "foo bar" {
    case "a":
        var v93 int16 = -1
        _ = v93
    case "\n\t":
        v1 = (v1 + v1)
    }
    switch 42 {
    case 42:
        if true {
            v3 += v3
            v2 = 255
        }
    case 2147483647:
        var v94 int64 = 42
        _ = v94
    case 128:
        for fi95 := 0; fi95 < 6; fi95++ {
            v96 := GenS3{F0: 42, F1: 255, F2: 1}
            _ = fi95
            _ = v96
        }
    default:
        var v97 uint64 = 4294967295
        _ = v97
    }
    v98 := []int{1, 2147483647, 127, 42}
    if len(v98) > 0 {
        v98[0] = 1
    }
    fmt.Println(v1, v3)
    for fi99 := 0; fi99 < 9; fi99++ {
        if true {
            var v100 float32 = 1e10
            _ = v100
        } else {
            if false {
                v92 = (v92 - v92)
                v101 := GenS11{F0: 389}
                v1 = "\n\t"
                _ = v101
            }
        }
        if (!false) {
            if false {
                if len(v98) >= 2 {
                    v98 = v98[1:len(v98)]
                }
                v2 = 159
            } else {
                v102 := GenS5{F0: 278}
                v1 = v1
                if (fi99 != fi99) {
                    switch v1 {
                    case "a":
                        var v103 int8 = 42
                        _ = v103
                    default:
                        var v104 int8 = v92
                        _ = v104
                    }
                } else {
                    v105 := GenS12{F0: 0, F1: 18446744073709551615, F2: v92}
                    for fi106 := 0; fi106 < 13; fi106++ {
                        v3 = (v3 * v3)
                        _ = fi106
                    }
                    var v107 uint8 = 0
                    _ = v105
                    _ = v107
                }
                _ = v102
            }
        }
        for fi108 := 0; fi108 < 7; fi108++ {
            v3 = -2.25
            v109 := GenS13{F0: (187 < 2), F1: 0}
            v110 := GenS14{F0: (v1 + v1)}
            _ = fi108
            _ = v109
            _ = v110
        }
        _ = fi99
    }
    v111 := GenS15{F0: -128, F1: 207}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v92:", v92)
    fmt.Println("slicelen v98:", len(v98), v98)
    _ = v1
    _ = v111
    _ = v2
    _ = v3
    _ = v92
    _ = v98
}
