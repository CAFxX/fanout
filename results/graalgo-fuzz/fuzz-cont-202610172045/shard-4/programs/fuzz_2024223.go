package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 int64; F2 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 uint64; F1 bool; F2 int }
type GenS11 struct { F0 int16; F1 uint8; F2 int32 }
func genFn0(p0 int32) uint {
    if (true && (!true)) {
        p0 -= p0
        v6 := []int{128}
        v7 := GenS1{F0: true, F1: true, F2: 65535}
        _ = v6
        _ = v7
    } else {
        fmt.Println(p0)
    }
    v8 := GenS2{F0: 0, F1: 65535, F2: 393}
    v9 := GenS2{F0: 1, F1: 1, F2: 153}
    _ = p0
    _ = v8
    _ = v9
    return 680
}
func genFn1(p0 int, p1 uint32, p2 bool) int {
    p2 = true
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := []int{fi10}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint8 = 127
        v14 := []int{(-(p0)), fi12}
        _ = fi12
        _ = v13
        _ = v14
    }
    switch ("hello" + "foo bar") {
    case "hello":
        p1 = (p1 | p1)
    default:
        if (1.5 == -2.25) {
            fmt.Println(p2)
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 int8 = 1
                fmt.Println(p2)
                switch fi15 {
                case 128:
                    if false {
                        switch "a" {
                        case "\n\t":
                            var v17 uint8 = 1
                            _ = v17
                        case "foo bar":
                            p1 = 4294967295
                        default:
                            var v18 string = "hello"
                            _ = v18
                        }
                    } else {
                        v19 := GenS2{F0: -1, F1: 0, F2: 42}
                        _ = v19
                    }
                case 2147483647:
                    for fi20 := 0; fi20 < 17; fi20++ {
                        switch "hello" {
                        case "a":
                            var v21 uint16 = 289
                            _ = v21
                        }
                        _ = fi20
                    }
                case -2147483648:
                    v22 := []int{p0, p0, fi15}
                    _ = v22
                default:
                    if ((p1 + p1) < (p1 - (p1 + p1))) {
                        for fi23 := 0; fi23 < 6; fi23++ {
                            if (p2 || (42 < 4294967295)) {
                                fmt.Println(v16, p0)
                                var v24 int32 = 0
                                switch "\n\t" {
                                case "":
                                    var v25 int8 = v16
                                    _ = v25
                                default:
                                    v26 := GenS2{F0: 1, F1: 0, F2: 438}
                                    _ = v26
                                }
                                _ = v24
                            } else {
                                var v27 float64 = (1.5 * (1e10 - float64((float32(float64(0.1)) * -2.25))))
                                var v28 uint16 = 32767
                                _ = v27
                                _ = v28
                            }
                            switch "foo bar" {
                            case "foo bar":
                                v29 := []int{(fi15 | 621), 2147483647, (fi23 / ((fi23) | 1)), (fi23 * 32768)}
                                _ = v29
                            default:
                                if p2 {
                                    switch (fi23 | 2147483647) {
                                    case 32768:
                                        var v30 uint32 = p1
                                        _ = v30
                                    case -164:
                                        p0 += (fi23 ^ (p0 ^ (fi15 ^ fi15)))
                                    }
                                    var v31 string = ("" + ((((("\n\t" + "hello") + (("\n\t" + "foo bar") + "")) + "\n\t") + "\n\t") + ""))
                                    switch "\n\t" {
                                    case "a":
                                        var v32 uint8 = 2
                                        _ = v32
                                    }
                                    v33 := GenS1{F0: true, F1: (p2 || false), F2: 1}
                                    var v34 uint = 1
                                    _ = v31
                                    _ = v33
                                    _ = v34
                                }
                            }
                            _ = fi23
                        }
                        fmt.Println(v16)
                    } else {
                        for fi35 := 0; fi35 < 3; fi35++ {
                            p2 = false
                            fmt.Println(fi35)
                            _ = fi35
                        }
                        var v36 float32 = 1e10
                        v37 := []int{32768, 9223372036854775807, -596, -68}
                        _ = v36
                        _ = v37
                    }
                }
                _ = fi15
                _ = v16
            }
            if (!(p1 >= 42)) {
                p0 = (^(p0))
            } else {
                for fi38 := 0; fi38 < 5; fi38++ {
                    p1 = (p1 & p1)
                    var v39 int16 = 792
                    _ = fi38
                    _ = v39
                }
            }
        }
    }
    v40 := GenS1{F0: (!p2), F1: (false && p2), F2: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v40
    return (p0 ^ p0)
}

func main() {
    var v1 uint32 = 42
    var v2 string = ("\n\t" + "hello")
    var v3 int = -2147483648
    var v4 bool = false
    var v5 int32 = -1
    v2 = ("\n\t" + v2)
    v41 := []int{1, v3, (v3 / ((1) | 1))}
    var v42 string = ("hello" + "a")
    v2 = v42
    var v43 int64 = 9223372036854775807
    for fi44 := 0; fi44 < 5; fi44++ {
        if len(v41) >= 2 {
            v41 = v41[1:len(v41)]
        }
        _ = fi44
    }
    for fi45 := 0; fi45 < 1; fi45++ {
        if (false && (v4 && (-368 == 0))) {
            v1 -= v1
            for fi46 := 0; fi46 < 19; fi46++ {
                if true {
                    switch v2 {
                    case "a":
                        var v47 uint32 = (v1 - (v1 - 2))
                        _ = v47
                    case "hello":
                        v42 = ((v2 + (v2 + (v2 + v42))) + v2)
                    case "":
                        v5 -= 0
                    default:
                        fmt.Println(fi45, v1, fi45)
                    }
                    v48 := GenS3{F0: 981}
                    v43 -= (v43 << ((4294967295) & 63))
                    _ = v48
                }
                if v4 {
                    switch "hello" {
                    case "hello":
                        v49 := GenS3{F0: -243}
                        _ = v49
                    case "":
                        v50 := GenS3{F0: 42}
                        _ = v50
                    case "a":
                        if v4 {
                            if v4 {
                                var v51 uint32 = (v1 - 4294967295)
                                _ = v51
                            }
                            switch fi46 {
                            case -56:
                                fmt.Println(fi45)
                            default:
                                var v52 int64 = (v43 | v43)
                                fmt.Println(v52, fi46, fi45)
                                var v53 string = ((((v42 + v42) + v2) + v42) + v42)
                                switch fi46 {
                                case 981:
                                    v4 = false
                                case 9223372036854775807:
                                    var v54 int64 = 1
                                    var v55 int8 = -44
                                    if false {
                                        v56 := GenS3{F0: -32768}
                                        switch genFn1(fi45, 65535, false) {
                                        case 127:
                                            v43 = v43
                                        case 128:
                                            if len(v41) > 0 {
                                                v41[0] = (v3 * genFn1(2147483647, v1, (0.0 >= 1e10)))
                                            }
                                        case 42:
                                            if true {
                                                if len(v41) >= 2 {
                                                    v41 = v41[1:len(v41)]
                                                }
                                                v53 = v42
                                            }
                                        }
                                        _ = v56
                                    }
                                    fmt.Println(v1, v1)
                                    var v57 int16 = 1
                                    if len(v41) >= 2 {
                                        v41 = v41[1:len(v41)]
                                    }
                                    if v4 {
                                        if (1.5 < 0.0) {
                                            v52 -= v52
                                        }
                                        fmt.Println(v42, v2, v1)
                                    } else {
                                        v57 = (-(v57))
                                        if (!(0 < 4294967295)) {
                                            var v58 int = (v3 >> ((2) & 63))
                                            if len(v41) >= 2 {
                                                v41 = v41[1:len(v41)]
                                            }
                                            _ = v58
                                        } else {
                                            v57 += (v57 & (v57 + 1000))
                                            if v4 {
                                                var v59 int64 = 9223372036854775807
                                                v1 = 1
                                                var v60 uint8 = 2
                                                _ = v59
                                                _ = v60
                                            }
                                            switch (^(fi45)) {
                                            case -799:
                                                switch (fi45 & fi46) {
                                                case 127:
                                                    var v61 int = 9223372036854775807
                                                    _ = v61
                                                case -9223372036854775808:
                                                    var v62 float64 = (0.1 + (3.14159 - float64(1e10)))
                                                    _ = v62
                                                }
                                            case 2147483647:
                                                fmt.Println(v57, v54)
                                            }
                                        }
                                        fmt.Println(fi46, v52)
                                    }
                                    _ = v54
                                    _ = v55
                                    _ = v57
                                case 128:
                                    if (!(-2.25 > 1e10)) {
                                        v41 = append(v41, (fi45 >> ((genFn0(-1)) & 63)))
                                        v52 += (v52 - (v43 + (^(v43))))
                                        v63 := GenS4{F0: (("" + v2) + "a")}
                                        _ = v63
                                    }
                                default:
                                    fmt.Println(v4, v5, v52)
                                }
                                switch (v2 + (v42 + v2)) {
                                case "hello":
                                    v64 := GenS1{F0: v4, F1: (2 == 0), F2: 1}
                                    _ = v64
                                case "foo bar":
                                    v65 := GenS5{F0: 666}
                                    _ = v65
                                default:
                                    if (true || (false || v4)) {
                                        var v66 int = genFn1(fi45, (v1 >> ((42) & 63)), (((v52 % ((v52) | 1)) >= 1) || (4294967295 <= 4294967295)))
                                        _ = v66
                                    }
                                }
                                _ = v52
                                _ = v53
                            }
                        }
                    }
                    if (v42 > v2) {
                        fmt.Println(v3, v43)
                        v41 = append(v41, genFn1(-1, v1, false))
                    } else {
                        v67 := GenS4{F0: ("" + (v2 + (v2 + ((v2 + (v42 + v42)) + ""))))}
                        v43 -= v43
                        _ = len(v41)
                        _ = v67
                    }
                } else {
                    v5 += (v5 & -1)
                    v68 := GenS6{F0: 42}
                    if v4 {
                        if (1e10 < 1.5) {
                            switch "a" {
                            case "hello":
                                v69 := GenS7{F0: 1, F1: 42, F2: 100}
                                _ = v69
                            default:
                                if ((v43 + v43) < (v43 << ((0) & 63))) {
                                    var v70 int32 = v5
                                    _ = v70
                                } else {
                                    v3 = genFn1(127, (v1 << ((65535) & 63)), v4)
                                    if len(v41) >= 2 {
                                        v41 = v41[1:len(v41)]
                                    }
                                    if v4 {
                                        var v71 int32 = (v5 << ((0) & 63))
                                        v5 = v71
                                        _ = v71
                                    }
                                }
                            }
                        } else {
                            var v72 float64 = 0.0
                            var v73 int64 = -9223372036854775808
                            if (42 == -128) {
                                v41 = append(v41, fi45)
                                var v74 float64 = (v72 + (float64((1e10 + 1.5)) * v72))
                                _ = v74
                            }
                            _ = v72
                            _ = v73
                        }
                        v75 := GenS8{F0: v5}
                        _ = v75
                    } else {
                        v43 = (v43 << ((65535) & 63))
                        if len(v41) >= 2 {
                            v41 = v41[1:len(v41)]
                        }
                        v41 = append(v41, fi45)
                    }
                    _ = v68
                }
                var v76 int = 691
                _ = fi46
                _ = v76
            }
            if ((fi45 != v3) || v4) {
                var v77 uint8 = 102
                v78 := GenS9{F0: -113, F1: v1}
                _ = v77
                _ = v78
            }
        } else {
            switch 42 {
            case 721:
                fmt.Println(v42)
            case -622:
                v42 = v42
            }
        }
        v4 = (!(!v4))
        v79 := GenS10{F0: 1, F1: v4, F2: v3}
        _ = fi45
        _ = v79
    }
    v1 = (^(v1))
    for fi80 := 0; fi80 < 2; fi80++ {
        var v81 float32 = 1e10
        var v82 int8 = 127
        for fi83 := 0; fi83 < 9; fi83++ {
            fmt.Println(v5, v81, v81)
            _ = fi83
        }
        _ = fi80
        _ = v81
        _ = v82
    }
    if false {
        if (95 <= -11) {
            v84 := GenS8{F0: 0}
            fmt.Println(v43, v5)
            _ = v84
        }
        var v85 bool = v4
        v86 := GenS9{F0: 42, F1: (v1 - 12)}
        _ = v85
        _ = v86
    }
    fmt.Println(v43)
    switch v42 {
    case "":
        v2 = "foo bar"
    case "\n\t":
        var v87 uint32 = 2147483647
        _ = v87
    case "hello":
        if v4 {
            v2 = v2
        } else {
            switch (v42 + v2) {
            case "":
                for fi88 := 0; fi88 < 13; fi88++ {
                    v89 := GenS11{F0: 161, F1: uint8(v1), F2: 1}
                    _ = fi88
                    _ = v89
                }
            case "\n\t":
                switch v42 {
                case "hello":
                    var v90 float64 = 0.1
                    _ = v90
                }
            case "a":
                v2 = "a"
            }
        }
    default:
        fmt.Println(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
