package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int32; F1 bool }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint32; F1 bool }

func main() {
    v1 := []int{128}
    v2 := []string{("hello" + "hello"), "foo bar", ("\n\t" + "a"), ("a" + ("" + ("a" + ("hello" + "hello"))))}
    var v3 uint64 = 42
    v4 := make(map[int]uint32)
    switch "" {
    case "foo bar":
        if true {
            v5 := GenS1{F0: 42, F1: 0, F2: 127}
            var v6 uint16 = 401
            if (!false) {
                if false {
                    for fi7 := 0; fi7 < 13; fi7++ {
                        _ = len(v4)
                        switch "foo bar" {
                        case "":
                            fmt.Println(v6, v3, v6)
                        case "a":
                            if true {
                                if (34 >= 255) {
                                    v8 := GenS1{F0: fi7, F1: 4294967295, F2: 42}
                                    _ = v8
                                } else {
                                    v6 = v6
                                    v6 -= (v6 << ((42) & 63))
                                }
                                v9 := GenS1{F0: fi7, F1: v3, F2: 1}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v9
                            } else {
                                v3 -= v3
                                switch fi7 {
                                case 290:
                                    v3 = (v3 & (v3 & (v3 % ((v3) | 1))))
                                case 563:
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v6)
                                    }
                                }
                                for fi10 := 0; fi10 < 9; fi10++ {
                                    var v11 uint32 = 594
                                    _ = fi10
                                    _ = v11
                                }
                            }
                        }
                        var v12 bool = false
                        _ = fi7
                        _ = v12
                    }
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                    var v13 int16 = 32767
                    _ = v13
                }
            }
            _ = v5
            _ = v6
        }
    case "\n\t":
        v4[127] = 42
    default:
        var v14 float64 = 0.0
        _ = v14
    }
    if (!(1.5 >= ((1e10 * 0.0) * 3.14159))) {
        var v15 uint64 = (^(v3))
        if false {
            v15 = (v15 ^ v3)
            var v16 int64 = 1
            _ = v16
        } else {
            v2 = append(v2, "a")
            for fi17 := 0; fi17 < 17; fi17++ {
                if true {
                    var v18 int16 = 284
                    v3 = v15
                    var v19 string = "\n\t"
                    _ = v18
                    _ = v19
                } else {
                    fmt.Println(v3)
                    if (true || true) {
                        for fi20 := 0; fi20 < 5; fi20++ {
                            v15 -= v3
                            if ((-1 != 775) && (true || true)) {
                                var v21 bool = (-491 <= 1)
                                _ = v21
                            } else {
                                v22 := GenS2{F0: 844}
                                if (30 > 42) {
                                    var v23 uint32 = 2147483647
                                    _ = v23
                                }
                                var v24 int = fi17
                                _ = v22
                                _ = v24
                            }
                            var v25 bool = ((973 != 42) || false)
                            var v26 uint = 42
                            if (42 == 4294967295) {
                                switch fi17 {
                                case 9223372036854775807:
                                    v4[(-(fi20))] = 0
                                    switch fi20 {
                                    case 62:
                                        _ = len(v4)
                                    case -103:
                                        var v27 int8 = 1
                                        _ = v27
                                    case 4294967295:
                                        v4[fi20] = 148
                                    }
                                    var v28 int16 = 1000
                                    _ = v28
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v26 += (^(v26))
                            }
                            _ = fi20
                            _ = v25
                            _ = v26
                        }
                        for fi29 := 0; fi29 < 6; fi29++ {
                            v30 := GenS2{F0: 133}
                            v31 := v4[(fi29 >> ((793) & 63))]
                            if false {
                                var v32 uint64 = v15
                                fmt.Println(v32, v32, v3)
                                var v33 float64 = (3.14159 + 3.14159)
                                delete(v4, fi17)
                                switch fi29 {
                                case 42:
                                    v34 := GenS3{F0: 168, F1: 42}
                                    _ = v34
                                case -111:
                                    v35 := GenS4{F0: 255, F1: 103, F2: -1}
                                    _ = v35
                                default:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                _ = v32
                                _ = v33
                            } else {
                                fmt.Println(v3)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if (-1 == -34) {
                                    if (!((!true) && (200 <= 255))) {
                                        fmt.Println(fi29, v31, fi17)
                                        var v36 int32 = -1
                                        _ = v36
                                    }
                                }
                            }
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    } else {
                        var v37 uint32 = 586
                        _ = v37
                    }
                    var v38 bool = true
                    _ = v38
                }
                v39 := GenS3{F0: 1, F1: 1}
                for fi40 := 0; fi40 < 2; fi40++ {
                    switch "foo bar" {
                    case "foo bar":
                        var v41 int32 = 42
                        _ = v41
                    case "":
                        v15 = v15
                    }
                    v42 := GenS5{F0: 1}
                    _ = len(v2)
                    _ = fi40
                    _ = v42
                }
                _ = fi17
                _ = v39
            }
        }
        var v43 bool = ((!(0 <= -1)) || (true || (((0.0 + 1.5) - (0.0 * 3.14159)) == 0.0)))
        _ = v15
        _ = v43
    } else {
        fmt.Println(v3)
        v44 := GenS6{F0: 32767}
        for fi45 := 0; fi45 < 9; fi45++ {
            for fi46 := 0; fi46 < 10; fi46++ {
                var v47 uint64 = v3
                var v48 int8 = -123
                var v49 int = 2147483647
                _ = fi46
                _ = v47
                _ = v48
                _ = v49
            }
            var v50 uint32 = 0
            _ = fi45
            _ = v50
        }
        _ = v44
    }
    var v51 float32 = 0.1
    fmt.Println(v51, v3)
    v4[488] = 4294967295
    fmt.Println(v3, v3)
    switch "a" {
    case "a":
        if (!(42 > 1)) {
            v3 = (v3 >> ((260) & 63))
            v51 -= (v51 * (v51 / float32(1e10)))
            var v52 string = "\n\t"
            _ = v52
        } else {
            if false {
                v53 := GenS6{F0: 1}
                _ = len(v1)
                _ = v53
            } else {
                v1 = append(v1, 9223372036854775807)
                if false {
                    switch -2147483648 {
                    case -9223372036854775808:
                        v51 = v51
                    default:
                        var v54 int64 = -202
                        _ = v54
                    }
                    v55 := GenS7{F0: 255}
                    switch 34 {
                    case -771:
                        var v56 uint = 42
                        _ = v56
                    case 597:
                        v57 := GenS8{F0: -2147483648, F1: (!false)}
                        _ = v57
                    default:
                        for fi58 := 0; fi58 < 18; fi58++ {
                            for fi59 := 0; fi59 < 13; fi59++ {
                                if (109 >= 255) {
                                    v51 += (v51 * 0.1)
                                    fmt.Println(v3, fi59, fi58)
                                } else {
                                    v3 += (v3 + (^(v3)))
                                    v51 = (float32(0.1) - (v51 - ((1e10 - v51) * v51)))
                                    fmt.Println(fi58, fi59)
                                }
                                _ = fi59
                            }
                            fmt.Println(v3, v3, v51)
                            _ = fi58
                        }
                    }
                    _ = v55
                } else {
                    v60 := GenS1{F0: 2147483647, F1: v3, F2: -1}
                    v51 -= v51
                    if (2147483648 > 42) {
                        for fi61 := 0; fi61 < 12; fi61++ {
                            var v62 uint8 = 127
                            switch 42 {
                            case -1:
                                v62 = 42
                            case 1:
                                var v63 int8 = -84
                                _ = v63
                            default:
                                switch (fi61 & fi61) {
                                case 0:
                                    fmt.Println(v51)
                                case 32768:
                                    v51 -= (v51 * 1.5)
                                }
                            }
                            switch "a" {
                            case "hello":
                                v64 := GenS9{F0: 235}
                                _ = v64
                            case "\n\t":
                                var v65 uint16 = 65535
                                _ = v65
                            default:
                                delete(v4, 128)
                            }
                            _ = fi61
                            _ = v62
                        }
                    } else {
                        for fi66 := 0; fi66 < 4; fi66++ {
                            var v67 uint = 4294967295
                            _ = fi66
                            _ = v67
                        }
                        delete(v4, -752)
                        var v68 uint8 = 128
                        _ = v68
                    }
                    _ = v60
                }
                var v69 int32 = 42
                _ = v69
            }
            v70 := v4[127]
            v70 -= uint32(v3)
            _ = v70
        }
    }
    switch ("a" + "hello") {
    case "foo bar":
        if false {
            v3 = v3
            v3 += v3
            var v71 uint16 = 255
            _ = v71
        } else {
            for fi72 := 0; fi72 < 8; fi72++ {
                var v73 int32 = 486
                v74 := v4[4294967295]
                _ = fi72
                _ = v73
                _ = v74
            }
            switch -1 {
            case 127:
                for fi75 := 0; fi75 < 12; fi75++ {
                    var v76 bool = true
                    _ = fi75
                    _ = v76
                }
            case -2147483648:
                for fi77 := 0; fi77 < 5; fi77++ {
                    switch "a" {
                    case "a":
                        fmt.Println(v51, fi77)
                    default:
                        v78 := GenS10{F0: 296, F1: (!((("foo bar" + "hello") + "a") != "\n\t"))}
                        _ = v78
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if ((v51 + (v51 * 0.1)) > -2.25) {
                        v3 -= 42
                        v3 -= (v3 * (^(v3)))
                    }
                    _ = fi77
                }
            default:
                for fi79 := 0; fi79 < 11; fi79++ {
                    if false {
                        for fi80 := 0; fi80 < 14; fi80++ {
                            _ = len(v4)
                            var v81 uint32 = 2147483647
                            _ = fi80
                            _ = v81
                        }
                    }
                    _ = fi79
                }
            }
            for fi82 := 0; fi82 < 7; fi82++ {
                var v83 uint64 = v3
                _ = fi82
                _ = v83
            }
        }
    case "hello":
        if (!(383 > 866)) {
            var v84 bool = (-964 == 376)
            v84 = v84
            for fi85 := 0; fi85 < 17; fi85++ {
                v84 = (914 > 65535)
                _ = fi85
            }
            _ = v84
        }
    default:
        v51 -= v51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
}
