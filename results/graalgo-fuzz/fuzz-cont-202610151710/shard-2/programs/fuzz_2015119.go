package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 string; F1 bool; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int64; F1 int16; F2 uint16 }
type GenS6 struct { F0 int16; F1 int }
func genFn0(p0 int32, p1 uint, p2 uint16) uint8 {
    var v7 float32 = 0.1
    p2 = 42
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 uint64) int64 {
    v10 := []int{}
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint32 = 422
        if len(v10) > 0 {
            v10[0] = 128
        }
        v12 += (v12 - (v12 << (((p0 >> (((p0 / (((^(p0))) | 1))) & 63))) & 63)))
        _ = fi11
        _ = v12
    }
    var v13 uint16 = 255
    var v14 uint16 = v13
    _ = p0
    _ = p1
    _ = v10
    _ = v13
    _ = v14
    return 433
}

func main() {
    var v1 uint8 = 128
    var v2 int = 32768
    v3 := make(map[string]bool)
    v4 := make(map[string]string)
    var v5 int32 = -1
    v6 := []uint32{621}
    var v15 float64 = -2.25
    fmt.Println(v15, v15, v1)
    fmt.Println(v5, v2)
    v16 := GenS1{F0: -32768, F1: 61, F2: 42}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (-1 <= 32767) {
        v1 = v1
        switch ("hello" + "hello") {
        case "hello":
            var v17 float64 = v15
            _ = v17
        case "a":
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 int8 = 111
                _ = fi18
                _ = v19
            }
        case "foo bar":
            switch (v2 * (v2 - (v2 >> ((uint(v1)) & 63)))) {
            case 32767:
                fmt.Println(v15, v5, v15)
            case -335:
                switch (v2 / ((v2) | 1)) {
                case 931:
                    if (!(834 <= 0)) {
                        for fi20 := 0; fi20 < 7; fi20++ {
                            v21 := GenS2{F0: "hello", F1: (false && (("\n\t" != ("foo bar" + ("hello" + "hello"))) && (-550 < 0))), F2: 0}
                            var v22 int64 = 653
                            switch (("" + "") + "a") {
                            case "foo bar":
                                var v23 float32 = 3.14159
                                _ = v23
                            case "hello":
                                delete(v3, "a")
                            default:
                                v3["hello"] = (-97 == 1)
                            }
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        v3["foo bar"] = true
                        switch "a" {
                        case "\n\t":
                            for fi24 := 0; fi24 < 8; fi24++ {
                                if true {
                                    switch (-(fi24)) {
                                    case 2147483647:
                                        var v25 string = "a"
                                        _ = v25
                                    default:
                                        switch -462 {
                                        case -2147483648:
                                            v3["a"] = false
                                        case 128:
                                            var v26 bool = true
                                            _ = v26
                                        }
                                    }
                                    fmt.Println(v1, v5, fi24)
                                    var v27 uint64 = 2
                                    _ = v27
                                }
                                var v28 int8 = 127
                                var v29 int32 = v5
                                _ = fi24
                                _ = v28
                                _ = v29
                            }
                        case "":
                            v30 := GenS3{F0: 65535}
                            _ = v30
                        }
                    }
                case 9223372036854775807:
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v15 -= ((v15 + float64((0.1 + (1.5 / (float32(v15) / float32(-2.25)))))) / 1.5)
                        switch (^(v2)) {
                        case 4294967295:
                            if (100 > -1) {
                                if (65535 > 708) {
                                    fmt.Println(fi31, v2, v2)
                                    delete(v3, "a")
                                }
                                if (1 < 18446744073709551615) {
                                    switch "a" {
                                    case "":
                                        switch "a" {
                                        case "\n\t":
                                            fmt.Println(v15)
                                        case "hello":
                                            for fi32 := 0; fi32 < 18; fi32++ {
                                                v33 := GenS3{F0: 0}
                                                if true {
                                                    v34 := GenS2{F0: ("hello" + ""), F1: (false || (false || (9223372036854775807 > -1))), F2: 0}
                                                    var v35 uint16 = 0
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    v5 = (^(v5))
                                                    var v36 uint64 = 662
                                                    fmt.Println(fi31)
                                                    _ = v36
                                                }
                                                _ = fi32
                                                _ = v33
                                            }
                                        default:
                                            for fi37 := 0; fi37 < 12; fi37++ {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                _ = len(v4)
                                                _ = fi37
                                            }
                                        }
                                    case "foo bar":
                                        var v38 uint8 = 2
                                        _ = v38
                                    default:
                                        v5 -= (^(v5))
                                    }
                                    v2 -= (v2 & v2)
                                    if (1.5 != ((((1.5 - float32(v15)) + 0.1) / 1e10) * 3.14159)) {
                                        fmt.Println(v2, v2, v2)
                                    } else {
                                        for fi39 := 0; fi39 < 12; fi39++ {
                                            var v40 float64 = float64(((float32(v15) * (((-2.25 / float32(v15)) - 0.0) * -2.25)) - -2.25))
                                            _ = fi39
                                            _ = v40
                                        }
                                        var v41 uint8 = v1
                                        _ = v41
                                    }
                                } else {
                                    v3["\n\t"] = ("\n\t" > ((("hello" + ("hello" + "\n\t")) + "\n\t") + ("a" + "a")))
                                    _ = len(v6)
                                }
                                v1 -= (v1 | (^(v1)))
                            }
                        case 1:
                            for fi42 := 0; fi42 < 15; fi42++ {
                                if (127 <= 1) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v43 string = "hello"
                                        _ = v43
                                    }
                                    fmt.Println(v2, v5, fi31)
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v5 = (v5 << ((351) & 63))
                                }
                                v44 := GenS1{F0: -155, F1: -1, F2: 174}
                                _ = fi42
                                _ = v44
                            }
                        case 127:
                            for fi45 := 0; fi45 < 6; fi45++ {
                                if false {
                                    var v46 float32 = 3.14159
                                    _ = v46
                                }
                                v47 := GenS2{F0: ("\n\t" + ""), F1: true, F2: -1}
                                delete(v3, "hello")
                                _ = fi45
                                _ = v47
                            }
                        default:
                            v48 := GenS4{F0: 2147483648, F1: v2}
                            _ = v48
                        }
                        var v49 uint16 = 65535
                        _ = fi31
                        _ = v49
                    }
                }
            }
        }
    } else {
        switch 32768 {
        case 127:
            fmt.Println(v5, v2)
        default:
            var v50 int32 = 1
            _ = v50
        }
    }
    var v51 bool = true
    v15 = 0.1
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if true {
        var v52 int8 = 42
        fmt.Println(v1, v1, v52)
        _ = v52
    } else {
        switch v2 {
        case 632:
            v53 := GenS5{F0: -742, F1: -32768, F2: 255}
            _ = v53
        case 1:
            switch ("a" + "\n\t") {
            case "hello":
                v54 := GenS6{F0: -32768, F1: (^(v2))}
                _ = v54
            case "foo bar":
                var v55 float64 = (v15 + 0.0)
                _ = v55
            case "\n\t":
                v1 = (v1 % (((v1 >> ((4294967295) & 63))) | 1))
            default:
                switch ("\n\t" + "\n\t") {
                case "":
                    delete(v3, "hello")
                case "hello":
                    var v56 uint = 1
                    _ = v56
                default:
                    delete(v3, "foo bar")
                }
            }
        }
        var v57 int32 = v5
        var v58 string = "a"
        _ = v57
        _ = v58
    }
    _ = len(v3)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if v51 {
        var v59 int64 = genFn1(255, 2)
        switch v2 {
        case 42:
            for fi60 := 0; fi60 < 20; fi60++ {
                var v61 float32 = 3.14159
                _ = fi60
                _ = v61
            }
        }
        _ = v59
    }
    v5 = v5
    if (42 <= 1) {
        if (v51 && (!(329 != 145))) {
            v5 = 42
            switch 0 {
            case 2147483647:
                v5 = v5
            case 127:
                var v62 uint8 = genFn0((v5 + (-(v5))), 0, 681)
                _ = v62
            case 0:
                v5 = 1
            default:
                if ((float32((float64(0.0) + 0.1)) * 3.14159) != 1.5) {
                    v51 = false
                }
            }
            for fi63 := 0; fi63 < 2; fi63++ {
                if true {
                    switch -2147483648 {
                    case -528:
                        switch v2 {
                        case 128:
                            var v64 uint32 = 596
                            _ = v64
                        case -2147483648:
                            v15 += (v15 - (v15 + v15))
                        case 288:
                            switch fi63 {
                            case -228:
                                v65 := v4["\n\t"]
                                _ = v65
                            }
                        default:
                            for fi66 := 0; fi66 < 6; fi66++ {
                                v1 = genFn0(-2147483648, 0, 65535)
                                v67 := GenS4{F0: 0, F1: (-(fi66))}
                                _ = fi66
                                _ = v67
                            }
                        }
                    case 32768:
                        v68 := GenS2{F0: (("" + ("foo bar" + ("a" + "foo bar"))) + "hello"), F1: (v15 > (v15 * v15)), F2: -1}
                        _ = v68
                    default:
                        var v69 int64 = 1
                        _ = v69
                    }
                    v15 = 0.0
                }
                var v70 string = "hello"
                var v71 int64 = -270
                _ = fi63
                _ = v70
                _ = v71
            }
        } else {
            if (!(!(((v5 ^ 2147483647) < (v5 / ((v5) | 1))) && v51))) {
                for fi72 := 0; fi72 < 2; fi72++ {
                    delete(v4, ((("foo bar" + ("\n\t" + "\n\t")) + (("" + "a") + "foo bar")) + ("foo bar" + "hello")))
                    if v51 {
                        v73 := GenS4{F0: 4294967295, F1: (-(fi72))}
                        fmt.Println(v2)
                        switch "foo bar" {
                        case "foo bar":
                            if (!false) {
                                v5 += (-(v5))
                                var v74 bool = ((-(v2)) == (fi72 >> ((360) & 63)))
                                fmt.Println(v5)
                                _ = v74
                            }
                        case "hello":
                            var v75 uint64 = 2
                            _ = v75
                        case "a":
                            var v76 int64 = 1
                            _ = v76
                        default:
                            v51 = (((^(v1)) > (v1 | (v1 >> ((843) & 63)))) || (51 != 4294967295))
                        }
                        _ = v73
                    } else {
                        if true {
                            var v77 float64 = (v15 * 1.5)
                            fmt.Println(v1, v51)
                            if false {
                                var v78 float32 = -2.25
                                for fi79 := 0; fi79 < 9; fi79++ {
                                    var v80 uint16 = 255
                                    _ = fi79
                                    _ = v80
                                }
                                _ = v78
                            } else {
                                v6 = append(v6, 4294967295)
                            }
                            _ = v77
                        } else {
                            v2 -= (fi72 * int(v1))
                            for fi81 := 0; fi81 < 7; fi81++ {
                                if (!v51) {
                                    if (2 >= 0) {
                                        if true {
                                            v2 = (fi81 << ((740) & 63))
                                            v5 -= v5
                                            fmt.Println(v2)
                                            var v82 int8 = -128
                                            _ = v82
                                        }
                                        if (false && ((^(v1)) > 1)) {
                                            delete(v4, "hello")
                                        }
                                        var v83 float32 = 0.1
                                        v1 = (v1 - (v1 / ((genFn0(-2147483648, 18446744073709551615, 1)) | 1)))
                                        _ = v83
                                    }
                                    switch "a" {
                                    case "":
                                        v5 -= (v5 - v5)
                                    case "\n\t":
                                        fmt.Println(fi81, fi81)
                                    default:
                                        v2 = 0
                                        fmt.Println(fi72, v5, v2)
                                    }
                                }
                                _ = fi81
                            }
                        }
                        v84 := GenS1{F0: -1, F1: 1, F2: 1}
                        _ = v84
                    }
                    _ = fi72
                }
                var v85 float64 = 0.1
                if v51 {
                    v1 = (v1 & (v1 * (^(v1))))
                    v5 = 1
                    for fi86 := 0; fi86 < 8; fi86++ {
                        v2 = (^(fi86))
                        _ = fi86
                    }
                } else {
                    v51 = ((v1 * 0) >= genFn0((v5 / ((v5) | 1)), 592, 42))
                }
                _ = v85
            } else {
                for fi87 := 0; fi87 < 6; fi87++ {
                    switch "" {
                    case "a":
                        if (1 == (^(v5))) {
                            switch "" {
                            case "foo bar":
                                v6 = append(v6, 80)
                            default:
                                for fi88 := 0; fi88 < 7; fi88++ {
                                    if false {
                                        var v89 int32 = -2147483648
                                        v90 := GenS5{F0: -9223372036854775808, F1: 32767, F2: 845}
                                        v91 := GenS5{F0: 1, F1: 0, F2: 2}
                                        _ = v89
                                        _ = v90
                                        _ = v91
                                    } else {
                                        v1 += v1
                                        var v92 int64 = 9223372036854775807
                                        v51 = ((56 >= (^(v1))) || v51)
                                        v93 := GenS6{F0: -1, F1: int(v5)}
                                        _ = v92
                                        _ = v93
                                    }
                                    v94 := GenS1{F0: int16(fi87), F1: 0, F2: 0}
                                    var v95 bool = (!(!(-536 < 9223372036854775807)))
                                    _ = fi88
                                    _ = v94
                                    _ = v95
                                }
                            }
                            v96 := v4[((("" + "foo bar") + "a") + ("" + ("" + "a")))]
                            _ = len(v6)
                            _ = v96
                        }
                    case "hello":
                        var v97 int = (v2 / ((int(v5)) | 1))
                        _ = v97
                    case "":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    v6 = append(v6, 295)
                    _ = fi87
                }
            }
        }
        if (924 >= 42) {
            var v98 uint16 = 1
            _ = v98
        }
    } else {
        if (v51 || v51) {
            v99 := GenS2{F0: ("\n\t" + "\n\t"), F1: (v51 && v51), F2: 32767}
            _ = v99
        }
        var v100 uint64 = 18446744073709551615
        _ = v100
    }
    delete(v3, "hello")
    v101 := GenS3{F0: 571}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v101
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
