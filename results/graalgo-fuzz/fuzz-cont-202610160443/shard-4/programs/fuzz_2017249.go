package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int64; F2 string }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 uint16; F1 string; F2 uint32 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 int64; F1 int8; F2 uint64 }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 uint16; F1 int16 }
func genFn0(p0 int64, p1 uint8) uint {
    v6 := GenS1{F0: 104, F1: 422, F2: 2}
    var v7 uint16 = 255
    if (!(!false)) {
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := []int{(fi8 * (fi8 * (^(fi8)))), 4294967295, fi8}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            v7 -= (v7 << ((uint(p1)) & 63))
            _ = fi10
        }
        p0 = (p0 << ((18446744073709551615) & 63))
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 bool, p1 bool) int8 {
    v11 := GenS1{F0: 128, F1: 1, F2: 42}
    var v12 uint = 829
    fmt.Println(v12)
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 87
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    v3 := []int{9223372036854775807, -315}
    var v4 int = 127
    v5 := make(map[int]bool)
    v4 += v4
    v13 := GenS1{F0: 4294967295, F1: v4, F2: 18446744073709551615}
    v4 -= v4
    if (!true) {
        var v14 string = "hello"
        _ = v14
    }
    v15 := GenS2{F0: v4, F1: false}
    var v16 bool = false
    for fi17 := 0; fi17 < 18; fi17++ {
        switch (fi17 ^ (fi17 + (^(fi17)))) {
        case -1:
            v3 = append(v3, -2147483648)
        case 2147483647:
            fmt.Println(v4, v4)
        default:
            var v18 int = (fi17 | (fi17 ^ (-(fi17))))
            _ = v18
        }
        _ = fi17
    }
    v4 = v4
    var v19 int32 = -1
    v19 = (v19 & v19)
    if (!((42 != 42) && false)) {
        if true {
            v20 := v2[1]
            v21 := GenS1{F0: v4, F1: v4, F2: 2}
            _ = v20
            _ = v21
        } else {
            v19 += v19
            fmt.Println(v19, v19, v16)
            if ((!(100 <= 100)) || v16) {
                var v22 int32 = 1
                var v23 string = "hello"
                for fi24 := 0; fi24 < 6; fi24++ {
                    var v25 int = (v4 / ((-2147483648) | 1))
                    _ = len(v1)
                    switch ("hello" + ("hello" + v23)) {
                    case "foo bar":
                        fmt.Println(v25)
                    case "hello":
                        v19 -= (v22 ^ -1)
                    default:
                        switch "\n\t" {
                        case "hello":
                            v26 := GenS3{F0: 65535}
                            _ = v26
                        case "a":
                            var v27 int64 = 42
                            _ = v27
                        default:
                            var v28 uint16 = 42
                            _ = v28
                        }
                    }
                    _ = fi24
                    _ = v25
                }
                _ = v22
                _ = v23
            }
        }
        v4 = -9223372036854775808
    } else {
        if v16 {
            if (368 <= 0) {
                v19 = v19
                var v29 uint32 = 1
                switch (v4 * 32768) {
                case -559:
                    fmt.Println(v29)
                case 4294967295:
                    if v16 {
                        if len(v3) > 0 {
                            v3[0] = 518
                        }
                        var v30 uint32 = v29
                        if (false && (-10 == -105)) {
                            _ = len(v5)
                        } else {
                            for fi31 := 0; fi31 < 4; fi31++ {
                                if (v16 || true) {
                                    _ = len(v2)
                                } else {
                                    v29 -= (v29 | v29)
                                }
                                for fi32 := 0; fi32 < 10; fi32++ {
                                    v4 = (fi31 & (v4 / (((v4 | (v4 % (((-(v4))) | 1)))) | 1)))
                                    _ = fi32
                                }
                                _ = fi31
                            }
                            if (genFn0(295, 216) != 4294967295) {
                                v33 := GenS3{F0: 65535}
                                _ = v33
                            }
                        }
                        _ = v30
                    } else {
                        fmt.Println(v19, v19, v29)
                    }
                }
                _ = v29
            } else {
                if len(v3) > 0 {
                    v3[0] = (v4 >> ((255) & 63))
                }
            }
            fmt.Println(v16, v19)
        } else {
            var v34 uint = genFn0(-379, 0)
            _ = v34
        }
        if (42 > -1) {
            if true {
                for fi35 := 0; fi35 < 15; fi35++ {
                    for fi36 := 0; fi36 < 11; fi36++ {
                        var v37 string = (((("hello" + "\n\t") + "hello") + "a") + "a")
                        var v38 int32 = v19
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                    var v39 string = ""
                    v4 = 2147483647
                    _ = fi35
                    _ = v39
                }
                for fi40 := 0; fi40 < 14; fi40++ {
                    v4 -= v4
                    var v41 int32 = (v19 / (((v19 | (v19 / (((v19 >> ((genFn0(9223372036854775807, 255)) & 63))) | 1)))) | 1))
                    _ = fi40
                    _ = v41
                }
                if (2 <= 459) {
                    switch -162 {
                    case 571:
                        if v16 {
                            var v42 int16 = -1
                            v16 = (-76 >= -1)
                            _ = v42
                        } else {
                            var v43 string = "hello"
                            _ = v43
                        }
                    case 128:
                        var v44 int8 = -49
                        _ = v44
                    case -1:
                        fmt.Println(v19, v4)
                    default:
                        for fi45 := 0; fi45 < 19; fi45++ {
                            v1[(fi45 % (((-(fi45))) | 1))] = "hello"
                            v19 = (-(v19))
                            _ = fi45
                        }
                    }
                }
            }
        } else {
            if v16 {
                if false {
                    v19 -= -2147483648
                    delete(v2, 1)
                } else {
                    if true {
                        var v46 int8 = genFn1((9223372036854775807 > 1), (134 <= 35))
                        v4 = (v4 * (-(v4)))
                        _ = v46
                    } else {
                        var v47 uint16 = 65535
                        v48 := GenS1{F0: (v4 - 9223372036854775807), F1: 32767, F2: 0}
                        _ = v47
                        _ = v48
                    }
                    delete(v5, v4)
                }
            }
            var v49 string = ("\n\t" + "")
            _ = v49
        }
    }
    var v50 int8 = 42
    v51 := GenS2{F0: (v4 % (((v4 / ((9223372036854775807) | 1))) | 1)), F1: (false || true)}
    switch "" {
    case "hello":
        for fi52 := 0; fi52 < 3; fi52++ {
            switch "a" {
            case "\n\t":
                var v53 uint8 = 127
                _ = v53
            case "hello":
                var v54 int32 = v19
                _ = v54
            case "":
                for fi55 := 0; fi55 < 4; fi55++ {
                    v56 := GenS4{F0: 42}
                    if len(v3) > 0 {
                        v3[0] = 32767
                    }
                    if true {
                        delete(v5, fi52)
                        switch (fi55 ^ 125) {
                        case 9223372036854775807:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case -656:
                            switch ("foo bar" + "\n\t") {
                            case "foo bar":
                                v57 := GenS5{F0: v16, F1: 9223372036854775807, F2: ((((("" + "") + "") + "foo bar") + "foo bar") + ("" + ""))}
                                _ = v57
                            case "\n\t":
                                v58 := GenS5{F0: ((1e10 == float64((0.1 * ((3.14159 - 3.14159) + -2.25)))) || (2147483647 <= (^(v19)))), F1: 0, F2: ""}
                                _ = v58
                            default:
                                v59 := GenS6{F0: ("" + "\n\t"), F1: true}
                                var v60 bool = (v16 && (!v16))
                                _ = v59
                                _ = v60
                            }
                        case 127:
                            if true {
                                if false {
                                    v16 = (v16 && v16)
                                    fmt.Println(fi55)
                                    var v61 uint32 = 4294967295
                                    _ = v61
                                } else {
                                    if v16 {
                                        var v62 int64 = 0
                                        var v63 uint64 = 0
                                        v63 -= (^(v63))
                                        _ = v62
                                        _ = v63
                                    }
                                    if v16 {
                                        switch ("\n\t" + "\n\t") {
                                        case "hello":
                                            var v64 uint16 = 921
                                            _ = v64
                                        case "a":
                                            if v16 {
                                                var v65 int32 = (v19 & v19)
                                                v66 := GenS4{F0: -733}
                                                _ = v65
                                                _ = v66
                                            } else {
                                                _ = len(v1)
                                                v16 = (4294967295 != 562)
                                            }
                                        case "":
                                            var v67 bool = (!(0 >= 1))
                                            _ = v67
                                        }
                                        var v68 string = ""
                                        _ = v68
                                    }
                                }
                                var v69 bool = v16
                                fmt.Println(fi52, v50, fi52)
                                _ = v69
                            } else {
                                v50 = (v50 & genFn1((2 < 18446744073709551615), (!v16)))
                                if (v16 || ((("" + "\n\t") > "\n\t") && v16)) {
                                    switch ("foo bar" + ((("hello" + "hello") + ("a" + "foo bar")) + "foo bar")) {
                                    case "":
                                        delete(v5, (fi52 - 2147483647))
                                    }
                                } else {
                                    var v70 uint32 = 432
                                    if len(v3) > 0 {
                                        v3[0] = fi52
                                    }
                                    _ = v70
                                }
                                _ = len(v3)
                                fmt.Println(fi52)
                                var v71 int = (v4 / (((v4 - 32768)) | 1))
                                _ = v71
                            }
                            fmt.Println(fi52, v50, v50)
                            var v72 int = (v4 | 1)
                            _ = v72
                        }
                        if v16 {
                            var v73 uint32 = 65535
                            switch (^(fi52)) {
                            case 32767:
                                var v74 uint = 914
                                _ = v74
                            case 128:
                                if false {
                                    if v16 {
                                        v19 = -661
                                        if ((-2.25 != (-2.25 / -2.25)) && (v16 || (false && (1 <= 487)))) {
                                            v16 = (816 > (v19 >> ((genFn0(634, 255)) & 63)))
                                            var v75 uint64 = 42
                                            fmt.Println(v16, fi55)
                                            _ = v75
                                        }
                                    }
                                    v16 = (v50 == v50)
                                    v76 := GenS5{F0: (!(!(753 == 2147483647))), F1: -9223372036854775808, F2: (("a" + ("" + "hello")) + "foo bar")}
                                    var v77 uint32 = v73
                                    _ = v76
                                    _ = v77
                                }
                                _ = len(v3)
                                if false {
                                    if v16 {
                                        if (v73 == (v73 ^ (^(v73)))) {
                                            if v16 {
                                                v73 += (v73 & v73)
                                            } else {
                                                v3 = append(v3, 2147483647)
                                            }
                                            v19 = v19
                                            var v78 bool = (42 > 42)
                                            v79 := GenS7{F0: 1, F1: ("\n\t" + "\n\t"), F2: (^(v73))}
                                            _ = v78
                                            _ = v79
                                        } else {
                                            if ((v73 >> ((229) & 63)) < (v73 | v73)) {
                                                switch "hello" {
                                                case "":
                                                    v80 := GenS5{F0: v16, F1: 42, F2: (("foo bar" + ("\n\t" + "foo bar")) + "hello")}
                                                    _ = v80
                                                case "\n\t":
                                                    fmt.Println(v4, v4)
                                                }
                                                v81 := GenS8{F0: 255, F1: 65535}
                                                v82 := GenS9{F0: 1, F1: v50, F2: 1}
                                                _ = v81
                                                _ = v82
                                            } else {
                                                fmt.Println(v73, v4, v16)
                                                _ = len(v5)
                                            }
                                        }
                                        v4 = fi55
                                    } else {
                                        if v16 {
                                            fmt.Println(fi55, v73, v73)
                                        } else {
                                            var v83 float64 = (float64(0.1) - (((0.1 - 0.0) - float64(1e10)) - 1e10))
                                            v50 += v50
                                            _ = v83
                                        }
                                    }
                                    _ = len(v1)
                                    switch (("" + "\n\t") + "\n\t") {
                                    case "":
                                        v5[-1] = v16
                                    case "\n\t":
                                        if true {
                                            var v84 uint32 = v73
                                            if v16 {
                                                switch v4 {
                                                case 49:
                                                    v85 := GenS10{F0: 2, F1: (v50 / (((v50 ^ -128)) | 1))}
                                                    if (v16 && v16) {
                                                        if (v16 && v16) {
                                                            switch (fi52 + (v4 & -214)) {
                                                            case 128:
                                                                v50 += v50
                                                            case -2147483648:
                                                                v19 = (v19 - (v19 % ((-1) | 1)))
                                                            case -229:
                                                                v16 = (v19 < v19)
                                                            default:
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            }
                                                            v86 := GenS7{F0: 32767, F1: "hello", F2: v73}
                                                            _ = v86
                                                        }
                                                        v19 = (v19 * v19)
                                                        var v87 string = "hello"
                                                        _ = v87
                                                    }
                                                    v88 := GenS5{F0: (v16 && v16), F1: -196, F2: ""}
                                                    _ = v85
                                                    _ = v88
                                                case -2147483648:
                                                    v73 = (^(v84))
                                                case -948:
                                                    v19 = (-(v19))
                                                }
                                            } else {
                                                v50 += (v50 % ((1) | 1))
                                                fmt.Println(v16, fi55, v4)
                                            }
                                            _ = v84
                                        }
                                    default:
                                        v16 = v16
                                    }
                                    fmt.Println(fi52)
                                } else {
                                    if (-9223372036854775808 < 0) {
                                        v4 -= (v4 | (v4 / ((fi52) | 1)))
                                        var v89 bool = (52 != 42)
                                        fmt.Println(v19)
                                        _ = v89
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                }
                            case 9223372036854775807:
                                switch 579 {
                                case 32768:
                                    switch "\n\t" {
                                    case "foo bar":
                                        v90 := GenS9{F0: 9223372036854775807, F1: (v50 - (v50 << ((0) & 63))), F2: 644}
                                        _ = v90
                                    }
                                    var v91 uint = 65535
                                    var v92 int64 = 241
                                    _ = v91
                                    _ = v92
                                case 0:
                                    v4 = 117
                                case 127:
                                    if (9223372036854775807 > 325) {
                                        var v93 float64 = float64(float32(float64(0.1)))
                                        _ = v93
                                    } else {
                                        if v16 {
                                            v94 := GenS9{F0: 549, F1: -128, F2: 481}
                                            _ = v94
                                        }
                                    }
                                }
                            }
                            if ((v73 * v73) > v73) {
                                v4 = (v4 * fi55)
                                fmt.Println(v50)
                            }
                            _ = v73
                        }
                    } else {
                        var v95 int8 = -128
                        v4 = v4
                        var v96 uint32 = 918
                        _ = v95
                        _ = v96
                    }
                    _ = fi55
                    _ = v56
                }
            }
            v4 -= 32768
            _ = fi52
        }
    case "\n\t":
        var v97 uint16 = 1
        _ = v97
    default:
        var v98 uint32 = 778
        _ = v98
    }
    switch 127 {
    case 597:
        switch 794 {
        case 42:
            for fi99 := 0; fi99 < 7; fi99++ {
                v100 := GenS11{F0: 945, F1: -1}
                v101 := GenS4{F0: 58}
                _ = fi99
                _ = v100
                _ = v101
            }
        default:
            var v102 bool = v16
            _ = v102
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
}
