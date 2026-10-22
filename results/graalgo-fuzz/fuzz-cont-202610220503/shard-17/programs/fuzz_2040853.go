package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int8; F1 string; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16; F1 bool; F2 uint }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32; F1 string }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int64; F1 int8; F2 uint8 }
func genFn0(p0 int16, p1 uint16) int16 {
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 += p0
        if (true && (true && true)) {
            p1 = p1
        }
        v7 := []int{(fi6 ^ 127), fi6}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        switch "hello" {
        case "hello":
            v9 := GenS1{F0: 2, F1: 728, F2: 0}
            _ = v9
        case "":
            var v10 uint32 = 0
            _ = v10
        case "foo bar":
            for fi11 := 0; fi11 < 1; fi11++ {
                fmt.Println(p1)
                if true {
                    var v12 int = 1
                    v12 += fi11
                    p0 -= 1
                    v13 := []int{}
                    if len(v13) > 0 {
                        v13[0] = 42
                    }
                    v14 := GenS1{F0: 88, F1: 1, F2: 42}
                    p1 -= 32767
                    v13 = append(v13, fi8)
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    if false {
                        v15 := []int{fi11, (fi8 | (-(fi8))), fi8, fi11}
                        _ = v15
                    }
                    var v16 uint64 = 2
                    if false {
                        var v17 int = (fi8 >> ((uint(v16)) & 63))
                        var v18 uint64 = 0
                        var v19 int = -729
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        var v20 uint64 = (^(v16))
                        _ = v20
                    }
                    _ = v16
                }
                p0 = 32767
                _ = fi11
            }
        default:
            var v21 int = (fi8 - (fi8 >> ((429) & 63)))
            _ = v21
        }
        p1 = p1
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := []int{fi22, (fi22 & fi8)}
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            p1 = 61
            _ = fi22
            _ = v23
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    return int16(p1)
}
func genFn1(p0 int64, p1 int8, p2 uint) uint {
    v24 := []int{0, -415, -9223372036854775808, 32768}
    var v25 uint16 = 416
    var v26 float32 = (3.14159 * ((0.1 - (3.14159 - -2.25)) * float32(0.0)))
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    switch "foo bar" {
    case "":
        var v27 uint64 = 571
        _ = v27
    case "a":
        v24 = append(v24, 32768)
    case "hello":
        p2 = 742
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    _ = v26
    return p2
}

func main() {
    var v1 int32 = 1
    var v2 bool = ((v1 > -1) || (false || true))
    var v3 float32 = float32(-2.25)
    v4 := make(map[uint32]uint32)
    var v5 int32 = v1
    v28 := GenS1{F0: 0, F1: 1, F2: 2}
    var v29 uint32 = 0
    var v30 float32 = 1e10
    if (("foo bar" <= "a") || v2) {
        v3 += (v30 / (1.5 + (v30 - (v30 * v3))))
    } else {
        v1 += -407
        v29 = v29
        if v2 {
            if false {
                fmt.Println(v1)
                fmt.Println(v5, v2)
                var v31 uint = 4294967295
                _ = v31
            } else {
                delete(v4, (v29 & 641))
            }
            switch -9223372036854775808 {
            case 0:
                v1 += (v1 + 2147483647)
            case 32767:
                if false {
                    var v32 uint = 88
                    switch ("a" + ("foo bar" + ("" + (("a" + "") + "hello")))) {
                    case "a":
                        v30 = float32(3.14159)
                    default:
                        if v2 {
                            v29 = 65535
                            switch -321 {
                            case 128:
                                if v2 {
                                    v33 := GenS1{F0: 1, F1: (v29 / (((v29 & (^(v29)))) | 1)), F2: 4294967295}
                                    switch "a" {
                                    case "foo bar":
                                        var v34 uint8 = 127
                                        _ = v34
                                    }
                                    if ("hello" == "foo bar") {
                                        if (61 == 255) {
                                            var v35 int8 = -128
                                            _ = v35
                                        } else {
                                            var v36 float64 = 1e10
                                            _ = v36
                                        }
                                        switch 32767 {
                                        case 9223372036854775807:
                                            switch 2147483647 {
                                            case 127:
                                                v32 = (v32 - 143)
                                            case 32768:
                                                if (1 < -626) {
                                                    if true {
                                                        v3 = v30
                                                    } else {
                                                        var v37 uint64 = 536
                                                        if v2 {
                                                            var v38 int16 = 75
                                                            for fi39 := 0; fi39 < 16; fi39++ {
                                                                v40 := []int{}
                                                                v40 = append(v40, (fi39 % ((fi39) | 1)))
                                                                v32 = (v32 - (^(v32)))
                                                                _ = fi39
                                                                _ = v40
                                                            }
                                                            fmt.Println(v3, v5)
                                                            _ = v38
                                                        }
                                                        if ("hello" <= ("hello" + ("\n\t" + "foo bar"))) {
                                                            v4[(v29 & v29)] = v29
                                                        } else {
                                                            fmt.Println(v37, v3, v1)
                                                        }
                                                        _ = v37
                                                    }
                                                    if false {
                                                        fmt.Println(v29, v32)
                                                        fmt.Println(v32, v3)
                                                    } else {
                                                        for fi41 := 0; fi41 < 16; fi41++ {
                                                            for fi42 := 0; fi42 < 16; fi42++ {
                                                                fmt.Println(v29, v1)
                                                                _ = fi42
                                                            }
                                                            v43 := []int{(fi41 >> ((255) & 63)), fi41}
                                                            switch (fi41 - (fi41 ^ (^(fi41)))) {
                                                            case -364:
                                                                v44 := GenS2{F0: 77, F1: "", F2: 1}
                                                                _ = v44
                                                            case -9223372036854775808:
                                                                var v45 int32 = 30
                                                                _ = v45
                                                            case 4294967295:
                                                                var v46 int64 = -9223372036854775808
                                                                _ = v46
                                                            default:
                                                                v47 := GenS3{F0: true}
                                                                _ = v47
                                                            }
                                                            _ = fi41
                                                            _ = v43
                                                        }
                                                        v29 += 0
                                                        var v48 int16 = 1
                                                        _ = v48
                                                    }
                                                }
                                            case -805:
                                                v4[(v29 % ((v29) | 1))] = (v29 + 42)
                                            default:
                                                var v49 uint8 = 128
                                                _ = v49
                                            }
                                        case 2147483647:
                                            v30 = v3
                                        case 42:
                                            delete(v4, 2147483648)
                                        default:
                                            fmt.Println(v2, v1, v3)
                                        }
                                        switch 0 {
                                        case 32767:
                                            v50 := GenS4{F0: -47}
                                            _ = v50
                                        }
                                    }
                                    _ = v33
                                } else {
                                    for fi51 := 0; fi51 < 18; fi51++ {
                                        v3 -= v30
                                        v30 = (v3 * (1.5 + 0.1))
                                        v52 := v4[v29]
                                        _ = fi51
                                        _ = v52
                                    }
                                }
                            case 32767:
                                v53 := v4[(^(v29))]
                                _ = v53
                            case 91:
                                if (!(-24 >= -128)) {
                                    for fi54 := 0; fi54 < 20; fi54++ {
                                        var v55 uint32 = 443
                                        _ = fi54
                                        _ = v55
                                    }
                                } else {
                                    for fi56 := 0; fi56 < 5; fi56++ {
                                        if (74 <= (v29 & 232)) {
                                            if v2 {
                                                if v2 {
                                                    v57 := GenS5{F0: 128}
                                                    v2 = (false || (65535 >= 255))
                                                    _ = v57
                                                } else {
                                                    if (-99 != 32) {
                                                        v58 := []int{1}
                                                        if (v29 != v29) {
                                                            for fi59 := 0; fi59 < 17; fi59++ {
                                                                v29 = (^(v29))
                                                                if v2 {
                                                                    if false {
                                                                        if (v2 || v2) {
                                                                            v3 = (v3 + ((1.5 + 0.0) + float32(0.1)))
                                                                            v1 = (v1 - v1)
                                                                            v60 := v4[v29]
                                                                            _ = v60
                                                                        } else {
                                                                            var v61 float64 = -2.25
                                                                            var v62 int32 = v5
                                                                            _ = v61
                                                                            _ = v62
                                                                        }
                                                                        switch (fi59 ^ fi56) {
                                                                        case -1:
                                                                            v63 := GenS4{F0: 1000}
                                                                            _ = v63
                                                                        default:
                                                                            if (-80 < 127) {
                                                                                if ((v32 << (((v32 % (((v32 + (v32 | (v32 / ((v32) | 1))))) | 1))) & 63)) != genFn1(76, -61, (v32 ^ 42))) {
                                                                                    var v64 uint = v32
                                                                                    if len(v58) > 0 {
                                                                                        v58[0] = fi56
                                                                                    }
                                                                                    _ = v64
                                                                                } else {
                                                                                    v32 += (^(v32))
                                                                                    v5 = v1
                                                                                    var v65 int8 = 1
                                                                                    _ = v65
                                                                                }
                                                                                if len(v58) > 0 {
                                                                                    v58[0] = (fi59 + fi56)
                                                                                }
                                                                                var v66 int64 = 9223372036854775807
                                                                                _ = v66
                                                                            } else {
                                                                                fmt.Println(v29, v29, v30)
                                                                                fmt.Println(v29, v32, fi56)
                                                                                var v67 int = 128
                                                                                var v68 int = 127
                                                                                fmt.Println(v67, v30, v67)
                                                                                v30 = v3
                                                                                v69 := GenS6{F0: 1000}
                                                                                v68 += fi59
                                                                                v68 = (v68 ^ v67)
                                                                                _ = v67
                                                                                _ = v68
                                                                                _ = v69
                                                                            }
                                                                        }
                                                                        if (199 < 243) {
                                                                            v29 -= v29
                                                                            if (1 < 18446744073709551615) {
                                                                                var v70 string = "\n\t"
                                                                                if ((false || (52 != 42)) && (58 == 128)) {
                                                                                    v71 := GenS7{F0: genFn0(-237, 1), F1: (-592 < 0), F2: (v32 ^ v32)}
                                                                                    _ = v71
                                                                                } else {
                                                                                    if false {
                                                                                        v70 = "\n\t"
                                                                                        var v72 string = ""
                                                                                        _ = v72
                                                                                    }
                                                                                }
                                                                                _ = v70
                                                                            } else {
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    fmt.Println(v32)
                                                                                case "hello":
                                                                                    v58 = append(v58, (fi56 % (((fi59 % ((-1) | 1))) | 1)))
                                                                                case "":
                                                                                    if (v2 && ((v1 > v1) && (v2 && true))) {
                                                                                        switch 128 {
                                                                                        case 128:
                                                                                            _ = len(v4)
                                                                                        default:
                                                                                            v73 := GenS5{F0: 0}
                                                                                            _ = v73
                                                                                        }
                                                                                        _ = len(v58)
                                                                                        v74 := GenS4{F0: 874}
                                                                                        _ = v74
                                                                                    }
                                                                                }
                                                                                v75 := GenS2{F0: 1, F1: "foo bar", F2: 0}
                                                                                _ = v75
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v76 float64 = (float64(v30) + (float64((v3 + float32(float64(0.0)))) * 1.5))
                                                                        _ = v76
                                                                    }
                                                                    v30 += (-2.25 / 0.1)
                                                                    v77 := GenS7{F0: int16(v32), F1: ("a" > "a"), F2: (v32 / (((v32 | v32)) | 1))}
                                                                    v78 := GenS2{F0: -128, F1: "foo bar", F2: 4294967295}
                                                                    v3 -= ((0.0 - v30) / (float32(((-2.25 / (0.1 / float64((v3 * -2.25)))) + 0.0)) * ((1.5 + v3) * 1.5)))
                                                                    _ = v77
                                                                    _ = v78
                                                                } else {
                                                                    var v79 string = "\n\t"
                                                                    _ = v79
                                                                }
                                                                v80 := GenS5{F0: 42}
                                                                _ = fi59
                                                                _ = v80
                                                            }
                                                            v81 := GenS8{F0: fi56, F1: -1}
                                                            _ = v81
                                                        } else {
                                                            var v82 int32 = v1
                                                            _ = v82
                                                        }
                                                        for fi83 := 0; fi83 < 8; fi83++ {
                                                            v1 = (v1 & (-(v5)))
                                                            _ = fi83
                                                        }
                                                        _ = v58
                                                    } else {
                                                        switch 32767 {
                                                        case 9223372036854775807:
                                                            v1 += 2147483647
                                                        case -1:
                                                            v4[v29] = v29
                                                        case -627:
                                                            var v84 float64 = -2.25
                                                            _ = v84
                                                        default:
                                                            fmt.Println(v32, v5)
                                                        }
                                                        for fi85 := 0; fi85 < 1; fi85++ {
                                                            fmt.Println(v29, v1)
                                                            var v86 uint8 = 255
                                                            switch "hello" {
                                                            case "hello":
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v2 = v2
                                                                case "":
                                                                    delete(v4, v29)
                                                                default:
                                                                    fmt.Println(v32, fi85, fi56)
                                                                }
                                                            case "":
                                                                var v87 bool = (!v2)
                                                                _ = v87
                                                            case "a":
                                                                var v88 int16 = 897
                                                                _ = v88
                                                            default:
                                                                v89 := GenS9{F0: 42}
                                                                _ = v89
                                                            }
                                                            var v90 int = (^(fi85))
                                                            v2 = (v2 || (1e10 == 1.5))
                                                            if v2 {
                                                                v91 := v4[(v29 << ((v32) & 63))]
                                                                if false {
                                                                    v92 := GenS3{F0: v2}
                                                                    var v93 string = "hello"
                                                                    v5 = 0
                                                                    switch v93 {
                                                                    case "foo bar":
                                                                        v2 = false
                                                                    }
                                                                    _ = v92
                                                                    _ = v93
                                                                }
                                                                _ = v91
                                                            }
                                                            _ = fi85
                                                            _ = v86
                                                            _ = v90
                                                        }
                                                    }
                                                    if true {
                                                        var v94 uint = v32
                                                        _ = v94
                                                    }
                                                    v29 = uint32(v32)
                                                }
                                                var v95 int32 = v5
                                                v4[(v29 / ((2) | 1))] = 4294967295
                                                _ = v95
                                            }
                                        } else {
                                            var v96 bool = (1 <= 35)
                                            _ = len(v4)
                                            _ = v96
                                        }
                                        for fi97 := 0; fi97 < 18; fi97++ {
                                            v98 := GenS4{F0: 0}
                                            var v99 int32 = (v1 - 1)
                                            v5 -= (v99 >> (((^(v32))) & 63))
                                            _ = fi97
                                            _ = v98
                                            _ = v99
                                        }
                                        for fi100 := 0; fi100 < 18; fi100++ {
                                            v101 := []int{(fi56 & 127), (fi100 & fi56), (fi100 / ((-128) | 1))}
                                            var v102 float64 = 0.1
                                            v5 -= (v1 | -1)
                                            _ = len(v4)
                                            _ = fi100
                                            _ = v101
                                            _ = v102
                                        }
                                        _ = fi56
                                    }
                                    var v103 int = -9223372036854775808
                                    switch (v103 << ((genFn1(42, 113, (v32 | v32))) & 63)) {
                                    case -752:
                                        v32 += (v32 % ((2) | 1))
                                    }
                                    _ = v103
                                }
                            default:
                                var v104 int = -2147483648
                                _ = v104
                            }
                        }
                    }
                    if (-914 <= 4294967295) {
                        v105 := GenS5{F0: 2}
                        v29 = 247
                        v4[0] = v29
                        _ = v105
                    }
                    _ = v32
                } else {
                    var v106 uint = 98
                    _ = v106
                }
            }
            for fi107 := 0; fi107 < 11; fi107++ {
                v108 := GenS6{F0: 1000}
                if ((100 != 100) || v2) {
                    v109 := v4[v29]
                    v110 := []int{(fi107 * (fi107 & fi107))}
                    _ = v109
                    _ = v110
                }
                switch "a" {
                case "hello":
                    v5 = v1
                }
                _ = fi107
                _ = v108
            }
        } else {
            var v111 string = "hello"
            fmt.Println(v2, v29, v2)
            switch 4294967295 {
            case 9223372036854775807:
                v111 = ((v111 + v111) + v111)
            case 0:
                var v112 uint16 = 42
                _ = v112
            case -1:
                _ = len(v4)
            }
            _ = v111
        }
    }
    var v113 uint16 = 32767
    for fi114 := 0; fi114 < 20; fi114++ {
        v113 -= (v113 << ((65535) & 63))
        _ = fi114
    }
    v3 += v30
    delete(v4, 1)
    var v115 int64 = -1
    if ((("\n\t" + "hello") <= "hello") || false) {
        for fi116 := 0; fi116 < 15; fi116++ {
            v117 := GenS10{F0: 773, F1: "\n\t"}
            _ = fi116
            _ = v117
        }
        var v118 string = "hello"
        switch 892 {
        case 4294967295:
            if true {
                var v119 uint64 = 18446744073709551615
                v115 = -9223372036854775808
                for fi120 := 0; fi120 < 5; fi120++ {
                    switch (fi120 / (((fi120 / ((2147483647) | 1))) | 1)) {
                    case 0:
                        if ((!v2) && ((v115 != v115) || false)) {
                            v2 = (!(!false))
                        } else {
                            v121 := GenS10{F0: (v29 << ((65535) & 63)), F1: ""}
                            _ = v121
                        }
                    }
                    _ = fi120
                }
                _ = v119
            } else {
                for fi122 := 0; fi122 < 8; fi122++ {
                    v123 := []int{32767, fi122, fi122}
                    v124 := GenS11{F0: (v113 << ((65535) & 63))}
                    _ = fi122
                    _ = v123
                    _ = v124
                }
                v125 := GenS12{F0: v115, F1: 127, F2: uint8(v29)}
                if v2 {
                    var v126 uint8 = 127
                    var v127 float64 = (float64(((((1.5 + v30) - 1.5) / (float32(1.5) * v30)) * v3)) / (1e10 / 3.14159))
                    _ = v126
                    _ = v127
                } else {
                    switch ("" + v118) {
                    case "":
                        for fi128 := 0; fi128 < 16; fi128++ {
                            var v129 int8 = -128
                            v30 = ((0.0 * v3) + v3)
                            var v130 int8 = (v129 & (v129 - 42))
                            _ = fi128
                            _ = v129
                            _ = v130
                        }
                    case "a":
                        v2 = v2
                    case "hello":
                        v115 -= v115
                    }
                    var v131 string = (v118 + "hello")
                    _ = v131
                }
                _ = v125
            }
        }
        _ = v118
    }
    fmt.Println(v113, v1)
    v4[v29] = 655
    v1 = (v1 - v5)
    var v132 uint32 = 906
    fmt.Println(v113, v5)
    var v133 int64 = (v115 - v115)
    v134 := v4[(v29 / ((v29) | 1))]
    v135 := GenS6{F0: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v115:", v115)
    fmt.Println("v132:", v132)
    fmt.Println("v133:", v133)
    fmt.Println("v134:", v134)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v113
    _ = v115
    _ = v132
    _ = v133
    _ = v134
    _ = v135
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
