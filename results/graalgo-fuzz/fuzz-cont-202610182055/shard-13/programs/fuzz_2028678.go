package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool; F1 uint }
type GenS10 struct { F0 int64; F1 uint; F2 uint }
type GenS11 struct { F0 uint8; F1 int8 }
type GenS12 struct { F0 bool; F1 int }
type GenS13 struct { F0 int16; F1 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 32767
    var v3 float64 = 3.14159
    var v4 uint32 = 0
    for fi5 := 0; fi5 < 16; fi5++ {
        if false {
            var v6 int8 = 100
            v7 := v1[(fi5 - (fi5 >> ((18446744073709551615) & 63)))]
            _ = v6
            _ = v7
        } else {
            v2 += v2
            var v8 int8 = -128
            _ = v8
        }
        _ = fi5
    }
    var v9 uint64 = 0
    if (42 != -9223372036854775808) {
        fmt.Println(v3, v2)
        switch "foo bar" {
        case "a":
            var v10 uint16 = (^(v2))
            _ = v10
        case "hello":
            for fi11 := 0; fi11 < 8; fi11++ {
                fmt.Println(fi11, fi11)
                _ = fi11
            }
        default:
            if (uint8(v4) < 0) {
                v2 -= 78
                v3 = 1.5
            } else {
                v12 := []int{127, -1}
                switch "a" {
                case "a":
                    v13 := GenS1{F0: (false || false), F1: 42, F2: 1}
                    _ = v13
                case "\n\t":
                    var v14 uint16 = v2
                    _ = v14
                case "foo bar":
                    if false {
                        switch -2147483648 {
                        case 32767:
                            _ = len(v12)
                        }
                        var v15 uint = uint(v2)
                        var v16 uint32 = (v4 >> ((v15) & 63))
                        _ = v15
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 10; fi17++ {
                            if (0 >= 100) {
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    if ((("hello" < ("\n\t" + "")) || true) || false) {
                                        if true {
                                            var v19 float64 = ((((v3 * 1.5) / (-2.25 + (((float64(1e10) + v3) / ((v3 * (v3 * v3)) / 1e10)) / v3))) / (v3 - (1.5 * ((-2.25 + 3.14159) / v3)))) + ((v3 / (0.0 * v3)) / ((1e10 - -2.25) + 3.14159)))
                                            v20 := GenS1{F0: false, F1: 42, F2: 0}
                                            if (25 == fi17) {
                                                switch "hello" {
                                                case "foo bar":
                                                    if false {
                                                        var v21 uint = 42
                                                        var v22 uint16 = v2
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                    if len(v12) > 0 {
                                                        v12[0] = (fi18 ^ (^(fi18)))
                                                    }
                                                case "":
                                                    var v23 uint64 = 1
                                                    _ = v23
                                                }
                                            } else {
                                                var v24 int16 = 32767
                                                if false {
                                                    fmt.Println(v2, v4, v19)
                                                    _ = len(v1)
                                                    v25 := GenS2{F0: 653}
                                                    _ = v25
                                                } else {
                                                    if false {
                                                        fmt.Println(fi17, v24, fi17)
                                                        v26 := GenS3{F0: 5}
                                                        var v27 int64 = 9223372036854775807
                                                        _ = v26
                                                        _ = v27
                                                    }
                                                }
                                                switch "hello" {
                                                case "a":
                                                    v28 := GenS3{F0: 42}
                                                    _ = v28
                                                case "hello":
                                                    v29 := GenS2{F0: (^(v2))}
                                                    _ = v29
                                                }
                                                if (-1 >= -1) {
                                                    if len(v12) > 0 {
                                                        v12[0] = (fi17 | fi17)
                                                    }
                                                    var v30 float32 = 0.0
                                                    v9 += v9
                                                    _ = v30
                                                }
                                                v31 := GenS4{F0: 810, F1: 2}
                                                _ = v24
                                                _ = v31
                                            }
                                            _ = v19
                                            _ = v20
                                        }
                                    }
                                    switch fi17 {
                                    case -9223372036854775808:
                                        if (0.0 > 1.5) {
                                            var v32 int8 = 1
                                            var v33 uint16 = v2
                                            _ = v32
                                            _ = v33
                                        } else {
                                            v34 := GenS2{F0: (v2 << ((168) & 63))}
                                            v1[0] = ("foo bar" + "foo bar")
                                            _ = v34
                                        }
                                    case 590:
                                        var v35 uint16 = v2
                                        _ = v35
                                    case 128:
                                        if (0 == 1) {
                                            fmt.Println(v2, fi18)
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v36 uint32 = uint32(v2)
                                                _ = v36
                                            case "a":
                                                v37 := GenS4{F0: 2, F1: 966}
                                                _ = v37
                                            }
                                            fmt.Println(v9, fi17)
                                        }
                                        v38 := GenS3{F0: 1}
                                        _ = v38
                                    }
                                    _ = fi18
                                }
                                for fi39 := 0; fi39 < 17; fi39++ {
                                    if (0 < 9223372036854775807) {
                                        v3 = v3
                                        if (1.5 <= 3.14159) {
                                            v2 = (v2 * (v2 << ((255) & 63)))
                                        }
                                        v1[(fi39 >> ((85) & 63))] = "hello"
                                    }
                                    if (4294967295 < v9) {
                                        v9 = 1
                                    } else {
                                        v3 = float64(((0.0 + 0.1) - (0.1 - 0.0)))
                                        fmt.Println(v2, v2)
                                        v40 := GenS3{F0: -945}
                                        _ = v40
                                    }
                                    _ = fi39
                                }
                            } else {
                                var v41 int32 = -1
                                if ((0.0 <= (((v3 + (v3 / (((v3 + 0.0) + (v3 / float64(3.14159))) * v3))) + v3) / 1.5)) && (int8(v2) >= 68)) {
                                    switch ("a" + "\n\t") {
                                    case "foo bar":
                                        v2 = v2
                                    case "a":
                                        fmt.Println(v2, v4)
                                    }
                                    var v42 int8 = 81
                                    v43 := GenS5{F0: -32768, F1: v9}
                                    _ = v42
                                    _ = v43
                                }
                                _ = v41
                            }
                            _ = fi17
                        }
                        v2 = v2
                        if (-111 != 1) {
                            if (!false) {
                                v44 := GenS3{F0: 2147483647}
                                _ = v44
                            } else {
                                for fi45 := 0; fi45 < 10; fi45++ {
                                    v46 := GenS6{F0: -40}
                                    v47 := GenS3{F0: 1}
                                    _ = fi45
                                    _ = v46
                                    _ = v47
                                }
                            }
                            v48 := GenS7{F0: v2, F1: -1, F2: 0}
                            _ = v48
                        } else {
                            fmt.Println(v9, v4, v2)
                            var v49 uint16 = (v2 ^ 372)
                            _ = v49
                        }
                    }
                }
                _ = v12
            }
        }
        if false {
            v9 += (^(v9))
            v50 := []int{}
            if ((127 <= -2147483648) || (127 == 497)) {
                if ((19 <= v9) && true) {
                    if true {
                        var v51 int8 = 127
                        for fi52 := 0; fi52 < 8; fi52++ {
                            v53 := GenS3{F0: int32(fi52)}
                            _ = fi52
                            _ = v53
                        }
                        if true {
                            for fi54 := 0; fi54 < 1; fi54++ {
                                if false {
                                    if (true || true) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v3 = (3.14159 / -2.25)
                                        case "hello":
                                            fmt.Println(v3, v51, v2)
                                        }
                                        var v55 float64 = v3
                                        var v56 bool = true
                                        _ = v55
                                        _ = v56
                                    }
                                    var v57 string = "foo bar"
                                    v4 = (v4 + (^(v4)))
                                    _ = v57
                                } else {
                                    for fi58 := 0; fi58 < 1; fi58++ {
                                        if false {
                                            if true {
                                                if (-32768 >= 1) {
                                                    v59 := GenS8{F0: 4294967295}
                                                    v60 := GenS4{F0: 18446744073709551615, F1: 259}
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    var v61 bool = true
                                                    v62 := GenS9{F0: false, F1: 65535}
                                                    _ = v61
                                                    _ = v62
                                                }
                                                fmt.Println(v9, fi58, fi58)
                                            }
                                            var v63 uint64 = (v9 + (v9 << ((1) & 63)))
                                            _ = v63
                                        } else {
                                            var v64 uint64 = v9
                                            _ = v64
                                        }
                                        _ = fi58
                                    }
                                    if true {
                                        if (v4 < v4) {
                                            _ = len(v1)
                                            v4 -= 2
                                            v65 := v1[fi54]
                                            _ = v65
                                        }
                                        if false {
                                            var v66 uint64 = v9
                                            _ = v66
                                        } else {
                                            v67 := GenS10{F0: -1, F1: 261, F2: 109}
                                            var v68 bool = (((float32(v3) / 0.1) * 1.5) > (0.0 / 1.0))
                                            for fi69 := 0; fi69 < 9; fi69++ {
                                                var v70 float64 = (v3 + v3)
                                                _ = fi69
                                                _ = v70
                                            }
                                            _ = v67
                                            _ = v68
                                        }
                                    } else {
                                        fmt.Println(v3)
                                    }
                                    var v71 uint16 = v2
                                    _ = v71
                                }
                                var v72 int64 = 0
                                _ = fi54
                                _ = v72
                            }
                        } else {
                            delete(v1, 2147483647)
                            v73 := GenS9{F0: false, F1: 254}
                            _ = v73
                        }
                        _ = v51
                    } else {
                        var v74 uint64 = v9
                        var v75 uint32 = (v4 & (v4 / ((2) | 1)))
                        fmt.Println(v4)
                        _ = v74
                        _ = v75
                    }
                }
            } else {
                v2 = uint16(v4)
                _ = len(v1)
            }
            _ = v50
        }
    } else {
        switch 0 {
        case 32767:
            switch 32768 {
            case 4294967295:
                fmt.Println(v9, v2)
            }
        case 9223372036854775807:
            for fi76 := 0; fi76 < 19; fi76++ {
                v3 += 1e10
                v77 := []int{fi76, 250, fi76}
                _ = fi76
                _ = v77
            }
        default:
            fmt.Println(v9, v9)
        }
        v3 -= v3
        if true {
            v3 = 3.14159
            switch "\n\t" {
            case "\n\t":
                v78 := GenS4{F0: 4294967295, F1: 42}
                _ = v78
            case "":
                var v79 int8 = 127
                _ = v79
            }
            v80 := GenS4{F0: 65535, F1: 2}
            _ = v80
        }
    }
    var v81 int16 = -1
    var v82 int16 = v81
    if false {
        v83 := GenS11{F0: 0, F1: 100}
        v84 := GenS11{F0: 255, F1: 0}
        _ = v83
        _ = v84
    } else {
        v82 = -1
        for fi85 := 0; fi85 < 17; fi85++ {
            v81 = -639
            if false {
                for fi86 := 0; fi86 < 11; fi86++ {
                    v82 += (v82 & (v81 / (((^(v82))) | 1)))
                    v9 = (v9 | (v9 / ((v9) | 1)))
                    var v87 bool = (false || (161 >= uint8(v2)))
                    _ = fi86
                    _ = v87
                }
                v81 -= (v82 % ((v81) | 1))
                for fi88 := 0; fi88 < 3; fi88++ {
                    var v89 int64 = -1
                    var v90 uint64 = v9
                    var v91 uint = 255
                    _ = fi88
                    _ = v89
                    _ = v90
                    _ = v91
                }
            } else {
                delete(v1, (fi85 ^ fi85))
            }
            if (1 > 0) {
                v9 += (v9 + (v9 >> ((18446744073709551615) & 63)))
                for fi92 := 0; fi92 < 17; fi92++ {
                    var v93 int8 = 0
                    v81 = 42
                    var v94 float32 = 1.5
                    _ = fi92
                    _ = v93
                    _ = v94
                }
                v95 := []int{(-(fi85)), fi85}
                _ = v95
            } else {
                for fi96 := 0; fi96 < 19; fi96++ {
                    v97 := GenS10{F0: -124, F1: 65535, F2: uint(v9)}
                    v98 := GenS12{F0: (int16(fi85) < -1), F1: fi96}
                    if false {
                        v99 := GenS11{F0: 30, F1: 42}
                        _ = len(v1)
                        if true {
                            var v100 int16 = v81
                            v9 -= (v9 << ((255) & 63))
                            _ = v100
                        }
                        _ = v99
                    } else {
                        var v101 string = "a"
                        var v102 uint8 = 127
                        _ = v101
                        _ = v102
                    }
                    _ = fi96
                    _ = v97
                    _ = v98
                }
                switch fi85 {
                case 42:
                    v103 := []int{(fi85 % ((32768) | 1)), fi85}
                    _ = v103
                default:
                    var v104 float64 = -2.25
                    _ = v104
                }
            }
            _ = fi85
        }
        v82 = (v81 >> ((994) & 63))
    }
    switch "\n\t" {
    case "hello":
        v9 = (v9 | 42)
    case "\n\t":
        if true {
            var v105 uint8 = 42
            if (true && false) {
                if false {
                    for fi106 := 0; fi106 < 11; fi106++ {
                        for fi107 := 0; fi107 < 11; fi107++ {
                            var v108 uint64 = (v9 / (((v9 % ((v9) | 1))) | 1))
                            v109 := GenS5{F0: v81, F1: v9}
                            _ = fi107
                            _ = v108
                            _ = v109
                        }
                        _ = fi106
                    }
                    var v110 int16 = 41
                    _ = v110
                }
                if false {
                    if (-9223372036854775808 <= 0) {
                        for fi111 := 0; fi111 < 17; fi111++ {
                            for fi112 := 0; fi112 < 7; fi112++ {
                                v113 := []int{int(v105), fi112, fi111}
                                delete(v1, (fi112 - fi112))
                                v3 -= (v3 + 1e10)
                                v114 := GenS11{F0: (v105 / ((uint8(v4)) | 1)), F1: -128}
                                v9 -= v9
                                fmt.Println(v81)
                                _ = fi112
                                _ = v113
                                _ = v114
                            }
                            delete(v1, fi111)
                            _ = fi111
                        }
                        if (v105 > 127) {
                            v82 += (-(v81))
                        }
                    }
                }
                for fi115 := 0; fi115 < 1; fi115++ {
                    if false {
                        fmt.Println(v2, v3, v9)
                        v105 += (v105 << ((4294967295) & 63))
                    } else {
                        delete(v1, 32767)
                    }
                    var v116 bool = false
                    _ = fi115
                    _ = v116
                }
            }
            for fi117 := 0; fi117 < 13; fi117++ {
                for fi118 := 0; fi118 < 3; fi118++ {
                    var v119 int = fi118
                    _ = fi118
                    _ = v119
                }
                if false {
                    var v120 bool = (338 > fi117)
                    v121 := []int{fi117, (-(fi117)), -93}
                    _ = v120
                    _ = v121
                }
                _ = fi117
            }
            _ = v105
        } else {
            v9 = (v9 | v9)
        }
    default:
        v4 -= v4
    }
    for fi122 := 0; fi122 < 9; fi122++ {
        var v123 bool = (true || (-1 < 1))
        _ = fi122
        _ = v123
    }
    for fi124 := 0; fi124 < 1; fi124++ {
        fmt.Println(v9, v9)
        fmt.Println(v3)
        if (!true) {
            var v125 int16 = 105
            v126 := GenS13{F0: v82, F1: 1}
            for fi127 := 0; fi127 < 8; fi127++ {
                v4 += (v4 - v4)
                switch "" {
                case "\n\t":
                    var v128 uint64 = (v9 * 412)
                    _ = v128
                case "hello":
                    v2 = v2
                default:
                    v129 := []int{-71, int(v2)}
                    _ = v129
                }
                _ = fi127
            }
            _ = v125
            _ = v126
        } else {
            v130 := GenS13{F0: v81, F1: int32(v4)}
            v2 = (v2 % (((^(v2))) | 1))
            var v131 uint = 301
            _ = v130
            _ = v131
        }
        _ = fi124
    }
    v132 := GenS3{F0: 42}
    v9 = v9
    var v133 float32 = float32(v3)
    v1[127] = ("hello" + ("a" + (((((("foo bar" + "\n\t") + "foo bar") + "hello") + (("hello" + "foo bar") + "hello")) + ("" + "")) + "")))
    fmt.Println(v2)
    if (1 == -1) {
        v81 = -546
        if false {
            v9 = (v9 >> ((0) & 63))
            fmt.Println(v9, v81)
            if (0 != -18) {
                var v134 uint8 = 0
                v135 := []int{2147483647, 590}
                _ = v134
                _ = v135
            } else {
                for fi136 := 0; fi136 < 2; fi136++ {
                    if (false || (true && true)) {
                        var v137 int16 = 1000
                        var v138 int8 = -128
                        _ = v137
                        _ = v138
                    }
                    _ = fi136
                }
                var v139 int64 = 0
                _ = v139
            }
        }
        if ((true && true) || ((v3 / v3) != -2.25)) {
            switch -2147483648 {
            case -363:
                v2 += (^(v2))
            }
            var v140 uint32 = (v4 >> ((612) & 63))
            _ = v140
        } else {
            _ = len(v1)
            v3 = -2.25
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v133:", v133)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v132
    _ = v133
    _ = v2
    _ = v3
    _ = v4
    _ = v81
    _ = v82
    _ = v9
}
