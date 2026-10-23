package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int8; F1 bool; F2 int64 }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS11 struct { F0 int64 }
func genFn0(p0 uint32, p1 int32, p2 uint8) int32 {
    fmt.Println(p2, p1)
    var v6 int8 = -128
    switch "hello" {
    case "foo bar":
        var v7 uint32 = p0
        _ = v7
    case "hello":
        var v8 bool = false
        _ = v8
    default:
        fmt.Println(p1, p1)
    }
    if true {
        if false {
            for fi9 := 0; fi9 < 16; fi9++ {
                p0 = p0
                v10 := GenS1{F0: 42}
                v11 := []int{}
                _ = fi9
                _ = v10
                _ = v11
            }
            p0 -= (p0 >> ((2) & 63))
            if (true && (798 <= 1)) {
                p2 = 115
            } else {
                var v12 int8 = 100
                switch -2147483648 {
                case 32768:
                    p0 += (^(p0))
                case 0:
                    v13 := GenS2{F0: -943, F1: p0}
                    _ = v13
                }
                _ = v12
            }
        } else {
            var v14 int64 = 42
            switch -9223372036854775808 {
            case -805:
                var v15 bool = true
                _ = v15
            case -785:
                if false {
                    if (!false) {
                        v16 := []int{42}
                        fmt.Println(p2)
                        if true {
                            for fi17 := 0; fi17 < 6; fi17++ {
                                v14 -= 415
                                var v18 uint64 = 369
                                _ = fi17
                                _ = v18
                            }
                        } else {
                            if true {
                                v14 = 1
                            } else {
                                fmt.Println(v14, p0)
                                _ = len(v16)
                            }
                        }
                        _ = v16
                    }
                }
            case -352:
                switch 2147483647 {
                case 127:
                    switch 168 {
                    case 864:
                        var v19 int16 = 32767
                        _ = v19
                    case -737:
                        v20 := []int{-1, -488, 32768}
                        _ = v20
                    }
                case 33:
                    v21 := []int{8}
                    _ = v21
                default:
                    for fi22 := 0; fi22 < 17; fi22++ {
                        p2 -= uint8(p0)
                        for fi23 := 0; fi23 < 14; fi23++ {
                            if false {
                                switch fi23 {
                                case -536:
                                    fmt.Println(p1, v6)
                                case 9223372036854775807:
                                    var v24 float64 = ((0.1 * ((float64((3.14159 * (float32(((1.5 - -2.25) + (3.14159 + 3.14159))) * -2.25))) + (((0.0 / 3.14159) - 0.1) - 1e10)) / ((-2.25 / 1.0) / ((3.14159 + 0.1) * (3.14159 - 0.0))))) + ((float64(0.0) - 1e10) / 1e10))
                                    _ = v24
                                }
                            }
                            if false {
                                fmt.Println(p0)
                                v14 += (v14 ^ (v14 % (((v14 ^ (v14 | v14))) | 1)))
                            } else {
                                var v25 uint8 = p2
                                var v26 int64 = v14
                                _ = v25
                                _ = v26
                            }
                            p0 -= p0
                            _ = fi23
                        }
                        _ = fi22
                    }
                }
            default:
                if true {
                    v27 := []int{128, 32767, 1, 127}
                    if true {
                        var v28 float64 = (1e10 - float64(float32(1.5)))
                        fmt.Println(v28, v28, p1)
                        fmt.Println(v14, v28, v14)
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 20; fi29++ {
                        p0 += 0
                        v6 -= v6
                        _ = fi29
                    }
                    _ = v27
                }
            }
            v30 := GenS3{F0: 1, F1: 813}
            _ = v14
            _ = v30
        }
        var v31 uint8 = 127
        var v32 bool = (32768 >= 32768)
        _ = v31
        _ = v32
    } else {
        if false {
            v33 := GenS4{F0: 65535, F1: int16(v6)}
            _ = v33
        } else {
            v34 := GenS5{F0: 79, F1: (v6 < v6), F2: 0}
            _ = v34
        }
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1(p0 int, p1 uint16, p2 uint8) uint32 {
    var v35 float64 = 0.1
    if (true && false) {
        for fi36 := 0; fi36 < 13; fi36++ {
            var v37 int64 = 1
            _ = fi36
            _ = v37
        }
        v38 := GenS6{F0: -115, F1: -1}
        for fi39 := 0; fi39 < 16; fi39++ {
            var v40 uint16 = p1
            _ = fi39
            _ = v40
        }
        _ = v38
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    return 995
}

func main() {
    var v1 int8 = -1
    v2 := make(map[string]string)
    v3 := []uint32{65535}
    var v4 int32 = 0
    var v5 int32 = (v4 | 1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi41 := 0; fi41 < 5; fi41++ {
        switch fi41 {
        case 4294967295:
            for fi42 := 0; fi42 < 6; fi42++ {
                if (true && true) {
                    if false {
                        delete(v2, "foo bar")
                        var v43 float32 = -2.25
                        switch (fi42 / (((fi42 / ((fi41) | 1))) | 1)) {
                        case -773:
                            v43 += float32(3.14159)
                        default:
                            var v44 uint = 1
                            if true {
                                if (0.0 == 0.0) {
                                    v45 := GenS4{F0: 2, F1: 814}
                                    _ = v45
                                }
                            }
                            var v46 string = "foo bar"
                            _ = v44
                            _ = v46
                        }
                        var v47 int16 = 1000
                        delete(v2, "foo bar")
                        _ = v43
                        _ = v47
                    }
                    v48 := GenS7{F0: v1, F1: 1000}
                    var v49 uint = 121
                    _ = v48
                    _ = v49
                }
                if (fi42 <= (^(fi42))) {
                    if (523 <= 2147483647) {
                        v5 += (v4 + v4)
                        switch (fi41 % ((fi41) | 1)) {
                        case 0:
                            var v50 uint32 = 42
                            _ = v50
                        default:
                            delete(v2, ("hello" + "hello"))
                            if false {
                                var v51 float32 = 1e10
                                var v52 int16 = 0
                                _ = v51
                                _ = v52
                            }
                        }
                        fmt.Println(v5, v5, fi42)
                    } else {
                        v1 = (v1 - (v1 * v1))
                        v5 = -703
                    }
                    fmt.Println(v5, v5, fi41)
                }
                _ = fi42
            }
        case 9223372036854775807:
            v1 = -1
        }
        _ = fi41
    }
    if (!(((18446744073709551615 < 18446744073709551615) && (65535 >= 0)) && (!(-9223372036854775808 <= 0)))) {
        if true {
            switch "" {
            case "foo bar":
                v4 -= v5
            case "a":
                if (838 == 65535) {
                    var v53 int64 = -530
                    if false {
                        v5 = 2147483647
                    }
                    _ = v53
                }
            default:
                var v54 uint32 = 0
                _ = v54
            }
            fmt.Println(v4)
            v1 -= (v1 + v1)
        } else {
            if true {
                for fi55 := 0; fi55 < 1; fi55++ {
                    var v56 int8 = v1
                    for fi57 := 0; fi57 < 7; fi57++ {
                        v5 += (-(v5))
                        if (-1 >= 218) {
                            switch 9223372036854775807 {
                            case -1:
                                var v58 uint64 = 18446744073709551615
                                _ = v58
                            case 32767:
                                switch "\n\t" {
                                case "hello":
                                    switch (fi55 | 0) {
                                    case 0:
                                        var v59 uint = 912
                                        _ = v59
                                    case 127:
                                        var v60 string = ("foo bar" + "")
                                        _ = v60
                                    default:
                                        v61 := GenS7{F0: v1, F1: -36}
                                        _ = v61
                                    }
                                case "a":
                                    if (("foo bar" + "") < (((("hello" + "a") + ("" + ("\n\t" + "foo bar"))) + "foo bar") + ("" + "\n\t"))) {
                                        _ = len(v3)
                                        switch ("foo bar" + ("foo bar" + ("hello" + "\n\t"))) {
                                        case "a":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case "hello":
                                            v62 := GenS1{F0: 42}
                                            fmt.Println(v1, v56)
                                            _ = v62
                                        }
                                        var v63 int8 = (-(v56))
                                        v56 -= v56
                                        _ = v63
                                    } else {
                                        fmt.Println(v56, v56)
                                        delete(v2, ("hello" + "a"))
                                        if (!false) {
                                            fmt.Println(fi57, v56)
                                        } else {
                                            if true {
                                                fmt.Println(v4, v4, fi57)
                                                if false {
                                                    v64 := v2[("hello" + "")]
                                                    _ = v64
                                                }
                                                var v65 int32 = 1
                                                _ = v65
                                            } else {
                                                v5 = 0
                                                var v66 float32 = float32((3.14159 / 1.0))
                                                if len(v3) > 0 {
                                                    v3[0] = 0
                                                }
                                                v3 = append(v3, 1)
                                                _ = v66
                                            }
                                        }
                                    }
                                case "foo bar":
                                    v1 = 1
                                }
                            case -9223372036854775808:
                                v4 += (v5 - v4)
                            default:
                                v1 = 1
                            }
                            var v67 uint64 = 1
                            v67 = (v67 - 4294967295)
                            _ = v67
                        }
                        _ = fi57
                    }
                    v68 := GenS5{F0: v56, F1: true, F2: -365}
                    _ = fi55
                    _ = v56
                    _ = v68
                }
                var v69 float64 = (-2.25 / (1.5 * ((((float64(3.14159) / 3.14159) - 0.1) / 1e10) - ((float64(0.0) * 0.0) * 0.0))))
                if ((!false) && false) {
                    var v70 int = 360
                    var v71 int16 = 42
                    var v72 int64 = 752
                    _ = v70
                    _ = v71
                    _ = v72
                } else {
                    for fi73 := 0; fi73 < 9; fi73++ {
                        v69 -= (v69 + -2.25)
                        var v74 int8 = (v1 % ((-1) | 1))
                        _ = fi73
                        _ = v74
                    }
                    if (1 >= 0) {
                        _ = len(v3)
                        for fi75 := 0; fi75 < 4; fi75++ {
                            switch fi75 {
                            case 32767:
                                v76 := GenS7{F0: (v1 % ((v1) | 1)), F1: 1000}
                                _ = v76
                            default:
                                v4 = (v4 / ((634) | 1))
                            }
                            _ = fi75
                        }
                    } else {
                        v4 = v4
                        if true {
                            v4 -= (v5 >> ((65535) & 63))
                            fmt.Println(v69, v69, v4)
                        } else {
                            v77 := GenS8{F0: 2147483647}
                            _ = len(v3)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v77
                        }
                        v1 -= v1
                    }
                    var v78 int = 4294967295
                    _ = v78
                }
                _ = v69
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    }
    fmt.Println(v4, v1, v1)
    switch "a" {
    case "foo bar":
        var v79 bool = ((("a" + ("\n\t" + "a")) == "\n\t") || (0 < 704))
        _ = v79
    case "a":
        for fi80 := 0; fi80 < 17; fi80++ {
            var v81 int16 = 0
            _ = fi80
            _ = v81
        }
    case "\n\t":
        switch ("hello" + "\n\t") {
        case "\n\t":
            v82 := GenS1{F0: 0}
            _ = v82
        default:
            var v83 int32 = (v4 >> ((18446744073709551615) & 63))
            _ = v83
        }
    }
    v84 := GenS9{F0: 127}
    v85 := GenS7{F0: (-(v1)), F1: 1}
    v4 += v4
    v5 = (v4 & v5)
    v4 = v4
    var v86 uint = 373
    v4 = (v4 * v4)
    var v87 uint8 = 42
    v4 = (v4 & (v4 ^ 0))
    v88 := GenS10{F0: -1, F1: 882, F2: 42}
    _ = len(v3)
    for fi89 := 0; fi89 < 8; fi89++ {
        switch ("\n\t" + "hello") {
        case "":
            v90 := GenS11{F0: -9223372036854775808}
            _ = v90
        }
        _ = fi89
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v84
    _ = v85
    _ = v86
    _ = v87
    _ = v88
}
