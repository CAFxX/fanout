package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint8; F2 int8 }
type GenS5 struct { F0 int16; F1 int16; F2 int16 }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 int64; F1 int64; F2 int8 }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 int8; F1 bool; F2 int }

func main() {
    var v1 bool = true
    v2 := []string{}
    var v3 uint64 = 18446744073709551615
    v4 := []uint32{4294967295}
    v5 := make(map[string]string)
    v6 := GenS1{F0: 864, F1: "hello"}
    v4 = append(v4, 65535)
    if v1 {
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: -1, F1: ""}
            switch "foo bar" {
            case "":
                v9 := GenS1{F0: 1, F1: "a"}
                _ = v9
            case "a":
                v1 = (-346 >= 2147483647)
            case "hello":
                var v10 float32 = 3.14159
                _ = v10
            default:
                if false {
                    v4 = append(v4, 2)
                    if v1 {
                        v1 = v1
                        var v11 int16 = -816
                        v12 := v5[(((("foo bar" + "a") + "\n\t") + "a") + "")]
                        _ = v11
                        _ = v12
                    }
                    var v13 uint16 = 1
                    _ = v13
                } else {
                    v14 := GenS2{F0: ((fi7 == (^(fi7))) || v1), F1: 2147483648, F2: 736}
                    var v15 float64 = -2.25
                    _ = v14
                    _ = v15
                }
            }
            _ = fi7
            _ = v8
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            v2 = append(v2, "\n\t")
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := GenS2{F0: (!v1), F1: 574, F2: 32767}
                if false {
                    v1 = (-2147483648 < fi16)
                    fmt.Println(fi16, fi17)
                    var v19 int64 = 42
                    _ = v19
                }
                switch ("foo bar" + (("" + "foo bar") + "")) {
                case "\n\t":
                    v3 += (v3 - v3)
                case "":
                    if (!(-2.25 <= 0.1)) {
                        switch "" {
                        case "":
                            delete(v5, ("\n\t" + "foo bar"))
                        case "\n\t":
                            if (!v1) {
                                fmt.Println(v1, v3, fi16)
                            } else {
                                v1 = false
                                var v20 int64 = 0
                                _ = v20
                            }
                        }
                        v1 = (!(100 <= 1))
                        switch fi16 {
                        case -486:
                            fmt.Println(fi17)
                        case -2147483648:
                            v21 := GenS3{F0: (1.5 > ((1.5 / float64((1.5 - (1e10 / -2.25)))) - -2.25))}
                            _ = v21
                        default:
                            if (1000 < 967) {
                                switch 1 {
                                case -436:
                                    if true {
                                        v22 := GenS2{F0: (2 == 0), F1: 4294967295, F2: 65535}
                                        v1 = (0 >= 65535)
                                        v3 = 42
                                        _ = v22
                                    } else {
                                        v1 = v1
                                    }
                                case -1:
                                    v4 = append(v4, 42)
                                }
                                switch "foo bar" {
                                case "a":
                                    v1 = (2147483647 > 2)
                                    v23 := GenS3{F0: v1}
                                    var v24 uint64 = (v3 | 207)
                                    _ = v23
                                    _ = v24
                                default:
                                    var v25 int32 = 1
                                    _ = v25
                                }
                                var v26 uint32 = 2147483647
                                _ = v26
                            }
                        }
                    }
                }
                _ = fi17
                _ = v18
            }
            v27 := GenS4{F0: uint(v3), F1: 117, F2: int8(fi16)}
            _ = fi16
            _ = v27
        }
    }
    v28 := GenS5{F0: 255, F1: -32768, F2: 1000}
    if (v1 && v1) {
        v29 := GenS6{F0: v3, F1: "foo bar"}
        switch "foo bar" {
        case "foo bar":
            v30 := GenS6{F0: v3, F1: ("hello" + "\n\t")}
            _ = v30
        case "":
            var v31 float64 = 0.1
            _ = v31
        case "a":
            switch "a" {
            case "\n\t":
                fmt.Println(v3, v1)
            case "a":
                v1 = (2147483648 == 2)
            default:
                v3 = 42
            }
        default:
            if (v1 || false) {
                switch -108 {
                case 4294967295:
                    for fi32 := 0; fi32 < 1; fi32++ {
                        fmt.Println(v1, v3, v3)
                        _ = fi32
                    }
                case 730:
                    v1 = true
                case -838:
                    v3 = 1
                default:
                    var v33 uint8 = 0
                    _ = v33
                }
            } else {
                if v1 {
                    var v34 float64 = 0.1
                    _ = v34
                } else {
                    if v1 {
                        var v35 int8 = -1
                        var v36 int16 = 1000
                        _ = v35
                        _ = v36
                    } else {
                        v1 = true
                        for fi37 := 0; fi37 < 7; fi37++ {
                            switch "foo bar" {
                            case "":
                                fmt.Println(v3, v3)
                            case "hello":
                                if (42 < 190) {
                                    if (65535 < 917) {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v3 = (v3 & (v3 << ((0) & 63)))
                                        v38 := GenS7{F0: -1, F1: -9223372036854775808, F2: 100}
                                        _ = v38
                                    } else {
                                        v39 := GenS1{F0: 592, F1: "foo bar"}
                                        _ = v39
                                    }
                                    fmt.Println(fi37)
                                } else {
                                    v40 := GenS8{F0: v1, F1: 42}
                                    var v41 float32 = (0.0 + 1e10)
                                    fmt.Println(v3, v1, v3)
                                    _ = v40
                                    _ = v41
                                }
                            case "a":
                                v3 = v3
                            }
                            _ = fi37
                        }
                    }
                }
                if true {
                    switch -820 {
                    case 2147483647:
                        fmt.Println(v3, v3)
                    default:
                        var v42 uint64 = (v3 * 18446744073709551615)
                        _ = v42
                    }
                } else {
                    for fi43 := 0; fi43 < 9; fi43++ {
                        v1 = (42 > 2)
                        v3 = v3
                        v1 = (-992 > 0)
                        _ = fi43
                    }
                    var v44 string = "a"
                    _ = v44
                }
            }
        }
        switch "\n\t" {
        case "a":
            for fi45 := 0; fi45 < 11; fi45++ {
                v3 -= v3
                var v46 int8 = 100
                switch "foo bar" {
                case "\n\t":
                    var v47 int = fi45
                    _ = v47
                case "foo bar":
                    var v48 bool = true
                    _ = v48
                }
                _ = fi45
                _ = v46
            }
        case "":
            var v49 int8 = -2
            _ = v49
        case "foo bar":
            fmt.Println(v3, v3)
        }
        _ = v29
    }
    v50 := GenS9{F0: 1, F1: (((("a" + "a") + "foo bar") + "foo bar") != "hello"), F2: -33}
    delete(v5, "foo bar")
    if ((v3 / ((0) | 1)) == (v3 + v3)) {
        fmt.Println(v1, v3)
        v51 := GenS3{F0: v1}
        _ = v51
    }
    if (((3.14159 + 1e10) > (3.14159 - 3.14159)) && v1) {
        if v1 {
            fmt.Println(v1, v1)
        } else {
            var v52 int64 = -9223372036854775808
            v2 = append(v2, "a")
            _ = v52
        }
    }
    v53 := GenS5{F0: 438, F1: 42, F2: 1}
    if len(v4) > 0 {
        v4[0] = 2147483648
    }
    fmt.Println(v3)
    var v54 int16 = 771
    v1 = v1
    for fi55 := 0; fi55 < 14; fi55++ {
        v1 = true
        for fi56 := 0; fi56 < 4; fi56++ {
            v5[("" + "a")] = "\n\t"
            _ = fi56
        }
        fmt.Println(v1, v1, v1)
        _ = fi55
    }
    var v57 float64 = 3.14159
    for fi58 := 0; fi58 < 15; fi58++ {
        for fi59 := 0; fi59 < 5; fi59++ {
            var v60 uint8 = 1
            _ = fi59
            _ = v60
        }
        v3 += v3
        if (-1 <= 42) {
            var v61 int16 = v54
            fmt.Println(fi58)
            var v62 int64 = -692
            _ = v61
            _ = v62
        } else {
            if v1 {
                var v63 int32 = -2147483648
                _ = v63
            }
            fmt.Println(v54, v57, v3)
            var v64 int64 = -603
            _ = v64
        }
        _ = fi58
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v53
    _ = v54
    _ = v57
    _ = v6
}
