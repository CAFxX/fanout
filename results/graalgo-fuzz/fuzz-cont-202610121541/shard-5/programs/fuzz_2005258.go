package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint16; F1 uint32; F2 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 int32; F1 int32 }
func genFn0() uint64 {
    var v7 int64 = 1
    var v8 string = "hello"
    if false {
        fmt.Println(v7)
        var v9 uint8 = 128
        _ = v9
    }
    if (-833 != 1) {
        v7 += (^(v7))
        var v10 uint64 = 1
        v7 += (v7 & v7)
        _ = v10
    } else {
        v11 := GenS1{F0: 1}
        for fi12 := 0; fi12 < 4; fi12++ {
            if (2 >= 1) {
                var v13 uint64 = 160
                _ = v13
            } else {
                var v14 string = ("foo bar" + v8)
                _ = v14
            }
            _ = fi12
        }
        switch 9223372036854775807 {
        case -989:
            switch v8 {
            case "a":
                v8 = ((v8 + ("foo bar" + "\n\t")) + v8)
            }
        case 4294967295:
            if ((v7 / ((v7) | 1)) <= (v7 + -1)) {
                var v15 uint16 = 641
                fmt.Println(v15)
                _ = v15
            }
        case 9223372036854775807:
            if (932 <= 4294967295) {
                if false {
                    v16 := GenS1{F0: 744}
                    v17 := GenS1{F0: 2}
                    _ = v16
                    _ = v17
                } else {
                    switch -917 {
                    case 4294967295:
                        v7 += (v7 & -1)
                    case 650:
                        v18 := GenS1{F0: 232}
                        _ = v18
                    default:
                        if (!false) {
                            v7 = 1
                            for fi19 := 0; fi19 < 2; fi19++ {
                                switch v8 {
                                case "hello":
                                    v20 := GenS1{F0: 1}
                                    _ = v20
                                case "\n\t":
                                    if ((((1.5 - ((1e10 * float64(((0.1 - 0.1) + 0.0))) + 0.0)) - 1e10) - (((((float64(0.0) + 3.14159) - float64(0.0)) * 3.14159) - -2.25) + (((-2.25 / 1.0) / (((1e10 + 1.5) - 3.14159) + float64(0.0))) - 3.14159))) < 0.0) {
                                        fmt.Println(v8, v7, fi19)
                                    }
                                case "foo bar":
                                    for fi21 := 0; fi21 < 19; fi21++ {
                                        if (!false) {
                                            v22 := GenS1{F0: 4294967295}
                                            switch "a" {
                                            case "\n\t":
                                                v23 := GenS1{F0: 4294967295}
                                                _ = v23
                                            case "":
                                                v7 = (-(v7))
                                            }
                                            var v24 bool = (519 >= 1)
                                            _ = v22
                                            _ = v24
                                        } else {
                                            v25 := []int{(fi19 | (fi19 - fi19)), 1, fi19}
                                            _ = v25
                                        }
                                        var v26 float64 = 0.0
                                        _ = fi21
                                        _ = v26
                                    }
                                }
                                _ = fi19
                            }
                        }
                    }
                }
                if (-128 == -82) {
                    var v27 uint32 = 2147483648
                    v27 -= 2147483648
                    _ = v27
                } else {
                    switch ("\n\t" + ((("foo bar" + (v8 + (("\n\t" + "hello") + v8))) + (v8 + ("a" + "hello"))) + ((((("a" + v8) + "a") + "") + "") + (("a" + ("a" + v8)) + v8)))) {
                    case "\n\t":
                        if true {
                            v28 := []int{-1}
                            v7 = v7
                            switch 555 {
                            case 299:
                                v28 = append(v28, 32768)
                            case -865:
                                if true {
                                    var v29 uint16 = 1
                                    v30 := GenS1{F0: 222}
                                    v8 = ""
                                    _ = v29
                                    _ = v30
                                } else {
                                    fmt.Println(v8)
                                }
                            default:
                                var v31 int = 367
                                _ = v31
                            }
                            _ = v28
                        }
                    case "a":
                        for fi32 := 0; fi32 < 13; fi32++ {
                            v8 = v8
                            _ = fi32
                        }
                    case "hello":
                        v33 := GenS2{F0: (v8 + "foo bar")}
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 13; fi34++ {
                        v35 := GenS3{F0: -68}
                        fmt.Println(v8, v7, v7)
                        _ = fi34
                        _ = v35
                    }
                    if (1 < 0) {
                        v36 := GenS4{F0: 0}
                        var v37 uint8 = 46
                        v38 := []int{}
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                }
                fmt.Println(v7, v8)
            } else {
                v39 := GenS5{F0: 0, F1: 42, F2: 0}
                v40 := []int{42}
                for fi41 := 0; fi41 < 9; fi41++ {
                    switch v8 {
                    case "hello":
                        v42 := GenS3{F0: 42}
                        _ = v42
                    case "\n\t":
                        v7 -= 1
                    case "foo bar":
                        v7 += (v7 / ((v7) | 1))
                    }
                    _ = fi41
                }
                _ = v39
                _ = v40
            }
        default:
            var v43 float32 = 1e10
            _ = v43
        }
        _ = v11
    }
    _ = v7
    _ = v8
    return 432
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 2
    var v3 uint = 65535
    var v4 uint = (v3 / ((0) | 1))
    v5 := make(map[uint32]string)
    var v6 int64 = 173
    var v44 bool = true
    var v45 string = "\n\t"
    if true {
        if v44 {
            if (v44 && true) {
                switch v45 {
                case "":
                    if true {
                        var v46 float32 = 1.5
                        switch 32768 {
                        case 2147483647:
                            v6 -= (^(v6))
                        case -82:
                            fmt.Println(v6, v4)
                        case 0:
                            v3 -= v3
                        default:
                            v2 = v2
                        }
                        _ = v46
                    }
                case "foo bar":
                    v47 := v1[960]
                    _ = v47
                case "hello":
                    fmt.Println(v3)
                default:
                    switch -528 {
                    case 4294967295:
                        var v48 int = 2147483647
                        _ = v48
                    case 678:
                        v4 = (^(v3))
                    case -904:
                        var v49 uint = (v4 | 42)
                        _ = v49
                    default:
                        v50 := GenS5{F0: 525, F1: 1, F2: 127}
                        _ = v50
                    }
                }
                var v51 string = v45
                if v44 {
                    if v44 {
                        v51 = ((("" + v45) + "") + "hello")
                    } else {
                        if (!v44) {
                            v52 := GenS3{F0: -1}
                            _ = v52
                        } else {
                            v6 -= v6
                            var v53 bool = v44
                            v51 = v51
                            _ = v53
                        }
                        delete(v5, 2147483647)
                    }
                    v6 += v6
                    v54 := GenS5{F0: v2, F1: 0, F2: 127}
                    _ = v54
                } else {
                    v55 := GenS6{F0: 1, F1: 4294967295, F2: v45}
                    v56 := GenS7{F0: 0}
                    _ = v55
                    _ = v56
                }
                _ = v51
            }
        }
        for fi57 := 0; fi57 < 19; fi57++ {
            v58 := []int{(fi57 / (((fi57 / ((int(v6)) | 1))) | 1)), 231}
            var v59 int8 = int8(v3)
            _ = fi57
            _ = v58
            _ = v59
        }
        switch ((((("hello" + (v45 + v45)) + ((v45 + ((v45 + "\n\t") + ("foo bar" + "foo bar"))) + "\n\t")) + "foo bar") + v45) + (("\n\t" + v45) + v45)) {
        case "\n\t":
            v60 := []int{}
            _ = v60
        case "foo bar":
            for fi61 := 0; fi61 < 10; fi61++ {
                switch (fi61 >> (((v4 | 600)) & 63)) {
                case -9223372036854775808:
                    for fi62 := 0; fi62 < 17; fi62++ {
                        if v44 {
                            v63 := GenS6{F0: 202, F1: 4294967295, F2: ""}
                            if (-698 < 598) {
                                v64 := GenS3{F0: -32}
                                _ = v64
                            }
                            if v44 {
                                v65 := GenS1{F0: (v4 + v4)}
                                fmt.Println(v44)
                                _ = v65
                            }
                            _ = v63
                        } else {
                            var v66 uint8 = 128
                            fmt.Println(v3, v4, v4)
                            _ = v66
                        }
                        _ = fi62
                    }
                default:
                    v4 = (v3 / ((v4) | 1))
                }
                for fi67 := 0; fi67 < 3; fi67++ {
                    v68 := []int{(fi67 * fi67), fi61, (fi61 + fi67), -1}
                    fmt.Println(v44)
                    _ = fi67
                    _ = v68
                }
                _ = fi61
            }
        }
    }
    fmt.Println(v44, v3)
    fmt.Println(v4)
    var v69 int16 = 1
    if v44 {
        v70 := []int{}
        _ = v70
    } else {
        v71 := []int{-1, 2147483647, int(v2)}
        _ = v71
    }
    v6 = -86
    v72 := GenS8{F0: 4294967295, F1: 165, F2: 782}
    for fi73 := 0; fi73 < 3; fi73++ {
        if (v2 > (v2 ^ (v2 & genFn0()))) {
            var v74 uint16 = 174
            var v75 int32 = 0
            for fi76 := 0; fi76 < 6; fi76++ {
                v3 = 0
                if v44 {
                    v74 = 468
                    v77 := GenS9{F0: v74, F1: v6}
                    var v78 uint = (^(v3))
                    _ = v77
                    _ = v78
                }
                var v79 float64 = (1e10 - 3.14159)
                _ = fi76
                _ = v79
            }
            _ = v74
            _ = v75
        } else {
            fmt.Println(v3)
            v80 := GenS10{F0: 2147483647, F1: 2147483647}
            v81 := GenS5{F0: uint64(v4), F1: -1, F2: 0}
            _ = v80
            _ = v81
        }
        _ = fi73
    }
    var v82 int16 = (v69 >> ((v4) & 63))
    var v83 int64 = -1
    switch v45 {
    case "a":
        var v84 bool = v44
        _ = v84
    case "\n\t":
        v6 += (v6 + (^(v6)))
    case "":
        switch v45 {
        case "":
            for fi85 := 0; fi85 < 4; fi85++ {
                v83 = v6
                _ = fi85
            }
        case "hello":
            fmt.Println(v83, v82)
        default:
            switch -386 {
            case 9223372036854775807:
                if false {
                    if (v83 > v83) {
                        for fi86 := 0; fi86 < 11; fi86++ {
                            var v87 int32 = 531
                            v88 := GenS2{F0: ("a" + ("a" + v45))}
                            _ = fi86
                            _ = v87
                            _ = v88
                        }
                    } else {
                        v3 = v4
                        var v89 int64 = (v83 % ((v83) | 1))
                        _ = v89
                    }
                    v4 = (v3 + v3)
                }
            case 32767:
                switch 4294967295 {
                case -2147483648:
                    var v90 uint32 = 403
                    _ = v90
                }
            case -9223372036854775808:
                switch 42 {
                case 4294967295:
                    v83 += -748
                }
            default:
                var v91 int64 = (v6 * v6)
                _ = v91
            }
        }
    default:
        v92 := []int{int(v2), 854, -503, 128}
        _ = v92
    }
    fmt.Println(v44)
    fmt.Println(v3, v44, v44)
    v3 = 27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v69
    _ = v72
    _ = v82
    _ = v83
}
