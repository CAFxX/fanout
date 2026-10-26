package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 string }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 string; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8; F1 uint16; F2 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int8; F1 int8; F2 uint64 }
func genFn0(p0 string, p1 bool) int8 {
    var v7 float32 = ((3.14159 / 1.5) * (((1e10 + 3.14159) + -2.25) / 3.14159))
    v8 := []int{9223372036854775807, 32768, -677, 767}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -12
}
func genFn1() string {
    switch "a" {
    case "\n\t":
        if (!((false && (true || (!(51 >= 18446744073709551615)))) || (198 <= 2))) {
            v9 := []int{32767, 44, 128, -789}
            _ = v9
        } else {
            var v10 int = 128
            var v11 uint64 = 206
            var v12 int = (v10 / ((v10) | 1))
            _ = v10
            _ = v11
            _ = v12
        }
    case "":
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 bool = true
            _ = fi13
            _ = v14
        }
    default:
        var v15 float32 = -2.25
        _ = v15
    }
    if (266 < 65535) {
        var v16 bool = false
        v16 = v16
        _ = v16
    }
    return ("\n\t" + ("" + ("a" + "foo bar")))
}

func main() {
    var v1 int = 627
    var v2 uint8 = 255
    v3 := []uint32{42}
    var v4 uint16 = 32767
    v5 := make(map[int]bool)
    var v6 int32 = 42
    v4 = (v4 - (v4 - 255))
    v2 = v2
    if true {
        var v17 bool = ((false && (!true)) && true)
        var v18 float32 = ((0.0 * (1e10 - (((0.0 + 1.5) / float32((1e10 / 1e10))) + 0.1))) * 0.0)
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 15; fi19++ {
            _ = len(v5)
            var v20 uint16 = v4
            for fi21 := 0; fi21 < 15; fi21++ {
                var v22 int8 = 127
                var v23 int32 = v6
                v24 := GenS1{F0: v2, F1: 4294967295, F2: "\n\t"}
                v5[fi21] = true
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            _ = fi19
            _ = v20
        }
        switch (v1 >> ((255) & 63)) {
        case 4294967295:
            v6 -= (v6 - (v6 * int32(v2)))
        }
        switch 42 {
        case 127:
            if ((v1 < v1) || (v4 < 0)) {
                v25 := GenS1{F0: (v2 >> ((2) & 63)), F1: 2147483648, F2: ((("" + "hello") + (("\n\t" + "\n\t") + "")) + ("\n\t" + ((("foo bar" + "") + "hello") + (("\n\t" + ("foo bar" + "")) + ("foo bar" + "\n\t")))))}
                if (("\n\t" + "foo bar") <= "\n\t") {
                    v26 := GenS1{F0: v2, F1: 65535, F2: "foo bar"}
                    var v27 uint8 = (v2 + (v2 >> ((769) & 63)))
                    _ = v26
                    _ = v27
                }
                v6 = (-(v6))
                _ = v25
            }
        case -1:
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "hello" {
                case "hello":
                    var v28 uint16 = v4
                    _ = v28
                }
            }
        }
    }
    if (42 >= 692) {
        var v29 uint64 = 42
        delete(v5, v1)
        switch v1 {
        case 0:
            fmt.Println(v6, v1, v2)
        case 127:
            var v30 int32 = (-(v6))
            _ = v30
        case 4294967295:
            v1 -= v1
        }
        _ = v29
    } else {
        for fi31 := 0; fi31 < 8; fi31++ {
            var v32 int = 9223372036854775807
            switch fi31 {
            case 560:
                for fi33 := 0; fi33 < 19; fi33++ {
                    var v34 int16 = 1
                    var v35 uint32 = 479
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            default:
                v32 += (v32 * 2147483647)
            }
            for fi36 := 0; fi36 < 15; fi36++ {
                var v37 bool = (255 >= (^(v2)))
                _ = len(v3)
                var v38 bool = v37
                var v39 int8 = genFn0(("\n\t" + "\n\t"), (!false))
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
            }
            _ = fi31
            _ = v32
        }
        v6 -= v6
    }
    var v40 float32 = 1.5
    v6 = 1
    if (!false) {
        for fi41 := 0; fi41 < 7; fi41++ {
            fmt.Println(fi41, v40, v1)
            v42 := GenS2{F0: -32768, F1: uint32(v2)}
            v6 -= 42
            _ = fi41
            _ = v42
        }
    } else {
        var v43 uint16 = 2
        fmt.Println(v40)
        _ = v43
    }
    v2 += 19
    if false {
        var v44 uint64 = 2
        _ = v44
    }
    v45 := GenS3{F0: "foo bar", F1: v4, F2: uint32(v4)}
    for fi46 := 0; fi46 < 7; fi46++ {
        if false {
            var v47 uint64 = 4294967295
            _ = v47
        }
        v48 := GenS4{F0: 508, F1: v2}
        v4 = 32767
        _ = fi46
        _ = v48
    }
    fmt.Println(v4)
    delete(v5, v1)
    if true {
        v49 := GenS5{F0: (v4 ^ (v4 & (v4 | (v4 + 65535)))), F1: 65535}
        if false {
            if true {
                fmt.Println(v6, v4)
            } else {
                v6 = 1
            }
            switch "hello" {
            case "foo bar":
                switch "" {
                case "a":
                    var v50 uint16 = v4
                    _ = v50
                case "\n\t":
                    var v51 float64 = float64(0.1)
                    _ = v51
                case "foo bar":
                    var v52 int = (v1 - (v1 - v1))
                    _ = v52
                default:
                    if (1 > 685) {
                        v2 += 130
                        switch v1 {
                        case 127:
                            switch "hello" {
                            case "hello":
                                var v53 float64 = -2.25
                                _ = v53
                            case "foo bar":
                                v54 := GenS3{F0: ("\n\t" + ("\n\t" + "a")), F1: v4, F2: 2}
                                _ = v54
                            }
                        default:
                            v4 = uint16(v2)
                        }
                    }
                }
            case "a":
                fmt.Println(v4, v4)
            }
            fmt.Println(v4, v40)
        }
        v55 := GenS6{F0: v6}
        _ = v49
        _ = v55
    } else {
        var v56 int = -9223372036854775808
        switch v56 {
        case -463:
            var v57 int = v1
            _ = v57
        case 4294967295:
            for fi58 := 0; fi58 < 15; fi58++ {
                switch -2147483648 {
                case 9223372036854775807:
                    v59 := GenS6{F0: (v6 - v6)}
                    _ = v59
                }
                v60 := GenS2{F0: 42, F1: 689}
                fmt.Println(v6, v1)
                _ = fi58
                _ = v60
            }
        }
        _ = v56
    }
    if false {
        if (false || (42 != 2)) {
            if (100 != -1) {
                v4 = v4
                if true {
                    v5[(v1 - (v1 % ((v1) | 1)))] = (-2.25 > float32(1e10))
                }
                v61 := GenS5{F0: (v4 - v4), F1: 1}
                _ = v61
            } else {
                v1 = 2147483647
            }
        }
        v6 = 2147483647
        if false {
            if (-128 > 0) {
                _ = len(v3)
                var v62 int16 = 0
                _ = v62
            }
            switch -9223372036854775808 {
            case -858:
                delete(v5, v1)
            case 42:
                switch "hello" {
                case "foo bar":
                    var v63 int32 = int32(v4)
                    _ = v63
                case "":
                    for fi64 := 0; fi64 < 2; fi64++ {
                        v65 := GenS1{F0: 245, F1: 42, F2: "\n\t"}
                        v66 := GenS7{F0: 1, F1: (v4 >> ((334) & 63)), F2: "a"}
                        v4 -= v4
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                case "\n\t":
                    v4 -= (v4 / ((v4) | 1))
                default:
                    var v67 uint32 = uint32(v2)
                    _ = v67
                }
            }
        }
    } else {
        v40 = ((v40 + v40) + v40)
        if ("\n\t" > "foo bar") {
            for fi68 := 0; fi68 < 1; fi68++ {
                v69 := v5[(fi68 & (fi68 << ((1) & 63)))]
                for fi70 := 0; fi70 < 6; fi70++ {
                    v4 = v4
                    switch ((("a" + ("\n\t" + "foo bar")) + "a") + "foo bar") {
                    case "":
                        v6 += 449
                    case "\n\t":
                        v71 := GenS4{F0: (v6 - v6), F1: v2}
                        _ = v71
                    case "hello":
                        var v72 int16 = 0
                        _ = v72
                    }
                    _ = fi70
                }
                _ = fi68
                _ = v69
            }
        } else {
            if ((v2 ^ v2) < v2) {
                for fi73 := 0; fi73 < 5; fi73++ {
                    v1 = 141
                    _ = fi73
                }
                switch -644 {
                case -1:
                    v40 += -2.25
                case 1:
                    if (!((1000 > -707) && ("\n\t" >= ""))) {
                        v74 := v5[(v1 * 561)]
                        _ = v74
                    } else {
                        if true {
                            if false {
                                v4 = (v4 & v4)
                                for fi75 := 0; fi75 < 1; fi75++ {
                                    if (18446744073709551615 != 551) {
                                        var v76 int = (fi75 ^ (fi75 + fi75))
                                        v2 = (^(v2))
                                        v77 := GenS1{F0: v2, F1: 2147483647, F2: (("a" + "") + ("foo bar" + ""))}
                                        _ = v76
                                        _ = v77
                                    } else {
                                        fmt.Println(v40, v40)
                                        v78 := GenS4{F0: v6, F1: v2}
                                        var v79 int8 = 103
                                        _ = v78
                                        _ = v79
                                    }
                                    v3 = append(v3, uint32(v4))
                                    _ = fi75
                                }
                            }
                        } else {
                            for fi80 := 0; fi80 < 15; fi80++ {
                                switch "" {
                                case "foo bar":
                                    var v81 int8 = 42
                                    _ = v81
                                case "\n\t":
                                    var v82 int32 = 0
                                    _ = v82
                                case "":
                                    fmt.Println(v4, v1)
                                default:
                                    v83 := GenS1{F0: 36, F1: 1, F2: "foo bar"}
                                    _ = v83
                                }
                                for fi84 := 0; fi84 < 4; fi84++ {
                                    v85 := GenS8{F0: 42}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if ((9223372036854775807 > (fi84 >> ((195) & 63))) || (!(844 < 1))) {
                                        if (1 < 626) {
                                            v86 := GenS7{F0: 1, F1: v4, F2: "foo bar"}
                                            var v87 uint = 0
                                            v88 := GenS9{F0: 44, F1: 0, F2: 980}
                                            fmt.Println(v4)
                                            _ = v86
                                            _ = v87
                                            _ = v88
                                        } else {
                                            fmt.Println(v1, v6, v40)
                                        }
                                        fmt.Println(v1, v2)
                                    }
                                    _ = fi84
                                    _ = v85
                                }
                                _ = fi80
                            }
                            for fi89 := 0; fi89 < 8; fi89++ {
                                if true {
                                    fmt.Println(v4)
                                    if ((!false) && true) {
                                        var v90 int32 = -2147483648
                                        delete(v5, -310)
                                        var v91 uint = 4294967295
                                        _ = v90
                                        _ = v91
                                    }
                                } else {
                                    var v92 uint16 = v4
                                    _ = v92
                                }
                                _ = fi89
                            }
                        }
                        v4 += 65535
                    }
                default:
                    v93 := v5[(v1 % (((v1 >> ((523) & 63))) | 1))]
                    _ = v93
                }
                var v94 uint64 = 120
                _ = v94
            } else {
                _ = len(v3)
            }
        }
    }
    var v95 float32 = float32(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v45
    _ = v5
    _ = v6
    _ = v95
}
