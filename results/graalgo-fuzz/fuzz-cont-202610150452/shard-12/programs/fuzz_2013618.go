package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 int8; F1 int16 }

func main() {
    v1 := []string{"", ""}
    var v2 uint16 = 255
    var v3 float32 = -2.25
    var v4 string = "a"
    v5 := []int{128, -2147483648, 856, -2147483648}
    var v6 int64 = -9223372036854775808
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (v3 <= (((1.5 + v3) + 0.1) * v3)) {
        v6 = 438
    } else {
        for fi7 := 0; fi7 < 3; fi7++ {
            switch 288 {
            case 32768:
                if false {
                    v8 := GenS1{F0: 2147483647}
                    switch (fi7 + (fi7 ^ (fi7 & fi7))) {
                    case 32768:
                        var v9 bool = true
                        _ = v9
                    case 32767:
                        var v10 bool = (1 <= 0)
                        _ = v10
                    }
                    if ((false && false) || (42 != -32768)) {
                        switch fi7 {
                        case 2147483647:
                            var v11 int = (fi7 & fi7)
                            _ = v11
                        case 1:
                            for fi12 := 0; fi12 < 17; fi12++ {
                                switch v4 {
                                case "":
                                    var v13 int = (fi12 + -710)
                                    _ = v13
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = len(v1)
                                var v14 float32 = 1.5
                                switch ("foo bar" + ("" + ("foo bar" + (v4 + v4)))) {
                                case "foo bar":
                                    var v15 uint32 = 154
                                    _ = v15
                                default:
                                    if (-128 != 43) {
                                        var v16 int32 = -1
                                        v17 := GenS1{F0: v16}
                                        var v18 float64 = 0.0
                                        v3 -= (((v3 / (float32(float64((v14 + (-2.25 + 0.0)))) * v14)) - (v14 - v3)) * ((-2.25 / float32(((v18 + (0.1 / (v18 * v18))) - (v18 - v18)))) * (0.1 + (0.1 + float32((v18 + ((((-2.25 / v18) + -2.25) + float64(v3)) * (-2.25 - float64(((((v14 + v14) + 1.5) * 3.14159) + v3))))))))))
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                    } else {
                                        if true {
                                            fmt.Println(v2, fi12, v2)
                                            v4 = ""
                                            if len(v1) > 0 {
                                                v1[0] = v4
                                            }
                                            v19 := GenS1{F0: 2147483647}
                                            if false {
                                                var v20 float32 = v14
                                                var v21 float64 = 3.14159
                                                v22 := GenS2{F0: 1}
                                                switch (fi7 & (fi7 - fi12)) {
                                                case 451:
                                                    var v23 uint8 = 0
                                                    _ = v23
                                                case -508:
                                                    var v24 uint32 = 42
                                                    _ = v24
                                                }
                                                _ = v20
                                                _ = v21
                                                _ = v22
                                            } else {
                                                var v25 int32 = -405
                                                _ = v25
                                            }
                                            v26 := GenS2{F0: v6}
                                            switch (v4 + ("a" + "foo bar")) {
                                            case "a":
                                                v3 = 1e10
                                            case "foo bar":
                                                fmt.Println(fi7, fi12, fi7)
                                            case "hello":
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            v14 += float32(0.0)
                                            _ = v19
                                            _ = v26
                                        } else {
                                            v27 := GenS3{F0: 196, F1: (v4 + "hello")}
                                            _ = v27
                                        }
                                        fmt.Println(v3)
                                    }
                                }
                                _ = fi12
                                _ = v14
                            }
                        case -9223372036854775808:
                            var v28 int64 = (-(v6))
                            _ = v28
                        }
                        _ = len(v1)
                    }
                    _ = v8
                } else {
                    var v29 uint16 = (v2 << ((110) & 63))
                    _ = v29
                }
            case 127:
                v4 = ((v4 + (v4 + ("a" + v4))) + (("" + v4) + v4))
            default:
                var v30 uint = 4294967295
                _ = v30
            }
            var v31 float64 = 3.14159
            for fi32 := 0; fi32 < 6; fi32++ {
                if true {
                    switch (^(fi7)) {
                    case 255:
                        if false {
                            v6 += -9223372036854775808
                            var v33 int64 = v6
                            if len(v1) > 0 {
                                v1[0] = "a"
                            }
                            _ = v33
                        }
                    case 32767:
                        v34 := GenS4{F0: 255, F1: v4}
                        _ = v34
                    }
                }
                _ = fi32
            }
            _ = fi7
            _ = v31
        }
    }
    v35 := GenS4{F0: 185, F1: (((v4 + v4) + (v4 + (v4 + ""))) + ((("\n\t" + v4) + "\n\t") + v4))}
    v36 := GenS4{F0: 177, F1: v4}
    v6 = v6
    for fi37 := 0; fi37 < 4; fi37++ {
        v6 = -214
        _ = len(v5)
        var v38 bool = true
        _ = fi37
        _ = v38
    }
    switch -2147483648 {
    case -424:
        v6 -= (v6 + 560)
    default:
        v3 = -2.25
    }
    if (false && (false && false)) {
        v3 += 3.14159
        for fi39 := 0; fi39 < 18; fi39++ {
            switch fi39 {
            case 9223372036854775807:
                v4 = (("" + "\n\t") + v4)
            case 32768:
                v40 := GenS2{F0: v6}
                _ = v40
            }
            _ = fi39
        }
        switch 825 {
        case 127:
            for fi41 := 0; fi41 < 18; fi41++ {
                if (false && (!false)) {
                    v42 := GenS5{F0: int32(fi41), F1: 18446744073709551615}
                    _ = v42
                }
                var v43 uint32 = 2
                var v44 int8 = 42
                _ = fi41
                _ = v43
                _ = v44
            }
        case -9223372036854775808:
            for fi45 := 0; fi45 < 2; fi45++ {
                var v46 uint64 = 4294967295
                v5 = append(v5, (fi45 ^ fi45))
                _ = fi45
                _ = v46
            }
        case 42:
            if false {
                if ((1 < 84) || ((0 != 0) || false)) {
                    v6 = (v6 >> ((255) & 63))
                    v2 += v2
                    v2 += 116
                }
                fmt.Println(v3)
                fmt.Println(v2)
            }
        default:
            _ = len(v1)
        }
    }
    v47 := GenS2{F0: (v6 << ((0) & 63))}
    v1 = append(v1, v4)
    if (585 > 42) {
        if false {
            v2 = (v2 - (v2 * v2))
            switch -351 {
            case 2147483647:
                switch 496 {
                case -2147483648:
                    if ((1e10 < -2.25) && true) {
                        for fi48 := 0; fi48 < 11; fi48++ {
                            var v49 string = ""
                            var v50 string = v4
                            for fi51 := 0; fi51 < 8; fi51++ {
                                var v52 int8 = 0
                                _ = len(v1)
                                _ = fi51
                                _ = v52
                            }
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                        v6 += -779
                    } else {
                        for fi53 := 0; fi53 < 10; fi53++ {
                            for fi54 := 0; fi54 < 6; fi54++ {
                                if true {
                                    _ = len(v1)
                                    v55 := GenS1{F0: 770}
                                    _ = v55
                                } else {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if true {
                                        var v56 int64 = v6
                                        v3 = v3
                                        v2 = v2
                                        _ = v56
                                    } else {
                                        v3 = 0.1
                                    }
                                    if (127 == 42) {
                                        _ = len(v5)
                                    } else {
                                        var v57 uint8 = 165
                                        if true {
                                            v2 -= v2
                                        } else {
                                            v3 = float32(1e10)
                                            v58 := GenS3{F0: (v57 / (((v57 ^ v57)) | 1)), F1: "\n\t"}
                                            _ = v58
                                        }
                                        var v59 int32 = 2147483647
                                        _ = v57
                                        _ = v59
                                    }
                                }
                                var v60 bool = (!true)
                                var v61 uint8 = 128
                                _ = fi54
                                _ = v60
                                _ = v61
                            }
                            v2 -= v2
                            v62 := GenS5{F0: 2147483647, F1: 4294967295}
                            _ = fi53
                            _ = v62
                        }
                    }
                case -9223372036854775808:
                    if (false || (v3 != v3)) {
                        fmt.Println(v3, v6)
                    } else {
                        var v63 int8 = -128
                        _ = v63
                    }
                case 127:
                    for fi64 := 0; fi64 < 14; fi64++ {
                        var v65 int32 = 2147483647
                        v65 += 2147483647
                        v65 = (v65 / (((v65 | v65)) | 1))
                        _ = fi64
                        _ = v65
                    }
                default:
                    var v66 int16 = 42
                    _ = v66
                }
            case 42:
                fmt.Println(v3)
            }
        }
        v2 += v2
        v6 = (v6 >> ((uint(v2)) & 63))
    } else {
        v4 = v4
        v3 -= float32(0.0)
        var v67 bool = (true && true)
        _ = v67
    }
    v68 := GenS6{F0: 2, F1: 65535}
    if (!(!true)) {
        for fi69 := 0; fi69 < 2; fi69++ {
            var v70 uint = 990
            v70 = 137
            var v71 string = "a"
            _ = fi69
            _ = v70
            _ = v71
        }
        switch 2147483647 {
        case 42:
            var v72 int8 = 127
            _ = v72
        case 1:
            if (9223372036854775807 >= (v6 >> ((653) & 63))) {
                switch v4 {
                case "a":
                    if len(v5) > 0 {
                        v5[0] = 722
                    }
                case "hello":
                    for fi73 := 0; fi73 < 11; fi73++ {
                        var v74 uint = 685
                        switch int(v6) {
                        case 781:
                            var v75 uint = v74
                            _ = v75
                        }
                        _ = fi73
                        _ = v74
                    }
                case "foo bar":
                    var v76 int = -1
                    _ = v76
                }
                var v77 int32 = 2147483647
                v3 += (((0.0 - (v3 - v3)) / v3) - float32(1.5))
                _ = v77
            }
        }
    } else {
        for fi78 := 0; fi78 < 3; fi78++ {
            var v79 uint16 = (v2 & v2)
            v80 := GenS7{F0: 101, F1: 42}
            _ = fi78
            _ = v79
            _ = v80
        }
        var v81 uint64 = 264
        for fi82 := 0; fi82 < 3; fi82++ {
            if (606 != v6) {
                switch fi82 {
                case 386:
                    var v83 int16 = 1000
                    _ = v83
                }
                v84 := GenS6{F0: 1, F1: 843}
                _ = v84
            }
            _ = fi82
        }
        _ = v81
    }
    switch 127 {
    case 1:
        for fi85 := 0; fi85 < 7; fi85++ {
            switch fi85 {
            case 32767:
                if len(v5) > 0 {
                    v5[0] = fi85
                }
            case 1:
                v1 = append(v1, "hello")
            default:
                switch v4 {
                case "hello":
                    for fi86 := 0; fi86 < 2; fi86++ {
                        v87 := GenS4{F0: 34, F1: (v4 + (v4 + (v4 + ("\n\t" + "hello"))))}
                        switch (^(fi85)) {
                        case 1:
                            v88 := GenS6{F0: 149, F1: 2}
                            _ = v88
                        case 42:
                            var v89 int32 = -2147483648
                            _ = v89
                        case 0:
                            var v90 float32 = 1e10
                            _ = v90
                        }
                        _ = fi86
                        _ = v87
                    }
                }
            }
            for fi91 := 0; fi91 < 20; fi91++ {
                switch v4 {
                case "\n\t":
                    fmt.Println(v3)
                case "foo bar":
                    fmt.Println(v3, fi85)
                default:
                    v92 := GenS6{F0: 0, F1: 2147483647}
                    _ = v92
                }
                var v93 uint32 = 0
                _ = fi91
                _ = v93
            }
            _ = fi85
        }
    case 9223372036854775807:
        var v94 string = "hello"
        _ = v94
    case 2147483647:
        for fi95 := 0; fi95 < 11; fi95++ {
            v4 = "foo bar"
            _ = fi95
        }
    }
    var v96 int64 = 1
    for fi97 := 0; fi97 < 12; fi97++ {
        v4 = "\n\t"
        var v98 int16 = 1
        _ = fi97
        _ = v98
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v68
    _ = v96
}
