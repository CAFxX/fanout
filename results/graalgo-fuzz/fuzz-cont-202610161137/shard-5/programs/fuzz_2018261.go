package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int16; F1 uint64; F2 uint }
type GenS10 struct { F0 int16; F1 int8; F2 int16 }
type GenS11 struct { F0 int32; F1 uint32 }
type GenS12 struct { F0 string; F1 uint8 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int; F1 int64; F2 uint32 }
type GenS18 struct { F0 uint8; F1 uint; F2 int32 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 int; F1 uint8; F2 int }
func genFn0(p0 uint64, p1 uint16) uint64 {
    var v5 uint64 = p0
    switch 2147483647 {
    case 127:
        p0 += v5
    case 222:
        var v6 int = 866
        _ = v6
    case 4294967295:
        var v7 int16 = -21
        _ = v7
    default:
        var v8 int32 = 2147483647
        _ = v8
    }
    var v9 int = -132
    p0 += p0
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 0
}

func main() {
    var v1 string = "a"
    var v2 float32 = (0.1 / (1.5 * -2.25))
    v3 := make(map[int]string)
    var v4 uint16 = 0
    v10 := []int{9223372036854775807, -2147483648, 127, 9223372036854775807}
    v4 = v4
    v4 = v4
    var v11 uint16 = (v4 & v4)
    var v12 int16 = int16(v4)
    if false {
        v1 = v1
    } else {
        if true {
            for fi13 := 0; fi13 < 8; fi13++ {
                fmt.Println(fi13, v12, v12)
                _ = fi13
            }
            v12 = v12
        }
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 uint32 = 2
        for fi16 := 0; fi16 < 20; fi16++ {
            if false {
                var v17 float32 = 0.1
                if ((18446744073709551615 < 255) || (true || true)) {
                    var v18 uint16 = v4
                    switch fi14 {
                    case -62:
                        if true {
                            _ = len(v10)
                        } else {
                            delete(v3, (fi16 >> ((65535) & 63)))
                        }
                    case 729:
                        var v19 string = v1
                        _ = v19
                    case 128:
                        if true {
                            var v20 uint64 = 864
                            if len(v10) > 0 {
                                v10[0] = fi16
                            }
                            v18 = 255
                            _ = v20
                        }
                    default:
                        fmt.Println(v18)
                    }
                    if false {
                        var v21 float32 = v2
                        v22 := GenS1{F0: false, F1: fi16}
                        _ = v21
                        _ = v22
                    } else {
                        var v23 string = v1
                        var v24 int32 = 1
                        if (-221 < (v12 * v12)) {
                            var v25 uint32 = (v15 - 4294967295)
                            v17 -= v17
                            var v26 int = fi14
                            _ = v25
                            _ = v26
                        }
                        _ = v23
                        _ = v24
                    }
                    _ = v18
                } else {
                    v1 = "foo bar"
                    if false {
                        v1 = ((v1 + ("foo bar" + (v1 + v1))) + (v1 + (v1 + (v1 + v1))))
                        v17 += v2
                    } else {
                        v10 = append(v10, (fi14 & (-(fi14))))
                        if false {
                            var v27 int32 = -1
                            _ = v27
                        } else {
                            switch ("\n\t" + v1) {
                            case "a":
                                if true {
                                    var v28 int32 = -236
                                    fmt.Println(v15, v4, v2)
                                    _ = v28
                                } else {
                                    var v29 int32 = -330
                                    var v30 uint16 = 1
                                    if len(v10) > 0 {
                                        v10[0] = fi14
                                    }
                                    v15 += (v15 * uint32(v30))
                                    switch fi14 {
                                    case 127:
                                        switch -192 {
                                        case 42:
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                        }
                                        var v31 bool = (4294967295 < 741)
                                        _ = v31
                                    case -990:
                                        v12 = (v12 >> ((42) & 63))
                                    case -996:
                                        v30 = 42
                                    default:
                                        fmt.Println(v12)
                                    }
                                    var v32 float64 = (3.14159 / ((0.1 + ((0.1 / 1e10) * (3.14159 - (1.5 + -2.25)))) - 3.14159))
                                    _ = v29
                                    _ = v30
                                    _ = v32
                                }
                            default:
                                fmt.Println(fi14)
                            }
                            var v33 int8 = 1
                            if true {
                                switch v1 {
                                case "foo bar":
                                    _ = len(v3)
                                case "\n\t":
                                    fmt.Println(v33, v2)
                                case "a":
                                    fmt.Println(v15)
                                default:
                                    if (!false) {
                                        var v34 int8 = (v33 >> ((734) & 63))
                                        _ = v34
                                    }
                                }
                                fmt.Println(v17, fi16)
                                if false {
                                    v10 = append(v10, (fi14 << ((1) & 63)))
                                    if (!true) {
                                        switch v1 {
                                        case "a":
                                            var v35 int64 = int64(fi14)
                                            _ = v35
                                        }
                                        if true {
                                            var v36 int8 = v33
                                            if (!((!(false && false)) || false)) {
                                                v17 = (v17 + v2)
                                                fmt.Println(v2, v17, v17)
                                                var v37 int32 = -1
                                                _ = v37
                                            }
                                            v3[fi14] = ("\n\t" + v1)
                                            if ((((v2 / v2) - v17) >= 3.14159) && ((float32(((0.0 / (0.1 * (-2.25 * 1.5))) * 0.0)) + (v17 / (3.14159 - float32(((1e10 * float64((-2.25 - v2))) / (float64(1.5) - (1e10 - -2.25))))))) >= (v2 / v17))) {
                                                v4 = (v4 << ((42) & 63))
                                                var v38 uint = 18446744073709551615
                                                switch v1 {
                                                case "foo bar":
                                                    if false {
                                                        var v39 int8 = (v36 ^ 127)
                                                        if (1 <= -9223372036854775808) {
                                                            v39 -= 42
                                                            fmt.Println(v1, v15)
                                                            if false {
                                                                if true {
                                                                    _ = len(v3)
                                                                }
                                                            } else {
                                                                v10 = append(v10, (fi16 / ((fi16) | 1)))
                                                                v40 := GenS1{F0: ((fi16 - (^(fi16))) > (fi14 | (fi14 ^ fi14))), F1: fi14}
                                                                v17 = v2
                                                                switch (fi16 | (fi14 << ((v38) & 63))) {
                                                                case -521:
                                                                    if ((v15 * 1) >= v15) {
                                                                        v41 := GenS1{F0: false, F1: (fi16 ^ (fi14 + 9223372036854775807))}
                                                                        _ = len(v3)
                                                                        v36 -= v36
                                                                        _ = v41
                                                                    } else {
                                                                        v42 := GenS2{F0: (fi16 + fi16)}
                                                                        switch (fi16 & -520) {
                                                                        case 128:
                                                                            v33 = (v33 / ((v33) | 1))
                                                                        case 2147483647:
                                                                            var v43 int64 = int64(v15)
                                                                            _ = v43
                                                                        case 0:
                                                                            if false {
                                                                                switch fi16 {
                                                                                case -9223372036854775808:
                                                                                    v15 += (v15 << ((2) & 63))
                                                                                case 308:
                                                                                    var v44 uint8 = 127
                                                                                    _ = v44
                                                                                case 878:
                                                                                    if (0.0 > 0.1) {
                                                                                        var v45 string = (v1 + (("hello" + ("foo bar" + v1)) + "\n\t"))
                                                                                        var v46 int8 = (v36 >> (((v38 | v38)) & 63))
                                                                                        _ = v45
                                                                                        _ = v46
                                                                                    } else {
                                                                                        v47 := GenS3{F0: -1}
                                                                                        if len(v10) > 0 {
                                                                                            v10[0] = (fi14 ^ (fi16 << ((42) & 63)))
                                                                                        }
                                                                                        if false {
                                                                                            v48 := GenS4{F0: true}
                                                                                            v10 = append(v10, fi14)
                                                                                            _ = v48
                                                                                        } else {
                                                                                            v33 -= v36
                                                                                        }
                                                                                        _ = v47
                                                                                    }
                                                                                default:
                                                                                    if false {
                                                                                        switch (((v1 + (v1 + v1)) + "\n\t") + v1) {
                                                                                        case "foo bar":
                                                                                            delete(v3, fi14)
                                                                                        case "hello":
                                                                                            delete(v3, 42)
                                                                                        default:
                                                                                            var v49 float64 = 0.1
                                                                                            _ = v49
                                                                                        }
                                                                                        if (0.0 != float64((float32(1.5) + 0.0))) {
                                                                                            var v50 float32 = float32(((0.1 + 3.14159) + (1e10 + -2.25)))
                                                                                            var v51 int = (fi16 * (fi16 - 127))
                                                                                            _ = v50
                                                                                            _ = v51
                                                                                        } else {
                                                                                            switch (fi16 * fi14) {
                                                                                            case 0:
                                                                                                switch -9223372036854775808 {
                                                                                                case -9223372036854775808:
                                                                                                    v15 += 2147483648
                                                                                                case -1:
                                                                                                    var v52 uint16 = v11
                                                                                                    var v53 int8 = 1
                                                                                                    _ = v52
                                                                                                    _ = v53
                                                                                                case 32767:
                                                                                                    v54 := GenS5{F0: 42, F1: v33}
                                                                                                    v1 = v1
                                                                                                    _ = v54
                                                                                                default:
                                                                                                    v55 := GenS2{F0: (fi16 & (-(fi16)))}
                                                                                                    _ = v55
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v1)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if len(v10) > 0 {
                                                                                    v10[0] = fi14
                                                                                }
                                                                            } else {
                                                                                switch (^(fi14)) {
                                                                                case 127:
                                                                                    v56 := GenS6{F0: (v11 & (v11 << ((255) & 63))), F1: (^(v4))}
                                                                                    _ = v56
                                                                                }
                                                                                switch fi14 {
                                                                                case 128:
                                                                                    v33 -= -128
                                                                                }
                                                                                var v57 int = (fi16 * (fi16 & fi14))
                                                                                var v58 int = fi14
                                                                                _ = v57
                                                                                _ = v58
                                                                            }
                                                                        }
                                                                        v1 = v1
                                                                        _ = v42
                                                                    }
                                                                }
                                                                if len(v10) > 0 {
                                                                    v10[0] = 42
                                                                }
                                                                _ = v40
                                                            }
                                                            fmt.Println(fi16, v38)
                                                            if (240 > 0) {
                                                                fmt.Println(v38)
                                                                var v59 uint32 = 818
                                                                v60 := GenS3{F0: 2147483647}
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                            if ((uint64(v38) != genFn0(genFn0(18446744073709551615, (v4 | (v11 & (v11 + uint16(v15))))), v4)) && ((!(false || true)) && ((-980 < 1) || false))) {
                                                                var v61 int64 = 42
                                                                _ = v61
                                                            } else {
                                                                if (!true) {
                                                                    var v62 int32 = 42
                                                                    v2 += float32(1.5)
                                                                    _ = v62
                                                                }
                                                                if len(v10) >= 2 {
                                                                    v10 = v10[1:len(v10)]
                                                                }
                                                            }
                                                            switch "hello" {
                                                            case "hello":
                                                                switch fi14 {
                                                                case 4294967295:
                                                                    v36 = (v39 + (v36 % ((v36) | 1)))
                                                                case 42:
                                                                    fmt.Println(v17, v17, v12)
                                                                default:
                                                                    if true {
                                                                        switch fi14 {
                                                                        case 48:
                                                                            var v63 uint32 = 830
                                                                            _ = v63
                                                                        case 294:
                                                                            var v64 uint32 = (v15 >> ((v38) & 63))
                                                                            _ = v64
                                                                        case 128:
                                                                            var v65 uint8 = 2
                                                                            _ = v65
                                                                        default:
                                                                            v38 += v38
                                                                        }
                                                                        v66 := v3[(-(fi14))]
                                                                        if false {
                                                                            if (int32(v33) > 42) {
                                                                                var v67 uint8 = 128
                                                                                _ = v67
                                                                            } else {
                                                                                fmt.Println(v38)
                                                                            }
                                                                        }
                                                                        _ = v66
                                                                    }
                                                                }
                                                                var v68 int8 = (v39 & (-(v36)))
                                                                v11 += (v4 % ((v4) | 1))
                                                                _ = v68
                                                            case "":
                                                                v17 -= 1.5
                                                            case "\n\t":
                                                                var v69 string = "hello"
                                                                _ = v69
                                                            }
                                                            v70 := GenS2{F0: 128}
                                                            fmt.Println(fi16, v36, v39)
                                                            v71 := GenS1{F0: (869 != -9223372036854775808), F1: fi14}
                                                            v10 = append(v10, fi16)
                                                            v36 -= v39
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        if len(v10) > 0 {
                                                            v10[0] = fi14
                                                        }
                                                        var v72 string = v1
                                                        _ = v39
                                                        _ = v72
                                                    } else {
                                                        _ = len(v10)
                                                    }
                                                case "\n\t":
                                                    v12 = v12
                                                default:
                                                    v73 := GenS6{F0: (v11 & v11), F1: v4}
                                                    _ = v73
                                                }
                                                _ = v38
                                            } else {
                                                v74 := GenS7{F0: v36}
                                                v75 := GenS8{F0: 4294967295}
                                                v33 -= (v33 ^ (^(v36)))
                                                _ = v74
                                                _ = v75
                                            }
                                            if true {
                                                v76 := GenS3{F0: 1}
                                                var v77 int64 = int64(v36)
                                                if true {
                                                    switch int(v33) {
                                                    case -1:
                                                        var v78 uint64 = 4294967295
                                                        _ = v78
                                                    case 128:
                                                        v79 := GenS4{F0: (!((fi14 | 80) <= fi16))}
                                                        _ = v79
                                                    default:
                                                        _ = len(v10)
                                                    }
                                                    if (!true) {
                                                        var v80 uint16 = (v4 % ((1) | 1))
                                                        _ = v80
                                                    } else {
                                                        if false {
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                        }
                                                        v81 := GenS9{F0: v12, F1: 902, F2: 0}
                                                        switch v1 {
                                                        case "foo bar":
                                                            fmt.Println(v1, v15, v2)
                                                        case "":
                                                            var v82 uint = 42
                                                            _ = v82
                                                        case "a":
                                                            fmt.Println(v4, v2)
                                                        }
                                                        _ = v81
                                                    }
                                                    var v83 uint32 = (v15 & (v15 | (^(v15))))
                                                    _ = v83
                                                } else {
                                                    fmt.Println(v1, v1, v77)
                                                    if ((v77 | (v77 + 754)) >= v77) {
                                                        _ = len(v3)
                                                    } else {
                                                        switch (fi16 ^ fi16) {
                                                        case 32768:
                                                            if ((v33 + (v36 | 0)) <= -14) {
                                                                v84 := GenS10{F0: v12, F1: v36, F2: v12}
                                                                var v85 uint64 = 42
                                                                if false {
                                                                    v86 := GenS9{F0: int16(v77), F1: genFn0((^(v85)), v4), F2: 18446744073709551615}
                                                                    _ = v86
                                                                } else {
                                                                    var v87 uint8 = 240
                                                                    if (-1 < 1) {
                                                                        v87 = 2
                                                                    } else {
                                                                        switch fi14 {
                                                                        case 1:
                                                                            var v88 int = 4294967295
                                                                            _ = v88
                                                                        default:
                                                                            fmt.Println(v12, v12)
                                                                        }
                                                                        v3[(fi16 / ((-1) | 1))] = v1
                                                                    }
                                                                    v4 = (v11 << ((731) & 63))
                                                                    _ = v87
                                                                }
                                                                if (42 > 42) {
                                                                    var v89 uint = 656
                                                                    v90 := GenS2{F0: fi14}
                                                                    _ = v89
                                                                    _ = v90
                                                                }
                                                                _ = v84
                                                                _ = v85
                                                            }
                                                        default:
                                                            if (1 > uint8(v4)) {
                                                                var v91 bool = ((-(fi16)) < fi16)
                                                                fmt.Println(v91, fi14)
                                                                _ = v91
                                                            }
                                                        }
                                                    }
                                                    switch fi16 {
                                                    case 32768:
                                                        v12 += v12
                                                    case -184:
                                                        v11 = v11
                                                    default:
                                                        v12 = (^(v12))
                                                    }
                                                }
                                                _ = v76
                                                _ = v77
                                            }
                                            _ = v36
                                        }
                                        fmt.Println(fi14, v12)
                                    } else {
                                        v11 += (v11 * 1)
                                        fmt.Println(v1, fi14, v15)
                                        if false {
                                            fmt.Println(v17, fi16, fi16)
                                            _ = len(v3)
                                            _ = len(v3)
                                        }
                                        var v92 uint64 = 232
                                        if (!(!false)) {
                                            fmt.Println(v2, fi16, v92)
                                            v17 -= v2
                                            v4 = (v11 / (((v4 + v4)) | 1))
                                            switch (fi14 << ((666) & 63)) {
                                            case 4294967295:
                                                var v93 int = fi14
                                                _ = v93
                                            case 128:
                                                v92 += (v92 - v92)
                                            case -847:
                                                fmt.Println(v11)
                                            }
                                            if false {
                                                if (-1 <= 42) {
                                                    v94 := GenS11{F0: -850, F1: 287}
                                                    v2 = (v17 - v17)
                                                    _ = v94
                                                } else {
                                                    var v95 float32 = ((v17 - (1e10 * (float32(3.14159) - (v17 / 1e10)))) - (v17 - v2))
                                                    var v96 bool = false
                                                    _ = v95
                                                    _ = v96
                                                }
                                            } else {
                                                v97 := GenS5{F0: -9223372036854775808, F1: -51}
                                                if (!true) {
                                                    v11 = v11
                                                    v15 = (v15 + 0)
                                                    if false {
                                                        _ = len(v10)
                                                        var v98 float64 = 3.14159
                                                        var v99 uint16 = v4
                                                        if len(v10) >= 2 {
                                                            v10 = v10[1:len(v10)]
                                                        }
                                                        fmt.Println(fi14)
                                                        _ = v98
                                                        _ = v99
                                                    } else {
                                                        if (!true) {
                                                            v2 = v2
                                                        }
                                                        v100 := GenS9{F0: (-(v12)), F1: 1, F2: 18446744073709551615}
                                                        _ = v100
                                                    }
                                                } else {
                                                    var v101 int64 = 0
                                                    _ = v101
                                                }
                                                fmt.Println(v11, fi14, v15)
                                                v33 += (^(v33))
                                                if false {
                                                    if true {
                                                        v15 = v15
                                                    } else {
                                                        if true {
                                                            fmt.Println(v17, v11, v17)
                                                            var v102 float32 = (v17 + 0.0)
                                                            v4 -= (^(v4))
                                                            if ((947 >= 18446744073709551615) || true) {
                                                                v103 := v3[-602]
                                                                v104 := v3[(^(fi16))]
                                                                _ = v103
                                                                _ = v104
                                                            }
                                                            switch fi16 {
                                                            case 42:
                                                                v17 = 0.1
                                                            case -2147483648:
                                                                if true {
                                                                    delete(v3, (fi16 % ((-504) | 1)))
                                                                    v102 -= (-2.25 / (v2 - v17))
                                                                    v33 = v33
                                                                } else {
                                                                    switch (v1 + (v1 + "")) {
                                                                    case "foo bar":
                                                                        fmt.Println(v33, v2)
                                                                    case "a":
                                                                        v33 = v33
                                                                    }
                                                                    switch fi14 {
                                                                    case -1:
                                                                        var v105 float32 = float32(0.0)
                                                                        _ = v105
                                                                    case 2147483647:
                                                                        if false {
                                                                            v1 = ((("foo bar" + "foo bar") + v1) + "a")
                                                                            var v106 string = ((v1 + (("\n\t" + "hello") + ((((v1 + v1) + ((("foo bar" + "foo bar") + v1) + ((("\n\t" + v1) + v1) + "\n\t"))) + v1) + v1))) + (("\n\t" + (v1 + ((((("a" + ("foo bar" + "a")) + "a") + "foo bar") + "\n\t") + "\n\t"))) + ("" + v1)))
                                                                            switch (((v1 + v106) + v1) + (v106 + "hello")) {
                                                                            case "foo bar":
                                                                                v107 := GenS8{F0: uint64(v15)}
                                                                                _ = v107
                                                                            case "\n\t":
                                                                                var v108 uint = 0
                                                                                _ = v108
                                                                            case "a":
                                                                                _ = len(v3)
                                                                            }
                                                                            _ = v106
                                                                        } else {
                                                                            if true {
                                                                                var v109 int = (fi16 * fi14)
                                                                                _ = v109
                                                                            }
                                                                        }
                                                                    case 423:
                                                                        var v110 float64 = float64(((0.1 * v2) * v102))
                                                                        _ = v110
                                                                    default:
                                                                        v111 := GenS12{F0: "\n\t", F1: 128}
                                                                        _ = v111
                                                                    }
                                                                    v4 = v11
                                                                    v12 += (v12 - v12)
                                                                }
                                                            case 115:
                                                                v1 = (v1 + v1)
                                                            }
                                                            _ = v102
                                                        } else {
                                                            if (!((!false) && (2 == v11))) {
                                                                if (!((!((^(v15)) > 2147483647)) || (42 < 215))) {
                                                                    v112 := GenS13{F0: v33}
                                                                    if (!true) {
                                                                        if true {
                                                                            v113 := GenS14{F0: -343}
                                                                            v114 := GenS15{F0: (!(uint16(v15) > v4))}
                                                                            if false {
                                                                                var v115 uint = 744
                                                                                v92 = (v92 | genFn0(v92, (v4 % ((30) | 1))))
                                                                                _ = v115
                                                                            } else {
                                                                                var v116 uint64 = v92
                                                                                _ = len(v3)
                                                                                _ = v116
                                                                            }
                                                                            if ((int32(v33) <= -2147483648) && true) {
                                                                                switch ("hello" + ("\n\t" + ("\n\t" + v1))) {
                                                                                case "":
                                                                                    if (-821 <= 42) {
                                                                                        v117 := v3[-2147483648]
                                                                                        if false {
                                                                                            var v118 uint32 = v15
                                                                                            if true {
                                                                                                if true {
                                                                                                    v3[(fi16 >> ((2) & 63))] = (v117 + v1)
                                                                                                    fmt.Println(v92)
                                                                                                    switch v117 {
                                                                                                    case "\n\t":
                                                                                                        var v119 float64 = 3.14159
                                                                                                        _ = v119
                                                                                                    case "foo bar":
                                                                                                        var v120 float64 = 0.1
                                                                                                        _ = v120
                                                                                                    }
                                                                                                    var v121 int8 = v33
                                                                                                    _ = v121
                                                                                                }
                                                                                                v122 := GenS8{F0: v92}
                                                                                                switch (^(fi14)) {
                                                                                                case 32768:
                                                                                                    delete(v3, 618)
                                                                                                case 2147483647:
                                                                                                    fmt.Println(v11, v17, v117)
                                                                                                case 32767:
                                                                                                    v123 := GenS9{F0: v12, F1: v92, F2: 65535}
                                                                                                    _ = v123
                                                                                                }
                                                                                                var v124 int64 = 806
                                                                                                _ = v122
                                                                                                _ = v124
                                                                                            }
                                                                                            _ = v118
                                                                                        }
                                                                                        _ = v117
                                                                                    } else {
                                                                                        if (1 < (v33 ^ -1)) {
                                                                                            v12 = (v12 + (v12 + int16(v92)))
                                                                                            v125 := v3[fi16]
                                                                                            _ = v125
                                                                                        } else {
                                                                                            v126 := GenS10{F0: (v12 - v12), F1: 0, F2: v12}
                                                                                            _ = v126
                                                                                        }
                                                                                        if true {
                                                                                            v4 = v11
                                                                                            v33 -= -16
                                                                                        } else {
                                                                                            var v127 int = 459
                                                                                            if len(v10) > 0 {
                                                                                                v10[0] = v127
                                                                                            }
                                                                                            v3[(-(fi16))] = (v1 + "hello")
                                                                                            _ = v127
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v128 := GenS16{F0: true}
                                                                                    _ = v128
                                                                                }
                                                                            }
                                                                            v2 = (0.1 * ((((v2 * 0.0) * v2) - v17) * (v2 + v17)))
                                                                            _ = v113
                                                                            _ = v114
                                                                        }
                                                                    } else {
                                                                        v10 = append(v10, (fi16 << ((704) & 63)))
                                                                        var v129 float64 = 0.0
                                                                        v4 = uint16(v15)
                                                                        _ = v129
                                                                    }
                                                                    _ = v112
                                                                }
                                                                if (842 >= 2) {
                                                                    var v130 uint16 = (^(v11))
                                                                    v131 := GenS17{F0: (fi16 / ((int(v15)) | 1)), F1: -1, F2: v15}
                                                                    var v132 float64 = (1.5 - 0.1)
                                                                    switch -2147483648 {
                                                                    case 2147483647:
                                                                        switch fi14 {
                                                                        case 4294967295:
                                                                            v33 += v33
                                                                        default:
                                                                            var v133 float32 = v2
                                                                            _ = v133
                                                                        }
                                                                    }
                                                                    _ = v130
                                                                    _ = v131
                                                                    _ = v132
                                                                }
                                                            }
                                                            if false {
                                                                var v134 uint16 = (v11 % ((v4) | 1))
                                                                if (152 <= 1) {
                                                                    v135 := v3[(fi16 / ((fi16) | 1))]
                                                                    var v136 int32 = 836
                                                                    if len(v10) >= 2 {
                                                                        v10 = v10[1:len(v10)]
                                                                    }
                                                                    v17 = v2
                                                                    var v137 float32 = 1.5
                                                                    switch v1 {
                                                                    case "":
                                                                        switch (^(fi14)) {
                                                                        case -193:
                                                                            v135 = v135
                                                                        }
                                                                    case "a":
                                                                        v138 := GenS4{F0: (true && true)}
                                                                        _ = v138
                                                                    case "foo bar":
                                                                        v139 := GenS18{F0: 47, F1: 0, F2: (-(v136))}
                                                                        _ = v139
                                                                    }
                                                                    _ = v135
                                                                    _ = v136
                                                                    _ = v137
                                                                }
                                                                _ = v134
                                                            }
                                                            if (!true) {
                                                                var v140 float64 = 1.5
                                                                _ = v140
                                                            }
                                                        }
                                                    }
                                                    switch (fi16 % (((fi14 >> ((0) & 63))) | 1)) {
                                                    case 32768:
                                                        switch v1 {
                                                        case "foo bar":
                                                            _ = len(v3)
                                                        case "\n\t":
                                                            delete(v3, fi14)
                                                        default:
                                                            if true {
                                                                switch v1 {
                                                                case "hello":
                                                                    fmt.Println(v17, v33, v12)
                                                                }
                                                                v141 := GenS14{F0: -9223372036854775808}
                                                                _ = v141
                                                            }
                                                        }
                                                    case -9223372036854775808:
                                                        v142 := GenS18{F0: 227, F1: 1, F2: 527}
                                                        _ = v142
                                                    }
                                                    _ = len(v3)
                                                } else {
                                                    var v143 int32 = 0
                                                    _ = v143
                                                }
                                                _ = v97
                                            }
                                        }
                                        _ = v92
                                    }
                                    if (!true) {
                                        switch v1 {
                                        case "a":
                                            if (false && (!true)) {
                                                v144 := GenS19{F0: 42}
                                                v1 = ((("hello" + (v1 + ((((((("\n\t" + v1) + "a") + v1) + ("a" + v1)) + v1) + v1) + (v1 + ((v1 + v1) + v1))))) + "a") + (((v1 + v1) + v1) + ("\n\t" + "\n\t")))
                                                var v145 uint16 = (v11 + (v11 - 255))
                                                _ = v144
                                                _ = v145
                                            } else {
                                                var v146 float64 = 0.1
                                                v147 := GenS19{F0: -1}
                                                v15 = (v15 >> ((1) & 63))
                                                _ = v146
                                                _ = v147
                                            }
                                        case "\n\t":
                                            var v148 string = v1
                                            _ = v148
                                        default:
                                            if (v1 == "hello") {
                                                var v149 float32 = v17
                                                v15 += v15
                                                v150 := GenS2{F0: 32767}
                                                _ = v149
                                                _ = v150
                                            }
                                            v2 = v2
                                            v151 := GenS5{F0: -9223372036854775808, F1: 15}
                                            _ = v151
                                        }
                                    } else {
                                        fmt.Println(v4, v17, v15)
                                    }
                                    v2 = 1.5
                                    v15 -= 4294967295
                                }
                                var v152 string = ("a" + "foo bar")
                                v153 := GenS11{F0: 0, F1: (^(v15))}
                                _ = v152
                                _ = v153
                            } else {
                                switch fi16 {
                                case 268:
                                    if true {
                                        v12 -= (v12 & 1000)
                                        switch (fi14 >> ((4294967295) & 63)) {
                                        case 32768:
                                            v154 := GenS18{F0: 117, F1: uint(v11), F2: -190}
                                            _ = v154
                                        default:
                                            var v155 int32 = 1
                                            v156 := GenS20{F0: 2}
                                            _ = v155
                                            _ = v156
                                        }
                                    } else {
                                        if true {
                                            if ((v1 + (v1 + v1)) <= (((v1 + "\n\t") + v1) + "foo bar")) {
                                                var v157 float32 = 0.1
                                                v158 := GenS21{F0: (fi16 | (fi16 * fi16)), F1: 1, F2: 32768}
                                                _ = v157
                                                _ = v158
                                            }
                                            var v159 int = fi14
                                            _ = v159
                                        }
                                    }
                                case 516:
                                    if ((v33 % (((v33 + v33)) | 1)) != v33) {
                                        fmt.Println(v33)
                                        if true {
                                            var v160 float32 = 1e10
                                            var v161 string = v1
                                            fmt.Println(v1)
                                            _ = v160
                                            _ = v161
                                        }
                                    }
                                }
                            }
                            _ = v33
                        }
                    }
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
                var v162 uint16 = v11
                _ = v162
                _ = v17
            } else {
                v10 = append(v10, 0)
            }
            v11 -= v4
            v163 := GenS14{F0: -254}
            _ = fi16
            _ = v163
        }
        _ = fi14
        _ = v15
    }
    v164 := v3[-1]
    v10 = append(v10, 1)
    for fi165 := 0; fi165 < 12; fi165++ {
        v12 = 32767
        var v166 int32 = 816
        switch (fi165 & fi165) {
        case 370:
            for fi167 := 0; fi167 < 20; fi167++ {
                v12 = 1
                if false {
                    switch fi165 {
                    case -949:
                        var v168 uint16 = v11
                        _ = v168
                    case 916:
                        if (float64(v2) >= 1.5) {
                            switch ("\n\t" + v164) {
                            case "":
                                v4 = (v4 ^ v4)
                            case "\n\t":
                                if true {
                                    delete(v3, int(v166))
                                    if ((v12 + v12) <= v12) {
                                        var v169 uint16 = v11
                                        var v170 int32 = (v166 + 2147483647)
                                        _ = v169
                                        _ = v170
                                    } else {
                                        v11 = v11
                                    }
                                } else {
                                    var v171 float64 = float64(0.0)
                                    v166 -= 915
                                    _ = v171
                                }
                            default:
                                fmt.Println(v4, v164, v2)
                            }
                        } else {
                            fmt.Println(v166, v2, v166)
                            v1 = v1
                        }
                    case 32768:
                        var v172 int32 = -2147483648
                        _ = v172
                    }
                    var v173 int = fi167
                    v174 := GenS7{F0: 86}
                    _ = v173
                    _ = v174
                } else {
                    var v175 float32 = 1.5
                    _ = v175
                }
                fmt.Println(v11)
                _ = len(v10)
                var v176 float64 = -2.25
                _ = fi167
                _ = v176
            }
        case 127:
            if ((v11 - v4) <= (^(v4))) {
                if ((255 <= 65535) && (((67 != 127) || (858 != 0)) || false)) {
                    v12 = (-(v12))
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    if false {
                        var v177 string = (("hello" + v164) + v164)
                        var v178 uint32 = 2
                        _ = v177
                        _ = v178
                    } else {
                        switch (fi165 * (fi165 % ((int(v4)) | 1))) {
                        case -9223372036854775808:
                            v166 += v166
                        }
                    }
                }
                v3[fi165] = "\n\t"
                if false {
                    var v179 uint64 = 0
                    for fi180 := 0; fi180 < 19; fi180++ {
                        var v181 string = v1
                        var v182 int8 = 100
                        _ = fi180
                        _ = v181
                        _ = v182
                    }
                    var v183 int32 = v166
                    _ = v179
                    _ = v183
                }
            }
        case 42:
            var v184 uint64 = 4294967295
            _ = v184
        }
        _ = fi165
        _ = v166
    }
    if true {
        v2 -= ((v2 + v2) + 0.1)
    } else {
        switch -995 {
        case -980:
            var v185 float32 = v2
            _ = v185
        case 4294967295:
            for fi186 := 0; fi186 < 16; fi186++ {
                switch int(v4) {
                case 32768:
                    fmt.Println(v164, v2)
                default:
                    fmt.Println(v164, v4)
                }
                var v187 int64 = 0
                fmt.Println(v187, fi186, v11)
                _ = fi186
                _ = v187
            }
        default:
            fmt.Println(v2, v1, v12)
        }
    }
    for fi188 := 0; fi188 < 7; fi188++ {
        var v189 uint = 601
        _ = fi188
        _ = v189
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v164:", v164)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v164
    _ = v2
    _ = v3
    _ = v4
}
