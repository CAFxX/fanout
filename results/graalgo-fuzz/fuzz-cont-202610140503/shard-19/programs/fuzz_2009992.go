package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int32; F2 int }
type GenS4 struct { F0 uint64; F1 uint64; F2 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 int64; F1 int16 }
type GenS10 struct { F0 uint16; F1 uint }
type GenS11 struct { F0 uint; F1 bool; F2 uint64 }
type GenS12 struct { F0 uint; F1 bool }
func genFn0(p0 uint8, p1 bool) uint64 {
    var v4 uint32 = 2147483647
    v5 := GenS1{F0: -1, F1: 4294967295}
    fmt.Println(p1, v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int64, p1 uint8, p2 uint) int32 {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 float64 = float64(float32(1e10))
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 int32 = -1
        v10 := []int{(fi8 + fi8), (fi8 << (((p2 - 2)) & 63)), (fi8 - 1), (fi8 - fi8)}
        v11 := GenS2{F0: v9}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println(p0)
    switch ("foo bar" + "hello") {
    case "a":
        if false {
            if true {
                v12 := GenS3{F0: -66, F1: 1, F2: 42}
                v13 := GenS4{F0: genFn0(128, false), F1: 0, F2: false}
                _ = v12
                _ = v13
            } else {
                if ((false || false) || false) {
                    fmt.Println(p1)
                    p1 = (p1 - p1)
                } else {
                    p2 += 611
                    var v14 uint = p2
                    for fi15 := 0; fi15 < 11; fi15++ {
                        for fi16 := 0; fi16 < 9; fi16++ {
                            v17 := []int{}
                            if true {
                                var v18 int16 = 1
                                switch ("foo bar" + "") {
                                case "\n\t":
                                    if ((p1 + p1) < (p1 / ((p1) | 1))) {
                                        v19 := GenS5{F0: 1}
                                        v20 := GenS6{F0: (p0 % ((-1) | 1)), F1: -458}
                                        if true {
                                            v18 -= v18
                                            fmt.Println(p0, fi15, p2)
                                        }
                                        _ = v19
                                        _ = v20
                                    }
                                case "foo bar":
                                    if len(v17) > 0 {
                                        v17[0] = -695
                                    }
                                }
                                _ = v18
                            }
                            _ = fi16
                            _ = v17
                        }
                        var v21 float64 = ((3.14159 * 0.0) + -2.25)
                        _ = fi15
                        _ = v21
                    }
                    _ = v14
                }
                var v22 int8 = 0
                _ = v22
            }
        }
    case "hello":
        var v23 int8 = 42
        _ = v23
    default:
        switch "" {
        case "":
            if false {
                for fi24 := 0; fi24 < 13; fi24++ {
                    switch (-(fi24)) {
                    case 9223372036854775807:
                        v25 := GenS5{F0: 18446744073709551615}
                        _ = v25
                    default:
                        p1 = (p1 | p1)
                    }
                    v26 := []int{fi24}
                    _ = fi24
                    _ = v26
                }
            } else {
                var v27 uint16 = 2
                var v28 int32 = 2147483647
                for fi29 := 0; fi29 < 13; fi29++ {
                    if true {
                        if true {
                            p0 -= (p0 & p0)
                            if ((p2 | 65535) > (p2 + (p2 ^ (p2 / (((p2 >> ((p2) & 63))) | 1))))) {
                                v30 := GenS6{F0: p0, F1: p0}
                                var v31 uint16 = (v27 | (v27 ^ (^(v27))))
                                _ = v30
                                _ = v31
                            } else {
                                v27 = 1
                                if (119 == 1) {
                                    v32 := GenS7{F0: p1}
                                    for fi33 := 0; fi33 < 1; fi33++ {
                                        v34 := GenS8{F0: 774, F1: (p2 + (p2 ^ (p2 << ((p2) & 63))))}
                                        v27 = uint16(p1)
                                        _ = fi33
                                        _ = v34
                                    }
                                    _ = v32
                                } else {
                                    var v35 int64 = (p0 / ((int64(v28)) | 1))
                                    for fi36 := 0; fi36 < 7; fi36++ {
                                        if true {
                                            v37 := GenS6{F0: v35, F1: (p0 * (v35 - v35))}
                                            _ = v37
                                        } else {
                                            if true {
                                                fmt.Println(v28, p2, p2)
                                                if false {
                                                    fmt.Println(v27)
                                                } else {
                                                    v38 := []int{(fi36 & 4294967295)}
                                                    if true {
                                                        fmt.Println(v28)
                                                        if true {
                                                            var v39 float32 = 3.14159
                                                            if len(v38) > 0 {
                                                                v38[0] = (fi29 & 520)
                                                            }
                                                            _ = v39
                                                        } else {
                                                            if len(v38) > 0 {
                                                                v38[0] = (^(fi36))
                                                            }
                                                        }
                                                        var v40 float32 = -2.25
                                                        _ = v40
                                                    } else {
                                                        var v41 int8 = 69
                                                        _ = v41
                                                    }
                                                    v38 = append(v38, 641)
                                                    v42 := GenS6{F0: p0, F1: p0}
                                                    _ = v38
                                                    _ = v42
                                                }
                                                v43 := GenS2{F0: v28}
                                                var v44 int32 = v28
                                                if true {
                                                    var v45 int16 = 1000
                                                    v46 := []int{}
                                                    switch 556 {
                                                    case 965:
                                                        switch fi36 {
                                                        case 961:
                                                            v47 := GenS9{F0: (p0 >> (((^(p2))) & 63)), F1: (v45 - (v45 >> (((p2 - 745)) & 63)))}
                                                            _ = v47
                                                        case -1:
                                                            v28 += v44
                                                        default:
                                                            var v48 uint = (p2 % ((p2) | 1))
                                                            _ = v48
                                                        }
                                                    case 9223372036854775807:
                                                        var v49 float64 = (float64(1e10) + float64((3.14159 * float32(-2.25))))
                                                        _ = v49
                                                    }
                                                    _ = v45
                                                    _ = v46
                                                }
                                                _ = v43
                                                _ = v44
                                            } else {
                                                var v50 float32 = 1e10
                                                if (-413 >= 0) {
                                                    var v51 int = fi29
                                                    _ = v51
                                                } else {
                                                    var v52 int32 = v28
                                                    v53 := GenS9{F0: v35, F1: -616}
                                                    var v54 uint32 = 2
                                                    _ = v52
                                                    _ = v53
                                                    _ = v54
                                                }
                                                v55 := []int{91, 42}
                                                _ = v50
                                                _ = v55
                                            }
                                            fmt.Println(v27, p2)
                                        }
                                        var v56 float32 = ((0.1 + ((float32(0.0) + 0.1) + float32(0.1))) - -2.25)
                                        v28 += v28
                                        var v57 int = -365
                                        var v58 int16 = 64
                                        var v59 uint16 = (v27 - v27)
                                        _ = fi36
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = v35
                                }
                                v27 += v27
                            }
                            var v60 uint16 = (^(v27))
                            _ = v60
                        } else {
                            var v61 uint32 = 752
                            var v62 uint32 = 4294967295
                            var v63 bool = true
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                        var v64 float64 = 1.5
                        _ = v64
                    } else {
                        v65 := []int{}
                        if (!(-2.25 <= (float32((0.1 * (0.0 + 1e10))) + 3.14159))) {
                            v28 -= -2147483648
                            if false {
                                if (true && false) {
                                    v66 := GenS10{F0: v27, F1: p2}
                                    _ = v66
                                } else {
                                    var v67 int = (fi29 - (-(fi29)))
                                    if len(v65) > 0 {
                                        v65[0] = int(p1)
                                    }
                                    _ = v67
                                }
                            } else {
                                v27 -= v27
                                for fi68 := 0; fi68 < 17; fi68++ {
                                    v69 := GenS1{F0: 179, F1: 520}
                                    p0 = p0
                                    var v70 uint8 = (p1 >> (((^(p2))) & 63))
                                    p0 = p0
                                    p1 = p1
                                    _ = fi68
                                    _ = v69
                                    _ = v70
                                }
                            }
                            if (false || (false && true)) {
                                p0 = 257
                                for fi71 := 0; fi71 < 1; fi71++ {
                                    v28 -= (v28 * 0)
                                    var v72 int8 = -128
                                    _ = fi71
                                    _ = v72
                                }
                                p0 += p0
                            }
                        }
                        v73 := GenS11{F0: p2, F1: true, F2: genFn0((^(p1)), false)}
                        _ = v65
                        _ = v73
                    }
                    _ = fi29
                }
                _ = v27
                _ = v28
            }
        case "\n\t":
            var v74 uint16 = 65535
            _ = v74
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 int64 = 0
    v2 := []uint32{4294967295, 869}
    v3 := make(map[int]int)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v75 := GenS1{F0: -9223372036854775808, F1: 104}
    for fi76 := 0; fi76 < 9; fi76++ {
        var v77 int8 = -128
        _ = fi76
        _ = v77
    }
    var v78 int = -167
    v79 := GenS12{F0: 4294967295, F1: (!false)}
    if false {
        v80 := GenS1{F0: 32, F1: 0}
        v1 -= (v1 & (-(v1)))
        _ = v80
    }
    fmt.Println(v78)
    v1 -= (v1 << ((0) & 63))
    if (2147483648 != 869) {
        var v81 uint32 = 2147483647
        var v82 uint8 = 96
        for fi83 := 0; fi83 < 12; fi83++ {
            var v84 uint32 = 65535
            _ = fi83
            _ = v84
        }
        _ = v81
        _ = v82
    } else {
        var v85 uint64 = 0
        if true {
            if (1 == -261) {
                var v86 float64 = 0.0
                _ = v86
            } else {
                var v87 int8 = 42
                _ = v87
            }
            var v88 uint = 1
            delete(v3, 9223372036854775807)
            _ = v88
        }
        _ = v85
    }
    switch "\n\t" {
    case "":
        if true {
            var v89 int64 = v1
            if true {
                var v90 int64 = (v1 | 42)
                delete(v3, (v78 % (((v78 % ((v78) | 1))) | 1)))
                fmt.Println(v78, v90, v90)
                _ = v90
            }
            _ = v89
        }
    case "foo bar":
        v1 = 0
    default:
        var v91 uint8 = 1
        _ = v91
    }
    var v92 uint = 255
    var v93 uint = (v92 ^ 18446744073709551615)
    delete(v3, v78)
    if false {
        v94 := GenS2{F0: 392}
        _ = v94
    } else {
        v1 = v1
        v95 := GenS4{F0: 923, F1: 1, F2: false}
        _ = v95
    }
    fmt.Println(v1, v78, v93)
    for fi96 := 0; fi96 < 16; fi96++ {
        fmt.Println(fi96, fi96, v93)
        _ = fi96
    }
    fmt.Println(v78)
    if false {
        if false {
            fmt.Println(v93, v93, v78)
            for fi97 := 0; fi97 < 12; fi97++ {
                var v98 uint8 = 1
                v93 = v92
                if (-32768 < 1) {
                    if (32767 < 65535) {
                        v98 -= (v98 / ((0) | 1))
                    } else {
                        fmt.Println(v93, fi97)
                    }
                    v1 += (v1 + (v1 >> (((v92 >> (((^(v92))) & 63))) & 63)))
                }
                _ = fi97
                _ = v98
            }
            v93 = v93
        }
        fmt.Println(v93, v78)
        v99 := GenS7{F0: 42}
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v78:", v78)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v75
    _ = v78
    _ = v79
    _ = v92
    _ = v93
}
