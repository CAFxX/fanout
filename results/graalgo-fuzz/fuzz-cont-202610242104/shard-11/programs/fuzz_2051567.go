package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 int; F2 string }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int64; F1 int64 }
type GenS10 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 bool; F1 uint32 }
func genFn0(p0 string) bool {
    v6 := []int{4294967295, -2147483648, 9223372036854775807, -195}
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = p0
    p0 = p0
    _ = p0
    _ = v6
    return ((true || true) || (0 != 42))
}
func genFn1(p0 int8, p1 uint8) uint64 {
    for fi7 := 0; fi7 < 19; fi7++ {
        if false {
            v8 := []int{fi7, (fi7 * (fi7 & (fi7 % (((fi7 >> ((18446744073709551615) & 63))) | 1)))), fi7}
            _ = v8
        }
        if false {
            switch "hello" {
            case "\n\t":
                for fi9 := 0; fi9 < 18; fi9++ {
                    p0 = 1
                    var v10 int32 = 42
                    fmt.Println(fi7, p0, v10)
                    _ = fi9
                    _ = v10
                }
            case "foo bar":
                v11 := GenS1{F0: 116}
                _ = v11
            case "a":
                switch ("a" + "a") {
                case "":
                    p0 = -1
                default:
                    p1 = (p1 ^ (^(p1)))
                }
            default:
                p1 += p1
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := []int{}
                var v14 uint32 = uint32(p1)
                p0 = p0
                _ = fi12
                _ = v13
                _ = v14
            }
            p1 += p1
        }
        _ = fi7
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        if true {
            if (p1 >= 255) {
                for fi16 := 0; fi16 < 13; fi16++ {
                    v17 := GenS2{F0: 2}
                    p0 = p0
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 19; fi18++ {
                    p1 -= 1
                    v19 := GenS2{F0: 0}
                    _ = fi18
                    _ = v19
                }
            } else {
                p0 -= (^(p0))
                if (p1 != (p1 | (^(p1)))) {
                    var v20 int8 = 1
                    var v21 int64 = -863
                    _ = v20
                    _ = v21
                }
            }
            if true {
                var v22 float32 = ((-2.25 + 0.0) / 1.5)
                _ = v22
            }
        } else {
            v23 := []int{(-(fi15)), (fi15 | -809), (fi15 | 32768)}
            p0 = 42
            p0 = 100
            _ = v23
        }
        var v24 int = (fi15 * (fi15 ^ (fi15 << ((255) & 63))))
        for fi25 := 0; fi25 < 10; fi25++ {
            p1 += 1
            fmt.Println(fi15, fi25, fi15)
            v26 := []int{fi25, fi15}
            _ = fi25
            _ = v26
        }
        _ = fi15
        _ = v24
    }
    p1 = (p1 / (((p1 | 128)) | 1))
    for fi27 := 0; fi27 < 7; fi27++ {
        var v28 bool = true
        v29 := GenS3{F0: (p0 >> ((2) & 63))}
        v30 := GenS2{F0: 501}
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    _ = p0
    _ = p1
    return 816
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 65535
    v3 := make(map[int]int)
    var v4 uint = 96
    var v5 uint8 = 225
    if (32767 > -32768) {
        var v31 uint = 839
        if ((-1 > -1) || true) {
            fmt.Println(v2, v2, v2)
        } else {
            switch 42 {
            case 4294967295:
                if (151 <= 1) {
                    for fi32 := 0; fi32 < 16; fi32++ {
                        var v33 uint = (v4 % (((^(v31))) | 1))
                        v4 += v4
                        _ = fi32
                        _ = v33
                    }
                }
            }
            v2 = v2
        }
        _ = v31
    } else {
        for fi34 := 0; fi34 < 13; fi34++ {
            for fi35 := 0; fi35 < 8; fi35++ {
                v5 = (v5 << ((v4) & 63))
                var v36 bool = false
                _ = fi35
                _ = v36
            }
            v37 := []int{}
            switch -1 {
            case -9223372036854775808:
                var v38 uint16 = v2
                _ = v38
            default:
                fmt.Println(fi34, v5)
            }
            _ = fi34
            _ = v37
        }
    }
    v5 = 128
    v39 := GenS2{F0: 65535}
    var v40 uint8 = v5
    if true {
        fmt.Println(v2)
    }
    v40 += (v5 * v40)
    if (9223372036854775807 == 4294967295) {
        switch "hello" {
        case "hello":
            if (v5 < v5) {
                var v41 uint = v4
                _ = v41
            } else {
                switch (("hello" + ("a" + "foo bar")) + "hello") {
                case "\n\t":
                    if false {
                        v42 := GenS4{F0: 2147483647, F1: 42}
                        v5 = 127
                        _ = len(v1)
                        _ = v42
                    } else {
                        var v43 string = "a"
                        v4 = 745
                        v44 := GenS5{F0: 1000}
                        _ = v43
                        _ = v44
                    }
                case "hello":
                    if (-2147483648 <= 80) {
                        for fi45 := 0; fi45 < 17; fi45++ {
                            v2 = (^(v2))
                            _ = fi45
                        }
                        for fi46 := 0; fi46 < 16; fi46++ {
                            var v47 uint64 = genFn1(1, (v5 & v40))
                            if true {
                                _ = len(v1)
                                for fi48 := 0; fi48 < 18; fi48++ {
                                    var v49 uint16 = (v2 + (v2 ^ (v2 >> ((42) & 63))))
                                    v49 = 42
                                    if (-2.25 == float64(0.1)) {
                                        switch (("hello" + ("a" + (("a" + ("a" + ("a" + ("a" + ("" + ("foo bar" + "a")))))) + ("\n\t" + "\n\t")))) + "hello") {
                                        case "":
                                            v1[4294967295] = 2
                                        case "hello":
                                            var v50 uint8 = (^(v40))
                                            _ = v50
                                        default:
                                            v51 := []int{(fi48 & fi48)}
                                            _ = v51
                                        }
                                    } else {
                                        fmt.Println(v47, v4, fi46)
                                        var v52 int = fi46
                                        v53 := v1[(fi48 >> (((v4 / (((v4 / ((v4) | 1))) | 1))) & 63))]
                                        _ = v52
                                        _ = v53
                                    }
                                    _ = fi48
                                    _ = v49
                                }
                                if ((v4 << ((4294967295) & 63)) != v4) {
                                    v54 := []int{fi46, fi46, (fi46 / ((fi46) | 1)), (fi46 / ((fi46) | 1))}
                                    _ = v54
                                }
                            } else {
                                v5 -= (v40 & (v5 >> ((v4) & 63)))
                                v55 := GenS6{F0: 1, F1: (fi46 >> ((v4) & 63)), F2: (("foo bar" + "a") + "a")}
                                if true {
                                    if ((1 != 0) || true) {
                                        fmt.Println(v5, v4, v2)
                                    } else {
                                        v2 = v2
                                    }
                                    for fi56 := 0; fi56 < 1; fi56++ {
                                        v40 = v5
                                        switch 1 {
                                        case -1:
                                            delete(v3, (fi46 + (fi46 & (-(fi46)))))
                                        case -915:
                                            v57 := []int{}
                                            if (397 < 551) {
                                                fmt.Println(v4, v4, v47)
                                            }
                                            _ = v57
                                        default:
                                            var v58 uint = v4
                                            var v59 uint64 = (v47 / (((v47 ^ (v47 + 887))) | 1))
                                            fmt.Println(v40)
                                            _ = v58
                                            _ = v59
                                        }
                                        var v60 float32 = float32((0.1 + 1e10))
                                        _ = fi56
                                        _ = v60
                                    }
                                    v2 = (v2 >> ((v4) & 63))
                                } else {
                                    switch 147 {
                                    case 127:
                                        v61 := GenS2{F0: 787}
                                        _ = v61
                                    case 876:
                                        v62 := GenS4{F0: -193, F1: uint64(v4)}
                                        _ = v62
                                    case -244:
                                        fmt.Println(v5, v47, v40)
                                    default:
                                        var v63 uint16 = (v2 / ((v2) | 1))
                                        _ = v63
                                    }
                                    if (-1 > int8(v2)) {
                                        if false {
                                            v64 := GenS7{F0: 42, F1: (!("" == ("" + "hello")))}
                                            _ = v64
                                        } else {
                                            var v65 int64 = 316
                                            var v66 uint8 = (v5 * 156)
                                            fmt.Println(v65)
                                            _ = v65
                                            _ = v66
                                        }
                                        switch fi46 {
                                        case 4294967295:
                                            if (0 >= -128) {
                                                var v67 uint16 = 42
                                                _ = v67
                                            }
                                        case -323:
                                            if (1.5 != 0.1) {
                                                v5 = v5
                                                if (!(-9223372036854775808 > -1)) {
                                                    var v68 uint32 = 59
                                                    v69 := GenS8{F0: 0}
                                                    _ = v68
                                                    _ = v69
                                                }
                                                switch "hello" {
                                                case "a":
                                                    delete(v3, (fi46 + (fi46 | (fi46 / (((-(fi46))) | 1)))))
                                                case "\n\t":
                                                    v70 := GenS3{F0: 44}
                                                    _ = v70
                                                case "":
                                                    for fi71 := 0; fi71 < 2; fi71++ {
                                                        v1[(fi46 + fi71)] = 1
                                                        _ = fi71
                                                    }
                                                }
                                            }
                                        case 0:
                                            var v72 int32 = 1
                                            _ = v72
                                        default:
                                            fmt.Println(v4, fi46, v2)
                                        }
                                        v5 = 128
                                    }
                                    v2 += v2
                                }
                                _ = v55
                            }
                            v73 := v1[fi46]
                            _ = fi46
                            _ = v47
                            _ = v73
                        }
                    }
                }
                fmt.Println(v2, v2)
                v4 -= (v4 & (v4 << (((v4 + v4)) & 63)))
            }
        case "\n\t":
            v4 = (v4 % (((v4 * v4)) | 1))
        default:
            for fi74 := 0; fi74 < 19; fi74++ {
                v75 := v1[(-(fi74))]
                var v76 int32 = 0
                _ = fi74
                _ = v75
                _ = v76
            }
        }
        if ((false && (1.5 < 0.0)) && false) {
            v4 = (v4 | 2)
            var v77 uint16 = (^(v2))
            fmt.Println(v77)
            _ = v77
        } else {
            if false {
                v2 = (^(v2))
                var v78 uint = v4
                _ = v78
            } else {
                v2 = 313
                var v79 bool = (int32(v2) <= 50)
                v79 = v79
                _ = v79
            }
            var v80 int8 = 1
            switch 9223372036854775807 {
            case -9223372036854775808:
                v81 := GenS9{F0: 42, F1: int64(v4)}
                _ = v81
            default:
                switch 1 {
                case 42:
                    var v82 uint = 0
                    _ = v82
                case -624:
                    v80 += -128
                case 174:
                    v2 += (v2 | 26)
                default:
                    for fi83 := 0; fi83 < 12; fi83++ {
                        var v84 float64 = (0.0 * (float64(0.1) - (float64(-2.25) / 1.5)))
                        v85 := GenS10{F0: -280, F1: v2, F2: (v40 + (v5 | 255))}
                        switch "hello" {
                        case "\n\t":
                            v86 := []int{(fi83 % ((fi83) | 1)), 9223372036854775807}
                            _ = v86
                        default:
                            for fi87 := 0; fi87 < 4; fi87++ {
                                v88 := GenS11{F0: 42}
                                var v89 uint = 42
                                v90 := GenS12{F0: true, F1: 42}
                                _ = fi87
                                _ = v88
                                _ = v89
                                _ = v90
                            }
                        }
                        _ = fi83
                        _ = v84
                        _ = v85
                    }
                }
            }
            _ = v80
        }
    } else {
        for fi91 := 0; fi91 < 15; fi91++ {
            if (false || ((^(v4)) < 619)) {
                v92 := v3[(fi91 + 826)]
                var v93 uint16 = (v2 << (((v4 & (v4 * v4))) & 63))
                _ = v92
                _ = v93
            } else {
                _ = len(v1)
            }
            fmt.Println(v40, v5)
            _ = fi91
        }
    }
    v94 := GenS8{F0: -2147483648}
    v95 := []int{439, 9223372036854775807, 663, 1}
    for fi96 := 0; fi96 < 14; fi96++ {
        v2 += 2
        fmt.Println(v5)
        _ = fi96
    }
    var v97 float32 = 3.14159
    v98 := v3[0]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v95:", len(v95), v95)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v94
    _ = v95
    _ = v97
    _ = v98
}
