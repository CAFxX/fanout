package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int32 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 int8; F1 int64 }
type GenS9 struct { F0 int32; F1 int8 }
type GenS10 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 string; F1 int32; F2 string }
type GenS13 struct { F0 int64; F1 string }
func genFn0(p0 uint, p1 string, p2 int16) int {
    switch p1 {
    case "hello":
        if false {
            fmt.Println(p2, p1, p0)
        } else {
            var v4 int32 = 1
            var v5 uint32 = 2147483647
            for fi6 := 0; fi6 < 16; fi6++ {
                if (65535 > (v5 / ((v5) | 1))) {
                    v5 -= 65535
                }
                _ = fi6
            }
            _ = v4
            _ = v5
        }
    case "\n\t":
        v7 := []int{128}
        _ = v7
    case "a":
        v8 := GenS1{F0: 38, F1: -677, F2: 42}
        _ = v8
    default:
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint8 = 2
            fmt.Println(p2, v10, p0)
            v10 = (v10 | (v10 & v10))
            _ = fi9
            _ = v10
        }
    }
    p0 = (p0 - (p0 % ((2) | 1)))
    _ = p0
    _ = p1
    _ = p2
    return 243
}
func genFn1(p0 int16, p1 int) uint16 {
    if true {
        p1 = p1
        for fi11 := 0; fi11 < 5; fi11++ {
            fmt.Println(p0, fi11, p1)
            _ = fi11
        }
    } else {
        if false {
            v12 := []int{-9223372036854775808, (-(p1))}
            p1 = p1
            _ = v12
        } else {
            for fi13 := 0; fi13 < 5; fi13++ {
                p1 = genFn0(42, "", p0)
                var v14 bool = (!(!true))
                _ = fi13
                _ = v14
            }
        }
        p0 = (p0 | p0)
    }
    v15 := GenS1{F0: 128, F1: (p0 + p0), F2: -2147483648}
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 uint64 = 787
            var v19 int64 = -9223372036854775808
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 float64 = 1e10
        _ = fi16
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v15
    return 932
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = float32(float64((3.14159 + 0.1)))
    var v3 bool = false
    var v21 string = "\n\t"
    var v22 uint64 = 0
    var v23 uint64 = (v22 + v22)
    var v24 uint16 = 590
    v22 = (v22 << ((18446744073709551615) & 63))
    if (-571 != int64(v23)) {
        v25 := GenS2{F0: v24, F1: 42}
        if false {
            v26 := []int{genFn0(836, v21, 1), 493, -2147483648}
            _ = v26
        } else {
            v3 = v3
            var v27 float64 = 0.0
            _ = v27
        }
        _ = v25
    }
    v28 := GenS3{F0: v22, F1: 4294967295}
    var v29 uint8 = 127
    v23 -= 66
    for fi30 := 0; fi30 < 5; fi30++ {
        if false {
            var v31 string = v21
            v32 := GenS4{F0: (v29 | 0), F1: v29, F2: (v23 << ((808) & 63))}
            _ = v31
            _ = v32
        } else {
            var v33 int16 = -1
            if v3 {
                var v34 int = fi30
                v2 += (0.0 * (v2 * v2))
                _ = v34
            }
            _ = v33
        }
        _ = fi30
    }
    v35 := []int{32767, genFn0(65535, v21, 1000)}
    v29 = 127
    v36 := GenS5{F0: v24, F1: v23}
    for fi37 := 0; fi37 < 19; fi37++ {
        if (0.1 == -2.25) {
            for fi38 := 0; fi38 < 8; fi38++ {
                switch fi37 {
                case -1:
                    fmt.Println(v24, v2)
                }
                fmt.Println(fi38, v22)
                delete(v1, 2147483647)
                var v39 uint16 = v24
                var v40 uint8 = (v29 & v29)
                if ((v2 + ((-2.25 - (v2 * (3.14159 - ((0.0 + float32((1e10 * ((0.1 + 3.14159) / 1.0)))) * (3.14159 / float32(3.14159)))))) * (v2 / (0.0 + ((v2 - v2) * (1.5 / -2.25)))))) < 1.5) {
                    var v41 int32 = 1
                    if v3 {
                        v39 -= 32767
                    }
                    _ = v41
                }
                switch ((v21 + v21) + "foo bar") {
                case "a":
                    v2 -= 1e10
                }
                _ = fi38
                _ = v39
                _ = v40
            }
            var v42 uint64 = v23
            switch v21 {
            case "":
                if (v3 && ((0 < 57) && v3)) {
                    var v43 string = "foo bar"
                    v24 = v24
                    v29 = v29
                    _ = v43
                }
            case "hello":
                for fi44 := 0; fi44 < 6; fi44++ {
                    if ((1e10 - v2) < (v2 / v2)) {
                        if v3 {
                            v29 = v29
                        } else {
                            switch ("a" + v21) {
                            case "hello":
                                var v45 string = ""
                                _ = v45
                            case "foo bar":
                                var v46 int8 = -1
                                _ = v46
                            }
                        }
                        if (42 >= 65) {
                            var v47 int64 = 42
                            var v48 float32 = float32(((1e10 / ((1e10 / (0.1 + (0.1 - 0.1))) * (1e10 - 0.0))) - 0.1))
                            _ = v47
                            _ = v48
                        }
                    } else {
                        if len(v35) > 0 {
                            v35[0] = 2147483647
                        }
                        if (0.1 != 0.0) {
                            switch "hello" {
                            case "\n\t":
                                v49 := GenS1{F0: v29, F1: int16(v24), F2: -1}
                                _ = v49
                            case "a":
                                v50 := GenS5{F0: v24, F1: v22}
                                _ = len(v35)
                                _ = v50
                            case "hello":
                                fmt.Println(v42)
                            default:
                                if false {
                                    if (!v3) {
                                        if true {
                                            var v51 int8 = 74
                                            var v52 int16 = 42
                                            v53 := GenS5{F0: uint16(v42), F1: 0}
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                        }
                                        if v3 {
                                            v54 := GenS1{F0: v29, F1: 32767, F2: 1}
                                            v55 := GenS6{F0: (v24 << ((1) & 63)), F1: v3}
                                            _ = v54
                                            _ = v55
                                        } else {
                                            if (((v24 / ((v24) | 1)) > 4) && v3) {
                                                if len(v35) > 0 {
                                                    v35[0] = (^(fi44))
                                                }
                                            } else {
                                                if false {
                                                    v3 = (2147483647 > 2147483647)
                                                    fmt.Println(fi44)
                                                }
                                                v56 := GenS2{F0: (v24 * 32767), F1: -9223372036854775808}
                                                _ = len(v1)
                                                _ = v56
                                            }
                                            fmt.Println(v42, v22, v3)
                                        }
                                    } else {
                                        v2 = v2
                                        var v57 float32 = float32(0.0)
                                        v21 = "foo bar"
                                        _ = len(v1)
                                        _ = v57
                                    }
                                    fmt.Println(v22, fi37)
                                    var v58 float32 = ((v2 / 1.5) - (((1.5 - v2) / (3.14159 * v2)) - float32(3.14159)))
                                    var v59 bool = (int32(v23) > 0)
                                    v60 := GenS7{F0: v21, F1: 2}
                                    fmt.Println(v22, v42, fi44)
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                            }
                            v2 = (v2 + (((1e10 / 3.14159) + float32(0.0)) - v2))
                            v61 := GenS8{F0: 100, F1: -9223372036854775808}
                            _ = v61
                        } else {
                            _ = len(v1)
                            fmt.Println(v29, v2, v2)
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                        }
                        v21 = v21
                        switch 32767 {
                        case 1:
                            var v62 uint16 = v24
                            _ = v62
                        default:
                            v29 += uint8(v24)
                        }
                        switch fi37 {
                        case 128:
                            v63 := GenS9{F0: int32(fi44), F1: 1}
                            _ = v63
                        case 1:
                            if (!v3) {
                                v3 = false
                                fmt.Println(v22, fi44, v23)
                                v29 += v29
                            } else {
                                var v64 uint64 = v23
                                if (!false) {
                                    if v3 {
                                        v24 = genFn1(1, (fi37 + fi44))
                                        var v65 int = 9223372036854775807
                                        if ((9223372036854775807 > 1) || v3) {
                                            v64 += (v64 << ((552) & 63))
                                            v3 = (0.1 < 1.5)
                                            v2 = (1e10 - ((0.1 / -2.25) + 0.0))
                                        } else {
                                            if (v2 == (v2 + v2)) {
                                                var v66 uint16 = v24
                                                _ = v66
                                            }
                                        }
                                        var v67 float32 = float32(1.5)
                                        v2 = (3.14159 / v2)
                                        var v68 uint32 = 42
                                        if true {
                                            v3 = v3
                                            if (1.5 != 1.5) {
                                                v69 := v1[1]
                                                var v70 int64 = 0
                                                _ = v69
                                                _ = v70
                                            } else {
                                                v65 = (fi37 >> ((255) & 63))
                                                var v71 float64 = float64(v2)
                                                _ = v71
                                            }
                                            v72 := GenS10{F0: (v29 | (v29 & v29)), F1: uint32(v29), F2: int16(v24)}
                                            _ = v72
                                        }
                                        v73 := GenS7{F0: v21, F1: (v68 - 871)}
                                        var v74 uint32 = (v68 + 2147483647)
                                        _ = v65
                                        _ = v67
                                        _ = v68
                                        _ = v73
                                        _ = v74
                                    } else {
                                        var v75 uint16 = v24
                                        if len(v35) > 0 {
                                            v35[0] = fi44
                                        }
                                        fmt.Println(v75)
                                        var v76 uint32 = 761
                                        if len(v35) >= 2 {
                                            v35 = v35[1:len(v35)]
                                        }
                                        _ = v75
                                        _ = v76
                                    }
                                    v77 := v1[0]
                                    v78 := GenS2{F0: genFn1(841, (fi44 | (fi37 * (^(fi37))))), F1: 478}
                                    _ = v77
                                    _ = v78
                                } else {
                                    switch (^(fi44)) {
                                    case 0:
                                        v21 = v21
                                    case 636:
                                        if (!(!v3)) {
                                            var v79 int64 = 9223372036854775807
                                            v64 -= (^(v22))
                                            if (!(v29 <= 1)) {
                                                switch v21 {
                                                case "hello":
                                                    v80 := GenS8{F0: -1, F1: -941}
                                                    _ = v80
                                                case "":
                                                    v42 += (v64 >> ((2) & 63))
                                                case "foo bar":
                                                    if (-2.25 > 0.1) {
                                                        var v81 int32 = 0
                                                        if (!(1 >= 1)) {
                                                            v24 = v24
                                                        }
                                                        v23 = v23
                                                        _ = v81
                                                    }
                                                }
                                            } else {
                                                v82 := GenS8{F0: 127, F1: -533}
                                                _ = v82
                                            }
                                            _ = v79
                                        } else {
                                            fmt.Println(v2, v22)
                                            delete(v1, 702)
                                        }
                                    default:
                                        switch fi44 {
                                        case 127:
                                            if (0 == 1) {
                                                var v83 float32 = (0.1 - 0.1)
                                                v84 := GenS10{F0: v29, F1: uint32(v23), F2: 1}
                                                v35 = append(v35, (fi37 - genFn0(uint(v24), "a", 1000)))
                                                switch (fi37 * (fi37 & (fi44 * (fi44 / (((^(fi44))) | 1))))) {
                                                case -1:
                                                    v85 := GenS4{F0: 255, F1: (v29 - 1), F2: v42}
                                                    _ = v85
                                                default:
                                                    if len(v35) >= 2 {
                                                        v35 = v35[1:len(v35)]
                                                    }
                                                    v86 := GenS1{F0: (v29 * (v29 << ((18446744073709551615) & 63))), F1: 42, F2: -1}
                                                    _ = v86
                                                }
                                                if (!false) {
                                                    if v3 {
                                                        if (!(1e10 != -2.25)) {
                                                            fmt.Println(v21, v42)
                                                            if false {
                                                                var v87 uint = 4294967295
                                                                v1[455] = 65535
                                                                _ = v87
                                                            } else {
                                                                switch v21 {
                                                                case "a":
                                                                    v29 = v29
                                                                default:
                                                                    v42 += 1
                                                                }
                                                                fmt.Println(v23, fi37)
                                                            }
                                                        } else {
                                                            _ = len(v1)
                                                        }
                                                        var v88 int = (fi37 % ((0) | 1))
                                                        if (v3 || true) {
                                                            fmt.Println(v24)
                                                            var v89 int = genFn0(65535, ((v21 + ((v21 + v21) + "\n\t")) + v21), -730)
                                                            _ = v89
                                                        }
                                                        _ = v88
                                                    }
                                                    switch fi44 {
                                                    case -582:
                                                        if (true && true) {
                                                            fmt.Println(v23, fi44, fi44)
                                                            var v90 uint32 = 134
                                                            v91 := GenS11{F0: v90}
                                                            var v92 uint8 = v29
                                                            v3 = false
                                                            _ = v90
                                                            _ = v91
                                                            _ = v92
                                                        } else {
                                                            var v93 int32 = int32(fi37)
                                                            var v94 int = int(v93)
                                                            switch v94 {
                                                            case -9223372036854775808:
                                                                v22 += (v64 | v64)
                                                            }
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                    case 32768:
                                                        if (533 == 401) {
                                                            v95 := GenS12{F0: "", F1: -235, F2: v21}
                                                            if ((fi44 | (fi44 << ((uint(v29)) & 63))) >= -181) {
                                                                switch v21 {
                                                                case "":
                                                                    var v96 int16 = 0
                                                                    _ = v96
                                                                default:
                                                                    var v97 uint32 = 700
                                                                    _ = v97
                                                                }
                                                                v98 := GenS13{F0: -1, F1: "a"}
                                                                fmt.Println(v2, v2, v42)
                                                                _ = v98
                                                            } else {
                                                                fmt.Println(v64, v3)
                                                                v21 = v21
                                                                var v99 uint16 = v24
                                                                v2 += (v2 + (float32(float64((v83 / float32(((0.1 - 3.14159) + 3.14159))))) - v83))
                                                                _ = v99
                                                            }
                                                            _ = v95
                                                        }
                                                    }
                                                    v22 = 1
                                                    if v3 {
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v100 uint32 = 4294967295
                                                            _ = v100
                                                        case "hello":
                                                            v24 += 459
                                                        case "\n\t":
                                                            v64 = v64
                                                        default:
                                                            v64 += (v23 & v64)
                                                        }
                                                        if false {
                                                            var v101 int16 = 1
                                                            if len(v35) > 0 {
                                                                v35[0] = (fi44 & (fi37 << ((65535) & 63)))
                                                            }
                                                            if len(v35) > 0 {
                                                                v35[0] = -152
                                                            }
                                                            _ = v101
                                                        } else {
                                                            switch v21 {
                                                            case "hello":
                                                                if v3 {
                                                                    var v102 int64 = -9223372036854775808
                                                                    _ = v102
                                                                }
                                                            }
                                                            var v103 uint16 = genFn1(694, (-(fi44)))
                                                            _ = v103
                                                        }
                                                    }
                                                    v1[600] = 2
                                                    v104 := GenS10{F0: (^(v29)), F1: 0, F2: 42}
                                                    if len(v35) >= 2 {
                                                        v35 = v35[1:len(v35)]
                                                    }
                                                    v105 := GenS1{F0: (v29 & (v29 << ((4294967295) & 63))), F1: -1, F2: 1}
                                                    _ = v104
                                                    _ = v105
                                                } else {
                                                    v23 = (^(v42))
                                                    v24 -= v24
                                                    var v106 int = genFn0(67, "", 32767)
                                                    switch v21 {
                                                    case "\n\t":
                                                        var v107 uint32 = 246
                                                        _ = v107
                                                    }
                                                    _ = v106
                                                }
                                                _ = v83
                                                _ = v84
                                            } else {
                                                _ = len(v35)
                                                v108 := GenS1{F0: v29, F1: 84, F2: 42}
                                                _ = v108
                                            }
                                        case 2147483647:
                                            var v109 int = 0
                                            _ = v109
                                        default:
                                            switch (fi44 & (fi44 % ((fi44) | 1))) {
                                            case -979:
                                                v64 = (v22 & v23)
                                            case 32767:
                                                if len(v35) >= 2 {
                                                    v35 = v35[1:len(v35)]
                                                }
                                                var v110 int16 = 0
                                                _ = v110
                                            case 4294967295:
                                                if len(v35) > 0 {
                                                    v35[0] = fi37
                                                }
                                                v111 := GenS9{F0: -1, F1: -128}
                                                fmt.Println(v29)
                                                _ = v111
                                            }
                                        }
                                    }
                                    var v112 float64 = 1e10
                                    v2 = (3.14159 - 1.5)
                                    _ = v112
                                }
                                v2 += float32(1e10)
                                _ = v64
                            }
                        case -597:
                            fmt.Println(v2)
                        }
                        switch fi37 {
                        case 0:
                            switch "hello" {
                            case "":
                                v113 := GenS8{F0: 42, F1: -9223372036854775808}
                                _ = v113
                            case "hello":
                                v3 = (!(1.5 < v2))
                            case "foo bar":
                                v24 = (v24 >> ((42) & 63))
                            }
                        case -2147483648:
                            v24 += (v24 + (v24 ^ v24))
                        }
                    }
                    if len(v35) > 0 {
                        v35[0] = (^(fi37))
                    }
                    if true {
                        switch ((v21 + v21) + "") {
                        case "hello":
                            v2 += 0.1
                        }
                    }
                    v23 = (v23 | v42)
                    _ = fi44
                }
            }
            _ = v42
        } else {
            if (((!true) && v3) && true) {
                v3 = (!(0.0 < float64(float32(0.1))))
            } else {
                fmt.Println(v21, v3)
            }
        }
        _ = fi37
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
}
