package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint; F2 string }
type GenS4 struct { F0 int16; F1 bool; F2 int8 }
type GenS5 struct { F0 int; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 int32; F1 string }
type GenS9 struct { F0 uint32 }
func genFn0(p0 string) int8 {
    p0 = "a"
    v5 := []int{}
    var v6 int8 = 1
    if true {
        fmt.Println(p0, p0)
    } else {
        fmt.Println(v6, v6)
        if true {
            p0 = (p0 + p0)
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        for fi7 := 0; fi7 < 14; fi7++ {
            if false {
                var v8 int16 = 1
                v9 := GenS1{F0: 9223372036854775807, F1: 83}
                _ = v8
                _ = v9
            }
            _ = fi7
        }
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        _ = len(v5)
        for fi11 := 0; fi11 < 18; fi11++ {
            p0 = ("\n\t" + p0)
            switch fi11 {
            case -2147483648:
                var v12 int8 = (v6 * (v6 * -115))
                v12 += (v6 & (^(v6)))
                var v13 float64 = 1e10
                var v14 int32 = -304
                _ = v12
                _ = v13
                _ = v14
            }
            v6 = v6
            _ = fi11
        }
        _ = fi10
    }
    _ = p0
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 bool) int64 {
    var v15 uint16 = 65535
    var v16 uint = 4294967295
    if p0 {
        switch "a" {
        case "foo bar":
            var v17 bool = (0 <= genFn0("a"))
            _ = v17
        case "hello":
            v18 := GenS2{F0: 65535}
            _ = v18
        case "a":
            switch "" {
            case "a":
                v19 := GenS1{F0: int64(v15), F1: 127}
                _ = v19
            case "hello":
                v15 -= v15
            case "\n\t":
                var v20 int32 = 0
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            p0 = (!true)
            if (!true) {
                var v22 float32 = 0.0
                var v23 uint16 = (v15 ^ (^(v15)))
                var v24 float64 = 1.5
                _ = v22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
    } else {
        fmt.Println(p0, v15)
    }
    _ = p0
    _ = v15
    _ = v16
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 3.14159
    var v3 int32 = 42
    var v4 string = "\n\t"
    var v25 uint64 = 18446744073709551615
    var v26 int8 = 100
    if true {
        for fi27 := 0; fi27 < 9; fi27++ {
            v25 -= v25
            var v28 int = (fi27 | 9223372036854775807)
            v29 := []int{fi27, v28, v28}
            _ = fi27
            _ = v28
            _ = v29
        }
        if (v2 <= 0.0) {
            v26 -= (^(v26))
            var v30 int16 = -32768
            v26 -= (v26 | -2)
            _ = v30
        } else {
            var v31 float64 = (1.5 + -2.25)
            _ = v31
        }
        v3 += 877
    } else {
        if (18446744073709551615 > 18446744073709551615) {
            var v32 int8 = v26
            _ = v32
        }
    }
    var v33 float64 = 1e10
    for fi34 := 0; fi34 < 17; fi34++ {
        fmt.Println(v26, v3, v25)
        v35 := GenS3{F0: (fi34 >> ((18446744073709551615) & 63)), F1: 703, F2: v4}
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        v37 := []int{}
        for fi38 := 0; fi38 < 16; fi38++ {
            var v39 uint16 = 114
            v40 := GenS4{F0: 1000, F1: true, F2: int8(fi36)}
            var v41 uint64 = (v25 - (v25 | v25))
            switch "\n\t" {
            case "a":
                fmt.Println(v2, v2, v39)
            case "foo bar":
                if ((0.1 + (v33 / v33)) <= (v33 - (float64(((((((1e10 - -2.25) - float32(((v33 / (-2.25 / 3.14159)) + v33))) - v2) - 1.5) / ((v2 * v2) / v2)) * (3.14159 + v2))) * v33))) {
                    var v42 int8 = v26
                    fmt.Println(v39)
                    _ = v42
                } else {
                    var v43 float64 = float64((v2 - (-2.25 - v2)))
                    var v44 float64 = 0.0
                    v45 := v1[v4]
                    if true {
                        if false {
                            var v46 int8 = (-(v26))
                            _ = v46
                        } else {
                            var v47 uint16 = v39
                            v48 := GenS4{F0: 1, F1: (0 != 1), F2: v26}
                            if (true && false) {
                                v1[""] = (^(v45))
                                if len(v37) > 0 {
                                    v37[0] = (^(fi36))
                                }
                            }
                            _ = v47
                            _ = v48
                        }
                    }
                    var v49 uint = 0
                    if (!(!(false || (!((v41 - uint64(v39)) != v25))))) {
                        var v50 int = (fi38 ^ fi36)
                        v51 := v1[v4]
                        v52 := GenS1{F0: genFn1((int(v26) != (fi36 - (fi38 << ((v49) & 63))))), F1: 1}
                        fmt.Println(v39, v49, v49)
                        if len(v37) > 0 {
                            v37[0] = -1
                        }
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                    v39 -= v39
                    var v53 uint16 = (v39 | v39)
                    var v54 int64 = -1
                    _ = v43
                    _ = v44
                    _ = v45
                    _ = v49
                    _ = v53
                    _ = v54
                }
            default:
                if false {
                    v55 := GenS5{F0: (fi38 * (fi38 % ((32768) | 1))), F1: -262, F2: (v41 ^ v41)}
                    v4 = ("hello" + "\n\t")
                    _ = v55
                } else {
                    var v56 uint16 = v39
                    v57 := GenS1{F0: int64(v26), F1: 0}
                    _ = v56
                    _ = v57
                }
            }
            fmt.Println(v33, v39)
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
        _ = fi36
        _ = v37
    }
    v2 = 0.1
    v58 := []int{-540}
    if false {
        if (false || false) {
            if true {
                fmt.Println(v4)
                if false {
                    var v59 string = v4
                    switch 128 {
                    case 4294967295:
                        v58 = append(v58, 97)
                    }
                    var v60 int8 = 29
                    _ = v59
                    _ = v60
                } else {
                    switch (v4 + v4) {
                    case "a":
                        v61 := GenS5{F0: 32768, F1: 1000, F2: (v25 ^ (v25 << ((4294967295) & 63)))}
                        _ = v61
                    case "foo bar":
                        if false {
                            v62 := GenS2{F0: 4294967295}
                            if (!true) {
                                for fi63 := 0; fi63 < 11; fi63++ {
                                    v64 := GenS2{F0: 500}
                                    fmt.Println(v33)
                                    _ = fi63
                                    _ = v64
                                }
                                var v65 int32 = v3
                                v1["foo bar"] = 1
                                _ = v65
                            } else {
                                var v66 int32 = 1
                                var v67 int32 = 802
                                _ = v66
                                _ = v67
                            }
                            if true {
                                v2 += v2
                            } else {
                                for fi68 := 0; fi68 < 10; fi68++ {
                                    if len(v58) >= 2 {
                                        v58 = v58[1:len(v58)]
                                    }
                                    for fi69 := 0; fi69 < 10; fi69++ {
                                        fmt.Println(v26, v2)
                                        fmt.Println(fi68)
                                        _ = fi69
                                    }
                                    _ = fi68
                                }
                            }
                            _ = v62
                        } else {
                            var v70 float64 = (v33 + v33)
                            v71 := GenS2{F0: 65535}
                            _ = len(v58)
                            _ = v70
                            _ = v71
                        }
                    case "":
                        switch v4 {
                        case "foo bar":
                            var v72 bool = true
                            _ = v72
                        case "\n\t":
                            for fi73 := 0; fi73 < 7; fi73++ {
                                if ((v3 & v3) < (v3 % ((v3) | 1))) {
                                    fmt.Println(v3, v4)
                                    var v74 uint8 = 202
                                    _ = v74
                                } else {
                                    v75 := GenS6{F0: (v25 & (v25 >> ((381) & 63))), F1: (fi73 ^ fi73)}
                                    _ = v75
                                }
                                _ = fi73
                            }
                        case "a":
                            switch v4 {
                            case "hello":
                                if (-9223372036854775808 >= 4294967295) {
                                    for fi76 := 0; fi76 < 18; fi76++ {
                                        v77 := GenS4{F0: -1, F1: true, F2: v26}
                                        fmt.Println(v4, v3)
                                        _ = fi76
                                        _ = v77
                                    }
                                } else {
                                    v2 += (0.0 * (v2 / 3.14159))
                                    fmt.Println(v33, v33)
                                    if (!false) {
                                        for fi78 := 0; fi78 < 15; fi78++ {
                                            for fi79 := 0; fi79 < 15; fi79++ {
                                                v33 = (v33 - v33)
                                                fmt.Println(v25, v33, v25)
                                                v80 := v1[("" + "a")]
                                                _ = len(v1)
                                                v81 := GenS7{F0: 65535, F1: -9223372036854775808}
                                                _ = fi79
                                                _ = v80
                                                _ = v81
                                            }
                                            _ = fi78
                                        }
                                        switch v4 {
                                        case "a":
                                            if (((v2 * 3.14159) + 0.1) < v2) {
                                                for fi82 := 0; fi82 < 4; fi82++ {
                                                    if false {
                                                        v25 += v25
                                                    }
                                                    v3 += (v3 >> ((65535) & 63))
                                                    var v83 float32 = 3.14159
                                                    _ = fi82
                                                    _ = v83
                                                }
                                                var v84 uint32 = 42
                                                v25 -= v25
                                                _ = v84
                                            }
                                        case "\n\t":
                                            if ((v3 << ((703) & 63)) > (v3 + 173)) {
                                                for fi85 := 0; fi85 < 7; fi85++ {
                                                    var v86 float32 = v2
                                                    if false {
                                                        v3 = v3
                                                        v2 -= ((((((v86 - (v2 / 0.1)) * v2) / v86) + ((float32(v33) + float32((0.0 + (-2.25 + ((-2.25 + (0.0 / (1e10 - (float64(v86) / 0.1)))) / 1e10))))) + float32(v33))) - v2) / (v2 / -2.25))
                                                        if len(v58) >= 2 {
                                                            v58 = v58[1:len(v58)]
                                                        }
                                                    }
                                                    _ = fi85
                                                    _ = v86
                                                }
                                            }
                                        default:
                                            for fi87 := 0; fi87 < 19; fi87++ {
                                                _ = len(v1)
                                                if true {
                                                    v25 -= v25
                                                    for fi88 := 0; fi88 < 6; fi88++ {
                                                        v2 = v2
                                                        _ = fi88
                                                    }
                                                    var v89 float32 = v2
                                                    _ = v89
                                                } else {
                                                    var v90 bool = true
                                                    for fi91 := 0; fi91 < 16; fi91++ {
                                                        if (v90 && v90) {
                                                            v92 := GenS5{F0: (^(fi91)), F1: int16(v25), F2: 4294967295}
                                                            delete(v1, (v4 + v4))
                                                            _ = v92
                                                        }
                                                        switch 0 {
                                                        case 127:
                                                            switch 42 {
                                                            case -620:
                                                                var v93 uint = 18446744073709551615
                                                                _ = v93
                                                            case 4294967295:
                                                                v58 = append(v58, (fi87 & fi91))
                                                            case 32768:
                                                                v26 = (v26 | genFn0(v4))
                                                            default:
                                                                if (42 != 128) {
                                                                    if len(v58) > 0 {
                                                                        v58[0] = (fi91 << ((4294967295) & 63))
                                                                    }
                                                                    var v94 float64 = (v33 * v33)
                                                                    if v90 {
                                                                        var v95 uint64 = v25
                                                                        fmt.Println(v26, fi91, v95)
                                                                        _ = v95
                                                                    }
                                                                    _ = v94
                                                                }
                                                            }
                                                        case 25:
                                                            if ((v26 | (v26 ^ 42)) == (v26 - (v26 << ((165) & 63)))) {
                                                                v96 := GenS8{F0: 42, F1: (v4 + v4)}
                                                                _ = v96
                                                            }
                                                        case 0:
                                                            if (v4 <= ((v4 + (v4 + "foo bar")) + "")) {
                                                                if v90 {
                                                                    v2 -= -2.25
                                                                    v26 -= (v26 << ((4294967295) & 63))
                                                                    var v97 uint8 = 1
                                                                    switch v4 {
                                                                    case "a":
                                                                        var v98 string = v4
                                                                        _ = v98
                                                                    }
                                                                    v90 = (!(108 >= 98))
                                                                    _ = v97
                                                                }
                                                            } else {
                                                                v99 := GenS6{F0: (v25 * 4294967295), F1: (fi91 / ((fi87) | 1))}
                                                                v3 = v3
                                                                _ = v99
                                                            }
                                                        default:
                                                            if v90 {
                                                                v1[(((v4 + (((v4 + v4) + v4) + "\n\t")) + v4) + v4)] = 2147483647
                                                                var v100 int = (fi87 >> ((255) & 63))
                                                                switch (^(fi91)) {
                                                                case 2147483647:
                                                                    _ = len(v58)
                                                                case 1:
                                                                    v3 = (v3 + v3)
                                                                case -9223372036854775808:
                                                                    fmt.Println(v3)
                                                                default:
                                                                    if true {
                                                                        var v101 uint32 = 903
                                                                        fmt.Println(v26, v26, v26)
                                                                        if (!true) {
                                                                            v26 = v26
                                                                            v3 = -2147483648
                                                                            var v102 float64 = v33
                                                                            _ = v102
                                                                        }
                                                                        _ = v101
                                                                    } else {
                                                                        _ = len(v58)
                                                                        var v103 uint16 = 32767
                                                                        fmt.Println(v100, v26, v100)
                                                                        fmt.Println(fi87, v103, v25)
                                                                        v100 = 257
                                                                        _ = v103
                                                                    }
                                                                }
                                                                _ = v100
                                                            }
                                                        }
                                                        _ = fi91
                                                    }
                                                    var v104 uint64 = 1
                                                    _ = v104
                                                    _ = v90
                                                }
                                                if false {
                                                    var v105 float32 = v2
                                                    switch fi87 {
                                                    case 2147483647:
                                                        if len(v58) > 0 {
                                                            v58[0] = (fi87 % (((-(fi87))) | 1))
                                                        }
                                                    case 1:
                                                        var v106 int8 = v26
                                                        _ = v106
                                                    }
                                                    _ = v105
                                                } else {
                                                    var v107 int16 = 42
                                                    _ = v107
                                                }
                                                _ = fi87
                                            }
                                        }
                                    }
                                }
                            }
                        default:
                            var v108 float32 = v2
                            _ = v108
                        }
                    }
                    v3 += (v3 / (((v3 ^ 42)) | 1))
                }
                var v109 string = (("a" + ("foo bar" + v4)) + ("hello" + "a"))
                _ = v109
            } else {
                if ((v26 & v26) <= (v26 / (((v26 % ((v26) | 1))) | 1))) {
                    var v110 uint16 = 2
                    var v111 int32 = (v3 ^ (v3 >> ((208) & 63)))
                    _ = v110
                    _ = v111
                } else {
                    if (false || true) {
                        var v112 int16 = -32768
                        _ = v112
                    } else {
                        var v113 int = 0
                        v114 := GenS9{F0: 1}
                        _ = v113
                        _ = v114
                    }
                    fmt.Println(v33, v2, v33)
                    v26 = (v26 % ((v26) | 1))
                }
                for fi115 := 0; fi115 < 4; fi115++ {
                    fmt.Println(v26, v3)
                    switch "\n\t" {
                    case "foo bar":
                        if (("foo bar" >= ("foo bar" + "foo bar")) || true) {
                            var v116 string = "a"
                            _ = v116
                        }
                    case "a":
                        if true {
                            v26 += (v26 + v26)
                            for fi117 := 0; fi117 < 18; fi117++ {
                                if (1 == 2) {
                                    if true {
                                        var v118 uint = 65535
                                        if ((v2 + v2) == (v2 / 0.1)) {
                                            var v119 int64 = 0
                                            v118 = v118
                                            _ = v119
                                        }
                                        switch int(v118) {
                                        case 1:
                                            v26 = (v26 * v26)
                                        case 390:
                                            fmt.Println(v2)
                                        }
                                        _ = v118
                                    }
                                    v120 := GenS5{F0: (fi115 / (((fi115 >> ((492) & 63))) | 1)), F1: 1000, F2: (v25 | v25)}
                                    _ = v120
                                }
                                _ = fi117
                            }
                            v25 = v25
                        }
                    }
                    var v121 uint16 = 928
                    _ = fi115
                    _ = v121
                }
                for fi122 := 0; fi122 < 2; fi122++ {
                    fmt.Println(v3, v26)
                    var v123 float64 = v33
                    for fi124 := 0; fi124 < 15; fi124++ {
                        var v125 bool = (-1 >= v3)
                        _ = fi124
                        _ = v125
                    }
                    _ = fi122
                    _ = v123
                }
            }
        } else {
            var v126 uint64 = (v25 % ((v25) | 1))
            _ = v126
        }
    } else {
        for fi127 := 0; fi127 < 4; fi127++ {
            var v128 int32 = v3
            var v129 uint8 = 42
            var v130 int64 = int64(v26)
            _ = fi127
            _ = v128
            _ = v129
            _ = v130
        }
    }
    for fi131 := 0; fi131 < 16; fi131++ {
        var v132 int32 = -552
        _ = fi131
        _ = v132
    }
    var v133 int8 = genFn0(v4)
    v134 := GenS7{F0: 2, F1: 0}
    v135 := v1[v4]
    var v136 int8 = v26
    v26 = (v136 << ((0) & 63))
    switch "foo bar" {
    case "a":
        v137 := GenS6{F0: v25, F1: 861}
        _ = v137
    default:
        for fi138 := 0; fi138 < 13; fi138++ {
            v139 := GenS7{F0: 65535, F1: -1}
            v133 = 100
            _ = fi138
            _ = v139
        }
    }
    for fi140 := 0; fi140 < 15; fi140++ {
        var v141 int = (^(fi140))
        v3 = (v3 | (v3 / ((v3) | 1)))
        _ = fi140
        _ = v141
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v133:", v133)
    fmt.Println("v135:", v135)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v58:", len(v58), v58)
    _ = v1
    _ = v133
    _ = v134
    _ = v135
    _ = v136
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v58
}
