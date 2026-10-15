package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int64; F2 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16; F1 uint }
type GenS9 struct { F0 int8; F1 int8 }
type GenS10 struct { F0 string; F1 uint32; F2 int8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS13 struct { F0 bool; F1 int8; F2 uint16 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int8; F1 uint32; F2 uint }

func main() {
    var v1 uint64 = 0
    var v2 uint = 255
    v3 := make(map[int]int)
    var v4 uint64 = (^(v1))
    var v5 uint64 = v4
    var v6 int16 = 42
    v7 := []int{32767, -419, 2147483647}
    v8 := GenS1{F0: 2}
    var v9 uint8 = uint8(v4)
    if len(v7) > 0 {
        v7[0] = 896
    }
    switch 168 {
    case 274:
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
    }
    switch 643 {
    case 32768:
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := v3[fi10]
            switch fi10 {
            case 9223372036854775807:
                for fi12 := 0; fi12 < 10; fi12++ {
                    fmt.Println(v4)
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    _ = fi12
                }
            case 346:
                switch "" {
                case "\n\t":
                    v5 -= 0
                case "":
                    var v13 int32 = 42
                    _ = v13
                case "hello":
                    var v14 string = "hello"
                    _ = v14
                }
            case 42:
                v2 = (^(v2))
            default:
                v1 = (v5 * v5)
            }
            _ = fi10
            _ = v11
        }
    case 64:
        _ = len(v7)
    case -920:
        var v15 int8 = -1
        _ = v15
    }
    var v16 float64 = 1.5
    var v17 int = 171
    for fi18 := 0; fi18 < 7; fi18++ {
        if (false || (!(!true))) {
            v6 -= v6
            v17 -= (fi18 + v17)
            var v19 float64 = float64((0.0 - 3.14159))
            _ = v19
        }
        switch "foo bar" {
        case "a":
            var v20 int64 = 1
            _ = v20
        case "\n\t":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        default:
            _ = len(v3)
        }
        for fi21 := 0; fi21 < 13; fi21++ {
            v22 := GenS2{F0: fi18}
            switch 32767 {
            case 32767:
                switch (v17 & (fi21 - (v17 ^ 128))) {
                case -2147483648:
                    v3[9223372036854775807] = (fi18 + (fi21 & 32768))
                case 9223372036854775807:
                    var v23 int16 = v6
                    _ = v23
                default:
                    v17 = (fi21 % ((fi21) | 1))
                    var v24 uint64 = (v5 * 565)
                    _ = v24
                }
            case -916:
                v9 = v9
            default:
                fmt.Println(v4, v5)
                var v25 float64 = ((v16 + -2.25) * v16)
                if true {
                    if (!true) {
                        var v26 uint32 = 2
                        if true {
                            fmt.Println(v2, fi18, fi18)
                            var v27 uint32 = 4294967295
                            _ = v27
                        } else {
                            fmt.Println(v4)
                            fmt.Println(v16)
                        }
                        _ = v26
                    } else {
                        var v28 int8 = -27
                        _ = v28
                    }
                }
                _ = v25
            }
            switch "hello" {
            case "foo bar":
                v17 += (-(fi18))
            case "":
                v29 := GenS1{F0: 222}
                _ = v29
            default:
                var v30 int16 = int16(v2)
                _ = v30
            }
            _ = fi21
            _ = v22
        }
        _ = fi18
    }
    v31 := GenS3{F0: (v17 + (v17 >> ((v2) & 63))), F1: -1, F2: -128}
    for fi32 := 0; fi32 < 4; fi32++ {
        if (!(fi32 <= v17)) {
            var v33 int64 = -556
            fmt.Println(fi32, v16, v2)
            v34 := GenS4{F0: 9223372036854775807}
            _ = v33
            _ = v34
        } else {
            if (!((true && (!true)) || false)) {
                if true {
                    if ((fi32 & (fi32 / (((^(v17))) | 1))) < 1) {
                        fmt.Println(v6, v16)
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        for fi35 := 0; fi35 < 18; fi35++ {
                            if (-251 != -1) {
                                if len(v7) > 0 {
                                    v7[0] = v17
                                }
                                v5 -= (v5 ^ v5)
                                var v36 uint = (v2 * (v2 * uint(v9)))
                                _ = v36
                            } else {
                                v37 := GenS3{F0: (^(fi32)), F1: 42, F2: int8(v9)}
                                if true {
                                    v5 = 42
                                }
                                v16 = 3.14159
                                switch "foo bar" {
                                case "hello":
                                    var v38 uint16 = 1
                                    _ = v38
                                case "":
                                    fmt.Println(fi32)
                                case "foo bar":
                                    _ = len(v7)
                                default:
                                    var v39 float32 = 0.1
                                    _ = v39
                                }
                                if ((v17 & fi32) > 32767) {
                                    var v40 uint32 = 341
                                    var v41 int = (fi32 - fi32)
                                    _ = v40
                                    _ = v41
                                } else {
                                    v3[(v17 ^ fi35)] = 1
                                    if (false || (-32768 > (v6 - v6))) {
                                        var v42 uint = (v2 << ((v2) & 63))
                                        var v43 bool = ((1e10 * v16) > 0.1)
                                        if v43 {
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            if ("hello" != "") {
                                                switch "a" {
                                                case "a":
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                }
                                                v44 := GenS2{F0: fi32}
                                                _ = v44
                                            } else {
                                                if false {
                                                    var v45 uint32 = 0
                                                    fmt.Println(v17, v2)
                                                    var v46 int8 = 42
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    if v43 {
                                                        v1 += 131
                                                        switch (("a" + "\n\t") + ("hello" + "foo bar")) {
                                                        case "":
                                                            var v47 int16 = 42
                                                            _ = v47
                                                        case "a":
                                                            v48 := GenS1{F0: 2}
                                                            _ = v48
                                                        case "\n\t":
                                                            var v49 int32 = 0
                                                            _ = v49
                                                        }
                                                        if true {
                                                            if true {
                                                                v5 = 42
                                                                var v50 uint16 = 32767
                                                                v1 -= 490
                                                                _ = v50
                                                            }
                                                        } else {
                                                            var v51 int = (fi35 ^ 32768)
                                                            _ = v51
                                                        }
                                                    }
                                                    fmt.Println(fi32, v9)
                                                }
                                                if false {
                                                    if (!(2147483648 == uint32(v9))) {
                                                        v52 := GenS3{F0: (^(fi35)), F1: 9223372036854775807, F2: 0}
                                                        v2 = v2
                                                        v3[(-(fi32))] = (fi32 * fi32)
                                                        v53 := GenS4{F0: fi32}
                                                        v54 := GenS1{F0: 0}
                                                        v55 := GenS2{F0: -202}
                                                        switch 127 {
                                                        case -1:
                                                            if (4294967295 >= 42) {
                                                                if (42 != -128) {
                                                                    var v56 float32 = 3.14159
                                                                    v57 := GenS5{F0: 539, F1: (v9 >> ((v2) & 63))}
                                                                    v6 = (v6 >> ((42) & 63))
                                                                    v58 := GenS5{F0: 2, F1: v9}
                                                                    v16 += -2.25
                                                                    _ = v56
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    if v43 {
                                                                        v43 = (468 == 0)
                                                                    } else {
                                                                        fmt.Println(fi32, v5, v43)
                                                                        v4 = v5
                                                                    }
                                                                }
                                                                v5 = (^(v5))
                                                                var v59 uint16 = 59
                                                                _ = v59
                                                            } else {
                                                                switch (-(fi35)) {
                                                                case 9223372036854775807:
                                                                    v60 := GenS6{F0: 37, F1: (v9 - v9)}
                                                                    _ = v60
                                                                case 2147483647:
                                                                    v61 := GenS7{F0: (v2 >> (((v42 << (((v42 | v42)) & 63))) & 63))}
                                                                    _ = v61
                                                                }
                                                                if (((v2 - v42) != 706) || (-12 < -45)) {
                                                                    v6 -= v6
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        if (!v43) {
                                                                            var v62 int8 = -7
                                                                            var v63 int64 = 9223372036854775807
                                                                            v64 := GenS8{F0: 42, F1: v2}
                                                                            _ = v62
                                                                            _ = v63
                                                                            _ = v64
                                                                        } else {
                                                                            var v65 int = (v17 << (((v42 ^ v42)) & 63))
                                                                            _ = v65
                                                                        }
                                                                    case "foo bar":
                                                                        v66 := GenS3{F0: fi32, F1: -164, F2: 64}
                                                                        _ = v66
                                                                    case "a":
                                                                        v6 -= 1000
                                                                    default:
                                                                        v17 += -1
                                                                    }
                                                                }
                                                                v17 += -9223372036854775808
                                                            }
                                                        }
                                                        _ = v52
                                                        _ = v53
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                } else {
                                                    var v67 int = fi35
                                                    if v43 {
                                                        switch v17 {
                                                        case -502:
                                                            var v68 uint32 = uint32(v42)
                                                            v17 = fi32
                                                            if true {
                                                                switch "foo bar" {
                                                                case "":
                                                                    fmt.Println(v16)
                                                                case "hello":
                                                                    var v69 bool = (v4 <= v5)
                                                                    _ = v69
                                                                case "foo bar":
                                                                    v68 = v68
                                                                default:
                                                                    v70 := GenS3{F0: v17, F1: 42, F2: 1}
                                                                    _ = v70
                                                                }
                                                                v17 += fi35
                                                                var v71 int16 = 692
                                                                _ = v71
                                                            }
                                                            _ = v68
                                                        case -174:
                                                            var v72 uint16 = 967
                                                            _ = v72
                                                        }
                                                        fmt.Println(v1, v16, fi32)
                                                    } else {
                                                        var v73 float64 = (0.0 - 1.5)
                                                        _ = v73
                                                    }
                                                    v9 = v9
                                                    _ = v67
                                                }
                                                var v74 int8 = 0
                                                _ = v74
                                            }
                                        } else {
                                            v75 := GenS9{F0: -86, F1: -34}
                                            var v76 int8 = -67
                                            _ = v75
                                            _ = v76
                                        }
                                        _ = v42
                                        _ = v43
                                    }
                                    var v77 uint = (v2 | (v2 & v2))
                                    _ = v77
                                }
                                _ = v37
                            }
                            v17 -= (^(fi35))
                            _ = fi35
                        }
                    }
                    if true {
                        switch "foo bar" {
                        case "\n\t":
                            v4 -= 4294967295
                        case "foo bar":
                            for fi78 := 0; fi78 < 17; fi78++ {
                                fmt.Println(v17, v6)
                                _ = fi78
                            }
                        case "a":
                            v4 = v1
                        }
                        if false {
                            for fi79 := 0; fi79 < 4; fi79++ {
                                if false {
                                    if true {
                                        v17 += int(v6)
                                        v5 = (v4 / ((1) | 1))
                                    }
                                    switch 399 {
                                    case -839:
                                        var v80 uint = (v2 & v2)
                                        var v81 float32 = 1e10
                                        fmt.Println(v16)
                                        var v82 uint64 = 42
                                        switch fi79 {
                                        case -825:
                                            var v83 uint8 = (v9 + (v9 >> ((v2) & 63)))
                                            if ((!(true && (v83 != (v83 + (v9 ^ 240))))) && true) {
                                                v84 := GenS10{F0: ("a" + "foo bar"), F1: 0, F2: -128}
                                                _ = v84
                                            } else {
                                                switch "foo bar" {
                                                case "hello":
                                                    v16 += (v16 - 1.5)
                                                case "":
                                                    var v85 int = -1
                                                    _ = v85
                                                case "foo bar":
                                                    v1 = uint64(v80)
                                                default:
                                                    if false {
                                                        _ = len(v7)
                                                    }
                                                }
                                                var v86 bool = (uint8(v80) <= 2)
                                                _ = v86
                                            }
                                            _ = v83
                                        case 930:
                                            if (true && false) {
                                                v87 := GenS11{F0: ((783 > 2147483647) && false)}
                                                if true {
                                                    _ = len(v7)
                                                } else {
                                                    v88 := GenS6{F0: 70, F1: v9}
                                                    var v89 uint16 = 32767
                                                    switch 9223372036854775807 {
                                                    case 32767:
                                                        v1 -= v5
                                                    case -912:
                                                        v90 := GenS12{F0: 1, F1: 81, F2: (v4 * uint64(v9))}
                                                        v91 := GenS4{F0: (fi79 ^ (v17 + v17))}
                                                        if true {
                                                            v81 -= v81
                                                        }
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                    _ = v88
                                                    _ = v89
                                                }
                                                v16 += 1e10
                                                v17 += 42
                                                v92 := GenS2{F0: (^(fi32))}
                                                _ = v87
                                                _ = v92
                                            }
                                        default:
                                            switch 127 {
                                            case -327:
                                                fmt.Println(v81, v80, v4)
                                            default:
                                                v6 += v6
                                            }
                                        }
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    case 0:
                                        _ = len(v7)
                                        var v93 uint64 = v4
                                        v4 += uint64(v9)
                                        _ = v93
                                    case 32767:
                                        v94 := GenS11{F0: (!(1 >= (v1 >> ((0) & 63))))}
                                        _ = v94
                                    default:
                                        if true {
                                            v95 := GenS1{F0: 2147483648}
                                            var v96 int64 = 1
                                            _ = v95
                                            _ = v96
                                        } else {
                                            v5 += v1
                                            if len(v7) > 0 {
                                                v7[0] = fi79
                                            }
                                            v97 := GenS12{F0: v9, F1: 128, F2: v4}
                                            _ = v97
                                        }
                                    }
                                } else {
                                    v4 = (v5 & (v5 - v4))
                                }
                                switch "a" {
                                case "foo bar":
                                    if false {
                                        v98 := GenS1{F0: 2}
                                        _ = len(v7)
                                        _ = v98
                                    } else {
                                        v2 -= (^(v2))
                                        var v99 string = ""
                                        var v100 uint16 = 967
                                        _ = v100
                                        _ = v99
                                    }
                                    delete(v3, (fi79 & (fi79 | fi79)))
                                }
                                _ = fi79
                            }
                            var v101 int8 = 35
                            for fi102 := 0; fi102 < 15; fi102++ {
                                v2 += (v2 * v2)
                                if (0 <= -796) {
                                    if false {
                                        if false {
                                            v2 = (v2 << ((uint(v5)) & 63))
                                            var v103 uint32 = 65535
                                            _ = v103
                                        } else {
                                            var v104 int16 = 32767
                                            _ = v104
                                        }
                                    } else {
                                        fmt.Println(v2, v1, fi32)
                                        fmt.Println(v2, v5, v5)
                                    }
                                }
                                v105 := GenS5{F0: 658, F1: 1}
                                _ = fi102
                                _ = v105
                            }
                            _ = v101
                        } else {
                            if false {
                                var v106 uint32 = 4294967295
                                fmt.Println(v9)
                                _ = v106
                            } else {
                                delete(v3, 230)
                            }
                        }
                        fmt.Println(v6, v17)
                    } else {
                        var v107 int = int(v9)
                        _ = v107
                    }
                    for fi108 := 0; fi108 < 10; fi108++ {
                        switch 127 {
                        case 9223372036854775807:
                            v109 := GenS13{F0: (false || true), F1: -114, F2: 16}
                            _ = v109
                        }
                        var v110 uint32 = 1
                        if ((-128 == -82) || (false || false)) {
                            fmt.Println(v5, v6)
                            var v111 float64 = v16
                            _ = v111
                        } else {
                            v112 := GenS12{F0: (v9 & (v9 & v9)), F1: (v9 % (((v9 / ((84) | 1))) | 1)), F2: v1}
                            v113 := GenS14{F0: (v6 << (((v2 >> (((v2 ^ v2)) & 63))) & 63))}
                            _ = v112
                            _ = v113
                        }
                        _ = fi108
                        _ = v110
                    }
                } else {
                    v5 += 2
                    v2 -= v2
                }
                switch "hello" {
                case "\n\t":
                    switch fi32 {
                    case -9223372036854775808:
                        if (!false) {
                            var v114 int32 = 505
                            var v115 int16 = v6
                            _ = v114
                            _ = v115
                        } else {
                            var v116 uint32 = 2147483647
                            fmt.Println(fi32, v4)
                            _ = v116
                        }
                    }
                case "hello":
                    v117 := GenS12{F0: (v9 / ((v9) | 1)), F1: v9, F2: (^(v4))}
                    _ = v117
                case "a":
                    v118 := GenS15{F0: int8(v9), F1: 1, F2: 18446744073709551615}
                    _ = v118
                }
                v119 := GenS13{F0: true, F1: 42, F2: 65535}
                _ = v119
            }
        }
        var v120 int8 = 0
        _ = fi32
        _ = v120
    }
    if (1e10 >= 0.1) {
        var v121 uint32 = 2147483648
        v122 := GenS3{F0: (v17 - v17), F1: 42, F2: 127}
        v123 := GenS3{F0: v17, F1: 42, F2: 35}
        _ = v121
        _ = v122
        _ = v123
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
