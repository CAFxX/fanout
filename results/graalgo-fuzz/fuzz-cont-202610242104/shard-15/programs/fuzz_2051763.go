package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS3 struct { F0 int16; F1 string; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64; F1 string }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint64; F1 string; F2 uint64 }
type GenS11 struct { F0 string; F1 int64 }
type GenS12 struct { F0 uint8; F1 uint64 }
type GenS13 struct { F0 int16 }
func genFn0(p0 uint64) int {
    v7 := []int{}
    for fi8 := 0; fi8 < 7; fi8++ {
        p0 -= 330
        fmt.Println(p0)
        _ = fi8
    }
    if false {
        var v9 uint8 = 128
        _ = v9
    }
    p0 = p0
    v10 := GenS1{F0: -1}
    _ = p0
    _ = v10
    _ = v7
    return 96
}
func genFn1(p0 int64, p1 int32, p2 uint8) int16 {
    p1 -= (p1 * (p1 + p1))
    fmt.Println(p1, p2, p0)
    if true {
        v11 := GenS2{F0: 106, F1: 18446744073709551615, F2: 45}
        switch 770 {
        case 9223372036854775807:
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 float32 = 1e10
                if (!true) {
                    var v14 uint32 = 1
                    if (0 != 1) {
                        v14 += (v14 << ((255) & 63))
                    } else {
                        v15 := GenS2{F0: (v14 * 567), F1: 42, F2: -107}
                        for fi16 := 0; fi16 < 18; fi16++ {
                            v17 := []int{1, -1, (fi16 & (^(fi16)))}
                            _ = fi16
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 1; fi18++ {
                            p0 = (p0 & p0)
                            fmt.Println(p2, v14)
                            v19 := GenS3{F0: 1000, F1: ((("" + "") + "") + "foo bar"), F2: (v14 * v14)}
                            _ = fi18
                            _ = v19
                        }
                        _ = v15
                    }
                    _ = v14
                } else {
                    var v20 uint = 255
                    switch "a" {
                    case "foo bar":
                        v21 := []int{fi12}
                        _ = v21
                    case "\n\t":
                        if false {
                            fmt.Println(p2)
                        } else {
                            v22 := []int{(fi12 & (^(fi12))), (fi12 & (fi12 % (((fi12 << ((0) & 63))) | 1)))}
                            for fi23 := 0; fi23 < 15; fi23++ {
                                fmt.Println(v13, p2, v20)
                                _ = fi23
                            }
                            _ = v22
                        }
                    case "":
                        p0 = (p0 | int64(p1))
                    default:
                        if false {
                            v24 := []int{(fi12 & fi12)}
                            _ = v24
                        } else {
                            switch "a" {
                            case "foo bar":
                                p1 += -2147483648
                            default:
                                v25 := GenS2{F0: 583, F1: 2, F2: 42}
                                _ = v25
                            }
                            v20 += (^(v20))
                        }
                    }
                    v13 -= v13
                    _ = v20
                }
                _ = fi12
                _ = v13
            }
        case 42:
            for fi26 := 0; fi26 < 10; fi26++ {
                switch fi26 {
                case 401:
                    var v27 string = "foo bar"
                    _ = v27
                case 0:
                    var v28 int = genFn0(2)
                    _ = v28
                }
                _ = fi26
            }
        case -841:
            p1 += p1
        }
        _ = v11
    } else {
        if (42 == 1) {
            fmt.Println(p1)
            if false {
                var v29 bool = ((p0 + (p0 + p0)) != p0)
                var v30 string = (("a" + "") + "a")
                _ = v29
                _ = v30
            } else {
                p2 = p2
                var v31 float64 = (1e10 * 0.0)
                _ = v31
            }
            v32 := []int{128, 9223372036854775807, 9223372036854775807, 127}
            _ = v32
        } else {
            switch 1 {
            case 671:
                v33 := []int{genFn0(18446744073709551615), -9223372036854775808}
                _ = v33
            }
            for fi34 := 0; fi34 < 17; fi34++ {
                p1 += (^(p1))
                _ = fi34
            }
            var v35 uint8 = (p2 ^ 127)
            _ = v35
        }
        v36 := []int{1, 1, -1, 262}
        _ = v36
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint16 = 65535
    var v2 uint8 = 128
    v3 := make(map[uint32]int)
    v4 := []int{42}
    v5 := []string{"\n\t"}
    var v6 float64 = 3.14159
    var v37 int8 = -128
    if true {
        var v38 float32 = 0.1
        switch 42 {
        case 720:
            switch "\n\t" {
            case "hello":
                v39 := GenS2{F0: 156, F1: 42, F2: (v37 * (v37 + v37))}
                _ = v39
            case "foo bar":
                if true {
                    var v40 int8 = 42
                    if len(v5) > 0 {
                        v5[0] = "foo bar"
                    }
                    if ((!(1 <= 1)) && false) {
                        v37 -= (v37 + 127)
                        for fi41 := 0; fi41 < 15; fi41++ {
                            var v42 string = "hello"
                            var v43 int32 = -1
                            v44 := GenS1{F0: fi41}
                            _ = fi41
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                    } else {
                        for fi45 := 0; fi45 < 8; fi45++ {
                            fmt.Println(v40)
                            var v46 uint16 = v1
                            _ = fi45
                            _ = v46
                        }
                        var v47 int64 = -71
                        switch genFn0(18446744073709551615) {
                        case 1:
                            switch -587 {
                            case -9223372036854775808:
                                switch -2147483648 {
                                case 4294967295:
                                    if (v47 < (v47 ^ (v47 ^ (v47 ^ v47)))) {
                                        switch "a" {
                                        case "":
                                            v47 -= (v47 - v47)
                                        case "a":
                                            v48 := GenS1{F0: -434}
                                            _ = v48
                                        }
                                    }
                                default:
                                    v38 -= ((v38 + v38) + v38)
                                }
                            case 9223372036854775807:
                                v3[1] = 4294967295
                            case 1:
                                var v49 int = 32767
                                _ = v49
                            default:
                                var v50 int16 = 1
                                _ = v50
                            }
                        }
                        _ = v47
                    }
                    _ = v40
                } else {
                    fmt.Println(v37, v37, v37)
                    var v51 uint64 = 646
                    if (!(1 > 1000)) {
                        if (((v2 & (^(v2))) > (v2 ^ 218)) || false) {
                            fmt.Println(v38)
                        } else {
                            v1 = 255
                        }
                        if (128 > 647) {
                            for fi52 := 0; fi52 < 18; fi52++ {
                                fmt.Println(fi52, v51, fi52)
                                v53 := GenS1{F0: (^(fi52))}
                                _ = fi52
                                _ = v53
                            }
                            var v54 uint64 = v51
                            _ = v54
                        } else {
                            if true {
                                fmt.Println(v37)
                            } else {
                                v38 = v38
                                v51 += 0
                            }
                            _ = len(v5)
                        }
                    }
                    _ = v51
                }
            }
        case 32768:
            v55 := GenS4{F0: 4294967295, F1: (v37 ^ (v37 + (-(v37))))}
            _ = v55
        case 127:
            v6 = float64((3.14159 + float32((v6 + float64(((((v38 / (((v38 - 0.0) * 0.1) - ((0.0 - ((0.0 - v38) * (v38 / v38))) / v38))) - (1e10 / (float32(0.0) + (v38 / v38)))) / (v38 + ((1.5 / (3.14159 * (3.14159 - v38))) / (1e10 - 1.5)))) - (float32(v6) - ((float32(((1.5 / 1.0) / (float64(v38) / float64(((1.5 / 1e10) / v38))))) + 1.5) - v38))))))))
        default:
            if (!(!false)) {
                if (486 <= -32768) {
                    var v56 int16 = -919
                    switch "" {
                    case "a":
                        v4 = append(v4, 168)
                    default:
                        v57 := GenS5{F0: 668, F1: 42, F2: 32768}
                        _ = v57
                    }
                    _ = v56
                } else {
                    v58 := GenS4{F0: 248, F1: -111}
                    switch "\n\t" {
                    case "hello":
                        _ = len(v5)
                    case "\n\t":
                        v2 = (^(v2))
                    }
                    if false {
                        fmt.Println(v38, v1)
                    }
                    _ = v58
                }
                switch -2147483648 {
                case -198:
                    var v59 bool = false
                    _ = v59
                case -9223372036854775808:
                    fmt.Println(v6)
                default:
                    for fi60 := 0; fi60 < 18; fi60++ {
                        v61 := GenS6{F0: 2147483647}
                        if (!false) {
                            var v62 float64 = (v6 / (((3.14159 * v6) / v6) * 1e10))
                            _ = v62
                        } else {
                            for fi63 := 0; fi63 < 9; fi63++ {
                                var v64 string = ("hello" + ("\n\t" + ""))
                                var v65 uint16 = v1
                                fmt.Println(v64, fi60)
                                _ = fi63
                                _ = v64
                                _ = v65
                            }
                            switch fi60 {
                            case -2147483648:
                                var v66 int16 = 32767
                                _ = v66
                            case 9223372036854775807:
                                v38 = 1.5
                            }
                            for fi67 := 0; fi67 < 15; fi67++ {
                                fmt.Println(v1, v6)
                                if (4294967295 >= 65535) {
                                    fmt.Println(v37, fi67, v37)
                                    v68 := GenS7{F0: false}
                                    v69 := GenS3{F0: -1, F1: "", F2: 2147483647}
                                    _ = v68
                                    _ = v69
                                }
                                _ = fi67
                            }
                        }
                        v6 -= 1.5
                        _ = fi60
                        _ = v61
                    }
                }
                v6 += (((-2.25 * -2.25) + -2.25) * v6)
            } else {
                for fi70 := 0; fi70 < 14; fi70++ {
                    v38 = v38
                    _ = fi70
                }
                v71 := GenS6{F0: 982}
                _ = v71
            }
        }
        _ = v38
    } else {
        switch int(v2) {
        case 9223372036854775807:
            var v72 bool = true
            _ = v72
        case 98:
            var v73 uint = 2
            _ = v73
        case -1:
            for fi74 := 0; fi74 < 14; fi74++ {
                var v75 int32 = -1
                _ = fi74
                _ = v75
            }
        }
    }
    v2 = (v2 % (((^(v2))) | 1))
    v6 += (((1e10 - -2.25) + -2.25) / v6)
    if len(v4) > 0 {
        v4[0] = -1
    }
    switch -9223372036854775808 {
    case -1:
        switch ("a" + "\n\t") {
        case "hello":
            v76 := GenS6{F0: 2}
            _ = v76
        default:
            if len(v4) > 0 {
                v4[0] = 298
            }
        }
    case 127:
        for fi77 := 0; fi77 < 5; fi77++ {
            _ = len(v3)
            switch genFn0(835) {
            case 484:
                var v78 float64 = -2.25
                _ = v78
            }
            v79 := GenS8{F0: 42, F1: "hello"}
            _ = fi77
            _ = v79
        }
    default:
        v3[54] = -2147483648
    }
    v80 := GenS2{F0: 807, F1: 42, F2: (v37 ^ (v37 << ((4294967295) & 63)))}
    if (!true) {
        if (-2147483648 <= -825) {
            var v81 uint64 = 122
            switch 9223372036854775807 {
            case -408:
                var v82 int32 = -673
                _ = v82
            default:
                var v83 int8 = (v37 & 42)
                _ = v83
            }
            for fi84 := 0; fi84 < 1; fi84++ {
                v2 -= v2
                var v85 int8 = v37
                v86 := GenS1{F0: (fi84 * (fi84 << ((255) & 63)))}
                _ = fi84
                _ = v85
                _ = v86
            }
            _ = v81
        } else {
            var v87 float64 = float64((-2.25 / float32(-2.25)))
            _ = v87
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch "a" {
        case "a":
            for fi88 := 0; fi88 < 7; fi88++ {
                for fi89 := 0; fi89 < 19; fi89++ {
                    v37 -= 41
                    _ = fi89
                }
                for fi90 := 0; fi90 < 15; fi90++ {
                    var v91 uint16 = 42
                    var v92 uint16 = 572
                    v93 := GenS1{F0: genFn0(18446744073709551615)}
                    v94 := GenS9{F0: 42}
                    if false {
                        v91 = v1
                    }
                    if true {
                        v95 := GenS9{F0: -662}
                        v1 += v91
                        _ = v95
                    } else {
                        switch genFn0(42) {
                        case 0:
                            v92 += v1
                            v91 = 65535
                        case 42:
                            switch 32767 {
                            case -1:
                                v1 -= (v1 << ((912) & 63))
                            case 32768:
                                switch ("" + ("" + ("a" + "foo bar"))) {
                                case "\n\t":
                                    var v96 int = fi88
                                    _ = v96
                                case "":
                                    v97 := GenS10{F0: 114, F1: "a", F2: 141}
                                    _ = v97
                                case "a":
                                    var v98 uint64 = 18446744073709551615
                                    _ = v98
                                default:
                                    var v99 uint64 = 42
                                    _ = v99
                                }
                            case 9223372036854775807:
                                v91 = v91
                            }
                            v100 := GenS5{F0: (v1 - (v91 - v91)), F1: uint32(v2), F2: (fi90 | (fi90 | fi88))}
                            _ = v100
                        }
                        _ = len(v3)
                    }
                    v101 := GenS11{F0: ((("" + (("" + "hello") + "a")) + "\n\t") + (("" + "") + ("\n\t" + "a"))), F1: 9223372036854775807}
                    _ = fi90
                    _ = v101
                    _ = v91
                    _ = v92
                    _ = v93
                    _ = v94
                }
                _ = fi88
            }
        case "hello":
            if true {
                fmt.Println(v1, v2)
                switch 500 {
                case 1:
                    v102 := GenS12{F0: (v2 * v2), F1: 1}
                    _ = v102
                case -9223372036854775808:
                    if true {
                        if len(v4) > 0 {
                            v4[0] = 441
                        }
                    }
                default:
                    var v103 uint = 127
                    _ = v103
                }
            } else {
                var v104 uint64 = uint64(v1)
                _ = v104
            }
        default:
            if (502 > 0) {
                v37 = 127
            } else {
                var v105 int16 = -32768
                for fi106 := 0; fi106 < 10; fi106++ {
                    for fi107 := 0; fi107 < 7; fi107++ {
                        fmt.Println(v105)
                        v37 += v37
                        if (true && (!(-32768 == (v105 << ((1) & 63))))) {
                            if true {
                                v105 += v105
                            } else {
                                fmt.Println(v6, fi107)
                                v1 = 32767
                            }
                            var v108 uint8 = 0
                            _ = v108
                        } else {
                            switch "" {
                            case "":
                                fmt.Println(fi107)
                            case "\n\t":
                                var v109 float64 = v6
                                _ = v109
                            case "a":
                                var v110 float32 = (((0.0 * (1e10 + 0.0)) / (1.5 / 0.1)) + 3.14159)
                                _ = v110
                            default:
                                var v111 uint16 = (v1 ^ 42)
                                _ = v111
                            }
                            var v112 string = "a"
                            _ = v112
                        }
                        v105 -= (v105 >> ((242) & 63))
                        var v113 int8 = 33
                        fmt.Println(v2, v1)
                        _ = fi107
                        _ = v113
                    }
                    for fi114 := 0; fi114 < 13; fi114++ {
                        v1 -= v1
                        var v115 int16 = 0
                        v2 += v2
                        _ = fi114
                        _ = v115
                    }
                    _ = fi106
                }
                _ = v105
            }
        }
    }
    var v116 uint = 255
    switch 127 {
    case 1:
        var v117 uint = (v116 + uint(v1))
        _ = v117
    case 611:
        switch 42 {
        case 9223372036854775807:
            for fi118 := 0; fi118 < 16; fi118++ {
                var v119 float64 = ((0.1 + (v6 * (3.14159 * ((v6 + v6) + ((v6 - v6) * v6))))) + v6)
                _ = fi118
                _ = v119
            }
        case -210:
            v1 += (v1 << ((v116) & 63))
        case 4294967295:
            fmt.Println(v1)
        }
    case 32768:
        switch 9223372036854775807 {
        case 4294967295:
            fmt.Println(v6, v116)
        case 42:
            switch 861 {
            case 1:
                var v120 int8 = (v37 % ((-99) | 1))
                _ = v120
            case 127:
                if false {
                    fmt.Println(v37, v1, v2)
                    fmt.Println(v116, v2, v116)
                    var v121 int8 = 100
                    _ = v121
                } else {
                    var v122 uint8 = 111
                    v122 = v122
                    _ = v122
                }
            }
        case -77:
            fmt.Println(v116, v2)
        }
    default:
        v123 := GenS13{F0: 0}
        _ = v123
    }
    v37 -= v37
    v116 += 42
    v37 += (v37 * v37)
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v116
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v80
}
