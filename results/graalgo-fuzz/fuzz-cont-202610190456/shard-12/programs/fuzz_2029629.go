package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 uint8; F2 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS9 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16; F1 int16 }
type GenS12 struct { F0 int16; F1 string; F2 int8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint32; F1 uint16; F2 uint }
type GenS16 struct { F0 int32; F1 int64; F2 uint32 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 bool; F1 uint8; F2 int8 }
type GenS19 struct { F0 int; F1 uint64 }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 int }
type GenS22 struct { F0 uint; F1 int32 }
func genFn0() uint8 {
    v5 := GenS1{F0: (!false), F1: 2147483647}
    if true {
        v6 := GenS2{F0: 0}
        var v7 bool = false
        _ = v6
        _ = v7
    }
    _ = v5
    return 0
}

func main() {
    var v1 int32 = -2147483648
    var v2 int8 = 2
    var v3 string = "foo bar"
    var v4 string = (v3 + v3)
    if ((!true) || true) {
        if ((v1 * v1) >= (v1 & (v1 >> ((42) & 63)))) {
            v4 = v4
        } else {
            if (v3 >= v4) {
                v4 = (v3 + (v4 + "a"))
            } else {
                fmt.Println(v3)
                var v8 uint16 = 1
                _ = v8
            }
            v9 := GenS2{F0: -2147483648}
            v10 := GenS3{F0: 1}
            _ = v10
            _ = v9
        }
    } else {
        if (((^(v1)) == v1) && true) {
            v11 := []int{-926, 42, 9223372036854775807}
            _ = v11
        } else {
            var v12 int16 = 42
            var v13 uint32 = 1
            _ = v12
            _ = v13
        }
    }
    if (128 <= 0) {
        if false {
            if true {
                if true {
                    v1 -= -1
                } else {
                    v14 := []int{32767, -9223372036854775808, 0}
                    var v15 int32 = v1
                    _ = v14
                    _ = v15
                }
            } else {
                v16 := []int{-9223372036854775808, -9223372036854775808}
                switch 128 {
                case 32767:
                    if false {
                        v3 = v4
                        v3 = v4
                    } else {
                        v16 = append(v16, 963)
                    }
                case 4294967295:
                    var v17 int16 = 42
                    _ = v17
                case 729:
                    v1 -= (v1 >> ((1) & 63))
                }
                fmt.Println(v1, v4)
                _ = v16
            }
        } else {
            v2 = (v2 & (^(v2)))
            v18 := []int{9223372036854775807, 1}
            _ = v18
        }
        if false {
            for fi19 := 0; fi19 < 7; fi19++ {
                v2 -= (v2 >> ((0) & 63))
                _ = fi19
            }
            fmt.Println(v4)
            var v20 uint32 = 501
            _ = v20
        } else {
            var v21 int32 = (v1 + (v1 % ((1) | 1)))
            v22 := []int{2147483647, -2147483648, 32768}
            v4 = v3
            _ = v21
            _ = v22
        }
    }
    v23 := []int{1}
    v2 = (v2 / (((-(v2))) | 1))
    var v24 int16 = 42
    v24 -= 32767
    switch 127 {
    case 32768:
        v3 = (((("hello" + v4) + (v3 + v4)) + "hello") + (((v3 + (v4 + "a")) + "\n\t") + "\n\t"))
    case 42:
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
    case -9223372036854775808:
        v25 := GenS3{F0: 42}
        _ = v25
    }
    switch 4294967295 {
    case 259:
        var v26 uint = 42
        _ = v26
    }
    fmt.Println(v24, v2, v4)
    v27 := GenS2{F0: (v1 ^ (-(v1)))}
    switch v3 {
    case "":
        var v28 float64 = (-2.25 * -2.25)
        _ = v28
    case "\n\t":
        if (1.5 >= 1.5) {
            v4 = v4
            switch v4 {
            case "\n\t":
                fmt.Println(v2, v1)
            case "":
                v23 = append(v23, 193)
            default:
                var v29 float64 = -2.25
                _ = v29
            }
            v2 = (-(v2))
        } else {
            v30 := GenS4{F0: (^(v24)), F1: genFn0(), F2: false}
            _ = v30
        }
    case "a":
        switch -474 {
        case -470:
            for fi31 := 0; fi31 < 20; fi31++ {
                if true {
                    for fi32 := 0; fi32 < 16; fi32++ {
                        v2 = v2
                        v24 = 0
                        _ = fi32
                    }
                } else {
                    for fi33 := 0; fi33 < 7; fi33++ {
                        v34 := GenS3{F0: -443}
                        if false {
                            switch ("a" + "hello") {
                            case "a":
                                var v35 uint64 = 657
                                _ = v35
                            case "foo bar":
                                _ = len(v23)
                            }
                        } else {
                            if true {
                                v3 = ("hello" + (v3 + ""))
                                var v36 uint64 = 0
                                switch "hello" {
                                case "a":
                                    fmt.Println(v1, v2)
                                    v36 = (v36 ^ v36)
                                    var v37 uint64 = (^(v36))
                                    _ = v37
                                }
                                if false {
                                    if true {
                                        switch "foo bar" {
                                        case "":
                                            _ = len(v23)
                                        case "\n\t":
                                            v38 := GenS5{F0: 662}
                                            _ = v38
                                        }
                                        switch (v3 + "foo bar") {
                                        case "a":
                                            if false {
                                                v4 = ((v4 + (((v4 + "") + v4) + v3)) + (((((v4 + v3) + (((((v3 + "foo bar") + ("" + "hello")) + (((v4 + v3) + v4) + v3)) + v4) + v4)) + v4) + (((v4 + "hello") + ("foo bar" + (v4 + v4))) + v3)) + (("" + v3) + "a")))
                                            } else {
                                                if len(v23) >= 2 {
                                                    v23 = v23[1:len(v23)]
                                                }
                                                var v39 uint = 1
                                                var v40 uint8 = 127
                                                if false {
                                                    if (false || false) {
                                                        if len(v23) > 0 {
                                                            v23[0] = (fi33 >> ((v39) & 63))
                                                        }
                                                        if len(v23) > 0 {
                                                            v23[0] = fi31
                                                        }
                                                        var v41 int8 = -69
                                                        v41 -= v41
                                                        var v42 float32 = 0.1
                                                        _ = v41
                                                        _ = v42
                                                    } else {
                                                        _ = len(v23)
                                                        var v43 int = (fi33 / (((fi31 ^ -378)) | 1))
                                                        _ = v43
                                                    }
                                                    if false {
                                                        v23 = append(v23, int(v2))
                                                        var v44 int16 = (v24 + (v24 + 32767))
                                                        if (255 < v39) {
                                                            if (0 != 1) {
                                                                var v45 int16 = 1000
                                                                _ = v45
                                                            } else {
                                                                if false {
                                                                    var v46 uint32 = 42
                                                                    switch -2147483648 {
                                                                    case 972:
                                                                        var v47 uint32 = v46
                                                                        _ = v47
                                                                    case 42:
                                                                        v1 = v1
                                                                    case 118:
                                                                        v39 -= (v39 & (v39 * v39))
                                                                        switch -1 {
                                                                        case 2147483647:
                                                                            var v48 uint8 = (v40 & v40)
                                                                            if (2 == 858) {
                                                                                if true {
                                                                                    switch (^(fi31)) {
                                                                                    case 97:
                                                                                        fmt.Println(v3)
                                                                                    case 9223372036854775807:
                                                                                        var v49 uint = 65535
                                                                                        _ = v49
                                                                                    default:
                                                                                        if len(v23) > 0 {
                                                                                            v23[0] = (fi31 * fi31)
                                                                                        }
                                                                                    }
                                                                                    switch fi33 {
                                                                                    case -803:
                                                                                        if (float32(0.0) >= float32(1e10)) {
                                                                                            switch v4 {
                                                                                            case "":
                                                                                                if (0.1 != 3.14159) {
                                                                                                    v46 = (v46 + (v46 ^ (v46 % (((v46 + (v46 << (((v39 / ((v39) | 1))) & 63)))) | 1))))
                                                                                                } else {
                                                                                                    switch v3 {
                                                                                                    case "\n\t":
                                                                                                        var v50 int8 = v2
                                                                                                        _ = v50
                                                                                                    case "":
                                                                                                        var v51 int8 = int8(v1)
                                                                                                        _ = v51
                                                                                                    case "foo bar":
                                                                                                        if true {
                                                                                                            switch "a" {
                                                                                                            case "\n\t":
                                                                                                                v39 += v39
                                                                                                            }
                                                                                                            v4 = "\n\t"
                                                                                                            v52 := GenS6{F0: (true && (!false)), F1: v40}
                                                                                                            _ = v52
                                                                                                        } else {
                                                                                                            v53 := GenS4{F0: v24, F1: 127, F2: (v24 <= int16(v36))}
                                                                                                            _ = len(v23)
                                                                                                            _ = v53
                                                                                                        }
                                                                                                    default:
                                                                                                        var v54 uint8 = (^(v40))
                                                                                                        _ = v54
                                                                                                    }
                                                                                                }
                                                                                            case "\n\t":
                                                                                                switch "hello" {
                                                                                                case "hello":
                                                                                                    v46 = v46
                                                                                                    var v55 int32 = int32(v46)
                                                                                                    v24 = v44
                                                                                                    v46 = 2
                                                                                                    var v56 int8 = v2
                                                                                                    _ = v55
                                                                                                    _ = v56
                                                                                                case "foo bar":
                                                                                                    if true {
                                                                                                        if false {
                                                                                                            fmt.Println(v1, v36, v40)
                                                                                                            fmt.Println(v1, v36, v1)
                                                                                                            if ("\n\t" == (v4 + v3)) {
                                                                                                                _ = len(v23)
                                                                                                                var v57 uint16 = uint16(v39)
                                                                                                                _ = v57
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch v4 {
                                                                                                            case "foo bar":
                                                                                                                v48 = (v40 << ((uint(v46)) & 63))
                                                                                                            default:
                                                                                                                var v58 uint16 = 2
                                                                                                                _ = v58
                                                                                                            }
                                                                                                            v59 := GenS6{F0: (410 != 42), F1: v40}
                                                                                                            _ = v59
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    v36 = (v36 * 18446744073709551615)
                                                                                                }
                                                                                                v48 = (v40 / ((212) | 1))
                                                                                            case "hello":
                                                                                                switch (fi33 + fi31) {
                                                                                                case 42:
                                                                                                    var v60 int64 = 9223372036854775807
                                                                                                    var v61 string = v3
                                                                                                    _ = v60
                                                                                                    _ = v61
                                                                                                default:
                                                                                                    fmt.Println(v36, v48, v2)
                                                                                                }
                                                                                                var v62 int64 = int64(v48)
                                                                                                v24 -= (v44 * 32767)
                                                                                                switch 181 {
                                                                                                case 802:
                                                                                                    v63 := GenS1{F0: (1e10 <= (0.1 - 1.5)), F1: v1}
                                                                                                    _ = v63
                                                                                                case 42:
                                                                                                    fmt.Println(v4, v40)
                                                                                                default:
                                                                                                    var v64 float32 = (0.1 * (3.14159 + 0.1))
                                                                                                    _ = v64
                                                                                                }
                                                                                                if (v2 != 78) {
                                                                                                    fmt.Println(v48, v44, v39)
                                                                                                }
                                                                                                _ = v62
                                                                                            }
                                                                                            v65 := GenS6{F0: (uint32(v39) < (v46 / ((42) | 1))), F1: genFn0()}
                                                                                            _ = v65
                                                                                        } else {
                                                                                            if ((v40 ^ 91) >= (v48 & 168)) {
                                                                                                var v66 int = 1
                                                                                                _ = v66
                                                                                            } else {
                                                                                                if false {
                                                                                                    v1 -= v1
                                                                                                    if true {
                                                                                                        var v67 int = (fi33 | (fi33 & fi33))
                                                                                                        if len(v23) >= 2 {
                                                                                                            v23 = v23[1:len(v23)]
                                                                                                        }
                                                                                                        _ = v67
                                                                                                    } else {
                                                                                                        if len(v23) > 0 {
                                                                                                            v23[0] = -743
                                                                                                        }
                                                                                                    }
                                                                                                    var v68 float32 = 1e10
                                                                                                    _ = v68
                                                                                                }
                                                                                            }
                                                                                            v36 += (v36 / (((v36 ^ 1)) | 1))
                                                                                        }
                                                                                    case 4294967295:
                                                                                        v36 = (^(v36))
                                                                                    }
                                                                                    if len(v23) >= 2 {
                                                                                        v23 = v23[1:len(v23)]
                                                                                    }
                                                                                } else {
                                                                                    v40 += v40
                                                                                    if (((0.0 / 1.0) + float32((3.14159 * float64(3.14159)))) == 0.1) {
                                                                                        v39 -= 827
                                                                                        var v69 uint16 = 0
                                                                                        v69 -= (v69 >> ((2) & 63))
                                                                                        _ = v69
                                                                                    } else {
                                                                                        if len(v23) >= 2 {
                                                                                            v23 = v23[1:len(v23)]
                                                                                        }
                                                                                    }
                                                                                    var v70 int = -2147483648
                                                                                    _ = v70
                                                                                }
                                                                                v44 -= (v44 >> ((v39) & 63))
                                                                            } else {
                                                                                v4 = "hello"
                                                                            }
                                                                            if ("\n\t" == ((v4 + ((v4 + ((v4 + (v4 + ("a" + v4))) + v4)) + (v3 + v4))) + ("foo bar" + v3))) {
                                                                                fmt.Println(v44)
                                                                                v71 := GenS3{F0: -9223372036854775808}
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    v72 := GenS6{F0: false, F1: genFn0()}
                                                                                    _ = v72
                                                                                case "foo bar":
                                                                                    var v73 uint16 = 0
                                                                                    _ = v73
                                                                                default:
                                                                                    v44 = (v44 + -781)
                                                                                }
                                                                                _ = v71
                                                                            } else {
                                                                                if true {
                                                                                    var v74 uint8 = 59
                                                                                    if (v1 < 0) {
                                                                                        fmt.Println(v39, v46, v36)
                                                                                    } else {
                                                                                        if ((v1 + (v1 << ((uint(v40)) & 63))) >= (v1 >> (((^(v39))) & 63))) {
                                                                                            fmt.Println(v2, v74)
                                                                                            var v75 bool = (!true)
                                                                                            switch v3 {
                                                                                            case "\n\t":
                                                                                                v44 += v44
                                                                                            default:
                                                                                                var v76 int32 = (v1 ^ v1)
                                                                                                _ = v76
                                                                                            }
                                                                                            _ = v75
                                                                                        }
                                                                                        if (v48 > v40) {
                                                                                            v46 -= (v46 ^ (v46 & 2147483647))
                                                                                            var v77 int8 = 1
                                                                                            v78 := GenS3{F0: -9223372036854775808}
                                                                                            v36 = uint64(v46)
                                                                                            v36 = (v36 / (((v36 % (((v36 ^ (v36 - (v36 * v36)))) | 1))) | 1))
                                                                                            _ = v77
                                                                                            _ = v78
                                                                                        } else {
                                                                                            if (2147483647 >= 1) {
                                                                                                switch fi31 {
                                                                                                case 730:
                                                                                                    v4 = ((((v4 + ("a" + v3)) + (v4 + v3)) + (v3 + v3)) + (v3 + ((v4 + v4) + v3)))
                                                                                                    v3 = "hello"
                                                                                                    var v79 int32 = v1
                                                                                                    if len(v23) > 0 {
                                                                                                        v23[0] = (fi31 / ((298) | 1))
                                                                                                    }
                                                                                                    var v80 bool = true
                                                                                                    _ = v79
                                                                                                    _ = v80
                                                                                                case -705:
                                                                                                    var v81 uint16 = 255
                                                                                                    _ = v81
                                                                                                }
                                                                                                if len(v23) > 0 {
                                                                                                    v23[0] = fi31
                                                                                                }
                                                                                                var v82 float32 = 0.1
                                                                                                v24 -= (v24 << (((v39 >> ((uint(v46)) & 63))) & 63))
                                                                                                _ = v82
                                                                                            } else {
                                                                                                var v83 bool = (!(false || true))
                                                                                                _ = v83
                                                                                            }
                                                                                            fmt.Println(v2)
                                                                                        }
                                                                                    }
                                                                                    v23 = append(v23, fi31)
                                                                                    _ = v74
                                                                                }
                                                                                if (true || true) {
                                                                                    v84 := GenS7{F0: fi33, F1: 255}
                                                                                    v85 := GenS8{F0: v36, F1: (^(v36)), F2: (^(v46))}
                                                                                    _ = v84
                                                                                    _ = v85
                                                                                } else {
                                                                                    fmt.Println(v24, v24)
                                                                                }
                                                                                v86 := GenS9{F0: v48, F1: (v46 / (((^(v46))) | 1)), F2: 2147483647}
                                                                                v23 = append(v23, 1)
                                                                                _ = v86
                                                                            }
                                                                            _ = v48
                                                                        case 282:
                                                                            if ((v1 - -2147483648) < (v1 << (((v39 * 4294967295)) & 63))) {
                                                                                var v87 uint32 = (v46 << ((v39) & 63))
                                                                                v2 += v2
                                                                                fmt.Println(v24)
                                                                                _ = v87
                                                                            }
                                                                        }
                                                                        v4 = (v4 + (v4 + v3))
                                                                        v24 = v24
                                                                        if (!false) {
                                                                            switch (((v3 + (v4 + (((((("foo bar" + v3) + "foo bar") + "a") + ("foo bar" + "")) + (("\n\t" + v3) + (v3 + (v3 + (v3 + "foo bar"))))) + v3))) + ("" + "hello")) + ("hello" + "foo bar")) {
                                                                            case "\n\t":
                                                                                var v88 string = (v4 + v3)
                                                                                _ = v88
                                                                            case "foo bar":
                                                                                v89 := GenS4{F0: (v24 / ((v44) | 1)), F1: v40, F2: false}
                                                                                fmt.Println(v3, fi31)
                                                                                _ = v89
                                                                            case "":
                                                                                v36 += (v36 | 4294967295)
                                                                            }
                                                                        } else {
                                                                            v40 -= v40
                                                                        }
                                                                        if len(v23) > 0 {
                                                                            v23[0] = fi31
                                                                        }
                                                                        fmt.Println(v44)
                                                                        switch (fi31 / ((fi31) | 1)) {
                                                                        case 9223372036854775807:
                                                                            var v90 float32 = ((-2.25 / 1.5) - 0.1)
                                                                            _ = v90
                                                                        case 2147483647:
                                                                            v23 = append(v23, 127)
                                                                        default:
                                                                            var v91 int32 = v1
                                                                            _ = v91
                                                                        }
                                                                        if (fi33 >= fi31) {
                                                                            v36 = uint64(v46)
                                                                            if true {
                                                                                var v92 uint32 = (v46 << ((v39) & 63))
                                                                                var v93 uint8 = 128
                                                                                switch v3 {
                                                                                case "foo bar":
                                                                                    if ((v44 | int16(v36)) == 1) {
                                                                                        if ((3.14159 * -2.25) != 3.14159) {
                                                                                            v44 += v24
                                                                                            if (true && false) {
                                                                                                v94 := GenS2{F0: -750}
                                                                                                _ = v94
                                                                                            }
                                                                                            v95 := GenS10{F0: (v39 & (v39 / ((v39) | 1)))}
                                                                                            _ = v95
                                                                                        } else {
                                                                                            v3 = (v3 + v4)
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v3, v46)
                                                                                            case "hello":
                                                                                                switch int(v1) {
                                                                                                case 4294967295:
                                                                                                    v40 = (v93 + (v40 ^ 0))
                                                                                                    v2 += v2
                                                                                                default:
                                                                                                    v96 := GenS11{F0: 42, F1: (^(v44))}
                                                                                                    _ = v96
                                                                                                }
                                                                                            default:
                                                                                                v24 = 1
                                                                                            }
                                                                                            var v97 float64 = 3.14159
                                                                                            if (!(uint(v36) == 574)) {
                                                                                                v23 = append(v23, (fi31 >> (((v39 % ((v39) | 1))) & 63)))
                                                                                                var v98 int16 = (v24 + (-(v24)))
                                                                                                _ = v98
                                                                                            } else {
                                                                                                switch 128 {
                                                                                                case -9223372036854775808:
                                                                                                    v99 := GenS12{F0: v44, F1: v3, F2: (v2 >> (((^(v39))) & 63))}
                                                                                                    _ = v99
                                                                                                case 0:
                                                                                                    if ((!(false || ((671 >= 42) || (false || (!(true && true)))))) && (507 <= 1)) {
                                                                                                        switch (((v3 + "a") + "hello") + v3) {
                                                                                                        case "":
                                                                                                            fmt.Println(v24)
                                                                                                        case "\n\t":
                                                                                                            v1 += (v1 % ((int32(v2)) | 1))
                                                                                                        default:
                                                                                                            v100 := GenS12{F0: 1, F1: ("\n\t" + (v4 + "\n\t")), F2: (v2 & v2)}
                                                                                                            _ = v100
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            switch fi31 {
                                                                                            case 818:
                                                                                                if true {
                                                                                                    v36 = (v36 + (v36 | v36))
                                                                                                    fmt.Println(fi33)
                                                                                                }
                                                                                            case 2147483647:
                                                                                                var v101 int8 = (v2 / ((int8(v39)) | 1))
                                                                                                _ = v101
                                                                                            default:
                                                                                                switch (fi31 ^ -9223372036854775808) {
                                                                                                case -427:
                                                                                                    v102 := GenS13{F0: v1}
                                                                                                    _ = v102
                                                                                                case 4294967295:
                                                                                                    var v103 uint = 18446744073709551615
                                                                                                    _ = v103
                                                                                                }
                                                                                            }
                                                                                            _ = v97
                                                                                        }
                                                                                        var v104 uint = (^(v39))
                                                                                        _ = v104
                                                                                    }
                                                                                case "\n\t":
                                                                                    var v105 float64 = 0.1
                                                                                    _ = v105
                                                                                case "hello":
                                                                                    v106 := GenS14{F0: 0}
                                                                                    _ = v106
                                                                                default:
                                                                                    v107 := GenS2{F0: v1}
                                                                                    _ = v107
                                                                                }
                                                                                _ = v92
                                                                                _ = v93
                                                                            }
                                                                            v108 := GenS15{F0: v46, F1: 0, F2: (^(v39))}
                                                                            _ = v108
                                                                        } else {
                                                                            if len(v23) > 0 {
                                                                                v23[0] = -9223372036854775808
                                                                            }
                                                                            switch (fi31 - fi31) {
                                                                            case 9223372036854775807:
                                                                                v109 := GenS16{F0: (v1 + (^(v1))), F1: -620, F2: v46}
                                                                                _ = v109
                                                                            case -809:
                                                                                _ = len(v23)
                                                                                if len(v23) > 0 {
                                                                                    v23[0] = int(v40)
                                                                                }
                                                                                var v110 int = -1
                                                                                switch "a" {
                                                                                case "":
                                                                                    var v111 uint64 = 32
                                                                                    _ = v111
                                                                                }
                                                                                fmt.Println(v2)
                                                                                if true {
                                                                                    switch v4 {
                                                                                    case "":
                                                                                        var v112 int8 = v2
                                                                                        _ = v112
                                                                                    case "hello":
                                                                                        _ = len(v23)
                                                                                    }
                                                                                }
                                                                                _ = v110
                                                                            case 128:
                                                                                fmt.Println(v4, v36)
                                                                            default:
                                                                                v39 -= v39
                                                                            }
                                                                        }
                                                                    default:
                                                                        v23 = append(v23, fi33)
                                                                    }
                                                                    fmt.Println(v3, v40)
                                                                    _ = v46
                                                                } else {
                                                                    fmt.Println(fi33)
                                                                }
                                                                v23 = append(v23, fi31)
                                                                fmt.Println(v24)
                                                            }
                                                            if false {
                                                                v4 = "hello"
                                                                v24 = -70
                                                                v1 -= (^(v1))
                                                            } else {
                                                                var v113 float32 = 0.0
                                                                v40 = (v40 << (((v39 >> (((v39 + (v39 | v39))) & 63))) & 63))
                                                                var v114 uint64 = 487
                                                                _ = v113
                                                                _ = v114
                                                            }
                                                            v115 := GenS15{F0: 2, F1: 42, F2: (v39 << ((v39) & 63))}
                                                            v116 := GenS6{F0: true, F1: (v40 | v40)}
                                                            _ = v115
                                                            _ = v116
                                                        } else {
                                                            v4 = ""
                                                        }
                                                        var v117 bool = true
                                                        _ = v117
                                                        _ = v44
                                                    }
                                                } else {
                                                    if false {
                                                        var v118 uint64 = v36
                                                        _ = v118
                                                    } else {
                                                        var v119 float64 = (-2.25 / 1e10)
                                                        _ = v119
                                                    }
                                                    var v120 float32 = float32((float64(0.0) - (3.14159 - (0.0 + float64(float32(0.1))))))
                                                    _ = v120
                                                }
                                                _ = v39
                                                _ = v40
                                            }
                                        default:
                                            var v121 uint16 = 735
                                            _ = v121
                                        }
                                    }
                                    var v122 int8 = (v2 / (((v2 << ((3) & 63))) | 1))
                                    _ = len(v23)
                                    if len(v23) > 0 {
                                        v23[0] = (fi33 * fi31)
                                    }
                                    _ = v122
                                } else {
                                    fmt.Println(v1, fi33, v1)
                                    v123 := GenS17{F0: 900}
                                    var v124 string = v3
                                    _ = v123
                                    _ = v124
                                }
                                _ = v36
                            }
                            fmt.Println(v24)
                            var v125 float32 = 3.14159
                            _ = v125
                        }
                        _ = fi33
                        _ = v34
                    }
                    switch v3 {
                    case "":
                        v1 -= (v1 << ((42) & 63))
                    case "foo bar":
                        if false {
                            switch fi31 {
                            case 4294967295:
                                var v126 float32 = 0.0
                                _ = v126
                            case -20:
                                if len(v23) > 0 {
                                    v23[0] = (fi31 << ((1) & 63))
                                }
                            case -278:
                                v3 = "a"
                            }
                        } else {
                            if (2147483648 <= 451) {
                                fmt.Println(v3, v24, fi31)
                                switch fi31 {
                                case 42:
                                    var v127 uint64 = 2
                                    _ = v127
                                default:
                                    _ = len(v23)
                                }
                                _ = len(v23)
                            } else {
                                v128 := GenS1{F0: true, F1: v1}
                                _ = v128
                            }
                            if (((1e10 + 1e10) >= 1.5) || true) {
                                v1 = v1
                                var v129 bool = false
                                v24 = (v24 ^ v24)
                                _ = v129
                            }
                            v4 = v3
                        }
                    case "\n\t":
                        fmt.Println(v2, v2)
                    default:
                        _ = len(v23)
                    }
                }
                for fi130 := 0; fi130 < 11; fi130++ {
                    var v131 int32 = v1
                    v132 := GenS10{F0: 18446744073709551615}
                    var v133 uint16 = 0
                    _ = fi130
                    _ = v131
                    _ = v132
                    _ = v133
                }
                v4 = v3
                _ = fi31
            }
        case 32767:
            var v134 uint = 255
            _ = v134
        default:
            if (1 == 42) {
                fmt.Println(v2, v1, v4)
                fmt.Println(v24)
                var v135 bool = false
                _ = v135
            } else {
                for fi136 := 0; fi136 < 12; fi136++ {
                    if (v1 > v1) {
                        if false {
                            var v137 int64 = 373
                            var v138 bool = true
                            v137 += v137
                            _ = v137
                            _ = v138
                        }
                    }
                    _ = fi136
                }
            }
        }
    }
    fmt.Println(v2)
    if (v1 != v1) {
        v1 = v1
        var v139 string = (v3 + v4)
        if (0 <= 751) {
            if false {
                v140 := GenS18{F0: true, F1: 255, F2: v2}
                _ = v140
            } else {
                for fi141 := 0; fi141 < 9; fi141++ {
                    for fi142 := 0; fi142 < 2; fi142++ {
                        v143 := GenS4{F0: 42, F1: 105, F2: (-17 > (v2 + (v2 & v2)))}
                        var v144 string = v4
                        _ = fi142
                        _ = v143
                        _ = v144
                    }
                    for fi145 := 0; fi145 < 16; fi145++ {
                        if (138 < 196) {
                            var v146 uint16 = 65535
                            switch -219 {
                            case 42:
                                fmt.Println(v2, v1, v1)
                            case 127:
                                var v147 uint = 4294967295
                                _ = v147
                            case 32767:
                                v1 = 2147483647
                                var v148 uint8 = 128
                                _ = v148
                            }
                            if len(v23) > 0 {
                                v23[0] = (fi141 + fi141)
                            }
                            _ = v146
                        } else {
                            v149 := GenS19{F0: fi141, F1: 42}
                            _ = v149
                        }
                        var v150 int64 = -404
                        v24 = int16(v150)
                        var v151 int = fi145
                        _ = fi145
                        _ = v150
                        _ = v151
                    }
                    _ = fi141
                }
                v2 = v2
                if ((-9223372036854775808 >= -486) || (false && true)) {
                    v152 := GenS20{F0: (^(v24))}
                    _ = v152
                } else {
                    if false {
                        if false {
                            v153 := GenS10{F0: 2}
                            _ = v153
                        } else {
                            var v154 int64 = 1
                            v1 = 63
                            switch (((v4 + v139) + (v3 + "")) + "hello") {
                            case "hello":
                                v3 = v139
                            case "\n\t":
                                v155 := GenS17{F0: 1}
                                _ = v155
                            default:
                                for fi156 := 0; fi156 < 14; fi156++ {
                                    switch (fi156 | int(v1)) {
                                    case -826:
                                        v139 = ((v139 + v139) + v139)
                                    case 32768:
                                        for fi157 := 0; fi157 < 15; fi157++ {
                                            var v158 string = v3
                                            if len(v23) >= 2 {
                                                v23 = v23[1:len(v23)]
                                            }
                                            v23 = append(v23, (fi156 >> ((520) & 63)))
                                            var v159 uint8 = 255
                                            _ = fi157
                                            _ = v158
                                            _ = v159
                                        }
                                    case 127:
                                        v4 = "foo bar"
                                    }
                                    if len(v23) > 0 {
                                        v23[0] = fi156
                                    }
                                    fmt.Println(v154, fi156)
                                    _ = fi156
                                }
                            }
                            _ = v154
                        }
                        fmt.Println(v4, v4)
                        if ((-268 <= -1) && false) {
                            v2 = v2
                            v3 = ((v139 + (((v139 + (v3 + v4)) + v4) + v4)) + v3)
                        } else {
                            fmt.Println(v139, v24, v139)
                            for fi160 := 0; fi160 < 9; fi160++ {
                                var v161 uint = 0
                                if len(v23) > 0 {
                                    v23[0] = -527
                                }
                                _ = fi160
                                _ = v161
                            }
                            v1 += v1
                        }
                    }
                    v162 := GenS21{F0: 853}
                    _ = v162
                }
            }
        } else {
            fmt.Println(v4, v24)
            var v163 uint64 = 0
            for fi164 := 0; fi164 < 12; fi164++ {
                if false {
                    switch "\n\t" {
                    case "\n\t":
                        switch -451 {
                        case -3:
                            v139 = v4
                        case 2147483647:
                            _ = len(v23)
                        default:
                            v4 = "foo bar"
                        }
                    case "hello":
                        v165 := GenS22{F0: 255, F1: (^(v1))}
                        _ = v165
                    }
                    _ = len(v23)
                    fmt.Println(v3, v24)
                }
                var v166 float32 = 0.1
                _ = fi164
                _ = v166
            }
            _ = v163
        }
        _ = v139
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
}
