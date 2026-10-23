package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int32; F1 int64; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string; F1 int64; F2 uint32 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 string; F1 bool; F2 int }
func genFn0() int8 {
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    var v8 uint16 = 2
    v9 := GenS1{F0: 42, F1: 0, F2: 665}
    _ = v8
    _ = v9
    return 19
}

func main() {
    var v1 int64 = 823
    var v2 float32 = (-2.25 / (0.1 / (1e10 * (0.1 / 1.0))))
    var v3 int = -9223372036854775808
    v4 := []string{((("" + "hello") + "hello") + ""), "a", ""}
    v5 := make(map[uint32]int)
    var v10 int16 = 32767
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if ((-(v1)) <= -1) {
        var v11 int = -1
        _ = v11
    }
    _ = len(v5)
    var v12 int8 = -115
    v12 += genFn0()
    var v13 int32 = 1
    _ = len(v4)
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            switch v3 {
            case 1:
                v16 := GenS2{F0: 2147483648, F1: genFn0()}
                _ = v16
            case 32768:
                v10 -= -32768
            case -2147483648:
                v17 := GenS2{F0: 743, F1: 42}
                _ = v17
            default:
                if (0 > 2) {
                    if (false || (1 == 2)) {
                        fmt.Println(fi15, v10)
                        fmt.Println(v10, v3, v2)
                        if (32767 < 357) {
                            var v18 int32 = v13
                            _ = v18
                        }
                    } else {
                        if true {
                            var v19 float64 = (((3.14159 / float64(1.5)) * 1.5) * (1.5 * 1e10))
                            _ = v19
                        } else {
                            if true {
                                var v20 uint16 = 2
                                _ = v20
                            } else {
                                v21 := GenS2{F0: 4294967295, F1: -1}
                                var v22 uint64 = 4294967295
                                var v23 bool = false
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            _ = len(v4)
                            v24 := v5[2147483647]
                            _ = v24
                        }
                        if (false || true) {
                            switch fi14 {
                            case -638:
                                if false {
                                    if false {
                                        fmt.Println(fi14, v2)
                                        v1 -= (^(v1))
                                    }
                                    v25 := GenS3{F0: (v13 << ((471) & 63)), F1: v1, F2: 930}
                                    v26 := v5[397]
                                    v27 := GenS2{F0: 2, F1: -128}
                                    v26 -= 128
                                    var v28 bool = true
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                } else {
                                    if false {
                                        _ = len(v4)
                                        switch v3 {
                                        case -422:
                                            var v29 uint8 = 2
                                            _ = v29
                                        case -334:
                                            v4 = append(v4, ("a" + "\n\t"))
                                        }
                                        if ((v2 == 0.1) || true) {
                                            v2 -= (float32((((-2.25 - 0.0) - (float64(0.0) * 0.0)) * (3.14159 + 1e10))) - v2)
                                            v5[459] = (v3 & v3)
                                            switch ("a" + ("hello" + (("foo bar" + "foo bar") + ("" + ("" + "\n\t"))))) {
                                            case "hello":
                                                if false {
                                                    if true {
                                                        if len(v4) > 0 {
                                                            v4[0] = ""
                                                        }
                                                    } else {
                                                        v30 := GenS4{F0: 1}
                                                        fmt.Println(v2, v1, fi14)
                                                        switch "\n\t" {
                                                        case "a":
                                                            if (true || false) {
                                                                v10 -= v10
                                                                fmt.Println(v10, v12, v2)
                                                            } else {
                                                                if ((("foo bar" + "hello") + ("" + "")) > "\n\t") {
                                                                    fmt.Println(fi14, fi15, fi14)
                                                                    v2 += (v2 - float32((3.14159 - 1e10)))
                                                                } else {
                                                                    if (!(0 != v10)) {
                                                                        v2 -= float32(1e10)
                                                                        var v31 float64 = -2.25
                                                                        fmt.Println(v12, v1)
                                                                        _ = v31
                                                                    }
                                                                }
                                                                v13 -= v13
                                                                var v32 uint16 = 881
                                                                _ = v32
                                                            }
                                                        case "foo bar":
                                                            fmt.Println(v13, v12, v1)
                                                        default:
                                                            if true {
                                                                v4 = append(v4, "hello")
                                                            } else {
                                                                v33 := GenS5{F0: (v13 * v13)}
                                                                _ = v33
                                                            }
                                                        }
                                                        var v34 uint64 = 4294967295
                                                        var v35 uint64 = (v34 ^ (v34 & v34))
                                                        _ = v30
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                } else {
                                                    var v36 uint32 = 573
                                                    fmt.Println(fi14)
                                                    var v37 float32 = (v2 + 0.0)
                                                    if (!(42 <= 1)) {
                                                        fmt.Println(v2, v36, v3)
                                                    }
                                                    switch fi14 {
                                                    case 792:
                                                        v10 -= v10
                                                    default:
                                                        v38 := GenS3{F0: v13, F1: v1, F2: 650}
                                                        _ = v38
                                                    }
                                                    _ = v36
                                                    _ = v37
                                                }
                                            default:
                                                v39 := GenS4{F0: 4294967295}
                                                _ = v39
                                            }
                                        } else {
                                            v3 += (fi14 >> ((2) & 63))
                                            v40 := GenS1{F0: (v13 | v13), F1: 318, F2: 1}
                                            switch (v3 - (^(fi14))) {
                                            case 32767:
                                                v2 = (1.5 / v2)
                                                if (((("" + (("a" + (("hello" + "foo bar") + ("\n\t" + "hello"))) + "hello")) + "foo bar") + "hello") != ("hello" + ("\n\t" + (("\n\t" + "hello") + ("" + "\n\t"))))) {
                                                    delete(v5, 4294967295)
                                                    v10 -= v10
                                                }
                                            case 32768:
                                                if (false && (185 > v10)) {
                                                    if true {
                                                        v12 -= v12
                                                    } else {
                                                        var v41 float32 = (float32(float64((3.14159 / float32(0.1)))) * 0.0)
                                                        if len(v4) > 0 {
                                                            v4[0] = ("" + "foo bar")
                                                        }
                                                        v42 := GenS6{F0: ("foo bar" + "\n\t"), F1: (v1 >> ((2) & 63)), F2: 42}
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = ("foo bar" + "hello")
                                                    }
                                                }
                                            default:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            var v43 int16 = (v10 / ((v10) | 1))
                                            _ = v40
                                            _ = v43
                                        }
                                        v44 := GenS7{F0: v10, F1: (^(v10))}
                                        if ((1e10 - (v2 / v2)) >= 0.0) {
                                            switch "foo bar" {
                                            case "hello":
                                                if true {
                                                    switch fi15 {
                                                    case -9223372036854775808:
                                                        v45 := GenS4{F0: 42}
                                                        _ = v45
                                                    case -1:
                                                        fmt.Println(v12, v13, v3)
                                                    case -77:
                                                        fmt.Println(v10, v2, v3)
                                                    }
                                                }
                                            default:
                                                var v46 uint8 = 127
                                                _ = v46
                                            }
                                        }
                                        _ = v44
                                    }
                                    v47 := GenS5{F0: (-(v13))}
                                    fmt.Println(v13)
                                    if true {
                                        if true {
                                            switch 127 {
                                            case -2147483648:
                                                var v48 int32 = v13
                                                _ = v48
                                            case 646:
                                                var v49 uint32 = 65
                                                _ = v49
                                            }
                                            v12 = 42
                                            switch (fi14 ^ fi14) {
                                            case 9223372036854775807:
                                                var v50 uint8 = 228
                                                _ = v50
                                            case 278:
                                                v1 += (v1 - 809)
                                            case 1:
                                                fmt.Println(v13, fi14)
                                            }
                                            var v51 bool = false
                                            _ = v51
                                        }
                                        fmt.Println(fi14, v13)
                                    }
                                    var v52 float64 = ((-2.25 - 0.1) / -2.25)
                                    _ = v47
                                    _ = v52
                                }
                            case 127:
                                switch ("\n\t" + ("a" + "foo bar")) {
                                case "":
                                    _ = len(v4)
                                case "a":
                                    v5[793] = fi14
                                default:
                                    var v53 uint8 = 1
                                    fmt.Println(v53)
                                    _ = v53
                                }
                                if (44 < 1) {
                                    var v54 int = 614
                                    _ = v54
                                } else {
                                    v10 -= (^(v10))
                                }
                            case 42:
                                v4 = append(v4, "")
                            }
                            var v55 int64 = (v1 << ((65535) & 63))
                            v5[1] = (^(fi15))
                            _ = v55
                        }
                        v56 := GenS7{F0: v10, F1: 397}
                        if (!(v1 >= 144)) {
                            switch ((("a" + "foo bar") + "a") + "") {
                            case "\n\t":
                                var v57 uint32 = 1
                                _ = v57
                            default:
                                v3 -= v3
                            }
                            v13 = (v13 ^ v13)
                            fmt.Println(fi15, v3)
                        }
                        _ = v56
                    }
                } else {
                    v58 := v5[0]
                    if true {
                        v59 := GenS3{F0: 1, F1: (v1 << ((454) & 63)), F2: 717}
                        var v60 float64 = 3.14159
                        _ = v59
                        _ = v60
                    }
                    _ = v58
                }
                v13 = int32(v12)
                fmt.Println(fi14, v12)
            }
            if true {
                delete(v5, 40)
            } else {
                if false {
                    fmt.Println(fi15)
                }
                fmt.Println(v12, v2)
                v13 = -197
                v12 = (v12 << ((18446744073709551615) & 63))
            }
            v3 = v3
            _ = fi15
        }
        var v61 uint32 = 2147483647
        _ = fi14
        _ = v61
    }
    for fi62 := 0; fi62 < 5; fi62++ {
        var v63 uint32 = 4294967295
        _ = len(v5)
        _ = fi62
        _ = v63
    }
    for fi64 := 0; fi64 < 5; fi64++ {
        var v65 int = (fi64 << ((65535) & 63))
        if true {
            fmt.Println(v2)
            switch 9223372036854775807 {
            case 4294967295:
                fmt.Println(v65, fi64, v2)
            }
            if false {
                fmt.Println(v13)
                if (v13 != 42) {
                    v66 := GenS8{F0: ("hello" + ("foo bar" + ("a" + ("a" + (("" + "foo bar") + ("hello" + "hello")))))), F1: (v12 <= 100), F2: fi64}
                    _ = v66
                } else {
                    var v67 uint64 = 42
                    v67 -= (v67 & 42)
                    _ = v67
                }
                v13 += v13
            }
        }
        _ = fi64
        _ = v65
    }
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
