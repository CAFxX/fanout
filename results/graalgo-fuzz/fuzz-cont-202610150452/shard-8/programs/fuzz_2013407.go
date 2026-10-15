package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 int; F2 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int32 }
func genFn0(p0 int16, p1 uint16) bool {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: 127, F1: 4294967295, F2: 255}
        var v6 uint64 = 2
        if true {
            v6 = 1
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(p1, p0)
    if (true || (1 <= 0)) {
        if false {
            if (18446744073709551615 >= 65535) {
                if (0.1 <= -2.25) {
                    if false {
                        v7 := GenS1{F0: 255, F1: 128, F2: 4294967295}
                        switch (("a" + "a") + "foo bar") {
                        case "\n\t":
                            var v8 uint = 1
                            _ = v8
                        }
                        fmt.Println(p1, p0)
                        _ = v7
                    }
                    v9 := GenS2{F0: "hello", F1: 866}
                    v10 := []int{-2147483648, -2147483648, -671}
                    _ = v10
                    _ = v9
                } else {
                    var v11 int32 = -1
                    if (-1 == 128) {
                        p1 = 255
                        fmt.Println(p0, v11, p1)
                    }
                    _ = v11
                }
            } else {
                p1 = (^(p1))
            }
        }
        if (42 <= 42) {
            switch -866 {
            case 521:
                p1 = 1
            case -1:
                if ((18446744073709551615 <= 879) || (207 <= 1)) {
                    var v12 uint32 = 1
                    if (((1 > 1) || ((0 != 42) || (v12 != 394))) || true) {
                        v13 := GenS3{F0: 1}
                        v14 := []int{127}
                        var v15 int = -9223372036854775808
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    var v16 int32 = 984
                    _ = v12
                    _ = v16
                }
            case -2147483648:
                var v17 int8 = 1
                _ = v17
            default:
                v18 := []int{32767}
                _ = v18
            }
            var v19 uint8 = 176
            fmt.Println(v19, v19, p0)
            _ = v19
        } else {
            v20 := GenS3{F0: 127}
            if (2147483647 != 534) {
                v21 := GenS3{F0: 42}
                p0 += (p0 | p0)
                var v22 string = "foo bar"
                _ = v21
                _ = v22
            } else {
                v23 := GenS4{F0: 2147483647}
                v24 := []int{-66}
                _ = v23
                _ = v24
            }
            v25 := GenS1{F0: 173, F1: 577, F2: 586}
            _ = v20
            _ = v25
        }
    } else {
        fmt.Println(p0)
        if true {
            for fi26 := 0; fi26 < 4; fi26++ {
                v27 := []int{fi26}
                var v28 float32 = ((3.14159 * -2.25) - (float32(0.0) + (0.1 * 1.5)))
                for fi29 := 0; fi29 < 14; fi29++ {
                    p1 -= (p1 ^ 255)
                    switch -9223372036854775808 {
                    case 904:
                        var v30 int32 = -109
                        _ = v30
                    }
                    if len(v27) > 0 {
                        v27[0] = fi29
                    }
                    _ = fi29
                }
                _ = fi26
                _ = v27
                _ = v28
            }
            p1 += (p1 ^ (p1 * 255))
        }
        p0 += p0
    }
    switch "foo bar" {
    case "a":
        for fi31 := 0; fi31 < 14; fi31++ {
            for fi32 := 0; fi32 < 4; fi32++ {
                v33 := []int{32768, (fi31 | (fi32 & (-(fi31)))), (-(fi32))}
                _ = fi32
                _ = v33
            }
            fmt.Println(fi31, p0, p0)
            var v34 bool = false
            _ = fi31
            _ = v34
        }
    case "\n\t":
        var v35 uint32 = 2147483648
        _ = v35
    }
    _ = p0
    _ = p1
    return true
}
func genFn1(p0 uint8) uint8 {
    fmt.Println(p0)
    if true {
        if false {
            p0 -= (p0 / ((p0) | 1))
            var v36 int16 = -1
            fmt.Println(p0)
            _ = v36
        } else {
            var v37 float32 = 0.0
            v37 -= -2.25
            switch ("hello" + "foo bar") {
            case "\n\t":
                for fi38 := 0; fi38 < 6; fi38++ {
                    p0 = 0
                    v39 := GenS3{F0: (p0 ^ (p0 / ((1) | 1)))}
                    _ = fi38
                    _ = v39
                }
            case "foo bar":
                switch "hello" {
                case "hello":
                    fmt.Println(p0, p0)
                case "":
                    var v40 uint16 = 405
                    _ = v40
                case "\n\t":
                    if (false && ((!false) && true)) {
                        for fi41 := 0; fi41 < 4; fi41++ {
                            if ("" > "hello") {
                                var v42 uint16 = 1
                                fmt.Println(p0, v42, fi41)
                                v43 := GenS4{F0: 42}
                                _ = v42
                                _ = v43
                            }
                            var v44 float32 = v37
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(fi41, p0, p0)
                            case "hello":
                                if (true && (true || true)) {
                                    for fi45 := 0; fi45 < 4; fi45++ {
                                        var v46 int = (fi41 % ((-154) | 1))
                                        _ = fi45
                                        _ = v46
                                    }
                                    switch (fi41 | fi41) {
                                    case 4294967295:
                                        p0 = 1
                                    case -9223372036854775808:
                                        v47 := GenS5{F0: 42}
                                        _ = v47
                                    default:
                                        v37 -= v44
                                    }
                                }
                            default:
                                v37 = v37
                            }
                            _ = fi41
                            _ = v44
                        }
                        for fi48 := 0; fi48 < 16; fi48++ {
                            var v49 uint8 = p0
                            for fi50 := 0; fi50 < 6; fi50++ {
                                if (-1 <= 42) {
                                    if (true && false) {
                                        switch (fi50 + fi50) {
                                        case 128:
                                            fmt.Println(fi48)
                                        }
                                        if false {
                                            var v51 bool = true
                                            switch (fi48 ^ -772) {
                                            case -869:
                                                if v51 {
                                                    switch ("hello" + ("a" + ("hello" + "\n\t"))) {
                                                    case "\n\t":
                                                        v52 := []int{}
                                                        _ = v52
                                                    case "hello":
                                                        v53 := []int{}
                                                        if v51 {
                                                            var v54 int = fi48
                                                            _ = v54
                                                        }
                                                        _ = v53
                                                    }
                                                    v55 := GenS6{F0: true, F1: 9223372036854775807, F2: 1}
                                                    v49 -= 255
                                                    _ = v55
                                                } else {
                                                    v56 := GenS7{F0: 2}
                                                    _ = v56
                                                }
                                            case 726:
                                                var v57 int64 = -9223372036854775808
                                                _ = v57
                                            }
                                            if v51 {
                                                var v58 int16 = 0
                                                if v51 {
                                                    var v59 float64 = ((0.0 / ((-2.25 / 1e10) - (-2.25 + float64((1e10 / v37))))) + 0.0)
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v60 string = "hello"
                                                        _ = v60
                                                    }
                                                    v59 = (3.14159 - -2.25)
                                                    _ = v59
                                                } else {
                                                    var v61 uint32 = 514
                                                    _ = v61
                                                }
                                                if v51 {
                                                    v62 := GenS8{F0: 2147483648, F1: v58}
                                                    v63 := GenS6{F0: v51, F1: fi48, F2: int64(fi48)}
                                                    v64 := []int{}
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                }
                                                _ = v58
                                            }
                                            _ = v51
                                        }
                                    } else {
                                        if false {
                                            var v65 bool = false
                                            v65 = v65
                                            _ = v65
                                        } else {
                                            fmt.Println(v37, fi50, v37)
                                            var v66 uint64 = 0
                                            if ((true && (0 <= (fi48 + fi50))) && true) {
                                                fmt.Println(v66, fi50)
                                                if ((float64(1.5) + 0.0) > (-2.25 * (0.1 - 0.1))) {
                                                    v67 := GenS4{F0: 83}
                                                    _ = v67
                                                } else {
                                                    v68 := GenS9{F0: 435, F1: -2147483648}
                                                    v69 := GenS4{F0: 2147483647}
                                                    v70 := []int{(fi48 + -2147483648), 767, fi48, fi50}
                                                    _ = v68
                                                    _ = v69
                                                    _ = v70
                                                }
                                                v71 := []int{}
                                                _ = v71
                                            } else {
                                                fmt.Println(fi50)
                                                var v72 float64 = 1e10
                                                if true {
                                                    v73 := GenS1{F0: (v49 * 255), F1: fi50, F2: 18446744073709551615}
                                                    var v74 int8 = 95
                                                    var v75 uint8 = p0
                                                    _ = v73
                                                    _ = v74
                                                    _ = v75
                                                }
                                                switch "" {
                                                case "\n\t":
                                                    switch -9223372036854775808 {
                                                    case 2147483647:
                                                        v76 := []int{(fi48 + (fi50 << ((255) & 63))), (fi48 >> ((18446744073709551615) & 63)), (fi50 - fi48), (fi48 ^ -259)}
                                                        _ = v76
                                                    case 105:
                                                        v37 += (1e10 * 1.5)
                                                    }
                                                default:
                                                    var v77 uint64 = v66
                                                    _ = v77
                                                }
                                                _ = v72
                                            }
                                            var v78 uint = 4294967295
                                            _ = v66
                                            _ = v78
                                        }
                                        if false {
                                            var v79 uint8 = (v49 - (^(v49)))
                                            _ = v79
                                        } else {
                                            var v80 string = "a"
                                            _ = v80
                                        }
                                        v81 := GenS9{F0: 1000, F1: -1}
                                        v49 = v49
                                        var v82 int64 = -1
                                        _ = v81
                                        _ = v82
                                    }
                                }
                                var v83 string = ""
                                _ = fi50
                                _ = v83
                            }
                            if (18446744073709551615 < 1) {
                                for fi84 := 0; fi84 < 14; fi84++ {
                                    var v85 int16 = -172
                                    var v86 int64 = -9223372036854775808
                                    v87 := GenS1{F0: p0, F1: fi48, F2: 2}
                                    _ = fi84
                                    _ = v85
                                    _ = v86
                                    _ = v87
                                }
                                p0 += p0
                            }
                            _ = fi48
                            _ = v49
                        }
                    }
                }
            }
            _ = v37
        }
    } else {
        v88 := GenS10{F0: -729}
        _ = v88
    }
    _ = p0
    return 1
}

func main() {
    v1 := []int{9223372036854775807, -855, -9223372036854775808}
    var v2 bool = true
    v3 := make(map[string]int)
    _ = len(v1)
    v89 := v3[("hello" + "")]
    var v90 uint8 = 224
    _ = len(v1)
    if ((2 < 0) || (!false)) {
        switch "hello" {
        case "\n\t":
            if (!(-1 < -1)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v91 int16 = -1
                _ = v91
            }
        case "a":
            v92 := GenS11{F0: 2147483648}
            _ = v92
        case "":
            if (985 > -314) {
                v93 := GenS6{F0: (255 < 0), F1: (v89 << ((1) & 63)), F2: 312}
                v1 = append(v1, 32767)
                _ = v93
            } else {
                v2 = (157 != 42)
            }
        }
    } else {
        v89 = v89
    }
    fmt.Println(v89, v90)
    fmt.Println(v90)
    for fi94 := 0; fi94 < 1; fi94++ {
        v95 := GenS12{F0: 0}
        _ = fi94
        _ = v95
    }
    v90 = genFn1(v90)
    v90 -= (v90 - v90)
    fmt.Println(v2)
    v2 = (!v2)
    fmt.Println(v90, v89, v89)
    v96 := GenS13{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v89
    _ = v90
    _ = v96
}
