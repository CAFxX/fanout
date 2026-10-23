package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 bool; F1 uint16; F2 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8; F1 int16; F2 uint32 }
type GenS10 struct { F0 int32; F1 bool; F2 int16 }
func genFn0(p0 bool) uint32 {
    switch 2147483647 {
    case 32767:
        var v5 int8 = 42
        _ = v5
    case 32768:
        for fi6 := 0; fi6 < 5; fi6++ {
            for fi7 := 0; fi7 < 18; fi7++ {
                p0 = (267 == 996)
                _ = fi7
            }
            _ = fi6
        }
    default:
        if p0 {
            for fi8 := 0; fi8 < 10; fi8++ {
                for fi9 := 0; fi9 < 15; fi9++ {
                    var v10 uint8 = 2
                    _ = fi9
                    _ = v10
                }
                var v11 bool = true
                _ = fi8
                _ = v11
            }
            var v12 uint16 = 255
            _ = v12
        }
    }
    var v13 string = "foo bar"
    for fi14 := 0; fi14 < 9; fi14++ {
        if (100 > 96) {
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := []int{(fi14 * (^(fi15))), 430, 0, 32768}
                if p0 {
                    if p0 {
                        fmt.Println(fi14)
                        p0 = (fi15 == (fi14 - fi15))
                    } else {
                        var v17 float32 = 1e10
                        var v18 string = ((v13 + "\n\t") + "")
                        v16 = append(v16, fi15)
                        fmt.Println(fi15, p0)
                        v19 := GenS1{F0: fi14, F1: -90}
                        var v20 float32 = 3.14159
                        _ = v17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    switch "" {
                    case "hello":
                        p0 = (1 < 18446744073709551615)
                    case "a":
                        var v21 int64 = int64(fi15)
                        _ = v21
                    default:
                        fmt.Println(v13, fi15, fi14)
                    }
                    v13 = v13
                    var v22 bool = (42 != -9223372036854775808)
                    if (826 > 918) {
                        var v23 uint16 = 457
                        var v24 uint = 473
                        _ = v23
                        _ = v24
                    } else {
                        if p0 {
                            fmt.Println(v13, v13)
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(fi14)
                            case "hello":
                                _ = len(v16)
                            }
                        } else {
                            switch "hello" {
                            case "foo bar":
                                switch 42 {
                                case 42:
                                    var v25 uint16 = 42
                                    _ = v25
                                }
                            default:
                                _ = len(v16)
                            }
                        }
                        v26 := GenS2{F0: 127, F1: 2}
                        _ = v26
                    }
                    v27 := GenS3{F0: 179}
                    if v22 {
                        p0 = v22
                    }
                    var v28 float32 = 1e10
                    _ = v22
                    _ = v27
                    _ = v28
                } else {
                    v13 = (v13 + (v13 + "a"))
                    p0 = (42 != 85)
                }
                var v29 int = fi14
                _ = fi15
                _ = v16
                _ = v29
            }
            if p0 {
                switch (fi14 * (fi14 | (fi14 / ((128) | 1)))) {
                case -9223372036854775808:
                    var v30 uint8 = 42
                    _ = v30
                case 128:
                    v31 := []int{fi14}
                    _ = v31
                default:
                    for fi32 := 0; fi32 < 3; fi32++ {
                        v33 := []int{(fi32 / ((fi14) | 1))}
                        switch (fi32 >> ((1) & 63)) {
                        case 127:
                            if len(v33) > 0 {
                                v33[0] = -182
                            }
                        case -2147483648:
                            var v34 string = "hello"
                            _ = v34
                        case 128:
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                        }
                        _ = fi32
                        _ = v33
                    }
                }
                if p0 {
                    p0 = (true && false)
                    var v35 bool = (255 <= 1)
                    _ = v35
                }
                for fi36 := 0; fi36 < 1; fi36++ {
                    if (!true) {
                        switch (^(fi14)) {
                        case -518:
                            var v37 uint64 = 86
                            _ = v37
                        default:
                            var v38 int32 = 42
                            _ = v38
                        }
                    }
                    p0 = true
                    v39 := []int{(fi14 + fi36)}
                    _ = fi36
                    _ = v39
                }
            }
            fmt.Println(fi14, fi14)
        }
        p0 = false
        _ = fi14
    }
    _ = p0
    _ = v13
    return 2147483647
}

func main() {
    var v1 bool = true
    var v2 uint = 4294967295
    var v3 int64 = 359
    var v4 float32 = (float32((1.5 / float64((1e10 + float32(0.1))))) * (-2.25 * 1.5))
    v40 := []int{-27, 2147483647, -43}
    v4 = 0.0
    if (254 >= 437) {
        if len(v40) > 0 {
            v40[0] = -359
        }
    } else {
        _ = len(v40)
        fmt.Println(v2)
    }
    for fi41 := 0; fi41 < 11; fi41++ {
        var v42 int8 = 1
        v43 := GenS2{F0: 127, F1: 65535}
        for fi44 := 0; fi44 < 5; fi44++ {
            if (!(0 > 2)) {
                v45 := GenS3{F0: 724}
                v46 := GenS4{F0: 637, F1: -1}
                if (!v1) {
                    switch int(v3) {
                    case 1:
                        if (!v1) {
                            v42 += v42
                            var v47 string = (("a" + "hello") + ("\n\t" + "foo bar"))
                            v4 = -2.25
                            _ = v47
                        }
                        switch "a" {
                        case "hello":
                            var v48 float64 = -2.25
                            if len(v40) > 0 {
                                v40[0] = fi44
                            }
                            var v49 float32 = float32(v48)
                            _ = v48
                            _ = v49
                        default:
                            v2 = v2
                        }
                    case 2147483647:
                        switch "a" {
                        case "":
                            v50 := GenS1{F0: 9223372036854775807, F1: v42}
                            _ = v50
                        case "a":
                            if (!((65535 == uint32(v2)) || (2 == 872))) {
                                v51 := GenS5{F0: 100, F1: 167}
                                switch "foo bar" {
                                case "a":
                                    if v1 {
                                        _ = len(v40)
                                        var v52 string = ("foo bar" + "foo bar")
                                        v4 = v4
                                        _ = v52
                                    } else {
                                        v53 := GenS4{F0: -1, F1: -1}
                                        var v54 int32 = 42
                                        _ = v53
                                        _ = v54
                                    }
                                case "\n\t":
                                    if (genFn0((!(v1 && ((637 != 0) && ((v1 && false) && false))))) == 2) {
                                        var v55 bool = ((2147483647 <= 42) && (!true))
                                        fmt.Println(fi41, v42, v3)
                                        v56 := GenS2{F0: 255, F1: 2147483648}
                                        _ = v55
                                        _ = v56
                                    } else {
                                        v1 = true
                                        var v57 string = "foo bar"
                                        if ((!(42 >= v3)) || false) {
                                            v58 := GenS3{F0: 18446744073709551615}
                                            switch (fi41 >> ((18446744073709551615) & 63)) {
                                            case 173:
                                                v57 = "foo bar"
                                            }
                                            if (!true) {
                                                var v59 int8 = v42
                                                var v60 int8 = v59
                                                if len(v40) > 0 {
                                                    v40[0] = 9223372036854775807
                                                }
                                                v61 := GenS3{F0: 388}
                                                _ = v59
                                                _ = v60
                                                _ = v61
                                            } else {
                                                switch fi41 {
                                                case 32767:
                                                    var v62 float32 = (float32(0.1) * 3.14159)
                                                    _ = v62
                                                case 0:
                                                    v3 = v3
                                                default:
                                                    v63 := GenS6{F0: (true || (float32(0.0) <= 0.0)), F1: 65535, F2: v1}
                                                    _ = v63
                                                }
                                                fmt.Println(fi41, fi44)
                                                var v64 uint32 = 2147483648
                                                _ = v64
                                            }
                                            _ = v58
                                        } else {
                                            v3 += v3
                                        }
                                        _ = v57
                                    }
                                case "foo bar":
                                    fmt.Println(v4, fi41, v1)
                                default:
                                    fmt.Println(v1)
                                }
                                _ = v51
                            }
                        case "hello":
                            if (!(!false)) {
                                _ = len(v40)
                                switch "a" {
                                case "":
                                    fmt.Println(v3, v3)
                                case "foo bar":
                                    var v65 int32 = -2147483648
                                    _ = v65
                                default:
                                    if ("foo bar" < ("hello" + "foo bar")) {
                                        v42 = (v42 / (((v42 >> (((v2 >> (((v2 & (v2 / (((v2 % (((v2 + 0)) | 1))) | 1)))) & 63))) & 63))) | 1))
                                    } else {
                                        var v66 uint32 = 2147483647
                                        _ = v66
                                    }
                                }
                                v67 := GenS2{F0: 42, F1: 1}
                                _ = v67
                            } else {
                                v40 = append(v40, (fi44 << ((0) & 63)))
                                if (0 < 282) {
                                    fmt.Println(v3, v42)
                                    switch (fi44 | fi44) {
                                    case 4294967295:
                                        switch "hello" {
                                        case "a":
                                            v2 = 801
                                        case "foo bar":
                                            v68 := GenS2{F0: 186, F1: 2}
                                            _ = v68
                                        case "":
                                            v4 = -2.25
                                        default:
                                            fmt.Println(v2)
                                        }
                                    }
                                } else {
                                    v69 := GenS7{F0: 2}
                                    _ = v69
                                }
                                if len(v40) >= 2 {
                                    v40 = v40[1:len(v40)]
                                }
                            }
                        }
                    case 42:
                        var v70 float32 = (v4 - v4)
                        _ = v70
                    default:
                        v1 = ((v3 | (v3 % (((v3 << ((v2) & 63))) | 1))) > (v3 >> ((v2) & 63)))
                    }
                }
                _ = v45
                _ = v46
            }
            v71 := GenS5{F0: (v42 - v42), F1: 0}
            var v72 uint32 = 42
            var v73 int32 = 42
            var v74 uint64 = 54
            var v75 int = fi41
            var v76 string = "hello"
            var v77 int8 = v42
            v78 := GenS8{F0: "a"}
            _ = fi44
            _ = v71
            _ = v72
            _ = v73
            _ = v74
            _ = v75
            _ = v76
            _ = v77
            _ = v78
        }
        _ = fi41
        _ = v42
        _ = v43
    }
    if v1 {
        for fi79 := 0; fi79 < 15; fi79++ {
            if len(v40) >= 2 {
                v40 = v40[1:len(v40)]
            }
            _ = fi79
        }
        var v80 int = 871
        _ = v80
    } else {
        var v81 uint8 = 2
        for fi82 := 0; fi82 < 10; fi82++ {
            v4 = v4
            var v83 float64 = 0.1
            v84 := GenS1{F0: 32767, F1: -1}
            _ = fi82
            _ = v83
            _ = v84
        }
        var v85 uint64 = 736
        _ = v81
        _ = v85
    }
    if (((float32(1e10) * float32(0.0)) >= v4) && ("foo bar" > "hello")) {
        if v1 {
            var v86 string = "a"
            if len(v40) >= 2 {
                v40 = v40[1:len(v40)]
            }
            _ = v86
        } else {
            fmt.Println(v1, v1)
            for fi87 := 0; fi87 < 18; fi87++ {
                if len(v40) > 0 {
                    v40[0] = fi87
                }
                _ = fi87
            }
        }
        switch -192 {
        case 42:
            v88 := GenS7{F0: 2}
            _ = v88
        case 614:
            if v1 {
                switch "a" {
                case "foo bar":
                    for fi89 := 0; fi89 < 8; fi89++ {
                        v4 = 3.14159
                        switch fi89 {
                        case 0:
                            var v90 uint16 = 403
                            _ = v90
                        default:
                            v91 := GenS1{F0: (fi89 + fi89), F1: -128}
                            _ = v91
                        }
                        _ = len(v40)
                        _ = fi89
                    }
                default:
                    var v92 int8 = 1
                    _ = v92
                }
                v93 := GenS2{F0: 42, F1: genFn0(false)}
                _ = v93
            } else {
                for fi94 := 0; fi94 < 1; fi94++ {
                    v95 := GenS9{F0: -128, F1: int16(fi94), F2: 1}
                    fmt.Println(v1)
                    _ = fi94
                    _ = v95
                }
            }
        case 1:
            var v96 int32 = 380
            _ = v96
        }
    } else {
        var v97 int8 = 127
        _ = v97
    }
    v98 := GenS10{F0: -459, F1: v1, F2: 877}
    switch 369 {
    case 0:
        v2 = v2
    default:
        switch ((("foo bar" + "\n\t") + "\n\t") + "foo bar") {
        case "a":
            v40 = append(v40, 9223372036854775807)
        case "hello":
            switch "a" {
            case "a":
                v4 = 3.14159
            case "":
                var v99 int8 = -112
                _ = v99
            default:
                fmt.Println(v2, v3)
            }
        case "foo bar":
            var v100 int32 = 42
            _ = v100
        }
    }
    fmt.Println(v4, v1)
    var v101 int64 = (v3 % ((0) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v98
}
