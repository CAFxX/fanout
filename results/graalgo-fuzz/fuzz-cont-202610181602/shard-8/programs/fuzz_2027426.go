package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 uint32; F1 int32; F2 bool }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 int64; F1 uint; F2 uint16 }
type GenS10 struct { F0 int16; F1 int; F2 uint }
type GenS11 struct { F0 int8; F1 bool; F2 uint32 }
type GenS12 struct { F0 int8; F1 int16; F2 uint }
type GenS13 struct { F0 uint16; F1 uint; F2 bool }
type GenS14 struct { F0 uint64; F1 int16 }
type GenS15 struct { F0 uint8; F1 bool; F2 uint32 }

func main() {
    var v1 int8 = 0
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]uint32)
    var v4 uint32 = 4294967295
    var v5 float64 = 3.14159
    v5 = -2.25
    v2 -= v2
    var v6 uint32 = 2147483648
    v5 -= float64(1.5)
    v7 := GenS1{F0: 772, F1: 2147483647}
    var v8 int8 = v1
    v6 = v6
    for fi9 := 0; fi9 < 8; fi9++ {
        v2 = 467
        v10 := GenS2{F0: 351}
        _ = fi9
        _ = v10
    }
    if true {
        var v11 uint = 0
        v12 := GenS2{F0: 0}
        v13 := GenS2{F0: -2147483648}
        _ = v11
        _ = v12
        _ = v13
    } else {
        var v14 uint8 = 0
        if true {
            if true {
                var v15 int64 = -9223372036854775808
                delete(v3, 128)
                _ = v15
            } else {
                if (!(-1 <= -622)) {
                    var v16 uint32 = (v4 >> ((2) & 63))
                    _ = v16
                } else {
                    if false {
                        delete(v3, 128)
                        var v17 string = ("\n\t" + "\n\t")
                        _ = v17
                    } else {
                        if (!false) {
                            v18 := v3[-1]
                            var v19 bool = (!(-683 <= 32767))
                            if ((!true) || ((v2 | v2) != uint64(v18))) {
                                for fi20 := 0; fi20 < 5; fi20++ {
                                    if (v19 || v19) {
                                        var v21 uint32 = v6
                                        _ = v21
                                    } else {
                                        fmt.Println(v19, fi20, fi20)
                                    }
                                    v22 := []int{}
                                    for fi23 := 0; fi23 < 15; fi23++ {
                                        if v19 {
                                            if v19 {
                                                var v24 int64 = 0
                                                var v25 int16 = 1000
                                                var v26 bool = ((!((^(v2)) <= (v2 / ((v2) | 1)))) || (!v19))
                                                _ = v24
                                                _ = v25
                                                _ = v26
                                            } else {
                                                switch 9223372036854775807 {
                                                case -9223372036854775808:
                                                    v27 := GenS2{F0: -1}
                                                    _ = v27
                                                case 0:
                                                    v28 := GenS3{F0: v19, F1: 599}
                                                    _ = v28
                                                default:
                                                    switch (fi20 >> ((1) & 63)) {
                                                    case 127:
                                                        var v29 uint8 = v14
                                                        _ = v29
                                                    case -875:
                                                        if v19 {
                                                            _ = len(v3)
                                                            switch 530 {
                                                            case 0:
                                                                fmt.Println(v14, v6, v5)
                                                                _ = len(v3)
                                                            }
                                                            v30 := GenS3{F0: (903 <= 455), F1: 600}
                                                            if ((uint32(v2) == (v4 - uint32(v14))) || ((v1 - (v8 >> ((4294967295) & 63))) == v1)) {
                                                                if (uint64(v14) >= (v2 >> ((490) & 63))) {
                                                                    var v31 int64 = -958
                                                                    if (!v19) {
                                                                        v19 = (-103 <= 42)
                                                                        v8 = v1
                                                                        fmt.Println(v5, fi23, v14)
                                                                    }
                                                                    _ = v31
                                                                }
                                                                v14 = uint8(v6)
                                                                v5 = float64(-2.25)
                                                            }
                                                            switch 174 {
                                                            case 761:
                                                                if ((v8 >> ((18446744073709551615) & 63)) >= v1) {
                                                                    v32 := v3[fi23]
                                                                    switch ("\n\t" + "hello") {
                                                                    case "\n\t":
                                                                        var v33 int8 = (-(v8))
                                                                        _ = v33
                                                                    case "hello":
                                                                        switch ("\n\t" + ("" + "")) {
                                                                        case "a":
                                                                            fmt.Println(v8, v6)
                                                                        case "hello":
                                                                            fmt.Println(fi23, v8)
                                                                        default:
                                                                            if v19 {
                                                                                v18 = (v18 | v32)
                                                                                v34 := GenS4{F0: (v6 % (((v4 - (v4 + v32))) | 1)), F1: 7}
                                                                                _ = v34
                                                                            } else {
                                                                                if false {
                                                                                    v8 -= (v8 << ((697) & 63))
                                                                                    _ = len(v22)
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v32
                                                                }
                                                            case 128:
                                                                v2 -= uint64(v14)
                                                            default:
                                                                v18 -= v18
                                                            }
                                                            fmt.Println(fi20)
                                                            _ = v30
                                                        }
                                                    default:
                                                        var v35 float32 = 1e10
                                                        _ = v35
                                                    }
                                                }
                                                v19 = ((v1 < 1) && v19)
                                                if false {
                                                    switch (-(fi23)) {
                                                    case -476:
                                                        switch 42 {
                                                        case -1:
                                                            if (false && false) {
                                                                var v36 uint8 = 65
                                                                _ = v36
                                                            }
                                                        case -447:
                                                            if v19 {
                                                                if (0.1 > float32((0.1 / v5))) {
                                                                    if len(v22) >= 2 {
                                                                        v22 = v22[1:len(v22)]
                                                                    }
                                                                }
                                                                v18 = v6
                                                                var v37 int64 = 429
                                                                _ = v37
                                                            } else {
                                                                var v38 int16 = -1
                                                                _ = v38
                                                            }
                                                        case 127:
                                                            switch "a" {
                                                            case "foo bar":
                                                                var v39 float32 = 1e10
                                                                _ = v39
                                                            default:
                                                                if (254 >= -1) {
                                                                    switch "a" {
                                                                    case "a":
                                                                        switch fi20 {
                                                                        case 127:
                                                                            delete(v3, (fi23 & (fi23 / (((-(fi23))) | 1))))
                                                                            v3[-2147483648] = (v18 * (^(v18)))
                                                                            var v40 uint8 = v14
                                                                            _ = v40
                                                                        case 9223372036854775807:
                                                                            if (v19 || true) {
                                                                                v22 = append(v22, 469)
                                                                                v41 := GenS5{F0: (v2 | (v2 * (v2 | v2))), F1: v2, F2: (v4 * v4)}
                                                                                switch fi20 {
                                                                                case 42:
                                                                                    v42 := GenS6{F0: 755, F1: (v14 + 128)}
                                                                                    _ = v42
                                                                                case -139:
                                                                                    fmt.Println(v14)
                                                                                case 44:
                                                                                    _ = len(v3)
                                                                                    fmt.Println(v5, v14, v5)
                                                                                    var v43 bool = false
                                                                                    v44 := GenS7{F0: 785, F1: 42, F2: (("hello" != ("" + "foo bar")) && v43)}
                                                                                    _ = v43
                                                                                    _ = v44
                                                                                default:
                                                                                    switch "" {
                                                                                    case "hello":
                                                                                        v19 = ((v19 || (!(0 >= 925))) || (float32(v5) > -2.25))
                                                                                    case "a":
                                                                                        v45 := GenS8{F0: 212, F1: v14}
                                                                                        v4 -= 0
                                                                                        v5 = -2.25
                                                                                        v46 := GenS9{F0: 1, F1: 2, F2: 0}
                                                                                        v18 += v4
                                                                                        v3[fi20] = 2
                                                                                        var v47 int8 = v8
                                                                                        _ = v45
                                                                                        _ = v46
                                                                                        _ = v47
                                                                                    default:
                                                                                        if false {
                                                                                            v48 := GenS10{F0: 42, F1: fi20, F2: 955}
                                                                                            v49 := GenS3{F0: (1.5 > 0.1), F1: 942}
                                                                                            _ = v48
                                                                                            _ = v49
                                                                                        }
                                                                                    }
                                                                                }
                                                                                fmt.Println(v8, v18, v19)
                                                                                _ = v41
                                                                            }
                                                                        default:
                                                                            v50 := GenS1{F0: 359, F1: fi23}
                                                                            _ = v50
                                                                        }
                                                                    case "hello":
                                                                        switch ("foo bar" + "a") {
                                                                        case "":
                                                                            fmt.Println(fi20)
                                                                        case "hello":
                                                                            delete(v3, fi20)
                                                                        default:
                                                                            var v51 uint8 = (v14 / (((v14 * 92)) | 1))
                                                                            _ = v51
                                                                        }
                                                                    default:
                                                                        var v52 bool = v19
                                                                        _ = v52
                                                                    }
                                                                    v53 := GenS7{F0: v18, F1: 176, F2: v19}
                                                                    if false {
                                                                        switch "" {
                                                                        case "hello":
                                                                            var v54 bool = ((((1.5 > -2.25) && (-881 == 42)) || false) || v19)
                                                                            _ = v54
                                                                        }
                                                                        var v55 string = ("\n\t" + "a")
                                                                        v56 := GenS4{F0: (v6 / ((v6) | 1)), F1: 0}
                                                                        var v57 uint32 = (v18 >> ((2) & 63))
                                                                        v57 = (v57 | (v18 / ((v57) | 1)))
                                                                        v58 := GenS11{F0: (v1 + (v8 * 41)), F1: (fi20 >= fi23), F2: v6}
                                                                        _ = len(v22)
                                                                        _ = v55
                                                                        _ = v56
                                                                        _ = v57
                                                                        _ = v58
                                                                    }
                                                                    var v59 int64 = 9223372036854775807
                                                                    _ = v53
                                                                    _ = v59
                                                                } else {
                                                                    var v60 uint64 = (v2 * (v2 * (v2 << ((18446744073709551615) & 63))))
                                                                    if ("foo bar" <= "a") {
                                                                        if false {
                                                                            v22 = append(v22, fi20)
                                                                            fmt.Println(v18)
                                                                            if len(v22) >= 2 {
                                                                                v22 = v22[1:len(v22)]
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(v6, v14, v6)
                                                                    _ = v60
                                                                }
                                                            }
                                                        default:
                                                            var v61 bool = (v19 || (1 < uint16(v2)))
                                                            _ = v61
                                                        }
                                                        v62 := GenS2{F0: 1}
                                                        v19 = (v19 || (!v19))
                                                        _ = v62
                                                    case 32768:
                                                        switch (((("a" + "foo bar") + ("hello" + ("" + "\n\t"))) + "foo bar") + ("" + "hello")) {
                                                        case "a":
                                                            _ = len(v22)
                                                        case "foo bar":
                                                            v19 = (!v19)
                                                        default:
                                                            var v63 string = ""
                                                            _ = v63
                                                        }
                                                    case 0:
                                                        v5 += 1.5
                                                    default:
                                                        v6 = v18
                                                    }
                                                    delete(v3, -566)
                                                }
                                            }
                                            v5 -= v5
                                            v64 := GenS12{F0: (v8 ^ v1), F1: 1, F2: 42}
                                            _ = v64
                                        } else {
                                            v4 -= (v18 & 601)
                                        }
                                        v4 = v6
                                        _ = fi23
                                    }
                                    _ = fi20
                                    _ = v22
                                }
                                for fi65 := 0; fi65 < 4; fi65++ {
                                    for fi66 := 0; fi66 < 19; fi66++ {
                                        switch fi66 {
                                        case -2147483648:
                                            if ((-(fi66)) < 4294967295) {
                                                switch "a" {
                                                case "\n\t":
                                                    switch ((("foo bar" + "a") + ("" + "a")) + "") {
                                                    case "":
                                                        v14 += (v14 << ((2) & 63))
                                                    }
                                                case "":
                                                    if (v19 || false) {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v5 -= (((v5 * 3.14159) + (v5 * float64(3.14159))) - v5)
                                                        case "a":
                                                            if v19 {
                                                                fmt.Println(v2)
                                                            } else {
                                                                v8 -= 1
                                                            }
                                                        default:
                                                            v18 = v6
                                                        }
                                                        v2 -= (v2 % ((v2) | 1))
                                                        switch 306 {
                                                        case 0:
                                                            var v67 int64 = 42
                                                            _ = v67
                                                        case -2147483648:
                                                            v4 = (v18 ^ 0)
                                                        }
                                                    } else {
                                                        v5 -= float64(0.1)
                                                        if (42 < -672) {
                                                            v68 := []int{}
                                                            var v69 int64 = int64(fi65)
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                    }
                                                case "hello":
                                                    var v70 float64 = ((v5 / 0.1) / v5)
                                                    _ = v70
                                                default:
                                                    v71 := GenS6{F0: int16(v8), F1: (v14 & v14)}
                                                    _ = v71
                                                }
                                                v5 = 1.5
                                                v14 -= (^(v14))
                                            } else {
                                                if v19 {
                                                    fmt.Println(v18)
                                                } else {
                                                    var v72 int64 = -407
                                                    var v73 uint8 = (v14 >> ((1) & 63))
                                                    _ = v72
                                                    _ = v73
                                                }
                                                delete(v3, fi65)
                                                v2 = v2
                                                v74 := GenS9{F0: 982, F1: 0, F2: 65535}
                                                if false {
                                                    switch ("\n\t" + "foo bar") {
                                                    case "\n\t":
                                                        if v19 {
                                                            if (!(0 > -2147483648)) {
                                                                v75 := GenS13{F0: 0, F1: 42, F2: (0.0 > 1e10)}
                                                                fmt.Println(v1, v8, fi65)
                                                                var v76 int16 = 1
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                        }
                                                        v77 := GenS6{F0: 42, F1: (^(v14))}
                                                        _ = v77
                                                    case "hello":
                                                        v1 += int8(v6)
                                                    default:
                                                        fmt.Println(v18, v6)
                                                    }
                                                } else {
                                                    fmt.Println(v1, fi66, v14)
                                                    v78 := GenS5{F0: v2, F1: (^(v2)), F2: 42}
                                                    v79 := []int{}
                                                    var v80 int = fi66
                                                    fmt.Println(fi66, v14)
                                                    switch -486 {
                                                    case 5:
                                                        var v81 uint32 = v6
                                                        _ = v81
                                                    }
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                }
                                                fmt.Println(v6)
                                                _ = v74
                                            }
                                        }
                                        v82 := GenS14{F0: v2, F1: -32768}
                                        var v83 uint8 = v14
                                        _ = fi66
                                        _ = v82
                                        _ = v83
                                    }
                                    for fi84 := 0; fi84 < 7; fi84++ {
                                        v85 := []int{(fi65 / (((fi65 ^ fi65)) | 1)), fi84}
                                        _ = fi84
                                        _ = v85
                                    }
                                    _ = fi65
                                }
                            } else {
                                var v86 int8 = v8
                                v87 := []int{32767, 4294967295, 2147483647, 4294967295}
                                _ = v86
                                _ = v87
                            }
                            _ = v18
                            _ = v19
                        } else {
                            var v88 bool = true
                            _ = v88
                        }
                        fmt.Println(v4, v5, v1)
                        v8 -= (v8 / (((v8 >> ((42) & 63))) | 1))
                    }
                    v8 += (v8 + v8)
                    v89 := GenS7{F0: v4, F1: 42, F2: true}
                    _ = v89
                }
                _ = len(v3)
                v90 := GenS11{F0: v8, F1: false, F2: 2147483647}
                _ = v90
            }
            var v91 uint32 = v6
            v92 := v3[119]
            _ = v91
            _ = v92
        }
        _ = v14
    }
    fmt.Println(v6, v1)
    v8 = v1
    v93 := []int{}
    v94 := GenS15{F0: 28, F1: true, F2: 2}
    fmt.Println(v4, v1)
    fmt.Println(v8)
    v95 := GenS5{F0: v2, F1: v2, F2: 454}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v93:", len(v93), v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v93
    _ = v94
    _ = v95
}
