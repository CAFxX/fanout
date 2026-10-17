package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint16 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int32; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint16; F2 uint }
type GenS6 struct { F0 uint16; F1 bool; F2 uint }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint32; F1 int32; F2 int32 }
type GenS10 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS11 struct { F0 uint8; F1 uint64; F2 string }
type GenS12 struct { F0 int16; F1 bool }
type GenS13 struct { F0 int8; F1 uint8; F2 int64 }

func main() {
    var v1 uint64 = 1
    var v2 float64 = 0.0
    var v3 int8 = 1
    var v4 bool = (4294967295 < 2147483647)
    var v5 float64 = v2
    var v6 uint8 = 0
    v7 := []int{}
    v7 = append(v7, 1)
    v1 += v1
    v8 := GenS1{F0: 922, F1: 32767, F2: 1}
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi10
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        v5 = v5
        for fi12 := 0; fi12 < 2; fi12++ {
            if ("hello" >= ("hello" + "\n\t")) {
                fmt.Println(v3, v2)
                if ((((fi11 | fi12) != fi12) && v4) || (620 > 191)) {
                    v13 := GenS2{F0: 400, F1: (-(v3))}
                    var v14 int32 = -420
                    _ = v13
                    _ = v14
                }
            } else {
                var v15 int64 = 976
                if (true && (v15 <= v15)) {
                    v4 = (v4 && true)
                    v15 += (v15 + 42)
                }
                v6 += (v6 & v6)
                fmt.Println(v15)
                _ = v15
            }
            fmt.Println(v5)
            switch 1 {
            case 32767:
                if false {
                    switch (("foo bar" + ("" + "")) + "") {
                    case "\n\t":
                        var v16 uint = 819
                        _ = v16
                    case "":
                        switch (fi12 % ((1) | 1)) {
                        case 4294967295:
                            var v17 int = fi12
                            _ = v17
                        case 2147483647:
                            v5 += (((v5 * v5) / float64(float32(((v2 - (1.5 + v5)) * 0.0)))) + v5)
                        case -302:
                            var v18 uint16 = 32767
                            _ = v18
                        default:
                            v19 := GenS2{F0: fi12, F1: (v3 % ((v3) | 1))}
                            _ = v19
                        }
                    default:
                        v2 = (v5 * 1e10)
                    }
                    if v4 {
                        switch (fi12 % ((fi12) | 1)) {
                        case 32768:
                            if len(v7) > 0 {
                                v7[0] = (fi12 % (((fi11 - fi11)) | 1))
                            }
                        }
                        _ = len(v7)
                    }
                    if (4294967295 <= 255) {
                        if v4 {
                            var v20 float32 = (0.0 - 1.5)
                            switch fi12 {
                            case 32768:
                                var v21 uint8 = (^(v6))
                                _ = v21
                            default:
                                v4 = v4
                            }
                            v2 -= 3.14159
                            v2 -= 0.1
                            fmt.Println(v6)
                            switch 2147483647 {
                            case -9223372036854775808:
                                v2 -= v5
                            case 2147483647:
                                v22 := GenS1{F0: 42, F1: 0, F2: 424}
                                _ = v22
                            }
                            if (547 != 65535) {
                                v4 = v4
                                if v4 {
                                    var v23 int8 = (v3 << ((4294967295) & 63))
                                    var v24 uint8 = v6
                                    v25 := GenS1{F0: 42, F1: 834, F2: 65535}
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 int64 = 9223372036854775807
                                    switch "hello" {
                                    case "\n\t":
                                        v27 := GenS3{F0: -1, F1: 2, F2: 2}
                                        _ = v27
                                    case "":
                                        v1 += (^(v1))
                                    }
                                    _ = v26
                                }
                                v4 = v4
                            } else {
                                var v28 uint16 = 2
                                if (((v6 & (v6 ^ v6)) >= (v6 >> ((1) & 63))) || v4) {
                                    fmt.Println(v4, v5, v6)
                                }
                                switch 127 {
                                case 865:
                                    v4 = (788 != 1)
                                case 1:
                                    if true {
                                        _ = len(v7)
                                        switch (fi12 % (((fi11 * fi11)) | 1)) {
                                        case -821:
                                            v4 = false
                                        case 58:
                                            var v29 uint = 42
                                            _ = v29
                                        case 11:
                                            v1 += 887
                                        default:
                                            v6 = (v6 & v6)
                                            v30 := GenS1{F0: 539, F1: -876, F2: v28}
                                            _ = v30
                                        }
                                        var v31 string = ""
                                        _ = v31
                                    } else {
                                        switch fi12 {
                                        case 127:
                                            var v32 uint8 = (v6 + (v6 + (v6 & v6)))
                                            _ = v32
                                        case 4294967295:
                                            v7 = append(v7, fi11)
                                        }
                                        v2 -= v2
                                    }
                                case -9223372036854775808:
                                    v3 = 100
                                default:
                                    if len(v7) > 0 {
                                        v7[0] = (fi11 + -9223372036854775808)
                                    }
                                }
                                v33 := GenS1{F0: 901, F1: 0, F2: (v28 >> ((581) & 63))}
                                v1 += 42
                                v34 := GenS4{F0: 4294967295, F1: (v28 | (v28 % ((458) | 1)))}
                                _ = v28
                                _ = v33
                                _ = v34
                            }
                            var v35 float64 = ((0.0 - v5) + 1e10)
                            _ = v20
                            _ = v35
                        } else {
                            var v36 uint8 = (v6 * (v6 % ((128) | 1)))
                            var v37 int8 = (v3 + v3)
                            var v38 uint16 = 65535
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                        var v39 int32 = -685
                        _ = v39
                    }
                }
            default:
                v5 -= v2
            }
            _ = fi12
        }
        for fi40 := 0; fi40 < 5; fi40++ {
            if (v1 == (^(v1))) {
                if ((v5 - ((v5 * v5) / ((v2 - (float64((float32(v5) - ((0.1 * 3.14159) + ((0.0 - 1.5) - float32(float64(0.0)))))) + v2)) * v5))) >= (v5 * v5)) {
                    v3 += (v3 << ((269) & 63))
                }
                var v41 int16 = -519
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v41
            } else {
                v2 -= v2
            }
            switch "foo bar" {
            case "hello":
                fmt.Println(v5, v1, fi40)
            case "":
                var v42 uint = 541
                _ = v42
            case "foo bar":
                v6 = (v6 & v6)
            }
            _ = fi40
        }
        _ = fi11
    }
    if (!(65535 > 4294967295)) {
        v43 := GenS3{F0: 1, F1: 2, F2: 2147483648}
        var v44 float32 = 0.1
        for fi45 := 0; fi45 < 18; fi45++ {
            v46 := GenS5{F0: v3, F1: 1, F2: 42}
            for fi47 := 0; fi47 < 19; fi47++ {
                v44 += v44
                _ = fi47
            }
            var v48 int64 = -969
            _ = fi45
            _ = v46
            _ = v48
        }
        _ = v43
        _ = v44
    } else {
        v49 := GenS2{F0: 0, F1: -128}
        v50 := GenS6{F0: 438, F1: (!v4), F2: 42}
        _ = v49
        _ = v50
    }
    var v51 bool = (-2.25 <= float64((0.0 - 0.0)))
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if (0.1 > (v5 * v2)) {
        v1 = v1
    } else {
        if (42 == int(v1)) {
            var v52 uint8 = v6
            for fi53 := 0; fi53 < 1; fi53++ {
                for fi54 := 0; fi54 < 14; fi54++ {
                    var v55 int8 = (v3 ^ (v3 / (((v3 * 1)) | 1)))
                    _ = fi54
                    _ = v55
                }
                var v56 uint8 = (v6 | v52)
                _ = fi53
                _ = v56
            }
            for fi57 := 0; fi57 < 8; fi57++ {
                for fi58 := 0; fi58 < 14; fi58++ {
                    v59 := GenS7{F0: uint(v1), F1: v3}
                    fmt.Println(v6)
                    v5 = (v2 + float64(0.1))
                    _ = fi58
                    _ = v59
                }
                _ = fi57
            }
            _ = v52
        }
        var v60 uint64 = 42
        if ((v5 + 1.5) > (v5 + v5)) {
            fmt.Println(v51, v60, v2)
            v61 := GenS8{F0: (uint16(v1) != 32767)}
            fmt.Println(v4, v2)
            _ = v61
        }
        _ = v60
    }
    v62 := GenS9{F0: 2, F1: 42, F2: 2147483647}
    if (!v51) {
        v6 -= 128
        if v51 {
            for fi63 := 0; fi63 < 9; fi63++ {
                v64 := GenS7{F0: 65535, F1: v3}
                var v65 float32 = (0.0 + 0.1)
                v3 += -11
                _ = fi63
                _ = v64
                _ = v65
            }
        } else {
            v6 = v6
            switch (("\n\t" + "a") + "foo bar") {
            case "hello":
                v66 := GenS10{F0: 65535, F1: 0, F2: 2}
                _ = v66
            case "":
                if v51 {
                    switch 4294967295 {
                    case -9223372036854775808:
                        switch 42 {
                        case 287:
                            v2 = v5
                        default:
                            var v67 bool = v51
                            _ = v67
                        }
                    case 127:
                        v4 = false
                    case 0:
                        switch 2147483647 {
                        case -2147483648:
                            var v68 uint8 = (^(v6))
                            _ = v68
                        }
                    default:
                        for fi69 := 0; fi69 < 10; fi69++ {
                            var v70 bool = v51
                            _ = fi69
                            _ = v70
                        }
                    }
                } else {
                    var v71 string = "a"
                    _ = v71
                }
            default:
                for fi72 := 0; fi72 < 2; fi72++ {
                    if v51 {
                        v73 := GenS11{F0: 2, F1: 460, F2: (("hello" + "hello") + "a")}
                        _ = v73
                    }
                    var v74 float64 = (v2 - (v5 / 1.5))
                    v51 = (-1 >= -627)
                    _ = fi72
                    _ = v74
                }
            }
            var v75 string = "hello"
            _ = v75
        }
        switch -9223372036854775808 {
        case 9223372036854775807:
            if (!true) {
                if (v1 >= (^(v1))) {
                    for fi76 := 0; fi76 < 16; fi76++ {
                        var v77 int16 = 1
                        for fi78 := 0; fi78 < 5; fi78++ {
                            v1 = v1
                            fmt.Println(v2, v6)
                            _ = fi78
                        }
                        v6 += v6
                        _ = fi76
                        _ = v77
                    }
                }
                var v79 uint16 = 794
                var v80 int = 42
                _ = v79
                _ = v80
            } else {
                fmt.Println(v2)
                if false {
                    v2 = float64(((((1.5 + (0.1 - 0.0)) - 1.5) / (0.1 + float32(0.1))) + 0.0))
                } else {
                    v5 = v2
                    v7 = append(v7, 42)
                    v81 := GenS12{F0: -314, F1: ((!(0 == 0)) || (!(v51 || (2147483647 > -2147483648))))}
                    _ = v81
                }
            }
        default:
            switch "a" {
            case "hello":
                if v4 {
                    if (!v51) {
                        var v82 uint64 = v1
                        switch 42 {
                        case 4294967295:
                            fmt.Println(v5, v3, v4)
                        case 265:
                            fmt.Println(v51)
                        }
                        v7 = append(v7, 632)
                        _ = v82
                    }
                    switch -659 {
                    case 127:
                        switch ("a" + "a") {
                        case "":
                            var v83 uint32 = 378
                            _ = v83
                        case "hello":
                            if false {
                                for fi84 := 0; fi84 < 12; fi84++ {
                                    for fi85 := 0; fi85 < 16; fi85++ {
                                        v3 -= -66
                                        v5 -= 0.0
                                        _ = fi85
                                    }
                                    _ = fi84
                                }
                                for fi86 := 0; fi86 < 4; fi86++ {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    _ = fi86
                                }
                                var v87 string = "foo bar"
                                _ = v87
                            } else {
                                for fi88 := 0; fi88 < 16; fi88++ {
                                    var v89 int32 = 134
                                    var v90 int8 = (v3 << ((1) & 63))
                                    for fi91 := 0; fi91 < 14; fi91++ {
                                        _ = len(v7)
                                        _ = fi91
                                    }
                                    _ = fi88
                                    _ = v89
                                    _ = v90
                                }
                                v92 := GenS13{F0: (-(v3)), F1: (v6 - (^(v6))), F2: 1}
                                var v93 float32 = (3.14159 / float32(((v5 - 1.5) - v5)))
                                _ = v92
                                _ = v93
                            }
                        default:
                            v94 := GenS1{F0: 9223372036854775807, F1: -1, F2: 333}
                            _ = v94
                        }
                    case -814:
                        var v95 uint = 2
                        _ = v95
                    case -2147483648:
                        v7 = append(v7, 2147483647)
                    }
                }
            case "":
                var v96 int32 = -157
                _ = v96
            default:
                v1 = (v1 & (v1 ^ v1))
            }
        }
    } else {
        v6 += (v6 ^ v6)
        var v97 float32 = (0.1 + 0.0)
        if v4 {
            if len(v7) > 0 {
                v7[0] = 898
            }
            var v98 int8 = (v3 ^ -128)
            switch ("hello" + ((("a" + ("hello" + "\n\t")) + (("a" + "\n\t") + "hello")) + "hello")) {
            case "foo bar":
                v97 = v97
            case "a":
                for fi99 := 0; fi99 < 19; fi99++ {
                    v7 = append(v7, fi99)
                    for fi100 := 0; fi100 < 18; fi100++ {
                        v7 = append(v7, fi100)
                        var v101 uint8 = v6
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        fmt.Println(v5, v97)
                        var v102 uint64 = (v1 - v1)
                        var v103 int32 = 2147483647
                        _ = fi100
                        _ = v101
                        _ = v102
                        _ = v103
                    }
                    v104 := GenS6{F0: 42, F1: true, F2: 4294967295}
                    _ = fi99
                    _ = v104
                }
            case "":
                for fi105 := 0; fi105 < 7; fi105++ {
                    if (!v51) {
                        v4 = true
                    } else {
                        v98 += v98
                        var v106 uint = 65535
                        v7 = append(v7, (-(fi105)))
                        _ = v106
                    }
                    _ = fi105
                }
            default:
                if true {
                    for fi107 := 0; fi107 < 9; fi107++ {
                        v7 = append(v7, -9223372036854775808)
                        _ = fi107
                    }
                }
            }
            _ = v98
        }
        _ = v97
    }
    var v108 int8 = 47
    var v109 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v108
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v62
    _ = v7
    _ = v8
}
