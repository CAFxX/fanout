package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }
type GenS2 struct { F0 int64; F1 uint32; F2 int64 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint; F1 uint8; F2 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 uint64; F1 int16; F2 uint64 }
func genFn0(p0 uint64) int16 {
    var v5 int = 1
    var v6 uint16 = 714
    v7 := []int{(v5 | v5)}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -426
}
func genFn1(p0 string, p1 uint64, p2 uint8) int16 {
    if (false && false) {
        for fi8 := 0; fi8 < 14; fi8++ {
            switch (fi8 / ((fi8) | 1)) {
            case 42:
                v9 := []int{fi8, fi8, (fi8 / (((fi8 * fi8)) | 1))}
                _ = v9
            case -1:
                v10 := []int{fi8, fi8, fi8}
                _ = v10
            case 4294967295:
                var v11 uint16 = 255
                _ = v11
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                if (1 < 1) {
                    v13 := GenS1{F0: int8(p2), F1: 0, F2: 2147483647}
                    _ = v13
                }
                _ = fi12
            }
            switch "hello" {
            case "foo bar":
                v14 := []int{(fi8 % ((fi8) | 1))}
                _ = v14
            default:
                var v15 bool = true
                _ = v15
            }
            _ = fi8
        }
        var v16 uint32 = 2
        fmt.Println(p2, v16, p2)
        _ = v16
    } else {
        if (42 == (^(p1))) {
            fmt.Println(p1)
            for fi17 := 0; fi17 < 13; fi17++ {
                var v18 float32 = 3.14159
                v19 := []int{(-(fi17)), fi17, fi17}
                _ = len(v19)
                _ = fi17
                _ = v18
                _ = v19
            }
        } else {
            switch 0 {
            case 32768:
                if true {
                    v20 := GenS2{F0: 397, F1: 1, F2: -1}
                    v21 := []int{}
                    for fi22 := 0; fi22 < 1; fi22++ {
                        var v23 uint32 = 2
                        var v24 int32 = int32(fi22)
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = v20
                    _ = v21
                }
            case 241:
                p2 = p2
            case 0:
                var v25 int64 = -9223372036854775808
                _ = v25
            }
        }
        p1 -= (p1 - p1)
    }
    var v26 uint64 = p1
    var v27 uint64 = p1
    if true {
        var v28 uint16 = 0
        _ = v28
    } else {
        var v29 int32 = 324
        fmt.Println(p2, v29)
        for fi30 := 0; fi30 < 10; fi30++ {
            v26 -= v26
            var v31 bool = false
            var v32 uint = 4294967295
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = v29
    }
    v27 += (p1 + v27)
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return 42
}

func main() {
    var v1 uint32 = 947
    v2 := []string{("a" + "a"), "foo bar", (("\n\t" + "") + "")}
    var v3 uint64 = 977
    var v4 uint8 = 128
    var v33 uint = 65535
    if false {
        var v34 float64 = 0.1
        switch -298 {
        case 2147483647:
            v35 := GenS1{F0: 0, F1: (v1 / ((uint32(v3)) | 1)), F2: 42}
            _ = v35
        case 4294967295:
            v36 := GenS3{F0: (3.14159 <= (3.14159 * (((0.1 * (((0.1 - 3.14159) * 0.1) / -2.25)) * 1e10) + 0.0))), F1: v33}
            _ = v36
        }
        if true {
            v37 := GenS4{F0: 4294967295, F1: 93, F2: -580}
            switch ("hello" + ("a" + ("hello" + ("\n\t" + "foo bar")))) {
            case "a":
                var v38 uint64 = 255
                _ = v38
            case "":
                var v39 float32 = (((-2.25 / 1.0) - 1e10) / 1.5)
                _ = v39
            }
            v40 := GenS4{F0: (v33 >> (((^(v33))) & 63)), F1: v4, F2: -48}
            _ = v37
            _ = v40
        } else {
            v33 = (v33 << (((^(v33))) & 63))
        }
        _ = v34
    }
    var v41 uint16 = 255
    if (!((0.0 >= 0.1) || (42 == -223))) {
        v41 = (v41 | v41)
    } else {
        v1 -= (v1 * uint32(v33))
        var v42 int64 = -1
        switch 226 {
        case 9223372036854775807:
            v43 := GenS1{F0: 100, F1: 1, F2: -2147483648}
            _ = v43
        }
        _ = v42
    }
    if (!false) {
        for fi44 := 0; fi44 < 11; fi44++ {
            fmt.Println(v3)
            _ = fi44
        }
        fmt.Println(v41, v3, v1)
        var v45 float64 = 1e10
        _ = v45
    }
    v46 := GenS3{F0: ((v3 * 42) != v3), F1: 417}
    _ = len(v2)
    if (652 < 0) {
        v33 = 159
        if true {
            v3 = 853
            v47 := GenS4{F0: v33, F1: v4, F2: -2147483648}
            _ = v47
        } else {
            for fi48 := 0; fi48 < 3; fi48++ {
                if true {
                    v1 -= uint32(v3)
                    v49 := GenS3{F0: (!(1 >= 253)), F1: v33}
                    _ = v49
                } else {
                    v2 = append(v2, "foo bar")
                }
                for fi50 := 0; fi50 < 8; fi50++ {
                    var v51 int = (fi48 + (fi48 / ((fi48) | 1)))
                    v2 = append(v2, "foo bar")
                    v33 -= (v33 + 0)
                    _ = fi50
                    _ = v51
                }
                switch "foo bar" {
                case "":
                    if false {
                        if (!true) {
                            switch "" {
                            case "":
                                if false {
                                    switch "a" {
                                    case "foo bar":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "":
                                        var v52 uint16 = v41
                                        _ = v52
                                    }
                                } else {
                                    var v53 float64 = -2.25
                                    switch "hello" {
                                    case "hello":
                                        switch fi48 {
                                        case 393:
                                            v54 := GenS4{F0: (v33 % (((v33 + (v33 & (^(v33))))) | 1)), F1: 42, F2: 42}
                                            _ = v54
                                        case -977:
                                            v4 += (v4 ^ (v4 | 2))
                                        default:
                                            _ = len(v2)
                                        }
                                    case "\n\t":
                                        fmt.Println(v3, v4, v41)
                                    case "a":
                                        if len(v2) > 0 {
                                            v2[0] = "a"
                                        }
                                    default:
                                        v2 = append(v2, "hello")
                                    }
                                    _ = v53
                                }
                            case "foo bar":
                                for fi55 := 0; fi55 < 9; fi55++ {
                                    var v56 float32 = 3.14159
                                    if true {
                                        _ = len(v2)
                                        v56 += v56
                                        v2 = append(v2, (("foo bar" + "foo bar") + ""))
                                        v33 = uint(v1)
                                        v41 += (v41 % ((32767) | 1))
                                    } else {
                                        if false {
                                            var v57 float32 = v56
                                            switch "\n\t" {
                                            case "foo bar":
                                                if (365 == v3) {
                                                    fmt.Println(v56, fi55, fi48)
                                                    fmt.Println(v56)
                                                    var v58 int = fi55
                                                    _ = v58
                                                }
                                            case "a":
                                                v59 := GenS1{F0: -128, F1: (v1 * v1), F2: -1}
                                                _ = v59
                                            case "":
                                                var v60 int64 = int64(fi48)
                                                _ = v60
                                            default:
                                                fmt.Println(v3)
                                            }
                                            var v61 bool = ((v4 >> ((408) & 63)) == 1)
                                            var v62 float64 = (3.14159 + float64((1.5 * (v56 - -2.25))))
                                            _ = v57
                                            _ = v61
                                            _ = v62
                                        } else {
                                            var v63 bool = false
                                            _ = v63
                                        }
                                    }
                                    if (0 < 2147483647) {
                                        fmt.Println(fi55, v41)
                                        fmt.Println(v1, v4, v3)
                                    } else {
                                        v4 += v4
                                        fmt.Println(fi55, fi48, v1)
                                        var v64 string = ("hello" + "\n\t")
                                        _ = v64
                                    }
                                    _ = fi55
                                    _ = v56
                                }
                            case "hello":
                                if len(v2) > 0 {
                                    v2[0] = ("a" + "hello")
                                }
                            }
                            for fi65 := 0; fi65 < 15; fi65++ {
                                fmt.Println(v3, v1)
                                if false {
                                    v4 = (^(v4))
                                    v33 = v33
                                    _ = len(v2)
                                } else {
                                    var v66 uint16 = uint16(v33)
                                    _ = v66
                                }
                                _ = fi65
                            }
                        } else {
                            if true {
                                var v67 float32 = (-2.25 + (1.5 + (0.0 / float32(((1.5 / float64(1e10)) / (1.5 + 3.14159))))))
                                v41 += (^(v41))
                                _ = v67
                            } else {
                                var v68 int8 = 84
                                for fi69 := 0; fi69 < 3; fi69++ {
                                    v70 := GenS5{F0: 9223372036854775807}
                                    v71 := GenS4{F0: v33, F1: 146, F2: 161}
                                    _ = fi69
                                    _ = v70
                                    _ = v71
                                }
                                var v72 bool = (false && (-1 >= int16(v68)))
                                _ = v68
                                _ = v72
                            }
                        }
                        v73 := GenS6{F0: -1}
                        _ = v73
                    } else {
                        _ = len(v2)
                        var v74 int32 = 0
                        v33 += v33
                        _ = v74
                    }
                case "foo bar":
                    v2 = append(v2, "")
                }
                _ = fi48
            }
            v3 -= v3
            switch 2147483647 {
            case -1:
                v75 := GenS5{F0: 42}
                _ = v75
            case -370:
                if true {
                    v3 += v3
                }
            case 32768:
                if ((-21 != 1) && false) {
                    for fi76 := 0; fi76 < 14; fi76++ {
                        v33 -= (v33 + 1)
                        v33 = 255
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi76
                    }
                    fmt.Println(v3, v3, v41)
                    v41 -= v41
                }
            default:
                for fi77 := 0; fi77 < 9; fi77++ {
                    v78 := GenS7{F0: v3, F1: -319}
                    if len(v2) > 0 {
                        v2[0] = ("hello" + "a")
                    }
                    if (!(true && (!(32767 != 699)))) {
                        v79 := GenS1{F0: -84, F1: v1, F2: -2147483648}
                        var v80 int16 = 329
                        v2 = append(v2, "a")
                        _ = v79
                        _ = v80
                    } else {
                        if false {
                            var v81 uint64 = 42
                            for fi82 := 0; fi82 < 4; fi82++ {
                                switch fi82 {
                                case 2147483647:
                                    if len(v2) > 0 {
                                        v2[0] = ("foo bar" + "hello")
                                    }
                                case -833:
                                    var v83 uint32 = v1
                                    _ = v83
                                default:
                                    v33 = (v33 * 1)
                                    var v84 uint64 = (v3 * uint64(v41))
                                    var v85 uint = (v33 * v33)
                                    _ = v84
                                    _ = v85
                                }
                                v86 := GenS2{F0: 284, F1: v1, F2: 9223372036854775807}
                                var v87 int64 = 9223372036854775807
                                if false {
                                    fmt.Println(fi82, fi77, v1)
                                    switch "hello" {
                                    case "foo bar":
                                        v3 = (v3 - (v81 * (v81 | 1)))
                                    case "hello":
                                        switch fi82 {
                                        case 128:
                                            v81 += (v3 - (^(v3)))
                                        case 2147483647:
                                            fmt.Println(v1, v87)
                                        default:
                                            v88 := GenS7{F0: 2, F1: -2147483648}
                                            _ = v88
                                        }
                                    case "\n\t":
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                    default:
                                        switch (^(fi77)) {
                                        case -1:
                                            v2 = append(v2, ("" + "\n\t"))
                                        case 4294967295:
                                            v3 = (^(v3))
                                            var v89 bool = true
                                            var v90 int32 = -1
                                            _ = v89
                                            _ = v90
                                        default:
                                            var v91 float64 = float64(0.1)
                                            _ = v91
                                        }
                                    }
                                    var v92 int8 = 100
                                    _ = v92
                                }
                                _ = fi82
                                _ = v86
                                _ = v87
                            }
                            _ = v81
                        }
                        fmt.Println(v3, v3)
                        if true {
                            if (3.14159 >= 1.5) {
                                v93 := GenS8{F0: v3, F1: -1, F2: (v3 * (v3 * v3))}
                                var v94 uint64 = v3
                                var v95 int16 = 535
                                _ = v93
                                _ = v94
                                _ = v95
                            }
                            v3 = uint64(v1)
                        }
                    }
                    _ = fi77
                    _ = v78
                }
            }
        }
    } else {
        if (0 != -58) {
            if false {
                if false {
                    v33 += v33
                    var v96 int8 = 0
                    var v97 uint16 = v41
                    _ = v96
                    _ = v97
                } else {
                    var v98 float32 = -2.25
                    _ = v98
                }
                v2 = append(v2, ("hello" + ""))
            } else {
                var v99 float64 = (3.14159 * ((1e10 * 1e10) * ((0.0 * -2.25) - 3.14159)))
                var v100 uint16 = v41
                _ = v100
                _ = v99
            }
        }
    }
    if false {
        if true {
            for fi101 := 0; fi101 < 18; fi101++ {
                var v102 uint16 = 65535
                _ = fi101
                _ = v102
            }
            var v103 bool = (!false)
            v103 = (!((-32768 == 1) && v103))
            _ = v103
        }
    }
    v104 := GenS4{F0: v33, F1: v4, F2: 42}
    for fi105 := 0; fi105 < 14; fi105++ {
        fmt.Println(v1, v33)
        if (false && (0.0 >= float64(0.1))) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v106 string = "foo bar"
            _ = v106
        }
        var v107 uint = uint(v4)
        _ = fi105
        _ = v107
    }
    v108 := GenS3{F0: ((!(1 >= 1)) || (!false)), F1: uint(v3)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v104
    _ = v108
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v41
    _ = v46
}
