package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint8; F2 int32 }
type GenS6 struct { F0 uint64; F1 string; F2 uint8 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 uint64; F1 uint32 }
type GenS10 struct { F0 uint32; F1 uint32 }
type GenS11 struct { F0 uint8; F1 bool }
type GenS12 struct { F0 uint; F1 uint }
type GenS13 struct { F0 int32; F1 int; F2 int64 }
func genFn0(p0 uint) int64 {
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(fi8)
            var v9 uint64 = 1
            fmt.Println(fi8)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: 32767}
    _ = p0
    _ = v10
    return -1
}

func main() {
    var v1 int16 = 341
    var v2 int64 = 0
    var v3 int64 = (v2 << ((18) & 63))
    var v4 int64 = v2
    var v5 uint32 = 941
    v6 := []uint32{(v5 & 2), 1, (^(v5)), (v5 >> ((1) & 63))}
    fmt.Println(v3, v5)
    var v11 float64 = (0.0 * 0.0)
    v12 := GenS1{F0: (v1 ^ 1000)}
    if true {
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 uint8 = 0
            _ = fi13
            _ = v14
        }
    }
    v15 := GenS2{F0: 42}
    for fi16 := 0; fi16 < 10; fi16++ {
        v1 -= v1
        switch ((("" + "\n\t") + ("foo bar" + "a")) + ("foo bar" + "")) {
        case "\n\t":
            for fi17 := 0; fi17 < 19; fi17++ {
                if (!true) {
                    if ((true || (-44 != -128)) && (true && false)) {
                        v18 := GenS3{F0: 1, F1: (("" + ("a" + "a")) + "")}
                        _ = v18
                    }
                    v19 := GenS2{F0: 142}
                    var v20 uint64 = 2
                    _ = v19
                    _ = v20
                }
                switch "\n\t" {
                case "\n\t":
                    var v21 uint32 = v5
                    _ = v21
                case "a":
                    if false {
                        if false {
                            if (0 != 42) {
                                if true {
                                    _ = len(v6)
                                } else {
                                    v6 = append(v6, v5)
                                    v22 := GenS1{F0: 0}
                                    v4 -= (v4 | v2)
                                    v23 := GenS4{F0: 255}
                                    _ = len(v6)
                                    _ = v22
                                    _ = v23
                                }
                                v24 := GenS5{F0: 2, F1: 178, F2: -2147483648}
                                _ = v24
                            }
                            var v25 bool = ("" >= "\n\t")
                            var v26 int = 9223372036854775807
                            switch -782 {
                            case 127:
                                v27 := GenS6{F0: 2, F1: ((("" + "") + ("foo bar" + "\n\t")) + "foo bar"), F2: 42}
                                _ = v27
                            }
                            var v28 uint = 2
                            v5 += (v5 / (((^(v5))) | 1))
                            _ = v25
                            _ = v26
                            _ = v28
                        }
                        v11 += 1.5
                        var v29 int64 = v2
                        _ = v29
                    } else {
                        v30 := GenS4{F0: 4294967295}
                        switch int(v3) {
                        case 378:
                            switch (fi17 << ((42) & 63)) {
                            case 32768:
                                v1 -= (v1 >> ((4294967295) & 63))
                            case -606:
                                if len(v6) > 0 {
                                    v6[0] = (v5 / ((2147483647) | 1))
                                }
                            default:
                                v2 = (v3 & (v4 * genFn0(806)))
                            }
                        }
                        if true {
                            var v31 int64 = -848
                            _ = v31
                        } else {
                            var v32 uint32 = (v5 - 2147483647)
                            v1 += 460
                            v33 := GenS1{F0: v1}
                            _ = v32
                            _ = v33
                        }
                        var v34 float32 = 3.14159
                        v34 -= (v34 - (0.0 / v34))
                        v11 = v11
                        if false {
                            _ = len(v6)
                            v5 = (^(v5))
                            v3 = v2
                        }
                        var v35 int = (fi17 & 2147483647)
                        _ = v30
                        _ = v34
                        _ = v35
                    }
                case "":
                    var v36 int32 = 2147483647
                    _ = v36
                }
                _ = fi17
            }
        }
        _ = fi16
    }
    var v37 int64 = v4
    if ((float32((((v11 / (v11 / 0.1)) + v11) + 1.5)) * 0.1) < 3.14159) {
        var v38 uint = 563
        for fi39 := 0; fi39 < 18; fi39++ {
            switch "a" {
            case "hello":
                for fi40 := 0; fi40 < 11; fi40++ {
                    switch "" {
                    case "hello":
                        v37 += (-(v3))
                    case "foo bar":
                        if (v5 < (v5 ^ (v5 % (((v5 % (((^(v5))) | 1))) | 1)))) {
                            v3 = v37
                            var v41 int32 = 367
                            var v42 uint8 = 2
                            _ = v41
                            _ = v42
                        }
                    default:
                        fmt.Println(fi39, v3, fi40)
                        var v43 int64 = v3
                        fmt.Println(fi40, v5, v2)
                        _ = v43
                    }
                    v6 = append(v6, v5)
                    fmt.Println(v37)
                    _ = fi40
                }
            case "foo bar":
                switch (fi39 << (((v38 >> (((v38 | v38)) & 63))) & 63)) {
                case -1:
                    var v44 int8 = 0
                    _ = v44
                case 93:
                    for fi45 := 0; fi45 < 17; fi45++ {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        fmt.Println(fi45, v38, v11)
                        switch "" {
                        case "":
                            if (805 < 606) {
                                v11 -= -2.25
                                v46 := GenS2{F0: 1}
                                _ = v46
                            }
                        case "a":
                            if len(v6) > 0 {
                                v6[0] = v5
                            }
                        case "foo bar":
                            var v47 int64 = v4
                            if true {
                                v48 := GenS5{F0: v38, F1: 153, F2: -319}
                                _ = v48
                            } else {
                                if len(v6) > 0 {
                                    v6[0] = (v5 | (v5 >> (((v38 % ((v38) | 1))) & 63)))
                                }
                                var v49 uint16 = 32
                                v50 := GenS7{F0: 0, F1: fi45}
                                _ = v49
                                _ = v50
                            }
                            if false {
                                switch fi39 {
                                case 4294967295:
                                    v2 = v2
                                case -2147483648:
                                    if len(v6) > 0 {
                                        v6[0] = (v5 * (v5 >> ((1) & 63)))
                                    }
                                default:
                                    fmt.Println(v38, v47, v1)
                                }
                                if false {
                                    var v51 int16 = v1
                                    _ = v51
                                } else {
                                    var v52 uint16 = 139
                                    v37 += (v2 * genFn0((^(v38))))
                                    v38 = (v38 >> ((705) & 63))
                                    v53 := GenS1{F0: (v1 >> ((v38) & 63))}
                                    _ = v52
                                    _ = v53
                                }
                                var v54 string = ("a" + "")
                                v55 := GenS2{F0: 42}
                                _ = v54
                                _ = v55
                            }
                            var v56 uint16 = 570
                            _ = v47
                            _ = v56
                        default:
                            var v57 float64 = (v11 - float64(1e10))
                            _ = v57
                        }
                        _ = fi45
                    }
                case 127:
                    var v58 bool = ((v4 - (^(v37))) <= genFn0((v38 - 533)))
                    _ = v58
                default:
                    switch fi39 {
                    case 0:
                        switch fi39 {
                        case 32768:
                            v37 += v37
                        default:
                            var v59 int32 = -959
                            _ = v59
                        }
                    case 26:
                        for fi60 := 0; fi60 < 18; fi60++ {
                            if ("\n\t" < ("foo bar" + "")) {
                                v37 += (v3 - (v37 & (v3 & (-(v4)))))
                                v11 = 1.5
                            } else {
                                switch (fi60 & int(v38)) {
                                case 842:
                                    fmt.Println(v37, fi60, fi39)
                                case -9223372036854775808:
                                    if true {
                                        var v61 uint64 = 4294967295
                                        v4 += int64(v5)
                                        _ = v61
                                    }
                                }
                                var v62 float64 = (1.5 + v11)
                                v63 := GenS8{F0: 713, F1: (^(v5))}
                                _ = v62
                                _ = v63
                            }
                            v2 += v2
                            v64 := GenS2{F0: 2}
                            v4 = (v37 >> ((v38) & 63))
                            _ = fi60
                            _ = v64
                        }
                    case 4294967295:
                        fmt.Println(v2)
                    }
                }
            case "\n\t":
                for fi65 := 0; fi65 < 9; fi65++ {
                    v11 = ((1.5 / -2.25) / (v11 + 0.0))
                    v66 := GenS6{F0: 474, F1: "foo bar", F2: 42}
                    if (v1 < v1) {
                        v67 := GenS9{F0: 0, F1: v5}
                        _ = v67
                    } else {
                        fmt.Println(v4, v1, v11)
                        v68 := GenS10{F0: 65535, F1: 2}
                        if false {
                            v2 += (v3 & (v2 & v4))
                            var v69 uint64 = 0
                            _ = v69
                        } else {
                            v6 = append(v6, v5)
                            var v70 uint = 18446744073709551615
                            v71 := GenS7{F0: 822, F1: (fi39 % (((fi65 >> (((v70 - v38)) & 63))) | 1))}
                            v70 = (v70 >> (((v70 * (^(v70)))) & 63))
                            _ = v70
                            _ = v71
                        }
                        _ = v68
                    }
                    _ = fi65
                    _ = v66
                }
            default:
                var v72 int32 = 2147483647
                _ = v72
            }
            _ = fi39
        }
        v73 := GenS11{F0: 127, F1: false}
        _ = v38
        _ = v73
    }
    switch "foo bar" {
    case "foo bar":
        if len(v6) > 0 {
            v6[0] = v5
        }
    case "":
        switch ("hello" + ("" + ("hello" + "\n\t"))) {
        case "a":
            v1 -= (v1 / ((60) | 1))
        default:
            v74 := GenS12{F0: 18446744073709551615, F1: 894}
            _ = v74
        }
    default:
        switch "foo bar" {
        case "hello":
            for fi75 := 0; fi75 < 10; fi75++ {
                if true {
                    v5 = v5
                    var v76 uint16 = 93
                    _ = v76
                } else {
                    v77 := GenS1{F0: v1}
                    var v78 float64 = (v11 * (float64(0.0) / 1.0))
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = v77
                    _ = v78
                }
                v79 := GenS13{F0: -1, F1: (fi75 | (fi75 * -2147483648)), F2: v37}
                _ = fi75
                _ = v79
            }
        default:
            var v80 string = ("a" + "hello")
            _ = v80
        }
    }
    v37 += v37
    v6 = append(v6, v5)
    for fi81 := 0; fi81 < 15; fi81++ {
        var v82 bool = false
        var v83 uint32 = (v5 % ((v5) | 1))
        var v84 uint16 = 42
        _ = fi81
        _ = v82
        _ = v83
        _ = v84
    }
    switch ("\n\t" + "hello") {
    case "hello":
        for fi85 := 0; fi85 < 2; fi85++ {
            if (!false) {
                var v86 bool = true
                var v87 float32 = ((0.1 - (3.14159 + (0.0 / ((0.0 + -2.25) / 1.5)))) * 3.14159)
                _ = v86
                _ = v87
            } else {
                fmt.Println(v11)
            }
            var v88 uint64 = uint64(v5)
            _ = fi85
            _ = v88
        }
    default:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
