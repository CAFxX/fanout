package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 string; F1 int32; F2 bool }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 int; F1 int8; F2 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int64; F1 int64; F2 int8 }
type GenS10 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS11 struct { F0 int8; F1 uint; F2 int }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64; F1 int; F2 uint }
func genFn0(p0 int64, p1 int16, p2 uint32) int {
    var v6 int64 = p0
    var v7 uint64 = 508
    var v8 uint = 18446744073709551615
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int = fi9
        var v11 int = 9223372036854775807
        switch "foo bar" {
        case "":
            var v12 uint8 = 128
            _ = v12
        case "hello":
            fmt.Println(v10, p2, p2)
        default:
            p2 = 1
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 uint64) int8 {
    v13 := GenS1{F0: (!false), F1: 128, F2: 2147483647}
    var v14 uint64 = (p0 >> ((42) & 63))
    if ((1e10 / (-2.25 * 1.5)) <= (1e10 - 3.14159)) {
        if false {
            for fi15 := 0; fi15 < 20; fi15++ {
                v16 := GenS2{F0: 1, F1: 851}
                v17 := []int{fi15, fi15, fi15, (fi15 * (^(fi15)))}
                v18 := GenS2{F0: 100, F1: -593}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            var v19 uint32 = 259
            _ = v19
        } else {
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := []int{(fi20 >> ((18446744073709551615) & 63)), fi20, (fi20 / ((0) | 1)), fi20}
                var v22 string = "hello"
                _ = fi20
                _ = v21
                _ = v22
            }
            v14 -= p0
            v14 += (p0 % ((0) | 1))
        }
    } else {
        v23 := []int{545, 162, -2147483648, 0}
        v23 = append(v23, -2147483648)
        _ = v23
    }
    _ = p0
    _ = v13
    _ = v14
    return 83
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]int)
    var v4 uint32 = 4294967295
    v5 := make(map[uint32]bool)
    var v24 string = ("\n\t" + (("" + ("hello" + (("a" + "a") + "a"))) + "foo bar"))
    v4 = 0
    v25 := v5[(v4 & (v4 - (v4 - (v4 | (^(v4))))))]
    for fi26 := 0; fi26 < 18; fi26++ {
        for fi27 := 0; fi27 < 14; fi27++ {
            var v28 float32 = ((1.5 / 1.0) / (3.14159 + 1.5))
            _ = fi27
            _ = v28
        }
        if (!(477 != 0)) {
            if false {
                for fi29 := 0; fi29 < 4; fi29++ {
                    var v30 int16 = -487
                    var v31 float64 = -2.25
                    var v32 float64 = (1e10 + ((-2.25 * 0.1) - v31))
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                for fi33 := 0; fi33 < 8; fi33++ {
                    switch (fi33 >> ((0) & 63)) {
                    case 9223372036854775807:
                        v25 = v25
                    case -335:
                        v24 = v24
                        v34 := []int{9223372036854775807}
                        var v35 uint64 = v2
                        _ = v34
                        _ = v35
                    case 456:
                        var v36 int8 = -98
                        _ = v36
                    }
                    v24 = v24
                    switch fi33 {
                    case 32768:
                        v2 += (v2 << ((175) & 63))
                        if (fi33 <= fi33) {
                            v25 = v25
                            fmt.Println(v4)
                            if v25 {
                                v25 = ((!(!(v25 && (fi26 > genFn0(930, 32767, v4))))) && v25)
                            }
                        } else {
                            switch genFn0(-124, 124, v4) {
                            case 32767:
                                fmt.Println(v4, fi33)
                            case 42:
                                var v37 int = fi26
                                _ = v37
                            case 32768:
                                if (true || true) {
                                    v3[(v4 * 0)] = -652
                                } else {
                                    if v25 {
                                        switch v24 {
                                        case "foo bar":
                                            v38 := []int{(fi33 ^ (fi33 ^ (fi26 % ((fi26) | 1)))), 42}
                                            _ = v38
                                        case "":
                                            switch v24 {
                                            case "\n\t":
                                                v39 := v3[(v4 >> ((0) & 63))]
                                                _ = v39
                                            case "":
                                                v40 := []int{fi33, (fi33 - fi26), fi26, 32767}
                                                _ = v40
                                            default:
                                                if (42 != 250) {
                                                    v41 := v5[v4]
                                                    fmt.Println(v41, fi33)
                                                    _ = v41
                                                } else {
                                                    v2 += (^(v2))
                                                    v42 := []int{fi26, (-(fi26))}
                                                    _ = v42
                                                }
                                            }
                                            v43 := []int{32767, 42, (-(fi26))}
                                            _ = v43
                                        default:
                                            fmt.Println(v24)
                                        }
                                        switch fi33 {
                                        case 4294967295:
                                            if ((v25 && false) || true) {
                                                v44 := []int{-9223372036854775808}
                                                _ = v44
                                            } else {
                                                switch "\n\t" {
                                                case "":
                                                    v24 = (((v24 + v24) + v24) + (v24 + ("\n\t" + "")))
                                                case "\n\t":
                                                    v45 := []int{int(v4), (fi26 ^ fi33), fi26}
                                                    _ = v45
                                                case "a":
                                                    switch genFn0(897, 333, 2147483647) {
                                                    case 291:
                                                        v46 := GenS3{F0: ("foo bar" + v24), F1: 1, F2: v25}
                                                        _ = v46
                                                    case -247:
                                                        if v25 {
                                                            v24 = (v24 + v24)
                                                            if v25 {
                                                                delete(v5, v4)
                                                                if v25 {
                                                                    if v25 {
                                                                        v4 += v4
                                                                    } else {
                                                                        fmt.Println(v2, fi33)
                                                                    }
                                                                    v2 = v2
                                                                } else {
                                                                    var v47 int = -2147483648
                                                                    switch ((v24 + "a") + v24) {
                                                                    case "foo bar":
                                                                        v48 := GenS3{F0: v24, F1: 42, F2: v25}
                                                                        if (v2 != (v2 & (v2 - (v2 * 0)))) {
                                                                            fmt.Println(v47, v2, fi26)
                                                                            v47 = genFn0(9223372036854775807, -32768, 4294967295)
                                                                        }
                                                                        v49 := GenS4{F0: (fi26 * genFn0(-1, -622, v4)), F1: 171}
                                                                        var v50 float64 = 0.1
                                                                        _ = v48
                                                                        _ = v49
                                                                        _ = v50
                                                                    default:
                                                                        fmt.Println(v25, v47)
                                                                    }
                                                                    _ = v47
                                                                }
                                                            }
                                                            if v25 {
                                                                var v51 float64 = ((1.5 * -2.25) * float64(1.5))
                                                                _ = v51
                                                            }
                                                        }
                                                    case 32768:
                                                        v52 := []int{(fi26 + (^(fi33))), (fi26 / ((genFn0(9223372036854775807, -1, (v4 << ((18446744073709551615) & 63)))) | 1)), -2147483648}
                                                        _ = v52
                                                    default:
                                                        v53 := GenS5{F0: 1, F1: 128}
                                                        _ = v53
                                                    }
                                                default:
                                                    var v54 int8 = 127
                                                    _ = v54
                                                }
                                                var v55 uint = 18446744073709551615
                                                _ = v55
                                            }
                                        }
                                        v56 := []int{fi26, genFn0(1, 1000, uint32(v2)), (fi33 + (^(fi26)))}
                                        switch "hello" {
                                        case "foo bar":
                                            v4 += (v4 % ((v4) | 1))
                                            v25 = true
                                            var v57 int8 = 100
                                            _ = v57
                                        case "":
                                            v24 = v24
                                        }
                                        _ = v56
                                    }
                                    switch (fi33 + -9223372036854775808) {
                                    case 776:
                                        var v58 float64 = 0.1
                                        _ = v58
                                    case -9223372036854775808:
                                        var v59 float32 = 1.5
                                        _ = v59
                                    case 9223372036854775807:
                                        v60 := []int{}
                                        fmt.Println(v24, v4, v4)
                                        var v61 string = (("hello" + v24) + v24)
                                        _ = v60
                                        _ = v61
                                    default:
                                        var v62 uint = 582
                                        _ = v62
                                    }
                                    v2 -= v2
                                    v63 := GenS6{F0: fi26, F1: 99, F2: false}
                                    v1["a"] = false
                                    var v64 uint = 18446744073709551615
                                    _ = v63
                                    _ = v64
                                }
                            }
                            v65 := GenS4{F0: (fi26 / ((1) | 1)), F1: 1}
                            _ = v65
                        }
                        v66 := GenS5{F0: 128, F1: 179}
                        _ = v66
                    case -2147483648:
                        switch v24 {
                        case "":
                            delete(v3, v4)
                        case "\n\t":
                            fmt.Println(v4, v2, v24)
                        case "a":
                            v67 := v3[v4]
                            fmt.Println(v2, fi26, v2)
                            var v68 float32 = -2.25
                            if ((!true) || (!true)) {
                                var v69 int64 = 42
                                var v70 int32 = 42
                                if ((fi33 & genFn0(-401, 32767, (v4 % ((65535) | 1)))) != (v67 & v67)) {
                                    v71 := []int{(fi33 - fi33), fi33, (v67 >> ((1) & 63)), (^(fi26))}
                                    if v25 {
                                        v5[(v4 | v4)] = true
                                    } else {
                                        v67 += (fi33 << ((4294967295) & 63))
                                    }
                                    _ = v71
                                } else {
                                    v72 := GenS3{F0: v24, F1: (v70 / (((v70 | (v70 / ((2147483647) | 1)))) | 1)), F2: true}
                                    _ = v72
                                }
                                var v73 uint16 = 222
                                if v25 {
                                    switch 42 {
                                    case 1:
                                        switch 32768 {
                                        case 2147483647:
                                            var v74 uint = 18446744073709551615
                                            _ = v74
                                        default:
                                            if (v25 || (2 != 1)) {
                                                var v75 float32 = 0.1
                                                if false {
                                                    v76 := []int{fi33, int(v73)}
                                                    _ = v76
                                                }
                                                _ = v75
                                            } else {
                                                v73 = v73
                                                var v77 float32 = 0.1
                                                var v78 bool = (!true)
                                                _ = v77
                                                _ = v78
                                            }
                                        }
                                    case -2147483648:
                                        v2 += 547
                                    default:
                                        fmt.Println(v4)
                                    }
                                    switch v24 {
                                    case "hello":
                                        var v79 uint32 = 4294967295
                                        _ = v79
                                    case "foo bar":
                                        v80 := GenS7{F0: v73}
                                        _ = v80
                                    default:
                                        var v81 int8 = -59
                                        _ = v81
                                    }
                                    if v25 {
                                        fmt.Println(v70, v67)
                                        fmt.Println(v70, fi33)
                                    }
                                }
                                v82 := []int{(fi33 >> ((4294967295) & 63)), -1}
                                _ = v69
                                _ = v70
                                _ = v73
                                _ = v82
                            } else {
                                v83 := []int{9223372036854775807, 444}
                                v84 := GenS8{F0: 658}
                                v68 -= v68
                                _ = v83
                                _ = v84
                            }
                            var v85 uint8 = 128
                            _ = v67
                            _ = v68
                            _ = v85
                        default:
                            var v86 string = "hello"
                            _ = v86
                        }
                    case 1:
                        v24 = v24
                    default:
                        v2 = v2
                    }
                    if v25 {
                        v24 = v24
                        fmt.Println(fi33)
                        v87 := []int{genFn0(-1, -1, (^(v4))), fi26}
                        _ = v87
                    }
                    _ = fi33
                }
                var v88 int32 = 648
                _ = v88
            } else {
                switch "a" {
                case "hello":
                    v4 -= v4
                case "\n\t":
                    v2 += (^(v2))
                }
            }
            v89 := GenS8{F0: 42}
            v2 += (v2 % (((^(v2))) | 1))
            _ = v89
        }
        v2 = (v2 * 4294967295)
        _ = fi26
    }
    for fi90 := 0; fi90 < 8; fi90++ {
        switch v24 {
        case "hello":
            var v91 uint16 = 313
            _ = v91
        case "foo bar":
            fmt.Println(fi90, v25, v25)
        case "\n\t":
            v25 = false
        }
        v92 := []int{(fi90 | 345), (fi90 << ((42) & 63)), fi90, fi90}
        v93 := GenS9{F0: 1, F1: -1, F2: 14}
        _ = fi90
        _ = v92
        _ = v93
    }
    var v94 float64 = float64(1.5)
    v25 = v25
    for fi95 := 0; fi95 < 5; fi95++ {
        if (((855 != 65535) || (0 >= v4)) && v25) {
            for fi96 := 0; fi96 < 1; fi96++ {
                switch ((("a" + (v24 + (v24 + v24))) + "\n\t") + v24) {
                case "a":
                    fmt.Println(v24, v4, v4)
                case "":
                    v97 := GenS10{F0: 9223372036854775807, F1: v2, F2: 106}
                    _ = v97
                }
                var v98 int = fi96
                fmt.Println(v4)
                _ = fi96
                _ = v98
            }
            v99 := GenS11{F0: 0, F1: 745, F2: fi95}
            _ = v99
        }
        v100 := GenS7{F0: 42}
        _ = fi95
        _ = v100
    }
    fmt.Println(v4, v4)
    if (!(-851 <= 32767)) {
        for fi101 := 0; fi101 < 10; fi101++ {
            v102 := GenS12{F0: v24}
            v24 = "\n\t"
            v25 = ((float64(1e10) / (-2.25 / v94)) > (v94 + float64(1.5)))
            _ = fi101
            _ = v102
        }
    } else {
        for fi103 := 0; fi103 < 7; fi103++ {
            var v104 int32 = 1
            _ = fi103
            _ = v104
        }
        var v105 int32 = -673
        var v106 uint32 = v4
        _ = v105
        _ = v106
    }
    var v107 uint16 = 32767
    v108 := GenS13{F0: 9223372036854775807, F1: 32767, F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v94:", v94)
    _ = v1
    _ = v107
    _ = v108
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v94
}
