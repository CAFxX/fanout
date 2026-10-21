package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int32; F1 int64; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int; F2 int32 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 uint; F1 bool; F2 uint64 }
type GenS8 struct { F0 uint; F1 int16; F2 bool }
type GenS9 struct { F0 int32; F1 int8 }
type GenS10 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 string; F1 uint16 }
type GenS15 struct { F0 int16; F1 int32 }

func main() {
    var v1 int8 = 39
    v2 := make(map[uint32]int)
    v3 := []int{32768}
    var v4 int16 = 32767
    var v5 uint32 = 988
    var v6 uint32 = v5
    var v7 int8 = -128
    switch ((((("" + "") + "foo bar") + ("\n\t" + "\n\t")) + "\n\t") + "hello") {
    case "\n\t":
        v8 := GenS1{F0: 128, F1: v5}
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        switch 438 {
        case 9223372036854775807:
            switch 127 {
            case 1:
                _ = len(v3)
            case 32767:
                v10 := v2[(v6 & v5)]
                _ = v10
            case -9223372036854775808:
                var v11 float32 = float32(3.14159)
                _ = v11
            default:
                var v12 float32 = ((((1e10 / 1.0) - 1e10) * 1.5) * 1.5)
                _ = v12
            }
        case 47:
            if (18446744073709551615 < 18446744073709551615) {
                v6 = v6
                fmt.Println(v5)
                for fi13 := 0; fi13 < 9; fi13++ {
                    var v14 float32 = (3.14159 / float32(-2.25))
                    v5 += v5
                    fmt.Println(v14, v4)
                    _ = fi13
                    _ = v14
                }
            } else {
                v5 -= (v6 / ((v6) | 1))
                fmt.Println(fi9)
            }
        case 664:
            var v15 float32 = -2.25
            _ = v15
        }
        _ = fi9
    }
    v5 = 42
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(v6, v5)
        if ("hello" != (("a" + ("" + "hello")) + ("hello" + ""))) {
            v17 := GenS1{F0: 255, F1: (v5 >> ((42) & 63))}
            var v18 string = "a"
            _ = v17
            _ = v18
        } else {
            fmt.Println(v4, v1, v4)
            if true {
                switch ("hello" + "") {
                case "\n\t":
                    if (("hello" + ("hello" + "foo bar")) >= "hello") {
                        if false {
                            var v19 int = 0
                            _ = v19
                        } else {
                            v2[(^(v5))] = 156
                            var v20 int16 = int16(v7)
                            v2[(v5 | (^(v6)))] = fi16
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 4; fi21++ {
                            v1 += int8(v6)
                            var v22 float32 = 1e10
                            _ = len(v3)
                            if len(v3) > 0 {
                                v3[0] = (fi16 * fi21)
                            }
                            fmt.Println(v22, v22, v6)
                            if true {
                                v6 += (v5 + v6)
                                switch "hello" {
                                case "":
                                    var v23 bool = true
                                    v6 += (v5 / (((v5 >> ((383) & 63))) | 1))
                                    var v24 int16 = (v4 | v4)
                                    v25 := GenS1{F0: 80, F1: v6}
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                case "\n\t":
                                    fmt.Println(v7, fi16, fi16)
                                default:
                                    v26 := GenS1{F0: 0, F1: 2}
                                    if (735 == 18446744073709551615) {
                                        if ((v5 & (v6 - 843)) >= v6) {
                                            v5 -= (v6 + (^(v6)))
                                            v27 := GenS1{F0: 42, F1: (v5 | (v6 ^ (^(v6))))}
                                            _ = v27
                                        } else {
                                            var v28 float64 = 0.1
                                            v29 := GenS2{F0: -1, F1: 1, F2: (false && true)}
                                            _ = v28
                                            _ = v29
                                        }
                                        v4 += (v4 + (v4 ^ (v4 + 734)))
                                        var v30 uint8 = 92
                                        v2[v5] = 465
                                        _ = v30
                                    } else {
                                        v31 := GenS2{F0: 42, F1: 1, F2: false}
                                        v32 := GenS3{F0: (2 != 18446744073709551615)}
                                        var v33 int16 = (v4 ^ (v4 | 32767))
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v26
                                }
                            }
                            _ = fi21
                            _ = v22
                        }
                    } else {
                        if false {
                            switch "a" {
                            case "a":
                                fmt.Println(v5, v5, v6)
                            case "\n\t":
                                for fi34 := 0; fi34 < 13; fi34++ {
                                    if false {
                                        v1 = (v7 & 42)
                                        var v35 uint = 2
                                        v6 = 2147483648
                                        _ = v35
                                    }
                                    _ = fi34
                                }
                            }
                        }
                        v7 = (-(v1))
                    }
                default:
                    for fi36 := 0; fi36 < 13; fi36++ {
                        if false {
                            var v37 int32 = 0
                            switch ("a" + "\n\t") {
                            case "a":
                                fmt.Println(v6)
                            case "\n\t":
                                v38 := GenS2{F0: (v37 * v37), F1: -9223372036854775808, F2: false}
                                _ = v38
                            }
                            _ = v37
                        } else {
                            switch fi16 {
                            case 717:
                                switch fi36 {
                                case 32767:
                                    switch "\n\t" {
                                    case "":
                                        v39 := v2[v6]
                                        _ = v39
                                    case "foo bar":
                                        if ((v7 ^ 100) != (-(v7))) {
                                            v6 += 2147483647
                                        }
                                    case "hello":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        v40 := GenS4{F0: 375, F1: 32767, F2: -354}
                                        _ = v40
                                    }
                                }
                            }
                            var v41 int = fi36
                            fmt.Println(v1)
                            if (!false) {
                                var v42 string = (("foo bar" + ("foo bar" + "foo bar")) + ("foo bar" + ("\n\t" + "foo bar")))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v42
                            } else {
                                switch 128 {
                                case 301:
                                    var v43 int = (fi16 ^ fi16)
                                    _ = v43
                                case 2147483647:
                                    v44 := v2[641]
                                    _ = v44
                                case 786:
                                    var v45 int32 = -1
                                    _ = v45
                                default:
                                    v1 = v1
                                }
                            }
                            _ = v41
                        }
                        _ = fi36
                    }
                }
                var v46 uint16 = 32767
                v4 -= v4
                _ = v46
            }
        }
        _ = fi16
    }
    if ((v5 >= (^(v5))) || false) {
        v4 -= v4
    } else {
        v47 := GenS5{F0: 18446744073709551615, F1: v6}
        if ((!(true && (0 == 255))) && (-262 > 500)) {
            v7 = int8(v5)
            if false {
                var v48 int64 = 875
                switch int(v6) {
                case 443:
                    v5 -= v6
                case -489:
                    delete(v2, v5)
                default:
                    var v49 bool = true
                    _ = v49
                }
                _ = v48
            } else {
                v1 -= (v7 - -68)
            }
            v6 -= (v5 / ((v6) | 1))
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch "" {
            case "a":
                for fi50 := 0; fi50 < 14; fi50++ {
                    if false {
                        v5 -= (v5 % ((v5) | 1))
                    } else {
                        v4 = v4
                        var v51 bool = true
                        _ = v51
                    }
                    var v52 uint16 = 2
                    _ = fi50
                    _ = v52
                }
            case "hello":
                switch ("a" + ("hello" + ("foo bar" + ("foo bar" + "foo bar")))) {
                case "foo bar":
                    var v53 bool = false
                    _ = v53
                case "a":
                    var v54 uint16 = 121
                    _ = v54
                default:
                    v55 := GenS4{F0: 1, F1: -288, F2: -2147483648}
                    _ = v55
                }
            case "\n\t":
                v56 := GenS3{F0: ((-(v4)) != (v4 << ((48) & 63)))}
                _ = v56
            }
            switch ("foo bar" + "hello") {
            case "a":
                fmt.Println(v5)
            case "foo bar":
                fmt.Println(v7, v6)
            case "":
                if true {
                    var v57 bool = (!false)
                    if true {
                        v5 = v5
                    } else {
                        v58 := GenS6{F0: 2, F1: 763}
                        v59 := GenS6{F0: 231, F1: 421}
                        v57 = v57
                        _ = v58
                        _ = v59
                    }
                    _ = v57
                } else {
                    v60 := GenS7{F0: 980, F1: (65535 >= 65535), F2: 4294967295}
                    _ = v60
                }
            }
        }
        v61 := GenS6{F0: 1, F1: 4294967295}
        _ = v47
        _ = v61
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v7 += (v7 - v7)
    switch 42 {
    case -563:
        v62 := GenS8{F0: 1, F1: (v4 & 32767), F2: false}
        _ = v62
    default:
        switch 373 {
        case 127:
            for fi63 := 0; fi63 < 4; fi63++ {
                _ = len(v2)
                delete(v2, (v6 | v5))
                _ = fi63
            }
        case 32767:
            v64 := GenS9{F0: 0, F1: 1}
            _ = v64
        }
    }
    fmt.Println(v7)
    for fi65 := 0; fi65 < 4; fi65++ {
        var v66 int8 = (v1 | v7)
        _ = fi65
        _ = v66
    }
    v67 := GenS1{F0: 127, F1: v6}
    v7 = v7
    if (true && (65535 > 302)) {
        v68 := GenS6{F0: 65535, F1: 1}
        for fi69 := 0; fi69 < 6; fi69++ {
            switch fi69 {
            case 127:
                v5 += 2147483647
            }
            if true {
                var v70 float32 = (0.1 * 0.0)
                switch "foo bar" {
                case "hello":
                    v71 := GenS10{F0: v6, F1: 42, F2: 42}
                    _ = v71
                case "foo bar":
                    if false {
                        if (true || (-1 < (v7 & (v1 & (v1 / (((v1 % ((-128) | 1))) | 1)))))) {
                            var v72 int16 = v4
                            if false {
                                var v73 float64 = (float64((v70 + (0.0 - float32(0.1)))) + (float64(v70) - 1.5))
                                var v74 uint = 0
                                _ = v73
                                _ = v74
                            } else {
                                var v75 int64 = 1
                                _ = v75
                            }
                            _ = v72
                        }
                    } else {
                        if true {
                            fmt.Println(v70, fi69)
                            switch (fi69 % ((127) | 1)) {
                            case -221:
                                switch fi69 {
                                case 2147483647:
                                    for fi76 := 0; fi76 < 17; fi76++ {
                                        var v77 float32 = v70
                                        _ = fi76
                                        _ = v77
                                    }
                                case 1:
                                    v2[v5] = fi69
                                }
                            }
                            if ((("hello" + "foo bar") + "hello") > ((("a" + ("hello" + "hello")) + ("\n\t" + "foo bar")) + (("\n\t" + "hello") + "\n\t"))) {
                                if false {
                                    var v78 uint8 = 169
                                    for fi79 := 0; fi79 < 12; fi79++ {
                                        v78 = v78
                                        v80 := GenS4{F0: 42, F1: 32768, F2: -3}
                                        v4 = (v4 * v4)
                                        _ = fi79
                                        _ = v80
                                    }
                                    var v81 bool = false
                                    _ = v78
                                    _ = v81
                                }
                                v7 = (v1 | (v1 ^ v7))
                            }
                        }
                        v82 := GenS11{F0: 0}
                        _ = v82
                    }
                case "a":
                    var v83 int32 = 42
                    _ = v83
                default:
                    if (("\n\t" + ("a" + "\n\t")) > "a") {
                        switch "\n\t" {
                        case "\n\t":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "":
                            var v84 uint8 = 79
                            _ = v84
                        case "hello":
                            var v85 int8 = v1
                            _ = v85
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v86 float64 = 0.0
                        _ = v86
                    }
                }
                if (true || true) {
                    var v87 string = ("a" + (("" + "hello") + "a"))
                    _ = v87
                }
                _ = v70
            }
            if true {
                v7 += v1
            } else {
                v88 := GenS12{F0: -255, F1: 0, F2: (v6 / ((42) | 1))}
                for fi89 := 0; fi89 < 8; fi89++ {
                    var v90 int = fi69
                    if false {
                        if false {
                            if (76 >= 127) {
                                var v91 int32 = -1
                                v90 -= 0
                                _ = v91
                            }
                            v1 = 42
                        } else {
                            fmt.Println(v6, v4)
                            var v92 float32 = 1.5
                            _ = v92
                        }
                        var v93 int64 = -908
                        _ = v93
                    } else {
                        var v94 uint32 = v5
                        _ = v94
                    }
                    v6 -= (v5 - v5)
                    _ = fi89
                    _ = v90
                }
                _ = v88
            }
            _ = fi69
        }
        _ = v68
    } else {
        v95 := GenS13{F0: ("\n\t" + "a")}
        var v96 uint16 = 2
        fmt.Println(v6, v96, v1)
        _ = v95
        _ = v96
    }
    v97 := v2[v6]
    v98 := GenS14{F0: (("a" + "a") + "a"), F1: 1}
    v99 := GenS15{F0: 32767, F1: -2147483648}
    for fi100 := 0; fi100 < 7; fi100++ {
        fmt.Println(v6)
        _ = fi100
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v7
    _ = v97
    _ = v98
    _ = v99
}
