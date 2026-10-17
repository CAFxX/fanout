package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int16; F2 int8 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 string }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 string; F1 int8 }
type GenS14 struct { F0 string; F1 int }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int64; F1 int16; F2 uint64 }
type GenS17 struct { F0 uint8; F1 string }
type GenS18 struct { F0 int32; F1 uint8; F2 uint }
type GenS19 struct { F0 int32; F1 int64 }
func genFn0() uint64 {
    v4 := []int{0, -9223372036854775808, 9223372036854775807}
    if false {
        var v5 int32 = 150
        v5 = -208
        _ = v5
    }
    var v6 uint = 4294967295
    if false {
        switch "a" {
        case "hello":
            v6 = v6
        case "a":
            var v7 uint64 = 1
            _ = v7
        default:
            var v8 int64 = -1
            _ = v8
        }
        var v9 float32 = (0.1 + -2.25)
        _ = v9
    }
    _ = v4
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 int8, p1 uint, p2 int8) int8 {
    switch (("" + "hello") + "") {
    case "\n\t":
        p1 = (^(p1))
    default:
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS1{F0: (fi10 / ((fi10) | 1)), F1: p1, F2: false}
            switch "a" {
            case "":
                switch (-(fi10)) {
                case -1:
                    v12 := []int{fi10, (fi10 >> ((255) & 63)), (^(fi10)), -700}
                    _ = v12
                case 9223372036854775807:
                    switch -1 {
                    case 1:
                        var v13 int64 = 9223372036854775807
                        _ = v13
                    case 2147483647:
                        v14 := []int{}
                        _ = v14
                    }
                case 707:
                    for fi15 := 0; fi15 < 9; fi15++ {
                        if false {
                            var v16 uint = 2
                            if false {
                                v17 := []int{(fi15 - fi15), fi10}
                                _ = v17
                            } else {
                                fmt.Println(p2, p1)
                                p1 += 65535
                                fmt.Println(v16)
                            }
                            _ = v16
                        } else {
                            p1 -= p1
                            if ((((1.5 - 1e10) * (1e10 - -2.25)) / 0.1) <= 0.1) {
                                switch (fi15 - fi10) {
                                case -9223372036854775808:
                                    v18 := GenS1{F0: fi10, F1: p1, F2: (false || true)}
                                    _ = v18
                                default:
                                    var v19 int16 = -32768
                                    _ = v19
                                }
                                if true {
                                    var v20 bool = (!false)
                                    var v21 int = fi15
                                    var v22 int32 = 1
                                    v20 = v20
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                } else {
                                    switch 642 {
                                    case 42:
                                        if false {
                                            v23 := GenS2{F0: (p1 << ((0) & 63))}
                                            if false {
                                                var v24 uint = (p1 / ((p1) | 1))
                                                v25 := []int{fi10, (fi10 >> ((v24) & 63)), 127}
                                                if len(v25) > 0 {
                                                    v25[0] = (fi10 | 32767)
                                                }
                                                _ = v24
                                                _ = v25
                                            } else {
                                                p0 = 1
                                                var v26 int64 = 9223372036854775807
                                                v27 := []int{}
                                                _ = v26
                                                _ = v27
                                            }
                                            switch (fi10 ^ fi15) {
                                            case 32767:
                                                v28 := []int{fi15, -2147483648}
                                                v29 := GenS3{F0: 0, F1: 0, F2: 255}
                                                if false {
                                                    switch "" {
                                                    case "foo bar":
                                                        if (65535 <= 0) {
                                                            var v30 int32 = -1
                                                            _ = v30
                                                        }
                                                    default:
                                                        if (false && (true || false)) {
                                                            if (((float32(((1e10 + 3.14159) + 3.14159)) > ((-2.25 / 1.5) + -2.25)) && false) && false) {
                                                                p1 += p1
                                                                v31 := GenS4{F0: 1, F1: 1, F2: p2}
                                                                _ = v31
                                                            } else {
                                                                p2 = -1
                                                                fmt.Println(fi10)
                                                                p2 -= (^(p2))
                                                                var v32 uint16 = 856
                                                                if len(v28) > 0 {
                                                                    v28[0] = fi15
                                                                }
                                                                _ = v32
                                                            }
                                                        } else {
                                                            v33 := GenS5{F0: 761, F1: -1}
                                                            p0 = 100
                                                            _ = len(v28)
                                                            _ = v33
                                                        }
                                                    }
                                                    p0 = (p0 - 11)
                                                }
                                                _ = v28
                                                _ = v29
                                            default:
                                                p0 = p0
                                            }
                                            fmt.Println(p1)
                                            _ = v23
                                        } else {
                                            var v34 int16 = 822
                                            _ = v34
                                        }
                                    case 678:
                                        var v35 uint = 2
                                        _ = v35
                                    }
                                    p0 += 21
                                }
                                v36 := GenS2{F0: 956}
                                _ = v36
                            } else {
                                var v37 int64 = 42
                                p0 = (p2 ^ 0)
                                fmt.Println(p0)
                                _ = v37
                            }
                            v38 := GenS3{F0: 45, F1: p1, F2: 42}
                            _ = v38
                        }
                        p1 += 18446744073709551615
                        _ = fi15
                    }
                }
            case "\n\t":
                var v39 string = (("\n\t" + "\n\t") + ("foo bar" + "foo bar"))
                _ = v39
            case "a":
                switch (fi10 & (fi10 % ((int(p0)) | 1))) {
                case -514:
                    if (p0 >= 100) {
                        v40 := GenS4{F0: 0, F1: 42, F2: (p0 * (p2 * -128))}
                        v41 := GenS6{F0: fi10}
                        var v42 int32 = 42
                        _ = v40
                        _ = v41
                        _ = v42
                    } else {
                        if false {
                            v43 := GenS7{F0: 127}
                            _ = v43
                        }
                        var v44 uint64 = 18446744073709551615
                        _ = v44
                    }
                default:
                    if ((51 < -199) || false) {
                        v45 := []int{2147483647, 896, 9223372036854775807}
                        _ = v45
                    } else {
                        if false {
                            if (-339 <= 657) {
                                switch "" {
                                case "":
                                    var v46 uint16 = 65535
                                    _ = v46
                                case "a":
                                    v47 := []int{fi10}
                                    _ = v47
                                }
                            } else {
                                if (!((188 == 255) && (!(!(true && (!(!true))))))) {
                                    switch "foo bar" {
                                    case "":
                                        v48 := GenS8{F0: int16(p2), F1: ""}
                                        _ = v48
                                    }
                                    fmt.Println(p0, fi10)
                                }
                                v49 := GenS9{F0: "\n\t", F1: 127}
                                p2 += (p2 ^ (p2 % (((-(p0))) | 1)))
                                _ = v49
                            }
                            if true {
                                var v50 bool = false
                                for fi51 := 0; fi51 < 20; fi51++ {
                                    v52 := []int{(fi51 & fi51), (-(fi10)), (fi10 >> (((^(p1))) & 63))}
                                    if (42 != 819) {
                                        fmt.Println(p1, p2)
                                        var v53 uint8 = uint8(p1)
                                        _ = v53
                                    } else {
                                        _ = len(v52)
                                    }
                                    fmt.Println(p0, p1, p0)
                                    _ = fi51
                                    _ = v52
                                }
                                _ = v50
                            }
                        } else {
                            v54 := GenS2{F0: 1}
                            if (594 >= -757) {
                                v55 := []int{-824}
                                var v56 bool = (0.1 >= 3.14159)
                                for fi57 := 0; fi57 < 13; fi57++ {
                                    if v56 {
                                        var v58 bool = false
                                        v59 := GenS10{F0: (81 > 42)}
                                        if false {
                                            if (((1 != 0) || (2 >= 1)) && v56) {
                                                if ((-(p2)) <= p0) {
                                                    v60 := GenS8{F0: 42, F1: ("\n\t" + "")}
                                                    var v61 uint8 = 6
                                                    _ = v60
                                                    _ = v61
                                                }
                                            }
                                        } else {
                                            if (!v58) {
                                                p0 = (p2 / ((p2) | 1))
                                                var v62 float64 = (3.14159 * 0.0)
                                                switch -519 {
                                                case 4294967295:
                                                    var v63 int8 = (^(p0))
                                                    if v58 {
                                                        var v64 uint = (p1 % ((892) | 1))
                                                        var v65 uint8 = 42
                                                        _ = v64
                                                        _ = v65
                                                    } else {
                                                        if (!v56) {
                                                            var v66 int = 32767
                                                            v67 := GenS3{F0: 207, F1: (p1 - (p1 - (p1 & p1))), F2: 0}
                                                            var v68 uint32 = 2
                                                            _ = v66
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                    }
                                                    fmt.Println(v58, v62)
                                                    v55 = append(v55, fi57)
                                                    var v69 float32 = (((0.1 * ((1.5 / (3.14159 * 1e10)) * 1.5)) - 0.0) / 1.5)
                                                    p2 = -1
                                                    if v58 {
                                                        if (((fi57 - fi10) == (fi57 & (fi10 << ((p1) & 63)))) || (v69 == v69)) {
                                                            if (0 <= 549) {
                                                                if v58 {
                                                                    if len(v55) > 0 {
                                                                        v55[0] = (-(fi57))
                                                                    }
                                                                    switch ("" + ("a" + (("\n\t" + (("a" + ((("foo bar" + ("foo bar" + "")) + "foo bar") + (("a" + "\n\t") + (("hello" + "foo bar") + ("" + "\n\t"))))) + ("" + "\n\t"))) + ("foo bar" + "hello")))) {
                                                                    case "a":
                                                                        p1 = (p1 % ((p1) | 1))
                                                                    }
                                                                    if len(v55) > 0 {
                                                                        v55[0] = fi10
                                                                    }
                                                                }
                                                            } else {
                                                                if len(v55) > 0 {
                                                                    v55[0] = (fi10 & 42)
                                                                }
                                                                v63 -= (p2 ^ (v63 & p0))
                                                            }
                                                        } else {
                                                            v56 = (!true)
                                                            if v58 {
                                                                var v70 uint8 = 255
                                                                if len(v55) > 0 {
                                                                    v55[0] = fi57
                                                                }
                                                                _ = v70
                                                            } else {
                                                                if ((1 <= 0) || (255 > 65535)) {
                                                                    switch fi57 {
                                                                    case 0:
                                                                        v71 := GenS11{F0: 234}
                                                                        _ = v71
                                                                    case 127:
                                                                        v72 := GenS12{F0: (fi57 | (fi57 / ((-508) | 1)))}
                                                                        if (!(985 < 65535)) {
                                                                            p1 -= p1
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                var v73 float64 = v62
                                                                                _ = v73
                                                                            case "hello":
                                                                                _ = len(v55)
                                                                            }
                                                                        } else {
                                                                            var v74 uint32 = 1
                                                                            fmt.Println(fi10, p2)
                                                                            _ = v74
                                                                        }
                                                                        _ = v72
                                                                    }
                                                                } else {
                                                                    if len(v55) >= 2 {
                                                                        v55 = v55[1:len(v55)]
                                                                    }
                                                                    fmt.Println(fi10, p1, v69)
                                                                }
                                                                var v75 uint = p1
                                                                fmt.Println(v75, v58, p1)
                                                                if (42 > int32(p0)) {
                                                                    if (v58 || (2147483648 >= 2147483648)) {
                                                                        var v76 int64 = 0
                                                                        var v77 bool = v58
                                                                        _ = v76
                                                                        _ = v77
                                                                    }
                                                                } else {
                                                                    v55 = append(v55, (fi10 - (fi57 >> ((p1) & 63))))
                                                                    switch "hello" {
                                                                    case "":
                                                                        if len(v55) > 0 {
                                                                            v55[0] = (fi57 | (fi57 & (fi57 / ((fi57) | 1))))
                                                                        }
                                                                    default:
                                                                        if len(v55) > 0 {
                                                                            v55[0] = -2147483648
                                                                        }
                                                                    }
                                                                }
                                                                switch "" {
                                                                case "a":
                                                                    if (25 < -928) {
                                                                        var v78 string = ("" + "a")
                                                                        var v79 uint = (^(p1))
                                                                        v80 := GenS6{F0: (fi57 << ((v75) & 63))}
                                                                        var v81 bool = (255 <= 128)
                                                                        _ = v78
                                                                        _ = v79
                                                                        _ = v80
                                                                        _ = v81
                                                                    }
                                                                }
                                                                _ = v75
                                                            }
                                                            v82 := GenS13{F0: "", F1: (v63 / (((v63 ^ p0)) | 1))}
                                                            _ = v82
                                                        }
                                                    }
                                                    _ = len(v55)
                                                    _ = v63
                                                    _ = v69
                                                case 1:
                                                    fmt.Println(p1, p0, v56)
                                                case 0:
                                                    v62 += (v62 - 3.14159)
                                                    v55 = append(v55, 128)
                                                }
                                                if len(v55) > 0 {
                                                    v55[0] = 513
                                                }
                                                _ = v62
                                            } else {
                                                v55 = append(v55, -2147483648)
                                                v56 = ("hello" < ("a" + "a"))
                                                p0 += (p2 | -128)
                                            }
                                            p1 += (p1 | (p1 >> (((p1 % (((p1 - 65535)) | 1))) & 63)))
                                        }
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = fi57
                                }
                                _ = v55
                                _ = v56
                            }
                            for fi83 := 0; fi83 < 18; fi83++ {
                                v84 := []int{fi10, fi10, fi10}
                                p1 -= p1
                                var v85 int16 = 32767
                                _ = fi83
                                _ = v84
                                _ = v85
                            }
                            _ = v54
                        }
                        var v86 uint8 = 127
                        for fi87 := 0; fi87 < 2; fi87++ {
                            v88 := GenS14{F0: "a", F1: (fi87 >> ((1) & 63))}
                            var v89 uint = p1
                            _ = fi87
                            _ = v88
                            _ = v89
                        }
                        _ = v86
                    }
                }
            default:
                var v90 int16 = 1
                _ = v90
            }
            _ = fi10
            _ = v11
        }
    }
    var v91 int = 48
    _ = p0
    _ = p1
    _ = p2
    _ = v91
    return p0
}

func main() {
    v1 := []uint32{2147483647, 42, 4294967295, 2147483647}
    v2 := make(map[uint32]bool)
    var v3 uint8 = 127
    var v92 int32 = 468
    v92 += v92
    fmt.Println(v3)
    _ = len(v2)
    v3 += v3
    fmt.Println(v3, v92)
    switch (("\n\t" + "foo bar") + "hello") {
    case "\n\t":
        switch 9223372036854775807 {
        case 4294967295:
            v93 := GenS13{F0: "foo bar", F1: -1}
            _ = v93
        case 32768:
            if true {
                var v94 int = 339
                var v95 string = "foo bar"
                _ = v94
                _ = v95
            } else {
                v92 -= 1
                v96 := GenS15{F0: v92}
                v2[835] = (-915 == 9223372036854775807)
                _ = v96
            }
        case -2147483648:
            v92 += 0
        default:
            switch "a" {
            case "":
                v3 += (v3 % ((v3) | 1))
            case "foo bar":
                v97 := GenS6{F0: -1}
                _ = v97
            }
        }
    }
    fmt.Println(v92, v3)
    switch "a" {
    case "hello":
        for fi98 := 0; fi98 < 16; fi98++ {
            if (!true) {
                _ = len(v1)
            } else {
                var v99 string = "hello"
                fmt.Println(v3)
                _ = v99
            }
            _ = fi98
        }
    default:
        v100 := GenS16{F0: 973, F1: 32767, F2: 4294967295}
        _ = v100
    }
    if (true || false) {
        var v101 float64 = (0.0 + 3.14159)
        _ = v101
    } else {
        v92 = (v92 / ((v92) | 1))
    }
    if (4294967295 > 1) {
        switch -2147483648 {
        case 9223372036854775807:
            if (!false) {
                for fi102 := 0; fi102 < 2; fi102++ {
                    fmt.Println(fi102, v92)
                    _ = fi102
                }
            } else {
                var v103 int32 = v92
                _ = v103
            }
        case 32767:
            switch "foo bar" {
            case "a":
                for fi104 := 0; fi104 < 5; fi104++ {
                    v3 = v3
                    _ = fi104
                }
            case "hello":
                var v105 uint16 = 255
                _ = v105
            case "foo bar":
                var v106 uint32 = 358
                _ = v106
            default:
                var v107 float64 = 3.14159
                _ = v107
            }
        }
        var v108 uint = 42
        for fi109 := 0; fi109 < 10; fi109++ {
            v3 += 242
            for fi110 := 0; fi110 < 9; fi110++ {
                var v111 uint32 = 2
                if len(v1) > 0 {
                    v1[0] = (v111 * (v111 | (v111 / ((v111) | 1))))
                }
                var v112 int8 = genFn1(105, uint(v3), 126)
                switch ("hello" + (("foo bar" + "\n\t") + "a")) {
                case "\n\t":
                    var v113 int8 = 65
                    _ = v113
                case "foo bar":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if true {
                        _ = len(v1)
                    }
                    if ((^(v111)) != 2147483647) {
                        v3 += 151
                        v108 += 745
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            var v114 uint16 = 65535
                            var v115 uint16 = 496
                            if false {
                                v115 -= v115
                            } else {
                                var v116 string = "\n\t"
                                v115 += (^(v115))
                                v117 := GenS15{F0: -2147483648}
                                _ = v116
                                _ = v117
                            }
                            _ = v114
                            _ = v115
                        case "hello":
                            v3 -= v3
                        case "":
                            var v118 uint8 = (v3 / (((^(v3))) | 1))
                            _ = v118
                        default:
                            if false {
                                switch fi110 {
                                case -9223372036854775808:
                                    v112 = genFn1(97, (^(v108)), v112)
                                case 959:
                                    if (genFn0() > 1) {
                                        v3 += (v3 + (v3 ^ (v3 / (((v3 - (v3 % ((uint8(v108)) | 1)))) | 1))))
                                        var v119 float32 = 1e10
                                        _ = v119
                                    } else {
                                        if ((((("foo bar" + "hello") + "\n\t") + ("a" + "a")) + ("" + "foo bar")) < "") {
                                            v3 += 100
                                            var v120 int = fi110
                                            _ = len(v1)
                                            _ = v120
                                        } else {
                                            if true {
                                                v112 += genFn1(int8(v108), 18446744073709551615, v112)
                                                v92 += v92
                                                if len(v1) > 0 {
                                                    v1[0] = (^(v111))
                                                }
                                            }
                                        }
                                        v112 -= 78
                                    }
                                    if true {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v108 += 2
                                        }
                                        if (!true) {
                                            fmt.Println(v3, v112)
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        var v121 float32 = 0.0
                                        _ = v121
                                    }
                                default:
                                    switch 42 {
                                    case -9223372036854775808:
                                        if (!(!((v111 >= 478) && false))) {
                                            var v122 int = -1
                                            fmt.Println(fi110, v108)
                                            var v123 int8 = v112
                                            _ = v122
                                            _ = v123
                                        }
                                    case 4294967295:
                                        fmt.Println(v92)
                                    }
                                }
                                if (true && (0 != 9223372036854775807)) {
                                    if false {
                                        var v124 uint8 = v3
                                        _ = v124
                                    } else {
                                        var v125 uint32 = 2147483647
                                        var v126 uint64 = 2
                                        _ = v125
                                        _ = v126
                                    }
                                    v1 = append(v1, (v111 | v111))
                                } else {
                                    if (false || true) {
                                        v127 := GenS1{F0: fi110, F1: 276, F2: false}
                                        _ = v127
                                    }
                                    v128 := GenS9{F0: "hello", F1: 52}
                                    v108 -= uint(v111)
                                    _ = v128
                                }
                                var v129 uint16 = 2
                                fmt.Println(v108, v108, v3)
                                var v130 int8 = (v112 - v112)
                                if false {
                                    v3 = v3
                                    if ((!(true && (-97 <= -9223372036854775808))) && true) {
                                        v131 := GenS9{F0: "hello", F1: uint8(v129)}
                                        v132 := v2[v111]
                                        fmt.Println(v132, v129)
                                        _ = v131
                                        _ = v132
                                    } else {
                                        var v133 float64 = ((3.14159 * 3.14159) + 1e10)
                                        var v134 int = fi110
                                        var v135 uint8 = (^(v3))
                                        if false {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        } else {
                                            v130 = 127
                                            v130 = (v112 << (((^(v108))) & 63))
                                            v136 := GenS8{F0: -357, F1: ""}
                                            if ((v108 - (^(v108))) == (v108 & (v108 >> ((0) & 63)))) {
                                                v137 := GenS17{F0: v135, F1: "foo bar"}
                                                _ = v137
                                            }
                                            switch 127 {
                                            case 9223372036854775807:
                                                v111 -= 821
                                            case 1:
                                                fmt.Println(v108, v134, v92)
                                            default:
                                                var v138 int32 = (v92 & (v92 * (-(v92))))
                                                _ = v138
                                            }
                                            _ = v136
                                        }
                                        _ = v133
                                        _ = v134
                                        _ = v135
                                    }
                                }
                                switch (^(fi110)) {
                                case -860:
                                    var v139 int64 = -1
                                    switch (fi110 - fi110) {
                                    case 518:
                                        fmt.Println(v139, v130)
                                    case -345:
                                        var v140 int16 = -484
                                        _ = v140
                                    case 32768:
                                        v141 := GenS18{F0: (v92 / ((1) | 1)), F1: v3, F2: 42}
                                        _ = v141
                                    default:
                                        if true {
                                            v1 = append(v1, v111)
                                            v3 = (v3 * (v3 | v3))
                                            fmt.Println(v108, v3, v92)
                                        } else {
                                            var v142 int16 = 960
                                            v111 = (v111 * (v111 % (((v111 ^ (v111 % ((v111) | 1)))) | 1)))
                                            fmt.Println(v130, v111, v142)
                                            _ = v142
                                        }
                                    }
                                    var v143 int64 = -1
                                    _ = v139
                                    _ = v143
                                case 629:
                                    var v144 string = (("" + ((("foo bar" + "hello") + "\n\t") + "")) + "foo bar")
                                    _ = v144
                                default:
                                    if false {
                                        var v145 int16 = 1
                                        _ = v145
                                    }
                                }
                                fmt.Println(fi109, v111, v111)
                                _ = v129
                                _ = v130
                            } else {
                                switch (fi110 ^ (fi110 + 639)) {
                                case 4294967295:
                                    if ("a" > "hello") {
                                        if false {
                                            fmt.Println(v112, fi110)
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v146 uint = (v108 & 255)
                                                _ = v146
                                            case "hello":
                                                v3 -= v3
                                            case "":
                                                v3 += 255
                                            }
                                            fmt.Println(v3, v111)
                                        } else {
                                            var v147 int32 = 2147483647
                                            v148 := GenS10{F0: ("\n\t" > "foo bar")}
                                            var v149 uint16 = 255
                                            _ = v147
                                            _ = v148
                                            _ = v149
                                        }
                                        fmt.Println(v111, v92)
                                    }
                                    fmt.Println(v3)
                                    var v150 int = int(v108)
                                    _ = v150
                                default:
                                    if (("a" < "foo bar") && (3.14159 < -2.25)) {
                                        var v151 uint = 84
                                        var v152 float64 = float64((1e10 - 0.0))
                                        _ = v151
                                        _ = v152
                                    }
                                }
                            }
                        }
                        v153 := GenS13{F0: "\n\t", F1: v112}
                        var v154 int64 = 772
                        _ = v153
                        _ = v154
                    }
                case "hello":
                    v3 += 127
                }
                _ = fi110
                _ = v111
                _ = v112
            }
            var v155 int8 = 64
            _ = fi109
            _ = v155
        }
        _ = v108
    } else {
        v156 := GenS17{F0: v3, F1: "\n\t"}
        switch "\n\t" {
        case "a":
            switch ("foo bar" + "\n\t") {
            case "":
                _ = len(v1)
            case "a":
                if false {
                    var v157 string = "hello"
                    var v158 int16 = 42
                    v159 := GenS10{F0: false}
                    _ = v157
                    _ = v158
                    _ = v159
                }
            case "\n\t":
                var v160 string = "\n\t"
                _ = v160
            default:
                for fi161 := 0; fi161 < 19; fi161++ {
                    if (2147483648 != 857) {
                        fmt.Println(v92, v92)
                        v162 := GenS8{F0: 1, F1: ("hello" + "a")}
                        v92 += (v92 + v92)
                        _ = v162
                    }
                    v1 = append(v1, 2)
                    _ = fi161
                }
            }
        case "\n\t":
            var v163 uint = 255
            _ = v163
        default:
            var v164 uint32 = 2147483647
            _ = v164
        }
        v92 = (v92 & v92)
        _ = v156
    }
    v3 = v3
    fmt.Println(v3, v3)
    if ("\n\t" <= "foo bar") {
        v165 := GenS7{F0: 1}
        _ = v165
    }
    for fi166 := 0; fi166 < 2; fi166++ {
        v167 := GenS19{F0: v92, F1: 42}
        _ = fi166
        _ = v167
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v92
}
