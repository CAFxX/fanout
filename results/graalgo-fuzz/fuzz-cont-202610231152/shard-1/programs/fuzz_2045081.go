package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint64; F2 int64 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS7 struct { F0 int64; F1 string; F2 int }
type GenS8 struct { F0 uint; F1 int; F2 int16 }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS12 struct { F0 bool; F1 bool }
func genFn0(p0 uint64) string {
    v4 := GenS1{F0: 255}
    var v5 int32 = -980
    v6 := GenS1{F0: uint16(p0)}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return "foo bar"
}

func main() {
    v1 := []int{0, 935}
    v2 := []uint32{98, 144, 2147483647, 2147483647}
    var v3 uint16 = 2
    if true {
        switch 4294967295 {
        case 9223372036854775807:
            v7 := GenS2{F0: 42, F1: 0, F2: 9223372036854775807}
            _ = v7
        case 42:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 2147483647:
            var v8 uint32 = 2147483648
            _ = v8
        default:
            for fi9 := 0; fi9 < 2; fi9++ {
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 int32 = 0
                    v12 := GenS1{F0: 2}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                if (-1 >= 1) {
                    switch (-(fi9)) {
                    case 4294967295:
                        v13 := GenS2{F0: 1, F1: 1, F2: 9223372036854775807}
                        _ = v13
                    case 9223372036854775807:
                        var v14 int64 = -9223372036854775808
                        _ = v14
                    case 572:
                        for fi15 := 0; fi15 < 16; fi15++ {
                            v3 -= 255
                            _ = fi15
                        }
                    }
                    v3 = (v3 & (^(v3)))
                }
                v16 := GenS3{F0: -1, F1: (fi9 & fi9)}
                _ = fi9
                _ = v16
            }
        }
        switch -2147483648 {
        case 9223372036854775807:
            var v17 uint16 = v3
            _ = v17
        default:
            switch 1000 {
            case 127:
                v18 := GenS2{F0: -1, F1: 0, F2: -1}
                _ = v18
            case -9223372036854775808:
                for fi19 := 0; fi19 < 8; fi19++ {
                    if true {
                        v3 = v3
                    } else {
                        fmt.Println(fi19, fi19)
                        for fi20 := 0; fi20 < 12; fi20++ {
                            fmt.Println(v3)
                            _ = fi20
                        }
                        var v21 uint16 = v3
                        _ = v21
                    }
                    _ = fi19
                }
            case 128:
                var v22 uint = 42
                _ = v22
            }
        }
    } else {
        v23 := GenS2{F0: 32767, F1: 309, F2: 1}
        v24 := GenS2{F0: 1000, F1: 122, F2: 1}
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 uint = 65535
        v3 -= v3
        _ = fi25
        _ = v26
    }
    var v27 bool = ((42 < v3) && (!(245 < 42)))
    _ = len(v1)
    v27 = (-1 == 32767)
    v28 := GenS2{F0: -32768, F1: 0, F2: -566}
    for fi29 := 0; fi29 < 17; fi29++ {
        _ = len(v1)
        _ = fi29
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        v3 += v3
        v27 = (-26 < 1)
        v3 += v3
        _ = fi30
    }
    if (!v27) {
        for fi31 := 0; fi31 < 12; fi31++ {
            v3 = v3
            v27 = v27
            v3 -= 693
            _ = fi31
        }
        switch "a" {
        case "hello":
            var v32 uint32 = 4294967295
            _ = v32
        case "\n\t":
            switch 32767 {
            case 32768:
                v2 = append(v2, 42)
            case 841:
                v1 = append(v1, -1)
            default:
                v33 := GenS3{F0: -32768, F1: -2147483648}
                _ = v33
            }
        default:
            for fi34 := 0; fi34 < 11; fi34++ {
                v35 := GenS3{F0: 42, F1: -2147483648}
                for fi36 := 0; fi36 < 15; fi36++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(fi36)
                    switch (fi36 >> ((840) & 63)) {
                    case -9223372036854775808:
                        v3 += v3
                    case 127:
                        var v37 int8 = 127
                        v38 := GenS1{F0: (v3 ^ (^(v3)))}
                        v27 = (0.0 == (((((float64(-2.25) + float64(-2.25)) - -2.25) + 0.0) + 3.14159) * 0.0))
                        _ = v37
                        _ = v38
                    }
                    _ = fi36
                }
                _ = fi34
                _ = v35
            }
        }
    }
    v39 := GenS4{F0: 0, F1: uint64(v3)}
    fmt.Println(v3, v27)
    for fi40 := 0; fi40 < 13; fi40++ {
        if (!((0.1 / (1e10 - -2.25)) < 3.14159)) {
            var v41 int64 = 0
            for fi42 := 0; fi42 < 9; fi42++ {
                fmt.Println(v27, fi40)
                if v27 {
                    if (v27 && (-63 <= 42)) {
                        v43 := GenS4{F0: 127, F1: 202}
                        fmt.Println(fi40, v27, v3)
                        var v44 uint32 = 1
                        _ = v43
                        _ = v44
                    }
                    if v27 {
                        if v27 {
                            v27 = (!true)
                            if v27 {
                                var v45 uint32 = 2
                                fmt.Println(v45, v3)
                                var v46 uint32 = v45
                                switch -1 {
                                case 32768:
                                    v47 := GenS3{F0: 1000, F1: 32768}
                                    _ = v47
                                case -274:
                                    switch "foo bar" {
                                    case "a":
                                        v2 = append(v2, v45)
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                default:
                                    if false {
                                        switch "hello" {
                                        case "\n\t":
                                            var v48 float64 = -2.25
                                            _ = v48
                                        case "":
                                            var v49 uint16 = (v3 * v3)
                                            _ = v49
                                        case "a":
                                            if len(v2) > 0 {
                                                v2[0] = (v46 | (v46 + v46))
                                            }
                                        default:
                                            v46 += (v45 >> ((595) & 63))
                                        }
                                        var v50 uint64 = 18446744073709551615
                                        v51 := GenS5{F0: 0, F1: (^(v45))}
                                        _ = v50
                                        _ = v51
                                    } else {
                                        if true {
                                            var v52 uint = 1
                                            var v53 int = (^(fi42))
                                            v54 := GenS6{F0: 896, F1: 65535, F2: 4294967295}
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        }
                                    }
                                }
                                v55 := GenS5{F0: 4294967295, F1: (v46 * 2147483647)}
                                v46 += v46
                                v56 := GenS7{F0: 718, F1: ((("hello" + "hello") + "\n\t") + "a"), F2: (^(fi42))}
                                _ = v45
                                _ = v46
                                _ = v55
                                _ = v56
                            }
                            var v57 uint8 = 255
                            _ = v57
                        } else {
                            v3 = v3
                            v58 := GenS2{F0: -32768, F1: uint64(v3), F2: v41}
                            var v59 int = int(v3)
                            _ = v58
                            _ = v59
                        }
                        var v60 uint64 = 2
                        _ = v60
                    } else {
                        v27 = false
                    }
                    v41 += v41
                }
                v41 += -85
                _ = fi42
            }
            _ = v41
        } else {
            var v61 int64 = 42
            _ = v61
        }
        if v27 {
            switch "" {
            case "a":
                switch "\n\t" {
                case "hello":
                    var v62 bool = (1 != 1)
                    _ = v62
                case "\n\t":
                    fmt.Println(v27, v3)
                }
            case "":
                switch (fi40 ^ (fi40 + fi40)) {
                case 9223372036854775807:
                    v63 := GenS8{F0: 18446744073709551615, F1: fi40, F2: 32767}
                    _ = v63
                case -92:
                    for fi64 := 0; fi64 < 11; fi64++ {
                        fmt.Println(v3)
                        _ = fi64
                    }
                case -393:
                    v27 = v27
                }
            }
            v3 = (^(v3))
            switch ("" + "hello") {
            case "hello":
                var v65 string = "foo bar"
                _ = v65
            case "":
                for fi66 := 0; fi66 < 11; fi66++ {
                    v67 := GenS9{F0: 0, F1: 1}
                    _ = fi66
                    _ = v67
                }
            }
        }
        switch fi40 {
        case 0:
            v3 = 94
        case -102:
            switch "" {
            case "foo bar":
                v68 := GenS2{F0: -872, F1: 4294967295, F2: 9223372036854775807}
                _ = v68
            default:
                fmt.Println(v3, v3)
            }
        default:
            v27 = (!((v3 ^ v3) < (v3 % ((v3) | 1))))
        }
        _ = fi40
    }
    for fi69 := 0; fi69 < 8; fi69++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v70 uint32 = 2
        _ = fi69
        _ = v70
    }
    switch (("\n\t" + ((("foo bar" + ("\n\t" + "foo bar")) + "a") + "\n\t")) + "\n\t") {
    case "a":
        var v71 bool = ((0.1 >= (3.14159 - 3.14159)) && (v27 || (1 >= 151)))
        _ = v71
    case "foo bar":
        for fi72 := 0; fi72 < 12; fi72++ {
            fmt.Println(v3)
            _ = fi72
        }
    case "hello":
        if (1e10 > -2.25) {
            var v73 uint32 = 799
            switch "\n\t" {
            case "a":
                v27 = (v27 || (!(!((v73 << ((18446744073709551615) & 63)) <= v73))))
            }
            fmt.Println(v3, v27)
            _ = v73
        } else {
            var v74 uint64 = 2
            fmt.Println(v3, v74, v27)
            _ = v74
        }
    }
    switch 9223372036854775807 {
    case 1:
        if true {
            v27 = true
        } else {
            v3 -= 210
            v75 := GenS10{F0: 2}
            if (!true) {
                var v76 bool = true
                v27 = v76
                v77 := GenS11{F0: 0, F1: 1, F2: 2147483647}
                _ = v76
                _ = v77
            } else {
                for fi78 := 0; fi78 < 3; fi78++ {
                    if len(v1) > 0 {
                        v1[0] = (fi78 + fi78)
                    }
                    var v79 int8 = -1
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi78
                    _ = v79
                }
                fmt.Println(v27, v3)
                for fi80 := 0; fi80 < 8; fi80++ {
                    for fi81 := 0; fi81 < 15; fi81++ {
                        var v82 int = (fi80 % ((fi81) | 1))
                        switch -230 {
                        case -9223372036854775808:
                            switch "a" {
                            case "a":
                                v3 = (v3 / (((^(v3))) | 1))
                            case "foo bar":
                                if v27 {
                                    fmt.Println(v27, fi80, v82)
                                } else {
                                    fmt.Println(v82, v3, v82)
                                }
                            case "":
                                v27 = true
                                v82 -= 9223372036854775807
                            }
                        default:
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                var v83 int32 = 2147483647
                                _ = v83
                            case 0:
                                v84 := GenS5{F0: 1, F1: 215}
                                _ = v84
                            case 32767:
                                _ = len(v2)
                            default:
                                v1 = append(v1, v82)
                            }
                        }
                        switch "\n\t" {
                        case "":
                            var v85 uint8 = 255
                            _ = v85
                        case "hello":
                            fmt.Println(fi80)
                        case "a":
                            if v27 {
                                _ = len(v2)
                                v2 = append(v2, 765)
                                _ = len(v2)
                            }
                        default:
                            if v27 {
                                var v86 int = -164
                                v86 = (fi80 | 981)
                                _ = v86
                            }
                            v82 -= 4294967295
                            switch int(v3) {
                            case -2147483648:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case 128:
                                var v87 uint16 = 255
                                _ = v87
                            default:
                                v82 = fi80
                            }
                            v88 := GenS12{F0: false, F1: (1 != -9223372036854775808)}
                            _ = v88
                        }
                        _ = fi81
                        _ = v82
                    }
                    _ = fi80
                }
            }
            _ = v75
        }
    case 127:
        for fi89 := 0; fi89 < 7; fi89++ {
            var v90 string = "foo bar"
            _ = fi89
            _ = v90
        }
    case 0:
        v3 += 255
    }
    var v91 uint = 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v39
    _ = v91
}
