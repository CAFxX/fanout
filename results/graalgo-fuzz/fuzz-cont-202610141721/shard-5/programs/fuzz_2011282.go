package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 string; F2 uint16 }
type GenS9 struct { F0 int16; F1 uint64; F2 int64 }
type GenS10 struct { F0 string; F1 int16; F2 int }
type GenS11 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS12 struct { F0 uint64; F1 string; F2 int }
type GenS13 struct { F0 int8; F1 uint16; F2 uint64 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]bool)
    var v3 int64 = 85
    v2[2147483647] = true
    if (!false) {
        _ = len(v1)
        fmt.Println(v3)
    } else {
        v3 += 9223372036854775807
    }
    switch 32767 {
    case -9223372036854775808:
        v3 -= 9223372036854775807
    case -388:
        var v4 uint32 = 950
        _ = v4
    }
    v5 := GenS1{F0: -1, F1: 42, F2: 1000}
    v1 = append(v1, "\n\t")
    var v6 int32 = 0
    v6 = (v6 * (v6 >> ((1) & 63)))
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            if (!(185 > 42)) {
                var v8 int = (fi7 << ((0) & 63))
                v2[2147483647] = (3.14159 == -2.25)
                _ = v8
            } else {
                for fi9 := 0; fi9 < 6; fi9++ {
                    v10 := GenS1{F0: -24, F1: 42, F2: 123}
                    _ = fi9
                    _ = v10
                }
            }
            _ = fi7
        }
    } else {
        switch 32768 {
        case 535:
            var v11 uint64 = 0
            _ = v11
        case 32767:
            v12 := GenS2{F0: "\n\t"}
            _ = v12
        case -929:
            v2[188] = ((-1 >= -1) && false)
        default:
            var v13 float32 = ((float32(1.5) / 1.5) + -2.25)
            _ = v13
        }
    }
    switch "" {
    case "a":
        v14 := GenS3{F0: 1}
        _ = v14
    case "\n\t":
        for fi15 := 0; fi15 < 10; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                if false {
                    v17 := GenS2{F0: "\n\t"}
                    fmt.Println(fi16, fi16)
                    var v18 uint32 = 65535
                    _ = v17
                    _ = v18
                } else {
                    if (!(("" + ((((("\n\t" + "foo bar") + "a") + "a") + ("foo bar" + "a")) + "foo bar")) <= ("\n\t" + "\n\t"))) {
                        fmt.Println(v6, v3, v3)
                        var v19 uint16 = 578
                        _ = v19
                    } else {
                        if (2147483647 <= fi15) {
                            v6 += (v6 % ((int32(fi15)) | 1))
                            var v20 uint = 268
                            _ = v20
                        } else {
                            v6 -= (v6 >> ((255) & 63))
                            if len(v1) > 0 {
                                v1[0] = ((("hello" + ("" + ("hello" + ("\n\t" + "foo bar")))) + "foo bar") + ("\n\t" + "\n\t"))
                            }
                            var v21 uint64 = 1
                            _ = v21
                        }
                        switch "" {
                        case "foo bar":
                            fmt.Println(fi16)
                        case "\n\t":
                            v22 := v2[2]
                            _ = v22
                        case "hello":
                            var v23 float32 = float32(1.5)
                            _ = v23
                        default:
                            if true {
                                if ((v3 / ((-80) | 1)) > v3) {
                                    var v24 uint64 = 704
                                    _ = v24
                                } else {
                                    _ = len(v1)
                                    var v25 float32 = (0.1 - 0.0)
                                    v25 += 0.0
                                    _ = v25
                                }
                                v26 := GenS4{F0: 1, F1: 42, F2: 65535}
                                var v27 uint64 = 966
                                _ = v26
                                _ = v27
                            } else {
                                if (("a" >= "") && (255 < 0)) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v28 uint8 = 49
                                    v29 := GenS5{F0: 18446744073709551615}
                                    _ = v28
                                    _ = v29
                                }
                            }
                            if true {
                                v6 -= 0
                            }
                        }
                    }
                }
                var v30 uint32 = 2147483647
                var v31 uint16 = 261
                _ = fi16
                _ = v30
                _ = v31
            }
            fmt.Println(fi15)
            var v32 int32 = 0
            _ = fi15
            _ = v32
        }
    case "hello":
        if (!true) {
            v33 := v2[4294967295]
            _ = v33
        } else {
            if (!true) {
                if (42 == 995) {
                    v3 = 1
                    switch 493 {
                    case -2147483648:
                        v34 := GenS6{F0: (v6 - (v6 % (((v6 ^ -680)) | 1)))}
                        _ = v34
                    }
                }
                if true {
                    v3 += v3
                    if (3.14159 >= (((1e10 * (3.14159 + (1e10 * float64((1.5 / 1e10))))) + (((1.5 / float64(1.5)) / 0.1) + float64(0.1))) / 1.0)) {
                        v3 += v3
                        switch "" {
                        case "hello":
                            for fi35 := 0; fi35 < 13; fi35++ {
                                v6 = v6
                                _ = fi35
                            }
                        }
                    } else {
                        var v36 uint16 = 859
                        var v37 string = ""
                        _ = v36
                        _ = v37
                    }
                    v2[2147483648] = (42 >= 1)
                }
                v3 += (v3 % ((v3) | 1))
            } else {
                if true {
                    switch 4294967295 {
                    case 128:
                        var v38 uint8 = 0
                        _ = v38
                    case 399:
                        fmt.Println(v6, v3)
                    default:
                        for fi39 := 0; fi39 < 17; fi39++ {
                            var v40 uint16 = 733
                            v41 := GenS7{F0: fi39}
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                    }
                    v42 := GenS1{F0: 1, F1: (v6 ^ v6), F2: 1}
                    for fi43 := 0; fi43 < 17; fi43++ {
                        switch (fi43 - 2147483647) {
                        case 0:
                            var v44 int8 = 127
                            _ = v44
                        case -9223372036854775808:
                            switch "" {
                            case "":
                                var v45 uint64 = 4294967295
                                _ = v45
                            case "foo bar":
                                v6 = (v6 % ((v6) | 1))
                            }
                        case 658:
                            _ = len(v1)
                        default:
                            for fi46 := 0; fi46 < 11; fi46++ {
                                switch (^(fi43)) {
                                case -861:
                                    v3 = 622
                                case -703:
                                    v47 := GenS7{F0: 9223372036854775807}
                                    _ = v47
                                case 892:
                                    if true {
                                        switch fi46 {
                                        case 929:
                                            v3 -= 991
                                        }
                                        if (!true) {
                                            switch "\n\t" {
                                            case "a":
                                                var v48 float32 = ((0.0 / float32(1e10)) / -2.25)
                                                _ = v48
                                            case "\n\t":
                                                v49 := GenS8{F0: (v3 > (v3 ^ v3)), F1: ("hello" + "a"), F2: 65535}
                                                _ = v49
                                            default:
                                                if false {
                                                    v50 := GenS9{F0: -32768, F1: 604, F2: v3}
                                                    var v51 string = ("a" + "a")
                                                    if (false && (!false)) {
                                                        if ((!true) && (420 >= 255)) {
                                                            v3 -= (v3 / ((-808) | 1))
                                                        } else {
                                                            v52 := GenS10{F0: (v51 + v51), F1: 42, F2: (fi43 << ((18446744073709551615) & 63))}
                                                            fmt.Println(v6, fi43, v3)
                                                            _ = v52
                                                        }
                                                        fmt.Println(v51, v3)
                                                    } else {
                                                        if false {
                                                            v3 -= 42
                                                            _ = len(v1)
                                                            var v53 int = fi43
                                                            _ = v53
                                                        }
                                                        delete(v2, 42)
                                                        var v54 uint = 2
                                                        v51 = "foo bar"
                                                        _ = v54
                                                    }
                                                    _ = v50
                                                    _ = v51
                                                }
                                            }
                                            switch fi43 {
                                            case 0:
                                                if true {
                                                    v6 += (v6 & (v6 ^ (v6 >> ((0) & 63))))
                                                    if true {
                                                        switch fi46 {
                                                        case 2147483647:
                                                            v2[42] = true
                                                        case 32768:
                                                            if true {
                                                                var v55 bool = (18446744073709551615 <= 699)
                                                                v3 = -1
                                                                v3 = (v3 % (((v3 | (v3 | v3))) | 1))
                                                                var v56 int8 = -1
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                        default:
                                                            v57 := GenS11{F0: 780, F1: -97, F2: 28}
                                                            _ = v57
                                                        }
                                                        var v58 int64 = v3
                                                        _ = len(v2)
                                                        _ = v58
                                                    }
                                                }
                                            case -309:
                                                switch 32767 {
                                                case -801:
                                                    fmt.Println(fi43, fi46)
                                                case 42:
                                                    fmt.Println(v3, v6, v3)
                                                case 251:
                                                    var v59 int = (fi43 | fi46)
                                                    _ = v59
                                                }
                                            }
                                            v6 = 1
                                        }
                                        if (int64(fi43) != (^(v3))) {
                                            if (!((-2.25 * (((0.1 / -2.25) - 1.5) / 0.1)) == 1e10)) {
                                                fmt.Println(fi46, v3, v6)
                                                delete(v2, 590)
                                                fmt.Println(v6, v6)
                                            } else {
                                                var v60 uint64 = 900
                                                if (2147483647 == 0) {
                                                    v60 = 168
                                                } else {
                                                    var v61 int64 = (v3 + v3)
                                                    switch fi43 {
                                                    case 127:
                                                        if false {
                                                            v62 := GenS10{F0: "", F1: -32768, F2: -9223372036854775808}
                                                            fmt.Println(v6, fi43)
                                                            _ = v62
                                                        }
                                                    case 42:
                                                        var v63 uint8 = 0
                                                        _ = v63
                                                    }
                                                    var v64 string = ""
                                                    _ = v61
                                                    _ = v64
                                                }
                                                v65 := GenS12{F0: (^(v60)), F1: "foo bar", F2: fi43}
                                                _ = v60
                                                _ = v65
                                            }
                                            fmt.Println(v3, v3, v3)
                                            v6 = (v6 * 0)
                                        } else {
                                            var v66 int64 = v3
                                            _ = v66
                                        }
                                    } else {
                                        _ = len(v1)
                                        v6 += (v6 / (((v6 >> ((0) & 63))) | 1))
                                        var v67 float32 = (-2.25 * -2.25)
                                        v2[656] = ((v67 + 1e10) != 1e10)
                                        v6 = v6
                                        if false {
                                            fmt.Println(v67, fi46, v3)
                                            var v68 uint = 4294967295
                                            _ = v68
                                        } else {
                                            var v69 uint64 = 1
                                            _ = v69
                                        }
                                        _ = v67
                                    }
                                default:
                                    var v70 uint16 = 0
                                    _ = v70
                                }
                                fmt.Println(v3, fi43)
                                switch "\n\t" {
                                case "hello":
                                    switch fi46 {
                                    case 0:
                                        if ((fi43 + (fi46 / ((4294967295) | 1))) < 2147483647) {
                                            v71 := GenS11{F0: 882, F1: -110, F2: 42}
                                            var v72 bool = false
                                            _ = v71
                                            _ = v72
                                        } else {
                                            if true {
                                                v6 += v6
                                                var v73 uint32 = 1
                                                _ = v73
                                            } else {
                                                var v74 int64 = (-(v3))
                                                _ = v74
                                            }
                                            var v75 int = 103
                                            v75 += (^(v75))
                                            _ = v75
                                        }
                                    case -2147483648:
                                        v76 := GenS10{F0: ("hello" + (("foo bar" + "hello") + "a")), F1: -1, F2: fi46}
                                        _ = v76
                                    }
                                case "":
                                    var v77 float32 = float32(1e10)
                                    _ = v77
                                default:
                                    v1 = append(v1, "\n\t")
                                }
                                fmt.Println(fi43)
                                _ = fi46
                            }
                        }
                        switch (-(fi43)) {
                        case 1:
                            for fi78 := 0; fi78 < 13; fi78++ {
                                v79 := GenS1{F0: 127, F1: v6, F2: 1000}
                                delete(v2, 2)
                                v80 := GenS9{F0: 1000, F1: 632, F2: (v3 - (v3 | v3))}
                                _ = fi78
                                _ = v79
                                _ = v80
                            }
                        case 32768:
                            _ = len(v2)
                        case -827:
                            var v81 int8 = -128
                            _ = v81
                        default:
                            fmt.Println(fi43)
                        }
                        if true {
                            v3 += (-(v3))
                            var v82 uint8 = 249
                            _ = v82
                        } else {
                            switch (fi43 - fi43) {
                            case 32768:
                                var v83 uint64 = 4294967295
                                _ = v83
                            case 514:
                                v6 = -1
                            default:
                                switch "\n\t" {
                                case "hello":
                                    v84 := GenS10{F0: "foo bar", F1: 0, F2: fi43}
                                    _ = v84
                                case "a":
                                    v6 -= (v6 * (v6 * (v6 ^ v6)))
                                default:
                                    var v85 string = "foo bar"
                                    _ = v85
                                }
                            }
                        }
                        _ = fi43
                    }
                    _ = v42
                } else {
                    v2[538] = false
                    v86 := GenS13{F0: 42, F1: 917, F2: 1}
                    _ = v86
                }
                if true {
                    v87 := GenS13{F0: -68, F1: 341, F2: 374}
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    _ = v87
                }
            }
        }
    default:
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
