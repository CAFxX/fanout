package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int16 }
type GenS2 struct { F0 uint8; F1 int; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 string; F2 uint8 }
type GenS6 struct { F0 uint; F1 uint8; F2 int64 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16; F1 int; F2 uint32 }
type GenS10 struct { F0 int8; F1 int32 }
type GenS11 struct { F0 int; F1 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 int = -1
    var v3 uint = 0
    v2 = (v2 | v2)
    v3 -= (v3 % (((v3 + v3)) | 1))
    if (int64(v3) != 9223372036854775807) {
        v4 := GenS1{F0: 18446744073709551615, F1: 631, F2: -32768}
        v5 := GenS1{F0: 2, F1: 1, F2: 834}
        _ = v4
        _ = v5
    } else {
        v6 := GenS1{F0: 0, F1: 1, F2: -1}
        v2 -= v2
        fmt.Println(v3, v2)
        _ = v6
    }
    switch ("" + "foo bar") {
    case "hello":
        switch 9223372036854775807 {
        case 552:
            v7 := []int{v2, v2, (-(v2))}
            _ = v7
        case 9223372036854775807:
            var v8 uint16 = 2
            _ = v8
        case 516:
            var v9 float64 = 3.14159
            _ = v9
        }
    default:
        if true {
            v3 -= (v3 >> ((v3) & 63))
            switch v2 {
            case 42:
                if ((float64(0.1) - (float64(float32(1e10)) - 0.1)) == (1.5 / 1.5)) {
                    v2 = v2
                    if false {
                        v10 := GenS1{F0: 4294967295, F1: 1, F2: -225}
                        var v11 string = ""
                        _ = v10
                        _ = v11
                    } else {
                        v12 := GenS1{F0: 334, F1: 1000, F2: 32767}
                        _ = v12
                    }
                }
            case 0:
                var v13 uint = 488
                _ = v13
            case 2147483647:
                v2 -= (v2 - (v2 - (v2 % ((v2) | 1))))
            default:
                delete(v1, "foo bar")
            }
        } else {
            var v14 uint64 = 0
            _ = v14
        }
    }
    v3 += (v3 << ((v3) & 63))
    v3 = v3
    switch -1 {
    case 128:
        for fi15 := 0; fi15 < 16; fi15++ {
            if (47 <= 42) {
                v16 := []int{(fi15 << (((v3 ^ (v3 % ((v3) | 1)))) & 63)), 32767, -9223372036854775808}
                _ = v16
            } else {
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 int64 = 42
                    if false {
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v18, fi15)
                        case "":
                            var v19 string = (("hello" + "hello") + "a")
                            _ = v19
                        case "foo bar":
                            switch ("foo bar" + "a") {
                            case "a":
                                if (7 <= 42) {
                                    fmt.Println(v2, v18, fi17)
                                } else {
                                    v20 := GenS1{F0: 0, F1: -980, F2: -1}
                                    if (true && true) {
                                        v1["foo bar"] = (fi17 << (((v3 >> ((v3) & 63))) & 63))
                                    } else {
                                        if false {
                                            if (true || (-2.25 > (0.0 / 1.0))) {
                                                var v21 uint = (v3 | 65535)
                                                _ = v21
                                            }
                                            var v22 string = "foo bar"
                                            if ((1 != 1000) && true) {
                                                switch ((v22 + "hello") + v22) {
                                                case "a":
                                                    if true {
                                                        fmt.Println(v2)
                                                        switch "hello" {
                                                        case "hello":
                                                            var v23 int16 = -1
                                                            v24 := GenS1{F0: 0, F1: (v23 - v23), F2: 1000}
                                                            _ = v23
                                                            _ = v24
                                                        case "":
                                                            v25 := GenS1{F0: 96, F1: 550, F2: -32768}
                                                            _ = v25
                                                        }
                                                    } else {
                                                        v26 := GenS1{F0: 764, F1: 0, F2: 1}
                                                        switch v22 {
                                                        case "\n\t":
                                                            if false {
                                                                v18 = (v18 * v18)
                                                                v3 = v3
                                                            }
                                                        case "a":
                                                            switch v22 {
                                                            case "\n\t":
                                                                var v27 bool = false
                                                                _ = v27
                                                            case "foo bar":
                                                                v28 := GenS1{F0: 1, F1: 42, F2: 441}
                                                                _ = v28
                                                            }
                                                        case "foo bar":
                                                            v29 := []int{-1, fi17, v2, (fi17 % ((v2) | 1))}
                                                            _ = v29
                                                        default:
                                                            fmt.Println(v22)
                                                        }
                                                        _ = v26
                                                    }
                                                case "foo bar":
                                                    v3 -= (v3 * (^(v3)))
                                                case "":
                                                    v30 := GenS1{F0: 0, F1: 147, F2: 32767}
                                                    _ = v30
                                                default:
                                                    v18 += (v18 & 509)
                                                }
                                                v3 = (v3 - (v3 % (((v3 + v3)) | 1)))
                                                v31 := GenS1{F0: 815, F1: 1, F2: -540}
                                                _ = v31
                                            } else {
                                                v3 = v3
                                                fmt.Println(v22)
                                            }
                                            _ = v22
                                        }
                                        v32 := []int{fi17, 642}
                                        _ = v32
                                    }
                                    _ = len(v1)
                                    v1[("\n\t" + "\n\t")] = 772
                                    _ = v20
                                }
                            default:
                                v2 -= (^(fi17))
                            }
                        }
                    }
                    _ = fi17
                    _ = v18
                }
            }
            var v33 int = 32768
            _ = fi15
            _ = v33
        }
    case -9223372036854775808:
        v3 = (v3 & v3)
    case -1:
        v34 := GenS2{F0: uint8(v3), F1: (v2 & (v2 << ((18446744073709551615) & 63))), F2: 127}
        _ = v34
    default:
        var v35 int32 = 1
        _ = v35
    }
    v36 := GenS3{F0: 32767}
    v3 = (v3 & (v3 + (^(v3))))
    var v37 int = (-(v2))
    if true {
        v37 += (v37 % ((-9223372036854775808) | 1))
    } else {
        fmt.Println(v37, v2)
        if (!(-721 < 1)) {
            v2 = (-(v37))
            var v38 int64 = 9223372036854775807
            _ = v38
        }
        for fi39 := 0; fi39 < 18; fi39++ {
            v40 := []int{(v37 - v2), fi39}
            if len(v40) >= 2 {
                v40 = v40[1:len(v40)]
            }
            _ = fi39
            _ = v40
        }
    }
    if (18446744073709551615 > 404) {
        for fi41 := 0; fi41 < 1; fi41++ {
            v42 := GenS4{F0: "foo bar"}
            v43 := GenS5{F0: 255, F1: "", F2: 82}
            switch ("" + "\n\t") {
            case "":
                v1["\n\t"] = v2
            default:
                v2 = -232
            }
            _ = fi41
            _ = v42
            _ = v43
        }
        if (false && (false && (1 < -466))) {
            var v44 bool = true
            v45 := GenS1{F0: 0, F1: 32767, F2: -1}
            switch (("a" + ("a" + "hello")) + "a") {
            case "a":
                for fi46 := 0; fi46 < 12; fi46++ {
                    switch "foo bar" {
                    case "":
                        v1[(("\n\t" + "") + "")] = (v37 + 128)
                    default:
                        _ = len(v1)
                    }
                    v3 -= (v3 >> ((42) & 63))
                    v2 = v37
                    _ = fi46
                }
            case "":
                var v47 float32 = float32(((1.5 + 1e10) - 1e10))
                _ = v47
            default:
                switch "\n\t" {
                case "\n\t":
                    v2 -= v2
                }
            }
            _ = v44
            _ = v45
        }
    } else {
        for fi48 := 0; fi48 < 10; fi48++ {
            fmt.Println(fi48)
            v49 := GenS2{F0: 127, F1: 32768, F2: -128}
            _ = fi48
            _ = v49
        }
        if true {
            v50 := GenS6{F0: (v3 ^ v3), F1: 88, F2: 9223372036854775807}
            if (2147483647 <= -2147483648) {
                v51 := GenS3{F0: uint16(v3)}
                _ = v51
            } else {
                if true {
                    v1[("" + "foo bar")] = v37
                    v52 := GenS7{F0: v2, F1: 2}
                    _ = v52
                }
                delete(v1, "hello")
            }
            _ = v50
        } else {
            v2 = v2
        }
        if (true || (!false)) {
            var v53 int = 21
            _ = v53
        } else {
            switch ("\n\t" + "a") {
            case "":
                v54 := GenS2{F0: 246, F1: (v37 >> ((v3) & 63)), F2: 42}
                _ = v54
            case "a":
                v55 := GenS8{F0: 128}
                _ = v55
            }
            if false {
                switch -378 {
                case 675:
                    if false {
                        fmt.Println(v2)
                        v3 = 33
                        if ((v2 << ((v3) & 63)) == 4294967295) {
                            switch "" {
                            case "hello":
                                var v56 uint = v3
                                _ = v56
                            case "\n\t":
                                for fi57 := 0; fi57 < 9; fi57++ {
                                    var v58 string = ""
                                    v59 := GenS2{F0: 42, F1: v37, F2: -99}
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                }
                            default:
                                if (false && false) {
                                    for fi60 := 0; fi60 < 8; fi60++ {
                                        delete(v1, ("\n\t" + ((("" + "foo bar") + ("\n\t" + ("a" + "\n\t"))) + (("" + "foo bar") + ""))))
                                        _ = fi60
                                    }
                                    if (!(!(-128 >= 0))) {
                                        switch v2 {
                                        case -9223372036854775808:
                                            v61 := []int{v37}
                                            _ = v61
                                        case 1:
                                            v62 := GenS9{F0: 1, F1: (v2 >> (((v3 / (((^(v3))) | 1))) & 63)), F2: 4294967295}
                                            _ = v62
                                        case -571:
                                            fmt.Println(v37)
                                        default:
                                            if true {
                                                switch (v37 * v37) {
                                                case 127:
                                                    var v63 int8 = 100
                                                    _ = v63
                                                case -1:
                                                    for fi64 := 0; fi64 < 13; fi64++ {
                                                        var v65 float32 = 0.0
                                                        fmt.Println(v2, v65)
                                                        _ = fi64
                                                        _ = v65
                                                    }
                                                case 32768:
                                                    v66 := GenS1{F0: 754, F1: -32768, F2: 32767}
                                                    _ = v66
                                                default:
                                                    v67 := GenS10{F0: 8, F1: 0}
                                                    _ = v67
                                                }
                                            }
                                        }
                                        fmt.Println(v37)
                                    }
                                } else {
                                    fmt.Println(v3, v3, v37)
                                    var v68 int64 = 462
                                    switch ("foo bar" + ("a" + "a")) {
                                    case "foo bar":
                                        v69 := GenS11{F0: v2, F1: v68}
                                        _ = v69
                                    case "hello":
                                        if false {
                                            switch "a" {
                                            case "hello":
                                                v3 = 772
                                            default:
                                                v68 = (v68 ^ (v68 ^ v68))
                                            }
                                        }
                                    case "":
                                        if true {
                                            if (-2.25 < (0.1 * 0.1)) {
                                                var v70 float64 = 1e10
                                                _ = v70
                                            }
                                        }
                                    default:
                                        if (false && false) {
                                            var v71 uint32 = 4294967295
                                            fmt.Println(v37, v71)
                                            _ = v71
                                        }
                                    }
                                    _ = v68
                                }
                            }
                            fmt.Println(v3, v3)
                            v2 += v37
                        }
                    }
                case -2147483648:
                    for fi72 := 0; fi72 < 11; fi72++ {
                        for fi73 := 0; fi73 < 11; fi73++ {
                            v2 = int(v3)
                            v2 = -375
                            switch "hello" {
                            case "a":
                                var v74 int64 = -1
                                _ = v74
                            case "":
                                v37 = (^(fi73))
                            default:
                                v2 += fi72
                            }
                            _ = fi73
                        }
                        _ = fi72
                    }
                case 1:
                    var v75 uint16 = 2
                    _ = v75
                }
                v37 = -366
                fmt.Println(v3)
            } else {
                var v76 uint = (v3 / ((216) | 1))
                v77 := []int{-956, int(v76), v2}
                if false {
                    var v78 int32 = -337
                    v79 := GenS11{F0: (v37 | 32768), F1: 718}
                    _ = v78
                    _ = v79
                }
                _ = v76
                _ = v77
            }
            var v80 uint64 = 42
            _ = v80
        }
    }
    fmt.Println(v3, v2, v2)
    for fi81 := 0; fi81 < 16; fi81++ {
        for fi82 := 0; fi82 < 16; fi82++ {
            switch fi82 {
            case 715:
                var v83 uint32 = 721
                _ = v83
            case -792:
                var v84 bool = (1e10 < ((3.14159 + 3.14159) / 1e10))
                _ = v84
            case 42:
                if (42 < 0) {
                    var v85 int32 = 2147483647
                    _ = v85
                }
            }
            var v86 uint = (v3 - v3)
            var v87 int = fi81
            _ = fi82
            _ = v86
            _ = v87
        }
        if false {
            if (42 > 65535) {
                fmt.Println(v3, v37, v3)
                if ((1e10 / 1.0) == (float32(1e10) - 3.14159)) {
                    var v88 uint8 = 1
                    for fi89 := 0; fi89 < 19; fi89++ {
                        v37 = fi89
                        _ = fi89
                    }
                    _ = v88
                } else {
                    v90 := []int{}
                    _ = v90
                }
                if (1 > -213) {
                    var v91 int = (fi81 | v2)
                    _ = v91
                }
            }
            v1[(((("foo bar" + (("hello" + "hello") + "\n\t")) + "a") + "\n\t") + "a")] = (fi81 / ((v37) | 1))
        }
        var v92 bool = true
        _ = fi81
        _ = v92
    }
    v93 := []int{-705}
    v1[""] = (-(v37))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v93:", len(v93), v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v93
}
