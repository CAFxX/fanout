package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int8; F1 int16; F2 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS7 struct { F0 uint; F1 int; F2 uint }
type GenS8 struct { F0 uint32; F1 uint; F2 int64 }
type GenS9 struct { F0 uint8; F1 int64 }
type GenS10 struct { F0 bool; F1 uint64 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint; F1 string; F2 uint64 }
func genFn0(p0 bool, p1 int, p2 int16) int64 {
    p0 = p0
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int = fi5
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p1, p2)
    v7 := []int{128, (p1 | (p1 % ((32767) | 1)))}
    p2 = (p2 - p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -263
}
func genFn1(p0 int64, p1 uint32, p2 uint16) int16 {
    switch ("a" + "a") {
    case "a":
        p2 += p2
    }
    if false {
        var v8 int16 = 1000
        var v9 uint16 = p2
        _ = v8
        _ = v9
    } else {
        if (!true) {
            var v10 uint64 = 4294967295
            for fi11 := 0; fi11 < 2; fi11++ {
                switch -9223372036854775808 {
                case 1:
                    for fi12 := 0; fi12 < 5; fi12++ {
                        var v13 int32 = 313
                        _ = fi12
                        _ = v13
                    }
                case -2147483648:
                    v14 := GenS1{F0: genFn0((((3.14159 / 1.0) / -2.25) < 0.1), (fi11 << ((255) & 63)), -288)}
                    _ = v14
                case -155:
                    v15 := GenS1{F0: (p0 ^ (p0 * (p0 / ((0) | 1))))}
                    _ = v15
                }
                v16 := GenS1{F0: p0}
                _ = fi11
                _ = v16
            }
            _ = v10
        }
        switch 32767 {
        case -1:
            v17 := []int{9223372036854775807}
            _ = v17
        default:
            for fi18 := 0; fi18 < 4; fi18++ {
                p1 += p1
                var v19 uint32 = 2
                p0 += (p0 * (p0 * genFn0(false, fi18, 42)))
                _ = fi18
                _ = v19
            }
        }
        v20 := GenS2{F0: p2, F1: 0, F2: (p0 | genFn0((genFn0((!((0.0 + ((((3.14159 / 1.5) - 1.5) + 3.14159) * 1e10)) != ((1e10 * 0.1) * float32(0.0)))), 1, 1000) <= p0), int(p2), 0))}
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint64 = 0
    var v2 float32 = 0.1
    var v3 bool = ((false || (!false)) && (435 == 65535))
    var v4 string = "foo bar"
    v1 -= 4294967295
    var v21 uint32 = 1
    v3 = v3
    var v22 int16 = 32767
    v3 = false
    v23 := []int{32767, -2147483648}
    var v24 int = 0
    for fi25 := 0; fi25 < 14; fi25++ {
        v4 = v4
        for fi26 := 0; fi26 < 18; fi26++ {
            if true {
                var v27 uint32 = v21
                v22 += (^(v22))
                switch (fi26 << ((718) & 63)) {
                case -1:
                    fmt.Println(v24, fi26, fi26)
                    var v28 uint64 = 538
                    _ = v28
                case 407:
                    var v29 int8 = 19
                    _ = v29
                case 32768:
                    switch (fi26 & 32768) {
                    case 127:
                        v30 := GenS1{F0: 42}
                        _ = v30
                    case 0:
                        var v31 int64 = 9223372036854775807
                        _ = v31
                    }
                }
                _ = v27
            }
            v32 := GenS3{F0: (v1 - (v1 << ((1) & 63))), F1: 2}
            switch v4 {
            case "":
                fmt.Println(v2)
                var v33 uint = 2
                _ = v33
            case "foo bar":
                if v3 {
                    v34 := GenS4{F0: 39, F1: v22, F2: v21}
                    fmt.Println(fi26, v21, v21)
                    _ = v34
                }
            case "\n\t":
                v3 = (127 <= 242)
            }
            v21 = v21
            var v35 float32 = float32((1.5 / (((1e10 * 3.14159) + 0.0) - (-2.25 + -2.25))))
            _ = fi26
            _ = v32
            _ = v35
        }
        switch ("" + "a") {
        case "foo bar":
            v36 := GenS4{F0: 127, F1: int16(v21), F2: 2147483647}
            _ = v36
        case "":
            switch v4 {
            case "a":
                var v37 uint16 = 42
                _ = v37
            case "foo bar":
                switch "foo bar" {
                case "\n\t":
                    v22 -= 32767
                }
            case "\n\t":
                var v38 float32 = 3.14159
                _ = v38
            default:
                var v39 uint32 = 2
                _ = v39
            }
        case "hello":
            if true {
                var v40 uint16 = 1
                v41 := GenS5{F0: 0}
                _ = v40
                _ = v41
            } else {
                if (v3 && false) {
                    if (0 == (v22 + (v22 + (v22 - v22)))) {
                        var v42 int = -851
                        _ = v42
                    } else {
                        var v43 int = (v24 / ((127) | 1))
                        if v3 {
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            if (!v3) {
                                if (v3 || (-586 > 42)) {
                                    switch v24 {
                                    case 1:
                                        fmt.Println(v1, v2)
                                    }
                                    v24 = int(v22)
                                    switch ((v4 + v4) + v4) {
                                    case "hello":
                                        if (238 < 208) {
                                            fmt.Println(v24)
                                            var v44 uint32 = 0
                                            for fi45 := 0; fi45 < 14; fi45++ {
                                                v46 := GenS6{F0: (1.5 == v2), F1: 42, F2: 0}
                                                _ = fi45
                                                _ = v46
                                            }
                                            _ = v44
                                        } else {
                                            v3 = (v2 >= v2)
                                        }
                                    case "\n\t":
                                        v3 = true
                                    case "a":
                                        v47 := GenS5{F0: -2147483648}
                                        _ = v47
                                    }
                                }
                            }
                        } else {
                            if len(v23) > 0 {
                                v23[0] = (^(v24))
                            }
                        }
                        _ = len(v23)
                        _ = v43
                    }
                    switch v4 {
                    case "foo bar":
                        if false {
                            fmt.Println(v21, fi25, v4)
                            v48 := GenS7{F0: 4294967295, F1: (v24 - 127), F2: 2}
                            _ = v48
                        }
                    }
                } else {
                    if v3 {
                        for fi49 := 0; fi49 < 8; fi49++ {
                            switch -776 {
                            case 1:
                                var v50 string = (v4 + "")
                                _ = v50
                            case 9223372036854775807:
                                var v51 uint32 = (v21 + (v21 << ((24) & 63)))
                                _ = v51
                            default:
                                v22 = -1
                            }
                            _ = fi49
                        }
                        var v52 uint8 = 1
                        _ = v52
                    } else {
                        for fi53 := 0; fi53 < 5; fi53++ {
                            if v3 {
                                switch 9223372036854775807 {
                                case -537:
                                    fmt.Println(v2, v24)
                                case -207:
                                    var v54 uint = 0
                                    _ = v54
                                default:
                                    v23 = append(v23, 770)
                                }
                                v55 := GenS6{F0: (1000 != (v22 - (v22 + -832))), F1: 126, F2: (v1 | 4294967295)}
                                if ((v3 && v3) || ((2 > 1) || (!false))) {
                                    if (!((4294967295 != 2) && (100 > 42))) {
                                        var v56 uint64 = (^(v1))
                                        var v57 int8 = 42
                                        if (307 <= uint(v21)) {
                                            var v58 uint16 = 268
                                            v24 = -1
                                            v59 := GenS5{F0: -793}
                                            _ = v58
                                            _ = v59
                                        }
                                        _ = v56
                                        _ = v57
                                    }
                                    if len(v23) > 0 {
                                        v23[0] = v24
                                    }
                                } else {
                                    switch -361 {
                                    case -1:
                                        if (128 != 98) {
                                            if (v22 == 310) {
                                                fmt.Println(v2)
                                                if (255 == 1) {
                                                    var v60 string = "a"
                                                    var v61 int32 = 2147483647
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    var v62 float32 = 1.5
                                                    v24 += v24
                                                    v63 := GenS1{F0: -866}
                                                    _ = v62
                                                    _ = v63
                                                }
                                                if false {
                                                    if false {
                                                        var v64 bool = (85 != 101)
                                                        _ = v64
                                                    }
                                                    var v65 int64 = 42
                                                    _ = v65
                                                } else {
                                                    switch (fi25 << ((4294967295) & 63)) {
                                                    case 32767:
                                                        var v66 string = "a"
                                                        _ = v66
                                                    }
                                                }
                                            } else {
                                                v2 = v2
                                                v21 = (v21 / (((v21 & (v21 % ((v21) | 1)))) | 1))
                                                fmt.Println(v24, fi25)
                                            }
                                            var v67 string = (v4 + v4)
                                            _ = v67
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(v21)
                                    default:
                                        _ = len(v23)
                                    }
                                    var v68 uint32 = uint32(v1)
                                    var v69 int8 = -128
                                    _ = v68
                                    _ = v69
                                }
                                _ = v55
                            }
                            _ = fi53
                        }
                    }
                    if v3 {
                        for fi70 := 0; fi70 < 14; fi70++ {
                            v4 = (v4 + "foo bar")
                            v24 -= fi25
                            var v71 string = v4
                            _ = fi70
                            _ = v71
                        }
                        var v72 int = 4294967295
                        _ = v72
                    } else {
                        if ((v24 >> ((42) & 63)) < v24) {
                            for fi73 := 0; fi73 < 6; fi73++ {
                                v74 := GenS6{F0: (65535 != 65535), F1: 210, F2: 303}
                                _ = fi73
                                _ = v74
                            }
                            var v75 int = (v24 % (((v24 % (((v24 % ((v24) | 1))) | 1))) | 1))
                            _ = v75
                        } else {
                            v76 := GenS8{F0: 42, F1: 192, F2: -1}
                            if len(v23) > 0 {
                                v23[0] = (fi25 ^ -76)
                            }
                            _ = v76
                        }
                        v21 = 2
                        if (!v3) {
                            v24 += 127
                        }
                    }
                }
                for fi77 := 0; fi77 < 2; fi77++ {
                    fmt.Println(v21)
                    if v3 {
                        _ = len(v23)
                        if v3 {
                            v23 = append(v23, (fi77 - fi77))
                            switch 4294967295 {
                            case 319:
                                if (((-2.25 - (3.14159 + 3.14159)) / 3.14159) < (3.14159 - 0.1)) {
                                    var v78 float32 = v2
                                    _ = v78
                                } else {
                                    v1 += (v1 ^ v1)
                                    fmt.Println(fi77, fi25)
                                }
                            }
                            switch v4 {
                            case "hello":
                                if (("\n\t" < v4) || true) {
                                    fmt.Println(v4, v24, v24)
                                    v79 := GenS6{F0: v3, F1: 1, F2: (v1 ^ (v1 % ((v1) | 1)))}
                                    var v80 bool = (125 == 2)
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    var v81 bool = (255 > uint8(v21))
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                } else {
                                    var v82 string = ((v4 + "hello") + "hello")
                                    var v83 uint32 = (v21 >> ((0) & 63))
                                    var v84 float32 = v2
                                    _ = v82
                                    _ = v83
                                    _ = v84
                                }
                            case "":
                                v1 -= 42
                            default:
                                var v85 uint = 0
                                _ = v85
                            }
                        } else {
                            v86 := GenS9{F0: 42, F1: 9223372036854775807}
                            _ = v86
                        }
                        var v87 uint8 = uint8(v1)
                        _ = v87
                    } else {
                        if len(v23) > 0 {
                            v23[0] = -1
                        }
                    }
                    v88 := GenS2{F0: 32767, F1: 32767, F2: 0}
                    _ = fi77
                    _ = v88
                }
                v1 -= (v1 >> ((4294967295) & 63))
            }
        default:
            var v89 float32 = v2
            _ = v89
        }
        _ = fi25
    }
    for fi90 := 0; fi90 < 16; fi90++ {
        _ = len(v23)
        _ = fi90
    }
    v3 = v3
    var v91 int8 = 100
    if (true && false) {
        if len(v23) > 0 {
            v23[0] = 78
        }
    } else {
        for fi92 := 0; fi92 < 1; fi92++ {
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            _ = fi92
        }
        v1 = 18446744073709551615
        if (!(!false)) {
            if (v3 && v3) {
                if v3 {
                    v23 = append(v23, (v24 >> ((510) & 63)))
                } else {
                    var v93 uint64 = 384
                    if (32767 != int(v91)) {
                        var v94 int16 = genFn1(1, 2147483648, 381)
                        _ = v94
                    } else {
                        v22 = v22
                        v95 := GenS10{F0: (v2 <= (v2 / v2)), F1: 4294967295}
                        if ((!(!((v93 + (v93 + v1)) >= v1))) && v3) {
                            v2 -= (v2 + (v2 * (0.1 - v2)))
                        }
                        _ = v95
                    }
                    _ = v93
                }
                if ((3.14159 - ((1.5 / 1.5) / ((1e10 * (0.1 + 0.0)) + -2.25))) >= -2.25) {
                    v21 += v21
                    if true {
                        v1 += v1
                        switch (v24 - (v24 - int(v1))) {
                        case 240:
                            for fi96 := 0; fi96 < 5; fi96++ {
                                if v3 {
                                    var v97 uint64 = 42
                                    _ = v97
                                } else {
                                    if v3 {
                                        v98 := GenS5{F0: 1}
                                        var v99 int64 = 42
                                        _ = v98
                                        _ = v99
                                    }
                                    for fi100 := 0; fi100 < 5; fi100++ {
                                        v101 := GenS7{F0: 4294967295, F1: (fi100 << ((895) & 63)), F2: 981}
                                        _ = fi100
                                        _ = v101
                                    }
                                    var v102 int64 = -459
                                    _ = v102
                                }
                                _ = fi96
                            }
                        }
                    } else {
                        switch int(v22) {
                        case 2147483647:
                            v103 := GenS11{F0: ""}
                            _ = v103
                        case 42:
                            var v104 bool = ((v3 && v3) || v3)
                            _ = v104
                        }
                        var v105 uint64 = uint64(v21)
                        switch (v24 & v24) {
                        case -451:
                            var v106 float64 = 1.5
                            _ = v106
                        case 32767:
                            var v107 uint = 2
                            _ = v107
                        case -438:
                            v108 := GenS12{F0: 692}
                            _ = v108
                        default:
                            v109 := GenS13{F0: (v91 >> ((255) & 63)), F1: (^(v21)), F2: v105}
                            _ = v109
                        }
                        _ = v105
                    }
                    var v110 int64 = genFn0((true || (1 != (v91 / ((v91) | 1)))), 843, (v22 % (((v22 >> ((42) & 63))) | 1)))
                    _ = v110
                } else {
                    var v111 float32 = v2
                    v112 := GenS14{F0: (2 > 0)}
                    v24 = 340
                    _ = v111
                    _ = v112
                }
            } else {
                v113 := GenS15{F0: 42, F1: v4, F2: 2}
                var v114 uint8 = 42
                v115 := GenS15{F0: 333, F1: (v4 + "foo bar"), F2: v1}
                _ = v113
                _ = v114
                _ = v115
            }
            for fi116 := 0; fi116 < 10; fi116++ {
                var v117 string = v4
                _ = fi116
                _ = v117
            }
        } else {
            var v118 bool = v3
            v118 = v118
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            _ = v118
        }
    }
    var v119 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v119
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v91
}
