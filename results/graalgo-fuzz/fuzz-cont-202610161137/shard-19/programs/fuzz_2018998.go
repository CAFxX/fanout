package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 string; F1 uint64; F2 uint16 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 int16; F2 int8 }
type GenS8 struct { F0 int64 }

func main() {
    var v1 int32 = -73
    var v2 string = ("a" + ("foo bar" + "\n\t"))
    var v3 int64 = 42
    var v4 uint32 = 220
    var v5 uint64 = 0
    v6 := []int{-9223372036854775808, 4294967295, 9223372036854775807, 0}
    var v7 int16 = -32768
    fmt.Println(v1)
    var v8 int = -1
    switch 127 {
    case 128:
        fmt.Println(v7)
    case -461:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case 1:
        if (255 >= 32767) {
            _ = len(v6)
            if true {
                switch v8 {
                case -9223372036854775808:
                    fmt.Println(v8, v5, v8)
                default:
                    switch v8 {
                    case -340:
                        v3 = (v3 - v3)
                    default:
                        for fi9 := 0; fi9 < 17; fi9++ {
                            for fi10 := 0; fi10 < 8; fi10++ {
                                var v11 uint32 = 0
                                var v12 float64 = 3.14159
                                var v13 uint32 = 569
                                _ = fi10
                                _ = v11
                                _ = v12
                                _ = v13
                            }
                            if true {
                                var v14 float64 = 1.5
                                if len(v6) > 0 {
                                    v6[0] = (v8 & 1)
                                }
                                for fi15 := 0; fi15 < 11; fi15++ {
                                    var v16 uint64 = 178
                                    _ = fi15
                                    _ = v16
                                }
                                _ = v14
                            }
                            _ = fi9
                        }
                    }
                }
                if ((v2 >= "") && true) {
                    var v17 uint32 = (v4 << ((4294967295) & 63))
                    _ = v17
                } else {
                    v3 -= v3
                    v4 = v4
                    var v18 uint32 = v4
                    _ = v18
                }
            }
        }
    }
    var v19 int16 = -32768
    fmt.Println(v7, v1)
    for fi20 := 0; fi20 < 1; fi20++ {
        for fi21 := 0; fi21 < 17; fi21++ {
            v22 := GenS1{F0: "a", F1: v1}
            v23 := GenS1{F0: (v2 + v2), F1: (v1 % ((903) | 1))}
            if false {
                var v24 uint = 451
                var v25 int32 = (v1 & (v1 / ((v1) | 1)))
                var v26 uint16 = 0
                _ = v24
                _ = v25
                _ = v26
            } else {
                switch v8 {
                case 4294967295:
                    switch (^(v8)) {
                    case -1:
                        var v27 float64 = (1.5 * 0.1)
                        _ = v27
                    case 2147483647:
                        var v28 string = v2
                        _ = v28
                    case 1:
                        var v29 int32 = (v1 >> ((947) & 63))
                        var v30 uint32 = 4294967295
                        v7 -= v19
                        if (float32(((float64((0.1 / 1.0)) + 3.14159) + float64(3.14159))) < 0.1) {
                            switch (fi21 | (fi21 / (((fi20 | fi21)) | 1))) {
                            case 810:
                                v1 = (v1 >> ((337) & 63))
                            case -2147483648:
                                switch (v2 + v2) {
                                case "\n\t":
                                    _ = len(v6)
                                case "hello":
                                    fmt.Println(v19)
                                }
                            case 32768:
                                var v31 int16 = (v19 >> ((4294967295) & 63))
                                _ = v31
                            }
                            var v32 int16 = (^(v19))
                            _ = len(v6)
                            _ = v32
                        } else {
                            var v33 float64 = -2.25
                            v34 := GenS2{F0: v2, F1: 213, F2: 789}
                            _ = v33
                            _ = v34
                        }
                        _ = v29
                        _ = v30
                    default:
                        v6 = append(v6, fi20)
                    }
                    v35 := GenS2{F0: v2, F1: 1, F2: 253}
                    v19 = -587
                    if ((3.14159 * 1.5) != 0.0) {
                        if true {
                            var v36 uint16 = 0
                            fmt.Println(v3, v19, v5)
                            v2 = v2
                            _ = v36
                        } else {
                            fmt.Println(v8, fi20, v3)
                            if true {
                                if true {
                                    v2 = ((((v2 + (((v2 + v2) + "a") + "")) + v2) + v2) + (v2 + (v2 + "a")))
                                    fmt.Println(v4)
                                    v37 := GenS1{F0: v2, F1: 0}
                                    _ = v37
                                } else {
                                    if true {
                                        var v38 int32 = (^(v1))
                                        _ = v38
                                    }
                                    if true {
                                        v6 = append(v6, (fi21 | fi21))
                                    } else {
                                        switch ((v2 + "hello") + v2) {
                                        case "a":
                                            if len(v6) > 0 {
                                                v6[0] = (fi20 - 612)
                                            }
                                        }
                                    }
                                    if true {
                                        if (!(false || (v7 >= 0))) {
                                            if len(v6) > 0 {
                                                v6[0] = (fi20 << ((1) & 63))
                                            }
                                            if false {
                                                if false {
                                                    v1 -= v1
                                                    v39 := GenS3{F0: true, F1: false}
                                                    _ = v39
                                                }
                                            }
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            fmt.Println(v4, fi20)
                                        }
                                        v40 := GenS4{F0: "foo bar", F1: 2}
                                        if (!((-128 <= 42) || false)) {
                                            switch 128 {
                                            case -9223372036854775808:
                                                fmt.Println(v4, v2, v19)
                                            case 2147483647:
                                                if (true || true) {
                                                    v1 = v1
                                                }
                                                fmt.Println(v4)
                                            case 42:
                                                if ((0.1 / 3.14159) > (1.5 - 0.1)) {
                                                    if false {
                                                        fmt.Println(v1, v4, v4)
                                                        if true {
                                                            var v41 float64 = (((3.14159 / ((-2.25 + (1e10 - ((1e10 - 0.1) * float64(float32(1e10))))) + 0.1)) - -2.25) / -2.25)
                                                            if false {
                                                                v6 = append(v6, (fi21 | (^(fi20))))
                                                                var v42 float32 = (((float32(v41) + 1e10) - 0.1) - float32(float64((1e10 * (0.1 - (1e10 * ((1e10 + 0.1) + 1.5)))))))
                                                                v41 += (float64((v42 * (v42 * v42))) + ((v41 + v41) + 3.14159))
                                                                v7 = v19
                                                                _ = v42
                                                            } else {
                                                                fmt.Println(v41, v5)
                                                                if true {
                                                                    v43 := GenS2{F0: "", F1: v5, F2: 0}
                                                                    var v44 uint16 = 42
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                                if ((v41 - v41) >= (v41 * -2.25)) {
                                                                    v8 += int(v5)
                                                                    v41 = v41
                                                                    var v45 int = int(v19)
                                                                    _ = v45
                                                                }
                                                                if len(v6) > 0 {
                                                                    v6[0] = int(v5)
                                                                }
                                                            }
                                                            _ = v41
                                                        } else {
                                                            if (((v5 ^ (v5 & (v5 % ((4294967295) | 1)))) > v5) && false) {
                                                                v2 = "hello"
                                                            } else {
                                                                if false {
                                                                    var v46 bool = false
                                                                    var v47 float32 = (-2.25 * ((0.0 / 0.1) / float32(-2.25)))
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                            }
                                                        }
                                                        if true {
                                                            if (((v3 * (-(v3))) >= -9223372036854775808) || ((18446744073709551615 != 2) || true)) {
                                                                v6 = append(v6, (v8 & -2147483648))
                                                            } else {
                                                                var v48 int8 = int8(v7)
                                                                var v49 uint8 = 174
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                            var v50 uint = 414
                                                            _ = v50
                                                        }
                                                    }
                                                    v51 := GenS2{F0: v2, F1: (v5 ^ (v5 | (v5 ^ (v5 << ((255) & 63))))), F2: 722}
                                                    fmt.Println(v8, v7)
                                                    v52 := GenS4{F0: ("hello" + (("hello" + (v2 + (v2 + (("" + v2) + ("foo bar" + v2))))) + v2)), F1: 1}
                                                    v53 := GenS1{F0: v2, F1: v1}
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                }
                                            default:
                                                var v54 float64 = 3.14159
                                                switch v2 {
                                                case "foo bar":
                                                    if (!(!(-128 != 0))) {
                                                        if (v5 < 1) {
                                                            var v55 int16 = 431
                                                            if true {
                                                                v6 = append(v6, 441)
                                                            }
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            if (-49 < 42) {
                                                                var v56 string = v2
                                                                fmt.Println(v55)
                                                                v4 = v4
                                                                _ = v56
                                                            }
                                                            v4 += (v4 % ((2147483647) | 1))
                                                            fmt.Println(v8, fi20, v54)
                                                            if (1e10 != v54) {
                                                                if false {
                                                                    var v57 float32 = ((float32(v54) * 3.14159) + (3.14159 + 0.0))
                                                                    fmt.Println(v55, v19)
                                                                    _ = v57
                                                                } else {
                                                                    switch "" {
                                                                    case "a":
                                                                        v8 = v8
                                                                    case "hello":
                                                                        v4 -= 1
                                                                    default:
                                                                        var v58 string = ("\n\t" + v2)
                                                                        _ = v58
                                                                    }
                                                                    var v59 uint16 = 255
                                                                    v59 = (v59 | (v59 / ((v59) | 1)))
                                                                    _ = v59
                                                                }
                                                                v60 := GenS2{F0: (v2 + v2), F1: (^(v5)), F2: 65535}
                                                                if len(v6) > 0 {
                                                                    v6[0] = int(v1)
                                                                }
                                                                _ = v60
                                                            }
                                                            v7 = -1
                                                            switch "" {
                                                            case "hello":
                                                                var v61 int8 = -1
                                                                _ = len(v6)
                                                                if (v61 <= (v61 % ((v61) | 1))) {
                                                                    switch v8 {
                                                                    case 0:
                                                                        if (v2 < (("" + (v2 + v2)) + (v2 + v2))) {
                                                                            if (true && true) {
                                                                                v6 = append(v6, (fi21 >> ((18446744073709551615) & 63)))
                                                                            } else {
                                                                                fmt.Println(v55)
                                                                            }
                                                                            v62 := GenS3{F0: (-9223372036854775808 >= fi20), F1: ((v3 & (v3 % ((v3) | 1))) == 98)}
                                                                            _ = v62
                                                                        }
                                                                    case 42:
                                                                        v63 := GenS5{F0: v5}
                                                                        _ = v63
                                                                    }
                                                                    var v64 int8 = (v61 << ((65535) & 63))
                                                                    v3 -= v3
                                                                    _ = v64
                                                                }
                                                                _ = v61
                                                            case "\n\t":
                                                                if len(v6) > 0 {
                                                                    v6[0] = 42
                                                                }
                                                            case "foo bar":
                                                                v2 = ("\n\t" + v2)
                                                                v3 = 9223372036854775807
                                                                switch (((v2 + ("" + v2)) + v2) + v2) {
                                                                case "foo bar":
                                                                    var v65 int64 = v3
                                                                    _ = v65
                                                                case "":
                                                                    v6 = append(v6, fi21)
                                                                case "\n\t":
                                                                    if (!false) {
                                                                        var v66 uint16 = 32767
                                                                        if len(v6) > 0 {
                                                                            v6[0] = (fi21 % ((128) | 1))
                                                                        }
                                                                        _ = v66
                                                                    } else {
                                                                        fmt.Println(v19, v55, v1)
                                                                        if (false && true) {
                                                                            v67 := GenS4{F0: v2, F1: 2}
                                                                            _ = v67
                                                                        }
                                                                    }
                                                                default:
                                                                    v5 = (v5 & 18446744073709551615)
                                                                }
                                                            default:
                                                                switch ((v2 + (v2 + v2)) + v2) {
                                                                case "a":
                                                                    var v68 uint8 = 9
                                                                    _ = v68
                                                                case "foo bar":
                                                                    _ = len(v6)
                                                                case "hello":
                                                                    if false {
                                                                        var v69 float32 = -2.25
                                                                        var v70 float32 = (1e10 / float32((0.1 * -2.25)))
                                                                        v4 = v4
                                                                        _ = v69
                                                                        _ = v70
                                                                    }
                                                                default:
                                                                    var v71 float64 = 3.14159
                                                                    _ = v71
                                                                }
                                                            }
                                                            var v72 uint64 = v5
                                                            _ = v55
                                                            _ = v72
                                                        } else {
                                                            fmt.Println(fi20)
                                                        }
                                                        _ = len(v6)
                                                    }
                                                case "":
                                                    fmt.Println(v8, v4)
                                                    v7 += v7
                                                    v6 = append(v6, int(v3))
                                                case "\n\t":
                                                    v1 = int32(v19)
                                                default:
                                                    switch (v8 ^ fi21) {
                                                    case 60:
                                                        var v73 float32 = ((((0.1 - 0.1) - (1e10 * 1.5)) * 1.5) + ((float32((0.1 + v54)) + 0.1) + (-2.25 / (0.0 + (3.14159 + (0.1 + 1e10))))))
                                                        _ = v73
                                                    case 4294967295:
                                                        v4 = 31
                                                    case -581:
                                                        switch ("hello" + "foo bar") {
                                                        case "hello":
                                                            switch (-(v8)) {
                                                            case 127:
                                                                fmt.Println(v4, v54)
                                                            default:
                                                                var v74 int16 = (v7 - (v7 + (-(v7))))
                                                                _ = v74
                                                            }
                                                            var v75 int = (fi20 * v8)
                                                            _ = v75
                                                        case "\n\t":
                                                            var v76 uint8 = 2
                                                            _ = v76
                                                        case "foo bar":
                                                            v77 := GenS6{F0: 4294967295}
                                                            _ = v77
                                                        default:
                                                            if (true || (v1 == (v1 % ((0) | 1)))) {
                                                                v6 = append(v6, v8)
                                                                var v78 uint16 = 42
                                                                v6 = append(v6, fi21)
                                                                fmt.Println(v4)
                                                                _ = v78
                                                            } else {
                                                                v79 := GenS3{F0: (!false), F1: false}
                                                                v8 -= int(v3)
                                                                _ = v79
                                                            }
                                                        }
                                                    default:
                                                        v5 -= v5
                                                    }
                                                }
                                                _ = v54
                                            }
                                        }
                                        _ = v40
                                    }
                                }
                                if (127 == 42) {
                                    fmt.Println(v5, fi21, v7)
                                }
                            } else {
                                v80 := GenS4{F0: (((v2 + v2) + v2) + ("" + "hello")), F1: 922}
                                var v81 uint64 = v5
                                v82 := GenS4{F0: (v2 + (((("foo bar" + (v2 + v2)) + v2) + (v2 + v2)) + v2)), F1: 1}
                                var v83 uint16 = uint16(v5)
                                _ = v80
                                _ = v81
                                _ = v82
                                _ = v83
                            }
                            v84 := GenS2{F0: v2, F1: (v5 & 42), F2: 473}
                            _ = v84
                        }
                        v6 = append(v6, (fi21 / ((128) | 1)))
                        if true {
                            v5 += v5
                        } else {
                            if true {
                                fmt.Println(v2, v7)
                                v19 = (v19 ^ 42)
                                v2 = v2
                            } else {
                                fmt.Println(fi21, v7)
                                _ = len(v6)
                            }
                            v4 -= v4
                            v8 = 9223372036854775807
                        }
                        switch fi20 {
                        case 128:
                            var v85 float64 = 0.1
                            _ = v85
                        case -1:
                            switch v2 {
                            case "a":
                                v7 += 0
                            case "foo bar":
                                v86 := GenS7{F0: (fi20 - int(v1)), F1: (v7 ^ 0), F2: 1}
                                _ = v86
                            default:
                                if false {
                                    fmt.Println(v7, v1, fi20)
                                }
                            }
                            var v87 float32 = (float32(1e10) - (0.0 / 1.0))
                            _ = v87
                        case -398:
                            fmt.Println(v4, v2, v8)
                        }
                        var v88 uint16 = 825
                        _ = v88
                    }
                    if false {
                        var v89 int = (fi21 - fi20)
                        fmt.Println(v4, v5)
                        var v90 int32 = (-(v1))
                        v1 += (v1 ^ 2147483647)
                        v2 = v2
                        _ = v89
                        _ = v90
                    } else {
                        v6 = append(v6, fi20)
                        var v91 float64 = (-2.25 - 1.5)
                        v5 += 42
                        _ = v91
                    }
                    var v92 string = "hello"
                    fmt.Println(v1)
                    v93 := GenS8{F0: v3}
                    _ = v35
                    _ = v92
                    _ = v93
                case 128:
                    v2 = "\n\t"
                default:
                    var v94 float32 = -2.25
                    _ = v94
                }
                var v95 uint32 = 42
                _ = v95
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        v7 += v19
        _ = fi20
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v96 uint16 = 255
    var v97 string = "\n\t"
    v4 -= (v4 | v4)
    for fi98 := 0; fi98 < 10; fi98++ {
        if (((0.0 == 3.14159) && false) || true) {
            var v99 int8 = 0
            _ = v99
        }
        _ = fi98
    }
    fmt.Println(v8)
    var v100 uint32 = (^(v4))
    var v101 float32 = float32(float64(1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v100
    _ = v101
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v96
    _ = v97
}
