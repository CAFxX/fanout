package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 bool; F1 uint; F2 string }
type GenS5 struct { F0 string; F1 uint16; F2 int32 }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS8 struct { F0 int32; F1 uint32; F2 string }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int; F1 int64; F2 uint16 }
type GenS11 struct { F0 uint32; F1 int32; F2 uint }
type GenS12 struct { F0 string; F1 bool }
type GenS13 struct { F0 uint8; F1 uint; F2 int64 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int16; F1 uint8; F2 int }
type GenS16 struct { F0 int64; F1 uint32; F2 uint16 }
func genFn0(p0 uint64, p1 int32) int {
    v7 := []int{-2147483648, 128}
    for fi8 := 0; fi8 < 17; fi8++ {
        if true {
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 uint8 = 26
                var v11 float32 = 1.5
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return 2147483647
}
func genFn1(p0 uint8, p1 bool) int16 {
    p0 = (p0 | p0)
    v12 := []int{}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    var v1 bool = ((!(!(!true))) || false)
    var v2 uint64 = 2
    v3 := []uint32{42, 2147483648, 42, 73}
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    v6 := make(map[uint32]uint32)
    fmt.Println(v2, v2)
    if true {
        v13 := GenS1{F0: 1, F1: genFn1(14, ((42 >= -128) || ((v1 || v1) || (v1 || false))))}
        fmt.Println(v2, v1)
        _ = v13
    } else {
        var v14 int = 455
        _ = v14
    }
    if v1 {
        v15 := GenS1{F0: -1, F1: 1}
        switch "foo bar" {
        case "\n\t":
            v16 := GenS1{F0: 32767, F1: 238}
            _ = v16
        case "hello":
            v17 := GenS1{F0: 1, F1: -115}
            _ = v17
        case "a":
            delete(v4, 127)
        default:
            if ((1 != 2) || false) {
                v18 := GenS2{F0: (!(("\n\t" + "a") <= "a")), F1: 24}
                v3 = append(v3, 903)
                for fi19 := 0; fi19 < 3; fi19++ {
                    v20 := GenS2{F0: (!v1), F1: 10}
                    v1 = (32767 > 1000)
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
        }
        v2 += 2
        _ = v15
    }
    var v21 int = 0
    if ((0 <= 9223372036854775807) || (0 <= 42)) {
        v21 += (v21 << ((2) & 63))
        v1 = (65535 <= 65535)
    }
    v22 := GenS1{F0: -810, F1: 0}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v23 := GenS3{F0: (112 < -128), F1: 32767}
    v24 := v5[v21]
    delete(v5, (v21 % ((genFn0(v2, 2147483647)) | 1)))
    v25 := GenS2{F0: v1, F1: 42}
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 int64 = 631
        if (-2.25 == 0.1) {
            v28 := GenS4{F0: (!(255 <= 32767)), F1: 238, F2: ""}
            if (!(v24 && v24)) {
                v21 = -1
                var v29 int32 = 342
                _ = v29
            } else {
                v30 := v4[(fi26 ^ (fi26 / ((9223372036854775807) | 1)))]
                if (!true) {
                    switch 128 {
                    case 1:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -630:
                        var v31 float32 = float32((float64(1.5) - (0.1 + 0.1)))
                        _ = v31
                    case -837:
                        v32 := GenS5{F0: v30, F1: 255, F2: -1}
                        _ = v32
                    default:
                        v33 := GenS6{F0: 0, F1: 65535}
                        _ = v33
                    }
                    v34 := GenS4{F0: v1, F1: 18446744073709551615, F2: "\n\t"}
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 19; fi35++ {
                        v36 := GenS7{F0: 0, F1: (v2 + v2), F2: 42}
                        _ = fi35
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 8; fi37++ {
                        if (1 > 942) {
                            v1 = (1 <= -635)
                            if true {
                                var v38 uint = 681
                                _ = v38
                            } else {
                                var v39 int16 = -1
                                if ((!(!(false || v1))) || true) {
                                    var v40 int8 = 0
                                    _ = v40
                                } else {
                                    v41 := GenS4{F0: v1, F1: 2, F2: (v30 + v30)}
                                    v42 := GenS5{F0: ((v30 + "a") + "\n\t"), F1: 204, F2: -898}
                                    fmt.Println(v24, v1, v24)
                                    var v43 uint = 1
                                    v21 = genFn0(v2, -707)
                                    v3 = append(v3, 2)
                                    if (0 <= int8(v43)) {
                                        v44 := GenS6{F0: -2147483648, F1: 42}
                                        if len(v3) > 0 {
                                            v3[0] = uint32(v43)
                                        }
                                        _ = v44
                                    }
                                    switch (v21 >> ((1) & 63)) {
                                    case -9223372036854775808:
                                        if ((4294967295 > 777) || v1) {
                                            fmt.Println(v43, v27)
                                            v45 := GenS8{F0: 68, F1: 2147483647, F2: (v30 + "hello")}
                                            var v46 uint = 65535
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                                switch fi26 {
                                case 4294967295:
                                    v47 := GenS2{F0: (v39 <= v39), F1: 32767}
                                    _ = v47
                                case 2147483647:
                                    switch ("foo bar" + v30) {
                                    case "a":
                                        if ((v39 << ((970) & 63)) <= v39) {
                                            fmt.Println(v24, v24)
                                            v1 = (85 == 122)
                                        }
                                    case "hello":
                                        var v48 uint32 = 4294967295
                                        _ = v48
                                    case "\n\t":
                                        v24 = ((fi37 % ((genFn0(1, 1)) | 1)) == (v21 - -9223372036854775808))
                                        if (!v1) {
                                            if v24 {
                                                if len(v3) > 0 {
                                                    v3[0] = 931
                                                }
                                                v49 := GenS6{F0: 489, F1: 0}
                                                var v50 uint64 = v2
                                                switch -1 {
                                                case -890:
                                                    v39 += (v39 << ((255) & 63))
                                                default:
                                                    fmt.Println(v2, fi37, v50)
                                                }
                                                var v51 uint8 = 128
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                            } else {
                                                fmt.Println(v2, v2, v1)
                                            }
                                            v21 = v21
                                            if len(v3) > 0 {
                                                v3[0] = 2147483648
                                            }
                                        } else {
                                            fmt.Println(fi37, fi26, v21)
                                        }
                                    default:
                                        var v52 int64 = (v27 + v27)
                                        _ = v52
                                    }
                                }
                                v53 := GenS9{F0: v27}
                                _ = v39
                                _ = v53
                            }
                            v1 = v24
                        } else {
                            v24 = v24
                            if (!(32767 != -32768)) {
                                switch fi37 {
                                case -667:
                                    if ((fi26 & (fi37 >> ((65535) & 63))) <= fi37) {
                                        var v54 uint32 = 2147483647
                                        _ = len(v3)
                                        _ = v54
                                    }
                                case 4294967295:
                                    switch "hello" {
                                    case "a":
                                        delete(v6, 4294967295)
                                        v55 := GenS10{F0: (fi26 + fi37), F1: v27, F2: 2}
                                        fmt.Println(v24, fi37, fi26)
                                        v4[32768] = v30
                                        if ((v2 * 743) <= 42) {
                                            var v56 int8 = 33
                                            var v57 int64 = 42
                                            v58 := GenS11{F0: 2147483648, F1: 0, F2: 4294967295}
                                            _ = v56
                                            _ = v57
                                            _ = v58
                                        } else {
                                            v21 -= (fi37 + 353)
                                            switch "\n\t" {
                                            case "":
                                                v6[4294967295] = 1
                                            case "a":
                                                var v59 uint32 = 248
                                                _ = v59
                                            case "hello":
                                                var v60 string = (v30 + v30)
                                                _ = v60
                                            default:
                                                v21 = v21
                                            }
                                            v2 += (v2 % (((v2 << ((925) & 63))) | 1))
                                        }
                                        var v61 float64 = 1.5
                                        _ = v55
                                        _ = v61
                                    case "hello":
                                        switch genFn0((v2 << ((2) & 63)), 0) {
                                        case 4294967295:
                                            if (1 > 168) {
                                                if v1 {
                                                    var v62 uint16 = 0
                                                    _ = v62
                                                } else {
                                                    v63 := GenS1{F0: 42, F1: 42}
                                                    if ((fi37 - fi26) <= (fi37 >> ((169) & 63))) {
                                                        var v64 uint = 255
                                                        var v65 int16 = 42
                                                        _ = v64
                                                        _ = v65
                                                    } else {
                                                        var v66 float32 = (((1.5 + float32(float64(3.14159))) * 0.1) + 3.14159)
                                                        v3 = append(v3, 2147483648)
                                                        _ = v66
                                                    }
                                                    if (255 < 2) {
                                                        if true {
                                                            if v1 {
                                                                var v67 string = (v30 + v30)
                                                                _ = v67
                                                            }
                                                        }
                                                        var v68 uint8 = 228
                                                        _ = v68
                                                    } else {
                                                        fmt.Println(v2)
                                                        v69 := GenS12{F0: (v30 + "\n\t"), F1: false}
                                                        if (!(v1 || (!true))) {
                                                            v70 := GenS11{F0: 215, F1: 608, F2: 0}
                                                            switch genFn0(737, -123) {
                                                            case 42:
                                                                if (70 > 56) {
                                                                    v21 -= fi37
                                                                    v71 := GenS13{F0: 1, F1: 399, F2: v27}
                                                                    _ = v71
                                                                }
                                                            case 128:
                                                                var v72 int = (v21 * genFn0(v2, 2147483647))
                                                                _ = v72
                                                            case 813:
                                                                v1 = (!v24)
                                                            default:
                                                                var v73 int32 = 2147483647
                                                                _ = v73
                                                            }
                                                            var v74 uint = 964
                                                            _ = v70
                                                            _ = v74
                                                        } else {
                                                            if (0 >= genFn1(127, true)) {
                                                                v21 -= 277
                                                            } else {
                                                                v75 := GenS2{F0: v1, F1: 65535}
                                                                _ = v75
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = 621
                                                            }
                                                            fmt.Println(v30)
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = uint32(v2)
                                                        }
                                                        _ = v69
                                                    }
                                                    _ = v63
                                                }
                                                v1 = false
                                                v21 = (fi26 ^ v21)
                                            }
                                        case 32767:
                                            v3 = append(v3, 4294967295)
                                        case 685:
                                            v1 = (1.5 > 1e10)
                                        default:
                                            v27 = v27
                                        }
                                    default:
                                        v76 := GenS1{F0: 1, F1: -720}
                                        _ = v76
                                    }
                                case -9223372036854775808:
                                    v77 := v5[1]
                                    _ = v77
                                }
                                fmt.Println(v2, v1, v1)
                            }
                        }
                        var v78 int = 9223372036854775807
                        v5[fi26] = (((1.5 * 0.1) / 1e10) >= 0.1)
                        var v79 bool = (((^(v2)) >= 1) && v24)
                        _ = fi37
                        _ = v78
                        _ = v79
                    }
                    if (-2.25 > -2.25) {
                        v80 := GenS10{F0: (fi26 + 4294967295), F1: (^(v27)), F2: 1}
                        v81 := GenS14{F0: v24}
                        for fi82 := 0; fi82 < 2; fi82++ {
                            var v83 uint32 = 419
                            var v84 int = fi26
                            v85 := v6[v83]
                            var v86 float32 = (-2.25 - (1e10 + float32(-2.25)))
                            _ = fi82
                            _ = v83
                            _ = v84
                            _ = v85
                            _ = v86
                        }
                        _ = v80
                        _ = v81
                    }
                }
                _ = v30
            }
            if (v1 && v1) {
                v2 = 596
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                v2 -= v2
                var v87 uint8 = 128
                fmt.Println(v27)
                _ = v87
            }
            _ = v28
        } else {
            var v88 int64 = v27
            if (42 > uint(v2)) {
                if (65535 <= 1) {
                    var v89 uint8 = 127
                    v27 += (v88 ^ -137)
                    var v90 int16 = 1
                    _ = v89
                    _ = v90
                }
                var v91 int16 = 1000
                _ = v91
            }
            var v92 uint8 = 25
            _ = v88
            _ = v92
        }
        _ = fi26
        _ = v27
    }
    for fi93 := 0; fi93 < 18; fi93++ {
        v94 := GenS15{F0: 1000, F1: 20, F2: (fi93 ^ v21)}
        for fi95 := 0; fi95 < 17; fi95++ {
            var v96 int = (-(fi95))
            if true {
                v97 := GenS16{F0: -9223372036854775808, F1: 42, F2: 32767}
                delete(v6, 2147483647)
                var v98 string = "hello"
                _ = v97
                _ = v98
            } else {
                v2 -= 1
            }
            _ = fi95
            _ = v96
        }
        for fi99 := 0; fi99 < 16; fi99++ {
            v100 := GenS7{F0: 232, F1: (v2 << ((65535) & 63)), F2: 42}
            v1 = (!(v24 || true))
            v101 := GenS5{F0: "a", F1: uint16(v2), F2: 2147483647}
            v24 = (!(v2 != v2))
            v24 = ((("foo bar" + "foo bar") + "") > "hello")
            _ = fi99
            _ = v100
            _ = v101
        }
        _ = fi93
        _ = v94
    }
    var v102 uint8 = 127
    fmt.Println(v2, v21, v1)
    for fi103 := 0; fi103 < 9; fi103++ {
        var v104 uint16 = 42
        _ = fi103
        _ = v104
    }
    var v105 int8 = -128
    var v106 int8 = v105
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v102
    _ = v105
    _ = v106
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
