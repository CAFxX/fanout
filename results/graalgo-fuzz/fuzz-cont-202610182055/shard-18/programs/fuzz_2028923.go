package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int; F2 int }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 uint; F1 int16 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint8; F1 uint64; F2 string }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint32; F1 int }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int8; F1 uint32 }
type GenS17 struct { F0 int16; F1 string; F2 int64 }
type GenS18 struct { F0 uint8 }
func genFn0(p0 uint16, p1 uint8, p2 uint32) int16 {
    v5 := []int{127, 42}
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint8 = (^(p1))
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int = (fi8 & 42)
            _ = fi8
            _ = v9
        }
        p1 = 212
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 699
}
func genFn1(p0 uint8, p1 int64, p2 int) int32 {
    v10 := []int{-2147483648, p2}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -26
}

func main() {
    v1 := []string{("foo bar" + "\n\t")}
    var v2 uint32 = 65535
    var v3 float64 = (0.1 - float64(0.1))
    var v4 float64 = (-2.25 * (float64(0.1) - v3))
    fmt.Println(v4, v4, v2)
    var v11 string = ("a" + "foo bar")
    v2 = (v2 * 2147483647)
    var v12 int64 = 9223372036854775807
    v1 = append(v1, "\n\t")
    for fi13 := 0; fi13 < 10; fi13++ {
        switch 4294967295 {
        case 128:
            for fi14 := 0; fi14 < 10; fi14++ {
                if false {
                    if len(v1) > 0 {
                        v1[0] = v11
                    }
                    v15 := GenS1{F0: (v2 / ((0) | 1)), F1: 1, F2: -9223372036854775808}
                    _ = v15
                } else {
                    switch fi13 {
                    case -925:
                        switch (fi13 ^ fi14) {
                        case 42:
                            var v16 string = "a"
                            _ = v16
                        case -1:
                            v4 += ((((v4 / ((((3.14159 - ((float64(0.1) + float64((((1.5 - 3.14159) - (1.5 - 1e10)) + float32(v3)))) - v4)) - (((float64(3.14159) * v4) + (float64((-2.25 * ((1.5 + 3.14159) * 1.5))) + v4)) * 0.1)) + -2.25) * (v3 + v4))) / 1.0) / float64(3.14159)) - (3.14159 - v3))
                        default:
                            if len(v1) > 0 {
                                v1[0] = v11
                            }
                        }
                    default:
                        v17 := GenS2{F0: v11}
                        _ = v17
                    }
                    v2 = (^(v2))
                    var v18 float64 = 0.0
                    _ = v18
                }
                if false {
                    fmt.Println(v2, fi13, v2)
                    var v19 float32 = float32(0.1)
                    _ = v19
                } else {
                    switch "" {
                    case "foo bar":
                        v3 -= (0.0 / 1.0)
                        fmt.Println(v2)
                    case "\n\t":
                        if true {
                            var v20 int64 = (^(v12))
                            v3 -= v3
                            _ = v20
                        } else {
                            switch (fi13 * (fi13 | 128)) {
                            case -2147483648:
                                fmt.Println(fi14, fi14, v4)
                            default:
                                v2 = 363
                            }
                        }
                        fmt.Println(v3)
                        if false {
                            if (!(v11 < v11)) {
                                var v21 int64 = (^(v12))
                                _ = v21
                            } else {
                                v3 = v4
                                var v22 string = v11
                                _ = v22
                            }
                            fmt.Println(v2, v4)
                        }
                    case "":
                        switch v11 {
                        case "":
                            if false {
                                var v23 uint8 = 2
                                var v24 bool = ((^(v23)) < (v23 | (v23 >> ((0) & 63))))
                                if len(v1) > 0 {
                                    v1[0] = ("\n\t" + v11)
                                }
                                v1 = append(v1, v11)
                                v25 := GenS3{F0: v23, F1: (-(fi13)), F2: fi13}
                                if v24 {
                                    switch v11 {
                                    case "\n\t":
                                        v24 = false
                                    default:
                                        var v26 uint = 65535
                                        _ = v26
                                    }
                                    v27 := GenS1{F0: v2, F1: 42, F2: fi14}
                                    _ = v27
                                } else {
                                    var v28 int8 = -128
                                    v29 := GenS4{F0: genFn1(0, (v12 / ((v12) | 1)), fi14), F1: int64(fi13)}
                                    _ = v28
                                    _ = v29
                                }
                                v2 -= (^(v2))
                                var v30 float32 = 0.0
                                _ = v23
                                _ = v24
                                _ = v25
                                _ = v30
                            }
                        case "foo bar":
                            var v31 int = (^(fi13))
                            v12 += (v12 << ((1) & 63))
                            if true {
                                var v32 int64 = v12
                                v1 = append(v1, (v11 + "\n\t"))
                                if ((v11 + v11) < (v11 + ("a" + ("a" + v11)))) {
                                    if false {
                                        fmt.Println(v4)
                                        v12 = 42
                                        var v33 int8 = -1
                                        _ = v33
                                    } else {
                                        v1 = append(v1, "foo bar")
                                        _ = len(v1)
                                    }
                                    if true {
                                        v12 += (^(v32))
                                        var v34 uint8 = 2
                                        _ = v34
                                    }
                                }
                                _ = v32
                            } else {
                                var v35 uint = 4294967295
                                _ = v35
                            }
                            _ = v31
                        }
                    }
                    if ((!(v2 > 2147483647)) || false) {
                        if (246 != 2) {
                            if false {
                                v12 = (v12 / ((v12) | 1))
                                v2 += (^(v2))
                            }
                            if false {
                                if (true || true) {
                                    if false {
                                        v12 -= v12
                                        switch v11 {
                                        case "hello":
                                            v36 := GenS3{F0: 2, F1: (fi14 - (fi14 - fi14)), F2: -570}
                                            _ = v36
                                        case "":
                                            if (!true) {
                                                var v37 uint32 = v2
                                                _ = v37
                                            } else {
                                                v12 = (-(v12))
                                            }
                                        case "\n\t":
                                            switch (^(fi13)) {
                                            case -573:
                                                v12 = v12
                                            case -611:
                                                if false {
                                                    v12 -= v12
                                                    var v38 uint8 = uint8(v2)
                                                    var v39 float64 = (float64(-2.25) * v3)
                                                    switch 1 {
                                                    case 379:
                                                        fmt.Println(v11)
                                                    case 639:
                                                        if len(v1) > 0 {
                                                            v1[0] = v11
                                                        }
                                                    }
                                                    v12 += (v12 >> ((959) & 63))
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    fmt.Println(v3, v4)
                                                    switch 9223372036854775807 {
                                                    case 747:
                                                        v1 = append(v1, v11)
                                                    }
                                                    _ = len(v1)
                                                }
                                            }
                                        }
                                        var v40 uint32 = (^(v2))
                                        _ = v40
                                    } else {
                                        switch v11 {
                                        case "hello":
                                            switch (fi13 / ((157) | 1)) {
                                            case 193:
                                                if false {
                                                    v41 := GenS5{F0: 2}
                                                    if false {
                                                        if false {
                                                            if false {
                                                                v11 = v11
                                                                v2 = v2
                                                                fmt.Println(v4)
                                                            }
                                                            switch -9223372036854775808 {
                                                            case 524:
                                                                var v42 uint32 = v2
                                                                _ = v42
                                                            case -851:
                                                                var v43 int64 = (v12 >> ((4294967295) & 63))
                                                                _ = v43
                                                            case -156:
                                                                if (255 <= 1) {
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        fmt.Println(v3, v4)
                                                                    case "a":
                                                                        fmt.Println(v11)
                                                                    case "foo bar":
                                                                        fmt.Println(fi14)
                                                                    }
                                                                    v1 = append(v1, "foo bar")
                                                                }
                                                            default:
                                                                v1 = append(v1, "\n\t")
                                                            }
                                                        }
                                                        if (true && (4294967295 != 0)) {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v44 uint = 65535
                                                            var v45 int64 = v12
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                    }
                                                    v46 := GenS1{F0: 2147483647, F1: -32768, F2: -613}
                                                    var v47 int32 = -1
                                                    _ = v41
                                                    _ = v46
                                                    _ = v47
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = v11
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v48 uint8 = 2
                                                    _ = v48
                                                }
                                            case -2147483648:
                                                var v49 uint8 = 127
                                                _ = v49
                                            }
                                        case "":
                                            if (v2 == (^(v2))) {
                                                if len(v1) > 0 {
                                                    v1[0] = "a"
                                                }
                                            } else {
                                                v3 += v4
                                            }
                                        case "foo bar":
                                            v50 := GenS6{F0: (fi14 ^ (fi14 / ((fi14) | 1))), F1: uint16(v2)}
                                            var v51 float64 = ((v3 * v4) / (v4 + (v3 - 1e10)))
                                            fmt.Println(v11, v2, v12)
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v50
                                            _ = v51
                                        }
                                        var v52 int = fi14
                                        v2 = v2
                                        _ = v52
                                    }
                                    switch 4294967295 {
                                    case 2147483647:
                                        _ = len(v1)
                                    case -35:
                                        if (int8(v2) < 1) {
                                            if len(v1) > 0 {
                                                v1[0] = (v11 + ((v11 + v11) + ""))
                                            }
                                        }
                                    case -768:
                                        v12 = (v12 & (v12 | 9223372036854775807))
                                    default:
                                        v1 = append(v1, ("foo bar" + v11))
                                    }
                                    var v53 uint16 = 2
                                    _ = v53
                                } else {
                                    v3 = 1.5
                                }
                                var v54 int8 = 22
                                fmt.Println(v11)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch v11 {
                                case "\n\t":
                                    v12 = -97
                                case "":
                                    v3 = (v4 + (v4 + v4))
                                case "hello":
                                    v55 := GenS7{F0: 334, F1: -507}
                                    _ = v55
                                default:
                                    if (true || ((1 < 18446744073709551615) && (!true))) {
                                        v56 := GenS3{F0: 14, F1: 128, F2: -2147483648}
                                        switch "hello" {
                                        case "hello":
                                            v1 = append(v1, ("a" + "a"))
                                            v4 += ((0.0 - v3) + 0.0)
                                            fmt.Println(fi13)
                                        case "foo bar":
                                            switch (v11 + (v11 + (((("foo bar" + "a") + "\n\t") + v11) + v11))) {
                                            case "foo bar":
                                                fmt.Println(v2)
                                            case "hello":
                                                var v57 int64 = v12
                                                _ = v57
                                            case "a":
                                                v11 = (v11 + (v11 + v11))
                                            }
                                        }
                                        _ = v56
                                    }
                                    var v58 bool = true
                                    _ = v58
                                }
                                _ = v54
                            }
                            v59 := GenS2{F0: ""}
                            if ((1 < -2147483648) || ((128 != 1) && (fi13 < fi13))) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v60 bool = false
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                                v11 = v11
                                _ = v60
                            } else {
                                if (836 == 65535) {
                                    v61 := GenS4{F0: 42, F1: -1}
                                    v2 -= (v2 * v2)
                                    v62 := GenS7{F0: 18446744073709551615, F1: 1000}
                                    _ = v61
                                    _ = v62
                                } else {
                                    v1 = append(v1, "")
                                }
                            }
                            var v63 int = 128
                            _ = v59
                            _ = v63
                        }
                        if true {
                            v2 = v2
                            switch 42 {
                            case -436:
                                v64 := GenS8{F0: 689, F1: -47}
                                _ = len(v1)
                                v65 := GenS9{F0: v2}
                                if true {
                                    v66 := GenS3{F0: 128, F1: (fi14 % ((fi13) | 1)), F2: fi14}
                                    if (1000 != 473) {
                                        var v67 uint = 433
                                        v68 := GenS4{F0: 2147483647, F1: v12}
                                        v67 += (v67 | v67)
                                        _ = v67
                                        _ = v68
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = v11
                                        }
                                    }
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v3 -= (float64(1e10) / (v4 + v4))
                                        fmt.Println(v2, v3, fi14)
                                        if false {
                                            fmt.Println(v4, v4, v2)
                                            var v69 float64 = 0.0
                                            _ = v69
                                        } else {
                                            if true {
                                                if false {
                                                    v3 += (v4 / float64(1.5))
                                                    switch (fi14 | fi13) {
                                                    case 32767:
                                                        var v70 float32 = 0.1
                                                        _ = v70
                                                    default:
                                                        fmt.Println(v2, v11)
                                                        switch ("foo bar" + ((v11 + (((v11 + ("hello" + v11)) + "hello") + v11)) + v11)) {
                                                        case "foo bar":
                                                            v4 -= 3.14159
                                                        case "\n\t":
                                                            if false {
                                                                v4 = v3
                                                                fmt.Println(v11, v11)
                                                                v71 := GenS8{F0: 2, F1: -128}
                                                                _ = v71
                                                            }
                                                        case "hello":
                                                            var v72 uint8 = 127
                                                            _ = v72
                                                        }
                                                        v73 := GenS8{F0: uint16(v2), F1: 1}
                                                        _ = v73
                                                    }
                                                }
                                                v74 := GenS10{F0: 0, F1: 42, F2: (v11 + v11)}
                                                v3 -= v3
                                                _ = v74
                                            } else {
                                                var v75 int = (fi13 >> ((1) & 63))
                                                switch (fi13 / ((fi13) | 1)) {
                                                case 32767:
                                                    var v76 string = ("hello" + v11)
                                                    _ = v76
                                                case 32768:
                                                    v77 := GenS11{F0: 2147483647}
                                                    _ = v77
                                                case 69:
                                                    v78 := GenS12{F0: (!(v4 >= (v3 + (((((v3 - v3) + float64(((1e10 + 0.1) - -2.25))) + (v3 - v3)) + v4) * v3))))}
                                                    _ = v78
                                                }
                                                _ = v75
                                            }
                                            var v79 uint8 = 0
                                            var v80 int8 = 82
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v81 uint = 65535
                                            _ = v79
                                            _ = v80
                                            _ = v81
                                        }
                                        var v82 int8 = -65
                                        var v83 int8 = v82
                                        _ = v82
                                        _ = v83
                                    }
                                    _ = v66
                                } else {
                                    var v84 string = ((v11 + v11) + v11)
                                    _ = v84
                                }
                                _ = v64
                                _ = v65
                            case -2147483648:
                                var v85 int16 = 1
                                _ = v85
                            case 127:
                                if (!false) {
                                    _ = len(v1)
                                    var v86 uint = 493
                                    v3 -= (v3 - v3)
                                    _ = v86
                                } else {
                                    var v87 int16 = 0
                                    v88 := GenS2{F0: v11}
                                    _ = v87
                                    _ = v88
                                }
                            }
                        }
                    } else {
                        v12 = (^(v12))
                        var v89 int = fi14
                        var v90 int32 = 1
                        _ = v89
                        _ = v90
                    }
                    var v91 uint = 1
                    _ = v91
                }
                v3 -= (v3 - v3)
                switch fi13 {
                case 0:
                    if (!false) {
                        if false {
                            switch fi14 {
                            case -2147483648:
                                var v92 int32 = -2147483648
                                _ = v92
                            case -2:
                                v93 := GenS13{F0: 0}
                                _ = v93
                            case -1:
                                v11 = ("\n\t" + v11)
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    } else {
                        v3 += 0.0
                    }
                case -1:
                    if ((true && false) && true) {
                        fmt.Println(v11, fi13, v3)
                        if (!(1 != 24)) {
                            var v94 uint16 = 42
                            v4 -= v3
                            _ = len(v1)
                            v2 = v2
                            fmt.Println(fi13, v11)
                            v2 -= v2
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if true {
                                var v95 uint16 = (^(v94))
                                _ = v95
                            } else {
                                v2 = 4294967295
                                switch "hello" {
                                case "foo bar":
                                    if true {
                                        var v96 int = 128
                                        _ = v96
                                    } else {
                                        v4 = v3
                                        if false {
                                            if ((true && (495 <= 4294967295)) || true) {
                                                if true {
                                                    var v97 int16 = 32767
                                                    if (42 <= 104) {
                                                        v98 := GenS9{F0: v2}
                                                        _ = v98
                                                    }
                                                    _ = v97
                                                } else {
                                                    v3 = (3.14159 * (((0.1 / v3) + float64(0.1)) / (1.5 + v4)))
                                                    fmt.Println(v12, v11)
                                                    v12 = (v12 + v12)
                                                    var v99 bool = (0.0 >= (1.5 * ((1e10 / 3.14159) + float32(((0.0 / v4) + -2.25)))))
                                                    _ = v99
                                                }
                                                fmt.Println(v3, v12)
                                            }
                                        }
                                    }
                                case "hello":
                                    v94 = 65535
                                default:
                                    var v100 float32 = 1.5
                                    _ = v100
                                }
                            }
                            v101 := GenS12{F0: (true && (127 < 159))}
                            _ = len(v1)
                            v3 = (0.0 / ((v4 * -2.25) * v4))
                            var v102 int32 = 139
                            v94 += (v94 & v94)
                            _ = v101
                            _ = v102
                            _ = v94
                        } else {
                            switch "hello" {
                            case "a":
                                if ((!true) && (((((v2 >> ((730) & 63)) <= (v2 & (^(v2)))) && true) || (!true)) && true)) {
                                    switch (fi14 - (fi14 | (fi13 / (((fi13 + 0)) | 1)))) {
                                    case -625:
                                        v4 -= 0.0
                                    case -1:
                                        fmt.Println(fi14, fi13)
                                    case 127:
                                        fmt.Println(v2, fi13)
                                    }
                                    var v103 int16 = -1
                                    if len(v1) > 0 {
                                        v1[0] = (((v11 + "hello") + (v11 + (v11 + "a"))) + v11)
                                    }
                                    _ = v103
                                } else {
                                    if true {
                                        v104 := GenS14{F0: (v2 - (v2 >> ((1) & 63))), F1: (fi13 ^ fi13)}
                                        v11 = (v11 + v11)
                                        switch ((((("hello" + (v11 + (("" + ("a" + "foo bar")) + (v11 + "\n\t")))) + ((("" + ("a" + "")) + v11) + "")) + (v11 + "\n\t")) + "a") + v11) {
                                        case "a":
                                            if len(v1) > 0 {
                                                v1[0] = (v11 + "\n\t")
                                            }
                                        default:
                                            if true {
                                                if len(v1) > 0 {
                                                    v1[0] = "foo bar"
                                                }
                                                switch "" {
                                                case "a":
                                                    switch v11 {
                                                    case "\n\t":
                                                        v2 = 65535
                                                    case "a":
                                                        v3 = v4
                                                        var v105 uint8 = 2
                                                        _ = v105
                                                    default:
                                                        if (!true) {
                                                            v1 = append(v1, "\n\t")
                                                        }
                                                    }
                                                default:
                                                    var v106 uint8 = 1
                                                    _ = v106
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = ("\n\t" + v11)
                                                }
                                            }
                                        }
                                        _ = v104
                                    } else {
                                        switch 9223372036854775807 {
                                        case 128:
                                            v107 := GenS5{F0: 173}
                                            _ = v107
                                        case -9223372036854775808:
                                            var v108 string = v11
                                            _ = v108
                                        }
                                        v109 := GenS15{F0: 1}
                                        var v110 float32 = -2.25
                                        _ = v109
                                        _ = v110
                                    }
                                    _ = len(v1)
                                    switch fi14 {
                                    case 31:
                                        v111 := GenS16{F0: 92, F1: (v2 ^ v2)}
                                        _ = v111
                                    case 42:
                                        v11 = ("\n\t" + (v11 + v11))
                                    }
                                    v112 := GenS14{F0: v2, F1: (fi14 ^ (fi13 & (fi13 + fi14)))}
                                    _ = v112
                                }
                            case "hello":
                                var v113 uint8 = 26
                                _ = v113
                            default:
                                if false {
                                    if (170 >= 2) {
                                        v114 := GenS17{F0: -1, F1: v11, F2: v12}
                                        v2 -= v2
                                        fmt.Println(v11, v11, v4)
                                        _ = v114
                                    }
                                    v4 -= 1.5
                                    v11 = (v11 + ("\n\t" + "\n\t"))
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = v11
                                    }
                                    var v115 int8 = -39
                                    v116 := GenS8{F0: 32767, F1: v115}
                                    switch fi13 {
                                    case -9223372036854775808:
                                        v117 := GenS10{F0: 152, F1: 50, F2: "foo bar"}
                                        _ = v117
                                    case 380:
                                        v4 -= v4
                                    case -1:
                                        var v118 uint8 = 1
                                        _ = v118
                                    default:
                                        v4 += 1.5
                                    }
                                    _ = v115
                                    _ = v116
                                }
                            }
                            v12 -= -9223372036854775808
                        }
                    } else {
                        v12 = v12
                    }
                case 4294967295:
                    v119 := GenS2{F0: v11}
                    if ((v2 / ((42) | 1)) != v2) {
                        v4 -= v4
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v120 := GenS17{F0: int16(v2), F1: "a", F2: int64(v2)}
                        _ = v120
                    }
                    v2 = (^(v2))
                    _ = v119
                default:
                    switch 128 {
                    case -6:
                        v121 := GenS18{F0: 0}
                        _ = v121
                    default:
                        v11 = v11
                    }
                }
                _ = fi14
            }
        case -988:
            v12 = v12
        case 0:
            v2 = v2
        }
        _ = fi13
    }
    fmt.Println(v4, v11)
    v122 := GenS11{F0: 2147483647}
    var v123 uint16 = 255
    v124 := GenS1{F0: v2, F1: -375, F2: 979}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v122
    _ = v123
    _ = v124
    _ = v2
    _ = v3
    _ = v4
}
