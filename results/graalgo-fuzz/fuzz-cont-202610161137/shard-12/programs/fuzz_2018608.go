package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint64; F1 uint8 }
func genFn0(p0 bool) uint16 {
    var v7 int32 = 2147483647
    var v8 int16 = 0
    _ = p0
    _ = v7
    _ = v8
    return 942
}

func main() {
    v1 := []uint32{8, 65535}
    v2 := make(map[string]bool)
    var v3 float32 = (-2.25 * (1e10 - 1.5))
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    v6 := []uint32{}
    v9 := GenS1{F0: 518}
    var v10 int32 = 42
    var v11 int16 = -1
    v12 := v4["a"]
    var v13 float32 = v3
    for fi14 := 0; fi14 < 12; fi14++ {
        if len(v6) > 0 {
            v6[0] = 2
        }
        _ = fi14
    }
    v12 = 42
    v13 = 1e10
    if (true && true) {
        switch ("foo bar" + ("a" + ("a" + "foo bar"))) {
        case "":
            var v15 uint32 = 2147483647
            _ = v15
        case "\n\t":
            fmt.Println(v11, v11)
        }
        v16 := GenS1{F0: 32767}
        _ = v16
    } else {
        fmt.Println(v10, v11, v11)
        v17 := GenS1{F0: 255}
        var v18 bool = (true && false)
        _ = v17
        _ = v18
    }
    switch ("" + ("hello" + ("\n\t" + ""))) {
    case "hello":
        var v19 uint = 694
        _ = v19
    case "":
        switch "foo bar" {
        case "\n\t":
            if len(v6) > 0 {
                v6[0] = 510
            }
        case "hello":
            v20 := GenS1{F0: 255}
            _ = v20
        case "foo bar":
            fmt.Println(v10, v3)
        default:
            for fi21 := 0; fi21 < 8; fi21++ {
                fmt.Println(fi21, v3, fi21)
                _ = fi21
            }
        }
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = 0
        }
    default:
        if true {
            v3 = v13
            var v22 int16 = (v11 + v11)
            switch 2147483647 {
            case 539:
                var v23 int16 = (v22 * 0)
                _ = v23
            case 9223372036854775807:
                switch (v12 << ((573) & 63)) {
                case 378:
                    v2["a"] = (true && ("hello" >= "a"))
                case 9223372036854775807:
                    var v24 string = "hello"
                    _ = v24
                case 32768:
                    var v25 float64 = 3.14159
                    _ = v25
                default:
                    for fi26 := 0; fi26 < 15; fi26++ {
                        for fi27 := 0; fi27 < 11; fi27++ {
                            v12 -= (fi26 ^ (fi27 / ((1) | 1)))
                            v22 -= (v11 + (v11 >> ((4294967295) & 63)))
                            _ = len(v4)
                            _ = fi27
                        }
                        v28 := GenS1{F0: 65535}
                        _ = fi26
                        _ = v28
                    }
                }
            case -677:
                v3 = ((float32((float64(1e10) * float64(v13))) / 1.5) / 1.0)
            }
            _ = v22
        }
    }
    v11 -= 764
    if true {
        for fi29 := 0; fi29 < 20; fi29++ {
            v12 -= -2147483648
            v12 = v12
            _ = fi29
        }
    } else {
        switch (v12 - (v12 >> ((255) & 63))) {
        case -1:
            v12 = (^(v12))
        case -9223372036854775808:
            switch "foo bar" {
            case "\n\t":
                if (!true) {
                    v2[(("\n\t" + ("hello" + "a")) + "foo bar")] = false
                } else {
                    if true {
                        var v30 int8 = -85
                        _ = v30
                    }
                    if (float32(-2.25) <= (v3 * 0.0)) {
                        fmt.Println(v11, v3, v13)
                    } else {
                        var v31 float32 = v13
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                        _ = v31
                    }
                }
            case "foo bar":
                fmt.Println(v12)
            case "":
                var v32 int16 = 709
                _ = v32
            default:
                fmt.Println(v13)
            }
        default:
            for fi33 := 0; fi33 < 12; fi33++ {
                fmt.Println(v12)
                _ = fi33
            }
        }
        if ((98 == -128) && ((2 < 2) || (v11 >= v11))) {
            fmt.Println(v10)
            if false {
                v13 = v13
                fmt.Println(v11, v10, v11)
                if ("hello" >= "\n\t") {
                    v34 := GenS2{F0: (v10 << ((4294967295) & 63)), F1: 700, F2: 391}
                    var v35 int8 = 1
                    _ = v34
                    _ = v35
                }
            } else {
                var v36 uint32 = 65535
                v37 := GenS1{F0: 228}
                var v38 uint8 = 140
                _ = v36
                _ = v37
                _ = v38
            }
            if false {
                v39 := GenS3{F0: 18446744073709551615}
                _ = v39
            }
        }
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        if ((((v13 * float32(((float64(1.5) * 1.5) / 1.5))) * (float32(-2.25) + (3.14159 / 1.0))) / v3) < v13) {
            var v41 int8 = -4
            _ = v41
        } else {
            for fi42 := 0; fi42 < 1; fi42++ {
                var v43 int8 = 0
                var v44 int = (-(fi42))
                v43 = (-(v43))
                v45 := GenS1{F0: genFn0((553 >= 0))}
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
        }
        v10 -= 42
        if true {
            switch (("a" + "foo bar") + "") {
            case "hello":
                v10 -= v10
            case "":
                v6 = append(v6, 4294967295)
            }
            v46 := GenS4{F0: (v11 + int16(v10)), F1: 0}
            var v47 string = "\n\t"
            _ = v46
            _ = v47
        }
        _ = fi40
    }
    for fi48 := 0; fi48 < 8; fi48++ {
        if true {
            switch v12 {
            case -9223372036854775808:
                switch "a" {
                case "foo bar":
                    v3 += v13
                default:
                    for fi49 := 0; fi49 < 4; fi49++ {
                        v50 := GenS4{F0: (v11 % ((v11) | 1)), F1: 42}
                        _ = fi49
                        _ = v50
                    }
                }
            case 128:
                v13 = (1.5 - (3.14159 / (float32(0.0) * (v3 * (v3 / v3)))))
            case -2147483648:
                _ = len(v6)
            default:
                switch "\n\t" {
                case "hello":
                    for fi51 := 0; fi51 < 11; fi51++ {
                        v10 = (v10 % ((2147483647) | 1))
                        _ = fi51
                    }
                default:
                    v52 := GenS1{F0: 255}
                    _ = v52
                }
            }
            if (3.14159 >= 1.5) {
                if true {
                    if ((!true) || false) {
                        v12 = (fi48 & (^(fi48)))
                        switch ((("hello" + (("foo bar" + ("hello" + "foo bar")) + "a")) + (((("foo bar" + ("hello" + "\n\t")) + "hello") + "\n\t") + "")) + "hello") {
                        case "":
                            v53 := GenS4{F0: (v11 << ((481) & 63)), F1: 523}
                            _ = v53
                        }
                        var v54 float32 = 0.1
                        _ = v54
                    } else {
                        if true {
                            if len(v1) > 0 {
                                v1[0] = 686
                            }
                        } else {
                            var v55 int16 = 134
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if true {
                                for fi56 := 0; fi56 < 7; fi56++ {
                                    var v57 float64 = 1.5
                                    v58 := GenS4{F0: (v55 % ((v11) | 1)), F1: 157}
                                    var v59 uint64 = 2
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                                v60 := GenS1{F0: genFn0(true)}
                                fmt.Println(fi48)
                                _ = v60
                            }
                            _ = v55
                        }
                        var v61 int8 = 42
                        _ = v61
                    }
                    fmt.Println(v10, v10)
                    var v62 int16 = 0
                    _ = v62
                }
                fmt.Println(v12)
            }
        }
        if (!(861 == 18446744073709551615)) {
            for fi63 := 0; fi63 < 17; fi63++ {
                delete(v4, "a")
                var v64 string = "a"
                var v65 bool = ((v11 << ((2) & 63)) < v11)
                var v66 int8 = 10
                _ = fi63
                _ = v64
                _ = v65
                _ = v66
            }
            v13 += v13
        } else {
            for fi67 := 0; fi67 < 10; fi67++ {
                var v68 uint8 = 255
                _ = fi67
                _ = v68
            }
        }
        _ = fi48
    }
    var v69 float32 = ((0.1 / (v3 / ((((v13 / 3.14159) / v13) * (float32((-2.25 + float64(v13))) * v3)) / ((v3 * ((float32((0.1 + -2.25)) * v13) * v3)) * float32(1.5))))) + float32(0.0))
    v70 := GenS3{F0: 0}
    v71 := GenS2{F0: v10, F1: 50, F2: 939}
    v72 := GenS5{F0: 1, F1: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v71
    _ = v72
    _ = v9
}
