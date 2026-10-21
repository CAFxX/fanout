package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 uint8; F1 int32; F2 string }
type GenS7 struct { F0 int; F1 uint64; F2 uint8 }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 int16; F1 uint; F2 int8 }
type GenS10 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS11 struct { F0 string; F1 uint16; F2 uint8 }
type GenS12 struct { F0 uint8; F1 string }
func genFn0(p0 uint) int16 {
    switch -141 {
    case -9223372036854775808:
        fmt.Println(p0)
    case -1:
        p0 -= p0
    case 9223372036854775807:
        var v5 int = -320
        _ = v5
    default:
        var v6 int8 = 1
        _ = v6
    }
    v7 := GenS1{F0: true, F1: 1}
    switch "hello" {
    case "foo bar":
        if true {
            var v8 string = "hello"
            fmt.Println(v8)
            if ((false || (!false)) && (false && (32767 > 288))) {
                switch 841 {
                case -874:
                    var v9 uint8 = 252
                    _ = v9
                case -2147483648:
                    v8 = v8
                }
            }
            _ = v8
        } else {
            var v10 int = -9223372036854775808
            _ = v10
        }
    default:
        v11 := []int{4294967295, -9223372036854775808, -2147483648}
        _ = v11
    }
    v12 := GenS1{F0: false, F1: 224}
    if (("\n\t" + "foo bar") != ("foo bar" + "hello")) {
        var v13 int64 = -9223372036854775808
        if (((0.0 - 1.5) != 0.0) || true) {
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 float32 = 3.14159
                v13 -= (v13 + v13)
                _ = fi14
                _ = v15
            }
        } else {
            v16 := []int{}
            _ = v16
        }
        _ = v13
    }
    _ = p0
    _ = v12
    _ = v7
    return 1
}
func genFn1(p0 int8, p1 bool, p2 uint16) int64 {
    var v17 uint32 = 721
    var v18 int = -101
    v19 := GenS1{F0: (1e10 == 0.1), F1: 0}
    v18 = 4294967295
    v20 := GenS2{F0: 813}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v19
    _ = v20
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 255
    var v3 uint64 = 42
    var v4 int = 957
    v21 := GenS1{F0: true, F1: 42}
    v2 = uint8(v3)
    fmt.Println(v2)
    var v22 uint = 1
    fmt.Println(v2)
    var v23 uint16 = 8
    v23 = (v23 / (((^(v23))) | 1))
    if (true || false) {
        var v24 float32 = (0.1 / 1.0)
        v25 := []int{int(v2), v4, int(v23)}
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 15; fi26++ {
            if true {
                switch "a" {
                case "foo bar":
                    v4 -= (fi26 & (-(v4)))
                case "\n\t":
                    v27 := GenS3{F0: (v3 * (v3 & v3)), F1: 42}
                    _ = v27
                case "hello":
                    if (1 <= 42) {
                        for fi28 := 0; fi28 < 8; fi28++ {
                            _ = len(v1)
                            if false {
                                if (true && (false && false)) {
                                    v29 := GenS1{F0: true, F1: 796}
                                    var v30 float32 = (1.5 / float32((3.14159 + 3.14159)))
                                    if true {
                                        switch "a" {
                                        case "":
                                            v31 := GenS1{F0: ((1 >= 0) || (((("" + (("foo bar" + "foo bar") + "foo bar")) + "foo bar") + "") != "hello")), F1: -153}
                                            _ = v31
                                        }
                                        var v32 uint32 = 2
                                        var v33 int8 = 100
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v34 bool = (uint8(v3) < uint8(v23))
                                    fmt.Println(fi26, v3)
                                    if v34 {
                                        if v34 {
                                            v35 := GenS4{F0: true}
                                            fmt.Println(fi26, v3)
                                            var v36 string = ""
                                            var v37 int16 = 32767
                                            v37 += (-(v37))
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        }
                                        switch ("a" + ("" + "\n\t")) {
                                        case "":
                                            v23 = (v23 & (v23 + (v23 >> (((v22 << ((v22) & 63))) & 63))))
                                        case "a":
                                            fmt.Println(v3)
                                        case "\n\t":
                                            v4 = fi26
                                        }
                                    } else {
                                        v2 = 247
                                        var v38 uint8 = v2
                                        switch "a" {
                                        case "":
                                            v4 += (-(v4))
                                        }
                                        v39 := GenS4{F0: (((((((float64(1.5) + 3.14159) + -2.25) - 3.14159) + float64((1e10 - ((1e10 * 0.0) + (float32(3.14159) * (-2.25 / (0.1 / -2.25))))))) + 1.5) * float64(-2.25)) == float64(1e10))}
                                        delete(v1, (fi28 + v4))
                                        var v40 float32 = (0.1 * 0.0)
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                    if (((65535 <= v22) || true) || v34) {
                                        var v41 bool = v34
                                        switch "" {
                                        case "a":
                                            switch "\n\t" {
                                            case "":
                                                v42 := GenS3{F0: (^(v3)), F1: 0}
                                                _ = v42
                                            case "hello":
                                                v43 := GenS5{F0: 128, F1: 1}
                                                _ = v43
                                            default:
                                                var v44 int64 = -1
                                                _ = v44
                                            }
                                        case "\n\t":
                                            v45 := []int{(v4 * 9223372036854775807), fi26}
                                            _ = v45
                                        }
                                        _ = v41
                                    }
                                    switch "" {
                                    case "":
                                        if (((v2 - 2) <= (v2 + v2)) || ((fi28 % ((fi28) | 1)) < v4)) {
                                            v22 = v22
                                            v46 := GenS6{F0: (v2 ^ v2), F1: 659, F2: "a"}
                                            switch ("" + "hello") {
                                            case "a":
                                                var v47 uint64 = v3
                                                _ = v47
                                            case "\n\t":
                                                var v48 float64 = (0.1 * float64((3.14159 * float32(1e10))))
                                                _ = v48
                                            default:
                                                _ = len(v1)
                                            }
                                            fmt.Println(fi26, v3, v4)
                                            var v49 int8 = 42
                                            var v50 bool = v34
                                            v51 := v1[fi28]
                                            _ = v46
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        } else {
                                            fmt.Println(fi28, fi28, fi28)
                                            delete(v1, fi28)
                                            var v52 int16 = int16(fi26)
                                            _ = v52
                                        }
                                    case "a":
                                        var v53 string = "\n\t"
                                        _ = v53
                                    }
                                    var v54 string = ""
                                    if (uint64(v22) == v3) {
                                        v55 := []int{v4, 128, (-(fi28)), (fi28 + (v4 / (((-(fi26))) | 1)))}
                                        switch (v4 << ((18446744073709551615) & 63)) {
                                        case 667:
                                            fmt.Println(v34, fi28)
                                        case 1:
                                            delete(v1, fi28)
                                            if (127 == 1) {
                                                v2 = v2
                                                v56 := GenS7{F0: (fi26 >> ((v22) & 63)), F1: v3, F2: 19}
                                                _ = v56
                                            }
                                            if true {
                                                var v57 uint = (v22 ^ (^(v22)))
                                                v54 = ""
                                                _ = v57
                                            }
                                        }
                                        delete(v1, (fi28 - v4))
                                        _ = v55
                                    }
                                    _ = v34
                                    _ = v54
                                }
                            } else {
                                v3 = v3
                                var v58 float64 = 0.0
                                if false {
                                    var v59 int32 = -1
                                    _ = v59
                                }
                                _ = v58
                            }
                            var v60 float64 = 1.5
                            _ = fi28
                            _ = v60
                        }
                    } else {
                        var v61 uint64 = v3
                        _ = v61
                    }
                }
                fmt.Println(v4)
                var v62 uint32 = 319
                _ = v62
            } else {
                v3 = v3
                var v63 int8 = 42
                v3 = 42
                _ = v63
            }
            v3 += (v3 >> ((uint(v23)) & 63))
            _ = fi26
        }
    }
    v64 := GenS6{F0: (^(v2)), F1: 42, F2: ""}
    fmt.Println(v3, v3, v3)
    if (2147483647 > 2147483648) {
        for fi65 := 0; fi65 < 17; fi65++ {
            var v66 uint = (v22 & v22)
            _ = fi65
            _ = v66
        }
        v67 := GenS1{F0: (!(!(273 < v3))), F1: -477}
        _ = v67
    } else {
        v68 := GenS6{F0: v2, F1: 0, F2: ("a" + "foo bar")}
        v22 = 292
        switch "" {
        case "hello":
            var v69 int8 = 127
            _ = v69
        case "a":
            switch (v4 << (((v22 | 109)) & 63)) {
            case 128:
                fmt.Println(v3, v3)
            case -952:
                v70 := GenS7{F0: -398, F1: (v3 >> ((v22) & 63)), F2: 168}
                _ = v70
            case 4294967295:
                switch int(v22) {
                case -604:
                    switch v4 {
                    case 2147483647:
                        for fi71 := 0; fi71 < 7; fi71++ {
                            if (-9223372036854775808 < int64(v2)) {
                                v72 := []int{4294967295, (v4 ^ -9223372036854775808)}
                                v73 := GenS8{F0: uint(v3), F1: -1}
                                fmt.Println(v22, v2, v4)
                                _ = v72
                                _ = v73
                            } else {
                                v74 := GenS9{F0: -1, F1: v22, F2: 0}
                                _ = v74
                            }
                            fmt.Println(fi71, v2, v22)
                            switch "hello" {
                            case "a":
                                if (66 <= 67) {
                                    v75 := GenS1{F0: true, F1: -556}
                                    _ = v75
                                }
                            case "\n\t":
                                if false {
                                    var v76 uint32 = 2147483648
                                    if true {
                                        fmt.Println(v23)
                                        if true {
                                            v77 := GenS10{F0: uint8(v22), F1: 370, F2: v3}
                                            v3 = 528
                                            _ = v77
                                        } else {
                                            v78 := []int{(^(v4)), 847, -1}
                                            _ = v78
                                        }
                                    } else {
                                        fmt.Println(v76)
                                    }
                                    var v79 uint8 = (v2 * v2)
                                    _ = v76
                                    _ = v79
                                } else {
                                    v23 -= 42
                                    v23 += 32767
                                    switch -642 {
                                    case -2147483648:
                                        var v80 uint32 = 99
                                        _ = v80
                                    case -344:
                                        var v81 string = "a"
                                        _ = v81
                                    case 32767:
                                        for fi82 := 0; fi82 < 14; fi82++ {
                                            fmt.Println(v23, fi71)
                                            _ = fi82
                                        }
                                    }
                                }
                            case "":
                                var v83 string = ("foo bar" + "a")
                                _ = v83
                            default:
                                for fi84 := 0; fi84 < 12; fi84++ {
                                    v23 = 995
                                    var v85 uint16 = (v23 + (v23 % (((v23 << ((1) & 63))) | 1)))
                                    _ = fi84
                                    _ = v85
                                }
                            }
                            _ = fi71
                        }
                    default:
                        v86 := GenS11{F0: "", F1: v23, F2: 1}
                        _ = v86
                    }
                default:
                    var v87 bool = true
                    _ = v87
                }
            default:
                switch "" {
                case "foo bar":
                    v2 -= 153
                case "\n\t":
                    v23 = v23
                }
            }
        case "foo bar":
            var v88 float64 = 3.14159
            _ = v88
        default:
            var v89 uint = (v22 / (((v22 % ((4294967295) | 1))) | 1))
            _ = v89
        }
        _ = v68
    }
    if true {
        v22 = v22
    } else {
        v2 = (v2 >> ((942) & 63))
        fmt.Println(v23, v2, v4)
        if ((!true) || false) {
            v90 := GenS12{F0: 132, F1: ""}
            fmt.Println(v3, v3, v3)
            _ = v90
        } else {
            var v91 uint = v22
            _ = v91
        }
    }
    var v92 int8 = -1
    var v93 uint64 = (v3 % ((v3) | 1))
    v94 := []int{}
    var v95 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("slicelen v94:", len(v94), v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v64
    _ = v92
    _ = v93
    _ = v94
    _ = v95
}
