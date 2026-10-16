package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint64; F1 uint16; F2 string }
type GenS3 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 int64; F1 int }
type GenS10 struct { F0 int64; F1 int8 }
func genFn0(p0 int) uint32 {
    var v7 float32 = 3.14159
    v8 := []int{}
    fmt.Println(v7, p0)
    var v9 float32 = (v7 - (0.0 - v7))
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := []int{-664, 128, 2147483647, -467}
    var v2 int32 = -406
    var v3 int64 = -372
    var v4 int8 = -1
    var v5 float64 = -2.25
    v6 := make(map[int]int)
    v3 = -290
    var v10 int32 = v2
    v11 := GenS1{F0: 1, F1: -578}
    v3 += 978
    var v12 int64 = v3
    var v13 uint8 = 245
    v10 -= (v2 ^ v10)
    v5 = 3.14159
    v14 := GenS1{F0: 42, F1: 32767}
    v5 = v5
    for fi15 := 0; fi15 < 11; fi15++ {
        switch "a" {
        case "foo bar":
            var v16 int = (fi15 | (fi15 * (-(fi15))))
            _ = v16
        case "hello":
            v17 := GenS1{F0: 2, F1: -2147483648}
            _ = v17
        default:
            if ("hello" != "\n\t") {
                v12 = v3
                _ = len(v6)
            }
        }
        switch "hello" {
        case "\n\t":
            if false {
                v18 := GenS2{F0: 18446744073709551615, F1: 65535, F2: ("\n\t" + "\n\t")}
                v13 = v13
                var v19 uint64 = 0
                _ = v18
                _ = v19
            } else {
                v10 = 42
            }
        default:
            if false {
                v20 := v6[fi15]
                if len(v1) > 0 {
                    v1[0] = v20
                }
                v21 := GenS1{F0: 1, F1: fi15}
                _ = v20
                _ = v21
            }
        }
        _ = fi15
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        if false {
            v23 := GenS1{F0: 0, F1: 0}
            _ = v23
        }
        _ = len(v6)
        _ = fi22
    }
    v13 = v13
    switch "\n\t" {
    case "":
        if false {
            v24 := GenS2{F0: 0, F1: 255, F2: "foo bar"}
            if ((-740 > -1) && true) {
                var v25 float64 = (((3.14159 + v5) - v5) + (1.5 - 3.14159))
                _ = len(v1)
                for fi26 := 0; fi26 < 2; fi26++ {
                    var v27 int64 = v3
                    v5 += v25
                    delete(v6, (fi26 & (fi26 - (fi26 | 4294967295))))
                    _ = fi26
                    _ = v27
                }
                _ = v25
            } else {
                for fi28 := 0; fi28 < 11; fi28++ {
                    var v29 bool = (!(true || true))
                    var v30 bool = ((fi28 + (-(fi28))) <= -6)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                fmt.Println(v10, v2)
            }
            _ = v24
        } else {
            if (true && ((!(!(true && (624 > 2147483647)))) || (!((2147483647 < 1) && false)))) {
                var v31 int = -108
                _ = v31
            } else {
                switch 32768 {
                case 32768:
                    for fi32 := 0; fi32 < 14; fi32++ {
                        for fi33 := 0; fi33 < 8; fi33++ {
                            var v34 uint = 0
                            var v35 uint16 = 284
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        _ = fi32
                    }
                case 123:
                    switch "foo bar" {
                    case "":
                        if ((v13 < (v13 >> ((1) & 63))) || true) {
                            v36 := GenS2{F0: 2, F1: 630, F2: "hello"}
                            v5 -= 3.14159
                            var v37 uint64 = 18446744073709551615
                            _ = v36
                            _ = v37
                        }
                    case "hello":
                        var v38 bool = (v10 == 0)
                        _ = v38
                    }
                case 454:
                    v39 := GenS3{F0: 785, F1: 42, F2: (v13 >> ((255) & 63))}
                    _ = v39
                }
                fmt.Println(v2, v10, v4)
                v40 := GenS1{F0: 4294967295, F1: 430}
                _ = v40
            }
        }
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            if (1 == 496) {
                var v41 uint = 2
                v1 = append(v1, 42)
                _ = v41
            } else {
                if ((v4 - (v4 - 56)) <= v4) {
                    switch "\n\t" {
                    case "\n\t":
                        if (42 <= 2) {
                            switch "" {
                            case "a":
                                var v42 uint = 1
                                _ = v42
                            case "foo bar":
                                for fi43 := 0; fi43 < 20; fi43++ {
                                    switch fi43 {
                                    case -924:
                                        v2 = (v2 / ((v2) | 1))
                                    case 128:
                                        v44 := GenS4{F0: (v10 * (^(v2))), F1: 4294967295, F2: 1}
                                        _ = v44
                                    }
                                    v4 += (^(v4))
                                    _ = fi43
                                }
                            case "hello":
                                v45 := v6[-911]
                                _ = v45
                            }
                            if (2 <= 255) {
                                switch (("foo bar" + ((("a" + (("a" + ("hello" + ("\n\t" + "hello"))) + "a")) + "hello") + ((("\n\t" + "hello") + "a") + "hello"))) + ("a" + (("foo bar" + "") + "\n\t"))) {
                                case "a":
                                    for fi46 := 0; fi46 < 3; fi46++ {
                                        var v47 uint8 = v13
                                        v6[-434] = (fi46 * 42)
                                        _ = fi46
                                        _ = v47
                                    }
                                case "foo bar":
                                    _ = len(v1)
                                default:
                                    for fi48 := 0; fi48 < 14; fi48++ {
                                        v13 += v13
                                        for fi49 := 0; fi49 < 9; fi49++ {
                                            if false {
                                                v1 = append(v1, (fi49 ^ fi48))
                                            } else {
                                                v13 += (v13 * (v13 & v13))
                                                var v50 bool = (2 == 2)
                                                v51 := GenS5{F0: 0}
                                                _ = v50
                                                _ = v51
                                            }
                                            v52 := v6[fi48]
                                            if (269 < 2147483647) {
                                                var v53 int16 = -761
                                                _ = v53
                                            }
                                            _ = fi49
                                            _ = v52
                                        }
                                        v54 := GenS5{F0: 663}
                                        _ = fi48
                                        _ = v54
                                    }
                                }
                                _ = len(v6)
                                for fi55 := 0; fi55 < 6; fi55++ {
                                    if (!(3.14159 != -2.25)) {
                                        var v56 int64 = (v12 * (v12 & v12))
                                        _ = v56
                                    }
                                    var v57 bool = (!false)
                                    fmt.Println(fi55)
                                    _ = fi55
                                    _ = v57
                                }
                            } else {
                                if true {
                                    fmt.Println(v4, v3, v2)
                                    if (v5 > v5) {
                                        var v58 int64 = v12
                                        for fi59 := 0; fi59 < 8; fi59++ {
                                            var v60 int16 = 0
                                            _ = fi59
                                            _ = v60
                                        }
                                        v5 = (float64(float32((float64((0.0 + 0.0)) * 0.1))) * 1.5)
                                        _ = v58
                                    } else {
                                        var v61 int8 = (v4 / ((v4) | 1))
                                        var v62 uint8 = (v13 & (v13 - v13))
                                        _ = v61
                                        _ = v62
                                    }
                                    v63 := GenS3{F0: 941, F1: (^(v4)), F2: (^(v13))}
                                    _ = v63
                                } else {
                                    v3 = (v12 & v3)
                                    if (1e10 != 3.14159) {
                                        for fi64 := 0; fi64 < 3; fi64++ {
                                            v1 = append(v1, 0)
                                            fmt.Println(v4)
                                            _ = fi64
                                        }
                                    } else {
                                        v4 -= 127
                                        var v65 uint = 18446744073709551615
                                        _ = v65
                                    }
                                    v5 = 1e10
                                }
                                v66 := GenS6{F0: 315}
                                fmt.Println(v12, v2, v10)
                                _ = v66
                            }
                        }
                    }
                    for fi67 := 0; fi67 < 1; fi67++ {
                        v68 := v6[fi67]
                        var v69 uint16 = 0
                        if true {
                            if false {
                                if false {
                                    if (-32768 > 221) {
                                        delete(v6, fi67)
                                    } else {
                                        if false {
                                            switch 128 {
                                            case 2147483647:
                                                fmt.Println(fi67, v5)
                                            case -550:
                                                for fi70 := 0; fi70 < 1; fi70++ {
                                                    v71 := GenS7{F0: 42}
                                                    _ = fi70
                                                    _ = v71
                                                }
                                            case -9223372036854775808:
                                                v72 := GenS7{F0: 20}
                                                _ = v72
                                            }
                                        }
                                        for fi73 := 0; fi73 < 18; fi73++ {
                                            switch ("a" + "foo bar") {
                                            case "\n\t":
                                                if true {
                                                    var v74 int8 = v4
                                                    var v75 string = ("hello" + "")
                                                    _ = v74
                                                    _ = v75
                                                }
                                            default:
                                                v12 = v3
                                                v5 = (v5 * 1.5)
                                                v1 = append(v1, 2147483647)
                                            }
                                            _ = fi73
                                        }
                                        v68 = (fi67 & v68)
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = 128
                                    }
                                } else {
                                    if true {
                                        var v76 string = "a"
                                        var v77 uint = 474
                                        _ = v76
                                        _ = v77
                                    } else {
                                        fmt.Println(v3, v12)
                                        v69 = v69
                                    }
                                    v10 = (v10 + v2)
                                    for fi78 := 0; fi78 < 9; fi78++ {
                                        if true {
                                            v2 = (v2 ^ v2)
                                            if ((int16(v3) == -1) || (432 != 1)) {
                                                v13 = (^(v13))
                                            } else {
                                                var v79 int64 = v3
                                                v69 -= 2
                                                _ = v79
                                            }
                                            v3 += v12
                                        }
                                        var v80 uint16 = (v69 + v69)
                                        _ = fi78
                                        _ = v80
                                    }
                                }
                            }
                            for fi81 := 0; fi81 < 6; fi81++ {
                                var v82 bool = (v2 != (v2 << ((80) & 63)))
                                var v83 bool = (false || (!(42 > 4294967295)))
                                v1 = append(v1, (v68 / ((fi67) | 1)))
                                fmt.Println(v10, fi81, v82)
                                _ = fi81
                                _ = v82
                                _ = v83
                            }
                            v84 := GenS6{F0: uint64(v13)}
                            _ = v84
                        } else {
                            if (!(502 != 34)) {
                                v85 := v6[422]
                                _ = v85
                            } else {
                                v86 := GenS8{F0: false, F1: int8(v69)}
                                if true {
                                    var v87 int32 = v10
                                    for fi88 := 0; fi88 < 19; fi88++ {
                                        fmt.Println(v69, v5, fi88)
                                        v3 = v12
                                        _ = fi88
                                    }
                                    var v89 int32 = v10
                                    _ = v87
                                    _ = v89
                                }
                                v90 := GenS8{F0: (false && true), F1: 127}
                                _ = v86
                                _ = v90
                            }
                            var v91 uint32 = 4294967295
                            switch -9223372036854775808 {
                            case 128:
                                v92 := GenS1{F0: 4294967295, F1: (fi67 - fi67)}
                                _ = v92
                            }
                            _ = v91
                        }
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                    v93 := GenS2{F0: 321, F1: 32767, F2: ("foo bar" + "hello")}
                    _ = v93
                } else {
                    for fi94 := 0; fi94 < 13; fi94++ {
                        switch "\n\t" {
                        case "\n\t":
                            if (334 < 42) {
                                for fi95 := 0; fi95 < 11; fi95++ {
                                    if true {
                                        if (0 < 604) {
                                            var v96 int = fi94
                                            _ = v96
                                        } else {
                                            var v97 uint64 = 2
                                            v5 = v5
                                            if false {
                                                v2 = v2
                                            }
                                            _ = v97
                                        }
                                        fmt.Println(v12, v2)
                                    }
                                    switch "" {
                                    case "hello":
                                        v5 -= v5
                                    case "":
                                        v98 := GenS9{F0: -826, F1: fi94}
                                        _ = v98
                                    case "foo bar":
                                        fmt.Println(v5)
                                    default:
                                        var v99 uint = 282
                                        _ = v99
                                    }
                                    var v100 float64 = 0.1
                                    _ = fi95
                                    _ = v100
                                }
                            } else {
                                var v101 float64 = ((1e10 * (0.0 - ((v5 + v5) + (float64(float32(1.5)) + (v5 - v5))))) - 1.5)
                                fmt.Println(v10, v3, v3)
                                _ = v101
                            }
                        case "hello":
                            v10 = (v10 / ((v10) | 1))
                        case "foo bar":
                            var v102 uint = 18446744073709551615
                            _ = v102
                        default:
                            switch (fi94 | (fi94 << ((255) & 63))) {
                            case -1:
                                for fi103 := 0; fi103 < 1; fi103++ {
                                    v13 += (v13 / ((v13) | 1))
                                    v104 := GenS10{F0: -548, F1: v4}
                                    v2 = -999
                                    _ = fi103
                                    _ = v104
                                }
                            case -424:
                                v1 = append(v1, int(v2))
                            default:
                                var v105 int64 = v3
                                _ = v105
                            }
                        }
                        _ = fi94
                    }
                }
                delete(v6, 793)
            }
        }
    }
    var v106 uint = 394
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v106:", v106)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v106
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
