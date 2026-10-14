package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int32; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 string; F1 int16; F2 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 int8; F2 string }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 string }

func main() {
    v1 := []uint32{65535, 2147483648, 2147483647}
    var v2 string = (("foo bar" + "") + ("\n\t" + "\n\t"))
    var v3 uint = 65535
    var v4 int64 = 719
    switch v2 {
    case "":
        var v5 uint32 = 1
        _ = v5
    case "hello":
        if true {
            switch 961 {
            case 128:
                fmt.Println(v3, v4, v2)
            case 2147483647:
                for fi6 := 0; fi6 < 10; fi6++ {
                    for fi7 := 0; fi7 < 1; fi7++ {
                        v3 += 42
                        if ((-(v4)) < v4) {
                            if (true && (!false)) {
                                var v8 int = 32768
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    var v9 uint16 = 959
                                    _ = v9
                                case "\n\t":
                                    v10 := GenS1{F0: 100}
                                    _ = v10
                                case "hello":
                                    var v11 float64 = (0.1 / 1.5)
                                    _ = v11
                                }
                                _ = v8
                            } else {
                                fmt.Println(v3, fi7, v4)
                                fmt.Println(fi6)
                            }
                        }
                        _ = fi7
                    }
                    _ = fi6
                }
            case 9223372036854775807:
                if false {
                    switch 127 {
                    case -1:
                        v4 -= 966
                    case 0:
                        v4 = 636
                    case -191:
                        var v12 bool = ((v4 >= v4) && (!(v4 != -184)))
                        _ = v12
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 = v2
                }
            }
        } else {
            v4 = v4
            var v13 uint16 = 2
            _ = v13
        }
    case "foo bar":
        if false {
            if false {
                if (4294967295 <= 0) {
                    if (!false) {
                        var v14 bool = false
                        _ = len(v1)
                        _ = v14
                    }
                    var v15 int16 = -32768
                    _ = v15
                }
                v16 := GenS1{F0: 127}
                _ = v16
            } else {
                v17 := GenS1{F0: 100}
                _ = v17
            }
            v18 := GenS1{F0: 0}
            _ = v18
        }
    }
    if false {
        var v19 uint16 = 32767
        v20 := GenS1{F0: 42}
        v4 = 937
        _ = v19
        _ = v20
    }
    var v21 string = ((v2 + (v2 + v2)) + v2)
    for fi22 := 0; fi22 < 14; fi22++ {
        fmt.Println(v2)
        v3 -= v3
        _ = fi22
    }
    v23 := GenS2{F0: 255, F1: 1, F2: "foo bar"}
    for fi24 := 0; fi24 < 20; fi24++ {
        fmt.Println(fi24, v2, v2)
        fmt.Println(fi24, v3)
        _ = fi24
    }
    var v25 uint = (v3 ^ v3)
    v2 = v2
    if (false || (3.14159 != 3.14159)) {
        v26 := GenS1{F0: 0}
        var v27 uint8 = 128
        _ = v26
        _ = v27
    }
    var v28 uint16 = uint16(v25)
    v1 = append(v1, 4294967295)
    switch 32767 {
    case 128:
        var v29 int16 = -32768
        _ = v29
    }
    switch (v2 + ("\n\t" + (v2 + ((v2 + (v21 + v21)) + v21)))) {
    case "":
        if ((!true) && false) {
            if false {
                v21 = ("a" + v2)
                for fi30 := 0; fi30 < 4; fi30++ {
                    fmt.Println(v3, v25, v25)
                    for fi31 := 0; fi31 < 10; fi31++ {
                        v32 := GenS2{F0: v28, F1: 597, F2: (((v2 + ("\n\t" + "hello")) + ("hello" + v2)) + (v2 + ((("" + "\n\t") + ((v2 + (("foo bar" + (v21 + (((v21 + "\n\t") + ("\n\t" + "\n\t")) + v21))) + v21)) + (v2 + "a"))) + "foo bar")))}
                        v3 -= (^(v25))
                        v33 := GenS3{F0: (fi30 / ((-270) | 1))}
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    fmt.Println(v25)
                    _ = fi30
                }
            }
            v25 = v25
        } else {
            if false {
                for fi34 := 0; fi34 < 16; fi34++ {
                    v3 = v25
                    _ = fi34
                }
                switch (("" + (v2 + ((v2 + (v21 + "hello")) + (v2 + v2)))) + (v2 + v2)) {
                case "hello":
                    switch "" {
                    case "\n\t":
                        v2 = ("" + v21)
                    case "foo bar":
                        fmt.Println(v4, v4)
                    case "":
                        if (!("a" >= v21)) {
                            switch v21 {
                            case "":
                                v21 = (("" + v2) + (("\n\t" + v21) + ((v21 + v2) + (v2 + v21))))
                            }
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                        } else {
                            v35 := GenS2{F0: v28, F1: 2147483647, F2: (v21 + ((v2 + "") + (("" + v2) + ((("hello" + (v2 + "foo bar")) + v2) + ("" + (v2 + (v2 + v21)))))))}
                            v1 = append(v1, uint32(v3))
                            _ = v35
                        }
                    default:
                        if ((!true) && (67 > 115)) {
                            if true {
                                switch "hello" {
                                case "a":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case "":
                                    var v36 float32 = (3.14159 - (((0.1 - 1.5) - (-2.25 + -2.25)) + 0.0))
                                    _ = v36
                                default:
                                    var v37 uint = v3
                                    _ = v37
                                }
                            } else {
                                fmt.Println(v2, v3, v4)
                                if ((^(v4)) > v4) {
                                    if (!(18446744073709551615 == 4294967295)) {
                                        v28 = (^(v28))
                                    } else {
                                        var v38 float32 = 0.1
                                        v39 := GenS4{F0: (v2 + "a"), F1: 42}
                                        _ = v38
                                        _ = v39
                                    }
                                } else {
                                    for fi40 := 0; fi40 < 18; fi40++ {
                                        var v41 string = "\n\t"
                                        if (true && false) {
                                            for fi42 := 0; fi42 < 9; fi42++ {
                                                v43 := GenS4{F0: (v2 + (v2 + v21)), F1: 42}
                                                var v44 uint8 = 1
                                                _ = fi42
                                                _ = v43
                                                _ = v44
                                            }
                                            v45 := GenS1{F0: 100}
                                            _ = v45
                                        } else {
                                            var v46 uint16 = v28
                                            _ = v46
                                        }
                                        v3 -= (v25 << (((v3 ^ (^(v25)))) & 63))
                                        _ = fi40
                                        _ = v41
                                    }
                                    var v47 uint64 = 1
                                    v1 = append(v1, 2)
                                    _ = v47
                                }
                            }
                            var v48 float64 = 1.5
                            _ = v48
                        }
                    }
                case "foo bar":
                    for fi49 := 0; fi49 < 13; fi49++ {
                        switch ("foo bar" + v2) {
                        case "":
                            v50 := GenS5{F0: ((v2 + "") + ""), F1: 1000, F2: 2}
                            _ = v50
                        case "\n\t":
                            if true {
                                var v51 int64 = (v4 + 1)
                                v52 := GenS3{F0: fi49}
                                var v53 int32 = 0
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                        }
                        _ = fi49
                    }
                case "":
                    v54 := GenS4{F0: ((("" + v2) + v2) + v21), F1: 0}
                    _ = v54
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            } else {
                if false {
                    for fi55 := 0; fi55 < 19; fi55++ {
                        var v56 string = v21
                        switch fi55 {
                        case 32768:
                            fmt.Println(v25, v28)
                        case -9223372036854775808:
                            v3 += (v3 * (v3 / ((v25) | 1)))
                        }
                        _ = fi55
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 10; fi57++ {
                        switch 32768 {
                        case 9223372036854775807:
                            switch "\n\t" {
                            case "":
                                v3 = 175
                            case "a":
                                v21 = v2
                            case "hello":
                                v1 = append(v1, uint32(v3))
                            }
                        case 0:
                            if (v28 > v28) {
                                fmt.Println(v4, v4)
                                v58 := GenS6{F0: 127}
                                _ = v58
                            } else {
                                v1 = append(v1, 1)
                            }
                        case 532:
                            v59 := GenS4{F0: v2, F1: -1}
                            _ = v59
                        }
                        v60 := GenS6{F0: 1}
                        switch "foo bar" {
                        case "hello":
                            v3 += v3
                        case "":
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                        default:
                            if false {
                                v2 = "a"
                                if (v2 < (((v2 + v2) + "foo bar") + v21)) {
                                    fmt.Println(v3)
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                fmt.Println(v2, v28)
                            } else {
                                var v61 int = -509
                                for fi62 := 0; fi62 < 20; fi62++ {
                                    switch "foo bar" {
                                    case "\n\t":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "":
                                        v63 := GenS3{F0: (fi62 / ((fi57) | 1))}
                                        v64 := GenS1{F0: -109}
                                        _ = v63
                                        _ = v64
                                    default:
                                        v21 = ((v21 + v21) + v2)
                                    }
                                    _ = fi62
                                }
                                _ = v61
                            }
                        }
                        _ = fi57
                        _ = v60
                    }
                } else {
                    if (4294967295 > 42) {
                        for fi65 := 0; fi65 < 15; fi65++ {
                            switch (fi65 / ((32768) | 1)) {
                            case 2147483647:
                                switch (fi65 & fi65) {
                                case 9223372036854775807:
                                    var v66 float32 = 1.5
                                    _ = v66
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = 18
                                    }
                                }
                            case -1:
                                v2 = (v2 + "hello")
                            default:
                                v67 := GenS3{F0: fi65}
                                _ = v67
                            }
                            v68 := GenS7{F0: v28, F1: -128, F2: "hello"}
                            var v69 bool = (v2 >= v21)
                            _ = fi65
                            _ = v68
                            _ = v69
                        }
                        for fi70 := 0; fi70 < 5; fi70++ {
                            v71 := GenS8{F0: v2, F1: -392}
                            v72 := GenS9{F0: (("" + ((v2 + (v21 + v2)) + (((v2 + "foo bar") + v2) + v2))) + ((v2 + v2) + ""))}
                            var v73 int = (fi70 % ((2147483647) | 1))
                            _ = fi70
                            _ = v71
                            _ = v72
                            _ = v73
                        }
                        switch ("" + v2) {
                        case "hello":
                            v2 = v21
                        default:
                            switch 32768 {
                            case 2147483647:
                                if false {
                                    switch (v2 + ((v21 + (v21 + "a")) + ("hello" + v2))) {
                                    case "a":
                                        if (false && true) {
                                            switch v21 {
                                            case "":
                                                fmt.Println(v25)
                                            case "\n\t":
                                                var v74 uint8 = 199
                                                _ = v74
                                            default:
                                                var v75 uint = (^(v25))
                                                _ = v75
                                            }
                                            var v76 uint = 2
                                            _ = len(v1)
                                            _ = v76
                                        }
                                    case "hello":
                                        fmt.Println(v3, v4)
                                    case "":
                                        var v77 uint16 = v28
                                        _ = v77
                                    default:
                                        v4 = (v4 & (v4 >> ((255) & 63)))
                                    }
                                } else {
                                    for fi78 := 0; fi78 < 13; fi78++ {
                                        var v79 int = 128
                                        _ = fi78
                                        _ = v79
                                    }
                                    v4 = -9223372036854775808
                                }
                            case 42:
                                fmt.Println(v2)
                            case -646:
                                v25 += v3
                            }
                        }
                    }
                    v25 = v3
                }
                v4 += -1
            }
            v2 = ((v21 + ("" + v2)) + "\n\t")
        }
    default:
        var v80 float64 = 1.5
        _ = v80
    }
    for fi81 := 0; fi81 < 9; fi81++ {
        var v82 uint32 = 404
        fmt.Println(v82, v82)
        var v83 uint = v3
        _ = fi81
        _ = v82
        _ = v83
    }
    var v84 int = 995
    for fi85 := 0; fi85 < 10; fi85++ {
        for fi86 := 0; fi86 < 13; fi86++ {
            fmt.Println(fi85, fi86, v84)
            fmt.Println(v2, v3, fi86)
            _ = fi86
        }
        v21 = "\n\t"
        _ = fi85
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v84
}
