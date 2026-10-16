package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint8; F1 string; F2 int16 }
type GenS8 struct { F0 int64; F1 uint32; F2 uint }
type GenS9 struct { F0 int16; F1 uint; F2 bool }
type GenS10 struct { F0 int64; F1 int }

func main() {
    var v1 float32 = (((-2.25 / 1.0) / 1.5) / ((1e10 + -2.25) * ((3.14159 / 1.0) / 1e10)))
    v2 := []int{-2147483648, 1}
    v3 := []uint32{2}
    var v4 int8 = 40
    var v5 int16 = -695
    var v6 uint8 = 127
    v7 := GenS1{F0: "\n\t", F1: 0}
    var v8 uint8 = 146
    v6 = 255
    var v9 uint16 = 102
    switch 42 {
    case 127:
        switch 851 {
        case 128:
            if (127 == -2147483648) {
                var v10 uint8 = v6
                for fi11 := 0; fi11 < 16; fi11++ {
                    v12 := GenS1{F0: (("a" + ("hello" + "")) + "a"), F1: 1}
                    _ = fi11
                    _ = v12
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v10
            }
        case 32767:
            v4 = 42
        case 9223372036854775807:
            var v13 int32 = 1
            _ = v13
        default:
            var v14 uint64 = 2
            _ = v14
        }
    }
    var v15 uint16 = 255
    if (v1 > v1) {
        v15 = v15
        switch ("" + "a") {
        case "hello":
            v1 -= 0.0
        case "a":
            fmt.Println(v5, v15, v6)
        case "":
            var v16 float64 = -2.25
            _ = v16
        default:
            var v17 uint32 = 2147483648
            _ = v17
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            switch "hello" {
            case "":
                v19 := GenS2{F0: v4, F1: 18446744073709551615, F2: 42}
                _ = v19
            default:
                if true {
                    if ((true || ((766 == (fi18 & (fi18 ^ fi18))) || (!true))) || true) {
                        var v20 uint16 = v15
                        if (904 == 4294967295) {
                            v1 = (v1 + v1)
                            switch fi18 {
                            case 127:
                                var v21 uint64 = 251
                                _ = v21
                            case 32768:
                                if true {
                                    for fi22 := 0; fi22 < 1; fi22++ {
                                        v9 += 712
                                        _ = fi22
                                    }
                                    v3 = append(v3, 4294967295)
                                    for fi23 := 0; fi23 < 19; fi23++ {
                                        fmt.Println(fi18, fi18, v20)
                                        v15 += (v15 | v9)
                                        var v24 uint32 = 977
                                        _ = fi23
                                        _ = v24
                                    }
                                } else {
                                    _ = len(v2)
                                    var v25 uint8 = v6
                                    _ = v25
                                }
                            }
                        }
                        _ = v20
                    }
                    if (42 == 0) {
                        switch ("" + "a") {
                        case "foo bar":
                            v8 -= (v6 + v8)
                        default:
                            var v26 int = int(v4)
                            _ = v26
                        }
                        var v27 bool = false
                        _ = v27
                    } else {
                        switch fi18 {
                        case 32767:
                            if ((fi18 ^ 32768) != (fi18 - (fi18 >> ((493) & 63)))) {
                                var v28 uint64 = 4294967295
                                if (315 <= fi18) {
                                    for fi29 := 0; fi29 < 6; fi29++ {
                                        if false {
                                            var v30 uint64 = v28
                                            if true {
                                                fmt.Println(v4, v6, fi18)
                                            } else {
                                                if len(v2) > 0 {
                                                    v2[0] = fi18
                                                }
                                                var v31 float64 = 1.5
                                                switch ("foo bar" + "\n\t") {
                                                case "":
                                                    if false {
                                                        v5 = v5
                                                        var v32 int8 = v4
                                                        _ = v32
                                                    } else {
                                                        fmt.Println(v30)
                                                    }
                                                }
                                                _ = v31
                                            }
                                            if true {
                                                v33 := GenS3{F0: 2147483648, F1: 42}
                                                fmt.Println(v15, v5)
                                                v3 = append(v3, 2)
                                                _ = v33
                                            } else {
                                                v34 := GenS4{F0: v15, F1: 4294967295}
                                                v28 += v28
                                                v6 = v6
                                                _ = v34
                                            }
                                            _ = v30
                                        } else {
                                            v35 := GenS5{F0: (^(v4)), F1: (fi29 - fi29)}
                                            v15 -= 65535
                                            if ((v28 / (((v28 & v28)) | 1)) != v28) {
                                                var v36 uint64 = 2
                                                v9 -= (^(v9))
                                                v37 := GenS6{F0: 681}
                                                v1 += v1
                                                fmt.Println(v8, v15, fi18)
                                                _ = v36
                                                _ = v37
                                            }
                                            _ = v35
                                        }
                                        var v38 int8 = (v4 ^ -72)
                                        _ = fi29
                                        _ = v38
                                    }
                                    var v39 float64 = (float64((v1 - 3.14159)) + (((float64(3.14159) - (0.1 / (1e10 - float64(3.14159)))) + float64((v1 + (float32(-2.25) + (float32(((0.1 * 0.0) / 3.14159)) / v1))))) * 3.14159))
                                    v5 -= (v5 / ((v5) | 1))
                                    _ = v39
                                } else {
                                    var v40 uint32 = 773
                                    v15 = (v15 * v15)
                                    v1 -= (float32(1.5) / 1.0)
                                    _ = v40
                                }
                                _ = v28
                            } else {
                                switch fi18 {
                                case 657:
                                    var v41 uint32 = 65535
                                    _ = v41
                                }
                                switch (fi18 % ((-2147483648) | 1)) {
                                case 127:
                                    switch ("a" + "foo bar") {
                                    case "":
                                        var v42 float64 = (float64((1.5 + v1)) / 3.14159)
                                        _ = v42
                                    }
                                case 9223372036854775807:
                                    v43 := GenS7{F0: 68, F1: ((("hello" + "\n\t") + "foo bar") + "a"), F2: 0}
                                    _ = v43
                                case -809:
                                    if len(v2) > 0 {
                                        v2[0] = fi18
                                    }
                                default:
                                    fmt.Println(fi18, v9)
                                }
                                if false {
                                    if (v1 >= 0.0) {
                                        for fi44 := 0; fi44 < 6; fi44++ {
                                            var v45 uint = 2
                                            _ = fi44
                                            _ = v45
                                        }
                                        var v46 float32 = 1e10
                                        _ = v46
                                    }
                                    fmt.Println(v8, v15, v5)
                                }
                            }
                        case -464:
                            var v47 int32 = 924
                            _ = v47
                        case -1:
                            v3 = append(v3, 4294967295)
                        default:
                            for fi48 := 0; fi48 < 4; fi48++ {
                                _ = len(v2)
                                if (1 < 4294967295) {
                                    v49 := GenS8{F0: 9223372036854775807, F1: 507, F2: 596}
                                    _ = v49
                                } else {
                                    _ = len(v3)
                                }
                                _ = fi48
                            }
                        }
                        fmt.Println(v1, v6)
                        fmt.Println(v6, v8, v6)
                    }
                    var v50 uint = 4294967295
                    _ = v50
                } else {
                    var v51 uint64 = 0
                    switch (fi18 + fi18) {
                    case 1:
                        v52 := GenS9{F0: (-(v5)), F1: 65535, F2: (v4 == (v4 - (v4 & v4)))}
                        _ = v52
                    case 0:
                        var v53 uint8 = (v6 % (((v6 | v8)) | 1))
                        _ = v53
                    case -1:
                        var v54 int64 = -1
                        _ = v54
                    }
                    _ = v51
                }
            }
            _ = fi18
        }
    } else {
        fmt.Println(v9, v1)
    }
    v55 := GenS2{F0: (v4 / (((v4 - (-(v4)))) | 1)), F1: 2, F2: 4294967295}
    var v56 int16 = (^(v5))
    v57 := GenS5{F0: v4, F1: 128}
    switch "" {
    case "hello":
        if false {
            switch "foo bar" {
            case "hello":
                var v58 int = 128
                _ = v58
            case "\n\t":
                var v59 bool = (3.14159 == float64(((-2.25 / 1e10) / v1)))
                _ = v59
            default:
                fmt.Println(v8, v8)
            }
            var v60 uint16 = 452
            _ = v60
        }
    }
    if false {
        if len(v2) > 0 {
            v2[0] = -728
        }
    } else {
        v56 += v56
        _ = len(v2)
        switch ("foo bar" + "foo bar") {
        case "a":
            for fi61 := 0; fi61 < 2; fi61++ {
                switch "a" {
                case "a":
                    v4 = v4
                case "\n\t":
                    if true {
                        var v62 uint16 = (^(v9))
                        if false {
                            var v63 int32 = 2147483647
                            v4 = (v4 * (v4 - (^(v4))))
                            v1 = 3.14159
                            _ = v63
                        }
                        _ = v62
                    }
                case "hello":
                    var v64 int16 = -120
                    _ = v64
                default:
                    if false {
                        var v65 float64 = ((-2.25 - -2.25) + ((float64((v1 / (1e10 + float32(0.1)))) * (-2.25 - (0.0 * -2.25))) - 3.14159))
                        if (true && false) {
                            for fi66 := 0; fi66 < 4; fi66++ {
                                v65 = -2.25
                                if len(v3) > 0 {
                                    v3[0] = 2147483648
                                }
                                v67 := GenS8{F0: 1, F1: 621, F2: 798}
                                _ = fi66
                                _ = v67
                            }
                        } else {
                            var v68 float32 = v1
                            _ = v68
                        }
                        _ = v65
                    }
                }
                var v69 int16 = v56
                _ = fi61
                _ = v69
            }
        case "\n\t":
            v70 := GenS10{F0: 1, F1: 2147483647}
            _ = v70
        case "":
            var v71 int8 = int8(v5)
            _ = v71
        default:
            v1 += v1
        }
    }
    switch 1 {
    case 128:
        fmt.Println(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
