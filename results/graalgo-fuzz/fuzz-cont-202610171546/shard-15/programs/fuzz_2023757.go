package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int8; F2 int32 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint }
type GenS4 struct { F0 uint; F1 int16; F2 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool }

func main() {
    v1 := []string{(("a" + "") + (("\n\t" + ("hello" + "hello")) + "hello"))}
    v2 := make(map[uint32]int)
    v3 := []uint32{65535, 4294967295, 1, 65535}
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 int8 = 100
        if ((2 == 127) && (1e10 >= 0.0)) {
            for fi6 := 0; fi6 < 18; fi6++ {
                v5 += 0
                _ = fi6
            }
            switch (fi4 % (((fi4 >> ((979) & 63))) | 1)) {
            case 9223372036854775807:
                switch ("hello" + "hello") {
                case "":
                    switch fi4 {
                    case -671:
                        v3 = append(v3, 0)
                    default:
                        v7 := GenS1{F0: -1}
                        _ = v7
                    }
                }
            }
            for fi8 := 0; fi8 < 20; fi8++ {
                switch 629 {
                case 32768:
                    if true {
                        var v9 uint = 272
                        _ = v9
                    }
                default:
                    var v10 int = (fi4 - (fi4 & (fi4 << ((1) & 63))))
                    _ = v10
                }
                var v11 float32 = 1.5
                _ = fi8
                _ = v11
            }
        }
        v12 := GenS2{F0: ("" + "a"), F1: (v5 + v5), F2: 1}
        _ = fi4
        _ = v12
        _ = v5
    }
    switch 4294967295 {
    case 208:
        var v13 int16 = 1
        _ = v13
    case 1:
        var v14 int64 = 92
        _ = v14
    default:
        v2[65535] = 32768
    }
    var v15 uint = 2
    v15 += (v15 << (((^(v15))) & 63))
    switch 32767 {
    case 4294967295:
        if true {
            v15 = 1
            var v16 uint8 = 255
            for fi17 := 0; fi17 < 10; fi17++ {
                if true {
                    var v18 int16 = 1
                    _ = v18
                }
                if true {
                    v15 += (v15 * v15)
                    if (0 >= -2147483648) {
                        v16 += v16
                        v15 += v15
                        var v19 string = "foo bar"
                        _ = v19
                    } else {
                        var v20 bool = true
                        delete(v2, 0)
                        v2[2147483648] = fi17
                        _ = v20
                    }
                }
                _ = fi17
            }
            _ = v16
        } else {
            fmt.Println(v15)
        }
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        v15 -= 395
        _ = fi21
    }
    fmt.Println(v15)
    for fi22 := 0; fi22 < 12; fi22++ {
        for fi23 := 0; fi23 < 6; fi23++ {
            if true {
                if ((!false) || (!false)) {
                    switch "hello" {
                    case "\n\t":
                        v24 := GenS1{F0: 1}
                        _ = v24
                    case "hello":
                        fmt.Println(v15, fi22, fi22)
                    case "foo bar":
                        v15 = v15
                    default:
                        fmt.Println(v15, fi23)
                    }
                }
                if (0 != 1) {
                    _ = len(v3)
                    if true {
                        v15 -= 0
                        v15 += v15
                        v15 = 42
                        var v25 uint32 = 2147483648
                        v26 := GenS2{F0: ("a" + ""), F1: -1, F2: 1}
                        var v27 uint16 = 720
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        if true {
                            var v28 bool = (!(fi23 < (fi23 * 9223372036854775807)))
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                            v29 := GenS3{F0: 4294967295, F1: 42, F2: 590}
                            _ = v28
                            _ = v29
                        }
                    }
                }
            }
            switch "" {
            case "":
                var v30 uint = v15
                _ = v30
            case "hello":
                v1 = append(v1, "hello")
            default:
                v31 := GenS1{F0: 42}
                if (false && true) {
                    var v32 int64 = -9223372036854775808
                    if ((-(fi23)) > 32768) {
                        switch 32767 {
                        case 0:
                            v32 = 9223372036854775807
                        case 127:
                            v15 -= v15
                        }
                        var v33 int = (-(fi22))
                        var v34 int32 = 404
                        _ = v33
                        _ = v34
                    }
                    if (42 <= 42) {
                        var v35 uint8 = 2
                        if true {
                            v36 := GenS4{F0: (v15 ^ 0), F1: 1000, F2: 154}
                            if (!false) {
                                if (false || false) {
                                    v35 = v35
                                }
                                v37 := GenS2{F0: "hello", F1: -128, F2: 0}
                                _ = v37
                            }
                            _ = v36
                        } else {
                            var v38 float32 = ((3.14159 * float32(0.0)) / 1.0)
                            _ = v38
                        }
                        _ = v35
                    }
                    _ = v32
                } else {
                    _ = len(v2)
                    _ = len(v1)
                }
                _ = v31
            }
            var v39 uint8 = 127
            var v40 string = ("hello" + "hello")
            _ = fi23
            _ = v39
            _ = v40
        }
        _ = fi22
    }
    var v41 uint = v15
    v41 = v41
    var v42 int8 = 100
    if ((340 != 42) && (!((false && false) && true))) {
        switch 32768 {
        case 9223372036854775807:
            v43 := GenS5{F0: 0}
            _ = v43
        default:
            if (127 == v42) {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                if true {
                    switch "\n\t" {
                    case "a":
                        switch 32768 {
                        case 42:
                            switch ("hello" + "\n\t") {
                            case "":
                                v42 = (v42 >> ((94) & 63))
                            case "foo bar":
                                fmt.Println(v15, v15, v15)
                            default:
                                var v44 int = 42
                                _ = v44
                            }
                        case -209:
                            var v45 uint32 = 2
                            _ = v45
                        case 4294967295:
                            if false {
                                var v46 int16 = -32768
                                _ = v46
                            } else {
                                fmt.Println(v42)
                                if (-828 == -1) {
                                    for fi47 := 0; fi47 < 2; fi47++ {
                                        switch 32767 {
                                        case 0:
                                            v48 := GenS6{F0: (^(fi47)), F1: true}
                                            _ = v48
                                        case 4294967295:
                                            v15 = 255
                                        case 500:
                                            _ = len(v3)
                                        default:
                                            v49 := GenS2{F0: "", F1: (-(v42)), F2: -2147483648}
                                            _ = v49
                                        }
                                        v50 := GenS6{F0: (-(fi47)), F1: (255 > 2)}
                                        _ = fi47
                                        _ = v50
                                    }
                                    var v51 uint = 18446744073709551615
                                    _ = v51
                                } else {
                                    for fi52 := 0; fi52 < 3; fi52++ {
                                        if (!true) {
                                            var v53 bool = (!false)
                                            _ = v53
                                        }
                                        v1 = append(v1, "\n\t")
                                        _ = fi52
                                    }
                                }
                                v54 := GenS3{F0: 2, F1: -9223372036854775808, F2: (v15 * (v41 + (v15 / (((^(v41))) | 1))))}
                                _ = v54
                            }
                        default:
                            if false {
                                var v55 uint8 = 245
                                var v56 uint = v41
                                _ = v55
                                _ = v56
                            }
                        }
                    case "":
                        switch (((("hello" + "a") + "foo bar") + "hello") + "") {
                        case "a":
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                        case "\n\t":
                            fmt.Println(v15, v42, v15)
                        default:
                            var v57 uint = v41
                            _ = v57
                        }
                    case "\n\t":
                        var v58 uint16 = 65535
                        _ = v58
                    default:
                        fmt.Println(v41)
                    }
                } else {
                    switch 42 {
                    case 1:
                        var v59 float32 = ((((1.5 / 0.1) - (float32(-2.25) * 0.0)) - -2.25) + float32(1.5))
                        _ = v59
                    case 0:
                        v60 := GenS7{F0: "\n\t", F1: 32768}
                        _ = v60
                    default:
                        delete(v2, 2147483648)
                    }
                }
                if true {
                    if (-1 < 1) {
                        fmt.Println(v42, v42, v15)
                        v41 = v15
                        fmt.Println(v41)
                    } else {
                        if (!(true || (!false))) {
                            fmt.Println(v15, v15, v42)
                            var v61 int16 = -32768
                            _ = v61
                        } else {
                            var v62 uint8 = 0
                            switch -1 {
                            case -1:
                                v41 = (v15 | 1)
                            default:
                                switch ("foo bar" + "foo bar") {
                                case "":
                                    fmt.Println(v62)
                                case "a":
                                    v62 -= (v62 / (((v62 - (v62 - v62))) | 1))
                                case "foo bar":
                                    v2[2147483648] = 1
                                default:
                                    if ((false && true) && true) {
                                        fmt.Println(v42, v41, v41)
                                    }
                                }
                            }
                            if true {
                                v41 = 65535
                            }
                            _ = v62
                        }
                    }
                } else {
                    var v63 float64 = -2.25
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v63
                }
            }
        }
        switch 2147483647 {
        case 487:
            switch ("foo bar" + ("\n\t" + "a")) {
            case "hello":
                var v64 uint8 = 127
                _ = v64
            default:
                var v65 int8 = (^(v42))
                _ = v65
            }
        case 42:
            var v66 uint = v15
            _ = v66
        case -380:
            v15 -= 1
        default:
            switch ("a" + ("" + "hello")) {
            case "hello":
                v15 -= (v41 << ((v41) & 63))
            case "":
                v67 := GenS4{F0: v41, F1: 157, F2: 98}
                _ = v67
            case "a":
                if false {
                    v68 := v2[2147483648]
                    v68 -= v68
                    v15 = 0
                    _ = v68
                } else {
                    fmt.Println(v42, v41, v41)
                    v69 := GenS8{F0: 42}
                    _ = v69
                }
            default:
                fmt.Println(v15)
            }
        }
        if ((127 != 32767) || (v42 > v42)) {
            _ = len(v3)
            if (!(683 > 2147483647)) {
                for fi70 := 0; fi70 < 12; fi70++ {
                    v71 := GenS4{F0: (v15 >> (((v15 - v41)) & 63)), F1: 0, F2: 39}
                    _ = fi70
                    _ = v71
                }
                var v72 int8 = (v42 | (v42 * (v42 & (v42 ^ v42))))
                v73 := GenS9{F0: ((1000 <= 1000) || false)}
                _ = v72
                _ = v73
            } else {
                switch 439 {
                case -1:
                    _ = len(v2)
                case 32768:
                    switch -410 {
                    case -662:
                        switch "\n\t" {
                        case "":
                            var v74 int16 = 1
                            _ = v74
                        }
                    }
                case -2147483648:
                    v41 = (v15 >> ((42) & 63))
                default:
                    var v75 uint64 = 18
                    _ = v75
                }
                fmt.Println(v42, v41)
            }
        }
    }
    fmt.Println(v41, v15)
    v15 -= (^(v41))
    v42 = v42
    var v76 int64 = -9223372036854775808
    for fi77 := 0; fi77 < 20; fi77++ {
        var v78 uint8 = 124
        _ = len(v1)
        switch 4294967295 {
        case -322:
            v76 = v76
        case 127:
            v79 := GenS4{F0: (v15 - (^(v15))), F1: 1000, F2: 139}
            _ = v79
        case -930:
            fmt.Println(v15)
        }
        _ = fi77
        _ = v78
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v76
}
