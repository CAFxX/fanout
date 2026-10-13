package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int; F2 int64 }
type GenS3 struct { F0 int16; F1 int16; F2 string }
type GenS4 struct { F0 int; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int32; F1 uint64; F2 string }
type GenS6 struct { F0 int16; F1 string; F2 int }
type GenS7 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS8 struct { F0 uint; F1 uint32 }
type GenS9 struct { F0 uint; F1 int32; F2 int32 }
func genFn0() uint64 {
    var v5 int = 1
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: 1}
            v9 := GenS2{F0: -244, F1: (fi6 | (fi6 >> ((638) & 63))), F2: 1}
            v10 := []int{4294967295, fi7, 1, 32768}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v11 := []int{-9223372036854775808, -9223372036854775808}
    v12 := GenS1{F0: -488}
    fmt.Println(v5)
    _ = v11
    _ = v12
    _ = v5
    return 1
}
func genFn1() uint {
    switch 32767 {
    case 0:
        v13 := []int{127}
        _ = v13
    }
    var v14 int8 = 1
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            switch "hello" {
            case "a":
                var v17 uint8 = 255
                _ = v17
            case "foo bar":
                var v18 int = 32768
                _ = v18
            }
            var v19 uint = 664
            v14 = 127
            _ = fi16
            _ = v19
        }
        if (0 <= 0) {
            v14 -= v14
        }
        _ = fi15
    }
    v20 := []int{}
    if false {
        fmt.Println(v14)
        fmt.Println(v14)
    }
    _ = v14
    _ = v20
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{2147483647, 2147483648}
    var v3 uint64 = 1
    var v4 string = (("foo bar" + (((("" + ("a" + "foo bar")) + "a") + "") + "")) + "hello")
    var v21 bool = (!(!false))
    for fi22 := 0; fi22 < 17; fi22++ {
        switch fi22 {
        case 9223372036854775807:
            for fi23 := 0; fi23 < 4; fi23++ {
                v4 = (v4 + v4)
                if v21 {
                    v4 = (v4 + "")
                    v3 = 853
                } else {
                    var v24 int64 = 9223372036854775807
                    v21 = (!v21)
                    v24 += v24
                    _ = v24
                }
                _ = fi23
            }
        case -9223372036854775808:
            v3 -= v3
        default:
            for fi25 := 0; fi25 < 5; fi25++ {
                switch (("\n\t" + v4) + (v4 + (v4 + ""))) {
                case "hello":
                    v26 := GenS3{F0: 0, F1: -69, F2: ("" + (v4 + v4))}
                    _ = v26
                }
                switch v4 {
                case "":
                    var v27 int16 = 616
                    _ = v27
                case "a":
                    var v28 int32 = int32(fi22)
                    _ = v28
                default:
                    v3 += (v3 & v3)
                    v29 := GenS4{F0: (fi25 ^ fi25), F1: v3, F2: 2}
                    _ = v29
                }
                if v21 {
                    var v30 float64 = (3.14159 / 1e10)
                    v31 := GenS4{F0: (fi25 - 850), F1: (v3 - (v3 - 165)), F2: 65535}
                    v32 := GenS4{F0: (fi22 + 32767), F1: genFn0(), F2: 42}
                    _ = v30
                    _ = v31
                    _ = v32
                }
                delete(v1, fi22)
                _ = fi25
            }
        }
        var v33 int8 = -128
        _ = fi22
        _ = v33
    }
    fmt.Println(v4, v3)
    var v34 float32 = 1.5
    var v35 uint = 255
    v36 := GenS5{F0: -1, F1: v3, F2: ((("" + v4) + (v4 + v4)) + "")}
    switch 2147483647 {
    case 9223372036854775807:
        if (((1.5 / (3.14159 / 1.5)) != 0.1) && (!true)) {
            for fi37 := 0; fi37 < 7; fi37++ {
                if (v21 && (v3 < 2)) {
                    v38 := GenS1{F0: -9223372036854775808}
                    _ = v38
                }
                if ((!(!v21)) || v21) {
                    var v39 int16 = -1
                    _ = v39
                } else {
                    if ((fi37 < (fi37 ^ -1)) || v21) {
                        if (0 <= 1) {
                            fmt.Println(v21, v35, v3)
                            for fi40 := 0; fi40 < 5; fi40++ {
                                v34 = (v34 * v34)
                                switch int(v35) {
                                case 128:
                                    v21 = (float64(float32(1.5)) >= 1e10)
                                }
                                var v41 int64 = 1
                                _ = fi40
                                _ = v41
                            }
                        } else {
                            switch "hello" {
                            case "foo bar":
                                switch (fi37 ^ (fi37 << (((v35 ^ v35)) & 63))) {
                                case 127:
                                    switch v4 {
                                    case "hello":
                                        var v42 bool = (v3 < genFn0())
                                        _ = v42
                                    case "a":
                                        var v43 uint16 = 468
                                        _ = v43
                                    case "\n\t":
                                        switch (-(fi37)) {
                                        case 545:
                                            v21 = v21
                                        }
                                    }
                                case -2147483648:
                                    for fi44 := 0; fi44 < 6; fi44++ {
                                        var v45 bool = v21
                                        v4 = (v4 + ("" + (("" + ((v4 + (v4 + v4)) + v4)) + ("" + ((v4 + "") + ((v4 + v4) + v4))))))
                                        _ = fi44
                                        _ = v45
                                    }
                                case -9223372036854775808:
                                    _ = len(v1)
                                }
                            case "a":
                                fmt.Println(v21, fi37)
                            case "\n\t":
                                v3 = genFn0()
                            }
                            switch "a" {
                            case "foo bar":
                                if (true || false) {
                                    v46 := GenS1{F0: 360}
                                    var v47 uint = v35
                                    _ = v46
                                    _ = v47
                                }
                            case "":
                                if (3.14159 != 0.0) {
                                    var v48 uint16 = 2
                                    _ = v48
                                }
                            default:
                                v4 = v4
                            }
                            for fi49 := 0; fi49 < 10; fi49++ {
                                switch v4 {
                                case "\n\t":
                                    var v50 uint8 = 128
                                    _ = v50
                                case "hello":
                                    var v51 float64 = (0.1 / (-2.25 * 3.14159))
                                    _ = v51
                                case "foo bar":
                                    switch v4 {
                                    case "\n\t":
                                        delete(v1, -2147483648)
                                    }
                                }
                                v3 = v3
                                var v52 int16 = -1
                                _ = fi49
                                _ = v52
                            }
                        }
                        for fi53 := 0; fi53 < 16; fi53++ {
                            if true {
                                switch 32768 {
                                case -354:
                                    v54 := GenS6{F0: -843, F1: "foo bar", F2: -2147483648}
                                    _ = v54
                                }
                                _ = len(v2)
                            } else {
                                v4 = "foo bar"
                                var v55 int = fi37
                                _ = v55
                            }
                            if (886 < 2) {
                                var v56 bool = v21
                                v3 -= 18446744073709551615
                                v21 = true
                                v57 := GenS7{F0: 863, F1: 48, F2: 2}
                                _ = v56
                                _ = v57
                            }
                            _ = fi53
                        }
                        v58 := GenS7{F0: 2, F1: 169, F2: 2147483648}
                        _ = v58
                    }
                    v35 += (v35 * 660)
                    if false {
                        fmt.Println(v4)
                        v4 = "hello"
                        v1[(fi37 / ((42) | 1))] = true
                    } else {
                        if (2 >= 715) {
                            var v59 bool = (0.0 >= (1.5 / 1.0))
                            v21 = v59
                            _ = v59
                        } else {
                            if v21 {
                                var v60 float32 = 1.5
                                _ = v60
                            }
                            var v61 uint16 = 2
                            var v62 int32 = -1
                            _ = v61
                            _ = v62
                        }
                        _ = len(v1)
                    }
                }
                var v63 int32 = -2147483648
                _ = fi37
                _ = v63
            }
        } else {
            var v64 bool = false
            _ = v64
        }
    case 766:
        var v65 int = 32768
        _ = v65
    case 0:
        v34 = float32(0.1)
    default:
        for fi66 := 0; fi66 < 1; fi66++ {
            var v67 string = v4
            for fi68 := 0; fi68 < 6; fi68++ {
                if (0.0 >= 1.5) {
                    v4 = ((v4 + (v4 + v67)) + (v67 + v67))
                } else {
                    v69 := GenS2{F0: 9223372036854775807, F1: 32768, F2: 42}
                    v34 = -2.25
                    var v70 uint16 = 255
                    v21 = (380 == 203)
                    if (127 == 255) {
                        fmt.Println(v67, v67)
                    }
                    _ = v69
                    _ = v70
                }
                _ = fi68
            }
            switch (fi66 * fi66) {
            case 128:
                fmt.Println(v3)
            case 4294967295:
                v71 := GenS7{F0: 65535, F1: 2, F2: 836}
                _ = v71
            case -865:
                v3 -= v3
            }
            _ = fi66
            _ = v67
        }
    }
    if ((v4 + v4) >= v4) {
        if (v21 && true) {
            if (477 <= 65535) {
                var v72 uint64 = 386
                _ = v72
            } else {
                var v73 int8 = 0
                _ = v73
            }
            var v74 int32 = 0
            _ = v74
        }
    }
    for fi75 := 0; fi75 < 9; fi75++ {
        var v76 float64 = (3.14159 * (1.5 - 1e10))
        _ = fi75
        _ = v76
    }
    fmt.Println(v21)
    for fi77 := 0; fi77 < 16; fi77++ {
        var v78 int = (fi77 % (((fi77 >> ((genFn1()) & 63))) | 1))
        switch ("a" + v4) {
        case "a":
            _ = len(v1)
        case "":
            v4 = (v4 + v4)
        default:
            v79 := GenS1{F0: 400}
            _ = v79
        }
        _ = fi77
        _ = v78
    }
    v80 := GenS4{F0: -639, F1: (v3 / (((v3 >> ((49) & 63))) | 1)), F2: 42}
    v81 := GenS8{F0: v35, F1: 2}
    var v82 uint16 = 0
    v83 := GenS9{F0: v35, F1: 233, F2: -1}
    for fi84 := 0; fi84 < 13; fi84++ {
        for fi85 := 0; fi85 < 11; fi85++ {
            switch ("foo bar" + v4) {
            case "foo bar":
                _ = len(v2)
                var v86 int = -286
                _ = v86
            }
            var v87 uint16 = 2
            fmt.Println(v35, v3, v35)
            _ = fi85
            _ = v87
        }
        _ = fi84
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v80
    _ = v81
    _ = v82
    _ = v83
}
