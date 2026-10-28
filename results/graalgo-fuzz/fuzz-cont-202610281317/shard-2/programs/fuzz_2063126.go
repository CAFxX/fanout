package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int16; F1 uint32; F2 int32 }
type GenS6 struct { F0 string }

func main() {
    v1 := []uint32{4294967295}
    v2 := []int{9223372036854775807, 32768, 127, -423}
    var v3 int64 = 0
    v4 := []uint32{636, 2}
    var v5 int16 = 42
    _ = len(v2)
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint8 = 0
        _ = fi6
        _ = v7
    }
    v5 = v5
    if true {
        for fi8 := 0; fi8 < 2; fi8++ {
            if ((!((false && false) || ((1 < (v5 % ((v5) | 1))) && ((((true && ("" < "\n\t")) && true) && (false || (198 == 2))) || false)))) || true) {
                if (false && true) {
                    for fi9 := 0; fi9 < 20; fi9++ {
                        var v10 uint64 = 42
                        _ = fi9
                        _ = v10
                    }
                    var v11 int8 = 127
                    _ = v11
                } else {
                    var v12 string = "a"
                    _ = v12
                }
            }
            var v13 uint = 65535
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi8
            _ = v13
        }
    }
    _ = len(v2)
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS1{F0: 2, F1: 0}
        _ = fi14
        _ = v15
    }
    v3 = 42
    v5 -= v5
    fmt.Println(v5, v5)
    switch "\n\t" {
    case "\n\t":
        if (!(((!(127 < 1)) || true) && (false && true))) {
            v5 = -32768
            if false {
                if len(v4) > 0 {
                    v4[0] = 2147483648
                }
            }
            switch 32768 {
            case 2147483647:
                if true {
                    v5 -= v5
                    var v16 uint32 = 450
                    switch -204 {
                    case -9223372036854775808:
                        v16 -= (v16 | v16)
                    }
                    _ = v16
                }
            case 9223372036854775807:
                fmt.Println(v3, v5)
            }
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v5)
    if true {
        var v17 uint32 = 65535
        switch ("a" + ("a" + "hello")) {
        case "\n\t":
            v17 += (v17 * (v17 & v17))
        case "hello":
            if len(v2) > 0 {
                v2[0] = -831
            }
        }
        v1 = append(v1, 643)
        _ = v17
    } else {
        v18 := GenS2{F0: "foo bar", F1: 65535}
        if (2 < 0) {
            for fi19 := 0; fi19 < 5; fi19++ {
                for fi20 := 0; fi20 < 19; fi20++ {
                    v21 := GenS2{F0: "hello", F1: 255}
                    _ = fi20
                    _ = v21
                }
                if true {
                    if true {
                        switch "hello" {
                        case "a":
                            var v22 uint = 13
                            _ = v22
                        case "hello":
                            var v23 uint8 = 232
                            _ = v23
                        }
                    } else {
                        var v24 uint32 = 13
                        v24 = (^(v24))
                        _ = v24
                    }
                    v3 = 9223372036854775807
                }
                var v25 int16 = (v5 & v5)
                _ = fi19
                _ = v25
            }
            v26 := GenS3{F0: (110 >= 127), F1: 100, F2: 42}
            _ = v26
        } else {
            v5 = v5
            for fi27 := 0; fi27 < 6; fi27++ {
                for fi28 := 0; fi28 < 12; fi28++ {
                    v5 = -1
                    _ = fi28
                }
                _ = fi27
            }
            var v29 string = "a"
            _ = v29
        }
        _ = v18
    }
    switch 32768 {
    case 32767:
        for fi30 := 0; fi30 < 4; fi30++ {
            if (2147483647 > 2147483647) {
                if (100 < -111) {
                    for fi31 := 0; fi31 < 10; fi31++ {
                        v5 = v5
                        _ = fi31
                    }
                } else {
                    switch fi30 {
                    case -241:
                        var v32 int32 = 0
                        _ = v32
                    case -281:
                        v4 = append(v4, 614)
                    case -260:
                        switch (^(fi30)) {
                        case 29:
                            _ = len(v2)
                        case 4294967295:
                            for fi33 := 0; fi33 < 3; fi33++ {
                                v34 := GenS1{F0: 1, F1: v3}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v35 := GenS4{F0: 134, F1: 149}
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                        case 20:
                            v1 = append(v1, 2147483647)
                        default:
                            v3 += (^(v3))
                        }
                    default:
                        v4 = append(v4, 65535)
                    }
                    var v36 int8 = -1
                    var v37 bool = (2 == 663)
                    _ = v36
                    _ = v37
                }
                for fi38 := 0; fi38 < 5; fi38++ {
                    fmt.Println(v3, fi38, fi38)
                    _ = fi38
                }
            } else {
                v1 = append(v1, 2147483647)
            }
            v5 = v5
            if false {
                if len(v2) > 0 {
                    v2[0] = (fi30 - (fi30 / ((fi30) | 1)))
                }
                var v39 int8 = 17
                for fi40 := 0; fi40 < 13; fi40++ {
                    v3 = v3
                    fmt.Println(fi30)
                    v5 -= v5
                    _ = fi40
                }
                _ = v39
            } else {
                v4 = append(v4, 694)
            }
            _ = fi30
        }
    case -1:
        v41 := GenS4{F0: 0, F1: 15}
        _ = v41
    case -9223372036854775808:
        v4 = append(v4, 936)
    }
    switch 1 {
    case -865:
        switch 9223372036854775807 {
        case 587:
            v42 := GenS1{F0: 42, F1: v3}
            _ = v42
        case 2147483647:
            var v43 float64 = (0.1 + 3.14159)
            _ = v43
        default:
            if true {
                if false {
                    if (v3 <= (v3 / (((v3 - v3)) | 1))) {
                        v44 := GenS1{F0: 593, F1: 9223372036854775807}
                        if (592 == 18446744073709551615) {
                            switch 32768 {
                            case 0:
                                var v45 int64 = (v3 | -783)
                                _ = v45
                            default:
                                _ = len(v1)
                            }
                        }
                        _ = v44
                    }
                    if (670 > 0) {
                        v46 := GenS3{F0: (928 > 18446744073709551615), F1: 102, F2: 32767}
                        v47 := GenS1{F0: 35, F1: (^(v3))}
                        _ = v46
                        _ = v47
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            switch 569 {
                            case 698:
                                fmt.Println(v5)
                            case 32768:
                                v3 = v3
                            case 4294967295:
                                if false {
                                    fmt.Println(v3, v3)
                                    if true {
                                        if len(v4) > 0 {
                                            v4[0] = 1
                                        }
                                        if (!((false && true) && (0 <= 667))) {
                                            fmt.Println(v5, v5)
                                            var v48 int = -9223372036854775808
                                            _ = v48
                                        }
                                        v49 := GenS5{F0: (v5 - (v5 % ((v5) | 1))), F1: 2147483648, F2: 42}
                                        _ = v49
                                    }
                                    for fi50 := 0; fi50 < 5; fi50++ {
                                        v2 = append(v2, fi50)
                                        _ = fi50
                                    }
                                } else {
                                    if (3.14159 >= 1.5) {
                                        var v51 int16 = v5
                                        if (true && (1 >= v51)) {
                                            var v52 int16 = (v51 + (v51 - (-(v51))))
                                            v53 := GenS3{F0: (!true), F1: 127, F2: 42}
                                            _ = v52
                                            _ = v53
                                        }
                                        if true {
                                            fmt.Println(v5)
                                        }
                                        _ = v51
                                    } else {
                                        var v54 string = "\n\t"
                                        fmt.Println(v54, v3, v5)
                                        _ = v54
                                    }
                                    switch "hello" {
                                    case "":
                                        var v55 uint8 = 213
                                        _ = v55
                                    default:
                                        switch "foo bar" {
                                        case "foo bar":
                                            if true {
                                                var v56 int8 = 100
                                                switch "a" {
                                                case "hello":
                                                    var v57 uint16 = 65535
                                                    _ = v57
                                                case "":
                                                    for fi58 := 0; fi58 < 12; fi58++ {
                                                        var v59 string = "foo bar"
                                                        v56 = (^(v56))
                                                        fmt.Println(v5)
                                                        _ = fi58
                                                        _ = v59
                                                    }
                                                case "a":
                                                    v60 := GenS3{F0: (!(!false)), F1: v56, F2: 385}
                                                    _ = v60
                                                default:
                                                    v3 = -1
                                                }
                                                if (-1 > 32768) {
                                                    v61 := GenS6{F0: "a"}
                                                    _ = v61
                                                }
                                                _ = v56
                                            } else {
                                                v5 = (^(v5))
                                                fmt.Println(v5, v3)
                                                var v62 int16 = v5
                                                _ = v62
                                            }
                                        case "\n\t":
                                            for fi63 := 0; fi63 < 8; fi63++ {
                                                switch fi63 {
                                                case -476:
                                                    fmt.Println(v5, v3)
                                                case 127:
                                                    var v64 float32 = (1.5 - float32(0.1))
                                                    _ = v64
                                                case -1:
                                                    if false {
                                                        v65 := GenS6{F0: "foo bar"}
                                                        for fi66 := 0; fi66 < 9; fi66++ {
                                                            if (((((float64(float32(-2.25)) + ((-2.25 - -2.25) / 1.0)) * (((0.1 - (0.0 / ((float64(-2.25) / 1.0) * (3.14159 + -2.25)))) / (0.1 / (1.5 * float64(3.14159)))) - 3.14159)) * 0.1) > 1.5) && true) {
                                                                v5 = 42
                                                            }
                                                            _ = fi66
                                                        }
                                                        _ = v65
                                                    }
                                                default:
                                                    v3 = (^(v3))
                                                }
                                                for fi67 := 0; fi67 < 17; fi67++ {
                                                    v68 := GenS2{F0: "hello", F1: 42}
                                                    _ = fi67
                                                    _ = v68
                                                }
                                                for fi69 := 0; fi69 < 13; fi69++ {
                                                    v3 = (v3 - v3)
                                                    fmt.Println(fi63)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = fi69
                                                }
                                                _ = fi63
                                            }
                                        case "a":
                                            for fi70 := 0; fi70 < 14; fi70++ {
                                                _ = len(v4)
                                                v3 -= 315
                                                v71 := GenS5{F0: (v5 * v5), F1: 319, F2: 71}
                                                _ = fi70
                                                _ = v71
                                            }
                                        }
                                    }
                                    v5 = v5
                                }
                            }
                        default:
                            fmt.Println(v5)
                        }
                        var v72 uint = 65535
                        if false {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v73 uint32 = 42
                            if len(v2) > 0 {
                                v2[0] = 127
                            }
                            _ = v73
                        } else {
                            if (false || false) {
                                fmt.Println(v5, v3)
                            } else {
                                if ((!true) || (!(0.0 == 1.5))) {
                                    var v74 int32 = 2147483647
                                    var v75 float64 = -2.25
                                    _ = v74
                                    _ = v75
                                }
                                var v76 uint32 = 85
                                v72 = uint(v76)
                                _ = v76
                            }
                        }
                        _ = v72
                    }
                } else {
                    fmt.Println(v3)
                    v2 = append(v2, 144)
                }
                for fi77 := 0; fi77 < 17; fi77++ {
                    v3 = (v3 & -9223372036854775808)
                    var v78 int32 = int32(v3)
                    _ = fi77
                    _ = v78
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
