package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint }
type GenS2 struct { F0 uint; F1 string; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int; F2 uint }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 int16; F1 int64; F2 string }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS9 struct { F0 string; F1 uint32; F2 int64 }
type GenS10 struct { F0 uint; F1 int8 }
func genFn0(p0 uint) int64 {
    var v5 uint16 = 32767
    if (18446744073709551615 >= 994) {
        p0 = p0
        var v6 float64 = 3.14159
        _ = v6
    } else {
        v7 := GenS1{F0: 65535, F1: 678, F2: 0}
        v8 := []int{2147483647, 0, 2147483647, 4294967295}
        v5 = v5
        _ = v7
        _ = v8
    }
    if ("hello" == "hello") {
        v9 := []int{4294967295}
        v10 := GenS1{F0: 2, F1: 33, F2: (p0 % ((p0) | 1))}
        _ = v10
        _ = v9
    } else {
        var v11 uint = (p0 - 255)
        _ = v11
    }
    v5 -= (^(v5))
    var v12 uint64 = 4294967295
    _ = p0
    _ = v12
    _ = v5
    return 42
}
func genFn1() bool {
    var v13 int8 = 42
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := []int{(fi14 / (((fi14 / ((fi14) | 1))) | 1)), (fi14 ^ 9223372036854775807), -9223372036854775808}
        _ = fi14
        _ = v15
    }
    fmt.Println(v13)
    _ = v13
    return false
}

func main() {
    v1 := []uint32{888, 2147483648, 0}
    var v2 uint16 = 65535
    var v3 uint32 = 59
    var v4 uint16 = (^(v2))
    for fi16 := 0; fi16 < 13; fi16++ {
        if false {
            var v17 int32 = -1
            _ = v17
        }
        fmt.Println(v4, v4, fi16)
        v2 -= (v2 * 674)
        _ = fi16
    }
    if true {
        var v18 uint8 = 149
        var v19 float32 = ((0.0 + 0.1) * (1.5 - (3.14159 - -2.25)))
        _ = v18
        _ = v19
    }
    var v20 int16 = -1
    if (42 == -1) {
        for fi21 := 0; fi21 < 7; fi21++ {
            for fi22 := 0; fi22 < 10; fi22++ {
                var v23 int32 = -2147483648
                if ((-(v23)) == v23) {
                    switch "" {
                    case "foo bar":
                        if true {
                            if true {
                                switch (fi21 >> ((2) & 63)) {
                                case 128:
                                    v20 = (v20 >> ((1) & 63))
                                default:
                                    if false {
                                        v4 = (v2 % (((^(v4))) | 1))
                                        v4 -= (v2 | 32767)
                                    } else {
                                        switch fi21 {
                                        case 819:
                                            fmt.Println(v20)
                                        }
                                        v24 := GenS1{F0: v3, F1: 4294967295, F2: 4294967295}
                                        _ = v24
                                    }
                                }
                            } else {
                                var v25 string = "\n\t"
                                var v26 uint16 = (v4 >> ((2) & 63))
                                _ = v25
                                _ = v26
                            }
                            v27 := GenS1{F0: (^(v3)), F1: 2, F2: 4294967295}
                            _ = v27
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    case "":
                        if false {
                            switch ("" + "a") {
                            case "\n\t":
                                var v28 float64 = 0.1
                                _ = v28
                            default:
                                var v29 uint16 = 255
                                _ = v29
                            }
                            fmt.Println(v23, v20)
                        } else {
                            var v30 float64 = 0.0
                            v31 := GenS1{F0: (v3 - v3), F1: (v3 & v3), F2: 42}
                            _ = v30
                            _ = v31
                        }
                        switch "\n\t" {
                        case "foo bar":
                            switch "foo bar" {
                            case "a":
                                v32 := GenS2{F0: 592, F1: ("a" + "foo bar"), F2: 722}
                                _ = v32
                            }
                        case "a":
                            if false {
                                v20 -= v20
                                var v33 uint64 = 534
                                _ = v33
                            }
                        default:
                            if true {
                                fmt.Println(v23)
                                fmt.Println(v23, v2)
                            } else {
                                fmt.Println(v3)
                            }
                        }
                    }
                    v2 = (v4 % (((v2 / ((v4) | 1))) | 1))
                }
                _ = fi22
                _ = v23
            }
            for fi34 := 0; fi34 < 11; fi34++ {
                fmt.Println(fi21, fi34)
                v35 := GenS1{F0: (^(v3)), F1: 42, F2: 2}
                switch "" {
                case "a":
                    v3 = (v3 << ((42) & 63))
                case "hello":
                    if (!true) {
                        v4 -= (v4 << ((255) & 63))
                        var v36 int32 = 2147483647
                        _ = v36
                    } else {
                        fmt.Println(v20, v2)
                        fmt.Println(v4)
                        switch "hello" {
                        case "a":
                            v37 := GenS1{F0: v3, F1: (v3 - (v3 & 42)), F2: 42}
                            _ = v37
                        case "hello":
                            v1 = append(v1, 106)
                            switch ((("\n\t" + ("\n\t" + "\n\t")) + "a") + "\n\t") {
                            case "\n\t":
                                if false {
                                    var v38 int64 = -1
                                    fmt.Println(fi34, v4, v38)
                                    _ = v38
                                }
                            case "foo bar":
                                v39 := GenS1{F0: v3, F1: v3, F2: 65535}
                                _ = v39
                            }
                        case "foo bar":
                            v4 = 65535
                        default:
                            var v40 string = "a"
                            _ = v40
                        }
                    }
                }
                _ = fi34
                _ = v35
            }
            var v41 int = fi21
            _ = fi21
            _ = v41
        }
        for fi42 := 0; fi42 < 12; fi42++ {
            v1 = append(v1, (v3 >> ((255) & 63)))
            var v43 uint32 = 0
            var v44 uint32 = v3
            _ = fi42
            _ = v43
            _ = v44
        }
        if (!false) {
            fmt.Println(v3, v3)
            for fi45 := 0; fi45 < 12; fi45++ {
                v46 := GenS1{F0: 4294967295, F1: v3, F2: 0}
                var v47 int = (-(fi45))
                _ = len(v1)
                _ = fi45
                _ = v46
                _ = v47
            }
            v2 -= v2
        }
    } else {
        var v48 int32 = -640
        _ = v48
    }
    fmt.Println(v3, v20)
    v2 = v4
    v49 := GenS2{F0: 4294967295, F1: "", F2: 329}
    if true {
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v4 += 883
        }
    }
    for fi50 := 0; fi50 < 13; fi50++ {
        v4 = v4
        v3 -= (^(v3))
        if ((0.1 * -2.25) != ((1.5 / 1e10) / (0.1 * (0.1 - -2.25)))) {
            if (610 == 4294967295) {
                var v51 uint = uint(v4)
                _ = v51
            } else {
                if (676 != 1) {
                    _ = len(v1)
                    var v52 int16 = 42
                    _ = v52
                } else {
                    v53 := GenS3{F0: ("" + ("" + ("hello" + "a")))}
                    _ = v53
                }
                if (381 <= 1) {
                    var v54 int = (fi50 >> ((521) & 63))
                    _ = v54
                }
            }
            switch fi50 {
            case 128:
                v55 := GenS4{F0: 128, F1: (fi50 % ((32768) | 1)), F2: 80}
                _ = v55
            case 42:
                for fi56 := 0; fi56 < 10; fi56++ {
                    _ = len(v1)
                    v57 := GenS5{F0: genFn0(434), F1: ""}
                    switch fi50 {
                    case -373:
                        switch "\n\t" {
                        case "foo bar":
                            switch (-(fi50)) {
                            case 127:
                                v3 = 2
                            case 1:
                                v4 -= (^(v4))
                            case 0:
                                var v58 uint16 = (v4 - 32767)
                                _ = v58
                            default:
                                v3 += (^(v3))
                            }
                        case "":
                            var v59 string = ("" + "hello")
                            _ = v59
                        default:
                            v60 := GenS3{F0: "\n\t"}
                            _ = v60
                        }
                        var v61 uint16 = (v2 + (v4 | (v2 & 255)))
                        _ = v61
                    case 1:
                        v62 := GenS2{F0: 1, F1: ("a" + ("a" + ("foo bar" + ("a" + (("a" + "foo bar") + ""))))), F2: 18446744073709551615}
                        _ = v62
                    case 32768:
                        v1 = append(v1, (v3 / ((2) | 1)))
                    default:
                        if false {
                            if true {
                                v63 := GenS4{F0: fi50, F1: -9223372036854775808, F2: 2}
                                _ = v63
                            }
                            v3 = 4294967295
                            var v64 int8 = 100
                            v2 += 110
                            var v65 uint8 = uint8(v2)
                            _ = v64
                            _ = v65
                        } else {
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v66 := GenS4{F0: fi50, F1: (fi56 + fi50), F2: 255}
                                _ = v66
                            } else {
                                if true {
                                    var v67 uint16 = (v4 * v4)
                                    if ("\n\t" > "hello") {
                                        if false {
                                            var v68 int8 = -1
                                            switch (fi56 >> ((18446744073709551615) & 63)) {
                                            case -325:
                                                v69 := GenS6{F0: v20, F1: genFn0(42), F2: "hello"}
                                                _ = v69
                                            case 1:
                                                v67 = (v67 >> ((4294967295) & 63))
                                            }
                                            if (false && (572 == -75)) {
                                                v70 := GenS2{F0: 65535, F1: "foo bar", F2: 987}
                                                _ = v70
                                            }
                                            fmt.Println(v68, v2, v3)
                                            _ = v68
                                        }
                                        var v71 uint8 = 2
                                        v72 := GenS3{F0: (("hello" + ((("\n\t" + "") + "\n\t") + "a")) + ((("" + "a") + ("hello" + "a")) + ((("a" + "a") + (((("foo bar" + "a") + (("foo bar" + "\n\t") + "foo bar")) + "hello") + "hello")) + "a")))}
                                        if true {
                                            v73 := GenS5{F0: -1, F1: (("a" + "\n\t") + "")}
                                            fmt.Println(fi50, v67)
                                            v67 = 65535
                                            _ = v73
                                        }
                                        _ = v71
                                        _ = v72
                                    } else {
                                        v20 = (v20 & 42)
                                        switch (fi56 << ((0) & 63)) {
                                        case 2147483647:
                                            if ((0.1 + (3.14159 * (3.14159 / (-2.25 + (1e10 / 0.1))))) != 1e10) {
                                                if ((3.14159 >= 1.5) && (3.14159 == (1e10 * 3.14159))) {
                                                    if (true || (!true)) {
                                                        v74 := GenS6{F0: (v20 & (v20 - -1)), F1: 0, F2: ""}
                                                        switch fi56 {
                                                        case -974:
                                                            v75 := GenS7{F0: (v3 * 42), F1: 42}
                                                            _ = v75
                                                        case -396:
                                                            var v76 uint8 = 72
                                                            _ = v76
                                                        case 127:
                                                            if true {
                                                                fmt.Println(v20, v67)
                                                                var v77 uint64 = 1
                                                                switch -31 {
                                                                case -1:
                                                                    var v78 float64 = (0.1 * (1.5 + 0.1))
                                                                    _ = v78
                                                                case -9223372036854775808:
                                                                    var v79 uint8 = 196
                                                                    _ = v79
                                                                case 383:
                                                                    var v80 uint8 = uint8(v3)
                                                                    _ = v80
                                                                default:
                                                                    var v81 uint64 = (^(v77))
                                                                    _ = v81
                                                                }
                                                                _ = v77
                                                            } else {
                                                                switch "\n\t" {
                                                                case "":
                                                                    var v82 int8 = 1
                                                                    _ = v82
                                                                case "\n\t":
                                                                    fmt.Println(v2)
                                                                default:
                                                                    v3 += (v3 / ((644) | 1))
                                                                }
                                                            }
                                                        default:
                                                            var v83 uint16 = v67
                                                            _ = v83
                                                        }
                                                        if true {
                                                            var v84 int64 = 409
                                                            v20 -= -1
                                                            _ = v84
                                                        } else {
                                                            var v85 float64 = 1e10
                                                            if false {
                                                                v86 := GenS2{F0: 1, F1: "foo bar", F2: 255}
                                                                if true {
                                                                    v87 := GenS4{F0: (fi56 * fi56), F1: 810, F2: 2}
                                                                    v85 -= 0.1
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    _ = v87
                                                                } else {
                                                                    if false {
                                                                        if (true || true) {
                                                                            v88 := GenS6{F0: (v20 ^ v20), F1: 0, F2: "a"}
                                                                            fmt.Println(v4)
                                                                            _ = v88
                                                                        }
                                                                        if false {
                                                                            v67 = v4
                                                                        } else {
                                                                            v3 = (v3 >> ((622) & 63))
                                                                        }
                                                                    }
                                                                }
                                                                switch ("" + "") {
                                                                case "a":
                                                                    fmt.Println(v2, v85, v85)
                                                                default:
                                                                    fmt.Println(v2)
                                                                }
                                                                _ = v86
                                                            }
                                                            v3 = 712
                                                            _ = v85
                                                        }
                                                        var v89 int = -9223372036854775808
                                                        var v90 int32 = 45
                                                        v2 -= v4
                                                        _ = v74
                                                        _ = v89
                                                        _ = v90
                                                    } else {
                                                        var v91 uint = 255
                                                        if (!false) {
                                                            var v92 string = ("\n\t" + (("foo bar" + "\n\t") + "hello"))
                                                            switch (-(fi56)) {
                                                            case 1:
                                                                _ = len(v1)
                                                                v1 = append(v1, (^(v3)))
                                                                v1 = append(v1, (v3 * v3))
                                                                if false {
                                                                    v92 = v92
                                                                    v93 := GenS3{F0: v92}
                                                                    _ = v93
                                                                } else {
                                                                    v94 := GenS1{F0: 65535, F1: (v3 << ((65535) & 63)), F2: 1}
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (v3 | v3)
                                                                    }
                                                                    _ = v94
                                                                }
                                                                switch (v92 + ("" + "a")) {
                                                                case "hello":
                                                                    v1 = append(v1, (v3 - 595))
                                                                case "":
                                                                    v1 = append(v1, v3)
                                                                default:
                                                                    v3 += 42
                                                                }
                                                            case 128:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                v92 = v92
                                                            case -9223372036854775808:
                                                                var v95 bool = false
                                                                _ = v95
                                                            }
                                                            var v96 uint64 = 1
                                                            _ = v92
                                                            _ = v96
                                                        }
                                                        _ = v91
                                                    }
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        _ = len(v1)
                                                    default:
                                                        _ = len(v1)
                                                    }
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = (v3 % (((v3 % ((4294967295) | 1))) | 1))
                                                }
                                                var v97 uint64 = 42
                                                _ = v97
                                            } else {
                                                fmt.Println(v3, v4)
                                                var v98 uint = 65535
                                                _ = v98
                                            }
                                        case 4294967295:
                                            var v99 float32 = 3.14159
                                            _ = v99
                                        default:
                                            v100 := GenS5{F0: -540, F1: ""}
                                            v1 = append(v1, 4294967295)
                                            var v101 int32 = -1
                                            _ = v100
                                            _ = v101
                                        }
                                        v2 = 32767
                                        fmt.Println(v3, v3, v3)
                                        v102 := GenS8{F0: 2, F1: 1, F2: (v3 << ((uint(v4)) & 63))}
                                        _ = v102
                                    }
                                    _ = v67
                                } else {
                                    v2 = (^(v2))
                                }
                                v20 = (^(v20))
                                switch fi50 {
                                case 9223372036854775807:
                                    var v103 uint16 = (v2 & 2)
                                    _ = v103
                                default:
                                    v104 := GenS9{F0: "foo bar", F1: v3, F2: 1}
                                    _ = v104
                                }
                            }
                        }
                    }
                    _ = fi56
                    _ = v57
                }
            }
        } else {
            v105 := GenS10{F0: 921, F1: 108}
            _ = v105
        }
        _ = fi50
    }
    v106 := GenS2{F0: 78, F1: "\n\t", F2: 501}
    var v107 int32 = -148
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v106
    _ = v107
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v49
}
