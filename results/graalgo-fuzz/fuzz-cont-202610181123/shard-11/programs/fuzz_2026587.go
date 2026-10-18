package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 string; F1 int32 }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint16; F1 string }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint64, p1 int32, p2 uint64) uint {
    var v7 float32 = 3.14159
    if true {
        fmt.Println(v7, p0)
        v8 := []int{-854}
        v7 += v7
        _ = v8
    } else {
        v9 := []int{-2147483648, 494, 32768}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v9
    }
    p0 -= p2
    var v10 uint8 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 4294967295
    v3 := []string{"foo bar", "foo bar", "\n\t"}
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    v6 := make(map[string]bool)
    var v11 uint64 = 42
    var v12 int16 = 1000
    v11 += v11
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 17; fi14++ {
            v2 = 2147483647
            v15 := GenS1{F0: 2, F1: v12}
            v16 := GenS2{F0: 836}
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 uint64 = v11
        _ = fi13
        _ = v17
    }
    v12 = v12
    for fi18 := 0; fi18 < 12; fi18++ {
        switch (fi18 >> ((0) & 63)) {
        case 42:
            _ = len(v5)
        case -9223372036854775808:
            for fi19 := 0; fi19 < 16; fi19++ {
                switch "" {
                case "foo bar":
                    var v20 int64 = int64(v12)
                    _ = v20
                case "\n\t":
                    v21 := GenS1{F0: 65535, F1: v12}
                    _ = v21
                case "a":
                    if (true && (!true)) {
                        if true {
                            if (!(!true)) {
                                v11 = v11
                                v22 := GenS3{F0: 215, F1: 346}
                                _ = v22
                            } else {
                                var v23 uint64 = v11
                                var v24 int32 = 42
                                _ = v23
                                _ = v24
                            }
                            v25 := GenS1{F0: 2, F1: (v12 / (((v12 << ((2) & 63))) | 1))}
                            v11 = (v11 * v11)
                            _ = v25
                        } else {
                            var v26 uint8 = 1
                            var v27 uint8 = (v26 & v26)
                            _ = v26
                            _ = v27
                        }
                        if (42 >= 105) {
                            _ = len(v1)
                            if false {
                                fmt.Println(fi18)
                            } else {
                                if true {
                                    switch fi18 {
                                    case 42:
                                        v2 -= 1
                                    case 127:
                                        v11 -= v11
                                    case 32768:
                                        v28 := v5[(fi18 / ((fi19) | 1))]
                                        _ = v28
                                    default:
                                        delete(v4, (^(v2)))
                                    }
                                    if true {
                                        var v29 int64 = 0
                                        if ((!false) || (298 != v11)) {
                                            var v30 uint32 = (v2 * 789)
                                            _ = v30
                                        } else {
                                            v3 = append(v3, (("foo bar" + ("" + (((("\n\t" + ("foo bar" + "")) + "") + "hello") + "a"))) + "\n\t"))
                                            fmt.Println(v11)
                                            v29 -= (^(v29))
                                            v31 := GenS1{F0: 32767, F1: -977}
                                            _ = v31
                                        }
                                        var v32 float64 = -2.25
                                        _ = v29
                                        _ = v32
                                    }
                                }
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                                v2 = (v2 ^ (v2 << ((27) & 63)))
                            }
                            v33 := GenS3{F0: 1, F1: 1}
                            _ = v33
                        }
                        v5[-512] = false
                    } else {
                        var v34 int64 = 1
                        _ = v34
                    }
                }
                var v35 uint64 = (v11 & 0)
                if (!(true || true)) {
                    v11 -= (v11 / ((v35) | 1))
                    delete(v1, fi19)
                }
                switch "hello" {
                case "":
                    v36 := GenS1{F0: 65535, F1: 1000}
                    _ = v36
                default:
                    if (0.0 == (3.14159 + (1e10 + 0.0))) {
                        if (18446744073709551615 >= (v35 << ((265) & 63))) {
                            switch -2147483648 {
                            case 32767:
                                v3 = append(v3, (("hello" + (("\n\t" + ((("" + "hello") + "foo bar") + "")) + "\n\t")) + (("" + "foo bar") + ("" + "a"))))
                            default:
                                var v37 float32 = float32(1.5)
                                _ = v37
                            }
                            if len(v3) > 0 {
                                v3[0] = (("\n\t" + "a") + "foo bar")
                            }
                            if false {
                                v35 += v35
                                fmt.Println(v11)
                                v2 = 4294967295
                            } else {
                                v38 := GenS4{F0: 438, F1: 1}
                                fmt.Println(v12, v2)
                                if true {
                                    v11 -= v35
                                    v39 := GenS3{F0: -376, F1: -2147483648}
                                    _ = v39
                                }
                                _ = v38
                            }
                        } else {
                            v35 += v35
                        }
                    }
                }
                switch "\n\t" {
                case "hello":
                    v40 := v1[2147483647]
                    _ = v40
                case "":
                    var v41 int = fi19
                    _ = v41
                case "a":
                    switch "" {
                    case "\n\t":
                        v1[(fi18 % (((-(fi18))) | 1))] = ""
                    }
                    delete(v1, fi18)
                    switch fi18 {
                    case 2147483647:
                        v12 = -1
                    case 4294967295:
                        switch ("foo bar" + "a") {
                        case "hello":
                            v12 -= (v12 | (v12 ^ 42))
                        }
                    default:
                        fmt.Println(v12, v12)
                    }
                }
                if (0.1 != (0.1 * ((1e10 * 1.5) - (0.1 * -2.25)))) {
                    v42 := GenS1{F0: 255, F1: (v12 + v12)}
                    v2 = (v2 ^ 2147483648)
                    var v43 uint = 4294967295
                    _ = v42
                    _ = v43
                }
                if (!(int64(v2) < 1)) {
                    v11 = (^(v11))
                    if (!(true || true)) {
                        v35 -= v35
                        fmt.Println(fi18, v12, fi18)
                    }
                }
                v12 -= 42
                v44 := GenS5{F0: -1, F1: 32768}
                _ = fi19
                _ = v35
                _ = v44
            }
        case -654:
            v2 += v2
        default:
            v11 += (v11 >> ((42) & 63))
        }
        v2 -= (v2 | (v2 >> ((18446744073709551615) & 63)))
        _ = fi18
    }
    v12 = (v12 % (((v12 / ((-968) | 1))) | 1))
    switch ((("" + (("\n\t" + ((("foo bar" + "foo bar") + "hello") + "a")) + "")) + ("\n\t" + (((("\n\t" + "foo bar") + "a") + "") + ""))) + "foo bar") {
    case "\n\t":
        if (42 < 9223372036854775807) {
            if ((true && true) || (!true)) {
                if (true && ((true || true) || false)) {
                    if (!(42 <= 42)) {
                        for fi45 := 0; fi45 < 5; fi45++ {
                            switch "hello" {
                            case "\n\t":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "foo bar":
                                var v46 bool = false
                                _ = v46
                            default:
                                var v47 int64 = 152
                                _ = v47
                            }
                            _ = fi45
                        }
                        if (32767 > -9223372036854775808) {
                            var v48 uint16 = 0
                            v49 := GenS6{F0: 48}
                            fmt.Println(v48, v2)
                            _ = v48
                            _ = v49
                        } else {
                            var v50 float32 = (1e10 - 0.0)
                            delete(v5, -387)
                            _ = v50
                        }
                    }
                    var v51 float64 = (0.1 / 1.0)
                    v11 = v11
                    _ = v51
                } else {
                    for fi52 := 0; fi52 < 15; fi52++ {
                        v3 = append(v3, "")
                        var v53 int = fi52
                        if true {
                            for fi54 := 0; fi54 < 19; fi54++ {
                                switch int(v11) {
                                case 1:
                                    var v55 uint16 = 32767
                                    _ = v55
                                default:
                                    var v56 float64 = (0.1 * (0.0 - float64((-2.25 + 3.14159))))
                                    _ = v56
                                }
                                _ = fi54
                            }
                            if true {
                                for fi57 := 0; fi57 < 3; fi57++ {
                                    v3 = append(v3, "\n\t")
                                    _ = fi57
                                }
                                var v58 int64 = 42
                                v53 = int(v12)
                                _ = v58
                            } else {
                                v2 += v2
                                if (1 <= 1) {
                                    delete(v1, (-(v53)))
                                    v59 := GenS7{F0: 42, F1: (!(true && true))}
                                    v11 -= 18446744073709551615
                                    _ = v59
                                } else {
                                    v12 = 32767
                                    if (42 != -83) {
                                        var v60 bool = true
                                        v60 = (!(true && ((v11 / ((v11) | 1)) > 42)))
                                        var v61 int16 = (^(v12))
                                        _ = v60
                                        _ = v61
                                    }
                                }
                                fmt.Println(v12, fi52)
                            }
                            var v62 uint64 = v11
                            _ = v62
                        } else {
                            v11 = v11
                            v4[(^(v2))] = "hello"
                            v2 -= (v2 << ((997) & 63))
                        }
                        _ = fi52
                        _ = v53
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v12)
                }
                var v63 uint8 = 42
                _ = v63
            } else {
                for fi64 := 0; fi64 < 16; fi64++ {
                    v3 = append(v3, "hello")
                    _ = fi64
                }
            }
            switch "a" {
            case "a":
                if true {
                    if true {
                        var v65 int8 = 42
                        _ = v65
                    } else {
                        if (true && false) {
                            for fi66 := 0; fi66 < 2; fi66++ {
                                if (!(("hello" + "a") >= "\n\t")) {
                                    if (!(true && false)) {
                                        v67 := GenS2{F0: 1}
                                        _ = v67
                                    } else {
                                        var v68 int16 = (v12 << ((genFn0((v11 - 4294967295), -481, (v11 >> ((4294967295) & 63)))) & 63))
                                        fmt.Println(v11, v2, fi66)
                                        v69 := GenS8{F0: "\n\t", F1: 2147483647}
                                        _ = v68
                                        _ = v69
                                    }
                                    v12 += v12
                                }
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = ("\n\t" + "\n\t")
                                    }
                                } else {
                                    v11 -= v11
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v70 uint16 = 761
                                    _ = v70
                                }
                                _ = fi66
                            }
                        }
                        switch "a" {
                        case "\n\t":
                            for fi71 := 0; fi71 < 6; fi71++ {
                                var v72 int64 = 0
                                _ = fi71
                                _ = v72
                            }
                        }
                    }
                }
            case "foo bar":
                for fi73 := 0; fi73 < 20; fi73++ {
                    var v74 int16 = (-(v12))
                    for fi75 := 0; fi75 < 2; fi75++ {
                        if (1.5 >= 0.0) {
                            v74 = (^(v74))
                            v76 := GenS1{F0: 42, F1: (v12 + 733)}
                            if (127 >= 1) {
                                v2 -= (^(v2))
                            }
                            _ = v76
                        }
                        if true {
                            if true {
                                if (!(("foo bar" + "foo bar") >= ("hello" + ("foo bar" + "foo bar")))) {
                                    var v77 float64 = 1e10
                                    v12 += v12
                                    _ = v77
                                }
                                v78 := GenS6{F0: 1}
                                var v79 float32 = 0.0
                                _ = v78
                                _ = v79
                            } else {
                                var v80 int = 32767
                                _ = v80
                            }
                            switch int(v11) {
                            case 32767:
                                var v81 uint8 = 2
                                _ = v81
                            case 1:
                                delete(v1, 421)
                            }
                            var v82 string = "\n\t"
                            _ = v82
                        } else {
                            v12 += (v74 / ((int16(v11)) | 1))
                        }
                        if (v12 > v74) {
                            v83 := GenS5{F0: 1, F1: fi75}
                            _ = v83
                        }
                        _ = fi75
                    }
                    v2 = v2
                    _ = fi73
                    _ = v74
                }
            case "\n\t":
                var v84 float64 = 1.5
                _ = v84
            default:
                _ = len(v3)
            }
        } else {
            v12 += 32767
            var v85 float32 = -2.25
            for fi86 := 0; fi86 < 7; fi86++ {
                switch ("foo bar" + ("" + "hello")) {
                case "foo bar":
                    v87 := GenS9{F0: 1, F1: 1}
                    _ = v87
                }
                _ = fi86
            }
            _ = v85
        }
    }
    v88 := GenS10{F0: true}
    if (((false || true) && true) || (65535 == 0)) {
        v2 = v2
        v89 := GenS11{F0: 65535, F1: ""}
        _ = v89
    } else {
        var v90 int16 = v12
        v91 := GenS12{F0: (-699 != 128)}
        _ = v90
        _ = v91
    }
    for fi92 := 0; fi92 < 14; fi92++ {
        v2 += v2
        switch 4294967295 {
        case 630:
            v11 = v11
        case -1:
            for fi93 := 0; fi93 < 9; fi93++ {
                switch ("foo bar" + "hello") {
                case "\n\t":
                    v94 := GenS4{F0: 67, F1: 253}
                    _ = v94
                case "":
                    fmt.Println(v12, fi93)
                }
                v95 := GenS13{F0: 2, F1: 1}
                _ = fi93
                _ = v95
            }
        case 2147483647:
            if ((false || true) && true) {
                _ = len(v3)
                v12 = v12
            }
        default:
            v2 = (v2 % ((v2) | 1))
        }
        _ = fi92
    }
    var v96 int8 = 127
    v2 += (v2 + v2)
    _ = len(v3)
    var v97 int = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v88
    _ = v96
    _ = v97
}
