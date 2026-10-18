package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int; F1 bool; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS6 struct { F0 bool; F1 string; F2 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 string; F1 uint16 }
type GenS11 struct { F0 string; F1 int16; F2 uint32 }
type GenS12 struct { F0 string; F1 int16; F2 int }
type GenS13 struct { F0 int8; F1 uint32; F2 uint64 }
func genFn0(p0 int32, p1 int64, p2 uint64) uint {
    p2 -= 18446744073709551615
    fmt.Println(p0, p2, p1)
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 13; fi5++ {
            var v6 float64 = ((0.0 * 1.5) + (0.1 / ((-2.25 - 1.5) - -2.25)))
            var v7 int8 = -128
            v7 = v7
            _ = fi5
            _ = v6
            _ = v7
        }
        p1 = -596
        fmt.Println(fi4)
        _ = fi4
    }
    if (70 >= (^(p2))) {
        fmt.Println(p0, p0)
    }
    if true {
        v8 := GenS1{F0: 116, F1: 65535}
        if (true || (65535 != 42)) {
            p0 = (p0 % (((p0 << ((18446744073709551615) & 63))) | 1))
            var v9 uint32 = 42
            v9 = v9
            _ = v9
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 uint, p1 uint8) bool {
    p0 += p0
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := []int{fi10, int(p0), 32767}
        for fi12 := 0; fi12 < 20; fi12++ {
            if true {
                var v13 string = "hello"
                if true {
                    var v14 uint16 = 42
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = v14
                }
                _ = v13
            } else {
                if ((false && (4294967295 > 422)) && false) {
                    if true {
                        v15 := GenS2{F0: (fi10 ^ (fi10 * fi10)), F1: (!true)}
                        _ = v15
                    } else {
                        fmt.Println(fi10)
                        switch -954 {
                        case 9223372036854775807:
                            var v16 uint64 = 457
                            _ = v16
                        case 32768:
                            var v17 bool = (-450 <= 42)
                            _ = v17
                        }
                        var v18 int32 = 2147483647
                        v19 := GenS1{F0: 42, F1: 118}
                        p1 = (^(p1))
                        _ = v18
                        _ = v19
                    }
                    if len(v11) > 0 {
                        v11[0] = int(p0)
                    }
                    p0 = (p0 >> (((p0 + p0)) & 63))
                } else {
                    fmt.Println(p0)
                    p0 -= p0
                    if false {
                        var v20 uint32 = 387
                        v21 := GenS3{F0: (fi12 ^ fi10), F1: (-598 <= -1), F2: (^(p1))}
                        v11 = append(v11, (fi10 << ((p0) & 63)))
                        _ = v20
                        _ = v21
                    }
                }
                _ = len(v11)
                if true {
                    if (!false) {
                        p0 = (^(p0))
                        p1 = (p1 & p1)
                        v22 := GenS4{F0: 109}
                        _ = v22
                    } else {
                        if (-2147483648 <= 42) {
                            _ = len(v11)
                            var v23 int16 = 42
                            v24 := GenS5{F0: (v23 >> (((p0 >> ((p0) & 63))) & 63)), F1: 65535, F2: (p1 * (p1 | 255))}
                            v25 := GenS6{F0: true, F1: "\n\t", F2: 0}
                            p1 -= (p1 % ((p1) | 1))
                            fmt.Println(fi10, p0)
                            v11 = append(v11, fi12)
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        var v26 uint32 = 643
                        _ = v26
                    }
                }
            }
            fmt.Println(p1)
            p0 = p0
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    var v27 float64 = ((((float64((-2.25 / -2.25)) / ((0.0 + ((1.5 - 0.1) - 0.0)) * float64(1.5))) / (1e10 / 1e10)) * 0.1) / 1.5)
    _ = p0
    _ = p1
    _ = v27
    return false
}

func main() {
    v1 := []string{"", "hello", "a", ("hello" + "\n\t")}
    var v2 int64 = 434
    v3 := []string{"hello", "", (("hello" + "foo bar") + "hello")}
    v28 := GenS7{F0: ("\n\t" + "\n\t")}
    if false {
        v2 = (v2 ^ v2)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v29 := GenS5{F0: -32768, F1: 309, F2: 128}
    switch 96 {
    case 127:
        v2 = v2
    }
    v2 -= v2
    fmt.Println(v2)
    for fi30 := 0; fi30 < 16; fi30++ {
        _ = len(v3)
        if (!((0.1 + (0.0 - (0.1 + 0.0))) <= float32(float64(float32(0.1))))) {
            switch fi30 {
            case 2147483647:
                v31 := GenS8{F0: -1, F1: 255}
                _ = v31
            case -607:
                switch -9223372036854775808 {
                case 1:
                    v32 := GenS4{F0: 0}
                    _ = v32
                case 625:
                    v33 := GenS9{F0: 660, F1: fi30}
                    _ = v33
                default:
                    if (("foo bar" + ("a" + "")) >= "") {
                        var v34 string = "foo bar"
                        _ = v34
                    } else {
                        if (((float32(-2.25) + (1.5 - 1.5)) - (0.0 + (0.1 / (1e10 / 3.14159)))) == 0.1) {
                            if false {
                                fmt.Println(fi30, v2)
                            } else {
                                if (4294967295 >= 42) {
                                    if true {
                                        v2 = (v2 + (v2 ^ (-(v2))))
                                        _ = len(v3)
                                        for fi35 := 0; fi35 < 17; fi35++ {
                                            switch "foo bar" {
                                            case "a":
                                                if true {
                                                    switch (fi30 % ((fi35) | 1)) {
                                                    case 127:
                                                        fmt.Println(v2, fi35)
                                                    case 348:
                                                        if len(v1) > 0 {
                                                            v1[0] = (("" + ("\n\t" + "hello")) + "foo bar")
                                                        }
                                                    }
                                                    var v36 uint16 = 771
                                                    switch ("foo bar" + ("foo bar" + "a")) {
                                                    case "\n\t":
                                                        switch "" {
                                                        case "a":
                                                            var v37 bool = (861 < 4294967295)
                                                            _ = v37
                                                        case "foo bar":
                                                            if (("\n\t" + ("hello" + "foo bar")) >= "a") {
                                                                fmt.Println(fi35)
                                                                fmt.Println(v36, v2, v2)
                                                            } else {
                                                                switch "a" {
                                                                case "a":
                                                                    v36 -= (v36 % (((v36 | v36)) | 1))
                                                                case "\n\t":
                                                                    if false {
                                                                        switch (^(fi35)) {
                                                                        case 1:
                                                                            var v38 int16 = -517
                                                                            _ = v38
                                                                        }
                                                                        _ = len(v1)
                                                                    }
                                                                }
                                                                v39 := GenS5{F0: -389, F1: 2147483647, F2: 0}
                                                                _ = v39
                                                            }
                                                        }
                                                    case "foo bar":
                                                        var v40 float64 = float64(0.0)
                                                        _ = v40
                                                    case "hello":
                                                        v41 := GenS10{F0: "", F1: (v36 & (v36 << ((0) & 63)))}
                                                        _ = v41
                                                    }
                                                    _ = v36
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    switch (fi30 ^ (fi30 << ((65535) & 63))) {
                                                    case 32767:
                                                        v42 := GenS11{F0: "foo bar", F1: -960, F2: 0}
                                                        _ = v42
                                                    case -701:
                                                        var v43 uint32 = 1
                                                        _ = v43
                                                    case -1:
                                                        fmt.Println(fi35, v2)
                                                    }
                                                }
                                            case "":
                                                if true {
                                                    if ((-1 > 127) && true) {
                                                        v2 -= v2
                                                        v2 -= (v2 * (^(v2)))
                                                    } else {
                                                        if true {
                                                            fmt.Println(fi30, v2)
                                                            var v44 uint = 1
                                                            _ = v44
                                                        }
                                                    }
                                                }
                                            case "foo bar":
                                                switch "hello" {
                                                case "hello":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case "\n\t":
                                                    var v45 string = (((("" + ("" + "hello")) + "hello") + "") + "hello")
                                                    _ = v45
                                                case "a":
                                                    var v46 int64 = v2
                                                    v47 := GenS3{F0: 32767, F1: (((0.1 + (3.14159 + 0.1)) + float64(1.5)) <= float64(float32(3.14159))), F2: 127}
                                                    _ = v46
                                                    _ = v47
                                                default:
                                                    v2 += -9223372036854775808
                                                }
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = ""
                                            }
                                            fmt.Println(fi35, fi35)
                                            fmt.Println(v2, fi35, fi30)
                                            if (974 < 339) {
                                                var v48 uint32 = 42
                                                v48 = (v48 >> ((genFn0(-381, (v2 + -9223372036854775808), 1)) & 63))
                                                _ = v48
                                            }
                                            _ = fi35
                                        }
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        switch "hello" {
                                        case "":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case "foo bar":
                                            var v49 float64 = (1.5 - 0.1)
                                            _ = v49
                                        case "a":
                                            v50 := GenS12{F0: ("" + "foo bar"), F1: -1, F2: 4294967295}
                                            _ = v50
                                        }
                                    }
                                    if true {
                                        var v51 string = ("a" + ("foo bar" + "hello"))
                                        switch (fi30 & fi30) {
                                        case -897:
                                            var v52 float32 = ((((float32((0.0 / 1.5)) / 1.0) + 0.0) + 1e10) + 1e10)
                                            _ = v52
                                        case -2147483648:
                                            switch fi30 {
                                            case 42:
                                                v2 = v2
                                            }
                                        }
                                        _ = v51
                                    }
                                    var v53 float32 = -2.25
                                    _ = v53
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = (("" + (("\n\t" + "foo bar") + "a")) + "foo bar")
                                    }
                                    v2 -= (v2 >> ((18446744073709551615) & 63))
                                    v2 += v2
                                }
                            }
                        } else {
                            fmt.Println(v2, fi30)
                        }
                    }
                }
            case 694:
                v54 := GenS13{F0: 42, F1: 2, F2: 42}
                _ = v54
            default:
                v2 = v2
            }
            switch (fi30 - (fi30 + 32768)) {
            case 1:
                var v55 int32 = 0
                _ = v55
            }
            v2 -= v2
        }
        _ = fi30
    }
    fmt.Println(v2)
    switch -9223372036854775808 {
    case 52:
        var v56 string = "hello"
        _ = v56
    case 32767:
        var v57 bool = (1 >= 34)
        _ = v57
    default:
        if false {
            if (648 != 4294967295) {
                for fi58 := 0; fi58 < 11; fi58++ {
                    for fi59 := 0; fi59 < 7; fi59++ {
                        var v60 int = int(v2)
                        if (!true) {
                            fmt.Println(v2)
                            v61 := GenS1{F0: 42, F1: 793}
                            var v62 uint32 = 2147483647
                            _ = v61
                            _ = v62
                        }
                        _ = fi59
                        _ = v60
                    }
                    _ = fi58
                }
                var v63 uint8 = 127
                _ = v63
            } else {
                v64 := GenS1{F0: 934, F1: 666}
                _ = v64
            }
            var v65 bool = false
            v66 := GenS13{F0: -99, F1: 767, F2: 4294967295}
            _ = v65
            _ = v66
        } else {
            v2 = (-(v2))
            v2 = v2
            for fi67 := 0; fi67 < 10; fi67++ {
                v1 = append(v1, "foo bar")
                _ = fi67
            }
        }
    }
    fmt.Println(v2)
    var v68 int8 = 42
    for fi69 := 0; fi69 < 18; fi69++ {
        for fi70 := 0; fi70 < 4; fi70++ {
            _ = len(v3)
            switch "hello" {
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = (("hello" + "\n\t") + ("hello" + ("foo bar" + "")))
                }
            case "a":
                var v71 bool = (true && false)
                _ = v71
            default:
                v72 := GenS13{F0: (v68 & 0), F1: 65535, F2: 1}
                _ = v72
            }
            _ = fi70
        }
        if (32767 == -1) {
            v68 = v68
        }
        var v73 uint8 = 62
        _ = fi69
        _ = v73
    }
    for fi74 := 0; fi74 < 2; fi74++ {
        var v75 int32 = 555
        fmt.Println(v68, v68, fi74)
        if (false || (!(true && false))) {
            var v76 uint8 = 0
            var v77 uint64 = 403
            v77 = (v77 ^ 690)
            _ = v76
            _ = v77
        } else {
            for fi78 := 0; fi78 < 17; fi78++ {
                var v79 uint16 = 2
                if false {
                    if false {
                        v80 := GenS12{F0: "foo bar", F1: 32767, F2: 331}
                        if (v75 > -2147483648) {
                            var v81 string = "hello"
                            v82 := GenS4{F0: v2}
                            if len(v1) > 0 {
                                v1[0] = v81
                            }
                            fmt.Println(v68, fi74)
                            var v83 int64 = v2
                            _ = v81
                            _ = v82
                            _ = v83
                        } else {
                            v75 += (v75 / ((v75) | 1))
                            v1 = append(v1, (("a" + "") + "a"))
                        }
                        var v84 int32 = v75
                        _ = v80
                        _ = v84
                    }
                } else {
                    var v85 uint64 = 664
                    var v86 uint16 = v79
                    _ = v85
                    _ = v86
                }
                _ = fi78
                _ = v79
            }
            v75 -= 2147483647
            if true {
                var v87 string = "\n\t"
                _ = v87
            }
        }
        _ = fi74
        _ = v75
    }
    v68 -= v68
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v68
}
