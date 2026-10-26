package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8; F1 uint8 }
func genFn0(p0 bool) int64 {
    for fi6 := 0; fi6 < 16; fi6++ {
        p0 = p0
        var v7 int32 = 0
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 float64 = 3.14159
        _ = fi8
        _ = v9
    }
    var v10 float64 = 0.0
    _ = p0
    _ = v10
    return 77
}
func genFn1(p0 int64) uint16 {
    v11 := []int{32767, 0, 2147483647}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v12 int = -9223372036854775808
    _ = p0
    _ = v11
    _ = v12
    return 302
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483647, 42}
    var v3 uint16 = 2
    v4 := make(map[uint32]bool)
    var v5 string = "hello"
    v3 = (v3 - 32767)
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int32 = 191
        v15 := GenS1{F0: 1}
        _ = fi13
        _ = v14
        _ = v15
    }
    if ((-535 < -429) || false) {
        for fi16 := 0; fi16 < 12; fi16++ {
            if false {
                var v17 float32 = -2.25
                for fi18 := 0; fi18 < 1; fi18++ {
                    if true {
                        if (!false) {
                            fmt.Println(v5, v3)
                        } else {
                            v5 = "a"
                        }
                        v19 := v4[783]
                        if ((uint(v3) <= 2) && (v5 >= "a")) {
                            fmt.Println(v17)
                            fmt.Println(fi16, fi18)
                            if v19 {
                                fmt.Println(v19)
                                v3 = (v3 / ((v3) | 1))
                                switch 2147483647 {
                                case -979:
                                    if (42 == 180) {
                                        v20 := GenS2{F0: -2147483648, F1: ("foo bar" + "")}
                                        v21 := GenS2{F0: -401, F1: v5}
                                        _ = v20
                                        _ = v21
                                    } else {
                                        switch ((v5 + (v5 + (v5 + "hello"))) + "a") {
                                        case "a":
                                            v17 -= v17
                                        case "hello":
                                            _ = len(v2)
                                        case "foo bar":
                                            switch (v5 + v5) {
                                            case "hello":
                                                fmt.Println(v19)
                                            case "a":
                                                v3 = (v3 | v3)
                                            case "foo bar":
                                                v4[2] = (3.14159 <= -2.25)
                                            default:
                                                var v22 uint = 42
                                                _ = v22
                                            }
                                        }
                                        var v23 uint16 = 42
                                        v19 = (!(((!false) || (v19 && v19)) || v19))
                                        v23 = 2
                                        var v24 uint16 = (v23 - (v23 - genFn1(0)))
                                        v23 -= (v24 & (^(v23)))
                                        v17 = (1e10 - ((v17 - (v17 / (v17 * ((0.1 + v17) * v17)))) * 1e10))
                                        v25 := GenS2{F0: 0, F1: v5}
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                default:
                                    v26 := v1[753]
                                    _ = v26
                                }
                                var v27 int64 = -549
                                _ = v27
                            }
                        }
                        _ = v19
                    }
                    var v28 uint16 = (^(v3))
                    _ = fi18
                    _ = v28
                }
                _ = v17
            } else {
                var v29 int64 = 9223372036854775807
                v5 = ("\n\t" + v5)
                _ = v29
            }
            var v30 uint64 = 8
            v31 := GenS1{F0: 1}
            _ = fi16
            _ = v30
            _ = v31
        }
    } else {
        fmt.Println(v5, v3)
    }
    var v32 float32 = ((((0.1 * 0.0) - 0.1) - (1e10 / ((0.1 * 1.5) * 1.5))) / (((3.14159 / (0.1 + 0.0)) + float32(((((-2.25 + float64(float32((1.5 + float64(1e10))))) + 0.0) - 1e10) / -2.25))) / 3.14159))
    for fi33 := 0; fi33 < 12; fi33++ {
        if true {
            v3 = v3
            fmt.Println(v3)
            switch fi33 {
            case -902:
                var v34 string = v5
                _ = v34
            case 4294967295:
                for fi35 := 0; fi35 < 12; fi35++ {
                    var v36 uint = 458
                    _ = fi35
                    _ = v36
                }
            case 1:
                var v37 int32 = -2147483648
                _ = v37
            }
        }
        v38 := GenS3{F0: (^(v3)), F1: true}
        v3 = genFn1(-1)
        _ = fi33
        _ = v38
    }
    for fi39 := 0; fi39 < 5; fi39++ {
        fmt.Println(v32, fi39, v3)
        _ = fi39
    }
    switch "\n\t" {
    case "":
        for fi40 := 0; fi40 < 9; fi40++ {
            if (!((("" != "a") || true) || true)) {
                v5 = v5
                for fi41 := 0; fi41 < 12; fi41++ {
                    var v42 uint64 = 533
                    var v43 int8 = 127
                    if false {
                        v44 := GenS2{F0: 69, F1: v5}
                        v42 = v42
                        _ = v44
                    } else {
                        v45 := GenS4{F0: 255}
                        var v46 string = ("" + "a")
                        switch 4294967295 {
                        case 128:
                            if true {
                                v47 := v1[fi40]
                                if true {
                                    var v48 int = 42
                                    var v49 string = v46
                                    _ = v48
                                    _ = v49
                                } else {
                                    if ((v42 + (v42 % ((v42) | 1))) != v42) {
                                        delete(v4, 5)
                                    }
                                }
                                _ = v47
                            } else {
                                fmt.Println(v5)
                                if true {
                                    var v50 int32 = -2147483648
                                    switch (((((((v5 + v5) + "\n\t") + v46) + "foo bar") + (v46 + (v5 + v46))) + v46) + "a") {
                                    case "a":
                                        v50 = (v50 / ((0) | 1))
                                    default:
                                        if (false && (!((!true) && true))) {
                                            switch 0 {
                                            case -2147483648:
                                                if false {
                                                    switch (^(fi40)) {
                                                    case -147:
                                                        v32 = (1.5 / v32)
                                                    case 42:
                                                        v43 -= v43
                                                    case -2147483648:
                                                        var v51 uint32 = 0
                                                        _ = v51
                                                    default:
                                                        v52 := GenS3{F0: v3, F1: (uint8(v42) < 42)}
                                                        _ = v52
                                                    }
                                                }
                                            }
                                            if (!(-365 < 42)) {
                                                v4[2147483648] = (2147483648 >= 0)
                                            }
                                        } else {
                                            fmt.Println(fi41, v46)
                                            _ = len(v2)
                                        }
                                    }
                                    var v53 string = (v5 + v5)
                                    _ = v50
                                    _ = v53
                                }
                            }
                        case 4294967295:
                            var v54 int64 = 0
                            _ = v54
                        case 353:
                            v55 := GenS1{F0: 546}
                            v3 = uint16(v42)
                            var v56 bool = (527 <= -9223372036854775808)
                            v57 := GenS5{F0: 127}
                            _ = v55
                            _ = v56
                            _ = v57
                        }
                        v43 = (v43 - 100)
                        _ = v45
                        _ = v46
                    }
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                if false {
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    if (!true) {
                        if (1.5 == 0.1) {
                            for fi58 := 0; fi58 < 6; fi58++ {
                                v59 := GenS6{F0: v5, F1: 42}
                                if true {
                                    if (1 <= 100) {
                                        if (float32((1e10 - (0.0 * (-2.25 - (1e10 * (0.0 / 0.1)))))) <= v32) {
                                            delete(v1, (^(fi40)))
                                            var v60 uint8 = 1
                                            _ = v60
                                        } else {
                                            switch (fi40 + fi58) {
                                            case 32768:
                                                v32 += (v32 / (((((-2.25 * v32) * (((v32 - v32) - 0.1) - (0.1 * 0.0))) * v32) * 3.14159) * v32))
                                            case 4294967295:
                                                var v61 float32 = (v32 + v32)
                                                _ = v61
                                            case -1:
                                                fmt.Println(v5, v3, v5)
                                            }
                                            v32 += v32
                                        }
                                        fmt.Println(v5, fi58, fi58)
                                        v3 = v3
                                    } else {
                                        if (9223372036854775807 >= 9223372036854775807) {
                                            var v62 string = ((("a" + "hello") + "foo bar") + v5)
                                            v63 := GenS1{F0: 4294967295}
                                            switch (((v5 + v62) + v5) + (((v5 + v5) + "hello") + (v5 + "a"))) {
                                            case "\n\t":
                                                var v64 int = (fi40 | -927)
                                                _ = v64
                                            default:
                                                v62 = "hello"
                                            }
                                            v32 = (float32(1.5) / 1e10)
                                            fmt.Println(v32, v62)
                                            _ = v62
                                            _ = v63
                                        } else {
                                            var v65 uint32 = 2
                                            var v66 bool = false
                                            v65 = (v65 >> ((255) & 63))
                                            _ = v65
                                            _ = v66
                                        }
                                    }
                                }
                                var v67 int16 = -1
                                if (false || (fi40 < 128)) {
                                    var v68 uint = 18446744073709551615
                                    _ = v68
                                }
                                _ = fi58
                                _ = v59
                                _ = v67
                            }
                        } else {
                            switch (fi40 & (fi40 / ((fi40) | 1))) {
                            case -2147483648:
                                for fi69 := 0; fi69 < 20; fi69++ {
                                    if false {
                                        v70 := GenS5{F0: -128}
                                        v5 = v5
                                        var v71 string = (v5 + ("\n\t" + (((v5 + "\n\t") + (v5 + "hello")) + v5)))
                                        _ = v70
                                        _ = v71
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v72 int8 = 100
                                    _ = fi69
                                    _ = v72
                                }
                            }
                            var v73 int32 = 42
                            _ = v73
                        }
                        var v74 uint8 = 1
                        if false {
                            var v75 uint8 = v74
                            switch fi40 {
                            case -216:
                                fmt.Println(fi40, v5, v74)
                            case 32767:
                                var v76 int8 = 127
                                _ = v76
                            default:
                                for fi77 := 0; fi77 < 12; fi77++ {
                                    fmt.Println(v32, fi77, fi77)
                                    if true {
                                        v78 := GenS2{F0: -499, F1: ((("a" + "") + "hello") + v5)}
                                        var v79 float32 = v32
                                        v80 := GenS6{F0: v5, F1: 42}
                                        _ = v78
                                        _ = v79
                                        _ = v80
                                    } else {
                                        v32 = v32
                                        v5 = ((v5 + "a") + (v5 + v5))
                                        v32 += (v32 + ((v32 - ((float32((1.5 / 1.0)) - (((v32 - 3.14159) - (3.14159 * 3.14159)) / 1e10)) / v32)) * (1.5 * v32)))
                                        if false {
                                            v3 = v3
                                        } else {
                                            fmt.Println(v5, v75, v74)
                                            v3 += (v3 % ((2) | 1))
                                        }
                                        switch "a" {
                                        case "\n\t":
                                            v3 = (v3 - 1)
                                        default:
                                            var v81 float64 = (1e10 + 0.0)
                                            _ = v81
                                        }
                                    }
                                    if (true || true) {
                                        fmt.Println(v74, v32, v5)
                                    } else {
                                        var v82 bool = (1 <= 3)
                                        v32 = (v32 * ((0.0 / v32) / (v32 / -2.25)))
                                        switch (v5 + v5) {
                                        case "a":
                                            if ((v82 || v82) || (4294967295 > 2)) {
                                                v4[2147483648] = (!v82)
                                                v75 = v75
                                            }
                                        default:
                                            var v83 int32 = 0
                                            v84 := v4[65535]
                                            var v85 int16 = -724
                                            _ = v83
                                            _ = v84
                                            _ = v85
                                        }
                                        if (((1e10 / 1.5) < 0.1) || false) {
                                            fmt.Println(fi77, fi40)
                                            var v86 float32 = v32
                                            _ = v86
                                        } else {
                                            delete(v1, (fi77 ^ (fi77 & fi77)))
                                        }
                                        switch fi40 {
                                        case -2147483648:
                                            if len(v2) > 0 {
                                                v2[0] = 65535
                                            }
                                        }
                                        _ = v82
                                    }
                                    _ = fi77
                                }
                            }
                            _ = v75
                        } else {
                            v87 := GenS7{F0: 1}
                            var v88 uint64 = uint64(v74)
                            _ = v87
                            _ = v88
                        }
                        _ = v74
                    } else {
                        switch v5 {
                        case "":
                            var v89 uint64 = 2
                            _ = v89
                        }
                        var v90 string = (v5 + (v5 + (("" + ("\n\t" + v5)) + v5)))
                        v32 = 1.5
                        _ = v90
                    }
                    for fi91 := 0; fi91 < 6; fi91++ {
                        var v92 int32 = 42
                        _ = fi91
                        _ = v92
                    }
                }
            }
            fmt.Println(fi40, v3, v3)
            if ((0.1 >= (3.14159 - 3.14159)) && true) {
                var v93 float64 = 1.5
                _ = v93
            }
            _ = fi40
        }
    }
    delete(v1, -238)
    var v94 uint = 42
    if ((!(192 <= 0)) && true) {
        switch v5 {
        case "a":
            fmt.Println(v32, v94)
        case "hello":
            for fi95 := 0; fi95 < 16; fi95++ {
                if false {
                    var v96 uint64 = 42
                    _ = v96
                }
                _ = fi95
            }
        }
        v94 -= (v94 % (((v94 * v94)) | 1))
    }
    fmt.Println(v32)
    if (v32 != -2.25) {
        var v97 uint32 = 2147483648
        var v98 float64 = 1e10
        _ = v97
        _ = v98
    } else {
        v99 := GenS8{F0: 100, F1: uint8(v3)}
        switch 2147483647 {
        case 0:
            for fi100 := 0; fi100 < 8; fi100++ {
                for fi101 := 0; fi101 < 17; fi101++ {
                    v5 = ((v5 + v5) + (("" + (v5 + "\n\t")) + ("\n\t" + v5)))
                    fmt.Println(v5, v32, v5)
                    v5 = v5
                    _ = fi101
                }
                if (((255 < 892) || false) && (33 >= 127)) {
                    for fi102 := 0; fi102 < 11; fi102++ {
                        var v103 uint64 = 0
                        var v104 uint8 = 127
                        _ = fi102
                        _ = v103
                        _ = v104
                    }
                    v5 = "foo bar"
                    fmt.Println(v94, fi100)
                }
                var v105 string = (((v5 + (v5 + (((v5 + ((v5 + v5) + v5)) + v5) + v5))) + "a") + "hello")
                _ = fi100
                _ = v105
            }
        case -705:
            var v106 string = v5
            _ = v106
        case -9223372036854775808:
            for fi107 := 0; fi107 < 2; fi107++ {
                var v108 float64 = (0.0 - float64(-2.25))
                _ = fi107
                _ = v108
            }
        default:
            fmt.Println(v3)
        }
        _ = v99
    }
    v5 = v5
    v5 = ("" + v5)
    if false {
        fmt.Println(v3, v94, v3)
        v3 = (v3 / (((v3 - v3)) | 1))
    } else {
        v32 += v32
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v3 = (v3 / (((v3 * genFn1(0))) | 1))
    v109 := GenS2{F0: -696, F1: (v5 + (v5 + (v5 + "a")))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v109
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v94
}
