package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 bool; F2 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool; F1 int64; F2 string }
type GenS10 struct { F0 bool; F1 bool; F2 int8 }

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]bool)
    v3 := []uint32{1, 65535}
    v4 := make(map[int]string)
    fmt.Println(v1)
    fmt.Println(v1)
    _ = len(v3)
    if false {
        fmt.Println(v1)
        if ((true || ((32767 <= 65535) || (true && false))) || true) {
            var v5 float64 = -2.25
            _ = v5
        } else {
            _ = len(v3)
            var v6 int32 = -1
            if (-2.25 <= 0.1) {
                if false {
                    if (!true) {
                        if (!(!((float64(1.5) / 1.0) < 0.1))) {
                            v7 := GenS1{F0: 42}
                            if ((float32(-2.25) / (float32(1.5) + 0.1)) <= (float32(1.5) * (1e10 / 1.0))) {
                                _ = len(v3)
                                var v8 uint = 747
                                _ = v8
                            } else {
                                v9 := GenS1{F0: 927}
                                var v10 int64 = -9223372036854775808
                                _ = v10
                                _ = v9
                            }
                            _ = v7
                        } else {
                            var v11 int8 = 127
                            _ = v11
                        }
                    } else {
                        if true {
                            fmt.Println(v6, v1)
                        } else {
                            var v12 uint = 255
                            v1 = 586
                            switch -520 {
                            case 186:
                                var v13 bool = (979 <= 958)
                                _ = v13
                            case -338:
                                v6 = -1
                            }
                            _ = v12
                        }
                    }
                    v6 = (v6 >> ((2) & 63))
                }
                if (91 < 127) {
                    v1 = 1
                    for fi14 := 0; fi14 < 14; fi14++ {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            v16 := GenS2{F0: 1, F1: 0}
                            v17 := GenS2{F0: 1, F1: 0}
                            fmt.Println(fi14)
                            v18 := GenS1{F0: 2}
                            if ((-2.25 != 0.1) || true) {
                                v1 += 1
                            } else {
                                if ((!false) && true) {
                                    if (v1 > v1) {
                                        v19 := GenS1{F0: 781}
                                        switch "hello" {
                                        case "\n\t":
                                            switch -207 {
                                            case -236:
                                                v6 += int32(fi14)
                                            }
                                        case "":
                                            fmt.Println(fi14)
                                        default:
                                            switch fi14 {
                                            case 4294967295:
                                                switch (fi15 | 42) {
                                                case 2147483647:
                                                    fmt.Println(v1)
                                                case 32767:
                                                    v1 = v1
                                                case 127:
                                                    v6 += -1
                                                    var v20 uint64 = (v1 + v1)
                                                    v4[-1] = "a"
                                                    _ = v20
                                                default:
                                                    var v21 uint16 = 42
                                                    _ = v21
                                                }
                                                var v22 uint64 = (v1 ^ 2)
                                                _ = v22
                                            case 127:
                                                switch (fi15 + -1) {
                                                case 128:
                                                    var v23 int64 = -135
                                                    if true {
                                                        fmt.Println(v23, v1)
                                                        var v24 uint64 = v1
                                                        _ = v24
                                                    } else {
                                                        if ("\n\t" >= ("hello" + "hello")) {
                                                            if ((!false) && false) {
                                                                v23 = (^(v23))
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = 66
                                                                }
                                                                var v25 float32 = 1.5
                                                                v23 += -777
                                                                _ = v25
                                                            }
                                                            v3 = append(v3, 2)
                                                            if true {
                                                                if ((^(fi14)) >= (fi14 % (((fi14 ^ (fi14 / (((fi14 / ((774) | 1))) | 1)))) | 1))) {
                                                                    var v26 string = "foo bar"
                                                                    v23 = v23
                                                                    _ = v26
                                                                }
                                                                if true {
                                                                    var v27 uint64 = 835
                                                                    v28 := v4[(fi15 + (fi15 / ((fi15) | 1)))]
                                                                    var v29 uint64 = 18446744073709551615
                                                                    _ = v27
                                                                    _ = v28
                                                                    _ = v29
                                                                } else {
                                                                    var v30 string = "foo bar"
                                                                    v31 := GenS3{F0: 42}
                                                                    _ = v30
                                                                    _ = v31
                                                                }
                                                            } else {
                                                                v23 = -1
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    fmt.Println(fi15, v6)
                                                                default:
                                                                    v6 = (v6 & -942)
                                                                }
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = len(v3)
                                                    fmt.Println(v1)
                                                    if (31 != 0) {
                                                        v1 += v1
                                                        v23 = (v23 | (v23 + (^(v23))))
                                                        fmt.Println(fi15)
                                                    }
                                                    var v32 int = 2147483647
                                                    fmt.Println(v1, fi14, fi14)
                                                    _ = v23
                                                    _ = v32
                                                case 999:
                                                    v33 := GenS4{F0: 100}
                                                    _ = v33
                                                default:
                                                    v1 += 4294967295
                                                }
                                            case -9223372036854775808:
                                                switch "a" {
                                                case "":
                                                    var v34 uint16 = 1
                                                    _ = v34
                                                default:
                                                    switch "a" {
                                                    case "":
                                                        switch "a" {
                                                        case "hello":
                                                            fmt.Println(v6, v1, fi14)
                                                        case "\n\t":
                                                            var v35 int = (-(fi15))
                                                            _ = v35
                                                        default:
                                                            switch "\n\t" {
                                                            case "hello":
                                                                if (1e10 >= 3.14159) {
                                                                    v36 := GenS3{F0: 0}
                                                                    v6 = (v6 - (v6 ^ v6))
                                                                    _ = v36
                                                                } else {
                                                                    switch fi15 {
                                                                    case 744:
                                                                        var v37 int8 = 1
                                                                        _ = v37
                                                                    case 42:
                                                                        fmt.Println(v1, fi15)
                                                                    case 32767:
                                                                        fmt.Println(v1)
                                                                    }
                                                                    switch -9223372036854775808 {
                                                                    case -9223372036854775808:
                                                                        v1 += v1
                                                                    case -851:
                                                                        v38 := GenS5{F0: ("a" + ((("foo bar" + (("" + "\n\t") + "foo bar")) + (("a" + ("" + ("hello" + "foo bar"))) + "a")) + "")), F1: 0}
                                                                        _ = v38
                                                                    default:
                                                                        v1 -= v1
                                                                    }
                                                                }
                                                                v1 += 230
                                                            }
                                                        }
                                                    case "\n\t":
                                                        if len(v3) > 0 {
                                                            v3[0] = 0
                                                        }
                                                    case "a":
                                                        if len(v3) > 0 {
                                                            v3[0] = 456
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        var v39 int32 = -213
                                        _ = v19
                                        _ = v39
                                    } else {
                                        var v40 int16 = -32768
                                        _ = v40
                                    }
                                    v41 := GenS3{F0: 42}
                                    v42 := GenS6{F0: 661}
                                    _ = v41
                                    _ = v42
                                } else {
                                    var v43 int16 = -927
                                    _ = v43
                                }
                                if (42 == 65535) {
                                    var v44 uint16 = 834
                                    var v45 int = fi14
                                    _ = v44
                                    _ = v45
                                }
                                switch 9223372036854775807 {
                                case 32767:
                                    delete(v4, int(v6))
                                case -296:
                                    _ = len(v4)
                                default:
                                    if true {
                                        if (v6 > v6) {
                                            switch (^(fi15)) {
                                            case 42:
                                                v6 = 0
                                            default:
                                                if ("a" <= "hello") {
                                                    var v46 bool = true
                                                    var v47 uint32 = 1
                                                    _ = v46
                                                    _ = v47
                                                }
                                            }
                                        }
                                    } else {
                                        v6 -= -511
                                    }
                                }
                            }
                            var v48 uint8 = 1
                            _ = fi15
                            _ = v16
                            _ = v17
                            _ = v18
                            _ = v48
                        }
                        _ = fi14
                    }
                } else {
                    v6 += v6
                    switch 622 {
                    case 127:
                        switch "foo bar" {
                        case "":
                            switch ("a" + ("" + "a")) {
                            case "foo bar":
                                var v49 bool = false
                                _ = v49
                            case "\n\t":
                                v1 -= (v1 + 18446744073709551615)
                            case "a":
                                _ = len(v3)
                            }
                        case "\n\t":
                            for fi50 := 0; fi50 < 15; fi50++ {
                                v51 := GenS2{F0: 127, F1: -1}
                                _ = fi50
                                _ = v51
                            }
                        case "foo bar":
                            if true {
                                for fi52 := 0; fi52 < 7; fi52++ {
                                    var v53 uint64 = 18446744073709551615
                                    _ = fi52
                                    _ = v53
                                }
                            }
                        default:
                            var v54 uint8 = 255
                            _ = v54
                        }
                    case 2147483647:
                        var v55 uint16 = 0
                        _ = v55
                    case 42:
                        v1 = v1
                    }
                    _ = len(v2)
                }
                var v56 float32 = float32(1e10)
                _ = v56
            }
            _ = v6
        }
        v1 -= 4294967295
    }
    v57 := GenS7{F0: 1, F1: (2 < 4294967295), F2: 0}
    v58 := GenS8{F0: 9223372036854775807}
    switch ("a" + "") {
    case "":
        v59 := v2[42]
        _ = v59
    default:
        if len(v3) > 0 {
            v3[0] = 2
        }
    }
    var v60 float64 = (1e10 * 0.1)
    v61 := GenS9{F0: ("foo bar" != "foo bar"), F1: 1, F2: (((("foo bar" + "\n\t") + "\n\t") + ("" + "foo bar")) + ("hello" + "\n\t"))}
    _ = len(v3)
    if (true || (32767 == 1)) {
        var v62 float32 = 0.1
        _ = v62
    }
    fmt.Println(v1)
    switch "hello" {
    case "hello":
        delete(v4, 127)
    case "a":
        if true {
            for fi63 := 0; fi63 < 13; fi63++ {
                if (98 >= 42) {
                    if (127 < 187) {
                        var v64 int8 = 42
                        fmt.Println(v60)
                        fmt.Println(v60, v1)
                        _ = v64
                    } else {
                        var v65 int32 = -728
                        v60 -= ((v60 * v60) - (float64(1e10) - ((0.1 - ((0.0 - (((v60 + v60) - v60) * (float64(0.0) / ((v60 + ((v60 / v60) + -2.25)) + v60)))) * float64(3.14159))) + v60)))
                        _ = v65
                    }
                } else {
                    for fi66 := 0; fi66 < 7; fi66++ {
                        fmt.Println(v1, v60)
                        _ = fi66
                    }
                    fmt.Println(v1, v60)
                }
                fmt.Println(v1)
                _ = fi63
            }
        }
    case "\n\t":
        v1 = (v1 % ((846) | 1))
    default:
        if true {
            switch 32767 {
            case 414:
                var v67 uint32 = 42
                _ = v67
            }
            switch 835 {
            case 127:
                if true {
                    v60 = v60
                } else {
                    if (!false) {
                        fmt.Println(v1, v60)
                    }
                }
            }
            v1 += (v1 + v1)
        } else {
            v1 = (^(v1))
        }
    }
    delete(v4, 2147483647)
    var v68 int64 = -9223372036854775808
    var v69 uint64 = v1
    switch 32768 {
    case 1:
        v70 := GenS10{F0: true, F1: (-112 > -128), F2: 75}
        _ = v70
    case 127:
        var v71 int8 = 0
        _ = v71
    default:
        var v72 string = ("hello" + ("hello" + "foo bar"))
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v60:", v60)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
    _ = v60
    _ = v61
    _ = v68
    _ = v69
}
