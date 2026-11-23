package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 string }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int; F2 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 int; F2 uint8 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int8 }
func genFn0(p0 int64, p1 uint64, p2 bool) uint32 {
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi7
        _ = v8
    }
    p1 = (p1 / ((p1) | 1))
    if ((^(p1)) < (p1 ^ (p1 ^ (p1 ^ (p1 * p1))))) {
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS1{F0: 42, F1: (((781 < (^(p1))) && (!p2)) && (4294967295 > 0)), F2: "foo bar"}
            _ = fi9
            _ = v10
        }
        var v11 int32 = 1
        switch -2147483648 {
        case -927:
            switch "hello" {
            case "foo bar":
                fmt.Println(p2, v11, p1)
            case "\n\t":
                var v12 float32 = 1.5
                _ = v12
            }
        case 2147483647:
            var v13 uint64 = p1
            _ = v13
        case 127:
            v14 := GenS1{F0: 2, F1: false, F2: ""}
            _ = v14
        default:
            v15 := GenS1{F0: 255, F1: (v11 > v11), F2: ""}
            _ = v15
        }
        _ = v11
    }
    var v16 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 0
}
func genFn1(p0 int16, p1 int) uint16 {
    v17 := GenS2{F0: p0, F1: 100}
    v18 := []int{}
    switch 127 {
    case 47:
        fmt.Println(p0, p0)
    default:
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 484
}

func main() {
    var v1 int64 = 484
    var v2 int8 = 127
    v3 := []uint32{2147483648, 65535, 2, 42}
    var v4 uint8 = 101
    var v5 uint64 = 271
    var v6 uint = 795
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 int8 = (^(v2))
        _ = len(v3)
        _ = fi19
        _ = v20
    }
    if (v4 != 36) {
        _ = len(v3)
    }
    switch "" {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    default:
        var v21 int64 = v1
        _ = v21
    }
    v3 = append(v3, 1)
    var v22 int8 = v2
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 int8 = -1
        fmt.Println(v5)
        _ = len(v3)
        _ = fi23
        _ = v24
    }
    v25 := GenS2{F0: -32768, F1: v22}
    v6 = v6
    if (127 > (v2 - -128)) {
        if ((v4 - 190) < 128) {
            switch -57 {
            case 1:
                v26 := GenS1{F0: 0, F1: true, F2: "a"}
                _ = v26
            case 0:
                switch "foo bar" {
                case "\n\t":
                    for fi27 := 0; fi27 < 16; fi27++ {
                        var v28 float64 = (((0.1 * (0.0 - (0.0 + 1.5))) * (1.5 * (3.14159 * 0.1))) * 0.0)
                        var v29 uint64 = (^(v5))
                        v30 := GenS3{F0: 9223372036854775807}
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                case "hello":
                    switch "hello" {
                    case "":
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                    case "foo bar":
                        var v31 bool = (42 <= 554)
                        _ = v31
                    }
                default:
                    if (!(false || true)) {
                        var v32 float64 = float64(0.1)
                        _ = v32
                    } else {
                        v33 := GenS3{F0: 42}
                        v34 := GenS4{F0: 32767, F1: 642, F2: 65535}
                        _ = v33
                        _ = v34
                    }
                }
            }
            fmt.Println(v6, v4)
        }
        if true {
            var v35 float64 = 0.0
            _ = v35
        } else {
            var v36 float64 = 3.14159
            _ = v36
        }
    }
    if false {
        v3 = append(v3, 1)
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        switch fi37 {
        case 201:
            v38 := GenS5{F0: uint8(v5)}
            _ = v38
        default:
            for fi39 := 0; fi39 < 19; fi39++ {
                v40 := GenS3{F0: (-(v1))}
                var v41 int16 = 1
                v42 := GenS6{F0: (v2 | -101), F1: (fi37 / ((542) | 1)), F2: v4}
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                if true {
                    switch (fi39 & fi39) {
                    case -899:
                        var v43 int64 = (v1 | (^(v1)))
                        var v44 int = (fi37 - (fi39 | (fi39 & fi37)))
                        _ = len(v3)
                        _ = v43
                        _ = v44
                    case 32767:
                        switch (fi39 - fi39) {
                        case -1:
                            var v45 float32 = (((1e10 + (3.14159 / (1.5 * 1.5))) - 3.14159) + (3.14159 + 1e10))
                            _ = v45
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v6 -= 707
                            if len(v3) > 0 {
                                v3[0] = 2147483648
                            }
                        }
                    case 127:
                        switch int(v5) {
                        case 32768:
                            v41 = -502
                        }
                    default:
                        fmt.Println(v5, v5)
                        switch fi37 {
                        case 855:
                            v41 -= 42
                        case -1:
                            if true {
                                v41 = v41
                                if false {
                                    _ = len(v3)
                                    v6 = (v6 * (^(v6)))
                                    switch "" {
                                    case "hello":
                                        v41 = (v41 & v41)
                                    case "a":
                                        v22 += v22
                                        fmt.Println(fi37)
                                    }
                                    var v46 float32 = (0.0 + 3.14159)
                                    _ = v46
                                }
                            } else {
                                var v47 uint32 = 0
                                if (!true) {
                                    var v48 bool = false
                                    if v48 {
                                        var v49 int = (fi37 + (fi39 & (fi37 | (fi37 | fi37))))
                                        v50 := GenS4{F0: fi39, F1: (fi39 | (fi37 / (((fi39 * v49)) | 1))), F2: 939}
                                        var v51 int32 = -2147483648
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = genFn0(int64(v41), 0, (71 == 127))
                                        }
                                        if ((v47 << ((uint(v5)) & 63)) != v47) {
                                            v3 = append(v3, genFn0((v1 | 42), (v5 * (v5 % ((v5) | 1))), (!((fi37 << (((v6 * (v6 >> ((v6) & 63)))) & 63)) > fi37))))
                                            v41 -= 1000
                                        }
                                        if false {
                                            switch (("\n\t" + "a") + (("" + "\n\t") + "a")) {
                                            case "hello":
                                                var v52 float32 = (float32((1e10 - 1e10)) - float32(float64((float32(1.5) / (0.1 / (((-2.25 / 0.1) * (0.1 / 3.14159)) - -2.25))))))
                                                _ = v52
                                            case "a":
                                                var v53 int16 = int16(v47)
                                                var v54 float64 = 1e10
                                                v4 += 42
                                                _ = v53
                                                _ = v54
                                            case "foo bar":
                                                switch (fi39 >> (((v6 | v6)) & 63)) {
                                                case 4294967295:
                                                    var v55 int64 = v1
                                                    _ = v55
                                                }
                                            }
                                        }
                                    }
                                    _ = v48
                                } else {
                                    v5 = (^(v5))
                                    var v56 uint32 = v47
                                    var v57 float64 = -2.25
                                    _ = v56
                                    _ = v57
                                }
                                _ = v47
                            }
                        case 1:
                            var v58 uint32 = 2
                            _ = v58
                        }
                        var v59 uint16 = 481
                        _ = v59
                    }
                    switch ("" + "") {
                    case "foo bar":
                        v3 = append(v3, 0)
                    case "\n\t":
                        switch (fi37 | (fi37 / ((fi37) | 1))) {
                        case -582:
                            var v60 int32 = 0
                            v41 = v41
                            var v61 int64 = 774
                            _ = v60
                            _ = v61
                        case 128:
                            var v62 int8 = v22
                            _ = v62
                        }
                    case "":
                        fmt.Println(v6, v6)
                    }
                    if false {
                        v41 += v41
                        fmt.Println(v41, v1)
                    } else {
                        if ((v6 / (((v6 * v6)) | 1)) != 65535) {
                            var v63 bool = (float32(float64(1e10)) < 3.14159)
                            v4 = (v4 | v4)
                            if false {
                                if ((255 <= (v4 * v4)) || false) {
                                    v64 := GenS7{F0: (v4 * v4), F1: 1}
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                    if (1 != 218) {
                                        if true {
                                            var v65 uint64 = (v5 | v5)
                                            _ = v65
                                        } else {
                                            var v66 uint16 = 32767
                                            v67 := GenS5{F0: 132}
                                            var v68 string = ("hello" + "\n\t")
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        }
                                        v69 := GenS4{F0: 1, F1: fi39, F2: 255}
                                        _ = v69
                                    } else {
                                        switch (fi39 * (fi37 & fi39)) {
                                        case -826:
                                            v70 := GenS4{F0: fi39, F1: int(v41), F2: 255}
                                            _ = v70
                                        case 32768:
                                            v1 = (^(v1))
                                        }
                                        var v71 bool = (true || (!(v63 && (((!v63) || true) && (!v63)))))
                                        _ = v71
                                    }
                                    _ = v64
                                }
                                fmt.Println(v22)
                                v3 = append(v3, 961)
                                v72 := GenS5{F0: 2}
                                var v73 bool = (1e10 < 1e10)
                                _ = v72
                                _ = v73
                            }
                            _ = v63
                        }
                    }
                } else {
                    var v74 int64 = (v1 / ((v1) | 1))
                    var v75 uint8 = v4
                    var v76 int = -9223372036854775808
                    v76 += v76
                    _ = v74
                    _ = v75
                    _ = v76
                }
                v77 := GenS1{F0: 2, F1: (false || false), F2: (("a" + "foo bar") + "hello")}
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
                _ = v77
            }
        }
        var v78 int32 = -560
        v22 -= v22
        _ = fi37
        _ = v78
    }
    if (!(51 < (^(v6)))) {
        var v79 string = (("a" + "\n\t") + "")
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v79
    }
    fmt.Println(v4, v4)
    v6 += 4294967295
    var v80 float32 = ((0.0 - float32(0.0)) * -2.25)
    if (v80 == v80) {
        var v81 uint32 = 1
        for fi82 := 0; fi82 < 8; fi82++ {
            switch "\n\t" {
            case "a":
                switch ("hello" + ("a" + (((("foo bar" + "\n\t") + "foo bar") + ("a" + ("foo bar" + "foo bar"))) + ("" + ("" + (("\n\t" + "hello") + "a")))))) {
                case "foo bar":
                    v3 = append(v3, (v81 + (v81 >> (((^(v6))) & 63))))
                case "hello":
                    if true {
                        v22 -= (v22 + -7)
                    } else {
                        var v83 int = (^(fi82))
                        switch "hello" {
                        case "":
                            switch v83 {
                            case 424:
                                fmt.Println(v1, v22, v81)
                            case 578:
                                var v84 uint8 = v4
                                _ = v84
                            }
                        case "hello":
                            for fi85 := 0; fi85 < 8; fi85++ {
                                v2 += (v22 % ((v2) | 1))
                                _ = fi85
                            }
                        case "\n\t":
                            v83 -= (fi82 | v83)
                        default:
                            var v86 float32 = v80
                            _ = v86
                        }
                        v5 -= v5
                        _ = v83
                    }
                default:
                    v87 := GenS8{F0: 695}
                    _ = v87
                }
            case "":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                switch (fi82 << ((v6) & 63)) {
                case 128:
                    var v88 uint8 = v4
                    _ = v88
                case 0:
                    var v89 uint8 = v4
                    _ = v89
                case -990:
                    v6 = 192
                default:
                    var v90 uint = v6
                    _ = v90
                }
            }
            if true {
                v91 := GenS9{F0: 2147483647}
                if len(v3) > 0 {
                    v3[0] = (^(v81))
                }
                if (((1e10 - 0.0) * 1.5) != 1.5) {
                    if (!((v5 << ((v6) & 63)) < (v5 >> ((307) & 63)))) {
                        if false {
                            v6 = (v6 / ((4294967295) | 1))
                        } else {
                            fmt.Println(v4, v6, v2)
                        }
                    } else {
                        for fi92 := 0; fi92 < 18; fi92++ {
                            v93 := GenS10{F0: (v2 | v22)}
                            v94 := GenS10{F0: v2}
                            var v95 int = (^(fi92))
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v96 uint64 = 0
                            var v97 string = "a"
                            var v98 uint64 = v96
                            _ = fi92
                            _ = v93
                            _ = v94
                            _ = v95
                            _ = v96
                            _ = v97
                            _ = v98
                        }
                    }
                    v80 += v80
                } else {
                    var v99 float32 = v80
                    if len(v3) > 0 {
                        v3[0] = v81
                    }
                    v100 := GenS4{F0: (fi82 & (fi82 | (-(fi82)))), F1: int(v4), F2: genFn1(1, 787)}
                    _ = v100
                    _ = v99
                }
                _ = v91
            }
            _ = fi82
        }
        _ = v81
    }
    v101 := GenS2{F0: 1000, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v101
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v80
}
