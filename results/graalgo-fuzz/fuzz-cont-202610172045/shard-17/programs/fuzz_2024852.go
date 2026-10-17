package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint; F2 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 int32; F1 int8 }
type GenS11 struct { F0 int8; F1 bool; F2 uint }
type GenS12 struct { F0 bool; F1 bool }
type GenS13 struct { F0 int32; F1 string }
type GenS14 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint16; F1 int; F2 uint16 }
type GenS19 struct { F0 int16; F1 uint64; F2 bool }
type GenS20 struct { F0 uint }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 int32; F1 uint64 }
type GenS23 struct { F0 int64; F1 uint }
type GenS24 struct { F0 uint8; F1 string; F2 uint32 }
type GenS25 struct { F0 int16; F1 int; F2 uint8 }
type GenS26 struct { F0 int8; F1 int64; F2 uint }
func genFn0(p0 uint8) uint8 {
    if false {
        p0 -= (p0 / (((p0 ^ (p0 & (^(p0))))) | 1))
    }
    p0 = 128
    p0 = (p0 ^ (p0 - p0))
    _ = p0
    return (p0 >> ((65535) & 63))
}
func genFn1(p0 int) uint64 {
    var v5 int32 = 1
    var v6 string = "hello"
    v7 := GenS1{F0: 42, F1: 65535, F2: -9223372036854775808}
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(v6, p0, v6)
        fmt.Println(p0, p0, v5)
        var v9 uint16 = 32767
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := []int{-9223372036854775808, -1}
    v2 := []int{}
    var v3 uint16 = 206
    var v4 string = (("a" + "hello") + "a")
    if (false || true) {
        v4 = v4
    }
    var v10 uint16 = 42
    v4 = ("\n\t" + v4)
    v11 := GenS1{F0: 42, F1: 577, F2: -9223372036854775808}
    var v12 uint8 = 255
    for fi13 := 0; fi13 < 5; fi13++ {
        v4 = "foo bar"
        fmt.Println(v12, v10, v12)
        for fi14 := 0; fi14 < 5; fi14++ {
            _ = len(v1)
            var v15 uint32 = 10
            v16 := GenS2{F0: true}
            var v17 uint16 = v10
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    v18 := GenS1{F0: int8(v3), F1: 2147483648, F2: -9223372036854775808}
    v3 -= uint16(v12)
    fmt.Println(v12, v12)
    var v19 int32 = -800
    if (759 >= 475) {
        var v20 float64 = 1e10
        _ = v20
    } else {
        v21 := GenS2{F0: true}
        v4 = ((v4 + v4) + v4)
        _ = v21
    }
    var v22 int16 = -354
    for fi23 := 0; fi23 < 14; fi23++ {
        v19 -= (v19 / (((^(v19))) | 1))
        switch (^(fi23)) {
        case -9223372036854775808:
            switch v4 {
            case "":
                v1 = append(v1, fi23)
            case "hello":
                for fi24 := 0; fi24 < 14; fi24++ {
                    if (!true) {
                        var v25 float32 = 0.1
                        _ = v25
                    } else {
                        switch fi24 {
                        case 9223372036854775807:
                            var v26 float32 = 0.1
                            v27 := GenS3{F0: v3}
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                            _ = v26
                            _ = v27
                        case 354:
                            v12 -= genFn0(v12)
                            switch "hello" {
                            case "":
                                if true {
                                    if (865 < 0) {
                                        var v28 int32 = (^(v19))
                                        _ = v28
                                    } else {
                                        fmt.Println(v19, fi24)
                                        var v29 int8 = 12
                                        _ = v29
                                    }
                                    var v30 int8 = 42
                                    _ = v30
                                }
                                var v31 string = "\n\t"
                                v10 += (v3 + v10)
                                _ = v31
                            case "hello":
                                v22 = int16(v10)
                            default:
                                fmt.Println(fi23, v10)
                            }
                            var v32 float32 = 1e10
                            _ = v32
                        case -2147483648:
                            v33 := GenS4{F0: (0.1 >= (((3.14159 - ((float32((1.5 / 1.0)) + (1e10 - 1.5)) * (float32(0.1) / 3.14159))) / 3.14159) * (1e10 * ((-2.25 * (1.5 / -2.25)) + 1e10)))), F1: (v19 % ((int32(v10)) | 1))}
                            _ = v33
                        }
                        _ = len(v2)
                        var v34 uint8 = (v12 % ((v12) | 1))
                        _ = v34
                    }
                    v19 -= v19
                    _ = fi24
                }
            default:
                var v35 uint64 = 42
                _ = v35
            }
        case 128:
            var v36 bool = true
            _ = v36
        case 0:
            if len(v1) > 0 {
                v1[0] = (fi23 - 1)
            }
        }
        if true {
            if ((false || (!false)) || true) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (0.0 == (1.5 - 0.0)) {
                    if ((((1e10 + 0.0) - 3.14159) / 1.0) >= 0.0) {
                        var v37 int32 = v19
                        if (v37 < (v19 >> ((544) & 63))) {
                            for fi38 := 0; fi38 < 9; fi38++ {
                                if false {
                                    switch fi38 {
                                    case -427:
                                        fmt.Println(v37)
                                    case 1:
                                        v1 = append(v1, -857)
                                    case 718:
                                        fmt.Println(v37, fi23)
                                    }
                                    v37 -= v19
                                }
                                if (0 >= 686) {
                                    if (0 != -1) {
                                        if ((1 != v22) || (!((!((fi38 ^ -1) <= fi38)) || (false || false)))) {
                                            if ((true && (true || true)) && false) {
                                                v12 = (v12 & genFn0(v12))
                                            } else {
                                                var v39 uint = 4294967295
                                                _ = v39
                                            }
                                        }
                                    }
                                    v10 = v10
                                } else {
                                    if false {
                                        v1 = append(v1, fi23)
                                        if false {
                                            _ = len(v2)
                                            var v40 float32 = (((((((-2.25 / float32(1e10)) - (0.0 + -2.25)) + (((-2.25 / (1e10 / 1e10)) - 3.14159) * float32(-2.25))) / -2.25) + (1e10 / (((0.1 * ((float32((3.14159 + 1.5)) * 1e10) - -2.25)) - 1.5) - 0.1))) / (3.14159 / (0.1 + (0.0 * 1e10)))) + 3.14159)
                                            _ = v40
                                        }
                                        _ = len(v2)
                                    }
                                }
                                v37 = v37
                                _ = fi38
                            }
                        } else {
                            v4 = v4
                            switch "hello" {
                            case "a":
                                var v41 uint = 928
                                _ = v41
                            }
                            if (genFn0(v12) > v12) {
                                v42 := GenS3{F0: (v10 ^ (v10 + 135))}
                                var v43 uint64 = 1
                                _ = v42
                                _ = v43
                            } else {
                                fmt.Println(v12, v4)
                            }
                        }
                        if (genFn1(fi23) != 86) {
                            v3 = v10
                        } else {
                            if true {
                                switch fi23 {
                                case 9223372036854775807:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case 32767:
                                    if true {
                                        var v44 float64 = 3.14159
                                        fmt.Println(v22, v4, v19)
                                        _ = v44
                                    } else {
                                        v45 := GenS2{F0: (100 > -128)}
                                        v19 = (-(v37))
                                        _ = v45
                                    }
                                default:
                                    if (!false) {
                                        for fi46 := 0; fi46 < 10; fi46++ {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v47 := GenS4{F0: true, F1: (v19 / (((v37 ^ v19)) | 1))}
                                            v19 = v19
                                            _ = fi46
                                            _ = v47
                                        }
                                    } else {
                                        var v48 uint64 = 1
                                        switch v4 {
                                        case "foo bar":
                                            if (255 > 68) {
                                                v4 = "\n\t"
                                                for fi49 := 0; fi49 < 13; fi49++ {
                                                    var v50 uint8 = v12
                                                    _ = fi49
                                                    _ = v50
                                                }
                                                var v51 int32 = v37
                                                _ = v51
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                if true {
                                                    fmt.Println(v19, v4)
                                                    for fi52 := 0; fi52 < 14; fi52++ {
                                                        var v53 uint64 = (v48 ^ v48)
                                                        v22 -= 217
                                                        switch v4 {
                                                        case "foo bar":
                                                            if (!false) {
                                                                var v54 uint32 = 647
                                                                var v55 uint16 = 0
                                                                switch fi23 {
                                                                case -1:
                                                                    var v56 int64 = 9223372036854775807
                                                                    _ = v56
                                                                default:
                                                                    switch (fi23 << ((uint(v53)) & 63)) {
                                                                    case 1:
                                                                        v54 += (v54 ^ (v54 % (((v54 * v54)) | 1)))
                                                                        var v57 int = (fi52 + (fi52 ^ (fi52 - 2147483647)))
                                                                        if ((^(v37)) == v19) {
                                                                            if ((v3 >> ((65535) & 63)) == uint16(v54)) {
                                                                                if (855 <= v48) {
                                                                                    fmt.Println(v3, v48)
                                                                                } else {
                                                                                    v37 -= (-(v19))
                                                                                    if true {
                                                                                        if false {
                                                                                            if true {
                                                                                                if (v10 >= v10) {
                                                                                                    var v58 int = -1
                                                                                                    _ = v58
                                                                                                }
                                                                                                v3 = 195
                                                                                                fmt.Println(v57)
                                                                                                v54 -= 326
                                                                                            } else {
                                                                                                v59 := GenS5{F0: false}
                                                                                                _ = v59
                                                                                            }
                                                                                        }
                                                                                        v4 = (v4 + (v4 + v4))
                                                                                    } else {
                                                                                        if true {
                                                                                            var v60 float64 = 0.0
                                                                                            var v61 uint8 = 66
                                                                                            _ = v60
                                                                                            _ = v61
                                                                                        }
                                                                                        fmt.Println(v22, v12, v57)
                                                                                        fmt.Println(v10, v37)
                                                                                    }
                                                                                    var v62 int16 = 1000
                                                                                    v10 -= uint16(v12)
                                                                                    _ = v62
                                                                                }
                                                                                switch fi23 {
                                                                                case 1:
                                                                                    if true {
                                                                                        var v63 int64 = -1
                                                                                        if (!true) {
                                                                                            var v64 uint16 = (^(v10))
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = (fi23 * 9223372036854775807)
                                                                                            }
                                                                                            _ = v64
                                                                                        }
                                                                                        v55 = (^(v55))
                                                                                        _ = v63
                                                                                    } else {
                                                                                        var v65 uint32 = (v54 - (v54 | v54))
                                                                                        var v66 uint8 = (v12 & 166)
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                    }
                                                                                default:
                                                                                    v55 -= v3
                                                                                }
                                                                                v22 = int16(v48)
                                                                            } else {
                                                                                var v67 int = 730
                                                                                var v68 bool = false
                                                                                switch v67 {
                                                                                case -2147483648:
                                                                                    fmt.Println(v55, fi23, v12)
                                                                                    v1 = append(v1, int(v12))
                                                                                case -904:
                                                                                    var v69 int8 = -128
                                                                                    _ = v69
                                                                                case -845:
                                                                                    v70 := GenS6{F0: ("\n\t" + v4), F1: 42, F2: 1}
                                                                                    _ = v70
                                                                                default:
                                                                                    if ((0.1 <= float32(-2.25)) || false) {
                                                                                        switch (v4 + v4) {
                                                                                        case "hello":
                                                                                            v71 := GenS1{F0: -128, F1: (v54 % (((v54 - v54)) | 1)), F2: 541}
                                                                                            _ = v71
                                                                                        case "a":
                                                                                            if v68 {
                                                                                                var v72 string = (v4 + "\n\t")
                                                                                                fmt.Println(v57, v22, fi52)
                                                                                                _ = v72
                                                                                            } else {
                                                                                                if (!(v22 == (v22 & 42))) {
                                                                                                    v3 += (v10 + 876)
                                                                                                    v73 := GenS3{F0: 1}
                                                                                                    if (v12 <= 12) {
                                                                                                        v54 += 4294967295
                                                                                                        v10 += (v55 | v10)
                                                                                                        switch 128 {
                                                                                                        case 1:
                                                                                                            v57 -= (fi52 | (^(fi23)))
                                                                                                        case -835:
                                                                                                            v48 -= v48
                                                                                                        case 786:
                                                                                                            if (!v68) {
                                                                                                                v74 := GenS7{F0: 908}
                                                                                                                _ = len(v2)
                                                                                                                _ = v74
                                                                                                            } else {
                                                                                                                var v75 uint8 = 55
                                                                                                                var v76 bool = (0 > uint(v54))
                                                                                                                var v77 uint8 = v12
                                                                                                                _ = v75
                                                                                                                _ = v76
                                                                                                                _ = v77
                                                                                                            }
                                                                                                        default:
                                                                                                            v78 := GenS6{F0: v4, F1: 2, F2: int64(v57)}
                                                                                                            _ = v78
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v79 float32 = (((1.5 - 0.0) / 3.14159) + (1e10 * (float32(0.0) * (float32(1e10) * 3.14159))))
                                                                                                        v80 := GenS5{F0: (348 <= (v48 + v53))}
                                                                                                        var v81 int64 = 9223372036854775807
                                                                                                        _ = v79
                                                                                                        _ = v80
                                                                                                        _ = v81
                                                                                                    }
                                                                                                    _ = v73
                                                                                                }
                                                                                                if (!v68) {
                                                                                                    var v82 float32 = ((float32(3.14159) / (-2.25 + -2.25)) / 1.0)
                                                                                                    if true {
                                                                                                        v55 = v10
                                                                                                        if ((v22 >= v22) && v68) {
                                                                                                            var v83 int64 = 695
                                                                                                            if (!((v48 < genFn1(v67)) || false)) {
                                                                                                                if len(v2) >= 2 {
                                                                                                                    v2 = v2[1:len(v2)]
                                                                                                                }
                                                                                                            } else {
                                                                                                                fmt.Println(v19)
                                                                                                            }
                                                                                                            if (v68 && v68) {
                                                                                                                if v68 {
                                                                                                                    var v84 float64 = 0.1
                                                                                                                    if v68 {
                                                                                                                        if (9223372036854775807 == fi52) {
                                                                                                                            v85 := GenS2{F0: v68}
                                                                                                                            v2 = append(v2, (fi52 >> ((2) & 63)))
                                                                                                                            _ = v85
                                                                                                                        } else {
                                                                                                                            v86 := GenS8{F0: -2, F1: v22}
                                                                                                                            if (v68 && (122 != v37)) {
                                                                                                                                _ = len(v2)
                                                                                                                                var v87 int64 = 816
                                                                                                                                var v88 string = (v4 + (v4 + v4))
                                                                                                                                if v68 {
                                                                                                                                    fmt.Println(fi23)
                                                                                                                                    fmt.Println(fi52, fi23, v10)
                                                                                                                                }
                                                                                                                                var v89 int64 = 0
                                                                                                                                var v90 uint32 = (v54 >> ((2) & 63))
                                                                                                                                _ = v87
                                                                                                                                _ = v88
                                                                                                                                _ = v89
                                                                                                                                _ = v90
                                                                                                                            }
                                                                                                                            _ = v86
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if v68 {
                                                                                                                        v4 = ("\n\t" + "foo bar")
                                                                                                                        switch (-(v67)) {
                                                                                                                        case 32767:
                                                                                                                            v67 += -9223372036854775808
                                                                                                                        case -1:
                                                                                                                            var v91 int32 = -100
                                                                                                                            _ = v91
                                                                                                                        case -971:
                                                                                                                            if true {
                                                                                                                                if ((v68 || (v83 < v83)) && ((v12 / (((^(v12))) | 1)) <= (^(v12)))) {
                                                                                                                                    if v68 {
                                                                                                                                        if v68 {
                                                                                                                                            v92 := GenS8{F0: 1, F1: v22}
                                                                                                                                            if (((v68 && (((v82 * float32((v84 - v84))) / 1.0) == v82)) || true) || (false && (!false))) {
                                                                                                                                                v55 -= (^(v10))
                                                                                                                                                v83 = 1
                                                                                                                                            } else {
                                                                                                                                                fmt.Println(v55)
                                                                                                                                            }
                                                                                                                                            _ = v92
                                                                                                                                        }
                                                                                                                                        switch ("" + "") {
                                                                                                                                        case "":
                                                                                                                                            var v93 float32 = float32((v84 * v84))
                                                                                                                                            _ = v93
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    switch (fi23 | (fi52 * fi23)) {
                                                                                                                                    case 42:
                                                                                                                                        v12 += 163
                                                                                                                                    default:
                                                                                                                                        v94 := GenS9{F0: "a", F1: (^(v10))}
                                                                                                                                        _ = v94
                                                                                                                                    }
                                                                                                                                    _ = len(v1)
                                                                                                                                    v84 = v84
                                                                                                                                }
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            _ = len(v1)
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v84
                                                                                                                } else {
                                                                                                                    if ((!(255 != 2)) || true) {
                                                                                                                        var v95 string = v4
                                                                                                                        var v96 uint32 = (v54 % (((v54 % ((uint32(v3)) | 1))) | 1))
                                                                                                                        _ = v95
                                                                                                                        _ = v96
                                                                                                                    } else {
                                                                                                                        fmt.Println(v10, fi52)
                                                                                                                        var v97 uint16 = (^(v3))
                                                                                                                        _ = v97
                                                                                                                    }
                                                                                                                    var v98 int8 = int8(v54)
                                                                                                                    v99 := GenS4{F0: v68, F1: v37}
                                                                                                                    v100 := GenS7{F0: int64(v54)}
                                                                                                                    _ = v100
                                                                                                                    _ = v98
                                                                                                                    _ = v99
                                                                                                                }
                                                                                                                v101 := GenS9{F0: "foo bar", F1: uint16(v54)}
                                                                                                                _ = v101
                                                                                                            }
                                                                                                            v102 := GenS1{F0: 0, F1: v54, F2: (v83 & (v83 >> ((215) & 63)))}
                                                                                                            if v68 {
                                                                                                                if (v68 && v68) {
                                                                                                                    if true {
                                                                                                                        v4 = ("\n\t" + ((v4 + (v4 + ("\n\t" + v4))) + v4))
                                                                                                                        if false {
                                                                                                                            if v68 {
                                                                                                                                var v103 string = v4
                                                                                                                                _ = v103
                                                                                                                            } else {
                                                                                                                                v104 := GenS4{F0: (v54 == v54), F1: 0}
                                                                                                                                if ((v83 ^ v83) < -9223372036854775808) {
                                                                                                                                    var v105 string = "\n\t"
                                                                                                                                    _ = v105
                                                                                                                                }
                                                                                                                                switch (v67 ^ (fi52 >> ((18446744073709551615) & 63))) {
                                                                                                                                case 32768:
                                                                                                                                    var v106 uint32 = v54
                                                                                                                                    _ = v106
                                                                                                                                case -9223372036854775808:
                                                                                                                                    _ = len(v1)
                                                                                                                                case -427:
                                                                                                                                    v107 := GenS6{F0: "a", F1: 912, F2: -145}
                                                                                                                                    _ = v107
                                                                                                                                default:
                                                                                                                                    var v108 int = (v67 & v57)
                                                                                                                                    _ = v108
                                                                                                                                }
                                                                                                                                var v109 float64 = -2.25
                                                                                                                                var v110 int = (-(fi52))
                                                                                                                                _ = v104
                                                                                                                                _ = v109
                                                                                                                                _ = v110
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v4 = v4
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    fmt.Println(v12, v83)
                                                                                                                    v83 -= v83
                                                                                                                    _ = len(v2)
                                                                                                                    switch v57 {
                                                                                                                    case -867:
                                                                                                                        v67 = int(v53)
                                                                                                                    case 0:
                                                                                                                        if v68 {
                                                                                                                            fmt.Println(fi52)
                                                                                                                            fmt.Println(v83)
                                                                                                                            var v111 int8 = int8(v22)
                                                                                                                            _ = v111
                                                                                                                        }
                                                                                                                    case -9223372036854775808:
                                                                                                                        _ = len(v1)
                                                                                                                        v112 := GenS10{F0: 314, F1: 127}
                                                                                                                        _ = v112
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            if true {
                                                                                                                switch v4 {
                                                                                                                case "\n\t":
                                                                                                                    if v68 {
                                                                                                                        v48 += v48
                                                                                                                    } else {
                                                                                                                        v48 -= (v53 >> ((300) & 63))
                                                                                                                        switch (fi23 << ((2) & 63)) {
                                                                                                                        case 9223372036854775807:
                                                                                                                            var v113 float64 = (0.0 - (0.1 + (0.1 / ((0.0 - 3.14159) / 1.5))))
                                                                                                                            _ = v113
                                                                                                                        case 2147483647:
                                                                                                                            if v68 {
                                                                                                                                switch (v4 + ("foo bar" + v4)) {
                                                                                                                                case "hello":
                                                                                                                                    var v114 string = "a"
                                                                                                                                    _ = v114
                                                                                                                                }
                                                                                                                                _ = len(v2)
                                                                                                                                v2 = append(v2, 507)
                                                                                                                                switch fi23 {
                                                                                                                                case -36:
                                                                                                                                    v4 = v4
                                                                                                                                }
                                                                                                                                switch -372 {
                                                                                                                                case 519:
                                                                                                                                    var v115 int32 = v37
                                                                                                                                    _ = v115
                                                                                                                                }
                                                                                                                                if ((((v4 + "foo bar") + "\n\t") + v4) != "") {
                                                                                                                                    if (!v68) {
                                                                                                                                        v67 += (fi23 + v57)
                                                                                                                                    } else {
                                                                                                                                        if (v53 == v53) {
                                                                                                                                            v54 = uint32(v12)
                                                                                                                                            var v116 uint32 = (v54 * v54)
                                                                                                                                            var v117 int8 = -128
                                                                                                                                            _ = v116
                                                                                                                                            _ = v117
                                                                                                                                        }
                                                                                                                                        v4 = "a"
                                                                                                                                        if v68 {
                                                                                                                                            fmt.Println(v3, v22)
                                                                                                                                            v67 += (fi23 & fi23)
                                                                                                                                            var v118 uint16 = (^(v55))
                                                                                                                                            _ = v118
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    v19 = v37
                                                                                                                                    var v119 float64 = (-2.25 * ((1e10 - ((3.14159 - float64((float32((0.0 * ((float64(float32(float64((((((v82 / ((1.5 / (float32(1e10) / float32(float64(v82)))) / 3.14159)) * (float32(3.14159) * v82)) / v82) - (float32(3.14159) - 1.5)) / v82)))) / 1.0) / -2.25))) - float32(1e10)))) - (1.5 + (-2.25 * (float64(float32((-2.25 / float64(v82)))) * float64(v82)))))) / ((1e10 / 1.5) + 0.1)))
                                                                                                                                    _ = v119
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                switch "" {
                                                                                                                                case "foo bar":
                                                                                                                                    var v120 uint8 = uint8(v54)
                                                                                                                                    _ = v120
                                                                                                                                case "hello":
                                                                                                                                    if (9223372036854775807 != (v83 ^ 420)) {
                                                                                                                                        switch fi23 {
                                                                                                                                        case 2147483647:
                                                                                                                                            v121 := GenS9{F0: v4, F1: 423}
                                                                                                                                            _ = v121
                                                                                                                                        default:
                                                                                                                                            fmt.Println(v4)
                                                                                                                                            var v122 uint32 = (v54 - v54)
                                                                                                                                            var v123 string = (v4 + v4)
                                                                                                                                            _ = v122
                                                                                                                                            _ = v123
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                case "":
                                                                                                                                    v124 := GenS5{F0: (fi52 >= (fi23 | (fi52 - (fi23 & -1))))}
                                                                                                                                    v2 = append(v2, (^(v57)))
                                                                                                                                    v53 = v53
                                                                                                                                    var v125 int8 = 25
                                                                                                                                    var v126 uint16 = v3
                                                                                                                                    _ = v124
                                                                                                                                    _ = v125
                                                                                                                                    _ = v126
                                                                                                                                default:
                                                                                                                                    if len(v1) >= 2 {
                                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                        case 1:
                                                                                                                            v55 -= 0
                                                                                                                        default:
                                                                                                                            if len(v1) >= 2 {
                                                                                                                                v1 = v1[1:len(v1)]
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    if v68 {
                                                                                                                        if v68 {
                                                                                                                            v1 = append(v1, fi52)
                                                                                                                            var v127 float32 = float32(-2.25)
                                                                                                                            v128 := GenS11{F0: 1, F1: (v48 > 0), F2: 255}
                                                                                                                            fmt.Println(v83, v55, v57)
                                                                                                                            var v129 int8 = 24
                                                                                                                            _ = v127
                                                                                                                            _ = v128
                                                                                                                            _ = v129
                                                                                                                        }
                                                                                                                        switch (fi23 | (fi23 + (v57 % ((v57) | 1)))) {
                                                                                                                        case -609:
                                                                                                                            var v130 uint16 = (^(v55))
                                                                                                                            _ = v130
                                                                                                                        default:
                                                                                                                            var v131 uint64 = (^(v48))
                                                                                                                            _ = v131
                                                                                                                        }
                                                                                                                        var v132 int64 = (v83 / ((v83) | 1))
                                                                                                                        v54 -= v54
                                                                                                                        _ = v132
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v133 := GenS8{F0: -95, F1: 228}
                                                                                                                    _ = v133
                                                                                                                }
                                                                                                            }
                                                                                                            if ((v12 < genFn0((v12 & (^(v12))))) && false) {
                                                                                                                switch v57 {
                                                                                                                case -9223372036854775808:
                                                                                                                    v10 += v55
                                                                                                                case 167:
                                                                                                                    if v68 {
                                                                                                                        if v68 {
                                                                                                                            v19 = (v37 - int32(v57))
                                                                                                                            switch "a" {
                                                                                                                            case "foo bar":
                                                                                                                                _ = len(v1)
                                                                                                                                if ((151 > 143) || true) {
                                                                                                                                    v82 += v82
                                                                                                                                    fmt.Println(v22, v12)
                                                                                                                                    v134 := GenS12{F0: (v37 > (v37 | (v37 >> ((18446744073709551615) & 63)))), F1: (v48 == (v53 * v48))}
                                                                                                                                    _ = v134
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v67)
                                                                                                                                    v135 := GenS13{F0: v37, F1: v4}
                                                                                                                                    _ = v135
                                                                                                                                }
                                                                                                                                if true {
                                                                                                                                    v136 := GenS14{F0: 56, F1: v53, F2: (v53 | uint64(v54))}
                                                                                                                                    if len(v1) >= 2 {
                                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                                    }
                                                                                                                                    _ = v136
                                                                                                                                }
                                                                                                                            case "a":
                                                                                                                                if len(v2) >= 2 {
                                                                                                                                    v2 = v2[1:len(v2)]
                                                                                                                                }
                                                                                                                            case "hello":
                                                                                                                                var v137 uint32 = (^(v54))
                                                                                                                                _ = v137
                                                                                                                            default:
                                                                                                                                v10 = (v10 % ((v10) | 1))
                                                                                                                            }
                                                                                                                        }
                                                                                                                        var v138 uint16 = v3
                                                                                                                        _ = v138
                                                                                                                    }
                                                                                                                }
                                                                                                                var v139 string = ((v4 + v4) + v4)
                                                                                                                if false {
                                                                                                                    if (fi52 < fi23) {
                                                                                                                        v68 = (v10 < 42)
                                                                                                                    } else {
                                                                                                                        switch v139 {
                                                                                                                        case "\n\t":
                                                                                                                            if false {
                                                                                                                                var v140 uint32 = v54
                                                                                                                                _ = v140
                                                                                                                            } else {
                                                                                                                                if len(v2) > 0 {
                                                                                                                                    v2[0] = (fi23 ^ 2147483647)
                                                                                                                                }
                                                                                                                                v53 -= v48
                                                                                                                            }
                                                                                                                        case "":
                                                                                                                            v141 := GenS13{F0: v19, F1: ("\n\t" + v139)}
                                                                                                                            _ = v141
                                                                                                                        case "a":
                                                                                                                            if ("hello" != ("hello" + "")) {
                                                                                                                                fmt.Println(v139, v10)
                                                                                                                                v57 += (fi23 / ((32768) | 1))
                                                                                                                            } else {
                                                                                                                                v57 += 4294967295
                                                                                                                                if false {
                                                                                                                                    if true {
                                                                                                                                        var v142 float64 = 0.1
                                                                                                                                        var v143 bool = false
                                                                                                                                        v4 = v139
                                                                                                                                        _ = v142
                                                                                                                                        _ = v143
                                                                                                                                    } else {
                                                                                                                                        fmt.Println(v37, fi23, v82)
                                                                                                                                        var v144 int = (v57 - int(v10))
                                                                                                                                        _ = v144
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    var v145 int32 = (-(v19))
                                                                                                                                    var v146 string = v139
                                                                                                                                    var v147 float64 = float64(v82)
                                                                                                                                    _ = v145
                                                                                                                                    _ = v146
                                                                                                                                    _ = v147
                                                                                                                                }
                                                                                                                                fmt.Println(v10, v57, v54)
                                                                                                                                v148 := GenS15{F0: 42}
                                                                                                                                _ = v148
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v3 += v10
                                                                                                                    }
                                                                                                                    v149 := GenS16{F0: (v67 >> ((204) & 63))}
                                                                                                                    if v68 {
                                                                                                                        var v150 uint16 = (v10 / ((v10) | 1))
                                                                                                                        v151 := GenS17{F0: 127}
                                                                                                                        switch ((v139 + v139) + (("a" + "a") + "foo bar")) {
                                                                                                                        case "\n\t":
                                                                                                                            if len(v1) >= 2 {
                                                                                                                                v1 = v1[1:len(v1)]
                                                                                                                            }
                                                                                                                        case "":
                                                                                                                            var v152 uint8 = v12
                                                                                                                            _ = v152
                                                                                                                        case "foo bar":
                                                                                                                            var v153 int16 = (-(v22))
                                                                                                                            _ = v153
                                                                                                                        default:
                                                                                                                            v22 -= (v22 & v22)
                                                                                                                        }
                                                                                                                        _ = v150
                                                                                                                        _ = v151
                                                                                                                    } else {
                                                                                                                        if (!(genFn1((v67 | -2147483648)) != (^(v53)))) {
                                                                                                                            if false {
                                                                                                                                v82 = -2.25
                                                                                                                                v154 := GenS17{F0: -1}
                                                                                                                                v2 = append(v2, (-(fi52)))
                                                                                                                                _ = v154
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v22 -= v22
                                                                                                                            var v155 uint64 = genFn1(-1)
                                                                                                                            _ = v155
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v149
                                                                                                                } else {
                                                                                                                    var v156 float32 = v82
                                                                                                                    _ = v156
                                                                                                                }
                                                                                                                _ = v139
                                                                                                            } else {
                                                                                                                v53 = (v48 + (v48 / ((0) | 1)))
                                                                                                                switch ("foo bar" + v4) {
                                                                                                                case "a":
                                                                                                                    v4 = (((v4 + "a") + "foo bar") + v4)
                                                                                                                default:
                                                                                                                    switch (fi23 % ((v57) | 1)) {
                                                                                                                    case -9223372036854775808:
                                                                                                                        v2 = append(v2, int(v19))
                                                                                                                    case 32767:
                                                                                                                        if len(v1) >= 2 {
                                                                                                                            v1 = v1[1:len(v1)]
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v157 uint8 = (v12 - 72)
                                                                                                                        _ = v157
                                                                                                                    }
                                                                                                                }
                                                                                                                var v158 float64 = -2.25
                                                                                                                if false {
                                                                                                                    var v159 int32 = 42
                                                                                                                    if len(v1) >= 2 {
                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                    }
                                                                                                                    _ = v159
                                                                                                                } else {
                                                                                                                    if v68 {
                                                                                                                        v160 := GenS5{F0: v68}
                                                                                                                        _ = v160
                                                                                                                    }
                                                                                                                    v54 = 2147483648
                                                                                                                    switch fi52 {
                                                                                                                    case 596:
                                                                                                                        if len(v2) >= 2 {
                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                        }
                                                                                                                    case 32768:
                                                                                                                        v161 := GenS14{F0: -1, F1: 221, F2: v53}
                                                                                                                        _ = v161
                                                                                                                    case -156:
                                                                                                                        v54 -= 1
                                                                                                                    }
                                                                                                                    var v162 uint8 = v12
                                                                                                                    _ = v162
                                                                                                                }
                                                                                                                _ = v158
                                                                                                            }
                                                                                                            v22 -= (v22 << ((943) & 63))
                                                                                                            v19 = int32(v54)
                                                                                                            var v163 int64 = (-(v83))
                                                                                                            _ = v102
                                                                                                            _ = v163
                                                                                                            _ = v83
                                                                                                        } else {
                                                                                                            var v164 int = (fi52 & fi52)
                                                                                                            v1 = append(v1, (fi23 >> ((758) & 63)))
                                                                                                            v53 = 391
                                                                                                            v165 := GenS4{F0: (!(v82 == 0.1)), F1: int32(v48)}
                                                                                                            if (v82 == v82) {
                                                                                                                if (255 <= (v55 & (v10 | 0))) {
                                                                                                                    fmt.Println(fi52)
                                                                                                                    v166 := GenS13{F0: v37, F1: "a"}
                                                                                                                    if v68 {
                                                                                                                        v167 := GenS18{F0: v55, F1: (v57 | v67), F2: (^(v55))}
                                                                                                                        _ = v167
                                                                                                                    }
                                                                                                                    _ = v166
                                                                                                                } else {
                                                                                                                    v48 = 18446744073709551615
                                                                                                                    var v168 string = v4
                                                                                                                    _ = v168
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v164
                                                                                                            _ = v165
                                                                                                        }
                                                                                                    }
                                                                                                    v10 += (v3 % (((v3 / (((v3 ^ uint16(v54))) | 1))) | 1))
                                                                                                    _ = v82
                                                                                                }
                                                                                            }
                                                                                        case "":
                                                                                            v37 += -1
                                                                                        default:
                                                                                            v4 = v4
                                                                                        }
                                                                                        if (v67 == v57) {
                                                                                            v1 = append(v1, (fi52 << ((2) & 63)))
                                                                                            switch v4 {
                                                                                            case "foo bar":
                                                                                                switch "a" {
                                                                                                case "hello":
                                                                                                    var v169 uint64 = (v48 + genFn1((fi52 * v57)))
                                                                                                    _ = v169
                                                                                                }
                                                                                            case "":
                                                                                                v170 := GenS1{F0: 127, F1: (v54 * 2147483648), F2: int64(v10)}
                                                                                                _ = v170
                                                                                            default:
                                                                                                fmt.Println(v68, v4, fi23)
                                                                                            }
                                                                                            v57 = fi52
                                                                                        } else {
                                                                                            var v171 int16 = (v22 / ((335) | 1))
                                                                                            if ((3.14159 + 1.5) > (0.1 / float32((3.14159 - 0.1)))) {
                                                                                                fmt.Println(fi23)
                                                                                            }
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = fi52
                                                                                            }
                                                                                            _ = len(v1)
                                                                                            if (127 <= v12) {
                                                                                                v172 := GenS19{F0: 42, F1: (^(v53)), F2: (((^(v48)) >= (^(v53))) || (!v68))}
                                                                                                switch (v57 >> ((868) & 63)) {
                                                                                                case -2147483648:
                                                                                                    v173 := GenS1{F0: -1, F1: (v54 - (v54 ^ 875)), F2: -9223372036854775808}
                                                                                                    _ = v173
                                                                                                case 4294967295:
                                                                                                    var v174 int8 = -122
                                                                                                    _ = v174
                                                                                                case 2147483647:
                                                                                                    switch ((v4 + v4) + "a") {
                                                                                                    case "hello":
                                                                                                        v19 = v19
                                                                                                    case "":
                                                                                                        v4 = v4
                                                                                                    case "\n\t":
                                                                                                        fmt.Println(v57)
                                                                                                    default:
                                                                                                        var v175 int16 = (v171 + v22)
                                                                                                        _ = v175
                                                                                                    }
                                                                                                default:
                                                                                                    var v176 string = "hello"
                                                                                                    _ = v176
                                                                                                }
                                                                                                if true {
                                                                                                    if v68 {
                                                                                                        if (1e10 >= ((((1e10 * 0.0) + 0.0) - 0.0) + float32(0.1))) {
                                                                                                            var v177 uint64 = (v48 << ((1) & 63))
                                                                                                            var v178 int = int(v37)
                                                                                                            if (v68 || false) {
                                                                                                                v179 := GenS3{F0: 255}
                                                                                                                _ = v179
                                                                                                            }
                                                                                                            _ = v177
                                                                                                            _ = v178
                                                                                                        }
                                                                                                    }
                                                                                                    if v68 {
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = v172
                                                                                            } else {
                                                                                                if (float32(((1e10 * (1.5 - float64((1.5 / ((float32((-2.25 * -2.25)) * 1e10) + 3.14159))))) / (3.14159 + (1e10 / 3.14159)))) > 1.5) {
                                                                                                    fmt.Println(v67, v68, v19)
                                                                                                    var v180 bool = v68
                                                                                                    _ = v180
                                                                                                } else {
                                                                                                    v67 = (fi52 % ((int(v12)) | 1))
                                                                                                    var v181 uint16 = v3
                                                                                                    _ = v181
                                                                                                }
                                                                                                v22 += -953
                                                                                                var v182 int32 = -2147483648
                                                                                                v54 = (v54 / ((781) | 1))
                                                                                                if (!v68) {
                                                                                                    fmt.Println(v53)
                                                                                                    if ((v67 & (v67 | v57)) > 938) {
                                                                                                        v1 = append(v1, (v57 & (v57 / ((v67) | 1))))
                                                                                                        if (uint(v55) >= uint(v55)) {
                                                                                                            var v183 int16 = 0
                                                                                                            if v68 {
                                                                                                                v184 := GenS20{F0: 255}
                                                                                                                _ = v184
                                                                                                            }
                                                                                                            _ = v183
                                                                                                        }
                                                                                                        if false {
                                                                                                            var v185 uint = 369
                                                                                                            fmt.Println(v185, v68)
                                                                                                            var v186 uint16 = v55
                                                                                                            _ = v185
                                                                                                            _ = v186
                                                                                                        }
                                                                                                    } else {
                                                                                                        v187 := GenS16{F0: v67}
                                                                                                        fmt.Println(v57)
                                                                                                        switch 0 {
                                                                                                        case -9223372036854775808:
                                                                                                            var v188 int32 = 2147483647
                                                                                                            _ = v188
                                                                                                        case 0:
                                                                                                            var v189 int64 = 9223372036854775807
                                                                                                            _ = v189
                                                                                                        case 128:
                                                                                                            v1 = append(v1, 128)
                                                                                                        default:
                                                                                                            if (2 == v55) {
                                                                                                                if (true && v68) {
                                                                                                                    var v190 int16 = (v171 | (-(v22)))
                                                                                                                    var v191 int16 = (v190 * v171)
                                                                                                                    _ = v190
                                                                                                                    _ = v191
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        v192 := GenS13{F0: 2147483647, F1: v4}
                                                                                                        v57 = v67
                                                                                                        _ = v187
                                                                                                        _ = v192
                                                                                                    }
                                                                                                }
                                                                                                _ = v182
                                                                                            }
                                                                                            v1 = append(v1, 0)
                                                                                            if v68 {
                                                                                                if (!(!v68)) {
                                                                                                    var v193 uint16 = 352
                                                                                                    _ = v193
                                                                                                } else {
                                                                                                    switch 42 {
                                                                                                    case 930:
                                                                                                        var v194 bool = v68
                                                                                                        _ = v194
                                                                                                    case 0:
                                                                                                        var v195 uint64 = v53
                                                                                                        _ = v195
                                                                                                    }
                                                                                                    v196 := GenS17{F0: -80}
                                                                                                    v197 := GenS12{F0: (18446744073709551615 >= 1), F1: (v3 < (v55 | (v10 >> ((42) & 63))))}
                                                                                                    _ = v196
                                                                                                    _ = v197
                                                                                                }
                                                                                            } else {
                                                                                                _ = len(v1)
                                                                                                var v198 uint16 = v3
                                                                                                _ = v198
                                                                                            }
                                                                                            _ = v171
                                                                                        }
                                                                                        if ((!(int(v12) > -2147483648)) || v68) {
                                                                                            fmt.Println(v67, v4)
                                                                                            v19 += v37
                                                                                        }
                                                                                        fmt.Println(v55, v57, v22)
                                                                                        if ((1.5 > (1.5 * (1.5 + 0.0))) && v68) {
                                                                                            _ = len(v2)
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = fi23
                                                                                            }
                                                                                            fmt.Println(fi52, v37, v3)
                                                                                        } else {
                                                                                            if v68 {
                                                                                                switch (v67 >> ((225) & 63)) {
                                                                                                case -1:
                                                                                                    if false {
                                                                                                        switch v67 {
                                                                                                        case 402:
                                                                                                            v199 := GenS4{F0: (966 < -1), F1: v19}
                                                                                                            _ = v199
                                                                                                        }
                                                                                                    } else {
                                                                                                        v57 = (v57 + fi23)
                                                                                                        fmt.Println(v37)
                                                                                                        var v200 uint = 986
                                                                                                        _ = v200
                                                                                                    }
                                                                                                case -352:
                                                                                                    var v201 int8 = 1
                                                                                                    _ = v201
                                                                                                case 127:
                                                                                                    v202 := GenS21{F0: -154}
                                                                                                    _ = v202
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v12, v19, v37)
                                                                                            }
                                                                                            v203 := GenS22{F0: 480, F1: (^(v53))}
                                                                                            switch v4 {
                                                                                            case "a":
                                                                                                fmt.Println(fi52, v37)
                                                                                            default:
                                                                                                v2 = append(v2, 128)
                                                                                            }
                                                                                            _ = v203
                                                                                        }
                                                                                    }
                                                                                    v204 := GenS5{F0: v68}
                                                                                    switch (v67 >> ((65535) & 63)) {
                                                                                    case 1:
                                                                                        switch v4 {
                                                                                        case "":
                                                                                            switch ("a" + v4) {
                                                                                            case "foo bar":
                                                                                                fmt.Println(v67, v12)
                                                                                            case "hello":
                                                                                                v205 := GenS11{F0: 100, F1: (v48 == (^(v53))), F2: 4294967295}
                                                                                                _ = v205
                                                                                            }
                                                                                        case "\n\t":
                                                                                            var v206 int16 = v22
                                                                                            _ = v206
                                                                                        case "foo bar":
                                                                                            v207 := GenS6{F0: v4, F1: 42, F2: -681}
                                                                                            _ = v207
                                                                                        default:
                                                                                            v4 = (v4 + "hello")
                                                                                        }
                                                                                    case -557:
                                                                                        v208 := GenS23{F0: 42, F1: 1}
                                                                                        _ = v208
                                                                                    default:
                                                                                        var v209 int16 = (v22 | v22)
                                                                                        _ = v209
                                                                                    }
                                                                                    fmt.Println(v55, v57, v3)
                                                                                    v210 := GenS22{F0: v37, F1: (v48 * v48)}
                                                                                    _ = v204
                                                                                    _ = v210
                                                                                }
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                        } else {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            switch (v57 ^ v57) {
                                                                            case -993:
                                                                                v211 := GenS6{F0: "", F1: 904, F2: 42}
                                                                                switch ((((v4 + (v4 + v4)) + v4) + "\n\t") + "a") {
                                                                                case "":
                                                                                    v212 := GenS24{F0: 207, F1: ((v4 + v4) + v4), F2: v54}
                                                                                    _ = v212
                                                                                case "\n\t":
                                                                                    if ((v37 % (((v19 - (v19 * (v19 & (v37 - int32(v57)))))) | 1)) <= (v19 >> ((13) & 63))) {
                                                                                        var v213 float32 = float32(1e10)
                                                                                        _ = v213
                                                                                    }
                                                                                }
                                                                                _ = v211
                                                                            case 128:
                                                                                v214 := GenS19{F0: int16(v55), F1: 786, F2: (!(0.0 <= 1.5))}
                                                                                _ = v214
                                                                            case 32767:
                                                                                v22 -= (v22 & (v22 & (^(v22))))
                                                                            default:
                                                                                fmt.Println(fi52)
                                                                                if (v4 <= v4) {
                                                                                    switch v4 {
                                                                                    case "\n\t":
                                                                                        if true {
                                                                                            fmt.Println(v55)
                                                                                            v12 -= genFn0((v12 >> ((290) & 63)))
                                                                                        } else {
                                                                                            v22 -= (v22 & (v22 - v22))
                                                                                            var v215 uint32 = (v54 + 65535)
                                                                                            _ = v215
                                                                                        }
                                                                                    case "":
                                                                                        v216 := GenS25{F0: (v22 >> ((0) & 63)), F1: int(v54), F2: (v12 - genFn0(109))}
                                                                                        _ = v216
                                                                                    default:
                                                                                        var v217 int16 = (v22 ^ (v22 | v22))
                                                                                        _ = v217
                                                                                    }
                                                                                    v57 += (v57 / ((int(v12)) | 1))
                                                                                    var v218 uint = uint(v55)
                                                                                    _ = v218
                                                                                } else {
                                                                                    var v219 int32 = v19
                                                                                    fmt.Println(fi52, v37)
                                                                                    _ = v219
                                                                                }
                                                                                switch "\n\t" {
                                                                                case "\n\t":
                                                                                    v55 = v3
                                                                                default:
                                                                                    var v220 int8 = 42
                                                                                    _ = v220
                                                                                }
                                                                                var v221 uint16 = (v3 - 0)
                                                                                var v222 string = ""
                                                                                _ = v221
                                                                                _ = v222
                                                                            }
                                                                            var v223 uint8 = uint8(v55)
                                                                            _ = v223
                                                                        }
                                                                        _ = v57
                                                                    case 946:
                                                                        var v224 int = 0
                                                                        _ = v224
                                                                    case -1:
                                                                        var v225 float64 = 0.0
                                                                        _ = v225
                                                                    }
                                                                }
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                            v226 := GenS7{F0: 9223372036854775807}
                                                            _ = v226
                                                        }
                                                        _ = fi52
                                                        _ = v53
                                                    }
                                                    var v227 float64 = (1.5 / ((float64((3.14159 / 1.5)) + (0.1 / -2.25)) - 0.1))
                                                    _ = v227
                                                }
                                                if true {
                                                    v37 -= v19
                                                }
                                            }
                                        case "":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        default:
                                            switch (v4 + (v4 + "")) {
                                            case "foo bar":
                                                for fi228 := 0; fi228 < 1; fi228++ {
                                                    if ((0.1 / (-2.25 - 0.0)) <= float32((float64(float32(3.14159)) / ((0.1 - 1.5) / float64((float32((((-2.25 / -2.25) / (0.0 + 1e10)) / float64(0.1))) * float32(-2.25))))))) {
                                                        if ((fi228 == (fi23 ^ -428)) || true) {
                                                            fmt.Println(v12, v12)
                                                            var v229 uint32 = 661
                                                            switch ((v4 + (v4 + (v4 + v4))) + v4) {
                                                            case "hello":
                                                                var v230 bool = (!(836 >= -990))
                                                                _ = v230
                                                            case "foo bar":
                                                                var v231 float32 = ((1e10 / (((1e10 + (0.0 + float32((float64(1e10) * -2.25)))) + (0.0 * 1e10)) / (3.14159 + 1.5))) * (0.1 * 3.14159))
                                                                _ = v231
                                                            case "\n\t":
                                                                v232 := GenS13{F0: v19, F1: (v4 + v4)}
                                                                _ = v232
                                                            }
                                                            v22 -= -32768
                                                            _ = v229
                                                        } else {
                                                            v233 := GenS2{F0: false}
                                                            switch fi23 {
                                                            case -9223372036854775808:
                                                                if (!false) {
                                                                    v3 -= v10
                                                                    if true {
                                                                        v234 := GenS13{F0: -1, F1: v4}
                                                                        var v235 uint64 = v48
                                                                        _ = v234
                                                                        _ = v235
                                                                    } else {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 2147483647
                                                                        }
                                                                    }
                                                                }
                                                                v48 = genFn1(fi228)
                                                            case 127:
                                                                if (617 != 174) {
                                                                    v37 -= v19
                                                                    var v236 int16 = 1000
                                                                    v2 = append(v2, fi228)
                                                                    fmt.Println(v12, fi23)
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    _ = v236
                                                                }
                                                            }
                                                            _ = v233
                                                        }
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = (^(fi228))
                                                        }
                                                    }
                                                    v237 := GenS26{F0: 127, F1: 1, F2: 65535}
                                                    _ = fi228
                                                    _ = v237
                                                }
                                            case "\n\t":
                                                v10 -= (v3 + (v10 * 2))
                                            case "":
                                                v238 := GenS18{F0: (^(v3)), F1: int(v3), F2: 2}
                                                _ = v238
                                            default:
                                                for fi239 := 0; fi239 < 6; fi239++ {
                                                    switch 32767 {
                                                    case 32767:
                                                        if len(v1) > 0 {
                                                            v1[0] = 32768
                                                        }
                                                    case 0:
                                                        var v240 uint32 = uint32(v48)
                                                        _ = v240
                                                    case 32768:
                                                        if len(v2) > 0 {
                                                            v2[0] = 42
                                                        }
                                                    }
                                                    v4 = v4
                                                    if len(v1) > 0 {
                                                        v1[0] = fi239
                                                    }
                                                    var v241 string = (v4 + ("\n\t" + v4))
                                                    _ = fi239
                                                    _ = v241
                                                }
                                            }
                                        }
                                        _ = v48
                                    }
                                }
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                        _ = v37
                    }
                } else {
                    fmt.Println(v22)
                    var v242 float64 = -2.25
                    if ((v4 + v4) < "foo bar") {
                        if false {
                            v3 += v3
                            for fi243 := 0; fi243 < 15; fi243++ {
                                if true {
                                    v244 := GenS25{F0: (v22 & (^(v22))), F1: (-(fi23)), F2: genFn0(1)}
                                    if (fi23 <= fi23) {
                                        fmt.Println(v10, v3)
                                    }
                                    v3 = v10
                                    _ = v244
                                }
                                if len(v1) > 0 {
                                    v1[0] = -514
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi243
                            }
                        } else {
                            v22 += -32768
                        }
                        var v245 float64 = (((float64(0.0) + ((1e10 * 1.5) + v242)) * v242) - ((float64((0.1 - 3.14159)) / v242) / v242))
                        _ = v245
                    }
                    _ = v242
                }
            }
            v3 = (v3 & 0)
            fmt.Println(v4)
        }
        _ = fi23
    }
    if true {
        if false {
            if true {
                var v246 string = ""
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v246
            }
            v247 := GenS19{F0: (v22 / ((-1) | 1)), F1: 4294967295, F2: true}
            v248 := GenS24{F0: v12, F1: v4, F2: uint32(v12)}
            _ = v247
            _ = v248
        }
        var v249 int8 = -1
        _ = len(v1)
        _ = v249
    }
    var v250 uint8 = (^(v12))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v250:", v250)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v250
    _ = v3
    _ = v4
}
