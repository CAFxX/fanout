package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int16; F1 string; F2 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-2147483648, 127, 573, 42}
    var v3 string = ""
    v4 := []string{(v3 + ""), (v3 + ("a" + (v3 + ((v3 + v3) + "")))), "foo bar", ""}
    v5 := []int{32767, 128, 2147483647}
    if false {
        if (52 >= 2) {
            var v6 float32 = (float32(0.1) * ((3.14159 - 0.0) / 1.0))
            var v7 uint16 = 255
            switch "hello" {
            case "\n\t":
                if true {
                    var v8 int64 = 1
                    v9 := GenS1{F0: -1, F1: (-1 != 0), F2: 0}
                    if ("a" >= "foo bar") {
                        if ((1.5 < -2.25) && false) {
                            delete(v1, 2147483648)
                            if true {
                                for fi10 := 0; fi10 < 8; fi10++ {
                                    if len(v4) > 0 {
                                        v4[0] = "foo bar"
                                    }
                                    v11 := v1[42]
                                    var v12 int = v11
                                    _ = fi10
                                    _ = v11
                                    _ = v12
                                }
                                v7 = v7
                                if (-833 == 9223372036854775807) {
                                    v13 := GenS1{F0: -218, F1: true, F2: 1}
                                    _ = v13
                                } else {
                                    var v14 string = v3
                                    for fi15 := 0; fi15 < 15; fi15++ {
                                        v7 = v7
                                        v16 := GenS1{F0: 0, F1: false, F2: 42}
                                        v4 = append(v4, "a")
                                        _ = fi15
                                        _ = v16
                                    }
                                    var v17 bool = true
                                    _ = v14
                                    _ = v17
                                }
                            } else {
                                v18 := GenS1{F0: 42, F1: (72 <= -155), F2: 1}
                                if true {
                                    v19 := GenS1{F0: int32(v7), F1: true, F2: 98}
                                    switch v3 {
                                    case "a":
                                        if (!(v8 >= v8)) {
                                            v20 := GenS2{F0: 1, F1: (v7 >> ((714) & 63))}
                                            var v21 uint16 = 996
                                            switch "a" {
                                            case "foo bar":
                                                v1[1] = 32767
                                            case "a":
                                                var v22 uint32 = 868
                                                _ = v22
                                            case "\n\t":
                                                for fi23 := 0; fi23 < 15; fi23++ {
                                                    v21 = v7
                                                    v21 = v7
                                                    _ = fi23
                                                }
                                            default:
                                                switch ((v3 + "") + "hello") {
                                                case "\n\t":
                                                    v2 = append(v2, -326)
                                                case "a":
                                                    for fi24 := 0; fi24 < 3; fi24++ {
                                                        var v25 float32 = v6
                                                        _ = fi24
                                                        _ = v25
                                                    }
                                                default:
                                                    v7 = 863
                                                }
                                            }
                                            _ = v20
                                            _ = v21
                                        } else {
                                            v3 = v3
                                        }
                                    default:
                                        v3 = v3
                                    }
                                    _ = v19
                                }
                                var v26 int = 128
                                _ = v18
                                _ = v26
                            }
                        } else {
                            v7 += v7
                            fmt.Println(v7, v8)
                            v7 += v7
                        }
                        if true {
                            if ((!(false && ((v3 < "\n\t") && (float64(1e10) <= 1.5)))) && true) {
                                _ = len(v1)
                                for fi27 := 0; fi27 < 9; fi27++ {
                                    if (fi27 != 2147483647) {
                                        for fi28 := 0; fi28 < 16; fi28++ {
                                            var v29 uint16 = 2
                                            switch 419 {
                                            case 128:
                                                var v30 uint16 = v29
                                                _ = v30
                                            case 1000:
                                                fmt.Println(v6, v3, v7)
                                            default:
                                                var v31 uint16 = (^(v7))
                                                _ = v31
                                            }
                                            v8 = (v8 / ((0) | 1))
                                            _ = fi28
                                            _ = v29
                                        }
                                        if false {
                                            var v32 int64 = (^(v8))
                                            _ = v32
                                        } else {
                                            v3 = ((v3 + ((("" + v3) + "foo bar") + v3)) + "foo bar")
                                        }
                                    } else {
                                        if true {
                                            fmt.Println(fi27)
                                            for fi33 := 0; fi33 < 18; fi33++ {
                                                switch v3 {
                                                case "":
                                                    v6 -= v6
                                                case "\n\t":
                                                    if (89 < 100) {
                                                        var v34 uint32 = 2147483647
                                                        var v35 int16 = 1000
                                                        var v36 float64 = 1e10
                                                        _ = v34
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                default:
                                                    v3 = v3
                                                }
                                                v6 += float32(1.5)
                                                _ = fi33
                                            }
                                            v7 = v7
                                        }
                                    }
                                    _ = fi27
                                }
                                if true {
                                    var v37 uint8 = 128
                                    _ = v37
                                }
                            }
                            switch v3 {
                            case "\n\t":
                                var v38 int = 4294967295
                                _ = v38
                            case "":
                                if (v6 == (float32(float64(v6)) - 1.5)) {
                                    v39 := GenS3{F0: 100, F1: -1}
                                    v40 := GenS1{F0: 2147483647, F1: false, F2: 1}
                                    _ = v39
                                    _ = v40
                                }
                            case "a":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        }
                        var v41 int64 = 518
                        _ = v41
                    } else {
                        var v42 string = ("\n\t" + v3)
                        if (!true) {
                            var v43 string = v3
                            for fi44 := 0; fi44 < 12; fi44++ {
                                if (!(19 <= (-(fi44)))) {
                                    var v45 uint64 = 2
                                    if len(v5) > 0 {
                                        v5[0] = 1
                                    }
                                    v3 = v43
                                    _ = v45
                                }
                                var v46 bool = true
                                var v47 uint8 = 42
                                _ = fi44
                                _ = v46
                                _ = v47
                            }
                            switch -818 {
                            case 341:
                                fmt.Println(v3, v6, v6)
                            default:
                                for fi48 := 0; fi48 < 18; fi48++ {
                                    if false {
                                        var v49 uint = 1
                                        v50 := GenS4{F0: 42, F1: (("foo bar" + ("hello" + (v42 + (v42 + v43)))) + v43), F2: (^(fi48))}
                                        _ = v49
                                        _ = v50
                                    } else {
                                        v6 -= (0.0 + (float32(((-2.25 - (0.1 + float64(v6))) + (1e10 / float64(((v6 * 3.14159) - (float32(3.14159) - float32(1.5))))))) - (v6 * 0.1)))
                                        if (!true) {
                                            var v51 string = "hello"
                                            _ = v51
                                        } else {
                                            v52 := GenS1{F0: 2147483647, F1: true, F2: 0}
                                            _ = v52
                                        }
                                    }
                                    _ = fi48
                                }
                            }
                            _ = v43
                        } else {
                            if len(v4) > 0 {
                                v4[0] = "foo bar"
                            }
                        }
                        _ = v42
                    }
                    _ = v8
                    _ = v9
                }
            }
            _ = v6
            _ = v7
        } else {
            v3 = v3
            var v53 bool = (!(943 != 931))
            _ = v53
        }
        if (false && true) {
            if (3.14159 != 0.1) {
                v3 = ((v3 + v3) + (("a" + v3) + (v3 + (v3 + "foo bar"))))
            } else {
                _ = len(v1)
            }
            if len(v5) > 0 {
                v5[0] = 42
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            if (true || true) {
                v3 = v3
            } else {
                v3 = v3
                fmt.Println(v3)
                v3 = "hello"
            }
        }
        for fi54 := 0; fi54 < 12; fi54++ {
            v5 = append(v5, (fi54 - (fi54 - (fi54 << ((2) & 63)))))
            switch fi54 {
            case 487:
                switch (v3 + v3) {
                case "\n\t":
                    var v55 uint = 255
                    _ = v55
                case "foo bar":
                    v56 := GenS3{F0: -100, F1: 613}
                    _ = v56
                default:
                    delete(v1, 2147483648)
                }
            case -57:
                var v57 int = (fi54 / ((fi54) | 1))
                _ = v57
            }
            if false {
                var v58 int8 = 1
                var v59 int32 = 42
                _ = v58
                _ = v59
            } else {
                if (!false) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v3 = v3
                    if true {
                        for fi60 := 0; fi60 < 14; fi60++ {
                            v3 = v3
                            _ = fi60
                        }
                        var v61 uint8 = 42
                        _ = v61
                    }
                }
            }
            _ = fi54
        }
    }
    for fi62 := 0; fi62 < 10; fi62++ {
        switch v3 {
        case "":
            var v63 int = (fi62 & fi62)
            _ = v63
        }
        _ = fi62
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi64 := 0; fi64 < 15; fi64++ {
        v3 = ("\n\t" + v3)
        v3 = "hello"
        if (-1 >= -2147483648) {
            if (4294967295 > 0) {
                var v65 uint = 0
                var v66 uint = v65
                _ = v65
                _ = v66
            } else {
                v67 := GenS2{F0: 2, F1: 816}
                for fi68 := 0; fi68 < 12; fi68++ {
                    fmt.Println(fi64, v3, v3)
                    fmt.Println(v3, v3)
                    v5 = append(v5, fi68)
                    _ = fi68
                }
                _ = v67
            }
        }
        _ = fi64
    }
    switch v3 {
    case "":
        v3 = "\n\t"
    case "a":
        v69 := GenS1{F0: 0, F1: (!(!true)), F2: 34}
        _ = v69
    }
    v70 := GenS3{F0: 100, F1: -1}
    var v71 bool = true
    fmt.Println(v3)
    switch ("a" + ("foo bar" + (v3 + v3))) {
    case "":
        var v72 bool = v71
        _ = v72
    }
    v2 = append(v2, 188)
    var v73 int32 = 2147483647
    switch -609 {
    case 32767:
        v73 = (^(v73))
    case 128:
        switch v3 {
        case "foo bar":
            v74 := GenS5{F0: -99}
            _ = v74
        case "a":
            fmt.Println(v73, v73, v3)
        default:
            v75 := GenS4{F0: -32768, F1: v3, F2: -185}
            _ = v75
        }
    case -2147483648:
        if (!(v71 && false)) {
            if false {
                delete(v1, 9)
                v76 := GenS4{F0: -1, F1: "\n\t", F2: -677}
                _ = v76
            } else {
                var v77 float64 = (((0.0 * 1.5) / -2.25) * float64(3.14159))
                if true {
                    for fi78 := 0; fi78 < 9; fi78++ {
                        var v79 int64 = 268
                        delete(v1, 242)
                        _ = fi78
                        _ = v79
                    }
                    for fi80 := 0; fi80 < 16; fi80++ {
                        v81 := GenS6{F0: 1, F1: v71}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v82 int16 = -192
                        _ = fi80
                        _ = v81
                        _ = v82
                    }
                    for fi83 := 0; fi83 < 17; fi83++ {
                        switch fi83 {
                        case 2147483647:
                            fmt.Println(v73, v77)
                        case 0:
                            var v84 int = (fi83 + 566)
                            _ = v84
                        case -1:
                            var v85 int16 = 42
                            _ = v85
                        }
                        var v86 int = 2147483647
                        switch fi83 {
                        case 1:
                            if v71 {
                                for fi87 := 0; fi87 < 11; fi87++ {
                                    v1[2] = 42
                                    _ = fi87
                                }
                                for fi88 := 0; fi88 < 3; fi88++ {
                                    v89 := GenS5{F0: 87}
                                    switch ("" + "") {
                                    case "":
                                        switch fi88 {
                                        case -9223372036854775808:
                                            switch fi88 {
                                            case -9223372036854775808:
                                                var v90 int32 = (v73 << ((386) & 63))
                                                _ = v90
                                            case 32768:
                                                fmt.Println(v71, v3)
                                            case 9223372036854775807:
                                                fmt.Println(v71)
                                            default:
                                                fmt.Println(v71, v3)
                                            }
                                        }
                                    default:
                                        var v91 float64 = v77
                                        _ = v91
                                    }
                                    v92 := GenS2{F0: 127, F1: 478}
                                    _ = fi88
                                    _ = v89
                                    _ = v92
                                }
                                v93 := GenS2{F0: 128, F1: 65535}
                                _ = v93
                            }
                        default:
                            v94 := GenS3{F0: -128, F1: (v73 % (((-(v73))) | 1))}
                            _ = v94
                        }
                        _ = fi83
                        _ = v86
                    }
                } else {
                    v1[531] = 32767
                }
                _ = v77
            }
            var v95 bool = ((1e10 * 3.14159) == 3.14159)
            v96 := GenS3{F0: -128, F1: (v73 << ((181) & 63))}
            _ = v95
            _ = v96
        } else {
            v71 = (1e10 < (1.5 + 1.5))
            if v71 {
                fmt.Println(v73, v71, v71)
                v71 = ((!v71) && false)
                var v97 float64 = 1.5
                _ = v97
            }
            v3 = "hello"
        }
    }
    if (!v71) {
        v98 := GenS7{F0: 646}
        _ = v98
    } else {
        fmt.Println(v71)
    }
    v99 := GenS4{F0: 0, F1: v3, F2: 127}
    v73 -= v73
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v71
    _ = v73
    _ = v99
}
