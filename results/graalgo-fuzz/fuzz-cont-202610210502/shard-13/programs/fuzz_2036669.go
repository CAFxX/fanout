package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int8; F1 int16; F2 uint8 }
type GenS11 struct { F0 uint32; F1 int16 }
type GenS12 struct { F0 int32; F1 uint16; F2 int }
type GenS13 struct { F0 int; F1 uint }
func genFn0() uint64 {
    var v4 uint64 = 902
    var v5 int = 127
    var v6 uint16 = 136
    if (65535 <= 65) {
        if false {
            var v7 uint8 = 42
            v8 := []int{469}
            var v9 int = v5
            _ = v7
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 3; fi10++ {
                v5 = v5
                var v11 float32 = 0.0
                fmt.Println(v6, v11)
                _ = fi10
                _ = v11
            }
        }
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(fi12, v5)
            if true {
                v13 := GenS1{F0: 732, F1: 32767}
                _ = v13
            } else {
                v14 := GenS1{F0: 552, F1: -1}
                if false {
                    v15 := GenS2{F0: 127}
                    var v16 uint16 = (v6 + 800)
                    _ = v15
                    _ = v16
                }
                _ = v14
            }
            _ = fi12
        }
        switch "" {
        case "hello":
            v17 := GenS2{F0: 0}
            _ = v17
        case "":
            v18 := GenS1{F0: -1, F1: 1000}
            _ = v18
        case "foo bar":
            v4 = 42
        default:
            switch (v5 << ((602) & 63)) {
            case -735:
                v19 := GenS3{F0: 2}
                _ = v19
            case -9223372036854775808:
                var v20 int16 = 1000
                _ = v20
            default:
                for fi21 := 0; fi21 < 12; fi21++ {
                    for fi22 := 0; fi22 < 10; fi22++ {
                        var v23 string = "a"
                        var v24 uint8 = 128
                        fmt.Println(v24)
                        var v25 string = (v23 + v23)
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    _ = fi21
                }
            }
        }
        v26 := GenS1{F0: 32767, F1: -32768}
        _ = v26
    }
    v27 := GenS4{F0: (^(v6))}
    _ = v27
    _ = v4
    _ = v5
    _ = v6
    return v4
}
func genFn1(p0 uint32, p1 int8) uint {
    if false {
        fmt.Println(p0)
        v28 := []int{32767}
        _ = v28
    }
    v29 := []int{32767}
    var v30 uint32 = (^(p0))
    if false {
        v30 = (^(v30))
        if true {
            v30 = v30
        }
        var v31 int = -9223372036854775808
        _ = v31
    } else {
        var v32 uint64 = 0
        _ = v32
    }
    fmt.Println(p1, v30)
    _ = p0
    _ = p1
    _ = v29
    _ = v30
    return 4294967295
}

func main() {
    var v1 float32 = (0.1 + -2.25)
    v2 := []string{}
    v3 := make(map[int]uint32)
    var v33 int = -2147483648
    fmt.Println(v33, v1)
    v34 := GenS5{F0: -43, F1: 49}
    switch v33 {
    case 128:
        var v35 uint = genFn1(0, 42)
        _ = v35
    case -982:
        v36 := v3[32767]
        _ = v36
    case 179:
        v37 := GenS6{F0: -1, F1: 2147483648}
        _ = v37
    default:
        for fi38 := 0; fi38 < 15; fi38++ {
            v1 += v1
            for fi39 := 0; fi39 < 3; fi39++ {
                v33 += fi38
                if false {
                    delete(v3, (^(fi39)))
                    switch "" {
                    case "hello":
                        switch "foo bar" {
                        case "":
                            v33 = v33
                        case "hello":
                            if false {
                                v1 = 1.5
                            } else {
                                var v40 bool = ((v1 + v1) == ((1.5 * (0.1 / -2.25)) / (1e10 / v1)))
                                _ = v40
                            }
                            if (!(fi38 >= 9223372036854775807)) {
                                v41 := GenS7{F0: 0}
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = ("hello" + "a")
                                    }
                                    var v42 int64 = 286
                                    _ = v42
                                }
                                v43 := v3[-163]
                                _ = v41
                                _ = v43
                            } else {
                                v33 = (v33 % ((v33) | 1))
                                if false {
                                    v44 := GenS3{F0: 1}
                                    var v45 uint8 = 243
                                    var v46 int = fi38
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                            }
                            var v47 int64 = 0
                            v48 := GenS8{F0: "foo bar"}
                            v49 := GenS9{F0: (^(v33))}
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                    case "\n\t":
                        fmt.Println(fi39)
                    case "foo bar":
                        if false {
                            _ = len(v2)
                        }
                    default:
                        v33 -= -2147483648
                    }
                } else {
                    if (!(0 < -2147483648)) {
                        switch 294 {
                        case 9223372036854775807:
                            var v50 uint64 = 42
                            _ = v50
                        case 566:
                            if ((float32(0.0) - (v1 / v1)) < v1) {
                                if true {
                                    if true {
                                        var v51 int32 = -2147483648
                                        v52 := GenS10{F0: -48, F1: 42, F2: 30}
                                        _ = v51
                                        _ = v52
                                    }
                                    var v53 uint32 = 65535
                                    if false {
                                        fmt.Println(v53)
                                        _ = len(v2)
                                    }
                                    _ = v53
                                } else {
                                    switch ("" + (("foo bar" + "foo bar") + "hello")) {
                                    case "a":
                                        v1 += v1
                                    default:
                                        v1 += (v1 - -2.25)
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            } else {
                                if (!(1 >= 42)) {
                                    if (4294967295 != 2) {
                                        v54 := GenS10{F0: 80, F1: 1, F2: 0}
                                        _ = v54
                                    }
                                }
                            }
                        case 0:
                            if false {
                                if true {
                                    var v55 string = "foo bar"
                                    if ((!false) || true) {
                                        v56 := GenS11{F0: 2, F1: 73}
                                        _ = v56
                                    } else {
                                        v57 := v3[fi39]
                                        var v58 float64 = float64(0.1)
                                        if ((!false) && (v1 <= ((v1 / 1.5) + (v1 + (v1 / 1.0))))) {
                                            v59 := GenS11{F0: 2, F1: 1000}
                                            if true {
                                                switch 128 {
                                                case 128:
                                                    var v60 uint = 503
                                                    _ = v60
                                                }
                                                fmt.Println(fi39, v33)
                                            } else {
                                                if (-9223372036854775808 > -9223372036854775808) {
                                                    delete(v3, (fi39 >> ((0) & 63)))
                                                }
                                            }
                                            _ = v59
                                        } else {
                                            switch "hello" {
                                            case "a":
                                                if false {
                                                    fmt.Println(fi38, v57, fi39)
                                                    v2 = append(v2, v55)
                                                    var v61 int = (fi38 / ((2147483647) | 1))
                                                    _ = v61
                                                } else {
                                                    v1 = (((3.14159 * (1e10 - 0.1)) + v1) * (v1 - (v1 * (0.0 * -2.25))))
                                                    fmt.Println(v58, v57, v58)
                                                }
                                            case "foo bar":
                                                v33 = 32768
                                            case "hello":
                                                if (-2.25 < 1.5) {
                                                    _ = len(v2)
                                                }
                                            }
                                            var v62 float64 = ((v58 / 1.5) + v58)
                                            if true {
                                                v63 := GenS6{F0: -1, F1: (v57 << ((genFn1((v57 >> ((42) & 63)), -128)) & 63))}
                                                if len(v2) > 0 {
                                                    v2[0] = (v55 + v55)
                                                }
                                                v62 += v58
                                                _ = v63
                                            } else {
                                                if (!false) {
                                                    var v64 bool = (!false)
                                                    _ = v64
                                                } else {
                                                    v65 := v3[128]
                                                    v66 := GenS8{F0: ((v55 + "a") + (v55 + (("\n\t" + (v55 + v55)) + ((((("foo bar" + "hello") + v55) + ("foo bar" + ("a" + v55))) + ("hello" + (((v55 + (v55 + "\n\t")) + v55) + "foo bar"))) + ""))))}
                                                    _ = v65
                                                    _ = v66
                                                }
                                                switch ((("hello" + (v55 + v55)) + v55) + v55) {
                                                case "a":
                                                    if false {
                                                        if (!true) {
                                                            fmt.Println(v58, v62, v62)
                                                        }
                                                        var v67 int = fi38
                                                        _ = v67
                                                    }
                                                case "\n\t":
                                                    v57 -= 682
                                                }
                                            }
                                            var v68 float64 = -2.25
                                            v62 = (v68 / v68)
                                            _ = v62
                                            _ = v68
                                        }
                                        _ = v57
                                        _ = v58
                                    }
                                    fmt.Println(v55)
                                    _ = v55
                                }
                            }
                        }
                        var v69 int = (fi38 - v33)
                        switch ("foo bar" + "foo bar") {
                        case "foo bar":
                            if (128 != 85) {
                                v33 = -360
                                if false {
                                    fmt.Println(v69, v69, v1)
                                    v70 := GenS5{F0: 0, F1: 2}
                                    _ = v70
                                } else {
                                    v71 := v3[(-(v69))]
                                    _ = v71
                                }
                                switch v69 {
                                case 4294967295:
                                    switch "" {
                                    case "a":
                                        fmt.Println(v1, fi39, fi39)
                                    case "hello":
                                        _ = len(v3)
                                    default:
                                        var v72 float64 = 1.5
                                        _ = v72
                                    }
                                }
                            }
                            var v73 int = (fi39 & (v69 >> ((genFn1(42, 90)) & 63)))
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                            _ = v73
                        }
                        if (-2147483648 > 2147483647) {
                            var v74 int32 = -1
                            var v75 int = fi39
                            var v76 int32 = v74
                            _ = v74
                            _ = v75
                            _ = v76
                        } else {
                            if true {
                                v77 := GenS6{F0: 42, F1: 65535}
                                _ = v77
                            }
                            v78 := v3[v69]
                            v69 = v69
                            _ = v78
                        }
                        _ = v69
                    } else {
                        v3[9223372036854775807] = 65535
                        _ = len(v3)
                        var v79 float32 = float32((((0.0 * 3.14159) / -2.25) + (3.14159 + 1.5)))
                        _ = v79
                    }
                    v2 = append(v2, "a")
                    fmt.Println(fi38, v1)
                    var v80 uint64 = 305
                    _ = v80
                }
                v81 := GenS4{F0: 751}
                _ = fi39
                _ = v81
            }
            _ = fi38
        }
    }
    fmt.Println(v1)
    for fi82 := 0; fi82 < 12; fi82++ {
        v33 = 127
        if true {
            v33 = (fi82 % ((fi82) | 1))
            for fi83 := 0; fi83 < 20; fi83++ {
                fmt.Println(fi82, fi83, fi82)
                _ = fi83
            }
            for fi84 := 0; fi84 < 10; fi84++ {
                v3[(fi84 & fi84)] = 4294967295
                if false {
                    var v85 float64 = -2.25
                    var v86 int = fi82
                    switch (v86 | (-(fi84))) {
                    case -2147483648:
                        fmt.Println(v85, fi84)
                    case -9223372036854775808:
                        var v87 int8 = 127
                        _ = v87
                    }
                    var v88 bool = true
                    _ = v85
                    _ = v86
                    _ = v88
                }
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch v33 {
                    case 2147483647:
                        var v89 uint64 = 645
                        _ = v89
                    case -1:
                        v90 := GenS12{F0: 499, F1: 172, F2: (-(fi84))}
                        _ = v90
                    case -2147483648:
                        v91 := v3[-2147483648]
                        _ = v91
                    }
                } else {
                    fmt.Println(v33)
                    switch (fi84 * (v33 << ((255) & 63))) {
                    case 127:
                        v33 += (^(fi84))
                    case 42:
                        if (845 < 65535) {
                            v33 = fi84
                            var v92 float64 = (3.14159 - (-2.25 - (-2.25 - (0.1 + 0.0))))
                            _ = v92
                        }
                    case -9223372036854775808:
                        var v93 uint32 = 2147483647
                        _ = v93
                    default:
                        v94 := GenS13{F0: fi84, F1: 3}
                        _ = v94
                    }
                }
                var v95 uint32 = 371
                _ = fi84
                _ = v95
            }
        }
        v33 = (^(fi82))
        _ = fi82
    }
    var v96 float32 = 0.0
    switch (v33 | (-(v33))) {
    case 42:
        delete(v3, 1)
    default:
        for fi97 := 0; fi97 < 14; fi97++ {
            v98 := v3[4294967295]
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi97
            _ = v98
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v96
}
