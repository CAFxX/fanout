package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 uint32; F1 int32 }
type GenS9 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS10 struct { F0 uint; F1 int32 }
type GenS11 struct { F0 uint; F1 uint8 }
type GenS12 struct { F0 uint64; F1 int; F2 uint }

func main() {
    var v1 string = "\n\t"
    var v2 int64 = 710
    var v3 uint16 = 42
    v4 := []uint32{703, 4294967295, 42, 1}
    fmt.Println(v3)
    for fi5 := 0; fi5 < 11; fi5++ {
        if true {
            v6 := GenS1{F0: 2147483647, F1: 2, F2: 964}
            _ = v6
        } else {
            fmt.Println(fi5)
            for fi7 := 0; fi7 < 20; fi7++ {
                fmt.Println(v1, v1)
                var v8 uint8 = 128
                _ = fi7
                _ = v8
            }
            var v9 float64 = float64((-2.25 * 3.14159))
            _ = v9
        }
        switch (fi5 & (fi5 + fi5)) {
        case -437:
            var v10 bool = (((-2.25 + 0.0) - ((1.5 + ((float32(-2.25) / 3.14159) * 1e10)) / 1.0)) >= 0.0)
            _ = v10
        default:
            for fi11 := 0; fi11 < 11; fi11++ {
                switch 4294967295 {
                case 2147483647:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 128:
                    v3 += (^(v3))
                }
                var v12 uint32 = 547
                _ = fi11
                _ = v12
            }
        }
        var v13 int32 = 1
        _ = fi5
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 int32 = 42
            switch "" {
            case "\n\t":
                if false {
                    v17 := GenS2{F0: false}
                    var v18 bool = false
                    v19 := GenS3{F0: v16, F1: 588, F2: -1}
                    _ = v17
                    _ = v18
                    _ = v19
                }
            case "a":
                v20 := GenS4{F0: 132, F1: 209, F2: (^(v3))}
                var v21 bool = (((1.5 / 1.0) - 3.14159) == 3.14159)
                _ = v20
                _ = v21
            case "foo bar":
                if true {
                    var v22 float32 = 0.0
                    fmt.Println(v1)
                    fmt.Println(fi15, v16, v2)
                    _ = v22
                } else {
                    v4 = append(v4, 1)
                }
            default:
                if (-1 > -32768) {
                    switch (fi15 + -339) {
                    case -1:
                        if (v2 > (v2 >> ((2) & 63))) {
                            var v23 int64 = v2
                            _ = v23
                        } else {
                            switch (fi14 ^ -295) {
                            case 32768:
                                if false {
                                    var v24 int8 = 127
                                    switch (fi14 % (((fi15 ^ (fi14 | fi14))) | 1)) {
                                    case 4294967295:
                                        v4 = append(v4, 729)
                                    case -1:
                                        if false {
                                            var v25 uint8 = 102
                                            _ = v25
                                        } else {
                                            var v26 int16 = -32768
                                            switch int(v26) {
                                            case 127:
                                                var v27 string = v1
                                                _ = v27
                                            case 4294967295:
                                                v4 = append(v4, 4294967295)
                                            }
                                            v3 += (v3 | 0)
                                            _ = v26
                                        }
                                        var v28 string = "hello"
                                        v28 = v28
                                        _ = v28
                                    }
                                    _ = v24
                                }
                            case 833:
                                var v29 float64 = -2.25
                                _ = v29
                            case 2147483647:
                                var v30 uint16 = v3
                                _ = v30
                            default:
                                v31 := GenS5{F0: 2}
                                _ = v31
                            }
                            var v32 uint16 = (v3 | v3)
                            _ = v32
                        }
                    case -550:
                        if true {
                            var v33 uint = 923
                            v4 = append(v4, 65535)
                            if (887 < 2147483647) {
                                var v34 uint32 = 2147483648
                                _ = v34
                            }
                            _ = v33
                        } else {
                            if len(v4) > 0 {
                                v4[0] = 4294967295
                            }
                        }
                    case 9223372036854775807:
                        v35 := GenS5{F0: 686}
                        _ = v35
                    default:
                        fmt.Println(v3, v2)
                        v1 = v1
                        v4 = append(v4, 212)
                    }
                } else {
                    switch fi14 {
                    case 9223372036854775807:
                        if true {
                            v3 -= 1
                        } else {
                            v1 = "\n\t"
                            var v36 int = (fi14 * fi15)
                            fmt.Println(v36)
                            _ = v36
                        }
                    case 127:
                        var v37 uint16 = v3
                        _ = v37
                    default:
                        var v38 int64 = (v2 << ((65535) & 63))
                        _ = v38
                    }
                }
            }
            _ = fi15
            _ = v16
        }
        var v39 bool = false
        var v40 bool = true
        _ = fi14
        _ = v39
        _ = v40
    }
    v41 := GenS6{F0: 1, F1: -1}
    switch 820 {
    case -1:
        v42 := GenS2{F0: (v1 > "hello")}
        _ = v42
    case 128:
        switch "a" {
        case "\n\t":
            for fi43 := 0; fi43 < 11; fi43++ {
                if true {
                    v3 = (^(v3))
                    if true {
                        v44 := GenS5{F0: 255}
                        _ = v44
                    } else {
                        _ = len(v4)
                    }
                } else {
                    v45 := GenS5{F0: 42}
                    _ = v45
                }
                var v46 uint = 0
                var v47 int16 = -1
                _ = fi43
                _ = v46
                _ = v47
            }
        case "foo bar":
            v2 = v2
        case "a":
            v48 := GenS5{F0: 683}
            _ = v48
        }
    }
    switch v1 {
    case "hello":
        fmt.Println(v1, v2)
    }
    if len(v4) > 0 {
        v4[0] = 2147483648
    }
    var v49 uint8 = 198
    fmt.Println(v2, v49)
    v1 = "hello"
    for fi50 := 0; fi50 < 15; fi50++ {
        v1 = "hello"
        for fi51 := 0; fi51 < 15; fi51++ {
            if true {
                fmt.Println(v3)
            }
            _ = fi51
        }
        var v52 int = fi50
        _ = fi50
        _ = v52
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi53 := 0; fi53 < 20; fi53++ {
        var v54 uint64 = 2
        for fi55 := 0; fi55 < 13; fi55++ {
            if (!((v49 - (v49 & 1)) > (^(v49)))) {
                v4 = append(v4, 4294967295)
                v1 = v1
            } else {
                _ = len(v4)
                v54 -= v54
            }
            v56 := GenS7{F0: 843, F1: "\n\t"}
            var v57 int64 = -1
            _ = fi55
            _ = v56
            _ = v57
        }
        for fi58 := 0; fi58 < 20; fi58++ {
            if true {
                if true {
                    v4 = append(v4, 0)
                    v3 -= 858
                }
                var v59 uint16 = v3
                v54 = v54
                _ = v59
            }
            v1 = v1
            var v60 int8 = 127
            _ = fi58
            _ = v60
        }
        _ = fi53
        _ = v54
    }
    v61 := GenS8{F0: 816, F1: 2147483647}
    v62 := GenS7{F0: 132, F1: v1}
    fmt.Println(v1, v49, v3)
    for fi63 := 0; fi63 < 4; fi63++ {
        switch "\n\t" {
        case "foo bar":
            for fi64 := 0; fi64 < 13; fi64++ {
                if (2 == 12) {
                    if false {
                        switch ("hello" + v1) {
                        case "foo bar":
                            if true {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v4 = append(v4, 147)
                                if len(v4) > 0 {
                                    v4[0] = 42
                                }
                            }
                        case "a":
                            switch -9223372036854775808 {
                            case -1:
                                if false {
                                    fmt.Println(v2, fi64)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if true {
                                        v2 = v2
                                        if (v1 >= "a") {
                                            var v65 uint = 1
                                            v65 -= 42
                                            _ = v65
                                        }
                                    } else {
                                        v66 := GenS9{F0: v49, F1: 0, F2: uint32(v49)}
                                        switch fi64 {
                                        case 1:
                                            if true {
                                                var v67 uint8 = (v49 << ((4294967295) & 63))
                                                v68 := GenS1{F0: 94, F1: 852, F2: v2}
                                                v67 = v49
                                                _ = v67
                                                _ = v68
                                            }
                                        case 4294967295:
                                            var v69 string = v1
                                            _ = v69
                                        case -9223372036854775808:
                                            var v70 float32 = (3.14159 * 0.0)
                                            _ = v70
                                        default:
                                            v2 = (v2 + (v2 * v2))
                                        }
                                        _ = v66
                                    }
                                } else {
                                    switch (^(fi63)) {
                                    case 128:
                                        if (false || true) {
                                            if len(v4) > 0 {
                                                v4[0] = 2147483648
                                            }
                                        } else {
                                            v71 := GenS4{F0: (v49 / ((v49) | 1)), F1: 255, F2: 2}
                                            v72 := GenS10{F0: 65535, F1: -2147483648}
                                            var v73 string = v1
                                            _ = v71
                                            _ = v72
                                            _ = v73
                                        }
                                    case 9223372036854775807:
                                        if ((fi64 + (fi63 / ((42) | 1))) != -849) {
                                            v74 := GenS4{F0: v49, F1: 331, F2: (v3 >> ((328) & 63))}
                                            v2 = -174
                                            v4 = append(v4, 4294967295)
                                            _ = v74
                                        }
                                    case -2147483648:
                                        var v75 string = v1
                                        _ = v75
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = 1
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = 0
                                    }
                                    v1 = v1
                                    switch fi63 {
                                    case 0:
                                        v2 += (-(v2))
                                    case 770:
                                        var v76 int8 = -128
                                        _ = v76
                                    case 168:
                                        v77 := GenS2{F0: (!true)}
                                        _ = v77
                                    default:
                                        var v78 uint8 = 255
                                        _ = v78
                                    }
                                }
                            case 724:
                                var v79 float64 = float64(((0.0 + 0.1) / (0.0 + 3.14159)))
                                _ = v79
                            case 128:
                                switch (fi64 % ((2147483647) | 1)) {
                                case -369:
                                    var v80 float32 = -2.25
                                    _ = v80
                                }
                            }
                        }
                        _ = len(v4)
                        v81 := GenS11{F0: 18446744073709551615, F1: 61}
                        _ = v81
                    } else {
                        v3 += (v3 / (((^(v3))) | 1))
                    }
                    if ((v49 + (v49 & (v49 >> ((65535) & 63)))) >= v49) {
                        v2 = (v2 | 42)
                        v2 = -9223372036854775808
                        var v82 uint = 42
                        _ = v82
                    }
                    switch v1 {
                    case "":
                        var v83 int32 = -618
                        _ = v83
                    }
                }
                v84 := GenS12{F0: 42, F1: fi64, F2: 1}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v85 int16 = 473
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi64
                _ = v84
                _ = v85
            }
        case "":
            switch fi63 {
            case 127:
                if (!false) {
                    fmt.Println(v1, fi63, v49)
                    v1 = ("a" + "foo bar")
                    var v86 float64 = (float64((0.1 + 3.14159)) / (3.14159 - ((-2.25 / 3.14159) * 3.14159)))
                    _ = v86
                }
            default:
                v49 += (^(v49))
            }
        default:
            if (2 != 18446744073709551615) {
                v87 := GenS1{F0: 42, F1: 255, F2: v2}
                var v88 float64 = 0.0
                var v89 uint8 = (v49 >> ((347) & 63))
                _ = v87
                _ = v88
                _ = v89
            }
        }
        v90 := GenS5{F0: 120}
        _ = fi63
        _ = v90
    }
    v3 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v49
    _ = v61
    _ = v62
}
