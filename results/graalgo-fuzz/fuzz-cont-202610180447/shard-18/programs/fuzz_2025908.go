package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint; F1 int16; F2 uint32 }
type GenS3 struct { F0 int32; F1 int64; F2 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16; F1 int8; F2 uint }
type GenS10 struct { F0 bool }

func main() {
    var v1 float32 = (((-2.25 + (1.5 / 1.0)) + -2.25) / 1.0)
    var v2 bool = (-86 == 0)
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]uint32)
    v5 := make(map[int]int)
    v6 := []string{}
    v7 := GenS1{F0: 32767, F1: -1}
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: 1, F1: 42}
        v1 = float32(1e10)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        switch (("a" + (("hello" + ("\n\t" + "a")) + "a")) + "\n\t") {
        case "\n\t":
            var v11 uint16 = 2
            _ = v11
        case "":
            _ = len(v4)
        case "hello":
            if (!(!(1 <= 896))) {
                for fi12 := 0; fi12 < 18; fi12++ {
                    _ = len(v6)
                    fmt.Println(v2, fi10, v1)
                    v13 := GenS1{F0: -951, F1: 11}
                    var v14 uint32 = 336
                    var v15 uint32 = 545
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                var v16 uint16 = 291
                _ = v16
            } else {
                v17 := GenS1{F0: 0, F1: -89}
                _ = v17
            }
        default:
            v1 -= (v1 + (v1 + ((-2.25 / float32(float64(v1))) + float32(3.14159))))
        }
        _ = fi10
    }
    v1 += v1
    if v2 {
        var v18 int = 151
        var v19 uint = 255
        if false {
            fmt.Println(v2)
            fmt.Println(v19, v3, v3)
            for fi20 := 0; fi20 < 3; fi20++ {
                if false {
                    switch v18 {
                    case 32767:
                        v21 := GenS2{F0: (v19 & v19), F1: 1000, F2: 598}
                        _ = v21
                    default:
                        var v22 int16 = -1
                        _ = v22
                    }
                    var v23 uint32 = 2147483648
                    _ = v23
                }
                if v2 {
                    var v24 uint32 = uint32(v19)
                    var v25 string = (("foo bar" + "hello") + "hello")
                    _ = v24
                    _ = v25
                } else {
                    v1 = (-2.25 * (((v1 + ((float32((float64(float32(0.1)) - 3.14159)) / v1) / 3.14159)) + (float32(-2.25) / v1)) / float32(3.14159)))
                    v26 := GenS2{F0: v19, F1: 0, F2: 2147483647}
                    if v2 {
                        switch "foo bar" {
                        case "a":
                            switch (("\n\t" + "a") + (((("a" + "a") + ("a" + "a")) + ("\n\t" + ("hello" + "a"))) + "hello")) {
                            case "foo bar":
                                var v27 float64 = 1e10
                                _ = v27
                            case "a":
                                if (-139 == 2147483647) {
                                    var v28 bool = (0 > 128)
                                    v28 = true
                                    var v29 int8 = 100
                                    _ = v28
                                    _ = v29
                                } else {
                                    if (v2 && (-2147483648 <= 0)) {
                                        _ = len(v5)
                                        switch 32767 {
                                        case -1:
                                            var v30 int16 = 0
                                            _ = v30
                                        case 42:
                                            for fi31 := 0; fi31 < 6; fi31++ {
                                                if (v2 && v2) {
                                                    v32 := GenS3{F0: -2147483648, F1: -9223372036854775808, F2: (^(v3))}
                                                    v4[v18] = 2147483648
                                                    _ = v32
                                                }
                                                var v33 int = (^(v18))
                                                v34 := GenS4{F0: 0}
                                                _ = fi31
                                                _ = v33
                                                _ = v34
                                            }
                                        case 337:
                                            if (42 != 18446744073709551615) {
                                                v1 = (1.5 + (v1 + v1))
                                                var v35 int32 = 153
                                                _ = v35
                                            } else {
                                                if (222 < 128) {
                                                    v36 := GenS1{F0: 1000, F1: -1}
                                                    _ = v36
                                                } else {
                                                    v37 := GenS5{F0: (!v2)}
                                                    v3 = v3
                                                    v19 += (v19 | (v19 - v19))
                                                    _ = v37
                                                }
                                                for fi38 := 0; fi38 < 15; fi38++ {
                                                    if false {
                                                        var v39 int64 = v3
                                                        _ = len(v6)
                                                        v40 := GenS1{F0: 1, F1: -31}
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                    v4[(fi38 ^ v18)] = 2147483648
                                                    var v41 int8 = -128
                                                    _ = fi38
                                                    _ = v41
                                                }
                                                switch fi20 {
                                                case -9223372036854775808:
                                                    if (v2 && v2) {
                                                        var v42 int64 = v3
                                                        v3 -= (^(v42))
                                                        switch (v18 >> ((405) & 63)) {
                                                        case 127:
                                                            v43 := GenS3{F0: 0, F1: (v3 + v42), F2: v42}
                                                            _ = v43
                                                        case 433:
                                                            fmt.Println(v19, v3, v18)
                                                        case -908:
                                                            v19 = 255
                                                        default:
                                                            if (-1 > 42) {
                                                                if v2 {
                                                                    if v2 {
                                                                        v1 += v1
                                                                    }
                                                                    var v44 float32 = ((0.1 + 0.1) + 3.14159)
                                                                    _ = v44
                                                                } else {
                                                                    v18 += (v18 << ((v19) & 63))
                                                                }
                                                                v45 := GenS6{F0: 1}
                                                                if v2 {
                                                                    switch ("" + "") {
                                                                    case "":
                                                                        var v46 int16 = -498
                                                                        _ = v46
                                                                    default:
                                                                        v5[-765] = -1
                                                                    }
                                                                    v19 = (v19 & (^(v19)))
                                                                } else {
                                                                    v3 -= (v42 / ((423) | 1))
                                                                    if (18446744073709551615 < v19) {
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                        v47 := GenS2{F0: v19, F1: 1000, F2: 48}
                                                                        _ = v47
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        if false {
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        if v2 {
                                                                            var v48 int32 = 2147483647
                                                                            for fi49 := 0; fi49 < 11; fi49++ {
                                                                                if ((v19 * 2) != (v19 - v19)) {
                                                                                    v18 = -844
                                                                                    if len(v6) > 0 {
                                                                                        v6[0] = "a"
                                                                                    }
                                                                                }
                                                                                _ = fi49
                                                                            }
                                                                            _ = v48
                                                                        } else {
                                                                            v3 -= (v42 * 1)
                                                                            var v50 string = "a"
                                                                            var v51 int64 = (v3 ^ (v3 << ((v19) & 63)))
                                                                            _ = v50
                                                                            _ = v51
                                                                        }
                                                                    default:
                                                                        v2 = false
                                                                    }
                                                                }
                                                                _ = v45
                                                            }
                                                        }
                                                        _ = v42
                                                    }
                                                default:
                                                    for fi52 := 0; fi52 < 4; fi52++ {
                                                        v53 := v4[(fi52 ^ -9223372036854775808)]
                                                        var v54 float32 = v1
                                                        var v55 uint64 = 0
                                                        _ = fi52
                                                        _ = v53
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                }
                                            }
                                        }
                                    } else {
                                        var v56 int32 = -2147483648
                                        var v57 int32 = v56
                                        _ = v56
                                        _ = v57
                                    }
                                    var v58 uint8 = 255
                                    for fi59 := 0; fi59 < 15; fi59++ {
                                        if len(v6) > 0 {
                                            v6[0] = (("a" + "\n\t") + "foo bar")
                                        }
                                        v60 := GenS4{F0: 4294967295}
                                        _ = fi59
                                        _ = v60
                                    }
                                    _ = v58
                                }
                            }
                        }
                    }
                    _ = v26
                }
                _ = fi20
            }
        } else {
            for fi61 := 0; fi61 < 13; fi61++ {
                v2 = (((^(fi61)) > fi61) || (v2 && false))
                if true {
                    if true {
                        v3 = v3
                        v3 = v3
                    } else {
                        var v62 int16 = 1
                        var v63 uint8 = 182
                        _ = len(v4)
                        _ = v62
                        _ = v63
                    }
                    v64 := GenS7{F0: (^(v19)), F1: 804}
                    _ = v64
                } else {
                    if len(v6) > 0 {
                        v6[0] = "foo bar"
                    }
                    var v65 uint8 = 42
                    v65 -= 250
                    _ = v65
                }
                _ = fi61
            }
        }
        _ = v18
        _ = v19
    }
    v66 := GenS8{F0: 1}
    if false {
        v1 = 3.14159
    }
    if true {
        _ = len(v4)
        for fi67 := 0; fi67 < 19; fi67++ {
            for fi68 := 0; fi68 < 2; fi68++ {
                v2 = (false && (0 == 2))
                var v69 uint32 = 0
                v70 := v4[(fi68 | fi68)]
                _ = fi68
                _ = v69
                _ = v70
            }
            for fi71 := 0; fi71 < 11; fi71++ {
                v1 = 0.1
                v4[(fi71 / ((fi67) | 1))] = 0
                if false {
                    var v72 int32 = 1
                    _ = v72
                } else {
                    fmt.Println(v3, fi67, v2)
                }
                fmt.Println(fi71, fi67, v2)
                v3 = -9223372036854775808
                v3 += (v3 / ((9223372036854775807) | 1))
                var v73 float32 = (v1 / v1)
                v74 := GenS2{F0: 884, F1: 0, F2: 95}
                var v75 int32 = 831
                _ = fi71
                _ = v73
                _ = v74
                _ = v75
            }
            v76 := v4[fi67]
            _ = fi67
            _ = v76
        }
    }
    if v2 {
        if v2 {
            v1 = (((3.14159 * v1) * (v1 + float32((float64((float32(1e10) * 1.5)) + 0.0)))) - -2.25)
        } else {
            switch ("\n\t" + ("foo bar" + "foo bar")) {
            case "a":
                var v77 int8 = 42
                _ = v77
            }
            var v78 uint64 = 106
            v78 += 4294967295
            _ = v78
        }
        for fi79 := 0; fi79 < 6; fi79++ {
            if ((^(fi79)) <= 32768) {
                v80 := GenS5{F0: true}
                if (!v2) {
                    var v81 int = fi79
                    v82 := GenS4{F0: 1}
                    var v83 float32 = (1e10 / v1)
                    _ = v81
                    _ = v82
                    _ = v83
                }
                _ = v80
            }
            _ = fi79
        }
    }
    v84 := GenS9{F0: 0, F1: -128, F2: 1}
    v85 := GenS8{F0: 32767}
    var v86 uint64 = 1
    v87 := GenS10{F0: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v7
    _ = v84
    _ = v85
    _ = v86
    _ = v87
}
