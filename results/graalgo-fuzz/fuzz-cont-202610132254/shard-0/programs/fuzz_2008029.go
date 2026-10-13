package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int16; F2 bool }
type GenS3 struct { F0 int; F1 int; F2 uint32 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 bool; F2 uint64 }
type GenS7 struct { F0 string; F1 uint; F2 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint; F1 int; F2 uint8 }
type GenS10 struct { F0 int16; F1 uint16; F2 int16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS13 struct { F0 uint8; F1 int16 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int16; F1 uint }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 int16; F1 int64; F2 bool }
func genFn0(p0 int, p1 int) uint32 {
    var v5 int = (p0 * (p1 >> ((0) & 63)))
    p1 = (v5 % (((p0 + (v5 ^ 1))) | 1))
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint16 = 42
        if (594 < 32767) {
            v8 := GenS1{F0: 42}
            if (p1 == (-(p0))) {
                switch p1 {
                case 168:
                    v9 := []int{}
                    _ = v9
                case 4294967295:
                    p0 = (^(v5))
                case 32768:
                    if true {
                        for fi10 := 0; fi10 < 13; fi10++ {
                            v11 := GenS1{F0: 127}
                            _ = fi10
                            _ = v11
                        }
                    }
                }
                var v12 uint64 = 297
                _ = v12
            }
            v13 := []int{(v5 % (((fi6 + p1)) | 1)), -896, (fi6 | 127)}
            _ = v13
            _ = v8
        }
        p0 += v5
        _ = fi6
        _ = v7
    }
    if ((9223372036854775807 > -1) || ((667 >= 432) || (42 <= 888))) {
        v14 := []int{}
        for fi15 := 0; fi15 < 10; fi15++ {
            v5 = (v5 ^ (^(v5)))
            _ = fi15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint32 = 2147483648
            v18 := GenS2{F0: 1, F1: -861, F2: (!(1 <= -128))}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v14
    }
    p1 = v5
    _ = p0
    _ = p1
    _ = v5
    return 2
}

func main() {
    var v1 bool = (true || (0 <= -9223372036854775808))
    v2 := make(map[int]int)
    v3 := []string{"a"}
    v4 := make(map[string]uint32)
    var v19 string = ("a" + "")
    v20 := v4[v19]
    fmt.Println(v19, v20, v19)
    var v21 int32 = -1
    v1 = true
    var v22 int = -725
    var v23 int8 = 1
    v24 := v4[("hello" + (v19 + (("foo bar" + v19) + "foo bar")))]
    for fi25 := 0; fi25 < 18; fi25++ {
        var v26 string = (v19 + "")
        v27 := GenS2{F0: -735, F1: -1, F2: (366 < v24)}
        _ = fi25
        _ = v26
        _ = v27
    }
    v23 += (v23 / (((-(v23))) | 1))
    v28 := GenS3{F0: (v22 % (((v22 & (v22 & v22))) | 1)), F1: 1, F2: (v20 + (v24 ^ v24))}
    v29 := GenS4{F0: 90, F1: 1}
    v21 -= (v21 & v21)
    switch "foo bar" {
    case "hello":
        fmt.Println(v1, v24, v21)
    case "":
        _ = len(v2)
    default:
        v22 -= v22
    }
    var v30 bool = (false && true)
    fmt.Println(v24)
    if (0 < -392) {
        switch v22 {
        case -219:
            var v31 int16 = 0
            _ = v31
        case -2147483648:
            v30 = v1
        case 128:
            switch v22 {
            case 2147483647:
                if (1 <= 898) {
                    v32 := GenS4{F0: 205, F1: 91}
                    fmt.Println(v22)
                    for fi33 := 0; fi33 < 3; fi33++ {
                        switch v22 {
                        case 127:
                            v34 := GenS4{F0: 255, F1: 0}
                            _ = v34
                        case 24:
                            var v35 float32 = 1e10
                            _ = v35
                        case 32768:
                            for fi36 := 0; fi36 < 1; fi36++ {
                                var v37 int16 = -552
                                switch (fi33 + v22) {
                                case 2147483647:
                                    v23 = (v23 % (((v23 ^ (v23 ^ (v23 % ((0) | 1))))) | 1))
                                case -264:
                                    if (2 > 4294967295) {
                                        if (v21 <= -1) {
                                            if v30 {
                                                v19 = (("" + v19) + v19)
                                                v38 := GenS1{F0: v23}
                                                v39 := GenS5{F0: -341}
                                                _ = v38
                                                _ = v39
                                            } else {
                                                if true {
                                                    var v40 int16 = v37
                                                    v22 = int(v21)
                                                    _ = v40
                                                } else {
                                                    v22 = (-(fi33))
                                                    v41 := GenS6{F0: v21, F1: (int8(v21) < -1), F2: 42}
                                                    var v42 float64 = 1e10
                                                    _ = v41
                                                    _ = v42
                                                }
                                                var v43 float64 = (-2.25 / 1e10)
                                                var v44 float64 = v43
                                                _ = v43
                                                _ = v44
                                            }
                                        } else {
                                            delete(v2, (fi36 * 32767))
                                            var v45 uint = 255
                                            fmt.Println(v20, v1)
                                            _ = v45
                                        }
                                        var v46 float64 = 0.0
                                        _ = v46
                                    } else {
                                        var v47 int64 = 42
                                        fmt.Println(v22, v23, v23)
                                        _ = v47
                                    }
                                case 32767:
                                    v20 = (v20 + (v24 + (^(v20))))
                                default:
                                    switch 928 {
                                    case 4294967295:
                                        if (v1 && false) {
                                            var v48 uint8 = 255
                                            _ = v48
                                        } else {
                                            v24 += (v20 - (v24 - (v20 + 0)))
                                        }
                                    }
                                }
                                fmt.Println(fi33, v20)
                                _ = fi36
                                _ = v37
                            }
                        }
                        var v49 float32 = 0.0
                        _ = fi33
                        _ = v49
                    }
                    _ = v32
                } else {
                    v23 = v23
                    for fi50 := 0; fi50 < 10; fi50++ {
                        var v51 uint = 255
                        for fi52 := 0; fi52 < 19; fi52++ {
                            v19 = v19
                            _ = fi52
                        }
                        var v53 uint16 = 1
                        _ = fi50
                        _ = v51
                        _ = v53
                    }
                    v54 := GenS6{F0: v21, F1: (3.14159 < 0.0), F2: uint64(v20)}
                    _ = v54
                }
            case 4294967295:
                v55 := v2[596]
                _ = v55
            }
        default:
            for fi56 := 0; fi56 < 5; fi56++ {
                var v57 int32 = v21
                if ((32767 != 65535) || (953 > 4294967295)) {
                    var v58 string = v19
                    fmt.Println(v30, fi56, v19)
                    if v30 {
                        delete(v2, -9223372036854775808)
                        var v59 int32 = v21
                        _ = v59
                    }
                    _ = v58
                }
                _ = fi56
                _ = v57
            }
        }
    } else {
        for fi60 := 0; fi60 < 16; fi60++ {
            if ((!(-32768 == 0)) && v1) {
                for fi61 := 0; fi61 < 5; fi61++ {
                    v62 := GenS7{F0: "hello", F1: 0, F2: 42}
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(fi60)
                    case "a":
                        switch v19 {
                        case "hello":
                            switch -2147483648 {
                            case 0:
                                var v63 int8 = 42
                                _ = v63
                            case 127:
                                v20 += genFn0(fi61, (fi61 ^ (fi60 * (-(fi60)))))
                            }
                        case "a":
                            var v64 int16 = 42
                            _ = v64
                        case "":
                            v23 = v23
                        default:
                            v21 += 42
                        }
                    default:
                        switch (^(fi61)) {
                        case 2147483647:
                            fmt.Println(v1)
                        }
                    }
                    _ = fi61
                    _ = v62
                }
                if true {
                    v23 += v23
                } else {
                    v20 -= (v20 + v20)
                    var v65 int8 = (^(v23))
                    for fi66 := 0; fi66 < 8; fi66++ {
                        switch v19 {
                        case "a":
                            var v67 int16 = 1000
                            _ = v67
                        case "hello":
                            switch 0 {
                            case 32767:
                                switch 999 {
                                case 127:
                                    var v68 int8 = (v65 % (((-(v23))) | 1))
                                    _ = v68
                                }
                            default:
                                var v69 float32 = ((0.1 - (1e10 + -2.25)) * 0.0)
                                _ = v69
                            }
                        }
                        v70 := v2[(v22 / ((128) | 1))]
                        v71 := GenS4{F0: 0, F1: 243}
                        fmt.Println(fi60, fi66)
                        switch 32768 {
                        case 32768:
                            if true {
                                fmt.Println(v23)
                                v22 = fi66
                                var v72 float32 = 0.1
                                v73 := GenS8{F0: 1}
                                var v74 bool = (-1 <= 1)
                                if ((float64((v72 / (0.0 - (v72 + (float32((3.14159 / (1.5 / (((((1.5 + 3.14159) + (((0.0 / 1e10) - (3.14159 / 1.0)) - -2.25)) / float64(v72)) + -2.25) + 0.1)))) + (((1.5 + 1e10) * ((v72 - -2.25) * float32(-2.25))) - 1e10)))))) / 1.0) != 3.14159) {
                                    v30 = false
                                    v19 = ""
                                    if len(v3) > 0 {
                                        v3[0] = v19
                                    }
                                    var v75 int = (v70 ^ -9223372036854775808)
                                    _ = v75
                                }
                                v30 = v74
                                var v76 uint8 = 1
                                v77 := GenS2{F0: 0, F1: 0, F2: v30}
                                _ = v72
                                _ = v73
                                _ = v74
                                _ = v76
                                _ = v77
                            }
                        case 0:
                            if ((v24 * v20) < (v24 | genFn0(fi60, v22))) {
                                switch v19 {
                                case "a":
                                    v78 := GenS9{F0: 384, F1: (v70 & fi66), F2: 110}
                                    _ = v78
                                case "foo bar":
                                    v3 = append(v3, (((((((v19 + "\n\t") + (("a" + "hello") + "hello")) + "a") + (v19 + v19)) + ((v19 + "foo bar") + v19)) + "a") + v19))
                                    if v1 {
                                        v79 := GenS10{F0: 42, F1: 42, F2: -1}
                                        v24 += (v24 ^ (v24 + (v24 << ((51) & 63))))
                                        v20 = v20
                                        _ = v79
                                    } else {
                                        if v1 {
                                            var v80 int16 = -32768
                                            if v1 {
                                                v24 += (v20 / ((genFn0(42, 128)) | 1))
                                                if v1 {
                                                    if v1 {
                                                        if len(v3) > 0 {
                                                            v3[0] = v19
                                                        }
                                                        v81 := GenS11{F0: 62}
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v81
                                                    }
                                                    var v82 uint8 = 128
                                                    delete(v4, "")
                                                    _ = v82
                                                }
                                                fmt.Println(v30, v1)
                                            } else {
                                                var v83 float32 = 1e10
                                                v84 := GenS3{F0: v70, F1: 0, F2: (^(v24))}
                                                v85 := GenS10{F0: 1000, F1: 42, F2: v80}
                                                var v86 uint64 = 18446744073709551615
                                                _ = v83
                                                _ = v84
                                                _ = v85
                                                _ = v86
                                            }
                                            _ = v80
                                        }
                                    }
                                default:
                                    if (v30 || true) {
                                        fmt.Println(v30)
                                        var v87 int64 = 236
                                        switch "a" {
                                        case "\n\t":
                                            switch fi60 {
                                            case 32767:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    fmt.Println(v20, v19, v19)
                                                case "hello":
                                                    v88 := GenS12{F0: v20, F1: 65535, F2: (v21 / ((v21) | 1))}
                                                    _ = v88
                                                default:
                                                    switch ("a" + "\n\t") {
                                                    case "\n\t":
                                                        v89 := GenS13{F0: 2, F1: 623}
                                                        _ = v89
                                                    case "a":
                                                        fmt.Println(v1, v1, v30)
                                                    case "foo bar":
                                                        switch fi66 {
                                                        case 2147483647:
                                                            if (!false) {
                                                                v90 := GenS11{F0: 18446744073709551615}
                                                                v91 := GenS7{F0: (v19 + v19), F1: uint(v24), F2: v21}
                                                                switch v19 {
                                                                case "foo bar":
                                                                    v23 = 127
                                                                case "":
                                                                    if (!v30) {
                                                                        if (!(v87 <= 9223372036854775807)) {
                                                                            v1 = (!v1)
                                                                        }
                                                                        fmt.Println(v21, v22, v70)
                                                                        v20 -= (^(v20))
                                                                    }
                                                                case "a":
                                                                    if v30 {
                                                                        var v92 string = v19
                                                                        v92 = ("\n\t" + ((v19 + v19) + v92))
                                                                        _ = v92
                                                                    } else {
                                                                        fmt.Println(v70, fi66)
                                                                    }
                                                                }
                                                                var v93 int8 = 4
                                                                v94 := GenS10{F0: 1, F1: 255, F2: 32767}
                                                                _ = v90
                                                                _ = v91
                                                                _ = v93
                                                                _ = v94
                                                            } else {
                                                                switch v19 {
                                                                case "\n\t":
                                                                    var v95 string = v19
                                                                    v65 = v65
                                                                    _ = v95
                                                                case "foo bar":
                                                                    var v96 float64 = (-2.25 + (1.5 / (0.1 - 1.5)))
                                                                    _ = v96
                                                                }
                                                                if v30 {
                                                                    v1 = ("foo bar" == v19)
                                                                    switch "" {
                                                                    case "hello":
                                                                        v24 -= v24
                                                                    case "a":
                                                                        switch -257 {
                                                                        case -2147483648:
                                                                            v97 := GenS5{F0: (v22 - (v22 / ((int(v65)) | 1)))}
                                                                            _ = v97
                                                                        case 0:
                                                                            if (!v30) {
                                                                                switch 2147483647 {
                                                                                case 42:
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                }
                                                                                switch v19 {
                                                                                case "a":
                                                                                    if v30 {
                                                                                        if (int32(v24) == (v21 / ((0) | 1))) {
                                                                                            if true {
                                                                                                fmt.Println(v19, v19, v24)
                                                                                            } else {
                                                                                                v87 -= (v87 >> ((42) & 63))
                                                                                                v98 := GenS12{F0: (v24 >> ((4294967295) & 63)), F1: 42, F2: v21}
                                                                                                v99 := GenS9{F0: 18446744073709551615, F1: fi66, F2: 128}
                                                                                                _ = v98
                                                                                                _ = v99
                                                                                            }
                                                                                            v3 = append(v3, v19)
                                                                                        }
                                                                                        switch fi60 {
                                                                                        case 662:
                                                                                            _ = len(v3)
                                                                                        }
                                                                                        if v30 {
                                                                                            if v1 {
                                                                                                if v30 {
                                                                                                    if (v24 > 754) {
                                                                                                        if (-128 == v23) {
                                                                                                            fmt.Println(v87, fi60)
                                                                                                            var v100 string = (v19 + "")
                                                                                                            switch v70 {
                                                                                                            case -1:
                                                                                                                delete(v4, ((("\n\t" + "") + "foo bar") + ("foo bar" + v19)))
                                                                                                            case -81:
                                                                                                                fmt.Println(v21, fi60)
                                                                                                            case -74:
                                                                                                                var v101 string = "foo bar"
                                                                                                                _ = v101
                                                                                                            default:
                                                                                                                switch "a" {
                                                                                                                case "":
                                                                                                                    _ = len(v4)
                                                                                                                default:
                                                                                                                    fmt.Println(v87, v70, fi60)
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v100
                                                                                                        }
                                                                                                        fmt.Println(v87)
                                                                                                        var v102 int64 = (v87 + v87)
                                                                                                        _ = v102
                                                                                                    } else {
                                                                                                        switch "\n\t" {
                                                                                                        case "hello":
                                                                                                            v103 := GenS14{F0: (!(356 >= -1))}
                                                                                                            v24 = v24
                                                                                                            _ = v103
                                                                                                        case "":
                                                                                                            var v104 int8 = (v23 * v23)
                                                                                                            _ = v104
                                                                                                        }
                                                                                                    }
                                                                                                    var v105 uint16 = 42
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    _ = v105
                                                                                                } else {
                                                                                                    if ((!(100 < (v23 * v65))) && ("hello" == v19)) {
                                                                                                        v65 += v23
                                                                                                        var v106 float32 = (1.5 * 1.5)
                                                                                                        if (623 >= 18446744073709551615) {
                                                                                                            if v1 {
                                                                                                                fmt.Println(v23, v21, v70)
                                                                                                                v22 += (v70 | (v70 * fi66))
                                                                                                                var v107 uint8 = 42
                                                                                                                _ = v107
                                                                                                            } else {
                                                                                                                var v108 uint = 999
                                                                                                                var v109 int64 = v87
                                                                                                                _ = v108
                                                                                                                _ = v109
                                                                                                            }
                                                                                                            switch "a" {
                                                                                                            case "foo bar":
                                                                                                                v24 = (v24 % ((genFn0(fi60, 42)) | 1))
                                                                                                                switch fi60 {
                                                                                                                case -9223372036854775808:
                                                                                                                    v21 = v21
                                                                                                                case -2147483648:
                                                                                                                    fmt.Println(v24)
                                                                                                                }
                                                                                                            case "a":
                                                                                                                v1 = (!v1)
                                                                                                            default:
                                                                                                                delete(v2, 128)
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v110 int8 = (v65 + v23)
                                                                                                            fmt.Println(v23, v24)
                                                                                                            _ = v110
                                                                                                        }
                                                                                                        _ = v106
                                                                                                    } else {
                                                                                                        v111 := GenS14{F0: v30}
                                                                                                        _ = v111
                                                                                                    }
                                                                                                    v112 := GenS9{F0: 1, F1: (-(v22)), F2: 0}
                                                                                                    v30 = (3.14159 > 1e10)
                                                                                                    _ = v112
                                                                                                }
                                                                                                switch "a" {
                                                                                                case "a":
                                                                                                    var v113 bool = (195 >= 174)
                                                                                                    _ = v113
                                                                                                default:
                                                                                                    v23 = 50
                                                                                                    if false {
                                                                                                        var v114 uint32 = v24
                                                                                                        _ = v114
                                                                                                    } else {
                                                                                                        v1 = v30
                                                                                                        if v1 {
                                                                                                            var v115 uint32 = genFn0(v22, 32768)
                                                                                                            _ = v115
                                                                                                        } else {
                                                                                                            switch v22 {
                                                                                                            case -11:
                                                                                                                if ((3.14159 - ((-2.25 + ((1e10 / float32(-2.25)) / 1.0)) * float32(((((0.1 + 1e10) + -2.25) - 1e10) + 3.14159)))) >= -2.25) {
                                                                                                                    delete(v4, v19)
                                                                                                                }
                                                                                                            default:
                                                                                                                v1 = (!((float64(1e10) + 1.5) < ((float64(float32(float64(1.5))) / 1.0) / 3.14159)))
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    var v116 int8 = v23
                                                                                                    fmt.Println(v30, v30, v30)
                                                                                                    _ = v116
                                                                                                }
                                                                                            } else {
                                                                                                v117 := GenS15{F0: (v23 >> ((175) & 63))}
                                                                                                v118 := GenS6{F0: v21, F1: true, F2: 0}
                                                                                                var v119 uint64 = 712
                                                                                                _ = v117
                                                                                                _ = v118
                                                                                                _ = v119
                                                                                            }
                                                                                            switch "foo bar" {
                                                                                            case "\n\t":
                                                                                                delete(v4, "\n\t")
                                                                                            case "hello":
                                                                                                v24 += 2
                                                                                            case "a":
                                                                                                var v120 uint32 = (v20 / ((2) | 1))
                                                                                                _ = v120
                                                                                            }
                                                                                            var v121 int8 = v23
                                                                                            _ = v121
                                                                                        } else {
                                                                                            fmt.Println(v22, fi60)
                                                                                            fmt.Println(v1, v23)
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    var v122 float32 = (0.1 * (-2.25 + -2.25))
                                                                                    _ = v122
                                                                                default:
                                                                                    if v1 {
                                                                                        v30 = true
                                                                                    }
                                                                                }
                                                                                if (42 >= 846) {
                                                                                    v87 -= (-(v87))
                                                                                    var v123 uint = 18446744073709551615
                                                                                    _ = v123
                                                                                }
                                                                            } else {
                                                                                if (0 == v87) {
                                                                                    var v124 int32 = v21
                                                                                    _ = v124
                                                                                }
                                                                                v1 = (127 < v65)
                                                                            }
                                                                        default:
                                                                            var v125 uint64 = 644
                                                                            _ = v125
                                                                        }
                                                                        fmt.Println(v19, v1)
                                                                        switch (("hello" + "\n\t") + v19) {
                                                                        case "hello":
                                                                            delete(v2, (fi66 | (v22 & (fi60 + v22))))
                                                                        case "":
                                                                            var v126 uint16 = 356
                                                                            _ = v126
                                                                        case "\n\t":
                                                                            fmt.Println(v21)
                                                                        }
                                                                        var v127 int = int(v21)
                                                                        var v128 int = 2147483647
                                                                        if (float64((0.0 / 1.0)) > 1e10) {
                                                                            if v30 {
                                                                                var v129 int = fi60
                                                                                if true {
                                                                                    var v130 uint8 = 0
                                                                                    switch v19 {
                                                                                    case "":
                                                                                        var v131 int16 = 0
                                                                                        _ = v131
                                                                                    case "hello":
                                                                                        switch -983 {
                                                                                        case 9223372036854775807:
                                                                                            v128 = v129
                                                                                        case 42:
                                                                                            if false {
                                                                                                switch v19 {
                                                                                                case "a":
                                                                                                    v132 := GenS16{F0: -266}
                                                                                                    _ = v132
                                                                                                case "hello":
                                                                                                    var v133 int16 = -80
                                                                                                    _ = v133
                                                                                                default:
                                                                                                    v87 -= v87
                                                                                                }
                                                                                                v24 = 512
                                                                                                if v1 {
                                                                                                    if v1 {
                                                                                                        var v134 string = ""
                                                                                                        var v135 string = v19
                                                                                                        _ = v134
                                                                                                        _ = v135
                                                                                                    }
                                                                                                } else {
                                                                                                    var v136 uint32 = genFn0(fi60, (v127 + fi66))
                                                                                                    _ = v136
                                                                                                }
                                                                                            }
                                                                                        case -741:
                                                                                            switch v19 {
                                                                                            case "foo bar":
                                                                                                var v137 uint32 = uint32(v130)
                                                                                                if v30 {
                                                                                                    if (0 > -771) {
                                                                                                        var v138 int64 = -9223372036854775808
                                                                                                        if false {
                                                                                                            v3 = append(v3, v19)
                                                                                                            v139 := GenS17{F0: (1 >= 42)}
                                                                                                            _ = v139
                                                                                                        } else {
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = ("hello" + v19)
                                                                                                            }
                                                                                                        }
                                                                                                        if v30 {
                                                                                                            v87 = v87
                                                                                                        } else {
                                                                                                            switch v19 {
                                                                                                            case "":
                                                                                                                var v140 int8 = v23
                                                                                                                fmt.Println(v70, fi60, v127)
                                                                                                                _ = v140
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v138
                                                                                                    } else {
                                                                                                        v87 += (v87 << ((2) & 63))
                                                                                                        v141 := GenS18{F0: -400, F1: 942}
                                                                                                        fmt.Println(v23, v65, v1)
                                                                                                        var v142 int64 = 843
                                                                                                        fmt.Println(fi66)
                                                                                                        _ = v141
                                                                                                        _ = v142
                                                                                                    }
                                                                                                    var v143 uint = 4294967295
                                                                                                    switch v22 {
                                                                                                    case 127:
                                                                                                        var v144 uint16 = 854
                                                                                                        _ = v144
                                                                                                    case 471:
                                                                                                        v145 := v2[(fi66 << ((v143) & 63))]
                                                                                                        _ = v145
                                                                                                    default:
                                                                                                        var v146 int16 = -532
                                                                                                        _ = v146
                                                                                                    }
                                                                                                    _ = v143
                                                                                                }
                                                                                                _ = v137
                                                                                            case "\n\t":
                                                                                                fmt.Println(v65, v21)
                                                                                            }
                                                                                            if (v30 || (0 != 42)) {
                                                                                                v87 = 1
                                                                                                switch v19 {
                                                                                                case "\n\t":
                                                                                                    v23 -= v65
                                                                                                    var v147 int = (v70 >> ((2) & 63))
                                                                                                    var v148 uint16 = 0
                                                                                                    _ = v147
                                                                                                    _ = v148
                                                                                                case "a":
                                                                                                    _ = len(v2)
                                                                                                case "hello":
                                                                                                    var v149 float64 = 0.1
                                                                                                    _ = v149
                                                                                                }
                                                                                            } else {
                                                                                                switch v129 {
                                                                                                case 617:
                                                                                                    var v150 float32 = 1.5
                                                                                                    _ = v150
                                                                                                case 955:
                                                                                                    v151 := GenS19{F0: 0}
                                                                                                    _ = v151
                                                                                                default:
                                                                                                    v128 += 42
                                                                                                    v70 = (v70 % ((v127) | 1))
                                                                                                    v152 := GenS3{F0: v70, F1: v128, F2: (v24 / ((4294967295) | 1))}
                                                                                                    _ = v152
                                                                                                }
                                                                                            }
                                                                                            var v153 int64 = (v87 << ((166) & 63))
                                                                                            _ = v153
                                                                                        default:
                                                                                            v65 = (v65 + 0)
                                                                                        }
                                                                                    }
                                                                                    v154 := GenS7{F0: "foo bar", F1: 2, F2: -1}
                                                                                    _ = v130
                                                                                    _ = v154
                                                                                }
                                                                                _ = len(v2)
                                                                                _ = v129
                                                                            }
                                                                            v19 = ("a" + "a")
                                                                            v155 := GenS4{F0: 255, F1: 239}
                                                                            _ = v155
                                                                        }
                                                                        var v156 int8 = (^(v65))
                                                                        v21 -= v21
                                                                        var v157 string = "a"
                                                                        _ = len(v3)
                                                                        v156 -= (v156 & v156)
                                                                        var v158 int32 = -450
                                                                        _ = v127
                                                                        _ = v128
                                                                        _ = v156
                                                                        _ = v157
                                                                        _ = v158
                                                                    case "foo bar":
                                                                        v23 -= v23
                                                                    default:
                                                                        var v159 uint8 = 2
                                                                        _ = v159
                                                                    }
                                                                    switch fi60 {
                                                                    case 32767:
                                                                        v160 := GenS15{F0: v65}
                                                                        _ = v160
                                                                    case 9223372036854775807:
                                                                        var v161 bool = true
                                                                        _ = v161
                                                                    case 4294967295:
                                                                        v2[v70] = fi66
                                                                    default:
                                                                        v3 = append(v3, v19)
                                                                        v20 = v20
                                                                        var v162 uint8 = 42
                                                                        _ = v162
                                                                    }
                                                                }
                                                            }
                                                        case -93:
                                                            var v163 bool = ((!(1 == 18446744073709551615)) || true)
                                                            _ = v163
                                                        case 9223372036854775807:
                                                            var v164 int8 = (v65 * v65)
                                                            _ = v164
                                                        }
                                                        var v165 int8 = v65
                                                        _ = v165
                                                    }
                                                }
                                            case -296:
                                                v166 := GenS3{F0: (fi60 | (v22 * v70)), F1: 32768, F2: genFn0(0, (fi66 & -2147483648))}
                                                _ = v166
                                            case 9223372036854775807:
                                                fmt.Println(v87, v30, v70)
                                            default:
                                                v167 := GenS10{F0: 1000, F1: 2, F2: 1000}
                                                _ = v167
                                            }
                                            if v1 {
                                                if false {
                                                    _ = len(v3)
                                                    switch v19 {
                                                    case "a":
                                                        fmt.Println(v1, fi60)
                                                    default:
                                                        if v1 {
                                                            var v168 uint8 = 255
                                                            _ = v168
                                                        } else {
                                                            if true {
                                                                var v169 uint32 = (v20 % ((250) | 1))
                                                                fmt.Println(v30, fi60)
                                                                var v170 int32 = 2147483647
                                                                _ = v169
                                                                _ = v170
                                                            }
                                                        }
                                                    }
                                                    v171 := GenS2{F0: v21, F1: 783, F2: true}
                                                    delete(v2, (v22 & (v70 ^ (fi66 ^ fi66))))
                                                    _ = v171
                                                } else {
                                                    switch v19 {
                                                    case "\n\t":
                                                        v172 := v4[v19]
                                                        _ = v172
                                                    case "foo bar":
                                                        switch 0 {
                                                        case 32767:
                                                            var v173 float32 = (1.5 / 1.0)
                                                            _ = v173
                                                        case 409:
                                                            if (4294967295 > (-(v70))) {
                                                                var v174 bool = v30
                                                                fmt.Println(v1, fi66, v1)
                                                                _ = v174
                                                            }
                                                        case 1:
                                                            fmt.Println(v30, v30, v19)
                                                        default:
                                                            switch (v19 + ("a" + v19)) {
                                                            case "":
                                                                v175 := GenS12{F0: (v24 - v24), F1: 32767, F2: (v21 >> ((65535) & 63))}
                                                                _ = v175
                                                            case "hello":
                                                                var v176 float32 = 0.0
                                                                var v177 float64 = 1.5
                                                                var v178 string = (v19 + "hello")
                                                                v4["a"] = (^(v24))
                                                                var v179 uint64 = 51
                                                                _ = v176
                                                                _ = v177
                                                                _ = v178
                                                                _ = v179
                                                            case "\n\t":
                                                                var v180 string = v19
                                                                _ = v180
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(v23)
                                                    }
                                                    v181 := GenS20{F0: -128}
                                                    v30 = ((v65 == (v65 << ((0) & 63))) && (!(v1 && v1)))
                                                    _ = v181
                                                }
                                            } else {
                                                v30 = true
                                                fmt.Println(v22, v24, v23)
                                            }
                                            v65 -= v65
                                        case "hello":
                                            v20 = (v24 >> ((363) & 63))
                                        case "a":
                                            fmt.Println(fi60, v22, v30)
                                        default:
                                            switch "" {
                                            case "foo bar":
                                                var v182 int64 = v87
                                                _ = v182
                                            case "\n\t":
                                                var v183 string = v19
                                                _ = v183
                                            case "":
                                                if false {
                                                    v23 -= v65
                                                    var v184 float32 = ((1.5 - 1e10) + 3.14159)
                                                    var v185 string = "a"
                                                    _ = v184
                                                    _ = v185
                                                } else {
                                                    var v186 uint64 = 4294967295
                                                    if v1 {
                                                        fmt.Println(v30, v186)
                                                    } else {
                                                        v2[(v22 % ((int(v24)) | 1))] = 560
                                                        if (806 > (v186 >> ((0) & 63))) {
                                                            switch ("\n\t" + v19) {
                                                            case "hello":
                                                                var v187 int8 = -83
                                                                _ = v187
                                                            default:
                                                                var v188 int8 = 76
                                                                _ = v188
                                                            }
                                                        } else {
                                                            if true {
                                                                var v189 int64 = int64(v21)
                                                                _ = v189
                                                            }
                                                        }
                                                        if v1 {
                                                            v19 = (v19 + (v19 + v19))
                                                            var v190 string = v19
                                                            switch (fi60 - 770) {
                                                            case -9223372036854775808:
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            case 1:
                                                                var v191 float64 = 3.14159
                                                                _ = v191
                                                            default:
                                                                var v192 int64 = 0
                                                                _ = v192
                                                            }
                                                            _ = v190
                                                        }
                                                    }
                                                    _ = v186
                                                }
                                            }
                                        }
                                        _ = v87
                                    } else {
                                        switch v19 {
                                        case "hello":
                                            var v193 int32 = (v21 & v21)
                                            _ = v193
                                        case "\n\t":
                                            v22 = -1
                                        case "a":
                                            fmt.Println(v23, v1, fi66)
                                        }
                                        if v1 {
                                            v194 := GenS21{F0: 1000, F1: 1, F2: ((^(v24)) == (v24 << ((514) & 63)))}
                                            var v195 uint32 = (v20 / (((^(v24))) | 1))
                                            v22 -= fi60
                                            var v196 float64 = ((((1.5 + 0.1) - 0.1) + 1.5) - -2.25)
                                            _ = v194
                                            _ = v195
                                            _ = v196
                                        } else {
                                            v22 += (fi66 >> ((18446744073709551615) & 63))
                                            v23 += -128
                                            if len(v3) > 0 {
                                                v3[0] = v19
                                            }
                                        }
                                        var v197 string = v19
                                        _ = v197
                                    }
                                }
                                v198 := v4[v19]
                                var v199 uint64 = 442
                                var v200 int32 = (v21 / (((v21 % ((v21) | 1))) | 1))
                                v70 += 42
                                if (true && v30) {
                                    v201 := GenS15{F0: (-(v23))}
                                    v23 = -33
                                    v200 = v21
                                    _ = v201
                                }
                                _ = v198
                                _ = v199
                                _ = v200
                            }
                        case 476:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi66
                        _ = v70
                        _ = v71
                    }
                    _ = v65
                }
            }
            v202 := GenS11{F0: 84}
            _ = fi60
            _ = v202
        }
        v22 = (^(v22))
    }
    for fi203 := 0; fi203 < 4; fi203++ {
        var v204 uint32 = 627
        v30 = (!true)
        _ = fi203
        _ = v204
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
