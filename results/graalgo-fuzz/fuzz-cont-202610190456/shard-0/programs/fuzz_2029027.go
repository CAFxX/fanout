package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int16 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 bool }
type GenS4 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS7 struct { F0 uint64; F1 int32; F2 uint }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint; F1 uint8 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint64; F1 uint8; F2 int }
func genFn0() int16 {
    for fi6 := 0; fi6 < 17; fi6++ {
        switch 9223372036854775807 {
        case 126:
            switch "foo bar" {
            case "":
                v7 := GenS1{F0: 1, F1: 1, F2: 348}
                _ = v7
            case "hello":
                for fi8 := 0; fi8 < 6; fi8++ {
                    switch (("foo bar" + "\n\t") + "foo bar") {
                    case "\n\t":
                        var v9 uint = 72
                        _ = v9
                    }
                    v10 := []int{(fi6 * fi6), (fi6 >> ((287) & 63)), fi6, (fi8 * (fi8 | (^(fi6))))}
                    if len(v10) > 0 {
                        v10[0] = fi6
                    }
                    _ = fi8
                    _ = v10
                }
            case "a":
                var v11 float64 = float64(float32(3.14159))
                _ = v11
            default:
                v12 := GenS2{F0: (1000 < -32768), F1: 1}
                _ = v12
            }
        case 42:
            var v13 int16 = 42
            _ = v13
        }
        if (false || false) {
            var v14 uint = 65535
            _ = v14
        } else {
            var v15 uint = 4294967295
            _ = v15
        }
        var v16 float32 = (1e10 / 1.0)
        _ = fi6
        _ = v16
    }
    v17 := []int{848, -876, 1, 128}
    _ = v17
    return 128
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]int)
    v18 := GenS1{F0: 42, F1: 100, F2: 42}
    for fi19 := 0; fi19 < 7; fi19++ {
        switch 128 {
        case 1:
            fmt.Println(v1)
        }
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 uint16 = 1
            v22 := GenS3{F0: 4294967295, F1: v1, F2: (false || false)}
            v1 -= v1
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    v1 += v1
    for fi23 := 0; fi23 < 12; fi23++ {
        if (("" + ("" + ((("hello" + "foo bar") + "") + ("\n\t" + ("hello" + ("foo bar" + "foo bar")))))) > "a") {
            fmt.Println(v1, fi23)
            fmt.Println(v1)
        }
        var v24 uint = 18446744073709551615
        for fi25 := 0; fi25 < 12; fi25++ {
            switch "foo bar" {
            case "hello":
                if true {
                    v26 := GenS4{F0: 1, F1: true, F2: 65535}
                    _ = v26
                }
            case "":
                if (2 != 42) {
                    if false {
                        if false {
                            v24 += v24
                            var v27 int16 = int16(fi23)
                            var v28 int32 = 2147483647
                            _ = v27
                            _ = v28
                        } else {
                            v29 := GenS3{F0: 2147483647, F1: uint16(v24), F2: false}
                            var v30 int16 = -288
                            v30 -= (v30 << (((^(v24))) & 63))
                            var v31 bool = (!((("foo bar" + ("\n\t" + "a")) + "hello") == ("a" + ("foo bar" + (("hello" + ("\n\t" + (("hello" + ("" + "")) + "\n\t"))) + "\n\t")))))
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    } else {
                        v32 := []int{}
                        if (1 == 177) {
                            v33 := GenS5{F0: 42, F1: 567}
                            var v34 bool = ((255 >= (^(v1))) || false)
                            var v35 int = fi25
                            _ = v33
                            _ = v34
                            _ = v35
                        } else {
                            var v36 uint16 = (v1 % ((32767) | 1))
                            v37 := GenS6{F0: 934, F1: 2, F2: uint8(v1)}
                            _ = v36
                            _ = v37
                        }
                        _ = v32
                    }
                } else {
                    fmt.Println(fi25, v24, v24)
                    if ((1.5 != 1e10) || true) {
                        v1 += 0
                    }
                    v24 = v24
                    v38 := GenS7{F0: 177, F1: 0, F2: v24}
                    if (false && (0.1 > 0.0)) {
                        v39 := GenS8{F0: 32767}
                        _ = v39
                    }
                    var v40 float32 = (0.1 * 0.0)
                    _ = v38
                    _ = v40
                }
            }
            v41 := GenS2{F0: true, F1: 42}
            _ = fi25
            _ = v41
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v1)
    var v42 int32 = 0
    var v43 uint32 = 2147483648
    for fi44 := 0; fi44 < 15; fi44++ {
        v1 -= (v1 & (^(v1)))
        switch fi44 {
        case 32767:
            for fi45 := 0; fi45 < 11; fi45++ {
                var v46 bool = true
                var v47 uint32 = (v43 / (((v43 * (v43 ^ (v43 ^ v43)))) | 1))
                _ = fi45
                _ = v46
                _ = v47
            }
        }
        if true {
            if true {
                switch "foo bar" {
                case "hello":
                    v43 += (^(v43))
                }
                v48 := []int{(-(fi44))}
                _ = v48
            }
            var v49 uint16 = (v1 >> ((255) & 63))
            delete(v3, (fi44 / ((-351) | 1)))
            _ = v49
        }
        _ = fi44
    }
    v1 = v1
    v50 := []int{9223372036854775807}
    var v51 int8 = 51
    delete(v3, -558)
    if ("" == ("foo bar" + "\n\t")) {
        var v52 int = -24
        _ = v52
    } else {
        var v53 uint16 = 710
        for fi54 := 0; fi54 < 1; fi54++ {
            v55 := GenS8{F0: v53}
            _ = fi54
            _ = v55
        }
        if (!true) {
            fmt.Println(v43, v43, v51)
            v56 := GenS9{F0: 1}
            if len(v50) >= 2 {
                v50 = v50[1:len(v50)]
            }
            _ = v56
        }
        _ = v53
    }
    fmt.Println(v42, v43, v51)
    if false {
        fmt.Println(v51, v42, v1)
        if (("foo bar" + "foo bar") != "a") {
            if (42 > 2) {
                if true {
                    _ = len(v2)
                    for fi57 := 0; fi57 < 13; fi57++ {
                        switch (fi57 | -1) {
                        case 42:
                            for fi58 := 0; fi58 < 4; fi58++ {
                                fmt.Println(v42, v43)
                                fmt.Println(v51, v43, v42)
                                _ = fi58
                            }
                        case -225:
                            v5[v43] = fi57
                        case -1:
                            if ((0 >= int64(v42)) && true) {
                                var v59 string = "hello"
                                v1 = (v1 ^ 32767)
                                fmt.Println(fi57)
                                _ = v59
                            } else {
                                var v60 int = (fi57 / ((fi57) | 1))
                                _ = len(v50)
                                switch v60 {
                                case 1:
                                    var v61 int = -982
                                    _ = v61
                                }
                                _ = v60
                            }
                        }
                        if (true && (2 > 65535)) {
                            if true {
                                if len(v50) >= 2 {
                                    v50 = v50[1:len(v50)]
                                }
                                fmt.Println(v42, fi57, fi57)
                            }
                            for fi62 := 0; fi62 < 12; fi62++ {
                                fmt.Println(fi62, v42)
                                v3[fi57] = (("\n\t" + "hello") + "\n\t")
                                var v63 float64 = 1e10
                                var v64 int16 = 0
                                v65 := GenS1{F0: v51, F1: 127, F2: v64}
                                var v66 uint64 = 1
                                v66 = v66
                                _ = fi62
                                _ = v63
                                _ = v64
                                _ = v65
                                _ = v66
                            }
                        } else {
                            if true {
                                for fi67 := 0; fi67 < 15; fi67++ {
                                    v42 -= (v42 - (v42 ^ v42))
                                    v4[v43] = (v43 << ((4294967295) & 63))
                                    if len(v50) > 0 {
                                        v50[0] = fi67
                                    }
                                    _ = fi67
                                }
                                var v68 int8 = v51
                                _ = v68
                            }
                        }
                        switch "foo bar" {
                        case "\n\t":
                            switch "\n\t" {
                            case "foo bar":
                                if (42 > 255) {
                                    for fi69 := 0; fi69 < 8; fi69++ {
                                        if (!(!true)) {
                                            v1 += v1
                                        }
                                        var v70 string = "hello"
                                        _ = fi69
                                        _ = v70
                                    }
                                    fmt.Println(v42)
                                }
                            }
                        default:
                            switch 4294967295 {
                            case -981:
                                _ = len(v50)
                            }
                        }
                        _ = fi57
                    }
                } else {
                    _ = len(v50)
                    var v71 int64 = 42
                    var v72 uint64 = 2
                    _ = v71
                    _ = v72
                }
            }
        } else {
            v51 = -107
            v43 += v43
        }
        v51 -= (v51 * v51)
    } else {
        if true {
            if true {
                for fi73 := 0; fi73 < 13; fi73++ {
                    fmt.Println(v51, v1, v1)
                    var v74 uint64 = 2
                    var v75 uint64 = (v74 * v74)
                    _ = fi73
                    _ = v74
                    _ = v75
                }
            } else {
                v43 += 973
            }
        } else {
            switch "" {
            case "":
                v76 := v2[v43]
                _ = v76
            default:
                for fi77 := 0; fi77 < 7; fi77++ {
                    for fi78 := 0; fi78 < 6; fi78++ {
                        _ = len(v50)
                        switch "\n\t" {
                        case "a":
                            if true {
                                v42 = -464
                                switch "hello" {
                                case "hello":
                                    v1 = 156
                                case "\n\t":
                                    v42 -= (v42 / (((v42 & (v42 * (-(v42))))) | 1))
                                case "foo bar":
                                    switch (("foo bar" + "hello") + ("\n\t" + "")) {
                                    case "hello":
                                        var v79 bool = (v51 != (v51 >> ((2) & 63)))
                                        _ = v79
                                    case "\n\t":
                                        v51 = v51
                                    }
                                }
                                var v80 int32 = v42
                                v51 += v51
                                if true {
                                    fmt.Println(v1)
                                    v43 = v43
                                    delete(v5, v43)
                                    delete(v4, v43)
                                    v51 = v51
                                } else {
                                    v81 := GenS10{F0: 178, F1: 255}
                                    _ = v81
                                }
                                _ = v80
                            }
                        case "foo bar":
                            var v82 string = "foo bar"
                            _ = v82
                        case "\n\t":
                            if (2 > 255) {
                                if false {
                                    var v83 bool = ((((-2.25 + 1e10) * 3.14159) != -2.25) && false)
                                    v84 := GenS11{F0: (v43 + v43)}
                                    _ = len(v50)
                                    _ = v83
                                    _ = v84
                                } else {
                                    if true {
                                        v85 := GenS12{F0: 1}
                                        if true {
                                            switch 228 {
                                            case 32768:
                                                v42 += v42
                                            default:
                                                v5[v43] = -1
                                            }
                                            if (!true) {
                                                var v86 uint32 = 2147483648
                                                fmt.Println(fi78)
                                                _ = v86
                                            } else {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v43 = (v43 - (v43 >> ((18446744073709551615) & 63)))
                                                case "a":
                                                    var v87 int = 502
                                                    _ = v87
                                                case "hello":
                                                    var v88 float64 = 1e10
                                                    _ = v88
                                                }
                                                var v89 int8 = (^(v51))
                                                _ = v89
                                            }
                                            fmt.Println(v1)
                                        } else {
                                            var v90 uint64 = 4294967295
                                            if false {
                                                v42 = (^(v42))
                                                var v91 int32 = v42
                                                v91 -= -2147483648
                                                _ = v91
                                            } else {
                                                switch "foo bar" {
                                                case "":
                                                    if (fi78 > (fi77 + (^(fi78)))) {
                                                        v42 = v42
                                                    } else {
                                                        fmt.Println(v42, v1)
                                                        var v92 uint32 = v43
                                                        var v93 int8 = v51
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                case "\n\t":
                                                    var v94 uint64 = uint64(v43)
                                                    _ = v94
                                                case "foo bar":
                                                    if true {
                                                        v95 := GenS13{F0: (^(v90)), F1: 226, F2: int(v42)}
                                                        v90 = v90
                                                        _ = v95
                                                    } else {
                                                        if true {
                                                            if false {
                                                                delete(v2, (v43 % ((v43) | 1)))
                                                            } else {
                                                                v90 += (v90 + v90)
                                                                v96 := GenS7{F0: v90, F1: (v42 ^ (v42 << ((4294967295) & 63))), F2: 1}
                                                                fmt.Println(v90, v42)
                                                                _ = v96
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            var v97 uint16 = (v1 % ((2) | 1))
                                            _ = v90
                                            _ = v97
                                        }
                                        _ = v85
                                    }
                                    if len(v50) > 0 {
                                        v50[0] = -1
                                    }
                                }
                            }
                        default:
                            var v98 float64 = (0.1 * 3.14159)
                            _ = v98
                        }
                        _ = fi78
                    }
                    var v99 int = -1
                    _ = fi77
                    _ = v99
                }
            }
        }
        v51 -= 19
        _ = len(v3)
    }
    var v100 uint16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v100
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v50
    _ = v51
}
