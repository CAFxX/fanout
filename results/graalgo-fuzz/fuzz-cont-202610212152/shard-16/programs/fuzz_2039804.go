package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 bool; F1 uint; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 uint64; F2 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 int8 }
type GenS8 struct { F0 uint32; F1 int64; F2 int16 }
type GenS9 struct { F0 string; F1 uint; F2 int64 }
func genFn0(p0 uint32, p1 uint64, p2 int8) string {
    fmt.Println(p2, p0, p1)
    p2 = (p2 << ((2) & 63))
    for fi6 := 0; fi6 < 4; fi6++ {
        if false {
            var v7 float64 = 1.5
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := []int{fi6}
                v7 += 1e10
                var v10 int32 = 0
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}

func main() {
    var v1 float64 = 0.1
    var v2 uint = 255
    v3 := []uint32{2147483647}
    var v4 float64 = -2.25
    v5 := []string{(("\n\t" + "\n\t") + ("hello" + "\n\t")), "", "a"}
    switch "" {
    case "foo bar":
        if false {
            switch -71 {
            case -2147483648:
                switch 802 {
                case 0:
                    var v11 int16 = -1
                    _ = v11
                default:
                    switch "" {
                    case "hello":
                        v1 -= (((v1 / 1e10) + 0.0) + float64(((3.14159 - 1e10) - 0.1)))
                    default:
                        var v12 int = -995
                        _ = v12
                    }
                }
            case 746:
                switch "\n\t" {
                case "hello":
                    if ((v4 <= (0.0 / (3.14159 - 0.1))) && ((((1 < 1) || (!(196 < 1))) && (2 <= 714)) && (!(false || true)))) {
                        if (!true) {
                            for fi13 := 0; fi13 < 14; fi13++ {
                                fmt.Println(v2, v1)
                                if (32767 < 32767) {
                                    if false {
                                        v3 = append(v3, 562)
                                        switch "" {
                                        case "foo bar":
                                            for fi14 := 0; fi14 < 4; fi14++ {
                                                fmt.Println(fi13, fi14, v2)
                                                var v15 int = (^(fi14))
                                                v16 := GenS1{F0: 1, F1: 255, F2: 2}
                                                _ = fi14
                                                _ = v15
                                                _ = v16
                                            }
                                        case "hello":
                                            for fi17 := 0; fi17 < 10; fi17++ {
                                                var v18 int16 = 132
                                                _ = fi17
                                                _ = v18
                                            }
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = 2
                                            }
                                        }
                                        fmt.Println(v2)
                                    }
                                } else {
                                    v4 -= (1e10 / v4)
                                }
                                var v19 int8 = -128
                                _ = fi13
                                _ = v19
                            }
                        } else {
                            if (!false) {
                                var v20 int8 = 62
                                _ = v20
                            }
                            if (true || (!false)) {
                                var v21 uint8 = 1
                                fmt.Println(v21)
                                v2 = 370
                                _ = v21
                            }
                            v1 += v4
                        }
                    } else {
                        var v22 bool = true
                        var v23 uint32 = 4294967295
                        v4 = 0.0
                        _ = v22
                        _ = v23
                    }
                }
            default:
                if true {
                    switch "foo bar" {
                    case "a":
                        if (v1 >= 1e10) {
                            fmt.Println(v2, v1, v1)
                        }
                    case "foo bar":
                        var v24 float32 = (1e10 * (float32(1.5) - 0.1))
                        _ = v24
                    case "\n\t":
                        v25 := GenS1{F0: 255, F1: 32767, F2: 107}
                        _ = v25
                    default:
                        v3 = append(v3, 246)
                    }
                    fmt.Println(v4, v1)
                } else {
                    fmt.Println(v2)
                    fmt.Println(v1, v1)
                    v26 := GenS1{F0: 1, F1: 32767, F2: 4294967295}
                    _ = v26
                }
            }
            fmt.Println(v1, v1)
            switch "a" {
            case "a":
                var v27 int8 = -128
                _ = v27
            case "hello":
                var v28 int16 = 0
                _ = v28
            default:
                fmt.Println(v2)
            }
        } else {
            v29 := GenS1{F0: 67, F1: 65535, F2: 0}
            _ = v29
        }
    case "":
        var v30 float64 = v1
        _ = v30
    case "a":
        var v31 int = 4294967295
        _ = v31
    default:
        v2 -= (v2 << ((65535) & 63))
    }
    v32 := GenS1{F0: 2, F1: 2, F2: 18446744073709551615}
    var v33 int64 = 1
    switch "foo bar" {
    case "foo bar":
        var v34 int8 = 42
        _ = v34
    }
    if ((100 >= 52) || (352 != 2147483647)) {
        if false {
            v33 -= -269
            v33 -= 42
            if true {
                fmt.Println(v33, v2, v4)
                v4 = 0.0
            } else {
                v2 = (v2 & v2)
                for fi35 := 0; fi35 < 15; fi35++ {
                    v36 := GenS2{F0: (^(fi35)), F1: (^(v2))}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (true && true) {
                        v1 = v4
                        var v37 float32 = 0.0
                        var v38 float64 = (v1 / 1.0)
                        _ = v37
                        _ = v38
                    } else {
                        v39 := GenS3{F0: (1.5 < (1e10 - 3.14159)), F1: v2, F2: 42}
                        _ = v39
                    }
                    _ = fi35
                    _ = v36
                }
            }
        } else {
            v4 -= float64((0.0 - -2.25))
            var v40 uint = (v2 & (^(v2)))
            v41 := GenS1{F0: 1, F1: 32767, F2: 18446744073709551615}
            _ = v40
            _ = v41
        }
    } else {
        v33 += (v33 & (v33 + v33))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (true && ((float64(1.5) / ((1e10 + v1) + 3.14159)) <= v4)) {
            var v42 int8 = 112
            v43 := GenS4{F0: 1}
            for fi44 := 0; fi44 < 3; fi44++ {
                v45 := GenS4{F0: -32768}
                for fi46 := 0; fi46 < 5; fi46++ {
                    fmt.Println(v42, v33)
                    v47 := GenS5{F0: (v2 >> ((1) & 63)), F1: 4294967295, F2: "a"}
                    _ = fi46
                    _ = v47
                }
                _ = fi44
                _ = v45
            }
            _ = v42
            _ = v43
        }
    }
    v2 -= v2
    switch "foo bar" {
    case "a":
        switch 128 {
        case 0:
            var v48 uint32 = 65535
            _ = v48
        case 9223372036854775807:
            v1 = ((0.1 / (float64(1e10) - ((float64(3.14159) + v1) + v4))) - (v1 * (v4 - 0.0)))
        default:
            v49 := GenS6{F0: ""}
            _ = v49
        }
    case "hello":
        var v50 uint = 65535
        _ = v50
    case "":
        for fi51 := 0; fi51 < 20; fi51++ {
            fmt.Println(v2)
            var v52 float64 = (v4 + float64(0.1))
            v53 := GenS7{F0: ("hello" + ""), F1: 100}
            _ = fi51
            _ = v52
            _ = v53
        }
    default:
        for fi54 := 0; fi54 < 3; fi54++ {
            v55 := GenS8{F0: 65535, F1: v33, F2: -1}
            v33 -= -9223372036854775808
            _ = fi54
            _ = v55
        }
    }
    var v56 float32 = 0.1
    if true {
        v4 -= float64((1.5 + ((v56 - v56) - 3.14159)))
    }
    if (2 > 0) {
        if true {
            _ = len(v5)
        }
        v4 = v4
        v57 := GenS9{F0: ("foo bar" + "hello"), F1: v2, F2: v33}
        _ = v57
    }
    v1 = v1
    v33 = -1
    v3 = append(v3, 466)
    var v58 int16 = -388
    switch -562 {
    case 9223372036854775807:
        switch (("foo bar" + "hello") + (("foo bar" + ("a" + "\n\t")) + "\n\t")) {
        case "hello":
            if ((!true) && (!true)) {
                var v59 string = ""
                _ = v59
            }
        case "a":
            var v60 bool = (1 != -128)
            _ = v60
        case "":
            var v61 float64 = v4
            _ = v61
        }
    case 32768:
        if (true && (!true)) {
            for fi62 := 0; fi62 < 13; fi62++ {
                fmt.Println(v2)
                _ = fi62
            }
            var v63 uint32 = 4294967295
            _ = v63
        }
    }
    var v64 uint8 = 0
    if ("" == "hello") {
        if (42 <= 0) {
            var v65 int8 = 1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v65
        }
    } else {
        v2 -= v2
        v33 = (-(v33))
        var v66 uint32 = 4294967295
        _ = v66
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v56
    _ = v58
    _ = v64
}
