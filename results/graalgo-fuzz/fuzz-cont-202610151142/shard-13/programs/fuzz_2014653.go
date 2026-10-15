package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint; F1 uint; F2 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int16; F1 int16; F2 int64 }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 string; F1 uint; F2 uint }
func genFn0(p0 string, p1 string, p2 int32) int64 {
    p2 -= p2
    var v6 uint16 = 0
    if (100 == -128) {
        switch "foo bar" {
        case "hello":
            for fi7 := 0; fi7 < 11; fi7++ {
                if false {
                    for fi8 := 0; fi8 < 8; fi8++ {
                        var v9 int = (fi7 & (fi7 / ((fi7) | 1)))
                        _ = fi8
                        _ = v9
                    }
                    fmt.Println(p1)
                }
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 int32 = p2
                    v12 := GenS1{F0: (4294967295 <= 85)}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v13 := []int{(^(fi7))}
                _ = fi7
                _ = v13
            }
        case "foo bar":
            var v14 string = "hello"
            _ = v14
        case "a":
            var v15 int64 = 9223372036854775807
            _ = v15
        default:
            v16 := GenS1{F0: (false || true)}
            _ = v16
        }
        v17 := GenS1{F0: (50 < 4294967295)}
        _ = v17
    }
    v18 := GenS2{F0: (-(p2)), F1: 557}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v6
    return -1
}
func genFn1(p0 string, p1 uint64) int8 {
    v19 := []int{-771, 693, -769}
    p1 = 18446744073709551615
    p0 = (("" + p0) + ("\n\t" + p0))
    var v20 bool = true
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 100
}

func main() {
    var v1 uint8 = 255
    var v2 bool = ((true || false) && true)
    var v3 string = "a"
    v4 := make(map[uint32]int)
    v5 := []int{32767}
    if (0 <= 4294967295) {
        if (4294967295 < 508) {
            v1 = v1
        } else {
            fmt.Println(v3, v1, v1)
        }
        if (false || false) {
            var v21 uint = 514
            v5 = append(v5, -9223372036854775808)
            _ = v21
        }
        v3 = "\n\t"
    }
    v22 := GenS3{F0: v2, F1: 255}
    switch -1 {
    case 32767:
        var v23 string = (("foo bar" + v3) + (v3 + (v3 + (v3 + v3))))
        _ = v23
    case 32768:
        var v24 int64 = -9223372036854775808
        _ = v24
    case 4294967295:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        var v25 uint32 = 1
        _ = v25
    }
    var v26 int16 = 1
    for fi27 := 0; fi27 < 15; fi27++ {
        if false {
            var v28 string = "hello"
            v28 = ("foo bar" + v28)
            for fi29 := 0; fi29 < 2; fi29++ {
                fmt.Println(v28, v26, v2)
                v28 = ""
                switch 9223372036854775807 {
                case 102:
                    if (!false) {
                        if v2 {
                            var v30 bool = v2
                            v1 = 42
                            _ = v30
                        }
                        if v2 {
                            if false {
                                var v31 float64 = 3.14159
                                _ = v31
                            }
                        }
                        v3 = (v3 + v3)
                    } else {
                        _ = len(v5)
                        v32 := GenS4{F0: v26, F1: v1, F2: 1}
                        _ = v32
                    }
                case 970:
                    v33 := GenS5{F0: 962, F1: 1, F2: -9223372036854775808}
                    _ = v33
                case 2147483647:
                    fmt.Println(v2, v1)
                }
                var v34 string = "foo bar"
                _ = fi29
                _ = v34
            }
            _ = v28
        }
        fmt.Println(v3)
        _ = fi27
    }
    var v35 uint32 = 2147483647
    v35 += v35
    v4[v35] = 9223372036854775807
    switch 32768 {
    case -2147483648:
        if false {
            v36 := GenS1{F0: (0 == 42)}
            v37 := GenS1{F0: v2}
            var v38 int8 = genFn1((v3 + ("hello" + v3)), 18446744073709551615)
            _ = v36
            _ = v37
            _ = v38
        } else {
            switch ((v3 + v3) + v3) {
            case "foo bar":
                for fi39 := 0; fi39 < 2; fi39++ {
                    var v40 uint = 0
                    var v41 int64 = genFn0((v3 + (((v3 + (("a" + v3) + ("a" + (((("" + (v3 + (("a" + "foo bar") + ((((v3 + v3) + v3) + (v3 + v3)) + "\n\t")))) + "") + "hello") + (((v3 + "hello") + ("\n\t" + v3)) + ("a" + ("a" + v3))))))) + v3) + "a")), (v3 + "a"), 674)
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            case "a":
                var v42 uint16 = 0
                _ = v42
            case "hello":
                fmt.Println(v35, v2, v2)
            }
            if v2 {
                switch -594 {
                case 0:
                    switch -2147483648 {
                    case 127:
                        for fi43 := 0; fi43 < 14; fi43++ {
                            if true {
                                if true {
                                    if ((^(v26)) >= (v26 >> ((989) & 63))) {
                                        var v44 int32 = 42
                                        _ = v44
                                    }
                                    for fi45 := 0; fi45 < 9; fi45++ {
                                        if (2 > 18446744073709551615) {
                                            v46 := GenS6{F0: v35}
                                            _ = v46
                                        }
                                        var v47 uint8 = 196
                                        fmt.Println(fi45)
                                        _ = fi45
                                        _ = v47
                                    }
                                }
                                v48 := GenS1{F0: (-2.25 > 3.14159)}
                                var v49 int8 = -1
                                _ = v48
                                _ = v49
                            }
                            v50 := GenS3{F0: v2, F1: 65535}
                            _ = fi43
                            _ = v50
                        }
                    default:
                        v51 := GenS7{F0: v35, F1: 276}
                        _ = v51
                    }
                default:
                    fmt.Println(v1)
                }
            } else {
                for fi52 := 0; fi52 < 13; fi52++ {
                    var v53 int32 = -1
                    switch ((("foo bar" + v3) + (v3 + v3)) + (v3 + v3)) {
                    case "hello":
                        v54 := v4[uint32(v1)]
                        _ = v54
                    case "a":
                        v55 := GenS6{F0: (v35 | v35)}
                        _ = v55
                    default:
                        v56 := GenS8{F0: 723, F1: 42, F2: -926}
                        _ = v56
                    }
                    _ = fi52
                    _ = v53
                }
                var v57 uint8 = (v1 % (((^(v1))) | 1))
                _ = len(v5)
                _ = v57
            }
        }
    case -570:
        for fi58 := 0; fi58 < 17; fi58++ {
            var v59 float64 = 1e10
            if (0 <= -682) {
                var v60 bool = (0 <= 0)
                for fi61 := 0; fi61 < 16; fi61++ {
                    if v60 {
                        switch 4294967295 {
                        case 32768:
                            if false {
                                _ = len(v4)
                            }
                        case 574:
                            v2 = true
                        }
                        var v62 string = "a"
                        if false {
                            if ((genFn0(v3, "a", 1) <= -1) || true) {
                                fmt.Println(v62, fi58)
                                fmt.Println(fi61)
                            } else {
                                var v63 uint16 = 65535
                                _ = v63
                            }
                        } else {
                            switch (v3 + "\n\t") {
                            case "hello":
                                var v64 uint32 = v35
                                _ = v64
                            }
                            fmt.Println(v1, v2)
                            fmt.Println(v1, v59, v3)
                        }
                        _ = v62
                    } else {
                        v65 := GenS7{F0: v35, F1: 127}
                        v66 := GenS9{F0: 0, F1: 42}
                        switch (fi61 | fi61) {
                        case 501:
                            v67 := GenS3{F0: (!true), F1: 42}
                            _ = v67
                        }
                        _ = v65
                        _ = v66
                    }
                    _ = len(v4)
                    v68 := GenS9{F0: 18446744073709551615, F1: (v35 & v35)}
                    _ = fi61
                    _ = v68
                }
                var v69 uint8 = (v1 * 97)
                _ = v60
                _ = v69
            } else {
                if (329 != 244) {
                    if v2 {
                        for fi70 := 0; fi70 < 13; fi70++ {
                            if true {
                                fmt.Println(v35, fi70, v3)
                                var v71 bool = (v35 == uint32(v1))
                                var v72 uint8 = (^(v1))
                                _ = v71
                                _ = v72
                            }
                            _ = fi70
                        }
                    } else {
                        for fi73 := 0; fi73 < 9; fi73++ {
                            v35 += 42
                            _ = fi73
                        }
                        if false {
                            var v74 bool = (2147483647 <= 0)
                            v75 := GenS2{F0: -860, F1: (-(v26))}
                            v35 -= 697
                            _ = v74
                            _ = v75
                        }
                    }
                    for fi76 := 0; fi76 < 8; fi76++ {
                        delete(v4, (v35 / (((^(v35))) | 1)))
                        _ = fi76
                    }
                    for fi77 := 0; fi77 < 3; fi77++ {
                        var v78 int8 = 1
                        fmt.Println(fi77, fi58)
                        v79 := v4[v35]
                        _ = fi77
                        _ = v78
                        _ = v79
                    }
                } else {
                    v2 = false
                    switch "" {
                    case "\n\t":
                        for fi80 := 0; fi80 < 4; fi80++ {
                            var v81 float32 = 1e10
                            if (v2 && (!(2 >= 1))) {
                                switch int(v1) {
                                case 9223372036854775807:
                                    if false {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    } else {
                                        var v82 uint16 = 128
                                        v35 += 0
                                        _ = v82
                                    }
                                case 403:
                                    v83 := v4[v35]
                                    _ = v83
                                default:
                                    if (1000 <= (v26 ^ (^(v26)))) {
                                        var v84 int64 = -1
                                        _ = v84
                                    }
                                    v59 = (float64(((-2.25 * (1.5 / v81)) * (v81 - (v81 / (-2.25 * 3.14159))))) - float64(v81))
                                }
                                if len(v5) > 0 {
                                    v5[0] = (fi58 & (-(fi58)))
                                }
                            }
                            v85 := GenS1{F0: true}
                            _ = fi80
                            _ = v81
                            _ = v85
                        }
                    }
                }
                v4[65535] = fi58
                if ((!(!true)) || v2) {
                    v3 = "a"
                }
            }
            _ = fi58
            _ = v59
        }
    default:
        var v86 int32 = 162
        _ = v86
    }
    fmt.Println(v26, v1, v26)
    for fi87 := 0; fi87 < 3; fi87++ {
        var v88 uint = 0
        fmt.Println(fi87)
        _ = fi87
        _ = v88
    }
    v26 -= 32767
    for fi89 := 0; fi89 < 4; fi89++ {
        v90 := GenS10{F0: (v3 + v3), F1: 2, F2: 1}
        _ = fi89
        _ = v90
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
