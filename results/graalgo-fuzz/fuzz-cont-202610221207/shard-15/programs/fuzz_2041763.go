package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32; F1 int64; F2 int64 }
type GenS9 struct { F0 int64; F1 string; F2 uint32 }
type GenS10 struct { F0 bool; F1 bool; F2 int32 }
type GenS11 struct { F0 uint64; F1 bool }
type GenS12 struct { F0 string; F1 int64; F2 int64 }
func genFn0(p0 uint8, p1 string, p2 int16) uint16 {
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 string = "foo bar"
        _ = fi4
        _ = v5
    }
    v6 := []int{304, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 65535
}
func genFn1(p0 string, p1 string) uint32 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 string = (p0 + ((p1 + p0) + p0))
        if false {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 int8 = 127
                var v11 int8 = (v10 % (((v10 & (^(v10)))) | 1))
                v12 := GenS1{F0: ("foo bar" + "a"), F1: genFn0(2, p0, 1000)}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        } else {
            v13 := []int{fi7}
            if (float64(3.14159) < 3.14159) {
                _ = len(v13)
                _ = len(v13)
            }
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    v14 := []int{4294967295, 485, -9223372036854775808}
    p1 = (p0 + p1)
    _ = p0
    _ = p1
    _ = v14
    return 4294967295
}

func main() {
    v1 := []int{2147483647, 4294967295, 32768, 2147483647}
    var v2 int = -969
    v3 := make(map[int]int)
    v2 = v2
    v2 = 482
    v2 = (-(v2))
    fmt.Println(v2)
    var v15 uint16 = 42
    if (!true) {
        if (false || false) {
            var v16 uint32 = 899
            switch (v2 >> ((744) & 63)) {
            case 2147483647:
                v15 += v15
            case 9223372036854775807:
                if ((^(v15)) > v15) {
                    if len(v1) > 0 {
                        v1[0] = v2
                    }
                    if ((!((1e10 * (-2.25 - (1e10 + -2.25))) == -2.25)) && (!(v15 < uint16(v16)))) {
                        var v17 uint32 = v16
                        v18 := GenS1{F0: ("hello" + ("hello" + "foo bar")), F1: (v15 / ((genFn0(95, "\n\t", 42)) | 1))}
                        _ = v17
                        _ = v18
                    }
                    fmt.Println(v2, v15)
                } else {
                    var v19 int64 = -908
                    var v20 int = (v2 % (((v2 % ((-437) | 1))) | 1))
                    v15 += genFn0(uint8(v15), "foo bar", -32768)
                    _ = v19
                    _ = v20
                }
            default:
                for fi21 := 0; fi21 < 7; fi21++ {
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v23 := GenS2{F0: false}
                        v24 := GenS1{F0: "foo bar", F1: genFn0(42, "\n\t", -389)}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = fi21
                }
            }
            for fi25 := 0; fi25 < 9; fi25++ {
                v1 = append(v1, (v2 | (v2 * v2)))
                for fi26 := 0; fi26 < 2; fi26++ {
                    var v27 int8 = -1
                    fmt.Println(v2, v16, fi26)
                    var v28 int16 = -506
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                if (!true) {
                    v29 := GenS3{F0: 18446744073709551615, F1: 332}
                    if false {
                        v30 := GenS4{F0: -805, F1: true}
                        _ = len(v1)
                        switch "" {
                        case "foo bar":
                            for fi31 := 0; fi31 < 11; fi31++ {
                                if (!((!true) && false)) {
                                    var v32 uint64 = 1
                                    v33 := GenS5{F0: 100}
                                    var v34 uint8 = 128
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                                var v35 uint16 = v15
                                _ = fi31
                                _ = v35
                            }
                        case "":
                            if false {
                                for fi36 := 0; fi36 < 18; fi36++ {
                                    fmt.Println(v2, v15, v2)
                                    var v37 int = fi25
                                    _ = fi36
                                    _ = v37
                                }
                                switch "\n\t" {
                                case "a":
                                    fmt.Println(v2, v15)
                                case "":
                                    v15 = v15
                                case "foo bar":
                                    if len(v1) > 0 {
                                        v1[0] = fi25
                                    }
                                default:
                                    var v38 int64 = -16
                                    _ = v38
                                }
                            } else {
                                var v39 int32 = 2147483647
                                v39 += (v39 << ((18446744073709551615) & 63))
                                v15 = 893
                                _ = v39
                            }
                        case "a":
                            v3[-51] = fi25
                        default:
                            delete(v3, fi25)
                        }
                        _ = v30
                    }
                    v40 := GenS6{F0: (v16 >> ((uint(v15)) & 63)), F1: 2147483647}
                    _ = v29
                    _ = v40
                } else {
                    var v41 uint64 = 584
                    _ = v41
                }
                _ = fi25
            }
            _ = v16
        } else {
            for fi42 := 0; fi42 < 18; fi42++ {
                v15 = genFn0(255, "a", -590)
                _ = fi42
            }
            if true {
                if false {
                    var v43 uint64 = 18446744073709551615
                    for fi44 := 0; fi44 < 8; fi44++ {
                        for fi45 := 0; fi45 < 15; fi45++ {
                            v46 := GenS3{F0: 0, F1: (v43 >> ((4294967295) & 63))}
                            v47 := GenS4{F0: 1, F1: (!(1 < 65535))}
                            var v48 int32 = 172
                            _ = fi45
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                        _ = fi44
                    }
                    v49 := GenS7{F0: 2}
                    _ = v43
                    _ = v49
                }
                if true {
                    var v50 bool = true
                    _ = v50
                }
                var v51 uint = 1
                _ = v51
            }
            switch ("foo bar" + "foo bar") {
            case "":
                for fi52 := 0; fi52 < 8; fi52++ {
                    v53 := GenS4{F0: 540, F1: false}
                    if false {
                        v54 := GenS2{F0: (true && (false && (127 > 90)))}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = len(v1)
                        _ = v54
                    } else {
                        var v55 string = ("a" + "\n\t")
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi56 := 0; fi56 < 1; fi56++ {
                            var v57 int64 = -1
                            _ = fi56
                            _ = v57
                        }
                        _ = v55
                    }
                    _ = fi52
                    _ = v53
                }
            case "a":
                var v58 float64 = -2.25
                _ = v58
            }
        }
        for fi59 := 0; fi59 < 17; fi59++ {
            var v60 float32 = 3.14159
            _ = fi59
            _ = v60
        }
    } else {
        for fi61 := 0; fi61 < 9; fi61++ {
            fmt.Println(fi61, v2)
            v2 = v2
            _ = fi61
        }
        for fi62 := 0; fi62 < 15; fi62++ {
            switch "foo bar" {
            case "hello":
                switch fi62 {
                case -2147483648:
                    if (127 == 127) {
                        v3[fi62] = 286
                        var v63 int32 = 1
                        _ = v63
                    } else {
                        if false {
                            v2 += v2
                        } else {
                            var v64 int = 32768
                            _ = v64
                        }
                    }
                default:
                    var v65 uint32 = 545
                    _ = v65
                }
            case "":
                if (!(!(!(true && false)))) {
                    v15 = v15
                    switch "hello" {
                    case "\n\t":
                        switch (fi62 - v2) {
                        case -2147483648:
                            for fi66 := 0; fi66 < 11; fi66++ {
                                v2 = (v2 + 433)
                                v67 := GenS5{F0: 127}
                                _ = fi66
                                _ = v67
                            }
                        case 1:
                            for fi68 := 0; fi68 < 16; fi68++ {
                                v15 += (v15 - v15)
                                _ = fi68
                            }
                        }
                    case "hello":
                        for fi69 := 0; fi69 < 20; fi69++ {
                            if (!((v15 >= v15) && true)) {
                                var v70 bool = ((42 <= 255) && false)
                                v71 := GenS4{F0: 0, F1: (18446744073709551615 == 42)}
                                if v70 {
                                    if (203 < 255) {
                                        fmt.Println(fi62, v15)
                                    } else {
                                        delete(v3, v2)
                                        v72 := GenS2{F0: (uint32(v15) != 719)}
                                        if v70 {
                                            v70 = v70
                                        }
                                        _ = v72
                                    }
                                } else {
                                    fmt.Println(fi69)
                                    var v73 int64 = 693
                                    delete(v3, 848)
                                    _ = v73
                                }
                                _ = v70
                                _ = v71
                            } else {
                                v74 := GenS8{F0: 42, F1: 0, F2: -654}
                                if (-2.25 > 0.1) {
                                    fmt.Println(fi69, v15, fi69)
                                    if false {
                                        var v75 bool = false
                                        if v75 {
                                            v15 = (v15 ^ (v15 / ((1) | 1)))
                                            if v75 {
                                                v2 = (v2 * v2)
                                            }
                                            v15 -= v15
                                            v75 = false
                                            var v76 bool = ((2 <= 218) || true)
                                            _ = v76
                                        }
                                        var v77 int = 587
                                        _ = len(v1)
                                        fmt.Println(v75)
                                        fmt.Println(fi69, v77)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v75
                                        _ = v77
                                    }
                                }
                                if (!true) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v78 uint16 = 255
                                        _ = v78
                                    }
                                    var v79 int = fi62
                                    v15 = v15
                                    _ = v79
                                } else {
                                    v80 := GenS9{F0: 978, F1: "hello", F2: genFn1("", "\n\t")}
                                    v81 := GenS10{F0: ((v2 - fi69) > 475), F1: (!(-9223372036854775808 == 183)), F2: 0}
                                    _ = v80
                                    _ = v81
                                }
                                _ = v74
                            }
                            switch ("foo bar" + (("hello" + "a") + "foo bar")) {
                            case "hello":
                                v15 += 418
                            case "\n\t":
                                fmt.Println(fi69, v2)
                            case "a":
                                switch ("" + ("\n\t" + "foo bar")) {
                                case "foo bar":
                                    v15 = v15
                                case "a":
                                    if false {
                                        _ = len(v1)
                                    }
                                case "":
                                    v15 += (v15 / (((^(v15))) | 1))
                                }
                            }
                            var v82 uint = 65535
                            _ = fi69
                            _ = v82
                        }
                    case "a":
                        v2 += v2
                    }
                }
            default:
                for fi83 := 0; fi83 < 3; fi83++ {
                    v15 += (v15 >> ((2) & 63))
                    switch "foo bar" {
                    case "\n\t":
                        v15 = 0
                    case "hello":
                        v15 -= v15
                    case "foo bar":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        var v84 int64 = -9223372036854775808
                        _ = v84
                    }
                    switch 42 {
                    case -35:
                        v2 = (^(v2))
                    case 128:
                        switch v2 {
                        case 9223372036854775807:
                            var v85 int16 = 0
                            _ = v85
                        default:
                            if false {
                                var v86 int16 = -32768
                                v86 -= v86
                                _ = v86
                            } else {
                                var v87 uint32 = 4294967295
                                v88 := GenS11{F0: 260, F1: (205 < 128)}
                                _ = v87
                                _ = v88
                            }
                        }
                    default:
                        v89 := GenS8{F0: 1, F1: -9223372036854775808, F2: 501}
                        _ = v89
                    }
                    _ = fi83
                }
            }
            var v90 uint8 = 93
            v91 := GenS12{F0: "", F1: -122, F2: 1}
            _ = fi62
            _ = v90
            _ = v91
        }
        if true {
            if true {
                v15 = v15
                fmt.Println(v15, v2)
                var v92 int64 = -725
                _ = v92
            }
            switch ("\n\t" + ("a" + (("a" + "hello") + "hello"))) {
            case "":
                var v93 bool = (812 >= 2147483647)
                _ = v93
            case "\n\t":
                v1 = append(v1, v2)
            case "hello":
                _ = len(v1)
            default:
                var v94 int8 = 4
                _ = v94
            }
            var v95 string = ("" + "foo bar")
            _ = v95
        } else {
            var v96 int8 = -56
            fmt.Println(v15, v96)
            if (("foo bar" + "hello") <= "a") {
                v15 = v15
                switch ("\n\t" + (((("foo bar" + "\n\t") + "\n\t") + "") + "")) {
                case "hello":
                    v97 := GenS2{F0: false}
                    _ = v97
                case "\n\t":
                    var v98 int = (v2 / (((v2 - 32768)) | 1))
                    _ = v98
                }
            }
            _ = v96
        }
    }
    var v99 uint16 = 65535
    v100 := GenS6{F0: 1, F1: genFn1("a", "")}
    var v101 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v101:", v101)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v15
    _ = v2
    _ = v3
    _ = v99
}
