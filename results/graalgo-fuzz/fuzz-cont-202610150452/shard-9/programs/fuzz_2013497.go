package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS11 struct { F0 uint8; F1 uint8; F2 int16 }
func genFn0(p0 int8, p1 string) int8 {
    var v6 uint64 = 1
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: v6, F1: -956, F2: p1}
        v6 += (v6 | 4294967295)
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := GenS1{F0: v6, F1: int32(p0), F2: p1}
            p1 = ""
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    if true {
        v6 += 4294967295
        switch 298 {
        case 0:
            var v11 string = p1
            _ = v11
        case 127:
            v12 := []int{42, -1}
            _ = v12
        }
        v13 := []int{128, 9223372036854775807}
        _ = v13
    } else {
        switch "\n\t" {
        case "a":
            var v14 bool = (-2.25 > ((float32(0.0) + 0.1) + -2.25))
            _ = v14
        case "\n\t":
            if true {
                var v15 uint16 = 403
                _ = v15
            }
        }
        if true {
            var v16 int16 = -32768
            _ = v16
        } else {
            if true {
                v17 := []int{127, 4294967295}
                if true {
                    v17 = append(v17, -218)
                    if (42 <= 1) {
                        v18 := GenS1{F0: v6, F1: 482, F2: "a"}
                        if len(v17) > 0 {
                            v17[0] = 823
                        }
                        _ = v18
                    }
                    v19 := GenS2{F0: 42, F1: -661}
                    _ = v19
                }
                if (-9223372036854775808 > 480) {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        if (int32(fi20) == 42) {
                            for fi21 := 0; fi21 < 2; fi21++ {
                                p1 = ((p1 + p1) + ((p1 + p1) + ("\n\t" + p1)))
                                p0 = p0
                                _ = fi21
                            }
                            v17 = append(v17, fi20)
                            p1 = "\n\t"
                        } else {
                            var v22 uint32 = 725
                            for fi23 := 0; fi23 < 13; fi23++ {
                                switch p1 {
                                case "foo bar":
                                    v24 := GenS1{F0: (^(v6)), F1: 2147483647, F2: p1}
                                    _ = v24
                                case "hello":
                                    v6 += (v6 * v6)
                                case "":
                                    v17 = append(v17, int(v22))
                                }
                                var v25 int32 = -2147483648
                                var v26 string = ""
                                _ = fi23
                                _ = v25
                                _ = v26
                            }
                            _ = v22
                        }
                        _ = fi20
                    }
                    var v27 int8 = -30
                    _ = v27
                } else {
                    v17 = append(v17, 1)
                    for fi28 := 0; fi28 < 14; fi28++ {
                        v6 += (v6 * (v6 & v6))
                        _ = fi28
                    }
                    v6 += 1
                }
                _ = v17
            }
            v29 := []int{9223372036854775807, 32767, 127}
            switch "foo bar" {
            case "":
                v30 := GenS1{F0: v6, F1: 0, F2: p1}
                _ = v30
            case "foo bar":
                v6 -= (^(v6))
            case "\n\t":
                p1 = (("foo bar" + (p1 + ((p1 + p1) + "foo bar"))) + ("a" + "foo bar"))
            }
            _ = v29
        }
        v6 = v6
    }
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    var v1 bool = true
    var v2 int16 = 32767
    var v3 uint32 = 2147483647
    var v4 int32 = 1
    var v5 uint16 = 591
    var v31 uint32 = 65535
    for fi32 := 0; fi32 < 19; fi32++ {
        if (!(18446744073709551615 != 65535)) {
            v31 += (v3 * 65535)
        } else {
            for fi33 := 0; fi33 < 18; fi33++ {
                if false {
                    v34 := []int{4294967295}
                    _ = v34
                }
                if (!(!false)) {
                    v4 = -122
                    if false {
                        if (0 != 92) {
                            var v35 float64 = float64(1e10)
                            _ = v35
                        } else {
                            fmt.Println(v4)
                        }
                    }
                    if (100 <= -101) {
                        var v36 uint16 = 65535
                        var v37 bool = v1
                        _ = v36
                        _ = v37
                    } else {
                        var v38 string = "a"
                        var v39 float64 = 0.0
                        _ = v38
                        _ = v39
                    }
                }
                v40 := []int{(fi32 / ((-513) | 1)), fi32}
                _ = fi33
                _ = v40
            }
            for fi41 := 0; fi41 < 5; fi41++ {
                v42 := []int{(fi32 * fi41), (fi41 * 32767)}
                v42 = append(v42, (fi41 & (fi41 * fi41)))
                _ = fi41
                _ = v42
            }
        }
        v43 := []int{-577, -300, fi32, fi32}
        if len(v43) > 0 {
            v43[0] = (fi32 - (-(fi32)))
        }
        _ = fi32
        _ = v43
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        fmt.Println(v5)
        v45 := GenS3{F0: 278, F1: (fi44 | (fi44 * 1))}
        var v46 bool = (3.14159 >= 3.14159)
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println(v31, v3, v5)
    var v47 float32 = 3.14159
    for fi48 := 0; fi48 < 12; fi48++ {
        for fi49 := 0; fi49 < 15; fi49++ {
            if false {
                if false {
                    if false {
                        v50 := []int{(fi49 >> ((831) & 63))}
                        v5 -= (^(v5))
                        switch fi48 {
                        case 32768:
                            if len(v50) > 0 {
                                v50[0] = fi48
                            }
                        case 0:
                            if true {
                                v51 := GenS4{F0: 127, F1: 0}
                                _ = v51
                            } else {
                                var v52 int16 = v2
                                v1 = ((v2 ^ (v2 + (v2 ^ (v52 << ((1) & 63))))) == v52)
                                _ = v52
                            }
                        case -9223372036854775808:
                            var v53 float32 = ((v47 * 0.0) / ((v47 + v47) - (float32((1e10 / (-2.25 * (float64(float32(1.5)) * -2.25)))) - v47)))
                            _ = v53
                        default:
                            _ = len(v50)
                        }
                        _ = v50
                    } else {
                        var v54 int = (fi48 / ((-622) | 1))
                        _ = v54
                    }
                    v31 += (v3 >> ((306) & 63))
                } else {
                    switch -47 {
                    case 42:
                        fmt.Println(v4, v3, v2)
                    case 1:
                        var v55 uint8 = 127
                        _ = v55
                    }
                    v56 := GenS5{F0: 187}
                    _ = v56
                }
            } else {
                var v57 uint64 = 2
                var v58 bool = false
                _ = v57
                _ = v58
            }
            if (v3 != 447) {
                if true {
                    if true {
                        var v59 int8 = int8(fi48)
                        _ = v59
                    }
                    v60 := []int{fi48, fi49, int(v2)}
                    if len(v60) >= 2 {
                        v60 = v60[1:len(v60)]
                    }
                    _ = v60
                }
                fmt.Println(fi49)
                switch (fi49 ^ fi49) {
                case -982:
                    var v61 uint16 = v5
                    _ = v61
                case 2147483647:
                    v62 := []int{}
                    _ = v62
                default:
                    var v63 int16 = 1000
                    _ = v63
                }
            }
            _ = fi49
        }
        if v1 {
            if v1 {
                v64 := GenS6{F0: 4294967295, F1: (v2 % ((v2) | 1))}
                switch "foo bar" {
                case "foo bar":
                    for fi65 := 0; fi65 < 9; fi65++ {
                        v1 = (42 == 191)
                        v5 += (v5 * (v5 << ((870) & 63)))
                        var v66 bool = v1
                        _ = fi65
                        _ = v66
                    }
                case "a":
                    fmt.Println(v4, v1, fi48)
                default:
                    for fi67 := 0; fi67 < 15; fi67++ {
                        switch 32768 {
                        case 457:
                            if (1 <= 184) {
                                v68 := []int{467, (^(fi67))}
                                v47 += float32(0.1)
                                var v69 int64 = 1
                                fmt.Println(v4)
                                v3 -= 65535
                                _ = v68
                                _ = v69
                            }
                        case 1:
                            var v70 float32 = v47
                            _ = v70
                        case 2147483647:
                            v71 := []int{(^(fi67)), (fi48 / ((fi48) | 1))}
                            var v72 int32 = 0
                            _ = v71
                            _ = v72
                        }
                        _ = fi67
                    }
                }
                for fi73 := 0; fi73 < 20; fi73++ {
                    v1 = (!(("foo bar" + ("a" + ("foo bar" + "\n\t"))) <= "foo bar"))
                    _ = fi73
                }
                _ = v64
            }
            var v74 int8 = -69
            _ = v74
        }
        _ = fi48
    }
    v75 := GenS7{F0: v2}
    v4 -= v4
    for fi76 := 0; fi76 < 17; fi76++ {
        v77 := []int{}
        if (-128 == -67) {
            var v78 int32 = (^(v4))
            v79 := GenS5{F0: 365}
            _ = len(v77)
            _ = v78
            _ = v79
        } else {
            v80 := GenS8{F0: 2}
            for fi81 := 0; fi81 < 13; fi81++ {
                switch "hello" {
                case "hello":
                    v1 = (18446744073709551615 > 0)
                case "":
                    v4 = (v4 | 2147483647)
                case "foo bar":
                    switch "foo bar" {
                    case "\n\t":
                        if (!(true || false)) {
                            var v82 int = int(v4)
                            _ = v82
                        }
                    case "a":
                        if v1 {
                            v3 += (v31 >> ((639) & 63))
                            if v1 {
                                v77 = append(v77, (fi81 >> ((255) & 63)))
                            } else {
                                if v1 {
                                    v5 = 65535
                                    var v83 int8 = -118
                                    switch "" {
                                    case "a":
                                        v77 = append(v77, fi76)
                                    }
                                    _ = v83
                                } else {
                                    var v84 string = "\n\t"
                                    fmt.Println(fi81, v5)
                                    if (v3 > 1) {
                                        var v85 uint = 18446744073709551615
                                        v5 = (v5 << (((^(v85))) & 63))
                                        v86 := GenS9{F0: v2, F1: v84}
                                        var v87 int8 = 125
                                        v4 = (v4 >> ((584) & 63))
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                    }
                                    _ = v84
                                }
                            }
                            v31 = (v3 << ((65535) & 63))
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(v1, v2)
                            case "":
                                var v88 int32 = (v4 * 1)
                                _ = v88
                            default:
                                fmt.Println(v4, v3)
                            }
                            v5 -= v5
                            var v89 int16 = (v2 - (v2 >> ((242) & 63)))
                            _ = v89
                        }
                    case "foo bar":
                        v77 = append(v77, -9223372036854775808)
                    }
                }
                if v1 {
                    v90 := GenS1{F0: 1, F1: -2147483648, F2: ((("foo bar" + "foo bar") + "a") + "a")}
                    var v91 uint32 = (v31 - (v3 & v3))
                    _ = v90
                    _ = v91
                } else {
                    if len(v77) > 0 {
                        v77[0] = 245
                    }
                    v92 := GenS10{F0: 254, F1: (v2 ^ (v2 ^ v2)), F2: uint8(v31)}
                    _ = v92
                }
                v47 += v47
                _ = fi81
            }
            _ = v80
        }
        v3 += 821
        _ = fi76
        _ = v77
    }
    for fi93 := 0; fi93 < 12; fi93++ {
        v94 := GenS1{F0: 1, F1: v4, F2: "a"}
        v95 := []int{-2147483648, -8, fi93}
        switch "a" {
        case "\n\t":
            for fi96 := 0; fi96 < 1; fi96++ {
                if (!v1) {
                    v3 = 97
                    fmt.Println(v5, fi96)
                    if false {
                        if false {
                            v2 += (v2 << ((2) & 63))
                            v4 = 366
                            v47 = (v47 * float32(-2.25))
                        } else {
                            var v97 int16 = v2
                            v98 := GenS11{F0: 0, F1: 1, F2: v97}
                            _ = v97
                            _ = v98
                        }
                    }
                } else {
                    var v99 float64 = -2.25
                    var v100 int8 = 1
                    v99 += (0.1 + ((v99 + (float64(1.5) + v99)) * ((((v99 / (((float64((v47 / (v47 + (0.1 + (v47 * (((1e10 / v47) - v47) + v47)))))) - v99) / float64(0.1)) * (float64((v47 - (v47 + 0.0))) - (((v99 / 0.1) - (1.5 * 0.0)) + v99)))) / v99) - float64(((1e10 * v47) - 1.5))) * (float64(float32(((-2.25 + (1e10 + v99)) / v99))) * 1.5))))
                    _ = v100
                    _ = v99
                }
                var v101 uint = 18446744073709551615
                _ = fi96
                _ = v101
            }
        default:
            for fi102 := 0; fi102 < 16; fi102++ {
                fmt.Println(fi102, v1, v2)
                if len(v95) >= 2 {
                    v95 = v95[1:len(v95)]
                }
                _ = fi102
            }
        }
        _ = fi93
        _ = v94
        _ = v95
    }
    for fi103 := 0; fi103 < 11; fi103++ {
        fmt.Println(v5)
        _ = fi103
    }
    if (genFn0(-107, ("hello" + (("a" + ("a" + "\n\t")) + "hello"))) >= -128) {
        var v104 float64 = 3.14159
        _ = v104
    }
    v2 += v2
    var v105 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v105
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v47
    _ = v5
    _ = v75
}
