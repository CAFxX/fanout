package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 int16; F1 int64; F2 uint16 }
type GenS9 struct { F0 uint64; F1 uint64; F2 int }
type GenS10 struct { F0 int16; F1 int8 }
type GenS11 struct { F0 bool; F1 uint }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64; F1 int16; F2 uint64 }
type GenS14 struct { F0 string; F1 uint16; F2 int }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int8; F1 int8 }
func genFn0(p0 int8, p1 uint64, p2 uint8) int16 {
    v4 := []int{2147483647, 2147483647, 4294967295, -9223372036854775808}
    var v5 uint32 = 846
    if true {
        v6 := GenS1{F0: v5, F1: 220, F2: 2}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v6
    } else {
        _ = len(v4)
    }
    v5 -= v5
    v4 = append(v4, 32768)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -871
}
func genFn1(p0 bool) bool {
    v7 := GenS2{F0: 1, F1: 100}
    p0 = (!p0)
    v8 := GenS2{F0: 32767, F1: 42}
    switch "a" {
    case "\n\t":
        fmt.Println(p0)
    }
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := make(map[string]string)
    var v2 string = "foo bar"
    v3 := make(map[int]string)
    var v9 uint = 255
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS3{F0: (^(v9))}
        if (-9223372036854775808 == 0) {
            v12 := []int{(fi10 >> ((v9) & 63)), (fi10 | fi10), fi10}
            switch (fi10 << (((v9 << (((v9 / (((v9 * (^(v9)))) | 1))) & 63))) & 63)) {
            case 9223372036854775807:
                if (!true) {
                    v13 := GenS4{F0: 361, F1: -533, F2: 65535}
                    fmt.Println(v9)
                    _ = v13
                }
            default:
                v14 := GenS5{F0: 66}
                _ = v14
            }
            _ = v12
        } else {
            var v15 int64 = int64(fi10)
            if ((-32768 > -946) && (false || (!(-586 > 0)))) {
                var v16 float32 = float32(-2.25)
                v16 = v16
                _ = v16
            } else {
                v15 = -407
            }
            fmt.Println(v9, v2, fi10)
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    var v17 int = 128
    if true {
        var v18 bool = true
        fmt.Println(v9, v9)
        _ = v18
    }
    var v19 uint64 = 4294967295
    if false {
        var v20 uint16 = 680
        v21 := v3[32767]
        for fi22 := 0; fi22 < 8; fi22++ {
            v19 = (v19 * v19)
            v23 := GenS1{F0: 42, F1: v20, F2: uint32(v19)}
            v24 := GenS6{F0: v19, F1: -1}
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = v20
        _ = v21
    } else {
        switch (^(v17)) {
        case 372:
            var v25 int = v17
            _ = v25
        case 4294967295:
            if true {
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 uint8 = uint8(v19)
                    v27 = (^(v27))
                    _ = fi26
                    _ = v27
                }
                v17 = v17
                for fi28 := 0; fi28 < 8; fi28++ {
                    fmt.Println(fi28, v19, fi28)
                    v29 := GenS5{F0: 255}
                    fmt.Println(fi28)
                    _ = fi28
                    _ = v29
                }
            } else {
                if (((v17 - (v17 - v17)) > (-(v17))) || true) {
                    for fi30 := 0; fi30 < 12; fi30++ {
                        var v31 float64 = (float64(((0.1 * ((-2.25 * (0.1 + (-2.25 + -2.25))) - ((-2.25 / (1e10 + 3.14159)) / (0.1 * (0.1 + 1e10))))) - float32(-2.25))) * (1e10 / 0.1))
                        _ = fi30
                        _ = v31
                    }
                    if (0 <= 1) {
                        if false {
                            if ((v17 % (((v17 - v17)) | 1)) == (v17 & 42)) {
                                for fi32 := 0; fi32 < 4; fi32++ {
                                    v33 := []int{fi32, fi32}
                                    v34 := GenS7{F0: -15, F1: v2}
                                    v19 = (^(v19))
                                    _ = fi32
                                    _ = v33
                                    _ = v34
                                }
                            }
                            if (!true) {
                                switch ((v2 + "foo bar") + (v2 + (("hello" + ((((("" + v2) + (v2 + v2)) + "") + v2) + (v2 + ("\n\t" + (v2 + "foo bar"))))) + (("\n\t" + v2) + (v2 + v2))))) {
                                case "a":
                                    var v35 uint64 = (v19 & 0)
                                    _ = v35
                                case "hello":
                                    switch 508 {
                                    case 743:
                                        fmt.Println(v17, v19)
                                    case 9223372036854775807:
                                        v36 := v1["\n\t"]
                                        _ = v36
                                    }
                                case "foo bar":
                                    if false {
                                        for fi37 := 0; fi37 < 16; fi37++ {
                                            if (false || (235 < 127)) {
                                                var v38 int8 = 42
                                                for fi39 := 0; fi39 < 2; fi39++ {
                                                    if ((v38 << (((v9 * (v9 - 395))) & 63)) <= -128) {
                                                        v38 = v38
                                                        v40 := []int{128}
                                                        switch fi39 {
                                                        case -9223372036854775808:
                                                            var v41 int = v17
                                                            _ = v41
                                                        }
                                                        _ = v40
                                                    }
                                                    var v42 float32 = (float32((3.14159 + 3.14159)) - -2.25)
                                                    switch (fi37 - v17) {
                                                    case -9223372036854775808:
                                                        fmt.Println(fi39)
                                                    default:
                                                        fmt.Println(v17, fi39)
                                                    }
                                                    _ = fi39
                                                    _ = v42
                                                }
                                                _ = v38
                                            } else {
                                                v43 := GenS8{F0: -32768, F1: 0, F2: 32767}
                                                v44 := GenS3{F0: (v9 * v9)}
                                                if true {
                                                    if (false || true) {
                                                        var v45 uint8 = 0
                                                        _ = v45
                                                    } else {
                                                        switch v17 {
                                                        case 0:
                                                            var v46 float32 = (1e10 - 1.5)
                                                            _ = v46
                                                        }
                                                        var v47 int64 = 333
                                                        if (((0.0 * (3.14159 - ((((0.0 + 0.1) / (0.1 * (-2.25 + -2.25))) + (3.14159 - -2.25)) + -2.25))) == (1.5 * 0.1)) && false) {
                                                            switch (-(v17)) {
                                                            case 4294967295:
                                                                var v48 int32 = -981
                                                                _ = v48
                                                            default:
                                                                v19 += 620
                                                            }
                                                        }
                                                        _ = v47
                                                    }
                                                    v49 := []int{(fi37 % ((fi37) | 1)), (v17 | v17), 305}
                                                    v9 = 4294967295
                                                    _ = v49
                                                }
                                                _ = v43
                                                _ = v44
                                            }
                                            _ = fi37
                                        }
                                        v9 = 255
                                        for fi50 := 0; fi50 < 4; fi50++ {
                                            var v51 float32 = 0.0
                                            _ = fi50
                                            _ = v51
                                        }
                                    }
                                default:
                                    if true {
                                        v9 += (v9 + 2)
                                    }
                                }
                                if true {
                                    v52 := GenS7{F0: 42, F1: (v2 + (v2 + v2))}
                                    _ = v52
                                }
                                v53 := GenS9{F0: v19, F1: (v19 >> (((^(v9))) & 63)), F2: v17}
                                _ = v53
                            } else {
                                v19 -= 908
                                v54 := GenS10{F0: -451, F1: 111}
                                _ = v54
                            }
                        }
                        for fi55 := 0; fi55 < 13; fi55++ {
                            v19 = (v19 & 4294967295)
                            v9 -= v9
                            switch (v17 >> (((^(v9))) & 63)) {
                            case -9223372036854775808:
                                if (uint32(v9) > 340) {
                                    v2 = v2
                                    v56 := GenS11{F0: true, F1: (v9 + v9)}
                                    var v57 int32 = 0
                                    _ = v56
                                    _ = v57
                                } else {
                                    fmt.Println(v19, v17)
                                }
                            case 9223372036854775807:
                                v58 := []int{}
                                _ = v58
                            case 42:
                                for fi59 := 0; fi59 < 14; fi59++ {
                                    switch "" {
                                    case "\n\t":
                                        if (0.0 >= float64(3.14159)) {
                                            fmt.Println(v9, fi55)
                                            v2 = v2
                                            v17 = 0
                                        } else {
                                            v9 -= (v9 & v9)
                                            fmt.Println(v19, v17)
                                            v60 := GenS8{F0: 295, F1: 1, F2: 255}
                                            var v61 uint16 = 65535
                                            _ = v60
                                            _ = v61
                                        }
                                    case "foo bar":
                                        if (false && true) {
                                            var v62 float32 = (1.5 - (3.14159 - 1e10))
                                            v62 += 0.0
                                            fmt.Println(fi59, v17, v9)
                                            _ = v62
                                        }
                                    default:
                                        v63 := GenS12{F0: v2}
                                        _ = v63
                                    }
                                    v64 := []int{v17, v17, v17, 4294967295}
                                    if len(v64) > 0 {
                                        v64[0] = (^(fi59))
                                    }
                                    _ = fi59
                                    _ = v64
                                }
                            default:
                                var v65 int8 = -43
                                _ = v65
                            }
                            _ = fi55
                        }
                        v66 := GenS13{F0: -1, F1: 32767, F2: 0}
                        _ = v66
                    } else {
                        v9 = (v9 ^ v9)
                    }
                    for fi67 := 0; fi67 < 18; fi67++ {
                        var v68 int64 = 9223372036854775807
                        _ = fi67
                        _ = v68
                    }
                }
                var v69 bool = true
                if v69 {
                    v2 = "a"
                }
                _ = v69
            }
        case 9223372036854775807:
            for fi70 := 0; fi70 < 17; fi70++ {
                var v71 uint16 = 65535
                v71 -= v71
                var v72 uint16 = (v71 ^ (v71 / ((v71) | 1)))
                _ = fi70
                _ = v71
                _ = v72
            }
        default:
            v2 = ((v2 + v2) + v2)
        }
        switch 755 {
        case -457:
            if ((!(true || (1e10 >= 1e10))) && (false && (971 < uint32(v9)))) {
                for fi73 := 0; fi73 < 5; fi73++ {
                    var v74 float64 = -2.25
                    if ((255 > 567) || false) {
                        var v75 float64 = v74
                        _ = v75
                    } else {
                        for fi76 := 0; fi76 < 8; fi76++ {
                            switch (fi73 >> ((v9) & 63)) {
                            case 0:
                                var v77 int = (v17 ^ 378)
                                _ = v77
                            case -1:
                                v78 := GenS11{F0: false, F1: 42}
                                _ = v78
                            default:
                                if (!false) {
                                    var v79 float64 = v74
                                    _ = v79
                                } else {
                                    var v80 int16 = -489
                                    var v81 int32 = -499
                                    _ = v80
                                    _ = v81
                                }
                            }
                            _ = fi76
                        }
                    }
                    _ = fi73
                    _ = v74
                }
            }
        case -797:
            if true {
                v82 := GenS2{F0: genFn0(118, 18446744073709551615, 8), F1: 124}
                var v83 uint16 = 552
                if false {
                    switch (v2 + (v2 + (v2 + v2))) {
                    case "a":
                        for fi84 := 0; fi84 < 4; fi84++ {
                            switch (fi84 | fi84) {
                            case 32767:
                                var v85 float64 = float64((1e10 * (((3.14159 / -2.25) * 1e10) * 3.14159)))
                                _ = v85
                            case 850:
                                var v86 uint16 = (v83 / ((v83) | 1))
                                _ = v86
                            case -615:
                                fmt.Println(v19, v2, v9)
                            default:
                                v9 = v9
                            }
                            v19 += (v19 & v19)
                            _ = fi84
                        }
                    case "foo bar":
                        fmt.Println(v17)
                    case "hello":
                        v83 -= (v83 << ((v9) & 63))
                    }
                    if true {
                        v19 = v19
                        fmt.Println(v9, v9)
                        v87 := GenS14{F0: (v2 + v2), F1: (^(v83)), F2: int(v83)}
                        _ = v87
                    } else {
                        v19 -= v19
                        if true {
                            var v88 uint64 = (v19 + v19)
                            switch 128 {
                            case -2147483648:
                                fmt.Println(v88, v88)
                            case -615:
                                var v89 uint32 = uint32(v19)
                                _ = v89
                            case 0:
                                var v90 int = 4294967295
                                _ = v90
                            default:
                                v83 += v83
                            }
                            _ = v88
                        } else {
                            fmt.Println(v19, v19)
                            for fi91 := 0; fi91 < 6; fi91++ {
                                switch fi91 {
                                case 815:
                                    for fi92 := 0; fi92 < 19; fi92++ {
                                        var v93 uint8 = 0
                                        _ = fi92
                                        _ = v93
                                    }
                                case 32768:
                                    v94 := []int{v17}
                                    _ = v94
                                default:
                                    v95 := v3[0]
                                    _ = v95
                                }
                                _ = fi91
                            }
                            var v96 int64 = 9223372036854775807
                            _ = v96
                        }
                        if (v9 > v9) {
                            if (2 != 4294967295) {
                                v97 := []int{v17}
                                for fi98 := 0; fi98 < 8; fi98++ {
                                    var v99 string = v2
                                    var v100 int32 = 1
                                    _ = fi98
                                    _ = v100
                                    _ = v99
                                }
                                v101 := GenS15{F0: 42}
                                _ = v101
                                _ = v97
                            }
                            v83 += (v83 | (^(v83)))
                            for fi102 := 0; fi102 < 13; fi102++ {
                                v103 := []int{(v17 + 4294967295), (fi102 % ((fi102) | 1)), (^(fi102)), fi102}
                                if len(v103) >= 2 {
                                    v103 = v103[1:len(v103)]
                                }
                                v17 -= (v17 * (fi102 & (fi102 + (v17 >> ((0) & 63)))))
                                _ = fi102
                                _ = v103
                            }
                        } else {
                            var v104 string = ("\n\t" + "foo bar")
                            if ((42 < 42) && ((-2147483648 == 0) || true)) {
                                if true {
                                    v105 := GenS16{F0: -1, F1: -120}
                                    var v106 uint = v9
                                    _ = v105
                                    _ = v106
                                }
                            } else {
                                if ((0.1 != 0.0) && false) {
                                    var v107 int8 = 100
                                    _ = v107
                                }
                            }
                            var v108 int16 = 1000
                            _ = v104
                            _ = v108
                        }
                    }
                    v109 := GenS10{F0: -1, F1: 42}
                    _ = v109
                }
                _ = v82
                _ = v83
            } else {
                var v110 uint8 = 127
                _ = v110
            }
        case 128:
            v2 = ("foo bar" + (v2 + v2))
        }
    }
    for fi111 := 0; fi111 < 5; fi111++ {
        v17 -= (fi111 >> ((v9) & 63))
        v9 = 1
        var v112 bool = (!true)
        _ = fi111
        _ = v112
    }
    for fi113 := 0; fi113 < 11; fi113++ {
        fmt.Println(fi113, v2, v9)
        v9 += (v9 / (((v9 * 42)) | 1))
        _ = fi113
    }
    if false {
        for fi114 := 0; fi114 < 5; fi114++ {
            if false {
                v17 = 0
            }
            v115 := []int{598, (v17 % (((fi114 * v17)) | 1))}
            v17 = v17
            _ = fi114
            _ = v115
        }
        for fi116 := 0; fi116 < 10; fi116++ {
            if (!(!true)) {
                switch (v17 % (((fi116 << (((v9 << ((792) & 63))) & 63))) | 1)) {
                case 2147483647:
                    fmt.Println(v2)
                }
                v19 = v19
            }
            _ = fi116
        }
        v17 += v17
    }
    var v117 uint16 = 807
    v17 = v17
    var v118 int8 = 0
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v117
    _ = v118
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
