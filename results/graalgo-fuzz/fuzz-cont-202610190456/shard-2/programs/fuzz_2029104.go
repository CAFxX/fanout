package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int; F1 uint8; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 string; F1 int64 }
func genFn0(p0 int32, p1 uint16, p2 uint8) int32 {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{(^(fi6))}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return (p0 & -1)
}
func genFn1(p0 int8, p1 uint16, p2 uint16) uint {
    var v8 int16 = 32767
    if (0 <= -624) {
        fmt.Println(v8, p2, p1)
    } else {
        v9 := []int{42, 4294967295, -235, 32768}
        fmt.Println(p2, p0, p2)
        _ = v9
    }
    if (false && false) {
        p0 += (p0 / ((127) | 1))
        v10 := []int{int(v8)}
        switch (("hello" + ("foo bar" + "hello")) + "hello") {
        case "a":
            if (false && ("" == ("hello" + "foo bar"))) {
                var v11 int8 = (p0 + (-(p0)))
                var v12 uint64 = 0
                switch "foo bar" {
                case "\n\t":
                    var v13 int16 = v8
                    _ = v13
                case "":
                    if len(v10) > 0 {
                        v10[0] = 1
                    }
                }
                _ = v11
                _ = v12
            }
        case "foo bar":
            v14 := GenS1{F0: v8, F1: 127}
            _ = v14
        }
        _ = v10
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS2{F0: (-(fi15)), F1: 255, F2: p0}
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 bool = (981 < p2)
            _ = fi17
            _ = v18
        }
        switch (fi15 + fi15) {
        case 300:
            p2 = p1
        case 249:
            var v19 float32 = 1.5
            _ = v19
        case -689:
            switch (^(fi15)) {
            case 32768:
                var v20 float64 = float64(-2.25)
                _ = v20
            case 779:
                var v21 uint8 = 1
                _ = v21
            case -2147483648:
                for fi22 := 0; fi22 < 1; fi22++ {
                    switch "hello" {
                    case "a":
                        if (917 > -9223372036854775808) {
                            v23 := GenS1{F0: -746, F1: fi22}
                            p1 = p2
                            if (p0 < (p0 - 0)) {
                                v24 := GenS3{F0: (-(p0))}
                                v25 := GenS2{F0: fi22, F1: uint8(p2), F2: p0}
                                var v26 bool = (fi22 >= 1)
                                v27 := GenS3{F0: (p0 ^ p0)}
                                if true {
                                    fmt.Println(p1)
                                }
                                _ = v24
                                _ = v25
                                _ = v26
                                _ = v27
                            } else {
                                if (!true) {
                                    fmt.Println(fi15)
                                    p1 -= 116
                                    p1 = 1
                                }
                                if (!(v8 == v8)) {
                                    p0 += p0
                                    var v28 int = (^(fi22))
                                    _ = v28
                                } else {
                                    v29 := []int{-2147483648, fi15}
                                    var v30 int64 = -51
                                    if len(v29) >= 2 {
                                        v29 = v29[1:len(v29)]
                                    }
                                    _ = v29
                                    _ = v30
                                }
                                v31 := GenS2{F0: (-(fi15)), F1: 2, F2: (p0 / ((127) | 1))}
                                _ = v31
                            }
                            v32 := []int{fi22, (fi22 << ((1) & 63)), (^(fi15)), fi15}
                            p2 -= p1
                            v33 := GenS3{F0: p0}
                            _ = v23
                            _ = v32
                            _ = v33
                        } else {
                            if true {
                                switch 128 {
                                case 32767:
                                    v34 := GenS4{F0: p0}
                                    _ = v34
                                case -1:
                                    switch ("foo bar" + "a") {
                                    case "hello":
                                        if (!true) {
                                            var v35 int16 = v8
                                            switch fi15 {
                                            case 0:
                                                switch int(v35) {
                                                case 793:
                                                    if (42 > genFn0(2147483647, p1, 127)) {
                                                        p2 -= (p1 >> ((4294967295) & 63))
                                                        var v36 int = fi22
                                                        _ = v36
                                                    }
                                                case 0:
                                                    p0 += p0
                                                default:
                                                    if (42 >= 4294967295) {
                                                        if (!(!(237 <= 1))) {
                                                            var v37 bool = true
                                                            var v38 uint8 = 125
                                                            v39 := GenS3{F0: 100}
                                                            var v40 uint = 4294967295
                                                            _ = v37
                                                            _ = v38
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                    } else {
                                                        var v41 uint8 = 0
                                                        if (true || true) {
                                                            var v42 float32 = 1e10
                                                            p1 -= (p2 / (((p1 + (^(p1)))) | 1))
                                                            if false {
                                                                var v43 uint = 4294967295
                                                                _ = v43
                                                            }
                                                            p0 = (p0 ^ (-(p0)))
                                                            v44 := GenS2{F0: fi15, F1: (^(v41)), F2: 96}
                                                            var v45 int64 = 1
                                                            if (!false) {
                                                                fmt.Println(v45)
                                                                p0 -= (^(p0))
                                                            }
                                                            _ = v42
                                                            _ = v44
                                                            _ = v45
                                                        } else {
                                                            fmt.Println(fi15, fi15)
                                                            v46 := []int{-67}
                                                            switch ("foo bar" + ("" + "\n\t")) {
                                                            case "\n\t":
                                                                var v47 float64 = (float64(-2.25) * 0.0)
                                                                _ = v47
                                                            default:
                                                                if false {
                                                                    var v48 float32 = (-2.25 - ((-2.25 + float32(((-2.25 * 0.0) * (float64(float32(0.1)) - 1.5)))) * (3.14159 / 1.5)))
                                                                    v46 = append(v46, fi15)
                                                                    var v49 uint8 = 0
                                                                    var v50 string = ((("" + ("\n\t" + "a")) + "\n\t") + "a")
                                                                    _ = v48
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                            }
                                                            switch (("a" + ("a" + "hello")) + ("" + "foo bar")) {
                                                            case "\n\t":
                                                                p2 = 492
                                                            }
                                                            fmt.Println(p0)
                                                            _ = v46
                                                        }
                                                        _ = v41
                                                    }
                                                }
                                            case 687:
                                                if true {
                                                    var v51 uint32 = 65535
                                                    fmt.Println(p1)
                                                    var v52 int32 = 0
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    v53 := []int{0, (-(fi22)), (-(fi22))}
                                                    var v54 bool = false
                                                    if len(v53) >= 2 {
                                                        v53 = v53[1:len(v53)]
                                                    }
                                                    var v55 uint16 = p2
                                                    _ = v53
                                                    _ = v54
                                                    _ = v55
                                                }
                                            case -9223372036854775808:
                                                var v56 uint = 4294967295
                                                _ = v56
                                            }
                                            var v57 int64 = -646
                                            _ = v35
                                            _ = v57
                                        } else {
                                            switch (("" + "foo bar") + "") {
                                            case "\n\t":
                                                var v58 uint32 = 42
                                                _ = v58
                                            default:
                                                switch "a" {
                                                case "\n\t":
                                                    fmt.Println(v8)
                                                    var v59 uint8 = 2
                                                    v60 := []int{}
                                                    _ = v59
                                                    _ = v60
                                                }
                                            }
                                            p1 = (p1 >> ((600) & 63))
                                            var v61 uint16 = p1
                                            _ = v61
                                        }
                                    case "a":
                                        v62 := GenS2{F0: (fi22 << ((18446744073709551615) & 63)), F1: 42, F2: p0}
                                        _ = v62
                                    case "":
                                        v63 := GenS1{F0: (v8 | 32767), F1: fi15}
                                        _ = v63
                                    default:
                                        v64 := GenS5{F0: (-(v8)), F1: ("" + ("foo bar" + "a"))}
                                        _ = v64
                                    }
                                default:
                                    p1 = (^(p1))
                                }
                                fmt.Println(p1, p2)
                                fmt.Println(p2, p2)
                                v65 := GenS2{F0: fi15, F1: 0, F2: (p0 % ((p0) | 1))}
                                p2 += p2
                                _ = v65
                            } else {
                                var v66 uint = 65535
                                if false {
                                    p1 += (p1 + (p1 % ((42) | 1)))
                                    switch "a" {
                                    case "a":
                                        p0 -= p0
                                        v67 := []int{fi22, (-(fi22))}
                                        _ = v67
                                    case "hello":
                                        p1 = (p2 & 2)
                                    default:
                                        v66 = v66
                                    }
                                } else {
                                    v8 = (v8 % ((-32768) | 1))
                                    if ((p1 * (p1 << (((^(v66))) & 63))) <= 32767) {
                                        switch "\n\t" {
                                        case "hello":
                                            v66 = v66
                                        case "\n\t":
                                            var v68 uint16 = p2
                                            _ = v68
                                        case "foo bar":
                                            v69 := []int{}
                                            _ = v69
                                        default:
                                            p0 = p0
                                        }
                                        var v70 uint64 = 1
                                        _ = v70
                                    }
                                    var v71 float32 = (3.14159 - 1e10)
                                    _ = v71
                                }
                                var v72 uint = v66
                                _ = v66
                                _ = v72
                            }
                            switch "foo bar" {
                            case "\n\t":
                                var v73 int64 = 42
                                _ = v73
                            default:
                                switch 2147483647 {
                                case 42:
                                    v74 := GenS6{F0: 0, F1: 332}
                                    _ = v74
                                case 763:
                                    switch ("a" + "foo bar") {
                                    case "":
                                        if ((42 < -456) || (4 < 42)) {
                                            switch (fi15 << ((4294967295) & 63)) {
                                            case -366:
                                                v8 = v8
                                            }
                                            p1 = 660
                                        }
                                    case "\n\t":
                                        if (!true) {
                                            p1 = 42
                                            v75 := []int{32767, (fi15 & fi15), (fi22 * fi15)}
                                            if len(v75) >= 2 {
                                                v75 = v75[1:len(v75)]
                                            }
                                            _ = v75
                                        } else {
                                            if (false || (false && (-2.25 > float64((1.5 + float32(float64(float32(-2.25)))))))) {
                                                if false {
                                                    p2 -= p2
                                                    p1 -= p2
                                                    v76 := []int{(fi15 | (fi15 >> ((705) & 63))), 0, fi15}
                                                    _ = v76
                                                }
                                                var v77 string = "hello"
                                                _ = v77
                                            }
                                        }
                                    case "foo bar":
                                        var v78 bool = (66 <= 0)
                                        _ = v78
                                    default:
                                        v8 = (^(v8))
                                        var v79 int8 = (p0 - -128)
                                        _ = v79
                                    }
                                    var v80 int = 700
                                    if false {
                                        v81 := []int{888, -65}
                                        v82 := GenS1{F0: 353, F1: fi22}
                                        _ = v81
                                        _ = v82
                                    }
                                    _ = v80
                                }
                            }
                            var v83 int32 = genFn0(2147483647, (p1 / ((65535) | 1)), 255)
                            _ = v83
                        }
                    default:
                        var v84 uint16 = (p1 * (p2 & p2))
                        _ = v84
                    }
                    if false {
                        if true {
                            p2 -= p2
                        } else {
                            switch fi15 {
                            case 127:
                                p2 = 888
                            case 759:
                                var v85 int = -9223372036854775808
                                _ = v85
                            default:
                                if (true || false) {
                                    v86 := []int{}
                                    var v87 float32 = 3.14159
                                    p2 = 2
                                    _ = v86
                                    _ = v87
                                } else {
                                    fmt.Println(p0)
                                }
                            }
                            var v88 float64 = 0.0
                            v89 := []int{fi22, fi15, fi22, fi22}
                            var v90 float64 = (1e10 * (1e10 * 1.5))
                            var v91 uint16 = (p2 * p1)
                            _ = v88
                            _ = v89
                            _ = v90
                            _ = v91
                        }
                    }
                    p0 = (p0 + (p0 * (p0 + (p0 << ((970) & 63)))))
                    _ = fi22
                }
            }
        default:
            switch fi15 {
            case 128:
                v92 := []int{fi15, (fi15 << ((144) & 63))}
                _ = v92
            case -886:
                var v93 uint = 4294967295
                _ = v93
            }
        }
        _ = fi15
        _ = v16
    }
    v94 := GenS5{F0: 78, F1: "hello"}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v94
    return 1
}

func main() {
    v1 := []uint32{1}
    var v2 int8 = 79
    v3 := make(map[int]int)
    var v4 int32 = 42
    var v5 uint32 = 1
    for fi95 := 0; fi95 < 4; fi95++ {
        var v96 int16 = -506
        v97 := GenS7{F0: -9223372036854775808, F1: 267}
        fmt.Println(v5, v5, v2)
        _ = fi95
        _ = v96
        _ = v97
    }
    fmt.Println(v4)
    if (((v2 << ((genFn1(v2, 65535, 65535)) & 63)) >= (^(v2))) || true) {
        switch "foo bar" {
        case "a":
            if false {
                v98 := GenS8{F0: "", F1: -820}
                _ = v98
            } else {
                v4 = (v4 * v4)
            }
        case "hello":
            fmt.Println(v4, v5, v4)
        }
        v2 -= (v2 * (^(v2)))
        for fi99 := 0; fi99 < 8; fi99++ {
            if len(v1) > 0 {
                v1[0] = v5
            }
            var v100 uint32 = (v5 << ((genFn1(v2, 32767, 1)) & 63))
            _ = fi99
            _ = v100
        }
    } else {
        fmt.Println(v4, v4, v4)
    }
    switch 717 {
    case 2147483647:
        v2 -= v2
    case -1:
        if (!true) {
            v5 = v5
        } else {
            var v101 float32 = 1.5
            switch (("a" + "foo bar") + "hello") {
            case "\n\t":
                _ = len(v3)
            case "foo bar":
                switch 466 {
                case -2147483648:
                    var v102 uint8 = 73
                    _ = v102
                case 2147483647:
                    for fi103 := 0; fi103 < 1; fi103++ {
                        var v104 float32 = v101
                        _ = fi103
                        _ = v104
                    }
                }
            }
            _ = v101
        }
    default:
        if (int32(v2) >= genFn0(158, 0, 255)) {
            v105 := GenS4{F0: 1}
            _ = v105
        }
    }
    fmt.Println(v2)
    var v106 int8 = 1
    for fi107 := 0; fi107 < 4; fi107++ {
        v5 -= v5
        v2 = (v2 + -118)
        _ = fi107
    }
    v108 := GenS3{F0: (^(v106))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v106
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
