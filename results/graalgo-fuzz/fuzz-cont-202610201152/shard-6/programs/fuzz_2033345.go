package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32; F1 bool; F2 int16 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int; F1 int16; F2 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 int16; F2 uint32 }
type GenS7 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS10 struct { F0 uint8; F1 uint; F2 bool }
func genFn0(p0 int) int16 {
    v4 := GenS1{F0: 4294967295, F1: 1}
    if (!(!false)) {
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 uint8 = 128
            p0 -= fi5
            v6 -= 224
            _ = fi5
            _ = v6
        }
    }
    switch "\n\t" {
    case "foo bar":
        if (956 != 65535) {
            if false {
                switch "a" {
                case "\n\t":
                    fmt.Println(p0)
                case "a":
                    if (2 == 217) {
                        v7 := []int{}
                        var v8 float64 = -2.25
                        _ = v7
                        _ = v8
                    } else {
                        var v9 int32 = 168
                        _ = v9
                    }
                case "hello":
                    fmt.Println(p0)
                default:
                    p0 += 49
                }
                switch "hello" {
                case "":
                    for fi10 := 0; fi10 < 6; fi10++ {
                        v11 := []int{p0, (fi10 >> ((18446744073709551615) & 63)), p0, 30}
                        v12 := GenS2{F0: 42, F1: false, F2: 42}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                case "foo bar":
                    if false {
                        v13 := GenS3{F0: 656, F1: 281, F2: 0}
                        fmt.Println(p0)
                        _ = v13
                    }
                }
                p0 = p0
            }
            if (!true) {
                fmt.Println(p0)
                switch (p0 + (p0 >> ((2) & 63))) {
                case 9223372036854775807:
                    for fi14 := 0; fi14 < 5; fi14++ {
                        fmt.Println(p0, p0)
                        _ = fi14
                    }
                }
                for fi15 := 0; fi15 < 10; fi15++ {
                    v16 := GenS3{F0: 32767, F1: 42, F2: 555}
                    _ = fi15
                    _ = v16
                }
            } else {
                for fi17 := 0; fi17 < 12; fi17++ {
                    v18 := GenS4{F0: 0, F1: -1, F2: 830}
                    _ = fi17
                    _ = v18
                }
                var v19 bool = true
                if (!v19) {
                    if (1 >= 9223372036854775807) {
                        fmt.Println(p0)
                        for fi20 := 0; fi20 < 4; fi20++ {
                            v19 = true
                            var v21 uint64 = 549
                            v21 = v21
                            _ = fi20
                            _ = v21
                        }
                    } else {
                        v22 := GenS1{F0: 2, F1: 255}
                        for fi23 := 0; fi23 < 4; fi23++ {
                            v24 := GenS5{F0: ("\n\t" + "\n\t")}
                            var v25 uint8 = 1
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        _ = v22
                    }
                } else {
                    p0 = p0
                    for fi26 := 0; fi26 < 1; fi26++ {
                        for fi27 := 0; fi27 < 12; fi27++ {
                            p0 = p0
                            var v28 float32 = 3.14159
                            if v19 {
                                v29 := GenS2{F0: 0, F1: (((2147483647 < 329) && v19) && v19), F2: 1000}
                                _ = v29
                            } else {
                                v28 -= 0.0
                                var v30 uint8 = 52
                                fmt.Println(v30, v30, p0)
                                v31 := GenS6{F0: 1, F1: -32768, F2: 4294967295}
                                v32 := GenS3{F0: -509, F1: 697, F2: 2}
                                v28 -= v28
                                v33 := GenS2{F0: 0, F1: false, F2: 42}
                                v28 = ((v28 * (float32((1e10 - 0.1)) * ((-2.25 * float32(((1.5 * ((-2.25 + (((1.5 + 3.14159) - float64(0.0)) - (1.5 + 0.1))) * (3.14159 - 0.0))) * float64(0.0)))) - (v28 * 1e10)))) + (-2.25 + 0.1))
                                _ = v30
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            _ = fi27
                            _ = v28
                        }
                        _ = fi26
                    }
                }
                _ = v19
            }
        } else {
            var v34 int = (p0 & p0)
            _ = v34
        }
    case "hello":
        p0 = 127
    case "a":
        p0 -= (-(p0))
    }
    _ = p0
    _ = v4
    return 867
}

func main() {
    var v1 float32 = ((1e10 - (0.1 + 0.1)) * (-2.25 * -2.25))
    var v2 int32 = -761
    var v3 int8 = 100
    v1 += (3.14159 - -2.25)
    v1 = v1
    for fi35 := 0; fi35 < 1; fi35++ {
        if ((false || false) && (v1 != ((float32(3.14159) - v1) / (0.0 + v1)))) {
            v2 = v2
            v3 += 1
        } else {
            var v36 int16 = -32768
            v36 = int16(v2)
            _ = v36
        }
        var v37 int16 = 32767
        _ = fi35
        _ = v37
    }
    if (277 < 4294967295) {
        if ((127 > 42) && true) {
            var v38 int32 = v2
            var v39 float32 = v1
            switch "\n\t" {
            case "":
                v40 := GenS2{F0: v38, F1: false, F2: 1}
                _ = v40
            case "a":
                v1 += v1
            case "foo bar":
                v41 := []int{}
                _ = v41
            }
            _ = v38
            _ = v39
        } else {
            v1 -= v1
        }
        switch "foo bar" {
        case "hello":
            v2 -= (v2 - 126)
        case "":
            var v42 uint64 = 493
            _ = v42
        }
        var v43 uint64 = 4294967295
        _ = v43
    }
    v44 := []int{-2147483648}
    switch "\n\t" {
    case "":
        v45 := GenS3{F0: 0, F1: 1, F2: 435}
        _ = v45
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        var v47 uint32 = 815
        if (255 >= 0) {
            var v48 string = "foo bar"
            _ = v48
        } else {
            for fi49 := 0; fi49 < 19; fi49++ {
                if len(v44) >= 2 {
                    v44 = v44[1:len(v44)]
                }
                _ = fi49
            }
            fmt.Println(v3, fi46, v47)
            var v50 float32 = (v1 - 1e10)
            _ = v50
        }
        _ = fi46
        _ = v47
    }
    var v51 uint64 = 0
    if (true || true) {
        for fi52 := 0; fi52 < 14; fi52++ {
            if true {
                if (false && (!(4294967295 < 0))) {
                    if len(v44) >= 2 {
                        v44 = v44[1:len(v44)]
                    }
                    for fi53 := 0; fi53 < 15; fi53++ {
                        var v54 float64 = (0.1 + (3.14159 + (1e10 + 0.1)))
                        if len(v44) > 0 {
                            v44[0] = (fi53 & 9223372036854775807)
                        }
                        if false {
                            if ((2 == 826) || ((v54 > 3.14159) && false)) {
                                v55 := GenS3{F0: -139, F1: 210, F2: (v51 / ((2) | 1))}
                                v54 -= (v54 + (-2.25 + float64((v1 / float32(-2.25)))))
                                _ = v55
                            }
                            fmt.Println(v3)
                        } else {
                            _ = len(v44)
                        }
                        _ = fi53
                        _ = v54
                    }
                } else {
                    v56 := GenS5{F0: (("hello" + "") + ((("" + "foo bar") + ("\n\t" + "hello")) + "\n\t"))}
                    if len(v44) >= 2 {
                        v44 = v44[1:len(v44)]
                    }
                    _ = v56
                }
                fmt.Println(v2)
            } else {
                v57 := GenS1{F0: 65535, F1: 255}
                _ = v57
            }
            _ = fi52
        }
        for fi58 := 0; fi58 < 20; fi58++ {
            switch (-(fi58)) {
            case 42:
                _ = len(v44)
            case -478:
                fmt.Println(v2, v51, v3)
            }
            if len(v44) > 0 {
                v44[0] = 561
            }
            _ = fi58
        }
    }
    v51 -= v51
    if len(v44) > 0 {
        v44[0] = 1
    }
    v3 = (v3 + (v3 % ((v3) | 1)))
    for fi59 := 0; fi59 < 7; fi59++ {
        if false {
            v1 = v1
            v60 := GenS2{F0: v2, F1: ((342 >= 706) && true), F2: int16(v3)}
            _ = v60
        } else {
            v51 = v51
            v2 -= v2
            for fi61 := 0; fi61 < 15; fi61++ {
                v51 += v51
                switch fi61 {
                case 127:
                    var v62 int = 84
                    _ = v62
                case 1:
                    if false {
                        if true {
                            if (false || (187 > 139)) {
                                var v63 int64 = 1
                                _ = v63
                            } else {
                                switch (("hello" + "a") + ("a" + ("" + (("a" + "\n\t") + "")))) {
                                case "\n\t":
                                    if (v3 > v3) {
                                        if len(v44) >= 2 {
                                            v44 = v44[1:len(v44)]
                                        }
                                        var v64 uint16 = 2
                                        _ = v64
                                    } else {
                                        fmt.Println(fi59, v3)
                                    }
                                case "a":
                                    if (255 < 123) {
                                        v51 -= v51
                                    } else {
                                        v65 := GenS6{F0: v3, F1: -709, F2: 596}
                                        v1 = v1
                                        v1 = v1
                                        _ = v65
                                    }
                                case "foo bar":
                                    v2 = (v2 - (v2 & (v2 ^ (v2 ^ v2))))
                                }
                                v51 += 248
                                switch (fi59 / (((fi59 ^ fi61)) | 1)) {
                                case 4294967295:
                                    v1 -= 1.5
                                default:
                                    v2 = (v2 % (((v2 + v2)) | 1))
                                }
                            }
                            v66 := GenS2{F0: (-(v2)), F1: false, F2: 781}
                            _ = v66
                        } else {
                            v67 := GenS4{F0: fi61, F1: 0, F2: 804}
                            _ = len(v44)
                            var v68 uint64 = v51
                            _ = v67
                            _ = v68
                        }
                        if true {
                            var v69 uint8 = 128
                            _ = v69
                        }
                        var v70 uint8 = 42
                        _ = v70
                    } else {
                        switch "a" {
                        case "a":
                            v71 := GenS7{F0: 4294967295, F1: 527, F2: (^(v2))}
                            _ = v71
                        case "":
                            fmt.Println(v2, v51, v3)
                        }
                    }
                }
                switch (fi61 >> ((804) & 63)) {
                case 0:
                    switch (-(fi59)) {
                    case 32768:
                        v72 := GenS6{F0: (v3 | v3), F1: int16(v51), F2: 758}
                        fmt.Println(v51)
                        v73 := GenS2{F0: 0, F1: (2 != 425), F2: 1000}
                        _ = v72
                        _ = v73
                    case -691:
                        var v74 float64 = 1e10
                        switch "foo bar" {
                        case "\n\t":
                            v51 += 366
                        case "a":
                            var v75 string = "hello"
                            _ = v75
                        case "foo bar":
                            v3 = -124
                        }
                        var v76 uint8 = 42
                        _ = v74
                        _ = v76
                    }
                default:
                    v1 = v1
                }
                _ = fi61
            }
        }
        var v77 float64 = 3.14159
        v2 = -1
        _ = fi59
        _ = v77
    }
    if true {
        v1 = (v1 * 0.1)
    } else {
        var v78 float32 = (v1 * v1)
        v79 := GenS8{F0: 205}
        if true {
            for fi80 := 0; fi80 < 10; fi80++ {
                for fi81 := 0; fi81 < 10; fi81++ {
                    fmt.Println(v78)
                    v82 := GenS1{F0: 42, F1: 836}
                    if (!((-32768 < -1) && true)) {
                        _ = len(v44)
                        v83 := GenS9{F0: 255, F1: 255, F2: (^(v3))}
                        v84 := GenS10{F0: uint8(v51), F1: 255, F2: ((v2 != (v2 ^ (v2 % ((93) | 1)))) || false)}
                        _ = v83
                        _ = v84
                    }
                    _ = fi81
                    _ = v82
                }
                _ = fi80
            }
            v2 = (v2 ^ -1)
            var v85 uint16 = 899
            _ = v85
        }
        _ = v78
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v51
}
