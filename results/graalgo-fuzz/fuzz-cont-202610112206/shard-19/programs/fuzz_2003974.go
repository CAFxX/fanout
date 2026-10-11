package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint32; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int8; F1 int; F2 bool }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 uint8; F1 int32; F2 int32 }
type GenS10 struct { F0 bool; F1 uint64; F2 uint8 }
func genFn0(p0 uint32) uint8 {
    var v6 uint = 948
    v6 -= 4294967295
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 string = ""
        switch fi7 {
        case 9223372036854775807:
            if (false && (-255 <= -1)) {
                for fi9 := 0; fi9 < 13; fi9++ {
                    v10 := GenS1{F0: 32767}
                    _ = fi9
                    _ = v10
                }
                v11 := GenS1{F0: 2}
                var v12 bool = true
                _ = v11
                _ = v12
            }
        default:
            var v13 uint16 = 1
            _ = v13
        }
        fmt.Println(v8, v8)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return 255
}

func main() {
    v1 := []int{128, -2147483648, 9223372036854775807, 42}
    v2 := make(map[uint32]uint32)
    var v3 int = 32767
    v4 := []int{(v3 & 32767), (v3 << ((1) & 63))}
    v5 := make(map[uint32]uint32)
    for fi14 := 0; fi14 < 7; fi14++ {
        v3 = v3
        _ = fi14
    }
    var v15 uint32 = 598
    if true {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 bool = false
            for fi18 := 0; fi18 < 3; fi18++ {
                v19 := GenS1{F0: 255}
                v20 := v2[(v15 >> ((18446744073709551615) & 63))]
                var v21 uint = 2
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            _ = fi16
            _ = v17
        }
        v15 = 1
        var v22 bool = false
        _ = v22
    } else {
        if (false && (100 >= -1)) {
            for fi23 := 0; fi23 < 18; fi23++ {
                if (!(56 == 2)) {
                    v3 = (fi23 + (^(fi23)))
                    for fi24 := 0; fi24 < 1; fi24++ {
                        var v25 int8 = 0
                        fmt.Println(fi24, v15, fi23)
                        fmt.Println(fi23, v15, v15)
                        _ = fi24
                        _ = v25
                    }
                    if false {
                        v15 -= 0
                        switch (fi23 & 2147483647) {
                        case -490:
                            v26 := GenS1{F0: 1}
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 11; fi27++ {
                            if (2 < 4294967295) {
                                if false {
                                    var v28 uint64 = 42
                                    var v29 uint64 = v28
                                    v15 = (v15 | 416)
                                    _ = v28
                                    _ = v29
                                } else {
                                    if ((11 > 255) || false) {
                                        delete(v5, v15)
                                        v15 = (v15 % ((v15) | 1))
                                        if false {
                                            var v30 float64 = 1.5
                                            _ = v30
                                        } else {
                                            var v31 int8 = 42
                                            _ = v31
                                        }
                                        var v32 int16 = 42
                                        _ = v32
                                    } else {
                                        v15 = v15
                                        v33 := GenS1{F0: 0}
                                        var v34 uint32 = (v15 >> ((1) & 63))
                                        v35 := GenS2{F0: 28}
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                    switch (fi27 + fi23) {
                                    case 4294967295:
                                        v36 := v5[(v15 >> ((4294967295) & 63))]
                                        _ = v36
                                    case 32767:
                                        if (false || (1 != 42)) {
                                            if true {
                                                if true {
                                                    if true {
                                                        fmt.Println(v15, v3)
                                                        fmt.Println(v15, v3)
                                                        v3 = (fi27 & -842)
                                                    }
                                                    v15 = (v15 >> ((789) & 63))
                                                    v15 -= v15
                                                    v15 -= (v15 & v15)
                                                }
                                                v37 := v2[(v15 | 54)]
                                                switch (("foo bar" + "") + "") {
                                                case "\n\t":
                                                    var v38 int = (v3 * 4294967295)
                                                    _ = v38
                                                }
                                                _ = v37
                                            } else {
                                                var v39 int8 = 100
                                                var v40 int32 = -1
                                                switch (-(fi27)) {
                                                case -9223372036854775808:
                                                    if ((0 <= 0) || false) {
                                                        v40 -= (v40 | v40)
                                                    } else {
                                                        var v41 int = (fi23 - fi23)
                                                        var v42 int32 = (^(v40))
                                                        if len(v1) > 0 {
                                                            v1[0] = -9223372036854775808
                                                        }
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                case 582:
                                                    fmt.Println(fi27, fi23)
                                                default:
                                                    v43 := GenS1{F0: 255}
                                                    _ = v43
                                                }
                                                v3 = (fi27 | -549)
                                                fmt.Println(v39, v40)
                                                switch "\n\t" {
                                                case "":
                                                    fmt.Println(v40, v3, fi23)
                                                default:
                                                    var v44 int8 = (v39 * -126)
                                                    _ = v44
                                                }
                                                _ = v39
                                                _ = v40
                                            }
                                            v15 = (^(v15))
                                            v15 += v15
                                        }
                                        v45 := v2[(v15 + v15)]
                                        fmt.Println(v45)
                                        _ = v45
                                    }
                                }
                                v3 += 32768
                                v15 = (v15 / ((v15) | 1))
                            }
                            if false {
                                fmt.Println(fi27, fi27)
                                fmt.Println(fi23, v3, fi23)
                                fmt.Println(v3, v15)
                            } else {
                                if (!false) {
                                    var v46 uint = 141
                                    var v47 int = 128
                                    _ = v46
                                    _ = v47
                                }
                                var v48 uint8 = 176
                                _ = v48
                            }
                            v49 := GenS2{F0: genFn0((v15 / (((v15 & (^(v15)))) | 1)))}
                            _ = fi27
                            _ = v49
                        }
                    }
                } else {
                    v50 := GenS2{F0: 175}
                    for fi51 := 0; fi51 < 15; fi51++ {
                        var v52 uint = 2
                        fmt.Println(fi23)
                        _ = fi51
                        _ = v52
                    }
                    if false {
                        v15 -= v15
                    } else {
                        if false {
                            v15 = (^(v15))
                            var v53 string = "foo bar"
                            var v54 string = (v53 + (v53 + (((("hello" + ("a" + v53)) + v53) + v53) + (((v53 + "foo bar") + (v53 + v53)) + v53))))
                            _ = v53
                            _ = v54
                        } else {
                            v1 = append(v1, -432)
                        }
                        switch "a" {
                        case "hello":
                            v3 = v3
                        case "foo bar":
                            if (2 < 42) {
                                for fi55 := 0; fi55 < 1; fi55++ {
                                    if false {
                                        if ((true || true) && true) {
                                            _ = len(v2)
                                            var v56 int64 = -1
                                            _ = v56
                                        } else {
                                            v15 = 42
                                            v3 = v3
                                            v15 = v15
                                            switch "a" {
                                            case "hello":
                                                v57 := GenS3{F0: ((!true) || (!false)), F1: 962, F2: ""}
                                                _ = v57
                                            }
                                        }
                                        var v58 uint8 = 128
                                        var v59 int16 = 42
                                        _ = v58
                                        _ = v59
                                    } else {
                                        if (false || true) {
                                            if true {
                                                fmt.Println(v15)
                                            }
                                        } else {
                                            v15 = (v15 + v15)
                                            v3 = (v3 * (-(fi23)))
                                            var v60 string = ""
                                            _ = v60
                                        }
                                        if true {
                                            fmt.Println(v3, fi23, fi23)
                                            switch "" {
                                            case "foo bar":
                                                v61 := GenS4{F0: (fi23 % (((^(v3))) | 1))}
                                                _ = v61
                                            case "":
                                                fmt.Println(fi23, v3, fi23)
                                            case "hello":
                                                var v62 uint32 = v15
                                                _ = v62
                                            default:
                                                fmt.Println(v15, v15)
                                            }
                                            v63 := GenS1{F0: 2}
                                            switch fi23 {
                                            case 303:
                                                fmt.Println(fi23, fi55, v3)
                                            case 32768:
                                                v64 := GenS5{F0: -941}
                                                _ = v64
                                            case 4294967295:
                                                v65 := GenS6{F0: int16(v15), F1: -128}
                                                if (-32768 > -1) {
                                                    var v66 int16 = 1000
                                                    _ = v66
                                                } else {
                                                    v15 -= v15
                                                    if true {
                                                        if true {
                                                            var v67 int64 = 9223372036854775807
                                                            v1 = append(v1, v3)
                                                            _ = v67
                                                        }
                                                        _ = len(v2)
                                                    } else {
                                                        v3 += 4294967295
                                                        var v68 float64 = 0.0
                                                        v3 = 497
                                                        _ = v68
                                                    }
                                                    v69 := GenS7{F0: -107, F1: (fi55 % (((v3 % (((fi55 << ((782) & 63))) | 1))) | 1)), F2: (128 >= 2)}
                                                    if false {
                                                        var v70 float32 = 0.0
                                                        fmt.Println(v15, v70, v70)
                                                        fmt.Println(v70, v3, v15)
                                                        _ = v70
                                                    } else {
                                                        var v71 uint16 = 255
                                                        var v72 uint32 = v15
                                                        fmt.Println(v71, fi55)
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    _ = len(v2)
                                                    _ = v69
                                                }
                                                var v73 uint16 = 1
                                                _ = v65
                                                _ = v73
                                            default:
                                                v3 += fi23
                                            }
                                            v15 = (v15 | v15)
                                            _ = v63
                                        } else {
                                            v74 := GenS4{F0: fi23}
                                            fmt.Println(v15, v3, v3)
                                            _ = v74
                                        }
                                    }
                                    fmt.Println(fi55, v3, fi55)
                                    fmt.Println(v15, fi55, v3)
                                    _ = fi55
                                }
                                fmt.Println(v15)
                            }
                        }
                    }
                    _ = v50
                }
                _ = fi23
            }
            for fi75 := 0; fi75 < 1; fi75++ {
                v5[(v15 + (^(v15)))] = (^(v15))
                if true {
                    if false {
                        if (!false) {
                            var v76 int8 = 1
                            _ = v76
                        }
                    }
                } else {
                    var v77 int64 = 0
                    if true {
                        var v78 uint = 515
                        var v79 uint64 = 423
                        _ = len(v5)
                        _ = v78
                        _ = v79
                    }
                    _ = v77
                }
                v80 := GenS8{F0: 42, F1: 231}
                _ = fi75
                _ = v80
            }
            var v81 uint8 = 127
            _ = v81
        } else {
            fmt.Println(v15, v15)
            v82 := GenS2{F0: 127}
            _ = v82
        }
    }
    v15 = (v15 * (v15 << ((789) & 63)))
    var v83 float32 = ((3.14159 * (1e10 / (0.0 - (3.14159 * ((0.1 + (float32(-2.25) - 1e10)) / 0.1))))) - 1e10)
    for fi84 := 0; fi84 < 6; fi84++ {
        fmt.Println(v83, v15, v3)
        v4 = append(v4, fi84)
        _ = fi84
    }
    for fi85 := 0; fi85 < 15; fi85++ {
        var v86 uint = 18446744073709551615
        _ = fi85
        _ = v86
    }
    for fi87 := 0; fi87 < 4; fi87++ {
        v83 = (0.1 * v83)
        v2[v15] = 42
        _ = fi87
    }
    var v88 int64 = 1
    var v89 uint64 = 507
    delete(v5, 349)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v90 int16 = 32767
    switch 32767 {
    case 951:
        for fi91 := 0; fi91 < 15; fi91++ {
            for fi92 := 0; fi92 < 7; fi92++ {
                if ((v88 <= v88) || true) {
                    v2[1] = v15
                    var v93 uint32 = (^(v15))
                    _ = v93
                } else {
                    var v94 uint32 = v15
                    v95 := v5[929]
                    var v96 int16 = (^(v90))
                    _ = v94
                    _ = v95
                    _ = v96
                }
                fmt.Println(v88, v15)
                if (1 < 255) {
                    _ = len(v2)
                } else {
                    fmt.Println(fi92, v88, v3)
                    switch (fi91 << ((4294967295) & 63)) {
                    case -435:
                        v97 := GenS1{F0: 1}
                        _ = v97
                    case 32767:
                        v83 = 0.0
                    case 128:
                        switch (v3 ^ (fi91 / ((fi91) | 1))) {
                        case -9223372036854775808:
                            var v98 uint8 = 1
                            _ = v98
                        default:
                            if false {
                                var v99 float64 = 3.14159
                                switch v3 {
                                case 42:
                                    var v100 int64 = 42
                                    _ = v100
                                case 828:
                                    if len(v1) > 0 {
                                        v1[0] = 128
                                    }
                                case 9223372036854775807:
                                    fmt.Println(v99, v3, v99)
                                default:
                                    switch "" {
                                    case "foo bar":
                                        v101 := GenS3{F0: true, F1: 2, F2: "a"}
                                        _ = v101
                                    case "a":
                                        v3 -= (fi91 & v3)
                                        switch (-(v3)) {
                                        case 4294967295:
                                            v3 += (^(v3))
                                        default:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        v102 := GenS9{F0: 19, F1: 447, F2: 624}
                                        v103 := GenS2{F0: 127}
                                        var v104 float64 = v99
                                        fmt.Println(fi91, fi92)
                                        _ = v102
                                        _ = v103
                                        _ = v104
                                    case "hello":
                                        if false {
                                            v3 += fi92
                                            fmt.Println(v3, v88)
                                            v99 -= v99
                                        }
                                    }
                                }
                                v83 += float32((1e10 + 1.5))
                                _ = v99
                            }
                        }
                    default:
                        if false {
                            switch fi91 {
                            case 32767:
                                delete(v5, (v15 << ((119) & 63)))
                            case -2147483648:
                                delete(v2, v15)
                            case 9223372036854775807:
                                v89 += (v89 ^ 0)
                            default:
                                var v105 float32 = ((v83 - v83) - v83)
                                _ = v105
                            }
                        } else {
                            var v106 uint32 = v15
                            if (!(!true)) {
                                v107 := GenS10{F0: false, F1: v89, F2: 42}
                                _ = v107
                            }
                            _ = v106
                        }
                    }
                    v108 := GenS10{F0: (v88 <= -1), F1: 1, F2: 255}
                    _ = v108
                }
                v90 = (v90 / (((-(v90))) | 1))
                _ = fi92
            }
            _ = fi91
        }
    default:
        if (42 <= 42) {
            v90 = (v90 ^ (v90 * 185))
        } else {
            v89 = v89
        }
    }
    v89 += v89
    v83 = (float32(3.14159) / v83)
    for fi109 := 0; fi109 < 15; fi109++ {
        fmt.Println(v83, v90)
        _ = fi109
    }
    v110 := v5[(v15 + v15)]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v110:", v110)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v83:", v83)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v110
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v88
    _ = v89
    _ = v90
}
