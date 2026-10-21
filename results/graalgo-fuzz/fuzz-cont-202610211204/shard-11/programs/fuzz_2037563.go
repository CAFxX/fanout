package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 uint16; F1 int8 }
func genFn0(p0 int16, p1 int) uint32 {
    p0 = (^(p0))
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int16 = 61
        v9 := []int{(^(p1)), fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint64 = 42
        if false {
            if false {
                if ((0.0 - 3.14159) < (1e10 / 1.0)) {
                    v12 := []int{-9223372036854775808, (fi10 & (fi10 * 1)), fi10}
                    fmt.Println(p1, p0, fi10)
                    _ = v12
                } else {
                    var v13 uint = 255
                    _ = v13
                }
                if true {
                    var v14 int8 = -73
                    _ = v14
                } else {
                    var v15 string = "hello"
                    p1 = (p1 | fi10)
                    _ = v15
                }
            } else {
                var v16 float64 = (-2.25 / (3.14159 * -2.25))
                var v17 int16 = (-(p0))
                var v18 int = fi10
                _ = v16
                _ = v17
                _ = v18
            }
            v11 = (^(v11))
        }
        v19 := GenS1{F0: (("" + "") + "\n\t"), F1: 65535, F2: -2147483648}
        _ = fi10
        _ = v11
        _ = v19
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint = 255
    v2 := []uint32{41, 2147483648, 990, 2147483648}
    v3 := make(map[string]uint32)
    var v4 int64 = 220
    var v5 uint8 = 255
    v6 := make(map[int]int)
    v2 = append(v2, 2147483647)
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 uint8 = 2
        if true {
            var v22 int8 = 0
            if (!(3.14159 > 1.5)) {
                var v23 string = "foo bar"
                switch 936 {
                case 9223372036854775807:
                    v24 := GenS2{F0: (v1 | (v1 & v1)), F1: 1}
                    _ = v24
                }
                _ = v23
            } else {
                switch (("" + ("a" + "hello")) + ("" + ("hello" + (((("a" + "hello") + ("\n\t" + "hello")) + "") + ("\n\t" + "hello"))))) {
                case "":
                    delete(v3, "")
                case "hello":
                    v21 -= (v5 | v21)
                case "a":
                    v5 += (^(v21))
                default:
                    var v25 string = "foo bar"
                    _ = v25
                }
                for fi26 := 0; fi26 < 2; fi26++ {
                    switch (fi20 / ((fi26) | 1)) {
                    case -819:
                        var v27 bool = true
                        _ = v27
                    }
                    var v28 int16 = 0
                    var v29 int = fi26
                    var v30 uint64 = 42
                    var v31 uint32 = 4294967295
                    _ = fi26
                    _ = v28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                var v32 int8 = v22
                _ = v32
            }
            v5 = v5
            _ = v22
        } else {
            var v33 int = 2147483647
            _ = v33
        }
        var v34 bool = (-119 < int8(fi20))
        _ = fi20
        _ = v21
        _ = v34
    }
    v35 := GenS2{F0: v1, F1: (v1 * v1)}
    v5 = (v5 | uint8(v1))
    var v36 float32 = 0.1
    for fi37 := 0; fi37 < 1; fi37++ {
        if (32767 > -809) {
            v38 := v3["foo bar"]
            _ = v38
        } else {
            for fi39 := 0; fi39 < 14; fi39++ {
                v1 = (v1 ^ 0)
                if true {
                    if true {
                        var v40 int64 = (v4 | (-(v4)))
                        var v41 float32 = (v36 - 1e10)
                        var v42 int16 = 1
                        _ = v40
                        _ = v41
                        _ = v42
                    } else {
                        fmt.Println(v36)
                        if (genFn0(42, fi39) > 1) {
                            v43 := GenS3{F0: "a"}
                            _ = v43
                        }
                        v44 := GenS2{F0: 831, F1: (v1 & v1)}
                        _ = v44
                    }
                }
                _ = fi39
            }
        }
        v1 = 255
        var v45 int16 = 343
        _ = fi37
        _ = v45
    }
    var v46 float64 = (((1e10 / float64((((v36 + -2.25) * 1e10) - (v36 - 3.14159)))) * 0.1) - 0.0)
    switch 32 {
    case 308:
        v36 -= v36
    case -2147483648:
        for fi47 := 0; fi47 < 2; fi47++ {
            v48 := GenS3{F0: "a"}
            _ = fi47
            _ = v48
        }
    case 1:
        switch "" {
        case "":
            for fi49 := 0; fi49 < 12; fi49++ {
                switch "hello" {
                case "hello":
                    if (false || (-2147483648 <= 42)) {
                        v50 := GenS4{F0: ((("" + ("foo bar" + "")) + ("hello" + "foo bar")) + "a")}
                        for fi51 := 0; fi51 < 14; fi51++ {
                            v6[fi49] = fi49
                            _ = fi51
                        }
                        _ = v50
                    }
                case "foo bar":
                    v4 += v4
                case "":
                    if (v5 >= (^(v5))) {
                        var v52 uint8 = v5
                        _ = v52
                    }
                default:
                    for fi53 := 0; fi53 < 2; fi53++ {
                        v54 := GenS4{F0: ("hello" + "a")}
                        _ = fi53
                        _ = v54
                    }
                }
                if ((809 >= 42) && false) {
                    var v55 uint = (v1 % ((0) | 1))
                    if true {
                        for fi56 := 0; fi56 < 9; fi56++ {
                            if true {
                                switch "" {
                                case "foo bar":
                                    v57 := GenS3{F0: ((("a" + "hello") + "a") + "")}
                                    _ = v57
                                }
                                v55 -= 899
                            }
                            var v58 int8 = 32
                            _ = len(v3)
                            _ = fi56
                            _ = v58
                        }
                        _ = len(v6)
                        switch (-(fi49)) {
                        case -2147483648:
                            _ = len(v2)
                        case 446:
                            switch "a" {
                            case "a":
                                if false {
                                    for fi59 := 0; fi59 < 12; fi59++ {
                                        var v60 uint8 = 145
                                        if true {
                                            v61 := GenS2{F0: (v55 - (v55 - v1)), F1: (v1 / ((v1) | 1))}
                                            if true {
                                                var v62 int16 = 1
                                                v63 := GenS1{F0: ("a" + "foo bar"), F1: 849, F2: 2147483647}
                                                _ = v62
                                                _ = v63
                                            } else {
                                                var v64 float32 = (v36 * v36)
                                                v1 += uint(v5)
                                                var v65 uint16 = 65535
                                                _ = v64
                                                _ = v65
                                            }
                                            _ = v61
                                        }
                                        v2 = append(v2, 569)
                                        _ = fi59
                                        _ = v60
                                    }
                                }
                            case "hello":
                                fmt.Println(v5, v55, fi49)
                            }
                        case 32768:
                            for fi66 := 0; fi66 < 7; fi66++ {
                                var v67 int8 = -1
                                var v68 float64 = 1.5
                                var v69 uint32 = 42
                                _ = fi66
                                _ = v67
                                _ = v68
                                _ = v69
                            }
                        }
                    }
                    if (32767 != 2) {
                        for fi70 := 0; fi70 < 14; fi70++ {
                            v46 += v46
                            var v71 float32 = (v36 + 1.5)
                            v5 = 127
                            _ = fi70
                            _ = v71
                        }
                        if (47 >= (v55 ^ (v55 / (((^(v55))) | 1)))) {
                            var v72 uint16 = 65535
                            _ = v72
                        }
                        if len(v2) > 0 {
                            v2[0] = 2
                        }
                    } else {
                        v4 -= v4
                    }
                    _ = v55
                }
                _ = fi49
            }
        }
    }
    var v73 uint64 = 558
    for fi74 := 0; fi74 < 9; fi74++ {
        if false {
            v75 := GenS4{F0: (("\n\t" + "\n\t") + "foo bar")}
            v5 += v5
            _ = v75
        } else {
            v4 = 1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        switch (fi74 / ((-1) | 1)) {
        case 42:
            v3[("" + (("a" + "foo bar") + "a"))] = 932
        case 32767:
            v4 = (v4 / (((v4 | v4)) | 1))
        case 4294967295:
            v5 -= (v5 ^ v5)
        }
        v46 = 1e10
        _ = fi74
    }
    for fi76 := 0; fi76 < 11; fi76++ {
        v5 -= v5
        for fi77 := 0; fi77 < 3; fi77++ {
            if false {
                switch int(v73) {
                case -584:
                    var v78 int64 = (v4 / ((9223372036854775807) | 1))
                    _ = v78
                }
                v36 = (v36 * float32((float64(v36) + v46)))
                if false {
                    v79 := GenS4{F0: ("\n\t" + "\n\t")}
                    if (v46 <= ((v46 / (v46 * v46)) - v46)) {
                        v80 := v6[fi77]
                        _ = v80
                    } else {
                        v81 := GenS3{F0: ("\n\t" + "hello")}
                        v3[""] = 65535
                        _ = v81
                    }
                    v73 -= v73
                    _ = v79
                } else {
                    switch fi77 {
                    case 127:
                        _ = len(v2)
                    case -1:
                        var v82 uint = (v1 ^ (v1 << ((v1) & 63)))
                        _ = v82
                    }
                    v1 = (^(v1))
                    if (!false) {
                        if false {
                            if false {
                                fmt.Println(v5, v1, v1)
                                var v83 int8 = 100
                                v73 = (^(v73))
                                v5 += v5
                                fmt.Println(v5)
                                var v84 uint32 = 958
                                _ = len(v2)
                                v46 += v46
                                _ = v83
                                _ = v84
                            }
                            v85 := GenS4{F0: "\n\t"}
                            _ = v85
                        } else {
                            if (!false) {
                                if (100 > 89) {
                                    v5 = v5
                                    v86 := GenS1{F0: "hello", F1: uint16(v5), F2: 42}
                                    fmt.Println(v36, v46, v5)
                                    v87 := GenS3{F0: ""}
                                    var v88 string = "a"
                                    switch v88 {
                                    case "a":
                                        if true {
                                            v89 := GenS5{F0: -98, F1: v5, F2: 1000}
                                            v90 := GenS2{F0: v1, F1: v1}
                                            var v91 string = ("" + v88)
                                            _ = v89
                                            _ = v90
                                            _ = v91
                                        }
                                    default:
                                        v46 = v46
                                    }
                                    _ = v86
                                    _ = v87
                                    _ = v88
                                }
                                if (true || false) {
                                    fmt.Println(v73, v36, v46)
                                    v1 -= v1
                                }
                            } else {
                                var v92 int = fi77
                                switch v92 {
                                case 1:
                                    var v93 int = (fi77 & fi76)
                                    v2 = append(v2, uint32(v1))
                                    var v94 uint = (v1 | 1)
                                    _ = len(v3)
                                    _ = v93
                                    _ = v94
                                case 4294967295:
                                    v2 = append(v2, 995)
                                case 773:
                                    v36 = 1.5
                                default:
                                    var v95 uint8 = v5
                                    _ = v95
                                }
                                var v96 uint8 = v5
                                _ = v92
                                _ = v96
                            }
                            fmt.Println(v73, v73, v5)
                        }
                        switch 127 {
                        case 127:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    } else {
                        if (true || (!(((fi76 / ((fi76) | 1)) == -211) || true))) {
                            switch (fi77 >> (((v1 & (v1 >> ((18446744073709551615) & 63)))) & 63)) {
                            case 1:
                                if (((-2.25 / 3.14159) + float64(v36)) > (((v46 + v46) - (v46 + 3.14159)) - ((0.1 - v46) / v46))) {
                                    v97 := GenS2{F0: (^(v1)), F1: (v1 / ((v1) | 1))}
                                    v5 = (v5 >> ((v1) & 63))
                                    _ = v97
                                }
                            }
                            switch (fi77 + fi76) {
                            case 9223372036854775807:
                                if true {
                                    fmt.Println(v46)
                                    switch "" {
                                    case "a":
                                        var v98 int64 = (v4 + (^(v4)))
                                        _ = v98
                                    case "foo bar":
                                        switch 128 {
                                        case 32768:
                                            v46 = ((v46 * 1.5) + (((v46 / (v46 / v46)) - float64(v36)) * float64((((0.0 * (0.1 + 1.5)) / ((1e10 * (float32((((v46 / float64(v36)) * ((1e10 + v46) / -2.25)) * 0.0)) * (0.0 - v36))) + 0.1)) + 1e10))))
                                        case -724:
                                            switch (fi76 >> ((255) & 63)) {
                                            case -638:
                                                if (!false) {
                                                    var v99 uint64 = (v73 - 0)
                                                    var v100 int16 = 492
                                                    v4 -= v4
                                                    _ = v100
                                                    _ = v99
                                                }
                                            case 42:
                                                fmt.Println(v73, v4)
                                            case 194:
                                                var v101 uint64 = v73
                                                _ = v101
                                            }
                                        }
                                    case "\n\t":
                                        if true {
                                            v102 := GenS4{F0: "a"}
                                            _ = v102
                                        }
                                    }
                                    v36 = v36
                                } else {
                                    v5 = v5
                                }
                            default:
                                var v103 int32 = 1
                                _ = v103
                            }
                        }
                    }
                }
                if ((v46 / ((v46 - float64(-2.25)) + (((0.0 / -2.25) / float64(v36)) + v46))) == float64((-2.25 + 1e10))) {
                    v5 = v5
                } else {
                    v104 := GenS2{F0: (v1 * v1), F1: 631}
                    var v105 int8 = int8(v73)
                    var v106 uint64 = (v73 * 1)
                    _ = v104
                    _ = v105
                    _ = v106
                }
                var v107 uint = 854
                _ = v107
            }
            var v108 string = ((("foo bar" + ("foo bar" + "hello")) + "\n\t") + "a")
            _ = fi77
            _ = v108
        }
        switch fi76 {
        case 128:
            if (-39 > 32767) {
                if (0.1 <= v46) {
                    var v109 int32 = 2147483647
                    var v110 uint = (v1 << ((518) & 63))
                    if (!true) {
                        if true {
                            v4 -= (^(v4))
                            fmt.Println(v36, v109, v110)
                            for fi111 := 0; fi111 < 6; fi111++ {
                                var v112 uint8 = (v5 % ((v5) | 1))
                                _ = fi111
                                _ = v112
                            }
                        } else {
                            v113 := GenS6{F0: v5}
                            var v114 string = ""
                            v115 := GenS7{F0: v5, F1: (v110 == v1)}
                            _ = v113
                            _ = v114
                            _ = v115
                        }
                        for fi116 := 0; fi116 < 3; fi116++ {
                            if (752 > 32767) {
                                var v117 uint32 = genFn0(32767, fi76)
                                v110 = 42
                                _ = v117
                            } else {
                                v118 := GenS8{F0: 0, F1: 117}
                                var v119 uint64 = 18446744073709551615
                                _ = v118
                                _ = v119
                            }
                            var v120 int64 = (v4 * (^(v4)))
                            _ = fi116
                            _ = v120
                        }
                        if false {
                            var v121 uint16 = uint16(v73)
                            if true {
                                var v122 int64 = (v4 + (v4 ^ -1))
                                var v123 uint = uint(v121)
                                _ = v122
                                _ = v123
                            }
                            _ = v121
                        }
                    } else {
                        v124 := GenS6{F0: 199}
                        _ = v124
                    }
                    _ = v109
                    _ = v110
                }
            } else {
                for fi125 := 0; fi125 < 16; fi125++ {
                    v46 = v46
                    _ = fi125
                }
            }
        default:
            var v126 int32 = -438
            _ = v126
        }
        _ = fi76
    }
    var v127 float32 = 0.0
    fmt.Println(v5, v127, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v127:", v127)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v73:", v73)
    _ = v1
    _ = v127
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v73
}
