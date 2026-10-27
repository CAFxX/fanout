package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS2 struct { F0 string; F1 int64; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 uint; F2 uint }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32; F1 int8 }
type GenS9 struct { F0 int64; F1 int8; F2 int32 }
func genFn0(p0 int32) int16 {
    if (255 > 128) {
        v4 := GenS1{F0: 1, F1: 4294967295, F2: 42}
        switch -9223372036854775808 {
        case 2147483647:
            v5 := GenS1{F0: 110, F1: 495, F2: int16(p0)}
            _ = v5
        case 1:
            for fi6 := 0; fi6 < 2; fi6++ {
                for fi7 := 0; fi7 < 15; fi7++ {
                    if true {
                        var v8 bool = (!(4294967295 < 131))
                        v9 := []int{32768, (fi6 / (((-(fi6))) | 1)), (fi6 << ((0) & 63)), fi7}
                        var v10 bool = (!v8)
                        _ = v10
                        _ = v8
                        _ = v9
                    }
                    p0 += 0
                    _ = fi7
                }
                _ = fi6
            }
        case -1:
            p0 = p0
        default:
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 uint32 = 993
                fmt.Println(p0, fi11, fi11)
                v13 := GenS2{F0: "a", F1: -855, F2: true}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 1; fi14++ {
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := []int{}
                p0 += 42
                _ = fi15
                _ = v16
            }
            fmt.Println(fi14, p0)
            _ = fi14
        }
        _ = v4
    }
    if true {
        v17 := GenS1{F0: 992, F1: 2, F2: 1}
        var v18 float32 = 1e10
        _ = v17
        _ = v18
    } else {
        v19 := GenS3{F0: 65535}
        _ = v19
    }
    switch 0 {
    case 127:
        if true {
            switch "a" {
            case "hello":
                v20 := []int{1, 389, -9223372036854775808}
                _ = v20
            }
        } else {
            v21 := []int{}
            v22 := GenS3{F0: 4294967295}
            for fi23 := 0; fi23 < 14; fi23++ {
                if false {
                    v24 := GenS4{F0: (false || false)}
                    _ = v24
                } else {
                    if (42 > 0) {
                        for fi25 := 0; fi25 < 1; fi25++ {
                            var v26 float32 = float32((((1.5 / 1e10) * float64(0.0)) / (3.14159 / 1.5)))
                            _ = fi25
                            _ = v26
                        }
                        _ = len(v21)
                    }
                    for fi27 := 0; fi27 < 14; fi27++ {
                        var v28 uint32 = 800
                        v28 += (^(v28))
                        p0 += p0
                        _ = fi27
                        _ = v28
                    }
                }
                if false {
                    for fi29 := 0; fi29 < 17; fi29++ {
                        switch ((((("a" + "") + "") + "foo bar") + "") + "foo bar") {
                        case "hello":
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                        case "a":
                            var v30 int8 = 101
                            _ = v30
                        case "":
                            fmt.Println(fi23)
                        default:
                            switch (fi29 ^ (fi23 << ((42) & 63))) {
                            case 2147483647:
                                var v31 uint = 42
                                var v32 int64 = 9223372036854775807
                                var v33 uint8 = 127
                                _ = v31
                                _ = v32
                                _ = v33
                            case -51:
                                switch 42 {
                                case -982:
                                    v34 := GenS2{F0: ("\n\t" + ("hello" + ("a" + "foo bar"))), F1: 232, F2: false}
                                    _ = v34
                                case 9223372036854775807:
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                case 2147483647:
                                    p0 += -2147483648
                                    v35 := GenS1{F0: 42, F1: 1, F2: -1}
                                    _ = v35
                                }
                            }
                        }
                        p0 -= p0
                        _ = fi29
                    }
                }
                _ = len(v21)
                _ = fi23
            }
            _ = v21
            _ = v22
        }
    case -735:
        v36 := []int{0, 32767}
        _ = v36
    }
    var v37 uint16 = 65535
    _ = p0
    _ = v37
    return 1
}
func genFn1(p0 int, p1 uint64) string {
    switch "foo bar" {
    case "a":
        var v38 int8 = -22
        _ = v38
    case "\n\t":
        p1 += p1
    }
    var v39 uint16 = 42
    var v40 uint = 121
    _ = p0
    _ = p1
    _ = v39
    _ = v40
    return ("" + "a")
}

func main() {
    var v1 int16 = -1
    var v2 int32 = 1
    v3 := []string{("\n\t" + "\n\t"), "hello"}
    switch -2147483648 {
    case -516:
        v1 = v1
    case 0:
        var v41 int8 = 42
        _ = v41
    case 42:
        for fi42 := 0; fi42 < 19; fi42++ {
            v1 = genFn0(v2)
            v2 = v2
            fmt.Println(v1, v1)
            _ = fi42
        }
    default:
        var v43 int32 = 0
        _ = v43
    }
    v2 = 363
    if (-678 == -1) {
        fmt.Println(v1, v1)
    }
    v44 := GenS3{F0: 2147483647}
    switch (("\n\t" + ("" + "\n\t")) + "hello") {
    case "foo bar":
        v3 = append(v3, ((("hello" + "\n\t") + (("a" + ("hello" + "foo bar")) + "hello")) + "\n\t"))
    }
    for fi45 := 0; fi45 < 15; fi45++ {
        for fi46 := 0; fi46 < 7; fi46++ {
            if true {
                switch "a" {
                case "":
                    if false {
                        if false {
                            switch 9223372036854775807 {
                            case 42:
                                v2 = (v2 >> ((65535) & 63))
                            case 4294967295:
                                switch (fi46 ^ fi46) {
                                case 32767:
                                    if false {
                                        var v47 int = (^(fi45))
                                        _ = v47
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            fmt.Println(v2)
                                        case "\n\t":
                                            fmt.Println(fi45, v1, fi45)
                                        case "a":
                                            v2 += v2
                                            var v48 uint32 = 2147483647
                                            _ = v48
                                        }
                                        if true {
                                            if (127 > 100) {
                                                fmt.Println(fi46, fi46, fi46)
                                                v49 := GenS5{F0: genFn0((-(v2))), F1: 950, F2: 2}
                                                v3 = append(v3, "\n\t")
                                                _ = v49
                                            } else {
                                                if (false && false) {
                                                    v1 -= 0
                                                    v50 := GenS2{F0: "hello", F1: -9223372036854775808, F2: false}
                                                    if false {
                                                        v51 := GenS2{F0: (("" + "") + (("hello" + "") + (("hello" + ("\n\t" + (("a" + (("foo bar" + ((("\n\t" + "\n\t") + "foo bar") + "\n\t")) + "")) + ("hello" + "hello")))) + "hello"))), F1: 416, F2: ((42 < 252) || (!((v2 * (v2 >> ((18446744073709551615) & 63))) <= v2)))}
                                                        v52 := GenS3{F0: 2}
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    _ = v50
                                                }
                                            }
                                            if (0 >= 999) {
                                                v53 := GenS2{F0: "\n\t", F1: 944, F2: false}
                                                _ = v53
                                            }
                                            fmt.Println(v1, v1)
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            var v54 int32 = 2147483647
                                            _ = v54
                                        }
                                        var v55 int16 = 1000
                                        _ = v55
                                    }
                                }
                            case -157:
                                var v56 float32 = 3.14159
                                var v57 uint64 = 18446744073709551615
                                v58 := GenS6{F0: 42, F1: 100}
                                if true {
                                    if false {
                                        _ = len(v3)
                                        if (1.5 < ((v56 * v56) / float32((0.1 + ((0.0 + (float64(v56) / ((1e10 - float64((v56 + (float32((-2.25 * float64((float32(0.0) / v56)))) - (v56 - (0.0 * v56)))))) + 3.14159))) - (1e10 * (1.5 + 1e10))))))) {
                                            switch "a" {
                                            case "foo bar":
                                                v3 = append(v3, (("foo bar" + (("foo bar" + "a") + "a")) + ""))
                                                v2 += (v2 << ((255) & 63))
                                                v2 -= (v2 + (v2 ^ -2147483648))
                                            default:
                                                switch fi45 {
                                                case -636:
                                                    var v59 uint32 = 1
                                                    _ = v59
                                                case 127:
                                                    _ = len(v3)
                                                }
                                            }
                                            if (128 >= 0) {
                                                var v60 bool = ((115 < 65535) && (!true))
                                                if len(v3) > 0 {
                                                    v3[0] = ""
                                                }
                                                var v61 uint = 42
                                                _ = v60
                                                _ = v61
                                            } else {
                                                v1 = (v1 >> ((2) & 63))
                                                if (100 < 1) {
                                                    if len(v3) > 0 {
                                                        v3[0] = "a"
                                                    }
                                                }
                                            }
                                            v62 := GenS5{F0: (v1 + (^(v1))), F1: 0, F2: 1}
                                            _ = v62
                                        } else {
                                            var v63 uint32 = 1
                                            var v64 float64 = 1.5
                                            fmt.Println(fi46, fi45, v63)
                                            _ = v63
                                            _ = v64
                                        }
                                    }
                                    if true {
                                        var v65 string = ("hello" + ((("" + ("\n\t" + "foo bar")) + "") + "a"))
                                        _ = v65
                                    }
                                    v66 := GenS7{F0: (v1 + v1)}
                                    if (!true) {
                                        fmt.Println(fi45, v57)
                                        var v67 int16 = int16(v2)
                                        switch "a" {
                                        case "":
                                            fmt.Println(v2, v1, v57)
                                        case "\n\t":
                                            v68 := GenS4{F0: true}
                                            fmt.Println(v67)
                                            _ = v68
                                        case "a":
                                            fmt.Println(fi45, fi46, fi45)
                                        default:
                                            var v69 int32 = v2
                                            _ = v69
                                        }
                                        _ = v67
                                    }
                                    _ = v66
                                }
                                _ = v56
                                _ = v57
                                _ = v58
                            default:
                                v2 = v2
                            }
                            v1 += int16(v2)
                        } else {
                            switch int(v2) {
                            case 665:
                                var v70 string = ""
                                _ = v70
                            default:
                                var v71 string = ("a" + "\n\t")
                                if (0.1 == (((0.1 * (((-2.25 / (((0.1 * float32(1.5)) - 0.0) / 0.1)) - float32(3.14159)) / (0.1 / 1.0))) - 1e10) - 1.5)) {
                                    var v72 bool = true
                                    v2 = (v2 >> ((4294967295) & 63))
                                    _ = v72
                                } else {
                                    var v73 uint32 = 2147483647
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    fmt.Println(v71)
                                    _ = v73
                                }
                                _ = v71
                            }
                        }
                        if false {
                            v1 = (v1 + 417)
                            fmt.Println(v1, v1, fi45)
                            v74 := GenS2{F0: "foo bar", F1: 9223372036854775807, F2: false}
                            _ = v74
                        } else {
                            v1 += v1
                            v1 = v1
                            switch "\n\t" {
                            case "":
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                            }
                        }
                        if true {
                            v3 = append(v3, ("\n\t" + (("hello" + ("foo bar" + "a")) + ("foo bar" + "\n\t"))))
                            switch 9223372036854775807 {
                            case 4294967295:
                                v75 := GenS6{F0: 42, F1: 100}
                                _ = v75
                            }
                        } else {
                            v3 = append(v3, "")
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v1 = 1
                        }
                    } else {
                        switch (-(fi45)) {
                        case 127:
                            v1 += genFn0((v2 | 1))
                        default:
                            switch "a" {
                            case "a":
                                switch ("a" + "hello") {
                                case "foo bar":
                                    v1 -= (v1 / ((1000) | 1))
                                case "\n\t":
                                    switch ((("foo bar" + "foo bar") + "hello") + "a") {
                                    case "a":
                                        if false {
                                            if true {
                                                v2 -= 1
                                            } else {
                                                if (306 >= 2147483647) {
                                                    fmt.Println(fi45, fi45, fi46)
                                                } else {
                                                    v1 = -1
                                                    v2 -= 2147483647
                                                    var v76 string = "foo bar"
                                                    _ = v76
                                                }
                                            }
                                            if (!false) {
                                                switch ((("hello" + "\n\t") + "\n\t") + ("a" + "foo bar")) {
                                                case "\n\t":
                                                    var v77 int = (fi46 / (((fi45 << ((1) & 63))) | 1))
                                                    _ = v77
                                                case "a":
                                                    v1 += genFn0((v2 + v2))
                                                default:
                                                    var v78 uint64 = 18446744073709551615
                                                    _ = v78
                                                }
                                                switch (fi45 % (((^(fi46))) | 1)) {
                                                case 9223372036854775807:
                                                    var v79 int16 = (v1 % ((genFn0((v2 * (v2 + v2)))) | 1))
                                                    _ = v79
                                                case 32768:
                                                    var v80 int32 = v2
                                                    _ = v80
                                                case 994:
                                                    v3 = append(v3, "hello")
                                                }
                                                var v81 string = "\n\t"
                                                _ = v81
                                            } else {
                                                v82 := GenS5{F0: (v1 % (((v1 / ((-32768) | 1))) | 1)), F1: 576, F2: 1}
                                                var v83 float64 = 0.1
                                                fmt.Println(fi45, v2)
                                                v84 := GenS7{F0: (^(v1))}
                                                _ = v82
                                                _ = v83
                                                _ = v84
                                            }
                                        } else {
                                            v2 = v2
                                            v85 := GenS3{F0: 832}
                                            _ = v85
                                        }
                                    default:
                                        var v86 uint8 = 172
                                        _ = v86
                                    }
                                default:
                                    var v87 float64 = (1e10 / 0.1)
                                    _ = v87
                                }
                                var v88 float32 = 0.1
                                if (2 != 179) {
                                    var v89 int = (fi46 << ((0) & 63))
                                    _ = v89
                                } else {
                                    v90 := GenS8{F0: v2, F1: 69}
                                    _ = v90
                                }
                                fmt.Println(v2, v1)
                                var v91 float32 = (0.1 + (((0.1 + v88) + ((-2.25 + v88) / 0.1)) * 1.5))
                                var v92 uint = 1
                                v2 += 42
                                _ = v88
                                _ = v91
                                _ = v92
                            case "\n\t":
                                v2 -= 697
                            case "":
                                v1 -= 1000
                                v1 = (^(v1))
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                            }
                        }
                        var v93 float64 = 3.14159
                        _ = v93
                    }
                case "\n\t":
                    v2 -= v2
                }
            }
            switch "\n\t" {
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                switch (-(fi45)) {
                case 32767:
                    switch "foo bar" {
                    case "foo bar":
                        v3 = append(v3, ("\n\t" + "hello"))
                    case "":
                        v94 := GenS7{F0: genFn0((v2 ^ (v2 - 1)))}
                        _ = v94
                    case "a":
                        v95 := GenS9{F0: 1, F1: 1, F2: -670}
                        _ = v95
                    }
                }
            }
            _ = fi46
        }
        _ = fi45
    }
    var v96 bool = (32767 < 255)
    var v97 int64 = 42
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v96
    _ = v97
}
