package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int; F1 int16; F2 string }
type GenS4 struct { F0 uint16; F1 int8; F2 int8 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint8; F1 bool }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int64; F1 string; F2 int16 }
type GenS14 struct { F0 bool; F1 int16; F2 uint64 }
type GenS15 struct { F0 uint64; F1 uint }
type GenS16 struct { F0 uint16; F1 int8; F2 uint16 }
func genFn0() uint32 {
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := []int{-2147483648}
            var v8 float32 = 0.1
            var v9 uint8 = 146
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        var v10 float32 = (float32(((-2.25 / float64(1.5)) / (0.0 + ((((((0.0 / (0.1 * 1e10)) + 3.14159) * (1e10 / ((((0.1 - 0.0) - 1.5) * 1e10) - float64(((1.5 + ((0.1 / 1e10) - 0.1)) - 1e10))))) * (0.1 + ((-2.25 - 1e10) + 0.1))) / 0.1) * (float64(3.14159) + (3.14159 - (((1e10 * float64((0.1 / 1.0))) / 1e10) * (0.1 + 3.14159)))))))) / 1e10)
        _ = fi5
        _ = v10
    }
    var v11 int16 = -1
    if true {
        switch -876 {
        case 42:
            fmt.Println(v11)
        case 9223372036854775807:
            switch 396 {
            case 32768:
                var v12 int32 = -643
                _ = v12
            }
        case 2147483647:
            v13 := GenS1{F0: 543, F1: 42}
            _ = v13
        }
        if (v11 < 0) {
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 uint = 18446744073709551615
                if (127 <= 1) {
                    fmt.Println(v11)
                    var v16 int16 = (v11 + 32767)
                    _ = v16
                } else {
                    fmt.Println(v11, fi14)
                }
                _ = fi14
                _ = v15
            }
            v17 := []int{}
            _ = v17
        }
        switch "hello" {
        case "hello":
            v18 := []int{}
            _ = v18
        default:
            if (!(1 > 128)) {
                if false {
                    v11 = v11
                }
                v11 += (v11 << ((65535) & 63))
            } else {
                var v19 bool = (1 != 128)
                _ = v19
            }
        }
    }
    _ = v11
    return 124
}
func genFn1(p0 bool) int64 {
    if p0 {
        switch ("" + "") {
        case "a":
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 float32 = -2.25
                _ = fi20
                _ = v21
            }
        case "\n\t":
            v22 := GenS1{F0: 255, F1: -387}
            _ = v22
        default:
            switch 4294967295 {
            case 17:
                if (855 > 969) {
                    v23 := []int{722, 9223372036854775807}
                    _ = v23
                }
            case 1:
                for fi24 := 0; fi24 < 12; fi24++ {
                    fmt.Println(fi24, p0)
                    for fi25 := 0; fi25 < 19; fi25++ {
                        fmt.Println(fi25, fi25, fi24)
                        var v26 int = fi25
                        v27 := GenS2{F0: 183, F1: (-9223372036854775808 >= 929)}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    v28 := GenS3{F0: (fi24 / ((32767) | 1)), F1: -930, F2: "\n\t"}
                    _ = fi24
                    _ = v28
                }
            case -2147483648:
                if (!(2 != 231)) {
                    if ((false || p0) || (!false)) {
                        fmt.Println(p0)
                        v29 := []int{}
                        v30 := GenS1{F0: 65535, F1: 1}
                        _ = v29
                        _ = v30
                    } else {
                        var v31 uint = 709
                        fmt.Println(p0, p0)
                        if p0 {
                            v32 := []int{-1}
                            var v33 float64 = (0.1 - 1e10)
                            _ = v32
                            _ = v33
                        } else {
                            fmt.Println(p0, v31)
                        }
                        _ = v31
                    }
                    for fi34 := 0; fi34 < 15; fi34++ {
                        for fi35 := 0; fi35 < 6; fi35++ {
                            v36 := []int{(fi35 | fi34)}
                            if len(v36) > 0 {
                                v36[0] = (fi34 - 128)
                            }
                            fmt.Println(p0)
                            _ = fi35
                            _ = v36
                        }
                        _ = fi34
                    }
                }
            }
        }
        v37 := GenS2{F0: 255, F1: false}
        _ = v37
    } else {
        var v38 float32 = (3.14159 / -2.25)
        _ = v38
    }
    p0 = false
    for fi39 := 0; fi39 < 11; fi39++ {
        v40 := GenS1{F0: 1, F1: -1}
        _ = fi39
        _ = v40
    }
    _ = p0
    return 42
}

func main() {
    var v1 int32 = 1
    v2 := []int{32767, 1, 2147483647}
    var v3 uint32 = 108
    var v4 int16 = 205
    if false {
        var v41 float64 = 1.5
        if false {
            if ((100 >= 100) && ((2 <= 0) && false)) {
                if false {
                    var v42 int8 = -128
                    v41 -= (1.5 * ((1e10 - 3.14159) * v41))
                    v42 += int8(v3)
                    _ = v42
                } else {
                    var v43 bool = ((v1 + v1) <= v1)
                    if false {
                        for fi44 := 0; fi44 < 19; fi44++ {
                            v41 += v41
                            for fi45 := 0; fi45 < 6; fi45++ {
                                var v46 bool = v43
                                v47 := GenS4{F0: 521, F1: 61, F2: 100}
                                _ = fi45
                                _ = v46
                                _ = v47
                            }
                            var v48 int32 = v1
                            _ = fi44
                            _ = v48
                        }
                    } else {
                        v49 := GenS5{F0: 2, F1: 0}
                        v43 = v43
                        _ = v49
                    }
                    v4 = 0
                    _ = v43
                }
                v4 += (v4 - (v4 ^ (v4 % (((v4 & v4)) | 1))))
            }
            v4 = v4
            for fi50 := 0; fi50 < 17; fi50++ {
                var v51 uint32 = (^(v3))
                v1 += v1
                _ = fi50
                _ = v51
            }
        }
        _ = v41
    } else {
        v52 := GenS5{F0: 128, F1: 6}
        _ = v52
    }
    var v53 uint16 = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi54 := 0; fi54 < 12; fi54++ {
        if len(v2) > 0 {
            v2[0] = 32768
        }
        _ = fi54
    }
    var v55 uint16 = (^(v53))
    for fi56 := 0; fi56 < 2; fi56++ {
        if (170 < 255) {
            v55 -= (v53 * 152)
        }
        _ = fi56
    }
    var v57 int = -982
    v57 = v57
    var v58 int64 = 9223372036854775807
    var v59 int = v57
    if (!((v58 + (^(v58))) != (v58 + (v58 ^ (v58 ^ v58))))) {
        var v60 uint16 = v53
        _ = v60
    } else {
        v59 = v59
        for fi61 := 0; fi61 < 6; fi61++ {
            v62 := GenS1{F0: 0, F1: 47}
            var v63 float32 = (0.0 - (1.5 / 1.0))
            var v64 uint8 = 28
            _ = fi61
            _ = v62
            _ = v63
            _ = v64
        }
    }
    for fi65 := 0; fi65 < 1; fi65++ {
        if true {
            if false {
                if (false || true) {
                    v2 = append(v2, int(v3))
                } else {
                    if (v1 <= (-(v1))) {
                        _ = len(v2)
                        var v66 int32 = 2147483647
                        v67 := GenS6{F0: 785, F1: (v1 == -80)}
                        _ = v66
                        _ = v67
                    }
                }
                for fi68 := 0; fi68 < 9; fi68++ {
                    var v69 string = ("foo bar" + "\n\t")
                    var v70 string = ((v69 + ("" + "foo bar")) + v69)
                    switch v69 {
                    case "foo bar":
                        v4 += -152
                    case "":
                        switch int(v58) {
                        case 2147483647:
                            var v71 int8 = -117
                            _ = v71
                        case 1:
                            switch (v69 + "") {
                            case "foo bar":
                                if (v1 != -210) {
                                    if false {
                                        var v72 uint = uint(v53)
                                        var v73 int8 = 0
                                        fmt.Println(v69, fi65)
                                        v74 := GenS7{F0: (-(v1))}
                                        _ = v72
                                        _ = v73
                                        _ = v74
                                    } else {
                                        var v75 int = (fi65 | v57)
                                        _ = v75
                                    }
                                    v76 := GenS8{F0: (v3 & 2)}
                                    _ = v76
                                } else {
                                    fmt.Println(v1, v59, v69)
                                    var v77 int16 = v4
                                    fmt.Println(v4, v4, v4)
                                    _ = v77
                                }
                            case "\n\t":
                                v3 = 0
                            case "hello":
                                v78 := GenS3{F0: (v59 * v59), F1: v4, F2: "a"}
                                _ = len(v2)
                                v79 := GenS9{F0: v1}
                                _ = v78
                                _ = v79
                            }
                        default:
                            v80 := GenS10{F0: 2}
                            _ = v80
                        }
                        fmt.Println(v1)
                        var v81 bool = ((0.1 / (1e10 - (3.14159 + float64((0.0 / 1e10))))) == (float64((1e10 - 1.5)) * 3.14159))
                        _ = v81
                    case "hello":
                        v82 := GenS11{F0: 2, F1: false}
                        _ = v82
                    }
                    _ = fi68
                    _ = v69
                    _ = v70
                }
                v83 := GenS5{F0: 102, F1: 127}
                _ = v83
            }
            fmt.Println(v58, fi65)
            var v84 bool = false
            _ = v84
        } else {
            if ((!((!true) || (40 >= 629))) || (false && false)) {
                if true {
                    if false {
                        for fi85 := 0; fi85 < 17; fi85++ {
                            v53 -= (v55 / ((v55) | 1))
                            v55 = (^(v53))
                            switch "a" {
                            case "foo bar":
                                v4 = (^(v4))
                            case "hello":
                                v2 = append(v2, v59)
                            }
                            v86 := GenS12{F0: (v57 - (v59 ^ 9223372036854775807))}
                            _ = fi85
                            _ = v86
                        }
                        switch "" {
                        case "a":
                            var v87 int64 = genFn1(false)
                            _ = v87
                        }
                        var v88 uint64 = 2
                        _ = v88
                    }
                    v2 = append(v2, (v59 | fi65))
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch ("foo bar" + ("hello" + "hello")) {
                    case "foo bar":
                        if (!(v1 > (v1 << ((65535) & 63)))) {
                            var v89 int16 = v4
                            _ = v89
                        }
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "\n\t":
                        _ = len(v2)
                    default:
                        fmt.Println(v59)
                    }
                }
                v90 := GenS13{F0: (v58 / (((^(v58))) | 1)), F1: "", F2: -36}
                if len(v2) > 0 {
                    v2[0] = -1
                }
                _ = v90
            }
            for fi91 := 0; fi91 < 13; fi91++ {
                switch "a" {
                case "":
                    v92 := GenS14{F0: (120 < 79), F1: v4, F2: 2}
                    _ = v92
                case "a":
                    v55 = 2
                default:
                    v3 = v3
                }
                var v93 uint8 = 0
                v94 := GenS15{F0: 42, F1: 773}
                _ = fi91
                _ = v93
                _ = v94
            }
            switch (fi65 >> ((9) & 63)) {
            case 0:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 127:
                for fi95 := 0; fi95 < 7; fi95++ {
                    v1 += 896
                    var v96 bool = (1.5 != (0.1 + 3.14159))
                    var v97 int = (v59 >> ((2) & 63))
                    _ = fi95
                    _ = v96
                    _ = v97
                }
            }
        }
        v98 := GenS16{F0: (^(v53)), F1: 21, F2: (v53 / ((837) | 1))}
        var v99 string = "foo bar"
        _ = fi65
        _ = v98
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v55
    _ = v57
    _ = v58
    _ = v59
}
