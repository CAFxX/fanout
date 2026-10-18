package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 int8; F2 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8; F1 uint16; F2 string }
type GenS9 struct { F0 int16; F1 int8; F2 int }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int; F1 int64 }
type GenS12 struct { F0 uint16; F1 int32; F2 int8 }
func genFn0(p0 string, p1 bool, p2 uint16) uint64 {
    var v7 uint64 = 0
    switch (("hello" + p0) + (p0 + p0)) {
    case "hello":
        p1 = p1
    case "foo bar":
        switch "" {
        case "\n\t":
            v8 := GenS1{F0: 62}
            _ = v8
        case "":
            fmt.Println(p0)
        case "foo bar":
            v9 := GenS2{F0: 0}
            _ = v9
        }
    case "a":
        for fi10 := 0; fi10 < 20; fi10++ {
            if (((0.0 - 1e10) > ((1.5 / 0.1) + -2.25)) && ((float32(3.14159) * 1.5) != ((1e10 * 0.1) + 3.14159))) {
                if (-2.25 == 1.5) {
                    var v11 int32 = -2147483648
                    v11 = (-(v11))
                    var v12 bool = ((v11 * -2147483648) <= int32(v7))
                    _ = v11
                    _ = v12
                }
            } else {
                switch 9223372036854775807 {
                case 32767:
                    v13 := []int{4294967295, fi10}
                    _ = v13
                case -1:
                    var v14 int8 = 127
                    _ = v14
                case 135:
                    for fi15 := 0; fi15 < 5; fi15++ {
                        p0 = "hello"
                        fmt.Println(p1, p2, fi10)
                        _ = fi15
                    }
                }
            }
            var v16 uint = 0
            v17 := GenS1{F0: (p2 % ((1) | 1))}
            _ = fi10
            _ = v16
            _ = v17
        }
    }
    v7 = v7
    if p1 {
        v18 := GenS3{F0: 759, F1: -1, F2: -32768}
        _ = v18
    }
    if p1 {
        v19 := GenS2{F0: 65535}
        for fi20 := 0; fi20 < 18; fi20++ {
            v7 -= (v7 | 18446744073709551615)
            var v21 uint = 42
            _ = fi20
            _ = v21
        }
        if (!(131 <= 5)) {
            v22 := []int{127}
            if (42 >= 42) {
                _ = len(v22)
                if (128 < 1) {
                    if p1 {
                        if p1 {
                            for fi23 := 0; fi23 < 15; fi23++ {
                                for fi24 := 0; fi24 < 11; fi24++ {
                                    fmt.Println(p0, fi23)
                                    var v25 int8 = -1
                                    var v26 string = ((p0 + "foo bar") + (p0 + (("" + p0) + ((("foo bar" + ((p0 + "foo bar") + "foo bar")) + "\n\t") + p0))))
                                    p2 = p2
                                    switch ("foo bar" + v26) {
                                    case "hello":
                                        switch fi23 {
                                        case -2147483648:
                                            p1 = (1 >= 255)
                                        case 128:
                                            v26 = "\n\t"
                                        default:
                                            v27 := GenS4{F0: 18446744073709551615}
                                            _ = v27
                                        }
                                        if p1 {
                                            fmt.Println(fi24, v7, p0)
                                        } else {
                                            if len(v22) >= 2 {
                                                v22 = v22[1:len(v22)]
                                            }
                                            var v28 uint16 = p2
                                            switch v26 {
                                            case "hello":
                                                if p1 {
                                                    switch (fi23 ^ fi23) {
                                                    case 32768:
                                                        var v29 int8 = 100
                                                        _ = v29
                                                    case -9223372036854775808:
                                                        v26 = p0
                                                    default:
                                                        if len(v22) >= 2 {
                                                            v22 = v22[1:len(v22)]
                                                        }
                                                    }
                                                    p2 += 80
                                                    switch "" {
                                                    case "foo bar":
                                                        v7 += 42
                                                    case "a":
                                                        p2 += p2
                                                    }
                                                }
                                            case "\n\t":
                                                var v30 uint32 = 42
                                                _ = v30
                                            default:
                                                fmt.Println(v28)
                                            }
                                            _ = v28
                                        }
                                        p1 = (v7 != 42)
                                    case "":
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                    case "\n\t":
                                        fmt.Println(v7)
                                    }
                                    p0 = p0
                                    switch fi24 {
                                    case -1:
                                        var v31 string = p0
                                        _ = v31
                                    case 127:
                                        fmt.Println(fi23, fi24)
                                    }
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                                _ = fi23
                            }
                            v32 := GenS5{F0: -683, F1: 122, F2: 298}
                            _ = v32
                        }
                        switch "foo bar" {
                        case "\n\t":
                            v22 = append(v22, -1)
                        }
                    } else {
                        switch 128 {
                        case -568:
                            p2 = (p2 + p2)
                        case 42:
                            var v33 string = "hello"
                            _ = v33
                        }
                        var v34 uint64 = (^(v7))
                        p2 += (^(p2))
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 16; fi35++ {
                        p2 = 65535
                        var v36 int16 = 847
                        v37 := GenS6{F0: (v7 | 348)}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    var v38 string = p0
                    _ = v38
                } else {
                    var v39 string = p0
                    switch -622 {
                    case 0:
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                    default:
                        var v40 int64 = -1
                        _ = v40
                    }
                    if (((!p1) || (4294967295 <= 564)) || p1) {
                        fmt.Println(v39)
                        var v41 uint = 18446744073709551615
                        _ = len(v22)
                        _ = v41
                    } else {
                        v22 = append(v22, 621)
                    }
                    _ = v39
                }
                p2 += uint16(v7)
            } else {
                switch 127 {
                case -123:
                    fmt.Println(v7, p1)
                default:
                    v7 += 42
                }
                for fi42 := 0; fi42 < 16; fi42++ {
                    if (0 < 32767) {
                        if (((p0 + p0) + (p0 + p0)) == "foo bar") {
                            if len(v22) > 0 {
                                v22[0] = (fi42 ^ (fi42 >> ((222) & 63)))
                            }
                        }
                    }
                    var v43 float64 = (-2.25 * 0.0)
                    if ((983 >= 874) || (-122 != 0)) {
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                    }
                    _ = fi42
                    _ = v43
                }
            }
            var v44 string = p0
            _ = v22
            _ = v44
        }
        _ = v19
    } else {
        switch -123 {
        case -9223372036854775808:
            var v45 int64 = 889
            _ = v45
        case 0:
            p2 = p2
        case -388:
            for fi46 := 0; fi46 < 1; fi46++ {
                for fi47 := 0; fi47 < 8; fi47++ {
                    v48 := GenS7{F0: 1}
                    if true {
                        var v49 int32 = 0
                        fmt.Println(v49)
                        var v50 int32 = v49
                        _ = v49
                        _ = v50
                    } else {
                        p1 = p1
                    }
                    _ = fi47
                    _ = v48
                }
                _ = fi46
            }
        }
        if p1 {
            var v51 int8 = 0
            _ = v51
        }
        if p1 {
            v7 += 83
            switch "\n\t" {
            case "":
                p1 = p1
            default:
                v52 := GenS3{F0: 1000, F1: -396, F2: 0}
                _ = v52
            }
            v53 := GenS4{F0: 553}
            _ = v53
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 string, p1 uint32, p2 uint) bool {
    switch "hello" {
    case "hello":
        v54 := GenS8{F0: 0, F1: 587, F2: ""}
        _ = v54
    }
    for fi55 := 0; fi55 < 9; fi55++ {
        v56 := GenS1{F0: 263}
        for fi57 := 0; fi57 < 10; fi57++ {
            if false {
                var v58 float64 = ((1.5 - ((3.14159 + -2.25) + 0.1)) + 1e10)
                _ = v58
            }
            v59 := GenS9{F0: 42, F1: 42, F2: fi55}
            _ = fi57
            _ = v59
        }
        _ = fi55
        _ = v56
    }
    p2 += p2
    if false {
        var v60 int32 = 591
        var v61 float32 = (1e10 * (1.5 - (3.14159 * 3.14159)))
        v62 := []int{-9223372036854775808, 2147483647, -1, 2147483647}
        _ = v60
        _ = v61
        _ = v62
    }
    var v63 int = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v63
    return (!(!true))
}

func main() {
    v1 := []uint32{65535}
    var v2 int64 = 42
    var v3 float64 = 0.0
    var v4 float32 = 1e10
    v5 := make(map[string]string)
    v6 := make(map[int]string)
    switch "hello" {
    case "foo bar":
        if false {
            v64 := GenS3{F0: -32768, F1: 912, F2: 465}
            var v65 int = 32767
            _ = v64
            _ = v65
        }
    case "a":
        if true {
            v66 := GenS10{F0: 0}
            _ = v66
        } else {
            for fi67 := 0; fi67 < 12; fi67++ {
                if true {
                    fmt.Println(v4, v2)
                    if (((true && false) || ((false || true) || true)) && false) {
                        if (1000 < 1000) {
                            v3 -= 3.14159
                        }
                    } else {
                        _ = len(v1)
                        v1 = append(v1, 1)
                        fmt.Println(v4, v2)
                    }
                    v68 := GenS7{F0: 2}
                    _ = v68
                } else {
                    v2 -= 1
                    v69 := v5[(("" + ("a" + (("hello" + ("foo bar" + (("foo bar" + "") + ""))) + "hello"))) + "hello")]
                    _ = v69
                }
                v4 += 1e10
                fmt.Println(v2, fi67)
                _ = fi67
            }
            v70 := GenS1{F0: 32767}
            _ = v70
        }
    default:
        if true {
            if ("foo bar" != ("a" + ("" + ("foo bar" + "")))) {
                var v71 int64 = v2
                _ = v71
            }
            for fi72 := 0; fi72 < 2; fi72++ {
                v73 := GenS11{F0: 32768, F1: (v2 * 1)}
                _ = fi72
                _ = v73
            }
        }
    }
    for fi74 := 0; fi74 < 5; fi74++ {
        switch (((("\n\t" + "hello") + "a") + "foo bar") + "hello") {
        case "a":
            if false {
                var v75 float32 = 0.1
                v4 = ((v4 + ((v4 * v4) - 0.1)) * 1.5)
                _ = v75
            }
        case "hello":
            if false {
                if false {
                    for fi76 := 0; fi76 < 8; fi76++ {
                        fmt.Println(v3, v3)
                        _ = fi76
                    }
                    fmt.Println(v3)
                    v4 -= ((-2.25 * v4) - 0.0)
                }
                var v77 uint = 42
                _ = v77
            } else {
                v78 := GenS12{F0: 2, F1: 1, F2: 42}
                var v79 int16 = -1
                v2 = v2
                _ = v78
                _ = v79
            }
        default:
            v2 = v2
        }
        _ = len(v6)
        var v80 uint = 255
        _ = fi74
        _ = v80
    }
    var v81 bool = ((!(70 != 2)) || true)
    if (0 != 127) {
        switch 1 {
        case 1:
            v82 := GenS1{F0: 65535}
            _ = v82
        default:
            if v81 {
                for fi83 := 0; fi83 < 2; fi83++ {
                    fmt.Println(v3)
                    _ = fi83
                }
            } else {
                v4 = (v4 / 0.1)
            }
        }
    } else {
        var v84 uint16 = 696
        for fi85 := 0; fi85 < 17; fi85++ {
            v84 -= (v84 + 240)
            for fi86 := 0; fi86 < 7; fi86++ {
                v87 := GenS5{F0: -1, F1: -16, F2: 334}
                v81 = ((v81 || (493 > 1)) && false)
                switch "\n\t" {
                case "":
                    var v88 float64 = float64((0.0 * ((((v4 / v4) + v4) / v4) * 1.5)))
                    _ = v88
                default:
                    fmt.Println(v84, fi86)
                }
                _ = fi86
                _ = v87
            }
            _ = fi85
        }
        v89 := v5[("hello" + "\n\t")]
        _ = v84
        _ = v89
    }
    var v90 uint64 = 1
    for fi91 := 0; fi91 < 6; fi91++ {
        v92 := GenS3{F0: 452, F1: fi91, F2: 509}
        _ = fi91
        _ = v92
    }
    var v93 uint = 42
    v94 := GenS3{F0: -104, F1: 0, F2: -32768}
    for fi95 := 0; fi95 < 5; fi95++ {
        fmt.Println(v2, v81)
        _ = fi95
    }
    v2 -= v2
    var v96 int16 = 836
    v3 += v3
    v93 -= v93
    fmt.Println(v81, v90, v81)
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v81:", v81)
    fmt.Println("v90:", v90)
    fmt.Println("v93:", v93)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v90
    _ = v93
    _ = v94
    _ = v96
}
